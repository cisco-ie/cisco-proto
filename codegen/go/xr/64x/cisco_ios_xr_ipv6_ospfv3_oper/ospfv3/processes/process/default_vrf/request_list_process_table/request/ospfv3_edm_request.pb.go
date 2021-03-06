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
// source: ospfv3_edm_request.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_request_list_process_table_request

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

type Ospfv3EdmRequest_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmRequest_KEYS) Reset()         { *m = Ospfv3EdmRequest_KEYS{} }
func (m *Ospfv3EdmRequest_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmRequest_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmRequest_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_73bc4e6bb9c0a822, []int{0}
}

func (m *Ospfv3EdmRequest_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmRequest_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmRequest_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmRequest_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmRequest_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmRequest_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmRequest_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmRequest_KEYS.Size(m)
}
func (m *Ospfv3EdmRequest_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmRequest_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmRequest_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmRequest_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmRequest_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ospfv3EdmRequest_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type Ospfv3EdmLsaSum struct {
	HeaderLsaType           string   `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType,proto3" json:"header_lsa_type,omitempty"`
	HeaderLsaAge            uint32   `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge,proto3" json:"header_lsa_age,omitempty"`
	HeaderLsaId             string   `protobuf:"bytes,3,opt,name=header_lsa_id,json=headerLsaId,proto3" json:"header_lsa_id,omitempty"`
	HeaderAdvertisingRouter string   `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter,proto3" json:"header_advertising_router,omitempty"`
	HeaderSequenceNumber    int32    `protobuf:"zigzag32,5,opt,name=header_sequence_number,json=headerSequenceNumber,proto3" json:"header_sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Ospfv3EdmLsaSum) Reset()         { *m = Ospfv3EdmLsaSum{} }
func (m *Ospfv3EdmLsaSum) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmLsaSum) ProtoMessage()    {}
func (*Ospfv3EdmLsaSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_73bc4e6bb9c0a822, []int{1}
}

func (m *Ospfv3EdmLsaSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmLsaSum.Unmarshal(m, b)
}
func (m *Ospfv3EdmLsaSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmLsaSum.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmLsaSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmLsaSum.Merge(m, src)
}
func (m *Ospfv3EdmLsaSum) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmLsaSum.Size(m)
}
func (m *Ospfv3EdmLsaSum) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmLsaSum.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmLsaSum proto.InternalMessageInfo

func (m *Ospfv3EdmLsaSum) GetHeaderLsaType() string {
	if m != nil {
		return m.HeaderLsaType
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaAge() uint32 {
	if m != nil {
		return m.HeaderLsaAge
	}
	return 0
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaId() string {
	if m != nil {
		return m.HeaderLsaId
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderAdvertisingRouter() string {
	if m != nil {
		return m.HeaderAdvertisingRouter
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderSequenceNumber() int32 {
	if m != nil {
		return m.HeaderSequenceNumber
	}
	return 0
}

type Ospfv3EdmRequest struct {
	RequestNeighborAddress string             `protobuf:"bytes,50,opt,name=request_neighbor_address,json=requestNeighborAddress,proto3" json:"request_neighbor_address,omitempty"`
	IsRequestVirtualLink   bool               `protobuf:"varint,51,opt,name=is_request_virtual_link,json=isRequestVirtualLink,proto3" json:"is_request_virtual_link,omitempty"`
	RequestVirtualLinkId   uint32             `protobuf:"varint,52,opt,name=request_virtual_link_id,json=requestVirtualLinkId,proto3" json:"request_virtual_link_id,omitempty"`
	IsRequestShamLink      bool               `protobuf:"varint,53,opt,name=is_request_sham_link,json=isRequestShamLink,proto3" json:"is_request_sham_link,omitempty"`
	RequestShamLinkId      uint32             `protobuf:"varint,54,opt,name=request_sham_link_id,json=requestShamLinkId,proto3" json:"request_sham_link_id,omitempty"`
	Request                []*Ospfv3EdmLsaSum `protobuf:"bytes,55,rep,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *Ospfv3EdmRequest) Reset()         { *m = Ospfv3EdmRequest{} }
func (m *Ospfv3EdmRequest) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmRequest) ProtoMessage()    {}
func (*Ospfv3EdmRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73bc4e6bb9c0a822, []int{2}
}

func (m *Ospfv3EdmRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmRequest.Unmarshal(m, b)
}
func (m *Ospfv3EdmRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmRequest.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmRequest.Merge(m, src)
}
func (m *Ospfv3EdmRequest) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmRequest.Size(m)
}
func (m *Ospfv3EdmRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmRequest proto.InternalMessageInfo

func (m *Ospfv3EdmRequest) GetRequestNeighborAddress() string {
	if m != nil {
		return m.RequestNeighborAddress
	}
	return ""
}

func (m *Ospfv3EdmRequest) GetIsRequestVirtualLink() bool {
	if m != nil {
		return m.IsRequestVirtualLink
	}
	return false
}

func (m *Ospfv3EdmRequest) GetRequestVirtualLinkId() uint32 {
	if m != nil {
		return m.RequestVirtualLinkId
	}
	return 0
}

func (m *Ospfv3EdmRequest) GetIsRequestShamLink() bool {
	if m != nil {
		return m.IsRequestShamLink
	}
	return false
}

func (m *Ospfv3EdmRequest) GetRequestShamLinkId() uint32 {
	if m != nil {
		return m.RequestShamLinkId
	}
	return 0
}

func (m *Ospfv3EdmRequest) GetRequest() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3EdmRequest_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.request_list_process_table.request.ospfv3_edm_request_KEYS")
	proto.RegisterType((*Ospfv3EdmLsaSum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.request_list_process_table.request.ospfv3_edm_lsa_sum")
	proto.RegisterType((*Ospfv3EdmRequest)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.request_list_process_table.request.ospfv3_edm_request")
}

func init() { proto.RegisterFile("ospfv3_edm_request.proto", fileDescriptor_73bc4e6bb9c0a822) }

var fileDescriptor_73bc4e6bb9c0a822 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x15, 0xca, 0xaf, 0x3b, 0x9d, 0xa1, 0x56, 0x34, 0x0d, 0xbb, 0x52, 0x01, 0x2a, 0x9b,
	0x8c, 0x34, 0x9d, 0x19, 0x10, 0xbb, 0x2e, 0x58, 0x54, 0x8c, 0xba, 0x48, 0x11, 0x12, 0x2b, 0xcb,
	0xad, 0x6f, 0x5b, 0xab, 0xf9, 0xc3, 0xd7, 0x89, 0xe8, 0x4b, 0xb0, 0xe1, 0x55, 0x78, 0x38, 0x96,
	0x28, 0xb6, 0x13, 0x52, 0x85, 0x35, 0x3b, 0xeb, 0x9c, 0xef, 0xf6, 0x9c, 0x5e, 0xc7, 0x24, 0xc8,
	0x30, 0xdf, 0x96, 0x33, 0x06, 0x22, 0x61, 0x0a, 0xbe, 0x15, 0x80, 0x3a, 0xcc, 0x55, 0xa6, 0x33,
	0x0a, 0x1b, 0x89, 0x9b, 0x8c, 0xc9, 0x0c, 0xd9, 0x77, 0xc5, 0x64, 0x5e, 0xde, 0x31, 0xc7, 0x66,
	0x39, 0xa8, 0xd0, 0x9e, 0x2b, 0x76, 0x03, 0x88, 0x80, 0xf5, 0x29, 0x14, 0xb0, 0xe5, 0x45, 0xac,
	0x59, 0xa9, 0xb6, 0xa1, 0xfb, 0x45, 0x16, 0x4b, 0xd4, 0xcc, 0x01, 0x4c, 0xf3, 0x75, 0x0c, 0xb5,
	0x35, 0xf9, 0xe1, 0x91, 0x51, 0xb7, 0x03, 0xfb, 0xf4, 0xf1, 0xeb, 0x8a, 0xbe, 0x24, 0x67, 0xf5,
	0x50, 0xca, 0x13, 0x08, 0xbc, 0xb1, 0x37, 0x7d, 0x16, 0xf5, 0x9d, 0xb6, 0xe4, 0x09, 0xd0, 0xd7,
	0xe4, 0x5c, 0xa6, 0x1a, 0xd4, 0x96, 0x6f, 0xc0, 0x42, 0x0f, 0x0c, 0x34, 0x68, 0x54, 0x83, 0xbd,
	0x25, 0xcf, 0x53, 0x90, 0xbb, 0xfd, 0x3a, 0x53, 0x8c, 0x0b, 0xa1, 0x00, 0x31, 0xe8, 0x19, 0xf0,
	0xa2, 0xd6, 0xe7, 0x56, 0x9e, 0xfc, 0xf6, 0x08, 0x6d, 0x15, 0x8a, 0x91, 0x33, 0x2c, 0x12, 0xfa,
	0x86, 0x5c, 0xec, 0x81, 0x0b, 0x50, 0x46, 0xd1, 0xc7, 0xbc, 0xae, 0x33, 0xb0, 0xf2, 0x3d, 0xf2,
	0xcf, 0xc7, 0x1c, 0xe8, 0x2b, 0x72, 0xde, 0xe2, 0xf8, 0xce, 0x16, 0x1a, 0x44, 0x67, 0x0d, 0x36,
	0xdf, 0x01, 0x9d, 0x90, 0x41, 0x8b, 0x92, 0xc2, 0x95, 0xe9, 0x37, 0xd0, 0x42, 0xd0, 0x0f, 0xe4,
	0x85, 0x63, 0xb8, 0x28, 0x41, 0x69, 0x89, 0x32, 0xdd, 0x31, 0x95, 0x15, 0x1a, 0x54, 0xf0, 0xd0,
	0xf0, 0x23, 0x0b, 0xcc, 0xff, 0xfa, 0x91, 0xb1, 0xe9, 0x0d, 0xb9, 0x74, 0xb3, 0x58, 0x2d, 0x34,
	0xad, 0x96, 0x53, 0x24, 0x6b, 0x50, 0xc1, 0xa3, 0xb1, 0x37, 0x1d, 0x46, 0xbe, 0x75, 0x57, 0xce,
	0x5c, 0x1a, 0x6f, 0xf2, 0xab, 0x77, 0xf2, 0xd7, 0xdd, 0x5d, 0xd0, 0xf7, 0x24, 0xa8, 0xaf, 0xa5,
	0xb3, 0xc4, 0x6b, 0xd3, 0xe3, 0xd2, 0xf9, 0xcb, 0xd3, 0x5d, 0xd2, 0x5b, 0x32, 0x92, 0xd8, 0xdc,
	0x69, 0x29, 0x95, 0x2e, 0x78, 0xcc, 0x62, 0x99, 0x1e, 0x82, 0xd9, 0xd8, 0x9b, 0x3e, 0x8d, 0x7c,
	0x89, 0x91, 0x75, 0xbf, 0x58, 0xf3, 0x5e, 0xa6, 0x87, 0x6a, 0xec, 0x5f, 0x33, 0xd5, 0x9e, 0x6e,
	0xcc, 0x32, 0x7d, 0xd5, 0x19, 0x5a, 0x08, 0x7a, 0x45, 0xfc, 0x56, 0x1a, 0xee, 0x79, 0x62, 0xa3,
	0x6e, 0x4d, 0xd4, 0xb0, 0x89, 0x5a, 0xed, 0x79, 0x62, 0x72, 0xae, 0x88, 0xdf, 0xa1, 0xab, 0x90,
	0x3b, 0x13, 0x32, 0x54, 0xa7, 0xf8, 0x42, 0xd0, 0x9f, 0x1e, 0x79, 0xe2, 0xd4, 0xe0, 0xdd, 0xb8,
	0x37, 0xed, 0x5f, 0x1f, 0xc3, 0xff, 0xf2, 0x4c, 0xc2, 0xee, 0x17, 0x19, 0xd5, 0x4d, 0xd6, 0x8f,
	0xcd, 0x83, 0x9d, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x6d, 0xae, 0x5e, 0xcc, 0x03, 0x00,
	0x00,
}
