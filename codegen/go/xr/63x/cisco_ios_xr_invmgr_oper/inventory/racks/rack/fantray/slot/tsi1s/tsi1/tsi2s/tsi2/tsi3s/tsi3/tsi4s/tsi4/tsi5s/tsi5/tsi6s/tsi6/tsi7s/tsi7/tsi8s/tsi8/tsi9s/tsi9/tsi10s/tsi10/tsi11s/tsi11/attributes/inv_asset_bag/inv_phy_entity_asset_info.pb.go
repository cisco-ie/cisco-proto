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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_tsi10s_tsi10_tsi11s_tsi11_attributes_inv_asset_bag

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
	Name_8               string   `protobuf:"bytes,9,opt,name=name_8,json=name8,proto3" json:"name_8,omitempty"`
	Name_9               string   `protobuf:"bytes,10,opt,name=name_9,json=name9,proto3" json:"name_9,omitempty"`
	Name_10              string   `protobuf:"bytes,11,opt,name=name_10,json=name10,proto3" json:"name_10,omitempty"`
	Name_11              string   `protobuf:"bytes,12,opt,name=name_11,json=name11,proto3" json:"name_11,omitempty"`
	Name_12              string   `protobuf:"bytes,13,opt,name=name_12,json=name12,proto3" json:"name_12,omitempty"`
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

