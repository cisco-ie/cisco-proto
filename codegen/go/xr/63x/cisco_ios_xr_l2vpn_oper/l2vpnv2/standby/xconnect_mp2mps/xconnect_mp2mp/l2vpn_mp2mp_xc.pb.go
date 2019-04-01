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
// source: l2vpn_mp2mp_xc.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_standby_xconnect_mp2mps_xconnect_mp2mp

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

type L2VpnMp2MpXc_KEYS struct {
	GroupName            string   `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Mp2MpName            string   `protobuf:"bytes,2,opt,name=mp2_mp_name,json=mp2MpName,proto3" json:"mp2_mp_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMp2MpXc_KEYS) Reset()         { *m = L2VpnMp2MpXc_KEYS{} }
func (m *L2VpnMp2MpXc_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnMp2MpXc_KEYS) ProtoMessage()    {}
func (*L2VpnMp2MpXc_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{0}
}

func (m *L2VpnMp2MpXc_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMp2MpXc_KEYS.Unmarshal(m, b)
}
func (m *L2VpnMp2MpXc_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMp2MpXc_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnMp2MpXc_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMp2MpXc_KEYS.Merge(m, src)
}
func (m *L2VpnMp2MpXc_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnMp2MpXc_KEYS.Size(m)
}
func (m *L2VpnMp2MpXc_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMp2MpXc_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMp2MpXc_KEYS proto.InternalMessageInfo

func (m *L2VpnMp2MpXc_KEYS) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *L2VpnMp2MpXc_KEYS) GetMp2MpName() string {
	if m != nil {
		return m.Mp2MpName
	}
	return ""
}

type L2VpnRdAuto struct {
	RouterId             string   `protobuf:"bytes,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	AutoIndex            uint32   `protobuf:"varint,2,opt,name=auto_index,json=autoIndex,proto3" json:"auto_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRdAuto) Reset()         { *m = L2VpnRdAuto{} }
func (m *L2VpnRdAuto) String() string { return proto.CompactTextString(m) }
func (*L2VpnRdAuto) ProtoMessage()    {}
func (*L2VpnRdAuto) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{1}
}

func (m *L2VpnRdAuto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRdAuto.Unmarshal(m, b)
}
func (m *L2VpnRdAuto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRdAuto.Marshal(b, m, deterministic)
}
func (m *L2VpnRdAuto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRdAuto.Merge(m, src)
}
func (m *L2VpnRdAuto) XXX_Size() int {
	return xxx_messageInfo_L2VpnRdAuto.Size(m)
}
func (m *L2VpnRdAuto) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRdAuto.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRdAuto proto.InternalMessageInfo

func (m *L2VpnRdAuto) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *L2VpnRdAuto) GetAutoIndex() uint32 {
	if m != nil {
		return m.AutoIndex
	}
	return 0
}

