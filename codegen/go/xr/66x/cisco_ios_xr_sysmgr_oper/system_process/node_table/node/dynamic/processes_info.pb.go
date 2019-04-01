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
// source: processes_info.proto

package cisco_ios_xr_sysmgr_oper_system_process_node_table_node_dynamic

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

type ProcessesInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessesInfo_KEYS) Reset()         { *m = ProcessesInfo_KEYS{} }
func (m *ProcessesInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*ProcessesInfo_KEYS) ProtoMessage()    {}
func (*ProcessesInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f65e07866e53cde, []int{0}
}

func (m *ProcessesInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessesInfo_KEYS.Unmarshal(m, b)
}
func (m *ProcessesInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessesInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *ProcessesInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessesInfo_KEYS.Merge(m, src)
}
func (m *ProcessesInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_ProcessesInfo_KEYS.Size(m)
}
func (m *ProcessesInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessesInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessesInfo_KEYS proto.InternalMessageInfo

func (m *ProcessesInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type ProcessInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InstanceId           uint32   `protobuf:"varint,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Args                 string   `protobuf:"bytes,3,opt,name=args,proto3" json:"args,omitempty"`
	Jid                  uint32   `protobuf:"varint,4,opt,name=jid,proto3" json:"jid,omitempty"`
	State                string   `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	LastStarted          string   `protobuf:"bytes,6,opt,name=last_started,json=lastStarted,proto3" json:"last_started,omitempty"`
	RespawnCount         uint32   `protobuf:"varint,7,opt,name=respawn_count,json=respawnCount,proto3" json:"respawn_count,omitempty"`
	PlacementState       string   `protobuf:"bytes,8,opt,name=placement_state,json=placementState,proto3" json:"placement_state,omitempty"`
	IsMandatory          bool     `protobuf:"varint,9,opt,name=is_mandatory,json=isMandatory,proto3" json:"is_mandatory,omitempty"`
	IsMaintenance        bool     `protobuf:"varint,10,opt,name=is_maintenance,json=isMaintenance,proto3" json:"is_maintenance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessInfo) Reset()         { *m = ProcessInfo{} }
func (m *ProcessInfo) String() string { return proto.CompactTextString(m) }
func (*ProcessInfo) ProtoMessage()    {}
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f65e07866e53cde, []int{1}
}

func (m *ProcessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessInfo.Unmarshal(m, b)
}
func (m *ProcessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessInfo.Marshal(b, m, deterministic)
}
func (m *ProcessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessInfo.Merge(m, src)
}
func (m *ProcessInfo) XXX_Size() int {
	return xxx_messageInfo_ProcessInfo.Size(m)
}
func (m *ProcessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessInfo proto.InternalMessageInfo

func (m *ProcessInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProcessInfo) GetInstanceId() uint32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *ProcessInfo) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *ProcessInfo) GetJid() uint32 {
	if m != nil {
		return m.Jid
	}
	return 0
}

func (m *ProcessInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ProcessInfo) GetLastStarted() string {
	if m != nil {
		return m.LastStarted
	}
	return ""
}

func (m *ProcessInfo) GetRespawnCount() uint32 {
	if m != nil {
		return m.RespawnCount
	}
	return 0
}

func (m *ProcessInfo) GetPlacementState() string {
	if m != nil {
		return m.PlacementState
	}
	return ""
}

func (m *ProcessInfo) GetIsMandatory() bool {
	if m != nil {
		return m.IsMandatory
	}
	return false
}

func (m *ProcessInfo) GetIsMaintenance() bool {
	if m != nil {
		return m.IsMaintenance
	}
	return false
}

type ProcessesInfo struct {
	ProcessCount         uint32         `protobuf:"varint,50,opt,name=process_count,json=processCount,proto3" json:"process_count,omitempty"`
	Process              []*ProcessInfo `protobuf:"bytes,51,rep,name=process,proto3" json:"process,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProcessesInfo) Reset()         { *m = ProcessesInfo{} }
func (m *ProcessesInfo) String() string { return proto.CompactTextString(m) }
func (*ProcessesInfo) ProtoMessage()    {}
func (*ProcessesInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f65e07866e53cde, []int{2}
}

func (m *ProcessesInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessesInfo.Unmarshal(m, b)
}
func (m *ProcessesInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessesInfo.Marshal(b, m, deterministic)
}
func (m *ProcessesInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessesInfo.Merge(m, src)
}
func (m *ProcessesInfo) XXX_Size() int {
	return xxx_messageInfo_ProcessesInfo.Size(m)
}
func (m *ProcessesInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessesInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessesInfo proto.InternalMessageInfo

func (m *ProcessesInfo) GetProcessCount() uint32 {
	if m != nil {
		return m.ProcessCount
	}
	return 0
}

func (m *ProcessesInfo) GetProcess() []*ProcessInfo {
	if m != nil {
		return m.Process
	}
	return nil
}

func init() {
	proto.RegisterType((*ProcessesInfo_KEYS)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.dynamic.processes_info_KEYS")
	proto.RegisterType((*ProcessInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.dynamic.process_info")
	proto.RegisterType((*ProcessesInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.dynamic.processes_info")
}

func init() { proto.RegisterFile("processes_info.proto", fileDescriptor_9f65e07866e53cde) }

var fileDescriptor_9f65e07866e53cde = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbd, 0x6e, 0xdb, 0x30,
	0x14, 0x85, 0x21, 0xff, 0xfb, 0xca, 0x76, 0x0b, 0xd6, 0x03, 0x81, 0x0e, 0x55, 0x5d, 0x14, 0xd5,
	0xa4, 0xc1, 0x7e, 0x80, 0x0e, 0x45, 0x87, 0xa2, 0x70, 0x06, 0x79, 0xca, 0x44, 0xd0, 0xd2, 0x8d,
	0xc1, 0xc0, 0x22, 0x05, 0x5e, 0x06, 0x89, 0x5e, 0x26, 0x0f, 0x92, 0xa7, 0x0b, 0x48, 0x49, 0x86,
	0x3d, 0x67, 0xbb, 0xfc, 0x74, 0xce, 0xd1, 0x11, 0xaf, 0x60, 0x5d, 0x5b, 0x53, 0x20, 0x11, 0x92,
	0x50, 0xfa, 0xc1, 0x64, 0xb5, 0x35, 0xce, 0xb0, 0xdf, 0x85, 0xa2, 0xc2, 0x08, 0x65, 0x48, 0xbc,
	0x58, 0x41, 0x0d, 0x55, 0x27, 0x2b, 0x4c, 0x8d, 0x36, 0xa3, 0x86, 0x1c, 0x56, 0xa2, 0x73, 0x65,
	0xda, 0x94, 0x28, 0x9c, 0x3c, 0x9e, 0x31, 0x8c, 0x59, 0xd9, 0x68, 0x59, 0xa9, 0x62, 0xb3, 0x85,
	0x2f, 0xb7, 0xc1, 0xe2, 0xff, 0xdf, 0xfb, 0x03, 0xfb, 0x0a, 0xf3, 0xe0, 0xd0, 0xb2, 0x42, 0x1e,
	0x25, 0x51, 0x3a, 0xcf, 0x67, 0x1e, 0xdc, 0xc9, 0x0a, 0x37, 0x6f, 0x03, 0x58, 0x74, 0xa6, 0x60,
	0x61, 0x0c, 0x46, 0x57, 0xc2, 0x30, 0xb3, 0x6f, 0x10, 0x2b, 0x4d, 0x4e, 0xea, 0x02, 0x85, 0x2a,
	0xf9, 0x20, 0x89, 0xd2, 0x65, 0x0e, 0x3d, 0xfa, 0x57, 0x7a, 0x93, 0xb4, 0x27, 0xe2, 0xc3, 0xd6,
	0xe4, 0x67, 0xf6, 0x19, 0x86, 0x8f, 0xaa, 0xe4, 0xa3, 0x20, 0xf6, 0x23, 0x5b, 0xc3, 0x98, 0x9c,
	0x74, 0xc8, 0xc7, 0x41, 0xd6, 0x1e, 0xd8, 0x77, 0x58, 0x9c, 0x25, 0x39, 0x41, 0x4e, 0x5a, 0x87,
	0x25, 0x9f, 0x84, 0x87, 0xb1, 0x67, 0x87, 0x16, 0xb1, 0x1f, 0xb0, 0xb4, 0x48, 0xb5, 0x7c, 0xd6,
	0xa2, 0x30, 0x4f, 0xda, 0xf1, 0x69, 0x08, 0x5d, 0x74, 0xf0, 0x8f, 0x67, 0xec, 0x17, 0x7c, 0xaa,
	0xcf, 0xb2, 0xc0, 0x0a, 0x75, 0x08, 0x73, 0xc8, 0x67, 0x21, 0x6a, 0x75, 0xc1, 0x87, 0xfe, 0x85,
	0x8a, 0x44, 0x25, 0x75, 0x29, 0x9d, 0xb1, 0x0d, 0x9f, 0x27, 0x51, 0x3a, 0xcb, 0x63, 0x45, 0xfb,
	0x1e, 0xb1, 0x9f, 0xb0, 0x0a, 0x12, 0xa5, 0x1d, 0x6a, 0xff, 0x8d, 0x1c, 0x82, 0x68, 0xe9, 0x45,
	0x17, 0xb8, 0x79, 0x8d, 0x60, 0x75, 0x7b, 0xe3, 0xbe, 0x6a, 0x7f, 0x9d, 0x6d, 0xd5, 0x6d, 0x5b,
	0xb5, 0x83, 0x6d, 0xd5, 0x13, 0x4c, 0xbb, 0x33, 0xdf, 0x25, 0xc3, 0x34, 0xde, 0xee, 0xb3, 0x0f,
	0xee, 0x3e, 0xbb, 0xde, 0x61, 0xde, 0xa7, 0x1f, 0x27, 0xe1, 0xcf, 0xda, 0xbd, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xa6, 0x62, 0x9e, 0x91, 0x71, 0x02, 0x00, 0x00,
}