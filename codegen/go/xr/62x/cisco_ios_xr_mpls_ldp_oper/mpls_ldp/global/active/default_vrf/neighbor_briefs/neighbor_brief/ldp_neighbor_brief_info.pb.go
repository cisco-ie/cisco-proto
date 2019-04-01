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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_active_default_vrf_neighbor_briefs_neighbor_brief

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
	LsrId                string   `protobuf:"bytes,1,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,2,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
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
	proto.RegisterType((*LdpNeighborBriefInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.neighbor_briefs.neighbor_brief.ldp_neighbor_brief_info_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.neighbor_briefs.neighbor_brief.ldp_vrf_info")
	proto.RegisterType((*LdpNeighborBriefAfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.neighbor_briefs.neighbor_brief.ldp_neighbor_brief_af_info")
	proto.RegisterType((*LdpNeighborBriefInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.neighbor_briefs.neighbor_brief.ldp_neighbor_brief_info")
}

func init() { proto.RegisterFile("ldp_neighbor_brief_info.proto", fileDescriptor_57a8ad7f03439cf8) }

var fileDescriptor_57a8ad7f03439cf8 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xdb, 0x6a, 0xdc, 0x30,
	0x10, 0xc5, 0xbb, 0x69, 0x48, 0xb4, 0x59, 0x97, 0x8a, 0x86, 0x98, 0x5e, 0xc0, 0x6c, 0x2f, 0xf8,
	0xc9, 0x85, 0x4d, 0x9b, 0xf7, 0x94, 0x5e, 0x58, 0x5a, 0x5a, 0xf0, 0xf6, 0xa5, 0xb4, 0x30, 0x48,
	0x96, 0xb4, 0x55, 0x91, 0x25, 0x33, 0xb2, 0x4d, 0x02, 0xfd, 0x95, 0xfe, 0x45, 0xbf, 0xa2, 0x5f,
	0x55, 0xe4, 0xf5, 0x36, 0x17, 0x92, 0xc7, 0x7d, 0xf3, 0x9c, 0x33, 0xe7, 0xcc, 0x70, 0xc6, 0x22,
	0x8f, 0x8d, 0xa8, 0xc1, 0x4a, 0xbd, 0xfa, 0xc1, 0x1d, 0x02, 0x47, 0x2d, 0x15, 0x68, 0xab, 0x5c,
	0x5e, 0xa3, 0x6b, 0x1c, 0xfd, 0x5e, 0x6a, 0x5f, 0x3a, 0xd0, 0xce, 0xc3, 0x19, 0x42, 0x55, 0x1b,
	0x0f, 0x41, 0xe0, 0x6a, 0x89, 0xf9, 0xa6, 0xca, 0x57, 0xc6, 0x71, 0x66, 0x72, 0x56, 0x36, 0xba,
	0x93, 0xb9, 0x90, 0x8a, 0xb5, 0xa6, 0x81, 0x0e, 0x55, 0x7e, 0xd5, 0xd8, 0x5f, 0xab, 0x67, 0xdf,
	0xc8, 0xa3, 0x5b, 0xc6, 0xc3, 0x87, 0xb7, 0x5f, 0x97, 0xf4, 0x90, 0xec, 0x1a, 0x8f, 0xa0, 0x45,
	0x12, 0xa5, 0x51, 0xb6, 0x5f, 0xdc, 0x31, 0x1e, 0x17, 0x82, 0x3e, 0x25, 0xb1, 0x61, 0x5c, 0x1a,
	0xf0, 0x35, 0x2b, 0x65, 0xa0, 0x47, 0x69, 0x94, 0x4d, 0x8b, 0x83, 0x1e, 0x5d, 0x06, 0x70, 0x21,
	0x66, 0x73, 0x72, 0x10, 0xcc, 0x3b, 0x5c, 0x3b, 0x52, 0x4a, 0x76, 0x2c, 0xab, 0xe4, 0x60, 0xd5,
	0x7f, 0xd3, 0x98, 0x8c, 0xfe, 0xab, 0x47, 0x5a, 0xcc, 0xfe, 0x46, 0xe4, 0xc1, 0x0d, 0x1b, 0xb1,
	0xc1, 0xe2, 0x19, 0x89, 0x99, 0x10, 0x28, 0xbd, 0x07, 0xc5, 0x2a, 0x6d, 0xce, 0x07, 0xb3, 0xe9,
	0x80, 0xbe, 0xeb, 0x41, 0xfa, 0x82, 0xdc, 0xb7, 0x6d, 0x05, 0x4e, 0x81, 0xe5, 0x08, 0x22, 0x24,
	0xd8, 0x49, 0x3c, 0x1f, 0xe6, 0xdc, 0xb3, 0x6d, 0xf5, 0x59, 0x7d, 0xe2, 0xf8, 0x66, 0x43, 0x5c,
	0x13, 0x0c, 0x66, 0xd2, 0x27, 0xe3, 0xab, 0x82, 0xd3, 0x0d, 0x41, 0x9f, 0x90, 0xf8, 0x92, 0xc0,
	0x70, 0x93, 0xec, 0xf4, 0xad, 0x93, 0x4d, 0xeb, 0x47, 0x6e, 0x66, 0x7f, 0xc6, 0xe4, 0xe8, 0x96,
	0x78, 0xe9, 0x2f, 0x32, 0xee, 0x50, 0x25, 0xf3, 0x34, 0xca, 0x26, 0xf3, 0x9f, 0xf9, 0x36, 0xaf,
	0x9c, 0x5f, 0xbe, 0x42, 0x11, 0xc6, 0xd2, 0x13, 0x72, 0xa4, 0x3d, 0xac, 0x90, 0x95, 0x52, 0xb5,
	0x06, 0x50, 0xfa, 0x86, 0x61, 0xc3, 0xb8, 0x91, 0xc9, 0x71, 0x1a, 0x65, 0x7b, 0xc5, 0xa1, 0xf6,
	0xef, 0x07, 0xb6, 0xb8, 0x20, 0xe9, 0x43, 0xb2, 0x6f, 0x3d, 0x82, 0x6f, 0x58, 0x23, 0x93, 0x97,
	0x7d, 0xf4, 0x7b, 0xd6, 0xe3, 0x32, 0xd4, 0xf4, 0x39, 0xb9, 0xdb, 0xd6, 0xd0, 0xe8, 0x4a, 0x82,
	0x97, 0xa5, 0xb3, 0xc2, 0x27, 0xaf, 0xfa, 0x50, 0xa6, 0x6d, 0xfd, 0x45, 0x57, 0x72, 0xb9, 0x06,
	0xe9, 0xef, 0x88, 0xc4, 0x21, 0xb5, 0x90, 0xf4, 0x7a, 0xa9, 0xe4, 0x24, 0x1d, 0x67, 0x93, 0xf9,
	0xd9, 0xf6, 0x63, 0xb8, 0xf9, 0xbf, 0x2a, 0x26, 0x96, 0xe3, 0x6b, 0x3c, 0x55, 0x0b, 0xab, 0x1c,
	0xdf, 0xed, 0x5f, 0xde, 0xf1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x66, 0x7e, 0x67, 0x9a,
	0x03, 0x00, 0x00,
}
