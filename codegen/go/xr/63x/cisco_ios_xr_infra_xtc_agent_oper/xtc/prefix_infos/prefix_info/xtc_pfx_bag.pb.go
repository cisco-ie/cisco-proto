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
// source: xtc_pfx_bag.proto

package cisco_ios_xr_infra_xtc_agent_oper_xtc_prefix_infos_prefix_info

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

type XtcPfxBag_KEYS struct {
	NodeIdentifier       uint32   `protobuf:"varint,1,opt,name=node_identifier,json=nodeIdentifier,proto3" json:"node_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcPfxBag_KEYS) Reset()         { *m = XtcPfxBag_KEYS{} }
func (m *XtcPfxBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*XtcPfxBag_KEYS) ProtoMessage()    {}
func (*XtcPfxBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3971953fe4f71a8, []int{0}
}

func (m *XtcPfxBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcPfxBag_KEYS.Unmarshal(m, b)
}
func (m *XtcPfxBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcPfxBag_KEYS.Marshal(b, m, deterministic)
}
func (m *XtcPfxBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcPfxBag_KEYS.Merge(m, src)
}
func (m *XtcPfxBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_XtcPfxBag_KEYS.Size(m)
}
func (m *XtcPfxBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcPfxBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_XtcPfxBag_KEYS proto.InternalMessageInfo

func (m *XtcPfxBag_KEYS) GetNodeIdentifier() uint32 {
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
	return fileDescriptor_e3971953fe4f71a8, []int{1}
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
	return fileDescriptor_e3971953fe4f71a8, []int{2}
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
	return fileDescriptor_e3971953fe4f71a8, []int{3}
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
	return fileDescriptor_e3971953fe4f71a8, []int{4}
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
	return fileDescriptor_e3971953fe4f71a8, []int{5}
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
	return fileDescriptor_e3971953fe4f71a8, []int{6}
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
	return fileDescriptor_e3971953fe4f71a8, []int{7}
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

type XtcPfxEntry struct {
	Ip                   *XtcIpAddrType `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *XtcPfxEntry) Reset()         { *m = XtcPfxEntry{} }
func (m *XtcPfxEntry) String() string { return proto.CompactTextString(m) }
func (*XtcPfxEntry) ProtoMessage()    {}
func (*XtcPfxEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3971953fe4f71a8, []int{8}
}

func (m *XtcPfxEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcPfxEntry.Unmarshal(m, b)
}
func (m *XtcPfxEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcPfxEntry.Marshal(b, m, deterministic)
}
func (m *XtcPfxEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcPfxEntry.Merge(m, src)
}
func (m *XtcPfxEntry) XXX_Size() int {
	return xxx_messageInfo_XtcPfxEntry.Size(m)
}
func (m *XtcPfxEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcPfxEntry.DiscardUnknown(m)
}

var xxx_messageInfo_XtcPfxEntry proto.InternalMessageInfo

func (m *XtcPfxEntry) GetIp() *XtcIpAddrType {
	if m != nil {
		return m.Ip
	}
	return nil
}

