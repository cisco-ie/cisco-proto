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
// source: track_info_brief___entry.proto

package cisco_ios_xr_manageability_object_tracking_oper_object_tracking_track_type_ipv4_route_brief

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

type TrackInfoBrief__Entry_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackInfoBrief__Entry_KEYS) Reset()         { *m = TrackInfoBrief__Entry_KEYS{} }
func (m *TrackInfoBrief__Entry_KEYS) String() string { return proto.CompactTextString(m) }
func (*TrackInfoBrief__Entry_KEYS) ProtoMessage()    {}
func (*TrackInfoBrief__Entry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_03fe192acc1772f2, []int{0}
}

func (m *TrackInfoBrief__Entry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackInfoBrief__Entry_KEYS.Unmarshal(m, b)
}
func (m *TrackInfoBrief__Entry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackInfoBrief__Entry_KEYS.Marshal(b, m, deterministic)
}
func (m *TrackInfoBrief__Entry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackInfoBrief__Entry_KEYS.Merge(m, src)
}
func (m *TrackInfoBrief__Entry_KEYS) XXX_Size() int {
	return xxx_messageInfo_TrackInfoBrief__Entry_KEYS.Size(m)
}
func (m *TrackInfoBrief__Entry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackInfoBrief__Entry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TrackInfoBrief__Entry_KEYS proto.InternalMessageInfo

type IntfTrackInfo__ struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntfTrackInfo__) Reset()         { *m = IntfTrackInfo__{} }
func (m *IntfTrackInfo__) String() string { return proto.CompactTextString(m) }
func (*IntfTrackInfo__) ProtoMessage()    {}
func (*IntfTrackInfo__) Descriptor() ([]byte, []int) {
	return fileDescriptor_03fe192acc1772f2, []int{1}
}

func (m *IntfTrackInfo__) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntfTrackInfo__.Unmarshal(m, b)
}
func (m *IntfTrackInfo__) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntfTrackInfo__.Marshal(b, m, deterministic)
}
func (m *IntfTrackInfo__) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntfTrackInfo__.Merge(m, src)
}
func (m *IntfTrackInfo__) XXX_Size() int {
	return xxx_messageInfo_IntfTrackInfo__.Size(m)
}
func (m *IntfTrackInfo__) XXX_DiscardUnknown() {
	xxx_messageInfo_IntfTrackInfo__.DiscardUnknown(m)
}

var xxx_messageInfo_IntfTrackInfo__ proto.InternalMessageInfo

func (m *IntfTrackInfo__) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type RouteTrackInfo__ struct {
	Prefix               uint32   `protobuf:"varint,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	Vrf                  string   `protobuf:"bytes,3,opt,name=vrf,proto3" json:"vrf,omitempty"`
	NextHop              string   `protobuf:"bytes,4,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteTrackInfo__) Reset()         { *m = RouteTrackInfo__{} }
func (m *RouteTrackInfo__) String() string { return proto.CompactTextString(m) }
func (*RouteTrackInfo__) ProtoMessage()    {}
func (*RouteTrackInfo__) Descriptor() ([]byte, []int) {
	return fileDescriptor_03fe192acc1772f2, []int{2}
}

func (m *RouteTrackInfo__) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTrackInfo__.Unmarshal(m, b)
}
func (m *RouteTrackInfo__) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTrackInfo__.Marshal(b, m, deterministic)
}
func (m *RouteTrackInfo__) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTrackInfo__.Merge(m, src)
}
func (m *RouteTrackInfo__) XXX_Size() int {
	return xxx_messageInfo_RouteTrackInfo__.Size(m)
}
func (m *RouteTrackInfo__) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTrackInfo__.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTrackInfo__ proto.InternalMessageInfo

func (m *RouteTrackInfo__) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

func (m *RouteTrackInfo__) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *RouteTrackInfo__) GetVrf() string {
	if m != nil {
		return m.Vrf
	}
	return ""
}

