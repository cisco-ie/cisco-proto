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
// source: tcp_nsr_sscb_brief_bag.proto

package cisco_ios_xr_ip_tcp_oper_tcp_nsr_nodes_node_session_set_brief_sets_brief_set

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

type TcpNsrSscbBriefBag_KEYS struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Id_1                 string   `protobuf:"bytes,2,opt,name=id_1,json=id1,proto3" json:"id_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpNsrSscbBriefBag_KEYS) Reset()         { *m = TcpNsrSscbBriefBag_KEYS{} }
func (m *TcpNsrSscbBriefBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TcpNsrSscbBriefBag_KEYS) ProtoMessage()    {}
func (*TcpNsrSscbBriefBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bdacf1c7ababfa6, []int{0}
}

func (m *TcpNsrSscbBriefBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrSscbBriefBag_KEYS.Unmarshal(m, b)
}
func (m *TcpNsrSscbBriefBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrSscbBriefBag_KEYS.Marshal(b, m, deterministic)
}
func (m *TcpNsrSscbBriefBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrSscbBriefBag_KEYS.Merge(m, src)
}
func (m *TcpNsrSscbBriefBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TcpNsrSscbBriefBag_KEYS.Size(m)
}
func (m *TcpNsrSscbBriefBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrSscbBriefBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrSscbBriefBag_KEYS proto.InternalMessageInfo

func (m *TcpNsrSscbBriefBag_KEYS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TcpNsrSscbBriefBag_KEYS) GetId_1() string {
	if m != nil {
		return m.Id_1
	}
	return ""
}

type TcpNsrSscbBriefBag struct {
	Sscb                             uint64   `protobuf:"varint,50,opt,name=sscb,proto3" json:"sscb,omitempty"`
	Pid                              uint32   `protobuf:"varint,51,opt,name=pid,proto3" json:"pid,omitempty"`
	ClientName                       []string `protobuf:"bytes,52,rep,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	ClientInstance                   uint32   `protobuf:"varint,53,opt,name=client_instance,json=clientInstance,proto3" json:"client_instance,omitempty"`
	SetId                            uint32   `protobuf:"varint,54,opt,name=set_id,json=setId,proto3" json:"set_id,omitempty"`
	SsoRole                          uint32   `protobuf:"varint,55,opt,name=sso_role,json=ssoRole,proto3" json:"sso_role,omitempty"`
	Mode                             uint32   `protobuf:"varint,56,opt,name=mode,proto3" json:"mode,omitempty"`
	AddressFamily                    string   `protobuf:"bytes,57,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	WellKnownPort                    uint32   `protobuf:"varint,58,opt,name=well_known_port,json=wellKnownPort,proto3" json:"well_known_port,omitempty"`
	LocalNode                        string   `protobuf:"bytes,59,opt,name=local_node,json=localNode,proto3" json:"local_node,omitempty"`
	LocalInstance                    uint32   `protobuf:"varint,60,opt,name=local_instance,json=localInstance,proto3" json:"local_instance,omitempty"`
	ProtectNode                      string   `protobuf:"bytes,61,opt,name=protect_node,json=protectNode,proto3" json:"protect_node,omitempty"`
	ProtectInstance                  uint32   `protobuf:"varint,62,opt,name=protect_instance,json=protectInstance,proto3" json:"protect_instance,omitempty"`
	NumberOfSessions                 uint32   `protobuf:"varint,63,opt,name=number_of_sessions,json=numberOfSessions,proto3" json:"number_of_sessions,omitempty"`
	NumberOfSyncedSessionsUpStream   uint32   `protobuf:"varint,64,opt,name=number_of_synced_sessions_up_stream,json=numberOfSyncedSessionsUpStream,proto3" json:"number_of_synced_sessions_up_stream,omitempty"`
	NumberOfSyncedSessionsDownStream uint32   `protobuf:"varint,65,opt,name=number_of_synced_sessions_down_stream,json=numberOfSyncedSessionsDownStream,proto3" json:"number_of_synced_sessions_down_stream,omitempty"`
	IsInitSyncInProgress             bool     `protobuf:"varint,66,opt,name=is_init_sync_in_progress,json=isInitSyncInProgress,proto3" json:"is_init_sync_in_progress,omitempty"`
	IsSscbInitSyncReady              bool     `protobuf:"varint,67,opt,name=is_sscb_init_sync_ready,json=isSscbInitSyncReady,proto3" json:"is_sscb_init_sync_ready,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *TcpNsrSscbBriefBag) Reset()         { *m = TcpNsrSscbBriefBag{} }
func (m *TcpNsrSscbBriefBag) String() string { return proto.CompactTextString(m) }
func (*TcpNsrSscbBriefBag) ProtoMessage()    {}
func (*TcpNsrSscbBriefBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bdacf1c7ababfa6, []int{1}
}

func (m *TcpNsrSscbBriefBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrSscbBriefBag.Unmarshal(m, b)
}
func (m *TcpNsrSscbBriefBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrSscbBriefBag.Marshal(b, m, deterministic)
}
func (m *TcpNsrSscbBriefBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrSscbBriefBag.Merge(m, src)
}
func (m *TcpNsrSscbBriefBag) XXX_Size() int {
	return xxx_messageInfo_TcpNsrSscbBriefBag.Size(m)
}
func (m *TcpNsrSscbBriefBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrSscbBriefBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrSscbBriefBag proto.InternalMessageInfo

func (m *TcpNsrSscbBriefBag) GetSscb() uint64 {
	if m != nil {
		return m.Sscb
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetClientName() []string {
	if m != nil {
		return m.ClientName
	}
	return nil
}

func (m *TcpNsrSscbBriefBag) GetClientInstance() uint32 {
	if m != nil {
		return m.ClientInstance
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetSetId() uint32 {
	if m != nil {
		return m.SetId
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetSsoRole() uint32 {
	if m != nil {
		return m.SsoRole
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *TcpNsrSscbBriefBag) GetWellKnownPort() uint32 {
	if m != nil {
		return m.WellKnownPort
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetLocalNode() string {
	if m != nil {
		return m.LocalNode
	}
	return ""
}

func (m *TcpNsrSscbBriefBag) GetLocalInstance() uint32 {
	if m != nil {
		return m.LocalInstance
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetProtectNode() string {
	if m != nil {
		return m.ProtectNode
	}
	return ""
}

func (m *TcpNsrSscbBriefBag) GetProtectInstance() uint32 {
	if m != nil {
		return m.ProtectInstance
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetNumberOfSessions() uint32 {
	if m != nil {
		return m.NumberOfSessions
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetNumberOfSyncedSessionsUpStream() uint32 {
	if m != nil {
		return m.NumberOfSyncedSessionsUpStream
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetNumberOfSyncedSessionsDownStream() uint32 {
	if m != nil {
		return m.NumberOfSyncedSessionsDownStream
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetIsInitSyncInProgress() bool {
	if m != nil {
		return m.IsInitSyncInProgress
	}
	return false
}

func (m *TcpNsrSscbBriefBag) GetIsSscbInitSyncReady() bool {
	if m != nil {
		return m.IsSscbInitSyncReady
	}
	return false
}

func init() {
	proto.RegisterType((*TcpNsrSscbBriefBag_KEYS)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.session_set.brief_sets.brief_set.tcp_nsr_sscb_brief_bag_KEYS")
	proto.RegisterType((*TcpNsrSscbBriefBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.session_set.brief_sets.brief_set.tcp_nsr_sscb_brief_bag")
}

func init() { proto.RegisterFile("tcp_nsr_sscb_brief_bag.proto", fileDescriptor_6bdacf1c7ababfa6) }

var fileDescriptor_6bdacf1c7ababfa6 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5b, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xd5, 0xad, 0x1b, 0xab, 0x47, 0x2f, 0x98, 0x9b, 0x11, 0xb7, 0x30, 0x34, 0x28, 0x12,
	0xaa, 0x34, 0x36, 0xc6, 0x1d, 0xc6, 0x55, 0xaa, 0x8a, 0xb6, 0x29, 0x15, 0x0f, 0x3c, 0x1d, 0x25,
	0xb1, 0x3b, 0x59, 0x24, 0x76, 0xe4, 0xe3, 0xa9, 0xf4, 0x33, 0xf3, 0x25, 0x90, 0x4f, 0xd2, 0x8c,
	0x87, 0xf2, 0x12, 0x1d, 0xfd, 0xce, 0xf9, 0xff, 0x1c, 0xc7, 0x0e, 0xbb, 0xe3, 0xb3, 0x12, 0x0c,
	0x3a, 0x40, 0xcc, 0x52, 0x48, 0x9d, 0x56, 0x33, 0x48, 0x93, 0xb3, 0x51, 0xe9, 0xac, 0xb7, 0xfc,
	0x7b, 0xa6, 0x31, 0xb3, 0xa0, 0x2d, 0xc2, 0x6f, 0x07, 0xba, 0x84, 0x30, 0x6d, 0x4b, 0xe5, 0x46,
	0x75, 0x6c, 0x64, 0xac, 0x54, 0x48, 0xcf, 0x11, 0x2a, 0x44, 0x6d, 0x0d, 0xa0, 0xf2, 0xa3, 0x4a,
	0x84, 0xca, 0xe3, 0x45, 0xb9, 0x73, 0xc4, 0x6e, 0xaf, 0x5e, 0x0d, 0x26, 0x5f, 0x7f, 0x4e, 0x79,
	0x8f, 0xad, 0x69, 0x29, 0x5a, 0x51, 0x6b, 0xd8, 0x89, 0xd7, 0xb4, 0xe4, 0x57, 0x58, 0x5b, 0x4b,
	0xd8, 0x13, 0x6b, 0x44, 0xd6, 0xb5, 0xdc, 0xdb, 0xf9, 0xb3, 0xc1, 0x6e, 0xac, 0x56, 0x70, 0xce,
	0xda, 0x81, 0x88, 0x67, 0x51, 0x6b, 0xd8, 0x8e, 0xa9, 0xe6, 0x03, 0xb6, 0x5e, 0x6a, 0x29, 0xf6,
	0xa3, 0xd6, 0xb0, 0x1b, 0x87, 0x92, 0xdf, 0x67, 0xdb, 0x59, 0xae, 0x95, 0xf1, 0x60, 0x92, 0x42,
	0x89, 0x83, 0x68, 0x7d, 0xd8, 0x89, 0x59, 0x85, 0x8e, 0x93, 0x42, 0xf1, 0xc7, 0xac, 0x5f, 0x0f,
	0x68, 0x83, 0x3e, 0x31, 0x99, 0x12, 0xcf, 0x29, 0xde, 0xab, 0xf0, 0xb8, 0xa6, 0xfc, 0x3a, 0xdb,
	0x44, 0xe5, 0x41, 0x4b, 0x71, 0x48, 0xfd, 0x0d, 0x54, 0x7e, 0x2c, 0xf9, 0x2d, 0xb6, 0x85, 0x68,
	0xc1, 0xd9, 0x5c, 0x89, 0x17, 0xd4, 0xb8, 0x84, 0x68, 0x63, 0x9b, 0xab, 0xf0, 0x86, 0x85, 0x95,
	0x4a, 0xbc, 0x24, 0x4c, 0x35, 0xdf, 0x65, 0xbd, 0x44, 0x4a, 0xa7, 0x10, 0x61, 0x96, 0x14, 0x3a,
	0x5f, 0x88, 0x57, 0xb4, 0xdb, 0x6e, 0x4d, 0xbf, 0x11, 0xe4, 0x8f, 0x58, 0x7f, 0xae, 0xf2, 0x1c,
	0x7e, 0x19, 0x3b, 0x37, 0x50, 0x5a, 0xe7, 0xc5, 0x6b, 0xb2, 0x74, 0x03, 0x9e, 0x04, 0x7a, 0x6a,
	0x9d, 0xe7, 0x77, 0x19, 0xcb, 0x6d, 0x96, 0xe4, 0x10, 0x0e, 0x44, 0xbc, 0x21, 0x55, 0x87, 0xc8,
	0x71, 0xbd, 0x5a, 0xd5, 0x6e, 0xf6, 0xf6, 0xb6, 0xb2, 0x10, 0x6d, 0xb6, 0xf6, 0x80, 0x5d, 0x0e,
	0xc7, 0xaf, 0x32, 0x5f, 0x79, 0xde, 0x91, 0x67, 0xbb, 0x66, 0x64, 0x7a, 0xc2, 0x06, 0xcb, 0x91,
	0xc6, 0xf5, 0x9e, 0x5c, 0xfd, 0x9a, 0x37, 0xb6, 0xa7, 0x8c, 0x9b, 0xf3, 0x22, 0x55, 0x0e, 0x6c,
	0xb8, 0x06, 0x74, 0x47, 0x50, 0x7c, 0xa0, 0xe1, 0x41, 0xd5, 0x39, 0x99, 0x4d, 0x6b, 0xce, 0x27,
	0xec, 0xe1, 0x3f, 0xd3, 0x0b, 0x93, 0x29, 0xd9, 0x84, 0xe0, 0xbc, 0x04, 0xf4, 0x4e, 0x25, 0x85,
	0x38, 0xa2, 0xf8, 0xbd, 0x26, 0x4e, 0x83, 0x4b, 0xc9, 0x8f, 0x72, 0x4a, 0x53, 0xfc, 0x84, 0xed,
	0xfe, 0x5f, 0x26, 0xc3, 0xa7, 0xac, 0x75, 0x1f, 0x49, 0x17, 0xad, 0xd6, 0x7d, 0xb1, 0x73, 0x53,
	0x0b, 0x0f, 0x99, 0xd0, 0x08, 0xda, 0x68, 0x4f, 0x3a, 0xd0, 0x06, 0x4a, 0x67, 0xcf, 0xc2, 0x49,
	0x89, 0x4f, 0x51, 0x6b, 0xb8, 0x15, 0x5f, 0xd3, 0x38, 0x36, 0xda, 0x07, 0xc3, 0xd8, 0x9c, 0xd6,
	0x3d, 0x7e, 0xc0, 0x6e, 0x6a, 0xac, 0x6e, 0xec, 0x45, 0xd8, 0xa9, 0x44, 0x2e, 0xc4, 0x67, 0x8a,
	0x5d, 0xd5, 0x38, 0xc5, 0x2c, 0x5d, 0x46, 0xe3, 0xd0, 0x4a, 0x37, 0xe9, 0x27, 0xdc, 0xff, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xba, 0xdb, 0xe7, 0x10, 0xa4, 0x03, 0x00, 0x00,
}