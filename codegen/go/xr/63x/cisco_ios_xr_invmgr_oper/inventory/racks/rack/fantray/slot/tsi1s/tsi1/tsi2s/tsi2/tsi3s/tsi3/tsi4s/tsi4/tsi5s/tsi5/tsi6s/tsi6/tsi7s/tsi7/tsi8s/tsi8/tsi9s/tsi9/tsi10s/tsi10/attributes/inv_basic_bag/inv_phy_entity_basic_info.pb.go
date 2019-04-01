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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_tsi10s_tsi10_attributes_inv_basic_bag

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
	Name_11              string   `protobuf:"bytes,12,opt,name=name_11,json=name11,proto3" json:"name_11,omitempty"`
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

func (m *InvPhyEntityBasicInfo_KEYS) GetName_11() string {
	if m != nil {
		return m.Name_11
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
	proto.RegisterType((*InvPhyEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.attributes.inv_basic_bag.inv_phy_entity_basic_info_KEYS")
	proto.RegisterType((*InvPhyEntityBasicInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.attributes.inv_basic_bag.inv_phy_entity_basic_info")
}

func init() { proto.RegisterFile("inv_phy_entity_basic_info.proto", fileDescriptor_94f1af735311b03a) }

var fileDescriptor_94f1af735311b03a = []byte{
	// 853 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0xdb, 0x6e, 0x1b, 0x37,
	0x13, 0xc7, 0xa1, 0x7c, 0xb1, 0x23, 0xd1, 0xfe, 0x90, 0x64, 0xeb, 0x38, 0x4c, 0x8b, 0x36, 0x42,
	0xda, 0x0b, 0x17, 0x05, 0x04, 0x9d, 0x0f, 0x6d, 0x7a, 0x48, 0x9d, 0xb8, 0x71, 0xd3, 0x18, 0x85,
	0x9c, 0x5e, 0xf4, 0x8a, 0xa0, 0xc8, 0x59, 0x89, 0xe8, 0x2e, 0xb9, 0x25, 0xb9, 0xb2, 0xd7, 0xef,
	0xd5, 0x27, 0xe8, 0x43, 0xf4, 0x75, 0x0a, 0xce, 0x4a, 0xbb, 0x72, 0x8a, 0xdc, 0x8c, 0x87, 0xbf,
	0xff, 0x7f, 0xc9, 0x99, 0xd9, 0xa5, 0x45, 0x9e, 0x2a, 0xbd, 0x66, 0xd9, 0xaa, 0x60, 0xa0, 0xbd,
	0xf2, 0x05, 0x5b, 0x70, 0xa7, 0x04, 0x53, 0x3a, 0x36, 0x9d, 0xcc, 0x1a, 0x6f, 0xa2, 0xbf, 0x1b,
	0x42, 0x39, 0x61, 0x98, 0x32, 0x8e, 0x5d, 0x5b, 0xa6, 0xf4, 0x3a, 0x5d, 0x5a, 0x66, 0x32, 0xb0,
	0x1d, 0xa5, 0xd7, 0xa0, 0xbd, 0xb1, 0x45, 0xc7, 0x72, 0xf1, 0x87, 0xc3, 0xd8, 0x89, 0xb9, 0xf6,
	0x96, 0x17, 0x1d, 0x97, 0x18, 0xdf, 0xf1, 0x4e, 0xf5, 0x1c, 0xc6, 0x10, 0xfa, 0x98, 0xf6, 0x43,
	0x18, 0x60, 0x3a, 0x08, 0x61, 0x88, 0xe9, 0x30, 0x84, 0x11, 0xa6, 0xa3, 0x10, 0xc6, 0x98, 0x8e,
	0x43, 0x98, 0x60, 0x3a, 0x09, 0x61, 0x8a, 0xe9, 0x34, 0x84, 0x19, 0xa6, 0x33, 0xdc, 0xbc, 0x5b,
	0x9e, 0xd1, 0xed, 0x70, 0xef, 0xad, 0x5a, 0xe4, 0x1e, 0x5c, 0x28, 0x6e, 0xd3, 0xcc, 0x82, 0x2f,
	0x9f, 0xfd, 0x75, 0x87, 0x7c, 0xf6, 0xc1, 0x4e, 0xd9, 0x9b, 0x57, 0xbf, 0x5f, 0x46, 0x11, 0xb9,
	0xab, 0x79, 0x0a, 0xb4, 0xd1, 0x6e, 0x9c, 0xb4, 0xe6, 0x98, 0x47, 0x8f, 0xc8, 0x7e, 0xf8, 0xcb,
	0x7a, 0xf4, 0x0e, 0xd2, 0xbd, 0xb0, 0xea, 0x55, 0xb8, 0x4f, 0xff, 0x57, 0xe3, 0x7e, 0x85, 0x07,
	0xf4, 0x6e, 0x8d, 0x07, 0x15, 0x1e, 0xd2, 0xbd, 0x1a, 0x0f, 0x2b, 0x3c, 0xa2, 0xfb, 0x35, 0x1e,
	0x55, 0x78, 0x4c, 0xef, 0xd5, 0x78, 0x5c, 0xe1, 0x09, 0x6d, 0xd6, 0x78, 0x52, 0xe1, 0x29, 0x6d,
	0xd5, 0x78, 0x5a, 0xe1, 0x19, 0x25, 0x35, 0x9e, 0x45, 0x8f, 0xc9, 0xbd, 0xb2, 0x9d, 0x2e, 0x3d,
	0x40, 0x8e, 0xae, 0x5e, 0xb7, 0x16, 0x7a, 0xf4, 0x70, 0x47, 0xe8, 0x3d, 0xfb, 0xa7, 0x49, 0x9e,
	0x7c, 0x70, 0x6e, 0x51, 0x9b, 0x1c, 0x48, 0x70, 0xc2, 0xaa, 0xcc, 0x2b, 0xa3, 0x69, 0x1f, 0x1f,
	0xdd, 0x45, 0xd1, 0x53, 0x72, 0xb0, 0x06, 0x2d, 0x8d, 0x65, 0xbe, 0xc8, 0x80, 0x0e, 0xd0, 0x41,
	0x4a, 0xf4, 0xae, 0xc8, 0xa0, 0x9a, 0xfa, 0x70, 0x67, 0xea, 0x5f, 0x91, 0x87, 0x2b, 0x6e, 0xe5,
	0x15, 0xb7, 0xc0, 0x2c, 0xac, 0x95, 0x0b, 0x9b, 0x8f, 0xd0, 0xf0, 0x60, 0x2b, 0xcc, 0x37, 0x3c,
	0x98, 0x63, 0x65, 0xd3, 0xdb, 0xe6, 0x71, 0x69, 0xde, 0x0a, 0xbb, 0x66, 0x67, 0x62, 0x7f, 0xdb,
	0x3c, 0x29, 0xcd, 0x5b, 0xa1, 0x32, 0x0f, 0xc9, 0xb1, 0x58, 0xa9, 0x8c, 0xfd, 0xb7, 0x96, 0x29,
	0x3e, 0x71, 0x14, 0xd4, 0xd7, 0xef, 0xd7, 0xf3, 0x39, 0xf9, 0xbf, 0x03, 0xab, 0x78, 0xc2, 0x74,
	0x9e, 0x2e, 0xc0, 0xd2, 0x19, 0x9a, 0x0f, 0x4b, 0x78, 0x81, 0x2c, 0xd4, 0x91, 0x72, 0x9d, 0xc7,
	0x5c, 0xf8, 0xdc, 0x82, 0x65, 0x38, 0x82, 0xaf, 0xcb, 0x3a, 0x76, 0x85, 0x8b, 0x30, 0x8e, 0x4f,
	0x09, 0x49, 0x8d, 0x84, 0xa4, 0x74, 0x7d, 0x83, 0xae, 0x16, 0x12, 0x94, 0xdb, 0xe4, 0x90, 0x3b,
	0x07, 0x9e, 0x29, 0xc9, 0x9c, 0xb7, 0xf4, 0x79, 0x39, 0x63, 0x64, 0xe7, 0xf2, 0xd2, 0xdb, 0xa8,
	0x47, 0x8e, 0xb6, 0x8e, 0xf0, 0x0e, 0x63, 0x25, 0x38, 0xbe, 0xaf, 0x6f, 0xdb, 0x8d, 0x93, 0x87,
	0xf3, 0x8f, 0x36, 0xce, 0x5d, 0x29, 0x9a, 0x91, 0x27, 0xca, 0xb1, 0x58, 0x41, 0x22, 0x99, 0x85,
	0x2c, 0xe1, 0x02, 0xf8, 0x22, 0x01, 0x96, 0x6b, 0xe5, 0xe9, 0x77, 0xed, 0xc6, 0x49, 0x73, 0x7e,
	0xac, 0xdc, 0x59, 0xd0, 0xe7, 0xb5, 0xfc, 0x9b, 0x56, 0x3e, 0x1a, 0x93, 0xc7, 0xb7, 0x7a, 0x2b,
	0x8f, 0xf6, 0x7c, 0xe9, 0xe8, 0xf7, 0x78, 0xe0, 0xa3, 0x5d, 0xf9, 0x45, 0x50, 0xdf, 0xf1, 0xa5,
	0x8b, 0xba, 0xe4, 0x48, 0x98, 0x34, 0x33, 0x4e, 0x79, 0x60, 0x22, 0xe1, 0xce, 0x31, 0x61, 0x24,
	0xd0, 0x1f, 0xf0, 0xa1, 0xa8, 0xd2, 0x4e, 0x83, 0x74, 0x6a, 0x24, 0x84, 0x8f, 0x2b, 0x85, 0xd4,
	0xd8, 0x82, 0x39, 0x75, 0x03, 0xf4, 0x05, 0x1a, 0x49, 0x89, 0x2e, 0xd5, 0x0d, 0x44, 0xcf, 0xc9,
	0xc7, 0xa0, 0xd7, 0xca, 0x1a, 0x9d, 0x82, 0xf6, 0x3c, 0x61, 0xa9, 0xd1, 0xca, 0x1b, 0xcb, 0x32,
	0xee, 0x57, 0xf4, 0x47, 0x1c, 0x14, 0xbd, 0xe5, 0x78, 0x5b, 0x1a, 0x7e, 0xe5, 0x7e, 0x15, 0x1d,
	0x91, 0x3d, 0x9e, 0x28, 0xee, 0xe8, 0x69, 0x79, 0x87, 0x70, 0x11, 0xde, 0xc6, 0xd2, 0x9a, 0x3c,
	0x63, 0x71, 0xc2, 0x97, 0xf4, 0x25, 0x8e, 0xa2, 0x85, 0xe4, 0x2c, 0xe1, 0xcb, 0xd0, 0x85, 0x86,
	0x2b, 0x26, 0x61, 0xad, 0x70, 0x92, 0xdb, 0xaf, 0xe0, 0x55, 0xd9, 0x85, 0x86, 0xab, 0x97, 0x5b,
	0x69, 0xf3, 0x2d, 0x5c, 0x90, 0x2f, 0xb2, 0x55, 0xe1, 0x94, 0xe0, 0x09, 0x4b, 0x78, 0x01, 0xe1,
	0x7f, 0xac, 0x07, 0x1b, 0x73, 0x01, 0x2c, 0x35, 0x32, 0x4f, 0xa0, 0xbc, 0x3b, 0x67, 0xb8, 0x43,
	0x7b, 0xeb, 0xfd, 0x25, 0x58, 0xcf, 0xb7, 0xce, 0xb7, 0x68, 0xc4, 0x1b, 0xf5, 0x25, 0x79, 0x90,
	0x6b, 0x0b, 0xc2, 0x2c, 0xb5, 0xba, 0x01, 0xc9, 0x62, 0x9b, 0xd3, 0x9f, 0xb0, 0xcc, 0xfb, 0xbb,
	0xfc, 0xcc, 0xe6, 0xd1, 0x09, 0xb9, 0x6f, 0x41, 0xe6, 0x5a, 0x72, 0x2d, 0x0a, 0xe7, 0xb9, 0x07,
	0xfa, 0x1a, 0x4f, 0x79, 0x1f, 0x47, 0xc7, 0x64, 0x5f, 0x40, 0x66, 0xac, 0xa7, 0xe7, 0xb8, 0xd5,
	0x66, 0x15, 0x7d, 0x42, 0x5a, 0xd7, 0x96, 0x39, 0x61, 0x32, 0x90, 0xf4, 0x67, 0x94, 0x9a, 0xd7,
	0xf6, 0x12, 0xd7, 0x41, 0xcc, 0xb5, 0xfa, 0x33, 0x07, 0xa6, 0x24, 0x7d, 0x83, 0x1b, 0x37, 0x4b,
	0x70, 0x2e, 0x17, 0xfb, 0xf8, 0xf3, 0x32, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x40, 0x74, 0x8a,
	0x71, 0x81, 0x06, 0x00, 0x00,
}
