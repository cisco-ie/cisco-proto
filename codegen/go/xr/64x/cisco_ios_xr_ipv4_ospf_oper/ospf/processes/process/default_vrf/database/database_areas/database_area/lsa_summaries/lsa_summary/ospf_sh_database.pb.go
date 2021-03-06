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
// source: ospf_sh_database.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_database_database_areas_database_area_lsa_summaries_lsa_summary

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

type OspfShDatabase_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	AreaId_1             uint32   `protobuf:"varint,3,opt,name=area_id_1,json=areaId1,proto3" json:"area_id_1,omitempty"`
	LsType               string   `protobuf:"bytes,4,opt,name=ls_type,json=lsType,proto3" json:"ls_type,omitempty"`
	LsId                 string   `protobuf:"bytes,5,opt,name=ls_id,json=lsId,proto3" json:"ls_id,omitempty"`
	AdvertisingRouter    string   `protobuf:"bytes,6,opt,name=advertising_router,json=advertisingRouter,proto3" json:"advertising_router,omitempty"`
	InterfaceName        string   `protobuf:"bytes,7,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDatabase_KEYS) Reset()         { *m = OspfShDatabase_KEYS{} }
func (m *OspfShDatabase_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShDatabase_KEYS) ProtoMessage()    {}
func (*OspfShDatabase_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9f78b00611f1589, []int{0}
}

func (m *OspfShDatabase_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDatabase_KEYS.Unmarshal(m, b)
}
func (m *OspfShDatabase_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDatabase_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShDatabase_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDatabase_KEYS.Merge(m, src)
}
func (m *OspfShDatabase_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShDatabase_KEYS.Size(m)
}
func (m *OspfShDatabase_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDatabase_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDatabase_KEYS proto.InternalMessageInfo

func (m *OspfShDatabase_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShDatabase_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShDatabase_KEYS) GetAreaId_1() uint32 {
	if m != nil {
		return m.AreaId_1
	}
	return 0
}

func (m *OspfShDatabase_KEYS) GetLsType() string {
	if m != nil {
		return m.LsType
	}
	return ""
}

func (m *OspfShDatabase_KEYS) GetLsId() string {
	if m != nil {
		return m.LsId
	}
	return ""
}

func (m *OspfShDatabase_KEYS) GetAdvertisingRouter() string {
	if m != nil {
		return m.AdvertisingRouter
	}
	return ""
}

func (m *OspfShDatabase_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type OspfShDbHeader struct {
	LsType               string   `protobuf:"bytes,1,opt,name=ls_type,json=lsType,proto3" json:"ls_type,omitempty"`
	Lsid                 string   `protobuf:"bytes,2,opt,name=lsid,proto3" json:"lsid,omitempty"`
	AdvertisingRouter    string   `protobuf:"bytes,3,opt,name=advertising_router,json=advertisingRouter,proto3" json:"advertising_router,omitempty"`
	LsaAreaId            string   `protobuf:"bytes,4,opt,name=lsa_area_id,json=lsaAreaId,proto3" json:"lsa_area_id,omitempty"`
	LsaAge               uint32   `protobuf:"varint,5,opt,name=lsa_age,json=lsaAge,proto3" json:"lsa_age,omitempty"`
	DnAgeLsa             bool     `protobuf:"varint,6,opt,name=dn_age_lsa,json=dnAgeLsa,proto3" json:"dn_age_lsa,omitempty"`
	Nsf                  bool     `protobuf:"varint,7,opt,name=nsf,proto3" json:"nsf,omitempty"`
	SequenceNumber       uint32   `protobuf:"varint,8,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	Checksum             uint32   `protobuf:"varint,9,opt,name=checksum,proto3" json:"checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDbHeader) Reset()         { *m = OspfShDbHeader{} }
func (m *OspfShDbHeader) String() string { return proto.CompactTextString(m) }
func (*OspfShDbHeader) ProtoMessage()    {}
func (*OspfShDbHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9f78b00611f1589, []int{1}
}

func (m *OspfShDbHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbHeader.Unmarshal(m, b)
}
func (m *OspfShDbHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbHeader.Marshal(b, m, deterministic)
}
func (m *OspfShDbHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbHeader.Merge(m, src)
}
func (m *OspfShDbHeader) XXX_Size() int {
	return xxx_messageInfo_OspfShDbHeader.Size(m)
}
func (m *OspfShDbHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDbHeader.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDbHeader proto.InternalMessageInfo

func (m *OspfShDbHeader) GetLsType() string {
	if m != nil {
		return m.LsType
	}
	return ""
}

func (m *OspfShDbHeader) GetLsid() string {
	if m != nil {
		return m.Lsid
	}
	return ""
}

func (m *OspfShDbHeader) GetAdvertisingRouter() string {
	if m != nil {
		return m.AdvertisingRouter
	}
	return ""
}

func (m *OspfShDbHeader) GetLsaAreaId() string {
	if m != nil {
		return m.LsaAreaId
	}
	return ""
}

func (m *OspfShDbHeader) GetLsaAge() uint32 {
	if m != nil {
		return m.LsaAge
	}
	return 0
}

func (m *OspfShDbHeader) GetDnAgeLsa() bool {
	if m != nil {
		return m.DnAgeLsa
	}
	return false
}

func (m *OspfShDbHeader) GetNsf() bool {
	if m != nil {
		return m.Nsf
	}
	return false
}

func (m *OspfShDbHeader) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *OspfShDbHeader) GetChecksum() uint32 {
	if m != nil {
		return m.Checksum
	}
	return 0
}

