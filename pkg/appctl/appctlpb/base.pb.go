// Copyright (C) 2024  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: base.proto

package appctlpb

import (
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

type AppStatus int32

const (
	AppStatus_UNKNOWN AppStatus = 0
	// RPC service is available, but proxy is not started.
	AppStatus_IDLE AppStatus = 1
	// Proxy is starting.
	AppStatus_STARTING AppStatus = 2
	// Proxy is running.
	AppStatus_RUNNING AppStatus = 3
	// Proxy is being stopped.
	AppStatus_STOPPING AppStatus = 4
)

// Enum value maps for AppStatus.
var (
	AppStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "IDLE",
		2: "STARTING",
		3: "RUNNING",
		4: "STOPPING",
	}
	AppStatus_value = map[string]int32{
		"UNKNOWN":  0,
		"IDLE":     1,
		"STARTING": 2,
		"RUNNING":  3,
		"STOPPING": 4,
	}
)

func (x AppStatus) Enum() *AppStatus {
	p := new(AppStatus)
	*p = x
	return p
}

func (x AppStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[0].Descriptor()
}

func (AppStatus) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[0]
}

func (x AppStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppStatus.Descriptor instead.
func (AppStatus) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

type LoggingLevel int32

const (
	LoggingLevel_DEFAULT LoggingLevel = 0
	LoggingLevel_FATAL   LoggingLevel = 1
	LoggingLevel_ERROR   LoggingLevel = 2
	LoggingLevel_WARN    LoggingLevel = 3
	LoggingLevel_INFO    LoggingLevel = 4
	LoggingLevel_DEBUG   LoggingLevel = 5
	LoggingLevel_TRACE   LoggingLevel = 6
)

// Enum value maps for LoggingLevel.
var (
	LoggingLevel_name = map[int32]string{
		0: "DEFAULT",
		1: "FATAL",
		2: "ERROR",
		3: "WARN",
		4: "INFO",
		5: "DEBUG",
		6: "TRACE",
	}
	LoggingLevel_value = map[string]int32{
		"DEFAULT": 0,
		"FATAL":   1,
		"ERROR":   2,
		"WARN":    3,
		"INFO":    4,
		"DEBUG":   5,
		"TRACE":   6,
	}
)

func (x LoggingLevel) Enum() *LoggingLevel {
	p := new(LoggingLevel)
	*p = x
	return p
}

func (x LoggingLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoggingLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[1].Descriptor()
}

func (LoggingLevel) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[1]
}

func (x LoggingLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoggingLevel.Descriptor instead.
func (LoggingLevel) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

type TransportProtocol int32

const (
	TransportProtocol_UNKNOWN_TRANSPORT_PROTOCOL TransportProtocol = 0
	TransportProtocol_UDP                        TransportProtocol = 1
	TransportProtocol_TCP                        TransportProtocol = 2
)

// Enum value maps for TransportProtocol.
var (
	TransportProtocol_name = map[int32]string{
		0: "UNKNOWN_TRANSPORT_PROTOCOL",
		1: "UDP",
		2: "TCP",
	}
	TransportProtocol_value = map[string]int32{
		"UNKNOWN_TRANSPORT_PROTOCOL": 0,
		"UDP":                        1,
		"TCP":                        2,
	}
)

func (x TransportProtocol) Enum() *TransportProtocol {
	p := new(TransportProtocol)
	*p = x
	return p
}

func (x TransportProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransportProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[2].Descriptor()
}

func (TransportProtocol) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[2]
}

func (x TransportProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransportProtocol.Descriptor instead.
func (TransportProtocol) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

type AppStatusMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *AppStatus `protobuf:"varint,1,opt,name=status,proto3,enum=appctl.AppStatus,oneof" json:"status,omitempty"`
}

func (x *AppStatusMsg) Reset() {
	*x = AppStatusMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppStatusMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppStatusMsg) ProtoMessage() {}

func (x *AppStatusMsg) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppStatusMsg.ProtoReflect.Descriptor instead.
func (*AppStatusMsg) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *AppStatusMsg) GetStatus() AppStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return AppStatus_UNKNOWN
}

type ServerEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String representation of IP address.
	// It can be either IPv4 or IPv6.
	IpAddress *string `protobuf:"bytes,1,opt,name=ipAddress,proto3,oneof" json:"ipAddress,omitempty"`
	// Server's full qualified domain name.
	// When this is set, the `ipAddress` field is ignored.
	DomainName *string `protobuf:"bytes,2,opt,name=domainName,proto3,oneof" json:"domainName,omitempty"`
	// Server's port-protocol bindings.
	PortBindings []*PortBinding `protobuf:"bytes,3,rep,name=portBindings,proto3" json:"portBindings,omitempty"`
}

func (x *ServerEndpoint) Reset() {
	*x = ServerEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerEndpoint) ProtoMessage() {}

func (x *ServerEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerEndpoint.ProtoReflect.Descriptor instead.
func (*ServerEndpoint) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *ServerEndpoint) GetIpAddress() string {
	if x != nil && x.IpAddress != nil {
		return *x.IpAddress
	}
	return ""
}

func (x *ServerEndpoint) GetDomainName() string {
	if x != nil && x.DomainName != nil {
		return *x.DomainName
	}
	return ""
}

