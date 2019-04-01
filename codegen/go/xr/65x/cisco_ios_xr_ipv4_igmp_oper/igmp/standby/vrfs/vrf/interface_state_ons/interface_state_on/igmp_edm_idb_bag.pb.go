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

package cisco_ios_xr_ipv4_igmp_oper_igmp_standby_vrfs_vrf_interface_state_ons_interface_state_on

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
	proto.RegisterType((*IgmpEdmIdbBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.vrfs.vrf.interface_state_ons.interface_state_on.igmp_edm_idb_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.vrfs.vrf.interface_state_ons.interface_state_on.igmp_addrtype")
	proto.RegisterType((*IgmpEdmIdbBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.vrfs.vrf.interface_state_ons.interface_state_on.igmp_edm_idb_bag")
}

func init() { proto.RegisterFile("igmp_edm_idb_bag.proto", fileDescriptor_bfc702e9ae24d547) }

var fileDescriptor_bfc702e9ae24d547 = []byte{
	// 1206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0x5b, 0x73, 0xdb, 0x36,
	0x13, 0x86, 0x47, 0x5f, 0xe6, 0x4b, 0x1c, 0xd8, 0x92, 0x62, 0x24, 0x4e, 0x90, 0x53, 0xeb, 0x38,
	0x4d, 0xea, 0xe9, 0x85, 0x66, 0x72, 0x72, 0x4e, 0x4d, 0x5b, 0xd5, 0x75, 0x5d, 0x25, 0xb2, 0x1b,
	0x4b, 0xb6, 0x6b, 0xf7, 0x84, 0x42, 0xe2, 0x4a, 0x46, 0x87, 0x24, 0x18, 0x00, 0xe4, 0x58, 0x97,
	0xfd, 0x07, 0xb9, 0xea, 0xef, 0xed, 0x60, 0x41, 0x8a, 0xa2, 0xe2, 0x5e, 0x36, 0x37, 0x1a, 0xe9,
	0xdd, 0xe7, 0x5d, 0x2c, 0x96, 0xe0, 0x42, 0xe4, 0xaa, 0x1c, 0x47, 0x09, 0x87, 0x20, 0xe2, 0x32,
	0x18, 0xf0, 0x81, 0x18, 0xb7, 0x12, 0xad, 0xac, 0xa2, 0x47, 0x43, 0x69, 0x86, 0x8a, 0x4b, 0x65,
	0xf8, 0xa9, 0xe6, 0x32, 0xc9, 0x1e, 0x73, 0x24, 0x55, 0x02, 0xba, 0xe5, 0xbe, 0xb5, 0x8c, 0x15,
	0x71, 0x30, 0x98, 0xb4, 0x32, 0x3d, 0x32, 0xee, 0xa3, 0x25, 0x63, 0x0b, 0x7a, 0x24, 0x86, 0xc0,
	0x8d, 0x15, 0x16, 0xb8, 0x8a, 0xcd, 0x19, 0xda, 0xda, 0x31, 0x59, 0x99, 0x5f, 0x93, 0xbf, 0xd9,
	0x3a, 0xee, 0xd3, 0xeb, 0x64, 0x21, 0xd3, 0x23, 0x1e, 0x8b, 0x08, 0x58, 0x6d, 0xb5, 0xb6, 0x7e,
	0xb1, 0x77, 0x21, 0xd3, 0xa3, 0x5d, 0x11, 0x01, 0xbd, 0x47, 0x1a, 0x65, 0x26, 0x04, 0xfe, 0x87,
	0x40, 0x7d, 0xaa, 0x3a, 0x6c, 0x2d, 0x26, 0x75, 0x4c, 0x2d, 0x82, 0x40, 0xdb, 0x49, 0x02, 0xf4,
	0x1a, 0xb9, 0x20, 0x2a, 0x19, 0xcf, 0x0b, 0x9f, 0xf0, 0x0e, 0x59, 0xc2, 0x3d, 0x39, 0x12, 0x8c,
	0xc9, 0xd3, 0x2d, 0x3a, 0xad, 0xed, 0xa5, 0x1c, 0xd9, 0x98, 0x22, 0xe7, 0xa6, 0xc8, 0x46, 0x8e,
	0xac, 0xbd, 0xbf, 0x42, 0x2e, 0xcd, 0xef, 0x85, 0x7e, 0x41, 0x96, 0xab, 0xb5, 0xf2, 0x53, 0xcd,
	0x1e, 0xa2, 0xb9, 0x59, 0x29, 0xf7, 0x48, 0xd3, 0x2b, 0xe4, 0xff, 0xd8, 0x17, 0xf6, 0x68, 0xb5,
	0xb6, 0x5e, 0xef, 0xf9, 0x1f, 0xf4, 0xaf, 0x1a, 0xb9, 0x50, 0xac, 0xfa, 0x78, 0xb5, 0xb6, 0xbe,
	0xf8, 0x70, 0xdc, 0xfa, 0xaf, 0x1e, 0x47, 0xab, 0xd2, 0xb0, 0x5e, 0xb1, 0x2e, 0xbd, 0x4b, 0xea,
	0x89, 0x86, 0x91, 0x3c, 0xe5, 0x21, 0xc4, 0x63, 0x7b, 0xc2, 0x9e, 0x60, 0x85, 0x4b, 0x5e, 0xec,
	0xa2, 0x46, 0xef, 0x93, 0xa6, 0x34, 0xbc, 0x4c, 0x9a, 0x26, 0x6c, 0x63, 0xb5, 0xb6, 0xbe, 0xd0,
	0xab, 0x4b, 0xd3, 0x29, 0xd4, 0x83, 0x84, 0xae, 0x91, 0xba, 0xe3, 0x12, 0x0e, 0xb1, 0x18, 0x84,
	0x10, 0xb0, 0xa7, 0x48, 0x2d, 0x4a, 0xd3, 0x49, 0xb6, 0xbc, 0x84, 0x6d, 0x33, 0x5c, 0xab, 0xd4,
	0x82, 0x9e, 0x72, 0xcf, 0x90, 0x6b, 0x4a, 0xd3, 0x43, 0xbd, 0x60, 0xdd, 0xa3, 0x71, 0x65, 0x67,
	0xa0, 0x8d, 0x54, 0x31, 0x7b, 0x8e, 0xb5, 0x2d, 0x3a, 0xed, 0xd0, 0x4b, 0x0e, 0x39, 0x51, 0xc6,
	0x4e, 0x91, 0x17, 0x1e, 0x71, 0x5a, 0x81, 0xdc, 0x23, 0x8d, 0x77, 0x29, 0xe8, 0x89, 0xdf, 0x40,
	0x26, 0x42, 0xf6, 0x12, 0xa1, 0x3a, 0xaa, 0x9d, 0x5c, 0x74, 0x9d, 0xf0, 0x98, 0x95, 0x11, 0xa8,
	0xd4, 0xb2, 0x2f, 0x7d, 0x27, 0x50, 0xdc, 0xf7, 0x1a, 0x7d, 0x45, 0x6e, 0x7a, 0x28, 0x12, 0xa7,
	0x32, 0x4a, 0x23, 0xae, 0xc1, 0x24, 0x2a, 0x36, 0x80, 0x2e, 0xf6, 0x0a, 0x2d, 0x0c, 0x91, 0x1d,
	0x4f, 0xf4, 0x72, 0xc0, 0x65, 0xa0, 0x2f, 0xc9, 0x8d, 0x50, 0x18, 0xcb, 0x23, 0x88, 0x06, 0xa0,
	0xf9, 0x5c, 0x59, 0x5f, 0xa1, 0xfb, 0x9a, 0x23, 0x76, 0x10, 0xd8, 0xab, 0x14, 0xf8, 0x29, 0x59,
	0x1c, 0x6b, 0x95, 0x26, 0xfc, 0x4f, 0x25, 0x63, 0xc3, 0xbe, 0x46, 0x9a, 0xa0, 0xf4, 0xda, 0x29,
	0xae, 0x17, 0x1e, 0x08, 0x41, 0x64, 0x60, 0xd8, 0x37, 0xbe, 0x17, 0xa8, 0x75, 0x51, 0xa2, 0xb7,
	0x09, 0x91, 0x06, 0xd7, 0x95, 0xa0, 0x59, 0x1b, 0xdb, 0x7e, 0x51, 0x9a, 0x3d, 0x2f, 0xd0, 0xf7,
	0x35, 0xd2, 0xcc, 0x83, 0xd3, 0xf7, 0xe1, 0xdb, 0x8f, 0x7b, 0x32, 0x1b, 0xf9, 0xfa, 0xc5, 0xeb,
	0xd9, 0x22, 0x97, 0xad, 0xb2, 0x22, 0xe4, 0x62, 0x68, 0x65, 0x06, 0x1c, 0x77, 0x63, 0xd8, 0x26,
	0xee, 0x6d, 0x19, 0x43, 0x6d, 0x8c, 0x6c, 0x63, 0x80, 0x7e, 0x42, 0x88, 0x56, 0x83, 0xd4, 0xd8,
	0xd8, 0x15, 0xff, 0x9d, 0x6f, 0x52, 0xa9, 0xd0, 0xcf, 0x49, 0x33, 0xd1, 0xea, 0x74, 0x52, 0x1e,
	0x67, 0xb6, 0x85, 0x2f, 0x6d, 0x03, 0xe5, 0xe9, 0x71, 0x2e, 0x8e, 0x8d, 0x6b, 0x45, 0x9a, 0xe0,
	0xd3, 0xfd, 0xbe, 0x3c, 0x36, 0x12, 0xf4, 0x01, 0x8a, 0xf4, 0x39, 0xb9, 0x1e, 0x0a, 0xc3, 0xc3,
	0x90, 0x6b, 0x18, 0x4b, 0x63, 0xb5, 0xb0, 0x52, 0xc5, 0x7c, 0xa8, 0xd2, 0xd8, 0xb2, 0x6d, 0x74,
	0x5c, 0x0d, 0x85, 0xe9, 0x86, 0xbd, 0x99, 0xf0, 0xa6, 0x8b, 0xd2, 0x07, 0x64, 0xc5, 0x59, 0xc7,
	0x60, 0x8b, 0x66, 0xe7, 0xb6, 0x1f, 0xd0, 0x46, 0x43, 0x61, 0xb6, 0xc1, 0xe6, 0x8d, 0xf0, 0x96,
	0x75, 0x72, 0xc9, 0x59, 0xd2, 0x24, 0x70, 0xcd, 0xf3, 0x74, 0x07, 0xe9, 0x46, 0x28, 0xcc, 0x01,
	0xca, 0x9e, 0x7c, 0x81, 0x47, 0xcd, 0xd7, 0x15, 0xa9, 0x0c, 0xaa, 0x9e, 0xd7, 0x95, 0xc2, 0x5c,
	0x7c, 0xd6, 0xfb, 0x98, 0x5c, 0xcb, 0xbd, 0x22, 0x08, 0xaa, 0xc6, 0x37, 0x68, 0xbc, 0x8c, 0xc6,
	0x76, 0x10, 0xcc, 0xba, 0xf2, 0xed, 0xc4, 0x69, 0x18, 0x56, 0x3d, 0xdd, 0xe9, 0x76, 0x76, 0xd3,
	0x30, 0x9c, 0xb5, 0x3c, 0x23, 0x0c, 0xb7, 0x13, 0x9f, 0xd1, 0xbb, 0x9d, 0x69, 0x89, 0x07, 0x95,
	0xb0, 0x77, 0x7e, 0x46, 0x1a, 0xd2, 0x70, 0x67, 0xd6, 0xf0, 0x2e, 0x05, 0x63, 0xd9, 0x2e, 0x1e,
	0xe6, 0x25, 0x69, 0xba, 0xc2, 0xf4, 0xbc, 0x96, 0x0f, 0x1b, 0x4f, 0xb9, 0x0c, 0xa0, 0x21, 0x60,
	0x3f, 0x16, 0xc3, 0x06, 0xc1, 0x42, 0xa6, 0x2b, 0xe4, 0xbc, 0xbb, 0x96, 0x64, 0xc0, 0xde, 0xfa,
	0x21, 0x9d, 0xe9, 0x51, 0x27, 0xa0, 0xb7, 0x08, 0x89, 0x2c, 0xf0, 0x3c, 0xb4, 0x87, 0xa1, 0x85,
	0xc8, 0xc2, 0x21, 0x46, 0x6f, 0x90, 0x85, 0x50, 0x0d, 0xb1, 0x2e, 0xd6, 0xf3, 0xb1, 0xe2, 0x37,
	0xbd, 0x44, 0xce, 0x45, 0x36, 0x65, 0x7d, 0x94, 0xdd, 0x57, 0x7a, 0x93, 0x5c, 0x74, 0x79, 0xfc,
	0x55, 0xb0, 0xef, 0xf1, 0x4c, 0x8f, 0xfa, 0x78, 0x1b, 0x3c, 0x21, 0x57, 0xa5, 0x3b, 0x00, 0xf1,
	0x48, 0x8e, 0x53, 0xbf, 0xd5, 0x0c, 0xb4, 0x1c, 0x4d, 0xd8, 0x01, 0x16, 0xbc, 0x22, 0xcd, 0xe6,
	0x87, 0x41, 0xfa, 0x80, 0x5c, 0xa9, 0x7a, 0xdc, 0x02, 0x60, 0xd9, 0x21, 0x9a, 0x2e, 0xcf, 0xc7,
	0xfa, 0x60, 0xdd, 0x4a, 0x1f, 0x58, 0x40, 0x6b, 0xa5, 0xd9, 0x4f, 0x7e, 0xa5, 0xf9, 0xe8, 0x96,
	0x0b, 0xba, 0xd9, 0x57, 0x09, 0xf0, 0x68, 0xe8, 0x66, 0x59, 0xb1, 0xe0, 0x11, 0x7a, 0x59, 0x05,
	0xd9, 0x71, 0xc4, 0xa1, 0x5f, 0xb5, 0x4d, 0x6e, 0xff, 0x9b, 0xdd, 0x2f, 0x7e, 0x8c, 0x09, 0x6e,
	0x9c, 0x99, 0xc0, 0x57, 0xf0, 0x08, 0x5b, 0x24, 0xa3, 0x7c, 0x76, 0xcc, 0x3c, 0xd3, 0x9f, 0xfd,
	0x6e, 0xa5, 0xe9, 0x44, 0xd8, 0xcd, 0x99, 0xe7, 0x7a, 0x17, 0x2f, 0x25, 0x93, 0x0e, 0xcc, 0x50,
	0xcb, 0x01, 0x68, 0xf6, 0x4b, 0x71, 0x50, 0xfa, 0x53, 0xcd, 0x4d, 0x85, 0x92, 0xe0, 0x91, 0x0a,
	0x80, 0xfd, 0xea, 0x5f, 0xab, 0x52, 0xde, 0x51, 0x01, 0xb8, 0x71, 0xe4, 0x4a, 0x08, 0x20, 0xb6,
	0xd2, 0x4e, 0x78, 0xa2, 0xc1, 0x40, 0x6c, 0xd9, 0x6f, 0x98, 0x73, 0x59, 0x9a, 0x4e, 0x1e, 0x79,
	0xeb, 0x03, 0xf4, 0xef, 0x1a, 0xa1, 0x33, 0x99, 0x8b, 0xa1, 0xfa, 0xfb, 0xc7, 0x1d, 0xaa, 0xcb,
	0x65, 0x09, 0xed, 0xf2, 0xe2, 0x9f, 0xa9, 0x4b, 0x06, 0x8c, 0xe3, 0x14, 0x5c, 0x2a, 0xc5, 0x4e,
	0x80, 0xc3, 0x52, 0x68, 0x88, 0x2d, 0x97, 0xa3, 0x13, 0x11, 0x07, 0x21, 0xb0, 0x3f, 0xf2, 0x61,
	0x89, 0x72, 0x27, 0x57, 0xe9, 0x16, 0x59, 0x75, 0xd3, 0x90, 0x1b, 0x19, 0x0f, 0x81, 0xe3, 0x1d,
	0x57, 0x5c, 0x6e, 0xdc, 0xc0, 0x50, 0xc5, 0x81, 0x61, 0x02, 0x1b, 0x7a, 0xd3, 0x71, 0x7d, 0x87,
	0x75, 0x85, 0xb1, 0xf9, 0x0d, 0xd7, 0xf7, 0x08, 0xdd, 0x26, 0x77, 0xe6, 0xd3, 0x68, 0x48, 0x94,
	0xb6, 0xb3, 0x79, 0x06, 0x98, 0xe7, 0x56, 0x25, 0x4f, 0x0f, 0xa9, 0x32, 0xd1, 0x53, 0xc2, 0xf2,
	0xbf, 0x18, 0x7e, 0x76, 0xcf, 0xfa, 0x87, 0xe8, 0x5f, 0xf1, 0x71, 0x3f, 0xc5, 0x4b, 0xe3, 0x7d,
	0xd2, 0x74, 0xaf, 0xbb, 0x4d, 0x93, 0xb0, 0x18, 0x5f, 0x81, 0x1f, 0xfb, 0x91, 0x85, 0x7d, 0xa7,
	0xe2, 0xfc, 0x19, 0x9c, 0xc7, 0xbf, 0xcf, 0x8f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x91, 0xe1,
	0x73, 0x1c, 0x58, 0x0b, 0x00, 0x00,
}