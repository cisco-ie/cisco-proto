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
// source: tcp_nsr_pcb_brief_bag.proto

package cisco_ios_xr_ip_tcp_oper_tcp_nsr_nodes_node_session_brief_sessions_brief_session

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

type TcpNsrPcbBriefBag_KEYS struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Id_1                 string   `protobuf:"bytes,2,opt,name=id_1,json=id1,proto3" json:"id_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpNsrPcbBriefBag_KEYS) Reset()         { *m = TcpNsrPcbBriefBag_KEYS{} }
func (m *TcpNsrPcbBriefBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TcpNsrPcbBriefBag_KEYS) ProtoMessage()    {}
func (*TcpNsrPcbBriefBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d2614875c3448c3, []int{0}
}

func (m *TcpNsrPcbBriefBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrPcbBriefBag_KEYS.Unmarshal(m, b)
}
func (m *TcpNsrPcbBriefBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrPcbBriefBag_KEYS.Marshal(b, m, deterministic)
}
func (m *TcpNsrPcbBriefBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrPcbBriefBag_KEYS.Merge(m, src)
}
func (m *TcpNsrPcbBriefBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TcpNsrPcbBriefBag_KEYS.Size(m)
}
func (m *TcpNsrPcbBriefBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrPcbBriefBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrPcbBriefBag_KEYS proto.InternalMessageInfo

func (m *TcpNsrPcbBriefBag_KEYS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TcpNsrPcbBriefBag_KEYS) GetId_1() string {
	if m != nil {
		return m.Id_1
	}
	return ""
}

type TcpNsrPcbBriefBag struct {
	AddressFamily                string   `protobuf:"bytes,50,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	Pcb                          uint64   `protobuf:"varint,51,opt,name=pcb,proto3" json:"pcb,omitempty"`
	Sscb                         uint64   `protobuf:"varint,52,opt,name=sscb,proto3" json:"sscb,omitempty"`
	LocalAddress                 []string `protobuf:"bytes,53,rep,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	ForeignAddress               []string `protobuf:"bytes,54,rep,name=foreign_address,json=foreignAddress,proto3" json:"foreign_address,omitempty"`
	LocalPort                    uint32   `protobuf:"varint,55,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	ForeignPort                  uint32   `protobuf:"varint,56,opt,name=foreign_port,json=foreignPort,proto3" json:"foreign_port,omitempty"`
	VrfId                        uint32   `protobuf:"varint,57,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	IsAdminConfiguredUp          bool     `protobuf:"varint,58,opt,name=is_admin_configured_up,json=isAdminConfiguredUp,proto3" json:"is_admin_configured_up,omitempty"`
	IsUsOperationalUp            string   `protobuf:"bytes,59,opt,name=is_us_operational_up,json=isUsOperationalUp,proto3" json:"is_us_operational_up,omitempty"`
	IsDsOperationalUp            string   `protobuf:"bytes,60,opt,name=is_ds_operational_up,json=isDsOperationalUp,proto3" json:"is_ds_operational_up,omitempty"`
	IsOnlyReceivePathReplication bool     `protobuf:"varint,61,opt,name=is_only_receive_path_replication,json=isOnlyReceivePathReplication,proto3" json:"is_only_receive_path_replication,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *TcpNsrPcbBriefBag) Reset()         { *m = TcpNsrPcbBriefBag{} }
func (m *TcpNsrPcbBriefBag) String() string { return proto.CompactTextString(m) }
func (*TcpNsrPcbBriefBag) ProtoMessage()    {}
func (*TcpNsrPcbBriefBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d2614875c3448c3, []int{1}
}

func (m *TcpNsrPcbBriefBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrPcbBriefBag.Unmarshal(m, b)
}
func (m *TcpNsrPcbBriefBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrPcbBriefBag.Marshal(b, m, deterministic)
}
func (m *TcpNsrPcbBriefBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrPcbBriefBag.Merge(m, src)
}
func (m *TcpNsrPcbBriefBag) XXX_Size() int {
	return xxx_messageInfo_TcpNsrPcbBriefBag.Size(m)
}
func (m *TcpNsrPcbBriefBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrPcbBriefBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrPcbBriefBag proto.InternalMessageInfo

func (m *TcpNsrPcbBriefBag) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *TcpNsrPcbBriefBag) GetPcb() uint64 {
	if m != nil {
		return m.Pcb
	}
	return 0
}

func (m *TcpNsrPcbBriefBag) GetSscb() uint64 {
	if m != nil {
		return m.Sscb
	}
	return 0
}

func (m *TcpNsrPcbBriefBag) GetLocalAddress() []string {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *TcpNsrPcbBriefBag) GetForeignAddress() []string {
	if m != nil {
		return m.ForeignAddress
	}
	return nil
}

func (m *TcpNsrPcbBriefBag) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *TcpNsrPcbBriefBag) GetForeignPort() uint32 {
	if m != nil {
		return m.ForeignPort
	}
	return 0
}

func (m *TcpNsrPcbBriefBag) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *TcpNsrPcbBriefBag) GetIsAdminConfiguredUp() bool {
	if m != nil {
		return m.IsAdminConfiguredUp
	}
	return false
}

func (m *TcpNsrPcbBriefBag) GetIsUsOperationalUp() string {
	if m != nil {
		return m.IsUsOperationalUp
	}
	return ""
}

func (m *TcpNsrPcbBriefBag) GetIsDsOperationalUp() string {
	if m != nil {
		return m.IsDsOperationalUp
	}
	return ""
}

func (m *TcpNsrPcbBriefBag) GetIsOnlyReceivePathReplication() bool {
	if m != nil {
		return m.IsOnlyReceivePathReplication
	}
	return false
}

