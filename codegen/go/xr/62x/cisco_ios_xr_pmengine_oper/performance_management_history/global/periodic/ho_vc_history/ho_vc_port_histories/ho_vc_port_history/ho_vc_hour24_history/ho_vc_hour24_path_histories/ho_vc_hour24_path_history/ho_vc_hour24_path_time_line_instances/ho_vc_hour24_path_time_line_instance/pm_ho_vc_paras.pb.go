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
// source: pm_ho_vc_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_history_global_periodic_ho_vc_history_ho_vc_port_histories_ho_vc_port_history_ho_vc_hour24_history_ho_vc_hour24_path_histories_ho_vc_hour24_path_history_ho_vc_hour24_path_time_line_instances_ho_vc_hour24_path_time_line_instance

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

type PmHoVcParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Number_1             uint32   `protobuf:"varint,3,opt,name=number_1,json=number1,proto3" json:"number_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmHoVcParas_KEYS) Reset()         { *m = PmHoVcParas_KEYS{} }
func (m *PmHoVcParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmHoVcParas_KEYS) ProtoMessage()    {}
func (*PmHoVcParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb27aa255ef0874, []int{0}
}

func (m *PmHoVcParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmHoVcParas_KEYS.Unmarshal(m, b)
}
func (m *PmHoVcParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmHoVcParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmHoVcParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmHoVcParas_KEYS.Merge(m, src)
}
func (m *PmHoVcParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmHoVcParas_KEYS.Size(m)
}
func (m *PmHoVcParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmHoVcParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmHoVcParas_KEYS proto.InternalMessageInfo

func (m *PmHoVcParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmHoVcParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *PmHoVcParas_KEYS) GetNumber_1() uint32 {
	if m != nil {
		return m.Number_1
	}
	return 0
}

type PmHoVcParameter struct {
	Data                 uint32   `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            uint32   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,3,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmHoVcParameter) Reset()         { *m = PmHoVcParameter{} }
func (m *PmHoVcParameter) String() string { return proto.CompactTextString(m) }
func (*PmHoVcParameter) ProtoMessage()    {}
func (*PmHoVcParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb27aa255ef0874, []int{1}
}

func (m *PmHoVcParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmHoVcParameter.Unmarshal(m, b)
}
func (m *PmHoVcParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmHoVcParameter.Marshal(b, m, deterministic)
}
func (m *PmHoVcParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmHoVcParameter.Merge(m, src)
}
func (m *PmHoVcParameter) XXX_Size() int {
	return xxx_messageInfo_PmHoVcParameter.Size(m)
}
func (m *PmHoVcParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PmHoVcParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PmHoVcParameter proto.InternalMessageInfo

func (m *PmHoVcParameter) GetData() uint32 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PmHoVcParameter) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PmHoVcParameter) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

