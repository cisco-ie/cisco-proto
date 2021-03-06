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

package cisco_ios_xr_ipv4_igmp_oper_mld_active_default_context_interface_table_interface

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
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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
	proto.RegisterType((*IgmpEdmIdbBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.interface_table.interface.igmp_edm_idb_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.interface_table.interface.igmp_addrtype")
	proto.RegisterType((*IgmpEdmIdbBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.interface_table.interface.igmp_edm_idb_bag")
}

func init() { proto.RegisterFile("igmp_edm_idb_bag.proto", fileDescriptor_bfc702e9ae24d547) }

var fileDescriptor_bfc702e9ae24d547 = []byte{
	// 1188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5b, 0x73, 0xdb, 0x44,
	0x14, 0x1e, 0xd3, 0xa1, 0x97, 0x4d, 0x6c, 0x37, 0x9b, 0xa6, 0x5d, 0x7a, 0x81, 0x34, 0xa5, 0x25,
	0xc3, 0x83, 0x67, 0x7a, 0x4b, 0x6f, 0xb4, 0x60, 0x42, 0x08, 0x6e, 0x9d, 0x90, 0xd8, 0x49, 0x68,
	0xb9, 0x2d, 0x6b, 0xe9, 0xd8, 0x59, 0x46, 0xd2, 0xaa, 0xbb, 0x2b, 0x8f, 0xfd, 0x0f, 0x78, 0xe2,
	0xef, 0xf2, 0xca, 0xec, 0x59, 0xc9, 0xb2, 0xdc, 0xf2, 0xc6, 0xf4, 0x2d, 0xfe, 0x2e, 0xe7, 0x9c,
	0x3d, 0x5a, 0x7d, 0x0a, 0xb9, 0x2c, 0x47, 0x71, 0xca, 0x21, 0x8c, 0xb9, 0x0c, 0x07, 0x7c, 0x20,
	0x46, 0xad, 0x54, 0x2b, 0xab, 0xe8, 0x41, 0x20, 0x4d, 0xa0, 0xb8, 0x54, 0x86, 0x4f, 0x34, 0x97,
	0xe9, 0xf8, 0x01, 0x47, 0xa5, 0x4a, 0x41, 0xb7, 0xe2, 0x28, 0x6c, 0x89, 0xc0, 0xca, 0x31, 0xb4,
	0x42, 0x18, 0x8a, 0x2c, 0xb2, 0x3c, 0x50, 0x89, 0x85, 0x89, 0x6d, 0xc9, 0xc4, 0x82, 0x1e, 0x8a,
	0x00, 0xb8, 0x15, 0x83, 0x08, 0xca, 0xdf, 0x1b, 0x2f, 0xc8, 0xda, 0x62, 0x2f, 0xfe, 0x6a, 0xe7,
	0x4d, 0x9f, 0xde, 0x26, 0x8d, 0xd2, 0x95, 0x88, 0x18, 0x58, 0x6d, 0xbd, 0xb6, 0x79, 0xa1, 0x57,
	0x9f, 0xa1, 0xfb, 0x22, 0x86, 0x8d, 0x84, 0xd4, 0xd1, 0x2f, 0xc2, 0x50, 0xdb, 0x69, 0x0a, 0xf4,
	0x0a, 0x39, 0x27, 0x86, 0xf3, 0x86, 0xb3, 0x62, 0xe8, 0x94, 0xf4, 0x26, 0x59, 0xc6, 0x81, 0x9d,
	0x12, 0x8c, 0x61, 0x1f, 0x21, 0xbb, 0xe4, 0xb0, 0xb6, 0x87, 0x72, 0xc9, 0xd6, 0x4c, 0x72, 0x66,
	0x26, 0xd9, 0xca, 0x25, 0x1b, 0xff, 0xac, 0x92, 0x8b, 0x8b, 0x03, 0xd3, 0x2f, 0xc9, 0x4a, 0x75,
	0x56, 0x3e, 0xd1, 0xec, 0x1e, 0x9a, 0x9b, 0x95, 0x71, 0x5f, 0x6b, 0x7a, 0x89, 0x7c, 0x6c, 0xac,
	0xb0, 0xc0, 0xee, 0xaf, 0xd7, 0x36, 0xeb, 0x3d, 0xff, 0x83, 0x4e, 0xc9, 0xb9, 0xa2, 0xe9, 0x83,
	0xf5, 0xda, 0xe6, 0xd2, 0x3d, 0xde, 0xfa, 0xbf, 0x57, 0xdd, 0xaa, 0xec, 0xa9, 0x57, 0xf4, 0xa3,
	0xb7, 0x48, 0x3d, 0xd5, 0x30, 0x94, 0x13, 0x1e, 0x41, 0x32, 0xb2, 0xa7, 0xec, 0x21, 0x0e, 0xb6,
	0xec, 0xc1, 0x2e, 0x62, 0xf4, 0x0e, 0x69, 0x4a, 0xc3, 0xcb, 0xda, 0x59, 0xca, 0xb6, 0xd6, 0x6b,
	0x9b, 0xe7, 0x7b, 0x75, 0x69, 0x3a, 0x05, 0x7a, 0x9c, 0xd2, 0x0d, 0x52, 0x77, 0xba, 0x94, 0x43,
	0xe2, 0xba, 0x87, 0xec, 0x11, 0xaa, 0x96, 0xa4, 0xe9, 0xa4, 0x3b, 0x1e, 0xc2, 0x6d, 0x19, 0xae,
	0x55, 0x66, 0x41, 0xcf, 0x74, 0x8f, 0x51, 0xd7, 0x94, 0xa6, 0x87, 0x78, 0xa1, 0x75, 0x4f, 0xc4,
	0x8d, 0x3d, 0x06, 0x6d, 0xa4, 0x4a, 0xd8, 0x13, 0x9c, 0x6d, 0xc9, 0x61, 0x27, 0x1e, 0x72, 0x92,
	0x53, 0x65, 0xec, 0x4c, 0xf2, 0xd4, 0x4b, 0x1c, 0x56, 0x48, 0x6e, 0x93, 0xc6, 0xdb, 0x0c, 0xf4,
	0xd4, 0x1f, 0x60, 0x2c, 0x22, 0xf6, 0x0c, 0x45, 0x75, 0x44, 0x3b, 0x39, 0xe8, 0x36, 0xe1, 0x65,
	0x56, 0xc6, 0xa0, 0x32, 0xcb, 0xbe, 0xf2, 0x9b, 0x40, 0xf0, 0xc8, 0x63, 0xf4, 0x39, 0xb9, 0xe6,
	0x45, 0xb1, 0x98, 0xc8, 0x38, 0x8b, 0xb9, 0x06, 0x93, 0xaa, 0xc4, 0x00, 0xba, 0xd8, 0x73, 0xb4,
	0x30, 0x94, 0xec, 0x79, 0x45, 0x2f, 0x17, 0xb8, 0x0a, 0xf4, 0x19, 0xb9, 0x1a, 0x09, 0x63, 0x79,
	0x0c, 0xf1, 0x00, 0x34, 0x5f, 0x18, 0xeb, 0x05, 0xba, 0xaf, 0x38, 0xc5, 0x1e, 0x0a, 0x0e, 0x2b,
	0x03, 0x7e, 0x46, 0x96, 0x46, 0x5a, 0x65, 0x29, 0xff, 0x53, 0xc9, 0xc4, 0xb0, 0xaf, 0x51, 0x4d,
	0x10, 0x7a, 0xe9, 0x10, 0xb7, 0x0b, 0x2f, 0x88, 0x40, 0x8c, 0xc1, 0xb0, 0x6f, 0xfc, 0x2e, 0x10,
	0xeb, 0x22, 0x44, 0x6f, 0x10, 0x22, 0x0d, 0xf6, 0x95, 0xa0, 0x59, 0x1b, 0xd7, 0x7e, 0x41, 0x9a,
	0x43, 0x0f, 0xd0, 0xbf, 0x6a, 0xa4, 0x99, 0x93, 0xb3, 0xd7, 0xe0, 0xdb, 0x0f, 0x73, 0x23, 0x1b,
	0x79, 0xdf, 0xe2, 0x6d, 0x6c, 0x91, 0x55, 0xab, 0xac, 0x88, 0xb8, 0xaf, 0xcb, 0xf1, 0x14, 0x86,
	0x6d, 0xe3, 0x99, 0x56, 0x90, 0x6a, 0x23, 0xb3, 0x8b, 0x04, 0xfd, 0x94, 0x10, 0xad, 0x06, 0x99,
	0xb1, 0x89, 0x1b, 0xfa, 0x3b, 0xbf, 0x9c, 0x12, 0xa1, 0x5f, 0x90, 0x66, 0xaa, 0xd5, 0x64, 0x5a,
	0x5e, 0x63, 0xb6, 0x83, 0xef, 0x68, 0x03, 0xe1, 0xd9, 0x35, 0x2e, 0xae, 0x8b, 0x5b, 0x41, 0x96,
	0xe2, 0x53, 0xfd, 0xbe, 0xbc, 0x2e, 0x12, 0xf4, 0x31, 0x82, 0xf4, 0x09, 0xf9, 0x24, 0x12, 0x86,
	0x47, 0x11, 0xd7, 0x30, 0x92, 0xc6, 0x6a, 0x61, 0xa5, 0x4a, 0x78, 0xa0, 0xb2, 0xc4, 0xb2, 0x5d,
	0x74, 0x5c, 0x8e, 0x84, 0xe9, 0x46, 0xbd, 0x39, 0x7a, 0xdb, 0xb1, 0xf4, 0x2e, 0x59, 0x73, 0xd6,
	0x11, 0xd8, 0x62, 0xc9, 0xb9, 0xed, 0x07, 0xb4, 0xd1, 0x48, 0x98, 0x5d, 0xb0, 0xf9, 0x22, 0xbc,
	0x65, 0x93, 0x5c, 0x74, 0x96, 0x2c, 0x0d, 0x85, 0x85, 0x5c, 0xdd, 0x41, 0x75, 0x23, 0x12, 0xe6,
	0x18, 0x61, 0xaf, 0x7c, 0x8a, 0x57, 0xcc, 0xcf, 0x15, 0xab, 0x31, 0x54, 0x3d, 0x2f, 0x2b, 0x83,
	0x39, 0x7e, 0xde, 0xfb, 0x80, 0x5c, 0xc9, 0xbd, 0x22, 0x0c, 0xab, 0xc6, 0x57, 0x68, 0x5c, 0x45,
	0x63, 0x3b, 0x0c, 0xe7, 0x5d, 0xf9, 0x71, 0x92, 0x2c, 0x8a, 0xaa, 0x9e, 0xee, 0xec, 0x38, 0xfb,
	0x59, 0x14, 0xcd, 0x5b, 0x1e, 0x13, 0x86, 0xc7, 0x49, 0xde, 0xb3, 0xbb, 0xbd, 0xd9, 0x88, 0xc7,
	0x15, 0xda, 0x3b, 0x3f, 0x27, 0x0d, 0x69, 0xb8, 0x33, 0x6b, 0x78, 0x9b, 0x81, 0xb1, 0x6c, 0x1f,
	0x2f, 0xf1, 0xb2, 0x34, 0x5d, 0x61, 0x7a, 0x1e, 0xcb, 0x43, 0xc6, 0xab, 0x5c, 0x05, 0xd0, 0x10,
	0xb2, 0x1f, 0x8b, 0x90, 0x41, 0x61, 0x01, 0xd3, 0x35, 0x72, 0x76, 0xac, 0x87, 0x5c, 0x86, 0xec,
	0xc0, 0x67, 0xf2, 0x58, 0x0f, 0x3b, 0x21, 0xbd, 0x4e, 0x48, 0x6c, 0x81, 0xe7, 0xd4, 0x21, 0x52,
	0xe7, 0x63, 0x0b, 0x27, 0xc8, 0x5e, 0x25, 0xe7, 0x23, 0x15, 0xe0, 0x5c, 0xac, 0xe7, 0xb9, 0xe2,
	0x37, 0xbd, 0x48, 0xce, 0xc4, 0x36, 0x63, 0x7d, 0x84, 0xdd, 0x9f, 0xf4, 0x1a, 0xb9, 0xe0, 0xea,
	0xf8, 0xe4, 0x3f, 0xf2, 0xf2, 0xb1, 0x1e, 0xf6, 0x31, 0xfc, 0x1f, 0x92, 0xcb, 0xd2, 0x5d, 0x80,
	0x64, 0x28, 0x47, 0x99, 0x3f, 0xea, 0x18, 0xb4, 0x1c, 0x4e, 0xd9, 0x31, 0x0e, 0xbc, 0x26, 0xcd,
	0xf6, 0xbb, 0x24, 0xbd, 0x4b, 0x2e, 0x55, 0x3d, 0xae, 0x01, 0x58, 0x76, 0x82, 0xa6, 0xd5, 0x45,
	0xae, 0x0f, 0xd6, 0x75, 0x7a, 0xc7, 0x02, 0x5a, 0x2b, 0xcd, 0x7e, 0xf2, 0x9d, 0x16, 0xd9, 0x1d,
	0x47, 0xba, 0xcc, 0xab, 0x10, 0x3c, 0x0e, 0x5c, 0x86, 0x15, 0x0d, 0x5f, 0xa3, 0x97, 0x55, 0x24,
	0x7b, 0x4e, 0x71, 0xe2, 0xbb, 0xb6, 0xc9, 0x8d, 0xff, 0xb2, 0xfb, 0xe6, 0x6f, 0xb0, 0xc0, 0xd5,
	0xf7, 0x16, 0xf0, 0x13, 0xdc, 0xc7, 0x15, 0xc9, 0xd8, 0x6f, 0x70, 0xfe, 0x99, 0xfe, 0xec, 0x4f,
	0x2b, 0x4d, 0x27, 0xc6, 0x6d, 0xce, 0x3d, 0xd7, 0x5b, 0xf8, 0x31, 0x32, 0xd9, 0xc0, 0x04, 0x5a,
	0x0e, 0x40, 0xb3, 0x5f, 0x8a, 0x8b, 0xd2, 0x9f, 0x61, 0x2e, 0x15, 0x4a, 0x05, 0x8f, 0x55, 0x08,
	0xec, 0x57, 0xff, 0x5a, 0x95, 0xf0, 0x9e, 0x0a, 0xc1, 0xc5, 0x91, 0x1b, 0x21, 0x84, 0xc4, 0x4a,
	0x3b, 0xe5, 0xa9, 0x06, 0x03, 0x89, 0x65, 0xbf, 0x61, 0xcd, 0x15, 0x69, 0x3a, 0x39, 0x73, 0xe0,
	0x09, 0xfa, 0x77, 0x8d, 0xd0, 0xb9, 0xca, 0x45, 0x98, 0xfe, 0xfe, 0x61, 0xc2, 0x74, 0xa5, 0x6c,
	0xdd, 0x2e, 0x3f, 0xf4, 0x73, 0xf3, 0xc8, 0x90, 0x71, 0x4c, 0xbf, 0xe5, 0x12, 0xec, 0x84, 0x18,
	0x92, 0x42, 0x43, 0x62, 0xb9, 0x1c, 0x9e, 0x8a, 0x24, 0x8c, 0x80, 0xfd, 0x91, 0x87, 0x24, 0xc2,
	0x9d, 0x1c, 0xa5, 0x3b, 0x64, 0xdd, 0xa5, 0x20, 0x37, 0x32, 0x09, 0x80, 0xe3, 0x37, 0xad, 0xf8,
	0x98, 0x71, 0x03, 0x81, 0x4a, 0x42, 0xc3, 0x04, 0x2e, 0xf2, 0x9a, 0xd3, 0xf5, 0x9d, 0xac, 0x2b,
	0x8c, 0xcd, 0xbf, 0x68, 0x7d, 0x2f, 0xa1, 0xbb, 0xe4, 0xe6, 0x62, 0x19, 0x0d, 0xa9, 0xd2, 0x76,
	0xbe, 0xce, 0x00, 0xeb, 0x5c, 0xaf, 0xd4, 0xe9, 0xa1, 0xaa, 0x2c, 0xf4, 0x88, 0xb0, 0xfc, 0x5f,
	0x0a, 0x9f, 0xd9, 0xf3, 0xfe, 0x00, 0xfd, 0x6b, 0x9e, 0xf7, 0xe9, 0x5d, 0x1a, 0xef, 0x90, 0xa6,
	0x7b, 0xcd, 0x6d, 0x96, 0x46, 0x45, 0x6c, 0x85, 0x3e, 0xee, 0x63, 0x0b, 0x47, 0x0e, 0xc5, 0xdc,
	0x19, 0x9c, 0xc5, 0x7f, 0x81, 0xef, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x31, 0x49, 0x3f,
	0x1c, 0x0b, 0x00, 0x00,
}
