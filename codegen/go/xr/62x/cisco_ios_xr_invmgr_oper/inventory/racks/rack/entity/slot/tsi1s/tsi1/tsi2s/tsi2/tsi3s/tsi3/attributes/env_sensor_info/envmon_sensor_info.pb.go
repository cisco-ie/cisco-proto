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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_attributes_env_sensor_info

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

type EnvmonSensorInfo struct {
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

func (m *EnvmonSensorInfo) GetFieldValidityBitmap() []byte {
	if m != nil {
		return m.FieldValidityBitmap
	}
	return nil
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

func init() {
	proto.RegisterType((*EnvmonSensorInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.env_sensor_info.envmon_sensor_info_KEYS")
	proto.RegisterType((*EnvmonSensorInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.env_sensor_info.envmon_sensor_info")
}

func init() { proto.RegisterFile("envmon_sensor_info.proto", fileDescriptor_bc03e94ffc42a321) }

var fileDescriptor_bc03e94ffc42a321 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x89, 0x77, 0x2d, 0x76, 0x3d, 0x05, 0xd7, 0x3b, 0x5d, 0x50, 0xb1, 0x1c, 0x3e, 0xf4,
	0xc5, 0x40, 0x93, 0xfa, 0x5b, 0x5f, 0x44, 0x1f, 0xc4, 0xb7, 0xde, 0x21, 0xf8, 0xb4, 0x4c, 0x93,
	0xb9, 0x32, 0x98, 0xec, 0x86, 0xdd, 0x49, 0x30, 0x7f, 0x82, 0x4f, 0xfe, 0xcb, 0x92, 0x49, 0x69,
	0x0f, 0xfa, 0x32, 0x99, 0xf9, 0x7c, 0x92, 0xef, 0x4e, 0x60, 0x95, 0x41, 0xd7, 0xd5, 0xde, 0xd9,
	0x88, 0x2e, 0xfa, 0x60, 0xc9, 0xdd, 0xf8, 0xb4, 0x09, 0x9e, 0xbd, 0x6e, 0x0b, 0x8a, 0x85, 0xb7,
	0xe4, 0xa3, 0xfd, 0x33, 0x88, 0xae, 0xde, 0x06, 0xeb, 0x1b, 0x0c, 0x29, 0xb9, 0x0e, 0x1d, 0xfb,
	0xd0, 0xa7, 0x01, 0x8a, 0xdf, 0x51, 0x6a, 0x8a, 0x8e, 0x89, 0xfb, 0x34, 0x56, 0x9e, 0x53, 0x8e,
	0xb4, 0x8c, 0x52, 0x87, 0x92, 0x49, 0x9b, 0x0d, 0x25, 0x97, 0x36, 0x4f, 0x81, 0x39, 0xd0, 0xa6,
	0x65, 0x8c, 0x29, 0xba, 0xee, 0xf6, 0xe1, 0x97, 0x7f, 0x13, 0xf5, 0xe4, 0x78, 0x27, 0xfb, 0xe3,
	0xdb, 0xaf, 0x2b, 0xad, 0xd5, 0xa9, 0x83, 0x1a, 0x4d, 0x32, 0x4f, 0x16, 0xb3, 0xb5, 0xf4, 0xfa,
	0x42, 0x4d, 0x87, 0xa7, 0x5d, 0x9a, 0x3b, 0x42, 0x27, 0xc3, 0xb4, 0xdc, 0xe3, 0xcc, 0x9c, 0x1c,
	0x70, 0xb6, 0xc7, 0xb9, 0x39, 0x3d, 0xe0, 0x7c, 0x8f, 0x57, 0x66, 0x72, 0xc0, 0xab, 0xcb, 0x7f,
	0x27, 0x4a, 0x1f, 0xef, 0xa2, 0x33, 0x75, 0x71, 0x43, 0x58, 0x95, 0xb6, 0x83, 0x8a, 0x4a, 0xe2,
	0xde, 0x6e, 0x88, 0x6b, 0x68, 0x4c, 0x36, 0x4f, 0x16, 0x67, 0xeb, 0x47, 0x22, 0x7f, 0xee, 0xdc,
	0x17, 0x51, 0xfa, 0x95, 0xd2, 0x25, 0x76, 0x54, 0xa0, 0x2d, 0x31, 0x16, 0x81, 0x1a, 0x26, 0xef,
	0x4c, 0x2e, 0xa7, 0x3d, 0x1c, 0xcd, 0xd7, 0x83, 0xd0, 0xe7, 0x6a, 0xd2, 0x3a, 0xe2, 0x68, 0x56,
	0xe3, 0x3e, 0x32, 0xe8, 0xa7, 0x6a, 0xb6, 0x0b, 0xa1, 0xd2, 0xbc, 0x9e, 0x27, 0x8b, 0xfb, 0xeb,
	0xbb, 0x23, 0xf8, 0x5e, 0x0e, 0x9f, 0x74, 0x50, 0xb5, 0x68, 0xde, 0x88, 0x18, 0x07, 0xfd, 0x5c,
	0x29, 0xa8, 0x20, 0xd4, 0x96, 0xfb, 0x06, 0xcd, 0x5b, 0x51, 0x33, 0x21, 0xd7, 0x7d, 0x83, 0x92,
	0x08, 0x0c, 0xa3, 0x7d, 0xb7, 0x4b, 0x04, 0x06, 0x91, 0xe7, 0x6a, 0x12, 0x0b, 0xa8, 0xd0, 0xbc,
	0x1f, 0x13, 0x65, 0xd0, 0xcf, 0xd4, 0xac, 0x09, 0x58, 0x50, 0x1c, 0x7e, 0xe0, 0xc3, 0x18, 0xb8,
	0x07, 0xfa, 0xb1, 0x9a, 0x46, 0x06, 0x6e, 0xa3, 0xf9, 0x28, 0x6a, 0x37, 0xe9, 0x97, 0xea, 0x01,
	0x6c, 0xd1, 0x32, 0xd5, 0x68, 0x23, 0x43, 0xdd, 0x98, 0x4f, 0xe2, 0xcf, 0x60, 0x8b, 0xd7, 0x54,
	0xe3, 0xd5, 0xc0, 0xf4, 0x0b, 0x75, 0xaf, 0x6d, 0x4a, 0x60, 0xb4, 0x01, 0x18, 0xcd, 0x67, 0x79,
	0x45, 0x8d, 0x68, 0x0d, 0x8c, 0x9b, 0xa9, 0xdc, 0xcd, 0xfc, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe8, 0x2a, 0xe3, 0x7a, 0xb7, 0x02, 0x00, 0x00,
}
