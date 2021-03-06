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
// source: te_sr_prefix.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_segment_routing_sr_prefix_infos_sr_prefix_info

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

type TeSrPrefix_KEYS struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	IgpInstance          string   `protobuf:"bytes,2,opt,name=igp_instance,json=igpInstance,proto3" json:"igp_instance,omitempty"`
	Area                 uint32   `protobuf:"varint,3,opt,name=area,proto3" json:"area,omitempty"`
	Prefix               string   `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,5,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeSrPrefix_KEYS) Reset()         { *m = TeSrPrefix_KEYS{} }
func (m *TeSrPrefix_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeSrPrefix_KEYS) ProtoMessage()    {}
func (*TeSrPrefix_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0805029a80318ca8, []int{0}
}

func (m *TeSrPrefix_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSrPrefix_KEYS.Unmarshal(m, b)
}
func (m *TeSrPrefix_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSrPrefix_KEYS.Marshal(b, m, deterministic)
}
func (m *TeSrPrefix_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSrPrefix_KEYS.Merge(m, src)
}
func (m *TeSrPrefix_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeSrPrefix_KEYS.Size(m)
}
func (m *TeSrPrefix_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSrPrefix_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeSrPrefix_KEYS proto.InternalMessageInfo

func (m *TeSrPrefix_KEYS) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *TeSrPrefix_KEYS) GetIgpInstance() string {
	if m != nil {
		return m.IgpInstance
	}
	return ""
}

func (m *TeSrPrefix_KEYS) GetArea() uint32 {
	if m != nil {
		return m.Area
	}
	return 0
}

func (m *TeSrPrefix_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *TeSrPrefix_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type TeSrPrimaryPath struct {
	OutgoingInterface    string   `protobuf:"bytes,1,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	OutgoingLabel        uint32   `protobuf:"varint,2,opt,name=outgoing_label,json=outgoingLabel,proto3" json:"outgoing_label,omitempty"`
	OutgoingStrictLabel  uint32   `protobuf:"varint,3,opt,name=outgoing_strict_label,json=outgoingStrictLabel,proto3" json:"outgoing_strict_label,omitempty"`
	NextHop              string   `protobuf:"bytes,4,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	PathId               uint32   `protobuf:"varint,5,opt,name=path_id,json=pathId,proto3" json:"path_id,omitempty"`
	BackupPathId         uint32   `protobuf:"varint,6,opt,name=backup_path_id,json=backupPathId,proto3" json:"backup_path_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeSrPrimaryPath) Reset()         { *m = TeSrPrimaryPath{} }
func (m *TeSrPrimaryPath) String() string { return proto.CompactTextString(m) }
func (*TeSrPrimaryPath) ProtoMessage()    {}
func (*TeSrPrimaryPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_0805029a80318ca8, []int{1}
}

func (m *TeSrPrimaryPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSrPrimaryPath.Unmarshal(m, b)
}
func (m *TeSrPrimaryPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSrPrimaryPath.Marshal(b, m, deterministic)
}
func (m *TeSrPrimaryPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSrPrimaryPath.Merge(m, src)
}
func (m *TeSrPrimaryPath) XXX_Size() int {
	return xxx_messageInfo_TeSrPrimaryPath.Size(m)
}
func (m *TeSrPrimaryPath) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSrPrimaryPath.DiscardUnknown(m)
}

var xxx_messageInfo_TeSrPrimaryPath proto.InternalMessageInfo

func (m *TeSrPrimaryPath) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *TeSrPrimaryPath) GetOutgoingLabel() uint32 {
	if m != nil {
		return m.OutgoingLabel
	}
	return 0
}

func (m *TeSrPrimaryPath) GetOutgoingStrictLabel() uint32 {
	if m != nil {
		return m.OutgoingStrictLabel
	}
	return 0
}

func (m *TeSrPrimaryPath) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *TeSrPrimaryPath) GetPathId() uint32 {
	if m != nil {
		return m.PathId
	}
	return 0
}

