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
// source: pm_pcs_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_ethernet_ethernet_ports_ethernet_port_ethernet_current_ethernet_hour24_hour24pcses_hour24pcs

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

type PmPcsParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmPcsParas_KEYS) Reset()         { *m = PmPcsParas_KEYS{} }
func (m *PmPcsParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmPcsParas_KEYS) ProtoMessage()    {}
func (*PmPcsParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ee74ada23b448, []int{0}
}

func (m *PmPcsParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmPcsParas_KEYS.Unmarshal(m, b)
}
func (m *PmPcsParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmPcsParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmPcsParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmPcsParas_KEYS.Merge(m, src)
}
func (m *PmPcsParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmPcsParas_KEYS.Size(m)
}
func (m *PmPcsParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmPcsParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmPcsParas_KEYS proto.InternalMessageInfo

func (m *PmPcsParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmPcsParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PmPcsParameter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data                 uint64   `protobuf:"varint,2,opt,name=data,proto3" json:"data,omitempty"`
	Threshold            uint64   `protobuf:"varint,3,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TcaReport            bool     `protobuf:"varint,4,opt,name=tca_report,json=tcaReport,proto3" json:"tca_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmPcsParameter) Reset()         { *m = PmPcsParameter{} }
func (m *PmPcsParameter) String() string { return proto.CompactTextString(m) }
func (*PmPcsParameter) ProtoMessage()    {}
func (*PmPcsParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ee74ada23b448, []int{1}
}

func (m *PmPcsParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmPcsParameter.Unmarshal(m, b)
}
func (m *PmPcsParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmPcsParameter.Marshal(b, m, deterministic)
}
func (m *PmPcsParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmPcsParameter.Merge(m, src)
}
func (m *PmPcsParameter) XXX_Size() int {
	return xxx_messageInfo_PmPcsParameter.Size(m)
}
func (m *PmPcsParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PmPcsParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PmPcsParameter proto.InternalMessageInfo

func (m *PmPcsParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmPcsParameter) GetData() uint64 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *PmPcsParameter) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PmPcsParameter) GetTcaReport() bool {
	if m != nil {
		return m.TcaReport
	}
	return false
}

type PmPcsParas struct {
	Index                uint32            `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool              `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            string            `protobuf:"bytes,52,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        string            `protobuf:"bytes,53,opt,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   string            `protobuf:"bytes,54,opt,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear30SecTime   string            `protobuf:"bytes,55,opt,name=last_clear30_sec_time,json=lastClear30SecTime,proto3" json:"last_clear30_sec_time,omitempty"`
	LastClear24HrTime    string            `protobuf:"bytes,56,opt,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Bip                  []*PmPcsParameter `protobuf:"bytes,57,rep,name=bip,proto3" json:"bip,omitempty"`
	FrmErr               []*PmPcsParameter `protobuf:"bytes,58,rep,name=frm_err,json=frmErr,proto3" json:"frm_err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PmPcsParas) Reset()         { *m = PmPcsParas{} }
func (m *PmPcsParas) String() string { return proto.CompactTextString(m) }
func (*PmPcsParas) ProtoMessage()    {}
func (*PmPcsParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ee74ada23b448, []int{2}
}

func (m *PmPcsParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmPcsParas.Unmarshal(m, b)
}
func (m *PmPcsParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmPcsParas.Marshal(b, m, deterministic)
}
func (m *PmPcsParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmPcsParas.Merge(m, src)
}
func (m *PmPcsParas) XXX_Size() int {
	return xxx_messageInfo_PmPcsParas.Size(m)
}
func (m *PmPcsParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmPcsParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmPcsParas proto.InternalMessageInfo

func (m *PmPcsParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmPcsParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmPcsParas) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *PmPcsParas) GetLastClearTime() string {
	if m != nil {
		return m.LastClearTime
	}
	return ""
}

func (m *PmPcsParas) GetLastClear15MinTime() string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return ""
}

func (m *PmPcsParas) GetLastClear30SecTime() string {
	if m != nil {
		return m.LastClear30SecTime
	}
	return ""
}

func (m *PmPcsParas) GetLastClear24HrTime() string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return ""
}

func (m *PmPcsParas) GetBip() []*PmPcsParameter {
	if m != nil {
		return m.Bip
	}
	return nil
}

func (m *PmPcsParas) GetFrmErr() []*PmPcsParameter {
	if m != nil {
		return m.FrmErr
	}
	return nil
}

