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
// source: bmd_bfd_counter_bag_multiple.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_bfd_counters_bfd_counters_bundles_bfd_counters_bundle_bfd_counters_bundle_descendant

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

type BmdBfdCounterBagMultiple_KEYS struct {
	BundleInterface      string   `protobuf:"bytes,1,opt,name=bundle_interface,json=bundleInterface,proto3" json:"bundle_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdBfdCounterBagMultiple_KEYS) Reset()         { *m = BmdBfdCounterBagMultiple_KEYS{} }
func (m *BmdBfdCounterBagMultiple_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmdBfdCounterBagMultiple_KEYS) ProtoMessage()    {}
func (*BmdBfdCounterBagMultiple_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_555c9eddef818db9, []int{0}
}

func (m *BmdBfdCounterBagMultiple_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBfdCounterBagMultiple_KEYS.Unmarshal(m, b)
}
func (m *BmdBfdCounterBagMultiple_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBfdCounterBagMultiple_KEYS.Marshal(b, m, deterministic)
}
func (m *BmdBfdCounterBagMultiple_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBfdCounterBagMultiple_KEYS.Merge(m, src)
}
func (m *BmdBfdCounterBagMultiple_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmdBfdCounterBagMultiple_KEYS.Size(m)
}
func (m *BmdBfdCounterBagMultiple_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBfdCounterBagMultiple_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBfdCounterBagMultiple_KEYS proto.InternalMessageInfo

func (m *BmdBfdCounterBagMultiple_KEYS) GetBundleInterface() string {
	if m != nil {
		return m.BundleInterface
	}
	return ""
}

type BmNameBag struct {
	ItemName             string   `protobuf:"bytes,1,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmNameBag) Reset()         { *m = BmNameBag{} }
func (m *BmNameBag) String() string { return proto.CompactTextString(m) }
func (*BmNameBag) ProtoMessage()    {}
func (*BmNameBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_555c9eddef818db9, []int{1}
}

func (m *BmNameBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmNameBag.Unmarshal(m, b)
}
func (m *BmNameBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmNameBag.Marshal(b, m, deterministic)
}
func (m *BmNameBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmNameBag.Merge(m, src)
}
func (m *BmNameBag) XXX_Size() int {
	return xxx_messageInfo_BmNameBag.Size(m)
}
func (m *BmNameBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BmNameBag.DiscardUnknown(m)
}

var xxx_messageInfo_BmNameBag proto.InternalMessageInfo

func (m *BmNameBag) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

type BmdBfdCounterBag struct {
	MemberName                   string   `protobuf:"bytes,1,opt,name=member_name,json=memberName,proto3" json:"member_name,omitempty"`
	LastTimeCleared              uint64   `protobuf:"varint,2,opt,name=last_time_cleared,json=lastTimeCleared,proto3" json:"last_time_cleared,omitempty"`
	Starting                     uint32   `protobuf:"varint,3,opt,name=starting,proto3" json:"starting,omitempty"`
	Up                           uint32   `protobuf:"varint,4,opt,name=up,proto3" json:"up,omitempty"`
	Down                         uint32   `protobuf:"varint,5,opt,name=down,proto3" json:"down,omitempty"`
	NeighborUnconfigured         uint32   `protobuf:"varint,6,opt,name=neighbor_unconfigured,json=neighborUnconfigured,proto3" json:"neighbor_unconfigured,omitempty"`
	StartTimeouts                uint32   `protobuf:"varint,7,opt,name=start_timeouts,json=startTimeouts,proto3" json:"start_timeouts,omitempty"`
	NeighborUnconfiguredTimeouts uint32   `protobuf:"varint,8,opt,name=neighbor_unconfigured_timeouts,json=neighborUnconfiguredTimeouts,proto3" json:"neighbor_unconfigured_timeouts,omitempty"`
	TimeSinceCleared             uint64   `protobuf:"varint,9,opt,name=time_since_cleared,json=timeSinceCleared,proto3" json:"time_since_cleared,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *BmdBfdCounterBag) Reset()         { *m = BmdBfdCounterBag{} }
func (m *BmdBfdCounterBag) String() string { return proto.CompactTextString(m) }
func (*BmdBfdCounterBag) ProtoMessage()    {}
func (*BmdBfdCounterBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_555c9eddef818db9, []int{2}
}

func (m *BmdBfdCounterBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBfdCounterBag.Unmarshal(m, b)
}
func (m *BmdBfdCounterBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBfdCounterBag.Marshal(b, m, deterministic)
}
func (m *BmdBfdCounterBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBfdCounterBag.Merge(m, src)
}
func (m *BmdBfdCounterBag) XXX_Size() int {
	return xxx_messageInfo_BmdBfdCounterBag.Size(m)
}
func (m *BmdBfdCounterBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBfdCounterBag.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBfdCounterBag proto.InternalMessageInfo

func (m *BmdBfdCounterBag) GetMemberName() string {
	if m != nil {
		return m.MemberName
	}
	return ""
}

func (m *BmdBfdCounterBag) GetLastTimeCleared() uint64 {
	if m != nil {
		return m.LastTimeCleared
	}
	return 0
}

func (m *BmdBfdCounterBag) GetStarting() uint32 {
	if m != nil {
		return m.Starting
	}
	return 0
}

func (m *BmdBfdCounterBag) GetUp() uint32 {
	if m != nil {
		return m.Up
	}
	return 0
}

func (m *BmdBfdCounterBag) GetDown() uint32 {
	if m != nil {
		return m.Down
	}
	return 0
}

func (m *BmdBfdCounterBag) GetNeighborUnconfigured() uint32 {
	if m != nil {
		return m.NeighborUnconfigured
	}
	return 0
}

func (m *BmdBfdCounterBag) GetStartTimeouts() uint32 {
	if m != nil {
		return m.StartTimeouts
	}
	return 0
}

func (m *BmdBfdCounterBag) GetNeighborUnconfiguredTimeouts() uint32 {
	if m != nil {
		return m.NeighborUnconfiguredTimeouts
	}
	return 0
}

func (m *BmdBfdCounterBag) GetTimeSinceCleared() uint64 {
	if m != nil {
		return m.TimeSinceCleared
	}
	return 0
}

type BmdBfdCounterBagMultiple struct {
	BundleName           *BmNameBag          `protobuf:"bytes,50,opt,name=bundle_name,json=bundleName,proto3" json:"bundle_name,omitempty"`
	BfdCounter           []*BmdBfdCounterBag `protobuf:"bytes,51,rep,name=bfd_counter,json=bfdCounter,proto3" json:"bfd_counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BmdBfdCounterBagMultiple) Reset()         { *m = BmdBfdCounterBagMultiple{} }
func (m *BmdBfdCounterBagMultiple) String() string { return proto.CompactTextString(m) }
func (*BmdBfdCounterBagMultiple) ProtoMessage()    {}
func (*BmdBfdCounterBagMultiple) Descriptor() ([]byte, []int) {
	return fileDescriptor_555c9eddef818db9, []int{3}
}

func (m *BmdBfdCounterBagMultiple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBfdCounterBagMultiple.Unmarshal(m, b)
}
func (m *BmdBfdCounterBagMultiple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBfdCounterBagMultiple.Marshal(b, m, deterministic)
}
func (m *BmdBfdCounterBagMultiple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBfdCounterBagMultiple.Merge(m, src)
}
func (m *BmdBfdCounterBagMultiple) XXX_Size() int {
	return xxx_messageInfo_BmdBfdCounterBagMultiple.Size(m)
}
func (m *BmdBfdCounterBagMultiple) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBfdCounterBagMultiple.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBfdCounterBagMultiple proto.InternalMessageInfo

func (m *BmdBfdCounterBagMultiple) GetBundleName() *BmNameBag {
	if m != nil {
		return m.BundleName
	}
	return nil
}

func (m *BmdBfdCounterBagMultiple) GetBfdCounter() []*BmdBfdCounterBag {
	if m != nil {
		return m.BfdCounter
	}
	return nil
}

func init() {
	proto.RegisterType((*BmdBfdCounterBagMultiple_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bfd_counters.bfd_counters_bundles.bfd_counters_bundle.bfd_counters_bundle_descendant.bmd_bfd_counter_bag_multiple_KEYS")
	proto.RegisterType((*BmNameBag)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bfd_counters.bfd_counters_bundles.bfd_counters_bundle.bfd_counters_bundle_descendant.bm_name_bag")
	proto.RegisterType((*BmdBfdCounterBag)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bfd_counters.bfd_counters_bundles.bfd_counters_bundle.bfd_counters_bundle_descendant.bmd_bfd_counter_bag")
	proto.RegisterType((*BmdBfdCounterBagMultiple)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bfd_counters.bfd_counters_bundles.bfd_counters_bundle.bfd_counters_bundle_descendant.bmd_bfd_counter_bag_multiple")
}

func init() { proto.RegisterFile("bmd_bfd_counter_bag_multiple.proto", fileDescriptor_555c9eddef818db9) }

var fileDescriptor_555c9eddef818db9 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0xdd, 0xa5, 0xec, 0x4e, 0xd4, 0x3f, 0x18, 0x90, 0x22, 0xa8, 0x20, 0x44, 0x42,
	0x0a, 0x15, 0xca, 0x61, 0xf7, 0x11, 0x0a, 0x07, 0x84, 0xd4, 0x43, 0x5a, 0x0e, 0x9c, 0x2c, 0x3b,
	0x99, 0x04, 0x4b, 0xb1, 0x1d, 0x39, 0x8e, 0xe0, 0xce, 0x0b, 0xf4, 0xc8, 0x73, 0x70, 0xe4, 0xb1,
	0x78, 0x02, 0x64, 0x27, 0xdd, 0x5d, 0xd4, 0xa8, 0xe7, 0xbd, 0x65, 0x7e, 0x9f, 0xbf, 0x99, 0xd1,
	0x67, 0x2b, 0x90, 0x70, 0x59, 0x52, 0x5e, 0x95, 0xb4, 0xd0, 0xbd, 0xb2, 0x68, 0x28, 0x67, 0x35,
	0x95, 0x7d, 0x63, 0x45, 0xdb, 0x60, 0xd6, 0x1a, 0x6d, 0x35, 0xf9, 0x19, 0x14, 0xa2, 0x2b, 0x34,
	0x15, 0xba, 0xa3, 0x3f, 0x0c, 0xe5, 0xbd, 0x2a, 0x1b, 0x94, 0xb5, 0xa1, 0xba, 0x45, 0x93, 0x0d,
	0x25, 0x15, 0xaa, 0xd2, 0x46, 0x32, 0x2b, 0xb4, 0xca, 0xf6, 0xfa, 0x75, 0xff, 0x15, 0xa3, 0x77,
	0x12, 0x4e, 0x31, 0x5a, 0x62, 0x57, 0xa0, 0x2a, 0x99, 0xb2, 0xc9, 0x15, 0xbc, 0x79, 0x68, 0x57,
	0xfa, 0xf9, 0xe3, 0xd7, 0x6b, 0xf2, 0x0e, 0xce, 0xb6, 0xfb, 0x58, 0x34, 0x15, 0x2b, 0x30, 0x0a,
	0xe2, 0x20, 0x5d, 0xe5, 0xa7, 0x03, 0xff, 0x74, 0x87, 0x93, 0x0b, 0x08, 0xb9, 0xa4, 0x8a, 0x49,
	0x74, 0x7d, 0xc8, 0x4b, 0x58, 0x09, 0x8b, 0x03, 0x18, 0x2d, 0x4b, 0x07, 0xae, 0x98, 0xc4, 0xe4,
	0xef, 0x0c, 0x9e, 0x4e, 0x0c, 0x27, 0xaf, 0x21, 0x94, 0x28, 0x39, 0x9a, 0x7d, 0x1b, 0x0c, 0xc8,
	0x19, 0xc9, 0x05, 0x3c, 0x69, 0x58, 0x67, 0xa9, 0x15, 0x12, 0x69, 0xd1, 0x20, 0x33, 0x58, 0x46,
	0xb3, 0x38, 0x48, 0x17, 0xf9, 0xa9, 0x13, 0x6e, 0x84, 0xc4, 0xcb, 0x01, 0x93, 0x17, 0xb0, 0xec,
	0x2c, 0x33, 0x56, 0xa8, 0x3a, 0x9a, 0xc7, 0x41, 0x7a, 0x9c, 0x6f, 0x6b, 0x72, 0x02, 0xb3, 0xbe,
	0x8d, 0x16, 0x9e, 0xce, 0xfa, 0x96, 0x10, 0x58, 0x94, 0xfa, 0xbb, 0x8a, 0x1e, 0x79, 0xe2, 0xbf,
	0xc9, 0x06, 0x9e, 0x2b, 0x14, 0xf5, 0x37, 0xae, 0x0d, 0xed, 0x55, 0xa1, 0x55, 0x25, 0xea, 0xde,
	0xcd, 0x3b, 0xf2, 0x87, 0x9e, 0xdd, 0x89, 0x5f, 0xf6, 0x34, 0xf2, 0x16, 0x4e, 0xfc, 0x10, 0xbf,
	0xa1, 0xee, 0x6d, 0x17, 0x3d, 0xf6, 0xa7, 0x8f, 0x3d, 0xbd, 0x19, 0x21, 0xf9, 0x00, 0xaf, 0x26,
	0x7b, 0xef, 0x6c, 0x4b, 0x6f, 0x3b, 0x9f, 0x1a, 0xb2, 0xed, 0xf2, 0x1e, 0x88, 0x0f, 0xa2, 0x13,
	0xaa, 0xd8, 0xc5, 0xb1, 0xf2, 0x71, 0x9c, 0x39, 0xe5, 0xda, 0x09, 0x63, 0x1e, 0xc9, 0xed, 0x1c,
	0xce, 0x1f, 0xba, 0x71, 0xf2, 0x3b, 0x80, 0x70, 0xbc, 0x6d, 0x1f, 0xff, 0x3a, 0x0e, 0xd2, 0x70,
	0x7d, 0x1b, 0x64, 0x07, 0xf0, 0x5c, 0xb3, 0xbd, 0xb7, 0x95, 0xc3, 0x20, 0xfb, 0x17, 0xf1, 0xc7,
	0x2d, 0xbd, 0xb3, 0x46, 0x9b, 0x78, 0x9e, 0x86, 0xeb, 0x5f, 0x87, 0xb2, 0xf4, 0xbd, 0xb8, 0x73,
	0xe0, 0x55, 0x79, 0x39, 0xd4, 0xfc, 0xc8, 0xff, 0x10, 0x36, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xd9, 0xbd, 0xcb, 0x13, 0x36, 0x04, 0x00, 0x00,
}
