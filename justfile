# SPDX-FileCopyrightText: 2023 Siemens AG
#
# SPDX-License-Identifier: Apache-2.0
#
# Author: Michael Adler <michael.adler@siemens.com>

THISDIR := justfile_directory()
DOCKER := env_var_or_default("DOCKER", "docker")
NIX := "nix --experimental-features 'nix-command flakes'"

# postgres

export PGUSER := env_var_or_default("PGUSER", "wfx")
export PGPASSWORD := env_var_or_default("PGPASSWORD", "secret")
export PGHOST := env_var_or_default("PGHOST", "localhost")
export PGPORT := env_var_or_default("PGPORT", "5432")
export PGDATABASE := env_var_or_default("PGDATABASE", "wfx")
export PGSSLMODE := env_var_or_default("PGSSLMODE", "disable")

# mysql

export MYSQL_USER := env_var_or_default("MYSQL_USER", "root")
export MYSQL_PASSWORD := env_var_or_default("MYSQL_PASSWORD", "root")
export MYSQL_ROOT_PASSWORD := env_var_or_default("MYSQL_PASSWORD", "root")
export MYSQL_DATABASE := env_var_or_default("MYSQL_DATABASE", "wfx")
export MYSQL_HOST := env_var_or_default("MYSQL_HOST", "localhost")

build:
    #!/usr/bin/env bash
    set -euo pipefail
    # goreleaser requires an absolute path to the compiler
    /usr/bin/env CC={{ THISDIR }}/.ci/zcc goreleaser build --clean --single-target --snapshot
    go build -C example/plugin
    go build -C contrib/remote-access/client
    go build -C contrib/config-deployment/client

# Update dependencies
update-deps:
    #!/usr/bin/env bash
    set -euxo pipefail
    go get -u ./...
    go mod tidy

# Build the documentation
pages:
    #!/usr/bin/env bash
    set -euo pipefail
    rm -rf public hugo/public
    pushd hugo
    make clean && make -j`nproc`
    npm install postcss postcss-cli autoprefixer
    hugo --minify
    popd
    mv hugo/public .
    go tool github.com/wjdp/htmltest
    lychee public

# Serve docs
docs-serve:
    #!/bin/sh
    cd hugo && make clean && make -j`nproc` && hugo server -D

# Lint code
lint:
    #!/usr/bin/env bash
    set -euo pipefail
    export CGO_ENABLED=0
    golangci-lint run -v --build-tags=testing
    staticcheck -tags=testing ./...
    go list ./... 2>/dev/null | sed -e 's,github.com/siemens/wfx/,,' | grep -v "^generated" | sort | uniq | while read -r pkg; do
        if [[ "$pkg" == *tests* ]] || [[ ! -d "$pkg" ]]; then
            continue
        fi
        file_count=$(find "$pkg" -maxdepth 1 -type f -name "*.go" | wc -l)
        test_count=$(find "$pkg" -maxdepth 1 -type f -name "*_test.go" | wc -l)
        if [[ "$file_count" -gt 0 ]] && [[ "$test_count" -eq 0 ]]; then
            echo "WARN: package $pkg has $file_count file(s) but no tests"
        fi
        if [[ "$test_count" -gt 0 ]]; then
            grep -R -q goleak.VerifyTestMain "$pkg" || {
                echo "ERROR: package $pkg does not use goleak"
                exit 1
            }
        fi
    done

# Format code
format:
    #!/usr/bin/env bash
    set -eux
    go tool mvdan.cc/gofumpt -l -w .
    prettier -l -w .
    just --fmt --unstable

_generate-openapi:
    #!/bin/sh
    set -eu
    cd "{{ THISDIR }}/spec"
    go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml wfx.openapi.yml

_generate-ent:
    #!/usr/bin/env bash
    set -euxo pipefail
    cd "{{ THISDIR }}/generated/ent"
    find . -not -name generate.go -and -not -name main.go -and -not -path "**/schema/*" -type f -delete
    go tool entgo.io/ent/cmd/ent generate --header \
        "// SPDX-FileCopyrightText: The entgo authors
         // SPDX-License-Identifier: Apache-2.0

         // Code generated by ent, DO NOT EDIT." --feature sql/execquery,sql/versioned-migration ./schema

_generate-mockery:
    go tool github.com/vektra/mockery/v2 --all

_generate-flatbuffers:
    #!/usr/bin/env bash
    set -euo pipefail
    rm -rf generated/plugin
    find fbs -name "*.fbs" | xargs flatc -g --gen-object-api --go-module-name github.com/siemens/wfx
    go tool mvdan.cc/gofumpt -l -w generated/plugin

# Generate code
generate: _generate-openapi _generate-ent _generate-mockery _generate-flatbuffers

