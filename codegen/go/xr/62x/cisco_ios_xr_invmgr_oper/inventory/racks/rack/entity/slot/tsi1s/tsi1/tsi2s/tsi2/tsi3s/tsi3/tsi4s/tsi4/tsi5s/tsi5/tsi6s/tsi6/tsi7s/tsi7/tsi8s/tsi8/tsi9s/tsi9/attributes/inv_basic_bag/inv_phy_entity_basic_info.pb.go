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
// source: inv_phy_entity_basic_info.proto

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_attributes_inv_basic_bag

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

type InvPhyEntityBasicInfo_KEYS struct {
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvPhyEntityBasicInfo_KEYS) Reset()         { *m = InvPhyEntityBasicInfo_KEYS{} }
func (m *InvPhyEntityBasicInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*InvPhyEntityBasicInfo_KEYS) ProtoMessage()    {}
func (*InvPhyEntityBasicInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_94f1af735311b03a, []int{0}
}

func (m *InvPhyEntityBasicInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvPhyEntityBasicInfo_KEYS.Unmarshal(m, b)
}
func (m *InvPhyEntityBasicInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvPhyEntityBasicInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *InvPhyEntityBasicInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvPhyEntityBasicInfo_KEYS.Merge(m, src)
}
func (m *InvPhyEntityBasicInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_InvPhyEntityBasicInfo_KEYS.Size(m)
}
func (m *InvPhyEntityBasicInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InvPhyEntityBasicInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InvPhyEntityBasicInfo_KEYS proto.InternalMessageInfo

func (m *InvPhyEntityBasicInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvPhyEntityBasicInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *InvPhyEntityBasicInfo_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *InvPhyEntityBasicInfo_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
	}
	return ""
}

func (m *InvPhyEntityBasicInfo_KEYS) GetName_4() string {
	if m != nil {
		return m.Name_4
	}
	return ""
}

func (m *InvPhyEntityBasicInfo_KEYS) GetName_5() string {
	if m != nil {
		return m.Name_5
	}
	return ""
}

func (m *InvPhyEntityBasicInfo_KEYS) GetName_6() string {
	if m != nil {
		return m.Name_6
	}
	return ""
}

func (m *InvPhyEntityBasicInfo_KEYS) GetName_7() string {
	if m != nil {
		return m.Name_7
	}
	return ""
}

func (m *InvPhyEntityBasicInfo_KEYS) GetName_8() string {
	if m != nil {
		return m.Name_8
	}
	return ""
}

func (m *InvPhyEntityBasicInfo_KEYS) GetName_9() string {
	if m != nil {
		return m.Name_9
	}
	return ""
}

func (m *InvPhyEntityBasicInfo_KEYS) GetName_10() string {
	if m != nil {
		return m.Name_10
	}
	return ""
}