func (m *RouteTrackInfo__) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

type RtrTrackInfo__ struct {
	IpslaOpId            uint32   `protobuf:"varint,1,opt,name=ipsla_op_id,json=ipslaOpId,proto3" json:"ipsla_op_id,omitempty"`
	Rtt                  uint32   `protobuf:"varint,2,opt,name=rtt,proto3" json:"rtt,omitempty"`
	ReturnCode           uint32   `protobuf:"varint,3,opt,name=return_code,json=returnCode,proto3" json:"return_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RtrTrackInfo__) Reset()         { *m = RtrTrackInfo__{} }
func (m *RtrTrackInfo__) String() string { return proto.CompactTextString(m) }
func (*RtrTrackInfo__) ProtoMessage()    {}
func (*RtrTrackInfo__) Descriptor() ([]byte, []int) {
	return fileDescriptor_03fe192acc1772f2, []int{3}
}

func (m *RtrTrackInfo__) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RtrTrackInfo__.Unmarshal(m, b)
}
func (m *RtrTrackInfo__) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RtrTrackInfo__.Marshal(b, m, deterministic)
}
func (m *RtrTrackInfo__) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RtrTrackInfo__.Merge(m, src)
}
func (m *RtrTrackInfo__) XXX_Size() int {
	return xxx_messageInfo_RtrTrackInfo__.Size(m)
}
func (m *RtrTrackInfo__) XXX_DiscardUnknown() {
	xxx_messageInfo_RtrTrackInfo__.DiscardUnknown(m)
}

var xxx_messageInfo_RtrTrackInfo__ proto.InternalMessageInfo

func (m *RtrTrackInfo__) GetIpslaOpId() uint32 {
	if m != nil {
		return m.IpslaOpId
	}
	return 0
}

func (m *RtrTrackInfo__) GetRtt() uint32 {
	if m != nil {
		return m.Rtt
	}
	return 0
}

func (m *RtrTrackInfo__) GetReturnCode() uint32 {
	if m != nil {
		return m.ReturnCode
	}
	return 0
}

type BfdrtrTrackInfo__ struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	DestinationAddress   uint32   `protobuf:"varint,2,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	Rate                 uint32   `protobuf:"varint,3,opt,name=rate,proto3" json:"rate,omitempty"`
	DebounceCount        uint32   `protobuf:"varint,4,opt,name=debounce_count,json=debounceCount,proto3" json:"debounce_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdrtrTrackInfo__) Reset()         { *m = BfdrtrTrackInfo__{} }
func (m *BfdrtrTrackInfo__) String() string { return proto.CompactTextString(m) }
func (*BfdrtrTrackInfo__) ProtoMessage()    {}
func (*BfdrtrTrackInfo__) Descriptor() ([]byte, []int) {
	return fileDescriptor_03fe192acc1772f2, []int{4}
}

func (m *BfdrtrTrackInfo__) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdrtrTrackInfo__.Unmarshal(m, b)
}
func (m *BfdrtrTrackInfo__) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdrtrTrackInfo__.Marshal(b, m, deterministic)
}
func (m *BfdrtrTrackInfo__) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdrtrTrackInfo__.Merge(m, src)
}
func (m *BfdrtrTrackInfo__) XXX_Size() int {
	return xxx_messageInfo_BfdrtrTrackInfo__.Size(m)
}
func (m *BfdrtrTrackInfo__) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdrtrTrackInfo__.DiscardUnknown(m)
}

var xxx_messageInfo_BfdrtrTrackInfo__ proto.InternalMessageInfo

func (m *BfdrtrTrackInfo__) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *BfdrtrTrackInfo__) GetDestinationAddress() uint32 {
	if m != nil {
		return m.DestinationAddress
	}
	return 0
}

func (m *BfdrtrTrackInfo__) GetRate() uint32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *BfdrtrTrackInfo__) GetDebounceCount() uint32 {
	if m != nil {
		return m.DebounceCount
	}
	return 0
}

type TrackTypeInfoUnion struct {
	Discriminant         string             `protobuf:"bytes,1,opt,name=discriminant,proto3" json:"discriminant,omitempty"`
	InterfaceTracks      *IntfTrackInfo__   `protobuf:"bytes,2,opt,name=interface_tracks,json=interfaceTracks,proto3" json:"interface_tracks,omitempty"`
	RouteTracks          *RouteTrackInfo__  `protobuf:"bytes,3,opt,name=route_tracks,json=routeTracks,proto3" json:"route_tracks,omitempty"`
	IpslaTracks          *RtrTrackInfo__    `protobuf:"bytes,4,opt,name=ipsla_tracks,json=ipslaTracks,proto3" json:"ipsla_tracks,omitempty"`
	BfdTracks            *BfdrtrTrackInfo__ `protobuf:"bytes,5,opt,name=bfd_tracks,json=bfdTracks,proto3" json:"bfd_tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TrackTypeInfoUnion) Reset()         { *m = TrackTypeInfoUnion{} }
func (m *TrackTypeInfoUnion) String() string { return proto.CompactTextString(m) }
func (*TrackTypeInfoUnion) ProtoMessage()    {}
func (*TrackTypeInfoUnion) Descriptor() ([]byte, []int) {
	return fileDescriptor_03fe192acc1772f2, []int{5}
}

func (m *TrackTypeInfoUnion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackTypeInfoUnion.Unmarshal(m, b)
}
func (m *TrackTypeInfoUnion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackTypeInfoUnion.Marshal(b, m, deterministic)
}
func (m *TrackTypeInfoUnion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackTypeInfoUnion.Merge(m, src)
}
func (m *TrackTypeInfoUnion) XXX_Size() int {
	return xxx_messageInfo_TrackTypeInfoUnion.Size(m)
}
func (m *TrackTypeInfoUnion) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackTypeInfoUnion.DiscardUnknown(m)
}

