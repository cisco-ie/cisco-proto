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
// source: ospfv3_edm_spf_stats.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_vrf_statistics_spf_stats

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

type Ospfv3EdmSpfStats_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmSpfStats_KEYS) Reset()         { *m = Ospfv3EdmSpfStats_KEYS{} }
func (m *Ospfv3EdmSpfStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmSpfStats_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmSpfStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac3b88eb82da8de, []int{0}
}

func (m *Ospfv3EdmSpfStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmSpfStats_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmSpfStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmSpfStats_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmSpfStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmSpfStats_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmSpfStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmSpfStats_KEYS.Size(m)
}
func (m *Ospfv3EdmSpfStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmSpfStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmSpfStats_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmSpfStats_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmSpfStats_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type Ospfv3EdmAreaSumm struct {
	AreaId               uint32   `protobuf:"varint,1,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	SpFs                 uint32   `protobuf:"varint,2,opt,name=sp_fs,json=spFs,proto3" json:"sp_fs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmAreaSumm) Reset()         { *m = Ospfv3EdmAreaSumm{} }
func (m *Ospfv3EdmAreaSumm) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmAreaSumm) ProtoMessage()    {}
func (*Ospfv3EdmAreaSumm) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac3b88eb82da8de, []int{1}
}

func (m *Ospfv3EdmAreaSumm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmAreaSumm.Unmarshal(m, b)
}
func (m *Ospfv3EdmAreaSumm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmAreaSumm.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmAreaSumm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmAreaSumm.Merge(m, src)
}
func (m *Ospfv3EdmAreaSumm) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmAreaSumm.Size(m)
}
func (m *Ospfv3EdmAreaSumm) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmAreaSumm.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmAreaSumm proto.InternalMessageInfo

func (m *Ospfv3EdmAreaSumm) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Ospfv3EdmAreaSumm) GetSpFs() uint32 {
	if m != nil {
		return m.SpFs
	}
	return 0
}

