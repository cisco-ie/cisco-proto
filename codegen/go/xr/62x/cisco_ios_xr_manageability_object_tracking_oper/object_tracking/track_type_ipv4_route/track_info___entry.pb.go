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
// source: track_info___entry.proto

package cisco_ios_xr_manageability_object_tracking_oper_object_tracking_track_type_ipv4_route

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

type TrackInfo__Entry_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackInfo__Entry_KEYS) Reset()         { *m = TrackInfo__Entry_KEYS{} }
func (m *TrackInfo__Entry_KEYS) String() string { return proto.CompactTextString(m) }
func (*TrackInfo__Entry_KEYS) ProtoMessage()    {}
func (*TrackInfo__Entry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee7ca1ee6d013b, []int{0}
}

func (m *TrackInfo__Entry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackInfo__Entry_KEYS.Unmarshal(m, b)
}
func (m *TrackInfo__Entry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackInfo__Entry_KEYS.Marshal(b, m, deterministic)
}
func (m *TrackInfo__Entry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackInfo__Entry_KEYS.Merge(m, src)
}
func (m *TrackInfo__Entry_KEYS) XXX_Size() int {
	return xxx_messageInfo_TrackInfo__Entry_KEYS.Size(m)
}
func (m *TrackInfo__Entry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackInfo__Entry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TrackInfo__Entry_KEYS proto.InternalMessageInfo

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
	return fileDescriptor_94ee7ca1ee6d013b, []int{1}
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
	return fileDescriptor_94ee7ca1ee6d013b, []int{2}
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
	return fileDescriptor_94ee7ca1ee6d013b, []int{3}
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
	return fileDescriptor_94ee7ca1ee6d013b, []int{4}
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
	return fileDescriptor_94ee7ca1ee6d013b, []int{5}
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

type BoolTrackInfo__Item struct {
	ObjectName           string   `protobuf:"bytes,1,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	TrackState           bool     `protobuf:"varint,2,opt,name=track_state,json=trackState,proto3" json:"track_state,omitempty"`
	WithNot              bool     `protobuf:"varint,3,opt,name=with_not,json=withNot,proto3" json:"with_not,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolTrackInfo__Item) Reset()         { *m = BoolTrackInfo__Item{} }
func (m *BoolTrackInfo__Item) String() string { return proto.CompactTextString(m) }
func (*BoolTrackInfo__Item) ProtoMessage()    {}
func (*BoolTrackInfo__Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee7ca1ee6d013b, []int{6}
}

func (m *BoolTrackInfo__Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolTrackInfo__Item.Unmarshal(m, b)
}
func (m *BoolTrackInfo__Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolTrackInfo__Item.Marshal(b, m, deterministic)
}
func (m *BoolTrackInfo__Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolTrackInfo__Item.Merge(m, src)
}
func (m *BoolTrackInfo__Item) XXX_Size() int {
	return xxx_messageInfo_BoolTrackInfo__Item.Size(m)
}
func (m *BoolTrackInfo__Item) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolTrackInfo__Item.DiscardUnknown(m)
}

var xxx_messageInfo_BoolTrackInfo__Item proto.InternalMessageInfo

func (m *BoolTrackInfo__Item) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

func (m *BoolTrackInfo__Item) GetTrackState() bool {
	if m != nil {
		return m.TrackState
	}
	return false
}

func (m *BoolTrackInfo__Item) GetWithNot() bool {
	if m != nil {
		return m.WithNot
	}
	return false
}

type BoolTrackInfo__Entry struct {
	BoolTrackInfo__      []*BoolTrackInfo__Item `protobuf:"bytes,4,rep,name=bool_track_info__,json=boolTrackInfo,proto3" json:"bool_track_info__,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BoolTrackInfo__Entry) Reset()         { *m = BoolTrackInfo__Entry{} }
func (m *BoolTrackInfo__Entry) String() string { return proto.CompactTextString(m) }
func (*BoolTrackInfo__Entry) ProtoMessage()    {}
func (*BoolTrackInfo__Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee7ca1ee6d013b, []int{7}
}

func (m *BoolTrackInfo__Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolTrackInfo__Entry.Unmarshal(m, b)
}
func (m *BoolTrackInfo__Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolTrackInfo__Entry.Marshal(b, m, deterministic)
}
func (m *BoolTrackInfo__Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolTrackInfo__Entry.Merge(m, src)
}
func (m *BoolTrackInfo__Entry) XXX_Size() int {
	return xxx_messageInfo_BoolTrackInfo__Entry.Size(m)
}
func (m *BoolTrackInfo__Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolTrackInfo__Entry.DiscardUnknown(m)
}

var xxx_messageInfo_BoolTrackInfo__Entry proto.InternalMessageInfo

func (m *BoolTrackInfo__Entry) GetBoolTrackInfo__() []*BoolTrackInfo__Item {
	if m != nil {
		return m.BoolTrackInfo__
	}
	return nil
}

type ThresholdTrackInfo__Item struct {
	ObjectName           string   `protobuf:"bytes,1,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	TrackState           bool     `protobuf:"varint,2,opt,name=track_state,json=trackState,proto3" json:"track_state,omitempty"`
	Weight               uint32   `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThresholdTrackInfo__Item) Reset()         { *m = ThresholdTrackInfo__Item{} }
func (m *ThresholdTrackInfo__Item) String() string { return proto.CompactTextString(m) }
func (*ThresholdTrackInfo__Item) ProtoMessage()    {}
func (*ThresholdTrackInfo__Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee7ca1ee6d013b, []int{8}
}

func (m *ThresholdTrackInfo__Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThresholdTrackInfo__Item.Unmarshal(m, b)
}
func (m *ThresholdTrackInfo__Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThresholdTrackInfo__Item.Marshal(b, m, deterministic)
}
func (m *ThresholdTrackInfo__Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThresholdTrackInfo__Item.Merge(m, src)
}
func (m *ThresholdTrackInfo__Item) XXX_Size() int {
	return xxx_messageInfo_ThresholdTrackInfo__Item.Size(m)
}
func (m *ThresholdTrackInfo__Item) XXX_DiscardUnknown() {
	xxx_messageInfo_ThresholdTrackInfo__Item.DiscardUnknown(m)
}

var xxx_messageInfo_ThresholdTrackInfo__Item proto.InternalMessageInfo

func (m *ThresholdTrackInfo__Item) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

func (m *ThresholdTrackInfo__Item) GetTrackState() bool {
	if m != nil {
		return m.TrackState
	}
	return false
}

func (m *ThresholdTrackInfo__Item) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type ThresholdTrackInfo__Entry struct {
	ThresholdTrackInfo__ []*ThresholdTrackInfo__Item `protobuf:"bytes,4,rep,name=threshold_track_info__,json=thresholdTrackInfo,proto3" json:"threshold_track_info__,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ThresholdTrackInfo__Entry) Reset()         { *m = ThresholdTrackInfo__Entry{} }
func (m *ThresholdTrackInfo__Entry) String() string { return proto.CompactTextString(m) }
func (*ThresholdTrackInfo__Entry) ProtoMessage()    {}
func (*ThresholdTrackInfo__Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee7ca1ee6d013b, []int{9}
}

func (m *ThresholdTrackInfo__Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThresholdTrackInfo__Entry.Unmarshal(m, b)
}
func (m *ThresholdTrackInfo__Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThresholdTrackInfo__Entry.Marshal(b, m, deterministic)
}
func (m *ThresholdTrackInfo__Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThresholdTrackInfo__Entry.Merge(m, src)
}
func (m *ThresholdTrackInfo__Entry) XXX_Size() int {
	return xxx_messageInfo_ThresholdTrackInfo__Entry.Size(m)
}
func (m *ThresholdTrackInfo__Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_ThresholdTrackInfo__Entry.DiscardUnknown(m)
}

var xxx_messageInfo_ThresholdTrackInfo__Entry proto.InternalMessageInfo

func (m *ThresholdTrackInfo__Entry) GetThresholdTrackInfo__() []*ThresholdTrackInfo__Item {
	if m != nil {
		return m.ThresholdTrackInfo__
	}
	return nil
}

type InterfaceTrackingInfo__Item struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceTrackingInfo__Item) Reset()         { *m = InterfaceTrackingInfo__Item{} }
func (m *InterfaceTrackingInfo__Item) String() string { return proto.CompactTextString(m) }
func (*InterfaceTrackingInfo__Item) ProtoMessage()    {}
func (*InterfaceTrackingInfo__Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee7ca1ee6d013b, []int{10}
}

func (m *InterfaceTrackingInfo__Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceTrackingInfo__Item.Unmarshal(m, b)
}
func (m *InterfaceTrackingInfo__Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceTrackingInfo__Item.Marshal(b, m, deterministic)
}
func (m *InterfaceTrackingInfo__Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceTrackingInfo__Item.Merge(m, src)
}
func (m *InterfaceTrackingInfo__Item) XXX_Size() int {
	return xxx_messageInfo_InterfaceTrackingInfo__Item.Size(m)
}
func (m *InterfaceTrackingInfo__Item) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceTrackingInfo__Item.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceTrackingInfo__Item proto.InternalMessageInfo

func (m *InterfaceTrackingInfo__Item) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type InterfaceTrackingInfo__Entry struct {
	InterfaceTrackingInfo__ []*InterfaceTrackingInfo__Item `protobuf:"bytes,2,rep,name=interface_tracking_info__,json=interfaceTrackingInfo,proto3" json:"interface_tracking_info__,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                       `json:"-"`
	XXX_unrecognized        []byte                         `json:"-"`
	XXX_sizecache           int32                          `json:"-"`
}

func (m *InterfaceTrackingInfo__Entry) Reset()         { *m = InterfaceTrackingInfo__Entry{} }
func (m *InterfaceTrackingInfo__Entry) String() string { return proto.CompactTextString(m) }
func (*InterfaceTrackingInfo__Entry) ProtoMessage()    {}
func (*InterfaceTrackingInfo__Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee7ca1ee6d013b, []int{11}
}

func (m *InterfaceTrackingInfo__Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceTrackingInfo__Entry.Unmarshal(m, b)
}
func (m *InterfaceTrackingInfo__Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceTrackingInfo__Entry.Marshal(b, m, deterministic)
}
func (m *InterfaceTrackingInfo__Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceTrackingInfo__Entry.Merge(m, src)
}
func (m *InterfaceTrackingInfo__Entry) XXX_Size() int {
	return xxx_messageInfo_InterfaceTrackingInfo__Entry.Size(m)
}
func (m *InterfaceTrackingInfo__Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceTrackingInfo__Entry.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceTrackingInfo__Entry proto.InternalMessageInfo

func (m *InterfaceTrackingInfo__Entry) GetInterfaceTrackingInfo__() []*InterfaceTrackingInfo__Item {
	if m != nil {
		return m.InterfaceTrackingInfo__
	}
	return nil
}

type DelayedStateStatus__ struct {
	TimeRemaining        uint32   `protobuf:"varint,1,opt,name=time_remaining,json=timeRemaining,proto3" json:"time_remaining,omitempty"`
	TrackState           bool     `protobuf:"varint,2,opt,name=track_state,json=trackState,proto3" json:"track_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelayedStateStatus__) Reset()         { *m = DelayedStateStatus__{} }
func (m *DelayedStateStatus__) String() string { return proto.CompactTextString(m) }
func (*DelayedStateStatus__) ProtoMessage()    {}
func (*DelayedStateStatus__) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee7ca1ee6d013b, []int{12}
}

func (m *DelayedStateStatus__) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelayedStateStatus__.Unmarshal(m, b)
}
func (m *DelayedStateStatus__) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelayedStateStatus__.Marshal(b, m, deterministic)
}
func (m *DelayedStateStatus__) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelayedStateStatus__.Merge(m, src)
}
func (m *DelayedStateStatus__) XXX_Size() int {
	return xxx_messageInfo_DelayedStateStatus__.Size(m)
}
func (m *DelayedStateStatus__) XXX_DiscardUnknown() {
	xxx_messageInfo_DelayedStateStatus__.DiscardUnknown(m)
}

