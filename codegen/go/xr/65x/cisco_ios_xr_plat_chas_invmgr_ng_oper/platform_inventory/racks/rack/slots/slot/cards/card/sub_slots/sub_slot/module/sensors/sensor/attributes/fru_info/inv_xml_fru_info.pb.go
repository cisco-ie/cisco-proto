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

package cisco_ios_xr_plat_chas_invmgr_ng_oper_platform_inventory_racks_rack_slots_slot_cards_card_sub_slots_sub_slot_module_sensors_sensor_attributes_fru_info

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
	proto.RegisterType((*InvXmlFruInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_ng_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.sensors.sensor.attributes.fru_info.inv_xml_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_plat_chas_invmgr_ng_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.sensors.sensor.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvXmlFruInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_ng_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.sensors.sensor.attributes.fru_info.inv_xml_fru_info")
}

func init() { proto.RegisterFile("inv_xml_fru_info.proto", fileDescriptor_4c4a3ece7c164817) }

var fileDescriptor_4c4a3ece7c164817 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0x4d, 0x6e, 0x13, 0x31,
	0x14, 0xd6, 0x94, 0xb6, 0x12, 0xa6, 0xfc, 0xc4, 0xb4, 0x95, 0x01, 0x81, 0x4a, 0x16, 0xa8, 0x2b,
	0x43, 0x93, 0x00, 0x5d, 0x21, 0x21, 0xc4, 0xa2, 0x42, 0xfc, 0x68, 0x02, 0x0b, 0x56, 0x96, 0x33,
	0x71, 0x82, 0xc5, 0xcc, 0x7b, 0x23, 0x3f, 0x4f, 0x28, 0x27, 0xe0, 0x06, 0x9c, 0x80, 0xdb, 0x70,
	0x11, 0xd6, 0x9c, 0x00, 0xf9, 0xa7, 0x09, 0x84, 0x1e, 0x00, 0x36, 0xb6, 0xe7, 0xfb, 0x79, 0xfa,
	0xde, 0xd3, 0xd3, 0xb0, 0x7d, 0x0b, 0x0b, 0x75, 0xda, 0xd4, 0x6a, 0xe6, 0x3a, 0x65, 0x61, 0x86,
	0xb2, 0x75, 0xe8, 0x91, 0x7f, 0x2d, 0x2a, 0x4b, 0x15, 0x2a, 0x8b, 0xa4, 0x4e, 0x9d, 0x6a, 0x6b,
	0xed, 0x55, 0xf5, 0x41, 0x93, 0xb2, 0xb0, 0x68, 0xe6, 0x4e, 0xc1, 0x5c, 0x61, 0x6b, 0x9c, 0x0c,
	0xc4, 0x0c, 0x5d, 0x13, 0x70, 0x03, 0x1e, 0xdd, 0x67, 0xe9, 0x74, 0xf5, 0x91, 0xe2, 0x29, 0xa9,
	0x46, 0x4f, 0xf1, 0x94, 0x95, 0x76, 0x53, 0x8a, 0xa7, 0xa4, 0x6e, 0xa2, 0x32, 0x93, 0x5f, 0xb2,
	0xc1, 0x69, 0x57, 0x1b, 0x49, 0x06, 0x08, 0x1d, 0xe5, 0x5b, 0x6a, 0xef, 0x9d, 0x9d, 0x74, 0xde,
	0x90, 0x3c, 0x8b, 0xd7, 0xff, 0x52, 0xb0, 0xbd, 0xf5, 0xcc, 0xea, 0xc5, 0xf3, 0xf7, 0x63, 0xce,
	0xd9, 0x26, 0xe8, 0xc6, 0x88, 0xe2, 0xa0, 0x38, 0xbc, 0x58, 0xc6, 0x37, 0xdf, 0x63, 0xdb, 0xe1,
	0x56, 0x47, 0x62, 0x23, 0xa2, 0x5b, 0xe1, 0xeb, 0x68, 0x09, 0x0f, 0xc4, 0x85, 0x15, 0x3c, 0x58,
	0xc2, 0x43, 0xb1, 0xb9, 0x82, 0x87, 0x4b, 0x78, 0x24, 0xb6, 0x56, 0xf0, 0xa8, 0x3f, 0x67, 0x3b,
	0x21, 0x88, 0xb7, 0x8d, 0xa1, 0xd6, 0x54, 0xfc, 0x1e, 0xbb, 0x1a, 0xde, 0xca, 0x82, 0x22, 0x53,
	0x21, 0x4c, 0x29, 0x46, 0xe9, 0x95, 0x97, 0x03, 0x7c, 0x02, 0xe3, 0x04, 0xf2, 0xfb, 0x6c, 0xf7,
	0x4c, 0x07, 0x1a, 0x70, 0x29, 0xde, 0x88, 0xe2, 0x5e, 0x12, 0xbf, 0xd2, 0x80, 0xd9, 0xd0, 0xff,
	0xb9, 0xc5, 0xae, 0xad, 0xb7, 0xcc, 0x9f, 0xb0, 0x5b, 0x69, 0x68, 0x4a, 0x4f, 0x1b, 0x0b, 0x96,
	0xbc, 0xd3, 0xde, 0x2e, 0x8c, 0x22, 0xaf, 0xbd, 0x11, 0x83, 0x98, 0xf4, 0x46, 0x92, 0x3c, 0xfd,
	0x43, 0x31, 0x0e, 0x02, 0x7e, 0xc2, 0xee, 0x66, 0x7f, 0x8b, 0x9f, 0x8c, 0x3b, 0xbf, 0xca, 0x30,
	0x56, 0xb9, 0x93, 0x84, 0x6f, 0x82, 0xee, 0xbc, 0x52, 0xc7, 0x4c, 0xe4, 0x52, 0x61, 0x25, 0xb4,
	0xb7, 0x08, 0xba, 0xce, 0x15, 0x46, 0xb1, 0xc2, 0x7e, 0xe2, 0x5f, 0xaf, 0xe8, 0xe4, 0x7c, 0xc0,
	0x76, 0xb3, 0xb3, 0x41, 0xb0, 0x1e, 0x5d, 0x76, 0x3d, 0x8c, 0x2e, 0x9e, 0xb8, 0x97, 0x89, 0x4a,
	0x0e, 0xc9, 0xae, 0x67, 0x87, 0x33, 0x64, 0xbc, 0x72, 0x46, 0x13, 0x82, 0x78, 0x14, 0x0d, 0xbd,
	0x44, 0x95, 0x81, 0x29, 0x23, 0xc1, 0x7f, 0x14, 0xec, 0x76, 0xad, 0xc9, 0xff, 0x1d, 0x2d, 0xac,
	0x33, 0xcc, 0x8d, 0x78, 0x7c, 0x50, 0x1c, 0x5e, 0x1a, 0x7c, 0x2b, 0xe4, 0xbf, 0xb9, 0xf0, 0xf2,
	0xf7, 0x1d, 0x2b, 0x6f, 0x86, 0x5e, 0xd6, 0xc7, 0xf8, 0x2c, 0x36, 0xc2, 0xbf, 0x17, 0xec, 0x4a,
	0x9e, 0x4d, 0xd7, 0x46, 0x8b, 0x38, 0xfe, 0x9f, 0x7a, 0xdb, 0x49, 0xbe, 0x77, 0xed, 0x5b, 0xdb,
	0x98, 0xc9, 0x76, 0xfc, 0x0f, 0x0d, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x81, 0x1f, 0x40,
	0xa1, 0x04, 0x00, 0x00,
}