type XtcPfxBag struct {
	NodeIdentifierXr       uint32         `protobuf:"varint,50,opt,name=node_identifier_xr,json=nodeIdentifierXr,proto3" json:"node_identifier_xr,omitempty"`
	NodeProtocolIdentifier *XtcNodePidBag `protobuf:"bytes,51,opt,name=node_protocol_identifier,json=nodeProtocolIdentifier,proto3" json:"node_protocol_identifier,omitempty"`
	Address                []*XtcPfxEntry `protobuf:"bytes,52,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}       `json:"-"`
	XXX_unrecognized       []byte         `json:"-"`
	XXX_sizecache          int32          `json:"-"`
}

func (m *XtcPfxBag) Reset()         { *m = XtcPfxBag{} }
func (m *XtcPfxBag) String() string { return proto.CompactTextString(m) }
func (*XtcPfxBag) ProtoMessage()    {}
func (*XtcPfxBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3971953fe4f71a8, []int{9}
}

func (m *XtcPfxBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcPfxBag.Unmarshal(m, b)
}
func (m *XtcPfxBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcPfxBag.Marshal(b, m, deterministic)
}
func (m *XtcPfxBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcPfxBag.Merge(m, src)
}
func (m *XtcPfxBag) XXX_Size() int {
	return xxx_messageInfo_XtcPfxBag.Size(m)
}
func (m *XtcPfxBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcPfxBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcPfxBag proto.InternalMessageInfo

func (m *XtcPfxBag) GetNodeIdentifierXr() uint32 {
	if m != nil {
		return m.NodeIdentifierXr
	}
	return 0
}

func (m *XtcPfxBag) GetNodeProtocolIdentifier() *XtcNodePidBag {
	if m != nil {
		return m.NodeProtocolIdentifier
	}
	return nil
}

func (m *XtcPfxBag) GetAddress() []*XtcPfxEntry {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*XtcPfxBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.prefix_infos.prefix_info.xtc_pfx_bag_KEYS")
	proto.RegisterType((*XtcIgpInfoIsis)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.prefix_infos.prefix_info.xtc_igp_info_isis")
	proto.RegisterType((*XtcIgpInfoOspf)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.prefix_infos.prefix_info.xtc_igp_info_ospf")
	proto.RegisterType((*XtcIgpInfoBgp)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.prefix_infos.prefix_info.xtc_igp_info_bgp")
	proto.RegisterType((*XtcIgpInfoType)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.prefix_infos.prefix_info.xtc_igp_info_type")
	proto.RegisterType((*XtcIgpInfoBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.prefix_infos.prefix_info.xtc_igp_info_bag")
	proto.RegisterType((*XtcNodePidBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.prefix_infos.prefix_info.xtc_node_pid_bag")
	proto.RegisterType((*XtcIpAddrType)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.prefix_infos.prefix_info.xtc_ip_addr_type")
	proto.RegisterType((*XtcPfxEntry)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.prefix_infos.prefix_info.xtc_pfx_entry")
	proto.RegisterType((*XtcPfxBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.prefix_infos.prefix_info.xtc_pfx_bag")
}

func init() { proto.RegisterFile("xtc_pfx_bag.proto", fileDescriptor_e3971953fe4f71a8) }

var fileDescriptor_e3971953fe4f71a8 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xc1, 0x6e, 0xd4, 0x3c,
	0x10, 0xc7, 0xb5, 0xd9, 0x74, 0xdb, 0x9d, 0xaa, 0xed, 0xd6, 0xea, 0xd7, 0x2f, 0x52, 0x2f, 0x55,
	0x2e, 0x54, 0x2a, 0xda, 0x4a, 0x6d, 0xc5, 0x05, 0x89, 0x03, 0x02, 0xa4, 0x15, 0x02, 0x15, 0xf7,
	0x02, 0x27, 0xe3, 0x24, 0x8e, 0x65, 0x69, 0x37, 0x36, 0x8e, 0xa9, 0xb2, 0xaf, 0xc0, 0x73, 0xf0,
	0x02, 0xbc, 0x0f, 0x67, 0x9e, 0x03, 0x79, 0x92, 0x6c, 0xd3, 0x54, 0xe2, 0x42, 0xb8, 0xc5, 0x63,
	0xfb, 0xf7, 0x9f, 0x19, 0xff, 0x33, 0x70, 0x58, 0xb9, 0x94, 0x99, 0xbc, 0x62, 0x09, 0x97, 0x73,
	0x63, 0xb5, 0xd3, 0xe4, 0x45, 0xaa, 0xca, 0x54, 0x33, 0xa5, 0x4b, 0x56, 0x59, 0xa6, 0x8a, 0xdc,
	0x72, 0xe6, 0x4f, 0x71, 0x29, 0x0a, 0xc7, 0xb4, 0x11, 0x76, 0x5e, 0xb9, 0x74, 0x6e, 0xac, 0xc8,
	0x55, 0xe5, 0xf7, 0x75, 0xd9, 0x5d, 0xc4, 0xcf, 0x61, 0xd6, 0x81, 0xb2, 0xb7, 0xaf, 0x3f, 0xdd,
	0x92, 0x27, 0x70, 0x50, 0xe8, 0x4c, 0x30, 0x95, 0x89, 0xc2, 0xa9, 0x5c, 0x09, 0x1b, 0x8d, 0x4e,
	0x47, 0x67, 0x7b, 0x74, 0xdf, 0x87, 0x17, 0x9b, 0x68, 0xfc, 0xa6, 0xce, 0x48, 0x49, 0x83, 0x30,
	0xa6, 0x4a, 0x55, 0x92, 0x13, 0x98, 0x96, 0xeb, 0xd2, 0x89, 0x15, 0x53, 0x19, 0xde, 0x9b, 0xd2,
	0x9d, 0x3a, 0xb0, 0xc8, 0xc8, 0x11, 0x6c, 0x2d, 0xc5, 0x9d, 0x58, 0x46, 0x01, 0x02, 0xeb, 0x45,
	0xfc, 0xaa, 0xc7, 0xd1, 0xa5, 0xc9, 0x3d, 0xc7, 0xea, 0xaf, 0x4e, 0xd8, 0x0e, 0xa7, 0x0e, 0x2c,
	0x32, 0x42, 0x20, 0xe4, 0x56, 0xf0, 0x06, 0x83, 0xdf, 0xf1, 0x45, 0x5d, 0xca, 0x86, 0x92, 0x48,
	0xf3, 0x47, 0x48, 0xfc, 0x33, 0xe8, 0xe9, 0xba, 0xb5, 0x11, 0xe4, 0x3f, 0x98, 0x60, 0xa0, 0x3d,
	0xbf, 0xa5, 0xa4, 0x59, 0x64, 0x44, 0x40, 0xe8, 0xcb, 0x43, 0xc5, 0xdd, 0xcb, 0x0f, 0xf3, 0xbf,
	0xeb, 0xfb, 0xfc, 0x51, 0xdf, 0x28, 0xe2, 0xbd, 0x8c, 0xaf, 0x3e, 0x1a, 0xff, 0x03, 0x19, 0x0f,
	0xa6, 0x88, 0x27, 0x09, 0x8c, 0x13, 0x69, 0xa2, 0x10, 0x55, 0x6e, 0x06, 0x55, 0x49, 0xa4, 0xa1,
	0x1e, 0x1e, 0x7f, 0x1f, 0xf5, 0x1f, 0x84, 0x4b, 0x92, 0xc2, 0x58, 0x49, 0x83, 0xad, 0x1d, 0xba,
	0x3c, 0xff, 0x7a, 0xd4, 0xd3, 0xc9, 0x39, 0x1c, 0x66, 0x7a, 0xc5, 0x55, 0xd1, 0xb5, 0xb0, 0x7f,
	0xb8, 0x90, 0xce, 0xea, 0x8d, 0x8e, 0x89, 0x7f, 0x05, 0x75, 0x9a, 0x68, 0x79, 0xa3, 0x32, 0x4c,
	0xf3, 0x04, 0xa6, 0xb8, 0x2e, 0xf8, 0x4a, 0xb4, 0xbe, 0xf1, 0x81, 0xf7, 0x7c, 0x25, 0xc8, 0x25,
	0x1c, 0x2b, 0x73, 0x77, 0xed, 0x2b, 0x65, 0x1b, 0x77, 0xb1, 0x52, 0x38, 0xd4, 0xd8, 0xa1, 0xc4,
	0xef, 0xbe, 0x94, 0x86, 0x36, 0x46, 0xbb, 0x15, 0x8e, 0x9c, 0x03, 0x79, 0x7c, 0x07, 0x5f, 0x79,
	0x4a, 0x0f, 0x7a, 0xe7, 0xc9, 0x05, 0x1c, 0xf9, 0x90, 0x13, 0x3d, 0x7c, 0x88, 0xf8, 0xc3, 0x7a,
	0xaf, 0x4b, 0x3f, 0x83, 0x59, 0xff, 0x42, 0xb4, 0x85, 0xec, 0xfd, 0x87, 0x87, 0xc9, 0x1a, 0x0e,
	0xda, 0x86, 0xd9, 0x15, 0x77, 0x4a, 0x17, 0xd1, 0xe4, 0x74, 0x3c, 0xbc, 0x09, 0xb8, 0xa4, 0xfb,
	0xfe, 0xc7, 0xb9, 0xd7, 0x89, 0x6f, 0x1b, 0x3b, 0x18, 0xc6, 0xb3, 0xcc, 0xd6, 0x3f, 0xdb, 0xff,
	0xb0, 0xcd, 0xf3, 0x6e, 0x97, 0x27, 0x3c, 0xc7, 0x1e, 0x13, 0x08, 0x7d, 0xe6, 0xd8, 0xd1, 0x29,
	0xc5, 0xef, 0x26, 0xf6, 0xac, 0xe9, 0x1a, 0x7e, 0xc7, 0x5f, 0x60, 0xaf, 0x9d, 0x5f, 0xa2, 0x70,
	0x76, 0x4d, 0x3e, 0x43, 0xa0, 0x5a, 0x7f, 0x0d, 0x53, 0x53, 0x27, 0x5f, 0x1a, 0x28, 0x13, 0xff,
	0x08, 0x60, 0xb7, 0x33, 0x33, 0xc9, 0x53, 0x20, 0xbd, 0x71, 0xc9, 0x2a, 0x1b, 0x5d, 0xe2, 0x64,
	0x9a, 0x3d, 0x9c, 0x98, 0x1f, 0x2d, 0xf9, 0x36, 0x82, 0xa8, 0xb6, 0x9a, 0x9f, 0xdf, 0xa9, 0x5e,
	0x76, 0x3d, 0x7a, 0x35, 0x5c, 0xda, 0x5d, 0x3b, 0xd3, 0x63, 0xbf, 0xba, 0x69, 0x04, 0xef, 0xd3,
	0x21, 0x12, 0xb6, 0x7d, 0x6d, 0xa2, 0x2c, 0xa3, 0x6b, 0x74, 0xc1, 0xbb, 0x21, 0xa4, 0x37, 0x8f,
	0x41, 0x5b, 0x7a, 0x32, 0xc1, 0x6a, 0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x96, 0x4a, 0xee,
	0x00, 0xc2, 0x06, 0x00, 0x00,
}
