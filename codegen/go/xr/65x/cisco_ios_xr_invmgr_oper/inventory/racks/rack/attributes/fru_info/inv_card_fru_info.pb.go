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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_attributes_fru_info

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
	proto.RegisterType((*InvCardFruInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.attributes.fru_info.inv_card_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvCardFruInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.attributes.fru_info.inv_card_fru_info")
}

func init() { proto.RegisterFile("inv_card_fru_info.proto", fileDescriptor_3004b1b322cece9c) }

var fileDescriptor_3004b1b322cece9c = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x6a, 0x14, 0x41,
	0x10, 0xc6, 0x59, 0x51, 0xc1, 0x32, 0x12, 0xb7, 0x89, 0xc9, 0x18, 0x10, 0xc2, 0x1e, 0x24, 0x48,
	0x18, 0x21, 0x89, 0x51, 0x83, 0x97, 0x10, 0x3c, 0x88, 0xc4, 0xc0, 0xac, 0x1e, 0x3c, 0x35, 0x9d,
	0xd9, 0xca, 0xda, 0x98, 0xae, 0x1e, 0xab, 0x6b, 0x56, 0x7d, 0x0d, 0x6f, 0xbe, 0xad, 0x74, 0xf5,
	0xc6, 0x68, 0x76, 0x6f, 0xe2, 0x65, 0x69, 0xea, 0xfb, 0x7d, 0xf5, 0x77, 0x07, 0x36, 0x3c, 0xcd,
	0x6c, 0xeb, 0x78, 0x62, 0xcf, 0xb9, 0xb7, 0x9e, 0xce, 0x63, 0xdd, 0x71, 0x94, 0x68, 0x8e, 0x5a,
	0x9f, 0xda, 0x68, 0x7d, 0x4c, 0xf6, 0x1b, 0x5b, 0x4f, 0xb3, 0x30, 0x65, 0x1b, 0x3b, 0xe4, 0xda,
	0xd3, 0x0c, 0x49, 0x22, 0x7f, 0xaf, 0xd9, 0xb5, 0x9f, 0x93, 0xfe, 0xd6, 0x4e, 0x84, 0xfd, 0x59,
	0x2f, 0x98, 0xea, 0xcb, 0x44, 0xa3, 0x1d, 0x58, 0x5f, 0xc8, 0x6e, 0xdf, 0xbe, 0xfe, 0x38, 0x36,
	0x06, 0x6e, 0x92, 0x0b, 0x58, 0x0d, 0xb6, 0x06, 0xdb, 0x77, 0x1a, 0x7d, 0x8f, 0xa6, 0xb0, 0x92,
	0x69, 0xf1, 0x01, 0x53, 0x87, 0xad, 0x79, 0x0c, 0xab, 0xf9, 0x6d, 0x3d, 0xd9, 0x84, 0x6d, 0xa4,
	0x49, 0x52, 0x7c, 0xd8, 0xdc, 0xcb, 0xe1, 0x37, 0x34, 0x2e, 0x41, 0xf3, 0x14, 0xd6, 0x2e, 0x39,
	0x72, 0x14, 0x7f, 0xc3, 0x37, 0x14, 0x1e, 0x16, 0xf8, 0x9d, 0xa3, 0x38, 0x37, 0x8c, 0x7e, 0xde,
	0x82, 0xe1, 0x42, 0x5f, 0xe6, 0x10, 0x1e, 0x6a, 0xc0, 0x4d, 0x82, 0x27, 0x9f, 0x84, 0x9d, 0xf8,
	0x19, 0xda, 0x24, 0x4e, 0xb0, 0xda, 0xd5, 0x5c, 0x1b, 0x19, 0x38, 0xfa, 0x4b, 0x1f, 0x67, 0xd9,
	0xbc, 0x82, 0xcd, 0x2e, 0x7e, 0x45, 0x5e, 0x6e, 0xde, 0x53, 0x73, 0xa5, 0xc4, 0x32, 0xf7, 0x3e,
	0xac, 0x6b, 0xe5, 0xbc, 0x5c, 0x27, 0x3e, 0x92, 0xbb, 0x98, 0x3b, 0xf7, 0xd5, 0xb9, 0x96, 0xd5,
	0xd3, 0x2b, 0xb1, 0xb8, 0x76, 0xc0, 0xa8, 0x2b, 0x44, 0xf2, 0x12, 0x79, 0xee, 0x78, 0xa6, 0x8e,
	0xfb, 0x59, 0x39, 0x29, 0x42, 0xa1, 0x9f, 0xc0, 0x50, 0x69, 0xc6, 0x84, 0x62, 0x19, 0x5d, 0x8a,
	0x54, 0x1d, 0xe8, 0xf6, 0x57, 0xb3, 0xd0, 0xe4, 0x78, 0xa3, 0x61, 0xf3, 0x63, 0x00, 0x8f, 0x2e,
	0x5c, 0x92, 0xc5, 0x86, 0x6c, 0xfb, 0xc9, 0xd1, 0x14, 0xab, 0xe7, 0x5b, 0x83, 0xed, 0xbb, 0xbb,
	0xa7, 0xf5, 0x3f, 0xff, 0x45, 0xea, 0x3f, 0x2f, 0xde, 0x6c, 0xe6, 0xaa, 0xd7, 0x07, 0x3d, 0xd6,
	0x92, 0xf9, 0x3c, 0x65, 0xc5, 0x6d, 0xcf, 0x8c, 0x24, 0x36, 0xa0, 0x4b, 0x3d, 0x63, 0x40, 0x92,
	0xea, 0x45, 0x39, 0x8f, 0x02, 0xc7, 0x45, 0x3f, 0xb9, 0x92, 0xcd, 0x17, 0x58, 0xd1, 0xe1, 0xfb,
	0x4e, 0x6b, 0x55, 0x2f, 0xff, 0x4f, 0xfb, 0x90, 0x8b, 0x7c, 0xe8, 0xde, 0xfb, 0x80, 0xe6, 0x00,
	0x4a, 0x37, 0x4b, 0x8e, 0x7a, 0xa8, 0xcd, 0x3e, 0x50, 0xf9, 0xfa, 0xb0, 0x67, 0xb7, 0xf5, 0xe3,
	0xdb, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x57, 0xa7, 0xc7, 0x3c, 0x97, 0x03, 0x00, 0x00,
}
