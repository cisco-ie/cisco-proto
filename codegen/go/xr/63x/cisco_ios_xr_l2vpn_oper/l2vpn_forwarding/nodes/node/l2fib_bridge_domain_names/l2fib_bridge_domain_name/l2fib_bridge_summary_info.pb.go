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
// source: l2fib_bridge_summary_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_bridge_domain_names_l2fib_bridge_domain_name

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

type L2FibBridgeSummaryInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	BdGroupName          string   `protobuf:"bytes,2,opt,name=bd_group_name,json=bdGroupName,proto3" json:"bd_group_name,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibBridgeSummaryInfo_KEYS) Reset()         { *m = L2FibBridgeSummaryInfo_KEYS{} }
func (m *L2FibBridgeSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibBridgeSummaryInfo_KEYS) ProtoMessage()    {}
func (*L2FibBridgeSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f9edfcbf466325, []int{0}
}

func (m *L2FibBridgeSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibBridgeSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibBridgeSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibBridgeSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibBridgeSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibBridgeSummaryInfo_KEYS.Merge(m, src)
}
func (m *L2FibBridgeSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibBridgeSummaryInfo_KEYS.Size(m)
}
func (m *L2FibBridgeSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibBridgeSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibBridgeSummaryInfo_KEYS proto.InternalMessageInfo

func (m *L2FibBridgeSummaryInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibBridgeSummaryInfo_KEYS) GetBdGroupName() string {
	if m != nil {
		return m.BdGroupName
	}
	return ""
}

func (m *L2FibBridgeSummaryInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type L2FibBridgeSummaryInfo struct {
	BridgeId                 uint32   `protobuf:"varint,50,opt,name=bridge_id,json=bridgeId,proto3" json:"bridge_id,omitempty"`
	BridgeName               string   `protobuf:"bytes,51,opt,name=bridge_name,json=bridgeName,proto3" json:"bridge_name,omitempty"`
	MacLimit                 uint32   `protobuf:"varint,52,opt,name=mac_limit,json=macLimit,proto3" json:"mac_limit,omitempty"`
	MacLimitAction           string   `protobuf:"bytes,53,opt,name=mac_limit_action,json=macLimitAction,proto3" json:"mac_limit_action,omitempty"`
	FloodDisabled            bool     `protobuf:"varint,54,opt,name=flood_disabled,json=floodDisabled,proto3" json:"flood_disabled,omitempty"`
	MacLearningDisabled      bool     `protobuf:"varint,55,opt,name=mac_learning_disabled,json=macLearningDisabled,proto3" json:"mac_learning_disabled,omitempty"`
	MacPortDownFlushDisabled bool     `protobuf:"varint,56,opt,name=mac_port_down_flush_disabled,json=macPortDownFlushDisabled,proto3" json:"mac_port_down_flush_disabled,omitempty"`
	AdminDisabled            bool     `protobuf:"varint,57,opt,name=admin_disabled,json=adminDisabled,proto3" json:"admin_disabled,omitempty"`
	BridgePortCount          uint32   `protobuf:"varint,58,opt,name=bridge_port_count,json=bridgePortCount,proto3" json:"bridge_port_count,omitempty"`
	NumberOfShg              uint32   `protobuf:"varint,59,opt,name=number_of_shg,json=numberOfShg,proto3" json:"number_of_shg,omitempty"`
	NumberOfHwmac            uint32   `protobuf:"varint,60,opt,name=number_of_hwmac,json=numberOfHwmac,proto3" json:"number_of_hwmac,omitempty"`
	NumberOfSwmac            uint32   `protobuf:"varint,61,opt,name=number_of_swmac,json=numberOfSwmac,proto3" json:"number_of_swmac,omitempty"`
	Mtu                      uint32   `protobuf:"varint,62,opt,name=mtu,proto3" json:"mtu,omitempty"`
	AgingTimeOut             uint32   `protobuf:"varint,63,opt,name=aging_time_out,json=agingTimeOut,proto3" json:"aging_time_out,omitempty"`
	Msti                     uint32   `protobuf:"varint,64,opt,name=msti,proto3" json:"msti,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *L2FibBridgeSummaryInfo) Reset()         { *m = L2FibBridgeSummaryInfo{} }
func (m *L2FibBridgeSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibBridgeSummaryInfo) ProtoMessage()    {}
func (*L2FibBridgeSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f9edfcbf466325, []int{1}
}

