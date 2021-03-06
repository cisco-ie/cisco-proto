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
// source: pce_pfx_bag.proto

package cisco_ios_xr_infra_xtc_oper_pce_topology_prefix_infos_prefix_info

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

type PcePfxBag_KEYS struct {
	NodeIdentifier       uint32   `protobuf:"varint,1,opt,name=node_identifier,json=nodeIdentifier,proto3" json:"node_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcePfxBag_KEYS) Reset()         { *m = PcePfxBag_KEYS{} }
func (m *PcePfxBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PcePfxBag_KEYS) ProtoMessage()    {}
func (*PcePfxBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24ed4f5c420edb1, []int{0}
}

func (m *PcePfxBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePfxBag_KEYS.Unmarshal(m, b)
}
func (m *PcePfxBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePfxBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PcePfxBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePfxBag_KEYS.Merge(m, src)
}
func (m *PcePfxBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PcePfxBag_KEYS.Size(m)
}
func (m *PcePfxBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePfxBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PcePfxBag_KEYS proto.InternalMessageInfo

func (m *PcePfxBag_KEYS) GetNodeIdentifier() uint32 {
	if m != nil {
		return m.NodeIdentifier
	}
	return 0
}

type PceIgpInfoIsis struct {
	SystemId             string   `protobuf:"bytes,1,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	Level                uint32   `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceIgpInfoIsis) Reset()         { *m = PceIgpInfoIsis{} }
func (m *PceIgpInfoIsis) String() string { return proto.CompactTextString(m) }
func (*PceIgpInfoIsis) ProtoMessage()    {}
func (*PceIgpInfoIsis) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24ed4f5c420edb1, []int{1}
}

func (m *PceIgpInfoIsis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIgpInfoIsis.Unmarshal(m, b)
}
func (m *PceIgpInfoIsis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIgpInfoIsis.Marshal(b, m, deterministic)
}
func (m *PceIgpInfoIsis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIgpInfoIsis.Merge(m, src)
}
func (m *PceIgpInfoIsis) XXX_Size() int {
	return xxx_messageInfo_PceIgpInfoIsis.Size(m)
}
func (m *PceIgpInfoIsis) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIgpInfoIsis.DiscardUnknown(m)
}

var xxx_messageInfo_PceIgpInfoIsis proto.InternalMessageInfo

func (m *PceIgpInfoIsis) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *PceIgpInfoIsis) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type PceIgpInfoOspf struct {
	RouterId             string   `protobuf:"bytes,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	Area                 uint32   `protobuf:"varint,2,opt,name=area,proto3" json:"area,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceIgpInfoOspf) Reset()         { *m = PceIgpInfoOspf{} }
func (m *PceIgpInfoOspf) String() string { return proto.CompactTextString(m) }
func (*PceIgpInfoOspf) ProtoMessage()    {}
func (*PceIgpInfoOspf) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24ed4f5c420edb1, []int{2}
}

func (m *PceIgpInfoOspf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIgpInfoOspf.Unmarshal(m, b)
}
func (m *PceIgpInfoOspf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIgpInfoOspf.Marshal(b, m, deterministic)
}
func (m *PceIgpInfoOspf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIgpInfoOspf.Merge(m, src)
}
func (m *PceIgpInfoOspf) XXX_Size() int {
	return xxx_messageInfo_PceIgpInfoOspf.Size(m)
}
func (m *PceIgpInfoOspf) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIgpInfoOspf.DiscardUnknown(m)
}

var xxx_messageInfo_PceIgpInfoOspf proto.InternalMessageInfo

func (m *PceIgpInfoOspf) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *PceIgpInfoOspf) GetArea() uint32 {
	if m != nil {
		return m.Area
	}
	return 0
}

