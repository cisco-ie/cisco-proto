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
// source: ospf_sh_stats_spf.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_statistics_spf_stats

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

type OspfShStatsSpf_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShStatsSpf_KEYS) Reset()         { *m = OspfShStatsSpf_KEYS{} }
func (m *OspfShStatsSpf_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShStatsSpf_KEYS) ProtoMessage()    {}
func (*OspfShStatsSpf_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{0}
}

func (m *OspfShStatsSpf_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShStatsSpf_KEYS.Unmarshal(m, b)
}
func (m *OspfShStatsSpf_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShStatsSpf_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShStatsSpf_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShStatsSpf_KEYS.Merge(m, src)
}
func (m *OspfShStatsSpf_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShStatsSpf_KEYS.Size(m)
}
func (m *OspfShStatsSpf_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShStatsSpf_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShStatsSpf_KEYS proto.InternalMessageInfo

func (m *OspfShStatsSpf_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShStatsSpf_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type OspfShTime struct {
	Second               uint32   `protobuf:"varint,1,opt,name=second,proto3" json:"second,omitempty"`
	Nanosecond           uint32   `protobuf:"varint,2,opt,name=nanosecond,proto3" json:"nanosecond,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShTime) Reset()         { *m = OspfShTime{} }
func (m *OspfShTime) String() string { return proto.CompactTextString(m) }
func (*OspfShTime) ProtoMessage()    {}
func (*OspfShTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{1}
}

func (m *OspfShTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTime.Unmarshal(m, b)
}
func (m *OspfShTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTime.Marshal(b, m, deterministic)
}
func (m *OspfShTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTime.Merge(m, src)
}
func (m *OspfShTime) XXX_Size() int {
	return xxx_messageInfo_OspfShTime.Size(m)
}
func (m *OspfShTime) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTime.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTime proto.InternalMessageInfo

func (m *OspfShTime) GetSecond() uint32 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *OspfShTime) GetNanosecond() uint32 {
	if m != nil {
		return m.Nanosecond
	}
	return 0
}

type OspfShStatIntra struct {
	SiAreaIdStr          string      `protobuf:"bytes,1,opt,name=si_area_id_str,json=siAreaIdStr,proto3" json:"si_area_id_str,omitempty"`
	SiStartTime          *OspfShTime `protobuf:"bytes,2,opt,name=si_start_time,json=siStartTime,proto3" json:"si_start_time,omitempty"`
	SiChangeFlags        uint32      `protobuf:"varint,3,opt,name=si_change_flags,json=siChangeFlags,proto3" json:"si_change_flags,omitempty"`
	SiDuration           *OspfShTime `protobuf:"bytes,4,opt,name=si_duration,json=siDuration,proto3" json:"si_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OspfShStatIntra) Reset()         { *m = OspfShStatIntra{} }
func (m *OspfShStatIntra) String() string { return proto.CompactTextString(m) }
func (*OspfShStatIntra) ProtoMessage()    {}
func (*OspfShStatIntra) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{2}
}

func (m *OspfShStatIntra) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShStatIntra.Unmarshal(m, b)
}
func (m *OspfShStatIntra) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShStatIntra.Marshal(b, m, deterministic)
}
func (m *OspfShStatIntra) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShStatIntra.Merge(m, src)
}
func (m *OspfShStatIntra) XXX_Size() int {
	return xxx_messageInfo_OspfShStatIntra.Size(m)
}
func (m *OspfShStatIntra) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShStatIntra.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShStatIntra proto.InternalMessageInfo

func (m *OspfShStatIntra) GetSiAreaIdStr() string {
	if m != nil {
		return m.SiAreaIdStr
	}
	return ""
}

func (m *OspfShStatIntra) GetSiStartTime() *OspfShTime {
	if m != nil {
		return m.SiStartTime
	}
	return nil
}

func (m *OspfShStatIntra) GetSiChangeFlags() uint32 {
	if m != nil {
		return m.SiChangeFlags
	}
	return 0
}

func (m *OspfShStatIntra) GetSiDuration() *OspfShTime {
	if m != nil {
		return m.SiDuration
	}
	return nil
}

