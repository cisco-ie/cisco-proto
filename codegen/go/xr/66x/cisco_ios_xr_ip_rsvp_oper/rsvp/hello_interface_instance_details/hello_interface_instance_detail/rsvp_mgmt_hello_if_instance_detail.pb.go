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
// source: rsvp_mgmt_hello_if_instance_detail.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_hello_interface_instance_details_hello_interface_instance_detail

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

type RsvpMgmtHelloIfInstanceDetail_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtHelloIfInstanceDetail_KEYS) Reset()         { *m = RsvpMgmtHelloIfInstanceDetail_KEYS{} }
func (m *RsvpMgmtHelloIfInstanceDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtHelloIfInstanceDetail_KEYS) ProtoMessage()    {}
func (*RsvpMgmtHelloIfInstanceDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f1aa8e9c8575c3, []int{0}
}

func (m *RsvpMgmtHelloIfInstanceDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtHelloIfInstanceDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtHelloIfInstanceDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail_KEYS.Merge(m, src)
}
func (m *RsvpMgmtHelloIfInstanceDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail_KEYS.Size(m)
}
func (m *RsvpMgmtHelloIfInstanceDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtHelloIfInstanceDetail_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *RsvpMgmtHelloIfInstanceDetail_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

type RsvpMgmtTimespec struct {
	Seconds              int32    `protobuf:"zigzag32,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanoseconds          int32    `protobuf:"zigzag32,2,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtTimespec) Reset()         { *m = RsvpMgmtTimespec{} }
func (m *RsvpMgmtTimespec) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtTimespec) ProtoMessage()    {}
func (*RsvpMgmtTimespec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f1aa8e9c8575c3, []int{1}
}

func (m *RsvpMgmtTimespec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtTimespec.Unmarshal(m, b)
}
func (m *RsvpMgmtTimespec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtTimespec.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtTimespec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtTimespec.Merge(m, src)
}
func (m *RsvpMgmtTimespec) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtTimespec.Size(m)
}
func (m *RsvpMgmtTimespec) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtTimespec.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtTimespec proto.InternalMessageInfo

func (m *RsvpMgmtTimespec) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *RsvpMgmtTimespec) GetNanoseconds() int32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

type RsvpMgmtHelloIfInstanceDetail struct {
	SourceAddressXr       string            `protobuf:"bytes,50,opt,name=source_address_xr,json=sourceAddressXr,proto3" json:"source_address_xr,omitempty"`
	DestinationAddressXr  string            `protobuf:"bytes,51,opt,name=destination_address_xr,json=destinationAddressXr,proto3" json:"destination_address_xr,omitempty"`
	HelloGlobalNeighborId string            `protobuf:"bytes,52,opt,name=hello_global_neighbor_id,json=helloGlobalNeighborId,proto3" json:"hello_global_neighbor_id,omitempty"`
	InstanceType          string            `protobuf:"bytes,53,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	LastMessageSentTime   *RsvpMgmtTimespec `protobuf:"bytes,54,opt,name=last_message_sent_time,json=lastMessageSentTime,proto3" json:"last_message_sent_time,omitempty"`
	HelloInterface        string            `protobuf:"bytes,55,opt,name=hello_interface,json=helloInterface,proto3" json:"hello_interface,omitempty"`
	SourceInstance        uint32            `protobuf:"varint,56,opt,name=source_instance,json=sourceInstance,proto3" json:"source_instance,omitempty"`
	DestinationInstance   uint32            `protobuf:"varint,57,opt,name=destination_instance,json=destinationInstance,proto3" json:"destination_instance,omitempty"`
	HelloMessagesSent     uint64            `protobuf:"varint,58,opt,name=hello_messages_sent,json=helloMessagesSent,proto3" json:"hello_messages_sent,omitempty"`
	HelloMessagesReceived uint64            `protobuf:"varint,59,opt,name=hello_messages_received,json=helloMessagesReceived,proto3" json:"hello_messages_received,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *RsvpMgmtHelloIfInstanceDetail) Reset()         { *m = RsvpMgmtHelloIfInstanceDetail{} }
func (m *RsvpMgmtHelloIfInstanceDetail) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtHelloIfInstanceDetail) ProtoMessage()    {}
func (*RsvpMgmtHelloIfInstanceDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f1aa8e9c8575c3, []int{2}
}

