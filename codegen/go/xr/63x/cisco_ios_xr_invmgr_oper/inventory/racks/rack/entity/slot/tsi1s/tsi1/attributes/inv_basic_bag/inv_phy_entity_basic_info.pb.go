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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_attributes_inv_basic_bag

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
	proto.RegisterType((*InvPhyEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.attributes.inv_basic_bag.inv_phy_entity_basic_info_KEYS")
	proto.RegisterType((*InvPhyEntityBasicInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.attributes.inv_basic_bag.inv_phy_entity_basic_info")
}

func init() { proto.RegisterFile("inv_phy_entity_basic_info.proto", fileDescriptor_94f1af735311b03a) }

var fileDescriptor_94f1af735311b03a = []byte{
	// 734 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x6f, 0x1c, 0x35,
	0x10, 0xc7, 0x15, 0xa0, 0x51, 0xe2, 0x16, 0xb5, 0x5d, 0xd2, 0xe0, 0x82, 0xa0, 0xa7, 0xc2, 0x43,
	0x10, 0xd2, 0x89, 0xa4, 0xa5, 0x50, 0x28, 0x3f, 0x4a, 0xda, 0xd0, 0x50, 0x1a, 0xa1, 0x4b, 0x79,
	0xe0, 0x01, 0x8d, 0x7c, 0xf6, 0xec, 0xdd, 0x88, 0x5d, 0x7b, 0x19, 0x7b, 0x2f, 0xd9, 0xfc, 0xa3,
	0xfc, 0x3b, 0xc8, 0xb3, 0xb7, 0xe9, 0x5d, 0xab, 0xbe, 0xec, 0x9d, 0xbf, 0xdf, 0x8f, 0x67, 0xc6,
	0xe3, 0x1f, 0xea, 0x0e, 0xf9, 0x05, 0x34, 0xf3, 0x0e, 0xd0, 0x27, 0x4a, 0x1d, 0x4c, 0x4d, 0x24,
	0x0b, 0xe4, 0xcb, 0x30, 0x6e, 0x38, 0xa4, 0x50, 0xfc, 0x6d, 0x29, 0xda, 0x00, 0x14, 0x22, 0x9c,
	0x33, 0x90, 0x5f, 0xd4, 0x33, 0x86, 0xd0, 0x20, 0x8f, 0xc9, 0x2f, 0xd0, 0xa7, 0xc0, 0xdd, 0x98,
	0x8d, 0xfd, 0x27, 0xca, 0x77, 0xdc, 0x87, 0x19, 0xc7, 0x2a, 0xa4, 0x71, 0x8a, 0xb4, 0x1f, 0xe5,
	0x3b, 0x36, 0x29, 0x31, 0x4d, 0xdb, 0x84, 0x31, 0x4f, 0x5b, 0x66, 0x99, 0x9a, 0xd9, 0xdd, 0xa9,
	0xfa, 0xf4, 0xad, 0x15, 0xc0, 0xf3, 0xa7, 0x7f, 0x9d, 0x16, 0x85, 0x7a, 0xcf, 0x9b, 0x1a, 0xf5,
	0xc6, 0x68, 0x63, 0x6f, 0x7b, 0x22, 0xff, 0x8b, 0x5b, 0x6a, 0x33, 0xff, 0xc2, 0xbe, 0x7e, 0x47,
	0xd4, 0x2b, 0x79, 0xb4, 0x7f, 0x29, 0x1f, 0xe8, 0x77, 0x5f, 0xc9, 0x07, 0x77, 0xff, 0xdb, 0x52,
	0xb7, 0xdf, 0x9a, 0xa4, 0x18, 0xa9, 0xab, 0x0e, 0xa3, 0x65, 0x6a, 0x12, 0x05, 0xaf, 0x0f, 0x64,
	0xe6, 0xaa, 0x54, 0xdc, 0x51, 0x57, 0x17, 0xe8, 0x5d, 0x60, 0x48, 0x5d, 0x83, 0xfa, 0x9e, 0x10,
	0xaa, 0x97, 0x5e, 0x76, 0x0d, 0x5e, 0x96, 0x78, 0x7f, 0xa5, 0xc4, 0x2f, 0xd5, 0xcd, 0xb9, 0x61,
	0x77, 0x66, 0x18, 0x81, 0x71, 0x41, 0x31, 0x07, 0xff, 0x5a, 0x80, 0x1b, 0x83, 0x31, 0x59, 0xea,
	0x19, 0x2e, 0x89, 0xeb, 0x75, 0xf8, 0x41, 0x0f, 0x0f, 0xc6, 0x2a, 0x1c, 0x43, 0x99, 0xd6, 0xe1,
	0x6f, 0x7a, 0x78, 0x30, 0x2e, 0xe1, 0xfb, 0x6a, 0xd7, 0xce, 0xa9, 0x81, 0x37, 0x6b, 0xf9, 0x56,
	0x66, 0xec, 0x64, 0xf7, 0xd9, 0xeb, 0xf5, 0x7c, 0xa6, 0xde, 0x8f, 0xc8, 0x64, 0x2a, 0xf0, 0x6d,
	0x3d, 0x45, 0xd6, 0x0f, 0x05, 0xbe, 0xd6, 0x8b, 0x27, 0xa2, 0xe5, 0x3a, 0x6a, 0xe3, 0xdb, 0xd2,
	0xd8, 0xd4, 0x32, 0x32, 0x48, 0x0b, 0xbe, 0xeb, 0xeb, 0x58, 0x35, 0x4e, 0x72, 0x3b, 0x3e, 0x51,
	0xaa, 0x0e, 0x0e, 0xab, 0x9e, 0xfa, 0x5e, 0xa8, 0x6d, 0x51, 0xc4, 0x1e, 0xa9, 0x6b, 0x26, 0x46,
	0x4c, 0x40, 0x0e, 0x62, 0x62, 0xfd, 0xa8, 0xef, 0xb1, 0x68, 0xc7, 0xee, 0x34, 0x71, 0xb1, 0xaf,
	0x76, 0x06, 0x22, 0xef, 0x61, 0x49, 0xd6, 0xc8, 0x7e, 0xfd, 0x30, 0xda, 0xd8, 0xbb, 0x39, 0xf9,
	0x60, 0x49, 0xae, 0x5a, 0xc5, 0x43, 0x75, 0x9b, 0x22, 0x94, 0x84, 0x95, 0x03, 0xc6, 0xa6, 0x32,
	0x16, 0xcd, 0xb4, 0x42, 0x68, 0x3d, 0x25, 0xfd, 0xe3, 0x68, 0x63, 0x6f, 0x6b, 0xb2, 0x4b, 0xf1,
	0x28, 0xfb, 0x93, 0x57, 0xf6, 0x9f, 0x9e, 0x52, 0xf1, 0x40, 0x7d, 0xb8, 0xb6, 0xb6, 0x3e, 0x75,
	0x32, 0xb3, 0xa8, 0x7f, 0x92, 0x84, 0xb7, 0x56, 0xed, 0xc7, 0xd9, 0x7d, 0x69, 0x66, 0xb1, 0xf8,
	0x4a, 0xed, 0xd8, 0x50, 0x37, 0x21, 0x52, 0x42, 0xb0, 0x95, 0x89, 0x11, 0x6c, 0x70, 0xa8, 0x7f,
	0x96, 0x49, 0xc5, 0xa5, 0x77, 0x98, 0xad, 0xc3, 0xe0, 0x30, 0x1f, 0xae, 0x1a, 0xeb, 0xc0, 0x1d,
	0x44, 0xba, 0x40, 0xfd, 0x58, 0x40, 0xd5, 0x4b, 0xa7, 0x74, 0x81, 0xc5, 0x23, 0xf5, 0x11, 0xfa,
	0x05, 0x71, 0xf0, 0x35, 0xfa, 0x64, 0x2a, 0xa8, 0x83, 0xa7, 0x14, 0x18, 0x1a, 0x93, 0xe6, 0xfa,
	0x17, 0x69, 0x94, 0x5e, 0x23, 0x5e, 0xf4, 0xc0, 0x1f, 0x26, 0xcd, 0x8b, 0x1d, 0x75, 0xc5, 0x54,
	0x64, 0xa2, 0x3e, 0xec, 0x6f, 0x84, 0x0c, 0xf2, 0x6e, 0xcc, 0x38, 0xb4, 0x0d, 0x94, 0x95, 0x99,
	0xe9, 0x27, 0xd2, 0x8a, 0x6d, 0x51, 0x8e, 0x2a, 0x33, 0xcb, 0xab, 0xf0, 0x78, 0x06, 0x0e, 0x17,
	0x24, 0x9d, 0x1c, 0x4e, 0xc1, 0xd3, 0x7e, 0x15, 0x1e, 0xcf, 0x9e, 0x0c, 0xd6, 0xf2, 0x2c, 0x9c,
	0xa8, 0xcf, 0x9b, 0x79, 0x17, 0xc9, 0x9a, 0x0a, 0x2a, 0xd3, 0x61, 0x7e, 0x29, 0x12, 0x72, 0x69,
	0x2c, 0x42, 0x1d, 0x5c, 0x5b, 0x61, 0x7f, 0x77, 0x8e, 0x24, 0xc2, 0x68, 0x60, 0x7f, 0xcf, 0xe8,
	0xf1, 0x40, 0xbe, 0x10, 0x50, 0x6e, 0xd4, 0x17, 0xea, 0x46, 0xeb, 0x19, 0x6d, 0x98, 0x79, 0xba,
	0x40, 0x07, 0x25, 0xb7, 0xfa, 0x57, 0x29, 0xf3, 0xfa, 0xaa, 0x7e, 0xc4, 0x6d, 0xb1, 0xa7, 0xae,
	0x33, 0xba, 0xd6, 0x3b, 0xe3, 0x6d, 0x17, 0x93, 0x49, 0xa8, 0x9f, 0x49, 0x96, 0xd7, 0xe5, 0x62,
	0x57, 0x6d, 0x5a, 0x6c, 0x02, 0x27, 0x7d, 0x2c, 0xa1, 0x96, 0xa3, 0xe2, 0x63, 0xb5, 0x7d, 0xce,
	0x10, 0x6d, 0x68, 0xd0, 0xe9, 0xdf, 0xc4, 0xda, 0x3a, 0xe7, 0x53, 0x19, 0x67, 0xb3, 0xf5, 0xf4,
	0x6f, 0x8b, 0x40, 0x4e, 0x3f, 0x97, 0xc0, 0x5b, 0xbd, 0x70, 0xec, 0xa6, 0x9b, 0xf2, 0x46, 0xde,
	0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x12, 0x93, 0x23, 0x98, 0x46, 0x05, 0x00, 0x00,
}
