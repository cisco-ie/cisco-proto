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
// source: rsvp_mgmt_glbl_nbr_compact.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_global_neighbor_briefs_global_neighbor_brief

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

type RsvpMgmtGlblNbrCompact_KEYS struct {
	NeighborAddress      string   `protobuf:"bytes,1,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtGlblNbrCompact_KEYS) Reset()         { *m = RsvpMgmtGlblNbrCompact_KEYS{} }
func (m *RsvpMgmtGlblNbrCompact_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtGlblNbrCompact_KEYS) ProtoMessage()    {}
func (*RsvpMgmtGlblNbrCompact_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a433b059f4147, []int{0}
}

func (m *RsvpMgmtGlblNbrCompact_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtGlblNbrCompact_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtGlblNbrCompact_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtGlblNbrCompact_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtGlblNbrCompact_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtGlblNbrCompact_KEYS.Merge(m, src)
}
func (m *RsvpMgmtGlblNbrCompact_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtGlblNbrCompact_KEYS.Size(m)
}
func (m *RsvpMgmtGlblNbrCompact_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtGlblNbrCompact_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtGlblNbrCompact_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtGlblNbrCompact_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type RsvpMgmtGlblNbrFlags struct {
	IsApplicationOuni    bool     `protobuf:"varint,1,opt,name=is_application_ouni,json=isApplicationOuni,proto3" json:"is_application_ouni,omitempty"`
	IsApplicationMpls    bool     `protobuf:"varint,2,opt,name=is_application_mpls,json=isApplicationMpls,proto3" json:"is_application_mpls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtGlblNbrFlags) Reset()         { *m = RsvpMgmtGlblNbrFlags{} }
func (m *RsvpMgmtGlblNbrFlags) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtGlblNbrFlags) ProtoMessage()    {}
func (*RsvpMgmtGlblNbrFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a433b059f4147, []int{1}
}

func (m *RsvpMgmtGlblNbrFlags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtGlblNbrFlags.Unmarshal(m, b)
}
func (m *RsvpMgmtGlblNbrFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtGlblNbrFlags.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtGlblNbrFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtGlblNbrFlags.Merge(m, src)
}
func (m *RsvpMgmtGlblNbrFlags) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtGlblNbrFlags.Size(m)
}
func (m *RsvpMgmtGlblNbrFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtGlblNbrFlags.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtGlblNbrFlags proto.InternalMessageInfo

func (m *RsvpMgmtGlblNbrFlags) GetIsApplicationOuni() bool {
	if m != nil {
		return m.IsApplicationOuni
	}
	return false
}

func (m *RsvpMgmtGlblNbrFlags) GetIsApplicationMpls() bool {
	if m != nil {
		return m.IsApplicationMpls
	}
	return false
}

type RsvpMgmtTimespec struct {
	Seconds              int32    `protobuf:"zigzag32,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanoseconds          int32    `protobuf:"zigzag32,2,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtTimespec) Reset()         { *m = RsvpMgmtTimespec{} }
func (m *RsvpMgmtTimespec) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtTimespec) ProtoMessage()    {}
func (*RsvpMgmtTimespec) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a433b059f4147, []int{2}
}

func (m *RsvpMgmtTimespec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtTimespec.Unmarshal(m, b)
}
func (m *RsvpMgmtTimespec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtTimespec.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtTimespec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtTimespec.Merge(m, src)
}
func (m *RsvpMgmtTimespec) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtTimespec.Size(m)
}
func (m *RsvpMgmtTimespec) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtTimespec.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtTimespec proto.InternalMessageInfo

func (m *RsvpMgmtTimespec) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *RsvpMgmtTimespec) GetNanoseconds() int32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

