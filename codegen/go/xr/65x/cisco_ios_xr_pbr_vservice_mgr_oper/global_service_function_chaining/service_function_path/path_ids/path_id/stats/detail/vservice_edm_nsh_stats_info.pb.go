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
// source: vservice_edm_nsh_stats_info.proto

package cisco_ios_xr_pbr_vservice_mgr_oper_global_service_function_chaining_service_function_path_path_ids_path_id_stats_detail

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

type VserviceEdmNshStatsInfo_KEYS struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VserviceEdmNshStatsInfo_KEYS) Reset()         { *m = VserviceEdmNshStatsInfo_KEYS{} }
func (m *VserviceEdmNshStatsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*VserviceEdmNshStatsInfo_KEYS) ProtoMessage()    {}
func (*VserviceEdmNshStatsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6caaa09d9c151, []int{0}
}

func (m *VserviceEdmNshStatsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VserviceEdmNshStatsInfo_KEYS.Unmarshal(m, b)
}
func (m *VserviceEdmNshStatsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VserviceEdmNshStatsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *VserviceEdmNshStatsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VserviceEdmNshStatsInfo_KEYS.Merge(m, src)
}
func (m *VserviceEdmNshStatsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_VserviceEdmNshStatsInfo_KEYS.Size(m)
}
func (m *VserviceEdmNshStatsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_VserviceEdmNshStatsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_VserviceEdmNshStatsInfo_KEYS proto.InternalMessageInfo

func (m *VserviceEdmNshStatsInfo_KEYS) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type VserviceNshSpiSi struct {
	ProcessedPkts        uint64   `protobuf:"varint,1,opt,name=processed_pkts,json=processedPkts,proto3" json:"processed_pkts,omitempty"`
	ProcessedBytes       uint64   `protobuf:"varint,2,opt,name=processed_bytes,json=processedBytes,proto3" json:"processed_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VserviceNshSpiSi) Reset()         { *m = VserviceNshSpiSi{} }
func (m *VserviceNshSpiSi) String() string { return proto.CompactTextString(m) }
func (*VserviceNshSpiSi) ProtoMessage()    {}
func (*VserviceNshSpiSi) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6caaa09d9c151, []int{1}
}

func (m *VserviceNshSpiSi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VserviceNshSpiSi.Unmarshal(m, b)
}
func (m *VserviceNshSpiSi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VserviceNshSpiSi.Marshal(b, m, deterministic)
}
func (m *VserviceNshSpiSi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VserviceNshSpiSi.Merge(m, src)
}
func (m *VserviceNshSpiSi) XXX_Size() int {
	return xxx_messageInfo_VserviceNshSpiSi.Size(m)
}
func (m *VserviceNshSpiSi) XXX_DiscardUnknown() {
	xxx_messageInfo_VserviceNshSpiSi.DiscardUnknown(m)
}

var xxx_messageInfo_VserviceNshSpiSi proto.InternalMessageInfo

func (m *VserviceNshSpiSi) GetProcessedPkts() uint64 {
	if m != nil {
		return m.ProcessedPkts
	}
	return 0
}

func (m *VserviceNshSpiSi) GetProcessedBytes() uint64 {
	if m != nil {
		return m.ProcessedBytes
	}
	return 0
}

type VserviceNshSpiSiTerminate struct {
	TerminatedPkts       uint64   `protobuf:"varint,1,opt,name=terminated_pkts,json=terminatedPkts,proto3" json:"terminated_pkts,omitempty"`
	TerminatedBytes      uint64   `protobuf:"varint,2,opt,name=terminated_bytes,json=terminatedBytes,proto3" json:"terminated_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VserviceNshSpiSiTerminate) Reset()         { *m = VserviceNshSpiSiTerminate{} }
func (m *VserviceNshSpiSiTerminate) String() string { return proto.CompactTextString(m) }
func (*VserviceNshSpiSiTerminate) ProtoMessage()    {}
func (*VserviceNshSpiSiTerminate) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6caaa09d9c151, []int{2}
}

func (m *VserviceNshSpiSiTerminate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VserviceNshSpiSiTerminate.Unmarshal(m, b)
}
func (m *VserviceNshSpiSiTerminate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VserviceNshSpiSiTerminate.Marshal(b, m, deterministic)
}
func (m *VserviceNshSpiSiTerminate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VserviceNshSpiSiTerminate.Merge(m, src)
}
func (m *VserviceNshSpiSiTerminate) XXX_Size() int {
	return xxx_messageInfo_VserviceNshSpiSiTerminate.Size(m)
}
func (m *VserviceNshSpiSiTerminate) XXX_DiscardUnknown() {
	xxx_messageInfo_VserviceNshSpiSiTerminate.DiscardUnknown(m)
}

var xxx_messageInfo_VserviceNshSpiSiTerminate proto.InternalMessageInfo

func (m *VserviceNshSpiSiTerminate) GetTerminatedPkts() uint64 {
	if m != nil {
		return m.TerminatedPkts
	}
	return 0
}

func (m *VserviceNshSpiSiTerminate) GetTerminatedBytes() uint64 {
	if m != nil {
		return m.TerminatedBytes
	}
	return 0
}

type VserviceNshSfp struct {
	SpiSi                *VserviceNshSpiSi          `protobuf:"bytes,1,opt,name=spi_si,json=spiSi,proto3" json:"spi_si,omitempty"`
	Term                 *VserviceNshSpiSiTerminate `protobuf:"bytes,2,opt,name=term,proto3" json:"term,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *VserviceNshSfp) Reset()         { *m = VserviceNshSfp{} }
func (m *VserviceNshSfp) String() string { return proto.CompactTextString(m) }
func (*VserviceNshSfp) ProtoMessage()    {}
func (*VserviceNshSfp) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6caaa09d9c151, []int{3}
}

func (m *VserviceNshSfp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VserviceNshSfp.Unmarshal(m, b)
}
func (m *VserviceNshSfp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VserviceNshSfp.Marshal(b, m, deterministic)
}
func (m *VserviceNshSfp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VserviceNshSfp.Merge(m, src)
}
func (m *VserviceNshSfp) XXX_Size() int {
	return xxx_messageInfo_VserviceNshSfp.Size(m)
}
func (m *VserviceNshSfp) XXX_DiscardUnknown() {
	xxx_messageInfo_VserviceNshSfp.DiscardUnknown(m)
}

var xxx_messageInfo_VserviceNshSfp proto.InternalMessageInfo

func (m *VserviceNshSfp) GetSpiSi() *VserviceNshSpiSi {
	if m != nil {
		return m.SpiSi
	}
	return nil
}

func (m *VserviceNshSfp) GetTerm() *VserviceNshSpiSiTerminate {
	if m != nil {
		return m.Term
	}
	return nil
}

type VserviceNshSf struct {
	ProcessedPkts        uint64   `protobuf:"varint,1,opt,name=processed_pkts,json=processedPkts,proto3" json:"processed_pkts,omitempty"`
	ProcessedBytes       uint64   `protobuf:"varint,2,opt,name=processed_bytes,json=processedBytes,proto3" json:"processed_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VserviceNshSf) Reset()         { *m = VserviceNshSf{} }
func (m *VserviceNshSf) String() string { return proto.CompactTextString(m) }
func (*VserviceNshSf) ProtoMessage()    {}
func (*VserviceNshSf) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6caaa09d9c151, []int{4}
}

func (m *VserviceNshSf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VserviceNshSf.Unmarshal(m, b)
}
func (m *VserviceNshSf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VserviceNshSf.Marshal(b, m, deterministic)
}
func (m *VserviceNshSf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VserviceNshSf.Merge(m, src)
}
func (m *VserviceNshSf) XXX_Size() int {
	return xxx_messageInfo_VserviceNshSf.Size(m)
}
func (m *VserviceNshSf) XXX_DiscardUnknown() {
	xxx_messageInfo_VserviceNshSf.DiscardUnknown(m)
}

var xxx_messageInfo_VserviceNshSf proto.InternalMessageInfo

func (m *VserviceNshSf) GetProcessedPkts() uint64 {
	if m != nil {
		return m.ProcessedPkts
	}
	return 0
}

func (m *VserviceNshSf) GetProcessedBytes() uint64 {
	if m != nil {
		return m.ProcessedBytes
	}
	return 0
}

type VserviceNshSffLocal struct {
	MalformedErrPkts     uint64   `protobuf:"varint,1,opt,name=malformed_err_pkts,json=malformedErrPkts,proto3" json:"malformed_err_pkts,omitempty"`
	LookupErrPkts        uint64   `protobuf:"varint,2,opt,name=lookup_err_pkts,json=lookupErrPkts,proto3" json:"lookup_err_pkts,omitempty"`
	MalformedErrBytes    uint64   `protobuf:"varint,3,opt,name=malformed_err_bytes,json=malformedErrBytes,proto3" json:"malformed_err_bytes,omitempty"`
	LookupErrBytes       uint64   `protobuf:"varint,4,opt,name=lookup_err_bytes,json=lookupErrBytes,proto3" json:"lookup_err_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VserviceNshSffLocal) Reset()         { *m = VserviceNshSffLocal{} }
func (m *VserviceNshSffLocal) String() string { return proto.CompactTextString(m) }
func (*VserviceNshSffLocal) ProtoMessage()    {}
func (*VserviceNshSffLocal) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6caaa09d9c151, []int{5}
}

func (m *VserviceNshSffLocal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VserviceNshSffLocal.Unmarshal(m, b)
}
func (m *VserviceNshSffLocal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VserviceNshSffLocal.Marshal(b, m, deterministic)
}
func (m *VserviceNshSffLocal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VserviceNshSffLocal.Merge(m, src)
}
func (m *VserviceNshSffLocal) XXX_Size() int {
	return xxx_messageInfo_VserviceNshSffLocal.Size(m)
}
func (m *VserviceNshSffLocal) XXX_DiscardUnknown() {
	xxx_messageInfo_VserviceNshSffLocal.DiscardUnknown(m)
}

var xxx_messageInfo_VserviceNshSffLocal proto.InternalMessageInfo

func (m *VserviceNshSffLocal) GetMalformedErrPkts() uint64 {
	if m != nil {
		return m.MalformedErrPkts
	}
	return 0
}

func (m *VserviceNshSffLocal) GetLookupErrPkts() uint64 {
	if m != nil {
		return m.LookupErrPkts
	}
	return 0
}

func (m *VserviceNshSffLocal) GetMalformedErrBytes() uint64 {
	if m != nil {
		return m.MalformedErrBytes
	}
	return 0
}

func (m *VserviceNshSffLocal) GetLookupErrBytes() uint64 {
	if m != nil {
		return m.LookupErrBytes
	}
	return 0
}

type VsNshStatsData struct {
	Type                 string                     `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Sfp                  *VserviceNshSfp            `protobuf:"bytes,2,opt,name=sfp,proto3" json:"sfp,omitempty"`
	SpiSi                *VserviceNshSpiSi          `protobuf:"bytes,3,opt,name=spi_si,json=spiSi,proto3" json:"spi_si,omitempty"`
	Term                 *VserviceNshSpiSiTerminate `protobuf:"bytes,4,opt,name=term,proto3" json:"term,omitempty"`
	Sf                   *VserviceNshSf             `protobuf:"bytes,5,opt,name=sf,proto3" json:"sf,omitempty"`
	Sff                  *VserviceNshSf             `protobuf:"bytes,6,opt,name=sff,proto3" json:"sff,omitempty"`
	SffLocal             *VserviceNshSffLocal       `protobuf:"bytes,7,opt,name=sff_local,json=sffLocal,proto3" json:"sff_local,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *VsNshStatsData) Reset()         { *m = VsNshStatsData{} }
func (m *VsNshStatsData) String() string { return proto.CompactTextString(m) }
func (*VsNshStatsData) ProtoMessage()    {}
func (*VsNshStatsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6caaa09d9c151, []int{6}
}

func (m *VsNshStatsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VsNshStatsData.Unmarshal(m, b)
}
func (m *VsNshStatsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VsNshStatsData.Marshal(b, m, deterministic)
}
func (m *VsNshStatsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VsNshStatsData.Merge(m, src)
}
func (m *VsNshStatsData) XXX_Size() int {
	return xxx_messageInfo_VsNshStatsData.Size(m)
}
func (m *VsNshStatsData) XXX_DiscardUnknown() {
	xxx_messageInfo_VsNshStatsData.DiscardUnknown(m)
}

var xxx_messageInfo_VsNshStatsData proto.InternalMessageInfo

func (m *VsNshStatsData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VsNshStatsData) GetSfp() *VserviceNshSfp {
	if m != nil {
		return m.Sfp
	}
	return nil
}

func (m *VsNshStatsData) GetSpiSi() *VserviceNshSpiSi {
	if m != nil {
		return m.SpiSi
	}
	return nil
}

func (m *VsNshStatsData) GetTerm() *VserviceNshSpiSiTerminate {
	if m != nil {
		return m.Term
	}
	return nil
}

func (m *VsNshStatsData) GetSf() *VserviceNshSf {
	if m != nil {
		return m.Sf
	}
	return nil
}

func (m *VsNshStatsData) GetSff() *VserviceNshSf {
	if m != nil {
		return m.Sff
	}
	return nil
}

func (m *VsNshStatsData) GetSffLocal() *VserviceNshSffLocal {
	if m != nil {
		return m.SffLocal
	}
	return nil
}

type VsNshSiData struct {
	Type                 string                     `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	SpiSi                *VserviceNshSpiSi          `protobuf:"bytes,2,opt,name=spi_si,json=spiSi,proto3" json:"spi_si,omitempty"`
	Term                 *VserviceNshSpiSiTerminate `protobuf:"bytes,3,opt,name=term,proto3" json:"term,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *VsNshSiData) Reset()         { *m = VsNshSiData{} }
func (m *VsNshSiData) String() string { return proto.CompactTextString(m) }
func (*VsNshSiData) ProtoMessage()    {}
func (*VsNshSiData) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6caaa09d9c151, []int{7}
}

func (m *VsNshSiData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VsNshSiData.Unmarshal(m, b)
}
func (m *VsNshSiData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VsNshSiData.Marshal(b, m, deterministic)
}
func (m *VsNshSiData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VsNshSiData.Merge(m, src)
}
func (m *VsNshSiData) XXX_Size() int {
	return xxx_messageInfo_VsNshSiData.Size(m)
}
func (m *VsNshSiData) XXX_DiscardUnknown() {
	xxx_messageInfo_VsNshSiData.DiscardUnknown(m)
}

var xxx_messageInfo_VsNshSiData proto.InternalMessageInfo

func (m *VsNshSiData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VsNshSiData) GetSpiSi() *VserviceNshSpiSi {
	if m != nil {
		return m.SpiSi
	}
	return nil
}

func (m *VsNshSiData) GetTerm() *VserviceNshSpiSiTerminate {
	if m != nil {
		return m.Term
	}
	return nil
}

type VserviceNshSi struct {
	Si                   uint32       `protobuf:"varint,1,opt,name=si,proto3" json:"si,omitempty"`
	Data                 *VsNshSiData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *VserviceNshSi) Reset()         { *m = VserviceNshSi{} }
func (m *VserviceNshSi) String() string { return proto.CompactTextString(m) }
func (*VserviceNshSi) ProtoMessage()    {}
func (*VserviceNshSi) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6caaa09d9c151, []int{8}
}

func (m *VserviceNshSi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VserviceNshSi.Unmarshal(m, b)
}
func (m *VserviceNshSi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VserviceNshSi.Marshal(b, m, deterministic)
}
func (m *VserviceNshSi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VserviceNshSi.Merge(m, src)
}
func (m *VserviceNshSi) XXX_Size() int {
	return xxx_messageInfo_VserviceNshSi.Size(m)
}
func (m *VserviceNshSi) XXX_DiscardUnknown() {
	xxx_messageInfo_VserviceNshSi.DiscardUnknown(m)
}

var xxx_messageInfo_VserviceNshSi proto.InternalMessageInfo

func (m *VserviceNshSi) GetSi() uint32 {
	if m != nil {
		return m.Si
	}
	return 0
}

func (m *VserviceNshSi) GetData() *VsNshSiData {
	if m != nil {
		return m.Data
	}
	return nil
}

type VserviceEdmNshStatsInfo struct {
	Data                 *VsNshStatsData  `protobuf:"bytes,50,opt,name=data,proto3" json:"data,omitempty"`
	SiArr                []*VserviceNshSi `protobuf:"bytes,51,rep,name=si_arr,json=siArr,proto3" json:"si_arr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *VserviceEdmNshStatsInfo) Reset()         { *m = VserviceEdmNshStatsInfo{} }
func (m *VserviceEdmNshStatsInfo) String() string { return proto.CompactTextString(m) }
func (*VserviceEdmNshStatsInfo) ProtoMessage()    {}
func (*VserviceEdmNshStatsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6caaa09d9c151, []int{9}
}

func (m *VserviceEdmNshStatsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VserviceEdmNshStatsInfo.Unmarshal(m, b)
}
func (m *VserviceEdmNshStatsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VserviceEdmNshStatsInfo.Marshal(b, m, deterministic)
}
func (m *VserviceEdmNshStatsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VserviceEdmNshStatsInfo.Merge(m, src)
}
func (m *VserviceEdmNshStatsInfo) XXX_Size() int {
	return xxx_messageInfo_VserviceEdmNshStatsInfo.Size(m)
}
func (m *VserviceEdmNshStatsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VserviceEdmNshStatsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VserviceEdmNshStatsInfo proto.InternalMessageInfo

func (m *VserviceEdmNshStatsInfo) GetData() *VsNshStatsData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *VserviceEdmNshStatsInfo) GetSiArr() []*VserviceNshSi {
	if m != nil {
		return m.SiArr
	}
	return nil
}

func init() {
	proto.RegisterType((*VserviceEdmNshStatsInfo_KEYS)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function_path.path_ids.path_id.stats.detail.vservice_edm_nsh_stats_info_KEYS")
	proto.RegisterType((*VserviceNshSpiSi)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function_path.path_ids.path_id.stats.detail.vservice_nsh_spi_si")
	proto.RegisterType((*VserviceNshSpiSiTerminate)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function_path.path_ids.path_id.stats.detail.vservice_nsh_spi_si_terminate")
	proto.RegisterType((*VserviceNshSfp)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function_path.path_ids.path_id.stats.detail.vservice_nsh_sfp")
	proto.RegisterType((*VserviceNshSf)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function_path.path_ids.path_id.stats.detail.vservice_nsh_sf")
	proto.RegisterType((*VserviceNshSffLocal)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function_path.path_ids.path_id.stats.detail.vservice_nsh_sff_local")
	proto.RegisterType((*VsNshStatsData)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function_path.path_ids.path_id.stats.detail.vs_nsh_stats_data")
	proto.RegisterType((*VsNshSiData)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function_path.path_ids.path_id.stats.detail.vs_nsh_si_data")
	proto.RegisterType((*VserviceNshSi)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function_path.path_ids.path_id.stats.detail.vservice_nsh_si")
	proto.RegisterType((*VserviceEdmNshStatsInfo)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function_path.path_ids.path_id.stats.detail.vservice_edm_nsh_stats_info")
}

func init() { proto.RegisterFile("vservice_edm_nsh_stats_info.proto", fileDescriptor_50a6caaa09d9c151) }

var fileDescriptor_50a6caaa09d9c151 = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0x66, 0x7e, 0x36, 0xee, 0xd6, 0x92, 0x9f, 0xed, 0x05, 0x19, 0x10, 0x21, 0x0e, 0xa8, 0x2b,
	0xc8, 0x1c, 0xb2, 0x4f, 0xa0, 0xb0, 0x27, 0x3d, 0xc8, 0xec, 0xc9, 0x53, 0xd3, 0x99, 0xe9, 0x4e,
	0x9a, 0x4c, 0xa6, 0x9b, 0xee, 0xde, 0xe8, 0x3e, 0x80, 0xf8, 0x83, 0xe0, 0x4d, 0x10, 0x05, 0x3d,
	0x78, 0x13, 0xdf, 0xc1, 0xa3, 0xcf, 0xe1, 0xc5, 0x67, 0xf0, 0x26, 0xd3, 0x93, 0x4c, 0x26, 0x71,
	0xdd, 0xd3, 0x42, 0x26, 0x97, 0xd0, 0xa9, 0xf9, 0xba, 0xbe, 0xaf, 0xba, 0xab, 0xaa, 0x0b, 0x6e,
	0xcd, 0x34, 0x55, 0x33, 0x9e, 0x50, 0x4c, 0xd3, 0x29, 0xce, 0xf5, 0x18, 0x6b, 0x43, 0x8c, 0xc6,
	0x3c, 0x67, 0x22, 0x92, 0x4a, 0x18, 0x81, 0x9e, 0x25, 0x5c, 0x27, 0x02, 0x73, 0xa1, 0xf1, 0x73,
	0x85, 0xe5, 0x50, 0xe1, 0x6a, 0xcf, 0x74, 0xa4, 0xb0, 0x90, 0x54, 0x45, 0xa3, 0x4c, 0x0c, 0x49,
	0x86, 0x17, 0x76, 0x76, 0x96, 0x27, 0x86, 0x8b, 0x1c, 0x27, 0x63, 0xc2, 0x73, 0x9e, 0x8f, 0xa2,
	0x7f, 0xbe, 0x48, 0x62, 0xc6, 0x51, 0xf1, 0x83, 0x79, 0xaa, 0x17, 0x8b, 0xc8, 0x92, 0x47, 0x29,
	0x35, 0x84, 0x67, 0xe1, 0x00, 0xfa, 0x97, 0xa8, 0xc3, 0x8f, 0x4e, 0x9e, 0x9e, 0xa2, 0x0e, 0xb8,
	0x3c, 0x0d, 0x9c, 0xbe, 0x73, 0xd4, 0x8e, 0x5d, 0x9e, 0x86, 0x14, 0x0e, 0xab, 0x3d, 0x16, 0x2f,
	0x39, 0xd6, 0x1c, 0xdd, 0x86, 0x8e, 0x54, 0x22, 0xa1, 0x5a, 0xd3, 0x14, 0xcb, 0x89, 0xd1, 0x76,
	0x8b, 0x1f, 0xb7, 0x2b, 0xeb, 0x93, 0x89, 0xd1, 0xe8, 0x2e, 0x74, 0x97, 0xb0, 0xe1, 0xb9, 0xa1,
	0x3a, 0x70, 0x2d, 0x6e, 0xb9, 0xfb, 0x61, 0x61, 0x0d, 0x35, 0xdc, 0xbc, 0x80, 0x06, 0x1b, 0xaa,
	0xa6, 0x3c, 0x27, 0x86, 0x16, 0x9e, 0xaa, 0x3f, 0x2b, 0x8c, 0x9d, 0xa5, 0xd9, 0x52, 0xde, 0x83,
	0x5e, 0x0d, 0x58, 0xe7, 0xac, 0x39, 0x28, 0x49, 0x7f, 0xbb, 0xd0, 0x5b, 0x65, 0x65, 0x12, 0x7d,
	0x75, 0xa0, 0x55, 0xb2, 0x5b, 0x82, 0xfd, 0xc1, 0x5b, 0x27, 0xda, 0xd0, 0x85, 0x45, 0x17, 0x1c,
	0x49, 0xbc, 0xa3, 0x25, 0x3f, 0xe5, 0xe8, 0x9b, 0x03, 0x7e, 0x11, 0x8f, 0x8d, 0x6d, 0x7f, 0xf0,
	0xbe, 0x51, 0x22, 0x97, 0xf7, 0x16, 0x5b, 0x91, 0x21, 0x81, 0xee, 0xda, 0x41, 0x5f, 0x79, 0x06,
	0xfd, 0x74, 0xe0, 0xfa, 0x1a, 0x07, 0xc3, 0x99, 0x48, 0x48, 0x86, 0xee, 0x03, 0x9a, 0x92, 0x8c,
	0x09, 0x35, 0xa5, 0x29, 0xa6, 0x4a, 0xd5, 0xe9, 0x7a, 0xd5, 0x97, 0x13, 0xa5, 0x2c, 0xe3, 0x1d,
	0xe8, 0x66, 0x42, 0x4c, 0xce, 0xe4, 0x12, 0x5a, 0x32, 0xb6, 0x4b, 0xf3, 0x02, 0x17, 0xc1, 0xe1,
	0xaa, 0xd7, 0x52, 0x9d, 0x67, 0xb1, 0x07, 0x75, 0xb7, 0x56, 0x20, 0x3a, 0x82, 0x5e, 0xcd, 0x6f,
	0x09, 0xf6, 0xcb, 0x50, 0x2a, 0xc7, 0x65, 0x28, 0x2f, 0x76, 0xe1, 0x60, 0xa6, 0x6b, 0xe5, 0x99,
	0x12, 0x43, 0x10, 0x02, 0xdf, 0x9c, 0x4b, 0x6a, 0x75, 0xef, 0xc5, 0x76, 0x8d, 0x3e, 0x39, 0xe0,
	0x69, 0x26, 0xe7, 0x49, 0xf0, 0xba, 0x29, 0x49, 0xc0, 0x64, 0xec, 0xad, 0xd5, 0x92, 0xb7, 0x0d,
	0xb5, 0xe4, 0x6f, 0x41, 0x2d, 0xa1, 0x0f, 0x0e, 0xb8, 0x9a, 0x05, 0x3b, 0x56, 0xeb, 0xab, 0xc6,
	0x5c, 0x79, 0xec, 0x6a, 0x86, 0x3e, 0xda, 0x7c, 0x64, 0x41, 0xab, 0x69, 0xe2, 0x0a, 0x55, 0xe8,
	0xbb, 0x03, 0x7b, 0x55, 0x57, 0x08, 0xae, 0x59, 0x8d, 0xef, 0x1a, 0xa3, 0x71, 0xae, 0x2b, 0xde,
	0xd5, 0x8c, 0x3d, 0x2e, 0x56, 0xe1, 0x1f, 0x17, 0x3a, 0x8b, 0x3e, 0xc0, 0xff, 0xdf, 0x04, 0x6a,
	0x55, 0xe6, 0x6e, 0x43, 0x95, 0x79, 0xdb, 0xf0, 0x62, 0xfd, 0x70, 0xd6, 0x9f, 0x2c, 0x5e, 0xcc,
	0x46, 0xf3, 0xa9, 0xa0, 0x1d, 0xbb, 0x9a, 0x17, 0xdd, 0xd7, 0x2f, 0x6e, 0x65, 0x7e, 0xec, 0x2f,
	0x37, 0x19, 0x51, 0x3d, 0x4b, 0x62, 0xab, 0x2a, 0xfc, 0xe5, 0xc2, 0x8d, 0x4b, 0xe6, 0x3d, 0xf4,
	0x79, 0x21, 0x7f, 0x60, 0xe5, 0xbf, 0xd9, 0xbc, 0xfc, 0xea, 0xb1, 0x2b, 0x23, 0x40, 0x5f, 0x8a,
	0xcc, 0xe6, 0x98, 0x28, 0x15, 0x1c, 0xf7, 0xbd, 0x26, 0x75, 0x94, 0x22, 0xab, 0xf9, 0x03, 0xa5,
	0x86, 0x2d, 0x3b, 0xd3, 0x1f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x37, 0x97, 0x3b, 0xf8,
	0x0b, 0x00, 0x00,
}
