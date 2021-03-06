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
// source: pm_secy_if_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_macsec_macsec_ports_macsec_port_macsec_current_macsec_second30_macsec_second30secyifs_macsec_second30secyif

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

type PmSecyIfParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSecyIfParas_KEYS) Reset()         { *m = PmSecyIfParas_KEYS{} }
func (m *PmSecyIfParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmSecyIfParas_KEYS) ProtoMessage()    {}
func (*PmSecyIfParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_35a721e54758c8a5, []int{0}
}

func (m *PmSecyIfParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSecyIfParas_KEYS.Unmarshal(m, b)
}
func (m *PmSecyIfParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSecyIfParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmSecyIfParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSecyIfParas_KEYS.Merge(m, src)
}
func (m *PmSecyIfParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmSecyIfParas_KEYS.Size(m)
}
func (m *PmSecyIfParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSecyIfParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmSecyIfParas_KEYS proto.InternalMessageInfo

func (m *PmSecyIfParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmSecyIfParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PmSecyIfParam struct {
	Data                 uint64   `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            uint64   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,3,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	Valid                bool     `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSecyIfParam) Reset()         { *m = PmSecyIfParam{} }
func (m *PmSecyIfParam) String() string { return proto.CompactTextString(m) }
func (*PmSecyIfParam) ProtoMessage()    {}
func (*PmSecyIfParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_35a721e54758c8a5, []int{1}
}

func (m *PmSecyIfParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSecyIfParam.Unmarshal(m, b)
}
func (m *PmSecyIfParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSecyIfParam.Marshal(b, m, deterministic)
}
func (m *PmSecyIfParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSecyIfParam.Merge(m, src)
}
func (m *PmSecyIfParam) XXX_Size() int {
	return xxx_messageInfo_PmSecyIfParam.Size(m)
}
func (m *PmSecyIfParam) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSecyIfParam.DiscardUnknown(m)
}

var xxx_messageInfo_PmSecyIfParam proto.InternalMessageInfo

func (m *PmSecyIfParam) GetData() uint64 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PmSecyIfParam) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PmSecyIfParam) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

