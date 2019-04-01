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
// source: pm_sonet_path_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_history_global_periodic_sonet_history_sonet_port_histories_sonet_port_history_sonet_minute15_history_sonet_minute15_path_histories_sonet_minute15_path_history_sonet_minute15_path_time_line_instances_sonet_minute15_path_time_line_instance

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

type PmSonetPathParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Number_1             uint32   `protobuf:"varint,3,opt,name=number_1,json=number1,proto3" json:"number_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSonetPathParas_KEYS) Reset()         { *m = PmSonetPathParas_KEYS{} }
func (m *PmSonetPathParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmSonetPathParas_KEYS) ProtoMessage()    {}
func (*PmSonetPathParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e901a9804fd43c, []int{0}
}

func (m *PmSonetPathParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSonetPathParas_KEYS.Unmarshal(m, b)
}
func (m *PmSonetPathParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSonetPathParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmSonetPathParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSonetPathParas_KEYS.Merge(m, src)
}
func (m *PmSonetPathParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmSonetPathParas_KEYS.Size(m)
}
func (m *PmSonetPathParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSonetPathParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmSonetPathParas_KEYS proto.InternalMessageInfo

func (m *PmSonetPathParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmSonetPathParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *PmSonetPathParas_KEYS) GetNumber_1() uint32 {
	if m != nil {
		return m.Number_1
	}
	return 0
}

type PmSonetPathParameter struct {
	Data                 uint32   `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            uint32   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,3,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSonetPathParameter) Reset()         { *m = PmSonetPathParameter{} }
func (m *PmSonetPathParameter) String() string { return proto.CompactTextString(m) }
func (*PmSonetPathParameter) ProtoMessage()    {}
func (*PmSonetPathParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e901a9804fd43c, []int{1}
}

func (m *PmSonetPathParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSonetPathParameter.Unmarshal(m, b)
}
func (m *PmSonetPathParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSonetPathParameter.Marshal(b, m, deterministic)
}
func (m *PmSonetPathParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSonetPathParameter.Merge(m, src)
}
func (m *PmSonetPathParameter) XXX_Size() int {
	return xxx_messageInfo_PmSonetPathParameter.Size(m)
}
func (m *PmSonetPathParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSonetPathParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PmSonetPathParameter proto.InternalMessageInfo

func (m *PmSonetPathParameter) GetData() uint32 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PmSonetPathParameter) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PmSonetPathParameter) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