type Ospfv3HeaderInfo struct {
	HeaderRouterId       string               `protobuf:"bytes,1,opt,name=header_router_id,json=headerRouterId,proto3" json:"header_router_id,omitempty"`
	HeaderSpFs           uint32               `protobuf:"varint,2,opt,name=header_sp_fs,json=headerSpFs,proto3" json:"header_sp_fs,omitempty"`
	AreaSummary          []*Ospfv3EdmAreaSumm `protobuf:"bytes,3,rep,name=area_summary,json=areaSummary,proto3" json:"area_summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Ospfv3HeaderInfo) Reset()         { *m = Ospfv3HeaderInfo{} }
func (m *Ospfv3HeaderInfo) String() string { return proto.CompactTextString(m) }
func (*Ospfv3HeaderInfo) ProtoMessage()    {}
func (*Ospfv3HeaderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac3b88eb82da8de, []int{2}
}

func (m *Ospfv3HeaderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3HeaderInfo.Unmarshal(m, b)
}
func (m *Ospfv3HeaderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3HeaderInfo.Marshal(b, m, deterministic)
}
func (m *Ospfv3HeaderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3HeaderInfo.Merge(m, src)
}
func (m *Ospfv3HeaderInfo) XXX_Size() int {
	return xxx_messageInfo_Ospfv3HeaderInfo.Size(m)
}
func (m *Ospfv3HeaderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3HeaderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3HeaderInfo proto.InternalMessageInfo

func (m *Ospfv3HeaderInfo) GetHeaderRouterId() string {
	if m != nil {
		return m.HeaderRouterId
	}
	return ""
}

func (m *Ospfv3HeaderInfo) GetHeaderSpFs() uint32 {
	if m != nil {
		return m.HeaderSpFs
	}
	return 0
}

func (m *Ospfv3HeaderInfo) GetAreaSummary() []*Ospfv3EdmAreaSumm {
	if m != nil {
		return m.AreaSummary
	}
	return nil
}

type Ospfv3SpfTime struct {
	Dijkstra             uint32   `protobuf:"varint,1,opt,name=dijkstra,proto3" json:"dijkstra,omitempty"`
	IntraPrefix          uint32   `protobuf:"varint,2,opt,name=intra_prefix,json=intraPrefix,proto3" json:"intra_prefix,omitempty"`
	IntraPrefixDel       uint32   `protobuf:"varint,3,opt,name=intra_prefix_del,json=intraPrefixDel,proto3" json:"intra_prefix_del,omitempty"`
	InterPrefix          uint32   `protobuf:"varint,4,opt,name=inter_prefix,json=interPrefix,proto3" json:"inter_prefix,omitempty"`
	InterPrefixDel       uint32   `protobuf:"varint,5,opt,name=inter_prefix_del,json=interPrefixDel,proto3" json:"inter_prefix_del,omitempty"`
	ExternalPrefix       uint32   `protobuf:"varint,6,opt,name=external_prefix,json=externalPrefix,proto3" json:"external_prefix,omitempty"`
	ExternalPrefixDel    uint32   `protobuf:"varint,7,opt,name=external_prefix_del,json=externalPrefixDel,proto3" json:"external_prefix_del,omitempty"`
	RibAdd               uint32   `protobuf:"varint,8,opt,name=rib_add,json=ribAdd,proto3" json:"rib_add,omitempty"`
	RibDel               uint32   `protobuf:"varint,9,opt,name=rib_del,json=ribDel,proto3" json:"rib_del,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3SpfTime) Reset()         { *m = Ospfv3SpfTime{} }
func (m *Ospfv3SpfTime) String() string { return proto.CompactTextString(m) }
func (*Ospfv3SpfTime) ProtoMessage()    {}
func (*Ospfv3SpfTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac3b88eb82da8de, []int{3}
}

func (m *Ospfv3SpfTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3SpfTime.Unmarshal(m, b)
}
func (m *Ospfv3SpfTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3SpfTime.Marshal(b, m, deterministic)
}
func (m *Ospfv3SpfTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3SpfTime.Merge(m, src)
}
func (m *Ospfv3SpfTime) XXX_Size() int {
	return xxx_messageInfo_Ospfv3SpfTime.Size(m)
}
func (m *Ospfv3SpfTime) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3SpfTime.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3SpfTime proto.InternalMessageInfo

func (m *Ospfv3SpfTime) GetDijkstra() uint32 {
	if m != nil {
		return m.Dijkstra
	}
	return 0
}

func (m *Ospfv3SpfTime) GetIntraPrefix() uint32 {
	if m != nil {
		return m.IntraPrefix
	}
	return 0
}

func (m *Ospfv3SpfTime) GetIntraPrefixDel() uint32 {
	if m != nil {
		return m.IntraPrefixDel
	}
	return 0
}

func (m *Ospfv3SpfTime) GetInterPrefix() uint32 {
	if m != nil {
		return m.InterPrefix
	}
	return 0
}

func (m *Ospfv3SpfTime) GetInterPrefixDel() uint32 {
	if m != nil {
		return m.InterPrefixDel
	}
	return 0
}

func (m *Ospfv3SpfTime) GetExternalPrefix() uint32 {
	if m != nil {
		return m.ExternalPrefix
	}
	return 0
}

func (m *Ospfv3SpfTime) GetExternalPrefixDel() uint32 {
	if m != nil {
		return m.ExternalPrefixDel
	}
	return 0
}

func (m *Ospfv3SpfTime) GetRibAdd() uint32 {
	if m != nil {
		return m.RibAdd
	}
	return 0
}

func (m *Ospfv3SpfTime) GetRibDel() uint32 {
	if m != nil {
		return m.RibDel
	}
	return 0
}

type Ospfv3LsaChange struct {
	LsaAreaId            uint32   `protobuf:"varint,1,opt,name=lsa_area_id,json=lsaAreaId,proto3" json:"lsa_area_id,omitempty"`
	LsaAdvertisingRouter string   `protobuf:"bytes,2,opt,name=lsa_advertising_router,json=lsaAdvertisingRouter,proto3" json:"lsa_advertising_router,omitempty"`
	LsaId                string   `protobuf:"bytes,3,opt,name=lsa_id,json=lsaId,proto3" json:"lsa_id,omitempty"`
	LsaType              uint32   `protobuf:"varint,4,opt,name=lsa_type,json=lsaType,proto3" json:"lsa_type,omitempty"`
	LsaFlush             bool     `protobuf:"varint,5,opt,name=lsa_flush,json=lsaFlush,proto3" json:"lsa_flush,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3LsaChange) Reset()         { *m = Ospfv3LsaChange{} }
func (m *Ospfv3LsaChange) String() string { return proto.CompactTextString(m) }
func (*Ospfv3LsaChange) ProtoMessage()    {}
func (*Ospfv3LsaChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac3b88eb82da8de, []int{4}
}

func (m *Ospfv3LsaChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3LsaChange.Unmarshal(m, b)
}
func (m *Ospfv3LsaChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3LsaChange.Marshal(b, m, deterministic)
}
func (m *Ospfv3LsaChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3LsaChange.Merge(m, src)
}
func (m *Ospfv3LsaChange) XXX_Size() int {
	return xxx_messageInfo_Ospfv3LsaChange.Size(m)
}
func (m *Ospfv3LsaChange) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3LsaChange.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3LsaChange proto.InternalMessageInfo

func (m *Ospfv3LsaChange) GetLsaAreaId() uint32 {
	if m != nil {
		return m.LsaAreaId
	}
	return 0
}

func (m *Ospfv3LsaChange) GetLsaAdvertisingRouter() string {
	if m != nil {
		return m.LsaAdvertisingRouter
	}
	return ""
}

func (m *Ospfv3LsaChange) GetLsaId() string {
	if m != nil {
		return m.LsaId
	}
	return ""
}

func (m *Ospfv3LsaChange) GetLsaType() uint32 {
	if m != nil {
		return m.LsaType
	}
	return 0
}

func (m *Ospfv3LsaChange) GetLsaFlush() bool {
	if m != nil {
		return m.LsaFlush
	}
	return false
}

type Ospfv3EdmSpfAreaStats struct {
	SpfStatAreaId        uint32         `protobuf:"varint,1,opt,name=spf_stat_area_id,json=spfStatAreaId,proto3" json:"spf_stat_area_id,omitempty"`
	SpfStatTime          *Ospfv3SpfTime `protobuf:"bytes,2,opt,name=spf_stat_time,json=spfStatTime,proto3" json:"spf_stat_time,omitempty"`
	SpfStatLsaTypeCount  []uint32       `protobuf:"varint,3,rep,packed,name=spf_stat_lsa_type_count,json=spfStatLsaTypeCount,proto3" json:"spf_stat_lsa_type_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Ospfv3EdmSpfAreaStats) Reset()         { *m = Ospfv3EdmSpfAreaStats{} }
func (m *Ospfv3EdmSpfAreaStats) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmSpfAreaStats) ProtoMessage()    {}
func (*Ospfv3EdmSpfAreaStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac3b88eb82da8de, []int{5}
}

func (m *Ospfv3EdmSpfAreaStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmSpfAreaStats.Unmarshal(m, b)
}
func (m *Ospfv3EdmSpfAreaStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmSpfAreaStats.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmSpfAreaStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmSpfAreaStats.Merge(m, src)
}
func (m *Ospfv3EdmSpfAreaStats) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmSpfAreaStats.Size(m)
}
func (m *Ospfv3EdmSpfAreaStats) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmSpfAreaStats.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmSpfAreaStats proto.InternalMessageInfo