func (m *RsvpMgmtHelloIfInstanceDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail.Unmarshal(m, b)
}
func (m *RsvpMgmtHelloIfInstanceDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtHelloIfInstanceDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail.Merge(m, src)
}
func (m *RsvpMgmtHelloIfInstanceDetail) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail.Size(m)
}
func (m *RsvpMgmtHelloIfInstanceDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtHelloIfInstanceDetail proto.InternalMessageInfo

func (m *RsvpMgmtHelloIfInstanceDetail) GetSourceAddressXr() string {
	if m != nil {
		return m.SourceAddressXr
	}
	return ""
}

func (m *RsvpMgmtHelloIfInstanceDetail) GetDestinationAddressXr() string {
	if m != nil {
		return m.DestinationAddressXr
	}
	return ""
}

func (m *RsvpMgmtHelloIfInstanceDetail) GetHelloGlobalNeighborId() string {
	if m != nil {
		return m.HelloGlobalNeighborId
	}
	return ""
}

func (m *RsvpMgmtHelloIfInstanceDetail) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *RsvpMgmtHelloIfInstanceDetail) GetLastMessageSentTime() *RsvpMgmtTimespec {
	if m != nil {
		return m.LastMessageSentTime
	}
	return nil
}

func (m *RsvpMgmtHelloIfInstanceDetail) GetHelloInterface() string {
	if m != nil {
		return m.HelloInterface
	}
	return ""
}

func (m *RsvpMgmtHelloIfInstanceDetail) GetSourceInstance() uint32 {
	if m != nil {
		return m.SourceInstance
	}
	return 0
}

func (m *RsvpMgmtHelloIfInstanceDetail) GetDestinationInstance() uint32 {
	if m != nil {
		return m.DestinationInstance
	}
	return 0
}

func (m *RsvpMgmtHelloIfInstanceDetail) GetHelloMessagesSent() uint64 {
	if m != nil {
		return m.HelloMessagesSent
	}
	return 0
}

func (m *RsvpMgmtHelloIfInstanceDetail) GetHelloMessagesReceived() uint64 {
	if m != nil {
		return m.HelloMessagesReceived
	}
	return 0
}