var xxx_messageInfo_TrackTypeInfoUnion proto.InternalMessageInfo

func (m *TrackTypeInfoUnion) GetDiscriminant() string {
	if m != nil {
		return m.Discriminant
	}
	return ""
}

func (m *TrackTypeInfoUnion) GetInterfaceTracks() *IntfTrackInfo__ {
	if m != nil {
		return m.InterfaceTracks
	}
	return nil
}

func (m *TrackTypeInfoUnion) GetRouteTracks() *RouteTrackInfo__ {
	if m != nil {
		return m.RouteTracks
	}
	return nil
}

func (m *TrackTypeInfoUnion) GetIpslaTracks() *RtrTrackInfo__ {
	if m != nil {
		return m.IpslaTracks
	}
	return nil
}

func (m *TrackTypeInfoUnion) GetBfdTracks() *BfdrtrTrackInfo__ {
	if m != nil {
		return m.BfdTracks
	}
	return nil
}

type TrackInfoBrief__Item struct {
	TrackeName           string              `protobuf:"bytes,1,opt,name=tracke_name,json=trackeName,proto3" json:"tracke_name,omitempty"`
	Type                 string              `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	TrackState           bool                `protobuf:"varint,3,opt,name=track_state,json=trackState,proto3" json:"track_state,omitempty"`
	TrackTypeInfo        *TrackTypeInfoUnion `protobuf:"bytes,4,opt,name=track_type_info,json=trackTypeInfo,proto3" json:"track_type_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TrackInfoBrief__Item) Reset()         { *m = TrackInfoBrief__Item{} }
func (m *TrackInfoBrief__Item) String() string { return proto.CompactTextString(m) }
func (*TrackInfoBrief__Item) ProtoMessage()    {}
func (*TrackInfoBrief__Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_03fe192acc1772f2, []int{6}
}

func (m *TrackInfoBrief__Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackInfoBrief__Item.Unmarshal(m, b)
}
func (m *TrackInfoBrief__Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackInfoBrief__Item.Marshal(b, m, deterministic)
}
func (m *TrackInfoBrief__Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackInfoBrief__Item.Merge(m, src)
}
func (m *TrackInfoBrief__Item) XXX_Size() int {
	return xxx_messageInfo_TrackInfoBrief__Item.Size(m)
}
func (m *TrackInfoBrief__Item) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackInfoBrief__Item.DiscardUnknown(m)
}

