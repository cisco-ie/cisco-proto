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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_attributes_fru_info

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
	proto.RegisterType((*InvCardFruInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.fru_info.inv_card_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvCardFruInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.fru_info.inv_card_fru_info")
}

func init() { proto.RegisterFile("inv_card_fru_info.proto", fileDescriptor_3004b1b322cece9c) }

var fileDescriptor_3004b1b322cece9c = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x66, 0xab, 0x2d, 0x38, 0x56, 0x6a, 0x86, 0xfe, 0xac, 0x05, 0xa1, 0xe4, 0x42, 0x8a, 0xc8,
	0x48, 0xd2, 0x5a, 0xb5, 0x78, 0x23, 0xc5, 0x0b, 0x91, 0x2a, 0x6c, 0xf4, 0xa2, 0xde, 0x0c, 0xd3,
	0xcd, 0x49, 0x1c, 0xcc, 0x9e, 0x59, 0xce, 0x9c, 0x8d, 0xf6, 0x11, 0x7c, 0x15, 0x1f, 0xc5, 0xa7,
	0x92, 0x39, 0x9b, 0x36, 0x9a, 0xe4, 0x56, 0x6f, 0x26, 0x93, 0xef, 0x67, 0xce, 0xef, 0xaa, 0x3d,
	0x8f, 0x53, 0x5b, 0x3a, 0x1a, 0xda, 0x11, 0x35, 0xd6, 0xe3, 0x28, 0x98, 0x9a, 0x02, 0x07, 0x7d,
	0x51, 0xfa, 0x58, 0x06, 0xeb, 0x43, 0xb4, 0xdf, 0xc9, 0x7a, 0x9c, 0x56, 0x63, 0xb2, 0xa1, 0x06,
	0x32, 0x1e, 0xa7, 0x80, 0x1c, 0xe8, 0xca, 0x90, 0x2b, 0xbf, 0x46, 0x39, 0xcd, 0xc8, 0x21, 0x93,
	0xbb, 0x32, 0x71, 0x12, 0xd8, 0x70, 0xf4, 0xbd, 0x28, 0xa7, 0x71, 0xcc, 0xe4, 0x2f, 0x1b, 0x86,
	0x68, 0xae, 0x03, 0x74, 0x3f, 0xab, 0xdd, 0xa5, 0xa8, 0xf6, 0xdd, 0x9b, 0x8b, 0x81, 0xd6, 0xea,
	0x36, 0xba, 0x0a, 0xf2, 0xec, 0x20, 0x3b, 0xbc, 0x53, 0xc8, 0x5d, 0xef, 0xa8, 0x8d, 0xf4, 0x6b,
	0x7b, 0xf9, 0x9a, 0xa0, 0xeb, 0xe9, 0x5f, 0xef, 0x06, 0xee, 0xe7, 0xb7, 0xe6, 0x70, 0xbf, 0x3b,
	0x56, 0x9b, 0xe9, 0x6d, 0xf6, 0x15, 0xc4, 0x1a, 0x4a, 0xfd, 0x48, 0x6d, 0xa5, 0xbb, 0xf5, 0x68,
	0x23, 0x94, 0x01, 0x87, 0x51, 0x1e, 0xef, 0x14, 0xf7, 0x12, 0xfc, 0x16, 0x07, 0x2d, 0xa8, 0x9f,
	0xaa, 0xed, 0x6b, 0x1d, 0x3a, 0x0c, 0x37, 0xe2, 0x35, 0x11, 0x77, 0x5a, 0xf1, 0x7b, 0x87, 0x61,
	0x66, 0xe8, 0xfe, 0x5a, 0x57, 0x9d, 0xa5, 0x2a, 0xf4, 0xa9, 0x7a, 0x20, 0x80, 0x1b, 0x56, 0x1e,
	0x7d, 0x64, 0x72, 0xec, 0xa7, 0x60, 0x23, 0x3b, 0x86, 0xbc, 0x2f, 0x6f, 0xed, 0x25, 0xc1, 0xeb,
	0xbf, 0xf8, 0x41, 0xa2, 0xf5, 0x2b, 0xb5, 0x5f, 0x87, 0x6f, 0x40, 0xab, 0xcd, 0x47, 0x62, 0xce,
	0x45, 0xb1, 0xca, 0x7d, 0xac, 0x76, 0x25, 0x72, 0x1a, 0x91, 0x63, 0x1f, 0xd0, 0x4d, 0x66, 0xce,
	0x63, 0x71, 0x6e, 0x27, 0xf6, 0xc3, 0x9c, 0x6c, 0x5d, 0x4f, 0x94, 0x16, 0x57, 0x15, 0xd0, 0x73,
	0xa0, 0x99, 0xe3, 0x99, 0x38, 0xee, 0x27, 0xe6, 0xbc, 0x25, 0x5a, 0xf5, 0x63, 0xd5, 0x11, 0x35,
	0x41, 0x04, 0xb6, 0x04, 0x2e, 0x06, 0xcc, 0x4f, 0xa4, 0xfd, 0x5b, 0x89, 0x28, 0x12, 0x5e, 0x08,
	0xac, 0x7f, 0x66, 0xea, 0xe1, 0xc4, 0x45, 0x5e, 0x4e, 0xc8, 0x96, 0x5f, 0x1c, 0x8e, 0x21, 0x7f,
	0x7e, 0x90, 0x1d, 0xde, 0xed, 0x8f, 0xcd, 0x3f, 0x5b, 0x34, 0xf3, 0xe7, 0x26, 0x14, 0xfb, 0x29,
	0x9b, 0xc5, 0x06, 0x9c, 0x49, 0x2a, 0x69, 0x6c, 0x6d, 0xeb, 0xcb, 0x86, 0x08, 0x90, 0x6d, 0x05,
	0x2e, 0x36, 0x04, 0x15, 0x20, 0xe7, 0x2f, 0xda, 0xb1, 0x89, 0xe0, 0xac, 0xe5, 0xcf, 0xe7, 0xb4,
	0xfe, 0x91, 0xa9, 0x4d, 0xe9, 0x4a, 0x53, 0x4b, 0xb0, 0xfc, 0xe5, 0xff, 0xad, 0x4b, 0xa5, 0xe0,
	0x9f, 0xea, 0x8f, 0xbe, 0x02, 0x7d, 0xa2, 0xda, 0x34, 0x57, 0x6c, 0xc1, 0xa9, 0x54, 0xb1, 0x23,
	0xf4, 0x62, 0x17, 0x2e, 0x37, 0xe4, 0x9b, 0x3f, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x62, 0xc1,
	0x04, 0x01, 0x0e, 0x04, 0x00, 0x00,
}
