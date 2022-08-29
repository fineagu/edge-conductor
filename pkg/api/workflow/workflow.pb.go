//
// Copyright (c) 2022 Intel Corporation.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-rc.1
// 	protoc        v3.14.0
// source: api/proto/workflow.proto

package workflow

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Result_Return int32

const (
	Result_Success Result_Return = 0
	Result_Error   Result_Return = 1
)

// Enum value maps for Result_Return.
var (
	Result_Return_name = map[int32]string{
		0: "Success",
		1: "Error",
	}
	Result_Return_value = map[string]int32{
		"Success": 0,
		"Error":   1,
	}
)

func (x Result_Return) Enum() *Result_Return {
	p := new(Result_Return)
	*p = x
	return p
}

func (x Result_Return) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result_Return) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_workflow_proto_enumTypes[0].Descriptor()
}

func (Result_Return) Type() protoreflect.EnumType {
	return &file_api_proto_workflow_proto_enumTypes[0]
}

func (x Result_Return) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result_Return.Descriptor instead.
func (Result_Return) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_workflow_proto_rawDescGZIP(), []int{4, 0}
}

type ConnectResult_Return int32

const (
	ConnectResult_Connected ConnectResult_Return = 0
	ConnectResult_Completed ConnectResult_Return = 1
	ConnectResult_Error     ConnectResult_Return = 2
)

// Enum value maps for ConnectResult_Return.
var (
	ConnectResult_Return_name = map[int32]string{
		0: "Connected",
		1: "Completed",
		2: "Error",
	}
	ConnectResult_Return_value = map[string]int32{
		"Connected": 0,
		"Completed": 1,
		"Error":     2,
	}
)

func (x ConnectResult_Return) Enum() *ConnectResult_Return {
	p := new(ConnectResult_Return)
	*p = x
	return p
}

func (x ConnectResult_Return) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectResult_Return) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_workflow_proto_enumTypes[1].Descriptor()
}

func (ConnectResult_Return) Type() protoreflect.EnumType {
	return &file_api_proto_workflow_proto_enumTypes[1]
}

func (x ConnectResult_Return) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectResult_Return.Descriptor instead.
func (ConnectResult_Return) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_workflow_proto_rawDescGZIP(), []int{5, 0}
}

type PluginConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin *Plugin `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
}

func (x *PluginConnectRequest) Reset() {
	*x = PluginConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConnectRequest) ProtoMessage() {}

func (x *PluginConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConnectRequest.ProtoReflect.Descriptor instead.
func (*PluginConnectRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *PluginConnectRequest) GetPlugin() *Plugin {
	if x != nil {
		return x.Plugin
	}
	return nil
}

type PluginCompleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin       *Plugin       `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Result       *Result       `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	WorkflowData *WorkflowData `protobuf:"bytes,3,opt,name=workflow_data,json=workflowData,proto3" json:"workflow_data,omitempty"`
}

func (x *PluginCompleteRequest) Reset() {
	*x = PluginCompleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginCompleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginCompleteRequest) ProtoMessage() {}

func (x *PluginCompleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginCompleteRequest.ProtoReflect.Descriptor instead.
func (*PluginCompleteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *PluginCompleteRequest) GetPlugin() *Plugin {
	if x != nil {
		return x.Plugin
	}
	return nil
}

func (x *PluginCompleteRequest) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *PluginCompleteRequest) GetWorkflowData() *WorkflowData {
	if x != nil {
		return x.WorkflowData
	}
	return nil
}

type PluginConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowData *WorkflowData  `protobuf:"bytes,1,opt,name=workflow_data,json=workflowData,proto3" json:"workflow_data,omitempty"`
	Result       *ConnectResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PluginConnectResponse) Reset() {
	*x = PluginConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_workflow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConnectResponse) ProtoMessage() {}

func (x *PluginConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_workflow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConnectResponse.ProtoReflect.Descriptor instead.
func (*PluginConnectResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *PluginConnectResponse) GetWorkflowData() *WorkflowData {
	if x != nil {
		return x.WorkflowData
	}
	return nil
}

func (x *PluginConnectResponse) GetResult() *ConnectResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type WorkflowData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	PluginData []byte `protobuf:"bytes,2,opt,name=plugin_data,json=pluginData,proto3" json:"plugin_data,omitempty"`
}

func (x *WorkflowData) Reset() {
	*x = WorkflowData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_workflow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowData) ProtoMessage() {}

func (x *WorkflowData) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_workflow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowData.ProtoReflect.Descriptor instead.
func (*WorkflowData) Descriptor() ([]byte, []int) {
	return file_api_proto_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *WorkflowData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *WorkflowData) GetPluginData() []byte {
	if x != nil {
		return x.PluginData
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Return Result_Return `protobuf:"varint,1,opt,name=return,proto3,enum=workflow.Result_Return" json:"return,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_workflow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_workflow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_api_proto_workflow_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetReturn() Result_Return {
	if x != nil {
		return x.Return
	}
	return Result_Success
}

type ConnectResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Return ConnectResult_Return `protobuf:"varint,1,opt,name=return,proto3,enum=workflow.ConnectResult_Return" json:"return,omitempty"`
}

func (x *ConnectResult) Reset() {
	*x = ConnectResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_workflow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResult) ProtoMessage() {}

func (x *ConnectResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_workflow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResult.ProtoReflect.Descriptor instead.
func (*ConnectResult) Descriptor() ([]byte, []int) {
	return file_api_proto_workflow_proto_rawDescGZIP(), []int{5}
}

func (x *ConnectResult) GetReturn() ConnectResult_Return {
	if x != nil {
		return x.Return
	}
	return ConnectResult_Connected
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_workflow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_workflow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_api_proto_workflow_proto_rawDescGZIP(), []int{6}
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_workflow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_workflow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_api_proto_workflow_proto_rawDescGZIP(), []int{7}
}

func (x *Log) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

var File_api_proto_workflow_proto protoreflect.FileDescriptor

var file_api_proto_workflow_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x22, 0x40, 0x0a, 0x14, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x06,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0xa8, 0x01, 0x0a, 0x15, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x43, 0x0a, 0x0c, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x5b,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x52, 0x06, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x20, 0x0a, 0x06, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x22, 0x7a, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x06,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x22, 0x31, 0x0a, 0x06, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x22, 0x1c, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x32, 0xda,
	0x01, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x52, 0x0a, 0x0d, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x75, 0x74, 0x4c, 0x6f, 0x67, 0x12,
	0x0d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x1a, 0x10,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x28, 0x01, 0x12, 0x45, 0x0a, 0x0e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_workflow_proto_rawDescOnce sync.Once
	file_api_proto_workflow_proto_rawDescData = file_api_proto_workflow_proto_rawDesc
)

func file_api_proto_workflow_proto_rawDescGZIP() []byte {
	file_api_proto_workflow_proto_rawDescOnce.Do(func() {
		file_api_proto_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_workflow_proto_rawDescData)
	})
	return file_api_proto_workflow_proto_rawDescData
}

var file_api_proto_workflow_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_proto_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_workflow_proto_goTypes = []interface{}{
	(Result_Return)(0),            // 0: workflow.Result.Return
	(ConnectResult_Return)(0),     // 1: workflow.ConnectResult.Return
	(*PluginConnectRequest)(nil),  // 2: workflow.PluginConnectRequest
	(*PluginCompleteRequest)(nil), // 3: workflow.PluginCompleteRequest
	(*PluginConnectResponse)(nil), // 4: workflow.PluginConnectResponse
	(*WorkflowData)(nil),          // 5: workflow.WorkflowData
	(*Result)(nil),                // 6: workflow.Result
	(*ConnectResult)(nil),         // 7: workflow.ConnectResult
	(*Plugin)(nil),                // 8: workflow.Plugin
	(*Log)(nil),                   // 9: workflow.Log
}
var file_api_proto_workflow_proto_depIdxs = []int32{
	8,  // 0: workflow.PluginConnectRequest.plugin:type_name -> workflow.Plugin
	8,  // 1: workflow.PluginCompleteRequest.plugin:type_name -> workflow.Plugin
	6,  // 2: workflow.PluginCompleteRequest.result:type_name -> workflow.Result
	5,  // 3: workflow.PluginCompleteRequest.workflow_data:type_name -> workflow.WorkflowData
	5,  // 4: workflow.PluginConnectResponse.workflow_data:type_name -> workflow.WorkflowData
	7,  // 5: workflow.PluginConnectResponse.result:type_name -> workflow.ConnectResult
	0,  // 6: workflow.Result.return:type_name -> workflow.Result.Return
	1,  // 7: workflow.ConnectResult.return:type_name -> workflow.ConnectResult.Return
	2,  // 8: workflow.Workflow.PluginConnect:input_type -> workflow.PluginConnectRequest
	9,  // 9: workflow.Workflow.PluginPutLog:input_type -> workflow.Log
	3,  // 10: workflow.Workflow.PluginComplete:input_type -> workflow.PluginCompleteRequest
	4,  // 11: workflow.Workflow.PluginConnect:output_type -> workflow.PluginConnectResponse
	6,  // 12: workflow.Workflow.PluginPutLog:output_type -> workflow.Result
	6,  // 13: workflow.Workflow.PluginComplete:output_type -> workflow.Result
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_proto_workflow_proto_init() }
func file_api_proto_workflow_proto_init() {
	if File_api_proto_workflow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginConnectRequest); i {
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
		file_api_proto_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginCompleteRequest); i {
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
		file_api_proto_workflow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginConnectResponse); i {
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
		file_api_proto_workflow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowData); i {
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
		file_api_proto_workflow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_api_proto_workflow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResult); i {
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
		file_api_proto_workflow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_api_proto_workflow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_api_proto_workflow_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_workflow_proto_goTypes,
		DependencyIndexes: file_api_proto_workflow_proto_depIdxs,
		EnumInfos:         file_api_proto_workflow_proto_enumTypes,
		MessageInfos:      file_api_proto_workflow_proto_msgTypes,
	}.Build()
	File_api_proto_workflow_proto = out.File
	file_api_proto_workflow_proto_rawDesc = nil
	file_api_proto_workflow_proto_goTypes = nil
	file_api_proto_workflow_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkflowClient is the client API for Workflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkflowClient interface {
	PluginConnect(ctx context.Context, in *PluginConnectRequest, opts ...grpc.CallOption) (*PluginConnectResponse, error)
	PluginPutLog(ctx context.Context, opts ...grpc.CallOption) (Workflow_PluginPutLogClient, error)
	PluginComplete(ctx context.Context, in *PluginCompleteRequest, opts ...grpc.CallOption) (*Result, error)
}

type workflowClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowClient(cc grpc.ClientConnInterface) WorkflowClient {
	return &workflowClient{cc}
}

func (c *workflowClient) PluginConnect(ctx context.Context, in *PluginConnectRequest, opts ...grpc.CallOption) (*PluginConnectResponse, error) {
	out := new(PluginConnectResponse)
	err := c.cc.Invoke(ctx, "/workflow.Workflow/PluginConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) PluginPutLog(ctx context.Context, opts ...grpc.CallOption) (Workflow_PluginPutLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Workflow_serviceDesc.Streams[0], "/workflow.Workflow/PluginPutLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &workflowPluginPutLogClient{stream}
	return x, nil
}

type Workflow_PluginPutLogClient interface {
	Send(*Log) error
	CloseAndRecv() (*Result, error)
	grpc.ClientStream
}

type workflowPluginPutLogClient struct {
	grpc.ClientStream
}

func (x *workflowPluginPutLogClient) Send(m *Log) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workflowPluginPutLogClient) CloseAndRecv() (*Result, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workflowClient) PluginComplete(ctx context.Context, in *PluginCompleteRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/workflow.Workflow/PluginComplete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServer is the server API for Workflow service.
type WorkflowServer interface {
	PluginConnect(context.Context, *PluginConnectRequest) (*PluginConnectResponse, error)
	PluginPutLog(Workflow_PluginPutLogServer) error
	PluginComplete(context.Context, *PluginCompleteRequest) (*Result, error)
}

// UnimplementedWorkflowServer can be embedded to have forward compatible implementations.
type UnimplementedWorkflowServer struct {
}

func (*UnimplementedWorkflowServer) PluginConnect(context.Context, *PluginConnectRequest) (*PluginConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginConnect not implemented")
}
func (*UnimplementedWorkflowServer) PluginPutLog(Workflow_PluginPutLogServer) error {
	return status.Errorf(codes.Unimplemented, "method PluginPutLog not implemented")
}
func (*UnimplementedWorkflowServer) PluginComplete(context.Context, *PluginCompleteRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginComplete not implemented")
}

func RegisterWorkflowServer(s *grpc.Server, srv WorkflowServer) {
	s.RegisterService(&_Workflow_serviceDesc, srv)
}

func _Workflow_PluginConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).PluginConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.Workflow/PluginConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).PluginConnect(ctx, req.(*PluginConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_PluginPutLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkflowServer).PluginPutLog(&workflowPluginPutLogServer{stream})
}

type Workflow_PluginPutLogServer interface {
	SendAndClose(*Result) error
	Recv() (*Log, error)
	grpc.ServerStream
}

type workflowPluginPutLogServer struct {
	grpc.ServerStream
}

func (x *workflowPluginPutLogServer) SendAndClose(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workflowPluginPutLogServer) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Workflow_PluginComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginCompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).PluginComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.Workflow/PluginComplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).PluginComplete(ctx, req.(*PluginCompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workflow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.Workflow",
	HandlerType: (*WorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginConnect",
			Handler:    _Workflow_PluginConnect_Handler,
		},
		{
			MethodName: "PluginComplete",
			Handler:    _Workflow_PluginComplete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PluginPutLog",
			Handler:       _Workflow_PluginPutLog_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/proto/workflow.proto",
}