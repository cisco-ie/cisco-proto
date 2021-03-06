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
// source: bgp_label_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_label_entries_label_entry

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

type BgpLabelBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Label                uint32   `protobuf:"varint,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpLabelBag_KEYS) Reset()         { *m = BgpLabelBag_KEYS{} }
func (m *BgpLabelBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpLabelBag_KEYS) ProtoMessage()    {}
func (*BgpLabelBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dafb6ce8340483a, []int{0}
}

func (m *BgpLabelBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpLabelBag_KEYS.Unmarshal(m, b)
}
func (m *BgpLabelBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpLabelBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpLabelBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpLabelBag_KEYS.Merge(m, src)
}
func (m *BgpLabelBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpLabelBag_KEYS.Size(m)
}
func (m *BgpLabelBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpLabelBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpLabelBag_KEYS proto.InternalMessageInfo

func (m *BgpLabelBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpLabelBag_KEYS) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

type BgpEdmLabelEntry_ struct {
	Label                uint32   `protobuf:"varint,1,opt,name=label,proto3" json:"label,omitempty"`
	Rds                  string   `protobuf:"bytes,2,opt,name=rds,proto3" json:"rds,omitempty"`
	Vrf                  string   `protobuf:"bytes,3,opt,name=vrf,proto3" json:"vrf,omitempty"`
	Ip                   string   `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Ip6                  string   `protobuf:"bytes,5,opt,name=ip6,proto3" json:"ip6,omitempty"`
	RpcSetId             uint32   `protobuf:"varint,6,opt,name=rpc_set_id,json=rpcSetId,proto3" json:"rpc_set_id,omitempty"`
	Masklen              uint32   `protobuf:"varint,7,opt,name=masklen,proto3" json:"masklen,omitempty"`
	TsSec                uint32   `protobuf:"varint,8,opt,name=ts_sec,json=tsSec,proto3" json:"ts_sec,omitempty"`
	TsSsec               uint32   `protobuf:"varint,9,opt,name=ts_ssec,json=tsSsec,proto3" json:"ts_ssec,omitempty"`
	Info                 uint32   `protobuf:"varint,10,opt,name=info,proto3" json:"info,omitempty"`
	Refcount             uint32   `protobuf:"varint,11,opt,name=refcount,proto3" json:"refcount,omitempty"`
	Inactive             bool     `protobuf:"varint,12,opt,name=inactive,proto3" json:"inactive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpEdmLabelEntry_) Reset()         { *m = BgpEdmLabelEntry_{} }
func (m *BgpEdmLabelEntry_) String() string { return proto.CompactTextString(m) }
func (*BgpEdmLabelEntry_) ProtoMessage()    {}
func (*BgpEdmLabelEntry_) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dafb6ce8340483a, []int{1}
}

func (m *BgpEdmLabelEntry_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpEdmLabelEntry_.Unmarshal(m, b)
}
func (m *BgpEdmLabelEntry_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpEdmLabelEntry_.Marshal(b, m, deterministic)
}
func (m *BgpEdmLabelEntry_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpEdmLabelEntry_.Merge(m, src)
}
func (m *BgpEdmLabelEntry_) XXX_Size() int {
	return xxx_messageInfo_BgpEdmLabelEntry_.Size(m)
}
func (m *BgpEdmLabelEntry_) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpEdmLabelEntry_.DiscardUnknown(m)
}

var xxx_messageInfo_BgpEdmLabelEntry_ proto.InternalMessageInfo

func (m *BgpEdmLabelEntry_) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *BgpEdmLabelEntry_) GetRds() string {
	if m != nil {
		return m.Rds
	}
	return ""
}

func (m *BgpEdmLabelEntry_) GetVrf() string {
	if m != nil {
		return m.Vrf
	}
	return ""
}

func (m *BgpEdmLabelEntry_) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *BgpEdmLabelEntry_) GetIp6() string {
	if m != nil {
		return m.Ip6
	}
	return ""
}

func (m *BgpEdmLabelEntry_) GetRpcSetId() uint32 {
	if m != nil {
		return m.RpcSetId
	}
	return 0
}

func (m *BgpEdmLabelEntry_) GetMasklen() uint32 {
	if m != nil {
		return m.Masklen
	}
	return 0
}

func (m *BgpEdmLabelEntry_) GetTsSec() uint32 {
	if m != nil {
		return m.TsSec
	}
	return 0
}

func (m *BgpEdmLabelEntry_) GetTsSsec() uint32 {
	if m != nil {
		return m.TsSsec
	}
	return 0
}

func (m *BgpEdmLabelEntry_) GetInfo() uint32 {
	if m != nil {
		return m.Info
	}
	return 0
}

func (m *BgpEdmLabelEntry_) GetRefcount() uint32 {
	if m != nil {
		return m.Refcount
	}
	return 0
}

func (m *BgpEdmLabelEntry_) GetInactive() bool {
	if m != nil {
		return m.Inactive
	}
	return false
}

type BgpLabelBag struct {
	Entry                *BgpEdmLabelEntry_ `protobuf:"bytes,50,opt,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BgpLabelBag) Reset()         { *m = BgpLabelBag{} }
func (m *BgpLabelBag) String() string { return proto.CompactTextString(m) }
func (*BgpLabelBag) ProtoMessage()    {}
func (*BgpLabelBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dafb6ce8340483a, []int{2}
}

