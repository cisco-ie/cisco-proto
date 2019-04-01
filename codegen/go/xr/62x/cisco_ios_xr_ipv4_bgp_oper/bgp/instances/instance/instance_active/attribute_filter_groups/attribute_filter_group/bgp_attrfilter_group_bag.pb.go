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
// source: bgp_attrfilter_group_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_attribute_filter_groups_attribute_filter_group

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

type BgpAttrfilterGroupBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	GroupName            string   `protobuf:"bytes,2,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpAttrfilterGroupBag_KEYS) Reset()         { *m = BgpAttrfilterGroupBag_KEYS{} }
func (m *BgpAttrfilterGroupBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpAttrfilterGroupBag_KEYS) ProtoMessage()    {}
func (*BgpAttrfilterGroupBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e788e2cc752bf0dc, []int{0}
}

func (m *BgpAttrfilterGroupBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpAttrfilterGroupBag_KEYS.Unmarshal(m, b)
}
func (m *BgpAttrfilterGroupBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpAttrfilterGroupBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpAttrfilterGroupBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpAttrfilterGroupBag_KEYS.Merge(m, src)
}
func (m *BgpAttrfilterGroupBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpAttrfilterGroupBag_KEYS.Size(m)
}
func (m *BgpAttrfilterGroupBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpAttrfilterGroupBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpAttrfilterGroupBag_KEYS proto.InternalMessageInfo

func (m *BgpAttrfilterGroupBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpAttrfilterGroupBag_KEYS) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type BgpAttrfilterEntryBag struct {
	AttributeFilterEntryRequestedAction string   `protobuf:"bytes,1,opt,name=attribute_filter_entry_requested_action,json=attributeFilterEntryRequestedAction,proto3" json:"attribute_filter_entry_requested_action,omitempty"`
	AttributeFilterEntryRangeStart      uint32   `protobuf:"varint,2,opt,name=attribute_filter_entry_range_start,json=attributeFilterEntryRangeStart,proto3" json:"attribute_filter_entry_range_start,omitempty"`
	AttributeFilterEntryRangeEnd        uint32   `protobuf:"varint,3,opt,name=attribute_filter_entry_range_end,json=attributeFilterEntryRangeEnd,proto3" json:"attribute_filter_entry_range_end,omitempty"`
	XXX_NoUnkeyedLiteral                struct{} `json:"-"`
	XXX_unrecognized                    []byte   `json:"-"`
	XXX_sizecache                       int32    `json:"-"`
}

func (m *BgpAttrfilterEntryBag) Reset()         { *m = BgpAttrfilterEntryBag{} }
func (m *BgpAttrfilterEntryBag) String() string { return proto.CompactTextString(m) }
func (*BgpAttrfilterEntryBag) ProtoMessage()    {}
func (*BgpAttrfilterEntryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_e788e2cc752bf0dc, []int{1}
}

func (m *BgpAttrfilterEntryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpAttrfilterEntryBag.Unmarshal(m, b)
}
func (m *BgpAttrfilterEntryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpAttrfilterEntryBag.Marshal(b, m, deterministic)
}
func (m *BgpAttrfilterEntryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpAttrfilterEntryBag.Merge(m, src)
}
func (m *BgpAttrfilterEntryBag) XXX_Size() int {
	return xxx_messageInfo_BgpAttrfilterEntryBag.Size(m)
}
func (m *BgpAttrfilterEntryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpAttrfilterEntryBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpAttrfilterEntryBag proto.InternalMessageInfo

func (m *BgpAttrfilterEntryBag) GetAttributeFilterEntryRequestedAction() string {
	if m != nil {
		return m.AttributeFilterEntryRequestedAction
	}
	return ""
}

func (m *BgpAttrfilterEntryBag) GetAttributeFilterEntryRangeStart() uint32 {
	if m != nil {
		return m.AttributeFilterEntryRangeStart
	}
	return 0
}

func (m *BgpAttrfilterEntryBag) GetAttributeFilterEntryRangeEnd() uint32 {
	if m != nil {
		return m.AttributeFilterEntryRangeEnd
	}
	return 0
}

type BgpAttrfilterGroupBag struct {
	AttributeFilterGroupName       string                   `protobuf:"bytes,50,opt,name=attribute_filter_group_name,json=attributeFilterGroupName,proto3" json:"attribute_filter_group_name,omitempty"`
	AttributeFilterEntry           []*BgpAttrfilterEntryBag `protobuf:"bytes,51,rep,name=attribute_filter_entry,json=attributeFilterEntry,proto3" json:"attribute_filter_entry,omitempty"`
	AttributeFilterTotalGroupCount uint32                   `protobuf:"varint,52,opt,name=attribute_filter_total_group_count,json=attributeFilterTotalGroupCount,proto3" json:"attribute_filter_total_group_count,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}                 `json:"-"`
	XXX_unrecognized               []byte                   `json:"-"`
	XXX_sizecache                  int32                    `json:"-"`
}

