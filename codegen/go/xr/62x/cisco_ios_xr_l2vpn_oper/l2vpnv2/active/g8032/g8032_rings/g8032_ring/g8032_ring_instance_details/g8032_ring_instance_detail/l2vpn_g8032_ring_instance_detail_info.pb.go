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
// source: l2vpn_g8032_ring_instance_detail_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_active_g8032_g8032_rings_g8032_ring_g8032_ring_instance_details_g8032_ring_instance_detail

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

type L2VpnG8032RingInstanceDetailInfo_KEYS struct {
	RingName             string   `protobuf:"bytes,1,opt,name=ring_name,json=ringName,proto3" json:"ring_name,omitempty"`
	Instance             uint32   `protobuf:"varint,2,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnG8032RingInstanceDetailInfo_KEYS) Reset()         { *m = L2VpnG8032RingInstanceDetailInfo_KEYS{} }
func (m *L2VpnG8032RingInstanceDetailInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnG8032RingInstanceDetailInfo_KEYS) ProtoMessage()    {}
func (*L2VpnG8032RingInstanceDetailInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9766fb27c8712c17, []int{0}
}

func (m *L2VpnG8032RingInstanceDetailInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo_KEYS.Unmarshal(m, b)
}
func (m *L2VpnG8032RingInstanceDetailInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnG8032RingInstanceDetailInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo_KEYS.Merge(m, src)
}
func (m *L2VpnG8032RingInstanceDetailInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo_KEYS.Size(m)
}
func (m *L2VpnG8032RingInstanceDetailInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo_KEYS proto.InternalMessageInfo

func (m *L2VpnG8032RingInstanceDetailInfo_KEYS) GetRingName() string {
	if m != nil {
		return m.RingName
	}
	return ""
}

func (m *L2VpnG8032RingInstanceDetailInfo_KEYS) GetInstance() uint32 {
	if m != nil {
		return m.Instance
	}
	return 0
}

type L2VpnG8032RingInstanceDetailInfo struct {
	RingInstance         uint32   `protobuf:"varint,50,opt,name=ring_instance,json=ringInstance,proto3" json:"ring_instance,omitempty"`
	IncludedVlanId       []uint32 `protobuf:"varint,51,rep,packed,name=included_vlan_id,json=includedVlanId,proto3" json:"included_vlan_id,omitempty"`
	RingDescription      string   `protobuf:"bytes,52,opt,name=ring_description,json=ringDescription,proto3" json:"ring_description,omitempty"`
	RingProfile          string   `protobuf:"bytes,53,opt,name=ring_profile,json=ringProfile,proto3" json:"ring_profile,omitempty"`
	Rpl                  string   `protobuf:"bytes,54,opt,name=rpl,proto3" json:"rpl,omitempty"`
	ApsPort0             string   `protobuf:"bytes,55,opt,name=aps_port0,json=apsPort0,proto3" json:"aps_port0,omitempty"`
	ApsPort1             string   `protobuf:"bytes,56,opt,name=aps_port1,json=apsPort1,proto3" json:"aps_port1,omitempty"`
	ConfigState          bool     `protobuf:"varint,57,opt,name=config_state,json=configState,proto3" json:"config_state,omitempty"`
	UnresolvedReason     string   `protobuf:"bytes,58,opt,name=unresolved_reason,json=unresolvedReason,proto3" json:"unresolved_reason,omitempty"`
	ApsChannelLevel      uint32   `protobuf:"varint,59,opt,name=aps_channel_level,json=apsChannelLevel,proto3" json:"aps_channel_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnG8032RingInstanceDetailInfo) Reset()         { *m = L2VpnG8032RingInstanceDetailInfo{} }
func (m *L2VpnG8032RingInstanceDetailInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnG8032RingInstanceDetailInfo) ProtoMessage()    {}
func (*L2VpnG8032RingInstanceDetailInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9766fb27c8712c17, []int{1}
}

func (m *L2VpnG8032RingInstanceDetailInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo.Unmarshal(m, b)
}
func (m *L2VpnG8032RingInstanceDetailInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnG8032RingInstanceDetailInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo.Merge(m, src)
}
func (m *L2VpnG8032RingInstanceDetailInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo.Size(m)
}
func (m *L2VpnG8032RingInstanceDetailInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnG8032RingInstanceDetailInfo proto.InternalMessageInfo

func (m *L2VpnG8032RingInstanceDetailInfo) GetRingInstance() uint32 {
	if m != nil {
		return m.RingInstance
	}
	return 0
}

func (m *L2VpnG8032RingInstanceDetailInfo) GetIncludedVlanId() []uint32 {
	if m != nil {
		return m.IncludedVlanId
	}
	return nil
}

func (m *L2VpnG8032RingInstanceDetailInfo) GetRingDescription() string {
	if m != nil {
		return m.RingDescription
	}
	return ""
}

func (m *L2VpnG8032RingInstanceDetailInfo) GetRingProfile() string {
	if m != nil {
		return m.RingProfile
	}
	return ""
}

func (m *L2VpnG8032RingInstanceDetailInfo) GetRpl() string {
	if m != nil {
		return m.Rpl
	}
	return ""
}

func (m *L2VpnG8032RingInstanceDetailInfo) GetApsPort0() string {
	if m != nil {
		return m.ApsPort0
	}
	return ""
}

func (m *L2VpnG8032RingInstanceDetailInfo) GetApsPort1() string {
	if m != nil {
		return m.ApsPort1
	}
	return ""
}

func (m *L2VpnG8032RingInstanceDetailInfo) GetConfigState() bool {
	if m != nil {
		return m.ConfigState
	}
	return false
}

func (m *L2VpnG8032RingInstanceDetailInfo) GetUnresolvedReason() string {
	if m != nil {
		return m.UnresolvedReason
	}
	return ""
}

func (m *L2VpnG8032RingInstanceDetailInfo) GetApsChannelLevel() uint32 {
	if m != nil {
		return m.ApsChannelLevel
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnG8032RingInstanceDetailInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.g8032.g8032_rings.g8032_ring.g8032_ring_instance_details.g8032_ring_instance_detail.l2vpn_g8032_ring_instance_detail_info_KEYS")
	proto.RegisterType((*L2VpnG8032RingInstanceDetailInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.g8032.g8032_rings.g8032_ring.g8032_ring_instance_details.g8032_ring_instance_detail.l2vpn_g8032_ring_instance_detail_info")
}

func init() {
	proto.RegisterFile("l2vpn_g8032_ring_instance_detail_info.proto", fileDescriptor_9766fb27c8712c17)
}

var fileDescriptor_9766fb27c8712c17 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0xab, 0x13, 0x31,
	0x14, 0xc5, 0x19, 0x0b, 0xd2, 0x97, 0x67, 0x7d, 0xf3, 0xb2, 0x0a, 0xba, 0x19, 0x9f, 0x08, 0xe3,
	0x2b, 0x0c, 0xed, 0xd4, 0x3f, 0x55, 0x97, 0xea, 0xa2, 0x28, 0x52, 0xa6, 0x20, 0xb8, 0xba, 0xc4,
	0xcc, 0x6d, 0x0d, 0xa4, 0x49, 0x48, 0xd2, 0x41, 0xfc, 0x0c, 0x7e, 0x68, 0x49, 0xc6, 0xda, 0x76,
	0xe1, 0xa3, 0xbb, 0x9b, 0xdf, 0xb9, 0xf7, 0x9c, 0xcb, 0x25, 0x64, 0xac, 0xea, 0xce, 0x6a, 0xd8,
	0xcc, 0x27, 0xb3, 0x1a, 0x9c, 0xd4, 0x1b, 0x90, 0xda, 0x07, 0xae, 0x05, 0x42, 0x8b, 0x81, 0x4b,
	0x05, 0x52, 0xaf, 0x4d, 0x65, 0x9d, 0x09, 0x86, 0xfe, 0x12, 0xd2, 0x0b, 0x03, 0xd2, 0x78, 0xf8,
	0xe9, 0xa0, 0x9f, 0x34, 0x16, 0x5d, 0x95, 0xca, 0xae, 0xae, 0xb8, 0x08, 0xb2, 0xc3, 0x2a, 0xb9,
	0x55, 0x07, 0x4f, 0x7f, 0x54, 0x57, 0xff, 0x8f, 0xf2, 0x77, 0x68, 0x37, 0x48, 0x6e, 0xcf, 0x5a,
	0x15, 0x3e, 0x7d, 0xfc, 0xb6, 0xa2, 0x8f, 0xc9, 0x45, 0xea, 0xd0, 0x7c, 0x8b, 0x2c, 0x2b, 0xb2,
	0xf2, 0xa2, 0x19, 0x46, 0xf0, 0x85, 0x6f, 0x91, 0x3e, 0x22, 0xc3, 0xfd, 0x24, 0xbb, 0x57, 0x64,
	0xe5, 0xa8, 0xf9, 0xf7, 0xbe, 0xf9, 0x3d, 0x20, 0xcf, 0xce, 0xca, 0xa1, 0x4f, 0xc9, 0xe8, 0x44,
	0x64, 0x75, 0xb2, 0x7a, 0x10, 0xe1, 0xe2, 0x2f, 0xa3, 0x25, 0xc9, 0xa5, 0x16, 0x6a, 0xd7, 0x62,
	0x0b, 0x9d, 0xe2, 0x1a, 0x64, 0xcb, 0x66, 0xc5, 0xa0, 0x1c, 0x35, 0x0f, 0xf7, 0xfc, 0xab, 0xe2,
	0x7a, 0xd1, 0xd2, 0xe7, 0x24, 0x4f, 0x76, 0x2d, 0x7a, 0xe1, 0xa4, 0x0d, 0xd2, 0x68, 0xf6, 0x22,
	0x2d, 0x7e, 0x15, 0xf9, 0x87, 0x03, 0xa6, 0x4f, 0x48, 0x0a, 0x01, 0xeb, 0xcc, 0x5a, 0x2a, 0x64,
	0x2f, 0x53, 0xdb, 0x65, 0x64, 0xcb, 0x1e, 0xd1, 0x9c, 0x0c, 0x9c, 0x55, 0xec, 0x55, 0x52, 0x62,
	0x19, 0x2f, 0xc2, 0xad, 0x07, 0x6b, 0x5c, 0x98, 0xb0, 0xd7, 0xfd, 0x45, 0xb8, 0xf5, 0xcb, 0xf8,
	0x3e, 0x16, 0xa7, 0x6c, 0x7e, 0x22, 0x4e, 0x63, 0x9c, 0x30, 0x7a, 0x2d, 0x37, 0xe0, 0x03, 0x0f,
	0xc8, 0xde, 0x14, 0x59, 0x39, 0x6c, 0x2e, 0x7b, 0xb6, 0x8a, 0x88, 0x8e, 0xc9, 0xf5, 0x4e, 0x3b,
	0xf4, 0x46, 0x75, 0xd8, 0x82, 0x43, 0xee, 0x8d, 0x66, 0x6f, 0x93, 0x4f, 0x7e, 0x10, 0x9a, 0xc4,
	0xe9, 0x2d, 0xb9, 0x8e, 0x61, 0xe2, 0x07, 0xd7, 0x1a, 0x15, 0x28, 0xec, 0x50, 0xb1, 0x77, 0xe9,
	0x78, 0x57, 0xdc, 0xfa, 0xf7, 0x3d, 0xff, 0x1c, 0xf1, 0xf7, 0xfb, 0xe9, 0xe3, 0xcd, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x37, 0x0a, 0x17, 0x24, 0xa7, 0x02, 0x00, 0x00,
}