type PceIgpInfoBgp struct {
	RouterId             string   `protobuf:"bytes,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	ConfedAsn            uint32   `protobuf:"varint,2,opt,name=confed_asn,json=confedAsn,proto3" json:"confed_asn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceIgpInfoBgp) Reset()         { *m = PceIgpInfoBgp{} }
func (m *PceIgpInfoBgp) String() string { return proto.CompactTextString(m) }
func (*PceIgpInfoBgp) ProtoMessage()    {}
func (*PceIgpInfoBgp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24ed4f5c420edb1, []int{3}
}

func (m *PceIgpInfoBgp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIgpInfoBgp.Unmarshal(m, b)
}
func (m *PceIgpInfoBgp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIgpInfoBgp.Marshal(b, m, deterministic)
}
func (m *PceIgpInfoBgp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIgpInfoBgp.Merge(m, src)
}
func (m *PceIgpInfoBgp) XXX_Size() int {
	return xxx_messageInfo_PceIgpInfoBgp.Size(m)
}
func (m *PceIgpInfoBgp) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIgpInfoBgp.DiscardUnknown(m)
}

var xxx_messageInfo_PceIgpInfoBgp proto.InternalMessageInfo

func (m *PceIgpInfoBgp) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *PceIgpInfoBgp) GetConfedAsn() uint32 {
	if m != nil {
		return m.ConfedAsn
	}
	return 0
}

