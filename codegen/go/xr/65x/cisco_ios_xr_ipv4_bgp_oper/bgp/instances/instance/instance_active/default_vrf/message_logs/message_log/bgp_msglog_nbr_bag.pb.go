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
// source: bgp_msglog_nbr_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_message_logs_message_log

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

type BgpMsglogNbrBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,2,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	Direction            uint32   `protobuf:"varint,3,opt,name=direction,proto3" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpMsglogNbrBag_KEYS) Reset()         { *m = BgpMsglogNbrBag_KEYS{} }
func (m *BgpMsglogNbrBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpMsglogNbrBag_KEYS) ProtoMessage()    {}
func (*BgpMsglogNbrBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_365190dc70d53b1e, []int{0}
}

func (m *BgpMsglogNbrBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpMsglogNbrBag_KEYS.Unmarshal(m, b)
}
func (m *BgpMsglogNbrBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpMsglogNbrBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpMsglogNbrBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpMsglogNbrBag_KEYS.Merge(m, src)
}
func (m *BgpMsglogNbrBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpMsglogNbrBag_KEYS.Size(m)
}
func (m *BgpMsglogNbrBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpMsglogNbrBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpMsglogNbrBag_KEYS proto.InternalMessageInfo

func (m *BgpMsglogNbrBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpMsglogNbrBag_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *BgpMsglogNbrBag_KEYS) GetDirection() uint32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type BgpTimespec struct {
	Seconds              uint32   `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanoseconds          uint32   `protobuf:"varint,2,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpTimespec) Reset()         { *m = BgpTimespec{} }
func (m *BgpTimespec) String() string { return proto.CompactTextString(m) }
func (*BgpTimespec) ProtoMessage()    {}
func (*BgpTimespec) Descriptor() ([]byte, []int) {
	return fileDescriptor_365190dc70d53b1e, []int{1}
}

func (m *BgpTimespec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpTimespec.Unmarshal(m, b)
}
func (m *BgpTimespec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpTimespec.Marshal(b, m, deterministic)
}
func (m *BgpTimespec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpTimespec.Merge(m, src)
}
func (m *BgpTimespec) XXX_Size() int {
	return xxx_messageInfo_BgpTimespec.Size(m)
}
func (m *BgpTimespec) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpTimespec.DiscardUnknown(m)
}

var xxx_messageInfo_BgpTimespec proto.InternalMessageInfo

func (m *BgpTimespec) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *BgpTimespec) GetNanoseconds() uint32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

type BgpNbrMsg_ struct {
	MessageTypeReceived     uint32       `protobuf:"varint,1,opt,name=message_type_received,json=messageTypeReceived,proto3" json:"message_type_received,omitempty"`
	TotalLoggedMessageCount uint32       `protobuf:"varint,2,opt,name=total_logged_message_count,json=totalLoggedMessageCount,proto3" json:"total_logged_message_count,omitempty"`
	MessageTimestamp        *BgpTimespec `protobuf:"bytes,3,opt,name=message_timestamp,json=messageTimestamp,proto3" json:"message_timestamp,omitempty"`
	MessageDataLength       uint32       `protobuf:"varint,4,opt,name=message_data_length,json=messageDataLength,proto3" json:"message_data_length,omitempty"`
	LoggedMessageData       []uint32     `protobuf:"varint,5,rep,packed,name=logged_message_data,json=loggedMessageData,proto3" json:"logged_message_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}     `json:"-"`
	XXX_unrecognized        []byte       `json:"-"`
	XXX_sizecache           int32        `json:"-"`
}

func (m *BgpNbrMsg_) Reset()         { *m = BgpNbrMsg_{} }
func (m *BgpNbrMsg_) String() string { return proto.CompactTextString(m) }
func (*BgpNbrMsg_) ProtoMessage()    {}
func (*BgpNbrMsg_) Descriptor() ([]byte, []int) {
	return fileDescriptor_365190dc70d53b1e, []int{2}
}

func (m *BgpNbrMsg_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNbrMsg_.Unmarshal(m, b)
}
func (m *BgpNbrMsg_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNbrMsg_.Marshal(b, m, deterministic)
}
func (m *BgpNbrMsg_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNbrMsg_.Merge(m, src)
}
func (m *BgpNbrMsg_) XXX_Size() int {
	return xxx_messageInfo_BgpNbrMsg_.Size(m)
}
func (m *BgpNbrMsg_) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNbrMsg_.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNbrMsg_ proto.InternalMessageInfo

