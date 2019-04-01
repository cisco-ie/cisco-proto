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
// source: pm_fec_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_dwdm_dwdm_ports_dwdm_port_dwdm_current_dwdm_minute15_dwdm_minute15fecs_dwdm_minute15fec

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

type PmFecParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmFecParas_KEYS) Reset()         { *m = PmFecParas_KEYS{} }
func (m *PmFecParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmFecParas_KEYS) ProtoMessage()    {}
func (*PmFecParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e952374f03679f, []int{0}
}

func (m *PmFecParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmFecParas_KEYS.Unmarshal(m, b)
}
func (m *PmFecParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmFecParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmFecParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmFecParas_KEYS.Merge(m, src)
}
func (m *PmFecParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmFecParas_KEYS.Size(m)
}
func (m *PmFecParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmFecParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmFecParas_KEYS proto.InternalMessageInfo

func (m *PmFecParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmFecParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PmFecParam struct {
	Data                 uint64   `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            uint64   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,3,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	Valid                bool     `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmFecParam) Reset()         { *m = PmFecParam{} }
func (m *PmFecParam) String() string { return proto.CompactTextString(m) }
func (*PmFecParam) ProtoMessage()    {}
func (*PmFecParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e952374f03679f, []int{1}
}

func (m *PmFecParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmFecParam.Unmarshal(m, b)
}
func (m *PmFecParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmFecParam.Marshal(b, m, deterministic)
}
func (m *PmFecParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmFecParam.Merge(m, src)
}
func (m *PmFecParam) XXX_Size() int {
	return xxx_messageInfo_PmFecParam.Size(m)
}
func (m *PmFecParam) XXX_DiscardUnknown() {
	xxx_messageInfo_PmFecParam.DiscardUnknown(m)
}

var xxx_messageInfo_PmFecParam proto.InternalMessageInfo

func (m *PmFecParam) GetData() uint64 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PmFecParam) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PmFecParam) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

