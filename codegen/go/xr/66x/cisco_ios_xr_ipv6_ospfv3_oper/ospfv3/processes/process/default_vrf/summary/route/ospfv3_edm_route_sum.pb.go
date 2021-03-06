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
// source: ospfv3_edm_route_sum.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_summary_route

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

type Ospfv3EdmRouteSum_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmRouteSum_KEYS) Reset()         { *m = Ospfv3EdmRouteSum_KEYS{} }
func (m *Ospfv3EdmRouteSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmRouteSum_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmRouteSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_175224b5a8ab135f, []int{0}
}

func (m *Ospfv3EdmRouteSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmRouteSum_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmRouteSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmRouteSum_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmRouteSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmRouteSum_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmRouteSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmRouteSum_KEYS.Size(m)
}
func (m *Ospfv3EdmRouteSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmRouteSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmRouteSum_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmRouteSum_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

type Ospfv3EdmRouteSum struct {
	RouteId              string   `protobuf:"bytes,50,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	IntraAreaRoute       uint32   `protobuf:"varint,51,opt,name=intra_area_route,json=intraAreaRoute,proto3" json:"intra_area_route,omitempty"`
	InterAreaRoute       uint32   `protobuf:"varint,52,opt,name=inter_area_route,json=interAreaRoute,proto3" json:"inter_area_route,omitempty"`
	ExternOneRoute       uint32   `protobuf:"varint,53,opt,name=extern_one_route,json=externOneRoute,proto3" json:"extern_one_route,omitempty"`
	ExternTwoRoute       uint32   `protobuf:"varint,54,opt,name=extern_two_route,json=externTwoRoute,proto3" json:"extern_two_route,omitempty"`
	NssaOneRoute         uint32   `protobuf:"varint,55,opt,name=nssa_one_route,json=nssaOneRoute,proto3" json:"nssa_one_route,omitempty"`
	NssaTwoRoute         uint32   `protobuf:"varint,56,opt,name=nssa_two_route,json=nssaTwoRoute,proto3" json:"nssa_two_route,omitempty"`
	TotalSentRoute       uint32   `protobuf:"varint,57,opt,name=total_sent_route,json=totalSentRoute,proto3" json:"total_sent_route,omitempty"`
	RouteConnected       uint32   `protobuf:"varint,58,opt,name=route_connected,json=routeConnected,proto3" json:"route_connected,omitempty"`
	RedistributionRoute  uint32   `protobuf:"varint,59,opt,name=redistribution_route,json=redistributionRoute,proto3" json:"redistribution_route,omitempty"`
	TotalReceivedRoute   uint32   `protobuf:"varint,60,opt,name=total_received_route,json=totalReceivedRoute,proto3" json:"total_received_route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmRouteSum) Reset()         { *m = Ospfv3EdmRouteSum{} }
func (m *Ospfv3EdmRouteSum) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmRouteSum) ProtoMessage()    {}
func (*Ospfv3EdmRouteSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_175224b5a8ab135f, []int{1}
}

func (m *Ospfv3EdmRouteSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmRouteSum.Unmarshal(m, b)
}
func (m *Ospfv3EdmRouteSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmRouteSum.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmRouteSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmRouteSum.Merge(m, src)
}
func (m *Ospfv3EdmRouteSum) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmRouteSum.Size(m)
}
func (m *Ospfv3EdmRouteSum) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmRouteSum.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmRouteSum proto.InternalMessageInfo

func (m *Ospfv3EdmRouteSum) GetRouteId() string {
	if m != nil {
		return m.RouteId
	}
	return ""
}

