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

package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_slots_slot_cards_card_sub_slots_sub_slot_module_port_slots_port_slot_sensors_sensor_attributes_fru_info

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
	proto.RegisterType((*InvXmlFruInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.port_slots.port_slot.sensors.sensor.attributes.fru_info.inv_xml_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.port_slots.port_slot.sensors.sensor.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvXmlFruInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.port_slots.port_slot.sensors.sensor.attributes.fru_info.inv_xml_fru_info")
}

func init() { proto.RegisterFile("inv_xml_fru_info.proto", fileDescriptor_4c4a3ece7c164817) }

var fileDescriptor_4c4a3ece7c164817 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0x96, 0x36, 0x12, 0xa6, 0xfc, 0xc4, 0xb4, 0x95, 0x01, 0x81, 0x4a, 0x0e, 0xa8, 0x27,
	0x43, 0xf3, 0x03, 0x3d, 0x21, 0x21, 0xc4, 0xa1, 0x42, 0xfc, 0x68, 0x03, 0x07, 0x4e, 0x96, 0xb3,
	0x71, 0x82, 0xc5, 0xae, 0x67, 0xe5, 0xf1, 0x86, 0xf2, 0x22, 0x1c, 0x39, 0xf3, 0x04, 0xbc, 0x0d,
	0x2f, 0xc2, 0x09, 0xf9, 0xa7, 0xbb, 0x90, 0xf6, 0x05, 0x72, 0xb1, 0x27, 0xdf, 0x9f, 0x66, 0x46,
	0xce, 0x92, 0x03, 0x6d, 0x56, 0xe2, 0xac, 0x2a, 0xc5, 0xc2, 0x36, 0x42, 0x9b, 0x05, 0xf0, 0xda,
	0x82, 0x03, 0xfa, 0x33, 0x2b, 0x34, 0x16, 0x20, 0x34, 0xa0, 0x38, 0xb3, 0xa2, 0x2e, 0xa5, 0x13,
	0xc5, 0x67, 0x89, 0x42, 0x9b, 0x55, 0xb5, 0xb4, 0x02, 0x6a, 0x65, 0xb9, 0x47, 0x17, 0x60, 0x2b,
	0x0f, 0x2a, 0xe3, 0xc0, 0x7e, 0xe3, 0x56, 0x16, 0x5f, 0x30, 0x9c, 0x1c, 0x4b, 0x70, 0x18, 0x4e,
	0x5e, 0x48, 0x3b, 0xc7, 0x70, 0x72, 0x6c, 0x66, 0x22, 0x31, 0xa9, 0xe2, 0x15, 0xcc, 0x9b, 0x52,
	0xf1, 0x1a, 0xac, 0x4b, 0x54, 0x5b, 0x72, 0x54, 0x06, 0xc1, 0x62, 0xba, 0xb9, 0x74, 0xce, 0xea,
	0x59, 0xe3, 0x14, 0xf2, 0xf3, 0x86, 0x07, 0x3f, 0x32, 0xb2, 0xbf, 0x3e, 0x85, 0x78, 0xfd, 0xea,
	0xd3, 0x94, 0x52, 0xb2, 0x6d, 0x64, 0xa5, 0x58, 0x76, 0x98, 0x1d, 0x5d, 0xcd, 0x43, 0x4d, 0xf7,
	0x49, 0xcf, 0xdf, 0xe2, 0x98, 0x6d, 0x05, 0x74, 0xc7, 0xff, 0x3a, 0x6e, 0xe1, 0x21, 0xbb, 0xd2,
	0xc1, 0xc3, 0x16, 0x1e, 0xb1, 0xed, 0x0e, 0x1e, 0xb5, 0xf0, 0x98, 0xed, 0x74, 0xf0, 0xb8, 0x85,
	0x27, 0xac, 0xd7, 0xc1, 0x93, 0xc1, 0x92, 0xec, 0xfa, 0xfe, 0x9c, 0xae, 0x14, 0xd6, 0xaa, 0xa0,
	0x8f, 0xc8, 0x4d, 0x5f, 0x0b, 0x6d, 0x04, 0xaa, 0x02, 0xcc, 0x1c, 0x43, 0x87, 0xfd, 0xfc, 0xba,
	0x87, 0x4f, 0xcd, 0x34, 0x82, 0xf4, 0x31, 0xd9, 0x3b, 0xd7, 0x19, 0x69, 0xa0, 0x15, 0x6f, 0x05,
	0x71, 0x3f, 0x8a, 0xdf, 0x4a, 0x03, 0xc9, 0x30, 0xf8, 0xde, 0x23, 0xb7, 0xd6, 0x37, 0x41, 0x9f,
	0x93, 0x7b, 0x71, 0xc1, 0x42, 0xce, 0x2b, 0x6d, 0x34, 0x3a, 0x2b, 0x9d, 0x5e, 0x29, 0x81, 0x4e,
	0x3a, 0xc5, 0x86, 0xa1, 0xd3, 0x3b, 0x51, 0xf2, 0xe2, 0x3f, 0xc5, 0xd4, 0x0b, 0xe8, 0x29, 0x79,
	0x98, 0xfc, 0x35, 0x7c, 0x55, 0xf6, 0xf2, 0x94, 0x51, 0x48, 0x79, 0x10, 0x85, 0xef, 0xbd, 0xee,
	0xb2, 0xa8, 0x13, 0xc2, 0x52, 0x94, 0x7f, 0x3e, 0xd2, 0x69, 0x30, 0xb2, 0x4c, 0x09, 0xe3, 0x90,
	0x70, 0x10, 0xf9, 0x77, 0x1d, 0x1d, 0x9d, 0x4f, 0xc8, 0x5e, 0x72, 0x56, 0x60, 0xb4, 0x03, 0x9b,
	0x5c, 0x93, 0xe0, 0xa2, 0x91, 0x7b, 0x13, 0xa9, 0xe8, 0xe0, 0xe4, 0x76, 0x72, 0x58, 0x85, 0xca,
	0x09, 0xab, 0x24, 0x82, 0x61, 0x4f, 0x83, 0xa1, 0x1f, 0xa9, 0xdc, 0x33, 0x79, 0x20, 0xe8, 0x9f,
	0x8c, 0xdc, 0x2f, 0x25, 0xba, 0x8b, 0xad, 0xf9, 0x77, 0x6f, 0x96, 0x8a, 0x3d, 0x3b, 0xcc, 0x8e,
	0xae, 0x0d, 0x7f, 0x65, 0x7c, 0x53, 0xfe, 0x19, 0xfc, 0xdf, 0x57, 0x97, 0xdf, 0xf5, 0xd3, 0xad,
	0x2f, 0xf6, 0x65, 0x18, 0x8d, 0xfe, 0xce, 0xc8, 0x8d, 0xb4, 0xad, 0xa6, 0x0e, 0x16, 0x76, 0xb2,
	0xd9, 0xd3, 0xee, 0xc6, 0xb0, 0x8f, 0xf5, 0x07, 0x5d, 0xa9, 0x59, 0x2f, 0x7c, 0xd4, 0x46, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x46, 0x37, 0x29, 0x0b, 0xee, 0x04, 0x00, 0x00,
}
