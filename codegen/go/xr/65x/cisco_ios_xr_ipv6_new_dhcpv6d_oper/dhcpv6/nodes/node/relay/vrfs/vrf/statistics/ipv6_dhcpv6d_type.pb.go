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
// source: ipv6_dhcpv6d_type.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_vrfs_vrf_statistics

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

type Ipv6Dhcpv6DType_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DType_KEYS) Reset()         { *m = Ipv6Dhcpv6DType_KEYS{} }
func (m *Ipv6Dhcpv6DType_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DType_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DType_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2cb0c06afea8a3, []int{0}
}

func (m *Ipv6Dhcpv6DType_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DType_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DType_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DType_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DType_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DType_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DType_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DType_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DType_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DType_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DType_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DType_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DType_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type Ipv6Dhcpv6DFilteredStats struct {
	ReceivedPackets      uint64   `protobuf:"varint,1,opt,name=received_packets,json=receivedPackets,proto3" json:"received_packets,omitempty"`
	TransmittedPackets   uint64   `protobuf:"varint,2,opt,name=transmitted_packets,json=transmittedPackets,proto3" json:"transmitted_packets,omitempty"`
	DroppedPackets       uint64   `protobuf:"varint,3,opt,name=dropped_packets,json=droppedPackets,proto3" json:"dropped_packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DFilteredStats) Reset()         { *m = Ipv6Dhcpv6DFilteredStats{} }
func (m *Ipv6Dhcpv6DFilteredStats) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DFilteredStats) ProtoMessage()    {}
func (*Ipv6Dhcpv6DFilteredStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2cb0c06afea8a3, []int{1}
}

func (m *Ipv6Dhcpv6DFilteredStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DFilteredStats.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DFilteredStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DFilteredStats.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DFilteredStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DFilteredStats.Merge(m, src)
}
func (m *Ipv6Dhcpv6DFilteredStats) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DFilteredStats.Size(m)
}
func (m *Ipv6Dhcpv6DFilteredStats) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DFilteredStats.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DFilteredStats proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DFilteredStats) GetReceivedPackets() uint64 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *Ipv6Dhcpv6DFilteredStats) GetTransmittedPackets() uint64 {
	if m != nil {
		return m.TransmittedPackets
	}
	return 0
}

func (m *Ipv6Dhcpv6DFilteredStats) GetDroppedPackets() uint64 {
	if m != nil {
		return m.DroppedPackets
	}
	return 0
}

type Ipv6Dhcpv6DType struct {
	Solicit              *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,50,opt,name=solicit,proto3" json:"solicit,omitempty"`
	Advertise            *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,51,opt,name=advertise,proto3" json:"advertise,omitempty"`
	Request              *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,52,opt,name=request,proto3" json:"request,omitempty"`
	Reply                *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,53,opt,name=reply,proto3" json:"reply,omitempty"`
	Confirm              *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,54,opt,name=confirm,proto3" json:"confirm,omitempty"`
	Decline              *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,55,opt,name=decline,proto3" json:"decline,omitempty"`
	Renew                *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,56,opt,name=renew,proto3" json:"renew,omitempty"`
	Rebind               *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,57,opt,name=rebind,proto3" json:"rebind,omitempty"`
	Release              *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,58,opt,name=release,proto3" json:"release,omitempty"`
	Reconfig             *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,59,opt,name=reconfig,proto3" json:"reconfig,omitempty"`
	Inform               *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,60,opt,name=inform,proto3" json:"inform,omitempty"`
	RelayForward         *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,61,opt,name=relay_forward,json=relayForward,proto3" json:"relay_forward,omitempty"`
	RelayReply           *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,62,opt,name=relay_reply,json=relayReply,proto3" json:"relay_reply,omitempty"`
	LeaseQuery           *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,63,opt,name=lease_query,json=leaseQuery,proto3" json:"lease_query,omitempty"`
	LeaseQueryReply      *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,64,opt,name=lease_query_reply,json=leaseQueryReply,proto3" json:"lease_query_reply,omitempty"`
	LeaseQueryDone       *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,65,opt,name=lease_query_done,json=leaseQueryDone,proto3" json:"lease_query_done,omitempty"`
	LeaseQueryData       *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,66,opt,name=lease_query_data,json=leaseQueryData,proto3" json:"lease_query_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Ipv6Dhcpv6DType) Reset()         { *m = Ipv6Dhcpv6DType{} }
func (m *Ipv6Dhcpv6DType) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DType) ProtoMessage()    {}
func (*Ipv6Dhcpv6DType) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2cb0c06afea8a3, []int{2}
}