type PceIgpInfoType struct {
	IgpId                string          `protobuf:"bytes,1,opt,name=igp_id,json=igpId,proto3" json:"igp_id,omitempty"`
	Isis                 *PceIgpInfoIsis `protobuf:"bytes,2,opt,name=isis,proto3" json:"isis,omitempty"`
	Ospf                 *PceIgpInfoOspf `protobuf:"bytes,3,opt,name=ospf,proto3" json:"ospf,omitempty"`
	Bgp                  *PceIgpInfoBgp  `protobuf:"bytes,4,opt,name=bgp,proto3" json:"bgp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PceIgpInfoType) Reset()         { *m = PceIgpInfoType{} }
func (m *PceIgpInfoType) String() string { return proto.CompactTextString(m) }
func (*PceIgpInfoType) ProtoMessage()    {}
func (*PceIgpInfoType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24ed4f5c420edb1, []int{4}
}

func (m *PceIgpInfoType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIgpInfoType.Unmarshal(m, b)
}
func (m *PceIgpInfoType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIgpInfoType.Marshal(b, m, deterministic)
}
func (m *PceIgpInfoType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIgpInfoType.Merge(m, src)
}
func (m *PceIgpInfoType) XXX_Size() int {
	return xxx_messageInfo_PceIgpInfoType.Size(m)
}
func (m *PceIgpInfoType) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIgpInfoType.DiscardUnknown(m)
}

var xxx_messageInfo_PceIgpInfoType proto.InternalMessageInfo

func (m *PceIgpInfoType) GetIgpId() string {
	if m != nil {
		return m.IgpId
	}
	return ""
}

func (m *PceIgpInfoType) GetIsis() *PceIgpInfoIsis {
	if m != nil {
		return m.Isis
	}
	return nil
}

func (m *PceIgpInfoType) GetOspf() *PceIgpInfoOspf {
	if m != nil {
		return m.Ospf
	}
	return nil
}

func (m *PceIgpInfoType) GetBgp() *PceIgpInfoBgp {
	if m != nil {
		return m.Bgp
	}
	return nil
}

type PceIgpInfoBag struct {
	Igp                    *PceIgpInfoType `protobuf:"bytes,1,opt,name=igp,proto3" json:"igp,omitempty"`
	DomainIdentifier       uint64          `protobuf:"varint,2,opt,name=domain_identifier,json=domainIdentifier,proto3" json:"domain_identifier,omitempty"`
	AutonomousSystemNumber uint32          `protobuf:"varint,3,opt,name=autonomous_system_number,json=autonomousSystemNumber,proto3" json:"autonomous_system_number,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *PceIgpInfoBag) Reset()         { *m = PceIgpInfoBag{} }
func (m *PceIgpInfoBag) String() string { return proto.CompactTextString(m) }
func (*PceIgpInfoBag) ProtoMessage()    {}
func (*PceIgpInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24ed4f5c420edb1, []int{5}
}

func (m *PceIgpInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIgpInfoBag.Unmarshal(m, b)
}
func (m *PceIgpInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIgpInfoBag.Marshal(b, m, deterministic)
}
func (m *PceIgpInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIgpInfoBag.Merge(m, src)
}
func (m *PceIgpInfoBag) XXX_Size() int {
	return xxx_messageInfo_PceIgpInfoBag.Size(m)
}
func (m *PceIgpInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIgpInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceIgpInfoBag proto.InternalMessageInfo

func (m *PceIgpInfoBag) GetIgp() *PceIgpInfoType {
	if m != nil {
		return m.Igp
	}
	return nil
}

func (m *PceIgpInfoBag) GetDomainIdentifier() uint64 {
	if m != nil {
		return m.DomainIdentifier
	}
	return 0
}

func (m *PceIgpInfoBag) GetAutonomousSystemNumber() uint32 {
	if m != nil {
		return m.AutonomousSystemNumber
	}
	return 0
}

type PceSrgbInfoBag struct {
	Start                uint32          `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Size                 uint32          `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	IgpSrgb              *PceIgpInfoType `protobuf:"bytes,3,opt,name=igp_srgb,json=igpSrgb,proto3" json:"igp_srgb,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PceSrgbInfoBag) Reset()         { *m = PceSrgbInfoBag{} }
func (m *PceSrgbInfoBag) String() string { return proto.CompactTextString(m) }
func (*PceSrgbInfoBag) ProtoMessage()    {}
func (*PceSrgbInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24ed4f5c420edb1, []int{6}
}

func (m *PceSrgbInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceSrgbInfoBag.Unmarshal(m, b)
}
func (m *PceSrgbInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceSrgbInfoBag.Marshal(b, m, deterministic)
}
func (m *PceSrgbInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceSrgbInfoBag.Merge(m, src)
}
func (m *PceSrgbInfoBag) XXX_Size() int {
	return xxx_messageInfo_PceSrgbInfoBag.Size(m)
}
func (m *PceSrgbInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceSrgbInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceSrgbInfoBag proto.InternalMessageInfo

func (m *PceSrgbInfoBag) GetStart() uint32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *PceSrgbInfoBag) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PceSrgbInfoBag) GetIgpSrgb() *PceIgpInfoType {
	if m != nil {
		return m.IgpSrgb
	}
	return nil
}

type PceNodePidBag struct {
	NodeName             string            `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Ipv4BgpRouterIdSet   bool              `protobuf:"varint,2,opt,name=ipv4bgp_router_id_set,json=ipv4bgpRouterIdSet,proto3" json:"ipv4bgp_router_id_set,omitempty"`
	Ipv4BgpRouterId      string            `protobuf:"bytes,3,opt,name=ipv4bgp_router_id,json=ipv4bgpRouterId,proto3" json:"ipv4bgp_router_id,omitempty"`
	Ipv4TeRouterIdSet    bool              `protobuf:"varint,4,opt,name=ipv4te_router_id_set,json=ipv4teRouterIdSet,proto3" json:"ipv4te_router_id_set,omitempty"`
	Ipv4TeRouterId       string            `protobuf:"bytes,5,opt,name=ipv4te_router_id,json=ipv4teRouterId,proto3" json:"ipv4te_router_id,omitempty"`
	IgpInformation       []*PceIgpInfoBag  `protobuf:"bytes,6,rep,name=igp_information,json=igpInformation,proto3" json:"igp_information,omitempty"`
	SrgbInformation      []*PceSrgbInfoBag `protobuf:"bytes,7,rep,name=srgb_information,json=srgbInformation,proto3" json:"srgb_information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PceNodePidBag) Reset()         { *m = PceNodePidBag{} }
func (m *PceNodePidBag) String() string { return proto.CompactTextString(m) }
func (*PceNodePidBag) ProtoMessage()    {}
func (*PceNodePidBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24ed4f5c420edb1, []int{7}
}

func (m *PceNodePidBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceNodePidBag.Unmarshal(m, b)
}
func (m *PceNodePidBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceNodePidBag.Marshal(b, m, deterministic)
}
func (m *PceNodePidBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceNodePidBag.Merge(m, src)
}
func (m *PceNodePidBag) XXX_Size() int {
	return xxx_messageInfo_PceNodePidBag.Size(m)
}
func (m *PceNodePidBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceNodePidBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceNodePidBag proto.InternalMessageInfo

func (m *PceNodePidBag) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PceNodePidBag) GetIpv4BgpRouterIdSet() bool {
	if m != nil {
		return m.Ipv4BgpRouterIdSet
	}
	return false
}

func (m *PceNodePidBag) GetIpv4BgpRouterId() string {
	if m != nil {
		return m.Ipv4BgpRouterId
	}
	return ""
}

func (m *PceNodePidBag) GetIpv4TeRouterIdSet() bool {
	if m != nil {
		return m.Ipv4TeRouterIdSet
	}
	return false
}

func (m *PceNodePidBag) GetIpv4TeRouterId() string {
	if m != nil {
		return m.Ipv4TeRouterId
	}
	return ""
}

func (m *PceNodePidBag) GetIgpInformation() []*PceIgpInfoBag {
	if m != nil {
		return m.IgpInformation
	}
	return nil
}

func (m *PceNodePidBag) GetSrgbInformation() []*PceSrgbInfoBag {
	if m != nil {
		return m.SrgbInformation
	}
	return nil
}

type PceIpAddrType struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceIpAddrType) Reset()         { *m = PceIpAddrType{} }
func (m *PceIpAddrType) String() string { return proto.CompactTextString(m) }
func (*PceIpAddrType) ProtoMessage()    {}
func (*PceIpAddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24ed4f5c420edb1, []int{8}
}

func (m *PceIpAddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIpAddrType.Unmarshal(m, b)
}
func (m *PceIpAddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIpAddrType.Marshal(b, m, deterministic)
}
func (m *PceIpAddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIpAddrType.Merge(m, src)
}
func (m *PceIpAddrType) XXX_Size() int {
	return xxx_messageInfo_PceIpAddrType.Size(m)
}
func (m *PceIpAddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIpAddrType.DiscardUnknown(m)
}

var xxx_messageInfo_PceIpAddrType proto.InternalMessageInfo

func (m *PceIpAddrType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PceIpAddrType) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *PceIpAddrType) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type PcePfxBag struct {
	NodeIdentifierXr       uint32           `protobuf:"varint,50,opt,name=node_identifier_xr,json=nodeIdentifierXr,proto3" json:"node_identifier_xr,omitempty"`
	NodeProtocolIdentifier *PceNodePidBag   `protobuf:"bytes,51,opt,name=node_protocol_identifier,json=nodeProtocolIdentifier,proto3" json:"node_protocol_identifier,omitempty"`
	Address                []*PceIpAddrType `protobuf:"bytes,52,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}         `json:"-"`
	XXX_unrecognized       []byte           `json:"-"`
	XXX_sizecache          int32            `json:"-"`
}

func (m *PcePfxBag) Reset()         { *m = PcePfxBag{} }
func (m *PcePfxBag) String() string { return proto.CompactTextString(m) }
func (*PcePfxBag) ProtoMessage()    {}
func (*PcePfxBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24ed4f5c420edb1, []int{9}
}

func (m *PcePfxBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePfxBag.Unmarshal(m, b)
}
func (m *PcePfxBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePfxBag.Marshal(b, m, deterministic)
}
func (m *PcePfxBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePfxBag.Merge(m, src)
}
func (m *PcePfxBag) XXX_Size() int {
	return xxx_messageInfo_PcePfxBag.Size(m)
}
func (m *PcePfxBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePfxBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePfxBag proto.InternalMessageInfo

func (m *PcePfxBag) GetNodeIdentifierXr() uint32 {
	if m != nil {
		return m.NodeIdentifierXr
	}
	return 0
}

func (m *PcePfxBag) GetNodeProtocolIdentifier() *PceNodePidBag {
	if m != nil {
		return m.NodeProtocolIdentifier
	}
	return nil
}

func (m *PcePfxBag) GetAddress() []*PceIpAddrType {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*PcePfxBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.prefix_infos.prefix_info.pce_pfx_bag_KEYS")
	proto.RegisterType((*PceIgpInfoIsis)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.prefix_infos.prefix_info.pce_igp_info_isis")
	proto.RegisterType((*PceIgpInfoOspf)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.prefix_infos.prefix_info.pce_igp_info_ospf")
	proto.RegisterType((*PceIgpInfoBgp)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.prefix_infos.prefix_info.pce_igp_info_bgp")
	proto.RegisterType((*PceIgpInfoType)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.prefix_infos.prefix_info.pce_igp_info_type")
	proto.RegisterType((*PceIgpInfoBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.prefix_infos.prefix_info.pce_igp_info_bag")
	proto.RegisterType((*PceSrgbInfoBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.prefix_infos.prefix_info.pce_srgb_info_bag")
	proto.RegisterType((*PceNodePidBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.prefix_infos.prefix_info.pce_node_pid_bag")
	proto.RegisterType((*PceIpAddrType)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.prefix_infos.prefix_info.pce_ip_addr_type")
	proto.RegisterType((*PcePfxBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.prefix_infos.prefix_info.pce_pfx_bag")
}

func init() { proto.RegisterFile("pce_pfx_bag.proto", fileDescriptor_f24ed4f5c420edb1) }

var fileDescriptor_f24ed4f5c420edb1 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6a, 0x14, 0x31,
	0x14, 0x66, 0xbb, 0x7f, 0xdd, 0x53, 0xda, 0xdd, 0x0d, 0x6d, 0x1d, 0x28, 0x42, 0x99, 0x1b, 0x8b,
	0xca, 0x8a, 0x6d, 0x11, 0xc1, 0xab, 0x82, 0x0a, 0x8b, 0x50, 0x24, 0xeb, 0x85, 0x5e, 0x85, 0xcc,
	0x4e, 0x66, 0x0c, 0xec, 0x4c, 0x42, 0x92, 0x2d, 0x5b, 0x11, 0x7c, 0x02, 0x5f, 0xc5, 0x07, 0xf0,
	0x6d, 0xbc, 0x10, 0x7c, 0x0c, 0xc9, 0x99, 0xd9, 0xdd, 0xe9, 0x14, 0xbc, 0x71, 0xf7, 0x2e, 0x39,
	0x49, 0xbe, 0x2f, 0xe7, 0x7c, 0xdf, 0x49, 0x60, 0xa8, 0xa7, 0x82, 0xe9, 0x64, 0xc1, 0x22, 0x9e,
	0x8e, 0xb4, 0x51, 0x4e, 0x91, 0xab, 0xa9, 0xb4, 0x53, 0xc5, 0xa4, 0xb2, 0x6c, 0x61, 0x98, 0xcc,
	0x13, 0xc3, 0xd9, 0xc2, 0x4d, 0x99, 0xd2, 0xc2, 0x8c, 0xfc, 0x76, 0xa7, 0xb4, 0x9a, 0xa9, 0xf4,
	0x76, 0xa4, 0x8d, 0x48, 0xe4, 0xc2, 0x6f, 0x51, 0xb6, 0x3a, 0x09, 0x5f, 0xc1, 0xa0, 0x82, 0xcb,
	0xde, 0xbd, 0xf9, 0x34, 0x21, 0x8f, 0xa0, 0x9f, 0xab, 0x58, 0x30, 0x19, 0x8b, 0xdc, 0xc9, 0x44,
	0x0a, 0x13, 0x34, 0x4e, 0x1b, 0x67, 0xfb, 0xf4, 0xc0, 0x87, 0xc7, 0xab, 0x68, 0xf8, 0xb6, 0xb8,
	0x94, 0x4c, 0x35, 0x82, 0x31, 0x69, 0xa5, 0x25, 0x27, 0xd0, 0xb3, 0xb7, 0xd6, 0x89, 0x8c, 0xc9,
	0x18, 0xcf, 0xf5, 0xe8, 0x6e, 0x11, 0x18, 0xc7, 0xe4, 0x10, 0xda, 0x33, 0x71, 0x23, 0x66, 0xc1,
	0x0e, 0x02, 0x16, 0x93, 0xf0, 0x75, 0x0d, 0x47, 0x59, 0x9d, 0x78, 0x1c, 0xa3, 0xe6, 0x4e, 0x98,
	0x0a, 0x4e, 0x11, 0x18, 0xc7, 0x84, 0x40, 0x8b, 0x1b, 0xc1, 0x4b, 0x18, 0x1c, 0x87, 0xd7, 0x45,
	0x2a, 0x2b, 0x94, 0x28, 0xd5, 0xff, 0x06, 0x79, 0x08, 0x30, 0x55, 0x79, 0x22, 0x62, 0xc6, 0x6d,
	0x5e, 0x42, 0xf5, 0x8a, 0xc8, 0x95, 0xcd, 0xc3, 0xdf, 0x3b, 0xb5, 0x6b, 0xb9, 0x5b, 0x2d, 0xc8,
	0x11, 0x74, 0x30, 0xb0, 0x84, 0x6b, 0xcb, 0x54, 0x8f, 0x63, 0xf2, 0x19, 0x5a, 0x3e, 0x7b, 0x44,
	0xd9, 0x3b, 0xff, 0x30, 0xfa, 0x6f, 0x65, 0x46, 0xf7, 0x2a, 0x4b, 0x91, 0xc1, 0x33, 0xf9, 0xfa,
	0x04, 0xcd, 0xed, 0x30, 0x79, 0x6c, 0x8a, 0x0c, 0x44, 0x40, 0x33, 0x4a, 0x75, 0xd0, 0x42, 0xa2,
	0xc9, 0xa6, 0x89, 0xa2, 0x54, 0x53, 0x8f, 0x1f, 0xfe, 0x6a, 0xd4, 0x85, 0xe3, 0x29, 0x49, 0xa0,
	0x29, 0x53, 0x8d, 0x35, 0xde, 0x42, 0x92, 0x5e, 0x49, 0xea, 0x09, 0xc8, 0x13, 0x18, 0xc6, 0x2a,
	0xe3, 0x32, 0xaf, 0xba, 0xdd, 0x8b, 0xd8, 0xa2, 0x83, 0x62, 0x61, 0xed, 0x77, 0xf2, 0x12, 0x02,
	0x3e, 0x77, 0x2a, 0x57, 0x99, 0x9a, 0x5b, 0x56, 0xba, 0x3c, 0x9f, 0x67, 0x91, 0x30, 0x28, 0xc7,
	0x3e, 0x3d, 0x5e, 0xaf, 0x4f, 0x70, 0xf9, 0x1a, 0x57, 0xc3, 0x1f, 0x8d, 0xc2, 0x4b, 0xd6, 0xa4,
	0xd1, 0x3a, 0xc9, 0x43, 0x68, 0x5b, 0xc7, 0x8d, 0x2b, 0xdb, 0xab, 0x98, 0x78, 0x6f, 0x5b, 0xf9,
	0x45, 0x2c, 0xbd, 0xed, 0xc7, 0x44, 0xc1, 0xae, 0xbf, 0xbc, 0x3f, 0xbe, 0x2d, 0xe1, 0xb1, 0x26,
	0x5d, 0x99, 0xea, 0x89, 0x49, 0xa3, 0xf0, 0x4f, 0xb3, 0x10, 0x05, 0x1f, 0x02, 0x2d, 0x63, 0xbc,
	0xef, 0x09, 0xf4, 0x70, 0x9e, 0xf3, 0x4c, 0x2c, 0xbb, 0xc9, 0x07, 0xae, 0x79, 0x26, 0xc8, 0x73,
	0x38, 0x92, 0xfa, 0xe6, 0x32, 0x4a, 0x35, 0x5b, 0xb5, 0x1c, 0xb3, 0xc2, 0x61, 0x1e, 0xbb, 0x94,
	0x94, 0x8b, 0xb4, 0xec, 0xbe, 0x89, 0x70, 0xe4, 0x31, 0x0c, 0xef, 0x1d, 0xc1, 0xf4, 0x7a, 0xb4,
	0x5f, 0xdb, 0x4e, 0x9e, 0xc1, 0xa1, 0x0f, 0x39, 0x51, 0x43, 0x6f, 0x21, 0xfa, 0xb0, 0x58, 0xab,
	0x82, 0x9f, 0xc1, 0xa0, 0x7e, 0x20, 0x68, 0x23, 0xf6, 0xc1, 0xdd, 0xcd, 0xe4, 0x2b, 0xf4, 0x97,
	0x55, 0x30, 0x19, 0x77, 0x52, 0xe5, 0x41, 0xe7, 0xb4, 0xb9, 0x15, 0xcf, 0xf3, 0x94, 0x1e, 0xf8,
	0x07, 0x63, 0x4d, 0x45, 0xbe, 0xc1, 0x60, 0xe5, 0x8a, 0x25, 0x7d, 0x17, 0xe9, 0x37, 0x25, 0xf1,
	0x1d, 0xd3, 0xd1, 0xbe, 0x9f, 0x56, 0x2e, 0x10, 0x4e, 0xca, 0xf6, 0xd3, 0x8c, 0xc7, 0xb1, 0x29,
	0x5e, 0xb9, 0x07, 0xd0, 0xe5, 0x49, 0x55, 0xe7, 0x0e, 0x4f, 0x50, 0x65, 0x02, 0x2d, 0x5f, 0x3d,
	0x14, 0xb5, 0x47, 0x71, 0x5c, 0xc6, 0x5e, 0x94, 0xca, 0xe1, 0x38, 0xfc, 0xb9, 0x03, 0x7b, 0x95,
	0x8f, 0x85, 0x3c, 0x05, 0x52, 0xfb, 0x53, 0xd8, 0xc2, 0x04, 0xe7, 0x68, 0xf1, 0xc1, 0xdd, 0x6f,
	0xe5, 0xa3, 0x21, 0xdf, 0x1b, 0x10, 0x14, 0xce, 0xf3, 0xff, 0xdc, 0x54, 0xcd, 0xaa, 0xdd, 0x79,
	0xb1, 0xd1, 0xf7, 0xa8, 0x6a, 0x70, 0x7a, 0xec, 0x67, 0xef, 0x4b, 0xce, 0x4a, 0xe3, 0x67, 0xd0,
	0xf5, 0xb5, 0x11, 0xd6, 0x06, 0x97, 0x9b, 0x75, 0x46, 0xa5, 0xe8, 0x74, 0xc9, 0x11, 0x75, 0x30,
	0xed, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0x9b, 0x5f, 0x3f, 0xf3, 0x07, 0x00, 0x00,
}
