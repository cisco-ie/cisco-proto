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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_tsi10s_tsi10_attributes_env_sensor_info_xml

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
	proto.RegisterType((*EnvmonSensorInfoXml_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.attributes.env_sensor_info_xml.envmon_sensor_info_xml_KEYS")
	proto.RegisterType((*EnvmonThresholdType)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.attributes.env_sensor_info_xml.envmon_threshold_type")
	proto.RegisterType((*EnvmonThresholdInfoXml)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.attributes.env_sensor_info_xml.envmon_threshold_info_xml")
	proto.RegisterType((*EnvmonSensorInfoXml)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.attributes.env_sensor_info_xml.envmon_sensor_info_xml")
}

func init() { proto.RegisterFile("envmon_sensor_info_xml.proto", fileDescriptor_48a5d855b51dbaa9) }

var fileDescriptor_48a5d855b51dbaa9 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0x4f, 0x6e, 0x13, 0x31,
	0x14, 0xc6, 0x35, 0x49, 0x93, 0x36, 0x0e, 0x05, 0xd5, 0xa2, 0xc5, 0xa8, 0x48, 0x44, 0x91, 0x10,
	0xd9, 0x60, 0x65, 0x66, 0xf2, 0x77, 0xc9, 0x82, 0x15, 0x12, 0x8b, 0x29, 0x42, 0x62, 0x65, 0xb9,
	0x89, 0x4b, 0x2d, 0x12, 0x3b, 0xb2, 0x9d, 0x51, 0xb3, 0x66, 0xc7, 0x49, 0x58, 0x73, 0x04, 0x24,
	0x16, 0xdc, 0x82, 0x05, 0x07, 0x41, 0x7e, 0x93, 0xda, 0x81, 0xe6, 0x04, 0xdd, 0x7c, 0xf9, 0xfc,
	0x7b, 0x4f, 0x99, 0x7c, 0xef, 0x79, 0x82, 0x9e, 0x09, 0x55, 0x2e, 0xb5, 0x62, 0x56, 0x28, 0xab,
	0x0d, 0x93, 0xea, 0x4a, 0xb3, 0x9b, 0xe5, 0x82, 0xae, 0x8c, 0x76, 0x1a, 0xff, 0x4a, 0x66, 0xd2,
	0xce, 0x34, 0x93, 0xda, 0xb2, 0x1b, 0x5f, 0x2d, 0x97, 0x9f, 0x0c, 0xd3, 0x2b, 0x61, 0xa8, 0x54,
	0xa5, 0x50, 0x4e, 0x9b, 0x0d, 0x35, 0x7c, 0xf6, 0xd9, 0x82, 0xd2, 0x2b, 0xae, 0x9c, 0xe1, 0x1b,
	0x6a, 0x17, 0xda, 0x51, 0x67, 0x65, 0x6a, 0x41, 0xbd, 0x64, 0x60, 0x33, 0x2f, 0x39, 0xd8, 0xdc,
	0xcb, 0x00, 0xec, 0xc0, 0xcb, 0x10, 0xec, 0xd0, 0xcb, 0x08, 0xec, 0xc8, 0xcb, 0x18, 0xec, 0xd8,
	0xcb, 0x04, 0xec, 0xc4, 0xcb, 0x14, 0xec, 0x14, 0xbe, 0xbc, 0x5f, 0x3d, 0xa3, 0x4f, 0xb9, 0x73,
	0x46, 0x5e, 0xae, 0x9d, 0xb0, 0x54, 0xa8, 0xf2, 0xff, 0x44, 0xdd, 0xef, 0x35, 0x74, 0xbe, 0x3f,
	0x2c, 0x7b, 0xfb, 0xe6, 0xe3, 0x05, 0xc6, 0xe8, 0x40, 0xf1, 0xa5, 0x20, 0x49, 0x27, 0xe9, 0xb5,
	0x0a, 0xf0, 0xf8, 0x14, 0x35, 0xfd, 0x27, 0x4b, 0x49, 0x0d, 0x68, 0xc3, 0x9f, 0xd2, 0x80, 0x33,
	0x52, 0x8f, 0x38, 0x0b, 0x38, 0x27, 0x07, 0x11, 0xe7, 0x01, 0x0f, 0x48, 0x23, 0xe2, 0x41, 0xc0,
	0x43, 0xd2, 0x8c, 0x78, 0x18, 0xf0, 0x88, 0x1c, 0x46, 0x3c, 0x0a, 0x78, 0x4c, 0x8e, 0x22, 0x1e,
	0x07, 0x3c, 0x21, 0xad, 0x88, 0x27, 0x01, 0x4f, 0x09, 0x8a, 0x78, 0x8a, 0x9f, 0xa0, 0xc3, 0x2a,
	0x4e, 0x9f, 0xb4, 0x81, 0x43, 0x57, 0xda, 0x8f, 0x85, 0x94, 0x3c, 0xd8, 0x29, 0xa4, 0xdd, 0x9f,
	0x09, 0x3a, 0xdd, 0x0e, 0xcd, 0x5d, 0x1b, 0x61, 0xaf, 0xf5, 0x62, 0xce, 0xdc, 0x66, 0x25, 0xf0,
	0x2b, 0x84, 0x23, 0xb1, 0xa2, 0x14, 0x46, 0xba, 0xcd, 0x76, 0x78, 0x27, 0xa1, 0x72, 0xb1, 0x2d,
	0xfc, 0xdb, 0x6e, 0xc4, 0x82, 0x3b, 0xa9, 0xd5, 0x76, 0xaa, 0xb1, 0xbd, 0xd8, 0x16, 0xf0, 0x4b,
	0xf4, 0x28, 0xb6, 0x97, 0x7c, 0xb1, 0x16, 0x30, 0xea, 0x93, 0xe2, 0x61, 0xc0, 0x1f, 0x3c, 0xc5,
	0x2f, 0x50, 0x24, 0x0c, 0xf6, 0x57, 0xcd, 0xfe, 0x38, 0xd0, 0x77, 0x7c, 0x29, 0xba, 0xdf, 0x6a,
	0xe8, 0xe9, 0x9d, 0x1c, 0xb7, 0xfb, 0xc7, 0x5f, 0x6b, 0xbb, 0x8f, 0xe3, 0xc6, 0x70, 0x9f, 0xa4,
	0xde, 0x6b, 0x67, 0xbf, 0x13, 0x7a, 0x6f, 0xde, 0x00, 0xba, 0x77, 0x91, 0x3b, 0x13, 0x7d, 0xed,
	0x83, 0x77, 0x7f, 0xd4, 0xd1, 0xd9, 0xfe, 0xf7, 0x04, 0x77, 0x50, 0x7b, 0x2e, 0xec, 0xcc, 0xc8,
	0x15, 0x6c, 0x2f, 0x83, 0x49, 0xef, 0x22, 0xfc, 0x18, 0x35, 0xd6, 0x4a, 0x3a, 0x4b, 0xf2, 0xea,
	0xde, 0xc1, 0xc1, 0xd3, 0x6a, 0x87, 0x83, 0x8a, 0xc2, 0x01, 0x9f, 0xa3, 0xd6, 0x9c, 0x3b, 0x0e,
	0xbf, 0x82, 0x0c, 0xa1, 0x72, 0xe4, 0xc1, 0x7b, 0x7f, 0xbd, 0xce, 0x50, 0xd3, 0x3a, 0xee, 0xd6,
	0x96, 0x8c, 0xaa, 0x0b, 0x59, 0x9d, 0xf0, 0x73, 0xd4, 0x5e, 0xaf, 0xe6, 0xdc, 0x09, 0x66, 0xb8,
	0x13, 0x64, 0xdc, 0x49, 0x7a, 0xc7, 0x05, 0xaa, 0x50, 0xc1, 0x9d, 0xc0, 0x5f, 0x6a, 0xa8, 0x15,
	0x12, 0x91, 0x49, 0x27, 0xe9, 0xb5, 0xb3, 0x3f, 0xf7, 0x7a, 0x8b, 0xb7, 0x95, 0x22, 0xe6, 0xbe,
	0x6c, 0xc2, 0xff, 0x77, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x97, 0x5a, 0xf0, 0xd7, 0xdf, 0x05,
	0x00, 0x00,
}
