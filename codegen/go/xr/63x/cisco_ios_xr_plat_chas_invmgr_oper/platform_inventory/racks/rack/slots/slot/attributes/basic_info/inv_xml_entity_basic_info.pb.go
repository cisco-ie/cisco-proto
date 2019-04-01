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

package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_slots_slot_attributes_basic_info

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
	proto.RegisterType((*InvXmlEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.attributes.basic_info.inv_xml_entity_basic_info_KEYS")
	proto.RegisterType((*InvXmlEntityBasicInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.attributes.basic_info.inv_xml_entity_basic_info")
}

func init() { proto.RegisterFile("inv_xml_entity_basic_info.proto", fileDescriptor_637d7cec7f49a4b4) }

var fileDescriptor_637d7cec7f49a4b4 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbd, 0xcf, 0xd3, 0x30,
	0x10, 0xc6, 0xd5, 0x0a, 0x2a, 0xea, 0x82, 0x44, 0x23, 0x81, 0xd2, 0x01, 0x5a, 0x95, 0xa5, 0x52,
	0xa5, 0x48, 0xa5, 0x7c, 0x95, 0x1d, 0x96, 0x4a, 0x1d, 0x02, 0x0c, 0x4c, 0x27, 0x27, 0xb9, 0xd0,
	0x13, 0x8e, 0x1d, 0x9d, 0xdd, 0xd0, 0xfc, 0xe7, 0x8c, 0xaf, 0xec, 0xbc, 0xfd, 0x1a, 0xba, 0x5c,
	0xe2, 0xdf, 0xf3, 0xf8, 0x67, 0x0f, 0x16, 0x53, 0xd2, 0x0d, 0x1c, 0x2b, 0x05, 0xa8, 0x1d, 0xb9,
	0x16, 0x32, 0x69, 0x29, 0x07, 0xd2, 0xa5, 0x49, 0x6a, 0x36, 0xce, 0x44, 0x32, 0x27, 0x9b, 0x1b,
	0x20, 0x63, 0xe1, 0xc8, 0x50, 0x2b, 0xe9, 0x20, 0xdf, 0x4b, 0x0b, 0xa4, 0x9b, 0xea, 0x0f, 0x83,
	0xa9, 0x91, 0x13, 0x4f, 0x4b, 0xc3, 0x95, 0x87, 0xa8, 0x9d, 0xe1, 0x36, 0x61, 0x99, 0xff, 0xb5,
	0x61, 0x26, 0x56, 0x19, 0x67, 0xc3, 0x4c, 0xa4, 0x73, 0x4c, 0xd9, 0xc1, 0xa1, 0x4d, 0x2e, 0x07,
	0xcd, 0xb7, 0xe2, 0xed, 0xdd, 0x5b, 0xc0, 0xf6, 0xdb, 0xef, 0x1f, 0x51, 0x24, 0x9e, 0x68, 0x59,
	0x61, 0xdc, 0x9b, 0xf5, 0x16, 0xc3, 0x34, 0xfc, 0x47, 0xaf, 0xc4, 0xc0, 0x7f, 0x61, 0x15, 0xf7,
	0x03, 0x7d, 0xea, 0x57, 0xab, 0xf9, 0xff, 0xbe, 0x98, 0xdc, 0xb5, 0x9d, 0x45, 0xef, 0xaf, 0x44,
	0x33, 0x31, 0x2a, 0xd0, 0xe6, 0x4c, 0xb5, 0x23, 0xa3, 0xe3, 0x75, 0x88, 0xae, 0x51, 0xf4, 0x46,
	0x88, 0xca, 0x14, 0xa8, 0x20, 0xec, 0xfd, 0x10, 0x0a, 0xc3, 0x40, 0x76, 0x5e, 0xb0, 0x14, 0xe3,
	0xbd, 0xe4, 0xe2, 0x9f, 0x64, 0x04, 0xc6, 0x86, 0xac, 0xd7, 0x7c, 0x0c, 0xad, 0x97, 0xa7, 0x20,
	0x7d, 0xe4, 0xd1, 0x3b, 0xf1, 0xc2, 0x22, 0x93, 0x54, 0xa0, 0x0f, 0x55, 0x86, 0x1c, 0x7f, 0x0a,
	0xc5, 0xe7, 0x1d, 0xdc, 0x05, 0xe6, 0x8d, 0x25, 0x71, 0x75, 0x6b, 0xfc, 0xdc, 0x19, 0x4f, 0xc1,
	0xd9, 0xb8, 0x14, 0x63, 0x6b, 0x4a, 0x77, 0x5b, 0xfe, 0xd2, 0x95, 0x4f, 0xc1, 0xb9, 0x3c, 0x15,
	0xa3, 0x06, 0x75, 0x61, 0x18, 0x5c, 0x5b, 0x63, 0xbc, 0x09, 0x35, 0xd1, 0xa1, 0x9f, 0x6d, 0x8d,
	0xd1, 0x46, 0x4c, 0xc8, 0x42, 0x49, 0xa8, 0x0a, 0x60, 0xac, 0x95, 0xcc, 0x51, 0x66, 0x0a, 0xe1,
	0xa0, 0xc9, 0xc5, 0x5f, 0x67, 0xbd, 0xc5, 0xb3, 0xf4, 0x35, 0xd9, 0xef, 0x3e, 0x4f, 0x2f, 0xf1,
	0x2f, 0x4d, 0x2e, 0x1b, 0x84, 0x17, 0xb3, 0x7e, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x69, 0x4a,
	0x34, 0x54, 0x02, 0x00, 0x00,
}
