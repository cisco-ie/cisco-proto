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

package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_q_routes_q_route

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

type Ipv4RibEdmPathEntry struct {
	Ipv4RibEdmPath       []*Ipv4RibEdmPathItem `protobuf:"bytes,39,rep,name=ipv4_rib_edm_path,json=ipv4RibEdmPath,proto3" json:"ipv4_rib_edm_path,omitempty"`
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
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.q_routes.q_route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.q_routes.q_route.ipv4_rib_edm_path_item")
	proto.RegisterType((*Ipv4RibEdmPathEntry)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.q_routes.q_route.ipv4_rib_edm_path_entry")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.q_routes.q_route.ipv4_rib_edm_route")
}

func init() { proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor_44c1b2f42127243a) }

var fileDescriptor_44c1b2f42127243a = []byte{
	// 1426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdb, 0x72, 0x14, 0xb7,
	0x16, 0xad, 0x39, 0x06, 0x7b, 0xac, 0xf1, 0x6d, 0xda, 0xc6, 0x96, 0x31, 0x97, 0xc1, 0x1c, 0x38,
	0x86, 0x3a, 0x4c, 0x12, 0x43, 0x9c, 0x84, 0x5c, 0x8d, 0x31, 0x30, 0xc1, 0x80, 0x33, 0x50, 0x54,
	0xf2, 0xa4, 0xd2, 0xb4, 0xd4, 0x63, 0x15, 0xdd, 0xad, 0x46, 0xd2, 0x34, 0x76, 0x55, 0x3e, 0x23,
	0xf9, 0x8a, 0x3c, 0xe7, 0x3d, 0xbf, 0x90, 0xc7, 0x54, 0xe5, 0x63, 0x52, 0xda, 0x5b, 0xdd, 0x1e,
	0xdb, 0xe4, 0x39, 0x79, 0x81, 0xd9, 0x6b, 0xad, 0x2d, 0x4b, 0xfb, 0xda, 0x84, 0xaa, 0xa2, 0xbc,
	0xc7, 0x8c, 0x1a, 0x30, 0x29, 0x32, 0x66, 0xf4, 0xc8, 0xc9, 0x6e, 0x61, 0xb4, 0xd3, 0xd1, 0x8f,
	0xb1, 0xb2, 0xb1, 0x66, 0x4a, 0x5b, 0x76, 0x68, 0x98, 0x2a, 0x40, 0x04, 0x6a, 0x5d, 0x48, 0xd3,
	0x35, 0x6a, 0xd0, 0x2d, 0x4d, 0x62, 0xfd, 0x3f, 0x5d, 0x9e, 0xd8, 0x2e, 0x4f, 0xba, 0xd6, 0xff,
	0x6f, 0x79, 0xd2, 0x0d, 0x6a, 0x38, 0x8e, 0x39, 0x3e, 0x48, 0x25, 0xcb, 0x79, 0x26, 0xed, 0xdf,
	0x11, 0xdd, 0xb7, 0x08, 0xd9, 0xea, 0xc7, 0xfa, 0xef, 0x0d, 0xb2, 0x72, 0xf6, 0x6a, 0xec, 0xe9,
	0xee, 0x0f, 0x2f, 0xa3, 0x55, 0xd2, 0x2c, 0x4d, 0x02, 0x8e, 0xb4, 0xd1, 0x69, 0x6c, 0x4c, 0xf7,
	0xa7, 0x4a, 0x93, 0x3c, 0xe7, 0x99, 0x8c, 0x56, 0xc8, 0x14, 0x0f, 0xcc, 0x7f, 0x80, 0x99, 0xe4,
	0x48, 0xac, 0x92, 0xa6, 0xad, 0x98, 0x09, 0xf4, 0xb1, 0x81, 0xda, 0x20, 0x0b, 0xa7, 0xef, 0x43,
	0xcf, 0x81, 0x64, 0x0e, 0xf0, 0x57, 0x1e, 0x06, 0x25, 0x25, 0x53, 0x5c, 0x08, 0x23, 0xad, 0xa5,
	0xe7, 0xf1, 0x8c, 0x60, 0x46, 0xd7, 0xc9, 0x6c, 0x61, 0x64, 0xa2, 0x0e, 0x59, 0x2a, 0xf3, 0xa1,
	0x3b, 0xa0, 0x93, 0x9d, 0xc6, 0xc6, 0x6c, 0x7f, 0x06, 0xc1, 0x3d, 0xc0, 0xd6, 0x7f, 0x69, 0x91,
	0xe5, 0x13, 0x6f, 0x2a, 0xb8, 0x3b, 0x60, 0xca, 0xc9, 0x6c, 0xfc, 0xe4, 0xc6, 0xc9, 0x93, 0xef,
	0x90, 0x48, 0xe5, 0x89, 0x36, 0x19, 0x77, 0x4a, 0xe7, 0xcc, 0xea, 0x91, 0x89, 0xab, 0xc7, 0xb5,
	0xc7, 0x98, 0x97, 0x40, 0x44, 0x97, 0x09, 0x29, 0xb7, 0x58, 0x2e, 0x0f, 0xdd, 0x81, 0x2e, 0xc2,
	0x4b, 0xa7, 0xcb, 0xad, 0xe7, 0x08, 0x44, 0x9b, 0xe4, 0x42, 0xb9, 0xc5, 0xde, 0x73, 0x20, 0x3e,
	0x78, 0xb1, 0xdc, 0xea, 0x9d, 0x39, 0xf2, 0x06, 0x99, 0x53, 0xb9, 0x93, 0x26, 0xe1, 0x71, 0x88,
	0x0e, 0x3e, 0x7e, 0xb6, 0x46, 0x21, 0x38, 0xcb, 0x64, 0x32, 0x93, 0xce, 0xa8, 0x38, 0xbc, 0x3d,
	0x58, 0xd1, 0x55, 0xd2, 0x4a, 0x35, 0x17, 0x2c, 0x90, 0x53, 0x40, 0x12, 0x0f, 0x3d, 0x43, 0x01,
	0x25, 0x53, 0x49, 0xca, 0x87, 0x76, 0xeb, 0x1e, 0x6d, 0x76, 0x1a, 0x1b, 0xe7, 0xfa, 0x95, 0x19,
	0x2d, 0x91, 0xf3, 0xf0, 0x93, 0x4e, 0x83, 0x13, 0x1a, 0x18, 0x6b, 0x55, 0x72, 0x27, 0x19, 0xb2,
	0xa4, 0x8a, 0x35, 0x80, 0x8f, 0x40, 0xb4, 0x4c, 0x26, 0x53, 0xad, 0x0b, 0x29, 0x68, 0xab, 0xd3,
	0xd8, 0x68, 0xf6, 0x83, 0x15, 0xdd, 0x22, 0x6d, 0x1f, 0x1c, 0x76, 0xa0, 0x8b, 0x90, 0x6f, 0x25,
	0xe8, 0x0c, 0x1c, 0x30, 0xe7, 0x89, 0x27, 0xba, 0x80, 0x7c, 0xf7, 0x4e, 0x4a, 0xeb, 0x7a, 0x9b,
	0xc5, 0xc2, 0x08, 0xd2, 0xd7, 0xa1, 0xec, 0xee, 0x90, 0xc5, 0x53, 0xa7, 0x82, 0x78, 0x0e, 0xc4,
	0x0b, 0xe3, 0xe7, 0x82, 0xbc, 0x43, 0x66, 0x6a, 0x39, 0x4f, 0x14, 0x9d, 0xc7, 0x98, 0x04, 0xdd,
	0x76, 0xa2, 0xa2, 0x75, 0x32, 0x5b, 0x2b, 0xac, 0x97, 0x2c, 0x80, 0xa4, 0x15, 0x24, 0x2f, 0x79,
	0xa2, 0x7c, 0x60, 0xb1, 0x6e, 0x53, 0x3e, 0x90, 0x29, 0x6d, 0xe3, 0x21, 0x00, 0xed, 0x79, 0x24,
	0x5a, 0x23, 0xd3, 0x6e, 0x94, 0xe7, 0x32, 0xf5, 0x6f, 0x8c, 0x80, 0x6e, 0x22, 0xd0, 0x13, 0x3e,
	0x40, 0xbe, 0xfc, 0x94, 0xa0, 0x8b, 0x98, 0x2e, 0xb4, 0x7c, 0x74, 0x07, 0x3c, 0x7e, 0x33, 0x2a,
	0x58, 0xa0, 0x97, 0x30, 0xba, 0x08, 0xee, 0xa3, 0xe8, 0x16, 0x69, 0x1b, 0x99, 0xb0, 0x38, 0x77,
	0x4c, 0x27, 0x0c, 0x29, 0x7a, 0x01, 0xa3, 0x68, 0x64, 0xb2, 0x93, 0xbb, 0x17, 0xc9, 0x03, 0x40,
	0xa3, 0x1d, 0x72, 0x25, 0x1f, 0x65, 0x03, 0x69, 0xbc, 0x52, 0x1e, 0x3a, 0x99, 0x0b, 0x29, 0x58,
	0xac, 0xb3, 0x6c, 0x94, 0x2b, 0xa7, 0xa4, 0xa5, 0xcb, 0xe0, 0xb7, 0x86, 0xaa, 0x17, 0xc9, 0x6e,
	0xd0, 0xec, 0x1c, 0x4b, 0xa2, 0x6b, 0x64, 0x26, 0x2b, 0x8b, 0x9c, 0x15, 0x46, 0x5a, 0x99, 0x3b,
	0xba, 0x02, 0x39, 0x6d, 0x79, 0x6c, 0x1f, 0xa1, 0xe8, 0x26, 0x99, 0x87, 0x76, 0x32, 0xae, 0x56,
	0x51, 0x50, 0xcd, 0x7a, 0xb8, 0xef, 0x2a, 0xdd, 0x47, 0xe4, 0x82, 0x4f, 0xa6, 0xca, 0x0a, 0x6d,
	0xdc, 0xb8, 0x7a, 0x15, 0xd4, 0x51, 0x69, 0x92, 0x1e, 0x70, 0xc7, 0x2e, 0x5d, 0xb2, 0x88, 0x5d,
	0xc2, 0xb8, 0x1d, 0x73, 0xb8, 0x08, 0x0e, 0x6d, 0xa4, 0xb6, 0xad, 0xa9, 0xf5, 0xb7, 0x49, 0x00,
	0x99, 0x11, 0xb5, 0x7a, 0x0d, 0xd4, 0xf3, 0x48, 0xf4, 0x45, 0xa5, 0xbd, 0x4f, 0x56, 0xad, 0x1c,
	0x66, 0x32, 0x77, 0x52, 0x54, 0x6d, 0x5b, 0xfb, 0x5c, 0x02, 0x9f, 0x95, 0x5a, 0x10, 0xba, 0xb8,
	0xf2, 0xbd, 0x42, 0x5a, 0x75, 0x91, 0x28, 0x41, 0x2f, 0x43, 0x1c, 0xa7, 0x43, 0x89, 0xf4, 0x44,
	0xf4, 0x01, 0x59, 0x1a, 0xe3, 0x99, 0x91, 0x49, 0xac, 0x47, 0xb9, 0xa3, 0x57, 0x40, 0xd8, 0xae,
	0x85, 0xfd, 0x40, 0xf8, 0xba, 0xd4, 0xb6, 0x48, 0x18, 0x37, 0x92, 0xfb, 0x13, 0xaf, 0x42, 0xfd,
	0x12, 0x8f, 0x6d, 0x1b, 0xc9, 0x7b, 0x22, 0xfa, 0x3f, 0x89, 0x8c, 0xcc, 0xb4, 0x93, 0x21, 0xe9,
	0xcc, 0x8f, 0x29, 0xda, 0xe9, 0x4c, 0xf8, 0x3a, 0x47, 0x06, 0xf3, 0xbe, 0x2d, 0x84, 0xf1, 0x69,
	0x3b, 0xe0, 0x16, 0xeb, 0xd3, 0xba, 0x37, 0xf4, 0x1a, 0xa6, 0xed, 0x80, 0xdb, 0xbd, 0x00, 0xf9,
	0x79, 0x95, 0x8f, 0xb2, 0x20, 0xa1, 0xeb, 0xe1, 0x09, 0xa3, 0x0c, 0x05, 0xd1, 0x45, 0xd2, 0xac,
	0xbd, 0xaf, 0x77, 0x26, 0x7c, 0x05, 0x57, 0x36, 0x54, 0xaa, 0xca, 0x85, 0xca, 0x87, 0xa1, 0x03,
	0xfe, 0x1b, 0x2a, 0x15, 0xc1, 0xba, 0x07, 0xf2, 0x03, 0x25, 0x58, 0x22, 0x95, 0xa0, 0x37, 0x60,
	0xbc, 0x34, 0x3d, 0xf0, 0x48, 0x2a, 0xe1, 0xc9, 0xac, 0x48, 0x2d, 0x92, 0x37, 0x91, 0xf4, 0x80,
	0x27, 0xd7, 0xff, 0x3c, 0xbd, 0x81, 0xa0, 0xbc, 0x64, 0xee, 0xcc, 0x51, 0xf4, 0x5b, 0x83, 0xb4,
	0xcf, 0x70, 0xf4, 0x7f, 0x9d, 0x89, 0x8d, 0xd6, 0xe6, 0x4f, 0x8d, 0xee, 0x3f, 0xb9, 0x39, 0xbb,
	0xef, 0xdf, 0x30, 0xfd, 0x39, 0x8f, 0xf7, 0xd5, 0x60, 0x57, 0x64, 0xbe, 0x89, 0xd7, 0xff, 0x98,
	0x26, 0xd1, 0xd9, 0x05, 0x0b, 0x63, 0x01, 0x76, 0x16, 0xdd, 0xc4, 0xfd, 0x89, 0x96, 0x5f, 0x92,
	0x27, 0x16, 0x1c, 0x3b, 0x34, 0xf4, 0x2e, 0x36, 0xfc, 0xf8, 0x8e, 0xfb, 0xde, 0xf8, 0xb4, 0xe0,
	0x25, 0x4b, 0x69, 0xac, 0xd2, 0x39, 0xbd, 0x87, 0x69, 0x01, 0xf0, 0x35, 0x62, 0x7e, 0x76, 0xc1,
	0x57, 0x46, 0xac, 0x61, 0x38, 0x7d, 0x8c, 0xb3, 0xab, 0x82, 0x7a, 0x02, 0x87, 0x7c, 0x10, 0xc0,
	0x2c, 0xdd, 0x82, 0xeb, 0xcc, 0x54, 0x20, 0xcc, 0xd1, 0x8b, 0xa4, 0xa9, 0x72, 0xeb, 0x78, 0x1e,
	0x4b, 0xfa, 0x09, 0xf0, 0xb5, 0xed, 0x73, 0x1b, 0xa7, 0x4a, 0xe6, 0xce, 0x9f, 0xff, 0x29, 0x0e,
	0x3f, 0x04, 0x7a, 0xc2, 0x57, 0x5d, 0x08, 0xe4, 0x51, 0x21, 0xe9, 0x67, 0x58, 0x75, 0xb8, 0xec,
	0x8f, 0x0a, 0x38, 0xb7, 0x30, 0x4a, 0x1b, 0xe5, 0x8e, 0xe8, 0x7d, 0x74, 0xad, 0x6c, 0xf8, 0x90,
	0x28, 0x05, 0x3a, 0x7e, 0x0e, 0xdc, 0x94, 0x2d, 0x05, 0xb8, 0xd5, 0xeb, 0xea, 0x8b, 0xf1, 0x75,
	0x75, 0x83, 0xcc, 0xd5, 0x63, 0x0f, 0xe9, 0x2f, 0xa1, 0xd2, 0x66, 0x2b, 0x14, 0x17, 0xd6, 0x02,
	0x99, 0x70, 0x7c, 0x48, 0xbf, 0x02, 0x57, 0xff, 0xd3, 0xdf, 0x42, 0xa8, 0xf0, 0xba, 0xaf, 0xf1,
	0x16, 0x95, 0xed, 0xbf, 0x0a, 0x84, 0x0a, 0x01, 0x66, 0xb5, 0xea, 0x1b, 0x6c, 0xec, 0x9a, 0x79,
	0x58, 0xc9, 0x8f, 0x77, 0xf3, 0xf6, 0xe9, 0xdd, 0xec, 0x2b, 0xc4, 0x32, 0x1c, 0x0c, 0x0f, 0x42,
	0x1a, 0x3c, 0xb4, 0x03, 0x13, 0xe1, 0x0e, 0x89, 0xb8, 0x73, 0x46, 0x0d, 0x7c, 0xb0, 0x94, 0x90,
	0xb9, 0xf3, 0x31, 0xd9, 0xc1, 0xbf, 0x53, 0x33, 0xbd, 0x40, 0xf8, 0xac, 0x39, 0xc3, 0x93, 0x44,
	0xc5, 0x4c, 0xe5, 0x42, 0x1e, 0xd2, 0x87, 0x98, 0xfb, 0x00, 0xf6, 0x3c, 0x16, 0xdd, 0xaa, 0xbe,
	0xb7, 0x0a, 0x23, 0x63, 0x29, 0xa4, 0xbf, 0xf9, 0x2e, 0xe8, 0xe6, 0x01, 0xdf, 0xaf, 0x61, 0x9f,
	0xc4, 0xb7, 0xda, 0xb2, 0xa1, 0xd1, 0xa3, 0x82, 0x3e, 0xc2, 0x18, 0xbc, 0xd5, 0xf6, 0xb1, 0xb7,
	0x7d, 0x26, 0x92, 0x54, 0xbf, 0x63, 0x3e, 0x6c, 0x8f, 0x31, 0x13, 0xde, 0x7e, 0xc5, 0x87, 0xde,
	0x2f, 0x79, 0x27, 0x58, 0x9c, 0x72, 0x6b, 0xe9, 0x13, 0xf4, 0x4b, 0xde, 0x89, 0x1d, 0x6f, 0x7b,
	0xb2, 0x50, 0x71, 0x78, 0x72, 0x2f, 0xa4, 0x57, 0xc5, 0xf8, 0xe0, 0x65, 0x32, 0xc9, 0x63, 0xa7,
	0x4a, 0x49, 0xbf, 0xc5, 0xef, 0x06, 0xb4, 0xa2, 0x4b, 0x64, 0xba, 0x0e, 0x2b, 0x7d, 0x0a, 0xd4,
	0x31, 0x10, 0x7d, 0x48, 0x96, 0x8e, 0xd3, 0x01, 0x25, 0x8a, 0x45, 0xbb, 0x07, 0x45, 0x79, 0x9c,
	0xaa, 0x7d, 0x4f, 0x41, 0xe9, 0xae, 0x11, 0xac, 0x37, 0xc6, 0x87, 0x92, 0x3e, 0xc3, 0x4b, 0x00,
	0xb0, 0x3d, 0x94, 0xa7, 0x37, 0xfb, 0xf3, 0x33, 0x9b, 0x9d, 0x92, 0xa9, 0xea, 0x2e, 0x2f, 0xf0,
	0xe5, 0xe5, 0x71, 0x63, 0xb9, 0x41, 0x5a, 0xf7, 0xde, 0x3e, 0x94, 0x1a, 0x71, 0x83, 0xb4, 0xea,
	0xbc, 0xdb, 0xa4, 0x8d, 0x67, 0x67, 0x5a, 0xa8, 0xe4, 0x88, 0x39, 0x95, 0x49, 0xfa, 0x1d, 0xc8,
	0x30, 0xfc, 0xcf, 0x00, 0x7f, 0xa5, 0x32, 0x19, 0xfd, 0xda, 0xa8, 0xfa, 0x04, 0xe6, 0x5b, 0xbf,
	0xd3, 0xd8, 0x68, 0x6d, 0xfe, 0xfc, 0xaf, 0x9b, 0x6f, 0x30, 0x93, 0x43, 0xff, 0xfa, 0xd9, 0x36,
	0x98, 0x84, 0x14, 0xdc, 0xfd, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x0e, 0x88, 0xcb, 0xdd, 0x0c,
	0x00, 0x00,
}
