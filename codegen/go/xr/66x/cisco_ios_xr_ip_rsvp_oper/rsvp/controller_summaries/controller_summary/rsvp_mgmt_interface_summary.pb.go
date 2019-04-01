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
// source: rsvp_mgmt_interface_summary.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_controller_summaries_controller_summary

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

type RsvpMgmtInterfaceSummary_KEYS struct {
	ControllerName       string   `protobuf:"bytes,1,opt,name=controller_name,json=controllerName,proto3" json:"controller_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtInterfaceSummary_KEYS) Reset()         { *m = RsvpMgmtInterfaceSummary_KEYS{} }
func (m *RsvpMgmtInterfaceSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtInterfaceSummary_KEYS) ProtoMessage()    {}
func (*RsvpMgmtInterfaceSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1afdc8eaec0c32ba, []int{0}
}

func (m *RsvpMgmtInterfaceSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtInterfaceSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtInterfaceSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS.Merge(m, src)
}
func (m *RsvpMgmtInterfaceSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS.Size(m)
}
func (m *RsvpMgmtInterfaceSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtInterfaceSummary_KEYS) GetControllerName() string {
	if m != nil {
		return m.ControllerName
	}
	return ""
}

type RsvpMgmtInterfaceBwPrestdDste struct {
	AllocatedBitRate              uint64   `protobuf:"varint,1,opt,name=allocated_bit_rate,json=allocatedBitRate,proto3" json:"allocated_bit_rate,omitempty"`
	MaxFlowBandwidth              uint64   `protobuf:"varint,2,opt,name=max_flow_bandwidth,json=maxFlowBandwidth,proto3" json:"max_flow_bandwidth,omitempty"`
	MaxBandwidth                  uint64   `protobuf:"varint,3,opt,name=max_bandwidth,json=maxBandwidth,proto3" json:"max_bandwidth,omitempty"`
	MaxSubpoolBandwidth           uint64   `protobuf:"varint,4,opt,name=max_subpool_bandwidth,json=maxSubpoolBandwidth,proto3" json:"max_subpool_bandwidth,omitempty"`
	IsMaxBandwidthAbsolute        bool     `protobuf:"varint,5,opt,name=is_max_bandwidth_absolute,json=isMaxBandwidthAbsolute,proto3" json:"is_max_bandwidth_absolute,omitempty"`
	IsMaxSubpoolBandwidthAbsolute bool     `protobuf:"varint,6,opt,name=is_max_subpool_bandwidth_absolute,json=isMaxSubpoolBandwidthAbsolute,proto3" json:"is_max_subpool_bandwidth_absolute,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *RsvpMgmtInterfaceBwPrestdDste) Reset()         { *m = RsvpMgmtInterfaceBwPrestdDste{} }
func (m *RsvpMgmtInterfaceBwPrestdDste) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtInterfaceBwPrestdDste) ProtoMessage()    {}
func (*RsvpMgmtInterfaceBwPrestdDste) Descriptor() ([]byte, []int) {
	return fileDescriptor_1afdc8eaec0c32ba, []int{1}
}

func (m *RsvpMgmtInterfaceBwPrestdDste) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste.Unmarshal(m, b)
}
func (m *RsvpMgmtInterfaceBwPrestdDste) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtInterfaceBwPrestdDste) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste.Merge(m, src)
}
func (m *RsvpMgmtInterfaceBwPrestdDste) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste.Size(m)
}
func (m *RsvpMgmtInterfaceBwPrestdDste) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste proto.InternalMessageInfo

