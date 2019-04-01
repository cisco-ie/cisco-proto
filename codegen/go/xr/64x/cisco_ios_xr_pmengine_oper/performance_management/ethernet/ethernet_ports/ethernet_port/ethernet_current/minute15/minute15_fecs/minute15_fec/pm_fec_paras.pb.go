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

package cisco_ios_xr_pmengine_oper_performance_management_ethernet_ethernet_ports_ethernet_port_ethernet_current_minute15_minute15_fecs_minute15_fec

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
	proto.RegisterType((*PmFecParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ethernet.ethernet_ports.ethernet_port.ethernet_current.minute15.minute15_fecs.minute15_fec.pm_fec_paras_KEYS")
	proto.RegisterType((*PmFecParam)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ethernet.ethernet_ports.ethernet_port.ethernet_current.minute15.minute15_fecs.minute15_fec.pm_fec_param")
	proto.RegisterType((*PmFecParamString)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ethernet.ethernet_ports.ethernet_port.ethernet_current.minute15.minute15_fecs.minute15_fec.pm_fec_param_string")
	proto.RegisterType((*PmFecParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ethernet.ethernet_ports.ethernet_port.ethernet_current.minute15.minute15_fecs.minute15_fec.pm_fec_paras")
}

func init() { proto.RegisterFile("pm_fec_paras.proto", fileDescriptor_f9e952374f03679f) }

var fileDescriptor_f9e952374f03679f = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcd, 0x6e, 0x13, 0x31,
	0x18, 0x94, 0xdb, 0xb4, 0x69, 0x9c, 0x56, 0x50, 0x53, 0xa4, 0x3d, 0x80, 0x88, 0x8a, 0x84, 0x22,
	0x81, 0x42, 0x93, 0x34, 0xfc, 0xff, 0x48, 0x45, 0xa0, 0x4a, 0x08, 0x0e, 0xdb, 0x4a, 0x88, 0x93,
	0xe5, 0x3a, 0x5f, 0x13, 0x4b, 0xb1, 0x77, 0x6b, 0x7b, 0x4b, 0x1f, 0x82, 0x27, 0xe0, 0xc6, 0x95,
	0x0b, 0x1c, 0x10, 0x27, 0x4e, 0x88, 0x07, 0x43, 0xfe, 0x76, 0xb3, 0xbb, 0xa1, 0x2f, 0xd0, 0xde,
	0xfc, 0x8d, 0x67, 0x66, 0x67, 0x6c, 0xed, 0x6a, 0x29, 0x4b, 0x35, 0x3f, 0x06, 0xc9, 0x53, 0x61,
	0x85, 0xeb, 0xa5, 0x36, 0xf1, 0x09, 0xfb, 0x4c, 0xa4, 0x72, 0x32, 0xe1, 0x2a, 0x71, 0xfc, 0xcc,
	0xf2, 0x54, 0x83, 0x99, 0x28, 0x03, 0x3c, 0x49, 0xc1, 0xf6, 0x52, 0xb0, 0xc7, 0x89, 0xd5, 0xc2,
	0x48, 0xe0, 0x5a, 0x18, 0x31, 0x01, 0x0d, 0xc6, 0xf7, 0xc0, 0x4f, 0xc1, 0x1a, 0xa8, 0x16, 0x3c,
	0x4d, 0xac, 0x77, 0x8b, 0x63, 0x35, 0xc9, 0xcc, 0xda, 0x20, 0xd3, 0xca, 0x64, 0x1e, 0xfa, 0xa3,
	0x72, 0x11, 0xc2, 0xb8, 0x85, 0x69, 0x7b, 0x9f, 0x6e, 0xd6, 0x43, 0xf2, 0xb7, 0xaf, 0x3f, 0x1e,
	0x30, 0x46, 0x1b, 0x46, 0x68, 0x88, 0x48, 0x87, 0x74, 0x5b, 0x31, 0xae, 0xd9, 0x2d, 0xda, 0x76,
	0xb3, 0xc4, 0x73, 0x93, 0xe9, 0x23, 0xb0, 0xd1, 0x52, 0x87, 0x74, 0x37, 0x62, 0x1a, 0xa0, 0xf7,
	0x88, 0x6c, 0x67, 0x74, 0xbd, 0xe6, 0xa4, 0x83, 0xc9, 0x58, 0x78, 0x81, 0x26, 0x8d, 0x18, 0xd7,
	0xec, 0x06, 0x6d, 0xf9, 0xa9, 0x05, 0x37, 0x4d, 0x66, 0x63, 0xb4, 0x68, 0xc4, 0x15, 0xc0, 0x6e,
	0x52, 0xea, 0xa5, 0xe0, 0x16, 0x42, 0x9d, 0x68, 0xb9, 0x43, 0xba, 0x6b, 0x71, 0xcb, 0x4b, 0x11,
	0x23, 0xc0, 0xb6, 0xe8, 0xca, 0xa9, 0x98, 0xa9, 0x71, 0xd4, 0xc0, 0x9d, 0x7c, 0xd8, 0xfe, 0xbe,
	0x44, 0xaf, 0xd5, 0x9f, 0xcb, 0x9d, 0xb7, 0xca, 0x4c, 0x58, 0x44, 0x9b, 0x5a, 0x19, 0xa5, 0x33,
	0x5d, 0xd4, 0x98, 0x8f, 0x61, 0x47, 0x9c, 0x82, 0x15, 0x13, 0xc0, 0x08, 0xad, 0x78, 0x3e, 0xa2,
	0x46, 0x9c, 0xa1, 0x66, 0xb9, 0xd0, 0xe4, 0x23, 0xbb, 0x4b, 0x37, 0x0b, 0x39, 0xaf, 0x0a, 0x34,
	0x90, 0x73, 0xb5, 0xd8, 0x38, 0x2c, 0x7b, 0xdc, 0xa3, 0xac, 0x24, 0x57, 0x7d, 0x56, 0x30, 0x75,
	0xc9, 0x2e, 0x6b, 0x05, 0xeb, 0xfc, 0x29, 0x35, 0xeb, 0xd5, 0xc2, 0x3a, 0xdf, 0x58, 0xb4, 0x9e,
	0x93, 0x2b, 0xeb, 0x66, 0x61, 0x5d, 0xb0, 0xcf, 0x9f, 0xd8, 0x5a, 0xfd, 0xc4, 0xbe, 0xb5, 0x17,
	0x6e, 0xca, 0x05, 0x9a, 0x32, 0x63, 0x38, 0x8b, 0x06, 0x78, 0xa9, 0xf9, 0x50, 0x89, 0x87, 0x35,
	0x31, 0xde, 0xa0, 0xd2, 0xe0, 0xbc, 0xd0, 0x69, 0xb4, 0x8b, 0x29, 0x2b, 0x80, 0xdd, 0xa1, 0x57,
	0x66, 0xc2, 0x79, 0x2e, 0x67, 0x20, 0x2c, 0x0f, 0x78, 0x34, 0x42, 0xce, 0x46, 0x80, 0x5f, 0x05,
	0xf4, 0x50, 0x69, 0x60, 0x7d, 0x7a, 0xbd, 0xe2, 0xf5, 0x47, 0x5c, 0x2b, 0x93, 0xb3, 0x1f, 0x20,
	0x9b, 0x95, 0xec, 0xfe, 0xe8, 0x9d, 0x32, 0xe7, 0x25, 0xc3, 0x1d, 0xee, 0x40, 0xe6, 0x92, 0x87,
	0xff, 0x49, 0x86, 0x3b, 0x07, 0x20, 0x51, 0x72, 0x9f, 0x6e, 0x55, 0x92, 0xc1, 0x2e, 0x9f, 0x16,
	0x91, 0x1e, 0xa1, 0x62, 0xb3, 0x54, 0x0c, 0x76, 0xf7, 0xf3, 0x58, 0xb7, 0xe9, 0x86, 0x03, 0x19,
	0xcc, 0xb3, 0x14, 0x0f, 0xf6, 0x31, 0x56, 0x5f, 0x47, 0xf0, 0x20, 0xc7, 0xd8, 0x4f, 0x42, 0x9b,
	0x20, 0xf9, 0x91, 0xf2, 0x2e, 0x7a, 0xd2, 0x21, 0xdd, 0xf6, 0xe0, 0x0b, 0xe9, 0x5d, 0xa4, 0x77,
	0xba, 0x57, 0x7f, 0x1d, 0xe2, 0x55, 0x90, 0x7b, 0xca, 0x3b, 0xf6, 0x8b, 0xd0, 0xb5, 0x4c, 0xf2,
	0x4f, 0x89, 0x1d, 0xbb, 0xe8, 0xe9, 0xc5, 0xcf, 0xdd, 0xcc, 0xe4, 0x87, 0x90, 0x95, 0xfd, 0x21,
	0xb4, 0x9d, 0x5a, 0xc0, 0xad, 0xf0, 0xe5, 0x79, 0x86, 0xd9, 0xbf, 0x5e, 0xe0, 0xec, 0xc5, 0x27,
	0x28, 0x6e, 0xa5, 0x16, 0xde, 0x80, 0xdc, 0x03, 0xcb, 0xfe, 0x12, 0xba, 0x9e, 0x26, 0xce, 0x97,
	0x2d, 0x9e, 0x5f, 0x9a, 0x16, 0x34, 0xe4, 0x2e, 0x6a, 0xfc, 0x20, 0x94, 0x9c, 0x44, 0x2f, 0x2e,
	0x4d, 0x76, 0x72, 0xc2, 0x7e, 0x13, 0xda, 0x3c, 0xd1, 0xc2, 0x4e, 0x94, 0x89, 0x5e, 0x5e, 0x9a,
	0xe0, 0xf3, 0xc8, 0x47, 0xab, 0xf8, 0xd7, 0x30, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xd4,
	0xc5, 0x08, 0x4b, 0x08, 0x00, 0x00,
}
