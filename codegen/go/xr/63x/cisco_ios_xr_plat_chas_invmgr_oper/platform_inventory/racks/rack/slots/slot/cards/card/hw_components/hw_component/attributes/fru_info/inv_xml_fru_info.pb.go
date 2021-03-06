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
// source: inv_xml_fru_info.proto

package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_slots_slot_cards_card_hw_components_hw_component_attributes_fru_info

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

type InvXmlFruInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Name_1               string   `protobuf:"bytes,2,opt,name=name_1,json=name1,proto3" json:"name_1,omitempty"`
	Name_2               string   `protobuf:"bytes,3,opt,name=name_2,json=name2,proto3" json:"name_2,omitempty"`
	Name_3               string   `protobuf:"bytes,4,opt,name=name_3,json=name3,proto3" json:"name_3,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvXmlFruInfo_KEYS) Reset()         { *m = InvXmlFruInfo_KEYS{} }
func (m *InvXmlFruInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*InvXmlFruInfo_KEYS) ProtoMessage()    {}
func (*InvXmlFruInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c4a3ece7c164817, []int{0}
}

func (m *InvXmlFruInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvXmlFruInfo_KEYS.Unmarshal(m, b)
}
func (m *InvXmlFruInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvXmlFruInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *InvXmlFruInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvXmlFruInfo_KEYS.Merge(m, src)
}
func (m *InvXmlFruInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_InvXmlFruInfo_KEYS.Size(m)
}
func (m *InvXmlFruInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InvXmlFruInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InvXmlFruInfo_KEYS proto.InternalMessageInfo

func (m *InvXmlFruInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvXmlFruInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *InvXmlFruInfo_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *InvXmlFruInfo_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
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
	return fileDescriptor_4c4a3ece7c164817, []int{1}
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

type InvXmlFruInfo struct {
	ModuleAdministrativeState      string       `protobuf:"bytes,50,opt,name=module_administrative_state,json=moduleAdministrativeState,proto3" json:"module_administrative_state,omitempty"`
	ModulePowerAdministrativeState string       `protobuf:"bytes,51,opt,name=module_power_administrative_state,json=modulePowerAdministrativeState,proto3" json:"module_power_administrative_state,omitempty"`
	ModuleOperationalState         string       `protobuf:"bytes,52,opt,name=module_operational_state,json=moduleOperationalState,proto3" json:"module_operational_state,omitempty"`
	ModuleMonitorState             string       `protobuf:"bytes,53,opt,name=module_monitor_state,json=moduleMonitorState,proto3" json:"module_monitor_state,omitempty"`
	ModuleResetReason              string       `protobuf:"bytes,54,opt,name=module_reset_reason,json=moduleResetReason,proto3" json:"module_reset_reason,omitempty"`
	LastOperationalStateChange     *InvTimespec `protobuf:"bytes,55,opt,name=last_operational_state_change,json=lastOperationalStateChange,proto3" json:"last_operational_state_change,omitempty"`
	ModuleUpTime                   *InvTimespec `protobuf:"bytes,56,opt,name=module_up_time,json=moduleUpTime,proto3" json:"module_up_time,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}     `json:"-"`
	XXX_unrecognized               []byte       `json:"-"`
	XXX_sizecache                  int32        `json:"-"`
}

func (m *InvXmlFruInfo) Reset()         { *m = InvXmlFruInfo{} }
func (m *InvXmlFruInfo) String() string { return proto.CompactTextString(m) }
func (*InvXmlFruInfo) ProtoMessage()    {}
func (*InvXmlFruInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c4a3ece7c164817, []int{2}
}

