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

package cisco_ios_xr_pmengine_oper_performance_management_optics_optics_ports_optics_port_optics_current_optics_minute15_optics_minute15_optics_optics_minute15_optic

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
	ConfiguredMinThresh  string   `protobuf:"bytes,6,opt,name=configured_min_thresh,json=configuredMinThresh,proto3" json:"configured_min_thresh,omitempty"`
	MinimumTcaReport     bool     `protobuf:"varint,7,opt,name=minimum_tca_report,json=minimumTcaReport,proto3" json:"minimum_tca_report,omitempty"`
	MaximumThreshold     int32    `protobuf:"zigzag32,8,opt,name=maximum_threshold,json=maximumThreshold,proto3" json:"maximum_threshold,omitempty"`
	ConfiguredMaxThresh  string   `protobuf:"bytes,9,opt,name=configured_max_thresh,json=configuredMaxThresh,proto3" json:"configured_max_thresh,omitempty"`
	MaximumTcaReport     bool     `protobuf:"varint,10,opt,name=maximum_tca_report,json=maximumTcaReport,proto3" json:"maximum_tca_report,omitempty"`
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

func (m *PmOpticsParameter) GetConfiguredMinThresh() string {
	if m != nil {
		return m.ConfiguredMinThresh
	}
	return ""
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

func (m *PmOpticsParameter) GetConfiguredMaxThresh() string {
	if m != nil {
		return m.ConfiguredMaxThresh
	}
	return ""
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
	ConfiguredMinThresh  string   `protobuf:"bytes,6,opt,name=configured_min_thresh,json=configuredMinThresh,proto3" json:"configured_min_thresh,omitempty"`
	MinimumTcaReport     bool     `protobuf:"varint,7,opt,name=minimum_tca_report,json=minimumTcaReport,proto3" json:"minimum_tca_report,omitempty"`
	MaximumThreshold     string   `protobuf:"bytes,8,opt,name=maximum_threshold,json=maximumThreshold,proto3" json:"maximum_threshold,omitempty"`
	ConfiguredMaxThresh  string   `protobuf:"bytes,9,opt,name=configured_max_thresh,json=configuredMaxThresh,proto3" json:"configured_max_thresh,omitempty"`
	MaximumTcaReport     bool     `protobuf:"varint,10,opt,name=maximum_tca_report,json=maximumTcaReport,proto3" json:"maximum_tca_report,omitempty"`
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

func (m *PmOpticsParameterString) GetConfiguredMinThresh() string {
	if m != nil {
		return m.ConfiguredMinThresh
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

func (m *PmOpticsParameterString) GetConfiguredMaxThresh() string {
	if m != nil {
		return m.ConfiguredMaxThresh
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
	AmpliGain            *PmOpticsParameterString `protobuf:"bytes,72,opt,name=ampli_gain,json=ampliGain,proto3" json:"ampli_gain,omitempty"`
	AmpliGainTilt        *PmOpticsParameterString `protobuf:"bytes,73,opt,name=ampli_gain_tilt,json=ampliGainTilt,proto3" json:"ampli_gain_tilt,omitempty"`
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

func (m *PmOpticsParas) GetAmpliGain() *PmOpticsParameterString {
	if m != nil {
		return m.AmpliGain
	}
	return nil
}

func (m *PmOpticsParas) GetAmpliGainTilt() *PmOpticsParameterString {
	if m != nil {
		return m.AmpliGainTilt
	}
	return nil
}

func init() {
	proto.RegisterType((*PmOpticsParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management.optics.optics_ports.optics_port.optics_current.optics_minute15.optics_minute15_optics.optics_minute15_optic.pm_optics_paras_KEYS")
	proto.RegisterType((*PmOpticsParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management.optics.optics_ports.optics_port.optics_current.optics_minute15.optics_minute15_optics.optics_minute15_optic.pm_optics_parameter")
	proto.RegisterType((*PmOpticsParameterString)(nil), "cisco_ios_xr_pmengine_oper.performance_management.optics.optics_ports.optics_port.optics_current.optics_minute15.optics_minute15_optics.optics_minute15_optic.pm_optics_parameter_string")
	proto.RegisterType((*PmOpticsParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management.optics.optics_ports.optics_port.optics_current.optics_minute15.optics_minute15_optics.optics_minute15_optic.pm_optics_paras")
}

func init() { proto.RegisterFile("pm_optics_paras.proto", fileDescriptor_380a2e2d176d5301) }

var fileDescriptor_380a2e2d176d5301 = []byte{
	// 816 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x98, 0x4d, 0x6f, 0x23, 0x35,
	0x18, 0xc7, 0xe5, 0xa4, 0xcd, 0x76, 0x0c, 0x51, 0x1b, 0x6f, 0x8b, 0x2c, 0xb4, 0x87, 0xaa, 0x48,
	0xa8, 0x12, 0x28, 0x6c, 0x92, 0x86, 0xf7, 0xd7, 0x5d, 0x76, 0x5b, 0x84, 0x10, 0x55, 0x52, 0x09,
	0x71, 0xb2, 0x1c, 0x8f, 0x33, 0xb1, 0x34, 0x7e, 0xa9, 0xc7, 0x69, 0xf2, 0x4d, 0x38, 0xc1, 0x87,
	0xe0, 0xc6, 0x19, 0x71, 0x41, 0x08, 0x89, 0x1b, 0x48, 0x9c, 0xb8, 0xf0, 0x2d, 0x10, 0x1a, 0x7b,
	0x92, 0x69, 0xd2, 0x56, 0xea, 0x6d, 0x3b, 0xbd, 0xcd, 0xf3, 0xf8, 0x79, 0x1e, 0xff, 0x7f, 0xfe,
	0x8f, 0x27, 0x52, 0xe0, 0x9e, 0x91, 0x44, 0x1b, 0x27, 0x58, 0x46, 0x0c, 0xb5, 0x34, 0x6b, 0x1b,
	0xab, 0x9d, 0x46, 0xdf, 0x03, 0x26, 0x32, 0xa6, 0x89, 0xd0, 0x19, 0x99, 0x5b, 0x62, 0x24, 0x57,
	0x89, 0x50, 0x9c, 0x68, 0xc3, 0x6d, 0xdb, 0x70, 0x3b, 0xd6, 0x56, 0x52, 0xc5, 0x38, 0x91, 0x54,
	0xd1, 0x84, 0x4b, 0xae, 0x5c, 0x3b, 0x8c, 0x69, 0x2f, 0xa6, 0x69, 0xeb, 0x56, 0x82, 0xc5, 0x33,
	0x9b, 0x5a, 0x5b, 0x96, 0x13, 0x29, 0xd4, 0xd4, 0xf1, 0x4e, 0x7f, 0x3d, 0x26, 0xab, 0xe3, 0x56,
	0xd3, 0x07, 0x4f, 0xe0, 0xee, 0x9a, 0x70, 0xf2, 0xe5, 0xb3, 0x6f, 0x87, 0x08, 0xc1, 0x0d, 0x45,
	0x25, 0xc7, 0x60, 0x1f, 0x1c, 0x46, 0x03, 0xff, 0x8c, 0x5e, 0x81, 0x0d, 0x35, 0x95, 0x23, 0x6e,
	0x71, 0x6d, 0x1f, 0x1c, 0x36, 0x07, 0x45, 0x74, 0xf0, 0x5d, 0x1d, 0x3e, 0x5c, 0x1d, 0x22, 0xb9,
	0xe3, 0x16, 0xed, 0xc2, 0xcd, 0x0b, 0x9a, 0x8a, 0xd8, 0x0f, 0xd9, 0x1a, 0x84, 0x00, 0x61, 0xf8,
	0x40, 0x0a, 0x25, 0xe4, 0x54, 0xfa, 0x31, 0xad, 0xc1, 0x22, 0xcc, 0x57, 0xe8, 0x05, 0xb7, 0x34,
	0xe1, 0xb8, 0x1e, 0x56, 0x8a, 0xd0, 0xf7, 0xd0, 0xb9, 0xef, 0xd9, 0x28, 0x7a, 0x42, 0x88, 0xde,
	0x80, 0xad, 0xa2, 0x9d, 0xb8, 0x89, 0xe5, 0xd9, 0x44, 0xa7, 0x31, 0xde, 0xf4, 0x35, 0x3b, 0xc5,
	0xc2, 0xd9, 0x22, 0x8f, 0xba, 0x70, 0x8f, 0x69, 0x35, 0x16, 0xc9, 0xd4, 0xf2, 0x38, 0x3f, 0x89,
	0xa2, 0x07, 0x37, 0x3c, 0xe5, 0xc3, 0x72, 0xf1, 0x2b, 0xa1, 0x42, 0x1b, 0x7a, 0x13, 0xa2, 0xe5,
	0x06, 0x8c, 0x12, 0xcb, 0x73, 0x0b, 0xf0, 0x03, 0x4f, 0xb4, 0xdc, 0x81, 0xd1, 0x81, 0xcf, 0x7b,
	0x39, 0x41, 0xd9, 0x25, 0x39, 0x5b, 0x85, 0x9c, 0xb0, 0x70, 0xa3, 0x1c, 0x3a, 0x5f, 0xc8, 0x89,
	0xae, 0xc8, 0xa1, 0xf3, 0x4b, 0x72, 0x16, 0x1b, 0x94, 0x72, 0x60, 0x21, 0xa7, 0xd8, 0x61, 0x21,
	0xe7, 0xe0, 0x87, 0x3a, 0x7c, 0xf5, 0x1a, 0x67, 0x48, 0xe6, 0xac, 0x50, 0xc9, 0xed, 0x0c, 0x8a,
	0x6e, 0x34, 0x28, 0xba, 0xd1, 0xa0, 0xe8, 0x16, 0x06, 0x45, 0x77, 0xcb, 0xa0, 0xe8, 0x85, 0x18,
	0xf4, 0xdf, 0x23, 0xb8, 0xbd, 0x76, 0xff, 0x72, 0x57, 0x84, 0x8a, 0xf9, 0x1c, 0x77, 0xfd, 0x2d,
	0x0b, 0x41, 0xe9, 0x55, 0xef, 0xb2, 0x57, 0x8f, 0x60, 0xe4, 0x84, 0xe4, 0x99, 0xa3, 0xd2, 0xe0,
	0x23, 0xaf, 0xaa, 0x4c, 0xa0, 0xd7, 0xe1, 0x76, 0x4a, 0x33, 0x47, 0x58, 0xca, 0xa9, 0x25, 0x79,
	0x1e, 0xf7, 0x7d, 0x4d, 0x33, 0x4f, 0x3f, 0xcd, 0xb3, 0x67, 0x42, 0x72, 0xd4, 0x81, 0x7b, 0x65,
	0x5d, 0xa7, 0x1f, 0x0e, 0x3e, 0xaf, 0x7e, 0xdb, 0x57, 0xa3, 0x65, 0x75, 0xa7, 0x9f, 0x9f, 0xfb,
	0x95, 0x96, 0xde, 0x63, 0x92, 0x71, 0x16, 0x5a, 0xde, 0x59, 0x6b, 0xe9, 0x3d, 0x1e, 0x72, 0xe6,
	0x5b, 0xde, 0x82, 0xbb, 0x65, 0x4b, 0xf7, 0x88, 0x4c, 0x0a, 0x49, 0xef, 0xfa, 0x8e, 0xd6, 0xb2,
	0xa3, 0x7b, 0x74, 0x12, 0x64, 0xbd, 0x06, 0x9b, 0x19, 0x67, 0xf9, 0xf0, 0xa9, 0xf1, 0xa7, 0xf8,
	0x9e, 0x47, 0x7f, 0xd9, 0x27, 0x87, 0x21, 0x87, 0x7e, 0x01, 0xb0, 0x9e, 0x8e, 0x18, 0x7e, 0x7f,
	0x1f, 0x1c, 0xbe, 0xd4, 0xfd, 0x11, 0xb4, 0xef, 0xf4, 0xf7, 0xb6, 0x7d, 0xcd, 0x6d, 0x1c, 0xe4,
	0xfa, 0xd1, 0x1f, 0x00, 0x36, 0xd2, 0x11, 0x23, 0x86, 0xe1, 0x0f, 0x3c, 0xca, 0x4f, 0x15, 0x44,
	0x29, 0x3e, 0x2c, 0x83, 0xcd, 0x74, 0xc4, 0x4e, 0x19, 0xfa, 0x0d, 0xc0, 0xba, 0x36, 0x0e, 0x7f,
	0x58, 0x79, 0xa0, 0x1c, 0xa3, 0xc0, 0xb1, 0xf8, 0xa3, 0xfb, 0x80, 0x63, 0xd1, 0xcf, 0x00, 0xd6,
	0x58, 0x8c, 0x3f, 0xae, 0xee, 0xc5, 0xa9, 0xb1, 0xd8, 0x9b, 0x12, 0x27, 0x31, 0xfe, 0xa4, 0xfa,
	0xa6, 0xc4, 0x49, 0xc0, 0x31, 0x32, 0xc6, 0x9f, 0x56, 0x1f, 0xc7, 0xc8, 0x18, 0xfd, 0x0e, 0xe0,
	0x86, 0xce, 0x94, 0xc5, 0x9f, 0x55, 0x9e, 0xc7, 0x73, 0xa0, 0x7f, 0x01, 0x6c, 0x31, 0xae, 0xf2,
	0xfc, 0x8c, 0x5e, 0xf0, 0x94, 0xab, 0xc4, 0x4d, 0xf0, 0x93, 0xca, 0xd3, 0xed, 0x04, 0xa8, 0x6f,
	0x96, 0x4c, 0xe1, 0x4d, 0x8c, 0x53, 0xfc, 0xf4, 0x1e, 0xbc, 0x89, 0x71, 0x1a, 0x70, 0x98, 0xc5,
	0x9f, 0xdf, 0x03, 0x1c, 0x66, 0xd1, 0xaf, 0x00, 0xd6, 0x8c, 0xc2, 0xcf, 0x2a, 0x4f, 0x53, 0x33,
	0x0a, 0xfd, 0x09, 0x20, 0xb4, 0x73, 0x92, 0x89, 0x84, 0x18, 0x3d, 0xc3, 0xcf, 0x2b, 0x0f, 0xb5,
	0x65, 0xe7, 0x43, 0x91, 0x9c, 0xea, 0x19, 0xfa, 0x1b, 0xc0, 0x9d, 0x54, 0xcf, 0x3c, 0xdb, 0xd8,
	0xf2, 0x73, 0xa2, 0xc7, 0x63, 0x7c, 0x5c, 0xdd, 0x9f, 0xdc, 0x66, 0xaa, 0x67, 0x43, 0x91, 0x3c,
	0xb7, 0xfc, 0xfc, 0xeb, 0xf1, 0x18, 0xfd, 0x05, 0x20, 0xa4, 0xd2, 0xa4, 0x82, 0x24, 0x54, 0x28,
	0x7c, 0x52, 0x79, 0xe7, 0x22, 0x4f, 0x73, 0x4c, 0x85, 0x42, 0xff, 0x00, 0xb8, 0x5d, 0xb2, 0x11,
	0x27, 0x52, 0x87, 0xbf, 0xa8, 0x3c, 0x60, 0x73, 0x09, 0x78, 0x26, 0x52, 0x37, 0x6a, 0xf8, 0xbf,
	0xa9, 0x7a, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xcf, 0x5e, 0x54, 0xbf, 0x12, 0x00, 0x00,
}
