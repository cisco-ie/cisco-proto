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
// source: igmp_edm_idb_ifrs_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_mld_active_default_context_ifrs_interfaces_ifrs_interface

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

type IgmpEdmIdbIfrsBag_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmIdbIfrsBag_KEYS) Reset()         { *m = IgmpEdmIdbIfrsBag_KEYS{} }
func (m *IgmpEdmIdbIfrsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmIdbIfrsBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmIdbIfrsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd9ff375042b5111, []int{0}
}

func (m *IgmpEdmIdbIfrsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmIdbIfrsBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmIdbIfrsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmIdbIfrsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmIdbIfrsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmIdbIfrsBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmIdbIfrsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmIdbIfrsBag_KEYS.Size(m)
}
func (m *IgmpEdmIdbIfrsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmIdbIfrsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmIdbIfrsBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmIdbIfrsBag_KEYS) GetInterfaceName() string {
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
	return fileDescriptor_cd9ff375042b5111, []int{1}
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
	InterfaceNameXr              string        `protobuf:"bytes,1,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	State                        uint32        `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
	Address                      *IgmpAddrtype `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength                 uint32        `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	IsInterfaceUp                bool          `protobuf:"varint,5,opt,name=is_interface_up,json=isInterfaceUp,proto3" json:"is_interface_up,omitempty"`
	IsIpEnabled                  bool          `protobuf:"varint,6,opt,name=is_ip_enabled,json=isIpEnabled,proto3" json:"is_ip_enabled,omitempty"`
	IsRouterEnabled              bool          `protobuf:"varint,7,opt,name=is_router_enabled,json=isRouterEnabled,proto3" json:"is_router_enabled,omitempty"`
	IgmpVersion                  uint32        `protobuf:"varint,8,opt,name=igmp_version,json=igmpVersion,proto3" json:"igmp_version,omitempty"`
	HostVersion                  uint32        `protobuf:"varint,9,opt,name=host_version,json=hostVersion,proto3" json:"host_version,omitempty"`
	QueryInterval                uint32        `protobuf:"varint,10,opt,name=query_interval,json=queryInterval,proto3" json:"query_interval,omitempty"`
	QueryTimeout                 uint32        `protobuf:"varint,11,opt,name=query_timeout,json=queryTimeout,proto3" json:"query_timeout,omitempty"`
	QueryMaximumResponseTime     uint32        `protobuf:"varint,12,opt,name=query_maximum_response_time,json=queryMaximumResponseTime,proto3" json:"query_maximum_response_time,omitempty"`
	LastMemberQueryInterval      uint32        `protobuf:"varint,13,opt,name=last_member_query_interval,json=lastMemberQueryInterval,proto3" json:"last_member_query_interval,omitempty"`
	GroupJoins                   uint32        `protobuf:"varint,14,opt,name=group_joins,json=groupJoins,proto3" json:"group_joins,omitempty"`
	GroupLeaves                  uint32        `protobuf:"varint,15,opt,name=group_leaves,json=groupLeaves,proto3" json:"group_leaves,omitempty"`
	IsQuerier                    bool          `protobuf:"varint,16,opt,name=is_querier,json=isQuerier,proto3" json:"is_querier,omitempty"`
	QuerierAddress               *IgmpAddrtype `protobuf:"bytes,17,opt,name=querier_address,json=querierAddress,proto3" json:"querier_address,omitempty"`
	TotalActiveGroups            uint32        `protobuf:"varint,18,opt,name=total_active_groups,json=totalActiveGroups,proto3" json:"total_active_groups,omitempty"`
	Robustness                   uint32        `protobuf:"varint,19,opt,name=robustness,proto3" json:"robustness,omitempty"`
	ProxyInterface               string        `protobuf:"bytes,20,opt,name=proxy_interface,json=proxyInterface,proto3" json:"proxy_interface,omitempty"`
	QuerierUptime                uint32        `protobuf:"varint,21,opt,name=querier_uptime,json=querierUptime,proto3" json:"querier_uptime,omitempty"`
	LasLlRegistrationCount       uint32        `protobuf:"varint,22,opt,name=las_ll_registration_count,json=lasLlRegistrationCount,proto3" json:"las_ll_registration_count,omitempty"`
	LasGetAddressCount           uint32        `protobuf:"varint,23,opt,name=las_get_address_count,json=lasGetAddressCount,proto3" json:"las_get_address_count,omitempty"`
	LasUpdateCount               uint32        `protobuf:"varint,24,opt,name=las_update_count,json=lasUpdateCount,proto3" json:"las_update_count,omitempty"`
	LasLlRemoveUpdateCount       uint32        `protobuf:"varint,25,opt,name=las_ll_remove_update_count,json=lasLlRemoveUpdateCount,proto3" json:"las_ll_remove_update_count,omitempty"`
	LasLlAddUpdateCount          uint32        `protobuf:"varint,26,opt,name=las_ll_add_update_count,json=lasLlAddUpdateCount,proto3" json:"las_ll_add_update_count,omitempty"`
	LasNullUpdateCount           uint32        `protobuf:"varint,27,opt,name=las_null_update_count,json=lasNullUpdateCount,proto3" json:"las_null_update_count,omitempty"`
	LasUnregistrationCount       uint32        `protobuf:"varint,28,opt,name=las_unregistration_count,json=lasUnregistrationCount,proto3" json:"las_unregistration_count,omitempty"`
	IsLasRequest                 bool          `protobuf:"varint,29,opt,name=is_las_request,json=isLasRequest,proto3" json:"is_las_request,omitempty"`
	IsLasRegistered              bool          `protobuf:"varint,30,opt,name=is_las_registered,json=isLasRegistered,proto3" json:"is_las_registered,omitempty"`
	VrfId                        uint32        `protobuf:"varint,31,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	MteVrfId                     uint32        `protobuf:"varint,32,opt,name=mte_vrf_id,json=mteVrfId,proto3" json:"mte_vrf_id,omitempty"`
	Location                     uint32        `protobuf:"varint,33,opt,name=location,proto3" json:"location,omitempty"`
	Mtu                          uint32        `protobuf:"varint,34,opt,name=mtu,proto3" json:"mtu,omitempty"`
	VrfState                     uint32        `protobuf:"varint,35,opt,name=vrf_state,json=vrfState,proto3" json:"vrf_state,omitempty"`
	IsConfigurationverify        bool          `protobuf:"varint,36,opt,name=is_configurationverify,json=isConfigurationverify,proto3" json:"is_configurationverify,omitempty"`
	ConfigurationvrfSet          bool          `protobuf:"varint,37,opt,name=configurationvrf_set,json=configurationvrfSet,proto3" json:"configurationvrf_set,omitempty"`
	ConfigurationvrfError        bool          `protobuf:"varint,38,opt,name=configurationvrf_error,json=configurationvrfError,proto3" json:"configurationvrf_error,omitempty"`
	ConfigurationMcastVrfSet     bool          `protobuf:"varint,39,opt,name=configuration_mcast_vrf_set,json=configurationMcastVrfSet,proto3" json:"configuration_mcast_vrf_set,omitempty"`
	ConfigurationMcastVrfError   bool          `protobuf:"varint,40,opt,name=configuration_mcast_vrf_error,json=configurationMcastVrfError,proto3" json:"configuration_mcast_vrf_error,omitempty"`
	IsImStateRegistered          bool          `protobuf:"varint,41,opt,name=is_im_state_registered,json=isImStateRegistered,proto3" json:"is_im_state_registered,omitempty"`
	IsSubscriber                 bool          `protobuf:"varint,42,opt,name=is_subscriber,json=isSubscriber,proto3" json:"is_subscriber,omitempty"`
	SubscriberMode               uint32        `protobuf:"varint,43,opt,name=subscriber_mode,json=subscriberMode,proto3" json:"subscriber_mode,omitempty"`
	IsIdentityPresent            bool          `protobuf:"varint,44,opt,name=is_identity_present,json=isIdentityPresent,proto3" json:"is_identity_present,omitempty"`
	SubscriberAddress            *IgmpAddrtype `protobuf:"bytes,45,opt,name=subscriber_address,json=subscriberAddress,proto3" json:"subscriber_address,omitempty"`
	SubscriberId                 string        `protobuf:"bytes,46,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	ParentIfhandle               string        `protobuf:"bytes,47,opt,name=parent_ifhandle,json=parentIfhandle,proto3" json:"parent_ifhandle,omitempty"`
	TimeSinceLastQueryInSeconds  uint32        `protobuf:"varint,48,opt,name=time_since_last_query_in_seconds,json=timeSinceLastQueryInSeconds,proto3" json:"time_since_last_query_in_seconds,omitempty"`
	TimeSinceLastReportInSeconds uint32        `protobuf:"varint,49,opt,name=time_since_last_report_in_seconds,json=timeSinceLastReportInSeconds,proto3" json:"time_since_last_report_in_seconds,omitempty"`
	RouterUptimeInSeconds        uint32        `protobuf:"varint,50,opt,name=router_uptime_in_seconds,json=routerUptimeInSeconds,proto3" json:"router_uptime_in_seconds,omitempty"`
	MteTupleCount                uint32        `protobuf:"varint,51,opt,name=mte_tuple_count,json=mteTupleCount,proto3" json:"mte_tuple_count,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}      `json:"-"`
	XXX_unrecognized             []byte        `json:"-"`
	XXX_sizecache                int32         `json:"-"`
}

func (m *IgmpEdmIdbBag) Reset()         { *m = IgmpEdmIdbBag{} }
func (m *IgmpEdmIdbBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmIdbBag) ProtoMessage()    {}
func (*IgmpEdmIdbBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd9ff375042b5111, []int{2}
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

type IgmpEdmIdbIfrsBag struct {
	IgmpInterfaceEntry   *IgmpEdmIdbBag `protobuf:"bytes,50,opt,name=igmp_interface_entry,json=igmpInterfaceEntry,proto3" json:"igmp_interface_entry,omitempty"`
	JoinGroupCount       uint32         `protobuf:"varint,51,opt,name=join_group_count,json=joinGroupCount,proto3" json:"join_group_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IgmpEdmIdbIfrsBag) Reset()         { *m = IgmpEdmIdbIfrsBag{} }
func (m *IgmpEdmIdbIfrsBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmIdbIfrsBag) ProtoMessage()    {}
func (*IgmpEdmIdbIfrsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd9ff375042b5111, []int{3}
}

func (m *IgmpEdmIdbIfrsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmIdbIfrsBag.Unmarshal(m, b)
}
func (m *IgmpEdmIdbIfrsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmIdbIfrsBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmIdbIfrsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmIdbIfrsBag.Merge(m, src)
}
func (m *IgmpEdmIdbIfrsBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmIdbIfrsBag.Size(m)
}
func (m *IgmpEdmIdbIfrsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmIdbIfrsBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmIdbIfrsBag proto.InternalMessageInfo

func (m *IgmpEdmIdbIfrsBag) GetIgmpInterfaceEntry() *IgmpEdmIdbBag {
	if m != nil {
		return m.IgmpInterfaceEntry
	}
	return nil
}

func (m *IgmpEdmIdbIfrsBag) GetJoinGroupCount() uint32 {
	if m != nil {
		return m.JoinGroupCount
	}
	return 0
}

func init() {
	proto.RegisterType((*IgmpEdmIdbIfrsBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.ifrs_interfaces.ifrs_interface.igmp_edm_idb_ifrs_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.ifrs_interfaces.ifrs_interface.igmp_addrtype")
	proto.RegisterType((*IgmpEdmIdbBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.ifrs_interfaces.ifrs_interface.igmp_edm_idb_bag")
	proto.RegisterType((*IgmpEdmIdbIfrsBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.ifrs_interfaces.ifrs_interface.igmp_edm_idb_ifrs_bag")
}

func init() { proto.RegisterFile("igmp_edm_idb_ifrs_bag.proto", fileDescriptor_cd9ff375042b5111) }

var fileDescriptor_cd9ff375042b5111 = []byte{
	// 1242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4b, 0x73, 0xdb, 0xb6,
	0x13, 0x1f, 0xfd, 0xf3, 0x8f, 0x63, 0xaf, 0x2c, 0xc9, 0x86, 0xed, 0x18, 0xb1, 0xf3, 0x70, 0x9c,
	0x97, 0x9b, 0xb6, 0x6a, 0xf3, 0x68, 0xfa, 0x9a, 0x1e, 0x3c, 0x19, 0x4f, 0x46, 0xad, 0x9d, 0x69,
	0xe4, 0x38, 0xd3, 0x9e, 0x30, 0x90, 0xb8, 0x52, 0xd0, 0x21, 0x09, 0x06, 0x00, 0x35, 0xf6, 0xa5,
	0xc7, 0xde, 0x7a, 0xea, 0xa9, 0x1f, 0xae, 0xdf, 0xa5, 0x83, 0x05, 0x49, 0x91, 0x8e, 0x7b, 0x4c,
	0x6e, 0xd2, 0xef, 0xb1, 0xbb, 0x58, 0x82, 0xbb, 0x84, 0x6d, 0x35, 0x4d, 0x32, 0x81, 0x51, 0x22,
	0x54, 0x34, 0x12, 0x6a, 0x62, 0xac, 0x18, 0xc9, 0x69, 0x3f, 0x33, 0xda, 0x69, 0x76, 0x32, 0x56,
	0x76, 0xac, 0x85, 0xd2, 0x56, 0x9c, 0x1a, 0xa1, 0xb2, 0xd9, 0x53, 0x41, 0x72, 0x9d, 0xa1, 0xe9,
	0x27, 0x71, 0xd4, 0x97, 0x63, 0xa7, 0x66, 0xd8, 0x8f, 0x70, 0x22, 0xf3, 0xd8, 0x89, 0xb1, 0x4e,
	0x1d, 0x9e, 0xba, 0x3e, 0x85, 0x51, 0xa9, 0x43, 0x33, 0x91, 0x63, 0xb4, 0xe7, 0xfe, 0xef, 0x3e,
	0x87, 0xad, 0x0b, 0xb3, 0x8a, 0x9f, 0x0e, 0x7e, 0x3d, 0x66, 0xf7, 0xa0, 0x5b, 0x49, 0x45, 0x2a,
	0x13, 0xe4, 0xad, 0x9d, 0xd6, 0xde, 0xd2, 0xb0, 0x53, 0xa1, 0x2f, 0x65, 0x82, 0xbb, 0x29, 0x74,
	0x28, 0x88, 0x8c, 0x22, 0xe3, 0xce, 0x32, 0x64, 0x9b, 0x70, 0x45, 0x4e, 0xea, 0x86, 0x05, 0x39,
	0xf1, 0x4a, 0x76, 0x1b, 0x96, 0xa9, 0x74, 0xaf, 0x44, 0x6b, 0xf9, 0xff, 0x88, 0x6d, 0x7b, 0x6c,
	0x3f, 0x40, 0x85, 0xe4, 0x59, 0x25, 0xb9, 0x54, 0x49, 0x9e, 0x15, 0x92, 0xdd, 0x3f, 0xd6, 0x61,
	0xa5, 0x51, 0xf5, 0x48, 0x4e, 0xd9, 0x43, 0x58, 0x6d, 0xd6, 0x2a, 0x4e, 0x4d, 0x91, 0xbd, 0xd7,
	0x28, 0xf7, 0x17, 0xc3, 0xd6, 0xe1, 0xb2, 0x75, 0xd2, 0x21, 0xe5, 0xef, 0x0c, 0xc3, 0x1f, 0xf6,
	0x3b, 0x5c, 0xa9, 0x27, 0x6d, 0x3f, 0x8e, 0xfa, 0x1f, 0xa4, 0xe9, 0xfd, 0x46, 0xb3, 0x86, 0x65,
	0x52, 0x76, 0x07, 0x3a, 0x99, 0xc1, 0x89, 0x3a, 0x15, 0x31, 0xa6, 0x53, 0xf7, 0x96, 0xff, 0x9f,
	0xaa, 0x5b, 0x0e, 0xe0, 0x21, 0x61, 0xec, 0x3e, 0xf4, 0x54, 0x2d, 0x96, 0xc8, 0x33, 0x7e, 0x79,
	0xa7, 0xb5, 0xb7, 0x38, 0xec, 0x28, 0x3b, 0x28, 0xd1, 0x93, 0x8c, 0xed, 0x42, 0xc7, 0xeb, 0x32,
	0x81, 0xa9, 0x1c, 0xc5, 0x18, 0xf1, 0x05, 0x52, 0xb5, 0x95, 0x1d, 0x64, 0x07, 0x01, 0xa2, 0x96,
	0x59, 0x61, 0x74, 0xee, 0xd0, 0x54, 0xba, 0x2b, 0xa4, 0xeb, 0x29, 0x3b, 0x24, 0xbc, 0xd4, 0xfa,
	0xc7, 0xe2, 0xcb, 0x9e, 0xa1, 0xb1, 0x4a, 0xa7, 0x7c, 0x91, 0x6a, 0x6b, 0x7b, 0xec, 0x4d, 0x80,
	0xbc, 0xe4, 0xad, 0xb6, 0xae, 0x92, 0x2c, 0x05, 0x89, 0xc7, 0x4a, 0xc9, 0x3d, 0xe8, 0xbe, 0xcb,
	0xd1, 0x9c, 0x85, 0x03, 0xcc, 0x64, 0xcc, 0x81, 0x44, 0x1d, 0x42, 0x07, 0x05, 0xe8, 0x3b, 0x11,
	0x64, 0x4e, 0x25, 0xa8, 0x73, 0xc7, 0xdb, 0xa1, 0x13, 0x04, 0xbe, 0x0e, 0x18, 0xfb, 0x01, 0xb6,
	0x83, 0x28, 0x91, 0xa7, 0x2a, 0xc9, 0x13, 0x61, 0xd0, 0x66, 0x3a, 0xb5, 0x48, 0x2e, 0xbe, 0x4c,
	0x16, 0x4e, 0x92, 0xa3, 0xa0, 0x18, 0x16, 0x02, 0x1f, 0x81, 0x7d, 0x0f, 0x5b, 0xb1, 0xb4, 0x4e,
	0x24, 0x98, 0x8c, 0xd0, 0x88, 0x73, 0x65, 0x75, 0xc8, 0xbd, 0xe9, 0x15, 0x47, 0x24, 0x78, 0xd5,
	0x28, 0xf0, 0x16, 0xb4, 0xa7, 0x46, 0xe7, 0x99, 0xf8, 0x4d, 0xab, 0xd4, 0xf2, 0x2e, 0xa9, 0x81,
	0xa0, 0x1f, 0x3d, 0xe2, 0x7b, 0x11, 0x04, 0x31, 0xca, 0x19, 0x5a, 0xde, 0x0b, 0xbd, 0x20, 0xec,
	0x90, 0x20, 0x76, 0x03, 0x40, 0x59, 0xca, 0xab, 0xd0, 0xf0, 0x15, 0x6a, 0xfb, 0x92, 0xb2, 0xaf,
	0x02, 0xc0, 0xfe, 0x6c, 0x41, 0xaf, 0x20, 0xab, 0x77, 0x61, 0xf5, 0x23, 0x5e, 0xcb, 0x6e, 0x91,
	0xbc, 0x7c, 0x2f, 0xfb, 0xb0, 0xe6, 0xb4, 0x93, 0xb1, 0x08, 0xc1, 0x05, 0x1d, 0xc5, 0x72, 0x46,
	0x07, 0x5b, 0x25, 0x6a, 0x9f, 0x98, 0x17, 0x44, 0xb0, 0x9b, 0x00, 0x46, 0x8f, 0x72, 0xeb, 0x52,
	0x5f, 0xf9, 0x5a, 0xe8, 0xd0, 0x1c, 0x61, 0x0f, 0xa0, 0x97, 0x19, 0x7d, 0x7a, 0x36, 0x2f, 0x80,
	0xaf, 0xd3, 0xdb, 0xda, 0x25, 0xb8, 0xba, 0xcb, 0xe5, 0x9d, 0xf1, 0x7d, 0xc8, 0x33, 0x7a, 0xb4,
	0x1b, 0xf3, 0x3b, 0xa3, 0xd0, 0x9c, 0x10, 0xc8, 0xbe, 0x85, 0x6b, 0xb1, 0xb4, 0x22, 0x8e, 0x85,
	0xc1, 0xa9, 0xb2, 0xce, 0x48, 0xa7, 0x74, 0x2a, 0xc6, 0x3a, 0x4f, 0x1d, 0xbf, 0x4a, 0x8e, 0xab,
	0xb1, 0xb4, 0x87, 0xf1, 0xb0, 0x46, 0x3f, 0xf7, 0x2c, 0x7b, 0x04, 0x1b, 0xde, 0x3a, 0x45, 0x57,
	0x76, 0xba, 0xb0, 0x6d, 0x92, 0x8d, 0xc5, 0xd2, 0xbe, 0x40, 0x57, 0x34, 0x22, 0x58, 0xf6, 0x60,
	0xc5, 0x5b, 0xf2, 0x2c, 0x92, 0x0e, 0x0b, 0x35, 0x27, 0x75, 0x37, 0x96, 0xf6, 0x84, 0xe0, 0xa0,
	0xfc, 0x8e, 0xee, 0x59, 0xa8, 0x2b, 0xd1, 0x33, 0x6c, 0x7a, 0xae, 0x35, 0x0a, 0xf3, 0x7c, 0xdd,
	0xfb, 0x14, 0x36, 0x0b, 0xaf, 0x8c, 0xa2, 0xa6, 0x71, 0x8b, 0x8c, 0x6b, 0x64, 0xdc, 0x8f, 0xa2,
	0xba, 0xab, 0x38, 0x4e, 0x9a, 0xc7, 0x71, 0xd3, 0xb3, 0x5d, 0x1d, 0xe7, 0x65, 0x1e, 0xc7, 0x75,
	0xcb, 0x37, 0xc0, 0xe9, 0x38, 0xe9, 0x05, 0xbd, 0xbb, 0x5e, 0x95, 0x78, 0xd2, 0xa0, 0x83, 0xf3,
	0x2e, 0x74, 0x95, 0x15, 0xde, 0x6c, 0xf0, 0x5d, 0x8e, 0xd6, 0xf1, 0x1b, 0x74, 0x93, 0x97, 0x95,
	0x3d, 0x94, 0x76, 0x18, 0xb0, 0x62, 0xd2, 0x04, 0x95, 0x8f, 0x80, 0x06, 0x23, 0x7e, 0xb3, 0x9c,
	0x34, 0x24, 0x2c, 0x61, 0xb6, 0x01, 0x0b, 0x33, 0x33, 0x11, 0x2a, 0xe2, 0xb7, 0xc2, 0x74, 0x9e,
	0x99, 0xc9, 0x20, 0x62, 0xd7, 0x01, 0x12, 0x87, 0xa2, 0xa0, 0x76, 0x88, 0x5a, 0x4c, 0x1c, 0xbe,
	0x21, 0x76, 0x0b, 0x16, 0x63, 0x3d, 0xa6, 0xba, 0xf8, 0xed, 0xc0, 0x95, 0xff, 0xd9, 0x0a, 0x5c,
	0x4a, 0x5c, 0xce, 0x77, 0x09, 0xf6, 0x3f, 0xd9, 0x36, 0x2c, 0xf9, 0x38, 0x61, 0x07, 0xdc, 0x09,
	0xf2, 0x99, 0x99, 0x1c, 0xd3, 0x1a, 0xf8, 0x0a, 0xae, 0x2a, 0x7f, 0x01, 0xd2, 0x89, 0x9a, 0xe6,
	0xe1, 0xa8, 0x33, 0x34, 0x6a, 0x72, 0xc6, 0xef, 0x52, 0xc1, 0x1b, 0xca, 0x3e, 0x7f, 0x9f, 0x64,
	0x8f, 0x60, 0xbd, 0xe9, 0xf1, 0x09, 0xd0, 0xf1, 0x7b, 0x64, 0x5a, 0x3b, 0xcf, 0x1d, 0xa3, 0xf3,
	0x99, 0xde, 0xb3, 0xa0, 0x31, 0xda, 0xf0, 0xfb, 0x21, 0xd3, 0x79, 0xf6, 0xc0, 0x93, 0x7e, 0xf0,
	0x35, 0x08, 0x91, 0x8c, 0xfd, 0x20, 0x2b, 0x13, 0x3e, 0x20, 0x2f, 0x6f, 0x48, 0x8e, 0xbc, 0xe2,
	0x4d, 0xc8, 0xba, 0x0f, 0x37, 0xfe, 0xcb, 0x1e, 0x92, 0xef, 0x51, 0x80, 0xad, 0x0b, 0x03, 0x84,
	0x0a, 0x9e, 0x50, 0x8b, 0x54, 0x12, 0x3a, 0x58, 0x7f, 0xa6, 0x9f, 0x84, 0xd3, 0x2a, 0x3b, 0x48,
	0xa8, 0x9b, 0xb5, 0xe7, 0x7a, 0x87, 0x36, 0x92, 0xcd, 0x47, 0x76, 0x6c, 0xd4, 0x08, 0x0d, 0x7f,
	0x58, 0x5e, 0x94, 0xe3, 0x0a, 0xf3, 0x53, 0x61, 0xae, 0x10, 0x89, 0x8e, 0x90, 0x7f, 0x1a, 0x5e,
	0xab, 0x39, 0x7c, 0xa4, 0x23, 0xf4, 0xe3, 0xc8, 0x97, 0x10, 0x61, 0xea, 0x94, 0x3b, 0x13, 0x99,
	0x41, 0x8b, 0xa9, 0xe3, 0x9f, 0x51, 0xcc, 0x55, 0x65, 0x07, 0x05, 0xf3, 0x73, 0x20, 0xd8, 0x5f,
	0x2d, 0x60, 0xb5, 0xc8, 0xe5, 0x44, 0xfd, 0xfc, 0x23, 0x4e, 0xd4, 0xd5, 0x79, 0xfe, 0xfd, 0xf9,
	0xca, 0xaf, 0x15, 0xa5, 0x22, 0xde, 0xa7, 0x11, 0xb8, 0x3c, 0x07, 0x07, 0x11, 0x4d, 0x4a, 0x69,
	0x30, 0x75, 0x42, 0x4d, 0xde, 0xca, 0x34, 0x8a, 0x91, 0x7f, 0x51, 0x4c, 0x4a, 0x82, 0x07, 0x05,
	0xca, 0x0e, 0x60, 0xc7, 0x8f, 0x42, 0x61, 0x55, 0x3a, 0x46, 0x41, 0xdb, 0xad, 0x5c, 0x6b, 0xc2,
	0xe2, 0x58, 0xa7, 0x91, 0xe5, 0x5f, 0x52, 0x37, 0xb7, 0xbd, 0xee, 0xd8, 0xcb, 0x0e, 0xa5, 0x75,
	0xc5, 0x6e, 0x3b, 0x0e, 0x12, 0xf6, 0x02, 0x6e, 0x9f, 0x0f, 0x63, 0x30, 0xd3, 0xc6, 0xd5, 0xe3,
	0x3c, 0xa2, 0x38, 0xd7, 0x1b, 0x71, 0x86, 0xa4, 0x9a, 0x07, 0xfa, 0x1a, 0x78, 0xf1, 0x71, 0x11,
	0x06, 0x77, 0xdd, 0xff, 0x98, 0xfc, 0x1b, 0x81, 0x0f, 0x23, 0x7c, 0x6e, 0xbc, 0x0f, 0x3d, 0xff,
	0xae, 0xbb, 0x3c, 0x8b, 0xcb, 0xd9, 0xf5, 0x24, 0xcc, 0xfc, 0xc4, 0xe1, 0x6b, 0x8f, 0xd2, 0xf0,
	0xd9, 0xfd, 0xa7, 0x05, 0x1b, 0x17, 0x7e, 0xbe, 0xb2, 0xbf, 0x5b, 0xb0, 0x4e, 0xcc, 0xfc, 0x4b,
	0x09, 0x53, 0x67, 0xce, 0x28, 0x6f, 0xfb, 0xf1, 0xf4, 0x43, 0x3e, 0xf0, 0xda, 0x57, 0xe9, 0x90,
	0x79, 0xa4, 0xda, 0x65, 0x07, 0xbe, 0x04, 0xbf, 0x3b, 0xfc, 0x67, 0x43, 0xd8, 0xa0, 0x8d, 0xe3,
	0x75, 0x3d, 0x4e, 0xfb, 0x93, 0xce, 0x37, 0x5a, 0xa0, 0x6f, 0xff, 0x27, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x29, 0xca, 0x73, 0x03, 0x1a, 0x0c, 0x00, 0x00,
}
