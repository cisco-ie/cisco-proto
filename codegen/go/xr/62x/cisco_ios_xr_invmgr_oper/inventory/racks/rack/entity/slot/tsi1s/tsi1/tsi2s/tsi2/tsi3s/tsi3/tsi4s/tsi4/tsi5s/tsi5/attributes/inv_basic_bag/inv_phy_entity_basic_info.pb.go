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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_attributes_inv_basic_bag

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
	proto.RegisterType((*InvPhyEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.inv_basic_bag.inv_phy_entity_basic_info_KEYS")
	proto.RegisterType((*InvPhyEntityBasicInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.inv_basic_bag.inv_phy_entity_basic_info")
}

func init() { proto.RegisterFile("inv_phy_entity_basic_info.proto", fileDescriptor_94f1af735311b03a) }

var fileDescriptor_94f1af735311b03a = []byte{
	// 786 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x6d, 0x6f, 0x1c, 0x35,
	0x10, 0xc7, 0x75, 0xd0, 0x84, 0xc4, 0x2d, 0x6a, 0xbb, 0xa4, 0xc1, 0x05, 0x41, 0x4f, 0x85, 0x17,
	0x41, 0x48, 0x27, 0x72, 0x0f, 0x81, 0x42, 0x79, 0x28, 0x69, 0x43, 0x43, 0x69, 0x84, 0x2e, 0xe5,
	0x05, 0xaf, 0x2c, 0x9f, 0x3d, 0x7b, 0x37, 0x62, 0xd7, 0x5e, 0xc6, 0xde, 0x4b, 0x36, 0xdf, 0x80,
	0x2f, 0xc5, 0xd7, 0xe0, 0xeb, 0x20, 0xcf, 0xde, 0x53, 0x5a, 0xf5, 0xcd, 0x9c, 0xfd, 0xfb, 0xff,
	0xd7, 0x9e, 0x19, 0xdb, 0x27, 0x1e, 0xa0, 0x9b, 0xab, 0x6a, 0xd6, 0x28, 0x70, 0x11, 0x63, 0xa3,
	0x26, 0x3a, 0xa0, 0x51, 0xe8, 0x72, 0xdf, 0xab, 0xc8, 0x47, 0x9f, 0xfd, 0xd3, 0x31, 0x18, 0x8c,
	0x57, 0xe8, 0x83, 0xba, 0x24, 0x85, 0x6e, 0x5e, 0x4e, 0x49, 0xf9, 0x0a, 0xa8, 0x87, 0x6e, 0x0e,
	0x2e, 0x7a, 0x6a, 0x7a, 0xa4, 0xcd, 0x5f, 0x81, 0x63, 0xaf, 0x5d, 0xa7, 0x17, 0x0a, 0x1f, 0x7b,
	0x31, 0xe0, 0x61, 0xe0, 0x98, 0x42, 0x9f, 0x87, 0xfd, 0x14, 0x06, 0x3c, 0x1c, 0xa4, 0x30, 0xe4,
	0xe1, 0x30, 0x85, 0x11, 0x0f, 0x47, 0x3d, 0x1d, 0x23, 0xe1, 0xa4, 0x8e, 0x10, 0xd2, 0x16, 0x8b,
	0x94, 0x26, 0x7a, 0xfa, 0xf0, 0xdf, 0x8e, 0xf8, 0xf4, 0xad, 0xf9, 0xaa, 0x17, 0xcf, 0xfe, 0x3c,
	0xcf, 0x32, 0x71, 0xc3, 0xe9, 0x12, 0x64, 0xa7, 0xdb, 0x39, 0xd8, 0x1d, 0xf3, 0x38, 0xbb, 0x27,
	0xb6, 0xd3, 0xaf, 0x3a, 0x94, 0xef, 0x30, 0xdd, 0x4a, 0xb3, 0xc3, 0x15, 0xee, 0xcb, 0x77, 0xd7,
	0xb8, 0xbf, 0xc2, 0x03, 0x79, 0x63, 0x8d, 0x07, 0x2b, 0x3c, 0x94, 0x5b, 0x6b, 0x3c, 0x5c, 0xe1,
	0x91, 0xdc, 0x5e, 0xe3, 0xd1, 0x0a, 0x1f, 0xc9, 0xf7, 0xd6, 0xf8, 0xe8, 0xe1, 0x7f, 0x3b, 0xe2,
	0xfe, 0x5b, 0x0b, 0xc8, 0xba, 0xe2, 0xa6, 0x85, 0x60, 0x08, 0xab, 0x88, 0xde, 0xc9, 0x3e, 0x7f,
	0xb9, 0x89, 0xb2, 0x07, 0xe2, 0xe6, 0x1c, 0x9c, 0xf5, 0xa4, 0x62, 0x53, 0x81, 0x1c, 0xb0, 0x43,
	0xb4, 0xe8, 0x55, 0x53, 0xc1, 0xaa, 0xfc, 0xe1, 0x46, 0xf9, 0x5f, 0x8a, 0xbb, 0x33, 0x4d, 0xf6,
	0x42, 0x13, 0x28, 0x82, 0x39, 0x86, 0xb4, 0xf8, 0x88, 0x0d, 0x77, 0x96, 0xc2, 0x78, 0xc1, 0x93,
	0x39, 0x47, 0x2a, 0xaf, 0x9b, 0x8f, 0x5a, 0xf3, 0x52, 0xd8, 0x34, 0x07, 0x9f, 0xc7, 0xeb, 0xe6,
	0xaf, 0x5b, 0xf3, 0x52, 0x58, 0x99, 0x87, 0x62, 0xdf, 0xcc, 0xb0, 0x52, 0x6f, 0xe6, 0xf2, 0x0d,
	0x7f, 0xb1, 0x97, 0xd4, 0xe7, 0xaf, 0xe7, 0xf3, 0x99, 0x78, 0x3f, 0x00, 0xa1, 0x2e, 0x94, 0xab,
	0xcb, 0x09, 0x90, 0x7c, 0xc4, 0xe6, 0x5b, 0x2d, 0x3c, 0x63, 0x96, 0xf2, 0x28, 0xb5, 0xab, 0x73,
	0x6d, 0x62, 0x4d, 0x40, 0x8a, 0x5b, 0xf0, 0x6d, 0x9b, 0xc7, 0xa6, 0x70, 0x96, 0xda, 0xf1, 0x89,
	0x10, 0xa5, 0xb7, 0x50, 0xb4, 0xae, 0xef, 0xd8, 0xb5, 0xcb, 0x84, 0xe5, 0xae, 0xb8, 0xa5, 0x43,
	0x80, 0xa8, 0xd0, 0xaa, 0x10, 0x49, 0x3e, 0x6e, 0x7b, 0xcc, 0xec, 0xd4, 0x9e, 0x47, 0xca, 0x0e,
	0xc5, 0xde, 0xd2, 0x91, 0xce, 0x30, 0x47, 0xa3, 0xf9, 0xbc, 0xbe, 0xef, 0x76, 0x0e, 0xee, 0x8e,
	0x3f, 0x58, 0x38, 0x37, 0xa5, 0xec, 0x91, 0xb8, 0x8f, 0x41, 0xe5, 0x08, 0x85, 0x55, 0x04, 0x55,
	0xa1, 0x0d, 0xe8, 0x49, 0x01, 0xaa, 0x76, 0x18, 0xe5, 0x0f, 0xdd, 0xce, 0xc1, 0xce, 0x78, 0x1f,
	0xc3, 0x49, 0xd2, 0xc7, 0x6b, 0xf9, 0x0f, 0x87, 0x31, 0x3b, 0x12, 0x1f, 0x5e, 0xab, 0xad, 0xdd,
	0x3a, 0xea, 0x69, 0x90, 0x3f, 0xf2, 0x86, 0xf7, 0x36, 0xe5, 0x27, 0x49, 0x7d, 0xa5, 0xa7, 0x21,
	0xfb, 0x4a, 0xec, 0x19, 0x5f, 0x56, 0x3e, 0x60, 0x04, 0x65, 0x0a, 0x1d, 0x82, 0x32, 0xde, 0x82,
	0xfc, 0x89, 0x3f, 0xca, 0x56, 0xda, 0x71, 0x92, 0x8e, 0xbd, 0x85, 0x74, 0xb9, 0x4a, 0x28, 0x3d,
	0x35, 0x2a, 0xe0, 0x15, 0xc8, 0x27, 0x6c, 0x14, 0x2d, 0x3a, 0xc7, 0x2b, 0xc8, 0x1e, 0x8b, 0x8f,
	0xc0, 0xcd, 0x91, 0xbc, 0x2b, 0xc1, 0x45, 0x5d, 0xa8, 0xd2, 0x3b, 0x8c, 0x9e, 0x54, 0xa5, 0xe3,
	0x4c, 0xfe, 0xcc, 0x8d, 0x92, 0xd7, 0x1c, 0x2f, 0x5b, 0xc3, 0xef, 0x3a, 0xce, 0xb2, 0x3d, 0xb1,
	0xa5, 0x0b, 0xd4, 0x41, 0x1e, 0xb7, 0x2f, 0x82, 0x27, 0xe9, 0x34, 0xa6, 0xe4, 0xeb, 0x4a, 0xe5,
	0x85, 0x9e, 0xca, 0xa7, 0xdc, 0x8a, 0x5d, 0x26, 0x27, 0x85, 0x9e, 0xa6, 0x2a, 0x1c, 0x5c, 0x28,
	0x0b, 0x73, 0xe4, 0x4e, 0x2e, 0x6f, 0xc1, 0xb3, 0xb6, 0x0a, 0x07, 0x17, 0x4f, 0x97, 0xd2, 0xe2,
	0x2e, 0x9c, 0x89, 0xcf, 0xab, 0x59, 0x13, 0xd0, 0xe8, 0x42, 0x15, 0xba, 0x81, 0xf4, 0x97, 0x15,
	0x81, 0x72, 0x6d, 0x40, 0x95, 0xde, 0xd6, 0x05, 0xb4, 0x6f, 0xe7, 0x84, 0x57, 0xe8, 0x2e, 0xbd,
	0xbf, 0x25, 0xeb, 0xe9, 0xd2, 0xf9, 0x92, 0x8d, 0xfc, 0xa2, 0xbe, 0x10, 0x77, 0x6a, 0x47, 0x60,
	0xfc, 0xd4, 0xe1, 0x15, 0x58, 0x95, 0x53, 0x2d, 0x7f, 0xe1, 0x34, 0x6f, 0x6f, 0xf2, 0x13, 0xaa,
	0xb3, 0x03, 0x71, 0x9b, 0xc0, 0xd6, 0xce, 0x6a, 0x67, 0x9a, 0x10, 0x75, 0x04, 0xf9, 0x9c, 0x77,
	0x79, 0x1d, 0x67, 0xfb, 0x62, 0xdb, 0x40, 0xe5, 0x29, 0xca, 0x53, 0x5e, 0x6a, 0x31, 0xcb, 0x3e,
	0x16, 0xbb, 0x97, 0xa4, 0x82, 0xf1, 0x15, 0x58, 0xf9, 0x2b, 0x4b, 0x3b, 0x97, 0x74, 0xce, 0xf3,
	0x24, 0xd6, 0x0e, 0xff, 0xae, 0x41, 0xa1, 0x95, 0x2f, 0x78, 0xe1, 0x9d, 0x16, 0x9c, 0xda, 0xc9,
	0x36, 0xff, 0x5b, 0x0f, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x22, 0x54, 0xb5, 0x9e, 0xd0, 0x05,
	0x00, 0x00,
}