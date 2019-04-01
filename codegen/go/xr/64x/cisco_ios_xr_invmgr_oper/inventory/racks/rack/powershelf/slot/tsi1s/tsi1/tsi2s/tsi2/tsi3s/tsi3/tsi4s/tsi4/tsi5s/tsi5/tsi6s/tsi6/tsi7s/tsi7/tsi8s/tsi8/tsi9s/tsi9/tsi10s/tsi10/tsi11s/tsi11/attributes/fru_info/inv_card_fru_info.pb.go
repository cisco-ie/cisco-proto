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
// source: inv_card_fru_info.proto

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_tsi10s_tsi10_tsi11s_tsi11_attributes_fru_info

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

type InvCardFruInfo_KEYS struct {
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
	Name_10              string   `protobuf:"bytes,11,opt,name=name_10,json=name10,proto3" json:"name_10,omitempty"`
	Name_11              string   `protobuf:"bytes,12,opt,name=name_11,json=name11,proto3" json:"name_11,omitempty"`
	Name_12              string   `protobuf:"bytes,13,opt,name=name_12,json=name12,proto3" json:"name_12,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvCardFruInfo_KEYS) Reset()         { *m = InvCardFruInfo_KEYS{} }
func (m *InvCardFruInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*InvCardFruInfo_KEYS) ProtoMessage()    {}
func (*InvCardFruInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3004b1b322cece9c, []int{0}
}

func (m *InvCardFruInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvCardFruInfo_KEYS.Unmarshal(m, b)
}
func (m *InvCardFruInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvCardFruInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *InvCardFruInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvCardFruInfo_KEYS.Merge(m, src)
}
func (m *InvCardFruInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_InvCardFruInfo_KEYS.Size(m)
}
func (m *InvCardFruInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InvCardFruInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InvCardFruInfo_KEYS proto.InternalMessageInfo

func (m *InvCardFruInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_4() string {
	if m != nil {
		return m.Name_4
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_5() string {
	if m != nil {
		return m.Name_5
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_6() string {
	if m != nil {
		return m.Name_6
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_7() string {
	if m != nil {
		return m.Name_7
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_8() string {
	if m != nil {
		return m.Name_8
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_9() string {
	if m != nil {
		return m.Name_9
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_10() string {
	if m != nil {
		return m.Name_10
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_11() string {
	if m != nil {
		return m.Name_11
	}
	return ""
}

func (m *InvCardFruInfo_KEYS) GetName_12() string {
	if m != nil {
		return m.Name_12
	}
	return ""
}

type InvTimespec struct {
	TimeInSeconds        int32    `protobuf:"zigzag32,1,opt,name=time_in_seconds,json=timeInSeconds,proto3" json:"time_in_seconds,omitempty"`
	TimeInNanoSeconds    int32    `protobuf:"zigzag32,2,opt,name=time_in_nano_seconds,json=timeInNanoSeconds,proto3" json:"time_in_nano_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvTimespec) Reset()         { *m = InvTimespec{} }
func (m *InvTimespec) String() string { return proto.CompactTextString(m) }
func (*InvTimespec) ProtoMessage()    {}
func (*InvTimespec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3004b1b322cece9c, []int{1}
}

func (m *InvTimespec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvTimespec.Unmarshal(m, b)
}
func (m *InvTimespec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvTimespec.Marshal(b, m, deterministic)
}
func (m *InvTimespec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvTimespec.Merge(m, src)
}
func (m *InvTimespec) XXX_Size() int {
	return xxx_messageInfo_InvTimespec.Size(m)
}
func (m *InvTimespec) XXX_DiscardUnknown() {
	xxx_messageInfo_InvTimespec.DiscardUnknown(m)
}

var xxx_messageInfo_InvTimespec proto.InternalMessageInfo

func (m *InvTimespec) GetTimeInSeconds() int32 {
	if m != nil {
		return m.TimeInSeconds
	}
	return 0
}

func (m *InvTimespec) GetTimeInNanoSeconds() int32 {
	if m != nil {
		return m.TimeInNanoSeconds
	}
	return 0
}

type InvCardFruInfo struct {
	CardAdministrativeState    int32        `protobuf:"zigzag32,50,opt,name=card_administrative_state,json=cardAdministrativeState,proto3" json:"card_administrative_state,omitempty"`
	PowerAdministrativeState   int32        `protobuf:"zigzag32,51,opt,name=power_administrative_state,json=powerAdministrativeState,proto3" json:"power_administrative_state,omitempty"`
	CardOperationalState       int32        `protobuf:"zigzag32,52,opt,name=card_operational_state,json=cardOperationalState,proto3" json:"card_operational_state,omitempty"`
	CardMonitorState           int32        `protobuf:"zigzag32,53,opt,name=card_monitor_state,json=cardMonitorState,proto3" json:"card_monitor_state,omitempty"`
	CardResetReason            string       `protobuf:"bytes,54,opt,name=card_reset_reason,json=cardResetReason,proto3" json:"card_reset_reason,omitempty"`
	LastOperationalStateChange *InvTimespec `protobuf:"bytes,55,opt,name=last_operational_state_change,json=lastOperationalStateChange,proto3" json:"last_operational_state_change,omitempty"`
	PowerCurrentMeasurement    int32        `protobuf:"zigzag32,56,opt,name=power_current_measurement,json=powerCurrentMeasurement,proto3" json:"power_current_measurement,omitempty"`
	CardUpTime                 *InvTimespec `protobuf:"bytes,57,opt,name=card_up_time,json=cardUpTime,proto3" json:"card_up_time,omitempty"`
	PowerOperationalState      int32        `protobuf:"zigzag32,58,opt,name=power_operational_state,json=powerOperationalState,proto3" json:"power_operational_state,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}     `json:"-"`
	XXX_unrecognized           []byte       `json:"-"`
	XXX_sizecache              int32        `json:"-"`
}

func (m *InvCardFruInfo) Reset()         { *m = InvCardFruInfo{} }
func (m *InvCardFruInfo) String() string { return proto.CompactTextString(m) }
func (*InvCardFruInfo) ProtoMessage()    {}
func (*InvCardFruInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3004b1b322cece9c, []int{2}
}

func (m *InvCardFruInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvCardFruInfo.Unmarshal(m, b)
}
func (m *InvCardFruInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvCardFruInfo.Marshal(b, m, deterministic)
}
func (m *InvCardFruInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvCardFruInfo.Merge(m, src)
}
func (m *InvCardFruInfo) XXX_Size() int {
	return xxx_messageInfo_InvCardFruInfo.Size(m)
}
func (m *InvCardFruInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InvCardFruInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InvCardFruInfo proto.InternalMessageInfo

func (m *InvCardFruInfo) GetCardAdministrativeState() int32 {
	if m != nil {
		return m.CardAdministrativeState
	}
	return 0
}

func (m *InvCardFruInfo) GetPowerAdministrativeState() int32 {
	if m != nil {
		return m.PowerAdministrativeState
	}
	return 0
}

func (m *InvCardFruInfo) GetCardOperationalState() int32 {
	if m != nil {
		return m.CardOperationalState
	}
	return 0
}

func (m *InvCardFruInfo) GetCardMonitorState() int32 {
	if m != nil {
		return m.CardMonitorState
	}
	return 0
}

func (m *InvCardFruInfo) GetCardResetReason() string {
	if m != nil {
		return m.CardResetReason
	}
	return ""
}

func (m *InvCardFruInfo) GetLastOperationalStateChange() *InvTimespec {
	if m != nil {
		return m.LastOperationalStateChange
	}
	return nil
}

func (m *InvCardFruInfo) GetPowerCurrentMeasurement() int32 {
	if m != nil {
		return m.PowerCurrentMeasurement
	}
	return 0
}

func (m *InvCardFruInfo) GetCardUpTime() *InvTimespec {
	if m != nil {
		return m.CardUpTime
	}
	return nil
}

func (m *InvCardFruInfo) GetPowerOperationalState() int32 {
	if m != nil {
		return m.PowerOperationalState
	}
	return 0
}

func init() {
	proto.RegisterType((*InvCardFruInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.fru_info.inv_card_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvCardFruInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.fru_info.inv_card_fru_info")
}

func init() { proto.RegisterFile("inv_card_fru_info.proto", fileDescriptor_3004b1b322cece9c) }

var fileDescriptor_3004b1b322cece9c = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x95, 0xcd, 0x6e, 0x13, 0x3f,
	0x14, 0xc5, 0x35, 0xf9, 0xf7, 0xe3, 0x5f, 0xb7, 0x55, 0x89, 0xd5, 0x36, 0xa6, 0x12, 0x52, 0x95,
	0x05, 0xaa, 0x10, 0x1a, 0x32, 0x1f, 0x99, 0x24, 0x15, 0x1b, 0x54, 0xb1, 0x40, 0xa8, 0x20, 0x25,
	0xb0, 0x60, 0x65, 0xb9, 0x13, 0x27, 0xb5, 0xc8, 0xd8, 0x23, 0xdb, 0x09, 0xf0, 0x36, 0x3c, 0x05,
	0x8f, 0xc0, 0x92, 0x27, 0xe0, 0x5d, 0x40, 0xbe, 0x93, 0x8e, 0x43, 0x93, 0x27, 0x80, 0xcd, 0xc9,
	0xcd, 0xef, 0xdc, 0x3b, 0x8e, 0x8f, 0x3d, 0x0a, 0x6a, 0x09, 0xb9, 0xa0, 0x39, 0xd3, 0x63, 0x3a,
	0xd1, 0x73, 0x2a, 0xe4, 0x44, 0x85, 0xa5, 0x56, 0x56, 0xe1, 0x1f, 0x41, 0x2e, 0x4c, 0xae, 0xa8,
	0x50, 0x86, 0x7e, 0xd6, 0x54, 0xc8, 0x45, 0x31, 0xd5, 0x54, 0x95, 0x5c, 0x87, 0x42, 0x2e, 0xb8,
	0xb4, 0x4a, 0x7f, 0x09, 0x35, 0xcb, 0x3f, 0x1a, 0xd0, 0xb0, 0x54, 0x9f, 0xb8, 0x36, 0xb7, 0x7c,
	0x36, 0x09, 0xcd, 0x4c, 0xd9, 0xd0, 0x1a, 0x11, 0x19, 0x50, 0x27, 0x31, 0x94, 0xb1, 0x93, 0x04,
	0xca, 0xc4, 0x49, 0x0a, 0x65, 0xea, 0xa4, 0x0b, 0x65, 0xd7, 0x49, 0x06, 0x65, 0xe6, 0xa4, 0x07,
	0x65, 0xcf, 0x49, 0x1f, 0xca, 0xbe, 0x93, 0x01, 0x94, 0x03, 0x78, 0x78, 0xa7, 0x5a, 0xa3, 0x03,
	0xba, 0x5c, 0x30, 0x0a, 0x99, 0xb5, 0x5a, 0xdc, 0xcc, 0x2d, 0x37, 0xe1, 0xdd, 0xb6, 0xda, 0xdf,
	0x1b, 0xe8, 0x74, 0x6d, 0xb3, 0xf4, 0xf5, 0xcb, 0x0f, 0x23, 0x8c, 0xd1, 0x96, 0x64, 0x05, 0x27,
	0xc1, 0x79, 0x70, 0xb1, 0x37, 0x84, 0x1a, 0x9f, 0xa0, 0x1d, 0xf7, 0x49, 0x23, 0xd2, 0x00, 0xba,
	0xed, 0xbe, 0x45, 0x35, 0x8e, 0xc9, 0x7f, 0x1e, 0xc7, 0x35, 0x4e, 0xc8, 0x96, 0xc7, 0x49, 0x8d,
	0x53, 0xb2, 0xed, 0x71, 0x5a, 0xe3, 0x2e, 0xd9, 0xf1, 0xb8, 0x5b, 0xe3, 0x8c, 0xec, 0x7a, 0x9c,
	0xd5, 0xb8, 0x47, 0xfe, 0xf7, 0xb8, 0x57, 0xe3, 0x3e, 0xd9, 0xf3, 0xb8, 0x5f, 0xe3, 0x01, 0x41,
	0x1e, 0x0f, 0x70, 0x0b, 0xed, 0x56, 0xdb, 0xe9, 0x90, 0x7d, 0xe0, 0xd0, 0x15, 0x75, 0xbc, 0x11,
	0x91, 0x83, 0x15, 0x23, 0xf2, 0x46, 0x4c, 0x0e, 0x57, 0x8c, 0xb8, 0x3d, 0x45, 0x07, 0x2e, 0x47,
	0x2b, 0x0a, 0x6e, 0x4a, 0x9e, 0xe3, 0xc7, 0xe8, 0xc8, 0xd5, 0x54, 0x48, 0x6a, 0x78, 0xae, 0xe4,
	0xd8, 0x40, 0x90, 0xcd, 0xe1, 0xa1, 0xc3, 0xaf, 0xe4, 0xa8, 0x82, 0xf8, 0x19, 0x3a, 0xbe, 0xeb,
	0x93, 0x4c, 0xaa, 0xba, 0xb9, 0x01, 0xcd, 0xcd, 0xaa, 0xf9, 0x0d, 0x93, 0x6a, 0x39, 0xd0, 0xfe,
	0xb6, 0x8b, 0x9a, 0x6b, 0x27, 0x86, 0x2f, 0xd1, 0x43, 0x00, 0x6c, 0x5c, 0x08, 0x29, 0x8c, 0xd5,
	0xcc, 0x8a, 0x05, 0xa7, 0xc6, 0x32, 0xcb, 0x49, 0x0c, 0xcf, 0x6a, 0xb9, 0x86, 0x17, 0x7f, 0xf8,
	0x23, 0x67, 0xe3, 0xe7, 0xe8, 0x0c, 0xee, 0xe9, 0xe6, 0xe1, 0x04, 0x86, 0x09, 0x74, 0x6c, 0x9a,
	0x4e, 0xd1, 0x29, 0xac, 0xec, 0x5e, 0x02, 0x66, 0x85, 0x92, 0x6c, 0xb6, 0x9c, 0x4c, 0x61, 0xf2,
	0xd8, 0xb9, 0x6f, 0xbd, 0x59, 0x4d, 0x3d, 0x45, 0x18, 0xa6, 0x0a, 0x25, 0x85, 0x55, 0x7a, 0x39,
	0xd1, 0x85, 0x89, 0x07, 0xce, 0xb9, 0xae, 0x8c, 0xaa, 0xfb, 0x09, 0x6a, 0x42, 0xb7, 0xe6, 0x86,
	0x5b, 0xaa, 0x39, 0x33, 0x4a, 0x92, 0x0c, 0xf2, 0x3f, 0x72, 0xc6, 0xd0, 0xf1, 0x21, 0x60, 0xfc,
	0xb5, 0x81, 0x1e, 0xcd, 0x98, 0xb1, 0xeb, 0x3f, 0x88, 0xe6, 0xb7, 0x4c, 0x4e, 0x39, 0xe9, 0x9d,
	0x07, 0x17, 0xfb, 0xf1, 0xcf, 0x20, 0xfc, 0xbb, 0xde, 0xe5, 0x70, 0xf5, 0xfe, 0x0d, 0xcf, 0x5c,
	0x06, 0xf7, 0x63, 0xbf, 0x82, 0x00, 0xdc, 0x65, 0xa9, 0x0e, 0x3c, 0x9f, 0x6b, 0xcd, 0xa5, 0xa5,
	0x05, 0x67, 0x66, 0xae, 0x79, 0xc1, 0xa5, 0x25, 0xfd, 0xea, 0xb2, 0x40, 0xc3, 0x55, 0xe5, 0x5f,
	0x7b, 0x1b, 0xff, 0x0a, 0xd0, 0x01, 0x9c, 0xc5, 0xbc, 0x84, 0xc5, 0xc8, 0xe0, 0x5f, 0x48, 0x13,
	0xb9, 0x2d, 0xbf, 0x2f, 0xdf, 0x89, 0x82, 0xe3, 0x0c, 0x55, 0xe1, 0x6c, 0xb8, 0xf1, 0x97, 0x90,
	0xdd, 0x09, 0xd8, 0xf7, 0xb3, 0xbf, 0xd9, 0x81, 0xbf, 0x90, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x91, 0x3f, 0xea, 0xb5, 0x5d, 0x06, 0x00, 0x00,
}
