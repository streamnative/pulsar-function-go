//*
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.25.1
// source: context.proto

package pf

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SendMessageId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LedgerId       int64 `protobuf:"varint,1,opt,name=ledgerId,proto3" json:"ledgerId,omitempty"`
	EntryId        int64 `protobuf:"varint,2,opt,name=entryId,proto3" json:"entryId,omitempty"`
	BatchIndex     int32 `protobuf:"varint,3,opt,name=batchIndex,proto3" json:"batchIndex,omitempty"`
	PartitionIndex int32 `protobuf:"varint,4,opt,name=partitionIndex,proto3" json:"partitionIndex,omitempty"`
}

func (x *SendMessageId) Reset() {
	*x = SendMessageId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageId) ProtoMessage() {}

func (x *SendMessageId) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageId.ProtoReflect.Descriptor instead.
func (*SendMessageId) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageId) GetLedgerId() int64 {
	if x != nil {
		return x.LedgerId
	}
	return 0
}

func (x *SendMessageId) GetEntryId() int64 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *SendMessageId) GetBatchIndex() int32 {
	if x != nil {
		return x.BatchIndex
	}
	return 0
}

func (x *SendMessageId) GetPartitionIndex() int32 {
	if x != nil {
		return x.PartitionIndex
	}
	return 0
}

type StateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *StateKey) Reset() {
	*x = StateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateKey) ProtoMessage() {}

func (x *StateKey) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateKey.ProtoReflect.Descriptor instead.
func (*StateKey) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{1}
}

func (x *StateKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type StateKeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StateKeyValue) Reset() {
	*x = StateKeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateKeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateKeyValue) ProtoMessage() {}

func (x *StateKeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateKeyValue.ProtoReflect.Descriptor instead.
func (*StateKeyValue) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{2}
}

func (x *StateKeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StateKeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type StateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StateResult) Reset() {
	*x = StateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateResult) ProtoMessage() {}

func (x *StateResult) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateResult.ProtoReflect.Descriptor instead.
func (*StateResult) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{3}
}

func (x *StateResult) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type IncrStateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Amount int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *IncrStateKey) Reset() {
	*x = IncrStateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrStateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrStateKey) ProtoMessage() {}

func (x *IncrStateKey) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrStateKey.ProtoReflect.Descriptor instead.
func (*IncrStateKey) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{4}
}

func (x *IncrStateKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *IncrStateKey) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{5}
}

func (x *Counter) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{6}
}

func (x *Request) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type MessageId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MessageId) Reset() {
	*x = MessageId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageId) ProtoMessage() {}

func (x *MessageId) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageId.ProtoReflect.Descriptor instead.
func (*MessageId) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{7}
}

