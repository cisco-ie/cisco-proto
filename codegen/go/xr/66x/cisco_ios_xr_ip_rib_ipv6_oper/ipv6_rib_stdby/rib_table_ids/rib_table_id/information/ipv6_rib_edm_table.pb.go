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
// source: ipv6_rib_edm_table.proto

package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_rib_table_ids_rib_table_id_information

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

type Ipv6RibEdmTable_KEYS struct {
	Tableid              string   `protobuf:"bytes,1,opt,name=tableid,proto3" json:"tableid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6RibEdmTable_KEYS) Reset()         { *m = Ipv6RibEdmTable_KEYS{} }
func (m *Ipv6RibEdmTable_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmTable_KEYS) ProtoMessage()    {}
func (*Ipv6RibEdmTable_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f6744f6153a6679, []int{0}
}

func (m *Ipv6RibEdmTable_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmTable_KEYS.Unmarshal(m, b)
}
func (m *Ipv6RibEdmTable_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmTable_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6RibEdmTable_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmTable_KEYS.Merge(m, src)
}
func (m *Ipv6RibEdmTable_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmTable_KEYS.Size(m)
}
func (m *Ipv6RibEdmTable_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmTable_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmTable_KEYS proto.InternalMessageInfo

func (m *Ipv6RibEdmTable_KEYS) GetTableid() string {
	if m != nil {
		return m.Tableid
	}
	return ""
}

type Ipv6RibEdmTable struct {
	Tableid              uint32   `protobuf:"varint,50,opt,name=tableid,proto3" json:"tableid,omitempty"`
	Afi                  uint32   `protobuf:"varint,51,opt,name=afi,proto3" json:"afi,omitempty"`
	Safi                 uint32   `protobuf:"varint,52,opt,name=safi,proto3" json:"safi,omitempty"`
	VrfName              string   `protobuf:"bytes,53,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	TableName            string   `protobuf:"bytes,54,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Version              uint64   `protobuf:"varint,55,opt,name=version,proto3" json:"version,omitempty"`
	ConfPrefixLimit      uint32   `protobuf:"varint,56,opt,name=conf_prefix_limit,json=confPrefixLimit,proto3" json:"conf_prefix_limit,omitempty"`
	CurrentPrefixCount   uint32   `protobuf:"varint,57,opt,name=current_prefix_count,json=currentPrefixCount,proto3" json:"current_prefix_count,omitempty"`
	NumSvdlclPrefix      uint32   `protobuf:"varint,58,opt,name=num_svdlcl_prefix,json=numSvdlclPrefix,proto3" json:"num_svdlcl_prefix,omitempty"`
	NumSvdremPrefix      uint32   `protobuf:"varint,59,opt,name=num_svdrem_prefix,json=numSvdremPrefix,proto3" json:"num_svdrem_prefix,omitempty"`
	TableVersion         uint64   `protobuf:"varint,60,opt,name=table_version,json=tableVersion,proto3" json:"table_version,omitempty"`
	PrefixLimitNotified  bool     `protobuf:"varint,61,opt,name=prefix_limit_notified,json=prefixLimitNotified,proto3" json:"prefix_limit_notified,omitempty"`
	FwdReferenced        bool     `protobuf:"varint,62,opt,name=fwd_referenced,json=fwdReferenced,proto3" json:"fwd_referenced,omitempty"`
	Deleted              bool     `protobuf:"varint,63,opt,name=deleted,proto3" json:"deleted,omitempty"`
	InitialConverge      bool     `protobuf:"varint,64,opt,name=initial_converge,json=initialConverge,proto3" json:"initial_converge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6RibEdmTable) Reset()         { *m = Ipv6RibEdmTable{} }
func (m *Ipv6RibEdmTable) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmTable) ProtoMessage()    {}
func (*Ipv6RibEdmTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f6744f6153a6679, []int{1}
}

func (m *Ipv6RibEdmTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmTable.Unmarshal(m, b)
}
func (m *Ipv6RibEdmTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmTable.Marshal(b, m, deterministic)
}
func (m *Ipv6RibEdmTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmTable.Merge(m, src)
}
func (m *Ipv6RibEdmTable) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmTable.Size(m)
}
func (m *Ipv6RibEdmTable) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmTable.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmTable proto.InternalMessageInfo

func (m *Ipv6RibEdmTable) GetTableid() uint32 {
	if m != nil {
		return m.Tableid
	}
	return 0
}

func (m *Ipv6RibEdmTable) GetAfi() uint32 {
	if m != nil {
		return m.Afi
	}
	return 0
}

func (m *Ipv6RibEdmTable) GetSafi() uint32 {
	if m != nil {
		return m.Safi
	}
	return 0
}

func (m *Ipv6RibEdmTable) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6RibEdmTable) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *Ipv6RibEdmTable) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv6RibEdmTable) GetConfPrefixLimit() uint32 {
	if m != nil {
		return m.ConfPrefixLimit
	}
	return 0
}

func (m *Ipv6RibEdmTable) GetCurrentPrefixCount() uint32 {
	if m != nil {
		return m.CurrentPrefixCount
	}
	return 0
}

func (m *Ipv6RibEdmTable) GetNumSvdlclPrefix() uint32 {
	if m != nil {
		return m.NumSvdlclPrefix
	}
	return 0
}

func (m *Ipv6RibEdmTable) GetNumSvdremPrefix() uint32 {
	if m != nil {
		return m.NumSvdremPrefix
	}
	return 0
}

func (m *Ipv6RibEdmTable) GetTableVersion() uint64 {
	if m != nil {
		return m.TableVersion
	}
	return 0
}

func (m *Ipv6RibEdmTable) GetPrefixLimitNotified() bool {
	if m != nil {
		return m.PrefixLimitNotified
	}
	return false
}

func (m *Ipv6RibEdmTable) GetFwdReferenced() bool {
	if m != nil {
		return m.FwdReferenced
	}
	return false
}

func (m *Ipv6RibEdmTable) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *Ipv6RibEdmTable) GetInitialConverge() bool {
	if m != nil {
		return m.InitialConverge
	}
	return false
}

func init() {
	proto.RegisterType((*Ipv6RibEdmTable_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.rib_table_ids.rib_table_id.information.ipv6_rib_edm_table_KEYS")
	proto.RegisterType((*Ipv6RibEdmTable)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.rib_table_ids.rib_table_id.information.ipv6_rib_edm_table")
}

func init() { proto.RegisterFile("ipv6_rib_edm_table.proto", fileDescriptor_6f6744f6153a6679) }

var fileDescriptor_6f6744f6153a6679 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x09, 0x16, 0x77, 0xf7, 0x62, 0xdd, 0x75, 0x54, 0x1c, 0x1f, 0x84, 0xb2, 0x22, 0x54,
	0x1f, 0x8a, 0x6c, 0x75, 0xfd, 0xff, 0x07, 0x16, 0x9f, 0x94, 0x45, 0x5a, 0x10, 0x7c, 0x1a, 0xd2,
	0xcc, 0x1d, 0xb9, 0x90, 0x99, 0x09, 0x37, 0xd3, 0xec, 0xfa, 0x79, 0xfc, 0xa2, 0x92, 0x9b, 0xa4,
	0x46, 0xfa, 0x76, 0xef, 0xef, 0x9c, 0x33, 0xb7, 0xa7, 0x04, 0x34, 0x55, 0xcd, 0xb9, 0x61, 0xda,
	0x18, 0xb4, 0xde, 0xa4, 0x7c, 0x53, 0xe2, 0xa2, 0xe2, 0x98, 0xa2, 0x5a, 0x17, 0x54, 0x17, 0xd1,
	0x50, 0xac, 0xcd, 0x35, 0x1b, 0xaa, 0xc4, 0x24, 0xee, 0x58, 0x21, 0x2f, 0x76, 0xb9, 0x3a, 0xd9,
	0xcd, 0xef, 0x45, 0x3b, 0x49, 0xda, 0x90, 0xad, 0xff, 0xdb, 0x16, 0x14, 0x5c, 0x64, 0x9f, 0x27,
	0x8a, 0xe1, 0x74, 0x09, 0x0f, 0xf6, 0x0f, 0x9a, 0xaf, 0x5f, 0x7e, 0xae, 0x95, 0x86, 0x03, 0xd9,
	0xc8, 0xea, 0x6c, 0x96, 0xcd, 0x8f, 0x56, 0xc3, 0x7a, 0xfa, 0x67, 0x02, 0x6a, 0x3f, 0x35, 0x0e,
	0x9c, 0xcd, 0xb2, 0xf9, 0x74, 0x17, 0x50, 0x27, 0x70, 0x23, 0x77, 0xa4, 0x97, 0x42, 0xdb, 0x51,
	0x29, 0x98, 0xd4, 0x2d, 0x7a, 0x21, 0x48, 0x66, 0xf5, 0x10, 0x0e, 0x1b, 0x76, 0x26, 0xe4, 0x1e,
	0xf5, 0xcb, 0xee, 0x62, 0xc3, 0xee, 0x32, 0xf7, 0xa8, 0x1e, 0x01, 0x74, 0xbf, 0x4c, 0xc4, 0x73,
	0x11, 0x8f, 0x84, 0x88, 0xac, 0xe1, 0xa0, 0x41, 0xae, 0x29, 0x06, 0xfd, 0x6a, 0x96, 0xcd, 0x27,
	0xab, 0x61, 0x55, 0xcf, 0xe0, 0x4e, 0x11, 0x83, 0x33, 0x15, 0xa3, 0xa3, 0x6b, 0x53, 0x92, 0xa7,
	0xa4, 0x5f, 0xcb, 0xd1, 0xe3, 0x56, 0xf8, 0x2e, 0xfc, 0x5b, 0x8b, 0xd5, 0x73, 0xb8, 0x57, 0x6c,
	0x99, 0x31, 0xa4, 0xc1, 0x5e, 0xc4, 0x6d, 0x48, 0xfa, 0x8d, 0xd8, 0x55, 0xaf, 0x75, 0x89, 0x8b,
	0x56, 0x69, 0x5f, 0x0f, 0x5b, 0x6f, 0xea, 0xc6, 0x96, 0x45, 0xd9, 0x87, 0xf4, 0xdb, 0xee, 0xf5,
	0xb0, 0xf5, 0x6b, 0xe1, 0x5d, 0x60, 0xe4, 0x65, 0xf4, 0x83, 0xf7, 0xdd, 0xd8, 0xcb, 0xe8, 0x7b,
	0xef, 0x63, 0x98, 0x76, 0x75, 0x87, 0x56, 0xef, 0xa5, 0xd5, 0x2d, 0x81, 0x3f, 0xfa, 0x6a, 0x67,
	0x70, 0x7f, 0xdc, 0xca, 0x84, 0x98, 0xc8, 0x11, 0x5a, 0xfd, 0x61, 0x96, 0xcd, 0x0f, 0x57, 0x77,
	0xab, 0x7f, 0xd5, 0x2e, 0x7b, 0x49, 0x3d, 0x81, 0xdb, 0xee, 0xca, 0x1a, 0x46, 0x87, 0x8c, 0xa1,
	0x40, 0xab, 0x3f, 0x8a, 0x79, 0xea, 0xae, 0xec, 0x6a, 0x07, 0xdb, 0xff, 0xd3, 0x62, 0x89, 0x09,
	0xad, 0xfe, 0x24, 0xfa, 0xb0, 0xaa, 0xa7, 0x70, 0x42, 0x81, 0x12, 0xe5, 0xa5, 0x29, 0x62, 0x68,
	0x90, 0x7f, 0xa1, 0xfe, 0x2c, 0x96, 0xe3, 0x9e, 0x5f, 0xf4, 0x78, 0x73, 0x53, 0x3e, 0xdb, 0xe5,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x2f, 0xd3, 0x6c, 0xd2, 0x02, 0x00, 0x00,
}