var xxx_messageInfo_TrackInfoBrief__Item proto.InternalMessageInfo

func (m *TrackInfoBrief__Item) GetTrackeName() string {
	if m != nil {
		return m.TrackeName
	}
	return ""
}

func (m *TrackInfoBrief__Item) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TrackInfoBrief__Item) GetTrackState() bool {
	if m != nil {
		return m.TrackState
	}
	return false
}

func (m *TrackInfoBrief__Item) GetTrackTypeInfo() *TrackTypeInfoUnion {
	if m != nil {
		return m.TrackTypeInfo
	}
	return nil
}

type TrackInfoBrief__Entry struct {
	TrackInfoBrief__     []*TrackInfoBrief__Item `protobuf:"bytes,50,rep,name=track_info_brief__,json=trackInfoBrief,proto3" json:"track_info_brief__,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TrackInfoBrief__Entry) Reset()         { *m = TrackInfoBrief__Entry{} }
func (m *TrackInfoBrief__Entry) String() string { return proto.CompactTextString(m) }
func (*TrackInfoBrief__Entry) ProtoMessage()    {}
func (*TrackInfoBrief__Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_03fe192acc1772f2, []int{7}
}

func (m *TrackInfoBrief__Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackInfoBrief__Entry.Unmarshal(m, b)
}
func (m *TrackInfoBrief__Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackInfoBrief__Entry.Marshal(b, m, deterministic)
}
func (m *TrackInfoBrief__Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackInfoBrief__Entry.Merge(m, src)
}
func (m *TrackInfoBrief__Entry) XXX_Size() int {
	return xxx_messageInfo_TrackInfoBrief__Entry.Size(m)
}
func (m *TrackInfoBrief__Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackInfoBrief__Entry.DiscardUnknown(m)
}

var xxx_messageInfo_TrackInfoBrief__Entry proto.InternalMessageInfo

func (m *TrackInfoBrief__Entry) GetTrackInfoBrief__() []*TrackInfoBrief__Item {
	if m != nil {
		return m.TrackInfoBrief__
	}
	return nil
}

func init() {
	proto.RegisterType((*TrackInfoBrief__Entry_KEYS)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route_brief.track_info_brief___entry_KEYS")
	proto.RegisterType((*IntfTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route_brief.intf_track_info__")
	proto.RegisterType((*RouteTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route_brief.route_track_info__")
	proto.RegisterType((*RtrTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route_brief.rtr_track_info__")
	proto.RegisterType((*BfdrtrTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route_brief.bfdrtr_track_info__")
	proto.RegisterType((*TrackTypeInfoUnion)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route_brief.track_type_info_union")
	proto.RegisterType((*TrackInfoBrief__Item)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route_brief.track_info_brief___item")
	proto.RegisterType((*TrackInfoBrief__Entry)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route_brief.track_info_brief___entry")
}

func init() { proto.RegisterFile("track_info_brief___entry.proto", fileDescriptor_03fe192acc1772f2) }

var fileDescriptor_03fe192acc1772f2 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x6d, 0x5a, 0x9a, 0x49, 0xb6, 0x2d, 0xae, 0x80, 0x70, 0xa0, 0xad, 0x16, 0x21,
	0xf5, 0x14, 0xa4, 0xc0, 0x89, 0x1b, 0x54, 0x48, 0x54, 0x20, 0x90, 0xb6, 0xbd, 0x20, 0x0e, 0x23,
	0xef, 0x7a, 0xb6, 0x35, 0x34, 0xb6, 0xe5, 0x75, 0xaa, 0xe6, 0xc2, 0x13, 0x20, 0xf1, 0xe7, 0x06,
	0x27, 0x8e, 0xbc, 0x04, 0xef, 0x86, 0x6c, 0x6f, 0x4a, 0xd2, 0xb4, 0x12, 0x97, 0x70, 0x9b, 0xfd,
	0x26, 0xf6, 0xf7, 0x8b, 0xf7, 0x1b, 0x2f, 0x6c, 0x3b, 0xcb, 0xcb, 0x0f, 0x28, 0x55, 0xa5, 0xb1,
	0xb0, 0x92, 0x2a, 0x44, 0x24, 0xe5, 0xec, 0xb8, 0x6f, 0xac, 0x76, 0x9a, 0xbd, 0x2b, 0x65, 0x5d,
	0x6a, 0x94, 0xba, 0xc6, 0x73, 0x8b, 0x43, 0xae, 0xf8, 0x31, 0xf1, 0x42, 0x9e, 0x4a, 0x37, 0x46,
	0x5d, 0xbc, 0xa7, 0xd2, 0x61, 0xd8, 0x41, 0xaa, 0x63, 0xd4, 0x86, 0x6c, 0xff, 0x92, 0xd8, 0x8f,
	0xfb, 0xbb, 0xb1, 0x21, 0x94, 0xe6, 0xec, 0x31, 0x5a, 0x3d, 0x72, 0x14, 0xad, 0xb2, 0x1d, 0xb8,
	0x77, 0x9d, 0x3d, 0xbe, 0x7c, 0xfe, 0xf6, 0x30, 0x7b, 0x02, 0x37, 0xa5, 0x72, 0x15, 0x4e, 0xfd,
	0x0a, 0xd9, 0x03, 0x58, 0x97, 0xca, 0x91, 0xad, 0x78, 0x49, 0xa8, 0xf8, 0x90, 0x7a, 0xc9, 0x6e,
	0xb2, 0xd7, 0xce, 0xd3, 0x0b, 0xf5, 0x35, 0x1f, 0x52, 0xf6, 0x11, 0x58, 0xf4, 0x9a, 0x59, 0x7c,
	0x1b, 0x56, 0x8d, 0xa5, 0x4a, 0x9e, 0x87, 0x45, 0x69, 0xde, 0x3c, 0xb1, 0xfb, 0x90, 0xc6, 0x0a,
	0x4f, 0x49, 0x1d, 0xbb, 0x93, 0xde, 0x52, 0x68, 0x77, 0xa3, 0xf8, 0x2a, 0x68, 0x6c, 0x13, 0x96,
	0xcf, 0x6c, 0xd5, 0x5b, 0x0e, 0x76, 0xbe, 0x64, 0x77, 0x61, 0x4d, 0xd1, 0xb9, 0xc3, 0x13, 0x6d,
	0x7a, 0xad, 0x20, 0xdf, 0xf0, 0xcf, 0x2f, 0xb4, 0xc9, 0x08, 0x36, 0xad, 0xb3, 0xb3, 0xee, 0xdb,
	0xd0, 0x91, 0xa6, 0x3e, 0xe5, 0xa8, 0x0d, 0x4a, 0xd1, 0x20, 0xb4, 0x83, 0xf4, 0xc6, 0x1c, 0x08,
	0x6f, 0x60, 0x9d, 0x6b, 0xbc, 0x7d, 0xc9, 0x76, 0xa0, 0x63, 0xc9, 0x8d, 0xac, 0xc2, 0x52, 0x0b,
	0x0a, 0xd6, 0x69, 0x0e, 0x51, 0xda, 0xd7, 0x82, 0xb2, 0x5f, 0x09, 0x6c, 0x15, 0x95, 0x98, 0xb3,
	0xfa, 0xb7, 0x53, 0x62, 0x0f, 0x61, 0x4b, 0x50, 0xed, 0xa4, 0xe2, 0x4e, 0x6a, 0x85, 0x5c, 0x08,
	0x4b, 0x75, 0xdd, 0x10, 0xb0, 0xa9, 0xd6, 0xd3, 0xd8, 0x61, 0x0c, 0x5a, 0x96, 0xbb, 0x09, 0x49,
	0xa8, 0xbd, 0x97, 0xa0, 0x42, 0x8f, 0x54, 0x49, 0x58, 0xea, 0x91, 0x72, 0xe1, 0x2c, 0xd2, 0x3c,
	0x9d, 0xa8, 0xfb, 0x5e, 0xcc, 0xbe, 0xae, 0xc0, 0xad, 0xe9, 0x38, 0x78, 0xd0, 0x91, 0x92, 0x5a,
	0xb1, 0x0c, 0xba, 0x42, 0xd6, 0xa5, 0x95, 0x43, 0xa9, 0xb8, 0x72, 0x0d, 0xea, 0x8c, 0xc6, 0xbe,
	0x27, 0xb0, 0xf9, 0xf7, 0x1f, 0x85, 0x7d, 0x22, 0x67, 0x67, 0xa0, 0xfa, 0x0b, 0x4c, 0x69, 0x7f,
	0x2e, 0x81, 0xf9, 0xc6, 0x05, 0xc7, 0x51, 0xc0, 0x60, 0xdf, 0x12, 0xe8, 0x4e, 0x85, 0xad, 0x0e,
	0xa7, 0xd3, 0x19, 0xe8, 0x85, 0x72, 0xcd, 0xa7, 0x3b, 0xef, 0x04, 0xad, 0x81, 0xfa, 0x92, 0x40,
	0x37, 0xa6, 0xad, 0x81, 0x6a, 0x05, 0xa8, 0xe1, 0x62, 0xa1, 0x2e, 0xe5, 0x30, 0x8f, 0x81, 0x6f,
	0x90, 0x3e, 0x27, 0x00, 0x45, 0x25, 0x26, 0x40, 0x2b, 0x01, 0xc8, 0x2c, 0x14, 0xe8, 0x8a, 0xd9,
	0xc8, 0xdb, 0x45, 0x25, 0x22, 0x51, 0xf6, 0x69, 0x09, 0xee, 0x5c, 0x71, 0x07, 0x49, 0x47, 0x43,
	0x3f, 0x7b, 0xa1, 0x35, 0x33, 0x3f, 0x10, 0xa5, 0x30, 0x3c, 0x0c, 0x5a, 0xde, 0x34, 0xa4, 0xb0,
	0x9d, 0x87, 0xfa, 0x62, 0x11, 0xd6, 0x6e, 0x32, 0x26, 0x6b, 0xcd, 0xa2, 0x43, 0xaf, 0xb0, 0x1f,
	0x09, 0x6c, 0x5c, 0x9a, 0x82, 0xe6, 0xcd, 0xd8, 0x85, 0x1e, 0xc4, 0x95, 0x93, 0x97, 0xa7, 0x41,
	0x3e, 0x1a, 0x1b, 0x3a, 0x50, 0x95, 0xce, 0x7e, 0x27, 0xd0, 0xbb, 0xee, 0x4a, 0x66, 0x3f, 0x13,
	0x60, 0xf3, 0xcd, 0xde, 0x60, 0x77, 0x79, 0xaf, 0x33, 0x70, 0xff, 0x01, 0x7e, 0xee, 0x15, 0xe5,
	0xeb, 0xa1, 0xe1, 0xd1, 0x9f, 0x79, 0xb9, 0x58, 0x0d, 0x5f, 0xad, 0x47, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x78, 0xdc, 0x41, 0x42, 0xd7, 0x06, 0x00, 0x00,
}