type SonetPathEntry struct {
	PathWidth            string                `protobuf:"bytes,1,opt,name=path_width,json=pathWidth,proto3" json:"path_width,omitempty"`
	PathStatus           int32                 `protobuf:"zigzag32,2,opt,name=path_status,json=pathStatus,proto3" json:"path_status,omitempty"`
	PathESs              *PmSonetPathParameter `protobuf:"bytes,3,opt,name=path_e_ss,json=pathESs,proto3" json:"path_e_ss,omitempty"`
	PathSeSs             *PmSonetPathParameter `protobuf:"bytes,4,opt,name=path_se_ss,json=pathSeSs,proto3" json:"path_se_ss,omitempty"`
	PathCVs              *PmSonetPathParameter `protobuf:"bytes,5,opt,name=path_c_vs,json=pathCVs,proto3" json:"path_c_vs,omitempty"`
	PathUaSs             *PmSonetPathParameter `protobuf:"bytes,6,opt,name=path_ua_ss,json=pathUaSs,proto3" json:"path_ua_ss,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SonetPathEntry) Reset()         { *m = SonetPathEntry{} }
func (m *SonetPathEntry) String() string { return proto.CompactTextString(m) }
func (*SonetPathEntry) ProtoMessage()    {}
func (*SonetPathEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e901a9804fd43c, []int{2}
}

func (m *SonetPathEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SonetPathEntry.Unmarshal(m, b)
}
func (m *SonetPathEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SonetPathEntry.Marshal(b, m, deterministic)
}
func (m *SonetPathEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SonetPathEntry.Merge(m, src)
}
func (m *SonetPathEntry) XXX_Size() int {
	return xxx_messageInfo_SonetPathEntry.Size(m)
}
func (m *SonetPathEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SonetPathEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SonetPathEntry proto.InternalMessageInfo

func (m *SonetPathEntry) GetPathWidth() string {
	if m != nil {
		return m.PathWidth
	}
	return ""
}

func (m *SonetPathEntry) GetPathStatus() int32 {
	if m != nil {
		return m.PathStatus
	}
	return 0
}

func (m *SonetPathEntry) GetPathESs() *PmSonetPathParameter {
	if m != nil {
		return m.PathESs
	}
	return nil
}

func (m *SonetPathEntry) GetPathSeSs() *PmSonetPathParameter {
	if m != nil {
		return m.PathSeSs
	}
	return nil
}

func (m *SonetPathEntry) GetPathCVs() *PmSonetPathParameter {
	if m != nil {
		return m.PathCVs
	}
	return nil
}

func (m *SonetPathEntry) GetPathUaSs() *PmSonetPathParameter {
	if m != nil {
		return m.PathUaSs
	}
	return nil
}

type SonetFarEndPathEntry struct {
	FarEndPathESs        uint32   `protobuf:"varint,1,opt,name=far_end_path_e_ss,json=farEndPathESs,proto3" json:"far_end_path_e_ss,omitempty"`
	FarEndPathSeSs       uint32   `protobuf:"varint,2,opt,name=far_end_path_se_ss,json=farEndPathSeSs,proto3" json:"far_end_path_se_ss,omitempty"`
	FarEndPathCVs        uint32   `protobuf:"varint,3,opt,name=far_end_path_c_vs,json=farEndPathCVs,proto3" json:"far_end_path_c_vs,omitempty"`
	FarEndPathUaSs       uint32   `protobuf:"varint,4,opt,name=far_end_path_ua_ss,json=farEndPathUaSs,proto3" json:"far_end_path_ua_ss,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SonetFarEndPathEntry) Reset()         { *m = SonetFarEndPathEntry{} }
func (m *SonetFarEndPathEntry) String() string { return proto.CompactTextString(m) }
func (*SonetFarEndPathEntry) ProtoMessage()    {}
func (*SonetFarEndPathEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e901a9804fd43c, []int{3}
}

func (m *SonetFarEndPathEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SonetFarEndPathEntry.Unmarshal(m, b)
}
func (m *SonetFarEndPathEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SonetFarEndPathEntry.Marshal(b, m, deterministic)
}
func (m *SonetFarEndPathEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SonetFarEndPathEntry.Merge(m, src)
}
func (m *SonetFarEndPathEntry) XXX_Size() int {
	return xxx_messageInfo_SonetFarEndPathEntry.Size(m)
}
func (m *SonetFarEndPathEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SonetFarEndPathEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SonetFarEndPathEntry proto.InternalMessageInfo

func (m *SonetFarEndPathEntry) GetFarEndPathESs() uint32 {
	if m != nil {
		return m.FarEndPathESs
	}
	return 0
}

func (m *SonetFarEndPathEntry) GetFarEndPathSeSs() uint32 {
	if m != nil {
		return m.FarEndPathSeSs
	}
	return 0
}

func (m *SonetFarEndPathEntry) GetFarEndPathCVs() uint32 {
	if m != nil {
		return m.FarEndPathCVs
	}
	return 0
}

func (m *SonetFarEndPathEntry) GetFarEndPathUaSs() uint32 {
	if m != nil {
		return m.FarEndPathUaSs
	}
	return 0
}

type PmSonetPathParas struct {
	Index                uint32                `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool                  `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            []string              `protobuf:"bytes,52,rep,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        []string              `protobuf:"bytes,53,rep,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   []string              `protobuf:"bytes,54,rep,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear24HrTime    []string              `protobuf:"bytes,55,rep,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Path                 *SonetPathEntry       `protobuf:"bytes,56,opt,name=path,proto3" json:"path,omitempty"`
	FePath               *SonetFarEndPathEntry `protobuf:"bytes,57,opt,name=fe_path,json=fePath,proto3" json:"fe_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PmSonetPathParas) Reset()         { *m = PmSonetPathParas{} }
func (m *PmSonetPathParas) String() string { return proto.CompactTextString(m) }
func (*PmSonetPathParas) ProtoMessage()    {}
func (*PmSonetPathParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e901a9804fd43c, []int{4}
}

func (m *PmSonetPathParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSonetPathParas.Unmarshal(m, b)
}
func (m *PmSonetPathParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSonetPathParas.Marshal(b, m, deterministic)
}
func (m *PmSonetPathParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSonetPathParas.Merge(m, src)
}
func (m *PmSonetPathParas) XXX_Size() int {
	return xxx_messageInfo_PmSonetPathParas.Size(m)
}
func (m *PmSonetPathParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSonetPathParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmSonetPathParas proto.InternalMessageInfo

func (m *PmSonetPathParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmSonetPathParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmSonetPathParas) GetTimestamp() []string {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PmSonetPathParas) GetLastClearTime() []string {
	if m != nil {
		return m.LastClearTime
	}
	return nil
}

func (m *PmSonetPathParas) GetLastClear15MinTime() []string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return nil
}

func (m *PmSonetPathParas) GetLastClear24HrTime() []string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return nil
}

func (m *PmSonetPathParas) GetPath() *SonetPathEntry {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *PmSonetPathParas) GetFePath() *SonetFarEndPathEntry {
	if m != nil {
		return m.FePath
	}
	return nil
}

func init() {
	proto.RegisterType((*PmSonetPathParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.sonet_history.sonet_port_histories.sonet_port_history.sonet_minute15_history.sonet_minute15_path_histories.sonet_minute15_path_history.sonet_minute15_path_time_line_instances.sonet_minute15_path_time_line_instance.pm_sonet_path_paras_KEYS")
	proto.RegisterType((*PmSonetPathParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.sonet_history.sonet_port_histories.sonet_port_history.sonet_minute15_history.sonet_minute15_path_histories.sonet_minute15_path_history.sonet_minute15_path_time_line_instances.sonet_minute15_path_time_line_instance.pm_sonetPath_parameter")
	proto.RegisterType((*SonetPathEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.sonet_history.sonet_port_histories.sonet_port_history.sonet_minute15_history.sonet_minute15_path_histories.sonet_minute15_path_history.sonet_minute15_path_time_line_instances.sonet_minute15_path_time_line_instance.sonetPathEntry")
	proto.RegisterType((*SonetFarEndPathEntry)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.sonet_history.sonet_port_histories.sonet_port_history.sonet_minute15_history.sonet_minute15_path_histories.sonet_minute15_path_history.sonet_minute15_path_time_line_instances.sonet_minute15_path_time_line_instance.sonetFarEndPathEntry")
	proto.RegisterType((*PmSonetPathParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management_history.global.periodic.sonet_history.sonet_port_histories.sonet_port_history.sonet_minute15_history.sonet_minute15_path_histories.sonet_minute15_path_history.sonet_minute15_path_time_line_instances.sonet_minute15_path_time_line_instance.pm_sonet_path_paras")
}

func init() { proto.RegisterFile("pm_sonet_path_paras.proto", fileDescriptor_84e901a9804fd43c) }

var fileDescriptor_84e901a9804fd43c = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0xb6, 0x69, 0xd2, 0xb8, 0x4a, 0x51, 0x4d, 0xa9, 0xb6, 0x12, 0x88, 0x2a, 0x07, 0x54,
	0x71, 0x08, 0x4a, 0xda, 0xf0, 0x71, 0xae, 0x82, 0x90, 0x10, 0x12, 0xda, 0xf0, 0x21, 0x4e, 0x96,
	0xbb, 0x3b, 0x69, 0x2c, 0xad, 0xed, 0x95, 0xed, 0x94, 0xf6, 0xce, 0xff, 0xe0, 0x27, 0xc0, 0x05,
	0x71, 0x47, 0xdc, 0xb9, 0x20, 0x7e, 0x0f, 0xf2, 0x6c, 0x92, 0x6d, 0x97, 0x00, 0x95, 0x72, 0xdd,
	0xdb, 0xfa, 0xf9, 0xcd, 0xbc, 0xcc, 0xf3, 0xd8, 0x13, 0xb2, 0x97, 0x49, 0x66, 0xb5, 0x02, 0xc7,
	0x32, 0xee, 0x26, 0x2c, 0xe3, 0x86, 0xdb, 0x6e, 0x66, 0xb4, 0xd3, 0xf4, 0x63, 0x10, 0x0b, 0x1b,
	0x6b, 0x26, 0xb4, 0x65, 0xe7, 0x86, 0x65, 0x12, 0xd4, 0xa9, 0x50, 0xc0, 0x74, 0x06, 0xa6, 0x9b,
	0x81, 0x19, 0x6b, 0x23, 0xb9, 0x8a, 0x81, 0x49, 0xae, 0xf8, 0x29, 0x48, 0x50, 0x8e, 0x4d, 0x84,
	0x75, 0xda, 0x5c, 0x74, 0x4f, 0x53, 0x7d, 0xc2, 0x53, 0xcf, 0x12, 0x3a, 0x11, 0x71, 0x37, 0x97,
	0x98, 0xef, 0xce, 0x04, 0xb5, 0x99, 0x43, 0x02, 0xec, 0x9f, 0xe0, 0x9c, 0x27, 0x85, 0x9a, 0x3a,
	0xe8, 0x0d, 0xfe, 0x06, 0xe3, 0x0f, 0x2f, 0xe7, 0x59, 0xb6, 0xbb, 0x3c, 0xd2, 0x09, 0x09, 0x2c,
	0xf5, 0x85, 0x09, 0x65, 0x9d, 0x2f, 0xc8, 0x5e, 0x93, 0xd7, 0xe1, 0x24, 0x5c, 0x62, 0x1f, 0x7b,
	0x3e, 0x7c, 0x37, 0xa2, 0x94, 0xd4, 0x15, 0x97, 0x10, 0xd6, 0xf6, 0x6b, 0x07, 0xad, 0x08, 0xbf,
	0xe9, 0x2e, 0x69, 0xa8, 0xa9, 0x3c, 0x01, 0x13, 0x06, 0xfb, 0xb5, 0x83, 0x76, 0x34, 0x5b, 0xd1,
	0x3d, 0xb2, 0x91, 0x7f, 0xb1, 0x5e, 0xb8, 0x86, 0x3b, 0xcd, 0x7c, 0xdd, 0xeb, 0x08, 0xb2, 0x3b,
	0x97, 0x78, 0x39, 0x57, 0x90, 0xe0, 0xc0, 0x78, 0x81, 0x84, 0x3b, 0x8e, 0x02, 0xed, 0x08, 0xbf,
	0xe9, 0x6d, 0xd2, 0x72, 0x13, 0x03, 0x76, 0xa2, 0xd3, 0x64, 0xa6, 0x51, 0x00, 0xf4, 0x0e, 0x21,
	0x2e, 0xe6, 0xcc, 0x80, 0x37, 0x19, 0x85, 0x36, 0xa2, 0x96, 0x8b, 0x79, 0x84, 0x40, 0xe7, 0xf3,
	0x26, 0xd9, 0x5a, 0x08, 0x0d, 0x95, 0x33, 0x17, 0x3e, 0x02, 0xeb, 0x7a, 0x2f, 0x12, 0x37, 0x99,
	0x95, 0xd2, 0xf2, 0xc8, 0x5b, 0x0f, 0xd0, 0xbb, 0x64, 0x13, 0xb7, 0xad, 0xe3, 0x6e, 0x6a, 0x51,
	0x70, 0x3b, 0xc2, 0x88, 0x11, 0x22, 0xf4, 0x67, 0x40, 0x90, 0xce, 0x80, 0x59, 0x8b, 0x8a, 0x9b,
	0xfd, 0xaf, 0x41, 0xb7, 0xea, 0xab, 0x7f, 0xf2, 0xba, 0xcb, 0x4f, 0x3c, 0x6a, 0x7a, 0xfe, 0x70,
	0x64, 0xe9, 0xaf, 0x60, 0x76, 0x2e, 0x16, 0x7d, 0xad, 0x57, 0xbe, 0xae, 0xe4, 0xeb, 0x06, 0x76,
	0x2c, 0x8c, 0x2e, 0xf5, 0x6b, 0xcc, 0xce, 0x6c, 0xb8, 0x5e, 0xf9, 0xba, 0x7a, 0xbf, 0x1e, 0xbf,
	0xb9, 0xd4, 0xaf, 0x53, 0xee, 0xfb, 0xb5, 0x51, 0xf9, 0xba, 0x7a, 0xbf, 0xbe, 0xe6, 0x23, 0xdb,
	0xf9, 0x56, 0x23, 0x3b, 0xc8, 0x78, 0xca, 0xcd, 0x50, 0x25, 0xc5, 0xc3, 0x7d, 0x40, 0xb6, 0xc7,
	0xdc, 0x30, 0x50, 0x09, 0x2b, 0xde, 0xdf, 0x7c, 0x52, 0xb4, 0xc7, 0x05, 0x77, 0x64, 0xe9, 0x7d,
	0x42, 0xaf, 0x30, 0xf3, 0x27, 0x25, 0x9f, 0x1d, 0x5b, 0x05, 0x15, 0xaf, 0x47, 0x39, 0x2b, 0xde,
	0x92, 0xb5, 0x72, 0x56, 0x7f, 0xe2, 0xe5, 0xac, 0xf9, 0xc1, 0xd7, 0xcb, 0x59, 0xb1, 0x88, 0x0f,
	0x4d, 0x72, 0x73, 0xc9, 0x18, 0xa5, 0x3b, 0x64, 0x5d, 0xa8, 0x04, 0xce, 0xc3, 0x3e, 0x86, 0xe5,
	0x0b, 0x8f, 0x9e, 0xf1, 0x54, 0x24, 0xe1, 0x21, 0xce, 0xaf, 0x7c, 0x81, 0x83, 0x4f, 0x48, 0xb0,
	0x8e, 0xcb, 0x2c, 0x3c, 0xda, 0x5f, 0xf3, 0x73, 0x6a, 0x01, 0xd0, 0x7b, 0xe4, 0x46, 0xca, 0xad,
	0x63, 0x71, 0x0a, 0xdc, 0xa0, 0xd1, 0xe1, 0x00, 0x39, 0x6d, 0x0f, 0x1f, 0x7b, 0xf4, 0x95, 0x90,
	0x40, 0x7b, 0xe4, 0x56, 0xc1, 0xeb, 0x0d, 0xfc, 0x01, 0xe5, 0xec, 0x87, 0xc8, 0xa6, 0x0b, 0x76,
	0x6f, 0xf0, 0x42, 0x28, 0x0c, 0x79, 0x40, 0x76, 0x8a, 0x90, 0xfe, 0x11, 0x9b, 0xcc, 0xf2, 0x3f,
	0xc2, 0x88, 0xed, 0x45, 0x44, 0xff, 0xe8, 0x59, 0xae, 0xf1, 0x3d, 0x20, 0x75, 0x5f, 0x64, 0xf8,
	0x18, 0x6f, 0xc1, 0xa7, 0xea, 0x16, 0xfc, 0xef, 0x16, 0x5c, 0xfd, 0x4f, 0x12, 0xa1, 0x7b, 0xf4,
	0x47, 0x40, 0x9a, 0x63, 0xc0, 0x98, 0xf0, 0x09, 0x3a, 0xf9, 0xa5, 0x72, 0xf2, 0x5a, 0x4e, 0x96,
	0x9e, 0x8a, 0xa8, 0x31, 0x06, 0xbf, 0x38, 0x69, 0xe0, 0xbf, 0xfe, 0xc3, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x38, 0x63, 0x0a, 0x16, 0x12, 0x0c, 0x00, 0x00,
}