func (m *L2FibBridgeSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibBridgeSummaryInfo.Unmarshal(m, b)
}
func (m *L2FibBridgeSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibBridgeSummaryInfo.Marshal(b, m, deterministic)
}
func (m *L2FibBridgeSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibBridgeSummaryInfo.Merge(m, src)
}
func (m *L2FibBridgeSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibBridgeSummaryInfo.Size(m)
}
func (m *L2FibBridgeSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibBridgeSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibBridgeSummaryInfo proto.InternalMessageInfo

func (m *L2FibBridgeSummaryInfo) GetBridgeId() uint32 {
	if m != nil {
		return m.BridgeId
	}
	return 0
}

func (m *L2FibBridgeSummaryInfo) GetBridgeName() string {
	if m != nil {
		return m.BridgeName
	}
	return ""
}

func (m *L2FibBridgeSummaryInfo) GetMacLimit() uint32 {
	if m != nil {
		return m.MacLimit
	}
	return 0
}

func (m *L2FibBridgeSummaryInfo) GetMacLimitAction() string {
	if m != nil {
		return m.MacLimitAction
	}
	return ""
}

func (m *L2FibBridgeSummaryInfo) GetFloodDisabled() bool {
	if m != nil {
		return m.FloodDisabled
	}
	return false
}

func (m *L2FibBridgeSummaryInfo) GetMacLearningDisabled() bool {
	if m != nil {
		return m.MacLearningDisabled
	}
	return false
}

func (m *L2FibBridgeSummaryInfo) GetMacPortDownFlushDisabled() bool {
	if m != nil {
		return m.MacPortDownFlushDisabled
	}
	return false
}

func (m *L2FibBridgeSummaryInfo) GetAdminDisabled() bool {
	if m != nil {
		return m.AdminDisabled
	}
	return false
}

func (m *L2FibBridgeSummaryInfo) GetBridgePortCount() uint32 {
	if m != nil {
		return m.BridgePortCount
	}
	return 0
}

func (m *L2FibBridgeSummaryInfo) GetNumberOfShg() uint32 {
	if m != nil {
		return m.NumberOfShg
	}
	return 0
}

func (m *L2FibBridgeSummaryInfo) GetNumberOfHwmac() uint32 {
	if m != nil {
		return m.NumberOfHwmac
	}
	return 0
}

func (m *L2FibBridgeSummaryInfo) GetNumberOfSwmac() uint32 {
	if m != nil {
		return m.NumberOfSwmac
	}
	return 0
}

func (m *L2FibBridgeSummaryInfo) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *L2FibBridgeSummaryInfo) GetAgingTimeOut() uint32 {
	if m != nil {
		return m.AgingTimeOut
	}
	return 0
}

func (m *L2FibBridgeSummaryInfo) GetMsti() uint32 {
	if m != nil {
		return m.Msti
	}
	return 0
}

