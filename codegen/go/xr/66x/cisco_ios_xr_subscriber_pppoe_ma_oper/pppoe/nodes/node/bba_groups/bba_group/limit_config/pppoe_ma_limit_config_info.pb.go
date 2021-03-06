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
// source: pppoe_ma_limit_config_info.proto

package cisco_ios_xr_subscriber_pppoe_ma_oper_pppoe_nodes_node_bba_groups_bba_group_limit_config

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

type PppoeMaLimitConfigInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	BbaGroupName         string   `protobuf:"bytes,2,opt,name=bba_group_name,json=bbaGroupName,proto3" json:"bba_group_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PppoeMaLimitConfigInfo_KEYS) Reset()         { *m = PppoeMaLimitConfigInfo_KEYS{} }
func (m *PppoeMaLimitConfigInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PppoeMaLimitConfigInfo_KEYS) ProtoMessage()    {}
func (*PppoeMaLimitConfigInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a93f5baff7dc950f, []int{0}
}

func (m *PppoeMaLimitConfigInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PppoeMaLimitConfigInfo_KEYS.Unmarshal(m, b)
}
func (m *PppoeMaLimitConfigInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PppoeMaLimitConfigInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PppoeMaLimitConfigInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PppoeMaLimitConfigInfo_KEYS.Merge(m, src)
}
func (m *PppoeMaLimitConfigInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PppoeMaLimitConfigInfo_KEYS.Size(m)
}
func (m *PppoeMaLimitConfigInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PppoeMaLimitConfigInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PppoeMaLimitConfigInfo_KEYS proto.InternalMessageInfo

func (m *PppoeMaLimitConfigInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PppoeMaLimitConfigInfo_KEYS) GetBbaGroupName() string {
	if m != nil {
		return m.BbaGroupName
	}
	return ""
}

type PppoeMaLimitConfigSingle struct {
	MaxLimit              uint32   `protobuf:"varint,1,opt,name=max_limit,json=maxLimit,proto3" json:"max_limit,omitempty"`
	Threshold             uint32   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	RadiusOverrideEnabled int32    `protobuf:"zigzag32,3,opt,name=radius_override_enabled,json=radiusOverrideEnabled,proto3" json:"radius_override_enabled,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *PppoeMaLimitConfigSingle) Reset()         { *m = PppoeMaLimitConfigSingle{} }
func (m *PppoeMaLimitConfigSingle) String() string { return proto.CompactTextString(m) }
func (*PppoeMaLimitConfigSingle) ProtoMessage()    {}
func (*PppoeMaLimitConfigSingle) Descriptor() ([]byte, []int) {
	return fileDescriptor_a93f5baff7dc950f, []int{1}
}

func (m *PppoeMaLimitConfigSingle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PppoeMaLimitConfigSingle.Unmarshal(m, b)
}
func (m *PppoeMaLimitConfigSingle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PppoeMaLimitConfigSingle.Marshal(b, m, deterministic)
}
func (m *PppoeMaLimitConfigSingle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PppoeMaLimitConfigSingle.Merge(m, src)
}
func (m *PppoeMaLimitConfigSingle) XXX_Size() int {
	return xxx_messageInfo_PppoeMaLimitConfigSingle.Size(m)
}
func (m *PppoeMaLimitConfigSingle) XXX_DiscardUnknown() {
	xxx_messageInfo_PppoeMaLimitConfigSingle.DiscardUnknown(m)
}

var xxx_messageInfo_PppoeMaLimitConfigSingle proto.InternalMessageInfo

func (m *PppoeMaLimitConfigSingle) GetMaxLimit() uint32 {
	if m != nil {
		return m.MaxLimit
	}
	return 0
}

func (m *PppoeMaLimitConfigSingle) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PppoeMaLimitConfigSingle) GetRadiusOverrideEnabled() int32 {
	if m != nil {
		return m.RadiusOverrideEnabled
	}
	return 0
}

