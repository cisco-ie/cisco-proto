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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_attributes_env_sensor_info

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
	proto.RegisterType((*EnvmonSensorInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.env_sensor_info.envmon_sensor_info_KEYS")
	proto.RegisterType((*EnvmonSensorInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.env_sensor_info.envmon_sensor_info")
}

func init() { proto.RegisterFile("envmon_sensor_info.proto", fileDescriptor_bc03e94ffc42a321) }

var fileDescriptor_bc03e94ffc42a321 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcb, 0x6f, 0xd3, 0x4e,
	0x10, 0xc7, 0x95, 0x5f, 0x7f, 0x09, 0xcd, 0xf2, 0x90, 0xba, 0xb4, 0xb0, 0xe2, 0x21, 0xa2, 0x8a,
	0x43, 0x2e, 0x58, 0x4a, 0xca, 0xfb, 0x0d, 0x82, 0x03, 0xe2, 0x96, 0x56, 0x48, 0x88, 0xc3, 0x32,
	0xb1, 0x27, 0xd1, 0x08, 0xef, 0xae, 0xb5, 0x3b, 0xb6, 0xea, 0x3f, 0x9d, 0x1b, 0xf2, 0xd8, 0xa4,
	0x48, 0xbd, 0xac, 0xfd, 0xf9, 0x7e, 0x3c, 0xdf, 0xb1, 0x25, 0x2b, 0x83, 0xbe, 0x71, 0xc1, 0xdb,
	0x84, 0x3e, 0x85, 0x68, 0xc9, 0x6f, 0x42, 0x56, 0xc5, 0xc0, 0x41, 0xff, 0xcc, 0x29, 0xe5, 0xc1,
	0x52, 0x48, 0xf6, 0xbc, 0x13, 0x8d, 0xdb, 0x46, 0x1b, 0x2a, 0x8c, 0x19, 0xf9, 0x06, 0x3d, 0x87,
	0xd8, 0x66, 0x11, 0xf2, 0x5f, 0x49, 0xce, 0x6c, 0x03, 0x9e, 0x23, 0xb4, 0x59, 0x2a, 0x03, 0x67,
	0x9c, 0x68, 0x91, 0xe4, 0xcc, 0x80, 0x39, 0xd2, 0xba, 0x66, 0x4c, 0x19, 0xfa, 0xe6, 0xdf, 0x3d,
	0xc7, 0x3f, 0xd4, 0xed, 0xcb, 0xdb, 0xed, 0xd7, 0xcf, 0xdf, 0x4f, 0xb5, 0x56, 0xff, 0x7b, 0x70,
	0x68, 0x46, 0xb3, 0xd1, 0x7c, 0xba, 0x92, 0x7b, 0x7d, 0xa4, 0x26, 0xdd, 0xd5, 0x2e, 0xcc, 0x7f,
	0x92, 0x8e, 0x3b, 0x5a, 0xec, 0xe2, 0xa5, 0xd9, 0xbb, 0x88, 0x97, 0xc7, 0xbf, 0xf7, 0x94, 0xbe,
	0xdc, 0xae, 0x97, 0xea, 0x68, 0x43, 0x58, 0x16, 0xb6, 0x81, 0x92, 0x0a, 0xe2, 0xd6, 0xae, 0x89,
	0x1d, 0x54, 0x66, 0x29, 0xc3, 0x37, 0x45, 0x7e, 0x1b, 0xdc, 0x47, 0x51, 0xfa, 0x91, 0xd2, 0x05,
	0x36, 0x94, 0xa3, 0x2d, 0x30, 0xe5, 0x91, 0x2a, 0xa6, 0xe0, 0xcd, 0x89, 0x0c, 0x1c, 0xf4, 0xe6,
	0xd3, 0x85, 0xd0, 0x87, 0x6a, 0x5c, 0x7b, 0xe2, 0x64, 0x1e, 0xf7, 0xef, 0x23, 0xa0, 0xef, 0xaa,
	0xe9, 0x50, 0x42, 0x85, 0x79, 0x32, 0x1b, 0xcd, 0xaf, 0xaf, 0xf6, 0xfb, 0xe0, 0x4b, 0xd1, 0x8d,
	0x34, 0x50, 0xd6, 0x68, 0x9e, 0x8a, 0xe8, 0x41, 0xdf, 0x57, 0x0a, 0x4a, 0x88, 0xce, 0x72, 0x5b,
	0xa1, 0x79, 0x26, 0x6a, 0x2a, 0xc9, 0x59, 0x5b, 0xa1, 0x34, 0x02, 0x43, 0x6f, 0x9f, 0x0f, 0x8d,
	0xc0, 0x20, 0xf2, 0x50, 0x8d, 0x53, 0x0e, 0x25, 0x9a, 0x17, 0x7d, 0xa3, 0x80, 0xbe, 0xa7, 0xa6,
	0x55, 0xc4, 0x9c, 0x52, 0xf7, 0x01, 0x2f, 0xfb, 0xc2, 0x5d, 0xa0, 0x6f, 0xa9, 0x49, 0x62, 0xe0,
	0x3a, 0x99, 0x57, 0xa2, 0x06, 0xd2, 0x0f, 0xd5, 0x0d, 0xd8, 0xa2, 0x65, 0x72, 0x68, 0x13, 0x83,
	0xab, 0xcc, 0x6b, 0xf1, 0xd7, 0x60, 0x8b, 0x67, 0xe4, 0xf0, 0xb4, 0xcb, 0xf4, 0x03, 0x75, 0xb5,
	0xae, 0x0a, 0x60, 0xb4, 0x11, 0x18, 0xcd, 0x1b, 0x79, 0x44, 0xf5, 0xd1, 0x0a, 0x18, 0xb5, 0x51,
	0x57, 0xa0, 0xc1, 0x08, 0x5b, 0x34, 0x6f, 0x67, 0xa3, 0xf9, 0xc1, 0xea, 0x2f, 0x76, 0xc6, 0x91,
	0x27, 0x57, 0x3b, 0xf3, 0xae, 0x37, 0x03, 0x8a, 0x81, 0x73, 0x31, 0xef, 0x07, 0xd3, 0xa3, 0xbe,
	0xa3, 0xf6, 0xc9, 0x33, 0xc6, 0x06, 0x4a, 0xf3, 0x41, 0xd4, 0x8e, 0xd7, 0x13, 0xf9, 0x83, 0x4f,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x64, 0x5b, 0x8d, 0xd1, 0xdd, 0x02, 0x00, 0x00,
}