func init() {
	proto.RegisterType((*L2FibBridgeSummaryInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_bridge_domain_names.l2fib_bridge_domain_name.l2fib_bridge_summary_info_KEYS")
	proto.RegisterType((*L2FibBridgeSummaryInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_bridge_domain_names.l2fib_bridge_domain_name.l2fib_bridge_summary_info")
}

func init() { proto.RegisterFile("l2fib_bridge_summary_info.proto", fileDescriptor_23f9edfcbf466325) }

var fileDescriptor_23f9edfcbf466325 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4b, 0x6f, 0x13, 0x3f,
	0x14, 0xc5, 0x95, 0x7f, 0xab, 0xfe, 0x1b, 0x87, 0xa4, 0xc5, 0x08, 0x61, 0x54, 0x44, 0xa3, 0x88,
	0xa2, 0x88, 0x45, 0x16, 0x29, 0xef, 0x47, 0x01, 0x51, 0x1e, 0x15, 0x88, 0xa2, 0x84, 0x0d, 0xab,
	0x2b, 0xcf, 0xd8, 0x33, 0xb1, 0x34, 0xf6, 0x0d, 0xb6, 0x87, 0xc0, 0x97, 0xe4, 0x33, 0x21, 0xdf,
	0x49, 0x13, 0x65, 0x91, 0xcd, 0xe8, 0xfa, 0x9c, 0xdf, 0x39, 0xb6, 0x3c, 0x33, 0xec, 0xb8, 0x1a,
	0x17, 0x26, 0x83, 0xcc, 0x1b, 0x55, 0x6a, 0x08, 0xb5, 0xb5, 0xd2, 0xff, 0x01, 0xe3, 0x0a, 0x1c,
	0xcd, 0x3d, 0x46, 0xe4, 0x45, 0x6e, 0x42, 0x8e, 0x60, 0x30, 0xc0, 0x6f, 0x0f, 0xd5, 0xf8, 0xd7,
	0xdc, 0x01, 0xce, 0xb5, 0x1f, 0x35, 0x63, 0x81, 0x7e, 0x21, 0xbd, 0x32, 0xae, 0x1c, 0x39, 0x54,
	0x3a, 0xd0, 0x73, 0xb4, 0x51, 0xaa, 0xd0, 0x4a, 0xe3, 0xc0, 0x49, 0xab, 0xc3, 0x56, 0x67, 0xf0,
	0x93, 0xdd, 0xdd, 0x7a, 0x14, 0xf8, 0xfc, 0xfe, 0xc7, 0x94, 0xdf, 0x62, 0xff, 0xa7, 0x72, 0x30,
	0x4a, 0xb4, 0xfa, 0xad, 0x61, 0x7b, 0xb2, 0x97, 0x96, 0x17, 0x8a, 0x0f, 0x58, 0x37, 0x53, 0x50,
	0x7a, 0xac, 0xe7, 0xd4, 0x25, 0xfe, 0x23, 0xbb, 0x93, 0xa9, 0x8f, 0x49, 0xfb, 0x2a, 0xad, 0xe6,
	0x9c, 0xed, 0x92, 0xb5, 0x43, 0x16, 0xcd, 0x83, 0xbf, 0xbb, 0xec, 0xf6, 0xd6, 0x3d, 0xf9, 0x11,
	0x6b, 0x2f, 0x65, 0xa3, 0xc4, 0xb8, 0xdf, 0x1a, 0x76, 0x27, 0xfb, 0x8d, 0x70, 0xa1, 0xf8, 0x31,
	0xeb, 0x2c, 0x4d, 0x6a, 0x3d, 0xa5, 0x56, 0xd6, 0x48, 0xb4, 0xdf, 0x11, 0x6b, 0x5b, 0x99, 0x43,
	0x65, 0xac, 0x89, 0xe2, 0x61, 0x93, 0xb6, 0x32, 0xff, 0x92, 0xd6, 0x7c, 0xc8, 0x0e, 0x57, 0x26,
	0xc8, 0x3c, 0x1a, 0x74, 0xe2, 0x11, 0x55, 0xf4, 0xae, 0x98, 0xb7, 0xa4, 0xf2, 0x13, 0xd6, 0x2b,
	0x2a, 0x44, 0x05, 0xca, 0x04, 0x99, 0x55, 0x5a, 0x89, 0xc7, 0xfd, 0xd6, 0x70, 0x7f, 0xd2, 0x25,
	0xf5, 0x7c, 0x29, 0xf2, 0x31, 0xbb, 0x49, 0x85, 0x5a, 0x7a, 0x67, 0x5c, 0xb9, 0xa6, 0x9f, 0x10,
	0x7d, 0x23, 0xb5, 0x2e, 0xbd, 0x55, 0xe6, 0x8c, 0xdd, 0x49, 0x99, 0x39, 0xfa, 0x08, 0x0a, 0x17,
	0x0e, 0x8a, 0xaa, 0x0e, 0xb3, 0x75, 0xf4, 0x29, 0x45, 0x85, 0x95, 0xf9, 0x37, 0xf4, 0xf1, 0x1c,
	0x17, 0xee, 0x43, 0x02, 0x56, 0xf9, 0x13, 0xd6, 0x93, 0xca, 0x1a, 0xb7, 0x4e, 0x3c, 0x6b, 0x8e,
	0x46, 0xea, 0x0a, 0x7b, 0xc0, 0xae, 0x2f, 0x6f, 0x8a, 0x76, 0xca, 0xb1, 0x76, 0x51, 0x3c, 0xa7,
	0x0b, 0x39, 0x68, 0x8c, 0x54, 0xff, 0x2e, 0xc9, 0xe9, 0x45, 0xba, 0xda, 0x66, 0xda, 0x03, 0x16,
	0x10, 0x66, 0xa5, 0x78, 0x41, 0x5c, 0xa7, 0x11, 0x2f, 0x8b, 0xe9, 0xac, 0xe4, 0xf7, 0xd9, 0xc1,
	0x9a, 0x99, 0x2d, 0xac, 0xcc, 0xc5, 0x4b, 0xa2, 0xba, 0x57, 0xd4, 0xa7, 0x24, 0x6e, 0x72, 0x81,
	0xb8, 0x57, 0x9b, 0xdc, 0x94, 0xb8, 0x43, 0xb6, 0x63, 0x63, 0x2d, 0xce, 0xc8, 0x4b, 0x23, 0xbf,
	0xc7, 0x7a, 0xb2, 0x4c, 0xb7, 0x18, 0x8d, 0xd5, 0x80, 0x75, 0x14, 0xaf, 0xc9, 0xbc, 0x46, 0xea,
	0x77, 0x63, 0xf5, 0x65, 0x1d, 0xd3, 0x07, 0x65, 0x43, 0x34, 0xe2, 0x0d, 0x79, 0x34, 0x67, 0x7b,
	0xf4, 0xcb, 0x9c, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x39, 0x84, 0xb1, 0x55, 0x03, 0x00,
	0x00,
}
