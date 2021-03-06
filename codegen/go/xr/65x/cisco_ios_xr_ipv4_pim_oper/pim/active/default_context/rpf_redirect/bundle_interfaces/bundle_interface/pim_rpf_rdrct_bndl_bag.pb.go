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

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_rpf_redirect_bundle_interfaces_bundle_interface

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
	proto.RegisterType((*PimRpfRdrctBndlBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.rpf_redirect.bundle_interfaces.bundle_interface.pim_rpf_rdrct_bndl_bag_KEYS")
	proto.RegisterType((*PimRpfRdrctBndlBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.rpf_redirect.bundle_interfaces.bundle_interface.pim_rpf_rdrct_bndl_bag")
}

func init() { proto.RegisterFile("pim_rpf_rdrct_bndl_bag.proto", fileDescriptor_86ec24c5935320a1) }

var fileDescriptor_86ec24c5935320a1 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0xe9, 0xc7, 0x87, 0x60, 0xc4, 0x8a, 0x23, 0xd5, 0x81, 0x2a, 0x96, 0x82, 0xd8, 0xd5,
	0x08, 0xb6, 0xb6, 0x2a, 0xb8, 0x29, 0xb8, 0x10, 0xc1, 0x45, 0xd5, 0x85, 0xab, 0x90, 0x49, 0xce,
	0xb4, 0x81, 0x4c, 0x12, 0x32, 0x69, 0xad, 0xd7, 0xe1, 0x0d, 0xcb, 0xa4, 0xf3, 0x93, 0xfe, 0x6c,
	0xcf, 0xfb, 0x3c, 0xef, 0x39, 0x19, 0x06, 0x9d, 0x6b, 0x9e, 0x62, 0xa3, 0x13, 0x6c, 0x98, 0xa1,
	0x16, 0xc7, 0x92, 0x09, 0x1c, 0x93, 0x69, 0xa4, 0x8d, 0xb2, 0x2a, 0x00, 0xca, 0x33, 0xaa, 0x30,
	0x57, 0x19, 0x5e, 0x1a, 0xcc, 0xf5, 0x62, 0x80, 0x73, 0x5e, 0x69, 0x30, 0x91, 0xe6, 0x69, 0x44,
	0xa8, 0xe5, 0x0b, 0x88, 0x18, 0x24, 0x64, 0x2e, 0x2c, 0xa6, 0x4a, 0x5a, 0x58, 0xda, 0xc8, 0xf5,
	0x01, 0xe3, 0x06, 0xa8, 0x8d, 0xe2, 0xb9, 0x64, 0x02, 0x30, 0x97, 0x16, 0x4c, 0x42, 0x28, 0x64,
	0x5b, 0x93, 0x2e, 0xa0, 0xf6, 0xee, 0x33, 0xf0, 0xeb, 0xf3, 0xd7, 0x7b, 0x70, 0x89, 0x0e, 0x0a,
	0x45, 0x92, 0x14, 0xc2, 0x46, 0xa7, 0xd1, 0xdb, 0x9f, 0xa0, 0xd5, 0xe8, 0x8d, 0xa4, 0x10, 0x5c,
	0xa1, 0x66, 0x55, 0xb6, 0x62, 0xfe, 0x39, 0xe6, 0xb0, 0x9a, 0xe6, 0x58, 0xf7, 0xf7, 0x3f, 0x3a,
	0xdd, 0xbd, 0x27, 0x18, 0xa1, 0xd0, 0x3f, 0x1a, 0xfb, 0xfb, 0x6e, 0x5d, 0x57, 0xcb, 0xe8, 0x64,
	0x52, 0xc4, 0xe3, 0x7a, 0xf5, 0x13, 0x6a, 0xaf, 0x89, 0x1b, 0x77, 0xf4, 0x9d, 0x1b, 0x7a, 0xee,
	0x8b, 0x7f, 0x52, 0x70, 0x83, 0x4e, 0xc8, 0x82, 0x70, 0x41, 0x62, 0x01, 0x38, 0x26, 0x92, 0x7d,
	0x73, 0x66, 0x67, 0xe1, 0xa0, 0xd3, 0xe8, 0x1d, 0x4f, 0x82, 0x2a, 0x1a, 0x97, 0x89, 0x13, 0x84,
	0x50, 0x94, 0x58, 0x60, 0x9e, 0x70, 0x57, 0x08, 0x65, 0x54, 0x0b, 0xd7, 0xe8, 0xc8, 0x2a, 0x4b,
	0x84, 0x07, 0x0f, 0x1d, 0xdc, 0x74, 0xe3, 0x1a, 0x1c, 0xa2, 0x33, 0xab, 0xb4, 0x12, 0x6a, 0xfa,
	0x53, 0xb3, 0x78, 0x9e, 0x01, 0x0b, 0x47, 0x4e, 0x68, 0x95, 0x71, 0xe5, 0x7c, 0x66, 0xc0, 0x72,
	0x2f, 0x93, 0x4a, 0x69, 0x2e, 0xa7, 0x9b, 0xde, 0xfd, 0xca, 0x2b, 0xe3, 0x75, 0x6f, 0x8c, 0x2e,
	0xea, 0x97, 0xd8, 0x99, 0x81, 0x6c, 0xa6, 0x84, 0xff, 0xa6, 0x07, 0x67, 0xb7, 0x2b, 0xe8, 0xa3,
	0x64, 0xea, 0x9b, 0xf3, 0x8e, 0xea, 0xf3, 0xed, 0xea, 0x78, 0x2c, 0x3a, 0x4a, 0x68, 0xbb, 0x23,
	0xde, 0x73, 0xbf, 0x7a, 0xff, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x38, 0x13, 0xac, 0x0a, 0x03,
	0x00, 0x00,
}
