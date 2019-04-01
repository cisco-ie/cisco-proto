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
// source: rip_show_route_bd.proto

package cisco_ios_xr_ip_rip_oper_rip_default_vrf_routes_route

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

type RipShowRouteBd_KEYS struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RipShowRouteBd_KEYS) Reset()         { *m = RipShowRouteBd_KEYS{} }
func (m *RipShowRouteBd_KEYS) String() string { return proto.CompactTextString(m) }
func (*RipShowRouteBd_KEYS) ProtoMessage()    {}
func (*RipShowRouteBd_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4d66ed104ef7a31, []int{0}
}

func (m *RipShowRouteBd_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RipShowRouteBd_KEYS.Unmarshal(m, b)
}
func (m *RipShowRouteBd_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RipShowRouteBd_KEYS.Marshal(b, m, deterministic)
}
func (m *RipShowRouteBd_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RipShowRouteBd_KEYS.Merge(m, src)
}
func (m *RipShowRouteBd_KEYS) XXX_Size() int {
	return xxx_messageInfo_RipShowRouteBd_KEYS.Size(m)
}
func (m *RipShowRouteBd_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RipShowRouteBd_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RipShowRouteBd_KEYS proto.InternalMessageInfo

func (m *RipShowRouteBd_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RipShowRouteBd_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type RipPathBd struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	NextHopAddress       string   `protobuf:"bytes,2,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	Metric               uint32   `protobuf:"varint,3,opt,name=metric,proto3" json:"metric,omitempty"`
	Tag                  uint32   `protobuf:"varint,4,opt,name=tag,proto3" json:"tag,omitempty"`
	Interface            string   `protobuf:"bytes,5,opt,name=interface,proto3" json:"interface,omitempty"`
	Uptime               uint32   `protobuf:"varint,6,opt,name=uptime,proto3" json:"uptime,omitempty"`
	IsPermanent          bool     `protobuf:"varint,7,opt,name=is_permanent,json=isPermanent,proto3" json:"is_permanent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RipPathBd) Reset()         { *m = RipPathBd{} }
func (m *RipPathBd) String() string { return proto.CompactTextString(m) }
func (*RipPathBd) ProtoMessage()    {}
func (*RipPathBd) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4d66ed104ef7a31, []int{1}
}

func (m *RipPathBd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RipPathBd.Unmarshal(m, b)
}
func (m *RipPathBd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RipPathBd.Marshal(b, m, deterministic)
}
func (m *RipPathBd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RipPathBd.Merge(m, src)
}
func (m *RipPathBd) XXX_Size() int {
	return xxx_messageInfo_RipPathBd.Size(m)
}
func (m *RipPathBd) XXX_DiscardUnknown() {
	xxx_messageInfo_RipPathBd.DiscardUnknown(m)
}

var xxx_messageInfo_RipPathBd proto.InternalMessageInfo

func (m *RipPathBd) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *RipPathBd) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *RipPathBd) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *RipPathBd) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *RipPathBd) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *RipPathBd) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *RipPathBd) GetIsPermanent() bool {
	if m != nil {
		return m.IsPermanent
	}
	return false
}