func (m *Ipv6Dhcpv6DType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DType.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DType.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DType.Merge(m, src)
}
func (m *Ipv6Dhcpv6DType) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DType.Size(m)
}
func (m *Ipv6Dhcpv6DType) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DType.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DType proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DType) GetSolicit() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Solicit
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetAdvertise() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Advertise
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetRequest() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetReply() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetConfirm() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Confirm
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetDecline() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Decline
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetRenew() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Renew
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetRebind() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Rebind
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetRelease() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetReconfig() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Reconfig
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetInform() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Inform
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetRelayForward() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.RelayForward
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetRelayReply() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.RelayReply
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetLeaseQuery() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.LeaseQuery
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetLeaseQueryReply() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.LeaseQueryReply
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetLeaseQueryDone() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.LeaseQueryDone
	}
	return nil
}

func (m *Ipv6Dhcpv6DType) GetLeaseQueryData() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.LeaseQueryData
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DType_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.vrfs.vrf.statistics.ipv6_dhcpv6d_type_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DFilteredStats)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.vrfs.vrf.statistics.ipv6_dhcpv6d_filtered_stats")
	proto.RegisterType((*Ipv6Dhcpv6DType)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.vrfs.vrf.statistics.ipv6_dhcpv6d_type")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_type.proto", fileDescriptor_aa2cb0c06afea8a3) }

