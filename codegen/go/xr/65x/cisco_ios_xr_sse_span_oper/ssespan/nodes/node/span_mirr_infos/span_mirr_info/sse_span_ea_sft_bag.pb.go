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
// source: sse_span_ea_sft_bag.proto

package cisco_ios_xr_sse_span_oper_ssespan_nodes_node_span_mirr_infos_span_mirr_info

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

type SseSpanEaSftBag_KEYS struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	IntfName             string   `protobuf:"bytes,2,opt,name=intf_name,json=intfName,proto3" json:"intf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SseSpanEaSftBag_KEYS) Reset()         { *m = SseSpanEaSftBag_KEYS{} }
func (m *SseSpanEaSftBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*SseSpanEaSftBag_KEYS) ProtoMessage()    {}
func (*SseSpanEaSftBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_97e89aae0549e7b0, []int{0}
}

func (m *SseSpanEaSftBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SseSpanEaSftBag_KEYS.Unmarshal(m, b)
}
func (m *SseSpanEaSftBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SseSpanEaSftBag_KEYS.Marshal(b, m, deterministic)
}
func (m *SseSpanEaSftBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SseSpanEaSftBag_KEYS.Merge(m, src)
}
func (m *SseSpanEaSftBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_SseSpanEaSftBag_KEYS.Size(m)
}
func (m *SseSpanEaSftBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SseSpanEaSftBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SseSpanEaSftBag_KEYS proto.InternalMessageInfo

func (m *SseSpanEaSftBag_KEYS) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *SseSpanEaSftBag_KEYS) GetIntfName() string {
	if m != nil {
		return m.IntfName
	}
	return ""
}

type SseSpanEaSftBag struct {
	SrcIfh               uint32   `protobuf:"varint,50,opt,name=src_ifh,json=srcIfh,proto3" json:"src_ifh,omitempty"`
	IntfNameXr           []uint32 `protobuf:"varint,51,rep,packed,name=intf_name_xr,json=intfNameXr,proto3" json:"intf_name_xr,omitempty"`
	V4AclFlag            uint32   `protobuf:"varint,52,opt,name=v4_acl_flag,json=v4AclFlag,proto3" json:"v4_acl_flag,omitempty"`
	V6AclFlag            uint32   `protobuf:"varint,53,opt,name=v6_acl_flag,json=v6AclFlag,proto3" json:"v6_acl_flag,omitempty"`
	GreAclFlag           uint32   `protobuf:"varint,54,opt,name=gre_acl_flag,json=greAclFlag,proto3" json:"gre_acl_flag,omitempty"`
	V4State              uint32   `protobuf:"varint,55,opt,name=v4state,proto3" json:"v4state,omitempty"`
	V6State              uint32   `protobuf:"varint,56,opt,name=v6state,proto3" json:"v6state,omitempty"`
	GreState             uint32   `protobuf:"varint,57,opt,name=gre_state,json=greState,proto3" json:"gre_state,omitempty"`
	V4Sessid             uint32   `protobuf:"varint,58,opt,name=v4sessid,proto3" json:"v4sessid,omitempty"`
	V6Sessid             uint32   `protobuf:"varint,59,opt,name=v6sessid,proto3" json:"v6sessid,omitempty"`
	GreSessid            uint32   `protobuf:"varint,60,opt,name=gre_sessid,json=greSessid,proto3" json:"gre_sessid,omitempty"`
	V4DstType            uint32   `protobuf:"varint,61,opt,name=v4dst_type,json=v4dstType,proto3" json:"v4dst_type,omitempty"`
	V6DstType            uint32   `protobuf:"varint,62,opt,name=v6dst_type,json=v6dstType,proto3" json:"v6dst_type,omitempty"`
	GredstType           uint32   `protobuf:"varint,63,opt,name=gredst_type,json=gredstType,proto3" json:"gredst_type,omitempty"`
	V4Statsptr           uint32   `protobuf:"varint,64,opt,name=v4statsptr,proto3" json:"v4statsptr,omitempty"`
	V6Statsptr           uint32   `protobuf:"varint,65,opt,name=v6statsptr,proto3" json:"v6statsptr,omitempty"`
	Grev4Statsptr        uint32   `protobuf:"varint,66,opt,name=grev4statsptr,proto3" json:"grev4statsptr,omitempty"`
	Grev6Statsptr        uint32   `protobuf:"varint,67,opt,name=grev6statsptr,proto3" json:"grev6statsptr,omitempty"`
	Mplsv4Stats          uint32   `protobuf:"varint,68,opt,name=mplsv4stats,proto3" json:"mplsv4stats,omitempty"`
	Mplsv6Pkts           uint32   `protobuf:"varint,69,opt,name=mplsv6pkts,proto3" json:"mplsv6pkts,omitempty"`
	NpUmask              uint32   `protobuf:"varint,70,opt,name=np_umask,json=npUmask,proto3" json:"np_umask,omitempty"`
	Uidb                 []uint32 `protobuf:"varint,71,rep,packed,name=uidb,proto3" json:"uidb,omitempty"`
	SftHwData            []uint32 `protobuf:"varint,72,rep,packed,name=sft_hw_data,json=sftHwData,proto3" json:"sft_hw_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SseSpanEaSftBag) Reset()         { *m = SseSpanEaSftBag{} }