func (m *PmSecyIfParam) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type PmSecyIfParas struct {
	Index                uint32         `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool           `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            []string       `protobuf:"bytes,52,rep,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        []string       `protobuf:"bytes,53,rep,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   []string       `protobuf:"bytes,54,rep,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear30SecTime   []string       `protobuf:"bytes,55,rep,name=last_clear30_sec_time,json=lastClear30SecTime,proto3" json:"last_clear30_sec_time,omitempty"`
	LastClear24HrTime    []string       `protobuf:"bytes,56,rep,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Sec30Support         bool           `protobuf:"varint,57,opt,name=sec30_support,json=sec30Support,proto3" json:"sec30_support,omitempty"`
	SampleCount          uint64         `protobuf:"varint,58,opt,name=sample_count,json=sampleCount,proto3" json:"sample_count,omitempty"`
	InPktsUntagged       *PmSecyIfParam `protobuf:"bytes,59,opt,name=in_pkts_untagged,json=inPktsUntagged,proto3" json:"in_pkts_untagged,omitempty"`
	InPktsNoTag          *PmSecyIfParam `protobuf:"bytes,60,opt,name=in_pkts_no_tag,json=inPktsNoTag,proto3" json:"in_pkts_no_tag,omitempty"`
	InPktsBadTag         *PmSecyIfParam `protobuf:"bytes,61,opt,name=in_pkts_bad_tag,json=inPktsBadTag,proto3" json:"in_pkts_bad_tag,omitempty"`
	InPktsUnknownSci     *PmSecyIfParam `protobuf:"bytes,62,opt,name=in_pkts_unknown_sci,json=inPktsUnknownSci,proto3" json:"in_pkts_unknown_sci,omitempty"`
	InPktsNoSci          *PmSecyIfParam `protobuf:"bytes,63,opt,name=in_pkts_no_sci,json=inPktsNoSci,proto3" json:"in_pkts_no_sci,omitempty"`
	InPktsOverrun        *PmSecyIfParam `protobuf:"bytes,64,opt,name=in_pkts_overrun,json=inPktsOverrun,proto3" json:"in_pkts_overrun,omitempty"`
	InOctetsValidated    *PmSecyIfParam `protobuf:"bytes,65,opt,name=in_octets_validated,json=inOctetsValidated,proto3" json:"in_octets_validated,omitempty"`
	InOctetsDecrypted    *PmSecyIfParam `protobuf:"bytes,66,opt,name=in_octets_decrypted,json=inOctetsDecrypted,proto3" json:"in_octets_decrypted,omitempty"`
	OutPktsUntagged      *PmSecyIfParam `protobuf:"bytes,67,opt,name=out_pkts_untagged,json=outPktsUntagged,proto3" json:"out_pkts_untagged,omitempty"`
	OutPktsTooLong       *PmSecyIfParam `protobuf:"bytes,68,opt,name=out_pkts_too_long,json=outPktsTooLong,proto3" json:"out_pkts_too_long,omitempty"`
	OutOctetsProtected   *PmSecyIfParam `protobuf:"bytes,69,opt,name=out_octets_protected,json=outOctetsProtected,proto3" json:"out_octets_protected,omitempty"`
	OutOctetsEncrypted   *PmSecyIfParam `protobuf:"bytes,70,opt,name=out_octets_encrypted,json=outOctetsEncrypted,proto3" json:"out_octets_encrypted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PmSecyIfParas) Reset()         { *m = PmSecyIfParas{} }
func (m *PmSecyIfParas) String() string { return proto.CompactTextString(m) }
func (*PmSecyIfParas) ProtoMessage()    {}
func (*PmSecyIfParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_35a721e54758c8a5, []int{2}
}

func (m *PmSecyIfParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSecyIfParas.Unmarshal(m, b)
}
func (m *PmSecyIfParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSecyIfParas.Marshal(b, m, deterministic)
}
func (m *PmSecyIfParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSecyIfParas.Merge(m, src)
}
func (m *PmSecyIfParas) XXX_Size() int {
	return xxx_messageInfo_PmSecyIfParas.Size(m)
}
func (m *PmSecyIfParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSecyIfParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmSecyIfParas proto.InternalMessageInfo

func (m *PmSecyIfParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmSecyIfParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmSecyIfParas) GetTimestamp() []string {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PmSecyIfParas) GetLastClearTime() []string {
	if m != nil {
		return m.LastClearTime
	}
	return nil
}

func (m *PmSecyIfParas) GetLastClear15MinTime() []string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return nil
}

func (m *PmSecyIfParas) GetLastClear30SecTime() []string {
	if m != nil {
		return m.LastClear30SecTime
	}
	return nil
}

func (m *PmSecyIfParas) GetLastClear24HrTime() []string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return nil
}

func (m *PmSecyIfParas) GetSec30Support() bool {
	if m != nil {
		return m.Sec30Support
	}
	return false
}

func (m *PmSecyIfParas) GetSampleCount() uint64 {
	if m != nil {
		return m.SampleCount
	}
	return 0
}

func (m *PmSecyIfParas) GetInPktsUntagged() *PmSecyIfParam {
	if m != nil {
		return m.InPktsUntagged
	}
	return nil
}

func (m *PmSecyIfParas) GetInPktsNoTag() *PmSecyIfParam {
	if m != nil {
		return m.InPktsNoTag
	}
	return nil
}

func (m *PmSecyIfParas) GetInPktsBadTag() *PmSecyIfParam {
	if m != nil {
		return m.InPktsBadTag
	}
	return nil
}

func (m *PmSecyIfParas) GetInPktsUnknownSci() *PmSecyIfParam {
	if m != nil {
		return m.InPktsUnknownSci
	}
	return nil
}

