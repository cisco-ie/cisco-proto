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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_attributes_inv_asset_bag

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
	proto.RegisterType((*InvPhyEntityAssetInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.inv_asset_bag.inv_phy_entity_asset_info_KEYS")
	proto.RegisterType((*InvPhyEntityAssetInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.inv_asset_bag.inv_phy_entity_asset_info")
}

func init() { proto.RegisterFile("inv_phy_entity_asset_info.proto", fileDescriptor_e4a1ef5b8d4a67fa) }

var fileDescriptor_e4a1ef5b8d4a67fa = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4f, 0x8b, 0xd4, 0x40,
	0x10, 0xc5, 0x19, 0x9d, 0x5d, 0xb0, 0xbd, 0x05, 0x84, 0x56, 0x96, 0xdd, 0x65, 0x05, 0xf1, 0x62,
	0x60, 0x92, 0xf1, 0xcf, 0x51, 0xd1, 0x15, 0xc5, 0x3f, 0x87, 0x9d, 0x93, 0xa7, 0xa6, 0x13, 0x3b,
	0xd9, 0xc2, 0xe9, 0xea, 0x58, 0x5d, 0x19, 0xcd, 0xb7, 0xf0, 0xe6, 0xd7, 0x95, 0x14, 0x99, 0x4c,
	0xc6, 0xdd, 0xb9, 0x54, 0x2a, 0x8f, 0xf7, 0x7e, 0x55, 0x75, 0x68, 0x75, 0x06, 0xb8, 0x31, 0xcd,
	0x75, 0x67, 0x1c, 0x32, 0x70, 0x67, 0x6c, 0x8c, 0x8e, 0x0d, 0x60, 0x15, 0xd2, 0x86, 0x02, 0x87,
	0x84, 0x4b, 0x88, 0x65, 0x30, 0x10, 0xa2, 0xf9, 0x4d, 0x06, 0x70, 0xe3, 0x6b, 0x32, 0xa1, 0x71,
	0x94, 0x02, 0x6e, 0x1c, 0x72, 0xa0, 0x2e, 0x25, 0x5b, 0xfe, 0x88, 0x52, 0xd3, 0xca, 0x22, 0x93,
	0xed, 0xd2, 0xb8, 0x0e, 0x9c, 0x72, 0x84, 0x45, 0x94, 0xda, 0x97, 0x4c, 0xda, 0xac, 0x2f, 0xb9,
	0xb4, 0x79, 0x6a, 0x99, 0x09, 0x8a, 0x96, 0x5d, 0xec, 0x69, 0xc3, 0xf0, 0xc2, 0xd6, 0x17, 0x7f,
	0x66, 0xea, 0xf4, 0xe0, 0x66, 0xe6, 0xd3, 0xe5, 0xb7, 0x55, 0x92, 0xa8, 0x39, 0x5a, 0xef, 0xf4,
	0xec, 0x7c, 0xf6, 0xf4, 0xde, 0x95, 0xf4, 0xc9, 0x03, 0x75, 0xdc, 0x7f, 0xcd, 0x42, 0xdf, 0x11,
	0xf5, 0xa8, 0xff, 0x5b, 0x8c, 0x72, 0xa6, 0xef, 0xee, 0xe4, 0x6c, 0x94, 0x73, 0x3d, 0xdf, 0xc9,
	0xf9, 0x28, 0x2f, 0xf5, 0xd1, 0x4e, 0x5e, 0x5e, 0xfc, 0x9d, 0xab, 0x87, 0x07, 0x57, 0x4a, 0xce,
	0xd4, 0xfd, 0xc6, 0x12, 0x1b, 0x6c, 0x7d, 0xe1, 0x48, 0x67, 0x92, 0x54, 0xbd, 0xf4, 0x55, 0x94,
	0xe4, 0xb5, 0x3a, 0xf1, 0x16, 0xdb, 0xca, 0x96, 0xdc, 0x92, 0x23, 0xc9, 0xfa, 0x62, 0xdd, 0x6d,
	0x13, 0xb9, 0x24, 0x1e, 0x4d, 0x3d, 0x6f, 0x06, 0xcb, 0x40, 0x78, 0xa7, 0x4e, 0x6f, 0x27, 0x90,
	0xdb, 0x40, 0x84, 0x80, 0x7a, 0x29, 0x8c, 0x93, 0xdb, 0x18, 0x57, 0x83, 0x27, 0xf9, 0xa0, 0xce,
	0xf7, 0x28, 0x15, 0x90, 0xff, 0x65, 0xc9, 0x19, 0xf8, 0xde, 0x5f, 0x55, 0x81, 0x23, 0xfd, 0x5c,
	0x38, 0x7b, 0xd3, 0xde, 0x0f, 0xb6, 0x8f, 0xa3, 0xeb, 0x06, 0x29, 0x86, 0x8a, 0xff, 0x27, 0xbd,
	0xb8, 0x49, 0x5a, 0x0d, 0xb6, 0x09, 0xe9, 0x5a, 0x2d, 0xf6, 0x48, 0x65, 0xf0, 0x3e, 0xa0, 0x59,
	0x5b, 0xac, 0x5b, 0x5b, 0x3b, 0xe3, 0x7e, 0xb6, 0xd0, 0x78, 0x87, 0x3c, 0x45, 0xbf, 0x14, 0xf4,
	0xb3, 0x69, 0xf0, 0xad, 0xe4, 0x3e, 0x0f, 0xb1, 0xcb, 0x6d, 0x6a, 0x32, 0x69, 0xa5, 0x9e, 0x04,
	0x82, 0x1a, 0xd0, 0xae, 0x27, 0xd4, 0xfd, 0x33, 0x98, 0x00, 0x6b, 0xfd, 0x4a, 0xf0, 0x8f, 0xb7,
	0xee, 0x11, 0xf6, 0x65, 0x7a, 0x8a, 0x58, 0x8b, 0x63, 0x79, 0x29, 0xf9, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x82, 0x97, 0x0e, 0xf2, 0x4c, 0x03, 0x00, 0x00,
}