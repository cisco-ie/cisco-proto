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
// source: ciherr_nodes_bg.proto

package cisco_ios_xr_asic_errors_oper_asic_errors_nodes_node_asic_information_instances_instance_error_path_multiple_bit_hard_errors

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

type CiherrNodesBg_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Asic                 string   `protobuf:"bytes,2,opt,name=asic,proto3" json:"asic,omitempty"`
	AsicInstance         uint32   `protobuf:"varint,3,opt,name=asic_instance,json=asicInstance,proto3" json:"asic_instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CiherrNodesBg_KEYS) Reset()         { *m = CiherrNodesBg_KEYS{} }
func (m *CiherrNodesBg_KEYS) String() string { return proto.CompactTextString(m) }
func (*CiherrNodesBg_KEYS) ProtoMessage()    {}
func (*CiherrNodesBg_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ca32ee840cc252f, []int{0}
}

func (m *CiherrNodesBg_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CiherrNodesBg_KEYS.Unmarshal(m, b)
}
func (m *CiherrNodesBg_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CiherrNodesBg_KEYS.Marshal(b, m, deterministic)
}
func (m *CiherrNodesBg_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CiherrNodesBg_KEYS.Merge(m, src)
}
func (m *CiherrNodesBg_KEYS) XXX_Size() int {
	return xxx_messageInfo_CiherrNodesBg_KEYS.Size(m)
}
func (m *CiherrNodesBg_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_CiherrNodesBg_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_CiherrNodesBg_KEYS proto.InternalMessageInfo

func (m *CiherrNodesBg_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *CiherrNodesBg_KEYS) GetAsic() string {
	if m != nil {
		return m.Asic
	}
	return ""
}

func (m *CiherrNodesBg_KEYS) GetAsicInstance() uint32 {
	if m != nil {
		return m.AsicInstance
	}
	return 0
}

type CsrsInfoBg struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              uint64   `protobuf:"varint,2,opt,name=address,proto3" json:"address,omitempty"`
	Width                uint32   `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsrsInfoBg) Reset()         { *m = CsrsInfoBg{} }
func (m *CsrsInfoBg) String() string { return proto.CompactTextString(m) }
func (*CsrsInfoBg) ProtoMessage()    {}
func (*CsrsInfoBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ca32ee840cc252f, []int{1}
}

func (m *CsrsInfoBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsrsInfoBg.Unmarshal(m, b)
}
func (m *CsrsInfoBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsrsInfoBg.Marshal(b, m, deterministic)
}
func (m *CsrsInfoBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsrsInfoBg.Merge(m, src)
}
func (m *CsrsInfoBg) XXX_Size() int {
	return xxx_messageInfo_CsrsInfoBg.Size(m)
}
func (m *CsrsInfoBg) XXX_DiscardUnknown() {
	xxx_messageInfo_CsrsInfoBg.DiscardUnknown(m)
}

var xxx_messageInfo_CsrsInfoBg proto.InternalMessageInfo

func (m *CsrsInfoBg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CsrsInfoBg) GetAddress() uint64 {
	if m != nil {
		return m.Address
	}
	return 0
}

func (m *CsrsInfoBg) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

