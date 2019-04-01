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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_attributes_inv_asset_bag

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
	proto.RegisterType((*InvPhyEntityAssetInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.attributes.inv_asset_bag.inv_phy_entity_asset_info_KEYS")
	proto.RegisterType((*InvPhyEntityAssetInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.attributes.inv_asset_bag.inv_phy_entity_asset_info")
}

func init() { proto.RegisterFile("inv_phy_entity_asset_info.proto", fileDescriptor_e4a1ef5b8d4a67fa) }

var fileDescriptor_e4a1ef5b8d4a67fa = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd2, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0x06, 0x70, 0x5c, 0x5c, 0x43, 0xb7, 0x37, 0x41, 0x41, 0x2d, 0xc6, 0x36, 0x2e, 0x94, 0x5e,
	0x2a, 0x70, 0x9d, 0x7f, 0xc7, 0x84, 0xc4, 0x21, 0xc1, 0x49, 0x0e, 0x36, 0x04, 0x72, 0x5a, 0x56,
	0xca, 0x48, 0x1e, 0xa2, 0xdd, 0x55, 0x66, 0x57, 0x72, 0xf4, 0x24, 0x79, 0xdd, 0xa0, 0x45, 0x36,
	0x52, 0x6c, 0x5f, 0x6c, 0x69, 0xf8, 0xbe, 0xdf, 0x30, 0x20, 0x36, 0x44, 0x55, 0xf0, 0x6c, 0x55,
	0x72, 0x50, 0x16, 0x6d, 0xc9, 0x85, 0x31, 0x60, 0x39, 0xaa, 0x58, 0x07, 0x19, 0x69, 0xab, 0xbd,
	0xc7, 0x08, 0x4d, 0xa4, 0x39, 0x6a, 0xc3, 0xdf, 0x88, 0xa3, 0x2a, 0x64, 0x42, 0x5c, 0x67, 0x40,
	0x01, 0xaa, 0x02, 0x94, 0xd5, 0x54, 0x06, 0x24, 0xa2, 0x17, 0xe3, 0x7e, 0x83, 0x4c, 0xaf, 0x81,
	0xcc, 0x0a, 0xd2, 0x38, 0x30, 0xa9, 0xb6, 0x81, 0xb0, 0x96, 0x30, 0xcc, 0x2d, 0x98, 0x2a, 0x5f,
	0xf3, 0xa1, 0x48, 0xc6, 0x73, 0x36, 0x38, 0xb8, 0x9a, 0xcf, 0x67, 0x4f, 0x4b, 0xcf, 0x63, 0x5d,
	0x25, 0x24, 0xf8, 0x9d, 0x51, 0xe7, 0xef, 0xb7, 0x85, 0x7b, 0xf6, 0x7e, 0xb0, 0x5e, 0xf5, 0xcf,
	0x27, 0xfe, 0x17, 0x37, 0xfd, 0x5a, 0xbd, 0x4d, 0xc6, 0xef, 0x5d, 0xf6, 0xf3, 0xa0, 0xe6, 0x0d,
	0xd9, 0xf7, 0x4c, 0x90, 0xe5, 0x2a, 0x97, 0x21, 0x90, 0xff, 0xdf, 0x35, 0x59, 0x35, 0x7a, 0x70,
	0x13, 0xef, 0x9c, 0xf5, 0xa5, 0x50, 0x79, 0x2c, 0x22, 0x9b, 0x13, 0x90, 0xeb, 0xca, 0x30, 0x2d,
	0x37, 0x8d, 0xa9, 0x6b, 0xfc, 0x6a, 0x66, 0x2e, 0xea, 0x48, 0x2d, 0x5c, 0xb1, 0xc1, 0x7e, 0x81,
	0xa0, 0x40, 0x83, 0x5a, 0xf9, 0x47, 0xce, 0xe8, 0xef, 0x33, 0x16, 0x75, 0xc6, 0xbb, 0x61, 0xa3,
	0x96, 0x12, 0x23, 0xc9, 0xb5, 0x20, 0xe0, 0xf8, 0x5c, 0x5d, 0x15, 0x23, 0x90, 0x7f, 0xec, 0x9c,
	0xd6, 0xb6, 0xeb, 0x3a, 0x76, 0xbb, 0x4d, 0xed, 0x48, 0x46, 0xc7, 0xf6, 0xb3, 0x74, 0xb2, 0x2b,
	0x2d, 0xeb, 0x58, 0x43, 0x5a, 0xb1, 0x49, 0x4b, 0x8a, 0xb4, 0x94, 0x5a, 0xf1, 0x54, 0xa8, 0x24,
	0x17, 0x09, 0x70, 0x78, 0xcd, 0x31, 0x93, 0xa0, 0x6c, 0x93, 0x3e, 0x75, 0xf4, 0xbf, 0x66, 0xf1,
	0xd2, 0xf5, 0xee, 0xea, 0xda, 0x6c, 0xd3, 0x6a, 0x6c, 0x5a, 0xb2, 0x3f, 0x9a, 0x30, 0x41, 0x25,
	0xd2, 0x86, 0xda, 0x3e, 0xc3, 0x12, 0xaa, 0xc4, 0x3f, 0x73, 0xfc, 0xef, 0x4d, 0x7a, 0x8b, 0xdd,
	0x37, 0x4f, 0x71, 0xd1, 0xb0, 0xe7, 0xbe, 0xe2, 0xe9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a,
	0x0d, 0xaa, 0xbc, 0xe8, 0x02, 0x00, 0x00,
}
