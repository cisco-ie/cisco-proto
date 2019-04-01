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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_attributes_env_sensor_info_xml

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
	proto.RegisterType((*EnvmonSensorInfoXml_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.env_sensor_info_xml.envmon_sensor_info_xml_KEYS")
	proto.RegisterType((*EnvmonThresholdType)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.env_sensor_info_xml.envmon_threshold_type")
	proto.RegisterType((*EnvmonThresholdInfoXml)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.env_sensor_info_xml.envmon_threshold_info_xml")
	proto.RegisterType((*EnvmonSensorInfoXml)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.env_sensor_info_xml.envmon_sensor_info_xml")
}

func init() { proto.RegisterFile("envmon_sensor_info_xml.proto", fileDescriptor_48a5d855b51dbaa9) }

var fileDescriptor_48a5d855b51dbaa9 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xbf, 0x8e, 0x13, 0x3f,
	0x10, 0xc7, 0xe5, 0x5f, 0x2e, 0xf9, 0x11, 0x47, 0x07, 0x3a, 0x8b, 0x3b, 0x19, 0x1d, 0x12, 0x51,
	0x24, 0x44, 0x1a, 0x2c, 0x65, 0xff, 0x04, 0x5a, 0x0a, 0x2a, 0x24, 0x8a, 0x3d, 0x84, 0x44, 0x65,
	0xf9, 0x12, 0xc3, 0x59, 0xec, 0xda, 0x2b, 0x7b, 0x76, 0x75, 0x79, 0x03, 0x3a, 0xde, 0x82, 0x82,
	0x57, 0xa0, 0xa6, 0xa0, 0xe3, 0x49, 0x78, 0x06, 0x64, 0xef, 0xde, 0x3a, 0x40, 0x1e, 0x00, 0x9a,
	0xc9, 0xcc, 0xe7, 0x6b, 0x29, 0xf3, 0x9d, 0x19, 0x2d, 0xbe, 0x2f, 0x75, 0x5b, 0x19, 0xcd, 0x9d,
	0xd4, 0xce, 0x58, 0xae, 0xf4, 0x5b, 0xc3, 0xaf, 0xab, 0x92, 0xd5, 0xd6, 0x80, 0x21, 0x1f, 0xd1,
	0x46, 0xb9, 0x8d, 0xe1, 0xca, 0x38, 0x7e, 0xed, 0xd5, 0xb6, 0x7a, 0x67, 0xb9, 0xa9, 0xa5, 0x65,
	0x4a, 0xb7, 0x52, 0x83, 0xb1, 0x3b, 0x66, 0xc5, 0xe6, 0xbd, 0x0b, 0x91, 0x49, 0x0d, 0x0a, 0x76,
	0xcc, 0x95, 0x06, 0x18, 0x38, 0xb5, 0x72, 0x21, 0xfa, 0x90, 0x84, 0x34, 0xf1, 0x21, 0x0d, 0x69,
	0xea, 0x43, 0x16, 0xd2, 0xcc, 0x87, 0x3c, 0xa4, 0x39, 0x13, 0x00, 0x56, 0x5d, 0x36, 0x20, 0x1d,
	0x93, 0xba, 0xfd, 0xbd, 0xaf, 0xc5, 0x17, 0x84, 0xcf, 0x0f, 0xb7, 0xcc, 0x5f, 0x3c, 0x7f, 0x73,
	0x41, 0x08, 0x3e, 0xd2, 0xa2, 0x92, 0x14, 0xcd, 0xd1, 0x72, 0x5a, 0x84, 0x9c, 0x9c, 0xe2, 0x89,
	0xff, 0xe5, 0x2b, 0xfa, 0x5f, 0xa0, 0x63, 0x5f, 0xad, 0x06, 0x9c, 0xd0, 0x51, 0xc4, 0xc9, 0x80,
	0x53, 0x7a, 0x14, 0x71, 0x3a, 0xe0, 0x8c, 0x8e, 0x23, 0xce, 0x06, 0x9c, 0xd3, 0x49, 0xc4, 0xf9,
	0x80, 0xd7, 0xf4, 0xff, 0x88, 0xd7, 0x8b, 0xaf, 0x08, 0x9f, 0xf6, 0xdd, 0xc3, 0x95, 0x95, 0xee,
	0xca, 0x94, 0x5b, 0x0e, 0xbb, 0x5a, 0x92, 0xc7, 0x98, 0x44, 0xe2, 0x64, 0x2b, 0xad, 0x82, 0x5d,
	0xef, 0xe2, 0x64, 0x50, 0x2e, 0x7a, 0xe1, 0xd7, 0xe7, 0x56, 0x96, 0x02, 0x94, 0xd1, 0xbd, 0xbd,
	0xf8, 0xbc, 0xe8, 0x05, 0xf2, 0x08, 0xdf, 0x89, 0xcf, 0x5b, 0x51, 0x36, 0x32, 0x78, 0x3e, 0x29,
	0x6e, 0x0f, 0xf8, 0xb5, 0xa7, 0xe4, 0x21, 0x8e, 0x84, 0x87, 0x41, 0x76, 0x43, 0x38, 0x1e, 0xe8,
	0x4b, 0x51, 0xc9, 0xc5, 0x0f, 0x84, 0xef, 0xfd, 0xe1, 0xe3, 0x66, 0x11, 0xe4, 0x3b, 0xda, 0xff,
	0x3b, 0x61, 0xad, 0xf0, 0x4e, 0x46, 0xcb, 0x59, 0xf2, 0x09, 0xb1, 0xbf, 0xec, 0xa0, 0xd8, 0xc1,
	0x75, 0xec, 0xcd, 0xe5, 0x99, 0x6f, 0x7f, 0xf1, 0x61, 0x84, 0xcf, 0x0e, 0x9f, 0x1d, 0x99, 0xe3,
	0xd9, 0x56, 0xba, 0x8d, 0x55, 0x75, 0xd8, 0x41, 0x12, 0xe6, 0xb5, 0x8f, 0xc8, 0x5d, 0x3c, 0x6e,
	0xb4, 0x02, 0x47, 0xd3, 0xee, 0x16, 0x42, 0xe1, 0x69, 0xb7, 0x89, 0xac, 0xa3, 0xa1, 0x20, 0xe7,
	0x78, 0xba, 0x15, 0x20, 0x42, 0x17, 0x34, 0x0f, 0xca, 0x2d, 0x0f, 0x5e, 0xf9, 0x23, 0x39, 0xc3,
	0x13, 0x07, 0x02, 0x1a, 0x47, 0xd7, 0x41, 0xe9, 0x2b, 0xf2, 0x00, 0xcf, 0x9a, 0x7a, 0x2b, 0x40,
	0x72, 0x2b, 0x40, 0xd2, 0x27, 0x73, 0xb4, 0x3c, 0x2e, 0x70, 0x87, 0x0a, 0x01, 0x92, 0x7c, 0x43,
	0x78, 0x3a, 0x38, 0xa2, 0x4f, 0xe7, 0x68, 0x39, 0x4b, 0x3e, 0xff, 0x03, 0xbb, 0xb8, 0x51, 0x8a,
	0xd8, 0xfd, 0xe5, 0x24, 0x7c, 0x9a, 0xd2, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81, 0x4c, 0x8e,
	0x61, 0xba, 0x04, 0x00, 0x00,
}
