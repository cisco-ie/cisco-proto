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
// source: envmon_sensor_info_xml.proto

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_tsi10s_tsi10_tsi11s_tsi11_attributes_env_sensor_info_xml

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

type EnvmonSensorInfoXml_KEYS struct {
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
	Name_12              string   `protobuf:"bytes,13,opt,name=name_12,json=name12,proto3" json:"name_12,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvmonSensorInfoXml_KEYS) Reset()         { *m = EnvmonSensorInfoXml_KEYS{} }
func (m *EnvmonSensorInfoXml_KEYS) String() string { return proto.CompactTextString(m) }
func (*EnvmonSensorInfoXml_KEYS) ProtoMessage()    {}
func (*EnvmonSensorInfoXml_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a5d855b51dbaa9, []int{0}
}

func (m *EnvmonSensorInfoXml_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvmonSensorInfoXml_KEYS.Unmarshal(m, b)
}
func (m *EnvmonSensorInfoXml_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvmonSensorInfoXml_KEYS.Marshal(b, m, deterministic)
}
func (m *EnvmonSensorInfoXml_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvmonSensorInfoXml_KEYS.Merge(m, src)
}
func (m *EnvmonSensorInfoXml_KEYS) XXX_Size() int {
	return xxx_messageInfo_EnvmonSensorInfoXml_KEYS.Size(m)
}
func (m *EnvmonSensorInfoXml_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvmonSensorInfoXml_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_EnvmonSensorInfoXml_KEYS proto.InternalMessageInfo

func (m *EnvmonSensorInfoXml_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_4() string {
	if m != nil {
		return m.Name_4
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_5() string {
	if m != nil {
		return m.Name_5
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_6() string {
	if m != nil {
		return m.Name_6
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_7() string {
	if m != nil {
		return m.Name_7
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_8() string {
	if m != nil {
		return m.Name_8
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_9() string {
	if m != nil {
		return m.Name_9
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_10() string {
	if m != nil {
		return m.Name_10
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_11() string {
	if m != nil {
		return m.Name_11
	}
	return ""
}

func (m *EnvmonSensorInfoXml_KEYS) GetName_12() string {
	if m != nil {
		return m.Name_12
	}
	return ""
}

type EnvmonThresholdType struct {
	ThresholdSeverity    string   `protobuf:"bytes,1,opt,name=threshold_severity,json=thresholdSeverity,proto3" json:"threshold_severity,omitempty"`
	ThresholdRelation    string   `protobuf:"bytes,2,opt,name=threshold_relation,json=thresholdRelation,proto3" json:"threshold_relation,omitempty"`
	ThresholdValue       int32    `protobuf:"zigzag32,3,opt,name=threshold_value,json=thresholdValue,proto3" json:"threshold_value,omitempty"`
	ThresholdName        string   `protobuf:"bytes,4,opt,name=threshold_name,json=thresholdName,proto3" json:"threshold_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvmonThresholdType) Reset()         { *m = EnvmonThresholdType{} }
func (m *EnvmonThresholdType) String() string { return proto.CompactTextString(m) }
func (*EnvmonThresholdType) ProtoMessage()    {}
func (*EnvmonThresholdType) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a5d855b51dbaa9, []int{1}
}

func (m *EnvmonThresholdType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvmonThresholdType.Unmarshal(m, b)
}
func (m *EnvmonThresholdType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvmonThresholdType.Marshal(b, m, deterministic)
}
func (m *EnvmonThresholdType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvmonThresholdType.Merge(m, src)
}
func (m *EnvmonThresholdType) XXX_Size() int {
	return xxx_messageInfo_EnvmonThresholdType.Size(m)
}
func (m *EnvmonThresholdType) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvmonThresholdType.DiscardUnknown(m)
}

var xxx_messageInfo_EnvmonThresholdType proto.InternalMessageInfo

func (m *EnvmonThresholdType) GetThresholdSeverity() string {
	if m != nil {
		return m.ThresholdSeverity
	}
	return ""
}

func (m *EnvmonThresholdType) GetThresholdRelation() string {
	if m != nil {
		return m.ThresholdRelation
	}
	return ""
}