var xxx_messageInfo_DelayedStateStatus__ proto.InternalMessageInfo

func (m *DelayedStateStatus__) GetTimeRemaining() uint32 {
	if m != nil {
		return m.TimeRemaining
	}
	return 0
}

func (m *DelayedStateStatus__) GetTrackState() bool {
	if m != nil {
		return m.TrackState
	}
	return false
}

type TrackInfo__Item struct {
	TrackeName           string                        `protobuf:"bytes,1,opt,name=tracke_name,json=trackeName,proto3" json:"tracke_name,omitempty"`
	Type                 string                        `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	TrackState           bool                          `protobuf:"varint,3,opt,name=track_state,json=trackState,proto3" json:"track_state,omitempty"`
	StateChangeCounter   uint32                        `protobuf:"varint,4,opt,name=state_change_counter,json=stateChangeCounter,proto3" json:"state_change_counter,omitempty"`
	SecondsLastChange    uint64                        `protobuf:"varint,5,opt,name=seconds_last_change,json=secondsLastChange,proto3" json:"seconds_last_change,omitempty"`
	TrackTypeInfo        *TrackTypeInfoUnion           `protobuf:"bytes,6,opt,name=track_type_info,json=trackTypeInfo,proto3" json:"track_type_info,omitempty"`
	BoolTracks           *BoolTrackInfo__Entry         `protobuf:"bytes,7,opt,name=bool_tracks,json=boolTracks,proto3" json:"bool_tracks,omitempty"`
	ThresholdTracks      *ThresholdTrackInfo__Entry    `protobuf:"bytes,8,opt,name=threshold_tracks,json=thresholdTracks,proto3" json:"threshold_tracks,omitempty"`
	ThresholdUp          uint32                        `protobuf:"varint,9,opt,name=threshold_up,json=thresholdUp,proto3" json:"threshold_up,omitempty"`
	ThresholdDown        uint32                        `protobuf:"varint,10,opt,name=threshold_down,json=thresholdDown,proto3" json:"threshold_down,omitempty"`
	TrackingInteraces    *InterfaceTrackingInfo__Entry `protobuf:"bytes,11,opt,name=tracking_interaces,json=trackingInteraces,proto3" json:"tracking_interaces,omitempty"`
	Delayed              *DelayedStateStatus__         `protobuf:"bytes,12,opt,name=delayed,proto3" json:"delayed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TrackInfo__Item) Reset()         { *m = TrackInfo__Item{} }
func (m *TrackInfo__Item) String() string { return proto.CompactTextString(m) }
func (*TrackInfo__Item) ProtoMessage()    {}
func (*TrackInfo__Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee7ca1ee6d013b, []int{13}
}

func (m *TrackInfo__Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackInfo__Item.Unmarshal(m, b)
}
func (m *TrackInfo__Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackInfo__Item.Marshal(b, m, deterministic)
}
func (m *TrackInfo__Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackInfo__Item.Merge(m, src)
}
func (m *TrackInfo__Item) XXX_Size() int {
	return xxx_messageInfo_TrackInfo__Item.Size(m)
}
func (m *TrackInfo__Item) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackInfo__Item.DiscardUnknown(m)
}

