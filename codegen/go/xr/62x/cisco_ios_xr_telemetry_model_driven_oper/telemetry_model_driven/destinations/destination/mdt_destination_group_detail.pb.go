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
// source: mdt_destination_group_detail.proto

package cisco_ios_xr_telemetry_model_driven_oper_telemetry_model_driven_destinations_destination

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

type MdtDestinationGroupDetail_KEYS struct {
	DestinationId        string   `protobuf:"bytes,1,opt,name=destination_id,json=destinationId,proto3" json:"destination_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MdtDestinationGroupDetail_KEYS) Reset()         { *m = MdtDestinationGroupDetail_KEYS{} }
func (m *MdtDestinationGroupDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*MdtDestinationGroupDetail_KEYS) ProtoMessage()    {}
func (*MdtDestinationGroupDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11db6a3ca922731, []int{0}
}

func (m *MdtDestinationGroupDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdtDestinationGroupDetail_KEYS.Unmarshal(m, b)
}
func (m *MdtDestinationGroupDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdtDestinationGroupDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *MdtDestinationGroupDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdtDestinationGroupDetail_KEYS.Merge(m, src)
}
func (m *MdtDestinationGroupDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_MdtDestinationGroupDetail_KEYS.Size(m)
}
func (m *MdtDestinationGroupDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MdtDestinationGroupDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MdtDestinationGroupDetail_KEYS proto.InternalMessageInfo

func (m *MdtDestinationGroupDetail_KEYS) GetDestinationId() string {
	if m != nil {
		return m.DestinationId
	}
	return ""
}

type MdtDestinationIPAddress struct {
	IpType               string   `protobuf:"bytes,1,opt,name=ip_type,json=ipType,proto3" json:"ip_type,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MdtDestinationIPAddress) Reset()         { *m = MdtDestinationIPAddress{} }
func (m *MdtDestinationIPAddress) String() string { return proto.CompactTextString(m) }
func (*MdtDestinationIPAddress) ProtoMessage()    {}
func (*MdtDestinationIPAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11db6a3ca922731, []int{1}
}

func (m *MdtDestinationIPAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdtDestinationIPAddress.Unmarshal(m, b)
}
func (m *MdtDestinationIPAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdtDestinationIPAddress.Marshal(b, m, deterministic)
}
func (m *MdtDestinationIPAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdtDestinationIPAddress.Merge(m, src)
}
func (m *MdtDestinationIPAddress) XXX_Size() int {
	return xxx_messageInfo_MdtDestinationIPAddress.Size(m)
}
func (m *MdtDestinationIPAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_MdtDestinationIPAddress.DiscardUnknown(m)
}

var xxx_messageInfo_MdtDestinationIPAddress proto.InternalMessageInfo

func (m *MdtDestinationIPAddress) GetIpType() string {
	if m != nil {
		return m.IpType
	}
	return ""
}

func (m *MdtDestinationIPAddress) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *MdtDestinationIPAddress) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type MdtDestination struct {
	Id                    string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubIdStr              string                   `protobuf:"bytes,2,opt,name=sub_id_str,json=subIdStr,proto3" json:"sub_id_str,omitempty"`
	SubId                 []uint64                 `protobuf:"varint,3,rep,packed,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	DestIpAddress         *MdtDestinationIPAddress `protobuf:"bytes,4,opt,name=dest_ip_address,json=destIpAddress,proto3" json:"dest_ip_address,omitempty"`
	DestPort              uint32                   `protobuf:"varint,5,opt,name=dest_port,json=destPort,proto3" json:"dest_port,omitempty"`
	Encoding              string                   `protobuf:"bytes,6,opt,name=encoding,proto3" json:"encoding,omitempty"`
	Transport             string                   `protobuf:"bytes,7,opt,name=transport,proto3" json:"transport,omitempty"`
	Vrf                   string                   `protobuf:"bytes,8,opt,name=vrf,proto3" json:"vrf,omitempty"`
	VrfId                 uint32                   `protobuf:"varint,9,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	State                 string                   `protobuf:"bytes,10,opt,name=state,proto3" json:"state,omitempty"`
	UdpMtu                uint32                   `protobuf:"varint,11,opt,name=udp_mtu,json=udpMtu,proto3" json:"udp_mtu,omitempty"`
	Dscp                  uint32                   `protobuf:"varint,12,opt,name=dscp,proto3" json:"dscp,omitempty"`
	Tls                   uint32                   `protobuf:"varint,13,opt,name=tls,proto3" json:"tls,omitempty"`
	TlsHost               string                   `protobuf:"bytes,14,opt,name=tls_host,json=tlsHost,proto3" json:"tls_host,omitempty"`
	TotalNumOfPacketsSent uint64                   `protobuf:"varint,15,opt,name=total_num_of_packets_sent,json=totalNumOfPacketsSent,proto3" json:"total_num_of_packets_sent,omitempty"`
	TotalNumOfBytesSent   uint64                   `protobuf:"varint,16,opt,name=total_num_of_bytes_sent,json=totalNumOfBytesSent,proto3" json:"total_num_of_bytes_sent,omitempty"`
	LastCollectionTime    uint64                   `protobuf:"varint,17,opt,name=last_collection_time,json=lastCollectionTime,proto3" json:"last_collection_time,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *MdtDestination) Reset()         { *m = MdtDestination{} }
func (m *MdtDestination) String() string { return proto.CompactTextString(m) }
func (*MdtDestination) ProtoMessage()    {}
func (*MdtDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11db6a3ca922731, []int{2}
}

func (m *MdtDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdtDestination.Unmarshal(m, b)
}
func (m *MdtDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdtDestination.Marshal(b, m, deterministic)
}
func (m *MdtDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdtDestination.Merge(m, src)
}
func (m *MdtDestination) XXX_Size() int {
	return xxx_messageInfo_MdtDestination.Size(m)
}
func (m *MdtDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_MdtDestination.DiscardUnknown(m)
}

var xxx_messageInfo_MdtDestination proto.InternalMessageInfo

func (m *MdtDestination) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MdtDestination) GetSubIdStr() string {
	if m != nil {
		return m.SubIdStr
	}
	return ""
}

func (m *MdtDestination) GetSubId() []uint64 {
	if m != nil {
		return m.SubId
	}
	return nil
}

func (m *MdtDestination) GetDestIpAddress() *MdtDestinationIPAddress {
	if m != nil {
		return m.DestIpAddress
	}
	return nil
}

func (m *MdtDestination) GetDestPort() uint32 {
	if m != nil {
		return m.DestPort
	}
	return 0
}

func (m *MdtDestination) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *MdtDestination) GetTransport() string {
	if m != nil {
		return m.Transport
	}
	return ""
}

func (m *MdtDestination) GetVrf() string {
	if m != nil {
		return m.Vrf
	}
	return ""
}

func (m *MdtDestination) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *MdtDestination) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *MdtDestination) GetUdpMtu() uint32 {
	if m != nil {
		return m.UdpMtu
	}
	return 0
}

func (m *MdtDestination) GetDscp() uint32 {
	if m != nil {
		return m.Dscp
	}
	return 0
}

func (m *MdtDestination) GetTls() uint32 {
	if m != nil {
		return m.Tls
	}
	return 0
}

func (m *MdtDestination) GetTlsHost() string {
	if m != nil {
		return m.TlsHost
	}
	return ""
}

func (m *MdtDestination) GetTotalNumOfPacketsSent() uint64 {
	if m != nil {
		return m.TotalNumOfPacketsSent
	}
	return 0
}

func (m *MdtDestination) GetTotalNumOfBytesSent() uint64 {
	if m != nil {
		return m.TotalNumOfBytesSent
	}
	return 0
}

func (m *MdtDestination) GetLastCollectionTime() uint64 {
	if m != nil {
		return m.LastCollectionTime
	}
	return 0
}

type MdtSensorPaths struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	State                bool     `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
	StatusStr            string   `protobuf:"bytes,3,opt,name=status_str,json=statusStr,proto3" json:"status_str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MdtSensorPaths) Reset()         { *m = MdtSensorPaths{} }
func (m *MdtSensorPaths) String() string { return proto.CompactTextString(m) }
func (*MdtSensorPaths) ProtoMessage()    {}
func (*MdtSensorPaths) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11db6a3ca922731, []int{3}
}

func (m *MdtSensorPaths) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdtSensorPaths.Unmarshal(m, b)
}
func (m *MdtSensorPaths) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdtSensorPaths.Marshal(b, m, deterministic)
}
func (m *MdtSensorPaths) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdtSensorPaths.Merge(m, src)
}
func (m *MdtSensorPaths) XXX_Size() int {
	return xxx_messageInfo_MdtSensorPaths.Size(m)
}
func (m *MdtSensorPaths) XXX_DiscardUnknown() {
	xxx_messageInfo_MdtSensorPaths.DiscardUnknown(m)
}

var xxx_messageInfo_MdtSensorPaths proto.InternalMessageInfo

func (m *MdtSensorPaths) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MdtSensorPaths) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

func (m *MdtSensorPaths) GetStatusStr() string {
	if m != nil {
		return m.StatusStr
	}
	return ""
}

type MdtCollectionSysdbGroup struct {
	Path                   string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Cadence                uint64   `protobuf:"varint,2,opt,name=cadence,proto3" json:"cadence,omitempty"`
	TotalGetCount          uint64   `protobuf:"varint,3,opt,name=total_get_count,json=totalGetCount,proto3" json:"total_get_count,omitempty"`
	TotalListCount         uint64   `protobuf:"varint,4,opt,name=total_list_count,json=totalListCount,proto3" json:"total_list_count,omitempty"`
	TotalDatalistCount     uint64   `protobuf:"varint,5,opt,name=total_datalist_count,json=totalDatalistCount,proto3" json:"total_datalist_count,omitempty"`
	TotalFinddataCount     uint64   `protobuf:"varint,6,opt,name=total_finddata_count,json=totalFinddataCount,proto3" json:"total_finddata_count,omitempty"`
	TotalGetBulkCount      uint64   `protobuf:"varint,7,opt,name=total_get_bulk_count,json=totalGetBulkCount,proto3" json:"total_get_bulk_count,omitempty"`
	TotalItemCount         uint64   `protobuf:"varint,8,opt,name=total_item_count,json=totalItemCount,proto3" json:"total_item_count,omitempty"`
	TotalGetErrors         uint64   `protobuf:"varint,9,opt,name=total_get_errors,json=totalGetErrors,proto3" json:"total_get_errors,omitempty"`
	TotalListErrors        uint64   `protobuf:"varint,10,opt,name=total_list_errors,json=totalListErrors,proto3" json:"total_list_errors,omitempty"`
	TotalDatalistErrors    uint64   `protobuf:"varint,11,opt,name=total_datalist_errors,json=totalDatalistErrors,proto3" json:"total_datalist_errors,omitempty"`
	TotalFinddataErrors    uint64   `protobuf:"varint,12,opt,name=total_finddata_errors,json=totalFinddataErrors,proto3" json:"total_finddata_errors,omitempty"`
	TotalGetBulkErrors     uint64   `protobuf:"varint,13,opt,name=total_get_bulk_errors,json=totalGetBulkErrors,proto3" json:"total_get_bulk_errors,omitempty"`
	TotalEncodeErrors      uint64   `protobuf:"varint,14,opt,name=total_encode_errors,json=totalEncodeErrors,proto3" json:"total_encode_errors,omitempty"`
	TotalEncodeNotready    uint64   `protobuf:"varint,15,opt,name=total_encode_notready,json=totalEncodeNotready,proto3" json:"total_encode_notready,omitempty"`
	TotalSendErrors        uint64   `protobuf:"varint,16,opt,name=total_send_errors,json=totalSendErrors,proto3" json:"total_send_errors,omitempty"`
	TotalSendDrops         uint64   `protobuf:"varint,17,opt,name=total_send_drops,json=totalSendDrops,proto3" json:"total_send_drops,omitempty"`
	TotalSentBytes         uint64   `protobuf:"varint,18,opt,name=total_sent_bytes,json=totalSentBytes,proto3" json:"total_sent_bytes,omitempty"`
	TotalSendPackets       uint64   `protobuf:"varint,19,opt,name=total_send_packets,json=totalSendPackets,proto3" json:"total_send_packets,omitempty"`
	TotalSendBytesDropped  uint64   `protobuf:"varint,20,opt,name=total_send_bytes_dropped,json=totalSendBytesDropped,proto3" json:"total_send_bytes_dropped,omitempty"`
	TotalCollections       uint64   `protobuf:"varint,21,opt,name=total_collections,json=totalCollections,proto3" json:"total_collections,omitempty"`
	TotalCollectionsMissed uint64   `protobuf:"varint,22,opt,name=total_collections_missed,json=totalCollectionsMissed,proto3" json:"total_collections_missed,omitempty"`
	MaxCollectionTime      uint64   `protobuf:"varint,23,opt,name=max_collection_time,json=maxCollectionTime,proto3" json:"max_collection_time,omitempty"`
	MinCollectionTime      uint64   `protobuf:"varint,24,opt,name=min_collection_time,json=minCollectionTime,proto3" json:"min_collection_time,omitempty"`
	AvgCollectionTime      uint64   `protobuf:"varint,25,opt,name=avg_collection_time,json=avgCollectionTime,proto3" json:"avg_collection_time,omitempty"`
	CollectionMethod       uint64   `protobuf:"varint,26,opt,name=collection_method,json=collectionMethod,proto3" json:"collection_method,omitempty"`
	Status                 string   `protobuf:"bytes,27,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *MdtCollectionSysdbGroup) Reset()         { *m = MdtCollectionSysdbGroup{} }
func (m *MdtCollectionSysdbGroup) String() string { return proto.CompactTextString(m) }
func (*MdtCollectionSysdbGroup) ProtoMessage()    {}
func (*MdtCollectionSysdbGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11db6a3ca922731, []int{4}
}

func (m *MdtCollectionSysdbGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdtCollectionSysdbGroup.Unmarshal(m, b)
}
func (m *MdtCollectionSysdbGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdtCollectionSysdbGroup.Marshal(b, m, deterministic)
}
func (m *MdtCollectionSysdbGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdtCollectionSysdbGroup.Merge(m, src)
}
func (m *MdtCollectionSysdbGroup) XXX_Size() int {
	return xxx_messageInfo_MdtCollectionSysdbGroup.Size(m)
}
func (m *MdtCollectionSysdbGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MdtCollectionSysdbGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MdtCollectionSysdbGroup proto.InternalMessageInfo

func (m *MdtCollectionSysdbGroup) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MdtCollectionSysdbGroup) GetCadence() uint64 {
	if m != nil {
		return m.Cadence
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalGetCount() uint64 {
	if m != nil {
		return m.TotalGetCount
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalListCount() uint64 {
	if m != nil {
		return m.TotalListCount
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalDatalistCount() uint64 {
	if m != nil {
		return m.TotalDatalistCount
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalFinddataCount() uint64 {
	if m != nil {
		return m.TotalFinddataCount
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalGetBulkCount() uint64 {
	if m != nil {
		return m.TotalGetBulkCount
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalItemCount() uint64 {
	if m != nil {
		return m.TotalItemCount
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalGetErrors() uint64 {
	if m != nil {
		return m.TotalGetErrors
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalListErrors() uint64 {
	if m != nil {
		return m.TotalListErrors
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalDatalistErrors() uint64 {
	if m != nil {
		return m.TotalDatalistErrors
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalFinddataErrors() uint64 {
	if m != nil {
		return m.TotalFinddataErrors
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalGetBulkErrors() uint64 {
	if m != nil {
		return m.TotalGetBulkErrors
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalEncodeErrors() uint64 {
	if m != nil {
		return m.TotalEncodeErrors
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalEncodeNotready() uint64 {
	if m != nil {
		return m.TotalEncodeNotready
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalSendErrors() uint64 {
	if m != nil {
		return m.TotalSendErrors
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalSendDrops() uint64 {
	if m != nil {
		return m.TotalSendDrops
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalSentBytes() uint64 {
	if m != nil {
		return m.TotalSentBytes
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalSendPackets() uint64 {
	if m != nil {
		return m.TotalSendPackets
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalSendBytesDropped() uint64 {
	if m != nil {
		return m.TotalSendBytesDropped
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalCollections() uint64 {
	if m != nil {
		return m.TotalCollections
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetTotalCollectionsMissed() uint64 {
	if m != nil {
		return m.TotalCollectionsMissed
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetMaxCollectionTime() uint64 {
	if m != nil {
		return m.MaxCollectionTime
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetMinCollectionTime() uint64 {
	if m != nil {
		return m.MinCollectionTime
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetAvgCollectionTime() uint64 {
	if m != nil {
		return m.AvgCollectionTime
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetCollectionMethod() uint64 {
	if m != nil {
		return m.CollectionMethod
	}
	return 0
}

func (m *MdtCollectionSysdbGroup) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type MdtCollectionGroup struct {
	Id                      uint64                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Cadence                 uint32                     `protobuf:"varint,2,opt,name=cadence,proto3" json:"cadence,omitempty"`
	TotalCollections        uint32                     `protobuf:"varint,3,opt,name=total_collections,json=totalCollections,proto3" json:"total_collections,omitempty"`
	Encoding                string                     `protobuf:"bytes,4,opt,name=encoding,proto3" json:"encoding,omitempty"`
	LastCollectionStartTime uint64                     `protobuf:"varint,5,opt,name=last_collection_start_time,json=lastCollectionStartTime,proto3" json:"last_collection_start_time,omitempty"`
	LastCollectionEndTime   uint64                     `protobuf:"varint,6,opt,name=last_collection_end_time,json=lastCollectionEndTime,proto3" json:"last_collection_end_time,omitempty"`
	MaxCollectionTime       uint32                     `protobuf:"varint,7,opt,name=max_collection_time,json=maxCollectionTime,proto3" json:"max_collection_time,omitempty"`
	MinCollectionTime       uint32                     `protobuf:"varint,8,opt,name=min_collection_time,json=minCollectionTime,proto3" json:"min_collection_time,omitempty"`
	MinTotalTime            uint32                     `protobuf:"varint,9,opt,name=min_total_time,json=minTotalTime,proto3" json:"min_total_time,omitempty"`
	MaxTotalTime            uint32                     `protobuf:"varint,10,opt,name=max_total_time,json=maxTotalTime,proto3" json:"max_total_time,omitempty"`
	AvgTotalTime            uint32                     `protobuf:"varint,11,opt,name=avg_total_time,json=avgTotalTime,proto3" json:"avg_total_time,omitempty"`
	TotalOtherErrors        uint32                     `protobuf:"varint,12,opt,name=total_other_errors,json=totalOtherErrors,proto3" json:"total_other_errors,omitempty"`
	TotalOnDataInstances    uint32                     `protobuf:"varint,13,opt,name=total_on_data_instances,json=totalOnDataInstances,proto3" json:"total_on_data_instances,omitempty"`
	TotalNotReady           uint32                     `protobuf:"varint,14,opt,name=total_not_ready,json=totalNotReady,proto3" json:"total_not_ready,omitempty"`
	TotalSendErrors         uint32                     `protobuf:"varint,15,opt,name=total_send_errors,json=totalSendErrors,proto3" json:"total_send_errors,omitempty"`
	TotalSendDrops          uint32                     `protobuf:"varint,16,opt,name=total_send_drops,json=totalSendDrops,proto3" json:"total_send_drops,omitempty"`
	CollectionPath          []*MdtSensorPaths          `protobuf:"bytes,17,rep,name=collection_path,json=collectionPath,proto3" json:"collection_path,omitempty"`
	InternalCollectionGroup []*MdtCollectionSysdbGroup `protobuf:"bytes,18,rep,name=internal_collection_group,json=internalCollectionGroup,proto3" json:"internal_collection_group,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                   `json:"-"`
	XXX_unrecognized        []byte                     `json:"-"`
	XXX_sizecache           int32                      `json:"-"`
}

func (m *MdtCollectionGroup) Reset()         { *m = MdtCollectionGroup{} }
func (m *MdtCollectionGroup) String() string { return proto.CompactTextString(m) }
func (*MdtCollectionGroup) ProtoMessage()    {}
func (*MdtCollectionGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11db6a3ca922731, []int{5}
}

func (m *MdtCollectionGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdtCollectionGroup.Unmarshal(m, b)
}
func (m *MdtCollectionGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdtCollectionGroup.Marshal(b, m, deterministic)
}
func (m *MdtCollectionGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdtCollectionGroup.Merge(m, src)
}
func (m *MdtCollectionGroup) XXX_Size() int {
	return xxx_messageInfo_MdtCollectionGroup.Size(m)
}
func (m *MdtCollectionGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MdtCollectionGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MdtCollectionGroup proto.InternalMessageInfo

func (m *MdtCollectionGroup) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MdtCollectionGroup) GetCadence() uint32 {
	if m != nil {
		return m.Cadence
	}
	return 0
}

func (m *MdtCollectionGroup) GetTotalCollections() uint32 {
	if m != nil {
		return m.TotalCollections
	}
	return 0
}

func (m *MdtCollectionGroup) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *MdtCollectionGroup) GetLastCollectionStartTime() uint64 {
	if m != nil {
		return m.LastCollectionStartTime
	}
	return 0
}

func (m *MdtCollectionGroup) GetLastCollectionEndTime() uint64 {
	if m != nil {
		return m.LastCollectionEndTime
	}
	return 0
}

func (m *MdtCollectionGroup) GetMaxCollectionTime() uint32 {
	if m != nil {
		return m.MaxCollectionTime
	}
	return 0
}

func (m *MdtCollectionGroup) GetMinCollectionTime() uint32 {
	if m != nil {
		return m.MinCollectionTime
	}
	return 0
}

func (m *MdtCollectionGroup) GetMinTotalTime() uint32 {
	if m != nil {
		return m.MinTotalTime
	}
	return 0
}

func (m *MdtCollectionGroup) GetMaxTotalTime() uint32 {
	if m != nil {
		return m.MaxTotalTime
	}
	return 0
}

func (m *MdtCollectionGroup) GetAvgTotalTime() uint32 {
	if m != nil {
		return m.AvgTotalTime
	}
	return 0
}

func (m *MdtCollectionGroup) GetTotalOtherErrors() uint32 {
	if m != nil {
		return m.TotalOtherErrors
	}
	return 0
}

func (m *MdtCollectionGroup) GetTotalOnDataInstances() uint32 {
	if m != nil {
		return m.TotalOnDataInstances
	}
	return 0
}

func (m *MdtCollectionGroup) GetTotalNotReady() uint32 {
	if m != nil {
		return m.TotalNotReady
	}
	return 0
}

func (m *MdtCollectionGroup) GetTotalSendErrors() uint32 {
	if m != nil {
		return m.TotalSendErrors
	}
	return 0
}

func (m *MdtCollectionGroup) GetTotalSendDrops() uint32 {
	if m != nil {
		return m.TotalSendDrops
	}
	return 0
}

func (m *MdtCollectionGroup) GetCollectionPath() []*MdtSensorPaths {
	if m != nil {
		return m.CollectionPath
	}
	return nil
}

func (m *MdtCollectionGroup) GetInternalCollectionGroup() []*MdtCollectionSysdbGroup {
	if m != nil {
		return m.InternalCollectionGroup
	}
	return nil
}

type MdtDestinationDetail struct {
	Destination          *MdtDestination       `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	CollectionGroup      []*MdtCollectionGroup `protobuf:"bytes,2,rep,name=collection_group,json=collectionGroup,proto3" json:"collection_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MdtDestinationDetail) Reset()         { *m = MdtDestinationDetail{} }
func (m *MdtDestinationDetail) String() string { return proto.CompactTextString(m) }
func (*MdtDestinationDetail) ProtoMessage()    {}
func (*MdtDestinationDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11db6a3ca922731, []int{6}
}

func (m *MdtDestinationDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdtDestinationDetail.Unmarshal(m, b)
}
func (m *MdtDestinationDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdtDestinationDetail.Marshal(b, m, deterministic)
}
func (m *MdtDestinationDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdtDestinationDetail.Merge(m, src)
}
func (m *MdtDestinationDetail) XXX_Size() int {
	return xxx_messageInfo_MdtDestinationDetail.Size(m)
}
func (m *MdtDestinationDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_MdtDestinationDetail.DiscardUnknown(m)
}

var xxx_messageInfo_MdtDestinationDetail proto.InternalMessageInfo

func (m *MdtDestinationDetail) GetDestination() *MdtDestination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *MdtDestinationDetail) GetCollectionGroup() []*MdtCollectionGroup {
	if m != nil {
		return m.CollectionGroup
	}
	return nil
}

type MdtDestinationGroupDetail struct {
	Id                   string                  `protobuf:"bytes,50,opt,name=id,proto3" json:"id,omitempty"`
	Configured           uint32                  `protobuf:"varint,51,opt,name=configured,proto3" json:"configured,omitempty"`
	Destination          []*MdtDestinationDetail `protobuf:"bytes,52,rep,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MdtDestinationGroupDetail) Reset()         { *m = MdtDestinationGroupDetail{} }
func (m *MdtDestinationGroupDetail) String() string { return proto.CompactTextString(m) }
func (*MdtDestinationGroupDetail) ProtoMessage()    {}
func (*MdtDestinationGroupDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11db6a3ca922731, []int{7}
}

func (m *MdtDestinationGroupDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdtDestinationGroupDetail.Unmarshal(m, b)
}
func (m *MdtDestinationGroupDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdtDestinationGroupDetail.Marshal(b, m, deterministic)
}
func (m *MdtDestinationGroupDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdtDestinationGroupDetail.Merge(m, src)
}
func (m *MdtDestinationGroupDetail) XXX_Size() int {
	return xxx_messageInfo_MdtDestinationGroupDetail.Size(m)
}
func (m *MdtDestinationGroupDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_MdtDestinationGroupDetail.DiscardUnknown(m)
}

var xxx_messageInfo_MdtDestinationGroupDetail proto.InternalMessageInfo

func (m *MdtDestinationGroupDetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MdtDestinationGroupDetail) GetConfigured() uint32 {
	if m != nil {
		return m.Configured
	}
	return 0
}

func (m *MdtDestinationGroupDetail) GetDestination() []*MdtDestinationDetail {
	if m != nil {
		return m.Destination
	}
	return nil
}

func init() {
	proto.RegisterType((*MdtDestinationGroupDetail_KEYS)(nil), "cisco_ios_xr_telemetry_model_driven_oper.telemetry_model_driven.destinations.destination.mdt_destination_group_detail_KEYS")
	proto.RegisterType((*MdtDestinationIPAddress)(nil), "cisco_ios_xr_telemetry_model_driven_oper.telemetry_model_driven.destinations.destination.MdtDestinationIPAddress")
	proto.RegisterType((*MdtDestination)(nil), "cisco_ios_xr_telemetry_model_driven_oper.telemetry_model_driven.destinations.destination.mdt_destination")
	proto.RegisterType((*MdtSensorPaths)(nil), "cisco_ios_xr_telemetry_model_driven_oper.telemetry_model_driven.destinations.destination.mdt_sensor_paths")
	proto.RegisterType((*MdtCollectionSysdbGroup)(nil), "cisco_ios_xr_telemetry_model_driven_oper.telemetry_model_driven.destinations.destination.mdt_collection_sysdb_group")
	proto.RegisterType((*MdtCollectionGroup)(nil), "cisco_ios_xr_telemetry_model_driven_oper.telemetry_model_driven.destinations.destination.mdt_collection_group")
	proto.RegisterType((*MdtDestinationDetail)(nil), "cisco_ios_xr_telemetry_model_driven_oper.telemetry_model_driven.destinations.destination.mdt_destination_detail")
	proto.RegisterType((*MdtDestinationGroupDetail)(nil), "cisco_ios_xr_telemetry_model_driven_oper.telemetry_model_driven.destinations.destination.mdt_destination_group_detail")
}

func init() { proto.RegisterFile("mdt_destination_group_detail.proto", fileDescriptor_d11db6a3ca922731) }

var fileDescriptor_d11db6a3ca922731 = []byte{
	// 1333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xcd, 0x6e, 0x1b, 0x37,
	0x17, 0x85, 0x6c, 0xd9, 0x92, 0x29, 0x5b, 0x92, 0x19, 0xff, 0x8c, 0x9d, 0x7c, 0x1f, 0x1c, 0xa1,
	0x2d, 0x84, 0xa6, 0x70, 0x53, 0x27, 0x4d, 0x02, 0x74, 0xd5, 0xc4, 0x6e, 0xea, 0xb6, 0x4e, 0x0c,
	0x39, 0x8b, 0x16, 0x5d, 0x10, 0x63, 0x91, 0x92, 0xd9, 0x68, 0xc8, 0x29, 0xc9, 0x11, 0xec, 0x7d,
	0x77, 0xdd, 0xb5, 0xab, 0xbe, 0x45, 0x9f, 0xa1, 0xaf, 0xd0, 0x47, 0xe8, 0x73, 0x14, 0x28, 0x78,
	0xc9, 0x99, 0xa1, 0x46, 0x4e, 0x90, 0x8d, 0xb3, 0x9b, 0xb9, 0xf7, 0x1c, 0xf2, 0xf2, 0xf0, 0xf2,
	0x70, 0x06, 0xf5, 0x12, 0x6a, 0x08, 0x65, 0xda, 0x70, 0x11, 0x1b, 0x2e, 0x05, 0x19, 0x2b, 0x99,
	0xa5, 0x84, 0x32, 0x13, 0xf3, 0xc9, 0x7e, 0xaa, 0xa4, 0x91, 0xf8, 0xfb, 0x21, 0xd7, 0x43, 0x49,
	0xb8, 0xd4, 0xe4, 0x52, 0x11, 0xc3, 0x26, 0x2c, 0x61, 0x46, 0x5d, 0x91, 0x44, 0x52, 0x36, 0x21,
	0x54, 0xf1, 0x29, 0x13, 0x44, 0xa6, 0x4c, 0xed, 0x5f, 0x9f, 0xdb, 0x0f, 0xc6, 0xd7, 0xe1, 0x4b,
	0xef, 0x1b, 0x74, 0xf7, 0x6d, 0xf3, 0x93, 0x6f, 0x8f, 0x7e, 0x38, 0xc3, 0x1f, 0xa2, 0x76, 0x08,
	0xe0, 0x34, 0xaa, 0xed, 0xd5, 0xfa, 0x2b, 0x83, 0xb5, 0x20, 0x7a, 0x4c, 0x7b, 0x97, 0x68, 0xfb,
	0x84, 0x9a, 0xc3, 0x20, 0x76, 0xfa, 0x25, 0xa5, 0x8a, 0x69, 0x8d, 0xb7, 0x51, 0x83, 0xa7, 0xc4,
	0x5c, 0xa5, 0xcc, 0x53, 0x97, 0x79, 0xfa, 0xea, 0x2a, 0x65, 0xf8, 0x2e, 0x5a, 0xe5, 0xe9, 0xf4,
	0x21, 0x89, 0x1d, 0x30, 0x5a, 0x80, 0x6c, 0xcb, 0xc6, 0x72, 0xae, 0x83, 0x3c, 0x2a, 0x20, 0x8b,
	0x05, 0xe4, 0x91, 0x87, 0xf4, 0x7e, 0x59, 0x42, 0x9d, 0xca, 0x32, 0x70, 0x1b, 0x2d, 0x14, 0x85,
	0x2e, 0x70, 0x8a, 0xef, 0x20, 0xa4, 0xb3, 0x73, 0xc2, 0x29, 0xd1, 0x46, 0xf9, 0x79, 0x9a, 0x3a,
	0x3b, 0x3f, 0xa6, 0x67, 0x46, 0xe1, 0x4d, 0xb4, 0xec, 0xb2, 0xd1, 0xe2, 0xde, 0x62, 0xbf, 0x3e,
	0x58, 0x82, 0x0c, 0xfe, 0xa3, 0x86, 0x3a, 0x76, 0x50, 0xc2, 0xd3, 0x62, 0xfe, 0xfa, 0x5e, 0xad,
	0xdf, 0x3a, 0xf8, 0x79, 0xff, 0xa6, 0xf6, 0x64, 0xff, 0x0d, 0x22, 0x3a, 0xb9, 0x8f, 0xd3, 0x5c,
	0x97, 0xdb, 0x68, 0x05, 0x4a, 0x4b, 0xa5, 0x32, 0xd1, 0xd2, 0x5e, 0xad, 0xbf, 0x36, 0x68, 0xda,
	0xc0, 0xa9, 0x54, 0x06, 0xef, 0xa2, 0x26, 0x13, 0x43, 0x49, 0xb9, 0x18, 0x47, 0xcb, 0x6e, 0xad,
	0xf9, 0x3b, 0xbe, 0x83, 0x56, 0x8c, 0x8a, 0x85, 0x06, 0x62, 0x03, 0x92, 0x65, 0x00, 0x77, 0xd1,
	0xe2, 0x54, 0x8d, 0xa2, 0x26, 0xc4, 0xed, 0xa3, 0xd5, 0x66, 0xaa, 0x46, 0x56, 0x9b, 0x15, 0x98,
	0x65, 0x69, 0xaa, 0x46, 0xc7, 0x14, 0x6f, 0xa0, 0x25, 0x6d, 0x62, 0xc3, 0x22, 0x04, 0x50, 0xf7,
	0x62, 0x77, 0x3a, 0xa3, 0x29, 0x49, 0x4c, 0x16, 0xb5, 0x00, 0xbd, 0x9c, 0xd1, 0xf4, 0xc4, 0x64,
	0x18, 0xa3, 0x3a, 0xd5, 0xc3, 0x34, 0x5a, 0x85, 0x28, 0x3c, 0xdb, 0xb9, 0xcc, 0x44, 0x47, 0x6b,
	0x10, 0xb2, 0x8f, 0x78, 0x07, 0x35, 0xcd, 0x44, 0x93, 0x0b, 0xa9, 0x4d, 0xd4, 0x86, 0x71, 0x1b,
	0x66, 0xa2, 0xbf, 0x96, 0xda, 0xe0, 0x27, 0x68, 0xc7, 0x48, 0x13, 0x4f, 0x88, 0xc8, 0x12, 0x22,
	0x47, 0x24, 0x8d, 0x87, 0xaf, 0x99, 0xd1, 0x44, 0x33, 0x61, 0xa2, 0xce, 0x5e, 0xad, 0x5f, 0x1f,
	0x6c, 0x02, 0xe0, 0x45, 0x96, 0xbc, 0x1c, 0x9d, 0xba, 0xec, 0x19, 0x13, 0x06, 0x3f, 0x44, 0xdb,
	0x33, 0xcc, 0xf3, 0x2b, 0xc3, 0x3c, 0xaf, 0x0b, 0xbc, 0x5b, 0x25, 0xef, 0xa9, 0xcd, 0x01, 0xeb,
	0x3e, 0xda, 0x98, 0xc4, 0xda, 0x90, 0xa1, 0x9c, 0x4c, 0xd8, 0x10, 0x3a, 0xdf, 0xf0, 0x84, 0x45,
	0xeb, 0x40, 0xc1, 0x36, 0xf7, 0xac, 0x48, 0xbd, 0xe2, 0x09, 0xeb, 0xfd, 0x88, 0xba, 0xb6, 0x0b,
	0x35, 0x13, 0x5a, 0x2a, 0x92, 0xc6, 0xe6, 0x42, 0xdb, 0x65, 0xdb, 0x07, 0xdf, 0x88, 0xf0, 0x5c,
	0x2a, 0x67, 0xbb, 0xb0, 0x99, 0x2b, 0xf7, 0x3f, 0x84, 0xec, 0x43, 0xa6, 0xa1, 0x41, 0x5d, 0x97,
	0xaf, 0xb8, 0xc8, 0x99, 0x51, 0xbd, 0xbf, 0x57, 0xd0, 0xae, 0x1d, 0x3d, 0x28, 0x47, 0x5f, 0x69,
	0x7a, 0xee, 0xce, 0xeb, 0xb5, 0xf3, 0x44, 0xa8, 0x31, 0x8c, 0x29, 0x13, 0x43, 0x37, 0x53, 0x7d,
	0x90, 0xbf, 0xe2, 0x8f, 0x50, 0xc7, 0x29, 0x32, 0x66, 0x76, 0xc4, 0x4c, 0x18, 0x98, 0xb0, 0x3e,
	0x58, 0x83, 0xf0, 0x73, 0x66, 0x9e, 0xd9, 0x20, 0xee, 0xa3, 0xae, 0xc3, 0x4d, 0xb8, 0xce, 0x81,
	0x75, 0x00, 0xb6, 0x21, 0xfe, 0x1d, 0xd7, 0x1e, 0x79, 0x1f, 0x6d, 0x38, 0x24, 0x8d, 0x4d, 0x1c,
	0xa0, 0x97, 0x9c, 0x5a, 0x90, 0x3b, 0xf4, 0xa9, 0x0a, 0x63, 0xc4, 0x05, 0xb5, 0x2c, 0xcf, 0x58,
	0x0e, 0x18, 0x5f, 0xf9, 0x94, 0x63, 0x7c, 0x9a, 0x33, 0x6c, 0xd5, 0xe7, 0xd9, 0xe4, 0xb5, 0x67,
	0x34, 0x80, 0xb1, 0x9e, 0x97, 0xfe, 0x34, 0x9b, 0xbc, 0xae, 0x94, 0xcf, 0x0d, 0x4b, 0x3c, 0xb8,
	0x19, 0x94, 0x7f, 0x6c, 0x58, 0x52, 0x41, 0xda, 0xa1, 0x99, 0x52, 0x52, 0x69, 0xe8, 0xf6, 0x1c,
	0xf9, 0x9c, 0x99, 0x23, 0x88, 0xe2, 0x8f, 0xd1, 0x7a, 0x20, 0x89, 0x87, 0x22, 0x80, 0x76, 0x0a,
	0x4d, 0x3c, 0xf6, 0x00, 0x6d, 0x56, 0x44, 0xf1, 0xf8, 0x56, 0xd0, 0x76, 0xb9, 0x2a, 0x55, 0x4e,
	0x21, 0x8b, 0xe7, 0xac, 0x06, 0x9c, 0x5c, 0x17, 0xcf, 0xf9, 0x2c, 0xe7, 0x14, 0xc2, 0x78, 0xce,
	0x5a, 0xa0, 0xa5, 0x57, 0xc6, 0x53, 0xf6, 0x91, 0x1b, 0x89, 0x80, 0x2d, 0xb0, 0x9c, 0xd0, 0x0e,
	0xa4, 0x3c, 0x82, 0x4c, 0xb5, 0x2c, 0x8f, 0x17, 0xd2, 0x28, 0x16, 0xd3, 0x2b, 0x7f, 0xf2, 0x6e,
	0x05, 0x8c, 0x17, 0x3e, 0x55, 0x4a, 0xa5, 0x99, 0xa0, 0xf9, 0x0c, 0xdd, 0x40, 0xaa, 0x33, 0x26,
	0xa8, 0x1f, 0xbf, 0xd8, 0x00, 0xc0, 0x52, 0x25, 0x53, 0xed, 0x4f, 0x5a, 0xbb, 0x80, 0x1e, 0xda,
	0xe8, 0x0c, 0xd2, 0xb8, 0xb3, 0x1c, 0xe1, 0x59, 0xa4, 0x81, 0x53, 0x8c, 0x3f, 0x41, 0x38, 0x18,
	0xd3, 0xfb, 0x45, 0x74, 0x0b, 0xb0, 0xdd, 0x62, 0x54, 0xef, 0x14, 0xf8, 0x31, 0x8a, 0x02, 0xb4,
	0xf3, 0x08, 0x5b, 0x47, 0xca, 0x68, 0xb4, 0x11, 0xd8, 0x8b, 0xe5, 0xc0, 0xf8, 0x87, 0x2e, 0x89,
	0xef, 0xe5, 0xcb, 0x2c, 0x8f, 0xa6, 0x8e, 0x36, 0x83, 0x59, 0x4a, 0x9b, 0xd0, 0xf8, 0x49, 0x3e,
	0x4b, 0x00, 0x26, 0x09, 0xd7, 0x9a, 0xd1, 0x68, 0x0b, 0x38, 0x5b, 0x55, 0xce, 0x09, 0x64, 0xed,
	0x8e, 0x25, 0xf1, 0xe5, 0x9c, 0x1d, 0x6d, 0xbb, 0x1d, 0x4b, 0xe2, 0xcb, 0x59, 0x37, 0x02, 0x3c,
	0x17, 0x73, 0xf8, 0xc8, 0xe3, 0xb9, 0x98, 0xc7, 0xc7, 0xd3, 0xf1, 0x1c, 0x7e, 0xc7, 0xe1, 0xe3,
	0xe9, 0xb8, 0x82, 0xbf, 0x87, 0xd6, 0x03, 0x6c, 0xc2, 0xcc, 0x85, 0xa4, 0xd1, 0xae, 0x5b, 0x76,
	0x99, 0x38, 0x81, 0x38, 0xde, 0x42, 0xcb, 0xce, 0xca, 0xa2, 0xdb, 0xee, 0xfe, 0x77, 0x6f, 0xbd,
	0x7f, 0x1b, 0x68, 0xa3, 0xe2, 0x6a, 0xce, 0xcf, 0xca, 0xeb, 0xbb, 0x0e, 0xd7, 0x77, 0xc5, 0xcb,
	0xd6, 0x4a, 0x2f, 0xbb, 0x56, 0xfe, 0x45, 0xc0, 0xcc, 0xcb, 0x1f, 0xde, 0x8b, 0xf5, 0xca, 0xbd,
	0xf8, 0x05, 0xda, 0xad, 0x1a, 0xbe, 0x36, 0xb1, 0x32, 0x4e, 0x07, 0x67, 0x64, 0xdb, 0xb3, 0xb6,
	0x7f, 0x66, 0xf3, 0xa0, 0xc6, 0x63, 0x14, 0x55, 0xc9, 0xb6, 0x8d, 0x80, 0xea, 0x1c, 0x6d, 0x73,
	0x96, 0x7a, 0x24, 0x68, 0xb1, 0x4d, 0xd7, 0x6c, 0x6b, 0x03, 0x16, 0xf0, 0xee, 0xdb, 0xda, 0xf4,
	0xf8, 0xb9, 0x6d, 0xfd, 0x00, 0xb5, 0x2d, 0xde, 0x49, 0x04, 0x50, 0x77, 0x8b, 0xaf, 0x26, 0x5c,
	0xbc, 0xb2, 0xc1, 0x02, 0x15, 0x5f, 0x86, 0x28, 0xe4, 0x51, 0xf1, 0xe5, 0x0c, 0xca, 0xb6, 0x48,
	0x80, 0x72, 0x77, 0xfc, 0x6a, 0x3c, 0x1d, 0x97, 0xa8, 0xe2, 0xd8, 0x49, 0x73, 0xc1, 0x54, 0x68,
	0x5f, 0xf9, 0x8e, 0xbc, 0xb4, 0x09, 0x7f, 0xf0, 0x3f, 0xcf, 0x2f, 0x67, 0x29, 0xc0, 0x26, 0x09,
	0x17, 0xda, 0xc4, 0x62, 0xc8, 0xf2, 0xef, 0x02, 0xe7, 0xf9, 0x2f, 0x85, 0xf5, 0xc9, 0xe3, 0x3c,
	0x57, 0xde, 0x60, 0x42, 0x1a, 0xe2, 0x9c, 0xa8, 0x0d, 0x70, 0x77, 0x83, 0xbd, 0x90, 0x66, 0xf0,
	0x66, 0x0f, 0xea, 0x00, 0xf2, 0x9d, 0x3c, 0xa8, 0x0b, 0xd0, 0xaa, 0x07, 0xfd, 0x5e, 0x43, 0x9d,
	0x60, 0x07, 0xe0, 0xe6, 0x5d, 0xdf, 0x5b, 0xec, 0xb7, 0x0e, 0x7e, 0xba, 0xb9, 0xef, 0xc2, 0xea,
	0xb7, 0xc5, 0xa0, 0x5d, 0x96, 0x70, 0x6a, 0xef, 0xfb, 0x3f, 0x6b, 0x68, 0x87, 0x0b, 0xc3, 0x94,
	0x98, 0x39, 0x0d, 0xee, 0x44, 0x45, 0x18, 0xea, 0x33, 0x37, 0x5b, 0xdf, 0xf5, 0x5f, 0x27, 0x83,
	0xed, 0xbc, 0xac, 0xb2, 0x39, 0x9f, 0xdb, 0x44, 0xef, 0xaf, 0x05, 0xb4, 0x55, 0xfd, 0x01, 0x71,
	0xbf, 0x1e, 0xf8, 0xd7, 0x1a, 0x6a, 0x05, 0x61, 0xf0, 0x82, 0xd6, 0x01, 0xbf, 0xd9, 0xfa, 0x83,
	0xf7, 0x41, 0x38, 0xbb, 0xfd, 0x13, 0xe8, 0xce, 0x49, 0xba, 0x00, 0x92, 0x8a, 0xf7, 0x26, 0xa9,
	0x13, 0x33, 0xe8, 0x3c, 0x27, 0xe2, 0x3f, 0x35, 0x74, 0xe7, 0x6d, 0x7f, 0x71, 0xde, 0x4c, 0x0f,
	0x8a, 0x7f, 0xa1, 0xff, 0x23, 0x34, 0x94, 0x62, 0xc4, 0xc7, 0x99, 0x62, 0x34, 0x7a, 0x00, 0x2d,
	0x1e, 0x44, 0xf0, 0x6f, 0x15, 0xe9, 0x1f, 0xc2, 0x3a, 0xd3, 0xf7, 0x26, 0xbd, 0xaf, 0x7b, 0x66,
	0x07, 0xce, 0x97, 0xe1, 0x5f, 0xf8, 0xc1, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xd6, 0x2e,
	0xc9, 0x31, 0x0f, 0x00, 0x00,
}
