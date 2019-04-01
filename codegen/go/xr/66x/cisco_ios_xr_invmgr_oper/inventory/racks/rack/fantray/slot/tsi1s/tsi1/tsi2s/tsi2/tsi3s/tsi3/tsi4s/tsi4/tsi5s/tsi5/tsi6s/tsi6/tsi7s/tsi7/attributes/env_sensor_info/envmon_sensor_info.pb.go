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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_attributes_env_sensor_info

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
	proto.RegisterType((*EnvmonSensorInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.attributes.env_sensor_info.envmon_sensor_info_KEYS")
	proto.RegisterType((*EnvmonSensorInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.attributes.env_sensor_info.envmon_sensor_info")
}

func init() { proto.RegisterFile("envmon_sensor_info.proto", fileDescriptor_bc03e94ffc42a321) }

var fileDescriptor_bc03e94ffc42a321 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0x35, 0xb4, 0x33, 0xed, 0x98, 0x8b, 0x54, 0xd3, 0xc2, 0x11, 0x17, 0x51, 0x55, 0x2c,
	0xba, 0x21, 0x52, 0x93, 0xb9, 0x94, 0xfb, 0x45, 0xb0, 0x40, 0xec, 0xa6, 0x15, 0x12, 0x2b, 0xeb,
	0x34, 0x39, 0x33, 0xb2, 0x48, 0x9c, 0xc8, 0x76, 0xa2, 0xce, 0xeb, 0xf0, 0x6e, 0x3c, 0x04, 0x3b,
	0xe4, 0x93, 0x90, 0x20, 0x75, 0xf3, 0xcf, 0xf9, 0xbf, 0x6f, 0x7c, 0x64, 0x2f, 0x22, 0x80, 0x4c,
	0x53, 0x94, 0x46, 0x39, 0x32, 0xae, 0xb4, 0x4a, 0x9b, 0x75, 0x19, 0x55, 0xb6, 0xf4, 0xa5, 0xfc,
	0x35, 0x4a, 0xb5, 0x4b, 0x4b, 0xa5, 0x4b, 0xa7, 0xae, 0x83, 0x69, 0x8a, 0x8d, 0x55, 0x65, 0x45,
	0x36, 0xd2, 0xa6, 0x21, 0xe3, 0x4b, 0xbb, 0x8d, 0x2c, 0xa6, 0x3f, 0x1d, 0x67, 0xb4, 0x46, 0xe3,
	0x2d, 0x6e, 0x23, 0x97, 0x97, 0x3e, 0xf2, 0x4e, 0x9f, 0x39, 0xce, 0x10, 0x31, 0x8f, 0x71, 0x88,
	0x84, 0xc7, 0x24, 0xc4, 0x8c, 0xc7, 0x59, 0x88, 0x39, 0x8f, 0xf3, 0x10, 0x0b, 0x1e, 0x17, 0x21,
	0x96, 0x3c, 0x2e, 0x23, 0xf4, 0xde, 0xea, 0xab, 0xda, 0x93, 0x8b, 0xc8, 0x34, 0xff, 0x5f, 0xf5,
	0xe4, 0xf7, 0x48, 0x3c, 0xbc, 0xf9, 0x02, 0xf5, 0xed, 0xcb, 0x8f, 0x0b, 0x29, 0xc5, 0xae, 0xc1,
	0x82, 0x60, 0x74, 0x3c, 0x3a, 0x9d, 0xae, 0x78, 0x96, 0x47, 0x62, 0x12, 0x7e, 0xd5, 0x19, 0xdc,
	0x62, 0x3a, 0x0e, 0xed, 0xac, 0xc7, 0x31, 0xec, 0x0c, 0x38, 0xee, 0x71, 0x02, 0xbb, 0x03, 0x4e,
	0x7a, 0x3c, 0x83, 0xf1, 0x80, 0x67, 0x3d, 0x9e, 0xc3, 0x64, 0xc0, 0xf3, 0x1e, 0x2f, 0x60, 0x6f,
	0xc0, 0x8b, 0x1e, 0x2f, 0x61, 0x7f, 0xc0, 0xcb, 0x1e, 0x9f, 0xc3, 0x74, 0xc0, 0xe7, 0x27, 0x7f,
	0x76, 0x84, 0xbc, 0xf9, 0x4e, 0x19, 0x8b, 0xa3, 0xb5, 0xa6, 0x3c, 0x53, 0x0d, 0xe6, 0x3a, 0xd3,
	0x7e, 0xab, 0xae, 0xb4, 0x2f, 0xb0, 0x82, 0x98, 0x0f, 0xdf, 0x67, 0xf9, 0xbd, 0x73, 0x9f, 0x58,
	0xc9, 0x17, 0x42, 0x66, 0xd4, 0xe8, 0x94, 0x54, 0x46, 0x2e, 0xb5, 0xba, 0xf2, 0xba, 0x34, 0x90,
	0xf0, 0x81, 0x83, 0xd6, 0x7c, 0x1e, 0x84, 0x3c, 0x14, 0xe3, 0xda, 0x68, 0xef, 0x60, 0xd6, 0xde,
	0x87, 0x8b, 0x7c, 0x2c, 0xa6, 0xdd, 0x12, 0x9d, 0xc1, 0xfc, 0x78, 0x74, 0x7a, 0x77, 0xb5, 0xdf,
	0x82, 0xaf, 0x59, 0x38, 0xd2, 0x60, 0x5e, 0x13, 0x2c, 0x58, 0xb4, 0x45, 0x3e, 0x15, 0x02, 0x73,
	0xb4, 0x85, 0xf2, 0xdb, 0x8a, 0x60, 0xc9, 0x6a, 0xca, 0xe4, 0x72, 0x5b, 0x11, 0x6f, 0x44, 0x8f,
	0xad, 0x3d, 0xef, 0x36, 0xa2, 0x47, 0x96, 0x87, 0x62, 0xec, 0x52, 0xcc, 0x09, 0x5e, 0xb6, 0x1b,
	0xb9, 0xc8, 0x27, 0x62, 0x5a, 0x59, 0x4a, 0xb5, 0x0b, 0x0f, 0x78, 0xd5, 0x2e, 0xec, 0x81, 0x7c,
	0x20, 0x26, 0xce, 0xa3, 0xaf, 0x1d, 0xbc, 0x66, 0xd5, 0x35, 0xf9, 0x5c, 0xdc, 0xc3, 0x0d, 0x29,
	0xaf, 0x0b, 0x52, 0xce, 0x63, 0x51, 0xc1, 0x1b, 0xf6, 0x77, 0x70, 0x43, 0x97, 0xba, 0xa0, 0x8b,
	0xc0, 0xe4, 0x33, 0x71, 0xbb, 0xae, 0x32, 0xf4, 0xa4, 0x2c, 0x7a, 0x82, 0xb7, 0xfc, 0x17, 0xd1,
	0xa2, 0x15, 0x7a, 0x92, 0x20, 0xf6, 0xb0, 0x21, 0x8b, 0x1b, 0x82, 0x77, 0xc7, 0xa3, 0xd3, 0x83,
	0xd5, 0xbf, 0x1a, 0x4c, 0xa1, 0x8d, 0x2e, 0xea, 0x02, 0xde, 0xb7, 0xa6, 0xab, 0x6c, 0xf0, 0x9a,
	0xcd, 0x87, 0xce, 0xb4, 0x55, 0x3e, 0x12, 0xfb, 0xda, 0x78, 0xb2, 0x0d, 0xe6, 0xf0, 0x91, 0x55,
	0xdf, 0xaf, 0x26, 0xfc, 0x3d, 0x26, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x79, 0x2c, 0xa0, 0xf7,
	0xab, 0x03, 0x00, 0x00,
}