type OspfShStatPartial struct {
	SpAdvRtrId           uint32      `protobuf:"varint,1,opt,name=sp_adv_rtr_id,json=spAdvRtrId,proto3" json:"sp_adv_rtr_id,omitempty"`
	SpStartTime          *OspfShTime `protobuf:"bytes,2,opt,name=sp_start_time,json=spStartTime,proto3" json:"sp_start_time,omitempty"`
	SpDestCount          uint32      `protobuf:"varint,3,opt,name=sp_dest_count,json=spDestCount,proto3" json:"sp_dest_count,omitempty"`
	SpDestAddr           uint32      `protobuf:"varint,4,opt,name=sp_dest_addr,json=spDestAddr,proto3" json:"sp_dest_addr,omitempty"`
	SpDuration           *OspfShTime `protobuf:"bytes,5,opt,name=sp_duration,json=spDuration,proto3" json:"sp_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OspfShStatPartial) Reset()         { *m = OspfShStatPartial{} }
func (m *OspfShStatPartial) String() string { return proto.CompactTextString(m) }
func (*OspfShStatPartial) ProtoMessage()    {}
func (*OspfShStatPartial) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{3}
}

func (m *OspfShStatPartial) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShStatPartial.Unmarshal(m, b)
}
func (m *OspfShStatPartial) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShStatPartial.Marshal(b, m, deterministic)
}
func (m *OspfShStatPartial) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShStatPartial.Merge(m, src)
}
func (m *OspfShStatPartial) XXX_Size() int {
	return xxx_messageInfo_OspfShStatPartial.Size(m)
}
func (m *OspfShStatPartial) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShStatPartial.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShStatPartial proto.InternalMessageInfo

func (m *OspfShStatPartial) GetSpAdvRtrId() uint32 {
	if m != nil {
		return m.SpAdvRtrId
	}
	return 0
}

func (m *OspfShStatPartial) GetSpStartTime() *OspfShTime {
	if m != nil {
		return m.SpStartTime
	}
	return nil
}

func (m *OspfShStatPartial) GetSpDestCount() uint32 {
	if m != nil {
		return m.SpDestCount
	}
	return 0
}

func (m *OspfShStatPartial) GetSpDestAddr() uint32 {
	if m != nil {
		return m.SpDestAddr
	}
	return 0
}

func (m *OspfShStatPartial) GetSpDuration() *OspfShTime {
	if m != nil {
		return m.SpDuration
	}
	return nil
}

type OspfSpfTime struct {
	SpfDijkstra          uint32   `protobuf:"varint,1,opt,name=spf_dijkstra,json=spfDijkstra,proto3" json:"spf_dijkstra,omitempty"`
	SpfIntraPrefix       uint32   `protobuf:"varint,2,opt,name=spf_intra_prefix,json=spfIntraPrefix,proto3" json:"spf_intra_prefix,omitempty"`
	SpfIntraPrefixDel    uint32   `protobuf:"varint,3,opt,name=spf_intra_prefix_del,json=spfIntraPrefixDel,proto3" json:"spf_intra_prefix_del,omitempty"`
	SpfInterPrefix       uint32   `protobuf:"varint,4,opt,name=spf_inter_prefix,json=spfInterPrefix,proto3" json:"spf_inter_prefix,omitempty"`
	SpfInterPrefixDel    uint32   `protobuf:"varint,5,opt,name=spf_inter_prefix_del,json=spfInterPrefixDel,proto3" json:"spf_inter_prefix_del,omitempty"`
	SpfExtPrefix         uint32   `protobuf:"varint,6,opt,name=spf_ext_prefix,json=spfExtPrefix,proto3" json:"spf_ext_prefix,omitempty"`
	SpfExtPrefixDel      uint32   `protobuf:"varint,7,opt,name=spf_ext_prefix_del,json=spfExtPrefixDel,proto3" json:"spf_ext_prefix_del,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfSpfTime) Reset()         { *m = OspfSpfTime{} }
func (m *OspfSpfTime) String() string { return proto.CompactTextString(m) }
func (*OspfSpfTime) ProtoMessage()    {}
func (*OspfSpfTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{4}
}

func (m *OspfSpfTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfSpfTime.Unmarshal(m, b)
}
func (m *OspfSpfTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfSpfTime.Marshal(b, m, deterministic)
}
func (m *OspfSpfTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfSpfTime.Merge(m, src)
}
func (m *OspfSpfTime) XXX_Size() int {
	return xxx_messageInfo_OspfSpfTime.Size(m)
}
func (m *OspfSpfTime) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfSpfTime.DiscardUnknown(m)
}

var xxx_messageInfo_OspfSpfTime proto.InternalMessageInfo

