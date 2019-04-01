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

package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_slots_slot_attributes_fru_info

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
	proto.RegisterType((*InvXmlFruInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.attributes.fru_info.inv_xml_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvXmlFruInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.attributes.fru_info.inv_xml_fru_info")
}

func init() { proto.RegisterFile("inv_xml_fru_info.proto", fileDescriptor_4c4a3ece7c164817) }

var fileDescriptor_4c4a3ece7c164817 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x49, 0x6f, 0xd4, 0x30,
	0x14, 0x56, 0x2a, 0x5a, 0x09, 0x53, 0x96, 0x31, 0x6d, 0x15, 0x40, 0xa0, 0x32, 0x07, 0xd4, 0x93,
	0xa1, 0x2d, 0x4b, 0x4f, 0x48, 0x80, 0x38, 0x54, 0x88, 0x45, 0x19, 0x38, 0x70, 0x7a, 0x72, 0x33,
	0x9e, 0xc1, 0x22, 0x7e, 0x2f, 0xb2, 0x5f, 0x86, 0x72, 0xe5, 0x2f, 0xf0, 0x47, 0xf8, 0x89, 0xc8,
	0x4b, 0x59, 0x86, 0xb9, 0xa2, 0x5e, 0x1c, 0xe7, 0xdb, 0xfc, 0x25, 0x79, 0x11, 0x3b, 0x16, 0x17,
	0x70, 0xea, 0x3a, 0x98, 0xf9, 0x01, 0x2c, 0xce, 0x48, 0xf5, 0x9e, 0x98, 0x24, 0xb4, 0x36, 0xb4,
	0x04, 0x96, 0x02, 0x9c, 0x7a, 0xe8, 0x3b, 0xcd, 0xd0, 0x7e, 0xd2, 0x01, 0x2c, 0x2e, 0xdc, 0xdc,
	0x03, 0xf5, 0xc6, 0xab, 0x88, 0xce, 0xc8, 0xbb, 0x08, 0x1a, 0x64, 0xf2, 0x5f, 0x95, 0xd7, 0xed,
	0xe7, 0x90, 0x56, 0x15, 0x3a, 0xe2, 0x90, 0x56, 0xa5, 0x99, 0xbd, 0x3d, 0x19, 0xd8, 0x04, 0x75,
	0x76, 0xcc, 0xf8, 0xb9, 0xd8, 0x5e, 0x3e, 0x1a, 0x5e, 0xbd, 0xfc, 0x38, 0x91, 0x52, 0x5c, 0x40,
	0xed, 0x4c, 0x5d, 0xed, 0x56, 0x7b, 0x17, 0x9b, 0xb4, 0x97, 0xdb, 0x62, 0x23, 0x5e, 0x61, 0xbf,
	0x5e, 0x4b, 0xe8, 0x7a, 0xbc, 0xdb, 0x1f, 0xcf, 0xc5, 0x66, 0xcc, 0x60, 0xeb, 0x4c, 0xe8, 0x4d,
	0x2b, 0xef, 0x89, 0xab, 0x71, 0x0f, 0x16, 0x21, 0x98, 0x96, 0x70, 0x1a, 0x52, 0xca, 0xa8, 0xb9,
	0x1c, 0xe1, 0x63, 0x9c, 0x64, 0x50, 0xde, 0x17, 0x5b, 0x67, 0x3a, 0xd4, 0x48, 0xbf, 0xc4, 0x6b,
	0x49, 0x3c, 0xca, 0xe2, 0x37, 0x1a, 0xa9, 0x18, 0xc6, 0xdf, 0xd6, 0xc5, 0xb5, 0xe5, 0xb6, 0xf2,
	0xa9, 0xb8, 0xe5, 0x68, 0x3a, 0x74, 0x06, 0xf4, 0xd4, 0x59, 0xb4, 0x81, 0xbd, 0x66, 0xbb, 0x30,
	0x10, 0x58, 0xb3, 0xa9, 0x0f, 0x52, 0xd3, 0x1b, 0x59, 0xf2, 0xec, 0x2f, 0xc5, 0x24, 0x0a, 0xe4,
	0xb1, 0xb8, 0x5b, 0xfc, 0x3d, 0x7d, 0x31, 0x7e, 0x75, 0xca, 0x61, 0x4a, 0xb9, 0x93, 0x85, 0xef,
	0xa2, 0x6e, 0x55, 0xd4, 0x91, 0xa8, 0x4b, 0x54, 0xfc, 0x30, 0x9a, 0x2d, 0xa1, 0xee, 0x4a, 0xc2,
	0xc3, 0x94, 0xb0, 0x93, 0xf9, 0xb7, 0xbf, 0xe9, 0xec, 0x7c, 0x20, 0xb6, 0x8a, 0xd3, 0x11, 0x5a,
	0x26, 0x5f, 0x5c, 0x8f, 0x92, 0x4b, 0x66, 0xee, 0x75, 0xa6, 0xb2, 0x43, 0x89, 0xeb, 0xc5, 0xe1,
	0x4d, 0x30, 0x0c, 0xde, 0xe8, 0x40, 0x58, 0x3f, 0x4e, 0x86, 0x51, 0xa6, 0x9a, 0xc8, 0x34, 0x89,
	0x90, 0x3f, 0x2a, 0x71, 0xbb, 0xd3, 0x81, 0xff, 0xad, 0x16, 0x27, 0x0a, 0xe7, 0xa6, 0x7e, 0xb2,
	0x5b, 0xed, 0x5d, 0x3a, 0x70, 0xea, 0x3f, 0x8f, 0x9c, 0xfa, 0x73, 0x56, 0x9a, 0x9b, 0xb1, 0xd3,
	0xf2, 0xeb, 0x78, 0x91, 0x0a, 0xc9, 0xef, 0x95, 0xb8, 0x52, 0x9e, 0x71, 0xe8, 0x93, 0xa5, 0x3e,
	0x3a, 0x8f, 0x8e, 0x9b, 0xb9, 0xc4, 0x87, 0xfe, 0xbd, 0x75, 0xe6, 0x64, 0x23, 0xfd, 0x99, 0x87,
	0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x59, 0xf3, 0x72, 0xb3, 0x03, 0x00, 0x00,
}
