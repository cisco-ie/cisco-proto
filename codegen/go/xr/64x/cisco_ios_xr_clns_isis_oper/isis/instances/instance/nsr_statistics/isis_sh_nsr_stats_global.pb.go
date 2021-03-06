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
// source: isis_sh_nsr_stats_global.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics

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

type IsisShNsrStatsGlobal_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShNsrStatsGlobal_KEYS) Reset()         { *m = IsisShNsrStatsGlobal_KEYS{} }
func (m *IsisShNsrStatsGlobal_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShNsrStatsGlobal_KEYS) ProtoMessage()    {}
func (*IsisShNsrStatsGlobal_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1105235fa2f58eb9, []int{0}
}

func (m *IsisShNsrStatsGlobal_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShNsrStatsGlobal_KEYS.Unmarshal(m, b)
}
func (m *IsisShNsrStatsGlobal_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShNsrStatsGlobal_KEYS.Marshal(b, m, deterministic)
}
func (m *IsisShNsrStatsGlobal_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShNsrStatsGlobal_KEYS.Merge(m, src)
}
func (m *IsisShNsrStatsGlobal_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShNsrStatsGlobal_KEYS.Size(m)
}
func (m *IsisShNsrStatsGlobal_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShNsrStatsGlobal_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShNsrStatsGlobal_KEYS proto.InternalMessageInfo

func (m *IsisShNsrStatsGlobal_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type NsrStatsGblType struct {
	NoOfL1Lsp             uint32   `protobuf:"varint,1,opt,name=no_of_l1_lsp,json=noOfL1Lsp,proto3" json:"no_of_l1_lsp,omitempty"`
	NoOfL2Lsp             uint32   `protobuf:"varint,2,opt,name=no_of_l2_lsp,json=noOfL2Lsp,proto3" json:"no_of_l2_lsp,omitempty"`
	NoOfL1Adj             uint32   `protobuf:"varint,3,opt,name=no_of_l1_adj,json=noOfL1Adj,proto3" json:"no_of_l1_adj,omitempty"`
	NoOfL2Adj             uint32   `protobuf:"varint,4,opt,name=no_of_l2_adj,json=noOfL2Adj,proto3" json:"no_of_l2_adj,omitempty"`
	NoOfLiveInterface     uint32   `protobuf:"varint,5,opt,name=no_of_live_interface,json=noOfLiveInterface,proto3" json:"no_of_live_interface,omitempty"`
	NoOfPtpInterface      uint32   `protobuf:"varint,6,opt,name=no_of_ptp_interface,json=noOfPtpInterface,proto3" json:"no_of_ptp_interface,omitempty"`
	NoOfLanInterface      uint32   `protobuf:"varint,7,opt,name=no_of_lan_interface,json=noOfLanInterface,proto3" json:"no_of_lan_interface,omitempty"`
	NoOfLoopbackInterface uint32   `protobuf:"varint,8,opt,name=no_of_loopback_interface,json=noOfLoopbackInterface,proto3" json:"no_of_loopback_interface,omitempty"`
	NoOfTeTunnels         uint32   `protobuf:"varint,9,opt,name=no_of_te_tunnels,json=noOfTeTunnels,proto3" json:"no_of_te_tunnels,omitempty"`
	NoOfTeLinks           uint32   `protobuf:"varint,10,opt,name=no_of_te_links,json=noOfTeLinks,proto3" json:"no_of_te_links,omitempty"`
	NoOfIpv4Routes        uint32   `protobuf:"varint,11,opt,name=no_of_ipv4_routes,json=noOfIpv4Routes,proto3" json:"no_of_ipv4_routes,omitempty"`
	NoOfIpv6Routes        uint32   `protobuf:"varint,12,opt,name=no_of_ipv6_routes,json=noOfIpv6Routes,proto3" json:"no_of_ipv6_routes,omitempty"`
	Seqnum                uint32   `protobuf:"varint,13,opt,name=seqnum,proto3" json:"seqnum,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *NsrStatsGblType) Reset()         { *m = NsrStatsGblType{} }
func (m *NsrStatsGblType) String() string { return proto.CompactTextString(m) }
func (*NsrStatsGblType) ProtoMessage()    {}
func (*NsrStatsGblType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1105235fa2f58eb9, []int{1}
}

func (m *NsrStatsGblType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NsrStatsGblType.Unmarshal(m, b)
}
func (m *NsrStatsGblType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NsrStatsGblType.Marshal(b, m, deterministic)
}
func (m *NsrStatsGblType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NsrStatsGblType.Merge(m, src)
}
func (m *NsrStatsGblType) XXX_Size() int {
	return xxx_messageInfo_NsrStatsGblType.Size(m)
}
func (m *NsrStatsGblType) XXX_DiscardUnknown() {
	xxx_messageInfo_NsrStatsGblType.DiscardUnknown(m)
}

var xxx_messageInfo_NsrStatsGblType proto.InternalMessageInfo

func (m *NsrStatsGblType) GetNoOfL1Lsp() uint32 {
	if m != nil {
		return m.NoOfL1Lsp
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfL2Lsp() uint32 {
	if m != nil {
		return m.NoOfL2Lsp
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfL1Adj() uint32 {
	if m != nil {
		return m.NoOfL1Adj
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfL2Adj() uint32 {
	if m != nil {
		return m.NoOfL2Adj
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfLiveInterface() uint32 {
	if m != nil {
		return m.NoOfLiveInterface
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfPtpInterface() uint32 {
	if m != nil {
		return m.NoOfPtpInterface
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfLanInterface() uint32 {
	if m != nil {
		return m.NoOfLanInterface
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfLoopbackInterface() uint32 {
	if m != nil {
		return m.NoOfLoopbackInterface
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfTeTunnels() uint32 {
	if m != nil {
		return m.NoOfTeTunnels
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfTeLinks() uint32 {
	if m != nil {
		return m.NoOfTeLinks
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfIpv4Routes() uint32 {
	if m != nil {
		return m.NoOfIpv4Routes
	}
	return 0
}

func (m *NsrStatsGblType) GetNoOfIpv6Routes() uint32 {
	if m != nil {
		return m.NoOfIpv6Routes
	}
	return 0
}

func (m *NsrStatsGblType) GetSeqnum() uint32 {
	if m != nil {
		return m.Seqnum
	}
	return 0
}

type NsrStatsGblArrType struct {
	Self                 *NsrStatsGblType   `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	Peer                 []*NsrStatsGblType `protobuf:"bytes,2,rep,name=peer,proto3" json:"peer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NsrStatsGblArrType) Reset()         { *m = NsrStatsGblArrType{} }
func (m *NsrStatsGblArrType) String() string { return proto.CompactTextString(m) }
func (*NsrStatsGblArrType) ProtoMessage()    {}
func (*NsrStatsGblArrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1105235fa2f58eb9, []int{2}
}

func (m *NsrStatsGblArrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NsrStatsGblArrType.Unmarshal(m, b)
}
func (m *NsrStatsGblArrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NsrStatsGblArrType.Marshal(b, m, deterministic)
}
func (m *NsrStatsGblArrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NsrStatsGblArrType.Merge(m, src)
}
func (m *NsrStatsGblArrType) XXX_Size() int {
	return xxx_messageInfo_NsrStatsGblArrType.Size(m)
}
func (m *NsrStatsGblArrType) XXX_DiscardUnknown() {
	xxx_messageInfo_NsrStatsGblArrType.DiscardUnknown(m)
}

var xxx_messageInfo_NsrStatsGblArrType proto.InternalMessageInfo

func (m *NsrStatsGblArrType) GetSelf() *NsrStatsGblType {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *NsrStatsGblArrType) GetPeer() []*NsrStatsGblType {
	if m != nil {
		return m.Peer
	}
	return nil
}

type IsisShNsrStatsGlobal struct {
	IsisVmState          uint32              `protobuf:"varint,50,opt,name=isis_vm_state,json=isisVmState,proto3" json:"isis_vm_state,omitempty"`
	IsisNsrStatsData     *NsrStatsGblArrType `protobuf:"bytes,51,opt,name=isis_nsr_stats_data,json=isisNsrStatsData,proto3" json:"isis_nsr_stats_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IsisShNsrStatsGlobal) Reset()         { *m = IsisShNsrStatsGlobal{} }
func (m *IsisShNsrStatsGlobal) String() string { return proto.CompactTextString(m) }
func (*IsisShNsrStatsGlobal) ProtoMessage()    {}
func (*IsisShNsrStatsGlobal) Descriptor() ([]byte, []int) {
	return fileDescriptor_1105235fa2f58eb9, []int{3}
}

func (m *IsisShNsrStatsGlobal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShNsrStatsGlobal.Unmarshal(m, b)
}
func (m *IsisShNsrStatsGlobal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShNsrStatsGlobal.Marshal(b, m, deterministic)
}
func (m *IsisShNsrStatsGlobal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShNsrStatsGlobal.Merge(m, src)
}
func (m *IsisShNsrStatsGlobal) XXX_Size() int {
	return xxx_messageInfo_IsisShNsrStatsGlobal.Size(m)
}
func (m *IsisShNsrStatsGlobal) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShNsrStatsGlobal.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShNsrStatsGlobal proto.InternalMessageInfo

func (m *IsisShNsrStatsGlobal) GetIsisVmState() uint32 {
	if m != nil {
		return m.IsisVmState
	}
	return 0
}

func (m *IsisShNsrStatsGlobal) GetIsisNsrStatsData() *NsrStatsGblArrType {
	if m != nil {
		return m.IsisNsrStatsData
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShNsrStatsGlobal_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.isis_sh_nsr_stats_global_KEYS")
	proto.RegisterType((*NsrStatsGblType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.nsr_stats_gbl_type")
	proto.RegisterType((*NsrStatsGblArrType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.nsr_stats_gbl_arr_type")
	proto.RegisterType((*IsisShNsrStatsGlobal)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.isis_sh_nsr_stats_global")
}

func init() { proto.RegisterFile("isis_sh_nsr_stats_global.proto", fileDescriptor_1105235fa2f58eb9) }

var fileDescriptor_1105235fa2f58eb9 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0x95, 0x36, 0x04, 0xe2, 0x24, 0x55, 0xba, 0x85, 0xca, 0x17, 0xa0, 0x4a, 0x0f, 0xc0,
	0x81, 0x45, 0x49, 0xab, 0x72, 0x06, 0x95, 0x43, 0x45, 0x54, 0xd0, 0xb6, 0xaa, 0x04, 0x17, 0xcb,
	0xd9, 0x38, 0xe0, 0xd4, 0x6b, 0x1b, 0x8f, 0xb3, 0x82, 0x47, 0xe0, 0x1d, 0x78, 0x22, 0x9e, 0x86,
	0x47, 0x40, 0x9e, 0xdd, 0x65, 0x77, 0x05, 0x3d, 0x01, 0x37, 0xe7, 0x9f, 0xef, 0x9b, 0xb1, 0xe2,
	0x49, 0xc8, 0x03, 0x09, 0x12, 0x18, 0x7c, 0x64, 0x1a, 0x1c, 0x03, 0xcf, 0x3d, 0xb0, 0x0f, 0xca,
	0x2c, 0xb8, 0x8a, 0xad, 0x33, 0xde, 0x44, 0x2f, 0x53, 0x09, 0xa9, 0x61, 0xd2, 0x00, 0xfb, 0xec,
	0x58, 0xaa, 0x34, 0x30, 0x34, 0x8c, 0x15, 0x2e, 0x0e, 0xa7, 0x58, 0x6a, 0xf0, 0x5c, 0xa7, 0xa2,
	0x3e, 0xc5, 0x55, 0x2f, 0x09, 0x5e, 0xa6, 0x30, 0x39, 0x25, 0xf7, 0x6f, 0x9a, 0xc2, 0x5e, 0xbf,
	0x7a, 0x77, 0x11, 0x1d, 0x92, 0x51, 0xe5, 0x32, 0xcd, 0x33, 0x41, 0x3b, 0x07, 0x9d, 0xc7, 0xfd,
	0x64, 0x58, 0x85, 0xe7, 0x3c, 0x13, 0x93, 0x6f, 0x5d, 0x12, 0x35, 0xf4, 0x85, 0x62, 0xfe, 0x8b,
	0x15, 0xd1, 0x43, 0x32, 0xd4, 0x86, 0x99, 0x15, 0x53, 0x53, 0xa6, 0xc0, 0xa2, 0x3a, 0x4a, 0xfa,
	0xda, 0xbc, 0x59, 0xcd, 0xa7, 0x73, 0xb0, 0x0d, 0x60, 0x86, 0xc0, 0x56, 0x03, 0x98, 0xb5, 0x81,
	0x29, 0xe3, 0xcb, 0x35, 0xdd, 0x6e, 0x76, 0x78, 0xb1, 0x5c, 0xb7, 0x3a, 0x04, 0xa0, 0xdb, 0xec,
	0x10, 0x80, 0x67, 0xe4, 0x6e, 0x09, 0xc8, 0x5c, 0x30, 0xa9, 0xbd, 0x70, 0x2b, 0x9e, 0x0a, 0x7a,
	0x0b, 0xc1, 0x5d, 0x04, 0x65, 0x2e, 0xce, 0xaa, 0x42, 0xf4, 0x94, 0xec, 0x15, 0x82, 0xf5, 0xb6,
	0xc1, 0xf7, 0x90, 0x1f, 0x07, 0xfe, 0xad, 0xb7, 0x7f, 0xc0, 0x15, 0xd7, 0x0d, 0xfc, 0x76, 0x8d,
	0xcf, 0xb9, 0xae, 0xf1, 0xe7, 0x84, 0x96, 0xb8, 0x31, 0x76, 0xc1, 0xd3, 0xeb, 0x86, 0x73, 0x07,
	0x9d, 0x7b, 0xe8, 0x94, 0xd5, 0x5a, 0x7c, 0x44, 0xc6, 0x85, 0xe8, 0x05, 0xf3, 0x1b, 0xad, 0x85,
	0x02, 0xda, 0x47, 0x61, 0x14, 0x84, 0x4b, 0x71, 0x59, 0x84, 0xd1, 0x21, 0xd9, 0xf9, 0x05, 0x2a,
	0xa9, 0xaf, 0x81, 0x12, 0xc4, 0x06, 0x05, 0x36, 0x0f, 0x51, 0xf4, 0x84, 0xec, 0x16, 0x90, 0xb4,
	0xf9, 0x31, 0x73, 0x66, 0xe3, 0x05, 0xd0, 0x01, 0x72, 0x3b, 0x81, 0x3b, 0xb3, 0xf9, 0x71, 0x82,
	0x69, 0x0b, 0x3d, 0xa9, 0xd0, 0x61, 0x0b, 0x3d, 0x29, 0xd1, 0x7d, 0xd2, 0x03, 0xf1, 0x49, 0x6f,
	0x32, 0x3a, 0xc2, 0x7a, 0xf9, 0x69, 0xf2, 0xa3, 0x43, 0xf6, 0xdb, 0xeb, 0xc1, 0x9d, 0x2b, 0x56,
	0x64, 0x4d, 0xba, 0x20, 0xd4, 0x0a, 0x57, 0x63, 0x30, 0xbb, 0x8a, 0xff, 0x7e, 0xa5, 0xe3, 0xdf,
	0x17, 0x31, 0xc1, 0x19, 0x61, 0x96, 0x15, 0xc2, 0xd1, 0xad, 0x83, 0xed, 0xff, 0x39, 0x2b, 0xcc,
	0x98, 0x7c, 0xef, 0x10, 0x7a, 0xd3, 0x0f, 0x2b, 0x9a, 0x90, 0x11, 0xd6, 0xf2, 0x0c, 0x73, 0x41,
	0x67, 0xc5, 0x0b, 0x85, 0xf0, 0x2a, 0xbb, 0x08, 0x51, 0xf4, 0xb5, 0x43, 0xf6, 0x10, 0xaa, 0xed,
	0x25, 0xf7, 0x9c, 0x1e, 0xe1, 0x17, 0xf5, 0xfe, 0xdf, 0x5f, 0xbe, 0x7a, 0x92, 0x64, 0x1c, 0xf4,
	0x73, 0x70, 0xe1, 0x1e, 0x70, 0xca, 0x3d, 0x5f, 0xf4, 0xf0, 0xff, 0xe6, 0xe8, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe5, 0x29, 0x9a, 0xfe, 0x91, 0x04, 0x00, 0x00,
}
