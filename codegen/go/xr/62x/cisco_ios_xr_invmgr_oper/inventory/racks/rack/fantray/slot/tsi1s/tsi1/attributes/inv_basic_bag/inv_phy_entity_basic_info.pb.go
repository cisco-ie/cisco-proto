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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_attributes_inv_basic_bag

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
	proto.RegisterType((*InvPhyEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.inv_basic_bag.inv_phy_entity_basic_info_KEYS")
	proto.RegisterType((*InvPhyEntityBasicInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.inv_basic_bag.inv_phy_entity_basic_info")
}

func init() { proto.RegisterFile("inv_phy_entity_basic_info.proto", fileDescriptor_94f1af735311b03a) }

var fileDescriptor_94f1af735311b03a = []byte{
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x6f, 0x1c, 0x35,
	0x10, 0xc7, 0x15, 0xa0, 0x51, 0xe2, 0x16, 0xb5, 0x5d, 0xd2, 0xe0, 0x82, 0xa0, 0xa7, 0xc2, 0x43,
	0x10, 0xd2, 0x89, 0xa4, 0xa5, 0x50, 0x28, 0x3f, 0x4a, 0xda, 0xd0, 0x50, 0x1a, 0xa1, 0x4b, 0x79,
	0xe0, 0x85, 0x91, 0xcf, 0x9e, 0xbd, 0x1b, 0xb1, 0x6b, 0x2f, 0x63, 0xef, 0x25, 0x9b, 0x7f, 0x94,
	0x7f, 0x07, 0x79, 0xf6, 0x36, 0xbd, 0x6b, 0xd5, 0x97, 0xbd, 0xf3, 0xf7, 0xfb, 0xf1, 0xcc, 0x78,
	0xfc, 0x43, 0xdd, 0x21, 0xbf, 0x80, 0x66, 0xde, 0x01, 0xfa, 0x44, 0xa9, 0x83, 0xa9, 0x89, 0x64,
	0x81, 0x7c, 0x19, 0xc6, 0x0d, 0x87, 0x14, 0x8a, 0xbf, 0x2d, 0x45, 0x1b, 0x80, 0x42, 0x84, 0x73,
	0x06, 0xf2, 0x8b, 0x7a, 0xc6, 0x10, 0x1a, 0xe4, 0x31, 0xf9, 0x05, 0xfa, 0x14, 0xb8, 0x1b, 0xb3,
	0xb1, 0xff, 0x44, 0xf9, 0x8e, 0x4b, 0xe3, 0x13, 0x9b, 0x6e, 0x1c, 0xab, 0x90, 0xc6, 0x29, 0xd2,
	0x7e, 0x94, 0xef, 0xd8, 0xa4, 0xc4, 0x34, 0x6d, 0x13, 0xc6, 0x3c, 0x6f, 0x99, 0x66, 0x6a, 0x66,
	0x77, 0xa7, 0xea, 0xd3, 0xb7, 0x96, 0x00, 0xcf, 0x9f, 0xfe, 0x75, 0x5a, 0x14, 0xea, 0x3d, 0x6f,
	0x6a, 0xd4, 0x1b, 0xa3, 0x8d, 0xbd, 0xed, 0x89, 0xfc, 0x2f, 0x6e, 0xa9, 0xcd, 0xfc, 0x0b, 0xfb,
	0xfa, 0x1d, 0x51, 0xaf, 0xe4, 0xd1, 0xfe, 0xa5, 0x7c, 0xa0, 0xdf, 0x7d, 0x25, 0x1f, 0xdc, 0xfd,
	0x6f, 0x4b, 0xdd, 0x7e, 0x6b, 0x92, 0x62, 0xa4, 0xae, 0x3a, 0x8c, 0x96, 0xa9, 0x49, 0x14, 0xbc,
	0x3e, 0x90, 0x99, 0xab, 0x52, 0x71, 0x47, 0x5d, 0x5d, 0xa0, 0x77, 0x81, 0x21, 0x75, 0x0d, 0xea,
	0x7b, 0x42, 0xa8, 0x5e, 0x7a, 0xd9, 0x35, 0x78, 0x59, 0xe2, 0xfd, 0x95, 0x12, 0xbf, 0x54, 0x37,
	0xe7, 0x86, 0xdd, 0x99, 0x61, 0x04, 0xc6, 0x05, 0xc5, 0x1c, 0xfc, 0x6b, 0x01, 0x6e, 0x0c, 0xc6,
	0x64, 0xa9, 0x67, 0xb8, 0x24, 0xae, 0xd7, 0xe1, 0x07, 0x3d, 0x3c, 0x18, 0xab, 0x70, 0x0c, 0x65,
	0x5a, 0x87, 0xbf, 0xe9, 0xe1, 0xc1, 0xb8, 0x84, 0xef, 0xab, 0x5d, 0x3b, 0xa7, 0x06, 0xde, 0xac,
	0xe5, 0x5b, 0x99, 0xb1, 0x93, 0xdd, 0x67, 0xaf, 0xd7, 0xf3, 0x99, 0x7a, 0x3f, 0x22, 0x93, 0xa9,
	0xc0, 0xb7, 0xf5, 0x14, 0x59, 0x3f, 0x14, 0xf8, 0x5a, 0x2f, 0x9e, 0x88, 0x96, 0xeb, 0xa8, 0x8d,
	0x6f, 0x4b, 0x63, 0x53, 0xcb, 0xc8, 0x20, 0x2d, 0xf8, 0xae, 0xaf, 0x63, 0xd5, 0x38, 0xc9, 0xed,
	0xf8, 0x44, 0xa9, 0x3a, 0x38, 0xac, 0x7a, 0xea, 0x7b, 0xa1, 0xb6, 0x45, 0x11, 0x7b, 0xa4, 0xae,
	0x99, 0x18, 0x31, 0x01, 0x39, 0x88, 0x89, 0xf5, 0xa3, 0xbe, 0xc7, 0xa2, 0x1d, 0xbb, 0xd3, 0xc4,
	0xc5, 0xbe, 0xda, 0x19, 0x88, 0xbc, 0x87, 0x25, 0x59, 0x23, 0xfb, 0xf5, 0xc3, 0x68, 0x63, 0xef,
	0xe6, 0xe4, 0x83, 0x25, 0xb9, 0x6a, 0x15, 0x0f, 0xd5, 0x6d, 0x8a, 0x50, 0x12, 0x56, 0x0e, 0x18,
	0x9b, 0xca, 0x58, 0x34, 0xd3, 0x0a, 0xa1, 0xf5, 0x94, 0xf4, 0x8f, 0xa3, 0x8d, 0xbd, 0xad, 0xc9,
	0x2e, 0xc5, 0xa3, 0xec, 0x4f, 0x5e, 0xd9, 0x7f, 0x7a, 0x4a, 0xc5, 0x03, 0xf5, 0xe1, 0xda, 0xda,
	0xfa, 0xd4, 0xc9, 0xcc, 0xa2, 0xfe, 0x49, 0x12, 0xde, 0x5a, 0xb5, 0x1f, 0x67, 0xf7, 0xa5, 0x99,
	0xc5, 0xe2, 0x2b, 0xb5, 0x63, 0x43, 0xdd, 0x84, 0x48, 0x09, 0xc1, 0x56, 0x26, 0x46, 0xb0, 0xc1,
	0xa1, 0xfe, 0x59, 0x26, 0x15, 0x97, 0xde, 0x61, 0xb6, 0x0e, 0x83, 0xc3, 0x7c, 0xb8, 0x6a, 0xac,
	0x03, 0x77, 0x10, 0xe9, 0x02, 0xf5, 0x63, 0x01, 0x55, 0x2f, 0x9d, 0xd2, 0x05, 0x16, 0x8f, 0xd4,
	0x47, 0xe8, 0x17, 0xc4, 0xc1, 0xd7, 0xe8, 0x93, 0xa9, 0xa0, 0x0e, 0x9e, 0x52, 0x60, 0x68, 0x4c,
	0x9a, 0xeb, 0x5f, 0xa4, 0x51, 0x7a, 0x8d, 0x78, 0xd1, 0x03, 0x7f, 0x98, 0x34, 0x2f, 0x76, 0xd4,
	0x15, 0x53, 0x91, 0x89, 0xfa, 0xb0, 0xbf, 0x11, 0x32, 0xc8, 0xbb, 0x31, 0xe3, 0xd0, 0x36, 0x50,
	0x56, 0x66, 0xa6, 0x9f, 0x48, 0x2b, 0xb6, 0x45, 0x39, 0xaa, 0xcc, 0x2c, 0xaf, 0xc2, 0xe3, 0x19,
	0x38, 0x5c, 0x90, 0x74, 0x72, 0x38, 0x05, 0x4f, 0xfb, 0x55, 0x78, 0x3c, 0x7b, 0x32, 0x58, 0xcb,
	0xb3, 0x70, 0xa2, 0x3e, 0x6f, 0xe6, 0x5d, 0x24, 0x6b, 0x2a, 0xa8, 0x4c, 0x87, 0xf9, 0xa9, 0x48,
	0xc8, 0xa5, 0xb1, 0x08, 0x75, 0x70, 0x6d, 0x85, 0xfd, 0xdd, 0x39, 0x92, 0x08, 0xa3, 0x81, 0xfd,
	0x3d, 0xa3, 0xc7, 0x03, 0xf9, 0x42, 0x40, 0xb9, 0x51, 0x5f, 0xa8, 0x1b, 0xad, 0x67, 0xb4, 0x61,
	0xe6, 0xe9, 0x02, 0x1d, 0x94, 0xdc, 0xea, 0x5f, 0xa5, 0xcc, 0xeb, 0xab, 0xfa, 0x11, 0xb7, 0xc5,
	0x9e, 0xba, 0xce, 0xe8, 0x5a, 0xef, 0x8c, 0xb7, 0x5d, 0x4c, 0x26, 0xa1, 0x7e, 0x26, 0x59, 0x5e,
	0x97, 0x8b, 0x5d, 0xb5, 0x69, 0xb1, 0x09, 0x9c, 0xf4, 0xb1, 0x84, 0x5a, 0x8e, 0x8a, 0x8f, 0xd5,
	0xf6, 0x39, 0x43, 0xb4, 0xa1, 0x41, 0xa7, 0x7f, 0x13, 0x6b, 0xeb, 0x9c, 0x4f, 0x65, 0x9c, 0xcd,
	0xd6, 0xd3, 0xbf, 0x2d, 0x02, 0x39, 0xfd, 0x5c, 0x02, 0x6f, 0xf5, 0xc2, 0xb1, 0x9b, 0x6e, 0xca,
	0x23, 0x79, 0xef, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x23, 0x9a, 0x0b, 0x47, 0x05, 0x00,
	0x00,
}
