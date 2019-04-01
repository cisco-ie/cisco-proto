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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_attributes_inv_asset_bag

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
	proto.RegisterType((*InvPhyEntityAssetInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.inv_asset_bag.inv_phy_entity_asset_info_KEYS")
	proto.RegisterType((*InvPhyEntityAssetInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.inv_asset_bag.inv_phy_entity_asset_info")
}

func init() { proto.RegisterFile("inv_phy_entity_asset_info.proto", fileDescriptor_e4a1ef5b8d4a67fa) }

var fileDescriptor_e4a1ef5b8d4a67fa = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x86, 0xa9, 0xd6, 0x82, 0xe3, 0x2e, 0x20, 0x44, 0x29, 0x6d, 0xa9, 0x20, 0x6e, 0x0c, 0xa4,
	0xf5, 0x6b, 0xa9, 0x68, 0x45, 0xf1, 0x63, 0xd1, 0xae, 0xdc, 0x38, 0x4c, 0xe2, 0x24, 0x3d, 0x98,
	0x39, 0x13, 0xcf, 0x9c, 0xc4, 0x9b, 0x5f, 0x72, 0xff, 0xee, 0x25, 0x43, 0xda, 0x9b, 0xdc, 0xb6,
	0x9b, 0x49, 0xf2, 0xf2, 0x3e, 0xcf, 0xc9, 0x81, 0x11, 0x73, 0xc0, 0x5a, 0x96, 0xfb, 0x46, 0x6a,
	0x64, 0xe0, 0x46, 0x2a, 0xe7, 0x34, 0x4b, 0xc0, 0xcc, 0x46, 0x25, 0x59, 0xb6, 0xc1, 0xef, 0x14,
	0x5c, 0x6a, 0x25, 0x58, 0x27, 0xaf, 0x48, 0x02, 0xd6, 0x26, 0x27, 0x69, 0x4b, 0x4d, 0x11, 0x60,
	0xad, 0x91, 0x2d, 0x35, 0x11, 0xa9, 0xf4, 0xaf, 0xf3, 0x67, 0x94, 0x29, 0x64, 0x52, 0x4d, 0xe4,
	0x0a, 0xcb, 0x11, 0x3b, 0x88, 0x9d, 0x3f, 0x23, 0xc5, 0x4c, 0x90, 0x54, 0xac, 0x5d, 0xcb, 0x75,
	0x63, 0x12, 0x95, 0x2f, 0x13, 0x31, 0xbb, 0xf8, 0x0b, 0xf2, 0xdb, 0xe6, 0xd7, 0x2e, 0x08, 0xc4,
	0x18, 0x95, 0xd1, 0xe1, 0x68, 0x31, 0x7a, 0xf1, 0x70, 0xeb, 0xdf, 0x83, 0xc7, 0x62, 0xd2, 0x3e,
	0x65, 0x1c, 0xde, 0xf3, 0xe9, 0x83, 0xf6, 0x2b, 0x3e, 0xc6, 0xab, 0xf0, 0xfe, 0x6d, 0xbc, 0x5a,
	0x5e, 0x8f, 0xc5, 0x93, 0x8b, 0x43, 0x82, 0xb9, 0x78, 0x54, 0x2a, 0x62, 0x89, 0x95, 0x49, 0x34,
	0x85, 0x2b, 0x4f, 0x8a, 0x36, 0xfa, 0xe9, 0x93, 0xe0, 0xbd, 0x98, 0x1a, 0x85, 0x55, 0xa6, 0x52,
	0xae, 0x48, 0x93, 0x67, 0x4d, 0x52, 0x34, 0x07, 0x62, 0xed, 0x89, 0xa7, 0xfd, 0xce, 0x87, 0xae,
	0xd2, 0x19, 0x3e, 0x89, 0xd9, 0x79, 0x03, 0xe9, 0x1a, 0x1c, 0x58, 0x0c, 0x5f, 0x79, 0xc7, 0xf4,
	0x9c, 0x63, 0xdb, 0x75, 0x82, 0x2f, 0x62, 0x31, 0xb0, 0x64, 0x40, 0xe6, 0xbf, 0x22, 0x2d, 0xe1,
	0x4f, 0xbb, 0x55, 0x06, 0x9a, 0xc2, 0xd7, 0xde, 0x33, 0x98, 0xf6, 0xb9, 0xab, 0x7d, 0x3d, 0xb6,
	0x4e, 0x4c, 0xce, 0x66, 0x7c, 0xd7, 0xf4, 0xe6, 0xd4, 0xb4, 0xeb, 0x6a, 0x3d, 0xd3, 0x5e, 0xc4,
	0x03, 0x53, 0x6a, 0x8d, 0xb1, 0x28, 0x0b, 0x85, 0x79, 0xa5, 0x72, 0x2d, 0xf5, 0xbf, 0x0a, 0x4a,
	0xa3, 0x91, 0xfb, 0xea, 0xb7, 0x5e, 0xfd, 0xb2, 0x0f, 0x7e, 0xf4, 0xdc, 0xf7, 0x0e, 0xdb, 0x1c,
	0xa8, 0xde, 0xa4, 0x9d, 0x78, 0x6e, 0x09, 0x72, 0x40, 0x55, 0xf4, 0xac, 0xc3, 0x35, 0x98, 0x00,
	0xf3, 0xf0, 0x9d, 0xd7, 0x3f, 0x3b, 0xb4, 0x8f, 0xb2, 0x1f, 0xfd, 0x55, 0x7c, 0x35, 0x99, 0xf8,
	0x4b, 0xbe, 0xbe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x56, 0xce, 0xf7, 0x07, 0x03, 0x00, 0x00,
}
