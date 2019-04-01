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
// source: inv_phy_entity_asset_info.proto

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_attributes_inv_asset_bag

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

type InvPhyEntityAssetInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Name_1               string   `protobuf:"bytes,2,opt,name=name_1,json=name1,proto3" json:"name_1,omitempty"`
	Name_2               string   `protobuf:"bytes,3,opt,name=name_2,json=name2,proto3" json:"name_2,omitempty"`
	Name_3               string   `protobuf:"bytes,4,opt,name=name_3,json=name3,proto3" json:"name_3,omitempty"`
	Name_4               string   `protobuf:"bytes,5,opt,name=name_4,json=name4,proto3" json:"name_4,omitempty"`
	Name_5               string   `protobuf:"bytes,6,opt,name=name_5,json=name5,proto3" json:"name_5,omitempty"`
	Name_6               string   `protobuf:"bytes,7,opt,name=name_6,json=name6,proto3" json:"name_6,omitempty"`
	Name_7               string   `protobuf:"bytes,8,opt,name=name_7,json=name7,proto3" json:"name_7,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvPhyEntityAssetInfo_KEYS) Reset()         { *m = InvPhyEntityAssetInfo_KEYS{} }
func (m *InvPhyEntityAssetInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*InvPhyEntityAssetInfo_KEYS) ProtoMessage()    {}
func (*InvPhyEntityAssetInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4a1ef5b8d4a67fa, []int{0}
}

func (m *InvPhyEntityAssetInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvPhyEntityAssetInfo_KEYS.Unmarshal(m, b)
}
func (m *InvPhyEntityAssetInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvPhyEntityAssetInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *InvPhyEntityAssetInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvPhyEntityAssetInfo_KEYS.Merge(m, src)
}
func (m *InvPhyEntityAssetInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_InvPhyEntityAssetInfo_KEYS.Size(m)
}
func (m *InvPhyEntityAssetInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InvPhyEntityAssetInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InvPhyEntityAssetInfo_KEYS proto.InternalMessageInfo

func (m *InvPhyEntityAssetInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_4() string {
	if m != nil {
		return m.Name_4
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_5() string {
	if m != nil {
		return m.Name_5
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_6() string {
	if m != nil {
		return m.Name_6
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_7() string {
	if m != nil {
		return m.Name_7
	}
	return ""
}

type InvPhyEntityAssetInfo struct {
	PartNumber                                    string   `protobuf:"bytes,50,opt,name=part_number,json=partNumber,proto3" json:"part_number,omitempty"`
	ManufacturerAssemblyNumber                    string   `protobuf:"bytes,51,opt,name=manufacturer_assembly_number,json=manufacturerAssemblyNumber,proto3" json:"manufacturer_assembly_number,omitempty"`
	ManufacturerAssemblyRevision                  string   `protobuf:"bytes,52,opt,name=manufacturer_assembly_revision,json=manufacturerAssemblyRevision,proto3" json:"manufacturer_assembly_revision,omitempty"`
	ManufacturerFirmwareIdentifier                string   `protobuf:"bytes,53,opt,name=manufacturer_firmware_identifier,json=manufacturerFirmwareIdentifier,proto3" json:"manufacturer_firmware_identifier,omitempty"`
	ManufacturerSoftwareIdentifier                string   `protobuf:"bytes,54,opt,name=manufacturer_software_identifier,json=manufacturerSoftwareIdentifier,proto3" json:"manufacturer_software_identifier,omitempty"`
	ManufacturerCommonLanguageEquipmentIdentifier string   `protobuf:"bytes,55,opt,name=manufacturer_common_language_equipment_identifier,json=manufacturerCommonLanguageEquipmentIdentifier,proto3" json:"manufacturer_common_language_equipment_identifier,omitempty"`
	OriginalEquipmentManufacturerString           string   `protobuf:"bytes,56,opt,name=original_equipment_manufacturer_string,json=originalEquipmentManufacturerString,proto3" json:"original_equipment_manufacturer_string,omitempty"`
	XXX_NoUnkeyedLiteral                          struct{} `json:"-"`
	XXX_unrecognized                              []byte   `json:"-"`
	XXX_sizecache                                 int32    `json:"-"`
}

func (m *InvPhyEntityAssetInfo) Reset()         { *m = InvPhyEntityAssetInfo{} }
func (m *InvPhyEntityAssetInfo) String() string { return proto.CompactTextString(m) }
func (*InvPhyEntityAssetInfo) ProtoMessage()    {}
func (*InvPhyEntityAssetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4a1ef5b8d4a67fa, []int{1}
}

func (m *InvPhyEntityAssetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvPhyEntityAssetInfo.Unmarshal(m, b)
}
func (m *InvPhyEntityAssetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvPhyEntityAssetInfo.Marshal(b, m, deterministic)
}
func (m *InvPhyEntityAssetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvPhyEntityAssetInfo.Merge(m, src)
}
func (m *InvPhyEntityAssetInfo) XXX_Size() int {
	return xxx_messageInfo_InvPhyEntityAssetInfo.Size(m)
}
func (m *InvPhyEntityAssetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InvPhyEntityAssetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InvPhyEntityAssetInfo proto.InternalMessageInfo

func (m *InvPhyEntityAssetInfo) GetPartNumber() string {
	if m != nil {
		return m.PartNumber
	}
	return ""
}

func (m *InvPhyEntityAssetInfo) GetManufacturerAssemblyNumber() string {
	if m != nil {
		return m.ManufacturerAssemblyNumber
	}
	return ""
}

func (m *InvPhyEntityAssetInfo) GetManufacturerAssemblyRevision() string {
	if m != nil {
		return m.ManufacturerAssemblyRevision
	}
	return ""
}

func (m *InvPhyEntityAssetInfo) GetManufacturerFirmwareIdentifier() string {
	if m != nil {
		return m.ManufacturerFirmwareIdentifier
	}
	return ""
}

func (m *InvPhyEntityAssetInfo) GetManufacturerSoftwareIdentifier() string {
	if m != nil {
		return m.ManufacturerSoftwareIdentifier
	}
	return ""
}

func (m *InvPhyEntityAssetInfo) GetManufacturerCommonLanguageEquipmentIdentifier() string {
	if m != nil {
		return m.ManufacturerCommonLanguageEquipmentIdentifier
	}
	return ""
}

func (m *InvPhyEntityAssetInfo) GetOriginalEquipmentManufacturerString() string {
	if m != nil {
		return m.OriginalEquipmentManufacturerString
	}
	return ""
}

func init() {
	proto.RegisterType((*InvPhyEntityAssetInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.attributes.inv_asset_bag.inv_phy_entity_asset_info_KEYS")
	proto.RegisterType((*InvPhyEntityAssetInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.attributes.inv_asset_bag.inv_phy_entity_asset_info")
}

func init() { proto.RegisterFile("inv_phy_entity_asset_info.proto", fileDescriptor_e4a1ef5b8d4a67fa) }

var fileDescriptor_e4a1ef5b8d4a67fa = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0x55, 0xe8, 0x14, 0x30, 0xbb, 0x48, 0x48, 0x06, 0x8d, 0x66, 0x46, 0x83, 0x84, 0xd8,
	0x10, 0xa9, 0xb9, 0x95, 0x25, 0x08, 0x06, 0x81, 0xb8, 0x2c, 0xa6, 0x2b, 0x56, 0x96, 0x13, 0x9c,
	0xf4, 0x88, 0xc4, 0x0e, 0xc7, 0x4e, 0x87, 0x3e, 0x09, 0xbc, 0x18, 0xef, 0x83, 0x7c, 0x94, 0x26,
	0xe9, 0x5c, 0x36, 0xa7, 0xa7, 0x9f, 0xfe, 0xff, 0xb3, 0xbd, 0x08, 0x3b, 0x05, 0xbd, 0x15, 0xed,
	0x66, 0x27, 0x94, 0x76, 0xe0, 0x76, 0x42, 0x5a, 0xab, 0x9c, 0x00, 0x5d, 0x9a, 0xb0, 0x45, 0xe3,
	0x4c, 0xf0, 0x77, 0x56, 0x80, 0x2d, 0x8c, 0x00, 0x63, 0xc5, 0x6f, 0x14, 0xa0, 0xb7, 0x4d, 0x85,
	0xc2, 0xb4, 0x0a, 0x43, 0xd0, 0x5b, 0xa5, 0x9d, 0xc1, 0x5d, 0x88, 0xb2, 0xf8, 0x69, 0x69, 0x86,
	0xad, 0xb9, 0x52, 0x68, 0x37, 0xaa, 0x2e, 0x43, 0x5b, 0x1b, 0x17, 0x3a, 0x0b, 0x4b, 0x4b, 0xd3,
	0x8f, 0x88, 0xd6, 0xc8, 0x8f, 0x98, 0xd6, 0xd8, 0x8f, 0x84, 0xd6, 0xc4, 0x8f, 0x94, 0xd6, 0xd4,
	0x8f, 0x8c, 0xd6, 0x2c, 0x94, 0xce, 0x21, 0xe4, 0x9d, 0x53, 0xd6, 0x9f, 0xd8, 0xdf, 0x30, 0x97,
	0xd5, 0xf9, 0xbf, 0x19, 0x3b, 0xb9, 0xf3, 0xfa, 0xe2, 0xf3, 0xc5, 0xf7, 0x75, 0x10, 0xb0, 0xb9,
	0x96, 0x8d, 0xe2, 0xb3, 0xb3, 0xd9, 0xcb, 0x47, 0x97, 0xb4, 0x07, 0x4f, 0xd8, 0xc2, 0xff, 0x8a,
	0x25, 0xbf, 0x47, 0xf4, 0xc8, 0xff, 0x5b, 0x0e, 0x38, 0xe2, 0xf7, 0x47, 0x1c, 0x0d, 0x38, 0xe6,
	0xf3, 0x11, 0xc7, 0x03, 0x4e, 0xf8, 0xd1, 0x88, 0x93, 0x01, 0xa7, 0x7c, 0x31, 0xe2, 0x74, 0xc0,
	0x19, 0x7f, 0x30, 0xe2, 0x6c, 0xc0, 0x2b, 0xfe, 0x70, 0xc4, 0xab, 0xf3, 0x3f, 0x73, 0xf6, 0xf4,
	0xce, 0x77, 0x05, 0xa7, 0xec, 0x71, 0x2b, 0xd1, 0x09, 0xdd, 0x35, 0xb9, 0x42, 0x1e, 0x51, 0x93,
	0x79, 0xf4, 0x8d, 0x48, 0xf0, 0x86, 0x1d, 0x37, 0x52, 0x77, 0xa5, 0x2c, 0x5c, 0x87, 0x0a, 0xa9,
	0xdb, 0xe4, 0xf5, 0x6e, 0xdf, 0x88, 0xa9, 0xf1, 0x6c, 0x9a, 0x79, 0xdb, 0x47, 0x7a, 0xc3, 0x7b,
	0x76, 0x72, 0xbb, 0x01, 0xd5, 0x16, 0x2c, 0x18, 0xcd, 0x13, 0x72, 0x1c, 0xdf, 0xe6, 0xb8, 0xec,
	0x33, 0xc1, 0x47, 0x76, 0x76, 0x60, 0x29, 0x01, 0x9b, 0x2b, 0x89, 0x4a, 0xc0, 0x0f, 0xff, 0xaa,
	0x12, 0x14, 0xf2, 0x94, 0x3c, 0x07, 0xa7, 0x7d, 0xe8, 0x63, 0x9f, 0x86, 0xd4, 0x0d, 0x93, 0x35,
	0xa5, 0xbb, 0x6e, 0xca, 0x6e, 0x9a, 0xd6, 0x7d, 0x6c, 0x62, 0xda, 0xb0, 0xe5, 0x81, 0xa9, 0x30,
	0x4d, 0x63, 0xb4, 0xa8, 0xa5, 0xae, 0x3a, 0x59, 0x29, 0xa1, 0x7e, 0x75, 0xd0, 0x36, 0x4a, 0xbb,
	0xa9, 0x7a, 0x45, 0xea, 0x57, 0xd3, 0xe2, 0x3b, 0xea, 0x7d, 0xe9, 0x6b, 0x17, 0xfb, 0xd6, 0xe4,
	0xa4, 0x35, 0x7b, 0x61, 0x10, 0x2a, 0xd0, 0xb2, 0x9e, 0x58, 0x0f, 0x9f, 0xe1, 0x10, 0x74, 0xc5,
	0x5f, 0x93, 0xfe, 0xf9, 0x3e, 0x3d, 0xc8, 0xbe, 0x4e, 0x9f, 0x42, 0xd1, 0x7c, 0x41, 0xdf, 0x64,
	0xfc, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xc5, 0x97, 0x78, 0xb6, 0x03, 0x00, 0x00,
}