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
// source: current_install_operation.proto

package cisco_ios_xr_spirit_install_instmgr_oper_software_install_request

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

type CurrentInstallOperation_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentInstallOperation_KEYS) Reset()         { *m = CurrentInstallOperation_KEYS{} }
func (m *CurrentInstallOperation_KEYS) String() string { return proto.CompactTextString(m) }
func (*CurrentInstallOperation_KEYS) ProtoMessage()    {}
func (*CurrentInstallOperation_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25440856242766a, []int{0}
}

func (m *CurrentInstallOperation_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentInstallOperation_KEYS.Unmarshal(m, b)
}
func (m *CurrentInstallOperation_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentInstallOperation_KEYS.Marshal(b, m, deterministic)
}
func (m *CurrentInstallOperation_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentInstallOperation_KEYS.Merge(m, src)
}
func (m *CurrentInstallOperation_KEYS) XXX_Size() int {
	return xxx_messageInfo_CurrentInstallOperation_KEYS.Size(m)
}
func (m *CurrentInstallOperation_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentInstallOperation_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentInstallOperation_KEYS proto.InternalMessageInfo

type CurrentInstallOperation struct {
	CurrInstOper         string   `protobuf:"bytes,50,opt,name=curr_inst_oper,json=currInstOper,proto3" json:"curr_inst_oper,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentInstallOperation) Reset()         { *m = CurrentInstallOperation{} }
func (m *CurrentInstallOperation) String() string { return proto.CompactTextString(m) }
func (*CurrentInstallOperation) ProtoMessage()    {}
func (*CurrentInstallOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25440856242766a, []int{1}
}

func (m *CurrentInstallOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentInstallOperation.Unmarshal(m, b)
}
func (m *CurrentInstallOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentInstallOperation.Marshal(b, m, deterministic)
}
func (m *CurrentInstallOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentInstallOperation.Merge(m, src)
}
func (m *CurrentInstallOperation) XXX_Size() int {
	return xxx_messageInfo_CurrentInstallOperation.Size(m)
}
func (m *CurrentInstallOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentInstallOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentInstallOperation proto.InternalMessageInfo

func (m *CurrentInstallOperation) GetCurrInstOper() string {
	if m != nil {
		return m.CurrInstOper
	}
	return ""
}

func init() {
	proto.RegisterType((*CurrentInstallOperation_KEYS)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.request.current_install_operation_KEYS")
	proto.RegisterType((*CurrentInstallOperation)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.request.current_install_operation")
}

func init() { proto.RegisterFile("current_install_operation.proto", fileDescriptor_d25440856242766a) }

var fileDescriptor_d25440856242766a = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2e, 0x2d, 0x2a,
	0x4a, 0xcd, 0x2b, 0x89, 0xcf, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x89, 0xcf, 0x2f, 0x48, 0x2d,
	0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x72, 0x4c, 0xce, 0x2c,
	0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x2f, 0x2e, 0xc8, 0x2c, 0xca, 0x44,
	0x28, 0x06, 0xd1, 0xb9, 0xe9, 0x45, 0x60, 0x4d, 0x7a, 0xc5, 0xf9, 0x69, 0x25, 0xe5, 0x89, 0x45,
	0xa9, 0x30, 0x59, 0xbd, 0xa2, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x05, 0x2e, 0x39, 0x9c,
	0xb6, 0xc4, 0x7b, 0xbb, 0x46, 0x06, 0x2b, 0x39, 0x72, 0x49, 0xe2, 0x54, 0x21, 0xa4, 0xc2, 0xc5,
	0x07, 0x92, 0x04, 0xcb, 0x80, 0x85, 0x25, 0x8c, 0x14, 0x18, 0x35, 0x38, 0x83, 0x78, 0x40, 0xa2,
	0x9e, 0x79, 0xc5, 0x25, 0xfe, 0x05, 0xa9, 0x45, 0x49, 0x6c, 0x60, 0xe7, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x88, 0xff, 0x5d, 0x35, 0xd1, 0x00, 0x00, 0x00,
}