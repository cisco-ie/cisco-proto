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
// source: ldp_igp_sync_delay_restart_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_active_default_vrf_afs_af_igp_sync_delay_restart

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

type LdpIgpSyncDelayRestartInfo_KEYS struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpIgpSyncDelayRestartInfo_KEYS) Reset()         { *m = LdpIgpSyncDelayRestartInfo_KEYS{} }
func (m *LdpIgpSyncDelayRestartInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpIgpSyncDelayRestartInfo_KEYS) ProtoMessage()    {}
func (*LdpIgpSyncDelayRestartInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_01b4a17085537294, []int{0}
}

func (m *LdpIgpSyncDelayRestartInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIgpSyncDelayRestartInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpIgpSyncDelayRestartInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIgpSyncDelayRestartInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpIgpSyncDelayRestartInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIgpSyncDelayRestartInfo_KEYS.Merge(m, src)
}
func (m *LdpIgpSyncDelayRestartInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpIgpSyncDelayRestartInfo_KEYS.Size(m)
}
func (m *LdpIgpSyncDelayRestartInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIgpSyncDelayRestartInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIgpSyncDelayRestartInfo_KEYS proto.InternalMessageInfo

func (m *LdpIgpSyncDelayRestartInfo_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

type LdpIgpSyncDelayRestartInfo struct {
	Configured           bool     `protobuf:"varint,50,opt,name=configured,proto3" json:"configured,omitempty"`
	DelaySecs            uint32   `protobuf:"varint,51,opt,name=delay_secs,json=delaySecs,proto3" json:"delay_secs,omitempty"`
	TimerRunning         bool     `protobuf:"varint,52,opt,name=timer_running,json=timerRunning,proto3" json:"timer_running,omitempty"`
	RemainingSecs        uint32   `protobuf:"varint,53,opt,name=remaining_secs,json=remainingSecs,proto3" json:"remaining_secs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpIgpSyncDelayRestartInfo) Reset()         { *m = LdpIgpSyncDelayRestartInfo{} }
func (m *LdpIgpSyncDelayRestartInfo) String() string { return proto.CompactTextString(m) }
func (*LdpIgpSyncDelayRestartInfo) ProtoMessage()    {}
func (*LdpIgpSyncDelayRestartInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_01b4a17085537294, []int{1}
}

func (m *LdpIgpSyncDelayRestartInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIgpSyncDelayRestartInfo.Unmarshal(m, b)
}
func (m *LdpIgpSyncDelayRestartInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIgpSyncDelayRestartInfo.Marshal(b, m, deterministic)
}
func (m *LdpIgpSyncDelayRestartInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIgpSyncDelayRestartInfo.Merge(m, src)
}
func (m *LdpIgpSyncDelayRestartInfo) XXX_Size() int {
	return xxx_messageInfo_LdpIgpSyncDelayRestartInfo.Size(m)
}
func (m *LdpIgpSyncDelayRestartInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIgpSyncDelayRestartInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIgpSyncDelayRestartInfo proto.InternalMessageInfo

func (m *LdpIgpSyncDelayRestartInfo) GetConfigured() bool {
	if m != nil {
		return m.Configured
	}
	return false
}

func (m *LdpIgpSyncDelayRestartInfo) GetDelaySecs() uint32 {
	if m != nil {
		return m.DelaySecs
	}
	return 0
}

func (m *LdpIgpSyncDelayRestartInfo) GetTimerRunning() bool {
	if m != nil {
		return m.TimerRunning
	}
	return false
}

func (m *LdpIgpSyncDelayRestartInfo) GetRemainingSecs() uint32 {
	if m != nil {
		return m.RemainingSecs
	}
	return 0
}

func init() {
	proto.RegisterType((*LdpIgpSyncDelayRestartInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.igp.sync_delay_restart.ldp_igp_sync_delay_restart_info_KEYS")
	proto.RegisterType((*LdpIgpSyncDelayRestartInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.igp.sync_delay_restart.ldp_igp_sync_delay_restart_info")
}

func init() {
	proto.RegisterFile("ldp_igp_sync_delay_restart_info.proto", fileDescriptor_01b4a17085537294)
}

var fileDescriptor_01b4a17085537294 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xa5, 0x50, 0x8b, 0x32, 0x78, 0x21, 0x0b, 0x10, 0x15, 0x2a, 0x65, 0xf2, 0x40,
	0x61, 0x66, 0x62, 0x42, 0x62, 0x48, 0x27, 0xc4, 0x70, 0xba, 0x3a, 0xe7, 0xc8, 0x92, 0x63, 0x5b,
	0x67, 0xa7, 0xa2, 0xff, 0x89, 0x1f, 0x89, 0x92, 0x08, 0x84, 0xc4, 0xd0, 0xf1, 0x7d, 0xd2, 0xfb,
	0xee, 0xf4, 0xc4, 0xc6, 0xb5, 0x11, 0x6c, 0x17, 0x21, 0x1d, 0xbd, 0x86, 0x96, 0x1c, 0x1e, 0x81,
	0x29, 0x65, 0xe4, 0x0c, 0xd6, 0x9b, 0xa0, 0x22, 0x87, 0x1c, 0xe4, 0x87, 0xb6, 0x49, 0x07, 0xb0,
	0x21, 0xc1, 0x27, 0x43, 0x1f, 0x5d, 0x82, 0xb1, 0x18, 0x22, 0xb1, 0xfa, 0x49, 0xaa, 0x73, 0x61,
	0x8f, 0x4e, 0xa1, 0xce, 0xf6, 0x40, 0xaa, 0x25, 0x83, 0x83, 0xcb, 0x70, 0x60, 0xa3, 0xd0, 0x24,
	0x85, 0x46, 0xd9, 0x2e, 0xaa, 0xff, 0x77, 0xd6, 0xcf, 0xe2, 0xfe, 0xc4, 0x17, 0xf0, 0xfa, 0xf2,
	0xbe, 0x93, 0x57, 0xe2, 0x0c, 0x0d, 0x78, 0xec, 0xa9, 0x2c, 0xaa, 0xa2, 0x5e, 0x36, 0x0b, 0x34,
	0x6f, 0xd8, 0xd3, 0xfa, 0xab, 0x10, 0xb7, 0x27, 0x0c, 0xf2, 0x46, 0x08, 0x1d, 0xbc, 0xb1, 0xdd,
	0xc0, 0xd4, 0x96, 0x0f, 0x55, 0x51, 0x9f, 0x37, 0x7f, 0x88, 0xbc, 0x16, 0x62, 0x6e, 0x25, 0xd2,
	0xa9, 0xdc, 0x56, 0x45, 0xbd, 0x6a, 0x96, 0x13, 0xd9, 0x91, 0x4e, 0xf2, 0x4e, 0xac, 0xb2, 0xed,
	0x89, 0x81, 0x07, 0xef, 0xad, 0xef, 0xca, 0xc7, 0xc9, 0x70, 0x31, 0xc1, 0x66, 0x66, 0x72, 0x23,
	0x2e, 0x99, 0x7a, 0xb4, 0x63, 0x98, 0x3d, 0x4f, 0x93, 0x67, 0xf5, 0x4b, 0x47, 0xd7, 0x7e, 0x31,
	0x6d, 0xba, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x48, 0x13, 0x04, 0xce, 0x7c, 0x01, 0x00, 0x00,
}