func (m *OspfSpfTime) GetSpfDijkstra() uint32 {
	if m != nil {
		return m.SpfDijkstra
	}
	return 0
}

func (m *OspfSpfTime) GetSpfIntraPrefix() uint32 {
	if m != nil {
		return m.SpfIntraPrefix
	}
	return 0
}

func (m *OspfSpfTime) GetSpfIntraPrefixDel() uint32 {
	if m != nil {
		return m.SpfIntraPrefixDel
	}
	return 0
}

func (m *OspfSpfTime) GetSpfInterPrefix() uint32 {
	if m != nil {
		return m.SpfInterPrefix
	}
	return 0
}

func (m *OspfSpfTime) GetSpfInterPrefixDel() uint32 {
	if m != nil {
		return m.SpfInterPrefixDel
	}
	return 0
}

func (m *OspfSpfTime) GetSpfExtPrefix() uint32 {
	if m != nil {
		return m.SpfExtPrefix
	}
	return 0
}

func (m *OspfSpfTime) GetSpfExtPrefixDel() uint32 {
	if m != nil {
		return m.SpfExtPrefixDel
	}
	return 0
}

type OspfLsaChange struct {
	AreaId               uint32   `protobuf:"varint,1,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	AdvRtr               string   `protobuf:"bytes,2,opt,name=adv_rtr,json=advRtr,proto3" json:"adv_rtr,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Flush                bool     `protobuf:"varint,5,opt,name=flush,proto3" json:"flush,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfLsaChange) Reset()         { *m = OspfLsaChange{} }
func (m *OspfLsaChange) String() string { return proto.CompactTextString(m) }
func (*OspfLsaChange) ProtoMessage()    {}
func (*OspfLsaChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{5}
}

func (m *OspfLsaChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfLsaChange.Unmarshal(m, b)
}
func (m *OspfLsaChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfLsaChange.Marshal(b, m, deterministic)
}
func (m *OspfLsaChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfLsaChange.Merge(m, src)
}
func (m *OspfLsaChange) XXX_Size() int {
	return xxx_messageInfo_OspfLsaChange.Size(m)
}
func (m *OspfLsaChange) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfLsaChange.DiscardUnknown(m)
}

var xxx_messageInfo_OspfLsaChange proto.InternalMessageInfo

func (m *OspfLsaChange) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfLsaChange) GetAdvRtr() string {
	if m != nil {
		return m.AdvRtr
	}
	return ""
}

func (m *OspfLsaChange) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OspfLsaChange) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *OspfLsaChange) GetFlush() bool {
	if m != nil {
		return m.Flush
	}
	return false
}

type OspfEdmSpfAreaStats struct {
	AreaId               uint32       `protobuf:"varint,1,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	SpfTime              *OspfSpfTime `protobuf:"bytes,2,opt,name=spf_time,json=spfTime,proto3" json:"spf_time,omitempty"`
	LsaTypeCnt           []uint32     `protobuf:"varint,3,rep,packed,name=lsa_type_cnt,json=lsaTypeCnt,proto3" json:"lsa_type_cnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OspfEdmSpfAreaStats) Reset()         { *m = OspfEdmSpfAreaStats{} }
func (m *OspfEdmSpfAreaStats) String() string { return proto.CompactTextString(m) }
func (*OspfEdmSpfAreaStats) ProtoMessage()    {}
func (*OspfEdmSpfAreaStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{6}
}

func (m *OspfEdmSpfAreaStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfEdmSpfAreaStats.Unmarshal(m, b)
}
func (m *OspfEdmSpfAreaStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfEdmSpfAreaStats.Marshal(b, m, deterministic)
}
func (m *OspfEdmSpfAreaStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfEdmSpfAreaStats.Merge(m, src)
}
func (m *OspfEdmSpfAreaStats) XXX_Size() int {
	return xxx_messageInfo_OspfEdmSpfAreaStats.Size(m)
}
func (m *OspfEdmSpfAreaStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfEdmSpfAreaStats.DiscardUnknown(m)
}

var xxx_messageInfo_OspfEdmSpfAreaStats proto.InternalMessageInfo

func (m *OspfEdmSpfAreaStats) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfEdmSpfAreaStats) GetSpfTime() *OspfSpfTime {
	if m != nil {
		return m.SpfTime
	}
	return nil
}

func (m *OspfEdmSpfAreaStats) GetLsaTypeCnt() []uint32 {
	if m != nil {
		return m.LsaTypeCnt
	}
	return nil
}