type PmHoVcParameterRatio struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            string   `protobuf:"bytes,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,3,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmHoVcParameterRatio) Reset()         { *m = PmHoVcParameterRatio{} }
func (m *PmHoVcParameterRatio) String() string { return proto.CompactTextString(m) }
func (*PmHoVcParameterRatio) ProtoMessage()    {}
func (*PmHoVcParameterRatio) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb27aa255ef0874, []int{2}
}

func (m *PmHoVcParameterRatio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmHoVcParameterRatio.Unmarshal(m, b)
}
func (m *PmHoVcParameterRatio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmHoVcParameterRatio.Marshal(b, m, deterministic)
}
func (m *PmHoVcParameterRatio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmHoVcParameterRatio.Merge(m, src)
}
func (m *PmHoVcParameterRatio) XXX_Size() int {
	return xxx_messageInfo_PmHoVcParameterRatio.Size(m)
}
func (m *PmHoVcParameterRatio) XXX_DiscardUnknown() {
	xxx_messageInfo_PmHoVcParameterRatio.DiscardUnknown(m)
}

var xxx_messageInfo_PmHoVcParameterRatio proto.InternalMessageInfo

func (m *PmHoVcParameterRatio) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *PmHoVcParameterRatio) GetThreshold() string {
	if m != nil {
		return m.Threshold
	}
	return ""
}

func (m *PmHoVcParameterRatio) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

type HoVcEntry struct {
	PathStatus           int32                 `protobuf:"zigzag32,1,opt,name=path_status,json=pathStatus,proto3" json:"path_status,omitempty"`
	PathESs              *PmHoVcParameter      `protobuf:"bytes,2,opt,name=path_e_ss,json=pathESs,proto3" json:"path_e_ss,omitempty"`
	PathEsRs             *PmHoVcParameterRatio `protobuf:"bytes,3,opt,name=path_es_rs,json=pathEsRs,proto3" json:"path_es_rs,omitempty"`
	PathSeSs             *PmHoVcParameter      `protobuf:"bytes,4,opt,name=path_se_ss,json=pathSeSs,proto3" json:"path_se_ss,omitempty"`
	PathSesRs            *PmHoVcParameterRatio `protobuf:"bytes,5,opt,name=path_ses_rs,json=pathSesRs,proto3" json:"path_ses_rs,omitempty"`
	PathEBs              *PmHoVcParameter      `protobuf:"bytes,6,opt,name=path_e_bs,json=pathEBs,proto3" json:"path_e_bs,omitempty"`
	PathUaSs             *PmHoVcParameter      `protobuf:"bytes,7,opt,name=path_ua_ss,json=pathUaSs,proto3" json:"path_ua_ss,omitempty"`
	PathBbEs             *PmHoVcParameter      `protobuf:"bytes,8,opt,name=path_bb_es,json=pathBbEs,proto3" json:"path_bb_es,omitempty"`
	PathBbeRs            *PmHoVcParameterRatio `protobuf:"bytes,9,opt,name=path_bbe_rs,json=pathBbeRs,proto3" json:"path_bbe_rs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HoVcEntry) Reset()         { *m = HoVcEntry{} }
func (m *HoVcEntry) String() string { return proto.CompactTextString(m) }
func (*HoVcEntry) ProtoMessage()    {}
func (*HoVcEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb27aa255ef0874, []int{3}
}

func (m *HoVcEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HoVcEntry.Unmarshal(m, b)
}
func (m *HoVcEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HoVcEntry.Marshal(b, m, deterministic)
}
func (m *HoVcEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HoVcEntry.Merge(m, src)
}
func (m *HoVcEntry) XXX_Size() int {
	return xxx_messageInfo_HoVcEntry.Size(m)
}
func (m *HoVcEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_HoVcEntry.DiscardUnknown(m)
}

var xxx_messageInfo_HoVcEntry proto.InternalMessageInfo

func (m *HoVcEntry) GetPathStatus() int32 {
	if m != nil {
		return m.PathStatus
	}
	return 0
}

func (m *HoVcEntry) GetPathESs() *PmHoVcParameter {
	if m != nil {
		return m.PathESs
	}
	return nil
}

func (m *HoVcEntry) GetPathEsRs() *PmHoVcParameterRatio {
	if m != nil {
		return m.PathEsRs
	}
	return nil
}

func (m *HoVcEntry) GetPathSeSs() *PmHoVcParameter {
	if m != nil {
		return m.PathSeSs
	}
	return nil
}

func (m *HoVcEntry) GetPathSesRs() *PmHoVcParameterRatio {
	if m != nil {
		return m.PathSesRs
	}
	return nil
}

func (m *HoVcEntry) GetPathEBs() *PmHoVcParameter {
	if m != nil {
		return m.PathEBs
	}
	return nil
}

func (m *HoVcEntry) GetPathUaSs() *PmHoVcParameter {
	if m != nil {
		return m.PathUaSs
	}
	return nil
}

func (m *HoVcEntry) GetPathBbEs() *PmHoVcParameter {
	if m != nil {
		return m.PathBbEs
	}
	return nil
}

func (m *HoVcEntry) GetPathBbeRs() *PmHoVcParameterRatio {
	if m != nil {
		return m.PathBbeRs
	}
	return nil
}

