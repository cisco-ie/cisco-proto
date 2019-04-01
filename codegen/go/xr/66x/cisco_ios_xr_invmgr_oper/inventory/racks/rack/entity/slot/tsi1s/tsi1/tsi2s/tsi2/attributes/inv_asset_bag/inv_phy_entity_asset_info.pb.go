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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_attributes_inv_asset_bag

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
	proto.RegisterType((*InvPhyEntityAssetInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.inv_asset_bag.inv_phy_entity_asset_info_KEYS")
	proto.RegisterType((*InvPhyEntityAssetInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.inv_asset_bag.inv_phy_entity_asset_info")
}

func init() { proto.RegisterFile("inv_phy_entity_asset_info.proto", fileDescriptor_e4a1ef5b8d4a67fa) }

var fileDescriptor_e4a1ef5b8d4a67fa = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0xef, 0x93, 0x40,
	0x10, 0xc5, 0xf3, 0xd5, 0xda, 0xc4, 0xf5, 0x46, 0x62, 0xb2, 0x9a, 0xa6, 0x6d, 0x6a, 0x62, 0xbc,
	0x48, 0x02, 0xf8, 0xeb, 0xa8, 0xd1, 0x1a, 0x8d, 0x3f, 0x0e, 0xed, 0xc9, 0xd3, 0x66, 0xc1, 0x85,
	0x4e, 0x64, 0x77, 0x71, 0x76, 0x41, 0xf9, 0x4b, 0xfc, 0x77, 0x0d, 0x23, 0xa5, 0x60, 0xdb, 0xcb,
	0x30, 0xbc, 0xbc, 0xf7, 0x79, 0xcc, 0x01, 0xb6, 0x02, 0xd3, 0x88, 0xea, 0xd0, 0x0a, 0x65, 0x3c,
	0xf8, 0x56, 0x48, 0xe7, 0x94, 0x17, 0x60, 0x72, 0x1b, 0x56, 0x68, 0xbd, 0x0d, 0x0e, 0x19, 0xb8,
	0xcc, 0x0a, 0xb0, 0x4e, 0xfc, 0x46, 0x01, 0xa6, 0xd1, 0x05, 0x0a, 0x5b, 0x29, 0x0c, 0xc1, 0x34,
	0xca, 0x78, 0x8b, 0x6d, 0x88, 0x32, 0xfb, 0xe1, 0x68, 0x86, 0xff, 0x30, 0xa1, 0x2b, 0xad, 0x0f,
	0xbd, 0x83, 0xc8, 0xd1, 0xec, 0x46, 0x4c, 0x6b, 0x1c, 0x4a, 0xef, 0x11, 0xd2, 0xda, 0x2b, 0xd7,
	0x11, 0xfa, 0xc2, 0x54, 0x16, 0x9b, 0x96, 0x2d, 0xaf, 0x7e, 0x8c, 0xf8, 0xb4, 0xfd, 0xb6, 0x0f,
	0x02, 0x36, 0x33, 0x52, 0x2b, 0x7e, 0xb3, 0xbe, 0x79, 0x72, 0x77, 0x47, 0x7b, 0x70, 0x9f, 0xcd,
	0xbb, 0xa7, 0x88, 0xf8, 0x2d, 0x52, 0xef, 0x74, 0x6f, 0xd1, 0x20, 0xc7, 0xfc, 0xf6, 0x49, 0x8e,
	0x07, 0x39, 0xe1, 0xb3, 0x93, 0x9c, 0x6c, 0xfe, 0xcc, 0xd8, 0x83, 0xab, 0xdd, 0xc1, 0x8a, 0xdd,
	0xab, 0x24, 0x7a, 0x61, 0x6a, 0x9d, 0x2a, 0xe4, 0x31, 0x25, 0x59, 0x27, 0x7d, 0x25, 0x25, 0x78,
	0xcd, 0x16, 0x5a, 0x9a, 0x3a, 0x97, 0x99, 0xaf, 0x51, 0x21, 0x65, 0x75, 0x5a, 0xb6, 0xc7, 0x44,
	0x42, 0x89, 0x87, 0x63, 0xcf, 0x9b, 0xde, 0xd2, 0x13, 0xde, 0xb1, 0xe5, 0x65, 0x02, 0xaa, 0x06,
	0x1c, 0x58, 0xc3, 0x9f, 0x11, 0x63, 0x71, 0x89, 0xb1, 0xeb, 0x3d, 0xc1, 0x07, 0xb6, 0x9e, 0x50,
	0x72, 0x40, 0xfd, 0x4b, 0xa2, 0x12, 0xf0, 0xbd, 0xbb, 0x2a, 0x07, 0x85, 0xfc, 0x39, 0x71, 0x26,
	0x6d, 0xef, 0x7b, 0xdb, 0xc7, 0xc1, 0x75, 0x46, 0x72, 0x36, 0xf7, 0xff, 0x93, 0x5e, 0x9c, 0x93,
	0xf6, 0xbd, 0x6d, 0x44, 0x3a, 0xb0, 0x68, 0x42, 0xca, 0xac, 0xd6, 0xd6, 0x88, 0x52, 0x9a, 0xa2,
	0x96, 0x85, 0x12, 0xea, 0x67, 0x0d, 0x95, 0x56, 0xc6, 0x8f, 0xd1, 0x2f, 0x09, 0xfd, 0x74, 0x1c,
	0x7c, 0x4b, 0xb9, 0xcf, 0x7d, 0x6c, 0x7b, 0x4c, 0x8d, 0x9a, 0xf6, 0xec, 0xb1, 0x45, 0x28, 0xc0,
	0xc8, 0x72, 0x44, 0x9d, 0x9e, 0xe1, 0x11, 0x4c, 0xc1, 0x5f, 0x11, 0xfe, 0xd1, 0xd1, 0x3d, 0xc0,
	0xbe, 0x8c, 0x4f, 0x21, 0x6b, 0x3a, 0xa7, 0xbf, 0x20, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xda,
	0xa0, 0x33, 0xa8, 0x28, 0x03, 0x00, 0x00,
}
