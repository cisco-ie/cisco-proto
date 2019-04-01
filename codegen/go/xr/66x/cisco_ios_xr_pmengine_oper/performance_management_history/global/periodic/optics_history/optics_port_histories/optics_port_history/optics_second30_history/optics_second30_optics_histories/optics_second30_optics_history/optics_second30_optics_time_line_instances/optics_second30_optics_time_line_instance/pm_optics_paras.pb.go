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

package cisco_ios_xr_pmengine_oper_performance_management_history_global_periodic_optics_history_optics_port_histories_optics_port_history_optics_second30_history_optics_second30_optics_histories_optics_second30_optics_history_optics_second30_optics_time_line_instances_optics_second30_optics_time_line_instance

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
	Number_1             uint32   `protobuf:"varint,3,opt,name=number_1,json=number1,proto3" json:"number_1,omitempty"`
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

func (m *PmOpticsParas_KEYS) GetNumber_1() uint32 {
	if m != nil {
		return m.Number_1
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
	proto.RegisterType((*PmOpticsParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.optics_history.optics_port_histories.optics_port_history.optics_second30_history.optics_second30_optics_histories.optics_second30_optics_history.optics_second30_optics_time_line_instances.optics_second30_optics_time_line_instance.pm_optics_paras_KEYS")
	proto.RegisterType((*PmOpticsParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.optics_history.optics_port_histories.optics_port_history.optics_second30_history.optics_second30_optics_histories.optics_second30_optics_history.optics_second30_optics_time_line_instances.optics_second30_optics_time_line_instance.pm_optics_parameter")
	proto.RegisterType((*PmOpticsParameterString)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.optics_history.optics_port_histories.optics_port_history.optics_second30_history.optics_second30_optics_histories.optics_second30_optics_history.optics_second30_optics_time_line_instances.optics_second30_optics_time_line_instance.pm_optics_parameter_string")
	proto.RegisterType((*PmOpticsParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.optics_history.optics_port_histories.optics_port_history.optics_second30_history.optics_second30_optics_histories.optics_second30_optics_history.optics_second30_optics_time_line_instances.optics_second30_optics_time_line_instance.pm_optics_paras")
}

func init() { proto.RegisterFile("pm_optics_paras.proto", fileDescriptor_380a2e2d176d5301) }

var fileDescriptor_380a2e2d176d5301 = []byte{
	// 901 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x99, 0x4b, 0x6b, 0x2c, 0x45,
	0x14, 0xc7, 0x49, 0x4d, 0x5e, 0x5d, 0x3a, 0x24, 0x53, 0x37, 0x91, 0x52, 0x5c, 0x84, 0x08, 0x12,
	0x50, 0xe6, 0x66, 0x66, 0x12, 0xdf, 0xef, 0xeb, 0x7d, 0x88, 0x88, 0x97, 0x9e, 0x80, 0xb8, 0x90,
	0xa2, 0xa6, 0xba, 0xa6, 0xa7, 0xa0, 0x5e, 0xb7, 0xba, 0x93, 0x19, 0x3f, 0x89, 0x3b, 0xbf, 0x82,
	0x5f, 0xc2, 0x9d, 0x6b, 0x71, 0x23, 0x22, 0xb8, 0x53, 0x70, 0xe5, 0xc6, 0x9d, 0x74, 0x75, 0xf7,
	0x4c, 0x6e, 0x32, 0x03, 0x11, 0x3a, 0xe8, 0xa2, 0x76, 0x73, 0x4e, 0x9d, 0x7f, 0x9d, 0x53, 0xbf,
	0x53, 0x7d, 0x9a, 0x66, 0xe0, 0xbe, 0x55, 0xc4, 0xd8, 0x5c, 0xb0, 0x8c, 0x58, 0xea, 0x68, 0xd6,
	0xb5, 0xce, 0xe4, 0x06, 0x7d, 0x07, 0x98, 0xc8, 0x98, 0x21, 0xc2, 0x64, 0x64, 0xe6, 0x88, 0x55,
	0x5c, 0xa7, 0x42, 0x73, 0x62, 0x2c, 0x77, 0x5d, 0xcb, 0xdd, 0xd8, 0x38, 0x45, 0x35, 0xe3, 0x44,
	0x51, 0x4d, 0x53, 0xae, 0xb8, 0xce, 0xc9, 0x44, 0x64, 0xb9, 0x71, 0x5f, 0x77, 0x53, 0x69, 0x46,
	0x54, 0x16, 0x51, 0xc2, 0x24, 0x82, 0x75, 0xab, 0xed, 0xeb, 0xe5, 0x3a, 0x9b, 0x71, 0xb5, 0x44,
	0xf0, 0x6c, 0x89, 0x77, 0x1e, 0x99, 0x71, 0x66, 0x74, 0x32, 0x38, 0x5e, 0xe9, 0x7f, 0x2a, 0xc1,
	0xa5, 0xcd, 0x96, 0x07, 0xac, 0xd4, 0xe7, 0x42, 0x71, 0x22, 0x8b, 0x53, 0x0a, 0x9d, 0xe5, 0xc5,
	0xe9, 0xb2, 0x9b, 0x87, 0x1e, 0x7e, 0x05, 0xf7, 0xae, 0xa0, 0x24, 0x9f, 0xde, 0xff, 0x72, 0x88,
	0x10, 0x5c, 0xd7, 0x54, 0x71, 0xbc, 0x76, 0xb0, 0x76, 0x14, 0xc5, 0xfe, 0x37, 0x7a, 0x0e, 0x6e,
	0xea, 0x73, 0x35, 0xe2, 0x0e, 0x83, 0x83, 0xb5, 0xa3, 0x76, 0x5c, 0x59, 0xe8, 0x79, 0xb8, 0x5d,
	0xfe, 0x22, 0x3d, 0xdc, 0xf2, 0x2b, 0x5b, 0xa5, 0xdd, 0x3b, 0xfc, 0xa6, 0x05, 0xef, 0x3c, 0xbd,
	0xbf, 0xe2, 0x39, 0x77, 0x68, 0x0f, 0x6e, 0x5c, 0x50, 0x29, 0x12, 0xbf, 0xff, 0x76, 0x5c, 0x1a,
	0x08, 0xc3, 0x2d, 0x25, 0xb4, 0x50, 0xe7, 0xca, 0x67, 0xe8, 0xc4, 0xb5, 0x59, 0xac, 0xd0, 0x0b,
	0xee, 0x68, 0xca, 0x7d, 0x86, 0x4e, 0x5c, 0x9b, 0x5e, 0x43, 0x67, 0x5e, 0xb3, 0x5e, 0x69, 0x4a,
	0x13, 0xbd, 0x02, 0x3b, 0x95, 0x9c, 0xe4, 0x13, 0xc7, 0xb3, 0x89, 0x91, 0x09, 0xde, 0xf0, 0x31,
	0xbb, 0xd5, 0xc2, 0x59, 0xed, 0x47, 0x7d, 0xb8, 0xcf, 0x8c, 0x1e, 0x8b, 0xf4, 0xdc, 0xf1, 0x84,
	0x28, 0xa1, 0x2b, 0x0d, 0xde, 0xf4, 0x00, 0xee, 0x2c, 0x16, 0x3f, 0x13, 0xba, 0x94, 0xa1, 0x57,
	0x21, 0x9a, 0x27, 0x60, 0x94, 0x38, 0x5e, 0x5c, 0x03, 0xbc, 0xe5, 0x4f, 0x34, 0xcf, 0xc0, 0x68,
	0xec, 0xfd, 0xbe, 0x9c, 0xb2, 0xb2, 0x4b, 0xe5, 0x6c, 0x57, 0xe5, 0x94, 0x0b, 0x2b, 0xcb, 0xa1,
	0xb3, 0xba, 0x9c, 0xe8, 0x5a, 0x39, 0x74, 0x76, 0xa9, 0x9c, 0x3a, 0xc1, 0xa2, 0x1c, 0x58, 0x95,
	0x53, 0x65, 0xa8, 0xcb, 0x39, 0xfc, 0xb6, 0x05, 0x5f, 0x58, 0xd2, 0x19, 0x92, 0xe5, 0x4e, 0xe8,
	0xf4, 0x66, 0x0d, 0x8a, 0x56, 0x36, 0x28, 0x5a, 0xd9, 0xa0, 0xe8, 0x06, 0x0d, 0x8a, 0xfe, 0x5f,
	0x0d, 0x8a, 0xfe, 0x9b, 0x06, 0xdd, 0x85, 0x3b, 0x57, 0x1e, 0xcd, 0xa2, 0x2b, 0x42, 0x27, 0x7c,
	0x86, 0xfb, 0xfe, 0x31, 0x2b, 0x8d, 0x45, 0xaf, 0x06, 0x97, 0x7b, 0xf5, 0x22, 0x8c, 0x8a, 0xe7,
	0x3d, 0xcb, 0xa9, 0xb2, 0xf8, 0xc4, 0x57, 0xb5, 0x70, 0xa0, 0x97, 0xe1, 0x8e, 0xa4, 0x59, 0x4e,
	0x98, 0xe4, 0xd4, 0xf9, 0xc1, 0x80, 0x4f, 0x7d, 0x4c, 0xbb, 0x70, 0xdf, 0x2b, 0xbc, 0x67, 0x42,
	0x71, 0xd4, 0x83, 0xfb, 0x8b, 0xb8, 0xde, 0x69, 0x09, 0xbe, 0x88, 0x7e, 0xcd, 0x47, 0xa3, 0x79,
	0x74, 0xef, 0xb4, 0xe0, 0x7e, 0x4d, 0x32, 0x38, 0x2e, 0xa6, 0x50, 0x29, 0x79, 0xfd, 0x8a, 0x64,
	0x70, 0x3c, 0xe4, 0xcc, 0x4b, 0xee, 0xc2, 0xbd, 0x85, 0xa4, 0x7f, 0x42, 0x26, 0x55, 0x49, 0x6f,
	0x78, 0x45, 0x67, 0xae, 0xe8, 0x9f, 0x3c, 0x2a, 0xcb, 0x7a, 0x09, 0xb6, 0x33, 0xce, 0x8a, 0xcd,
	0xcf, 0xad, 0xa7, 0xf8, 0xa6, 0x3f, 0xfa, 0xb3, 0xde, 0x39, 0x2c, 0x7d, 0xe8, 0x27, 0x00, 0x5b,
	0x72, 0xc4, 0xf0, 0x5b, 0x07, 0x6b, 0x47, 0xcf, 0xf4, 0xbf, 0x07, 0xdd, 0xf0, 0x72, 0xf8, 0x57,
	0x2f, 0x87, 0xee, 0x92, 0xf9, 0x10, 0x17, 0x44, 0xd1, 0x6f, 0x00, 0x6e, 0xca, 0x11, 0x23, 0x96,
	0xe1, 0xb7, 0x3d, 0xdc, 0x1f, 0x02, 0xdc, 0x06, 0xe0, 0x56, 0xc3, 0x37, 0xde, 0x90, 0x23, 0xf6,
	0x98, 0xa1, 0x5f, 0x00, 0x6c, 0x19, 0x9b, 0xe3, 0x77, 0x02, 0xe2, 0xc6, 0x11, 0x17, 0x60, 0x2b,
	0xc0, 0x0e, 0xbf, 0x1b, 0x00, 0xdf, 0x06, 0x60, 0x87, 0x7e, 0x04, 0x10, 0xb0, 0x04, 0xbf, 0x17,
	0x06, 0x70, 0x73, 0x03, 0x18, 0xb0, 0xc4, 0x5f, 0xdc, 0x24, 0x4d, 0xf0, 0xfb, 0xe1, 0xe2, 0x36,
	0x7f, 0x71, 0x93, 0xb4, 0x04, 0x6c, 0x55, 0x82, 0x3f, 0x08, 0x80, 0x9b, 0x07, 0x6c, 0x55, 0x82,
	0x7e, 0x05, 0x70, 0xdd, 0x64, 0xda, 0xe1, 0x0f, 0x03, 0xe1, 0xc6, 0x09, 0x7b, 0xb2, 0xe8, 0x6f,
	0x00, 0x3b, 0x8c, 0xeb, 0xc2, 0x3f, 0xa5, 0x17, 0x5c, 0x72, 0x9d, 0xe6, 0x13, 0xfc, 0x51, 0xe0,
	0xdd, 0x38, 0xef, 0xdd, 0x12, 0xf3, 0x17, 0x73, 0xca, 0xe5, 0xfc, 0x48, 0x24, 0xbe, 0x17, 0x68,
	0xdf, 0xc2, 0xfc, 0x48, 0x64, 0x09, 0x98, 0x39, 0xfc, 0x71, 0x00, 0x7c, 0x0b, 0x80, 0x99, 0x43,
	0x3f, 0x03, 0x08, 0xac, 0xc6, 0xf7, 0x03, 0xdf, 0xc6, 0xf9, 0x02, 0xab, 0xd1, 0xef, 0x00, 0x42,
	0x37, 0x23, 0x99, 0x48, 0x89, 0x35, 0x53, 0xfc, 0x20, 0x60, 0x6e, 0x1c, 0xf3, 0xb6, 0x9b, 0x0d,
	0x45, 0xfa, 0xd8, 0x4c, 0xd1, 0x9f, 0x00, 0xee, 0x4a, 0x33, 0xf5, 0xb4, 0xc7, 0x8e, 0x3f, 0x21,
	0x66, 0x3c, 0xc6, 0x0f, 0xc3, 0x47, 0x49, 0x73, 0x1f, 0x25, 0x6d, 0x69, 0xa6, 0x43, 0x91, 0x3e,
	0x70, 0xfc, 0xc9, 0xe7, 0xe3, 0x31, 0xfa, 0x03, 0x40, 0x48, 0x95, 0x95, 0x82, 0xa4, 0x54, 0x68,
	0xfc, 0x28, 0xdc, 0xee, 0xc6, 0x6f, 0x77, 0xe4, 0xf9, 0x3e, 0xa4, 0x42, 0xa3, 0xbf, 0x00, 0xdc,
	0x59, 0xd0, 0x26, 0xb9, 0x90, 0x39, 0xfe, 0x24, 0x20, 0x6f, 0x1c, 0x79, 0x7b, 0x8e, 0xfc, 0x4c,
	0xc8, 0x7c, 0xb4, 0xe9, 0xff, 0x73, 0x1c, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xfe, 0xfc,
	0x63, 0x8c, 0x1c, 0x00, 0x00,
}