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
// source: pm_prbs_paras.proto

package cisco_ios_xr_pmengine_oper_performance_management_odu_odu_ports_odu_port_odu_current_odu_minute15_odu_minute15prbses_odu_minute15prbs

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

type PmPrbsParas_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmPrbsParas_KEYS) Reset()         { *m = PmPrbsParas_KEYS{} }
func (m *PmPrbsParas_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmPrbsParas_KEYS) ProtoMessage()    {}
func (*PmPrbsParas_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2609214b0ba90b4, []int{0}
}

func (m *PmPrbsParas_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmPrbsParas_KEYS.Unmarshal(m, b)
}
func (m *PmPrbsParas_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmPrbsParas_KEYS.Marshal(b, m, deterministic)
}
func (m *PmPrbsParas_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmPrbsParas_KEYS.Merge(m, src)
}
func (m *PmPrbsParas_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmPrbsParas_KEYS.Size(m)
}
func (m *PmPrbsParas_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmPrbsParas_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmPrbsParas_KEYS proto.InternalMessageInfo

func (m *PmPrbsParas_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmPrbsParas_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PmPrbsParameter struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	RcvPatt              string   `protobuf:"bytes,2,opt,name=rcv_patt,json=rcvPatt,proto3" json:"rcv_patt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmPrbsParameter) Reset()         { *m = PmPrbsParameter{} }
func (m *PmPrbsParameter) String() string { return proto.CompactTextString(m) }
func (*PmPrbsParameter) ProtoMessage()    {}
func (*PmPrbsParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2609214b0ba90b4, []int{1}
}

func (m *PmPrbsParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmPrbsParameter.Unmarshal(m, b)
}
func (m *PmPrbsParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmPrbsParameter.Marshal(b, m, deterministic)
}
func (m *PmPrbsParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmPrbsParameter.Merge(m, src)
}
func (m *PmPrbsParameter) XXX_Size() int {
	return xxx_messageInfo_PmPrbsParameter.Size(m)
}
func (m *PmPrbsParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PmPrbsParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PmPrbsParameter proto.InternalMessageInfo

func (m *PmPrbsParameter) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmPrbsParameter) GetRcvPatt() string {
	if m != nil {
		return m.RcvPatt
	}
	return ""
}

type PmPrbsParas struct {
	Index                uint32           `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool             `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            string           `protobuf:"bytes,52,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        string           `protobuf:"bytes,53,opt,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   string           `protobuf:"bytes,54,opt,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear24HrTime    string           `protobuf:"bytes,55,opt,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Ebc                  uint64           `protobuf:"varint,56,opt,name=ebc,proto3" json:"ebc,omitempty"`
	FoundCount           uint32           `protobuf:"varint,57,opt,name=found_count,json=foundCount,proto3" json:"found_count,omitempty"`
	LostCount            uint32           `protobuf:"varint,58,opt,name=lost_count,json=lostCount,proto3" json:"lost_count,omitempty"`
	FoundAtTime          string           `protobuf:"bytes,59,opt,name=found_at_time,json=foundAtTime,proto3" json:"found_at_time,omitempty"`
	LostAtTime           string           `protobuf:"bytes,60,opt,name=lost_at_time,json=lostAtTime,proto3" json:"lost_at_time,omitempty"`
	ConfPatt             string           `protobuf:"bytes,61,opt,name=conf_patt,json=confPatt,proto3" json:"conf_patt,omitempty"`
	RcvPatt              *PmPrbsParameter `protobuf:"bytes,62,opt,name=rcv_patt,json=rcvPatt,proto3" json:"rcv_patt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PmPrbsParas) Reset()         { *m = PmPrbsParas{} }
func (m *PmPrbsParas) String() string { return proto.CompactTextString(m) }
func (*PmPrbsParas) ProtoMessage()    {}
func (*PmPrbsParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2609214b0ba90b4, []int{2}
}

func (m *PmPrbsParas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmPrbsParas.Unmarshal(m, b)
}
func (m *PmPrbsParas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmPrbsParas.Marshal(b, m, deterministic)
}
func (m *PmPrbsParas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmPrbsParas.Merge(m, src)
}
func (m *PmPrbsParas) XXX_Size() int {
	return xxx_messageInfo_PmPrbsParas.Size(m)
}
func (m *PmPrbsParas) XXX_DiscardUnknown() {
	xxx_messageInfo_PmPrbsParas.DiscardUnknown(m)
}

var xxx_messageInfo_PmPrbsParas proto.InternalMessageInfo

func (m *PmPrbsParas) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PmPrbsParas) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmPrbsParas) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *PmPrbsParas) GetLastClearTime() string {
	if m != nil {
		return m.LastClearTime
	}
	return ""
}

func (m *PmPrbsParas) GetLastClear15MinTime() string {
	if m != nil {
		return m.LastClear15MinTime
	}
	return ""
}

func (m *PmPrbsParas) GetLastClear24HrTime() string {
	if m != nil {
		return m.LastClear24HrTime
	}
	return ""
}

func (m *PmPrbsParas) GetEbc() uint64 {
	if m != nil {
		return m.Ebc
	}
	return 0
}

func (m *PmPrbsParas) GetFoundCount() uint32 {
	if m != nil {
		return m.FoundCount
	}
	return 0
}

func (m *PmPrbsParas) GetLostCount() uint32 {
	if m != nil {
		return m.LostCount
	}
	return 0
}

func (m *PmPrbsParas) GetFoundAtTime() string {
	if m != nil {
		return m.FoundAtTime
	}
	return ""
}

func (m *PmPrbsParas) GetLostAtTime() string {
	if m != nil {
		return m.LostAtTime
	}
	return ""
}

func (m *PmPrbsParas) GetConfPatt() string {
	if m != nil {
		return m.ConfPatt
	}
	return ""
}

func (m *PmPrbsParas) GetRcvPatt() *PmPrbsParameter {
	if m != nil {
		return m.RcvPatt
	}
	return nil
}

func init() {
	proto.RegisterType((*PmPrbsParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management.odu.odu_ports.odu_port.odu_current.odu_minute15.odu_minute15prbses.odu_minute15prbs.pm_prbs_paras_KEYS")
	proto.RegisterType((*PmPrbsParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management.odu.odu_ports.odu_port.odu_current.odu_minute15.odu_minute15prbses.odu_minute15prbs.pm_prbs_parameter")
	proto.RegisterType((*PmPrbsParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management.odu.odu_ports.odu_port.odu_current.odu_minute15.odu_minute15prbses.odu_minute15prbs.pm_prbs_paras")
}

func init() { proto.RegisterFile("pm_prbs_paras.proto", fileDescriptor_c2609214b0ba90b4) }

var fileDescriptor_c2609214b0ba90b4 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0x5d, 0x8a, 0x13, 0x41,
	0x10, 0x66, 0x34, 0xae, 0x99, 0x5a, 0x83, 0x6e, 0xbb, 0xca, 0x88, 0x8a, 0x43, 0x1e, 0x24, 0x4f,
	0x23, 0xc9, 0x26, 0xfe, 0x2b, 0xca, 0x2a, 0x08, 0x22, 0xc8, 0xe8, 0x8b, 0x4f, 0x4d, 0xa7, 0x53,
	0x59, 0x1b, 0xd2, 0x3f, 0xf4, 0xf4, 0x84, 0xbd, 0x80, 0x77, 0xf0, 0x20, 0x5e, 0xc4, 0x1b, 0x49,
	0x57, 0xc7, 0x4c, 0xc2, 0x5e, 0x60, 0x1f, 0x06, 0xaa, 0xbe, 0xfa, 0xea, 0xfb, 0x6a, 0xa6, 0x6a,
	0xe0, 0xb6, 0xd3, 0xdc, 0xf9, 0x79, 0xc3, 0x9d, 0xf0, 0xa2, 0xa9, 0x9c, 0xb7, 0xc1, 0xb2, 0x5f,
	0x99, 0x54, 0x8d, 0xb4, 0x5c, 0xd9, 0x86, 0x9f, 0x7b, 0xee, 0x34, 0x9a, 0x33, 0x65, 0x90, 0x5b,
	0x87, 0xbe, 0x72, 0xe8, 0x97, 0xd6, 0x6b, 0x61, 0x24, 0x72, 0x2d, 0x8c, 0x38, 0x43, 0x8d, 0x26,
	0x54, 0x76, 0xd1, 0xc6, 0x87, 0x3b, 0xeb, 0x43, 0xb3, 0x8d, 0x28, 0x90, 0xad, 0xf7, 0x1b, 0x0a,
	0xd7, 0xca, 0xb4, 0x01, 0xc7, 0xb3, 0xbd, 0x24, 0x0e, 0x80, 0xcd, 0x05, 0x68, 0xf8, 0x0e, 0xd8,
	0xde, 0x78, 0xfc, 0xf3, 0xc7, 0x1f, 0xdf, 0x18, 0x83, 0x9e, 0x11, 0x1a, 0x8b, 0xac, 0xcc, 0x46,
	0x79, 0x4d, 0x31, 0xbb, 0x0b, 0x07, 0xa6, 0xd5, 0x73, 0xf4, 0xc5, 0x95, 0x32, 0x1b, 0x0d, 0xea,
	0x4d, 0x36, 0xfc, 0x00, 0x47, 0xbb, 0x0a, 0x1a, 0x03, 0x7a, 0x76, 0x0c, 0xd7, 0xd6, 0x62, 0xa5,
	0x16, 0xa4, 0xd0, 0xaf, 0x53, 0xc2, 0xee, 0x41, 0xdf, 0xcb, 0x35, 0x77, 0x22, 0x04, 0x12, 0xc9,
	0xeb, 0xeb, 0x5e, 0xae, 0xbf, 0x8a, 0x10, 0x86, 0x7f, 0x7b, 0x30, 0xd8, 0x1b, 0x24, 0x4a, 0x28,
	0xb3, 0xc0, 0xf3, 0x62, 0x42, 0x76, 0x29, 0xe9, 0x84, 0x4f, 0x76, 0x85, 0x1f, 0x40, 0x1e, 0x94,
	0xc6, 0x26, 0x08, 0xed, 0x8a, 0x29, 0x29, 0x77, 0x00, 0x7b, 0x0c, 0x37, 0x57, 0xa2, 0x09, 0x5c,
	0xae, 0x50, 0x78, 0x1e, 0xf1, 0x62, 0x46, 0x9c, 0x41, 0x84, 0x4f, 0x23, 0xfa, 0x5d, 0x69, 0x64,
	0x63, 0xb8, 0xd3, 0xf1, 0xc6, 0xb3, 0xf8, 0xa1, 0x12, 0xfb, 0x29, 0xb1, 0xd9, 0x96, 0x3d, 0x9e,
	0x7d, 0x51, 0x86, 0x5a, 0x9e, 0xc0, 0x71, 0xd7, 0x32, 0x99, 0xf2, 0x9f, 0x1b, 0xfd, 0x67, 0xd4,
	0x71, 0xb4, 0xed, 0x98, 0x4c, 0x3f, 0x25, 0x8f, 0x5b, 0x70, 0x15, 0xe7, 0xb2, 0x78, 0x5e, 0x66,
	0xa3, 0x5e, 0x1d, 0x43, 0xf6, 0x08, 0x0e, 0x97, 0xb6, 0x35, 0x0b, 0x2e, 0x6d, 0x6b, 0x42, 0xf1,
	0x82, 0xde, 0x16, 0x08, 0x3a, 0x8d, 0x08, 0x7b, 0x08, 0xb0, 0xb2, 0xd1, 0x83, 0xea, 0x2f, 0xa9,
	0x9e, 0x47, 0x24, 0x95, 0x87, 0x30, 0x48, 0xfd, 0x22, 0x24, 0xef, 0x57, 0xe4, 0x9d, 0x44, 0xdf,
	0x07, 0x72, 0x2d, 0xe1, 0x06, 0x49, 0xfc, 0xa7, 0xbc, 0x26, 0x0a, 0xc9, 0x6e, 0x18, 0xf7, 0x21,
	0x97, 0xd6, 0x2c, 0xd3, 0x6e, 0xde, 0x50, 0xb9, 0x1f, 0x81, 0xb8, 0x1c, 0xf6, 0x27, 0xdb, 0x59,
	0xdc, 0xdb, 0x32, 0x1b, 0x1d, 0x4e, 0x7e, 0x67, 0xd5, 0xa5, 0x38, 0xe0, 0xea, 0xc2, 0xed, 0x6d,
	0x6f, 0x6a, 0x7e, 0x40, 0xbf, 0xda, 0xc9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x4e, 0xd2,
	0x84, 0x81, 0x03, 0x00, 0x00,
}
