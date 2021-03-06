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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_attributes_inv_asset_bag

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
	proto.RegisterType((*InvPhyEntityAssetInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.attributes.inv_asset_bag.inv_phy_entity_asset_info_KEYS")
	proto.RegisterType((*InvPhyEntityAssetInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.attributes.inv_asset_bag.inv_phy_entity_asset_info")
}

func init() { proto.RegisterFile("inv_phy_entity_asset_info.proto", fileDescriptor_e4a1ef5b8d4a67fa) }

var fileDescriptor_e4a1ef5b8d4a67fa = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd2, 0xcf, 0x4b, 0xe3, 0x40,
	0x14, 0x07, 0x70, 0xba, 0x74, 0x0b, 0x3b, 0x7b, 0x0b, 0x2c, 0x64, 0x97, 0xd2, 0x96, 0x2e, 0x2c,
	0x7b, 0x31, 0x50, 0xeb, 0xaf, 0xa3, 0xa2, 0x15, 0xa5, 0xea, 0xa1, 0x3d, 0x79, 0x1a, 0x26, 0x71,
	0x92, 0x3e, 0xcc, 0xbc, 0x89, 0x6f, 0x26, 0xd1, 0xfc, 0x25, 0xfe, 0xbb, 0x92, 0x31, 0x2d, 0x89,
	0x6d, 0x2f, 0x6d, 0xf2, 0xf8, 0x7e, 0x3f, 0x8f, 0x07, 0x61, 0x43, 0xc0, 0x82, 0x67, 0xab, 0x92,
	0x4b, 0xb4, 0x60, 0x4b, 0x2e, 0x8c, 0x91, 0x96, 0x03, 0xc6, 0x3a, 0xc8, 0x48, 0x5b, 0xed, 0x2d,
	0x22, 0x30, 0x91, 0xe6, 0xa0, 0x0d, 0x7f, 0x23, 0x0e, 0x58, 0xa8, 0x84, 0xb8, 0xce, 0x24, 0x05,
	0x80, 0x85, 0x44, 0xab, 0xa9, 0x0c, 0x48, 0x44, 0xcf, 0xc6, 0xfd, 0x06, 0x9f, 0x4c, 0x60, 0x52,
	0x6d, 0x03, 0x61, 0x2d, 0x41, 0x98, 0x5b, 0x69, 0xaa, 0x6c, 0x4d, 0x87, 0x22, 0x19, 0xcf, 0xd9,
	0x60, 0xef, 0x5a, 0x3e, 0x9f, 0x3d, 0x2e, 0x3d, 0x8f, 0x75, 0x51, 0x28, 0xe9, 0x77, 0x46, 0x9d,
	0xff, 0x3f, 0x16, 0xee, 0xd9, 0xfb, 0xc5, 0x7a, 0xd5, 0x3f, 0x9f, 0xf8, 0xdf, 0xdc, 0xf4, 0x7b,
	0xf5, 0x36, 0x19, 0xbf, 0x77, 0xd9, 0xef, 0xbd, 0x9a, 0x37, 0x64, 0x3f, 0x33, 0x41, 0x96, 0x63,
	0xae, 0x42, 0x49, 0xfe, 0xa1, 0x6b, 0xb2, 0x6a, 0xf4, 0xe0, 0x26, 0xde, 0x39, 0xeb, 0x2b, 0x81,
	0x79, 0x2c, 0x22, 0x9b, 0x93, 0x24, 0xd7, 0x55, 0x61, 0x5a, 0xae, 0x1b, 0x53, 0xd7, 0xf8, 0xd3,
	0xcc, 0x5c, 0xd4, 0x91, 0x5a, 0xb8, 0x62, 0x83, 0xdd, 0x02, 0xc9, 0x02, 0x0c, 0x68, 0xf4, 0x8f,
	0x9c, 0xd1, 0xdf, 0x65, 0x2c, 0xea, 0x8c, 0x77, 0xc3, 0x46, 0x2d, 0x25, 0x06, 0x52, 0xaf, 0x82,
	0x24, 0x87, 0xa7, 0xea, 0xaa, 0x18, 0x24, 0xf9, 0xc7, 0xce, 0x69, 0x6d, 0xbb, 0xae, 0x63, 0xb7,
	0x9b, 0xd4, 0x96, 0x64, 0x74, 0x6c, 0xbf, 0x4a, 0x27, 0xdb, 0xd2, 0xb2, 0x8e, 0x35, 0xa4, 0x15,
	0x9b, 0xb4, 0xa4, 0x48, 0x2b, 0xa5, 0x91, 0xa7, 0x02, 0x93, 0x5c, 0x24, 0x92, 0xcb, 0x97, 0x1c,
	0x32, 0x25, 0xd1, 0x36, 0xe9, 0x53, 0x47, 0x1f, 0x34, 0x8b, 0x97, 0xae, 0x77, 0x57, 0xd7, 0x66,
	0xeb, 0x56, 0x63, 0xd3, 0x92, 0xfd, 0xd3, 0x04, 0x09, 0xa0, 0x48, 0x1b, 0x6a, 0xfb, 0x0c, 0x4b,
	0x80, 0x89, 0x7f, 0xe6, 0xf8, 0xbf, 0xeb, 0xf4, 0x06, 0xbb, 0x6f, 0x9e, 0xe2, 0xa2, 0x61, 0xcf,
	0x7d, 0xc1, 0xd3, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x5b, 0xe5, 0x6d, 0xe4, 0x02, 0x00,
	0x00,
}
