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
// source: rsvp_mgmt_nbr_detail.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_interface_neighbor_details_interface_neighbor_detail

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

type RsvpMgmtNbrDetail_KEYS struct {
	NeighborAddress      string   `protobuf:"bytes,1,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtNbrDetail_KEYS) Reset()         { *m = RsvpMgmtNbrDetail_KEYS{} }
func (m *RsvpMgmtNbrDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtNbrDetail_KEYS) ProtoMessage()    {}
func (*RsvpMgmtNbrDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_94bb71eff5358a73, []int{0}
}

func (m *RsvpMgmtNbrDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtNbrDetail_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtNbrDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtNbrDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtNbrDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtNbrDetail_KEYS.Merge(m, src)
}
func (m *RsvpMgmtNbrDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtNbrDetail_KEYS.Size(m)
}
func (m *RsvpMgmtNbrDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtNbrDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtNbrDetail_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtNbrDetail_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type RsvpMgmtIfNbrDetail struct {
	InterfaceNeighborAddress string   `protobuf:"bytes,1,opt,name=interface_neighbor_address,json=interfaceNeighborAddress,proto3" json:"interface_neighbor_address,omitempty"`
	NeighborInterfaceName    string   `protobuf:"bytes,2,opt,name=neighbor_interface_name,json=neighborInterfaceName,proto3" json:"neighbor_interface_name,omitempty"`
	IsRrEnabled              bool     `protobuf:"varint,3,opt,name=is_rr_enabled,json=isRrEnabled,proto3" json:"is_rr_enabled,omitempty"`
	NeighborEpoch            uint32   `protobuf:"varint,4,opt,name=neighbor_epoch,json=neighborEpoch,proto3" json:"neighbor_epoch,omitempty"`
	OutOfOrderMessages       uint32   `protobuf:"varint,5,opt,name=out_of_order_messages,json=outOfOrderMessages,proto3" json:"out_of_order_messages,omitempty"`
	RetransmittedMessages    uint32   `protobuf:"varint,6,opt,name=retransmitted_messages,json=retransmittedMessages,proto3" json:"retransmitted_messages,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *RsvpMgmtIfNbrDetail) Reset()         { *m = RsvpMgmtIfNbrDetail{} }
func (m *RsvpMgmtIfNbrDetail) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtIfNbrDetail) ProtoMessage()    {}
func (*RsvpMgmtIfNbrDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_94bb71eff5358a73, []int{1}
}

func (m *RsvpMgmtIfNbrDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtIfNbrDetail.Unmarshal(m, b)
}
func (m *RsvpMgmtIfNbrDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtIfNbrDetail.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtIfNbrDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtIfNbrDetail.Merge(m, src)
}
func (m *RsvpMgmtIfNbrDetail) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtIfNbrDetail.Size(m)
}
func (m *RsvpMgmtIfNbrDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtIfNbrDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtIfNbrDetail proto.InternalMessageInfo

func (m *RsvpMgmtIfNbrDetail) GetInterfaceNeighborAddress() string {
	if m != nil {
		return m.InterfaceNeighborAddress
	}
	return ""
}

func (m *RsvpMgmtIfNbrDetail) GetNeighborInterfaceName() string {
	if m != nil {
		return m.NeighborInterfaceName
	}
	return ""
}

func (m *RsvpMgmtIfNbrDetail) GetIsRrEnabled() bool {
	if m != nil {
		return m.IsRrEnabled
	}
	return false
}

func (m *RsvpMgmtIfNbrDetail) GetNeighborEpoch() uint32 {
	if m != nil {
		return m.NeighborEpoch
	}
	return 0
}

func (m *RsvpMgmtIfNbrDetail) GetOutOfOrderMessages() uint32 {
	if m != nil {
		return m.OutOfOrderMessages
	}
	return 0
}

func (m *RsvpMgmtIfNbrDetail) GetRetransmittedMessages() uint32 {
	if m != nil {
		return m.RetransmittedMessages
	}
	return 0
}

type RsvpMgmtNbrDetail struct {
	NodeAddress                 string                 `protobuf:"bytes,50,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	InterfaceNeighborListDetail []*RsvpMgmtIfNbrDetail `protobuf:"bytes,51,rep,name=interface_neighbor_list_detail,json=interfaceNeighborListDetail,proto3" json:"interface_neighbor_list_detail,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}               `json:"-"`
	XXX_unrecognized            []byte                 `json:"-"`
	XXX_sizecache               int32                  `json:"-"`
}

func (m *RsvpMgmtNbrDetail) Reset()         { *m = RsvpMgmtNbrDetail{} }
func (m *RsvpMgmtNbrDetail) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtNbrDetail) ProtoMessage()    {}
func (*RsvpMgmtNbrDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_94bb71eff5358a73, []int{2}
}

func (m *RsvpMgmtNbrDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtNbrDetail.Unmarshal(m, b)
}
func (m *RsvpMgmtNbrDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtNbrDetail.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtNbrDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtNbrDetail.Merge(m, src)
}
func (m *RsvpMgmtNbrDetail) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtNbrDetail.Size(m)
}
func (m *RsvpMgmtNbrDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtNbrDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtNbrDetail proto.InternalMessageInfo

func (m *RsvpMgmtNbrDetail) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *RsvpMgmtNbrDetail) GetInterfaceNeighborListDetail() []*RsvpMgmtIfNbrDetail {
	if m != nil {
		return m.InterfaceNeighborListDetail
	}
	return nil
}

