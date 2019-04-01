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
// source: isis_sh_label.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_sr_labels_sr_label

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

type IsisShLabel_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Label                uint32   `protobuf:"varint,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShLabel_KEYS) Reset()         { *m = IsisShLabel_KEYS{} }
func (m *IsisShLabel_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShLabel_KEYS) ProtoMessage()    {}
func (*IsisShLabel_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f52a5f0783e9d4c7, []int{0}
}

func (m *IsisShLabel_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShLabel_KEYS.Unmarshal(m, b)
}
func (m *IsisShLabel_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShLabel_KEYS.Marshal(b, m, deterministic)
}
func (m *IsisShLabel_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShLabel_KEYS.Merge(m, src)
}
func (m *IsisShLabel_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShLabel_KEYS.Size(m)
}
func (m *IsisShLabel_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShLabel_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShLabel_KEYS proto.InternalMessageInfo

func (m *IsisShLabel_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShLabel_KEYS) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

type IsisIpv4PrefixType struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisIpv4PrefixType) Reset()         { *m = IsisIpv4PrefixType{} }
func (m *IsisIpv4PrefixType) String() string { return proto.CompactTextString(m) }
func (*IsisIpv4PrefixType) ProtoMessage()    {}
func (*IsisIpv4PrefixType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f52a5f0783e9d4c7, []int{1}
}

func (m *IsisIpv4PrefixType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisIpv4PrefixType.Unmarshal(m, b)
}
func (m *IsisIpv4PrefixType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisIpv4PrefixType.Marshal(b, m, deterministic)
}
func (m *IsisIpv4PrefixType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisIpv4PrefixType.Merge(m, src)
}
func (m *IsisIpv4PrefixType) XXX_Size() int {
	return xxx_messageInfo_IsisIpv4PrefixType.Size(m)
}
func (m *IsisIpv4PrefixType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisIpv4PrefixType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisIpv4PrefixType proto.InternalMessageInfo

func (m *IsisIpv4PrefixType) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *IsisIpv4PrefixType) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type IsisIpv6PrefixType struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisIpv6PrefixType) Reset()         { *m = IsisIpv6PrefixType{} }
func (m *IsisIpv6PrefixType) String() string { return proto.CompactTextString(m) }
func (*IsisIpv6PrefixType) ProtoMessage()    {}
func (*IsisIpv6PrefixType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f52a5f0783e9d4c7, []int{2}
}

func (m *IsisIpv6PrefixType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisIpv6PrefixType.Unmarshal(m, b)
}
func (m *IsisIpv6PrefixType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisIpv6PrefixType.Marshal(b, m, deterministic)
}
func (m *IsisIpv6PrefixType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisIpv6PrefixType.Merge(m, src)
}
func (m *IsisIpv6PrefixType) XXX_Size() int {
	return xxx_messageInfo_IsisIpv6PrefixType.Size(m)
}
func (m *IsisIpv6PrefixType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisIpv6PrefixType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisIpv6PrefixType proto.InternalMessageInfo

func (m *IsisIpv6PrefixType) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *IsisIpv6PrefixType) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type IsisIpPrefixType struct {
	AfName               string              `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4                 *IsisIpv4PrefixType `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 *IsisIpv6PrefixType `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IsisIpPrefixType) Reset()         { *m = IsisIpPrefixType{} }
func (m *IsisIpPrefixType) String() string { return proto.CompactTextString(m) }
func (*IsisIpPrefixType) ProtoMessage()    {}
func (*IsisIpPrefixType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f52a5f0783e9d4c7, []int{3}
}