func (m *EnvmonThresholdType) GetThresholdValue() int32 {
	if m != nil {
		return m.ThresholdValue
	}
	return 0
}

func (m *EnvmonThresholdType) GetThresholdName() string {
	if m != nil {
		return m.ThresholdName
	}
	return ""
}

type EnvmonThresholdInfoXml struct {
	ThresholdArray       []*EnvmonThresholdType `protobuf:"bytes,1,rep,name=threshold_array,json=thresholdArray,proto3" json:"threshold_array,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EnvmonThresholdInfoXml) Reset()         { *m = EnvmonThresholdInfoXml{} }
func (m *EnvmonThresholdInfoXml) String() string { return proto.CompactTextString(m) }
func (*EnvmonThresholdInfoXml) ProtoMessage()    {}
func (*EnvmonThresholdInfoXml) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a5d855b51dbaa9, []int{2}
}

func (m *EnvmonThresholdInfoXml) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvmonThresholdInfoXml.Unmarshal(m, b)
}
func (m *EnvmonThresholdInfoXml) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvmonThresholdInfoXml.Marshal(b, m, deterministic)
}
func (m *EnvmonThresholdInfoXml) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvmonThresholdInfoXml.Merge(m, src)
}
func (m *EnvmonThresholdInfoXml) XXX_Size() int {
	return xxx_messageInfo_EnvmonThresholdInfoXml.Size(m)
}
func (m *EnvmonThresholdInfoXml) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvmonThresholdInfoXml.DiscardUnknown(m)
}

var xxx_messageInfo_EnvmonThresholdInfoXml proto.InternalMessageInfo

func (m *EnvmonThresholdInfoXml) GetThresholdArray() []*EnvmonThresholdType {
	if m != nil {
		return m.ThresholdArray
	}
	return nil
}

type EnvmonSensorInfoXml struct {
	Description          string                  `protobuf:"bytes,50,opt,name=description,proto3" json:"description,omitempty"`
	Units                string                  `protobuf:"bytes,51,opt,name=units,proto3" json:"units,omitempty"`
	Value                string                  `protobuf:"bytes,52,opt,name=value,proto3" json:"value,omitempty"`
	DataType             string                  `protobuf:"bytes,53,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Status               string                  `protobuf:"bytes,54,opt,name=status,proto3" json:"status,omitempty"`
	UpdateRate           uint32                  `protobuf:"varint,55,opt,name=update_rate,json=updateRate,proto3" json:"update_rate,omitempty"`
	Threshold            *EnvmonThresholdInfoXml `protobuf:"bytes,56,opt,name=threshold,proto3" json:"threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *EnvmonSensorInfoXml) Reset()         { *m = EnvmonSensorInfoXml{} }
func (m *EnvmonSensorInfoXml) String() string { return proto.CompactTextString(m) }
func (*EnvmonSensorInfoXml) ProtoMessage()    {}
func (*EnvmonSensorInfoXml) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a5d855b51dbaa9, []int{3}
}

func (m *EnvmonSensorInfoXml) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvmonSensorInfoXml.Unmarshal(m, b)
}
func (m *EnvmonSensorInfoXml) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvmonSensorInfoXml.Marshal(b, m, deterministic)
}
func (m *EnvmonSensorInfoXml) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvmonSensorInfoXml.Merge(m, src)
}
func (m *EnvmonSensorInfoXml) XXX_Size() int {
	return xxx_messageInfo_EnvmonSensorInfoXml.Size(m)
}
func (m *EnvmonSensorInfoXml) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvmonSensorInfoXml.DiscardUnknown(m)
}

var xxx_messageInfo_EnvmonSensorInfoXml proto.InternalMessageInfo

func (m *EnvmonSensorInfoXml) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EnvmonSensorInfoXml) GetUnits() string {
	if m != nil {
		return m.Units
	}
	return ""
}

func (m *EnvmonSensorInfoXml) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *EnvmonSensorInfoXml) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *EnvmonSensorInfoXml) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *EnvmonSensorInfoXml) GetUpdateRate() uint32 {
	if m != nil {
		return m.UpdateRate
	}
	return 0
}

func (m *EnvmonSensorInfoXml) GetThreshold() *EnvmonThresholdInfoXml {
	if m != nil {
		return m.Threshold
	}
	return nil
}

func init() {
	proto.RegisterType((*EnvmonSensorInfoXml_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.env_sensor_info_xml.envmon_sensor_info_xml_KEYS")
	proto.RegisterType((*EnvmonThresholdType)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.env_sensor_info_xml.envmon_threshold_type")
	proto.RegisterType((*EnvmonThresholdInfoXml)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.env_sensor_info_xml.envmon_threshold_info_xml")
	proto.RegisterType((*EnvmonSensorInfoXml)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.env_sensor_info_xml.envmon_sensor_info_xml")
}

func init() { proto.RegisterFile("envmon_sensor_info_xml.proto", fileDescriptor_48a5d855b51dbaa9) }

var fileDescriptor_48a5d855b51dbaa9 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x95, 0x3f, 0x8e, 0x13, 0x31,
	0x14, 0xc6, 0x35, 0xc9, 0x26, 0xbb, 0x71, 0x08, 0x68, 0x2d, 0x76, 0x31, 0x5a, 0x24, 0xa2, 0x48,
	0x88, 0x34, 0x8c, 0x32, 0x33, 0xf9, 0x5b, 0x52, 0x50, 0x21, 0x51, 0xcc, 0x22, 0x24, 0x2a, 0xcb,
	0x9b, 0x78, 0xc9, 0x88, 0x89, 0x3d, 0xb2, 0x9d, 0x61, 0x73, 0x0e, 0x3a, 0x4e, 0xc2, 0x05, 0xa8,
	0xe8, 0xe8, 0xb8, 0x05, 0x05, 0x07, 0x40, 0x7e, 0x33, 0x6b, 0x07, 0xc8, 0x09, 0xa0, 0xf9, 0xf2,
	0xf9, 0xf7, 0x59, 0x71, 0xde, 0x7b, 0xb6, 0x82, 0x1e, 0x71, 0x51, 0x6e, 0xa4, 0xa0, 0x9a, 0x0b,
	0x2d, 0x15, 0xcd, 0xc4, 0xb5, 0xa4, 0x37, 0x9b, 0x3c, 0x2c, 0x94, 0x34, 0x12, 0x7f, 0x0f, 0x96,
	0x99, 0x5e, 0x4a, 0x9a, 0x49, 0x4d, 0x6f, 0x6c, 0x5a, 0x6e, 0xde, 0x29, 0x2a, 0x0b, 0xae, 0xc2,
	0x4c, 0x94, 0x5c, 0x18, 0xa9, 0x76, 0xa1, 0x62, 0xcb, 0xf7, 0x1a, 0x34, 0x2c, 0xe4, 0x07, 0xae,
	0xf4, 0x9a, 0xe7, 0xd7, 0xa1, 0xce, 0xa5, 0x09, 0x8d, 0xce, 0x22, 0x0d, 0x6a, 0x25, 0x06, 0x1b,
	0x5b, 0x49, 0xc0, 0x26, 0x56, 0xc6, 0x60, 0xc7, 0x56, 0x26, 0x60, 0x27, 0x56, 0xa6, 0x60, 0xa7,
	0x56, 0x66, 0x60, 0x67, 0x56, 0xe6, 0x60, 0xe7, 0x56, 0x16, 0x60, 0x17, 0xf0, 0xe5, 0xa3, 0xea,
	0x8c, 0x11, 0x68, 0x7d, 0x60, 0x14, 0x32, 0x63, 0x54, 0x76, 0xb5, 0x35, 0x5c, 0x87, 0x5c, 0x94,
	0x7f, 0x56, 0x38, 0xf8, 0xda, 0x40, 0x17, 0x87, 0x8b, 0xa7, 0x2f, 0x5f, 0xbc, 0xbd, 0xc4, 0x18,
	0x1d, 0x09, 0xb6, 0xe1, 0x24, 0xe8, 0x07, 0xc3, 0x4e, 0x0a, 0x1e, 0x9f, 0xa1, 0xb6, 0xfd, 0xa4,
	0x11, 0x69, 0x00, 0x6d, 0xd9, 0x55, 0xe4, 0x70, 0x4c, 0x9a, 0x1e, 0xc7, 0x0e, 0x27, 0xe4, 0xc8,
	0xe3, 0xc4, 0xe1, 0x31, 0x69, 0x79, 0x3c, 0x76, 0x78, 0x42, 0xda, 0x1e, 0x4f, 0x1c, 0x9e, 0x92,
	0x63, 0x8f, 0xa7, 0x0e, 0xcf, 0xc8, 0x89, 0xc7, 0x33, 0x87, 0xe7, 0xa4, 0xe3, 0xf1, 0xdc, 0xe1,
	0x05, 0x41, 0x1e, 0x2f, 0xf0, 0x03, 0x74, 0x5c, 0x95, 0x33, 0x22, 0x5d, 0xe0, 0xb0, 0x2b, 0x1a,
	0xf9, 0x20, 0x22, 0x77, 0xf6, 0x82, 0xc8, 0x07, 0x31, 0xe9, 0xed, 0x05, 0xf1, 0xe0, 0x4b, 0x80,
	0xce, 0xea, 0x6e, 0x9a, 0xb5, 0xe2, 0x7a, 0x2d, 0xf3, 0x15, 0x35, 0xbb, 0x82, 0xe3, 0x67, 0x08,
	0x7b, 0xa2, 0x79, 0xc9, 0x55, 0x66, 0x76, 0x75, 0x57, 0x4f, 0x5d, 0x72, 0x59, 0x07, 0xbf, 0x6f,
	0x57, 0x3c, 0x67, 0x26, 0x93, 0xa2, 0x6e, 0xb7, 0xdf, 0x9e, 0xd6, 0x01, 0x7e, 0x8a, 0xee, 0xf9,
	0xed, 0x25, 0xcb, 0xb7, 0x1c, 0x66, 0x70, 0x9a, 0xde, 0x75, 0xf8, 0x8d, 0xa5, 0xf8, 0x09, 0xf2,
	0x84, 0xc2, 0x60, 0xab, 0xa1, 0xf4, 0x1c, 0x7d, 0xc5, 0x36, 0x7c, 0xf0, 0xb9, 0x81, 0x1e, 0xfe,
	0x55, 0xc7, 0xed, 0xc5, 0xc0, 0x9f, 0x1a, 0xfb, 0xc7, 0x31, 0xa5, 0x98, 0xad, 0xa4, 0x39, 0xec,
	0xc6, 0x3f, 0x82, 0xf0, 0x9f, 0x7d, 0x2a, 0xe1, 0xc1, 0xc1, 0xee, 0x75, 0xf8, 0xb9, 0x6d, 0xc4,
	0xe0, 0x5b, 0x13, 0x9d, 0x1f, 0x7e, 0x50, 0xb8, 0x8f, 0xba, 0x2b, 0xae, 0x97, 0x2a, 0x2b, 0x60,
	0x9a, 0x31, 0x74, 0x7e, 0x1f, 0xe1, 0xfb, 0xa8, 0xb5, 0x15, 0x99, 0xd1, 0x24, 0xa9, 0x2e, 0x28,
	0x2c, 0x2c, 0xad, 0x66, 0x3a, 0xae, 0x28, 0x2c, 0xf0, 0x05, 0xea, 0xac, 0x98, 0x61, 0xf0, 0x2b,
	0xc8, 0x04, 0x92, 0x13, 0x0b, 0x5e, 0xdb, 0xeb, 0x76, 0x8e, 0xda, 0xda, 0x30, 0xb3, 0xd5, 0x64,
	0x5a, 0x5d, 0xd0, 0x6a, 0x85, 0x1f, 0xa3, 0xee, 0xb6, 0x58, 0x31, 0xc3, 0xa9, 0x62, 0x86, 0x93,
	0x59, 0x3f, 0x18, 0xf6, 0x52, 0x54, 0xa1, 0x94, 0x19, 0x8e, 0x3f, 0x36, 0x50, 0xc7, 0x55, 0x44,
	0xe6, 0xfd, 0x60, 0xd8, 0x8d, 0x7f, 0xfe, 0x57, 0x53, 0xbd, 0x4d, 0x52, 0xdf, 0x87, 0xab, 0x36,
	0xfc, 0x11, 0x24, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x24, 0xae, 0xff, 0x28, 0x06, 0x00,
	0x00,
}