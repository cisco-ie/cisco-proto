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
// source: igmp_edm_nsf_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_igmp_active_process_nsf

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

type IgmpEdmNsfBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmNsfBag_KEYS) Reset()         { *m = IgmpEdmNsfBag_KEYS{} }
func (m *IgmpEdmNsfBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmNsfBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmNsfBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_225a750b16131656, []int{0}
}

func (m *IgmpEdmNsfBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmNsfBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmNsfBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmNsfBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmNsfBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmNsfBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmNsfBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmNsfBag_KEYS.Size(m)
}
func (m *IgmpEdmNsfBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmNsfBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmNsfBag_KEYS proto.InternalMessageInfo

type IgmpEdmNsfBag struct {
	IsMulticastNsfActive bool     `protobuf:"varint,50,opt,name=is_multicast_nsf_active,json=isMulticastNsfActive,proto3" json:"is_multicast_nsf_active,omitempty"`
	MulticastNsfTimeout  uint32   `protobuf:"varint,51,opt,name=multicast_nsf_timeout,json=multicastNsfTimeout,proto3" json:"multicast_nsf_timeout,omitempty"`
	MulticastNsfTimeLeft uint32   `protobuf:"varint,52,opt,name=multicast_nsf_time_left,json=multicastNsfTimeLeft,proto3" json:"multicast_nsf_time_left,omitempty"`
	RespawnCount         uint32   `protobuf:"varint,53,opt,name=respawn_count,json=respawnCount,proto3" json:"respawn_count,omitempty"`
	LastNsfOn            string   `protobuf:"bytes,54,opt,name=last_nsf_on,json=lastNsfOn,proto3" json:"last_nsf_on,omitempty"`
	LastNsfOff           string   `protobuf:"bytes,55,opt,name=last_nsf_off,json=lastNsfOff,proto3" json:"last_nsf_off,omitempty"`
	LastNsfOnMin         int32    `protobuf:"zigzag32,56,opt,name=last_nsf_on_min,json=lastNsfOnMin,proto3" json:"last_nsf_on_min,omitempty"`
	LastNsfOffMin        int32    `protobuf:"zigzag32,57,opt,name=last_nsf_off_min,json=lastNsfOffMin,proto3" json:"last_nsf_off_min,omitempty"`
	LastIcdNotifRecv     string   `protobuf:"bytes,58,opt,name=last_icd_notif_recv,json=lastIcdNotifRecv,proto3" json:"last_icd_notif_recv,omitempty"`
	LastIcdNotifRecvMin  int32    `protobuf:"zigzag32,59,opt,name=last_icd_notif_recv_min,json=lastIcdNotifRecvMin,proto3" json:"last_icd_notif_recv_min,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmNsfBag) Reset()         { *m = IgmpEdmNsfBag{} }
func (m *IgmpEdmNsfBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmNsfBag) ProtoMessage()    {}
func (*IgmpEdmNsfBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_225a750b16131656, []int{1}
}

func (m *IgmpEdmNsfBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmNsfBag.Unmarshal(m, b)
}
func (m *IgmpEdmNsfBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmNsfBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmNsfBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmNsfBag.Merge(m, src)
}
func (m *IgmpEdmNsfBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmNsfBag.Size(m)
}
func (m *IgmpEdmNsfBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmNsfBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmNsfBag proto.InternalMessageInfo

func (m *IgmpEdmNsfBag) GetIsMulticastNsfActive() bool {
	if m != nil {
		return m.IsMulticastNsfActive
	}
	return false
}

func (m *IgmpEdmNsfBag) GetMulticastNsfTimeout() uint32 {
	if m != nil {
		return m.MulticastNsfTimeout
	}
	return 0
}

func (m *IgmpEdmNsfBag) GetMulticastNsfTimeLeft() uint32 {
	if m != nil {
		return m.MulticastNsfTimeLeft
	}
	return 0
}

func (m *IgmpEdmNsfBag) GetRespawnCount() uint32 {
	if m != nil {
		return m.RespawnCount
	}
	return 0
}

func (m *IgmpEdmNsfBag) GetLastNsfOn() string {
	if m != nil {
		return m.LastNsfOn
	}
	return ""
}

func (m *IgmpEdmNsfBag) GetLastNsfOff() string {
	if m != nil {
		return m.LastNsfOff
	}
	return ""
}

func (m *IgmpEdmNsfBag) GetLastNsfOnMin() int32 {
	if m != nil {
		return m.LastNsfOnMin
	}
	return 0
}

func (m *IgmpEdmNsfBag) GetLastNsfOffMin() int32 {
	if m != nil {
		return m.LastNsfOffMin
	}
	return 0
}

func (m *IgmpEdmNsfBag) GetLastIcdNotifRecv() string {
	if m != nil {
		return m.LastIcdNotifRecv
	}
	return ""
}

func (m *IgmpEdmNsfBag) GetLastIcdNotifRecvMin() int32 {
	if m != nil {
		return m.LastIcdNotifRecvMin
	}
	return 0
}

func init() {
	proto.RegisterType((*IgmpEdmNsfBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.process.nsf.igmp_edm_nsf_bag_KEYS")
	proto.RegisterType((*IgmpEdmNsfBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.process.nsf.igmp_edm_nsf_bag")
}

func init() { proto.RegisterFile("igmp_edm_nsf_bag.proto", fileDescriptor_225a750b16131656) }

var fileDescriptor_225a750b16131656 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0x09, 0x2f, 0xbc, 0xd8, 0xb1, 0xc5, 0x3a, 0x6d, 0xed, 0xac, 0x24, 0x54, 0xc4, 0x6c,
	0xcc, 0xa2, 0x1f, 0x7e, 0xae, 0x44, 0x5c, 0x88, 0xb6, 0x42, 0x74, 0xe3, 0x6a, 0x48, 0xa7, 0x33,
	0xe5, 0x42, 0x32, 0x13, 0x32, 0xd3, 0xe8, 0xcf, 0xf4, 0x27, 0x49, 0x6e, 0xad, 0x8d, 0xd1, 0xdd,
	0x70, 0xce, 0x73, 0x9e, 0xbb, 0x19, 0x72, 0x00, 0xcb, 0x34, 0xe3, 0x72, 0x91, 0x72, 0x6d, 0x15,
	0x9f, 0xc7, 0xcb, 0x30, 0xcb, 0x8d, 0x33, 0x74, 0x24, 0xc0, 0x0a, 0xc3, 0xc1, 0x58, 0xfe, 0x9e,
	0x73, 0xc8, 0x8a, 0x31, 0x47, 0xd2, 0x64, 0x32, 0x0f, 0xcb, 0x57, 0x18, 0x0b, 0x07, 0x85, 0x2c,
	0x71, 0x21, 0xad, 0x0d, 0xb5, 0x55, 0x83, 0x3e, 0xe9, 0xd5, 0x75, 0xfc, 0xe1, 0xee, 0xf5, 0x79,
	0xf0, 0xf1, 0x8f, 0xb4, 0xeb, 0x0d, 0x9d, 0x90, 0x3e, 0x58, 0x9e, 0xae, 0x12, 0x07, 0x22, 0xb6,
	0x0e, 0xf3, 0xb5, 0x94, 0x0d, 0x7d, 0x2f, 0xd8, 0x89, 0xba, 0x60, 0xa7, 0x9b, 0x76, 0x66, 0xd5,
	0x0d, 0x76, 0x74, 0x48, 0x7a, 0x3f, 0x37, 0x0e, 0x52, 0x69, 0x56, 0x8e, 0x8d, 0x7c, 0x2f, 0x68,
	0x45, 0x9d, 0xb4, 0x32, 0x79, 0x59, 0x57, 0xe5, 0xa9, 0xdf, 0x1b, 0x9e, 0x48, 0xe5, 0xd8, 0x18,
	0x57, 0xdd, 0xfa, 0xea, 0x51, 0x2a, 0x47, 0x8f, 0x48, 0x2b, 0x97, 0x36, 0x8b, 0xdf, 0x34, 0x17,
	0x66, 0xa5, 0x1d, 0x9b, 0x20, 0xdc, 0xfc, 0x0a, 0x6f, 0xcb, 0x8c, 0x1e, 0x92, 0xdd, 0x64, 0xa3,
	0x35, 0x9a, 0x9d, 0xf9, 0x5e, 0xd0, 0x88, 0x1a, 0xc9, 0x5a, 0xf5, 0xa4, 0xa9, 0x4f, 0x9a, 0xdb,
	0x5e, 0x29, 0x76, 0x8e, 0x00, 0xd9, 0x00, 0x4a, 0xd1, 0x63, 0xb2, 0x57, 0x31, 0xf0, 0x14, 0x34,
	0xbb, 0xf0, 0xbd, 0x60, 0x3f, 0x6a, 0x7e, 0x5b, 0xa6, 0xa0, 0xe9, 0x09, 0x69, 0x57, 0x45, 0xc8,
	0x5d, 0x22, 0xd7, 0xda, 0xca, 0x4a, 0xf0, 0x94, 0x74, 0x10, 0x04, 0xb1, 0xe0, 0xda, 0x38, 0x50,
	0x3c, 0x97, 0xa2, 0x60, 0x57, 0x78, 0x18, 0x1d, 0xf7, 0x62, 0x31, 0x2b, 0x8b, 0x48, 0x8a, 0x82,
	0x8e, 0x49, 0xff, 0x0f, 0x1c, 0xf5, 0xd7, 0xa8, 0xef, 0xd4, 0x27, 0x53, 0xd0, 0xf3, 0xff, 0xf8,
	0x4f, 0x46, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x30, 0xa8, 0x80, 0x41, 0x02, 0x00, 0x00,
}