func (m *SseSpanEaSftBag) String() string { return proto.CompactTextString(m) }
func (*SseSpanEaSftBag) ProtoMessage()    {}
func (*SseSpanEaSftBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_97e89aae0549e7b0, []int{1}
}

func (m *SseSpanEaSftBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SseSpanEaSftBag.Unmarshal(m, b)
}
func (m *SseSpanEaSftBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SseSpanEaSftBag.Marshal(b, m, deterministic)
}
func (m *SseSpanEaSftBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SseSpanEaSftBag.Merge(m, src)
}
func (m *SseSpanEaSftBag) XXX_Size() int {
	return xxx_messageInfo_SseSpanEaSftBag.Size(m)
}
func (m *SseSpanEaSftBag) XXX_DiscardUnknown() {
	xxx_messageInfo_SseSpanEaSftBag.DiscardUnknown(m)
}

var xxx_messageInfo_SseSpanEaSftBag proto.InternalMessageInfo

func (m *SseSpanEaSftBag) GetSrcIfh() uint32 {
	if m != nil {
		return m.SrcIfh
	}
	return 0
}

func (m *SseSpanEaSftBag) GetIntfNameXr() []uint32 {
	if m != nil {
		return m.IntfNameXr
	}
	return nil
}

func (m *SseSpanEaSftBag) GetV4AclFlag() uint32 {
	if m != nil {
		return m.V4AclFlag
	}
	return 0
}

func (m *SseSpanEaSftBag) GetV6AclFlag() uint32 {
	if m != nil {
		return m.V6AclFlag
	}
	return 0
}

func (m *SseSpanEaSftBag) GetGreAclFlag() uint32 {
	if m != nil {
		return m.GreAclFlag
	}
	return 0
}

func (m *SseSpanEaSftBag) GetV4State() uint32 {
	if m != nil {
		return m.V4State
	}
	return 0
}

func (m *SseSpanEaSftBag) GetV6State() uint32 {
	if m != nil {
		return m.V6State
	}
	return 0
}

func (m *SseSpanEaSftBag) GetGreState() uint32 {
	if m != nil {
		return m.GreState
	}
	return 0
}

func (m *SseSpanEaSftBag) GetV4Sessid() uint32 {
	if m != nil {
		return m.V4Sessid
	}
	return 0
}

func (m *SseSpanEaSftBag) GetV6Sessid() uint32 {
	if m != nil {
		return m.V6Sessid
	}
	return 0
}

func (m *SseSpanEaSftBag) GetGreSessid() uint32 {
	if m != nil {
		return m.GreSessid
	}
	return 0
}

func (m *SseSpanEaSftBag) GetV4DstType() uint32 {
	if m != nil {
		return m.V4DstType
	}
	return 0
}

func (m *SseSpanEaSftBag) GetV6DstType() uint32 {
	if m != nil {
		return m.V6DstType
	}
	return 0
}

func (m *SseSpanEaSftBag) GetGredstType() uint32 {
	if m != nil {
		return m.GredstType
	}
	return 0
}

func (m *SseSpanEaSftBag) GetV4Statsptr() uint32 {
	if m != nil {
		return m.V4Statsptr
	}
	return 0
}