func (m *Ospfv3EdmRouteSum) GetIntraAreaRoute() uint32 {
	if m != nil {
		return m.IntraAreaRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetInterAreaRoute() uint32 {
	if m != nil {
		return m.InterAreaRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetExternOneRoute() uint32 {
	if m != nil {
		return m.ExternOneRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetExternTwoRoute() uint32 {
	if m != nil {
		return m.ExternTwoRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetNssaOneRoute() uint32 {
	if m != nil {
		return m.NssaOneRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetNssaTwoRoute() uint32 {
	if m != nil {
		return m.NssaTwoRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetTotalSentRoute() uint32 {
	if m != nil {
		return m.TotalSentRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetRouteConnected() uint32 {
	if m != nil {
		return m.RouteConnected
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetRedistributionRoute() uint32 {
	if m != nil {
		return m.RedistributionRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetTotalReceivedRoute() uint32 {
	if m != nil {
		return m.TotalReceivedRoute
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3EdmRouteSum_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.summary.route.ospfv3_edm_route_sum_KEYS")
	proto.RegisterType((*Ospfv3EdmRouteSum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.summary.route.ospfv3_edm_route_sum")
}

func init() { proto.RegisterFile("ospfv3_edm_route_sum.proto", fileDescriptor_175224b5a8ab135f) }

var fileDescriptor_175224b5a8ab135f = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd2, 0x5d, 0x4b, 0xf3, 0x30,
	0x14, 0x07, 0x70, 0xc6, 0x03, 0x8f, 0x1a, 0xe7, 0x1c, 0x75, 0x17, 0x9d, 0x57, 0x73, 0x08, 0xf6,
	0xaa, 0xa8, 0xd3, 0xf9, 0x8a, 0x20, 0xe2, 0x85, 0x08, 0x2a, 0x9d, 0x37, 0x5e, 0x85, 0xac, 0x3d,
	0x83, 0xc0, 0x9a, 0x94, 0x93, 0xd3, 0x6e, 0x7e, 0x58, 0xbf, 0x8b, 0x2c, 0xc9, 0xde, 0x64, 0x77,
	0xc9, 0xff, 0xfc, 0xce, 0x3f, 0x63, 0x94, 0x1d, 0x6a, 0x53, 0x8c, 0xaa, 0x1e, 0x87, 0x2c, 0xe7,
	0xa8, 0x4b, 0x02, 0x6e, 0xca, 0x3c, 0x2e, 0x50, 0x93, 0x0e, 0x3e, 0x52, 0x69, 0x52, 0xcd, 0xa5,
	0x36, 0x7c, 0x8a, 0x5c, 0x16, 0x55, 0x9f, 0x7b, 0xad, 0x0b, 0xc0, 0xd8, 0x9d, 0x67, 0x36, 0x05,
	0x63, 0xc0, 0xcc, 0x4f, 0x71, 0x06, 0x23, 0x51, 0x8e, 0x89, 0x57, 0x38, 0x8a, 0x4d, 0x99, 0xe7,
	0x02, 0xbf, 0x63, 0xdb, 0xdd, 0x7d, 0x60, 0xed, 0x4d, 0xef, 0xf1, 0xd7, 0xe7, 0xaf, 0x41, 0x70,
	0xc4, 0xea, 0xbe, 0x81, 0x2b, 0x91, 0x43, 0x58, 0xeb, 0xd4, 0xa2, 0x9d, 0x64, 0xd7, 0x67, 0x6f,
	0x22, 0x87, 0xee, 0xcf, 0x3f, 0xd6, 0xda, 0x54, 0x10, 0xb4, 0xd9, 0xb6, 0xbb, 0xc8, 0x2c, 0x3c,
	0xb7, 0x7b, 0x5b, 0xf6, 0xfe, 0x92, 0x05, 0x11, 0x6b, 0x4a, 0x45, 0x28, 0xb8, 0x40, 0x10, 0x6e,
	0x25, 0xec, 0x75, 0x6a, 0xd1, 0x5e, 0xd2, 0xb0, 0xf9, 0x23, 0x82, 0x48, 0x66, 0xa9, 0x97, 0x80,
	0xab, 0xf2, 0x62, 0x21, 0x01, 0xd7, 0x24, 0x4c, 0x09, 0x50, 0x71, 0xad, 0xc0, 0xcb, 0x4b, 0x27,
	0x5d, 0xfe, 0xae, 0xe0, 0xaf, 0xa4, 0x89, 0xf6, 0xb2, 0xbf, 0x2a, 0x3f, 0x27, 0xda, 0xc9, 0x63,
	0xd6, 0x50, 0xc6, 0x88, 0x95, 0xc6, 0x2b, 0xeb, 0xea, 0xb3, 0x74, 0xd1, 0x37, 0x57, 0xcb, 0xb6,
	0xeb, 0xa5, 0x5a, 0x74, 0x45, 0xac, 0x49, 0x9a, 0xc4, 0x98, 0x1b, 0x50, 0xe4, 0xdd, 0x8d, 0x7b,
	0xd5, 0xe6, 0x03, 0x50, 0xe4, 0xe4, 0x09, 0xdb, 0x77, 0x7f, 0x5c, 0xaa, 0x95, 0x82, 0x94, 0x20,
	0x0b, 0x6f, 0x1d, 0xb4, 0xf1, 0xd3, 0x3c, 0x0d, 0xce, 0x58, 0x0b, 0x21, 0x93, 0x86, 0x50, 0x0e,
	0x4b, 0x92, 0x5a, 0xf9, 0xda, 0x3b, 0xab, 0x0f, 0xd6, 0x67, 0xae, 0xfb, 0x94, 0xb5, 0xdc, 0xaf,
	0x40, 0x48, 0x41, 0x56, 0x90, 0xf9, 0x95, 0x7b, 0xbb, 0x12, 0xd8, 0x59, 0xe2, 0x47, 0x76, 0x63,
	0xf8, 0xdf, 0x7e, 0x78, 0xbd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x21, 0xaf, 0xc9, 0x96,
	0x02, 0x00, 0x00,
}
