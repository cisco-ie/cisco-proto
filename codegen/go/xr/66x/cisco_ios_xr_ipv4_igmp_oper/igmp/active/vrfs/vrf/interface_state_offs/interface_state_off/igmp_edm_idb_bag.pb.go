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
// source: igmp_edm_idb_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_igmp_active_vrfs_vrf_interface_state_offs_interface_state_off

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

type IgmpEdmIdbBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmIdbBag_KEYS) Reset()         { *m = IgmpEdmIdbBag_KEYS{} }
func (m *IgmpEdmIdbBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmIdbBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmIdbBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc702e9ae24d547, []int{0}
}

func (m *IgmpEdmIdbBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmIdbBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmIdbBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmIdbBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmIdbBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmIdbBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmIdbBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmIdbBag_KEYS.Size(m)
}
func (m *IgmpEdmIdbBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmIdbBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmIdbBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmIdbBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *IgmpEdmIdbBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IgmpAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpAddrtype) Reset()         { *m = IgmpAddrtype{} }
func (m *IgmpAddrtype) String() string { return proto.CompactTextString(m) }
func (*IgmpAddrtype) ProtoMessage()    {}
func (*IgmpAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc702e9ae24d547, []int{1}
}

func (m *IgmpAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpAddrtype.Unmarshal(m, b)
}
func (m *IgmpAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpAddrtype.Marshal(b, m, deterministic)
}
func (m *IgmpAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpAddrtype.Merge(m, src)
}
func (m *IgmpAddrtype) XXX_Size() int {
	return xxx_messageInfo_IgmpAddrtype.Size(m)
}
func (m *IgmpAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpAddrtype proto.InternalMessageInfo

func (m *IgmpAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IgmpAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *IgmpAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type IgmpEdmIdbBag struct {
	InterfaceNameXr              string        `protobuf:"bytes,50,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	State                        uint32        `protobuf:"varint,51,opt,name=state,proto3" json:"state,omitempty"`
	Address                      *IgmpAddrtype `protobuf:"bytes,52,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength                 uint32        `protobuf:"varint,53,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	IsInterfaceUp                bool          `protobuf:"varint,54,opt,name=is_interface_up,json=isInterfaceUp,proto3" json:"is_interface_up,omitempty"`
	IsIpEnabled                  bool          `protobuf:"varint,55,opt,name=is_ip_enabled,json=isIpEnabled,proto3" json:"is_ip_enabled,omitempty"`
	IsRouterEnabled              bool          `protobuf:"varint,56,opt,name=is_router_enabled,json=isRouterEnabled,proto3" json:"is_router_enabled,omitempty"`
	IgmpVersion                  uint32        `protobuf:"varint,57,opt,name=igmp_version,json=igmpVersion,proto3" json:"igmp_version,omitempty"`
	HostVersion                  uint32        `protobuf:"varint,58,opt,name=host_version,json=hostVersion,proto3" json:"host_version,omitempty"`
	QueryInterval                uint32        `protobuf:"varint,59,opt,name=query_interval,json=queryInterval,proto3" json:"query_interval,omitempty"`
	QueryTimeout                 uint32        `protobuf:"varint,60,opt,name=query_timeout,json=queryTimeout,proto3" json:"query_timeout,omitempty"`
	QueryMaximumResponseTime     uint32        `protobuf:"varint,61,opt,name=query_maximum_response_time,json=queryMaximumResponseTime,proto3" json:"query_maximum_response_time,omitempty"`
	LastMemberQueryInterval      uint32        `protobuf:"varint,62,opt,name=last_member_query_interval,json=lastMemberQueryInterval,proto3" json:"last_member_query_interval,omitempty"`
	GroupJoins                   uint32        `protobuf:"varint,63,opt,name=group_joins,json=groupJoins,proto3" json:"group_joins,omitempty"`
	GroupLeaves                  uint32        `protobuf:"varint,64,opt,name=group_leaves,json=groupLeaves,proto3" json:"group_leaves,omitempty"`
	IsQuerier                    bool          `protobuf:"varint,65,opt,name=is_querier,json=isQuerier,proto3" json:"is_querier,omitempty"`
	QuerierAddress               *IgmpAddrtype `protobuf:"bytes,66,opt,name=querier_address,json=querierAddress,proto3" json:"querier_address,omitempty"`
	TotalActiveGroups            uint32        `protobuf:"varint,67,opt,name=total_active_groups,json=totalActiveGroups,proto3" json:"total_active_groups,omitempty"`
	Robustness                   uint32        `protobuf:"varint,68,opt,name=robustness,proto3" json:"robustness,omitempty"`
	ProxyInterface               string        `protobuf:"bytes,69,opt,name=proxy_interface,json=proxyInterface,proto3" json:"proxy_interface,omitempty"`
	QuerierUptime                uint32        `protobuf:"varint,70,opt,name=querier_uptime,json=querierUptime,proto3" json:"querier_uptime,omitempty"`
	LasLlRegistrationCount       uint32        `protobuf:"varint,71,opt,name=las_ll_registration_count,json=lasLlRegistrationCount,proto3" json:"las_ll_registration_count,omitempty"`
	LasGetAddressCount           uint32        `protobuf:"varint,72,opt,name=las_get_address_count,json=lasGetAddressCount,proto3" json:"las_get_address_count,omitempty"`
	LasUpdateCount               uint32        `protobuf:"varint,73,opt,name=las_update_count,json=lasUpdateCount,proto3" json:"las_update_count,omitempty"`
	LasLlRemoveUpdateCount       uint32        `protobuf:"varint,74,opt,name=las_ll_remove_update_count,json=lasLlRemoveUpdateCount,proto3" json:"las_ll_remove_update_count,omitempty"`
	LasLlAddUpdateCount          uint32        `protobuf:"varint,75,opt,name=las_ll_add_update_count,json=lasLlAddUpdateCount,proto3" json:"las_ll_add_update_count,omitempty"`
	LasNullUpdateCount           uint32        `protobuf:"varint,76,opt,name=las_null_update_count,json=lasNullUpdateCount,proto3" json:"las_null_update_count,omitempty"`
	LasUnregistrationCount       uint32        `protobuf:"varint,77,opt,name=las_unregistration_count,json=lasUnregistrationCount,proto3" json:"las_unregistration_count,omitempty"`
	IsLasRequest                 bool          `protobuf:"varint,78,opt,name=is_las_request,json=isLasRequest,proto3" json:"is_las_request,omitempty"`
	IsLasRegistered              bool          `protobuf:"varint,79,opt,name=is_las_registered,json=isLasRegistered,proto3" json:"is_las_registered,omitempty"`
	VrfId                        uint32        `protobuf:"varint,80,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	MteVrfId                     uint32        `protobuf:"varint,81,opt,name=mte_vrf_id,json=mteVrfId,proto3" json:"mte_vrf_id,omitempty"`
	Location                     uint32        `protobuf:"varint,82,opt,name=location,proto3" json:"location,omitempty"`
	Mtu                          uint32        `protobuf:"varint,83,opt,name=mtu,proto3" json:"mtu,omitempty"`
	VrfState                     uint32        `protobuf:"varint,84,opt,name=vrf_state,json=vrfState,proto3" json:"vrf_state,omitempty"`
	IsConfigurationverify        bool          `protobuf:"varint,85,opt,name=is_configurationverify,json=isConfigurationverify,proto3" json:"is_configurationverify,omitempty"`
	ConfigurationvrfSet          bool          `protobuf:"varint,86,opt,name=configurationvrf_set,json=configurationvrfSet,proto3" json:"configurationvrf_set,omitempty"`
	ConfigurationvrfError        bool          `protobuf:"varint,87,opt,name=configurationvrf_error,json=configurationvrfError,proto3" json:"configurationvrf_error,omitempty"`
	ConfigurationMcastVrfSet     bool          `protobuf:"varint,88,opt,name=configuration_mcast_vrf_set,json=configurationMcastVrfSet,proto3" json:"configuration_mcast_vrf_set,omitempty"`
	ConfigurationMcastVrfError   bool          `protobuf:"varint,89,opt,name=configuration_mcast_vrf_error,json=configurationMcastVrfError,proto3" json:"configuration_mcast_vrf_error,omitempty"`
	IsImStateRegistered          bool          `protobuf:"varint,90,opt,name=is_im_state_registered,json=isImStateRegistered,proto3" json:"is_im_state_registered,omitempty"`
	IsSubscriber                 bool          `protobuf:"varint,91,opt,name=is_subscriber,json=isSubscriber,proto3" json:"is_subscriber,omitempty"`
	SubscriberMode               uint32        `protobuf:"varint,92,opt,name=subscriber_mode,json=subscriberMode,proto3" json:"subscriber_mode,omitempty"`
	IsIdentityPresent            bool          `protobuf:"varint,93,opt,name=is_identity_present,json=isIdentityPresent,proto3" json:"is_identity_present,omitempty"`
	SubscriberAddress            *IgmpAddrtype `protobuf:"bytes,94,opt,name=subscriber_address,json=subscriberAddress,proto3" json:"subscriber_address,omitempty"`
	SubscriberId                 string        `protobuf:"bytes,95,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	ParentIfhandle               string        `protobuf:"bytes,96,opt,name=parent_ifhandle,json=parentIfhandle,proto3" json:"parent_ifhandle,omitempty"`
	TimeSinceLastQueryInSeconds  uint32        `protobuf:"varint,97,opt,name=time_since_last_query_in_seconds,json=timeSinceLastQueryInSeconds,proto3" json:"time_since_last_query_in_seconds,omitempty"`
	TimeSinceLastReportInSeconds uint32        `protobuf:"varint,98,opt,name=time_since_last_report_in_seconds,json=timeSinceLastReportInSeconds,proto3" json:"time_since_last_report_in_seconds,omitempty"`
	RouterUptimeInSeconds        uint32        `protobuf:"varint,99,opt,name=router_uptime_in_seconds,json=routerUptimeInSeconds,proto3" json:"router_uptime_in_seconds,omitempty"`
	MteTupleCount                uint32        `protobuf:"varint,100,opt,name=mte_tuple_count,json=mteTupleCount,proto3" json:"mte_tuple_count,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}      `json:"-"`
	XXX_unrecognized             []byte        `json:"-"`
	XXX_sizecache                int32         `json:"-"`
}

func (m *IgmpEdmIdbBag) Reset()         { *m = IgmpEdmIdbBag{} }
func (m *IgmpEdmIdbBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmIdbBag) ProtoMessage()    {}
func (*IgmpEdmIdbBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc702e9ae24d547, []int{2}
}

func (m *IgmpEdmIdbBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmIdbBag.Unmarshal(m, b)
}
func (m *IgmpEdmIdbBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmIdbBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmIdbBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmIdbBag.Merge(m, src)
}
func (m *IgmpEdmIdbBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmIdbBag.Size(m)
}
func (m *IgmpEdmIdbBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmIdbBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmIdbBag proto.InternalMessageInfo

func (m *IgmpEdmIdbBag) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *IgmpEdmIdbBag) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetAddress() *IgmpAddrtype {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *IgmpEdmIdbBag) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetIsInterfaceUp() bool {
	if m != nil {
		return m.IsInterfaceUp
	}
	return false
}

func (m *IgmpEdmIdbBag) GetIsIpEnabled() bool {
	if m != nil {
		return m.IsIpEnabled
	}
	return false
}

func (m *IgmpEdmIdbBag) GetIsRouterEnabled() bool {
	if m != nil {
		return m.IsRouterEnabled
	}
	return false
}

func (m *IgmpEdmIdbBag) GetIgmpVersion() uint32 {
	if m != nil {
		return m.IgmpVersion
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetHostVersion() uint32 {
	if m != nil {
		return m.HostVersion
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetQueryInterval() uint32 {
	if m != nil {
		return m.QueryInterval
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetQueryTimeout() uint32 {
	if m != nil {
		return m.QueryTimeout
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetQueryMaximumResponseTime() uint32 {
	if m != nil {
		return m.QueryMaximumResponseTime
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetLastMemberQueryInterval() uint32 {
	if m != nil {
		return m.LastMemberQueryInterval
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetGroupJoins() uint32 {
	if m != nil {
		return m.GroupJoins
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetGroupLeaves() uint32 {
	if m != nil {
		return m.GroupLeaves
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetIsQuerier() bool {
	if m != nil {
		return m.IsQuerier
	}
	return false
}

func (m *IgmpEdmIdbBag) GetQuerierAddress() *IgmpAddrtype {
	if m != nil {
		return m.QuerierAddress
	}
	return nil
}

func (m *IgmpEdmIdbBag) GetTotalActiveGroups() uint32 {
	if m != nil {
		return m.TotalActiveGroups
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetRobustness() uint32 {
	if m != nil {
		return m.Robustness
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetProxyInterface() string {
	if m != nil {
		return m.ProxyInterface
	}
	return ""
}

func (m *IgmpEdmIdbBag) GetQuerierUptime() uint32 {
	if m != nil {
		return m.QuerierUptime
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetLasLlRegistrationCount() uint32 {
	if m != nil {
		return m.LasLlRegistrationCount
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetLasGetAddressCount() uint32 {
	if m != nil {
		return m.LasGetAddressCount
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetLasUpdateCount() uint32 {
	if m != nil {
		return m.LasUpdateCount
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetLasLlRemoveUpdateCount() uint32 {
	if m != nil {
		return m.LasLlRemoveUpdateCount
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetLasLlAddUpdateCount() uint32 {
	if m != nil {
		return m.LasLlAddUpdateCount
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetLasNullUpdateCount() uint32 {
	if m != nil {
		return m.LasNullUpdateCount
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetLasUnregistrationCount() uint32 {
	if m != nil {
		return m.LasUnregistrationCount
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetIsLasRequest() bool {
	if m != nil {
		return m.IsLasRequest
	}
	return false
}

func (m *IgmpEdmIdbBag) GetIsLasRegistered() bool {
	if m != nil {
		return m.IsLasRegistered
	}
	return false
}

func (m *IgmpEdmIdbBag) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetMteVrfId() uint32 {
	if m != nil {
		return m.MteVrfId
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetLocation() uint32 {
	if m != nil {
		return m.Location
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetVrfState() uint32 {
	if m != nil {
		return m.VrfState
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetIsConfigurationverify() bool {
	if m != nil {
		return m.IsConfigurationverify
	}
	return false
}

func (m *IgmpEdmIdbBag) GetConfigurationvrfSet() bool {
	if m != nil {
		return m.ConfigurationvrfSet
	}
	return false
}

func (m *IgmpEdmIdbBag) GetConfigurationvrfError() bool {
	if m != nil {
		return m.ConfigurationvrfError
	}
	return false
}

func (m *IgmpEdmIdbBag) GetConfigurationMcastVrfSet() bool {
	if m != nil {
		return m.ConfigurationMcastVrfSet
	}
	return false
}

func (m *IgmpEdmIdbBag) GetConfigurationMcastVrfError() bool {
	if m != nil {
		return m.ConfigurationMcastVrfError
	}
	return false
}

func (m *IgmpEdmIdbBag) GetIsImStateRegistered() bool {
	if m != nil {
		return m.IsImStateRegistered
	}
	return false
}

func (m *IgmpEdmIdbBag) GetIsSubscriber() bool {
	if m != nil {
		return m.IsSubscriber
	}
	return false
}

func (m *IgmpEdmIdbBag) GetSubscriberMode() uint32 {
	if m != nil {
		return m.SubscriberMode
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetIsIdentityPresent() bool {
	if m != nil {
		return m.IsIdentityPresent
	}
	return false
}

func (m *IgmpEdmIdbBag) GetSubscriberAddress() *IgmpAddrtype {
	if m != nil {
		return m.SubscriberAddress
	}
	return nil
}

func (m *IgmpEdmIdbBag) GetSubscriberId() string {
	if m != nil {
		return m.SubscriberId
	}
	return ""
}

func (m *IgmpEdmIdbBag) GetParentIfhandle() string {
	if m != nil {
		return m.ParentIfhandle
	}
	return ""
}

func (m *IgmpEdmIdbBag) GetTimeSinceLastQueryInSeconds() uint32 {
	if m != nil {
		return m.TimeSinceLastQueryInSeconds
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetTimeSinceLastReportInSeconds() uint32 {
	if m != nil {
		return m.TimeSinceLastReportInSeconds
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetRouterUptimeInSeconds() uint32 {
	if m != nil {
		return m.RouterUptimeInSeconds
	}
	return 0
}

func (m *IgmpEdmIdbBag) GetMteTupleCount() uint32 {
	if m != nil {
		return m.MteTupleCount
	}
	return 0
}

func init() {
	proto.RegisterType((*IgmpEdmIdbBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.vrfs.vrf.interface_state_offs.interface_state_off.igmp_edm_idb_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.vrfs.vrf.interface_state_offs.interface_state_off.igmp_addrtype")
	proto.RegisterType((*IgmpEdmIdbBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.vrfs.vrf.interface_state_offs.interface_state_off.igmp_edm_idb_bag")
}

func init() { proto.RegisterFile("igmp_edm_idb_bag.proto", fileDescriptor_bfc702e9ae24d547) }

var fileDescriptor_bfc702e9ae24d547 = []byte{
	// 1204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0x5b, 0x73, 0xdb, 0x36,
	0x13, 0x86, 0x47, 0x5f, 0xe6, 0x4b, 0x1c, 0xd8, 0x92, 0x62, 0x24, 0x4e, 0x90, 0x53, 0xeb, 0x38,
	0x4d, 0xea, 0xe9, 0x85, 0x66, 0x72, 0x3e, 0x35, 0x6d, 0x55, 0xd7, 0x75, 0x95, 0xc8, 0x6e, 0x2c,
	0xd9, 0x6e, 0xdc, 0x13, 0x0a, 0x89, 0x4b, 0x19, 0x1d, 0x92, 0x60, 0x00, 0x90, 0x63, 0xdf, 0xf6,
	0x1f, 0x74, 0x7a, 0xd1, 0xbf, 0xdb, 0xc1, 0x82, 0x14, 0x45, 0xc5, 0xbd, 0x6c, 0x6e, 0x34, 0xd2,
	0xbb, 0xcf, 0xbb, 0x58, 0x2c, 0xc1, 0x85, 0xc8, 0x65, 0x39, 0x89, 0x53, 0x0e, 0x41, 0xcc, 0x65,
	0x30, 0xe2, 0x23, 0x31, 0xe9, 0xa4, 0x5a, 0x59, 0x45, 0x0f, 0xc7, 0xd2, 0x8c, 0x15, 0x97, 0xca,
	0xf0, 0x63, 0xcd, 0x65, 0x9a, 0x3f, 0xe4, 0x48, 0xaa, 0x14, 0x74, 0xc7, 0x7d, 0xeb, 0x88, 0xb1,
	0x95, 0x39, 0x74, 0x72, 0x1d, 0x1a, 0xf7, 0xd1, 0x91, 0x89, 0x05, 0x1d, 0x8a, 0x31, 0x70, 0x63,
	0x85, 0x05, 0xae, 0xc2, 0xd0, 0x9c, 0x26, 0xae, 0x1d, 0x92, 0x95, 0xf9, 0x45, 0xf9, 0xeb, 0xcd,
	0xc3, 0x21, 0xbd, 0x4a, 0x16, 0x72, 0x1d, 0xf2, 0x44, 0xc4, 0xc0, 0x1a, 0xab, 0x8d, 0xf5, 0xf3,
	0x83, 0x73, 0xb9, 0x0e, 0x77, 0x44, 0x0c, 0xf4, 0x0e, 0x69, 0x55, 0xa9, 0x10, 0xf8, 0x1f, 0x02,
	0xcd, 0xa9, 0xea, 0xb0, 0xb5, 0x84, 0x34, 0x31, 0xb5, 0x08, 0x02, 0x6d, 0x4f, 0x52, 0xa0, 0x57,
	0xc8, 0x39, 0x51, 0xcb, 0x78, 0x56, 0xf8, 0x84, 0xb7, 0xc8, 0x12, 0x6e, 0xca, 0x91, 0x60, 0x4c,
	0x91, 0x6e, 0xd1, 0x69, 0x5d, 0x2f, 0x15, 0xc8, 0xe3, 0x29, 0x72, 0x66, 0x8a, 0x3c, 0x2e, 0x90,
	0xb5, 0xbf, 0x2e, 0x91, 0x0b, 0xf3, 0x7b, 0xa1, 0x9f, 0x91, 0xe5, 0x7a, 0xad, 0xfc, 0x58, 0xb3,
	0xfb, 0x68, 0x6e, 0xd7, 0xca, 0x7d, 0xab, 0xe9, 0x25, 0xf2, 0x7f, 0x6c, 0x0c, 0x7b, 0xb0, 0xda,
	0x58, 0x6f, 0x0e, 0xfc, 0x0f, 0xfa, 0x47, 0x83, 0x9c, 0x2b, 0x57, 0x7d, 0xb8, 0xda, 0x58, 0x5f,
	0xbc, 0x7f, 0xd4, 0xf9, 0xcf, 0x9e, 0x47, 0xa7, 0xd6, 0xb1, 0x41, 0xb9, 0x30, 0xbd, 0x4d, 0x9a,
	0xa9, 0x86, 0x50, 0x1e, 0xf3, 0x08, 0x92, 0x89, 0x3d, 0x62, 0x8f, 0xb0, 0xc4, 0x25, 0x2f, 0xf6,
	0x51, 0xa3, 0x77, 0x49, 0x5b, 0x1a, 0x5e, 0x65, 0xcd, 0x52, 0xf6, 0x78, 0xb5, 0xb1, 0xbe, 0x30,
	0x68, 0x4a, 0xd3, 0x2b, 0xd5, 0xfd, 0x94, 0xae, 0x91, 0xa6, 0xe3, 0x52, 0x0e, 0x89, 0x18, 0x45,
	0x10, 0xb0, 0x27, 0x48, 0x2d, 0x4a, 0xd3, 0x4b, 0x37, 0xbd, 0x84, 0x7d, 0x33, 0x5c, 0xab, 0xcc,
	0x82, 0x9e, 0x72, 0x4f, 0x91, 0x6b, 0x4b, 0x33, 0x40, 0xbd, 0x64, 0xdd, 0xb3, 0x71, 0x65, 0xe7,
	0xa0, 0x8d, 0x54, 0x09, 0x7b, 0x86, 0xb5, 0x2d, 0x3a, 0xed, 0xc0, 0x4b, 0x0e, 0x39, 0x52, 0xc6,
	0x4e, 0x91, 0xe7, 0x1e, 0x71, 0x5a, 0x89, 0xdc, 0x21, 0xad, 0x77, 0x19, 0xe8, 0x13, 0xbf, 0x81,
	0x5c, 0x44, 0xec, 0x05, 0x42, 0x4d, 0x54, 0x7b, 0x85, 0xe8, 0x3a, 0xe1, 0x31, 0x2b, 0x63, 0x50,
	0x99, 0x65, 0x9f, 0xfb, 0x4e, 0xa0, 0xb8, 0xe7, 0x35, 0xfa, 0x92, 0x5c, 0xf7, 0x50, 0x2c, 0x8e,
	0x65, 0x9c, 0xc5, 0x5c, 0x83, 0x49, 0x55, 0x62, 0x00, 0x5d, 0xec, 0x25, 0x5a, 0x18, 0x22, 0xdb,
	0x9e, 0x18, 0x14, 0x80, 0xcb, 0x40, 0x5f, 0x90, 0x6b, 0x91, 0x30, 0x96, 0xc7, 0x10, 0x8f, 0x40,
	0xf3, 0xb9, 0xb2, 0xbe, 0x40, 0xf7, 0x15, 0x47, 0x6c, 0x23, 0xb0, 0x5b, 0x2b, 0xf0, 0x63, 0xb2,
	0x38, 0xd1, 0x2a, 0x4b, 0xf9, 0xef, 0x4a, 0x26, 0x86, 0x7d, 0x89, 0x34, 0x41, 0xe9, 0x95, 0x53,
	0x5c, 0x2f, 0x3c, 0x10, 0x81, 0xc8, 0xc1, 0xb0, 0xaf, 0x7c, 0x2f, 0x50, 0xeb, 0xa3, 0x44, 0x6f,
	0x12, 0x22, 0x0d, 0xae, 0x2b, 0x41, 0xb3, 0x2e, 0xb6, 0xfd, 0xbc, 0x34, 0xbb, 0x5e, 0xa0, 0x7f,
	0x36, 0x48, 0xbb, 0x08, 0x4e, 0x5f, 0x88, 0xaf, 0x3f, 0xf0, 0xd1, 0x6c, 0x15, 0x05, 0x94, 0x2f,
	0x68, 0x87, 0x5c, 0xb4, 0xca, 0x8a, 0x88, 0xfb, 0xfc, 0x1c, 0xb7, 0x63, 0xd8, 0x06, 0x6e, 0x6e,
	0x19, 0x43, 0x5d, 0x8c, 0x6c, 0x61, 0x80, 0x7e, 0x44, 0x88, 0x56, 0xa3, 0xcc, 0xd8, 0xc4, 0x55,
	0xff, 0x8d, 0xef, 0x52, 0xa5, 0xd0, 0x4f, 0x49, 0x3b, 0xd5, 0xea, 0xf8, 0xa4, 0x3a, 0xcf, 0x6c,
	0x13, 0x5f, 0xdb, 0x16, 0xca, 0xd3, 0xf3, 0x5c, 0x9e, 0x1b, 0xd7, 0x8b, 0x2c, 0xc5, 0xc7, 0xfb,
	0x6d, 0x75, 0x6e, 0x24, 0xe8, 0x7d, 0x14, 0xe9, 0x33, 0x72, 0x35, 0x12, 0x86, 0x47, 0x11, 0xd7,
	0x30, 0x91, 0xc6, 0x6a, 0x61, 0xa5, 0x4a, 0xf8, 0x58, 0x65, 0x89, 0x65, 0x5b, 0xe8, 0xb8, 0x1c,
	0x09, 0xd3, 0x8f, 0x06, 0x33, 0xe1, 0x0d, 0x17, 0xa5, 0xf7, 0xc8, 0x8a, 0xb3, 0x4e, 0xc0, 0x96,
	0xdd, 0x2e, 0x6c, 0xdf, 0xa1, 0x8d, 0x46, 0xc2, 0x6c, 0x81, 0x2d, 0x1a, 0xe1, 0x2d, 0xeb, 0xe4,
	0x82, 0xb3, 0x64, 0x69, 0xe0, 0xba, 0xe7, 0xe9, 0x1e, 0xd2, 0xad, 0x48, 0x98, 0x7d, 0x94, 0x3d,
	0xf9, 0x1c, 0xcf, 0x9a, 0xaf, 0x2b, 0x56, 0x39, 0xd4, 0x3d, 0xaf, 0x6a, 0x85, 0xb9, 0xf8, 0xac,
	0xf7, 0x21, 0xb9, 0x52, 0x78, 0x45, 0x10, 0xd4, 0x8d, 0xaf, 0xd1, 0x78, 0x11, 0x8d, 0xdd, 0x20,
	0x98, 0x75, 0x15, 0xdb, 0x49, 0xb2, 0x28, 0xaa, 0x7b, 0xfa, 0xd3, 0xed, 0xec, 0x64, 0x51, 0x34,
	0x6b, 0x79, 0x4a, 0x18, 0x6e, 0x27, 0x39, 0xa5, 0x77, 0xdb, 0xd3, 0x12, 0xf7, 0x6b, 0x61, 0xef,
	0xfc, 0x84, 0xb4, 0xa4, 0xe1, 0xce, 0xac, 0xe1, 0x5d, 0x06, 0xc6, 0xb2, 0x1d, 0x3c, 0xcd, 0x4b,
	0xd2, 0xf4, 0x85, 0x19, 0x78, 0xad, 0x98, 0x36, 0x9e, 0x72, 0x19, 0x40, 0x43, 0xc0, 0xbe, 0x2f,
	0xa7, 0x0d, 0x82, 0xa5, 0x4c, 0x57, 0xc8, 0x59, 0x77, 0x31, 0xc9, 0x80, 0xbd, 0xf1, 0x63, 0x3a,
	0xd7, 0x61, 0x2f, 0xa0, 0x37, 0x08, 0x89, 0x2d, 0xf0, 0x22, 0xb4, 0x8b, 0xa1, 0x85, 0xd8, 0xc2,
	0x01, 0x46, 0xaf, 0x91, 0x85, 0x48, 0x8d, 0xb1, 0x2e, 0x36, 0xf0, 0xb1, 0xf2, 0x37, 0xbd, 0x40,
	0xce, 0xc4, 0x36, 0x63, 0x43, 0x94, 0xdd, 0x57, 0x7a, 0x9d, 0x9c, 0x77, 0x79, 0xfc, 0x65, 0xb0,
	0xe7, 0xf1, 0x5c, 0x87, 0x43, 0xbc, 0x0f, 0x1e, 0x91, 0xcb, 0xd2, 0x1d, 0x80, 0x24, 0x94, 0x93,
	0xcc, 0x6f, 0x35, 0x07, 0x2d, 0xc3, 0x13, 0xb6, 0x8f, 0x05, 0xaf, 0x48, 0xb3, 0xf1, 0x7e, 0x90,
	0xde, 0x23, 0x97, 0xea, 0x1e, 0xb7, 0x00, 0x58, 0x76, 0x80, 0xa6, 0x8b, 0xf3, 0xb1, 0x21, 0x58,
	0xb7, 0xd2, 0x7b, 0x16, 0xd0, 0x5a, 0x69, 0xf6, 0x83, 0x5f, 0x69, 0x3e, 0xba, 0xe9, 0x82, 0x6e,
	0xf8, 0xd5, 0x02, 0x3c, 0x1e, 0xbb, 0x61, 0x56, 0x2e, 0xf8, 0x16, 0xbd, 0xac, 0x86, 0x6c, 0x3b,
	0xe2, 0xc0, 0xaf, 0xda, 0x25, 0x37, 0xff, 0xcd, 0xee, 0x17, 0x3f, 0xc4, 0x04, 0xd7, 0x4e, 0x4d,
	0xe0, 0x2b, 0x78, 0x80, 0x2d, 0x92, 0x71, 0x31, 0x3c, 0x66, 0x9e, 0xe9, 0x8f, 0x7e, 0xb7, 0xd2,
	0xf4, 0x62, 0xec, 0xe6, 0xcc, 0x73, 0xbd, 0x8d, 0xb7, 0x92, 0xc9, 0x46, 0x66, 0xac, 0xe5, 0x08,
	0x34, 0xfb, 0xa9, 0x3c, 0x28, 0xc3, 0xa9, 0xe6, 0xa6, 0x42, 0x45, 0xf0, 0x58, 0x05, 0xc0, 0x7e,
	0xf6, 0xaf, 0x55, 0x25, 0x6f, 0xab, 0x00, 0xdc, 0x38, 0x72, 0x25, 0x04, 0x90, 0x58, 0x69, 0x4f,
	0x78, 0xaa, 0xc1, 0x40, 0x62, 0xd9, 0x2f, 0x98, 0x73, 0x59, 0x9a, 0x5e, 0x11, 0x79, 0xe3, 0x03,
	0xf4, 0xef, 0x06, 0xa1, 0x33, 0x99, 0xcb, 0xa9, 0xfa, 0xeb, 0x07, 0x9e, 0xaa, 0xcb, 0x55, 0x0d,
	0xdd, 0xea, 0xea, 0x9f, 0x29, 0x4c, 0x06, 0x8c, 0xe3, 0x18, 0x5c, 0xaa, 0xc4, 0x5e, 0x80, 0xd3,
	0x52, 0x68, 0x48, 0x2c, 0x97, 0xe1, 0x91, 0x48, 0x82, 0x08, 0xd8, 0x6f, 0xc5, 0xb4, 0x44, 0xb9,
	0x57, 0xa8, 0x74, 0x93, 0xac, 0xba, 0x71, 0xc8, 0x8d, 0x4c, 0xc6, 0xc0, 0xf1, 0x96, 0x2b, 0xaf,
	0x37, 0x6e, 0x60, 0xac, 0x92, 0xc0, 0x30, 0x81, 0x1d, 0xbd, 0xee, 0xb8, 0xa1, 0xc3, 0xfa, 0xc2,
	0xd8, 0xe2, 0x8e, 0x1b, 0x7a, 0x84, 0x6e, 0x91, 0x5b, 0xf3, 0x69, 0x34, 0xa4, 0x4a, 0xdb, 0xd9,
	0x3c, 0x23, 0xcc, 0x73, 0xa3, 0x96, 0x67, 0x80, 0x54, 0x95, 0xe8, 0x09, 0x61, 0xc5, 0x9f, 0x0c,
	0x3f, 0xbc, 0x67, 0xfd, 0x63, 0xf4, 0xaf, 0xf8, 0xb8, 0x1f, 0xe3, 0x95, 0xf1, 0x2e, 0x69, 0xbb,
	0xf7, 0xdd, 0x66, 0x69, 0x54, 0xce, 0xaf, 0xc0, 0xcf, 0xfd, 0xd8, 0xc2, 0x9e, 0x53, 0x71, 0x00,
	0x8d, 0xce, 0xe2, 0x5f, 0xe8, 0x07, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x47, 0x2b, 0xaf, 0x19,
	0x5c, 0x0b, 0x00, 0x00,
}
