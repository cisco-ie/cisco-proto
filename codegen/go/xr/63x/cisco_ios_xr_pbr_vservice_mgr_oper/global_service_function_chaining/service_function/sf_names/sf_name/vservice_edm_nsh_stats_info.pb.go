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

package cisco_ios_xr_pbr_vservice_mgr_oper_global_service_function_chaining_service_function_sf_names_sf_name

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
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
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

func (m *VserviceEdmNshStatsInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
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
	proto.RegisterType((*VserviceEdmNshStatsInfo_KEYS)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function.sf_names.sf_name.vservice_edm_nsh_stats_info_KEYS")
	proto.RegisterType((*VserviceNshSpiSi)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function.sf_names.sf_name.vservice_nsh_spi_si")
	proto.RegisterType((*VserviceNshSpiSiTerminate)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function.sf_names.sf_name.vservice_nsh_spi_si_terminate")
	proto.RegisterType((*VserviceNshSfp)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function.sf_names.sf_name.vservice_nsh_sfp")
	proto.RegisterType((*VserviceNshSf)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function.sf_names.sf_name.vservice_nsh_sf")
	proto.RegisterType((*VserviceNshSffLocal)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function.sf_names.sf_name.vservice_nsh_sff_local")
	proto.RegisterType((*VsNshStatsData)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function.sf_names.sf_name.vs_nsh_stats_data")
	proto.RegisterType((*VsNshSiData)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function.sf_names.sf_name.vs_nsh_si_data")
	proto.RegisterType((*VserviceNshSi)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function.sf_names.sf_name.vservice_nsh_si")
	proto.RegisterType((*VserviceEdmNshStatsInfo)(nil), "cisco_ios_xr_pbr_vservice_mgr_oper.global_service_function_chaining.service_function.sf_names.sf_name.vservice_edm_nsh_stats_info")
}

func init() { proto.RegisterFile("vservice_edm_nsh_stats_info.proto", fileDescriptor_50a6caaa09d9c151) }

var fileDescriptor_50a6caaa09d9c151 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0xc9, 0x8f, 0xdd, 0x75, 0xdf, 0xb2, 0x6d, 0x77, 0x16, 0x24, 0x20, 0x42, 0x0d, 0xa8,
	0x2b, 0x48, 0x0e, 0x5d, 0xf0, 0xae, 0xb0, 0x27, 0x3d, 0x48, 0xf6, 0xe4, 0x69, 0x48, 0xd3, 0x99,
	0xee, 0xd0, 0x24, 0x33, 0xcc, 0x4b, 0xeb, 0xae, 0xe0, 0x4d, 0x44, 0x41, 0x0f, 0x82, 0x17, 0xfd,
	0x1f, 0xfc, 0x0f, 0x3c, 0x79, 0xf2, 0x0f, 0xf0, 0x0f, 0x92, 0x99, 0xb4, 0x69, 0x5a, 0xd6, 0x3d,
	0x49, 0xd3, 0xdb, 0xf4, 0xe5, 0xbd, 0xf7, 0xfd, 0x4c, 0xde, 0x7c, 0x3b, 0x81, 0x7b, 0x33, 0x64,
	0x7a, 0x26, 0x52, 0x46, 0xd9, 0x28, 0xa7, 0x05, 0x5e, 0x50, 0x2c, 0x93, 0x12, 0xa9, 0x28, 0xb8,
	0x8c, 0x94, 0x96, 0xa5, 0x24, 0x2c, 0x15, 0x98, 0x4a, 0x2a, 0x24, 0xd2, 0x4b, 0x4d, 0xd5, 0x50,
	0xd3, 0xba, 0x26, 0x1f, 0x6b, 0x2a, 0x15, 0xd3, 0xd1, 0x38, 0x93, 0xc3, 0x24, 0xa3, 0x8b, 0x38,
	0x9f, 0x16, 0x69, 0x29, 0x64, 0x41, 0xd3, 0x8b, 0x44, 0x14, 0xa2, 0x18, 0x47, 0xeb, 0x4f, 0x22,
	0xe4, 0xb4, 0x48, 0x72, 0x86, 0x8b, 0x45, 0xf8, 0x04, 0xfa, 0x37, 0xb0, 0xd0, 0xe7, 0x67, 0xaf,
	0xce, 0x09, 0x01, 0xdf, 0xe4, 0x06, 0x4e, 0xdf, 0x39, 0xd9, 0x8f, 0xed, 0x3a, 0x64, 0x70, 0x5c,
	0xd7, 0xd9, 0x1a, 0x25, 0x28, 0x0a, 0x72, 0x1f, 0x3a, 0x4a, 0xcb, 0x94, 0x21, 0xb2, 0x11, 0x55,
	0x93, 0x12, 0x6d, 0x91, 0x1f, 0x1f, 0xd6, 0xd1, 0x97, 0x93, 0x12, 0xc9, 0x43, 0xe8, 0x2e, 0xd3,
	0x86, 0x57, 0x25, 0xc3, 0xc0, 0xb5, 0x79, 0xcb, 0xea, 0x67, 0x26, 0x1a, 0x22, 0xdc, 0xbd, 0x46,
	0x86, 0x96, 0x4c, 0xe7, 0xa2, 0x48, 0x4a, 0x66, 0x3a, 0xd5, 0x3f, 0x56, 0x14, 0x3b, 0xcb, 0xb0,
	0x95, 0x7c, 0x04, 0xbd, 0x46, 0x62, 0x53, 0xb3, 0xd1, 0xa0, 0x12, 0xfd, 0xe5, 0x42, 0x6f, 0x55,
	0x95, 0x2b, 0xf2, 0xc5, 0x81, 0xdd, 0x4a, 0xdd, 0x0a, 0x1c, 0x0c, 0xde, 0x44, 0x1b, 0x99, 0x50,
	0x74, 0xcd, 0xfe, 0xe3, 0x1d, 0x54, 0xe2, 0x5c, 0x90, 0x6f, 0x0e, 0xf8, 0x06, 0xde, 0x6e, 0xe4,
	0x60, 0xf0, 0xce, 0x69, 0x0f, 0x69, 0x39, 0x92, 0xd8, 0x22, 0x85, 0x09, 0x74, 0xd7, 0xde, 0xe1,
	0x7f, 0x3f, 0x1c, 0xbf, 0x1d, 0xb8, 0xbd, 0xa6, 0xc1, 0x69, 0x26, 0xd3, 0x24, 0x23, 0x8f, 0x81,
	0xe4, 0x49, 0xc6, 0xa5, 0xce, 0xd9, 0x88, 0x32, 0xad, 0x9b, 0x72, 0xbd, 0xfa, 0xc9, 0x99, 0xd6,
	0x56, 0xf1, 0x01, 0x74, 0x33, 0x29, 0x27, 0x53, 0xb5, 0x4c, 0xad, 0x14, 0x0f, 0xab, 0xf0, 0x22,
	0x2f, 0x82, 0xe3, 0xd5, 0xae, 0x15, 0x9d, 0x67, 0x73, 0x8f, 0x9a, 0x6d, 0x2d, 0x20, 0x39, 0x81,
	0x5e, 0xa3, 0x6f, 0x95, 0xec, 0x57, 0x5b, 0xa9, 0x1b, 0x57, 0x5b, 0xf9, 0xba, 0x07, 0x47, 0x33,
	0x6c, 0xb8, 0x6f, 0x94, 0x94, 0x89, 0x31, 0x5e, 0x79, 0xa5, 0x6a, 0xe3, 0x99, 0x35, 0xf9, 0xe8,
	0x80, 0x87, 0x5c, 0xcd, 0x47, 0xfe, 0xba, 0x95, 0x89, 0x73, 0x15, 0x7b, 0x6b, 0x9e, 0xf0, 0xb6,
	0xce, 0x13, 0xfe, 0xd6, 0x79, 0x82, 0xbc, 0x77, 0xc0, 0x45, 0x1e, 0xec, 0x58, 0xb2, 0x59, 0x3b,
	0xa3, 0x8b, 0x5d, 0xe4, 0xe4, 0x83, 0x3d, 0x44, 0x3c, 0xd8, 0x6d, 0x95, 0xc4, 0x20, 0x90, 0xef,
	0x0e, 0xec, 0xd7, 0xbe, 0x0d, 0xf6, 0x2c, 0xd0, 0xdb, 0x76, 0x80, 0xe6, 0x10, 0xf1, 0x2d, 0xe4,
	0xfc, 0x85, 0x59, 0x85, 0x7f, 0x5c, 0xe8, 0x2c, 0x6c, 0x29, 0xfe, 0xed, 0xc9, 0x86, 0x0f, 0xdc,
	0xad, 0xf3, 0x81, 0xb7, 0x7d, 0x77, 0xc3, 0x0f, 0x67, 0xfd, 0x72, 0x10, 0xa4, 0x03, 0xee, 0xfc,
	0x6a, 0x3d, 0x8c, 0x5d, 0x14, 0xe6, 0x7f, 0xce, 0x37, 0x2f, 0x7c, 0xfe, 0x46, 0xa7, 0x1b, 0xc3,
	0x6f, 0x4e, 0x3b, 0xb6, 0x08, 0xe1, 0x4f, 0x17, 0xee, 0xdc, 0xf0, 0x95, 0x44, 0x3e, 0x2d, 0x58,
	0x07, 0x96, 0xf5, 0x72, 0xc3, 0xac, 0xf5, 0x85, 0x51, 0xe1, 0x92, 0xcf, 0xe6, 0x38, 0x0a, 0x9a,
	0x68, 0x1d, 0x9c, 0xf6, 0xbd, 0xd6, 0x0c, 0x6e, 0x8e, 0xa2, 0x78, 0xaa, 0xf5, 0x70, 0xd7, 0x7e,
	0xd1, 0x9e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x18, 0xbb, 0x89, 0xe0, 0xf6, 0x0a, 0x00, 0x00,
}
