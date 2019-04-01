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
// source: bgp_rtset_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_afs_af_rt_set_counters_rt_set_counter

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

type BgpRtsetBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	RtSetId              uint32   `protobuf:"varint,3,opt,name=rt_set_id,json=rtSetId,proto3" json:"rt_set_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpRtsetBag_KEYS) Reset()         { *m = BgpRtsetBag_KEYS{} }
func (m *BgpRtsetBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpRtsetBag_KEYS) ProtoMessage()    {}
func (*BgpRtsetBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_82cb80d94ad8d5ee, []int{0}
}

func (m *BgpRtsetBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRtsetBag_KEYS.Unmarshal(m, b)
}
func (m *BgpRtsetBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRtsetBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpRtsetBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRtsetBag_KEYS.Merge(m, src)
}
func (m *BgpRtsetBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpRtsetBag_KEYS.Size(m)
}
func (m *BgpRtsetBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRtsetBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRtsetBag_KEYS proto.InternalMessageInfo

func (m *BgpRtsetBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpRtsetBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpRtsetBag_KEYS) GetRtSetId() uint32 {
	if m != nil {
		return m.RtSetId
	}
	return 0
}

type BgpRtset_ struct {
	RtSet                []uint32 `protobuf:"varint,1,rep,packed,name=rt_set,json=rtSet,proto3" json:"rt_set,omitempty"`
	RtSetLen             uint32   `protobuf:"varint,2,opt,name=rt_set_len,json=rtSetLen,proto3" json:"rt_set_len,omitempty"`
	RtSetId              uint32   `protobuf:"varint,3,opt,name=rt_set_id,json=rtSetId,proto3" json:"rt_set_id,omitempty"`
	RtSetNetCount        uint32   `protobuf:"varint,4,opt,name=rt_set_net_count,json=rtSetNetCount,proto3" json:"rt_set_net_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpRtset_) Reset()         { *m = BgpRtset_{} }
func (m *BgpRtset_) String() string { return proto.CompactTextString(m) }
func (*BgpRtset_) ProtoMessage()    {}
func (*BgpRtset_) Descriptor() ([]byte, []int) {
	return fileDescriptor_82cb80d94ad8d5ee, []int{1}
}

func (m *BgpRtset_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRtset_.Unmarshal(m, b)
}
func (m *BgpRtset_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRtset_.Marshal(b, m, deterministic)
}
func (m *BgpRtset_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRtset_.Merge(m, src)
}
func (m *BgpRtset_) XXX_Size() int {
	return xxx_messageInfo_BgpRtset_.Size(m)
}
func (m *BgpRtset_) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRtset_.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRtset_ proto.InternalMessageInfo

func (m *BgpRtset_) GetRtSet() []uint32 {
	if m != nil {
		return m.RtSet
	}
	return nil
}

func (m *BgpRtset_) GetRtSetLen() uint32 {
	if m != nil {
		return m.RtSetLen
	}
	return 0
}

func (m *BgpRtset_) GetRtSetId() uint32 {
	if m != nil {
		return m.RtSetId
	}
	return 0
}

func (m *BgpRtset_) GetRtSetNetCount() uint32 {
	if m != nil {
		return m.RtSetNetCount
	}
	return 0
}