func init() {
	proto.RegisterType((*RsvpMgmtNbrDetail_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.interface_neighbor_details.interface_neighbor_detail.rsvp_mgmt_nbr_detail_KEYS")
	proto.RegisterType((*RsvpMgmtIfNbrDetail)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.interface_neighbor_details.interface_neighbor_detail.rsvp_mgmt_if_nbr_detail")
	proto.RegisterType((*RsvpMgmtNbrDetail)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.interface_neighbor_details.interface_neighbor_detail.rsvp_mgmt_nbr_detail")
}

func init() { proto.RegisterFile("rsvp_mgmt_nbr_detail.proto", fileDescriptor_94bb71eff5358a73) }

var fileDescriptor_94bb71eff5358a73 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcb, 0x4a, 0xf3, 0x40,
	0x14, 0xc7, 0x49, 0xfb, 0x7d, 0x45, 0xa7, 0x56, 0x65, 0xb0, 0x36, 0x56, 0x90, 0x18, 0x10, 0xe2,
	0x26, 0x60, 0x8b, 0xae, 0xdc, 0x08, 0x56, 0x10, 0x2f, 0x85, 0x74, 0xe5, 0x6a, 0x98, 0x24, 0x27,
	0xed, 0x40, 0x92, 0x09, 0x73, 0xa6, 0xe2, 0x6b, 0xf9, 0x0e, 0xbe, 0x8e, 0xef, 0x20, 0x49, 0x9b,
	0xb4, 0xc5, 0x74, 0xe9, 0x2e, 0xfc, 0x2f, 0xbf, 0x24, 0xe7, 0x1c, 0xd2, 0x57, 0xf8, 0x9e, 0xb1,
	0x64, 0x9a, 0x68, 0x96, 0xfa, 0x8a, 0x85, 0xa0, 0xb9, 0x88, 0xdd, 0x4c, 0x49, 0x2d, 0xe9, 0x24,
	0x10, 0x18, 0x48, 0x26, 0x24, 0xb2, 0x0f, 0xc5, 0x44, 0xc6, 0x8a, 0xac, 0xcc, 0x40, 0xb9, 0xf9,
	0x93, 0x2b, 0x52, 0x0d, 0x2a, 0xe2, 0x01, 0xb0, 0x14, 0xc4, 0x74, 0xe6, 0xcb, 0xb2, 0x8f, 0xdb,
	0x2d, 0xfb, 0x81, 0x9c, 0xd4, 0xbd, 0x92, 0x3d, 0x8d, 0xde, 0x26, 0xf4, 0x92, 0x1c, 0x56, 0x79,
	0x1e, 0x86, 0x0a, 0x10, 0x4d, 0xc3, 0x32, 0x9c, 0x5d, 0xef, 0xa0, 0xd4, 0xef, 0x16, 0xb2, 0xfd,
	0xd5, 0x20, 0xbd, 0x15, 0x48, 0x44, 0x6b, 0x2c, 0x7a, 0x4b, 0xfa, 0x35, 0x1f, 0xb0, 0x09, 0x34,
	0xab, 0xc4, 0xeb, 0x26, 0x99, 0xde, 0x90, 0x5e, 0xd5, 0x59, 0xc3, 0xf0, 0x04, 0xcc, 0x46, 0x51,
	0xed, 0x96, 0xf6, 0x63, 0x85, 0xe0, 0x09, 0x50, 0x9b, 0x74, 0x04, 0x32, 0xa5, 0x18, 0xa4, 0xdc,
	0x8f, 0x21, 0x34, 0x9b, 0x96, 0xe1, 0xec, 0x78, 0x6d, 0x81, 0x9e, 0x1a, 0x2d, 0x24, 0x7a, 0x41,
	0xf6, 0x2b, 0x36, 0x64, 0x32, 0x98, 0x99, 0xff, 0x2c, 0xc3, 0xe9, 0x78, 0x9d, 0x52, 0x1d, 0xe5,
	0x22, 0xbd, 0x22, 0x5d, 0x39, 0xd7, 0x4c, 0x46, 0x4c, 0xaa, 0x10, 0x14, 0x4b, 0x00, 0x91, 0x4f,
	0x01, 0xcd, 0xff, 0x45, 0x9a, 0xca, 0xb9, 0x1e, 0x47, 0xe3, 0xdc, 0x7a, 0x59, 0x3a, 0xf4, 0x9a,
	0x1c, 0x2b, 0xd0, 0x8a, 0xa7, 0x98, 0x08, 0xad, 0x21, 0x5c, 0x75, 0x5a, 0x45, 0xa7, 0xbb, 0xe1,
	0x96, 0x35, 0xfb, 0xdb, 0x20, 0x47, 0x75, 0xfb, 0xa0, 0xe7, 0x64, 0x2f, 0x95, 0x21, 0x54, 0x53,
	0x1b, 0x14, 0xbf, 0xde, 0xce, 0xb5, 0x72, 0x50, 0x9f, 0x06, 0x39, 0xab, 0x99, 0x73, 0x2c, 0x50,
	0x2f, 0x29, 0xe6, 0xd0, 0x6a, 0x3a, 0xed, 0x41, 0xec, 0xfe, 0xc1, 0x25, 0xb9, 0x5b, 0xb6, 0xef,
	0x9d, 0xfe, 0xda, 0xec, 0xb3, 0x40, 0x7d, 0x5f, 0x98, 0x7e, 0xab, 0x38, 0xed, 0xe1, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd0, 0x6e, 0x54, 0x21, 0xf8, 0x02, 0x00, 0x00,
}