func (m *TeSrPrimaryPath) GetBackupPathId() uint32 {
	if m != nil {
		return m.BackupPathId
	}
	return 0
}

type TeSrBackupPath struct {
	OutgoingInterface         string   `protobuf:"bytes,1,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	OutgoingLabelsStack       []uint32 `protobuf:"varint,2,rep,packed,name=outgoing_labels_stack,json=outgoingLabelsStack,proto3" json:"outgoing_labels_stack,omitempty"`
	OutgoingStrictLabelsStack []uint32 `protobuf:"varint,3,rep,packed,name=outgoing_strict_labels_stack,json=outgoingStrictLabelsStack,proto3" json:"outgoing_strict_labels_stack,omitempty"`
	NextHop                   string   `protobuf:"bytes,4,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *TeSrBackupPath) Reset()         { *m = TeSrBackupPath{} }
func (m *TeSrBackupPath) String() string { return proto.CompactTextString(m) }
func (*TeSrBackupPath) ProtoMessage()    {}
func (*TeSrBackupPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_0805029a80318ca8, []int{2}
}

func (m *TeSrBackupPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSrBackupPath.Unmarshal(m, b)
}
func (m *TeSrBackupPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSrBackupPath.Marshal(b, m, deterministic)
}
func (m *TeSrBackupPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSrBackupPath.Merge(m, src)
}
func (m *TeSrBackupPath) XXX_Size() int {
	return xxx_messageInfo_TeSrBackupPath.Size(m)
}
func (m *TeSrBackupPath) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSrBackupPath.DiscardUnknown(m)
}

var xxx_messageInfo_TeSrBackupPath proto.InternalMessageInfo

func (m *TeSrBackupPath) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *TeSrBackupPath) GetOutgoingLabelsStack() []uint32 {
	if m != nil {
		return m.OutgoingLabelsStack
	}
	return nil
}

func (m *TeSrBackupPath) GetOutgoingStrictLabelsStack() []uint32 {
	if m != nil {
		return m.OutgoingStrictLabelsStack
	}
	return nil
}

func (m *TeSrBackupPath) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

