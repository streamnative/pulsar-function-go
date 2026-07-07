package pf

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// SinkSchemaSidecarFileName is the schema sidecar file consumed by the generic runtime.
const SinkSchemaSidecarFileName = "sink.schema.json"

// SinkSchema describes the sink schema definition written to the runtime sidecar file.
type SinkSchema struct {
	SchemaType string
	Name       string
	SchemaData string
	Properties map[string]string
}

type sinkSchemaSidecarPayload struct {
	SchemaType string            `json:"schemaType"`
	Name       string            `json:"name"`
	SchemaData string            `json:"schemaData"`
	Properties map[string]string `json:"properties"`
}

// StartWithSinkSchema writes the sink schema sidecar before starting the function.
func StartWithSinkSchema(funcName interface{}, schema SinkSchema) {
	if _, err := WriteSinkSchemaSidecar(schema); err != nil {
		panic(err)
	}
	Start(funcName)
}

// WriteSinkSchemaSidecar writes sink.schema.json next to the current executable.
func WriteSinkSchemaSidecar(schema SinkSchema) (string, error) {
	executable, err := os.Executable()
	if err != nil {
		return "", err
	}
	return writeSinkSchemaSidecarToDir(filepath.Dir(executable), schema)
}

func writeSinkSchemaSidecarToDir(workDir string, schema SinkSchema) (string, error) {
	schemaType := strings.TrimSpace(schema.SchemaType)
	if schemaType == "" || strings.EqualFold(schemaType, "bytes") {
		return "", nil
	}

	properties := schema.Properties
	if properties == nil {
		properties = map[string]string{}
	}
	payload := sinkSchemaSidecarPayload{
		SchemaType: schemaType,
		Name:       schema.Name,
		SchemaData: schema.SchemaData,
		Properties: properties,
	}
	content, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	sidecarPath := filepath.Join(workDir, SinkSchemaSidecarFileName)
	if err := os.WriteFile(sidecarPath, content, 0644); err != nil {
		return "", err
	}
	return sidecarPath, nil
}
