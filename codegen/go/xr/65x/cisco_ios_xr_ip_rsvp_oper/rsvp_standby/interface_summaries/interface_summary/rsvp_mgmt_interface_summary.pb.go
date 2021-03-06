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

package cisco_ios_xr_ip_rsvp_oper_rsvp_standby_interface_summaries_interface_summary

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
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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

func (m *RsvpMgmtInterfaceSummary_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
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
	proto.RegisterType((*RsvpMgmtInterfaceSummary_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.interface_summaries.interface_summary.rsvp_mgmt_interface_summary_KEYS")
	proto.RegisterType((*RsvpMgmtInterfaceBwPrestdDste)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.interface_summaries.interface_summary.rsvp_mgmt_interface_bw_prestd_dste")
	proto.RegisterType((*RsvpMgmtInterfaceBwStdDste)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.interface_summaries.interface_summary.rsvp_mgmt_interface_bw_std_dste")
	proto.RegisterType((*RsvpMgmtDsteModeInterfaceBw)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.interface_summaries.interface_summary.rsvp_mgmt_dste_mode_interface_bw")
	proto.RegisterType((*RsvpMgmtInterfaceSummary)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.interface_summaries.interface_summary.rsvp_mgmt_interface_summary")
}

func init() { proto.RegisterFile("rsvp_mgmt_interface_summary.proto", fileDescriptor_1afdc8eaec0c32ba) }

var fileDescriptor_1afdc8eaec0c32ba = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x95, 0x76, 0xeb, 0x3a, 0xff, 0x7e, 0xdd, 0x1f, 0x43, 0x21, 0x55, 0x35, 0xd1, 0x15,
	0x21, 0x0a, 0x42, 0x51, 0xdb, 0xf1, 0x47, 0xdc, 0xc1, 0x04, 0x88, 0x0a, 0xc6, 0x50, 0x7a, 0x03,
	0x57, 0x96, 0x93, 0xb8, 0xcc, 0x52, 0x1c, 0x47, 0xb6, 0xb3, 0xa6, 0x8f, 0xc2, 0x35, 0xe2, 0x9a,
	0x67, 0xe0, 0x45, 0x78, 0x11, 0x6e, 0x50, 0x9c, 0xb6, 0x49, 0xd6, 0xac, 0x70, 0x01, 0x88, 0xbb,
	0xf6, 0x9c, 0xcf, 0xf7, 0xeb, 0x63, 0x1f, 0xfb, 0x04, 0x1c, 0x0a, 0x79, 0x1e, 0x22, 0xf6, 0x81,
	0x29, 0x44, 0x03, 0x45, 0xc4, 0x04, 0xbb, 0x04, 0xc9, 0x88, 0x31, 0x2c, 0x66, 0x56, 0x28, 0xb8,
	0xe2, 0xf0, 0xb5, 0x4b, 0xa5, 0xcb, 0x11, 0xe5, 0x12, 0xc5, 0x02, 0xd1, 0x10, 0x69, 0x09, 0x0f,
	0x89, 0xb0, 0xf4, 0x2f, 0xa9, 0x70, 0xe0, 0x39, 0x33, 0xeb, 0xa2, 0x9e, 0x12, 0xb9, 0x12, 0x9b,
	0x75, 0x47, 0xa0, 0xb3, 0x66, 0x49, 0xf4, 0xea, 0xf9, 0xfb, 0x31, 0xbc, 0x05, 0x76, 0xb2, 0x4c,
	0x80, 0x19, 0x31, 0x8d, 0x8e, 0xd1, 0xdb, 0xb6, 0x1b, 0xcb, 0xe8, 0x1b, 0xcc, 0x48, 0xf7, 0x5b,
	0x05, 0x74, 0xcb, 0xbc, 0x9c, 0x29, 0x0a, 0x05, 0x91, 0xca, 0x43, 0x9e, 0x54, 0x04, 0xde, 0x03,
	0x10, 0xfb, 0x3e, 0x77, 0xb1, 0x22, 0x1e, 0x72, 0xa8, 0x42, 0x02, 0xab, 0xd4, 0x71, 0xc3, 0xde,
	0x5b, 0x66, 0x8e, 0xa9, 0xb2, 0x71, 0x4a, 0x33, 0x1c, 0xa3, 0x89, 0xcf, 0xa7, 0xc8, 0xc1, 0x81,
	0x37, 0xa5, 0x9e, 0x3a, 0x33, 0x2b, 0x29, 0xcd, 0x70, 0xfc, 0xc2, 0xe7, 0xd3, 0xe3, 0x45, 0x1c,
	0xde, 0x04, 0x8d, 0x84, 0xce, 0xc0, 0xaa, 0x06, 0xff, 0x67, 0x38, 0xce, 0xa0, 0x21, 0x68, 0x26,
	0x90, 0x8c, 0x9c, 0x90, 0x73, 0x3f, 0x07, 0x6f, 0x68, 0xf8, 0x0a, 0xc3, 0xf1, 0x38, 0xcd, 0x65,
	0x9a, 0xc7, 0xa0, 0x45, 0x25, 0x2a, 0x78, 0x23, 0xec, 0x48, 0xee, 0x47, 0x8a, 0x98, 0x9b, 0x1d,
	0xa3, 0x57, 0xb7, 0xaf, 0x51, 0x79, 0x92, 0x5b, 0xe6, 0xe9, 0x3c, 0x0b, 0x5f, 0x82, 0xc3, 0xb9,
	0x74, 0x65, 0xc5, 0xcc, 0xa2, 0xa6, 0x2d, 0x0e, 0xb4, 0xc5, 0xc5, 0xc5, 0x17, 0x4e, 0xdd, 0xaf,
	0x55, 0x70, 0xe3, 0x92, 0x03, 0xfe, 0x77, 0x4e, 0xd7, 0x02, 0xc9, 0x01, 0xa2, 0x64, 0x07, 0xfd,
	0x95, 0xb3, 0xdd, 0x67, 0x38, 0x7e, 0x9b, 0x64, 0x4a, 0xf9, 0x41, 0x8e, 0xdf, 0x2c, 0xf0, 0x83,
	0x5f, 0xec, 0x44, 0x6d, 0x6d, 0x27, 0x9e, 0x80, 0x83, 0x85, 0xd4, 0xed, 0x97, 0xc9, 0xb7, 0xb4,
	0xbc, 0x95, 0xca, 0xdd, 0xfe, 0x7a, 0x87, 0x41, 0x99, 0x43, 0xbd, 0xe0, 0x30, 0x58, 0xed, 0xe1,
	0xc7, 0x6a, 0xfe, 0xc1, 0x25, 0x2d, 0x43, 0x8c, 0x7b, 0xa4, 0xd0, 0x4d, 0xd8, 0x06, 0xdb, 0xcb,
	0xcc, 0xfc, 0xad, 0xd5, 0x93, 0xc0, 0x09, 0xf7, 0x08, 0xfc, 0x62, 0x80, 0x76, 0x28, 0x48, 0xfa,
	0xce, 0xb1, 0x48, 0xfb, 0x9e, 0xe9, 0x75, 0xf7, 0xfe, 0x1b, 0x86, 0xd6, 0xef, 0x1c, 0x13, 0xd6,
	0xcf, 0xdf, 0xb5, 0x6d, 0x86, 0x82, 0x8c, 0xe7, 0x35, 0x3d, 0x93, 0x8a, 0x8c, 0x16, 0x20, 0xfc,
	0x6c, 0x80, 0xeb, 0x97, 0x55, 0x5b, 0xd5, 0xd5, 0xb2, 0xbf, 0x52, 0xed, 0xb2, 0xd4, 0xa6, 0x2c,
	0xab, 0xb3, 0xfb, 0xbd, 0x02, 0xda, 0x6b, 0x86, 0x21, 0xbc, 0x0b, 0xf6, 0x8b, 0x73, 0x10, 0xc5,
	0xc2, 0x1c, 0xea, 0xf6, 0xec, 0x16, 0x46, 0xe1, 0x3b, 0x01, 0x3f, 0x19, 0xa0, 0x99, 0xdd, 0x0f,
	0x1a, 0x4c, 0xb8, 0x60, 0x58, 0x51, 0x1e, 0x98, 0x47, 0x7a, 0xc7, 0xc1, 0x9f, 0xda, 0x71, 0xf9,
	0x95, 0xb2, 0xaf, 0x2e, 0x8b, 0x19, 0x65, 0xb5, 0xc0, 0x16, 0xa8, 0x87, 0x58, 0x9d, 0x49, 0x44,
	0x03, 0xf3, 0x7e, 0xc7, 0xe8, 0x35, 0xec, 0x2d, 0xfd, 0x7f, 0x14, 0x24, 0x77, 0x30, 0x4d, 0xf1,
	0x48, 0x99, 0x0f, 0x74, 0x2e, 0x65, 0x4f, 0x23, 0x05, 0x6f, 0x83, 0x5d, 0x41, 0x24, 0x11, 0xe7,
	0xda, 0x46, 0xcb, 0x1f, 0x6a, 0x64, 0x27, 0x1f, 0x1e, 0x05, 0xf0, 0x0e, 0xd8, 0x2b, 0x80, 0x89,
	0xd9, 0x23, 0x4d, 0x16, 0x0c, 0x4e, 0x23, 0xe5, 0xd4, 0xf4, 0xe7, 0xed, 0xe8, 0x47, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x48, 0xfa, 0x5a, 0x13, 0x03, 0x07, 0x00, 0x00,
}
