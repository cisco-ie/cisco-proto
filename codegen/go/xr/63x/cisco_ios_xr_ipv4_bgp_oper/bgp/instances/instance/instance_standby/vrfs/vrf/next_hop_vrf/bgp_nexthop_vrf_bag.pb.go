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
// source: bgp_nexthop_vrf_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_next_hop_vrf

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

type BgpNexthopVrfBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpNexthopVrfBag_KEYS) Reset()         { *m = BgpNexthopVrfBag_KEYS{} }
func (m *BgpNexthopVrfBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpNexthopVrfBag_KEYS) ProtoMessage()    {}
func (*BgpNexthopVrfBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa3f4373201fb254, []int{0}
}

func (m *BgpNexthopVrfBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNexthopVrfBag_KEYS.Unmarshal(m, b)
}
func (m *BgpNexthopVrfBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNexthopVrfBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpNexthopVrfBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNexthopVrfBag_KEYS.Merge(m, src)
}
func (m *BgpNexthopVrfBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpNexthopVrfBag_KEYS.Size(m)
}
func (m *BgpNexthopVrfBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNexthopVrfBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNexthopVrfBag_KEYS proto.InternalMessageInfo

func (m *BgpNexthopVrfBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpNexthopVrfBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type BgpNexthopVrfBag struct {
	TotalProcessingTime            uint32   `protobuf:"varint,50,opt,name=total_processing_time,json=totalProcessingTime,proto3" json:"total_processing_time,omitempty"`
	MaxProcNotificationTime        uint32   `protobuf:"varint,51,opt,name=max_proc_notification_time,json=maxProcNotificationTime,proto3" json:"max_proc_notification_time,omitempty"`
	MaxNotificationBestpathDeletes uint32   `protobuf:"varint,52,opt,name=max_notification_bestpath_deletes,json=maxNotificationBestpathDeletes,proto3" json:"max_notification_bestpath_deletes,omitempty"`
	MaxNotificationBestpathChanges uint32   `protobuf:"varint,53,opt,name=max_notification_bestpath_changes,json=maxNotificationBestpathChanges,proto3" json:"max_notification_bestpath_changes,omitempty"`
	MaximumProcessingTime          uint32   `protobuf:"varint,54,opt,name=maximum_processing_time,json=maximumProcessingTime,proto3" json:"maximum_processing_time,omitempty"`
	LastNotificationicationTime    uint32   `protobuf:"varint,55,opt,name=last_notificationication_time,json=lastNotificationicationTime,proto3" json:"last_notificationication_time,omitempty"`
	LastNotificationProcessingTime uint32   `protobuf:"varint,56,opt,name=last_notification_processing_time,json=lastNotificationProcessingTime,proto3" json:"last_notification_processing_time,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *BgpNexthopVrfBag) Reset()         { *m = BgpNexthopVrfBag{} }
func (m *BgpNexthopVrfBag) String() string { return proto.CompactTextString(m) }
func (*BgpNexthopVrfBag) ProtoMessage()    {}
func (*BgpNexthopVrfBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa3f4373201fb254, []int{1}
}

func (m *BgpNexthopVrfBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNexthopVrfBag.Unmarshal(m, b)
}
func (m *BgpNexthopVrfBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNexthopVrfBag.Marshal(b, m, deterministic)
}
func (m *BgpNexthopVrfBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNexthopVrfBag.Merge(m, src)
}
func (m *BgpNexthopVrfBag) XXX_Size() int {
	return xxx_messageInfo_BgpNexthopVrfBag.Size(m)
}
func (m *BgpNexthopVrfBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNexthopVrfBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNexthopVrfBag proto.InternalMessageInfo

func (m *BgpNexthopVrfBag) GetTotalProcessingTime() uint32 {
	if m != nil {
		return m.TotalProcessingTime
	}
	return 0
}

func (m *BgpNexthopVrfBag) GetMaxProcNotificationTime() uint32 {
	if m != nil {
		return m.MaxProcNotificationTime
	}
	return 0
}

func (m *BgpNexthopVrfBag) GetMaxNotificationBestpathDeletes() uint32 {
	if m != nil {
		return m.MaxNotificationBestpathDeletes
	}
	return 0
}

func (m *BgpNexthopVrfBag) GetMaxNotificationBestpathChanges() uint32 {
	if m != nil {
		return m.MaxNotificationBestpathChanges
	}
	return 0
}

func (m *BgpNexthopVrfBag) GetMaximumProcessingTime() uint32 {
	if m != nil {
		return m.MaximumProcessingTime
	}
	return 0
}

func (m *BgpNexthopVrfBag) GetLastNotificationicationTime() uint32 {
	if m != nil {
		return m.LastNotificationicationTime
	}
	return 0
}

func (m *BgpNexthopVrfBag) GetLastNotificationProcessingTime() uint32 {
	if m != nil {
		return m.LastNotificationProcessingTime
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpNexthopVrfBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.next_hop_vrf.bgp_nexthop_vrf_bag_KEYS")
	proto.RegisterType((*BgpNexthopVrfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.next_hop_vrf.bgp_nexthop_vrf_bag")
}

func init() { proto.RegisterFile("bgp_nexthop_vrf_bag.proto", fileDescriptor_aa3f4373201fb254) }

var fileDescriptor_aa3f4373201fb254 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0xe9, 0xfb, 0x82, 0x1f, 0x83, 0xdd, 0xa4, 0x14, 0x53, 0x45, 0xa9, 0x75, 0xd3, 0x55,
	0x16, 0x6d, 0xad, 0x82, 0xbb, 0xaa, 0x0b, 0x11, 0x8a, 0x54, 0x17, 0xea, 0xe6, 0x32, 0x49, 0x6f,
	0xd2, 0x81, 0xce, 0x07, 0x99, 0x31, 0xc4, 0xdf, 0xeb, 0x1f, 0x91, 0x99, 0xb4, 0xa1, 0x69, 0x4a,
	0x37, 0xc9, 0x90, 0x73, 0x9e, 0x73, 0x6f, 0x0e, 0x43, 0x3a, 0x61, 0xa2, 0x40, 0x60, 0x6e, 0x16,
	0x52, 0x41, 0x96, 0xc6, 0x10, 0xd2, 0x24, 0x50, 0xa9, 0x34, 0xd2, 0xfb, 0x88, 0x98, 0x8e, 0x24,
	0x30, 0xa9, 0x21, 0x4f, 0x81, 0xa9, 0x6c, 0x04, 0xd6, 0x2c, 0x15, 0xa6, 0x41, 0x98, 0xa8, 0x80,
	0x09, 0x6d, 0xa8, 0x88, 0x50, 0x97, 0xa7, 0xf2, 0x00, 0xf6, 0x35, 0x0f, 0x7f, 0x82, 0x2c, 0x8d,
	0xb5, 0x7d, 0x04, 0x76, 0x04, 0xac, 0x66, 0xf4, 0xbe, 0x88, 0xbf, 0x63, 0x2c, 0xbc, 0x3c, 0x7d,
	0xbe, 0x79, 0xd7, 0xa4, 0x59, 0xa6, 0x08, 0xca, 0xd1, 0x6f, 0x74, 0x1b, 0xfd, 0xe3, 0xd9, 0xc9,
	0xfa, 0xe3, 0x94, 0x72, 0xf4, 0x3a, 0xe4, 0xc8, 0x42, 0x4e, 0xff, 0xe7, 0xf4, 0xc3, 0x2c, 0x8d,
	0xad, 0xd4, 0xfb, 0xfd, 0x4f, 0x5a, 0x3b, 0xc2, 0xbd, 0x01, 0x69, 0x1b, 0x69, 0xe8, 0x12, 0x54,
	0x2a, 0x23, 0xd4, 0x9a, 0x89, 0x04, 0x0c, 0xe3, 0xe8, 0x0f, 0xba, 0x8d, 0x7e, 0x73, 0xd6, 0x72,
	0xe2, 0x6b, 0xa9, 0xbd, 0x33, 0x8e, 0xde, 0x3d, 0x39, 0xe3, 0x34, 0x77, 0x04, 0x08, 0x69, 0x58,
	0xcc, 0x22, 0x6a, 0x98, 0x14, 0x05, 0x38, 0x74, 0xe0, 0x29, 0xa7, 0xb9, 0xc5, 0xa6, 0x1b, 0xba,
	0x83, 0x9f, 0xc9, 0x95, 0x85, 0x2b, 0x5c, 0x88, 0xda, 0x28, 0x6a, 0x16, 0x30, 0xc7, 0x25, 0x1a,
	0xd4, 0xfe, 0xc8, 0x65, 0x5c, 0x72, 0x9a, 0x6f, 0xf2, 0x93, 0x95, 0xed, 0xb1, 0x70, 0xed, 0x8f,
	0x8a, 0x16, 0x54, 0x24, 0xa8, 0xfd, 0x9b, 0xbd, 0x51, 0x0f, 0x85, 0xcb, 0x1b, 0x13, 0xbb, 0x30,
	0xe3, 0xdf, 0xbc, 0x56, 0xc4, 0xd8, 0x05, 0xb4, 0x57, 0xf2, 0x56, 0x15, 0x13, 0x72, 0xb1, 0xa4,
	0xda, 0x54, 0x76, 0xa8, 0xb4, 0x71, 0xeb, 0xe8, 0x73, 0x6b, 0x9a, 0xd6, 0x3d, 0xeb, 0x46, 0x6a,
	0x19, 0xb5, 0x2d, 0xee, 0x8a, 0xdf, 0xd8, 0xce, 0xa9, 0xae, 0x13, 0x1e, 0xb8, 0x2b, 0x3a, 0xfc,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x86, 0x1d, 0x09, 0xbf, 0x02, 0x00, 0x00,
}