var xxx_messageInfo_TrackInfo__Item proto.InternalMessageInfo

func (m *TrackInfo__Item) GetTrackeName() string {
	if m != nil {
		return m.TrackeName
	}
	return ""
}

func (m *TrackInfo__Item) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TrackInfo__Item) GetTrackState() bool {
	if m != nil {
		return m.TrackState
	}
	return false
}

func (m *TrackInfo__Item) GetStateChangeCounter() uint32 {
	if m != nil {
		return m.StateChangeCounter
	}
	return 0
}

func (m *TrackInfo__Item) GetSecondsLastChange() uint64 {
	if m != nil {
		return m.SecondsLastChange
	}
	return 0
}

func (m *TrackInfo__Item) GetTrackTypeInfo() *TrackTypeInfoUnion {
	if m != nil {
		return m.TrackTypeInfo
	}
	return nil
}

func (m *TrackInfo__Item) GetBoolTracks() *BoolTrackInfo__Entry {
	if m != nil {
		return m.BoolTracks
	}
	return nil
}

func (m *TrackInfo__Item) GetThresholdTracks() *ThresholdTrackInfo__Entry {
	if m != nil {
		return m.ThresholdTracks
	}
	return nil
}

func (m *TrackInfo__Item) GetThresholdUp() uint32 {
	if m != nil {
		return m.ThresholdUp
	}
	return 0
}

