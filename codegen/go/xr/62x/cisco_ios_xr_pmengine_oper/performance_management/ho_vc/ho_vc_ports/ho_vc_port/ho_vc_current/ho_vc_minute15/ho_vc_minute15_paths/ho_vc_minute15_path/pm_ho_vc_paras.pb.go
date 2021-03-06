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

package cisco_ios_xr_pmengine_oper_performance_management_ho_vc_ho_vc_ports_ho_vc_port_ho_vc_current_ho_vc_minute15_ho_vc_minute15_paths_ho_vc_minute15_path

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
	proto.RegisterType((*PmHoVcParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ho_vc.ho_vc_ports.ho_vc_port.ho_vc_current.ho_vc_minute15.ho_vc_minute15_paths.ho_vc_minute15_path.pm_ho_vc_paras_KEYS")
	proto.RegisterType((*PmHoVcParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ho_vc.ho_vc_ports.ho_vc_port.ho_vc_current.ho_vc_minute15.ho_vc_minute15_paths.ho_vc_minute15_path.pm_ho_vc_parameter")
	proto.RegisterType((*PmHoVcParameterRatio)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ho_vc.ho_vc_ports.ho_vc_port.ho_vc_current.ho_vc_minute15.ho_vc_minute15_paths.ho_vc_minute15_path.pm_ho_vc_parameter_ratio")
	proto.RegisterType((*HoVcEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ho_vc.ho_vc_ports.ho_vc_port.ho_vc_current.ho_vc_minute15.ho_vc_minute15_paths.ho_vc_minute15_path.ho_vcEntry")
	proto.RegisterType((*HoVcFarEndPathEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ho_vc.ho_vc_ports.ho_vc_port.ho_vc_current.ho_vc_minute15.ho_vc_minute15_paths.ho_vc_minute15_path.ho_vcFarEndPathEntry")
	proto.RegisterType((*PmHoVcParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ho_vc.ho_vc_ports.ho_vc_port.ho_vc_current.ho_vc_minute15.ho_vc_minute15_paths.ho_vc_minute15_path.pm_ho_vc_paras")
}

func init() { proto.RegisterFile("pm_ho_vc_paras.proto", fileDescriptor_7eb27aa255ef0874) }

var fileDescriptor_7eb27aa255ef0874 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0x7f, 0x4d, 0xd3, 0x7a, 0xaa, 0xf4, 0xa7, 0x2e, 0x01, 0xf9, 0x00, 0x22, 0xca, 0x01,
	0x45, 0x1c, 0x82, 0x92, 0x36, 0xfc, 0x39, 0xd2, 0x2a, 0x08, 0x09, 0x21, 0xa1, 0x35, 0x20, 0x71,
	0x5a, 0xad, 0x9d, 0x49, 0x63, 0x11, 0xaf, 0xad, 0xdd, 0x4d, 0x55, 0x9e, 0x82, 0x0b, 0x0f, 0xc0,
	0x11, 0x71, 0xe0, 0xc6, 0x0b, 0x20, 0x10, 0x27, 0xc4, 0x53, 0xf0, 0x1e, 0x68, 0xc7, 0x69, 0x9c,
	0xa4, 0x70, 0xe0, 0x16, 0xdf, 0x76, 0x3f, 0x7f, 0xf3, 0xcd, 0xce, 0xce, 0x8c, 0x67, 0xa1, 0x99,
	0xa7, 0x62, 0x92, 0x89, 0xb3, 0x58, 0xe4, 0x52, 0x4b, 0xd3, 0xcd, 0x75, 0x66, 0x33, 0xf6, 0xce,
	0x8b, 0x13, 0x13, 0x67, 0x22, 0xc9, 0x8c, 0x38, 0xd7, 0x22, 0x4f, 0x51, 0x9d, 0x26, 0x0a, 0x45,
	0x96, 0xa3, 0xee, 0xe6, 0xa8, 0xc7, 0x99, 0x4e, 0xa5, 0x8a, 0x51, 0xa4, 0x52, 0xc9, 0x53, 0x4c,
	0x51, 0xd9, 0x2e, 0xa9, 0x74, 0xe7, 0x5a, 0x99, 0xb6, 0x66, 0x69, 0x3d, 0x5f, 0xc6, 0x33, 0xad,
	0x17, 0x54, 0x91, 0x26, 0x6a, 0x66, 0xb1, 0x37, 0x58, 0xdb, 0x8a, 0x5c, 0xda, 0x89, 0xf9, 0x13,
	0xd8, 0x7e, 0x08, 0x57, 0x56, 0x8f, 0x2b, 0x9e, 0x0c, 0x5f, 0x85, 0x8c, 0x41, 0x4d, 0xc9, 0x14,
	0x03, 0xaf, 0xe5, 0x75, 0x7c, 0x4e, 0x6b, 0x76, 0x0d, 0xea, 0x6a, 0x96, 0x46, 0xa8, 0x83, 0xff,
	0x5a, 0x5e, 0xa7, 0xc1, 0xe7, 0xbb, 0x36, 0x02, 0x5b, 0x91, 0x48, 0xd1, 0xa2, 0x76, 0x0a, 0x23,
	0x69, 0x25, 0x29, 0x34, 0x38, 0xad, 0xd9, 0x75, 0xf0, 0xed, 0x44, 0xa3, 0x99, 0x64, 0xd3, 0xd1,
	0x5c, 0xa4, 0x04, 0xd8, 0x0d, 0x00, 0x1b, 0x4b, 0xa1, 0xd1, 0x85, 0x17, 0x6c, 0xb5, 0xbc, 0xce,
	0x2e, 0xf7, 0x6d, 0x2c, 0x39, 0x01, 0xed, 0xd7, 0x10, 0x5c, 0x76, 0x23, 0xb4, 0xb4, 0x49, 0xb6,
	0xe2, 0xcc, 0xff, 0x9b, 0x33, 0xff, 0x1f, 0x9c, 0xbd, 0x6d, 0x00, 0x90, 0xab, 0xa1, 0xb2, 0xfa,
	0x0d, 0xbb, 0x09, 0x7b, 0xee, 0xb6, 0x84, 0xb1, 0xd2, 0xce, 0x0c, 0xb9, 0x39, 0xe0, 0xe0, 0xa0,
	0x90, 0x10, 0xf6, 0xcd, 0x03, 0x9f, 0x18, 0x28, 0x8c, 0x21, 0x6f, 0x7b, 0xfd, 0x0f, 0x5e, 0x77,
	0x13, 0x53, 0xde, 0xbd, 0x7c, 0x8b, 0x7c, 0xc7, 0x7d, 0x18, 0x86, 0x86, 0xfd, 0xf0, 0x00, 0x8a,
	0x38, 0x8c, 0xd0, 0x86, 0xee, 0x65, 0xaf, 0xff, 0xa9, 0x32, 0x81, 0x14, 0xe5, 0xc0, 0x77, 0x29,
	0x1c, 0xc3, 0x0d, 0xfb, 0x7e, 0x11, 0x8f, 0xa1, 0xc4, 0xd4, 0xaa, 0x96, 0x18, 0x8a, 0x24, 0xc4,
	0xd0, 0xb0, 0x9f, 0xde, 0x45, 0x0d, 0x16, 0xa9, 0xd9, 0xae, 0x66, 0x6a, 0xfc, 0x22, 0x20, 0x97,
	0x9b, 0xa5, 0x9e, 0x89, 0x4c, 0x50, 0xaf, 0x66, 0xcf, 0x1c, 0x2f, 0xd5, 0xd8, 0x4c, 0xba, 0x1a,
	0xdb, 0xa9, 0x64, 0x8d, 0xbd, 0x90, 0xe1, 0x52, 0x24, 0x51, 0x24, 0xd0, 0x04, 0xbb, 0x95, 0x8c,
	0xe4, 0x38, 0x1a, 0x2e, 0x75, 0x4b, 0x14, 0xa1, 0xeb, 0x16, 0xbf, 0xc2, 0xdd, 0x72, 0x1c, 0x21,
	0x37, 0xed, 0x2f, 0x1e, 0x34, 0x89, 0xf4, 0x48, 0xea, 0xa1, 0x1a, 0x3d, 0x73, 0xc5, 0x47, 0xb3,
	0xa9, 0x03, 0x07, 0x63, 0xa9, 0x05, 0xaa, 0x91, 0x28, 0x27, 0x50, 0x31, 0x75, 0x1b, 0xe3, 0x92,
	0x1b, 0x1a, 0x76, 0x1b, 0xd8, 0x0a, 0xd3, 0x2c, 0x86, 0x55, 0x83, 0xef, 0x97, 0x54, 0xfa, 0xdd,
	0xac, 0xab, 0xc6, 0xe2, 0xac, 0x18, 0x07, 0x2b, 0xaa, 0x27, 0x2f, 0x2f, 0xab, 0x16, 0x5d, 0x50,
	0x5b, 0x57, 0x75, 0x05, 0xd6, 0xfe, 0x55, 0x83, 0xfd, 0xd5, 0xe7, 0x06, 0x6b, 0xc2, 0x76, 0xa2,
	0x46, 0x78, 0x1e, 0xf4, 0xc9, 0xa2, 0xd8, 0x38, 0xf4, 0x4c, 0x4e, 0x93, 0x51, 0x70, 0x48, 0x93,
	0xb9, 0xd8, 0xd0, 0x48, 0x4f, 0x52, 0x34, 0x56, 0xa6, 0x79, 0x70, 0xd4, 0xda, 0xa2, 0x91, 0x7e,
	0x01, 0xb0, 0x5b, 0xf0, 0xff, 0x54, 0x1a, 0x2b, 0xe2, 0x29, 0x4a, 0x2d, 0x1c, 0x1e, 0x0c, 0x88,
	0xd3, 0x70, 0xf0, 0x89, 0x43, 0x9f, 0x27, 0x29, 0xb2, 0x1e, 0x5c, 0x2d, 0x79, 0xbd, 0x81, 0xcb,
	0x45, 0xc1, 0xbe, 0x4b, 0x6c, 0xb6, 0x60, 0xf7, 0x06, 0x4f, 0x13, 0x45, 0x26, 0x77, 0xa0, 0x59,
	0x9a, 0xf4, 0x8f, 0xc4, 0x64, 0xae, 0x7f, 0x8f, 0x2c, 0x0e, 0x16, 0x16, 0xfd, 0xa3, 0xc7, 0x85,
	0x8f, 0xcf, 0x1e, 0xd4, 0xdc, 0x6d, 0x04, 0xf7, 0xa9, 0xf0, 0xde, 0x6f, 0x68, 0xe1, 0x95, 0x4f,
	0x1c, 0x4e, 0xc7, 0x65, 0x5f, 0x3d, 0xd8, 0x19, 0x23, 0x7d, 0x0c, 0x1e, 0xd0, 0xd1, 0x3f, 0x6e,
	0xf2, 0xd1, 0xd7, 0x7a, 0x81, 0xd7, 0xc7, 0xe8, 0x36, 0x51, 0x9d, 0xde, 0xdc, 0x87, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x8f, 0xaf, 0x4f, 0xe9, 0x8b, 0x0b, 0x00, 0x00,
}
