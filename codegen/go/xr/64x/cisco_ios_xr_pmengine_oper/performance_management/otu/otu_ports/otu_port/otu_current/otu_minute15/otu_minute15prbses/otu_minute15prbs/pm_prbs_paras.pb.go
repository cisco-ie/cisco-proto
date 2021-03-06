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

package cisco_ios_xr_pmengine_oper_performance_management_otu_otu_ports_otu_port_otu_current_otu_minute15_otu_minute15prbses_otu_minute15prbs

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

type PmPrbsStatusParameter struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	PrbsStatus           string   `protobuf:"bytes,2,opt,name=prbs_status,json=prbsStatus,proto3" json:"prbs_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmPrbsStatusParameter) Reset()         { *m = PmPrbsStatusParameter{} }
func (m *PmPrbsStatusParameter) String() string { return proto.CompactTextString(m) }
func (*PmPrbsStatusParameter) ProtoMessage()    {}
func (*PmPrbsStatusParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2609214b0ba90b4, []int{2}
}

func (m *PmPrbsStatusParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmPrbsStatusParameter.Unmarshal(m, b)
}
func (m *PmPrbsStatusParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmPrbsStatusParameter.Marshal(b, m, deterministic)
}
func (m *PmPrbsStatusParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmPrbsStatusParameter.Merge(m, src)
}
func (m *PmPrbsStatusParameter) XXX_Size() int {
	return xxx_messageInfo_PmPrbsStatusParameter.Size(m)
}
func (m *PmPrbsStatusParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PmPrbsStatusParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PmPrbsStatusParameter proto.InternalMessageInfo

func (m *PmPrbsStatusParameter) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PmPrbsStatusParameter) GetPrbsStatus() string {
	if m != nil {
		return m.PrbsStatus
	}
	return ""
}

type PmPrbsParas struct {
	Index                uint32                 `protobuf:"varint,50,opt,name=index,proto3" json:"index,omitempty"`
	Valid                bool                   `protobuf:"varint,51,opt,name=valid,proto3" json:"valid,omitempty"`
	Timestamp            string                 `protobuf:"bytes,52,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LastClearTime        string                 `protobuf:"bytes,53,opt,name=last_clear_time,json=lastClearTime,proto3" json:"last_clear_time,omitempty"`
	LastClear15MinTime   string                 `protobuf:"bytes,54,opt,name=last_clear15_min_time,json=lastClear15MinTime,proto3" json:"last_clear15_min_time,omitempty"`
	LastClear24HrTime    string                 `protobuf:"bytes,55,opt,name=last_clear24_hr_time,json=lastClear24HrTime,proto3" json:"last_clear24_hr_time,omitempty"`
	Ebc                  uint64                 `protobuf:"varint,56,opt,name=ebc,proto3" json:"ebc,omitempty"`
	FoundCount           uint32                 `protobuf:"varint,57,opt,name=found_count,json=foundCount,proto3" json:"found_count,omitempty"`
	LostCount            uint32                 `protobuf:"varint,58,opt,name=lost_count,json=lostCount,proto3" json:"lost_count,omitempty"`
	FoundAtTime          uint64                 `protobuf:"varint,59,opt,name=found_at_time,json=foundAtTime,proto3" json:"found_at_time,omitempty"`
	LostAtTime           uint64                 `protobuf:"varint,60,opt,name=lost_at_time,json=lostAtTime,proto3" json:"lost_at_time,omitempty"`
	ConfPatt             string                 `protobuf:"bytes,61,opt,name=conf_patt,json=confPatt,proto3" json:"conf_patt,omitempty"`
	RcvPatt              *PmPrbsParameter       `protobuf:"bytes,62,opt,name=rcv_patt,json=rcvPatt,proto3" json:"rcv_patt,omitempty"`
	PrbsStatus           *PmPrbsStatusParameter `protobuf:"bytes,63,opt,name=prbs_status,json=prbsStatus,proto3" json:"prbs_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PmPrbsParas) Reset()         { *m = PmPrbsParas{} }
func (m *PmPrbsParas) String() string { return proto.CompactTextString(m) }
func (*PmPrbsParas) ProtoMessage()    {}
func (*PmPrbsParas) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2609214b0ba90b4, []int{3}
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

func (m *PmPrbsParas) GetFoundAtTime() uint64 {
	if m != nil {
		return m.FoundAtTime
	}
	return 0
}

func (m *PmPrbsParas) GetLostAtTime() uint64 {
	if m != nil {
		return m.LostAtTime
	}
	return 0
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

func (m *PmPrbsParas) GetPrbsStatus() *PmPrbsStatusParameter {
	if m != nil {
		return m.PrbsStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*PmPrbsParas_KEYS)(nil), "cisco_ios_xr_pmengine_oper.performance_management.otu.otu_ports.otu_port.otu_current.otu_minute15.otu_minute15prbses.otu_minute15prbs.pm_prbs_paras_KEYS")
	proto.RegisterType((*PmPrbsParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management.otu.otu_ports.otu_port.otu_current.otu_minute15.otu_minute15prbses.otu_minute15prbs.pm_prbs_parameter")
	proto.RegisterType((*PmPrbsStatusParameter)(nil), "cisco_ios_xr_pmengine_oper.performance_management.otu.otu_ports.otu_port.otu_current.otu_minute15.otu_minute15prbses.otu_minute15prbs.pm_prbs_status_parameter")
	proto.RegisterType((*PmPrbsParas)(nil), "cisco_ios_xr_pmengine_oper.performance_management.otu.otu_ports.otu_port.otu_current.otu_minute15.otu_minute15prbses.otu_minute15prbs.pm_prbs_paras")
}

func init() { proto.RegisterFile("pm_prbs_paras.proto", fileDescriptor_c2609214b0ba90b4) }

var fileDescriptor_c2609214b0ba90b4 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0x15, 0x58, 0x4a, 0x33, 0x4b, 0x05, 0x6b, 0x16, 0x14, 0x04, 0x88, 0x28, 0x07, 0xd4,
	0x53, 0x50, 0xbb, 0x2d, 0xdf, 0x9f, 0x5a, 0x90, 0x90, 0x10, 0x12, 0x64, 0xb9, 0x70, 0xb2, 0x5c,
	0xd7, 0x5d, 0x22, 0xd5, 0x1f, 0x72, 0x26, 0xd5, 0xbe, 0x00, 0x8f, 0x80, 0xc4, 0x8d, 0x97, 0xe0,
	0x05, 0x78, 0x33, 0xe4, 0x49, 0xb6, 0x4d, 0xd9, 0x03, 0xd7, 0x3d, 0x54, 0xf2, 0xfc, 0xfd, 0x9f,
	0xdf, 0x8c, 0xe3, 0x71, 0xe1, 0xba, 0xd3, 0xdc, 0xf9, 0x59, 0xc5, 0x9d, 0xf0, 0xa2, 0xca, 0x9d,
	0xb7, 0x68, 0xd9, 0xf7, 0x48, 0x96, 0x95, 0xb4, 0xbc, 0xb4, 0x15, 0x3f, 0xf1, 0xdc, 0x69, 0x65,
	0x8e, 0x4b, 0xa3, 0xb8, 0x75, 0xca, 0xe7, 0x4e, 0xf9, 0x85, 0xf5, 0x5a, 0x18, 0xa9, 0xb8, 0x16,
	0x46, 0x1c, 0x2b, 0xad, 0x0c, 0xe6, 0x16, 0xeb, 0xf0, 0xe3, 0xce, 0x7a, 0xac, 0xd6, 0x2b, 0x5a,
	0xc8, 0xda, 0xfb, 0xd6, 0xc2, 0x75, 0x69, 0x6a, 0x54, 0xa3, 0xe9, 0x56, 0x10, 0x1a, 0x50, 0xd5,
	0x19, 0x29, 0x7b, 0x0d, 0x6c, 0xab, 0x3d, 0xfe, 0xe1, 0xdd, 0xd7, 0x23, 0xc6, 0x60, 0xc7, 0x08,
	0xad, 0x92, 0x28, 0x8d, 0x86, 0x71, 0x41, 0x6b, 0x76, 0x13, 0x7a, 0xa6, 0xd6, 0x33, 0xe5, 0x93,
	0x0b, 0x69, 0x34, 0x1c, 0x14, 0x6d, 0x94, 0xbd, 0x85, 0xbd, 0x2e, 0x41, 0x2b, 0x54, 0x9e, 0xed,
	0xc3, 0xa5, 0x95, 0x58, 0x96, 0x73, 0x22, 0xf4, 0x8b, 0x26, 0x60, 0xb7, 0xa0, 0xef, 0xe5, 0x8a,
	0x3b, 0x81, 0x48, 0x90, 0xb8, 0xb8, 0xec, 0xe5, 0xea, 0x93, 0x40, 0xcc, 0x3e, 0x43, 0x72, 0x4a,
	0xa9, 0x50, 0x60, 0xfd, 0x7f, 0xd8, 0x3d, 0xd8, 0xed, 0xd8, 0x5b, 0x1e, 0x04, 0xe9, 0x88, 0x94,
	0xec, 0x47, 0x0f, 0x06, 0x5b, 0x67, 0x0b, 0xa0, 0xd2, 0xcc, 0xd5, 0x49, 0x32, 0xa6, 0x13, 0x34,
	0xc1, 0x06, 0x7f, 0xd0, 0xc5, 0xdf, 0x81, 0x18, 0x4b, 0xad, 0x2a, 0x14, 0xda, 0x25, 0x13, 0x82,
	0x6f, 0x04, 0x76, 0x1f, 0xae, 0x2e, 0x45, 0x85, 0x5c, 0x2e, 0x95, 0xf0, 0x3c, 0xe8, 0xc9, 0x94,
	0x3c, 0x83, 0x20, 0x1f, 0x06, 0xf5, 0x4b, 0xa9, 0x15, 0x1b, 0xc1, 0x8d, 0x8d, 0x6f, 0x34, 0x0d,
	0xdf, 0xbe, 0x71, 0x3f, 0x24, 0x37, 0x5b, 0xbb, 0x47, 0xd3, 0x8f, 0xa5, 0xa1, 0x94, 0x07, 0xb0,
	0xbf, 0x49, 0x19, 0x4f, 0xf8, 0xb7, 0x96, 0xff, 0x88, 0x32, 0xf6, 0xd6, 0x19, 0xe3, 0xc9, 0xfb,
	0xa6, 0xc6, 0x35, 0xb8, 0xa8, 0x66, 0x32, 0x79, 0x9c, 0x46, 0xc3, 0x9d, 0x22, 0x2c, 0xc3, 0xa7,
	0x59, 0xd8, 0xda, 0xcc, 0xb9, 0xb4, 0xb5, 0xc1, 0xe4, 0x09, 0x9d, 0x16, 0x48, 0x3a, 0x0c, 0x0a,
	0xbb, 0x0b, 0xb0, 0xb4, 0xa1, 0x06, 0xed, 0x3f, 0xa5, 0xfd, 0x38, 0x28, 0xcd, 0x76, 0x06, 0x83,
	0x26, 0x5f, 0x60, 0x53, 0xfb, 0x19, 0xb1, 0x1b, 0xe8, 0x1b, 0xa4, 0xaa, 0x29, 0x5c, 0x21, 0xc4,
	0xa9, 0xe5, 0x39, 0x59, 0x08, 0xdb, 0x3a, 0x6e, 0x43, 0x2c, 0xad, 0x59, 0x34, 0xd7, 0xfd, 0x82,
	0xba, 0xef, 0x07, 0x21, 0xdc, 0x37, 0xfb, 0x1d, 0x75, 0x66, 0xe1, 0x65, 0x1a, 0x0d, 0x77, 0xc7,
	0x3f, 0xa3, 0xfc, 0x5c, 0xbc, 0x89, 0xfc, 0xcc, 0x38, 0xaf, 0xc7, 0x94, 0xfd, 0x89, 0xb6, 0xa7,
	0xee, 0x15, 0x75, 0xfe, 0xeb, 0xbc, 0x75, 0xfe, 0xef, 0x13, 0xea, 0xbe, 0x8b, 0x59, 0x8f, 0xfe,
	0x81, 0x0e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x7b, 0xcc, 0x60, 0x98, 0x04, 0x00, 0x00,
}
