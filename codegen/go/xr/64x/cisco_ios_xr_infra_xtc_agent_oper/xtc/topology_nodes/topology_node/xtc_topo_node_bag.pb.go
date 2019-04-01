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
// source: xtc_topo_node_bag.proto

package cisco_ios_xr_infra_xtc_agent_oper_xtc_topology_nodes_topology_node

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

type XtcTopoNodeBag_KEYS struct {
	NodeIdentifier       uint32   `protobuf:"varint,1,opt,name=node_identifier,json=nodeIdentifier,proto3" json:"node_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcTopoNodeBag_KEYS) Reset()         { *m = XtcTopoNodeBag_KEYS{} }
func (m *XtcTopoNodeBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*XtcTopoNodeBag_KEYS) ProtoMessage()    {}
func (*XtcTopoNodeBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{0}
}

func (m *XtcTopoNodeBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcTopoNodeBag_KEYS.Unmarshal(m, b)
}
func (m *XtcTopoNodeBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcTopoNodeBag_KEYS.Marshal(b, m, deterministic)
}
func (m *XtcTopoNodeBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcTopoNodeBag_KEYS.Merge(m, src)
}
func (m *XtcTopoNodeBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_XtcTopoNodeBag_KEYS.Size(m)
}
func (m *XtcTopoNodeBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcTopoNodeBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_XtcTopoNodeBag_KEYS proto.InternalMessageInfo

func (m *XtcTopoNodeBag_KEYS) GetNodeIdentifier() uint32 {
	if m != nil {
		return m.NodeIdentifier
	}
	return 0
}

type XtcIgpInfoIsis struct {
	SystemId             string   `protobuf:"bytes,1,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	Level                uint32   `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcIgpInfoIsis) Reset()         { *m = XtcIgpInfoIsis{} }
func (m *XtcIgpInfoIsis) String() string { return proto.CompactTextString(m) }
func (*XtcIgpInfoIsis) ProtoMessage()    {}
func (*XtcIgpInfoIsis) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{1}
}

func (m *XtcIgpInfoIsis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcIgpInfoIsis.Unmarshal(m, b)
}
func (m *XtcIgpInfoIsis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcIgpInfoIsis.Marshal(b, m, deterministic)
}
func (m *XtcIgpInfoIsis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcIgpInfoIsis.Merge(m, src)
}
func (m *XtcIgpInfoIsis) XXX_Size() int {
	return xxx_messageInfo_XtcIgpInfoIsis.Size(m)
}
func (m *XtcIgpInfoIsis) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcIgpInfoIsis.DiscardUnknown(m)
}

var xxx_messageInfo_XtcIgpInfoIsis proto.InternalMessageInfo

func (m *XtcIgpInfoIsis) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *XtcIgpInfoIsis) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type XtcIgpInfoOspf struct {
	RouterId             string   `protobuf:"bytes,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	Area                 uint32   `protobuf:"varint,2,opt,name=area,proto3" json:"area,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcIgpInfoOspf) Reset()         { *m = XtcIgpInfoOspf{} }
func (m *XtcIgpInfoOspf) String() string { return proto.CompactTextString(m) }
func (*XtcIgpInfoOspf) ProtoMessage()    {}
func (*XtcIgpInfoOspf) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{2}
}

func (m *XtcIgpInfoOspf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcIgpInfoOspf.Unmarshal(m, b)
}
func (m *XtcIgpInfoOspf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcIgpInfoOspf.Marshal(b, m, deterministic)
}
func (m *XtcIgpInfoOspf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcIgpInfoOspf.Merge(m, src)
}
func (m *XtcIgpInfoOspf) XXX_Size() int {
	return xxx_messageInfo_XtcIgpInfoOspf.Size(m)
}
func (m *XtcIgpInfoOspf) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcIgpInfoOspf.DiscardUnknown(m)
}

var xxx_messageInfo_XtcIgpInfoOspf proto.InternalMessageInfo

func (m *XtcIgpInfoOspf) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *XtcIgpInfoOspf) GetArea() uint32 {
	if m != nil {
		return m.Area
	}
	return 0
}

type XtcIgpInfoBgp struct {
	RouterId             string   `protobuf:"bytes,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcIgpInfoBgp) Reset()         { *m = XtcIgpInfoBgp{} }
func (m *XtcIgpInfoBgp) String() string { return proto.CompactTextString(m) }
func (*XtcIgpInfoBgp) ProtoMessage()    {}
func (*XtcIgpInfoBgp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{3}
}

func (m *XtcIgpInfoBgp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcIgpInfoBgp.Unmarshal(m, b)
}
func (m *XtcIgpInfoBgp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcIgpInfoBgp.Marshal(b, m, deterministic)
}
func (m *XtcIgpInfoBgp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcIgpInfoBgp.Merge(m, src)
}
func (m *XtcIgpInfoBgp) XXX_Size() int {
	return xxx_messageInfo_XtcIgpInfoBgp.Size(m)
}
func (m *XtcIgpInfoBgp) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcIgpInfoBgp.DiscardUnknown(m)
}

var xxx_messageInfo_XtcIgpInfoBgp proto.InternalMessageInfo

func (m *XtcIgpInfoBgp) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

type XtcIgpInfoType struct {
	IgpId                string          `protobuf:"bytes,1,opt,name=igp_id,json=igpId,proto3" json:"igp_id,omitempty"`
	Isis                 *XtcIgpInfoIsis `protobuf:"bytes,2,opt,name=isis,proto3" json:"isis,omitempty"`
	Ospf                 *XtcIgpInfoOspf `protobuf:"bytes,3,opt,name=ospf,proto3" json:"ospf,omitempty"`
	Bgp                  *XtcIgpInfoBgp  `protobuf:"bytes,4,opt,name=bgp,proto3" json:"bgp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *XtcIgpInfoType) Reset()         { *m = XtcIgpInfoType{} }
func (m *XtcIgpInfoType) String() string { return proto.CompactTextString(m) }
func (*XtcIgpInfoType) ProtoMessage()    {}
func (*XtcIgpInfoType) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{4}
}

