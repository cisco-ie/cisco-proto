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
// source: pppoe_ma_limit_info.proto

package cisco_ios_xr_subscriber_pppoe_ma_oper_pppoe_nodes_node_bba_groups_bba_group_limits_limit

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

type PppoeMaLimitInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	BbaGroupName         string   `protobuf:"bytes,2,opt,name=bba_group_name,json=bbaGroupName,proto3" json:"bba_group_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	MacAddress           string   `protobuf:"bytes,4,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Iwf                  bool     `protobuf:"varint,5,opt,name=iwf,proto3" json:"iwf,omitempty"`
	CircuitId            string   `protobuf:"bytes,6,opt,name=circuit_id,json=circuitId,proto3" json:"circuit_id,omitempty"`
	RemoteId             string   `protobuf:"bytes,7,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
	OuterVlanId          uint32   `protobuf:"varint,8,opt,name=outer_vlan_id,json=outerVlanId,proto3" json:"outer_vlan_id,omitempty"`
	InnerVlanId          uint32   `protobuf:"varint,9,opt,name=inner_vlan_id,json=innerVlanId,proto3" json:"inner_vlan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PppoeMaLimitInfo_KEYS) Reset()         { *m = PppoeMaLimitInfo_KEYS{} }
func (m *PppoeMaLimitInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PppoeMaLimitInfo_KEYS) ProtoMessage()    {}
func (*PppoeMaLimitInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef55ad5832e06799, []int{0}
}

func (m *PppoeMaLimitInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PppoeMaLimitInfo_KEYS.Unmarshal(m, b)
}
func (m *PppoeMaLimitInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PppoeMaLimitInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PppoeMaLimitInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PppoeMaLimitInfo_KEYS.Merge(m, src)
}
func (m *PppoeMaLimitInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PppoeMaLimitInfo_KEYS.Size(m)
}
func (m *PppoeMaLimitInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PppoeMaLimitInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PppoeMaLimitInfo_KEYS proto.InternalMessageInfo

func (m *PppoeMaLimitInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PppoeMaLimitInfo_KEYS) GetBbaGroupName() string {
	if m != nil {
		return m.BbaGroupName
	}
	return ""
}

func (m *PppoeMaLimitInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *PppoeMaLimitInfo_KEYS) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *PppoeMaLimitInfo_KEYS) GetIwf() bool {
	if m != nil {
		return m.Iwf
	}
	return false
}

func (m *PppoeMaLimitInfo_KEYS) GetCircuitId() string {
	if m != nil {
		return m.CircuitId
	}
	return ""
}

func (m *PppoeMaLimitInfo_KEYS) GetRemoteId() string {
	if m != nil {
		return m.RemoteId
	}
	return ""
}

func (m *PppoeMaLimitInfo_KEYS) GetOuterVlanId() uint32 {
	if m != nil {
		return m.OuterVlanId
	}
	return 0
}

func (m *PppoeMaLimitInfo_KEYS) GetInnerVlanId() uint32 {
	if m != nil {
		return m.InnerVlanId
	}
	return 0
}

