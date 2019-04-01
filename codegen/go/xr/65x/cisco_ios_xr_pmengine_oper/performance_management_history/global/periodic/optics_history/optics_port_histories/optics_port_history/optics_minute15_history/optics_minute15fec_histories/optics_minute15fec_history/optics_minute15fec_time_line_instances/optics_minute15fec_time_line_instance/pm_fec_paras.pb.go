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

package cisco_ios_xr_pmengine_oper_performance_management_history_global_periodic_optics_history_optics_port_histories_optics_port_history_optics_minute15_history_optics_minute15fec_histories_optics_minute15fec_history_optics_minute15fec_time_line_instances_optics_minute15fec_time_line_instance

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
	proto.RegisterType((*PmFecParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.optics_history.optics_port_histories.optics_port_history.optics_minute15_history.optics_minute15fec_histories.optics_minute15fec_history.optics_minute15fec_time_line_instances.optics_minute15fec_time_line_instance.pm_fec_paras_KEYS")
	proto.RegisterType((*PmFecParam)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.optics_history.optics_port_histories.optics_port_history.optics_minute15_history.optics_minute15fec_histories.optics_minute15fec_history.optics_minute15fec_time_line_instances.optics_minute15fec_time_line_instance.pm_fec_param")
	proto.RegisterType((*PmFecParamString)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.optics_history.optics_port_histories.optics_port_history.optics_minute15_history.optics_minute15fec_histories.optics_minute15fec_history.optics_minute15fec_time_line_instances.optics_minute15fec_time_line_instance.pm_fec_param_string")
	proto.RegisterType((*PmFecParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.optics_history.optics_port_histories.optics_port_history.optics_minute15_history.optics_minute15fec_histories.optics_minute15fec_history.optics_minute15fec_time_line_instances.optics_minute15fec_time_line_instance.pm_fec_paras")
}

func init() { proto.RegisterFile("pm_fec_paras.proto", fileDescriptor_f9e952374f03679f) }

var fileDescriptor_f9e952374f03679f = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x5d, 0x6b, 0x14, 0x3d,
	0x14, 0x66, 0xa7, 0xdb, 0xfd, 0x48, 0x5b, 0xde, 0xb7, 0xb1, 0xca, 0x08, 0x0a, 0xa5, 0x82, 0x14,
	0x94, 0xb5, 0xbb, 0xdb, 0xf5, 0xfb, 0x03, 0x2a, 0x8a, 0x20, 0xde, 0x4c, 0x0b, 0xa2, 0x37, 0x21,
	0x9b, 0x3d, 0xdd, 0x0d, 0x4c, 0x92, 0x69, 0x92, 0xa9, 0xf5, 0xbf, 0x08, 0xfe, 0x02, 0xf1, 0xc6,
	0x7f, 0xe0, 0x85, 0x78, 0xe5, 0x95, 0xf8, 0x77, 0x24, 0x99, 0xd9, 0x99, 0xa9, 0x5b, 0x45, 0xf7,
	0x7a, 0xef, 0x72, 0xce, 0x79, 0x9e, 0x27, 0xe7, 0x3c, 0xc9, 0x66, 0x19, 0x84, 0x13, 0x41, 0x0e,
	0x81, 0x91, 0x84, 0x6a, 0x6a, 0x3a, 0x89, 0x56, 0x56, 0xe1, 0xf7, 0x01, 0xe3, 0x86, 0x29, 0xc2,
	0x95, 0x21, 0x27, 0x9a, 0x24, 0x02, 0xe4, 0x98, 0x4b, 0x20, 0x2a, 0x01, 0xdd, 0x49, 0x40, 0x1f,
	0x2a, 0x2d, 0xa8, 0x64, 0x40, 0x04, 0x95, 0x74, 0x0c, 0x02, 0xa4, 0x25, 0x13, 0x6e, 0xac, 0xd2,
	0x6f, 0x3b, 0xe3, 0x58, 0x0d, 0x69, 0xec, 0x50, 0x5c, 0x8d, 0x38, 0xeb, 0xa8, 0xc4, 0x72, 0x66,
	0x8a, 0x72, 0x1e, 0x26, 0x4a, 0x4f, 0x29, 0x1c, 0xcc, 0x19, 0xd9, 0x02, 0x29, 0xb8, 0x4c, 0x2d,
	0x74, 0x07, 0xbf, 0xcb, 0xbb, 0xde, 0x67, 0x84, 0x66, 0x8b, 0x67, 0xf2, 0x2c, 0x17, 0x40, 0x62,
	0x37, 0x19, 0x97, 0xc6, 0xba, 0x89, 0xcc, 0xdf, 0xc1, 0xb6, 0x5e, 0xa3, 0xf5, 0xaa, 0x6f, 0xe4,
	0xf9, 0x93, 0x57, 0xfb, 0x18, 0xa3, 0xba, 0xa4, 0x02, 0xc2, 0xda, 0x66, 0x6d, 0xbb, 0x1d, 0xf9,
	0x35, 0xbe, 0x80, 0x1a, 0x32, 0x15, 0x43, 0xd0, 0x61, 0xb0, 0x59, 0xdb, 0x5e, 0x8b, 0xf2, 0x08,
	0x5f, 0x44, 0xad, 0x6c, 0x45, 0xba, 0xe1, 0x92, 0xaf, 0x34, 0xb3, 0xb8, 0xbb, 0x95, 0xa2, 0xd5,
	0x8a, 0xb6, 0x70, 0xb2, 0x23, 0x6a, 0xa9, 0x97, 0xad, 0x47, 0x7e, 0x8d, 0x2f, 0xa1, 0xb6, 0x9d,
	0x68, 0x30, 0x13, 0x15, 0x8f, 0xbc, 0x72, 0x3d, 0x2a, 0x13, 0xf8, 0x32, 0x42, 0x96, 0x51, 0xa2,
	0xc1, 0x19, 0xea, 0xe5, 0x5b, 0x51, 0xdb, 0x32, 0x1a, 0xf9, 0x04, 0xde, 0x40, 0xcb, 0xc7, 0x34,
	0xe6, 0xa3, 0xb0, 0xee, 0x2b, 0x59, 0xb0, 0xf5, 0x31, 0x40, 0xe7, 0xaa, 0xfb, 0x12, 0x63, 0x35,
	0x97, 0x63, 0x1c, 0xa2, 0xa6, 0xe0, 0x92, 0x8b, 0x54, 0xe4, 0x83, 0x4d, 0x43, 0x57, 0xa1, 0xc7,
	0xa0, 0xe9, 0x18, 0x7c, 0x0b, 0xed, 0x68, 0x1a, 0x7a, 0x0e, 0x3d, 0xf1, 0x9c, 0xa5, 0x9c, 0x93,
	0x85, 0xf8, 0x1a, 0x5a, 0xcf, 0xe9, 0xa4, 0x1c, 0xa0, 0xee, 0x31, 0xff, 0xe7, 0x85, 0x83, 0x62,
	0x8e, 0xeb, 0x08, 0x17, 0xe0, 0x72, 0x9e, 0x65, 0xdf, 0x75, 0x81, 0x2e, 0xc6, 0x72, 0xd2, 0xd9,
	0x2e, 0x15, 0xe9, 0x46, 0x2e, 0x9d, 0x15, 0x4e, 0x4b, 0x4f, 0xc1, 0xa5, 0x74, 0x33, 0x97, 0xce,
	0xd1, 0xb3, 0x8e, 0xb5, 0xaa, 0x8e, 0xbd, 0xc3, 0xa7, 0x4e, 0xca, 0x38, 0x18, 0x97, 0x23, 0x38,
	0x09, 0x7b, 0xfe, 0x44, 0xb3, 0xa0, 0x24, 0xf7, 0x2b, 0x64, 0x7f, 0x82, 0x5c, 0x80, 0xb1, 0x54,
	0x24, 0xe1, 0xae, 0xef, 0xb2, 0x4c, 0xe0, 0xab, 0xe8, 0xbf, 0x98, 0x1a, 0x4b, 0x58, 0x0c, 0x54,
	0xfb, 0x0b, 0x18, 0x0e, 0x3c, 0x66, 0xcd, 0xa5, 0x1f, 0xbb, 0xec, 0x01, 0x17, 0x80, 0xbb, 0xe8,
	0x7c, 0x89, 0xeb, 0x0e, 0xdc, 0xb5, 0xcd, 0xd0, 0x37, 0x3d, 0x1a, 0x17, 0xe8, 0xee, 0xe0, 0x05,
	0x97, 0xb3, 0x94, 0xfe, 0x0e, 0x31, 0xf9, 0x0d, 0x0f, 0x6f, 0xfd, 0x42, 0xe9, 0xef, 0xec, 0x03,
	0xf3, 0x94, 0x1b, 0x68, 0xa3, 0xa4, 0xf4, 0x76, 0xc9, 0x24, 0x6f, 0xe9, 0xb6, 0x67, 0xac, 0x17,
	0x8c, 0xde, 0xee, 0xb3, 0xac, 0xad, 0x2b, 0x68, 0xcd, 0x00, 0x73, 0xe2, 0x69, 0xe2, 0x8d, 0xbd,
	0xe3, 0x47, 0x5f, 0xf5, 0xc9, 0xfd, 0x2c, 0x87, 0xbf, 0x04, 0xa8, 0x09, 0x8c, 0x0c, 0xb9, 0x35,
	0xe1, 0xdd, 0xcd, 0xda, 0xf6, 0x4a, 0xef, 0x43, 0xd0, 0x59, 0x3c, 0x3c, 0x7f, 0x84, 0x75, 0xaa,
	0xbf, 0xd0, 0xa8, 0x01, 0x6c, 0x8f, 0x5b, 0x83, 0xbf, 0x06, 0xa8, 0x95, 0x32, 0xf2, 0x46, 0xe9,
	0x91, 0x09, 0xef, 0x2d, 0xac, 0x9c, 0xc3, 0xca, 0x66, 0xca, 0x5e, 0x3a, 0xfb, 0xf0, 0xf7, 0x00,
	0xad, 0x24, 0x1a, 0x7c, 0xc9, 0x3d, 0xdb, 0xf7, 0xbd, 0x9d, 0x9f, 0x16, 0x76, 0xfe, 0x8b, 0x9d,
	0xf9, 0x7f, 0x47, 0xd4, 0x4e, 0x34, 0x3c, 0x05, 0xb6, 0x07, 0x1a, 0xff, 0x08, 0xd0, 0x6a, 0xa2,
	0x8c, 0x2d, 0x8c, 0x7d, 0xb0, 0x30, 0x76, 0x7e, 0x63, 0x91, 0xb3, 0x32, 0x77, 0xf6, 0x73, 0x80,
	0x6a, 0x47, 0xe1, 0xc3, 0x85, 0x9d, 0xf3, 0xdb, 0x59, 0x3b, 0xc2, 0xdf, 0x02, 0xd4, 0x3c, 0x12,
	0x54, 0x8f, 0xb9, 0x0c, 0x1f, 0x2d, 0xbc, 0x9c, 0xdf, 0xcb, 0xa9, 0x8b, 0xc3, 0x86, 0xff, 0x98,
	0xe8, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x30, 0xeb, 0x08, 0x62, 0x0c, 0x00, 0x00,
}
