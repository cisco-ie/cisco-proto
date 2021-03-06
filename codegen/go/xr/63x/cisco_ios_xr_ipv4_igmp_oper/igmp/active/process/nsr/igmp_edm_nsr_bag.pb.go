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
// source: igmp_edm_nsr_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_igmp_active_process_nsr

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

type IgmpEdmNsrBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmNsrBag_KEYS) Reset()         { *m = IgmpEdmNsrBag_KEYS{} }
func (m *IgmpEdmNsrBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmNsrBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmNsrBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_893a9427d26c16e0, []int{0}
}

func (m *IgmpEdmNsrBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmNsrBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmNsrBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmNsrBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmNsrBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmNsrBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmNsrBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmNsrBag_KEYS.Size(m)
}
func (m *IgmpEdmNsrBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmNsrBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmNsrBag_KEYS proto.InternalMessageInfo

type IgmpEdmNsrBag struct {
	State                 uint32   `protobuf:"varint,50,opt,name=state,proto3" json:"state,omitempty"`
	PartnerProcConnected  bool     `protobuf:"varint,51,opt,name=partner_proc_connected,json=partnerProcConnected,proto3" json:"partner_proc_connected,omitempty"`
	CollabConvDone        bool     `protobuf:"varint,52,opt,name=collab_conv_done,json=collabConvDone,proto3" json:"collab_conv_done,omitempty"`
	RmfNotificationDone   bool     `protobuf:"varint,53,opt,name=rmf_notification_done,json=rmfNotificationDone,proto3" json:"rmf_notification_done,omitempty"`
	LastProc              uint64   `protobuf:"varint,54,opt,name=last_proc,json=lastProc,proto3" json:"last_proc,omitempty"`
	LastProcConnectionUp  uint64   `protobuf:"varint,55,opt,name=last_proc_connection_up,json=lastProcConnectionUp,proto3" json:"last_proc_connection_up,omitempty"`
	LastProcConnectionDn  uint64   `protobuf:"varint,56,opt,name=last_proc_connection_dn,json=lastProcConnectionDn,proto3" json:"last_proc_connection_dn,omitempty"`
	LastRmfReady          uint64   `protobuf:"varint,57,opt,name=last_rmf_ready,json=lastRmfReady,proto3" json:"last_rmf_ready,omitempty"`
	LastRmfNotReady       uint64   `protobuf:"varint,58,opt,name=last_rmf_not_ready,json=lastRmfNotReady,proto3" json:"last_rmf_not_ready,omitempty"`
	CountProcConnectionUp uint32   `protobuf:"varint,59,opt,name=count_proc_connection_up,json=countProcConnectionUp,proto3" json:"count_proc_connection_up,omitempty"`
	CountProcConnectionDn uint32   `protobuf:"varint,60,opt,name=count_proc_connection_dn,json=countProcConnectionDn,proto3" json:"count_proc_connection_dn,omitempty"`
	CountRmfReady         uint32   `protobuf:"varint,61,opt,name=count_rmf_ready,json=countRmfReady,proto3" json:"count_rmf_ready,omitempty"`
	CountRmfNotReady      uint32   `protobuf:"varint,62,opt,name=count_rmf_not_ready,json=countRmfNotReady,proto3" json:"count_rmf_not_ready,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *IgmpEdmNsrBag) Reset()         { *m = IgmpEdmNsrBag{} }
func (m *IgmpEdmNsrBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmNsrBag) ProtoMessage()    {}
func (*IgmpEdmNsrBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_893a9427d26c16e0, []int{1}
}

func (m *IgmpEdmNsrBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmNsrBag.Unmarshal(m, b)
}
func (m *IgmpEdmNsrBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmNsrBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmNsrBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmNsrBag.Merge(m, src)
}
func (m *IgmpEdmNsrBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmNsrBag.Size(m)
}
func (m *IgmpEdmNsrBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmNsrBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmNsrBag proto.InternalMessageInfo

func (m *IgmpEdmNsrBag) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *IgmpEdmNsrBag) GetPartnerProcConnected() bool {
	if m != nil {
		return m.PartnerProcConnected
	}
	return false
}

func (m *IgmpEdmNsrBag) GetCollabConvDone() bool {
	if m != nil {
		return m.CollabConvDone
	}
	return false
}

func (m *IgmpEdmNsrBag) GetRmfNotificationDone() bool {
	if m != nil {
		return m.RmfNotificationDone
	}
	return false
}

func (m *IgmpEdmNsrBag) GetLastProc() uint64 {
	if m != nil {
		return m.LastProc
	}
	return 0
}

func (m *IgmpEdmNsrBag) GetLastProcConnectionUp() uint64 {
	if m != nil {
		return m.LastProcConnectionUp
	}
	return 0
}

func (m *IgmpEdmNsrBag) GetLastProcConnectionDn() uint64 {
	if m != nil {
		return m.LastProcConnectionDn
	}
	return 0
}

func (m *IgmpEdmNsrBag) GetLastRmfReady() uint64 {
	if m != nil {
		return m.LastRmfReady
	}
	return 0
}

func (m *IgmpEdmNsrBag) GetLastRmfNotReady() uint64 {
	if m != nil {
		return m.LastRmfNotReady
	}
	return 0
}

func (m *IgmpEdmNsrBag) GetCountProcConnectionUp() uint32 {
	if m != nil {
		return m.CountProcConnectionUp
	}
	return 0
}

func (m *IgmpEdmNsrBag) GetCountProcConnectionDn() uint32 {
	if m != nil {
		return m.CountProcConnectionDn
	}
	return 0
}

func (m *IgmpEdmNsrBag) GetCountRmfReady() uint32 {
	if m != nil {
		return m.CountRmfReady
	}
	return 0
}

func (m *IgmpEdmNsrBag) GetCountRmfNotReady() uint32 {
	if m != nil {
		return m.CountRmfNotReady
	}
	return 0
}

func init() {
	proto.RegisterType((*IgmpEdmNsrBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.process.nsr.igmp_edm_nsr_bag_KEYS")
	proto.RegisterType((*IgmpEdmNsrBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.process.nsr.igmp_edm_nsr_bag")
}

func init() { proto.RegisterFile("igmp_edm_nsr_bag.proto", fileDescriptor_893a9427d26c16e0) }

var fileDescriptor_893a9427d26c16e0 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4b, 0xeb, 0x13, 0x31,
	0x14, 0xc5, 0x29, 0xfc, 0x95, 0x1a, 0xec, 0x83, 0xf4, 0x15, 0x70, 0x53, 0x8a, 0xc8, 0x80, 0x38,
	0x8b, 0x3e, 0xac, 0xef, 0x4d, 0xeb, 0x4a, 0x28, 0x32, 0xe2, 0xc2, 0x55, 0x48, 0x33, 0x99, 0x12,
	0xe8, 0xdc, 0x3b, 0x24, 0xe9, 0xa0, 0x9f, 0xd7, 0x2f, 0x22, 0xc9, 0x74, 0xa6, 0x52, 0x9d, 0x5d,
	0x72, 0xce, 0xef, 0x5c, 0xee, 0x21, 0x21, 0x53, 0x7d, 0xca, 0x0b, 0xae, 0xd2, 0x9c, 0x83, 0x35,
	0xfc, 0x28, 0x4e, 0x71, 0x61, 0xd0, 0x21, 0x5d, 0x49, 0x6d, 0x25, 0x72, 0x8d, 0x96, 0xff, 0x34,
	0x5c, 0x17, 0xe5, 0x9a, 0x07, 0x12, 0x0b, 0x65, 0x62, 0x7f, 0x8a, 0x85, 0x74, 0xba, 0x54, 0x1e,
	0x97, 0xca, 0xda, 0x18, 0xac, 0x59, 0xcc, 0xc8, 0xe4, 0x7e, 0x1c, 0xff, 0xf2, 0xf9, 0xc7, 0xb7,
	0xc5, 0xef, 0x07, 0x32, 0xbc, 0x77, 0xe8, 0x98, 0x3c, 0xb2, 0x4e, 0x38, 0xc5, 0x96, 0xf3, 0x4e,
	0xd4, 0x4b, 0xaa, 0x0b, 0x5d, 0x93, 0x69, 0x21, 0x8c, 0x03, 0x65, 0xb8, 0x1f, 0xcd, 0x25, 0x02,
	0x28, 0xe9, 0x54, 0xca, 0x56, 0xf3, 0x4e, 0xd4, 0x4d, 0xc6, 0x57, 0xf7, 0xab, 0x41, 0xb9, 0xab,
	0x3d, 0x1a, 0x91, 0xa1, 0xc4, 0xf3, 0x59, 0x1c, 0x3d, 0x5f, 0xf2, 0x14, 0x41, 0xb1, 0x75, 0xe0,
	0xfb, 0x95, 0xbe, 0x43, 0x28, 0xf7, 0x08, 0x8a, 0x2e, 0xc9, 0xc4, 0xe4, 0x19, 0x07, 0x74, 0x3a,
	0xd3, 0x52, 0x38, 0x8d, 0x50, 0xe1, 0x9b, 0x80, 0x8f, 0x4c, 0x9e, 0x1d, 0xfe, 0xf2, 0x42, 0xe6,
	0x19, 0x79, 0x72, 0x16, 0xd6, 0x85, 0x85, 0xd8, 0xeb, 0x79, 0x27, 0x7a, 0x48, 0xba, 0x5e, 0xf0,
	0x3b, 0xd0, 0x0d, 0x99, 0x35, 0x66, 0xbd, 0xad, 0x1f, 0x7a, 0x29, 0xd8, 0x36, 0xa0, 0xe3, 0x1a,
	0xdd, 0x35, 0xe6, 0xf7, 0xa2, 0x35, 0x96, 0x02, 0x7b, 0xd3, 0x16, 0xdb, 0x03, 0x7d, 0x4e, 0xfa,
	0x21, 0xe6, 0x3b, 0x18, 0x25, 0xd2, 0x5f, 0xec, 0x6d, 0xa0, 0x9f, 0x7a, 0x35, 0xc9, 0xb3, 0xc4,
	0x6b, 0xf4, 0x25, 0xa1, 0x0d, 0x05, 0xe8, 0xae, 0xe4, 0xbb, 0x40, 0x0e, 0xae, 0xe4, 0x01, 0x5d,
	0x05, 0x6f, 0x09, 0x93, 0x78, 0x81, 0xff, 0x36, 0x78, 0x1f, 0x9e, 0x66, 0x12, 0xfc, 0x7f, 0x2a,
	0xb4, 0x06, 0x53, 0x60, 0x1f, 0x5a, 0x83, 0x7b, 0xa0, 0x2f, 0xc8, 0xa0, 0x0a, 0xde, 0x5a, 0x7c,
	0x0c, 0x7c, 0x2f, 0xc8, 0x4d, 0x8d, 0x57, 0x64, 0x74, 0xe3, 0x6e, 0x3d, 0x3e, 0x05, 0x76, 0x58,
	0xb3, 0x75, 0x91, 0xe3, 0xe3, 0xf0, 0x75, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x69, 0x55,
	0x75, 0x3c, 0xd4, 0x02, 0x00, 0x00,
}
