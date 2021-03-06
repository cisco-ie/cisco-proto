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
// source: rib_edm_opaque_obj.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_opaques_opaque

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

type RibEdmOpaqueObj_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName       string   `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	OpaqueClientid       uint32   `protobuf:"varint,5,opt,name=opaque_clientid,json=opaqueClientid,proto3" json:"opaque_clientid,omitempty"`
	OpaqueProtoid        uint32   `protobuf:"varint,6,opt,name=opaque_protoid,json=opaqueProtoid,proto3" json:"opaque_protoid,omitempty"`
	OpaqueKeyType        uint32   `protobuf:"varint,7,opt,name=opaque_key_type,json=opaqueKeyType,proto3" json:"opaque_key_type,omitempty"`
	OpaqueKeySize        uint32   `protobuf:"varint,8,opt,name=opaque_key_size,json=opaqueKeySize,proto3" json:"opaque_key_size,omitempty"`
	OpaqueDataSize       uint32   `protobuf:"varint,9,opt,name=opaque_data_size,json=opaqueDataSize,proto3" json:"opaque_data_size,omitempty"`
	OpaqueString         string   `protobuf:"bytes,10,opt,name=opaque_string,json=opaqueString,proto3" json:"opaque_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RibEdmOpaqueObj_KEYS) Reset()         { *m = RibEdmOpaqueObj_KEYS{} }
func (m *RibEdmOpaqueObj_KEYS) String() string { return proto.CompactTextString(m) }
func (*RibEdmOpaqueObj_KEYS) ProtoMessage()    {}
func (*RibEdmOpaqueObj_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_14e185dab9eb515e, []int{0}
}