type OspfShDatabase struct {
	LsaHeader            *OspfShDbHeader `protobuf:"bytes,50,opt,name=lsa_header,json=lsaHeader,proto3" json:"lsa_header,omitempty"`
	ExternalTag          uint32          `protobuf:"varint,51,opt,name=external_tag,json=externalTag,proto3" json:"external_tag,omitempty"`
	LinkCount            uint32          `protobuf:"varint,52,opt,name=link_count,json=linkCount,proto3" json:"link_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OspfShDatabase) Reset()         { *m = OspfShDatabase{} }
func (m *OspfShDatabase) String() string { return proto.CompactTextString(m) }
func (*OspfShDatabase) ProtoMessage()    {}
func (*OspfShDatabase) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9f78b00611f1589, []int{2}
}

func (m *OspfShDatabase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDatabase.Unmarshal(m, b)
}
func (m *OspfShDatabase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDatabase.Marshal(b, m, deterministic)
}
func (m *OspfShDatabase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDatabase.Merge(m, src)
}
func (m *OspfShDatabase) XXX_Size() int {
	return xxx_messageInfo_OspfShDatabase.Size(m)
}
func (m *OspfShDatabase) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDatabase.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDatabase proto.InternalMessageInfo

func (m *OspfShDatabase) GetLsaHeader() *OspfShDbHeader {
	if m != nil {
		return m.LsaHeader
	}
	return nil
}

func (m *OspfShDatabase) GetExternalTag() uint32 {
	if m != nil {
		return m.ExternalTag
	}
	return 0
}

func (m *OspfShDatabase) GetLinkCount() uint32 {
	if m != nil {
		return m.LinkCount
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShDatabase_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_areas.database_area.lsa_summaries.lsa_summary.ospf_sh_database_KEYS")
	proto.RegisterType((*OspfShDbHeader)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_areas.database_area.lsa_summaries.lsa_summary.ospf_sh_db_header")
	proto.RegisterType((*OspfShDatabase)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_areas.database_area.lsa_summaries.lsa_summary.ospf_sh_database")
}

func init() { proto.RegisterFile("ospf_sh_database.proto", fileDescriptor_b9f78b00611f1589) }

var fileDescriptor_b9f78b00611f1589 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x9b, 0x34, 0x3f, 0x93, 0xa6, 0xb4, 0x8b, 0x00, 0xab, 0x02, 0x94, 0x46, 0x42, 0xe4,
	0x82, 0xa5, 0xfe, 0xbc, 0x40, 0x84, 0x90, 0xa8, 0x40, 0x3d, 0x84, 0x5e, 0x38, 0x8d, 0x26, 0xf6,
	0xc4, 0xb1, 0xba, 0xfe, 0x61, 0xc7, 0x8e, 0x9a, 0x0b, 0x67, 0xae, 0x3c, 0x00, 0x0f, 0xc1, 0xb3,
	0xf1, 0x02, 0x68, 0xd7, 0x71, 0x94, 0x16, 0xf5, 0xcc, 0x6d, 0xe6, 0x9b, 0xcf, 0xb3, 0xdf, 0xf7,
	0xed, 0x1a, 0x9e, 0xe7, 0x52, 0x2c, 0x50, 0x96, 0x18, 0x51, 0x49, 0x73, 0x12, 0x0e, 0x0a, 0x93,
	0x97, 0xb9, 0xfa, 0x1e, 0x26, 0x12, 0xe6, 0x98, 0xe4, 0x82, 0x77, 0x06, 0x93, 0x62, 0x75, 0x89,
	0x8e, 0x99, 0x17, 0x6c, 0x02, 0x5b, 0x59, 0x5e, 0xc8, 0x22, 0x2c, 0x4d, 0x15, 0x44, 0xbc, 0xa0,
	0x4a, 0x97, 0xb8, 0x32, 0x8b, 0x60, 0xbb, 0xae, 0x29, 0x90, 0x0c, 0x93, 0xdc, 0x6f, 0x03, 0x2d,
	0x84, 0x52, 0xa5, 0x29, 0x99, 0x84, 0x65, 0xa7, 0x5b, 0x8f, 0xff, 0x78, 0xf0, 0xec, 0xa1, 0x34,
	0xfc, 0xf4, 0xe1, 0xeb, 0x17, 0x75, 0x0a, 0x07, 0x9b, 0x03, 0x31, 0xa3, 0x94, 0x7d, 0x6f, 0xe4,
	0x4d, 0xfa, 0xb3, 0xc1, 0x06, 0xbb, 0xa6, 0x94, 0xd5, 0x0b, 0xe8, 0xda, 0xf5, 0x98, 0x44, 0xfe,
	0xde, 0xc8, 0x9b, 0x0c, 0x67, 0x1d, 0xdb, 0x5e, 0x45, 0xea, 0x04, 0xfa, 0x9b, 0x01, 0x9e, 0xf9,
	0x2d, 0x37, 0xea, 0xd6, 0xa3, 0x33, 0xfb, 0x91, 0x16, 0x2c, 0xd7, 0x05, 0xfb, 0x6d, 0xb7, 0xb2,
	0xa3, 0xe5, 0x66, 0x5d, 0xb0, 0x7a, 0x0a, 0xfb, 0x5a, 0xec, 0xae, 0x7d, 0x07, 0xb7, 0xb5, 0x5c,
	0x45, 0xea, 0x1d, 0x28, 0x8a, 0x56, 0x6c, 0xca, 0x44, 0x92, 0x2c, 0x46, 0x93, 0x57, 0x25, 0x1b,
	0xbf, 0xe3, 0x18, 0xc7, 0x3b, 0x93, 0x99, 0x1b, 0xa8, 0x37, 0x70, 0x98, 0x64, 0x25, 0x9b, 0x05,
	0x85, 0x5c, 0xcb, 0xee, 0x3a, 0xea, 0x70, 0x8b, 0x5a, 0xe1, 0xe3, 0x5f, 0x7b, 0x70, 0xbc, 0x75,
	0x3d, 0xc7, 0x25, 0x53, 0xc4, 0x66, 0x57, 0x99, 0x77, 0x4f, 0x99, 0x82, 0xb6, 0x96, 0x8d, 0x49,
	0x27, 0x2c, 0x79, 0x4c, 0x58, 0xeb, 0x31, 0x61, 0xaf, 0x61, 0x60, 0x63, 0x6f, 0xe2, 0xaa, 0x9d,
	0xf7, 0xb5, 0xd0, 0xb4, 0x4e, 0xcc, 0x9d, 0x4d, 0x48, 0x31, 0x3b, 0xfb, 0x43, 0x7b, 0x36, 0x4d,
	0x63, 0x56, 0x2f, 0x01, 0xa2, 0xcc, 0xe2, 0xa8, 0x85, 0x9c, 0xf1, 0xde, 0xac, 0x17, 0x65, 0xd3,
	0x98, 0x3f, 0x0b, 0xa9, 0x23, 0x68, 0x65, 0xb2, 0x70, 0x26, 0x7b, 0x33, 0x5b, 0xaa, 0xb7, 0xf0,
	0x44, 0xf8, 0x5b, 0xc5, 0x99, 0x0d, 0xa0, 0x4a, 0xe7, 0x6c, 0xfc, 0x9e, 0x5b, 0x78, 0xd8, 0xc0,
	0xd7, 0x0e, 0x55, 0x27, 0xd0, 0x0b, 0x97, 0x1c, 0xde, 0x4a, 0x95, 0xfa, 0x7d, 0xc7, 0xd8, 0xf6,
	0xe3, 0x1f, 0x7b, 0x70, 0xf4, 0xf0, 0x55, 0xa8, 0xdf, 0x1e, 0x80, 0xd5, 0x58, 0xa7, 0xe5, 0x9f,
	0x8f, 0xbc, 0xc9, 0xe0, 0xfc, 0xa7, 0x17, 0xfc, 0xdf, 0x17, 0x1c, 0xfc, 0x73, 0x8f, 0x2e, 0xd6,
	0x8f, 0xf5, 0x95, 0x9e, 0xc2, 0x01, 0xdf, 0x95, 0x6c, 0x32, 0xd2, 0x58, 0x52, 0xec, 0x5f, 0x38,
	0xa3, 0x83, 0x06, 0xbb, 0xa1, 0x58, 0xbd, 0x02, 0xd0, 0x49, 0x76, 0x8b, 0x61, 0x5e, 0x65, 0xa5,
	0x7f, 0xe9, 0x08, 0x7d, 0x8b, 0xbc, 0xb7, 0xc0, 0xbc, 0xe3, 0xfe, 0xd3, 0x8b, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x36, 0xd6, 0xb7, 0xfa, 0xc1, 0x03, 0x00, 0x00,
}
