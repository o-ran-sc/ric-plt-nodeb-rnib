//
// Copyright 2019 AT&T Intellectual Property
// Copyright 2019 Nokia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//
// This source code is part of the near-RT RIC (RAN Intelligent Controller)
// platform project (RICP).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: nb_identity.proto

package entities

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ConnectionStatus int32

const (
	ConnectionStatus_UNKNOWN_CONNECTION_STATUS ConnectionStatus = 0
	ConnectionStatus_CONNECTED                 ConnectionStatus = 1
	ConnectionStatus_DISCONNECTED              ConnectionStatus = 2
	ConnectionStatus_CONNECTED_SETUP_FAILED    ConnectionStatus = 3
	ConnectionStatus_CONNECTING                ConnectionStatus = 4
	ConnectionStatus_SHUTTING_DOWN             ConnectionStatus = 5
	ConnectionStatus_SHUT_DOWN                 ConnectionStatus = 6
)

// Enum value maps for ConnectionStatus.
var (
	ConnectionStatus_name = map[int32]string{
		0: "UNKNOWN_CONNECTION_STATUS",
		1: "CONNECTED",
		2: "DISCONNECTED",
		3: "CONNECTED_SETUP_FAILED",
		4: "CONNECTING",
		5: "SHUTTING_DOWN",
		6: "SHUT_DOWN",
	}
	ConnectionStatus_value = map[string]int32{
		"UNKNOWN_CONNECTION_STATUS": 0,
		"CONNECTED":                 1,
		"DISCONNECTED":              2,
		"CONNECTED_SETUP_FAILED":    3,
		"CONNECTING":                4,
		"SHUTTING_DOWN":             5,
		"SHUT_DOWN":                 6,
	}
)

func (x ConnectionStatus) Enum() *ConnectionStatus {
	p := new(ConnectionStatus)
	*p = x
	return p
}

func (x ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_nb_identity_proto_enumTypes[0].Descriptor()
}

func (ConnectionStatus) Type() protoreflect.EnumType {
	return &file_nb_identity_proto_enumTypes[0]
}

func (x ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionStatus.Descriptor instead.
func (ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_nb_identity_proto_rawDescGZIP(), []int{0}
}

type GlobalNbId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlmnId string `protobuf:"bytes,1,opt,name=plmn_id,json=plmnId,proto3" json:"plmn_id,omitempty"`
	NbId   string `protobuf:"bytes,2,opt,name=nb_id,json=nbId,proto3" json:"nb_id,omitempty"`
}

func (x *GlobalNbId) Reset() {
	*x = GlobalNbId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nb_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalNbId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalNbId) ProtoMessage() {}

func (x *GlobalNbId) ProtoReflect() protoreflect.Message {
	mi := &file_nb_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalNbId.ProtoReflect.Descriptor instead.
func (*GlobalNbId) Descriptor() ([]byte, []int) {
	return file_nb_identity_proto_rawDescGZIP(), []int{0}
}

func (x *GlobalNbId) GetPlmnId() string {
	if x != nil {
		return x.PlmnId
	}
	return ""
}

func (x *GlobalNbId) GetNbId() string {
	if x != nil {
		return x.NbId
	}
	return ""
}

type NbIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InventoryName                string           `protobuf:"bytes,1,opt,name=inventory_name,json=inventoryName,proto3" json:"inventory_name,omitempty"`
	GlobalNbId                   *GlobalNbId      `protobuf:"bytes,2,opt,name=global_nb_id,json=globalNbId,proto3" json:"global_nb_id,omitempty"`
	ConnectionStatus             ConnectionStatus `protobuf:"varint,3,opt,name=connection_status,json=connectionStatus,proto3,enum=entities.ConnectionStatus" json:"connection_status,omitempty"`
	HealthCheckTimestampSent     int64            `protobuf:"varint,4,opt,name=health_check_timestamp_sent,json=healthCheckTimestampSent,proto3" json:"health_check_timestamp_sent,omitempty"`
	HealthCheckTimestampReceived int64            `protobuf:"varint,5,opt,name=health_check_timestamp_received,json=healthCheckTimestampReceived,proto3" json:"health_check_timestamp_received,omitempty"`
}

