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
// source: ospf_sh_route_summary.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_summary_information

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

type OspfShRouteSummary_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRouteSummary_KEYS) Reset()         { *m = OspfShRouteSummary_KEYS{} }
func (m *OspfShRouteSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShRouteSummary_KEYS) ProtoMessage()    {}
func (*OspfShRouteSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_56147bc752a09782, []int{0}
}

func (m *OspfShRouteSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRouteSummary_KEYS.Unmarshal(m, b)
}
func (m *OspfShRouteSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRouteSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShRouteSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRouteSummary_KEYS.Merge(m, src)
}
func (m *OspfShRouteSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShRouteSummary_KEYS.Size(m)
}
func (m *OspfShRouteSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRouteSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRouteSummary_KEYS proto.InternalMessageInfo

func (m *OspfShRouteSummary_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShRouteSummary_KEYS) GetVrfName() string {
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
	return fileDescriptor_56147bc752a09782, []int{1}
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

type OspfShRouteSummCommon struct {
	ExternalType1S       uint32   `protobuf:"varint,1,opt,name=external_type1s,json=externalType1s,proto3" json:"external_type1s,omitempty"`
	ExternalType2S       uint32   `protobuf:"varint,2,opt,name=external_type2s,json=externalType2s,proto3" json:"external_type2s,omitempty"`
	ExternalNssaType1S   uint32   `protobuf:"varint,3,opt,name=external_nssa_type1s,json=externalNssaType1s,proto3" json:"external_nssa_type1s,omitempty"`
	ExternalNssaType2S   uint32   `protobuf:"varint,4,opt,name=external_nssa_type2s,json=externalNssaType2s,proto3" json:"external_nssa_type2s,omitempty"`
	InterAreas           uint32   `protobuf:"varint,5,opt,name=inter_areas,json=interAreas,proto3" json:"inter_areas,omitempty"`
	IntraAreas           uint32   `protobuf:"varint,6,opt,name=intra_areas,json=intraAreas,proto3" json:"intra_areas,omitempty"`
	Total                uint32   `protobuf:"varint,7,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRouteSummCommon) Reset()         { *m = OspfShRouteSummCommon{} }
func (m *OspfShRouteSummCommon) String() string { return proto.CompactTextString(m) }
func (*OspfShRouteSummCommon) ProtoMessage()    {}
func (*OspfShRouteSummCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_56147bc752a09782, []int{2}
}

func (m *OspfShRouteSummCommon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRouteSummCommon.Unmarshal(m, b)
}
func (m *OspfShRouteSummCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRouteSummCommon.Marshal(b, m, deterministic)
}
func (m *OspfShRouteSummCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRouteSummCommon.Merge(m, src)
}
func (m *OspfShRouteSummCommon) XXX_Size() int {
	return xxx_messageInfo_OspfShRouteSummCommon.Size(m)
}
func (m *OspfShRouteSummCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRouteSummCommon.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRouteSummCommon proto.InternalMessageInfo

func (m *OspfShRouteSummCommon) GetExternalType1S() uint32 {
	if m != nil {
		return m.ExternalType1S
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetExternalType2S() uint32 {
	if m != nil {
		return m.ExternalType2S
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetExternalNssaType1S() uint32 {
	if m != nil {
		return m.ExternalNssaType1S
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetExternalNssaType2S() uint32 {
	if m != nil {
		return m.ExternalNssaType2S
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetInterAreas() uint32 {
	if m != nil {
		return m.InterAreas
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetIntraAreas() uint32 {
	if m != nil {
		return m.IntraAreas
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type OspfShRouteSummary struct {
	Failures             uint32                 `protobuf:"varint,50,opt,name=failures,proto3" json:"failures,omitempty"`
	FailureTime          *OspfShTime            `protobuf:"bytes,51,opt,name=failure_time,json=failureTime,proto3" json:"failure_time,omitempty"`
	FailureAddress       string                 `protobuf:"bytes,52,opt,name=failure_address,json=failureAddress,proto3" json:"failure_address,omitempty"`
	Common               *OspfShRouteSummCommon `protobuf:"bytes,53,opt,name=common,proto3" json:"common,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *OspfShRouteSummary) Reset()         { *m = OspfShRouteSummary{} }
func (m *OspfShRouteSummary) String() string { return proto.CompactTextString(m) }
func (*OspfShRouteSummary) ProtoMessage()    {}
func (*OspfShRouteSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_56147bc752a09782, []int{3}
}

func (m *OspfShRouteSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRouteSummary.Unmarshal(m, b)
}
func (m *OspfShRouteSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRouteSummary.Marshal(b, m, deterministic)
}
func (m *OspfShRouteSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRouteSummary.Merge(m, src)
}
func (m *OspfShRouteSummary) XXX_Size() int {
	return xxx_messageInfo_OspfShRouteSummary.Size(m)
}
func (m *OspfShRouteSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRouteSummary.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRouteSummary proto.InternalMessageInfo

func (m *OspfShRouteSummary) GetFailures() uint32 {
	if m != nil {
		return m.Failures
	}
	return 0
}

func (m *OspfShRouteSummary) GetFailureTime() *OspfShTime {
	if m != nil {
		return m.FailureTime
	}
	return nil
}

func (m *OspfShRouteSummary) GetFailureAddress() string {
	if m != nil {
		return m.FailureAddress
	}
	return ""
}

func (m *OspfShRouteSummary) GetCommon() *OspfShRouteSummCommon {
	if m != nil {
		return m.Common
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShRouteSummary_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.summary_information.ospf_sh_route_summary_KEYS")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.summary_information.ospf_sh_time")
	proto.RegisterType((*OspfShRouteSummCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.summary_information.ospf_sh_route_summ_common")
	proto.RegisterType((*OspfShRouteSummary)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.summary_information.ospf_sh_route_summary")
}

func init() { proto.RegisterFile("ospf_sh_route_summary.proto", fileDescriptor_56147bc752a09782) }

var fileDescriptor_56147bc752a09782 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x55, 0xb6, 0x74, 0x5b, 0x66, 0x97, 0x56, 0xb2, 0x0a, 0x72, 0x8b, 0x04, 0x65, 0x2f, 0xed,
	0x29, 0x82, 0xb4, 0x7c, 0x40, 0x0f, 0x70, 0x41, 0xea, 0x21, 0xf4, 0x02, 0x17, 0x6b, 0xc8, 0x3a,
	0xc2, 0xd2, 0xda, 0x8e, 0x66, 0xbc, 0x51, 0x2b, 0xf1, 0x01, 0x9c, 0x11, 0x3f, 0xc1, 0x5f, 0xa2,
	0x38, 0xce, 0x6a, 0xd1, 0xe6, 0x4a, 0x2f, 0x91, 0xe7, 0xbd, 0x37, 0x2f, 0x2f, 0xf2, 0x0b, 0xbc,
	0xf4, 0xdc, 0xd4, 0x8a, 0xbf, 0x2b, 0xf2, 0xeb, 0xa0, 0x15, 0xaf, 0xad, 0x45, 0x7a, 0xc8, 0x1b,
	0xf2, 0xc1, 0x0b, 0xac, 0x0c, 0x57, 0x5e, 0x19, 0xcf, 0xea, 0x9e, 0x94, 0x69, 0xda, 0x6b, 0x15,
	0xe5, 0xbe, 0xd1, 0x94, 0x77, 0xa7, 0x4e, 0x57, 0x69, 0x66, 0xcd, 0xc3, 0x29, 0x6f, 0xa9, 0x8e,
	0x8f, 0xbc, 0x37, 0x34, 0xae, 0xf6, 0x64, 0x31, 0x18, 0xef, 0xf2, 0x64, 0xbe, 0x8d, 0x2d, 0xbe,
	0xc2, 0xd9, 0x68, 0x02, 0xf5, 0xe9, 0xc3, 0x97, 0xcf, 0xe2, 0x0d, 0xcc, 0x93, 0xaf, 0x72, 0x68,
	0xb5, 0xcc, 0xce, 0xb3, 0xcb, 0xa7, 0xe5, 0x2c, 0x61, 0xb7, 0x68, 0xb5, 0x38, 0x85, 0xc3, 0x96,
	0xea, 0x9e, 0x9e, 0x44, 0xfa, 0xa0, 0xa5, 0xba, 0xa3, 0x16, 0x1f, 0x61, 0x3e, 0x78, 0x07, 0x63,
	0xb5, 0x78, 0x01, 0x53, 0xd6, 0x95, 0x77, 0xcb, 0xe8, 0xf3, 0xac, 0x4c, 0x93, 0x78, 0x05, 0xe0,
	0xd0, 0xf9, 0xc4, 0x4d, 0x22, 0xb7, 0x85, 0x2c, 0xfe, 0x4c, 0xe0, 0x74, 0x37, 0xa4, 0xaa, 0xbc,
	0xb5, 0xde, 0x89, 0x0b, 0x38, 0xd6, 0xf7, 0x41, 0x93, 0xc3, 0x95, 0x0a, 0x0f, 0x8d, 0x7e, 0xc7,
	0xc9, 0xfe, 0x68, 0x80, 0xef, 0x22, 0xba, 0x23, 0x2c, 0x38, 0xbd, 0xeb, 0x1f, 0x61, 0xc1, 0xe2,
	0x2d, 0x9c, 0x6c, 0x84, 0x8e, 0x19, 0x07, 0xdb, 0xbd, 0xa8, 0x16, 0x03, 0x77, 0xcb, 0x8c, 0xc9,
	0x7a, 0x74, 0xa3, 0x60, 0xf9, 0x64, 0x7c, 0xa3, 0x60, 0xf1, 0x1a, 0x66, 0xc6, 0x05, 0x4d, 0x0a,
	0x49, 0x23, 0xcb, 0xfd, 0xfe, 0xa3, 0x23, 0x74, 0xd3, 0x21, 0x49, 0x40, 0x98, 0x04, 0xd3, 0x8d,
	0x80, 0xb0, 0x17, 0x9c, 0xc0, 0x7e, 0xf0, 0x01, 0x57, 0xf2, 0x20, 0x52, 0xfd, 0xb0, 0xf8, 0xb9,
	0x07, 0xcf, 0x47, 0x2f, 0x54, 0x9c, 0xc1, 0x61, 0x8d, 0x66, 0xb5, 0x26, 0xcd, 0xb2, 0x88, 0x2b,
	0x9b, 0x59, 0xfc, 0xca, 0x60, 0x9e, 0x86, 0x78, 0x55, 0xf2, 0xea, 0x3c, 0xbb, 0x9c, 0x15, 0x3e,
	0xff, 0xef, 0x05, 0xcc, 0xb7, 0x1b, 0x52, 0xce, 0x52, 0x88, 0xbb, 0xae, 0x2e, 0x17, 0x70, 0x3c,
	0x64, 0xc2, 0xe5, 0x92, 0x34, 0xb3, 0xbc, 0x8e, 0x05, 0x3b, 0x4a, 0xf0, 0x4d, 0x8f, 0x8a, 0xdf,
	0x19, 0x4c, 0xfb, 0x32, 0xc8, 0xf7, 0x31, 0xf7, 0x8f, 0x47, 0xcc, 0xbd, 0x53, 0xc8, 0x32, 0x65,
	0xf9, 0x36, 0x8d, 0x3f, 0xf1, 0xd5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xec, 0x05, 0xf3,
	0xe3, 0x03, 0x00, 0x00,
}
