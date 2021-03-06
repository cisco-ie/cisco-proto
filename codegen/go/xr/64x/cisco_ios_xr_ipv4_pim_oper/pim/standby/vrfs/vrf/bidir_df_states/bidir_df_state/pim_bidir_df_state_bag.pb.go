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
// source: pim_bidir_df_state_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_standby_vrfs_vrf_bidir_df_states_bidir_df_state

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

type PimBidirDfStateBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	RpAddress            string   `protobuf:"bytes,2,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimBidirDfStateBag_KEYS) Reset()         { *m = PimBidirDfStateBag_KEYS{} }
func (m *PimBidirDfStateBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimBidirDfStateBag_KEYS) ProtoMessage()    {}
func (*PimBidirDfStateBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd765c293947f0a, []int{0}
}

func (m *PimBidirDfStateBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBidirDfStateBag_KEYS.Unmarshal(m, b)
}
func (m *PimBidirDfStateBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBidirDfStateBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimBidirDfStateBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBidirDfStateBag_KEYS.Merge(m, src)
}
func (m *PimBidirDfStateBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimBidirDfStateBag_KEYS.Size(m)
}
func (m *PimBidirDfStateBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBidirDfStateBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimBidirDfStateBag_KEYS proto.InternalMessageInfo

func (m *PimBidirDfStateBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *PimBidirDfStateBag_KEYS) GetRpAddress() string {
	if m != nil {
		return m.RpAddress
	}
	return ""
}

func (m *PimBidirDfStateBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
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
	return fileDescriptor_2bd765c293947f0a, []int{1}
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

type PimBidirDfStateBag struct {
	RpAddressXr          *PimAddrtype `protobuf:"bytes,50,opt,name=rp_address_xr,json=rpAddressXr,proto3" json:"rp_address_xr,omitempty"`
	PimInterfaceName     string       `protobuf:"bytes,51,opt,name=pim_interface_name,json=pimInterfaceName,proto3" json:"pim_interface_name,omitempty"`
	ElectionState        string       `protobuf:"bytes,52,opt,name=election_state,json=electionState,proto3" json:"election_state,omitempty"`
	TimeSeconds          uint64       `protobuf:"varint,53,opt,name=time_seconds,json=timeSeconds,proto3" json:"time_seconds,omitempty"`
	TimeNanoSeconds      uint64       `protobuf:"varint,54,opt,name=time_nano_seconds,json=timeNanoSeconds,proto3" json:"time_nano_seconds,omitempty"`
	OurMetric            uint32       `protobuf:"varint,55,opt,name=our_metric,json=ourMetric,proto3" json:"our_metric,omitempty"`
	OurMetricPreference  uint32       `protobuf:"varint,56,opt,name=our_metric_preference,json=ourMetricPreference,proto3" json:"our_metric_preference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimBidirDfStateBag) Reset()         { *m = PimBidirDfStateBag{} }
func (m *PimBidirDfStateBag) String() string { return proto.CompactTextString(m) }
func (*PimBidirDfStateBag) ProtoMessage()    {}
func (*PimBidirDfStateBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd765c293947f0a, []int{2}
}

