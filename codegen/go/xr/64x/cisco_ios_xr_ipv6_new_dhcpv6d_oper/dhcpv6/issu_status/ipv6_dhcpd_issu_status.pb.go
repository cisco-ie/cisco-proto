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
// source: ipv6_dhcpd_issu_status.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_issu_status

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

type Ipv6DhcpdIssuStatus_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6DhcpdIssuStatus_KEYS) Reset()         { *m = Ipv6DhcpdIssuStatus_KEYS{} }
func (m *Ipv6DhcpdIssuStatus_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6DhcpdIssuStatus_KEYS) ProtoMessage()    {}
func (*Ipv6DhcpdIssuStatus_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_109d42765a8c27b2, []int{0}
}

func (m *Ipv6DhcpdIssuStatus_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6DhcpdIssuStatus_KEYS.Unmarshal(m, b)
}
func (m *Ipv6DhcpdIssuStatus_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6DhcpdIssuStatus_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6DhcpdIssuStatus_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6DhcpdIssuStatus_KEYS.Merge(m, src)
}
func (m *Ipv6DhcpdIssuStatus_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6DhcpdIssuStatus_KEYS.Size(m)
}
func (m *Ipv6DhcpdIssuStatus_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6DhcpdIssuStatus_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6DhcpdIssuStatus_KEYS proto.InternalMessageInfo

type Ipv6DhcpdIssuStatus struct {
	ProcessStartTime           uint64   `protobuf:"varint,50,opt,name=process_start_time,json=processStartTime,proto3" json:"process_start_time,omitempty"`
	IssuSyncCompleteTime       uint64   `protobuf:"varint,51,opt,name=issu_sync_complete_time,json=issuSyncCompleteTime,proto3" json:"issu_sync_complete_time,omitempty"`
	IssuSyncStartTime          uint64   `protobuf:"varint,52,opt,name=issu_sync_start_time,json=issuSyncStartTime,proto3" json:"issu_sync_start_time,omitempty"`
	IssuReadyTime              uint64   `protobuf:"varint,53,opt,name=issu_ready_time,json=issuReadyTime,proto3" json:"issu_ready_time,omitempty"`
	BigBangTime                uint64   `protobuf:"varint,54,opt,name=big_bang_time,json=bigBangTime,proto3" json:"big_bang_time,omitempty"`
	PrimaryRoleTime            uint64   `protobuf:"varint,55,opt,name=primary_role_time,json=primaryRoleTime,proto3" json:"primary_role_time,omitempty"`
	IssuReadyIssuMgrConnection bool     `protobuf:"varint,56,opt,name=issu_ready_issu_mgr_connection,json=issuReadyIssuMgrConnection,proto3" json:"issu_ready_issu_mgr_connection,omitempty"`
	Role                       string   `protobuf:"bytes,57,opt,name=role,proto3" json:"role,omitempty"`
	Phase                      string   `protobuf:"bytes,58,opt,name=phase,proto3" json:"phase,omitempty"`
	Version                    string   `protobuf:"bytes,59,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Ipv6DhcpdIssuStatus) Reset()         { *m = Ipv6DhcpdIssuStatus{} }
func (m *Ipv6DhcpdIssuStatus) String() string { return proto.CompactTextString(m) }
func (*Ipv6DhcpdIssuStatus) ProtoMessage()    {}
func (*Ipv6DhcpdIssuStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_109d42765a8c27b2, []int{1}
}

func (m *Ipv6DhcpdIssuStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6DhcpdIssuStatus.Unmarshal(m, b)
}
func (m *Ipv6DhcpdIssuStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6DhcpdIssuStatus.Marshal(b, m, deterministic)
}
func (m *Ipv6DhcpdIssuStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6DhcpdIssuStatus.Merge(m, src)
}
func (m *Ipv6DhcpdIssuStatus) XXX_Size() int {
	return xxx_messageInfo_Ipv6DhcpdIssuStatus.Size(m)
}
func (m *Ipv6DhcpdIssuStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6DhcpdIssuStatus.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6DhcpdIssuStatus proto.InternalMessageInfo

func (m *Ipv6DhcpdIssuStatus) GetProcessStartTime() uint64 {
	if m != nil {
		return m.ProcessStartTime
	}
	return 0
}

func (m *Ipv6DhcpdIssuStatus) GetIssuSyncCompleteTime() uint64 {
	if m != nil {
		return m.IssuSyncCompleteTime
	}
	return 0
}

func (m *Ipv6DhcpdIssuStatus) GetIssuSyncStartTime() uint64 {
	if m != nil {
		return m.IssuSyncStartTime
	}
	return 0
}

func (m *Ipv6DhcpdIssuStatus) GetIssuReadyTime() uint64 {
	if m != nil {
		return m.IssuReadyTime
	}
	return 0
}

func (m *Ipv6DhcpdIssuStatus) GetBigBangTime() uint64 {
	if m != nil {
		return m.BigBangTime
	}
	return 0
}

func (m *Ipv6DhcpdIssuStatus) GetPrimaryRoleTime() uint64 {
	if m != nil {
		return m.PrimaryRoleTime
	}
	return 0
}

func (m *Ipv6DhcpdIssuStatus) GetIssuReadyIssuMgrConnection() bool {
	if m != nil {
		return m.IssuReadyIssuMgrConnection
	}
	return false
}

func (m *Ipv6DhcpdIssuStatus) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Ipv6DhcpdIssuStatus) GetPhase() string {
	if m != nil {
		return m.Phase
	}
	return ""
}

func (m *Ipv6DhcpdIssuStatus) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv6DhcpdIssuStatus_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.issu_status.ipv6_dhcpd_issu_status_KEYS")
	proto.RegisterType((*Ipv6DhcpdIssuStatus)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.issu_status.ipv6_dhcpd_issu_status")
}

func init() { proto.RegisterFile("ipv6_dhcpd_issu_status.proto", fileDescriptor_109d42765a8c27b2) }

var fileDescriptor_109d42765a8c27b2 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x4b, 0x2b, 0x31,
	0x14, 0x86, 0x29, 0xb7, 0xf7, 0x5e, 0x8d, 0x94, 0xda, 0x50, 0x74, 0xf0, 0x8b, 0xd2, 0x85, 0x14,
	0x91, 0x0a, 0xd6, 0xd6, 0xaf, 0x5d, 0x8b, 0x0b, 0x11, 0x37, 0x53, 0x37, 0xae, 0x0e, 0x33, 0x99,
	0x30, 0x0d, 0x74, 0x92, 0x90, 0xa4, 0xd5, 0xf9, 0x53, 0xfe, 0x46, 0xc9, 0xc9, 0xf4, 0x63, 0xd1,
	0xdd, 0x9c, 0xf7, 0x7d, 0x9e, 0x9c, 0xb3, 0x18, 0x72, 0x26, 0xf4, 0x72, 0x04, 0xd9, 0x8c, 0xe9,
	0x0c, 0x84, 0xb5, 0x0b, 0xb0, 0x2e, 0x71, 0x0b, 0xdb, 0xd7, 0x46, 0x39, 0x45, 0x87, 0x4c, 0x58,
	0xa6, 0x40, 0x28, 0x0b, 0xdf, 0x06, 0x10, 0x95, 0xfc, 0x0b, 0xf1, 0xe5, 0x28, 0x03, 0xa5, 0xb9,
	0xe9, 0x87, 0xa1, 0xbf, 0x25, 0x77, 0xcf, 0xc9, 0xe9, 0xee, 0x67, 0xe1, 0xed, 0xe5, 0x73, 0xda,
	0xfd, 0xf9, 0x43, 0x8e, 0x76, 0xf7, 0xf4, 0x9a, 0x50, 0x6d, 0x14, 0xe3, 0xd6, 0xfa, 0xc4, 0x38,
	0x70, 0xa2, 0xe0, 0xd1, 0x6d, 0xa7, 0xd6, 0xab, 0xc7, 0x87, 0x55, 0x33, 0xf5, 0xc5, 0x87, 0x28,
	0x38, 0x1d, 0x92, 0xe3, 0x20, 0x97, 0x92, 0x01, 0x53, 0x85, 0x9e, 0x73, 0xc7, 0x83, 0x32, 0x40,
	0xa5, 0xed, 0xeb, 0x69, 0x29, 0xd9, 0xa4, 0x2a, 0x51, 0xbb, 0x21, 0xed, 0x8d, 0xb6, 0xb5, 0xe6,
	0x0e, 0x9d, 0xd6, 0xca, 0xd9, 0xec, 0xb9, 0x24, 0x4d, 0x14, 0x0c, 0x4f, 0xb2, 0x32, 0xb0, 0x43,
	0x64, 0x1b, 0x3e, 0x8e, 0x7d, 0x8a, 0x5c, 0x97, 0x34, 0x52, 0x91, 0x43, 0x9a, 0xc8, 0x3c, 0x50,
	0x23, 0xa4, 0x0e, 0x52, 0x91, 0x8f, 0x13, 0x99, 0x23, 0x73, 0x45, 0x5a, 0xda, 0x88, 0x22, 0x31,
	0x25, 0x18, 0x35, 0xaf, 0xae, 0xbd, 0x47, 0xae, 0x59, 0x15, 0xb1, 0x9a, 0x87, 0x43, 0xc7, 0xe4,
	0x62, 0x6b, 0x2f, 0x7e, 0x16, 0xb9, 0x01, 0xa6, 0xa4, 0xe4, 0xcc, 0x09, 0x25, 0xa3, 0x87, 0x4e,
	0xad, 0xb7, 0x17, 0x9f, 0xac, 0xcf, 0x78, 0xb5, 0x76, 0xf1, 0x9e, 0x9b, 0xc9, 0x9a, 0xa0, 0x94,
	0xd4, 0xfd, 0x9e, 0xe8, 0xb1, 0x53, 0xeb, 0xed, 0xc7, 0xf8, 0x4d, 0xdb, 0xe4, 0xaf, 0x9e, 0x25,
	0x96, 0x47, 0x4f, 0x18, 0x86, 0x81, 0x46, 0xe4, 0xff, 0x92, 0x1b, 0xeb, 0x9f, 0x7d, 0xc6, 0x7c,
	0x35, 0xa6, 0xff, 0xf0, 0x6f, 0x18, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x25, 0xe2, 0xc1,
	0x2d, 0x02, 0x00, 0x00,
}
