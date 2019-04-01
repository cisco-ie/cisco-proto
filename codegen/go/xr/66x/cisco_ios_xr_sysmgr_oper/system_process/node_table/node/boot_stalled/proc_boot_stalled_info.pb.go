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
// source: proc_boot_stalled_info.proto

package cisco_ios_xr_sysmgr_oper_system_process_node_table_node_boot_stalled

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

type ProcBootStalledInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcBootStalledInfo_KEYS) Reset()         { *m = ProcBootStalledInfo_KEYS{} }
func (m *ProcBootStalledInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*ProcBootStalledInfo_KEYS) ProtoMessage()    {}
func (*ProcBootStalledInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b27412dfd269c426, []int{0}
}

func (m *ProcBootStalledInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcBootStalledInfo_KEYS.Unmarshal(m, b)
}
func (m *ProcBootStalledInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcBootStalledInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *ProcBootStalledInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcBootStalledInfo_KEYS.Merge(m, src)
}
func (m *ProcBootStalledInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_ProcBootStalledInfo_KEYS.Size(m)
}
func (m *ProcBootStalledInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcBootStalledInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ProcBootStalledInfo_KEYS proto.InternalMessageInfo

func (m *ProcBootStalledInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type BootHoldInfo struct {
	BootHeldByName       string   `protobuf:"bytes,1,opt,name=boot_held_by_name,json=bootHeldByName,proto3" json:"boot_held_by_name,omitempty"`
	InstanceId           uint32   `protobuf:"varint,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Jid                  uint32   `protobuf:"varint,3,opt,name=jid,proto3" json:"jid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BootHoldInfo) Reset()         { *m = BootHoldInfo{} }
func (m *BootHoldInfo) String() string { return proto.CompactTextString(m) }
func (*BootHoldInfo) ProtoMessage()    {}
func (*BootHoldInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b27412dfd269c426, []int{1}
}

func (m *BootHoldInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootHoldInfo.Unmarshal(m, b)
}
func (m *BootHoldInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootHoldInfo.Marshal(b, m, deterministic)
}
func (m *BootHoldInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootHoldInfo.Merge(m, src)
}
func (m *BootHoldInfo) XXX_Size() int {
	return xxx_messageInfo_BootHoldInfo.Size(m)
}
func (m *BootHoldInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BootHoldInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BootHoldInfo proto.InternalMessageInfo

func (m *BootHoldInfo) GetBootHeldByName() string {
	if m != nil {
		return m.BootHeldByName
	}
	return ""
}

func (m *BootHoldInfo) GetInstanceId() uint32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *BootHoldInfo) GetJid() uint32 {
	if m != nil {
		return m.Jid
	}
	return 0
}

type ProcBootStalledInfo struct {
	SpawnStatus          string          `protobuf:"bytes,50,opt,name=spawn_status,json=spawnStatus,proto3" json:"spawn_status,omitempty"`
	BootHold             []*BootHoldInfo `protobuf:"bytes,51,rep,name=boot_hold,json=bootHold,proto3" json:"boot_hold,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProcBootStalledInfo) Reset()         { *m = ProcBootStalledInfo{} }
func (m *ProcBootStalledInfo) String() string { return proto.CompactTextString(m) }
func (*ProcBootStalledInfo) ProtoMessage()    {}
func (*ProcBootStalledInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b27412dfd269c426, []int{2}
}

func (m *ProcBootStalledInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcBootStalledInfo.Unmarshal(m, b)
}
func (m *ProcBootStalledInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcBootStalledInfo.Marshal(b, m, deterministic)
}
func (m *ProcBootStalledInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcBootStalledInfo.Merge(m, src)
}
func (m *ProcBootStalledInfo) XXX_Size() int {
	return xxx_messageInfo_ProcBootStalledInfo.Size(m)
}
func (m *ProcBootStalledInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcBootStalledInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcBootStalledInfo proto.InternalMessageInfo

func (m *ProcBootStalledInfo) GetSpawnStatus() string {
	if m != nil {
		return m.SpawnStatus
	}
	return ""
}

func (m *ProcBootStalledInfo) GetBootHold() []*BootHoldInfo {
	if m != nil {
		return m.BootHold
	}
	return nil
}

func init() {
	proto.RegisterType((*ProcBootStalledInfo_KEYS)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.boot_stalled.proc_boot_stalled_info_KEYS")
	proto.RegisterType((*BootHoldInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.boot_stalled.boot_hold_info")
	proto.RegisterType((*ProcBootStalledInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.boot_stalled.proc_boot_stalled_info")
}

func init() { proto.RegisterFile("proc_boot_stalled_info.proto", fileDescriptor_b27412dfd269c426) }

var fileDescriptor_b27412dfd269c426 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x0b, 0xb2, 0x9b, 0xaa, 0x68, 0x0e, 0x52, 0x58, 0xc1, 0xda, 0x53, 0xbd, 0xe4,
	0xb0, 0x7b, 0xf3, 0x28, 0x0a, 0x8a, 0xe0, 0xa1, 0xeb, 0xc5, 0xd3, 0x90, 0x36, 0xa3, 0x46, 0xd2,
	0x4c, 0xed, 0x44, 0xb4, 0x2f, 0xe5, 0x33, 0x4a, 0xb3, 0x28, 0x2b, 0xec, 0xd1, 0xdb, 0xf0, 0x25,
	0x5f, 0x66, 0xf2, 0x8f, 0x38, 0xe9, 0x7a, 0x6a, 0xa0, 0x26, 0x0a, 0xc0, 0x41, 0x3b, 0x87, 0x06,
	0xac, 0x7f, 0x22, 0xd5, 0xf5, 0x14, 0x48, 0x5e, 0x35, 0x96, 0x1b, 0x02, 0x4b, 0x0c, 0x9f, 0x3d,
	0xf0, 0xc0, 0xed, 0x73, 0x0f, 0xd4, 0x61, 0xaf, 0x78, 0xe0, 0x80, 0x2d, 0x8c, 0x36, 0x32, 0x2b,
	0x4f, 0x06, 0x21, 0xe8, 0xda, 0x61, 0x2c, 0xd5, 0xe6, 0x83, 0xc5, 0x85, 0x98, 0x6f, 0xef, 0x02,
	0x77, 0xd7, 0x8f, 0x2b, 0x39, 0x17, 0xb3, 0xa8, 0x7b, 0xdd, 0x62, 0x96, 0xe4, 0x49, 0x39, 0xab,
	0xa6, 0x23, 0xb8, 0xd7, 0x2d, 0x16, 0x5e, 0x1c, 0x44, 0xed, 0x85, 0xdc, 0xda, 0x91, 0xe7, 0xe2,
	0x68, 0x4d, 0xd0, 0x19, 0xa8, 0x87, 0x4d, 0x2d, 0x5e, 0xbd, 0x41, 0x67, 0x2e, 0x87, 0x51, 0x96,
	0xa7, 0x22, 0xb5, 0x9e, 0x83, 0xf6, 0x0d, 0x82, 0x35, 0xd9, 0x4e, 0x9e, 0x94, 0xfb, 0x95, 0xf8,
	0x41, 0xb7, 0x46, 0x1e, 0x8a, 0xc9, 0xab, 0x35, 0xd9, 0x24, 0x1e, 0x8c, 0x65, 0xf1, 0x95, 0x88,
	0xe3, 0xed, 0xc3, 0xca, 0x33, 0xb1, 0xc7, 0x9d, 0xfe, 0xf0, 0x23, 0x0d, 0xef, 0x9c, 0x2d, 0x62,
	0xcf, 0x34, 0xb2, 0x55, 0x44, 0xf2, 0x4d, 0xcc, 0x7e, 0xa7, 0xcd, 0x96, 0xf9, 0xa4, 0x4c, 0x17,
	0x0f, 0xea, 0x3f, 0x32, 0x54, 0x7f, 0x43, 0xa8, 0xa6, 0xf1, 0xa7, 0xe4, 0x4c, 0xbd, 0x1b, 0x37,
	0xb5, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x65, 0xd9, 0xe8, 0xc9, 0x01, 0x00, 0x00,
}
