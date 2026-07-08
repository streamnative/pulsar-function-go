//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package pf

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteSinkSchemaSidecarToDirWritesCommonSidecar(t *testing.T) {
	workDir := t.TempDir()
	schema := SinkSchema{
		SchemaType: SchemaTypeJSON,
		Name:       "example.Student",
		SchemaData: `{"type":"record","name":"Student","fields":[]}`,
		Properties: map[string]string{
			"custom": "value",
		},
	}

	sidecarPath, err := writeSinkSchemaSidecarToDir(workDir, schema)
	if err != nil {
		t.Fatalf("writeSinkSchemaSidecarToDir returned error: %v", err)
	}

	if sidecarPath != filepath.Join(workDir, SinkSchemaSidecarFileName) {
		t.Fatalf("sidecarPath = %q, want %q", sidecarPath, filepath.Join(workDir, SinkSchemaSidecarFileName))
	}

	content, err := os.ReadFile(sidecarPath)
	if err != nil {
		t.Fatalf("read sidecar: %v", err)
	}
	var payload map[string]interface{}
	if err := json.Unmarshal(content, &payload); err != nil {
		t.Fatalf("unmarshal sidecar: %v", err)
	}

	if payload["schemaType"] != "json" {
		t.Fatalf("schemaType = %v, want json", payload["schemaType"])
	}
	if payload["name"] != "example.Student" {
		t.Fatalf("name = %v, want example.Student", payload["name"])
	}
	if payload["schemaData"] != `{"type":"record","name":"Student","fields":[]}` {
		t.Fatalf("schemaData = %v, want record schema", payload["schemaData"])
	}
	properties, ok := payload["properties"].(map[string]interface{})
	if !ok {
		t.Fatalf("properties = %T, want object", payload["properties"])
	}
	if properties["custom"] != "value" {
		t.Fatalf("properties[custom] = %v, want value", properties["custom"])
	}
}

func TestWriteSinkSchemaSidecarToDirSkipsNoSchemaTypesAndRemovesStaleSidecar(t *testing.T) {
	tests := []struct {
		name       string
		schemaType string
	}{
		{name: "empty", schemaType: ""},
		{name: "bytes", schemaType: SchemaTypeBytes},
		{name: "uppercase bytes", schemaType: "BYTES"},
		{name: "none", schemaType: SchemaTypeNone},
		{name: "uppercase none", schemaType: "NONE"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workDir := t.TempDir()
			sidecarFile := filepath.Join(workDir, SinkSchemaSidecarFileName)
			if err := os.WriteFile(sidecarFile, []byte("stale"), 0644); err != nil {
				t.Fatalf("write stale sidecar: %v", err)
			}

			sidecarPath, err := writeSinkSchemaSidecarToDir(workDir, SinkSchema{SchemaType: tt.schemaType})
			if err != nil {
				t.Fatalf("writeSinkSchemaSidecarToDir returned error: %v", err)
			}
			if sidecarPath != "" {
				t.Fatalf("sidecarPath = %q, want empty", sidecarPath)
			}
			if _, err := os.Stat(sidecarFile); !os.IsNotExist(err) {
				t.Fatalf("sidecar stat error = %v, want not exist", err)
			}
		})
	}
}

func TestWriteSinkSchemaSidecarToDirTrimsSchemaTypeAndPreservesCasing(t *testing.T) {
	workDir := t.TempDir()

	sidecarPath, err := writeSinkSchemaSidecarToDir(workDir, SinkSchema{
		SchemaType: " JSON ",
		SchemaData: `{"type":"record","name":"Student","fields":[]}`,
	})
	if err != nil {
		t.Fatalf("writeSinkSchemaSidecarToDir returned error: %v", err)
	}

	content, err := os.ReadFile(sidecarPath)
	if err != nil {
		t.Fatalf("read sidecar: %v", err)
	}
	var payload struct {
		SchemaType string `json:"schemaType"`
	}
	if err := json.Unmarshal(content, &payload); err != nil {
		t.Fatalf("unmarshal sidecar: %v", err)
	}
	if payload.SchemaType != "JSON" {
		t.Fatalf("schemaType = %q, want JSON", payload.SchemaType)
	}
}

func TestSchemaTypeBoolUsesPulsarBooleanName(t *testing.T) {
	if SchemaTypeBool != "boolean" {
		t.Fatalf("SchemaTypeBool = %q, want boolean", SchemaTypeBool)
	}
}

func TestWriteSinkSchemaSidecarToDirWritesEmptyPropertiesObject(t *testing.T) {
	workDir := t.TempDir()

	sidecarPath, err := writeSinkSchemaSidecarToDir(workDir, SinkSchema{
		SchemaType: SchemaTypeString,
		SchemaData: "",
	})
	if err != nil {
		t.Fatalf("writeSinkSchemaSidecarToDir returned error: %v", err)
	}

	content, err := os.ReadFile(sidecarPath)
	if err != nil {
		t.Fatalf("read sidecar: %v", err)
	}
	var payload struct {
		Properties map[string]string `json:"properties"`
	}
	if err := json.Unmarshal(content, &payload); err != nil {
		t.Fatalf("unmarshal sidecar: %v", err)
	}
	if payload.Properties == nil {
		t.Fatal("properties is nil, want empty object")
	}
	if len(payload.Properties) != 0 {
		t.Fatalf("properties length = %d, want 0", len(payload.Properties))
	}
}
