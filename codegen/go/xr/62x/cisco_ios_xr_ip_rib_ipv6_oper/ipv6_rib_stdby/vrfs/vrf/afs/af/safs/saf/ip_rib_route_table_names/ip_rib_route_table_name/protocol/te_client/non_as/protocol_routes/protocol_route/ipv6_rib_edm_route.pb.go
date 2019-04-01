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

package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_te_client_non_as_protocol_routes_protocol_route

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
	Address                     string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	InformationSource           string   `protobuf:"bytes,2,opt,name=information_source,json=informationSource,proto3" json:"information_source,omitempty"`
	V6Nexthop                   string   `protobuf:"bytes,3,opt,name=v6_nexthop,json=v6Nexthop,proto3" json:"v6_nexthop,omitempty"`
	V6InformationSource         string   `protobuf:"bytes,4,opt,name=v6_information_source,json=v6InformationSource,proto3" json:"v6_information_source,omitempty"`
	InterfaceName               string   `protobuf:"bytes,5,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Metric                      uint32   `protobuf:"varint,6,opt,name=metric,proto3" json:"metric,omitempty"`
	LoadMetric                  uint32   `protobuf:"varint,7,opt,name=load_metric,json=loadMetric,proto3" json:"load_metric,omitempty"`
	Flags64                     uint64   `protobuf:"varint,8,opt,name=flags64,proto3" json:"flags64,omitempty"`
	Flags                       uint32   `protobuf:"varint,9,opt,name=flags,proto3" json:"flags,omitempty"`
	PrivateFlags                uint32   `protobuf:"varint,10,opt,name=private_flags,json=privateFlags,proto3" json:"private_flags,omitempty"`
	Looped                      bool     `protobuf:"varint,11,opt,name=looped,proto3" json:"looped,omitempty"`
	NextHopTableId              uint32   `protobuf:"varint,12,opt,name=next_hop_table_id,json=nextHopTableId,proto3" json:"next_hop_table_id,omitempty"`
	NextHopVrfName              string   `protobuf:"bytes,13,opt,name=next_hop_vrf_name,json=nextHopVrfName,proto3" json:"next_hop_vrf_name,omitempty"`
	NextHopTableName            string   `protobuf:"bytes,14,opt,name=next_hop_table_name,json=nextHopTableName,proto3" json:"next_hop_table_name,omitempty"`
	NextHopAfi                  uint32   `protobuf:"varint,15,opt,name=next_hop_afi,json=nextHopAfi,proto3" json:"next_hop_afi,omitempty"`
	NextHopSafi                 uint32   `protobuf:"varint,16,opt,name=next_hop_safi,json=nextHopSafi,proto3" json:"next_hop_safi,omitempty"`
	RouteLabel                  uint32   `protobuf:"varint,17,opt,name=route_label,json=routeLabel,proto3" json:"route_label,omitempty"`
	TunnelId                    uint32   `protobuf:"varint,18,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	Pathid                      uint32   `protobuf:"varint,19,opt,name=pathid,proto3" json:"pathid,omitempty"`
	BackupPathid                uint32   `protobuf:"varint,20,opt,name=backup_pathid,json=backupPathid,proto3" json:"backup_pathid,omitempty"`
	RefCntOfBackup              uint32   `protobuf:"varint,21,opt,name=ref_cnt_of_backup,json=refCntOfBackup,proto3" json:"ref_cnt_of_backup,omitempty"`
	NumberOfExtendedCommunities uint32   `protobuf:"varint,22,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities,proto3" json:"number_of_extended_communities,omitempty"`
	MvpnPresent                 bool     `protobuf:"varint,23,opt,name=mvpn_present,json=mvpnPresent,proto3" json:"mvpn_present,omitempty"`
	PathRtPresent               bool     `protobuf:"varint,24,opt,name=path_rt_present,json=pathRtPresent,proto3" json:"path_rt_present,omitempty"`
	VrfImportRtPresent          bool     `protobuf:"varint,25,opt,name=vrf_import_rt_present,json=vrfImportRtPresent,proto3" json:"vrf_import_rt_present,omitempty"`
	SourceAsrtPresent           bool     `protobuf:"varint,26,opt,name=source_asrt_present,json=sourceAsrtPresent,proto3" json:"source_asrt_present,omitempty"`
	SourceRdPresent             bool     `protobuf:"varint,27,opt,name=source_rd_present,json=sourceRdPresent,proto3" json:"source_rd_present,omitempty"`
	SegmentedNexthopPresent     bool     `protobuf:"varint,28,opt,name=segmented_nexthop_present,json=segmentedNexthopPresent,proto3" json:"segmented_nexthop_present,omitempty"`
	NextHopId                   uint32   `protobuf:"varint,29,opt,name=next_hop_id,json=nextHopId,proto3" json:"next_hop_id,omitempty"`
	NextHopIdRefcount           uint32   `protobuf:"varint,30,opt,name=next_hop_id_refcount,json=nextHopIdRefcount,proto3" json:"next_hop_id_refcount,omitempty"`
	OspfAreaId                  string   `protobuf:"bytes,31,opt,name=ospf_area_id,json=ospfAreaId,proto3" json:"ospf_area_id,omitempty"`
	RemoteBackupAddr            []string `protobuf:"bytes,32,rep,name=remote_backup_addr,json=remoteBackupAddr,proto3" json:"remote_backup_addr,omitempty"`
	HasLabelstk                 bool     `protobuf:"varint,33,opt,name=has_labelstk,json=hasLabelstk,proto3" json:"has_labelstk,omitempty"`
	NumLabels                   uint32   `protobuf:"varint,34,opt,name=num_labels,json=numLabels,proto3" json:"num_labels,omitempty"`
	Labelstk                    []uint32 `protobuf:"varint,35,rep,packed,name=labelstk,proto3" json:"labelstk,omitempty"`
	BindingLabel                uint32   `protobuf:"varint,36,opt,name=binding_label,json=bindingLabel,proto3" json:"binding_label,omitempty"`
	NhidFeid                    uint64   `protobuf:"varint,37,opt,name=nhid_feid,json=nhidFeid,proto3" json:"nhid_feid,omitempty"`
	MplsFeid                    uint64   `protobuf:"varint,38,opt,name=mpls_feid,json=mplsFeid,proto3" json:"mpls_feid,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
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

type Ipv6RibEdmPathEntry struct {
	Ipv6RibEdmPath       []*Ipv6RibEdmPathItem `protobuf:"bytes,39,rep,name=ipv6_rib_edm_path,json=ipv6RibEdmPath,proto3" json:"ipv6_rib_edm_path,omitempty"`
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
	proto.RegisterType((*Ipv6RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.te_client.non_as.protocol_routes.protocol_route.ipv6_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv6RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.te_client.non_as.protocol_routes.protocol_route.ipv6_rib_edm_path_item")
	proto.RegisterType((*Ipv6RibEdmPathEntry)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.te_client.non_as.protocol_routes.protocol_route.ipv6_rib_edm_path_entry")
	proto.RegisterType((*Ipv6RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.te_client.non_as.protocol_routes.protocol_route.ipv6_rib_edm_route")
}

func init() { proto.RegisterFile("ipv6_rib_edm_route.proto", fileDescriptor_97e3b1af6daa47d0) }

var fileDescriptor_97e3b1af6daa47d0 = []byte{
	// 1447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xdd, 0x72, 0x14, 0xb7,
	0x12, 0xae, 0xc5, 0x60, 0xaf, 0xb5, 0xfe, 0xdb, 0xb1, 0xb1, 0x65, 0xcc, 0xcf, 0x62, 0x0e, 0x1c,
	0x43, 0x1d, 0xf6, 0x9c, 0x63, 0x88, 0x93, 0x90, 0x5f, 0x63, 0x0c, 0x6c, 0x30, 0xe0, 0x2c, 0x14,
	0x95, 0x5c, 0xa9, 0xb4, 0x23, 0xcd, 0x5a, 0xc5, 0xcc, 0x68, 0x90, 0xb4, 0x83, 0xfd, 0x0e, 0xa9,
	0xca, 0x83, 0xe4, 0x22, 0x8f, 0x91, 0xdc, 0xe6, 0x2d, 0x72, 0x9f, 0x17, 0x48, 0xa9, 0x5b, 0x33,
	0x5e, 0xdb, 0xe4, 0x01, 0x92, 0x1b, 0x70, 0x7f, 0xdf, 0xd7, 0x5a, 0xf5, 0x8f, 0xba, 0x87, 0x50,
	0x55, 0x94, 0x5b, 0xcc, 0xa8, 0x01, 0x93, 0x22, 0x63, 0x46, 0x8f, 0x9c, 0xec, 0x16, 0x46, 0x3b,
	0x1d, 0xfd, 0xdc, 0x88, 0x95, 0x8d, 0x35, 0x53, 0xda, 0xb2, 0x43, 0xc3, 0x54, 0x01, 0x2a, 0x90,
	0xeb, 0x42, 0x9a, 0x6e, 0xed, 0x68, 0x9d, 0x18, 0x1c, 0x75, 0x4b, 0x93, 0x58, 0xff, 0x4f, 0x97,
	0x27, 0xb6, 0xcb, 0x93, 0xae, 0xf5, 0xff, 0x5b, 0x9e, 0x74, 0x83, 0x23, 0x1c, 0xcd, 0x1c, 0x1f,
	0xa4, 0x92, 0xe5, 0x3c, 0x93, 0xf6, 0xaf, 0x08, 0xfc, 0xf9, 0x58, 0xa7, 0x5d, 0x27, 0x59, 0x9c,
	0x2a, 0x99, 0xbb, 0x6e, 0xae, 0x73, 0xc6, 0x6d, 0xcd, 0xa0, 0xcf, 0x69, 0x7b, 0xfd, 0xb7, 0x06,
	0x59, 0x39, 0x1b, 0x0e, 0x7b, 0xb6, 0xfb, 0xfd, 0xab, 0x68, 0x95, 0x34, 0x4b, 0x93, 0xc0, 0x0f,
	0xd0, 0x46, 0xa7, 0xb1, 0x31, 0xdd, 0x9f, 0x2a, 0x4d, 0xf2, 0x82, 0x67, 0x32, 0x5a, 0x21, 0x53,
	0x3c, 0x30, 0xe7, 0x80, 0x99, 0xe4, 0x48, 0xac, 0x92, 0xa6, 0xad, 0x98, 0x09, 0xf4, 0xb1, 0x81,
	0xda, 0x20, 0x0b, 0xa7, 0xef, 0x4d, 0xcf, 0x83, 0x64, 0x0e, 0xf0, 0xd7, 0x1e, 0x06, 0x25, 0x25,
	0x53, 0x5c, 0x08, 0x23, 0xad, 0xa5, 0x17, 0xf0, 0x8c, 0x60, 0x46, 0x37, 0xc8, 0x6c, 0x61, 0x64,
	0xa2, 0x0e, 0x59, 0x2a, 0xf3, 0xa1, 0x3b, 0xa0, 0x93, 0x9d, 0xc6, 0xc6, 0x6c, 0x7f, 0x06, 0xc1,
	0x3d, 0xc0, 0xd6, 0x7f, 0x6a, 0x91, 0xe5, 0x13, 0x31, 0x15, 0xdc, 0x1d, 0x30, 0xe5, 0x64, 0x36,
	0x7e, 0x72, 0xe3, 0xe4, 0xc9, 0x77, 0x49, 0xa4, 0xf2, 0x44, 0x9b, 0x8c, 0x3b, 0xa5, 0x73, 0x66,
	0xf5, 0xc8, 0xc4, 0x55, 0x70, 0xed, 0x31, 0xe6, 0x15, 0x10, 0xd1, 0x15, 0x42, 0xca, 0x2d, 0x96,
	0xcb, 0x43, 0x77, 0xa0, 0x8b, 0x10, 0xe9, 0x74, 0xb9, 0xf5, 0x02, 0x81, 0x68, 0x93, 0x5c, 0x2c,
	0xb7, 0xd8, 0x07, 0x0e, 0xc4, 0x80, 0x17, 0xcb, 0xad, 0xde, 0x99, 0x23, 0x6f, 0x92, 0x39, 0x95,
	0x3b, 0x69, 0x12, 0x1e, 0x87, 0xec, 0x60, 0xf0, 0xb3, 0x35, 0x0a, 0xc9, 0x59, 0x26, 0x93, 0x99,
	0x74, 0x46, 0xc5, 0x21, 0xf6, 0x60, 0x45, 0xd7, 0x48, 0x2b, 0xd5, 0x5c, 0xb0, 0x40, 0x4e, 0x01,
	0x49, 0x3c, 0xf4, 0x1c, 0x05, 0x94, 0x4c, 0x25, 0x29, 0x1f, 0xda, 0xad, 0xfb, 0xb4, 0xd9, 0x69,
	0x6c, 0x9c, 0xef, 0x57, 0x66, 0xb4, 0x44, 0x2e, 0xc0, 0x9f, 0x74, 0x1a, 0x9c, 0xd0, 0xc0, 0x5c,
	0xab, 0x92, 0x3b, 0xc9, 0x90, 0x25, 0x55, 0xae, 0x01, 0x7c, 0x0c, 0xa2, 0x65, 0x32, 0x99, 0x6a,
	0x5d, 0x48, 0x41, 0x5b, 0x9d, 0xc6, 0x46, 0xb3, 0x1f, 0xac, 0xe8, 0x36, 0x69, 0xfb, 0xe4, 0xb0,
	0x03, 0x5d, 0x84, 0x7a, 0x2b, 0x41, 0x67, 0xe0, 0x80, 0x39, 0x4f, 0x3c, 0xd5, 0x05, 0xd4, 0xbb,
	0x77, 0x52, 0x5a, 0xf7, 0xdb, 0x2c, 0x36, 0x46, 0x90, 0xbe, 0x09, 0x6d, 0x77, 0x97, 0x2c, 0x9e,
	0x3a, 0x15, 0xc4, 0x73, 0x20, 0x5e, 0x18, 0x3f, 0x17, 0xe4, 0x1d, 0x32, 0x53, 0xcb, 0x79, 0xa2,
	0xe8, 0x3c, 0xe6, 0x24, 0xe8, 0xb6, 0x13, 0x15, 0xad, 0x93, 0xd9, 0x5a, 0x61, 0xbd, 0x64, 0x01,
	0x24, 0xad, 0x20, 0x79, 0xc5, 0x13, 0xe5, 0x13, 0x8b, 0x7d, 0x9b, 0xf2, 0x81, 0x4c, 0x69, 0x1b,
	0x0f, 0x01, 0x68, 0xcf, 0x23, 0xd1, 0x1a, 0x99, 0x76, 0xa3, 0x3c, 0x97, 0xa9, 0x8f, 0x31, 0x02,
	0xba, 0x89, 0x40, 0x4f, 0xf8, 0x04, 0xf9, 0xf6, 0x53, 0x82, 0x2e, 0x62, 0xb9, 0xd0, 0xf2, 0xd9,
	0x1d, 0xf0, 0xf8, 0xed, 0xa8, 0x60, 0x81, 0x5e, 0xc2, 0xec, 0x22, 0xb8, 0x8f, 0xa2, 0xdb, 0xa4,
	0x6d, 0x64, 0xc2, 0xe2, 0xdc, 0x31, 0x9d, 0x30, 0xa4, 0xe8, 0x45, 0xcc, 0xa2, 0x91, 0xc9, 0x4e,
	0xee, 0x5e, 0x26, 0x0f, 0x01, 0x8d, 0x76, 0xc8, 0xd5, 0x7c, 0x94, 0x0d, 0xa4, 0xf1, 0x4a, 0x79,
	0xe8, 0x64, 0x2e, 0xa4, 0x60, 0xb1, 0xce, 0xb2, 0x51, 0xae, 0x9c, 0x92, 0x96, 0x2e, 0x83, 0xdf,
	0x1a, 0xaa, 0x5e, 0x26, 0xbb, 0x41, 0xb3, 0x73, 0x2c, 0x89, 0xae, 0x93, 0x99, 0xac, 0x2c, 0x72,
	0x56, 0x18, 0x69, 0x65, 0xee, 0xe8, 0x0a, 0xd4, 0xb4, 0xe5, 0xb1, 0x7d, 0x84, 0xa2, 0x5b, 0x64,
	0x1e, 0x9e, 0x93, 0x71, 0xb5, 0x8a, 0x82, 0x6a, 0xd6, 0xc3, 0x7d, 0x57, 0xe9, 0xfe, 0x4f, 0x2e,
	0xfa, 0x62, 0xaa, 0xac, 0xd0, 0xc6, 0x8d, 0xab, 0x57, 0x41, 0x1d, 0x95, 0x26, 0xe9, 0x01, 0x77,
	0xec, 0xd2, 0x25, 0x8b, 0xf8, 0x4a, 0x18, 0xb7, 0x63, 0x0e, 0x97, 0xc0, 0xa1, 0x8d, 0xd4, 0xb6,
	0x35, 0xb5, 0xfe, 0x0e, 0x09, 0x20, 0x33, 0xa2, 0x56, 0xaf, 0x81, 0x7a, 0x1e, 0x89, 0xbe, 0xa8,
	0xb4, 0x0f, 0xc8, 0xaa, 0x95, 0xc3, 0x4c, 0xe6, 0x4e, 0x8a, 0xea, 0xd9, 0xd6, 0x3e, 0x97, 0xc1,
	0x67, 0xa5, 0x16, 0x84, 0x57, 0x5c, 0xf9, 0x5e, 0x25, 0xad, 0xba, 0x49, 0x94, 0xa0, 0x57, 0x20,
	0x8f, 0xd3, 0xa1, 0x45, 0x7a, 0x22, 0xfa, 0x2f, 0x59, 0x1a, 0xe3, 0x99, 0x91, 0x49, 0xac, 0x47,
	0xb9, 0xa3, 0x57, 0x41, 0xd8, 0xae, 0x85, 0xfd, 0x40, 0xf8, 0xbe, 0xd4, 0xb6, 0x48, 0x18, 0x37,
	0x92, 0xfb, 0x13, 0xaf, 0x41, 0xff, 0x12, 0x8f, 0x6d, 0x1b, 0xc9, 0x7b, 0x22, 0xfa, 0x0f, 0x89,
	0x8c, 0xcc, 0xb4, 0x93, 0xa1, 0xe8, 0xcc, 0x8f, 0x29, 0xda, 0xe9, 0x4c, 0xf8, 0x3e, 0x47, 0x06,
	0xeb, 0xbe, 0x2d, 0x84, 0xf1, 0x65, 0x3b, 0xe0, 0x16, 0xfb, 0xd3, 0xba, 0xb7, 0xf4, 0x3a, 0x96,
	0xed, 0x80, 0xdb, 0xbd, 0x00, 0xf9, 0x79, 0x95, 0x8f, 0xb2, 0x20, 0xa1, 0xeb, 0x21, 0x84, 0x51,
	0x86, 0x82, 0xe8, 0x12, 0x69, 0xd6, 0xde, 0x37, 0x3a, 0x13, 0xbe, 0x83, 0x2b, 0x1b, 0x3a, 0x55,
	0xe5, 0x42, 0xe5, 0xc3, 0xf0, 0x02, 0xfe, 0x15, 0x3a, 0x15, 0xc1, 0xfa, 0x0d, 0xe4, 0x07, 0x4a,
	0xb0, 0x44, 0x2a, 0x41, 0x6f, 0xc2, 0x78, 0x69, 0x7a, 0xe0, 0xb1, 0x54, 0xc2, 0x93, 0x59, 0x91,
	0x5a, 0x24, 0x6f, 0x21, 0xe9, 0x01, 0x4f, 0xae, 0xff, 0x78, 0xee, 0xd4, 0x06, 0x82, 0xf6, 0x92,
	0xb9, 0x33, 0x47, 0xd1, 0x1f, 0x0d, 0xd2, 0x3e, 0xc3, 0xd1, 0x7f, 0x77, 0x26, 0x36, 0x5a, 0x9b,
	0xbf, 0x34, 0xba, 0x7f, 0xb3, 0x65, 0xdb, 0xfd, 0xf0, 0x52, 0xea, 0xcf, 0x79, 0xbc, 0xaf, 0x06,
	0xbb, 0x22, 0xf3, 0xef, 0x7e, 0xfd, 0x07, 0x42, 0xa2, 0xb3, 0x3b, 0x19, 0x26, 0x09, 0xac, 0x39,
	0xba, 0x89, 0x2b, 0x17, 0x2d, 0xbf, 0x57, 0x4f, 0xec, 0x44, 0x76, 0x68, 0xe8, 0x3d, 0x9c, 0x11,
	0xe3, 0x6b, 0xf1, 0x3b, 0xe3, 0x2b, 0x89, 0xc1, 0x94, 0xd2, 0x58, 0xa5, 0x73, 0x7a, 0x1f, 0x2b,
	0x09, 0xe0, 0x1b, 0xc4, 0xfc, 0xb8, 0xab, 0xaf, 0xad, 0x04, 0xfd, 0x08, 0xc7, 0x5d, 0x05, 0xf5,
	0x04, 0xee, 0x85, 0x20, 0x80, 0xf1, 0xbb, 0x05, 0xd7, 0x99, 0xa9, 0x40, 0x18, 0xbd, 0x97, 0x48,
	0x53, 0xe5, 0xd6, 0xf1, 0x3c, 0x96, 0xf4, 0x63, 0xe0, 0x6b, 0xdb, 0xb7, 0x03, 0xa6, 0xcd, 0x9f,
	0xff, 0x09, 0xce, 0x4b, 0x04, 0x7a, 0xc2, 0x37, 0x6a, 0x48, 0xf8, 0x51, 0x21, 0xe9, 0xa7, 0xd8,
	0xa8, 0xf8, 0x7d, 0x70, 0x54, 0xc0, 0xb9, 0x85, 0x51, 0xda, 0x28, 0x77, 0x44, 0x1f, 0xa0, 0x6b,
	0x65, 0xc3, 0xb7, 0x47, 0x29, 0xd0, 0xf1, 0x33, 0xe0, 0xa6, 0x6c, 0x29, 0xc0, 0xad, 0xde, 0x70,
	0x9f, 0x8f, 0x6f, 0xb8, 0x9b, 0x64, 0xae, 0x9e, 0x94, 0x48, 0x7f, 0x01, 0xcd, 0x39, 0x5b, 0xa1,
	0xb8, 0xe3, 0x16, 0xc8, 0x84, 0xe3, 0x43, 0xfa, 0x25, 0xb8, 0xfa, 0x3f, 0xfd, 0x2d, 0x84, 0x0a,
	0xd1, 0x7d, 0x85, 0xb7, 0xa8, 0x6c, 0xff, 0x21, 0x21, 0x54, 0x48, 0x30, 0xab, 0x55, 0x5f, 0xe3,
	0x2c, 0xa8, 0x99, 0x47, 0x95, 0xfc, 0x78, 0x9d, 0x6f, 0x9f, 0x5e, 0xe7, 0xbe, 0x43, 0x2c, 0xc3,
	0x59, 0xf2, 0x30, 0x94, 0xc1, 0x43, 0x3b, 0x30, 0x44, 0xee, 0x92, 0x88, 0x3b, 0x67, 0xd4, 0xc0,
	0x27, 0x4b, 0x09, 0x99, 0x3b, 0x9f, 0x93, 0x1d, 0xfc, 0x9d, 0x9a, 0xe9, 0x05, 0xc2, 0x57, 0xcd,
	0x19, 0x9e, 0x24, 0x2a, 0x66, 0x2a, 0x17, 0xf2, 0x90, 0x3e, 0xc2, 0xda, 0x07, 0xb0, 0xe7, 0xb1,
	0xe8, 0x76, 0xf5, 0x89, 0x56, 0x18, 0x19, 0x4b, 0x21, 0xfd, 0xcd, 0x77, 0x41, 0x37, 0x0f, 0xf8,
	0x7e, 0x0d, 0xfb, 0x22, 0xbe, 0xd3, 0x96, 0x0d, 0x8d, 0x1e, 0x15, 0xf4, 0x31, 0xe6, 0xe0, 0x9d,
	0xb6, 0x4f, 0xbc, 0xed, 0x2b, 0x91, 0xa4, 0xfa, 0x3d, 0xf3, 0x69, 0x7b, 0x82, 0x95, 0xf0, 0xf6,
	0x6b, 0x3e, 0xf4, 0x7e, 0xc9, 0x7b, 0xc1, 0xe2, 0x94, 0x5b, 0x4b, 0x9f, 0xa2, 0x5f, 0xf2, 0x5e,
	0xec, 0x78, 0xdb, 0x93, 0x85, 0x8a, 0x43, 0xc8, 0xbd, 0x50, 0x5e, 0x15, 0x63, 0xc0, 0xcb, 0x64,
	0x92, 0xc7, 0x4e, 0x95, 0x92, 0x7e, 0x83, 0x9f, 0x1a, 0x68, 0x45, 0x97, 0xc9, 0x74, 0x9d, 0x56,
	0xfa, 0x0c, 0xa8, 0x63, 0x20, 0xfa, 0x1f, 0x59, 0x3a, 0x2e, 0x07, 0xb4, 0x28, 0x36, 0xed, 0x1e,
	0x34, 0xe5, 0x71, 0xa9, 0xf6, 0x3d, 0x05, 0xad, 0xbb, 0x46, 0xb0, 0xdf, 0x18, 0x1f, 0x4a, 0xfa,
	0x1c, 0x2f, 0x01, 0xc0, 0xf6, 0x50, 0x9e, 0xfe, 0x18, 0x78, 0x71, 0xe6, 0x63, 0x80, 0x92, 0xa9,
	0xea, 0x2e, 0x2f, 0x31, 0xf2, 0xf2, 0xf8, 0x61, 0xb9, 0x41, 0x5a, 0xbf, 0xbd, 0x7d, 0x68, 0x35,
	0xe2, 0x06, 0x69, 0xf5, 0xf2, 0xee, 0x90, 0x36, 0x9e, 0x9d, 0x69, 0xa1, 0x92, 0x23, 0xe6, 0x54,
	0x26, 0xe9, 0xb7, 0x20, 0xc3, 0xf4, 0x3f, 0x07, 0xfc, 0xb5, 0xca, 0x64, 0xf4, 0x7b, 0xa3, 0x7a,
	0x27, 0x30, 0x12, 0xfb, 0x9d, 0xc6, 0x46, 0x6b, 0xf3, 0xd7, 0x7f, 0xc2, 0x48, 0x84, 0xc9, 0x1f,
	0x9e, 0xbc, 0x1f, 0x87, 0x83, 0x49, 0x70, 0xb9, 0xf7, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0x17, 0x6c, 0xba, 0x77, 0x0d, 0x00, 0x00,
}
