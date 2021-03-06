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
// source: client_entry.proto

package cisco_ios_xr_lpts_pa_oper_lpts_pa_entries_entry

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

type ClientEntry_KEYS struct {
	Entry                string   `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientEntry_KEYS) Reset()         { *m = ClientEntry_KEYS{} }
func (m *ClientEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*ClientEntry_KEYS) ProtoMessage()    {}
func (*ClientEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb18fde65cb7a314, []int{0}
}

func (m *ClientEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEntry_KEYS.Unmarshal(m, b)
}
func (m *ClientEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *ClientEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEntry_KEYS.Merge(m, src)
}
func (m *ClientEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_ClientEntry_KEYS.Size(m)
}
func (m *ClientEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEntry_KEYS proto.InternalMessageInfo

func (m *ClientEntry_KEYS) GetEntry() string {
	if m != nil {
		return m.Entry
	}
	return ""
}

type ClientEntry struct {
	Flags                uint32   `protobuf:"varint,50,opt,name=flags,proto3" json:"flags,omitempty"`
	OpenFlags            uint32   `protobuf:"varint,51,opt,name=open_flags,json=openFlags,proto3" json:"open_flags,omitempty"`
	Location             uint32   `protobuf:"varint,52,opt,name=location,proto3" json:"location,omitempty"`
	ClientId             uint32   `protobuf:"varint,53,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Times                string   `protobuf:"bytes,54,opt,name=times,proto3" json:"times,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientEntry) Reset()         { *m = ClientEntry{} }
func (m *ClientEntry) String() string { return proto.CompactTextString(m) }
func (*ClientEntry) ProtoMessage()    {}
func (*ClientEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb18fde65cb7a314, []int{1}
}

func (m *ClientEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEntry.Unmarshal(m, b)
}
func (m *ClientEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEntry.Marshal(b, m, deterministic)
}
func (m *ClientEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEntry.Merge(m, src)
}
func (m *ClientEntry) XXX_Size() int {
	return xxx_messageInfo_ClientEntry.Size(m)
}
func (m *ClientEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEntry proto.InternalMessageInfo

func (m *ClientEntry) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *ClientEntry) GetOpenFlags() uint32 {
	if m != nil {
		return m.OpenFlags
	}
	return 0
}

func (m *ClientEntry) GetLocation() uint32 {
	if m != nil {
		return m.Location
	}
	return 0
}

func (m *ClientEntry) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *ClientEntry) GetTimes() string {
	if m != nil {
		return m.Times
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientEntry_KEYS)(nil), "cisco_ios_xr_lpts_pa_oper.lpts_pa.entries.entry.client_entry_KEYS")
	proto.RegisterType((*ClientEntry)(nil), "cisco_ios_xr_lpts_pa_oper.lpts_pa.entries.entry.client_entry")
}

func init() { proto.RegisterFile("client_entry.proto", fileDescriptor_cb18fde65cb7a314) }

var fileDescriptor_cb18fde65cb7a314 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0x89, 0x4f, 0xcd, 0x2b, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2,
	0x4f, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0xcf, 0x29, 0x28,
	0x29, 0x8e, 0x2f, 0x48, 0x8c, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0x83, 0x72, 0xf4, 0x40, 0xca, 0x33,
	0x53, 0x8b, 0xc1, 0x74, 0xa5, 0x92, 0x26, 0x97, 0x20, 0xb2, 0x31, 0xf1, 0xde, 0xae, 0x91, 0xc1,
	0x42, 0x22, 0x5c, 0xac, 0x60, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0xa3, 0x34,
	0x89, 0x91, 0x8b, 0x07, 0x59, 0x2d, 0x48, 0x59, 0x5a, 0x4e, 0x62, 0x7a, 0xb1, 0x84, 0x91, 0x02,
	0xa3, 0x06, 0x6f, 0x10, 0x84, 0x23, 0x24, 0xcb, 0xc5, 0x95, 0x5f, 0x90, 0x9a, 0x17, 0x0f, 0x91,
	0x32, 0x06, 0x4b, 0x71, 0x82, 0x44, 0xdc, 0xc0, 0xd2, 0x52, 0x5c, 0x1c, 0x39, 0xf9, 0xc9, 0x89,
	0x25, 0x99, 0xf9, 0x79, 0x12, 0x26, 0x60, 0x49, 0x38, 0x5f, 0x48, 0x9a, 0x8b, 0x13, 0x6a, 0x41,
	0x66, 0x8a, 0x84, 0x29, 0x44, 0x12, 0x22, 0xe0, 0x99, 0x02, 0xb2, 0xad, 0x24, 0x33, 0x37, 0xb5,
	0x58, 0xc2, 0x0c, 0xe2, 0x28, 0x30, 0x27, 0x89, 0x0d, 0xec, 0x6f, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x11, 0x93, 0x3a, 0xa8, 0x0d, 0x01, 0x00, 0x00,
}
