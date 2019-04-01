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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_tsi10s_tsi10_tsi11s_tsi11_attributes_fru_info

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
	proto.RegisterType((*InvCardFruInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.fru_info.inv_card_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvCardFruInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.fru_info.inv_card_fru_info")
}

func init() { proto.RegisterFile("inv_card_fru_info.proto", fileDescriptor_3004b1b322cece9c) }

var fileDescriptor_3004b1b322cece9c = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x95, 0xcd, 0x6e, 0x13, 0x3f,
	0x14, 0xc5, 0x35, 0xf9, 0xf7, 0xe3, 0x5f, 0x37, 0x55, 0x89, 0xd5, 0x36, 0xa6, 0x12, 0x52, 0x95,
	0x05, 0xaa, 0x10, 0x1a, 0x32, 0x1f, 0x99, 0x24, 0x15, 0x1b, 0x54, 0xb1, 0x40, 0xa8, 0x20, 0x25,
	0xb0, 0x60, 0x65, 0xb9, 0x13, 0x37, 0x58, 0x64, 0xec, 0x91, 0xed, 0x04, 0xfa, 0x2e, 0x3c, 0x03,
	0x8f, 0xc1, 0x82, 0x27, 0xe0, 0x5d, 0xd8, 0x20, 0xdf, 0x49, 0xc7, 0xa1, 0xc9, 0x13, 0x74, 0x73,
	0x72, 0xf3, 0x3b, 0xf7, 0x8c, 0xe3, 0x6b, 0x8f, 0x82, 0xda, 0x42, 0x2e, 0x68, 0xce, 0xf4, 0x84,
	0xde, 0xe8, 0x39, 0x15, 0xf2, 0x46, 0x85, 0xa5, 0x56, 0x56, 0xe1, 0x5f, 0x41, 0x2e, 0x4c, 0xae,
	0xa8, 0x50, 0x86, 0x7e, 0xd3, 0x54, 0xc8, 0x45, 0x31, 0xd5, 0x54, 0x95, 0x5c, 0x87, 0x42, 0x2e,
	0xb8, 0xb4, 0x4a, 0xdf, 0x86, 0x9a, 0xe5, 0x5f, 0x0c, 0x68, 0xc8, 0xa5, 0x15, 0xf6, 0x36, 0x34,
	0x33, 0x65, 0x43, 0x6b, 0x44, 0x64, 0x40, 0x9d, 0xc4, 0x50, 0xc6, 0x4e, 0x12, 0x28, 0x13, 0x27,
	0x29, 0x94, 0xa9, 0x93, 0x1e, 0x94, 0x3d, 0x27, 0x19, 0x94, 0x99, 0x93, 0x3e, 0x94, 0x7d, 0x27,
	0x03, 0x28, 0x07, 0x4e, 0x86, 0x50, 0x0e, 0xe1, 0xe1, 0xdd, 0x6a, 0x8d, 0x2e, 0xe8, 0x72, 0xc1,
	0x28, 0x64, 0xd6, 0x6a, 0x71, 0x3d, 0xb7, 0xdc, 0x84, 0x77, 0x5b, 0xea, 0xfc, 0x6c, 0xa0, 0x93,
	0xb5, 0x8d, 0xd2, 0xb7, 0xaf, 0x3f, 0x8d, 0x31, 0x46, 0x5b, 0x92, 0x15, 0x9c, 0x04, 0x67, 0xc1,
	0xf9, 0xde, 0x08, 0x6a, 0x7c, 0x8c, 0x76, 0xdc, 0x27, 0x8d, 0x48, 0x03, 0xe8, 0xb6, 0xfb, 0x16,
	0xd5, 0x38, 0x26, 0xff, 0x79, 0x1c, 0xd7, 0x38, 0x21, 0x5b, 0x1e, 0x27, 0x35, 0x4e, 0xc9, 0xb6,
	0xc7, 0x69, 0x8d, 0x7b, 0x64, 0xc7, 0xe3, 0x5e, 0x8d, 0x33, 0xb2, 0xeb, 0x71, 0x56, 0xe3, 0x3e,
	0xf9, 0xdf, 0xe3, 0x7e, 0x8d, 0x07, 0x64, 0xcf, 0xe3, 0x41, 0x8d, 0x87, 0x04, 0x79, 0x3c, 0xc4,
	0x6d, 0xb4, 0x5b, 0x6d, 0xa7, 0x4b, 0xf6, 0x81, 0x43, 0x57, 0xd4, 0xf5, 0x46, 0x44, 0x9a, 0x2b,
	0x46, 0xe4, 0x8d, 0x98, 0x1c, 0xac, 0x18, 0x71, 0x67, 0x8a, 0x9a, 0x6e, 0x8e, 0x56, 0x14, 0xdc,
	0x94, 0x3c, 0xc7, 0x4f, 0xd1, 0xa1, 0xab, 0xa9, 0x90, 0xd4, 0xf0, 0x5c, 0xc9, 0x89, 0x81, 0x41,
	0xb6, 0x46, 0x07, 0x0e, 0xbf, 0x91, 0xe3, 0x0a, 0xe2, 0x17, 0xe8, 0xe8, 0xae, 0x4f, 0x32, 0xa9,
	0xea, 0xe6, 0x06, 0x34, 0xb7, 0xaa, 0xe6, 0x77, 0x4c, 0xaa, 0x65, 0xa0, 0xf3, 0x63, 0x17, 0xb5,
	0xd6, 0x4e, 0x0c, 0x5f, 0xa0, 0xc7, 0x00, 0xd8, 0xa4, 0x10, 0x52, 0x18, 0xab, 0x99, 0x15, 0x0b,
	0x4e, 0x8d, 0x65, 0x96, 0x93, 0x18, 0x9e, 0xd5, 0x76, 0x0d, 0xaf, 0xfe, 0xf1, 0xc7, 0xce, 0xc6,
	0x2f, 0xd1, 0x69, 0xa9, 0xbe, 0x72, 0xbd, 0x39, 0x9c, 0x40, 0x98, 0x40, 0xc7, 0xa6, 0x74, 0x8a,
	0x4e, 0x60, 0x65, 0xf7, 0x02, 0x30, 0x2b, 0x94, 0x64, 0xb3, 0x65, 0x32, 0x85, 0xe4, 0x91, 0x73,
	0xdf, 0x7b, 0xb3, 0x4a, 0x3d, 0x47, 0x18, 0x52, 0x85, 0x92, 0xc2, 0x2a, 0xbd, 0x4c, 0xf4, 0x20,
	0xf1, 0xc8, 0x39, 0x57, 0x95, 0x51, 0x75, 0x3f, 0x43, 0x2d, 0xe8, 0xd6, 0xdc, 0x70, 0x4b, 0x35,
	0x67, 0x46, 0x49, 0x92, 0xc1, 0xfc, 0x0f, 0x9d, 0x31, 0x72, 0x7c, 0x04, 0x18, 0x7f, 0x6f, 0xa0,
	0x27, 0x33, 0x66, 0xec, 0xfa, 0x0f, 0xa2, 0xf9, 0x67, 0x26, 0xa7, 0x9c, 0xf4, 0xcf, 0x82, 0xf3,
	0xfd, 0xf8, 0x77, 0x10, 0x3e, 0x9c, 0xf7, 0x38, 0x5c, 0xbd, 0x7b, 0xa3, 0x53, 0xb7, 0xff, 0xfb,
	0x23, 0xbf, 0x84, 0xcd, 0xbb, 0x8b, 0x52, 0x1d, 0x76, 0x3e, 0xd7, 0x9a, 0x4b, 0x4b, 0x0b, 0xce,
	0xcc, 0x5c, 0xf3, 0x82, 0x4b, 0x4b, 0x06, 0xd5, 0x45, 0x81, 0x86, 0xcb, 0xca, 0xbf, 0xf2, 0x36,
	0xfe, 0x13, 0xa0, 0x26, 0x9c, 0xc3, 0xbc, 0x84, 0xc5, 0xc8, 0xf0, 0xa1, 0x4f, 0x12, 0xb9, 0xed,
	0x7e, 0x2c, 0x3f, 0x88, 0x82, 0xe3, 0x0c, 0x55, 0x83, 0xd9, 0x70, 0xd3, 0x2f, 0x60, 0x6e, 0xc7,
	0x60, 0xdf, 0x9f, 0xfb, 0xf5, 0x0e, 0xfc, 0x6d, 0x24, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4b,
	0xe4, 0x46, 0x5f, 0x51, 0x06, 0x00, 0x00,
}
