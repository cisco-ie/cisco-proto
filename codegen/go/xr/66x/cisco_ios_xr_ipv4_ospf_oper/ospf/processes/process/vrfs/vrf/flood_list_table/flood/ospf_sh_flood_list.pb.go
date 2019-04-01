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
// source: ospf_sh_flood_list.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_flood_list_table_flood

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

type OspfShFloodList_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShFloodList_KEYS) Reset()         { *m = OspfShFloodList_KEYS{} }
func (m *OspfShFloodList_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShFloodList_KEYS) ProtoMessage()    {}
func (*OspfShFloodList_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8800631f4d99f94, []int{0}
}

func (m *OspfShFloodList_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShFloodList_KEYS.Unmarshal(m, b)
}
func (m *OspfShFloodList_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShFloodList_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShFloodList_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShFloodList_KEYS.Merge(m, src)
}
func (m *OspfShFloodList_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShFloodList_KEYS.Size(m)
}
func (m *OspfShFloodList_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShFloodList_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShFloodList_KEYS proto.InternalMessageInfo

func (m *OspfShFloodList_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShFloodList_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShFloodList_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
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
	return fileDescriptor_d8800631f4d99f94, []int{1}
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

type OspfShFloodList struct {
	FloodInterfaceName   string          `protobuf:"bytes,50,opt,name=flood_interface_name,json=floodInterfaceName,proto3" json:"flood_interface_name,omitempty"`
	FloodPacingTimer     uint32          `protobuf:"varint,51,opt,name=flood_pacing_timer,json=floodPacingTimer,proto3" json:"flood_pacing_timer,omitempty"`
	FloodLsaCount        uint32          `protobuf:"varint,52,opt,name=flood_lsa_count,json=floodLsaCount,proto3" json:"flood_lsa_count,omitempty"`
	AreaFlood            []*OspfShLsaSum `protobuf:"bytes,53,rep,name=area_flood,json=areaFlood,proto3" json:"area_flood,omitempty"`
	AsFlood              []*OspfShLsaSum `protobuf:"bytes,54,rep,name=as_flood,json=asFlood,proto3" json:"as_flood,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OspfShFloodList) Reset()         { *m = OspfShFloodList{} }
func (m *OspfShFloodList) String() string { return proto.CompactTextString(m) }
func (*OspfShFloodList) ProtoMessage()    {}
func (*OspfShFloodList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8800631f4d99f94, []int{2}
}

func (m *OspfShFloodList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShFloodList.Unmarshal(m, b)
}
func (m *OspfShFloodList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShFloodList.Marshal(b, m, deterministic)
}
func (m *OspfShFloodList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShFloodList.Merge(m, src)
}
func (m *OspfShFloodList) XXX_Size() int {
	return xxx_messageInfo_OspfShFloodList.Size(m)
}
func (m *OspfShFloodList) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShFloodList.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShFloodList proto.InternalMessageInfo

func (m *OspfShFloodList) GetFloodInterfaceName() string {
	if m != nil {
		return m.FloodInterfaceName
	}
	return ""
}

func (m *OspfShFloodList) GetFloodPacingTimer() uint32 {
	if m != nil {
		return m.FloodPacingTimer
	}
	return 0
}

func (m *OspfShFloodList) GetFloodLsaCount() uint32 {
	if m != nil {
		return m.FloodLsaCount
	}
	return 0
}

func (m *OspfShFloodList) GetAreaFlood() []*OspfShLsaSum {
	if m != nil {
		return m.AreaFlood
	}
	return nil
}

func (m *OspfShFloodList) GetAsFlood() []*OspfShLsaSum {
	if m != nil {
		return m.AsFlood
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShFloodList_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.flood_list_table.flood.ospf_sh_flood_list_KEYS")
	proto.RegisterType((*OspfShLsaSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.flood_list_table.flood.ospf_sh_lsa_sum")
	proto.RegisterType((*OspfShFloodList)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.flood_list_table.flood.ospf_sh_flood_list")
}

func init() { proto.RegisterFile("ospf_sh_flood_list.proto", fileDescriptor_d8800631f4d99f94) }

var fileDescriptor_d8800631f4d99f94 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x95, 0x06, 0xfa, 0x73, 0x9b, 0xb4, 0x60, 0x2a, 0x3a, 0xdd, 0x85, 0x08, 0x50, 0x17,
	0x68, 0x84, 0xda, 0xc0, 0x82, 0x5d, 0x85, 0x40, 0xaa, 0x40, 0x15, 0x9a, 0x76, 0xc3, 0xca, 0x72,
	0x66, 0xee, 0x24, 0x16, 0x99, 0xf1, 0xe0, 0xeb, 0x19, 0xd1, 0x4d, 0x17, 0xbc, 0x09, 0xe2, 0x45,
	0x91, 0xaf, 0x9d, 0x34, 0x6a, 0xb7, 0xa8, 0x9b, 0xc8, 0x39, 0xe7, 0xb3, 0xcf, 0xf1, 0xcf, 0x40,
	0x62, 0xa8, 0x29, 0x25, 0xcd, 0x65, 0xb9, 0x30, 0xa6, 0x90, 0x0b, 0x4d, 0x2e, 0x6d, 0xac, 0x71,
	0x46, 0x64, 0xb9, 0xa6, 0xdc, 0x48, 0x6d, 0x48, 0xfe, 0xb2, 0x52, 0x37, 0xdd, 0x44, 0x32, 0x6b,
	0x1a, 0xb4, 0xa9, 0x1f, 0x79, 0x2e, 0x47, 0x22, 0xa4, 0xe5, 0x28, 0xed, 0x6c, 0xc9, 0x3f, 0xe9,
	0xed, 0x6a, 0xd2, 0xa9, 0xe9, 0x02, 0x83, 0x30, 0xbe, 0x81, 0xc3, 0xfb, 0x79, 0xf2, 0xcb, 0xa7,
	0xef, 0x97, 0xe2, 0x05, 0x0c, 0xe2, 0x2a, 0xb2, 0x56, 0x15, 0x26, 0xbd, 0x51, 0xef, 0x78, 0x27,
	0xdb, 0x8d, 0xda, 0x85, 0xaa, 0x50, 0x1c, 0xc1, 0x76, 0x67, 0xcb, 0x60, 0x6f, 0xb0, 0xbd, 0xd5,
	0xd9, 0x92, 0xad, 0x57, 0xb0, 0xa7, 0x6b, 0x87, 0xb6, 0x54, 0x39, 0x06, 0xa0, 0xcf, 0xc0, 0x70,
	0xa5, 0x7a, 0x6c, 0xfc, 0x77, 0x03, 0xf6, 0x97, 0x05, 0x16, 0xa4, 0x24, 0xb5, 0x95, 0x78, 0x0d,
	0xfb, 0x73, 0x54, 0x05, 0x5a, 0x56, 0xdc, 0x75, 0xb3, 0xcc, 0x1e, 0x06, 0xf9, 0x2b, 0xa9, 0xab,
	0xeb, 0x06, 0xc5, 0x4b, 0xd8, 0x5b, 0xe3, 0xd4, 0x2c, 0x74, 0x18, 0x66, 0x83, 0x15, 0x76, 0x36,
	0x43, 0x31, 0x82, 0xc1, 0x8a, 0x92, 0xba, 0x88, 0x35, 0x60, 0xc9, 0x9c, 0x17, 0xe2, 0x03, 0x1c,
	0x45, 0x42, 0x15, 0x1d, 0x5a, 0xa7, 0x49, 0xd7, 0x33, 0x69, 0x4d, 0xeb, 0xd0, 0x26, 0x8f, 0x18,
	0x3f, 0x0c, 0xc0, 0xd9, 0xad, 0x9f, 0xb1, 0x2d, 0x26, 0xf0, 0x3c, 0xce, 0x25, 0xfc, 0xd9, 0x62,
	0xed, 0x37, 0xdb, 0x56, 0x53, 0xb4, 0xc9, 0x63, 0xee, 0x72, 0x10, 0xdc, 0xcb, 0x68, 0x5e, 0xb0,
	0x27, 0x52, 0x78, 0xb6, 0xd6, 0x3c, 0x9f, 0x63, 0xfe, 0x83, 0xda, 0x2a, 0xd9, 0xe4, 0x29, 0x4f,
	0x57, 0xf5, 0x3f, 0x46, 0x63, 0xfc, 0xa7, 0x0f, 0xe2, 0xfe, 0x35, 0x89, 0xb7, 0x70, 0x10, 0xfe,
	0xdd, 0x39, 0xe9, 0x13, 0xee, 0x2c, 0xd8, 0x3b, 0x5f, 0x3f, 0x6e, 0xf1, 0x06, 0x82, 0x2a, 0x1b,
	0x95, 0xfb, 0x4d, 0x3a, 0x5d, 0xa1, 0x4d, 0x4e, 0x39, 0xf7, 0x09, 0x3b, 0xdf, 0xd8, 0xb8, 0xf2,
	0xba, 0xbf, 0x88, 0x98, 0xe6, 0x5b, 0x9a, 0xb6, 0x76, 0xc9, 0x84, 0xd1, 0x21, 0xcb, 0xbe, 0xa1,
	0x17, 0xc5, 0xef, 0x1e, 0x80, 0xb2, 0xa8, 0x42, 0xb7, 0xe4, 0xdd, 0xa8, 0x7f, 0xbc, 0x7b, 0x92,
	0xa7, 0xff, 0xff, 0xb9, 0xa6, 0x77, 0x9e, 0x4a, 0xb6, 0xe3, 0x63, 0x3f, 0x7b, 0x4b, 0xdc, 0xc0,
	0xb6, 0xa2, 0xd8, 0xe0, 0xfd, 0xc3, 0x35, 0xd8, 0x52, 0xc4, 0xf9, 0xd3, 0x4d, 0xfe, 0x48, 0x4f,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x2f, 0x27, 0x38, 0xc0, 0x03, 0x00, 0x00,
}
