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
// source: snmp_sensor_info_xml.proto

package cisco_ios_xr_snmp_sensormib_oper_augment_sensor_mib_ent_phy_indexes_ent_phy_index

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

type SnmpSensorInfoXml_KEYS struct {
	Index                string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpSensorInfoXml_KEYS) Reset()         { *m = SnmpSensorInfoXml_KEYS{} }
func (m *SnmpSensorInfoXml_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpSensorInfoXml_KEYS) ProtoMessage()    {}
func (*SnmpSensorInfoXml_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ed9185fed7012cd, []int{0}
}

func (m *SnmpSensorInfoXml_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpSensorInfoXml_KEYS.Unmarshal(m, b)
}
func (m *SnmpSensorInfoXml_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpSensorInfoXml_KEYS.Marshal(b, m, deterministic)
}
func (m *SnmpSensorInfoXml_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpSensorInfoXml_KEYS.Merge(m, src)
}
func (m *SnmpSensorInfoXml_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpSensorInfoXml_KEYS.Size(m)
}
func (m *SnmpSensorInfoXml_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpSensorInfoXml_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpSensorInfoXml_KEYS proto.InternalMessageInfo

func (m *SnmpSensorInfoXml_KEYS) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type SnmpSensorInfoXml struct {
	FieldValidityBitmap  string   `protobuf:"bytes,50,opt,name=field_validity_bitmap,json=fieldValidityBitmap,proto3" json:"field_validity_bitmap,omitempty"`
	DeviceDescription    string   `protobuf:"bytes,51,opt,name=device_description,json=deviceDescription,proto3" json:"device_description,omitempty"`
	Units                string   `protobuf:"bytes,52,opt,name=units,proto3" json:"units,omitempty"`
	DeviceId             uint32   `protobuf:"varint,53,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Value                uint32   `protobuf:"varint,54,opt,name=value,proto3" json:"value,omitempty"`
	AlarmType            uint32   `protobuf:"varint,55,opt,name=alarm_type,json=alarmType,proto3" json:"alarm_type,omitempty"`
	DataType             uint32   `protobuf:"varint,56,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Scale                uint32   `protobuf:"varint,57,opt,name=scale,proto3" json:"scale,omitempty"`
	Precision            uint32   `protobuf:"varint,58,opt,name=precision,proto3" json:"precision,omitempty"`
	Status               uint32   `protobuf:"varint,59,opt,name=status,proto3" json:"status,omitempty"`
	AgeTimeStamp         uint32   `protobuf:"varint,60,opt,name=age_time_stamp,json=ageTimeStamp,proto3" json:"age_time_stamp,omitempty"`
	UpdateRate           uint32   `protobuf:"varint,61,opt,name=update_rate,json=updateRate,proto3" json:"update_rate,omitempty"`
	MeasuredEntity       uint32   `protobuf:"varint,62,opt,name=measured_entity,json=measuredEntity,proto3" json:"measured_entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpSensorInfoXml) Reset()         { *m = SnmpSensorInfoXml{} }
func (m *SnmpSensorInfoXml) String() string { return proto.CompactTextString(m) }
func (*SnmpSensorInfoXml) ProtoMessage()    {}
func (*SnmpSensorInfoXml) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ed9185fed7012cd, []int{1}
}

func (m *SnmpSensorInfoXml) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpSensorInfoXml.Unmarshal(m, b)
}
func (m *SnmpSensorInfoXml) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpSensorInfoXml.Marshal(b, m, deterministic)
}
func (m *SnmpSensorInfoXml) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpSensorInfoXml.Merge(m, src)
}
func (m *SnmpSensorInfoXml) XXX_Size() int {
	return xxx_messageInfo_SnmpSensorInfoXml.Size(m)
}
func (m *SnmpSensorInfoXml) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpSensorInfoXml.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpSensorInfoXml proto.InternalMessageInfo

func (m *SnmpSensorInfoXml) GetFieldValidityBitmap() string {
	if m != nil {
		return m.FieldValidityBitmap
	}
	return ""
}

func (m *SnmpSensorInfoXml) GetDeviceDescription() string {
	if m != nil {
		return m.DeviceDescription
	}
	return ""
}

func (m *SnmpSensorInfoXml) GetUnits() string {
	if m != nil {
		return m.Units
	}
	return ""
}

func (m *SnmpSensorInfoXml) GetDeviceId() uint32 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *SnmpSensorInfoXml) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *SnmpSensorInfoXml) GetAlarmType() uint32 {
	if m != nil {
		return m.AlarmType
	}
	return 0
}

func (m *SnmpSensorInfoXml) GetDataType() uint32 {
	if m != nil {
		return m.DataType
	}
	return 0
}

func (m *SnmpSensorInfoXml) GetScale() uint32 {
	if m != nil {
		return m.Scale
	}
	return 0
}

func (m *SnmpSensorInfoXml) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *SnmpSensorInfoXml) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SnmpSensorInfoXml) GetAgeTimeStamp() uint32 {
	if m != nil {
		return m.AgeTimeStamp
	}
	return 0
}

func (m *SnmpSensorInfoXml) GetUpdateRate() uint32 {
	if m != nil {
		return m.UpdateRate
	}
	return 0
}

func (m *SnmpSensorInfoXml) GetMeasuredEntity() uint32 {
	if m != nil {
		return m.MeasuredEntity
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpSensorInfoXml_KEYS)(nil), "cisco_ios_xr_snmp_sensormib_oper.augment.sensor_mib.ent_phy_indexes.ent_phy_index.snmp_sensor_info_xml_KEYS")
	proto.RegisterType((*SnmpSensorInfoXml)(nil), "cisco_ios_xr_snmp_sensormib_oper.augment.sensor_mib.ent_phy_indexes.ent_phy_index.snmp_sensor_info_xml")
}

func init() { proto.RegisterFile("snmp_sensor_info_xml.proto", fileDescriptor_0ed9185fed7012cd) }

var fileDescriptor_0ed9185fed7012cd = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5b, 0x6b, 0x14, 0x41,
	0x10, 0x85, 0x59, 0x24, 0xc1, 0x6d, 0x35, 0x62, 0xbb, 0x4a, 0x7b, 0xc3, 0x10, 0x04, 0xf3, 0xe2,
	0x82, 0x89, 0xf7, 0xdb, 0x83, 0x98, 0x07, 0xf1, 0xc9, 0x4d, 0x10, 0x7c, 0x2a, 0x6a, 0xa7, 0x2b,
	0x6b, 0xc1, 0x74, 0x4f, 0xd3, 0x5d, 0xb3, 0xec, 0xfc, 0x35, 0x7f, 0x9d, 0x6c, 0xf5, 0x78, 0x83,
	0x3c, 0x9e, 0xf3, 0x9d, 0x73, 0xa6, 0x06, 0xda, 0xdc, 0x2d, 0x31, 0x24, 0x28, 0x14, 0x4b, 0x97,
	0x81, 0xe3, 0x79, 0x07, 0x9b, 0xd0, 0xce, 0x53, 0xee, 0xa4, 0xb3, 0x5f, 0x1b, 0x2e, 0x4d, 0x07,
	0xdc, 0x15, 0xd8, 0x64, 0xf8, 0x27, 0x18, 0x78, 0x09, 0x5d, 0xa2, 0x3c, 0xc7, 0x7e, 0x15, 0x28,
	0xca, 0x7c, 0xec, 0x07, 0x5e, 0xce, 0x29, 0x0a, 0xa4, 0x1f, 0x03, 0x70, 0xf4, 0xb4, 0xa1, 0xf2,
	0xbf, 0x3e, 0x78, 0x6a, 0xee, 0x5c, 0xf4, 0x41, 0xf8, 0x72, 0xf2, 0xfd, 0xd4, 0xce, 0xcc, 0x8e,
	0xa6, 0xdc, 0x64, 0x7f, 0x72, 0x38, 0x5d, 0x54, 0x71, 0xf0, 0xf3, 0x92, 0x99, 0x5d, 0xd4, 0xb1,
	0x47, 0xe6, 0xd6, 0x39, 0x53, 0xeb, 0x61, 0x8d, 0x2d, 0x7b, 0x96, 0x01, 0x96, 0x2c, 0x01, 0x93,
	0x3b, 0xd2, 0xfa, 0x4d, 0x85, 0xdf, 0x46, 0xf6, 0x51, 0x91, 0x7d, 0x62, 0xac, 0xa7, 0x35, 0x37,
	0x04, 0x9e, 0x4a, 0x93, 0x39, 0x09, 0x77, 0xd1, 0x1d, 0x6b, 0xe1, 0x46, 0x25, 0x9f, 0xfe, 0x82,
	0xed, 0x45, 0x7d, 0x64, 0x29, 0xee, 0x59, 0xbd, 0x48, 0x85, 0xbd, 0x67, 0xa6, 0xe3, 0x08, 0x7b,
	0xf7, 0x7c, 0x7f, 0x72, 0x78, 0x6d, 0x71, 0xb9, 0x1a, 0x9f, 0xfd, 0xb6, 0xb2, 0xc6, 0xb6, 0x27,
	0xf7, 0x42, 0x41, 0x15, 0xf6, 0x81, 0x31, 0xd8, 0x62, 0x0e, 0x20, 0x43, 0x22, 0xf7, 0x52, 0xd1,
	0x54, 0x9d, 0xb3, 0x21, 0x91, 0x2e, 0xa2, 0x60, 0xa5, 0xaf, 0xc6, 0x45, 0x14, 0x54, 0x38, 0x33,
	0x3b, 0xa5, 0xc1, 0x96, 0xdc, 0xeb, 0xba, 0xa8, 0xc2, 0xde, 0x37, 0xd3, 0x94, 0xa9, 0xe1, 0xb2,
	0xfd, 0x81, 0x37, 0x75, 0xf0, 0x8f, 0x61, 0x6f, 0x9b, 0xdd, 0x22, 0x28, 0x7d, 0x71, 0x6f, 0x15,
	0x8d, 0xca, 0x3e, 0x32, 0x7b, 0xb8, 0x22, 0x10, 0x0e, 0x04, 0x45, 0x30, 0x24, 0xf7, 0x4e, 0xf9,
	0x55, 0x5c, 0xd1, 0x19, 0x07, 0x3a, 0xdd, 0x7a, 0xf6, 0xa1, 0xb9, 0xd2, 0x27, 0x8f, 0x42, 0x90,
	0x51, 0xc8, 0xbd, 0xd7, 0x88, 0xa9, 0xd6, 0x02, 0x85, 0xec, 0x63, 0x73, 0x3d, 0x10, 0x96, 0x3e,
	0x93, 0x07, 0x8a, 0xc2, 0x32, 0xb8, 0x0f, 0x1a, 0xda, 0xfb, 0x6d, 0x9f, 0xa8, 0xbb, 0xdc, 0xd5,
	0x97, 0x74, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xf3, 0xc6, 0x92, 0x67, 0x02, 0x00, 0x00,
}