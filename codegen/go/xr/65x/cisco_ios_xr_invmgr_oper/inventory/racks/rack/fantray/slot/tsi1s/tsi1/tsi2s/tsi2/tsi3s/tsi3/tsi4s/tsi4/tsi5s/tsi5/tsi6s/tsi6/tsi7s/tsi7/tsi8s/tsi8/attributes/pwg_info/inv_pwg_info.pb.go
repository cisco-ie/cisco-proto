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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_attributes_pwg_info

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
	Name_4               string   `protobuf:"bytes,5,opt,name=name_4,json=name4,proto3" json:"name_4,omitempty"`
	Name_5               string   `protobuf:"bytes,6,opt,name=name_5,json=name5,proto3" json:"name_5,omitempty"`
	Name_6               string   `protobuf:"bytes,7,opt,name=name_6,json=name6,proto3" json:"name_6,omitempty"`
	Name_7               string   `protobuf:"bytes,8,opt,name=name_7,json=name7,proto3" json:"name_7,omitempty"`
	Name_8               string   `protobuf:"bytes,9,opt,name=name_8,json=name8,proto3" json:"name_8,omitempty"`
	Name_9               string   `protobuf:"bytes,10,opt,name=name_9,json=name9,proto3" json:"name_9,omitempty"`
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

func (m *InvPwgInfo_KEYS) GetName_4() string {
	if m != nil {
		return m.Name_4
	}
	return ""
}

func (m *InvPwgInfo_KEYS) GetName_5() string {
	if m != nil {
		return m.Name_5
	}
	return ""
}

func (m *InvPwgInfo_KEYS) GetName_6() string {
	if m != nil {
		return m.Name_6
	}
	return ""
}

func (m *InvPwgInfo_KEYS) GetName_7() string {
	if m != nil {
		return m.Name_7
	}
	return ""
}

func (m *InvPwgInfo_KEYS) GetName_8() string {
	if m != nil {
		return m.Name_8
	}
	return ""
}

func (m *InvPwgInfo_KEYS) GetName_9() string {
	if m != nil {
		return m.Name_9
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
	proto.RegisterType((*InvPwgInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.attributes.pwg_info.inv_pwg_info_KEYS")
	proto.RegisterType((*InvPwgInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.attributes.pwg_info.inv_pwg_info")
}

func init() { proto.RegisterFile("inv_pwg_info.proto", fileDescriptor_aa6fe1f6f982d195) }

var fileDescriptor_aa6fe1f6f982d195 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xcd, 0xaa, 0x13, 0x31,
	0x18, 0x86, 0x99, 0xe3, 0x39, 0xd5, 0x06, 0x37, 0x0d, 0xfe, 0x44, 0x45, 0x28, 0x5d, 0x75, 0x15,
	0xe8, 0xfc, 0x77, 0xe1, 0xa2, 0xa8, 0xb8, 0x10, 0x41, 0x2a, 0x2e, 0x5c, 0x85, 0x74, 0x26, 0x2d,
	0xc1, 0x36, 0x19, 0x92, 0xcc, 0xd4, 0x6e, 0xbd, 0x18, 0x2f, 0xd1, 0xb5, 0xe4, 0xab, 0x4d, 0x66,
	0xf3, 0xce, 0x33, 0x4f, 0xbe, 0x77, 0x92, 0x81, 0x20, 0x2c, 0xd5, 0xc0, 0xba, 0xf3, 0x81, 0x49,
	0xb5, 0xd7, 0xb4, 0x33, 0xda, 0x69, 0xfc, 0x27, 0x69, 0xa4, 0x6d, 0x34, 0x93, 0xda, 0xb2, 0x5f,
	0x86, 0x49, 0x35, 0x9c, 0x0e, 0x86, 0xe9, 0x4e, 0x18, 0x2a, 0xd5, 0x20, 0x94, 0xd3, 0xe6, 0x42,
	0x0d, 0x6f, 0x7e, 0x5a, 0x48, 0xba, 0xe7, 0xca, 0x19, 0x7e, 0xa1, 0xf6, 0xa8, 0x1d, 0x75, 0x56,
	0xae, 0x2c, 0xa4, 0x8f, 0x14, 0x30, 0xf5, 0x91, 0x01, 0x66, 0x3e, 0x72, 0xc0, 0xdc, 0x47, 0x01,
	0x58, 0xf8, 0x28, 0x01, 0x4b, 0x1f, 0x15, 0x60, 0xe5, 0xa3, 0x06, 0xac, 0x29, 0x77, 0xce, 0xc8,
	0x5d, 0xef, 0x84, 0xa5, 0xb7, 0xe3, 0x2e, 0xfe, 0x26, 0x68, 0x36, 0x3e, 0x3f, 0xfb, 0xfc, 0xf1,
	0xc7, 0x37, 0x8c, 0xd1, 0xbd, 0xe2, 0x27, 0x41, 0x92, 0x79, 0xb2, 0x9c, 0x6e, 0x81, 0xf1, 0x73,
	0x34, 0xf1, 0x4f, 0xb6, 0x22, 0x77, 0x60, 0x1f, 0xfc, 0xdb, 0x2a, 0xe8, 0x94, 0x3c, 0x8a, 0x3a,
	0x0d, 0x3a, 0x23, 0xf7, 0x51, 0x67, 0x41, 0xe7, 0xe4, 0x21, 0xea, 0x3c, 0xe8, 0x82, 0x4c, 0xa2,
	0x2e, 0x82, 0x2e, 0xc9, 0xe3, 0xa8, 0xcb, 0xa0, 0x2b, 0xf2, 0x24, 0xea, 0x2a, 0xe8, 0x9a, 0x4c,
	0xa3, 0xae, 0x83, 0x5e, 0x13, 0x14, 0xf5, 0x7a, 0xf1, 0xfb, 0x0e, 0x3d, 0x1d, 0xff, 0x38, 0x7e,
	0x87, 0xde, 0x74, 0xfa, 0x2c, 0x0c, 0x3b, 0x18, 0xdd, 0x77, 0xcc, 0x88, 0xb6, 0x57, 0x2d, 0x57,
	0xcd, 0x85, 0x9d, 0x74, 0x2b, 0x48, 0x3a, 0x4f, 0x96, 0xb3, 0x2d, 0x81, 0x91, 0x4f, 0x7e, 0x62,
	0x1b, 0x06, 0xbe, 0xe8, 0x56, 0xe0, 0x02, 0xbd, 0x1c, 0xd7, 0xaf, 0xdc, 0x2b, 0xe9, 0x2c, 0xc9,
	0x60, 0xdf, 0x67, 0xb1, 0xfa, 0xd5, 0xd3, 0x77, 0xbf, 0x86, 0x37, 0xe8, 0xed, 0xb8, 0xc6, 0x07,
	0x2e, 0x8f, 0x7c, 0x77, 0x14, 0xac, 0xe9, 0x8d, 0x11, 0xca, 0x91, 0x1c, 0xf6, 0x7d, 0x1d, 0xcb,
	0x9b, 0xdb, 0xc8, 0xfb, 0xeb, 0x04, 0x5e, 0xa3, 0x57, 0xe3, 0x4f, 0xb4, 0x86, 0x9f, 0x55, 0xa8,
	0x17, 0x50, 0x7f, 0x11, 0xeb, 0x1f, 0xfc, 0xf2, 0xff, 0xea, 0x6e, 0x02, 0xb7, 0x35, 0xfb, 0x17,
	0x00, 0x00, 0xff, 0xff, 0xf1, 0x77, 0x26, 0x65, 0xc3, 0x02, 0x00, 0x00,
}
