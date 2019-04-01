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
// source: pm_secy_rx_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_history_global_periodic_macsec_history_macsec_port_histories_macsec_port_history_macsec_hour24_history_macsec_hour24secyrx_histories_macsec_hour24secyrx_history_macsec_hour24secyrx_time_line_instances_macsec_hour24secyrx_time_line_instance

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

type PmSecyRxParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Number_1             uint32   `protobuf:"varint,3,opt,name=number_1,json=number1,proto3" json:"number_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSecyRxParas_KEYS) Reset()         { *m = PmSecyRxParas_KEYS{} }
func (m *PmSecyRxParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmSecyRxParas_KEYS) ProtoMessage()    {}
func (*PmSecyRxParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e3fa764ed03632f, []int{0}
}

func (m *PmSecyRxParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSecyRxParas_KEYS.Unmarshal(m, b)
}
func (m *PmSecyRxParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSecyRxParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmSecyRxParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSecyRxParas_KEYS.Merge(m, src)
}
func (m *PmSecyRxParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmSecyRxParas_KEYS.Size(m)
}
func (m *PmSecyRxParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSecyRxParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmSecyRxParas_KEYS proto.InternalMessageInfo

func (m *PmSecyRxParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmSecyRxParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *PmSecyRxParas_KEYS) GetNumber_1() uint32 {
	if m != nil {
		return m.Number_1
	}
	return 0
}

type PmSecyRxParam struct {
	Data                 uint64   `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            uint64   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,3,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	Valid                bool     `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSecyRxParam) Reset()         { *m = PmSecyRxParam{} }
func (m *PmSecyRxParam) String() string { return proto.CompactTextString(m) }
func (*PmSecyRxParam) ProtoMessage()    {}
func (*PmSecyRxParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e3fa764ed03632f, []int{1}
}

func (m *PmSecyRxParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSecyRxParam.Unmarshal(m, b)
}
func (m *PmSecyRxParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSecyRxParam.Marshal(b, m, deterministic)
}
func (m *PmSecyRxParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSecyRxParam.Merge(m, src)
}
func (m *PmSecyRxParam) XXX_Size() int {
	return xxx_messageInfo_PmSecyRxParam.Size(m)
}
func (m *PmSecyRxParam) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSecyRxParam.DiscardUnknown(m)
}

var xxx_messageInfo_PmSecyRxParam proto.InternalMessageInfo

func (m *PmSecyRxParam) GetData() uint64 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PmSecyRxParam) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PmSecyRxParam) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

func (m *PmSecyRxParam) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type PmSecyRxParas struct {
	Index                uint32         `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool           `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            []string       `protobuf:"bytes,52,rep,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        []string       `protobuf:"bytes,53,rep,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   []string       `protobuf:"bytes,54,rep,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear30SecTime   []string       `protobuf:"bytes,55,rep,name=last_clear30_sec_time,json=lastClear30SecTime,proto3" json:"last_clear30_sec_time,omitempty"`
	LastClear24HrTime    []string       `protobuf:"bytes,56,rep,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Sec30Support         bool           `protobuf:"varint,57,opt,name=sec30_support,json=sec30Support,proto3" json:"sec30_support,omitempty"`
	SampleCount          uint64         `protobuf:"varint,58,opt,name=sample_count,json=sampleCount,proto3" json:"sample_count,omitempty"`
	InPktsUnchecked      *PmSecyRxParam `protobuf:"bytes,59,opt,name=in_pkts_unchecked,json=inPktsUnchecked,proto3" json:"in_pkts_unchecked,omitempty"`
	InPktsDelayed        *PmSecyRxParam `protobuf:"bytes,60,opt,name=in_pkts_delayed,json=inPktsDelayed,proto3" json:"in_pkts_delayed,omitempty"`
	InPktsLate           *PmSecyRxParam `protobuf:"bytes,61,opt,name=in_pkts_late,json=inPktsLate,proto3" json:"in_pkts_late,omitempty"`
	InPktsOk             *PmSecyRxParam `protobuf:"bytes,62,opt,name=in_pkts_ok,json=inPktsOk,proto3" json:"in_pkts_ok,omitempty"`
	InPktsInvalid        *PmSecyRxParam `protobuf:"bytes,63,opt,name=in_pkts_invalid,json=inPktsInvalid,proto3" json:"in_pkts_invalid,omitempty"`
	InPktsNotValid       *PmSecyRxParam `protobuf:"bytes,64,opt,name=in_pkts_not_valid,json=inPktsNotValid,proto3" json:"in_pkts_not_valid,omitempty"`
	InPktsNotUsingSa     *PmSecyRxParam `protobuf:"bytes,65,opt,name=in_pkts_not_using_sa,json=inPktsNotUsingSa,proto3" json:"in_pkts_not_using_sa,omitempty"`
	InPktsUnusedSa       *PmSecyRxParam `protobuf:"bytes,66,opt,name=in_pkts_unused_sa,json=inPktsUnusedSa,proto3" json:"in_pkts_unused_sa,omitempty"`
	InPktsUntaggedHit    *PmSecyRxParam `protobuf:"bytes,67,opt,name=in_pkts_untagged_hit,json=inPktsUntaggedHit,proto3" json:"in_pkts_untagged_hit,omitempty"`
	InOctetsValidated    *PmSecyRxParam `protobuf:"bytes,68,opt,name=in_octets_validated,json=inOctetsValidated,proto3" json:"in_octets_validated,omitempty"`
	InOctetsDecrypted    *PmSecyRxParam `protobuf:"bytes,69,opt,name=in_octets_decrypted,json=inOctetsDecrypted,proto3" json:"in_octets_decrypted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PmSecyRxParas) Reset()         { *m = PmSecyRxParas{} }
func (m *PmSecyRxParas) String() string { return proto.CompactTextString(m) }
func (*PmSecyRxParas) ProtoMessage()    {}
func (*PmSecyRxParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e3fa764ed03632f, []int{2}
}

func (m *PmSecyRxParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSecyRxParas.Unmarshal(m, b)
}
func (m *PmSecyRxParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSecyRxParas.Marshal(b, m, deterministic)
}
func (m *PmSecyRxParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSecyRxParas.Merge(m, src)
}
func (m *PmSecyRxParas) XXX_Size() int {
	return xxx_messageInfo_PmSecyRxParas.Size(m)
}
func (m *PmSecyRxParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSecyRxParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmSecyRxParas proto.InternalMessageInfo

func (m *PmSecyRxParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmSecyRxParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmSecyRxParas) GetTimestamp() []string {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PmSecyRxParas) GetLastClearTime() []string {
	if m != nil {
		return m.LastClearTime
	}
	return nil
}

func (m *PmSecyRxParas) GetLastClear15MinTime() []string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return nil
}

func (m *PmSecyRxParas) GetLastClear30SecTime() []string {
	if m != nil {
		return m.LastClear30SecTime
	}
	return nil
}

func (m *PmSecyRxParas) GetLastClear24HrTime() []string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return nil
}

func (m *PmSecyRxParas) GetSec30Support() bool {
	if m != nil {
		return m.Sec30Support
	}
	return false
}

func (m *PmSecyRxParas) GetSampleCount() uint64 {
	if m != nil {
		return m.SampleCount
	}
	return 0
}

func (m *PmSecyRxParas) GetInPktsUnchecked() *PmSecyRxParam {
	if m != nil {
		return m.InPktsUnchecked
	}
	return nil
}

func (m *PmSecyRxParas) GetInPktsDelayed() *PmSecyRxParam {
	if m != nil {
		return m.InPktsDelayed
	}
	return nil
}

func (m *PmSecyRxParas) GetInPktsLate() *PmSecyRxParam {
	if m != nil {
		return m.InPktsLate
	}
	return nil
}

func (m *PmSecyRxParas) GetInPktsOk() *PmSecyRxParam {
	if m != nil {
		return m.InPktsOk
	}
	return nil
}

func (m *PmSecyRxParas) GetInPktsInvalid() *PmSecyRxParam {
	if m != nil {
		return m.InPktsInvalid
	}
	return nil
}

func (m *PmSecyRxParas) GetInPktsNotValid() *PmSecyRxParam {
	if m != nil {
		return m.InPktsNotValid
	}
	return nil
}

func (m *PmSecyRxParas) GetInPktsNotUsingSa() *PmSecyRxParam {
	if m != nil {
		return m.InPktsNotUsingSa
	}
	return nil
}

func (m *PmSecyRxParas) GetInPktsUnusedSa() *PmSecyRxParam {
	if m != nil {
		return m.InPktsUnusedSa
	}
	return nil
}

func (m *PmSecyRxParas) GetInPktsUntaggedHit() *PmSecyRxParam {
	if m != nil {
		return m.InPktsUntaggedHit
	}
	return nil
}

func (m *PmSecyRxParas) GetInOctetsValidated() *PmSecyRxParam {
	if m != nil {
		return m.InOctetsValidated
	}
	return nil
}

func (m *PmSecyRxParas) GetInOctetsDecrypted() *PmSecyRxParam {
	if m != nil {
		return m.InOctetsDecrypted
	}
	return nil
}

func init() {
	proto.RegisterType((*PmSecyRxParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.macsec_history.macsec_port_histories.macsec_port_history.macsec_hour24_history.macsec_hour24secyrx_histories.macsec_hour24secyrx_history.macsec_hour24secyrx_time_line_instances.macsec_hour24secyrx_time_line_instance.pm_secy_rx_paras_KEYS")
	proto.RegisterType((*PmSecyRxParam)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.macsec_history.macsec_port_histories.macsec_port_history.macsec_hour24_history.macsec_hour24secyrx_histories.macsec_hour24secyrx_history.macsec_hour24secyrx_time_line_instances.macsec_hour24secyrx_time_line_instance.pm_secy_rx_param")
	proto.RegisterType((*PmSecyRxParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.macsec_history.macsec_port_histories.macsec_port_history.macsec_hour24_history.macsec_hour24secyrx_histories.macsec_hour24secyrx_history.macsec_hour24secyrx_time_line_instances.macsec_hour24secyrx_time_line_instance.pm_secy_rx_paras")
}

func init() { proto.RegisterFile("pm_secy_rx_paras.proto", fileDescriptor_6e3fa764ed03632f) }

var fileDescriptor_6e3fa764ed03632f = []byte{
	// 743 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0xd8, 0xdd, 0x4e, 0x23, 0x37,
	0x14, 0x00, 0x60, 0x65, 0x48, 0x29, 0x31, 0xa4, 0x80, 0x0b, 0x68, 0x2a, 0xb5, 0x52, 0xa0, 0x52,
	0x95, 0xab, 0x94, 0x24, 0xd0, 0xff, 0x7f, 0x40, 0xa2, 0xea, 0x0f, 0xd5, 0xa4, 0x20, 0xf5, 0xa6,
	0x96, 0xf1, 0x9c, 0x26, 0x56, 0x66, 0xec, 0x91, 0xed, 0xa9, 0x92, 0xe7, 0xd8, 0x27, 0xd8, 0x67,
	0xd8, 0x37, 0xd9, 0x8b, 0xbd, 0x5a, 0x69, 0xf7, 0x62, 0xf7, 0x3d, 0x56, 0xb6, 0x99, 0x24, 0x84,
	0x48, 0xbb, 0x4a, 0x6e, 0xe7, 0x6e, 0x7c, 0x7c, 0xce, 0x89, 0xfd, 0xe1, 0xf1, 0x48, 0xa0, 0x83,
	0x2c, 0x25, 0x1a, 0xd8, 0x98, 0xa8, 0x11, 0xc9, 0xa8, 0xa2, 0xba, 0x95, 0x29, 0x69, 0x24, 0x7e,
	0x1c, 0x30, 0xae, 0x99, 0x24, 0x5c, 0x6a, 0x32, 0x52, 0x24, 0x4b, 0x41, 0xf4, 0xb9, 0x00, 0x22,
	0x33, 0x50, 0xad, 0x0c, 0xd4, 0x7f, 0x52, 0xa5, 0x54, 0x30, 0x20, 0x29, 0x15, 0xb4, 0x0f, 0x29,
	0x08, 0x43, 0x06, 0x5c, 0x1b, 0xa9, 0xc6, 0xad, 0x7e, 0x22, 0x6f, 0x69, 0x62, 0xb3, 0xb8, 0x8c,
	0x39, 0x6b, 0xa5, 0x94, 0x69, 0x60, 0x93, 0xe9, 0xbb, 0x61, 0x26, 0x55, 0x51, 0xc2, 0x41, 0x2f,
	0x88, 0x4e, 0x32, 0x07, 0x32, 0x57, 0x9d, 0x93, 0xc5, 0x51, 0xbb, 0x76, 0x35, 0x7a, 0xd8, 0x67,
	0xc1, 0xec, 0xe2, 0x4a, 0xc3, 0x53, 0x20, 0x89, 0xdd, 0x1a, 0x17, 0xda, 0xd8, 0x2d, 0xe9, 0x77,
	0xcc, 0x3b, 0xfa, 0x17, 0xed, 0xcf, 0xeb, 0x91, 0xdf, 0x2e, 0xfe, 0xe9, 0x61, 0x8c, 0xaa, 0x82,
	0xa6, 0x10, 0x56, 0x1a, 0x95, 0x66, 0x2d, 0x72, 0xcf, 0xf8, 0x00, 0xad, 0x8b, 0x3c, 0xbd, 0x05,
	0x15, 0x06, 0x8d, 0x4a, 0xb3, 0x1e, 0xdd, 0x8d, 0xf0, 0x47, 0x68, 0xc3, 0x3f, 0x91, 0x76, 0xb8,
	0xe6, 0x66, 0xde, 0xf7, 0xe3, 0xf6, 0xd1, 0x18, 0xed, 0xcc, 0xf5, 0x4f, 0x6d, 0xeb, 0x98, 0x1a,
	0xea, 0x5a, 0x57, 0x23, 0xf7, 0x8c, 0x3f, 0x46, 0x35, 0x33, 0x50, 0xa0, 0x07, 0x32, 0x89, 0x5d,
	0xf7, 0x6a, 0x34, 0x0d, 0xe0, 0x4f, 0x10, 0x32, 0x8c, 0x12, 0x05, 0x96, 0xd6, 0xfd, 0xc4, 0x46,
	0x54, 0x33, 0x8c, 0x46, 0x2e, 0x80, 0xf7, 0xd0, 0x7b, 0xff, 0xd3, 0x84, 0xc7, 0x61, 0xd5, 0xcd,
	0xf8, 0xc1, 0xd1, 0xa3, 0xc3, 0x07, 0xbf, 0xad, 0x6d, 0x2a, 0x17, 0x31, 0x8c, 0xc2, 0x8e, 0x5b,
	0xa7, 0x1f, 0x4c, 0x1b, 0x74, 0x67, 0x1a, 0xb8, 0x35, 0xf1, 0x14, 0xb4, 0xa1, 0x69, 0x16, 0x9e,
	0x34, 0xd6, 0x9a, 0xb5, 0x68, 0x1a, 0xc0, 0x9f, 0xa1, 0xed, 0x84, 0x6a, 0x43, 0x58, 0x02, 0x54,
	0x39, 0xda, 0xf0, 0xd4, 0xe5, 0xd4, 0x6d, 0xf8, 0xcc, 0x46, 0xff, 0xe6, 0x29, 0xe0, 0x36, 0xda,
	0x9f, 0xe6, 0xb5, 0x4f, 0x49, 0xca, 0x85, 0xcf, 0xfe, 0xc2, 0x65, 0xe3, 0x49, 0x76, 0xfb, 0xf4,
	0x0f, 0x2e, 0x1e, 0x96, 0x74, 0x8f, 0xed, 0x16, 0x7c, 0xc9, 0x97, 0x73, 0x25, 0xdd, 0xe3, 0x1e,
	0x30, 0x57, 0xf2, 0x39, 0xda, 0x9b, 0x96, 0xd8, 0x83, 0x76, 0xb7, 0xa4, 0xaf, 0x5c, 0xc5, 0xee,
	0xa4, 0xa2, 0x73, 0x72, 0xe9, 0x97, 0xf5, 0x29, 0xaa, 0x6b, 0x60, 0xb6, 0x79, 0x9e, 0x39, 0xd5,
	0xaf, 0xdd, 0xd6, 0xb7, 0x5c, 0xb0, 0xe7, 0x63, 0xf8, 0x10, 0x6d, 0x69, 0x9a, 0x66, 0x09, 0x10,
	0x26, 0x73, 0x61, 0xc2, 0x6f, 0xdc, 0x1f, 0x66, 0xd3, 0xc7, 0xce, 0x6c, 0x08, 0xbf, 0x0c, 0xd0,
	0x2e, 0x17, 0x24, 0x1b, 0x1a, 0x4d, 0x72, 0xc1, 0x06, 0xc0, 0x86, 0x10, 0x87, 0xdf, 0x36, 0x2a,
	0xcd, 0xcd, 0xce, 0x93, 0xa0, 0x55, 0xbe, 0x81, 0x6f, 0xc9, 0x6b, 0xcd, 0xbf, 0x1e, 0xd1, 0x36,
	0x17, 0x7f, 0x0d, 0x8d, 0xbe, 0x2e, 0x30, 0xf1, 0xf3, 0x00, 0x6d, 0x17, 0xc4, 0x31, 0x24, 0x74,
	0x0c, 0x71, 0xf8, 0x5d, 0x09, 0xbc, 0x34, 0x70, 0xdd, 0x03, 0x9f, 0x7b, 0x4a, 0xfc, 0x2c, 0x40,
	0x5b, 0x05, 0x6f, 0x42, 0x0d, 0x84, 0xdf, 0x97, 0xb6, 0x4b, 0xdb, 0x22, 0x6f, 0xfb, 0x3b, 0x35,
	0x80, 0x9f, 0x06, 0x08, 0x15, 0xb0, 0x72, 0x18, 0xfe, 0x50, 0xb2, 0x2e, 0xcd, 0xba, 0xe1, 0x59,
	0xaf, 0x86, 0xf7, 0x2e, 0x03, 0x2e, 0xfc, 0x57, 0xeb, 0xc7, 0x52, 0x76, 0xd5, 0xcb, 0xe0, 0x57,
	0x4f, 0x89, 0x5f, 0xcc, 0x7c, 0xce, 0x84, 0x34, 0xc4, 0x03, 0xff, 0x54, 0x02, 0x2f, 0x0d, 0xfc,
	0x81, 0x07, 0xfe, 0x53, 0x9a, 0x1b, 0x27, 0xfc, 0x2a, 0x40, 0x7b, 0xb3, 0xc2, 0xb9, 0xe6, 0xa2,
	0x4f, 0x34, 0x0d, 0x7f, 0x2e, 0x91, 0x97, 0x46, 0xde, 0x99, 0x20, 0x5f, 0x5b, 0xce, 0x1e, 0xbd,
	0x77, 0x90, 0x73, 0x91, 0x6b, 0x88, 0xad, 0xf1, 0x2f, 0xa5, 0xf1, 0xaa, 0x07, 0xf9, 0xda, 0x61,
	0xf6, 0x28, 0x7e, 0x3d, 0x73, 0x90, 0x73, 0x61, 0x68, 0xbf, 0x0f, 0x31, 0x19, 0x70, 0x13, 0x9e,
	0x95, 0xc8, 0x4b, 0x23, 0xef, 0x16, 0xc8, 0xde, 0xf3, 0x92, 0x1b, 0x7b, 0x61, 0x7c, 0xc8, 0x05,
	0x91, 0xcc, 0x80, 0xd1, 0xfe, 0x42, 0xa6, 0x06, 0xe2, 0xf0, 0xbc, 0x64, 0x5e, 0x85, 0xf9, 0xca,
	0x79, 0xde, 0x14, 0x9c, 0x73, 0xcc, 0x31, 0x30, 0x35, 0xce, 0x2c, 0xf3, 0x45, 0xc9, 0xbc, 0x3a,
	0xf3, 0x79, 0xc1, 0x79, 0xbb, 0xee, 0xfe, 0x37, 0xd5, 0x7d, 0x13, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0x40, 0xa1, 0xc6, 0xb5, 0x12, 0x00, 0x00,
}
