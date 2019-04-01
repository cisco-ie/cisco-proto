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

package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information

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
	As                   string   `protobuf:"bytes,5,opt,name=as,proto3" json:"as,omitempty"`
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

func (m *Ipv6RibEdmProto_KEYS) GetAs() string {
	if m != nil {
		return m.As
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
	proto.RegisterType((*Ipv6RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.bgp.as.information.ipv6_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv6RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.bgp.as.information.ipv6_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv6_rib_edm_proto.proto", fileDescriptor_63cb6e91b99f67ad) }

var fileDescriptor_63cb6e91b99f67ad = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0x14, 0x9a, 0x30, 0x90, 0x08, 0xb6, 0x85, 0x6e, 0xe0, 0x40, 0x55, 0x84, 0x94,
	0xd3, 0x0a, 0xb5, 0xa5, 0xfc, 0x15, 0x97, 0x8a, 0x13, 0x82, 0x43, 0xe0, 0xc2, 0x69, 0x35, 0x76,
	0xd6, 0xb0, 0x22, 0xf6, 0x5a, 0x3b, 0x6b, 0x0b, 0xde, 0x81, 0xa7, 0xe0, 0x15, 0x79, 0x01, 0xe4,
	0x59, 0x6f, 0x48, 0xa8, 0xb8, 0x38, 0xf2, 0x7c, 0xbf, 0xdf, 0xe4, 0xdb, 0x35, 0x48, 0xdb, 0x74,
	0x17, 0xda, 0xdb, 0x5c, 0x9b, 0x55, 0xa5, 0x1b, 0xef, 0x82, 0x53, 0xfc, 0x14, 0x3f, 0xb3, 0xc2,
	0x52, 0xe1, 0xb4, 0x75, 0xa4, 0xbf, 0x7b, 0x6d, 0x1b, 0xa6, 0x18, 0x77, 0x8d, 0xf1, 0x2a, 0x89,
	0xaa, 0xf3, 0x25, 0xf5, 0x0f, 0x85, 0x25, 0x29, 0x2c, 0x15, 0xf5, 0xbf, 0x84, 0xa5, 0x1a, 0x14,
	0xef, 0xda, 0x60, 0x74, 0xc0, 0x7c, 0x6d, 0x74, 0x8d, 0x95, 0xa1, 0xff, 0x05, 0xf1, 0x8f, 0x0b,
	0xb7, 0x56, 0xf9, 0x97, 0x46, 0x21, 0x29, 0x5b, 0x97, 0xce, 0x57, 0x18, 0xac, 0xab, 0x4f, 0x7e,
	0x65, 0x70, 0x74, 0xb5, 0xab, 0x7e, 0xf7, 0xf6, 0xf3, 0x47, 0x31, 0x87, 0x49, 0xe7, 0x4b, 0xde,
	0x21, 0xb3, 0xe3, 0x6c, 0x71, 0x63, 0x39, 0xee, 0x7c, 0xf9, 0x01, 0x2b, 0x23, 0x8e, 0x60, 0x8c,
	0x43, 0x32, 0xe2, 0x64, 0x1f, 0x63, 0x30, 0x87, 0x09, 0xa5, 0x64, 0x2f, 0x3a, 0x34, 0x44, 0x0b,
	0xb8, 0xfd, 0x6f, 0x35, 0x79, 0x8d, 0x91, 0x19, 0xcf, 0x3f, 0xf5, 0x63, 0x26, 0x67, 0x30, 0x42,
	0x92, 0xd7, 0x39, 0x1b, 0x21, 0x9d, 0xfc, 0xde, 0x03, 0x71, 0xb5, 0xa4, 0x78, 0x0c, 0xb3, 0x74,
	0xb4, 0x78, 0x03, 0xf2, 0x94, 0x95, 0x69, 0x9a, 0xf6, 0xcb, 0x48, 0xdc, 0x87, 0x89, 0xad, 0x29,
	0x60, 0x5d, 0x18, 0x79, 0xc6, 0xc0, 0xe6, 0x5d, 0x48, 0x18, 0x77, 0xc6, 0x93, 0x75, 0xb5, 0x3c,
	0x3f, 0xce, 0x16, 0xd3, 0x65, 0x7a, 0x15, 0x6f, 0xe0, 0x81, 0x37, 0x2b, 0x4b, 0xc1, 0xdb, 0xbc,
	0xed, 0xaf, 0x4a, 0x17, 0x6b, 0x6b, 0xea, 0xa0, 0x0b, 0xd7, 0xd6, 0x41, 0x3e, 0x65, 0x7a, 0xbe,
	0x8b, 0x5c, 0x32, 0x71, 0xd9, 0x03, 0xe2, 0x1c, 0xee, 0x6d, 0xca, 0x45, 0x93, 0x06, 0xf5, 0x82,
	0xd5, 0xc3, 0x94, 0x46, 0x89, 0xa2, 0xf5, 0x08, 0xa6, 0x7c, 0x17, 0x03, 0x4b, 0xf2, 0x19, 0xc3,
	0xb7, 0xe2, 0x90, 0x19, 0x12, 0x0a, 0x0e, 0xb0, 0x08, 0xb6, 0x33, 0x7a, 0x9b, 0x95, 0xcf, 0x19,
	0xbd, 0x13, 0xa3, 0xe5, 0x5f, 0x41, 0x3c, 0x81, 0xc3, 0x95, 0x59, 0x9b, 0x60, 0x56, 0xbb, 0xc2,
	0x0b, 0x16, 0xc4, 0x90, 0x6d, 0x1b, 0x0f, 0xe1, 0x66, 0x83, 0xe1, 0x6b, 0x02, 0x5f, 0x32, 0x08,
	0x3c, 0x8a, 0xc0, 0x29, 0xdc, 0xdd, 0x9c, 0x2e, 0x7e, 0xd4, 0xca, 0x54, 0xce, 0xff, 0x90, 0xaf,
	0x18, 0x3d, 0x48, 0x21, 0x2f, 0x7d, 0xcf, 0x51, 0x5f, 0x3b, 0xc7, 0xe2, 0x5b, 0xdb, 0xec, 0xb6,
	0x78, 0x1d, 0x6b, 0xc7, 0x68, 0xab, 0x44, 0xbe, 0xcf, 0x4b, 0xce, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x19, 0x98, 0x92, 0x87, 0x4c, 0x03, 0x00, 0x00,
}