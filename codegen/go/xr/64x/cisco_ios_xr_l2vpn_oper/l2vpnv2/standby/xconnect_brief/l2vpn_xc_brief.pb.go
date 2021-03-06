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
// source: l2vpn_xc_brief.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_standby_xconnect_brief

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

type L2VpnXcBrief_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnXcBrief_KEYS) Reset()         { *m = L2VpnXcBrief_KEYS{} }
func (m *L2VpnXcBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnXcBrief_KEYS) ProtoMessage()    {}
func (*L2VpnXcBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d19351c63b40399, []int{0}
}

func (m *L2VpnXcBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnXcBrief_KEYS.Unmarshal(m, b)
}
func (m *L2VpnXcBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnXcBrief_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnXcBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnXcBrief_KEYS.Merge(m, src)
}
func (m *L2VpnXcBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnXcBrief_KEYS.Size(m)
}
func (m *L2VpnXcBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnXcBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnXcBrief_KEYS proto.InternalMessageInfo

type L2VpnXcBriefLine struct {
	PsnType              string   `protobuf:"bytes,1,opt,name=psn_type,json=psnType,proto3" json:"psn_type,omitempty"`
	Ac1Encapsulation     string   `protobuf:"bytes,2,opt,name=ac1_encapsulation,json=ac1Encapsulation,proto3" json:"ac1_encapsulation,omitempty"`
	Ac2Encapsulation     string   `protobuf:"bytes,3,opt,name=ac2_encapsulation,json=ac2Encapsulation,proto3" json:"ac2_encapsulation,omitempty"`
	UpCount              []uint32 `protobuf:"varint,4,rep,packed,name=up_count,json=upCount,proto3" json:"up_count,omitempty"`
	DownCount            []uint32 `protobuf:"varint,5,rep,packed,name=down_count,json=downCount,proto3" json:"down_count,omitempty"`
	UnresolvedCount      []uint32 `protobuf:"varint,6,rep,packed,name=unresolved_count,json=unresolvedCount,proto3" json:"unresolved_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnXcBriefLine) Reset()         { *m = L2VpnXcBriefLine{} }
func (m *L2VpnXcBriefLine) String() string { return proto.CompactTextString(m) }
func (*L2VpnXcBriefLine) ProtoMessage()    {}
func (*L2VpnXcBriefLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d19351c63b40399, []int{1}
}

func (m *L2VpnXcBriefLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnXcBriefLine.Unmarshal(m, b)
}
func (m *L2VpnXcBriefLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnXcBriefLine.Marshal(b, m, deterministic)
}
func (m *L2VpnXcBriefLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnXcBriefLine.Merge(m, src)
}
func (m *L2VpnXcBriefLine) XXX_Size() int {
	return xxx_messageInfo_L2VpnXcBriefLine.Size(m)
}
func (m *L2VpnXcBriefLine) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnXcBriefLine.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnXcBriefLine proto.InternalMessageInfo

func (m *L2VpnXcBriefLine) GetPsnType() string {
	if m != nil {
		return m.PsnType
	}
	return ""
}

func (m *L2VpnXcBriefLine) GetAc1Encapsulation() string {
	if m != nil {
		return m.Ac1Encapsulation
	}
	return ""
}

func (m *L2VpnXcBriefLine) GetAc2Encapsulation() string {
	if m != nil {
		return m.Ac2Encapsulation
	}
	return ""
}

func (m *L2VpnXcBriefLine) GetUpCount() []uint32 {
	if m != nil {
		return m.UpCount
	}
	return nil
}

func (m *L2VpnXcBriefLine) GetDownCount() []uint32 {
	if m != nil {
		return m.DownCount
	}
	return nil
}

func (m *L2VpnXcBriefLine) GetUnresolvedCount() []uint32 {
	if m != nil {
		return m.UnresolvedCount
	}
	return nil
}

type L2VpnXcBriefFirstLayer struct {
	Ac2                  []*L2VpnXcBriefLine `protobuf:"bytes,1,rep,name=ac2,proto3" json:"ac2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *L2VpnXcBriefFirstLayer) Reset()         { *m = L2VpnXcBriefFirstLayer{} }
func (m *L2VpnXcBriefFirstLayer) String() string { return proto.CompactTextString(m) }
func (*L2VpnXcBriefFirstLayer) ProtoMessage()    {}
func (*L2VpnXcBriefFirstLayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d19351c63b40399, []int{2}
}

func (m *L2VpnXcBriefFirstLayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnXcBriefFirstLayer.Unmarshal(m, b)
}
func (m *L2VpnXcBriefFirstLayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnXcBriefFirstLayer.Marshal(b, m, deterministic)
}
func (m *L2VpnXcBriefFirstLayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnXcBriefFirstLayer.Merge(m, src)
}
func (m *L2VpnXcBriefFirstLayer) XXX_Size() int {
	return xxx_messageInfo_L2VpnXcBriefFirstLayer.Size(m)
}
func (m *L2VpnXcBriefFirstLayer) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnXcBriefFirstLayer.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnXcBriefFirstLayer proto.InternalMessageInfo

func (m *L2VpnXcBriefFirstLayer) GetAc2() []*L2VpnXcBriefLine {
	if m != nil {
		return m.Ac2
	}
	return nil
}

type L2VpnXcBriefMatrix struct {
	Ac1                  []*L2VpnXcBriefFirstLayer `protobuf:"bytes,1,rep,name=ac1,proto3" json:"ac1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *L2VpnXcBriefMatrix) Reset()         { *m = L2VpnXcBriefMatrix{} }
func (m *L2VpnXcBriefMatrix) String() string { return proto.CompactTextString(m) }
func (*L2VpnXcBriefMatrix) ProtoMessage()    {}
func (*L2VpnXcBriefMatrix) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d19351c63b40399, []int{3}
}

func (m *L2VpnXcBriefMatrix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnXcBriefMatrix.Unmarshal(m, b)
}
func (m *L2VpnXcBriefMatrix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnXcBriefMatrix.Marshal(b, m, deterministic)
}
func (m *L2VpnXcBriefMatrix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnXcBriefMatrix.Merge(m, src)
}
func (m *L2VpnXcBriefMatrix) XXX_Size() int {
	return xxx_messageInfo_L2VpnXcBriefMatrix.Size(m)
}
func (m *L2VpnXcBriefMatrix) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnXcBriefMatrix.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnXcBriefMatrix proto.InternalMessageInfo

func (m *L2VpnXcBriefMatrix) GetAc1() []*L2VpnXcBriefFirstLayer {
	if m != nil {
		return m.Ac1
	}
	return nil
}

type L2VpnXcBrief struct {
	EncapsulationReportMatrix []*L2VpnXcBriefMatrix `protobuf:"bytes,50,rep,name=encapsulation_report_matrix,json=encapsulationReportMatrix,proto3" json:"encapsulation_report_matrix,omitempty"`
	EncapsulationTotal        []*L2VpnXcBriefLine   `protobuf:"bytes,51,rep,name=encapsulation_total,json=encapsulationTotal,proto3" json:"encapsulation_total,omitempty"`
	MainTotalUp               uint32                `protobuf:"varint,52,opt,name=main_total_up,json=mainTotalUp,proto3" json:"main_total_up,omitempty"`
	MainTotalDown             uint32                `protobuf:"varint,53,opt,name=main_total_down,json=mainTotalDown,proto3" json:"main_total_down,omitempty"`
	MainTotalUnresolved       uint32                `protobuf:"varint,54,opt,name=main_total_unresolved,json=mainTotalUnresolved,proto3" json:"main_total_unresolved,omitempty"`
	UndefinedXc               uint32                `protobuf:"varint,55,opt,name=undefined_xc,json=undefinedXc,proto3" json:"undefined_xc,omitempty"`
	MemoryState               string                `protobuf:"bytes,56,opt,name=memory_state,json=memoryState,proto3" json:"memory_state,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}              `json:"-"`
	XXX_unrecognized          []byte                `json:"-"`
	XXX_sizecache             int32                 `json:"-"`
}

func (m *L2VpnXcBrief) Reset()         { *m = L2VpnXcBrief{} }
func (m *L2VpnXcBrief) String() string { return proto.CompactTextString(m) }
func (*L2VpnXcBrief) ProtoMessage()    {}
func (*L2VpnXcBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d19351c63b40399, []int{4}
}

func (m *L2VpnXcBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnXcBrief.Unmarshal(m, b)
}
func (m *L2VpnXcBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnXcBrief.Marshal(b, m, deterministic)
}
func (m *L2VpnXcBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnXcBrief.Merge(m, src)
}
func (m *L2VpnXcBrief) XXX_Size() int {
	return xxx_messageInfo_L2VpnXcBrief.Size(m)
}
func (m *L2VpnXcBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnXcBrief.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnXcBrief proto.InternalMessageInfo

func (m *L2VpnXcBrief) GetEncapsulationReportMatrix() []*L2VpnXcBriefMatrix {
	if m != nil {
		return m.EncapsulationReportMatrix
	}
	return nil
}

func (m *L2VpnXcBrief) GetEncapsulationTotal() []*L2VpnXcBriefLine {
	if m != nil {
		return m.EncapsulationTotal
	}
	return nil
}

func (m *L2VpnXcBrief) GetMainTotalUp() uint32 {
	if m != nil {
		return m.MainTotalUp
	}
	return 0
}

func (m *L2VpnXcBrief) GetMainTotalDown() uint32 {
	if m != nil {
		return m.MainTotalDown
	}
	return 0
}

func (m *L2VpnXcBrief) GetMainTotalUnresolved() uint32 {
	if m != nil {
		return m.MainTotalUnresolved
	}
	return 0
}

func (m *L2VpnXcBrief) GetUndefinedXc() uint32 {
	if m != nil {
		return m.UndefinedXc
	}
	return 0
}

func (m *L2VpnXcBrief) GetMemoryState() string {
	if m != nil {
		return m.MemoryState
	}
	return ""
}

func init() {
	proto.RegisterType((*L2VpnXcBrief_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_brief.l2vpn_xc_brief_KEYS")
	proto.RegisterType((*L2VpnXcBriefLine)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_brief.l2vpn_xc_brief_line")
	proto.RegisterType((*L2VpnXcBriefFirstLayer)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_brief.l2vpn_xc_brief_first_layer")
	proto.RegisterType((*L2VpnXcBriefMatrix)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_brief.l2vpn_xc_brief_matrix")
	proto.RegisterType((*L2VpnXcBrief)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_brief.l2vpn_xc_brief")
}

func init() { proto.RegisterFile("l2vpn_xc_brief.proto", fileDescriptor_8d19351c63b40399) }

var fileDescriptor_8d19351c63b40399 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x8e, 0xd3, 0x3c,
	0x14, 0xc5, 0x95, 0x2f, 0x1f, 0xf3, 0xe7, 0x96, 0x32, 0x83, 0x4b, 0x25, 0x17, 0x84, 0x54, 0xb2,
	0x40, 0x45, 0x48, 0x91, 0x9a, 0x81, 0x81, 0x3d, 0xcc, 0x6a, 0x34, 0x9b, 0xcc, 0x20, 0xc1, 0x02,
	0x59, 0xae, 0xe3, 0x4a, 0x96, 0x52, 0xdb, 0xb2, 0x9d, 0x4e, 0x22, 0xe0, 0x11, 0x78, 0x4b, 0xf6,
	0xbc, 0x02, 0xb2, 0x13, 0x3a, 0x4d, 0x61, 0x85, 0xba, 0x4b, 0xce, 0xfd, 0xe5, 0x9c, 0xeb, 0x6b,
	0xc7, 0xf0, 0xa8, 0xcc, 0xd6, 0x5a, 0x92, 0x9a, 0x91, 0x85, 0x11, 0x7c, 0x99, 0x6a, 0xa3, 0x9c,
	0x42, 0xe7, 0x4c, 0x58, 0xa6, 0x88, 0x50, 0x96, 0xd4, 0x86, 0xb4, 0x88, 0xd2, 0xdc, 0xa4, 0xe1,
	0x71, 0x9d, 0xa5, 0xd6, 0x51, 0x59, 0x2c, 0x9a, 0xb4, 0x66, 0x4a, 0x4a, 0xce, 0x5c, 0xfb, 0x75,
	0x32, 0x86, 0x51, 0xdf, 0x8f, 0x5c, 0x5e, 0x7c, 0xba, 0x4e, 0x7e, 0x46, 0x7f, 0xe8, 0xa5, 0x90,
	0x1c, 0x4d, 0xe0, 0x48, 0x5b, 0x49, 0x5c, 0xa3, 0x39, 0x8e, 0xa6, 0xd1, 0xec, 0x38, 0x3f, 0xd4,
	0x56, 0xde, 0x34, 0x9a, 0xa3, 0x97, 0xf0, 0x90, 0xb2, 0x39, 0xe1, 0x92, 0x51, 0x6d, 0xab, 0x92,
	0x3a, 0xa1, 0x24, 0xfe, 0x2f, 0x30, 0xa7, 0x94, 0xcd, 0x2f, 0xb6, 0xf5, 0x16, 0xce, 0x76, 0xe0,
	0xf8, 0x37, 0x9c, 0xf5, 0xe1, 0x09, 0x1c, 0x55, 0x9a, 0x30, 0x55, 0x49, 0x87, 0xff, 0x9f, 0xc6,
	0xb3, 0x61, 0x7e, 0x58, 0xe9, 0x77, 0xfe, 0x15, 0x3d, 0x05, 0x28, 0xd4, 0xad, 0xec, 0x8a, 0xf7,
	0x42, 0xf1, 0xd8, 0x2b, 0x6d, 0xf9, 0x05, 0x9c, 0x56, 0xd2, 0x70, 0xab, 0xca, 0x35, 0x2f, 0x3a,
	0xe8, 0x20, 0x40, 0x27, 0x77, 0x7a, 0x40, 0x93, 0x2f, 0xf0, 0x78, 0x67, 0xc1, 0x4b, 0x61, 0xac,
	0x23, 0x25, 0x6d, 0xb8, 0x41, 0x9f, 0x21, 0xa6, 0x2c, 0xc3, 0xd1, 0x34, 0x9e, 0x0d, 0xb2, 0xcb,
	0xf4, 0xdf, 0x86, 0x9d, 0xfe, 0x65, 0xa2, 0xb9, 0xf7, 0x4d, 0xbe, 0xc1, 0x78, 0xa7, 0xb6, 0xa2,
	0xce, 0x88, 0x1a, 0x15, 0x3e, 0x77, 0xde, 0xe5, 0xe6, 0x7b, 0xca, 0xdd, 0x5a, 0x98, 0x8f, 0x9f,
	0x27, 0x3f, 0x62, 0x78, 0xd0, 0x67, 0xd0, 0xf7, 0x08, 0x9e, 0xf4, 0x76, 0x87, 0x18, 0xae, 0x95,
	0x71, 0x5d, 0x63, 0x38, 0x0b, 0x1d, 0x5d, 0xed, 0xa9, 0xa3, 0xd6, 0x34, 0x9f, 0xf4, 0x12, 0xf3,
	0x10, 0x78, 0xd5, 0x0e, 0xe2, 0x2b, 0x8c, 0xfa, 0xed, 0x38, 0xe5, 0x68, 0x89, 0xcf, 0xf6, 0xbf,
	0x21, 0xa8, 0x97, 0x73, 0xe3, 0x63, 0x50, 0x02, 0xc3, 0x15, 0x15, 0x5d, 0x28, 0xa9, 0x34, 0x7e,
	0x35, 0x8d, 0x66, 0xc3, 0x7c, 0xe0, 0xc5, 0x40, 0x7c, 0xd0, 0xe8, 0x39, 0x9c, 0x6c, 0x31, 0xfe,
	0x0c, 0xe2, 0xd7, 0x81, 0x1a, 0x6e, 0xa8, 0xf7, 0xea, 0x56, 0xa2, 0x0c, 0xc6, 0xdb, 0x5e, 0x9b,
	0x63, 0x88, 0xcf, 0x03, 0x3d, 0xba, 0xf3, 0xdc, 0x94, 0xd0, 0x33, 0xb8, 0x5f, 0xc9, 0x82, 0x2f,
	0x85, 0xe4, 0x05, 0xa9, 0x19, 0x7e, 0xd3, 0xc6, 0x6f, 0xb4, 0x8f, 0xcc, 0x23, 0x2b, 0xbe, 0x52,
	0xa6, 0x21, 0xd6, 0x51, 0xc7, 0xf1, 0xdb, 0xf0, 0x33, 0x0d, 0x5a, 0xed, 0xda, 0x4b, 0x8b, 0x83,
	0x70, 0x55, 0x9c, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x54, 0xf5, 0xa9, 0x6e, 0x42, 0x04, 0x00,
	0x00,
}
