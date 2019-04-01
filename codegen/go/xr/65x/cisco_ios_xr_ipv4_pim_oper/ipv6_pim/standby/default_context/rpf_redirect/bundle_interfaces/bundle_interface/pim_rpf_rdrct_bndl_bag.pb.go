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
// source: pim_rpf_rdrct_bndl_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_default_context_rpf_redirect_bundle_interfaces_bundle_interface

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

type PimRpfRdrctBndlBag_KEYS struct {
	BundleName           string   `protobuf:"bytes,1,opt,name=bundle_name,json=bundleName,proto3" json:"bundle_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimRpfRdrctBndlBag_KEYS) Reset()         { *m = PimRpfRdrctBndlBag_KEYS{} }
func (m *PimRpfRdrctBndlBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimRpfRdrctBndlBag_KEYS) ProtoMessage()    {}
func (*PimRpfRdrctBndlBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_86ec24c5935320a1, []int{0}
}

func (m *PimRpfRdrctBndlBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfRdrctBndlBag_KEYS.Unmarshal(m, b)
}
func (m *PimRpfRdrctBndlBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfRdrctBndlBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimRpfRdrctBndlBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfRdrctBndlBag_KEYS.Merge(m, src)
}
func (m *PimRpfRdrctBndlBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimRpfRdrctBndlBag_KEYS.Size(m)
}
func (m *PimRpfRdrctBndlBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfRdrctBndlBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfRdrctBndlBag_KEYS proto.InternalMessageInfo

func (m *PimRpfRdrctBndlBag_KEYS) GetBundleName() string {
	if m != nil {
		return m.BundleName
	}
	return ""
}

func (m *PimRpfRdrctBndlBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type PimRpfRdrctBndlBag struct {
	RpfRedirectBundleName       string   `protobuf:"bytes,50,opt,name=rpf_redirect_bundle_name,json=rpfRedirectBundleName,proto3" json:"rpf_redirect_bundle_name,omitempty"`
	RpfRedirectInterfaceName    string   `protobuf:"bytes,51,opt,name=rpf_redirect_interface_name,json=rpfRedirectInterfaceName,proto3" json:"rpf_redirect_interface_name,omitempty"`
	AvailableBandwidth          int32    `protobuf:"zigzag32,52,opt,name=available_bandwidth,json=availableBandwidth,proto3" json:"available_bandwidth,omitempty"`
	AllocatedBandwidth          int32    `protobuf:"zigzag32,53,opt,name=allocated_bandwidth,json=allocatedBandwidth,proto3" json:"allocated_bandwidth,omitempty"`
	TotalBandwidth              int32    `protobuf:"zigzag32,54,opt,name=total_bandwidth,json=totalBandwidth,proto3" json:"total_bandwidth,omitempty"`
	TopologyBandwidthUsed       int32    `protobuf:"zigzag32,55,opt,name=topology_bandwidth_used,json=topologyBandwidthUsed,proto3" json:"topology_bandwidth_used,omitempty"`
	SnoopingBandwidthUsed       int32    `protobuf:"zigzag32,56,opt,name=snooping_bandwidth_used,json=snoopingBandwidthUsed,proto3" json:"snooping_bandwidth_used,omitempty"`
	AllocatedThresholdBandwidth int32    `protobuf:"zigzag32,57,opt,name=allocated_threshold_bandwidth,json=allocatedThresholdBandwidth,proto3" json:"allocated_threshold_bandwidth,omitempty"`
	AvailableThresholdBandwidth int32    `protobuf:"zigzag32,58,opt,name=available_threshold_bandwidth,json=availableThresholdBandwidth,proto3" json:"available_threshold_bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *PimRpfRdrctBndlBag) Reset()         { *m = PimRpfRdrctBndlBag{} }
func (m *PimRpfRdrctBndlBag) String() string { return proto.CompactTextString(m) }
func (*PimRpfRdrctBndlBag) ProtoMessage()    {}
func (*PimRpfRdrctBndlBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_86ec24c5935320a1, []int{1}
}

func (m *PimRpfRdrctBndlBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfRdrctBndlBag.Unmarshal(m, b)
}
func (m *PimRpfRdrctBndlBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfRdrctBndlBag.Marshal(b, m, deterministic)
}
func (m *PimRpfRdrctBndlBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfRdrctBndlBag.Merge(m, src)
}
func (m *PimRpfRdrctBndlBag) XXX_Size() int {
	return xxx_messageInfo_PimRpfRdrctBndlBag.Size(m)
}
func (m *PimRpfRdrctBndlBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfRdrctBndlBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfRdrctBndlBag proto.InternalMessageInfo

func (m *PimRpfRdrctBndlBag) GetRpfRedirectBundleName() string {
	if m != nil {
		return m.RpfRedirectBundleName
	}
	return ""
}

func (m *PimRpfRdrctBndlBag) GetRpfRedirectInterfaceName() string {
	if m != nil {
		return m.RpfRedirectInterfaceName
	}
	return ""
}

func (m *PimRpfRdrctBndlBag) GetAvailableBandwidth() int32 {
	if m != nil {
		return m.AvailableBandwidth
	}
	return 0
}

func (m *PimRpfRdrctBndlBag) GetAllocatedBandwidth() int32 {
	if m != nil {
		return m.AllocatedBandwidth
	}
	return 0
}

func (m *PimRpfRdrctBndlBag) GetTotalBandwidth() int32 {
	if m != nil {
		return m.TotalBandwidth
	}
	return 0
}

func (m *PimRpfRdrctBndlBag) GetTopologyBandwidthUsed() int32 {
	if m != nil {
		return m.TopologyBandwidthUsed
	}
	return 0
}

func (m *PimRpfRdrctBndlBag) GetSnoopingBandwidthUsed() int32 {
	if m != nil {
		return m.SnoopingBandwidthUsed
	}
	return 0
}

func (m *PimRpfRdrctBndlBag) GetAllocatedThresholdBandwidth() int32 {
	if m != nil {
		return m.AllocatedThresholdBandwidth
	}
	return 0
}

func (m *PimRpfRdrctBndlBag) GetAvailableThresholdBandwidth() int32 {
	if m != nil {
		return m.AvailableThresholdBandwidth
	}
	return 0
}

func init() {
	proto.RegisterType((*PimRpfRdrctBndlBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.rpf_redirect.bundle_interfaces.bundle_interface.pim_rpf_rdrct_bndl_bag_KEYS")
	proto.RegisterType((*PimRpfRdrctBndlBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.rpf_redirect.bundle_interfaces.bundle_interface.pim_rpf_rdrct_bndl_bag")
}

func init() { proto.RegisterFile("pim_rpf_rdrct_bndl_bag.proto", fileDescriptor_86ec24c5935320a1) }

var fileDescriptor_86ec24c5935320a1 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0xe9, 0xc7, 0x87, 0x60, 0xc4, 0x8a, 0x23, 0xd5, 0x81, 0x2a, 0x96, 0x82, 0xd8, 0xd5,
	0x08, 0xb6, 0xb6, 0x2a, 0xb8, 0x29, 0xb8, 0x10, 0xc1, 0x45, 0xd5, 0x85, 0xab, 0x90, 0x99, 0x9c,
	0x69, 0x83, 0x69, 0x12, 0x92, 0xb4, 0xb6, 0xd7, 0xe1, 0x0d, 0xcb, 0xa4, 0xf3, 0x93, 0xfe, 0x2c,
	0x7b, 0xde, 0xe7, 0x79, 0xcf, 0x49, 0x19, 0x74, 0xae, 0xd8, 0x14, 0x6b, 0x95, 0x62, 0x4d, 0x75,
	0x62, 0x71, 0x2c, 0x28, 0xc7, 0x31, 0x19, 0x47, 0x4a, 0x4b, 0x2b, 0x83, 0xef, 0x84, 0x99, 0x44,
	0x62, 0x26, 0x0d, 0x5e, 0x68, 0xcc, 0xd4, 0xbc, 0x87, 0x33, 0x5e, 0x2a, 0xd0, 0x11, 0x53, 0xf3,
	0x7e, 0xf6, 0x2b, 0x32, 0x96, 0x08, 0x1a, 0x2f, 0x23, 0x0a, 0x29, 0x99, 0x71, 0x8b, 0x13, 0x29,
	0x2c, 0x2c, 0x6c, 0xe4, 0x5a, 0x81, 0x32, 0x0d, 0x89, 0x8d, 0xe2, 0x99, 0xa0, 0x1c, 0x30, 0x13,
	0x16, 0x74, 0x4a, 0x12, 0x30, 0x5b, 0x93, 0x36, 0xa0, 0xe6, 0xee, 0x63, 0xf0, 0xeb, 0xf3, 0xd7,
	0x7b, 0x70, 0x89, 0x0e, 0x72, 0x45, 0x90, 0x29, 0x84, 0xb5, 0x56, 0xad, 0xb3, 0x3f, 0x42, 0xab,
	0xd1, 0x1b, 0x99, 0x42, 0x70, 0x85, 0xea, 0x65, 0xd9, 0x8a, 0xf9, 0xe7, 0x98, 0xc3, 0x72, 0x9a,
	0x61, 0xed, 0xdf, 0xff, 0xe8, 0x74, 0xf7, 0x9e, 0x60, 0x80, 0x42, 0xff, 0x68, 0xec, 0xef, 0xbb,
	0x75, 0x5d, 0x0d, 0xad, 0xd2, 0x51, 0x1e, 0x0f, 0xab, 0xd5, 0x4f, 0xa8, 0xb9, 0x26, 0x6e, 0xdc,
	0xd1, 0x75, 0x6e, 0xe8, 0xb9, 0x2f, 0xfe, 0x49, 0xc1, 0x0d, 0x3a, 0x21, 0x73, 0xc2, 0x38, 0x89,
	0x39, 0xe0, 0x98, 0x08, 0xfa, 0xc3, 0xa8, 0x9d, 0x84, 0xbd, 0x56, 0xad, 0x73, 0x3c, 0x0a, 0xca,
	0x68, 0x58, 0x24, 0x4e, 0xe0, 0x5c, 0x26, 0xc4, 0x02, 0xf5, 0x84, 0xbb, 0x5c, 0x28, 0xa2, 0x4a,
	0xb8, 0x46, 0x47, 0x56, 0x5a, 0xc2, 0x3d, 0xb8, 0xef, 0xe0, 0xba, 0x1b, 0x57, 0x60, 0x1f, 0x9d,
	0x59, 0xa9, 0x24, 0x97, 0xe3, 0x65, 0xc5, 0xe2, 0x99, 0x01, 0x1a, 0x0e, 0x9c, 0xd0, 0x28, 0xe2,
	0xd2, 0xf9, 0x34, 0x40, 0x33, 0xcf, 0x08, 0x29, 0x15, 0x13, 0xe3, 0x4d, 0xef, 0x7e, 0xe5, 0x15,
	0xf1, 0xba, 0x37, 0x44, 0x17, 0xd5, 0x4b, 0xec, 0x44, 0x83, 0x99, 0x48, 0xee, 0xbf, 0xe9, 0xc1,
	0xd9, 0xcd, 0x12, 0xfa, 0x28, 0x98, 0xea, 0xe6, 0xac, 0xa3, 0xfc, 0xfb, 0x76, 0x75, 0x3c, 0xe6,
	0x1d, 0x05, 0xb4, 0xdd, 0x11, 0xef, 0xb9, 0x0f, 0xbe, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x32,
	0x2c, 0x38, 0x32, 0x10, 0x03, 0x00, 0x00,
}
