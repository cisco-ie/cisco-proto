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
// source: mpls_vpn_vrf.proto

package cisco_ios_xr_mpls_vpn_oper_l3vpn_invalid_vrfs_invalid_vrf

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

type MplsVpnVrf_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsVpnVrf_KEYS) Reset()         { *m = MplsVpnVrf_KEYS{} }
func (m *MplsVpnVrf_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsVpnVrf_KEYS) ProtoMessage()    {}
func (*MplsVpnVrf_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a6b37da191da58, []int{0}
}

func (m *MplsVpnVrf_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsVpnVrf_KEYS.Unmarshal(m, b)
}
func (m *MplsVpnVrf_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsVpnVrf_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsVpnVrf_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsVpnVrf_KEYS.Merge(m, src)
}
func (m *MplsVpnVrf_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsVpnVrf_KEYS.Size(m)
}
func (m *MplsVpnVrf_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsVpnVrf_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsVpnVrf_KEYS proto.InternalMessageInfo

func (m *MplsVpnVrf_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type MplsVpnInterfaces struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsVpnInterfaces) Reset()         { *m = MplsVpnInterfaces{} }
func (m *MplsVpnInterfaces) String() string { return proto.CompactTextString(m) }
func (*MplsVpnInterfaces) ProtoMessage()    {}
func (*MplsVpnInterfaces) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a6b37da191da58, []int{1}
}

func (m *MplsVpnInterfaces) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsVpnInterfaces.Unmarshal(m, b)
}
func (m *MplsVpnInterfaces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsVpnInterfaces.Marshal(b, m, deterministic)
}
func (m *MplsVpnInterfaces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsVpnInterfaces.Merge(m, src)
}
func (m *MplsVpnInterfaces) XXX_Size() int {
	return xxx_messageInfo_MplsVpnInterfaces.Size(m)
}
func (m *MplsVpnInterfaces) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsVpnInterfaces.DiscardUnknown(m)
}

var xxx_messageInfo_MplsVpnInterfaces proto.InternalMessageInfo

func (m *MplsVpnInterfaces) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type MplsVpnRt struct {
	RouteTargetType      string   `protobuf:"bytes,1,opt,name=route_target_type,json=routeTargetType,proto3" json:"route_target_type,omitempty"`
	RouteTargetValue     string   `protobuf:"bytes,2,opt,name=route_target_value,json=routeTargetValue,proto3" json:"route_target_value,omitempty"`
	AfName               string   `protobuf:"bytes,3,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,4,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsVpnRt) Reset()         { *m = MplsVpnRt{} }
func (m *MplsVpnRt) String() string { return proto.CompactTextString(m) }
func (*MplsVpnRt) ProtoMessage()    {}
func (*MplsVpnRt) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a6b37da191da58, []int{2}
}

func (m *MplsVpnRt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsVpnRt.Unmarshal(m, b)
}
func (m *MplsVpnRt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsVpnRt.Marshal(b, m, deterministic)
}
func (m *MplsVpnRt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsVpnRt.Merge(m, src)
}
func (m *MplsVpnRt) XXX_Size() int {
	return xxx_messageInfo_MplsVpnRt.Size(m)
}
func (m *MplsVpnRt) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsVpnRt.DiscardUnknown(m)
}

var xxx_messageInfo_MplsVpnRt proto.InternalMessageInfo

func (m *MplsVpnRt) GetRouteTargetType() string {
	if m != nil {
		return m.RouteTargetType
	}
	return ""
}

func (m *MplsVpnRt) GetRouteTargetValue() string {
	if m != nil {
		return m.RouteTargetValue
	}
	return ""
}

func (m *MplsVpnRt) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *MplsVpnRt) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

type MplsVpnAfiSafi struct {
	AfName               string       `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string       `protobuf:"bytes,2,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTarget          []*MplsVpnRt `protobuf:"bytes,3,rep,name=route_target,json=routeTarget,proto3" json:"route_target,omitempty"`
	ImportRoutePolicy    string       `protobuf:"bytes,4,opt,name=import_route_policy,json=importRoutePolicy,proto3" json:"import_route_policy,omitempty"`
	ExportRoutePolicy    string       `protobuf:"bytes,5,opt,name=export_route_policy,json=exportRoutePolicy,proto3" json:"export_route_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MplsVpnAfiSafi) Reset()         { *m = MplsVpnAfiSafi{} }
func (m *MplsVpnAfiSafi) String() string { return proto.CompactTextString(m) }
func (*MplsVpnAfiSafi) ProtoMessage()    {}
func (*MplsVpnAfiSafi) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a6b37da191da58, []int{3}
}

