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
// source: ipv6_rib_edm_proto.proto

package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_local_lspv_information

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

type Ipv6RibEdmProto_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName       string   `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6RibEdmProto_KEYS) Reset()         { *m = Ipv6RibEdmProto_KEYS{} }
func (m *Ipv6RibEdmProto_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmProto_KEYS) ProtoMessage()    {}
func (*Ipv6RibEdmProto_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cb6e91b99f67ad, []int{0}
}

func (m *Ipv6RibEdmProto_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmProto_KEYS.Unmarshal(m, b)
}
func (m *Ipv6RibEdmProto_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmProto_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6RibEdmProto_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmProto_KEYS.Merge(m, src)
}
func (m *Ipv6RibEdmProto_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmProto_KEYS.Size(m)
}
func (m *Ipv6RibEdmProto_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmProto_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmProto_KEYS proto.InternalMessageInfo

func (m *Ipv6RibEdmProto_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6RibEdmProto_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv6RibEdmProto_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv6RibEdmProto_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

type Ipv6RibEdmProto struct {
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

func (m *Ipv6RibEdmProto) Reset()         { *m = Ipv6RibEdmProto{} }
func (m *Ipv6RibEdmProto) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmProto) ProtoMessage()    {}
func (*Ipv6RibEdmProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cb6e91b99f67ad, []int{1}
}

func (m *Ipv6RibEdmProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmProto.Unmarshal(m, b)
}
func (m *Ipv6RibEdmProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmProto.Marshal(b, m, deterministic)
}
func (m *Ipv6RibEdmProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmProto.Merge(m, src)
}
func (m *Ipv6RibEdmProto) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmProto.Size(m)
}
func (m *Ipv6RibEdmProto) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmProto.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmProto proto.InternalMessageInfo

func (m *Ipv6RibEdmProto) GetProtocolNames() string {
	if m != nil {
		return m.ProtocolNames
	}
	return ""
}

func (m *Ipv6RibEdmProto) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv6RibEdmProto) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetRedistributionClientCount() uint32 {
	if m != nil {
		return m.RedistributionClientCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetProtocolClientsCount() uint32 {
	if m != nil {
		return m.ProtocolClientsCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetRoutesCounts() uint32 {
	if m != nil {
		return m.RoutesCounts
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetDeletedRoutesCount() uint32 {
	if m != nil {
		return m.DeletedRoutesCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetProtocolRouteMemory() uint32 {
	if m != nil {
		return m.ProtocolRouteMemory
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetBackupRoutesCount() uint32 {
	if m != nil {
		return m.BackupRoutesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.information.ipv6_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv6RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.information.ipv6_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv6_rib_edm_proto.proto", fileDescriptor_63cb6e91b99f67ad) }

var fileDescriptor_63cb6e91b99f67ad = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x14, 0x35, 0x61, 0x20, 0x15, 0xb8, 0x85, 0x3a, 0x70, 0xa0, 0x2a, 0x42, 0xca,
	0xc9, 0x42, 0x6d, 0x29, 0x7f, 0xc5, 0xa5, 0xe2, 0x84, 0xe0, 0x10, 0xb8, 0x70, 0xb2, 0xbc, 0x8e,
	0x57, 0x58, 0xec, 0xae, 0x57, 0x1e, 0xef, 0x0a, 0xde, 0x02, 0xf1, 0xaa, 0xbc, 0x00, 0xf2, 0x78,
	0x1d, 0x12, 0xa2, 0x5e, 0x1c, 0x79, 0xbe, 0xdf, 0xcf, 0xf9, 0x26, 0x01, 0x6e, 0xbb, 0xe1, 0x52,
	0x7a, 0x5b, 0x4a, 0xb3, 0x6a, 0x64, 0xe7, 0x5d, 0x70, 0x82, 0x4e, 0xf6, 0xab, 0xd0, 0x16, 0xb5,
	0x93, 0xd6, 0xa1, 0xfc, 0xe1, 0xa5, 0xed, 0x88, 0x22, 0xdc, 0x75, 0xc6, 0x8b, 0x2c, 0x8a, 0xc1,
	0x57, 0x18, 0x0f, 0xa1, 0x2a, 0x14, 0xaa, 0x12, 0x18, 0x3f, 0x51, 0x55, 0x62, 0x54, 0xbc, 0xeb,
	0x83, 0x91, 0x41, 0x95, 0xb5, 0x91, 0xad, 0x6a, 0x0c, 0x5e, 0x17, 0xa4, 0x2f, 0xd6, 0xae, 0x16,
	0xb5, 0xd3, 0xaa, 0x16, 0x35, 0x76, 0x83, 0xb0, 0x6d, 0xe5, 0x7c, 0xa3, 0x82, 0x75, 0xed, 0xe9,
	0xef, 0x02, 0x8e, 0x77, 0xfb, 0xca, 0x0f, 0xef, 0xbf, 0x7e, 0x66, 0x73, 0x98, 0x0e, 0xbe, 0xa2,
	0x77, 0x78, 0x71, 0x52, 0x2c, 0x6e, 0x2d, 0x27, 0x83, 0xaf, 0x3e, 0xa9, 0xc6, 0xb0, 0x63, 0x98,
	0xa8, 0x31, 0xb9, 0x41, 0xc9, 0xbe, 0x4a, 0xc1, 0x1c, 0xa6, 0x98, 0x93, 0xbd, 0xe4, 0xe0, 0x18,
	0x2d, 0xe0, 0xee, 0xff, 0xf5, 0xf8, 0x4d, 0x42, 0x0e, 0x68, 0xfe, 0x25, 0x8e, 0x23, 0x79, 0xfa,
	0x67, 0x0f, 0xd8, 0x6e, 0x29, 0xf6, 0x14, 0x0e, 0xf2, 0x3a, 0x69, 0x6b, 0x7e, 0x46, 0xfa, 0x2c,
	0x4f, 0xa3, 0x8c, 0xec, 0x21, 0x4c, 0x6d, 0x8b, 0x41, 0xb5, 0xda, 0xf0, 0x73, 0x02, 0xd6, 0x77,
	0xc6, 0x61, 0x32, 0x18, 0x8f, 0xd6, 0xb5, 0xfc, 0xe2, 0xa4, 0x58, 0xcc, 0x96, 0xf9, 0xca, 0xde,
	0xc1, 0x23, 0x6f, 0x56, 0x16, 0x83, 0xb7, 0x65, 0x1f, 0x7f, 0x1a, 0xa9, 0x6b, 0x6b, 0xda, 0x20,
	0xb5, 0xeb, 0xdb, 0xc0, 0x9f, 0x13, 0x3d, 0xdf, 0x46, 0xae, 0x88, 0xb8, 0x8a, 0x00, 0xbb, 0x80,
	0x07, 0xeb, 0x72, 0xc9, 0xc4, 0x51, 0xbd, 0x24, 0xf5, 0x28, 0xa7, 0x49, 0xc2, 0x64, 0x3d, 0x81,
	0x19, 0xed, 0x3e, 0xb2, 0xc8, 0x5f, 0x10, 0x7c, 0x27, 0x0d, 0x89, 0x41, 0x26, 0xe0, 0x50, 0xe9,
	0x60, 0x07, 0x23, 0x37, 0x59, 0xfe, 0x92, 0xd0, 0x7b, 0x29, 0x5a, 0xfe, 0x13, 0xd8, 0x33, 0x38,
	0x5a, 0x99, 0xda, 0x04, 0xb3, 0xda, 0x16, 0x5e, 0x91, 0xc0, 0xc6, 0x6c, 0xd3, 0x78, 0x0c, 0xb7,
	0x3b, 0x15, 0xbe, 0x65, 0xf0, 0x35, 0x81, 0x40, 0xa3, 0x04, 0x9c, 0xc1, 0xfd, 0xf5, 0x76, 0xe9,
	0x4f, 0x6c, 0x4c, 0xe3, 0xfc, 0x4f, 0xfe, 0x86, 0xd0, 0xc3, 0x1c, 0xd2, 0xa3, 0x1f, 0x29, 0x8a,
	0xb5, 0x4b, 0xa5, 0xbf, 0xf7, 0xdd, 0x76, 0x8b, 0xb7, 0xa9, 0x76, 0x8a, 0x36, 0x4a, 0x94, 0xfb,
	0xf4, 0xc8, 0xf9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x59, 0xd3, 0xae, 0x40, 0x03, 0x00,
	0x00,
}