func (m *BgpNbrMsg_) GetMessageTypeReceived() uint32 {
	if m != nil {
		return m.MessageTypeReceived
	}
	return 0
}

func (m *BgpNbrMsg_) GetTotalLoggedMessageCount() uint32 {
	if m != nil {
		return m.TotalLoggedMessageCount
	}
	return 0
}

func (m *BgpNbrMsg_) GetMessageTimestamp() *BgpTimespec {
	if m != nil {
		return m.MessageTimestamp
	}
	return nil
}

func (m *BgpNbrMsg_) GetMessageDataLength() uint32 {
	if m != nil {
		return m.MessageDataLength
	}
	return 0
}

func (m *BgpNbrMsg_) GetLoggedMessageData() []uint32 {
	if m != nil {
		return m.LoggedMessageData
	}
	return nil
}

type BgpMsglogNbrBag struct {
	NeighborMessage      []*BgpNbrMsg_ `protobuf:"bytes,50,rep,name=neighbor_message,json=neighborMessage,proto3" json:"neighbor_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BgpMsglogNbrBag) Reset()         { *m = BgpMsglogNbrBag{} }
func (m *BgpMsglogNbrBag) String() string { return proto.CompactTextString(m) }
func (*BgpMsglogNbrBag) ProtoMessage()    {}
func (*BgpMsglogNbrBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_365190dc70d53b1e, []int{3}
}

func (m *BgpMsglogNbrBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpMsglogNbrBag.Unmarshal(m, b)
}
func (m *BgpMsglogNbrBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpMsglogNbrBag.Marshal(b, m, deterministic)
}
func (m *BgpMsglogNbrBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpMsglogNbrBag.Merge(m, src)
}
func (m *BgpMsglogNbrBag) XXX_Size() int {
	return xxx_messageInfo_BgpMsglogNbrBag.Size(m)
}
func (m *BgpMsglogNbrBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpMsglogNbrBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpMsglogNbrBag proto.InternalMessageInfo

func (m *BgpMsglogNbrBag) GetNeighborMessage() []*BgpNbrMsg_ {
	if m != nil {
		return m.NeighborMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpMsglogNbrBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.message_logs.message_log.bgp_msglog_nbr_bag_KEYS")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.message_logs.message_log.bgp_timespec")
	proto.RegisterType((*BgpNbrMsg_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.message_logs.message_log.bgp_nbr_msg_")
	proto.RegisterType((*BgpMsglogNbrBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.message_logs.message_log.bgp_msglog_nbr_bag")
}

func init() { proto.RegisterFile("bgp_msglog_nbr_bag.proto", fileDescriptor_365190dc70d53b1e) }

var fileDescriptor_365190dc70d53b1e = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x3d, 0x6f, 0x14, 0x31,
	0x10, 0xd5, 0xe6, 0xf8, 0x50, 0x9c, 0x9c, 0x08, 0x46, 0x28, 0x16, 0xa2, 0x58, 0x1d, 0x4d, 0x68,
	0x5c, 0x1c, 0x74, 0x54, 0x08, 0x68, 0x20, 0x50, 0x2c, 0x34, 0x54, 0x23, 0xaf, 0x3d, 0xe7, 0x58,
	0xda, 0xb5, 0x2d, 0xdb, 0x39, 0x91, 0x9a, 0x82, 0x9f, 0x81, 0xf8, 0x13, 0xfc, 0x3e, 0x64, 0xdf,
	0xfa, 0xb2, 0x22, 0xfd, 0x75, 0xf6, 0xbc, 0xe7, 0x99, 0xf7, 0xde, 0xce, 0x12, 0xd6, 0x6b, 0x0f,
	0x63, 0xd4, 0x83, 0xd3, 0x60, 0xfb, 0x00, 0xbd, 0xd0, 0xdc, 0x07, 0x97, 0x1c, 0xdd, 0x48, 0x13,
	0xa5, 0x03, 0xe3, 0x22, 0xfc, 0x08, 0x60, 0xfc, 0xf6, 0x35, 0x64, 0xae, 0xf3, 0x18, 0x78, 0xaf,
	0x3d, 0x37, 0x36, 0x26, 0x61, 0x25, 0xc6, 0xfd, 0x69, 0x7f, 0x00, 0x21, 0x93, 0xd9, 0x22, 0x57,
	0xb8, 0x11, 0xd7, 0x43, 0x82, 0x6d, 0xd8, 0xf0, 0x11, 0x63, 0x14, 0x1a, 0x61, 0x70, 0x3a, 0xce,
	0x2f, 0xab, 0x5f, 0x0d, 0x39, 0xbf, 0x2b, 0x02, 0x3e, 0x7d, 0xf8, 0xfe, 0x95, 0xbe, 0x20, 0xcb,
	0x7d, 0x4f, 0x2b, 0x46, 0x64, 0x4d, 0xdb, 0x5c, 0x1c, 0x77, 0xa7, 0xb5, 0xf8, 0x45, 0x8c, 0x48,
	0x5f, 0x92, 0x33, 0x8b, 0x46, 0x5f, 0xf5, 0x2e, 0x80, 0x50, 0x2a, 0x60, 0x8c, 0xec, 0xa8, 0xf0,
	0x1e, 0xd5, 0xfa, 0xdb, 0x5d, 0x99, 0x3e, 0x27, 0xc7, 0xca, 0x04, 0x94, 0xc9, 0x38, 0xcb, 0x16,
	0x6d, 0x73, 0xb1, 0xec, 0x6e, 0x0b, 0xab, 0x8f, 0xe4, 0x34, 0x0b, 0x49, 0x66, 0xc4, 0xe8, 0x51,
	0x52, 0x46, 0x1e, 0x46, 0x94, 0xce, 0xaa, 0x58, 0xe6, 0x2e, 0xbb, 0x7a, 0xa5, 0x2d, 0x39, 0xb1,
	0xc2, 0xba, 0x8a, 0x1e, 0x15, 0x74, 0x5e, 0x5a, 0xfd, 0x5c, 0xec, 0x9a, 0x65, 0x3b, 0x63, 0xd4,
	0x40, 0xd7, 0xe4, 0x69, 0x75, 0x9d, 0x6e, 0x3c, 0x42, 0x40, 0x89, 0x66, 0x8b, 0x6a, 0x6a, 0xfd,
	0x64, 0x02, 0xbf, 0xdd, 0x78, 0xec, 0x26, 0x88, 0xbe, 0x21, 0xcf, 0x92, 0x4b, 0x62, 0xc8, 0x39,
	0x69, 0x54, 0x50, 0x1b, 0x48, 0x77, 0x6d, 0xd3, 0x34, 0xf5, 0xbc, 0x30, 0x2e, 0x0b, 0xe1, 0xf3,
	0x0e, 0x7f, 0x97, 0x61, 0xfa, 0xa7, 0x21, 0x8f, 0xf7, 0x13, 0xb3, 0xa5, 0x24, 0x46, 0x5f, 0x4c,
	0x9f, 0xac, 0x13, 0x3f, 0xcc, 0xc7, 0xe5, 0xf3, 0x3c, 0xbb, 0xb3, 0xea, 0xb1, 0xaa, 0xa1, 0x9c,
	0x54, 0xdf, 0xa0, 0x44, 0x12, 0x30, 0xa0, 0xd5, 0xe9, 0x8a, 0xdd, 0x2b, 0xce, 0xaa, 0xfa, 0xf7,
	0x22, 0x89, 0xcb, 0x02, 0x64, 0xfe, 0x7f, 0x51, 0xe4, 0x67, 0xec, 0x7e, 0xbb, 0xc8, 0xfc, 0x61,
	0x1e, 0x42, 0x7e, 0xb5, 0xfa, 0xdb, 0x10, 0x7a, 0x77, 0xb7, 0xe8, 0xef, 0x66, 0xb6, 0x32, 0x53,
	0x27, 0xb6, 0x6e, 0x17, 0x87, 0x4e, 0xa6, 0x2e, 0xc7, 0xed, 0xa2, 0x4e, 0xe2, 0xfb, 0x07, 0xe5,
	0x1f, 0x7c, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x6a, 0x0d, 0x38, 0x9f, 0x03, 0x00, 0x00,
}