func (m *IsisIpPrefixType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisIpPrefixType.Unmarshal(m, b)
}
func (m *IsisIpPrefixType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisIpPrefixType.Marshal(b, m, deterministic)
}
func (m *IsisIpPrefixType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisIpPrefixType.Merge(m, src)
}
func (m *IsisIpPrefixType) XXX_Size() int {
	return xxx_messageInfo_IsisIpPrefixType.Size(m)
}
func (m *IsisIpPrefixType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisIpPrefixType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisIpPrefixType proto.InternalMessageInfo

func (m *IsisIpPrefixType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisIpPrefixType) GetIpv4() *IsisIpv4PrefixType {
	if m != nil {
		return m.Ipv4
	}
	return nil
}

func (m *IsisIpPrefixType) GetIpv6() *IsisIpv6PrefixType {
	if m != nil {
		return m.Ipv6
	}
	return nil
}

type IsisShLabel struct {
	LabelType            string            `protobuf:"bytes,50,opt,name=label_type,json=labelType,proto3" json:"label_type,omitempty"`
	LabelFlags           uint32            `protobuf:"varint,51,opt,name=label_flags,json=labelFlags,proto3" json:"label_flags,omitempty"`
	LabelRefcount        uint32            `protobuf:"varint,52,opt,name=label_refcount,json=labelRefcount,proto3" json:"label_refcount,omitempty"`
	LabelValue           uint32            `protobuf:"varint,53,opt,name=label_value,json=labelValue,proto3" json:"label_value,omitempty"`
	LabelInterface       string            `protobuf:"bytes,54,opt,name=label_interface,json=labelInterface,proto3" json:"label_interface,omitempty"`
	LabelIfh             string            `protobuf:"bytes,55,opt,name=label_ifh,json=labelIfh,proto3" json:"label_ifh,omitempty"`
	LabelTableId         uint32            `protobuf:"varint,56,opt,name=label_table_id,json=labelTableId,proto3" json:"label_table_id,omitempty"`
	LabelAfId            string            `protobuf:"bytes,57,opt,name=label_af_id,json=labelAfId,proto3" json:"label_af_id,omitempty"`
	LabelPrefix          *IsisIpPrefixType `protobuf:"bytes,58,opt,name=label_prefix,json=labelPrefix,proto3" json:"label_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IsisShLabel) Reset()         { *m = IsisShLabel{} }
func (m *IsisShLabel) String() string { return proto.CompactTextString(m) }
func (*IsisShLabel) ProtoMessage()    {}
func (*IsisShLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_f52a5f0783e9d4c7, []int{4}
}

func (m *IsisShLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShLabel.Unmarshal(m, b)
}
func (m *IsisShLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShLabel.Marshal(b, m, deterministic)
}
func (m *IsisShLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShLabel.Merge(m, src)
}
func (m *IsisShLabel) XXX_Size() int {
	return xxx_messageInfo_IsisShLabel.Size(m)
}
func (m *IsisShLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShLabel.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShLabel proto.InternalMessageInfo

func (m *IsisShLabel) GetLabelType() string {
	if m != nil {
		return m.LabelType
	}
	return ""
}

func (m *IsisShLabel) GetLabelFlags() uint32 {
	if m != nil {
		return m.LabelFlags
	}
	return 0
}

func (m *IsisShLabel) GetLabelRefcount() uint32 {
	if m != nil {
		return m.LabelRefcount
	}
	return 0
}

func (m *IsisShLabel) GetLabelValue() uint32 {
	if m != nil {
		return m.LabelValue
	}
	return 0
}

func (m *IsisShLabel) GetLabelInterface() string {
	if m != nil {
		return m.LabelInterface
	}
	return ""
}

func (m *IsisShLabel) GetLabelIfh() string {
	if m != nil {
		return m.LabelIfh
	}
	return ""
}

func (m *IsisShLabel) GetLabelTableId() uint32 {
	if m != nil {
		return m.LabelTableId
	}
	return 0
}

func (m *IsisShLabel) GetLabelAfId() string {
	if m != nil {
		return m.LabelAfId
	}
	return ""
}

func (m *IsisShLabel) GetLabelPrefix() *IsisIpPrefixType {
	if m != nil {
		return m.LabelPrefix
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShLabel_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.sr_labels.sr_label.isis_sh_label_KEYS")
	proto.RegisterType((*IsisIpv4PrefixType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.sr_labels.sr_label.isis_ipv4_prefix_type")
	proto.RegisterType((*IsisIpv6PrefixType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.sr_labels.sr_label.isis_ipv6_prefix_type")
	proto.RegisterType((*IsisIpPrefixType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.sr_labels.sr_label.isis_ip_prefix_type")
	proto.RegisterType((*IsisShLabel)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.sr_labels.sr_label.isis_sh_label")
}

func init() { proto.RegisterFile("isis_sh_label.proto", fileDescriptor_f52a5f0783e9d4c7) }

var fileDescriptor_f52a5f0783e9d4c7 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x5d, 0xab, 0xd3, 0x30,
	0x18, 0xc7, 0xd9, 0xa6, 0xd3, 0x3d, 0x5b, 0x15, 0xe2, 0x5b, 0x40, 0xd4, 0x31, 0x15, 0x77, 0xd5,
	0x8b, 0x6d, 0xd6, 0x97, 0x3b, 0x2f, 0x1c, 0x0c, 0x45, 0xa5, 0x0e, 0x41, 0x44, 0x42, 0xd6, 0x25,
	0x6b, 0xa0, 0x6b, 0x6b, 0xd3, 0x8d, 0x79, 0xe3, 0xa7, 0xf0, 0x1b, 0x9c, 0x2f, 0x7a, 0xc8, 0x93,
	0xf4, 0xac, 0xe5, 0x9c, 0xcb, 0x9d, 0xbb, 0xe4, 0xdf, 0x7f, 0x7e, 0xcf, 0x2b, 0x85, 0x7b, 0x4a,
	0x2b, 0xcd, 0x74, 0xcc, 0x12, 0xbe, 0x12, 0x89, 0x9f, 0x17, 0x59, 0x99, 0x91, 0x79, 0xa4, 0x74,
	0x94, 0x31, 0x95, 0x69, 0x76, 0x28, 0x58, 0x94, 0xa4, 0x9a, 0xa1, 0x2d, 0xcb, 0x45, 0xe1, 0x9b,
	0x93, 0xaf, 0x52, 0x5d, 0xf2, 0x34, 0x12, 0xc7, 0x93, 0xaf, 0x0b, 0xcb, 0xd0, 0x17, 0xa7, 0xd1,
	0x57, 0x20, 0x0d, 0x3c, 0xfb, 0xf4, 0xf1, 0xe7, 0x77, 0xf2, 0x1c, 0xbc, 0xea, 0x11, 0x4b, 0xf9,
	0x56, 0xd0, 0xd6, 0xb0, 0x35, 0xee, 0x85, 0x83, 0x4a, 0xfc, 0xc2, 0xb7, 0x82, 0xdc, 0x87, 0x9b,
	0xf8, 0x84, 0xb6, 0x87, 0xad, 0xb1, 0x17, 0xda, 0xcb, 0x68, 0x09, 0x0f, 0x10, 0xa8, 0xf2, 0xfd,
	0x8c, 0xe5, 0x85, 0x90, 0xea, 0xc0, 0xca, 0xbf, 0xb9, 0x20, 0x0f, 0xa1, 0x6b, 0xaf, 0x0e, 0xe6,
	0x6e, 0x26, 0x96, 0xb3, 0x25, 0x22, 0xdd, 0x94, 0xb1, 0xc3, 0x0d, 0xac, 0xf8, 0x19, 0xb5, 0x3a,
	0x35, 0x38, 0x1d, 0xf5, 0x7f, 0xdb, 0x35, 0x57, 0xe5, 0x0d, 0xe8, 0x23, 0xb8, 0xc5, 0x65, 0xbd,
	0xf0, 0x2e, 0x97, 0x58, 0xf2, 0x1f, 0xb8, 0x61, 0xea, 0x42, 0x58, 0x7f, 0xf2, 0xdb, 0x3f, 0xcd,
	0x10, 0xfc, 0x2b, 0x1b, 0x16, 0x62, 0x28, 0x17, 0x32, 0xa0, 0x9d, 0xeb, 0x09, 0x19, 0x5c, 0x0a,
	0x19, 0x8c, 0xce, 0x3a, 0xe0, 0x35, 0x96, 0x82, 0x3c, 0x01, 0xb0, 0xdb, 0x61, 0x5c, 0x74, 0x82,
	0x3d, 0xe9, 0xa1, 0xb2, 0x34, 0xfd, 0x7a, 0x06, 0x7d, 0xfb, 0x59, 0x26, 0x7c, 0xa3, 0xe9, 0x14,
	0x5b, 0x6d, 0x5f, 0xcc, 0x8d, 0x42, 0x5e, 0xc2, 0x1d, 0x6b, 0x28, 0x84, 0x8c, 0xb2, 0x5d, 0x5a,
	0xd2, 0x19, 0x7a, 0x3c, 0x54, 0x43, 0x27, 0x1e, 0x39, 0x7b, 0x9e, 0xec, 0x04, 0x7d, 0x5d, 0xe3,
	0xfc, 0x30, 0x0a, 0x79, 0x05, 0x77, 0xad, 0x41, 0xa5, 0xa5, 0x28, 0x24, 0x8f, 0x04, 0x0d, 0x30,
	0x19, 0x8b, 0x5f, 0x54, 0x2a, 0x79, 0x0c, 0x3d, 0x67, 0x94, 0x31, 0x7d, 0x83, 0x96, 0xdb, 0xd6,
	0x22, 0x63, 0xf2, 0xa2, 0xca, 0xa6, 0xe4, 0xab, 0x44, 0x30, 0xb5, 0xa6, 0x6f, 0xed, 0x72, 0xd8,
	0x8a, 0x8c, 0xb8, 0x58, 0x93, 0xa7, 0x55, 0x32, 0x5c, 0x1a, 0xcb, 0xbb, 0x5a, 0xd1, 0x1f, 0xe4,
	0x62, 0x4d, 0xfe, 0x81, 0xf5, 0xbb, 0x06, 0xd2, 0xf7, 0x38, 0xa0, 0x5f, 0x27, 0x1e, 0x50, 0x63,
	0x3c, 0x36, 0xa1, 0x6f, 0xa8, 0xac, 0xba, 0xf8, 0x23, 0x98, 0x9e, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x2b, 0xd9, 0x36, 0xe3, 0x1f, 0x04, 0x00, 0x00,
}
