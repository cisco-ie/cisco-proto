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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_attributes_fru_info

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
	proto.RegisterType((*InvCardFruInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.attributes.fru_info.inv_card_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvCardFruInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.attributes.fru_info.inv_card_fru_info")
}

func init() { proto.RegisterFile("inv_card_fru_info.proto", fileDescriptor_3004b1b322cece9c) }

var fileDescriptor_3004b1b322cece9c = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x4b, 0x6b, 0x14, 0x41,
	0x10, 0x66, 0x83, 0x09, 0xd8, 0x46, 0xe2, 0x36, 0x79, 0x8c, 0x01, 0x21, 0xec, 0x41, 0x82, 0x48,
	0x4b, 0x1e, 0x46, 0x0d, 0x5e, 0x64, 0xf1, 0x20, 0x92, 0x08, 0xb3, 0x7a, 0x10, 0x84, 0xa6, 0x33,
	0x5b, 0xbb, 0x36, 0xee, 0x54, 0x0f, 0xd5, 0x35, 0xa3, 0xb9, 0xf8, 0x57, 0xfc, 0x11, 0xfe, 0x41,
	0xe9, 0xea, 0x8d, 0xd1, 0xec, 0x1e, 0xc5, 0xcb, 0x4c, 0xcf, 0xf7, 0xa8, 0xe7, 0xb4, 0xda, 0xf1,
	0xd8, 0xd9, 0xca, 0xd1, 0xd8, 0x4e, 0xa8, 0xb5, 0x1e, 0x27, 0xc1, 0x34, 0x14, 0x38, 0xe8, 0xf3,
	0xca, 0xc7, 0x2a, 0x58, 0x1f, 0xa2, 0xfd, 0x46, 0xd6, 0x63, 0x57, 0x4f, 0xc9, 0x86, 0x06, 0xc8,
	0x78, 0xec, 0x00, 0x39, 0xd0, 0xa5, 0x21, 0x57, 0x7d, 0x89, 0xf2, 0x34, 0x13, 0x87, 0x4c, 0xee,
	0xd2, 0xc4, 0x59, 0x60, 0xe3, 0x98, 0xc9, 0x5f, 0xb4, 0x0c, 0xd1, 0x5c, 0x45, 0x1d, 0x0c, 0xd5,
	0xf6, 0x42, 0x2a, 0xfb, 0xf6, 0xf5, 0xc7, 0x91, 0xd6, 0xea, 0x16, 0xba, 0x1a, 0x8a, 0xde, 0x5e,
	0x6f, 0xff, 0x76, 0x29, 0x67, 0xbd, 0xa5, 0xd6, 0xd2, 0xdb, 0x1e, 0x14, 0x2b, 0x82, 0xae, 0xa6,
	0xaf, 0x83, 0xc1, 0x54, 0xad, 0xa7, 0x20, 0xec, 0x6b, 0x88, 0x0d, 0x54, 0xfa, 0xa1, 0xda, 0x48,
	0x67, 0xeb, 0xd1, 0x46, 0xa8, 0x02, 0x8e, 0xa3, 0x44, 0xe9, 0x97, 0x77, 0x13, 0xfc, 0x06, 0x47,
	0x19, 0xd4, 0x4f, 0xd4, 0xe6, 0x95, 0x0e, 0x1d, 0x86, 0xdf, 0xe2, 0x15, 0x11, 0xf7, 0xb3, 0xf8,
	0xdc, 0x61, 0x98, 0x1b, 0x06, 0x3f, 0x57, 0x55, 0x7f, 0xa1, 0x5c, 0x7d, 0xaa, 0xee, 0x0b, 0xe0,
	0xc6, 0xb5, 0x47, 0x1f, 0x99, 0x1c, 0xfb, 0x0e, 0x6c, 0x64, 0xc7, 0x50, 0x1c, 0x4a, 0xac, 0x9d,
	0x24, 0x78, 0xf5, 0x17, 0x3f, 0x4a, 0xb4, 0x7e, 0xa9, 0x76, 0x9b, 0xf0, 0x15, 0x68, 0xb9, 0xf9,
	0x48, 0xcc, 0x85, 0x28, 0x96, 0xb9, 0x8f, 0xd5, 0xb6, 0x64, 0x4e, 0x0b, 0x70, 0xec, 0x03, 0xba,
	0xd9, 0xdc, 0x79, 0x2c, 0xce, 0xcd, 0xc4, 0xbe, 0xbb, 0x26, 0xb3, 0xeb, 0xb1, 0xd2, 0xe2, 0xaa,
	0x03, 0x7a, 0x0e, 0x34, 0x77, 0x3c, 0x15, 0xc7, 0xbd, 0xc4, 0x9c, 0x65, 0x22, 0xab, 0x1f, 0xa9,
	0xbe, 0xa8, 0x09, 0x22, 0xb0, 0x25, 0x70, 0x31, 0x60, 0x71, 0x22, 0xe3, 0xdf, 0x48, 0x44, 0x99,
	0xf0, 0x52, 0x60, 0xfd, 0xa3, 0xa7, 0x1e, 0xcc, 0x5c, 0xe4, 0xc5, 0x82, 0x6c, 0xf5, 0xd9, 0xe1,
	0x14, 0x8a, 0x67, 0x7b, 0xbd, 0xfd, 0x3b, 0x87, 0x9f, 0xcc, 0xbf, 0xfd, 0x8d, 0xcc, 0x9f, 0xeb,
	0x2f, 0x77, 0x53, 0x09, 0x37, 0xbb, 0x1e, 0x4a, 0xfe, 0xb4, 0xab, 0x3c, 0xef, 0xaa, 0x25, 0x02,
	0x64, 0x5b, 0x83, 0x8b, 0x2d, 0x41, 0x0d, 0xc8, 0xc5, 0xf3, 0xbc, 0x2b, 0x11, 0x0c, 0x33, 0x7f,
	0x76, 0x4d, 0xeb, 0xef, 0x6a, 0x5d, 0x26, 0xd1, 0x36, 0x92, 0xab, 0x78, 0xf1, 0x1f, 0x7a, 0x51,
	0x29, 0xe3, 0x87, 0xe6, 0xbd, 0xaf, 0x41, 0x9f, 0xa8, 0x5c, 0xda, 0x92, 0x75, 0x9f, 0x4a, 0xe5,
	0x5b, 0x42, 0xdf, 0xec, 0xfc, 0x62, 0x4d, 0xae, 0xee, 0xd1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x56, 0x6e, 0x72, 0xb6, 0xd5, 0x03, 0x00, 0x00,
}