var fileDescriptor_aa2cb0c06afea8a3 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x96, 0x4d, 0x6b, 0x14, 0x31,
	0x18, 0xc7, 0x49, 0xd5, 0x76, 0x37, 0xd5, 0x6e, 0x1b, 0x41, 0x23, 0xbd, 0x94, 0xbd, 0x58, 0x2f,
	0x23, 0xb4, 0xba, 0xbe, 0xbf, 0xa2, 0x5e, 0x84, 0x52, 0xd7, 0x93, 0xa7, 0x90, 0x4e, 0x9e, 0xd1,
	0xd0, 0x99, 0x64, 0xfa, 0x24, 0x9d, 0x75, 0x8f, 0x22, 0x22, 0x82, 0x07, 0x2f, 0xe2, 0xdd, 0x4f,
	0xe6, 0x47, 0x91, 0xc9, 0xcc, 0xee, 0x8e, 0x2c, 0x78, 0x2a, 0x33, 0x97, 0xc0, 0x3e, 0x2f, 0xf9,
	0xff, 0xf6, 0xff, 0x24, 0x61, 0xe8, 0x55, 0x9d, 0x17, 0x23, 0xa1, 0x3e, 0xc4, 0x79, 0x31, 0x52,
	0xc2, 0x4f, 0x73, 0x88, 0x72, 0xb4, 0xde, 0xb2, 0x83, 0x58, 0xbb, 0xd8, 0x0a, 0x6d, 0x9d, 0xf8,
	0x88, 0x22, 0x54, 0x19, 0x98, 0xcc, 0x2b, 0x6d, 0x0e, 0x18, 0x55, 0x3f, 0x22, 0x63, 0x15, 0xb8,
	0xb0, 0x46, 0x08, 0xa9, 0x9c, 0x46, 0x05, 0x26, 0xae, 0x5c, 0x22, 0xe7, 0xa5, 0xd7, 0xce, 0xeb,
	0xd8, 0x0d, 0x0f, 0xe9, 0x95, 0x25, 0x29, 0xf1, 0xfa, 0xe5, 0xbb, 0xb7, 0x6c, 0x9b, 0xf6, 0xcb,
	0x5e, 0x61, 0x64, 0x06, 0x9c, 0xec, 0x90, 0xdd, 0xfe, 0xb8, 0x57, 0x06, 0x0e, 0x64, 0x06, 0xec,
	0x1a, 0xed, 0x15, 0x98, 0x54, 0xb9, 0x95, 0x90, 0x5b, 0x2b, 0x30, 0x29, 0x53, 0xc3, 0xdf, 0x84,
	0x6e, 0xff, 0xb3, 0x65, 0xa2, 0x53, 0x0f, 0x08, 0x4a, 0x94, 0xb2, 0x8e, 0xdd, 0xa0, 0x9b, 0x08,
	0x31, 0xe8, 0x02, 0x94, 0xc8, 0x65, 0x7c, 0x0c, 0xde, 0x85, 0xed, 0xcf, 0x8f, 0x07, 0xb3, 0xf8,
	0x61, 0x15, 0x66, 0x37, 0xe9, 0x65, 0x8f, 0xd2, 0xb8, 0x4c, 0x7b, 0xdf, 0xa8, 0x5e, 0x09, 0xd5,
	0xac, 0x91, 0x9a, 0x35, 0x5c, 0xa7, 0x03, 0x85, 0x36, 0xcf, 0x1b, 0xc5, 0xe7, 0x42, 0xf1, 0x46,
	0x1d, 0xae, 0x0b, 0x87, 0x7f, 0x18, 0xdd, 0x5a, 0xfa, 0xdf, 0xec, 0x0b, 0xa1, 0x6b, 0xce, 0xa6,
	0x3a, 0xd6, 0x9e, 0xef, 0xed, 0x90, 0xdd, 0xf5, 0xbd, 0xe3, 0xe8, 0x6c, 0xfd, 0x8e, 0xfe, 0xe3,
	0xcc, 0x78, 0xa6, 0xcd, 0xbe, 0x11, 0xda, 0x97, 0xaa, 0x00, 0xf4, 0xda, 0x01, 0xdf, 0x6f, 0x9f,
	0x64, 0xa1, 0x1e, 0x3c, 0x41, 0x38, 0x39, 0x05, 0xe7, 0xf9, 0xad, 0x0e, 0x3c, 0xa9, 0xb5, 0xd9,
	0x27, 0x42, 0x2f, 0x20, 0xe4, 0xe9, 0x94, 0xdf, 0x6e, 0x9f, 0xa2, 0x52, 0x0e, 0x5e, 0xc4, 0xd6,
	0x24, 0x1a, 0x33, 0x3e, 0xea, 0xc0, 0x8b, 0x5a, 0x3b, 0x70, 0x28, 0x88, 0x53, 0x6d, 0x80, 0xdf,
	0xe9, 0x80, 0xa3, 0xd6, 0xae, 0x67, 0x62, 0x60, 0xc2, 0xef, 0x76, 0x32, 0x13, 0x03, 0x13, 0xf6,
	0x99, 0xd0, 0x55, 0x84, 0x23, 0x6d, 0x14, 0xbf, 0xd7, 0x3e, 0x44, 0x2d, 0x5d, 0xdf, 0x92, 0x14,
	0xa4, 0x03, 0x7e, 0xbf, 0x93, 0x5b, 0x12, 0xb4, 0xd9, 0x57, 0x42, 0x7b, 0x08, 0xe1, 0x9c, 0xbc,
	0xe7, 0x0f, 0xda, 0x07, 0x99, 0x8b, 0x87, 0xb9, 0x68, 0x93, 0x58, 0xcc, 0xf8, 0xc3, 0x0e, 0xe6,
	0x52, 0x49, 0xb3, 0x1f, 0x84, 0x5e, 0x0a, 0x3b, 0x88, 0xc4, 0xe2, 0x44, 0xa2, 0xe2, 0x8f, 0xda,
	0x87, 0xb9, 0x18, 0xda, 0x5e, 0x55, 0x00, 0xec, 0x3b, 0xa1, 0xeb, 0x15, 0x52, 0xf5, 0x9c, 0x3d,
	0x6e, 0x1f, 0x88, 0x86, 0xb6, 0x71, 0x78, 0xd3, 0x4a, 0x9c, 0x70, 0x76, 0xc4, 0xc9, 0x29, 0xe0,
	0x94, 0x3f, 0xe9, 0x00, 0x27, 0xe8, 0xbf, 0x29, 0xe5, 0xd9, 0x2f, 0x42, 0xb7, 0x1a, 0x38, 0xb5,
	0x47, 0x4f, 0xdb, 0x87, 0x1a, 0x2c, 0xa0, 0x2a, 0xa3, 0x7e, 0x12, 0xba, 0xd9, 0x24, 0x53, 0xd6,
	0x00, 0x7f, 0xd6, 0x3e, 0xd8, 0xc6, 0x02, 0xec, 0x85, 0x35, 0xb0, 0xcc, 0x25, 0xbd, 0xe4, 0xcf,
	0xbb, 0xe5, 0x92, 0x5e, 0x1e, 0xad, 0x86, 0x0f, 0xd6, 0xfd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf5, 0x73, 0x6d, 0x3b, 0xcb, 0x0a, 0x00, 0x00,
}
