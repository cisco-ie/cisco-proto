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

package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_slots_slot_cards_card_sub_slots_sub_slot_module_port_slots_port_slot_portses_ports_hw_components_hw_component_attributes_fru_info

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
	Name_4               string   `protobuf:"bytes,5,opt,name=name_4,json=name4,proto3" json:"name_4,omitempty"`
	Name_5               string   `protobuf:"bytes,6,opt,name=name_5,json=name5,proto3" json:"name_5,omitempty"`
	Name_6               string   `protobuf:"bytes,7,opt,name=name_6,json=name6,proto3" json:"name_6,omitempty"`
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

func (m *InvXmlFruInfo_KEYS) GetName_4() string {
	if m != nil {
		return m.Name_4
	}
	return ""
}

func (m *InvXmlFruInfo_KEYS) GetName_5() string {
	if m != nil {
		return m.Name_5
	}
	return ""
}

func (m *InvXmlFruInfo_KEYS) GetName_6() string {
	if m != nil {
		return m.Name_6
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
	proto.RegisterType((*InvXmlFruInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.port_slots.port_slot.portses.ports.hw_components.hw_component.attributes.fru_info.inv_xml_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.port_slots.port_slot.portses.ports.hw_components.hw_component.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvXmlFruInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.port_slots.port_slot.portses.ports.hw_components.hw_component.attributes.fru_info.inv_xml_fru_info")
}

func init() { proto.RegisterFile("inv_xml_fru_info.proto", fileDescriptor_4c4a3ece7c164817) }

var fileDescriptor_4c4a3ece7c164817 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x86, 0x35, 0xa1, 0x0d, 0xc2, 0x94, 0x4b, 0x4c, 0x5b, 0x19, 0x10, 0xa8, 0x64, 0x81, 0xba,
	0x32, 0x34, 0x37, 0xba, 0x42, 0x42, 0x88, 0x45, 0x85, 0xb8, 0x68, 0x02, 0x0b, 0x56, 0x96, 0x33,
	0x71, 0x52, 0x8b, 0xb1, 0xcf, 0xc8, 0xf6, 0xa4, 0xe5, 0x3d, 0x78, 0x15, 0x5e, 0x80, 0x25, 0x2b,
	0xde, 0x86, 0x2d, 0xf2, 0xa5, 0x33, 0x10, 0xf2, 0x02, 0x74, 0xe3, 0x9c, 0xf9, 0xfe, 0xff, 0x3f,
	0xb1, 0x4f, 0xc6, 0x41, 0xfb, 0x52, 0xaf, 0xd8, 0xb9, 0x2a, 0xd9, 0xc2, 0xd4, 0x4c, 0xea, 0x05,
	0xd0, 0xca, 0x80, 0x03, 0xfc, 0x3d, 0x2b, 0xa4, 0x2d, 0x80, 0x49, 0xb0, 0xec, 0xdc, 0xb0, 0xaa,
	0xe4, 0x8e, 0x15, 0xa7, 0xdc, 0x32, 0xa9, 0x57, 0x6a, 0x69, 0x18, 0x54, 0xc2, 0x50, 0x4f, 0x17,
	0x60, 0x94, 0x87, 0x42, 0x3b, 0x30, 0x5f, 0xa8, 0xe1, 0xc5, 0x67, 0x1b, 0x56, 0x6a, 0x4b, 0x70,
	0x36, 0xac, 0xb4, 0xe0, 0x66, 0x6e, 0xc3, 0x4a, 0x6d, 0x3d, 0x63, 0x49, 0x49, 0x15, 0x55, 0x30,
	0xaf, 0x4b, 0x41, 0x2b, 0x30, 0x2e, 0x49, 0x4d, 0x19, 0x2a, 0x2b, 0x22, 0xb1, 0xf4, 0xf4, 0x8c,
	0x15, 0xa0, 0x2a, 0xd0, 0x42, 0xaf, 0x3d, 0x51, 0xee, 0x9c, 0x91, 0xb3, 0xda, 0x09, 0x4b, 0x2f,
	0x8e, 0xd2, 0xff, 0x96, 0xa1, 0xbd, 0xf5, 0xf3, 0xb1, 0xd7, 0xaf, 0x3e, 0x4d, 0x31, 0x46, 0x5b,
	0x9a, 0x2b, 0x41, 0xb2, 0x83, 0xec, 0xf0, 0x5a, 0x1e, 0x6a, 0xbc, 0x87, 0xba, 0xfe, 0x93, 0x1d,
	0x91, 0x4e, 0xa0, 0xdb, 0xfe, 0xe9, 0xa8, 0xc1, 0x03, 0x72, 0xa5, 0xc5, 0x83, 0x06, 0x0f, 0xc9,
	0x56, 0x8b, 0x87, 0x0d, 0x1e, 0x91, 0xed, 0x16, 0x8f, 0x1a, 0x3c, 0x26, 0xdd, 0x16, 0x8f, 0x1b,
	0x3c, 0x21, 0x57, 0x5b, 0x3c, 0xe9, 0x2f, 0xd1, 0x8e, 0xdf, 0xb6, 0x93, 0x4a, 0xd8, 0x4a, 0x14,
	0xf8, 0x31, 0xba, 0xe5, 0x6b, 0x26, 0x35, 0xb3, 0xa2, 0x00, 0x3d, 0xb7, 0x61, 0xe3, 0xbd, 0xfc,
	0x86, 0xc7, 0x27, 0x7a, 0x1a, 0x21, 0x7e, 0x82, 0x76, 0x2f, 0x7c, 0x9a, 0x6b, 0x68, 0xcc, 0x9d,
	0x60, 0xee, 0x45, 0xf3, 0x5b, 0xae, 0x21, 0x05, 0xfa, 0x3f, 0xba, 0xe8, 0xf6, 0xfa, 0x80, 0xf0,
	0x73, 0x74, 0x3f, 0xfe, 0x22, 0x8c, 0xcf, 0x95, 0xd4, 0xd2, 0x3a, 0xc3, 0x9d, 0x5c, 0x09, 0x66,
	0x1d, 0x77, 0x82, 0x0c, 0xc2, 0x4e, 0xef, 0x46, 0xcb, 0x8b, 0xbf, 0x1c, 0x53, 0x6f, 0xc0, 0x27,
	0xe8, 0x51, 0xca, 0x57, 0x70, 0x26, 0xcc, 0xe6, 0x2e, 0xc3, 0xd0, 0xe5, 0x61, 0x34, 0xbe, 0xf7,
	0xbe, 0x4d, 0xad, 0x8e, 0x11, 0x49, 0xad, 0xfc, 0xfb, 0xc6, 0x9d, 0x04, 0xcd, 0xcb, 0xd4, 0x61,
	0x14, 0x3a, 0xec, 0x47, 0xfd, 0x5d, 0x2b, 0xc7, 0xe4, 0x53, 0xb4, 0x9b, 0x92, 0x0a, 0xb4, 0x74,
	0x60, 0x52, 0x6a, 0x1c, 0x52, 0x38, 0x6a, 0x6f, 0xa2, 0x14, 0x13, 0x14, 0xdd, 0x49, 0x09, 0x23,
	0xac, 0x70, 0xcc, 0x08, 0x6e, 0x41, 0x93, 0x49, 0x08, 0xf4, 0xa2, 0x94, 0x7b, 0x25, 0x0f, 0x02,
	0xfe, 0xda, 0x41, 0x0f, 0x4a, 0x6e, 0xdd, 0xbf, 0x5b, 0xf3, 0x17, 0x45, 0x2f, 0x05, 0x79, 0x76,
	0x90, 0x1d, 0x5e, 0x1f, 0xfc, 0xcc, 0xe8, 0xff, 0x7f, 0x95, 0xe8, 0x9f, 0xef, 0x63, 0x7e, 0xcf,
	0x9f, 0x7b, 0x7d, 0xe4, 0x2f, 0xc3, 0xa1, 0xf1, 0xaf, 0x0c, 0xdd, 0x4c, 0x73, 0xac, 0xab, 0x10,
	0x21, 0xc7, 0x97, 0x75, 0x0e, 0x3b, 0xf1, 0x6b, 0x3e, 0x56, 0x1f, 0xa4, 0x12, 0xb3, 0x6e, 0xf8,
	0xe7, 0x1c, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x81, 0x67, 0x43, 0xcf, 0x53, 0x05, 0x00, 0x00,
}