func (m *MplsVpnAfiSafi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsVpnAfiSafi.Unmarshal(m, b)
}
func (m *MplsVpnAfiSafi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsVpnAfiSafi.Marshal(b, m, deterministic)
}
func (m *MplsVpnAfiSafi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsVpnAfiSafi.Merge(m, src)
}
func (m *MplsVpnAfiSafi) XXX_Size() int {
	return xxx_messageInfo_MplsVpnAfiSafi.Size(m)
}
func (m *MplsVpnAfiSafi) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsVpnAfiSafi.DiscardUnknown(m)
}

var xxx_messageInfo_MplsVpnAfiSafi proto.InternalMessageInfo

func (m *MplsVpnAfiSafi) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *MplsVpnAfiSafi) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *MplsVpnAfiSafi) GetRouteTarget() []*MplsVpnRt {
	if m != nil {
		return m.RouteTarget
	}
	return nil
}

func (m *MplsVpnAfiSafi) GetImportRoutePolicy() string {
	if m != nil {
		return m.ImportRoutePolicy
	}
	return ""
}

func (m *MplsVpnAfiSafi) GetExportRoutePolicy() string {
	if m != nil {
		return m.ExportRoutePolicy
	}
	return ""
}

type MplsVpnVrf struct {
	VrfNameXr            string               `protobuf:"bytes,50,opt,name=vrf_name_xr,json=vrfNameXr,proto3" json:"vrf_name_xr,omitempty"`
	VrfDescription       string               `protobuf:"bytes,51,opt,name=vrf_description,json=vrfDescription,proto3" json:"vrf_description,omitempty"`
	RouteDistinguisher   string               `protobuf:"bytes,52,opt,name=route_distinguisher,json=routeDistinguisher,proto3" json:"route_distinguisher,omitempty"`
	Interface            []*MplsVpnInterfaces `protobuf:"bytes,53,rep,name=interface,proto3" json:"interface,omitempty"`
	Af                   []*MplsVpnAfiSafi    `protobuf:"bytes,54,rep,name=af,proto3" json:"af,omitempty"`
	IsBigVrf             bool                 `protobuf:"varint,55,opt,name=is_big_vrf,json=isBigVrf,proto3" json:"is_big_vrf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MplsVpnVrf) Reset()         { *m = MplsVpnVrf{} }
func (m *MplsVpnVrf) String() string { return proto.CompactTextString(m) }
func (*MplsVpnVrf) ProtoMessage()    {}
func (*MplsVpnVrf) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a6b37da191da58, []int{4}
}

func (m *MplsVpnVrf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsVpnVrf.Unmarshal(m, b)
}
func (m *MplsVpnVrf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsVpnVrf.Marshal(b, m, deterministic)
}
func (m *MplsVpnVrf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsVpnVrf.Merge(m, src)
}
func (m *MplsVpnVrf) XXX_Size() int {
	return xxx_messageInfo_MplsVpnVrf.Size(m)
}
func (m *MplsVpnVrf) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsVpnVrf.DiscardUnknown(m)
}

var xxx_messageInfo_MplsVpnVrf proto.InternalMessageInfo

func (m *MplsVpnVrf) GetVrfNameXr() string {
	if m != nil {
		return m.VrfNameXr
	}
	return ""
}

func (m *MplsVpnVrf) GetVrfDescription() string {
	if m != nil {
		return m.VrfDescription
	}
	return ""
}

func (m *MplsVpnVrf) GetRouteDistinguisher() string {
	if m != nil {
		return m.RouteDistinguisher
	}
	return ""
}

func (m *MplsVpnVrf) GetInterface() []*MplsVpnInterfaces {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *MplsVpnVrf) GetAf() []*MplsVpnAfiSafi {
	if m != nil {
		return m.Af
	}
	return nil
}

func (m *MplsVpnVrf) GetIsBigVrf() bool {
	if m != nil {
		return m.IsBigVrf
	}
	return false
}

func init() {
	proto.RegisterType((*MplsVpnVrf_KEYS)(nil), "cisco_ios_xr_mpls_vpn_oper.l3vpn.invalid_vrfs.invalid_vrf.mpls_vpn_vrf_KEYS")
	proto.RegisterType((*MplsVpnInterfaces)(nil), "cisco_ios_xr_mpls_vpn_oper.l3vpn.invalid_vrfs.invalid_vrf.mpls_vpn_interfaces")
	proto.RegisterType((*MplsVpnRt)(nil), "cisco_ios_xr_mpls_vpn_oper.l3vpn.invalid_vrfs.invalid_vrf.mpls_vpn_rt")
	proto.RegisterType((*MplsVpnAfiSafi)(nil), "cisco_ios_xr_mpls_vpn_oper.l3vpn.invalid_vrfs.invalid_vrf.mpls_vpn_afi_safi")
	proto.RegisterType((*MplsVpnVrf)(nil), "cisco_ios_xr_mpls_vpn_oper.l3vpn.invalid_vrfs.invalid_vrf.mpls_vpn_vrf")
}

func init() { proto.RegisterFile("mpls_vpn_vrf.proto", fileDescriptor_99a6b37da191da58) }

var fileDescriptor_99a6b37da191da58 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdf, 0x8b, 0x13, 0x31,
	0x10, 0xc7, 0xd9, 0xad, 0xde, 0xb5, 0xd3, 0xf3, 0xce, 0xa6, 0x0f, 0x46, 0x10, 0x29, 0x05, 0xb1,
	0x88, 0xac, 0x70, 0xf5, 0x07, 0x82, 0x4f, 0x72, 0xfa, 0xa2, 0x1c, 0xb2, 0x1e, 0x87, 0x82, 0x10,
	0x72, 0xdb, 0xa4, 0x0e, 0x6c, 0x37, 0x21, 0x49, 0x97, 0xf6, 0xdf, 0xf0, 0xd5, 0xff, 0xd2, 0xbf,
	0x40, 0x36, 0xe9, 0x6e, 0x73, 0x88, 0x2f, 0xc7, 0xbd, 0x6d, 0xe6, 0xfb, 0x99, 0x99, 0xcc, 0xe4,
	0xbb, 0x40, 0x56, 0xba, 0xb4, 0xac, 0xd6, 0x15, 0xab, 0x8d, 0xcc, 0xb4, 0x51, 0x4e, 0x91, 0xb7,
	0x05, 0xda, 0x42, 0x31, 0x54, 0x96, 0x6d, 0x0c, 0xeb, 0x00, 0xa5, 0x85, 0xc9, 0xca, 0x79, 0xad,
	0xab, 0x0c, 0xab, 0x9a, 0x97, 0xb8, 0x68, 0x72, 0x6c, 0x7c, 0x98, 0x66, 0x30, 0x8a, 0x0b, 0xb2,
	0x4f, 0x1f, 0xbe, 0x7f, 0x25, 0x0f, 0xa1, 0xdf, 0x7c, 0x57, 0x7c, 0x25, 0x68, 0x32, 0x49, 0x66,
	0x83, 0xfc, 0xb0, 0x36, 0xf2, 0x9c, 0xaf, 0xc4, 0xf4, 0x1d, 0x8c, 0x3b, 0x1e, 0x2b, 0x27, 0x8c,
	0xe4, 0x85, 0xb0, 0xe4, 0x09, 0x1c, 0x77, 0xa7, 0x38, 0xef, 0x5e, 0x17, 0xf5, 0xd9, 0xbf, 0x13,
	0x18, 0x76, 0xe9, 0xc6, 0x91, 0x67, 0x30, 0x32, 0x6a, 0xed, 0x04, 0x73, 0xdc, 0x2c, 0x85, 0x63,
	0x6e, 0xab, 0xdb, 0xcc, 0x13, 0x2f, 0x5c, 0xf8, 0xf8, 0xc5, 0x56, 0x0b, 0xf2, 0x1c, 0xc8, 0x35,
	0xb6, 0xe6, 0xe5, 0x5a, 0xd0, 0xd4, 0xc3, 0xf7, 0x23, 0xf8, 0xb2, 0x89, 0x93, 0x07, 0x70, 0xc8,
	0x77, 0x13, 0xf4, 0x3c, 0x72, 0xc0, 0xfd, 0x00, 0xcd, 0x6c, 0xb6, 0x55, 0xee, 0x84, 0xd9, 0x6c,
	0x90, 0xa6, 0xbf, 0xd2, 0x68, 0x19, 0x5c, 0x22, 0xb3, 0x5c, 0x62, 0x5c, 0x29, 0xf9, 0x6f, 0xa5,
	0xf4, 0x5a, 0x25, 0x82, 0x70, 0x14, 0xdf, 0x95, 0xf6, 0x26, 0xbd, 0xd9, 0xf0, 0xf4, 0x63, 0x76,
	0xe3, 0x77, 0xca, 0xa2, 0xad, 0xe5, 0xc3, 0x68, 0x5a, 0x92, 0xc1, 0x18, 0x57, 0x5a, 0x19, 0xc7,
	0x42, 0x47, 0xad, 0x4a, 0x2c, 0xb6, 0xbb, 0xd1, 0x46, 0x41, 0xca, 0x1b, 0xe5, 0x8b, 0x17, 0x1a,
	0x5e, 0x6c, 0xfe, 0xe5, 0xef, 0x06, 0x3e, 0x48, 0x11, 0x3f, 0xfd, 0x93, 0xc2, 0x51, 0xec, 0x10,
	0xf2, 0x18, 0x86, 0xad, 0x39, 0xd8, 0xc6, 0xd0, 0x53, 0x9f, 0x38, 0xd8, 0xf9, 0xe3, 0x9b, 0x21,
	0x4f, 0xe1, 0xa4, 0xd1, 0x17, 0xc2, 0x16, 0x06, 0xb5, 0x43, 0x55, 0xd1, 0xb9, 0x67, 0x8e, 0x6b,
	0x23, 0xcf, 0xf6, 0x51, 0xf2, 0x02, 0xc6, 0xe1, 0x0a, 0x0b, 0xb4, 0x0e, 0xab, 0xe5, 0x1a, 0xed,
	0x4f, 0x61, 0xe8, 0x4b, 0x0f, 0x87, 0xb7, 0x3e, 0x8b, 0x15, 0x52, 0xc2, 0xa0, 0xb3, 0x13, 0x7d,
	0xe5, 0x57, 0x7a, 0x7e, 0x1b, 0x2b, 0xdd, 0xfb, 0x38, 0xdf, 0x37, 0x20, 0x3f, 0x20, 0xe5, 0x92,
	0xbe, 0xf6, 0x6d, 0x3e, 0xdf, 0x46, 0x9b, 0xd6, 0x51, 0x79, 0xca, 0x25, 0x79, 0x04, 0x80, 0x96,
	0x5d, 0xe1, 0xb2, 0x61, 0xe8, 0x9b, 0x49, 0x32, 0xeb, 0xe7, 0x7d, 0xb4, 0xef, 0x71, 0x79, 0x69,
	0xe4, 0xd5, 0x81, 0xff, 0xaf, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x71, 0x3c, 0xab, 0x3a,
	0xed, 0x03, 0x00, 0x00,
}