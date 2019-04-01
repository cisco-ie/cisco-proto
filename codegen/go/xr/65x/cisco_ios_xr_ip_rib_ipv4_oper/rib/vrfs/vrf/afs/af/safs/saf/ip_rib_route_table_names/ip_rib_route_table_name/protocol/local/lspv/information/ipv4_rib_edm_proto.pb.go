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

package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_local_lspv_information

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
	proto.RegisterType((*Ipv4RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.information.ipv4_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv4RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.information.ipv4_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv4_rib_edm_proto.proto", fileDescriptor_9d4183d9ee6d1994) }

var fileDescriptor_9d4183d9ee6d1994 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x14, 0x35, 0x61, 0x20, 0x15, 0xb8, 0x85, 0x3a, 0x70, 0xa0, 0x2a, 0x42, 0xca,
	0xc9, 0x42, 0x6d, 0xf8, 0x8f, 0xb8, 0x54, 0x9c, 0x10, 0x1c, 0x02, 0x17, 0x4e, 0x96, 0xd7, 0xf1,
	0x0a, 0x8b, 0x5d, 0x7b, 0xe5, 0xf1, 0xae, 0xe0, 0x19, 0xb8, 0xf1, 0xaa, 0xbc, 0x00, 0xf2, 0x78,
	0x37, 0x24, 0x44, 0x5c, 0x1c, 0x79, 0xbe, 0xdf, 0xcf, 0xf9, 0x26, 0x01, 0x6e, 0xdb, 0x7e, 0x29,
	0x83, 0x2d, 0xa5, 0x59, 0x37, 0xb2, 0x0d, 0x3e, 0x7a, 0x41, 0x27, 0xfb, 0x59, 0x68, 0x8b, 0xda,
	0x4b, 0xeb, 0x51, 0x7e, 0x0f, 0xd2, 0xb6, 0x44, 0x11, 0xee, 0x5b, 0x13, 0x44, 0xb0, 0xa5, 0xe8,
	0x43, 0x85, 0xe9, 0x10, 0xaa, 0x42, 0xa1, 0x2a, 0x81, 0xe9, 0x13, 0x55, 0x25, 0x06, 0x3a, 0xf8,
	0x2e, 0x1a, 0x19, 0x55, 0x59, 0x1b, 0xe9, 0x54, 0x63, 0xf0, 0x7f, 0x41, 0xfe, 0x4e, 0xed, 0x6b,
	0x51, 0x7b, 0xad, 0x6a, 0x51, 0x63, 0xdb, 0x0b, 0xeb, 0x2a, 0x1f, 0x1a, 0x15, 0xad, 0x77, 0xe7,
	0xbf, 0x0a, 0x38, 0xdd, 0xaf, 0x2a, 0xdf, 0xbf, 0xfb, 0xf2, 0x89, 0xcd, 0x61, 0xda, 0x87, 0x8a,
	0xde, 0xe1, 0xc5, 0x59, 0xb1, 0xb8, 0xb1, 0x9a, 0xf4, 0xa1, 0xfa, 0xa8, 0x1a, 0xc3, 0x4e, 0x61,
	0xa2, 0x86, 0xe4, 0x1a, 0x25, 0x87, 0x2a, 0x07, 0x73, 0x98, 0xe2, 0x98, 0x1c, 0x64, 0x07, 0x87,
	0x68, 0x01, 0xb7, 0xff, 0xad, 0xc7, 0xaf, 0x13, 0x72, 0x44, 0xf3, 0xcf, 0x69, 0x9c, 0xc8, 0xf3,
	0xdf, 0x07, 0xc0, 0xf6, 0x4b, 0xb1, 0xc7, 0x70, 0x34, 0xae, 0x93, 0xb7, 0xe6, 0x17, 0xa4, 0xcf,
	0xc6, 0x69, 0x92, 0x91, 0xdd, 0x87, 0xa9, 0x75, 0x18, 0x95, 0xd3, 0x86, 0x5f, 0x12, 0xb0, 0xb9,
	0x33, 0x0e, 0x93, 0xde, 0x04, 0xb4, 0xde, 0xf1, 0xe5, 0x59, 0xb1, 0x98, 0xad, 0xc6, 0x2b, 0x7b,
	0x0b, 0x0f, 0x82, 0x59, 0x5b, 0x8c, 0xc1, 0x96, 0x5d, 0xfa, 0x69, 0xa4, 0xae, 0xad, 0x71, 0x51,
	0x6a, 0xdf, 0xb9, 0xc8, 0x9f, 0x12, 0x3d, 0xdf, 0x45, 0xae, 0x88, 0xb8, 0x4a, 0x00, 0x5b, 0xc2,
	0xbd, 0x4d, 0xb9, 0x6c, 0xe2, 0xa0, 0x3e, 0x23, 0xf5, 0x64, 0x4c, 0xb3, 0x84, 0xd9, 0x7a, 0x04,
	0x33, 0xda, 0x7d, 0x60, 0x91, 0x3f, 0x27, 0xf8, 0x56, 0x1e, 0x12, 0x83, 0x4c, 0xc0, 0xb1, 0xd2,
	0xd1, 0xf6, 0x46, 0x6e, 0xb3, 0xfc, 0x05, 0xa1, 0x77, 0x72, 0xb4, 0xfa, 0x2b, 0xb0, 0x27, 0x70,
	0xb2, 0x36, 0xb5, 0x89, 0x66, 0xbd, 0x2b, 0xbc, 0x24, 0x81, 0x0d, 0xd9, 0xb6, 0xf1, 0x10, 0x6e,
	0xb6, 0x2a, 0x7e, 0x1d, 0xc1, 0x57, 0x04, 0x02, 0x8d, 0x32, 0x70, 0x01, 0x77, 0x37, 0xdb, 0xe5,
	0x3f, 0xb1, 0x31, 0x8d, 0x0f, 0x3f, 0xf8, 0x6b, 0x42, 0x8f, 0xc7, 0x90, 0x1e, 0xfd, 0x40, 0x51,
	0xaa, 0x5d, 0x2a, 0xfd, 0xad, 0x6b, 0x77, 0x5b, 0xbc, 0xc9, 0xb5, 0x73, 0xb4, 0x55, 0xa2, 0x3c,
	0xa4, 0x47, 0x2e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x62, 0xf0, 0xe9, 0x3b, 0x03, 0x00,
	0x00,
}
