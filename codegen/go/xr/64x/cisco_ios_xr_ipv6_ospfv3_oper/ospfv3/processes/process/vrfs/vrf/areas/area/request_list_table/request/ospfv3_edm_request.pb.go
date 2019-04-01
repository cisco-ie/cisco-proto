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

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_areas_area_request_list_table_request

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
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,5,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
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

func (m *Ospfv3EdmRequest_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3EdmRequest_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
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
	proto.RegisterType((*Ospfv3EdmRequest_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.request_list_table.request.ospfv3_edm_request_KEYS")
	proto.RegisterType((*Ospfv3EdmLsaSum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.request_list_table.request.ospfv3_edm_lsa_sum")
	proto.RegisterType((*Ospfv3EdmRequest)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.request_list_table.request.ospfv3_edm_request")
}

func init() { proto.RegisterFile("ospfv3_edm_request.proto", fileDescriptor_73bc4e6bb9c0a822) }

var fileDescriptor_73bc4e6bb9c0a822 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0x56, 0x7e, 0xfd, 0xb1, 0x6e, 0xee, 0xb2, 0x51, 0x2b, 0x5a, 0xb3, 0x5b, 0xa9, 0x00, 0x95,
	0x4b, 0x26, 0xad, 0xdb, 0x40, 0xdc, 0x7a, 0xe0, 0x50, 0x31, 0xf5, 0x90, 0x22, 0x24, 0x4e, 0x96,
	0x5b, 0xbf, 0x69, 0xad, 0xe5, 0x1f, 0x7e, 0x93, 0x88, 0x7e, 0x0e, 0xbe, 0x0a, 0x9f, 0x80, 0x4f,
	0xc5, 0x11, 0xc5, 0x76, 0x42, 0xab, 0x70, 0xe6, 0x12, 0x25, 0xcf, 0x9f, 0x3c, 0x7e, 0x5e, 0xdb,
	0xc4, 0xcf, 0x30, 0x8f, 0xaa, 0x19, 0x03, 0x91, 0x30, 0x05, 0x5f, 0x4b, 0xc0, 0x22, 0xc8, 0x55,
	0x56, 0x64, 0x14, 0x36, 0x12, 0x37, 0x19, 0x93, 0x19, 0xb2, 0x6f, 0x8a, 0xc9, 0xbc, 0x7a, 0x60,
	0x56, 0x9b, 0xe5, 0xa0, 0x02, 0xf3, 0x5e, 0x6b, 0x37, 0x80, 0x08, 0xd8, 0xbc, 0x05, 0x95, 0x8a,
	0xf4, 0x23, 0xe0, 0x0a, 0x38, 0xea, 0x67, 0x60, 0xff, 0xcc, 0x62, 0x89, 0x05, 0x2b, 0xf8, 0x3a,
	0x86, 0x06, 0x9a, 0xfc, 0x74, 0xc8, 0xa8, 0xbb, 0x06, 0xf6, 0xf1, 0xc3, 0x97, 0x15, 0x7d, 0x41,
	0xce, 0xed, 0x5f, 0x59, 0xca, 0x13, 0xf0, 0x9d, 0xb1, 0x33, 0x3d, 0x0b, 0x07, 0x16, 0x5b, 0xf2,
	0x04, 0xe8, 0x35, 0x39, 0xad, 0x54, 0x64, 0xe8, 0xff, 0x34, 0xdd, 0xaf, 0x54, 0xa4, 0xa9, 0x11,
	0xe9, 0xd7, 0xe1, 0x4c, 0x0a, 0xbf, 0x37, 0x76, 0xa6, 0x6e, 0x78, 0x52, 0x7f, 0x2e, 0x04, 0x7d,
	0x45, 0x2e, 0x64, 0x5a, 0x80, 0x8a, 0xf8, 0x06, 0x8c, 0xf3, 0x7f, 0xed, 0x74, 0x5b, 0x54, 0xfb,
	0xdf, 0x90, 0xe7, 0x29, 0xc8, 0xed, 0x6e, 0x9d, 0x29, 0xc6, 0x85, 0x50, 0x80, 0xe8, 0x3f, 0xd3,
	0xc2, 0xcb, 0x06, 0x9f, 0x1b, 0x78, 0xf2, 0xcb, 0x21, 0xf4, 0xa0, 0x44, 0x8c, 0x9c, 0x61, 0x99,
	0xd0, 0xd7, 0xe4, 0x72, 0x07, 0x5c, 0x80, 0xd2, 0x48, 0xb1, 0xcf, 0x9b, 0x0a, 0xae, 0x81, 0x1f,
	0x91, 0x7f, 0xda, 0xe7, 0x40, 0x5f, 0x92, 0x8b, 0x03, 0x1d, 0xdf, 0x9a, 0x2a, 0x6e, 0x78, 0xde,
	0xca, 0xe6, 0x5b, 0xa0, 0x13, 0xe2, 0x1e, 0xa8, 0x6c, 0xab, 0xb3, 0x70, 0xd0, 0x8a, 0x16, 0x82,
	0xbe, 0x27, 0xd7, 0x56, 0xc3, 0x45, 0x05, 0xaa, 0x90, 0x28, 0xd3, 0x2d, 0x53, 0x59, 0x59, 0x80,
	0xb2, 0x2d, 0x47, 0x46, 0x30, 0xff, 0xc3, 0x87, 0x9a, 0xa6, 0x77, 0xe4, 0xca, 0x7a, 0xb1, 0xde,
	0x84, 0xb4, 0x1e, 0x4e, 0x99, 0xac, 0x41, 0xe9, 0xd6, 0xc3, 0xd0, 0x33, 0xec, 0xca, 0x92, 0x4b,
	0xcd, 0x4d, 0x7e, 0xf4, 0x8e, 0xaa, 0xdb, 0xfd, 0xa3, 0xef, 0x88, 0xdf, 0x6c, 0x65, 0x67, 0x88,
	0xb7, 0x7a, 0x1d, 0x57, 0x96, 0x5f, 0x1e, 0xcf, 0x92, 0xde, 0x93, 0x91, 0xc4, 0xf6, 0x1c, 0x54,
	0x52, 0x15, 0x25, 0x8f, 0x59, 0x2c, 0xd3, 0x27, 0x7f, 0x36, 0x76, 0xa6, 0xa7, 0xa1, 0x27, 0x31,
	0x34, 0xec, 0x67, 0x43, 0x3e, 0xca, 0xf4, 0xa9, 0xb6, 0xfd, 0xcd, 0x53, 0xcf, 0xe9, 0x4e, 0x0f,
	0xd3, 0x53, 0x1d, 0xd3, 0x42, 0xd0, 0x1b, 0xe2, 0x1d, 0xa4, 0xe1, 0x8e, 0x27, 0x26, 0xea, 0x5e,
	0x47, 0x0d, 0xdb, 0xa8, 0xd5, 0x8e, 0x27, 0x3a, 0xe7, 0x86, 0x78, 0x1d, 0x75, 0x1d, 0xf2, 0xa0,
	0x43, 0x86, 0xea, 0x58, 0xbe, 0x10, 0xf4, 0xbb, 0x43, 0xfa, 0x16, 0xf5, 0xdf, 0x8e, 0x7b, 0xd3,
	0xc1, 0xed, 0x3e, 0xf8, 0x27, 0x57, 0x2b, 0xe8, 0x9e, 0xc8, 0xb0, 0x59, 0xc9, 0xfa, 0x44, 0x5f,
	0xf2, 0xd9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x07, 0x9f, 0x62, 0x00, 0x04, 0x00, 0x00,
}