package pf

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
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

func TestReadInputPreservesBinaryV2Payload(t *testing.T) {
	payload := []byte{0, 'a', '\n', '\r', 255}
	input := binaryV2InputFrame("1:2:3@topic", payload)

	protocol, msgID, got, err := readInput(bufioReader(input))
	if err != nil {
		t.Fatalf("readInput returned error: %v", err)
	}
	if protocol != childProtocolBinaryV2 {
		t.Fatalf("protocol = %v, want %v", protocol, childProtocolBinaryV2)
	}
	if msgID != "1:2:3" {
		t.Fatalf("msgID = %q, want %q", msgID, "1:2:3")
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("payload = %v, want %v", got, payload)
	}
}

func TestWriteResultPreservesBinaryV2Payload(t *testing.T) {
	payload := []byte{0, 'a', '\n', '\r', 255}
	var out bytes.Buffer

	writeResult(&out, childProtocolBinaryV2, payload, nil)

	status, got, err := readBinaryV2OutputFrame(&out)
	if err != nil {
		t.Fatalf("readBinaryV2OutputFrame returned error: %v", err)
	}
	if status != binaryV2StatusOK {
		t.Fatalf("status = %d, want %d", status, binaryV2StatusOK)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("payload = %v, want %v", got, payload)
	}
}

func TestWriteResultUsesBinaryV2EmptyAndErrorStatuses(t *testing.T) {
	var emptyOut bytes.Buffer
	writeResult(&emptyOut, childProtocolBinaryV2, nil, nil)
	status, body, err := readBinaryV2OutputFrame(&emptyOut)
	if err != nil {
		t.Fatalf("readBinaryV2OutputFrame returned error: %v", err)
	}
	if status != binaryV2StatusEmpty {
		t.Fatalf("status = %d, want %d", status, binaryV2StatusEmpty)
	}
	if len(body) != 0 {
		t.Fatalf("empty body length = %d, want 0", len(body))
	}

	var errorOut bytes.Buffer
	writeResult(&errorOut, childProtocolBinaryV2, nil, io.ErrUnexpectedEOF)
	status, body, err = readBinaryV2OutputFrame(&errorOut)
	if err != nil {
		t.Fatalf("readBinaryV2OutputFrame returned error: %v", err)
	}
	if status != binaryV2StatusError {
		t.Fatalf("status = %d, want %d", status, binaryV2StatusError)
	}
	if string(body) != io.ErrUnexpectedEOF.Error() {
		t.Fatalf("error body = %q, want %q", body, io.ErrUnexpectedEOF.Error())
	}
}

func binaryV2InputFrame(metadata string, payload []byte) []byte {
	frame := append([]byte{}, binaryV2InputMagic[:]...)
	frame = binary.BigEndian.AppendUint32(frame, uint32(len(metadata)))
	frame = binary.BigEndian.AppendUint64(frame, uint64(len(payload)))
	frame = append(frame, []byte(metadata)...)
	frame = append(frame, payload...)
	return frame
}

func readBinaryV2OutputFrame(reader io.Reader) (byte, []byte, error) {
	var magic [4]byte
	if _, err := io.ReadFull(reader, magic[:]); err != nil {
		return 0, nil, err
	}
	var status [1]byte
	if _, err := io.ReadFull(reader, status[:]); err != nil {
		return 0, nil, err
	}
	var bodyLen [8]byte
	if _, err := io.ReadFull(reader, bodyLen[:]); err != nil {
		return 0, nil, err
	}
	body := make([]byte, binary.BigEndian.Uint64(bodyLen[:]))
	if _, err := io.ReadFull(reader, body); err != nil {
		return 0, nil, err
	}
	return status[0], body, nil
}

func bufioReader(data []byte) *bufio.Reader {
	return bufio.NewReader(bytes.NewReader(data))
}