type RsvpMgmtGlblNbrCompact struct {
	IsGrEnabled             bool                  `protobuf:"varint,50,opt,name=is_gr_enabled,json=isGrEnabled,proto3" json:"is_gr_enabled,omitempty"`
	NodeAddress             string                `protobuf:"bytes,51,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	GlobalNeighborFlags     *RsvpMgmtGlblNbrFlags `protobuf:"bytes,52,opt,name=global_neighbor_flags,json=globalNeighborFlags,proto3" json:"global_neighbor_flags,omitempty"`
	LocalNodeAddress        []string              `protobuf:"bytes,53,rep,name=local_node_address,json=localNodeAddress,proto3" json:"local_node_address,omitempty"`
	RestartState            string                `protobuf:"bytes,54,opt,name=restart_state,json=restartState,proto3" json:"restart_state,omitempty"`
	NeighborHelloState      []string              `protobuf:"bytes,55,rep,name=neighbor_hello_state,json=neighborHelloState,proto3" json:"neighbor_hello_state,omitempty"`
	UpTime                  []*RsvpMgmtTimespec   `protobuf:"bytes,56,rep,name=up_time,json=upTime,proto3" json:"up_time,omitempty"`
	LostCommunicationTime   []*RsvpMgmtTimespec   `protobuf:"bytes,57,rep,name=lost_communication_time,json=lostCommunicationTime,proto3" json:"lost_communication_time,omitempty"`
	LostCommunicationReason []string              `protobuf:"bytes,58,rep,name=lost_communication_reason,json=lostCommunicationReason,proto3" json:"lost_communication_reason,omitempty"`
	LostCommunicationTotal  []uint32              `protobuf:"varint,59,rep,packed,name=lost_communication_total,json=lostCommunicationTotal,proto3" json:"lost_communication_total,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *RsvpMgmtGlblNbrCompact) Reset()         { *m = RsvpMgmtGlblNbrCompact{} }
func (m *RsvpMgmtGlblNbrCompact) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtGlblNbrCompact) ProtoMessage()    {}
func (*RsvpMgmtGlblNbrCompact) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a433b059f4147, []int{3}
}

func (m *RsvpMgmtGlblNbrCompact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtGlblNbrCompact.Unmarshal(m, b)
}
func (m *RsvpMgmtGlblNbrCompact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtGlblNbrCompact.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtGlblNbrCompact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtGlblNbrCompact.Merge(m, src)
}
func (m *RsvpMgmtGlblNbrCompact) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtGlblNbrCompact.Size(m)
}
func (m *RsvpMgmtGlblNbrCompact) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtGlblNbrCompact.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtGlblNbrCompact proto.InternalMessageInfo

func (m *RsvpMgmtGlblNbrCompact) GetIsGrEnabled() bool {
	if m != nil {
		return m.IsGrEnabled
	}
	return false
}

func (m *RsvpMgmtGlblNbrCompact) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *RsvpMgmtGlblNbrCompact) GetGlobalNeighborFlags() *RsvpMgmtGlblNbrFlags {
	if m != nil {
		return m.GlobalNeighborFlags
	}
	return nil
}

func (m *RsvpMgmtGlblNbrCompact) GetLocalNodeAddress() []string {
	if m != nil {
		return m.LocalNodeAddress
	}
	return nil
}

func (m *RsvpMgmtGlblNbrCompact) GetRestartState() string {
	if m != nil {
		return m.RestartState
	}
	return ""
}

func (m *RsvpMgmtGlblNbrCompact) GetNeighborHelloState() []string {
	if m != nil {
		return m.NeighborHelloState
	}
	return nil
}

func (m *RsvpMgmtGlblNbrCompact) GetUpTime() []*RsvpMgmtTimespec {
	if m != nil {
		return m.UpTime
	}
	return nil
}

func (m *RsvpMgmtGlblNbrCompact) GetLostCommunicationTime() []*RsvpMgmtTimespec {
	if m != nil {
		return m.LostCommunicationTime
	}
	return nil
}