func (m *InvXmlFruInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvXmlFruInfo.Unmarshal(m, b)
}
func (m *InvXmlFruInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvXmlFruInfo.Marshal(b, m, deterministic)
}
func (m *InvXmlFruInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvXmlFruInfo.Merge(m, src)
}
func (m *InvXmlFruInfo) XXX_Size() int {
	return xxx_messageInfo_InvXmlFruInfo.Size(m)
}
func (m *InvXmlFruInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InvXmlFruInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InvXmlFruInfo proto.InternalMessageInfo

func (m *InvXmlFruInfo) GetModuleAdministrativeState() string {
	if m != nil {
		return m.ModuleAdministrativeState
	}
	return ""
}

func (m *InvXmlFruInfo) GetModulePowerAdministrativeState() string {
	if m != nil {
		return m.ModulePowerAdministrativeState
	}
	return ""
}

func (m *InvXmlFruInfo) GetModuleOperationalState() string {
	if m != nil {
		return m.ModuleOperationalState
	}
	return ""
}

func (m *InvXmlFruInfo) GetModuleMonitorState() string {
	if m != nil {
		return m.ModuleMonitorState
	}
	return ""
}

func (m *InvXmlFruInfo) GetModuleResetReason() string {
	if m != nil {
		return m.ModuleResetReason
	}
	return ""
}

func (m *InvXmlFruInfo) GetLastOperationalStateChange() *InvTimespec {
	if m != nil {
		return m.LastOperationalStateChange
	}
	return nil
}

func (m *InvXmlFruInfo) GetModuleUpTime() *InvTimespec {
	if m != nil {
		return m.ModuleUpTime
	}
	return nil
}

func init() {
	proto.RegisterType((*InvXmlFruInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.hw_components.hw_component.attributes.fru_info.inv_xml_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.hw_components.hw_component.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvXmlFruInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.hw_components.hw_component.attributes.fru_info.inv_xml_fru_info")
}

func init() { proto.RegisterFile("inv_xml_fru_info.proto", fileDescriptor_4c4a3ece7c164817) }

var fileDescriptor_4c4a3ece7c164817 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4b, 0x6e, 0x53, 0x3d,
	0x14, 0xd6, 0xed, 0xdf, 0x56, 0xfa, 0x4d, 0x79, 0xc4, 0xb4, 0x95, 0x01, 0x81, 0x4a, 0x06, 0xa8,
	0x23, 0x43, 0x13, 0x1e, 0x1d, 0x21, 0x21, 0xc4, 0xa0, 0x42, 0x3c, 0x74, 0x03, 0x03, 0x46, 0x96,
	0x7b, 0xe3, 0xa4, 0x16, 0xd7, 0xe7, 0x58, 0xf6, 0x49, 0x5a, 0x16, 0xc0, 0x0a, 0xd8, 0x0d, 0xab,
	0x60, 0xc6, 0x76, 0x90, 0xed, 0xdb, 0x84, 0x86, 0x2e, 0xa0, 0x13, 0xc7, 0xfe, 0x5e, 0xfe, 0xec,
	0x58, 0x97, 0xed, 0x5a, 0x98, 0xab, 0x33, 0xd7, 0xaa, 0x49, 0x98, 0x29, 0x0b, 0x13, 0x94, 0x3e,
	0x20, 0x21, 0xff, 0x5e, 0x35, 0x36, 0x36, 0xa8, 0x2c, 0x46, 0x75, 0x16, 0x94, 0x6f, 0x35, 0xa9,
	0xe6, 0x44, 0x47, 0x65, 0x61, 0xee, 0xa6, 0x41, 0xa1, 0x37, 0x41, 0x26, 0x74, 0x82, 0xc1, 0x25,
	0xd0, 0x00, 0x61, 0xf8, 0x26, 0x83, 0x6e, 0xbe, 0xc6, 0x3c, 0xca, 0xd8, 0x22, 0xc5, 0x3c, 0xca,
	0x46, 0x87, 0x71, 0xcc, 0xa3, 0x3c, 0x39, 0x55, 0x0d, 0x3a, 0x8f, 0x60, 0x80, 0xe2, 0x85, 0x95,
	0xd4, 0x44, 0xc1, 0x1e, 0xcf, 0xc8, 0x44, 0x79, 0xde, 0xa6, 0xef, 0xd9, 0xce, 0x6a, 0x43, 0xf5,
	0xf6, 0xcd, 0x97, 0x11, 0xe7, 0x6c, 0x1d, 0xb4, 0x33, 0xa2, 0xda, 0xab, 0xf6, 0xff, 0xaf, 0xf3,
	0x9c, 0xef, 0xb0, 0xcd, 0xf4, 0xab, 0x0e, 0xc4, 0x5a, 0x46, 0x37, 0xd2, 0xea, 0x60, 0x01, 0x0f,
	0xc4, 0x7f, 0x4b, 0x78, 0xb0, 0x80, 0x87, 0x62, 0x7d, 0x09, 0x0f, 0xfb, 0x53, 0xb6, 0x95, 0x76,
	0x24, 0xeb, 0x4c, 0xf4, 0xa6, 0xe1, 0x8f, 0xd8, 0xcd, 0x34, 0x57, 0x16, 0x54, 0x34, 0x0d, 0xc2,
	0x38, 0xe6, 0x3d, 0x7b, 0xf5, 0xf5, 0x04, 0x1f, 0xc1, 0xa8, 0x80, 0xfc, 0x31, 0xdb, 0x3e, 0xd7,
	0x81, 0x06, 0x5c, 0x88, 0xd7, 0xb2, 0xb8, 0x57, 0xc4, 0xef, 0x35, 0x60, 0x67, 0xe8, 0xff, 0xda,
	0x60, 0xb7, 0x56, 0xcf, 0xc6, 0x5f, 0xb2, 0x7b, 0x0e, 0xc7, 0xb3, 0xd6, 0x28, 0x3d, 0x76, 0x16,
	0x6c, 0xa4, 0xa0, 0xc9, 0xce, 0x8d, 0x8a, 0xa4, 0xc9, 0x88, 0x41, 0x6e, 0x7a, 0xa7, 0x48, 0x5e,
	0x5d, 0x50, 0x8c, 0x92, 0x80, 0x1f, 0xb1, 0x87, 0x9d, 0xdf, 0xe3, 0xa9, 0x09, 0x97, 0xa7, 0x0c,
	0x73, 0xca, 0x83, 0x22, 0xfc, 0x98, 0x74, 0x97, 0x45, 0x1d, 0x32, 0xd1, 0x45, 0xa5, 0x3f, 0x5b,
	0x93, 0x45, 0xd0, 0x6d, 0x97, 0xf0, 0x34, 0x27, 0xec, 0x16, 0xfe, 0xc3, 0x92, 0x2e, 0xce, 0x27,
	0x6c, 0xbb, 0x73, 0x3a, 0x04, 0x4b, 0x18, 0x3a, 0xd7, 0xb3, 0xec, 0xe2, 0x85, 0x7b, 0x57, 0xa8,
	0xe2, 0x90, 0xec, 0x76, 0xe7, 0x08, 0x26, 0x1a, 0x52, 0xc1, 0xe8, 0x88, 0x20, 0x9e, 0x67, 0x43,
	0xaf, 0x50, 0x75, 0x62, 0xea, 0x4c, 0xf0, 0xdf, 0x15, 0xbb, 0xdf, 0xea, 0x48, 0xff, 0x56, 0x4b,
	0xaf, 0x14, 0xa6, 0x46, 0xbc, 0xd8, 0xab, 0xf6, 0xaf, 0x0d, 0x7e, 0x54, 0xf2, 0x4a, 0xbc, 0x63,
	0xf9, 0xf7, 0x93, 0xaa, 0xef, 0xa6, 0xea, 0xab, 0xb7, 0xf6, 0x3a, 0xf7, 0xe6, 0x3f, 0x2b, 0x76,
	0xa3, 0xbb, 0x8a, 0x99, 0xcf, 0x16, 0x71, 0x78, 0x85, 0x8f, 0xb2, 0x55, 0xba, 0x7e, 0xf6, 0x9f,
	0xac, 0x33, 0xc7, 0x9b, 0xf9, 0xe3, 0x31, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x68, 0x10,
	0xc8, 0x56, 0x04, 0x00, 0x00,
}