func (m *XtcIgpInfoType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcIgpInfoType.Unmarshal(m, b)
}
func (m *XtcIgpInfoType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcIgpInfoType.Marshal(b, m, deterministic)
}
func (m *XtcIgpInfoType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcIgpInfoType.Merge(m, src)
}
func (m *XtcIgpInfoType) XXX_Size() int {
	return xxx_messageInfo_XtcIgpInfoType.Size(m)
}
func (m *XtcIgpInfoType) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcIgpInfoType.DiscardUnknown(m)
}

var xxx_messageInfo_XtcIgpInfoType proto.InternalMessageInfo

func (m *XtcIgpInfoType) GetIgpId() string {
	if m != nil {
		return m.IgpId
	}
	return ""
}

func (m *XtcIgpInfoType) GetIsis() *XtcIgpInfoIsis {
	if m != nil {
		return m.Isis
	}
	return nil
}

func (m *XtcIgpInfoType) GetOspf() *XtcIgpInfoOspf {
	if m != nil {
		return m.Ospf
	}
	return nil
}

func (m *XtcIgpInfoType) GetBgp() *XtcIgpInfoBgp {
	if m != nil {
		return m.Bgp
	}
	return nil
}

type XtcIgpInfoBag struct {
	Igp                  *XtcIgpInfoType `protobuf:"bytes,1,opt,name=igp,proto3" json:"igp,omitempty"`
	DomainIdentifier     uint64          `protobuf:"varint,2,opt,name=domain_identifier,json=domainIdentifier,proto3" json:"domain_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *XtcIgpInfoBag) Reset()         { *m = XtcIgpInfoBag{} }
func (m *XtcIgpInfoBag) String() string { return proto.CompactTextString(m) }
func (*XtcIgpInfoBag) ProtoMessage()    {}
func (*XtcIgpInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{5}
}

func (m *XtcIgpInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcIgpInfoBag.Unmarshal(m, b)
}
func (m *XtcIgpInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcIgpInfoBag.Marshal(b, m, deterministic)
}
func (m *XtcIgpInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcIgpInfoBag.Merge(m, src)
}
func (m *XtcIgpInfoBag) XXX_Size() int {
	return xxx_messageInfo_XtcIgpInfoBag.Size(m)
}
func (m *XtcIgpInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcIgpInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcIgpInfoBag proto.InternalMessageInfo

func (m *XtcIgpInfoBag) GetIgp() *XtcIgpInfoType {
	if m != nil {
		return m.Igp
	}
	return nil
}

func (m *XtcIgpInfoBag) GetDomainIdentifier() uint64 {
	if m != nil {
		return m.DomainIdentifier
	}
	return 0
}

type XtcNodePidBag struct {
	NodeName             string           `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Ipv4BgpRouterIdSet   bool             `protobuf:"varint,2,opt,name=ipv4_bgp_router_id_set,json=ipv4BgpRouterIdSet,proto3" json:"ipv4_bgp_router_id_set,omitempty"`
	Ipv4BgpRouterId      string           `protobuf:"bytes,3,opt,name=ipv4_bgp_router_id,json=ipv4BgpRouterId,proto3" json:"ipv4_bgp_router_id,omitempty"`
	Ipv4TeRouterIdSet    bool             `protobuf:"varint,4,opt,name=ipv4te_router_id_set,json=ipv4teRouterIdSet,proto3" json:"ipv4te_router_id_set,omitempty"`
	Ipv4TeRouterId       string           `protobuf:"bytes,5,opt,name=ipv4te_router_id,json=ipv4teRouterId,proto3" json:"ipv4te_router_id,omitempty"`
	IgpInformation       []*XtcIgpInfoBag `protobuf:"bytes,6,rep,name=igp_information,json=igpInformation,proto3" json:"igp_information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *XtcNodePidBag) Reset()         { *m = XtcNodePidBag{} }
func (m *XtcNodePidBag) String() string { return proto.CompactTextString(m) }
func (*XtcNodePidBag) ProtoMessage()    {}
func (*XtcNodePidBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{6}
}

func (m *XtcNodePidBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcNodePidBag.Unmarshal(m, b)
}
func (m *XtcNodePidBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcNodePidBag.Marshal(b, m, deterministic)
}
func (m *XtcNodePidBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcNodePidBag.Merge(m, src)
}
func (m *XtcNodePidBag) XXX_Size() int {
	return xxx_messageInfo_XtcNodePidBag.Size(m)
}
func (m *XtcNodePidBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcNodePidBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcNodePidBag proto.InternalMessageInfo

func (m *XtcNodePidBag) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *XtcNodePidBag) GetIpv4BgpRouterIdSet() bool {
	if m != nil {
		return m.Ipv4BgpRouterIdSet
	}
	return false
}

func (m *XtcNodePidBag) GetIpv4BgpRouterId() string {
	if m != nil {
		return m.Ipv4BgpRouterId
	}
	return ""
}

func (m *XtcNodePidBag) GetIpv4TeRouterIdSet() bool {
	if m != nil {
		return m.Ipv4TeRouterIdSet
	}
	return false
}

func (m *XtcNodePidBag) GetIpv4TeRouterId() string {
	if m != nil {
		return m.Ipv4TeRouterId
	}
	return ""
}

func (m *XtcNodePidBag) GetIgpInformation() []*XtcIgpInfoBag {
	if m != nil {
		return m.IgpInformation
	}
	return nil
}

type XtcIpAddrType struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcIpAddrType) Reset()         { *m = XtcIpAddrType{} }
func (m *XtcIpAddrType) String() string { return proto.CompactTextString(m) }
func (*XtcIpAddrType) ProtoMessage()    {}
func (*XtcIpAddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{7}
}

func (m *XtcIpAddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcIpAddrType.Unmarshal(m, b)
}
func (m *XtcIpAddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcIpAddrType.Marshal(b, m, deterministic)
}
func (m *XtcIpAddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcIpAddrType.Merge(m, src)
}
func (m *XtcIpAddrType) XXX_Size() int {
	return xxx_messageInfo_XtcIpAddrType.Size(m)
}
func (m *XtcIpAddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcIpAddrType.DiscardUnknown(m)
}

var xxx_messageInfo_XtcIpAddrType proto.InternalMessageInfo

func (m *XtcIpAddrType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *XtcIpAddrType) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *XtcIpAddrType) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type XtcSidBag struct {
	SidType              string         `protobuf:"bytes,1,opt,name=sid_type,json=sidType,proto3" json:"sid_type,omitempty"`
	SidPrefix            *XtcIpAddrType `protobuf:"bytes,2,opt,name=sid_prefix,json=sidPrefix,proto3" json:"sid_prefix,omitempty"`
	Algorithm            uint32         `protobuf:"varint,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	MplsLabel            uint32         `protobuf:"varint,4,opt,name=mpls_label,json=mplsLabel,proto3" json:"mpls_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *XtcSidBag) Reset()         { *m = XtcSidBag{} }
