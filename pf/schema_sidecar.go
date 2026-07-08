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
	"strings"
)

const (
	// SinkSchemaSidecarFileName is the schema sidecar file consumed by the generic runtime.
	SinkSchemaSidecarFileName = "sink.schema.json"

	// SchemaTypeNone disables the sink schema sidecar.
	SchemaTypeNone = "none"
	// SchemaTypeBytes disables the sink schema sidecar.
	SchemaTypeBytes     = "bytes"
	SchemaTypeJSON      = "json"
	SchemaTypeAvro      = "avro"
	SchemaTypeString    = "string"
	SchemaTypeBool      = "boolean"
	SchemaTypeInt8      = "int8"
	SchemaTypeInt16     = "int16"
	SchemaTypeInt32     = "int32"
	SchemaTypeInt64     = "int64"
	SchemaTypeFloat     = "float"
	SchemaTypeDouble    = "double"
	SchemaTypeDate      = "date"
	SchemaTypeTime      = "time"
	SchemaTypeTimestamp = "timestamp"
)

// SinkSchema describes the sink schema definition written to the runtime sidecar file.
// SchemaType should use one of the SchemaType* constants for JSON, Avro, or
// primitive schemas. Empty, none, and bytes schema types do not need sidecars.
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
	sidecarPath := filepath.Join(workDir, SinkSchemaSidecarFileName)
	if isNoSchemaType(schemaType) {
		if err := os.Remove(sidecarPath); err != nil && !os.IsNotExist(err) {
			return "", err
		}
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

	if err := writeFileAtomic(sidecarPath, content, 0644); err != nil {
		return "", err
	}
	return sidecarPath, nil
}

func writeFileAtomic(path string, content []byte, perm os.FileMode) error {
	tmpFile, err := os.CreateTemp(filepath.Dir(path), filepath.Base(path)+".*.tmp")
	if err != nil {
		return err
	}
	tmpPath := tmpFile.Name()
	defer func() {
		_ = os.Remove(tmpPath)
	}()

	_, writeErr := tmpFile.Write(content)
	chmodErr := tmpFile.Chmod(perm)
	closeErr := tmpFile.Close()
	if writeErr != nil {
		return writeErr
	}
	if chmodErr != nil {
		return chmodErr
	}
	if closeErr != nil {
		return closeErr
	}

	return os.Rename(tmpPath, path)
}

func isNoSchemaType(schemaType string) bool {
	switch strings.ToLower(strings.TrimSpace(schemaType)) {
	case "", SchemaTypeBytes, SchemaTypeNone:
		return true
	default:
		return false
	}
}