func (m *BgpAttrfilterGroupBag) Reset()         { *m = BgpAttrfilterGroupBag{} }
func (m *BgpAttrfilterGroupBag) String() string { return proto.CompactTextString(m) }
func (*BgpAttrfilterGroupBag) ProtoMessage()    {}
func (*BgpAttrfilterGroupBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_e788e2cc752bf0dc, []int{2}
}

func (m *BgpAttrfilterGroupBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpAttrfilterGroupBag.Unmarshal(m, b)
}
func (m *BgpAttrfilterGroupBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpAttrfilterGroupBag.Marshal(b, m, deterministic)
}
func (m *BgpAttrfilterGroupBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpAttrfilterGroupBag.Merge(m, src)
}
func (m *BgpAttrfilterGroupBag) XXX_Size() int {
	return xxx_messageInfo_BgpAttrfilterGroupBag.Size(m)
}
func (m *BgpAttrfilterGroupBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpAttrfilterGroupBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpAttrfilterGroupBag proto.InternalMessageInfo

func (m *BgpAttrfilterGroupBag) GetAttributeFilterGroupName() string {
	if m != nil {
		return m.AttributeFilterGroupName
	}
	return ""
}

func (m *BgpAttrfilterGroupBag) GetAttributeFilterEntry() []*BgpAttrfilterEntryBag {
	if m != nil {
		return m.AttributeFilterEntry
	}
	return nil
}

func (m *BgpAttrfilterGroupBag) GetAttributeFilterTotalGroupCount() uint32 {
	if m != nil {
		return m.AttributeFilterTotalGroupCount
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpAttrfilterGroupBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.attribute_filter_groups.attribute_filter_group.bgp_attrfilter_group_bag_KEYS")
	proto.RegisterType((*BgpAttrfilterEntryBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.attribute_filter_groups.attribute_filter_group.bgp_attrfilter_entry_bag")
	proto.RegisterType((*BgpAttrfilterGroupBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.attribute_filter_groups.attribute_filter_group.bgp_attrfilter_group_bag")
}

func init() { proto.RegisterFile("bgp_attrfilter_group_bag.proto", fileDescriptor_e788e2cc752bf0dc) }

var fileDescriptor_e788e2cc752bf0dc = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0x06, 0xc2, 0xa2, 0xbb, 0x04, 0x91, 0x82, 0x6e, 0x8c, 0xee, 0xe0, 0x4e, 0x3d,
	0x6c, 0xbb, 0x7a, 0x10, 0xd9, 0x04, 0x05, 0x0f, 0xdd, 0x2e, 0x9e, 0x42, 0xda, 0x3d, 0x4b, 0x60,
	0x4b, 0x62, 0xf2, 0x3a, 0xf4, 0xee, 0xc9, 0x6f, 0xe5, 0xc7, 0xf1, 0x5b, 0x48, 0xb2, 0xb5, 0xc8,
	0xd6, 0x7a, 0xf4, 0x16, 0x5e, 0x7e, 0xf9, 0xff, 0xdf, 0xfb, 0xbf, 0x96, 0xf4, 0xd3, 0x5c, 0x33,
	0x8e, 0x68, 0x5e, 0xc4, 0x1a, 0xc1, 0xb0, 0xdc, 0xa8, 0x42, 0xb3, 0x94, 0xe7, 0xb1, 0x36, 0x0a,
	0x15, 0xd5, 0x99, 0xb0, 0x99, 0x62, 0x42, 0x59, 0xf6, 0x66, 0x98, 0xd0, 0xdb, 0x29, 0x73, 0x2f,
	0x94, 0x06, 0x13, 0xa7, 0xb9, 0x8e, 0x85, 0xb4, 0xc8, 0x65, 0x06, 0xb6, 0x3a, 0x55, 0x07, 0xc6,
	0x33, 0x14, 0x5b, 0x88, 0x9d, 0xba, 0x48, 0x0b, 0x04, 0xf6, 0xdb, 0xc3, 0x36, 0xd4, 0xa3, 0x8c,
	0xf4, 0x9a, 0x7a, 0x62, 0x8f, 0xb3, 0xe7, 0x05, 0x1d, 0x92, 0x6e, 0x65, 0x21, 0xf9, 0x06, 0xc2,
	0x60, 0x10, 0x8c, 0x3a, 0xc9, 0x59, 0x59, 0x7c, 0xe2, 0x1b, 0xa0, 0x3d, 0x42, 0x76, 0xcf, 0x3c,
	0xd1, 0xf2, 0x44, 0xc7, 0x57, 0xdc, 0x75, 0xf4, 0xd1, 0x22, 0xe1, 0x81, 0x0b, 0x48, 0x34, 0xef,
	0xce, 0x85, 0x2e, 0xc9, 0xf5, 0x51, 0x6f, 0xbb, 0x5b, 0x03, 0xaf, 0x05, 0x58, 0x84, 0x95, 0x9f,
	0x4d, 0xc9, 0xbd, 0xf5, 0xb0, 0xc2, 0xe7, 0x9e, 0x9e, 0x39, 0x38, 0x29, 0xd9, 0x5b, 0x8f, 0xd2,
	0x07, 0x12, 0x35, 0xa9, 0x72, 0x99, 0x03, 0xb3, 0xc8, 0x0d, 0xfa, 0x4e, 0xbb, 0x49, 0xbf, 0x56,
	0xd0, 0x61, 0x0b, 0x47, 0xd1, 0x39, 0x19, 0xfc, 0xa9, 0x05, 0x72, 0x15, 0xb6, 0xbd, 0xd2, 0x55,
	0xa3, 0xd2, 0x4c, 0xae, 0xa2, 0xef, 0xe3, 0x18, 0xaa, 0xb0, 0xe9, 0x0d, 0xb9, 0xac, 0x5f, 0xd1,
	0x2e, 0xd3, 0xb1, 0x1f, 0x3d, 0x3c, 0xd0, 0xbf, 0x2f, 0x23, 0xa6, 0x5f, 0x01, 0xb9, 0xa8, 0x6f,
	0x32, 0x9c, 0x0c, 0xda, 0xa3, 0xd3, 0xf1, 0x67, 0x10, 0xff, 0xf7, 0xc7, 0x15, 0x37, 0xed, 0x3c,
	0x39, 0xaf, 0xcb, 0xa9, 0x76, 0x67, 0xa8, 0x90, 0xaf, 0xf7, 0x41, 0x64, 0xaa, 0x90, 0x18, 0x4e,
	0x6b, 0x77, 0xb6, 0x74, 0x9c, 0x8f, 0xe3, 0xce, 0x51, 0xe9, 0x89, 0xff, 0xa1, 0x26, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xf2, 0x78, 0x0e, 0xa9, 0x72, 0x03, 0x00, 0x00,
}
