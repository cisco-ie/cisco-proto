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
// source: ipv4_rib_edm_proto.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_srv6_local_non_as_information

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

type Ipv4RibEdmProto_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName       string   `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4RibEdmProto_KEYS) Reset()         { *m = Ipv4RibEdmProto_KEYS{} }
func (m *Ipv4RibEdmProto_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto_KEYS) ProtoMessage()    {}
func (*Ipv4RibEdmProto_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4183d9ee6d1994, []int{0}
}

func (m *Ipv4RibEdmProto_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmProto_KEYS.Unmarshal(m, b)
}
func (m *Ipv4RibEdmProto_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmProto_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmProto_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmProto_KEYS.Merge(m, src)
}
func (m *Ipv4RibEdmProto_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmProto_KEYS.Size(m)
}
func (m *Ipv4RibEdmProto_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmProto_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmProto_KEYS proto.InternalMessageInfo

func (m *Ipv4RibEdmProto_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

type Ipv4RibEdmProto struct {
	ProtocolNames             string   `protobuf:"bytes,50,opt,name=protocol_names,json=protocolNames,proto3" json:"protocol_names,omitempty"`
	Instance                  string   `protobuf:"bytes,51,opt,name=instance,proto3" json:"instance,omitempty"`
	Version                   uint32   `protobuf:"varint,52,opt,name=version,proto3" json:"version,omitempty"`
	RedistributionClientCount uint32   `protobuf:"varint,53,opt,name=redistribution_client_count,json=redistributionClientCount,proto3" json:"redistribution_client_count,omitempty"`
	ProtocolClientsCount      uint32   `protobuf:"varint,54,opt,name=protocol_clients_count,json=protocolClientsCount,proto3" json:"protocol_clients_count,omitempty"`
	RoutesCounts              uint32   `protobuf:"varint,55,opt,name=routes_counts,json=routesCounts,proto3" json:"routes_counts,omitempty"`
	ActiveRoutesCount         uint32   `protobuf:"varint,56,opt,name=active_routes_count,json=activeRoutesCount,proto3" json:"active_routes_count,omitempty"`
	DeletedRoutesCount        uint32   `protobuf:"varint,57,opt,name=deleted_routes_count,json=deletedRoutesCount,proto3" json:"deleted_routes_count,omitempty"`
	PathsCount                uint32   `protobuf:"varint,58,opt,name=paths_count,json=pathsCount,proto3" json:"paths_count,omitempty"`
	ProtocolRouteMemory       uint32   `protobuf:"varint,59,opt,name=protocol_route_memory,json=protocolRouteMemory,proto3" json:"protocol_route_memory,omitempty"`
	BackupRoutesCount         uint32   `protobuf:"varint,60,opt,name=backup_routes_count,json=backupRoutesCount,proto3" json:"backup_routes_count,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Ipv4RibEdmProto) Reset()         { *m = Ipv4RibEdmProto{} }
func (m *Ipv4RibEdmProto) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto) ProtoMessage()    {}
func (*Ipv4RibEdmProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4183d9ee6d1994, []int{1}
}

func (m *Ipv4RibEdmProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmProto.Unmarshal(m, b)
}
func (m *Ipv4RibEdmProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmProto.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmProto.Merge(m, src)
}
func (m *Ipv4RibEdmProto) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmProto.Size(m)
}
func (m *Ipv4RibEdmProto) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmProto.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmProto proto.InternalMessageInfo

