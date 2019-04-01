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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_attributes_inv_basic_bag

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
	AllocatedPower                   int32    `protobuf:"zigzag32,76,opt,name=allocated_power,json=allocatedPower,proto3" json:"allocated_power,omitempty"`
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

func (m *InvPhyEntityBasicInfo) GetAllocatedPower() int32 {
	if m != nil {
		return m.AllocatedPower
	}
	return 0
}

func init() {
	proto.RegisterType((*InvPhyEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.attributes.inv_basic_bag.inv_phy_entity_basic_info_KEYS")
	proto.RegisterType((*InvPhyEntityBasicInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.attributes.inv_basic_bag.inv_phy_entity_basic_info")
}

func init() { proto.RegisterFile("inv_phy_entity_basic_info.proto", fileDescriptor_94f1af735311b03a) }

var fileDescriptor_94f1af735311b03a = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x6f, 0x1c, 0x35,
	0x10, 0xc7, 0x15, 0xa0, 0x51, 0xce, 0x2d, 0xa4, 0x35, 0x69, 0x70, 0x41, 0xd0, 0x53, 0x41, 0x22,
	0x08, 0xe9, 0x44, 0xd2, 0x52, 0x28, 0x94, 0x1f, 0x25, 0x6d, 0x68, 0x68, 0x1b, 0x55, 0x97, 0xf2,
	0xc0, 0x93, 0xe5, 0xb3, 0x67, 0xef, 0x46, 0x78, 0xed, 0x65, 0xec, 0xbd, 0x64, 0xf3, 0xc4, 0x9f,
	0x8e, 0x3c, 0x7b, 0x77, 0xbd, 0x14, 0xf5, 0x65, 0xef, 0xfc, 0xfd, 0x7e, 0x3c, 0x33, 0x1e, 0xff,
	0x10, 0xb7, 0x31, 0xcc, 0x75, 0x33, 0xeb, 0x34, 0x84, 0x8c, 0xb9, 0xd3, 0x13, 0x93, 0xd0, 0x6a,
	0x0c, 0x55, 0x1c, 0x35, 0x14, 0x73, 0x94, 0xc6, 0x62, 0xb2, 0x51, 0x63, 0x4c, 0xfa, 0x9c, 0x34,
	0x86, 0x79, 0x3d, 0x25, 0x1d, 0x1b, 0xa0, 0x11, 0x86, 0x39, 0x84, 0x1c, 0xa9, 0x1b, 0x91, 0xb1,
	0x7f, 0x27, 0xfe, 0x8e, 0x9a, 0x78, 0x06, 0x94, 0x66, 0xe0, 0xab, 0x51, 0xf2, 0x31, 0x8f, 0x72,
	0xc2, 0xfd, 0xc4, 0xdf, 0x91, 0xc9, 0x99, 0x70, 0xd2, 0x66, 0x48, 0x65, 0xea, 0x22, 0xd3, 0xc4,
	0x4c, 0xef, 0x4c, 0xc4, 0x67, 0x6f, 0xad, 0x42, 0x3f, 0x7b, 0xf2, 0xd7, 0xa9, 0x94, 0xe2, 0xbd,
	0x60, 0x6a, 0x50, 0x1b, 0xc3, 0x8d, 0xbd, 0xc1, 0x98, 0xff, 0xcb, 0x9b, 0x62, 0xb3, 0xfc, 0xea,
	0x7d, 0xf5, 0x0e, 0xab, 0x57, 0xca, 0x68, 0x7f, 0x25, 0x1f, 0xa8, 0x77, 0x5f, 0xcb, 0x07, 0x77,
	0xfe, 0x1d, 0x88, 0x5b, 0x6f, 0x4d, 0x22, 0x87, 0xe2, 0xaa, 0x83, 0x64, 0x09, 0x9b, 0x8c, 0x31,
	0xa8, 0x03, 0x9e, 0xb9, 0x2e, 0xc9, 0xdb, 0xe2, 0xea, 0x1c, 0x82, 0x8b, 0xa4, 0x73, 0xd7, 0x80,
	0xba, 0xcb, 0x84, 0xe8, 0xa5, 0x57, 0x5d, 0x03, 0xab, 0x12, 0xef, 0xad, 0x95, 0xf8, 0xb5, 0xb8,
	0x31, 0x33, 0xe4, 0xce, 0x0c, 0x81, 0x26, 0x98, 0x63, 0x2a, 0xc1, 0xbf, 0x65, 0xe0, 0xfa, 0xd2,
	0x18, 0x2f, 0xf4, 0x02, 0x57, 0x48, 0xf5, 0x65, 0xf8, 0x7e, 0x0f, 0x2f, 0x8d, 0x75, 0x38, 0xc5,
	0x2a, 0x5f, 0x86, 0xbf, 0xeb, 0xe1, 0xa5, 0xb1, 0x82, 0xef, 0x89, 0x5d, 0x3b, 0xc3, 0x46, 0xff,
	0xbf, 0x96, 0xef, 0x79, 0xc6, 0x4e, 0x71, 0x9f, 0xbe, 0x59, 0xcf, 0xe7, 0xe2, 0xfd, 0x04, 0x84,
	0xc6, 0xeb, 0xd0, 0xd6, 0x13, 0x20, 0xf5, 0x80, 0xe1, 0x6b, 0xbd, 0x78, 0xc2, 0x5a, 0xa9, 0xa3,
	0x36, 0xa1, 0xad, 0x8c, 0xcd, 0x2d, 0x01, 0x69, 0x6e, 0xc1, 0x0f, 0x7d, 0x1d, 0xeb, 0xc6, 0x49,
	0x69, 0xc7, 0xa7, 0x42, 0xd4, 0xd1, 0x81, 0xef, 0xa9, 0x1f, 0x99, 0x1a, 0xb0, 0xc2, 0xf6, 0x50,
	0x5c, 0x33, 0x29, 0x41, 0xd6, 0xe8, 0x74, 0xca, 0xa4, 0x1e, 0xf6, 0x3d, 0x66, 0xed, 0xd8, 0x9d,
	0x66, 0x92, 0xfb, 0x62, 0x67, 0x49, 0x94, 0x3d, 0xac, 0xd0, 0x1a, 0xde, 0xaf, 0x9f, 0x86, 0x1b,
	0x7b, 0x37, 0xc6, 0x1f, 0x2e, 0xc8, 0x75, 0x4b, 0x3e, 0x10, 0xb7, 0x30, 0xe9, 0x0a, 0xc1, 0x3b,
	0x4d, 0xd0, 0x78, 0x63, 0xc1, 0x4c, 0x3c, 0xe8, 0x36, 0x60, 0x56, 0x3f, 0x0f, 0x37, 0xf6, 0xb6,
	0xc6, 0xbb, 0x98, 0x8e, 0x8a, 0x3f, 0x7e, 0x6d, 0xff, 0x19, 0x30, 0xcb, 0xfb, 0xe2, 0xa3, 0x4b,
	0x6b, 0xeb, 0x53, 0x67, 0x33, 0x4d, 0xea, 0x17, 0x4e, 0x78, 0x73, 0xdd, 0x7e, 0x54, 0xdc, 0x57,
	0x66, 0x9a, 0xe4, 0x37, 0x62, 0xc7, 0xc6, 0xba, 0x89, 0x09, 0x33, 0x68, 0xeb, 0x4d, 0x4a, 0xda,
	0x46, 0x07, 0xea, 0x57, 0x9e, 0x24, 0x57, 0xde, 0x61, 0xb1, 0x0e, 0xa3, 0x83, 0x72, 0xb8, 0x6a,
	0xa8, 0x23, 0x75, 0x3a, 0xe1, 0x05, 0xa8, 0x47, 0x0c, 0x8a, 0x5e, 0x3a, 0xc5, 0x0b, 0x90, 0x0f,
	0xc5, 0xc7, 0x10, 0xe6, 0x48, 0x31, 0xd4, 0x10, 0xb2, 0xf1, 0xba, 0x8e, 0x01, 0x73, 0x24, 0xdd,
	0x98, 0x3c, 0x53, 0xbf, 0x71, 0xa3, 0xd4, 0x25, 0xe2, 0x45, 0x0f, 0xbc, 0x34, 0x79, 0x26, 0x77,
	0xc4, 0x15, 0xe3, 0xd1, 0x24, 0x75, 0xd8, 0xdf, 0x08, 0x1e, 0x94, 0xdd, 0x98, 0x52, 0x6c, 0x1b,
	0x5d, 0x79, 0x33, 0x55, 0x8f, 0xb9, 0x15, 0x03, 0x56, 0x8e, 0xbc, 0x99, 0x96, 0x55, 0x04, 0x38,
	0xd3, 0x0e, 0xe6, 0xc8, 0x9d, 0x5c, 0x9e, 0x82, 0x27, 0xfd, 0x2a, 0x02, 0x9c, 0x3d, 0x5e, 0x5a,
	0x8b, 0xb3, 0x70, 0x22, 0xbe, 0x68, 0x66, 0x5d, 0x42, 0x6b, 0xbc, 0xf6, 0xa6, 0x83, 0xf2, 0x5a,
	0x64, 0xa0, 0xca, 0x58, 0xd0, 0x75, 0x74, 0xad, 0x87, 0xfe, 0xee, 0x1c, 0x71, 0x84, 0xe1, 0x92,
	0x7d, 0x5e, 0xd0, 0xe3, 0x25, 0xf9, 0x82, 0x41, 0xbe, 0x51, 0x5f, 0x89, 0xeb, 0x6d, 0x20, 0xb0,
	0x71, 0x1a, 0xf0, 0x02, 0x9c, 0xae, 0xa8, 0x55, 0xbf, 0x73, 0x99, 0xdb, 0xeb, 0xfa, 0x11, 0xb5,
	0x72, 0x4f, 0x6c, 0x13, 0xb8, 0x36, 0x38, 0x13, 0x6c, 0x97, 0xb2, 0xc9, 0xa0, 0x9e, 0x72, 0x96,
	0x37, 0x65, 0xb9, 0x2b, 0x36, 0x2d, 0x34, 0x91, 0xb2, 0x3a, 0xe6, 0x50, 0x8b, 0x91, 0xfc, 0x44,
	0x0c, 0xce, 0x49, 0x27, 0x1b, 0x1b, 0x70, 0xea, 0x0f, 0xb6, 0xb6, 0xce, 0xe9, 0x94, 0xc7, 0xc5,
	0x6c, 0x03, 0xfe, 0xd3, 0x82, 0x46, 0xa7, 0x9e, 0x71, 0xe0, 0xad, 0x5e, 0x38, 0x76, 0xf2, 0x4b,
	0xb1, 0x6d, 0xbc, 0x8f, 0xd6, 0x64, 0x70, 0x9a, 0x9f, 0x3f, 0xf5, 0x9c, 0x91, 0x0f, 0x56, 0xf2,
	0xcb, 0xa2, 0x4e, 0x36, 0xf9, 0x41, 0xbd, 0xfb, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xbe,
	0x52, 0x26, 0x73, 0x05, 0x00, 0x00,
}
