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
// source: autorp_map_rp_bag.proto

package cisco_ios_xr_ipv4_autorp_oper_auto_rp_active_mapping_agent_rp_addresses_rp_address

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

type AutorpMapRpBag_KEYS struct {
	RpAddress            string   `protobuf:"bytes,1,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutorpMapRpBag_KEYS) Reset()         { *m = AutorpMapRpBag_KEYS{} }
func (m *AutorpMapRpBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*AutorpMapRpBag_KEYS) ProtoMessage()    {}
func (*AutorpMapRpBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4f5e9425a65ef50, []int{0}
}

func (m *AutorpMapRpBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutorpMapRpBag_KEYS.Unmarshal(m, b)
}
func (m *AutorpMapRpBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutorpMapRpBag_KEYS.Marshal(b, m, deterministic)
}
func (m *AutorpMapRpBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutorpMapRpBag_KEYS.Merge(m, src)
}
func (m *AutorpMapRpBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_AutorpMapRpBag_KEYS.Size(m)
}
func (m *AutorpMapRpBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AutorpMapRpBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AutorpMapRpBag_KEYS proto.InternalMessageInfo

func (m *AutorpMapRpBag_KEYS) GetRpAddress() string {
	if m != nil {
		return m.RpAddress
	}
	return ""
}

type AutorpMapRangeBag struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	ProtocolMode         string   `protobuf:"bytes,3,opt,name=protocol_mode,json=protocolMode,proto3" json:"protocol_mode,omitempty"`
	IsAdvertised         bool     `protobuf:"varint,4,opt,name=is_advertised,json=isAdvertised,proto3" json:"is_advertised,omitempty"`
	CreateType           uint32   `protobuf:"varint,5,opt,name=create_type,json=createType,proto3" json:"create_type,omitempty"`
	CheckPointObjectId   uint32   `protobuf:"varint,6,opt,name=check_point_object_id,json=checkPointObjectId,proto3" json:"check_point_object_id,omitempty"`
	Uptime               uint64   `protobuf:"varint,7,opt,name=uptime,proto3" json:"uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutorpMapRangeBag) Reset()         { *m = AutorpMapRangeBag{} }
func (m *AutorpMapRangeBag) String() string { return proto.CompactTextString(m) }
func (*AutorpMapRangeBag) ProtoMessage()    {}
func (*AutorpMapRangeBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4f5e9425a65ef50, []int{1}
}

func (m *AutorpMapRangeBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutorpMapRangeBag.Unmarshal(m, b)
}
func (m *AutorpMapRangeBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutorpMapRangeBag.Marshal(b, m, deterministic)
}
func (m *AutorpMapRangeBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutorpMapRangeBag.Merge(m, src)
}
func (m *AutorpMapRangeBag) XXX_Size() int {
	return xxx_messageInfo_AutorpMapRangeBag.Size(m)
}
func (m *AutorpMapRangeBag) XXX_DiscardUnknown() {
	xxx_messageInfo_AutorpMapRangeBag.DiscardUnknown(m)
}

var xxx_messageInfo_AutorpMapRangeBag proto.InternalMessageInfo

func (m *AutorpMapRangeBag) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *AutorpMapRangeBag) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *AutorpMapRangeBag) GetProtocolMode() string {
	if m != nil {
		return m.ProtocolMode
	}
	return ""
}

func (m *AutorpMapRangeBag) GetIsAdvertised() bool {
	if m != nil {
		return m.IsAdvertised
	}
	return false
}

func (m *AutorpMapRangeBag) GetCreateType() uint32 {
	if m != nil {
		return m.CreateType
	}
	return 0
}

func (m *AutorpMapRangeBag) GetCheckPointObjectId() uint32 {
	if m != nil {
		return m.CheckPointObjectId
	}
	return 0
}

func (m *AutorpMapRangeBag) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

