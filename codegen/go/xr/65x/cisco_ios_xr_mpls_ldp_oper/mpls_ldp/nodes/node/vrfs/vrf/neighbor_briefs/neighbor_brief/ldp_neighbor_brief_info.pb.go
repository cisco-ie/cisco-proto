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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_vrfs_vrf_neighbor_briefs_neighbor_brief

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
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	LsrId                string   `protobuf:"bytes,3,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,4,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
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

func (m *LdpNeighborBriefInfo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	proto.RegisterType((*LdpNeighborBriefInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.neighbor_briefs.neighbor_brief.ldp_neighbor_brief_info_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.neighbor_briefs.neighbor_brief.ldp_vrf_info")
	proto.RegisterType((*LdpNeighborBriefAfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.neighbor_briefs.neighbor_brief.ldp_neighbor_brief_af_info")
	proto.RegisterType((*LdpNeighborBriefInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.neighbor_briefs.neighbor_brief.ldp_neighbor_brief_info")
}

func init() { proto.RegisterFile("ldp_neighbor_brief_info.proto", fileDescriptor_57a8ad7f03439cf8) }

var fileDescriptor_57a8ad7f03439cf8 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xdd, 0x6a, 0x1b, 0x3d,
	0x10, 0x65, 0xed, 0x7c, 0xf9, 0x1c, 0x39, 0x76, 0xa9, 0x68, 0xc8, 0xf6, 0x0f, 0x8c, 0xfb, 0x83,
	0xaf, 0xb6, 0xe0, 0xb4, 0xb9, 0x4f, 0xe9, 0x0f, 0xa6, 0x25, 0x85, 0x75, 0x29, 0xf4, 0x6a, 0x90,
	0x2c, 0x29, 0x15, 0x68, 0xa5, 0x65, 0x66, 0xbd, 0x34, 0xcf, 0xd1, 0x8b, 0xbe, 0x42, 0x9f, 0xa5,
	0x4f, 0x55, 0x24, 0xaf, 0xa9, 0x1b, 0x92, 0xcb, 0xdc, 0x88, 0x9d, 0x73, 0x76, 0xce, 0x19, 0x9d,
	0x9d, 0x65, 0x8f, 0x9d, 0xaa, 0xc1, 0x6b, 0x7b, 0xf1, 0x4d, 0x06, 0x04, 0x89, 0x56, 0x1b, 0xb0,
	0xde, 0x84, 0xa2, 0xc6, 0xd0, 0x04, 0xfe, 0x65, 0x65, 0x69, 0x15, 0xc0, 0x06, 0x82, 0xef, 0x08,
	0x55, 0xed, 0x08, 0x62, 0x43, 0xa8, 0x35, 0x16, 0xdb, 0xaa, 0xf0, 0x41, 0x69, 0x4a, 0x67, 0xd1,
	0xa2, 0xa1, 0x78, 0x14, 0xff, 0x4a, 0xd2, 0x95, 0x7a, 0xfa, 0x23, 0x63, 0x8f, 0x6e, 0x70, 0x86,
	0x0f, 0x6f, 0xbf, 0x2e, 0xf9, 0x43, 0x76, 0x10, 0x05, 0xc1, 0x8b, 0x4a, 0xe7, 0xd9, 0x24, 0x9b,
	0x1d, 0x94, 0x83, 0x08, 0x9c, 0x8b, 0x4a, 0xf3, 0xfb, 0x6c, 0xd0, 0xa2, 0xd9, 0x70, 0xbd, 0xc4,
	0xfd, 0xdf, 0xa2, 0x49, 0xd4, 0x11, 0xdb, 0x77, 0x84, 0x60, 0x55, 0xde, 0x4f, 0xc4, 0x7f, 0x8e,
	0x70, 0xa1, 0xf8, 0x53, 0x36, 0x76, 0x42, 0x6a, 0x07, 0x54, 0x8b, 0x95, 0x8e, 0xf4, 0xde, 0x24,
	0x9b, 0x8d, 0xca, 0xc3, 0x84, 0x2e, 0x23, 0xb8, 0x50, 0xd3, 0x39, 0x3b, 0x8c, 0x43, 0x45, 0xed,
	0x38, 0x09, 0xe7, 0x6c, 0x6f, 0xc7, 0x3f, 0x3d, 0xf3, 0x31, 0xeb, 0x59, 0x95, 0x5c, 0x47, 0x65,
	0xcf, 0xaa, 0xe9, 0xef, 0x8c, 0x3d, 0xb8, 0xe6, 0x26, 0xa2, 0x93, 0x78, 0xc6, 0xc6, 0x42, 0x29,
	0xd4, 0x44, 0x60, 0x44, 0x65, 0xdd, 0x65, 0x27, 0x36, 0xea, 0xd0, 0x77, 0x09, 0xe4, 0x2f, 0xd8,
	0x3d, 0xbf, 0xae, 0x20, 0x18, 0xf0, 0x12, 0x41, 0xc5, 0xd0, 0x5b, 0x8d, 0x97, 0x9d, 0xcf, 0x5d,
	0xbf, 0xae, 0x3e, 0x99, 0x73, 0x89, 0x6f, 0xb6, 0xc4, 0x95, 0x86, 0x4e, 0x4c, 0x53, 0xba, 0xf5,
	0x4e, 0xc3, 0xd9, 0x96, 0xe0, 0x4f, 0xd8, 0x78, 0xa7, 0xc1, 0x49, 0xd7, 0x25, 0x30, 0xdc, 0xbe,
	0xfa, 0x51, 0xba, 0xe9, 0xaf, 0x3e, 0x3b, 0xbe, 0xe1, 0xb3, 0xf0, 0x96, 0xf5, 0x5b, 0x34, 0xf9,
	0x7c, 0x92, 0xcd, 0x86, 0x73, 0x55, 0xdc, 0xce, 0x62, 0x14, 0xbb, 0xf9, 0x97, 0xd1, 0x90, 0x9f,
	0xb2, 0x63, 0x4b, 0x70, 0x81, 0x62, 0xa5, 0xcd, 0xda, 0x01, 0x6a, 0x6a, 0x04, 0x36, 0x42, 0x3a,
	0x9d, 0x9f, 0x4c, 0xb2, 0xd9, 0xa0, 0x3c, 0xb2, 0xf4, 0xbe, 0x63, 0xcb, 0xbf, 0x64, 0xda, 0x20,
	0x42, 0xa0, 0x46, 0x34, 0x3a, 0x7f, 0xd9, 0x6d, 0x10, 0xe1, 0x32, 0xd6, 0xfc, 0x39, 0xbb, 0xb3,
	0xae, 0xa1, 0xb1, 0x95, 0x06, 0xd2, 0xab, 0xe0, 0x15, 0xe5, 0xaf, 0x52, 0x1c, 0xa3, 0x75, 0xfd,
	0xd9, 0x56, 0x7a, 0xb9, 0x01, 0xf9, 0xcf, 0x8c, 0x8d, 0x63, 0x5e, 0x31, 0xe3, 0xcd, 0x50, 0xf9,
	0xe9, 0xa4, 0x3f, 0x1b, 0xce, 0xf1, 0x36, 0x03, 0xb8, 0x7e, 0x97, 0xca, 0xa1, 0x97, 0xf8, 0x1a,
	0xcf, 0xcc, 0xc2, 0x9b, 0x20, 0xf7, 0xd3, 0x0f, 0x7a, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x1e,
	0x56, 0xd7, 0x30, 0xc1, 0x03, 0x00, 0x00,
}