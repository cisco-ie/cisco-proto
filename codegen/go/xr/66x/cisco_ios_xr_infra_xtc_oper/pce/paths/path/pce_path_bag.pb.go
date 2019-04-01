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
// source: pce_path_bag.proto

package cisco_ios_xr_infra_xtc_oper_pce_paths_path

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

type PcePathBag_KEYS struct {
	Af                   uint32   `protobuf:"varint,1,opt,name=af,proto3" json:"af,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Destination          string   `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcePathBag_KEYS) Reset()         { *m = PcePathBag_KEYS{} }
func (m *PcePathBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PcePathBag_KEYS) ProtoMessage()    {}
func (*PcePathBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bad7b98055410ed, []int{0}
}

func (m *PcePathBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePathBag_KEYS.Unmarshal(m, b)
}
func (m *PcePathBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePathBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PcePathBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePathBag_KEYS.Merge(m, src)
}
func (m *PcePathBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PcePathBag_KEYS.Size(m)
}
func (m *PcePathBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePathBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PcePathBag_KEYS proto.InternalMessageInfo

func (m *PcePathBag_KEYS) GetAf() uint32 {
	if m != nil {
		return m.Af
	}
	return 0
}

func (m *PcePathBag_KEYS) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *PcePathBag_KEYS) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type PceAddrBag struct {
	AddressFamily        uint32   `protobuf:"varint,1,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	Ipv4Prefix           string   `protobuf:"bytes,2,opt,name=ipv4_prefix,json=ipv4Prefix,proto3" json:"ipv4_prefix,omitempty"`
	Ipv6Prefix           string   `protobuf:"bytes,3,opt,name=ipv6_prefix,json=ipv6Prefix,proto3" json:"ipv6_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceAddrBag) Reset()         { *m = PceAddrBag{} }
func (m *PceAddrBag) String() string { return proto.CompactTextString(m) }
func (*PceAddrBag) ProtoMessage()    {}
func (*PceAddrBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bad7b98055410ed, []int{1}
}

func (m *PceAddrBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceAddrBag.Unmarshal(m, b)
}
func (m *PceAddrBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceAddrBag.Marshal(b, m, deterministic)
}
func (m *PceAddrBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceAddrBag.Merge(m, src)
}
func (m *PceAddrBag) XXX_Size() int {
	return xxx_messageInfo_PceAddrBag.Size(m)
}
func (m *PceAddrBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceAddrBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceAddrBag proto.InternalMessageInfo

func (m *PceAddrBag) GetAddressFamily() uint32 {
	if m != nil {
		return m.AddressFamily
	}
	return 0
}

func (m *PceAddrBag) GetIpv4Prefix() string {
	if m != nil {
		return m.Ipv4Prefix
	}
	return ""
}

func (m *PceAddrBag) GetIpv6Prefix() string {
	if m != nil {
		return m.Ipv6Prefix
	}
	return ""
}

type PceIpAddrType struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceIpAddrType) Reset()         { *m = PceIpAddrType{} }
func (m *PceIpAddrType) String() string { return proto.CompactTextString(m) }
func (*PceIpAddrType) ProtoMessage()    {}
func (*PceIpAddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bad7b98055410ed, []int{2}
}

func (m *PceIpAddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIpAddrType.Unmarshal(m, b)
}
func (m *PceIpAddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIpAddrType.Marshal(b, m, deterministic)
}
func (m *PceIpAddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIpAddrType.Merge(m, src)
}
func (m *PceIpAddrType) XXX_Size() int {
	return xxx_messageInfo_PceIpAddrType.Size(m)
}
func (m *PceIpAddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIpAddrType.DiscardUnknown(m)
}

var xxx_messageInfo_PceIpAddrType proto.InternalMessageInfo

func (m *PceIpAddrType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PceIpAddrType) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *PceIpAddrType) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type PcePathBag struct {
	Hops                 []*PceAddrBag  `protobuf:"bytes,50,rep,name=hops,proto3" json:"hops,omitempty"`
	Cost                 uint64         `protobuf:"varint,51,opt,name=cost,proto3" json:"cost,omitempty"`
	SourceXr             *PceIpAddrType `protobuf:"bytes,52,opt,name=source_xr,json=sourceXr,proto3" json:"source_xr,omitempty"`
	DestinationXr        *PceIpAddrType `protobuf:"bytes,53,opt,name=destination_xr,json=destinationXr,proto3" json:"destination_xr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PcePathBag) Reset()         { *m = PcePathBag{} }
func (m *PcePathBag) String() string { return proto.CompactTextString(m) }
func (*PcePathBag) ProtoMessage()    {}
func (*PcePathBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bad7b98055410ed, []int{3}
}

func (m *PcePathBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePathBag.Unmarshal(m, b)
}
func (m *PcePathBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePathBag.Marshal(b, m, deterministic)
}
func (m *PcePathBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePathBag.Merge(m, src)
}
func (m *PcePathBag) XXX_Size() int {
	return xxx_messageInfo_PcePathBag.Size(m)
}
func (m *PcePathBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePathBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePathBag proto.InternalMessageInfo

func (m *PcePathBag) GetHops() []*PceAddrBag {
	if m != nil {
		return m.Hops
	}
	return nil
}

func (m *PcePathBag) GetCost() uint64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *PcePathBag) GetSourceXr() *PceIpAddrType {
	if m != nil {
		return m.SourceXr
	}
	return nil
}

func (m *PcePathBag) GetDestinationXr() *PceIpAddrType {
	if m != nil {
		return m.DestinationXr
	}
	return nil
}

func init() {
	proto.RegisterType((*PcePathBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_oper.pce.paths.path.pce_path_bag_KEYS")
	proto.RegisterType((*PceAddrBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.paths.path.pce_addr_bag")
	proto.RegisterType((*PceIpAddrType)(nil), "cisco_ios_xr_infra_xtc_oper.pce.paths.path.pce_ip_addr_type")
	proto.RegisterType((*PcePathBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.paths.path.pce_path_bag")
}

func init() { proto.RegisterFile("pce_path_bag.proto", fileDescriptor_9bad7b98055410ed) }

var fileDescriptor_9bad7b98055410ed = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4b, 0xfb, 0x30,
	0x1c, 0xc5, 0x69, 0x37, 0xf6, 0xfb, 0xed, 0x5b, 0x37, 0x34, 0x07, 0xed, 0xcd, 0x52, 0x10, 0x8a,
	0x87, 0x1e, 0xb6, 0x39, 0x3c, 0x78, 0xd5, 0x8b, 0x22, 0xd2, 0x5d, 0xb6, 0x83, 0x84, 0x2c, 0x4b,
	0x5c, 0xc0, 0x35, 0x21, 0x89, 0xda, 0xfd, 0x57, 0xfe, 0x89, 0x92, 0x34, 0x93, 0x1e, 0x15, 0xbc,
	0x94, 0xd7, 0xd7, 0xc7, 0xe7, 0xd1, 0x97, 0x00, 0x52, 0x94, 0x61, 0x45, 0xec, 0x16, 0xaf, 0xc9,
	0x4b, 0xa9, 0xb4, 0xb4, 0x12, 0x5d, 0x52, 0x61, 0xa8, 0xc4, 0x42, 0x1a, 0xdc, 0x68, 0x2c, 0x6a,
	0xae, 0x09, 0x6e, 0x2c, 0xc5, 0x52, 0x31, 0x5d, 0x2a, 0xca, 0x4a, 0x97, 0x37, 0xfe, 0x99, 0x3f,
	0xc3, 0x49, 0x97, 0x80, 0xef, 0x6f, 0x57, 0x0b, 0x34, 0x86, 0x98, 0xf0, 0x34, 0xca, 0xa2, 0x62,
	0x54, 0xc5, 0x84, 0xa3, 0x53, 0x18, 0x18, 0xf9, 0xa6, 0x29, 0x4b, 0xe3, 0x2c, 0x2a, 0x86, 0x55,
	0x78, 0x43, 0x19, 0x24, 0x1b, 0x66, 0xac, 0xa8, 0x89, 0x15, 0xb2, 0x4e, 0x7b, 0xfe, 0x63, 0xd7,
	0xca, 0x3f, 0xe0, 0xc8, 0xe1, 0xc9, 0x66, 0xa3, 0x1d, 0x1e, 0x5d, 0xc0, 0xd8, 0x69, 0x66, 0x0c,
	0xe6, 0x64, 0x27, 0x5e, 0xf7, 0xa1, 0x65, 0x14, 0xdc, 0x3b, 0x6f, 0xa2, 0x73, 0x48, 0x84, 0x7a,
	0x9f, 0x61, 0xa5, 0x19, 0x17, 0x4d, 0x68, 0x05, 0x67, 0x3d, 0x79, 0x27, 0x04, 0xe6, 0x87, 0x40,
	0xef, 0x3b, 0x30, 0x6f, 0x03, 0xf9, 0x02, 0x8e, 0x5d, 0xb1, 0x50, 0x6d, 0xb7, 0xdd, 0x2b, 0x86,
	0xce, 0xe0, 0x1f, 0xe1, 0xb8, 0x26, 0x3b, 0xe6, 0x5b, 0x87, 0xd5, 0x80, 0xf0, 0x47, 0xb2, 0x63,
	0x08, 0x41, 0xdf, 0xb1, 0x43, 0x8f, 0xd7, 0xc1, 0x9b, 0x07, 0xb4, 0xd7, 0xf9, 0x67, 0xdc, 0xfe,
	0xce, 0x61, 0x2d, 0xf4, 0x00, 0xfd, 0xad, 0x54, 0x26, 0x9d, 0x64, 0xbd, 0x22, 0x99, 0x5c, 0x97,
	0x3f, 0x1f, 0xbe, 0xec, 0xce, 0x52, 0x79, 0x8a, 0xab, 0xa4, 0xd2, 0xd8, 0x74, 0x9a, 0x45, 0x45,
	0xbf, 0xf2, 0x1a, 0xad, 0x60, 0xd8, 0x8e, 0x8d, 0x1b, 0x9d, 0xce, 0xb2, 0xa8, 0x48, 0x26, 0x37,
	0xbf, 0xad, 0xe9, 0x8e, 0x50, 0xfd, 0x6f, 0x71, 0x4b, 0x8d, 0x28, 0x8c, 0x3b, 0x47, 0xe5, 0xf8,
	0x57, 0x7f, 0xc0, 0x1f, 0x75, 0x98, 0x4b, 0xbd, 0x1e, 0xf8, 0x2b, 0x39, 0xfd, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xf3, 0x25, 0x68, 0xf0, 0xa8, 0x02, 0x00, 0x00,
}