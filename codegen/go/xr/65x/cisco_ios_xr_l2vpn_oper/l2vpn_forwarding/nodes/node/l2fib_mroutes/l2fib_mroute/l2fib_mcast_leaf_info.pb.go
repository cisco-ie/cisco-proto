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
// source: l2fib_mcast_leaf_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_mroutes_l2fib_mroute

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

type L2FibMcastLeafInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Group                string   `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	BdGroupName          string   `protobuf:"bytes,4,opt,name=bd_group_name,json=bdGroupName,proto3" json:"bd_group_name,omitempty"`
	BdName               string   `protobuf:"bytes,5,opt,name=bd_name,json=bdName,proto3" json:"bd_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibMcastLeafInfo_KEYS) Reset()         { *m = L2FibMcastLeafInfo_KEYS{} }
func (m *L2FibMcastLeafInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibMcastLeafInfo_KEYS) ProtoMessage()    {}
func (*L2FibMcastLeafInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_96583c98adae0568, []int{0}
}

func (m *L2FibMcastLeafInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMcastLeafInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibMcastLeafInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMcastLeafInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibMcastLeafInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMcastLeafInfo_KEYS.Merge(m, src)
}
func (m *L2FibMcastLeafInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibMcastLeafInfo_KEYS.Size(m)
}
func (m *L2FibMcastLeafInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMcastLeafInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMcastLeafInfo_KEYS proto.InternalMessageInfo

func (m *L2FibMcastLeafInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibMcastLeafInfo_KEYS) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *L2FibMcastLeafInfo_KEYS) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *L2FibMcastLeafInfo_KEYS) GetBdGroupName() string {
	if m != nil {
		return m.BdGroupName
	}
	return ""
}

func (m *L2FibMcastLeafInfo_KEYS) GetBdName() string {
	if m != nil {
		return m.BdName
	}
	return ""
}

type L2FibBaseInfo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibBaseInfo) Reset()         { *m = L2FibBaseInfo{} }
func (m *L2FibBaseInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibBaseInfo) ProtoMessage()    {}
func (*L2FibBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96583c98adae0568, []int{1}
}