type HoVcFarEndPathEntry struct {
	FarEndPathESs        uint32   `protobuf:"varint,1,opt,name=far_end_path_e_ss,json=farEndPathESs,proto3" json:"far_end_path_e_ss,omitempty"`
	FarEndPathSeSs       uint32   `protobuf:"varint,2,opt,name=far_end_path_se_ss,json=farEndPathSeSs,proto3" json:"far_end_path_se_ss,omitempty"`
	FarEndPathCVs        uint32   `protobuf:"varint,3,opt,name=far_end_path_c_vs,json=farEndPathCVs,proto3" json:"far_end_path_c_vs,omitempty"`
	FarEndPathUaSs       uint32   `protobuf:"varint,4,opt,name=far_end_path_ua_ss,json=farEndPathUaSs,proto3" json:"far_end_path_ua_ss,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HoVcFarEndPathEntry) Reset()         { *m = HoVcFarEndPathEntry{} }
func (m *HoVcFarEndPathEntry) String() string { return proto.CompactTextString(m) }
func (*HoVcFarEndPathEntry) ProtoMessage()    {}
func (*HoVcFarEndPathEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb27aa255ef0874, []int{4}
}

func (m *HoVcFarEndPathEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HoVcFarEndPathEntry.Unmarshal(m, b)
}
func (m *HoVcFarEndPathEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HoVcFarEndPathEntry.Marshal(b, m, deterministic)
}
func (m *HoVcFarEndPathEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HoVcFarEndPathEntry.Merge(m, src)
}
func (m *HoVcFarEndPathEntry) XXX_Size() int {
	return xxx_messageInfo_HoVcFarEndPathEntry.Size(m)
}
func (m *HoVcFarEndPathEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_HoVcFarEndPathEntry.DiscardUnknown(m)
}

var xxx_messageInfo_HoVcFarEndPathEntry proto.InternalMessageInfo

func (m *HoVcFarEndPathEntry) GetFarEndPathESs() uint32 {
	if m != nil {
		return m.FarEndPathESs
	}
	return 0
}

func (m *HoVcFarEndPathEntry) GetFarEndPathSeSs() uint32 {
	if m != nil {
		return m.FarEndPathSeSs
	}
	return 0
}

func (m *HoVcFarEndPathEntry) GetFarEndPathCVs() uint32 {
	if m != nil {
		return m.FarEndPathCVs
	}
	return 0
}

func (m *HoVcFarEndPathEntry) GetFarEndPathUaSs() uint32 {
	if m != nil {
		return m.FarEndPathUaSs
	}
	return 0
}

type PmHoVcParas struct {
	Index                uint32               `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool                 `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            []string             `protobuf:"bytes,52,rep,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        []string             `protobuf:"bytes,53,rep,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   []string             `protobuf:"bytes,54,rep,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear24HrTime    []string             `protobuf:"bytes,55,rep,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Path                 *HoVcEntry           `protobuf:"bytes,56,opt,name=path,proto3" json:"path,omitempty"`
	FePath               *HoVcFarEndPathEntry `protobuf:"bytes,57,opt,name=fe_path,json=fePath,proto3" json:"fe_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PmHoVcParas) Reset()         { *m = PmHoVcParas{} }
func (m *PmHoVcParas) String() string { return proto.CompactTextString(m) }
func (*PmHoVcParas) ProtoMessage()    {}
func (*PmHoVcParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb27aa255ef0874, []int{5}
}

func (m *PmHoVcParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmHoVcParas.Unmarshal(m, b)
}
func (m *PmHoVcParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmHoVcParas.Marshal(b, m, deterministic)
}
func (m *PmHoVcParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmHoVcParas.Merge(m, src)
}
func (m *PmHoVcParas) XXX_Size() int {
	return xxx_messageInfo_PmHoVcParas.Size(m)
}
func (m *PmHoVcParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmHoVcParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmHoVcParas proto.InternalMessageInfo

func (m *PmHoVcParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmHoVcParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmHoVcParas) GetTimestamp() []string {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PmHoVcParas) GetLastClearTime() []string {
	if m != nil {
		return m.LastClearTime
	}
	return nil
}

func (m *PmHoVcParas) GetLastClear15MinTime() []string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return nil
}

func (m *PmHoVcParas) GetLastClear24HrTime() []string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return nil
}

func (m *PmHoVcParas) GetPath() *HoVcEntry {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *PmHoVcParas) GetFePath() *HoVcFarEndPathEntry {
	if m != nil {
		return m.FePath
	}
	return nil
}

func init() {
	proto.RegisterType((*PmHoVcParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ho_vc_history.ho_vc_port_histories.ho_vc_port_history.ho_vc_hour24_history.ho_vc_hour24_path_histories.ho_vc_hour24_path_history.ho_vc_hour24_path_time_line_instances.ho_vc_hour24_path_time_line_instance.pm_ho_vc_paras_KEYS")
	proto.RegisterType((*PmHoVcParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ho_vc_history.ho_vc_port_histories.ho_vc_port_history.ho_vc_hour24_history.ho_vc_hour24_path_histories.ho_vc_hour24_path_history.ho_vc_hour24_path_time_line_instances.ho_vc_hour24_path_time_line_instance.pm_ho_vc_parameter")
	proto.RegisterType((*PmHoVcParameterRatio)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ho_vc_history.ho_vc_port_histories.ho_vc_port_history.ho_vc_hour24_history.ho_vc_hour24_path_histories.ho_vc_hour24_path_history.ho_vc_hour24_path_time_line_instances.ho_vc_hour24_path_time_line_instance.pm_ho_vc_parameter_ratio")
	proto.RegisterType((*HoVcEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ho_vc_history.ho_vc_port_histories.ho_vc_port_history.ho_vc_hour24_history.ho_vc_hour24_path_histories.ho_vc_hour24_path_history.ho_vc_hour24_path_time_line_instances.ho_vc_hour24_path_time_line_instance.ho_vcEntry")
	proto.RegisterType((*HoVcFarEndPathEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ho_vc_history.ho_vc_port_histories.ho_vc_port_history.ho_vc_hour24_history.ho_vc_hour24_path_histories.ho_vc_hour24_path_history.ho_vc_hour24_path_time_line_instances.ho_vc_hour24_path_time_line_instance.ho_vcFarEndPathEntry")
	proto.RegisterType((*PmHoVcParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.ho_vc_history.ho_vc_port_histories.ho_vc_port_history.ho_vc_hour24_history.ho_vc_hour24_path_histories.ho_vc_hour24_path_history.ho_vc_hour24_path_time_line_instances.ho_vc_hour24_path_time_line_instance.pm_ho_vc_paras")
}

func init() { proto.RegisterFile("pm_ho_vc_paras.proto", fileDescriptor_7eb27aa255ef0874) }

var fileDescriptor_7eb27aa255ef0874 = []byte{
	// 733 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0x27, 0xdb, 0x34, 0xc9, 0x4e, 0x49, 0xa5, 0x63, 0x2a, 0x2b, 0x28, 0x96, 0x1c, 0xa4, 0x78,
	0x88, 0x24, 0x6d, 0xfd, 0x73, 0x6d, 0x89, 0x08, 0x22, 0xc8, 0x46, 0x05, 0x41, 0x18, 0x66, 0x37,
	0x2f, 0xcd, 0x60, 0x76, 0x66, 0x99, 0x99, 0x94, 0xf6, 0x83, 0x88, 0x67, 0x8f, 0xe2, 0xc1, 0x9b,
	0x1f, 0x40, 0x50, 0x3c, 0x89, 0xdf, 0xc2, 0xaf, 0x21, 0xf3, 0x36, 0xe9, 0x26, 0x4d, 0x94, 0x36,
	0xe7, 0xbd, 0xed, 0xbc, 0xfd, 0xbd, 0xdf, 0x63, 0x7f, 0xfb, 0x7b, 0x2f, 0x6f, 0x43, 0x1a, 0x69,
	0xc2, 0x86, 0x8a, 0x9d, 0xc4, 0x2c, 0xe5, 0x9a, 0x9b, 0x56, 0xaa, 0x95, 0x55, 0xf4, 0xbd, 0x17,
	0x0b, 0x13, 0x2b, 0x26, 0x94, 0x61, 0xa7, 0x9a, 0xa5, 0x09, 0xc8, 0x63, 0x21, 0x81, 0xa9, 0x14,
	0x74, 0x2b, 0x05, 0x3d, 0x50, 0x3a, 0xe1, 0x32, 0x06, 0x96, 0x70, 0xc9, 0x8f, 0x21, 0x01, 0x69,
	0xd9, 0x50, 0x18, 0xab, 0xf4, 0x59, 0xeb, 0x78, 0xa4, 0x22, 0x3e, 0x72, 0x28, 0xa1, 0xfa, 0x22,
	0x6e, 0x65, 0xec, 0xd3, 0xbb, 0x93, 0x5a, 0x4a, 0x4f, 0x13, 0x04, 0x98, 0xc5, 0xe0, 0x14, 0x37,
	0x54, 0x63, 0xdd, 0xd9, 0x5f, 0x1e, 0x4c, 0xb9, 0x1d, 0x2e, 0x70, 0x2c, 0xde, 0x5b, 0x96, 0x65,
	0x45, 0x02, 0x6c, 0xe4, 0x1e, 0x48, 0x48, 0x63, 0xdd, 0x83, 0x98, 0x4b, 0xa1, 0x9a, 0x6f, 0xc9,
	0xf5, 0x79, 0xbd, 0xd8, 0xb3, 0xee, 0x9b, 0x1e, 0xa5, 0xa4, 0x2c, 0x79, 0x02, 0x41, 0x69, 0xa7,
	0xb4, 0xeb, 0x87, 0x78, 0x4d, 0x6f, 0x90, 0x8a, 0x1c, 0x27, 0x11, 0xe8, 0xc0, 0xdb, 0x29, 0xed,
	0xd6, 0xc3, 0xc9, 0x89, 0xde, 0x24, 0xb5, 0xec, 0x8a, 0xb5, 0x83, 0x35, 0xbc, 0x53, 0xcd, 0xce,
	0xed, 0x26, 0x10, 0x3a, 0xc7, 0x9e, 0x80, 0x05, 0xed, 0xc8, 0xfb, 0xdc, 0x72, 0x24, 0xaf, 0x87,
	0x78, 0x4d, 0x6f, 0x11, 0xdf, 0x0e, 0x35, 0x98, 0xa1, 0x1a, 0xf5, 0x27, 0xfc, 0x79, 0x80, 0xde,
	0x26, 0xc4, 0xc6, 0x9c, 0x69, 0x70, 0x8a, 0x62, 0x91, 0x5a, 0xe8, 0xdb, 0x98, 0x87, 0x18, 0x68,
	0xbe, 0x23, 0xc1, 0x62, 0x19, 0xa6, 0xb9, 0x15, 0x6a, 0xae, 0x98, 0xff, 0xaf, 0x62, 0xfe, 0x15,
	0x8a, 0x7d, 0xd8, 0x26, 0x04, 0x4b, 0x75, 0xa5, 0xd5, 0x67, 0xf4, 0x0e, 0xd9, 0x40, 0x6d, 0x8d,
	0xe5, 0x76, 0x6c, 0xb0, 0xcc, 0x56, 0x48, 0x5c, 0xa8, 0x87, 0x11, 0xfa, 0xc3, 0x23, 0x3e, 0x22,
	0x80, 0x19, 0x83, 0xd5, 0x36, 0x3a, 0x9f, 0xbc, 0x56, 0x61, 0xc7, 0x25, 0xa8, 0xd6, 0xe2, 0x6b,
	0x0c, 0xab, 0x0e, 0xdb, 0xed, 0x19, 0xfa, 0xcb, 0x23, 0x24, 0x13, 0xd2, 0x30, 0x6d, 0xf0, 0xc5,
	0x6c, 0x74, 0xbe, 0x14, 0x4a, 0x5e, 0x56, 0xc9, 0xac, 0x21, 0xc2, 0x1a, 0xea, 0x69, 0x42, 0x43,
	0x7f, 0x4e, 0x05, 0x35, 0x68, 0xcd, 0x72, 0x61, 0xcd, 0x2b, 0x5a, 0x13, 0xa5, 0xec, 0x41, 0xcf,
	0xd0, 0xdf, 0xde, 0x74, 0x0c, 0x64, 0xe6, 0x5c, 0x2f, 0xcc, 0xb9, 0x92, 0x39, 0xfd, 0x4c, 0x51,
	0xe7, 0xce, 0x99, 0xb9, 0x19, 0x99, 0xa0, 0x52, 0x98, 0x73, 0xa5, 0xb9, 0x79, 0x38, 0xd3, 0xe6,
	0x63, 0xee, 0xda, 0xbc, 0x5a, 0x28, 0xb9, 0x4a, 0x9b, 0xbf, 0xe2, 0xbd, 0x19, 0x29, 0xa3, 0x88,
	0x81, 0x09, 0x6a, 0x85, 0x94, 0xab, 0x48, 0x79, 0x18, 0x75, 0x67, 0x26, 0x66, 0x14, 0x81, 0x9b,
	0x98, 0x7e, 0x31, 0x31, 0x57, 0x9f, 0x98, 0x87, 0x11, 0x84, 0xa6, 0xf9, 0xad, 0x44, 0x1a, 0x08,
	0x7a, 0xc2, 0x75, 0x57, 0xf6, 0x5f, 0xb8, 0xfe, 0xc7, 0x1d, 0x75, 0x97, 0x6c, 0x0d, 0xb8, 0x66,
	0x20, 0xfb, 0x2c, 0xdf, 0x44, 0xb3, 0xed, 0xbb, 0x3e, 0xc8, 0xb1, 0x3d, 0x43, 0xef, 0x11, 0x3a,
	0x87, 0x34, 0xe7, 0x4b, 0x6b, 0x3d, 0xdc, 0xcc, 0xa1, 0xf8, 0x9b, 0x77, 0x91, 0x35, 0x66, 0x27,
	0x66, 0xf2, 0x01, 0x30, 0xc3, 0x7a, 0xf4, 0x7a, 0x91, 0x35, 0x1b, 0x44, 0xe5, 0x8b, 0xac, 0xae,
	0xc5, 0x9a, 0x7f, 0x2a, 0x64, 0x73, 0xfe, 0x8b, 0x84, 0x36, 0xc8, 0xba, 0x90, 0x7d, 0x38, 0x0d,
	0x3a, 0x98, 0x91, 0x1d, 0x5c, 0xf4, 0x84, 0x8f, 0x44, 0x3f, 0xd8, 0xc3, 0x0d, 0x3d, 0x3b, 0xe0,
	0x6a, 0x2f, 0x12, 0x30, 0x96, 0x27, 0x69, 0xb0, 0xbf, 0xb3, 0x86, 0xab, 0xfd, 0x34, 0x40, 0xef,
	0x92, 0x6b, 0x23, 0x6e, 0x2c, 0x8b, 0x47, 0xc0, 0x35, 0xca, 0x1c, 0x1c, 0x20, 0xa6, 0xee, 0xc2,
	0x47, 0x2e, 0xfa, 0x52, 0x24, 0x40, 0xdb, 0x64, 0x3b, 0xc7, 0xb5, 0x0f, 0x58, 0x22, 0x64, 0x86,
	0x7e, 0x80, 0x68, 0x7a, 0x8e, 0x6e, 0x1f, 0x3c, 0x17, 0x12, 0x53, 0xee, 0x93, 0x46, 0x9e, 0xe2,
	0x9c, 0x33, 0xe1, 0x7f, 0x88, 0x19, 0x5b, 0xe7, 0x19, 0x9d, 0xfd, 0xa7, 0x59, 0x8d, 0xaf, 0x1e,
	0x29, 0x3b, 0x35, 0x82, 0x47, 0xe8, 0xfc, 0x8f, 0x85, 0xf3, 0x97, 0x3b, 0x3f, 0xff, 0xd6, 0x0a,
	0x51, 0x2f, 0xfa, 0xdd, 0x23, 0xd5, 0x01, 0x20, 0x3e, 0x78, 0x8c, 0xda, 0x7d, 0x2e, 0xb4, 0xfb,
	0x8f, 0x76, 0x17, 0xa6, 0x41, 0x58, 0x19, 0x80, 0x3b, 0x44, 0x15, 0xfc, 0x67, 0x64, 0xef, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x04, 0x65, 0x50, 0x31, 0x11, 0x00, 0x00,
}
