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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_attributes_env_sensor_info

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
	proto.RegisterType((*EnvmonSensorInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.attributes.env_sensor_info.envmon_sensor_info_KEYS")
	proto.RegisterType((*EnvmonSensorInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.attributes.env_sensor_info.envmon_sensor_info")
}

func init() { proto.RegisterFile("envmon_sensor_info.proto", fileDescriptor_bc03e94ffc42a321) }

var fileDescriptor_bc03e94ffc42a321 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x8f, 0xd3, 0x3c,
	0x10, 0x87, 0x95, 0x77, 0xb7, 0xd5, 0x5b, 0xb3, 0x20, 0x61, 0x76, 0xc1, 0x12, 0x20, 0xaa, 0x15,
	0x87, 0x5e, 0x88, 0xd4, 0xa4, 0xe5, 0x3f, 0x17, 0x04, 0x07, 0xc4, 0xad, 0xbb, 0x42, 0xe2, 0x64,
	0x4d, 0x93, 0x69, 0xb1, 0x48, 0x6c, 0xcb, 0x9e, 0x04, 0x7a, 0xe7, 0xce, 0x57, 0xe0, 0xa3, 0xa2,
	0x4c, 0xaa, 0x04, 0x69, 0x2f, 0x93, 0x99, 0xe7, 0x49, 0x7e, 0x19, 0x4b, 0x16, 0x0a, 0x6d, 0x5b,
	0x3b, 0xab, 0x23, 0xda, 0xe8, 0x82, 0x36, 0x76, 0xe7, 0x52, 0x1f, 0x1c, 0x39, 0xf9, 0x2b, 0x29,
	0x4c, 0x2c, 0x9c, 0x36, 0x2e, 0xea, 0x9f, 0x9d, 0x69, 0xeb, 0x7d, 0xd0, 0xce, 0x63, 0x48, 0x8d,
	0x6d, 0xd1, 0x92, 0x0b, 0x87, 0x34, 0x40, 0xf1, 0x3d, 0x72, 0x4d, 0xbd, 0xfb, 0x81, 0x21, 0x7e,
	0xc3, 0x6a, 0x97, 0xc6, 0xca, 0x51, 0x4a, 0xd1, 0x2c, 0x23, 0xd7, 0xae, 0x64, 0xdc, 0x66, 0x5d,
	0xc9, 0xb9, 0xcd, 0xbb, 0xb2, 0xe2, 0x76, 0x95, 0x02, 0x51, 0x30, 0xdb, 0x86, 0x30, 0xa6, 0x68,
	0xdb, 0x7f, 0x97, 0xb9, 0xfc, 0x93, 0x88, 0x07, 0x37, 0x77, 0xd4, 0x9f, 0x3f, 0x7e, 0xbd, 0x92,
	0x52, 0x9c, 0x5a, 0xa8, 0x51, 0x25, 0xf3, 0x64, 0x31, 0xdb, 0x70, 0x2f, 0x2f, 0xc4, 0xb4, 0x7b,
	0xea, 0xa5, 0xfa, 0x8f, 0xe9, 0xa4, 0x9b, 0x96, 0x03, 0xce, 0xd4, 0xc9, 0x88, 0xb3, 0x01, 0xe7,
	0xea, 0x74, 0xc4, 0xf9, 0x80, 0x57, 0x6a, 0x32, 0xe2, 0xd5, 0x80, 0xd7, 0x6a, 0x3a, 0xe2, 0xf5,
	0xe5, 0xef, 0x13, 0x21, 0x6f, 0xae, 0x28, 0x33, 0x71, 0xb1, 0x33, 0x58, 0x95, 0xba, 0x85, 0xca,
	0x94, 0x86, 0x0e, 0x7a, 0x6b, 0xa8, 0x06, 0xaf, 0xb2, 0x79, 0xb2, 0x38, 0xdb, 0xdc, 0x63, 0xf9,
	0xe5, 0xe8, 0xde, 0xb3, 0x92, 0xcf, 0x84, 0x2c, 0xb1, 0x35, 0x05, 0xea, 0x12, 0x63, 0x11, 0x8c,
	0x27, 0xe3, 0xac, 0xca, 0xf9, 0x6f, 0x77, 0x7b, 0xf3, 0x61, 0x14, 0xf2, 0x5c, 0x4c, 0x1a, 0x6b,
	0x28, 0xaa, 0x55, 0xbf, 0x0f, 0x0f, 0xf2, 0xa1, 0x98, 0x1d, 0x43, 0x4c, 0xa9, 0xd6, 0xf3, 0x64,
	0x71, 0x7b, 0xf3, 0x7f, 0x0f, 0x3e, 0x95, 0xdd, 0x27, 0x2d, 0x54, 0x0d, 0xaa, 0xe7, 0x2c, 0xfa,
	0x41, 0x3e, 0x16, 0x02, 0x2a, 0x08, 0xb5, 0xa6, 0x83, 0x47, 0xf5, 0x82, 0xd5, 0x8c, 0xc9, 0xf5,
	0xc1, 0x23, 0x27, 0x02, 0x41, 0x6f, 0x5f, 0x1e, 0x13, 0x81, 0x80, 0xe5, 0xb9, 0x98, 0xc4, 0x02,
	0x2a, 0x54, 0xaf, 0xfa, 0x44, 0x1e, 0xe4, 0x23, 0x31, 0xf3, 0x01, 0x0b, 0x13, 0xbb, 0x03, 0xbc,
	0xee, 0x03, 0x07, 0x20, 0xef, 0x8b, 0x69, 0x24, 0xa0, 0x26, 0xaa, 0x37, 0xac, 0x8e, 0x93, 0x7c,
	0x2a, 0xee, 0xc0, 0x1e, 0x35, 0x99, 0x1a, 0x75, 0x24, 0xa8, 0xbd, 0x7a, 0xcb, 0xfe, 0x0c, 0xf6,
	0x78, 0x6d, 0x6a, 0xbc, 0xea, 0x98, 0x7c, 0x22, 0x6e, 0x35, 0xbe, 0x04, 0x42, 0x1d, 0x80, 0x50,
	0xbd, 0xe3, 0x57, 0x44, 0x8f, 0x36, 0x40, 0xb8, 0x9d, 0xf2, 0x15, 0xce, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xd8, 0x68, 0x93, 0xa3, 0xde, 0x02, 0x00, 0x00,
}