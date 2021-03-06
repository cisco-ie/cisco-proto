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
// source: ospfv3_edm_border_router.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_border_routers_border_router

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

type Ospfv3EdmBorderRouter_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	BorderRouterId       string   `protobuf:"bytes,3,opt,name=border_router_id,json=borderRouterId,proto3" json:"border_router_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmBorderRouter_KEYS) Reset()         { *m = Ospfv3EdmBorderRouter_KEYS{} }
func (m *Ospfv3EdmBorderRouter_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmBorderRouter_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmBorderRouter_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_de51da64d25cb95f, []int{0}
}

func (m *Ospfv3EdmBorderRouter_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmBorderRouter_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmBorderRouter_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmBorderRouter_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmBorderRouter_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmBorderRouter_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmBorderRouter_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmBorderRouter_KEYS.Size(m)
}
func (m *Ospfv3EdmBorderRouter_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmBorderRouter_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmBorderRouter_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmBorderRouter_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmBorderRouter_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3EdmBorderRouter_KEYS) GetBorderRouterId() string {
	if m != nil {
		return m.BorderRouterId
	}
	return ""
}

type Ospfv3EdmBrPath struct {
	IsIntraAreaRouter       bool     `protobuf:"varint,1,opt,name=is_intra_area_router,json=isIntraAreaRouter,proto3" json:"is_intra_area_router,omitempty"`
	BorderRouterRouteMetric uint32   `protobuf:"varint,2,opt,name=border_router_route_metric,json=borderRouterRouteMetric,proto3" json:"border_router_route_metric,omitempty"`
	BorderRouterNextHop     string   `protobuf:"bytes,3,opt,name=border_router_next_hop,json=borderRouterNextHop,proto3" json:"border_router_next_hop,omitempty"`
	InterfaceName           string   `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	BorderRouterType        string   `protobuf:"bytes,5,opt,name=border_router_type,json=borderRouterType,proto3" json:"border_router_type,omitempty"`
	BorderRouterAreaId      string   `protobuf:"bytes,6,opt,name=border_router_area_id,json=borderRouterAreaId,proto3" json:"border_router_area_id,omitempty"`
	SpfVersion              uint32   `protobuf:"varint,7,opt,name=spf_version,json=spfVersion,proto3" json:"spf_version,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Ospfv3EdmBrPath) Reset()         { *m = Ospfv3EdmBrPath{} }
func (m *Ospfv3EdmBrPath) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmBrPath) ProtoMessage()    {}
func (*Ospfv3EdmBrPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_de51da64d25cb95f, []int{1}
}

func (m *Ospfv3EdmBrPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmBrPath.Unmarshal(m, b)
}
func (m *Ospfv3EdmBrPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmBrPath.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmBrPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmBrPath.Merge(m, src)
}
func (m *Ospfv3EdmBrPath) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmBrPath.Size(m)
}
func (m *Ospfv3EdmBrPath) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmBrPath.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmBrPath proto.InternalMessageInfo

func (m *Ospfv3EdmBrPath) GetIsIntraAreaRouter() bool {
	if m != nil {
		return m.IsIntraAreaRouter
	}
	return false
}

func (m *Ospfv3EdmBrPath) GetBorderRouterRouteMetric() uint32 {
	if m != nil {
		return m.BorderRouterRouteMetric
	}
	return 0
}

func (m *Ospfv3EdmBrPath) GetBorderRouterNextHop() string {
	if m != nil {
		return m.BorderRouterNextHop
	}
	return ""
}

func (m *Ospfv3EdmBrPath) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ospfv3EdmBrPath) GetBorderRouterType() string {
	if m != nil {
		return m.BorderRouterType
	}
	return ""
}

func (m *Ospfv3EdmBrPath) GetBorderRouterAreaId() string {
	if m != nil {
		return m.BorderRouterAreaId
	}
	return ""
}

func (m *Ospfv3EdmBrPath) GetSpfVersion() uint32 {
	if m != nil {
		return m.SpfVersion
	}
	return 0
}

type Ospfv3EdmBorderRouter struct {
	BorderRouterPath     []*Ospfv3EdmBrPath `protobuf:"bytes,50,rep,name=border_router_path,json=borderRouterPath,proto3" json:"border_router_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Ospfv3EdmBorderRouter) Reset()         { *m = Ospfv3EdmBorderRouter{} }
func (m *Ospfv3EdmBorderRouter) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmBorderRouter) ProtoMessage()    {}
func (*Ospfv3EdmBorderRouter) Descriptor() ([]byte, []int) {
	return fileDescriptor_de51da64d25cb95f, []int{2}
}

func (m *Ospfv3EdmBorderRouter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmBorderRouter.Unmarshal(m, b)
}
func (m *Ospfv3EdmBorderRouter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmBorderRouter.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmBorderRouter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmBorderRouter.Merge(m, src)
}
func (m *Ospfv3EdmBorderRouter) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmBorderRouter.Size(m)
}
func (m *Ospfv3EdmBorderRouter) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmBorderRouter.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmBorderRouter proto.InternalMessageInfo

