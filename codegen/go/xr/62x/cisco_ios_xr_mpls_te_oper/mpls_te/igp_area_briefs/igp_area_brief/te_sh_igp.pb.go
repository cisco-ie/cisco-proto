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
// source: te_sh_igp.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_igp_area_briefs_igp_area_brief

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

type TeShIgp_KEYS struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ProcessTag           string   `protobuf:"bytes,2,opt,name=process_tag,json=processTag,proto3" json:"process_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeShIgp_KEYS) Reset()         { *m = TeShIgp_KEYS{} }
func (m *TeShIgp_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeShIgp_KEYS) ProtoMessage()    {}
func (*TeShIgp_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d3a45ebb1af4017, []int{0}
}

func (m *TeShIgp_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeShIgp_KEYS.Unmarshal(m, b)
}
func (m *TeShIgp_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeShIgp_KEYS.Marshal(b, m, deterministic)
}
func (m *TeShIgp_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeShIgp_KEYS.Merge(m, src)
}
func (m *TeShIgp_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeShIgp_KEYS.Size(m)
}
func (m *TeShIgp_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeShIgp_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeShIgp_KEYS proto.InternalMessageInfo

func (m *TeShIgp_KEYS) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *TeShIgp_KEYS) GetProcessTag() string {
	if m != nil {
		return m.ProcessTag
	}
	return ""
}

type TeShIgpAreaStats struct {
	AreaAdjacencyRequestMessages    uint32   `protobuf:"varint,1,opt,name=area_adjacency_request_messages,json=areaAdjacencyRequestMessages,proto3" json:"area_adjacency_request_messages,omitempty"`
	AreaAdjacencyAnnounceMessages   uint32   `protobuf:"varint,2,opt,name=area_adjacency_announce_messages,json=areaAdjacencyAnnounceMessages,proto3" json:"area_adjacency_announce_messages,omitempty"`
	AreaFloodMessages               uint32   `protobuf:"varint,3,opt,name=area_flood_messages,json=areaFloodMessages,proto3" json:"area_flood_messages,omitempty"`
	AreaLsaAnnounceMessages         uint32   `protobuf:"varint,4,opt,name=area_lsa_announce_messages,json=areaLsaAnnounceMessages,proto3" json:"area_lsa_announce_messages,omitempty"`
	AreaLsaFragmentAnnounceMessages uint32   `protobuf:"varint,5,opt,name=area_lsa_fragment_announce_messages,json=areaLsaFragmentAnnounceMessages,proto3" json:"area_lsa_fragment_announce_messages,omitempty"`
	AreaLsaDeleteMessages           uint32   `protobuf:"varint,6,opt,name=area_lsa_delete_messages,json=areaLsaDeleteMessages,proto3" json:"area_lsa_delete_messages,omitempty"`
	AreaLsaFragmentDeleteMessages   uint32   `protobuf:"varint,7,opt,name=area_lsa_fragment_delete_messages,json=areaLsaFragmentDeleteMessages,proto3" json:"area_lsa_fragment_delete_messages,omitempty"`
	AreaTunnelAnnounceMessages      uint32   `protobuf:"varint,8,opt,name=area_tunnel_announce_messages,json=areaTunnelAnnounceMessages,proto3" json:"area_tunnel_announce_messages,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *TeShIgpAreaStats) Reset()         { *m = TeShIgpAreaStats{} }
func (m *TeShIgpAreaStats) String() string { return proto.CompactTextString(m) }
func (*TeShIgpAreaStats) ProtoMessage()    {}
func (*TeShIgpAreaStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d3a45ebb1af4017, []int{1}
}

