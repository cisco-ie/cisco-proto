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
// source: pcep_plsp_bag.proto

package cisco_ios_xr_infra_xtc_agent_oper_pcc_plsps_plsp

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

type PcepPlspBag_KEYS struct {
	PlspId               uint32   `protobuf:"varint,1,opt,name=plsp_id,json=plspId,proto3" json:"plsp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcepPlspBag_KEYS) Reset()         { *m = PcepPlspBag_KEYS{} }
func (m *PcepPlspBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PcepPlspBag_KEYS) ProtoMessage()    {}
func (*PcepPlspBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fedae5cf2bc176a, []int{0}
}

func (m *PcepPlspBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcepPlspBag_KEYS.Unmarshal(m, b)
}
func (m *PcepPlspBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcepPlspBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PcepPlspBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcepPlspBag_KEYS.Merge(m, src)
}
func (m *PcepPlspBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PcepPlspBag_KEYS.Size(m)
}
func (m *PcepPlspBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PcepPlspBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PcepPlspBag_KEYS proto.InternalMessageInfo

func (m *PcepPlspBag_KEYS) GetPlspId() uint32 {
	if m != nil {
		return m.PlspId
	}
	return 0
}

type PcepPlspStatsBag struct {
	PathsCreated         uint64   `protobuf:"varint,1,opt,name=paths_created,json=pathsCreated,proto3" json:"paths_created,omitempty"`
	PathsDestroyed       uint64   `protobuf:"varint,2,opt,name=paths_destroyed,json=pathsDestroyed,proto3" json:"paths_destroyed,omitempty"`
	PathCreateErrors     uint64   `protobuf:"varint,3,opt,name=path_create_errors,json=pathCreateErrors,proto3" json:"path_create_errors,omitempty"`
	PathDestroyErrors    uint64   `protobuf:"varint,4,opt,name=path_destroy_errors,json=pathDestroyErrors,proto3" json:"path_destroy_errors,omitempty"`
	RequestsCreated      uint64   `protobuf:"varint,5,opt,name=requests_created,json=requestsCreated,proto3" json:"requests_created,omitempty"`
	RequestsDestroyed    uint64   `protobuf:"varint,6,opt,name=requests_destroyed,json=requestsDestroyed,proto3" json:"requests_destroyed,omitempty"`
	RequestsFailed       uint64   `protobuf:"varint,7,opt,name=requests_failed,json=requestsFailed,proto3" json:"requests_failed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcepPlspStatsBag) Reset()         { *m = PcepPlspStatsBag{} }
func (m *PcepPlspStatsBag) String() string { return proto.CompactTextString(m) }
func (*PcepPlspStatsBag) ProtoMessage()    {}
func (*PcepPlspStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fedae5cf2bc176a, []int{1}
}

func (m *PcepPlspStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcepPlspStatsBag.Unmarshal(m, b)
}
func (m *PcepPlspStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcepPlspStatsBag.Marshal(b, m, deterministic)
}
func (m *PcepPlspStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcepPlspStatsBag.Merge(m, src)
}
func (m *PcepPlspStatsBag) XXX_Size() int {
	return xxx_messageInfo_PcepPlspStatsBag.Size(m)
}
func (m *PcepPlspStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcepPlspStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcepPlspStatsBag proto.InternalMessageInfo

func (m *PcepPlspStatsBag) GetPathsCreated() uint64 {
	if m != nil {
		return m.PathsCreated
	}
	return 0
}

func (m *PcepPlspStatsBag) GetPathsDestroyed() uint64 {
	if m != nil {
		return m.PathsDestroyed
	}
	return 0
}

func (m *PcepPlspStatsBag) GetPathCreateErrors() uint64 {
	if m != nil {
		return m.PathCreateErrors
	}
	return 0
}

func (m *PcepPlspStatsBag) GetPathDestroyErrors() uint64 {
	if m != nil {
		return m.PathDestroyErrors
	}
	return 0
}

func (m *PcepPlspStatsBag) GetRequestsCreated() uint64 {
	if m != nil {
		return m.RequestsCreated
	}
	return 0
}

func (m *PcepPlspStatsBag) GetRequestsDestroyed() uint64 {
	if m != nil {
		return m.RequestsDestroyed
	}
	return 0
}

func (m *PcepPlspStatsBag) GetRequestsFailed() uint64 {
	if m != nil {
		return m.RequestsFailed
	}
	return 0
}

type PcepPlspEventBag struct {
	Ts                   uint64   `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcepPlspEventBag) Reset()         { *m = PcepPlspEventBag{} }
func (m *PcepPlspEventBag) String() string { return proto.CompactTextString(m) }
func (*PcepPlspEventBag) ProtoMessage()    {}
func (*PcepPlspEventBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fedae5cf2bc176a, []int{2}
}

func (m *PcepPlspEventBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcepPlspEventBag.Unmarshal(m, b)
}
func (m *PcepPlspEventBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcepPlspEventBag.Marshal(b, m, deterministic)
}
func (m *PcepPlspEventBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcepPlspEventBag.Merge(m, src)
}
func (m *PcepPlspEventBag) XXX_Size() int {
	return xxx_messageInfo_PcepPlspEventBag.Size(m)
}
func (m *PcepPlspEventBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcepPlspEventBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcepPlspEventBag proto.InternalMessageInfo

func (m *PcepPlspEventBag) GetTs() uint64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *PcepPlspEventBag) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type PcepPathStatsBag struct {
	ReportsRequested     uint64   `protobuf:"varint,1,opt,name=reports_requested,json=reportsRequested,proto3" json:"reports_requested,omitempty"`
	ReportsSent          uint64   `protobuf:"varint,2,opt,name=reports_sent,json=reportsSent,proto3" json:"reports_sent,omitempty"`
	ReportsFailedToSend  uint64   `protobuf:"varint,3,opt,name=reports_failed_to_send,json=reportsFailedToSend,proto3" json:"reports_failed_to_send,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcepPathStatsBag) Reset()         { *m = PcepPathStatsBag{} }
func (m *PcepPathStatsBag) String() string { return proto.CompactTextString(m) }
func (*PcepPathStatsBag) ProtoMessage()    {}
func (*PcepPathStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fedae5cf2bc176a, []int{3}
}

func (m *PcepPathStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcepPathStatsBag.Unmarshal(m, b)
}
func (m *PcepPathStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcepPathStatsBag.Marshal(b, m, deterministic)
}
func (m *PcepPathStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcepPathStatsBag.Merge(m, src)
}
func (m *PcepPathStatsBag) XXX_Size() int {
	return xxx_messageInfo_PcepPathStatsBag.Size(m)
}
func (m *PcepPathStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcepPathStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcepPathStatsBag proto.InternalMessageInfo

func (m *PcepPathStatsBag) GetReportsRequested() uint64 {
	if m != nil {
		return m.ReportsRequested
	}
	return 0
}

func (m *PcepPathStatsBag) GetReportsSent() uint64 {
	if m != nil {
		return m.ReportsSent
	}
	return 0
}

func (m *PcepPathStatsBag) GetReportsFailedToSend() uint64 {
	if m != nil {
		return m.ReportsFailedToSend
	}
	return 0
}

type PcepHopIpv4 struct {
	V4Addr               uint32   `protobuf:"varint,1,opt,name=v4_addr,json=v4Addr,proto3" json:"v4_addr,omitempty"`
	PrefixLen            uint32   `protobuf:"varint,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcepHopIpv4) Reset()         { *m = PcepHopIpv4{} }
func (m *PcepHopIpv4) String() string { return proto.CompactTextString(m) }
func (*PcepHopIpv4) ProtoMessage()    {}
func (*PcepHopIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fedae5cf2bc176a, []int{4}
}

func (m *PcepHopIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcepHopIpv4.Unmarshal(m, b)
}
func (m *PcepHopIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcepHopIpv4.Marshal(b, m, deterministic)
}
func (m *PcepHopIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcepHopIpv4.Merge(m, src)
}
func (m *PcepHopIpv4) XXX_Size() int {
	return xxx_messageInfo_PcepHopIpv4.Size(m)
}
func (m *PcepHopIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_PcepHopIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_PcepHopIpv4 proto.InternalMessageInfo

func (m *PcepHopIpv4) GetV4Addr() uint32 {
	if m != nil {
		return m.V4Addr
	}
	return 0
}

func (m *PcepHopIpv4) GetPrefixLen() uint32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

type PcepHopSrIpv4 struct {
	Type                 uint32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Cflag                bool     `protobuf:"varint,2,opt,name=cflag,proto3" json:"cflag,omitempty"`
	Sid                  uint32   `protobuf:"varint,3,opt,name=sid,proto3" json:"sid,omitempty"`
	RemoteAddr           uint32   `protobuf:"varint,4,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	LocalAddr            uint32   `protobuf:"varint,5,opt,name=local_addr,json=localAddr,proto3" json:"local_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcepHopSrIpv4) Reset()         { *m = PcepHopSrIpv4{} }
func (m *PcepHopSrIpv4) String() string { return proto.CompactTextString(m) }
func (*PcepHopSrIpv4) ProtoMessage()    {}
func (*PcepHopSrIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fedae5cf2bc176a, []int{5}
}

func (m *PcepHopSrIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcepHopSrIpv4.Unmarshal(m, b)
}
func (m *PcepHopSrIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcepHopSrIpv4.Marshal(b, m, deterministic)
}
func (m *PcepHopSrIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcepHopSrIpv4.Merge(m, src)
}
func (m *PcepHopSrIpv4) XXX_Size() int {
	return xxx_messageInfo_PcepHopSrIpv4.Size(m)
}
func (m *PcepHopSrIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_PcepHopSrIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_PcepHopSrIpv4 proto.InternalMessageInfo

func (m *PcepHopSrIpv4) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *PcepHopSrIpv4) GetCflag() bool {
	if m != nil {
		return m.Cflag
	}
	return false
}

func (m *PcepHopSrIpv4) GetSid() uint32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

func (m *PcepHopSrIpv4) GetRemoteAddr() uint32 {
	if m != nil {
		return m.RemoteAddr
	}
	return 0
}

func (m *PcepHopSrIpv4) GetLocalAddr() uint32 {
	if m != nil {
		return m.LocalAddr
	}
	return 0
}

type PcepHopData struct {
	HopType              uint32         `protobuf:"varint,1,opt,name=hop_type,json=hopType,proto3" json:"hop_type,omitempty"`
	Ipv4                 *PcepHopIpv4   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	SrV4                 *PcepHopSrIpv4 `protobuf:"bytes,3,opt,name=sr_v4,json=srV4,proto3" json:"sr_v4,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PcepHopData) Reset()         { *m = PcepHopData{} }
func (m *PcepHopData) String() string { return proto.CompactTextString(m) }
func (*PcepHopData) ProtoMessage()    {}
func (*PcepHopData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fedae5cf2bc176a, []int{6}
}

func (m *PcepHopData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcepHopData.Unmarshal(m, b)
}
func (m *PcepHopData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcepHopData.Marshal(b, m, deterministic)
}
func (m *PcepHopData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcepHopData.Merge(m, src)
}
func (m *PcepHopData) XXX_Size() int {
	return xxx_messageInfo_PcepHopData.Size(m)
}
func (m *PcepHopData) XXX_DiscardUnknown() {
	xxx_messageInfo_PcepHopData.DiscardUnknown(m)
}

var xxx_messageInfo_PcepHopData proto.InternalMessageInfo

func (m *PcepHopData) GetHopType() uint32 {
	if m != nil {
		return m.HopType
	}
	return 0
}

func (m *PcepHopData) GetIpv4() *PcepHopIpv4 {
	if m != nil {
		return m.Ipv4
	}
	return nil
}

func (m *PcepHopData) GetSrV4() *PcepHopSrIpv4 {
	if m != nil {
		return m.SrV4
	}
	return nil
}

type PcepHopBag struct {
	Data                 *PcepHopData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Loose                bool         `protobuf:"varint,2,opt,name=loose,proto3" json:"loose,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PcepHopBag) Reset()         { *m = PcepHopBag{} }
func (m *PcepHopBag) String() string { return proto.CompactTextString(m) }
func (*PcepHopBag) ProtoMessage()    {}
func (*PcepHopBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fedae5cf2bc176a, []int{7}
}

func (m *PcepHopBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcepHopBag.Unmarshal(m, b)
}
func (m *PcepHopBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcepHopBag.Marshal(b, m, deterministic)
}
func (m *PcepHopBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcepHopBag.Merge(m, src)
}
func (m *PcepHopBag) XXX_Size() int {
	return xxx_messageInfo_PcepHopBag.Size(m)
}
func (m *PcepHopBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcepHopBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcepHopBag proto.InternalMessageInfo

func (m *PcepHopBag) GetData() *PcepHopData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PcepHopBag) GetLoose() bool {
	if m != nil {
		return m.Loose
	}
	return false
}

type PcepPathBag struct {
	Stats                         *PcepPathStatsBag `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	EroHop                        []*PcepHopBag     `protobuf:"bytes,2,rep,name=ero_hop,json=eroHop,proto3" json:"ero_hop,omitempty"`
	RroHop                        []*PcepHopBag     `protobuf:"bytes,3,rep,name=rro_hop,json=rroHop,proto3" json:"rro_hop,omitempty"`
	UsedBw                        string            `protobuf:"bytes,4,opt,name=used_bw,json=usedBw,proto3" json:"used_bw,omitempty"`
	RequestedBw                   string            `protobuf:"bytes,5,opt,name=requested_bw,json=requestedBw,proto3" json:"requested_bw,omitempty"`
	MetricValue                   string            `protobuf:"bytes,6,opt,name=metric_value,json=metricValue,proto3" json:"metric_value,omitempty"`
	Refcnt                        string            `protobuf:"bytes,7,opt,name=refcnt,proto3" json:"refcnt,omitempty"`
	LspPlspId                     uint32            `protobuf:"varint,8,opt,name=lsp_plsp_id,json=lspPlspId,proto3" json:"lsp_plsp_id,omitempty"`
	BindingSidValue               uint32            `protobuf:"varint,9,opt,name=binding_sid_value,json=bindingSidValue,proto3" json:"binding_sid_value,omitempty"`
	LspIdTlvExtTunnelId           uint32            `protobuf:"varint,10,opt,name=lsp_id_tlv_ext_tunnel_id,json=lspIdTlvExtTunnelId,proto3" json:"lsp_id_tlv_ext_tunnel_id,omitempty"`
	LspIdTlvTunnelEndpointAddress uint32            `protobuf:"varint,11,opt,name=lsp_id_tlv_tunnel_endpoint_address,json=lspIdTlvTunnelEndpointAddress,proto3" json:"lsp_id_tlv_tunnel_endpoint_address,omitempty"`
	LspIdTlvTunnelSenderAddress   uint32            `protobuf:"varint,12,opt,name=lsp_id_tlv_tunnel_sender_address,json=lspIdTlvTunnelSenderAddress,proto3" json:"lsp_id_tlv_tunnel_sender_address,omitempty"`
	SrpId                         uint32            `protobuf:"varint,13,opt,name=srp_id,json=srpId,proto3" json:"srp_id,omitempty"`
	LspIdTlvLspId                 uint32            `protobuf:"varint,14,opt,name=lsp_id_tlv_lsp_id,json=lspIdTlvLspId,proto3" json:"lsp_id_tlv_lsp_id,omitempty"`
	LspIdTlvTunnelId              uint32            `protobuf:"varint,15,opt,name=lsp_id_tlv_tunnel_id,json=lspIdTlvTunnelId,proto3" json:"lsp_id_tlv_tunnel_id,omitempty"`
	LspId                         uint32            `protobuf:"varint,16,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	BindingSidType                uint32            `protobuf:"varint,17,opt,name=binding_sid_type,json=bindingSidType,proto3" json:"binding_sid_type,omitempty"`
	LspOper                       uint32            `protobuf:"varint,18,opt,name=lsp_oper,json=lspOper,proto3" json:"lsp_oper,omitempty"`
	PathSetupType                 uint32            `protobuf:"varint,19,opt,name=path_setup_type,json=pathSetupType,proto3" json:"path_setup_type,omitempty"`
	MetricType                    uint32            `protobuf:"varint,20,opt,name=metric_type,json=metricType,proto3" json:"metric_type,omitempty"`
	IsReported                    bool              `protobuf:"varint,21,opt,name=is_reported,json=isReported,proto3" json:"is_reported,omitempty"`
	LspAFlag                      bool              `protobuf:"varint,22,opt,name=lsp_a_flag,json=lspAFlag,proto3" json:"lsp_a_flag,omitempty"`
	LspRFlag                      bool              `protobuf:"varint,23,opt,name=lsp_r_flag,json=lspRFlag,proto3" json:"lsp_r_flag,omitempty"`
	LspSFlag                      bool              `protobuf:"varint,24,opt,name=lsp_s_flag,json=lspSFlag,proto3" json:"lsp_s_flag,omitempty"`
	LspDFlag                      bool              `protobuf:"varint,25,opt,name=lsp_d_flag,json=lspDFlag,proto3" json:"lsp_d_flag,omitempty"`
	LspCFlag                      bool              `protobuf:"varint,26,opt,name=lsp_c_flag,json=lspCFlag,proto3" json:"lsp_c_flag,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}          `json:"-"`
	XXX_unrecognized              []byte            `json:"-"`
	XXX_sizecache                 int32             `json:"-"`
}

func (m *PcepPathBag) Reset()         { *m = PcepPathBag{} }
func (m *PcepPathBag) String() string { return proto.CompactTextString(m) }
func (*PcepPathBag) ProtoMessage()    {}
func (*PcepPathBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fedae5cf2bc176a, []int{8}
}

func (m *PcepPathBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcepPathBag.Unmarshal(m, b)
}
func (m *PcepPathBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcepPathBag.Marshal(b, m, deterministic)
}
func (m *PcepPathBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcepPathBag.Merge(m, src)
}
func (m *PcepPathBag) XXX_Size() int {
	return xxx_messageInfo_PcepPathBag.Size(m)
}
func (m *PcepPathBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcepPathBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcepPathBag proto.InternalMessageInfo

func (m *PcepPathBag) GetStats() *PcepPathStatsBag {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *PcepPathBag) GetEroHop() []*PcepHopBag {
	if m != nil {
		return m.EroHop
	}
	return nil
}

func (m *PcepPathBag) GetRroHop() []*PcepHopBag {
	if m != nil {
		return m.RroHop
	}
	return nil
}

func (m *PcepPathBag) GetUsedBw() string {
	if m != nil {
		return m.UsedBw
	}
	return ""
}

func (m *PcepPathBag) GetRequestedBw() string {
	if m != nil {
		return m.RequestedBw
	}
	return ""
}

func (m *PcepPathBag) GetMetricValue() string {
	if m != nil {
		return m.MetricValue
	}
	return ""
}

func (m *PcepPathBag) GetRefcnt() string {
	if m != nil {
		return m.Refcnt
	}
	return ""
}

func (m *PcepPathBag) GetLspPlspId() uint32 {
	if m != nil {
		return m.LspPlspId
	}
	return 0
}

func (m *PcepPathBag) GetBindingSidValue() uint32 {
	if m != nil {
		return m.BindingSidValue
	}
	return 0
}

func (m *PcepPathBag) GetLspIdTlvExtTunnelId() uint32 {
	if m != nil {
		return m.LspIdTlvExtTunnelId
	}
	return 0
}

func (m *PcepPathBag) GetLspIdTlvTunnelEndpointAddress() uint32 {
	if m != nil {
		return m.LspIdTlvTunnelEndpointAddress
	}
	return 0
}

func (m *PcepPathBag) GetLspIdTlvTunnelSenderAddress() uint32 {
	if m != nil {
		return m.LspIdTlvTunnelSenderAddress
	}
	return 0
}

func (m *PcepPathBag) GetSrpId() uint32 {
	if m != nil {
		return m.SrpId
	}
	return 0
}

func (m *PcepPathBag) GetLspIdTlvLspId() uint32 {
	if m != nil {
		return m.LspIdTlvLspId
	}
	return 0
}

func (m *PcepPathBag) GetLspIdTlvTunnelId() uint32 {
	if m != nil {
		return m.LspIdTlvTunnelId
	}
	return 0
}

func (m *PcepPathBag) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *PcepPathBag) GetBindingSidType() uint32 {
	if m != nil {
		return m.BindingSidType
	}
	return 0
}

func (m *PcepPathBag) GetLspOper() uint32 {
	if m != nil {
		return m.LspOper
	}
	return 0
}

func (m *PcepPathBag) GetPathSetupType() uint32 {
	if m != nil {
		return m.PathSetupType
	}
	return 0
}

func (m *PcepPathBag) GetMetricType() uint32 {
	if m != nil {
		return m.MetricType
	}
	return 0
}

func (m *PcepPathBag) GetIsReported() bool {
	if m != nil {
		return m.IsReported
	}
	return false
}

func (m *PcepPathBag) GetLspAFlag() bool {
	if m != nil {
		return m.LspAFlag
	}
	return false
}

func (m *PcepPathBag) GetLspRFlag() bool {
	if m != nil {
		return m.LspRFlag
	}
	return false
}

func (m *PcepPathBag) GetLspSFlag() bool {
	if m != nil {
		return m.LspSFlag
	}
	return false
}

func (m *PcepPathBag) GetLspDFlag() bool {
	if m != nil {
		return m.LspDFlag
	}
	return false
}

func (m *PcepPathBag) GetLspCFlag() bool {
	if m != nil {
		return m.LspCFlag
	}
	return false
}

type PcepPlspBag struct {
	PlspIdXr             uint32              `protobuf:"varint,50,opt,name=plsp_id_xr,json=plspIdXr,proto3" json:"plsp_id_xr,omitempty"`
	SymPathName          string              `protobuf:"bytes,51,opt,name=sym_path_name,json=symPathName,proto3" json:"sym_path_name,omitempty"`
	Stats                *PcepPlspStatsBag   `protobuf:"bytes,52,opt,name=stats,proto3" json:"stats,omitempty"`
	Refcnt               string              `protobuf:"bytes,53,opt,name=refcnt,proto3" json:"refcnt,omitempty"`
	ConnDelegatedTo      uint32              `protobuf:"varint,54,opt,name=conn_delegated_to,json=connDelegatedTo,proto3" json:"conn_delegated_to,omitempty"`
	EventHistory         []*PcepPlspEventBag `protobuf:"bytes,55,rep,name=event_history,json=eventHistory,proto3" json:"event_history,omitempty"`
	Path                 []*PcepPathBag      `protobuf:"bytes,56,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PcepPlspBag) Reset()         { *m = PcepPlspBag{} }
func (m *PcepPlspBag) String() string { return proto.CompactTextString(m) }
func (*PcepPlspBag) ProtoMessage()    {}
func (*PcepPlspBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fedae5cf2bc176a, []int{9}
}

func (m *PcepPlspBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcepPlspBag.Unmarshal(m, b)
}
func (m *PcepPlspBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcepPlspBag.Marshal(b, m, deterministic)
}
func (m *PcepPlspBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcepPlspBag.Merge(m, src)
}
func (m *PcepPlspBag) XXX_Size() int {
	return xxx_messageInfo_PcepPlspBag.Size(m)
}
func (m *PcepPlspBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcepPlspBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcepPlspBag proto.InternalMessageInfo

func (m *PcepPlspBag) GetPlspIdXr() uint32 {
	if m != nil {
		return m.PlspIdXr
	}
	return 0
}

func (m *PcepPlspBag) GetSymPathName() string {
	if m != nil {
		return m.SymPathName
	}
	return ""
}

func (m *PcepPlspBag) GetStats() *PcepPlspStatsBag {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *PcepPlspBag) GetRefcnt() string {
	if m != nil {
		return m.Refcnt
	}
	return ""
}

func (m *PcepPlspBag) GetConnDelegatedTo() uint32 {
	if m != nil {
		return m.ConnDelegatedTo
	}
	return 0
}

func (m *PcepPlspBag) GetEventHistory() []*PcepPlspEventBag {
	if m != nil {
		return m.EventHistory
	}
	return nil
}

func (m *PcepPlspBag) GetPath() []*PcepPathBag {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*PcepPlspBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_agent_oper.pcc.plsps.plsp.pcep_plsp_bag_KEYS")
	proto.RegisterType((*PcepPlspStatsBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.pcc.plsps.plsp.pcep_plsp_stats_bag")
	proto.RegisterType((*PcepPlspEventBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.pcc.plsps.plsp.pcep_plsp_event_bag")
	proto.RegisterType((*PcepPathStatsBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.pcc.plsps.plsp.pcep_path_stats_bag")
	proto.RegisterType((*PcepHopIpv4)(nil), "cisco_ios_xr_infra_xtc_agent_oper.pcc.plsps.plsp.pcep_hop_ipv4")
	proto.RegisterType((*PcepHopSrIpv4)(nil), "cisco_ios_xr_infra_xtc_agent_oper.pcc.plsps.plsp.pcep_hop_sr_ipv4")
	proto.RegisterType((*PcepHopData)(nil), "cisco_ios_xr_infra_xtc_agent_oper.pcc.plsps.plsp.pcep_hop_data")
	proto.RegisterType((*PcepHopBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.pcc.plsps.plsp.pcep_hop_bag")
	proto.RegisterType((*PcepPathBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.pcc.plsps.plsp.pcep_path_bag")
	proto.RegisterType((*PcepPlspBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.pcc.plsps.plsp.pcep_plsp_bag")
}

func init() { proto.RegisterFile("pcep_plsp_bag.proto", fileDescriptor_4fedae5cf2bc176a) }

var fileDescriptor_4fedae5cf2bc176a = []byte{
	// 1130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x72, 0x1b, 0xb5,
	0x17, 0x9e, 0x24, 0xb6, 0x63, 0x1f, 0xc7, 0x89, 0xad, 0xb4, 0xc9, 0xe6, 0xf7, 0xa3, 0x34, 0x2c,
	0x33, 0x34, 0x14, 0xea, 0x61, 0x12, 0x97, 0x3f, 0x37, 0x74, 0xd2, 0xc6, 0xa5, 0x19, 0x3a, 0xd0,
	0x59, 0x67, 0x5a, 0x18, 0x2e, 0x34, 0x9b, 0x95, 0x92, 0x2c, 0xb3, 0x5e, 0x2d, 0x92, 0xe2, 0xda,
	0x2f, 0xc1, 0x03, 0xf0, 0x1a, 0xbc, 0x02, 0x0f, 0xc1, 0x9b, 0x70, 0xcb, 0x9c, 0x23, 0xed, 0xda,
	0x9e, 0x5e, 0x81, 0xb9, 0xf1, 0x78, 0xbf, 0xef, 0xd3, 0xa7, 0xa3, 0x23, 0x9d, 0x23, 0xc1, 0x6e,
	0x91, 0xc8, 0x82, 0x17, 0x99, 0x29, 0xf8, 0x65, 0x7c, 0xdd, 0x2f, 0xb4, 0xb2, 0x8a, 0x7d, 0x96,
	0xa4, 0x26, 0x51, 0x3c, 0x55, 0x86, 0x4f, 0x35, 0x4f, 0xf3, 0x2b, 0x1d, 0xf3, 0xa9, 0x4d, 0x78,
	0x7c, 0x2d, 0x73, 0xcb, 0x55, 0x21, 0x75, 0xbf, 0x48, 0x92, 0x3e, 0x8e, 0x32, 0xf4, 0x1b, 0x3e,
	0x02, 0xb6, 0x64, 0xc4, 0xbf, 0x1d, 0xfe, 0x38, 0x62, 0xfb, 0xb0, 0x49, 0x40, 0x2a, 0x82, 0xb5,
	0xc3, 0xb5, 0xa3, 0x4e, 0xd4, 0xc0, 0xcf, 0x73, 0x11, 0xfe, 0xb1, 0xbe, 0x38, 0xb1, 0xb1, 0xb1,
	0x35, 0x38, 0x8a, 0x7d, 0x08, 0x9d, 0x22, 0xb6, 0x37, 0x86, 0x27, 0x5a, 0xc6, 0x56, 0xba, 0x61,
	0xb5, 0x68, 0x8b, 0xc0, 0x67, 0x0e, 0x63, 0x0f, 0x60, 0xc7, 0x89, 0x84, 0x34, 0x56, 0xab, 0x99,
	0x14, 0xc1, 0x3a, 0xc9, 0xb6, 0x09, 0x3e, 0x2b, 0x51, 0xf6, 0x29, 0x30, 0x44, 0xbc, 0x19, 0x97,
	0x5a, 0x2b, 0x6d, 0x82, 0x0d, 0xd2, 0x76, 0x91, 0x71, 0x8e, 0x43, 0xc2, 0x59, 0x1f, 0x76, 0x49,
	0xed, 0x5d, 0x4b, 0x79, 0x8d, 0xe4, 0x3d, 0xa4, 0xbc, 0xb3, 0xd7, 0x7f, 0x0c, 0x5d, 0x2d, 0x7f,
	0xb9, 0x95, 0xc6, 0xce, 0xc3, 0xad, 0x93, 0x78, 0xa7, 0xc4, 0xcb, 0x88, 0x1f, 0x01, 0xab, 0xa4,
	0xf3, 0xa0, 0x1b, 0xce, 0xb9, 0x64, 0xe6, 0x71, 0x3f, 0x80, 0xca, 0x81, 0x5f, 0xc5, 0x69, 0x26,
	0x45, 0xb0, 0xe9, 0x16, 0x58, 0xc2, 0xcf, 0x09, 0x0d, 0xbf, 0x5a, 0xcc, 0xa2, 0x9c, 0xe0, 0xe6,
	0x60, 0x16, 0xb7, 0x61, 0xdd, 0x1a, 0x9f, 0xba, 0x75, 0x6b, 0x18, 0x83, 0x9a, 0x90, 0x26, 0xa1,
	0x2c, 0xb5, 0x22, 0xfa, 0x1f, 0xfe, 0xb6, 0x56, 0x8e, 0xc5, 0x35, 0xcf, 0x77, 0xe0, 0x13, 0xe8,
	0x69, 0x59, 0x28, 0x6d, 0x0d, 0xf7, 0x93, 0x55, 0xbb, 0xd0, 0xf5, 0x44, 0x54, 0xe2, 0xec, 0x03,
	0xd8, 0x2a, 0xc5, 0x46, 0xe6, 0xd6, 0x6f, 0x43, 0xdb, 0x63, 0x23, 0x99, 0x5b, 0x76, 0x02, 0x7b,
	0xa5, 0xc4, 0x2d, 0x85, 0x5b, 0x85, 0x62, 0xe1, 0xf7, 0x61, 0xd7, 0xb3, 0x6e, 0x45, 0x17, 0x6a,
	0x24, 0x73, 0x11, 0x7e, 0x03, 0x1d, 0x8a, 0xed, 0x46, 0x15, 0x3c, 0x2d, 0x26, 0x03, 0x3c, 0x48,
	0x93, 0x01, 0x8f, 0x85, 0xd0, 0xe5, 0x41, 0x9a, 0x0c, 0x4e, 0x85, 0xd0, 0xec, 0x1e, 0x40, 0xa1,
	0xe5, 0x55, 0x3a, 0xe5, 0x99, 0xcc, 0x69, 0xfe, 0x4e, 0xd4, 0x72, 0xc8, 0x4b, 0x99, 0x87, 0xbf,
	0xae, 0x41, 0xb7, 0x72, 0x32, 0xda, 0x99, 0x31, 0xa8, 0xd9, 0x59, 0x21, 0xbd, 0x13, 0xfd, 0x67,
	0x77, 0xa0, 0x9e, 0x5c, 0x65, 0xf1, 0x35, 0x59, 0x34, 0x23, 0xf7, 0xc1, 0xba, 0xb0, 0x61, 0x52,
	0x17, 0x69, 0x27, 0xc2, 0xbf, 0xec, 0x3e, 0xb4, 0xb5, 0x1c, 0x2b, 0x2b, 0x5d, 0x30, 0x35, 0x62,
	0xc0, 0x41, 0x65, 0x40, 0x99, 0x4a, 0xe2, 0xcc, 0xf1, 0x75, 0x17, 0x10, 0x21, 0x48, 0x87, 0x7f,
	0xae, 0x2d, 0x2c, 0x4d, 0xc4, 0x36, 0x66, 0x07, 0xd0, 0xc4, 0xff, 0x0b, 0x11, 0x6d, 0xde, 0xa8,
	0xe2, 0x02, 0x83, 0x1a, 0x41, 0x0d, 0x03, 0xa6, 0x98, 0xda, 0xc7, 0x4f, 0xfa, 0xff, 0xb4, 0x2a,
	0xfb, 0x4b, 0x49, 0x8c, 0xc8, 0x8c, 0xbd, 0x81, 0xba, 0xd1, 0x7c, 0x32, 0xa0, 0x55, 0xb5, 0x8f,
	0x9f, 0xae, 0xe0, 0xea, 0x13, 0x1a, 0xd5, 0x8c, 0x7e, 0x3d, 0x08, 0x67, 0xb0, 0x55, 0x31, 0x78,
	0x92, 0x46, 0x50, 0xc3, 0x05, 0xd2, 0xa2, 0x56, 0x8b, 0x1e, 0x6d, 0x22, 0x32, 0xc3, 0x7d, 0xca,
	0x94, 0x32, 0xb2, 0xdc, 0x27, 0xfa, 0x08, 0xff, 0x6a, 0xfa, 0xac, 0xd2, 0x61, 0xc6, 0xc9, 0x7f,
	0x82, 0x3a, 0x9d, 0x69, 0x3f, 0xfb, 0xf0, 0x5f, 0xce, 0xbe, 0x5c, 0x1c, 0x91, 0xf3, 0x64, 0x6f,
	0x60, 0x53, 0x6a, 0x85, 0xa1, 0x05, 0xeb, 0x87, 0x1b, 0x47, 0xed, 0xe3, 0xaf, 0x57, 0x58, 0x1c,
	0xfa, 0x36, 0xa4, 0x56, 0x2f, 0x54, 0x81, 0xc6, 0xda, 0x1b, 0x6f, 0xfc, 0x37, 0xc6, 0xda, 0x19,
	0xef, 0xc3, 0xe6, 0xad, 0x91, 0x82, 0x5f, 0xbe, 0xa5, 0x23, 0xdb, 0x8a, 0x1a, 0xf8, 0xf9, 0xf4,
	0xad, 0xab, 0x60, 0x5f, 0xce, 0xc8, 0xd6, 0x89, 0x6d, 0x57, 0x98, 0x93, 0x8c, 0xa5, 0xd5, 0x69,
	0xc2, 0x27, 0x71, 0x76, 0x2b, 0xa9, 0x6d, 0xb5, 0xa2, 0xb6, 0xc3, 0x5e, 0x23, 0xc4, 0xf6, 0xa0,
	0xa1, 0xe5, 0x55, 0x92, 0x5b, 0xea, 0x53, 0xad, 0xc8, 0x7f, 0xb1, 0xf7, 0xa1, 0x8d, 0x9d, 0xa9,
	0xbc, 0x03, 0x9a, 0xbe, 0x1a, 0x4c, 0xf1, 0x8a, 0xae, 0x01, 0xf6, 0x10, 0x7a, 0x97, 0x69, 0x2e,
	0xd2, 0xfc, 0x9a, 0x9b, 0x54, 0x78, 0xff, 0x16, 0xa9, 0x76, 0x3c, 0x31, 0x4a, 0x85, 0x9b, 0xe3,
	0x31, 0x04, 0xce, 0x86, 0xdb, 0x6c, 0xc2, 0xe5, 0xd4, 0x72, 0x7b, 0x9b, 0xe7, 0x32, 0x43, 0x63,
	0xa0, 0x21, 0xbb, 0x64, 0x7a, 0x91, 0x4d, 0x86, 0x53, 0x7b, 0x41, 0xdc, 0xb9, 0x60, 0xe7, 0x10,
	0x2e, 0x0c, 0xf3, 0x43, 0x64, 0x2e, 0x0a, 0x95, 0xe6, 0x96, 0x8a, 0x54, 0x1a, 0x13, 0xb4, 0xc9,
	0xe0, 0x5e, 0x69, 0xe0, 0x46, 0x0f, 0xbd, 0xea, 0xd4, 0x89, 0xd8, 0x10, 0x0e, 0xdf, 0xb5, 0xc2,
	0x56, 0x26, 0x75, 0x65, 0xb4, 0x45, 0x46, 0xff, 0x5f, 0x36, 0x1a, 0x91, 0xa6, 0xb4, 0xb9, 0x0b,
	0x0d, 0xa3, 0x29, 0x1f, 0x1d, 0x12, 0xd7, 0x8d, 0xc6, 0x5c, 0x1c, 0x41, 0x6f, 0xc1, 0xdd, 0x67,
	0x6c, 0x9b, 0x14, 0x9d, 0xd2, 0xee, 0x25, 0x65, 0xad, 0x0f, 0x77, 0xde, 0x8d, 0x23, 0x15, 0xc1,
	0x0e, 0x89, 0xbb, 0xcb, 0x73, 0x9f, 0x0b, 0x9c, 0xd0, 0xdb, 0x75, 0xdd, 0x84, 0x2e, 0xf9, 0x47,
	0xd0, 0x5d, 0x4c, 0x3e, 0x35, 0xa0, 0x1e, 0x09, 0xb6, 0xe7, 0xb9, 0xa7, 0x3e, 0x74, 0x00, 0x4d,
	0x34, 0xc0, 0xd3, 0x16, 0x30, 0xd7, 0xa2, 0x32, 0x53, 0x7c, 0x5f, 0x48, 0xcd, 0x3e, 0x72, 0x77,
	0x31, 0x37, 0xd2, 0xde, 0xfa, 0x26, 0xb6, 0xeb, 0x62, 0x46, 0x78, 0x84, 0x28, 0x59, 0xdc, 0x07,
	0x7f, 0x60, 0x9c, 0xe6, 0x8e, 0xeb, 0x9b, 0x0e, 0x2a, 0x05, 0x29, 0x5e, 0x39, 0x78, 0x19, 0x48,
	0x11, 0xdc, 0xa5, 0xf2, 0x86, 0xd4, 0x44, 0x1e, 0x61, 0xef, 0x01, 0x60, 0x10, 0x31, 0xa7, 0x36,
	0xbd, 0x47, 0x3c, 0x86, 0x75, 0xfa, 0x1c, 0x3b, 0xb5, 0x67, 0xb5, 0x63, 0xf7, 0x2b, 0x36, 0x5a,
	0x64, 0x8d, 0x63, 0x83, 0x8a, 0x1d, 0x2d, 0xb2, 0xc2, 0xb1, 0x07, 0x15, 0x7b, 0xb6, 0xc8, 0x26,
	0x8e, 0xfd, 0x5f, 0xc5, 0x3e, 0x43, 0x36, 0xfc, 0x7d, 0xa3, 0xec, 0x3c, 0xfe, 0xe1, 0x83, 0x7a,
	0x7f, 0xde, 0xf9, 0x54, 0x07, 0xc7, 0xb4, 0xd0, 0xa6, 0x7b, 0xf6, 0xfc, 0xa0, 0x59, 0x08, 0x1d,
	0x33, 0x1b, 0xbb, 0xbe, 0x92, 0xc7, 0x63, 0x19, 0x9c, 0xb8, 0x6a, 0x32, 0xb3, 0xf1, 0xab, 0xd8,
	0xde, 0x7c, 0x17, 0x8f, 0xe5, 0xbc, 0x77, 0x0d, 0x56, 0xeb, 0x5d, 0x4b, 0x4f, 0xab, 0xb2, 0x77,
	0xcd, 0x4b, 0xf5, 0xf1, 0x52, 0xa9, 0x3e, 0x84, 0x5e, 0xa2, 0xf2, 0x9c, 0x0b, 0x99, 0xc9, 0x6b,
	0x7c, 0xb4, 0x70, 0xab, 0x82, 0xcf, 0x5d, 0x29, 0x22, 0x71, 0x56, 0xe2, 0x17, 0x8a, 0xfd, 0x0c,
	0x1d, 0xf7, 0xd8, 0xb8, 0x49, 0x8d, 0x55, 0x7a, 0x16, 0x7c, 0x41, 0xcd, 0x6a, 0xa5, 0x40, 0xab,
	0xd7, 0x4b, 0xb4, 0x45, 0x7f, 0x5f, 0x38, 0x6b, 0xbc, 0x45, 0x30, 0x59, 0xc1, 0x97, 0x34, 0xc5,
	0x93, 0x55, 0xfa, 0x38, 0x9a, 0x93, 0xd9, 0x65, 0x83, 0x9e, 0xb9, 0x27, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x24, 0x35, 0x49, 0xe6, 0xfd, 0x0a, 0x00, 0x00,
}
