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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_attributes_env_sensor_info_xml

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
	proto.RegisterType((*EnvmonSensorInfoXml_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.env_sensor_info_xml.envmon_sensor_info_xml_KEYS")
	proto.RegisterType((*EnvmonThresholdType)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.env_sensor_info_xml.envmon_threshold_type")
	proto.RegisterType((*EnvmonThresholdInfoXml)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.env_sensor_info_xml.envmon_threshold_info_xml")
	proto.RegisterType((*EnvmonSensorInfoXml)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.env_sensor_info_xml.envmon_sensor_info_xml")
}

func init() { proto.RegisterFile("envmon_sensor_info_xml.proto", fileDescriptor_48a5d855b51dbaa9) }

var fileDescriptor_48a5d855b51dbaa9 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x49, 0x93, 0x36, 0x8e, 0x0a, 0xaa, 0x45, 0x8b, 0x51, 0x91, 0x88, 0x22, 0x21,
	0x72, 0x61, 0xd5, 0xec, 0xe6, 0xef, 0x91, 0x03, 0x27, 0x24, 0x0e, 0x5b, 0x84, 0xc4, 0xc9, 0x72,
	0x13, 0x43, 0x2d, 0x12, 0x3b, 0xb2, 0x27, 0xab, 0xe6, 0x51, 0xe0, 0x39, 0xb8, 0x71, 0xe6, 0x05,
	0x38, 0xf0, 0x26, 0xdc, 0x38, 0x20, 0xcf, 0xa6, 0x76, 0x80, 0xbc, 0x00, 0xbd, 0x7c, 0xf9, 0xfc,
	0x1b, 0x4b, 0xd9, 0x6f, 0x66, 0x76, 0xc9, 0x63, 0xa9, 0xcb, 0xa5, 0xd1, 0xdc, 0x49, 0xed, 0x8c,
	0xe5, 0x4a, 0xbf, 0x37, 0xfc, 0x66, 0xb9, 0x48, 0x57, 0xd6, 0x80, 0xa1, 0x5f, 0x93, 0x99, 0x72,
	0x33, 0xc3, 0x95, 0x71, 0xfc, 0xc6, 0x57, 0xcb, 0xe5, 0x07, 0xcb, 0xcd, 0x4a, 0xda, 0x54, 0xe9,
	0x52, 0x6a, 0x30, 0x76, 0x93, 0x5a, 0x31, 0xfb, 0xe8, 0x50, 0x53, 0xa9, 0x41, 0xc1, 0x26, 0x75,
	0x0b, 0x03, 0x29, 0x38, 0xd5, 0x77, 0xa8, 0x5e, 0x32, 0xb4, 0x99, 0x97, 0x1c, 0x6d, 0xee, 0x65,
	0x80, 0x76, 0xe0, 0x65, 0x88, 0x76, 0xe8, 0x65, 0x84, 0x76, 0xe4, 0x65, 0x8c, 0x76, 0xec, 0x65,
	0x82, 0x76, 0xe2, 0x65, 0x8a, 0x76, 0x9a, 0x0a, 0x00, 0xab, 0xae, 0xd6, 0x20, 0x5d, 0x2a, 0x75,
	0xf9, 0x77, 0x86, 0xee, 0xa7, 0x1a, 0x39, 0xdf, 0x1f, 0x8f, 0xbf, 0x7a, 0xf9, 0xee, 0x92, 0x52,
	0x72, 0xa0, 0xc5, 0x52, 0xb2, 0xa4, 0x93, 0xf4, 0x5a, 0x05, 0x7a, 0x7a, 0x4a, 0x9a, 0xfe, 0x97,
	0xf7, 0x59, 0x0d, 0x69, 0xc3, 0x9f, 0xfa, 0x01, 0x67, 0xac, 0x1e, 0x71, 0x16, 0x70, 0xce, 0x0e,
	0x22, 0xce, 0x03, 0x1e, 0xb0, 0x46, 0xc4, 0x83, 0x80, 0x87, 0xac, 0x19, 0xf1, 0x30, 0xe0, 0x11,
	0x3b, 0x8c, 0x78, 0x14, 0xf0, 0x98, 0x1d, 0x45, 0x3c, 0x0e, 0x78, 0xc2, 0x5a, 0x11, 0x4f, 0x02,
	0x9e, 0x32, 0x12, 0xf1, 0x94, 0x3e, 0x24, 0x87, 0x55, 0x9c, 0x0b, 0xd6, 0x46, 0x8e, 0xb7, 0xfa,
	0x17, 0xdd, 0x6f, 0x09, 0x39, 0xdd, 0xf6, 0x06, 0xae, 0xad, 0x74, 0xd7, 0x66, 0x31, 0xe7, 0xb0,
	0x59, 0x49, 0xfa, 0x9c, 0xd0, 0x48, 0x9c, 0x2c, 0xa5, 0x55, 0xb0, 0xd9, 0xf6, 0xe8, 0x24, 0x54,
	0x2e, 0xb7, 0x85, 0x3f, 0xaf, 0x5b, 0xb9, 0x10, 0xa0, 0x8c, 0xde, 0x36, 0x2f, 0x5e, 0x2f, 0xb6,
	0x05, 0xfa, 0x8c, 0xdc, 0x8f, 0xd7, 0x4b, 0xb1, 0x58, 0x4b, 0xec, 0xe8, 0x49, 0x71, 0x2f, 0xe0,
	0xb7, 0x9e, 0xd2, 0xa7, 0x24, 0x12, 0x8e, 0x63, 0xaa, 0x5a, 0x7c, 0x1c, 0xe8, 0x6b, 0xb1, 0x94,
	0xdd, 0xcf, 0x35, 0xf2, 0xe8, 0x9f, 0x1c, 0xb7, 0x63, 0xa6, 0xbf, 0x92, 0xdd, 0xbf, 0x13, 0xd6,
	0x0a, 0x9f, 0xa4, 0xde, 0x6b, 0x67, 0xdf, 0x93, 0xf4, 0x3f, 0x5e, 0xed, 0x74, 0xef, 0xe8, 0x76,
	0x7a, 0xf8, 0xc2, 0x47, 0xed, 0x7e, 0xa9, 0x93, 0xb3, 0xfd, 0x2f, 0x00, 0xed, 0x90, 0xf6, 0x5c,
	0xba, 0x99, 0x55, 0x2b, 0x9c, 0x57, 0x86, 0xbd, 0xdd, 0x45, 0xf4, 0x01, 0x69, 0xac, 0xb5, 0x02,
	0xc7, 0xf2, 0x6a, 0xa1, 0xf0, 0xe0, 0x69, 0x35, 0xb5, 0x41, 0x45, 0xf1, 0x40, 0xcf, 0x49, 0x6b,
	0x2e, 0x40, 0xe0, 0x53, 0xb0, 0x21, 0x56, 0x8e, 0x3c, 0x78, 0xe3, 0x17, 0xea, 0x8c, 0x34, 0x1d,
	0x08, 0x58, 0x3b, 0x36, 0xaa, 0x56, 0xb0, 0x3a, 0xd1, 0x27, 0xa4, 0xbd, 0x5e, 0xcd, 0x05, 0x48,
	0x6e, 0x05, 0x48, 0x36, 0xee, 0x24, 0xbd, 0xe3, 0x82, 0x54, 0xa8, 0x10, 0x20, 0xe9, 0xcf, 0x84,
	0xb4, 0x42, 0x22, 0x36, 0xe9, 0x24, 0xbd, 0x76, 0xf6, 0xe3, 0x8e, 0xcd, 0xed, 0xb6, 0x52, 0xc4,
	0xa4, 0x57, 0x4d, 0xfc, 0xf8, 0xe6, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xa1, 0xff, 0xa9,
	0x9c, 0x05, 0x00, 0x00,
}