type TeSrPath struct {
	PrimaryPath          *TeSrPrimaryPath `protobuf:"bytes,1,opt,name=primary_path,json=primaryPath,proto3" json:"primary_path,omitempty"`
	BackupPath           *TeSrBackupPath  `protobuf:"bytes,2,opt,name=backup_path,json=backupPath,proto3" json:"backup_path,omitempty"`
	HasBackupPath        bool             `protobuf:"varint,3,opt,name=has_backup_path,json=hasBackupPath,proto3" json:"has_backup_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TeSrPath) Reset()         { *m = TeSrPath{} }
func (m *TeSrPath) String() string { return proto.CompactTextString(m) }
func (*TeSrPath) ProtoMessage()    {}
func (*TeSrPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_0805029a80318ca8, []int{3}
}

func (m *TeSrPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSrPath.Unmarshal(m, b)
}
func (m *TeSrPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSrPath.Marshal(b, m, deterministic)
}
func (m *TeSrPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSrPath.Merge(m, src)
}
func (m *TeSrPath) XXX_Size() int {
	return xxx_messageInfo_TeSrPath.Size(m)
}
func (m *TeSrPath) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSrPath.DiscardUnknown(m)
}

var xxx_messageInfo_TeSrPath proto.InternalMessageInfo

func (m *TeSrPath) GetPrimaryPath() *TeSrPrimaryPath {
	if m != nil {
		return m.PrimaryPath
	}
	return nil
}

func (m *TeSrPath) GetBackupPath() *TeSrBackupPath {
	if m != nil {
		return m.BackupPath
	}
	return nil
}

func (m *TeSrPath) GetHasBackupPath() bool {
	if m != nil {
		return m.HasBackupPath
	}
	return false
}

type TeSrIgpArea struct {
	SegmentRoutingEnabled bool     `protobuf:"varint,1,opt,name=segment_routing_enabled,json=segmentRoutingEnabled,proto3" json:"segment_routing_enabled,omitempty"`
	StrictSpfEnabled      bool     `protobuf:"varint,2,opt,name=strict_spf_enabled,json=strictSpfEnabled,proto3" json:"strict_spf_enabled,omitempty"`
	IgpType               string   `protobuf:"bytes,3,opt,name=igp_type,json=igpType,proto3" json:"igp_type,omitempty"`
	IgpInstanceXr         string   `protobuf:"bytes,4,opt,name=igp_instance_xr,json=igpInstanceXr,proto3" json:"igp_instance_xr,omitempty"`
	IgpArea               uint32   `protobuf:"varint,5,opt,name=igp_area,json=igpArea,proto3" json:"igp_area,omitempty"`
	IgpAreaFormat         string   `protobuf:"bytes,6,opt,name=igp_area_format,json=igpAreaFormat,proto3" json:"igp_area_format,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *TeSrIgpArea) Reset()         { *m = TeSrIgpArea{} }
func (m *TeSrIgpArea) String() string { return proto.CompactTextString(m) }
func (*TeSrIgpArea) ProtoMessage()    {}
func (*TeSrIgpArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_0805029a80318ca8, []int{4}
}

func (m *TeSrIgpArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSrIgpArea.Unmarshal(m, b)
}
func (m *TeSrIgpArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSrIgpArea.Marshal(b, m, deterministic)
}
func (m *TeSrIgpArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSrIgpArea.Merge(m, src)
}
func (m *TeSrIgpArea) XXX_Size() int {
	return xxx_messageInfo_TeSrIgpArea.Size(m)
}
func (m *TeSrIgpArea) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSrIgpArea.DiscardUnknown(m)
}

var xxx_messageInfo_TeSrIgpArea proto.InternalMessageInfo

func (m *TeSrIgpArea) GetSegmentRoutingEnabled() bool {
	if m != nil {
		return m.SegmentRoutingEnabled
	}
	return false
}

func (m *TeSrIgpArea) GetStrictSpfEnabled() bool {
	if m != nil {
		return m.StrictSpfEnabled
	}
	return false
}

func (m *TeSrIgpArea) GetIgpType() string {
	if m != nil {
		return m.IgpType
	}
	return ""
}

func (m *TeSrIgpArea) GetIgpInstanceXr() string {
	if m != nil {
		return m.IgpInstanceXr
	}
	return ""
}

func (m *TeSrIgpArea) GetIgpArea() uint32 {
	if m != nil {
		return m.IgpArea
	}
	return 0
}

func (m *TeSrIgpArea) GetIgpAreaFormat() string {
	if m != nil {
		return m.IgpAreaFormat
	}
	return ""
}

type NodeIdEntry struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeIdEntry) Reset()         { *m = NodeIdEntry{} }
func (m *NodeIdEntry) String() string { return proto.CompactTextString(m) }
func (*NodeIdEntry) ProtoMessage()    {}
func (*NodeIdEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_0805029a80318ca8, []int{5}
}

func (m *NodeIdEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeIdEntry.Unmarshal(m, b)
}
func (m *NodeIdEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeIdEntry.Marshal(b, m, deterministic)
}
func (m *NodeIdEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeIdEntry.Merge(m, src)
}
func (m *NodeIdEntry) XXX_Size() int {
	return xxx_messageInfo_NodeIdEntry.Size(m)
}
func (m *NodeIdEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeIdEntry.DiscardUnknown(m)
}

var xxx_messageInfo_NodeIdEntry proto.InternalMessageInfo

