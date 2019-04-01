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
// source: mpls_lsd_frr_db_entry.proto

package cisco_ios_xr_mpls_lsd_oper_mpls_lsd_nodes_mpls_lsd_node_frr_database_tunnel_heads_tunnel_head

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

type MplsLsdFrrDbEntry_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdFrrDbEntry_KEYS) Reset()         { *m = MplsLsdFrrDbEntry_KEYS{} }
func (m *MplsLsdFrrDbEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLsdFrrDbEntry_KEYS) ProtoMessage()    {}
func (*MplsLsdFrrDbEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f95d019357bb9c54, []int{0}
}

func (m *MplsLsdFrrDbEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdFrrDbEntry_KEYS.Unmarshal(m, b)
}
func (m *MplsLsdFrrDbEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdFrrDbEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLsdFrrDbEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdFrrDbEntry_KEYS.Merge(m, src)
}
func (m *MplsLsdFrrDbEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLsdFrrDbEntry_KEYS.Size(m)
}
func (m *MplsLsdFrrDbEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdFrrDbEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdFrrDbEntry_KEYS proto.InternalMessageInfo

func (m *MplsLsdFrrDbEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MplsLsdFrrDbEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type MplsLsdFrrDbEntryHeadId struct {
	DestinationPrefix       string   `protobuf:"bytes,1,opt,name=destination_prefix,json=destinationPrefix,proto3" json:"destination_prefix,omitempty"`
	DestinationPrefixLength uint32   `protobuf:"varint,2,opt,name=destination_prefix_length,json=destinationPrefixLength,proto3" json:"destination_prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *MplsLsdFrrDbEntryHeadId) Reset()         { *m = MplsLsdFrrDbEntryHeadId{} }
func (m *MplsLsdFrrDbEntryHeadId) String() string { return proto.CompactTextString(m) }
func (*MplsLsdFrrDbEntryHeadId) ProtoMessage()    {}
func (*MplsLsdFrrDbEntryHeadId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f95d019357bb9c54, []int{1}
}

func (m *MplsLsdFrrDbEntryHeadId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdFrrDbEntryHeadId.Unmarshal(m, b)
}
func (m *MplsLsdFrrDbEntryHeadId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdFrrDbEntryHeadId.Marshal(b, m, deterministic)
}
func (m *MplsLsdFrrDbEntryHeadId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdFrrDbEntryHeadId.Merge(m, src)
}
func (m *MplsLsdFrrDbEntryHeadId) XXX_Size() int {
	return xxx_messageInfo_MplsLsdFrrDbEntryHeadId.Size(m)
}
func (m *MplsLsdFrrDbEntryHeadId) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdFrrDbEntryHeadId.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdFrrDbEntryHeadId proto.InternalMessageInfo

func (m *MplsLsdFrrDbEntryHeadId) GetDestinationPrefix() string {
	if m != nil {
		return m.DestinationPrefix
	}
	return ""
}

func (m *MplsLsdFrrDbEntryHeadId) GetDestinationPrefixLength() uint32 {
	if m != nil {
		return m.DestinationPrefixLength
	}
	return 0
}

type MplsLsdFrrDbEntryMidId struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	Lspid                uint32   `protobuf:"varint,2,opt,name=lspid,proto3" json:"lspid,omitempty"`
	TunnelId             uint32   `protobuf:"varint,3,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdFrrDbEntryMidId) Reset()         { *m = MplsLsdFrrDbEntryMidId{} }
func (m *MplsLsdFrrDbEntryMidId) String() string { return proto.CompactTextString(m) }
func (*MplsLsdFrrDbEntryMidId) ProtoMessage()    {}
func (*MplsLsdFrrDbEntryMidId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f95d019357bb9c54, []int{2}
}

func (m *MplsLsdFrrDbEntryMidId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdFrrDbEntryMidId.Unmarshal(m, b)
}
func (m *MplsLsdFrrDbEntryMidId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdFrrDbEntryMidId.Marshal(b, m, deterministic)
}
func (m *MplsLsdFrrDbEntryMidId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdFrrDbEntryMidId.Merge(m, src)
}
func (m *MplsLsdFrrDbEntryMidId) XXX_Size() int {
	return xxx_messageInfo_MplsLsdFrrDbEntryMidId.Size(m)
}
func (m *MplsLsdFrrDbEntryMidId) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdFrrDbEntryMidId.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdFrrDbEntryMidId proto.InternalMessageInfo

func (m *MplsLsdFrrDbEntryMidId) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsLsdFrrDbEntryMidId) GetLspid() uint32 {
	if m != nil {
		return m.Lspid
	}
	return 0
}

func (m *MplsLsdFrrDbEntryMidId) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

type MplsLsdFrrDbEntryGenId struct {
	Role                 string                   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Head                 *MplsLsdFrrDbEntryHeadId `protobuf:"bytes,2,opt,name=head,proto3" json:"head,omitempty"`
	Midpoint             *MplsLsdFrrDbEntryMidId  `protobuf:"bytes,3,opt,name=midpoint,proto3" json:"midpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MplsLsdFrrDbEntryGenId) Reset()         { *m = MplsLsdFrrDbEntryGenId{} }
func (m *MplsLsdFrrDbEntryGenId) String() string { return proto.CompactTextString(m) }
func (*MplsLsdFrrDbEntryGenId) ProtoMessage()    {}
func (*MplsLsdFrrDbEntryGenId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f95d019357bb9c54, []int{3}
}

func (m *MplsLsdFrrDbEntryGenId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdFrrDbEntryGenId.Unmarshal(m, b)
}
func (m *MplsLsdFrrDbEntryGenId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdFrrDbEntryGenId.Marshal(b, m, deterministic)
}
func (m *MplsLsdFrrDbEntryGenId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdFrrDbEntryGenId.Merge(m, src)
}
func (m *MplsLsdFrrDbEntryGenId) XXX_Size() int {
	return xxx_messageInfo_MplsLsdFrrDbEntryGenId.Size(m)
}
func (m *MplsLsdFrrDbEntryGenId) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdFrrDbEntryGenId.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdFrrDbEntryGenId proto.InternalMessageInfo

func (m *MplsLsdFrrDbEntryGenId) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *MplsLsdFrrDbEntryGenId) GetHead() *MplsLsdFrrDbEntryHeadId {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *MplsLsdFrrDbEntryGenId) GetMidpoint() *MplsLsdFrrDbEntryMidId {
	if m != nil {
		return m.Midpoint
	}
	return nil
}

type MplsLsdFrrOutPath struct {
	OutInterface         string   `protobuf:"bytes,1,opt,name=out_interface,json=outInterface,proto3" json:"out_interface,omitempty"`
	OutLabel             uint32   `protobuf:"varint,2,opt,name=out_label,json=outLabel,proto3" json:"out_label,omitempty"`
	Ipv4NextHop          string   `protobuf:"bytes,3,opt,name=ipv4_next_hop,json=ipv4NextHop,proto3" json:"ipv4_next_hop,omitempty"`
	FrrInterface         string   `protobuf:"bytes,4,opt,name=frr_interface,json=frrInterface,proto3" json:"frr_interface,omitempty"`
	FrrLabel             uint32   `protobuf:"varint,5,opt,name=frr_label,json=frrLabel,proto3" json:"frr_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdFrrOutPath) Reset()         { *m = MplsLsdFrrOutPath{} }
func (m *MplsLsdFrrOutPath) String() string { return proto.CompactTextString(m) }
func (*MplsLsdFrrOutPath) ProtoMessage()    {}
func (*MplsLsdFrrOutPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_f95d019357bb9c54, []int{4}
}

func (m *MplsLsdFrrOutPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdFrrOutPath.Unmarshal(m, b)
}
func (m *MplsLsdFrrOutPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdFrrOutPath.Marshal(b, m, deterministic)
}
func (m *MplsLsdFrrOutPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdFrrOutPath.Merge(m, src)
}
func (m *MplsLsdFrrOutPath) XXX_Size() int {
	return xxx_messageInfo_MplsLsdFrrOutPath.Size(m)
}
func (m *MplsLsdFrrOutPath) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdFrrOutPath.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdFrrOutPath proto.InternalMessageInfo

func (m *MplsLsdFrrOutPath) GetOutInterface() string {
	if m != nil {
		return m.OutInterface
	}
	return ""
}

func (m *MplsLsdFrrOutPath) GetOutLabel() uint32 {
	if m != nil {
		return m.OutLabel
	}
	return 0
}

func (m *MplsLsdFrrOutPath) GetIpv4NextHop() string {
	if m != nil {
		return m.Ipv4NextHop
	}
	return ""
}

func (m *MplsLsdFrrOutPath) GetFrrInterface() string {
	if m != nil {
		return m.FrrInterface
	}
	return ""
}

func (m *MplsLsdFrrOutPath) GetFrrLabel() uint32 {
	if m != nil {
		return m.FrrLabel
	}
	return 0
}

type MplsLsdFrrDbEntry struct {
	FrrEntryId           *MplsLsdFrrDbEntryGenId `protobuf:"bytes,50,opt,name=frr_entry_id,json=frrEntryId,proto3" json:"frr_entry_id,omitempty"`
	TunnelInterface      string                  `protobuf:"bytes,51,opt,name=tunnel_interface,json=tunnelInterface,proto3" json:"tunnel_interface,omitempty"`
	InputLabel           uint32                  `protobuf:"varint,52,opt,name=input_label,json=inputLabel,proto3" json:"input_label,omitempty"`
	OutPath              []*MplsLsdFrrOutPath    `protobuf:"bytes,53,rep,name=out_path,json=outPath,proto3" json:"out_path,omitempty"`
	EntryFrrState        string                  `protobuf:"bytes,54,opt,name=entry_frr_state,json=entryFrrState,proto3" json:"entry_frr_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MplsLsdFrrDbEntry) Reset()         { *m = MplsLsdFrrDbEntry{} }
func (m *MplsLsdFrrDbEntry) String() string { return proto.CompactTextString(m) }
func (*MplsLsdFrrDbEntry) ProtoMessage()    {}
func (*MplsLsdFrrDbEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f95d019357bb9c54, []int{5}
}

func (m *MplsLsdFrrDbEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdFrrDbEntry.Unmarshal(m, b)
}
func (m *MplsLsdFrrDbEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdFrrDbEntry.Marshal(b, m, deterministic)
}
func (m *MplsLsdFrrDbEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdFrrDbEntry.Merge(m, src)
}
func (m *MplsLsdFrrDbEntry) XXX_Size() int {
	return xxx_messageInfo_MplsLsdFrrDbEntry.Size(m)
}
func (m *MplsLsdFrrDbEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdFrrDbEntry.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdFrrDbEntry proto.InternalMessageInfo

func (m *MplsLsdFrrDbEntry) GetFrrEntryId() *MplsLsdFrrDbEntryGenId {
	if m != nil {
		return m.FrrEntryId
	}
	return nil
}

func (m *MplsLsdFrrDbEntry) GetTunnelInterface() string {
	if m != nil {
		return m.TunnelInterface
	}
	return ""
}

func (m *MplsLsdFrrDbEntry) GetInputLabel() uint32 {
	if m != nil {
		return m.InputLabel
	}
	return 0
}

func (m *MplsLsdFrrDbEntry) GetOutPath() []*MplsLsdFrrOutPath {
	if m != nil {
		return m.OutPath
	}
	return nil
}

func (m *MplsLsdFrrDbEntry) GetEntryFrrState() string {
	if m != nil {
		return m.EntryFrrState
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsLsdFrrDbEntry_KEYS)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.frr_database.tunnel_heads.tunnel_head.mpls_lsd_frr_db_entry_KEYS")
	proto.RegisterType((*MplsLsdFrrDbEntryHeadId)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.frr_database.tunnel_heads.tunnel_head.mpls_lsd_frr_db_entry_head_id")
	proto.RegisterType((*MplsLsdFrrDbEntryMidId)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.frr_database.tunnel_heads.tunnel_head.mpls_lsd_frr_db_entry_mid_id")
	proto.RegisterType((*MplsLsdFrrDbEntryGenId)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.frr_database.tunnel_heads.tunnel_head.mpls_lsd_frr_db_entry_gen_id")
	proto.RegisterType((*MplsLsdFrrOutPath)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.frr_database.tunnel_heads.tunnel_head.mpls_lsd_frr_out_path")
	proto.RegisterType((*MplsLsdFrrDbEntry)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.frr_database.tunnel_heads.tunnel_head.mpls_lsd_frr_db_entry")
}

func init() { proto.RegisterFile("mpls_lsd_frr_db_entry.proto", fileDescriptor_f95d019357bb9c54) }

var fileDescriptor_f95d019357bb9c54 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x55, 0xba, 0x5b, 0xd8, 0x9d, 0x65, 0x5b, 0xb0, 0x40, 0x04, 0x0a, 0xa2, 0x0a, 0x2a, 0x2a,
	0x07, 0x72, 0xd8, 0x16, 0x0e, 0xdc, 0x38, 0x14, 0xb1, 0xa2, 0xaa, 0xaa, 0xf4, 0xc4, 0x01, 0x19,
	0xef, 0xda, 0xe9, 0x5a, 0x4a, 0x6c, 0xcb, 0x76, 0x50, 0x10, 0x7c, 0x01, 0x17, 0xb8, 0x71, 0xe2,
	0x53, 0x10, 0x9f, 0xc2, 0xaf, 0x20, 0xdb, 0x49, 0xe8, 0xd2, 0x85, 0x1b, 0xf4, 0xe6, 0x79, 0x33,
	0xf1, 0x9b, 0xf7, 0x66, 0x62, 0xd8, 0x2a, 0x55, 0x61, 0x70, 0x61, 0x28, 0xce, 0xb5, 0xc6, 0x74,
	0x86, 0x99, 0xb0, 0xfa, 0x5d, 0xaa, 0xb4, 0xb4, 0x12, 0xbd, 0x9e, 0x73, 0x33, 0x97, 0x98, 0x4b,
	0x83, 0x6b, 0x8d, 0xbb, 0x4a, 0xa9, 0x98, 0x4e, 0xbb, 0x48, 0x48, 0xca, 0xcc, 0x72, 0x98, 0xfa,
	0xbb, 0x88, 0x25, 0x33, 0x62, 0x58, 0x6a, 0x2b, 0x21, 0x58, 0x81, 0x17, 0x8c, 0x50, 0x73, 0x36,
	0x48, 0xde, 0xc0, 0xed, 0x95, 0xec, 0xf8, 0xe5, 0xc1, 0xab, 0x13, 0xb4, 0x05, 0x43, 0x77, 0x17,
	0x16, 0xa4, 0x64, 0x71, 0xb4, 0x1d, 0xed, 0x0e, 0xb3, 0x81, 0x03, 0x8e, 0x48, 0xc9, 0xd0, 0x0e,
	0x6c, 0x70, 0x61, 0x99, 0xce, 0xc9, 0xbc, 0xa9, 0x58, 0xf3, 0x15, 0xe3, 0x0e, 0x75, 0x65, 0xc9,
	0xc7, 0x08, 0xee, 0xae, 0xa6, 0x70, 0x0d, 0x60, 0x4e, 0xd1, 0x23, 0x40, 0x94, 0x19, 0xcb, 0x05,
	0xb1, 0x5c, 0x0a, 0xac, 0x34, 0xcb, 0x79, 0xdd, 0xd0, 0x5d, 0x3b, 0x93, 0x39, 0xf6, 0x09, 0xf4,
	0x14, 0x6e, 0x9d, 0x2f, 0xc7, 0x05, 0x13, 0xa7, 0x76, 0xe1, 0x5b, 0x18, 0x67, 0x37, 0xcf, 0x7d,
	0x75, 0xe8, 0xd3, 0x49, 0x0d, 0x77, 0x56, 0xf7, 0x52, 0x72, 0xdf, 0xca, 0x0e, 0x6c, 0x18, 0x59,
	0xe9, 0x39, 0xc3, 0x84, 0x52, 0xcd, 0x8c, 0x69, 0xda, 0x18, 0x07, 0xf4, 0x59, 0x00, 0xd1, 0x75,
	0x58, 0x2f, 0x8c, 0xe2, 0xb4, 0xa1, 0x0b, 0x81, 0x73, 0xab, 0xb1, 0x96, 0xd3, 0xb8, 0xe7, 0x33,
	0x83, 0x00, 0x4c, 0x69, 0xf2, 0x63, 0xed, 0x4f, 0xd4, 0xa7, 0x4c, 0x38, 0x6a, 0x04, 0x7d, 0x2d,
	0x8b, 0xd6, 0x66, 0x7f, 0x46, 0x9f, 0x23, 0xe8, 0x3b, 0x97, 0x3c, 0xcf, 0x68, 0xf2, 0x21, 0xfd,
	0xa7, 0xcb, 0x90, 0xfe, 0x75, 0x4c, 0x99, 0xef, 0x04, 0x7d, 0x89, 0x60, 0x50, 0x72, 0xaa, 0x24,
	0x17, 0xd6, 0x8b, 0x1c, 0x4d, 0xde, 0x5f, 0x48, 0x5b, 0x61, 0x62, 0x59, 0xd7, 0x4c, 0xf2, 0x3d,
	0x82, 0x1b, 0x4b, 0xa5, 0xb2, 0xb2, 0x58, 0x11, 0xbb, 0x40, 0xf7, 0x61, 0xec, 0xce, 0xdd, 0x5e,
	0x36, 0x1e, 0x5f, 0x91, 0x95, 0x9d, 0xb6, 0x98, 0x9b, 0x9e, 0x2b, 0x2a, 0xc8, 0x8c, 0x15, 0xcd,
	0x5c, 0x07, 0xb2, 0xb2, 0x87, 0x2e, 0x46, 0x09, 0x8c, 0xb9, 0x7a, 0xbb, 0x8f, 0x05, 0xab, 0x2d,
	0x5e, 0x48, 0xe5, 0x95, 0x0f, 0xb3, 0x91, 0x03, 0x8f, 0x58, 0x6d, 0x5f, 0x48, 0xe5, 0x58, 0x1c,
	0xeb, 0x2f, 0x96, 0x7e, 0x60, 0xc9, 0xb5, 0x5e, 0x62, 0x71, 0x45, 0x81, 0x65, 0x3d, 0xb0, 0xe4,
	0x5a, 0x7b, 0x96, 0xe4, 0x5b, 0xef, 0x37, 0x05, 0xad, 0x58, 0xf4, 0x35, 0x02, 0x77, 0x4f, 0x23,
	0x9d, 0xd3, 0x78, 0x72, 0x81, 0xce, 0x87, 0x85, 0xcd, 0x20, 0xd7, 0xfa, 0xc0, 0x01, 0x53, 0x8a,
	0x1e, 0xc2, 0xd5, 0x76, 0xf5, 0x3b, 0xf9, 0x7b, 0x5e, 0xfe, 0x66, 0xf3, 0x07, 0x74, 0x0e, 0xdc,
	0x83, 0x11, 0x17, 0xaa, 0x73, 0x7a, 0xdf, 0x7b, 0x00, 0x1e, 0x0a, 0x5e, 0x7f, 0x8a, 0x60, 0xd0,
	0x8e, 0x2e, 0x7e, 0xbc, 0xdd, 0xdb, 0x1d, 0x4d, 0xec, 0xff, 0xd4, 0xd9, 0x72, 0x67, 0x97, 0x65,
	0x65, 0x8f, 0xdd, 0xfe, 0x3c, 0x80, 0xcd, 0xa0, 0xdc, 0xa5, 0x8d, 0x25, 0x96, 0xc5, 0x4f, 0xc2,
	0xb3, 0xe0, 0xe1, 0xe7, 0x5a, 0x9f, 0x38, 0x70, 0x76, 0xc9, 0x3f, 0xd9, 0x7b, 0x3f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x76, 0xd9, 0x36, 0x15, 0xd1, 0x05, 0x00, 0x00,
}