func (x *NbIdentity) Reset() {
	*x = NbIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nb_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NbIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NbIdentity) ProtoMessage() {}

func (x *NbIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_nb_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NbIdentity.ProtoReflect.Descriptor instead.
func (*NbIdentity) Descriptor() ([]byte, []int) {
	return file_nb_identity_proto_rawDescGZIP(), []int{1}
}

func (x *NbIdentity) GetInventoryName() string {
	if x != nil {
		return x.InventoryName
	}
	return ""
}

func (x *NbIdentity) GetGlobalNbId() *GlobalNbId {
	if x != nil {
		return x.GlobalNbId
	}
	return nil
}

func (x *NbIdentity) GetConnectionStatus() ConnectionStatus {
	if x != nil {
		return x.ConnectionStatus
	}
	return ConnectionStatus_UNKNOWN_CONNECTION_STATUS
}

func (x *NbIdentity) GetHealthCheckTimestampSent() int64 {
	if x != nil {
		return x.HealthCheckTimestampSent
	}
	return 0
}

func (x *NbIdentity) GetHealthCheckTimestampReceived() int64 {
	if x != nil {
		return x.HealthCheckTimestampReceived
	}
	return 0
}

var File_nb_identity_proto protoreflect.FileDescriptor

var file_nb_identity_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x62, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x0a,
	0x0a, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4e, 0x62, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x6c, 0x6d, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c,
	0x6d, 0x6e, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x62, 0x49, 0x64, 0x22, 0xba, 0x02, 0x0a, 0x0a, 0x4e, 0x62,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x36, 0x0a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6e, 0x62, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4e, 0x62, 0x49, 0x64, 0x52, 0x0a, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x4e, 0x62, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3d, 0x0a, 0x1b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x53, 0x65, 0x6e, 0x74, 0x12,
	0x45, 0x0a, 0x1f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x2a, 0xa0, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x45, 0x54, 0x55, 0x50, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x48, 0x55, 0x54, 0x54,
	0x49, 0x4e, 0x47, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x48,
	0x55, 0x54, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x06, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_nb_identity_proto_rawDescOnce sync.Once
	file_nb_identity_proto_rawDescData = file_nb_identity_proto_rawDesc
)

func file_nb_identity_proto_rawDescGZIP() []byte {
	file_nb_identity_proto_rawDescOnce.Do(func() {
		file_nb_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_nb_identity_proto_rawDescData)
	})
	return file_nb_identity_proto_rawDescData
}

var file_nb_identity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nb_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nb_identity_proto_goTypes = []interface{}{
	(ConnectionStatus)(0), // 0: entities.ConnectionStatus
	(*GlobalNbId)(nil),    // 1: entities.GlobalNbId
	(*NbIdentity)(nil),    // 2: entities.NbIdentity
}
var file_nb_identity_proto_depIdxs = []int32{
	1, // 0: entities.NbIdentity.global_nb_id:type_name -> entities.GlobalNbId
	0, // 1: entities.NbIdentity.connection_status:type_name -> entities.ConnectionStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_nb_identity_proto_init() }
func file_nb_identity_proto_init() {
	if File_nb_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nb_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalNbId); i {
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
		file_nb_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NbIdentity); i {
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
			RawDescriptor: file_nb_identity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nb_identity_proto_goTypes,
		DependencyIndexes: file_nb_identity_proto_depIdxs,
		EnumInfos:         file_nb_identity_proto_enumTypes,
		MessageInfos:      file_nb_identity_proto_msgTypes,
	}.Build()
	File_nb_identity_proto = out.File
	file_nb_identity_proto_rawDesc = nil
	file_nb_identity_proto_goTypes = nil
	file_nb_identity_proto_depIdxs = nil
}