func init() {
	proto.RegisterType((*TcpNsrPcbBriefBag_KEYS)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.session.brief_sessions.brief_session.tcp_nsr_pcb_brief_bag_KEYS")
	proto.RegisterType((*TcpNsrPcbBriefBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.session.brief_sessions.brief_session.tcp_nsr_pcb_brief_bag")
}

func init() { proto.RegisterFile("tcp_nsr_pcb_brief_bag.proto", fileDescriptor_2d2614875c3448c3) }

var fileDescriptor_2d2614875c3448c3 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x86, 0x49, 0x5b, 0x17, 0x3b, 0x6e, 0xab, 0x3b, 0xba, 0x32, 0xf8, 0x01, 0x71, 0x45, 0xcc,
	0x55, 0x64, 0xad, 0xdf, 0x1f, 0xc8, 0xa2, 0x2e, 0x88, 0x17, 0x5b, 0x22, 0x7b, 0xe1, 0xd5, 0x21,
	0x99, 0x99, 0x74, 0x0f, 0x64, 0x67, 0x86, 0x39, 0x69, 0xb1, 0x7f, 0xd9, 0x5f, 0x21, 0x3d, 0x4d,
	0xbb, 0xa8, 0xbd, 0x09, 0x73, 0x9e, 0xf3, 0xbc, 0x6f, 0x42, 0x12, 0x71, 0xbf, 0xd5, 0x01, 0x1c,
	0x45, 0x08, 0xba, 0x82, 0x2a, 0xa2, 0xad, 0xa1, 0x2a, 0x67, 0x79, 0x88, 0xbe, 0xf5, 0x72, 0xaa,
	0x91, 0xb4, 0x07, 0xf4, 0x04, 0xbf, 0x22, 0x60, 0x80, 0x95, 0xec, 0x83, 0x8d, 0x79, 0x97, 0xca,
	0x9d, 0x37, 0x96, 0xf8, 0x9a, 0x93, 0x25, 0x42, 0xef, 0xf2, 0x75, 0x49, 0x37, 0xd1, 0xdf, 0xe3,
	0xd1, 0x27, 0x71, 0x6f, 0xe7, 0x0d, 0xe1, 0xfb, 0xd7, 0x9f, 0x3f, 0xe4, 0x58, 0xf4, 0xd0, 0xa8,
	0x24, 0x4d, 0xb2, 0x61, 0xd1, 0x43, 0x23, 0x0f, 0xc4, 0x00, 0x0d, 0x1c, 0xab, 0x1e, 0x93, 0x3e,
	0x9a, 0xe3, 0xa3, 0xdf, 0x7d, 0x71, 0xb8, 0xb3, 0x41, 0x3e, 0x11, 0xe3, 0xd2, 0x98, 0x68, 0x89,
	0xa0, 0x2e, 0x2f, 0xb1, 0x59, 0xaa, 0xe7, 0x1c, 0x1b, 0x75, 0xf4, 0x94, 0xa1, 0xbc, 0x25, 0xfa,
	0x41, 0x57, 0x6a, 0x92, 0x26, 0xd9, 0xa0, 0x58, 0x1d, 0xa5, 0x14, 0x03, 0x22, 0x5d, 0xa9, 0x17,
	0x8c, 0xf8, 0x2c, 0x1f, 0x8b, 0x51, 0xe3, 0x75, 0xd9, 0x40, 0x17, 0x56, 0x2f, 0xd3, 0x7e, 0x36,
	0x2c, 0xf6, 0x19, 0x9e, 0xac, 0x99, 0x7c, 0x2a, 0x6e, 0xd6, 0x3e, 0x5a, 0x9c, 0xb9, 0xad, 0xf6,
	0x8a, 0xb5, 0x71, 0x87, 0x37, 0xe2, 0x43, 0x21, 0xd6, 0x6d, 0xc1, 0xc7, 0x56, 0xbd, 0x4e, 0x93,
	0x6c, 0x54, 0x0c, 0x99, 0x4c, 0x7d, 0x6c, 0xe5, 0x23, 0xb1, 0xbf, 0xe9, 0x61, 0xe1, 0x0d, 0x0b,
	0x37, 0x3a, 0xc6, 0xca, 0xa1, 0xd8, 0x5b, 0xc4, 0x1a, 0xd0, 0xa8, 0xb7, 0xbc, 0xbc, 0xb6, 0x88,
	0xf5, 0x37, 0x23, 0x27, 0xe2, 0x2e, 0x12, 0x94, 0xe6, 0x12, 0x1d, 0x68, 0xef, 0x6a, 0x9c, 0xcd,
	0xa3, 0x35, 0x30, 0x0f, 0xea, 0x5d, 0x9a, 0x64, 0xd7, 0x8b, 0xdb, 0x48, 0x27, 0xab, 0xe5, 0xe7,
	0xed, 0xee, 0x3c, 0xc8, 0x67, 0xe2, 0x0e, 0x12, 0xcc, 0x89, 0xbf, 0x64, 0xd9, 0xa2, 0x77, 0x65,
	0xb3, 0x8a, 0xbc, 0xe7, 0xd7, 0x75, 0x80, 0x74, 0x4e, 0x67, 0x57, 0x9b, 0x6d, 0xc0, 0xfc, 0x17,
	0xf8, 0xb0, 0x09, 0x7c, 0xf9, 0x27, 0x70, 0x2a, 0x52, 0x24, 0xf0, 0xae, 0x59, 0x42, 0xb4, 0xda,
	0xe2, 0xc2, 0x42, 0x28, 0xdb, 0x0b, 0x88, 0x36, 0x34, 0xa8, 0xd9, 0x53, 0x1f, 0xf9, 0x01, 0x1f,
	0x20, 0x9d, 0xb9, 0x66, 0x59, 0xac, 0xad, 0x69, 0xd9, 0x5e, 0x14, 0x57, 0x4e, 0xb5, 0xc7, 0xbf,
	0xe1, 0xe4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xa8, 0x21, 0xf0, 0xa5, 0x02, 0x00, 0x00,
}
