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
// source: l2vpn_pwhe_detail.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_active_pseudowire_headend_detail_interfaces_detail_interface

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

type L2VpnPwheDetail_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnPwheDetail_KEYS) Reset()         { *m = L2VpnPwheDetail_KEYS{} }
func (m *L2VpnPwheDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwheDetail_KEYS) ProtoMessage()    {}
func (*L2VpnPwheDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8dc59e93eff627, []int{0}
}

func (m *L2VpnPwheDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwheDetail_KEYS.Unmarshal(m, b)
}
func (m *L2VpnPwheDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwheDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnPwheDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwheDetail_KEYS.Merge(m, src)
}
func (m *L2VpnPwheDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwheDetail_KEYS.Size(m)
}
func (m *L2VpnPwheDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwheDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwheDetail_KEYS proto.InternalMessageInfo

func (m *L2VpnPwheDetail_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type L2VpnPwheMacAddrT struct {
	MacAddress           string   `protobuf:"bytes,1,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnPwheMacAddrT) Reset()         { *m = L2VpnPwheMacAddrT{} }
func (m *L2VpnPwheMacAddrT) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwheMacAddrT) ProtoMessage()    {}
func (*L2VpnPwheMacAddrT) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8dc59e93eff627, []int{1}
}

func (m *L2VpnPwheMacAddrT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwheMacAddrT.Unmarshal(m, b)
}
func (m *L2VpnPwheMacAddrT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwheMacAddrT.Marshal(b, m, deterministic)
}
func (m *L2VpnPwheMacAddrT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwheMacAddrT.Merge(m, src)
}
func (m *L2VpnPwheMacAddrT) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwheMacAddrT.Size(m)
}
func (m *L2VpnPwheMacAddrT) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwheMacAddrT.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwheMacAddrT proto.InternalMessageInfo

func (m *L2VpnPwheMacAddrT) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

type L2VpnL2AddrPwheDetail struct {
	InterfaceType        string             `protobuf:"bytes,1,opt,name=interface_type,json=interfaceType,proto3" json:"interface_type,omitempty"`
	MacAddress           *L2VpnPwheMacAddrT `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L2VpnL2AddrPwheDetail) Reset()         { *m = L2VpnL2AddrPwheDetail{} }
func (m *L2VpnL2AddrPwheDetail) String() string { return proto.CompactTextString(m) }
func (*L2VpnL2AddrPwheDetail) ProtoMessage()    {}
func (*L2VpnL2AddrPwheDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8dc59e93eff627, []int{2}
}

func (m *L2VpnL2AddrPwheDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnL2AddrPwheDetail.Unmarshal(m, b)
}
func (m *L2VpnL2AddrPwheDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnL2AddrPwheDetail.Marshal(b, m, deterministic)
}
func (m *L2VpnL2AddrPwheDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnL2AddrPwheDetail.Merge(m, src)
}
func (m *L2VpnL2AddrPwheDetail) XXX_Size() int {
	return xxx_messageInfo_L2VpnL2AddrPwheDetail.Size(m)
}
func (m *L2VpnL2AddrPwheDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnL2AddrPwheDetail.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnL2AddrPwheDetail proto.InternalMessageInfo

func (m *L2VpnL2AddrPwheDetail) GetInterfaceType() string {
	if m != nil {
		return m.InterfaceType
	}
	return ""
}

func (m *L2VpnL2AddrPwheDetail) GetMacAddress() *L2VpnPwheMacAddrT {
	if m != nil {
		return m.MacAddress
	}
	return nil
}

type L2VpnPwheGenericIflistIntf struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	BundleInterfaceName  string   `protobuf:"bytes,2,opt,name=bundle_interface_name,json=bundleInterfaceName,proto3" json:"bundle_interface_name,omitempty"`
	InterfaceState       string   `protobuf:"bytes,3,opt,name=interface_state,json=interfaceState,proto3" json:"interface_state,omitempty"`
	ReplicateState       string   `protobuf:"bytes,4,opt,name=replicate_state,json=replicateState,proto3" json:"replicate_state,omitempty"`
	Misconfigured        bool     `protobuf:"varint,5,opt,name=misconfigured,proto3" json:"misconfigured,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnPwheGenericIflistIntf) Reset()         { *m = L2VpnPwheGenericIflistIntf{} }
func (m *L2VpnPwheGenericIflistIntf) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwheGenericIflistIntf) ProtoMessage()    {}
func (*L2VpnPwheGenericIflistIntf) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8dc59e93eff627, []int{3}
}

func (m *L2VpnPwheGenericIflistIntf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwheGenericIflistIntf.Unmarshal(m, b)
}
func (m *L2VpnPwheGenericIflistIntf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwheGenericIflistIntf.Marshal(b, m, deterministic)
}
func (m *L2VpnPwheGenericIflistIntf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwheGenericIflistIntf.Merge(m, src)
}
func (m *L2VpnPwheGenericIflistIntf) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwheGenericIflistIntf.Size(m)
}
func (m *L2VpnPwheGenericIflistIntf) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwheGenericIflistIntf.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwheGenericIflistIntf proto.InternalMessageInfo

func (m *L2VpnPwheGenericIflistIntf) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *L2VpnPwheGenericIflistIntf) GetBundleInterfaceName() string {
	if m != nil {
		return m.BundleInterfaceName
	}
	return ""
}

func (m *L2VpnPwheGenericIflistIntf) GetInterfaceState() string {
	if m != nil {
		return m.InterfaceState
	}
	return ""
}

func (m *L2VpnPwheGenericIflistIntf) GetReplicateState() string {
	if m != nil {
		return m.ReplicateState
	}
	return ""
}

func (m *L2VpnPwheGenericIflistIntf) GetMisconfigured() bool {
	if m != nil {
		return m.Misconfigured
	}
	return false
}

type L2VpnPwheGenericIflist struct {
	GenericInterfaceListName string                        `protobuf:"bytes,1,opt,name=generic_interface_list_name,json=genericInterfaceListName,proto3" json:"generic_interface_list_name,omitempty"`
	Id                       uint32                        `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	InterfaceDetail          []*L2VpnPwheGenericIflistIntf `protobuf:"bytes,3,rep,name=interface_detail,json=interfaceDetail,proto3" json:"interface_detail,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                      `json:"-"`
	XXX_unrecognized         []byte                        `json:"-"`
	XXX_sizecache            int32                         `json:"-"`
}

func (m *L2VpnPwheGenericIflist) Reset()         { *m = L2VpnPwheGenericIflist{} }
func (m *L2VpnPwheGenericIflist) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwheGenericIflist) ProtoMessage()    {}
func (*L2VpnPwheGenericIflist) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8dc59e93eff627, []int{4}
}

func (m *L2VpnPwheGenericIflist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwheGenericIflist.Unmarshal(m, b)
}
func (m *L2VpnPwheGenericIflist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwheGenericIflist.Marshal(b, m, deterministic)
}
func (m *L2VpnPwheGenericIflist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwheGenericIflist.Merge(m, src)
}
func (m *L2VpnPwheGenericIflist) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwheGenericIflist.Size(m)
}
func (m *L2VpnPwheGenericIflist) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwheGenericIflist.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwheGenericIflist proto.InternalMessageInfo

func (m *L2VpnPwheGenericIflist) GetGenericInterfaceListName() string {
	if m != nil {
		return m.GenericInterfaceListName
	}
	return ""
}

func (m *L2VpnPwheGenericIflist) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *L2VpnPwheGenericIflist) GetInterfaceDetail() []*L2VpnPwheGenericIflistIntf {
	if m != nil {
		return m.InterfaceDetail
	}
	return nil
}

type L2VpnPwheDetail struct {
	InterfaceState       string                  `protobuf:"bytes,50,opt,name=interface_state,json=interfaceState,proto3" json:"interface_state,omitempty"`
	AdminState           string                  `protobuf:"bytes,51,opt,name=admin_state,json=adminState,proto3" json:"admin_state,omitempty"`
	Mtu                  uint32                  `protobuf:"varint,52,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Bandwidth            uint32                  `protobuf:"varint,53,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	Label                uint32                  `protobuf:"varint,54,opt,name=label,proto3" json:"label,omitempty"`
	L2Overhead           uint32                  `protobuf:"varint,55,opt,name=l2_overhead,json=l2Overhead,proto3" json:"l2_overhead,omitempty"`
	Vctype               uint32                  `protobuf:"varint,56,opt,name=vctype,proto3" json:"vctype,omitempty"`
	ControlWord          bool                    `protobuf:"varint,57,opt,name=control_word,json=controlWord,proto3" json:"control_word,omitempty"`
	L2Address            *L2VpnL2AddrPwheDetail  `protobuf:"bytes,58,opt,name=l2_address,json=l2Address,proto3" json:"l2_address,omitempty"`
	GenericInterfaceList *L2VpnPwheGenericIflist `protobuf:"bytes,59,opt,name=generic_interface_list,json=genericInterfaceList,proto3" json:"generic_interface_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *L2VpnPwheDetail) Reset()         { *m = L2VpnPwheDetail{} }
func (m *L2VpnPwheDetail) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwheDetail) ProtoMessage()    {}
func (*L2VpnPwheDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8dc59e93eff627, []int{5}
}

func (m *L2VpnPwheDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwheDetail.Unmarshal(m, b)
}
func (m *L2VpnPwheDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwheDetail.Marshal(b, m, deterministic)
}
func (m *L2VpnPwheDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwheDetail.Merge(m, src)
}
func (m *L2VpnPwheDetail) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwheDetail.Size(m)
}
func (m *L2VpnPwheDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwheDetail.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwheDetail proto.InternalMessageInfo

func (m *L2VpnPwheDetail) GetInterfaceState() string {
	if m != nil {
		return m.InterfaceState
	}
	return ""
}

func (m *L2VpnPwheDetail) GetAdminState() string {
	if m != nil {
		return m.AdminState
	}
	return ""
}

func (m *L2VpnPwheDetail) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *L2VpnPwheDetail) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *L2VpnPwheDetail) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *L2VpnPwheDetail) GetL2Overhead() uint32 {
	if m != nil {
		return m.L2Overhead
	}
	return 0
}

func (m *L2VpnPwheDetail) GetVctype() uint32 {
	if m != nil {
		return m.Vctype
	}
	return 0
}

func (m *L2VpnPwheDetail) GetControlWord() bool {
	if m != nil {
		return m.ControlWord
	}
	return false
}

func (m *L2VpnPwheDetail) GetL2Address() *L2VpnL2AddrPwheDetail {
	if m != nil {
		return m.L2Address
	}
	return nil
}

func (m *L2VpnPwheDetail) GetGenericInterfaceList() *L2VpnPwheGenericIflist {
	if m != nil {
		return m.GenericInterfaceList
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnPwheDetail_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pseudowire_headend.detail_interfaces.detail_interface.l2vpn_pwhe_detail_KEYS")
	proto.RegisterType((*L2VpnPwheMacAddrT)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pseudowire_headend.detail_interfaces.detail_interface.l2vpn_pwhe_mac_addr_t")
	proto.RegisterType((*L2VpnL2AddrPwheDetail)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pseudowire_headend.detail_interfaces.detail_interface.l2vpn_l2_addr_pwhe_detail")
	proto.RegisterType((*L2VpnPwheGenericIflistIntf)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pseudowire_headend.detail_interfaces.detail_interface.l2vpn_pwhe_generic_iflist_intf")
	proto.RegisterType((*L2VpnPwheGenericIflist)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pseudowire_headend.detail_interfaces.detail_interface.l2vpn_pwhe_generic_iflist")
	proto.RegisterType((*L2VpnPwheDetail)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pseudowire_headend.detail_interfaces.detail_interface.l2vpn_pwhe_detail")
}

func init() { proto.RegisterFile("l2vpn_pwhe_detail.proto", fileDescriptor_7a8dc59e93eff627) }

var fileDescriptor_7a8dc59e93eff627 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x66, 0x93, 0xb6, 0x98, 0x13, 0x53, 0xeb, 0xd8, 0xc6, 0x15, 0xc5, 0xc6, 0xa0, 0x98, 0xab,
	0xbd, 0xd8, 0xfa, 0x53, 0x15, 0x11, 0x41, 0x2f, 0x44, 0x51, 0xd8, 0x0a, 0x22, 0x08, 0xc3, 0x64,
	0xe7, 0xa4, 0x19, 0x98, 0xdd, 0x59, 0x66, 0x27, 0x89, 0x05, 0x1f, 0xa3, 0xf8, 0x0a, 0x5e, 0xf9,
	0x1e, 0xbe, 0x82, 0xcf, 0xe1, 0x0b, 0xc8, 0xce, 0x6c, 0x36, 0x89, 0xd9, 0x0b, 0x6f, 0x9a, 0xbb,
	0x9d, 0xef, 0x7c, 0xe7, 0x9c, 0x39, 0xe7, 0x9b, 0x2f, 0x81, 0xeb, 0x32, 0x9c, 0x66, 0x29, 0xcd,
	0x66, 0x63, 0xa4, 0x1c, 0x0d, 0x13, 0x32, 0xc8, 0xb4, 0x32, 0x8a, 0x7c, 0x89, 0x45, 0x1e, 0x2b,
	0x2a, 0x54, 0x4e, 0xbf, 0x6a, 0xea, 0x58, 0x2a, 0x43, 0x1d, 0xd8, 0xcf, 0x69, 0x18, 0xb0, 0xd8,
	0x88, 0x29, 0x06, 0x59, 0x8e, 0x13, 0xae, 0x66, 0x42, 0x23, 0x1d, 0x23, 0xe3, 0x98, 0xf2, 0xc0,
	0xd5, 0xa1, 0x22, 0x35, 0xa8, 0x47, 0x2c, 0xc6, 0x7c, 0x0d, 0xe9, 0xbf, 0x80, 0xee, 0x5a, 0x63,
	0xfa, 0xf6, 0xf5, 0xe7, 0x13, 0x72, 0x0f, 0x76, 0x2b, 0x1a, 0x4d, 0x59, 0x82, 0xbe, 0xd7, 0xf3,
	0x06, 0xad, 0xa8, 0x53, 0xa1, 0xef, 0x59, 0x82, 0xfd, 0x63, 0x38, 0x58, 0x2a, 0x90, 0xb0, 0x98,
	0x32, 0xce, 0x35, 0x35, 0xe4, 0x10, 0xda, 0xf3, 0x13, 0xe6, 0x79, 0x99, 0x0c, 0x09, 0x8b, 0x5f,
	0x3a, 0xa4, 0xff, 0xdb, 0x83, 0x1b, 0x2e, 0x55, 0x86, 0x2e, 0x69, 0xe9, 0x0e, 0xab, 0xed, 0xcd,
	0x59, 0xb6, 0xde, 0xfe, 0xe3, 0x59, 0x86, 0xe4, 0xdc, 0x5b, 0x6d, 0xd3, 0xe8, 0x79, 0x83, 0x76,
	0x98, 0x07, 0x17, 0xb9, 0xb4, 0xa0, 0x76, 0xe0, 0x95, 0xd9, 0xfe, 0x78, 0x70, 0x7b, 0x89, 0x75,
	0x8a, 0x29, 0x6a, 0x11, 0x53, 0x31, 0x92, 0x22, 0x37, 0x45, 0xa5, 0xd1, 0x7f, 0xee, 0x97, 0x84,
	0x70, 0x30, 0x9c, 0xa4, 0x5c, 0x22, 0xfd, 0x87, 0xdd, 0xb0, 0xec, 0x6b, 0x2e, 0xf8, 0x66, 0x25,
	0xe7, 0x3e, 0x5c, 0x59, 0x90, 0x73, 0xc3, 0x0c, 0xfa, 0x4d, 0xcb, 0x5e, 0x74, 0x3c, 0x29, 0xd0,
	0x82, 0xa8, 0x31, 0x93, 0x22, 0x66, 0x66, 0x4e, 0xdc, 0x72, 0xc4, 0x0a, 0x76, 0xc4, 0xbb, 0xd0,
	0x49, 0x8a, 0x8d, 0xa6, 0x23, 0x71, 0x3a, 0xd1, 0xc8, 0xfd, 0xed, 0x9e, 0x37, 0xb8, 0x14, 0xad,
	0x82, 0xfd, 0xf3, 0xc6, 0x5c, 0xd1, 0x9a, 0xa9, 0xc9, 0x73, 0xb8, 0x59, 0x21, 0xd5, 0xed, 0xec,
	0x3e, 0x96, 0xa6, 0xf7, 0x4b, 0x4a, 0x35, 0xd0, 0x3b, 0x91, 0x1b, 0x3b, 0xd4, 0x2e, 0x34, 0x04,
	0xb7, 0x53, 0x77, 0xa2, 0x86, 0xe0, 0xe4, 0x87, 0x07, 0x7b, 0x8b, 0x3a, 0x4e, 0x23, 0xbf, 0xd9,
	0x6b, 0x0e, 0xda, 0xe1, 0xb7, 0x8d, 0xc9, 0x5f, 0x23, 0x6c, 0xb4, 0xd8, 0xfd, 0x2b, 0x9b, 0xd9,
	0xff, 0xb5, 0x05, 0x57, 0xd7, 0x4c, 0x56, 0x27, 0x52, 0x58, 0x2b, 0xd2, 0x21, 0xb4, 0x19, 0x4f,
	0x44, 0x5a, 0x92, 0x8e, 0x9c, 0x91, 0x2c, 0xe4, 0x08, 0x7b, 0xd0, 0x4c, 0xcc, 0xc4, 0x7f, 0x60,
	0x57, 0x53, 0x7c, 0x92, 0x5b, 0xd0, 0x1a, 0xb2, 0x94, 0xcf, 0x04, 0x37, 0x63, 0xff, 0xa1, 0xc5,
	0x17, 0x00, 0xd9, 0x87, 0x6d, 0xc9, 0x86, 0x28, 0xfd, 0x47, 0x36, 0xe2, 0x0e, 0x45, 0x1b, 0x19,
	0x52, 0x35, 0x45, 0x5d, 0xac, 0xc2, 0x7f, 0x6c, 0x63, 0x20, 0xc3, 0x0f, 0x25, 0x42, 0xba, 0xb0,
	0x33, 0x8d, 0xad, 0x13, 0x8f, 0x6d, 0xac, 0x3c, 0x91, 0x3b, 0x70, 0x39, 0x56, 0xa9, 0xd1, 0x4a,
	0xd2, 0x99, 0xd2, 0xdc, 0x7f, 0x62, 0x9f, 0x46, 0xbb, 0xc4, 0x3e, 0x29, 0xcd, 0xc9, 0x77, 0x0f,
	0xa0, 0x34, 0x79, 0x61, 0xd2, 0xa7, 0xd6, 0xa4, 0xb3, 0x4d, 0xa8, 0x54, 0xf3, 0xd3, 0x12, 0xb5,
	0x64, 0x58, 0xfa, 0x94, 0xfc, 0xf4, 0xa0, 0x5b, 0xff, 0x28, 0xfd, 0x67, 0x9b, 0xbb, 0x64, 0xcd,
	0x53, 0x8a, 0xf6, 0xeb, 0x8c, 0x30, 0xdc, 0xb1, 0xff, 0x09, 0x47, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0xec, 0xd5, 0x22, 0x2e, 0x06, 0x00, 0x00,
}