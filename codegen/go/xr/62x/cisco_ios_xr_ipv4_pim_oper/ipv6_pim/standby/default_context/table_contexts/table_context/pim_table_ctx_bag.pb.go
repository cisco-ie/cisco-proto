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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_default_context_table_contexts_table_context

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
	proto.RegisterType((*PimTableCtxBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.table_contexts.table_context.pim_table_ctx_bag_KEYS")
	proto.RegisterType((*PimTableCtxBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.table_contexts.table_context.pim_table_ctx_bag")
}

func init() { proto.RegisterFile("pim_table_ctx_bag.proto", fileDescriptor_89acfafa72caf04a) }

var fileDescriptor_89acfafa72caf04a = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6b, 0x14, 0x41,
	0x10, 0xc5, 0x59, 0xa3, 0xc9, 0x6e, 0x61, 0x60, 0xd3, 0x90, 0xd8, 0x21, 0x08, 0x4b, 0xbc, 0x2c,
	0x28, 0x73, 0x30, 0x31, 0x31, 0xfe, 0x39, 0x88, 0x7a, 0x58, 0x04, 0x0f, 0xe3, 0x25, 0x9e, 0x8a,
	0x9e, 0x99, 0x9a, 0xa5, 0x60, 0x76, 0xba, 0xe9, 0x6a, 0x87, 0xcd, 0x07, 0xf3, 0xfb, 0xc9, 0xd4,
	0xae, 0x98, 0x75, 0x72, 0xeb, 0x57, 0xef, 0xfd, 0x5e, 0x15, 0xc3, 0xc0, 0xb3, 0xc0, 0x2b, 0x4c,
	0xae, 0x68, 0x08, 0xcb, 0xb4, 0xc6, 0xc2, 0x2d, 0xb3, 0x10, 0x7d, 0xf2, 0xe6, 0xb6, 0x64, 0x29,
	0x3d, 0xb2, 0x17, 0x5c, 0x47, 0xe4, 0xd0, 0x5d, 0x62, 0x1f, 0xf5, 0x81, 0x62, 0xc6, 0xa1, 0xbb,
	0xea, 0x55, 0x26, 0xc9, 0xb5, 0x55, 0x71, 0x97, 0x55, 0x54, 0xbb, 0x5f, 0x4d, 0xc2, 0xd2, 0xb7,
	0x89, 0xd6, 0x29, 0xdb, 0x16, 0x6e, 0x94, 0xec, 0xca, 0xf3, 0x5b, 0x38, 0x19, 0x2c, 0xc5, 0x6f,
	0x5f, 0x7f, 0xfe, 0x30, 0xa7, 0x30, 0x16, 0x57, 0x63, 0xeb, 0x56, 0x64, 0x47, 0xb3, 0xd1, 0x7c,
	0x92, 0x1f, 0x88, 0xab, 0xbf, 0xbb, 0x15, 0x99, 0x17, 0x70, 0x98, 0x7c, 0xf0, 0x8d, 0x5f, 0xde,
	0x6d, 0xfc, 0x47, 0xea, 0x3f, 0xfd, 0x3b, 0xec, 0x43, 0xe7, 0xbf, 0xf7, 0xe0, 0x68, 0x50, 0x6d,
	0xa6, 0xb0, 0xe7, 0x6a, 0xb6, 0xaf, 0x67, 0xa3, 0xf9, 0x61, 0xde, 0x3f, 0x8d, 0x81, 0xc7, 0xd2,
	0x8f, 0x2e, 0x74, 0xa4, 0x6f, 0xf3, 0x1c, 0x60, 0x83, 0x69, 0xfb, 0xa5, 0xb6, 0x4f, 0x74, 0xa2,
	0xfb, 0x8f, 0x61, 0xbf, 0x8b, 0x35, 0x72, 0x65, 0xdf, 0x28, 0xf4, 0xa4, 0x8b, 0xf5, 0xa2, 0xea,
	0x2f, 0xde, 0x50, 0x5c, 0xd9, 0x2b, 0x35, 0x0e, 0x54, 0x2f, 0x2a, 0x73, 0x06, 0x13, 0x16, 0x74,
	0x65, 0xe2, 0x8e, 0xec, 0xf5, 0x6c, 0x34, 0x1f, 0xe7, 0x63, 0x96, 0x4f, 0xaa, 0xcd, 0x0d, 0x9c,
	0xb2, 0x20, 0x27, 0xd7, 0x60, 0xa4, 0x25, 0x4b, 0x8a, 0x2e, 0xb1, 0x6f, 0xb1, 0xf2, 0x2d, 0xd9,
	0xb7, 0x1a, 0x3e, 0x61, 0x59, 0x24, 0xd7, 0xe4, 0xf7, 0xec, 0x2f, 0xbe, 0x25, 0x73, 0x0d, 0x96,
	0x05, 0x23, 0x17, 0x0f, 0x90, 0x37, 0x4a, 0x1e, 0xb3, 0xe4, 0x5c, 0x0c, 0xc0, 0x8f, 0x70, 0xb6,
	0x05, 0x4b, 0xdf, 0x76, 0x14, 0x97, 0xd4, 0x96, 0x84, 0x91, 0x4a, 0xe2, 0x8e, 0x2a, 0xfb, 0x4e,
	0x59, 0xab, 0xec, 0xe7, 0x7f, 0x81, 0x7c, 0xeb, 0x9b, 0x57, 0x60, 0x86, 0xb8, 0x7d, 0xaf, 0xd4,
	0xf4, 0x7f, 0xca, 0xbc, 0x84, 0xa3, 0x18, 0xea, 0x9d, 0x13, 0xc5, 0x7e, 0xd0, 0x2f, 0x34, 0x8d,
	0xa1, 0xbe, 0x7f, 0x9c, 0x14, 0xfb, 0xfa, 0xcb, 0x5d, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc7,
	0xb7, 0x48, 0xe7, 0x8d, 0x02, 0x00, 0x00,
}