func (m *XtcSidBag) String() string { return proto.CompactTextString(m) }
func (*XtcSidBag) ProtoMessage()    {}
func (*XtcSidBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{8}
}

func (m *XtcSidBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcSidBag.Unmarshal(m, b)
}
func (m *XtcSidBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcSidBag.Marshal(b, m, deterministic)
}
func (m *XtcSidBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcSidBag.Merge(m, src)
}
func (m *XtcSidBag) XXX_Size() int {
	return xxx_messageInfo_XtcSidBag.Size(m)
}
func (m *XtcSidBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcSidBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcSidBag proto.InternalMessageInfo

func (m *XtcSidBag) GetSidType() string {
	if m != nil {
		return m.SidType
	}
	return ""
}

func (m *XtcSidBag) GetSidPrefix() *XtcIpAddrType {
	if m != nil {
		return m.SidPrefix
	}
	return nil
}

func (m *XtcSidBag) GetAlgorithm() uint32 {
	if m != nil {
		return m.Algorithm
	}
	return 0
}

func (m *XtcSidBag) GetMplsLabel() uint32 {
	if m != nil {
		return m.MplsLabel
	}
	return 0
}

type XtcIpv4LinkBag struct {
	LocalIpv4Address             string         `protobuf:"bytes,1,opt,name=local_ipv4_address,json=localIpv4Address,proto3" json:"local_ipv4_address,omitempty"`
	RemoteIpv4Address            string         `protobuf:"bytes,2,opt,name=remote_ipv4_address,json=remoteIpv4Address,proto3" json:"remote_ipv4_address,omitempty"`
	LocalIgpInformation          *XtcIgpInfoBag `protobuf:"bytes,3,opt,name=local_igp_information,json=localIgpInformation,proto3" json:"local_igp_information,omitempty"`
	RemoteNodeProtocolIdentifier *XtcNodePidBag `protobuf:"bytes,4,opt,name=remote_node_protocol_identifier,json=remoteNodeProtocolIdentifier,proto3" json:"remote_node_protocol_identifier,omitempty"`
	IgpMetric                    uint32         `protobuf:"varint,5,opt,name=igp_metric,json=igpMetric,proto3" json:"igp_metric,omitempty"`
	TeMetric                     uint32         `protobuf:"varint,6,opt,name=te_metric,json=teMetric,proto3" json:"te_metric,omitempty"`
	MaximumLinkBandwidth         uint64         `protobuf:"varint,7,opt,name=maximum_link_bandwidth,json=maximumLinkBandwidth,proto3" json:"maximum_link_bandwidth,omitempty"`
	MaxReservableBandwidth       uint64         `protobuf:"varint,8,opt,name=max_reservable_bandwidth,json=maxReservableBandwidth,proto3" json:"max_reservable_bandwidth,omitempty"`
	AdministrativeGroups         uint32         `protobuf:"varint,9,opt,name=administrative_groups,json=administrativeGroups,proto3" json:"administrative_groups,omitempty"`
	AdjacencySid                 []*XtcSidBag   `protobuf:"bytes,10,rep,name=adjacency_sid,json=adjacencySid,proto3" json:"adjacency_sid,omitempty"`
	Srlgs                        []uint32       `protobuf:"varint,11,rep,packed,name=srlgs,proto3" json:"srlgs,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}       `json:"-"`
	XXX_unrecognized             []byte         `json:"-"`
	XXX_sizecache                int32          `json:"-"`
}

func (m *XtcIpv4LinkBag) Reset()         { *m = XtcIpv4LinkBag{} }
func (m *XtcIpv4LinkBag) String() string { return proto.CompactTextString(m) }
func (*XtcIpv4LinkBag) ProtoMessage()    {}
func (*XtcIpv4LinkBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{9}
}

func (m *XtcIpv4LinkBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcIpv4LinkBag.Unmarshal(m, b)
}
func (m *XtcIpv4LinkBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcIpv4LinkBag.Marshal(b, m, deterministic)
}
func (m *XtcIpv4LinkBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcIpv4LinkBag.Merge(m, src)
}
func (m *XtcIpv4LinkBag) XXX_Size() int {
	return xxx_messageInfo_XtcIpv4LinkBag.Size(m)
}
func (m *XtcIpv4LinkBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcIpv4LinkBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcIpv4LinkBag proto.InternalMessageInfo

func (m *XtcIpv4LinkBag) GetLocalIpv4Address() string {
	if m != nil {
		return m.LocalIpv4Address
	}
	return ""
}

func (m *XtcIpv4LinkBag) GetRemoteIpv4Address() string {
	if m != nil {
		return m.RemoteIpv4Address
	}
	return ""
}

func (m *XtcIpv4LinkBag) GetLocalIgpInformation() *XtcIgpInfoBag {
	if m != nil {
		return m.LocalIgpInformation
	}
	return nil
}

func (m *XtcIpv4LinkBag) GetRemoteNodeProtocolIdentifier() *XtcNodePidBag {
	if m != nil {
		return m.RemoteNodeProtocolIdentifier
	}
	return nil
}

func (m *XtcIpv4LinkBag) GetIgpMetric() uint32 {
	if m != nil {
		return m.IgpMetric
	}
	return 0
}

func (m *XtcIpv4LinkBag) GetTeMetric() uint32 {
	if m != nil {
		return m.TeMetric
	}
	return 0
}

func (m *XtcIpv4LinkBag) GetMaximumLinkBandwidth() uint64 {
	if m != nil {
		return m.MaximumLinkBandwidth
	}
	return 0
}

func (m *XtcIpv4LinkBag) GetMaxReservableBandwidth() uint64 {
	if m != nil {
		return m.MaxReservableBandwidth
	}
	return 0
}

func (m *XtcIpv4LinkBag) GetAdministrativeGroups() uint32 {
	if m != nil {
		return m.AdministrativeGroups
	}
	return 0
}

func (m *XtcIpv4LinkBag) GetAdjacencySid() []*XtcSidBag {
	if m != nil {
		return m.AdjacencySid
	}
	return nil
}

func (m *XtcIpv4LinkBag) GetSrlgs() []uint32 {
	if m != nil {
		return m.Srlgs
	}
	return nil
}

type XtcIpv6LinkBag struct {
	LocalIpv6Address             string         `protobuf:"bytes,1,opt,name=local_ipv6_address,json=localIpv6Address,proto3" json:"local_ipv6_address,omitempty"`
	RemoteIpv6Address            string         `protobuf:"bytes,2,opt,name=remote_ipv6_address,json=remoteIpv6Address,proto3" json:"remote_ipv6_address,omitempty"`
	LocalIgpInformation          *XtcIgpInfoBag `protobuf:"bytes,3,opt,name=local_igp_information,json=localIgpInformation,proto3" json:"local_igp_information,omitempty"`
	RemoteNodeProtocolIdentifier *XtcNodePidBag `protobuf:"bytes,4,opt,name=remote_node_protocol_identifier,json=remoteNodeProtocolIdentifier,proto3" json:"remote_node_protocol_identifier,omitempty"`
	IgpMetric                    uint32         `protobuf:"varint,5,opt,name=igp_metric,json=igpMetric,proto3" json:"igp_metric,omitempty"`
	TeMetric                     uint32         `protobuf:"varint,6,opt,name=te_metric,json=teMetric,proto3" json:"te_metric,omitempty"`
	MaximumLinkBandwidth         uint64         `protobuf:"varint,7,opt,name=maximum_link_bandwidth,json=maximumLinkBandwidth,proto3" json:"maximum_link_bandwidth,omitempty"`
	MaxReservableBandwidth       uint64         `protobuf:"varint,8,opt,name=max_reservable_bandwidth,json=maxReservableBandwidth,proto3" json:"max_reservable_bandwidth,omitempty"`
	AdjacencySid                 []*XtcSidBag   `protobuf:"bytes,9,rep,name=adjacency_sid,json=adjacencySid,proto3" json:"adjacency_sid,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}       `json:"-"`
	XXX_unrecognized             []byte         `json:"-"`
	XXX_sizecache                int32          `json:"-"`
}

func (m *XtcIpv6LinkBag) Reset()         { *m = XtcIpv6LinkBag{} }
func (m *XtcIpv6LinkBag) String() string { return proto.CompactTextString(m) }
func (*XtcIpv6LinkBag) ProtoMessage()    {}
func (*XtcIpv6LinkBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{10}
}

func (m *XtcIpv6LinkBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcIpv6LinkBag.Unmarshal(m, b)
}
func (m *XtcIpv6LinkBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcIpv6LinkBag.Marshal(b, m, deterministic)
}
func (m *XtcIpv6LinkBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcIpv6LinkBag.Merge(m, src)
}
func (m *XtcIpv6LinkBag) XXX_Size() int {
	return xxx_messageInfo_XtcIpv6LinkBag.Size(m)
}
func (m *XtcIpv6LinkBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcIpv6LinkBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcIpv6LinkBag proto.InternalMessageInfo

func (m *XtcIpv6LinkBag) GetLocalIpv6Address() string {
	if m != nil {
		return m.LocalIpv6Address
	}
	return ""
}

func (m *XtcIpv6LinkBag) GetRemoteIpv6Address() string {
	if m != nil {
		return m.RemoteIpv6Address
	}
	return ""
}

func (m *XtcIpv6LinkBag) GetLocalIgpInformation() *XtcIgpInfoBag {
	if m != nil {
		return m.LocalIgpInformation
	}
	return nil
}

func (m *XtcIpv6LinkBag) GetRemoteNodeProtocolIdentifier() *XtcNodePidBag {
	if m != nil {
		return m.RemoteNodeProtocolIdentifier
	}
	return nil
}

func (m *XtcIpv6LinkBag) GetIgpMetric() uint32 {
	if m != nil {
		return m.IgpMetric
	}
	return 0
}

func (m *XtcIpv6LinkBag) GetTeMetric() uint32 {
	if m != nil {
		return m.TeMetric
	}
	return 0
}

func (m *XtcIpv6LinkBag) GetMaximumLinkBandwidth() uint64 {
	if m != nil {
		return m.MaximumLinkBandwidth
	}
	return 0
}

func (m *XtcIpv6LinkBag) GetMaxReservableBandwidth() uint64 {
	if m != nil {
		return m.MaxReservableBandwidth
	}
	return 0
}

func (m *XtcIpv6LinkBag) GetAdjacencySid() []*XtcSidBag {
	if m != nil {
		return m.AdjacencySid
	}
	return nil
}

type XtcTopoNodeBag struct {
	NodeIdentifierXr       uint32            `protobuf:"varint,50,opt,name=node_identifier_xr,json=nodeIdentifierXr,proto3" json:"node_identifier_xr,omitempty"`
	NodeProtocolIdentifier *XtcNodePidBag    `protobuf:"bytes,51,opt,name=node_protocol_identifier,json=nodeProtocolIdentifier,proto3" json:"node_protocol_identifier,omitempty"`
	PrefixSid              []*XtcSidBag      `protobuf:"bytes,52,rep,name=prefix_sid,json=prefixSid,proto3" json:"prefix_sid,omitempty"`
	Ipv4Link               []*XtcIpv4LinkBag `protobuf:"bytes,53,rep,name=ipv4_link,json=ipv4Link,proto3" json:"ipv4_link,omitempty"`
	Ipv6Link               []*XtcIpv6LinkBag `protobuf:"bytes,54,rep,name=ipv6_link,json=ipv6Link,proto3" json:"ipv6_link,omitempty"`
	Overload               bool              `protobuf:"varint,55,opt,name=overload,proto3" json:"overload,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *XtcTopoNodeBag) Reset()         { *m = XtcTopoNodeBag{} }
func (m *XtcTopoNodeBag) String() string { return proto.CompactTextString(m) }
func (*XtcTopoNodeBag) ProtoMessage()    {}
func (*XtcTopoNodeBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6714c208070abd, []int{11}
}

func (m *XtcTopoNodeBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcTopoNodeBag.Unmarshal(m, b)
}
func (m *XtcTopoNodeBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcTopoNodeBag.Marshal(b, m, deterministic)
}
func (m *XtcTopoNodeBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcTopoNodeBag.Merge(m, src)
}
func (m *XtcTopoNodeBag) XXX_Size() int {
	return xxx_messageInfo_XtcTopoNodeBag.Size(m)
}
func (m *XtcTopoNodeBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcTopoNodeBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcTopoNodeBag proto.InternalMessageInfo

func (m *XtcTopoNodeBag) GetNodeIdentifierXr() uint32 {
	if m != nil {
		return m.NodeIdentifierXr
	}
	return 0
}

func (m *XtcTopoNodeBag) GetNodeProtocolIdentifier() *XtcNodePidBag {
	if m != nil {
		return m.NodeProtocolIdentifier
	}
	return nil
}

func (m *XtcTopoNodeBag) GetPrefixSid() []*XtcSidBag {
	if m != nil {
		return m.PrefixSid
	}
	return nil
}

func (m *XtcTopoNodeBag) GetIpv4Link() []*XtcIpv4LinkBag {
	if m != nil {
		return m.Ipv4Link
	}
	return nil
}

func (m *XtcTopoNodeBag) GetIpv6Link() []*XtcIpv6LinkBag {
	if m != nil {
		return m.Ipv6Link
	}
	return nil
}

func (m *XtcTopoNodeBag) GetOverload() bool {
	if m != nil {
		return m.Overload
	}
	return false
}

func init() {
	proto.RegisterType((*XtcTopoNodeBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_topo_node_bag_KEYS")
	proto.RegisterType((*XtcIgpInfoIsis)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_igp_info_isis")
	proto.RegisterType((*XtcIgpInfoOspf)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_igp_info_ospf")
	proto.RegisterType((*XtcIgpInfoBgp)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_igp_info_bgp")
	proto.RegisterType((*XtcIgpInfoType)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_igp_info_type")
	proto.RegisterType((*XtcIgpInfoBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_igp_info_bag")
	proto.RegisterType((*XtcNodePidBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_node_pid_bag")
	proto.RegisterType((*XtcIpAddrType)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_ip_addr_type")
	proto.RegisterType((*XtcSidBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_sid_bag")
	proto.RegisterType((*XtcIpv4LinkBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_ipv4_link_bag")
	proto.RegisterType((*XtcIpv6LinkBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_ipv6_link_bag")
	proto.RegisterType((*XtcTopoNodeBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_nodes.topology_node.xtc_topo_node_bag")
}

func init() { proto.RegisterFile("xtc_topo_node_bag.proto", fileDescriptor_ab6714c208070abd) }

var fileDescriptor_ab6714c208070abd = []byte{
	// 984 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x96, 0x63, 0xc7, 0xf1, 0x54, 0x48, 0xe2, 0xf4, 0x26, 0xd9, 0x81, 0x5d, 0x44, 0x34, 0x17,
	0x56, 0x5a, 0xe4, 0x95, 0x92, 0x60, 0xb8, 0x6e, 0xc4, 0x8f, 0x2c, 0x96, 0x65, 0xd5, 0x59, 0x24,
	0x38, 0x8d, 0xda, 0x9e, 0xf6, 0x6c, 0xb3, 0x33, 0xd3, 0xad, 0xee, 0x8e, 0x71, 0x0e, 0x9c, 0xe1,
	0xc4, 0x0b, 0xf0, 0x02, 0x3c, 0x04, 0x67, 0x1e, 0x82, 0x27, 0x40, 0x3c, 0x05, 0xaa, 0x9a, 0xb1,
	0x33, 0x76, 0x12, 0x2e, 0xd8, 0x1c, 0xd0, 0xde, 0xa6, 0xab, 0xba, 0xbf, 0xaf, 0xeb, 0xe7, 0xab,
	0x69, 0xb8, 0x3f, 0xf5, 0xa3, 0xd8, 0x6b, 0xa3, 0xe3, 0x42, 0x27, 0x32, 0x1e, 0x8a, 0xb4, 0x67,
	0xac, 0xf6, 0x9a, 0x9d, 0x8f, 0x94, 0x1b, 0xe9, 0x58, 0x69, 0x17, 0x4f, 0x6d, 0xac, 0x8a, 0xb1,
	0x15, 0x31, 0xee, 0x15, 0xa9, 0x2c, 0x7c, 0xac, 0x8d, 0xb4, 0xbd, 0xa9, 0x1f, 0xf5, 0xf0, 0x68,
	0xa6, 0xd3, 0x2b, 0x3a, 0xee, 0x16, 0x97, 0xd1, 0x53, 0x38, 0xba, 0x01, 0x1f, 0x7f, 0xf1, 0xe9,
	0xb7, 0x17, 0xec, 0x7d, 0xd8, 0x23, 0x83, 0x4a, 0x64, 0xe1, 0xd5, 0x58, 0x49, 0x1b, 0x36, 0x8e,
	0x1b, 0x8f, 0x76, 0xf8, 0x2e, 0x9a, 0x07, 0x73, 0x6b, 0xf4, 0x19, 0xec, 0x23, 0x84, 0x4a, 0x0d,
	0xde, 0x41, 0xc7, 0xca, 0x29, 0xc7, 0x1e, 0x40, 0xe0, 0xae, 0x9c, 0x97, 0x79, 0xac, 0x12, 0x3a,
	0x17, 0xf0, 0x4e, 0x69, 0x18, 0x24, 0xec, 0x00, 0x36, 0x33, 0x39, 0x91, 0x59, 0xb8, 0x41, 0x80,
	0xe5, 0x22, 0xfa, 0x64, 0x09, 0x47, 0x3b, 0x33, 0x46, 0x1c, 0xab, 0x2f, 0xbd, 0xb4, 0x35, 0x9c,
	0xd2, 0x30, 0x48, 0x18, 0x83, 0x96, 0xb0, 0x52, 0x54, 0x30, 0xf4, 0x1d, 0x3d, 0x81, 0xee, 0x02,
	0xca, 0x30, 0x35, 0xff, 0x08, 0x12, 0xfd, 0xb9, 0xb1, 0xc4, 0xeb, 0xaf, 0x8c, 0x64, 0x87, 0xd0,
	0x26, 0xc3, 0x6c, 0xff, 0xa6, 0x4a, 0xcd, 0x20, 0x61, 0x0a, 0x5a, 0x18, 0x1e, 0x31, 0x6e, 0x9f,
	0x7c, 0xdd, 0xfb, 0xf7, 0x15, 0xe8, 0xdd, 0xc8, 0x1d, 0x27, 0x0a, 0xa4, 0xc2, 0x0c, 0x84, 0xcd,
	0x35, 0x51, 0x21, 0x38, 0x27, 0x0a, 0x36, 0x86, 0xe6, 0x30, 0x35, 0x61, 0x8b, 0x98, 0x5e, 0xae,
	0x9c, 0x69, 0x98, 0x1a, 0x8e, 0x04, 0xd1, 0xaf, 0x8d, 0xe5, 0xe2, 0x88, 0x94, 0xa5, 0xd0, 0x54,
	0xa9, 0xa1, 0x34, 0xaf, 0x23, 0x4c, 0xac, 0x26, 0x47, 0x06, 0xf6, 0x18, 0xf6, 0x13, 0x9d, 0x0b,
	0x55, 0xd4, 0x5b, 0x1a, 0x0b, 0xd9, 0xe2, 0xdd, 0xd2, 0x51, 0x6b, 0xea, 0xbf, 0x36, 0xca, 0xab,
	0x92, 0x04, 0x8c, 0x4a, 0xe8, 0xaa, 0x0f, 0x20, 0xa0, 0x75, 0x21, 0x72, 0x39, 0xeb, 0x23, 0x34,
	0x3c, 0x17, 0xb9, 0x64, 0x27, 0x70, 0xa4, 0xcc, 0xe4, 0x0c, 0xa3, 0x8d, 0xe7, 0xdd, 0x16, 0x3b,
	0xe9, 0x89, 0xa3, 0xc3, 0x19, 0x7a, 0xcf, 0x53, 0xc3, 0xab, 0xc6, 0xbb, 0x90, 0x9e, 0x3d, 0x06,
	0x76, 0xf3, 0x0c, 0x55, 0x3c, 0xe0, 0x7b, 0x4b, 0xfb, 0xd9, 0x13, 0x38, 0x40, 0x93, 0x97, 0x4b,
	0xf0, 0x2d, 0x82, 0xdf, 0x2f, 0x7d, 0x75, 0xf4, 0x47, 0xd0, 0x5d, 0x3e, 0x10, 0x6e, 0x12, 0xf6,
	0xee, 0xe2, 0x66, 0xf6, 0x03, 0xec, 0xcd, 0x12, 0x66, 0x73, 0xe1, 0x95, 0x2e, 0xc2, 0xf6, 0x71,
	0x73, 0x3d, 0xcd, 0x20, 0x52, 0xbe, 0x8b, 0x62, 0xba, 0xe6, 0x8a, 0x2e, 0xaa, 0xb6, 0x30, 0xb1,
	0x48, 0x12, 0x5b, 0x0a, 0xf0, 0x3e, 0x6c, 0x89, 0x71, 0x3d, 0xd3, 0x6d, 0x31, 0xa6, 0x3c, 0x33,
	0x68, 0xe1, 0xed, 0x29, 0xab, 0x01, 0xa7, 0xef, 0xca, 0xd6, 0xaf, 0x32, 0x47, 0xdf, 0xd1, 0x1f,
	0x0d, 0xd8, 0x46, 0x54, 0x57, 0x15, 0xef, 0x6d, 0xe8, 0xe0, 0x27, 0x82, 0x57, 0x88, 0x5b, 0x4e,
	0x25, 0x2f, 0x91, 0xcb, 0x01, 0xa0, 0xcb, 0x58, 0x39, 0x56, 0xd3, 0x4a, 0xdb, 0xab, 0x8b, 0xbc,
	0x16, 0x15, 0x0f, 0x9c, 0x4a, 0x5e, 0x10, 0x0d, 0x7b, 0x08, 0x81, 0xc8, 0x52, 0x6d, 0x95, 0x7f,
	0x95, 0xd3, 0xc5, 0x77, 0xf8, 0xb5, 0x81, 0xbd, 0x0b, 0x90, 0x9b, 0xcc, 0xc5, 0x99, 0x18, 0xca,
	0x8c, 0x4a, 0xbc, 0xc3, 0x03, 0xb4, 0x3c, 0x43, 0x43, 0xf4, 0x63, 0xbb, 0x1a, 0x5a, 0xd8, 0x3d,
	0x99, 0x2a, 0x5e, 0x53, 0x88, 0x1f, 0x00, 0xcb, 0xf4, 0x48, 0x64, 0xa5, 0x19, 0x59, 0xa5, 0x73,
	0x55, 0xb0, 0x5d, 0xf2, 0x0c, 0xcc, 0xe4, 0xec, 0x69, 0x69, 0x67, 0x3d, 0xb8, 0x67, 0x65, 0xae,
	0xbd, 0x5c, 0xdc, 0x5e, 0xe6, 0x75, 0xbf, 0x74, 0xd5, 0xf7, 0xff, 0xd4, 0x80, 0xc3, 0x0a, 0x7e,
	0xa9, 0x57, 0x9a, 0xeb, 0x1a, 0x1c, 0x22, 0xe5, 0xf7, 0xca, 0x7b, 0x2f, 0x34, 0x0c, 0xfb, 0xa5,
	0x01, 0xef, 0x55, 0x77, 0x2f, 0x05, 0x8a, 0xff, 0xc3, 0x91, 0xce, 0xea, 0xca, 0x5e, 0xf1, 0x34,
	0xab, 0x0f, 0x02, 0xfe, 0xb0, 0x24, 0x7f, 0xae, 0x13, 0xf9, 0xa2, 0xa2, 0xbe, 0x9e, 0x1d, 0x58,
	0x3b, 0x0c, 0x21, 0x97, 0xde, 0xaa, 0x11, 0x29, 0x6e, 0x87, 0x07, 0x2a, 0x35, 0x5f, 0x92, 0x01,
	0xa7, 0x88, 0x97, 0x33, 0x6f, 0x9b, 0xbc, 0x1d, 0x2f, 0x2b, 0xe7, 0x19, 0x1c, 0xe5, 0x62, 0xaa,
	0xf2, 0xcb, 0x7c, 0x56, 0xd6, 0x22, 0xf9, 0x5e, 0x25, 0xfe, 0x55, 0xb8, 0x45, 0x93, 0xea, 0xa0,
	0xf2, 0x3e, 0x53, 0xc5, 0xeb, 0xf3, 0x99, 0x8f, 0x7d, 0x0c, 0x61, 0x2e, 0xa6, 0xb1, 0x95, 0x4e,
	0xda, 0x89, 0x18, 0x66, 0xb2, 0x76, 0xae, 0x43, 0xe7, 0x10, 0x95, 0xcf, 0xdd, 0xd7, 0x27, 0x4f,
	0xe1, 0x50, 0x24, 0xb9, 0x2a, 0x94, 0xf3, 0x56, 0x78, 0x35, 0x91, 0x71, 0x6a, 0xf5, 0xa5, 0x71,
	0x61, 0x40, 0x17, 0x3b, 0x58, 0x74, 0x7e, 0x4e, 0x3e, 0xe6, 0x61, 0x47, 0x24, 0xdf, 0x89, 0x91,
	0x2c, 0x46, 0x57, 0xa8, 0xaf, 0x10, 0x68, 0x58, 0x7c, 0xb5, 0xaa, 0x5c, 0x57, 0x92, 0xe5, 0x6f,
	0xcd, 0x59, 0x2e, 0x14, 0xbd, 0x1a, 0x9c, 0xcd, 0x52, 0x17, 0x6e, 0x1f, 0x37, 0xf1, 0xd5, 0x40,
	0x8b, 0xe8, 0xb7, 0xcd, 0xb9, 0x12, 0xfa, 0xb7, 0x2b, 0xa1, 0x7f, 0x97, 0x12, 0xfa, 0xb7, 0x2a,
	0xa1, 0x7f, 0xa7, 0x12, 0xfa, 0x6f, 0x94, 0xf0, 0x3f, 0x56, 0xc2, 0x8d, 0xa6, 0x0e, 0xfe, 0x83,
	0xa6, 0x8e, 0x7e, 0x6f, 0x95, 0xed, 0xbb, 0xf0, 0x00, 0xc7, 0xf6, 0x5d, 0x7a, 0x7b, 0xc7, 0x53,
	0x1b, 0x9e, 0x50, 0x86, 0xba, 0x8b, 0xcf, 0xef, 0x6f, 0x2c, 0xfb, 0xb9, 0x01, 0xe1, 0x9d, 0xc5,
	0x3f, 0x5d, 0x63, 0xf1, 0x8f, 0x8a, 0xdb, 0xcb, 0x5e, 0x00, 0x94, 0xff, 0x52, 0xca, 0xe3, 0xd9,
	0x7a, 0xf2, 0x18, 0x94, 0x14, 0x38, 0x19, 0x2c, 0x04, 0xf3, 0x1f, 0x61, 0xf8, 0x21, 0xd1, 0xad,
	0xee, 0x21, 0x59, 0xff, 0xc3, 0xf2, 0x0e, 0x2e, 0xb1, 0xe3, 0x2a, 0xce, 0x72, 0xe4, 0x84, 0xfd,
	0x95, 0x73, 0xf6, 0x17, 0x39, 0xfb, 0xc4, 0xf9, 0x0e, 0x74, 0xf4, 0x44, 0xda, 0x4c, 0x8b, 0x24,
	0xfc, 0x88, 0x5e, 0x7d, 0xf3, 0xf5, 0xb0, 0x4d, 0xc5, 0x3f, 0xfd, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x1f, 0x3e, 0x61, 0x19, 0x2e, 0x0e, 0x00, 0x00,
}
