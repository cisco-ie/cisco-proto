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
	proto.RegisterType((*InvPhyEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.attributes.inv_basic_bag.inv_phy_entity_basic_info_KEYS")
	proto.RegisterType((*InvPhyEntityBasicInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.attributes.inv_basic_bag.inv_phy_entity_basic_info")
}

func init() { proto.RegisterFile("inv_phy_entity_basic_info.proto", fileDescriptor_94f1af735311b03a) }

var fileDescriptor_94f1af735311b03a = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x6f, 0x1c, 0x35,
	0x10, 0xc7, 0x15, 0xa0, 0x51, 0xce, 0x2d, 0xa4, 0x35, 0x69, 0x70, 0x41, 0xd0, 0x53, 0x41, 0x22,
	0x08, 0xe9, 0x44, 0xd2, 0x52, 0x28, 0x94, 0x1f, 0x25, 0x6d, 0x68, 0x68, 0x1b, 0x55, 0x97, 0xf2,
	0xc0, 0x03, 0xb2, 0x7c, 0xf6, 0xec, 0xdd, 0x88, 0x5d, 0x7b, 0x19, 0x7b, 0x2f, 0xd9, 0x3c, 0xf1,
	0xa7, 0x23, 0xcf, 0xde, 0x5e, 0xef, 0x5a, 0xe5, 0x65, 0xef, 0xfc, 0xfd, 0x7e, 0x3c, 0x33, 0x1e,
	0xff, 0x10, 0xb7, 0xd1, 0xcf, 0x75, 0x3d, 0x6b, 0x35, 0xf8, 0x84, 0xa9, 0xd5, 0x13, 0x13, 0xd1,
	0x6a, 0xf4, 0x45, 0x18, 0xd5, 0x14, 0x52, 0x90, 0x7f, 0x5b, 0x8c, 0x36, 0x68, 0x0c, 0x51, 0x9f,
	0x93, 0x46, 0x3f, 0xaf, 0xa6, 0xa4, 0x43, 0x0d, 0x34, 0x42, 0x3f, 0x07, 0x9f, 0x02, 0xb5, 0x23,
	0x32, 0xf6, 0x9f, 0xc8, 0xdf, 0x51, 0x17, 0x66, 0x14, 0xcb, 0x90, 0x46, 0x29, 0xe2, 0x7e, 0xe4,
	0xef, 0xc8, 0xa4, 0x44, 0x38, 0x69, 0x12, 0xc4, 0x3c, 0x6d, 0x91, 0x65, 0x62, 0xa6, 0x77, 0x26,
	0xe2, 0xb3, 0x4b, 0x2b, 0xd0, 0xcf, 0x9e, 0xfc, 0x75, 0x2a, 0xa5, 0x78, 0xcf, 0x9b, 0x0a, 0xd4,
	0xc6, 0x70, 0x63, 0x6f, 0x30, 0xe6, 0xff, 0xf2, 0xa6, 0xd8, 0xcc, 0xbf, 0x7a, 0x5f, 0xbd, 0xc3,
	0xea, 0x95, 0x3c, 0xda, 0x5f, 0xca, 0x07, 0xea, 0xdd, 0xd7, 0xf2, 0xc1, 0x9d, 0xff, 0x06, 0xe2,
	0xd6, 0xa5, 0x49, 0xe4, 0x50, 0x5c, 0x75, 0x10, 0x2d, 0x61, 0x9d, 0x30, 0x78, 0x75, 0xc0, 0x33,
	0x57, 0x25, 0x79, 0x5b, 0x5c, 0x9d, 0x83, 0x77, 0x81, 0x74, 0x6a, 0x6b, 0x50, 0x77, 0x99, 0x10,
	0x9d, 0xf4, 0xaa, 0xad, 0x61, 0x59, 0xe2, 0xbd, 0x95, 0x12, 0xbf, 0x16, 0x37, 0x66, 0x86, 0xdc,
	0x99, 0x21, 0xd0, 0x04, 0x73, 0x8c, 0x39, 0xf8, 0xb7, 0x0c, 0x5c, 0xef, 0x8d, 0xf1, 0x42, 0xcf,
	0x70, 0x81, 0x54, 0xad, 0xc3, 0xf7, 0x3b, 0xb8, 0x37, 0x56, 0xe1, 0x18, 0x8a, 0xb4, 0x0e, 0x7f,
	0xd7, 0xc1, 0xbd, 0xb1, 0x84, 0xef, 0x89, 0x5d, 0x3b, 0xc3, 0x5a, 0xbf, 0x5d, 0xcb, 0xf7, 0x3c,
	0x63, 0x27, 0xbb, 0x4f, 0xdf, 0xac, 0xe7, 0x73, 0xf1, 0x7e, 0x04, 0x42, 0x53, 0x6a, 0xdf, 0x54,
	0x13, 0x20, 0xf5, 0x80, 0xe1, 0x6b, 0x9d, 0x78, 0xc2, 0x5a, 0xae, 0xa3, 0x32, 0xbe, 0x29, 0x8c,
	0x4d, 0x0d, 0x01, 0x69, 0x6e, 0xc1, 0x0f, 0x5d, 0x1d, 0xab, 0xc6, 0x49, 0x6e, 0xc7, 0xa7, 0x42,
	0x54, 0xc1, 0x41, 0xd9, 0x51, 0x3f, 0x32, 0x35, 0x60, 0x85, 0xed, 0xa1, 0xb8, 0x66, 0x62, 0x84,
	0xa4, 0xd1, 0xe9, 0x98, 0x48, 0x3d, 0xec, 0x7a, 0xcc, 0xda, 0xb1, 0x3b, 0x4d, 0x24, 0xf7, 0xc5,
	0x4e, 0x4f, 0xe4, 0x3d, 0x2c, 0xd0, 0x1a, 0xde, 0xaf, 0x9f, 0x86, 0x1b, 0x7b, 0x37, 0xc6, 0x1f,
	0x2e, 0xc8, 0x55, 0x4b, 0x3e, 0x10, 0xb7, 0x30, 0xea, 0x02, 0xa1, 0x74, 0x9a, 0xa0, 0x2e, 0x8d,
	0x05, 0x33, 0x29, 0x41, 0x37, 0x1e, 0x93, 0xfa, 0x79, 0xb8, 0xb1, 0xb7, 0x35, 0xde, 0xc5, 0x78,
	0x94, 0xfd, 0xf1, 0x6b, 0xfb, 0x4f, 0x8f, 0x49, 0xde, 0x17, 0x1f, 0xad, 0xad, 0xad, 0x4b, 0x9d,
	0xcc, 0x34, 0xaa, 0x5f, 0x38, 0xe1, 0xcd, 0x55, 0xfb, 0x51, 0x76, 0x5f, 0x99, 0x69, 0x94, 0xdf,
	0x88, 0x1d, 0x1b, 0xaa, 0x3a, 0x44, 0x4c, 0xa0, 0x6d, 0x69, 0x62, 0xd4, 0x36, 0x38, 0x50, 0xbf,
	0xf2, 0x24, 0xb9, 0xf4, 0x0e, 0xb3, 0x75, 0x18, 0x1c, 0xe4, 0xc3, 0x55, 0x41, 0x15, 0xa8, 0xd5,
	0x11, 0x2f, 0x40, 0x3d, 0x62, 0x50, 0x74, 0xd2, 0x29, 0x5e, 0x80, 0x7c, 0x28, 0x3e, 0x06, 0x3f,
	0x47, 0x0a, 0xbe, 0x02, 0x9f, 0x4c, 0xa9, 0xab, 0xe0, 0x31, 0x05, 0xd2, 0xb5, 0x49, 0x33, 0xf5,
	0x1b, 0x37, 0x4a, 0xad, 0x11, 0x2f, 0x3a, 0xe0, 0xa5, 0x49, 0x33, 0xb9, 0x23, 0xae, 0x98, 0x12,
	0x4d, 0x54, 0x87, 0xdd, 0x8d, 0xe0, 0x41, 0xde, 0x8d, 0x29, 0x85, 0xa6, 0xd6, 0x45, 0x69, 0xa6,
	0xea, 0x31, 0xb7, 0x62, 0xc0, 0xca, 0x51, 0x69, 0xa6, 0x79, 0x15, 0x1e, 0xce, 0xb4, 0x83, 0x39,
	0x72, 0x27, 0xfb, 0x53, 0xf0, 0xa4, 0x5b, 0x85, 0x87, 0xb3, 0xc7, 0xbd, 0xb5, 0x38, 0x0b, 0x27,
	0xe2, 0x8b, 0x7a, 0xd6, 0x46, 0xb4, 0xa6, 0xd4, 0xa5, 0x69, 0x21, 0xbf, 0x14, 0x09, 0xa8, 0x30,
	0x16, 0x74, 0x15, 0x5c, 0x53, 0x42, 0x77, 0x77, 0x8e, 0x38, 0xc2, 0xb0, 0x67, 0x9f, 0x67, 0xf4,
	0xb8, 0x27, 0x5f, 0x30, 0xc8, 0x37, 0xea, 0x2b, 0x71, 0xbd, 0xf1, 0x04, 0x36, 0x4c, 0x3d, 0x5e,
	0x80, 0xd3, 0x05, 0x35, 0xea, 0x77, 0x2e, 0x73, 0x7b, 0x55, 0x3f, 0xa2, 0x46, 0xee, 0x89, 0x6d,
	0x02, 0xd7, 0x78, 0x67, 0xbc, 0x6d, 0x63, 0x32, 0x09, 0xd4, 0x53, 0xce, 0xf2, 0xa6, 0x2c, 0x77,
	0xc5, 0xa6, 0x85, 0x3a, 0x50, 0x52, 0xc7, 0x1c, 0x6a, 0x31, 0x92, 0x9f, 0x88, 0xc1, 0x39, 0xe9,
	0x68, 0x43, 0x0d, 0x4e, 0xfd, 0xc1, 0xd6, 0xd6, 0x39, 0x9d, 0xf2, 0x38, 0x9b, 0x8d, 0xc7, 0x7f,
	0x1b, 0xd0, 0xe8, 0xd4, 0x33, 0x0e, 0xbc, 0xd5, 0x09, 0xc7, 0x4e, 0x7e, 0x29, 0xb6, 0x4d, 0x59,
	0x06, 0x6b, 0x12, 0x38, 0x5d, 0x87, 0x33, 0x20, 0xf5, 0x9c, 0x91, 0x0f, 0x96, 0xf2, 0xcb, 0xac,
	0x4e, 0x36, 0xf9, 0x31, 0xbd, 0xfb, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xab, 0x76, 0x56,
	0x6f, 0x05, 0x00, 0x00,
}
