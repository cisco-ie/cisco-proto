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
// source: pim_rpf_policy_test_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_default_context_route_policy_route_policy_tests_route_policy_test

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

type PimRpfPolicyTestBag_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	GroupAddress         string   `protobuf:"bytes,2,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	ExtComm              string   `protobuf:"bytes,3,opt,name=ext_comm,json=extComm,proto3" json:"ext_comm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimRpfPolicyTestBag_KEYS) Reset()         { *m = PimRpfPolicyTestBag_KEYS{} }
func (m *PimRpfPolicyTestBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimRpfPolicyTestBag_KEYS) ProtoMessage()    {}
func (*PimRpfPolicyTestBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a2a94ff48f1cb7, []int{0}
}

func (m *PimRpfPolicyTestBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfPolicyTestBag_KEYS.Unmarshal(m, b)
}
func (m *PimRpfPolicyTestBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfPolicyTestBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimRpfPolicyTestBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfPolicyTestBag_KEYS.Merge(m, src)
}
func (m *PimRpfPolicyTestBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimRpfPolicyTestBag_KEYS.Size(m)
}
func (m *PimRpfPolicyTestBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfPolicyTestBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfPolicyTestBag_KEYS proto.InternalMessageInfo

func (m *PimRpfPolicyTestBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *PimRpfPolicyTestBag_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
	}
	return ""
}

func (m *PimRpfPolicyTestBag_KEYS) GetExtComm() string {
	if m != nil {
		return m.ExtComm
	}
	return ""
}

type PimAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimAddrtype) Reset()         { *m = PimAddrtype{} }
func (m *PimAddrtype) String() string { return proto.CompactTextString(m) }
func (*PimAddrtype) ProtoMessage()    {}
func (*PimAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a2a94ff48f1cb7, []int{1}
}