func (x *MessageId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PulsarMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic               string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Payload             []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	MessageId           string `protobuf:"bytes,3,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Properties          string `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
	PartitionKey        string `protobuf:"bytes,5,opt,name=partitionKey,proto3" json:"partitionKey,omitempty"`
	SequenceId          string `protobuf:"bytes,6,opt,name=sequenceId,proto3" json:"sequenceId,omitempty"`
	ReplicationClusters string `protobuf:"bytes,7,opt,name=replicationClusters,proto3" json:"replicationClusters,omitempty"`
	DisableReplication  bool   `protobuf:"varint,8,opt,name=disableReplication,proto3" json:"disableReplication,omitempty"`
	EventTimestamp      int32  `protobuf:"varint,9,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	DeliverAt           int32  `protobuf:"varint,10,opt,name=deliverAt,proto3" json:"deliverAt,omitempty"`
	DeliverAfter        int32  `protobuf:"varint,11,opt,name=deliverAfter,proto3" json:"deliverAfter,omitempty"`
}

func (x *PulsarMessage) Reset() {
	*x = PulsarMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PulsarMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PulsarMessage) ProtoMessage() {}

func (x *PulsarMessage) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PulsarMessage.ProtoReflect.Descriptor instead.
func (*PulsarMessage) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{8}
}

func (x *PulsarMessage) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PulsarMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *PulsarMessage) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *PulsarMessage) GetProperties() string {
	if x != nil {
		return x.Properties
	}
	return ""
}

func (x *PulsarMessage) GetPartitionKey() string {
	if x != nil {
		return x.PartitionKey
	}
	return ""
}

func (x *PulsarMessage) GetSequenceId() string {
	if x != nil {
		return x.SequenceId
	}
	return ""
}

func (x *PulsarMessage) GetReplicationClusters() string {
	if x != nil {
		return x.ReplicationClusters
	}
	return ""
}

func (x *PulsarMessage) GetDisableReplication() bool {
	if x != nil {
		return x.DisableReplication
	}
	return false
}

func (x *PulsarMessage) GetEventTimestamp() int32 {
	if x != nil {
		return x.EventTimestamp
	}
	return 0
}

func (x *PulsarMessage) GetDeliverAt() int32 {
	if x != nil {
		return x.DeliverAt
	}
	return 0
}

func (x *PulsarMessage) GetDeliverAfter() int32 {
	if x != nil {
		return x.DeliverAfter
	}
	return 0
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload        []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	MessageId      string `protobuf:"bytes,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Properties     string `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
	Key            string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	PartitionId    string `protobuf:"bytes,5,opt,name=partitionId,proto3" json:"partitionId,omitempty"`
	TopicName      string `protobuf:"bytes,6,opt,name=topicName,proto3" json:"topicName,omitempty"`
	EventTimestamp int32  `protobuf:"varint,7,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{9}
}

func (x *Record) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Record) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Record) GetProperties() string {
	if x != nil {
		return x.Properties
	}
	return ""
}

func (x *Record) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Record) GetPartitionId() string {
	if x != nil {
		return x.PartitionId
	}
	return ""
}

func (x *Record) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *Record) GetEventTimestamp() int32 {
	if x != nil {
		return x.EventTimestamp
	}
	return 0
}

type MetricData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricName string  `protobuf:"bytes,1,opt,name=metricName,proto3" json:"metricName,omitempty"`
	Value      float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MetricData) Reset() {
	*x = MetricData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricData) ProtoMessage() {}

func (x *MetricData) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricData.ProtoReflect.Descriptor instead.
func (*MetricData) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{10}
}

func (x *MetricData) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *MetricData) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Partition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicName      string `protobuf:"bytes,1,opt,name=topicName,proto3" json:"topicName,omitempty"`
	PartitionIndex int32  `protobuf:"varint,2,opt,name=partitionIndex,proto3" json:"partitionIndex,omitempty"`
	MessageId      []byte `protobuf:"bytes,3,opt,name=messageId,proto3" json:"messageId,omitempty"`
}

func (x *Partition) Reset() {
	*x = Partition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partition) ProtoMessage() {}

func (x *Partition) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partition.ProtoReflect.Descriptor instead.
func (*Partition) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{11}
}

func (x *Partition) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *Partition) GetPartitionIndex() int32 {
	if x != nil {
		return x.PartitionIndex
	}
	return 0
}

func (x *Partition) GetMessageId() []byte {
	if x != nil {
		return x.MessageId
	}
	return nil
}

var File_context_proto protoreflect.FileDescriptor

var file_context_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a,
	0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x1c, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x37, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x38, 0x0a, 0x0c, 0x49, 0x6e, 0x63, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1f, 0x0a, 0x07, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x1b, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8d,
	0x03, 0x0a, 0x0d, 0x50, 0x75, 0x6c, 0x73, 0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x41, 0x66, 0x74, 0x65, 0x72, 0x22, 0xda,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x42, 0x0a, 0x0a, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x6f, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x32, 0x90, 0x05, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x16,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x50, 0x75, 0x6c, 0x73, 0x61, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x6b,
	0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x05, 0x70, 0x61, 0x75, 0x73, 0x65, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12,
	0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x08, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x14, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x75, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x0a, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x69, 0x6e, 0x63, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x49, 0x6e, 0x63,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_context_proto_rawDescOnce sync.Once
	file_context_proto_rawDescData = file_context_proto_rawDesc
)

func file_context_proto_rawDescGZIP() []byte {
	file_context_proto_rawDescOnce.Do(func() {
		file_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_context_proto_rawDescData)
	})
	return file_context_proto_rawDescData
}

var file_context_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_context_proto_goTypes = []interface{}{
	(*SendMessageId)(nil), // 0: context.SendMessageId
	(*StateKey)(nil),      // 1: context.StateKey
	(*StateKeyValue)(nil), // 2: context.StateKeyValue
	(*StateResult)(nil),   // 3: context.StateResult
	(*IncrStateKey)(nil),  // 4: context.IncrStateKey
	(*Counter)(nil),       // 5: context.Counter
	(*Request)(nil),       // 6: context.Request
	(*MessageId)(nil),     // 7: context.MessageId
	(*PulsarMessage)(nil), // 8: context.PulsarMessage
	(*Record)(nil),        // 9: context.Record
	(*MetricData)(nil),    // 10: context.MetricData
	(*Partition)(nil),     // 11: context.Partition
	(*emptypb.Empty)(nil), // 12: google.protobuf.Empty
}
var file_context_proto_depIdxs = []int32{
	8,  // 0: context.ContextService.publish:input_type -> context.PulsarMessage
	7,  // 1: context.ContextService.currentRecord:input_type -> context.MessageId
	10, // 2: context.ContextService.recordMetrics:input_type -> context.MetricData
	11, // 3: context.ContextService.seek:input_type -> context.Partition
	11, // 4: context.ContextService.pause:input_type -> context.Partition
	11, // 5: context.ContextService.resume:input_type -> context.Partition
	1,  // 6: context.ContextService.getState:input_type -> context.StateKey
	2,  // 7: context.ContextService.putState:input_type -> context.StateKeyValue
	1,  // 8: context.ContextService.deleteState:input_type -> context.StateKey
	1,  // 9: context.ContextService.getCounter:input_type -> context.StateKey
	4,  // 10: context.ContextService.incrCounter:input_type -> context.IncrStateKey
	0,  // 11: context.ContextService.publish:output_type -> context.SendMessageId
	9,  // 12: context.ContextService.currentRecord:output_type -> context.Record
	12, // 13: context.ContextService.recordMetrics:output_type -> google.protobuf.Empty
	12, // 14: context.ContextService.seek:output_type -> google.protobuf.Empty
	12, // 15: context.ContextService.pause:output_type -> google.protobuf.Empty
	12, // 16: context.ContextService.resume:output_type -> google.protobuf.Empty
	3,  // 17: context.ContextService.getState:output_type -> context.StateResult
	12, // 18: context.ContextService.putState:output_type -> google.protobuf.Empty
	12, // 19: context.ContextService.deleteState:output_type -> google.protobuf.Empty
	5,  // 20: context.ContextService.getCounter:output_type -> context.Counter
	12, // 21: context.ContextService.incrCounter:output_type -> google.protobuf.Empty
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_context_proto_init() }
func file_context_proto_init() {
	if File_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateKeyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrStateKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PulsarMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_context_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_context_proto_goTypes,
		DependencyIndexes: file_context_proto_depIdxs,
		MessageInfos:      file_context_proto_msgTypes,
	}.Build()
	File_context_proto = out.File
	file_context_proto_rawDesc = nil
	file_context_proto_goTypes = nil
	file_context_proto_depIdxs = nil
}
