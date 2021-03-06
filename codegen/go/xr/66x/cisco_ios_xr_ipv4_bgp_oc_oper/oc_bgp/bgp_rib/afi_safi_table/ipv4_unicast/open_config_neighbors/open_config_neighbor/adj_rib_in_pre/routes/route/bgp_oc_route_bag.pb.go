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
// source: bgp_oc_route_bag.proto

package cisco_ios_xr_ipv4_bgp_oc_oper_oc_bgp_bgp_rib_afi_safi_table_ipv4_unicast_open_config_neighbors_open_config_neighbor_adj_rib_in_pre_routes_route

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

type BgpOcRouteBag_KEYS struct {
	NeighborAddress      string   `protobuf:"bytes,1,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	Route                string   `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	NeighborAddress_1    string   `protobuf:"bytes,3,opt,name=neighbor_address_1,json=neighborAddress1,proto3" json:"neighbor_address_1,omitempty"`
	PathId               uint32   `protobuf:"varint,4,opt,name=path_id,json=pathId,proto3" json:"path_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpOcRouteBag_KEYS) Reset()         { *m = BgpOcRouteBag_KEYS{} }
func (m *BgpOcRouteBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpOcRouteBag_KEYS) ProtoMessage()    {}
func (*BgpOcRouteBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7038c6d4a30a63, []int{0}
}

func (m *BgpOcRouteBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcRouteBag_KEYS.Unmarshal(m, b)
}
func (m *BgpOcRouteBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcRouteBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpOcRouteBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcRouteBag_KEYS.Merge(m, src)
}
func (m *BgpOcRouteBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpOcRouteBag_KEYS.Size(m)
}
func (m *BgpOcRouteBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcRouteBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcRouteBag_KEYS proto.InternalMessageInfo

func (m *BgpOcRouteBag_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *BgpOcRouteBag_KEYS) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *BgpOcRouteBag_KEYS) GetNeighborAddress_1() string {
	if m != nil {
		return m.NeighborAddress_1
	}
	return ""
}

func (m *BgpOcRouteBag_KEYS) GetPathId() uint32 {
	if m != nil {
		return m.PathId
	}
	return 0
}

type BgpOcAddrtype struct {
	Afi                  string   `protobuf:"bytes,1,opt,name=afi,proto3" json:"afi,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpOcAddrtype) Reset()         { *m = BgpOcAddrtype{} }
func (m *BgpOcAddrtype) String() string { return proto.CompactTextString(m) }
func (*BgpOcAddrtype) ProtoMessage()    {}
func (*BgpOcAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7038c6d4a30a63, []int{1}
}

func (m *BgpOcAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcAddrtype.Unmarshal(m, b)
}
func (m *BgpOcAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcAddrtype.Marshal(b, m, deterministic)
}
func (m *BgpOcAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcAddrtype.Merge(m, src)
}
func (m *BgpOcAddrtype) XXX_Size() int {
	return xxx_messageInfo_BgpOcAddrtype.Size(m)
}
func (m *BgpOcAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcAddrtype proto.InternalMessageInfo

func (m *BgpOcAddrtype) GetAfi() string {
	if m != nil {
		return m.Afi
	}
	return ""
}

func (m *BgpOcAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *BgpOcAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type BgpOcPrefixtype struct {
	Prefix               *BgpOcAddrtype `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32         `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BgpOcPrefixtype) Reset()         { *m = BgpOcPrefixtype{} }
func (m *BgpOcPrefixtype) String() string { return proto.CompactTextString(m) }
func (*BgpOcPrefixtype) ProtoMessage()    {}
func (*BgpOcPrefixtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7038c6d4a30a63, []int{2}
}

func (m *BgpOcPrefixtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcPrefixtype.Unmarshal(m, b)
}
func (m *BgpOcPrefixtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcPrefixtype.Marshal(b, m, deterministic)
}
func (m *BgpOcPrefixtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcPrefixtype.Merge(m, src)
}
func (m *BgpOcPrefixtype) XXX_Size() int {
	return xxx_messageInfo_BgpOcPrefixtype.Size(m)
}
func (m *BgpOcPrefixtype) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcPrefixtype.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcPrefixtype proto.InternalMessageInfo

func (m *BgpOcPrefixtype) GetPrefix() *BgpOcAddrtype {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *BgpOcPrefixtype) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type BgpOcAggregatorType struct {
	As                   uint32   `protobuf:"varint,1,opt,name=as,proto3" json:"as,omitempty"`
	As4                  uint32   `protobuf:"varint,2,opt,name=as4,proto3" json:"as4,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpOcAggregatorType) Reset()         { *m = BgpOcAggregatorType{} }
func (m *BgpOcAggregatorType) String() string { return proto.CompactTextString(m) }
func (*BgpOcAggregatorType) ProtoMessage()    {}
func (*BgpOcAggregatorType) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7038c6d4a30a63, []int{3}
}

func (m *BgpOcAggregatorType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcAggregatorType.Unmarshal(m, b)
}
func (m *BgpOcAggregatorType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcAggregatorType.Marshal(b, m, deterministic)
}
func (m *BgpOcAggregatorType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcAggregatorType.Merge(m, src)
}
func (m *BgpOcAggregatorType) XXX_Size() int {
	return xxx_messageInfo_BgpOcAggregatorType.Size(m)
}
func (m *BgpOcAggregatorType) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcAggregatorType.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcAggregatorType proto.InternalMessageInfo

func (m *BgpOcAggregatorType) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

func (m *BgpOcAggregatorType) GetAs4() uint32 {
	if m != nil {
		return m.As4
	}
	return 0
}

func (m *BgpOcAggregatorType) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type BgpOcStringType struct {
	Objects              string   `protobuf:"bytes,1,opt,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpOcStringType) Reset()         { *m = BgpOcStringType{} }
func (m *BgpOcStringType) String() string { return proto.CompactTextString(m) }
func (*BgpOcStringType) ProtoMessage()    {}
func (*BgpOcStringType) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7038c6d4a30a63, []int{4}
}

func (m *BgpOcStringType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcStringType.Unmarshal(m, b)
}
func (m *BgpOcStringType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcStringType.Marshal(b, m, deterministic)
}
func (m *BgpOcStringType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcStringType.Merge(m, src)
}
func (m *BgpOcStringType) XXX_Size() int {
	return xxx_messageInfo_BgpOcStringType.Size(m)
}
func (m *BgpOcStringType) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcStringType.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcStringType proto.InternalMessageInfo

func (m *BgpOcStringType) GetObjects() string {
	if m != nil {
		return m.Objects
	}
	return ""
}

type BgpOcRouteAttr struct {
	OriginType            string               `protobuf:"bytes,1,opt,name=origin_type,json=originType,proto3" json:"origin_type,omitempty"`
	AsPath                string               `protobuf:"bytes,2,opt,name=as_path,json=asPath,proto3" json:"as_path,omitempty"`
	As4Path               string               `protobuf:"bytes,3,opt,name=as4_path,json=as4Path,proto3" json:"as4_path,omitempty"`
	NextHop               *BgpOcAddrtype       `protobuf:"bytes,4,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	Med                   uint32               `protobuf:"varint,5,opt,name=med,proto3" json:"med,omitempty"`
	LocalPref             uint32               `protobuf:"varint,6,opt,name=local_pref,json=localPref,proto3" json:"local_pref,omitempty"`
	AtomicAggr            bool                 `protobuf:"varint,7,opt,name=atomic_aggr,json=atomicAggr,proto3" json:"atomic_aggr,omitempty"`
	AggregratorAttributes *BgpOcAggregatorType `protobuf:"bytes,8,opt,name=aggregrator_attributes,json=aggregratorAttributes,proto3" json:"aggregrator_attributes,omitempty"`
	Community             []*BgpOcStringType   `protobuf:"bytes,9,rep,name=community,proto3" json:"community,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-"`
	XXX_unrecognized      []byte               `json:"-"`
	XXX_sizecache         int32                `json:"-"`
}

func (m *BgpOcRouteAttr) Reset()         { *m = BgpOcRouteAttr{} }
func (m *BgpOcRouteAttr) String() string { return proto.CompactTextString(m) }
func (*BgpOcRouteAttr) ProtoMessage()    {}
func (*BgpOcRouteAttr) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7038c6d4a30a63, []int{5}
}

func (m *BgpOcRouteAttr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcRouteAttr.Unmarshal(m, b)
}
func (m *BgpOcRouteAttr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcRouteAttr.Marshal(b, m, deterministic)
}
func (m *BgpOcRouteAttr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcRouteAttr.Merge(m, src)
}
func (m *BgpOcRouteAttr) XXX_Size() int {
	return xxx_messageInfo_BgpOcRouteAttr.Size(m)
}
func (m *BgpOcRouteAttr) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcRouteAttr.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcRouteAttr proto.InternalMessageInfo

func (m *BgpOcRouteAttr) GetOriginType() string {
	if m != nil {
		return m.OriginType
	}
	return ""
}

func (m *BgpOcRouteAttr) GetAsPath() string {
	if m != nil {
		return m.AsPath
	}
	return ""
}

func (m *BgpOcRouteAttr) GetAs4Path() string {
	if m != nil {
		return m.As4Path
	}
	return ""
}

func (m *BgpOcRouteAttr) GetNextHop() *BgpOcAddrtype {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *BgpOcRouteAttr) GetMed() uint32 {
	if m != nil {
		return m.Med
	}
	return 0
}

func (m *BgpOcRouteAttr) GetLocalPref() uint32 {
	if m != nil {
		return m.LocalPref
	}
	return 0
}

func (m *BgpOcRouteAttr) GetAtomicAggr() bool {
	if m != nil {
		return m.AtomicAggr
	}
	return false
}

func (m *BgpOcRouteAttr) GetAggregratorAttributes() *BgpOcAggregatorType {
	if m != nil {
		return m.AggregratorAttributes
	}
	return nil
}

func (m *BgpOcRouteAttr) GetCommunity() []*BgpOcStringType {
	if m != nil {
		return m.Community
	}
	return nil
}

type BgpOcUnknownAttrType struct {
	AttributeType        uint32   `protobuf:"varint,1,opt,name=attribute_type,json=attributeType,proto3" json:"attribute_type,omitempty"`
	AttributeLength      uint32   `protobuf:"varint,2,opt,name=attribute_length,json=attributeLength,proto3" json:"attribute_length,omitempty"`
	AttributeValue       string   `protobuf:"bytes,3,opt,name=attribute_value,json=attributeValue,proto3" json:"attribute_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpOcUnknownAttrType) Reset()         { *m = BgpOcUnknownAttrType{} }
func (m *BgpOcUnknownAttrType) String() string { return proto.CompactTextString(m) }
func (*BgpOcUnknownAttrType) ProtoMessage()    {}
func (*BgpOcUnknownAttrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7038c6d4a30a63, []int{6}
}

func (m *BgpOcUnknownAttrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcUnknownAttrType.Unmarshal(m, b)
}
func (m *BgpOcUnknownAttrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcUnknownAttrType.Marshal(b, m, deterministic)
}
func (m *BgpOcUnknownAttrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcUnknownAttrType.Merge(m, src)
}
func (m *BgpOcUnknownAttrType) XXX_Size() int {
	return xxx_messageInfo_BgpOcUnknownAttrType.Size(m)
}
func (m *BgpOcUnknownAttrType) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcUnknownAttrType.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcUnknownAttrType proto.InternalMessageInfo

func (m *BgpOcUnknownAttrType) GetAttributeType() uint32 {
	if m != nil {
		return m.AttributeType
	}
	return 0
}

func (m *BgpOcUnknownAttrType) GetAttributeLength() uint32 {
	if m != nil {
		return m.AttributeLength
	}
	return 0
}

func (m *BgpOcUnknownAttrType) GetAttributeValue() string {
	if m != nil {
		return m.AttributeValue
	}
	return ""
}

type BgpOcExtAttr struct {
	OriginatorId         string                  `protobuf:"bytes,1,opt,name=originator_id,json=originatorId,proto3" json:"originator_id,omitempty"`
	Cluster              []string                `protobuf:"bytes,2,rep,name=cluster,proto3" json:"cluster,omitempty"`
	ExtCommunity         []*BgpOcStringType      `protobuf:"bytes,3,rep,name=ext_community,json=extCommunity,proto3" json:"ext_community,omitempty"`
	Aigp                 uint64                  `protobuf:"varint,4,opt,name=aigp,proto3" json:"aigp,omitempty"`
	PathId               uint32                  `protobuf:"varint,5,opt,name=path_id,json=pathId,proto3" json:"path_id,omitempty"`
	UnknownAttributes    []*BgpOcUnknownAttrType `protobuf:"bytes,6,rep,name=unknown_attributes,json=unknownAttributes,proto3" json:"unknown_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *BgpOcExtAttr) Reset()         { *m = BgpOcExtAttr{} }
func (m *BgpOcExtAttr) String() string { return proto.CompactTextString(m) }
func (*BgpOcExtAttr) ProtoMessage()    {}
func (*BgpOcExtAttr) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7038c6d4a30a63, []int{7}
}

func (m *BgpOcExtAttr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcExtAttr.Unmarshal(m, b)
}
func (m *BgpOcExtAttr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcExtAttr.Marshal(b, m, deterministic)
}
func (m *BgpOcExtAttr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcExtAttr.Merge(m, src)
}
func (m *BgpOcExtAttr) XXX_Size() int {
	return xxx_messageInfo_BgpOcExtAttr.Size(m)
}
func (m *BgpOcExtAttr) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcExtAttr.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcExtAttr proto.InternalMessageInfo

func (m *BgpOcExtAttr) GetOriginatorId() string {
	if m != nil {
		return m.OriginatorId
	}
	return ""
}

func (m *BgpOcExtAttr) GetCluster() []string {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *BgpOcExtAttr) GetExtCommunity() []*BgpOcStringType {
	if m != nil {
		return m.ExtCommunity
	}
	return nil
}

func (m *BgpOcExtAttr) GetAigp() uint64 {
	if m != nil {
		return m.Aigp
	}
	return 0
}

func (m *BgpOcExtAttr) GetPathId() uint32 {
	if m != nil {
		return m.PathId
	}
	return 0
}

func (m *BgpOcExtAttr) GetUnknownAttributes() []*BgpOcUnknownAttrType {
	if m != nil {
		return m.UnknownAttributes
	}
	return nil
}

type BgpOcDateAndTime struct {
	TimeValue            string   `protobuf:"bytes,1,opt,name=time_value,json=timeValue,proto3" json:"time_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpOcDateAndTime) Reset()         { *m = BgpOcDateAndTime{} }
func (m *BgpOcDateAndTime) String() string { return proto.CompactTextString(m) }
func (*BgpOcDateAndTime) ProtoMessage()    {}
func (*BgpOcDateAndTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7038c6d4a30a63, []int{8}
}

func (m *BgpOcDateAndTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcDateAndTime.Unmarshal(m, b)
}
func (m *BgpOcDateAndTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcDateAndTime.Marshal(b, m, deterministic)
}
func (m *BgpOcDateAndTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcDateAndTime.Merge(m, src)
}
func (m *BgpOcDateAndTime) XXX_Size() int {
	return xxx_messageInfo_BgpOcDateAndTime.Size(m)
}
func (m *BgpOcDateAndTime) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcDateAndTime.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcDateAndTime proto.InternalMessageInfo

func (m *BgpOcDateAndTime) GetTimeValue() string {
	if m != nil {
		return m.TimeValue
	}
	return ""
}

type BgpOcRouteBag struct {
	PrefixName           *BgpOcPrefixtype  `protobuf:"bytes,50,opt,name=prefix_name,json=prefixName,proto3" json:"prefix_name,omitempty"`
	RouteAttrList        *BgpOcRouteAttr   `protobuf:"bytes,51,opt,name=route_attr_list,json=routeAttrList,proto3" json:"route_attr_list,omitempty"`
	ExtAttributesList    *BgpOcExtAttr     `protobuf:"bytes,52,opt,name=ext_attributes_list,json=extAttributesList,proto3" json:"ext_attributes_list,omitempty"`
	LastModifiedDate     *BgpOcDateAndTime `protobuf:"bytes,53,opt,name=last_modified_date,json=lastModifiedDate,proto3" json:"last_modified_date,omitempty"`
	LastUpdateRecieved   *BgpOcDateAndTime `protobuf:"bytes,54,opt,name=last_update_recieved,json=lastUpdateRecieved,proto3" json:"last_update_recieved,omitempty"`
	ValidRoute           bool              `protobuf:"varint,55,opt,name=valid_route,json=validRoute,proto3" json:"valid_route,omitempty"`
	InvalidReason        string            `protobuf:"bytes,56,opt,name=invalid_reason,json=invalidReason,proto3" json:"invalid_reason,omitempty"`
	BestPath             bool              `protobuf:"varint,57,opt,name=best_path,json=bestPath,proto3" json:"best_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BgpOcRouteBag) Reset()         { *m = BgpOcRouteBag{} }
func (m *BgpOcRouteBag) String() string { return proto.CompactTextString(m) }
func (*BgpOcRouteBag) ProtoMessage()    {}
func (*BgpOcRouteBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7038c6d4a30a63, []int{9}
}

func (m *BgpOcRouteBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcRouteBag.Unmarshal(m, b)
}
func (m *BgpOcRouteBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcRouteBag.Marshal(b, m, deterministic)
}
func (m *BgpOcRouteBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcRouteBag.Merge(m, src)
}
func (m *BgpOcRouteBag) XXX_Size() int {
	return xxx_messageInfo_BgpOcRouteBag.Size(m)
}
func (m *BgpOcRouteBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcRouteBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcRouteBag proto.InternalMessageInfo

func (m *BgpOcRouteBag) GetPrefixName() *BgpOcPrefixtype {
	if m != nil {
		return m.PrefixName
	}
	return nil
}

func (m *BgpOcRouteBag) GetRouteAttrList() *BgpOcRouteAttr {
	if m != nil {
		return m.RouteAttrList
	}
	return nil
}

func (m *BgpOcRouteBag) GetExtAttributesList() *BgpOcExtAttr {
	if m != nil {
		return m.ExtAttributesList
	}
	return nil
}

func (m *BgpOcRouteBag) GetLastModifiedDate() *BgpOcDateAndTime {
	if m != nil {
		return m.LastModifiedDate
	}
	return nil
}

func (m *BgpOcRouteBag) GetLastUpdateRecieved() *BgpOcDateAndTime {
	if m != nil {
		return m.LastUpdateRecieved
	}
	return nil
}

func (m *BgpOcRouteBag) GetValidRoute() bool {
	if m != nil {
		return m.ValidRoute
	}
	return false
}

func (m *BgpOcRouteBag) GetInvalidReason() string {
	if m != nil {
		return m.InvalidReason
	}
	return ""
}

func (m *BgpOcRouteBag) GetBestPath() bool {
	if m != nil {
		return m.BestPath
	}
	return false
}

func init() {
	proto.RegisterType((*BgpOcRouteBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.routes.route.bgp_oc_route_bag_KEYS")
	proto.RegisterType((*BgpOcAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.routes.route.bgp_oc_addrtype")
	proto.RegisterType((*BgpOcPrefixtype)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.routes.route.bgp_oc_prefixtype")
	proto.RegisterType((*BgpOcAggregatorType)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.routes.route.bgp_oc_aggregator_type")
	proto.RegisterType((*BgpOcStringType)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.routes.route.bgp_oc_string_type")
	proto.RegisterType((*BgpOcRouteAttr)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.routes.route.bgp_oc_route_attr")
	proto.RegisterType((*BgpOcUnknownAttrType)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.routes.route.bgp_oc_unknown_attr_type")
	proto.RegisterType((*BgpOcExtAttr)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.routes.route.bgp_oc_ext_attr")
	proto.RegisterType((*BgpOcDateAndTime)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.routes.route.bgp_oc_date_and_time")
	proto.RegisterType((*BgpOcRouteBag)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.routes.route.bgp_oc_route_bag")
}

func init() { proto.RegisterFile("bgp_oc_route_bag.proto", fileDescriptor_4a7038c6d4a30a63) }

var fileDescriptor_4a7038c6d4a30a63 = []byte{
	// 981 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xd1, 0x8a, 0x23, 0x45,
	0x17, 0xa6, 0x26, 0x33, 0xc9, 0xe4, 0x64, 0x7a, 0x26, 0x53, 0xff, 0xec, 0x6e, 0xff, 0x88, 0x18,
	0xb3, 0x2c, 0xce, 0x82, 0x34, 0xec, 0x6c, 0x76, 0xd4, 0xcb, 0x41, 0x05, 0x17, 0x57, 0x59, 0xda,
	0x55, 0xf0, 0xaa, 0xa8, 0x74, 0x57, 0x3a, 0xb5, 0x26, 0x5d, 0x4d, 0x55, 0x65, 0xcc, 0x3e, 0x85,
	0x20, 0x78, 0xed, 0xa2, 0x2c, 0xe2, 0x23, 0x08, 0x5e, 0x89, 0x37, 0xe2, 0x95, 0xe0, 0x85, 0x5e,
	0xf9, 0x12, 0x3e, 0x80, 0xd4, 0xa9, 0xea, 0x74, 0x26, 0xfa, 0x00, 0x99, 0x9b, 0x99, 0xaa, 0xaf,
	0x4f, 0x9f, 0x3e, 0xdf, 0xf9, 0xce, 0x39, 0x55, 0x81, 0x9b, 0xe3, 0xa2, 0x62, 0x2a, 0x63, 0x5a,
	0x2d, 0xac, 0x60, 0x63, 0x5e, 0x24, 0x95, 0x56, 0x56, 0xd1, 0x2f, 0x48, 0x26, 0x4d, 0xa6, 0x98,
	0x54, 0x86, 0x2d, 0x35, 0x93, 0xd5, 0xe5, 0x88, 0x05, 0x53, 0x55, 0x09, 0x9d, 0xa8, 0xcc, 0x6d,
	0x13, 0x07, 0x69, 0x39, 0x4e, 0xf8, 0x44, 0x32, 0xe3, 0xfe, 0x58, 0x3e, 0x9e, 0x89, 0x04, 0xcd,
	0x17, 0xa5, 0xcc, 0xb8, 0xb1, 0x89, 0xaa, 0x44, 0xc9, 0x32, 0x55, 0x4e, 0x64, 0xc1, 0x4a, 0x21,
	0x8b, 0xe9, 0x58, 0x69, 0xf3, 0x9f, 0x68, 0xc2, 0xf3, 0xa7, 0xce, 0x1d, 0x93, 0x25, 0xab, 0xb4,
	0x48, 0x30, 0x28, 0xe3, 0xff, 0x0d, 0xbf, 0x26, 0x70, 0x63, 0x33, 0x58, 0xf6, 0xfe, 0xbb, 0x9f,
	0x7e, 0x44, 0xef, 0x42, 0xbf, 0xf6, 0xc0, 0x78, 0x9e, 0x6b, 0x61, 0x4c, 0x4c, 0x06, 0xe4, 0xb4,
	0x9b, 0x1e, 0xd5, 0xf8, 0x85, 0x87, 0xe9, 0x09, 0xec, 0xe1, 0xcb, 0xf1, 0x0e, 0x3e, 0xf7, 0x1b,
	0xfa, 0x3a, 0xd0, 0x4d, 0x07, 0xec, 0x5e, 0xdc, 0x42, 0x93, 0xfe, 0x86, 0x8b, 0x7b, 0xf4, 0x16,
	0x74, 0x2a, 0x6e, 0xa7, 0x4c, 0xe6, 0xf1, 0xee, 0x80, 0x9c, 0x46, 0x69, 0xdb, 0x6d, 0x1f, 0xe6,
	0x43, 0x09, 0x47, 0x21, 0x40, 0xe7, 0xc4, 0x3e, 0xab, 0x04, 0xed, 0x43, 0x8b, 0x4f, 0x64, 0x88,
	0xc6, 0x2d, 0xe9, 0xab, 0x70, 0x80, 0xd9, 0xa9, 0x03, 0xf5, 0x81, 0xf4, 0x1c, 0x56, 0x07, 0xe9,
	0x4d, 0xce, 0x57, 0x26, 0xad, 0x95, 0xc9, 0x79, 0x30, 0x19, 0xfe, 0x4d, 0xe0, 0x38, 0x7c, 0xab,
	0xd2, 0x62, 0x22, 0x97, 0xf8, 0xb5, 0x1f, 0x08, 0xb4, 0xfd, 0x16, 0xbf, 0xd8, 0x3b, 0x7b, 0x4e,
	0x92, 0x2d, 0x93, 0x31, 0xd9, 0xc8, 0x50, 0x1a, 0x02, 0xa6, 0xb7, 0x21, 0xf2, 0x2b, 0x36, 0x13,
	0x65, 0x61, 0xa7, 0x98, 0x98, 0x28, 0x3d, 0xf0, 0xe0, 0x23, 0xc4, 0x86, 0x4f, 0x56, 0xf5, 0xca,
	0x8b, 0x42, 0x8b, 0x82, 0x5b, 0xa5, 0x19, 0x52, 0x3f, 0x84, 0x1d, 0xee, 0x55, 0x8f, 0xd2, 0x1d,
	0x6e, 0x30, 0xf1, 0x66, 0x14, 0x9c, 0xb8, 0x25, 0x8d, 0xa1, 0x73, 0x35, 0xa1, 0xf5, 0x76, 0x98,
	0x00, 0x0d, 0x5e, 0x8d, 0xd5, 0xb2, 0x2c, 0xbc, 0xc7, 0x18, 0x3a, 0x6a, 0xfc, 0x54, 0x64, 0xb6,
	0x2e, 0xa6, 0x7a, 0x3b, 0x7c, 0xde, 0x5e, 0x25, 0xdf, 0x57, 0x22, 0xb7, 0x56, 0xd3, 0x57, 0xa0,
	0xa7, 0xb4, 0x2c, 0x64, 0x89, 0xaf, 0x87, 0x77, 0xc0, 0x43, 0x4f, 0x9c, 0xc3, 0x5b, 0xd0, 0xe1,
	0x86, 0xb9, 0x5a, 0x09, 0xa2, 0xb7, 0xb9, 0x79, 0xcc, 0xed, 0x94, 0xfe, 0x1f, 0xf6, 0xb9, 0x19,
	0xf9, 0x27, 0x75, 0x68, 0x66, 0x84, 0x8f, 0x7e, 0x24, 0xb0, 0x5f, 0x8a, 0xa5, 0x65, 0x53, 0x55,
	0x61, 0xb5, 0x5d, 0x0b, 0x4d, 0x3b, 0x2e, 0xe4, 0xf7, 0x54, 0xe5, 0x54, 0x98, 0x8b, 0x3c, 0xde,
	0xf3, 0x2a, 0xcc, 0x45, 0x4e, 0x5f, 0x06, 0x98, 0xa9, 0x8c, 0xcf, 0xb0, 0x6c, 0xe3, 0x36, 0x3e,
	0xe8, 0x22, 0xf2, 0x58, 0x8b, 0x89, 0x4b, 0x22, 0xb7, 0x6a, 0x2e, 0xbd, 0xc0, 0x71, 0x67, 0x40,
	0x4e, 0xf7, 0x53, 0xf0, 0xd0, 0x45, 0x51, 0x68, 0xfa, 0x17, 0x81, 0x9b, 0x5e, 0x7b, 0x8d, 0xe2,
	0xbb, 0xd4, 0xcb, 0xb1, 0x8b, 0x25, 0xde, 0xc7, 0xf4, 0x7c, 0xb7, 0xbd, 0xe9, 0xb9, 0x5a, 0xb2,
	0xe9, 0x8d, 0x35, 0x1e, 0x17, 0x2b, 0x1a, 0xf4, 0x27, 0x02, 0xdd, 0x4c, 0xcd, 0xe7, 0x8b, 0x52,
	0xda, 0x67, 0x71, 0x77, 0xd0, 0x3a, 0xed, 0x9d, 0x7d, 0xbb, 0xb5, 0xa4, 0xd6, 0x3a, 0x26, 0x6d,
	0xc2, 0x1e, 0x7e, 0x45, 0x20, 0x0e, 0x16, 0x8b, 0xf2, 0xb3, 0x52, 0x7d, 0x5e, 0xa2, 0x52, 0xbe,
	0xb3, 0xee, 0xc0, 0xe1, 0x4a, 0xb6, 0xa6, 0x59, 0xa2, 0x34, 0x5a, 0xa1, 0xd8, 0x2f, 0x77, 0xa1,
	0xdf, 0x98, 0x5d, 0x19, 0x0a, 0x47, 0x2b, 0xdc, 0xcf, 0x05, 0xfa, 0x1a, 0x34, 0x10, 0xbb, 0xe4,
	0xb3, 0x85, 0x08, 0x8d, 0xd4, 0x7c, 0xe8, 0x13, 0x87, 0x0e, 0x5f, 0xec, 0xae, 0x66, 0xb4, 0xeb,
	0x2a, 0x6c, 0xdc, 0xdb, 0x10, 0xf9, 0x2e, 0x45, 0x69, 0x64, 0x1e, 0x5a, 0xf7, 0xa0, 0x01, 0x1f,
	0xe6, 0x6e, 0x1a, 0x64, 0xb3, 0x85, 0xb1, 0x42, 0xc7, 0x3b, 0x83, 0x96, 0x6b, 0xd1, 0xb0, 0xa5,
	0xbf, 0x10, 0x88, 0x9c, 0xaf, 0x46, 0xb3, 0xd6, 0x35, 0xd2, 0xec, 0x40, 0x2c, 0xed, 0xdb, 0x75,
	0xe4, 0x94, 0xc2, 0x2e, 0x97, 0x85, 0x9f, 0x34, 0xbb, 0x29, 0xae, 0xd7, 0x8f, 0xbb, 0xbd, 0xf5,
	0xe3, 0x8e, 0xfe, 0x49, 0x80, 0xae, 0x8b, 0x1b, 0xda, 0xb0, 0x8d, 0xec, 0xbf, 0xdf, 0x5a, 0xf6,
	0xff, 0xaa, 0xc7, 0xf4, 0x38, 0x40, 0x4d, 0x13, 0x0e, 0x1f, 0xc0, 0x49, 0x30, 0xcf, 0xb9, 0x1b,
	0xf0, 0x65, 0xce, 0xac, 0x9c, 0x0b, 0x37, 0xbe, 0xdc, 0xff, 0x50, 0x63, 0xbe, 0x50, 0xba, 0x0e,
	0xf1, 0xe5, 0xf5, 0x65, 0x17, 0xfa, 0x9b, 0x77, 0x14, 0xfa, 0x33, 0x81, 0x5e, 0x38, 0xda, 0x4a,
	0x3e, 0x17, 0xf1, 0x19, 0xce, 0xa9, 0x6f, 0xb6, 0x36, 0x41, 0xcd, 0x85, 0x22, 0x05, 0xbf, 0xfe,
	0x90, 0xcf, 0x05, 0xfd, 0x95, 0xc0, 0x51, 0x73, 0xdc, 0xb1, 0x99, 0x34, 0x36, 0xbe, 0xbf, 0xe5,
	0x54, 0x9a, 0x78, 0xd3, 0x08, 0xd7, 0x4e, 0xe2, 0x47, 0xd2, 0x58, 0xfa, 0x1b, 0x81, 0xff, 0xd5,
	0x13, 0xc0, 0x6b, 0xee, 0x19, 0x8d, 0xb6, 0xfc, 0x8c, 0xad, 0x63, 0x4e, 0x8f, 0xc5, 0xd2, 0x36,
	0x05, 0x8b, 0x9c, 0x7e, 0x27, 0x40, 0x67, 0xdc, 0x58, 0x36, 0x57, 0xb9, 0x9c, 0x48, 0x91, 0x63,
	0xf1, 0xc6, 0x0f, 0x90, 0xd2, 0x8b, 0xad, 0xa5, 0x74, 0xa5, 0xc3, 0xd2, 0xbe, 0x63, 0xf0, 0x41,
	0x20, 0xf0, 0x0e, 0xb7, 0x82, 0xfe, 0x41, 0xe0, 0x04, 0x69, 0x2d, 0x2a, 0x34, 0xd5, 0x22, 0x93,
	0xe2, 0x52, 0xe4, 0xf1, 0xf9, 0xb5, 0x22, 0x86, 0xd2, 0x7c, 0x8c, 0x14, 0xd2, 0xc0, 0xc0, 0x5d,
	0x77, 0x2e, 0xf9, 0x4c, 0xe6, 0xbe, 0x50, 0xe3, 0x37, 0xfc, 0x75, 0x07, 0xa1, 0x14, 0x7f, 0x99,
	0xdc, 0x81, 0x43, 0x59, 0x06, 0x13, 0xc1, 0x8d, 0x2a, 0xe3, 0x37, 0x71, 0xe6, 0x44, 0x01, 0x4d,
	0x11, 0xa4, 0x2f, 0x41, 0x77, 0x2c, 0x8c, 0xf5, 0x57, 0xc8, 0xb7, 0xd0, 0xcb, 0xbe, 0x03, 0xdc,
	0x1d, 0x72, 0xdc, 0xc6, 0x5f, 0x74, 0xf7, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x07, 0xf4,
	0x63, 0xeb, 0x0d, 0x00, 0x00,
}
