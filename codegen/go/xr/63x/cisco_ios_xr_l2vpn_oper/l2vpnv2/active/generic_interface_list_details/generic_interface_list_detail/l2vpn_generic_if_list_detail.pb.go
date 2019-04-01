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
// source: l2vpn_generic_if_list_detail.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_active_generic_interface_list_details_generic_interface_list_detail

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

type L2VpnGenericIfListDetail_KEYS struct {
	InterfaceListName    string   `protobuf:"bytes,1,opt,name=interface_list_name,json=interfaceListName,proto3" json:"interface_list_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnGenericIfListDetail_KEYS) Reset()         { *m = L2VpnGenericIfListDetail_KEYS{} }
func (m *L2VpnGenericIfListDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnGenericIfListDetail_KEYS) ProtoMessage()    {}
func (*L2VpnGenericIfListDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d38fde8f332160a, []int{0}
}

func (m *L2VpnGenericIfListDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnGenericIfListDetail_KEYS.Unmarshal(m, b)
}
func (m *L2VpnGenericIfListDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnGenericIfListDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnGenericIfListDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnGenericIfListDetail_KEYS.Merge(m, src)
}
func (m *L2VpnGenericIfListDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnGenericIfListDetail_KEYS.Size(m)
}
func (m *L2VpnGenericIfListDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnGenericIfListDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnGenericIfListDetail_KEYS proto.InternalMessageInfo

func (m *L2VpnGenericIfListDetail_KEYS) GetInterfaceListName() string {
	if m != nil {
		return m.InterfaceListName
	}
	return ""
}

type InterfacesItem struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfacesItem) Reset()         { *m = InterfacesItem{} }
func (m *InterfacesItem) String() string { return proto.CompactTextString(m) }
func (*InterfacesItem) ProtoMessage()    {}
func (*InterfacesItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d38fde8f332160a, []int{1}
}

func (m *InterfacesItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesItem.Unmarshal(m, b)
}
func (m *InterfacesItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesItem.Marshal(b, m, deterministic)
}
func (m *InterfacesItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesItem.Merge(m, src)
}
func (m *InterfacesItem) XXX_Size() int {
	return xxx_messageInfo_InterfacesItem.Size(m)
}
func (m *InterfacesItem) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesItem.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesItem proto.InternalMessageInfo

func (m *InterfacesItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type L2VpnGenericIfList struct {
	InterfaceListNameXr  string            `protobuf:"bytes,1,opt,name=interface_list_name_xr,json=interfaceListNameXr,proto3" json:"interface_list_name_xr,omitempty"`
	InterfaceListId      uint32            `protobuf:"varint,2,opt,name=interface_list_id,json=interfaceListId,proto3" json:"interface_list_id,omitempty"`
	Interface            []*InterfacesItem `protobuf:"bytes,3,rep,name=interface,proto3" json:"interface,omitempty"`
	NumberOfInterfaces   uint32            `protobuf:"varint,4,opt,name=number_of_interfaces,json=numberOfInterfaces,proto3" json:"number_of_interfaces,omitempty"`
	Items                uint32            `protobuf:"varint,5,opt,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L2VpnGenericIfList) Reset()         { *m = L2VpnGenericIfList{} }
func (m *L2VpnGenericIfList) String() string { return proto.CompactTextString(m) }
func (*L2VpnGenericIfList) ProtoMessage()    {}
func (*L2VpnGenericIfList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d38fde8f332160a, []int{2}
}

func (m *L2VpnGenericIfList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnGenericIfList.Unmarshal(m, b)
}
func (m *L2VpnGenericIfList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnGenericIfList.Marshal(b, m, deterministic)
}
func (m *L2VpnGenericIfList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnGenericIfList.Merge(m, src)
}
func (m *L2VpnGenericIfList) XXX_Size() int {
	return xxx_messageInfo_L2VpnGenericIfList.Size(m)
}
func (m *L2VpnGenericIfList) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnGenericIfList.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnGenericIfList proto.InternalMessageInfo

func (m *L2VpnGenericIfList) GetInterfaceListNameXr() string {
	if m != nil {
		return m.InterfaceListNameXr
	}
	return ""
}

func (m *L2VpnGenericIfList) GetInterfaceListId() uint32 {
	if m != nil {
		return m.InterfaceListId
	}
	return 0
}

func (m *L2VpnGenericIfList) GetInterface() []*InterfacesItem {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *L2VpnGenericIfList) GetNumberOfInterfaces() uint32 {
	if m != nil {
		return m.NumberOfInterfaces
	}
	return 0
}

func (m *L2VpnGenericIfList) GetItems() uint32 {
	if m != nil {
		return m.Items
	}
	return 0
}

type L2VpnPwheRange struct {
	Lower                uint32   `protobuf:"varint,1,opt,name=lower,proto3" json:"lower,omitempty"`
	Upper                uint32   `protobuf:"varint,2,opt,name=upper,proto3" json:"upper,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnPwheRange) Reset()         { *m = L2VpnPwheRange{} }
func (m *L2VpnPwheRange) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwheRange) ProtoMessage()    {}
func (*L2VpnPwheRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d38fde8f332160a, []int{3}
}

func (m *L2VpnPwheRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwheRange.Unmarshal(m, b)
}
func (m *L2VpnPwheRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwheRange.Marshal(b, m, deterministic)
}
func (m *L2VpnPwheRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwheRange.Merge(m, src)
}
func (m *L2VpnPwheRange) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwheRange.Size(m)
}
func (m *L2VpnPwheRange) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwheRange.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwheRange proto.InternalMessageInfo

func (m *L2VpnPwheRange) GetLower() uint32 {
	if m != nil {
		return m.Lower
	}
	return 0
}

func (m *L2VpnPwheRange) GetUpper() uint32 {
	if m != nil {
		return m.Upper
	}
	return 0
}

type L2VpnGenericIfListDetail struct {
	Summary              *L2VpnGenericIfList `protobuf:"bytes,50,opt,name=summary,proto3" json:"summary,omitempty"`
	Interface            []*InterfacesItem   `protobuf:"bytes,51,rep,name=interface,proto3" json:"interface,omitempty"`
	Items                uint32              `protobuf:"varint,52,opt,name=items,proto3" json:"items,omitempty"`
	IsProvisioned        bool                `protobuf:"varint,53,opt,name=is_provisioned,json=isProvisioned,proto3" json:"is_provisioned,omitempty"`
	PsedowireEtherItems  uint32              `protobuf:"varint,54,opt,name=psedowire_ether_items,json=psedowireEtherItems,proto3" json:"psedowire_ether_items,omitempty"`
	PseudowireEtherRange []*L2VpnPwheRange   `protobuf:"bytes,55,rep,name=pseudowire_ether_range,json=pseudowireEtherRange,proto3" json:"pseudowire_ether_range,omitempty"`
	PseudowireIwItems    uint32              `protobuf:"varint,56,opt,name=pseudowire_iw_items,json=pseudowireIwItems,proto3" json:"pseudowire_iw_items,omitempty"`
	PseudowireIwRange    []*L2VpnPwheRange   `protobuf:"bytes,57,rep,name=pseudowire_iw_range,json=pseudowireIwRange,proto3" json:"pseudowire_iw_range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *L2VpnGenericIfListDetail) Reset()         { *m = L2VpnGenericIfListDetail{} }
func (m *L2VpnGenericIfListDetail) String() string { return proto.CompactTextString(m) }
func (*L2VpnGenericIfListDetail) ProtoMessage()    {}
func (*L2VpnGenericIfListDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d38fde8f332160a, []int{4}
}

func (m *L2VpnGenericIfListDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnGenericIfListDetail.Unmarshal(m, b)
}
func (m *L2VpnGenericIfListDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnGenericIfListDetail.Marshal(b, m, deterministic)
}
func (m *L2VpnGenericIfListDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnGenericIfListDetail.Merge(m, src)
}
func (m *L2VpnGenericIfListDetail) XXX_Size() int {
	return xxx_messageInfo_L2VpnGenericIfListDetail.Size(m)
}
func (m *L2VpnGenericIfListDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnGenericIfListDetail.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnGenericIfListDetail proto.InternalMessageInfo

func (m *L2VpnGenericIfListDetail) GetSummary() *L2VpnGenericIfList {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *L2VpnGenericIfListDetail) GetInterface() []*InterfacesItem {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *L2VpnGenericIfListDetail) GetItems() uint32 {
	if m != nil {
		return m.Items
	}
	return 0
}

func (m *L2VpnGenericIfListDetail) GetIsProvisioned() bool {
	if m != nil {
		return m.IsProvisioned
	}
	return false
}

func (m *L2VpnGenericIfListDetail) GetPsedowireEtherItems() uint32 {
	if m != nil {
		return m.PsedowireEtherItems
	}
	return 0
}

func (m *L2VpnGenericIfListDetail) GetPseudowireEtherRange() []*L2VpnPwheRange {
	if m != nil {
		return m.PseudowireEtherRange
	}
	return nil
}

func (m *L2VpnGenericIfListDetail) GetPseudowireIwItems() uint32 {
	if m != nil {
		return m.PseudowireIwItems
	}
	return 0
}

func (m *L2VpnGenericIfListDetail) GetPseudowireIwRange() []*L2VpnPwheRange {
	if m != nil {
		return m.PseudowireIwRange
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnGenericIfListDetail_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.generic_interface_list_details.generic_interface_list_detail.l2vpn_generic_if_list_detail_KEYS")
	proto.RegisterType((*InterfacesItem)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.generic_interface_list_details.generic_interface_list_detail.interfaces_item")
	proto.RegisterType((*L2VpnGenericIfList)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.generic_interface_list_details.generic_interface_list_detail.l2vpn_generic_if_list")
	proto.RegisterType((*L2VpnPwheRange)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.generic_interface_list_details.generic_interface_list_detail.l2vpn_pwhe_range")
	proto.RegisterType((*L2VpnGenericIfListDetail)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.generic_interface_list_details.generic_interface_list_detail.l2vpn_generic_if_list_detail")
}

func init() { proto.RegisterFile("l2vpn_generic_if_list_detail.proto", fileDescriptor_0d38fde8f332160a) }

var fileDescriptor_0d38fde8f332160a = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x95, 0x5b, 0x5a, 0xa8, 0xab, 0xa1, 0x8c, 0x67, 0x5a, 0x65, 0xc1, 0x22, 0x44, 0x42, 0x44,
	0x2c, 0x22, 0x94, 0xe1, 0xb9, 0x61, 0xd7, 0xc5, 0x08, 0x04, 0x28, 0xdd, 0xc0, 0xca, 0x4a, 0x93,
	0x9b, 0xd6, 0x52, 0x12, 0x5b, 0xb6, 0x93, 0x14, 0x3e, 0x03, 0xf1, 0x0b, 0xac, 0xf8, 0x0a, 0x3e,
	0x82, 0xff, 0x41, 0xb1, 0x67, 0x92, 0x99, 0x30, 0x9a, 0xe5, 0xd0, 0x5d, 0x7c, 0xaf, 0xcf, 0xf1,
	0x39, 0xe7, 0x5e, 0x05, 0x7b, 0x79, 0x58, 0x8b, 0x92, 0x5e, 0x41, 0x09, 0x92, 0x25, 0x94, 0x65,
	0x34, 0x67, 0x4a, 0xd3, 0x14, 0x74, 0xcc, 0xf2, 0x40, 0x48, 0xae, 0x39, 0x49, 0x12, 0xa6, 0x12,
	0x4e, 0x19, 0x57, 0xf4, 0x46, 0x52, 0x0b, 0xe0, 0x02, 0x64, 0x60, 0x3e, 0xeb, 0x30, 0x88, 0x13,
	0xcd, 0x6a, 0x08, 0x3a, 0x92, 0x52, 0x83, 0xcc, 0xe2, 0x04, 0x56, 0xb9, 0xd4, 0xf6, 0xb6, 0x77,
	0x81, 0x1f, 0x6d, 0x93, 0x42, 0xdf, 0x9d, 0x7f, 0xb9, 0x20, 0x01, 0x9e, 0x0c, 0xd0, 0x65, 0x5c,
	0x80, 0x83, 0x5c, 0xe4, 0x1f, 0x45, 0xe3, 0xae, 0xf5, 0x9e, 0x29, 0xfd, 0x21, 0x2e, 0xc0, 0x7b,
	0x82, 0x4f, 0xba, 0xa2, 0xa2, 0x4c, 0x43, 0x41, 0xa6, 0xf8, 0xa0, 0x8e, 0xf3, 0x6a, 0x09, 0xb2,
	0x07, 0xef, 0xcf, 0x1e, 0x3e, 0xdd, 0xf8, 0x3c, 0x99, 0xe1, 0xb3, 0x0d, 0x4f, 0xd2, 0x1b, 0xb9,
	0x20, 0x98, 0xfc, 0xf3, 0xea, 0x67, 0x49, 0x9e, 0xe2, 0xf1, 0x00, 0xc4, 0x52, 0x67, 0xcf, 0x45,
	0xfe, 0x28, 0x3a, 0x59, 0xbb, 0x3f, 0x4f, 0xc9, 0x77, 0x84, 0x8f, 0xba, 0x9a, 0xb3, 0xef, 0xee,
	0xfb, 0xc7, 0xa1, 0x0e, 0x76, 0x10, 0x79, 0x30, 0x88, 0x26, 0xea, 0x65, 0x90, 0x67, 0x78, 0x5a,
	0x56, 0xc5, 0x25, 0x48, 0xca, 0xb3, 0x1e, 0xad, 0x9c, 0x3b, 0xc6, 0x03, 0xb1, 0xbd, 0x8f, 0xd9,
	0xbc, 0xeb, 0xb4, 0xb9, 0xb6, 0x24, 0xca, 0x39, 0x30, 0x57, 0xec, 0xc1, 0x7b, 0x8b, 0x1f, 0x58,
	0xf1, 0xa2, 0xb9, 0x06, 0x2a, 0xe3, 0xf2, 0x0a, 0xda, 0x9b, 0x39, 0x6f, 0xc0, 0x06, 0x38, 0x8a,
	0xec, 0xa1, 0xad, 0x56, 0x42, 0x80, 0x5c, 0xc4, 0x64, 0x0f, 0xde, 0xef, 0x43, 0xfc, 0x70, 0xdb,
	0x5a, 0x90, 0x1f, 0x08, 0xdf, 0x55, 0x55, 0x51, 0xc4, 0xf2, 0xab, 0x13, 0xba, 0xc8, 0x3f, 0x0e,
	0xbf, 0xed, 0x24, 0xbb, 0x8d, 0xa2, 0xa2, 0xa5, 0x94, 0xc1, 0x50, 0x67, 0xb7, 0x63, 0xa8, 0xdd,
	0x88, 0x9e, 0xaf, 0x8c, 0x88, 0x3c, 0xc6, 0xf7, 0x99, 0xa2, 0x42, 0xf2, 0x9a, 0x29, 0xc6, 0x4b,
	0x48, 0x9d, 0x17, 0x2e, 0xf2, 0xef, 0x45, 0x23, 0xa6, 0x3e, 0xf5, 0x45, 0x12, 0xe2, 0x53, 0xa1,
	0x20, 0xe5, 0x0d, 0x93, 0x40, 0x41, 0x5f, 0x83, 0xa4, 0x96, 0xec, 0xa5, 0x21, 0x9b, 0x74, 0xcd,
	0xf3, 0xb6, 0x37, 0x37, 0xd4, 0xbf, 0x10, 0x3e, 0x13, 0x0a, 0xaa, 0x35, 0x94, 0x59, 0x02, 0xe7,
	0x95, 0x89, 0xa4, 0xda, 0xe1, 0xac, 0xfa, 0x0d, 0x8c, 0xa6, 0xbd, 0x28, 0xa3, 0x36, 0x32, 0x7b,
	0x19, 0xe0, 0xc9, 0x8a, 0x58, 0xd6, 0x2c, 0xfc, 0xbd, 0x36, 0xfe, 0xc6, 0x7d, 0x6b, 0xde, 0x58,
	0x77, 0x3f, 0xd1, 0x10, 0x60, 0xad, 0xbd, 0xf9, 0x9f, 0xd6, 0xd6, 0x74, 0x1a, 0x5f, 0x97, 0x87,
	0xe6, 0x2f, 0x3e, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xce, 0x9a, 0xe4, 0x1d, 0xeb, 0x05, 0x00,
	0x00,
}