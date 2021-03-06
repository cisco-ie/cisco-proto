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

package cisco_ios_xr_pmengine_oper_performance_management_history_global_periodic_ethernet_history_ethernet_port_histories_ethernet_port_history_ethernet_minute15_history_minute15_fec_histories_minute15_fec_history_minute15_fec_time_line_instances_minute15_fec_time_line_instance

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
	SlotNumber           uint32   `protobuf:"varint,2,opt,name=slot_number,json=slotNumber,proto3" json:"slot_number,omitempty"`
	Number               uint32   `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
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

func (m *PmFecParas_KEYS) GetSlotNumber() uint32 {
	if m != nil {
		return m.SlotNumber
	}
	return 0
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
	proto.RegisterType((*PmFecParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ethernet_history.ethernet_port_histories.ethernet_port_history.ethernet_minute15_history.minute15_fec_histories.minute15_fec_history.minute15_fec_time_line_instances.minute15_fec_time_line_instance.pm_fec_paras_KEYS")
	proto.RegisterType((*PmFecParam)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ethernet_history.ethernet_port_histories.ethernet_port_history.ethernet_minute15_history.minute15_fec_histories.minute15_fec_history.minute15_fec_time_line_instances.minute15_fec_time_line_instance.pm_fec_param")
	proto.RegisterType((*PmFecParamString)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ethernet_history.ethernet_port_histories.ethernet_port_history.ethernet_minute15_history.minute15_fec_histories.minute15_fec_history.minute15_fec_time_line_instances.minute15_fec_time_line_instance.pm_fec_param_string")
	proto.RegisterType((*PmFecParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ethernet_history.ethernet_port_histories.ethernet_port_history.ethernet_minute15_history.minute15_fec_histories.minute15_fec_history.minute15_fec_time_line_instances.minute15_fec_time_line_instance.pm_fec_paras")
}

func init() { proto.RegisterFile("pm_fec_paras.proto", fileDescriptor_f9e952374f03679f) }

var fileDescriptor_f9e952374f03679f = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x55, 0xdc, 0x34, 0x69, 0xa6, 0xad, 0xde, 0xcb, 0xbc, 0xbe, 0x27, 0x2f, 0x1e, 0xa2, 0x2a,
	0x12, 0xaa, 0x04, 0x0a, 0x4d, 0xd2, 0xf0, 0xfd, 0x21, 0x15, 0x81, 0x90, 0x10, 0x2c, 0xdc, 0x4a,
	0x88, 0xd5, 0x30, 0x99, 0xdc, 0x26, 0x23, 0x79, 0xc6, 0xee, 0xcc, 0xa4, 0x94, 0x5f, 0xc1, 0x2f,
	0x60, 0xc7, 0x9e, 0x05, 0x0b, 0x84, 0x84, 0xc4, 0x0a, 0xf1, 0xb7, 0xd0, 0x5c, 0x3b, 0xb6, 0xfb,
	0x21, 0xc8, 0x0f, 0xc8, 0xce, 0xf7, 0xdc, 0x73, 0x8e, 0xef, 0x3d, 0x63, 0x3b, 0x21, 0x34, 0x55,
	0xec, 0x10, 0x04, 0x4b, 0xb9, 0xe1, 0xb6, 0x93, 0x9a, 0xc4, 0x25, 0xf4, 0x7d, 0x20, 0xa4, 0x15,
	0x09, 0x93, 0x89, 0x65, 0x27, 0x86, 0xa5, 0x0a, 0xf4, 0x58, 0x6a, 0x60, 0x49, 0x0a, 0xa6, 0x93,
	0x82, 0x39, 0x4c, 0x8c, 0xe2, 0x5a, 0x00, 0x53, 0x5c, 0xf3, 0x31, 0x28, 0xd0, 0x8e, 0x4d, 0xa4,
	0x75, 0x89, 0x79, 0xd7, 0x19, 0xc7, 0xc9, 0x90, 0xc7, 0x9e, 0x25, 0x93, 0x91, 0x14, 0x1d, 0x70,
	0x13, 0x30, 0x1a, 0x4a, 0x42, 0x01, 0xa4, 0x89, 0x99, 0xa1, 0x12, 0xec, 0x85, 0x78, 0x85, 0xad,
	0xa4, 0x9e, 0x3a, 0xe8, 0x0e, 0x8a, 0x4e, 0x01, 0xf8, 0xe9, 0x4b, 0x9b, 0x0b, 0xe0, 0x33, 0x5c,
	0x27, 0x15, 0xb0, 0xd8, 0xef, 0x23, 0xb5, 0x75, 0x7e, 0x0f, 0xfb, 0x27, 0xc2, 0xd6, 0x1b, 0xd2,
	0xae, 0xe6, 0xc4, 0x9e, 0x3f, 0x79, 0xbd, 0x4f, 0x29, 0xa9, 0x6b, 0xae, 0x20, 0xac, 0x6d, 0xd6,
	0xb6, 0x5b, 0x11, 0x5e, 0xd3, 0xcb, 0x64, 0xd5, 0xc6, 0x89, 0x63, 0x7a, 0xaa, 0x86, 0x60, 0xc2,
	0x60, 0xb3, 0xb6, 0xbd, 0x1e, 0x11, 0x0f, 0xbd, 0x44, 0x84, 0xfe, 0x47, 0x1a, 0x79, 0x6f, 0x09,
	0x7b, 0x79, 0xb5, 0x35, 0x25, 0x6b, 0x95, 0x3b, 0x28, 0x6f, 0x3e, 0xe2, 0x8e, 0xa3, 0x79, 0x3d,
	0xc2, 0x6b, 0xfa, 0x3f, 0x69, 0xb9, 0x89, 0x01, 0x3b, 0x49, 0xe2, 0x11, 0x5a, 0xd7, 0xa3, 0x12,
	0xa0, 0x97, 0x08, 0x71, 0x82, 0x33, 0x03, 0x3e, 0x40, 0x74, 0x5f, 0x89, 0x5a, 0x4e, 0xf0, 0x08,
	0x01, 0xba, 0x41, 0x96, 0x8f, 0x79, 0x2c, 0x47, 0x61, 0x1d, 0x3b, 0x59, 0xb1, 0xf5, 0x29, 0x20,
	0xff, 0x54, 0xef, 0xcb, 0xac, 0x33, 0x52, 0x8f, 0x69, 0x48, 0x9a, 0x4a, 0x6a, 0xa9, 0xa6, 0x2a,
	0x5f, 0x6f, 0x56, 0xfa, 0x0e, 0x3f, 0x06, 0xc3, 0xc7, 0x80, 0x23, 0xb4, 0xa2, 0x59, 0x89, 0x1a,
	0x7e, 0x82, 0x9a, 0xa5, 0x5c, 0x93, 0x95, 0xf4, 0x1a, 0x69, 0xe7, 0x72, 0x56, 0x2e, 0x50, 0x47,
	0xce, 0xdf, 0x79, 0xe3, 0xa0, 0xd8, 0xe3, 0x3a, 0xa1, 0x05, 0xb9, 0xdc, 0x67, 0x19, 0xa7, 0x2e,
	0xd8, 0xc5, 0x5a, 0xde, 0x3a, 0xbb, 0x4b, 0xc5, 0xba, 0x91, 0x5b, 0x67, 0x8d, 0xd3, 0xd6, 0x33,
	0x72, 0x69, 0xdd, 0xcc, 0xad, 0x73, 0xf6, 0xf9, 0xc4, 0x56, 0xaa, 0x89, 0x7d, 0x6b, 0x9f, 0x3a,
	0x29, 0xeb, 0x69, 0x52, 0x8f, 0xe0, 0x24, 0xec, 0xe1, 0x81, 0x66, 0x45, 0x29, 0xee, 0x57, 0xc4,
	0x78, 0x82, 0x52, 0x81, 0x75, 0x5c, 0xa5, 0xe1, 0x2e, 0x4e, 0x59, 0x02, 0xf4, 0x2a, 0xf9, 0x2b,
	0xe6, 0xd6, 0x31, 0x11, 0x03, 0x37, 0xf8, 0x18, 0x86, 0x03, 0xe4, 0xac, 0x7b, 0xf8, 0xb1, 0x47,
	0x0f, 0xa4, 0x02, 0xda, 0x25, 0xff, 0x96, 0xbc, 0xee, 0xc0, 0xbf, 0x1b, 0x19, 0xfb, 0x26, 0xb2,
	0x69, 0xc1, 0xee, 0x0e, 0x5e, 0x48, 0x7d, 0x5e, 0xd2, 0xdf, 0x61, 0x36, 0x7f, 0xce, 0xc3, 0x5b,
	0x67, 0x24, 0xfd, 0x9d, 0x7d, 0x10, 0x28, 0xb9, 0x41, 0x36, 0x4a, 0x49, 0x6f, 0x97, 0x4d, 0xf2,
	0x91, 0x6e, 0xa3, 0xa2, 0x5d, 0x28, 0x7a, 0xbb, 0xcf, 0xb2, 0xb1, 0xae, 0x90, 0x75, 0x0b, 0xc2,
	0x9b, 0x4f, 0x53, 0x0c, 0xf6, 0x0e, 0xae, 0xbe, 0x86, 0xe0, 0x7e, 0x86, 0xd1, 0x2f, 0x01, 0x69,
	0x82, 0x60, 0x43, 0xe9, 0x6c, 0x78, 0x77, 0xb3, 0xb6, 0xbd, 0xda, 0xfb, 0x10, 0x74, 0x16, 0x9f,
	0x9b, 0x53, 0x84, 0x4e, 0xf5, 0x8d, 0x8c, 0x1a, 0x20, 0xf6, 0xa4, 0xb3, 0xf4, 0x6b, 0x40, 0x56,
	0xa6, 0x82, 0xbd, 0x4d, 0xcc, 0xc8, 0x86, 0xf7, 0x16, 0xd1, 0xcd, 0x11, 0x5d, 0x73, 0x2a, 0x5e,
	0xf9, 0xb8, 0xe8, 0x8f, 0x80, 0xac, 0xa6, 0x06, 0xb0, 0xe5, 0xbf, 0xbd, 0xf7, 0x31, 0xbe, 0x8f,
	0x8b, 0xf8, 0x7e, 0x17, 0x5f, 0xfe, 0x5b, 0x10, 0xb5, 0x52, 0x03, 0x4f, 0x41, 0xec, 0x81, 0xa1,
	0x3f, 0x03, 0xb2, 0x96, 0x26, 0xd6, 0x15, 0x41, 0x3e, 0x58, 0x04, 0x39, 0x7f, 0x90, 0xc4, 0x47,
	0x97, 0x27, 0xf9, 0x39, 0x20, 0xb5, 0xa3, 0xf0, 0xe1, 0x22, 0xbe, 0xf9, 0xe3, 0xab, 0x1d, 0xd1,
	0xef, 0x01, 0x69, 0x1e, 0x29, 0x6e, 0xc6, 0x52, 0x87, 0x8f, 0x16, 0xd9, 0xcd, 0x9f, 0xdd, 0x2c,
	0xb5, 0x61, 0x03, 0xff, 0xe2, 0xf7, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x96, 0x49, 0xa3, 0x23,
	0xf8, 0x0b, 0x00, 0x00,
}
