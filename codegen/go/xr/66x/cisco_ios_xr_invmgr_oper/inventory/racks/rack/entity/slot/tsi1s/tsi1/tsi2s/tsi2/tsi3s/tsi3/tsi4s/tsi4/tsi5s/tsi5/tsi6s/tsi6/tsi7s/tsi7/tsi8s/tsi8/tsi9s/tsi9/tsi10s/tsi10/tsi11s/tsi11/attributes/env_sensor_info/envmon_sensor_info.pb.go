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
// source: envmon_sensor_info.proto

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_tsi10s_tsi10_tsi11s_tsi11_attributes_env_sensor_info

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

type EnvmonSensorInfo_KEYS struct {
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

func (m *EnvmonSensorInfo_KEYS) Reset()         { *m = EnvmonSensorInfo_KEYS{} }
func (m *EnvmonSensorInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*EnvmonSensorInfo_KEYS) ProtoMessage()    {}
func (*EnvmonSensorInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc03e94ffc42a321, []int{0}
}

func (m *EnvmonSensorInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvmonSensorInfo_KEYS.Unmarshal(m, b)
}
func (m *EnvmonSensorInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvmonSensorInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *EnvmonSensorInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvmonSensorInfo_KEYS.Merge(m, src)
}
func (m *EnvmonSensorInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_EnvmonSensorInfo_KEYS.Size(m)
}
func (m *EnvmonSensorInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvmonSensorInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_EnvmonSensorInfo_KEYS proto.InternalMessageInfo

func (m *EnvmonSensorInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_4() string {
	if m != nil {
		return m.Name_4
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_5() string {
	if m != nil {
		return m.Name_5
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_6() string {
	if m != nil {
		return m.Name_6
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_7() string {
	if m != nil {
		return m.Name_7
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_8() string {
	if m != nil {
		return m.Name_8
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_9() string {
	if m != nil {
		return m.Name_9
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_10() string {
	if m != nil {
		return m.Name_10
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_11() string {
	if m != nil {
		return m.Name_11
	}
	return ""
}

func (m *EnvmonSensorInfo_KEYS) GetName_12() string {
	if m != nil {
		return m.Name_12
	}
	return ""
}

type EnvmonSensorInfo struct {
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
	Average              int32    `protobuf:"zigzag32,62,opt,name=average,proto3" json:"average,omitempty"`
	Minimum              int32    `protobuf:"zigzag32,63,opt,name=minimum,proto3" json:"minimum,omitempty"`
	Maximum              int32    `protobuf:"zigzag32,64,opt,name=maximum,proto3" json:"maximum,omitempty"`
	Interval             int32    `protobuf:"zigzag32,65,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvmonSensorInfo) Reset()         { *m = EnvmonSensorInfo{} }
func (m *EnvmonSensorInfo) String() string { return proto.CompactTextString(m) }
func (*EnvmonSensorInfo) ProtoMessage()    {}
func (*EnvmonSensorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc03e94ffc42a321, []int{1}
}

func (m *EnvmonSensorInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvmonSensorInfo.Unmarshal(m, b)
}
func (m *EnvmonSensorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvmonSensorInfo.Marshal(b, m, deterministic)
}
func (m *EnvmonSensorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvmonSensorInfo.Merge(m, src)
}
func (m *EnvmonSensorInfo) XXX_Size() int {
	return xxx_messageInfo_EnvmonSensorInfo.Size(m)
}
func (m *EnvmonSensorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvmonSensorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EnvmonSensorInfo proto.InternalMessageInfo

func (m *EnvmonSensorInfo) GetFieldValidityBitmap() string {
	if m != nil {
		return m.FieldValidityBitmap
	}
	return ""
}

func (m *EnvmonSensorInfo) GetDeviceDescription() string {
	if m != nil {
		return m.DeviceDescription
	}
	return ""
}

func (m *EnvmonSensorInfo) GetUnits() string {
	if m != nil {
		return m.Units
	}
	return ""
}

func (m *EnvmonSensorInfo) GetDeviceId() uint32 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *EnvmonSensorInfo) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *EnvmonSensorInfo) GetAlarmType() uint32 {
	if m != nil {
		return m.AlarmType
	}
	return 0
}

func (m *EnvmonSensorInfo) GetDataType() uint32 {
	if m != nil {
		return m.DataType
	}
	return 0
}

func (m *EnvmonSensorInfo) GetScale() uint32 {
	if m != nil {
		return m.Scale
	}
	return 0
}

func (m *EnvmonSensorInfo) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *EnvmonSensorInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *EnvmonSensorInfo) GetAgeTimeStamp() uint32 {
	if m != nil {
		return m.AgeTimeStamp
	}
	return 0
}

func (m *EnvmonSensorInfo) GetUpdateRate() uint32 {
	if m != nil {
		return m.UpdateRate
	}
	return 0
}

func (m *EnvmonSensorInfo) GetAverage() int32 {
	if m != nil {
		return m.Average
	}
	return 0
}

func (m *EnvmonSensorInfo) GetMinimum() int32 {
	if m != nil {
		return m.Minimum
	}
	return 0
}

func (m *EnvmonSensorInfo) GetMaximum() int32 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func (m *EnvmonSensorInfo) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func init() {
	proto.RegisterType((*EnvmonSensorInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.env_sensor_info.envmon_sensor_info_KEYS")
	proto.RegisterType((*EnvmonSensorInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.env_sensor_info.envmon_sensor_info")
}

func init() { proto.RegisterFile("envmon_sensor_info.proto", fileDescriptor_bc03e94ffc42a321) }

var fileDescriptor_bc03e94ffc42a321 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x94, 0x4b, 0x6b, 0xdb, 0x4c,
	0x14, 0x86, 0x71, 0x12, 0x3b, 0xf1, 0x24, 0xf9, 0x20, 0xf3, 0x25, 0xcd, 0xa1, 0x17, 0x1a, 0x42,
	0x17, 0xd9, 0x54, 0x58, 0xf2, 0xbd, 0xf7, 0x96, 0x76, 0x51, 0xba, 0x73, 0x42, 0xa1, 0xab, 0x61,
	0x22, 0x9d, 0x98, 0xa1, 0xd2, 0x48, 0xcc, 0x8c, 0x44, 0xfc, 0x0b, 0xbb, 0xee, 0xcf, 0xe9, 0xae,
	0xcc, 0x91, 0x2a, 0x19, 0xb2, 0x79, 0x7d, 0xde, 0xe7, 0xd1, 0x1c, 0x31, 0x5e, 0x88, 0x01, 0xea,
	0x2a, 0xcb, 0xb5, 0xb0, 0xa8, 0x6d, 0x6e, 0x84, 0xd2, 0x77, 0x79, 0x50, 0x98, 0xdc, 0xe5, 0xfc,
	0x77, 0x2f, 0x56, 0x36, 0xce, 0x85, 0xca, 0xad, 0xb8, 0xf7, 0xa6, 0xca, 0xd6, 0x46, 0xe4, 0x05,
	0x9a, 0x40, 0xe9, 0x0a, 0xb5, 0xcb, 0xcd, 0x26, 0x30, 0x32, 0xfe, 0x69, 0x29, 0x03, 0xd4, 0x4e,
	0xb9, 0x4d, 0x60, 0xd3, 0xdc, 0x05, 0xce, 0xaa, 0xd0, 0x52, 0xfa, 0x88, 0x68, 0x8c, 0x7c, 0x8c,
	0x69, 0x1c, 0xfb, 0x98, 0xd0, 0x38, 0xf1, 0x31, 0xa5, 0x71, 0xea, 0x63, 0x46, 0xe3, 0xcc, 0xc7,
	0x9c, 0xc6, 0xb9, 0x8f, 0x05, 0x8d, 0x0b, 0x1f, 0x4b, 0x1a, 0x97, 0xb4, 0x7c, 0x54, 0xbf, 0x63,
	0x44, 0xd9, 0xbc, 0x30, 0x0c, 0xa4, 0x73, 0x46, 0xdd, 0x96, 0x0e, 0x6d, 0x80, 0xba, 0xda, 0xbe,
	0xd9, 0xe5, 0xaf, 0x1d, 0x76, 0xfe, 0xf0, 0xc2, 0xe2, 0xdb, 0x97, 0x1f, 0xd7, 0x9c, 0xb3, 0x3d,
	0x2d, 0x33, 0x84, 0xde, 0x45, 0xef, 0x6a, 0xb8, 0xa2, 0x99, 0x9f, 0xb1, 0x81, 0xff, 0x15, 0x21,
	0xec, 0x10, 0xed, 0xfb, 0x16, 0xb6, 0x38, 0x82, 0xdd, 0x0e, 0x47, 0x2d, 0x1e, 0xc3, 0x5e, 0x87,
	0xc7, 0x2d, 0x9e, 0x40, 0xbf, 0xc3, 0x93, 0x16, 0x4f, 0x61, 0xd0, 0xe1, 0x69, 0x8b, 0x67, 0xb0,
	0xdf, 0xe1, 0x59, 0x8b, 0xe7, 0x70, 0xd0, 0xe1, 0x79, 0x8b, 0x17, 0x30, 0xec, 0xf0, 0xa2, 0xc5,
	0x4b, 0x60, 0x1d, 0x5e, 0xf2, 0x73, 0xb6, 0x5f, 0x5f, 0x67, 0x04, 0x87, 0xc4, 0xe9, 0xa9, 0x70,
	0xd4, 0x89, 0x10, 0x8e, 0xb6, 0x44, 0xd8, 0x89, 0x08, 0x8e, 0xb7, 0x44, 0x74, 0xf9, 0x67, 0x97,
	0xf1, 0x87, 0xff, 0x24, 0x8f, 0xd8, 0xd9, 0x9d, 0xc2, 0x34, 0x11, 0x95, 0x4c, 0x55, 0xa2, 0xdc,
	0x46, 0xdc, 0x2a, 0x97, 0xc9, 0x02, 0x22, 0x3a, 0xfd, 0x3f, 0xc9, 0xef, 0x8d, 0xfb, 0x44, 0x8a,
	0xbf, 0x64, 0x3c, 0xc1, 0x4a, 0xc5, 0x28, 0x12, 0xb4, 0xb1, 0x51, 0x85, 0x53, 0xb9, 0x86, 0x31,
	0x1d, 0x38, 0xa9, 0xcd, 0xe7, 0x4e, 0xf0, 0x53, 0xd6, 0x2f, 0xb5, 0x72, 0x16, 0x26, 0xf5, 0xd5,
	0xa8, 0xf0, 0x27, 0x6c, 0xd8, 0x2c, 0x51, 0x09, 0x4c, 0x2f, 0x7a, 0x57, 0xc7, 0xab, 0x83, 0x1a,
	0x7c, 0x4d, 0xfc, 0x91, 0x4a, 0xa6, 0x25, 0xc2, 0x8c, 0x44, 0x5d, 0xf8, 0x33, 0xc6, 0x64, 0x2a,
	0x4d, 0x26, 0xdc, 0xa6, 0x40, 0x98, 0x93, 0x1a, 0x12, 0xb9, 0xd9, 0x14, 0x48, 0x1b, 0xa5, 0x93,
	0xb5, 0x5d, 0x34, 0x1b, 0xa5, 0x93, 0x24, 0x4f, 0x59, 0xdf, 0xc6, 0x32, 0x45, 0x58, 0xd6, 0x1b,
	0xa9, 0xf0, 0xa7, 0x6c, 0x58, 0x18, 0x8c, 0x95, 0xf5, 0x17, 0x78, 0x55, 0x2f, 0x6c, 0x01, 0x7f,
	0xc4, 0x06, 0xd6, 0x49, 0x57, 0x5a, 0x78, 0x4d, 0xaa, 0x69, 0xfc, 0x05, 0xfb, 0x4f, 0xae, 0x51,
	0x38, 0x95, 0xa1, 0xb0, 0x4e, 0x66, 0x05, 0xbc, 0x21, 0x7f, 0x24, 0xd7, 0x78, 0xa3, 0x32, 0xbc,
	0xf6, 0x8c, 0x3f, 0x67, 0x87, 0x65, 0x91, 0x48, 0x87, 0xc2, 0x48, 0x87, 0xf0, 0x96, 0x1e, 0x61,
	0x35, 0x5a, 0x49, 0x87, 0x1c, 0xd8, 0xbe, 0xac, 0xd0, 0xc8, 0x35, 0xc2, 0xbb, 0x8b, 0xde, 0xd5,
	0xc9, 0xea, 0x5f, 0xf5, 0x26, 0x53, 0x5a, 0x65, 0x65, 0x06, 0xef, 0x6b, 0xd3, 0x54, 0x32, 0xf2,
	0x9e, 0xcc, 0x87, 0xc6, 0xd4, 0x95, 0x3f, 0x66, 0x07, 0x4a, 0x3b, 0x34, 0x95, 0x4c, 0xe1, 0x23,
	0xa9, 0xb6, 0xdf, 0x0e, 0xe8, 0x03, 0x31, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x92, 0x25,
	0x60, 0x3c, 0x04, 0x00, 0x00,
}