func (m *PmSecyIfParas) GetInPktsNoSci() *PmSecyIfParam {
	if m != nil {
		return m.InPktsNoSci
	}
	return nil
}

func (m *PmSecyIfParas) GetInPktsOverrun() *PmSecyIfParam {
	if m != nil {
		return m.InPktsOverrun
	}
	return nil
}

func (m *PmSecyIfParas) GetInOctetsValidated() *PmSecyIfParam {
	if m != nil {
		return m.InOctetsValidated
	}
	return nil
}

func (m *PmSecyIfParas) GetInOctetsDecrypted() *PmSecyIfParam {
	if m != nil {
		return m.InOctetsDecrypted
	}
	return nil
}

func (m *PmSecyIfParas) GetOutPktsUntagged() *PmSecyIfParam {
	if m != nil {
		return m.OutPktsUntagged
	}
	return nil
}

func (m *PmSecyIfParas) GetOutPktsTooLong() *PmSecyIfParam {
	if m != nil {
		return m.OutPktsTooLong
	}
	return nil
}

func (m *PmSecyIfParas) GetOutOctetsProtected() *PmSecyIfParam {
	if m != nil {
		return m.OutOctetsProtected
	}
	return nil
}

func (m *PmSecyIfParas) GetOutOctetsEncrypted() *PmSecyIfParam {
	if m != nil {
		return m.OutOctetsEncrypted
	}
	return nil
}