func (m *RsvpMgmtInterfaceBwPrestdDste) GetAllocatedBitRate() uint64 {
	if m != nil {
		return m.AllocatedBitRate
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwPrestdDste) GetMaxFlowBandwidth() uint64 {
	if m != nil {
		return m.MaxFlowBandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwPrestdDste) GetMaxBandwidth() uint64 {
	if m != nil {
		return m.MaxBandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwPrestdDste) GetMaxSubpoolBandwidth() uint64 {
	if m != nil {
		return m.MaxSubpoolBandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwPrestdDste) GetIsMaxBandwidthAbsolute() bool {
	if m != nil {
		return m.IsMaxBandwidthAbsolute
	}
	return false
}

func (m *RsvpMgmtInterfaceBwPrestdDste) GetIsMaxSubpoolBandwidthAbsolute() bool {
	if m != nil {
		return m.IsMaxSubpoolBandwidthAbsolute
	}
	return false
}

type RsvpMgmtInterfaceBwStdDste struct {
	AllocatedBitRate          uint64   `protobuf:"varint,1,opt,name=allocated_bit_rate,json=allocatedBitRate,proto3" json:"allocated_bit_rate,omitempty"`
	MaxFlowBandwidth          uint64   `protobuf:"varint,2,opt,name=max_flow_bandwidth,json=maxFlowBandwidth,proto3" json:"max_flow_bandwidth,omitempty"`
	MaxBandwidth              uint64   `protobuf:"varint,3,opt,name=max_bandwidth,json=maxBandwidth,proto3" json:"max_bandwidth,omitempty"`
	MaxPool0Bandwidth         uint64   `protobuf:"varint,4,opt,name=max_pool0_bandwidth,json=maxPool0Bandwidth,proto3" json:"max_pool0_bandwidth,omitempty"`
	MaxPool1Bandwidth         uint64   `protobuf:"varint,5,opt,name=max_pool1_bandwidth,json=maxPool1Bandwidth,proto3" json:"max_pool1_bandwidth,omitempty"`
	IsMaxBandwidthAbsolute    bool     `protobuf:"varint,6,opt,name=is_max_bandwidth_absolute,json=isMaxBandwidthAbsolute,proto3" json:"is_max_bandwidth_absolute,omitempty"`
	IsMaxBc0BandwidthAbsolute bool     `protobuf:"varint,7,opt,name=is_max_bc0_bandwidth_absolute,json=isMaxBc0BandwidthAbsolute,proto3" json:"is_max_bc0_bandwidth_absolute,omitempty"`
	IsMaxBc1BandwidthAbsolute bool     `protobuf:"varint,8,opt,name=is_max_bc1_bandwidth_absolute,json=isMaxBc1BandwidthAbsolute,proto3" json:"is_max_bc1_bandwidth_absolute,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *RsvpMgmtInterfaceBwStdDste) Reset()         { *m = RsvpMgmtInterfaceBwStdDste{} }
func (m *RsvpMgmtInterfaceBwStdDste) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtInterfaceBwStdDste) ProtoMessage()    {}
func (*RsvpMgmtInterfaceBwStdDste) Descriptor() ([]byte, []int) {
	return fileDescriptor_1afdc8eaec0c32ba, []int{2}
}

func (m *RsvpMgmtInterfaceBwStdDste) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtInterfaceBwStdDste.Unmarshal(m, b)
}
func (m *RsvpMgmtInterfaceBwStdDste) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtInterfaceBwStdDste.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtInterfaceBwStdDste) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtInterfaceBwStdDste.Merge(m, src)
}
func (m *RsvpMgmtInterfaceBwStdDste) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtInterfaceBwStdDste.Size(m)
}
func (m *RsvpMgmtInterfaceBwStdDste) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtInterfaceBwStdDste.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtInterfaceBwStdDste proto.InternalMessageInfo

func (m *RsvpMgmtInterfaceBwStdDste) GetAllocatedBitRate() uint64 {
	if m != nil {
		return m.AllocatedBitRate
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwStdDste) GetMaxFlowBandwidth() uint64 {
	if m != nil {
		return m.MaxFlowBandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwStdDste) GetMaxBandwidth() uint64 {
	if m != nil {
		return m.MaxBandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwStdDste) GetMaxPool0Bandwidth() uint64 {
	if m != nil {
		return m.MaxPool0Bandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwStdDste) GetMaxPool1Bandwidth() uint64 {
	if m != nil {
		return m.MaxPool1Bandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwStdDste) GetIsMaxBandwidthAbsolute() bool {
	if m != nil {
		return m.IsMaxBandwidthAbsolute
	}
	return false
}

func (m *RsvpMgmtInterfaceBwStdDste) GetIsMaxBc0BandwidthAbsolute() bool {
	if m != nil {
		return m.IsMaxBc0BandwidthAbsolute
	}
	return false
}

func (m *RsvpMgmtInterfaceBwStdDste) GetIsMaxBc1BandwidthAbsolute() bool {
	if m != nil {
		return m.IsMaxBc1BandwidthAbsolute
	}
	return false
}

type RsvpMgmtDsteModeInterfaceBw struct {
	DsteMode                 string                         `protobuf:"bytes,1,opt,name=dste_mode,json=dsteMode,proto3" json:"dste_mode,omitempty"`
	PreStandardDsteInterface *RsvpMgmtInterfaceBwPrestdDste `protobuf:"bytes,2,opt,name=pre_standard_dste_interface,json=preStandardDsteInterface,proto3" json:"pre_standard_dste_interface,omitempty"`
	StandardDsteInterface    *RsvpMgmtInterfaceBwStdDste    `protobuf:"bytes,3,opt,name=standard_dste_interface,json=standardDsteInterface,proto3" json:"standard_dste_interface,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                       `json:"-"`
	XXX_unrecognized         []byte                         `json:"-"`
	XXX_sizecache            int32                          `json:"-"`
}

func (m *RsvpMgmtDsteModeInterfaceBw) Reset()         { *m = RsvpMgmtDsteModeInterfaceBw{} }
func (m *RsvpMgmtDsteModeInterfaceBw) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtDsteModeInterfaceBw) ProtoMessage()    {}
func (*RsvpMgmtDsteModeInterfaceBw) Descriptor() ([]byte, []int) {
	return fileDescriptor_1afdc8eaec0c32ba, []int{3}
}

func (m *RsvpMgmtDsteModeInterfaceBw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw.Unmarshal(m, b)
}
func (m *RsvpMgmtDsteModeInterfaceBw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtDsteModeInterfaceBw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw.Merge(m, src)
}
func (m *RsvpMgmtDsteModeInterfaceBw) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw.Size(m)
}
func (m *RsvpMgmtDsteModeInterfaceBw) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw proto.InternalMessageInfo

func (m *RsvpMgmtDsteModeInterfaceBw) GetDsteMode() string {
	if m != nil {
		return m.DsteMode
	}
	return ""
}

func (m *RsvpMgmtDsteModeInterfaceBw) GetPreStandardDsteInterface() *RsvpMgmtInterfaceBwPrestdDste {
	if m != nil {
		return m.PreStandardDsteInterface
	}
	return nil
}

func (m *RsvpMgmtDsteModeInterfaceBw) GetStandardDsteInterface() *RsvpMgmtInterfaceBwStdDste {
	if m != nil {
		return m.StandardDsteInterface
	}
	return nil
}

type RsvpMgmtInterfaceSummary struct {
	InterfaceNameXr      string                       `protobuf:"bytes,50,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	BandwidthInformation *RsvpMgmtDsteModeInterfaceBw `protobuf:"bytes,51,opt,name=bandwidth_information,json=bandwidthInformation,proto3" json:"bandwidth_information,omitempty"`
	PathsIn              uint32                       `protobuf:"varint,52,opt,name=paths_in,json=pathsIn,proto3" json:"paths_in,omitempty"`
	PathsOut             uint32                       `protobuf:"varint,53,opt,name=paths_out,json=pathsOut,proto3" json:"paths_out,omitempty"`
	ReservationsIn       uint32                       `protobuf:"varint,54,opt,name=reservations_in,json=reservationsIn,proto3" json:"reservations_in,omitempty"`
	ReservationsOut      uint32                       `protobuf:"varint,55,opt,name=reservations_out,json=reservationsOut,proto3" json:"reservations_out,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RsvpMgmtInterfaceSummary) Reset()         { *m = RsvpMgmtInterfaceSummary{} }
func (m *RsvpMgmtInterfaceSummary) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtInterfaceSummary) ProtoMessage()    {}
func (*RsvpMgmtInterfaceSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_1afdc8eaec0c32ba, []int{4}
}

func (m *RsvpMgmtInterfaceSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary.Unmarshal(m, b)
}
func (m *RsvpMgmtInterfaceSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtInterfaceSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtInterfaceSummary.Merge(m, src)
}
func (m *RsvpMgmtInterfaceSummary) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary.Size(m)
}
func (m *RsvpMgmtInterfaceSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtInterfaceSummary.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtInterfaceSummary proto.InternalMessageInfo

func (m *RsvpMgmtInterfaceSummary) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *RsvpMgmtInterfaceSummary) GetBandwidthInformation() *RsvpMgmtDsteModeInterfaceBw {
	if m != nil {
		return m.BandwidthInformation
	}
	return nil
}

func (m *RsvpMgmtInterfaceSummary) GetPathsIn() uint32 {
	if m != nil {
		return m.PathsIn
	}
	return 0
}

func (m *RsvpMgmtInterfaceSummary) GetPathsOut() uint32 {
	if m != nil {
		return m.PathsOut
	}
	return 0
}

func (m *RsvpMgmtInterfaceSummary) GetReservationsIn() uint32 {
	if m != nil {
		return m.ReservationsIn
	}
	return 0
}

func (m *RsvpMgmtInterfaceSummary) GetReservationsOut() uint32 {
	if m != nil {
		return m.ReservationsOut
	}
	return 0
}

func init() {
	proto.RegisterType((*RsvpMgmtInterfaceSummary_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.controller_summaries.controller_summary.rsvp_mgmt_interface_summary_KEYS")
	proto.RegisterType((*RsvpMgmtInterfaceBwPrestdDste)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.controller_summaries.controller_summary.rsvp_mgmt_interface_bw_prestd_dste")
	proto.RegisterType((*RsvpMgmtInterfaceBwStdDste)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.controller_summaries.controller_summary.rsvp_mgmt_interface_bw_std_dste")
	proto.RegisterType((*RsvpMgmtDsteModeInterfaceBw)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.controller_summaries.controller_summary.rsvp_mgmt_dste_mode_interface_bw")
	proto.RegisterType((*RsvpMgmtInterfaceSummary)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.controller_summaries.controller_summary.rsvp_mgmt_interface_summary")
}

func init() { proto.RegisterFile("rsvp_mgmt_interface_summary.proto", fileDescriptor_1afdc8eaec0c32ba) }

var fileDescriptor_1afdc8eaec0c32ba = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0x75, 0x1f, 0x9d, 0x61, 0x5f, 0x86, 0x41, 0xa6, 0x6a, 0xa2, 0x2b, 0x17, 0x14,
	0x84, 0xa2, 0xb6, 0xe3, 0x43, 0xdc, 0xc1, 0x04, 0x13, 0xd5, 0x34, 0x86, 0xd2, 0x1b, 0xb8, 0xb2,
	0x9c, 0xc4, 0x5d, 0x8d, 0xe2, 0x38, 0xb2, 0x9d, 0x35, 0x7b, 0x08, 0x1e, 0x01, 0x71, 0xcd, 0x63,
	0xf0, 0x22, 0xbc, 0x05, 0xf7, 0xc8, 0x4e, 0xdb, 0x24, 0x34, 0x2b, 0x5c, 0x0c, 0x89, 0xbb, 0xf6,
	0x9c, 0xdf, 0xff, 0xef, 0x63, 0x1f, 0xfb, 0x04, 0x1c, 0x08, 0x79, 0x11, 0x23, 0x76, 0xce, 0x14,
	0xa2, 0x91, 0x22, 0x62, 0x88, 0x7d, 0x82, 0x64, 0xc2, 0x18, 0x16, 0x97, 0x4e, 0x2c, 0xb8, 0xe2,
	0xf0, 0xd8, 0xa7, 0xd2, 0xe7, 0x88, 0x72, 0x89, 0x52, 0x81, 0x68, 0x8c, 0x8c, 0x84, 0xc7, 0x44,
	0x38, 0xfa, 0x97, 0xe3, 0xf3, 0x48, 0x09, 0x1e, 0x86, 0x44, 0x4c, 0x84, 0x94, 0xc8, 0xf9, 0xe0,
	0x65, 0xeb, 0x04, 0x34, 0x17, 0x2c, 0x86, 0x4e, 0xde, 0x7c, 0x1c, 0xc0, 0x07, 0x60, 0xab, 0xa0,
	0x8c, 0x30, 0x23, 0xb6, 0xd5, 0xb4, 0xda, 0xeb, 0xee, 0x66, 0x1e, 0x7e, 0x87, 0x19, 0x69, 0xfd,
	0x58, 0x02, 0xad, 0x2a, 0x37, 0x6f, 0x8c, 0x62, 0x41, 0xa4, 0x0a, 0x50, 0x20, 0x15, 0x81, 0x8f,
	0x01, 0xc4, 0x61, 0xc8, 0x7d, 0xac, 0x48, 0x80, 0x3c, 0xaa, 0x90, 0xc0, 0x2a, 0xb3, 0x5c, 0x76,
	0xb7, 0x67, 0x99, 0x23, 0xaa, 0x5c, 0x9c, 0xd1, 0x0c, 0xa7, 0x68, 0x18, 0xf2, 0x31, 0xf2, 0x70,
	0x14, 0x8c, 0x69, 0xa0, 0x46, 0xf6, 0x52, 0x46, 0x33, 0x9c, 0x1e, 0x87, 0x7c, 0x7c, 0x34, 0x8d,
	0xc3, 0xfb, 0x60, 0x43, 0xd3, 0x39, 0x58, 0x33, 0xe0, 0x4d, 0x86, 0xd3, 0x1c, 0xea, 0x81, 0x5d,
	0x0d, 0xc9, 0xc4, 0x8b, 0x39, 0x0f, 0x0b, 0xf0, 0xb2, 0x81, 0x6f, 0x31, 0x9c, 0x0e, 0xb2, 0x5c,
	0xae, 0x79, 0x01, 0xf6, 0xa8, 0x44, 0x25, 0x6f, 0x84, 0x3d, 0xc9, 0xc3, 0x44, 0x11, 0x7b, 0xa5,
	0x69, 0xb5, 0xeb, 0xee, 0x1d, 0x2a, 0x4f, 0x0b, 0xcb, 0xbc, 0x9a, 0x64, 0xe1, 0x5b, 0x70, 0x30,
	0x91, 0xce, 0xad, 0x98, 0x5b, 0xac, 0x1a, 0x8b, 0x7d, 0x63, 0xf1, 0xfb, 0xe2, 0x53, 0xa7, 0xd6,
	0xf7, 0x1a, 0xb8, 0x77, 0xc5, 0x01, 0xff, 0x3f, 0xa7, 0xeb, 0x00, 0x7d, 0x80, 0x48, 0xef, 0xa0,
	0x33, 0x77, 0xb6, 0x3b, 0x0c, 0xa7, 0xef, 0x75, 0xa6, 0x92, 0xef, 0x16, 0xf8, 0x95, 0x12, 0xdf,
	0xfd, 0xcb, 0x4e, 0xac, 0x2e, 0xec, 0xc4, 0x4b, 0xb0, 0x3f, 0x95, 0xfa, 0x9d, 0x2a, 0xf9, 0x9a,
	0x91, 0xef, 0x65, 0x72, 0xbf, 0xb3, 0xd8, 0xa1, 0x5b, 0xe5, 0x50, 0x2f, 0x39, 0x74, 0xe7, 0x7b,
	0xf8, 0xb9, 0x56, 0x7c, 0x72, 0xba, 0x65, 0x88, 0xf1, 0x80, 0x94, 0xba, 0x09, 0x1b, 0x60, 0x7d,
	0x96, 0x99, 0x3c, 0xb6, 0xba, 0x0e, 0x9c, 0xf2, 0x80, 0xc0, 0x6f, 0x16, 0x68, 0xc4, 0x82, 0x20,
	0xa9, 0x70, 0x14, 0x60, 0x91, 0xf5, 0x3d, 0xd7, 0x9b, 0xee, 0xdd, 0xe8, 0x7d, 0x72, 0xae, 0x67,
	0x44, 0x38, 0x7f, 0x7e, 0xd1, 0xae, 0x1d, 0x0b, 0x32, 0x98, 0x54, 0xf3, 0x5a, 0x2a, 0xd2, 0x9f,
	0x82, 0xf0, 0xab, 0x05, 0xee, 0x5e, 0x55, 0x67, 0xcd, 0xd4, 0x79, 0xfe, 0x8f, 0xeb, 0x9c, 0x15,
	0xb9, 0x2b, 0xab, 0x2a, 0x6c, 0xfd, 0x5c, 0x02, 0x8d, 0x05, 0x23, 0x10, 0x3e, 0x02, 0x3b, 0x79,
	0x50, 0x0f, 0x3f, 0x94, 0x0a, 0xbb, 0x67, 0x5a, 0xb2, 0x35, 0x4b, 0xe8, 0xf1, 0xf7, 0x41, 0xc0,
	0x2f, 0x16, 0xd8, 0xcd, 0xef, 0x04, 0x8d, 0x86, 0x5c, 0x30, 0xac, 0x28, 0x8f, 0xec, 0x43, 0xb3,
	0xd7, 0xd1, 0xf5, 0xef, 0xb5, 0xfa, 0x02, 0xb9, 0xb7, 0x67, 0x65, 0xf4, 0xf3, 0x2a, 0xe0, 0x1e,
	0xa8, 0xc7, 0x58, 0x8d, 0x24, 0xa2, 0x91, 0xfd, 0xa4, 0x69, 0xb5, 0x37, 0xdc, 0x35, 0xf3, 0xbf,
	0x1f, 0xe9, 0x1b, 0x97, 0xa5, 0x78, 0xa2, 0xec, 0xa7, 0x26, 0x97, 0xb1, 0x67, 0x89, 0xd2, 0x5f,
	0x00, 0x41, 0x24, 0x11, 0x17, 0xc6, 0xc6, 0xc8, 0x9f, 0x19, 0x64, 0xb3, 0x18, 0xee, 0x47, 0xf0,
	0x21, 0xd8, 0x2e, 0x81, 0xda, 0xec, 0xb9, 0x21, 0x4b, 0x06, 0x67, 0x89, 0xf2, 0x56, 0xcd, 0x87,
	0xec, 0xf0, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xf6, 0x70, 0x88, 0xed, 0x06, 0x00, 0x00,
}