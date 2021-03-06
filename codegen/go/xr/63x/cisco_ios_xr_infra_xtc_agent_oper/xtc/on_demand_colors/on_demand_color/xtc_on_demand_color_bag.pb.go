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
// source: xtc_on_demand_color_bag.proto

package cisco_ios_xr_infra_xtc_agent_oper_xtc_on_demand_colors_on_demand_color

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

type XtcOnDemandColorBag_KEYS struct {
	Color                uint32   `protobuf:"varint,1,opt,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcOnDemandColorBag_KEYS) Reset()         { *m = XtcOnDemandColorBag_KEYS{} }
func (m *XtcOnDemandColorBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*XtcOnDemandColorBag_KEYS) ProtoMessage()    {}
func (*XtcOnDemandColorBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_342dd36a2fe8b485, []int{0}
}

func (m *XtcOnDemandColorBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcOnDemandColorBag_KEYS.Unmarshal(m, b)
}
func (m *XtcOnDemandColorBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcOnDemandColorBag_KEYS.Marshal(b, m, deterministic)
}
func (m *XtcOnDemandColorBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcOnDemandColorBag_KEYS.Merge(m, src)
}
func (m *XtcOnDemandColorBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_XtcOnDemandColorBag_KEYS.Size(m)
}
func (m *XtcOnDemandColorBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcOnDemandColorBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_XtcOnDemandColorBag_KEYS proto.InternalMessageInfo

func (m *XtcOnDemandColorBag_KEYS) GetColor() uint32 {
	if m != nil {
		return m.Color
	}
	return 0
}

type XtcDistInfoBag struct {
	DisjointnessType     string   `protobuf:"bytes,1,opt,name=disjointness_type,json=disjointnessType,proto3" json:"disjointness_type,omitempty"`
	GroupId              uint32   `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	SubId                uint32   `protobuf:"varint,3,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcDistInfoBag) Reset()         { *m = XtcDistInfoBag{} }
func (m *XtcDistInfoBag) String() string { return proto.CompactTextString(m) }
func (*XtcDistInfoBag) ProtoMessage()    {}
func (*XtcDistInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_342dd36a2fe8b485, []int{1}
}

func (m *XtcDistInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcDistInfoBag.Unmarshal(m, b)
}
func (m *XtcDistInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcDistInfoBag.Marshal(b, m, deterministic)
}
func (m *XtcDistInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcDistInfoBag.Merge(m, src)
}
func (m *XtcDistInfoBag) XXX_Size() int {
	return xxx_messageInfo_XtcDistInfoBag.Size(m)
}
func (m *XtcDistInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcDistInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcDistInfoBag proto.InternalMessageInfo

func (m *XtcDistInfoBag) GetDisjointnessType() string {
	if m != nil {
		return m.DisjointnessType
	}
	return ""
}

func (m *XtcDistInfoBag) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *XtcDistInfoBag) GetSubId() uint32 {
	if m != nil {
		return m.SubId
	}
	return 0
}

type XtcOnDemandColorBag struct {
	ColorXr              uint32          `protobuf:"varint,50,opt,name=color_xr,json=colorXr,proto3" json:"color_xr,omitempty"`
	DisjointPathInfo     *XtcDistInfoBag `protobuf:"bytes,51,opt,name=disjoint_path_info,json=disjointPathInfo,proto3" json:"disjoint_path_info,omitempty"`
	MaximumSidDepth      uint32          `protobuf:"varint,52,opt,name=maximum_sid_depth,json=maximumSidDepth,proto3" json:"maximum_sid_depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *XtcOnDemandColorBag) Reset()         { *m = XtcOnDemandColorBag{} }
func (m *XtcOnDemandColorBag) String() string { return proto.CompactTextString(m) }
func (*XtcOnDemandColorBag) ProtoMessage()    {}
func (*XtcOnDemandColorBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_342dd36a2fe8b485, []int{2}
}

func (m *XtcOnDemandColorBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcOnDemandColorBag.Unmarshal(m, b)
}
func (m *XtcOnDemandColorBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcOnDemandColorBag.Marshal(b, m, deterministic)
}
func (m *XtcOnDemandColorBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcOnDemandColorBag.Merge(m, src)
}
func (m *XtcOnDemandColorBag) XXX_Size() int {
	return xxx_messageInfo_XtcOnDemandColorBag.Size(m)
}
func (m *XtcOnDemandColorBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcOnDemandColorBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcOnDemandColorBag proto.InternalMessageInfo

func (m *XtcOnDemandColorBag) GetColorXr() uint32 {
	if m != nil {
		return m.ColorXr
	}
	return 0
}

func (m *XtcOnDemandColorBag) GetDisjointPathInfo() *XtcDistInfoBag {
	if m != nil {
		return m.DisjointPathInfo
	}
	return nil
}

func (m *XtcOnDemandColorBag) GetMaximumSidDepth() uint32 {
	if m != nil {
		return m.MaximumSidDepth
	}
	return 0
}

func init() {
	proto.RegisterType((*XtcOnDemandColorBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.on_demand_colors.on_demand_color.xtc_on_demand_color_bag_KEYS")
	proto.RegisterType((*XtcDistInfoBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.on_demand_colors.on_demand_color.xtc_dist_info_bag")
	proto.RegisterType((*XtcOnDemandColorBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.on_demand_colors.on_demand_color.xtc_on_demand_color_bag")
}

func init() { proto.RegisterFile("xtc_on_demand_color_bag.proto", fileDescriptor_342dd36a2fe8b485) }

var fileDescriptor_342dd36a2fe8b485 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x51, 0xdd, 0x4a, 0xc3, 0x30,
	0x18, 0xa5, 0x8a, 0x53, 0x23, 0xa2, 0x0b, 0x8a, 0x15, 0x14, 0x46, 0xaf, 0x86, 0x42, 0x2f, 0xb6,
	0x3d, 0x82, 0x0a, 0xc5, 0x1b, 0xd9, 0xbc, 0x70, 0x57, 0x1f, 0x69, 0x93, 0xb5, 0x11, 0x9b, 0x2f,
	0x24, 0x29, 0x64, 0x4f, 0xe0, 0xbb, 0xfa, 0x14, 0x92, 0x54, 0x51, 0x26, 0xbb, 0xf3, 0xf2, 0x9c,
	0xd3, 0xf3, 0xd3, 0x2f, 0xe4, 0xda, 0xbb, 0x0a, 0x50, 0x01, 0x17, 0x2d, 0x53, 0x1c, 0x2a, 0x7c,
	0x43, 0x03, 0x25, 0xab, 0x73, 0x6d, 0xd0, 0x21, 0x7d, 0xa8, 0xa4, 0xad, 0x10, 0x24, 0x5a, 0xf0,
	0x06, 0xa4, 0x5a, 0x19, 0x06, 0xc1, 0xc1, 0x6a, 0xa1, 0x1c, 0xa0, 0x16, 0x26, 0xf7, 0xae, 0xca,
	0x37, 0x02, 0xec, 0x26, 0x91, 0xcd, 0xc8, 0xd5, 0x96, 0x22, 0x78, 0xbc, 0x5f, 0x2e, 0xe8, 0x19,
	0xd9, 0x8b, 0x4c, 0x9a, 0x8c, 0x92, 0xf1, 0xf1, 0xbc, 0x07, 0x99, 0x21, 0xc3, 0xe0, 0xe2, 0xd2,
	0xba, 0xd0, 0x8d, 0xe1, 0x7b, 0x7a, 0x4b, 0x86, 0x5c, 0xda, 0x57, 0x94, 0xca, 0x29, 0x61, 0x2d,
	0xb8, 0xb5, 0x16, 0xd1, 0x76, 0x38, 0x3f, 0xfd, 0x2d, 0x3c, 0xaf, 0xb5, 0xa0, 0x97, 0xe4, 0xa0,
	0x36, 0xd8, 0x69, 0x90, 0x3c, 0xdd, 0x89, 0xd1, 0xfb, 0x11, 0x17, 0x9c, 0x9e, 0x93, 0x81, 0xed,
	0xca, 0x20, 0xec, 0xf6, 0x9d, 0xb6, 0x2b, 0x0b, 0x9e, 0x7d, 0x24, 0xe4, 0x62, 0xcb, 0xd4, 0x90,
	0xd6, 0x03, 0x6f, 0xd2, 0x49, 0x9f, 0x16, 0xf1, 0x8b, 0xa1, 0xef, 0x09, 0xa1, 0xdf, 0xed, 0xa0,
	0x99, 0x6b, 0xe2, 0xe0, 0x74, 0x3a, 0x4a, 0xc6, 0x47, 0x93, 0x65, 0xfe, 0x3f, 0x67, 0xcc, 0xff,
	0x5c, 0xe3, 0xe7, 0x97, 0x9f, 0x98, 0x6b, 0x0a, 0xb5, 0x42, 0x7a, 0x43, 0x86, 0x2d, 0xf3, 0xb2,
	0xed, 0x5a, 0xb0, 0x92, 0x03, 0x17, 0xda, 0x35, 0xe9, 0x2c, 0xae, 0x3d, 0xf9, 0x12, 0x16, 0x92,
	0xdf, 0x05, 0xba, 0x1c, 0xc4, 0x57, 0x9e, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xde, 0xc0,
	0x27, 0x06, 0x02, 0x00, 0x00,
}
