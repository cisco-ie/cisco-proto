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
// source: pm_stm_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_stm_stm_ports_stm_port_stm_current_stm_minute15_stm_minute15stms_stm_minute15stm

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

type PmStmParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmStmParas_KEYS) Reset()         { *m = PmStmParas_KEYS{} }
func (m *PmStmParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmStmParas_KEYS) ProtoMessage()    {}
func (*PmStmParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_448b83ca56a736d2, []int{0}
}

func (m *PmStmParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmStmParas_KEYS.Unmarshal(m, b)
}
func (m *PmStmParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmStmParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmStmParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmStmParas_KEYS.Merge(m, src)
}
func (m *PmStmParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmStmParas_KEYS.Size(m)
}
func (m *PmStmParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmStmParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmStmParas_KEYS proto.InternalMessageInfo

func (m *PmStmParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmStmParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PmStmParameter struct {
	Data                 uint32   `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            uint32   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,3,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmStmParameter) Reset()         { *m = PmStmParameter{} }
func (m *PmStmParameter) String() string { return proto.CompactTextString(m) }
func (*PmStmParameter) ProtoMessage()    {}
func (*PmStmParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_448b83ca56a736d2, []int{1}
}

func (m *PmStmParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmStmParameter.Unmarshal(m, b)
}
func (m *PmStmParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmStmParameter.Marshal(b, m, deterministic)
}
func (m *PmStmParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmStmParameter.Merge(m, src)
}
func (m *PmStmParameter) XXX_Size() int {
	return xxx_messageInfo_PmStmParameter.Size(m)
}
func (m *PmStmParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PmStmParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PmStmParameter proto.InternalMessageInfo

func (m *PmStmParameter) GetData() uint32 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PmStmParameter) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PmStmParameter) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

type PmStmParameterRatio struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            string   `protobuf:"bytes,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,3,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmStmParameterRatio) Reset()         { *m = PmStmParameterRatio{} }
func (m *PmStmParameterRatio) String() string { return proto.CompactTextString(m) }
func (*PmStmParameterRatio) ProtoMessage()    {}
func (*PmStmParameterRatio) Descriptor() ([]byte, []int) {
	return fileDescriptor_448b83ca56a736d2, []int{2}
}

func (m *PmStmParameterRatio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmStmParameterRatio.Unmarshal(m, b)
}
func (m *PmStmParameterRatio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmStmParameterRatio.Marshal(b, m, deterministic)
}
func (m *PmStmParameterRatio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmStmParameterRatio.Merge(m, src)
}
func (m *PmStmParameterRatio) XXX_Size() int {
	return xxx_messageInfo_PmStmParameterRatio.Size(m)
}
func (m *PmStmParameterRatio) XXX_DiscardUnknown() {
	xxx_messageInfo_PmStmParameterRatio.DiscardUnknown(m)
}

var xxx_messageInfo_PmStmParameterRatio proto.InternalMessageInfo

func (m *PmStmParameterRatio) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *PmStmParameterRatio) GetThreshold() string {
	if m != nil {
		return m.Threshold
	}
	return ""
}

func (m *PmStmParameterRatio) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

type StmSectionEntry struct {
	SectionStatus        int32                `protobuf:"zigzag32,1,opt,name=section_status,json=sectionStatus,proto3" json:"section_status,omitempty"`
	SectionESs           *PmStmParameter      `protobuf:"bytes,2,opt,name=section_e_ss,json=sectionESs,proto3" json:"section_e_ss,omitempty"`
	SectionEsRs          *PmStmParameterRatio `protobuf:"bytes,3,opt,name=section_es_rs,json=sectionEsRs,proto3" json:"section_es_rs,omitempty"`
	SectionBbEs          *PmStmParameter      `protobuf:"bytes,4,opt,name=section_bb_es,json=sectionBbEs,proto3" json:"section_bb_es,omitempty"`
	SectionBbeRs         *PmStmParameterRatio `protobuf:"bytes,5,opt,name=section_bbe_rs,json=sectionBbeRs,proto3" json:"section_bbe_rs,omitempty"`
	SectionSeSs          *PmStmParameter      `protobuf:"bytes,6,opt,name=section_se_ss,json=sectionSeSs,proto3" json:"section_se_ss,omitempty"`
	SectionSesRs         *PmStmParameterRatio `protobuf:"bytes,7,opt,name=section_ses_rs,json=sectionSesRs,proto3" json:"section_ses_rs,omitempty"`
	SectionUaSs          *PmStmParameter      `protobuf:"bytes,8,opt,name=section_ua_ss,json=sectionUaSs,proto3" json:"section_ua_ss,omitempty"`
	SectionEBs           *PmStmParameter      `protobuf:"bytes,9,opt,name=section_e_bs,json=sectionEBs,proto3" json:"section_e_bs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StmSectionEntry) Reset()         { *m = StmSectionEntry{} }
func (m *StmSectionEntry) String() string { return proto.CompactTextString(m) }
func (*StmSectionEntry) ProtoMessage()    {}
func (*StmSectionEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_448b83ca56a736d2, []int{3}
}

func (m *StmSectionEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StmSectionEntry.Unmarshal(m, b)
}
func (m *StmSectionEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StmSectionEntry.Marshal(b, m, deterministic)
}
func (m *StmSectionEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StmSectionEntry.Merge(m, src)
}
func (m *StmSectionEntry) XXX_Size() int {
	return xxx_messageInfo_StmSectionEntry.Size(m)
}
func (m *StmSectionEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_StmSectionEntry.DiscardUnknown(m)
}

var xxx_messageInfo_StmSectionEntry proto.InternalMessageInfo

func (m *StmSectionEntry) GetSectionStatus() int32 {
	if m != nil {
		return m.SectionStatus
	}
	return 0
}

func (m *StmSectionEntry) GetSectionESs() *PmStmParameter {
	if m != nil {
		return m.SectionESs
	}
	return nil
}

func (m *StmSectionEntry) GetSectionEsRs() *PmStmParameterRatio {
	if m != nil {
		return m.SectionEsRs
	}
	return nil
}

func (m *StmSectionEntry) GetSectionBbEs() *PmStmParameter {
	if m != nil {
		return m.SectionBbEs
	}
	return nil
}

func (m *StmSectionEntry) GetSectionBbeRs() *PmStmParameterRatio {
	if m != nil {
		return m.SectionBbeRs
	}
	return nil
}

func (m *StmSectionEntry) GetSectionSeSs() *PmStmParameter {
	if m != nil {
		return m.SectionSeSs
	}
	return nil
}

func (m *StmSectionEntry) GetSectionSesRs() *PmStmParameterRatio {
	if m != nil {
		return m.SectionSesRs
	}
	return nil
}

func (m *StmSectionEntry) GetSectionUaSs() *PmStmParameter {
	if m != nil {
		return m.SectionUaSs
	}
	return nil
}

func (m *StmSectionEntry) GetSectionEBs() *PmStmParameter {
	if m != nil {
		return m.SectionEBs
	}
	return nil
}

type StmLineEntry struct {
	LineStatus           int32                `protobuf:"zigzag32,1,opt,name=line_status,json=lineStatus,proto3" json:"line_status,omitempty"`
	LineESs              *PmStmParameter      `protobuf:"bytes,2,opt,name=line_e_ss,json=lineESs,proto3" json:"line_e_ss,omitempty"`
	LineEsRs             *PmStmParameterRatio `protobuf:"bytes,3,opt,name=line_es_rs,json=lineEsRs,proto3" json:"line_es_rs,omitempty"`
	LineBbEs             *PmStmParameter      `protobuf:"bytes,4,opt,name=line_bb_es,json=lineBbEs,proto3" json:"line_bb_es,omitempty"`
	LineBbeRs            *PmStmParameterRatio `protobuf:"bytes,5,opt,name=line_bbe_rs,json=lineBbeRs,proto3" json:"line_bbe_rs,omitempty"`
	LineSeSs             *PmStmParameter      `protobuf:"bytes,6,opt,name=line_se_ss,json=lineSeSs,proto3" json:"line_se_ss,omitempty"`
	LineSesRs            *PmStmParameterRatio `protobuf:"bytes,7,opt,name=line_ses_rs,json=lineSesRs,proto3" json:"line_ses_rs,omitempty"`
	LineUaSs             *PmStmParameter      `protobuf:"bytes,8,opt,name=line_ua_ss,json=lineUaSs,proto3" json:"line_ua_ss,omitempty"`
	LineEBs              *PmStmParameter      `protobuf:"bytes,9,opt,name=line_e_bs,json=lineEBs,proto3" json:"line_e_bs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StmLineEntry) Reset()         { *m = StmLineEntry{} }
func (m *StmLineEntry) String() string { return proto.CompactTextString(m) }
func (*StmLineEntry) ProtoMessage()    {}
func (*StmLineEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_448b83ca56a736d2, []int{4}
}

func (m *StmLineEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StmLineEntry.Unmarshal(m, b)
}
func (m *StmLineEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StmLineEntry.Marshal(b, m, deterministic)
}
func (m *StmLineEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StmLineEntry.Merge(m, src)
}
func (m *StmLineEntry) XXX_Size() int {
	return xxx_messageInfo_StmLineEntry.Size(m)
}
func (m *StmLineEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_StmLineEntry.DiscardUnknown(m)
}

var xxx_messageInfo_StmLineEntry proto.InternalMessageInfo

func (m *StmLineEntry) GetLineStatus() int32 {
	if m != nil {
		return m.LineStatus
	}
	return 0
}

func (m *StmLineEntry) GetLineESs() *PmStmParameter {
	if m != nil {
		return m.LineESs
	}
	return nil
}

func (m *StmLineEntry) GetLineEsRs() *PmStmParameterRatio {
	if m != nil {
		return m.LineEsRs
	}
	return nil
}

func (m *StmLineEntry) GetLineBbEs() *PmStmParameter {
	if m != nil {
		return m.LineBbEs
	}
	return nil
}

func (m *StmLineEntry) GetLineBbeRs() *PmStmParameterRatio {
	if m != nil {
		return m.LineBbeRs
	}
	return nil
}

func (m *StmLineEntry) GetLineSeSs() *PmStmParameter {
	if m != nil {
		return m.LineSeSs
	}
	return nil
}

func (m *StmLineEntry) GetLineSesRs() *PmStmParameterRatio {
	if m != nil {
		return m.LineSesRs
	}
	return nil
}

func (m *StmLineEntry) GetLineUaSs() *PmStmParameter {
	if m != nil {
		return m.LineUaSs
	}
	return nil
}

func (m *StmLineEntry) GetLineEBs() *PmStmParameter {
	if m != nil {
		return m.LineEBs
	}
	return nil
}

type StmFarEndLineEntry struct {
	FarEndLineESs        *PmStmParameter      `protobuf:"bytes,1,opt,name=far_end_line_e_ss,json=farEndLineESs,proto3" json:"far_end_line_e_ss,omitempty"`
	FarEndLineEsRs       *PmStmParameterRatio `protobuf:"bytes,2,opt,name=far_end_line_es_rs,json=farEndLineEsRs,proto3" json:"far_end_line_es_rs,omitempty"`
	FarEndLineEbbEs      *PmStmParameter      `protobuf:"bytes,3,opt,name=far_end_line_ebb_es,json=farEndLineEbbEs,proto3" json:"far_end_line_ebb_es,omitempty"`
	FarEndLineBbeRs      *PmStmParameterRatio `protobuf:"bytes,4,opt,name=far_end_line_bbe_rs,json=farEndLineBbeRs,proto3" json:"far_end_line_bbe_rs,omitempty"`
	FarEndLineSeSs       *PmStmParameter      `protobuf:"bytes,5,opt,name=far_end_line_se_ss,json=farEndLineSeSs,proto3" json:"far_end_line_se_ss,omitempty"`
	FarEndLineSesRs      *PmStmParameterRatio `protobuf:"bytes,6,opt,name=far_end_line_ses_rs,json=farEndLineSesRs,proto3" json:"far_end_line_ses_rs,omitempty"`
	FarEndLineUaSs       *PmStmParameter      `protobuf:"bytes,7,opt,name=far_end_line_ua_ss,json=farEndLineUaSs,proto3" json:"far_end_line_ua_ss,omitempty"`
	FarEndLineEBs        *PmStmParameter      `protobuf:"bytes,8,opt,name=far_end_line_e_bs,json=farEndLineEBs,proto3" json:"far_end_line_e_bs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StmFarEndLineEntry) Reset()         { *m = StmFarEndLineEntry{} }
func (m *StmFarEndLineEntry) String() string { return proto.CompactTextString(m) }
func (*StmFarEndLineEntry) ProtoMessage()    {}
func (*StmFarEndLineEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_448b83ca56a736d2, []int{5}
}

func (m *StmFarEndLineEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StmFarEndLineEntry.Unmarshal(m, b)
}
func (m *StmFarEndLineEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StmFarEndLineEntry.Marshal(b, m, deterministic)
}
func (m *StmFarEndLineEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StmFarEndLineEntry.Merge(m, src)
}
func (m *StmFarEndLineEntry) XXX_Size() int {
	return xxx_messageInfo_StmFarEndLineEntry.Size(m)
}
func (m *StmFarEndLineEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_StmFarEndLineEntry.DiscardUnknown(m)
}

var xxx_messageInfo_StmFarEndLineEntry proto.InternalMessageInfo

func (m *StmFarEndLineEntry) GetFarEndLineESs() *PmStmParameter {
	if m != nil {
		return m.FarEndLineESs
	}
	return nil
}

func (m *StmFarEndLineEntry) GetFarEndLineEsRs() *PmStmParameterRatio {
	if m != nil {
		return m.FarEndLineEsRs
	}
	return nil
}

func (m *StmFarEndLineEntry) GetFarEndLineEbbEs() *PmStmParameter {
	if m != nil {
		return m.FarEndLineEbbEs
	}
	return nil
}

func (m *StmFarEndLineEntry) GetFarEndLineBbeRs() *PmStmParameterRatio {
	if m != nil {
		return m.FarEndLineBbeRs
	}
	return nil
}

func (m *StmFarEndLineEntry) GetFarEndLineSeSs() *PmStmParameter {
	if m != nil {
		return m.FarEndLineSeSs
	}
	return nil
}

func (m *StmFarEndLineEntry) GetFarEndLineSesRs() *PmStmParameterRatio {
	if m != nil {
		return m.FarEndLineSesRs
	}
	return nil
}

func (m *StmFarEndLineEntry) GetFarEndLineUaSs() *PmStmParameter {
	if m != nil {
		return m.FarEndLineUaSs
	}
	return nil
}

func (m *StmFarEndLineEntry) GetFarEndLineEBs() *PmStmParameter {
	if m != nil {
		return m.FarEndLineEBs
	}
	return nil
}

type PmStmParas struct {
	Index                uint32              `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool                `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            []string            `protobuf:"bytes,52,rep,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        []string            `protobuf:"bytes,53,rep,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   []string            `protobuf:"bytes,54,rep,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear24HrTime    []string            `protobuf:"bytes,55,rep,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Section              *StmSectionEntry    `protobuf:"bytes,56,opt,name=section,proto3" json:"section,omitempty"`
	Line                 *StmLineEntry       `protobuf:"bytes,57,opt,name=line,proto3" json:"line,omitempty"`
	FeLine               *StmFarEndLineEntry `protobuf:"bytes,58,opt,name=fe_line,json=feLine,proto3" json:"fe_line,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PmStmParas) Reset()         { *m = PmStmParas{} }
func (m *PmStmParas) String() string { return proto.CompactTextString(m) }
func (*PmStmParas) ProtoMessage()    {}
func (*PmStmParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_448b83ca56a736d2, []int{6}
}

func (m *PmStmParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmStmParas.Unmarshal(m, b)
}
func (m *PmStmParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmStmParas.Marshal(b, m, deterministic)
}
func (m *PmStmParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmStmParas.Merge(m, src)
}
func (m *PmStmParas) XXX_Size() int {
	return xxx_messageInfo_PmStmParas.Size(m)
}
func (m *PmStmParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmStmParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmStmParas proto.InternalMessageInfo

func (m *PmStmParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmStmParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmStmParas) GetTimestamp() []string {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PmStmParas) GetLastClearTime() []string {
	if m != nil {
		return m.LastClearTime
	}
	return nil
}

func (m *PmStmParas) GetLastClear15MinTime() []string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return nil
}

func (m *PmStmParas) GetLastClear24HrTime() []string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return nil
}

func (m *PmStmParas) GetSection() *StmSectionEntry {
	if m != nil {
		return m.Section
	}
	return nil
}

func (m *PmStmParas) GetLine() *StmLineEntry {
	if m != nil {
		return m.Line
	}
	return nil
}

func (m *PmStmParas) GetFeLine() *StmFarEndLineEntry {
	if m != nil {
		return m.FeLine
	}
	return nil
}

func init() {
	proto.RegisterType((*PmStmParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management.stm.stm_ports.stm_port.stm_current.stm_minute15.stm_minute15stms.stm_minute15stm.pm_stm_paras_KEYS")
	proto.RegisterType((*PmStmParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management.stm.stm_ports.stm_port.stm_current.stm_minute15.stm_minute15stms.stm_minute15stm.pm_stm_parameter")
	proto.RegisterType((*PmStmParameterRatio)(nil), "cisco_ios_xr_pmengine_oper.performance_management.stm.stm_ports.stm_port.stm_current.stm_minute15.stm_minute15stms.stm_minute15stm.pm_stm_parameter_ratio")
	proto.RegisterType((*StmSectionEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management.stm.stm_ports.stm_port.stm_current.stm_minute15.stm_minute15stms.stm_minute15stm.stmSectionEntry")
	proto.RegisterType((*StmLineEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management.stm.stm_ports.stm_port.stm_current.stm_minute15.stm_minute15stms.stm_minute15stm.stmLineEntry")
	proto.RegisterType((*StmFarEndLineEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management.stm.stm_ports.stm_port.stm_current.stm_minute15.stm_minute15stms.stm_minute15stm.stmFarEndLineEntry")
	proto.RegisterType((*PmStmParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management.stm.stm_ports.stm_port.stm_current.stm_minute15.stm_minute15stms.stm_minute15stm.pm_stm_paras")
}

func init() { proto.RegisterFile("pm_stm_paras.proto", fileDescriptor_448b83ca56a736d2) }

var fileDescriptor_448b83ca56a736d2 = []byte{
	// 865 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x58, 0xbd, 0x8e, 0x23, 0x45,
	0x10, 0x56, 0xdf, 0xee, 0xda, 0x3b, 0x65, 0xfb, 0x96, 0x6d, 0x8e, 0xd3, 0x04, 0x20, 0x2c, 0x4b,
	0xa0, 0x8d, 0x8c, 0xec, 0x5b, 0xf3, 0x97, 0x20, 0x19, 0x19, 0x21, 0x01, 0xc9, 0x0c, 0x04, 0x44,
	0xad, 0x9e, 0x71, 0xfb, 0x76, 0x24, 0xf7, 0x8c, 0xd5, 0xd5, 0x46, 0x47, 0xca, 0x13, 0x20, 0x81,
	0x08, 0xc8, 0x49, 0x08, 0x10, 0x88, 0x00, 0x9d, 0x08, 0x4e, 0x27, 0x02, 0xe0, 0x6d, 0xe0, 0x0d,
	0x50, 0x57, 0x8f, 0xed, 0xf1, 0x2c, 0x04, 0x44, 0x37, 0x93, 0x4d, 0x97, 0xaa, 0xbb, 0xbf, 0xfa,
	0xaa, 0xba, 0x3e, 0xbb, 0x80, 0x6f, 0xb4, 0x40, 0xab, 0xc5, 0x46, 0x1a, 0x89, 0xe3, 0x8d, 0x29,
	0x6c, 0xc1, 0xbf, 0x60, 0x69, 0x86, 0x69, 0x21, 0xb2, 0x02, 0xc5, 0x23, 0x23, 0x36, 0x5a, 0xe5,
	0x0f, 0xb3, 0x5c, 0x89, 0x62, 0xa3, 0xcc, 0x78, 0xa3, 0xcc, 0xaa, 0x30, 0x5a, 0xe6, 0xa9, 0x12,
	0x5a, 0xe6, 0xf2, 0xa1, 0xd2, 0x2a, 0xb7, 0x63, 0xb4, 0x7a, 0x4c, 0xe7, 0x14, 0xc6, 0xe2, 0xfe,
	0x8b, 0x3e, 0xd2, 0xad, 0x31, 0xa5, 0x8b, 0xd0, 0x59, 0xbe, 0xb5, 0x6a, 0x32, 0x3b, 0x5a, 0xa0,
	0xd5, 0x58, 0x37, 0x8c, 0xde, 0x81, 0xcb, 0x2a, 0x34, 0xf1, 0xc1, 0xe2, 0xd3, 0x98, 0x73, 0x38,
	0xcd, 0xa5, 0x56, 0x21, 0x1b, 0xb2, 0xab, 0x20, 0xa2, 0x6f, 0x7e, 0x1f, 0x3a, 0xf9, 0x56, 0x27,
	0xca, 0x84, 0x77, 0x86, 0xec, 0x6a, 0x10, 0x95, 0xab, 0x51, 0x0a, 0xcf, 0x55, 0x0e, 0xd0, 0xca,
	0x2a, 0xe3, 0xf6, 0x2f, 0xa5, 0x95, 0xb4, 0x7f, 0x10, 0xd1, 0x37, 0x7f, 0x11, 0x02, 0x7b, 0x63,
	0x14, 0xde, 0x14, 0xeb, 0x65, 0x79, 0xc4, 0xc1, 0xc0, 0x5f, 0x02, 0xb0, 0xa9, 0x14, 0x46, 0xb9,
	0x70, 0xc2, 0x93, 0x21, 0xbb, 0x3a, 0x8f, 0x02, 0x9b, 0xca, 0x88, 0x0c, 0xa3, 0x0c, 0xee, 0xd7,
	0x2f, 0x11, 0x46, 0xda, 0xac, 0x38, 0xba, 0x2a, 0xf8, 0xaf, 0xab, 0x82, 0xff, 0x71, 0xd5, 0xf7,
	0x7d, 0xb8, 0x40, 0xab, 0x63, 0x95, 0xda, 0xac, 0xc8, 0x17, 0xb9, 0x35, 0x9f, 0xf3, 0x57, 0xe0,
	0x2e, 0xfa, 0xb5, 0x40, 0x2b, 0xed, 0x16, 0xe9, 0xba, 0xcb, 0x68, 0x50, 0x5a, 0x63, 0x32, 0xf2,
	0x5f, 0x18, 0xf4, 0x77, 0x7e, 0x4a, 0x20, 0xd2, 0xdd, 0xbd, 0xe9, 0xd7, 0x6c, 0xfc, 0xec, 0x13,
	0x3d, 0xae, 0xf3, 0x17, 0x41, 0x09, 0x75, 0x11, 0x23, 0x7f, 0xc2, 0x60, 0xb0, 0x47, 0x8e, 0xc2,
	0x20, 0xf1, 0xd2, 0x9b, 0x7e, 0xdb, 0x48, 0xe8, 0x3e, 0xf5, 0x51, 0x6f, 0x17, 0x00, 0x46, 0xc8,
	0x1f, 0x57, 0x22, 0x48, 0x12, 0xa1, 0x30, 0x3c, 0x6d, 0x32, 0xf9, 0x3b, 0xec, 0xf3, 0x64, 0x81,
	0xfc, 0x29, 0x3b, 0xd4, 0x57, 0x92, 0x28, 0x47, 0xff, 0x59, 0xf3, 0xe9, 0xef, 0xef, 0x43, 0x50,
	0x35, 0xfe, 0x91, 0x8a, 0xbf, 0xd3, 0x06, 0xfe, 0x63, 0x15, 0x1f, 0xf3, 0x8f, 0xbe, 0xfc, 0xbb,
	0xed, 0xe1, 0x3f, 0x56, 0xf5, 0xfa, 0xdf, 0x4a, 0xc7, 0xff, 0x79, 0x1b, 0xf8, 0xff, 0x44, 0xc6,
	0xf5, 0xbe, 0x99, 0x60, 0x18, 0xb4, 0xa2, 0x6f, 0xce, 0x71, 0xf4, 0x77, 0x0f, 0xfa, 0x68, 0xf5,
	0x87, 0x59, 0xae, 0xbc, 0x52, 0xbc, 0x0c, 0xbd, 0xb5, 0x83, 0x78, 0x24, 0x13, 0xe0, 0x4c, 0xa5,
	0x46, 0xfc, 0xc4, 0x20, 0x20, 0x8f, 0xe6, 0x0b, 0x44, 0xd7, 0xe1, 0x74, 0xea, 0xf0, 0x98, 0x01,
	0x78, 0xcc, 0x6d, 0x91, 0x86, 0x73, 0x82, 0xee, 0xde, 0xc5, 0xcf, 0x3b, 0xec, 0x2d, 0x10, 0x05,
	0x42, 0x4d, 0x8a, 0xf0, 0x2b, 0x2b, 0xeb, 0xa8, 0x3d, 0x72, 0x10, 0x78, 0xf0, 0xaa, 0xca, 0x79,
	0x0b, 0x84, 0x80, 0x38, 0x27, 0x15, 0xd8, 0x73, 0xde, 0x1e, 0x09, 0x08, 0x3c, 0xf8, 0xa3, 0x3a,
	0x6f, 0x41, 0xf3, 0x27, 0xce, 0xa9, 0xf3, 0x57, 0xba, 0x61, 0xd3, 0xdb, 0xbe, 0xef, 0x86, 0x73,
	0x1c, 0xfd, 0xd1, 0x07, 0x8e, 0x56, 0xbf, 0x27, 0xcd, 0x22, 0x5f, 0x1e, 0x3a, 0xff, 0x13, 0x06,
	0x97, 0x2b, 0x69, 0x84, 0xca, 0x97, 0xe2, 0xd0, 0xe0, 0x59, 0x93, 0x43, 0x1a, 0xac, 0x0e, 0x31,
	0xc4, 0xc8, 0x7f, 0x67, 0xc0, 0x8f, 0x23, 0xa0, 0x77, 0x70, 0xa7, 0xf9, 0xef, 0xe0, 0x6e, 0x25,
	0x10, 0xf7, 0x18, 0x7e, 0x63, 0xf0, 0xfc, 0x71, 0x24, 0xbe, 0xfb, 0x9f, 0x34, 0x39, 0x1b, 0x17,
	0x95, 0x20, 0x12, 0x27, 0x02, 0x7f, 0xd6, 0xa3, 0x28, 0xc5, 0xe0, 0xb4, 0xf9, 0x09, 0xa9, 0xc4,
	0xe2, 0x25, 0xe1, 0x69, 0xbd, 0xb6, 0xbc, 0x34, 0x9c, 0x35, 0x39, 0x21, 0x95, 0xaa, 0x22, 0x81,
	0xb8, 0x95, 0x8f, 0x52, 0x28, 0x3a, 0xad, 0xca, 0x87, 0x97, 0x8b, 0x5b, 0xf9, 0xf0, 0xb2, 0xd1,
	0x6d, 0x49, 0x3e, 0x48, 0x3c, 0xfe, 0xa5, 0xe3, 0x26, 0x0d, 0x57, 0xbe, 0x6a, 0xc7, 0x9d, 0xe3,
	0xe8, 0xaf, 0x33, 0xe8, 0x57, 0xa7, 0x6f, 0xfc, 0x1e, 0x9c, 0x65, 0xf9, 0x52, 0x3d, 0x0a, 0xa7,
	0x34, 0x20, 0xf3, 0x0b, 0x67, 0xfd, 0x4c, 0xae, 0xb3, 0x65, 0xf8, 0x80, 0x86, 0x55, 0x7e, 0x41,
	0x53, 0xae, 0x4c, 0x2b, 0xb4, 0x52, 0x6f, 0xc2, 0xeb, 0xe1, 0x09, 0x4d, 0xb9, 0x76, 0x06, 0xfe,
	0x2a, 0x5c, 0xac, 0x25, 0x5a, 0x91, 0xae, 0x95, 0x34, 0xc2, 0xd9, 0xc3, 0x19, 0xf9, 0x0c, 0x9c,
	0xf9, 0x5d, 0x67, 0xfd, 0x38, 0xd3, 0x8a, 0x4f, 0xe0, 0x85, 0x83, 0xdf, 0x64, 0xe6, 0xe0, 0x7b,
	0xef, 0xd7, 0xc9, 0x9b, 0xef, 0xbd, 0x27, 0xb3, 0x8f, 0xb2, 0x9c, 0xb6, 0xbc, 0x06, 0xf7, 0x0e,
	0x5b, 0xa6, 0xd7, 0xe2, 0xa6, 0x3c, 0xff, 0x0d, 0xda, 0x71, 0xb9, 0xdf, 0x31, 0xbd, 0x7e, 0xdf,
	0xdf, 0xf1, 0x03, 0x83, 0x6e, 0xf9, 0xa7, 0x29, 0x7c, 0x93, 0xd2, 0xf3, 0x55, 0x23, 0xd2, 0x53,
	0x1b, 0xf3, 0x45, 0x3b, 0x90, 0xfc, 0x3b, 0x06, 0xa7, 0xae, 0xa2, 0xc2, 0xb7, 0x08, 0xed, 0x97,
	0x4d, 0x41, 0xbb, 0xff, 0xb5, 0x11, 0x11, 0x3c, 0xfe, 0x23, 0x83, 0xee, 0x4a, 0x51, 0xf1, 0x87,
	0x6f, 0x13, 0xd4, 0x6f, 0x9a, 0x02, 0xb5, 0xf6, 0xf3, 0x28, 0xea, 0xac, 0x94, 0x5b, 0x24, 0x1d,
	0x9a, 0x7d, 0x3f, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xed, 0x26, 0xa6, 0x11, 0x17, 0x00,
	0x00,
}
