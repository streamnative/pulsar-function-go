package pf

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestRecordEventTimestampIsInt64(t *testing.T) {
	record := &Record{EventTimestamp: 5}
	var timestamp int64 = record.GetEventTimestamp()

	if timestamp != 5 {
		t.Fatalf("timestamp = %d, want 5", timestamp)
	}
}

func TestSetMessageIdClearsCachedRecord(t *testing.T) {
	functionContext := &FunctionContext{
		message: &Record{MessageId: "previous"},
	}

	functionContext.setMessageId(&MessageId{Id: "next"})

	if functionContext.message != nil {
		t.Fatalf("cached message = %+v, want nil", functionContext.message)
	}
}

func TestPublishSendsRawPayloadWithoutSchema(t *testing.T) {
	stub := &publishCaptureStub{}
	functionContext := &FunctionContext{
		ctx:  context.Background(),
		stub: stub,
	}

	messageID, err := functionContext.Publish("persistent://public/default/raw", []byte("payload"))
	if err != nil {
		t.Fatalf("Publish returned error: %v", err)
	}

	if messageID.GetEntryId() != 7 {
		t.Fatalf("entryId = %d, want 7", messageID.GetEntryId())
	}
	if stub.message == nil {
		t.Fatal("published message is nil")
	}
	if stub.message.GetTopic() != "persistent://public/default/raw" {
		t.Fatalf("topic = %q, want raw topic", stub.message.GetTopic())
	}
	if string(stub.message.GetPayload()) != "payload" {
		t.Fatalf("payload = %q, want payload", stub.message.GetPayload())
	}
	if stub.message.GetSchema() != nil {
		t.Fatalf("schema = %+v, want nil", stub.message.GetSchema())
	}
}

func TestPublishWithSchemaSendsSchemaMetadata(t *testing.T) {
	stub := &publishCaptureStub{}
	functionContext := &FunctionContext{
		ctx:  context.Background(),
		stub: stub,
	}
	schema := PublishMessageSchema{
		SchemaType: SchemaTypeAvro,
		Name:       "example.Event",
		SchemaData: `{"type":"record","name":"Event","fields":[]}`,
		Properties: map[string]string{
			"encoding": "binary",
		},
		Subject: "events-value",
	}

	_, err := functionContext.PublishWithSchema(
		"persistent://public/default/events",
		[]byte("encoded"),
		schema,
	)
	if err != nil {
		t.Fatalf("PublishWithSchema returned error: %v", err)
	}

	got := stub.message.GetSchema()
	if got == nil {
		t.Fatal("schema is nil, want schema metadata")
	}
	if got.GetSchemaType() != SchemaTypeAvro {
		t.Fatalf("schemaType = %q, want %q", got.GetSchemaType(), SchemaTypeAvro)
	}
	if got.GetName() != "example.Event" {
		t.Fatalf("schema name = %q, want example.Event", got.GetName())
	}
	if string(got.GetSchemaData()) != `{"type":"record","name":"Event","fields":[]}` {
		t.Fatalf("schemaData = %q, want Avro definition", got.GetSchemaData())
	}
	if got.GetProperties()["encoding"] != "binary" {
		t.Fatalf("schema properties = %+v, want encoding=binary", got.GetProperties())
	}
	if got.GetSubject() != "events-value" {
		t.Fatalf("subject = %q, want events-value", got.GetSubject())
	}
	if string(stub.message.GetPayload()) != "encoded" {
		t.Fatalf("payload = %q, want encoded", stub.message.GetPayload())
	}
}

type publishCaptureStub struct {
	message *PulsarMessage
}

func (s *publishCaptureStub) Publish(_ context.Context, message *PulsarMessage, _ ...grpc.CallOption) (*SendMessageId, error) {
	s.message = message
	return &SendMessageId{EntryId: 7}, nil
}

func (s *publishCaptureStub) CurrentRecord(context.Context, *MessageId, ...grpc.CallOption) (*Record, error) {
	return nil, nil
}

func (s *publishCaptureStub) RecordMetrics(context.Context, *MetricData, ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *publishCaptureStub) Seek(context.Context, *Partition, ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *publishCaptureStub) Pause(context.Context, *Partition, ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *publishCaptureStub) Resume(context.Context, *Partition, ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *publishCaptureStub) GetState(context.Context, *StateKey, ...grpc.CallOption) (*StateResult, error) {
	return nil, nil
}

func (s *publishCaptureStub) PutState(context.Context, *StateKeyValue, ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *publishCaptureStub) DeleteState(context.Context, *StateKey, ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *publishCaptureStub) GetCounter(context.Context, *StateKey, ...grpc.CallOption) (*Counter, error) {
	return nil, nil
}

func (s *publishCaptureStub) IncrCounter(context.Context, *IncrStateKey, ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
