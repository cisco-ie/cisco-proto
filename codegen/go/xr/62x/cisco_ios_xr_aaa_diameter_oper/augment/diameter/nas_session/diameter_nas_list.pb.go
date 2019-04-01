/*
Copyright 2019 Cisco Systems

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: diameter_nas_list.proto

package cisco_ios_xr_aaa_diameter_oper_augment_diameter_nas_session

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

type DiameterNasList_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiameterNasList_KEYS) Reset()         { *m = DiameterNasList_KEYS{} }
func (m *DiameterNasList_KEYS) String() string { return proto.CompactTextString(m) }
func (*DiameterNasList_KEYS) ProtoMessage()    {}
func (*DiameterNasList_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c5c5cf3e64bf30, []int{0}
}

func (m *DiameterNasList_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiameterNasList_KEYS.Unmarshal(m, b)
}
func (m *DiameterNasList_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiameterNasList_KEYS.Marshal(b, m, deterministic)
}
func (m *DiameterNasList_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiameterNasList_KEYS.Merge(m, src)
}
func (m *DiameterNasList_KEYS) XXX_Size() int {
	return xxx_messageInfo_DiameterNasList_KEYS.Size(m)
}
func (m *DiameterNasList_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_DiameterNasList_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_DiameterNasList_KEYS proto.InternalMessageInfo

type DiameterNas struct {
	AaaSessionId               string   `protobuf:"bytes,1,opt,name=aaa_session_id,json=aaaSessionId,proto3" json:"aaa_session_id,omitempty"`
	DiameterSessionId          string   `protobuf:"bytes,2,opt,name=diameter_session_id,json=diameterSessionId,proto3" json:"diameter_session_id,omitempty"`
	AuthenticationStatus       uint32   `protobuf:"varint,3,opt,name=authentication_status,json=authenticationStatus,proto3" json:"authentication_status,omitempty"`
	AuthorizationStatus        uint32   `protobuf:"varint,4,opt,name=authorization_status,json=authorizationStatus,proto3" json:"authorization_status,omitempty"`
	AccountingStatus           uint32   `protobuf:"varint,5,opt,name=accounting_status,json=accountingStatus,proto3" json:"accounting_status,omitempty"`
	AccountingStatusStop       uint32   `protobuf:"varint,6,opt,name=accounting_status_stop,json=accountingStatusStop,proto3" json:"accounting_status_stop,omitempty"`
	DisconnectStatus           uint32   `protobuf:"varint,7,opt,name=disconnect_status,json=disconnectStatus,proto3" json:"disconnect_status,omitempty"`
	AccountingIntrimInPackets  uint32   `protobuf:"varint,8,opt,name=accounting_intrim_in_packets,json=accountingIntrimInPackets,proto3" json:"accounting_intrim_in_packets,omitempty"`
	AccountingIntrimOutPackets uint32   `protobuf:"varint,9,opt,name=accounting_intrim_out_packets,json=accountingIntrimOutPackets,proto3" json:"accounting_intrim_out_packets,omitempty"`
	MethodList                 string   `protobuf:"bytes,10,opt,name=method_list,json=methodList,proto3" json:"method_list,omitempty"`
	ServerUsedList             string   `protobuf:"bytes,11,opt,name=server_used_list,json=serverUsedList,proto3" json:"server_used_list,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *DiameterNas) Reset()         { *m = DiameterNas{} }
func (m *DiameterNas) String() string { return proto.CompactTextString(m) }
func (*DiameterNas) ProtoMessage()    {}
func (*DiameterNas) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c5c5cf3e64bf30, []int{1}
}

func (m *DiameterNas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiameterNas.Unmarshal(m, b)
}
func (m *DiameterNas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiameterNas.Marshal(b, m, deterministic)
}
func (m *DiameterNas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiameterNas.Merge(m, src)
}
func (m *DiameterNas) XXX_Size() int {
	return xxx_messageInfo_DiameterNas.Size(m)
}
func (m *DiameterNas) XXX_DiscardUnknown() {
	xxx_messageInfo_DiameterNas.DiscardUnknown(m)
}

var xxx_messageInfo_DiameterNas proto.InternalMessageInfo

func (m *DiameterNas) GetAaaSessionId() string {
	if m != nil {
		return m.AaaSessionId
	}
	return ""
}

func (m *DiameterNas) GetDiameterSessionId() string {
	if m != nil {
		return m.DiameterSessionId
	}
	return ""
}

func (m *DiameterNas) GetAuthenticationStatus() uint32 {
	if m != nil {
		return m.AuthenticationStatus
	}
	return 0
}

func (m *DiameterNas) GetAuthorizationStatus() uint32 {
	if m != nil {
		return m.AuthorizationStatus
	}
	return 0
}

func (m *DiameterNas) GetAccountingStatus() uint32 {
	if m != nil {
		return m.AccountingStatus
	}
	return 0
}

func (m *DiameterNas) GetAccountingStatusStop() uint32 {
	if m != nil {
		return m.AccountingStatusStop
	}
	return 0
}

func (m *DiameterNas) GetDisconnectStatus() uint32 {
	if m != nil {
		return m.DisconnectStatus
	}
	return 0
}

func (m *DiameterNas) GetAccountingIntrimInPackets() uint32 {
	if m != nil {
		return m.AccountingIntrimInPackets
	}
	return 0
}

func (m *DiameterNas) GetAccountingIntrimOutPackets() uint32 {
	if m != nil {
		return m.AccountingIntrimOutPackets
	}
	return 0
}

func (m *DiameterNas) GetMethodList() string {
	if m != nil {
		return m.MethodList
	}
	return ""
}

func (m *DiameterNas) GetServerUsedList() string {
	if m != nil {
		return m.ServerUsedList
	}
	return ""
}

type DiameterNasList struct {
	AaanasId             string         `protobuf:"bytes,50,opt,name=aaanas_id,json=aaanasId,proto3" json:"aaanas_id,omitempty"`
	ListOfNas            []*DiameterNas `protobuf:"bytes,51,rep,name=list_of_nas,json=listOfNas,proto3" json:"list_of_nas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiameterNasList) Reset()         { *m = DiameterNasList{} }
func (m *DiameterNasList) String() string { return proto.CompactTextString(m) }
func (*DiameterNasList) ProtoMessage()    {}
func (*DiameterNasList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c5c5cf3e64bf30, []int{2}
}

func (m *DiameterNasList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiameterNasList.Unmarshal(m, b)
}
func (m *DiameterNasList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiameterNasList.Marshal(b, m, deterministic)
}
func (m *DiameterNasList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiameterNasList.Merge(m, src)
}
func (m *DiameterNasList) XXX_Size() int {
	return xxx_messageInfo_DiameterNasList.Size(m)
}
func (m *DiameterNasList) XXX_DiscardUnknown() {
	xxx_messageInfo_DiameterNasList.DiscardUnknown(m)
}

var xxx_messageInfo_DiameterNasList proto.InternalMessageInfo

func (m *DiameterNasList) GetAaanasId() string {
	if m != nil {
		return m.AaanasId
	}
	return ""
}

func (m *DiameterNasList) GetListOfNas() []*DiameterNas {
	if m != nil {
		return m.ListOfNas
	}
	return nil
}

func init() {
	proto.RegisterType((*DiameterNasList_KEYS)(nil), "cisco_ios_xr_aaa_diameter_oper.augment.diameter.nas_session.diameter_nas_list_KEYS")
	proto.RegisterType((*DiameterNas)(nil), "cisco_ios_xr_aaa_diameter_oper.augment.diameter.nas_session.diameter_nas")
	proto.RegisterType((*DiameterNasList)(nil), "cisco_ios_xr_aaa_diameter_oper.augment.diameter.nas_session.diameter_nas_list")
}

func init() { proto.RegisterFile("diameter_nas_list.proto", fileDescriptor_f7c5c5cf3e64bf30) }

var fileDescriptor_f7c5c5cf3e64bf30 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x15, 0x5a, 0x4a, 0x33, 0x29, 0x55, 0xb3, 0x85, 0xb2, 0xfc, 0x13, 0x51, 0xc4, 0x21,
	0x12, 0x92, 0x25, 0x1a, 0x6e, 0x1c, 0x10, 0x07, 0x0e, 0x16, 0x88, 0xa2, 0x58, 0x1c, 0x38, 0xad,
	0x06, 0x7b, 0xdb, 0xae, 0xc0, 0xbb, 0xd6, 0xce, 0x18, 0x21, 0xbe, 0x0a, 0x1f, 0x81, 0x2f, 0x89,
	0x76, 0x6d, 0xc7, 0x4e, 0x7d, 0xec, 0xd1, 0xef, 0xfd, 0xde, 0x9b, 0x95, 0x66, 0x0c, 0x8f, 0x0a,
	0x83, 0xa5, 0x66, 0xed, 0x95, 0x45, 0x52, 0x3f, 0x0d, 0x71, 0x52, 0x79, 0xc7, 0x4e, 0xbc, 0xcd,
	0x0d, 0xe5, 0x4e, 0x19, 0x47, 0xea, 0xb7, 0x57, 0x88, 0xa8, 0xb6, 0xa4, 0xab, 0xb4, 0x4f, 0xb0,
	0xbe, 0x2a, 0xb5, 0xe5, 0xa4, 0x53, 0x93, 0x90, 0x27, 0x4d, 0x64, 0x9c, 0x5d, 0x4a, 0x38, 0x1b,
	0xf5, 0xaa, 0x8f, 0x1f, 0xbe, 0x65, 0xcb, 0x7f, 0xfb, 0x70, 0x34, 0xb4, 0xc4, 0x4b, 0x38, 0x0e,
	0xe5, 0x6d, 0x52, 0x99, 0x42, 0x4e, 0x16, 0x93, 0xd5, 0x74, 0x73, 0x84, 0x88, 0x59, 0x23, 0xa6,
	0x85, 0x48, 0xe0, 0x74, 0x9b, 0x1a, 0xa0, 0x77, 0x22, 0x3a, 0xef, 0xac, 0x9e, 0x5f, 0xc3, 0x43,
	0xac, 0xf9, 0x5a, 0x5b, 0x36, 0x39, 0x72, 0xa0, 0x89, 0x91, 0x6b, 0x92, 0x7b, 0x8b, 0xc9, 0xea,
	0xfe, 0xe6, 0xc1, 0xae, 0x99, 0x45, 0x4f, 0xbc, 0x86, 0xa8, 0x3b, 0x6f, 0xfe, 0xec, 0x64, 0xf6,
	0x63, 0xe6, 0x74, 0xc7, 0x6b, 0x23, 0xaf, 0x60, 0x8e, 0x79, 0xee, 0x6a, 0xcb, 0xc6, 0x5e, 0x75,
	0xfc, 0xdd, 0xc8, 0x9f, 0xf4, 0x46, 0x0b, 0xbf, 0x81, 0xb3, 0x11, 0xac, 0x88, 0x5d, 0x25, 0x0f,
	0xda, 0x57, 0xdd, 0x48, 0x64, 0xec, 0xaa, 0x30, 0xa2, 0x08, 0xab, 0xb0, 0x56, 0xe7, 0xdc, 0x8d,
	0xb8, 0xd7, 0x8c, 0xe8, 0x8d, 0x76, 0xc4, 0x3b, 0x78, 0x36, 0x18, 0x61, 0x2c, 0x7b, 0x53, 0x2a,
	0x63, 0x55, 0x85, 0xf9, 0x0f, 0xcd, 0x24, 0x0f, 0x63, 0xee, 0x71, 0xcf, 0xa4, 0x11, 0x49, 0xed,
	0x97, 0x06, 0x10, 0xef, 0xe1, 0xf9, 0xb8, 0xc0, 0xd5, 0xbc, 0x6d, 0x98, 0xc6, 0x86, 0x27, 0x37,
	0x1b, 0x2e, 0x6a, 0xee, 0x2a, 0x5e, 0xc0, 0xac, 0xd4, 0x7c, 0xed, 0x8a, 0xb8, 0x76, 0x09, 0x71,
	0x47, 0xd0, 0x48, 0x9f, 0x0c, 0xb1, 0x58, 0xc1, 0x09, 0x69, 0xff, 0x4b, 0x7b, 0x55, 0x93, 0x6e,
	0xa9, 0x59, 0xa4, 0x8e, 0x1b, 0xfd, 0x2b, 0xe9, 0x48, 0x2e, 0xff, 0x4e, 0x60, 0x3e, 0x3a, 0x24,
	0xf1, 0x14, 0xa6, 0x88, 0x18, 0x3e, 0x4d, 0x21, 0xcf, 0x63, 0xf0, 0xb0, 0x11, 0xd2, 0x42, 0x18,
	0x98, 0xc5, 0x6b, 0x73, 0x97, 0x21, 0x20, 0xd7, 0x8b, 0xbd, 0xd5, 0xec, 0x3c, 0x4d, 0x6e, 0x71,
	0xcd, 0xc9, 0xf0, 0x05, 0x9b, 0x69, 0x68, 0xbf, 0xb8, 0xfc, 0x8c, 0xf4, 0xfd, 0x20, 0xfe, 0x29,
	0xeb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0x13, 0x4a, 0x50, 0x44, 0x03, 0x00, 0x00,
}
