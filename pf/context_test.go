package pf

import "testing"

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
