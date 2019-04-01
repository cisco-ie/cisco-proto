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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_attributes_env_sensor_info

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
	proto.RegisterType((*EnvmonSensorInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.attributes.env_sensor_info.envmon_sensor_info_KEYS")
	proto.RegisterType((*EnvmonSensorInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.attributes.env_sensor_info.envmon_sensor_info")
}

func init() { proto.RegisterFile("envmon_sensor_info.proto", fileDescriptor_bc03e94ffc42a321) }

var fileDescriptor_bc03e94ffc42a321 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x86, 0x19, 0x75, 0x06, 0xa7, 0x5c, 0x05, 0xdb, 0x5d, 0x6d, 0x50, 0x71, 0x58, 0x3c, 0xcc,
	0xc5, 0xc0, 0xee, 0xfa, 0xfd, 0x71, 0x91, 0xf5, 0x20, 0xde, 0x66, 0x07, 0xc1, 0x53, 0x53, 0x93,
	0xd4, 0x0e, 0x85, 0x49, 0x77, 0xd3, 0x5d, 0x09, 0xe6, 0x57, 0xf8, 0x97, 0x25, 0x95, 0xa0, 0xc2,
	0x5c, 0x92, 0xbc, 0xef, 0x93, 0x7e, 0xaa, 0x1b, 0x1a, 0x2c, 0xf9, 0xae, 0x09, 0xde, 0x65, 0xf2,
	0x39, 0x24, 0xc7, 0xfe, 0x3a, 0x14, 0x31, 0x05, 0x09, 0x66, 0x5b, 0x72, 0x2e, 0x83, 0xe3, 0x90,
	0xdd, 0xaf, 0x01, 0x74, 0xcd, 0x3e, 0xb9, 0x10, 0x29, 0x15, 0xec, 0x3b, 0xf2, 0x12, 0x52, 0x5f,
	0x24, 0x2c, 0x7f, 0x66, 0x7d, 0x16, 0xe4, 0x85, 0xa5, 0x2f, 0x72, 0x1d, 0xa4, 0x40, 0x91, 0xc4,
	0xbb, 0x56, 0x28, 0x17, 0xe4, 0xbb, 0xff, 0xdd, 0xa7, 0x97, 0xf0, 0xe8, 0x70, 0xa2, 0xfb, 0xf6,
	0xe5, 0xc7, 0x95, 0x31, 0x70, 0xcb, 0x63, 0x43, 0x76, 0xb6, 0x9a, 0xad, 0x97, 0x1b, 0xfd, 0x36,
	0x27, 0xb0, 0x18, 0xde, 0xee, 0xcc, 0xde, 0xd0, 0x76, 0x3e, 0xa4, 0xb3, 0xd3, 0xdf, 0x37, 0xc1,
	0x1c, 0x6a, 0xcc, 0x39, 0x9c, 0x5c, 0x33, 0xd5, 0x95, 0xeb, 0xb0, 0xe6, 0x8a, 0xa5, 0x77, 0x3b,
	0x96, 0x06, 0xa3, 0x3d, 0x5f, 0xcd, 0xd6, 0x47, 0x9b, 0x07, 0x0a, 0xbf, 0x4f, 0xec, 0xb3, 0x22,
	0xf3, 0x02, 0x4c, 0x45, 0x1d, 0x97, 0xe4, 0x2a, 0xca, 0x65, 0xe2, 0x28, 0x1c, 0xbc, 0xbd, 0xd0,
	0x69, 0xf7, 0x47, 0x72, 0xf9, 0x0f, 0x98, 0x63, 0x98, 0xb7, 0x9e, 0x25, 0xdb, 0x97, 0xe3, 0x7e,
	0x34, 0x98, 0xc7, 0xb0, 0x9c, 0x24, 0x5c, 0xd9, 0x57, 0xab, 0xd9, 0xfa, 0xee, 0xe6, 0xf6, 0x58,
	0x7c, 0xad, 0x86, 0x25, 0x1d, 0xd6, 0x2d, 0xd9, 0xd7, 0x0a, 0xc6, 0x60, 0x9e, 0x02, 0x60, 0x8d,
	0xa9, 0x71, 0xd2, 0x47, 0xb2, 0x6f, 0x14, 0x2d, 0xb5, 0xd9, 0xf6, 0x91, 0xd4, 0x88, 0x82, 0x23,
	0x7d, 0x3b, 0x19, 0x51, 0x50, 0xe1, 0x31, 0xcc, 0x73, 0x89, 0x35, 0xd9, 0x77, 0xa3, 0x51, 0x83,
	0x79, 0x02, 0xcb, 0x98, 0xa8, 0xe4, 0x3c, 0x1c, 0xe0, 0xfd, 0x28, 0xfc, 0x5b, 0x98, 0x87, 0xb0,
	0xc8, 0x82, 0xd2, 0x66, 0xfb, 0x41, 0xd1, 0x94, 0xcc, 0x73, 0xb8, 0x87, 0x7b, 0x72, 0xc2, 0x0d,
	0xb9, 0x2c, 0xd8, 0x44, 0xfb, 0x51, 0xf9, 0x11, 0xee, 0x69, 0xcb, 0x0d, 0x5d, 0x0d, 0x9d, 0x79,
	0x06, 0x77, 0xda, 0x58, 0xa1, 0x90, 0x4b, 0x28, 0x64, 0x3f, 0xe9, 0x2f, 0x30, 0x56, 0x1b, 0x14,
	0xda, 0x2d, 0xf4, 0xd2, 0x5c, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x87, 0xff, 0xbc, 0x50,
	0x02, 0x00, 0x00,
}
