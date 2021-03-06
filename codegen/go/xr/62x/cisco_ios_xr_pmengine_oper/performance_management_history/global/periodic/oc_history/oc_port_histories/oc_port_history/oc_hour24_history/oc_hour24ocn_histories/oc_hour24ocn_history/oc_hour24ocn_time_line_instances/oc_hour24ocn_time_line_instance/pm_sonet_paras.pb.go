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
// source: pm_sonet_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_history_global_periodic_oc_history_oc_port_histories_oc_port_history_oc_hour24_history_oc_hour24ocn_histories_oc_hour24ocn_history_oc_hour24ocn_time_line_instances_oc_hour24ocn_time_line_instance

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

type PmSonetParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Number_1             uint32   `protobuf:"varint,3,opt,name=number_1,json=number1,proto3" json:"number_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSonetParas_KEYS) Reset()         { *m = PmSonetParas_KEYS{} }
func (m *PmSonetParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmSonetParas_KEYS) ProtoMessage()    {}
func (*PmSonetParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af52c194751c16c, []int{0}
}

func (m *PmSonetParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSonetParas_KEYS.Unmarshal(m, b)
}
func (m *PmSonetParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSonetParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmSonetParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSonetParas_KEYS.Merge(m, src)
}
func (m *PmSonetParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmSonetParas_KEYS.Size(m)
}
func (m *PmSonetParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSonetParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmSonetParas_KEYS proto.InternalMessageInfo

func (m *PmSonetParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmSonetParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *PmSonetParas_KEYS) GetNumber_1() uint32 {
	if m != nil {
		return m.Number_1
	}
	return 0
}

type PmSonetParameter struct {
	Data                 uint32   `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            uint32   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,3,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSonetParameter) Reset()         { *m = PmSonetParameter{} }
func (m *PmSonetParameter) String() string { return proto.CompactTextString(m) }
func (*PmSonetParameter) ProtoMessage()    {}
func (*PmSonetParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af52c194751c16c, []int{1}
}

func (m *PmSonetParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSonetParameter.Unmarshal(m, b)
}
func (m *PmSonetParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSonetParameter.Marshal(b, m, deterministic)
}
func (m *PmSonetParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSonetParameter.Merge(m, src)
}
func (m *PmSonetParameter) XXX_Size() int {
	return xxx_messageInfo_PmSonetParameter.Size(m)
}
func (m *PmSonetParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSonetParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PmSonetParameter proto.InternalMessageInfo

func (m *PmSonetParameter) GetData() uint32 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PmSonetParameter) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PmSonetParameter) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