func (m *PimAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimAddrtype.Unmarshal(m, b)
}
func (m *PimAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimAddrtype.Marshal(b, m, deterministic)
}
func (m *PimAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimAddrtype.Merge(m, src)
}
func (m *PimAddrtype) XXX_Size() int {
	return xxx_messageInfo_PimAddrtype.Size(m)
}
func (m *PimAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_PimAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_PimAddrtype proto.InternalMessageInfo

func (m *PimAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PimAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *PimAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type PimRpfPolicyTestBag struct {
	PolicyName           string       `protobuf:"bytes,50,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	SourceAddressXr      *PimAddrtype `protobuf:"bytes,51,opt,name=source_address_xr,json=sourceAddressXr,proto3" json:"source_address_xr,omitempty"`
	GroupAddressXr       *PimAddrtype `protobuf:"bytes,52,opt,name=group_address_xr,json=groupAddressXr,proto3" json:"group_address_xr,omitempty"`
	ExtendedCommunityRt  string       `protobuf:"bytes,53,opt,name=extended_community_rt,json=extendedCommunityRt,proto3" json:"extended_community_rt,omitempty"`
	Pass                 bool         `protobuf:"varint,54,opt,name=pass,proto3" json:"pass,omitempty"`
	DefaultTable         bool         `protobuf:"varint,55,opt,name=default_table,json=defaultTable,proto3" json:"default_table,omitempty"`
	VrfName              string       `protobuf:"bytes,56,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Afi                  uint32       `protobuf:"varint,57,opt,name=afi,proto3" json:"afi,omitempty"`
	Safi                 uint32       `protobuf:"varint,58,opt,name=safi,proto3" json:"safi,omitempty"`
	TableName            string       `protobuf:"bytes,59,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	TableExists          bool         `protobuf:"varint,60,opt,name=table_exists,json=tableExists,proto3" json:"table_exists,omitempty"`
	TableActive          bool         `protobuf:"varint,61,opt,name=table_active,json=tableActive,proto3" json:"table_active,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimRpfPolicyTestBag) Reset()         { *m = PimRpfPolicyTestBag{} }
func (m *PimRpfPolicyTestBag) String() string { return proto.CompactTextString(m) }
func (*PimRpfPolicyTestBag) ProtoMessage()    {}
func (*PimRpfPolicyTestBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a2a94ff48f1cb7, []int{2}
}

func (m *PimRpfPolicyTestBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfPolicyTestBag.Unmarshal(m, b)
}
func (m *PimRpfPolicyTestBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfPolicyTestBag.Marshal(b, m, deterministic)
}
func (m *PimRpfPolicyTestBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfPolicyTestBag.Merge(m, src)
}
func (m *PimRpfPolicyTestBag) XXX_Size() int {
	return xxx_messageInfo_PimRpfPolicyTestBag.Size(m)
}
func (m *PimRpfPolicyTestBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfPolicyTestBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfPolicyTestBag proto.InternalMessageInfo

func (m *PimRpfPolicyTestBag) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *PimRpfPolicyTestBag) GetSourceAddressXr() *PimAddrtype {
	if m != nil {
		return m.SourceAddressXr
	}
	return nil
}

func (m *PimRpfPolicyTestBag) GetGroupAddressXr() *PimAddrtype {
	if m != nil {
		return m.GroupAddressXr
	}
	return nil
}

func (m *PimRpfPolicyTestBag) GetExtendedCommunityRt() string {
	if m != nil {
		return m.ExtendedCommunityRt
	}
	return ""
}

func (m *PimRpfPolicyTestBag) GetPass() bool {
	if m != nil {
		return m.Pass
	}
	return false
}

func (m *PimRpfPolicyTestBag) GetDefaultTable() bool {
	if m != nil {
		return m.DefaultTable
	}
	return false
}

func (m *PimRpfPolicyTestBag) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *PimRpfPolicyTestBag) GetAfi() uint32 {
	if m != nil {
		return m.Afi
	}
	return 0
}

func (m *PimRpfPolicyTestBag) GetSafi() uint32 {
	if m != nil {
		return m.Safi
	}
	return 0
}

func (m *PimRpfPolicyTestBag) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *PimRpfPolicyTestBag) GetTableExists() bool {
	if m != nil {
		return m.TableExists
	}
	return false
}

func (m *PimRpfPolicyTestBag) GetTableActive() bool {
	if m != nil {
		return m.TableActive
	}
	return false
}

func init() {
	proto.RegisterType((*PimRpfPolicyTestBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.route_policy.route_policy_tests.route_policy_test.pim_rpf_policy_test_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.route_policy.route_policy_tests.route_policy_test.pim_addrtype")
	proto.RegisterType((*PimRpfPolicyTestBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.route_policy.route_policy_tests.route_policy_test.pim_rpf_policy_test_bag")
}

func init() { proto.RegisterFile("pim_rpf_policy_test_bag.proto", fileDescriptor_19a2a94ff48f1cb7) }

var fileDescriptor_19a2a94ff48f1cb7 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x15, 0x0a, 0xfb, 0x67, 0xda, 0x2e, 0x8b, 0x11, 0xda, 0x20, 0xb1, 0xa2, 0x14, 0x21,
	0xf5, 0x94, 0x43, 0x77, 0x29, 0xff, 0x0f, 0xab, 0xd5, 0x9e, 0x90, 0x38, 0x04, 0x0e, 0xcb, 0xc9,
	0x72, 0x93, 0xc9, 0xca, 0x52, 0x53, 0x5b, 0xb6, 0x13, 0xa5, 0xdc, 0xf9, 0x18, 0xdc, 0x10, 0x9f,
	0x13, 0x79, 0x9c, 0x94, 0x46, 0x68, 0xaf, 0x70, 0x73, 0x7e, 0xf3, 0x3c, 0xf3, 0x3a, 0x7e, 0x2a,
	0x9c, 0x6a, 0x59, 0x72, 0xa3, 0x0b, 0xae, 0xd5, 0x4a, 0x66, 0x1b, 0xee, 0xd0, 0x3a, 0xbe, 0x14,
	0x37, 0x89, 0x36, 0xca, 0x29, 0xb6, 0xca, 0xa4, 0xcd, 0x14, 0x97, 0xca, 0xf2, 0xc6, 0x70, 0xa9,
	0xeb, 0x73, 0xee, 0x2f, 0x28, 0x8d, 0x26, 0x91, 0xba, 0x5e, 0xf8, 0xaf, 0x44, 0x64, 0x4e, 0xd6,
	0x98, 0xe4, 0x58, 0x88, 0x6a, 0xe5, 0x78, 0xa6, 0xd6, 0x0e, 0x1b, 0x97, 0x18, 0x55, 0x39, 0x6c,
	0xfb, 0xf6, 0x3e, 0x68, 0x88, 0xfd, 0x1b, 0x4d, 0xbf, 0x47, 0xf0, 0xe4, 0x16, 0x3f, 0xfc, 0xe3,
	0xd5, 0xd7, 0xcf, 0xec, 0x05, 0x1c, 0x59, 0x55, 0x99, 0x0c, 0xb9, 0xc8, 0x73, 0x83, 0xd6, 0xc6,
	0xd1, 0x24, 0x9a, 0x1d, 0xa6, 0xe3, 0x40, 0x2f, 0x02, 0x64, 0xcf, 0x61, 0x7c, 0x63, 0x54, 0xa5,
	0xb7, 0xaa, 0x3b, 0xa4, 0x1a, 0x11, 0xec, 0x44, 0x8f, 0xe1, 0x00, 0x1b, 0xef, 0xb6, 0x2c, 0xe3,
	0x01, 0xd5, 0xf7, 0xb1, 0x71, 0x97, 0xaa, 0x2c, 0xa7, 0x25, 0x8c, 0xbc, 0x0d, 0x7f, 0xdb, 0x6d,
	0x34, 0xb2, 0x13, 0xd8, 0x17, 0x05, 0x5f, 0x8b, 0x12, 0xdb, 0x79, 0x7b, 0xa2, 0xf8, 0x24, 0x4a,
	0x64, 0xcf, 0x60, 0x44, 0x3b, 0xe9, 0xcf, 0x19, 0x7a, 0xd6, 0x8d, 0x09, 0x92, 0xc5, 0x56, 0x32,
	0xd8, 0x4a, 0x16, 0xad, 0x64, 0xfa, 0xe3, 0x1e, 0x9c, 0xdc, 0xf2, 0xb3, 0xd9, 0x53, 0x18, 0xb6,
	0x88, 0xc6, 0xcf, 0xe9, 0x36, 0x04, 0x44, 0x16, 0x7e, 0x45, 0xf0, 0xa0, 0xbf, 0x13, 0xde, 0x98,
	0xf8, 0x6c, 0x12, 0xcd, 0x86, 0xf3, 0x6f, 0xc9, 0xbf, 0x7c, 0xbe, 0x64, 0x77, 0x67, 0xe9, 0xfd,
	0xde, 0x93, 0x5c, 0x1b, 0xf6, 0x33, 0x82, 0xe3, 0xde, 0xab, 0x78, 0x9f, 0xe7, 0xff, 0xdd, 0xe7,
	0xd1, 0x6e, 0x28, 0xae, 0x0d, 0x9b, 0xc3, 0x23, 0x6c, 0x1c, 0xae, 0x73, 0xcc, 0x29, 0x1b, 0xd5,
	0x5a, 0xba, 0x0d, 0x37, 0x2e, 0x7e, 0x49, 0xab, 0x7f, 0xd8, 0x15, 0x2f, 0xbb, 0x5a, 0xea, 0x18,
	0x83, 0xbb, 0x5a, 0x58, 0x1b, 0x2f, 0x26, 0xd1, 0xec, 0x20, 0xa5, 0xb3, 0xcf, 0x60, 0xe7, 0xd4,
	0x89, 0xe5, 0x0a, 0xe3, 0x57, 0x54, 0x1c, 0xb5, 0xf0, 0x8b, 0x67, 0x3e, 0x83, 0xb5, 0x69, 0x93,
	0xf5, 0x3a, 0x64, 0xb0, 0x36, 0x21, 0x5a, 0xc7, 0x30, 0x10, 0x85, 0x8c, 0xdf, 0x4c, 0xa2, 0xd9,
	0x38, 0xf5, 0x47, 0x3f, 0xc5, 0x7a, 0xf4, 0x96, 0x10, 0x9d, 0xd9, 0x29, 0x00, 0x75, 0x0f, 0x2d,
	0xde, 0x51, 0x8b, 0x43, 0x22, 0x5d, 0x3e, 0x43, 0x19, 0x1b, 0x69, 0x9d, 0x8d, 0xdf, 0x93, 0x87,
	0x21, 0xb1, 0x2b, 0x42, 0x7f, 0x24, 0x61, 0xbd, 0xf1, 0x87, 0x1d, 0xc9, 0x05, 0xa1, 0xe5, 0x1e,
	0xfd, 0x17, 0x9c, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x25, 0xa5, 0xfd, 0x2c, 0x04, 0x00,
	0x00,
}
