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

package main

import (
	"context"
	"fmt"

	"github.com/linkedin/goavro/v2"
	"github.com/streamnative/pulsar-function-go/pf"
)

const studentAvroSchema = `{"type":"record","name":"Student","fields":[{"name":"name","type":["null","string"]},{"name":"age","type":["null","int"]},{"name":"grade","type":["null","int"]}]}`
const studentAvroInputFallbackSchema = `{"type":"record","name":"Student","fields":[{"name":"name","type":["null","string"]},{"name":"age","type":"int"},{"name":"grade","type":"int"}]}`

var studentCodec = mustAvroCodec(studentAvroSchema)
var studentInputFallbackCodec = mustAvroCodec(studentAvroInputFallbackSchema)

func mustAvroCodec(schema string) *goavro.Codec {
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		panic(err)
	}
	return codec
}

func HandleContextPublishAvro(ctx context.Context, in []byte) error {
	record, err := decodeStudent(in)
	if err != nil {
		return err
	}

	name, err := avroUnionString(record["name"])
	if err != nil {
		return err
	}
	age, err := avroUnionInt(record["age"])
	if err != nil {
		return err
	}
	grade, err := avroUnionInt(record["grade"])
	if err != nil {
		return err
	}

	outputRecord := map[string]interface{}{
		"name":  map[string]interface{}{"string": name},
		"age":   map[string]interface{}{"int": age},
		"grade": map[string]interface{}{"int": grade + 1},
	}
	output, err := studentCodec.BinaryFromNative(nil, outputRecord)
	if err != nil {
		return err
	}

	fc, ok := pf.FromContext(ctx)
	if !ok {
		return fmt.Errorf("missing function context")
	}

	publishTopic, ok := fc.GetUserConfValue("publishTopic").(string)
	if !ok || publishTopic == "" {
		return fmt.Errorf("missing publishTopic user config")
	}

	_, err = fc.PublishWithSchema(publishTopic, output, pf.PublishMessageSchema{
		SchemaType: pf.SchemaTypeAvro,
		Name:       "Student",
		SchemaData: studentAvroSchema,
	})
	return err
}

func decodeStudent(data []byte) (map[string]interface{}, error) {
	for _, codec := range []*goavro.Codec{studentCodec, studentInputFallbackCodec} {
		native, _, err := codec.NativeFromBinary(data)
		if err != nil {
			continue
		}
		record, ok := native.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected Avro record, got %T", native)
		}
		return record, nil
	}
	_, _, err := studentCodec.NativeFromBinary(data)
	return nil, err
}

func avroUnionString(value interface{}) (string, error) {
	switch typed := value.(type) {
	case map[string]interface{}:
		raw, ok := typed["string"]
		if !ok {
			return "", fmt.Errorf("expected string union branch, got %v", typed)
		}
		return avroUnionString(raw)
	case string:
		return typed, nil
	default:
		return "", fmt.Errorf("expected Avro string, got %T", value)
	}
}

func avroUnionInt(value interface{}) (int32, error) {
	switch typed := value.(type) {
	case map[string]interface{}:
		raw, ok := typed["int"]
		if !ok {
			return 0, fmt.Errorf("expected int union branch, got %v", typed)
		}
		return avroUnionInt(raw)
	case int32:
		return typed, nil
	case int:
		return int32(typed), nil
	case int64:
		return int32(typed), nil
	default:
		return 0, fmt.Errorf("expected Avro int, got %T", value)
	}
}

func main() {
	pf.Start(HandleContextPublishAvro)
}
