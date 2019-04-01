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
// source: ipv6_rib_edm_route.proto

package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_local_lspv_protocol_routes_protocol_route

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

type Ipv6RibEdmRoute_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName       string   `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6RibEdmRoute_KEYS) Reset()         { *m = Ipv6RibEdmRoute_KEYS{} }
func (m *Ipv6RibEdmRoute_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmRoute_KEYS) ProtoMessage()    {}
func (*Ipv6RibEdmRoute_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_97e3b1af6daa47d0, []int{0}
}

func (m *Ipv6RibEdmRoute_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmRoute_KEYS.Unmarshal(m, b)
}
func (m *Ipv6RibEdmRoute_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmRoute_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6RibEdmRoute_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmRoute_KEYS.Merge(m, src)
}
func (m *Ipv6RibEdmRoute_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmRoute_KEYS.Size(m)
}
func (m *Ipv6RibEdmRoute_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmRoute_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmRoute_KEYS proto.InternalMessageInfo

func (m *Ipv6RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ipv6RibEdmPathItem struct {
	Address                      string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	InformationSource            string   `protobuf:"bytes,2,opt,name=information_source,json=informationSource,proto3" json:"information_source,omitempty"`
	V6Nexthop                    string   `protobuf:"bytes,3,opt,name=v6_nexthop,json=v6Nexthop,proto3" json:"v6_nexthop,omitempty"`
	V6InformationSource          string   `protobuf:"bytes,4,opt,name=v6_information_source,json=v6InformationSource,proto3" json:"v6_information_source,omitempty"`
	InterfaceName                string   `protobuf:"bytes,5,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Metric                       uint32   `protobuf:"varint,6,opt,name=metric,proto3" json:"metric,omitempty"`
	LoadMetric                   uint32   `protobuf:"varint,7,opt,name=load_metric,json=loadMetric,proto3" json:"load_metric,omitempty"`
	Flags64                      uint64   `protobuf:"varint,8,opt,name=flags64,proto3" json:"flags64,omitempty"`
	Flags                        uint32   `protobuf:"varint,9,opt,name=flags,proto3" json:"flags,omitempty"`
	PrivateFlags                 uint32   `protobuf:"varint,10,opt,name=private_flags,json=privateFlags,proto3" json:"private_flags,omitempty"`
	Looped                       bool     `protobuf:"varint,11,opt,name=looped,proto3" json:"looped,omitempty"`
	NextHopTableId               uint32   `protobuf:"varint,12,opt,name=next_hop_table_id,json=nextHopTableId,proto3" json:"next_hop_table_id,omitempty"`
	NextHopVrfName               string   `protobuf:"bytes,13,opt,name=next_hop_vrf_name,json=nextHopVrfName,proto3" json:"next_hop_vrf_name,omitempty"`
	NextHopTableName             string   `protobuf:"bytes,14,opt,name=next_hop_table_name,json=nextHopTableName,proto3" json:"next_hop_table_name,omitempty"`
	NextHopAfi                   uint32   `protobuf:"varint,15,opt,name=next_hop_afi,json=nextHopAfi,proto3" json:"next_hop_afi,omitempty"`
	NextHopSafi                  uint32   `protobuf:"varint,16,opt,name=next_hop_safi,json=nextHopSafi,proto3" json:"next_hop_safi,omitempty"`
	RouteLabel                   uint32   `protobuf:"varint,17,opt,name=route_label,json=routeLabel,proto3" json:"route_label,omitempty"`
	TunnelId                     uint32   `protobuf:"varint,18,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	Pathid                       uint32   `protobuf:"varint,19,opt,name=pathid,proto3" json:"pathid,omitempty"`
	BackupPathid                 uint32   `protobuf:"varint,20,opt,name=backup_pathid,json=backupPathid,proto3" json:"backup_pathid,omitempty"`
	RefCntOfBackup               uint32   `protobuf:"varint,21,opt,name=ref_cnt_of_backup,json=refCntOfBackup,proto3" json:"ref_cnt_of_backup,omitempty"`
	NumberOfExtendedCommunities  uint32   `protobuf:"varint,22,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities,proto3" json:"number_of_extended_communities,omitempty"`
	MvpnPresent                  bool     `protobuf:"varint,23,opt,name=mvpn_present,json=mvpnPresent,proto3" json:"mvpn_present,omitempty"`
	PathRtPresent                bool     `protobuf:"varint,24,opt,name=path_rt_present,json=pathRtPresent,proto3" json:"path_rt_present,omitempty"`
	VrfImportRtPresent           bool     `protobuf:"varint,25,opt,name=vrf_import_rt_present,json=vrfImportRtPresent,proto3" json:"vrf_import_rt_present,omitempty"`
	SourceAsrtPresent            bool     `protobuf:"varint,26,opt,name=source_asrt_present,json=sourceAsrtPresent,proto3" json:"source_asrt_present,omitempty"`
	SourceRdPresent              bool     `protobuf:"varint,27,opt,name=source_rd_present,json=sourceRdPresent,proto3" json:"source_rd_present,omitempty"`
	SegmentedNexthopPresent      bool     `protobuf:"varint,28,opt,name=segmented_nexthop_present,json=segmentedNexthopPresent,proto3" json:"segmented_nexthop_present,omitempty"`
	NextHopId                    uint32   `protobuf:"varint,29,opt,name=next_hop_id,json=nextHopId,proto3" json:"next_hop_id,omitempty"`
	NextHopIdRefcount            uint32   `protobuf:"varint,30,opt,name=next_hop_id_refcount,json=nextHopIdRefcount,proto3" json:"next_hop_id_refcount,omitempty"`
	OspfAreaId                   string   `protobuf:"bytes,31,opt,name=ospf_area_id,json=ospfAreaId,proto3" json:"ospf_area_id,omitempty"`
	RemoteBackupAddr             []string `protobuf:"bytes,32,rep,name=remote_backup_addr,json=remoteBackupAddr,proto3" json:"remote_backup_addr,omitempty"`
	HasLabelstk                  bool     `protobuf:"varint,33,opt,name=has_labelstk,json=hasLabelstk,proto3" json:"has_labelstk,omitempty"`
	NumLabels                    uint32   `protobuf:"varint,34,opt,name=num_labels,json=numLabels,proto3" json:"num_labels,omitempty"`
	Labelstk                     []uint32 `protobuf:"varint,35,rep,packed,name=labelstk,proto3" json:"labelstk,omitempty"`
	BindingLabel                 uint32   `protobuf:"varint,36,opt,name=binding_label,json=bindingLabel,proto3" json:"binding_label,omitempty"`
	NhidFeid                     uint64   `protobuf:"varint,37,opt,name=nhid_feid,json=nhidFeid,proto3" json:"nhid_feid,omitempty"`
	MplsFeid                     uint64   `protobuf:"varint,38,opt,name=mpls_feid,json=mplsFeid,proto3" json:"mpls_feid,omitempty"`
	HasVxlanNetworkId            bool     `protobuf:"varint,39,opt,name=has_vxlan_network_id,json=hasVxlanNetworkId,proto3" json:"has_vxlan_network_id,omitempty"`
	VxlanNetworkId               uint32   `protobuf:"varint,40,opt,name=vxlan_network_id,json=vxlanNetworkId,proto3" json:"vxlan_network_id,omitempty"`
	HasXcid                      bool     `protobuf:"varint,41,opt,name=has_xcid,json=hasXcid,proto3" json:"has_xcid,omitempty"`
	Xcid                         uint32   `protobuf:"varint,42,opt,name=xcid,proto3" json:"xcid,omitempty"`
	HasSpanDiagInterface         bool     `protobuf:"varint,43,opt,name=has_span_diag_interface,json=hasSpanDiagInterface,proto3" json:"has_span_diag_interface,omitempty"`
	SpanDiagInterface            string   `protobuf:"bytes,44,opt,name=span_diag_interface,json=spanDiagInterface,proto3" json:"span_diag_interface,omitempty"`
	HasSubscriberParentInterface bool     `protobuf:"varint,45,opt,name=has_subscriber_parent_interface,json=hasSubscriberParentInterface,proto3" json:"has_subscriber_parent_interface,omitempty"`
	SubscriberParentInterface    string   `protobuf:"bytes,46,opt,name=subscriber_parent_interface,json=subscriberParentInterface,proto3" json:"subscriber_parent_interface,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *Ipv6RibEdmPathItem) Reset()         { *m = Ipv6RibEdmPathItem{} }