type BgpRtsetBag struct {
	RouteTargetSet       *BgpRtset_ `protobuf:"bytes,50,opt,name=route_target_set,json=routeTargetSet,proto3" json:"route_target_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BgpRtsetBag) Reset()         { *m = BgpRtsetBag{} }
func (m *BgpRtsetBag) String() string { return proto.CompactTextString(m) }
func (*BgpRtsetBag) ProtoMessage()    {}
func (*BgpRtsetBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_82cb80d94ad8d5ee, []int{2}
}

func (m *BgpRtsetBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRtsetBag.Unmarshal(m, b)
}
func (m *BgpRtsetBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRtsetBag.Marshal(b, m, deterministic)
}
func (m *BgpRtsetBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRtsetBag.Merge(m, src)
}
func (m *BgpRtsetBag) XXX_Size() int {
	return xxx_messageInfo_BgpRtsetBag.Size(m)
}
func (m *BgpRtsetBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRtsetBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRtsetBag proto.InternalMessageInfo

func (m *BgpRtsetBag) GetRouteTargetSet() *BgpRtset_ {
	if m != nil {
		return m.RouteTargetSet
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpRtsetBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.rt_set_counters.rt_set_counter.bgp_rtset_bag_KEYS")
	proto.RegisterType((*BgpRtset_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.rt_set_counters.rt_set_counter.bgp_rtset_")
	proto.RegisterType((*BgpRtsetBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.rt_set_counters.rt_set_counter.bgp_rtset_bag")
}

func init() { proto.RegisterFile("bgp_rtset_bag.proto", fileDescriptor_82cb80d94ad8d5ee) }

var fileDescriptor_82cb80d94ad8d5ee = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x65, 0x0a, 0x2d, 0x3d, 0x08, 0xaa, 0x8c, 0x10, 0x11, 0x62, 0x88, 0xca, 0x40, 0x26,
	0x0f, 0x85, 0x37, 0x40, 0x0c, 0x08, 0xd4, 0x21, 0x65, 0x61, 0xb2, 0x9c, 0xf4, 0x1c, 0x59, 0x6a,
	0xed, 0xc8, 0xbe, 0x56, 0x6c, 0x6c, 0xbc, 0x0b, 0xcf, 0xc0, 0xcb, 0xa1, 0xba, 0x6d, 0xaa, 0x32,
	0xb0, 0xb2, 0x5d, 0xee, 0xff, 0xf5, 0xff, 0xdf, 0x29, 0x86, 0xf3, 0xb2, 0x6e, 0xa4, 0xa7, 0x80,
	0x24, 0x4b, 0x55, 0x8b, 0xc6, 0x3b, 0x72, 0x3c, 0x54, 0x26, 0x54, 0x4e, 0x1a, 0x17, 0xe4, 0xbb,
	0x97, 0xa6, 0x59, 0xde, 0xcb, 0x95, 0xcd, 0x35, 0xe8, 0x45, 0x59, 0x37, 0xc2, 0xd8, 0x40, 0xca,
	0x56, 0x18, 0xda, 0xa9, 0x1d, 0xa4, 0xaa, 0xc8, 0x2c, 0x51, 0x4c, 0x51, 0xab, 0xc5, 0x8c, 0xe4,
	0xd2, 0x6b, 0xa1, 0x74, 0x10, 0x4a, 0x0b, 0x4f, 0x72, 0x55, 0x53, 0xb9, 0x85, 0x25, 0xf4, 0xe1,
	0xd7, 0xf7, 0xd0, 0x02, 0xdf, 0x63, 0x91, 0xcf, 0x8f, 0x6f, 0x13, 0x7e, 0x03, 0x49, 0x1b, 0x6d,
	0xd5, 0x1c, 0x53, 0x96, 0xb1, 0xbc, 0x5f, 0x9c, 0x6e, 0x97, 0x63, 0x35, 0x47, 0x7e, 0x09, 0x3d,
	0xa5, 0xd7, 0xf2, 0x41, 0x94, 0xbb, 0x4a, 0x47, 0xe1, 0x0a, 0xfa, 0x9b, 0x16, 0x33, 0x4d, 0x3b,
	0x19, 0xcb, 0x93, 0xa2, 0xe7, 0x69, 0x82, 0xf4, 0x34, 0x1d, 0x7e, 0x32, 0x80, 0x5d, 0x21, 0xbf,
	0x80, 0xee, 0xda, 0x9a, 0xb2, 0xac, 0x93, 0x27, 0xc5, 0x51, 0xf4, 0xf1, 0x6b, 0x80, 0x4d, 0xc2,
	0x0c, 0x6d, 0x4c, 0x4f, 0x8a, 0xe3, 0x28, 0xbd, 0xa0, 0xfd, 0x2b, 0x9f, 0xdf, 0xc2, 0x60, 0xa3,
	0xd9, 0xed, 0x95, 0xe9, 0x61, 0xb4, 0x24, 0xd1, 0x32, 0x46, 0x7a, 0x58, 0x2d, 0x87, 0xdf, 0x0c,
	0x92, 0xbd, 0xcb, 0xf9, 0x17, 0x83, 0x81, 0x77, 0x0b, 0x42, 0x49, 0xca, 0xd7, 0xb8, 0xc6, 0x1a,
	0x65, 0x2c, 0x3f, 0x19, 0x7d, 0x88, 0x7f, 0xf8, 0x37, 0x62, 0x87, 0x57, 0x9c, 0x45, 0xb0, 0xd7,
	0xc8, 0x35, 0x41, 0x2a, 0xbb, 0xf1, 0xc9, 0xdc, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x05, 0x53,
	0x8e, 0x6c, 0x49, 0x02, 0x00, 0x00,
}