type AutorpMapRpBag struct {
	RpAddressXr          string               `protobuf:"bytes,50,opt,name=rp_address_xr,json=rpAddressXr,proto3" json:"rp_address_xr,omitempty"`
	ExpiryTime           uint64               `protobuf:"varint,51,opt,name=expiry_time,json=expiryTime,proto3" json:"expiry_time,omitempty"`
	PimVersion           uint32               `protobuf:"varint,52,opt,name=pim_version,json=pimVersion,proto3" json:"pim_version,omitempty"`
	Range                []*AutorpMapRangeBag `protobuf:"bytes,53,rep,name=range,proto3" json:"range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AutorpMapRpBag) Reset()         { *m = AutorpMapRpBag{} }
func (m *AutorpMapRpBag) String() string { return proto.CompactTextString(m) }
func (*AutorpMapRpBag) ProtoMessage()    {}
func (*AutorpMapRpBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4f5e9425a65ef50, []int{2}
}

func (m *AutorpMapRpBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutorpMapRpBag.Unmarshal(m, b)
}
func (m *AutorpMapRpBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutorpMapRpBag.Marshal(b, m, deterministic)
}
func (m *AutorpMapRpBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutorpMapRpBag.Merge(m, src)
}
func (m *AutorpMapRpBag) XXX_Size() int {
	return xxx_messageInfo_AutorpMapRpBag.Size(m)
}
func (m *AutorpMapRpBag) XXX_DiscardUnknown() {
	xxx_messageInfo_AutorpMapRpBag.DiscardUnknown(m)
}

var xxx_messageInfo_AutorpMapRpBag proto.InternalMessageInfo

func (m *AutorpMapRpBag) GetRpAddressXr() string {
	if m != nil {
		return m.RpAddressXr
	}
	return ""
}

func (m *AutorpMapRpBag) GetExpiryTime() uint64 {
	if m != nil {
		return m.ExpiryTime
	}
	return 0
}

func (m *AutorpMapRpBag) GetPimVersion() uint32 {
	if m != nil {
		return m.PimVersion
	}
	return 0
}

func (m *AutorpMapRpBag) GetRange() []*AutorpMapRangeBag {
	if m != nil {
		return m.Range
	}
	return nil
}

func init() {
	proto.RegisterType((*AutorpMapRpBag_KEYS)(nil), "cisco_ios_xr_ipv4_autorp_oper.auto_rp.active.mapping_agent.rp_addresses.rp_address.autorp_map_rp_bag_KEYS")
	proto.RegisterType((*AutorpMapRangeBag)(nil), "cisco_ios_xr_ipv4_autorp_oper.auto_rp.active.mapping_agent.rp_addresses.rp_address.autorp_map_range_bag")
	proto.RegisterType((*AutorpMapRpBag)(nil), "cisco_ios_xr_ipv4_autorp_oper.auto_rp.active.mapping_agent.rp_addresses.rp_address.autorp_map_rp_bag")
}

func init() { proto.RegisterFile("autorp_map_rp_bag.proto", fileDescriptor_d4f5e9425a65ef50) }

var fileDescriptor_d4f5e9425a65ef50 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0x6f, 0xd4, 0x40,
	0x10, 0x95, 0xf3, 0x71, 0x90, 0xbd, 0xbb, 0x82, 0x15, 0x84, 0x6d, 0x10, 0xd6, 0xd1, 0xb8, 0xb2,
	0x44, 0x12, 0x44, 0x9d, 0x82, 0x02, 0x01, 0x02, 0x99, 0x08, 0x41, 0x35, 0xda, 0xb3, 0x07, 0xdf,
	0x40, 0xec, 0x1d, 0xcd, 0x6e, 0x4e, 0xbe, 0x86, 0x9e, 0x9f, 0x4c, 0x87, 0xbc, 0x6b, 0x92, 0x48,
	0xa1, 0x4c, 0x37, 0xf3, 0xe6, 0xbd, 0x99, 0x9d, 0x37, 0xab, 0x9e, 0xda, 0xab, 0xe0, 0x84, 0xa1,
	0xb3, 0x0c, 0xc2, 0xb0, 0xb6, 0x6d, 0xc9, 0xe2, 0x82, 0xd3, 0x55, 0x4d, 0xbe, 0x76, 0x40, 0xce,
	0xc3, 0x20, 0x40, 0xbc, 0x3d, 0x83, 0x89, 0xea, 0x18, 0xa5, 0x1c, 0x63, 0x10, 0x2e, 0x6d, 0x1d,
	0x68, 0x8b, 0x65, 0x67, 0x99, 0xa9, 0x6f, 0xc1, 0xb6, 0xd8, 0x87, 0x52, 0x18, 0x6c, 0xd3, 0x08,
	0x7a, 0x8f, 0xfe, 0x56, 0xb2, 0x7a, 0xad, 0x8e, 0xef, 0x8c, 0x83, 0x77, 0x6f, 0xbe, 0x7d, 0xd6,
	0xcf, 0x94, 0xba, 0xe1, 0x99, 0x2c, 0xcf, 0x8a, 0xa3, 0xea, 0x48, 0xf8, 0x7c, 0x12, 0xfe, 0xde,
	0x53, 0x8f, 0x6f, 0x2b, 0x6d, 0xdf, 0xe2, 0x28, 0xd6, 0xc7, 0x6a, 0xc6, 0x82, 0xdf, 0x69, 0x98,
	0x34, 0x53, 0xa6, 0x5f, 0xa8, 0x65, 0x8a, 0xe0, 0x12, 0xfb, 0x36, 0x6c, 0xcc, 0x5e, 0x9e, 0x15,
	0xcb, 0x6a, 0x91, 0xc0, 0xf7, 0x11, 0x4b, 0x24, 0x17, 0x5c, 0xed, 0x2e, 0xa1, 0x73, 0x0d, 0x9a,
	0xfd, 0xd8, 0x63, 0xf1, 0x0f, 0xfc, 0xe0, 0x1a, 0x1c, 0x49, 0xe4, 0xc1, 0x36, 0x5b, 0x94, 0x40,
	0x1e, 0x1b, 0x73, 0x90, 0x67, 0xc5, 0xc3, 0x6a, 0x41, 0xfe, 0xfc, 0x1a, 0xd3, 0xcf, 0xd5, 0xbc,
	0x16, 0xb4, 0x01, 0x21, 0xec, 0x18, 0xcd, 0x61, 0x1c, 0xa6, 0x12, 0x74, 0xb1, 0x63, 0xd4, 0x2f,
	0xd5, 0x93, 0x7a, 0x83, 0xf5, 0x4f, 0x60, 0x47, 0x7d, 0x00, 0xb7, 0xfe, 0x81, 0x75, 0x00, 0x6a,
	0xcc, 0x2c, 0x52, 0x75, 0x2c, 0x7e, 0x1a, 0x6b, 0x1f, 0x63, 0xe9, 0x6d, 0x33, 0xae, 0x76, 0xc5,
	0x81, 0x3a, 0x34, 0x0f, 0xf2, 0xac, 0x38, 0xa8, 0xa6, 0x6c, 0xf5, 0x27, 0x53, 0x8f, 0xee, 0xb8,
	0xa8, 0x57, 0x6a, 0x79, 0x63, 0x20, 0x0c, 0x62, 0x4e, 0xe2, 0x2e, 0xf3, 0x6b, 0x0f, 0xbf, 0xca,
	0xf8, 0x4a, 0x1c, 0x98, 0x64, 0x07, 0xb1, 0xed, 0x69, 0x6c, 0xab, 0x12, 0x74, 0x41, 0x1d, 0x8e,
	0x04, 0xa6, 0x0e, 0xb6, 0x28, 0x9e, 0x5c, 0x6f, 0xce, 0xd2, 0x1a, 0x4c, 0xdd, 0x97, 0x84, 0xe8,
	0x5f, 0xea, 0x30, 0x7a, 0x6f, 0x5e, 0xe5, 0xfb, 0xc5, 0xfc, 0x64, 0x53, 0xde, 0xff, 0x27, 0x29,
	0xff, 0x77, 0xe7, 0x2a, 0x8d, 0x5d, 0xcf, 0xe2, 0x69, 0x4e, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0xd9, 0x75, 0xc5, 0xb6, 0x02, 0x00, 0x00,
}