func (m *TrackInfo__Item) GetThresholdDown() uint32 {
	if m != nil {
		return m.ThresholdDown
	}
	return 0
}

func (m *TrackInfo__Item) GetTrackingInteraces() *InterfaceTrackingInfo__Entry {
	if m != nil {
		return m.TrackingInteraces
	}
	return nil
}

func (m *TrackInfo__Item) GetDelayed() *DelayedStateStatus__ {
	if m != nil {
		return m.Delayed
	}
	return nil
}

type TrackInfo__Entry struct {
	TrackInfo__          []*TrackInfo__Item `protobuf:"bytes,50,rep,name=track_info__,json=trackInfo,proto3" json:"track_info__,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TrackInfo__Entry) Reset()         { *m = TrackInfo__Entry{} }
func (m *TrackInfo__Entry) String() string { return proto.CompactTextString(m) }
func (*TrackInfo__Entry) ProtoMessage()    {}
func (*TrackInfo__Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee7ca1ee6d013b, []int{14}
}

func (m *TrackInfo__Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackInfo__Entry.Unmarshal(m, b)
}
func (m *TrackInfo__Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackInfo__Entry.Marshal(b, m, deterministic)
}
func (m *TrackInfo__Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackInfo__Entry.Merge(m, src)
}
func (m *TrackInfo__Entry) XXX_Size() int {
	return xxx_messageInfo_TrackInfo__Entry.Size(m)
}
func (m *TrackInfo__Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackInfo__Entry.DiscardUnknown(m)
}

var xxx_messageInfo_TrackInfo__Entry proto.InternalMessageInfo

func (m *TrackInfo__Entry) GetTrackInfo__() []*TrackInfo__Item {
	if m != nil {
		return m.TrackInfo__
	}
	return nil
}

func init() {
	proto.RegisterType((*TrackInfo__Entry_KEYS)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.track_info___entry_KEYS")
	proto.RegisterType((*IntfTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.intf_track_info__")
	proto.RegisterType((*RouteTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.route_track_info__")
	proto.RegisterType((*RtrTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.rtr_track_info__")
	proto.RegisterType((*BfdrtrTrackInfo__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.bfdrtr_track_info__")
	proto.RegisterType((*TrackTypeInfoUnion)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.track_type_info_union")
	proto.RegisterType((*BoolTrackInfo__Item)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.bool_track_info___item")
	proto.RegisterType((*BoolTrackInfo__Entry)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.bool_track_info___entry")
	proto.RegisterType((*ThresholdTrackInfo__Item)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.threshold_track_info___item")
	proto.RegisterType((*ThresholdTrackInfo__Entry)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.threshold_track_info___entry")
	proto.RegisterType((*InterfaceTrackingInfo__Item)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.interface_tracking_info___item")
	proto.RegisterType((*InterfaceTrackingInfo__Entry)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.interface_tracking_info___entry")
	proto.RegisterType((*DelayedStateStatus__)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.delayed_state_status__")
	proto.RegisterType((*TrackInfo__Item)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.track_info___item")
	proto.RegisterType((*TrackInfo__Entry)(nil), "cisco_ios_xr_manageability_object_tracking_oper.object_tracking.track_type_ipv4_route.track_info___entry")
}

func init() { proto.RegisterFile("track_info___entry.proto", fileDescriptor_94ee7ca1ee6d013b) }

var fileDescriptor_94ee7ca1ee6d013b = []byte{
	// 976 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x41, 0x6f, 0x23, 0x35,
	0x14, 0x96, 0xb7, 0xa1, 0x4d, 0xde, 0x24, 0xb4, 0x71, 0xd9, 0xec, 0x54, 0xa0, 0xed, 0x32, 0x08,
	0x69, 0x4f, 0x01, 0x15, 0x4e, 0xdc, 0x50, 0x41, 0xb0, 0x62, 0xb5, 0x48, 0xde, 0xdd, 0x03, 0x27,
	0x33, 0x99, 0x71, 0x12, 0x2f, 0x89, 0x3d, 0xb2, 0x9d, 0xa6, 0xb9, 0x70, 0x45, 0x8b, 0x90, 0x10,
	0x02, 0x89, 0x1b, 0xe2, 0x02, 0xe2, 0xc6, 0x81, 0x1f, 0xc1, 0x0f, 0xe0, 0xd7, 0x70, 0x43, 0xf3,
	0xec, 0xa4, 0x99, 0x24, 0x5d, 0xf6, 0x32, 0xbd, 0x54, 0xf6, 0xf7, 0xc6, 0x7e, 0x9f, 0xbf, 0xbc,
	0xf7, 0x3d, 0x15, 0x62, 0x67, 0xd2, 0xec, 0x2b, 0x2e, 0xd5, 0x50, 0x73, 0xce, 0x85, 0x72, 0x66,
	0xd1, 0x2f, 0x8c, 0x76, 0x9a, 0x3e, 0xcd, 0xa4, 0xcd, 0x34, 0x97, 0xda, 0xf2, 0x4b, 0xc3, 0xa7,
	0xa9, 0x4a, 0x47, 0x22, 0x1d, 0xc8, 0x89, 0x74, 0x0b, 0xae, 0x07, 0xcf, 0x44, 0xe6, 0x38, 0x9e,
	0x95, 0x6a, 0xc4, 0x75, 0x21, 0x4c, 0x7f, 0x03, 0xec, 0xfb, 0x9b, 0xdd, 0xa2, 0x10, 0x5c, 0x16,
	0x17, 0xef, 0x73, 0xa3, 0x67, 0x4e, 0x24, 0x27, 0x70, 0x67, 0x3b, 0x25, 0xff, 0xec, 0xe3, 0x2f,
	0x1e, 0x27, 0x1f, 0x40, 0x57, 0x2a, 0x37, 0xe4, 0xeb, 0x71, 0xfa, 0x36, 0xbc, 0x2a, 0x95, 0x13,
	0x66, 0x98, 0x66, 0x82, 0xab, 0x74, 0x2a, 0x62, 0x72, 0x8f, 0xdc, 0x6f, 0xb1, 0xce, 0x0a, 0x7d,
	0x94, 0x4e, 0x45, 0xf2, 0x35, 0x50, 0xbc, 0xbf, 0x7a, 0xb8, 0x07, 0xfb, 0x85, 0x11, 0x43, 0x79,
	0x89, 0x87, 0x3a, 0x2c, 0xec, 0xe8, 0x5b, 0xd0, 0xf1, 0x2b, 0x3e, 0x11, 0x6a, 0xe4, 0xc6, 0xf1,
	0x2d, 0x0c, 0xb7, 0x3d, 0xf8, 0x10, 0x31, 0x7a, 0x04, 0x7b, 0x17, 0x66, 0x18, 0xef, 0x61, 0xba,
	0x72, 0x49, 0x4f, 0xa0, 0xa9, 0xc4, 0xa5, 0xe3, 0x63, 0x5d, 0xc4, 0x0d, 0x84, 0x0f, 0xca, 0xfd,
	0xa7, 0xba, 0x48, 0x04, 0x1c, 0x19, 0x67, 0xaa, 0xd9, 0xef, 0x42, 0x24, 0x0b, 0x3b, 0x49, 0xb9,
	0x2e, 0xb8, 0xcc, 0x03, 0x85, 0x16, 0x42, 0x9f, 0x17, 0x0f, 0xf2, 0x32, 0x81, 0x71, 0x2e, 0xe4,
	0x2e, 0x97, 0xf4, 0x14, 0x22, 0x23, 0xdc, 0xcc, 0x28, 0x9e, 0xe9, 0x5c, 0x60, 0xea, 0x0e, 0x03,
	0x0f, 0x9d, 0xeb, 0x5c, 0x24, 0x7f, 0x10, 0x38, 0x1e, 0x0c, 0xf3, 0xad, 0x54, 0x2f, 0xa7, 0x12,
	0x7d, 0x07, 0x8e, 0x73, 0x61, 0x9d, 0x54, 0xa9, 0x93, 0x5a, 0xf1, 0x34, 0xcf, 0x8d, 0xb0, 0x36,
	0x30, 0xa0, 0x6b, 0xa1, 0x0f, 0x7d, 0x84, 0x52, 0x68, 0x98, 0xd4, 0x2d, 0x99, 0xe0, 0xba, 0xcc,
	0x95, 0x8b, 0x81, 0x9e, 0xa9, 0x4c, 0xf0, 0x4c, 0xcf, 0x94, 0x43, 0x2d, 0x3a, 0xac, 0xb3, 0x44,
	0xcf, 0x4b, 0x30, 0xf9, 0xb7, 0x01, 0xb7, 0xd7, 0x4b, 0xa0, 0x24, 0x3a, 0x53, 0x52, 0x2b, 0x9a,
	0x40, 0x3b, 0x97, 0x36, 0x33, 0x72, 0x2a, 0x55, 0xaa, 0x5c, 0xa0, 0x5a, 0xc1, 0xe8, 0x8f, 0x04,
	0x8e, 0xae, 0x5e, 0x84, 0xf7, 0x78, 0x9e, 0xd1, 0xd9, 0xb8, 0x5f, 0x4b, 0x65, 0xf6, 0xb7, 0x6a,
	0x8f, 0x1d, 0xae, 0x18, 0x3c, 0x41, 0x02, 0xf4, 0x3b, 0x02, 0xed, 0xb5, 0x32, 0xb3, 0xa8, 0x4b,
	0x74, 0x26, 0x6b, 0x62, 0xb4, 0x5d, 0xd1, 0x2c, 0x42, 0x2c, 0xd0, 0xf9, 0x96, 0x40, 0xdb, 0x57,
	0x58, 0xa0, 0xd3, 0x40, 0x3a, 0xa3, 0xba, 0xe8, 0x6c, 0x54, 0x1d, 0xf3, 0xe5, 0x1d, 0xc8, 0x3c,
	0x27, 0x00, 0x83, 0x61, 0xbe, 0xa4, 0xf2, 0x0a, 0x52, 0x79, 0x56, 0x13, 0x95, 0x1d, 0x3d, 0xc0,
	0x5a, 0x83, 0x61, 0xee, 0xb9, 0x24, 0x33, 0xe8, 0x0d, 0xb4, 0x9e, 0x54, 0xe2, 0x5c, 0x3a, 0x31,
	0x2d, 0x3b, 0x2c, 0x64, 0x58, 0xeb, 0x12, 0xf0, 0x10, 0xb6, 0xc8, 0x29, 0x44, 0xfe, 0x94, 0x75,
	0x65, 0xe1, 0x97, 0x25, 0xd7, 0x64, 0x80, 0xd0, 0xe3, 0x12, 0x29, 0x4d, 0x60, 0x2e, 0xdd, 0x98,
	0x2b, 0xed, 0xf0, 0xe7, 0x6f, 0xb2, 0x83, 0x72, 0xff, 0x48, 0xbb, 0xe4, 0x2f, 0x02, 0x77, 0xb6,
	0xf3, 0xa2, 0xc3, 0xd1, 0x9f, 0x09, 0x74, 0xb7, 0x62, 0x71, 0xe3, 0xde, 0xde, 0xfd, 0xe8, 0x6c,
	0x5a, 0x97, 0x4a, 0x3b, 0x35, 0x60, 0x9d, 0x12, 0x47, 0xa5, 0x1e, 0xa8, 0xa1, 0x4e, 0xe6, 0xf0,
	0xba, 0x1b, 0x1b, 0x61, 0xc7, 0x7a, 0x92, 0xd7, 0xa2, 0x58, 0x0f, 0xf6, 0xe7, 0x42, 0x8e, 0xc6,
	0x2e, 0xd8, 0x48, 0xd8, 0x25, 0x7f, 0x13, 0x78, 0xe3, 0x9a, 0xcc, 0x5e, 0xb3, 0xdf, 0x09, 0xf4,
	0x76, 0x7f, 0x10, 0x84, 0x33, 0x35, 0x09, 0xf7, 0x02, 0x3d, 0x18, 0x5d, 0x05, 0xaf, 0x24, 0xfc,
	0x04, 0xee, 0x6e, 0x98, 0x55, 0x99, 0x70, 0x5d, 0xc5, 0x97, 0x1c, 0x63, 0xff, 0x10, 0x38, 0xbd,
	0xfe, 0x26, 0xaf, 0xca, 0x9f, 0x04, 0x4e, 0xae, 0xfd, 0x26, 0xbe, 0x85, 0xc2, 0xcc, 0xea, 0xf3,
	0xc8, 0x17, 0xbc, 0x92, 0xdd, 0xae, 0x1a, 0xa6, 0x54, 0x23, 0x94, 0xe7, 0x4b, 0xe8, 0xe5, 0x62,
	0x92, 0x2e, 0x44, 0xee, 0x6b, 0x04, 0xff, 0xce, 0xac, 0x9f, 0x5b, 0x4e, 0x4e, 0x05, 0x37, 0x62,
	0x9a, 0x4a, 0x25, 0xd5, 0x28, 0x4c, 0xc9, 0x4e, 0x89, 0xb2, 0x25, 0xf8, 0xbf, 0x25, 0x96, 0xfc,
	0xd0, 0x84, 0xee, 0xce, 0xd2, 0x45, 0xb0, 0xa2, 0xb8, 0x3f, 0xe6, 0xe7, 0x21, 0x85, 0x46, 0xf9,
	0x42, 0xbc, 0xb0, 0xc5, 0x70, 0xbd, 0x99, 0x6b, 0x6f, 0xab, 0x9c, 0xdf, 0x85, 0xd7, 0xfc, 0x2b,
	0xb2, 0x71, 0xaa, 0x46, 0x61, 0x06, 0x0a, 0x13, 0xa6, 0x20, 0xc5, 0xd8, 0x39, 0x86, 0xce, 0x7d,
	0x84, 0xf6, 0xe1, 0xd8, 0x8a, 0x4c, 0xab, 0xdc, 0xf2, 0x49, 0x6a, 0x5d, 0x38, 0x88, 0x16, 0xd9,
	0x60, 0xdd, 0x10, 0x7a, 0x98, 0x5a, 0xe7, 0x8f, 0xd1, 0x9f, 0x08, 0x1c, 0x6e, 0x8c, 0xce, 0x78,
	0x1f, 0xfd, 0x74, 0x52, 0x57, 0xc1, 0xef, 0x1a, 0xd4, 0xac, 0x83, 0xf0, 0x93, 0x45, 0x21, 0xca,
	0x9f, 0x91, 0x7e, 0x4f, 0x20, 0xba, 0xb2, 0x14, 0x1b, 0x1f, 0x20, 0x25, 0x75, 0x63, 0xe6, 0x85,
	0xe5, 0xcf, 0x60, 0xe5, 0x5e, 0x96, 0xfe, 0x42, 0xe0, 0x68, 0xa3, 0x57, 0x6d, 0xdc, 0x44, 0x5a,
	0xf6, 0x66, 0xad, 0xc1, 0x73, 0x3b, 0xac, 0x7a, 0x83, 0xa5, 0x6f, 0x42, 0xfb, 0xea, 0xc0, 0xac,
	0x88, 0x5b, 0x58, 0x23, 0xd1, 0x0a, 0x7b, 0x5a, 0x60, 0x0b, 0xac, 0x3e, 0xc9, 0xf5, 0x5c, 0xc5,
	0x10, 0x5a, 0x60, 0x89, 0x7e, 0xa4, 0xe7, 0x8a, 0xfe, 0x46, 0x80, 0xae, 0xf5, 0x9c, 0x13, 0x26,
	0xcd, 0x84, 0x8d, 0x23, 0x7c, 0xec, 0xc5, 0x8d, 0xb7, 0xbb, 0x7f, 0x6f, 0xd7, 0xad, 0xda, 0x3c,
	0x10, 0xa2, 0xdf, 0x10, 0x38, 0x08, 0xcd, 0x1e, 0xb7, 0x91, 0x5c, 0x5d, 0xd3, 0x6d, 0xb7, 0xa5,
	0xb0, 0x65, 0xf6, 0xe4, 0xd7, 0xa5, 0x62, 0x55, 0xfb, 0x7c, 0x4e, 0xa0, 0x5d, 0x19, 0x25, 0x67,
	0xe8, 0x98, 0xe3, 0x5a, 0x3b, 0x6b, 0xdd, 0x24, 0x5b, 0x6e, 0x39, 0x37, 0x06, 0xfb, 0xf8, 0xaf,
	0xd6, 0x7b, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xce, 0x30, 0x65, 0x86, 0x0d, 0x00, 0x00,
}