type CiherrErrorDataBg struct {
	AtTime               uint64   `protobuf:"varint,1,opt,name=at_time,json=atTime,proto3" json:"at_time,omitempty"`
	AtTimeNsec           uint64   `protobuf:"varint,2,opt,name=at_time_nsec,json=atTimeNsec,proto3" json:"at_time_nsec,omitempty"`
	CounterVal           uint32   `protobuf:"varint,3,opt,name=counter_val,json=counterVal,proto3" json:"counter_val,omitempty"`
	ErrorRegval          []uint32 `protobuf:"varint,4,rep,packed,name=error_regval,json=errorRegval,proto3" json:"error_regval,omitempty"`
	ErrorDesc            string   `protobuf:"bytes,5,opt,name=error_desc,json=errorDesc,proto3" json:"error_desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CiherrErrorDataBg) Reset()         { *m = CiherrErrorDataBg{} }
func (m *CiherrErrorDataBg) String() string { return proto.CompactTextString(m) }
func (*CiherrErrorDataBg) ProtoMessage()    {}
func (*CiherrErrorDataBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ca32ee840cc252f, []int{2}
}

func (m *CiherrErrorDataBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CiherrErrorDataBg.Unmarshal(m, b)
}
func (m *CiherrErrorDataBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CiherrErrorDataBg.Marshal(b, m, deterministic)
}
func (m *CiherrErrorDataBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CiherrErrorDataBg.Merge(m, src)
}
func (m *CiherrErrorDataBg) XXX_Size() int {
	return xxx_messageInfo_CiherrErrorDataBg.Size(m)
}
func (m *CiherrErrorDataBg) XXX_DiscardUnknown() {
	xxx_messageInfo_CiherrErrorDataBg.DiscardUnknown(m)
}

var xxx_messageInfo_CiherrErrorDataBg proto.InternalMessageInfo

func (m *CiherrErrorDataBg) GetAtTime() uint64 {
	if m != nil {
		return m.AtTime
	}
	return 0
}

func (m *CiherrErrorDataBg) GetAtTimeNsec() uint64 {
	if m != nil {
		return m.AtTimeNsec
	}
	return 0
}

func (m *CiherrErrorDataBg) GetCounterVal() uint32 {
	if m != nil {
		return m.CounterVal
	}
	return 0
}

func (m *CiherrErrorDataBg) GetErrorRegval() []uint32 {
	if m != nil {
		return m.ErrorRegval
	}
	return nil
}

func (m *CiherrErrorDataBg) GetErrorDesc() string {
	if m != nil {
		return m.ErrorDesc
	}
	return ""
}

type CiherrNodeBg struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AsicInfo             string               `protobuf:"bytes,2,opt,name=asic_info,json=asicInfo,proto3" json:"asic_info,omitempty"`
	NodeKey              uint32               `protobuf:"varint,3,opt,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
	AlarmOn              bool                 `protobuf:"varint,4,opt,name=alarm_on,json=alarmOn,proto3" json:"alarm_on,omitempty"`
	ThreshHi             uint32               `protobuf:"varint,5,opt,name=thresh_hi,json=threshHi,proto3" json:"thresh_hi,omitempty"`
	PeriodHi             uint32               `protobuf:"varint,6,opt,name=period_hi,json=periodHi,proto3" json:"period_hi,omitempty"`
	ThreshLo             uint32               `protobuf:"varint,7,opt,name=thresh_lo,json=threshLo,proto3" json:"thresh_lo,omitempty"`
	PeriodLo             uint32               `protobuf:"varint,8,opt,name=period_lo,json=periodLo,proto3" json:"period_lo,omitempty"`
	Count                uint32               `protobuf:"varint,9,opt,name=count,proto3" json:"count,omitempty"`
	IntrType             uint32               `protobuf:"varint,10,opt,name=intr_type,json=intrType,proto3" json:"intr_type,omitempty"`
	LeafId               uint32               `protobuf:"varint,11,opt,name=leaf_id,json=leafId,proto3" json:"leaf_id,omitempty"`
	LastCleared          uint64               `protobuf:"varint,12,opt,name=last_cleared,json=lastCleared,proto3" json:"last_cleared,omitempty"`
	CsrsInfo             []*CsrsInfoBg        `protobuf:"bytes,13,rep,name=csrs_info,json=csrsInfo,proto3" json:"csrs_info,omitempty"`
	LastErr              []*CiherrErrorDataBg `protobuf:"bytes,14,rep,name=last_err,json=lastErr,proto3" json:"last_err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CiherrNodeBg) Reset()         { *m = CiherrNodeBg{} }
func (m *CiherrNodeBg) String() string { return proto.CompactTextString(m) }
func (*CiherrNodeBg) ProtoMessage()    {}
func (*CiherrNodeBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ca32ee840cc252f, []int{3}
}

func (m *CiherrNodeBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CiherrNodeBg.Unmarshal(m, b)
}
func (m *CiherrNodeBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CiherrNodeBg.Marshal(b, m, deterministic)
}
func (m *CiherrNodeBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CiherrNodeBg.Merge(m, src)
}
func (m *CiherrNodeBg) XXX_Size() int {
	return xxx_messageInfo_CiherrNodeBg.Size(m)
}
func (m *CiherrNodeBg) XXX_DiscardUnknown() {
	xxx_messageInfo_CiherrNodeBg.DiscardUnknown(m)
}

var xxx_messageInfo_CiherrNodeBg proto.InternalMessageInfo

func (m *CiherrNodeBg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CiherrNodeBg) GetAsicInfo() string {
	if m != nil {
		return m.AsicInfo
	}
	return ""
}

func (m *CiherrNodeBg) GetNodeKey() uint32 {
	if m != nil {
		return m.NodeKey
	}
	return 0
}

func (m *CiherrNodeBg) GetAlarmOn() bool {
	if m != nil {
		return m.AlarmOn
	}
	return false
}

func (m *CiherrNodeBg) GetThreshHi() uint32 {
	if m != nil {
		return m.ThreshHi
	}
	return 0
}

func (m *CiherrNodeBg) GetPeriodHi() uint32 {
	if m != nil {
		return m.PeriodHi
	}
	return 0
}

func (m *CiherrNodeBg) GetThreshLo() uint32 {
	if m != nil {
		return m.ThreshLo
	}
	return 0
}

func (m *CiherrNodeBg) GetPeriodLo() uint32 {
	if m != nil {
		return m.PeriodLo
	}
	return 0
}

func (m *CiherrNodeBg) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *CiherrNodeBg) GetIntrType() uint32 {
	if m != nil {
		return m.IntrType
	}
	return 0
}

func (m *CiherrNodeBg) GetLeafId() uint32 {
	if m != nil {
		return m.LeafId
	}
	return 0
}

func (m *CiherrNodeBg) GetLastCleared() uint64 {
	if m != nil {
		return m.LastCleared
	}
	return 0
}

func (m *CiherrNodeBg) GetCsrsInfo() []*CsrsInfoBg {
	if m != nil {
		return m.CsrsInfo
	}
	return nil
}

func (m *CiherrNodeBg) GetLastErr() []*CiherrErrorDataBg {
	if m != nil {
		return m.LastErr
	}
	return nil
}

type CiherrNodesBg struct {
	Error                []*CiherrNodeBg `protobuf:"bytes,50,rep,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CiherrNodesBg) Reset()         { *m = CiherrNodesBg{} }
func (m *CiherrNodesBg) String() string { return proto.CompactTextString(m) }
func (*CiherrNodesBg) ProtoMessage()    {}
func (*CiherrNodesBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ca32ee840cc252f, []int{4}
}

func (m *CiherrNodesBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CiherrNodesBg.Unmarshal(m, b)
}
func (m *CiherrNodesBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CiherrNodesBg.Marshal(b, m, deterministic)
}
func (m *CiherrNodesBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CiherrNodesBg.Merge(m, src)
}
func (m *CiherrNodesBg) XXX_Size() int {
	return xxx_messageInfo_CiherrNodesBg.Size(m)
}
func (m *CiherrNodesBg) XXX_DiscardUnknown() {
	xxx_messageInfo_CiherrNodesBg.DiscardUnknown(m)
}

var xxx_messageInfo_CiherrNodesBg proto.InternalMessageInfo

func (m *CiherrNodesBg) GetError() []*CiherrNodeBg {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*CiherrNodesBg_KEYS)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.multiple_bit_hard_errors.ciherr_nodes_bg_KEYS")
	proto.RegisterType((*CsrsInfoBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.multiple_bit_hard_errors.csrs_info_bg")
	proto.RegisterType((*CiherrErrorDataBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.multiple_bit_hard_errors.ciherr_error_data_bg")
	proto.RegisterType((*CiherrNodeBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.multiple_bit_hard_errors.ciherr_node_bg")
	proto.RegisterType((*CiherrNodesBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.multiple_bit_hard_errors.ciherr_nodes_bg")
}

func init() { proto.RegisterFile("ciherr_nodes_bg.proto", fileDescriptor_1ca32ee840cc252f) }

var fileDescriptor_1ca32ee840cc252f = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcd, 0x6a, 0x14, 0x41,
	0x10, 0xc7, 0x19, 0xb3, 0x1f, 0xb3, 0xb5, 0xbb, 0x11, 0x9a, 0x88, 0x2d, 0x41, 0x5c, 0xd7, 0xcb,
	0x9e, 0xe6, 0x10, 0x1f, 0x41, 0x03, 0x86, 0x84, 0x08, 0x63, 0x10, 0x3c, 0x35, 0xbd, 0x3d, 0x95,
	0x4c, 0xe3, 0xec, 0xf4, 0xd0, 0xdd, 0x51, 0x17, 0x7c, 0x02, 0xf1, 0xe4, 0x2b, 0x78, 0xf1, 0xe4,
	0x49, 0xf4, 0xf5, 0xa4, 0xab, 0xdb, 0x65, 0x08, 0x7a, 0x8e, 0x97, 0xa1, 0xeb, 0x57, 0x53, 0x55,
	0xff, 0xaa, 0x9a, 0x1e, 0xb8, 0xa7, 0x74, 0x8d, 0xd6, 0x8a, 0xd6, 0x54, 0xe8, 0xc4, 0xfa, 0xaa,
	0xe8, 0xac, 0xf1, 0x86, 0x7d, 0x54, 0xda, 0x29, 0x23, 0xb4, 0x71, 0xe2, 0x83, 0x15, 0xd2, 0x69,
	0x25, 0xd0, 0x5a, 0x63, 0x9d, 0x30, 0x1d, 0xda, 0xa2, 0x07, 0x0a, 0x8a, 0xa4, 0x67, 0xc4, 0xba,
	0xbd, 0x34, 0x76, 0x23, 0xbd, 0x36, 0x6d, 0xa1, 0x5b, 0xe7, 0x65, 0xab, 0xd0, 0xed, 0x4e, 0x05,
	0x45, 0x89, 0x4e, 0xfa, 0xba, 0xd8, 0x5c, 0x37, 0x5e, 0x77, 0x0d, 0x8a, 0xb5, 0xf6, 0xa2, 0x96,
	0xb6, 0x4a, 0x29, 0x97, 0x0d, 0x1c, 0xdc, 0x90, 0x25, 0x4e, 0x8f, 0xdf, 0xbc, 0x62, 0x87, 0x30,
	0x09, 0x40, 0xb4, 0x72, 0x83, 0x3c, 0x5b, 0x64, 0xab, 0x49, 0x99, 0x07, 0x70, 0x2e, 0x37, 0xc8,
	0x18, 0x0c, 0x42, 0x7d, 0x7e, 0x87, 0x38, 0x9d, 0xd9, 0x13, 0x98, 0x27, 0x4d, 0xb1, 0x3a, 0xdf,
	0x5b, 0x64, 0xab, 0x79, 0x39, 0x0b, 0xf0, 0x24, 0xb1, 0x65, 0x09, 0x33, 0xe5, 0xac, 0x23, 0xe1,
	0x62, 0x7d, 0x15, 0x12, 0xf5, 0x0a, 0xd0, 0x99, 0x71, 0x18, 0xcb, 0xaa, 0xb2, 0xe8, 0x1c, 0xe5,
	0x1f, 0x94, 0x7f, 0x4c, 0x76, 0x00, 0xc3, 0xf7, 0xba, 0xf2, 0x75, 0x4a, 0x1d, 0x8d, 0xe5, 0x8f,
	0x6c, 0xd7, 0x42, 0x6c, 0xb8, 0x92, 0x5e, 0x86, 0xe4, 0xf7, 0x61, 0x2c, 0xbd, 0xf0, 0x3a, 0xe5,
	0x1f, 0x94, 0x23, 0xe9, 0x2f, 0xf4, 0x06, 0xd9, 0x02, 0x66, 0xc9, 0x21, 0x5a, 0x87, 0x2a, 0x95,
	0x81, 0xe8, 0x3d, 0x77, 0xa8, 0xd8, 0x23, 0x98, 0x2a, 0x73, 0xdd, 0x7a, 0xb4, 0xe2, 0x9d, 0x6c,
	0x52, 0x3d, 0x48, 0xe8, 0xb5, 0x6c, 0xd8, 0x63, 0x98, 0xc5, 0x62, 0x16, 0xaf, 0xc2, 0x1b, 0x83,
	0xc5, 0xde, 0x6a, 0x5e, 0x4e, 0x89, 0x95, 0x84, 0xd8, 0x43, 0x80, 0xa4, 0x07, 0x9d, 0xe2, 0x43,
	0xea, 0x70, 0x42, 0xe4, 0x39, 0x3a, 0xb5, 0xfc, 0x39, 0x84, 0xfd, 0xde, 0xe4, 0xff, 0x35, 0x8d,
	0x43, 0x98, 0xec, 0x56, 0x9d, 0xe6, 0x9d, 0xc7, 0x91, 0x5e, 0x1a, 0xf6, 0x00, 0x68, 0x27, 0xe2,
	0x2d, 0x6e, 0x93, 0xc6, 0x71, 0xb0, 0x4f, 0x71, 0x1b, 0x5c, 0xb2, 0x91, 0x76, 0x23, 0x4c, 0xcb,
	0x07, 0x8b, 0x6c, 0x95, 0x97, 0x63, 0xb2, 0x5f, 0xb6, 0x21, 0xa5, 0xaf, 0x2d, 0xba, 0x5a, 0xd4,
	0x9a, 0x74, 0xcd, 0xcb, 0x3c, 0x82, 0x17, 0x3a, 0x38, 0x3b, 0xb4, 0xda, 0x54, 0xc1, 0x39, 0x8a,
	0xce, 0x08, 0xa2, 0x33, 0x45, 0x36, 0x86, 0x8f, 0xfb, 0x91, 0x67, 0xa6, 0x17, 0xd9, 0x18, 0x9e,
	0xf7, 0x23, 0xcf, 0x4c, 0x58, 0x1d, 0x4d, 0x8f, 0x4f, 0xe2, 0xea, 0xc8, 0x08, 0x21, 0xba, 0xf5,
	0x56, 0xf8, 0x6d, 0x87, 0x1c, 0x62, 0x48, 0x00, 0x17, 0xdb, 0x0e, 0xc3, 0xfa, 0x1a, 0x94, 0x97,
	0x42, 0x57, 0x7c, 0x4a, 0xae, 0x51, 0x30, 0x4f, 0xaa, 0x30, 0xfb, 0x46, 0x3a, 0x2f, 0x54, 0x83,
	0xd2, 0x62, 0xc5, 0x67, 0xb4, 0xbe, 0x69, 0x60, 0xcf, 0x22, 0x62, 0xdf, 0x32, 0x98, 0xec, 0x3e,
	0x34, 0x3e, 0x5f, 0xec, 0xad, 0xa6, 0x47, 0x9f, 0xb2, 0xe2, 0x36, 0x6f, 0x5a, 0xd1, 0xff, 0xf0,
	0xcb, 0x3c, 0x58, 0xb4, 0xc3, 0xef, 0x19, 0xe4, 0xd4, 0x0e, 0x5a, 0xcb, 0xf7, 0x49, 0xe9, 0x97,
	0xdb, 0x56, 0xfa, 0x97, 0xdb, 0x54, 0x8e, 0x83, 0xc8, 0x63, 0x6b, 0x97, 0xbf, 0x32, 0xb8, 0x7b,
	0xe3, 0x97, 0xc1, 0xbe, 0x66, 0x30, 0xa4, 0xd7, 0xf9, 0x11, 0x75, 0xf0, 0xf9, 0xff, 0xe8, 0x20,
	0x5d, 0xac, 0x32, 0x6a, 0x5b, 0x8f, 0xe8, 0x87, 0xfb, 0xf4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x25, 0xfb, 0x27, 0x39, 0x89, 0x05, 0x00, 0x00,
}