func (m *PimBidirDfStateBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBidirDfStateBag.Unmarshal(m, b)
}
func (m *PimBidirDfStateBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBidirDfStateBag.Marshal(b, m, deterministic)
}
func (m *PimBidirDfStateBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBidirDfStateBag.Merge(m, src)
}
func (m *PimBidirDfStateBag) XXX_Size() int {
	return xxx_messageInfo_PimBidirDfStateBag.Size(m)
}
func (m *PimBidirDfStateBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBidirDfStateBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimBidirDfStateBag proto.InternalMessageInfo

func (m *PimBidirDfStateBag) GetRpAddressXr() *PimAddrtype {
	if m != nil {
		return m.RpAddressXr
	}
	return nil
}

func (m *PimBidirDfStateBag) GetPimInterfaceName() string {
	if m != nil {
		return m.PimInterfaceName
	}
	return ""
}

func (m *PimBidirDfStateBag) GetElectionState() string {
	if m != nil {
		return m.ElectionState
	}
	return ""
}

func (m *PimBidirDfStateBag) GetTimeSeconds() uint64 {
	if m != nil {
		return m.TimeSeconds
	}
	return 0
}

func (m *PimBidirDfStateBag) GetTimeNanoSeconds() uint64 {
	if m != nil {
		return m.TimeNanoSeconds
	}
	return 0
}

func (m *PimBidirDfStateBag) GetOurMetric() uint32 {
	if m != nil {
		return m.OurMetric
	}
	return 0
}

func (m *PimBidirDfStateBag) GetOurMetricPreference() uint32 {
	if m != nil {
		return m.OurMetricPreference
	}
	return 0
}

func init() {
	proto.RegisterType((*PimBidirDfStateBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.bidir_df_states.bidir_df_state.pim_bidir_df_state_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.bidir_df_states.bidir_df_state.pim_addrtype")
	proto.RegisterType((*PimBidirDfStateBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.bidir_df_states.bidir_df_state.pim_bidir_df_state_bag")
}

func init() { proto.RegisterFile("pim_bidir_df_state_bag.proto", fileDescriptor_2bd765c293947f0a) }

var fileDescriptor_2bd765c293947f0a = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4b, 0x8b, 0xd4, 0x40,
	0x14, 0x85, 0x89, 0x2d, 0x33, 0xf6, 0xcd, 0xb4, 0x8f, 0x12, 0x35, 0xa2, 0x03, 0xb1, 0x41, 0x08,
	0x22, 0x59, 0xf4, 0x8c, 0xad, 0x5b, 0x17, 0x2e, 0x44, 0x6c, 0x24, 0xbd, 0x51, 0x10, 0x8a, 0x4a,
	0x72, 0x23, 0x05, 0xd6, 0x83, 0x5b, 0x35, 0x61, 0x66, 0x23, 0xfe, 0x06, 0x7f, 0xb1, 0x54, 0xc5,
	0x4e, 0x3f, 0xe8, 0xa5, 0x9b, 0x40, 0xce, 0xf9, 0x72, 0x0f, 0xe7, 0x10, 0x78, 0x6e, 0xa5, 0xe2,
	0xb5, 0x6c, 0x25, 0xf1, 0xb6, 0xe3, 0xce, 0x0b, 0x8f, 0xbc, 0x16, 0x3f, 0x4a, 0x4b, 0xc6, 0x1b,
	0xb6, 0x6a, 0xa4, 0x6b, 0x0c, 0x97, 0xc6, 0xf1, 0x6b, 0xe2, 0xd2, 0xf6, 0x97, 0x3c, 0xf0, 0xc6,
	0x22, 0x95, 0x56, 0xaa, 0xd2, 0x79, 0xa1, 0xdb, 0xfa, 0xa6, 0xec, 0xa9, 0x73, 0xe1, 0x51, 0xee,
	0x5f, 0x72, 0x07, 0xef, 0xf3, 0x5f, 0xf0, 0xec, 0x78, 0x1e, 0xff, 0xf4, 0xe1, 0xdb, 0x9a, 0x3d,
	0x85, 0x3b, 0x3d, 0x75, 0x5c, 0x0b, 0x85, 0x59, 0x92, 0x27, 0xc5, 0xb4, 0x3a, 0xed, 0xa9, 0x5b,
	0x09, 0x85, 0xec, 0x1c, 0x80, 0x2c, 0x17, 0x6d, 0x4b, 0xe8, 0x5c, 0x76, 0x2b, 0x9a, 0x53, 0xb2,
	0xef, 0x07, 0x81, 0xbd, 0x84, 0xbb, 0x52, 0x7b, 0xa4, 0x4e, 0x34, 0x38, 0x7c, 0x3f, 0x89, 0xc8,
	0x6c, 0x54, 0xc3, 0x95, 0xb9, 0x82, 0xb3, 0x90, 0x1f, 0xce, 0xf8, 0x1b, 0x8b, 0xec, 0x09, 0x9c,
	0x8a, 0xbd, 0xbc, 0x13, 0x31, 0xc4, 0xbd, 0x80, 0xb3, 0xd8, 0x76, 0x3f, 0x30, 0x0d, 0xda, 0x26,
	0x72, 0x40, 0x96, 0x23, 0x32, 0x19, 0x91, 0xe5, 0x3f, 0x64, 0xfe, 0x67, 0x02, 0x8f, 0x8f, 0xf7,
	0x65, 0xbf, 0x13, 0x98, 0x6d, 0x0b, 0xf1, 0x6b, 0xca, 0x16, 0x79, 0x52, 0xa4, 0x8b, 0xef, 0xe5,
	0xff, 0x9d, 0xbc, 0xdc, 0xed, 0x5b, 0xa5, 0xe3, 0x62, 0x5f, 0x89, 0xbd, 0x06, 0x16, 0xcc, 0x83,
	0xdd, 0x2e, 0x62, 0x8d, 0xfb, 0x56, 0xaa, 0x8f, 0xbb, 0xd3, 0x85, 0x85, 0xf1, 0x27, 0x36, 0x5e,
	0x1a, 0x3d, 0x5c, 0xce, 0x2e, 0x87, 0x85, 0x37, 0xea, 0x3a, 0x88, 0x61, 0x15, 0x2f, 0x15, 0x72,
	0x87, 0x8d, 0xd1, 0xad, 0xcb, 0xde, 0xe4, 0x49, 0x71, 0xbb, 0x4a, 0x83, 0xb6, 0x1e, 0x24, 0xf6,
	0x0a, 0x1e, 0x44, 0x44, 0x0b, 0x6d, 0x46, 0x6e, 0x19, 0xb9, 0x7b, 0xc1, 0x58, 0x09, 0x6d, 0x36,
	0xec, 0x39, 0x80, 0xb9, 0x22, 0xae, 0xd0, 0x93, 0x6c, 0xb2, 0xb7, 0x79, 0x52, 0xcc, 0xaa, 0xa9,
	0xb9, 0xa2, 0xcf, 0x51, 0x60, 0x0b, 0x78, 0xb4, 0xb5, 0xb9, 0x25, 0xec, 0x90, 0x50, 0x37, 0x98,
	0xbd, 0x8b, 0xe4, 0xc3, 0x91, 0xfc, 0x32, 0x5a, 0xf5, 0x49, 0xfc, 0xb5, 0x2f, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x0d, 0x4a, 0xcb, 0x70, 0xfa, 0x02, 0x00, 0x00,
}
