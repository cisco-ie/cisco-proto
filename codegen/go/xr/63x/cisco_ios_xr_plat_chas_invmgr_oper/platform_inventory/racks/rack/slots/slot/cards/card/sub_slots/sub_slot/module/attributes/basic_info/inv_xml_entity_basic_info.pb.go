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

package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_slots_slot_cards_card_sub_slots_sub_slot_module_attributes_basic_info

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
	proto.RegisterType((*InvXmlEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.attributes.basic_info.inv_xml_entity_basic_info_KEYS")
	proto.RegisterType((*InvXmlEntityBasicInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.sub_slots.sub_slot.module.attributes.basic_info.inv_xml_entity_basic_info")
}

func init() { proto.RegisterFile("inv_xml_entity_basic_info.proto", fileDescriptor_637d7cec7f49a4b4) }

var fileDescriptor_637d7cec7f49a4b4 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3f, 0x6f, 0xd4, 0x40,
	0x10, 0xc5, 0x75, 0x21, 0x44, 0x64, 0x03, 0x12, 0xb1, 0x04, 0x72, 0x0a, 0xc8, 0x29, 0x34, 0x91,
	0x22, 0xad, 0x94, 0x3b, 0xfe, 0x85, 0x1e, 0x1a, 0xa4, 0x14, 0x07, 0x14, 0x54, 0xa3, 0xb5, 0x3d,
	0x26, 0x23, 0xf6, 0x8f, 0x35, 0xb3, 0x3e, 0xe2, 0x2f, 0xc0, 0x67, 0xa6, 0x44, 0xbb, 0x77, 0x67,
	0x27, 0x45, 0x9a, 0xd9, 0xd9, 0xdf, 0x7b, 0xfb, 0x66, 0x8a, 0x55, 0xa7, 0xe4, 0xd7, 0x70, 0xeb,
	0x2c, 0xa0, 0x8f, 0x14, 0x07, 0xa8, 0x8c, 0x50, 0x0d, 0xe4, 0xdb, 0xa0, 0x3b, 0x0e, 0x31, 0x14,
	0x7f, 0x67, 0x35, 0x49, 0x1d, 0x80, 0x82, 0xc0, 0x2d, 0x43, 0x67, 0x4d, 0x84, 0xfa, 0xc6, 0x08,
	0x90, 0x5f, 0xbb, 0x5f, 0x0c, 0xa1, 0x43, 0xd6, 0x89, 0xb6, 0x81, 0x5d, 0x82, 0xe8, 0x63, 0xe0,
	0x41, 0xb3, 0xa9, 0x7f, 0x4b, 0xae, 0x5a, 0x6c, 0x88, 0x92, 0xab, 0xae, 0x0d, 0x37, 0x92, 0xab,
	0x96, 0xbe, 0x82, 0xad, 0xb2, 0xed, 0xb4, 0x0b, 0x4d, 0x6f, 0x51, 0x9b, 0x18, 0x99, 0xaa, 0x3e,
	0xa2, 0xe8, 0x69, 0x9d, 0xb3, 0x41, 0xbd, 0x7e, 0x70, 0x57, 0xf8, 0xfa, 0xf9, 0xe7, 0xb7, 0xa2,
	0x50, 0xfb, 0xde, 0x38, 0x2c, 0x67, 0xf3, 0xd9, 0xf9, 0xe1, 0x2a, 0xf7, 0xc5, 0x0b, 0x75, 0x90,
	0x4e, 0xb8, 0x2c, 0xf7, 0x32, 0x7d, 0x9c, 0x6e, 0x97, 0x23, 0x5e, 0x94, 0x8f, 0x26, 0xbc, 0x18,
	0xf1, 0xb2, 0xdc, 0x9f, 0xf0, 0xf2, 0xec, 0xdf, 0x9e, 0x3a, 0x79, 0x70, 0xf6, 0x38, 0x76, 0x71,
	0x67, 0xec, 0x5c, 0x1d, 0x35, 0x28, 0x35, 0x53, 0x17, 0x29, 0xf8, 0x72, 0x99, 0xa5, 0xbb, 0xa8,
	0x78, 0xa5, 0x94, 0x0b, 0x0d, 0x5a, 0xc8, 0x6f, 0xdf, 0x66, 0xc3, 0x61, 0x26, 0xd7, 0x29, 0xe0,
	0x42, 0x1d, 0xdf, 0x18, 0x6e, 0xfe, 0x18, 0x46, 0x60, 0x5c, 0x93, 0xa4, 0x98, 0x77, 0xd9, 0xf5,
	0x7c, 0x27, 0xac, 0xb6, 0xbc, 0x78, 0xa3, 0x9e, 0x09, 0x32, 0x19, 0x0b, 0xbe, 0x77, 0x15, 0x72,
	0xf9, 0x3e, 0x1b, 0x9f, 0x6e, 0xe0, 0x75, 0x66, 0x29, 0xb1, 0x25, 0x76, 0xf7, 0x13, 0x3f, 0x6c,
	0x12, 0x77, 0xc2, 0x98, 0x78, 0xa1, 0x8e, 0x25, 0xb4, 0xf1, 0xbe, 0xf9, 0xe3, 0xc6, 0xbc, 0x13,
	0x46, 0xf3, 0xa9, 0x3a, 0x5a, 0xa3, 0x6f, 0x02, 0x43, 0x1c, 0x3a, 0x2c, 0xaf, 0xb2, 0x4d, 0x6d,
	0xd0, 0xf7, 0xa1, 0xc3, 0xe2, 0x4a, 0x9d, 0x90, 0x40, 0x4b, 0x68, 0x1b, 0x60, 0xec, 0xac, 0xa9,
	0xd1, 0x54, 0x16, 0xa1, 0xf7, 0x14, 0xcb, 0x4f, 0xf3, 0xd9, 0xf9, 0x93, 0xd5, 0x4b, 0x92, 0x2f,
	0x49, 0x5f, 0x4d, 0xf2, 0x0f, 0x4f, 0xb1, 0x3a, 0xc8, 0xbf, 0x70, 0xf9, 0x3f, 0x00, 0x00, 0xff,
	0xff, 0x3d, 0x9e, 0x9e, 0x88, 0xa8, 0x02, 0x00, 0x00,
}