func init() {
	proto.RegisterType((*RsvpMgmtHelloIfInstanceDetail_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.hello_interface_instance_details.hello_interface_instance_detail.rsvp_mgmt_hello_if_instance_detail_KEYS")
	proto.RegisterType((*RsvpMgmtTimespec)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.hello_interface_instance_details.hello_interface_instance_detail.rsvp_mgmt_timespec")
	proto.RegisterType((*RsvpMgmtHelloIfInstanceDetail)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.hello_interface_instance_details.hello_interface_instance_detail.rsvp_mgmt_hello_if_instance_detail")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_hello_if_instance_detail.proto", fileDescriptor_d8f1aa8e9c8575c3)
}

var fileDescriptor_d8f1aa8e9c8575c3 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0x96, 0xab, 0x0a, 0xc4, 0x94, 0xb4, 0xca, 0x86, 0x96, 0x3d, 0x5a, 0x41, 0xa8, 0x16, 0x07,
	0x23, 0xda, 0xd2, 0xf2, 0x38, 0x71, 0x40, 0x28, 0x42, 0x20, 0xe4, 0xf6, 0x50, 0x4e, 0x23, 0xc7,
	0x9e, 0xa6, 0x2b, 0xd9, 0xbb, 0xd6, 0xce, 0x52, 0xb5, 0x47, 0xfe, 0x0d, 0x3f, 0x92, 0x03, 0xf2,
	0x7a, 0x9d, 0xe6, 0x81, 0x94, 0x53, 0x6f, 0xc9, 0xf7, 0xd8, 0xfd, 0xbe, 0xd9, 0x31, 0x24, 0x96,
	0x6f, 0x1a, 0xac, 0x67, 0xb5, 0xc3, 0x6b, 0xaa, 0x2a, 0x83, 0xea, 0x0a, 0x95, 0x66, 0x97, 0xeb,
	0x82, 0xb0, 0x24, 0x97, 0xab, 0x2a, 0x6d, 0xac, 0x71, 0x46, 0x60, 0xa1, 0xb8, 0x30, 0xa8, 0x0c,
	0xe3, 0xad, 0x45, 0xd5, 0xa0, 0x77, 0x9a, 0x86, 0x6c, 0xda, 0xfe, 0x4a, 0x83, 0x5d, 0x3b, 0xb2,
	0x57, 0x79, 0x41, 0xab, 0xa7, 0xf0, 0x26, 0xc1, 0xf8, 0x77, 0x04, 0x87, 0x9b, 0xd3, 0xe0, 0xd7,
	0xcf, 0x3f, 0xcf, 0xc5, 0x4b, 0xd8, 0x65, 0xf3, 0xcb, 0x16, 0x84, 0x79, 0x59, 0x5a, 0x62, 0x96,
	0x51, 0x1c, 0x25, 0x4f, 0xb2, 0x41, 0x87, 0x7e, 0xea, 0x40, 0xf1, 0x1a, 0x46, 0x25, 0xb1, 0x53,
	0x3a, 0x77, 0xca, 0xe8, 0xb9, 0x76, 0xcb, 0x6b, 0xc5, 0x02, 0x15, 0x0c, 0xe3, 0x1f, 0x20, 0xee,
	0x23, 0x38, 0x55, 0x13, 0x37, 0x54, 0x08, 0x09, 0x8f, 0x99, 0x0a, 0xa3, 0xcb, 0xee, 0x9a, 0x61,
	0xd6, 0xff, 0x15, 0x31, 0xec, 0xe8, 0x5c, 0x9b, 0x9e, 0xdd, 0xf2, 0xec, 0x22, 0x34, 0xfe, 0xbb,
	0x0d, 0xe3, 0xcd, 0xad, 0xc4, 0x2b, 0x18, 0x2e, 0x17, 0xc2, 0x5b, 0x2b, 0x8f, 0x7c, 0xce, 0xbd,
	0xa5, 0x4e, 0x97, 0x56, 0x9c, 0xc0, 0xc1, 0x7f, 0x5a, 0xb5, 0x86, 0x63, 0x6f, 0x78, 0xb6, 0x5e,
	0xec, 0xd2, 0x8a, 0x33, 0x90, 0xdd, 0xed, 0xb3, 0xca, 0x4c, 0xf3, 0x0a, 0x35, 0xa9, 0xd9, 0xf5,
	0xd4, 0x58, 0x54, 0xa5, 0x3c, 0xf1, 0xbe, 0x7d, 0xcf, 0x7f, 0xf1, 0xf4, 0xf7, 0xc0, 0x4e, 0x4a,
	0xf1, 0x02, 0x06, 0xf3, 0xb4, 0xee, 0xae, 0x21, 0xf9, 0xd6, 0xab, 0x9f, 0xf6, 0xe0, 0xc5, 0x5d,
	0x43, 0xe2, 0x4f, 0x04, 0x07, 0x55, 0xce, 0x0e, 0x6b, 0x62, 0xce, 0x67, 0x84, 0x4c, 0xba, 0x9b,
	0xa0, 0x3c, 0x8d, 0xa3, 0x64, 0xe7, 0x88, 0xd3, 0x07, 0xde, 0x9f, 0x74, 0xfd, 0xe1, 0xb2, 0x51,
	0x1b, 0xe9, 0x5b, 0x97, 0xe8, 0x9c, 0xb4, 0xbb, 0x50, 0x35, 0x89, 0x43, 0xd8, 0x5b, 0x39, 0x4a,
	0x9e, 0xf9, 0x46, 0xbb, 0x1e, 0x9e, 0xf4, 0x68, 0x2b, 0x0c, 0x6f, 0xd2, 0x5f, 0x25, 0xdf, 0xc5,
	0x51, 0x32, 0xc8, 0xc2, 0xee, 0x4d, 0x02, 0x2a, 0xde, 0xc0, 0xe2, 0xc8, 0xef, 0xd5, 0xef, 0xbd,
	0x7a, 0x71, 0x05, 0xe7, 0x96, 0x14, 0x46, 0x5d, 0x88, 0x30, 0x2f, 0xf6, 0x03, 0x93, 0x1f, 0xe2,
	0x28, 0xd9, 0xce, 0x86, 0x9e, 0x0a, 0xb9, 0xb9, 0x0d, 0x2e, 0x4e, 0xe1, 0xf9, 0x8a, 0xde, 0x52,
	0x41, 0xea, 0x86, 0x4a, 0xf9, 0xd1, 0x7b, 0xf6, 0x97, 0x3c, 0x59, 0x20, 0xa7, 0x8f, 0xfc, 0xc7,
	0x7b, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x69, 0x29, 0x0a, 0xe8, 0x03, 0x00, 0x00,
}