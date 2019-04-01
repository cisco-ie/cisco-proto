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

package cisco_ios_xr_asic_errors_oper_asic_errors_nodes_node_asic_information_instances_instance_error_path_configuration_hard_errors

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
	proto.RegisterType((*CiherrNodesBg_KEYS)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.configuration_hard_errors.ciherr_nodes_bg_KEYS")
	proto.RegisterType((*CsrsInfoBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.configuration_hard_errors.csrs_info_bg")
	proto.RegisterType((*CiherrErrorDataBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.configuration_hard_errors.ciherr_error_data_bg")
	proto.RegisterType((*CiherrNodeBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.configuration_hard_errors.ciherr_node_bg")
	proto.RegisterType((*CiherrNodesBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.configuration_hard_errors.ciherr_nodes_bg")
}

func init() { proto.RegisterFile("ciherr_nodes_bg.proto", fileDescriptor_1ca32ee840cc252f) }

var fileDescriptor_1ca32ee840cc252f = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0xb5, 0x34, 0x1f, 0x1b, 0x27, 0x29, 0x92, 0x55, 0x84, 0x51, 0x85, 0x08, 0xe1, 0x92,
	0xd3, 0x1e, 0xca, 0x23, 0x40, 0x25, 0xaa, 0x56, 0x45, 0x5a, 0x2a, 0x24, 0x4e, 0x96, 0xe3, 0x9d,
	0x64, 0x2d, 0x36, 0xeb, 0x68, 0xec, 0x02, 0x39, 0xf0, 0x04, 0x88, 0x13, 0xaf, 0xc0, 0x85, 0x13,
	0x17, 0x0e, 0x7d, 0x3c, 0xe4, 0xb1, 0x89, 0x56, 0x15, 0x9c, 0xcb, 0x25, 0xf2, 0xfc, 0xc6, 0x33,
	0xf3, 0x9f, 0x99, 0x78, 0xd9, 0x03, 0x6d, 0x6a, 0x40, 0x94, 0xad, 0xad, 0xc0, 0xc9, 0xe5, 0xba,
	0xd8, 0xa2, 0xf5, 0x96, 0x7f, 0xd6, 0xc6, 0x69, 0x2b, 0x8d, 0x75, 0xf2, 0x13, 0x4a, 0xe5, 0x8c,
	0x96, 0x80, 0x68, 0xd1, 0x49, 0xbb, 0x05, 0x2c, 0x3a, 0xa0, 0xa0, 0x48, 0xfa, 0x8d, 0xd8, 0xb4,
	0x2b, 0x8b, 0x1b, 0xe5, 0x8d, 0x6d, 0x0b, 0xd3, 0x3a, 0xaf, 0x5a, 0x0d, 0x6e, 0x7f, 0x2a, 0x28,
	0x4a, 0x6e, 0x95, 0xaf, 0x0b, 0x6d, 0xdb, 0x95, 0x59, 0x5f, 0x23, 0xdd, 0x95, 0xb5, 0xc2, 0x2a,
	0xe5, 0x9c, 0x37, 0xec, 0xe8, 0x96, 0x2e, 0x79, 0x7e, 0xfa, 0xee, 0x0d, 0x3f, 0x66, 0xa3, 0x00,
	0x64, 0xab, 0x36, 0x20, 0xb2, 0x59, 0xb6, 0x18, 0x95, 0x79, 0x00, 0x97, 0x6a, 0x03, 0x9c, 0xb3,
	0x5e, 0x10, 0x20, 0xee, 0x11, 0xa7, 0x33, 0x7f, 0xc6, 0xa6, 0x49, 0x54, 0x2c, 0x2f, 0x0e, 0x66,
	0xd9, 0x62, 0x5a, 0x4e, 0x02, 0x3c, 0x4b, 0x6c, 0x5e, 0xb2, 0x89, 0x76, 0xe8, 0x48, 0xb9, 0x5c,
	0xae, 0x43, 0xa2, 0x4e, 0x01, 0x3a, 0x73, 0xc1, 0x86, 0xaa, 0xaa, 0x10, 0x9c, 0xa3, 0xfc, 0xbd,
	0xf2, 0x8f, 0xc9, 0x8f, 0x58, 0xff, 0xa3, 0xa9, 0x7c, 0x9d, 0x52, 0x47, 0x63, 0xfe, 0x2b, 0xdb,
	0xb7, 0x10, 0x3b, 0xae, 0x94, 0x57, 0x21, 0xf9, 0x43, 0x36, 0x54, 0x5e, 0x7a, 0x93, 0xf2, 0xf7,
	0xca, 0x81, 0xf2, 0x57, 0x66, 0x03, 0x7c, 0xc6, 0x26, 0xc9, 0x21, 0x5b, 0x07, 0x3a, 0x95, 0x61,
	0xd1, 0x7b, 0xe9, 0x40, 0xf3, 0x27, 0x6c, 0xac, 0xed, 0x75, 0xeb, 0x01, 0xe5, 0x07, 0xd5, 0xa4,
	0x7a, 0x2c, 0xa1, 0xb7, 0xaa, 0xe1, 0x4f, 0xd9, 0x24, 0x16, 0x43, 0x58, 0x87, 0x1b, 0xbd, 0xd9,
	0xc1, 0x62, 0x5a, 0x8e, 0x89, 0x95, 0x84, 0xf8, 0x63, 0xc6, 0x92, 0x1e, 0x70, 0x5a, 0xf4, 0xa9,
	0xc3, 0x11, 0x91, 0x97, 0xe0, 0xf4, 0xfc, 0xa6, 0xcf, 0x0e, 0x3b, 0x93, 0xff, 0xd7, 0x34, 0x8e,
	0xd9, 0x68, 0xbf, 0xeb, 0x34, 0xef, 0x3c, 0x8e, 0x74, 0x65, 0xf9, 0x23, 0x46, 0x3b, 0x91, 0xef,
	0x61, 0x97, 0x34, 0x0e, 0x83, 0x7d, 0x0e, 0xbb, 0xe0, 0x52, 0x8d, 0xc2, 0x8d, 0xb4, 0xad, 0xe8,
	0xcd, 0xb2, 0x45, 0x5e, 0x0e, 0xc9, 0x7e, 0xdd, 0x86, 0x94, 0xbe, 0x46, 0x70, 0xb5, 0xac, 0x0d,
	0xe9, 0x9a, 0x96, 0x79, 0x04, 0xaf, 0x4c, 0x70, 0x6e, 0x01, 0x8d, 0xad, 0x82, 0x73, 0x10, 0x9d,
	0x11, 0x44, 0x67, 0x8a, 0x6c, 0xac, 0x18, 0x76, 0x23, 0x2f, 0x6c, 0x27, 0xb2, 0xb1, 0x22, 0xef,
	0x46, 0x5e, 0xd8, 0xb0, 0x3a, 0x9a, 0x9e, 0x18, 0xc5, 0xd5, 0x91, 0x11, 0x42, 0x4c, 0xeb, 0x51,
	0xfa, 0xdd, 0x16, 0x04, 0x8b, 0x21, 0x01, 0x5c, 0xed, 0xb6, 0x10, 0xd6, 0xd7, 0x80, 0x5a, 0x49,
	0x53, 0x89, 0x31, 0xb9, 0x06, 0xc1, 0x3c, 0xab, 0xc2, 0xec, 0x1b, 0xe5, 0xbc, 0xd4, 0x0d, 0x28,
	0x84, 0x4a, 0x4c, 0x68, 0x7d, 0xe3, 0xc0, 0x5e, 0x44, 0xc4, 0x7f, 0x64, 0x6c, 0xb4, 0xff, 0xa3,
	0x89, 0xe9, 0xec, 0x60, 0x31, 0x3e, 0xf9, 0x92, 0x15, 0x77, 0xfa, 0xd4, 0x8a, 0xee, 0x3f, 0xbf,
	0xcc, 0x83, 0x45, 0x4b, 0xfc, 0x99, 0xb1, 0x9c, 0xfa, 0x01, 0x44, 0x71, 0x48, 0x52, 0xbf, 0xdd,
	0xb9, 0xd4, 0xbf, 0xbc, 0xa7, 0x72, 0x18, 0x54, 0x9e, 0x22, 0xce, 0x6f, 0x32, 0x76, 0xff, 0xd6,
	0x47, 0x83, 0x7f, 0xcf, 0x58, 0x9f, 0xae, 0x8b, 0x13, 0x6a, 0xe1, 0xeb, 0x7f, 0xd2, 0x42, 0x7a,
	0x5b, 0x65, 0x14, 0xb7, 0x1c, 0xd0, 0x47, 0xf7, 0xf9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3,
	0xc7, 0xe3, 0x0f, 0x8d, 0x05, 0x00, 0x00,
}