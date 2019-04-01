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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_attributes_pwg_info

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
	proto.RegisterType((*InvPwgInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.pwg_info.inv_pwg_info_KEYS")
	proto.RegisterType((*InvPwgInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.pwg_info.inv_pwg_info")
}

func init() { proto.RegisterFile("inv_pwg_info.proto", fileDescriptor_aa6fe1f6f982d195) }

var fileDescriptor_aa6fe1f6f982d195 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xbf, 0x4f, 0x3a, 0x41,
	0x10, 0xc5, 0x03, 0xdf, 0xaf, 0x24, 0x6e, 0x6c, 0xd8, 0xf8, 0x63, 0xd5, 0x98, 0x10, 0x2a, 0xaa,
	0x4d, 0x00, 0x29, 0x2c, 0x2c, 0x88, 0x1a, 0x0b, 0x63, 0x62, 0x30, 0x14, 0x54, 0x93, 0xe5, 0x6e,
	0x21, 0x1b, 0x61, 0xe6, 0x32, 0xbb, 0x77, 0x48, 0xeb, 0x5f, 0x6e, 0x6e, 0x91, 0xbb, 0x6b, 0xe6,
	0xde, 0xbd, 0x79, 0x9f, 0x79, 0xc5, 0x0a, 0xe9, 0xb0, 0x80, 0x6c, 0xb7, 0x06, 0x87, 0x2b, 0xd2,
	0x19, 0x53, 0x20, 0xb9, 0x48, 0x9c, 0x4f, 0x08, 0x1c, 0x79, 0xf8, 0x66, 0x70, 0x58, 0x6c, 0xd7,
	0x0c, 0x94, 0x59, 0xd6, 0x0e, 0x0b, 0x8b, 0x81, 0x78, 0xaf, 0xd9, 0x24, 0x5f, 0x3e, 0x4e, 0xbd,
	0x32, 0x18, 0xd8, 0xec, 0xb5, 0xdf, 0x50, 0xd0, 0xc1, 0xbb, 0xa1, 0x8f, 0x53, 0x9b, 0x10, 0xd8,
	0x2d, 0xf3, 0x60, 0xbd, 0x3e, 0x16, 0xf4, 0xe7, 0xa2, 0xdb, 0x2c, 0x84, 0xb7, 0x97, 0xc5, 0xa7,
	0x94, 0xe2, 0x3f, 0x9a, 0xad, 0x55, 0xad, 0x5e, 0x6b, 0x70, 0x3a, 0x8b, 0x5a, 0x5e, 0x88, 0x4e,
	0xf9, 0x85, 0xa1, 0x6a, 0x47, 0xf7, 0xa4, 0xfc, 0x1b, 0x56, 0xf6, 0x48, 0xfd, 0xab, 0xed, 0x51,
	0xff, 0xa7, 0x2d, 0xce, 0x9a, 0x77, 0xe5, 0xa3, 0xb8, 0xcd, 0x68, 0x67, 0x19, 0xd6, 0x4c, 0x79,
	0x06, 0x6c, 0xd3, 0x1c, 0x53, 0x83, 0xc9, 0x1e, 0xb6, 0x94, 0x5a, 0x35, 0xea, 0xb5, 0x06, 0xdd,
	0x99, 0x8a, 0x91, 0xd7, 0x32, 0x31, 0xab, 0x02, 0xef, 0x94, 0x5a, 0x39, 0x11, 0x57, 0x4d, 0xfc,
	0xa0, 0x73, 0x74, 0xc1, 0xab, 0x71, 0xec, 0x3d, 0xaf, 0xd1, 0x8f, 0x52, 0xcd, 0xcb, 0x9d, 0x9c,
	0x8a, 0xbb, 0x26, 0x66, 0x0a, 0xe3, 0x36, 0x66, 0xb9, 0xb1, 0x90, 0xe4, 0xcc, 0x16, 0x83, 0xba,
	0x8f, 0xbd, 0x37, 0x35, 0x3c, 0x3d, 0x46, 0x9e, 0x0e, 0x09, 0xf9, 0x20, 0xae, 0x9b, 0x27, 0x52,
	0x36, 0x3b, 0xac, 0xf0, 0x49, 0xc4, 0x2f, 0x6b, 0xfc, 0xb9, 0x5c, 0xff, 0xa1, 0xcb, 0x4e, 0x7c,
	0xbd, 0xf1, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x82, 0xfd, 0xa4, 0xd3, 0x01, 0x00, 0x00,
}