func (m *InvPhyEntityAssetInfo_KEYS) GetName_8() string {
	if m != nil {
		return m.Name_8
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_9() string {
	if m != nil {
		return m.Name_9
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_10() string {
	if m != nil {
		return m.Name_10
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_11() string {
	if m != nil {
		return m.Name_11
	}
	return ""
}

func (m *InvPhyEntityAssetInfo_KEYS) GetName_12() string {
	if m != nil {
		return m.Name_12
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
	proto.RegisterType((*InvPhyEntityAssetInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.inv_asset_bag.inv_phy_entity_asset_info_KEYS")
	proto.RegisterType((*InvPhyEntityAssetInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.inv_asset_bag.inv_phy_entity_asset_info")
}

func init() { proto.RegisterFile("inv_phy_entity_asset_info.proto", fileDescriptor_e4a1ef5b8d4a67fa) }

var fileDescriptor_e4a1ef5b8d4a67fa = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xd5, 0xd1, 0x75, 0xcc, 0x83, 0x4b, 0x24, 0x84, 0x41, 0xd3, 0x36, 0x0d, 0x09, 0x71,
	0x21, 0x6a, 0x92, 0xb6, 0x49, 0x6f, 0x20, 0x18, 0x02, 0xf1, 0xe7, 0xb0, 0x9e, 0x38, 0x59, 0x4e,
	0x70, 0x32, 0x8b, 0xc6, 0x0e, 0xb6, 0x53, 0xc8, 0x27, 0xe1, 0x5b, 0x21, 0x3e, 0x12, 0xf2, 0xbb,
	0xd4, 0x4e, 0xd9, 0x76, 0x79, 0xfa, 0xf4, 0xe7, 0xe7, 0x7d, 0x9c, 0x37, 0x87, 0xa0, 0x53, 0x2e,
	0x36, 0xa4, 0xb9, 0xea, 0x08, 0x13, 0x86, 0x9b, 0x8e, 0x50, 0xad, 0x99, 0x21, 0x5c, 0x94, 0x32,
	0x6c, 0x94, 0x34, 0x32, 0xf8, 0x3b, 0x2a, 0xb8, 0x2e, 0x24, 0xe1, 0x52, 0x93, 0x5f, 0x8a, 0x70,
	0xb1, 0xa9, 0x2b, 0x45, 0x64, 0xc3, 0x54, 0xc8, 0xc5, 0x86, 0x09, 0x23, 0x55, 0x17, 0x2a, 0x5a,
	0x7c, 0xd7, 0xa0, 0x61, 0x49, 0x85, 0x51, 0xb4, 0x0b, 0xf5, 0x5a, 0x9a, 0xd0, 0x68, 0x1e, 0x69,
	0x50, 0x2b, 0x31, 0xd8, 0xd8, 0x4a, 0x02, 0x36, 0xb1, 0x32, 0x03, 0x3b, 0xb3, 0x32, 0x07, 0x3b,
	0xb7, 0xb2, 0x00, 0xbb, 0xb0, 0x92, 0x82, 0x4d, 0xad, 0x64, 0x60, 0x33, 0x2b, 0x4b, 0xb0, 0x4b,
	0x28, 0x9f, 0x5e, 0xdf, 0x31, 0x05, 0xed, 0x2f, 0x8c, 0x42, 0x6a, 0x8c, 0xe2, 0x79, 0x6b, 0x98,
	0xb6, 0x4f, 0xda, 0x6f, 0x96, 0xd3, 0xea, 0xfc, 0xcf, 0x1e, 0x3a, 0xb9, 0x73, 0x6d, 0xf2, 0xf1,
	0xe2, 0xeb, 0x2a, 0x08, 0xd0, 0x58, 0xd0, 0x9a, 0xe1, 0xd1, 0xd9, 0xe8, 0xc5, 0xe1, 0x25, 0xf8,
	0xe0, 0x11, 0x9a, 0xd8, 0x5f, 0x12, 0xe1, 0x3d, 0xa0, 0xfb, 0xf6, 0x5f, 0xe4, 0x70, 0x8c, 0xef,
	0x79, 0x1c, 0x3b, 0x9c, 0xe0, 0xb1, 0xc7, 0x89, 0xc3, 0x33, 0xbc, 0xef, 0xf1, 0xcc, 0xe1, 0x39,
	0x9e, 0x78, 0x3c, 0x77, 0x78, 0x81, 0x0f, 0x3c, 0x5e, 0x38, 0x9c, 0xe2, 0xfb, 0x1e, 0xa7, 0x0e,
	0x67, 0xf8, 0xd0, 0xe3, 0xcc, 0xe1, 0x25, 0x46, 0x1e, 0x2f, 0x83, 0xc7, 0xe8, 0xe0, 0x7a, 0x9d,
	0x29, 0x3e, 0x02, 0x0e, 0xa9, 0x68, 0xea, 0x0f, 0x22, 0xfc, 0x60, 0x70, 0x10, 0xf9, 0x83, 0x18,
	0x3f, 0x1c, 0x1c, 0xc4, 0xe7, 0xbf, 0xc7, 0xe8, 0xc9, 0x9d, 0x2f, 0x34, 0x38, 0x45, 0x47, 0x0d,
	0x55, 0x86, 0x88, 0xb6, 0xce, 0x99, 0xc2, 0x31, 0x8c, 0x22, 0x8b, 0xbe, 0x00, 0x09, 0x5e, 0xa1,
	0xe3, 0x9a, 0x8a, 0xb6, 0xa4, 0x85, 0x69, 0x15, 0x53, 0x30, 0x5b, 0xe7, 0xeb, 0x6e, 0x3b, 0x91,
	0xc0, 0xc4, 0xd3, 0x61, 0xe6, 0x75, 0x1f, 0xe9, 0x1b, 0xde, 0xa2, 0x93, 0xdb, 0x1b, 0x14, 0xdb,
	0x70, 0xcd, 0xa5, 0xc0, 0x33, 0xe8, 0x38, 0xbe, 0xad, 0xe3, 0xb2, 0xcf, 0x04, 0xef, 0xd1, 0xd9,
	0x4e, 0x4b, 0xc9, 0x55, 0xfd, 0x93, 0x2a, 0x46, 0xf8, 0x37, 0xbb, 0x55, 0xc9, 0x99, 0xc2, 0x73,
	0xe8, 0xd9, 0xb9, 0xed, 0x5d, 0x1f, 0xfb, 0xe0, 0x52, 0x37, 0x9a, 0xb4, 0x2c, 0xcd, 0xff, 0x4d,
	0x8b, 0x9b, 0x4d, 0xab, 0x3e, 0x36, 0x68, 0xba, 0x42, 0xd1, 0x4e, 0x53, 0x21, 0xeb, 0x5a, 0x0a,
	0xb2, 0xa6, 0xa2, 0x6a, 0x69, 0xc5, 0x08, 0xfb, 0xd1, 0xf2, 0xa6, 0x66, 0xc2, 0x0c, 0xab, 0x53,
	0xa8, 0x7e, 0x39, 0x1c, 0x7c, 0x03, 0x73, 0x9f, 0xfa, 0xb1, 0x8b, 0xed, 0xd4, 0xe0, 0xa6, 0x15,
	0x7a, 0x2e, 0x15, 0xaf, 0xb8, 0xa0, 0xeb, 0x41, 0xeb, 0xee, 0x1a, 0x46, 0x71, 0x51, 0xe1, 0x0c,
	0xea, 0x9f, 0x6d, 0xd3, 0xae, 0xec, 0xf3, 0x70, 0x15, 0x88, 0xe6, 0x13, 0xf8, 0x88, 0x24, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xe1, 0x61, 0x2e, 0x67, 0x04, 0x00, 0x00,
}