func init() {
	proto.RegisterType((*PmSecyIfParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management.macsec.macsec_ports.macsec_port.macsec_current.macsec_second30.macsec_second30secyifs.macsec_second30secyif.pm_secy_if_paras_KEYS")
	proto.RegisterType((*PmSecyIfParam)(nil), "cisco_ios_xr_pmengine_oper.performance_management.macsec.macsec_ports.macsec_port.macsec_current.macsec_second30.macsec_second30secyifs.macsec_second30secyif.pm_secy_if_param")
	proto.RegisterType((*PmSecyIfParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management.macsec.macsec_ports.macsec_port.macsec_current.macsec_second30.macsec_second30secyifs.macsec_second30secyif.pm_secy_if_paras")
}

func init() { proto.RegisterFile("pm_secy_if_paras.proto", fileDescriptor_35a721e54758c8a5) }

var fileDescriptor_35a721e54758c8a5 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x97, 0xdb, 0x4e, 0x1b, 0x3b,
	0x14, 0x86, 0x35, 0x9b, 0x6c, 0xb4, 0x63, 0x08, 0x07, 0x6f, 0x60, 0xcf, 0xc5, 0xae, 0x94, 0x52,
	0xa9, 0xca, 0x55, 0x0a, 0x09, 0xf4, 0x7c, 0x24, 0x50, 0x55, 0xea, 0x01, 0x34, 0xa1, 0x95, 0x7a,
	0x65, 0x19, 0xcf, 0x62, 0xb0, 0x88, 0xed, 0x91, 0xed, 0xa1, 0xf0, 0x30, 0x7d, 0x86, 0xbe, 0x4c,
	0xab, 0x56, 0x6d, 0xa5, 0x56, 0x7d, 0x92, 0xca, 0x9e, 0x0c, 0x93, 0x0e, 0x7d, 0x80, 0xce, 0x15,
	0x5e, 0x3f, 0xeb, 0x5f, 0xb3, 0xbe, 0xe4, 0x8f, 0x9c, 0xa0, 0x95, 0x54, 0x10, 0x03, 0xec, 0x8c,
	0xf0, 0x43, 0x92, 0x52, 0x4d, 0x4d, 0x37, 0xd5, 0xca, 0x2a, 0xfc, 0x36, 0x60, 0xdc, 0x30, 0x45,
	0xb8, 0x32, 0xe4, 0x54, 0x93, 0x54, 0x80, 0x4c, 0xb8, 0x04, 0xa2, 0x52, 0xd0, 0xdd, 0x14, 0xf4,
	0xa1, 0xd2, 0x82, 0x4a, 0x06, 0x44, 0x50, 0x49, 0x13, 0x10, 0x20, 0x6d, 0x57, 0x50, 0x66, 0x80,
	0x8d, 0xff, 0x90, 0x54, 0x69, 0x6b, 0x26, 0x8b, 0xe2, 0xcc, 0x32, 0xad, 0xcb, 0x76, 0xf7, 0x74,
	0x25, 0xe3, 0xfe, 0x5a, 0xb5, 0x76, 0x5b, 0xf1, 0x43, 0xf3, 0x7b, 0x79, 0x75, 0x80, 0x96, 0xab,
	0x9b, 0x93, 0xa7, 0x3b, 0xaf, 0x87, 0x18, 0xa3, 0x86, 0xa4, 0x02, 0xc2, 0xa0, 0x1d, 0x74, 0x9a,
	0x91, 0x3f, 0xe3, 0x15, 0x34, 0x2d, 0x33, 0x71, 0x00, 0x3a, 0xfc, 0xab, 0x1d, 0x74, 0x5a, 0xd1,
	0xb8, 0x5a, 0x3d, 0x43, 0x0b, 0x95, 0x21, 0xc2, 0xf9, 0x63, 0x6a, 0xa9, 0xf7, 0x37, 0x22, 0x7f,
	0xc6, 0xff, 0xa3, 0xa6, 0x3d, 0xd2, 0x60, 0x8e, 0xd4, 0x28, 0xf6, 0x23, 0x1a, 0x51, 0x29, 0xe0,
	0x4b, 0x08, 0x59, 0x46, 0x89, 0x06, 0x07, 0x19, 0x4e, 0xb5, 0x83, 0xce, 0x3f, 0x51, 0xd3, 0x32,
	0x1a, 0x79, 0x01, 0x2f, 0xa1, 0xbf, 0x4f, 0xe8, 0x88, 0xc7, 0x61, 0xc3, 0xff, 0x27, 0x2f, 0x56,
	0xdf, 0xff, 0x77, 0xe1, 0xd9, 0xc6, 0xb5, 0x72, 0x19, 0xc3, 0x69, 0xd8, 0xf3, 0x6b, 0xe6, 0x45,
	0x39, 0xa0, 0x3f, 0x31, 0xc0, 0xef, 0xc4, 0x05, 0x18, 0x4b, 0x45, 0x1a, 0x6e, 0xb4, 0xa7, 0x3a,
	0xcd, 0xa8, 0x14, 0xf0, 0x55, 0x34, 0x3f, 0xa2, 0xc6, 0x12, 0x36, 0x02, 0xaa, 0x89, 0xd3, 0xc3,
	0x4d, 0xdf, 0xd3, 0x72, 0xf2, 0xc0, 0xa9, 0xfb, 0x5c, 0x00, 0x5e, 0x47, 0xcb, 0x65, 0xdf, 0xfa,
	0x26, 0x11, 0x5c, 0xe6, 0xdd, 0xd7, 0x7d, 0x37, 0x3e, 0xef, 0x5e, 0xdf, 0x7c, 0xce, 0xe5, 0x45,
	0x4b, 0x7f, 0xcd, 0x21, 0xe4, 0x96, 0x1b, 0x15, 0x4b, 0x7f, 0x6d, 0x08, 0xcc, 0x5b, 0xae, 0xa1,
	0xa5, 0xd2, 0xd2, 0xdb, 0x20, 0x47, 0xe3, 0x95, 0x6e, 0x7a, 0xc7, 0xe2, 0xb9, 0xa3, 0xb7, 0xf1,
	0x24, 0x5f, 0xeb, 0x0a, 0x6a, 0x19, 0x60, 0x6e, 0x78, 0x96, 0xfa, 0x57, 0xf5, 0x96, 0x47, 0x9f,
	0xf5, 0xe2, 0x30, 0xd7, 0xf0, 0x65, 0x34, 0x6b, 0xa8, 0x48, 0x47, 0x40, 0x98, 0xca, 0xa4, 0x0d,
	0x6f, 0xfb, 0x37, 0x66, 0x26, 0xd7, 0x06, 0x4e, 0xc2, 0x9f, 0x03, 0xb4, 0xc0, 0x25, 0x49, 0x8f,
	0xad, 0x21, 0x99, 0xb4, 0x34, 0x49, 0x20, 0x0e, 0xef, 0xb4, 0x83, 0xce, 0x4c, 0xef, 0x5d, 0xd0,
	0xfd, 0xa3, 0x13, 0xde, 0xad, 0x26, 0x33, 0x9a, 0xe3, 0x72, 0xef, 0xd8, 0x9a, 0x97, 0x63, 0x0e,
	0xfc, 0x21, 0x40, 0x73, 0x05, 0x9c, 0x54, 0xc4, 0xd2, 0x24, 0xbc, 0x5b, 0x53, 0xb4, 0x99, 0x1c,
	0xed, 0x85, 0xda, 0xa7, 0x09, 0xfe, 0x18, 0xa0, 0xf9, 0x82, 0xeb, 0x80, 0xc6, 0x1e, 0xec, 0x5e,
	0x4d, 0xc1, 0x66, 0x73, 0xb0, 0x2d, 0x1a, 0x3b, 0xb2, 0x6f, 0x01, 0xfa, 0xb7, 0x8c, 0xe3, 0xb1,
	0x54, 0x6f, 0x24, 0x31, 0x8c, 0x87, 0xf7, 0x6b, 0x4a, 0xb7, 0x50, 0x24, 0xd2, 0xa3, 0x0c, 0x19,
	0xaf, 0x66, 0xd2, 0xc1, 0x3d, 0xa8, 0x7b, 0x26, 0x1d, 0xd7, 0xa7, 0x89, 0x4c, 0xaa, 0x13, 0xd0,
	0x3a, 0x93, 0xe1, 0xc3, 0x9a, 0x82, 0xb5, 0x72, 0xb0, 0xdd, 0x1c, 0x03, 0x7f, 0xcf, 0x43, 0xa9,
	0x98, 0x05, 0x6b, 0x88, 0xbf, 0x5c, 0xa8, 0x85, 0x38, 0x7c, 0x54, 0x53, 0xbc, 0x45, 0x2e, 0x77,
	0x3d, 0xcb, 0xab, 0x02, 0xa5, 0x82, 0x18, 0x03, 0xd3, 0x67, 0xa9, 0x43, 0xdc, 0xaa, 0x3b, 0xe2,
	0x76, 0x81, 0x82, 0xbf, 0x06, 0x68, 0x51, 0x65, 0xb6, 0x72, 0xd5, 0x0d, 0x6a, 0x0a, 0x38, 0xaf,
	0x32, 0xfb, 0xcb, 0x5d, 0xf7, 0x65, 0x12, 0xcf, 0x2a, 0x45, 0x46, 0x4a, 0x26, 0xe1, 0x76, 0x5d,
	0x6f, 0xf2, 0x31, 0xde, 0xbe, 0x52, 0xcf, 0x94, 0x4c, 0xf0, 0x8f, 0x00, 0x2d, 0x39, 0xba, 0x71,
	0x40, 0xdd, 0x37, 0x70, 0x60, 0x2e, 0xa0, 0x3b, 0x35, 0x05, 0xc4, 0x2a, 0xb3, 0x79, 0x42, 0xf7,
	0x0a, 0x96, 0x2a, 0x24, 0xc8, 0xe2, 0x53, 0xf8, 0xb8, 0xf6, 0x90, 0x3b, 0x05, 0xcb, 0xc1, 0xb4,
	0xff, 0xf5, 0xd4, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0x95, 0xcc, 0x44, 0xe1, 0x57, 0x0d, 0x00,
	0x00,
}
