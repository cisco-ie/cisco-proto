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
// source: oc_mpls_te_lsp_counters.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_open_config_lsp_counters_lsp_counter

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

type OcMplsTeLspCounters_KEYS struct {
	TunnelName           string   `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	TunnelType           string   `protobuf:"bytes,2,opt,name=tunnel_type,json=tunnelType,proto3" json:"tunnel_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OcMplsTeLspCounters_KEYS) Reset()         { *m = OcMplsTeLspCounters_KEYS{} }
func (m *OcMplsTeLspCounters_KEYS) String() string { return proto.CompactTextString(m) }
func (*OcMplsTeLspCounters_KEYS) ProtoMessage()    {}
func (*OcMplsTeLspCounters_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f095a8d7436e24e, []int{0}
}

func (m *OcMplsTeLspCounters_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OcMplsTeLspCounters_KEYS.Unmarshal(m, b)
}
func (m *OcMplsTeLspCounters_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OcMplsTeLspCounters_KEYS.Marshal(b, m, deterministic)
}
func (m *OcMplsTeLspCounters_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OcMplsTeLspCounters_KEYS.Merge(m, src)
}
func (m *OcMplsTeLspCounters_KEYS) XXX_Size() int {
	return xxx_messageInfo_OcMplsTeLspCounters_KEYS.Size(m)
}
func (m *OcMplsTeLspCounters_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OcMplsTeLspCounters_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OcMplsTeLspCounters_KEYS proto.InternalMessageInfo

func (m *OcMplsTeLspCounters_KEYS) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *OcMplsTeLspCounters_KEYS) GetTunnelType() string {
	if m != nil {
		return m.TunnelType
	}
	return ""
}

type OcMplsTeLspCounters struct {
	Name                   string   `protobuf:"bytes,50,opt,name=name,proto3" json:"name,omitempty"`
	Type                   string   `protobuf:"bytes,51,opt,name=type,proto3" json:"type,omitempty"`
	Bytes                  uint64   `protobuf:"varint,52,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Packets                uint64   `protobuf:"varint,53,opt,name=packets,proto3" json:"packets,omitempty"`
	PathChanges            uint64   `protobuf:"varint,54,opt,name=path_changes,json=pathChanges,proto3" json:"path_changes,omitempty"`
	StateChanges           uint64   `protobuf:"varint,55,opt,name=state_changes,json=stateChanges,proto3" json:"state_changes,omitempty"`
	OnlineTime             string   `protobuf:"bytes,56,opt,name=online_time,json=onlineTime,proto3" json:"online_time,omitempty"`
	CurrentPathTime        string   `protobuf:"bytes,57,opt,name=current_path_time,json=currentPathTime,proto3" json:"current_path_time,omitempty"`
	NextReoptimizationTime string   `protobuf:"bytes,58,opt,name=next_reoptimization_time,json=nextReoptimizationTime,proto3" json:"next_reoptimization_time,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *OcMplsTeLspCounters) Reset()         { *m = OcMplsTeLspCounters{} }
func (m *OcMplsTeLspCounters) String() string { return proto.CompactTextString(m) }
func (*OcMplsTeLspCounters) ProtoMessage()    {}
func (*OcMplsTeLspCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f095a8d7436e24e, []int{1}
}

func (m *OcMplsTeLspCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OcMplsTeLspCounters.Unmarshal(m, b)
}
func (m *OcMplsTeLspCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OcMplsTeLspCounters.Marshal(b, m, deterministic)
}
func (m *OcMplsTeLspCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OcMplsTeLspCounters.Merge(m, src)
}
func (m *OcMplsTeLspCounters) XXX_Size() int {
	return xxx_messageInfo_OcMplsTeLspCounters.Size(m)
}
func (m *OcMplsTeLspCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_OcMplsTeLspCounters.DiscardUnknown(m)
}

var xxx_messageInfo_OcMplsTeLspCounters proto.InternalMessageInfo

func (m *OcMplsTeLspCounters) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OcMplsTeLspCounters) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OcMplsTeLspCounters) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *OcMplsTeLspCounters) GetPackets() uint64 {
	if m != nil {
		return m.Packets
	}
	return 0
}

func (m *OcMplsTeLspCounters) GetPathChanges() uint64 {
	if m != nil {
		return m.PathChanges
	}
	return 0
}

func (m *OcMplsTeLspCounters) GetStateChanges() uint64 {
	if m != nil {
		return m.StateChanges
	}
	return 0
}

func (m *OcMplsTeLspCounters) GetOnlineTime() string {
	if m != nil {
		return m.OnlineTime
	}
	return ""
}

func (m *OcMplsTeLspCounters) GetCurrentPathTime() string {
	if m != nil {
		return m.CurrentPathTime
	}
	return ""
}

