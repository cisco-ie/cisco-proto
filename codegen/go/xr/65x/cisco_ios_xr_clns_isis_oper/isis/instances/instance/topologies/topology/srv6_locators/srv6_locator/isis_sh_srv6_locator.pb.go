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
// source: isis_sh_srv6_locator.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_srv6_locators_srv6_locator

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

type IsisShSrv6Locator_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	TopologyName         string   `protobuf:"bytes,4,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	Prefix               string   `protobuf:"bytes,5,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShSrv6Locator_KEYS) Reset()         { *m = IsisShSrv6Locator_KEYS{} }
func (m *IsisShSrv6Locator_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShSrv6Locator_KEYS) ProtoMessage()    {}
func (*IsisShSrv6Locator_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdf5e03c041bf06, []int{0}
}

func (m *IsisShSrv6Locator_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShSrv6Locator_KEYS.Unmarshal(m, b)
}
func (m *IsisShSrv6Locator_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShSrv6Locator_KEYS.Marshal(b, m, deterministic)
}
func (m *IsisShSrv6Locator_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShSrv6Locator_KEYS.Merge(m, src)
}
func (m *IsisShSrv6Locator_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShSrv6Locator_KEYS.Size(m)
}
func (m *IsisShSrv6Locator_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShSrv6Locator_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShSrv6Locator_KEYS proto.InternalMessageInfo

func (m *IsisShSrv6Locator_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShSrv6Locator_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShSrv6Locator_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisShSrv6Locator_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *IsisShSrv6Locator_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *IsisShSrv6Locator_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type IsisSrv6RsvdOpcode struct {
	FunctionType         string   `protobuf:"bytes,1,opt,name=function_type,json=functionType,proto3" json:"function_type,omitempty"`
	FunctionOpcode       uint32   `protobuf:"varint,2,opt,name=function_opcode,json=functionOpcode,proto3" json:"function_opcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisSrv6RsvdOpcode) Reset()         { *m = IsisSrv6RsvdOpcode{} }
func (m *IsisSrv6RsvdOpcode) String() string { return proto.CompactTextString(m) }
func (*IsisSrv6RsvdOpcode) ProtoMessage()    {}
func (*IsisSrv6RsvdOpcode) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdf5e03c041bf06, []int{1}
}

func (m *IsisSrv6RsvdOpcode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisSrv6RsvdOpcode.Unmarshal(m, b)
}
func (m *IsisSrv6RsvdOpcode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisSrv6RsvdOpcode.Marshal(b, m, deterministic)
}
func (m *IsisSrv6RsvdOpcode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisSrv6RsvdOpcode.Merge(m, src)
}
func (m *IsisSrv6RsvdOpcode) XXX_Size() int {
	return xxx_messageInfo_IsisSrv6RsvdOpcode.Size(m)
}
func (m *IsisSrv6RsvdOpcode) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisSrv6RsvdOpcode.DiscardUnknown(m)
}

var xxx_messageInfo_IsisSrv6RsvdOpcode proto.InternalMessageInfo

func (m *IsisSrv6RsvdOpcode) GetFunctionType() string {
	if m != nil {
		return m.FunctionType
	}
	return ""
}

func (m *IsisSrv6RsvdOpcode) GetFunctionOpcode() uint32 {
	if m != nil {
		return m.FunctionOpcode
	}
	return 0
}

type IsisShSrv6Locator struct {
	LocatorName             string                `protobuf:"bytes,50,opt,name=locator_name,json=locatorName,proto3" json:"locator_name,omitempty"`
	LocatorId               uint32                `protobuf:"varint,51,opt,name=locator_id,json=locatorId,proto3" json:"locator_id,omitempty"`
	LocatorFormat           uint32                `protobuf:"varint,52,opt,name=locator_format,json=locatorFormat,proto3" json:"locator_format,omitempty"`
	LocatorFlags            uint32                `protobuf:"varint,53,opt,name=locator_flags,json=locatorFlags,proto3" json:"locator_flags,omitempty"`
	NumberOfReservedOpcodes uint32                `protobuf:"varint,54,opt,name=number_of_reserved_opcodes,json=numberOfReservedOpcodes,proto3" json:"number_of_reserved_opcodes,omitempty"`
	Opcode                  []*IsisSrv6RsvdOpcode `protobuf:"bytes,55,rep,name=opcode,proto3" json:"opcode,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *IsisShSrv6Locator) Reset()         { *m = IsisShSrv6Locator{} }
func (m *IsisShSrv6Locator) String() string { return proto.CompactTextString(m) }
func (*IsisShSrv6Locator) ProtoMessage()    {}
func (*IsisShSrv6Locator) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdf5e03c041bf06, []int{2}
}

func (m *IsisShSrv6Locator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShSrv6Locator.Unmarshal(m, b)
}
func (m *IsisShSrv6Locator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShSrv6Locator.Marshal(b, m, deterministic)
}
func (m *IsisShSrv6Locator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShSrv6Locator.Merge(m, src)
}
func (m *IsisShSrv6Locator) XXX_Size() int {
	return xxx_messageInfo_IsisShSrv6Locator.Size(m)
}
func (m *IsisShSrv6Locator) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShSrv6Locator.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShSrv6Locator proto.InternalMessageInfo

func (m *IsisShSrv6Locator) GetLocatorName() string {
	if m != nil {
		return m.LocatorName
	}
	return ""
}

func (m *IsisShSrv6Locator) GetLocatorId() uint32 {
	if m != nil {
		return m.LocatorId
	}
	return 0
}

func (m *IsisShSrv6Locator) GetLocatorFormat() uint32 {
	if m != nil {
		return m.LocatorFormat
	}
	return 0
}

func (m *IsisShSrv6Locator) GetLocatorFlags() uint32 {
	if m != nil {
		return m.LocatorFlags
	}
	return 0
}

func (m *IsisShSrv6Locator) GetNumberOfReservedOpcodes() uint32 {
	if m != nil {
		return m.NumberOfReservedOpcodes
	}
	return 0
}

func (m *IsisShSrv6Locator) GetOpcode() []*IsisSrv6RsvdOpcode {
	if m != nil {
		return m.Opcode
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShSrv6Locator_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.srv6_locators.srv6_locator.isis_sh_srv6_locator_KEYS")
	proto.RegisterType((*IsisSrv6RsvdOpcode)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.srv6_locators.srv6_locator.isis_srv6_rsvd_opcode")
	proto.RegisterType((*IsisShSrv6Locator)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.srv6_locators.srv6_locator.isis_sh_srv6_locator")
}

func init() { proto.RegisterFile("isis_sh_srv6_locator.proto", fileDescriptor_9cdf5e03c041bf06) }

var fileDescriptor_9cdf5e03c041bf06 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0x5d, 0xc8, 0xb2, 0xb3, 0xcd, 0x22, 0x59, 0xc0, 0x66, 0x2b, 0x21, 0x95, 0x56,
	0x88, 0x9e, 0x72, 0x68, 0xa1, 0x1c, 0x38, 0x83, 0x84, 0x40, 0x54, 0x0a, 0x5c, 0x38, 0x59, 0x6e,
	0xe2, 0xb4, 0x96, 0x12, 0x3b, 0xf2, 0xb8, 0x55, 0xf3, 0x18, 0xbc, 0x1c, 0x6f, 0xc2, 0x1d, 0x65,
	0x1c, 0x17, 0x90, 0x7a, 0xe5, 0x36, 0xfe, 0xfe, 0xdf, 0xe3, 0xf9, 0xa7, 0x0d, 0x8c, 0x15, 0x2a,
	0xe4, 0xb8, 0xe3, 0x68, 0x0f, 0x2b, 0x5e, 0x9b, 0x42, 0x38, 0x63, 0xb3, 0xd6, 0x1a, 0x67, 0xd8,
	0xa6, 0x50, 0x58, 0x18, 0xae, 0x0c, 0xf2, 0xa3, 0xe5, 0x45, 0xad, 0x91, 0x93, 0xdb, 0xb4, 0xd2,
	0x66, 0x7d, 0x95, 0x29, 0x8d, 0x4e, 0xe8, 0x42, 0xfe, 0xa9, 0x32, 0x67, 0x5a, 0x53, 0x9b, 0xad,
	0x92, 0x18, 0xca, 0x2e, 0xfb, 0xbb, 0x37, 0xfe, 0x73, 0x9a, 0xfe, 0x8c, 0xe0, 0xfe, 0xdc, 0x08,
	0xfc, 0xd3, 0xfb, 0xef, 0x5f, 0xd9, 0x0c, 0x92, 0xd0, 0x98, 0x6b, 0xd1, 0xc8, 0x34, 0x9a, 0x44,
	0xf3, 0xeb, 0x7c, 0x14, 0xe0, 0x17, 0xd1, 0x48, 0x76, 0x07, 0x57, 0xa2, 0xf2, 0xf2, 0x05, 0xc9,
	0xb1, 0xa8, 0x48, 0xb8, 0x87, 0x47, 0x18, 0x94, 0x4b, 0x52, 0xae, 0x70, 0x90, 0x66, 0x90, 0x84,
	0xe9, 0xbc, 0xfe, 0xc0, 0x37, 0x0e, 0x90, 0x4c, 0xcf, 0x20, 0x6e, 0xad, 0xac, 0xd4, 0x31, 0x7d,
	0xe8, 0xfb, 0xfa, 0x53, 0x7f, 0xd9, 0x57, 0xbc, 0x96, 0x7a, 0xeb, 0x76, 0x69, 0x3c, 0x89, 0xe6,
	0x49, 0x3e, 0xf2, 0xf0, 0x33, 0xb1, 0xa9, 0x84, 0xa7, 0x3e, 0x57, 0x1f, 0xca, 0xe2, 0xa1, 0xe4,
	0xa6, 0x2d, 0x4c, 0x49, 0x4f, 0x57, 0x7b, 0x5d, 0x38, 0x65, 0x34, 0x77, 0x5d, 0x7b, 0xca, 0x14,
	0xe0, 0xb7, 0xae, 0x95, 0xec, 0x15, 0x3c, 0x3e, 0x99, 0xfc, 0x3d, 0xca, 0x96, 0xe4, 0xb7, 0x01,
	0xaf, 0x89, 0x4e, 0x7f, 0x5d, 0xc0, 0x93, 0x73, 0xfb, 0x63, 0x2f, 0x60, 0x14, 0x56, 0x49, 0x01,
	0x17, 0xf4, 0xca, 0xcd, 0xc0, 0x28, 0xdf, 0x73, 0x80, 0x60, 0x51, 0x65, 0xba, 0xa4, 0xfe, 0xd7,
	0x03, 0xf9, 0x58, 0xb2, 0x97, 0x70, 0x1b, 0xe4, 0xca, 0xd8, 0x46, 0xb8, 0xf4, 0x35, 0x59, 0x92,
	0x81, 0x7e, 0x20, 0xd8, 0xe7, 0x39, 0xd9, 0x6a, 0xb1, 0xc5, 0xf4, 0x8d, 0xdf, 0x46, 0x70, 0xf5,
	0x8c, 0xbd, 0x83, 0xb1, 0xde, 0x37, 0x1b, 0x69, 0xb9, 0xa9, 0xb8, 0x95, 0x28, 0xed, 0x41, 0x86,
	0x8d, 0x60, 0xba, 0xa2, 0x1b, 0x77, 0xde, 0xb1, 0xae, 0xf2, 0x41, 0xf7, 0x11, 0x91, 0xfd, 0x88,
	0x20, 0x1e, 0x96, 0xf0, 0x76, 0x72, 0x39, 0xbf, 0x59, 0x74, 0xd9, 0xff, 0xff, 0x67, 0x66, 0x67,
	0x7f, 0xbd, 0x7c, 0x18, 0x64, 0x13, 0xd3, 0x27, 0xb2, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa8,
	0xf0, 0xf2, 0x2d, 0x40, 0x03, 0x00, 0x00,
}
