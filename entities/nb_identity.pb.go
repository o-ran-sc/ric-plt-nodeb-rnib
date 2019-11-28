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

//  This source code is part of the near-RT RIC (RAN Intelligent Controller)
//  platform project (RICP).

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nb_identity.proto

package entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NbIdentity struct {
	InventoryName        string      `protobuf:"bytes,1,opt,name=inventory_name,json=inventoryName,proto3" json:"inventory_name,omitempty"`
	GlobalNbId           *GlobalNbId `protobuf:"bytes,2,opt,name=global_nb_id,json=globalNbId,proto3" json:"global_nb_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NbIdentity) Reset()         { *m = NbIdentity{} }
func (m *NbIdentity) String() string { return proto.CompactTextString(m) }
func (*NbIdentity) ProtoMessage()    {}
func (*NbIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07701eb9efb4b89, []int{0}
}

func (m *NbIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NbIdentity.Unmarshal(m, b)
}
func (m *NbIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NbIdentity.Marshal(b, m, deterministic)
}
func (m *NbIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NbIdentity.Merge(m, src)
}
func (m *NbIdentity) XXX_Size() int {
	return xxx_messageInfo_NbIdentity.Size(m)
}
func (m *NbIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_NbIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_NbIdentity proto.InternalMessageInfo

func (m *NbIdentity) GetInventoryName() string {
	if m != nil {
		return m.InventoryName
	}
	return ""
}

func (m *NbIdentity) GetGlobalNbId() *GlobalNbId {
	if m != nil {
		return m.GlobalNbId
	}
	return nil
}

type GlobalNbId struct {
	PlmnId               string   `protobuf:"bytes,1,opt,name=plmn_id,json=plmnId,proto3" json:"plmn_id,omitempty"`
	NbId                 string   `protobuf:"bytes,2,opt,name=nb_id,json=nbId,proto3" json:"nb_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlobalNbId) Reset()         { *m = GlobalNbId{} }
func (m *GlobalNbId) String() string { return proto.CompactTextString(m) }
func (*GlobalNbId) ProtoMessage()    {}
func (*GlobalNbId) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07701eb9efb4b89, []int{1}
}

func (m *GlobalNbId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalNbId.Unmarshal(m, b)
}
func (m *GlobalNbId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalNbId.Marshal(b, m, deterministic)
}
func (m *GlobalNbId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalNbId.Merge(m, src)
}
func (m *GlobalNbId) XXX_Size() int {
	return xxx_messageInfo_GlobalNbId.Size(m)
}
func (m *GlobalNbId) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalNbId.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalNbId proto.InternalMessageInfo

func (m *GlobalNbId) GetPlmnId() string {
	if m != nil {
		return m.PlmnId
	}
	return ""
}

func (m *GlobalNbId) GetNbId() string {
	if m != nil {
		return m.NbId
	}
	return ""
}

func init() {
	proto.RegisterType((*NbIdentity)(nil), "entities.NbIdentity")
	proto.RegisterType((*GlobalNbId)(nil), "entities.GlobalNbId")
}

func init() { proto.RegisterFile("nb_identity.proto", fileDescriptor_a07701eb9efb4b89) }

var fileDescriptor_a07701eb9efb4b89 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x4b, 0x8a, 0xcf,
	0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00,
	0xf3, 0x32, 0x53, 0x8b, 0x95, 0xb2, 0xb9, 0xb8, 0xfc, 0x92, 0x3c, 0xa1, 0xb2, 0x42, 0xaa, 0x5c,
	0x7c, 0x99, 0x79, 0x65, 0xa9, 0x79, 0x25, 0xf9, 0x45, 0x95, 0xf1, 0x79, 0x89, 0xb9, 0xa9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xbc, 0x70, 0x51, 0xbf, 0xc4, 0xdc, 0x54, 0x21, 0x33, 0x2e,
	0x9e, 0xf4, 0x9c, 0xfc, 0xa4, 0xc4, 0x9c, 0x78, 0xb0, 0xd1, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0x22, 0x7a, 0x30, 0x53, 0xf5, 0xdc, 0xc1, 0xb2, 0x20, 0x83, 0x83, 0xb8, 0xd2, 0xe1, 0x6c,
	0x25, 0x2b, 0x2e, 0x2e, 0x84, 0x8c, 0x90, 0x38, 0x17, 0x7b, 0x41, 0x4e, 0x6e, 0x1e, 0xc8, 0x00,
	0x88, 0x2d, 0x6c, 0x20, 0xae, 0x67, 0x8a, 0x90, 0x30, 0x17, 0x2b, 0xc2, 0x5c, 0xce, 0x20, 0x96,
	0xbc, 0x24, 0xcf, 0x94, 0x24, 0x36, 0xb0, 0xcb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f,
	0x65, 0xd0, 0x4b, 0xce, 0x00, 0x00, 0x00,
}