func (m *OcMplsTeLspCounters) GetNextReoptimizationTime() string {
	if m != nil {
		return m.NextReoptimizationTime
	}
	return ""
}

func init() {
	proto.RegisterType((*OcMplsTeLspCounters_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.open_config.lsp_counters.lsp_counter.oc_mpls_te_lsp_counters_KEYS")
	proto.RegisterType((*OcMplsTeLspCounters)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.open_config.lsp_counters.lsp_counter.oc_mpls_te_lsp_counters")
}

func init() { proto.RegisterFile("oc_mpls_te_lsp_counters.proto", fileDescriptor_1f095a8d7436e24e) }

var fileDescriptor_1f095a8d7436e24e = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4d, 0x4f, 0x02, 0x31,
	0x10, 0x86, 0x03, 0x41, 0x8d, 0x05, 0x63, 0x6c, 0x8c, 0xf6, 0xa0, 0x01, 0xf1, 0x42, 0x3c, 0xec,
	0x41, 0xfc, 0x40, 0xaf, 0x46, 0x2f, 0x26, 0xc6, 0x20, 0x17, 0x4f, 0xb5, 0x34, 0x23, 0x34, 0xee,
	0xb6, 0x4d, 0x3b, 0x24, 0xe0, 0xaf, 0xf3, 0xa7, 0x99, 0x9d, 0x02, 0x59, 0x0e, 0xdc, 0x3a, 0xcf,
	0xfb, 0xbc, 0x33, 0xd9, 0x2c, 0x3b, 0x77, 0x5a, 0x16, 0x3e, 0x8f, 0x12, 0x41, 0xe6, 0xd1, 0x4b,
	0xed, 0x66, 0x16, 0x21, 0xc4, 0xcc, 0x07, 0x87, 0x8e, 0xbf, 0x68, 0x13, 0xb5, 0x93, 0xc6, 0x45,
	0x39, 0x0f, 0x6b, 0xd1, 0x79, 0x08, 0xd9, 0x72, 0xc8, 0x9c, 0x07, 0x2b, 0xb5, 0xb3, 0xdf, 0x66,
	0x92, 0x6d, 0x6c, 0xa8, 0x0c, 0xdd, 0x2f, 0x76, 0xb6, 0xe5, 0x90, 0x7c, 0x7d, 0xfe, 0xfc, 0xe0,
	0x6d, 0xd6, 0xc4, 0x99, 0xb5, 0x90, 0x4b, 0xab, 0x0a, 0x10, 0xb5, 0x4e, 0xad, 0xb7, 0x3f, 0x64,
	0x09, 0xbd, 0xa9, 0x02, 0x2a, 0x02, 0x2e, 0x3c, 0x88, 0x7a, 0x55, 0x18, 0x2d, 0x3c, 0x74, 0xff,
	0xea, 0xec, 0x74, 0xcb, 0x09, 0xce, 0x59, 0x83, 0xd6, 0x5e, 0x53, 0x8b, 0xde, 0x25, 0xa3, 0x4d,
	0xfd, 0xc4, 0xca, 0x37, 0x3f, 0x66, 0x3b, 0xe3, 0x05, 0x42, 0x14, 0x37, 0x9d, 0x5a, 0xaf, 0x31,
	0x4c, 0x03, 0x17, 0x6c, 0xcf, 0x2b, 0xfd, 0x03, 0x18, 0xc5, 0x2d, 0xf1, 0xd5, 0xc8, 0x2f, 0x58,
	0xcb, 0x2b, 0x9c, 0x4a, 0x3d, 0x55, 0x76, 0x02, 0x51, 0xdc, 0x51, 0xdc, 0x2c, 0xd9, 0x53, 0x42,
	0xfc, 0x92, 0x1d, 0x44, 0x54, 0x08, 0x6b, 0xe7, 0x9e, 0x9c, 0x16, 0xc1, 0x95, 0xd4, 0x66, 0x4d,
	0x67, 0x73, 0x63, 0x41, 0xa2, 0x29, 0x40, 0x0c, 0xd2, 0xc7, 0x25, 0x34, 0x32, 0x05, 0xf0, 0x2b,
	0x76, 0xa4, 0x67, 0x21, 0x80, 0x45, 0x49, 0x07, 0x49, 0x7b, 0x20, 0xed, 0x70, 0x19, 0xbc, 0x2b,
	0x9c, 0x92, 0x3b, 0x60, 0xc2, 0xc2, 0x1c, 0x65, 0x00, 0xe7, 0xd1, 0x14, 0xe6, 0x57, 0xa1, 0x71,
	0x36, 0x55, 0x1e, 0xa9, 0x72, 0x52, 0xe6, 0xc3, 0x8d, 0xb8, 0x6c, 0x8e, 0x77, 0xe9, 0x9f, 0xf7,
	0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x80, 0x21, 0x9f, 0xdb, 0x14, 0x02, 0x00, 0x00,
}
