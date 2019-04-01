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
// source: ipv4_rib_edm_route.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_connected_non_as_protocol_routes_protocol_route

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

type Ipv4RibEdmRoute_KEYS struct {
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

func (m *Ipv4RibEdmRoute_KEYS) Reset()         { *m = Ipv4RibEdmRoute_KEYS{} }
func (m *Ipv4RibEdmRoute_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute_KEYS) ProtoMessage()    {}
func (*Ipv4RibEdmRoute_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_44c1b2f42127243a, []int{0}
}

func (m *Ipv4RibEdmRoute_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Unmarshal(m, b)
}
func (m *Ipv4RibEdmRoute_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmRoute_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Merge(m, src)
}
func (m *Ipv4RibEdmRoute_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Size(m)
}
func (m *Ipv4RibEdmRoute_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmRoute_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmRoute_KEYS proto.InternalMessageInfo

func (m *Ipv4RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ipv4RibEdmPathItem struct {
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

func (m *Ipv4RibEdmPathItem) Reset()         { *m = Ipv4RibEdmPathItem{} }
func (m *Ipv4RibEdmPathItem) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPathItem) ProtoMessage()    {}
func (*Ipv4RibEdmPathItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_44c1b2f42127243a, []int{1}
}

func (m *Ipv4RibEdmPathItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmPathItem.Unmarshal(m, b)
}
func (m *Ipv4RibEdmPathItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmPathItem.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmPathItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmPathItem.Merge(m, src)
}
func (m *Ipv4RibEdmPathItem) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmPathItem.Size(m)
}
func (m *Ipv4RibEdmPathItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmPathItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmPathItem proto.InternalMessageInfo

func (m *Ipv4RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetV6InformationSource() string {
	if m != nil {
		return m.V6InformationSource
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetPathRtPresent() bool {
	if m != nil {
		return m.PathRtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetVrfImportRtPresent() bool {
	if m != nil {
		return m.VrfImportRtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourceAsrtPresent() bool {
	if m != nil {
		return m.SourceAsrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourceRdPresent() bool {
	if m != nil {
		return m.SourceRdPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetRemoteBackupAddr() []string {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMplsFeid() uint64 {
	if m != nil {
		return m.MplsFeid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetHasVxlanNetworkId() bool {
	if m != nil {
		return m.HasVxlanNetworkId
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetVxlanNetworkId() uint32 {
	if m != nil {
		return m.VxlanNetworkId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetHasXcid() bool {
	if m != nil {
		return m.HasXcid
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetXcid() uint32 {
	if m != nil {
		return m.Xcid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetHasSpanDiagInterface() bool {
	if m != nil {
		return m.HasSpanDiagInterface
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSpanDiagInterface() string {
	if m != nil {
		return m.SpanDiagInterface
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetHasSubscriberParentInterface() bool {
	if m != nil {
		return m.HasSubscriberParentInterface
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSubscriberParentInterface() string {
	if m != nil {
		return m.SubscriberParentInterface
	}
	return ""
}

type Ipv4RibEdmPathEntry struct {
	Ipv4RibEdmPath       []*Ipv4RibEdmPathItem `protobuf:"bytes,47,rep,name=ipv4_rib_edm_path,json=ipv4RibEdmPath,proto3" json:"ipv4_rib_edm_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Ipv4RibEdmPathEntry) Reset()         { *m = Ipv4RibEdmPathEntry{} }
func (m *Ipv4RibEdmPathEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPathEntry) ProtoMessage()    {}
func (*Ipv4RibEdmPathEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_44c1b2f42127243a, []int{2}
}

func (m *Ipv4RibEdmPathEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmPathEntry.Unmarshal(m, b)
}
func (m *Ipv4RibEdmPathEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmPathEntry.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmPathEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmPathEntry.Merge(m, src)
}
func (m *Ipv4RibEdmPathEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmPathEntry.Size(m)
}
func (m *Ipv4RibEdmPathEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmPathEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmPathEntry proto.InternalMessageInfo

func (m *Ipv4RibEdmPathEntry) GetIpv4RibEdmPath() []*Ipv4RibEdmPathItem {
	if m != nil {
		return m.Ipv4RibEdmPath
	}
	return nil
}

type Ipv4RibEdmRoute struct {
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
	RoutePath            *Ipv4RibEdmPathEntry `protobuf:"bytes,82,opt,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Ipv4RibEdmRoute) Reset()         { *m = Ipv4RibEdmRoute{} }
func (m *Ipv4RibEdmRoute) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute) ProtoMessage()    {}
func (*Ipv4RibEdmRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_44c1b2f42127243a, []int{3}
}

func (m *Ipv4RibEdmRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmRoute.Unmarshal(m, b)
}
func (m *Ipv4RibEdmRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmRoute.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmRoute.Merge(m, src)
}
func (m *Ipv4RibEdmRoute) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmRoute.Size(m)
}
func (m *Ipv4RibEdmRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmRoute.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmRoute proto.InternalMessageInfo

func (m *Ipv4RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetPrefixLengthXr() uint32 {
	if m != nil {
		return m.PrefixLengthXr
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePath() *Ipv4RibEdmPathEntry {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.connected.non_as.protocol_routes.protocol_route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.connected.non_as.protocol_routes.protocol_route.ipv4_rib_edm_path_item")
	proto.RegisterType((*Ipv4RibEdmPathEntry)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.connected.non_as.protocol_routes.protocol_route.ipv4_rib_edm_path_entry")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.connected.non_as.protocol_routes.protocol_route.ipv4_rib_edm_route")
}

func init() { proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor_44c1b2f42127243a) }

var fileDescriptor_44c1b2f42127243a = []byte{
	// 1599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x5d, 0x72, 0x1c, 0xb7,
	0x11, 0xae, 0x35, 0x65, 0x72, 0x09, 0xfe, 0x48, 0x1c, 0x51, 0x24, 0x28, 0xca, 0xd2, 0x9a, 0x8e,
	0x1c, 0x4a, 0xb1, 0x56, 0x89, 0x2c, 0x33, 0x89, 0x93, 0x38, 0xa1, 0x29, 0xca, 0xde, 0x58, 0x3f,
	0xcc, 0x4a, 0xa5, 0x72, 0x9e, 0x50, 0xd8, 0x01, 0x66, 0x17, 0xa5, 0x19, 0xcc, 0x18, 0xc0, 0x0e,
	0xc9, 0x13, 0xa4, 0x2a, 0xd7, 0xc8, 0x5b, 0x8e, 0x91, 0x9c, 0x20, 0x67, 0xc8, 0x7b, 0xce, 0x90,
	0xea, 0x6e, 0xcc, 0x70, 0x49, 0xca, 0x7e, 0xd7, 0x0b, 0xb9, 0xf8, 0xbe, 0xaf, 0x7b, 0x80, 0xfe,
	0x01, 0x9a, 0x71, 0x53, 0xd5, 0x8f, 0x85, 0x33, 0x23, 0xa1, 0x55, 0x21, 0x5c, 0x39, 0x0d, 0xba,
	0x5f, 0xb9, 0x32, 0x94, 0xc9, 0x3f, 0x3b, 0xa9, 0xf1, 0x69, 0x29, 0x4c, 0xe9, 0xc5, 0x89, 0x13,
	0xa6, 0x42, 0x15, 0xca, 0xcb, 0x4a, 0xbb, 0x3e, 0xac, 0x7c, 0x50, 0xa3, 0xd3, 0x7e, 0xed, 0x32,
	0x0f, 0x7f, 0xfa, 0x32, 0xf3, 0x7d, 0x99, 0xf5, 0x3d, 0xfc, 0xf7, 0x32, 0xeb, 0x47, 0x1b, 0xf4,
	0x2a, 0x82, 0x1c, 0xe5, 0x5a, 0x58, 0x59, 0x68, 0xff, 0x63, 0x04, 0x7d, 0x39, 0x2d, 0xf3, 0x7e,
	0x5a, 0x5a, 0xab, 0xd3, 0xa0, 0x55, 0xdf, 0x96, 0x56, 0x48, 0xdf, 0x32, 0x64, 0x73, 0x71, 0xbd,
	0xf3, 0x9f, 0x0e, 0xdb, 0xbc, 0x7c, 0x12, 0xf1, 0xdd, 0xe1, 0x5f, 0x5f, 0x25, 0x5b, 0xac, 0x5b,
	0xbb, 0x0c, 0x3f, 0xc0, 0x3b, 0xbd, 0xce, 0xee, 0xe2, 0x70, 0xa1, 0x76, 0xd9, 0x0b, 0x59, 0xe8,
	0x64, 0x93, 0x2d, 0xc8, 0xc8, 0x7c, 0x80, 0xcc, 0xbc, 0x24, 0x62, 0x8b, 0x75, 0x7d, 0xc3, 0xcc,
	0x91, 0x8d, 0x8f, 0xd4, 0x2e, 0xbb, 0x76, 0x71, 0xdf, 0xfc, 0x0a, 0x4a, 0x56, 0x11, 0x7f, 0x0d,
	0x30, 0x2a, 0x39, 0x5b, 0x90, 0x4a, 0x39, 0xed, 0x3d, 0xff, 0x90, 0x7c, 0xc4, 0x65, 0xf2, 0x09,
	0x5b, 0xa9, 0x9c, 0xce, 0xcc, 0x89, 0xc8, 0xb5, 0x1d, 0x87, 0x09, 0x9f, 0xef, 0x75, 0x76, 0x57,
	0x86, 0xcb, 0x04, 0x3e, 0x43, 0x6c, 0xe7, 0x1f, 0xab, 0x6c, 0xe3, 0xdc, 0x99, 0x2a, 0x19, 0x26,
	0xc2, 0x04, 0x5d, 0xcc, 0x7a, 0xee, 0x9c, 0xf7, 0xfc, 0x80, 0x25, 0xc6, 0x66, 0xa5, 0x2b, 0x64,
	0x30, 0xa5, 0x15, 0xbe, 0x9c, 0xba, 0xb4, 0x39, 0xdc, 0xda, 0x0c, 0xf3, 0x0a, 0x89, 0xe4, 0x23,
	0xc6, 0xea, 0x3d, 0x61, 0xf5, 0x49, 0x98, 0x94, 0x55, 0x3c, 0xe9, 0x62, 0xbd, 0xf7, 0x82, 0x80,
	0xe4, 0x11, 0xbb, 0x51, 0xef, 0x89, 0x77, 0x38, 0xa4, 0x03, 0x5f, 0xaf, 0xf7, 0x06, 0x97, 0x5c,
	0xde, 0x65, 0xab, 0xc6, 0x06, 0xed, 0x32, 0x99, 0xc6, 0xe8, 0xd0, 0xe1, 0x57, 0x5a, 0x14, 0x83,
	0xb3, 0xc1, 0xe6, 0x0b, 0x1d, 0x9c, 0x49, 0xe3, 0xd9, 0xe3, 0x2a, 0xb9, 0xc3, 0x96, 0xf2, 0x52,
	0x2a, 0x11, 0xc9, 0x05, 0x24, 0x19, 0x40, 0xcf, 0x49, 0xc0, 0xd9, 0x42, 0x96, 0xcb, 0xb1, 0xdf,
	0x7b, 0xcc, 0xbb, 0xbd, 0xce, 0xee, 0x95, 0x61, 0xb3, 0x4c, 0xd6, 0xd9, 0x87, 0xf8, 0x93, 0x2f,
	0xa2, 0x11, 0x2d, 0x28, 0xd6, 0xa6, 0x96, 0x41, 0x0b, 0x62, 0x59, 0x13, 0x6b, 0x04, 0x9f, 0xa2,
	0x68, 0x83, 0xcd, 0xe7, 0x65, 0x59, 0x69, 0xc5, 0x97, 0x7a, 0x9d, 0xdd, 0xee, 0x30, 0xae, 0x92,
	0x7b, 0x6c, 0x0d, 0x82, 0x23, 0x26, 0x65, 0x15, 0xf3, 0x6d, 0x14, 0x5f, 0x46, 0x07, 0xab, 0x40,
	0x7c, 0x5b, 0x56, 0x98, 0xef, 0xc1, 0x79, 0x69, 0x5b, 0x6f, 0x2b, 0x54, 0x18, 0x51, 0xfa, 0x26,
	0x96, 0xdd, 0x03, 0x76, 0xfd, 0x82, 0x57, 0x14, 0xaf, 0xa2, 0xf8, 0xda, 0xac, 0x5f, 0x94, 0xf7,
	0xd8, 0x72, 0x2b, 0x97, 0x99, 0xe1, 0x57, 0x29, 0x26, 0x51, 0xb7, 0x9f, 0x99, 0x64, 0x87, 0xad,
	0xb4, 0x0a, 0x0f, 0x92, 0x6b, 0x28, 0x59, 0x8a, 0x92, 0x57, 0x32, 0x33, 0x10, 0x58, 0xaa, 0xdb,
	0x5c, 0x8e, 0x74, 0xce, 0xd7, 0xc8, 0x09, 0x42, 0xcf, 0x00, 0x49, 0xb6, 0xd9, 0x62, 0x98, 0x5a,
	0xab, 0x73, 0x38, 0x63, 0x82, 0x74, 0x97, 0x80, 0x81, 0x82, 0x00, 0x41, 0xf9, 0x19, 0xc5, 0xaf,
	0x53, 0xba, 0x68, 0x05, 0xd1, 0x1d, 0xc9, 0xf4, 0xed, 0xb4, 0x12, 0x91, 0x5e, 0xa7, 0xe8, 0x12,
	0x78, 0x44, 0xa2, 0x7b, 0x6c, 0xcd, 0xe9, 0x4c, 0xa4, 0x36, 0x88, 0x32, 0x13, 0x44, 0xf1, 0x1b,
	0x14, 0x45, 0xa7, 0xb3, 0x03, 0x1b, 0x5e, 0x66, 0x5f, 0x23, 0x9a, 0x1c, 0xb0, 0xdb, 0x76, 0x5a,
	0x8c, 0xb4, 0x03, 0xa5, 0x3e, 0x09, 0xda, 0x2a, 0xad, 0x44, 0x5a, 0x16, 0xc5, 0xd4, 0x9a, 0x60,
	0xb4, 0xe7, 0x1b, 0x68, 0xb7, 0x4d, 0xaa, 0x97, 0xd9, 0x61, 0xd4, 0x1c, 0x9c, 0x49, 0x92, 0x8f,
	0xd9, 0x72, 0x51, 0x57, 0x56, 0x54, 0x4e, 0x7b, 0x6d, 0x03, 0xdf, 0xc4, 0x9c, 0x2e, 0x01, 0x76,
	0x44, 0x50, 0xf2, 0x29, 0xbb, 0x8a, 0xed, 0xe4, 0x42, 0xab, 0xe2, 0xa8, 0x5a, 0x01, 0x78, 0x18,
	0x1a, 0xdd, 0xaf, 0xd8, 0x0d, 0x48, 0xa6, 0x29, 0xaa, 0xd2, 0x85, 0x59, 0xf5, 0x16, 0xaa, 0x93,
	0xda, 0x65, 0x03, 0xe4, 0xce, 0x4c, 0xfa, 0xec, 0x3a, 0x75, 0x89, 0x90, 0x7e, 0xc6, 0xe0, 0x26,
	0x1a, 0xac, 0x11, 0xb5, 0xef, 0x5d, 0xab, 0xbf, 0xcf, 0x22, 0x28, 0x9c, 0x6a, 0xd5, 0xdb, 0xa8,
	0xbe, 0x4a, 0xc4, 0x50, 0x35, 0xda, 0x2f, 0xd9, 0x96, 0xd7, 0xe3, 0x42, 0xdb, 0xa0, 0x55, 0xd3,
	0xb6, 0xad, 0xcd, 0x2d, 0xb4, 0xd9, 0x6c, 0x05, 0xb1, 0x8b, 0x1b, 0xdb, 0xdb, 0x6c, 0xa9, 0x2d,
	0x12, 0xa3, 0xf8, 0x47, 0x18, 0xc7, 0xc5, 0x58, 0x22, 0x03, 0x95, 0x3c, 0x64, 0xeb, 0x33, 0xbc,
	0x70, 0x3a, 0x4b, 0xcb, 0xa9, 0x0d, 0xfc, 0x36, 0x0a, 0xd7, 0x5a, 0xe1, 0x30, 0x12, 0x50, 0x97,
	0xa5, 0xaf, 0x32, 0x21, 0x9d, 0x96, 0xe0, 0xf1, 0x0e, 0xd6, 0x2f, 0x03, 0x6c, 0xdf, 0x69, 0x39,
	0x50, 0xc9, 0x67, 0x2c, 0x71, 0xba, 0x28, 0x83, 0x8e, 0x49, 0x17, 0x70, 0x4d, 0xf1, 0x5e, 0x6f,
	0x0e, 0xea, 0x9c, 0x18, 0xca, 0xfb, 0xbe, 0x52, 0x0e, 0xd2, 0x36, 0x91, 0x9e, 0xea, 0xd3, 0x87,
	0xb7, 0xfc, 0x63, 0x4a, 0xdb, 0x44, 0xfa, 0x67, 0x11, 0x82, 0xfb, 0xca, 0x4e, 0x8b, 0x28, 0xe1,
	0x3b, 0xf1, 0x08, 0xd3, 0x82, 0x04, 0xc9, 0x4d, 0xd6, 0x6d, 0xad, 0x3f, 0xe9, 0xcd, 0x41, 0x05,
	0x37, 0x6b, 0xac, 0x54, 0x63, 0x95, 0xb1, 0xe3, 0xd8, 0x01, 0x3f, 0x8b, 0x95, 0x4a, 0x60, 0xdb,
	0x03, 0x76, 0x62, 0x94, 0xc8, 0xb4, 0x51, 0xfc, 0x2e, 0x5e, 0x2f, 0x5d, 0x00, 0x9e, 0x6a, 0xa3,
	0x80, 0x2c, 0xaa, 0xdc, 0x13, 0xf9, 0x29, 0x91, 0x00, 0x20, 0xf9, 0x90, 0xad, 0xc3, 0xe6, 0xeb,
	0x93, 0x5c, 0x5a, 0x61, 0x75, 0x38, 0x2e, 0xdd, 0x5b, 0x08, 0xca, 0xcf, 0x29, 0xed, 0x13, 0xe9,
	0xdf, 0x00, 0xf5, 0x82, 0x98, 0x81, 0x82, 0x77, 0xe4, 0x92, 0x78, 0x97, 0x7a, 0xa2, 0x3e, 0xaf,
	0xdc, 0x62, 0x5d, 0x70, 0x7d, 0x92, 0x1a, 0xc5, 0xef, 0xa1, 0xbb, 0x85, 0x89, 0xf4, 0xdf, 0xa7,
	0x46, 0x25, 0x09, 0xbb, 0x82, 0xf0, 0x7d, 0x34, 0xc4, 0xdf, 0xc9, 0x17, 0x6c, 0x13, 0xe4, 0xbe,
	0x92, 0x56, 0x28, 0x23, 0xc7, 0xa2, 0xbd, 0x78, 0xf9, 0x2f, 0xd0, 0x1a, 0x36, 0xfa, 0xaa, 0x92,
	0xf6, 0x89, 0x91, 0xe3, 0x41, 0xc3, 0x61, 0xd9, 0xbe, 0xc3, 0xe4, 0x33, 0x7a, 0x3a, 0xfc, 0x25,
	0xfd, 0x21, 0xbb, 0x83, 0x9f, 0x99, 0x8e, 0x7c, 0xea, 0x0c, 0x74, 0x6c, 0x25, 0x9d, 0xb6, 0x61,
	0xc6, 0xf6, 0x01, 0x7e, 0xee, 0x16, 0x7c, 0xae, 0x55, 0x1d, 0xa1, 0xe8, 0xcc, 0xcd, 0x57, 0x6c,
	0xfb, 0xa7, 0x5c, 0xf4, 0xf1, 0xf3, 0x5b, 0xfe, 0xc7, 0xec, 0x77, 0xfe, 0xfe, 0xc1, 0x85, 0x97,
	0x1f, 0xdb, 0x5a, 0xdb, 0xe0, 0x4e, 0x93, 0xff, 0x75, 0xd8, 0xda, 0x25, 0x8e, 0x3f, 0xec, 0xcd,
	0xed, 0x2e, 0x3d, 0xfa, 0x57, 0xa7, 0xff, 0xfe, 0xcc, 0x37, 0xfd, 0x77, 0xcf, 0x01, 0xc3, 0x55,
	0xc0, 0x87, 0x66, 0x74, 0xa8, 0x0a, 0xb8, 0x6a, 0x77, 0xfe, 0xc6, 0x58, 0x72, 0x79, 0x0c, 0xc2,
	0xcb, 0x1b, 0x27, 0x0b, 0xfe, 0x88, 0xa6, 0x1c, 0x5a, 0x41, 0x09, 0x9e, 0x1b, 0x43, 0xc4, 0x89,
	0xe3, 0x9f, 0x53, 0x09, 0xce, 0x4e, 0x22, 0xdf, 0x3b, 0x68, 0x1e, 0x3a, 0x4c, 0xad, 0x9d, 0x37,
	0xa5, 0xe5, 0x8f, 0xa9, 0x79, 0x10, 0x7c, 0x43, 0x18, 0xbc, 0x30, 0xed, 0xb6, 0x8d, 0xe2, 0x5f,
	0xd0, 0x0b, 0xd3, 0x40, 0x03, 0x45, 0x4f, 0x71, 0x14, 0xe0, 0x8b, 0xb7, 0x87, 0xdb, 0x59, 0x6e,
	0x40, 0x7c, 0xed, 0x6e, 0xb2, 0xae, 0xb1, 0x3e, 0x48, 0x9b, 0x6a, 0xfe, 0x6b, 0xe4, 0xdb, 0x35,
	0x74, 0x60, 0x9a, 0x1b, 0xac, 0x10, 0xc5, 0x7f, 0x43, 0x4f, 0x14, 0x01, 0x03, 0x05, 0x77, 0x43,
	0x0c, 0xf8, 0x69, 0xa5, 0xf9, 0x6f, 0xe9, 0x6e, 0xa0, 0x91, 0xec, 0xb4, 0x42, 0xbf, 0x95, 0x33,
	0xa5, 0x33, 0xe1, 0x94, 0x7f, 0x49, 0xa6, 0xcd, 0x1a, 0xc7, 0xbd, 0x5a, 0x91, 0xe1, 0xef, 0x90,
	0x5b, 0xf0, 0xb5, 0x42, 0xb3, 0x76, 0xa8, 0xf8, 0xfd, 0xec, 0x50, 0x71, 0x97, 0xad, 0xb6, 0x8f,
	0x13, 0xd1, 0x7f, 0xc0, 0xfb, 0x60, 0xa5, 0x41, 0x69, 0xac, 0xb8, 0xc6, 0xe6, 0x82, 0x1c, 0xf3,
	0xaf, 0xd0, 0x14, 0x7e, 0xc2, 0x2e, 0x94, 0x89, 0xa7, 0xfb, 0x23, 0xed, 0xa2, 0x59, 0xc3, 0xec,
	0xa6, 0x4c, 0x0c, 0xb0, 0x68, 0x55, 0x7f, 0xa2, 0xeb, 0xb7, 0x65, 0x9e, 0x34, 0xf2, 0xb3, 0x09,
	0x6a, 0xff, 0xe2, 0x04, 0x05, 0x15, 0xe2, 0x05, 0x5d, 0xdf, 0x5f, 0xc7, 0x34, 0x00, 0x74, 0x80,
	0xf7, 0xf6, 0x03, 0x96, 0xc8, 0x10, 0x9c, 0x19, 0x41, 0xb0, 0x8c, 0xd2, 0x36, 0x40, 0x4c, 0x0e,
	0xe8, 0x3b, 0x2d, 0x33, 0x88, 0x04, 0x64, 0x2d, 0x38, 0x99, 0x65, 0x26, 0x15, 0xc6, 0x2a, 0x7d,
	0xc2, 0x9f, 0x50, 0xee, 0x23, 0x38, 0x00, 0x2c, 0xb9, 0xd7, 0x4c, 0xc5, 0x95, 0xd3, 0xa9, 0x56,
	0x1a, 0x76, 0x7e, 0x88, 0xba, 0xab, 0x88, 0x1f, 0xb5, 0x30, 0x24, 0xf1, 0x87, 0xd2, 0x8b, 0xb1,
	0x2b, 0xa7, 0x15, 0x7f, 0x4a, 0x31, 0xf8, 0xa1, 0xf4, 0xdf, 0xc0, 0x1a, 0x32, 0x91, 0xe5, 0xe5,
	0xb1, 0x80, 0xb0, 0x7d, 0x43, 0x99, 0x80, 0xf5, 0x6b, 0x39, 0x06, 0xbb, 0xec, 0x58, 0x89, 0x34,
	0x97, 0xde, 0xf3, 0x6f, 0xc9, 0x2e, 0x3b, 0x56, 0x07, 0xb0, 0x06, 0xb2, 0x32, 0x69, 0x3c, 0xf2,
	0x20, 0xa6, 0xd7, 0xa4, 0x74, 0xe0, 0x0d, 0x36, 0x2f, 0xd3, 0x60, 0x6a, 0xcd, 0xff, 0x4c, 0xd3,
	0x1d, 0xad, 0x92, 0x5b, 0x6c, 0xb1, 0x0d, 0x2b, 0xff, 0x0e, 0xa9, 0x33, 0x20, 0xf9, 0x25, 0x5b,
	0x3f, 0x4b, 0x07, 0x96, 0x28, 0x15, 0xed, 0x33, 0x2c, 0xca, 0xb3, 0x54, 0x1d, 0x01, 0x85, 0xa5,
	0xbb, 0xcd, 0xa8, 0xde, 0x84, 0x1c, 0x6b, 0xfe, 0x9c, 0x36, 0x81, 0xc0, 0xfe, 0x58, 0x5f, 0x9c,
	0xbf, 0x5e, 0x5c, 0x9a, 0xbf, 0x38, 0x5b, 0x68, 0xf6, 0xf2, 0x92, 0x4e, 0x5e, 0x9f, 0x35, 0x56,
	0x18, 0xe5, 0x6d, 0xef, 0x1d, 0x61, 0xa9, 0xb1, 0x30, 0xca, 0x9b, 0xce, 0xbb, 0xcf, 0xd6, 0xc8,
	0x77, 0x51, 0x2a, 0x93, 0x9d, 0x8a, 0x60, 0x0a, 0xcd, 0xff, 0x82, 0x32, 0x0a, 0xff, 0x73, 0xc4,
	0x5f, 0x9b, 0x42, 0x27, 0xff, 0xed, 0x34, 0x7d, 0x82, 0xb7, 0xe1, 0xb0, 0xd7, 0xd9, 0x5d, 0x7a,
	0xf4, 0xef, 0xf7, 0xfc, 0x36, 0xc4, 0xfb, 0x3e, 0x76, 0x3b, 0xdc, 0x84, 0xa3, 0x79, 0x34, 0xf9,
	0xfc, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xb2, 0x28, 0x3e, 0xe0, 0x0e, 0x00, 0x00,
}