func (m *RibEdmOpaqueObj_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RibEdmOpaqueObj_KEYS.Unmarshal(m, b)
}
func (m *RibEdmOpaqueObj_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RibEdmOpaqueObj_KEYS.Marshal(b, m, deterministic)
}
func (m *RibEdmOpaqueObj_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RibEdmOpaqueObj_KEYS.Merge(m, src)
}
func (m *RibEdmOpaqueObj_KEYS) XXX_Size() int {
	return xxx_messageInfo_RibEdmOpaqueObj_KEYS.Size(m)
}
func (m *RibEdmOpaqueObj_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RibEdmOpaqueObj_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RibEdmOpaqueObj_KEYS proto.InternalMessageInfo

func (m *RibEdmOpaqueObj_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *RibEdmOpaqueObj_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *RibEdmOpaqueObj_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *RibEdmOpaqueObj_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueClientid() uint32 {
	if m != nil {
		return m.OpaqueClientid
	}
	return 0
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueProtoid() uint32 {
	if m != nil {
		return m.OpaqueProtoid
	}
	return 0
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueKeyType() uint32 {
	if m != nil {
		return m.OpaqueKeyType
	}
	return 0
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueKeySize() uint32 {
	if m != nil {
		return m.OpaqueKeySize
	}
	return 0
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueDataSize() uint32 {
	if m != nil {
		return m.OpaqueDataSize
	}
	return 0
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueString() string {
	if m != nil {
		return m.OpaqueString
	}
	return ""
}

type RibEdmOpaqueObj struct {
	Key                  string   `protobuf:"bytes,50,opt,name=key,proto3" json:"key,omitempty"`
	Data                 string   `protobuf:"bytes,51,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RibEdmOpaqueObj) Reset()         { *m = RibEdmOpaqueObj{} }
func (m *RibEdmOpaqueObj) String() string { return proto.CompactTextString(m) }
func (*RibEdmOpaqueObj) ProtoMessage()    {}
func (*RibEdmOpaqueObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_14e185dab9eb515e, []int{1}
}

func (m *RibEdmOpaqueObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RibEdmOpaqueObj.Unmarshal(m, b)
}
func (m *RibEdmOpaqueObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RibEdmOpaqueObj.Marshal(b, m, deterministic)
}
func (m *RibEdmOpaqueObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RibEdmOpaqueObj.Merge(m, src)
}
func (m *RibEdmOpaqueObj) XXX_Size() int {
	return xxx_messageInfo_RibEdmOpaqueObj.Size(m)
}
func (m *RibEdmOpaqueObj) XXX_DiscardUnknown() {
	xxx_messageInfo_RibEdmOpaqueObj.DiscardUnknown(m)
}

var xxx_messageInfo_RibEdmOpaqueObj proto.InternalMessageInfo

func (m *RibEdmOpaqueObj) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RibEdmOpaqueObj) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*RibEdmOpaqueObj_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.opaques.opaque.rib_edm_opaque_obj_KEYS")
	proto.RegisterType((*RibEdmOpaqueObj)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.opaques.opaque.rib_edm_opaque_obj")
}

func init() { proto.RegisterFile("rib_edm_opaque_obj.proto", fileDescriptor_14e185dab9eb515e) }

var fileDescriptor_14e185dab9eb515e = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4f, 0x2a, 0x31,
	0x10, 0xc7, 0xc3, 0x83, 0xc7, 0x8f, 0xe6, 0xc1, 0x23, 0xbd, 0x50, 0x6e, 0x04, 0xa3, 0x72, 0xea,
	0x41, 0x3c, 0x79, 0x55, 0x4f, 0x24, 0xc6, 0x00, 0x17, 0x4f, 0x93, 0x2e, 0x3b, 0x6b, 0x2a, 0x3f,
	0xba, 0xb6, 0x65, 0xe3, 0x72, 0xf3, 0x3f, 0x37, 0x3b, 0xad, 0x86, 0x40, 0xbc, 0x6c, 0x9b, 0xcf,
	0xf7, 0x33, 0x3b, 0x33, 0xbb, 0x4c, 0x58, 0x9d, 0x00, 0xa6, 0x5b, 0x30, 0xb9, 0x7a, 0xdf, 0x23,
	0x98, 0xe4, 0x4d, 0xe6, 0xd6, 0x78, 0xc3, 0x0f, 0x2b, 0xed, 0x56, 0x06, 0xb4, 0x71, 0xf0, 0x61,
	0x41, 0xe7, 0x50, 0x99, 0x3a, 0x2f, 0x6e, 0xc1, 0xe4, 0x68, 0xa5, 0xd5, 0x89, 0x2c, 0x6c, 0xe6,
	0xaa, 0x87, 0x54, 0x99, 0x93, 0x2a, 0x93, 0xae, 0x3a, 0x9d, 0xca, 0x64, 0xb4, 0xad, 0xd9, 0x7b,
	0x04, 0xaf, 0x92, 0x0d, 0xc2, 0x4e, 0x6d, 0xd1, 0xfd, 0x16, 0xc8, 0xd0, 0xdd, 0xc5, 0x73, 0xfc,
	0x59, 0x67, 0x83, 0xf3, 0xc1, 0x60, 0xf6, 0xf8, 0xb2, 0xe0, 0x43, 0xd6, 0x2e, 0x6c, 0x46, 0x65,
	0xa2, 0x36, 0xaa, 0x4d, 0x3a, 0xf3, 0x56, 0x61, 0xb3, 0x27, 0xb5, 0x45, 0x3e, 0x60, 0x2d, 0x15,
	0x93, 0x3f, 0x94, 0x34, 0x55, 0x08, 0x86, 0xac, 0xed, 0xbe, 0x93, 0x7a, 0xa8, 0x71, 0x31, 0x9a,
	0xb0, 0xfe, 0xe9, 0x34, 0xa2, 0x41, 0x4a, 0x8f, 0xf8, 0xb2, 0xc2, 0x64, 0x5e, 0xb3, 0xff, 0x71,
	0x96, 0xd5, 0x46, 0xe3, 0xce, 0xeb, 0x54, 0xfc, 0x1d, 0xd5, 0x26, 0xdd, 0x79, 0x2f, 0xe0, 0xfb,
	0x48, 0xf9, 0x25, 0x8b, 0x04, 0xe8, 0x4b, 0xea, 0x54, 0x34, 0xc9, 0xeb, 0x06, 0xfa, 0x1c, 0x20,
	0xbf, 0xfa, 0x79, 0xdf, 0x1a, 0x4b, 0xf0, 0x65, 0x8e, 0xa2, 0x75, 0xec, 0xcd, 0xb0, 0x5c, 0x96,
	0x39, 0x9e, 0x78, 0x4e, 0x1f, 0x50, 0xb4, 0x4f, 0xbc, 0x85, 0x3e, 0xd0, 0x26, 0xd1, 0x4b, 0x95,
	0x57, 0x41, 0xec, 0x1c, 0x0f, 0xf8, 0xa0, 0xbc, 0x22, 0xf3, 0x82, 0xc5, 0x52, 0x70, 0xde, 0xea,
	0xdd, 0xab, 0x60, 0xb4, 0xf0, 0xbf, 0x00, 0x17, 0xc4, 0xc6, 0x77, 0x8c, 0x9f, 0xff, 0x02, 0xde,
	0x67, 0xf5, 0x35, 0x96, 0xe2, 0x86, 0x0a, 0xaa, 0x2b, 0xe7, 0xac, 0x51, 0xf5, 0x13, 0x53, 0x42,
	0x74, 0x4f, 0x9a, 0xb4, 0xf8, 0xf4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xfd, 0xcf, 0x59, 0x5e,
	0x02, 0x00, 0x00,
}
