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
// source: ldp_neighbor_brief_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_default_vrf_neighbor_briefs_neighbor_brief

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

type LdpNeighborBriefInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	LsrId                string   `protobuf:"bytes,2,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,3,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpNeighborBriefInfo_KEYS) Reset()         { *m = LdpNeighborBriefInfo_KEYS{} }
func (m *LdpNeighborBriefInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpNeighborBriefInfo_KEYS) ProtoMessage()    {}
func (*LdpNeighborBriefInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_57a8ad7f03439cf8, []int{0}
}

func (m *LdpNeighborBriefInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNeighborBriefInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpNeighborBriefInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNeighborBriefInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpNeighborBriefInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNeighborBriefInfo_KEYS.Merge(m, src)
}
func (m *LdpNeighborBriefInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpNeighborBriefInfo_KEYS.Size(m)
}
func (m *LdpNeighborBriefInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNeighborBriefInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNeighborBriefInfo_KEYS proto.InternalMessageInfo

func (m *LdpNeighborBriefInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpNeighborBriefInfo_KEYS) GetLsrId() string {
	if m != nil {
		return m.LsrId
	}
	return ""
}

func (m *LdpNeighborBriefInfo_KEYS) GetLabelSpaceId() uint32 {
	if m != nil {
		return m.LabelSpaceId
	}
	return 0
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
	return fileDescriptor_57a8ad7f03439cf8, []int{1}
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

type LdpNeighborBriefAfInfo struct {
	AddressFamily        string   `protobuf:"bytes,1,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	NumOfNbrDiscovery    uint32   `protobuf:"varint,2,opt,name=num_of_nbr_discovery,json=numOfNbrDiscovery,proto3" json:"num_of_nbr_discovery,omitempty"`
	NumOfNbrAddresses    uint32   `protobuf:"varint,3,opt,name=num_of_nbr_addresses,json=numOfNbrAddresses,proto3" json:"num_of_nbr_addresses,omitempty"`
	NumOfNbrLbl          uint32   `protobuf:"varint,4,opt,name=num_of_nbr_lbl,json=numOfNbrLbl,proto3" json:"num_of_nbr_lbl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpNeighborBriefAfInfo) Reset()         { *m = LdpNeighborBriefAfInfo{} }
func (m *LdpNeighborBriefAfInfo) String() string { return proto.CompactTextString(m) }
func (*LdpNeighborBriefAfInfo) ProtoMessage()    {}
func (*LdpNeighborBriefAfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_57a8ad7f03439cf8, []int{2}
}

func (m *LdpNeighborBriefAfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNeighborBriefAfInfo.Unmarshal(m, b)
}
func (m *LdpNeighborBriefAfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNeighborBriefAfInfo.Marshal(b, m, deterministic)
}
func (m *LdpNeighborBriefAfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNeighborBriefAfInfo.Merge(m, src)
}
func (m *LdpNeighborBriefAfInfo) XXX_Size() int {
	return xxx_messageInfo_LdpNeighborBriefAfInfo.Size(m)
}
func (m *LdpNeighborBriefAfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNeighborBriefAfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNeighborBriefAfInfo proto.InternalMessageInfo

func (m *LdpNeighborBriefAfInfo) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *LdpNeighborBriefAfInfo) GetNumOfNbrDiscovery() uint32 {
	if m != nil {
		return m.NumOfNbrDiscovery
	}
	return 0
}

func (m *LdpNeighborBriefAfInfo) GetNumOfNbrAddresses() uint32 {
	if m != nil {
		return m.NumOfNbrAddresses
	}
	return 0
}

func (m *LdpNeighborBriefAfInfo) GetNumOfNbrLbl() uint32 {
	if m != nil {
		return m.NumOfNbrLbl
	}
	return 0
}

type LdpNeighborBriefInfo struct {
	Vrf                   *LdpVrfInfo               `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	IsGracefulRestartable bool                      `protobuf:"varint,51,opt,name=is_graceful_restartable,json=isGracefulRestartable,proto3" json:"is_graceful_restartable,omitempty"`
	NsrState              string                    `protobuf:"bytes,52,opt,name=nsr_state,json=nsrState,proto3" json:"nsr_state,omitempty"`
	UpTimeSeconds         uint32                    `protobuf:"varint,53,opt,name=up_time_seconds,json=upTimeSeconds,proto3" json:"up_time_seconds,omitempty"`
	NbrBrAfInfo           []*LdpNeighborBriefAfInfo `protobuf:"bytes,54,rep,name=nbr_br_af_info,json=nbrBrAfInfo,proto3" json:"nbr_br_af_info,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                  `json:"-"`
	XXX_unrecognized      []byte                    `json:"-"`
	XXX_sizecache         int32                     `json:"-"`
}

func (m *LdpNeighborBriefInfo) Reset()         { *m = LdpNeighborBriefInfo{} }
func (m *LdpNeighborBriefInfo) String() string { return proto.CompactTextString(m) }
func (*LdpNeighborBriefInfo) ProtoMessage()    {}
func (*LdpNeighborBriefInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_57a8ad7f03439cf8, []int{3}
}

func (m *LdpNeighborBriefInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNeighborBriefInfo.Unmarshal(m, b)
}
func (m *LdpNeighborBriefInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNeighborBriefInfo.Marshal(b, m, deterministic)
}
func (m *LdpNeighborBriefInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNeighborBriefInfo.Merge(m, src)
}
func (m *LdpNeighborBriefInfo) XXX_Size() int {
	return xxx_messageInfo_LdpNeighborBriefInfo.Size(m)
}
func (m *LdpNeighborBriefInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNeighborBriefInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNeighborBriefInfo proto.InternalMessageInfo

func (m *LdpNeighborBriefInfo) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpNeighborBriefInfo) GetIsGracefulRestartable() bool {
	if m != nil {
		return m.IsGracefulRestartable
	}
	return false
}

func (m *LdpNeighborBriefInfo) GetNsrState() string {
	if m != nil {
		return m.NsrState
	}
	return ""
}

func (m *LdpNeighborBriefInfo) GetUpTimeSeconds() uint32 {
	if m != nil {
		return m.UpTimeSeconds
	}
	return 0
}

func (m *LdpNeighborBriefInfo) GetNbrBrAfInfo() []*LdpNeighborBriefAfInfo {
	if m != nil {
		return m.NbrBrAfInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpNeighborBriefInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.neighbor_briefs.neighbor_brief.ldp_neighbor_brief_info_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.neighbor_briefs.neighbor_brief.ldp_vrf_info")
	proto.RegisterType((*LdpNeighborBriefAfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.neighbor_briefs.neighbor_brief.ldp_neighbor_brief_af_info")
	proto.RegisterType((*LdpNeighborBriefInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.neighbor_briefs.neighbor_brief.ldp_neighbor_brief_info")
}

func init() { proto.RegisterFile("ldp_neighbor_brief_info.proto", fileDescriptor_57a8ad7f03439cf8) }

var fileDescriptor_57a8ad7f03439cf8 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0x4d, 0x57, 0x5a, 0xa5, 0xc9, 0x98, 0x58, 0xa9, 0xd9, 0x07, 0x84, 0xec, 0x03, 0x3f,
	0x79, 0x90, 0x6e, 0x7d, 0xef, 0xd8, 0x07, 0x61, 0xa3, 0x03, 0x67, 0x2f, 0x7d, 0xba, 0x48, 0x96,
	0x94, 0x09, 0x64, 0xc9, 0xdc, 0x6b, 0x87, 0xe6, 0xaf, 0xec, 0x47, 0xec, 0x87, 0xec, 0x57, 0x0d,
	0x39, 0xce, 0x96, 0x95, 0xf6, 0xb1, 0x2f, 0xc6, 0xf7, 0x1c, 0xdd, 0x73, 0xe4, 0x73, 0xaf, 0xd9,
	0x73, 0xa7, 0x6a, 0xf0, 0xda, 0x2e, 0x7f, 0xc8, 0x80, 0x20, 0xd1, 0x6a, 0x03, 0xd6, 0x9b, 0x90,
	0xd7, 0x18, 0x9a, 0xc0, 0xaf, 0x4a, 0x4b, 0x65, 0x00, 0x1b, 0x08, 0xae, 0x11, 0xaa, 0xda, 0x11,
	0xc4, 0x86, 0x50, 0x6b, 0xcc, 0xb7, 0x55, 0xee, 0x83, 0xd2, 0xd4, 0x3d, 0x73, 0xa5, 0x8d, 0x68,
	0x5d, 0x03, 0x2b, 0x34, 0xf9, 0xff, 0xaa, 0x74, 0xa3, 0x9e, 0x5e, 0xb3, 0x67, 0x77, 0x78, 0xc3,
	0x97, 0x8f, 0x57, 0x0b, 0xfe, 0x94, 0x1d, 0x45, 0x49, 0xf0, 0xa2, 0xd2, 0x69, 0x32, 0x49, 0xb2,
	0xa3, 0xe2, 0x30, 0x02, 0x97, 0xa2, 0xd2, 0xfc, 0x84, 0x1d, 0x38, 0x42, 0xb0, 0x2a, 0xdd, 0xeb,
	0x98, 0x07, 0x8e, 0x70, 0xae, 0xf8, 0x4b, 0x36, 0x76, 0x42, 0x6a, 0x07, 0x54, 0x8b, 0x52, 0x47,
	0x7a, 0x30, 0x49, 0xb2, 0x51, 0x71, 0xdc, 0xa1, 0x8b, 0x08, 0xce, 0xd5, 0x74, 0xc6, 0x8e, 0xa3,
	0xf3, 0x0a, 0x37, 0x76, 0x9c, 0xb3, 0xfd, 0x1d, 0x93, 0xee, 0x9d, 0x8f, 0xd9, 0x5e, 0x2f, 0x3e,
	0x2a, 0xf6, 0xac, 0x9a, 0xfe, 0x4e, 0xd8, 0x93, 0x5b, 0xae, 0x2b, 0x7a, 0x89, 0x57, 0x6c, 0x2c,
	0x94, 0x42, 0x4d, 0x04, 0x46, 0x54, 0xd6, 0xad, 0x7b, 0xb1, 0x51, 0x8f, 0x7e, 0xea, 0x40, 0xfe,
	0x86, 0x3d, 0xf6, 0x6d, 0x05, 0xc1, 0x80, 0x97, 0x08, 0x2a, 0x66, 0xbb, 0xd2, 0xb8, 0xee, 0x7d,
	0x1e, 0xf9, 0xb6, 0xfa, 0x66, 0x2e, 0x25, 0x7e, 0xd8, 0x12, 0x37, 0x1a, 0x7a, 0x31, 0x4d, 0xfd,
	0x67, 0xfd, 0x6d, 0xb8, 0xd8, 0x12, 0xfc, 0x05, 0x1b, 0xef, 0x34, 0x38, 0xe9, 0xd2, 0xfd, 0xee,
	0xe8, 0x70, 0x7b, 0xf4, 0xab, 0x74, 0xd3, 0x5f, 0x03, 0x76, 0x7a, 0x47, 0xf6, 0x7c, 0xcd, 0x06,
	0x2b, 0x34, 0xe9, 0x6c, 0x92, 0x64, 0xc3, 0xd9, 0x32, 0xbf, 0xb7, 0xf9, 0xe7, 0xbb, 0x23, 0x28,
	0xa2, 0x27, 0x3f, 0x67, 0xa7, 0x96, 0x60, 0x89, 0xa2, 0xd4, 0xa6, 0x75, 0x80, 0x9a, 0x1a, 0x81,
	0x8d, 0x90, 0x4e, 0xa7, 0x67, 0x93, 0x24, 0x3b, 0x2c, 0x4e, 0x2c, 0x7d, 0xee, 0xd9, 0xe2, 0x1f,
	0xd9, 0x6d, 0x0a, 0x21, 0x50, 0x23, 0x1a, 0x9d, 0xbe, 0xed, 0x37, 0x85, 0x70, 0x11, 0x6b, 0xfe,
	0x9a, 0x3d, 0x6c, 0x6b, 0x68, 0x6c, 0xa5, 0x81, 0x74, 0x19, 0xbc, 0xa2, 0xf4, 0x5d, 0x97, 0xc8,
	0xa8, 0xad, 0xbf, 0xdb, 0x4a, 0x2f, 0x36, 0x20, 0xff, 0x99, 0xb0, 0x71, 0x8c, 0x2c, 0xc6, 0xbc,
	0xb9, 0x54, 0x7a, 0x3e, 0x19, 0x64, 0xc3, 0x59, 0x7b, 0xcf, 0x19, 0xdc, 0xbe, 0x51, 0xc5, 0xd0,
	0x4b, 0x7c, 0x8f, 0x17, 0x66, 0xee, 0x4d, 0x90, 0x07, 0xdd, 0xdf, 0x78, 0xf6, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0xca, 0x07, 0xe5, 0xae, 0x03, 0x00, 0x00,
}
