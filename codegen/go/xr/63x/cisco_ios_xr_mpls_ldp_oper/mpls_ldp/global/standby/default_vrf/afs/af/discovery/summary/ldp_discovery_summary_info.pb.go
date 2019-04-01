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
// source: ldp_discovery_summary_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_default_vrf_afs_af_discovery_summary

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

type LdpDiscoverySummaryInfo_KEYS struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpDiscoverySummaryInfo_KEYS) Reset()         { *m = LdpDiscoverySummaryInfo_KEYS{} }
func (m *LdpDiscoverySummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpDiscoverySummaryInfo_KEYS) ProtoMessage()    {}
func (*LdpDiscoverySummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad08bb298a90493, []int{0}
}

func (m *LdpDiscoverySummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpDiscoverySummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpDiscoverySummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpDiscoverySummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpDiscoverySummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpDiscoverySummaryInfo_KEYS.Merge(m, src)
}
func (m *LdpDiscoverySummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpDiscoverySummaryInfo_KEYS.Size(m)
}
func (m *LdpDiscoverySummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpDiscoverySummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpDiscoverySummaryInfo_KEYS proto.InternalMessageInfo

func (m *LdpDiscoverySummaryInfo_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
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
	return fileDescriptor_9ad08bb298a90493, []int{1}
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

type LdpDiscoverySummaryInfo struct {
	Vrf                        *LdpVrfInfo `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	LocalLdpId                 string      `protobuf:"bytes,51,opt,name=local_ldp_id,json=localLdpId,proto3" json:"local_ldp_id,omitempty"`
	NumOfLdpInterfaces         uint32      `protobuf:"varint,52,opt,name=num_of_ldp_interfaces,json=numOfLdpInterfaces,proto3" json:"num_of_ldp_interfaces,omitempty"`
	NumOfActiveLdpInterfaces   uint32      `protobuf:"varint,53,opt,name=num_of_active_ldp_interfaces,json=numOfActiveLdpInterfaces,proto3" json:"num_of_active_ldp_interfaces,omitempty"`
	NumOfLnkDiscXmit           uint32      `protobuf:"varint,54,opt,name=num_of_lnk_disc_xmit,json=numOfLnkDiscXmit,proto3" json:"num_of_lnk_disc_xmit,omitempty"`
	NumOfTgtDiscXmit           uint32      `protobuf:"varint,55,opt,name=num_of_tgt_disc_xmit,json=numOfTgtDiscXmit,proto3" json:"num_of_tgt_disc_xmit,omitempty"`
	NumOfLnkDiscRecv           uint32      `protobuf:"varint,56,opt,name=num_of_lnk_disc_recv,json=numOfLnkDiscRecv,proto3" json:"num_of_lnk_disc_recv,omitempty"`
	NumOfTgtDiscRecv           uint32      `protobuf:"varint,57,opt,name=num_of_tgt_disc_recv,json=numOfTgtDiscRecv,proto3" json:"num_of_tgt_disc_recv,omitempty"`
	NumOfDiscWithBadAddrRecv   uint32      `protobuf:"varint,58,opt,name=num_of_disc_with_bad_addr_recv,json=numOfDiscWithBadAddrRecv,proto3" json:"num_of_disc_with_bad_addr_recv,omitempty"`
	NumOfDiscWithBadHelloPdu   uint32      `protobuf:"varint,59,opt,name=num_of_disc_with_bad_hello_pdu,json=numOfDiscWithBadHelloPdu,proto3" json:"num_of_disc_with_bad_hello_pdu,omitempty"`
	NumOfDiscWithBadXportAddr  uint32      `protobuf:"varint,60,opt,name=num_of_disc_with_bad_xport_addr,json=numOfDiscWithBadXportAddr,proto3" json:"num_of_disc_with_bad_xport_addr,omitempty"`
	NumOfDiscWithSameRouterId  uint32      `protobuf:"varint,61,opt,name=num_of_disc_with_same_router_id,json=numOfDiscWithSameRouterId,proto3" json:"num_of_disc_with_same_router_id,omitempty"`
	NumOfDiscWithWrongRouterId uint32      `protobuf:"varint,62,opt,name=num_of_disc_with_wrong_router_id,json=numOfDiscWithWrongRouterId,proto3" json:"num_of_disc_with_wrong_router_id,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}    `json:"-"`
	XXX_unrecognized           []byte      `json:"-"`
	XXX_sizecache              int32       `json:"-"`
}

func (m *LdpDiscoverySummaryInfo) Reset()         { *m = LdpDiscoverySummaryInfo{} }
func (m *LdpDiscoverySummaryInfo) String() string { return proto.CompactTextString(m) }
func (*LdpDiscoverySummaryInfo) ProtoMessage()    {}
func (*LdpDiscoverySummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad08bb298a90493, []int{2}
}

func (m *LdpDiscoverySummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpDiscoverySummaryInfo.Unmarshal(m, b)
}
func (m *LdpDiscoverySummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpDiscoverySummaryInfo.Marshal(b, m, deterministic)
}
func (m *LdpDiscoverySummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpDiscoverySummaryInfo.Merge(m, src)
}
func (m *LdpDiscoverySummaryInfo) XXX_Size() int {
	return xxx_messageInfo_LdpDiscoverySummaryInfo.Size(m)
}
func (m *LdpDiscoverySummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpDiscoverySummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpDiscoverySummaryInfo proto.InternalMessageInfo

func (m *LdpDiscoverySummaryInfo) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpDiscoverySummaryInfo) GetLocalLdpId() string {
	if m != nil {
		return m.LocalLdpId
	}
	return ""
}

func (m *LdpDiscoverySummaryInfo) GetNumOfLdpInterfaces() uint32 {
	if m != nil {
		return m.NumOfLdpInterfaces
	}
	return 0
}

func (m *LdpDiscoverySummaryInfo) GetNumOfActiveLdpInterfaces() uint32 {
	if m != nil {
		return m.NumOfActiveLdpInterfaces
	}
	return 0
}

func (m *LdpDiscoverySummaryInfo) GetNumOfLnkDiscXmit() uint32 {
	if m != nil {
		return m.NumOfLnkDiscXmit
	}
	return 0
}

func (m *LdpDiscoverySummaryInfo) GetNumOfTgtDiscXmit() uint32 {
	if m != nil {
		return m.NumOfTgtDiscXmit
	}
	return 0
}

func (m *LdpDiscoverySummaryInfo) GetNumOfLnkDiscRecv() uint32 {
	if m != nil {
		return m.NumOfLnkDiscRecv
	}
	return 0
}

func (m *LdpDiscoverySummaryInfo) GetNumOfTgtDiscRecv() uint32 {
	if m != nil {
		return m.NumOfTgtDiscRecv
	}
	return 0
}

func (m *LdpDiscoverySummaryInfo) GetNumOfDiscWithBadAddrRecv() uint32 {
	if m != nil {
		return m.NumOfDiscWithBadAddrRecv
	}
	return 0
}

func (m *LdpDiscoverySummaryInfo) GetNumOfDiscWithBadHelloPdu() uint32 {
	if m != nil {
		return m.NumOfDiscWithBadHelloPdu
	}
	return 0
}

func (m *LdpDiscoverySummaryInfo) GetNumOfDiscWithBadXportAddr() uint32 {
	if m != nil {
		return m.NumOfDiscWithBadXportAddr
	}
	return 0
}

func (m *LdpDiscoverySummaryInfo) GetNumOfDiscWithSameRouterId() uint32 {
	if m != nil {
		return m.NumOfDiscWithSameRouterId
	}
	return 0
}

func (m *LdpDiscoverySummaryInfo) GetNumOfDiscWithWrongRouterId() uint32 {
	if m != nil {
		return m.NumOfDiscWithWrongRouterId
	}
	return 0
}

func init() {
	proto.RegisterType((*LdpDiscoverySummaryInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.discovery.summary.ldp_discovery_summary_info_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.discovery.summary.ldp_vrf_info")
	proto.RegisterType((*LdpDiscoverySummaryInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.discovery.summary.ldp_discovery_summary_info")
}

func init() { proto.RegisterFile("ldp_discovery_summary_info.proto", fileDescriptor_9ad08bb298a90493) }

var fileDescriptor_9ad08bb298a90493 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x5d, 0x8b, 0x13, 0x31,
	0x14, 0x86, 0x69, 0x95, 0x8a, 0x71, 0x15, 0x09, 0x8a, 0x71, 0x11, 0x77, 0xe8, 0x55, 0xaf, 0x02,
	0x76, 0xfd, 0x5c, 0x75, 0x71, 0x97, 0x15, 0x5c, 0x14, 0x95, 0x59, 0xa1, 0xeb, 0x55, 0x48, 0x27,
	0x49, 0x1b, 0x76, 0x32, 0x19, 0x92, 0xcc, 0xb4, 0xfd, 0x6f, 0xfe, 0x38, 0xc9, 0xe9, 0x07, 0xb3,
	0xdb, 0xa9, 0x77, 0xde, 0xb5, 0x93, 0xe7, 0x7d, 0xde, 0x73, 0xc2, 0x30, 0x28, 0xc9, 0x45, 0xc9,
	0x84, 0xf6, 0x99, 0xad, 0xa5, 0x5b, 0x30, 0x5f, 0x19, 0xc3, 0xdd, 0x82, 0xe9, 0x42, 0x59, 0x5a,
	0x3a, 0x1b, 0x2c, 0x1e, 0x65, 0xf1, 0x94, 0x69, 0xeb, 0xd9, 0xdc, 0x31, 0x53, 0xe6, 0x9e, 0xc5,
	0x8c, 0x2d, 0xa5, 0xa3, 0xeb, 0x7f, 0x74, 0x92, 0xdb, 0x31, 0xcf, 0xa9, 0x0f, 0xbc, 0x10, 0xe3,
	0x05, 0x15, 0x52, 0xf1, 0x2a, 0x0f, 0xac, 0x76, 0x8a, 0x72, 0xe5, 0x29, 0x57, 0x74, 0xd3, 0x41,
	0x57, 0x1d, 0xfd, 0x23, 0x74, 0xb0, 0xbb, 0x9c, 0x7d, 0xfd, 0xfc, 0xfb, 0x02, 0x3f, 0x41, 0x77,
	0xb8, 0x62, 0x05, 0x37, 0x92, 0x74, 0x92, 0xce, 0xe0, 0x6e, 0xda, 0xe3, 0xea, 0x3b, 0x37, 0xb2,
	0x3f, 0x44, 0x7b, 0x31, 0x5b, 0x3b, 0x05, 0x34, 0xc6, 0xe8, 0x76, 0x83, 0x82, 0xdf, 0xf8, 0x01,
	0xea, 0x6a, 0x41, 0xba, 0x49, 0x67, 0x70, 0x3f, 0xed, 0x6a, 0xd1, 0xff, 0xd3, 0x43, 0xfb, 0xbb,
	0x0b, 0xf1, 0x0c, 0xdd, 0xaa, 0x9d, 0x22, 0xc3, 0xa4, 0x33, 0xb8, 0x37, 0x94, 0xf4, 0x3f, 0x6d,
	0x4d, 0x9b, 0x63, 0xa7, 0xb1, 0x11, 0x27, 0x68, 0x2f, 0xb7, 0x19, 0xcf, 0xa1, 0x40, 0x0b, 0x72,
	0x08, 0x3b, 0x20, 0x78, 0xf6, 0x4d, 0x94, 0xe7, 0x02, 0xbf, 0x40, 0x8f, 0x8b, 0xca, 0x30, 0xab,
	0x96, 0x48, 0x11, 0xa4, 0x53, 0x3c, 0x93, 0x9e, 0xbc, 0x84, 0xe5, 0x70, 0x51, 0x99, 0x1f, 0x2a,
	0xa2, 0x9b, 0x13, 0x7c, 0x8c, 0x9e, 0xad, 0x22, 0x3c, 0x0b, 0xba, 0x96, 0x37, 0x93, 0xaf, 0x20,
	0x49, 0x20, 0x79, 0x02, 0xc4, 0xf5, 0x3c, 0x45, 0x8f, 0xd6, 0x95, 0xc5, 0x15, 0x5c, 0x19, 0x9b,
	0x1b, 0x1d, 0xc8, 0x6b, 0xc8, 0x3d, 0x5c, 0x36, 0x16, 0x57, 0x67, 0xda, 0x67, 0x97, 0x46, 0x87,
	0x06, 0x1f, 0x26, 0xa1, 0xc1, 0xbf, 0x69, 0xf0, 0xbf, 0x26, 0xa1, 0x85, 0xdf, 0xf8, 0x9d, 0xcc,
	0x6a, 0xf2, 0x76, 0xdb, 0x9f, 0xca, 0xac, 0x6e, 0xf3, 0x03, 0xff, 0x6e, 0xdb, 0x0f, 0xfc, 0x27,
	0xf4, 0x7c, 0xc5, 0x03, 0x3b, 0xd3, 0x61, 0xca, 0xc6, 0x5c, 0x30, 0x2e, 0x84, 0x5b, 0x26, 0x8f,
	0x1a, 0x37, 0x10, 0x63, 0x23, 0x1d, 0xa6, 0xa7, 0x5c, 0x9c, 0x08, 0xe1, 0xfe, 0x69, 0x98, 0xca,
	0x3c, 0xb7, 0xac, 0x14, 0x15, 0x79, 0xdf, 0x6e, 0xf8, 0x12, 0x81, 0x9f, 0xa2, 0xc2, 0xa7, 0xe8,
	0xa0, 0xd5, 0x30, 0x2f, 0xad, 0x0b, 0x30, 0x09, 0xf9, 0x00, 0x8a, 0xa7, 0x37, 0x15, 0x97, 0x91,
	0x88, 0x93, 0xb4, 0x3a, 0x3c, 0x37, 0x92, 0x39, 0x5b, 0x05, 0xe9, 0xe2, 0xfb, 0xf2, 0xb1, 0xc5,
	0x71, 0xc1, 0x8d, 0x4c, 0x81, 0x38, 0x17, 0xf8, 0x0c, 0x25, 0x5b, 0x8e, 0x99, 0xb3, 0xc5, 0xa4,
	0x21, 0x39, 0x06, 0xc9, 0xfe, 0x35, 0xc9, 0x28, 0x32, 0x6b, 0xcb, 0xb8, 0x07, 0x9f, 0x83, 0xc3,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x04, 0x24, 0x04, 0x7f, 0x32, 0x04, 0x00, 0x00,
}