type PppoeMaLimitConfigInfo struct {
	Card                  *PppoeMaLimitConfigSingle `protobuf:"bytes,50,opt,name=card,proto3" json:"card,omitempty"`
	AccessIntf            *PppoeMaLimitConfigSingle `protobuf:"bytes,51,opt,name=access_intf,json=accessIntf,proto3" json:"access_intf,omitempty"`
	Mac                   *PppoeMaLimitConfigSingle `protobuf:"bytes,52,opt,name=mac,proto3" json:"mac,omitempty"`
	MacIwf                *PppoeMaLimitConfigSingle `protobuf:"bytes,53,opt,name=mac_iwf,json=macIwf,proto3" json:"mac_iwf,omitempty"`
	MacAccessInterface    *PppoeMaLimitConfigSingle `protobuf:"bytes,54,opt,name=mac_access_interface,json=macAccessInterface,proto3" json:"mac_access_interface,omitempty"`
	MacIwfAccessInterface *PppoeMaLimitConfigSingle `protobuf:"bytes,55,opt,name=mac_iwf_access_interface,json=macIwfAccessInterface,proto3" json:"mac_iwf_access_interface,omitempty"`
	CircuitId             *PppoeMaLimitConfigSingle `protobuf:"bytes,56,opt,name=circuit_id,json=circuitId,proto3" json:"circuit_id,omitempty"`
	RemoteId              *PppoeMaLimitConfigSingle `protobuf:"bytes,57,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
	CircuitIdAndRemoteId  *PppoeMaLimitConfigSingle `protobuf:"bytes,58,opt,name=circuit_id_and_remote_id,json=circuitIdAndRemoteId,proto3" json:"circuit_id_and_remote_id,omitempty"`
	OuterVlanId           *PppoeMaLimitConfigSingle `protobuf:"bytes,59,opt,name=outer_vlan_id,json=outerVlanId,proto3" json:"outer_vlan_id,omitempty"`
	InnerVlanId           *PppoeMaLimitConfigSingle `protobuf:"bytes,60,opt,name=inner_vlan_id,json=innerVlanId,proto3" json:"inner_vlan_id,omitempty"`
	VlanId                *PppoeMaLimitConfigSingle `protobuf:"bytes,61,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                  `json:"-"`
	XXX_unrecognized      []byte                    `json:"-"`
	XXX_sizecache         int32                     `json:"-"`
}

func (m *PppoeMaLimitConfigInfo) Reset()         { *m = PppoeMaLimitConfigInfo{} }
func (m *PppoeMaLimitConfigInfo) String() string { return proto.CompactTextString(m) }
func (*PppoeMaLimitConfigInfo) ProtoMessage()    {}
func (*PppoeMaLimitConfigInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a93f5baff7dc950f, []int{2}
}

func (m *PppoeMaLimitConfigInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PppoeMaLimitConfigInfo.Unmarshal(m, b)
}
func (m *PppoeMaLimitConfigInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PppoeMaLimitConfigInfo.Marshal(b, m, deterministic)
}
func (m *PppoeMaLimitConfigInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PppoeMaLimitConfigInfo.Merge(m, src)
}
func (m *PppoeMaLimitConfigInfo) XXX_Size() int {
	return xxx_messageInfo_PppoeMaLimitConfigInfo.Size(m)
}
func (m *PppoeMaLimitConfigInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PppoeMaLimitConfigInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PppoeMaLimitConfigInfo proto.InternalMessageInfo

func (m *PppoeMaLimitConfigInfo) GetCard() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetAccessIntf() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.AccessIntf
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetMac() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.Mac
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetMacIwf() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.MacIwf
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetMacAccessInterface() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.MacAccessInterface
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetMacIwfAccessInterface() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.MacIwfAccessInterface
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetCircuitId() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.CircuitId
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetRemoteId() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.RemoteId
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetCircuitIdAndRemoteId() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.CircuitIdAndRemoteId
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetOuterVlanId() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.OuterVlanId
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetInnerVlanId() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.InnerVlanId
	}
	return nil
}

