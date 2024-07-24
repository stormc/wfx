//go:build swagger

package spec

/*
 * SPDX-FileCopyrightText: 2024 Siemens AG
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Author: Michael Adler <michael.adler@siemens.com>
 */

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

//go:embed wfx.swagger.yml
var swaggerYAML string

func init() {
	var yamlObj map[string]any
	if err := yaml.Unmarshal([]byte(swaggerYAML), &yamlObj); err != nil {
		panic(err)
	}
	basePath := yamlObj["basePath"].(string)

	jsonData, err := json.Marshal(yamlObj)
	if err != nil {
		panic(err)
	}

	Handlers[fmt.Sprintf("GET %s/swagger.json", basePath)] = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(jsonData)
	})
}

func yamlToJSON(yamlData []byte) ([]byte, error) {
	var yamlObj any
	if err := yaml.Unmarshal(yamlData, &yamlObj); err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(yamlObj)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