func (m *NodeIdEntry) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type TeSrPrefix struct {
	PrefixXr              string         `protobuf:"bytes,50,opt,name=prefix_xr,json=prefixXr,proto3" json:"prefix_xr,omitempty"`
	PrefixLengthXr        uint32         `protobuf:"varint,51,opt,name=prefix_length_xr,json=prefixLengthXr,proto3" json:"prefix_length_xr,omitempty"`
	PrefixSid             uint32         `protobuf:"varint,52,opt,name=prefix_sid,json=prefixSid,proto3" json:"prefix_sid,omitempty"`
	PrefixStrictSid       uint32         `protobuf:"varint,53,opt,name=prefix_strict_sid,json=prefixStrictSid,proto3" json:"prefix_strict_sid,omitempty"`
	SrPath                []*TeSrPath    `protobuf:"bytes,54,rep,name=sr_path,json=srPath,proto3" json:"sr_path,omitempty"`
	IgpArea               *TeSrIgpArea   `protobuf:"bytes,55,opt,name=igp_area,json=igpArea,proto3" json:"igp_area,omitempty"`
	AdvertizingNode       []*NodeIdEntry `protobuf:"bytes,56,rep,name=advertizing_node,json=advertizingNode,proto3" json:"advertizing_node,omitempty"`
	StrictAdvertizingNode []*NodeIdEntry `protobuf:"bytes,57,rep,name=strict_advertizing_node,json=strictAdvertizingNode,proto3" json:"strict_advertizing_node,omitempty"`
	FlagR                 bool           `protobuf:"varint,58,opt,name=flag_r,json=flagR,proto3" json:"flag_r,omitempty"`
	FlagN                 bool           `protobuf:"varint,59,opt,name=flag_n,json=flagN,proto3" json:"flag_n,omitempty"`
	FlagP                 bool           `protobuf:"varint,60,opt,name=flag_p,json=flagP,proto3" json:"flag_p,omitempty"`
	FlagE                 bool           `protobuf:"varint,61,opt,name=flag_e,json=flagE,proto3" json:"flag_e,omitempty"`
	FlagV                 bool           `protobuf:"varint,62,opt,name=flag_v,json=flagV,proto3" json:"flag_v,omitempty"`
	FlagL                 bool           `protobuf:"varint,63,opt,name=flag_l,json=flagL,proto3" json:"flag_l,omitempty"`
	StrictFlagP           bool           `protobuf:"varint,64,opt,name=strict_flag_p,json=strictFlagP,proto3" json:"strict_flag_p,omitempty"`
	StrictFlagE           bool           `protobuf:"varint,65,opt,name=strict_flag_e,json=strictFlagE,proto3" json:"strict_flag_e,omitempty"`
	StrictFlagV           bool           `protobuf:"varint,66,opt,name=strict_flag_v,json=strictFlagV,proto3" json:"strict_flag_v,omitempty"`
	StrictFlagL           bool           `protobuf:"varint,67,opt,name=strict_flag_l,json=strictFlagL,proto3" json:"strict_flag_l,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}       `json:"-"`
	XXX_unrecognized      []byte         `json:"-"`
	XXX_sizecache         int32          `json:"-"`
}

func (m *TeSrPrefix) Reset()         { *m = TeSrPrefix{} }
func (m *TeSrPrefix) String() string { return proto.CompactTextString(m) }
func (*TeSrPrefix) ProtoMessage()    {}
func (*TeSrPrefix) Descriptor() ([]byte, []int) {
	return fileDescriptor_0805029a80318ca8, []int{6}
}

func (m *TeSrPrefix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSrPrefix.Unmarshal(m, b)
}
func (m *TeSrPrefix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSrPrefix.Marshal(b, m, deterministic)
}
func (m *TeSrPrefix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSrPrefix.Merge(m, src)
}
func (m *TeSrPrefix) XXX_Size() int {
	return xxx_messageInfo_TeSrPrefix.Size(m)
}
func (m *TeSrPrefix) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSrPrefix.DiscardUnknown(m)
}

var xxx_messageInfo_TeSrPrefix proto.InternalMessageInfo

func (m *TeSrPrefix) GetPrefixXr() string {
	if m != nil {
		return m.PrefixXr
	}
	return ""
}

func (m *TeSrPrefix) GetPrefixLengthXr() uint32 {
	if m != nil {
		return m.PrefixLengthXr
	}
	return 0
}

func (m *TeSrPrefix) GetPrefixSid() uint32 {
	if m != nil {
		return m.PrefixSid
	}
	return 0
}

func (m *TeSrPrefix) GetPrefixStrictSid() uint32 {
	if m != nil {
		return m.PrefixStrictSid
	}
	return 0
}

func (m *TeSrPrefix) GetSrPath() []*TeSrPath {
	if m != nil {
		return m.SrPath
	}
	return nil
}

func (m *TeSrPrefix) GetIgpArea() *TeSrIgpArea {
	if m != nil {
		return m.IgpArea
	}
	return nil
}

func (m *TeSrPrefix) GetAdvertizingNode() []*NodeIdEntry {
	if m != nil {
		return m.AdvertizingNode
	}
	return nil
}

func (m *TeSrPrefix) GetStrictAdvertizingNode() []*NodeIdEntry {
	if m != nil {
		return m.StrictAdvertizingNode
	}
	return nil
}

func (m *TeSrPrefix) GetFlagR() bool {
	if m != nil {
		return m.FlagR
	}
	return false
}

func (m *TeSrPrefix) GetFlagN() bool {
	if m != nil {
		return m.FlagN
	}
	return false
}

func (m *TeSrPrefix) GetFlagP() bool {
	if m != nil {
		return m.FlagP
	}
	return false
}

func (m *TeSrPrefix) GetFlagE() bool {
	if m != nil {
		return m.FlagE
	}
	return false
}

func (m *TeSrPrefix) GetFlagV() bool {
	if m != nil {
		return m.FlagV
	}
	return false
}

func (m *TeSrPrefix) GetFlagL() bool {
	if m != nil {
		return m.FlagL
	}
	return false
}

func (m *TeSrPrefix) GetStrictFlagP() bool {
	if m != nil {
		return m.StrictFlagP
	}
	return false
}

func (m *TeSrPrefix) GetStrictFlagE() bool {
	if m != nil {
		return m.StrictFlagE
	}
	return false
}

func (m *TeSrPrefix) GetStrictFlagV() bool {
	if m != nil {
		return m.StrictFlagV
	}
	return false
}

func (m *TeSrPrefix) GetStrictFlagL() bool {
	if m != nil {
		return m.StrictFlagL
	}
	return false
}

func init() {
	proto.RegisterType((*TeSrPrefix_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.segment_routing.sr_prefix_infos.sr_prefix_info.te_sr_prefix_KEYS")
	proto.RegisterType((*TeSrPrimaryPath)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.segment_routing.sr_prefix_infos.sr_prefix_info.te_sr_primary_path")
	proto.RegisterType((*TeSrBackupPath)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.segment_routing.sr_prefix_infos.sr_prefix_info.te_sr_backup_path")
	proto.RegisterType((*TeSrPath)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.segment_routing.sr_prefix_infos.sr_prefix_info.te_sr_path")
	proto.RegisterType((*TeSrIgpArea)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.segment_routing.sr_prefix_infos.sr_prefix_info.te_sr_igp_area")
	proto.RegisterType((*NodeIdEntry)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.segment_routing.sr_prefix_infos.sr_prefix_info.node_id_entry")
	proto.RegisterType((*TeSrPrefix)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.segment_routing.sr_prefix_infos.sr_prefix_info.te_sr_prefix")
}

func init() { proto.RegisterFile("te_sr_prefix.proto", fileDescriptor_0805029a80318ca8) }

var fileDescriptor_0805029a80318ca8 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdb, 0x6a, 0x24, 0x45,
	0x18, 0xa6, 0x93, 0xdd, 0x99, 0xc9, 0x3f, 0xe9, 0x1c, 0x4a, 0x62, 0x2a, 0x1e, 0x20, 0xb6, 0x07,
	0x06, 0xd1, 0xb9, 0x98, 0xd5, 0xf5, 0xec, 0x9a, 0x95, 0x2c, 0x06, 0xc3, 0x12, 0x7a, 0x64, 0x89,
	0x20, 0x94, 0x9d, 0xee, 0x9a, 0x9e, 0x62, 0x7b, 0xba, 0x8b, 0xaa, 0x9a, 0x30, 0xa3, 0xb7, 0x82,
	0xe0, 0x0b, 0xf8, 0x00, 0x3e, 0x8b, 0x2f, 0x20, 0xf8, 0x2e, 0x5e, 0x78, 0x21, 0xf5, 0x77, 0xf5,
	0x61, 0x7b, 0xd7, 0x0b, 0x61, 0xe3, 0x5d, 0xff, 0xdf, 0xf7, 0x55, 0xfd, 0xe7, 0xa2, 0x81, 0x18,
	0xce, 0xb4, 0x62, 0x52, 0xf1, 0x99, 0x58, 0x8d, 0xa5, 0x2a, 0x4c, 0x41, 0x2e, 0x62, 0xa1, 0xe3,
	0x82, 0x89, 0x42, 0xb3, 0x95, 0x62, 0x0b, 0x99, 0x69, 0x66, 0x38, 0x2b, 0x24, 0x57, 0x63, 0x67,
	0x8c, 0x35, 0x4f, 0x17, 0x3c, 0x37, 0x4c, 0x15, 0x4b, 0x23, 0xf2, 0x74, 0x5c, 0x5f, 0xc1, 0x44,
	0x3e, 0x2b, 0x74, 0xc7, 0x0e, 0x7e, 0xf3, 0x60, 0xbf, 0xed, 0x88, 0x7d, 0x7d, 0xfa, 0xed, 0x94,
	0xbc, 0x04, 0x03, 0x74, 0x18, 0x17, 0x19, 0xf5, 0x8e, 0xbd, 0xd1, 0x56, 0x58, 0xdb, 0xe4, 0x35,
	0xd8, 0x16, 0xa9, 0x64, 0x22, 0xd7, 0x26, 0xca, 0x63, 0x4e, 0x37, 0x90, 0x1f, 0x8a, 0x54, 0x9e,
	0x39, 0x88, 0x10, 0xb8, 0x15, 0x29, 0x1e, 0xd1, 0xcd, 0x63, 0x6f, 0xe4, 0x87, 0xf8, 0x4d, 0x5e,
	0x84, 0x5e, 0xe9, 0x81, 0xde, 0xc2, 0x03, 0xce, 0x22, 0xaf, 0x83, 0xef, 0x3c, 0x67, 0x3c, 0x4f,
	0xcd, 0x9c, 0xde, 0xc6, 0x43, 0xdb, 0x25, 0x78, 0x8e, 0x58, 0xf0, 0x97, 0xd7, 0x94, 0x43, 0x2c,
	0x22, 0xb5, 0x66, 0x32, 0x32, 0x73, 0xf2, 0x2e, 0x90, 0x62, 0x69, 0xd2, 0x42, 0xe4, 0x29, 0x13,
	0xb9, 0xe1, 0x6a, 0x16, 0xc5, 0xdc, 0x05, 0xbc, 0x5f, 0x31, 0x67, 0x15, 0x41, 0xde, 0x84, 0x9d,
	0x5a, 0x9e, 0x45, 0x57, 0x3c, 0xc3, 0xd8, 0xfd, 0xd0, 0xaf, 0xd0, 0x73, 0x0b, 0x92, 0x09, 0x1c,
	0xd4, 0x32, 0x6d, 0x94, 0x88, 0x8d, 0x53, 0x97, 0xe9, 0xbc, 0x50, 0x91, 0x53, 0xe4, 0xca, 0x33,
	0x47, 0x30, 0xc8, 0xf9, 0xca, 0xb0, 0x79, 0x21, 0x5d, 0x7e, 0x7d, 0x6b, 0x7f, 0x55, 0x48, 0x72,
	0x08, 0x7d, 0x1b, 0x2c, 0x13, 0x89, 0x4b, 0xad, 0x67, 0xcd, 0xb3, 0x84, 0xbc, 0x01, 0x3b, 0x57,
	0x51, 0xfc, 0x78, 0x29, 0x59, 0xc5, 0xf7, 0xca, 0xd4, 0x4b, 0xf4, 0x02, 0x55, 0xc1, 0x1f, 0x75,
	0x83, 0x5a, 0xe2, 0xff, 0x9a, 0x79, 0x3b, 0x25, 0xcc, 0x45, 0x33, 0x6d, 0xa2, 0xf8, 0x31, 0xdd,
	0x38, 0xde, 0x6c, 0xa7, 0x84, 0xc9, 0xe8, 0xa9, 0xa5, 0xc8, 0x3d, 0x78, 0xe5, 0x99, 0x65, 0xa8,
	0x8e, 0x6e, 0xe2, 0xd1, 0xa3, 0x67, 0x54, 0xc3, 0x5d, 0xf0, 0xef, 0x35, 0x09, 0x7e, 0xdf, 0x00,
	0x70, 0xfd, 0xb4, 0xd9, 0xfc, 0xec, 0xc1, 0x76, 0xbb, 0xb1, 0x98, 0xc8, 0x70, 0x92, 0x8c, 0x9f,
	0xf7, 0xb8, 0x8f, 0x9f, 0x1e, 0xa2, 0x70, 0xe8, 0x2c, 0x5b, 0x6f, 0xf2, 0x93, 0x07, 0xc3, 0x56,
	0x9d, 0x71, 0x40, 0x86, 0x93, 0xf8, 0xa6, 0x02, 0x69, 0xb9, 0x0a, 0xa1, 0x69, 0x3b, 0x79, 0x0b,
	0x76, 0xe7, 0x91, 0x6e, 0xd3, 0x38, 0x7c, 0x83, 0xd0, 0x9f, 0x47, 0xfa, 0x7e, 0xad, 0x0b, 0xfe,
	0xf6, 0x60, 0xa7, 0xbc, 0xc9, 0xae, 0x24, 0xee, 0xd9, 0x5d, 0x38, 0xec, 0x84, 0xc2, 0x78, 0x1e,
	0x5d, 0x65, 0x3c, 0xc1, 0xaa, 0x0e, 0xc2, 0x03, 0x47, 0x87, 0x25, 0x7b, 0x5a, 0x92, 0xe4, 0x1d,
	0x20, 0xae, 0xcb, 0x5a, 0xce, 0xea, 0x23, 0x1b, 0x78, 0x64, 0xaf, 0x64, 0xa6, 0x72, 0x56, 0xa9,
	0x8f, 0x60, 0x60, 0x3d, 0x9a, 0xb5, 0xe4, 0x18, 0xd9, 0x56, 0xd8, 0x17, 0xa9, 0xfc, 0x66, 0x2d,
	0xb9, 0x8d, 0xbd, 0xfd, 0x3e, 0xb0, 0x95, 0x72, 0xdd, 0xf7, 0x5b, 0x4f, 0xc4, 0xa5, 0xaa, 0xae,
	0xc0, 0x87, 0xa2, 0x5c, 0x0c, 0x7b, 0xc5, 0x89, 0xcd, 0xc1, 0x5d, 0x61, 0x29, 0x36, 0x2b, 0xd4,
	0x22, 0x32, 0xb8, 0x1a, 0xe5, 0x15, 0x56, 0xf1, 0x00, 0xc1, 0x60, 0x04, 0x7e, 0x5e, 0x24, 0x9c,
	0x89, 0x84, 0xf1, 0xdc, 0xa8, 0xb5, 0xdd, 0x35, 0x07, 0xb8, 0x5d, 0xe8, 0x59, 0xf3, 0x2c, 0x09,
	0xfe, 0xec, 0xc3, 0x76, 0xfb, 0x99, 0x23, 0x2f, 0xc3, 0x96, 0x6b, 0xc7, 0x4a, 0xd1, 0x49, 0xf5,
	0xc4, 0x59, 0xe0, 0x52, 0x91, 0x11, 0xec, 0x3d, 0xf1, 0x26, 0x59, 0xcd, 0x1d, 0x0c, 0x71, 0xa7,
	0xfd, 0x2c, 0x5d, 0x2a, 0xf2, 0x2a, 0x80, 0x53, 0x6a, 0x91, 0xd0, 0xf7, 0x50, 0xe3, 0x2e, 0x9e,
	0x8a, 0x84, 0xbc, 0x0d, 0xfb, 0x15, 0xed, 0x6a, 0x2b, 0x12, 0xfa, 0x3e, 0xaa, 0x76, 0x9d, 0xaa,
	0xac, 0xac, 0x48, 0xc8, 0x12, 0xfa, 0x6e, 0x1f, 0xe8, 0xdd, 0xe3, 0xcd, 0xd1, 0x70, 0xf2, 0xdd,
	0x8d, 0x8d, 0xbf, 0x1d, 0xb7, 0x9e, 0x56, 0x38, 0x6a, 0x3f, 0xb6, 0xda, 0xf0, 0x01, 0x4e, 0xfb,
	0xf7, 0x37, 0xe5, 0xb7, 0xf2, 0xd3, 0x34, 0xfa, 0x17, 0x0f, 0xf6, 0xa2, 0xe4, 0x9a, 0x2b, 0x23,
	0x7e, 0xb0, 0x93, 0x6a, 0xbb, 0x45, 0x3f, 0xc4, 0xec, 0xd9, 0xf3, 0x8f, 0xe2, 0x89, 0x59, 0x09,
	0x77, 0x5b, 0x8e, 0x1f, 0x16, 0x09, 0x27, 0xbf, 0x7a, 0x70, 0xe8, 0xda, 0xf4, 0x54, 0x4c, 0x1f,
	0xfd, 0x3f, 0x31, 0x1d, 0x94, 0xfe, 0x4f, 0x3a, 0x91, 0x1d, 0x40, 0x6f, 0x96, 0x45, 0x29, 0x53,
	0xf4, 0x63, 0xdc, 0xc7, 0xdb, 0xd6, 0x0a, 0x6b, 0x38, 0xa7, 0x9f, 0x34, 0xf0, 0xc3, 0x1a, 0x96,
	0xf4, 0xd3, 0x06, 0xbe, 0xa8, 0x61, 0x4e, 0x3f, 0x6b, 0xe0, 0xd3, 0x1a, 0xbe, 0xa6, 0x9f, 0x37,
	0xf0, 0xa3, 0x1a, 0xce, 0xe8, 0xbd, 0x06, 0x3e, 0x27, 0x01, 0xf8, 0xae, 0x44, 0xce, 0xc5, 0x17,
	0xc8, 0x0e, 0x4b, 0xf0, 0x01, 0x3a, 0xea, 0x68, 0x38, 0x3d, 0xe9, 0x6a, 0x4e, 0xbb, 0x9a, 0x6b,
	0x7a, 0xbf, 0xab, 0x79, 0xd4, 0xd5, 0x64, 0xf4, 0xcb, 0xae, 0xe6, 0xfc, 0xaa, 0x87, 0xbf, 0x25,
	0x77, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xd0, 0x41, 0xd3, 0x2d, 0x09, 0x00, 0x00,
}
