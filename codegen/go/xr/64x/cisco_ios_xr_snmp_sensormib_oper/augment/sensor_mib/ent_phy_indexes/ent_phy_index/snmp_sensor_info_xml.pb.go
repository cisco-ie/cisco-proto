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
	FieldValidityBitmap  []byte   `protobuf:"bytes,50,opt,name=field_validity_bitmap,json=fieldValidityBitmap,proto3" json:"field_validity_bitmap,omitempty"`
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

func (m *SnmpSensorInfoXml) GetFieldValidityBitmap() []byte {
	if m != nil {
		return m.FieldValidityBitmap
	}
	return nil
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
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5b, 0x8b, 0x13, 0x41,
	0x10, 0x85, 0x09, 0xb2, 0x8b, 0x69, 0xd7, 0x15, 0xdb, 0x28, 0xed, 0x0d, 0xc3, 0x22, 0x98, 0x17,
	0x03, 0xee, 0x7a, 0xbf, 0x3d, 0x88, 0xfb, 0x20, 0x3e, 0x99, 0x5d, 0x04, 0x9f, 0x8a, 0xca, 0x4c,
	0x6d, 0x2c, 0x98, 0xee, 0x69, 0xba, 0x6b, 0x42, 0xe6, 0xaf, 0xf9, 0xeb, 0x24, 0xd5, 0xe3, 0x65,
	0x21, 0x8f, 0xe7, 0x7c, 0xa7, 0x4e, 0x55, 0x43, 0x9b, 0x7b, 0x39, 0xf8, 0x08, 0x99, 0x42, 0x6e,
	0x13, 0x70, 0xb8, 0x68, 0x61, 0xe3, 0x9b, 0x79, 0x4c, 0xad, 0xb4, 0xf6, 0x5b, 0xc5, 0xb9, 0x6a,
	0x81, 0xdb, 0x0c, 0x9b, 0x04, 0xff, 0x05, 0x3d, 0x2f, 0xa1, 0x8d, 0x94, 0xe6, 0xd8, 0xad, 0x3c,
	0x05, 0x99, 0x0f, 0xf3, 0x9e, 0x97, 0x73, 0x0a, 0x02, 0xf1, 0x67, 0x0f, 0x1c, 0x6a, 0xda, 0x50,
	0xbe, 0xac, 0x8f, 0x9e, 0x99, 0xbb, 0xbb, 0x16, 0xc2, 0xd7, 0xd3, 0x1f, 0x67, 0x76, 0x62, 0xf6,
	0x34, 0xe5, 0x46, 0xd3, 0xd1, 0x6c, 0xbc, 0x28, 0xe2, 0xe8, 0xd7, 0x15, 0x33, 0xd9, 0x35, 0x63,
	0x8f, 0xcd, 0xed, 0x0b, 0xa6, 0xa6, 0x86, 0x35, 0x36, 0x5c, 0xb3, 0xf4, 0xb0, 0x64, 0xf1, 0x18,
	0xdd, 0xf1, 0x74, 0x34, 0x3b, 0x58, 0xdc, 0x52, 0xf8, 0x7d, 0x60, 0x9f, 0x14, 0xd9, 0xa7, 0xc6,
	0xd6, 0xb4, 0xe6, 0x8a, 0xa0, 0xa6, 0x5c, 0x25, 0x8e, 0xc2, 0x6d, 0x70, 0x27, 0xba, 0xef, 0x66,
	0x21, 0x9f, 0xff, 0x81, 0xed, 0x45, 0x5d, 0x60, 0xc9, 0xee, 0x79, 0xb9, 0x48, 0x85, 0xbd, 0x6f,
	0xc6, 0x43, 0x09, 0xd7, 0xee, 0xc5, 0x74, 0x34, 0xbb, 0xbe, 0xb8, 0x5a, 0x8c, 0x2f, 0xf5, 0x76,
	0x64, 0x8d, 0x4d, 0x47, 0xee, 0xa5, 0x82, 0x22, 0xec, 0x43, 0x63, 0xb0, 0xc1, 0xe4, 0x41, 0xfa,
	0x48, 0xee, 0x95, 0xa2, 0xb1, 0x3a, 0xe7, 0x7d, 0x24, 0x6d, 0x44, 0xc1, 0x42, 0x5f, 0x0f, 0x8d,
	0x28, 0xa8, 0x70, 0x62, 0xf6, 0x72, 0x85, 0x0d, 0xb9, 0x37, 0xa5, 0x51, 0x85, 0x7d, 0x60, 0xc6,
	0x31, 0x51, 0xc5, 0x79, 0xfb, 0x80, 0xb7, 0xa5, 0xf0, 0xaf, 0x61, 0xef, 0x98, 0xfd, 0x2c, 0x28,
	0x5d, 0x76, 0xef, 0x14, 0x0d, 0xca, 0x3e, 0x36, 0x87, 0xb8, 0x22, 0x10, 0xf6, 0x04, 0x59, 0xd0,
	0x47, 0xf7, 0x5e, 0xf9, 0x01, 0xae, 0xe8, 0x9c, 0x3d, 0x9d, 0x6d, 0x3d, 0xfb, 0xc8, 0x5c, 0xeb,
	0x62, 0x8d, 0x42, 0x90, 0x50, 0xc8, 0x7d, 0xd0, 0x88, 0x29, 0xd6, 0x02, 0x85, 0xec, 0x13, 0x73,
	0xc3, 0x13, 0xe6, 0x2e, 0x51, 0x0d, 0x14, 0x84, 0xa5, 0x77, 0x1f, 0x35, 0x74, 0xf8, 0xc7, 0x3e,
	0x55, 0x77, 0xb9, 0xaf, 0x3f, 0xe9, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xb1, 0xbe,
	0x59, 0x67, 0x02, 0x00, 0x00,
}
