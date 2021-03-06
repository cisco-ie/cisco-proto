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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_default_context_route_policy_route_policy_tests_route_policy_test

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
	proto.RegisterType((*PimRpfPolicyTestBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.route_policy.route_policy_tests.route_policy_test.pim_rpf_policy_test_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.route_policy.route_policy_tests.route_policy_test.pim_addrtype")
	proto.RegisterType((*PimRpfPolicyTestBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.route_policy.route_policy_tests.route_policy_test.pim_rpf_policy_test_bag")
}

func init() { proto.RegisterFile("pim_rpf_policy_test_bag.proto", fileDescriptor_19a2a94ff48f1cb7) }

var fileDescriptor_19a2a94ff48f1cb7 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x15, 0x0a, 0xfb, 0x67, 0xda, 0x2e, 0x8b, 0x11, 0xda, 0x20, 0xb1, 0xa2, 0x14, 0x21,
	0xf5, 0x94, 0x43, 0x77, 0x29, 0xff, 0x0f, 0xab, 0xd5, 0x9e, 0x90, 0x38, 0x14, 0x0e, 0xcb, 0xc9,
	0x72, 0x93, 0xc9, 0xca, 0x52, 0x1d, 0x5b, 0xb6, 0x13, 0xa5, 0xe2, 0xcc, 0xe7, 0xe0, 0x84, 0xf8,
	0x9a, 0xc8, 0x93, 0xa4, 0x34, 0x42, 0x7b, 0x85, 0x9b, 0xf3, 0x9b, 0xe7, 0x99, 0xd7, 0xf1, 0x53,
	0xe1, 0xd4, 0x48, 0xc5, 0xad, 0xc9, 0xb9, 0xd1, 0x6b, 0x99, 0x6e, 0xb8, 0x47, 0xe7, 0xf9, 0x4a,
	0xdc, 0x24, 0xc6, 0x6a, 0xaf, 0x99, 0x4a, 0xa5, 0x4b, 0x35, 0x97, 0xda, 0xf1, 0xda, 0x72, 0x69,
	0xaa, 0x73, 0x1e, 0x2e, 0x68, 0x83, 0x36, 0x91, 0xa6, 0x5a, 0x84, 0xaf, 0xc4, 0x79, 0x51, 0x64,
	0xab, 0x4d, 0x92, 0x61, 0x2e, 0xca, 0xb5, 0xe7, 0xa9, 0x2e, 0x3c, 0xd6, 0x3e, 0xb1, 0xba, 0xf4,
	0xd8, 0x36, 0xee, 0x7d, 0xd0, 0x14, 0xf7, 0x37, 0x9a, 0x7e, 0x8f, 0xe0, 0xc9, 0x2d, 0x86, 0xf8,
	0xc7, 0xab, 0xaf, 0x9f, 0xd9, 0x0b, 0x38, 0x72, 0xba, 0xb4, 0x29, 0x72, 0x91, 0x65, 0x16, 0x9d,
	0x8b, 0xa3, 0x49, 0x34, 0x3b, 0x5c, 0x8e, 0x1b, 0x7a, 0xd1, 0x40, 0xf6, 0x1c, 0xc6, 0x37, 0x56,
	0x97, 0x66, 0xab, 0xba, 0x43, 0xaa, 0x11, 0xc1, 0x4e, 0xf4, 0x18, 0x0e, 0xb0, 0x0e, 0x6e, 0x95,
	0x8a, 0x07, 0x54, 0xdf, 0xc7, 0xda, 0x5f, 0x6a, 0xa5, 0xa6, 0x0a, 0x46, 0xc1, 0x46, 0xb8, 0xed,
	0x37, 0x06, 0xd9, 0x09, 0xec, 0x8b, 0x9c, 0x17, 0x42, 0x61, 0x3b, 0x6f, 0x4f, 0xe4, 0x9f, 0x84,
	0x42, 0xf6, 0x0c, 0x46, 0xb4, 0x94, 0xfe, 0x9c, 0x61, 0x60, 0xdd, 0x98, 0x46, 0xb2, 0xd8, 0x4a,
	0x06, 0x5b, 0xc9, 0xa2, 0x95, 0x4c, 0x7f, 0xdc, 0x83, 0x93, 0x5b, 0x7e, 0x36, 0x7b, 0x0a, 0xc3,
	0x16, 0xd1, 0xf8, 0x39, 0xdd, 0x86, 0x06, 0x91, 0x85, 0x5f, 0x11, 0x3c, 0xe8, 0xef, 0x84, 0xd7,
	0x36, 0x3e, 0x9b, 0x44, 0xb3, 0xe1, 0xfc, 0x5b, 0xf2, 0x4f, 0xdf, 0x2f, 0xd9, 0x5d, 0xda, 0xf2,
	0x7e, 0xef, 0x4d, 0xae, 0x2d, 0xfb, 0x19, 0xc1, 0x71, 0xef, 0x59, 0x82, 0xd1, 0xf3, 0xff, 0x6f,
	0xf4, 0x68, 0x37, 0x16, 0xd7, 0x96, 0xcd, 0xe1, 0x11, 0xd6, 0x1e, 0x8b, 0x0c, 0x33, 0x4a, 0x47,
	0x59, 0x48, 0xbf, 0xe1, 0xd6, 0xc7, 0x2f, 0x69, 0xf9, 0x0f, 0xbb, 0xe2, 0x65, 0x57, 0x5b, 0x7a,
	0xc6, 0xe0, 0xae, 0x11, 0xce, 0xc5, 0x8b, 0x49, 0x34, 0x3b, 0x58, 0xd2, 0x39, 0xa4, 0xb0, 0x73,
	0xea, 0xc5, 0x6a, 0x8d, 0xf1, 0x2b, 0x2a, 0x8e, 0x5a, 0xf8, 0x25, 0xb0, 0x90, 0xc2, 0xca, 0xb6,
	0xd9, 0x7a, 0xdd, 0xa4, 0xb0, 0xb2, 0x4d, 0xb8, 0x8e, 0x61, 0x20, 0x72, 0x19, 0xbf, 0x99, 0x44,
	0xb3, 0xf1, 0x32, 0x1c, 0xc3, 0x14, 0x17, 0xd0, 0x5b, 0x42, 0x74, 0x66, 0xa7, 0x00, 0xd4, 0xbd,
	0x69, 0xf1, 0x8e, 0x5a, 0x1c, 0x12, 0xe9, 0x12, 0xda, 0x94, 0xb1, 0x96, 0xce, 0xbb, 0xf8, 0x3d,
	0x79, 0x18, 0x12, 0xbb, 0x22, 0xf4, 0x47, 0x22, 0x52, 0x2f, 0x2b, 0x8c, 0x3f, 0xec, 0x48, 0x2e,
	0x08, 0xad, 0xf6, 0xe8, 0xef, 0xe0, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xab, 0x03,
	0x4d, 0x2f, 0x04, 0x00, 0x00,
}
