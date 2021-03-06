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
// source: ldp_summary_vrf.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_vrfs_vrf_summary

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

type LdpSummaryVrf_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpSummaryVrf_KEYS) Reset()         { *m = LdpSummaryVrf_KEYS{} }
func (m *LdpSummaryVrf_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpSummaryVrf_KEYS) ProtoMessage()    {}
func (*LdpSummaryVrf_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a81e59042442694, []int{0}
}

func (m *LdpSummaryVrf_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpSummaryVrf_KEYS.Unmarshal(m, b)
}
func (m *LdpSummaryVrf_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpSummaryVrf_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpSummaryVrf_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpSummaryVrf_KEYS.Merge(m, src)
}
func (m *LdpSummaryVrf_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpSummaryVrf_KEYS.Size(m)
}
func (m *LdpSummaryVrf_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpSummaryVrf_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpSummaryVrf_KEYS proto.InternalMessageInfo

func (m *LdpSummaryVrf_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpSummaryVrf_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type LdpVrfInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpVrfInfo) Reset()         { *m = LdpVrfInfo{} }
func (m *LdpVrfInfo) String() string { return proto.CompactTextString(m) }
func (*LdpVrfInfo) ProtoMessage()    {}
func (*LdpVrfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a81e59042442694, []int{1}
}

func (m *LdpVrfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpVrfInfo.Unmarshal(m, b)
}
func (m *LdpVrfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpVrfInfo.Marshal(b, m, deterministic)
}
func (m *LdpVrfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpVrfInfo.Merge(m, src)
}
func (m *LdpVrfInfo) XXX_Size() int {
	return xxx_messageInfo_LdpVrfInfo.Size(m)
}
func (m *LdpVrfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpVrfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpVrfInfo proto.InternalMessageInfo

func (m *LdpVrfInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LdpVrfInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LdpSummaryCommon struct {
	AddressFamilies                     string   `protobuf:"bytes,1,opt,name=address_families,json=addressFamilies,proto3" json:"address_families,omitempty"`
	NumberOfIpv4Af                      uint32   `protobuf:"varint,2,opt,name=number_of_ipv4af,json=numberOfIpv4af,proto3" json:"number_of_ipv4af,omitempty"`
	NumberOfIpv6Af                      uint32   `protobuf:"varint,3,opt,name=number_of_ipv6af,json=numberOfIpv6af,proto3" json:"number_of_ipv6af,omitempty"`
	NumberOfNeighbors                   uint32   `protobuf:"varint,4,opt,name=number_of_neighbors,json=numberOfNeighbors,proto3" json:"number_of_neighbors,omitempty"`
	NumberOfNsrSyncedNeighbors          uint32   `protobuf:"varint,5,opt,name=number_of_nsr_synced_neighbors,json=numberOfNsrSyncedNeighbors,proto3" json:"number_of_nsr_synced_neighbors,omitempty"`
	NumberOfGracefulRestartNeighbors    uint32   `protobuf:"varint,6,opt,name=number_of_graceful_restart_neighbors,json=numberOfGracefulRestartNeighbors,proto3" json:"number_of_graceful_restart_neighbors,omitempty"`
	NumberOfDownstreamOnDemandNeighbors uint32   `protobuf:"varint,7,opt,name=number_of_downstream_on_demand_neighbors,json=numberOfDownstreamOnDemandNeighbors,proto3" json:"number_of_downstream_on_demand_neighbors,omitempty"`
	NumberOfIpv4HelloAdj                uint32   `protobuf:"varint,8,opt,name=number_of_ipv4_hello_adj,json=numberOfIpv4HelloAdj,proto3" json:"number_of_ipv4_hello_adj,omitempty"`
	NumberOfIpv6HelloAdj                uint32   `protobuf:"varint,9,opt,name=number_of_ipv6_hello_adj,json=numberOfIpv6HelloAdj,proto3" json:"number_of_ipv6_hello_adj,omitempty"`
	NumberOfIpv4Routes                  uint32   `protobuf:"varint,10,opt,name=number_of_ipv4_routes,json=numberOfIpv4Routes,proto3" json:"number_of_ipv4_routes,omitempty"`
	NumberOfIpv6Routes                  uint32   `protobuf:"varint,11,opt,name=number_of_ipv6_routes,json=numberOfIpv6Routes,proto3" json:"number_of_ipv6_routes,omitempty"`
	NumberOfIpv4LocalAddresses          uint32   `protobuf:"varint,12,opt,name=number_of_ipv4_local_addresses,json=numberOfIpv4LocalAddresses,proto3" json:"number_of_ipv4_local_addresses,omitempty"`
	NumberOfIpv6LocalAddresses          uint32   `protobuf:"varint,13,opt,name=number_of_ipv6_local_addresses,json=numberOfIpv6LocalAddresses,proto3" json:"number_of_ipv6_local_addresses,omitempty"`
	NumberOfLdpInterfaces               uint32   `protobuf:"varint,14,opt,name=number_of_ldp_interfaces,json=numberOfLdpInterfaces,proto3" json:"number_of_ldp_interfaces,omitempty"`
	NumberOfIpv4LdpInterfaces           uint32   `protobuf:"varint,15,opt,name=number_of_ipv4ldp_interfaces,json=numberOfIpv4ldpInterfaces,proto3" json:"number_of_ipv4ldp_interfaces,omitempty"`
	NumberOfIpv6LdpInterfaces           uint32   `protobuf:"varint,16,opt,name=number_of_ipv6ldp_interfaces,json=numberOfIpv6ldpInterfaces,proto3" json:"number_of_ipv6ldp_interfaces,omitempty"`
	NumberOfBindingsIpv4                uint32   `protobuf:"varint,17,opt,name=number_of_bindings_ipv4,json=numberOfBindingsIpv4,proto3" json:"number_of_bindings_ipv4,omitempty"`
	NumberOfBindingsIpv6                uint32   `protobuf:"varint,18,opt,name=number_of_bindings_ipv6,json=numberOfBindingsIpv6,proto3" json:"number_of_bindings_ipv6,omitempty"`
	NumberOfLocalBindingsIpv4           uint32   `protobuf:"varint,19,opt,name=number_of_local_bindings_ipv4,json=numberOfLocalBindingsIpv4,proto3" json:"number_of_local_bindings_ipv4,omitempty"`
	NumberOfLocalBindingsIpv6           uint32   `protobuf:"varint,20,opt,name=number_of_local_bindings_ipv6,json=numberOfLocalBindingsIpv6,proto3" json:"number_of_local_bindings_ipv6,omitempty"`
	NumberOfRemoteBindingsIpv4          uint32   `protobuf:"varint,21,opt,name=number_of_remote_bindings_ipv4,json=numberOfRemoteBindingsIpv4,proto3" json:"number_of_remote_bindings_ipv4,omitempty"`
	NumberOfRemoteBindingsIpv6          uint32   `protobuf:"varint,22,opt,name=number_of_remote_bindings_ipv6,json=numberOfRemoteBindingsIpv6,proto3" json:"number_of_remote_bindings_ipv6,omitempty"`
	XXX_NoUnkeyedLiteral                struct{} `json:"-"`
	XXX_unrecognized                    []byte   `json:"-"`
	XXX_sizecache                       int32    `json:"-"`
}

func (m *LdpSummaryCommon) Reset()         { *m = LdpSummaryCommon{} }
func (m *LdpSummaryCommon) String() string { return proto.CompactTextString(m) }
func (*LdpSummaryCommon) ProtoMessage()    {}
func (*LdpSummaryCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a81e59042442694, []int{2}
}

func (m *LdpSummaryCommon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpSummaryCommon.Unmarshal(m, b)
}
func (m *LdpSummaryCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpSummaryCommon.Marshal(b, m, deterministic)
}
func (m *LdpSummaryCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpSummaryCommon.Merge(m, src)
}
func (m *LdpSummaryCommon) XXX_Size() int {
	return xxx_messageInfo_LdpSummaryCommon.Size(m)
}
func (m *LdpSummaryCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpSummaryCommon.DiscardUnknown(m)
}

var xxx_messageInfo_LdpSummaryCommon proto.InternalMessageInfo

func (m *LdpSummaryCommon) GetAddressFamilies() string {
	if m != nil {
		return m.AddressFamilies
	}
	return ""
}

func (m *LdpSummaryCommon) GetNumberOfIpv4Af() uint32 {
	if m != nil {
		return m.NumberOfIpv4Af
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv6Af() uint32 {
	if m != nil {
		return m.NumberOfIpv6Af
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfNeighbors() uint32 {
	if m != nil {
		return m.NumberOfNeighbors
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfNsrSyncedNeighbors() uint32 {
	if m != nil {
		return m.NumberOfNsrSyncedNeighbors
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfGracefulRestartNeighbors() uint32 {
	if m != nil {
		return m.NumberOfGracefulRestartNeighbors
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfDownstreamOnDemandNeighbors() uint32 {
	if m != nil {
		return m.NumberOfDownstreamOnDemandNeighbors
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv4HelloAdj() uint32 {
	if m != nil {
		return m.NumberOfIpv4HelloAdj
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv6HelloAdj() uint32 {
	if m != nil {
		return m.NumberOfIpv6HelloAdj
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv4Routes() uint32 {
	if m != nil {
		return m.NumberOfIpv4Routes
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv6Routes() uint32 {
	if m != nil {
		return m.NumberOfIpv6Routes
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv4LocalAddresses() uint32 {
	if m != nil {
		return m.NumberOfIpv4LocalAddresses
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv6LocalAddresses() uint32 {
	if m != nil {
		return m.NumberOfIpv6LocalAddresses
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfLdpInterfaces() uint32 {
	if m != nil {
		return m.NumberOfLdpInterfaces
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv4LdpInterfaces() uint32 {
	if m != nil {
		return m.NumberOfIpv4LdpInterfaces
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv6LdpInterfaces() uint32 {
	if m != nil {
		return m.NumberOfIpv6LdpInterfaces
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfBindingsIpv4() uint32 {
	if m != nil {
		return m.NumberOfBindingsIpv4
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfBindingsIpv6() uint32 {
	if m != nil {
		return m.NumberOfBindingsIpv6
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfLocalBindingsIpv4() uint32 {
	if m != nil {
		return m.NumberOfLocalBindingsIpv4
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfLocalBindingsIpv6() uint32 {
	if m != nil {
		return m.NumberOfLocalBindingsIpv6
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfRemoteBindingsIpv4() uint32 {
	if m != nil {
		return m.NumberOfRemoteBindingsIpv4
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfRemoteBindingsIpv6() uint32 {
	if m != nil {
		return m.NumberOfRemoteBindingsIpv6
	}
	return 0
}

type LdpSummaryVrf struct {
	Vrf                  *LdpVrfInfo       `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	Common               *LdpSummaryCommon `protobuf:"bytes,51,opt,name=common,proto3" json:"common,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LdpSummaryVrf) Reset()         { *m = LdpSummaryVrf{} }
func (m *LdpSummaryVrf) String() string { return proto.CompactTextString(m) }
func (*LdpSummaryVrf) ProtoMessage()    {}
func (*LdpSummaryVrf) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a81e59042442694, []int{3}
}

func (m *LdpSummaryVrf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpSummaryVrf.Unmarshal(m, b)
}
func (m *LdpSummaryVrf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpSummaryVrf.Marshal(b, m, deterministic)
}
func (m *LdpSummaryVrf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpSummaryVrf.Merge(m, src)
}
func (m *LdpSummaryVrf) XXX_Size() int {
	return xxx_messageInfo_LdpSummaryVrf.Size(m)
}
func (m *LdpSummaryVrf) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpSummaryVrf.DiscardUnknown(m)
}

var xxx_messageInfo_LdpSummaryVrf proto.InternalMessageInfo

func (m *LdpSummaryVrf) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpSummaryVrf) GetCommon() *LdpSummaryCommon {
	if m != nil {
		return m.Common
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpSummaryVrf_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.summary.ldp_summary_vrf_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.summary.ldp_vrf_info")
	proto.RegisterType((*LdpSummaryCommon)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.summary.ldp_summary_common")
	proto.RegisterType((*LdpSummaryVrf)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.summary.ldp_summary_vrf")
}

func init() { proto.RegisterFile("ldp_summary_vrf.proto", fileDescriptor_9a81e59042442694) }

var fileDescriptor_9a81e59042442694 = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x5b, 0x4f, 0xdb, 0x30,
	0x14, 0xc7, 0x55, 0x60, 0x5c, 0xcc, 0xdd, 0xd0, 0xcd, 0xec, 0x26, 0xd4, 0xed, 0xa1, 0x7b, 0x89,
	0x34, 0xd8, 0xbc, 0x47, 0x06, 0x62, 0x17, 0x34, 0x56, 0xa4, 0xa0, 0x3d, 0xec, 0xc9, 0x72, 0x63,
	0x1b, 0xc2, 0x12, 0xbb, 0xb2, 0xd3, 0x6e, 0x7c, 0xbc, 0x7d, 0x92, 0x7d, 0x95, 0xc9, 0x26, 0x69,
	0xec, 0x44, 0xc0, 0xc3, 0xf6, 0x12, 0xb5, 0xe7, 0xfc, 0xff, 0xbf, 0x73, 0x7a, 0x4e, 0x63, 0x83,
	0x6e, 0xc6, 0x46, 0xc4, 0x8c, 0xf3, 0x9c, 0xea, 0x6b, 0x32, 0xd1, 0x22, 0x1a, 0x69, 0x55, 0x28,
	0x78, 0x90, 0xa4, 0x26, 0x51, 0x24, 0x55, 0x86, 0xfc, 0xd2, 0x24, 0x1f, 0x65, 0x86, 0x58, 0xa1,
	0x1a, 0x71, 0x1d, 0x55, 0xdf, 0x22, 0xa9, 0x18, 0x37, 0xee, 0x19, 0x4d, 0xb4, 0x30, 0xf6, 0x11,
	0x95, 0xa8, 0xde, 0x00, 0x6c, 0x37, 0xc8, 0xe4, 0xcb, 0x87, 0xef, 0xe7, 0xf0, 0x09, 0x58, 0xb2,
	0x06, 0x22, 0x69, 0xce, 0x51, 0x67, 0xb7, 0xd3, 0x5f, 0x8a, 0x17, 0x6d, 0x60, 0x40, 0x73, 0x0e,
	0x77, 0xc0, 0xa2, 0x15, 0xba, 0xdc, 0x8c, 0xcb, 0x2d, 0x4c, 0xb4, 0xb0, 0xa9, 0xde, 0x1e, 0x58,
	0xb1, 0x3c, 0x9b, 0x4e, 0xa5, 0x50, 0x10, 0x82, 0x39, 0x0f, 0xe1, 0x3e, 0xc3, 0x35, 0x30, 0x93,
	0x32, 0x67, 0x5c, 0x8d, 0x67, 0x52, 0xd6, 0xfb, 0x0d, 0x00, 0xf4, 0x9b, 0x48, 0x54, 0x9e, 0x2b,
	0x09, 0x5f, 0x81, 0x0d, 0xca, 0x98, 0xe6, 0xc6, 0x10, 0x41, 0xf3, 0x34, 0x4b, 0xb9, 0x29, 0x31,
	0xeb, 0x65, 0xfc, 0x63, 0x19, 0x86, 0x7d, 0xb0, 0x21, 0xc7, 0xf9, 0x90, 0x6b, 0xa2, 0x04, 0x49,
	0x47, 0x93, 0x37, 0x54, 0x94, 0xfc, 0xb5, 0x9b, 0xf8, 0x99, 0x38, 0x71, 0xd1, 0x96, 0x12, 0x53,
	0x81, 0x66, 0x5b, 0x4a, 0x4c, 0x05, 0x8c, 0xc0, 0x56, 0xad, 0x94, 0x3c, 0xbd, 0xb8, 0x1c, 0x2a,
	0x6d, 0xd0, 0x9c, 0x13, 0x6f, 0x56, 0xe2, 0x41, 0x95, 0x80, 0x47, 0xe0, 0xb9, 0xa7, 0x37, 0x9a,
	0x98, 0x6b, 0x99, 0x70, 0xe6, 0x59, 0x1f, 0x38, 0xeb, 0xe3, 0xa9, 0xd5, 0xe8, 0x73, 0x27, 0xa9,
	0x19, 0x03, 0xf0, 0xb2, 0x66, 0x5c, 0x68, 0x9a, 0x70, 0x31, 0xce, 0x88, 0xe6, 0xa6, 0xa0, 0xba,
	0xf0, 0x48, 0xf3, 0x8e, 0xb4, 0x5b, 0x91, 0x3e, 0x95, 0xca, 0xf8, 0x46, 0x58, 0xf3, 0xbe, 0x81,
	0x7e, 0xcd, 0x63, 0xea, 0xa7, 0x34, 0x85, 0xe6, 0x34, 0x27, 0x4a, 0x12, 0xc6, 0x73, 0x2a, 0xfd,
	0xee, 0x16, 0x1c, 0xf3, 0x45, 0xc5, 0x3c, 0x9e, 0xaa, 0xcf, 0xe4, 0xb1, 0xd3, 0xd6, 0x58, 0x0c,
	0x50, 0x38, 0x6e, 0x72, 0xc9, 0xb3, 0x4c, 0x11, 0xca, 0xae, 0xd0, 0xa2, 0xc3, 0x6c, 0xfb, 0x63,
	0xff, 0x6c, 0x93, 0x87, 0xec, 0xaa, 0xe5, 0xc3, 0x9e, 0x6f, 0xa9, 0xe5, 0xc3, 0x53, 0xdf, 0x6b,
	0xd0, 0x6d, 0xd4, 0xd3, 0x6a, 0x5c, 0x70, 0x83, 0x80, 0x33, 0x41, 0xbf, 0x58, 0xec, 0x32, 0x2d,
	0x0b, 0xae, 0x2c, 0xcb, 0x2d, 0x0b, 0x2e, 0x2d, 0xc1, 0x02, 0x5d, 0x95, 0x4c, 0x25, 0x34, 0x23,
	0xe5, 0x9f, 0x8d, 0x1b, 0xb4, 0x12, 0x2e, 0xd0, 0x96, 0x3b, 0xb5, 0x92, 0xc3, 0x4a, 0xd1, 0x62,
	0xe0, 0x16, 0x63, 0xb5, 0xc5, 0xc0, 0x0d, 0xc6, 0x3b, 0x7f, 0x4a, 0xf6, 0xbd, 0x48, 0x65, 0xc1,
	0xb5, 0xa0, 0x09, 0x37, 0x68, 0xcd, 0xb9, 0xbb, 0x95, 0xfb, 0x94, 0x8d, 0x4e, 0xa6, 0x49, 0x78,
	0x00, 0x9e, 0x86, 0x3f, 0xa0, 0x61, 0x5e, 0x77, 0xe6, 0x1d, 0xbf, 0xfd, 0xec, 0x4e, 0x00, 0x6e,
	0x00, 0x36, 0x5a, 0x00, 0x1c, 0x02, 0xde, 0x82, 0x47, 0x35, 0x60, 0x98, 0x4a, 0x96, 0xca, 0x0b,
	0xe3, 0x5a, 0x41, 0x9b, 0xe1, 0x7e, 0x8f, 0xca, 0xa4, 0x6d, 0xe2, 0x76, 0x1b, 0x46, 0xf0, 0x56,
	0x1b, 0x86, 0xef, 0xc1, 0x33, 0x6f, 0x50, 0x6e, 0xce, 0x61, 0xcd, 0xad, 0xb0, 0x5f, 0x37, 0xe7,
	0xa0, 0xf0, 0x3d, 0x04, 0x8c, 0xb6, 0xef, 0x26, 0xe0, 0x70, 0xe1, 0x9a, 0xe7, 0xaa, 0xe0, 0x8d,
	0x26, 0xba, 0xe1, 0xc2, 0x63, 0xa7, 0x09, 0xba, 0xb8, 0x8f, 0x81, 0xd1, 0xc3, 0x7b, 0x18, 0xb8,
	0xf7, 0xa7, 0x03, 0xd6, 0x1b, 0x07, 0x39, 0x24, 0x60, 0x76, 0xa2, 0x05, 0xda, 0xdb, 0xed, 0xf4,
	0x97, 0xf7, 0xbe, 0x46, 0xff, 0x78, 0x55, 0x44, 0xfe, 0xb9, 0x1e, 0x5b, 0x32, 0xfc, 0x01, 0xe6,
	0x6f, 0xce, 0x6a, 0xb4, 0xef, 0x6a, 0x9c, 0xff, 0x97, 0x1a, 0xe1, 0x35, 0x10, 0x97, 0x25, 0x86,
	0xf3, 0xee, 0xc6, 0xdb, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x9d, 0xbc, 0x59, 0x0a, 0x07,
	0x00, 0x00,
}