func (m *Ospfv3EdmSpfAreaStats) GetSpfStatAreaId() uint32 {
	if m != nil {
		return m.SpfStatAreaId
	}
	return 0
}

func (m *Ospfv3EdmSpfAreaStats) GetSpfStatTime() *Ospfv3SpfTime {
	if m != nil {
		return m.SpfStatTime
	}
	return nil
}

func (m *Ospfv3EdmSpfAreaStats) GetSpfStatLsaTypeCount() []uint32 {
	if m != nil {
		return m.SpfStatLsaTypeCount
	}
	return nil
}

type Ospfv3RuntimeStats struct {
	StartTime            uint32                   `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	GlobalTime           *Ospfv3SpfTime           `protobuf:"bytes,2,opt,name=global_time,json=globalTime,proto3" json:"global_time,omitempty"`
	RibAddRoutes         uint32                   `protobuf:"varint,3,opt,name=rib_add_routes,json=ribAddRoutes,proto3" json:"rib_add_routes,omitempty"`
	RibDeleteRoutes      uint32                   `protobuf:"varint,4,opt,name=rib_delete_routes,json=ribDeleteRoutes,proto3" json:"rib_delete_routes,omitempty"`
	ReasonFlags          []byte                   `protobuf:"bytes,5,opt,name=reason_flags,json=reasonFlags,proto3" json:"reason_flags,omitempty"`
	LsaChanges           string                   `protobuf:"bytes,6,opt,name=lsa_changes,json=lsaChanges,proto3" json:"lsa_changes,omitempty"`
	Lsa                  []*Ospfv3LsaChange       `protobuf:"bytes,7,rep,name=lsa,proto3" json:"lsa,omitempty"`
	AreaStat             []*Ospfv3EdmSpfAreaStats `protobuf:"bytes,8,rep,name=area_stat,json=areaStat,proto3" json:"area_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Ospfv3RuntimeStats) Reset()         { *m = Ospfv3RuntimeStats{} }
func (m *Ospfv3RuntimeStats) String() string { return proto.CompactTextString(m) }
func (*Ospfv3RuntimeStats) ProtoMessage()    {}
func (*Ospfv3RuntimeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac3b88eb82da8de, []int{6}
}

func (m *Ospfv3RuntimeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3RuntimeStats.Unmarshal(m, b)
}
func (m *Ospfv3RuntimeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3RuntimeStats.Marshal(b, m, deterministic)
}
func (m *Ospfv3RuntimeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3RuntimeStats.Merge(m, src)
}
func (m *Ospfv3RuntimeStats) XXX_Size() int {
	return xxx_messageInfo_Ospfv3RuntimeStats.Size(m)
}
func (m *Ospfv3RuntimeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3RuntimeStats.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3RuntimeStats proto.InternalMessageInfo

func (m *Ospfv3RuntimeStats) GetStartTime() uint32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Ospfv3RuntimeStats) GetGlobalTime() *Ospfv3SpfTime {
	if m != nil {
		return m.GlobalTime
	}
	return nil
}

func (m *Ospfv3RuntimeStats) GetRibAddRoutes() uint32 {
	if m != nil {
		return m.RibAddRoutes
	}
	return 0
}

func (m *Ospfv3RuntimeStats) GetRibDeleteRoutes() uint32 {
	if m != nil {
		return m.RibDeleteRoutes
	}
	return 0
}

func (m *Ospfv3RuntimeStats) GetReasonFlags() []byte {
	if m != nil {
		return m.ReasonFlags
	}
	return nil
}

func (m *Ospfv3RuntimeStats) GetLsaChanges() string {
	if m != nil {
		return m.LsaChanges
	}
	return ""
}

func (m *Ospfv3RuntimeStats) GetLsa() []*Ospfv3LsaChange {
	if m != nil {
		return m.Lsa
	}
	return nil
}

func (m *Ospfv3RuntimeStats) GetAreaStat() []*Ospfv3EdmSpfAreaStats {
	if m != nil {
		return m.AreaStat
	}
	return nil
}

type Ospfv3EdmSpfStats struct {
	SpfHeader            *Ospfv3HeaderInfo     `protobuf:"bytes,50,opt,name=spf_header,json=spfHeader,proto3" json:"spf_header,omitempty"`
	SpfRuntime           []*Ospfv3RuntimeStats `protobuf:"bytes,51,rep,name=spf_runtime,json=spfRuntime,proto3" json:"spf_runtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Ospfv3EdmSpfStats) Reset()         { *m = Ospfv3EdmSpfStats{} }
func (m *Ospfv3EdmSpfStats) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmSpfStats) ProtoMessage()    {}
func (*Ospfv3EdmSpfStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac3b88eb82da8de, []int{7}
}

func (m *Ospfv3EdmSpfStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmSpfStats.Unmarshal(m, b)
}
func (m *Ospfv3EdmSpfStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmSpfStats.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmSpfStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmSpfStats.Merge(m, src)
}
func (m *Ospfv3EdmSpfStats) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmSpfStats.Size(m)
}
func (m *Ospfv3EdmSpfStats) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmSpfStats.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmSpfStats proto.InternalMessageInfo

func (m *Ospfv3EdmSpfStats) GetSpfHeader() *Ospfv3HeaderInfo {
	if m != nil {
		return m.SpfHeader
	}
	return nil
}

func (m *Ospfv3EdmSpfStats) GetSpfRuntime() []*Ospfv3RuntimeStats {
	if m != nil {
		return m.SpfRuntime
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3EdmSpfStats_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.vrf_statistics.spf_stats.ospfv3_edm_spf_stats_KEYS")
	proto.RegisterType((*Ospfv3EdmAreaSumm)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.vrf_statistics.spf_stats.ospfv3_edm_area_summ")
	proto.RegisterType((*Ospfv3HeaderInfo)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.vrf_statistics.spf_stats.ospfv3_header_info")
	proto.RegisterType((*Ospfv3SpfTime)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.vrf_statistics.spf_stats.ospfv3_spf_time")
	proto.RegisterType((*Ospfv3LsaChange)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.vrf_statistics.spf_stats.ospfv3_lsa_change")
	proto.RegisterType((*Ospfv3EdmSpfAreaStats)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.vrf_statistics.spf_stats.ospfv3_edm_spf_area_stats")
	proto.RegisterType((*Ospfv3RuntimeStats)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.vrf_statistics.spf_stats.ospfv3_runtime_stats")
	proto.RegisterType((*Ospfv3EdmSpfStats)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.vrf_statistics.spf_stats.ospfv3_edm_spf_stats")
}

func init() { proto.RegisterFile("ospfv3_edm_spf_stats.proto", fileDescriptor_cac3b88eb82da8de) }

var fileDescriptor_cac3b88eb82da8de = []byte{
	// 807 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcd, 0x4e, 0xf3, 0x46,
	0x14, 0x95, 0x13, 0xc8, 0xcf, 0x75, 0x02, 0x1f, 0xc3, 0xd7, 0x62, 0xa8, 0xda, 0x82, 0x55, 0x09,
	0xd4, 0x45, 0x16, 0x80, 0xba, 0x47, 0xd0, 0xa8, 0xa8, 0x55, 0x55, 0x4d, 0x58, 0x94, 0xd5, 0x68,
	0x12, 0x8f, 0xc3, 0x14, 0xc7, 0xb6, 0xe6, 0x4e, 0x22, 0x58, 0xf4, 0x05, 0xa8, 0xaa, 0xb6, 0x2f,
	0x53, 0xf5, 0x39, 0xfa, 0x18, 0x5d, 0xf7, 0x01, 0xaa, 0xf9, 0xb1, 0x49, 0x28, 0xdb, 0xf0, 0x2d,
	0x40, 0xbe, 0xe7, 0x1e, 0x9f, 0xb9, 0x73, 0x7d, 0xee, 0x4c, 0xe0, 0xa0, 0xc0, 0x32, 0x5d, 0x9c,
	0x31, 0x91, 0xcc, 0x18, 0x96, 0x29, 0x43, 0xcd, 0x35, 0x0e, 0x4a, 0x55, 0xe8, 0x82, 0xfc, 0x38,
	0x91, 0x38, 0x29, 0x98, 0x2c, 0x90, 0x3d, 0x28, 0x26, 0xcb, 0xc5, 0x57, 0xcc, 0xb3, 0x8b, 0x52,
	0xa8, 0x81, 0x7b, 0x36, 0xdc, 0x89, 0x40, 0x14, 0x58, 0x3d, 0x0d, 0x16, 0x2a, 0xb5, 0xff, 0xcc,
	0x9f, 0x15, 0x94, 0xa8, 0xe5, 0x04, 0x07, 0xb5, 0x7e, 0x7c, 0x0b, 0xfb, 0xaf, 0xad, 0xcb, 0xbe,
	0xfd, 0xfa, 0x76, 0x44, 0x8e, 0xa0, 0xe7, 0x95, 0x58, 0xce, 0x67, 0x22, 0x0a, 0x0e, 0x83, 0x93,
	0x2e, 0x0d, 0x3d, 0xf6, 0x3d, 0x9f, 0x09, 0xb2, 0x0f, 0x1d, 0xa3, 0x6d, 0xd3, 0x0d, 0x9b, 0x6e,
	0x2f, 0x54, 0x6a, 0x52, 0xf1, 0x15, 0xbc, 0x5f, 0x92, 0xe6, 0x4a, 0x70, 0x86, 0xf3, 0xd9, 0x8c,
	0xec, 0x41, 0xdb, 0x06, 0x32, 0xb1, 0x82, 0x7d, 0xda, 0x32, 0xe1, 0x75, 0x42, 0x76, 0x61, 0x13,
	0x4b, 0x96, 0xa2, 0x15, 0xea, 0xd3, 0x0d, 0x2c, 0x87, 0x18, 0xff, 0x1b, 0x00, 0xf1, 0x32, 0x77,
	0x82, 0x27, 0x42, 0x31, 0x99, 0xa7, 0x05, 0x39, 0x81, 0x77, 0x3e, 0x54, 0xc5, 0x5c, 0x1b, 0x34,
	0xf1, 0xe5, 0x6d, 0x39, 0x9c, 0x5a, 0xf8, 0x3a, 0x21, 0x87, 0xd0, 0xf3, 0xcc, 0x65, 0x71, 0x70,
	0xd8, 0xa8, 0x1c, 0x22, 0xf9, 0x23, 0x80, 0x5e, 0x5d, 0x1e, 0x57, 0x8f, 0x51, 0xf3, 0xb0, 0x79,
	0x12, 0x9e, 0xe6, 0x83, 0x75, 0x75, 0x7d, 0xf0, 0x5a, 0x5f, 0x68, 0x68, 0x1e, 0x47, 0xae, 0x84,
	0xf8, 0xef, 0x06, 0x6c, 0x7b, 0x96, 0x79, 0x4d, 0xcb, 0x99, 0x20, 0x07, 0xd0, 0x49, 0xe4, 0x4f,
	0xf7, 0xa8, 0x15, 0xf7, 0x9d, 0xab, 0x63, 0xf3, 0xa9, 0x64, 0xae, 0x15, 0x67, 0xa5, 0x12, 0xa9,
	0x7c, 0xf0, 0xbb, 0x0c, 0x2d, 0xf6, 0x83, 0x85, 0x4c, 0xcb, 0x96, 0x29, 0x2c, 0x11, 0x59, 0xd4,
	0xb4, 0xb4, 0xad, 0x25, 0xda, 0x95, 0xc8, 0xbc, 0x98, 0x50, 0x95, 0xd8, 0x46, 0x2d, 0x26, 0xd4,
	0x8a, 0x58, 0x4d, 0xb1, 0x62, 0x9b, 0xb5, 0x58, 0x45, 0x33, 0x62, 0xc7, 0xb0, 0x2d, 0x1e, 0xb4,
	0x50, 0x39, 0xcf, 0x2a, 0xbd, 0x96, 0x23, 0x56, 0xb0, 0x97, 0x1c, 0xc0, 0xee, 0x0b, 0xa2, 0x55,
	0x6d, 0x5b, 0xf2, 0xce, 0x2a, 0xd9, 0x08, 0xef, 0x41, 0x5b, 0xc9, 0x31, 0xe3, 0x49, 0x12, 0x75,
	0x9c, 0x8f, 0x94, 0x1c, 0x5f, 0x24, 0x49, 0x95, 0x30, 0x2f, 0x77, 0xeb, 0xc4, 0x95, 0xc8, 0xe2,
	0xbf, 0x02, 0xd8, 0xf1, 0x4d, 0xcd, 0x90, 0xb3, 0xc9, 0x1d, 0xcf, 0xa7, 0x82, 0x7c, 0x06, 0xa1,
	0x89, 0x56, 0x3d, 0xd9, 0xcd, 0x90, 0x5f, 0x38, 0x5b, 0x9e, 0xc3, 0xc7, 0x36, 0x9f, 0x2c, 0x84,
	0xd2, 0x12, 0x65, 0x3e, 0xf5, 0x9e, 0xf3, 0x86, 0x7f, 0x6f, 0xa8, 0xcf, 0x49, 0x67, 0x3c, 0xf2,
	0x11, 0xb4, 0xcc, 0x5b, 0x32, 0xb1, 0x3d, 0xee, 0xd2, 0xcd, 0x0c, 0x8d, 0xd8, 0x3e, 0x74, 0x0c,
	0xac, 0x1f, 0x4b, 0xe1, 0xdb, 0xda, 0xce, 0x90, 0xdf, 0x3c, 0x96, 0x82, 0x7c, 0x02, 0x66, 0x51,
	0x96, 0x66, 0x73, 0xbc, 0xb3, 0xbd, 0xec, 0x50, 0xc3, 0x1d, 0x9a, 0x38, 0x7e, 0x6a, 0xfc, 0x6f,
	0x50, 0x9d, 0x73, 0x8c, 0x9f, 0xc8, 0x31, 0xbc, 0xab, 0xcc, 0xf5, 0x62, 0x1f, 0x7d, 0x2c, 0xd3,
	0x91, 0xe6, 0xda, 0xef, 0xe5, 0xd7, 0x00, 0xfa, 0x35, 0xd3, 0x98, 0xca, 0xee, 0x21, 0x3c, 0x95,
	0x6b, 0xf7, 0x7a, 0xe5, 0x62, 0x1a, 0xfa, 0x8a, 0x6e, 0x8c, 0xa5, 0xcf, 0x61, 0xaf, 0x2e, 0xa7,
	0xea, 0x0b, 0x9b, 0x14, 0xf3, 0x5c, 0xdb, 0x21, 0xec, 0xd3, 0x5d, 0xcf, 0xfe, 0xce, 0x35, 0xe9,
	0xd2, 0xa4, 0xe2, 0x7f, 0x36, 0xea, 0xa3, 0x45, 0xcd, 0x73, 0xa3, 0xea, 0xfb, 0xf0, 0x29, 0x00,
	0x6a, 0xae, 0xfc, 0xd6, 0xfc, 0x97, 0xb4, 0x88, 0x5d, 0xed, 0x29, 0x80, 0x70, 0x9a, 0x15, 0x63,
	0x9e, 0x7d, 0xa0, 0xbd, 0x83, 0x5b, 0xdd, 0x16, 0xf3, 0x05, 0x6c, 0x79, 0xfb, 0x3a, 0x3b, 0xa1,
	0x1f, 0xc6, 0x9e, 0x73, 0xb1, 0xb5, 0x11, 0x92, 0x2f, 0x61, 0xc7, 0x7b, 0x59, 0x68, 0x51, 0x11,
	0x9d, 0x71, 0xb6, 0x9d, 0xab, 0x85, 0x16, 0x9e, 0x7b, 0x04, 0x3d, 0x25, 0x38, 0x16, 0x39, 0x4b,
	0x33, 0x3e, 0x45, 0xeb, 0xa1, 0x1e, 0x0d, 0x1d, 0x36, 0x34, 0x10, 0xf9, 0xdc, 0x79, 0xdd, 0x39,
	0x1f, 0xed, 0x20, 0x76, 0x29, 0x64, 0xc8, 0x2f, 0x1d, 0x42, 0x7e, 0x86, 0x66, 0x86, 0x3c, 0x6a,
	0xdb, 0x13, 0xf0, 0x7e, 0xed, 0x9d, 0x79, 0x2e, 0x86, 0x9a, 0x75, 0xc9, 0xef, 0x01, 0x74, 0x6b,
	0x5f, 0x47, 0x1d, 0x5b, 0x05, 0xbe, 0xc9, 0x39, 0xbc, 0x3a, 0x51, 0xb4, 0x63, 0x0f, 0x63, 0xcd,
	0x75, 0xfc, 0x67, 0x63, 0xe5, 0x1e, 0xab, 0x5f, 0x26, 0xbf, 0x04, 0x00, 0x26, 0x72, 0x37, 0x49,
	0x74, 0x6a, 0xcd, 0x94, 0xad, 0xbd, 0xd8, 0xa5, 0x5b, 0x90, 0x76, 0xb1, 0x4c, 0xbf, 0xb1, 0x31,
	0xf9, 0x2d, 0x00, 0x33, 0x59, 0xd5, 0x40, 0x44, 0x67, 0x6f, 0x74, 0x87, 0xad, 0x0c, 0x20, 0x35,
	0x0d, 0xa1, 0x0e, 0x19, 0xb7, 0xec, 0x6f, 0x97, 0xb3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0x80, 0xc6, 0x23, 0xd9, 0x08, 0x00, 0x00,
}