func (x *ServerEndpoint) GetPortBindings() []*PortBinding {
	if x != nil {
		return x.PortBindings
	}
	return nil
}

type PortBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A single port number.
	// This field can't be set with portRange at the same time.
	Port     *int32             `protobuf:"varint,1,opt,name=port,proto3,oneof" json:"port,omitempty"`
	Protocol *TransportProtocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=appctl.TransportProtocol,oneof" json:"protocol,omitempty"`
	// A port number range.
	// For example, "8000-9000" contains 1001 ports from 8000 to 9000.
	// This field can't be set with port at the same time.
	PortRange *string `protobuf:"bytes,3,opt,name=portRange,proto3,oneof" json:"portRange,omitempty"`
}

func (x *PortBinding) Reset() {
	*x = PortBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortBinding) ProtoMessage() {}

func (x *PortBinding) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortBinding.ProtoReflect.Descriptor instead.
func (*PortBinding) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *PortBinding) GetPort() int32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *PortBinding) GetProtocol() TransportProtocol {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return TransportProtocol_UNKNOWN_TRANSPORT_PROTOCOL
}

func (x *PortBinding) GetPortRange() string {
	if x != nil && x.PortRange != nil {
		return *x.PortRange
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User name is also the ID of user.
	// Typically this is an email address.
	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Raw password.
	// For safety this shouldn't be persisted at the server side.
	Password *string `protobuf:"bytes,2,opt,name=password,proto3,oneof" json:"password,omitempty"`
	// Hashed password.
	// Stored with hex encoding.
	HashedPassword *string `protobuf:"bytes,3,opt,name=hashedPassword,proto3,oneof" json:"hashedPassword,omitempty"`
	// User quotas.
	// This has no effect at the client side.
	Quotas []*Quota `protobuf:"bytes,4,rep,name=quotas,proto3" json:"quotas,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *User) GetHashedPassword() string {
	if x != nil && x.HashedPassword != nil {
		return *x.HashedPassword
	}
	return ""
}

func (x *User) GetQuotas() []*Quota {
	if x != nil {
		return x.Quotas
	}
	return nil
}

type Quota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of days to renew the quota.
	// The renew is rolling based.
	Days *int32 `protobuf:"varint,1,opt,name=days,proto3,oneof" json:"days,omitempty"`
	// Number of megabytes the user allowed to send and receive.
	Megabytes *int32 `protobuf:"varint,2,opt,name=megabytes,proto3,oneof" json:"megabytes,omitempty"`
}

func (x *Quota) Reset() {
	*x = Quota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quota) ProtoMessage() {}

func (x *Quota) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quota.ProtoReflect.Descriptor instead.
func (*Quota) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{5}
}

func (x *Quota) GetDays() int32 {
	if x != nil && x.Days != nil {
		return *x.Days
	}
	return 0
}

func (x *Quota) GetMegabytes() int32 {
	if x != nil && x.Megabytes != nil {
		return *x.Megabytes
	}
	return 0
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70,
	0x70, 0x63, 0x74, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x49, 0x0a,
	0x0c, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x2e, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x23,
	0x0a, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x63,
	0x74, 0x6c, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0c,
	0x70, 0x6f, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x50, 0x6f,
	0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x3a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48,
	0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x21,
	0x0a, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x52, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5a, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x17,
	0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04,
	0x64, 0x61, 0x79, 0x73, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6d, 0x65, 0x67, 0x61, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x09, 0x6d, 0x65,
	0x67, 0x61, 0x62, 0x79, 0x74, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64,
	0x61, 0x79, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6d, 0x65, 0x67, 0x61, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x2a, 0x4b, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x44, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x2a, 0x5b,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x46,
	0x41, 0x54, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x4e, 0x46, 0x4f, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x05,
	0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x06, 0x2a, 0x45, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x50, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x43, 0x50,
	0x10, 0x02, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x66, 0x65, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x65, 0x72, 0x75, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_base_proto_goTypes = []interface{}{
	(AppStatus)(0),         // 0: appctl.AppStatus
	(LoggingLevel)(0),      // 1: appctl.LoggingLevel
	(TransportProtocol)(0), // 2: appctl.TransportProtocol
	(*Empty)(nil),          // 3: appctl.Empty
	(*AppStatusMsg)(nil),   // 4: appctl.AppStatusMsg
	(*ServerEndpoint)(nil), // 5: appctl.ServerEndpoint
	(*PortBinding)(nil),    // 6: appctl.PortBinding
	(*User)(nil),           // 7: appctl.User
	(*Quota)(nil),          // 8: appctl.Quota
}
var file_base_proto_depIdxs = []int32{
	0, // 0: appctl.AppStatusMsg.status:type_name -> appctl.AppStatus
	6, // 1: appctl.ServerEndpoint.portBindings:type_name -> appctl.PortBinding
	2, // 2: appctl.PortBinding.protocol:type_name -> appctl.TransportProtocol
	8, // 3: appctl.User.quotas:type_name -> appctl.Quota
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppStatusMsg); i {
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
		file_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerEndpoint); i {
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
		file_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortBinding); i {
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
		file_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quota); i {
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
	file_base_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_base_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_base_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_base_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_base_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		EnumInfos:         file_base_proto_enumTypes,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