func (m *Ospfv3EdmBorderRouter) GetBorderRouterPath() []*Ospfv3EdmBrPath {
	if m != nil {
		return m.BorderRouterPath
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3EdmBorderRouter_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.border_routers.border_router.ospfv3_edm_border_router_KEYS")
	proto.RegisterType((*Ospfv3EdmBrPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.border_routers.border_router.ospfv3_edm_br_path")
	proto.RegisterType((*Ospfv3EdmBorderRouter)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.border_routers.border_router.ospfv3_edm_border_router")
}

func init() { proto.RegisterFile("ospfv3_edm_border_router.proto", fileDescriptor_de51da64d25cb95f) }

var fileDescriptor_de51da64d25cb95f = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0xc9, 0x96, 0x64, 0xca, 0x12, 0x32, 0xed, 0xcf, 0x1b, 0x6c, 0xcb, 0x02, 0x83, 0x5c,
	0x0c, 0x8f, 0x25, 0xb0, 0x9b, 0x5d, 0xed, 0x62, 0xb0, 0x30, 0x16, 0x86, 0x5b, 0x0a, 0x85, 0x82,
	0x50, 0xec, 0x63, 0xa2, 0x0b, 0x5b, 0xe2, 0x48, 0x35, 0xc9, 0x13, 0xf4, 0x2d, 0xfa, 0x22, 0x7d,
	0x99, 0x3e, 0x4a, 0xf1, 0xb1, 0x5b, 0xac, 0x96, 0x5c, 0xf6, 0x46, 0x96, 0xcf, 0xf7, 0xa3, 0xef,
	0x1c, 0xd9, 0xec, 0xa3, 0xb6, 0x26, 0x2b, 0x97, 0x02, 0xd2, 0x5c, 0x6c, 0x34, 0xa6, 0x80, 0x02,
	0xf5, 0xb9, 0x03, 0x8c, 0x0c, 0x6a, 0xa7, 0xf9, 0x59, 0xa2, 0x6c, 0xa2, 0x85, 0xd2, 0x56, 0xec,
	0x50, 0x28, 0x53, 0xfe, 0x10, 0x8d, 0x42, 0x1b, 0xc0, 0xa8, 0xde, 0x57, 0xdc, 0x04, 0xac, 0x05,
	0x7b, 0xbb, 0x8b, 0x4a, 0xcc, 0x68, 0x89, 0x3c, 0x53, 0xeb, 0xbf, 0xce, 0x2e, 0x02, 0xf6, 0xe1,
	0x50, 0x00, 0xf1, 0xf7, 0xf7, 0xe9, 0x11, 0xff, 0xcc, 0x9e, 0x37, 0x96, 0xa2, 0x90, 0x39, 0x84,
	0xc1, 0x34, 0x98, 0x3f, 0x8b, 0x87, 0x4d, 0x6d, 0x2d, 0x73, 0xe0, 0xef, 0xd8, 0xa0, 0xc4, 0xac,
	0x86, 0x3b, 0x04, 0xf7, 0x4b, 0xcc, 0x08, 0x9a, 0xb3, 0x89, 0xef, 0xa9, 0xd2, 0xb0, 0x4b, 0x94,
	0x71, 0x5d, 0x8f, 0xa9, 0xbc, 0x4a, 0x67, 0xd7, 0x1d, 0xc6, 0xdb, 0x49, 0x50, 0x18, 0xe9, 0xb6,
	0xfc, 0x1b, 0x7b, 0xa5, 0xac, 0x50, 0x85, 0x43, 0x29, 0x24, 0x82, 0x6c, 0x7c, 0x28, 0xc6, 0x20,
	0x7e, 0xa1, 0xec, 0xaa, 0x82, 0x7e, 0x21, 0xc8, 0xda, 0x89, 0xff, 0x64, 0xef, 0xfd, 0x13, 0xe9,
	0x21, 0x72, 0x70, 0xa8, 0x12, 0x8a, 0x37, 0x8a, 0xdf, 0xb6, 0xcf, 0xa6, 0xf5, 0x1f, 0xc1, 0x7c,
	0xc9, 0xde, 0xf8, 0xe2, 0x02, 0x76, 0x4e, 0x6c, 0xb5, 0x69, 0x42, 0xbf, 0x6c, 0x0b, 0xd7, 0xb0,
	0x73, 0x7f, 0xb4, 0xe1, 0x5f, 0xd8, 0x58, 0x15, 0x0e, 0x30, 0x93, 0x09, 0xd4, 0x43, 0x78, 0x42,
	0xe4, 0xd1, 0x5d, 0x95, 0x46, 0xf1, 0x95, 0x71, 0xdf, 0xdb, 0xed, 0x0d, 0x84, 0x4f, 0x89, 0x3a,
	0x69, 0xfb, 0x1e, 0xef, 0x0d, 0xf0, 0xef, 0xec, 0xb5, 0xcf, 0xa6, 0xe6, 0x55, 0x1a, 0xf6, 0x48,
	0xc0, 0xdb, 0x82, 0xaa, 0xfb, 0x55, 0xca, 0x3f, 0xb1, 0xa1, 0x35, 0x99, 0x28, 0x01, 0xad, 0xd2,
	0x45, 0xd8, 0xa7, 0x56, 0x99, 0x35, 0xd9, 0x49, 0x5d, 0x99, 0x5d, 0x05, 0x2c, 0x3c, 0x74, 0xd9,
	0xfc, 0x32, 0xb8, 0x9f, 0xaf, 0x9a, 0x7f, 0xb8, 0x98, 0x76, 0xe7, 0xc3, 0x85, 0x89, 0x1e, 0xf3,
	0x2b, 0x8c, 0x1e, 0xde, 0xbb, 0x3f, 0x91, 0xff, 0xd2, 0x6d, 0x37, 0x3d, 0xfa, 0x1f, 0x96, 0x37,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x4d, 0x23, 0x2a, 0x31, 0x03, 0x00, 0x00,
}
