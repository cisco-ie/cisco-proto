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
// source: obj_group_port_member_port.proto

package cisco_ios_xr_infra_objmgr_oper_object_group_port_objects_object_operators_operator

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

type ObjGroupPortMemberPort_KEYS struct {
	ObjectName           string   `protobuf:"bytes,1,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	OperatorType         string   `protobuf:"bytes,2,opt,name=operator_type,json=operatorType,proto3" json:"operator_type,omitempty"`
	Port                 string   `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjGroupPortMemberPort_KEYS) Reset()         { *m = ObjGroupPortMemberPort_KEYS{} }
func (m *ObjGroupPortMemberPort_KEYS) String() string { return proto.CompactTextString(m) }
func (*ObjGroupPortMemberPort_KEYS) ProtoMessage()    {}
func (*ObjGroupPortMemberPort_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fe91f737dc2c241, []int{0}
}

func (m *ObjGroupPortMemberPort_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjGroupPortMemberPort_KEYS.Unmarshal(m, b)
}
func (m *ObjGroupPortMemberPort_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjGroupPortMemberPort_KEYS.Marshal(b, m, deterministic)
}
func (m *ObjGroupPortMemberPort_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjGroupPortMemberPort_KEYS.Merge(m, src)
}
func (m *ObjGroupPortMemberPort_KEYS) XXX_Size() int {
	return xxx_messageInfo_ObjGroupPortMemberPort_KEYS.Size(m)
}
func (m *ObjGroupPortMemberPort_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjGroupPortMemberPort_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ObjGroupPortMemberPort_KEYS proto.InternalMessageInfo

func (m *ObjGroupPortMemberPort_KEYS) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

func (m *ObjGroupPortMemberPort_KEYS) GetOperatorType() string {
	if m != nil {
		return m.OperatorType
	}
	return ""
}

func (m *ObjGroupPortMemberPort_KEYS) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type ObjGroupPortMemberPort struct {
	OperatorTypeXr       uint32   `protobuf:"varint,50,opt,name=operator_type_xr,json=operatorTypeXr,proto3" json:"operator_type_xr,omitempty"`
	PortXr               uint32   `protobuf:"varint,51,opt,name=port_xr,json=portXr,proto3" json:"port_xr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjGroupPortMemberPort) Reset()         { *m = ObjGroupPortMemberPort{} }
func (m *ObjGroupPortMemberPort) String() string { return proto.CompactTextString(m) }
func (*ObjGroupPortMemberPort) ProtoMessage()    {}
func (*ObjGroupPortMemberPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fe91f737dc2c241, []int{1}
}

func (m *ObjGroupPortMemberPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjGroupPortMemberPort.Unmarshal(m, b)
}
func (m *ObjGroupPortMemberPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjGroupPortMemberPort.Marshal(b, m, deterministic)
}
func (m *ObjGroupPortMemberPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjGroupPortMemberPort.Merge(m, src)
}
func (m *ObjGroupPortMemberPort) XXX_Size() int {
	return xxx_messageInfo_ObjGroupPortMemberPort.Size(m)
}
func (m *ObjGroupPortMemberPort) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjGroupPortMemberPort.DiscardUnknown(m)
}

var xxx_messageInfo_ObjGroupPortMemberPort proto.InternalMessageInfo

func (m *ObjGroupPortMemberPort) GetOperatorTypeXr() uint32 {
	if m != nil {
		return m.OperatorTypeXr
	}
	return 0
}

func (m *ObjGroupPortMemberPort) GetPortXr() uint32 {
	if m != nil {
		return m.PortXr
	}
	return 0
}

func init() {
	proto.RegisterType((*ObjGroupPortMemberPort_KEYS)(nil), "cisco_ios_xr_infra_objmgr_oper.object_group.port.objects.object.operators.operator.obj_group_port_member_port_KEYS")
	proto.RegisterType((*ObjGroupPortMemberPort)(nil), "cisco_ios_xr_infra_objmgr_oper.object_group.port.objects.object.operators.operator.obj_group_port_member_port")
}

func init() { proto.RegisterFile("obj_group_port_member_port.proto", fileDescriptor_0fe91f737dc2c241) }

var fileDescriptor_0fe91f737dc2c241 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xa9, 0xca, 0x8a, 0xa3, 0x2b, 0x92, 0x8b, 0xc1, 0xcb, 0x96, 0xf5, 0xd2, 0x53, 0x0f,
	0xee, 0x6f, 0xf0, 0x24, 0x78, 0xa8, 0x1e, 0xd6, 0xd3, 0x90, 0x94, 0x71, 0x69, 0x21, 0x3b, 0x61,
	0x12, 0xa1, 0x8b, 0x7f, 0x5e, 0x3a, 0xdd, 0x15, 0x3d, 0xf4, 0x94, 0x97, 0x97, 0xf7, 0xbe, 0x07,
	0x81, 0x92, 0x7d, 0x8f, 0x3b, 0xe1, 0xaf, 0x88, 0x91, 0x25, 0x63, 0xa0, 0xe0, 0x49, 0x54, 0xd7,
	0x51, 0x38, 0xb3, 0x69, 0xda, 0x2e, 0xb5, 0x8c, 0x1d, 0x27, 0x1c, 0x04, 0xbb, 0xfd, 0xa7, 0x38,
	0x64, 0xdf, 0x87, 0x9d, 0x20, 0x47, 0x92, 0x9a, 0x7d, 0x4f, 0x6d, 0x9e, 0x18, 0xb5, 0xf6, 0x26,
	0x27, 0x1d, 0xcf, 0x7a, 0x4c, 0xb9, 0xcc, 0x92, 0x7e, 0xd5, 0xfa, 0x1b, 0x56, 0xf3, 0xbb, 0xf8,
	0xf2, 0xfc, 0xf1, 0x66, 0x56, 0x70, 0x7d, 0x24, 0xef, 0x5d, 0x20, 0x5b, 0x94, 0x45, 0x75, 0xd5,
	0xc0, 0x64, 0xbd, 0xba, 0x40, 0xe6, 0x11, 0x96, 0x27, 0x1e, 0xe6, 0x43, 0x24, 0x7b, 0xa6, 0x91,
	0x9b, 0x93, 0xf9, 0x7e, 0x88, 0x64, 0x0c, 0x5c, 0x8c, 0x48, 0x7b, 0xae, 0x6f, 0xaa, 0xd7, 0x08,
	0x0f, 0xf3, 0xe3, 0xa6, 0x82, 0xbb, 0x7f, 0x58, 0x1c, 0xc4, 0x3e, 0x95, 0x45, 0xb5, 0x6c, 0x6e,
	0xff, 0x92, 0xb7, 0x62, 0xee, 0xe1, 0x52, 0xdb, 0x83, 0xd8, 0x8d, 0x06, 0x16, 0xe3, 0x75, 0x2b,
	0x7e, 0xa1, 0x1f, 0xb7, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x38, 0x88, 0x2c, 0xe8, 0x5c, 0x01,
	0x00, 0x00,
}