# Start PostgreSQL container
postgres-start:
    #!/usr/bin/env bash
    count=`{{ DOCKER }} container ls --quiet --noheading --filter name=wfx-postgres --filter "health=healthy" | wc -l`
    image=`grep -m1 -o "image: postgres:.*" .github/workflows/ci.yml | sed -e 's/image:\s*//'`
    if [[ $count -eq 0 ]]; then
        echo "Starting $image"
        {{ DOCKER }} run -d --rm \
            --name wfx-postgres \
            -e "POSTGRES_USER=$PGUSER" \
            -e "POSTGRES_PASSWORD=$PGPASSWORD" \
            -e "POSTGRES_DB=$PGDATABASE" \
            -p 5432:5432 \
            --health-cmd pg_isready \
            --health-interval 3s \
            --health-timeout 5s \
            --health-retries 20 \
            "docker.io/library/$image"
    else
        echo "PostgreSQL is already running"
    fi
    while [[ $count -eq 0 ]]; do
        count=`{{ DOCKER }} container ls --quiet --noheading --filter name=wfx-postgres --filter "health=healthy" | wc -l`
        echo "Waiting for PostgreSQL to come up..."
        sleep 3
    done

# View PostgreSQL logs
@postgres-logs:
    {{ DOCKER }} logs --color -f wfx-postgres

# Enter PostgreSQL shell inside container
@postgres-shell:
    {{ DOCKER }} exec -it wfx-postgres psql -d $PGDATABASE -U $PGUSER

# Stop PostgreSQL container
postgres-stop: (_container-stop "wfx-postgres")

# Start wfx and connect to Postgres database
@postgres-wfx: postgres-start
    ./wfx --log-level debug \
        --storage postgres \
        --storage-opt "host=$PGHOST port=5432 database=$PGDATABASE user=$PGUSER password=$PGPASSWORD sslmode=disable"

# Generate schema definitions for PostgreSQL
postgres-generate-schema name:
    go run -mod=mod generated/ent/migrate/main.go postgres "{{ name }}"

# Stop a container with the given name.
_container-stop name:
    #!/usr/bin/env bash
    count=`{{ DOCKER }} ps --filter name={{ name }} --quiet --noheading | wc -l`
    if [[ $count -gt 0 ]]; then
        echo "Stopping {{ name }}"
        {{ DOCKER }} stop "{{ name }}" 1>/dev/null 2>/dev/null
        {{ DOCKER }} rm "{{ name }}" 2>/dev/null || true
    else
        echo "{{ name }} is already stopped"
    fi

# Start MySQL container
mysql-start:
    #!/usr/bin/env bash
    count=`{{ DOCKER }} container ls --quiet --noheading --filter name=wfx-mysql --filter "health=healthy" | wc -l`
    image=`grep -m1 -o "image: mysql:.*" .github/workflows/ci.yml | sed -e 's/image:\s*//'`
    if [[ $count -eq 0 ]]; then
        echo "Starting $image"
        {{ DOCKER }} run -d --rm \
            --name wfx-mysql \
            -e MYSQL_DATABASE \
            -e MYSQL_ROOT_PASSWORD \
            -p 3306:3306 \
            --health-cmd 'mysqladmin ping --silent' \
            --health-interval 3s \
            --health-timeout 5s \
            --health-retries 20 \
            "docker.io/library/$image"
    else
        echo "MySQL is already running"
    fi
    while [[ $count -eq 0 ]]; do
        count=`{{ DOCKER }} container ls --quiet --noheading --filter name=wfx-mysql --filter "health=healthy" | wc -l`
        echo "Waiting for MySQL to come up..."
        sleep 3
    done

# Stop MySQL container
mysql-stop: (_container-stop "wfx-mysql")

# View MySQL logs
@mysql-logs:
    {{ DOCKER }} logs --color -f wfx-mysql

# Enter MySQL shell
mysql-shell:
    {{ DOCKER }} exec -it wfx-mysql mysql -u$MYSQL_USER -p$MYSQL_PASSWORD -D $MYSQL_DATABASE

# Generate schema definitions for MySQL
mysql-generate-schema name:
    go run -mod=mod generated/ent/migrate/main.go mysql "{{ name }}"

# Start wfx and connect to MySQL container.
mysql-wfx: mysql-start
    ./wfx --log-level debug \
        --storage mysql \
        --storage-opt "$MYSQL_USER:$MYSQL_PASSWORD@/$MYSQL_DATABASE"

# Generate schema definitions for SQLite
sqlite-generate-schema name:
    go run -mod=mod generated/ent/migrate/main.go sqlite "{{ name }}"

# Check links used in Markdown files.
check-md-links:
    .ci/check-links.py

# vim: ts=4 sw=4 expandtab