type PppoeMaLimitInfo struct {
	State                string   `protobuf:"bytes,50,opt,name=state,proto3" json:"state,omitempty"`
	SessionCount         uint32   `protobuf:"varint,51,opt,name=session_count,json=sessionCount,proto3" json:"session_count,omitempty"`
	RadiusOverrideSet    int32    `protobuf:"zigzag32,52,opt,name=radius_override_set,json=radiusOverrideSet,proto3" json:"radius_override_set,omitempty"`
	OverrideLimit        uint32   `protobuf:"varint,53,opt,name=override_limit,json=overrideLimit,proto3" json:"override_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PppoeMaLimitInfo) Reset()         { *m = PppoeMaLimitInfo{} }
func (m *PppoeMaLimitInfo) String() string { return proto.CompactTextString(m) }
func (*PppoeMaLimitInfo) ProtoMessage()    {}
func (*PppoeMaLimitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef55ad5832e06799, []int{1}
}

func (m *PppoeMaLimitInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PppoeMaLimitInfo.Unmarshal(m, b)
}
func (m *PppoeMaLimitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PppoeMaLimitInfo.Marshal(b, m, deterministic)
}
func (m *PppoeMaLimitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PppoeMaLimitInfo.Merge(m, src)
}
func (m *PppoeMaLimitInfo) XXX_Size() int {
	return xxx_messageInfo_PppoeMaLimitInfo.Size(m)
}
func (m *PppoeMaLimitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PppoeMaLimitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PppoeMaLimitInfo proto.InternalMessageInfo

func (m *PppoeMaLimitInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *PppoeMaLimitInfo) GetSessionCount() uint32 {
	if m != nil {
		return m.SessionCount
	}
	return 0
}

func (m *PppoeMaLimitInfo) GetRadiusOverrideSet() int32 {
	if m != nil {
		return m.RadiusOverrideSet
	}
	return 0
}

func (m *PppoeMaLimitInfo) GetOverrideLimit() uint32 {
	if m != nil {
		return m.OverrideLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*PppoeMaLimitInfo_KEYS)(nil), "cisco_ios_xr_subscriber_pppoe_ma_oper.pppoe.nodes.node.bba_groups.bba_group.limits.limit.pppoe_ma_limit_info_KEYS")
	proto.RegisterType((*PppoeMaLimitInfo)(nil), "cisco_ios_xr_subscriber_pppoe_ma_oper.pppoe.nodes.node.bba_groups.bba_group.limits.limit.pppoe_ma_limit_info")
}

func init() { proto.RegisterFile("pppoe_ma_limit_info.proto", fileDescriptor_ef55ad5832e06799) }

var fileDescriptor_ef55ad5832e06799 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0xea, 0xd3, 0x30,
	0x14, 0x87, 0xe9, 0x7f, 0x6e, 0xae, 0xd9, 0x3a, 0x5c, 0xe6, 0x45, 0x44, 0xc4, 0x32, 0x15, 0x7a,
	0xd5, 0x0b, 0xa7, 0x0f, 0x20, 0x22, 0x32, 0x14, 0x85, 0x0e, 0x44, 0xaf, 0x42, 0x9a, 0x9c, 0x49,
	0x60, 0x4d, 0x4a, 0x92, 0x4e, 0xdf, 0xc8, 0x87, 0xf0, 0xe5, 0x24, 0x27, 0xb5, 0x7a, 0xb1, 0x9b,
	0x36, 0xf9, 0xce, 0x97, 0x5f, 0x0e, 0x87, 0x90, 0x47, 0x7d, 0xdf, 0x5b, 0xe0, 0x9d, 0xe0, 0x17,
	0xdd, 0xe9, 0xc0, 0xb5, 0x39, 0xdb, 0xba, 0x77, 0x36, 0x58, 0xfa, 0x55, 0x6a, 0x2f, 0x2d, 0xd7,
	0xd6, 0xf3, 0x9f, 0x8e, 0xfb, 0xa1, 0xf5, 0xd2, 0xe9, 0x16, 0x1c, 0x9f, 0x8e, 0xd8, 0x1e, 0x5c,
	0x8d, 0xbb, 0xda, 0x58, 0x05, 0x1e, 0xbf, 0x75, 0xdb, 0x0a, 0xfe, 0xdd, 0xd9, 0xa1, 0xf7, 0xff,
	0x96, 0x35, 0xa6, 0xfb, 0xf4, 0xdb, 0xff, 0xbe, 0x23, 0xec, 0xc6, 0xbd, 0xfc, 0xc3, 0xbb, 0x6f,
	0x27, 0xfa, 0x98, 0xe4, 0x31, 0x86, 0x1b, 0xd1, 0x01, 0xcb, 0xca, 0xac, 0xca, 0x9b, 0x65, 0x04,
	0x9f, 0x44, 0x07, 0xf4, 0x39, 0xd9, 0x4c, 0x99, 0xc9, 0xb8, 0x43, 0x63, 0xdd, 0xb6, 0xe2, 0x7d,
	0x84, 0x68, 0xbd, 0x20, 0x1b, 0x6d, 0x02, 0xb8, 0xb3, 0x90, 0x63, 0xce, 0x0c, 0xad, 0x62, 0xa2,
	0xa8, 0x3d, 0x25, 0xab, 0x4e, 0x48, 0x2e, 0x94, 0x72, 0xe0, 0x3d, 0xbb, 0x87, 0x0e, 0xe9, 0x84,
	0x7c, 0x93, 0x08, 0x7d, 0x40, 0x66, 0xfa, 0xc7, 0x99, 0xcd, 0xcb, 0xac, 0x5a, 0x36, 0x71, 0x49,
	0x9f, 0x10, 0x22, 0xb5, 0x93, 0x43, 0xec, 0x58, 0xb1, 0x05, 0x9e, 0xc8, 0x47, 0x72, 0x54, 0xb1,
	0x77, 0x07, 0x9d, 0x0d, 0x10, 0xab, 0xf7, 0x53, 0xef, 0x09, 0x1c, 0x15, 0xdd, 0x93, 0xc2, 0x0e,
	0x01, 0x1c, 0xbf, 0x5e, 0x84, 0x89, 0xc2, 0xb2, 0xcc, 0xaa, 0xa2, 0x59, 0x21, 0xfc, 0x72, 0x11,
	0x26, 0x39, 0xda, 0x98, 0xff, 0x9c, 0x3c, 0x39, 0x08, 0x93, 0xb3, 0xff, 0x95, 0x91, 0xdd, 0x8d,
	0xe9, 0xd1, 0x87, 0x64, 0xee, 0x83, 0x08, 0xc0, 0x5e, 0xe2, 0xc5, 0x69, 0x43, 0x9f, 0x91, 0xc2,
	0x83, 0xf7, 0xda, 0x1a, 0x2e, 0xed, 0x60, 0x02, 0x3b, 0x60, 0xe2, 0x7a, 0x84, 0x6f, 0x23, 0xa3,
	0x35, 0xd9, 0x39, 0xa1, 0xf4, 0xe0, 0xb9, 0xbd, 0x82, 0x73, 0x5a, 0x01, 0xf7, 0x10, 0xd8, 0xab,
	0x32, 0xab, 0xb6, 0xcd, 0x36, 0x95, 0x3e, 0x8f, 0x95, 0x13, 0x84, 0x38, 0xe0, 0x49, 0xc4, 0x0e,
	0xd8, 0x6b, 0x4c, 0x2d, 0xfe, 0xd2, 0x8f, 0x11, 0xb6, 0x0b, 0x7c, 0x48, 0x87, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x85, 0x97, 0x0c, 0x01, 0x65, 0x02, 0x00, 0x00,
}
