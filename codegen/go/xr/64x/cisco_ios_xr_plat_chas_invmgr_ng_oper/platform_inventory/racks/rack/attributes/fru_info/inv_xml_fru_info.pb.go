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

package cisco_ios_xr_plat_chas_invmgr_ng_oper_platform_inventory_racks_rack_attributes_fru_info

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
	proto.RegisterType((*InvXmlFruInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_ng_oper.platform_inventory.racks.rack.attributes.fru_info.inv_xml_fru_info_KEYS")
	proto.RegisterType((*InvTimespec)(nil), "cisco_ios_xr_plat_chas_invmgr_ng_oper.platform_inventory.racks.rack.attributes.fru_info.inv_timespec")
	proto.RegisterType((*InvXmlFruInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_ng_oper.platform_inventory.racks.rack.attributes.fru_info.inv_xml_fru_info")
}

func init() { proto.RegisterFile("inv_xml_fru_info.proto", fileDescriptor_4c4a3ece7c164817) }

var fileDescriptor_4c4a3ece7c164817 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcb, 0x8a, 0x14, 0x31,
	0x14, 0x86, 0x29, 0x19, 0x04, 0xe3, 0x78, 0xe9, 0x38, 0x0e, 0xa5, 0xa2, 0x8c, 0xbd, 0x90, 0x01,
	0x21, 0xca, 0x8c, 0x97, 0x59, 0x09, 0x22, 0x2e, 0x06, 0xf1, 0x42, 0xb5, 0x22, 0xae, 0x0e, 0x99,
	0xea, 0xd3, 0x65, 0xb0, 0x72, 0x4e, 0x91, 0xa4, 0xda, 0xf1, 0x19, 0x7c, 0x11, 0x9f, 0xcf, 0x27,
	0x90, 0x5c, 0xda, 0x4b, 0xdb, 0x5b, 0xdd, 0x34, 0xe1, 0xfc, 0xdf, 0xff, 0xe7, 0xef, 0x70, 0x4a,
	0xec, 0x1a, 0x5a, 0xc2, 0xa9, 0xed, 0x61, 0xe1, 0x46, 0x30, 0xb4, 0x60, 0x35, 0x38, 0x0e, 0x2c,
	0xdf, 0xb7, 0xc6, 0xb7, 0x0c, 0x86, 0x3d, 0x9c, 0x3a, 0x18, 0x7a, 0x1d, 0xa0, 0xfd, 0xa8, 0x3d,
	0x18, 0x5a, 0xda, 0xce, 0x01, 0x75, 0xc0, 0x03, 0x3a, 0x15, 0x85, 0x05, 0x3b, 0x1b, 0xe7, 0x48,
	0x81, 0xdd, 0x17, 0xe5, 0x74, 0xfb, 0xc9, 0xa7, 0x5f, 0xa5, 0x43, 0x70, 0xe6, 0x64, 0x0c, 0xe8,
	0xd5, 0x2a, 0x7e, 0x7a, 0x57, 0x5c, 0x5d, 0xbf, 0x12, 0x5e, 0x3c, 0xff, 0x30, 0x93, 0x52, 0x6c,
	0x91, 0xb6, 0x58, 0x57, 0x7b, 0xd5, 0xfe, 0xb9, 0x26, 0x9d, 0xa7, 0x9d, 0xd8, 0x8e, 0x70, 0x30,
	0x16, 0xfd, 0x80, 0xad, 0xbc, 0x23, 0x2e, 0xc5, 0x33, 0x18, 0x02, 0x8f, 0x2d, 0xd3, 0xdc, 0x27,
	0x7c, 0xd2, 0x5c, 0x88, 0xe3, 0x63, 0x9a, 0xe5, 0xa1, 0xbc, 0x27, 0x76, 0x56, 0x1c, 0x69, 0xe2,
	0x9f, 0xf0, 0x99, 0x04, 0x4f, 0x32, 0xfc, 0x4a, 0x13, 0x17, 0xc3, 0xf4, 0xfb, 0x96, 0xb8, 0xbc,
	0x5e, 0x4b, 0x3e, 0x11, 0x37, 0x2c, 0xcf, 0xc7, 0x1e, 0x41, 0xcf, 0xad, 0x21, 0xe3, 0x83, 0xd3,
	0xc1, 0x2c, 0x11, 0x7c, 0xd0, 0x01, 0xeb, 0x83, 0x54, 0xf4, 0x5a, 0x46, 0x9e, 0xfe, 0x41, 0xcc,
	0x22, 0x20, 0x8f, 0xc5, 0xed, 0xe2, 0x1f, 0xf8, 0x33, 0xba, 0xcd, 0x29, 0x87, 0x29, 0xe5, 0x56,
	0x06, 0xdf, 0x44, 0x6e, 0x53, 0xd4, 0x91, 0xa8, 0x4b, 0x54, 0x7c, 0x76, 0x1d, 0x0c, 0x93, 0xee,
	0x4b, 0xc2, 0x83, 0x94, 0xb0, 0x9b, 0xf5, 0xd7, 0xbf, 0xe4, 0xec, 0xbc, 0x2f, 0x76, 0x8a, 0xd3,
	0x32, 0x99, 0xc0, 0xae, 0xb8, 0x1e, 0x26, 0x97, 0xcc, 0xda, 0xcb, 0x2c, 0x65, 0x87, 0x12, 0x57,
	0x8a, 0xc3, 0xa1, 0xc7, 0x00, 0x0e, 0xb5, 0x67, 0xaa, 0x1f, 0x25, 0xc3, 0x24, 0x4b, 0x4d, 0x54,
	0x9a, 0x24, 0xc8, 0x6f, 0x95, 0xb8, 0xd9, 0x6b, 0x1f, 0xfe, 0xae, 0x16, 0x57, 0x86, 0x3a, 0xac,
	0x1f, 0xef, 0x55, 0xfb, 0xe7, 0x0f, 0x50, 0xfd, 0xa3, 0x9d, 0x52, 0xbf, 0xef, 0x48, 0x73, 0x3d,
	0x76, 0x59, 0x7f, 0x86, 0x67, 0xa9, 0x88, 0xfc, 0x5a, 0x89, 0x8b, 0xe5, 0xbf, 0x8d, 0x43, 0xb2,
	0xd4, 0x47, 0xff, 0xb3, 0xdb, 0x76, 0xbe, 0xfc, 0xdd, 0xf0, 0xd6, 0x58, 0x3c, 0x39, 0x9b, 0x3e,
	0xb5, 0xc3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x23, 0xef, 0x24, 0x37, 0x84, 0x03, 0x00, 0x00,
}
