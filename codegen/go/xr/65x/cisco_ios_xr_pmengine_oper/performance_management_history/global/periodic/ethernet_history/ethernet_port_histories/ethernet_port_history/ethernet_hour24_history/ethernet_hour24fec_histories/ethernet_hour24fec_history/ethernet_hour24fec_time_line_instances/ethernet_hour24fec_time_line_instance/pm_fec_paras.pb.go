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

package cisco_ios_xr_pmengine_oper_performance_management_history_global_periodic_ethernet_history_ethernet_port_histories_ethernet_port_history_ethernet_hour24_history_ethernet_hour24fec_histories_ethernet_hour24fec_history_ethernet_hour24fec_time_line_instances_ethernet_hour24fec_time_line_instance

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
	Number_1             uint32   `protobuf:"varint,3,opt,name=number_1,json=number1,proto3" json:"number_1,omitempty"`
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

func (m *PmFecParas_KEYS) GetNumber_1() uint32 {
	if m != nil {
		return m.Number_1
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
	proto.RegisterType((*PmFecParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ethernet_history.ethernet_port_histories.ethernet_port_history.ethernet_hour24_history.ethernet_hour24fec_histories.ethernet_hour24fec_history.ethernet_hour24fec_time_line_instances.ethernet_hour24fec_time_line_instance.pm_fec_paras_KEYS")
	proto.RegisterType((*PmFecParam)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ethernet_history.ethernet_port_histories.ethernet_port_history.ethernet_hour24_history.ethernet_hour24fec_histories.ethernet_hour24fec_history.ethernet_hour24fec_time_line_instances.ethernet_hour24fec_time_line_instance.pm_fec_param")
	proto.RegisterType((*PmFecParamString)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ethernet_history.ethernet_port_histories.ethernet_port_history.ethernet_hour24_history.ethernet_hour24fec_histories.ethernet_hour24fec_history.ethernet_hour24fec_time_line_instances.ethernet_hour24fec_time_line_instance.pm_fec_param_string")
	proto.RegisterType((*PmFecParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ethernet_history.ethernet_port_histories.ethernet_port_history.ethernet_hour24_history.ethernet_hour24fec_histories.ethernet_hour24fec_history.ethernet_hour24fec_time_line_instances.ethernet_hour24fec_time_line_instance.pm_fec_paras")
}

func init() { proto.RegisterFile("pm_fec_paras.proto", fileDescriptor_f9e952374f03679f) }

var fileDescriptor_f9e952374f03679f = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x4b, 0x6b, 0x14, 0x4d,
	0x14, 0x65, 0x3a, 0x93, 0x79, 0x54, 0x12, 0xbe, 0x2f, 0x65, 0x94, 0x16, 0x14, 0x42, 0x04, 0x09,
	0x28, 0x63, 0x66, 0x26, 0xf1, 0xfd, 0x80, 0x88, 0x22, 0x88, 0x9b, 0x4e, 0x40, 0x74, 0x53, 0xd4,
	0xd4, 0xdc, 0xcc, 0x14, 0x4c, 0x55, 0x77, 0xaa, 0xaa, 0x63, 0xfc, 0x41, 0xae, 0x5d, 0xb8, 0x73,
	0xe5, 0xd6, 0x85, 0xb8, 0x72, 0x25, 0xfe, 0x17, 0xa9, 0xdb, 0x3d, 0xdd, 0x9d, 0x87, 0x10, 0xcd,
	0x76, 0x76, 0x7d, 0xcf, 0x3d, 0xe7, 0xf4, 0xbd, 0xa7, 0xba, 0x7b, 0x18, 0x42, 0x13, 0xc5, 0xf6,
	0x40, 0xb0, 0x84, 0x1b, 0x6e, 0x3b, 0x89, 0x89, 0x5d, 0x4c, 0x3f, 0x04, 0x42, 0x5a, 0x11, 0x33,
	0x19, 0x5b, 0x76, 0x68, 0x58, 0xa2, 0x40, 0x8f, 0xa4, 0x06, 0x16, 0x27, 0x60, 0x3a, 0x09, 0x98,
	0xbd, 0xd8, 0x28, 0xae, 0x05, 0x30, 0xc5, 0x35, 0x1f, 0x81, 0x02, 0xed, 0xd8, 0x58, 0x5a, 0x17,
	0x9b, 0xf7, 0x9d, 0xd1, 0x24, 0x1e, 0xf0, 0x89, 0x67, 0xc9, 0x78, 0x28, 0x45, 0x07, 0xdc, 0x18,
	0x8c, 0x86, 0x92, 0x50, 0x00, 0x49, 0x6c, 0xa6, 0xa8, 0x04, 0x7b, 0x2a, 0x5e, 0x61, 0x8f, 0xe3,
	0xd4, 0xf4, 0x36, 0xff, 0x84, 0xfb, 0x0d, 0x4e, 0xb1, 0x3a, 0xde, 0x3c, 0x55, 0xe7, 0xa4, 0x02,
	0x36, 0xf1, 0xfb, 0x49, 0x6d, 0x9d, 0xdf, 0xcb, 0x9e, 0x8d, 0xb6, 0xf6, 0x96, 0x2c, 0x57, 0xd3,
	0x63, 0x2f, 0x9f, 0xbd, 0xd9, 0xa1, 0x94, 0xd4, 0x35, 0x57, 0x10, 0xd6, 0x56, 0x6b, 0xeb, 0xed,
	0x08, 0xaf, 0xe9, 0x25, 0xd2, 0xd0, 0xa9, 0x1a, 0x80, 0x09, 0x83, 0xd5, 0xda, 0xfa, 0x52, 0x94,
	0x57, 0xf4, 0x32, 0x69, 0x65, 0x57, 0xac, 0x1b, 0xce, 0x61, 0xa7, 0x99, 0xd5, 0xdd, 0xb5, 0x94,
	0x2c, 0x56, 0xbc, 0x95, 0xb7, 0x1d, 0x72, 0xc7, 0xd1, 0xb6, 0x1e, 0xe1, 0x35, 0xbd, 0x42, 0xda,
	0x6e, 0x6c, 0xc0, 0x8e, 0xe3, 0xc9, 0x10, 0x9d, 0xeb, 0x51, 0x09, 0xd0, 0xab, 0x84, 0x38, 0xc1,
	0x99, 0x01, 0x1f, 0x28, 0xda, 0xb7, 0xa2, 0xb6, 0x13, 0x3c, 0x42, 0x80, 0xae, 0x90, 0xf9, 0x03,
	0x3e, 0x91, 0xc3, 0xb0, 0x8e, 0x9d, 0xac, 0x58, 0xfb, 0x18, 0x90, 0x0b, 0xd5, 0xfb, 0x32, 0xeb,
	0x8c, 0xd4, 0x23, 0x1a, 0x92, 0xa6, 0x92, 0x5a, 0xaa, 0x54, 0xe5, 0x8b, 0x4d, 0x4b, 0xdf, 0xe1,
	0x07, 0x60, 0xf8, 0x08, 0x70, 0x84, 0x76, 0x34, 0x2d, 0x51, 0xc3, 0x0f, 0x51, 0x33, 0x97, 0x6b,
	0xb2, 0x92, 0xde, 0x20, 0xcb, 0xb9, 0x9c, 0x95, 0x0b, 0xd4, 0x91, 0xf3, 0x7f, 0xde, 0xd8, 0x2d,
	0xf6, 0xb8, 0x49, 0x68, 0x41, 0x2e, 0xf7, 0x99, 0xc7, 0xa9, 0x0b, 0x76, 0xb1, 0x96, 0xb7, 0xce,
	0xee, 0x52, 0xb1, 0x6e, 0xe4, 0xd6, 0x59, 0xe3, 0xa8, 0xf5, 0x94, 0x5c, 0x5a, 0x37, 0x73, 0xeb,
	0x9c, 0x7d, 0x32, 0xb1, 0x56, 0x35, 0xb1, 0x2f, 0xf4, 0xc8, 0x49, 0x59, 0x4f, 0x93, 0x7a, 0x08,
	0x87, 0x61, 0x0f, 0x4f, 0x34, 0x2b, 0x4a, 0x71, 0xbf, 0x22, 0xc6, 0x13, 0x94, 0x0a, 0xac, 0xe3,
	0x2a, 0x09, 0x37, 0x71, 0xca, 0x12, 0xa0, 0xd7, 0xc9, 0x7f, 0x13, 0x6e, 0x1d, 0x13, 0x13, 0xe0,
	0x06, 0x1f, 0xc0, 0x70, 0x0b, 0x39, 0x4b, 0x1e, 0x7e, 0xea, 0xd1, 0x5d, 0xa9, 0x80, 0x76, 0xc9,
	0xc5, 0x92, 0xd7, 0xdd, 0x62, 0x4a, 0xea, 0x8c, 0x7d, 0x1b, 0xd9, 0xb4, 0x60, 0x77, 0xb7, 0x5e,
	0x49, 0x7d, 0x52, 0xd2, 0xdf, 0x60, 0x36, 0x7f, 0xc2, 0xc3, 0x3b, 0xc7, 0x24, 0xfd, 0x8d, 0x1d,
	0x10, 0x28, 0xb9, 0x45, 0x56, 0x4a, 0x89, 0x7f, 0x15, 0xf3, 0x91, 0xee, 0xa2, 0x62, 0xb9, 0x50,
	0xf4, 0x36, 0x5f, 0x64, 0x63, 0x5d, 0x23, 0x4b, 0x16, 0x84, 0x37, 0x4f, 0x13, 0x0c, 0xf6, 0x1e,
	0xae, 0xbe, 0x88, 0xe0, 0x4e, 0x86, 0xd1, 0x6f, 0x01, 0x69, 0x82, 0x60, 0x03, 0xe9, 0x6c, 0x78,
	0x7f, 0xb5, 0xb6, 0xbe, 0xd0, 0xfb, 0x14, 0x74, 0x66, 0x9f, 0x9f, 0x33, 0xd0, 0x3a, 0xd5, 0xf7,
	0x34, 0x6a, 0x80, 0xd8, 0x96, 0xce, 0xd2, 0xef, 0x01, 0x69, 0xa5, 0x82, 0xbd, 0x8b, 0xcd, 0xd0,
	0x86, 0x0f, 0x66, 0x81, 0xfe, 0x73, 0xa0, 0xcd, 0x54, 0xbc, 0xf6, 0x21, 0xd2, 0x9f, 0x01, 0x59,
	0x48, 0x0c, 0x60, 0xcb, 0x7f, 0xc2, 0x1f, 0x62, 0xa8, 0x9f, 0x67, 0xa1, 0xfe, 0x7d, 0xa8, 0xf9,
	0xaf, 0x49, 0xd4, 0x4e, 0x0c, 0x3c, 0x07, 0xb1, 0x0d, 0x86, 0xfe, 0x0a, 0xc8, 0x62, 0x12, 0x5b,
	0x57, 0xc4, 0xfb, 0x68, 0x16, 0xef, 0x79, 0xe3, 0x25, 0x3e, 0xd0, 0x3c, 0xdf, 0xaf, 0x01, 0xa9,
	0xed, 0x87, 0x8f, 0x67, 0xa1, 0x9e, 0x37, 0xd4, 0xda, 0x3e, 0xfd, 0x11, 0x90, 0xe6, 0xbe, 0xe2,
	0x66, 0x24, 0x75, 0xf8, 0x64, 0x96, 0xe8, 0x79, 0x13, 0x9d, 0x66, 0x39, 0x68, 0xe0, 0xdf, 0x8e,
	0xfe, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xf5, 0xc9, 0x6e, 0x8c, 0x0c, 0x00, 0x00,
}