func (m *PmFecParam) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type PmFecParamString struct {
	Minimum              string   `protobuf:"bytes,1,opt,name=minimum,proto3" json:"minimum,omitempty"`
	Average              string   `protobuf:"bytes,2,opt,name=average,proto3" json:"average,omitempty"`
	Maximum              string   `protobuf:"bytes,3,opt,name=maximum,proto3" json:"maximum,omitempty"`
	MinimumThreshold     string   `protobuf:"bytes,4,opt,name=minimum_threshold,json=minimumThreshold,proto3" json:"minimum_threshold,omitempty"`
	MinimumTcaReport     bool     `protobuf:"varint,5,opt,name=minimum_tca_report,json=minimumTcaReport,proto3" json:"minimum_tca_report,omitempty"`
	MaximumThreshold     string   `protobuf:"bytes,6,opt,name=maximum_threshold,json=maximumThreshold,proto3" json:"maximum_threshold,omitempty"`
	MaximumTcaReport     bool     `protobuf:"varint,7,opt,name=maximum_tca_report,json=maximumTcaReport,proto3" json:"maximum_tca_report,omitempty"`
	Valid                bool     `protobuf:"varint,8,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmFecParamString) Reset()         { *m = PmFecParamString{} }
func (m *PmFecParamString) String() string { return proto.CompactTextString(m) }
func (*PmFecParamString) ProtoMessage()    {}
func (*PmFecParamString) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e952374f03679f, []int{2}
}

func (m *PmFecParamString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmFecParamString.Unmarshal(m, b)
}
func (m *PmFecParamString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmFecParamString.Marshal(b, m, deterministic)
}
func (m *PmFecParamString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmFecParamString.Merge(m, src)
}
func (m *PmFecParamString) XXX_Size() int {
	return xxx_messageInfo_PmFecParamString.Size(m)
}
func (m *PmFecParamString) XXX_DiscardUnknown() {
	xxx_messageInfo_PmFecParamString.DiscardUnknown(m)
}

var xxx_messageInfo_PmFecParamString proto.InternalMessageInfo

func (m *PmFecParamString) GetMinimum() string {
	if m != nil {
		return m.Minimum
	}
	return ""
}

func (m *PmFecParamString) GetAverage() string {
	if m != nil {
		return m.Average
	}
	return ""
}

func (m *PmFecParamString) GetMaximum() string {
	if m != nil {
		return m.Maximum
	}
	return ""
}

func (m *PmFecParamString) GetMinimumThreshold() string {
	if m != nil {
		return m.MinimumThreshold
	}
	return ""
}

func (m *PmFecParamString) GetMinimumTcaReport() bool {
	if m != nil {
		return m.MinimumTcaReport
	}
	return false
}

func (m *PmFecParamString) GetMaximumThreshold() string {
	if m != nil {
		return m.MaximumThreshold
	}
	return ""
}

func (m *PmFecParamString) GetMaximumTcaReport() bool {
	if m != nil {
		return m.MaximumTcaReport
	}
	return false
}

func (m *PmFecParamString) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type PmFecParas struct {
	Index                uint32            `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool              `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            string            `protobuf:"bytes,52,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        string            `protobuf:"bytes,53,opt,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   string            `protobuf:"bytes,54,opt,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear30SecTime   string            `protobuf:"bytes,55,opt,name=last_clear30_sec_time,json=lastClear30SecTime,proto3" json:"last_clear30_sec_time,omitempty"`
	LastClear24HrTime    string            `protobuf:"bytes,56,opt,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Sec30Support         bool              `protobuf:"varint,57,opt,name=sec30_support,json=sec30Support,proto3" json:"sec30_support,omitempty"`
	EcBits               *PmFecParam       `protobuf:"bytes,58,opt,name=ec_bits,json=ecBits,proto3" json:"ec_bits,omitempty"`
	UcWords              *PmFecParam       `protobuf:"bytes,59,opt,name=uc_words,json=ucWords,proto3" json:"uc_words,omitempty"`
	PreFecBer            *PmFecParamString `protobuf:"bytes,60,opt,name=pre_fec_ber,json=preFecBer,proto3" json:"pre_fec_ber,omitempty"`
	PostFecBer           *PmFecParamString `protobuf:"bytes,61,opt,name=post_fec_ber,json=postFecBer,proto3" json:"post_fec_ber,omitempty"`
	Q                    *PmFecParamString `protobuf:"bytes,62,opt,name=q,proto3" json:"q,omitempty"`
	Qmargin              *PmFecParamString `protobuf:"bytes,63,opt,name=qmargin,proto3" json:"qmargin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PmFecParas) Reset()         { *m = PmFecParas{} }
func (m *PmFecParas) String() string { return proto.CompactTextString(m) }
func (*PmFecParas) ProtoMessage()    {}
func (*PmFecParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e952374f03679f, []int{3}
}

func (m *PmFecParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmFecParas.Unmarshal(m, b)
}
func (m *PmFecParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmFecParas.Marshal(b, m, deterministic)
}
func (m *PmFecParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmFecParas.Merge(m, src)
}
func (m *PmFecParas) XXX_Size() int {
	return xxx_messageInfo_PmFecParas.Size(m)
}
func (m *PmFecParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmFecParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmFecParas proto.InternalMessageInfo

func (m *PmFecParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmFecParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmFecParas) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *PmFecParas) GetLastClearTime() string {
	if m != nil {
		return m.LastClearTime
	}
	return ""
}

func (m *PmFecParas) GetLastClear15MinTime() string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return ""
}

func (m *PmFecParas) GetLastClear30SecTime() string {
	if m != nil {
		return m.LastClear30SecTime
	}
	return ""
}

func (m *PmFecParas) GetLastClear24HrTime() string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return ""
}

func (m *PmFecParas) GetSec30Support() bool {
	if m != nil {
		return m.Sec30Support
	}
	return false
}

func (m *PmFecParas) GetEcBits() *PmFecParam {
	if m != nil {
		return m.EcBits
	}
	return nil
}

func (m *PmFecParas) GetUcWords() *PmFecParam {
	if m != nil {
		return m.UcWords
	}
	return nil
}

func (m *PmFecParas) GetPreFecBer() *PmFecParamString {
	if m != nil {
		return m.PreFecBer
	}
	return nil
}

func (m *PmFecParas) GetPostFecBer() *PmFecParamString {
	if m != nil {
		return m.PostFecBer
	}
	return nil
}

func (m *PmFecParas) GetQ() *PmFecParamString {
	if m != nil {
		return m.Q
	}
	return nil
}

func (m *PmFecParas) GetQmargin() *PmFecParamString {
	if m != nil {
		return m.Qmargin
	}
	return nil
}

func init() {
	proto.RegisterType((*PmFecParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management.dwdm.dwdm_ports.dwdm_port.dwdm_current.dwdm_minute15.dwdm_minute15fecs.dwdm_minute15fec.pm_fec_paras_KEYS")
	proto.RegisterType((*PmFecParam)(nil), "cisco_ios_xr_pmengine_oper.performance_management.dwdm.dwdm_ports.dwdm_port.dwdm_current.dwdm_minute15.dwdm_minute15fecs.dwdm_minute15fec.pm_fec_param")
	proto.RegisterType((*PmFecParamString)(nil), "cisco_ios_xr_pmengine_oper.performance_management.dwdm.dwdm_ports.dwdm_port.dwdm_current.dwdm_minute15.dwdm_minute15fecs.dwdm_minute15fec.pm_fec_param_string")
	proto.RegisterType((*PmFecParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management.dwdm.dwdm_ports.dwdm_port.dwdm_current.dwdm_minute15.dwdm_minute15fecs.dwdm_minute15fec.pm_fec_paras")
}

func init() { proto.RegisterFile("pm_fec_paras.proto", fileDescriptor_f9e952374f03679f) }

var fileDescriptor_f9e952374f03679f = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6a, 0x14, 0x31,
	0x14, 0x26, 0xed, 0x76, 0xb7, 0x9b, 0xb6, 0x68, 0x63, 0x95, 0xb9, 0x50, 0x58, 0x2a, 0xc8, 0x82,
	0xb2, 0x76, 0x77, 0xbb, 0xfe, 0x6b, 0xa1, 0xa2, 0x08, 0xe2, 0xcd, 0xb4, 0x20, 0x5e, 0x85, 0x34,
	0x73, 0xba, 0x0d, 0x6c, 0x32, 0xd3, 0x24, 0xd3, 0xf6, 0x15, 0x7c, 0x02, 0xbd, 0xf2, 0x0d, 0x44,
	0xc4, 0x2b, 0xaf, 0xc4, 0x37, 0x93, 0x9c, 0x99, 0x9d, 0x99, 0xda, 0x17, 0x68, 0x6f, 0x96, 0x9c,
	0xef, 0x7c, 0xdf, 0x99, 0xef, 0x4b, 0x76, 0xc2, 0x50, 0x96, 0x69, 0x7e, 0x08, 0x92, 0x67, 0xc2,
	0x0a, 0x37, 0xc8, 0x6c, 0xea, 0x53, 0xf6, 0x99, 0x48, 0xe5, 0x64, 0xca, 0x55, 0xea, 0xf8, 0x99,
	0xe5, 0x99, 0x06, 0x33, 0x55, 0x06, 0x78, 0x9a, 0x81, 0x1d, 0x64, 0x60, 0x0f, 0x53, 0xab, 0x85,
	0x91, 0xc0, 0xb5, 0x30, 0x62, 0x0a, 0x1a, 0x8c, 0x1f, 0x24, 0xa7, 0x89, 0xc6, 0x1f, 0x9e, 0xa5,
	0xd6, 0xbb, 0x7a, 0x59, 0xac, 0x64, 0x6e, 0xed, 0x9c, 0xc6, 0xb5, 0x32, 0xb9, 0x87, 0xe1, 0xe4,
	0x7c, 0x75, 0x08, 0xd2, 0x5d, 0x40, 0x36, 0x77, 0xe8, 0x7a, 0xd3, 0x21, 0x7f, 0xff, 0xe6, 0xd3,
	0x1e, 0x63, 0xb4, 0x65, 0x84, 0x86, 0x88, 0xf4, 0x48, 0xbf, 0x1b, 0xe3, 0x9a, 0xdd, 0xa2, 0x6d,
	0x93, 0xeb, 0x03, 0xb0, 0xd1, 0x42, 0x8f, 0xf4, 0xd7, 0xe2, 0xb2, 0xda, 0xcc, 0xe9, 0x6a, 0x63,
	0x80, 0x0e, 0xda, 0x44, 0x78, 0x81, 0xda, 0x56, 0x8c, 0x6b, 0x76, 0x9b, 0x76, 0xfd, 0x91, 0x05,
	0x77, 0x94, 0xce, 0x12, 0x94, 0xb7, 0xe2, 0x1a, 0x60, 0x77, 0x28, 0xf5, 0x52, 0x70, 0x0b, 0x21,
	0x4e, 0xb4, 0xd8, 0x23, 0xfd, 0xe5, 0xb8, 0xeb, 0xa5, 0x88, 0x11, 0x60, 0x1b, 0x74, 0xe9, 0x44,
	0xcc, 0x54, 0x12, 0xb5, 0xb0, 0x53, 0x14, 0x9b, 0x3f, 0x16, 0xe8, 0x8d, 0xe6, 0x73, 0xb9, 0xf3,
	0x56, 0x99, 0x29, 0x8b, 0x68, 0x47, 0x2b, 0xa3, 0x74, 0xae, 0x4b, 0xf7, 0xf3, 0x32, 0x74, 0xc4,
	0x09, 0x58, 0x31, 0x05, 0xb4, 0xd0, 0x8d, 0xe7, 0x25, 0x6a, 0xc4, 0x19, 0x6a, 0x16, 0x4b, 0x4d,
	0x51, 0xb2, 0xfb, 0x74, 0xbd, 0x94, 0xf3, 0x3a, 0x40, 0x0b, 0x39, 0xd7, 0xcb, 0xc6, 0x7e, 0x95,
	0xe3, 0x01, 0x65, 0x15, 0xb9, 0xce, 0xb3, 0x84, 0xae, 0x2b, 0x76, 0x15, 0x2b, 0x8c, 0x2e, 0x9e,
	0xd2, 0x18, 0xdd, 0x2e, 0x47, 0x17, 0x8d, 0xf3, 0xa3, 0xe7, 0xe4, 0x7a, 0x74, 0xa7, 0x1c, 0x5d,
	0xb2, 0x2f, 0xee, 0xd8, 0x72, 0x73, 0xc7, 0xbe, 0xae, 0x9c, 0x3b, 0x29, 0x17, 0x68, 0xca, 0x24,
	0x70, 0x16, 0x8d, 0xf0, 0x40, 0x8b, 0xa2, 0x16, 0x8f, 0x1b, 0x62, 0x3c, 0x41, 0xa5, 0xc1, 0x79,
	0xa1, 0xb3, 0x68, 0x1b, 0x5d, 0xd6, 0x00, 0xbb, 0x47, 0xaf, 0xcd, 0x84, 0xf3, 0x5c, 0xce, 0x40,
	0x58, 0x1e, 0xf0, 0x68, 0x82, 0x9c, 0xb5, 0x00, 0xbf, 0x0e, 0xe8, 0xbe, 0xd2, 0xc0, 0x86, 0xf4,
	0x66, 0xcd, 0x1b, 0x4e, 0xc2, 0x1f, 0xb1, 0x60, 0x3f, 0x42, 0x36, 0xab, 0xd8, 0xc3, 0xc9, 0x07,
	0x65, 0x2e, 0x4a, 0xc6, 0x5b, 0xdc, 0x81, 0x2c, 0x24, 0x8f, 0xff, 0x93, 0x8c, 0xb7, 0xf6, 0x40,
	0xa2, 0xe4, 0x21, 0xdd, 0xa8, 0x25, 0xa3, 0x6d, 0x7e, 0x54, 0x5a, 0x7a, 0x82, 0x8a, 0xf5, 0x4a,
	0x31, 0xda, 0x7e, 0x57, 0xd8, 0xba, 0x4b, 0xd7, 0x1c, 0xc8, 0x30, 0x3c, 0xcf, 0x70, 0x63, 0x9f,
	0x62, 0xf4, 0x55, 0x04, 0xf7, 0x0a, 0x8c, 0xfd, 0x24, 0xb4, 0x03, 0x92, 0x1f, 0x28, 0xef, 0xa2,
	0x67, 0x3d, 0xd2, 0x5f, 0x19, 0x7d, 0x21, 0x83, 0x4b, 0xf3, 0x1e, 0x0f, 0x9a, 0xef, 0x42, 0xdc,
	0x06, 0xb9, 0xab, 0xbc, 0x63, 0xbf, 0x08, 0x5d, 0xce, 0x25, 0x3f, 0x4d, 0x6d, 0xe2, 0xa2, 0xe7,
	0x97, 0xdc, 0x74, 0x27, 0x97, 0x1f, 0x83, 0x51, 0xf6, 0x87, 0xd0, 0x95, 0xcc, 0x02, 0xb6, 0xc2,
	0x7d, 0xf3, 0x02, 0x8d, 0x7f, 0xbb, 0xac, 0xc6, 0xcb, 0x9b, 0x27, 0xee, 0x66, 0x16, 0xde, 0x82,
	0xdc, 0x05, 0xcb, 0xfe, 0x12, 0xba, 0x9a, 0xa5, 0xce, 0x57, 0x11, 0x5e, 0x5e, 0x8d, 0x08, 0x34,
	0x98, 0x2e, 0x33, 0x7c, 0x27, 0x94, 0x1c, 0x47, 0xaf, 0xae, 0x86, 0x71, 0x72, 0xcc, 0x7e, 0x13,
	0xda, 0x39, 0xd6, 0xc2, 0x4e, 0x95, 0x89, 0x76, 0xae, 0x86, 0xeb, 0xb9, 0xdf, 0x83, 0x36, 0x7e,
	0x17, 0x8c, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x48, 0xe2, 0xef, 0x46, 0x2d, 0x08, 0x00, 0x00,
}