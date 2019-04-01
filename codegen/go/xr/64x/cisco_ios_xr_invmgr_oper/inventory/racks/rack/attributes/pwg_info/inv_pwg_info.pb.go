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
// source: inv_pwg_info.proto

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_attributes_pwg_info

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

type InvPwgInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvPwgInfo_KEYS) Reset()         { *m = InvPwgInfo_KEYS{} }
func (m *InvPwgInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*InvPwgInfo_KEYS) ProtoMessage()    {}
func (*InvPwgInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa6fe1f6f982d195, []int{0}
}

func (m *InvPwgInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvPwgInfo_KEYS.Unmarshal(m, b)
}
func (m *InvPwgInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvPwgInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *InvPwgInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvPwgInfo_KEYS.Merge(m, src)
}
func (m *InvPwgInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_InvPwgInfo_KEYS.Size(m)
}
func (m *InvPwgInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InvPwgInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InvPwgInfo_KEYS proto.InternalMessageInfo

func (m *InvPwgInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type InvPwgInfo struct {
	PowerGroupRedundancyMode   int32    `protobuf:"zigzag32,50,opt,name=power_group_redundancy_mode,json=powerGroupRedundancyMode,proto3" json:"power_group_redundancy_mode,omitempty"`
	PowerGroupPowerUnits       string   `protobuf:"bytes,51,opt,name=power_group_power_units,json=powerGroupPowerUnits,proto3" json:"power_group_power_units,omitempty"`
	PowerGroupAvailableCurrent int32    `protobuf:"zigzag32,52,opt,name=power_group_available_current,json=powerGroupAvailableCurrent,proto3" json:"power_group_available_current,omitempty"`
	PowerGroupDrawnCurrent     int32    `protobuf:"zigzag32,53,opt,name=power_group_drawn_current,json=powerGroupDrawnCurrent,proto3" json:"power_group_drawn_current,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *InvPwgInfo) Reset()         { *m = InvPwgInfo{} }
func (m *InvPwgInfo) String() string { return proto.CompactTextString(m) }
func (*InvPwgInfo) ProtoMessage()    {}
func (*InvPwgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa6fe1f6f982d195, []int{1}
}

func (m *InvPwgInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvPwgInfo.Unmarshal(m, b)
}
func (m *InvPwgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvPwgInfo.Marshal(b, m, deterministic)
}
func (m *InvPwgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvPwgInfo.Merge(m, src)
}
func (m *InvPwgInfo) XXX_Size() int {
	return xxx_messageInfo_InvPwgInfo.Size(m)
}
func (m *InvPwgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InvPwgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InvPwgInfo proto.InternalMessageInfo

func (m *InvPwgInfo) GetPowerGroupRedundancyMode() int32 {
	if m != nil {
		return m.PowerGroupRedundancyMode
	}
	return 0
}

func (m *InvPwgInfo) GetPowerGroupPowerUnits() string {
	if m != nil {
		return m.PowerGroupPowerUnits
	}
	return ""
}

func (m *InvPwgInfo) GetPowerGroupAvailableCurrent() int32 {
	if m != nil {
		return m.PowerGroupAvailableCurrent
	}
	return 0
}

func (m *InvPwgInfo) GetPowerGroupDrawnCurrent() int32 {
	if m != nil {
		return m.PowerGroupDrawnCurrent
	}
	return 0
}

func init() {
	proto.RegisterType((*InvPwgInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.attributes.pwg_info.inv_pwg_info_KEYS")
	proto.RegisterType((*InvPwgInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.attributes.pwg_info.inv_pwg_info")
}

func init() { proto.RegisterFile("inv_pwg_info.proto", fileDescriptor_aa6fe1f6f982d195) }

var fileDescriptor_aa6fe1f6f982d195 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x89, 0x88, 0xe0, 0x62, 0x73, 0x8b, 0x68, 0x54, 0x84, 0xe3, 0x1a, 0xaf, 0x4a, 0xe1,
	0x79, 0x85, 0x85, 0x45, 0x50, 0xb1, 0x10, 0x41, 0x22, 0x16, 0x56, 0xc3, 0x26, 0x19, 0xc3, 0xe2,
	0x65, 0x26, 0x4c, 0x36, 0x89, 0xd7, 0xfa, 0x97, 0x4b, 0xf6, 0x7e, 0x24, 0xcd, 0xf2, 0x60, 0xbe,
	0xef, 0x3d, 0x58, 0xa5, 0x2d, 0xb5, 0x50, 0x75, 0x05, 0x58, 0xfa, 0xe6, 0xa8, 0x12, 0x76, 0xac,
	0xe3, 0xcc, 0xd6, 0x19, 0x83, 0xe5, 0x1a, 0x7e, 0x05, 0x2c, 0xb5, 0x65, 0x21, 0xc0, 0x15, 0x4a,
	0x64, 0xa9, 0x45, 0x72, 0x2c, 0xeb, 0x48, 0x4c, 0xf6, 0x53, 0xfb, 0x37, 0x32, 0xce, 0x89, 0x4d,
	0x1b, 0x87, 0x75, 0xb4, 0x2b, 0x9a, 0xdd, 0xa8, 0xc9, 0xb8, 0x18, 0x5e, 0x9f, 0xbf, 0x3e, 0xb4,
	0x56, 0x87, 0x64, 0x4a, 0x0c, 0x83, 0x69, 0x30, 0x3f, 0x4e, 0x7c, 0x9e, 0xfd, 0x1d, 0xa8, 0x93,
	0x31, 0xa9, 0x1f, 0xd4, 0x55, 0xc5, 0x1d, 0x0a, 0x14, 0xc2, 0x4d, 0x05, 0x82, 0x79, 0x43, 0xb9,
	0xa1, 0x6c, 0x0d, 0x25, 0xe7, 0x18, 0xde, 0x4e, 0x83, 0xf9, 0x24, 0x09, 0x3d, 0xf2, 0xd2, 0x13,
	0xc9, 0x1e, 0x78, 0xe3, 0x1c, 0xf5, 0x52, 0x9d, 0x8f, 0xf5, 0x4d, 0x6e, 0xc8, 0xba, 0x3a, 0x5c,
	0xf8, 0xd9, 0xd3, 0x41, 0x7d, 0xef, 0xd3, 0x67, 0x7f, 0xd3, 0xb1, 0xba, 0x1e, 0x6b, 0xa6, 0x35,
	0x76, 0x65, 0xd2, 0x15, 0x42, 0xd6, 0x88, 0x20, 0xb9, 0xf0, 0xce, 0xef, 0x5e, 0x0e, 0x72, 0xbc,
	0x43, 0x1e, 0x37, 0x84, 0xbe, 0x57, 0x17, 0xe3, 0x8a, 0x5c, 0x4c, 0x47, 0x7b, 0x7d, 0xe9, 0xf5,
	0xb3, 0x41, 0x7f, 0xea, 0xcf, 0x5b, 0x35, 0x3d, 0xf2, 0xff, 0xbe, 0xf8, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0x09, 0x04, 0xf7, 0x10, 0x8d, 0x01, 0x00, 0x00,
}