func (m *BgpLabelBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpLabelBag.Unmarshal(m, b)
}
func (m *BgpLabelBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpLabelBag.Marshal(b, m, deterministic)
}
func (m *BgpLabelBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpLabelBag.Merge(m, src)
}
func (m *BgpLabelBag) XXX_Size() int {
	return xxx_messageInfo_BgpLabelBag.Size(m)
}
func (m *BgpLabelBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpLabelBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpLabelBag proto.InternalMessageInfo

func (m *BgpLabelBag) GetEntry() *BgpEdmLabelEntry_ {
	if m != nil {
		return m.Entry
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpLabelBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.label_entries.label_entry.bgp_label_bag_KEYS")
	proto.RegisterType((*BgpEdmLabelEntry_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.label_entries.label_entry.bgp_edm_label_entry_")
	proto.RegisterType((*BgpLabelBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.label_entries.label_entry.bgp_label_bag")
}

func init() { proto.RegisterFile("bgp_label_bag.proto", fileDescriptor_1dafb6ce8340483a) }

var fileDescriptor_1dafb6ce8340483a = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0xce, 0xd3, 0x30,
	0x14, 0x85, 0xe5, 0xb4, 0x4d, 0xd3, 0xdb, 0x16, 0x21, 0x53, 0x84, 0x85, 0x18, 0xa2, 0xb2, 0x64,
	0xca, 0x50, 0x50, 0xdf, 0x80, 0x01, 0x21, 0x81, 0x94, 0x4c, 0x88, 0xc1, 0x4a, 0x9c, 0x9b, 0xca,
	0xa2, 0x71, 0x8c, 0x6d, 0x22, 0x78, 0x02, 0x9e, 0x81, 0x17, 0xe0, 0x39, 0x91, 0x9d, 0x36, 0x7f,
	0x2b, 0xfd, 0xeb, 0xbf, 0xdd, 0xf3, 0xf9, 0xe4, 0x9c, 0xc8, 0xd7, 0xf0, 0xa2, 0x3e, 0x69, 0x7e,
	0xae, 0x6a, 0x3c, 0xf3, 0xba, 0x3a, 0xe5, 0xda, 0xf4, 0xae, 0xa7, 0xdf, 0x84, 0xb4, 0xa2, 0xe7,
	0xb2, 0xb7, 0xfc, 0x97, 0xe1, 0x52, 0x0f, 0xef, 0xb9, 0xb7, 0xf5, 0x1a, 0x4d, 0x5e, 0x9f, 0x74,
	0x2e, 0x95, 0x75, 0x95, 0x12, 0x68, 0xa7, 0x69, 0x1a, 0x78, 0x25, 0x9c, 0x1c, 0x30, 0x1f, 0x23,
	0x51, 0x39, 0x23, 0xd1, 0xde, 0xa8, 0xdf, 0xfb, 0x2f, 0x40, 0xef, 0x3a, 0xf9, 0xa7, 0x0f, 0x5f,
	0x4b, 0xfa, 0x16, 0xb6, 0x53, 0x84, 0xaa, 0x3a, 0x64, 0x24, 0x25, 0xd9, 0xaa, 0xd8, 0x5c, 0xe1,
	0xe7, 0xaa, 0x43, 0xba, 0x83, 0x45, 0xf8, 0x8c, 0x45, 0x29, 0xc9, 0xb6, 0xc5, 0x28, 0xf6, 0xff,
	0x22, 0xd8, 0xf9, 0x44, 0x6c, 0x3a, 0x7e, 0x53, 0xc4, 0x1f, 0xec, 0xe4, 0xc6, 0x4e, 0x9f, 0xc3,
	0xcc, 0x34, 0x36, 0x44, 0xac, 0x0a, 0x3f, 0x7a, 0x32, 0x98, 0x96, 0xcd, 0x46, 0x32, 0x98, 0x96,
	0x3e, 0x83, 0x48, 0x6a, 0x36, 0x0f, 0x20, 0x92, 0xda, 0x3b, 0xa4, 0x3e, 0xb2, 0xc5, 0xe8, 0x90,
	0xfa, 0x48, 0xdf, 0x00, 0x18, 0x2d, 0xb8, 0x45, 0xc7, 0x65, 0xc3, 0xe2, 0x50, 0x90, 0x18, 0x2d,
	0x4a, 0x74, 0x1f, 0x1b, 0xca, 0x60, 0xd9, 0x55, 0xf6, 0xfb, 0x19, 0x15, 0x5b, 0x86, 0xa3, 0xab,
	0xa4, 0x2f, 0x21, 0x76, 0x96, 0x5b, 0x14, 0x2c, 0x19, 0x7f, 0xca, 0xd9, 0x12, 0x05, 0x7d, 0x05,
	0x4b, 0x8f, 0x3d, 0x5f, 0x05, 0x1e, 0x3b, 0x5b, 0x5a, 0x14, 0x94, 0xc2, 0x5c, 0xaa, 0xb6, 0x67,
	0x10, 0x68, 0x98, 0xe9, 0x6b, 0x48, 0x0c, 0xb6, 0xa2, 0xff, 0xa9, 0x1c, 0x5b, 0x5f, 0x9a, 0x2f,
	0xda, 0x9f, 0x49, 0x35, 0xee, 0x80, 0x6d, 0x52, 0x92, 0x25, 0xc5, 0xa4, 0xf7, 0x7f, 0x09, 0x6c,
	0xef, 0xae, 0x9e, 0xfe, 0x21, 0xb0, 0x08, 0x97, 0xc5, 0x0e, 0x29, 0xc9, 0xd6, 0x87, 0x1f, 0xf9,
	0x13, 0x6e, 0x3e, 0x7f, 0x6c, 0x49, 0xc5, 0xd8, 0x5f, 0xc7, 0xe1, 0xe5, 0xbd, 0xfb, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0x57, 0x8a, 0x42, 0x06, 0x90, 0x02, 0x00, 0x00,
}
