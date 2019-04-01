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
// source: pm_optics_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_optics_optics_ports_optics_port_optics_current_optics_second30_optics_second30_optics_optics_second30_optic

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

type PmOpticsParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmOpticsParas_KEYS) Reset()         { *m = PmOpticsParas_KEYS{} }
func (m *PmOpticsParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmOpticsParas_KEYS) ProtoMessage()    {}
func (*PmOpticsParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_380a2e2d176d5301, []int{0}
}

func (m *PmOpticsParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmOpticsParas_KEYS.Unmarshal(m, b)
}
func (m *PmOpticsParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmOpticsParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmOpticsParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmOpticsParas_KEYS.Merge(m, src)
}
func (m *PmOpticsParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmOpticsParas_KEYS.Size(m)
}
func (m *PmOpticsParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmOpticsParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmOpticsParas_KEYS proto.InternalMessageInfo

func (m *PmOpticsParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmOpticsParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PmOpticsParameter struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Minimum              int32    `protobuf:"zigzag32,2,opt,name=minimum,proto3" json:"minimum,omitempty"`
	Average              int32    `protobuf:"zigzag32,3,opt,name=average,proto3" json:"average,omitempty"`
	Maximum              int32    `protobuf:"zigzag32,4,opt,name=maximum,proto3" json:"maximum,omitempty"`
	MinimumThreshold     int32    `protobuf:"zigzag32,5,opt,name=minimum_threshold,json=minimumThreshold,proto3" json:"minimum_threshold,omitempty"`
	MinimumTcaReport     bool     `protobuf:"varint,6,opt,name=minimum_tca_report,json=minimumTcaReport,proto3" json:"minimum_tca_report,omitempty"`
	MaximumThreshold     int32    `protobuf:"zigzag32,7,opt,name=maximum_threshold,json=maximumThreshold,proto3" json:"maximum_threshold,omitempty"`
	MaximumTcaReport     bool     `protobuf:"varint,8,opt,name=maximum_tca_report,json=maximumTcaReport,proto3" json:"maximum_tca_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmOpticsParameter) Reset()         { *m = PmOpticsParameter{} }
func (m *PmOpticsParameter) String() string { return proto.CompactTextString(m) }
func (*PmOpticsParameter) ProtoMessage()    {}
func (*PmOpticsParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_380a2e2d176d5301, []int{1}
}

func (m *PmOpticsParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmOpticsParameter.Unmarshal(m, b)
}
func (m *PmOpticsParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmOpticsParameter.Marshal(b, m, deterministic)
}
func (m *PmOpticsParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmOpticsParameter.Merge(m, src)
}
func (m *PmOpticsParameter) XXX_Size() int {
	return xxx_messageInfo_PmOpticsParameter.Size(m)
}
func (m *PmOpticsParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PmOpticsParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PmOpticsParameter proto.InternalMessageInfo

func (m *PmOpticsParameter) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmOpticsParameter) GetMinimum() int32 {
	if m != nil {
		return m.Minimum
	}
	return 0
}

func (m *PmOpticsParameter) GetAverage() int32 {
	if m != nil {
		return m.Average
	}
	return 0
}

func (m *PmOpticsParameter) GetMaximum() int32 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func (m *PmOpticsParameter) GetMinimumThreshold() int32 {
	if m != nil {
		return m.MinimumThreshold
	}
	return 0
}

func (m *PmOpticsParameter) GetMinimumTcaReport() bool {
	if m != nil {
		return m.MinimumTcaReport
	}
	return false
}

func (m *PmOpticsParameter) GetMaximumThreshold() int32 {
	if m != nil {
		return m.MaximumThreshold
	}
	return 0
}

func (m *PmOpticsParameter) GetMaximumTcaReport() bool {
	if m != nil {
		return m.MaximumTcaReport
	}
	return false
}

type PmOpticsParameterString struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Minimum              string   `protobuf:"bytes,2,opt,name=minimum,proto3" json:"minimum,omitempty"`
	Average              string   `protobuf:"bytes,3,opt,name=average,proto3" json:"average,omitempty"`
	Maximum              string   `protobuf:"bytes,4,opt,name=maximum,proto3" json:"maximum,omitempty"`
	MinimumThreshold     string   `protobuf:"bytes,5,opt,name=minimum_threshold,json=minimumThreshold,proto3" json:"minimum_threshold,omitempty"`
	MinimumTcaReport     bool     `protobuf:"varint,6,opt,name=minimum_tca_report,json=minimumTcaReport,proto3" json:"minimum_tca_report,omitempty"`
	MaximumThreshold     string   `protobuf:"bytes,7,opt,name=maximum_threshold,json=maximumThreshold,proto3" json:"maximum_threshold,omitempty"`
	MaximumTcaReport     bool     `protobuf:"varint,8,opt,name=maximum_tca_report,json=maximumTcaReport,proto3" json:"maximum_tca_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmOpticsParameterString) Reset()         { *m = PmOpticsParameterString{} }
func (m *PmOpticsParameterString) String() string { return proto.CompactTextString(m) }
func (*PmOpticsParameterString) ProtoMessage()    {}
func (*PmOpticsParameterString) Descriptor() ([]byte, []int) {
	return fileDescriptor_380a2e2d176d5301, []int{2}
}

func (m *PmOpticsParameterString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmOpticsParameterString.Unmarshal(m, b)
}
func (m *PmOpticsParameterString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmOpticsParameterString.Marshal(b, m, deterministic)
}
func (m *PmOpticsParameterString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmOpticsParameterString.Merge(m, src)
}
func (m *PmOpticsParameterString) XXX_Size() int {
	return xxx_messageInfo_PmOpticsParameterString.Size(m)
}
func (m *PmOpticsParameterString) XXX_DiscardUnknown() {
	xxx_messageInfo_PmOpticsParameterString.DiscardUnknown(m)
}

var xxx_messageInfo_PmOpticsParameterString proto.InternalMessageInfo

func (m *PmOpticsParameterString) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmOpticsParameterString) GetMinimum() string {
	if m != nil {
		return m.Minimum
	}
	return ""
}

func (m *PmOpticsParameterString) GetAverage() string {
	if m != nil {
		return m.Average
	}
	return ""
}

func (m *PmOpticsParameterString) GetMaximum() string {
	if m != nil {
		return m.Maximum
	}
	return ""
}

func (m *PmOpticsParameterString) GetMinimumThreshold() string {
	if m != nil {
		return m.MinimumThreshold
	}
	return ""
}

func (m *PmOpticsParameterString) GetMinimumTcaReport() bool {
	if m != nil {
		return m.MinimumTcaReport
	}
	return false
}

func (m *PmOpticsParameterString) GetMaximumThreshold() string {
	if m != nil {
		return m.MaximumThreshold
	}
	return ""
}

func (m *PmOpticsParameterString) GetMaximumTcaReport() bool {
	if m != nil {
		return m.MaximumTcaReport
	}
	return false
}

type PmOpticsParas struct {
	Index                uint32                   `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool                     `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            string                   `protobuf:"bytes,52,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        string                   `protobuf:"bytes,53,opt,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   string                   `protobuf:"bytes,54,opt,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear30SecTime   string                   `protobuf:"bytes,55,opt,name=last_clear30_sec_time,json=lastClear30SecTime,proto3" json:"last_clear30_sec_time,omitempty"`
	LastClear24HrTime    string                   `protobuf:"bytes,56,opt,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Sec30Support         bool                     `protobuf:"varint,57,opt,name=sec30_support,json=sec30Support,proto3" json:"sec30_support,omitempty"`
	Lbc                  *PmOpticsParameter       `protobuf:"bytes,58,opt,name=lbc,proto3" json:"lbc,omitempty"`
	LbcPc                *PmOpticsParameterString `protobuf:"bytes,59,opt,name=lbc_pc,json=lbcPc,proto3" json:"lbc_pc,omitempty"`
	Opt                  *PmOpticsParameterString `protobuf:"bytes,60,opt,name=opt,proto3" json:"opt,omitempty"`
	Opr                  *PmOpticsParameterString `protobuf:"bytes,61,opt,name=opr,proto3" json:"opr,omitempty"`
	Cd                   *PmOpticsParameter       `protobuf:"bytes,62,opt,name=cd,proto3" json:"cd,omitempty"`
	Dgd                  *PmOpticsParameterString `protobuf:"bytes,63,opt,name=dgd,proto3" json:"dgd,omitempty"`
	Pmd                  *PmOpticsParameterString `protobuf:"bytes,64,opt,name=pmd,proto3" json:"pmd,omitempty"`
	Osnr                 *PmOpticsParameterString `protobuf:"bytes,65,opt,name=osnr,proto3" json:"osnr,omitempty"`
	CenterWavelength     *PmOpticsParameterString `protobuf:"bytes,66,opt,name=center_wavelength,json=centerWavelength,proto3" json:"center_wavelength,omitempty"`
	Pdl                  *PmOpticsParameterString `protobuf:"bytes,67,opt,name=pdl,proto3" json:"pdl,omitempty"`
	Pcr                  *PmOpticsParameterString `protobuf:"bytes,68,opt,name=pcr,proto3" json:"pcr,omitempty"`
	Pn                   *PmOpticsParameterString `protobuf:"bytes,69,opt,name=pn,proto3" json:"pn,omitempty"`
	RxSigPow             *PmOpticsParameterString `protobuf:"bytes,70,opt,name=rx_sig_pow,json=rxSigPow,proto3" json:"rx_sig_pow,omitempty"`
	LowSigFreqOff        *PmOpticsParameter       `protobuf:"bytes,71,opt,name=low_sig_freq_off,json=lowSigFreqOff,proto3" json:"low_sig_freq_off,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PmOpticsParas) Reset()         { *m = PmOpticsParas{} }
func (m *PmOpticsParas) String() string { return proto.CompactTextString(m) }
func (*PmOpticsParas) ProtoMessage()    {}
func (*PmOpticsParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_380a2e2d176d5301, []int{3}
}

func (m *PmOpticsParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmOpticsParas.Unmarshal(m, b)
}
func (m *PmOpticsParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmOpticsParas.Marshal(b, m, deterministic)
}
func (m *PmOpticsParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmOpticsParas.Merge(m, src)
}
func (m *PmOpticsParas) XXX_Size() int {
	return xxx_messageInfo_PmOpticsParas.Size(m)
}
func (m *PmOpticsParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmOpticsParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmOpticsParas proto.InternalMessageInfo

func (m *PmOpticsParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmOpticsParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmOpticsParas) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *PmOpticsParas) GetLastClearTime() string {
	if m != nil {
		return m.LastClearTime
	}
	return ""
}

func (m *PmOpticsParas) GetLastClear15MinTime() string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return ""
}

func (m *PmOpticsParas) GetLastClear30SecTime() string {
	if m != nil {
		return m.LastClear30SecTime
	}
	return ""
}

func (m *PmOpticsParas) GetLastClear24HrTime() string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return ""
}