type SonetSectionEntry struct {
	SectionStatus        int32             `protobuf:"zigzag32,1,opt,name=section_status,json=sectionStatus,proto3" json:"section_status,omitempty"`
	SectionESs           *PmSonetParameter `protobuf:"bytes,2,opt,name=section_e_ss,json=sectionESs,proto3" json:"section_e_ss,omitempty"`
	SectionSeSs          *PmSonetParameter `protobuf:"bytes,3,opt,name=section_se_ss,json=sectionSeSs,proto3" json:"section_se_ss,omitempty"`
	SectionSefSs         *PmSonetParameter `protobuf:"bytes,4,opt,name=section_sef_ss,json=sectionSefSs,proto3" json:"section_sef_ss,omitempty"`
	SectionCVs           *PmSonetParameter `protobuf:"bytes,5,opt,name=section_c_vs,json=sectionCVs,proto3" json:"section_c_vs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SonetSectionEntry) Reset()         { *m = SonetSectionEntry{} }
func (m *SonetSectionEntry) String() string { return proto.CompactTextString(m) }
func (*SonetSectionEntry) ProtoMessage()    {}
func (*SonetSectionEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af52c194751c16c, []int{2}
}

func (m *SonetSectionEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SonetSectionEntry.Unmarshal(m, b)
}
func (m *SonetSectionEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SonetSectionEntry.Marshal(b, m, deterministic)
}
func (m *SonetSectionEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SonetSectionEntry.Merge(m, src)
}
func (m *SonetSectionEntry) XXX_Size() int {
	return xxx_messageInfo_SonetSectionEntry.Size(m)
}
func (m *SonetSectionEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SonetSectionEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SonetSectionEntry proto.InternalMessageInfo

func (m *SonetSectionEntry) GetSectionStatus() int32 {
	if m != nil {
		return m.SectionStatus
	}
	return 0
}

func (m *SonetSectionEntry) GetSectionESs() *PmSonetParameter {
	if m != nil {
		return m.SectionESs
	}
	return nil
}

func (m *SonetSectionEntry) GetSectionSeSs() *PmSonetParameter {
	if m != nil {
		return m.SectionSeSs
	}
	return nil
}

func (m *SonetSectionEntry) GetSectionSefSs() *PmSonetParameter {
	if m != nil {
		return m.SectionSefSs
	}
	return nil
}

func (m *SonetSectionEntry) GetSectionCVs() *PmSonetParameter {
	if m != nil {
		return m.SectionCVs
	}
	return nil
}

type SonetLineEntry struct {
	LineStatus           int32             `protobuf:"zigzag32,1,opt,name=line_status,json=lineStatus,proto3" json:"line_status,omitempty"`
	LineESs              *PmSonetParameter `protobuf:"bytes,2,opt,name=line_e_ss,json=lineESs,proto3" json:"line_e_ss,omitempty"`
	LineSeSs             *PmSonetParameter `protobuf:"bytes,3,opt,name=line_se_ss,json=lineSeSs,proto3" json:"line_se_ss,omitempty"`
	LineCVs              *PmSonetParameter `protobuf:"bytes,4,opt,name=line_c_vs,json=lineCVs,proto3" json:"line_c_vs,omitempty"`
	LineUaSs             *PmSonetParameter `protobuf:"bytes,5,opt,name=line_ua_ss,json=lineUaSs,proto3" json:"line_ua_ss,omitempty"`
	LineFcLs             *PmSonetParameter `protobuf:"bytes,6,opt,name=line_fc_ls,json=lineFcLs,proto3" json:"line_fc_ls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SonetLineEntry) Reset()         { *m = SonetLineEntry{} }
func (m *SonetLineEntry) String() string { return proto.CompactTextString(m) }
func (*SonetLineEntry) ProtoMessage()    {}
func (*SonetLineEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af52c194751c16c, []int{3}
}

func (m *SonetLineEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SonetLineEntry.Unmarshal(m, b)
}
func (m *SonetLineEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SonetLineEntry.Marshal(b, m, deterministic)
}
func (m *SonetLineEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SonetLineEntry.Merge(m, src)
}
func (m *SonetLineEntry) XXX_Size() int {
	return xxx_messageInfo_SonetLineEntry.Size(m)
}
func (m *SonetLineEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SonetLineEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SonetLineEntry proto.InternalMessageInfo

func (m *SonetLineEntry) GetLineStatus() int32 {
	if m != nil {
		return m.LineStatus
	}
	return 0
}

func (m *SonetLineEntry) GetLineESs() *PmSonetParameter {
	if m != nil {
		return m.LineESs
	}
	return nil
}

func (m *SonetLineEntry) GetLineSeSs() *PmSonetParameter {
	if m != nil {
		return m.LineSeSs
	}
	return nil
}

func (m *SonetLineEntry) GetLineCVs() *PmSonetParameter {
	if m != nil {
		return m.LineCVs
	}
	return nil
}

func (m *SonetLineEntry) GetLineUaSs() *PmSonetParameter {
	if m != nil {
		return m.LineUaSs
	}
	return nil
}

func (m *SonetLineEntry) GetLineFcLs() *PmSonetParameter {
	if m != nil {
		return m.LineFcLs
	}
	return nil
}

type SonetFarEndLineEntry struct {
	FarEndLineESs        *PmSonetParameter `protobuf:"bytes,1,opt,name=far_end_line_e_ss,json=farEndLineESs,proto3" json:"far_end_line_e_ss,omitempty"`
	FarEndLineSeSs       *PmSonetParameter `protobuf:"bytes,2,opt,name=far_end_line_se_ss,json=farEndLineSeSs,proto3" json:"far_end_line_se_ss,omitempty"`
	FarEndLineCVs        *PmSonetParameter `protobuf:"bytes,3,opt,name=far_end_line_c_vs,json=farEndLineCVs,proto3" json:"far_end_line_c_vs,omitempty"`
	FarEndLineUaSs       *PmSonetParameter `protobuf:"bytes,4,opt,name=far_end_line_ua_ss,json=farEndLineUaSs,proto3" json:"far_end_line_ua_ss,omitempty"`
	FarEndLineFcLs       *PmSonetParameter `protobuf:"bytes,5,opt,name=far_end_line_fc_ls,json=farEndLineFcLs,proto3" json:"far_end_line_fc_ls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SonetFarEndLineEntry) Reset()         { *m = SonetFarEndLineEntry{} }
func (m *SonetFarEndLineEntry) String() string { return proto.CompactTextString(m) }
func (*SonetFarEndLineEntry) ProtoMessage()    {}
func (*SonetFarEndLineEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af52c194751c16c, []int{4}
}

func (m *SonetFarEndLineEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SonetFarEndLineEntry.Unmarshal(m, b)
}
func (m *SonetFarEndLineEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SonetFarEndLineEntry.Marshal(b, m, deterministic)
}
func (m *SonetFarEndLineEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SonetFarEndLineEntry.Merge(m, src)
}
func (m *SonetFarEndLineEntry) XXX_Size() int {
	return xxx_messageInfo_SonetFarEndLineEntry.Size(m)
}
func (m *SonetFarEndLineEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SonetFarEndLineEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SonetFarEndLineEntry proto.InternalMessageInfo

func (m *SonetFarEndLineEntry) GetFarEndLineESs() *PmSonetParameter {
	if m != nil {
		return m.FarEndLineESs
	}
	return nil
}

func (m *SonetFarEndLineEntry) GetFarEndLineSeSs() *PmSonetParameter {
	if m != nil {
		return m.FarEndLineSeSs
	}
	return nil
}

func (m *SonetFarEndLineEntry) GetFarEndLineCVs() *PmSonetParameter {
	if m != nil {
		return m.FarEndLineCVs
	}
	return nil
}

func (m *SonetFarEndLineEntry) GetFarEndLineUaSs() *PmSonetParameter {
	if m != nil {
		return m.FarEndLineUaSs
	}
	return nil
}

func (m *SonetFarEndLineEntry) GetFarEndLineFcLs() *PmSonetParameter {
	if m != nil {
		return m.FarEndLineFcLs
	}
	return nil
}

type PmSonetParas struct {
	Index                uint32                `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool                  `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            []string              `protobuf:"bytes,52,rep,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        []string              `protobuf:"bytes,53,rep,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   []string              `protobuf:"bytes,54,rep,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear24HrTime    []string              `protobuf:"bytes,55,rep,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Section              *SonetSectionEntry    `protobuf:"bytes,56,opt,name=section,proto3" json:"section,omitempty"`
	Line                 *SonetLineEntry       `protobuf:"bytes,57,opt,name=line,proto3" json:"line,omitempty"`
	FeLine               *SonetFarEndLineEntry `protobuf:"bytes,58,opt,name=fe_line,json=feLine,proto3" json:"fe_line,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PmSonetParas) Reset()         { *m = PmSonetParas{} }
func (m *PmSonetParas) String() string { return proto.CompactTextString(m) }
func (*PmSonetParas) ProtoMessage()    {}
func (*PmSonetParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af52c194751c16c, []int{5}
}

func (m *PmSonetParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSonetParas.Unmarshal(m, b)
}
func (m *PmSonetParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSonetParas.Marshal(b, m, deterministic)
}
func (m *PmSonetParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSonetParas.Merge(m, src)
}
func (m *PmSonetParas) XXX_Size() int {
	return xxx_messageInfo_PmSonetParas.Size(m)
}
func (m *PmSonetParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSonetParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmSonetParas proto.InternalMessageInfo

func (m *PmSonetParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmSonetParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmSonetParas) GetTimestamp() []string {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PmSonetParas) GetLastClearTime() []string {
	if m != nil {
		return m.LastClearTime
	}
	return nil
}

func (m *PmSonetParas) GetLastClear15MinTime() []string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return nil
}

func (m *PmSonetParas) GetLastClear24HrTime() []string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return nil
}

func (m *PmSonetParas) GetSection() *SonetSectionEntry {
	if m != nil {
		return m.Section
	}
	return nil
}

func (m *PmSonetParas) GetLine() *SonetLineEntry {
	if m != nil {
		return m.Line
	}
	return nil
}

func (m *PmSonetParas) GetFeLine() *SonetFarEndLineEntry {
	if m != nil {
		return m.FeLine
	}
	return nil
}

func init() {
	proto.RegisterType((*PmSonetParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.oc_history.oc_port_histories.oc_port_history.oc_hour24_history.oc_hour24ocn_histories.oc_hour24ocn_history.oc_hour24ocn_time_line_instances.oc_hour24ocn_time_line_instance.pm_sonet_paras_KEYS")
	proto.RegisterType((*PmSonetParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.oc_history.oc_port_histories.oc_port_history.oc_hour24_history.oc_hour24ocn_histories.oc_hour24ocn_history.oc_hour24ocn_time_line_instances.oc_hour24ocn_time_line_instance.pm_sonet_parameter")
	proto.RegisterType((*SonetSectionEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.oc_history.oc_port_histories.oc_port_history.oc_hour24_history.oc_hour24ocn_histories.oc_hour24ocn_history.oc_hour24ocn_time_line_instances.oc_hour24ocn_time_line_instance.sonetSectionEntry")
	proto.RegisterType((*SonetLineEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.oc_history.oc_port_histories.oc_port_history.oc_hour24_history.oc_hour24ocn_histories.oc_hour24ocn_history.oc_hour24ocn_time_line_instances.oc_hour24ocn_time_line_instance.sonetLineEntry")
	proto.RegisterType((*SonetFarEndLineEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.oc_history.oc_port_histories.oc_port_history.oc_hour24_history.oc_hour24ocn_histories.oc_hour24ocn_history.oc_hour24ocn_time_line_instances.oc_hour24ocn_time_line_instance.sonetFarEndLineEntry")
	proto.RegisterType((*PmSonetParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.oc_history.oc_port_histories.oc_port_history.oc_hour24_history.oc_hour24ocn_histories.oc_hour24ocn_history.oc_hour24ocn_time_line_instances.oc_hour24ocn_time_line_instance.pm_sonet_paras")
}

func init() { proto.RegisterFile("pm_sonet_paras.proto", fileDescriptor_4af52c194751c16c) }

var fileDescriptor_4af52c194751c16c = []byte{
	// 783 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x98, 0x4d, 0x6b, 0x1b, 0x39,
	0x18, 0xc7, 0xf1, 0x24, 0xf1, 0xcb, 0x13, 0xdb, 0x8b, 0xb5, 0xde, 0x65, 0x16, 0x76, 0xd9, 0x60,
	0xd8, 0x25, 0x27, 0x2f, 0x76, 0x92, 0x7d, 0xbb, 0x06, 0x87, 0x85, 0x4d, 0x2f, 0xe3, 0x36, 0x50,
	0x28, 0x08, 0x65, 0xac, 0x89, 0x05, 0x33, 0xd2, 0x20, 0xc9, 0x21, 0xb9, 0xb7, 0xd0, 0xd2, 0x6f,
	0xd0, 0x6b, 0x69, 0xe9, 0xa1, 0xb4, 0x3d, 0x94, 0x7e, 0x8f, 0x7e, 0x98, 0xde, 0x7a, 0x28, 0x7a,
	0xc6, 0xf6, 0xc4, 0x2f, 0x4d, 0x3f, 0x80, 0xe7, 0x36, 0x7a, 0xf4, 0x97, 0x34, 0xf9, 0xe5, 0xaf,
	0xe7, 0x19, 0x3f, 0xd0, 0x4e, 0x13, 0x6a, 0x94, 0xe4, 0x96, 0xa6, 0x4c, 0x33, 0xd3, 0x4d, 0xb5,
	0xb2, 0x8a, 0x7c, 0x2a, 0x85, 0xc2, 0x84, 0x8a, 0x0a, 0x65, 0xe8, 0x95, 0xa6, 0x69, 0xc2, 0xe5,
	0x85, 0x90, 0x9c, 0xaa, 0x94, 0xeb, 0x6e, 0xca, 0x75, 0xa4, 0x74, 0xc2, 0x64, 0xc8, 0x69, 0xc2,
	0x24, 0xbb, 0xe0, 0x09, 0x97, 0x96, 0x8e, 0x85, 0xb1, 0x4a, 0x5f, 0x77, 0x2f, 0x62, 0x75, 0xce,
	0x62, 0xa7, 0x12, 0x6a, 0x24, 0xc2, 0xae, 0x0a, 0xe7, 0x53, 0x2a, 0xa4, 0xa9, 0xd2, 0x33, 0xa9,
	0xe0, 0x66, 0x29, 0x82, 0x8a, 0xb1, 0x9a, 0xe8, 0xfe, 0xe1, 0x6a, 0x44, 0x85, 0x72, 0x71, 0xe1,
	0x72, 0x78, 0x49, 0x6b, 0x45, 0xc2, 0x69, 0xec, 0xde, 0x5a, 0x48, 0x63, 0xdd, 0xdb, 0x9a, 0x6f,
	0x09, 0x3a, 0x0f, 0xe0, 0xfb, 0x45, 0x1e, 0xf4, 0xff, 0xc1, 0xfd, 0x21, 0x21, 0xb0, 0x2d, 0x59,
	0xc2, 0xfd, 0xd2, 0x5e, 0x69, 0xbf, 0x16, 0xe0, 0x33, 0xf9, 0x11, 0xca, 0x72, 0x92, 0x9c, 0x73,
	0xed, 0x7b, 0x7b, 0xa5, 0xfd, 0x46, 0x30, 0x1d, 0x91, 0x9f, 0xa0, 0x9a, 0x3d, 0xd1, 0x9e, 0xbf,
	0x85, 0x33, 0x95, 0x6c, 0xdc, 0xeb, 0x70, 0x20, 0x0b, 0xbb, 0x27, 0xdc, 0x72, 0xed, 0x36, 0x1f,
	0x31, 0xcb, 0x70, 0xf3, 0x46, 0x80, 0xcf, 0xe4, 0x67, 0xa8, 0xd9, 0xb1, 0xe6, 0x66, 0xac, 0xe2,
	0xd1, 0x74, 0xff, 0x3c, 0x40, 0x7e, 0x01, 0xb0, 0x21, 0xa3, 0x9a, 0x3b, 0x74, 0x78, 0x48, 0x35,
	0xa8, 0xd9, 0x90, 0x05, 0x18, 0xe8, 0x7c, 0xae, 0x41, 0x0b, 0x0f, 0x19, 0xf2, 0xd0, 0x0a, 0x25,
	0x07, 0xd2, 0xea, 0x6b, 0xf2, 0x1b, 0x34, 0x4d, 0x36, 0xa6, 0xc6, 0x32, 0x3b, 0x31, 0x78, 0x60,
	0x2b, 0x68, 0x4c, 0xa3, 0x43, 0x0c, 0x92, 0xd7, 0x1e, 0xd4, 0x67, 0x3a, 0x4e, 0x8d, 0xc1, 0xd3,
	0x77, 0xfb, 0x4f, 0xbc, 0xee, 0x46, 0x3a, 0xa2, 0xbb, 0xfa, 0x0f, 0x0b, 0x60, 0xca, 0x67, 0x30,
	0x34, 0xe4, 0x8d, 0x07, 0x8d, 0x39, 0x56, 0xe4, 0xb5, 0x55, 0xf0, 0x5a, 0xe2, 0xb5, 0x3b, 0x73,
	0x18, 0x1f, 0x1a, 0xf2, 0xce, 0xbb, 0xe1, 0x43, 0x1e, 0x39, 0x62, 0xdb, 0x05, 0xb1, 0x25, 0x62,
	0xf5, 0x39, 0xb1, 0x68, 0xb8, 0x78, 0x25, 0x43, 0x7a, 0x69, 0xfc, 0x9d, 0x02, 0xd8, 0x57, 0xae,
	0xe4, 0xf1, 0x99, 0xe9, 0x3c, 0xac, 0x43, 0x13, 0xe7, 0x4f, 0x85, 0xe4, 0x59, 0xee, 0xfb, 0x15,
	0x76, 0x71, 0x8b, 0x85, 0xc4, 0x07, 0x2e, 0x34, 0xcd, 0x7a, 0x2f, 0x3d, 0xa8, 0xa1, 0xa2, 0x48,
	0x79, 0x6b, 0xf9, 0x56, 0x9c, 0xc2, 0xe5, 0xbb, 0x57, 0x1e, 0x40, 0x86, 0xb2, 0x48, 0x76, 0x6b,
	0x49, 0x55, 0xd1, 0x55, 0x2e, 0xd3, 0xcd, 0x3d, 0x85, 0x77, 0xb6, 0x48, 0x72, 0x6b, 0x3d, 0x75,
	0x7c, 0x76, 0xc3, 0x53, 0x13, 0xe6, 0x3c, 0x55, 0x64, 0xb7, 0xb5, 0x9e, 0xba, 0xc7, 0x6e, 0x5e,
	0xbf, 0x28, 0xa4, 0xb1, 0xf1, 0xcb, 0x05, 0xaa, 0x75, 0xa8, 0x4e, 0xc2, 0x53, 0xd3, 0x79, 0x5b,
	0x87, 0x36, 0xce, 0x9e, 0x30, 0x3d, 0x90, 0xa3, 0xbc, 0x18, 0xbc, 0xf7, 0xa0, 0x15, 0x31, 0x4d,
	0xb9, 0x1c, 0xd1, 0x3c, 0xe7, 0x97, 0x0a, 0x94, 0x4b, 0x28, 0x1b, 0x51, 0x0e, 0x6e, 0x68, 0xc8,
	0x07, 0x0f, 0xc8, 0x02, 0x36, 0x53, 0xd4, 0xca, 0xb5, 0xdc, 0x9a, 0x39, 0x37, 0xac, 0x03, 0x2b,
	0x7e, 0xc3, 0x7a, 0x50, 0x54, 0xce, 0x5b, 0xfc, 0xe6, 0xaa, 0xc2, 0x8a, 0xdf, 0xb2, 0xea, 0x50,
	0xd4, 0xd1, 0x5b, 0xfc, 0x86, 0x35, 0x62, 0x05, 0x5c, 0x56, 0x2b, 0x8a, 0xb2, 0x7a, 0x0b, 0x38,
	0xac, 0x18, 0x1f, 0xab, 0xd0, 0x5c, 0xec, 0xfe, 0x90, 0x36, 0xec, 0x08, 0x39, 0xe2, 0x57, 0x7e,
	0x1f, 0x7b, 0x30, 0xd9, 0xc0, 0x45, 0x2f, 0x59, 0x2c, 0x46, 0xfe, 0x01, 0xb6, 0x5e, 0xb2, 0x01,
	0xf6, 0x6c, 0x44, 0xc2, 0x8d, 0x65, 0x49, 0xea, 0x1f, 0xee, 0x6d, 0xed, 0xd7, 0x82, 0x3c, 0x40,
	0x7e, 0x87, 0xef, 0x62, 0x66, 0x2c, 0x0d, 0x63, 0xce, 0x34, 0xbe, 0xa0, 0x7f, 0x84, 0x9a, 0x86,
	0x0b, 0x1f, 0xbb, 0xe8, 0x5d, 0x91, 0x70, 0xd2, 0x83, 0x1f, 0x72, 0x5d, 0xef, 0x88, 0x26, 0x22,
	0xfb, 0x73, 0xfc, 0x3f, 0x51, 0x4d, 0xe6, 0xea, 0xde, 0xd1, 0x1d, 0x21, 0x71, 0xc9, 0x1f, 0xd0,
	0xce, 0x97, 0x38, 0xbe, 0xd3, 0xfd, 0xff, 0xc2, 0x15, 0xad, 0xf9, 0x8a, 0xfe, 0xe1, 0x7f, 0xd9,
	0x19, 0xcf, 0x3d, 0xa8, 0x4c, 0x7f, 0x30, 0xf9, 0x7f, 0xa3, 0x2d, 0x1e, 0x6f, 0xaa, 0x2d, 0x56,
	0xfa, 0x64, 0xc1, 0x8c, 0x0c, 0x79, 0xe6, 0xc1, 0xb6, 0x13, 0xfb, 0xff, 0x20, 0xa2, 0x47, 0x1b,
	0x8d, 0x68, 0xfe, 0xf9, 0x14, 0x20, 0x13, 0xf2, 0xc2, 0x83, 0x4a, 0x94, 0x29, 0xfd, 0x7f, 0x91,
	0xcf, 0xd3, 0x8d, 0xe6, 0xb3, 0xf4, 0x91, 0x19, 0x94, 0x23, 0xee, 0x06, 0xe7, 0x65, 0xec, 0xa7,
	0x1f, 0x7c, 0x09, 0x00, 0x00, 0xff, 0xff, 0x80, 0x37, 0x40, 0xee, 0x67, 0x17, 0x00, 0x00,
}
