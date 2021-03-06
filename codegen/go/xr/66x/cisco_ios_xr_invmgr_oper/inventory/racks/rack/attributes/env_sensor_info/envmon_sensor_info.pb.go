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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_attributes_env_sensor_info

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
	proto.RegisterType((*EnvmonSensorInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.attributes.env_sensor_info.envmon_sensor_info_KEYS")
	proto.RegisterType((*EnvmonSensorInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.attributes.env_sensor_info.envmon_sensor_info")
}

func init() { proto.RegisterFile("envmon_sensor_info.proto", fileDescriptor_bc03e94ffc42a321) }

var fileDescriptor_bc03e94ffc42a321 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd2, 0xcb, 0x6b, 0x14, 0x41,
	0x10, 0x06, 0x70, 0x16, 0x4d, 0xcc, 0x96, 0x0f, 0x48, 0x1b, 0xb5, 0xf1, 0x81, 0x4b, 0xf0, 0xb0,
	0x97, 0xec, 0x21, 0xf1, 0xfd, 0x56, 0x14, 0x14, 0x6f, 0x9b, 0x20, 0x78, 0x6a, 0x6a, 0x67, 0x2a,
	0x4b, 0xe1, 0x74, 0xf7, 0xd0, 0x5d, 0x33, 0x64, 0xff, 0x74, 0x6f, 0x32, 0xd5, 0x63, 0x0c, 0xe4,
	0xb2, 0xec, 0x57, 0xbf, 0xaa, 0x6f, 0x18, 0x18, 0xb0, 0x14, 0x7a, 0x1f, 0x83, 0xcb, 0x14, 0x72,
	0x4c, 0x8e, 0xc3, 0x69, 0x5c, 0xb4, 0x29, 0x4a, 0x34, 0xdf, 0x2a, 0xce, 0x55, 0x74, 0x1c, 0xb3,
	0x3b, 0x1b, 0xa0, 0xf7, 0xeb, 0xe4, 0x62, 0x4b, 0x69, 0xc1, 0xa1, 0xa7, 0x20, 0x31, 0x6d, 0x16,
	0x09, 0xab, 0xdf, 0x59, 0x7f, 0x17, 0x28, 0x92, 0x78, 0xd5, 0x09, 0xe5, 0x05, 0x85, 0xfe, 0x62,
	0xdf, 0xfe, 0x01, 0xdc, 0xbb, 0xfc, 0x14, 0xf7, 0xe3, 0xeb, 0xaf, 0x63, 0x63, 0xe0, 0x6a, 0x40,
	0x4f, 0x76, 0x32, 0x9b, 0xcc, 0xa7, 0x4b, 0xfd, 0xbf, 0xff, 0xe7, 0x0a, 0x98, 0xcb, 0xfb, 0xe6,
	0x10, 0xee, 0x9c, 0x32, 0x35, 0xb5, 0xeb, 0xb1, 0xe1, 0x9a, 0x65, 0xe3, 0x56, 0x2c, 0x1e, 0x5b,
	0x7b, 0xa8, 0xb7, 0xb7, 0x15, 0x7f, 0x8e, 0xf6, 0x59, 0xc9, 0x1c, 0x80, 0xa9, 0xa9, 0xe7, 0x8a,
	0x5c, 0x4d, 0xb9, 0x4a, 0xdc, 0x0a, 0xc7, 0x60, 0x8f, 0xf4, 0x60, 0xb7, 0xc8, 0x97, 0xff, 0x60,
	0xf6, 0x60, 0xab, 0x0b, 0x2c, 0xd9, 0x3e, 0xd5, 0x8d, 0x12, 0xcc, 0x03, 0x98, 0x8e, 0x25, 0x5c,
	0xdb, 0x67, 0xb3, 0xc9, 0xfc, 0xe6, 0x72, 0xa7, 0x0c, 0xbe, 0xd7, 0xc3, 0x49, 0x8f, 0x4d, 0x47,
	0xf6, 0xb9, 0x42, 0x09, 0xe6, 0x11, 0x00, 0x36, 0x98, 0xbc, 0x93, 0x4d, 0x4b, 0xf6, 0x85, 0xd2,
	0x54, 0x27, 0x27, 0x9b, 0x96, 0xb4, 0x11, 0x05, 0x8b, 0xbe, 0x1c, 0x1b, 0x51, 0x50, 0x71, 0x0f,
	0xb6, 0x72, 0x85, 0x0d, 0xd9, 0x57, 0xa5, 0x51, 0x83, 0x79, 0x08, 0xd3, 0x36, 0x51, 0xc5, 0x79,
	0x78, 0x81, 0xd7, 0xa5, 0xf0, 0x7c, 0x60, 0xee, 0xc2, 0x76, 0x16, 0x94, 0x2e, 0xdb, 0x37, 0x4a,
	0x63, 0x32, 0x4f, 0xe0, 0x16, 0xae, 0xc9, 0x09, 0x7b, 0x72, 0x59, 0xd0, 0xb7, 0xf6, 0xad, 0xfa,
	0x0d, 0x5c, 0xd3, 0x09, 0x7b, 0x3a, 0x1e, 0x66, 0xe6, 0x31, 0x5c, 0xef, 0xda, 0x1a, 0x85, 0x5c,
	0x42, 0x21, 0xfb, 0x4e, 0x57, 0xa0, 0x8c, 0x96, 0x28, 0x64, 0x2c, 0x5c, 0xc3, 0x9e, 0x12, 0xae,
	0xc9, 0xbe, 0x9f, 0x4d, 0xe6, 0xbb, 0xcb, 0x7f, 0x71, 0x10, 0xcf, 0x81, 0x7d, 0xe7, 0xed, 0x87,
	0x22, 0x63, 0x54, 0xc1, 0x33, 0x95, 0x8f, 0xa3, 0x94, 0x68, 0xee, 0xc3, 0x0e, 0x07, 0xa1, 0xd4,
	0x63, 0x63, 0x3f, 0x29, 0x9d, 0xe7, 0xd5, 0xb6, 0x7e, 0x7b, 0x47, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0xeb, 0xda, 0xa2, 0x97, 0x02, 0x00, 0x00,
}
