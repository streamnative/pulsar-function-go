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
		SchemaType: "json",
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

func TestWriteSinkSchemaSidecarToDirSkipsBytesSchema(t *testing.T) {
	workDir := t.TempDir()

	sidecarPath, err := writeSinkSchemaSidecarToDir(workDir, SinkSchema{SchemaType: "bytes"})
	if err != nil {
		t.Fatalf("writeSinkSchemaSidecarToDir returned error: %v", err)
	}
	if sidecarPath != "" {
		t.Fatalf("sidecarPath = %q, want empty", sidecarPath)
	}
	if _, err := os.Stat(filepath.Join(workDir, SinkSchemaSidecarFileName)); !os.IsNotExist(err) {
		t.Fatalf("sidecar stat error = %v, want not exist", err)
	}
}

func TestWriteSinkSchemaSidecarToDirWritesEmptyPropertiesObject(t *testing.T) {
	workDir := t.TempDir()

	sidecarPath, err := writeSinkSchemaSidecarToDir(workDir, SinkSchema{
		SchemaType: "string",
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