func (m *PmOpticsParas) GetSec30Support() bool {
	if m != nil {
		return m.Sec30Support
	}
	return false
}

func (m *PmOpticsParas) GetLbc() *PmOpticsParameter {
	if m != nil {
		return m.Lbc
	}
	return nil
}

func (m *PmOpticsParas) GetLbcPc() *PmOpticsParameterString {
	if m != nil {
		return m.LbcPc
	}
	return nil
}

func (m *PmOpticsParas) GetOpt() *PmOpticsParameterString {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (m *PmOpticsParas) GetOpr() *PmOpticsParameterString {
	if m != nil {
		return m.Opr
	}
	return nil
}

func (m *PmOpticsParas) GetCd() *PmOpticsParameter {
	if m != nil {
		return m.Cd
	}
	return nil
}

func (m *PmOpticsParas) GetDgd() *PmOpticsParameterString {
	if m != nil {
		return m.Dgd
	}
	return nil
}

func (m *PmOpticsParas) GetPmd() *PmOpticsParameterString {
	if m != nil {
		return m.Pmd
	}
	return nil
}

func (m *PmOpticsParas) GetOsnr() *PmOpticsParameterString {
	if m != nil {
		return m.Osnr
	}
	return nil
}

func (m *PmOpticsParas) GetCenterWavelength() *PmOpticsParameterString {
	if m != nil {
		return m.CenterWavelength
	}
	return nil
}

func (m *PmOpticsParas) GetPdl() *PmOpticsParameterString {
	if m != nil {
		return m.Pdl
	}
	return nil
}

func (m *PmOpticsParas) GetPcr() *PmOpticsParameterString {
	if m != nil {
		return m.Pcr
	}
	return nil
}

func (m *PmOpticsParas) GetPn() *PmOpticsParameterString {
	if m != nil {
		return m.Pn
	}
	return nil
}

func (m *PmOpticsParas) GetRxSigPow() *PmOpticsParameterString {
	if m != nil {
		return m.RxSigPow
	}
	return nil
}

func (m *PmOpticsParas) GetLowSigFreqOff() *PmOpticsParameter {
	if m != nil {
		return m.LowSigFreqOff
	}
	return nil
}

func init() {
	proto.RegisterType((*PmOpticsParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management.optics.optics_ports.optics_port.optics_current.optics_second30.optics_second30_optics.optics_second30_optic.pm_optics_paras_KEYS")
	proto.RegisterType((*PmOpticsParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management.optics.optics_ports.optics_port.optics_current.optics_second30.optics_second30_optics.optics_second30_optic.pm_optics_parameter")
	proto.RegisterType((*PmOpticsParameterString)(nil), "cisco_ios_xr_pmengine_oper.performance_management.optics.optics_ports.optics_port.optics_current.optics_second30.optics_second30_optics.optics_second30_optic.pm_optics_parameter_string")
	proto.RegisterType((*PmOpticsParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management.optics.optics_ports.optics_port.optics_current.optics_second30.optics_second30_optics.optics_second30_optic.pm_optics_paras")
}

func init() { proto.RegisterFile("pm_optics_paras.proto", fileDescriptor_380a2e2d176d5301) }

var fileDescriptor_380a2e2d176d5301 = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x58, 0x4f, 0x6b, 0x13, 0x4d,
	0x1c, 0x66, 0xd2, 0x36, 0x6d, 0xe6, 0x7d, 0x43, 0x9b, 0x79, 0xdb, 0x97, 0x51, 0x3c, 0x94, 0x0a,
	0x52, 0x50, 0x62, 0x9b, 0xb4, 0xfe, 0xff, 0xdb, 0xda, 0x2a, 0x88, 0x58, 0x92, 0x82, 0x78, 0x1a,
	0x26, 0xb3, 0x93, 0xcd, 0xc0, 0xce, 0x9f, 0xce, 0x6e, 0x9a, 0x7c, 0x19, 0x3f, 0x83, 0xe0, 0x41,
	0xf0, 0xe0, 0x49, 0xbc, 0x88, 0x08, 0xde, 0x3c, 0x78, 0xf7, 0x6b, 0xc8, 0xce, 0x6c, 0x92, 0x26,
	0xb6, 0x50, 0x10, 0xc1, 0xed, 0x6d, 0x7f, 0xcf, 0x3c, 0xcf, 0x93, 0x79, 0x9e, 0x9d, 0x49, 0x20,
	0x70, 0xc9, 0x48, 0xa2, 0x4d, 0x22, 0x58, 0x4c, 0x0c, 0xb5, 0x34, 0xae, 0x1a, 0xab, 0x13, 0x8d,
	0x5e, 0x01, 0x26, 0x62, 0xa6, 0x89, 0xd0, 0x31, 0xe9, 0x5b, 0x62, 0x24, 0x57, 0xa1, 0x50, 0x9c,
	0x68, 0xc3, 0x6d, 0xd5, 0x70, 0xdb, 0xd6, 0x56, 0x52, 0xc5, 0x38, 0x91, 0x54, 0xd1, 0x90, 0x4b,
	0xae, 0x92, 0xaa, 0xb7, 0xa9, 0x0e, 0xdc, 0xb4, 0x4d, 0xc6, 0x86, 0xc1, 0x33, 0xeb, 0x5a, 0x3b,
	0xa2, 0x93, 0x98, 0x33, 0xad, 0x82, 0xfa, 0xda, 0xe4, 0x4c, 0xc6, 0xed, 0xc6, 0xe1, 0x95, 0x2d,
	0xb8, 0x38, 0xb1, 0x71, 0xf2, 0x74, 0xe7, 0x65, 0x13, 0x21, 0x38, 0xad, 0xa8, 0xe4, 0x18, 0x2c,
	0x83, 0xd5, 0x52, 0xc3, 0x3d, 0xa3, 0xff, 0x61, 0x51, 0x75, 0x65, 0x8b, 0x5b, 0x5c, 0x58, 0x06,
	0xab, 0xe5, 0x46, 0x36, 0xad, 0xbc, 0x2e, 0xc0, 0xff, 0xc6, 0x4d, 0x24, 0x4f, 0xb8, 0x45, 0x8b,
	0x70, 0xe6, 0x90, 0x46, 0x22, 0x70, 0x26, 0x73, 0x0d, 0x3f, 0x20, 0x0c, 0x67, 0xa5, 0x50, 0x42,
	0x76, 0xa5, 0xb3, 0xa9, 0x34, 0x06, 0x63, 0xba, 0x42, 0x0f, 0xb9, 0xa5, 0x21, 0xc7, 0x53, 0x7e,
	0x25, 0x1b, 0x9d, 0x86, 0xf6, 0x9d, 0x66, 0x3a, 0xd3, 0xf8, 0x11, 0x5d, 0x86, 0x95, 0x4c, 0x4e,
	0x92, 0x8e, 0xe5, 0x71, 0x47, 0x47, 0x01, 0x9e, 0x71, 0x9c, 0x85, 0x6c, 0x61, 0x7f, 0x80, 0xa3,
	0x2b, 0x10, 0x0d, 0xc9, 0x8c, 0x12, 0xcb, 0xd3, 0x3a, 0x71, 0xd1, 0xed, 0x6e, 0xc8, 0x66, 0xb4,
	0xe1, 0x70, 0x67, 0xed, 0x3f, 0xe5, 0x88, 0xf5, 0x6c, 0x66, 0xed, 0x17, 0xc6, 0xad, 0x07, 0xe4,
	0x91, 0xf5, 0x5c, 0x66, 0x9d, 0xb1, 0x07, 0xd6, 0x2b, 0x6f, 0x0b, 0xf0, 0xfc, 0x31, 0x8d, 0x91,
	0x38, 0xb1, 0x42, 0x85, 0xa7, 0x2b, 0xae, 0x74, 0x62, 0x71, 0xa5, 0x13, 0x8b, 0x2b, 0x9d, 0xa2,
	0xb8, 0xd2, 0x9f, 0x2b, 0xae, 0xf4, 0xdb, 0xc5, 0xbd, 0x3f, 0x07, 0xe7, 0x27, 0xce, 0x6b, 0xda,
	0x96, 0x50, 0x01, 0xef, 0xe3, 0x9a, 0x3b, 0x95, 0x7e, 0x18, 0x75, 0x58, 0x3f, 0xda, 0xe1, 0x05,
	0x58, 0x4a, 0x84, 0xe4, 0x71, 0x42, 0xa5, 0xc1, 0x1b, 0x6e, 0x4b, 0x23, 0x00, 0x5d, 0x82, 0xf3,
	0x11, 0x8d, 0x13, 0xc2, 0x22, 0x4e, 0x2d, 0x49, 0x71, 0xbc, 0xe9, 0x38, 0xe5, 0x14, 0xde, 0x4e,
	0xd1, 0x7d, 0x21, 0x39, 0x5a, 0x87, 0x4b, 0x23, 0xde, 0xfa, 0x26, 0x91, 0x42, 0x79, 0xf6, 0x35,
	0xc7, 0x46, 0x43, 0xf6, 0xfa, 0xe6, 0x33, 0xa1, 0x7e, 0x95, 0xd4, 0xd7, 0xd2, 0x6b, 0xe8, 0x25,
	0xd7, 0x27, 0x24, 0xf5, 0xb5, 0x26, 0x67, 0x4e, 0x72, 0x15, 0x2e, 0x8e, 0x24, 0xb5, 0x0d, 0xd2,
	0xc9, 0xb6, 0x74, 0xc3, 0x29, 0x2a, 0x43, 0x45, 0x6d, 0xe3, 0x89, 0xdf, 0xd6, 0x45, 0x58, 0x8e,
	0x39, 0x4b, 0xcd, 0xbb, 0xc6, 0xb5, 0x78, 0xd3, 0x45, 0xff, 0xd7, 0x81, 0x4d, 0x8f, 0xa1, 0x8f,
	0x00, 0x4e, 0x45, 0x2d, 0x86, 0x6f, 0x2d, 0x83, 0xd5, 0x7f, 0x6a, 0x6f, 0x40, 0xf5, 0xaf, 0xfe,
	0x7e, 0xaa, 0x1e, 0x73, 0x4b, 0x1a, 0xe9, 0xfe, 0xd1, 0x57, 0x00, 0x8b, 0x51, 0x8b, 0x11, 0xc3,
	0xf0, 0x6d, 0x17, 0xe5, 0x5d, 0x0e, 0xa3, 0x64, 0x17, 0xbe, 0x31, 0x13, 0xb5, 0xd8, 0x1e, 0x43,
	0x9f, 0x01, 0x9c, 0xd2, 0x26, 0xc1, 0x77, 0x72, 0x1f, 0x28, 0x8d, 0x91, 0xc5, 0xb1, 0xf8, 0xee,
	0x59, 0x88, 0x63, 0xd1, 0x07, 0x00, 0x0b, 0x2c, 0xc0, 0xf7, 0xf2, 0x7b, 0x71, 0x0a, 0x2c, 0x70,
	0x2f, 0x25, 0x08, 0x03, 0x7c, 0x3f, 0xff, 0x2f, 0x25, 0x08, 0x7d, 0x1c, 0x23, 0x03, 0xfc, 0x20,
	0xff, 0x71, 0x8c, 0x0c, 0xd0, 0x17, 0x00, 0xa7, 0x75, 0xac, 0x2c, 0x7e, 0x98, 0xfb, 0x3c, 0x2e,
	0x07, 0xfa, 0x01, 0x60, 0x85, 0x71, 0x95, 0xe2, 0x3d, 0x7a, 0xc8, 0x23, 0xae, 0xc2, 0xa4, 0x83,
	0xb7, 0x72, 0x9f, 0x6e, 0xc1, 0x87, 0x7a, 0x31, 0xcc, 0xe4, 0x4f, 0x62, 0x10, 0xe1, 0xed, 0x33,
	0x70, 0x12, 0x83, 0xc8, 0xc7, 0x61, 0x16, 0x3f, 0x3a, 0x03, 0x71, 0x98, 0x45, 0x9f, 0x00, 0x2c,
	0x18, 0x85, 0x77, 0x72, 0x9f, 0xa6, 0x60, 0x14, 0xfa, 0x06, 0x20, 0xb4, 0x7d, 0x12, 0x8b, 0x90,
	0x18, 0xdd, 0xc3, 0xbb, 0xb9, 0x0f, 0x35, 0x67, 0xfb, 0x4d, 0x11, 0xee, 0xe9, 0x1e, 0xfa, 0x0e,
	0xe0, 0x42, 0xa4, 0x7b, 0x2e, 0x5b, 0xdb, 0xf2, 0x03, 0xa2, 0xdb, 0x6d, 0xfc, 0x38, 0xbf, 0x3f,
	0xb9, 0xe5, 0x48, 0xf7, 0x9a, 0x22, 0xdc, 0xb5, 0xfc, 0xe0, 0x79, 0xbb, 0xdd, 0x2a, 0xba, 0x7f,
	0x05, 0xea, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x26, 0xd4, 0xf8, 0xec, 0x2e, 0x10, 0x00, 0x00,
}