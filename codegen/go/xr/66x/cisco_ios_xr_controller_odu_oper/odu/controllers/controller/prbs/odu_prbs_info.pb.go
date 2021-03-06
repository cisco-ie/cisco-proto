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
// source: odu_prbs_info.proto

package cisco_ios_xr_controller_odu_oper_odu_controllers_controller_prbs

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

type OduPrbsInfo_KEYS struct {
	ControllerName       string   `protobuf:"bytes,1,opt,name=controller_name,json=controllerName,proto3" json:"controller_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OduPrbsInfo_KEYS) Reset()         { *m = OduPrbsInfo_KEYS{} }
func (m *OduPrbsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*OduPrbsInfo_KEYS) ProtoMessage()    {}
func (*OduPrbsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_469e08ef878b6dbd, []int{0}
}

func (m *OduPrbsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OduPrbsInfo_KEYS.Unmarshal(m, b)
}
func (m *OduPrbsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OduPrbsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *OduPrbsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OduPrbsInfo_KEYS.Merge(m, src)
}
func (m *OduPrbsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_OduPrbsInfo_KEYS.Size(m)
}
func (m *OduPrbsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OduPrbsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OduPrbsInfo_KEYS proto.InternalMessageInfo

func (m *OduPrbsInfo_KEYS) GetControllerName() string {
	if m != nil {
		return m.ControllerName
	}
	return ""
}

type OduPrbsInfo struct {
	OduPrbsTest          string   `protobuf:"bytes,50,opt,name=odu_prbs_test,json=oduPrbsTest,proto3" json:"odu_prbs_test,omitempty"`
	OduPrbsMode          string   `protobuf:"bytes,51,opt,name=odu_prbs_mode,json=oduPrbsMode,proto3" json:"odu_prbs_mode,omitempty"`
	OduPrbsPattern       string   `protobuf:"bytes,52,opt,name=odu_prbs_pattern,json=oduPrbsPattern,proto3" json:"odu_prbs_pattern,omitempty"`
	OduPrbsStatus        string   `protobuf:"bytes,53,opt,name=odu_prbs_status,json=oduPrbsStatus,proto3" json:"odu_prbs_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OduPrbsInfo) Reset()         { *m = OduPrbsInfo{} }
func (m *OduPrbsInfo) String() string { return proto.CompactTextString(m) }
func (*OduPrbsInfo) ProtoMessage()    {}
func (*OduPrbsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_469e08ef878b6dbd, []int{1}
}

func (m *OduPrbsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OduPrbsInfo.Unmarshal(m, b)
}
func (m *OduPrbsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OduPrbsInfo.Marshal(b, m, deterministic)
}
func (m *OduPrbsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OduPrbsInfo.Merge(m, src)
}
func (m *OduPrbsInfo) XXX_Size() int {
	return xxx_messageInfo_OduPrbsInfo.Size(m)
}
func (m *OduPrbsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OduPrbsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OduPrbsInfo proto.InternalMessageInfo

func (m *OduPrbsInfo) GetOduPrbsTest() string {
	if m != nil {
		return m.OduPrbsTest
	}
	return ""
}

func (m *OduPrbsInfo) GetOduPrbsMode() string {
	if m != nil {
		return m.OduPrbsMode
	}
	return ""
}

func (m *OduPrbsInfo) GetOduPrbsPattern() string {
	if m != nil {
		return m.OduPrbsPattern
	}
	return ""
}

func (m *OduPrbsInfo) GetOduPrbsStatus() string {
	if m != nil {
		return m.OduPrbsStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*OduPrbsInfo_KEYS)(nil), "cisco_ios_xr_controller_odu_oper.odu.controllers.controller.prbs.odu_prbs_info_KEYS")
	proto.RegisterType((*OduPrbsInfo)(nil), "cisco_ios_xr_controller_odu_oper.odu.controllers.controller.prbs.odu_prbs_info")
}

func init() { proto.RegisterFile("odu_prbs_info.proto", fileDescriptor_469e08ef878b6dbd) }

var fileDescriptor_469e08ef878b6dbd = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xd0, 0xb1, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0x71, 0x6e, 0x11, 0x8c, 0xd4, 0x4a, 0x5c, 0x32, 0xca, 0x0d, 0xda, 0x29, 0x83, 0xd5,
	0x51, 0x70, 0x71, 0x12, 0xa5, 0x58, 0x17, 0xa7, 0x90, 0xbb, 0x3c, 0xe1, 0xa0, 0x97, 0x17, 0xde,
	0x7b, 0x01, 0xff, 0x25, 0xff, 0xcb, 0xd2, 0xb4, 0xf4, 0x2e, 0xdb, 0xf1, 0xe5, 0xf3, 0x7e, 0x1c,
	0x51, 0xb7, 0x18, 0xb2, 0x4b, 0xd4, 0xb1, 0x1b, 0xe2, 0x2f, 0xda, 0x44, 0x28, 0xa8, 0x5f, 0xfb,
	0x81, 0x7b, 0x74, 0x03, 0xb2, 0xfb, 0x23, 0xd7, 0x63, 0x14, 0xc2, 0xdd, 0x0e, 0xc8, 0x1d, 0x30,
	0x26, 0x20, 0x8b, 0x21, 0xdb, 0xa9, 0xf3, 0xec, 0xdb, 0x1e, 0xc6, 0xda, 0x17, 0xa5, 0xab, 0x61,
	0xf7, 0xfe, 0xf6, 0xb3, 0xd5, 0x0f, 0x6a, 0x39, 0x1b, 0x8b, 0x7e, 0x04, 0xd3, 0xdc, 0x35, 0xab,
	0xcb, 0xaf, 0xeb, 0x29, 0x7f, 0xfa, 0x11, 0xda, 0xff, 0x46, 0x2d, 0xaa, 0x7b, 0xdd, 0xce, 0x82,
	0x00, 0x8b, 0x79, 0x2c, 0x87, 0x57, 0x18, 0xf2, 0x86, 0x3a, 0xfe, 0x06, 0x96, 0xca, 0x8c, 0x18,
	0xc0, 0xac, 0x2b, 0xf3, 0x81, 0x01, 0xf4, 0x4a, 0xdd, 0x9c, 0x4d, 0xf2, 0x22, 0x40, 0xd1, 0x3c,
	0x1d, 0xff, 0xe1, 0xc4, 0x36, 0xc7, 0xaa, 0xef, 0xd5, 0xf2, 0x2c, 0x59, 0xbc, 0x64, 0x36, 0xcf,
	0x05, 0x2e, 0x4e, 0x70, 0x5b, 0x62, 0x77, 0x51, 0xde, 0x6c, 0xbd, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0xb8, 0x3f, 0x5e, 0x21, 0x4a, 0x01, 0x00, 0x00,
}