func (m *L2FibBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibBaseInfo.Unmarshal(m, b)
}
func (m *L2FibBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibBaseInfo.Marshal(b, m, deterministic)
}
func (m *L2FibBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibBaseInfo.Merge(m, src)
}
func (m *L2FibBaseInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibBaseInfo.Size(m)
}
func (m *L2FibBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibBaseInfo proto.InternalMessageInfo

type L2FibPrefixInfo struct {
	Proto                uint32   `protobuf:"varint,1,opt,name=proto,proto3" json:"proto,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	Prefix               []byte   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibPrefixInfo) Reset()         { *m = L2FibPrefixInfo{} }
func (m *L2FibPrefixInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibPrefixInfo) ProtoMessage()    {}
func (*L2FibPrefixInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96583c98adae0568, []int{2}
}

func (m *L2FibPrefixInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPrefixInfo.Unmarshal(m, b)
}
func (m *L2FibPrefixInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPrefixInfo.Marshal(b, m, deterministic)
}
func (m *L2FibPrefixInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPrefixInfo.Merge(m, src)
}
func (m *L2FibPrefixInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibPrefixInfo.Size(m)
}
func (m *L2FibPrefixInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPrefixInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPrefixInfo proto.InternalMessageInfo

func (m *L2FibPrefixInfo) GetProto() uint32 {
	if m != nil {
		return m.Proto
	}
	return 0
}

func (m *L2FibPrefixInfo) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *L2FibPrefixInfo) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

type L2FibMcastBagCounters struct {
	Packets              uint64   `protobuf:"varint,1,opt,name=packets,proto3" json:"packets,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibMcastBagCounters) Reset()         { *m = L2FibMcastBagCounters{} }
func (m *L2FibMcastBagCounters) String() string { return proto.CompactTextString(m) }
func (*L2FibMcastBagCounters) ProtoMessage()    {}
func (*L2FibMcastBagCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_96583c98adae0568, []int{3}
}

func (m *L2FibMcastBagCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMcastBagCounters.Unmarshal(m, b)
}
func (m *L2FibMcastBagCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMcastBagCounters.Marshal(b, m, deterministic)
}
func (m *L2FibMcastBagCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMcastBagCounters.Merge(m, src)
}
func (m *L2FibMcastBagCounters) XXX_Size() int {
	return xxx_messageInfo_L2FibMcastBagCounters.Size(m)
}
func (m *L2FibMcastBagCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMcastBagCounters.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMcastBagCounters proto.InternalMessageInfo

func (m *L2FibMcastBagCounters) GetPackets() uint64 {
	if m != nil {
		return m.Packets
	}
	return 0
}

func (m *L2FibMcastBagCounters) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

type L2FibMcastStatsBag struct {
	MulticastForwardStat     *L2FibMcastBagCounters `protobuf:"bytes,1,opt,name=multicast_forward_stat,json=multicastForwardStat,proto3" json:"multicast_forward_stat,omitempty"`
	ReceivedStat             *L2FibMcastBagCounters `protobuf:"bytes,2,opt,name=received_stat,json=receivedStat,proto3" json:"received_stat,omitempty"`
	Punt                     *L2FibMcastBagCounters `protobuf:"bytes,3,opt,name=punt,proto3" json:"punt,omitempty"`
	Drop                     *L2FibMcastBagCounters `protobuf:"bytes,4,opt,name=drop,proto3" json:"drop,omitempty"`
	MulticastCoreForwardStat *L2FibMcastBagCounters `protobuf:"bytes,5,opt,name=multicast_core_forward_stat,json=multicastCoreForwardStat,proto3" json:"multicast_core_forward_stat,omitempty"`
	CoreReceivedStat         *L2FibMcastBagCounters `protobuf:"bytes,6,opt,name=core_received_stat,json=coreReceivedStat,proto3" json:"core_received_stat,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}               `json:"-"`
	XXX_unrecognized         []byte                 `json:"-"`
	XXX_sizecache            int32                  `json:"-"`
}

func (m *L2FibMcastStatsBag) Reset()         { *m = L2FibMcastStatsBag{} }
func (m *L2FibMcastStatsBag) String() string { return proto.CompactTextString(m) }
func (*L2FibMcastStatsBag) ProtoMessage()    {}
func (*L2FibMcastStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_96583c98adae0568, []int{4}
}

func (m *L2FibMcastStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMcastStatsBag.Unmarshal(m, b)
}
func (m *L2FibMcastStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMcastStatsBag.Marshal(b, m, deterministic)
}
func (m *L2FibMcastStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMcastStatsBag.Merge(m, src)
}
func (m *L2FibMcastStatsBag) XXX_Size() int {
	return xxx_messageInfo_L2FibMcastStatsBag.Size(m)
}
func (m *L2FibMcastStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMcastStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMcastStatsBag proto.InternalMessageInfo

func (m *L2FibMcastStatsBag) GetMulticastForwardStat() *L2FibMcastBagCounters {
	if m != nil {
		return m.MulticastForwardStat
	}
	return nil
}

func (m *L2FibMcastStatsBag) GetReceivedStat() *L2FibMcastBagCounters {
	if m != nil {
		return m.ReceivedStat
	}
	return nil
}

func (m *L2FibMcastStatsBag) GetPunt() *L2FibMcastBagCounters {
	if m != nil {
		return m.Punt
	}
	return nil
}

func (m *L2FibMcastStatsBag) GetDrop() *L2FibMcastBagCounters {
	if m != nil {
		return m.Drop
	}
	return nil
}

func (m *L2FibMcastStatsBag) GetMulticastCoreForwardStat() *L2FibMcastBagCounters {
	if m != nil {
		return m.MulticastCoreForwardStat
	}
	return nil
}

func (m *L2FibMcastStatsBag) GetCoreReceivedStat() *L2FibMcastBagCounters {
	if m != nil {
		return m.CoreReceivedStat
	}
	return nil
}

type L2FibPlatformStatsUnionBag struct {
	DataType             string              `protobuf:"bytes,1,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Mulicast             *L2FibMcastStatsBag `protobuf:"bytes,2,opt,name=mulicast,proto3" json:"mulicast,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *L2FibPlatformStatsUnionBag) Reset()         { *m = L2FibPlatformStatsUnionBag{} }
func (m *L2FibPlatformStatsUnionBag) String() string { return proto.CompactTextString(m) }
func (*L2FibPlatformStatsUnionBag) ProtoMessage()    {}
func (*L2FibPlatformStatsUnionBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_96583c98adae0568, []int{5}
}

func (m *L2FibPlatformStatsUnionBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPlatformStatsUnionBag.Unmarshal(m, b)
}
func (m *L2FibPlatformStatsUnionBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPlatformStatsUnionBag.Marshal(b, m, deterministic)
}
func (m *L2FibPlatformStatsUnionBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPlatformStatsUnionBag.Merge(m, src)
}
func (m *L2FibPlatformStatsUnionBag) XXX_Size() int {
	return xxx_messageInfo_L2FibPlatformStatsUnionBag.Size(m)
}
func (m *L2FibPlatformStatsUnionBag) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPlatformStatsUnionBag.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPlatformStatsUnionBag proto.InternalMessageInfo

func (m *L2FibPlatformStatsUnionBag) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *L2FibPlatformStatsUnionBag) GetMulicast() *L2FibMcastStatsBag {
	if m != nil {
		return m.Mulicast
	}
	return nil
}

type L2FibPlatformStatsBag struct {
	ForwardStat          *L2FibPlatformStatsUnionBag `protobuf:"bytes,1,opt,name=forward_stat,json=forwardStat,proto3" json:"forward_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *L2FibPlatformStatsBag) Reset()         { *m = L2FibPlatformStatsBag{} }
func (m *L2FibPlatformStatsBag) String() string { return proto.CompactTextString(m) }
func (*L2FibPlatformStatsBag) ProtoMessage()    {}
func (*L2FibPlatformStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_96583c98adae0568, []int{6}
}

func (m *L2FibPlatformStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPlatformStatsBag.Unmarshal(m, b)
}
func (m *L2FibPlatformStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPlatformStatsBag.Marshal(b, m, deterministic)
}
func (m *L2FibPlatformStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPlatformStatsBag.Merge(m, src)
}
func (m *L2FibPlatformStatsBag) XXX_Size() int {
	return xxx_messageInfo_L2FibPlatformStatsBag.Size(m)
}
func (m *L2FibPlatformStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPlatformStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPlatformStatsBag proto.InternalMessageInfo

func (m *L2FibPlatformStatsBag) GetForwardStat() *L2FibPlatformStatsUnionBag {
	if m != nil {
		return m.ForwardStat
	}
	return nil
}

type L2FibMcastIrbInfo struct {
	MxidAcInterfaceHandle string   `protobuf:"bytes,1,opt,name=mxid_ac_interface_handle,json=mxidAcInterfaceHandle,proto3" json:"mxid_ac_interface_handle,omitempty"`
	MxidPwId              uint32   `protobuf:"varint,2,opt,name=mxid_pw_id,json=mxidPwId,proto3" json:"mxid_pw_id,omitempty"`
	MxidNextHopAddress    string   `protobuf:"bytes,3,opt,name=mxid_next_hop_address,json=mxidNextHopAddress,proto3" json:"mxid_next_hop_address,omitempty"`
	IrbPlatDataLen        uint32   `protobuf:"varint,4,opt,name=irb_plat_data_len,json=irbPlatDataLen,proto3" json:"irb_plat_data_len,omitempty"`
	IrbPlatData           []uint32 `protobuf:"varint,5,rep,packed,name=irb_plat_data,json=irbPlatData,proto3" json:"irb_plat_data,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *L2FibMcastIrbInfo) Reset()         { *m = L2FibMcastIrbInfo{} }
func (m *L2FibMcastIrbInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibMcastIrbInfo) ProtoMessage()    {}
func (*L2FibMcastIrbInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96583c98adae0568, []int{7}
}

func (m *L2FibMcastIrbInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMcastIrbInfo.Unmarshal(m, b)
}
func (m *L2FibMcastIrbInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMcastIrbInfo.Marshal(b, m, deterministic)
}
func (m *L2FibMcastIrbInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMcastIrbInfo.Merge(m, src)
}
func (m *L2FibMcastIrbInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibMcastIrbInfo.Size(m)
}
func (m *L2FibMcastIrbInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMcastIrbInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMcastIrbInfo proto.InternalMessageInfo

func (m *L2FibMcastIrbInfo) GetMxidAcInterfaceHandle() string {
	if m != nil {
		return m.MxidAcInterfaceHandle
	}
	return ""
}

func (m *L2FibMcastIrbInfo) GetMxidPwId() uint32 {
	if m != nil {
		return m.MxidPwId
	}
	return 0
}

func (m *L2FibMcastIrbInfo) GetMxidNextHopAddress() string {
	if m != nil {
		return m.MxidNextHopAddress
	}
	return ""
}

func (m *L2FibMcastIrbInfo) GetIrbPlatDataLen() uint32 {
	if m != nil {
		return m.IrbPlatDataLen
	}
	return 0
}

func (m *L2FibMcastIrbInfo) GetIrbPlatData() []uint32 {
	if m != nil {
		return m.IrbPlatData
	}
	return nil
}

type L2FibMcastLeafInfo struct {
	MulticastBaseInformation *L2FibBaseInfo         `protobuf:"bytes,50,opt,name=multicast_base_information,json=multicastBaseInformation,proto3" json:"multicast_base_information,omitempty"`
	SourcePrefix             *L2FibPrefixInfo       `protobuf:"bytes,51,opt,name=source_prefix,json=sourcePrefix,proto3" json:"source_prefix,omitempty"`
	DestinationPrefix        *L2FibPrefixInfo       `protobuf:"bytes,52,opt,name=destination_prefix,json=destinationPrefix,proto3" json:"destination_prefix,omitempty"`
	BridgeId                 uint32                 `protobuf:"varint,53,opt,name=bridge_id,json=bridgeId,proto3" json:"bridge_id,omitempty"`
	XidCount                 uint32                 `protobuf:"varint,54,opt,name=xid_count,json=xidCount,proto3" json:"xid_count,omitempty"`
	PlatformDataLength       uint32                 `protobuf:"varint,55,opt,name=platform_data_length,json=platformDataLength,proto3" json:"platform_data_length,omitempty"`
	PlatformData             []byte                 `protobuf:"bytes,56,opt,name=platform_data,json=platformData,proto3" json:"platform_data,omitempty"`
	HardwareInformation      []byte                 `protobuf:"bytes,57,opt,name=hardware_information,json=hardwareInformation,proto3" json:"hardware_information,omitempty"`
	BridgeDomainName         []string               `protobuf:"bytes,58,rep,name=bridge_domain_name,json=bridgeDomainName,proto3" json:"bridge_domain_name,omitempty"`
	ForwardStats             *L2FibPlatformStatsBag `protobuf:"bytes,59,opt,name=forward_stats,json=forwardStats,proto3" json:"forward_stats,omitempty"`
	IrbInfo                  *L2FibMcastIrbInfo     `protobuf:"bytes,60,opt,name=irb_info,json=irbInfo,proto3" json:"irb_info,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}               `json:"-"`
	XXX_unrecognized         []byte                 `json:"-"`
	XXX_sizecache            int32                  `json:"-"`
}

func (m *L2FibMcastLeafInfo) Reset()         { *m = L2FibMcastLeafInfo{} }
func (m *L2FibMcastLeafInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibMcastLeafInfo) ProtoMessage()    {}
func (*L2FibMcastLeafInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96583c98adae0568, []int{8}
}

func (m *L2FibMcastLeafInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMcastLeafInfo.Unmarshal(m, b)
}
func (m *L2FibMcastLeafInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMcastLeafInfo.Marshal(b, m, deterministic)
}
func (m *L2FibMcastLeafInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMcastLeafInfo.Merge(m, src)
}
func (m *L2FibMcastLeafInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibMcastLeafInfo.Size(m)
}
func (m *L2FibMcastLeafInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMcastLeafInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMcastLeafInfo proto.InternalMessageInfo

func (m *L2FibMcastLeafInfo) GetMulticastBaseInformation() *L2FibBaseInfo {
	if m != nil {
		return m.MulticastBaseInformation
	}
	return nil
}

func (m *L2FibMcastLeafInfo) GetSourcePrefix() *L2FibPrefixInfo {
	if m != nil {
		return m.SourcePrefix
	}
	return nil
}

func (m *L2FibMcastLeafInfo) GetDestinationPrefix() *L2FibPrefixInfo {
	if m != nil {
		return m.DestinationPrefix
	}
	return nil
}

func (m *L2FibMcastLeafInfo) GetBridgeId() uint32 {
	if m != nil {
		return m.BridgeId
	}
	return 0
}

func (m *L2FibMcastLeafInfo) GetXidCount() uint32 {
	if m != nil {
		return m.XidCount
	}
	return 0
}

func (m *L2FibMcastLeafInfo) GetPlatformDataLength() uint32 {
	if m != nil {
		return m.PlatformDataLength
	}
	return 0
}

func (m *L2FibMcastLeafInfo) GetPlatformData() []byte {
	if m != nil {
		return m.PlatformData
	}
	return nil
}

func (m *L2FibMcastLeafInfo) GetHardwareInformation() []byte {
	if m != nil {
		return m.HardwareInformation
	}
	return nil
}

func (m *L2FibMcastLeafInfo) GetBridgeDomainName() []string {
	if m != nil {
		return m.BridgeDomainName
	}
	return nil
}

func (m *L2FibMcastLeafInfo) GetForwardStats() *L2FibPlatformStatsBag {
	if m != nil {
		return m.ForwardStats
	}
	return nil
}

func (m *L2FibMcastLeafInfo) GetIrbInfo() *L2FibMcastIrbInfo {
	if m != nil {
		return m.IrbInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*L2FibMcastLeafInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mroutes.l2fib_mroute.l2fib_mcast_leaf_info_KEYS")
	proto.RegisterType((*L2FibBaseInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mroutes.l2fib_mroute.l2fib_base_info")
	proto.RegisterType((*L2FibPrefixInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mroutes.l2fib_mroute.l2fib_prefix_info")
	proto.RegisterType((*L2FibMcastBagCounters)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mroutes.l2fib_mroute.l2fib_mcast_bag_counters")
	proto.RegisterType((*L2FibMcastStatsBag)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mroutes.l2fib_mroute.l2fib_mcast_stats_bag")
	proto.RegisterType((*L2FibPlatformStatsUnionBag)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mroutes.l2fib_mroute.l2fib_platform_stats_union_bag")
	proto.RegisterType((*L2FibPlatformStatsBag)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mroutes.l2fib_mroute.l2fib_platform_stats_bag")
	proto.RegisterType((*L2FibMcastIrbInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mroutes.l2fib_mroute.l2fib_mcast_irb_info")
	proto.RegisterType((*L2FibMcastLeafInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mroutes.l2fib_mroute.l2fib_mcast_leaf_info")
}

func init() { proto.RegisterFile("l2fib_mcast_leaf_info.proto", fileDescriptor_96583c98adae0568) }

var fileDescriptor_96583c98adae0568 = []byte{
	// 907 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x8e, 0x1c, 0x35,
	0x10, 0x56, 0x67, 0xff, 0x6b, 0x67, 0x20, 0x6b, 0x26, 0xa1, 0x95, 0x45, 0x68, 0xd5, 0x5c, 0x16,
	0x09, 0x8d, 0xc8, 0x06, 0x08, 0x7f, 0x97, 0x90, 0x00, 0x19, 0x88, 0x56, 0xab, 0x0e, 0x17, 0x4e,
	0x96, 0xbb, 0xed, 0x9e, 0xb1, 0xe8, 0xb6, 0x5b, 0xb6, 0x27, 0x3b, 0x2b, 0x21, 0x21, 0x2e, 0x70,
	0x42, 0x82, 0x0b, 0x37, 0x10, 0x37, 0x6e, 0xbc, 0x03, 0x2f, 0xc1, 0x6b, 0xf0, 0x0a, 0xc8, 0xe5,
	0xee, 0x9e, 0x9e, 0x90, 0xdc, 0x32, 0x73, 0x19, 0x4d, 0x55, 0xb9, 0xfd, 0x7d, 0x76, 0x95, 0xbf,
	0x2a, 0x38, 0x2e, 0xcf, 0x0a, 0x99, 0xd1, 0x2a, 0x67, 0xd6, 0xd1, 0x52, 0xb0, 0x82, 0x4a, 0x55,
	0xe8, 0x71, 0x6d, 0xb4, 0xd3, 0xe4, 0x3c, 0x97, 0x36, 0xd7, 0x54, 0x6a, 0x4b, 0x17, 0x86, 0x96,
	0x67, 0x4f, 0x6a, 0x45, 0x75, 0x2d, 0xcc, 0x38, 0xfc, 0x2d, 0xb4, 0xb9, 0x64, 0x86, 0x4b, 0x35,
	0x1d, 0x2b, 0xcd, 0x85, 0xc5, 0xdf, 0x71, 0xb3, 0xa1, 0xd1, 0x73, 0x27, 0xec, 0x8a, 0x95, 0xfc,
	0x11, 0xc1, 0xad, 0x67, 0xe2, 0xd1, 0x2f, 0x3f, 0xfd, 0xfa, 0x31, 0x79, 0x15, 0xf6, 0xfc, 0x0e,
	0x54, 0xf2, 0x38, 0x3a, 0x89, 0x4e, 0x0f, 0xd2, 0x5d, 0x6f, 0x4e, 0x38, 0xb9, 0x09, 0xbb, 0x56,
	0xcf, 0x4d, 0x2e, 0xe2, 0x6b, 0xc1, 0x1f, 0x2c, 0x32, 0x82, 0x9d, 0xa9, 0xd1, 0xf3, 0x3a, 0xde,
	0x42, 0x77, 0x30, 0x48, 0x02, 0xc3, 0x8c, 0x53, 0xfc, 0x4f, 0x15, 0xab, 0x44, 0xbc, 0x8d, 0xd1,
	0xc3, 0x8c, 0x7f, 0xee, 0x7d, 0xe7, 0xac, 0x12, 0x1e, 0x2a, 0xe3, 0x21, 0xba, 0x13, 0xb6, 0xcc,
	0xb8, 0x0f, 0x24, 0x47, 0xf0, 0x72, 0x60, 0x98, 0x31, 0x2b, 0x90, 0x5b, 0x52, 0xc0, 0x51, 0x70,
	0xd5, 0x46, 0x14, 0x72, 0x81, 0x4e, 0x0f, 0x8d, 0x77, 0x84, 0x4c, 0x87, 0x69, 0x30, 0xc8, 0x1b,
	0x30, 0x6c, 0x16, 0x95, 0x42, 0x4d, 0xdd, 0x0c, 0xf9, 0x0e, 0xd3, 0x41, 0x70, 0x3e, 0x42, 0x9f,
	0x3f, 0x4d, 0xb0, 0x91, 0xf6, 0x20, 0x6d, 0xac, 0xe4, 0x0b, 0x88, 0xfb, 0x97, 0x93, 0xb1, 0x29,
	0xcd, 0xf5, 0x5c, 0x39, 0x61, 0x2c, 0x89, 0x61, 0xaf, 0x66, 0xf9, 0x37, 0xc2, 0x59, 0x04, 0xdc,
	0x4e, 0x5b, 0xd3, 0x13, 0xc9, 0xae, 0x9c, 0xb0, 0x08, 0xb5, 0x9d, 0x06, 0x23, 0xf9, 0x71, 0x0f,
	0x6e, 0xf4, 0x37, 0xb3, 0x8e, 0x39, 0xeb, 0xb7, 0x24, 0xbf, 0x47, 0x70, 0xb3, 0x9a, 0x97, 0x4e,
	0xa2, 0xbf, 0x49, 0x21, 0xc6, 0x71, 0xe7, 0xc3, 0xb3, 0xd9, 0xf8, 0xc5, 0x66, 0x7d, 0xfc, 0xbc,
	0x43, 0xa5, 0xa3, 0x8e, 0xc7, 0x67, 0x61, 0xcf, 0xc7, 0x8e, 0x39, 0xf2, 0x53, 0x04, 0x43, 0x23,
	0x72, 0x21, 0x9f, 0x88, 0x86, 0xd7, 0xb5, 0x0d, 0xf3, 0x1a, 0xb4, 0xf0, 0xc8, 0xe7, 0x5b, 0xd8,
	0xae, 0xe7, 0xca, 0x61, 0xb2, 0x36, 0xc9, 0x02, 0x51, 0x3d, 0x3a, 0x37, 0xba, 0xc6, 0x1a, 0xde,
	0x28, 0xba, 0x47, 0x25, 0x7f, 0x46, 0x70, 0xbc, 0x2c, 0x96, 0x5c, 0x1b, 0xb1, 0x5a, 0x31, 0x3b,
	0x1b, 0x66, 0x15, 0x77, 0x64, 0xee, 0x6b, 0x23, 0xfa, 0x55, 0xf3, 0x6b, 0x04, 0x04, 0xf9, 0xad,
	0x96, 0xce, 0xee, 0x86, 0x09, 0x5e, 0xf7, 0x1c, 0xd2, 0x5e, 0xf9, 0x24, 0x7f, 0x47, 0xf0, 0x7a,
	0x23, 0x1f, 0x25, 0x73, 0x85, 0x36, 0x55, 0xf3, 0x18, 0xe7, 0x4a, 0x6a, 0x85, 0x4f, 0xf2, 0x18,
	0x0e, 0x38, 0x73, 0x8c, 0xba, 0xab, 0x5a, 0x34, 0xca, 0xb7, 0xef, 0x1d, 0x5f, 0x5d, 0xd5, 0x82,
	0x7c, 0x1f, 0xc1, 0x7e, 0x35, 0x2f, 0xf1, 0xd0, 0xcd, 0x4b, 0x10, 0xeb, 0x3c, 0x4e, 0xa7, 0x14,
	0x69, 0x07, 0x9b, 0xfc, 0x15, 0xb5, 0xd2, 0xf4, 0xd4, 0x19, 0x3c, 0xfb, 0x5f, 0x22, 0x18, 0x3c,
	0x43, 0x46, 0xd4, 0x7a, 0x48, 0x3e, 0xef, 0x12, 0xd3, 0xc3, 0x62, 0x59, 0x0d, 0xc9, 0xbf, 0x11,
	0x8c, 0xfa, 0x87, 0x92, 0x26, 0x0b, 0xb2, 0x7d, 0x17, 0xe2, 0x6a, 0x21, 0x39, 0x65, 0x39, 0x95,
	0x3e, 0x63, 0x05, 0xcb, 0x05, 0x9d, 0x31, 0xc5, 0xcb, 0xf6, 0xe6, 0x6f, 0xf8, 0xf8, 0xbd, 0x7c,
	0xd2, 0x46, 0x1f, 0x62, 0x90, 0xbc, 0x06, 0x80, 0x1f, 0xd6, 0x97, 0xbe, 0x3d, 0x05, 0x59, 0xdf,
	0xf7, 0x9e, 0x8b, 0xcb, 0x09, 0x27, 0xb7, 0x01, 0x3f, 0xa3, 0x4a, 0x2c, 0x1c, 0x9d, 0xe9, 0x9a,
	0x32, 0xce, 0x8d, 0xb0, 0xb6, 0x69, 0x4c, 0xc4, 0x07, 0xcf, 0xc5, 0xc2, 0x3d, 0xd4, 0xf5, 0xbd,
	0x10, 0x21, 0x6f, 0xc2, 0x91, 0x67, 0xe5, 0xcf, 0x43, 0x31, 0xfb, 0xa5, 0x50, 0xf8, 0xca, 0x87,
	0xe9, 0x4b, 0xd2, 0x64, 0x17, 0x25, 0x73, 0x0f, 0x98, 0x63, 0x8f, 0x84, 0xf2, 0x0d, 0x6d, 0x65,
	0x69, 0xbc, 0x73, 0xb2, 0x75, 0x3a, 0x4c, 0x0f, 0x7b, 0xcb, 0x92, 0x7f, 0x9e, 0x12, 0xfc, 0xae,
	0xb5, 0x92, 0xdf, 0x22, 0xb8, 0xb5, 0x7c, 0xc3, 0x5d, 0x5b, 0x33, 0x15, 0x73, 0x52, 0xab, 0xf8,
	0x0c, 0xb3, 0x45, 0xd7, 0x93, 0xad, 0x0e, 0xad, 0xf7, 0x72, 0x3f, 0x61, 0x56, 0x4c, 0x96, 0x04,
	0xc8, 0x0f, 0x11, 0x0c, 0x43, 0x3f, 0x6f, 0x1a, 0x6c, 0x7c, 0x07, 0x29, 0xb1, 0x35, 0x15, 0xd0,
	0xb2, 0x89, 0xa7, 0x83, 0x80, 0x7b, 0x81, 0x2e, 0xf2, 0x73, 0x04, 0x84, 0x0b, 0xeb, 0xa4, 0x42,
	0x62, 0x2d, 0x9b, 0x77, 0x36, 0xc5, 0xe6, 0xa8, 0x07, 0xde, 0x50, 0x3a, 0x86, 0x83, 0xcc, 0x48,
	0x3e, 0xc5, 0x99, 0xe8, 0xdd, 0x50, 0x74, 0xc1, 0x31, 0xe1, 0x3e, 0xe8, 0x6b, 0x0e, 0xb5, 0x27,
	0x7e, 0x2f, 0x04, 0x17, 0x92, 0xdf, 0xf7, 0x36, 0x79, 0x1b, 0x46, 0xdd, 0x53, 0x69, 0xcb, 0xcb,
	0x0f, 0x24, 0x77, 0x71, 0x1d, 0x69, 0x63, 0x4d, 0x89, 0xf9, 0xb1, 0xc4, 0xcf, 0x2e, 0xfd, 0x2f,
	0xe2, 0xf7, 0x71, 0x3a, 0x19, 0xf4, 0x97, 0x92, 0xdb, 0x30, 0x9a, 0x31, 0xc3, 0x2f, 0x99, 0x59,
	0xad, 0xa2, 0x0f, 0x70, 0xed, 0x2b, 0x6d, 0xac, 0x9f, 0xdf, 0xb7, 0x80, 0x34, 0x67, 0xe0, 0xba,
	0x62, 0x52, 0x85, 0xa9, 0xeb, 0xc3, 0x93, 0xad, 0xd3, 0x83, 0xf4, 0x7a, 0x88, 0x3c, 0xc0, 0x00,
	0x0e, 0x66, 0xbe, 0xfb, 0xf7, 0xd5, 0xc4, 0xc6, 0x1f, 0xad, 0x53, 0xc2, 0xff, 0xaf, 0x67, 0xe9,
	0xa0, 0x27, 0x24, 0x96, 0x7c, 0x07, 0xfb, 0xad, 0x78, 0xc4, 0x1f, 0x23, 0x13, 0xbe, 0x4e, 0xf5,
	0x6d, 0xb1, 0xd2, 0x3d, 0x69, 0x32, 0x7f, 0x8b, 0xd9, 0x2e, 0x4e, 0x96, 0x77, 0xfe, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0x51, 0xb2, 0xd2, 0xa9, 0x0b, 0x00, 0x00,
}