func (m *Ipv4RibEdmProto) GetProtocolNames() string {
	if m != nil {
		return m.ProtocolNames
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRedistributionClientCount() uint32 {
	if m != nil {
		return m.RedistributionClientCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolClientsCount() uint32 {
	if m != nil {
		return m.ProtocolClientsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRoutesCounts() uint32 {
	if m != nil {
		return m.RoutesCounts
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetDeletedRoutesCount() uint32 {
	if m != nil {
		return m.DeletedRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolRouteMemory() uint32 {
	if m != nil {
		return m.ProtocolRouteMemory
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetBackupRoutesCount() uint32 {
	if m != nil {
		return m.BackupRoutesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.srv6_local.non_as.information.ipv4_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv4RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.srv6_local.non_as.information.ipv4_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv4_rib_edm_proto.proto", fileDescriptor_9d4183d9ee6d1994) }

var fileDescriptor_9d4183d9ee6d1994 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x14, 0x35, 0x61, 0x20, 0x15, 0xb8, 0x85, 0x3a, 0x70, 0xa0, 0x2a, 0x42, 0xca,
	0xc9, 0x42, 0x6d, 0x28, 0x7f, 0xc5, 0xa5, 0xe2, 0x84, 0xe0, 0x10, 0xb8, 0x70, 0xb2, 0xbc, 0x8e,
	0x57, 0x58, 0xec, 0xda, 0x2b, 0x8f, 0x77, 0x05, 0xaf, 0xd1, 0x57, 0xe5, 0x05, 0x90, 0xc7, 0xbb,
	0x21, 0xa1, 0xea, 0xc5, 0x91, 0xe7, 0xfb, 0xfd, 0x9c, 0x6f, 0x12, 0xe0, 0xb6, 0xed, 0x97, 0x32,
	0xd8, 0x52, 0x9a, 0x75, 0x23, 0xdb, 0xe0, 0xa3, 0x17, 0x74, 0xb2, 0xab, 0x42, 0x5b, 0xd4, 0x5e,
	0x5a, 0x8f, 0xf2, 0x57, 0x90, 0xb6, 0x25, 0x8a, 0x70, 0xdf, 0x9a, 0x20, 0x82, 0x2d, 0x45, 0x1f,
	0x2a, 0x4c, 0x87, 0x50, 0x15, 0x0a, 0x55, 0x09, 0x4c, 0x9f, 0xa8, 0x2a, 0x31, 0xd0, 0xc1, 0x77,
	0xd1, 0xc8, 0xa8, 0xca, 0xda, 0x48, 0xa7, 0x1a, 0x83, 0x37, 0x05, 0xf9, 0x3b, 0xb5, 0xaf, 0x05,
	0x86, 0xfe, 0x42, 0xd6, 0x5e, 0xab, 0x5a, 0x38, 0xef, 0xa4, 0x42, 0x61, 0x5d, 0xe5, 0x43, 0xa3,
	0xa2, 0xf5, 0xee, 0xf4, 0xaa, 0x80, 0xe3, 0xeb, 0x8d, 0xe5, 0xa7, 0x8f, 0xdf, 0xbf, 0xb2, 0x39,
	0x4c, 0xfb, 0x50, 0xd1, 0x73, 0xbc, 0x38, 0x29, 0x16, 0x77, 0x56, 0x93, 0x3e, 0x54, 0x5f, 0x54,
	0x63, 0xd8, 0x31, 0x4c, 0xd4, 0x90, 0xdc, 0xa2, 0x64, 0x5f, 0xe5, 0x60, 0x0e, 0x53, 0x1c, 0x93,
	0xbd, 0xec, 0xe0, 0x10, 0x2d, 0xe0, 0xfe, 0xff, 0x2d, 0xf9, 0x6d, 0x42, 0x0e, 0x68, 0xfe, 0x2d,
	0x8d, 0x13, 0x79, 0xfa, 0x67, 0x0f, 0xd8, 0xf5, 0x52, 0xec, 0x39, 0x1c, 0x8c, 0x5b, 0xe5, 0xe5,
	0xf9, 0x19, 0xe9, 0xb3, 0x71, 0x9a, 0x64, 0x64, 0x8f, 0x61, 0x6a, 0x1d, 0x46, 0xe5, 0xb4, 0xe1,
	0xe7, 0x04, 0x6c, 0xee, 0x8c, 0xc3, 0xa4, 0x37, 0x01, 0xad, 0x77, 0x7c, 0x79, 0x52, 0x2c, 0x66,
	0xab, 0xf1, 0xca, 0x3e, 0xc0, 0x93, 0x60, 0xd6, 0x16, 0x63, 0xb0, 0x65, 0x97, 0x7e, 0x1a, 0xa9,
	0x6b, 0x6b, 0x5c, 0x94, 0xda, 0x77, 0x2e, 0xf2, 0x97, 0x44, 0xcf, 0x77, 0x91, 0x4b, 0x22, 0x2e,
	0x13, 0xc0, 0x96, 0xf0, 0x68, 0x53, 0x2e, 0x9b, 0x38, 0xa8, 0x17, 0xa4, 0x1e, 0x8d, 0x69, 0x96,
	0x30, 0x5b, 0xcf, 0x60, 0x46, 0xbb, 0x0f, 0x2c, 0xf2, 0x57, 0x04, 0xdf, 0xcb, 0x43, 0x62, 0x90,
	0x09, 0x38, 0x54, 0x3a, 0xda, 0xde, 0xc8, 0x6d, 0x96, 0xbf, 0x26, 0xf4, 0x41, 0x8e, 0x56, 0xff,
	0x04, 0xf6, 0x02, 0x8e, 0xd6, 0xa6, 0x36, 0xd1, 0xac, 0x77, 0x85, 0x37, 0x24, 0xb0, 0x21, 0xdb,
	0x36, 0x9e, 0xc2, 0xdd, 0x56, 0xc5, 0x1f, 0x23, 0xf8, 0x96, 0x40, 0xa0, 0x51, 0x06, 0xce, 0xe0,
	0xe1, 0x66, 0xbb, 0xfc, 0x27, 0x36, 0xa6, 0xf1, 0xe1, 0x37, 0x7f, 0x47, 0xe8, 0xe1, 0x18, 0xd2,
	0xa3, 0x9f, 0x29, 0x4a, 0xb5, 0x4b, 0xa5, 0x7f, 0x76, 0xed, 0x6e, 0x8b, 0xf7, 0xb9, 0x76, 0x8e,
	0xb6, 0x4a, 0x94, 0xfb, 0xf4, 0xc8, 0xf9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x89, 0xd0,
	0x63, 0x42, 0x03, 0x00, 0x00,
}
