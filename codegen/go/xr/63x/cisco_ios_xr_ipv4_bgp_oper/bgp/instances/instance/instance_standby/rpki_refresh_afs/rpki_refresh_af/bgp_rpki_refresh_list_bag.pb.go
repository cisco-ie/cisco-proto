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
// source: bgp_rpki_refresh_list_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_rpki_refresh_afs_rpki_refresh_af

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

type BgpRpkiRefreshListBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpRpkiRefreshListBag_KEYS) Reset()         { *m = BgpRpkiRefreshListBag_KEYS{} }
func (m *BgpRpkiRefreshListBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpRpkiRefreshListBag_KEYS) ProtoMessage()    {}
func (*BgpRpkiRefreshListBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_018ceb9416958d1e, []int{0}
}

func (m *BgpRpkiRefreshListBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRpkiRefreshListBag_KEYS.Unmarshal(m, b)
}
func (m *BgpRpkiRefreshListBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRpkiRefreshListBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpRpkiRefreshListBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRpkiRefreshListBag_KEYS.Merge(m, src)
}
func (m *BgpRpkiRefreshListBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpRpkiRefreshListBag_KEYS.Size(m)
}
func (m *BgpRpkiRefreshListBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRpkiRefreshListBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRpkiRefreshListBag_KEYS proto.InternalMessageInfo

func (m *BgpRpkiRefreshListBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpRpkiRefreshListBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

type BgpEdmRpkiRefreshEntryType_ struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Neighbor             string   `protobuf:"bytes,2,opt,name=neighbor,proto3" json:"neighbor,omitempty"`
	Policy               string   `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
	Paths                uint32   `protobuf:"varint,4,opt,name=paths,proto3" json:"paths,omitempty"`
	Drop                 bool     `protobuf:"varint,5,opt,name=drop,proto3" json:"drop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpEdmRpkiRefreshEntryType_) Reset()         { *m = BgpEdmRpkiRefreshEntryType_{} }
func (m *BgpEdmRpkiRefreshEntryType_) String() string { return proto.CompactTextString(m) }
func (*BgpEdmRpkiRefreshEntryType_) ProtoMessage()    {}
func (*BgpEdmRpkiRefreshEntryType_) Descriptor() ([]byte, []int) {
	return fileDescriptor_018ceb9416958d1e, []int{1}
}

func (m *BgpEdmRpkiRefreshEntryType_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpEdmRpkiRefreshEntryType_.Unmarshal(m, b)
}
func (m *BgpEdmRpkiRefreshEntryType_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpEdmRpkiRefreshEntryType_.Marshal(b, m, deterministic)
}
func (m *BgpEdmRpkiRefreshEntryType_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpEdmRpkiRefreshEntryType_.Merge(m, src)
}
func (m *BgpEdmRpkiRefreshEntryType_) XXX_Size() int {
	return xxx_messageInfo_BgpEdmRpkiRefreshEntryType_.Size(m)
}
func (m *BgpEdmRpkiRefreshEntryType_) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpEdmRpkiRefreshEntryType_.DiscardUnknown(m)
}

var xxx_messageInfo_BgpEdmRpkiRefreshEntryType_ proto.InternalMessageInfo

func (m *BgpEdmRpkiRefreshEntryType_) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpEdmRpkiRefreshEntryType_) GetNeighbor() string {
	if m != nil {
		return m.Neighbor
	}
	return ""
}

func (m *BgpEdmRpkiRefreshEntryType_) GetPolicy() string {
	if m != nil {
		return m.Policy
	}
	return ""
}

func (m *BgpEdmRpkiRefreshEntryType_) GetPaths() uint32 {
	if m != nil {
		return m.Paths
	}
	return 0
}

func (m *BgpEdmRpkiRefreshEntryType_) GetDrop() bool {
	if m != nil {
		return m.Drop
	}
	return false
}

