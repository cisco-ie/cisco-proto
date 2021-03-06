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
// source: ipsla_history_bag.proto

package cisco_ios_xr_man_ipsla_oper_ipsla_operation_data_operations_operation_history_path_lifes_life_buckets_bucket_samples_sample

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

type IpslaHistoryBag_KEYS struct {
	OperationId          int32    `protobuf:"varint,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	LifeIndex            int32    `protobuf:"varint,2,opt,name=life_index,json=lifeIndex,proto3" json:"life_index,omitempty"`
	BucketIndex          int32    `protobuf:"varint,3,opt,name=bucket_index,json=bucketIndex,proto3" json:"bucket_index,omitempty"`
	SampleIndex          int32    `protobuf:"varint,4,opt,name=sample_index,json=sampleIndex,proto3" json:"sample_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpslaHistoryBag_KEYS) Reset()         { *m = IpslaHistoryBag_KEYS{} }
func (m *IpslaHistoryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpslaHistoryBag_KEYS) ProtoMessage()    {}
func (*IpslaHistoryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_199bbac23e97978b, []int{0}
}

func (m *IpslaHistoryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaHistoryBag_KEYS.Unmarshal(m, b)
}
func (m *IpslaHistoryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaHistoryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IpslaHistoryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaHistoryBag_KEYS.Merge(m, src)
}
func (m *IpslaHistoryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpslaHistoryBag_KEYS.Size(m)
}
func (m *IpslaHistoryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaHistoryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaHistoryBag_KEYS proto.InternalMessageInfo

func (m *IpslaHistoryBag_KEYS) GetOperationId() int32 {
	if m != nil {
		return m.OperationId
	}
	return 0
}

func (m *IpslaHistoryBag_KEYS) GetLifeIndex() int32 {
	if m != nil {
		return m.LifeIndex
	}
	return 0
}

func (m *IpslaHistoryBag_KEYS) GetBucketIndex() int32 {
	if m != nil {
		return m.BucketIndex
	}
	return 0
}

func (m *IpslaHistoryBag_KEYS) GetSampleIndex() int32 {
	if m != nil {
		return m.SampleIndex
	}
	return 0
}

type IpslaIpv4PrefixT struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	MaskLength           uint32   `protobuf:"varint,2,opt,name=mask_length,json=maskLength,proto3" json:"mask_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpslaIpv4PrefixT) Reset()         { *m = IpslaIpv4PrefixT{} }
func (m *IpslaIpv4PrefixT) String() string { return proto.CompactTextString(m) }
func (*IpslaIpv4PrefixT) ProtoMessage()    {}
func (*IpslaIpv4PrefixT) Descriptor() ([]byte, []int) {
	return fileDescriptor_199bbac23e97978b, []int{1}
}

func (m *IpslaIpv4PrefixT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaIpv4PrefixT.Unmarshal(m, b)
}
func (m *IpslaIpv4PrefixT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaIpv4PrefixT.Marshal(b, m, deterministic)
}
func (m *IpslaIpv4PrefixT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaIpv4PrefixT.Merge(m, src)
}
func (m *IpslaIpv4PrefixT) XXX_Size() int {
	return xxx_messageInfo_IpslaIpv4PrefixT.Size(m)
}
func (m *IpslaIpv4PrefixT) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaIpv4PrefixT.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaIpv4PrefixT proto.InternalMessageInfo

func (m *IpslaIpv4PrefixT) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IpslaIpv4PrefixT) GetMaskLength() uint32 {
	if m != nil {
		return m.MaskLength
	}
	return 0
}

type IpslaTunnelIdT struct {
	TunnelId             uint32   `protobuf:"varint,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpslaTunnelIdT) Reset()         { *m = IpslaTunnelIdT{} }
func (m *IpslaTunnelIdT) String() string { return proto.CompactTextString(m) }
func (*IpslaTunnelIdT) ProtoMessage()    {}
func (*IpslaTunnelIdT) Descriptor() ([]byte, []int) {
	return fileDescriptor_199bbac23e97978b, []int{2}
}

func (m *IpslaTunnelIdT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaTunnelIdT.Unmarshal(m, b)
}
func (m *IpslaTunnelIdT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaTunnelIdT.Marshal(b, m, deterministic)
}
func (m *IpslaTunnelIdT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaTunnelIdT.Merge(m, src)
}
func (m *IpslaTunnelIdT) XXX_Size() int {
	return xxx_messageInfo_IpslaTunnelIdT.Size(m)
}
func (m *IpslaTunnelIdT) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaTunnelIdT.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaTunnelIdT proto.InternalMessageInfo

func (m *IpslaTunnelIdT) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

type IpslaIpv4PwT struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	VirtualCircuitId     uint32   `protobuf:"varint,2,opt,name=virtual_circuit_id,json=virtualCircuitId,proto3" json:"virtual_circuit_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpslaIpv4PwT) Reset()         { *m = IpslaIpv4PwT{} }
func (m *IpslaIpv4PwT) String() string { return proto.CompactTextString(m) }
func (*IpslaIpv4PwT) ProtoMessage()    {}
func (*IpslaIpv4PwT) Descriptor() ([]byte, []int) {
	return fileDescriptor_199bbac23e97978b, []int{3}
}

func (m *IpslaIpv4PwT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaIpv4PwT.Unmarshal(m, b)
}
func (m *IpslaIpv4PwT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaIpv4PwT.Marshal(b, m, deterministic)
}
func (m *IpslaIpv4PwT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaIpv4PwT.Merge(m, src)
}
func (m *IpslaIpv4PwT) XXX_Size() int {
	return xxx_messageInfo_IpslaIpv4PwT.Size(m)
}
func (m *IpslaIpv4PwT) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaIpv4PwT.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaIpv4PwT proto.InternalMessageInfo

func (m *IpslaIpv4PwT) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IpslaIpv4PwT) GetVirtualCircuitId() uint32 {
	if m != nil {
		return m.VirtualCircuitId
	}
	return 0
}

type IpslaTargetUnion struct {
	TargetType           string            `protobuf:"bytes,1,opt,name=target_type,json=targetType,proto3" json:"target_type,omitempty"`
	Ipv4AddressTarget    string            `protobuf:"bytes,2,opt,name=ipv4_address_target,json=ipv4AddressTarget,proto3" json:"ipv4_address_target,omitempty"`
	Ipv4PrefixTarget     *IpslaIpv4PrefixT `protobuf:"bytes,3,opt,name=ipv4_prefix_target,json=ipv4PrefixTarget,proto3" json:"ipv4_prefix_target,omitempty"`
	TunnelIdTarget       *IpslaTunnelIdT   `protobuf:"bytes,4,opt,name=tunnel_id_target,json=tunnelIdTarget,proto3" json:"tunnel_id_target,omitempty"`
	Ipv4PseudowireTarget *IpslaIpv4PwT     `protobuf:"bytes,5,opt,name=ipv4_pseudowire_target,json=ipv4PseudowireTarget,proto3" json:"ipv4_pseudowire_target,omitempty"`
	Ipv6AddressTarget    string            `protobuf:"bytes,6,opt,name=ipv6_address_target,json=ipv6AddressTarget,proto3" json:"ipv6_address_target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IpslaTargetUnion) Reset()         { *m = IpslaTargetUnion{} }
func (m *IpslaTargetUnion) String() string { return proto.CompactTextString(m) }
func (*IpslaTargetUnion) ProtoMessage()    {}
func (*IpslaTargetUnion) Descriptor() ([]byte, []int) {
	return fileDescriptor_199bbac23e97978b, []int{4}
}

func (m *IpslaTargetUnion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaTargetUnion.Unmarshal(m, b)
}
func (m *IpslaTargetUnion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaTargetUnion.Marshal(b, m, deterministic)
}
func (m *IpslaTargetUnion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaTargetUnion.Merge(m, src)
}
func (m *IpslaTargetUnion) XXX_Size() int {
	return xxx_messageInfo_IpslaTargetUnion.Size(m)
}
func (m *IpslaTargetUnion) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaTargetUnion.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaTargetUnion proto.InternalMessageInfo

func (m *IpslaTargetUnion) GetTargetType() string {
	if m != nil {
		return m.TargetType
	}
	return ""
}

func (m *IpslaTargetUnion) GetIpv4AddressTarget() string {
	if m != nil {
		return m.Ipv4AddressTarget
	}
	return ""
}

func (m *IpslaTargetUnion) GetIpv4PrefixTarget() *IpslaIpv4PrefixT {
	if m != nil {
		return m.Ipv4PrefixTarget
	}
	return nil
}

func (m *IpslaTargetUnion) GetTunnelIdTarget() *IpslaTunnelIdT {
	if m != nil {
		return m.TunnelIdTarget
	}
	return nil
}

func (m *IpslaTargetUnion) GetIpv4PseudowireTarget() *IpslaIpv4PwT {
	if m != nil {
		return m.Ipv4PseudowireTarget
	}
	return nil
}

func (m *IpslaTargetUnion) GetIpv6AddressTarget() string {
	if m != nil {
		return m.Ipv6AddressTarget
	}
	return ""
}

type IpslaHistoryBag struct {
	StartTime            uint64            `protobuf:"varint,50,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	ResponseTime         uint32            `protobuf:"varint,51,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	ReturnCode           string            `protobuf:"bytes,52,opt,name=return_code,json=returnCode,proto3" json:"return_code,omitempty"`
	TargetAddress        *IpslaTargetUnion `protobuf:"bytes,53,opt,name=target_address,json=targetAddress,proto3" json:"target_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IpslaHistoryBag) Reset()         { *m = IpslaHistoryBag{} }
func (m *IpslaHistoryBag) String() string { return proto.CompactTextString(m) }
func (*IpslaHistoryBag) ProtoMessage()    {}
func (*IpslaHistoryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_199bbac23e97978b, []int{5}
}

func (m *IpslaHistoryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaHistoryBag.Unmarshal(m, b)
}
func (m *IpslaHistoryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaHistoryBag.Marshal(b, m, deterministic)
}
func (m *IpslaHistoryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaHistoryBag.Merge(m, src)
}
func (m *IpslaHistoryBag) XXX_Size() int {
	return xxx_messageInfo_IpslaHistoryBag.Size(m)
}
func (m *IpslaHistoryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaHistoryBag.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaHistoryBag proto.InternalMessageInfo

func (m *IpslaHistoryBag) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *IpslaHistoryBag) GetResponseTime() uint32 {
	if m != nil {
		return m.ResponseTime
	}
	return 0
}

func (m *IpslaHistoryBag) GetReturnCode() string {
	if m != nil {
		return m.ReturnCode
	}
	return ""
}

func (m *IpslaHistoryBag) GetTargetAddress() *IpslaTargetUnion {
	if m != nil {
		return m.TargetAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*IpslaHistoryBag_KEYS)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.operation_data.operations.operation.history.path.lifes.life.buckets.bucket.samples.sample.ipsla_history_bag_KEYS")
	proto.RegisterType((*IpslaIpv4PrefixT)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.operation_data.operations.operation.history.path.lifes.life.buckets.bucket.samples.sample.ipsla_ipv4_prefix_t")
	proto.RegisterType((*IpslaTunnelIdT)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.operation_data.operations.operation.history.path.lifes.life.buckets.bucket.samples.sample.ipsla_tunnel_id_t")
	proto.RegisterType((*IpslaIpv4PwT)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.operation_data.operations.operation.history.path.lifes.life.buckets.bucket.samples.sample.ipsla_ipv4_pw_t")
	proto.RegisterType((*IpslaTargetUnion)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.operation_data.operations.operation.history.path.lifes.life.buckets.bucket.samples.sample.ipsla_target_union")
	proto.RegisterType((*IpslaHistoryBag)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.operation_data.operations.operation.history.path.lifes.life.buckets.bucket.samples.sample.ipsla_history_bag")
}

func init() { proto.RegisterFile("ipsla_history_bag.proto", fileDescriptor_199bbac23e97978b) }

var fileDescriptor_199bbac23e97978b = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcf, 0x6e, 0x13, 0x3f,
	0x10, 0xd6, 0xf6, 0x97, 0xf4, 0x47, 0x9c, 0xa6, 0xa4, 0x06, 0x41, 0x24, 0x84, 0x5a, 0x96, 0x4b,
	0x0f, 0x68, 0x85, 0xda, 0xd0, 0x3b, 0xaa, 0x38, 0x44, 0x70, 0xa8, 0x96, 0x5c, 0x7a, 0xb2, 0x9c,
	0xf5, 0x34, 0xb1, 0xba, 0x59, 0xaf, 0x6c, 0x6f, 0xfe, 0x88, 0x57, 0x40, 0xc0, 0x23, 0xf0, 0x0c,
	0x5c, 0x10, 0x27, 0x0e, 0x3c, 0x05, 0x6f, 0x83, 0xec, 0xf1, 0xa6, 0xcb, 0x1f, 0x71, 0x25, 0x97,
	0xac, 0xe7, 0x9b, 0x2f, 0x33, 0xdf, 0x8c, 0x67, 0x76, 0xc9, 0x7d, 0x59, 0x9a, 0x9c, 0xb3, 0x99,
	0x34, 0x56, 0xe9, 0x35, 0x9b, 0xf0, 0x69, 0x52, 0x6a, 0x65, 0x15, 0x7d, 0x93, 0x49, 0x93, 0x29,
	0x26, 0x95, 0x61, 0x2b, 0xcd, 0xe6, 0xbc, 0x60, 0xc8, 0x54, 0x25, 0xe8, 0xc4, 0x1f, 0x13, 0x77,
	0xe4, 0x56, 0xaa, 0x82, 0x09, 0x6e, 0x1b, 0xa6, 0xb9, 0x39, 0x26, 0x21, 0x70, 0x52, 0x72, 0x3b,
	0x4b, 0x72, 0x79, 0x05, 0xc6, 0xff, 0x26, 0x93, 0x2a, 0xbb, 0x06, 0x6b, 0xc2, 0x33, 0x31, 0x7c,
	0x5e, 0xe6, 0x60, 0xc2, 0x33, 0xfe, 0x18, 0x91, 0x7b, 0xbf, 0x09, 0x63, 0x2f, 0x5f, 0x5c, 0xbe,
	0xa6, 0x8f, 0xc8, 0xde, 0x4d, 0x5e, 0x29, 0x06, 0xd1, 0x51, 0x74, 0xdc, 0x4e, 0xbb, 0x1b, 0x6c,
	0x24, 0xe8, 0x43, 0x42, 0x5c, 0x0e, 0x26, 0x0b, 0x01, 0xab, 0xc1, 0x8e, 0x27, 0x74, 0x1c, 0x32,
	0x72, 0x80, 0x8b, 0x80, 0x59, 0x03, 0xe1, 0x3f, 0x8c, 0x80, 0xd8, 0x86, 0x82, 0x4a, 0x02, 0xa5,
	0x85, 0x14, 0xc4, 0x3c, 0x25, 0xbe, 0x20, 0x77, 0x50, 0xa1, 0x2c, 0x17, 0x43, 0x56, 0x6a, 0xb8,
	0x92, 0x2b, 0x66, 0xe9, 0x80, 0xfc, 0xcf, 0x85, 0xd0, 0x60, 0x8c, 0x57, 0xd6, 0x49, 0x6b, 0x93,
	0x1e, 0x92, 0xee, 0x9c, 0x9b, 0x6b, 0x96, 0x43, 0x31, 0xb5, 0x33, 0x2f, 0xab, 0x97, 0x12, 0x07,
	0xbd, 0xf2, 0x48, 0xfc, 0x94, 0x1c, 0x60, 0x44, 0x5b, 0x15, 0x05, 0xe4, 0x4c, 0x0a, 0x66, 0xe9,
	0x03, 0xd2, 0xd9, 0x98, 0x3e, 0x62, 0x2f, 0xbd, 0x85, 0xc0, 0x48, 0xc4, 0x97, 0xe4, 0x76, 0x53,
	0xc3, 0xf2, 0xaf, 0xf9, 0x9f, 0x10, 0xba, 0x90, 0xda, 0x56, 0x3c, 0x67, 0x99, 0xd4, 0x59, 0x25,
	0xad, 0x0b, 0x89, 0x32, 0xfa, 0xc1, 0x73, 0x8e, 0x8e, 0x91, 0x88, 0xbf, 0xb7, 0x09, 0x0d, 0x6a,
	0xb8, 0x9e, 0x82, 0x65, 0x55, 0x21, 0x55, 0xe1, 0x8a, 0x08, 0xb6, 0x5d, 0x97, 0x10, 0x52, 0x10,
	0x84, 0xc6, 0xeb, 0x12, 0x68, 0xe2, 0xda, 0xb2, 0x18, 0xb2, 0x90, 0x35, 0xfc, 0xdb, 0xa7, 0xe9,
	0xa4, 0x07, 0xce, 0xf5, 0x1c, 0x3d, 0x63, 0xef, 0xa0, 0x5f, 0x23, 0x97, 0xa7, 0xd1, 0x41, 0xe4,
	0xbb, 0x3b, 0xe9, 0x9e, 0x7c, 0x88, 0x92, 0x7f, 0x38, 0x85, 0xc9, 0x1f, 0xee, 0x37, 0xed, 0x3b,
	0xf3, 0xc2, 0x5b, 0xa1, 0x82, 0x2f, 0x11, 0xe9, 0x37, 0x6e, 0x0c, 0xf5, 0xb7, 0xbc, 0xfe, 0x77,
	0xdb, 0xa0, 0xbf, 0xa1, 0x2d, 0xdd, 0xaf, 0x47, 0x27, 0x68, 0xff, 0xe6, 0xf7, 0xcc, 0xd5, 0x67,
	0xa0, 0x12, 0x6a, 0x29, 0x35, 0xd4, 0x15, 0xb4, 0x7d, 0x05, 0x6f, 0xb7, 0xe7, 0x06, 0x96, 0xcc,
	0xa6, 0x77, 0x7d, 0xf7, 0x37, 0x5a, 0x43, 0x15, 0x38, 0x73, 0x67, 0xbf, 0xce, 0xdc, 0xee, 0x66,
	0xe6, 0xce, 0x7e, 0x9a, 0xb9, 0xf8, 0xd3, 0x4e, 0xbd, 0x69, 0x8d, 0xb7, 0x8b, 0x7b, 0x6b, 0x18,
	0xcb, 0xb5, 0x65, 0x56, 0xce, 0x61, 0x70, 0x72, 0x14, 0x1d, 0xb7, 0xd2, 0x8e, 0x47, 0xc6, 0x72,
	0x0e, 0xf4, 0x31, 0xe9, 0x69, 0x30, 0xa5, 0x2a, 0x0c, 0x20, 0xe3, 0xd4, 0x6f, 0xce, 0x5e, 0x0d,
	0x7a, 0xd2, 0x21, 0xe9, 0x6a, 0xb0, 0x95, 0x2e, 0x58, 0xa6, 0x04, 0x0c, 0x86, 0xb8, 0x1e, 0x08,
	0x9d, 0x2b, 0x01, 0xf4, 0x73, 0x44, 0xf6, 0xc3, 0x02, 0xd5, 0x6b, 0xfa, 0xcc, 0x37, 0xfa, 0xfd,
	0x56, 0x8c, 0x4a, 0x63, 0xd5, 0xd3, 0x1e, 0x5a, 0xa1, 0x75, 0x93, 0x5d, 0xff, 0x59, 0x38, 0xfd,
	0x11, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x5b, 0xbf, 0xa0, 0x31, 0x06, 0x00, 0x00,
}
