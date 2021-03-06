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
// source: bgp_convergence_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_default_vrf_afs_af_convergence

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

type BgpConvergenceBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpConvergenceBag_KEYS) Reset()         { *m = BgpConvergenceBag_KEYS{} }
func (m *BgpConvergenceBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpConvergenceBag_KEYS) ProtoMessage()    {}
func (*BgpConvergenceBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d12437cce6d3364d, []int{0}
}

func (m *BgpConvergenceBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpConvergenceBag_KEYS.Unmarshal(m, b)
}
func (m *BgpConvergenceBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpConvergenceBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpConvergenceBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpConvergenceBag_KEYS.Merge(m, src)
}
func (m *BgpConvergenceBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpConvergenceBag_KEYS.Size(m)
}
func (m *BgpConvergenceBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpConvergenceBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpConvergenceBag_KEYS proto.InternalMessageInfo

func (m *BgpConvergenceBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpConvergenceBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

type BgpConvergenceBag struct {
	AfName               string   `protobuf:"bytes,50,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	HasConverged         bool     `protobuf:"varint,51,opt,name=has_converged,json=hasConverged,proto3" json:"has_converged,omitempty"`
	AreWriteQueuesEmpty  bool     `protobuf:"varint,52,opt,name=are_write_queues_empty,json=areWriteQueuesEmpty,proto3" json:"are_write_queues_empty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpConvergenceBag) Reset()         { *m = BgpConvergenceBag{} }
func (m *BgpConvergenceBag) String() string { return proto.CompactTextString(m) }
func (*BgpConvergenceBag) ProtoMessage()    {}
func (*BgpConvergenceBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d12437cce6d3364d, []int{1}
}

func (m *BgpConvergenceBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpConvergenceBag.Unmarshal(m, b)
}
func (m *BgpConvergenceBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpConvergenceBag.Marshal(b, m, deterministic)
}
func (m *BgpConvergenceBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpConvergenceBag.Merge(m, src)
}
func (m *BgpConvergenceBag) XXX_Size() int {
	return xxx_messageInfo_BgpConvergenceBag.Size(m)
}
func (m *BgpConvergenceBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpConvergenceBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpConvergenceBag proto.InternalMessageInfo

func (m *BgpConvergenceBag) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpConvergenceBag) GetHasConverged() bool {
	if m != nil {
		return m.HasConverged
	}
	return false
}

func (m *BgpConvergenceBag) GetAreWriteQueuesEmpty() bool {
	if m != nil {
		return m.AreWriteQueuesEmpty
	}
	return false
}

func init() {
	proto.RegisterType((*BgpConvergenceBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.afs.af.convergence.bgp_convergence_bag_KEYS")
	proto.RegisterType((*BgpConvergenceBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.afs.af.convergence.bgp_convergence_bag")
}

func init() { proto.RegisterFile("bgp_convergence_bag.proto", fileDescriptor_d12437cce6d3364d) }

var fileDescriptor_d12437cce6d3364d = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x15, 0x86, 0x02, 0x56, 0x59, 0x52, 0x09, 0xc2, 0x56, 0x95, 0xa5, 0x93, 0x07, 0xd2,
	0x37, 0x40, 0x9d, 0x90, 0x90, 0x28, 0x03, 0x30, 0x9d, 0xce, 0xce, 0x39, 0xb5, 0x44, 0x6c, 0x63,
	0x3b, 0x81, 0xbe, 0x01, 0x8f, 0x8d, 0x9c, 0xaa, 0x56, 0x86, 0x4e, 0x3e, 0x7d, 0xbf, 0xef, 0x8f,
	0x64, 0x76, 0x2f, 0x5a, 0x07, 0xd2, 0x9a, 0x81, 0x7c, 0x4b, 0x46, 0x12, 0x08, 0x6c, 0xb9, 0xf3,
	0x36, 0xda, 0x12, 0xa5, 0x0e, 0xd2, 0x82, 0xb6, 0x01, 0x7e, 0x3d, 0x68, 0x37, 0x6c, 0x20, 0x99,
	0xad, 0x23, 0xcf, 0x45, 0xeb, 0xb8, 0x36, 0x21, 0xa2, 0x91, 0x14, 0xf2, 0x95, 0x0f, 0x48, 0x4f,
	0x23, 0x0e, 0xbc, 0x21, 0x85, 0xfd, 0x57, 0x84, 0xc1, 0x2b, 0x8e, 0x2a, 0x70, 0x54, 0x7c, 0x32,
	0xb6, 0xfa, 0x60, 0xd5, 0x99, 0x7d, 0x78, 0xde, 0x7e, 0xbe, 0x95, 0x0f, 0xec, 0x26, 0xd7, 0x19,
	0xec, 0xa8, 0x2a, 0x96, 0xc5, 0xfa, 0x7a, 0x37, 0x3f, 0x89, 0x2f, 0xd8, 0x51, 0x79, 0xc7, 0x2e,
	0x51, 0x1d, 0xf1, 0xc5, 0x88, 0x67, 0xa8, 0x12, 0x58, 0xfd, 0x15, 0x6c, 0x71, 0xa6, 0x7a, 0x1a,
	0x78, 0x9c, 0x06, 0xd2, 0xdc, 0x1e, 0x43, 0xf6, 0x37, 0x55, 0xbd, 0x2c, 0xd6, 0x57, 0xbb, 0xf9,
	0x1e, 0xc3, 0xd3, 0x49, 0x2b, 0x6b, 0x76, 0x8b, 0x9e, 0xe0, 0xc7, 0xeb, 0x48, 0xf0, 0xdd, 0x53,
	0x4f, 0x01, 0xa8, 0x73, 0xf1, 0x50, 0x6d, 0x46, 0xf7, 0x02, 0x3d, 0xbd, 0x27, 0xf8, 0x3a, 0xb2,
	0x6d, 0x42, 0x62, 0x36, 0x7e, 0x67, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x88, 0x72, 0xb3, 0xd1,
	0x6b, 0x01, 0x00, 0x00,
}