func (m *PppoeMaLimitConfigInfo) GetVlanId() *PppoeMaLimitConfigSingle {
	if m != nil {
		return m.VlanId
	}
	return nil
}

func init() {
	proto.RegisterType((*PppoeMaLimitConfigInfo_KEYS)(nil), "cisco_ios_xr_subscriber_pppoe_ma_oper.pppoe.nodes.node.bba_groups.bba_group.limit_config.pppoe_ma_limit_config_info_KEYS")
	proto.RegisterType((*PppoeMaLimitConfigSingle)(nil), "cisco_ios_xr_subscriber_pppoe_ma_oper.pppoe.nodes.node.bba_groups.bba_group.limit_config.pppoe_ma_limit_config_single")
	proto.RegisterType((*PppoeMaLimitConfigInfo)(nil), "cisco_ios_xr_subscriber_pppoe_ma_oper.pppoe.nodes.node.bba_groups.bba_group.limit_config.pppoe_ma_limit_config_info")
}

func init() { proto.RegisterFile("pppoe_ma_limit_config_info.proto", fileDescriptor_a93f5baff7dc950f) }

var fileDescriptor_a93f5baff7dc950f = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0xd6, 0x4f, 0x8b, 0x13, 0x31,
	0x14, 0x00, 0x70, 0xe2, 0xca, 0xda, 0xbe, 0xda, 0x05, 0xc3, 0x2e, 0x06, 0x5d, 0xb0, 0x14, 0x0f,
	0x3d, 0xcd, 0x61, 0x57, 0xd7, 0xff, 0x87, 0x3d, 0x2c, 0x52, 0x14, 0x85, 0x11, 0x44, 0x4f, 0x21,
	0x93, 0x64, 0xba, 0x81, 0x49, 0x32, 0x24, 0xd3, 0x6e, 0x3f, 0x82, 0x78, 0x11, 0x11, 0x11, 0xfc,
	0x04, 0x9e, 0xf4, 0x6b, 0xf9, 0x31, 0x24, 0x99, 0xb6, 0xb3, 0xa8, 0xeb, 0x4d, 0x72, 0x29, 0x4d,
	0x5e, 0xe6, 0xbd, 0xdf, 0xbc, 0x07, 0x61, 0x60, 0x54, 0xd7, 0xb5, 0x95, 0x54, 0x33, 0x5a, 0x29,
	0xad, 0x1a, 0xca, 0xad, 0x29, 0xd5, 0x8c, 0x2a, 0x53, 0xda, 0xac, 0x76, 0xb6, 0xb1, 0xf8, 0x0d,
	0x57, 0x9e, 0x5b, 0xaa, 0xac, 0xa7, 0x4b, 0x47, 0xfd, 0xbc, 0xf0, 0xdc, 0xa9, 0x42, 0x3a, 0xba,
	0x79, 0xd2, 0xd6, 0xd2, 0x65, 0x71, 0x95, 0x19, 0x2b, 0xa4, 0x8f, 0xbf, 0x59, 0x51, 0x30, 0x3a,
	0x73, 0x76, 0x5e, 0xfb, 0xee, 0x6f, 0x76, 0xbe, 0xc8, 0x58, 0xc0, 0xad, 0x8b, 0xab, 0xd3, 0x67,
	0x27, 0x6f, 0x5f, 0xe1, 0x9b, 0xd0, 0x0f, 0xc9, 0xa8, 0x61, 0x5a, 0x12, 0x34, 0x42, 0x93, 0x7e,
	0xde, 0x0b, 0x1b, 0x2f, 0x98, 0x96, 0xf8, 0x36, 0xec, 0x6c, 0x32, 0xb7, 0x27, 0x2e, 0xc5, 0x13,
	0x57, 0x8b, 0x82, 0x3d, 0x0d, 0x9b, 0xe1, 0xd4, 0xf8, 0x23, 0x82, 0xfd, 0xbf, 0x97, 0xf1, 0xca,
	0xcc, 0x2a, 0x19, 0x6a, 0x68, 0xb6, 0x6c, 0x43, 0xb1, 0xc6, 0x30, 0xef, 0x69, 0xb6, 0x7c, 0x1e,
	0xd6, 0x78, 0x1f, 0xfa, 0xcd, 0xa9, 0x93, 0xfe, 0xd4, 0x56, 0x22, 0xa6, 0x1f, 0xe6, 0xdd, 0x06,
	0x3e, 0x82, 0xeb, 0x8e, 0x09, 0x35, 0xf7, 0xd4, 0x2e, 0xa4, 0x73, 0x4a, 0x48, 0x2a, 0x0d, 0x2b,
	0x2a, 0x29, 0xc8, 0xd6, 0x08, 0x4d, 0xae, 0xe5, 0x7b, 0x6d, 0xf8, 0xe5, 0x2a, 0x7a, 0xd2, 0x06,
	0xc7, 0x3f, 0x77, 0xe0, 0xc6, 0xc5, 0xaf, 0x8e, 0xdf, 0x23, 0xb8, 0xcc, 0x99, 0x13, 0xe4, 0x60,
	0x84, 0x26, 0x83, 0x83, 0x45, 0xf6, 0xbf, 0x46, 0x90, 0xfd, 0xab, 0x31, 0x79, 0x34, 0xe0, 0x2f,
	0x08, 0x06, 0x8c, 0x73, 0xe9, 0x3d, 0x55, 0xa6, 0x29, 0xc9, 0x61, 0x52, 0x13, 0xb4, 0x94, 0xa9,
	0x69, 0x4a, 0xfc, 0x0e, 0xc1, 0x96, 0x66, 0x9c, 0xdc, 0x49, 0x2a, 0x0a, 0x04, 0xfc, 0x01, 0xc1,
	0x15, 0xcd, 0x38, 0x55, 0x67, 0x25, 0xb9, 0x9b, 0x94, 0xb3, 0xad, 0x19, 0x9f, 0x9e, 0x95, 0xf8,
	0x1b, 0x82, 0xdd, 0x20, 0xea, 0x46, 0x27, 0x5d, 0xc9, 0xb8, 0x24, 0x47, 0x49, 0x79, 0x58, 0x33,
	0x7e, 0xbc, 0x1e, 0x61, 0x2b, 0xc2, 0x3f, 0x10, 0x90, 0x55, 0xf3, 0xfe, 0xe4, 0xde, 0x4b, 0xca,
	0xdd, 0x6b, 0xbb, 0xf9, 0xbb, 0xf8, 0x33, 0x02, 0xe0, 0xca, 0xf1, 0xb9, 0x6a, 0xa8, 0x12, 0xe4,
	0x7e, 0x52, 0x63, 0x7f, 0x25, 0x99, 0x0a, 0xfc, 0x09, 0x41, 0xdf, 0x49, 0x6d, 0x1b, 0x19, 0x58,
	0x0f, 0x92, 0xb2, 0x7a, 0x2d, 0x64, 0x2a, 0xf0, 0x77, 0x04, 0xa4, 0xeb, 0x16, 0x65, 0x46, 0xd0,
	0x0e, 0xf9, 0x30, 0x29, 0x72, 0x77, 0xd3, 0xbb, 0x63, 0x23, 0xf2, 0x35, 0xf8, 0x2b, 0x82, 0xa1,
	0x9d, 0x37, 0xd2, 0xd1, 0x45, 0xc5, 0x4c, 0x50, 0x3e, 0x4a, 0xaa, 0x1c, 0x44, 0xcc, 0xeb, 0x8a,
	0x99, 0x15, 0x4e, 0x19, 0x73, 0x0e, 0xf7, 0x38, 0x2d, 0x2e, 0x62, 0x56, 0xb8, 0x70, 0x0f, 0xae,
	0x59, 0x4f, 0xd2, 0xde, 0x83, 0x8b, 0x28, 0x2a, 0xb6, 0xe3, 0x57, 0xcc, 0xe1, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x54, 0x5b, 0xae, 0xa3, 0xe9, 0x08, 0x00, 0x00,
}
