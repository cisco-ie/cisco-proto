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
// source: ospf_sh_request_list.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_requests_request

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

type OspfShRequestList_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRequestList_KEYS) Reset()         { *m = OspfShRequestList_KEYS{} }
func (m *OspfShRequestList_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShRequestList_KEYS) ProtoMessage()    {}
func (*OspfShRequestList_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4609c816fd64cee, []int{0}
}

func (m *OspfShRequestList_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRequestList_KEYS.Unmarshal(m, b)
}
func (m *OspfShRequestList_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRequestList_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShRequestList_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRequestList_KEYS.Merge(m, src)
}
func (m *OspfShRequestList_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShRequestList_KEYS.Size(m)
}
func (m *OspfShRequestList_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRequestList_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRequestList_KEYS proto.InternalMessageInfo

func (m *OspfShRequestList_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShRequestList_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShRequestList_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type OspfShLsaSum struct {
	HeaderLsaType           string   `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType,proto3" json:"header_lsa_type,omitempty"`
	HeaderLsaAge            uint32   `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge,proto3" json:"header_lsa_age,omitempty"`
	HeaderLsId              string   `protobuf:"bytes,3,opt,name=header_ls_id,json=headerLsId,proto3" json:"header_ls_id,omitempty"`
	HeaderAdvertisingRouter string   `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter,proto3" json:"header_advertising_router,omitempty"`
	HeaderSequenceNumber    uint32   `protobuf:"varint,5,opt,name=header_sequence_number,json=headerSequenceNumber,proto3" json:"header_sequence_number,omitempty"`
	HeaderLsaChecksum       uint32   `protobuf:"varint,6,opt,name=header_lsa_checksum,json=headerLsaChecksum,proto3" json:"header_lsa_checksum,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *OspfShLsaSum) Reset()         { *m = OspfShLsaSum{} }
func (m *OspfShLsaSum) String() string { return proto.CompactTextString(m) }
func (*OspfShLsaSum) ProtoMessage()    {}
func (*OspfShLsaSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4609c816fd64cee, []int{1}
}

func (m *OspfShLsaSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShLsaSum.Unmarshal(m, b)
}
func (m *OspfShLsaSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShLsaSum.Marshal(b, m, deterministic)
}
func (m *OspfShLsaSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShLsaSum.Merge(m, src)
}
func (m *OspfShLsaSum) XXX_Size() int {
	return xxx_messageInfo_OspfShLsaSum.Size(m)
}
func (m *OspfShLsaSum) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShLsaSum.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShLsaSum proto.InternalMessageInfo

func (m *OspfShLsaSum) GetHeaderLsaType() string {
	if m != nil {
		return m.HeaderLsaType
	}
	return ""
}

func (m *OspfShLsaSum) GetHeaderLsaAge() uint32 {
	if m != nil {
		return m.HeaderLsaAge
	}
	return 0
}

func (m *OspfShLsaSum) GetHeaderLsId() string {
	if m != nil {
		return m.HeaderLsId
	}
	return ""
}

func (m *OspfShLsaSum) GetHeaderAdvertisingRouter() string {
	if m != nil {
		return m.HeaderAdvertisingRouter
	}
	return ""
}

func (m *OspfShLsaSum) GetHeaderSequenceNumber() uint32 {
	if m != nil {
		return m.HeaderSequenceNumber
	}
	return 0
}

func (m *OspfShLsaSum) GetHeaderLsaChecksum() uint32 {
	if m != nil {
		return m.HeaderLsaChecksum
	}
	return 0
}

type OspfShRequestList struct {
	RequestNeighborId      string          `protobuf:"bytes,50,opt,name=request_neighbor_id,json=requestNeighborId,proto3" json:"request_neighbor_id,omitempty"`
	RequestNeighborAddress string          `protobuf:"bytes,51,opt,name=request_neighbor_address,json=requestNeighborAddress,proto3" json:"request_neighbor_address,omitempty"`
	RequestInterfaceName   string          `protobuf:"bytes,52,opt,name=request_interface_name,json=requestInterfaceName,proto3" json:"request_interface_name,omitempty"`
	Request                []*OspfShLsaSum `protobuf:"bytes,53,rep,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *OspfShRequestList) Reset()         { *m = OspfShRequestList{} }
func (m *OspfShRequestList) String() string { return proto.CompactTextString(m) }
func (*OspfShRequestList) ProtoMessage()    {}
func (*OspfShRequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4609c816fd64cee, []int{2}
}

func (m *OspfShRequestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRequestList.Unmarshal(m, b)
}
func (m *OspfShRequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRequestList.Marshal(b, m, deterministic)
}
func (m *OspfShRequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRequestList.Merge(m, src)
}
func (m *OspfShRequestList) XXX_Size() int {
	return xxx_messageInfo_OspfShRequestList.Size(m)
}
func (m *OspfShRequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRequestList.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRequestList proto.InternalMessageInfo

func (m *OspfShRequestList) GetRequestNeighborId() string {
	if m != nil {
		return m.RequestNeighborId
	}
	return ""
}

func (m *OspfShRequestList) GetRequestNeighborAddress() string {
	if m != nil {
		return m.RequestNeighborAddress
	}
	return ""
}

func (m *OspfShRequestList) GetRequestInterfaceName() string {
	if m != nil {
		return m.RequestInterfaceName
	}
	return ""
}

func (m *OspfShRequestList) GetRequest() []*OspfShLsaSum {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShRequestList_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.requests.request.ospf_sh_request_list_KEYS")
	proto.RegisterType((*OspfShLsaSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.requests.request.ospf_sh_lsa_sum")
	proto.RegisterType((*OspfShRequestList)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.requests.request.ospf_sh_request_list")
}

func init() { proto.RegisterFile("ospf_sh_request_list.proto", fileDescriptor_e4609c816fd64cee) }

var fileDescriptor_e4609c816fd64cee = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x86, 0xd5, 0x0e, 0x86, 0xf8, 0xd6, 0xae, 0xcc, 0x54, 0x23, 0xe3, 0x54, 0x2a, 0x40, 0xe3,
	0x92, 0xc3, 0x56, 0x24, 0xc4, 0xad, 0x42, 0x1c, 0x2a, 0x50, 0x0f, 0x1d, 0x17, 0x4e, 0x96, 0x1b,
	0x7f, 0x69, 0x0d, 0x8d, 0x1d, 0xfc, 0x39, 0x15, 0xfd, 0x17, 0xfb, 0x0f, 0xfc, 0x1a, 0xfe, 0x15,
	0x8a, 0x63, 0x87, 0x51, 0x7a, 0xde, 0xcd, 0x79, 0xdf, 0xe7, 0x75, 0xde, 0xd8, 0x5f, 0xe0, 0xb9,
	0xa1, 0x32, 0xe7, 0xb4, 0xe6, 0x16, 0x7f, 0x54, 0x48, 0x8e, 0x6f, 0x14, 0xb9, 0xb4, 0xb4, 0xc6,
	0x19, 0x86, 0x99, 0xa2, 0xcc, 0x70, 0x65, 0x88, 0xff, 0xb4, 0x5c, 0x95, 0xdb, 0x09, 0xf7, 0xb4,
	0x29, 0xd1, 0xa6, 0xf5, 0xaa, 0xe6, 0x32, 0x24, 0x42, 0x8a, 0xab, 0x54, 0x62, 0x2e, 0xaa, 0x8d,
	0xe3, 0x5b, 0x9b, 0xa7, 0x42, 0x7e, 0x13, 0x19, 0xea, 0x6c, 0xc7, 0x95, 0xce, 0x8d, 0x2d, 0x84,
	0x53, 0x46, 0xa7, 0xe1, 0x45, 0x14, 0x17, 0xe3, 0xdb, 0x0e, 0x5c, 0x1c, 0x6a, 0xc1, 0x3f, 0x7d,
	0xfc, 0x7a, 0xc3, 0x5e, 0x40, 0x2f, 0xec, 0xcd, 0xb5, 0x28, 0x30, 0xe9, 0x8c, 0x3a, 0x97, 0x8f,
	0x17, 0x27, 0x41, 0x9b, 0x8b, 0x02, 0xd9, 0x2b, 0x38, 0x55, 0xda, 0xa1, 0xcd, 0x45, 0x86, 0x0d,
	0xd4, 0xf5, 0x50, 0xbf, 0x55, 0x3d, 0xf6, 0x06, 0x9e, 0x68, 0x54, 0xab, 0xf5, 0xd2, 0x58, 0x2e,
	0xa4, 0xb4, 0x48, 0x94, 0x1c, 0x79, 0x70, 0x10, 0xf5, 0x69, 0x23, 0x8f, 0x7f, 0x75, 0x61, 0x10,
	0x2b, 0x6d, 0x48, 0x70, 0xaa, 0x0a, 0xf6, 0x1a, 0x06, 0x6b, 0x14, 0x12, 0xad, 0x57, 0xdc, 0xae,
	0x8c, 0x5d, 0xfa, 0x8d, 0xfc, 0x99, 0xc4, 0x97, 0x5d, 0x89, 0xec, 0x25, 0x9c, 0xde, 0xe1, 0xc4,
	0xaa, 0x69, 0xd3, 0x5f, 0xf4, 0x5a, 0x6c, 0xba, 0x42, 0x36, 0x82, 0x5e, 0x4b, 0x71, 0x25, 0x43,
	0x11, 0x88, 0xcc, 0x4c, 0xb2, 0xf7, 0x70, 0x11, 0x08, 0x21, 0xb7, 0x68, 0x9d, 0x22, 0xa5, 0x57,
	0xdc, 0x9a, 0xca, 0xa1, 0x4d, 0x1e, 0x78, 0xfc, 0x59, 0x03, 0x4c, 0xff, 0xfa, 0x0b, 0x6f, 0xb3,
	0x09, 0x9c, 0x87, 0x2c, 0xd5, 0x07, 0xaa, 0xeb, 0x73, 0xa9, 0x8a, 0x25, 0xda, 0xe4, 0xa1, 0xef,
	0x32, 0x6c, 0xdc, 0x9b, 0x60, 0xce, 0xbd, 0xc7, 0x52, 0x78, 0x7a, 0xa7, 0x79, 0xb6, 0xc6, 0xec,
	0x3b, 0x55, 0x45, 0x72, 0xec, 0x23, 0x67, 0x6d, 0xfd, 0x0f, 0xc1, 0x18, 0xff, 0xee, 0xc2, 0xf0,
	0xd0, 0xc5, 0xd5, 0x1b, 0xc5, 0xe7, 0xf6, 0xc4, 0x95, 0x4c, 0xae, 0x7c, 0xe9, 0xb3, 0x60, 0xcd,
	0x83, 0x33, 0x93, 0xec, 0x1d, 0x24, 0xff, 0xf1, 0xf1, 0x86, 0xae, 0x7d, 0xe8, 0x7c, 0x2f, 0x14,
	0x2e, 0xaa, 0xfe, 0xd0, 0x98, 0xdc, 0x1b, 0x81, 0x89, 0xcf, 0x0d, 0x83, 0x3b, 0xfb, 0x67, 0x12,
	0x6e, 0x3b, 0xf0, 0x28, 0x18, 0xc9, 0xdb, 0xd1, 0xd1, 0xe5, 0xc9, 0xd5, 0x36, 0xbd, 0x97, 0x59,
	0x4f, 0xf7, 0x86, 0x6a, 0x11, 0x6b, 0x2c, 0x8f, 0xfd, 0x2f, 0x77, 0xfd, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x99, 0xe9, 0x15, 0x34, 0x90, 0x03, 0x00, 0x00,
}
