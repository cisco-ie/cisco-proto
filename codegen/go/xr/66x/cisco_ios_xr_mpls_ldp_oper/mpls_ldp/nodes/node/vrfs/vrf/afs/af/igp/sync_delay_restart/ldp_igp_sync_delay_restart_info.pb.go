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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_vrfs_vrf_afs_af_igp_sync_delay_restart

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
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,3,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
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

func (m *LdpIgpSyncDelayRestartInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpIgpSyncDelayRestartInfo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*LdpIgpSyncDelayRestartInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.afs.af.igp.sync_delay_restart.ldp_igp_sync_delay_restart_info_KEYS")
	proto.RegisterType((*LdpIgpSyncDelayRestartInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.afs.af.igp.sync_delay_restart.ldp_igp_sync_delay_restart_info")
}

func init() {
	proto.RegisterFile("ldp_igp_sync_delay_restart_info.proto", fileDescriptor_01b4a17085537294)
}

var fileDescriptor_01b4a17085537294 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xa9, 0xc2, 0xfe, 0x19, 0xac, 0x87, 0x5c, 0xac, 0x88, 0x5a, 0x56, 0x17, 0x7a, 0xca,
	0xc1, 0xd5, 0x47, 0xf0, 0x24, 0x78, 0xe8, 0xe2, 0xc1, 0xd3, 0x10, 0xd3, 0xa4, 0x04, 0xda, 0x24,
	0x4c, 0xba, 0xd5, 0x7d, 0x27, 0x1f, 0x52, 0x92, 0xaa, 0x08, 0x1e, 0xf6, 0x32, 0x30, 0xbf, 0x1f,
	0xf3, 0x7d, 0x87, 0x81, 0x75, 0xd7, 0x78, 0x34, 0xad, 0xc7, 0xb0, 0xb7, 0x12, 0x1b, 0xd5, 0x89,
	0x3d, 0x92, 0x0a, 0x83, 0xa0, 0x01, 0x8d, 0xd5, 0x8e, 0x7b, 0x72, 0x83, 0x63, 0x2f, 0xd2, 0x04,
	0xe9, 0xd0, 0xb8, 0x80, 0x1f, 0x84, 0xbd, 0xef, 0x02, 0xc6, 0x43, 0xe7, 0x15, 0xf1, 0x9f, 0x8d,
	0x5b, 0xd7, 0xa8, 0x90, 0x26, 0x1f, 0x49, 0x87, 0x38, 0xb8, 0xd0, 0x81, 0x0b, 0xcd, 0x4d, 0xeb,
	0xf9, 0xff, 0x86, 0xd5, 0x3b, 0xdc, 0x1e, 0xe8, 0xc7, 0xa7, 0xc7, 0xd7, 0x2d, 0xbb, 0x80, 0x65,
	0x8c, 0x45, 0x2b, 0x7a, 0x55, 0x64, 0x65, 0x56, 0x2d, 0xeb, 0x45, 0x04, 0xcf, 0xa2, 0x57, 0xec,
	0x1c, 0x16, 0x23, 0xe9, 0xc9, 0x1d, 0x25, 0x37, 0x1f, 0x49, 0x27, 0x75, 0x06, 0x73, 0xf1, 0x6d,
	0x8e, 0x93, 0x99, 0x89, 0x24, 0x56, 0x9f, 0x19, 0x5c, 0x1f, 0x68, 0x66, 0x57, 0x00, 0xd2, 0x59,
	0x6d, 0xda, 0x1d, 0xa9, 0xa6, 0xb8, 0x2b, 0xb3, 0x6a, 0x51, 0xff, 0x21, 0xec, 0x12, 0x60, 0xba,
	0x0a, 0x4a, 0x86, 0x62, 0x53, 0x66, 0x55, 0x5e, 0x2f, 0x13, 0xd9, 0x2a, 0x19, 0xd8, 0x0d, 0xe4,
	0x83, 0xe9, 0x15, 0x21, 0xed, 0xac, 0x35, 0xb6, 0x2d, 0xee, 0x53, 0xc2, 0x49, 0x82, 0xf5, 0xc4,
	0xd8, 0x1a, 0x4e, 0x49, 0xf5, 0xc2, 0xc4, 0x65, 0xca, 0x79, 0x48, 0x39, 0xf9, 0x2f, 0x8d, 0x59,
	0x6f, 0xb3, 0xf4, 0x85, 0xcd, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x7c, 0x2d, 0x27, 0xae,
	0x01, 0x00, 0x00,
}