type OspfRuntimeStats struct {
	SpfStartTime         uint32                 `protobuf:"varint,1,opt,name=spf_start_time,json=spfStartTime,proto3" json:"spf_start_time,omitempty"`
	GblSpfTime           *OspfSpfTime           `protobuf:"bytes,2,opt,name=gbl_spf_time,json=gblSpfTime,proto3" json:"gbl_spf_time,omitempty"`
	LsaChangeCnt         string                 `protobuf:"bytes,3,opt,name=lsa_change_cnt,json=lsaChangeCnt,proto3" json:"lsa_change_cnt,omitempty"`
	LsaInfo              []*OspfLsaChange       `protobuf:"bytes,4,rep,name=lsa_info,json=lsaInfo,proto3" json:"lsa_info,omitempty"`
	AreaStat             []*OspfEdmSpfAreaStats `protobuf:"bytes,5,rep,name=area_stat,json=areaStat,proto3" json:"area_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *OspfRuntimeStats) Reset()         { *m = OspfRuntimeStats{} }
func (m *OspfRuntimeStats) String() string { return proto.CompactTextString(m) }
func (*OspfRuntimeStats) ProtoMessage()    {}
func (*OspfRuntimeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{7}
}

func (m *OspfRuntimeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfRuntimeStats.Unmarshal(m, b)
}
func (m *OspfRuntimeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfRuntimeStats.Marshal(b, m, deterministic)
}
func (m *OspfRuntimeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfRuntimeStats.Merge(m, src)
}
func (m *OspfRuntimeStats) XXX_Size() int {
	return xxx_messageInfo_OspfRuntimeStats.Size(m)
}
func (m *OspfRuntimeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfRuntimeStats.DiscardUnknown(m)
}

var xxx_messageInfo_OspfRuntimeStats proto.InternalMessageInfo

func (m *OspfRuntimeStats) GetSpfStartTime() uint32 {
	if m != nil {
		return m.SpfStartTime
	}
	return 0
}

func (m *OspfRuntimeStats) GetGblSpfTime() *OspfSpfTime {
	if m != nil {
		return m.GblSpfTime
	}
	return nil
}

func (m *OspfRuntimeStats) GetLsaChangeCnt() string {
	if m != nil {
		return m.LsaChangeCnt
	}
	return ""
}

func (m *OspfRuntimeStats) GetLsaInfo() []*OspfLsaChange {
	if m != nil {
		return m.LsaInfo
	}
	return nil
}

func (m *OspfRuntimeStats) GetAreaStat() []*OspfEdmSpfAreaStats {
	if m != nil {
		return m.AreaStat
	}
	return nil
}

type OspfEdmAreaSumm struct {
	AreaId               uint32   `protobuf:"varint,1,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	SpfCount             uint32   `protobuf:"varint,2,opt,name=spf_count,json=spfCount,proto3" json:"spf_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfEdmAreaSumm) Reset()         { *m = OspfEdmAreaSumm{} }
func (m *OspfEdmAreaSumm) String() string { return proto.CompactTextString(m) }
func (*OspfEdmAreaSumm) ProtoMessage()    {}
func (*OspfEdmAreaSumm) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{8}
}

func (m *OspfEdmAreaSumm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfEdmAreaSumm.Unmarshal(m, b)
}
func (m *OspfEdmAreaSumm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfEdmAreaSumm.Marshal(b, m, deterministic)
}
func (m *OspfEdmAreaSumm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfEdmAreaSumm.Merge(m, src)
}
func (m *OspfEdmAreaSumm) XXX_Size() int {
	return xxx_messageInfo_OspfEdmAreaSumm.Size(m)
}
func (m *OspfEdmAreaSumm) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfEdmAreaSumm.DiscardUnknown(m)
}

var xxx_messageInfo_OspfEdmAreaSumm proto.InternalMessageInfo

func (m *OspfEdmAreaSumm) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfEdmAreaSumm) GetSpfCount() uint32 {
	if m != nil {
		return m.SpfCount
	}
	return 0
}

