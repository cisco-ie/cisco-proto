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
// source: rsi_srlg_rsip_show.proto

package cisco_ios_xr_infra_rsi_oper_srlg_rsips_rsip

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

type RsiSrlgRsipShow_KEYS struct {
	RsipName             string   `protobuf:"bytes,1,opt,name=rsip_name,json=rsipName,proto3" json:"rsip_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsiSrlgRsipShow_KEYS) Reset()         { *m = RsiSrlgRsipShow_KEYS{} }
func (m *RsiSrlgRsipShow_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsiSrlgRsipShow_KEYS) ProtoMessage()    {}
func (*RsiSrlgRsipShow_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c8fc18844c8d308, []int{0}
}

func (m *RsiSrlgRsipShow_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiSrlgRsipShow_KEYS.Unmarshal(m, b)
}
func (m *RsiSrlgRsipShow_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiSrlgRsipShow_KEYS.Marshal(b, m, deterministic)
}
func (m *RsiSrlgRsipShow_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiSrlgRsipShow_KEYS.Merge(m, src)
}
func (m *RsiSrlgRsipShow_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsiSrlgRsipShow_KEYS.Size(m)
}
func (m *RsiSrlgRsipShow_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiSrlgRsipShow_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsiSrlgRsipShow_KEYS proto.InternalMessageInfo

func (m *RsiSrlgRsipShow_KEYS) GetRsipName() string {
	if m != nil {
		return m.RsipName
	}
	return ""
}

type RsiSrlgAttrBrief struct {
	SrlgValue            uint32   `protobuf:"varint,1,opt,name=srlg_value,json=srlgValue,proto3" json:"srlg_value,omitempty"`
	Priority             string   `protobuf:"bytes,2,opt,name=priority,proto3" json:"priority,omitempty"`
	SrlgIndex            uint32   `protobuf:"varint,3,opt,name=srlg_index,json=srlgIndex,proto3" json:"srlg_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsiSrlgAttrBrief) Reset()         { *m = RsiSrlgAttrBrief{} }
func (m *RsiSrlgAttrBrief) String() string { return proto.CompactTextString(m) }
func (*RsiSrlgAttrBrief) ProtoMessage()    {}
func (*RsiSrlgAttrBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c8fc18844c8d308, []int{1}
}

func (m *RsiSrlgAttrBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiSrlgAttrBrief.Unmarshal(m, b)
}
func (m *RsiSrlgAttrBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiSrlgAttrBrief.Marshal(b, m, deterministic)
}
func (m *RsiSrlgAttrBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiSrlgAttrBrief.Merge(m, src)
}
func (m *RsiSrlgAttrBrief) XXX_Size() int {
	return xxx_messageInfo_RsiSrlgAttrBrief.Size(m)
}
func (m *RsiSrlgAttrBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiSrlgAttrBrief.DiscardUnknown(m)
}

var xxx_messageInfo_RsiSrlgAttrBrief proto.InternalMessageInfo

func (m *RsiSrlgAttrBrief) GetSrlgValue() uint32 {
	if m != nil {
		return m.SrlgValue
	}
	return 0
}

func (m *RsiSrlgAttrBrief) GetPriority() string {
	if m != nil {
		return m.Priority
	}
	return ""
}

func (m *RsiSrlgAttrBrief) GetSrlgIndex() uint32 {
	if m != nil {
		return m.SrlgIndex
	}
	return 0
}

type RsiSrlgRsipShow struct {
	OpticalLayerInterfaceName string              `protobuf:"bytes,50,opt,name=optical_layer_interface_name,json=opticalLayerInterfaceName,proto3" json:"optical_layer_interface_name,omitempty"`
	Registrations             uint32              `protobuf:"varint,51,opt,name=registrations,proto3" json:"registrations,omitempty"`
	InterfaceValues           uint32              `protobuf:"varint,52,opt,name=interface_values,json=interfaceValues,proto3" json:"interface_values,omitempty"`
	SrlgAttribute             []*RsiSrlgAttrBrief `protobuf:"bytes,53,rep,name=srlg_attribute,json=srlgAttribute,proto3" json:"srlg_attribute,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}            `json:"-"`
	XXX_unrecognized          []byte              `json:"-"`
	XXX_sizecache             int32               `json:"-"`
}

func (m *RsiSrlgRsipShow) Reset()         { *m = RsiSrlgRsipShow{} }
func (m *RsiSrlgRsipShow) String() string { return proto.CompactTextString(m) }
func (*RsiSrlgRsipShow) ProtoMessage()    {}
func (*RsiSrlgRsipShow) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c8fc18844c8d308, []int{2}
}

func (m *RsiSrlgRsipShow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiSrlgRsipShow.Unmarshal(m, b)
}
func (m *RsiSrlgRsipShow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiSrlgRsipShow.Marshal(b, m, deterministic)
}
func (m *RsiSrlgRsipShow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiSrlgRsipShow.Merge(m, src)
}
func (m *RsiSrlgRsipShow) XXX_Size() int {
	return xxx_messageInfo_RsiSrlgRsipShow.Size(m)
}
func (m *RsiSrlgRsipShow) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiSrlgRsipShow.DiscardUnknown(m)
}

var xxx_messageInfo_RsiSrlgRsipShow proto.InternalMessageInfo

func (m *RsiSrlgRsipShow) GetOpticalLayerInterfaceName() string {
	if m != nil {
		return m.OpticalLayerInterfaceName
	}
	return ""
}

func (m *RsiSrlgRsipShow) GetRegistrations() uint32 {
	if m != nil {
		return m.Registrations
	}
	return 0
}

func (m *RsiSrlgRsipShow) GetInterfaceValues() uint32 {
	if m != nil {
		return m.InterfaceValues
	}
	return 0
}

func (m *RsiSrlgRsipShow) GetSrlgAttribute() []*RsiSrlgAttrBrief {
	if m != nil {
		return m.SrlgAttribute
	}
	return nil
}

func init() {
	proto.RegisterType((*RsiSrlgRsipShow_KEYS)(nil), "cisco_ios_xr_infra_rsi_oper.srlg.rsips.rsip.rsi_srlg_rsip_show_KEYS")
	proto.RegisterType((*RsiSrlgAttrBrief)(nil), "cisco_ios_xr_infra_rsi_oper.srlg.rsips.rsip.rsi_srlg_attr_brief")
	proto.RegisterType((*RsiSrlgRsipShow)(nil), "cisco_ios_xr_infra_rsi_oper.srlg.rsips.rsip.rsi_srlg_rsip_show")
}

func init() { proto.RegisterFile("rsi_srlg_rsip_show.proto", fileDescriptor_2c8fc18844c8d308) }

var fileDescriptor_2c8fc18844c8d308 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x69, 0x0b, 0xd2, 0x8c, 0x44, 0x65, 0x3d, 0xb8, 0xfe, 0x83, 0x12, 0x3c, 0x54, 0x84,
	0x1c, 0x5a, 0xf5, 0xaa, 0x1e, 0x3c, 0x14, 0xc5, 0x43, 0x04, 0xc1, 0xd3, 0xb2, 0x89, 0x9b, 0x38,
	0x90, 0x66, 0xc3, 0xec, 0x56, 0xdb, 0x4f, 0xe1, 0x57, 0x96, 0xdd, 0x98, 0x14, 0xa9, 0x17, 0x2f,
	0x81, 0xfc, 0x78, 0xef, 0xed, 0xbc, 0x19, 0xe0, 0x64, 0x50, 0x18, 0x2a, 0x0b, 0x41, 0x06, 0x6b,
	0x61, 0xde, 0xf5, 0x67, 0x5c, 0x93, 0xb6, 0x9a, 0x5d, 0x64, 0x68, 0x32, 0x2d, 0x50, 0x1b, 0xb1,
	0x24, 0x81, 0x55, 0x4e, 0xd2, 0x69, 0x84, 0xae, 0x15, 0xc5, 0xce, 0x11, 0x3b, 0x87, 0xf1, 0xdf,
	0xe8, 0x1a, 0x0e, 0x36, 0x83, 0xc4, 0xc3, 0xfd, 0xeb, 0x33, 0x3b, 0x86, 0xc0, 0x93, 0x4a, 0xce,
	0x15, 0xef, 0x8d, 0x7a, 0xe3, 0x20, 0x19, 0x3a, 0xf0, 0x24, 0xe7, 0x2a, 0xd2, 0xb0, 0xdf, 0xf9,
	0xa4, 0xb5, 0x24, 0x52, 0x42, 0x95, 0xb3, 0x53, 0x00, 0x8f, 0x3e, 0x64, 0xb9, 0x68, 0x4c, 0x61,
	0x12, 0x38, 0xf2, 0xe2, 0x00, 0x3b, 0x82, 0x61, 0x4d, 0xa8, 0x09, 0xed, 0x8a, 0xf7, 0x9b, 0xc4,
	0xf6, 0xbf, 0xb3, 0x62, 0xf5, 0xa6, 0x96, 0x7c, 0xb0, 0xb6, 0xce, 0x1c, 0x88, 0xbe, 0xfa, 0xc0,
	0x36, 0x27, 0x65, 0x37, 0x70, 0xa2, 0x6b, 0x8b, 0x99, 0x2c, 0x45, 0x29, 0x57, 0xca, 0xf5, 0xb5,
	0x8a, 0x72, 0x99, 0xa9, 0x66, 0xee, 0x89, 0x7f, 0xe5, 0xf0, 0x47, 0xf3, 0xe8, 0x24, 0xb3, 0x56,
	0xe1, 0x8a, 0xb0, 0x33, 0x08, 0x49, 0x15, 0x68, 0x2c, 0x49, 0x8b, 0xba, 0x32, 0x7c, 0xea, 0x5f,
	0xfe, 0x0d, 0xd9, 0x39, 0xec, 0xad, 0x83, 0x7d, 0x39, 0xc3, 0x2f, 0xbd, 0x70, 0xb7, 0xe3, 0xbe,
	0xa2, 0x61, 0x05, 0xec, 0x74, 0x5b, 0xc1, 0x74, 0x61, 0x15, 0xbf, 0x1a, 0x0d, 0xc6, 0xdb, 0x93,
	0xdb, 0xf8, 0x1f, 0x77, 0x89, 0xff, 0x58, 0x6e, 0x12, 0x3a, 0x70, 0xd7, 0xc6, 0xa6, 0x5b, 0xfe,
	0xdc, 0xd3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x86, 0x15, 0x2d, 0x0a, 0x02, 0x00, 0x00,
}