type RipShowRouteBd struct {
	DestinationAddress   string       `protobuf:"bytes,50,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	PrefixLengthXr       uint32       `protobuf:"varint,51,opt,name=prefix_length_xr,json=prefixLengthXr,proto3" json:"prefix_length_xr,omitempty"`
	Distance             uint32       `protobuf:"varint,52,opt,name=distance,proto3" json:"distance,omitempty"`
	BgpCount             uint32       `protobuf:"varint,53,opt,name=bgp_count,json=bgpCount,proto3" json:"bgp_count,omitempty"`
	RouteType            uint32       `protobuf:"varint,54,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	RouteSummary         bool         `protobuf:"varint,55,opt,name=route_summary,json=routeSummary,proto3" json:"route_summary,omitempty"`
	RouteTag             uint32       `protobuf:"varint,56,opt,name=route_tag,json=routeTag,proto3" json:"route_tag,omitempty"`
	Version              uint32       `protobuf:"varint,57,opt,name=version,proto3" json:"version,omitempty"`
	Attributes           uint32       `protobuf:"varint,58,opt,name=attributes,proto3" json:"attributes,omitempty"`
	Active               bool         `protobuf:"varint,59,opt,name=active,proto3" json:"active,omitempty"`
	Paths                []*RipPathBd `protobuf:"bytes,60,rep,name=paths,proto3" json:"paths,omitempty"`
	PathOrigin           string       `protobuf:"bytes,61,opt,name=path_origin,json=pathOrigin,proto3" json:"path_origin,omitempty"`
	HoldDown             bool         `protobuf:"varint,62,opt,name=hold_down,json=holdDown,proto3" json:"hold_down,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RipShowRouteBd) Reset()         { *m = RipShowRouteBd{} }
func (m *RipShowRouteBd) String() string { return proto.CompactTextString(m) }
func (*RipShowRouteBd) ProtoMessage()    {}
func (*RipShowRouteBd) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4d66ed104ef7a31, []int{2}
}

func (m *RipShowRouteBd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RipShowRouteBd.Unmarshal(m, b)
}
func (m *RipShowRouteBd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RipShowRouteBd.Marshal(b, m, deterministic)
}
func (m *RipShowRouteBd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RipShowRouteBd.Merge(m, src)
}
func (m *RipShowRouteBd) XXX_Size() int {
	return xxx_messageInfo_RipShowRouteBd.Size(m)
}
func (m *RipShowRouteBd) XXX_DiscardUnknown() {
	xxx_messageInfo_RipShowRouteBd.DiscardUnknown(m)
}

var xxx_messageInfo_RipShowRouteBd proto.InternalMessageInfo

func (m *RipShowRouteBd) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RipShowRouteBd) GetPrefixLengthXr() uint32 {
	if m != nil {
		return m.PrefixLengthXr
	}
	return 0
}

func (m *RipShowRouteBd) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *RipShowRouteBd) GetBgpCount() uint32 {
	if m != nil {
		return m.BgpCount
	}
	return 0
}

func (m *RipShowRouteBd) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *RipShowRouteBd) GetRouteSummary() bool {
	if m != nil {
		return m.RouteSummary
	}
	return false
}

func (m *RipShowRouteBd) GetRouteTag() uint32 {
	if m != nil {
		return m.RouteTag
	}
	return 0
}

func (m *RipShowRouteBd) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RipShowRouteBd) GetAttributes() uint32 {
	if m != nil {
		return m.Attributes
	}
	return 0
}

func (m *RipShowRouteBd) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *RipShowRouteBd) GetPaths() []*RipPathBd {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *RipShowRouteBd) GetPathOrigin() string {
	if m != nil {
		return m.PathOrigin
	}
	return ""
}

func (m *RipShowRouteBd) GetHoldDown() bool {
	if m != nil {
		return m.HoldDown
	}
	return false
}

func init() {
	proto.RegisterType((*RipShowRouteBd_KEYS)(nil), "cisco_ios_xr_ip_rip_oper.rip.default_vrf.routes.route.rip_show_route_bd_KEYS")
	proto.RegisterType((*RipPathBd)(nil), "cisco_ios_xr_ip_rip_oper.rip.default_vrf.routes.route.rip_path_bd")
	proto.RegisterType((*RipShowRouteBd)(nil), "cisco_ios_xr_ip_rip_oper.rip.default_vrf.routes.route.rip_show_route_bd")
}

func init() { proto.RegisterFile("rip_show_route_bd.proto", fileDescriptor_f4d66ed104ef7a31) }

var fileDescriptor_f4d66ed104ef7a31 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x5b, 0x8b, 0x13, 0x31,
	0x14, 0xc7, 0xe9, 0xd6, 0xed, 0xb6, 0xa7, 0x17, 0xd6, 0x08, 0x35, 0x78, 0xad, 0x15, 0xa1, 0x4f,
	0x23, 0xec, 0xba, 0xde, 0x15, 0xbc, 0x81, 0xa0, 0xa0, 0x74, 0x15, 0xd6, 0xa7, 0x90, 0x99, 0x49,
	0xa7, 0x81, 0x36, 0x09, 0x49, 0xa6, 0x97, 0x8f, 0xeb, 0x57, 0xf0, 0x13, 0x48, 0x4e, 0xa6, 0x75,
	0x64, 0xdf, 0xf6, 0xa9, 0xf9, 0xff, 0xce, 0xe9, 0xff, 0xe4, 0x9c, 0x9c, 0x81, 0x9b, 0x56, 0x1a,
	0xe6, 0xe6, 0x7a, 0xcd, 0xac, 0x2e, 0xbd, 0x60, 0x69, 0x9e, 0x18, 0xab, 0xbd, 0x26, 0x67, 0x99,
	0x74, 0x99, 0x66, 0x52, 0x3b, 0xb6, 0xb1, 0x4c, 0x1a, 0x16, 0x12, 0xb5, 0x11, 0x36, 0xb1, 0xd2,
	0x24, 0xb9, 0x98, 0xf1, 0x72, 0xe1, 0xd9, 0xca, 0xce, 0x12, 0xfc, 0xa3, 0x8b, 0x3f, 0xe3, 0x9f,
	0x30, 0xbc, 0xe4, 0xc8, 0xbe, 0x7c, 0xfa, 0x75, 0x4e, 0x86, 0xd0, 0x32, 0x56, 0xcc, 0xe4, 0x86,
	0x36, 0x46, 0x8d, 0x49, 0x67, 0x5a, 0x29, 0xf2, 0x10, 0xfa, 0xf1, 0xc4, 0x16, 0x42, 0x15, 0x7e,
	0x4e, 0x0f, 0x46, 0x8d, 0x49, 0x7f, 0xda, 0x8b, 0xf0, 0x2b, 0xb2, 0xf1, 0xef, 0x06, 0x74, 0x83,
	0xaf, 0xe1, 0x7e, 0xce, 0xd2, 0x9c, 0x3c, 0x82, 0x81, 0xd3, 0xa5, 0xcd, 0x04, 0xe3, 0x79, 0x6e,
	0x85, 0x73, 0x95, 0x69, 0x3f, 0xd2, 0x77, 0x11, 0x92, 0x09, 0x1c, 0x2b, 0xb1, 0xf1, 0x6c, 0xae,
	0xcd, 0x3e, 0xf1, 0x00, 0x13, 0x07, 0x81, 0x7f, 0xd6, 0x66, 0x97, 0x39, 0x84, 0xd6, 0x52, 0x78,
	0x2b, 0x33, 0xda, 0xc4, 0xf2, 0x95, 0x22, 0xc7, 0xd0, 0xf4, 0xbc, 0xa0, 0xd7, 0x10, 0x86, 0x23,
	0xb9, 0x03, 0x1d, 0xa9, 0xbc, 0xb0, 0x33, 0x9e, 0x09, 0x7a, 0x88, 0x66, 0xff, 0x40, 0xf0, 0x29,
	0x8d, 0x97, 0x4b, 0x41, 0x5b, 0xd1, 0x27, 0x2a, 0xf2, 0x00, 0x7a, 0xd2, 0x31, 0x23, 0xec, 0x92,
	0x2b, 0xa1, 0x3c, 0x3d, 0x1a, 0x35, 0x26, 0xed, 0x69, 0x57, 0xba, 0xef, 0x3b, 0x34, 0xfe, 0xd3,
	0x84, 0xeb, 0x97, 0x66, 0x47, 0x1e, 0xc3, 0x8d, 0x5c, 0x38, 0x2f, 0x15, 0xf7, 0x52, 0xab, 0x7d,
	0x17, 0x27, 0x58, 0x98, 0xd4, 0x42, 0xb5, 0x9e, 0xff, 0x9b, 0x27, 0xdb, 0x58, 0x7a, 0x8a, 0x77,
	0x19, 0xd4, 0x47, 0x7a, 0x61, 0xc9, 0x2d, 0x68, 0xe7, 0xd2, 0x79, 0xae, 0x32, 0x41, 0x9f, 0x60,
	0xc6, 0x5e, 0x93, 0xdb, 0xd0, 0x49, 0x0b, 0xc3, 0x32, 0x5d, 0x2a, 0x4f, 0xcf, 0x62, 0x30, 0x2d,
	0xcc, 0x87, 0xa0, 0xc9, 0x5d, 0x80, 0x78, 0x3f, 0xbf, 0x35, 0x82, 0x3e, 0xc5, 0x68, 0x07, 0xc9,
	0x8f, 0xad, 0x11, 0xe1, 0x45, 0x63, 0xd8, 0x95, 0xcb, 0x25, 0xb7, 0x5b, 0xfa, 0x0c, 0x9b, 0xed,
	0x21, 0x3c, 0x8f, 0x2c, 0x14, 0xa8, 0x3c, 0x78, 0x41, 0x9f, 0xc7, 0x02, 0xd1, 0x82, 0x17, 0x84,
	0xc2, 0xd1, 0x4a, 0x58, 0x27, 0xb5, 0xa2, 0x2f, 0x30, 0xb4, 0x93, 0xe4, 0x1e, 0x00, 0xf7, 0xde,
	0xca, 0x34, 0xec, 0x1c, 0x7d, 0x89, 0xc1, 0x1a, 0x09, 0xf3, 0xe7, 0x99, 0x97, 0x2b, 0x41, 0x5f,
	0x61, 0xd1, 0x4a, 0x91, 0x0b, 0x38, 0x0c, 0xbb, 0xe3, 0xe8, 0xeb, 0x51, 0x73, 0xd2, 0x3d, 0x79,
	0x9f, 0x5c, 0x69, 0xbd, 0x93, 0xda, 0x0e, 0x4e, 0xa3, 0x21, 0xb9, 0x0f, 0x5d, 0x24, 0xda, 0xca,
	0x42, 0x2a, 0xfa, 0x06, 0x1f, 0x06, 0x02, 0xfa, 0x86, 0x24, 0x74, 0x3a, 0xd7, 0x8b, 0x9c, 0xe5,
	0x7a, 0xad, 0xe8, 0x5b, 0xbc, 0x55, 0x3b, 0x80, 0x8f, 0x7a, 0xad, 0xd2, 0x16, 0x7e, 0x6d, 0xa7,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xbd, 0xb0, 0x46, 0x88, 0x03, 0x00, 0x00,
}
