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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_attributes_pwg_info

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
	Name_1               string   `protobuf:"bytes,2,opt,name=name_1,json=name1,proto3" json:"name_1,omitempty"`
	Name_2               string   `protobuf:"bytes,3,opt,name=name_2,json=name2,proto3" json:"name_2,omitempty"`
	Name_3               string   `protobuf:"bytes,4,opt,name=name_3,json=name3,proto3" json:"name_3,omitempty"`
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

func (m *InvPwgInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *InvPwgInfo_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *InvPwgInfo_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
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
	proto.RegisterType((*InvPwgInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.pwg_info.inv_pwg_info_KEYS")
	proto.RegisterType((*InvPwgInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.pwg_info.inv_pwg_info")
}

func init() { proto.RegisterFile("inv_pwg_info.proto", fileDescriptor_aa6fe1f6f982d195) }

var fileDescriptor_aa6fe1f6f982d195 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0x69, 0xbf, 0xfd, 0x16, 0x0c, 0x5e, 0x1a, 0xfc, 0x11, 0x15, 0xa1, 0xf4, 0xd4, 0x53,
	0xa0, 0xbb, 0xf6, 0xe0, 0xc1, 0x43, 0x51, 0xf1, 0x20, 0x82, 0x54, 0x3c, 0x78, 0x0a, 0xe9, 0x6e,
	0x2c, 0xc1, 0xed, 0xcc, 0x32, 0xc9, 0x6e, 0xed, 0xd5, 0xbf, 0x5c, 0x92, 0xda, 0xdd, 0xbd, 0x4c,
	0x5e, 0xde, 0xbc, 0x4f, 0xde, 0x21, 0x8c, 0x5b, 0xa8, 0x55, 0xb9, 0x5d, 0x2b, 0x0b, 0x9f, 0x28,
	0x4b, 0x42, 0x8f, 0x3c, 0xcb, 0xac, 0xcb, 0x50, 0x59, 0x74, 0xea, 0x9b, 0x94, 0x85, 0x7a, 0xb3,
	0x26, 0x85, 0xa5, 0x21, 0x69, 0xa1, 0x36, 0xe0, 0x91, 0x76, 0x92, 0x74, 0xf6, 0xe5, 0xe2, 0x94,
	0x06, 0xbc, 0xf5, 0x3b, 0xe9, 0x0a, 0xf4, 0xd2, 0x3b, 0x3b, 0x73, 0x71, 0x86, 0x91, 0x44, 0x99,
	0x48, 0xed, 0x3d, 0xd9, 0x55, 0xe5, 0x8d, 0x93, 0x87, 0xaa, 0x49, 0xc1, 0x46, 0xdd, 0x6a, 0xf5,
	0xfc, 0xf8, 0xf1, 0xc6, 0x39, 0x1b, 0x80, 0xde, 0x18, 0xd1, 0x1b, 0xf7, 0xa6, 0x47, 0xcb, 0xa8,
	0xf9, 0x29, 0x1b, 0x86, 0x53, 0xcd, 0x44, 0x3f, 0xba, 0xff, 0xc3, 0x6d, 0xd6, 0xd8, 0x89, 0xf8,
	0xd7, 0xda, 0x49, 0x63, 0xa7, 0x62, 0xd0, 0xda, 0xe9, 0xe4, 0xa7, 0xcf, 0x8e, 0xbb, 0x75, 0xfc,
	0x8e, 0x5d, 0x95, 0xb8, 0x35, 0xa4, 0xd6, 0x84, 0x55, 0xa9, 0xc8, 0xe4, 0x15, 0xe4, 0x1a, 0xb2,
	0x9d, 0xda, 0x60, 0x6e, 0x44, 0x32, 0xee, 0x4d, 0x47, 0x4b, 0x11, 0x23, 0x4f, 0x21, 0xb1, 0x6c,
	0x02, 0x2f, 0x98, 0x1b, 0x3e, 0x67, 0xe7, 0x5d, 0x7c, 0xaf, 0x2b, 0xb0, 0xde, 0x89, 0x34, 0xf6,
	0x9e, 0xb4, 0xe8, 0x6b, 0x50, 0xef, 0x61, 0xc7, 0x17, 0xec, 0xba, 0x8b, 0xe9, 0x5a, 0xdb, 0x42,
	0xaf, 0x0a, 0xa3, 0xb2, 0x8a, 0xc8, 0x80, 0x17, 0x37, 0xb1, 0xf7, 0xb2, 0x85, 0x17, 0x87, 0xc8,
	0xfd, 0x3e, 0xc1, 0x6f, 0xd9, 0x45, 0xf7, 0x89, 0x9c, 0xf4, 0x16, 0x1a, 0x7c, 0x1e, 0xf1, 0xb3,
	0x16, 0x7f, 0x08, 0xeb, 0x3f, 0x74, 0x35, 0x8c, 0xdf, 0x9b, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xe5, 0x7e, 0xbe, 0xf9, 0xf4, 0x01, 0x00, 0x00,
}
