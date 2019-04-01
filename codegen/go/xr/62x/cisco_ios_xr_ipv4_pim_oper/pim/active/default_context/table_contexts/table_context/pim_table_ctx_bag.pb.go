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

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_table_contexts_table_context

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
	proto.RegisterType((*PimTableCtxBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.table_contexts.table_context.pim_table_ctx_bag_KEYS")
	proto.RegisterType((*PimTableCtxBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.table_contexts.table_context.pim_table_ctx_bag")
}

func init() { proto.RegisterFile("pim_table_ctx_bag.proto", fileDescriptor_89acfafa72caf04a) }

var fileDescriptor_89acfafa72caf04a = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6b, 0x14, 0x41,
	0x10, 0xc5, 0x59, 0xa3, 0xc9, 0x6e, 0x61, 0x60, 0xd3, 0x90, 0xd8, 0x21, 0x08, 0x4b, 0xbc, 0x2c,
	0x28, 0x73, 0x30, 0xd1, 0x18, 0xff, 0x1c, 0x44, 0x3d, 0x2c, 0x82, 0x87, 0xf1, 0xa2, 0xa7, 0xa2,
	0xa7, 0xa7, 0x7a, 0x29, 0x98, 0x9d, 0x6e, 0xba, 0xdb, 0x61, 0xfd, 0x60, 0x7e, 0x3f, 0x99, 0x9a,
	0x15, 0xb3, 0x8e, 0xb7, 0xae, 0xf7, 0xde, 0xef, 0x75, 0xd1, 0x34, 0x3c, 0x0a, 0xbc, 0xc1, 0x6c,
	0xaa, 0x86, 0xd0, 0xe6, 0x2d, 0x56, 0x66, 0x5d, 0x84, 0xe8, 0xb3, 0x57, 0xa5, 0xe5, 0x64, 0x3d,
	0xb2, 0x4f, 0xb8, 0x8d, 0xc8, 0xa1, 0xbb, 0xc6, 0x3e, 0xea, 0x03, 0xc5, 0x22, 0xf0, 0xa6, 0x30,
	0x36, 0x73, 0x47, 0x45, 0x4d, 0xce, 0xfc, 0x68, 0x32, 0x5a, 0xdf, 0x66, 0xda, 0xe6, 0x62, 0x57,
	0x35, 0x4c, 0x69, 0x7f, 0xbc, 0xfc, 0x06, 0x67, 0xa3, 0xeb, 0xf0, 0xf3, 0xa7, 0xef, 0x5f, 0xd5,
	0x39, 0x4c, 0x93, 0x71, 0xd8, 0x9a, 0x0d, 0xe9, 0xc9, 0x62, 0xb2, 0x9c, 0x95, 0x47, 0xc9, 0xb8,
	0x2f, 0x66, 0x43, 0xea, 0x09, 0x1c, 0x67, 0x1f, 0x7c, 0xe3, 0xd7, 0x3f, 0x07, 0xff, 0x9e, 0xf8,
	0x0f, 0xff, 0x88, 0x7d, 0xe8, 0xf2, 0xd7, 0x01, 0x9c, 0x8c, 0xaa, 0xd5, 0x1c, 0x0e, 0x8c, 0x63,
	0xfd, 0x7c, 0x31, 0x59, 0x1e, 0x97, 0xfd, 0x51, 0x29, 0xb8, 0x9f, 0x7a, 0xe9, 0x4a, 0x24, 0x39,
	0xab, 0xc7, 0x00, 0x03, 0x26, 0xed, 0xd7, 0xd2, 0x3e, 0x13, 0x45, 0xee, 0x3f, 0x85, 0xc3, 0x2e,
	0x3a, 0xe4, 0x5a, 0xbf, 0x10, 0xe8, 0x41, 0x17, 0xdd, 0xaa, 0xee, 0x37, 0x1e, 0x28, 0xae, 0xf5,
	0x4b, 0x31, 0x8e, 0x64, 0x5e, 0xd5, 0xea, 0x02, 0x66, 0x9c, 0x70, 0x78, 0x20, 0x7d, 0xb3, 0x98,
	0x2c, 0xa7, 0xe5, 0x94, 0xd3, 0x7b, 0x99, 0xd5, 0x2d, 0x9c, 0x73, 0x42, 0xce, 0xa6, 0xc1, 0x48,
	0x6b, 0x4e, 0x39, 0x9a, 0xcc, 0xbe, 0xc5, 0xda, 0xb7, 0xa4, 0x5f, 0x49, 0xf8, 0x8c, 0xd3, 0x2a,
	0x9b, 0xa6, 0xbc, 0x63, 0x7f, 0xf4, 0x2d, 0xa9, 0x1b, 0xd0, 0x9c, 0x30, 0x72, 0xf5, 0x1f, 0xf2,
	0x56, 0xc8, 0x53, 0x4e, 0x25, 0x57, 0x23, 0xf0, 0x1d, 0x5c, 0xec, 0x40, 0xeb, 0xdb, 0x8e, 0xe2,
	0x9a, 0x5a, 0x4b, 0x18, 0xc9, 0x12, 0x77, 0x54, 0xeb, 0xd7, 0xc2, 0x6a, 0x61, 0x3f, 0xfc, 0x0d,
	0x94, 0x3b, 0x5f, 0x3d, 0x03, 0x35, 0xc6, 0xf5, 0x1b, 0xa1, 0xe6, 0xff, 0x52, 0xea, 0x29, 0x9c,
	0xc4, 0xe0, 0xf6, 0x56, 0x4c, 0xfa, 0xad, 0xbc, 0xd0, 0x3c, 0x06, 0x77, 0x77, 0xb9, 0x54, 0x1d,
	0xca, 0x67, 0xbb, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x8a, 0x01, 0x93, 0x87, 0x02, 0x00,
	0x00,
}
