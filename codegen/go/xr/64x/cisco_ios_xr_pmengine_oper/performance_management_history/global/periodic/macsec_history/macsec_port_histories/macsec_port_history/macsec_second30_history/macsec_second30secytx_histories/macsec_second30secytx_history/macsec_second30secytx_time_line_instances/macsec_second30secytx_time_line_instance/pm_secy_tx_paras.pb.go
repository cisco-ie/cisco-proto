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
// source: pm_secy_tx_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_history_global_periodic_macsec_history_macsec_port_histories_macsec_port_history_macsec_second30_history_macsec_second30secytx_histories_macsec_second30secytx_history_macsec_second30secytx_time_line_instances_macsec_second30secytx_time_line_instance

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

type PmSecyTxParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Number_1             uint32   `protobuf:"varint,3,opt,name=number_1,json=number1,proto3" json:"number_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSecyTxParas_KEYS) Reset()         { *m = PmSecyTxParas_KEYS{} }
func (m *PmSecyTxParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmSecyTxParas_KEYS) ProtoMessage()    {}
func (*PmSecyTxParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e70c245f5df3d574, []int{0}
}

func (m *PmSecyTxParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSecyTxParas_KEYS.Unmarshal(m, b)
}
func (m *PmSecyTxParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSecyTxParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmSecyTxParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSecyTxParas_KEYS.Merge(m, src)
}
func (m *PmSecyTxParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmSecyTxParas_KEYS.Size(m)
}
func (m *PmSecyTxParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSecyTxParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmSecyTxParas_KEYS proto.InternalMessageInfo

func (m *PmSecyTxParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmSecyTxParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *PmSecyTxParas_KEYS) GetNumber_1() uint32 {
	if m != nil {
		return m.Number_1
	}
	return 0
}

type PmSecyTxParam struct {
	Data                 uint64   `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            uint64   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,3,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	Valid                bool     `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSecyTxParam) Reset()         { *m = PmSecyTxParam{} }
func (m *PmSecyTxParam) String() string { return proto.CompactTextString(m) }
func (*PmSecyTxParam) ProtoMessage()    {}
func (*PmSecyTxParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_e70c245f5df3d574, []int{1}
}

func (m *PmSecyTxParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSecyTxParam.Unmarshal(m, b)
}
func (m *PmSecyTxParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSecyTxParam.Marshal(b, m, deterministic)
}
func (m *PmSecyTxParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSecyTxParam.Merge(m, src)
}
func (m *PmSecyTxParam) XXX_Size() int {
	return xxx_messageInfo_PmSecyTxParam.Size(m)
}
func (m *PmSecyTxParam) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSecyTxParam.DiscardUnknown(m)
}

var xxx_messageInfo_PmSecyTxParam proto.InternalMessageInfo

func (m *PmSecyTxParam) GetData() uint64 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PmSecyTxParam) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PmSecyTxParam) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

func (m *PmSecyTxParam) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type PmSecyTxParas struct {
	Index                uint32         `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool           `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            []string       `protobuf:"bytes,52,rep,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        []string       `protobuf:"bytes,53,rep,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   []string       `protobuf:"bytes,54,rep,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear30SecTime   []string       `protobuf:"bytes,55,rep,name=last_clear30_sec_time,json=lastClear30SecTime,proto3" json:"last_clear30_sec_time,omitempty"`
	LastClear24HrTime    []string       `protobuf:"bytes,56,rep,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Sec30Support         bool           `protobuf:"varint,57,opt,name=sec30_support,json=sec30Support,proto3" json:"sec30_support,omitempty"`
	SampleCount          uint64         `protobuf:"varint,58,opt,name=sample_count,json=sampleCount,proto3" json:"sample_count,omitempty"`
	OutPktsProtected     *PmSecyTxParam `protobuf:"bytes,59,opt,name=out_pkts_protected,json=outPktsProtected,proto3" json:"out_pkts_protected,omitempty"`
	OutPktsEncrypted     *PmSecyTxParam `protobuf:"bytes,60,opt,name=out_pkts_encrypted,json=outPktsEncrypted,proto3" json:"out_pkts_encrypted,omitempty"`
	OutOctetsProtected   *PmSecyTxParam `protobuf:"bytes,61,opt,name=out_octets_protected,json=outOctetsProtected,proto3" json:"out_octets_protected,omitempty"`
	OutOctetsEncrypted   *PmSecyTxParam `protobuf:"bytes,62,opt,name=out_octets_encrypted,json=outOctetsEncrypted,proto3" json:"out_octets_encrypted,omitempty"`
	OutPktsTooLong       *PmSecyTxParam `protobuf:"bytes,63,opt,name=out_pkts_too_long,json=outPktsTooLong,proto3" json:"out_pkts_too_long,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PmSecyTxParas) Reset()         { *m = PmSecyTxParas{} }
func (m *PmSecyTxParas) String() string { return proto.CompactTextString(m) }
func (*PmSecyTxParas) ProtoMessage()    {}
func (*PmSecyTxParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_e70c245f5df3d574, []int{2}
}

func (m *PmSecyTxParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSecyTxParas.Unmarshal(m, b)
}
func (m *PmSecyTxParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSecyTxParas.Marshal(b, m, deterministic)
}
func (m *PmSecyTxParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSecyTxParas.Merge(m, src)
}
func (m *PmSecyTxParas) XXX_Size() int {
	return xxx_messageInfo_PmSecyTxParas.Size(m)
}
func (m *PmSecyTxParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSecyTxParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmSecyTxParas proto.InternalMessageInfo

func (m *PmSecyTxParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmSecyTxParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmSecyTxParas) GetTimestamp() []string {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PmSecyTxParas) GetLastClearTime() []string {
	if m != nil {
		return m.LastClearTime
	}
	return nil
}

func (m *PmSecyTxParas) GetLastClear15MinTime() []string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return nil
}

func (m *PmSecyTxParas) GetLastClear30SecTime() []string {
	if m != nil {
		return m.LastClear30SecTime
	}
	return nil
}

func (m *PmSecyTxParas) GetLastClear24HrTime() []string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return nil
}

func (m *PmSecyTxParas) GetSec30Support() bool {
	if m != nil {
		return m.Sec30Support
	}
	return false
}

func (m *PmSecyTxParas) GetSampleCount() uint64 {
	if m != nil {
		return m.SampleCount
	}
	return 0
}

func (m *PmSecyTxParas) GetOutPktsProtected() *PmSecyTxParam {
	if m != nil {
		return m.OutPktsProtected
	}
	return nil
}

func (m *PmSecyTxParas) GetOutPktsEncrypted() *PmSecyTxParam {
	if m != nil {
		return m.OutPktsEncrypted
	}
	return nil
}

func (m *PmSecyTxParas) GetOutOctetsProtected() *PmSecyTxParam {
	if m != nil {
		return m.OutOctetsProtected
	}
	return nil
}

func (m *PmSecyTxParas) GetOutOctetsEncrypted() *PmSecyTxParam {
	if m != nil {
		return m.OutOctetsEncrypted
	}
	return nil
}

func (m *PmSecyTxParas) GetOutPktsTooLong() *PmSecyTxParam {
	if m != nil {
		return m.OutPktsTooLong
	}
	return nil
}

func init() {
	proto.RegisterType((*PmSecyTxParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.macsec_history.macsec_port_histories.macsec_port_history.macsec_second30_history.macsec_second30secytx_histories.macsec_second30secytx_history.macsec_second30secytx_time_line_instances.macsec_second30secytx_time_line_instance.pm_secy_tx_paras_KEYS")
	proto.RegisterType((*PmSecyTxParam)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.macsec_history.macsec_port_histories.macsec_port_history.macsec_second30_history.macsec_second30secytx_histories.macsec_second30secytx_history.macsec_second30secytx_time_line_instances.macsec_second30secytx_time_line_instance.pm_secy_tx_param")
	proto.RegisterType((*PmSecyTxParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.macsec_history.macsec_port_histories.macsec_port_history.macsec_second30_history.macsec_second30secytx_histories.macsec_second30secytx_history.macsec_second30secytx_time_line_instances.macsec_second30secytx_time_line_instance.pm_secy_tx_paras")
}

func init() { proto.RegisterFile("pm_secy_tx_paras.proto", fileDescriptor_e70c245f5df3d574) }

var fileDescriptor_e70c245f5df3d574 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x96, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0x95, 0x6d, 0x28, 0x8d, 0xdb, 0x42, 0x6b, 0xb5, 0xd5, 0x22, 0x81, 0x14, 0x8a, 0x84,
	0x72, 0x5a, 0x9a, 0xa4, 0xe1, 0xfb, 0xe3, 0x50, 0x55, 0x42, 0x02, 0x44, 0xb5, 0xe9, 0x85, 0x0b,
	0x96, 0xe3, 0x1d, 0x12, 0xab, 0xeb, 0x0f, 0xd9, 0x0e, 0x4a, 0x9e, 0x83, 0xc7, 0xe0, 0x19, 0x78,
	0x20, 0x24, 0x0e, 0x3c, 0x02, 0xb2, 0x37, 0xc9, 0xd2, 0xa4, 0x95, 0x7a, 0xc8, 0x31, 0x37, 0xcf,
	0xdf, 0xff, 0x99, 0x1d, 0xff, 0x32, 0xb6, 0x82, 0x0e, 0xb4, 0x20, 0x16, 0xd8, 0x98, 0xb8, 0x11,
	0xd1, 0xd4, 0x50, 0x9b, 0x68, 0xa3, 0x9c, 0xc2, 0x3f, 0x23, 0xc6, 0x2d, 0x53, 0x84, 0x2b, 0x4b,
	0x46, 0x86, 0x68, 0x01, 0xb2, 0xcf, 0x25, 0x10, 0xa5, 0xc1, 0x24, 0x1a, 0xcc, 0x37, 0x65, 0x04,
	0x95, 0x0c, 0x88, 0xa0, 0x92, 0xf6, 0x41, 0x80, 0x74, 0x64, 0xc0, 0xad, 0x53, 0x66, 0x9c, 0xf4,
	0x73, 0xd5, 0xa3, 0xb9, 0x77, 0x71, 0x95, 0x71, 0x96, 0x08, 0xca, 0x2c, 0xb0, 0xd9, 0xf6, 0x24,
	0xd4, 0xca, 0x4c, 0x53, 0x38, 0xd8, 0x2b, 0xd4, 0x99, 0xd3, 0x02, 0x53, 0x32, 0x6b, 0x1f, 0x5d,
	0xa7, 0xfb, 0xfe, 0xdd, 0x68, 0xb1, 0xd6, 0x95, 0xfb, 0xd7, 0x65, 0x3b, 0x2e, 0x80, 0xe4, 0xfe,
	0x88, 0x5c, 0x5a, 0xe7, 0x8f, 0x66, 0x6f, 0xec, 0x3c, 0xfc, 0x8a, 0xf6, 0xe7, 0x39, 0x92, 0x0f,
	0xa7, 0x5f, 0xba, 0x18, 0xa3, 0xaa, 0xa4, 0x02, 0xe2, 0x4a, 0xbd, 0xd2, 0xa8, 0xa5, 0x61, 0x8d,
	0x0f, 0xd0, 0xba, 0x1c, 0x8a, 0x1e, 0x98, 0x38, 0xaa, 0x57, 0x1a, 0xdb, 0xe9, 0x24, 0xc2, 0xf7,
	0xd0, 0x46, 0xb1, 0x22, 0xcd, 0x78, 0x2d, 0xec, 0xdc, 0x2e, 0xe2, 0xe6, 0xe1, 0x18, 0xed, 0xcc,
	0xd5, 0x17, 0xbe, 0x74, 0x46, 0x1d, 0x0d, 0xa5, 0xab, 0x69, 0x58, 0xe3, 0xfb, 0xa8, 0xe6, 0x06,
	0x06, 0xec, 0x40, 0xe5, 0x59, 0xa8, 0x5e, 0x4d, 0x4b, 0x01, 0x3f, 0x40, 0xc8, 0x31, 0x4a, 0x0c,
	0x78, 0xc8, 0xe1, 0x13, 0x1b, 0x69, 0xcd, 0x31, 0x9a, 0x06, 0x01, 0xef, 0xa1, 0x5b, 0xdf, 0x69,
	0xce, 0xb3, 0xb8, 0x1a, 0x76, 0x8a, 0xe0, 0xf0, 0xc7, 0xee, 0xc2, 0xb7, 0xad, 0xb7, 0x72, 0x99,
	0xc1, 0x28, 0x6e, 0x85, 0x3e, 0x8b, 0xa0, 0x2c, 0xd0, 0xfe, 0xaf, 0x40, 0xe8, 0x89, 0x0b, 0xb0,
	0x8e, 0x0a, 0x1d, 0x1f, 0xd7, 0xd7, 0x1a, 0xb5, 0xb4, 0x14, 0xf0, 0x63, 0x74, 0x37, 0xa7, 0xd6,
	0x11, 0x96, 0x03, 0x35, 0x01, 0x6d, 0xdc, 0x09, 0x9e, 0x6d, 0x2f, 0x9f, 0x78, 0xf5, 0x9c, 0x0b,
	0xc0, 0x4d, 0xb4, 0x5f, 0xfa, 0x9a, 0x1d, 0x22, 0xb8, 0x2c, 0xdc, 0x4f, 0x83, 0x1b, 0xcf, 0xdc,
	0xcd, 0xce, 0x27, 0x2e, 0x17, 0x53, 0xda, 0x47, 0xfe, 0x08, 0x45, 0xca, 0xb3, 0xb9, 0x94, 0xf6,
	0x51, 0x17, 0x58, 0x48, 0x79, 0x82, 0xf6, 0xca, 0x94, 0xd6, 0x31, 0x19, 0x4c, 0x5a, 0x7a, 0x1e,
	0x32, 0x76, 0x67, 0x19, 0xad, 0xe3, 0xf7, 0x45, 0x5b, 0x8f, 0xd0, 0xb6, 0x05, 0xe6, 0x8b, 0x0f,
	0x75, 0xa0, 0xfa, 0x22, 0x1c, 0x7d, 0x2b, 0x88, 0xdd, 0x42, 0xc3, 0x0f, 0xd1, 0x96, 0xa5, 0x42,
	0xe7, 0x40, 0x98, 0x1a, 0x4a, 0x17, 0xbf, 0x0c, 0x3f, 0xcc, 0x66, 0xa1, 0x9d, 0x78, 0x09, 0xff,
	0x89, 0x10, 0x56, 0x43, 0x47, 0xf4, 0x85, 0xb3, 0xc4, 0xdf, 0x40, 0x60, 0x0e, 0xb2, 0xf8, 0x55,
	0xbd, 0xd2, 0xd8, 0x6c, 0xfd, 0x8a, 0x92, 0xd5, 0x65, 0xbc, 0xb1, 0x33, 0x99, 0xbf, 0x29, 0xe9,
	0x8e, 0x1a, 0xba, 0xb3, 0x0b, 0x67, 0xcf, 0xa6, 0x5c, 0x2f, 0xe3, 0x06, 0xc9, 0xcc, 0x58, 0x7b,
	0xdc, 0xaf, 0x57, 0xb8, 0x97, 0x88, 0xfb, 0x74, 0xca, 0x15, 0xff, 0x8d, 0xd0, 0x9e, 0xc7, 0xad,
	0x98, 0x83, 0x4b, 0xf3, 0xfd, 0x66, 0x05, 0x7c, 0x09, 0xc0, 0xfd, 0x28, 0x7f, 0x0e, 0x68, 0xcb,
	0x09, 0x9f, 0x43, 0x5e, 0xce, 0xf8, 0xdb, 0x15, 0xf2, 0xa5, 0x22, 0x2f, 0xa7, 0xfc, 0x77, 0x84,
	0x76, 0x67, 0x8f, 0x8a, 0x53, 0x8a, 0xe4, 0x4a, 0xf6, 0xe3, 0x77, 0x2b, 0xde, 0x4b, 0xe0, 0x7d,
	0x67, 0xf2, 0xa6, 0x9c, 0x2b, 0xf5, 0x51, 0xc9, 0x7e, 0x6f, 0x3d, 0xfc, 0x4b, 0x6d, 0xff, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x66, 0x5f, 0x12, 0xff, 0xbf, 0x0a, 0x00, 0x00,
}
