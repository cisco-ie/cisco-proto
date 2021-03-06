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
// source: pim_table_ctx_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_default_context_table_contexts_table_context

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

type PimTableCtxBag_KEYS struct {
	SafName              string   `protobuf:"bytes,1,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	TopologyName         string   `protobuf:"bytes,2,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimTableCtxBag_KEYS) Reset()         { *m = PimTableCtxBag_KEYS{} }
func (m *PimTableCtxBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimTableCtxBag_KEYS) ProtoMessage()    {}
func (*PimTableCtxBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_89acfafa72caf04a, []int{0}
}

func (m *PimTableCtxBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimTableCtxBag_KEYS.Unmarshal(m, b)
}
func (m *PimTableCtxBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimTableCtxBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimTableCtxBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimTableCtxBag_KEYS.Merge(m, src)
}
func (m *PimTableCtxBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimTableCtxBag_KEYS.Size(m)
}
func (m *PimTableCtxBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimTableCtxBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimTableCtxBag_KEYS proto.InternalMessageInfo

func (m *PimTableCtxBag_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *PimTableCtxBag_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

type PimTableCtxBag struct {
	Afi                      uint32   `protobuf:"varint,50,opt,name=afi,proto3" json:"afi,omitempty"`
	Safi                     uint32   `protobuf:"varint,51,opt,name=safi,proto3" json:"safi,omitempty"`
	TableName                string   `protobuf:"bytes,52,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	VrfId                    uint32   `protobuf:"varint,53,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	TableId                  uint32   `protobuf:"varint,54,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	IsActive                 bool     `protobuf:"varint,55,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	IsItalRegistrationDone   bool     `protobuf:"varint,56,opt,name=is_ital_registration_done,json=isItalRegistrationDone,proto3" json:"is_ital_registration_done,omitempty"`
	IsRibRegistrationDone    bool     `protobuf:"varint,57,opt,name=is_rib_registration_done,json=isRibRegistrationDone,proto3" json:"is_rib_registration_done,omitempty"`
	IsRibConvergenceReceived bool     `protobuf:"varint,58,opt,name=is_rib_convergence_received,json=isRibConvergenceReceived,proto3" json:"is_rib_convergence_received,omitempty"`
	IsRibConvergence         bool     `protobuf:"varint,59,opt,name=is_rib_convergence,json=isRibConvergence,proto3" json:"is_rib_convergence,omitempty"`
	RpfRegistrations         uint32   `protobuf:"varint,60,opt,name=rpf_registrations,json=rpfRegistrations,proto3" json:"rpf_registrations,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *PimTableCtxBag) Reset()         { *m = PimTableCtxBag{} }
func (m *PimTableCtxBag) String() string { return proto.CompactTextString(m) }
func (*PimTableCtxBag) ProtoMessage()    {}
func (*PimTableCtxBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_89acfafa72caf04a, []int{1}
}

func (m *PimTableCtxBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimTableCtxBag.Unmarshal(m, b)
}
func (m *PimTableCtxBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimTableCtxBag.Marshal(b, m, deterministic)
}
func (m *PimTableCtxBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimTableCtxBag.Merge(m, src)
}
func (m *PimTableCtxBag) XXX_Size() int {
	return xxx_messageInfo_PimTableCtxBag.Size(m)
}
func (m *PimTableCtxBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimTableCtxBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimTableCtxBag proto.InternalMessageInfo

func (m *PimTableCtxBag) GetAfi() uint32 {
	if m != nil {
		return m.Afi
	}
	return 0
}

func (m *PimTableCtxBag) GetSafi() uint32 {
	if m != nil {
		return m.Safi
	}
	return 0
}

func (m *PimTableCtxBag) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *PimTableCtxBag) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *PimTableCtxBag) GetTableId() uint32 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *PimTableCtxBag) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *PimTableCtxBag) GetIsItalRegistrationDone() bool {
	if m != nil {
		return m.IsItalRegistrationDone
	}
	return false
}

func (m *PimTableCtxBag) GetIsRibRegistrationDone() bool {
	if m != nil {
		return m.IsRibRegistrationDone
	}
	return false
}

func (m *PimTableCtxBag) GetIsRibConvergenceReceived() bool {
	if m != nil {
		return m.IsRibConvergenceReceived
	}
	return false
}

func (m *PimTableCtxBag) GetIsRibConvergence() bool {
	if m != nil {
		return m.IsRibConvergence
	}
	return false
}

func (m *PimTableCtxBag) GetRpfRegistrations() uint32 {
	if m != nil {
		return m.RpfRegistrations
	}
	return 0
}

func init() {
	proto.RegisterType((*PimTableCtxBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.table_contexts.table_context.pim_table_ctx_bag_KEYS")
	proto.RegisterType((*PimTableCtxBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.table_contexts.table_context.pim_table_ctx_bag")
}

func init() { proto.RegisterFile("pim_table_ctx_bag.proto", fileDescriptor_89acfafa72caf04a) }

var fileDescriptor_89acfafa72caf04a = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6b, 0x14, 0x41,
	0x10, 0xc5, 0x59, 0xa3, 0xc9, 0x6e, 0x61, 0x60, 0xd3, 0x90, 0xd8, 0x21, 0x08, 0x4b, 0xbc, 0x2c,
	0x28, 0x73, 0x30, 0x31, 0x31, 0xfe, 0x39, 0x88, 0x7a, 0x58, 0x04, 0x0f, 0xe3, 0x41, 0x3d, 0x15,
	0x3d, 0x3d, 0xd5, 0x4b, 0xc1, 0xec, 0x74, 0xd3, 0xdd, 0x0e, 0xeb, 0x07, 0xf3, 0xfb, 0xc9, 0xd4,
	0xac, 0x98, 0x38, 0xde, 0xfa, 0xd5, 0x7b, 0xbf, 0xd7, 0x45, 0xd3, 0xf0, 0x28, 0xf0, 0x06, 0xb3,
	0xa9, 0x1a, 0x42, 0x9b, 0xb7, 0x58, 0x99, 0x75, 0x11, 0xa2, 0xcf, 0x5e, 0x7d, 0xb5, 0x9c, 0xac,
	0x47, 0xf6, 0x09, 0xb7, 0x11, 0x39, 0x74, 0x97, 0xd8, 0x47, 0x7d, 0xa0, 0x58, 0x70, 0xe8, 0xae,
	0x7a, 0x55, 0x18, 0x9b, 0xb9, 0xa3, 0xa2, 0x26, 0x67, 0x7e, 0x34, 0x19, 0xad, 0x6f, 0x33, 0x6d,
	0x73, 0xb1, 0xeb, 0x1b, 0x54, 0xba, 0x2b, 0xcf, 0xbf, 0xc1, 0xc9, 0xe8, 0x4e, 0xfc, 0xf4, 0xf1,
	0xfb, 0x17, 0x75, 0x0a, 0xd3, 0x64, 0x1c, 0xb6, 0x66, 0x43, 0x7a, 0xb2, 0x98, 0x2c, 0x67, 0xe5,
	0x41, 0x32, 0xee, 0xb3, 0xd9, 0x90, 0x7a, 0x02, 0x87, 0xd9, 0x07, 0xdf, 0xf8, 0xf5, 0xcf, 0xc1,
	0xbf, 0x27, 0xfe, 0xc3, 0x3f, 0xc3, 0x3e, 0x74, 0xfe, 0x6b, 0x0f, 0x8e, 0x46, 0xd5, 0x6a, 0x0e,
	0x7b, 0xc6, 0xb1, 0x7e, 0xbe, 0x98, 0x2c, 0x0f, 0xcb, 0xfe, 0xa8, 0x14, 0xdc, 0x4f, 0xfd, 0xe8,
	0x42, 0x46, 0x72, 0x56, 0x8f, 0x01, 0x06, 0x4c, 0xda, 0x2f, 0xa5, 0x7d, 0x26, 0x13, 0xb9, 0xff,
	0x18, 0xf6, 0xbb, 0xe8, 0x90, 0x6b, 0xfd, 0x42, 0xa0, 0x07, 0x5d, 0x74, 0xab, 0xba, 0xdf, 0x78,
	0xa0, 0xb8, 0xd6, 0x57, 0x62, 0x1c, 0x88, 0x5e, 0xd5, 0xea, 0x0c, 0x66, 0x9c, 0x70, 0x78, 0x20,
	0x7d, 0xbd, 0x98, 0x2c, 0xa7, 0xe5, 0x94, 0xd3, 0x3b, 0xd1, 0xea, 0x06, 0x4e, 0x39, 0x21, 0x67,
	0xd3, 0x60, 0xa4, 0x35, 0xa7, 0x1c, 0x4d, 0x66, 0xdf, 0x62, 0xed, 0x5b, 0xd2, 0x2f, 0x25, 0x7c,
	0xc2, 0x69, 0x95, 0x4d, 0x53, 0xde, 0xb2, 0x3f, 0xf8, 0x96, 0xd4, 0x35, 0x68, 0x4e, 0x18, 0xb9,
	0xfa, 0x0f, 0x79, 0x23, 0xe4, 0x31, 0xa7, 0x92, 0xab, 0x11, 0xf8, 0x16, 0xce, 0x76, 0xa0, 0xf5,
	0x6d, 0x47, 0x71, 0x4d, 0xad, 0x25, 0x8c, 0x64, 0x89, 0x3b, 0xaa, 0xf5, 0x2b, 0x61, 0xb5, 0xb0,
	0xef, 0xff, 0x06, 0xca, 0x9d, 0xaf, 0x9e, 0x81, 0x1a, 0xe3, 0xfa, 0xb5, 0x50, 0xf3, 0x7f, 0x29,
	0xf5, 0x14, 0x8e, 0x62, 0x70, 0x77, 0x56, 0x4c, 0xfa, 0x8d, 0xbc, 0xd0, 0x3c, 0x06, 0x77, 0x7b,
	0xb9, 0x54, 0xed, 0xcb, 0x8f, 0xbb, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xb8, 0x13, 0x35,
	0x8c, 0x02, 0x00, 0x00,
}
