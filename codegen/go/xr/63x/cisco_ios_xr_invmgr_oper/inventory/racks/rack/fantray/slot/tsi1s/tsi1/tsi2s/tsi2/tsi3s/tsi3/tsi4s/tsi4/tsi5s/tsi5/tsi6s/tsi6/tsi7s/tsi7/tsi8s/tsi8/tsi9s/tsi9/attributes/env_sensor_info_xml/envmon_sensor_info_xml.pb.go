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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_attributes_env_sensor_info_xml

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
	proto.RegisterType((*EnvmonSensorInfoXml_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.env_sensor_info_xml.envmon_sensor_info_xml_KEYS")
	proto.RegisterType((*EnvmonThresholdType)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.env_sensor_info_xml.envmon_threshold_type")
	proto.RegisterType((*EnvmonThresholdInfoXml)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.env_sensor_info_xml.envmon_threshold_info_xml")
	proto.RegisterType((*EnvmonSensorInfoXml)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.env_sensor_info_xml.envmon_sensor_info_xml")
}

func init() { proto.RegisterFile("envmon_sensor_info_xml.proto", fileDescriptor_48a5d855b51dbaa9) }

var fileDescriptor_48a5d855b51dbaa9 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0x4f, 0x6e, 0x13, 0x31,
	0x14, 0xc6, 0x35, 0x49, 0x93, 0x36, 0x8e, 0x0a, 0xaa, 0x45, 0x8b, 0x51, 0x91, 0x88, 0x22, 0x21,
	0xb2, 0xc1, 0x6a, 0x66, 0xf2, 0x77, 0xc9, 0x82, 0x15, 0x12, 0x8b, 0x29, 0x42, 0x62, 0x65, 0xb9,
	0x89, 0x4b, 0x2d, 0x12, 0x3b, 0xb2, 0x3d, 0xa3, 0xe6, 0x28, 0x88, 0x73, 0xb0, 0x63, 0xcb, 0x09,
	0x58, 0x70, 0x13, 0x76, 0x48, 0xc8, 0x6f, 0xa6, 0x76, 0x80, 0x5c, 0x00, 0x36, 0x5f, 0x3e, 0xff,
	0x9e, 0xa5, 0xcc, 0xf7, 0xde, 0x9b, 0x41, 0x8f, 0x85, 0x2a, 0xd7, 0x5a, 0x31, 0x2b, 0x94, 0xd5,
	0x86, 0x49, 0x75, 0xad, 0xd9, 0xed, 0x7a, 0x45, 0x37, 0x46, 0x3b, 0x8d, 0xbf, 0x24, 0x0b, 0x69,
	0x17, 0x9a, 0x49, 0x6d, 0xd9, 0xad, 0xaf, 0x96, 0xeb, 0xf7, 0x86, 0xe9, 0x8d, 0x30, 0x54, 0xaa,
	0x52, 0x28, 0xa7, 0xcd, 0x96, 0x1a, 0xbe, 0xf8, 0x60, 0x41, 0xe9, 0x35, 0x57, 0xce, 0xf0, 0x2d,
	0xb5, 0x2b, 0xed, 0xa8, 0xb3, 0x72, 0x68, 0x41, 0xbd, 0xa4, 0x60, 0x53, 0x2f, 0x19, 0xd8, 0xcc,
	0xcb, 0x08, 0xec, 0xc8, 0xcb, 0x18, 0xec, 0xd8, 0xcb, 0x04, 0xec, 0xc4, 0xcb, 0x14, 0xec, 0xd4,
	0xcb, 0x0c, 0xec, 0xcc, 0xcb, 0x1c, 0xec, 0x9c, 0x72, 0xe7, 0x8c, 0xbc, 0x2a, 0x9c, 0xb0, 0x54,
	0xa8, 0xf2, 0xcf, 0x10, 0xfd, 0x8f, 0x0d, 0x74, 0xbe, 0x3f, 0x1f, 0x7b, 0xf5, 0xf2, 0xdd, 0x25,
	0xc6, 0xe8, 0x40, 0xf1, 0xb5, 0x20, 0x49, 0x2f, 0x19, 0x74, 0x72, 0xf0, 0xf8, 0x14, 0xb5, 0xfd,
	0x2f, 0x1b, 0x92, 0x06, 0xd0, 0x96, 0x3f, 0x0d, 0x03, 0x4e, 0x49, 0x33, 0xe2, 0x34, 0xe0, 0x8c,
	0x1c, 0x44, 0x9c, 0x05, 0x3c, 0x22, 0xad, 0x88, 0x47, 0x01, 0x8f, 0x49, 0x3b, 0xe2, 0x71, 0xc0,
	0x13, 0x72, 0x18, 0xf1, 0x24, 0xe0, 0x29, 0x39, 0x8a, 0x78, 0x1a, 0xf0, 0x8c, 0x74, 0x22, 0x9e,
	0x05, 0x3c, 0x27, 0x28, 0xe2, 0x39, 0x7e, 0x88, 0x0e, 0xab, 0x38, 0x17, 0xa4, 0x0b, 0x1c, 0x6e,
	0x0d, 0x2f, 0xfa, 0x5f, 0x13, 0x74, 0x5a, 0xf7, 0xc6, 0xdd, 0x18, 0x61, 0x6f, 0xf4, 0x6a, 0xc9,
	0xdc, 0x76, 0x23, 0xf0, 0x73, 0x84, 0x23, 0xb1, 0xa2, 0x14, 0x46, 0xba, 0x6d, 0xdd, 0xa3, 0x93,
	0x50, 0xb9, 0xac, 0x0b, 0xbf, 0x5f, 0x37, 0x62, 0xc5, 0x9d, 0xd4, 0xaa, 0x6e, 0x5e, 0xbc, 0x9e,
	0xd7, 0x05, 0xfc, 0x0c, 0xdd, 0x8f, 0xd7, 0x4b, 0xbe, 0x2a, 0x04, 0x74, 0xf4, 0x24, 0xbf, 0x17,
	0xf0, 0x5b, 0x4f, 0xf1, 0x53, 0x14, 0x09, 0x83, 0x31, 0x55, 0x2d, 0x3e, 0x0e, 0xf4, 0x35, 0x5f,
	0x8b, 0xfe, 0xa7, 0x06, 0x7a, 0xf4, 0x57, 0x8e, 0xbb, 0x31, 0xe3, 0x9f, 0xc9, 0xee, 0xdf, 0x71,
	0x63, 0xb8, 0x4f, 0xd2, 0x1c, 0x74, 0xd3, 0x6f, 0x09, 0xfd, 0x97, 0x77, 0x9b, 0xee, 0x9d, 0xdd,
	0x4e, 0x13, 0x5f, 0xf8, 0xac, 0xfd, 0xcf, 0x4d, 0x74, 0xb6, 0xff, 0x0d, 0xc0, 0x3d, 0xd4, 0x5d,
	0x0a, 0xbb, 0x30, 0x72, 0x03, 0x03, 0x4b, 0xa1, 0xb9, 0xbb, 0x08, 0x3f, 0x40, 0xad, 0x42, 0x49,
	0x67, 0x49, 0x56, 0x6d, 0x14, 0x1c, 0x3c, 0xad, 0xc6, 0x36, 0xaa, 0x28, 0x1c, 0xf0, 0x39, 0xea,
	0x2c, 0xb9, 0xe3, 0xf0, 0x14, 0x64, 0x0c, 0x95, 0x23, 0x0f, 0xde, 0xf8, 0x8d, 0x3a, 0x43, 0x6d,
	0xeb, 0xb8, 0x2b, 0x2c, 0x99, 0x54, 0x3b, 0x58, 0x9d, 0xf0, 0x13, 0xd4, 0x2d, 0x36, 0x4b, 0xee,
	0x04, 0x33, 0xdc, 0x09, 0x32, 0xed, 0x25, 0x83, 0xe3, 0x1c, 0x55, 0x28, 0xe7, 0x4e, 0xe0, 0x1f,
	0x09, 0xea, 0x84, 0x44, 0x64, 0xd6, 0x4b, 0x06, 0xdd, 0xf4, 0xfb, 0xff, 0x36, 0xb8, 0xbb, 0x4a,
	0x1e, 0xa3, 0x5e, 0xb5, 0xe1, 0xfb, 0x9b, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x95, 0xec, 0xa8,
	0x7b, 0x9f, 0x05, 0x00, 0x00,
}
