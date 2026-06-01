package pf

import (
	"encoding/binary"
	"strings"
	"testing"
)

func TestReadInputFrameParsesOneByteMetadataLength(t *testing.T) {
	frame := append([]byte{byte(len("1:2:3@topic"))}, []byte("1:2:3@topicpayload\n")...)

	msgID, payload, err := readInputFrame(frame)
	if err != nil {
		t.Fatalf("readInputFrame returned error: %v", err)
	}
	if msgID != "1:2:3" {
		t.Fatalf("msgID = %q, want %q", msgID, "1:2:3")
	}
	if string(payload) != "payload" {
		t.Fatalf("payload = %q, want %q", payload, "payload")
	}
}

func TestReadInputFrameParsesExtendedMetadataLength(t *testing.T) {
	metadata := "1:2:3@" + strings.Repeat("a", 300)
	frame := []byte{extendedStdinMetadataMarker, 0, 0, 0, 0}
	binary.BigEndian.PutUint32(frame[1:5], uint32(len(metadata)))
	frame = append(frame, []byte(metadata+"payload\n")...)

	msgID, payload, err := readInputFrame(frame)
	if err != nil {
		t.Fatalf("readInputFrame returned error: %v", err)
	}
	if msgID != "1:2:3" {
		t.Fatalf("msgID = %q, want %q", msgID, "1:2:3")
	}
	if string(payload) != "payload" {
		t.Fatalf("payload = %q, want %q", payload, "payload")
	}
}

func TestReadInputFrameRejectsTruncatedExtendedMetadataLength(t *testing.T) {
	_, _, err := readInputFrame([]byte{extendedStdinMetadataMarker, 0, 0, 1})
	if err == nil {
		t.Fatal("readInputFrame returned nil error, want truncated extended metadata error")
	}
}

func TestReadInputFrameRejectsMalformedMetadata(t *testing.T) {
	frame := append([]byte{byte(len("missing-topic"))}, []byte("missing-topicpayload\n")...)

	_, _, err := readInputFrame(frame)
	if err == nil {
		t.Fatal("readInputFrame returned nil error, want malformed metadata error")
	}
	if err.Error() != "invalid metadata format: expected message id and topic separated by @" {
		t.Fatalf("error = %q, want malformed metadata error", err.Error())
	}
}
