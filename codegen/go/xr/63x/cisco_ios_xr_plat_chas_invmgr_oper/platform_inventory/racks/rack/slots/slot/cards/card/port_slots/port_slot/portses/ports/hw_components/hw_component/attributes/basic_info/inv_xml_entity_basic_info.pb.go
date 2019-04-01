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
// source: inv_xml_entity_basic_info.proto

package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_slots_slot_cards_card_port_slots_port_slot_portses_ports_hw_components_hw_component_attributes_basic_info

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

type InvXmlEntityBasicInfo_KEYS struct {
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

func (m *InvXmlEntityBasicInfo_KEYS) Reset()         { *m = InvXmlEntityBasicInfo_KEYS{} }
func (m *InvXmlEntityBasicInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*InvXmlEntityBasicInfo_KEYS) ProtoMessage()    {}
func (*InvXmlEntityBasicInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_637d7cec7f49a4b4, []int{0}
}

func (m *InvXmlEntityBasicInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvXmlEntityBasicInfo_KEYS.Unmarshal(m, b)
}
func (m *InvXmlEntityBasicInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvXmlEntityBasicInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *InvXmlEntityBasicInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvXmlEntityBasicInfo_KEYS.Merge(m, src)
}
func (m *InvXmlEntityBasicInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_InvXmlEntityBasicInfo_KEYS.Size(m)
}
func (m *InvXmlEntityBasicInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InvXmlEntityBasicInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InvXmlEntityBasicInfo_KEYS proto.InternalMessageInfo

func (m *InvXmlEntityBasicInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvXmlEntityBasicInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *InvXmlEntityBasicInfo_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *InvXmlEntityBasicInfo_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
	}
	return ""
}

func (m *InvXmlEntityBasicInfo_KEYS) GetName_4() string {
	if m != nil {
		return m.Name_4
	}
	return ""
}

func (m *InvXmlEntityBasicInfo_KEYS) GetName_5() string {
	if m != nil {
		return m.Name_5
	}
	return ""
}

type InvXmlEntityBasicInfo struct {
	Name                   string   `protobuf:"bytes,50,opt,name=name,proto3" json:"name,omitempty"`
	Description            string   `protobuf:"bytes,51,opt,name=description,proto3" json:"description,omitempty"`
	ModelName              string   `protobuf:"bytes,52,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	HardwareRevision       string   `protobuf:"bytes,53,opt,name=hardware_revision,json=hardwareRevision,proto3" json:"hardware_revision,omitempty"`
	SerialNumber           string   `protobuf:"bytes,54,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	FirmwareRevision       string   `protobuf:"bytes,55,opt,name=firmware_revision,json=firmwareRevision,proto3" json:"firmware_revision,omitempty"`
	SoftwareRevision       string   `protobuf:"bytes,56,opt,name=software_revision,json=softwareRevision,proto3" json:"software_revision,omitempty"`
	VendorType             string   `protobuf:"bytes,57,opt,name=vendor_type,json=vendorType,proto3" json:"vendor_type,omitempty"`
	IsFieldReplaceableUnit bool     `protobuf:"varint,58,opt,name=is_field_replaceable_unit,json=isFieldReplaceableUnit,proto3" json:"is_field_replaceable_unit,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *InvXmlEntityBasicInfo) Reset()         { *m = InvXmlEntityBasicInfo{} }
func (m *InvXmlEntityBasicInfo) String() string { return proto.CompactTextString(m) }
func (*InvXmlEntityBasicInfo) ProtoMessage()    {}
func (*InvXmlEntityBasicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_637d7cec7f49a4b4, []int{1}
}

func (m *InvXmlEntityBasicInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvXmlEntityBasicInfo.Unmarshal(m, b)
}
func (m *InvXmlEntityBasicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvXmlEntityBasicInfo.Marshal(b, m, deterministic)
}
func (m *InvXmlEntityBasicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvXmlEntityBasicInfo.Merge(m, src)
}
func (m *InvXmlEntityBasicInfo) XXX_Size() int {
	return xxx_messageInfo_InvXmlEntityBasicInfo.Size(m)
}
func (m *InvXmlEntityBasicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InvXmlEntityBasicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InvXmlEntityBasicInfo proto.InternalMessageInfo

func (m *InvXmlEntityBasicInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetHardwareRevision() string {
	if m != nil {
		return m.HardwareRevision
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetFirmwareRevision() string {
	if m != nil {
		return m.FirmwareRevision
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetSoftwareRevision() string {
	if m != nil {
		return m.SoftwareRevision
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetVendorType() string {
	if m != nil {
		return m.VendorType
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetIsFieldReplaceableUnit() bool {
	if m != nil {
		return m.IsFieldReplaceableUnit
	}
	return false
}

func init() {
	proto.RegisterType((*InvXmlEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.port_slots.port_slot.portses.ports.hw_components.hw_component.attributes.basic_info.inv_xml_entity_basic_info_KEYS")
	proto.RegisterType((*InvXmlEntityBasicInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.port_slots.port_slot.portses.ports.hw_components.hw_component.attributes.basic_info.inv_xml_entity_basic_info")
}

func init() { proto.RegisterFile("inv_xml_entity_basic_info.proto", fileDescriptor_637d7cec7f49a4b4) }

var fileDescriptor_637d7cec7f49a4b4 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6b, 0xd5, 0x40,
	0x14, 0xc5, 0x49, 0x6d, 0x1f, 0x76, 0xaa, 0x60, 0x03, 0xca, 0x74, 0xa1, 0x7d, 0xd4, 0x4d, 0xa1,
	0x30, 0xd0, 0xf7, 0x47, 0xad, 0x7b, 0xdd, 0x08, 0x5d, 0x3c, 0x75, 0xe1, 0xea, 0x32, 0x99, 0xdc,
	0xf8, 0x2e, 0x26, 0x33, 0xc3, 0x9d, 0x79, 0xaf, 0xcd, 0x27, 0x12, 0xfc, 0x84, 0x2e, 0x25, 0x13,
	0x5f, 0xd2, 0x16, 0xba, 0xb9, 0xb9, 0xf7, 0x77, 0x4e, 0x0e, 0x27, 0x10, 0x71, 0x4a, 0x76, 0x0b,
	0xb7, 0x4d, 0x0d, 0x68, 0x23, 0xc5, 0x16, 0x0a, 0x1d, 0xc8, 0x00, 0xd9, 0xca, 0x29, 0xcf, 0x2e,
	0xba, 0xfc, 0x4f, 0x66, 0x28, 0x18, 0x07, 0xe4, 0x02, 0xdc, 0x32, 0xf8, 0x5a, 0x47, 0x30, 0x6b,
	0x1d, 0x80, 0xec, 0xb6, 0xf9, 0xc9, 0xe0, 0x3c, 0xb2, 0xea, 0x68, 0xe5, 0xb8, 0xe9, 0x20, 0xda,
	0xe8, 0xb8, 0x55, 0xac, 0xcd, 0xaf, 0x90, 0xa6, 0x0a, 0xb5, 0x8b, 0x21, 0x4d, 0x65, 0x34, 0x97,
	0x21, 0x4d, 0xe5, 0x1d, 0x47, 0xe8, 0xa5, 0x61, 0x4d, 0x5b, 0xc0, 0x9e, 0x04, 0xb5, 0xbe, 0x01,
	0xe3, 0x1a, 0xef, 0x2c, 0xda, 0x07, 0x97, 0xd2, 0x31, 0x32, 0x15, 0x9b, 0x88, 0x41, 0x8d, 0x95,
	0xcf, 0x7e, 0x67, 0xe2, 0xcd, 0xa3, 0x1f, 0x04, 0x5f, 0x3e, 0xfd, 0xf8, 0x9a, 0xe7, 0x62, 0xdf,
	0xea, 0x06, 0x65, 0x36, 0xcd, 0xce, 0x0f, 0x57, 0x69, 0xcf, 0x5f, 0x8a, 0x49, 0xf7, 0x84, 0x4b,
	0xb9, 0x97, 0xe8, 0x41, 0x77, 0x5d, 0x0e, 0x78, 0x26, 0x9f, 0x8c, 0x78, 0x36, 0xe0, 0xb9, 0xdc,
	0x1f, 0xf1, 0x7c, 0xc0, 0x0b, 0x79, 0x30, 0xe2, 0xc5, 0x80, 0x97, 0x72, 0x32, 0xe2, 0xe5, 0xd9,
	0xdf, 0x3d, 0x71, 0xf2, 0x68, 0xd3, 0xa1, 0xe4, 0xec, 0x4e, 0xc9, 0xa9, 0x38, 0x2a, 0x31, 0x18,
	0x26, 0x1f, 0xc9, 0x59, 0x39, 0x4f, 0xd2, 0x5d, 0x94, 0xbf, 0x16, 0xa2, 0x71, 0x25, 0xd6, 0x90,
	0xde, 0x5d, 0x24, 0xc3, 0x61, 0x22, 0xd7, 0x5d, 0xc0, 0x85, 0x38, 0x5e, 0x6b, 0x2e, 0x6f, 0x34,
	0x23, 0x30, 0x6e, 0x29, 0x74, 0x31, 0xcb, 0xe4, 0x7a, 0xb1, 0x13, 0x56, 0xff, 0x79, 0xfe, 0x56,
	0x3c, 0x0f, 0xc8, 0xa4, 0x6b, 0xb0, 0x9b, 0xa6, 0x40, 0x96, 0xef, 0x92, 0xf1, 0x59, 0x0f, 0xaf,
	0x13, 0xeb, 0x12, 0x2b, 0xe2, 0xe6, 0x7e, 0xe2, 0xfb, 0x3e, 0x71, 0x27, 0x0c, 0x89, 0x17, 0xe2,
	0x38, 0xb8, 0x2a, 0xde, 0x37, 0x7f, 0xe8, 0xcd, 0x3b, 0x61, 0x30, 0x9f, 0x8a, 0xa3, 0x2d, 0xda,
	0xd2, 0x31, 0xc4, 0xd6, 0xa3, 0xbc, 0x4a, 0x36, 0xd1, 0xa3, 0x6f, 0xad, 0xc7, 0xfc, 0x4a, 0x9c,
	0x50, 0x80, 0x8a, 0xb0, 0x2e, 0x81, 0xd1, 0xd7, 0xda, 0xa0, 0x2e, 0x6a, 0x84, 0x8d, 0xa5, 0x28,
	0x3f, 0x4e, 0xb3, 0xf3, 0xa7, 0xab, 0x57, 0x14, 0x3e, 0x77, 0xfa, 0x6a, 0x94, 0xbf, 0x5b, 0x8a,
	0xc5, 0x24, 0xfd, 0xd8, 0xf3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0x2b, 0x60, 0xc0, 0xfb,
	0x02, 0x00, 0x00,
}