type BgpRpkiRefreshListBag struct {
	RefreshEntry         []*BgpEdmRpkiRefreshEntryType_ `protobuf:"bytes,50,rep,name=refresh_entry,json=refreshEntry,proto3" json:"refresh_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *BgpRpkiRefreshListBag) Reset()         { *m = BgpRpkiRefreshListBag{} }
func (m *BgpRpkiRefreshListBag) String() string { return proto.CompactTextString(m) }
func (*BgpRpkiRefreshListBag) ProtoMessage()    {}
func (*BgpRpkiRefreshListBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_018ceb9416958d1e, []int{2}
}

func (m *BgpRpkiRefreshListBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRpkiRefreshListBag.Unmarshal(m, b)
}
func (m *BgpRpkiRefreshListBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRpkiRefreshListBag.Marshal(b, m, deterministic)
}
func (m *BgpRpkiRefreshListBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRpkiRefreshListBag.Merge(m, src)
}
func (m *BgpRpkiRefreshListBag) XXX_Size() int {
	return xxx_messageInfo_BgpRpkiRefreshListBag.Size(m)
}
func (m *BgpRpkiRefreshListBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRpkiRefreshListBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRpkiRefreshListBag proto.InternalMessageInfo

func (m *BgpRpkiRefreshListBag) GetRefreshEntry() []*BgpEdmRpkiRefreshEntryType_ {
	if m != nil {
		return m.RefreshEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpRpkiRefreshListBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.rpki_refresh_afs.rpki_refresh_af.bgp_rpki_refresh_list_bag_KEYS")
	proto.RegisterType((*BgpEdmRpkiRefreshEntryType_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.rpki_refresh_afs.rpki_refresh_af.bgp_edm_rpki_refresh_entry_type_")
	proto.RegisterType((*BgpRpkiRefreshListBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.rpki_refresh_afs.rpki_refresh_af.bgp_rpki_refresh_list_bag")
}

func init() { proto.RegisterFile("bgp_rpki_refresh_list_bag.proto", fileDescriptor_018ceb9416958d1e) }

var fileDescriptor_018ceb9416958d1e = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x95, 0xbf, 0xfe, 0x7c, 0xe5, 0xd2, 0x2e, 0x16, 0x82, 0xc0, 0x00, 0x51, 0x58, 0x32, 0x79,
	0x28, 0xbc, 0x42, 0x27, 0x24, 0x86, 0x30, 0xb1, 0x70, 0x65, 0x27, 0x4e, 0x62, 0xd1, 0xc4, 0x96,
	0x6d, 0x21, 0xf2, 0x00, 0x3c, 0x04, 0xcf, 0xc0, 0x93, 0xf0, 0x56, 0x28, 0x49, 0x89, 0x1a, 0xa4,
	0x8a, 0x8d, 0x29, 0xe7, 0x47, 0xe7, 0xdc, 0x1b, 0xdb, 0x70, 0x25, 0x0a, 0x83, 0xd6, 0x3c, 0x2b,
	0xb4, 0x32, 0xb7, 0xd2, 0x95, 0xb8, 0x55, 0xce, 0xa3, 0xe0, 0x05, 0x33, 0x56, 0x7b, 0x4d, 0xd3,
	0x54, 0xb9, 0x54, 0xa3, 0xd2, 0x0e, 0x5f, 0x2d, 0x2a, 0xf3, 0x72, 0x8b, 0x6d, 0x44, 0x1b, 0x69,
	0x99, 0x28, 0x0c, 0x53, 0xb5, 0xf3, 0xbc, 0x4e, 0xa5, 0x1b, 0xd0, 0x00, 0xb0, 0xfd, 0x64, 0xa2,
	0x61, 0xa3, 0x7e, 0x9e, 0xbb, 0x9f, 0x42, 0xf4, 0x04, 0x97, 0x07, 0xf7, 0xc0, 0xbb, 0xcd, 0xe3,
	0x03, 0xbd, 0x86, 0xd5, 0x50, 0x5b, 0xf3, 0x4a, 0x06, 0x24, 0x24, 0xf1, 0x51, 0xb2, 0xfc, 0x16,
	0xef, 0x79, 0x25, 0xe9, 0x19, 0xfc, 0xe7, 0x79, 0x6f, 0xff, 0xeb, 0xec, 0x39, 0xcf, 0x5b, 0x23,
	0x7a, 0x27, 0x10, 0xb6, 0x03, 0x64, 0x56, 0x8d, 0x87, 0xc8, 0xda, 0xdb, 0x06, 0x7d, 0x63, 0x24,
	0xee, 0xa7, 0xc9, 0x7e, 0x9a, 0x5e, 0xc0, 0xa2, 0x96, 0xaa, 0x28, 0x85, 0xb6, 0xbb, 0xde, 0x81,
	0xd3, 0x53, 0x98, 0x1b, 0xbd, 0x55, 0x69, 0x13, 0x4c, 0xfa, 0x4c, 0xcf, 0xe8, 0x09, 0xcc, 0x0c,
	0xf7, 0xa5, 0x0b, 0xa6, 0x21, 0x89, 0x57, 0x49, 0x4f, 0x28, 0x85, 0x69, 0x66, 0xb5, 0x09, 0x66,
	0x21, 0x89, 0x17, 0x49, 0x87, 0xa3, 0x4f, 0x02, 0xe7, 0x07, 0x7f, 0x9e, 0x7e, 0x10, 0x58, 0x8d,
	0x96, 0x0d, 0xd6, 0xe1, 0x24, 0x3e, 0x5e, 0xbf, 0x11, 0xf6, 0x07, 0x17, 0xc3, 0x7e, 0x3b, 0xb4,
	0x64, 0xb9, 0x13, 0x37, 0xad, 0x26, 0xe6, 0xdd, 0x9b, 0xb9, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x04, 0xc2, 0x71, 0x21, 0x56, 0x02, 0x00, 0x00,
}