type L2VpnRd_2ByteAs struct {
	TwoByteAs            uint32   `protobuf:"varint,1,opt,name=two_byte_as,json=twoByteAs,proto3" json:"two_byte_as,omitempty"`
	FourByteIndex        uint32   `protobuf:"varint,2,opt,name=four_byte_index,json=fourByteIndex,proto3" json:"four_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRd_2ByteAs) Reset()         { *m = L2VpnRd_2ByteAs{} }
func (m *L2VpnRd_2ByteAs) String() string { return proto.CompactTextString(m) }
func (*L2VpnRd_2ByteAs) ProtoMessage()    {}
func (*L2VpnRd_2ByteAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{2}
}

func (m *L2VpnRd_2ByteAs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRd_2ByteAs.Unmarshal(m, b)
}
func (m *L2VpnRd_2ByteAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRd_2ByteAs.Marshal(b, m, deterministic)
}
func (m *L2VpnRd_2ByteAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRd_2ByteAs.Merge(m, src)
}
func (m *L2VpnRd_2ByteAs) XXX_Size() int {
	return xxx_messageInfo_L2VpnRd_2ByteAs.Size(m)
}
func (m *L2VpnRd_2ByteAs) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRd_2ByteAs.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRd_2ByteAs proto.InternalMessageInfo

func (m *L2VpnRd_2ByteAs) GetTwoByteAs() uint32 {
	if m != nil {
		return m.TwoByteAs
	}
	return 0
}

func (m *L2VpnRd_2ByteAs) GetFourByteIndex() uint32 {
	if m != nil {
		return m.FourByteIndex
	}
	return 0
}

type L2VpnRd_4ByteAs struct {
	FourByteAs           uint32   `protobuf:"varint,1,opt,name=four_byte_as,json=fourByteAs,proto3" json:"four_byte_as,omitempty"`
	TwoByteIndex         uint32   `protobuf:"varint,2,opt,name=two_byte_index,json=twoByteIndex,proto3" json:"two_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRd_4ByteAs) Reset()         { *m = L2VpnRd_4ByteAs{} }
func (m *L2VpnRd_4ByteAs) String() string { return proto.CompactTextString(m) }
func (*L2VpnRd_4ByteAs) ProtoMessage()    {}
func (*L2VpnRd_4ByteAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{3}
}

func (m *L2VpnRd_4ByteAs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRd_4ByteAs.Unmarshal(m, b)
}
func (m *L2VpnRd_4ByteAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRd_4ByteAs.Marshal(b, m, deterministic)
}
func (m *L2VpnRd_4ByteAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRd_4ByteAs.Merge(m, src)
}
func (m *L2VpnRd_4ByteAs) XXX_Size() int {
	return xxx_messageInfo_L2VpnRd_4ByteAs.Size(m)
}
func (m *L2VpnRd_4ByteAs) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRd_4ByteAs.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRd_4ByteAs proto.InternalMessageInfo

func (m *L2VpnRd_4ByteAs) GetFourByteAs() uint32 {
	if m != nil {
		return m.FourByteAs
	}
	return 0
}

func (m *L2VpnRd_4ByteAs) GetTwoByteIndex() uint32 {
	if m != nil {
		return m.TwoByteIndex
	}
	return 0
}

type L2VpnRdV4Addr struct {
	Ipv4Address          string   `protobuf:"bytes,1,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	TwoByteIndex         uint32   `protobuf:"varint,2,opt,name=two_byte_index,json=twoByteIndex,proto3" json:"two_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRdV4Addr) Reset()         { *m = L2VpnRdV4Addr{} }
func (m *L2VpnRdV4Addr) String() string { return proto.CompactTextString(m) }
func (*L2VpnRdV4Addr) ProtoMessage()    {}
func (*L2VpnRdV4Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{4}
}

func (m *L2VpnRdV4Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRdV4Addr.Unmarshal(m, b)
}
func (m *L2VpnRdV4Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRdV4Addr.Marshal(b, m, deterministic)
}
func (m *L2VpnRdV4Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRdV4Addr.Merge(m, src)
}
func (m *L2VpnRdV4Addr) XXX_Size() int {
	return xxx_messageInfo_L2VpnRdV4Addr.Size(m)
}
func (m *L2VpnRdV4Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRdV4Addr.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRdV4Addr proto.InternalMessageInfo

func (m *L2VpnRdV4Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *L2VpnRdV4Addr) GetTwoByteIndex() uint32 {
	if m != nil {
		return m.TwoByteIndex
	}
	return 0
}

type L2VpnRd struct {
	Rd                   string           `protobuf:"bytes,1,opt,name=rd,proto3" json:"rd,omitempty"`
	Auto                 *L2VpnRdAuto     `protobuf:"bytes,2,opt,name=auto,proto3" json:"auto,omitempty"`
	TwoByteAs            *L2VpnRd_2ByteAs `protobuf:"bytes,3,opt,name=two_byte_as,json=twoByteAs,proto3" json:"two_byte_as,omitempty"`
	FourByteAs           *L2VpnRd_4ByteAs `protobuf:"bytes,4,opt,name=four_byte_as,json=fourByteAs,proto3" json:"four_byte_as,omitempty"`
	V4Addr               *L2VpnRdV4Addr   `protobuf:"bytes,5,opt,name=v4_addr,json=v4Addr,proto3" json:"v4_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *L2VpnRd) Reset()         { *m = L2VpnRd{} }
func (m *L2VpnRd) String() string { return proto.CompactTextString(m) }
func (*L2VpnRd) ProtoMessage()    {}
func (*L2VpnRd) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{5}
}

func (m *L2VpnRd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRd.Unmarshal(m, b)
}
func (m *L2VpnRd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRd.Marshal(b, m, deterministic)
}
func (m *L2VpnRd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRd.Merge(m, src)
}
func (m *L2VpnRd) XXX_Size() int {
	return xxx_messageInfo_L2VpnRd.Size(m)
}
func (m *L2VpnRd) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRd.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRd proto.InternalMessageInfo

func (m *L2VpnRd) GetRd() string {
	if m != nil {
		return m.Rd
	}
	return ""
}

func (m *L2VpnRd) GetAuto() *L2VpnRdAuto {
	if m != nil {
		return m.Auto
	}
	return nil
}

func (m *L2VpnRd) GetTwoByteAs() *L2VpnRd_2ByteAs {
	if m != nil {
		return m.TwoByteAs
	}
	return nil
}

func (m *L2VpnRd) GetFourByteAs() *L2VpnRd_4ByteAs {
	if m != nil {
		return m.FourByteAs
	}
	return nil
}

func (m *L2VpnRd) GetV4Addr() *L2VpnRdV4Addr {
	if m != nil {
		return m.V4Addr
	}
	return nil
}

type L2VpnRt_2ByteAs struct {
	TwoByteAs            uint32   `protobuf:"varint,1,opt,name=two_byte_as,json=twoByteAs,proto3" json:"two_byte_as,omitempty"`
	FourByteIndex        uint32   `protobuf:"varint,2,opt,name=four_byte_index,json=fourByteIndex,proto3" json:"four_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRt_2ByteAs) Reset()         { *m = L2VpnRt_2ByteAs{} }
func (m *L2VpnRt_2ByteAs) String() string { return proto.CompactTextString(m) }
func (*L2VpnRt_2ByteAs) ProtoMessage()    {}
func (*L2VpnRt_2ByteAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{6}
}

func (m *L2VpnRt_2ByteAs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRt_2ByteAs.Unmarshal(m, b)
}
func (m *L2VpnRt_2ByteAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRt_2ByteAs.Marshal(b, m, deterministic)
}
func (m *L2VpnRt_2ByteAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRt_2ByteAs.Merge(m, src)
}
func (m *L2VpnRt_2ByteAs) XXX_Size() int {
	return xxx_messageInfo_L2VpnRt_2ByteAs.Size(m)
}
func (m *L2VpnRt_2ByteAs) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRt_2ByteAs.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRt_2ByteAs proto.InternalMessageInfo

func (m *L2VpnRt_2ByteAs) GetTwoByteAs() uint32 {
	if m != nil {
		return m.TwoByteAs
	}
	return 0
}

func (m *L2VpnRt_2ByteAs) GetFourByteIndex() uint32 {
	if m != nil {
		return m.FourByteIndex
	}
	return 0
}

type L2VpnRt_4ByteAs struct {
	FourByteAs           uint32   `protobuf:"varint,1,opt,name=four_byte_as,json=fourByteAs,proto3" json:"four_byte_as,omitempty"`
	TwoByteIndex         uint32   `protobuf:"varint,2,opt,name=two_byte_index,json=twoByteIndex,proto3" json:"two_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRt_4ByteAs) Reset()         { *m = L2VpnRt_4ByteAs{} }
func (m *L2VpnRt_4ByteAs) String() string { return proto.CompactTextString(m) }
func (*L2VpnRt_4ByteAs) ProtoMessage()    {}
func (*L2VpnRt_4ByteAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{7}
}

func (m *L2VpnRt_4ByteAs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRt_4ByteAs.Unmarshal(m, b)
}
func (m *L2VpnRt_4ByteAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRt_4ByteAs.Marshal(b, m, deterministic)
}
func (m *L2VpnRt_4ByteAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRt_4ByteAs.Merge(m, src)
}
func (m *L2VpnRt_4ByteAs) XXX_Size() int {
	return xxx_messageInfo_L2VpnRt_4ByteAs.Size(m)
}
func (m *L2VpnRt_4ByteAs) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRt_4ByteAs.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRt_4ByteAs proto.InternalMessageInfo

func (m *L2VpnRt_4ByteAs) GetFourByteAs() uint32 {
	if m != nil {
		return m.FourByteAs
	}
	return 0
}

func (m *L2VpnRt_4ByteAs) GetTwoByteIndex() uint32 {
	if m != nil {
		return m.TwoByteIndex
	}
	return 0
}

type L2VpnRtV4Addr struct {
	Ipv4Address          string   `protobuf:"bytes,1,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	TwoByteIndex         uint32   `protobuf:"varint,2,opt,name=two_byte_index,json=twoByteIndex,proto3" json:"two_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRtV4Addr) Reset()         { *m = L2VpnRtV4Addr{} }
func (m *L2VpnRtV4Addr) String() string { return proto.CompactTextString(m) }
func (*L2VpnRtV4Addr) ProtoMessage()    {}
func (*L2VpnRtV4Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{8}
}

func (m *L2VpnRtV4Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRtV4Addr.Unmarshal(m, b)
}
func (m *L2VpnRtV4Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRtV4Addr.Marshal(b, m, deterministic)
}
func (m *L2VpnRtV4Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRtV4Addr.Merge(m, src)
}
func (m *L2VpnRtV4Addr) XXX_Size() int {
	return xxx_messageInfo_L2VpnRtV4Addr.Size(m)
}
func (m *L2VpnRtV4Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRtV4Addr.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRtV4Addr proto.InternalMessageInfo

func (m *L2VpnRtV4Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *L2VpnRtV4Addr) GetTwoByteIndex() uint32 {
	if m != nil {
		return m.TwoByteIndex
	}
	return 0
}

type L2VpnRtEsImport struct {
	HighBytes            uint32   `protobuf:"varint,1,opt,name=high_bytes,json=highBytes,proto3" json:"high_bytes,omitempty"`
	LowBytes             uint32   `protobuf:"varint,2,opt,name=low_bytes,json=lowBytes,proto3" json:"low_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRtEsImport) Reset()         { *m = L2VpnRtEsImport{} }
func (m *L2VpnRtEsImport) String() string { return proto.CompactTextString(m) }
func (*L2VpnRtEsImport) ProtoMessage()    {}
func (*L2VpnRtEsImport) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{9}
}

func (m *L2VpnRtEsImport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRtEsImport.Unmarshal(m, b)
}
func (m *L2VpnRtEsImport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRtEsImport.Marshal(b, m, deterministic)
}
func (m *L2VpnRtEsImport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRtEsImport.Merge(m, src)
}
func (m *L2VpnRtEsImport) XXX_Size() int {
	return xxx_messageInfo_L2VpnRtEsImport.Size(m)
}
func (m *L2VpnRtEsImport) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRtEsImport.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRtEsImport proto.InternalMessageInfo

func (m *L2VpnRtEsImport) GetHighBytes() uint32 {
	if m != nil {
		return m.HighBytes
	}
	return 0
}

func (m *L2VpnRtEsImport) GetLowBytes() uint32 {
	if m != nil {
		return m.LowBytes
	}
	return 0
}

type L2VpnRt struct {
	Rt                   string           `protobuf:"bytes,1,opt,name=rt,proto3" json:"rt,omitempty"`
	TwoByteAs            *L2VpnRt_2ByteAs `protobuf:"bytes,2,opt,name=two_byte_as,json=twoByteAs,proto3" json:"two_byte_as,omitempty"`
	FourByteAs           *L2VpnRt_4ByteAs `protobuf:"bytes,3,opt,name=four_byte_as,json=fourByteAs,proto3" json:"four_byte_as,omitempty"`
	V4Addr               *L2VpnRtV4Addr   `protobuf:"bytes,4,opt,name=v4_addr,json=v4Addr,proto3" json:"v4_addr,omitempty"`
	EsImport             *L2VpnRtEsImport `protobuf:"bytes,5,opt,name=es_import,json=esImport,proto3" json:"es_import,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *L2VpnRt) Reset()         { *m = L2VpnRt{} }
func (m *L2VpnRt) String() string { return proto.CompactTextString(m) }
func (*L2VpnRt) ProtoMessage()    {}
func (*L2VpnRt) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{10}
}

func (m *L2VpnRt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRt.Unmarshal(m, b)
}
func (m *L2VpnRt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRt.Marshal(b, m, deterministic)
}
func (m *L2VpnRt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRt.Merge(m, src)
}
func (m *L2VpnRt) XXX_Size() int {
	return xxx_messageInfo_L2VpnRt.Size(m)
}
func (m *L2VpnRt) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRt.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRt proto.InternalMessageInfo

func (m *L2VpnRt) GetRt() string {
	if m != nil {
		return m.Rt
	}
	return ""
}

func (m *L2VpnRt) GetTwoByteAs() *L2VpnRt_2ByteAs {
	if m != nil {
		return m.TwoByteAs
	}
	return nil
}

func (m *L2VpnRt) GetFourByteAs() *L2VpnRt_4ByteAs {
	if m != nil {
		return m.FourByteAs
	}
	return nil
}

func (m *L2VpnRt) GetV4Addr() *L2VpnRtV4Addr {
	if m != nil {
		return m.V4Addr
	}
	return nil
}

func (m *L2VpnRt) GetEsImport() *L2VpnRtEsImport {
	if m != nil {
		return m.EsImport
	}
	return nil
}

type L2VpnXcDisco struct {
	AdMethod              uint32     `protobuf:"varint,1,opt,name=ad_method,json=adMethod,proto3" json:"ad_method,omitempty"`
	VpnAdded              bool       `protobuf:"varint,2,opt,name=vpn_added,json=vpnAdded,proto3" json:"vpn_added,omitempty"`
	AdServiceConnected    bool       `protobuf:"varint,3,opt,name=ad_service_connected,json=adServiceConnected,proto3" json:"ad_service_connected,omitempty"`
	RdValue               *L2VpnRd   `protobuf:"bytes,4,opt,name=rd_value,json=rdValue,proto3" json:"rd_value,omitempty"`
	ImportRt              []*L2VpnRt `protobuf:"bytes,5,rep,name=import_rt,json=importRt,proto3" json:"import_rt,omitempty"`
	ExportRt              []*L2VpnRt `protobuf:"bytes,6,rep,name=export_rt,json=exportRt,proto3" json:"export_rt,omitempty"`
	AdSignallingMethod    uint32     `protobuf:"varint,7,opt,name=ad_signalling_method,json=adSignallingMethod,proto3" json:"ad_signalling_method,omitempty"`
	CeRange               uint32     `protobuf:"varint,8,opt,name=ce_range,json=ceRange,proto3" json:"ce_range,omitempty"`
	ExportRoutePolicy     string     `protobuf:"bytes,9,opt,name=export_route_policy,json=exportRoutePolicy,proto3" json:"export_route_policy,omitempty"`
	NumberACsUp           uint32     `protobuf:"varint,10,opt,name=number_a_cs_up,json=numberACsUp,proto3" json:"number_a_cs_up,omitempty"`
	NumberPWsUp           uint32     `protobuf:"varint,11,opt,name=number_p_ws_up,json=numberPWsUp,proto3" json:"number_p_ws_up,omitempty"`
	NumberCe2CeAdvertised uint32     `protobuf:"varint,12,opt,name=number_ce2ce_advertised,json=numberCe2ceAdvertised,proto3" json:"number_ce2ce_advertised,omitempty"`
	NumberACs             uint32     `protobuf:"varint,13,opt,name=number_a_cs,json=numberACs,proto3" json:"number_a_cs,omitempty"`
	NumberPseudowires     uint32     `protobuf:"varint,14,opt,name=number_pseudowires,json=numberPseudowires,proto3" json:"number_pseudowires,omitempty"`
	NumberCe2CEs          uint32     `protobuf:"varint,15,opt,name=number_ce2c_es,json=numberCe2cEs,proto3" json:"number_ce2c_es,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}   `json:"-"`
	XXX_unrecognized      []byte     `json:"-"`
	XXX_sizecache         int32      `json:"-"`
}

func (m *L2VpnXcDisco) Reset()         { *m = L2VpnXcDisco{} }
func (m *L2VpnXcDisco) String() string { return proto.CompactTextString(m) }
func (*L2VpnXcDisco) ProtoMessage()    {}
func (*L2VpnXcDisco) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{11}
}

func (m *L2VpnXcDisco) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnXcDisco.Unmarshal(m, b)
}
func (m *L2VpnXcDisco) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnXcDisco.Marshal(b, m, deterministic)
}
func (m *L2VpnXcDisco) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnXcDisco.Merge(m, src)
}
func (m *L2VpnXcDisco) XXX_Size() int {
	return xxx_messageInfo_L2VpnXcDisco.Size(m)
}
func (m *L2VpnXcDisco) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnXcDisco.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnXcDisco proto.InternalMessageInfo

func (m *L2VpnXcDisco) GetAdMethod() uint32 {
	if m != nil {
		return m.AdMethod
	}
	return 0
}

func (m *L2VpnXcDisco) GetVpnAdded() bool {
	if m != nil {
		return m.VpnAdded
	}
	return false
}

func (m *L2VpnXcDisco) GetAdServiceConnected() bool {
	if m != nil {
		return m.AdServiceConnected
	}
	return false
}

func (m *L2VpnXcDisco) GetRdValue() *L2VpnRd {
	if m != nil {
		return m.RdValue
	}
	return nil
}

func (m *L2VpnXcDisco) GetImportRt() []*L2VpnRt {
	if m != nil {
		return m.ImportRt
	}
	return nil
}

func (m *L2VpnXcDisco) GetExportRt() []*L2VpnRt {
	if m != nil {
		return m.ExportRt
	}
	return nil
}

func (m *L2VpnXcDisco) GetAdSignallingMethod() uint32 {
	if m != nil {
		return m.AdSignallingMethod
	}
	return 0
}

func (m *L2VpnXcDisco) GetCeRange() uint32 {
	if m != nil {
		return m.CeRange
	}
	return 0
}

func (m *L2VpnXcDisco) GetExportRoutePolicy() string {
	if m != nil {
		return m.ExportRoutePolicy
	}
	return ""
}

func (m *L2VpnXcDisco) GetNumberACsUp() uint32 {
	if m != nil {
		return m.NumberACsUp
	}
	return 0
}

func (m *L2VpnXcDisco) GetNumberPWsUp() uint32 {
	if m != nil {
		return m.NumberPWsUp
	}
	return 0
}

func (m *L2VpnXcDisco) GetNumberCe2CeAdvertised() uint32 {
	if m != nil {
		return m.NumberCe2CeAdvertised
	}
	return 0
}

func (m *L2VpnXcDisco) GetNumberACs() uint32 {
	if m != nil {
		return m.NumberACs
	}
	return 0
}

func (m *L2VpnXcDisco) GetNumberPseudowires() uint32 {
	if m != nil {
		return m.NumberPseudowires
	}
	return 0
}

func (m *L2VpnXcDisco) GetNumberCe2CEs() uint32 {
	if m != nil {
		return m.NumberCe2CEs
	}
	return 0
}

type L2VpnMp2MpXc struct {
	GroupNameXr          string        `protobuf:"bytes,50,opt,name=group_name_xr,json=groupNameXr,proto3" json:"group_name_xr,omitempty"`
	Mp2MpName            string        `protobuf:"bytes,51,opt,name=mp2mp_name,json=mp2mpName,proto3" json:"mp2mp_name,omitempty"`
	Mp2Mpid              uint32        `protobuf:"varint,52,opt,name=mp2mpid,proto3" json:"mp2mpid,omitempty"`
	VpnId                uint32        `protobuf:"varint,53,opt,name=vpn_id,json=vpnId,proto3" json:"vpn_id,omitempty"`
	VpnMtu               uint32        `protobuf:"varint,54,opt,name=vpn_mtu,json=vpnMtu,proto3" json:"vpn_mtu,omitempty"`
	L2Encapsulation      string        `protobuf:"bytes,55,opt,name=l2_encapsulation,json=l2Encapsulation,proto3" json:"l2_encapsulation,omitempty"`
	Discovery            *L2VpnXcDisco `protobuf:"bytes,56,opt,name=discovery,proto3" json:"discovery,omitempty"`
	XconnectShutdown     bool          `protobuf:"varint,57,opt,name=xconnect_shutdown,json=xconnectShutdown,proto3" json:"xconnect_shutdown,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *L2VpnMp2MpXc) Reset()         { *m = L2VpnMp2MpXc{} }
func (m *L2VpnMp2MpXc) String() string { return proto.CompactTextString(m) }
func (*L2VpnMp2MpXc) ProtoMessage()    {}
func (*L2VpnMp2MpXc) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f1c5dc6a27dd31, []int{12}
}

func (m *L2VpnMp2MpXc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMp2MpXc.Unmarshal(m, b)
}
func (m *L2VpnMp2MpXc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMp2MpXc.Marshal(b, m, deterministic)
}
func (m *L2VpnMp2MpXc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMp2MpXc.Merge(m, src)
}
func (m *L2VpnMp2MpXc) XXX_Size() int {
	return xxx_messageInfo_L2VpnMp2MpXc.Size(m)
}
func (m *L2VpnMp2MpXc) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMp2MpXc.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMp2MpXc proto.InternalMessageInfo

func (m *L2VpnMp2MpXc) GetGroupNameXr() string {
	if m != nil {
		return m.GroupNameXr
	}
	return ""
}

func (m *L2VpnMp2MpXc) GetMp2MpName() string {
	if m != nil {
		return m.Mp2MpName
	}
	return ""
}

func (m *L2VpnMp2MpXc) GetMp2Mpid() uint32 {
	if m != nil {
		return m.Mp2Mpid
	}
	return 0
}

func (m *L2VpnMp2MpXc) GetVpnId() uint32 {
	if m != nil {
		return m.VpnId
	}
	return 0
}

func (m *L2VpnMp2MpXc) GetVpnMtu() uint32 {
	if m != nil {
		return m.VpnMtu
	}
	return 0
}

func (m *L2VpnMp2MpXc) GetL2Encapsulation() string {
	if m != nil {
		return m.L2Encapsulation
	}
	return ""
}

func (m *L2VpnMp2MpXc) GetDiscovery() *L2VpnXcDisco {
	if m != nil {
		return m.Discovery
	}
	return nil
}

func (m *L2VpnMp2MpXc) GetXconnectShutdown() bool {
	if m != nil {
		return m.XconnectShutdown
	}
	return false
}

func init() {
	proto.RegisterType((*L2VpnMp2MpXc_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_mp2mp_xc_KEYS")
	proto.RegisterType((*L2VpnRdAuto)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_rd_auto")
	proto.RegisterType((*L2VpnRd_2ByteAs)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_rd_2byte_as")
	proto.RegisterType((*L2VpnRd_4ByteAs)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_rd_4byte_as")
	proto.RegisterType((*L2VpnRdV4Addr)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_rd_v4addr")
	proto.RegisterType((*L2VpnRd)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_rd")
	proto.RegisterType((*L2VpnRt_2ByteAs)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_rt_2byte_as")
	proto.RegisterType((*L2VpnRt_4ByteAs)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_rt_4byte_as")
	proto.RegisterType((*L2VpnRtV4Addr)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_rt_v4addr")
	proto.RegisterType((*L2VpnRtEsImport)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_rt_es_import")
	proto.RegisterType((*L2VpnRt)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_rt")
	proto.RegisterType((*L2VpnXcDisco)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_xc_disco")
	proto.RegisterType((*L2VpnMp2MpXc)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mps.xconnect_mp2mp.l2vpn_mp2mp_xc")
}

func init() { proto.RegisterFile("l2vpn_mp2mp_xc.proto", fileDescriptor_60f1c5dc6a27dd31) }

var fileDescriptor_60f1c5dc6a27dd31 = []byte{
	// 930 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0x5d, 0x6f, 0x1c, 0x35,
	0x14, 0x55, 0x93, 0x74, 0x77, 0xf6, 0x6e, 0x76, 0xd3, 0xb8, 0xad, 0x3a, 0x08, 0x81, 0xca, 0x82,
	0x10, 0x08, 0xb1, 0x42, 0xd3, 0x50, 0xe0, 0x71, 0x89, 0x82, 0x14, 0x55, 0x41, 0xd1, 0x84, 0xb6,
	0xb4, 0x3c, 0x58, 0xce, 0xd8, 0x6c, 0x46, 0xcc, 0x8c, 0x2d, 0xdb, 0xb3, 0x1f, 0xe2, 0xa7, 0xf1,
	0x93, 0x78, 0xe0, 0x17, 0x20, 0x21, 0x5f, 0x7b, 0x66, 0x27, 0xe1, 0xa1, 0x7d, 0xd8, 0x7d, 0x1b,
	0xdf, 0x73, 0x7c, 0x8e, 0xe3, 0xfb, 0x11, 0x2f, 0x3c, 0x2a, 0x92, 0x85, 0xaa, 0x68, 0xa9, 0x92,
	0x52, 0xd1, 0x55, 0x36, 0x55, 0x5a, 0x5a, 0x49, 0x7e, 0xca, 0x72, 0x93, 0x49, 0x9a, 0x4b, 0x43,
	0x57, 0x9a, 0x7a, 0x8a, 0x54, 0x42, 0x4f, 0xf1, 0x73, 0x91, 0x4c, 0x8d, 0x65, 0x15, 0xbf, 0x5e,
	0x4f, 0x57, 0x99, 0xac, 0x2a, 0x91, 0x59, 0x2f, 0x60, 0xee, 0xac, 0x27, 0xbf, 0xc0, 0xc3, 0xdb,
	0xfa, 0xf4, 0xc5, 0xd9, 0x9b, 0x2b, 0xf2, 0x11, 0xc0, 0x5c, 0xcb, 0x5a, 0xd1, 0x8a, 0x95, 0x22,
	0xbe, 0xf7, 0xf4, 0xde, 0x17, 0x83, 0x74, 0x80, 0x91, 0x9f, 0x59, 0x29, 0xc8, 0xc7, 0x30, 0x2c,
	0x55, 0x42, 0xcb, 0x80, 0xef, 0x79, 0xbc, 0x54, 0xc9, 0x05, 0xe2, 0x93, 0x17, 0x30, 0xf2, 0xaa,
	0x9a, 0x53, 0x56, 0x5b, 0x49, 0x3e, 0x84, 0x81, 0x96, 0xb5, 0x15, 0x9a, 0xe6, 0x3c, 0xc8, 0x45,
	0x3e, 0x70, 0xce, 0x9d, 0x99, 0x23, 0xd1, 0xbc, 0xe2, 0x62, 0x85, 0x62, 0xa3, 0x74, 0xe0, 0x22,
	0xe7, 0x2e, 0x30, 0xf9, 0x0d, 0x8e, 0x5b, 0xb1, 0xe4, 0x7a, 0x6d, 0x05, 0x65, 0xc6, 0x9d, 0xc0,
	0x2e, 0x25, 0x0d, 0x4b, 0x94, 0x1c, 0xa5, 0x03, 0xbb, 0x94, 0x3f, 0xae, 0xad, 0x98, 0x19, 0xf2,
	0x39, 0x1c, 0xfd, 0x2e, 0x6b, 0xed, 0x09, 0x5d, 0xe1, 0x91, 0x0b, 0x3b, 0xd2, 0xff, 0xc5, 0x4f,
	0x1a, 0xf1, 0xa7, 0x70, 0xb8, 0xd9, 0xdc, 0xaa, 0x43, 0xb3, 0x73, 0x66, 0xc8, 0x67, 0x30, 0x6e,
	0xed, 0xbb, 0xea, 0x87, 0xe1, 0x04, 0x5e, 0xfc, 0x2d, 0x1c, 0xb5, 0xe2, 0x8b, 0x13, 0xc6, 0xb9,
	0x26, 0x9f, 0xc0, 0x61, 0xae, 0x16, 0x27, 0xd4, 0x2d, 0x84, 0x31, 0xe1, 0x2e, 0x86, 0x2e, 0x36,
	0xf3, 0xa1, 0xf7, 0xd4, 0xfe, 0x67, 0x1f, 0xa2, 0x46, 0x9c, 0x8c, 0x61, 0x4f, 0x37, 0xf7, 0xba,
	0xa7, 0x39, 0xc9, 0xe1, 0xc0, 0xdd, 0x1f, 0x6e, 0x1c, 0x26, 0x2f, 0xa7, 0xdb, 0x29, 0x96, 0xe9,
	0xad, 0x9c, 0xa6, 0x68, 0x41, 0xd6, 0xb7, 0x13, 0xb1, 0x8f, 0x8e, 0x6f, 0xb6, 0xee, 0xd8, 0x24,
	0xbe, 0x9b, 0xe3, 0x3f, 0xef, 0xa4, 0xe9, 0x60, 0x47, 0xde, 0x4d, 0x5d, 0xdc, 0xaa, 0x00, 0x05,
	0xfd, 0x90, 0xc6, 0xf8, 0x3e, 0xfa, 0xbe, 0xde, 0xba, 0xaf, 0x2f, 0x99, 0xb4, 0xe7, 0x4b, 0xa3,
	0x53, 0xaa, 0x76, 0x97, 0x7d, 0x60, 0x77, 0xd9, 0x07, 0x76, 0xeb, 0x7d, 0x70, 0x09, 0xa4, 0xd5,
	0x16, 0x86, 0xe6, 0xa5, 0x92, 0xda, 0xba, 0x91, 0x72, 0x93, 0xcf, 0x6f, 0x70, 0x73, 0x7b, 0x2b,
	0x2e, 0xe2, 0x36, 0x1a, 0x37, 0x8e, 0x0a, 0xb9, 0x0c, 0xa8, 0x57, 0x8d, 0x0a, 0xb9, 0x44, 0x70,
	0xf2, 0xef, 0xa6, 0xb3, 0x2c, 0x76, 0x96, 0x6d, 0x3b, 0xcb, 0xde, 0x2d, 0xf7, 0xbd, 0x9d, 0x94,
	0x9c, 0x7d, 0xaf, 0x72, 0xdf, 0xdf, 0x91, 0xf7, 0xbb, 0xca, 0xfd, 0x60, 0x27, 0xe5, 0x6e, 0xef,
	0x94, 0x3b, 0x59, 0xc2, 0xa0, 0xcd, 0x67, 0x68, 0xb1, 0xb7, 0x5b, 0xf7, 0x6c, 0x1d, 0xd2, 0x48,
	0x98, 0x73, 0xfc, 0x9a, 0xfc, 0xd5, 0x83, 0xb1, 0x27, 0xac, 0x32, 0xca, 0x9d, 0xa1, 0xab, 0x17,
	0xc6, 0x69, 0x29, 0xec, 0x8d, 0xe4, 0xa1, 0x9a, 0x22, 0xc6, 0x2f, 0x70, 0xed, 0x40, 0x47, 0x66,
	0x9c, 0x0b, 0x8e, 0x05, 0x11, 0xa5, 0xd1, 0x42, 0x55, 0x33, 0xb7, 0x26, 0xdf, 0xc0, 0x23, 0xc6,
	0xa9, 0x11, 0x7a, 0x91, 0x67, 0x82, 0x86, 0x53, 0x08, 0x8e, 0xc9, 0x8b, 0x52, 0xc2, 0xf8, 0x95,
	0x87, 0x4e, 0x1b, 0x84, 0xfc, 0x01, 0x91, 0xeb, 0x7d, 0x56, 0xd4, 0x22, 0x5c, 0xf5, 0xe5, 0xb6,
	0x27, 0x4b, 0xda, 0xd7, 0xfc, 0x95, 0x33, 0x20, 0x25, 0x0c, 0xfc, 0xdf, 0x4f, 0xf1, 0x92, 0xf7,
	0x77, 0xe0, 0x66, 0xd3, 0x28, 0x5c, 0xb1, 0x75, 0x76, 0x62, 0xd5, 0xd8, 0xf5, 0x76, 0x65, 0xe7,
	0x2d, 0x52, 0xdb, 0x5c, 0x7e, 0x3e, 0xaf, 0x58, 0x51, 0xe4, 0xd5, 0xbc, 0xc9, 0x60, 0x1f, 0x33,
	0xe8, 0x2e, 0xbf, 0x85, 0x42, 0x2e, 0x3f, 0x80, 0x28, 0x13, 0x54, 0xb3, 0x6a, 0x2e, 0xe2, 0x08,
	0x59, 0xfd, 0x4c, 0xa4, 0x6e, 0x49, 0xa6, 0xf0, 0xb0, 0x39, 0xbb, 0x7b, 0xb8, 0x50, 0x25, 0x8b,
	0x3c, 0x5b, 0xc7, 0x03, 0x1c, 0x0d, 0xc7, 0xc1, 0xd3, 0x21, 0x97, 0x08, 0x90, 0x4f, 0x61, 0x5c,
	0xd5, 0xe5, 0xb5, 0xd0, 0x94, 0xd1, 0xcc, 0xd0, 0x5a, 0xc5, 0x80, 0x82, 0x43, 0x1f, 0x9d, 0x9d,
	0x9a, 0x97, 0xaa, 0x43, 0x52, 0x74, 0x89, 0xa4, 0x61, 0x97, 0x74, 0xf9, 0xda, 0x91, 0x9e, 0xc3,
	0x93, 0x40, 0xca, 0x44, 0x92, 0x09, 0xca, 0xf8, 0x42, 0x68, 0x9b, 0x1b, 0xc1, 0xe3, 0x43, 0x64,
	0x3f, 0xf6, 0xf0, 0xa9, 0x43, 0x67, 0x2d, 0xe8, 0xfe, 0x37, 0x74, 0x4e, 0x10, 0x8f, 0xfc, 0x14,
	0x6c, 0xed, 0xc9, 0xd7, 0x40, 0x1a, 0x73, 0x23, 0x6a, 0x2e, 0x97, 0xb9, 0x16, 0x26, 0x1e, 0x23,
	0xed, 0x38, 0x1c, 0x60, 0x03, 0xb8, 0x79, 0xdc, 0x39, 0x06, 0x15, 0x26, 0x3e, 0xf2, 0xf3, 0x78,
	0xe3, 0x7e, 0x66, 0x26, 0x7f, 0xef, 0x35, 0xdd, 0xd3, 0xbc, 0x28, 0xc9, 0x04, 0x46, 0x9b, 0xc7,
	0x24, 0x5d, 0xe9, 0x38, 0xf1, 0xc3, 0xbe, 0x7d, 0x4f, 0xfe, 0xaa, 0xdd, 0xc0, 0xf6, 0x7c, 0x7c,
	0x50, 0x3e, 0x6b, 0x1f, 0x94, 0xa5, 0x7f, 0x70, 0xc6, 0xd0, 0xc7, 0x45, 0xce, 0xe3, 0x13, 0x9f,
	0x96, 0xb0, 0x24, 0x8f, 0xa1, 0xe7, 0xcc, 0x72, 0x1e, 0x7f, 0x8b, 0xc0, 0xfd, 0x85, 0xaa, 0xce,
	0x39, 0x79, 0x02, 0x7d, 0x3c, 0x83, 0xad, 0xe3, 0xe7, 0x18, 0x77, 0xac, 0x0b, 0x5b, 0x93, 0x2f,
	0xe1, 0x41, 0x91, 0x50, 0x51, 0x65, 0x4c, 0x99, 0xba, 0x60, 0x36, 0x97, 0x55, 0xfc, 0x1d, 0xda,
	0x1d, 0x15, 0xc9, 0x59, 0x37, 0x4c, 0x2c, 0x0c, 0xb0, 0xfd, 0x17, 0x42, 0xaf, 0xe3, 0xef, 0xb1,
	0x15, 0x5f, 0x6d, 0xb7, 0x5a, 0x9b, 0x01, 0x93, 0x6e, 0x8c, 0xc8, 0x57, 0x70, 0xdc, 0x72, 0xcd,
	0x4d, 0x6d, 0xb9, 0x5c, 0x56, 0xf1, 0x0f, 0x38, 0x2e, 0x1e, 0x34, 0xc0, 0x55, 0x88, 0x5f, 0xf7,
	0xf0, 0xd7, 0xc0, 0xb3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x84, 0x03, 0x1c, 0x25, 0x0c,
	0x00, 0x00,
}
