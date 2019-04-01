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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_attributes_inv_asset_bag

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
	proto.RegisterType((*InvPhyEntityAssetInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.inv_asset_bag.inv_phy_entity_asset_info_KEYS")
	proto.RegisterType((*InvPhyEntityAssetInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.inv_asset_bag.inv_phy_entity_asset_info")
}

func init() { proto.RegisterFile("inv_phy_entity_asset_info.proto", fileDescriptor_e4a1ef5b8d4a67fa) }

var fileDescriptor_e4a1ef5b8d4a67fa = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x86, 0x15, 0x48, 0x83, 0x30, 0xbb, 0x91, 0x90, 0x0c, 0xaa, 0xda, 0xaa, 0x48, 0x88, 0x0d,
	0x23, 0x65, 0x2e, 0x81, 0x25, 0x08, 0x8a, 0x40, 0x5c, 0x16, 0xcd, 0x8a, 0x95, 0xe5, 0x19, 0x3c,
	0xd3, 0x23, 0x62, 0x7b, 0x38, 0xf6, 0x04, 0xf2, 0x08, 0x3c, 0x01, 0x6f, 0xc3, 0xb3, 0x21, 0x1f,
	0x26, 0x73, 0x69, 0x9b, 0xcd, 0xc9, 0xc9, 0xa7, 0xff, 0xff, 0x6c, 0x2f, 0x86, 0x9d, 0x82, 0xd9,
	0x8a, 0xe6, 0x6a, 0x27, 0x94, 0xf1, 0xe0, 0x77, 0x42, 0x3a, 0xa7, 0xbc, 0x00, 0x53, 0xd9, 0xb8,
	0x41, 0xeb, 0x6d, 0xf4, 0x7b, 0x56, 0x82, 0x2b, 0xad, 0x00, 0xeb, 0xc4, 0x2f, 0x14, 0x60, 0xb6,
	0xba, 0x46, 0x61, 0x1b, 0x85, 0x31, 0x98, 0xad, 0x32, 0xde, 0xe2, 0x2e, 0x46, 0x59, 0x7e, 0x77,
	0x34, 0xe3, 0xff, 0x9e, 0xd8, 0x6d, 0xac, 0x8f, 0xbd, 0x83, 0xa5, 0xa3, 0x19, 0x46, 0x42, 0x6b,
	0x12, 0x46, 0x4a, 0x6b, 0x1a, 0x46, 0x46, 0x6b, 0x16, 0x46, 0x4e, 0x6b, 0x1e, 0x4b, 0xef, 0x11,
	0x8a, 0xd6, 0x2b, 0x17, 0x8e, 0xe8, 0xae, 0x54, 0xc8, 0xfa, 0xfc, 0xef, 0x8c, 0x9d, 0x1c, 0xbc,
	0xaf, 0xf8, 0x78, 0xf1, 0x75, 0x1d, 0x45, 0x6c, 0x6e, 0xa4, 0x56, 0x7c, 0x76, 0x36, 0x7b, 0x76,
	0xff, 0x92, 0xf6, 0xe8, 0x21, 0x5b, 0x84, 0x5f, 0xb1, 0xe4, 0x77, 0x88, 0x1e, 0x85, 0x7f, 0xcb,
	0x1e, 0x27, 0xfc, 0xee, 0x80, 0x93, 0x1e, 0xa7, 0x7c, 0x3e, 0xe0, 0xb4, 0xc7, 0x19, 0x3f, 0x1a,
	0x70, 0xd6, 0xe3, 0x9c, 0x2f, 0x06, 0x9c, 0xf7, 0x78, 0xc5, 0xef, 0x0d, 0x78, 0x75, 0xfe, 0x67,
	0xce, 0x1e, 0x1d, 0x7c, 0x40, 0x74, 0xca, 0x1e, 0x34, 0x12, 0xbd, 0x30, 0xad, 0x2e, 0x14, 0xf2,
	0x84, 0x9a, 0x2c, 0xa0, 0x2f, 0x44, 0xa2, 0x57, 0xec, 0x58, 0x4b, 0xd3, 0x56, 0xb2, 0xf4, 0x2d,
	0x2a, 0xa4, 0xae, 0x2e, 0x36, 0xbb, 0x7d, 0x23, 0xa5, 0xc6, 0xe3, 0x71, 0xe6, 0x75, 0x17, 0xe9,
	0x0c, 0x6f, 0xd9, 0xc9, 0xed, 0x06, 0x54, 0x5b, 0x70, 0x60, 0x0d, 0xcf, 0xc8, 0x71, 0x7c, 0x9b,
	0xe3, 0xb2, 0xcb, 0x44, 0xef, 0xd9, 0xd9, 0xc4, 0x52, 0x01, 0xea, 0x9f, 0x12, 0x95, 0x80, 0x6f,
	0xe1, 0x55, 0x15, 0x28, 0xe4, 0x39, 0x79, 0x26, 0xa7, 0xbd, 0xeb, 0x62, 0x1f, 0xfa, 0xd4, 0x0d,
	0x93, 0xb3, 0x95, 0xbf, 0x6e, 0x5a, 0xdd, 0x34, 0xad, 0xbb, 0xd8, 0xc8, 0x74, 0xc5, 0x96, 0x13,
	0x53, 0x69, 0xb5, 0xb6, 0x46, 0x6c, 0xa4, 0xa9, 0x5b, 0x59, 0x2b, 0xa1, 0x7e, 0xb4, 0xd0, 0x68,
	0x65, 0xfc, 0x58, 0xfd, 0x82, 0xd4, 0xcf, 0xc7, 0xc5, 0x37, 0xd4, 0xfb, 0xd4, 0xd5, 0x2e, 0xf6,
	0xad, 0xd1, 0x49, 0x6b, 0xf6, 0xd4, 0x22, 0xd4, 0x60, 0xe4, 0x66, 0x64, 0x9d, 0x3e, 0xc3, 0x23,
	0x98, 0x9a, 0xbf, 0x24, 0xfd, 0x93, 0x7d, 0xba, 0x97, 0x7d, 0x1e, 0x3f, 0x85, 0xa2, 0xc5, 0x82,
	0xbe, 0xb6, 0xf4, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0xb6, 0xa1, 0x45, 0x90, 0x03, 0x00,
	0x00,
}