func (m *TeShIgpAreaStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeShIgpAreaStats.Unmarshal(m, b)
}
func (m *TeShIgpAreaStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeShIgpAreaStats.Marshal(b, m, deterministic)
}
func (m *TeShIgpAreaStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeShIgpAreaStats.Merge(m, src)
}
func (m *TeShIgpAreaStats) XXX_Size() int {
	return xxx_messageInfo_TeShIgpAreaStats.Size(m)
}
func (m *TeShIgpAreaStats) XXX_DiscardUnknown() {
	xxx_messageInfo_TeShIgpAreaStats.DiscardUnknown(m)
}

var xxx_messageInfo_TeShIgpAreaStats proto.InternalMessageInfo

func (m *TeShIgpAreaStats) GetAreaAdjacencyRequestMessages() uint32 {
	if m != nil {
		return m.AreaAdjacencyRequestMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaAdjacencyAnnounceMessages() uint32 {
	if m != nil {
		return m.AreaAdjacencyAnnounceMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaFloodMessages() uint32 {
	if m != nil {
		return m.AreaFloodMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaLsaAnnounceMessages() uint32 {
	if m != nil {
		return m.AreaLsaAnnounceMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaLsaFragmentAnnounceMessages() uint32 {
	if m != nil {
		return m.AreaLsaFragmentAnnounceMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaLsaDeleteMessages() uint32 {
	if m != nil {
		return m.AreaLsaDeleteMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaLsaFragmentDeleteMessages() uint32 {
	if m != nil {
		return m.AreaLsaFragmentDeleteMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaTunnelAnnounceMessages() uint32 {
	if m != nil {
		return m.AreaTunnelAnnounceMessages
	}
	return 0
}

type TeShIgpAreaActive struct {
	InterfacesCount             uint32            `protobuf:"varint,1,opt,name=interfaces_count,json=interfacesCount,proto3" json:"interfaces_count,omitempty"`
	LinkIdtReceived             bool              `protobuf:"varint,2,opt,name=link_idt_received,json=linkIdtReceived,proto3" json:"link_idt_received,omitempty"`
	TopologyIdtReceived         bool              `protobuf:"varint,3,opt,name=topology_idt_received,json=topologyIdtReceived,proto3" json:"topology_idt_received,omitempty"`
	SrStrict                    bool              `protobuf:"varint,4,opt,name=sr_strict,json=srStrict,proto3" json:"sr_strict,omitempty"`
	P2PHeadsCount               uint32            `protobuf:"varint,5,opt,name=p2p_heads_count,json=p2pHeadsCount,proto3" json:"p2p_heads_count,omitempty"`
	P2PAutorouteAnnouncedCount  uint32            `protobuf:"varint,6,opt,name=p2p_autoroute_announced_count,json=p2pAutorouteAnnouncedCount,proto3" json:"p2p_autoroute_announced_count,omitempty"`
	P2PForwardingAdjacencyCount uint32            `protobuf:"varint,7,opt,name=p2p_forwarding_adjacency_count,json=p2pForwardingAdjacencyCount,proto3" json:"p2p_forwarding_adjacency_count,omitempty"`
	P2MpDestinationCount        uint32            `protobuf:"varint,8,opt,name=p2mp_destination_count,json=p2mpDestinationCount,proto3" json:"p2mp_destination_count,omitempty"`
	TunnelLooseHops             uint32            `protobuf:"varint,9,opt,name=tunnel_loose_hops,json=tunnelLooseHops,proto3" json:"tunnel_loose_hops,omitempty"`
	AreaStatistics              *TeShIgpAreaStats `protobuf:"bytes,10,opt,name=area_statistics,json=areaStatistics,proto3" json:"area_statistics,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}          `json:"-"`
	XXX_unrecognized            []byte            `json:"-"`
	XXX_sizecache               int32             `json:"-"`
}

func (m *TeShIgpAreaActive) Reset()         { *m = TeShIgpAreaActive{} }
func (m *TeShIgpAreaActive) String() string { return proto.CompactTextString(m) }
func (*TeShIgpAreaActive) ProtoMessage()    {}
func (*TeShIgpAreaActive) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d3a45ebb1af4017, []int{2}
}

func (m *TeShIgpAreaActive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeShIgpAreaActive.Unmarshal(m, b)
}
func (m *TeShIgpAreaActive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeShIgpAreaActive.Marshal(b, m, deterministic)
}
func (m *TeShIgpAreaActive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeShIgpAreaActive.Merge(m, src)
}
func (m *TeShIgpAreaActive) XXX_Size() int {
	return xxx_messageInfo_TeShIgpAreaActive.Size(m)
}
func (m *TeShIgpAreaActive) XXX_DiscardUnknown() {
	xxx_messageInfo_TeShIgpAreaActive.DiscardUnknown(m)
}

var xxx_messageInfo_TeShIgpAreaActive proto.InternalMessageInfo

func (m *TeShIgpAreaActive) GetInterfacesCount() uint32 {
	if m != nil {
		return m.InterfacesCount
	}
	return 0
}

func (m *TeShIgpAreaActive) GetLinkIdtReceived() bool {
	if m != nil {
		return m.LinkIdtReceived
	}
	return false
}

func (m *TeShIgpAreaActive) GetTopologyIdtReceived() bool {
	if m != nil {
		return m.TopologyIdtReceived
	}
	return false
}

func (m *TeShIgpAreaActive) GetSrStrict() bool {
	if m != nil {
		return m.SrStrict
	}
	return false
}

func (m *TeShIgpAreaActive) GetP2PHeadsCount() uint32 {
	if m != nil {
		return m.P2PHeadsCount
	}
	return 0
}

func (m *TeShIgpAreaActive) GetP2PAutorouteAnnouncedCount() uint32 {
	if m != nil {
		return m.P2PAutorouteAnnouncedCount
	}
	return 0
}

func (m *TeShIgpAreaActive) GetP2PForwardingAdjacencyCount() uint32 {
	if m != nil {
		return m.P2PForwardingAdjacencyCount
	}
	return 0
}

func (m *TeShIgpAreaActive) GetP2MpDestinationCount() uint32 {
	if m != nil {
		return m.P2MpDestinationCount
	}
	return 0
}

func (m *TeShIgpAreaActive) GetTunnelLooseHops() uint32 {
	if m != nil {
		return m.TunnelLooseHops
	}
	return 0
}

func (m *TeShIgpAreaActive) GetAreaStatistics() *TeShIgpAreaStats {
	if m != nil {
		return m.AreaStatistics
	}
	return nil
}

type TeShIgpArea struct {
	AreaIndex            uint32             `protobuf:"varint,1,opt,name=area_index,json=areaIndex,proto3" json:"area_index,omitempty"`
	AreaNumber           uint32             `protobuf:"varint,2,opt,name=area_number,json=areaNumber,proto3" json:"area_number,omitempty"`
	AreaFormat           string             `protobuf:"bytes,3,opt,name=area_format,json=areaFormat,proto3" json:"area_format,omitempty"`
	IsConfigReady        bool               `protobuf:"varint,4,opt,name=is_config_ready,json=isConfigReady,proto3" json:"is_config_ready,omitempty"`
	ActiveData           *TeShIgpAreaActive `protobuf:"bytes,5,opt,name=active_data,json=activeData,proto3" json:"active_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TeShIgpArea) Reset()         { *m = TeShIgpArea{} }
func (m *TeShIgpArea) String() string { return proto.CompactTextString(m) }
func (*TeShIgpArea) ProtoMessage()    {}
func (*TeShIgpArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d3a45ebb1af4017, []int{3}
}

func (m *TeShIgpArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeShIgpArea.Unmarshal(m, b)
}
func (m *TeShIgpArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeShIgpArea.Marshal(b, m, deterministic)
}
func (m *TeShIgpArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeShIgpArea.Merge(m, src)
}
func (m *TeShIgpArea) XXX_Size() int {
	return xxx_messageInfo_TeShIgpArea.Size(m)
}
func (m *TeShIgpArea) XXX_DiscardUnknown() {
	xxx_messageInfo_TeShIgpArea.DiscardUnknown(m)
}

var xxx_messageInfo_TeShIgpArea proto.InternalMessageInfo

func (m *TeShIgpArea) GetAreaIndex() uint32 {
	if m != nil {
		return m.AreaIndex
	}
	return 0
}

func (m *TeShIgpArea) GetAreaNumber() uint32 {
	if m != nil {
		return m.AreaNumber
	}
	return 0
}

func (m *TeShIgpArea) GetAreaFormat() string {
	if m != nil {
		return m.AreaFormat
	}
	return ""
}

func (m *TeShIgpArea) GetIsConfigReady() bool {
	if m != nil {
		return m.IsConfigReady
	}
	return false
}

func (m *TeShIgpArea) GetActiveData() *TeShIgpAreaActive {
	if m != nil {
		return m.ActiveData
	}
	return nil
}

type TeShIgp struct {
	IgpType                string         `protobuf:"bytes,50,opt,name=igp_type,json=igpType,proto3" json:"igp_type,omitempty"`
	InstanceName           string         `protobuf:"bytes,51,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	IgpSystemId            string         `protobuf:"bytes,52,opt,name=igp_system_id,json=igpSystemId,proto3" json:"igp_system_id,omitempty"`
	ConfiguredTeRouterId   string         `protobuf:"bytes,53,opt,name=configured_te_router_id,json=configuredTeRouterId,proto3" json:"configured_te_router_id,omitempty"`
	GlobalRouterId         string         `protobuf:"bytes,54,opt,name=global_router_id,json=globalRouterId,proto3" json:"global_router_id,omitempty"`
	SecondaryRouterId      []string       `protobuf:"bytes,55,rep,name=secondary_router_id,json=secondaryRouterId,proto3" json:"secondary_router_id,omitempty"`
	GloballRouterIdOptical string         `protobuf:"bytes,56,opt,name=globall_router_id_optical,json=globallRouterIdOptical,proto3" json:"globall_router_id_optical,omitempty"`
	InUseTeRouterId        string         `protobuf:"bytes,57,opt,name=in_use_te_router_id,json=inUseTeRouterId,proto3" json:"in_use_te_router_id,omitempty"`
	IsConnectionUp         bool           `protobuf:"varint,58,opt,name=is_connection_up,json=isConnectionUp,proto3" json:"is_connection_up,omitempty"`
	ConnectionUpCount      uint32         `protobuf:"varint,59,opt,name=connection_up_count,json=connectionUpCount,proto3" json:"connection_up_count,omitempty"`
	ConnectionDownCount    uint32         `protobuf:"varint,60,opt,name=connection_down_count,json=connectionDownCount,proto3" json:"connection_down_count,omitempty"`
	Area                   []*TeShIgpArea `protobuf:"bytes,61,rep,name=area,proto3" json:"area,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}       `json:"-"`
	XXX_unrecognized       []byte         `json:"-"`
	XXX_sizecache          int32          `json:"-"`
}

func (m *TeShIgp) Reset()         { *m = TeShIgp{} }
func (m *TeShIgp) String() string { return proto.CompactTextString(m) }
func (*TeShIgp) ProtoMessage()    {}
func (*TeShIgp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d3a45ebb1af4017, []int{4}
}

func (m *TeShIgp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeShIgp.Unmarshal(m, b)
}
func (m *TeShIgp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeShIgp.Marshal(b, m, deterministic)
}
func (m *TeShIgp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeShIgp.Merge(m, src)
}
func (m *TeShIgp) XXX_Size() int {
	return xxx_messageInfo_TeShIgp.Size(m)
}
func (m *TeShIgp) XXX_DiscardUnknown() {
	xxx_messageInfo_TeShIgp.DiscardUnknown(m)
}

var xxx_messageInfo_TeShIgp proto.InternalMessageInfo

func (m *TeShIgp) GetIgpType() string {
	if m != nil {
		return m.IgpType
	}
	return ""
}

func (m *TeShIgp) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *TeShIgp) GetIgpSystemId() string {
	if m != nil {
		return m.IgpSystemId
	}
	return ""
}

func (m *TeShIgp) GetConfiguredTeRouterId() string {
	if m != nil {
		return m.ConfiguredTeRouterId
	}
	return ""
}

func (m *TeShIgp) GetGlobalRouterId() string {
	if m != nil {
		return m.GlobalRouterId
	}
	return ""
}

func (m *TeShIgp) GetSecondaryRouterId() []string {
	if m != nil {
		return m.SecondaryRouterId
	}
	return nil
}

func (m *TeShIgp) GetGloballRouterIdOptical() string {
	if m != nil {
		return m.GloballRouterIdOptical
	}
	return ""
}

func (m *TeShIgp) GetInUseTeRouterId() string {
	if m != nil {
		return m.InUseTeRouterId
	}
	return ""
}

func (m *TeShIgp) GetIsConnectionUp() bool {
	if m != nil {
		return m.IsConnectionUp
	}
	return false
}

func (m *TeShIgp) GetConnectionUpCount() uint32 {
	if m != nil {
		return m.ConnectionUpCount
	}
	return 0
}

func (m *TeShIgp) GetConnectionDownCount() uint32 {
	if m != nil {
		return m.ConnectionDownCount
	}
	return 0
}

func (m *TeShIgp) GetArea() []*TeShIgpArea {
	if m != nil {
		return m.Area
	}
	return nil
}

func init() {
	proto.RegisterType((*TeShIgp_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.igp_area_briefs.igp_area_brief.te_sh_igp_KEYS")
	proto.RegisterType((*TeShIgpAreaStats)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.igp_area_briefs.igp_area_brief.te_sh_igp_area_stats")
	proto.RegisterType((*TeShIgpAreaActive)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.igp_area_briefs.igp_area_brief.te_sh_igp_area_active")
	proto.RegisterType((*TeShIgpArea)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.igp_area_briefs.igp_area_brief.te_sh_igp_area")
	proto.RegisterType((*TeShIgp)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.igp_area_briefs.igp_area_brief.te_sh_igp")
}

func init() { proto.RegisterFile("te_sh_igp.proto", fileDescriptor_9d3a45ebb1af4017) }

var fileDescriptor_9d3a45ebb1af4017 = []byte{
	// 948 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdf, 0x72, 0x1b, 0x35,
	0x14, 0xc6, 0x27, 0x7f, 0x68, 0x6d, 0x99, 0xc4, 0xcd, 0x26, 0x69, 0xb7, 0x29, 0x21, 0xc1, 0x9d,
	0x61, 0x0c, 0xc3, 0xf8, 0x62, 0xdb, 0x52, 0x4a, 0x61, 0x06, 0x4f, 0xd2, 0x90, 0x0c, 0x69, 0x61,
	0x36, 0x29, 0x0c, 0x57, 0x1a, 0x65, 0x75, 0xbc, 0x11, 0xac, 0x25, 0x21, 0x69, 0x9b, 0xf8, 0x8a,
	0x87, 0xe0, 0x21, 0x78, 0x1f, 0xee, 0x79, 0x17, 0x46, 0x47, 0xeb, 0xf5, 0xda, 0xc9, 0x1d, 0xbd,
	0x4a, 0x74, 0xbe, 0xdf, 0xf9, 0xb4, 0xd2, 0x39, 0x47, 0x09, 0xe9, 0x3a, 0xa0, 0xf6, 0x92, 0x8a,
	0x5c, 0x0f, 0xb4, 0x51, 0x4e, 0x45, 0xdf, 0x65, 0xc2, 0x66, 0x8a, 0x0a, 0x65, 0xe9, 0xb5, 0xa1,
	0x63, 0x5d, 0x58, 0xea, 0x80, 0x2a, 0x0d, 0x66, 0x50, 0x2d, 0x06, 0x22, 0xd7, 0x94, 0x19, 0x60,
	0xf4, 0xc2, 0x08, 0x18, 0xd9, 0x85, 0x75, 0xef, 0x35, 0x59, 0xaf, 0x4d, 0xe9, 0x0f, 0xaf, 0x7e,
	0x3d, 0x8b, 0x76, 0x48, 0x0b, 0xcd, 0x33, 0x55, 0xc4, 0x4b, 0xfb, 0x4b, 0xfd, 0x76, 0x5a, 0xaf,
	0xa3, 0x3d, 0xd2, 0xd1, 0x46, 0x65, 0x60, 0x2d, 0x75, 0x2c, 0x8f, 0x97, 0x51, 0x26, 0x55, 0xe8,
	0x9c, 0xe5, 0xbd, 0xbf, 0x57, 0xc9, 0xd6, 0xcc, 0x0f, 0xf7, 0xb1, 0x8e, 0x39, 0x1b, 0xbd, 0x22,
	0x7b, 0xb8, 0x62, 0xfc, 0x37, 0x96, 0x81, 0xcc, 0x26, 0xd4, 0xc0, 0x1f, 0x25, 0x58, 0x47, 0xc7,
	0x60, 0x2d, 0xcb, 0xc1, 0xe2, 0x66, 0x6b, 0xe9, 0x47, 0x1e, 0x1b, 0x4e, 0xa9, 0x34, 0x40, 0xaf,
	0x2b, 0x26, 0xfa, 0x9e, 0xec, 0x2f, 0xd8, 0x30, 0x29, 0x55, 0x29, 0x33, 0x98, 0xf9, 0x2c, 0xa3,
	0xcf, 0xee, 0x9c, 0xcf, 0xb0, 0xa2, 0x6a, 0xa3, 0x01, 0xd9, 0x44, 0xa3, 0x51, 0xa1, 0x14, 0x9f,
	0xe5, 0xae, 0x60, 0xee, 0x86, 0x97, 0x8e, 0xbc, 0x52, 0xf3, 0x2f, 0xc9, 0x0e, 0xf2, 0x85, 0x65,
	0xb7, 0x6c, 0xb9, 0x8a, 0x69, 0x0f, 0x3c, 0x71, 0x6a, 0xd9, 0x8d, 0xcd, 0x4e, 0xc9, 0xe3, 0x3a,
	0x79, 0x64, 0x58, 0x3e, 0x06, 0xe9, 0x6e, 0x71, 0xf9, 0x00, 0x5d, 0xf6, 0x2a, 0x97, 0xa3, 0x0a,
	0xbc, 0xe1, 0xf6, 0x9c, 0xc4, 0xb5, 0x1b, 0x87, 0x02, 0x5c, 0xc3, 0xe2, 0x0e, 0x5a, 0x6c, 0x57,
	0x16, 0x87, 0xa8, 0xd6, 0x89, 0xc7, 0xe4, 0x93, 0x9b, 0x9f, 0xb1, 0xe8, 0x70, 0x77, 0x76, 0x7b,
	0x8d, 0x8f, 0x58, 0x70, 0x1a, 0x12, 0x04, 0xa8, 0x2b, 0xa5, 0x84, 0xe2, 0x96, 0xa3, 0xb4, 0xd0,
	0x05, 0xaf, 0xec, 0x1c, 0x99, 0xc5, 0x53, 0xf4, 0xfe, 0x59, 0x25, 0xdb, 0x0b, 0x9d, 0xc2, 0x32,
	0x27, 0xde, 0x41, 0xf4, 0x19, 0xb9, 0x27, 0xa4, 0x03, 0x33, 0x62, 0x19, 0x58, 0x9a, 0xa9, 0x52,
	0xba, 0xaa, 0x37, 0xba, 0xb3, 0xf8, 0x81, 0x0f, 0x47, 0x9f, 0x93, 0x8d, 0x42, 0xc8, 0xdf, 0xa9,
	0xe0, 0x8e, 0x1a, 0xc8, 0x40, 0xbc, 0x03, 0x8e, 0xf5, 0x6f, 0xa5, 0x5d, 0x2f, 0x9c, 0x70, 0x97,
	0x56, 0xe1, 0x28, 0x21, 0xdb, 0x4e, 0x69, 0x55, 0xa8, 0x7c, 0x32, 0xcf, 0xaf, 0x20, 0xbf, 0x39,
	0x15, 0x9b, 0x39, 0x8f, 0x48, 0xdb, 0x1a, 0x6a, 0x9d, 0x11, 0x99, 0xc3, 0x22, 0xb7, 0xd2, 0x96,
	0x35, 0x67, 0xb8, 0x8e, 0x3e, 0x25, 0x5d, 0x9d, 0x68, 0x7a, 0x09, 0x8c, 0x4f, 0x3f, 0x33, 0x54,
	0x70, 0x4d, 0x27, 0xfa, 0xd8, 0x47, 0xc3, 0x47, 0x0e, 0xc9, 0xae, 0xe7, 0x58, 0xe9, 0x94, 0x51,
	0xa5, 0x83, 0xfa, 0xba, 0x78, 0x95, 0x15, 0x8a, 0xb6, 0xa3, 0x13, 0x3d, 0x9c, 0x32, 0xd3, 0xeb,
	0xe2, 0xc1, 0xe2, 0x80, 0x7c, 0xec, 0x2d, 0x46, 0xca, 0x5c, 0x31, 0xc3, 0x85, 0xcc, 0x1b, 0x03,
	0x10, 0x3c, 0x42, 0xd9, 0x1e, 0xe9, 0x44, 0x1f, 0xd5, 0x50, 0xdd, 0xfd, 0xc1, 0xe4, 0x29, 0xb9,
	0xaf, 0x93, 0xb1, 0xa6, 0x1c, 0xac, 0x13, 0x92, 0x39, 0xa1, 0x64, 0x95, 0x1c, 0xaa, 0xb5, 0xe5,
	0xd5, 0xc3, 0x99, 0x58, 0x5f, 0x71, 0x55, 0xe5, 0x42, 0x29, 0x0b, 0xf4, 0x52, 0x69, 0x1b, 0xb7,
	0x43, 0x39, 0x82, 0x70, 0xea, 0xe3, 0xc7, 0x4a, 0xdb, 0xe8, 0x4f, 0xd2, 0xad, 0x47, 0x5e, 0x58,
	0x27, 0x32, 0x1b, 0x93, 0xfd, 0xa5, 0x7e, 0x27, 0xf9, 0x79, 0xf0, 0x7f, 0x1f, 0xaa, 0xc1, 0x6d,
	0xaf, 0x4a, 0xba, 0xee, 0x7f, 0x3f, 0xab, 0x77, 0xeb, 0xfd, 0xb5, 0xdc, 0x7c, 0xce, 0xbc, 0x18,
	0xed, 0x12, 0x82, 0x09, 0x42, 0x72, 0xb8, 0xae, 0xfa, 0xa8, 0xed, 0x23, 0x27, 0x3e, 0xe0, 0x5f,
	0x34, 0x94, 0x65, 0x39, 0xbe, 0x00, 0x53, 0xbd, 0x1d, 0x98, 0xf1, 0x06, 0x23, 0x35, 0x30, 0x52,
	0x66, 0xcc, 0x1c, 0x36, 0x4b, 0x3b, 0x00, 0x47, 0x18, 0xf1, 0x6d, 0x20, 0x7c, 0xfd, 0xe5, 0x48,
	0xe4, 0xd4, 0x00, 0xe3, 0x93, 0xaa, 0x53, 0xd6, 0x84, 0x3d, 0xc0, 0x68, 0xea, 0x83, 0xd1, 0x35,
	0xe9, 0x84, 0x06, 0xa7, 0x9c, 0x39, 0x86, 0xad, 0xd2, 0x49, 0x7e, 0x79, 0xef, 0x17, 0x13, 0xf6,
	0x48, 0x49, 0xf8, 0x79, 0xc8, 0x1c, 0xeb, 0xfd, 0xbb, 0x4a, 0xda, 0x35, 0x15, 0x3d, 0x24, 0x2d,
	0x0f, 0xbb, 0x89, 0x86, 0x38, 0xc1, 0xd3, 0xdc, 0x15, 0xb9, 0x3e, 0x9f, 0x68, 0x88, 0x1e, 0x93,
	0x35, 0x21, 0xad, 0x63, 0x7e, 0x94, 0x25, 0x1b, 0x43, 0xfc, 0x04, 0xf5, 0x0f, 0xa7, 0xc1, 0x37,
	0x6c, 0x0c, 0x51, 0x8f, 0xac, 0xf9, 0x7c, 0x3b, 0xb1, 0x0e, 0xc6, 0x54, 0xf0, 0xf8, 0x29, 0x42,
	0x1d, 0x91, 0xeb, 0x33, 0x8c, 0x9d, 0xf0, 0xe8, 0x19, 0x79, 0x10, 0x2e, 0xa4, 0x34, 0xc0, 0xfd,
	0x81, 0xb0, 0xa9, 0x8d, 0xa7, 0x9f, 0x21, 0xbd, 0x35, 0x93, 0xcf, 0x21, 0x45, 0xf1, 0x84, 0x47,
	0x7d, 0x72, 0x2f, 0x2f, 0xd4, 0x05, 0x2b, 0x1a, 0xfc, 0x97, 0xc8, 0xaf, 0x87, 0x78, 0x4d, 0x0e,
	0xc8, 0xa6, 0x85, 0x4c, 0x49, 0xce, 0xcc, 0xa4, 0x01, 0x3f, 0xdf, 0x5f, 0xe9, 0xb7, 0xd3, 0x8d,
	0x5a, 0xaa, 0xf9, 0x17, 0xe4, 0x61, 0x70, 0x68, 0x58, 0x53, 0xa5, 0x9d, 0xc8, 0x58, 0x11, 0x7f,
	0x85, 0x5b, 0xdc, 0xaf, 0x80, 0x69, 0xce, 0x8f, 0x41, 0x8d, 0xbe, 0x20, 0x9b, 0x42, 0xd2, 0xd2,
	0xc2, 0xfc, 0x39, 0x5e, 0x60, 0x52, 0x57, 0xc8, 0xb7, 0x16, 0xe6, 0x8f, 0x10, 0xba, 0x41, 0x42,
	0x86, 0x13, 0x56, 0xea, 0xf8, 0x6b, 0x6c, 0x87, 0x75, 0x6c, 0x87, 0x2a, 0xfc, 0x56, 0xfb, 0x23,
	0xcc, 0x61, 0xd5, 0x2c, 0xbe, 0x0c, 0x7f, 0x81, 0xb2, 0x06, 0x1a, 0x06, 0x31, 0x21, 0xdb, 0x0d,
	0x9e, 0xab, 0xab, 0xe9, 0xf4, 0x7e, 0x83, 0x19, 0x0d, 0xb3, 0x43, 0x75, 0x55, 0x0d, 0x2f, 0x27,
	0xab, 0xbe, 0x29, 0xe2, 0x6f, 0xf7, 0x57, 0xfa, 0x9d, 0xe4, 0xa7, 0xf7, 0xdd, 0x6c, 0x29, 0xba,
	0x5f, 0xdc, 0xc1, 0xff, 0x0f, 0x9e, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x07, 0xb7, 0xc3,
	0x9f, 0x08, 0x00, 0x00,
}