type OspfHeaderInfo struct {
	RouterId             string             `protobuf:"bytes,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	SpfCount             uint32             `protobuf:"varint,2,opt,name=spf_count,json=spfCount,proto3" json:"spf_count,omitempty"`
	AreaSumm             []*OspfEdmAreaSumm `protobuf:"bytes,3,rep,name=area_summ,json=areaSumm,proto3" json:"area_summ,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OspfHeaderInfo) Reset()         { *m = OspfHeaderInfo{} }
func (m *OspfHeaderInfo) String() string { return proto.CompactTextString(m) }
func (*OspfHeaderInfo) ProtoMessage()    {}
func (*OspfHeaderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{9}
}

func (m *OspfHeaderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfHeaderInfo.Unmarshal(m, b)
}
func (m *OspfHeaderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfHeaderInfo.Marshal(b, m, deterministic)
}
func (m *OspfHeaderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfHeaderInfo.Merge(m, src)
}
func (m *OspfHeaderInfo) XXX_Size() int {
	return xxx_messageInfo_OspfHeaderInfo.Size(m)
}
func (m *OspfHeaderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfHeaderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OspfHeaderInfo proto.InternalMessageInfo

func (m *OspfHeaderInfo) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *OspfHeaderInfo) GetSpfCount() uint32 {
	if m != nil {
		return m.SpfCount
	}
	return 0
}

func (m *OspfHeaderInfo) GetAreaSumm() []*OspfEdmAreaSumm {
	if m != nil {
		return m.AreaSumm
	}
	return nil
}

type OspfShStatsSpf struct {
	OsRuntime            []*OspfShStatIntra   `protobuf:"bytes,50,rep,name=os_runtime,json=osRuntime,proto3" json:"os_runtime,omitempty"`
	OsSumRuntime         []*OspfShStatPartial `protobuf:"bytes,51,rep,name=os_sum_runtime,json=osSumRuntime,proto3" json:"os_sum_runtime,omitempty"`
	OsExRuntime          []*OspfShStatPartial `protobuf:"bytes,52,rep,name=os_ex_runtime,json=osExRuntime,proto3" json:"os_ex_runtime,omitempty"`
	Runtime              []*OspfRuntimeStats  `protobuf:"bytes,53,rep,name=runtime,proto3" json:"runtime,omitempty"`
	SpfHeader            *OspfHeaderInfo      `protobuf:"bytes,54,opt,name=spf_header,json=spfHeader,proto3" json:"spf_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OspfShStatsSpf) Reset()         { *m = OspfShStatsSpf{} }
func (m *OspfShStatsSpf) String() string { return proto.CompactTextString(m) }
func (*OspfShStatsSpf) ProtoMessage()    {}
func (*OspfShStatsSpf) Descriptor() ([]byte, []int) {
	return fileDescriptor_885c7854df85c308, []int{10}
}

func (m *OspfShStatsSpf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShStatsSpf.Unmarshal(m, b)
}
func (m *OspfShStatsSpf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShStatsSpf.Marshal(b, m, deterministic)
}
func (m *OspfShStatsSpf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShStatsSpf.Merge(m, src)
}
func (m *OspfShStatsSpf) XXX_Size() int {
	return xxx_messageInfo_OspfShStatsSpf.Size(m)
}
func (m *OspfShStatsSpf) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShStatsSpf.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShStatsSpf proto.InternalMessageInfo

func (m *OspfShStatsSpf) GetOsRuntime() []*OspfShStatIntra {
	if m != nil {
		return m.OsRuntime
	}
	return nil
}

func (m *OspfShStatsSpf) GetOsSumRuntime() []*OspfShStatPartial {
	if m != nil {
		return m.OsSumRuntime
	}
	return nil
}

func (m *OspfShStatsSpf) GetOsExRuntime() []*OspfShStatPartial {
	if m != nil {
		return m.OsExRuntime
	}
	return nil
}

func (m *OspfShStatsSpf) GetRuntime() []*OspfRuntimeStats {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *OspfShStatsSpf) GetSpfHeader() *OspfHeaderInfo {
	if m != nil {
		return m.SpfHeader
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShStatsSpf_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_sh_stats_spf_KEYS")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_sh_time")
	proto.RegisterType((*OspfShStatIntra)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_sh_stat_intra")
	proto.RegisterType((*OspfShStatPartial)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_sh_stat_partial")
	proto.RegisterType((*OspfSpfTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_spf_time")
	proto.RegisterType((*OspfLsaChange)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_lsa_change")
	proto.RegisterType((*OspfEdmSpfAreaStats)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_edm_spf_area_stats")
	proto.RegisterType((*OspfRuntimeStats)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_runtime_stats")
	proto.RegisterType((*OspfEdmAreaSumm)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_edm_area_summ")
	proto.RegisterType((*OspfHeaderInfo)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_header_info")
	proto.RegisterType((*OspfShStatsSpf)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.spf_stats.ospf_sh_stats_spf")
}

func init() { proto.RegisterFile("ospf_sh_stats_spf.proto", fileDescriptor_885c7854df85c308) }

var fileDescriptor_885c7854df85c308 = []byte{
	// 937 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xdd, 0x6a, 0x24, 0x45,
	0x14, 0xa6, 0x27, 0x99, 0xcc, 0xcc, 0x99, 0x9f, 0xec, 0x16, 0x61, 0x77, 0x44, 0x90, 0xa4, 0x5d,
	0x24, 0x20, 0x8c, 0x90, 0x5d, 0xbd, 0x0f, 0x99, 0x2c, 0x46, 0x41, 0x96, 0x9e, 0x45, 0xf0, 0xc6,
	0xa2, 0x32, 0x55, 0x9d, 0x94, 0xf6, 0x74, 0x15, 0x75, 0x6a, 0x86, 0x59, 0xf1, 0x07, 0x15, 0xbd,
	0x10, 0x5f, 0x48, 0x7c, 0x09, 0xdf, 0xc2, 0x0b, 0x5f, 0x42, 0xea, 0x74, 0xf5, 0xfc, 0x18, 0x16,
	0x6f, 0x9c, 0x78, 0x13, 0xba, 0x4f, 0x7d, 0xfd, 0x9d, 0xef, 0xfc, 0xa6, 0x06, 0x1e, 0x1b, 0xb4,
	0x39, 0xc7, 0x5b, 0x8e, 0x5e, 0x78, 0xe4, 0x68, 0xf3, 0x91, 0x75, 0xc6, 0x1b, 0xf6, 0x62, 0xaa,
	0x71, 0x6a, 0xb8, 0x36, 0xc8, 0x97, 0x8e, 0x6b, 0xbb, 0x78, 0xc6, 0x09, 0x6a, 0xac, 0x72, 0x23,
	0x13, 0x71, 0x53, 0x85, 0xa8, 0xb0, 0x7e, 0x1a, 0x2d, 0x5c, 0x4e, 0x7f, 0x46, 0x81, 0x4c, 0xa3,
	0xd7, 0x53, 0x1c, 0x11, 0x7d, 0xe0, 0x4e, 0x3f, 0x85, 0x47, 0x77, 0x9c, 0xf1, 0x8f, 0x2f, 0x3f,
	0x9b, 0xb0, 0x13, 0xe8, 0x45, 0x0a, 0x5e, 0x8a, 0x99, 0x1a, 0x26, 0xc7, 0xc9, 0x69, 0x27, 0xeb,
	0x46, 0xdb, 0x27, 0x62, 0xa6, 0xd8, 0x1b, 0xd0, 0x5e, 0xb8, 0xbc, 0x3a, 0x6e, 0xd0, 0x71, 0x6b,
	0xe1, 0xf2, 0x70, 0x94, 0x3e, 0x87, 0x5e, 0xcd, 0xeb, 0xf5, 0x4c, 0xb1, 0x47, 0x70, 0x80, 0x6a,
	0x6a, 0x4a, 0x49, 0x3c, 0xfd, 0x2c, 0xbe, 0xb1, 0xb7, 0x00, 0x4a, 0x51, 0x9a, 0x78, 0xd6, 0xa0,
	0xb3, 0x0d, 0x4b, 0xfa, 0x67, 0x03, 0xd8, 0xa6, 0x40, 0xae, 0x4b, 0xef, 0x04, 0x7b, 0x1b, 0x06,
	0xa8, 0xb9, 0x70, 0x4a, 0x70, 0x2d, 0x39, 0x7a, 0x57, 0xcb, 0x43, 0x7d, 0xee, 0x94, 0xb8, 0x92,
	0x13, 0xef, 0xd8, 0x0f, 0x09, 0xf4, 0x51, 0x87, 0xcf, 0x9c, 0x27, 0x15, 0xc4, 0xdf, 0x3d, 0xfb,
	0x7c, 0xf4, 0x5f, 0xa7, 0x71, 0xb4, 0x19, 0x6b, 0x10, 0x31, 0x09, 0x3e, 0x5f, 0x86, 0xc0, 0xdf,
	0x81, 0x43, 0xd4, 0x7c, 0x7a, 0x2b, 0xca, 0x1b, 0xc5, 0xf3, 0x42, 0xdc, 0xe0, 0x70, 0x8f, 0xa2,
	0xec, 0xa3, 0xbe, 0x20, 0xeb, 0xf3, 0x60, 0x64, 0xdf, 0x41, 0x17, 0x35, 0x97, 0x73, 0x27, 0xbc,
	0x36, 0xe5, 0x70, 0xff, 0x5e, 0x94, 0x02, 0xea, 0x71, 0xf4, 0x98, 0xfe, 0xb4, 0x07, 0x47, 0x5b,
	0x99, 0xb6, 0xc2, 0x79, 0x2d, 0x0a, 0x76, 0x02, 0x7d, 0xb4, 0x5c, 0xc8, 0x05, 0x77, 0xde, 0x71,
	0x5d, 0x57, 0x10, 0xd0, 0x9e, 0xcb, 0x45, 0xe6, 0xdd, 0x95, 0xac, 0x32, 0x6d, 0xff, 0x8f, 0x4c,
	0xdb, 0x75, 0xa6, 0x53, 0xd2, 0x20, 0x15, 0x7a, 0x3e, 0x35, 0xf3, 0xd2, 0xc7, 0x3c, 0x77, 0xd1,
	0x8e, 0x15, 0xfa, 0x8b, 0x60, 0x62, 0xc7, 0xd0, 0xab, 0x31, 0x42, 0x4a, 0x47, 0x69, 0xa6, 0x50,
	0x02, 0xe4, 0x5c, 0x4a, 0x47, 0x75, 0xb0, 0xeb, 0x3a, 0x34, 0xef, 0xa9, 0x0e, 0x76, 0x55, 0x87,
	0xdf, 0x1b, 0xd0, 0xaf, 0x0e, 0x6d, 0x5e, 0xcd, 0xce, 0x49, 0x10, 0x9d, 0x73, 0xa9, 0xbf, 0xf8,
	0x12, 0xbd, 0x13, 0x31, 0xff, 0x5d, 0xb4, 0xf9, 0x38, 0x9a, 0xd8, 0x29, 0x3c, 0x08, 0x10, 0x1a,
	0x0e, 0x6e, 0x9d, 0xca, 0xf5, 0x32, 0x0e, 0xd3, 0x00, 0x6d, 0x7e, 0x15, 0xcc, 0x2f, 0xc8, 0xca,
	0xde, 0x83, 0xa3, 0x7f, 0x22, 0xb9, 0x54, 0x45, 0x4c, 0xd6, 0xc3, 0x6d, 0xf4, 0x58, 0x15, 0x1b,
	0xd4, 0xca, 0xd5, 0xd4, 0xfb, 0x9b, 0xd4, 0xca, 0xdd, 0xa1, 0x5e, 0x21, 0x89, 0xba, 0xb9, 0x49,
	0x5d, 0xa3, 0x03, 0xf5, 0x13, 0x08, 0x14, 0x5c, 0x2d, 0x7d, 0x4d, 0x7c, 0x40, 0xd0, 0x10, 0xee,
	0xe5, 0xd2, 0x47, 0xda, 0x77, 0x81, 0x6d, 0xa3, 0x88, 0xb4, 0x45, 0xc8, 0xc3, 0x4d, 0xe4, 0x58,
	0x15, 0xe9, 0x37, 0x70, 0x48, 0xc9, 0x2b, 0x50, 0xc4, 0xa1, 0x63, 0x8f, 0xa1, 0x15, 0x17, 0x45,
	0xbd, 0x7b, 0x04, 0xad, 0x08, 0x3a, 0xa8, 0xba, 0x3a, 0x6e, 0xaf, 0x03, 0x41, 0x0d, 0xcd, 0x06,
	0xd0, 0xd0, 0x92, 0x32, 0xd2, 0xc9, 0x1a, 0x5a, 0x32, 0x06, 0xfb, 0xfe, 0x95, 0x55, 0x31, 0x6c,
	0x7a, 0x66, 0x47, 0xd0, 0xcc, 0x8b, 0x39, 0xde, 0x52, 0x74, 0xed, 0xac, 0x7a, 0x49, 0xff, 0x48,
	0xe2, 0xf2, 0x56, 0x72, 0x46, 0x05, 0x24, 0xcf, 0x54, 0xf1, 0xd7, 0xeb, 0xf8, 0x0a, 0xda, 0x75,
	0xad, 0xe3, 0xdc, 0xf0, 0x5d, 0xf5, 0x5b, 0x74, 0x93, 0xb5, 0xd0, 0xe6, 0x34, 0x34, 0xc7, 0xd0,
	0x0b, 0xa9, 0x0a, 0x21, 0xf1, 0x29, 0xcd, 0xcc, 0x5e, 0x18, 0x88, 0x02, 0xc5, 0xcb, 0x57, 0x56,
	0x5d, 0x94, 0x3e, 0xfd, 0x6b, 0x2f, 0x6e, 0x60, 0x37, 0x2f, 0xc3, 0xb7, 0x31, 0x9a, 0x58, 0xbb,
	0x8d, 0x91, 0x4f, 0x56, 0xb5, 0x5b, 0xcf, 0xe4, 0xf7, 0x09, 0xf4, 0x6e, 0xae, 0x0b, 0x7e, 0xdf,
	0xf1, 0xc1, 0xcd, 0x75, 0x31, 0x89, 0x21, 0x3e, 0x81, 0xc1, 0xba, 0x1b, 0x62, 0x90, 0xa1, 0xb2,
	0x21, 0xf0, 0x6a, 0x03, 0x5f, 0x94, 0x9e, 0x7d, 0x0d, 0xed, 0x80, 0xd2, 0x65, 0x6e, 0x86, 0xfb,
	0xc7, 0x7b, 0xa7, 0xdd, 0x33, 0xb1, 0x23, 0x91, 0x6b, 0x31, 0x59, 0xab, 0x40, 0x71, 0x55, 0xe6,
	0x86, 0xfd, 0x9c, 0x40, 0x67, 0xd5, 0x2a, 0xc3, 0x26, 0xf9, 0xd7, 0x3b, 0xf2, 0x7f, 0xb7, 0x35,
	0xb3, 0x76, 0x78, 0x9e, 0x78, 0xe1, 0xd3, 0x8f, 0x62, 0xb1, 0x03, 0xa8, 0x02, 0xcc, 0x67, 0xb3,
	0xd7, 0xb7, 0xee, 0x9b, 0xd0, 0x09, 0xe8, 0x6a, 0xdf, 0x56, 0x0b, 0x27, 0xf4, 0x32, 0x2d, 0xdb,
	0x30, 0x0c, 0x0f, 0x88, 0xec, 0x56, 0x09, 0xa9, 0x1c, 0xe5, 0x36, 0x7c, 0xe1, 0xcc, 0x3c, 0x6c,
	0x88, 0x48, 0xd6, 0xc9, 0xda, 0x95, 0xe1, 0x5f, 0xe8, 0x42, 0x2f, 0x75, 0x56, 0x92, 0xa8, 0x51,
	0xbb, 0x67, 0x72, 0x87, 0x39, 0x5a, 0xf9, 0x8a, 0xe9, 0x99, 0xcf, 0x66, 0xe9, 0x6f, 0x4d, 0x78,
	0x78, 0xe7, 0xbe, 0xc4, 0x7e, 0x4c, 0x00, 0x0c, 0xd6, 0x03, 0x32, 0x3c, 0xdb, 0xa9, 0xb4, 0xad,
	0x8b, 0x50, 0xd6, 0x31, 0x98, 0x55, 0x6e, 0xd9, 0xaf, 0x09, 0x0c, 0x0c, 0x06, 0xc5, 0x2b, 0x25,
	0x4f, 0x49, 0x49, 0xbe, 0x63, 0x25, 0xf1, 0xa2, 0x90, 0xf5, 0x0c, 0x4e, 0xe6, 0xb3, 0x5a, 0xce,
	0x2f, 0x49, 0xf8, 0x3f, 0xc6, 0xd5, 0x72, 0xa5, 0xe6, 0xd9, 0xbd, 0xaa, 0xe9, 0x1a, 0xbc, 0x5c,
	0xd6, 0x62, 0xbe, 0x85, 0x56, 0xad, 0xe2, 0xfd, 0x9d, 0x56, 0x67, 0x6b, 0x49, 0x66, 0xb5, 0xd3,
	0xd0, 0xbb, 0xb0, 0x9e, 0x84, 0xe1, 0x07, 0xb4, 0x05, 0xaf, 0x77, 0xa4, 0x61, 0x63, 0xdc, 0xb2,
	0x30, 0x4f, 0x1f, 0xd2, 0xfb, 0xf5, 0x01, 0xfd, 0x86, 0x78, 0xfa, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0xfc, 0x90, 0x67, 0x5e, 0x0c, 0x00, 0x00,
}
