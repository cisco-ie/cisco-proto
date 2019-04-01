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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_attributes_inv_asset_bag

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
	proto.RegisterType((*InvPhyEntityAssetInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.inv_asset_bag.inv_phy_entity_asset_info_KEYS")
	proto.RegisterType((*InvPhyEntityAssetInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.inv_asset_bag.inv_phy_entity_asset_info")
}

func init() { proto.RegisterFile("inv_phy_entity_asset_info.proto", fileDescriptor_e4a1ef5b8d4a67fa) }

var fileDescriptor_e4a1ef5b8d4a67fa = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4f, 0x6b, 0xd5, 0x40,
	0x14, 0xc5, 0xa9, 0x3e, 0x0b, 0x8e, 0xbb, 0x80, 0x30, 0x4a, 0x69, 0x4b, 0x05, 0x71, 0x63, 0xe0,
	0xe5, 0xf9, 0x6f, 0xa9, 0x68, 0x45, 0xf1, 0xcf, 0xa2, 0x6f, 0xe5, 0x6a, 0x98, 0xc4, 0x49, 0x7a,
	0x31, 0x73, 0x27, 0xde, 0xb9, 0x89, 0xe6, 0x93, 0xf8, 0x75, 0x25, 0x97, 0xbc, 0x34, 0x69, 0xfb,
	0x36, 0x37, 0x37, 0x87, 0x73, 0x7e, 0x27, 0x77, 0x11, 0x75, 0x02, 0xd8, 0x99, 0xe6, 0xb2, 0x37,
	0x0e, 0x19, 0xb8, 0x37, 0x36, 0x46, 0xc7, 0x06, 0xb0, 0x0c, 0x69, 0x43, 0x81, 0x43, 0x02, 0x05,
	0xc4, 0x22, 0x18, 0x08, 0xd1, 0xfc, 0x25, 0x03, 0xd8, 0xf9, 0x8a, 0x4c, 0x68, 0x1c, 0xa5, 0x80,
	0x9d, 0x43, 0x0e, 0xd4, 0xa7, 0x64, 0x8b, 0x5f, 0x51, 0x66, 0x5a, 0x5a, 0x64, 0xb2, 0x7d, 0x1a,
	0xeb, 0xc0, 0x29, 0x47, 0x58, 0x47, 0x99, 0xc3, 0xc8, 0x64, 0xcd, 0x52, 0xcb, 0x4c, 0x90, 0xb7,
	0xec, 0xe2, 0x80, 0x18, 0x1b, 0x73, 0x5b, 0x9d, 0xf5, 0xea, 0x78, 0xef, 0xd7, 0x98, 0x2f, 0xe7,
	0x3f, 0xb6, 0x49, 0xa2, 0x56, 0x68, 0xbd, 0xd3, 0x07, 0xa7, 0x07, 0xcf, 0xee, 0x5f, 0xc8, 0x9e,
	0x3c, 0x54, 0x87, 0xc3, 0xd3, 0xac, 0xf5, 0x1d, 0x51, 0xef, 0x0d, 0x6f, 0xeb, 0x49, 0xce, 0xf4,
	0xdd, 0x2b, 0x39, 0x9b, 0xe4, 0x8d, 0x5e, 0x5d, 0xc9, 0x9b, 0xb3, 0x7f, 0x2b, 0xf5, 0x68, 0x6f,
	0x77, 0x72, 0xa2, 0x1e, 0x34, 0x96, 0xd8, 0x60, 0xeb, 0x73, 0x47, 0x3a, 0x93, 0xa4, 0x1a, 0xa4,
	0xef, 0xa2, 0x24, 0x6f, 0xd5, 0x91, 0xb7, 0xd8, 0x96, 0xb6, 0xe0, 0x96, 0x1c, 0x49, 0xd6, 0xe7,
	0x75, 0xbf, 0x4b, 0x6c, 0x24, 0xf1, 0x78, 0xee, 0x79, 0x37, 0x5a, 0x46, 0xc2, 0x07, 0x75, 0x7c,
	0x3b, 0x81, 0x5c, 0x07, 0x11, 0x02, 0xea, 0x17, 0xc2, 0x38, 0xba, 0x8d, 0x71, 0x31, 0x7a, 0x92,
	0x4f, 0xea, 0x74, 0x41, 0x29, 0x81, 0xfc, 0x1f, 0x4b, 0xce, 0xc0, 0xcf, 0xe1, 0xaa, 0x12, 0x1c,
	0xe9, 0x97, 0xc2, 0x59, 0xb4, 0x7d, 0x1c, 0x6d, 0x9f, 0x27, 0xd7, 0x0d, 0x52, 0x0c, 0x25, 0x5f,
	0x27, 0xbd, 0xba, 0x49, 0xda, 0x8e, 0xb6, 0x19, 0xe9, 0x52, 0xad, 0x17, 0xa4, 0x22, 0x78, 0x1f,
	0xd0, 0xd4, 0x16, 0xab, 0xd6, 0x56, 0xce, 0xb8, 0xdf, 0x2d, 0x34, 0xde, 0x21, 0xcf, 0xd1, 0xaf,
	0x05, 0xfd, 0x7c, 0x1e, 0x7c, 0x2f, 0xb9, 0xaf, 0x63, 0xec, 0x7c, 0x97, 0x9a, 0x35, 0x6d, 0xd5,
	0xd3, 0x40, 0x50, 0x01, 0xda, 0x7a, 0x46, 0x5d, 0x9e, 0xc1, 0x04, 0x58, 0xe9, 0x37, 0x82, 0x7f,
	0xb2, 0x73, 0x4f, 0xb0, 0x6f, 0xf3, 0x53, 0xc4, 0x9a, 0x1f, 0xca, 0x6f, 0xb0, 0xf9, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0xd8, 0x01, 0x22, 0x5b, 0x29, 0x03, 0x00, 0x00,
}