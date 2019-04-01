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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_default_vrf_afs_af_igp_sync_delay_restart

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
	proto.RegisterType((*LdpIgpSyncDelayRestartInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.igp.sync_delay_restart.ldp_igp_sync_delay_restart_info_KEYS")
	proto.RegisterType((*LdpIgpSyncDelayRestartInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.igp.sync_delay_restart.ldp_igp_sync_delay_restart_info")
}

func init() {
	proto.RegisterFile("ldp_igp_sync_delay_restart_info.proto", fileDescriptor_01b4a17085537294)
}

var fileDescriptor_01b4a17085537294 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0xf5, 0x30,
	0x14, 0x86, 0xe9, 0x72, 0xbf, 0xef, 0x06, 0xeb, 0x90, 0xc5, 0x2e, 0x6a, 0xb9, 0x7a, 0xa1, 0x53,
	0x06, 0xaf, 0xce, 0x4e, 0x4e, 0x82, 0x43, 0xef, 0x24, 0x08, 0x87, 0xd3, 0xf4, 0xa4, 0x04, 0xd2,
	0x24, 0x24, 0xa9, 0xd8, 0xff, 0xe4, 0x8f, 0x94, 0xa6, 0x28, 0x82, 0xc3, 0x1d, 0xdf, 0x07, 0xde,
	0xe7, 0x1c, 0x5e, 0xb6, 0x37, 0xbd, 0x07, 0x3d, 0x78, 0x88, 0xb3, 0x95, 0xd0, 0x93, 0xc1, 0x19,
	0x02, 0xc5, 0x84, 0x21, 0x81, 0xb6, 0xca, 0x09, 0x1f, 0x5c, 0x72, 0xfc, 0x4d, 0xea, 0x28, 0x1d,
	0x68, 0x17, 0xe1, 0x23, 0xc0, 0xe8, 0x4d, 0x84, 0xa5, 0xe8, 0x3c, 0x05, 0xf1, 0x9d, 0xc4, 0x60,
	0x5c, 0x87, 0x46, 0xc4, 0x84, 0xb6, 0xef, 0x66, 0xd1, 0x93, 0xc2, 0xc9, 0x24, 0x78, 0x0f, 0x4a,
	0xa0, 0x8a, 0x02, 0x95, 0xd0, 0x83, 0x17, 0x7f, 0x0f, 0xed, 0x1e, 0xd9, 0xed, 0x89, 0x37, 0xe0,
	0xf9, 0xe9, 0xf5, 0xc8, 0x2f, 0xd8, 0x3f, 0x54, 0x60, 0x71, 0xa4, 0xaa, 0xa8, 0x8b, 0x66, 0xdb,
	0x6e, 0x50, 0xbd, 0xe0, 0x48, 0xbb, 0xcf, 0x82, 0x5d, 0x9f, 0x30, 0xf0, 0x2b, 0xc6, 0xa4, 0xb3,
	0x4a, 0x0f, 0x53, 0xa0, 0xbe, 0xba, 0xab, 0x8b, 0xe6, 0x7f, 0xfb, 0x8b, 0xf0, 0x4b, 0xc6, 0xd6,
	0x56, 0x24, 0x19, 0xab, 0x43, 0x5d, 0x34, 0x65, 0xbb, 0xcd, 0xe4, 0x48, 0x32, 0xf2, 0x1b, 0x56,
	0x26, 0x3d, 0x52, 0x80, 0x30, 0x59, 0xab, 0xed, 0x50, 0xdd, 0x67, 0xc3, 0x59, 0x86, 0xed, 0xca,
	0xf8, 0x9e, 0x9d, 0x07, 0x1a, 0x51, 0x2f, 0x61, 0xf5, 0x3c, 0x64, 0x4f, 0xf9, 0x43, 0x17, 0x57,
	0xb7, 0xc9, 0xa3, 0x1e, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x78, 0x69, 0xf0, 0x3c, 0x7d, 0x01,
	0x00, 0x00,
}