type InvPhyEntityBasicInfo struct {
	Description                      string   `protobuf:"bytes,50,opt,name=description,proto3" json:"description,omitempty"`
	VendorType                       string   `protobuf:"bytes,51,opt,name=vendor_type,json=vendorType,proto3" json:"vendor_type,omitempty"`
	Name                             string   `protobuf:"bytes,52,opt,name=name,proto3" json:"name,omitempty"`
	HardwareRevision                 string   `protobuf:"bytes,53,opt,name=hardware_revision,json=hardwareRevision,proto3" json:"hardware_revision,omitempty"`
	FirmwareRevision                 string   `protobuf:"bytes,54,opt,name=firmware_revision,json=firmwareRevision,proto3" json:"firmware_revision,omitempty"`
	SoftwareRevision                 string   `protobuf:"bytes,55,opt,name=software_revision,json=softwareRevision,proto3" json:"software_revision,omitempty"`
	ChipHardwareRevision             string   `protobuf:"bytes,56,opt,name=chip_hardware_revision,json=chipHardwareRevision,proto3" json:"chip_hardware_revision,omitempty"`
	SerialNumber                     string   `protobuf:"bytes,57,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	ManufacturerName                 string   `protobuf:"bytes,58,opt,name=manufacturer_name,json=manufacturerName,proto3" json:"manufacturer_name,omitempty"`
	ModelName                        string   `protobuf:"bytes,59,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	AssetIdStr                       string   `protobuf:"bytes,60,opt,name=asset_id_str,json=assetIdStr,proto3" json:"asset_id_str,omitempty"`
	AssetIdentification              int32    `protobuf:"zigzag32,61,opt,name=asset_identification,json=assetIdentification,proto3" json:"asset_identification,omitempty"`
	IsFieldReplaceableUnit           bool     `protobuf:"varint,62,opt,name=is_field_replaceable_unit,json=isFieldReplaceableUnit,proto3" json:"is_field_replaceable_unit,omitempty"`
	ManufacturerAssetTags            int32    `protobuf:"zigzag32,63,opt,name=manufacturer_asset_tags,json=manufacturerAssetTags,proto3" json:"manufacturer_asset_tags,omitempty"`
	CompositeClassCode               int32    `protobuf:"zigzag32,64,opt,name=composite_class_code,json=compositeClassCode,proto3" json:"composite_class_code,omitempty"`
	MemorySize                       int32    `protobuf:"zigzag32,65,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	EnvironmentalMonitorPath         string   `protobuf:"bytes,66,opt,name=environmental_monitor_path,json=environmentalMonitorPath,proto3" json:"environmental_monitor_path,omitempty"`
	Alias                            string   `protobuf:"bytes,67,opt,name=alias,proto3" json:"alias,omitempty"`
	GroupFlag                        bool     `protobuf:"varint,68,opt,name=group_flag,json=groupFlag,proto3" json:"group_flag,omitempty"`
	NewDeviationNumber               int32    `protobuf:"zigzag32,69,opt,name=new_deviation_number,json=newDeviationNumber,proto3" json:"new_deviation_number,omitempty"`
	PhysicalLayerInterfaceModuleType int32    `protobuf:"zigzag32,70,opt,name=physical_layer_interface_module_type,json=physicalLayerInterfaceModuleType,proto3" json:"physical_layer_interface_module_type,omitempty"`
	UnrecognizedFru                  bool     `protobuf:"varint,71,opt,name=unrecognized_fru,json=unrecognizedFru,proto3" json:"unrecognized_fru,omitempty"`
	Redundancystate                  int32    `protobuf:"zigzag32,72,opt,name=redundancystate,proto3" json:"redundancystate,omitempty"`
	Ceport                           bool     `protobuf:"varint,73,opt,name=ceport,proto3" json:"ceport,omitempty"`
	XrScoped                         bool     `protobuf:"varint,74,opt,name=xr_scoped,json=xrScoped,proto3" json:"xr_scoped,omitempty"`
	UniqueId                         int32    `protobuf:"zigzag32,75,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *InvPhyEntityBasicInfo) Reset()         { *m = InvPhyEntityBasicInfo{} }
func (m *InvPhyEntityBasicInfo) String() string { return proto.CompactTextString(m) }
func (*InvPhyEntityBasicInfo) ProtoMessage()    {}
func (*InvPhyEntityBasicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_94f1af735311b03a, []int{1}
}

func (m *InvPhyEntityBasicInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvPhyEntityBasicInfo.Unmarshal(m, b)
}
func (m *InvPhyEntityBasicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvPhyEntityBasicInfo.Marshal(b, m, deterministic)
}
func (m *InvPhyEntityBasicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvPhyEntityBasicInfo.Merge(m, src)
}
func (m *InvPhyEntityBasicInfo) XXX_Size() int {
	return xxx_messageInfo_InvPhyEntityBasicInfo.Size(m)
}
func (m *InvPhyEntityBasicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InvPhyEntityBasicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InvPhyEntityBasicInfo proto.InternalMessageInfo

func (m *InvPhyEntityBasicInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetVendorType() string {
	if m != nil {
		return m.VendorType
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetHardwareRevision() string {
	if m != nil {
		return m.HardwareRevision
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetFirmwareRevision() string {
	if m != nil {
		return m.FirmwareRevision
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetSoftwareRevision() string {
	if m != nil {
		return m.SoftwareRevision
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetChipHardwareRevision() string {
	if m != nil {
		return m.ChipHardwareRevision
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetManufacturerName() string {
	if m != nil {
		return m.ManufacturerName
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetAssetIdStr() string {
	if m != nil {
		return m.AssetIdStr
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetAssetIdentification() int32 {
	if m != nil {
		return m.AssetIdentification
	}
	return 0
}

func (m *InvPhyEntityBasicInfo) GetIsFieldReplaceableUnit() bool {
	if m != nil {
		return m.IsFieldReplaceableUnit
	}
	return false
}

func (m *InvPhyEntityBasicInfo) GetManufacturerAssetTags() int32 {
	if m != nil {
		return m.ManufacturerAssetTags
	}
	return 0
}

func (m *InvPhyEntityBasicInfo) GetCompositeClassCode() int32 {
	if m != nil {
		return m.CompositeClassCode
	}
	return 0
}

func (m *InvPhyEntityBasicInfo) GetMemorySize() int32 {
	if m != nil {
		return m.MemorySize
	}
	return 0
}

func (m *InvPhyEntityBasicInfo) GetEnvironmentalMonitorPath() string {
	if m != nil {
		return m.EnvironmentalMonitorPath
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *InvPhyEntityBasicInfo) GetGroupFlag() bool {
	if m != nil {
		return m.GroupFlag
	}
	return false
}

func (m *InvPhyEntityBasicInfo) GetNewDeviationNumber() int32 {
	if m != nil {
		return m.NewDeviationNumber
	}
	return 0
}

func (m *InvPhyEntityBasicInfo) GetPhysicalLayerInterfaceModuleType() int32 {
	if m != nil {
		return m.PhysicalLayerInterfaceModuleType
	}
	return 0
}

func (m *InvPhyEntityBasicInfo) GetUnrecognizedFru() bool {
	if m != nil {
		return m.UnrecognizedFru
	}
	return false
}

func (m *InvPhyEntityBasicInfo) GetRedundancystate() int32 {
	if m != nil {
		return m.Redundancystate
	}
	return 0
}

func (m *InvPhyEntityBasicInfo) GetCeport() bool {
	if m != nil {
		return m.Ceport
	}
	return false
}

func (m *InvPhyEntityBasicInfo) GetXrScoped() bool {
	if m != nil {
		return m.XrScoped
	}
	return false
}

func (m *InvPhyEntityBasicInfo) GetUniqueId() int32 {
	if m != nil {
		return m.UniqueId
	}
	return 0
}

func init() {
	proto.RegisterType((*InvPhyEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.inv_basic_bag.inv_phy_entity_basic_info_KEYS")
	proto.RegisterType((*InvPhyEntityBasicInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.inv_basic_bag.inv_phy_entity_basic_info")
}

func init() { proto.RegisterFile("inv_phy_entity_basic_info.proto", fileDescriptor_94f1af735311b03a) }

var fileDescriptor_94f1af735311b03a = []byte{
	// 837 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0xdd, 0x72, 0x1b, 0x35,
	0x14, 0xc7, 0xc7, 0xa5, 0x49, 0x6d, 0xa5, 0x4c, 0xdb, 0x25, 0x4d, 0x55, 0x18, 0xa8, 0xa7, 0x70,
	0x11, 0x86, 0x19, 0x4f, 0xfc, 0x6d, 0x43, 0xf9, 0x28, 0x69, 0x43, 0x43, 0x69, 0x86, 0x71, 0xca,
	0x05, 0x57, 0x1a, 0x59, 0x3a, 0x6b, 0x6b, 0xd8, 0x95, 0x16, 0x49, 0xeb, 0x64, 0xf3, 0x2c, 0xbc,
	0x06, 0xcf, 0xc2, 0xeb, 0x30, 0x3a, 0x6b, 0xef, 0xda, 0xed, 0xf4, 0xe6, 0xe4, 0xe8, 0xf7, 0xff,
	0xaf, 0x74, 0x74, 0x24, 0xc5, 0xe4, 0x89, 0xd2, 0x2b, 0x96, 0x2d, 0x0b, 0x06, 0xda, 0x2b, 0x5f,
	0xb0, 0x39, 0x77, 0x4a, 0x30, 0xa5, 0x63, 0xd3, 0xc9, 0xac, 0xf1, 0x26, 0xfa, 0xb7, 0x21, 0x94,
	0x13, 0x86, 0x29, 0xe3, 0xd8, 0xb5, 0x65, 0x4a, 0xaf, 0xd2, 0x85, 0x65, 0x26, 0x03, 0xdb, 0x51,
	0x7a, 0x05, 0xda, 0x1b, 0x5b, 0x74, 0x2c, 0x17, 0x7f, 0x39, 0x8c, 0x9d, 0x72, 0x9e, 0x8e, 0x4b,
	0x8c, 0xef, 0x78, 0xa7, 0xba, 0x0e, 0x63, 0x08, 0x3d, 0x4c, 0x7b, 0x21, 0xf4, 0x31, 0xed, 0x87,
	0x30, 0xc0, 0x74, 0x10, 0xc2, 0x10, 0xd3, 0x61, 0x08, 0x23, 0x4c, 0x47, 0x21, 0x8c, 0x31, 0x1d,
	0x87, 0x30, 0xc1, 0x74, 0x12, 0xc2, 0x14, 0xd3, 0x69, 0x87, 0x7b, 0x6f, 0xd5, 0x3c, 0xf7, 0xe0,
	0x42, 0x39, 0xeb, 0xf2, 0xe7, 0x7c, 0xf1, 0xf4, 0x9f, 0x5b, 0xe4, 0x8b, 0x0f, 0xee, 0x8d, 0xbd,
	0x7e, 0xf9, 0xe7, 0x65, 0x14, 0x91, 0xdb, 0x9a, 0xa7, 0x40, 0x1b, 0xed, 0xc6, 0x71, 0x6b, 0x86,
	0x79, 0xf4, 0x90, 0xec, 0x87, 0xbf, 0xac, 0x4b, 0x6f, 0x21, 0xdd, 0x0b, 0xa3, 0x6e, 0x85, 0x7b,
	0xf4, 0xa3, 0x1a, 0xf7, 0x2a, 0xdc, 0xa7, 0xb7, 0x6b, 0xdc, 0xaf, 0xf0, 0x80, 0xee, 0xd5, 0x78,
	0x50, 0xe1, 0x21, 0xdd, 0xaf, 0xf1, 0xb0, 0xc2, 0x23, 0x7a, 0xa7, 0xc6, 0xa3, 0x0a, 0x8f, 0x69,
	0xb3, 0xc6, 0xe3, 0x0a, 0x4f, 0x68, 0xab, 0xc6, 0x93, 0x0a, 0x4f, 0x29, 0xa9, 0xf1, 0x34, 0x7a,
	0x44, 0xee, 0x94, 0xdb, 0x39, 0xa1, 0x07, 0xc8, 0xd1, 0xd5, 0x3d, 0x79, 0xfa, 0x5f, 0x93, 0x3c,
	0xfe, 0x60, 0x7b, 0xa2, 0x36, 0x39, 0x90, 0xe0, 0x84, 0x55, 0x99, 0x57, 0x46, 0xd3, 0x1e, 0x7e,
	0xba, 0x8d, 0xa2, 0x27, 0xe4, 0x60, 0x05, 0x5a, 0x1a, 0xcb, 0x7c, 0x91, 0x01, 0xed, 0xa3, 0x83,
	0x94, 0xe8, 0x6d, 0x91, 0x41, 0xd5, 0xdc, 0xc1, 0x56, 0x73, 0xbf, 0x21, 0x0f, 0x96, 0xdc, 0xca,
	0x2b, 0x6e, 0x81, 0x59, 0x58, 0x29, 0x17, 0x26, 0x1f, 0xa2, 0xe1, 0xfe, 0x46, 0x98, 0xad, 0x79,
	0x30, 0xc7, 0xca, 0xa6, 0xbb, 0xe6, 0x51, 0x69, 0xde, 0x08, 0xdb, 0x66, 0x67, 0x62, 0xbf, 0x6b,
	0x1e, 0x97, 0xe6, 0x8d, 0x50, 0x99, 0x07, 0xe4, 0x48, 0x2c, 0x55, 0xc6, 0xde, 0xaf, 0x65, 0x82,
	0x5f, 0x1c, 0x06, 0xf5, 0xd5, 0xbb, 0xf5, 0x7c, 0x49, 0x3e, 0x76, 0x60, 0x15, 0x4f, 0x98, 0xce,
	0xd3, 0x39, 0x58, 0x3a, 0x45, 0xf3, 0xdd, 0x12, 0x5e, 0x20, 0x0b, 0x75, 0xa4, 0x5c, 0xe7, 0x31,
	0x17, 0x3e, 0xb7, 0x60, 0x19, 0xb6, 0xe0, 0xdb, 0xb2, 0x8e, 0x6d, 0xe1, 0x22, 0xb4, 0xe3, 0x73,
	0x42, 0x52, 0x23, 0x21, 0x29, 0x5d, 0xdf, 0xa1, 0xab, 0x85, 0x04, 0xe5, 0x36, 0xb9, 0xcb, 0x9d,
	0x03, 0xcf, 0x94, 0x64, 0xce, 0x5b, 0xfa, 0xac, 0xec, 0x31, 0xb2, 0x73, 0x79, 0xe9, 0x6d, 0xd4,
	0x25, 0x87, 0x1b, 0x47, 0x38, 0xc3, 0x58, 0x09, 0x8e, 0xe7, 0xf5, 0x7d, 0xbb, 0x71, 0xfc, 0x60,
	0xf6, 0xc9, 0xda, 0xb9, 0x2d, 0x45, 0x53, 0xf2, 0x58, 0x39, 0x16, 0x2b, 0x48, 0x24, 0xb3, 0x90,
	0x25, 0x5c, 0x00, 0x9f, 0x27, 0xc0, 0x72, 0xad, 0x3c, 0xfd, 0xa1, 0xdd, 0x38, 0x6e, 0xce, 0x8e,
	0x94, 0x3b, 0x0b, 0xfa, 0xac, 0x96, 0xff, 0xd0, 0xca, 0x47, 0x23, 0xf2, 0x68, 0x67, 0x6f, 0xe5,
	0xd2, 0x9e, 0x2f, 0x1c, 0xfd, 0x11, 0x17, 0x7c, 0xb8, 0x2d, 0x3f, 0x0f, 0xea, 0x5b, 0xbe, 0x70,
	0xd1, 0x09, 0x39, 0x14, 0x26, 0xcd, 0x8c, 0x53, 0x1e, 0x98, 0x48, 0xb8, 0x73, 0x4c, 0x18, 0x09,
	0xf4, 0x27, 0xfc, 0x28, 0xaa, 0xb4, 0xd3, 0x20, 0x9d, 0x1a, 0x09, 0xe1, 0x72, 0xa5, 0x90, 0x1a,
	0x5b, 0x30, 0xa7, 0x6e, 0x80, 0x3e, 0x47, 0x23, 0x29, 0xd1, 0xa5, 0xba, 0x81, 0xe8, 0x19, 0xf9,
	0x14, 0xf4, 0x4a, 0x59, 0xa3, 0x53, 0xd0, 0x9e, 0x27, 0x2c, 0x35, 0x5a, 0x79, 0x63, 0x59, 0xc6,
	0xfd, 0x92, 0xfe, 0x8c, 0x8d, 0xa2, 0x3b, 0x8e, 0x37, 0xa5, 0xe1, 0x77, 0xee, 0x97, 0xd1, 0x21,
	0xd9, 0xe3, 0x89, 0xe2, 0x8e, 0x9e, 0x96, 0x4f, 0x05, 0x07, 0xe1, 0x34, 0x16, 0xd6, 0xe4, 0x19,
	0x8b, 0x13, 0xbe, 0xa0, 0x2f, 0xb0, 0x15, 0x2d, 0x24, 0x67, 0x09, 0x5f, 0x84, 0x5d, 0x68, 0xb8,
	0x62, 0x12, 0x56, 0x0a, 0x3b, 0xb9, 0xb9, 0x05, 0x2f, 0xcb, 0x5d, 0x68, 0xb8, 0x7a, 0xb1, 0x91,
	0xd6, 0x77, 0xe1, 0x82, 0x7c, 0x95, 0x2d, 0x0b, 0xa7, 0x04, 0x4f, 0x58, 0xc2, 0x0b, 0x08, 0xff,
	0x3c, 0x3d, 0xd8, 0x98, 0x0b, 0x60, 0xa9, 0x91, 0x79, 0x02, 0xe5, 0xdb, 0x39, 0xc3, 0x19, 0xda,
	0x1b, 0xef, 0x6f, 0xc1, 0x7a, 0xbe, 0x71, 0xbe, 0x41, 0x23, 0xbe, 0xa8, 0xaf, 0xc9, 0xfd, 0x5c,
	0x5b, 0x10, 0x66, 0xa1, 0xd5, 0x0d, 0x48, 0x16, 0xdb, 0x9c, 0xfe, 0x82, 0x65, 0xde, 0xdb, 0xe6,
	0x67, 0x36, 0x8f, 0x8e, 0xc9, 0x3d, 0x0b, 0x32, 0xd7, 0x92, 0x6b, 0x51, 0x38, 0xcf, 0x3d, 0xd0,
	0x57, 0xb8, 0xca, 0xbb, 0x38, 0x3a, 0x22, 0xfb, 0x02, 0x32, 0x63, 0x3d, 0x3d, 0xc7, 0xa9, 0xd6,
	0xa3, 0xe8, 0x33, 0xd2, 0xba, 0xb6, 0xcc, 0x09, 0x93, 0x81, 0xa4, 0xbf, 0xa2, 0xd4, 0xbc, 0xb6,
	0x97, 0x38, 0x0e, 0x62, 0xae, 0xd5, 0xdf, 0x39, 0x30, 0x25, 0xe9, 0x6b, 0x9c, 0xb8, 0x59, 0x82,
	0x73, 0x39, 0xdf, 0xc7, 0xdf, 0x8d, 0xfe, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x3d, 0x02,
	0x6a, 0x5a, 0x06, 0x00, 0x00,
}
