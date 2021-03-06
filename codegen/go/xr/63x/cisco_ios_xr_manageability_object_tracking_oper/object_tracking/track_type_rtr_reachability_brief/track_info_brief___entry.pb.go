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

package cisco_ios_xr_manageability_object_tracking_oper_object_tracking_track_type_rtr_reachability_brief

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
	proto.RegisterType((*TrackInfoBrief__Entry_KEYS)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_rtr_reachability_brief.track_info_brief___entry_KEYS")
	proto.RegisterType((*IntfTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_rtr_reachability_brief.intf_track_info__")
	proto.RegisterType((*RouteTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_rtr_reachability_brief.route_track_info__")
	proto.RegisterType((*RtrTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_rtr_reachability_brief.rtr_track_info__")
	proto.RegisterType((*BfdrtrTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_rtr_reachability_brief.bfdrtr_track_info__")
	proto.RegisterType((*TrackTypeInfoUnion)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_rtr_reachability_brief.track_type_info_union")
	proto.RegisterType((*TrackInfoBrief__Item)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_rtr_reachability_brief.track_info_brief___item")
	proto.RegisterType((*TrackInfoBrief__Entry)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_rtr_reachability_brief.track_info_brief___entry")
}

func init() { proto.RegisterFile("track_info_brief___entry.proto", fileDescriptor_03fe192acc1772f2) }

var fileDescriptor_03fe192acc1772f2 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xe5, 0x36, 0xed, 0xbf, 0x99, 0xc4, 0x6d, 0xff, 0x5b, 0x01, 0xe1, 0x40, 0x5b, 0x19,
	0x21, 0xf5, 0x14, 0xa4, 0x70, 0xe3, 0x06, 0x15, 0x12, 0x15, 0x08, 0x24, 0xb7, 0x17, 0x4e, 0xa3,
	0xb5, 0x77, 0xdc, 0x2e, 0x34, 0xbb, 0xd6, 0x7a, 0x53, 0x25, 0x1c, 0x78, 0x08, 0x24, 0x10, 0x37,
	0x38, 0x21, 0x78, 0x0e, 0x5e, 0x0c, 0xed, 0xd8, 0x29, 0x49, 0xd3, 0x4a, 0x5c, 0xc8, 0x6d, 0xfc,
	0x4d, 0x76, 0xbf, 0xdf, 0xae, 0xe7, 0x73, 0x60, 0xd7, 0x3b, 0x99, 0xbf, 0x43, 0x6d, 0x0a, 0x8b,
	0x99, 0xd3, 0x54, 0x20, 0x22, 0x19, 0xef, 0x26, 0xfd, 0xd2, 0x59, 0x6f, 0x85, 0xcc, 0x75, 0x95,
	0x5b, 0xd4, 0xb6, 0xc2, 0xb1, 0xc3, 0xa1, 0x34, 0xf2, 0x94, 0x64, 0xa6, 0xcf, 0xb5, 0x9f, 0xa0,
	0xcd, 0xde, 0x52, 0xee, 0x91, 0x77, 0xd0, 0xe6, 0x14, 0x6d, 0x49, 0xae, 0x7f, 0x45, 0xec, 0xd7,
	0xfb, 0xfb, 0x49, 0x49, 0xe8, 0xbc, 0x43, 0x47, 0x32, 0x3f, 0x9b, 0x6e, 0xc0, 0x86, 0xc9, 0x1e,
	0xdc, 0xbb, 0x09, 0x02, 0x5f, 0x3c, 0x7b, 0x73, 0x9c, 0x3c, 0x86, 0xff, 0xb5, 0xf1, 0x05, 0xce,
	0xfc, 0x0a, 0xc5, 0x03, 0xd8, 0xd4, 0xc6, 0x93, 0x2b, 0x64, 0x4e, 0x68, 0xe4, 0x90, 0x7a, 0xd1,
	0x7e, 0x74, 0xd0, 0x4e, 0xe3, 0x4b, 0xf5, 0x95, 0x1c, 0x52, 0xf2, 0x01, 0x84, 0xb3, 0x23, 0x4f,
	0xf3, 0x8b, 0x6f, 0xc3, 0x7a, 0xe9, 0xa8, 0xd0, 0x63, 0x5e, 0x14, 0xa7, 0xcd, 0x93, 0xb8, 0x0f,
	0x71, 0x5d, 0xe1, 0x39, 0x99, 0x53, 0x7f, 0xd6, 0x5b, 0xe1, 0x76, 0xb7, 0x16, 0x5f, 0xb2, 0x26,
	0xb6, 0x61, 0xf5, 0xc2, 0x15, 0xbd, 0x55, 0xb6, 0x0b, 0xa5, 0xb8, 0x0b, 0x1b, 0x86, 0xc6, 0x1e,
	0xcf, 0x6c, 0xd9, 0x6b, 0xb1, 0xfc, 0x5f, 0x78, 0x7e, 0x6e, 0xcb, 0x84, 0x60, 0x3b, 0x1c, 0x7b,
	0xce, 0x7d, 0x17, 0x3a, 0xba, 0xac, 0xce, 0x25, 0xda, 0x12, 0xb5, 0x6a, 0x10, 0xda, 0x2c, 0xbd,
	0x2e, 0x8f, 0x54, 0x30, 0x70, 0xde, 0x37, 0xde, 0xa1, 0x14, 0x7b, 0xd0, 0x71, 0xe4, 0x47, 0xce,
	0x60, 0x6e, 0x15, 0xb1, 0x75, 0x9c, 0x42, 0x2d, 0x1d, 0x5a, 0x45, 0xc9, 0x8f, 0x08, 0x76, 0xb2,
	0x42, 0x2d, 0x58, 0xfd, 0xdd, 0x2d, 0x89, 0x87, 0xb0, 0xa3, 0xa8, 0xf2, 0xda, 0x48, 0xaf, 0xad,
	0x41, 0xa9, 0x94, 0xa3, 0xaa, 0x6a, 0x08, 0xc4, 0x4c, 0xeb, 0x49, 0xdd, 0x11, 0x02, 0x5a, 0x4e,
	0xfa, 0x29, 0x09, 0xd7, 0xc1, 0x4b, 0x51, 0x66, 0x47, 0x26, 0x27, 0xcc, 0xed, 0xc8, 0x78, 0xbe,
	0x8b, 0x38, 0x8d, 0xa7, 0xea, 0x61, 0x10, 0x93, 0x9f, 0x6b, 0x70, 0x6b, 0x66, 0x28, 0x18, 0x74,
	0x64, 0xb4, 0x35, 0x22, 0x81, 0xae, 0xd2, 0x55, 0xee, 0xf4, 0x50, 0x1b, 0x69, 0x7c, 0x83, 0x3a,
	0xa7, 0x89, 0xaf, 0x11, 0x6c, 0xff, 0x39, 0x11, 0xef, 0x53, 0x73, 0x76, 0x06, 0xbe, 0xff, 0xcf,
	0x67, 0xb5, 0xbf, 0x30, 0x87, 0xe9, 0xd6, 0x25, 0xcd, 0x09, 0xc3, 0x88, 0x2f, 0x11, 0x74, 0x67,
	0x46, 0xae, 0xe2, 0x3b, 0xea, 0x0c, 0x46, 0x4b, 0xa0, 0x5b, 0x9c, 0xf4, 0xb4, 0xc3, 0x5a, 0x83,
	0xf6, 0x39, 0x82, 0x6e, 0x3d, 0x79, 0x0d, 0x5a, 0x8b, 0xd1, 0xaa, 0x65, 0xa0, 0x5d, 0x99, 0xcc,
	0xb4, 0x8e, 0x40, 0x03, 0xf6, 0x29, 0x02, 0xc8, 0x0a, 0x35, 0xc5, 0x5a, 0x63, 0xac, 0x8b, 0x25,
	0x60, 0x5d, 0x93, 0x99, 0xb4, 0x9d, 0x15, 0xaa, 0xe6, 0x4a, 0x3e, 0xae, 0xc0, 0x9d, 0x6b, 0xbe,
	0x4d, 0xda, 0xd3, 0x30, 0x64, 0x92, 0x5b, 0x73, 0xb9, 0x82, 0x5a, 0xe2, 0x50, 0x09, 0x68, 0x05,
	0x6b, 0x9e, 0xce, 0x76, 0xca, 0xf5, 0xe5, 0x22, 0xac, 0xfc, 0x34, 0x3e, 0x1b, 0xcd, 0xa2, 0xe3,
	0xa0, 0x88, 0x6f, 0x11, 0x6c, 0x5d, 0x49, 0x47, 0xf3, 0x96, 0xc6, 0x4b, 0xb8, 0x8e, 0x6b, 0x73,
	0x99, 0xc6, 0x2c, 0x9f, 0x4c, 0x4a, 0x3a, 0x32, 0x85, 0x4d, 0x7e, 0x45, 0xd0, 0xbb, 0xe9, 0x83,
	0x2d, 0xbe, 0x47, 0x20, 0x16, 0x9b, 0xbd, 0xc1, 0xfe, 0xea, 0x41, 0x67, 0xf0, 0x7e, 0x69, 0x47,
	0x58, 0x78, 0x5d, 0xe9, 0x26, 0x37, 0xc2, 0x01, 0x9e, 0x06, 0x39, 0x5b, 0xe7, 0xff, 0xb7, 0x47,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x33, 0xc6, 0x39, 0x01, 0x07, 0x00, 0x00,
}