func (m *SseSpanEaSftBag) GetV6Statsptr() uint32 {
	if m != nil {
		return m.V6Statsptr
	}
	return 0
}

func (m *SseSpanEaSftBag) GetGrev4Statsptr() uint32 {
	if m != nil {
		return m.Grev4Statsptr
	}
	return 0
}

func (m *SseSpanEaSftBag) GetGrev6Statsptr() uint32 {
	if m != nil {
		return m.Grev6Statsptr
	}
	return 0
}

func (m *SseSpanEaSftBag) GetMplsv4Stats() uint32 {
	if m != nil {
		return m.Mplsv4Stats
	}
	return 0
}

func (m *SseSpanEaSftBag) GetMplsv6Pkts() uint32 {
	if m != nil {
		return m.Mplsv6Pkts
	}
	return 0
}

func (m *SseSpanEaSftBag) GetNpUmask() uint32 {
	if m != nil {
		return m.NpUmask
	}
	return 0
}

func (m *SseSpanEaSftBag) GetUidb() []uint32 {
	if m != nil {
		return m.Uidb
	}
	return nil
}

func (m *SseSpanEaSftBag) GetSftHwData() []uint32 {
	if m != nil {
		return m.SftHwData
	}
	return nil
}

func init() {
	proto.RegisterType((*SseSpanEaSftBag_KEYS)(nil), "cisco_ios_xr_sse_span_oper.ssespan.nodes.node.span_mirr_infos.span_mirr_info.sse_span_ea_sft_bag_KEYS")
	proto.RegisterType((*SseSpanEaSftBag)(nil), "cisco_ios_xr_sse_span_oper.ssespan.nodes.node.span_mirr_infos.span_mirr_info.sse_span_ea_sft_bag")
}

func init() { proto.RegisterFile("sse_span_ea_sft_bag.proto", fileDescriptor_97e89aae0549e7b0) }

