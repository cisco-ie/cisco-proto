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

package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_rib_table_ids_rib_table_id_rib_table_itf_hndls_rib_table_itf_hndl_itf_route

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
	Tableid              string   `protobuf:"bytes,1,opt,name=tableid,proto3" json:"tableid,omitempty"`
	Handle               uint32   `protobuf:"varint,2,opt,name=handle,proto3" json:"handle,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
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

func (m *Ipv6RibEdmRoute_KEYS) GetTableid() string {
	if m != nil {
		return m.Tableid
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetHandle() uint32 {
	if m != nil {
		return m.Handle
	}
	return 0
}

func (m *Ipv6RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
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
	proto.RegisterType((*Ipv6RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.rib_table_ids.rib_table_id.rib_table_itf_hndls.rib_table_itf_hndl.itf_route.ipv6_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv6RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.rib_table_ids.rib_table_id.rib_table_itf_hndls.rib_table_itf_hndl.itf_route.ipv6_rib_edm_path_item")
	proto.RegisterType((*Ipv6RibEdmPathEntry)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.rib_table_ids.rib_table_id.rib_table_itf_hndls.rib_table_itf_hndl.itf_route.ipv6_rib_edm_path_entry")
	proto.RegisterType((*Ipv6RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.rib_table_ids.rib_table_id.rib_table_itf_hndls.rib_table_itf_hndl.itf_route.ipv6_rib_edm_route")
}

func init() { proto.RegisterFile("ipv6_rib_edm_route.proto", fileDescriptor_97e3b1af6daa47d0) }

var fileDescriptor_97e3b1af6daa47d0 = []byte{
	// 1388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x72, 0x13, 0x47,
	0x16, 0x2e, 0xad, 0xc1, 0xb6, 0x5a, 0xfe, 0x53, 0xdb, 0x98, 0x06, 0xf3, 0x23, 0xcc, 0xc2, 0x1a,
	0x6a, 0xf1, 0xee, 0x1a, 0xd6, 0xbb, 0x21, 0xbf, 0xc6, 0x18, 0x50, 0x30, 0xe0, 0x08, 0x17, 0x95,
	0x5c, 0x75, 0x8d, 0xa6, 0x7b, 0xa4, 0x2e, 0x66, 0x7a, 0x86, 0x9e, 0x96, 0xb0, 0xdf, 0x22, 0x79,
	0x8e, 0x5c, 0xe5, 0x3e, 0xaf, 0x91, 0xaa, 0x3c, 0x4e, 0xea, 0x9c, 0xd3, 0x33, 0x96, 0x6d, 0xee,
	0x93, 0xbb, 0x39, 0xdf, 0xf7, 0x9d, 0x56, 0xf7, 0xf9, 0x15, 0x13, 0xa6, 0x18, 0x6f, 0x4b, 0x67,
	0xfa, 0x52, 0xab, 0x4c, 0xba, 0x7c, 0xe4, 0xf5, 0x66, 0xe1, 0x72, 0x9f, 0xf3, 0xa3, 0xd8, 0x94,
	0x71, 0x2e, 0x4d, 0x5e, 0xca, 0x23, 0x27, 0x4d, 0x81, 0x22, 0x54, 0xe7, 0x85, 0x76, 0x9b, 0xb5,
	0x5f, 0xe9, 0x55, 0xff, 0x78, 0x13, 0xbe, 0x7c, 0xd4, 0x4f, 0xb5, 0x34, 0xaa, 0x3c, 0x65, 0x4d,
	0x1a, 0x3e, 0x91, 0x43, 0xab, 0xd2, 0xf2, 0x13, 0xd8, 0x26, 0x7c, 0xe0, 0xef, 0xaf, 0x6b, 0x76,
	0xf9, 0xfc, 0xad, 0xe4, 0xcb, 0xbd, 0x1f, 0xde, 0x72, 0xc1, 0x66, 0xd0, 0xcd, 0x28, 0xd1, 0xe8,
	0x34, 0x36, 0x9a, 0xbd, 0xca, 0xe4, 0xab, 0x6c, 0x7a, 0x18, 0x59, 0x95, 0x6a, 0xf1, 0xb7, 0x4e,
	0x63, 0x63, 0xbe, 0x17, 0x2c, 0xf0, 0x88, 0x94, 0x72, 0xba, 0x2c, 0xc5, 0x14, 0x79, 0x04, 0x73,
	0xfd, 0xe7, 0x16, 0x5b, 0x3d, 0xf5, 0x3b, 0x45, 0xe4, 0x87, 0xd2, 0x78, 0x9d, 0x4d, 0x3a, 0x35,
	0x4e, 0x39, 0xf1, 0x07, 0x8c, 0x1b, 0x9b, 0xe4, 0x2e, 0x8b, 0xbc, 0xc9, 0xad, 0x2c, 0xf3, 0x91,
	0x8b, 0xe9, 0x27, 0x9b, 0xbd, 0xf6, 0x04, 0xf3, 0x16, 0x09, 0x7e, 0x9d, 0xb1, 0xf1, 0xb6, 0xb4,
	0xfa, 0xc8, 0x0f, 0xf3, 0x22, 0x5c, 0xa0, 0x39, 0xde, 0x7e, 0x4d, 0x00, 0xdf, 0x62, 0x97, 0xc6,
	0xdb, 0xf2, 0x13, 0x07, 0x5e, 0x40, 0xe5, 0xf2, 0x78, 0xbb, 0x7b, 0xee, 0xc8, 0x3b, 0x6c, 0xc1,
	0x58, 0xaf, 0x5d, 0x12, 0xc5, 0x5a, 0xda, 0x28, 0xd3, 0xe2, 0x22, 0x8a, 0xe7, 0x6b, 0xf4, 0x75,
	0x94, 0x69, 0x88, 0x47, 0xa6, 0xbd, 0x33, 0xb1, 0x98, 0xa6, 0x78, 0x90, 0xc5, 0x6f, 0xb2, 0x56,
	0x9a, 0x47, 0x4a, 0x06, 0x72, 0x06, 0x49, 0x06, 0xd0, 0x2b, 0x12, 0x08, 0x36, 0x93, 0xa4, 0xd1,
	0xa0, 0xdc, 0x7e, 0x24, 0x66, 0x3b, 0x8d, 0x8d, 0x0b, 0xbd, 0xca, 0xe4, 0x2b, 0xec, 0x22, 0x7e,
	0x8a, 0x26, 0x3a, 0x91, 0xc1, 0x6f, 0xb3, 0xf9, 0xc2, 0x99, 0x71, 0xe4, 0xb5, 0x24, 0x96, 0x21,
	0x3b, 0x17, 0xc0, 0x67, 0x28, 0x5a, 0x65, 0xd3, 0x69, 0x9e, 0x17, 0x5a, 0x89, 0x56, 0xa7, 0xb1,
	0x31, 0xdb, 0x0b, 0x16, 0xbf, 0xc7, 0xda, 0x10, 0x1c, 0x39, 0xcc, 0x8b, 0xba, 0x60, 0xc4, 0x1c,
	0x1e, 0xb0, 0x00, 0xc4, 0x8b, 0xbc, 0x38, 0x04, 0xb8, 0x7b, 0x5a, 0x3a, 0x76, 0x09, 0x3d, 0x7d,
	0x1e, 0x9f, 0x5e, 0x49, 0xdf, 0xb9, 0x04, 0xdf, 0xfe, 0x80, 0x2d, 0x9f, 0x39, 0x15, 0xc5, 0x0b,
	0x28, 0x5e, 0x9a, 0x3c, 0x17, 0xe5, 0x1d, 0x36, 0x57, 0xcb, 0xa3, 0xc4, 0x88, 0x45, 0x8a, 0x49,
	0xd0, 0xed, 0x24, 0x86, 0xaf, 0xb3, 0xf9, 0x5a, 0x51, 0x82, 0x64, 0x09, 0x25, 0xad, 0x20, 0x79,
	0x1b, 0x25, 0x06, 0x02, 0x4b, 0x85, 0x9a, 0x46, 0x7d, 0x9d, 0x8a, 0x36, 0x1d, 0x82, 0xd0, 0x3e,
	0x20, 0x7c, 0x8d, 0x35, 0xfd, 0xc8, 0x5a, 0x9d, 0xc2, 0x1b, 0x39, 0xd2, 0xb3, 0x04, 0x74, 0xb1,
	0x7c, 0xa1, 0xfc, 0x8c, 0x12, 0xcb, 0x94, 0x2e, 0xb2, 0x20, 0xba, 0xfd, 0x28, 0x7e, 0x3f, 0x2a,
	0x64, 0xa0, 0x57, 0x28, 0xba, 0x04, 0x1e, 0x90, 0xe8, 0x1e, 0x6b, 0x3b, 0x9d, 0xc8, 0xd8, 0x7a,
	0x99, 0x27, 0x92, 0x28, 0x71, 0x89, 0xa2, 0xe8, 0x74, 0xb2, 0x6b, 0xfd, 0x9b, 0xe4, 0x09, 0xa2,
	0x7c, 0x97, 0xdd, 0xb0, 0xa3, 0xac, 0xaf, 0x1d, 0x28, 0xf5, 0x91, 0xd7, 0x56, 0x69, 0x25, 0xe3,
	0x3c, 0xcb, 0x46, 0xd6, 0x78, 0xa3, 0x4b, 0xb1, 0x8a, 0x7e, 0x6b, 0xa4, 0x7a, 0x93, 0xec, 0x05,
	0xcd, 0xee, 0x89, 0x84, 0xdf, 0x62, 0x73, 0xd9, 0xb8, 0xb0, 0xb2, 0x70, 0xba, 0xd4, 0xd6, 0x8b,
	0xcb, 0x98, 0xd3, 0x16, 0x60, 0x07, 0x04, 0xf1, 0xbb, 0x6c, 0x11, 0xdb, 0xc9, 0xf9, 0x5a, 0x25,
	0x50, 0x35, 0x0f, 0x70, 0xcf, 0x57, 0xba, 0xff, 0xb0, 0x4b, 0x90, 0x4c, 0x93, 0x15, 0xb9, 0xf3,
	0x93, 0xea, 0x2b, 0xa8, 0xe6, 0x63, 0x97, 0x74, 0x91, 0x3b, 0x71, 0xd9, 0x64, 0xcb, 0xd4, 0x25,
	0x32, 0x2a, 0x27, 0x1c, 0xae, 0xa2, 0x43, 0x9b, 0xa8, 0x9d, 0xd2, 0xd5, 0xfa, 0xfb, 0x2c, 0x80,
	0xd2, 0xa9, 0x5a, 0xbd, 0x86, 0xea, 0x45, 0x22, 0x7a, 0xaa, 0xd2, 0x3e, 0x66, 0x57, 0x4a, 0x3d,
	0xc8, 0xb4, 0xf5, 0x5a, 0x55, 0x6d, 0x5b, 0xfb, 0x5c, 0x43, 0x9f, 0xcb, 0xb5, 0x20, 0x74, 0x71,
	0xe5, 0x7b, 0x83, 0xb5, 0xea, 0x22, 0x31, 0x4a, 0x5c, 0xc7, 0x38, 0x36, 0x43, 0x89, 0x74, 0x15,
	0xff, 0x17, 0x5b, 0x99, 0xe0, 0xa5, 0xd3, 0x49, 0x9c, 0x8f, 0xac, 0x17, 0x37, 0x50, 0xd8, 0xae,
	0x85, 0xbd, 0x40, 0x40, 0x5d, 0xe6, 0x65, 0x91, 0xc8, 0xc8, 0xe9, 0x08, 0x4e, 0xbc, 0x89, 0xf5,
	0xcb, 0x00, 0xdb, 0x71, 0x3a, 0xea, 0x2a, 0xfe, 0x4f, 0xc6, 0x9d, 0xce, 0x72, 0xaf, 0x43, 0xd2,
	0x25, 0x8c, 0x29, 0xd1, 0xe9, 0x4c, 0x41, 0x9d, 0x13, 0x43, 0x79, 0xdf, 0x51, 0xca, 0x41, 0xda,
	0x86, 0x51, 0x49, 0xf5, 0x59, 0xfa, 0xf7, 0xe2, 0x16, 0xa5, 0x6d, 0x18, 0x95, 0xfb, 0x01, 0x82,
	0x79, 0x65, 0x47, 0x59, 0x90, 0x88, 0xf5, 0xf0, 0x84, 0x51, 0x46, 0x02, 0x7e, 0x95, 0xcd, 0xd6,
	0xde, 0xb7, 0x3b, 0x53, 0x50, 0xc1, 0x95, 0x8d, 0x95, 0x6a, 0xac, 0x32, 0x76, 0x10, 0x3a, 0xe0,
	0xef, 0xa1, 0x52, 0x09, 0xac, 0x7b, 0xc0, 0x0e, 0x8d, 0x92, 0x09, 0x4c, 0xf0, 0x3b, 0x38, 0x5e,
	0x66, 0x01, 0x78, 0x06, 0x23, 0x7c, 0x8d, 0x35, 0xb3, 0x22, 0x2d, 0x89, 0xbc, 0x4b, 0x24, 0x00,
	0x40, 0xae, 0xff, 0xde, 0x38, 0xb3, 0x15, 0xb0, 0xbc, 0xb4, 0xf5, 0xee, 0x98, 0xff, 0xda, 0x60,
	0xed, 0x73, 0x9c, 0xf8, 0x47, 0x67, 0x6a, 0xa3, 0xb5, 0xf5, 0x63, 0x63, 0xf3, 0xcf, 0x5a, 0x64,
	0x9b, 0x9f, 0xde, 0x2e, 0xbd, 0x05, 0xc0, 0x7b, 0xa6, 0xbf, 0xa7, 0x32, 0x68, 0xe0, 0xf5, 0xdf,
	0x9a, 0x8c, 0x9f, 0x5f, 0x78, 0x38, 0x12, 0x9c, 0x4e, 0xcc, 0x91, 0xd8, 0xc2, 0xc4, 0x07, 0x8b,
	0x6f, 0xb0, 0x25, 0xfa, 0x92, 0xa9, 0xb6, 0x03, 0x3f, 0x94, 0x47, 0x4e, 0x3c, 0xa4, 0x66, 0x27,
	0x7c, 0x1f, 0xe1, 0xef, 0x1d, 0xa4, 0x84, 0x46, 0xd2, 0x58, 0xbb, 0xd2, 0xe4, 0x56, 0x3c, 0xa2,
	0x94, 0x20, 0xf8, 0x8e, 0x30, 0x98, 0x5b, 0xb8, 0xf0, 0xe3, 0x1c, 0x07, 0xd3, 0x7f, 0x69, 0x6e,
	0x55, 0x50, 0x57, 0xd1, 0x80, 0x0f, 0x02, 0x9c, 0xa3, 0xdb, 0x78, 0x9d, 0xb9, 0x0a, 0xc4, 0x19,
	0x7a, 0x95, 0xcd, 0x1a, 0x5b, 0xfa, 0xc8, 0xc6, 0x5a, 0xfc, 0x0f, 0xf9, 0xda, 0x86, 0xbc, 0xc6,
	0xa9, 0xd1, 0xd6, 0xc3, 0xf9, 0xff, 0xa7, 0xc1, 0x47, 0x40, 0x57, 0x41, 0xc5, 0xd1, 0x1d, 0xfd,
	0x71, 0xa1, 0xc5, 0x67, 0x54, 0x71, 0x88, 0x1c, 0x1e, 0x17, 0x78, 0x6e, 0xe1, 0x4c, 0xee, 0x8c,
	0x3f, 0x16, 0x8f, 0xc9, 0xb5, 0xb2, 0xf9, 0x15, 0x36, 0x5b, 0x8e, 0x15, 0x39, 0x7e, 0x8e, 0xdc,
	0x4c, 0x39, 0x56, 0xe8, 0x56, 0xaf, 0xaa, 0x2f, 0x26, 0x57, 0xd5, 0x1d, 0xb6, 0x50, 0x8f, 0x3c,
	0xa2, 0xbf, 0xc4, 0x2a, 0x9b, 0xaf, 0x50, 0x5a, 0x56, 0x4b, 0x6c, 0xca, 0x47, 0x03, 0xf1, 0x15,
	0xba, 0xc2, 0x27, 0xdc, 0x42, 0x99, 0xf0, 0xba, 0xaf, 0xe9, 0x16, 0x95, 0x0d, 0xff, 0x08, 0x94,
	0x09, 0x01, 0x96, 0xb5, 0xea, 0x1b, 0x6a, 0xea, 0x9a, 0x79, 0x5a, 0xc9, 0x4f, 0xf6, 0xf2, 0xce,
	0xd9, 0xbd, 0x0c, 0x15, 0x52, 0x4a, 0x1a, 0x0a, 0x4f, 0x42, 0x1a, 0x00, 0xda, 0xc5, 0x69, 0xf0,
	0x80, 0xf1, 0xc8, 0x7b, 0x67, 0xfa, 0x10, 0x2c, 0xa3, 0xb4, 0xf5, 0x10, 0x93, 0x5d, 0xfa, 0x9d,
	0x9a, 0xe9, 0x06, 0x02, 0xb2, 0xe6, 0x5d, 0x94, 0x24, 0x26, 0x96, 0xc6, 0x2a, 0x7d, 0x24, 0x9e,
	0x52, 0xee, 0x03, 0xd8, 0x05, 0x8c, 0xdf, 0x63, 0x4b, 0x14, 0xfc, 0xc2, 0xe9, 0x58, 0x2b, 0x0d,
	0x37, 0xdf, 0x43, 0xdd, 0x22, 0xe2, 0x07, 0x35, 0x0c, 0x49, 0xfc, 0x90, 0x97, 0x72, 0xe0, 0xf2,
	0x51, 0x21, 0x9e, 0x51, 0x0c, 0x3e, 0xe4, 0xe5, 0x73, 0xb0, 0x21, 0x13, 0x49, 0x9a, 0x7f, 0x94,
	0x10, 0xb6, 0xe7, 0x94, 0x09, 0xb0, 0x0f, 0xa3, 0x01, 0xf8, 0x25, 0x1f, 0x95, 0x8c, 0xd3, 0xa8,
	0x2c, 0xc5, 0x0b, 0xf2, 0x4b, 0x3e, 0xaa, 0x5d, 0xb0, 0x81, 0x2c, 0x4c, 0x1c, 0x9e, 0xdc, 0x0d,
	0xe9, 0x35, 0x31, 0x3d, 0x78, 0x95, 0x4d, 0x47, 0xb1, 0x37, 0x63, 0x2d, 0xbe, 0xa5, 0xff, 0x0c,
	0x64, 0xf1, 0x6b, 0xac, 0x59, 0x87, 0x55, 0xbc, 0x44, 0xea, 0x04, 0xe0, 0xff, 0x66, 0x2b, 0x27,
	0xe9, 0xc0, 0x12, 0xa5, 0xa2, 0xdd, 0xc7, 0xa2, 0x3c, 0x49, 0xd5, 0x01, 0x50, 0x58, 0xba, 0x6b,
	0x8c, 0xea, 0x4d, 0x46, 0x03, 0x2d, 0x5e, 0xd1, 0x25, 0x10, 0xd8, 0x19, 0xe8, 0xb3, 0x5b, 0xfd,
	0xf5, 0xb9, 0xad, 0x2e, 0xd8, 0x4c, 0x75, 0x97, 0x37, 0xf4, 0xf2, 0xf1, 0x49, 0x63, 0xf9, 0x7e,
	0x5a, 0xf7, 0xde, 0x01, 0x96, 0x1a, 0xf3, 0xfd, 0xb4, 0xea, 0xbc, 0xfb, 0xac, 0x4d, 0x67, 0x67,
	0xb9, 0x32, 0xc9, 0xb1, 0xf4, 0x26, 0xd3, 0xe2, 0x3b, 0x94, 0x51, 0xf8, 0x5f, 0x21, 0x7e, 0x68,
	0x32, 0xcd, 0x7f, 0x69, 0x54, 0x7d, 0x82, 0xb3, 0xad, 0xd7, 0x69, 0x6c, 0xb4, 0xb6, 0x7e, 0xfa,
	0x4b, 0xcd, 0x36, 0x9c, 0xc5, 0xa1, 0x77, 0x61, 0xae, 0xf5, 0xa7, 0x31, 0xfc, 0x0f, 0xff, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x75, 0x63, 0xd6, 0x10, 0x64, 0x0c, 0x00, 0x00,
}