func (m *RsvpMgmtGlblNbrCompact) GetLostCommunicationReason() []string {
	if m != nil {
		return m.LostCommunicationReason
	}
	return nil
}

func (m *RsvpMgmtGlblNbrCompact) GetLostCommunicationTotal() []uint32 {
	if m != nil {
		return m.LostCommunicationTotal
	}
	return nil
}

func init() {
	proto.RegisterType((*RsvpMgmtGlblNbrCompact_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.global_neighbor_briefs.global_neighbor_brief.rsvp_mgmt_glbl_nbr_compact_KEYS")
	proto.RegisterType((*RsvpMgmtGlblNbrFlags)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.global_neighbor_briefs.global_neighbor_brief.rsvp_mgmt_glbl_nbr_flags")
	proto.RegisterType((*RsvpMgmtTimespec)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.global_neighbor_briefs.global_neighbor_brief.rsvp_mgmt_timespec")
	proto.RegisterType((*RsvpMgmtGlblNbrCompact)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.global_neighbor_briefs.global_neighbor_brief.rsvp_mgmt_glbl_nbr_compact")
}

func init() { proto.RegisterFile("rsvp_mgmt_glbl_nbr_compact.proto", fileDescriptor_767a433b059f4147) }

var fileDescriptor_767a433b059f4147 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x14, 0xd4, 0xb6, 0xd2, 0x96, 0xbe, 0xed, 0x8a, 0xd6, 0xa5, 0xd4, 0x70, 0x21, 0x84, 0x4b, 0x90,
	0x50, 0x84, 0xb6, 0x7c, 0x94, 0x72, 0xaa, 0x50, 0x01, 0xa9, 0x50, 0x50, 0xca, 0x85, 0x93, 0xe5,
	0x24, 0x6e, 0x6a, 0xc9, 0xb1, 0x2d, 0x3f, 0x07, 0x55, 0xfc, 0x08, 0xae, 0x88, 0x7f, 0x8b, 0xe2,
	0x6c, 0xb6, 0x2b, 0x9a, 0xbd, 0x01, 0xb7, 0xdd, 0x99, 0xf7, 0x31, 0x6f, 0xc6, 0x0a, 0x44, 0x0e,
	0xbf, 0x59, 0x56, 0x57, 0xb5, 0x67, 0x95, 0xca, 0x15, 0xd3, 0xb9, 0x63, 0x85, 0xa9, 0x2d, 0x2f,
	0x7c, 0x6a, 0x9d, 0xf1, 0x86, 0x9c, 0x16, 0x12, 0x0b, 0xc3, 0xa4, 0x41, 0x76, 0xe5, 0x98, 0xb4,
	0x2c, 0x74, 0x18, 0x2b, 0x5c, 0xda, 0xfe, 0x4a, 0x2b, 0x65, 0x72, 0xae, 0x98, 0x16, 0xb2, 0xba,
	0xcc, 0x8d, 0x63, 0xb9, 0x93, 0xe2, 0x02, 0x87, 0xe1, 0xf8, 0x03, 0x3c, 0x58, 0xbd, 0x90, 0x9d,
	0x9e, 0x7c, 0x3d, 0x27, 0x8f, 0x61, 0x7b, 0xd1, 0xc4, 0xcb, 0xd2, 0x09, 0x44, 0x3a, 0x8a, 0x46,
	0xc9, 0x66, 0x76, 0xbb, 0xc7, 0x8f, 0x3b, 0x38, 0xfe, 0x0e, 0x74, 0x60, 0xda, 0x85, 0xe2, 0x15,
	0x92, 0x14, 0x76, 0x25, 0x32, 0x6e, 0xad, 0x92, 0x05, 0xf7, 0xd2, 0x68, 0x66, 0x1a, 0x2d, 0xc3,
	0xa4, 0x5b, 0xd9, 0x8e, 0xc4, 0xe3, 0x6b, 0xe6, 0x53, 0xa3, 0xe5, 0x40, 0x7d, 0x6d, 0x15, 0xd2,
	0xb5, 0x81, 0xfa, 0x8f, 0x56, 0x61, 0xfc, 0x19, 0xc8, 0xf5, 0x6e, 0x2f, 0x6b, 0x81, 0x56, 0x14,
	0x84, 0xc2, 0x06, 0x8a, 0xc2, 0xe8, 0xb2, 0xd3, 0xbc, 0x93, 0xf5, 0x7f, 0x49, 0x04, 0x13, 0xcd,
	0xb5, 0xe9, 0xd9, 0xb5, 0xc0, 0x2e, 0x43, 0xf1, 0x8f, 0x31, 0xdc, 0x5f, 0x6d, 0x0e, 0x89, 0x61,
	0x2a, 0x91, 0x55, 0x8e, 0x09, 0xcd, 0x73, 0x25, 0x4a, 0x3a, 0x0b, 0xd2, 0x26, 0x12, 0xdf, 0xb9,
	0x93, 0x0e, 0x22, 0x0f, 0x61, 0x4b, 0x9b, 0x52, 0x2c, 0x7c, 0x3b, 0x08, 0xbe, 0x4d, 0x5a, 0x6c,
	0xee, 0x19, 0xf9, 0x35, 0x82, 0xbd, 0x3f, 0xb3, 0x09, 0x8e, 0xd1, 0x67, 0xd1, 0x28, 0x99, 0xcc,
	0x44, 0xfa, 0x17, 0xf3, 0x4e, 0x57, 0xc5, 0x93, 0xed, 0x76, 0xf5, 0x67, 0xf3, 0xf2, 0xb7, 0x21,
	0xb3, 0x27, 0x40, 0x94, 0x29, 0xda, 0x29, 0xcb, 0x47, 0x3c, 0x8f, 0xd6, 0x93, 0xcd, 0x6c, 0x3b,
	0x30, 0x67, 0x4b, 0x97, 0x3c, 0x82, 0xa9, 0x13, 0xe8, 0xb9, 0xf3, 0x0c, 0x3d, 0xf7, 0x82, 0xbe,
	0x08, 0xd7, 0x6e, 0xcd, 0xc1, 0xf3, 0x16, 0x23, 0x4f, 0xe1, 0xce, 0x42, 0xd2, 0xa5, 0x50, 0xca,
	0xcc, 0x6b, 0x5f, 0x86, 0xa1, 0xa4, 0xe7, 0xde, 0xb7, 0x54, 0xd7, 0x71, 0x05, 0x1b, 0x8d, 0x0d,
	0x89, 0xd2, 0xc3, 0x68, 0x3d, 0x99, 0xcc, 0xd8, 0x3f, 0x72, 0xa4, 0x7f, 0x34, 0xd9, 0xb8, 0xb1,
	0x5f, 0x64, 0x2d, 0xc8, 0xcf, 0x11, 0xec, 0x2b, 0x83, 0xbe, 0x8d, 0xbc, 0x6e, 0x74, 0xff, 0x0e,
	0x83, 0x94, 0x57, 0xff, 0x47, 0xca, 0x5e, 0xbb, 0xff, 0xcd, 0xf2, 0xfa, 0xa0, 0xec, 0x08, 0xee,
	0x0d, 0x08, 0x73, 0x82, 0xa3, 0xd1, 0xf4, 0x28, 0x58, 0xb9, 0x7f, 0xa3, 0x33, 0x0b, 0x34, 0x39,
	0x04, 0x3a, 0x74, 0x94, 0xf1, 0x5c, 0xd1, 0xd7, 0xd1, 0x7a, 0x32, 0xcd, 0xee, 0xde, 0x5c, 0xda,
	0xb2, 0xf9, 0x38, 0x7c, 0x80, 0x0e, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x72, 0x78, 0xa5, 0x85,
	0xa4, 0x04, 0x00, 0x00,
}