func (m *Ipv6RibEdmPathItem) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmPathItem) ProtoMessage()    {}
func (*Ipv6RibEdmPathItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_97e3b1af6daa47d0, []int{1}
}

func (m *Ipv6RibEdmPathItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmPathItem.Unmarshal(m, b)
}
func (m *Ipv6RibEdmPathItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmPathItem.Marshal(b, m, deterministic)
}
func (m *Ipv6RibEdmPathItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmPathItem.Merge(m, src)
}
func (m *Ipv6RibEdmPathItem) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmPathItem.Size(m)
}
func (m *Ipv6RibEdmPathItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmPathItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmPathItem proto.InternalMessageInfo

func (m *Ipv6RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetV6InformationSource() string {
	if m != nil {
		return m.V6InformationSource
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetPathRtPresent() bool {
	if m != nil {
		return m.PathRtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetVrfImportRtPresent() bool {
	if m != nil {
		return m.VrfImportRtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSourceAsrtPresent() bool {
	if m != nil {
		return m.SourceAsrtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSourceRdPresent() bool {
	if m != nil {
		return m.SourceRdPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetRemoteBackupAddr() []string {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv6RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv6RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetMplsFeid() uint64 {
	if m != nil {
		return m.MplsFeid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetHasVxlanNetworkId() bool {
	if m != nil {
		return m.HasVxlanNetworkId
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetVxlanNetworkId() uint32 {
	if m != nil {
		return m.VxlanNetworkId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetHasXcid() bool {
	if m != nil {
		return m.HasXcid
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetXcid() uint32 {
	if m != nil {
		return m.Xcid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetHasSpanDiagInterface() bool {
	if m != nil {
		return m.HasSpanDiagInterface
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSpanDiagInterface() string {
	if m != nil {
		return m.SpanDiagInterface
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetHasSubscriberParentInterface() bool {
	if m != nil {
		return m.HasSubscriberParentInterface
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSubscriberParentInterface() string {
	if m != nil {
		return m.SubscriberParentInterface
	}
	return ""
}

type Ipv6RibEdmPathEntry struct {
	Ipv6RibEdmPath       []*Ipv6RibEdmPathItem `protobuf:"bytes,47,rep,name=ipv6_rib_edm_path,json=ipv6RibEdmPath,proto3" json:"ipv6_rib_edm_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Ipv6RibEdmPathEntry) Reset()         { *m = Ipv6RibEdmPathEntry{} }
func (m *Ipv6RibEdmPathEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmPathEntry) ProtoMessage()    {}
func (*Ipv6RibEdmPathEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_97e3b1af6daa47d0, []int{2}
}

func (m *Ipv6RibEdmPathEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmPathEntry.Unmarshal(m, b)
}
func (m *Ipv6RibEdmPathEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmPathEntry.Marshal(b, m, deterministic)
}
func (m *Ipv6RibEdmPathEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmPathEntry.Merge(m, src)
}
func (m *Ipv6RibEdmPathEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmPathEntry.Size(m)
}
func (m *Ipv6RibEdmPathEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmPathEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmPathEntry proto.InternalMessageInfo

func (m *Ipv6RibEdmPathEntry) GetIpv6RibEdmPath() []*Ipv6RibEdmPathItem {
	if m != nil {
		return m.Ipv6RibEdmPath
	}
	return nil
}

type Ipv6RibEdmRoute struct {
	Prefix               string               `protobuf:"bytes,50,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLengthXr       uint32               `protobuf:"varint,51,opt,name=prefix_length_xr,json=prefixLengthXr,proto3" json:"prefix_length_xr,omitempty"`
	RouteVersion         uint32               `protobuf:"varint,52,opt,name=route_version,json=routeVersion,proto3" json:"route_version,omitempty"`
	ProtocolId           uint32               `protobuf:"varint,53,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"`
	ProtocolName         string               `protobuf:"bytes,54,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	Instance             string               `protobuf:"bytes,55,opt,name=instance,proto3" json:"instance,omitempty"`
	ClientId             uint32               `protobuf:"varint,56,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	RouteType            uint32               `protobuf:"varint,57,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	Priority             uint32               `protobuf:"varint,58,opt,name=priority,proto3" json:"priority,omitempty"`
	SvdType              uint32               `protobuf:"varint,59,opt,name=svd_type,json=svdType,proto3" json:"svd_type,omitempty"`
	Flags                uint32               `protobuf:"varint,60,opt,name=flags,proto3" json:"flags,omitempty"`
	ExtendedFlags        uint64               `protobuf:"varint,61,opt,name=extended_flags,json=extendedFlags,proto3" json:"extended_flags,omitempty"`
	Tag                  uint32               `protobuf:"varint,62,opt,name=tag,proto3" json:"tag,omitempty"`
	Distance             uint32               `protobuf:"varint,63,opt,name=distance,proto3" json:"distance,omitempty"`
	DiversionDistance    uint32               `protobuf:"varint,64,opt,name=diversion_distance,json=diversionDistance,proto3" json:"diversion_distance,omitempty"`
	Metric               uint32               `protobuf:"varint,65,opt,name=metric,proto3" json:"metric,omitempty"`
	PathsCount           uint32               `protobuf:"varint,66,opt,name=paths_count,json=pathsCount,proto3" json:"paths_count,omitempty"`
	AttributeIdentity    uint32               `protobuf:"varint,67,opt,name=attribute_identity,json=attributeIdentity,proto3" json:"attribute_identity,omitempty"`
	TrafficIndex         uint32               `protobuf:"varint,68,opt,name=traffic_index,json=trafficIndex,proto3" json:"traffic_index,omitempty"`
	RoutePrecedence      uint32               `protobuf:"varint,69,opt,name=route_precedence,json=routePrecedence,proto3" json:"route_precedence,omitempty"`
	QosGroup             uint32               `protobuf:"varint,70,opt,name=qos_group,json=qosGroup,proto3" json:"qos_group,omitempty"`
	FlowTag              uint32               `protobuf:"varint,71,opt,name=flow_tag,json=flowTag,proto3" json:"flow_tag,omitempty"`
	FwdClass             uint32               `protobuf:"varint,72,opt,name=fwd_class,json=fwdClass,proto3" json:"fwd_class,omitempty"`
	PicCount             uint32               `protobuf:"varint,73,opt,name=pic_count,json=picCount,proto3" json:"pic_count,omitempty"`
	Active               bool                 `protobuf:"varint,74,opt,name=active,proto3" json:"active,omitempty"`
	Diversion            bool                 `protobuf:"varint,75,opt,name=diversion,proto3" json:"diversion,omitempty"`
	DiversionProtoName   string               `protobuf:"bytes,76,opt,name=diversion_proto_name,json=diversionProtoName,proto3" json:"diversion_proto_name,omitempty"`
	RouteAge             uint32               `protobuf:"varint,77,opt,name=route_age,json=routeAge,proto3" json:"route_age,omitempty"`
	RouteLabel           uint32               `protobuf:"varint,78,opt,name=route_label,json=routeLabel,proto3" json:"route_label,omitempty"`
	Version              uint32               `protobuf:"varint,79,opt,name=version,proto3" json:"version,omitempty"`
	TblVersion           uint64               `protobuf:"varint,80,opt,name=tbl_version,json=tblVersion,proto3" json:"tbl_version,omitempty"`
	RouteModifyTime      uint64               `protobuf:"varint,81,opt,name=route_modify_time,json=routeModifyTime,proto3" json:"route_modify_time,omitempty"`
	RoutePath            *Ipv6RibEdmPathEntry `protobuf:"bytes,82,opt,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Ipv6RibEdmRoute) Reset()         { *m = Ipv6RibEdmRoute{} }
func (m *Ipv6RibEdmRoute) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmRoute) ProtoMessage()    {}
func (*Ipv6RibEdmRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_97e3b1af6daa47d0, []int{3}
}

func (m *Ipv6RibEdmRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmRoute.Unmarshal(m, b)
}
func (m *Ipv6RibEdmRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmRoute.Marshal(b, m, deterministic)
}
func (m *Ipv6RibEdmRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmRoute.Merge(m, src)
}
func (m *Ipv6RibEdmRoute) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmRoute.Size(m)
}
func (m *Ipv6RibEdmRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmRoute.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmRoute proto.InternalMessageInfo

func (m *Ipv6RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetPrefixLengthXr() uint32 {
	if m != nil {
		return m.PrefixLengthXr
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv6RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv6RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRoutePath() *Ipv6RibEdmPathEntry {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv6RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.protocol_routes.protocol_route.ipv6_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv6RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.protocol_routes.protocol_route.ipv6_rib_edm_path_item")
	proto.RegisterType((*Ipv6RibEdmPathEntry)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.protocol_routes.protocol_route.ipv6_rib_edm_path_entry")
	proto.RegisterType((*Ipv6RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.protocol_routes.protocol_route.ipv6_rib_edm_route")
}

func init() { proto.RegisterFile("ipv6_rib_edm_route.proto", fileDescriptor_97e3b1af6daa47d0) }

var fileDescriptor_97e3b1af6daa47d0 = []byte{
	// 1590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x72, 0x1b, 0xb7,
	0x15, 0x1e, 0x46, 0x8e, 0x44, 0x41, 0x3f, 0xb6, 0xd6, 0xb2, 0x04, 0x59, 0x8e, 0xcd, 0x28, 0x75,
	0x2a, 0xbb, 0x31, 0xdd, 0x3a, 0x89, 0xda, 0xa6, 0x6d, 0x5a, 0x45, 0x96, 0x13, 0x36, 0xfe, 0x51,
	0x69, 0x8f, 0x27, 0xbd, 0xc2, 0x80, 0x0b, 0x2c, 0x89, 0xf1, 0x2e, 0x76, 0x03, 0x80, 0x6b, 0xe9,
	0xae, 0x33, 0x7d, 0x94, 0x3c, 0x46, 0xfb, 0x12, 0xbd, 0xef, 0x7d, 0x5f, 0xa3, 0x73, 0xce, 0xc1,
	0x2e, 0x29, 0xc9, 0xc9, 0x7d, 0x6e, 0x24, 0xe2, 0xfb, 0xbe, 0x73, 0x16, 0x38, 0x3f, 0xc0, 0x61,
	0xdc, 0x54, 0xf5, 0x81, 0x70, 0x66, 0x24, 0xb4, 0x2a, 0x84, 0x2b, 0xa7, 0x41, 0xf7, 0x2b, 0x57,
	0x86, 0x32, 0xf9, 0xa1, 0x93, 0x1a, 0x9f, 0x96, 0xc2, 0x94, 0x5e, 0x9c, 0x3a, 0x61, 0x2a, 0x54,
	0xa1, 0xbc, 0xac, 0xb4, 0xeb, 0x37, 0x86, 0xfd, 0xda, 0x65, 0x1e, 0xfe, 0xf4, 0x65, 0xe6, 0xfb,
	0x32, 0xeb, 0x7b, 0xf8, 0xef, 0x65, 0xd6, 0x8f, 0x26, 0xe8, 0x54, 0x04, 0x39, 0xca, 0xb5, 0xb0,
	0xb2, 0xd0, 0xfe, 0xc7, 0x08, 0xfa, 0x70, 0x5a, 0xe6, 0xfd, 0xbc, 0x4c, 0x65, 0xde, 0xcf, 0x7d,
	0x55, 0xb7, 0x18, 0xa9, 0xfd, 0x85, 0xf5, 0xde, 0x7f, 0x3a, 0x6c, 0xfb, 0xf2, 0x11, 0xc4, 0xb7,
	0xc7, 0x7f, 0x7f, 0x99, 0xec, 0xb0, 0x6e, 0xed, 0x32, 0x74, 0xcd, 0x3b, 0xbd, 0xce, 0xfe, 0xf2,
	0x70, 0xa9, 0x76, 0xd9, 0x73, 0x59, 0xe8, 0x64, 0x9b, 0x2d, 0xc9, 0xc8, 0xbc, 0x87, 0xcc, 0xa2,
	0x24, 0x62, 0x87, 0x75, 0x7d, 0xc3, 0x2c, 0x90, 0x8d, 0x8f, 0xd4, 0x3e, 0xbb, 0x76, 0x71, 0xc7,
	0xfc, 0x0a, 0x4a, 0xd6, 0x11, 0x7f, 0x05, 0x30, 0x2a, 0x39, 0x5b, 0x92, 0x4a, 0x39, 0xed, 0x3d,
	0x7f, 0x9f, 0x7c, 0xc4, 0x65, 0xf2, 0x11, 0x5b, 0xab, 0x9c, 0xce, 0xcc, 0xa9, 0xc8, 0xb5, 0x1d,
	0x87, 0x09, 0x5f, 0xec, 0x75, 0xf6, 0xd7, 0x86, 0xab, 0x04, 0x3e, 0x45, 0x6c, 0xef, 0x87, 0x75,
	0xb6, 0x75, 0xee, 0x4c, 0x95, 0x0c, 0x13, 0x61, 0x82, 0x2e, 0xe6, 0x3d, 0x77, 0xce, 0x7b, 0x7e,
	0xc0, 0x12, 0x63, 0xb3, 0xd2, 0x15, 0x32, 0x98, 0xd2, 0x0a, 0x5f, 0x4e, 0x5d, 0xda, 0x1c, 0x6e,
	0x63, 0x8e, 0x79, 0x89, 0x44, 0xf2, 0x01, 0x63, 0xf5, 0x81, 0xb0, 0xfa, 0x34, 0x4c, 0xca, 0x2a,
	0x9e, 0x74, 0xb9, 0x3e, 0x78, 0x4e, 0x40, 0xf2, 0x88, 0xdd, 0xa8, 0x0f, 0xc4, 0x3b, 0x1c, 0xd2,
	0x81, 0xaf, 0xd7, 0x07, 0x83, 0x4b, 0x2e, 0xef, 0xb2, 0x75, 0x63, 0x83, 0x76, 0x99, 0x4c, 0x63,
	0x74, 0xe8, 0xf0, 0x6b, 0x2d, 0x8a, 0xc1, 0xd9, 0x62, 0x8b, 0x85, 0x0e, 0xce, 0xa4, 0xf1, 0xec,
	0x71, 0x95, 0xdc, 0x61, 0x2b, 0x79, 0x29, 0x95, 0x88, 0xe4, 0x12, 0x92, 0x0c, 0xa0, 0x67, 0x24,
	0xe0, 0x6c, 0x29, 0xcb, 0xe5, 0xd8, 0x1f, 0x7c, 0xc6, 0xbb, 0xbd, 0xce, 0xfe, 0x95, 0x61, 0xb3,
	0x4c, 0x36, 0xd9, 0xfb, 0xf8, 0x93, 0x2f, 0xa3, 0x11, 0x2d, 0x28, 0xd6, 0xa6, 0x96, 0x41, 0x0b,
	0x62, 0x59, 0x13, 0x6b, 0x04, 0x9f, 0xa0, 0x68, 0x8b, 0x2d, 0xe6, 0x65, 0x59, 0x69, 0xc5, 0x57,
	0x7a, 0x9d, 0xfd, 0xee, 0x30, 0xae, 0x92, 0x7b, 0x6c, 0x03, 0x82, 0x23, 0x26, 0x65, 0x15, 0xf3,
	0x6d, 0x14, 0x5f, 0x45, 0x07, 0xeb, 0x40, 0x7c, 0x53, 0x56, 0x98, 0xef, 0xc1, 0x79, 0x69, 0x5b,
	0x6f, 0x6b, 0x54, 0x18, 0x51, 0xfa, 0x3a, 0x96, 0xdd, 0x03, 0x76, 0xfd, 0x82, 0x57, 0x14, 0xaf,
	0xa3, 0xf8, 0xda, 0xbc, 0x5f, 0x94, 0xf7, 0xd8, 0x6a, 0x2b, 0x97, 0x99, 0xe1, 0x57, 0x29, 0x26,
	0x51, 0x77, 0x98, 0x99, 0x64, 0x8f, 0xad, 0xb5, 0x0a, 0x0f, 0x92, 0x6b, 0x28, 0x59, 0x89, 0x92,
	0x97, 0x32, 0x33, 0x10, 0x58, 0xaa, 0xdb, 0x5c, 0x8e, 0x74, 0xce, 0x37, 0xc8, 0x09, 0x42, 0x4f,
	0x01, 0x49, 0x76, 0xd9, 0x72, 0x98, 0x5a, 0xab, 0x73, 0x38, 0x63, 0x82, 0x74, 0x97, 0x80, 0x81,
	0x82, 0x00, 0x41, 0xf9, 0x19, 0xc5, 0xaf, 0x53, 0xba, 0x68, 0x05, 0xd1, 0x1d, 0xc9, 0xf4, 0xcd,
	0xb4, 0x12, 0x91, 0xde, 0xa4, 0xe8, 0x12, 0x78, 0x42, 0xa2, 0x7b, 0x6c, 0xc3, 0xe9, 0x4c, 0xa4,
	0x36, 0x88, 0x32, 0x13, 0x44, 0xf1, 0x1b, 0x14, 0x45, 0xa7, 0xb3, 0x23, 0x1b, 0x5e, 0x64, 0x5f,
	0x21, 0x9a, 0x1c, 0xb1, 0xdb, 0x76, 0x5a, 0x8c, 0xb4, 0x03, 0xa5, 0x3e, 0x0d, 0xda, 0x2a, 0xad,
	0x44, 0x5a, 0x16, 0xc5, 0xd4, 0x9a, 0x60, 0xb4, 0xe7, 0x5b, 0x68, 0xb7, 0x4b, 0xaa, 0x17, 0xd9,
	0x71, 0xd4, 0x1c, 0xcd, 0x24, 0xc9, 0x87, 0x6c, 0xb5, 0xa8, 0x2b, 0x2b, 0x2a, 0xa7, 0xbd, 0xb6,
	0x81, 0x6f, 0x63, 0x4e, 0x57, 0x00, 0x3b, 0x21, 0x28, 0xf9, 0x98, 0x5d, 0xc5, 0x76, 0x72, 0xa1,
	0x55, 0x71, 0x54, 0xad, 0x01, 0x3c, 0x0c, 0x8d, 0xee, 0x37, 0xec, 0x06, 0x24, 0xd3, 0x14, 0x55,
	0xe9, 0xc2, 0xbc, 0x7a, 0x07, 0xd5, 0x49, 0xed, 0xb2, 0x01, 0x72, 0x33, 0x93, 0x3e, 0xbb, 0x4e,
	0x5d, 0x22, 0xa4, 0x9f, 0x33, 0xb8, 0x89, 0x06, 0x1b, 0x44, 0x1d, 0x7a, 0xd7, 0xea, 0xef, 0xb3,
	0x08, 0x0a, 0xa7, 0x5a, 0xf5, 0x2e, 0xaa, 0xaf, 0x12, 0x31, 0x54, 0x8d, 0xf6, 0x0b, 0xb6, 0xe3,
	0xf5, 0xb8, 0xd0, 0x36, 0x68, 0xd5, 0xb4, 0x6d, 0x6b, 0x73, 0x0b, 0x6d, 0xb6, 0x5b, 0x41, 0xec,
	0xe2, 0xc6, 0xf6, 0x36, 0x5b, 0x69, 0x8b, 0xc4, 0x28, 0xfe, 0x01, 0xc6, 0x71, 0x39, 0x96, 0xc8,
	0x40, 0x25, 0x0f, 0xd9, 0xe6, 0x1c, 0x2f, 0x9c, 0xce, 0xd2, 0x72, 0x6a, 0x03, 0xbf, 0x8d, 0xc2,
	0x8d, 0x56, 0x38, 0x8c, 0x04, 0xd4, 0x65, 0xe9, 0xab, 0x4c, 0x48, 0xa7, 0x25, 0x78, 0xbc, 0x83,
	0xf5, 0xcb, 0x00, 0x3b, 0x74, 0x5a, 0x0e, 0x54, 0xf2, 0x09, 0x4b, 0x9c, 0x2e, 0xca, 0xa0, 0x63,
	0xd2, 0x05, 0x5c, 0x53, 0xbc, 0xd7, 0x5b, 0x80, 0x3a, 0x27, 0x86, 0xf2, 0x7e, 0xa8, 0x94, 0x83,
	0xb4, 0x4d, 0xa4, 0xa7, 0xfa, 0xf4, 0xe1, 0x0d, 0xff, 0x90, 0xd2, 0x36, 0x91, 0xfe, 0x69, 0x84,
	0xe0, 0xbe, 0xb2, 0xd3, 0x22, 0x4a, 0xf8, 0x5e, 0x3c, 0xc2, 0xb4, 0x20, 0x41, 0x72, 0x93, 0x75,
	0x5b, 0xeb, 0x8f, 0x7a, 0x0b, 0x50, 0xc1, 0xcd, 0x1a, 0x2b, 0xd5, 0x58, 0x65, 0xec, 0x38, 0x76,
	0xc0, 0x2f, 0x62, 0xa5, 0x12, 0xd8, 0xf6, 0x80, 0x9d, 0x18, 0x25, 0x32, 0x6d, 0x14, 0xbf, 0x8b,
	0xd7, 0x4b, 0x17, 0x80, 0x27, 0xda, 0x28, 0x20, 0x8b, 0x2a, 0xf7, 0x44, 0x7e, 0x4c, 0x24, 0x00,
	0x48, 0x3e, 0x64, 0x9b, 0xb0, 0xf9, 0xfa, 0x34, 0x97, 0x56, 0x58, 0x1d, 0xde, 0x96, 0xee, 0x0d,
	0x04, 0xe5, 0x97, 0x94, 0xf6, 0x89, 0xf4, 0xaf, 0x81, 0x7a, 0x4e, 0xcc, 0x40, 0xc1, 0x3b, 0x72,
	0x49, 0xbc, 0x4f, 0x3d, 0x51, 0x9f, 0x57, 0xee, 0xb0, 0x2e, 0xb8, 0x3e, 0x4d, 0x8d, 0xe2, 0xf7,
	0xd0, 0xdd, 0xd2, 0x44, 0xfa, 0xef, 0x52, 0xa3, 0x92, 0x84, 0x5d, 0x41, 0xf8, 0x3e, 0x1a, 0xe2,
	0xef, 0xe4, 0x73, 0xb6, 0x0d, 0x72, 0x5f, 0x49, 0x2b, 0x94, 0x91, 0x63, 0xd1, 0x5e, 0xbc, 0xfc,
	0x57, 0x68, 0x0d, 0x1b, 0x7d, 0x59, 0x49, 0xfb, 0xd8, 0xc8, 0xf1, 0xa0, 0xe1, 0xb0, 0x6c, 0xdf,
	0x61, 0xf2, 0x09, 0x3d, 0x1d, 0xfe, 0x92, 0xfe, 0x98, 0xdd, 0xc1, 0xcf, 0x4c, 0x47, 0x3e, 0x75,
	0x06, 0x3a, 0xb6, 0x92, 0x4e, 0xdb, 0x30, 0x67, 0xfb, 0x00, 0x3f, 0x77, 0x0b, 0x3e, 0xd7, 0xaa,
	0x4e, 0x50, 0x34, 0x73, 0xf3, 0x25, 0xdb, 0xfd, 0x29, 0x17, 0x7d, 0xfc, 0xfc, 0x8e, 0xff, 0x31,
	0xfb, 0xbd, 0x7f, 0xbe, 0x77, 0xe1, 0xe5, 0xc7, 0xb6, 0xd6, 0x36, 0xb8, 0xb3, 0xe4, 0x7f, 0x1d,
	0xb6, 0x71, 0x89, 0xe3, 0x0f, 0x7b, 0x0b, 0xfb, 0x2b, 0x8f, 0xfe, 0xd5, 0xe9, 0xff, 0x0c, 0x06,
	0x9b, 0xfe, 0xbb, 0x07, 0x80, 0xe1, 0x3a, 0xe0, 0x43, 0x33, 0x3a, 0x56, 0x05, 0xdc, 0xb1, 0x7b,
	0xff, 0x60, 0x2c, 0xb9, 0x3c, 0xff, 0xe0, 0xad, 0x8d, 0x23, 0x05, 0x7f, 0x44, 0xe3, 0x0d, 0xad,
	0xa0, 0xf6, 0xce, 0xcd, 0x1f, 0xe2, 0xd4, 0xf1, 0x4f, 0xa9, 0xf6, 0xe6, 0x47, 0x90, 0xef, 0x1c,
	0x74, 0x0d, 0x1d, 0xa3, 0xd6, 0xce, 0x9b, 0xd2, 0xf2, 0xcf, 0xa8, 0x6b, 0x10, 0x7c, 0x4d, 0x18,
	0x3c, 0x2d, 0xed, 0xb6, 0x8d, 0xe2, 0x9f, 0xd3, 0xd3, 0xd2, 0x40, 0x03, 0x45, 0x6f, 0x70, 0x14,
	0xe0, 0x53, 0x77, 0x80, 0xdb, 0x59, 0x6d, 0x40, 0x7c, 0xe6, 0x6e, 0xb2, 0xae, 0xb1, 0x3e, 0x48,
	0x9b, 0x6a, 0xfe, 0x5b, 0xe4, 0xdb, 0x35, 0xb4, 0x5e, 0x9a, 0x1b, 0x2c, 0x0d, 0xc5, 0x7f, 0x47,
	0x6f, 0x13, 0x01, 0x03, 0x05, 0x97, 0x42, 0x0c, 0xf5, 0x59, 0xa5, 0xf9, 0xef, 0xe9, 0x52, 0xa0,
	0x59, 0xec, 0xac, 0x42, 0xbf, 0x95, 0x33, 0xa5, 0x33, 0xe1, 0x8c, 0x7f, 0x41, 0xa6, 0xcd, 0x1a,
	0xe7, 0xbc, 0x5a, 0x91, 0xe1, 0x1f, 0x90, 0x5b, 0xf2, 0xb5, 0x42, 0xb3, 0x76, 0x9a, 0xf8, 0xe3,
	0xfc, 0x34, 0x71, 0x97, 0xad, 0xb7, 0xaf, 0x12, 0xd1, 0x7f, 0xc2, 0x8b, 0x60, 0xad, 0x41, 0x69,
	0x9e, 0xb8, 0xc6, 0x16, 0x82, 0x1c, 0xf3, 0x2f, 0xd1, 0x14, 0x7e, 0xc2, 0x2e, 0x94, 0x89, 0xa7,
	0xfb, 0x33, 0xed, 0xa2, 0x59, 0xc3, 0xd0, 0xa6, 0x4c, 0x0c, 0xb0, 0x68, 0x55, 0x7f, 0xa1, 0x7b,
	0xb7, 0x65, 0x1e, 0x37, 0xf2, 0xd9, 0xe8, 0x74, 0x78, 0x71, 0x74, 0x82, 0x0a, 0xf1, 0x82, 0xee,
	0xed, 0xaf, 0x62, 0x1a, 0x00, 0x3a, 0xc2, 0x0b, 0xfb, 0x01, 0x4b, 0x64, 0x08, 0xce, 0x8c, 0x20,
	0x58, 0x46, 0x69, 0x1b, 0x20, 0x26, 0x47, 0xf4, 0x9d, 0x96, 0x19, 0x44, 0x02, 0xb2, 0x16, 0x9c,
	0xcc, 0x32, 0x93, 0x0a, 0x63, 0x95, 0x3e, 0xe5, 0x8f, 0x29, 0xf7, 0x11, 0x1c, 0x00, 0x96, 0xdc,
	0x6b, 0xc6, 0xe1, 0xca, 0xe9, 0x54, 0x2b, 0x0d, 0x3b, 0x3f, 0x46, 0xdd, 0x55, 0xc4, 0x4f, 0x5a,
	0x18, 0x92, 0xf8, 0x7d, 0xe9, 0xc5, 0xd8, 0x95, 0xd3, 0x8a, 0x3f, 0xa1, 0x18, 0x7c, 0x5f, 0xfa,
	0xaf, 0x61, 0x0d, 0x99, 0xc8, 0xf2, 0xf2, 0xad, 0x80, 0xb0, 0x7d, 0x4d, 0x99, 0x80, 0xf5, 0x2b,
	0x39, 0x06, 0xbb, 0xec, 0xad, 0x12, 0x69, 0x2e, 0xbd, 0xe7, 0xdf, 0x90, 0x5d, 0xf6, 0x56, 0x1d,
	0xc1, 0x1a, 0xc8, 0xca, 0xa4, 0xf1, 0xc8, 0x83, 0x98, 0x5e, 0x93, 0xd2, 0x81, 0xb7, 0xd8, 0xa2,
	0x4c, 0x83, 0xa9, 0x35, 0xff, 0x2b, 0x8d, 0x75, 0xb4, 0x4a, 0x6e, 0xb1, 0xe5, 0x36, 0xac, 0xfc,
	0x5b, 0xa4, 0x66, 0x40, 0xf2, 0x6b, 0xb6, 0x39, 0x4b, 0x07, 0x96, 0x28, 0x15, 0xed, 0x53, 0x2c,
	0xca, 0x59, 0xaa, 0x4e, 0x80, 0xc2, 0xd2, 0xdd, 0x65, 0x54, 0x6f, 0x42, 0x8e, 0x35, 0x7f, 0x46,
	0x9b, 0x40, 0xe0, 0x70, 0xac, 0x2f, 0x0e, 0x5e, 0xcf, 0x2f, 0x0d, 0x5e, 0x9c, 0x2d, 0x35, 0x7b,
	0x79, 0x41, 0x27, 0xaf, 0x67, 0x8d, 0x15, 0x46, 0x79, 0xdb, 0x7b, 0x27, 0x58, 0x6a, 0x2c, 0x8c,
	0xf2, 0xa6, 0xf3, 0xee, 0xb3, 0x0d, 0xf2, 0x5d, 0x94, 0xca, 0x64, 0x67, 0x22, 0x98, 0x42, 0xf3,
	0xbf, 0xa1, 0x8c, 0xc2, 0xff, 0x0c, 0xf1, 0x57, 0xa6, 0xd0, 0xc9, 0x7f, 0x3b, 0x4d, 0x9f, 0xe0,
	0x35, 0x38, 0xec, 0x75, 0xf6, 0x57, 0x1e, 0xfd, 0xfb, 0xe7, 0x7a, 0x0d, 0xe2, 0x0d, 0x1f, 0xdb,
	0x1c, 0xae, 0xc0, 0xd1, 0x22, 0x9a, 0x7c, 0xfa, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0xe4,
	0x20, 0x01, 0xcb, 0x0e, 0x00, 0x00,
}