var fileDescriptor_97e89aae0549e7b0 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcb, 0x6f, 0xd4, 0x30,
	0x10, 0xc6, 0xb5, 0x3c, 0xda, 0xdd, 0xd9, 0xee, 0xc5, 0x1c, 0x70, 0x41, 0x94, 0xa8, 0xe2, 0xb0,
	0xa7, 0x3d, 0xd0, 0x10, 0xde, 0x8f, 0x42, 0x5b, 0x8a, 0x8a, 0x38, 0x6c, 0x41, 0x82, 0x93, 0xe5,
	0x4d, 0x9c, 0x6c, 0xd4, 0x3c, 0x2c, 0x8f, 0x9b, 0xb6, 0xff, 0x28, 0x7f, 0x0f, 0xf2, 0x24, 0x4e,
	0x4a, 0xd5, 0x4b, 0xe4, 0x6f, 0x7e, 0x9f, 0xbf, 0xd8, 0x93, 0x09, 0x6c, 0x23, 0x2a, 0x81, 0x5a,
	0x56, 0x42, 0x49, 0x81, 0xa9, 0x15, 0x2b, 0x99, 0x2d, 0xb4, 0xa9, 0x6d, 0xcd, 0xbe, 0xc7, 0x39,
	0xc6, 0xb5, 0xc8, 0x6b, 0x14, 0x97, 0x46, 0xf4, 0xbe, 0x5a, 0x2b, 0xb3, 0x40, 0x54, 0x4e, 0x2c,
	0xaa, 0x3a, 0x51, 0x48, 0xcf, 0x05, 0xc1, 0x32, 0x37, 0x46, 0xe4, 0x55, 0x5a, 0xe3, 0x0d, 0xbd,
	0x7b, 0x02, 0xfc, 0x96, 0x57, 0x89, 0x93, 0xc3, 0x3f, 0xa7, 0x8c, 0xc1, 0x3d, 0x17, 0xc1, 0x47,
	0xc1, 0x68, 0x3e, 0x59, 0xd2, 0x9a, 0x3d, 0x86, 0x49, 0x5e, 0xd9, 0x54, 0x54, 0xb2, 0x54, 0xfc,
	0x0e, 0x81, 0xb1, 0x2b, 0xfc, 0x90, 0xa5, 0xda, 0xfd, 0x7b, 0x1f, 0x1e, 0xdc, 0x92, 0xc6, 0x1e,
	0xc2, 0x26, 0x9a, 0x58, 0xe4, 0xe9, 0x9a, 0x3f, 0x0f, 0x46, 0xf3, 0xd9, 0x72, 0x03, 0x4d, 0xfc,
	0x2d, 0x5d, 0xb3, 0x00, 0xb6, 0xfa, 0x34, 0x71, 0x69, 0xf8, 0x5e, 0x70, 0x77, 0x3e, 0x5b, 0x82,
	0x0f, 0xfc, 0x6d, 0xd8, 0x0e, 0x4c, 0x9b, 0x50, 0xc8, 0xb8, 0x10, 0x69, 0x21, 0x33, 0x1e, 0xd2,
	0xf6, 0x49, 0x13, 0xee, 0xc7, 0xc5, 0x51, 0x21, 0x33, 0xe2, 0xd1, 0xc0, 0x5f, 0x74, 0x3c, 0xf2,
	0x3c, 0x80, 0xad, 0xcc, 0xa8, 0xc1, 0x10, 0x91, 0x01, 0x32, 0xa3, 0xbc, 0x83, 0xc3, 0x66, 0x13,
	0xa2, 0x95, 0x56, 0xf1, 0x97, 0x04, 0xbd, 0x24, 0x12, 0xb5, 0xe4, 0x55, 0x47, 0x5a, 0xe9, 0xba,
	0xe0, 0x52, 0x5b, 0xf6, 0x9a, 0xd8, 0x38, 0x33, 0xea, 0x94, 0xe0, 0x23, 0x18, 0x37, 0x21, 0x2a,
	0xc4, 0x3c, 0xe1, 0x6f, 0x5a, 0xe6, 0x35, 0xb1, 0xa8, 0x63, 0x6f, 0x3b, 0xd6, 0x69, 0xf6, 0x04,
	0x80, 0x42, 0x5b, 0xfa, 0xae, 0xbd, 0x89, 0x4b, 0xed, 0x71, 0x13, 0x26, 0x68, 0x85, 0xbd, 0xd2,
	0x8a, 0xbf, 0xf7, 0x8d, 0x48, 0xd0, 0xfe, 0xbc, 0xd2, 0x8a, 0x70, 0xd4, 0xe3, 0x0f, 0xbe, 0x0f,
	0x1e, 0x3f, 0x85, 0x69, 0x66, 0x54, 0xcf, 0x3f, 0xf6, 0x6d, 0xf0, 0x86, 0x1d, 0x17, 0xef, 0x2e,
	0x84, 0xda, 0x1a, 0xfe, 0xa9, 0xe5, 0x43, 0x85, 0x78, 0xd4, 0xf3, 0xfd, 0x8e, 0xf7, 0x15, 0xf6,
	0x0c, 0x66, 0x99, 0x51, 0xd7, 0x22, 0x3e, 0x93, 0xe5, 0xff, 0xa2, 0x77, 0x0d, 0x41, 0x5f, 0x06,
	0xd7, 0x90, 0x15, 0xc0, 0xb4, 0xd4, 0x05, 0x76, 0xfb, 0xf8, 0x01, 0x79, 0xae, 0x97, 0xdc, 0x69,
	0x48, 0x46, 0xfa, 0xcc, 0x22, 0x3f, 0x6c, 0x4f, 0x33, 0x54, 0xd8, 0x36, 0x8c, 0x2b, 0x2d, 0xce,
	0x4b, 0x89, 0x67, 0xfc, 0xa8, 0xfd, 0x76, 0x95, 0xfe, 0xe5, 0xa4, 0x9b, 0xea, 0xf3, 0x3c, 0x59,
	0xf1, 0xaf, 0x34, 0x6b, 0xb4, 0x76, 0x53, 0xe4, 0x66, 0x75, 0x7d, 0x21, 0x12, 0x69, 0x25, 0x3f,
	0x26, 0x34, 0xc1, 0xd4, 0x1e, 0x5f, 0x1c, 0x48, 0x2b, 0x57, 0x1b, 0xf4, 0xeb, 0xed, 0xfd, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0xae, 0x3f, 0x4e, 0x15, 0x97, 0x03, 0x00, 0x00,
}