func init() {
	proto.RegisterType((*PmPcsParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ethernet.ethernet_ports.ethernet_port.ethernet_current.ethernet_hour24.hour24pcses.hour24pcs.pm_pcs_paras_KEYS")
	proto.RegisterType((*PmPcsParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ethernet.ethernet_ports.ethernet_port.ethernet_current.ethernet_hour24.hour24pcses.hour24pcs.pm_pcs_parameter")
	proto.RegisterType((*PmPcsParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management.ethernet.ethernet_ports.ethernet_port.ethernet_current.ethernet_hour24.hour24pcses.hour24pcs.pm_pcs_paras")
}

func init() { proto.RegisterFile("pm_pcs_paras.proto", fileDescriptor_610ee74ada23b448) }

var fileDescriptor_610ee74ada23b448 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x65, 0xda, 0x29, 0xd4, 0x30, 0x82, 0xb1, 0x06, 0xe4, 0x05, 0x48, 0x55, 0x17, 0xa8,
	0xab, 0x30, 0x4d, 0x5b, 0xfe, 0x36, 0x2c, 0xd0, 0x48, 0x48, 0x88, 0x4d, 0x86, 0x0d, 0xab, 0x2b,
	0xd7, 0xbd, 0x9d, 0x58, 0xaa, 0x7f, 0x74, 0xed, 0xc2, 0x3c, 0x03, 0x0b, 0x9e, 0x81, 0x77, 0x60,
	0xc7, 0xd3, 0xa1, 0x38, 0xa3, 0x49, 0x28, 0x3c, 0x00, 0xac, 0x72, 0xcf, 0xc9, 0xf9, 0x9c, 0x13,
	0xe9, 0x9a, 0x8b, 0x60, 0x21, 0xe8, 0x08, 0x41, 0x91, 0x8a, 0x45, 0x20, 0x9f, 0xbc, 0xf8, 0xc6,
	0xb4, 0x89, 0xda, 0x83, 0xf1, 0x11, 0xae, 0x08, 0x82, 0x45, 0x77, 0x69, 0x1c, 0x82, 0x0f, 0x48,
	0x45, 0x40, 0xda, 0x7a, 0xb2, 0xca, 0x69, 0x04, 0xab, 0x9c, 0xba, 0x44, 0x8b, 0x2e, 0x15, 0x98,
	0x6a, 0x24, 0x87, 0xdd, 0x00, 0xc1, 0x53, 0x8a, 0xbf, 0xcb, 0x4e, 0xe9, 0x3d, 0x51, 0x1f, 0x83,
	0xda, 0xef, 0xa9, 0x5c, 0x16, 0xed, 0x23, 0xe8, 0x88, 0xb1, 0x9b, 0xa7, 0x6f, 0xf8, 0x49, 0xbf,
	0x26, 0xbc, 0x3f, 0xff, 0x74, 0x21, 0x04, 0x1f, 0x3a, 0x65, 0x51, 0xb2, 0x09, 0x9b, 0x8d, 0xab,
	0x3c, 0x8b, 0x47, 0x7c, 0xe4, 0xf6, 0x76, 0x8d, 0x24, 0x6f, 0x4d, 0xd8, 0xec, 0xb8, 0xba, 0x56,
	0xd3, 0x2f, 0xfc, 0x41, 0xef, 0x00, 0x8b, 0x09, 0xe9, 0xaf, 0xbc, 0xe0, 0xc3, 0x8d, 0x4a, 0x2a,
	0xd3, 0xc3, 0x2a, 0xcf, 0xe2, 0x31, 0x1f, 0xa7, 0x9a, 0x30, 0xd6, 0x7e, 0xb7, 0x91, 0x83, 0xfc,
	0xa2, 0x33, 0xc4, 0x13, 0xce, 0x93, 0x56, 0x40, 0xd8, 0xfc, 0x9f, 0x1c, 0x4e, 0xd8, 0xec, 0x4e,
	0x35, 0x4e, 0x5a, 0x55, 0xd9, 0x98, 0x7e, 0x3d, 0xe2, 0xf7, 0xfa, 0xd5, 0xc5, 0x29, 0x3f, 0x32,
	0x6e, 0x83, 0x57, 0xb2, 0xcc, 0x05, 0x5b, 0xd1, 0xb8, 0x9f, 0xd5, 0xce, 0x6c, 0xe4, 0x22, 0x1f,
	0xd0, 0x8a, 0xfc, 0x65, 0x63, 0x31, 0x26, 0x65, 0x83, 0x5c, 0xe6, 0x9a, 0x9d, 0x21, 0x9e, 0xf2,
	0xfb, 0x3b, 0x15, 0x13, 0xe8, 0x1d, 0x2a, 0x82, 0xc6, 0x97, 0xab, 0x9c, 0x39, 0x6e, 0xec, 0xb7,
	0x8d, 0xfb, 0xd1, 0x58, 0x14, 0x73, 0xfe, 0xb0, 0xcb, 0xcd, 0x57, 0x60, 0x8d, 0x6b, 0xd3, 0xcf,
	0x73, 0x5a, 0xdc, 0xa4, 0xe7, 0xab, 0x0f, 0xc6, 0xfd, 0x89, 0x2c, 0xce, 0x20, 0xa2, 0x6e, 0x91,
	0x17, 0x07, 0xc8, 0xe2, 0xec, 0x02, 0x75, 0x46, 0x9e, 0xf1, 0xd3, 0x0e, 0x29, 0x97, 0x50, 0x5f,
	0x57, 0x7a, 0x99, 0x89, 0x93, 0x1b, 0xa2, 0x5c, 0xbe, 0x6b, 0x6b, 0xfd, 0x60, 0x7c, 0xb0, 0x36,
	0x41, 0xbe, 0x9a, 0x0c, 0x66, 0x77, 0xcb, 0xef, 0xac, 0xf8, 0xb7, 0x76, 0xae, 0x38, 0xdc, 0x97,
	0xaa, 0x69, 0x2b, 0x7e, 0x32, 0x7e, 0x7b, 0x4b, 0x16, 0x90, 0x48, 0xbe, 0xfe, 0x5f, 0x9a, 0x8f,
	0xb6, 0x64, 0xcf, 0x89, 0xd6, 0xa3, 0x7c, 0xbd, 0x17, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x14,
	0x17, 0xd7, 0xde, 0xf4, 0x03, 0x00, 0x00,
}
