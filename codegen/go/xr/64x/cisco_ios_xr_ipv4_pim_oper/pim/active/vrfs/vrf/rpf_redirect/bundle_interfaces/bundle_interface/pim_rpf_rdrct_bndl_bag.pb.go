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

package cisco_ios_xr_ipv4_pim_oper_pim_active_vrfs_vrf_rpf_redirect_bundle_interfaces_bundle_interface

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	BundleName           string   `protobuf:"bytes,2,opt,name=bundle_name,json=bundleName,proto3" json:"bundle_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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

func (m *PimRpfRdrctBndlBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*PimRpfRdrctBndlBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.rpf_redirect.bundle_interfaces.bundle_interface.pim_rpf_rdrct_bndl_bag_KEYS")
	proto.RegisterType((*PimRpfRdrctBndlBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.rpf_redirect.bundle_interfaces.bundle_interface.pim_rpf_rdrct_bndl_bag")
}

func init() { proto.RegisterFile("pim_rpf_rdrct_bndl_bag.proto", fileDescriptor_86ec24c5935320a1) }

var fileDescriptor_86ec24c5935320a1 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x86, 0xc9, 0xbd, 0x97, 0xab, 0x8e, 0x58, 0x31, 0x52, 0x8d, 0x54, 0xb1, 0x14, 0xc4, 0xae,
	0x22, 0xd8, 0xda, 0xaa, 0xe0, 0xa6, 0xe0, 0x42, 0x04, 0x17, 0x55, 0x17, 0x6e, 0x1c, 0x26, 0x99,
	0x93, 0x76, 0x60, 0x92, 0x19, 0x66, 0xa6, 0x51, 0x77, 0xfe, 0x07, 0xff, 0xb0, 0x64, 0x9a, 0xaf,
	0x7e, 0x6c, 0xba, 0x38, 0xef, 0xf3, 0xbc, 0x73, 0x4e, 0x09, 0x3a, 0x96, 0x2c, 0xc6, 0x4a, 0x46,
	0x58, 0x51, 0x15, 0x1a, 0x1c, 0x24, 0x94, 0xe3, 0x80, 0x4c, 0x7c, 0xa9, 0x84, 0x11, 0xee, 0x7b,
	0xc8, 0x74, 0x28, 0x30, 0x13, 0x1a, 0x7f, 0x2a, 0xcc, 0x64, 0xda, 0xc7, 0x19, 0x2f, 0x24, 0x28,
	0x5f, 0xb2, 0xd8, 0x27, 0xa1, 0x61, 0x29, 0xf8, 0xa9, 0x8a, 0x74, 0xf6, 0xe3, 0xdb, 0x22, 0xa0,
	0x4c, 0x41, 0x68, 0xfc, 0x60, 0x96, 0x50, 0x0e, 0x98, 0x25, 0x06, 0x54, 0x44, 0x42, 0xd0, 0x2b,
	0x93, 0xce, 0xb7, 0x83, 0x5a, 0xeb, 0x17, 0xc0, 0x8f, 0xf7, 0x6f, 0xcf, 0xee, 0x11, 0xda, 0x4c,
	0x55, 0x84, 0x13, 0x12, 0x83, 0xe7, 0xb4, 0x9d, 0xee, 0xd6, 0x78, 0x23, 0x55, 0xd1, 0x13, 0x89,
	0xc1, 0x3d, 0x45, 0xdb, 0x79, 0x9d, 0x4d, 0xff, 0xd8, 0x14, 0xcd, 0x47, 0x16, 0x38, 0x43, 0x8d,
	0xf2, 0xa1, 0x39, 0xf3, 0xd7, 0x32, 0x3b, 0xe5, 0x34, 0xc3, 0x3a, 0x3f, 0xff, 0xd0, 0xc1, 0xfa,
	0x15, 0xdc, 0x21, 0xf2, 0xea, 0x07, 0xe1, 0xfa, 0x7b, 0x97, 0xb6, 0xab, 0xa9, 0x64, 0x34, 0xce,
	0xe3, 0x51, 0xf5, 0xf4, 0x1d, 0x6a, 0x2d, 0x88, 0x4b, 0x7b, 0xf4, 0xac, 0xeb, 0xd5, 0xdc, 0x87,
	0xfa, 0x4a, 0xee, 0x05, 0xda, 0x27, 0x29, 0x61, 0x9c, 0x04, 0x1c, 0x70, 0x40, 0x12, 0xfa, 0xc1,
	0xa8, 0x99, 0x7a, 0xfd, 0xb6, 0xd3, 0xdd, 0x1b, 0xbb, 0x65, 0x34, 0x2a, 0x12, 0x2b, 0x70, 0x2e,
	0x42, 0x62, 0x80, 0xd6, 0x84, 0xab, 0x5c, 0x28, 0xa2, 0x4a, 0x38, 0x47, 0xbb, 0x46, 0x18, 0xc2,
	0x6b, 0xf0, 0xc0, 0xc2, 0x0d, 0x3b, 0xae, 0xc0, 0x01, 0x3a, 0x34, 0x42, 0x0a, 0x2e, 0x26, 0x5f,
	0x15, 0x8b, 0x67, 0x1a, 0xa8, 0x37, 0xb4, 0x42, 0xb3, 0x88, 0x4b, 0xe7, 0x55, 0x03, 0xcd, 0x3c,
	0x9d, 0x08, 0x21, 0x59, 0x32, 0x59, 0xf6, 0xae, 0xe7, 0x5e, 0x11, 0x2f, 0x7a, 0x23, 0x74, 0x52,
	0x5d, 0x62, 0xa6, 0x0a, 0xf4, 0x54, 0xf0, 0xfa, 0x4d, 0x37, 0xd6, 0x6e, 0x95, 0xd0, 0x4b, 0xc1,
	0x54, 0x3b, 0x67, 0x1d, 0xe5, 0xdf, 0xb7, 0xae, 0xe3, 0x36, 0xef, 0x28, 0xa0, 0xd5, 0x8e, 0xe0,
	0xbf, 0xfd, 0xfe, 0x7b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0xfb, 0x11, 0xda, 0x1f, 0x03,
	0x00, 0x00,
}