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
// source: bmd_event_bag.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_events_events_bundles_events_bundle_events_bundle_children_members_events_bundle_children_member

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

type BmdEventBag_KEYS struct {
	BundleInterface      string   `protobuf:"bytes,1,opt,name=bundle_interface,json=bundleInterface,proto3" json:"bundle_interface,omitempty"`
	MemberInterface      string   `protobuf:"bytes,2,opt,name=member_interface,json=memberInterface,proto3" json:"member_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdEventBag_KEYS) Reset()         { *m = BmdEventBag_KEYS{} }
func (m *BmdEventBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmdEventBag_KEYS) ProtoMessage()    {}
func (*BmdEventBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_22002626393f0e45, []int{0}
}

func (m *BmdEventBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdEventBag_KEYS.Unmarshal(m, b)
}
func (m *BmdEventBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdEventBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BmdEventBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdEventBag_KEYS.Merge(m, src)
}
func (m *BmdEventBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmdEventBag_KEYS.Size(m)
}
func (m *BmdEventBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdEventBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmdEventBag_KEYS proto.InternalMessageInfo

func (m *BmdEventBag_KEYS) GetBundleInterface() string {
	if m != nil {
		return m.BundleInterface
	}
	return ""
}

func (m *BmdEventBag_KEYS) GetMemberInterface() string {
	if m != nil {
		return m.MemberInterface
	}
	return ""
}

type BmdBagEventData struct {
	DataType             string   `protobuf:"bytes,1,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	NoData               uint32   `protobuf:"varint,2,opt,name=no_data,json=noData,proto3" json:"no_data,omitempty"`
	Error                uint32   `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
	StringData           string   `protobuf:"bytes,4,opt,name=string_data,json=stringData,proto3" json:"string_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdBagEventData) Reset()         { *m = BmdBagEventData{} }
func (m *BmdBagEventData) String() string { return proto.CompactTextString(m) }
func (*BmdBagEventData) ProtoMessage()    {}
func (*BmdBagEventData) Descriptor() ([]byte, []int) {
	return fileDescriptor_22002626393f0e45, []int{1}
}

func (m *BmdBagEventData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagEventData.Unmarshal(m, b)
}
func (m *BmdBagEventData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagEventData.Marshal(b, m, deterministic)
}
func (m *BmdBagEventData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagEventData.Merge(m, src)
}
func (m *BmdBagEventData) XXX_Size() int {
	return xxx_messageInfo_BmdBagEventData.Size(m)
}
func (m *BmdBagEventData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagEventData.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagEventData proto.InternalMessageInfo

func (m *BmdBagEventData) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *BmdBagEventData) GetNoData() uint32 {
	if m != nil {
		return m.NoData
	}
	return 0
}

func (m *BmdBagEventData) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *BmdBagEventData) GetStringData() string {
	if m != nil {
		return m.StringData
	}
	return ""
}

type BmdBagEventItemMbr struct {
	MemberEventType      string           `protobuf:"bytes,1,opt,name=member_event_type,json=memberEventType,proto3" json:"member_event_type,omitempty"`
	TimeStamp            uint64           `protobuf:"varint,2,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Data                 *BmdBagEventData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BmdBagEventItemMbr) Reset()         { *m = BmdBagEventItemMbr{} }
func (m *BmdBagEventItemMbr) String() string { return proto.CompactTextString(m) }
func (*BmdBagEventItemMbr) ProtoMessage()    {}
func (*BmdBagEventItemMbr) Descriptor() ([]byte, []int) {
	return fileDescriptor_22002626393f0e45, []int{2}
}

func (m *BmdBagEventItemMbr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagEventItemMbr.Unmarshal(m, b)
}
func (m *BmdBagEventItemMbr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagEventItemMbr.Marshal(b, m, deterministic)
}
func (m *BmdBagEventItemMbr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagEventItemMbr.Merge(m, src)
}
func (m *BmdBagEventItemMbr) XXX_Size() int {
	return xxx_messageInfo_BmdBagEventItemMbr.Size(m)
}
func (m *BmdBagEventItemMbr) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagEventItemMbr.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagEventItemMbr proto.InternalMessageInfo

func (m *BmdBagEventItemMbr) GetMemberEventType() string {
	if m != nil {
		return m.MemberEventType
	}
	return ""
}

func (m *BmdBagEventItemMbr) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *BmdBagEventItemMbr) GetData() *BmdBagEventData {
	if m != nil {
		return m.Data
	}
	return nil
}

type BmdBagEventItemBdl struct {
	BundleEventType      string           `protobuf:"bytes,1,opt,name=bundle_event_type,json=bundleEventType,proto3" json:"bundle_event_type,omitempty"`
	TimeStamp            uint64           `protobuf:"varint,2,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Data                 *BmdBagEventData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BmdBagEventItemBdl) Reset()         { *m = BmdBagEventItemBdl{} }
func (m *BmdBagEventItemBdl) String() string { return proto.CompactTextString(m) }
func (*BmdBagEventItemBdl) ProtoMessage()    {}
func (*BmdBagEventItemBdl) Descriptor() ([]byte, []int) {
	return fileDescriptor_22002626393f0e45, []int{3}
}

func (m *BmdBagEventItemBdl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagEventItemBdl.Unmarshal(m, b)
}
func (m *BmdBagEventItemBdl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagEventItemBdl.Marshal(b, m, deterministic)
}
func (m *BmdBagEventItemBdl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagEventItemBdl.Merge(m, src)
}
func (m *BmdBagEventItemBdl) XXX_Size() int {
	return xxx_messageInfo_BmdBagEventItemBdl.Size(m)
}
func (m *BmdBagEventItemBdl) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagEventItemBdl.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagEventItemBdl proto.InternalMessageInfo

func (m *BmdBagEventItemBdl) GetBundleEventType() string {
	if m != nil {
		return m.BundleEventType
	}
	return ""
}

func (m *BmdBagEventItemBdl) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *BmdBagEventItemBdl) GetData() *BmdBagEventData {
	if m != nil {
		return m.Data
	}
	return nil
}

type BmdBagEventItemRg struct {
	RgEventType          string           `protobuf:"bytes,1,opt,name=rg_event_type,json=rgEventType,proto3" json:"rg_event_type,omitempty"`
	TimeStamp            uint64           `protobuf:"varint,2,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Data                 *BmdBagEventData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BmdBagEventItemRg) Reset()         { *m = BmdBagEventItemRg{} }
func (m *BmdBagEventItemRg) String() string { return proto.CompactTextString(m) }
func (*BmdBagEventItemRg) ProtoMessage()    {}
func (*BmdBagEventItemRg) Descriptor() ([]byte, []int) {
	return fileDescriptor_22002626393f0e45, []int{4}
}

func (m *BmdBagEventItemRg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagEventItemRg.Unmarshal(m, b)
}
func (m *BmdBagEventItemRg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagEventItemRg.Marshal(b, m, deterministic)
}
func (m *BmdBagEventItemRg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagEventItemRg.Merge(m, src)
}
func (m *BmdBagEventItemRg) XXX_Size() int {
	return xxx_messageInfo_BmdBagEventItemRg.Size(m)
}
func (m *BmdBagEventItemRg) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagEventItemRg.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagEventItemRg proto.InternalMessageInfo

func (m *BmdBagEventItemRg) GetRgEventType() string {
	if m != nil {
		return m.RgEventType
	}
	return ""
}

func (m *BmdBagEventItemRg) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *BmdBagEventItemRg) GetData() *BmdBagEventData {
	if m != nil {
		return m.Data
	}
	return nil
}

type BmdBagEventItem struct {
	EventType            string              `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	MemberEvtInfo        *BmdBagEventItemMbr `protobuf:"bytes,2,opt,name=member_evt_info,json=memberEvtInfo,proto3" json:"member_evt_info,omitempty"`
	BundleEvtInfo        *BmdBagEventItemBdl `protobuf:"bytes,3,opt,name=bundle_evt_info,json=bundleEvtInfo,proto3" json:"bundle_evt_info,omitempty"`
	RgEvtInfo            *BmdBagEventItemRg  `protobuf:"bytes,4,opt,name=rg_evt_info,json=rgEvtInfo,proto3" json:"rg_evt_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BmdBagEventItem) Reset()         { *m = BmdBagEventItem{} }
func (m *BmdBagEventItem) String() string { return proto.CompactTextString(m) }
func (*BmdBagEventItem) ProtoMessage()    {}
func (*BmdBagEventItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_22002626393f0e45, []int{5}
}

func (m *BmdBagEventItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagEventItem.Unmarshal(m, b)
}
func (m *BmdBagEventItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagEventItem.Marshal(b, m, deterministic)
}
func (m *BmdBagEventItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagEventItem.Merge(m, src)
}
func (m *BmdBagEventItem) XXX_Size() int {
	return xxx_messageInfo_BmdBagEventItem.Size(m)
}
func (m *BmdBagEventItem) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagEventItem.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagEventItem proto.InternalMessageInfo

func (m *BmdBagEventItem) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *BmdBagEventItem) GetMemberEvtInfo() *BmdBagEventItemMbr {
	if m != nil {
		return m.MemberEvtInfo
	}
	return nil
}

func (m *BmdBagEventItem) GetBundleEvtInfo() *BmdBagEventItemBdl {
	if m != nil {
		return m.BundleEvtInfo
	}
	return nil
}

func (m *BmdBagEventItem) GetRgEvtInfo() *BmdBagEventItemRg {
	if m != nil {
		return m.RgEvtInfo
	}
	return nil
}

type BmdEventBag struct {
	ItemName             string             `protobuf:"bytes,50,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	Items                []*BmdBagEventItem `protobuf:"bytes,51,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BmdEventBag) Reset()         { *m = BmdEventBag{} }
func (m *BmdEventBag) String() string { return proto.CompactTextString(m) }
func (*BmdEventBag) ProtoMessage()    {}
func (*BmdEventBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_22002626393f0e45, []int{6}
}

func (m *BmdEventBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdEventBag.Unmarshal(m, b)
}
func (m *BmdEventBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdEventBag.Marshal(b, m, deterministic)
}
func (m *BmdEventBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdEventBag.Merge(m, src)
}
func (m *BmdEventBag) XXX_Size() int {
	return xxx_messageInfo_BmdEventBag.Size(m)
}
func (m *BmdEventBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdEventBag.DiscardUnknown(m)
}

var xxx_messageInfo_BmdEventBag proto.InternalMessageInfo

func (m *BmdEventBag) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

func (m *BmdEventBag) GetItems() []*BmdBagEventItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*BmdEventBag_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_children_members.events_bundle_children_member.bmd_event_bag_KEYS")
	proto.RegisterType((*BmdBagEventData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_children_members.events_bundle_children_member.bmd_bag_event_data")
	proto.RegisterType((*BmdBagEventItemMbr)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_children_members.events_bundle_children_member.bmd_bag_event_item_mbr")
	proto.RegisterType((*BmdBagEventItemBdl)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_children_members.events_bundle_children_member.bmd_bag_event_item_bdl")
	proto.RegisterType((*BmdBagEventItemRg)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_children_members.events_bundle_children_member.bmd_bag_event_item_rg")
	proto.RegisterType((*BmdBagEventItem)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_children_members.events_bundle_children_member.bmd_bag_event_item")
	proto.RegisterType((*BmdEventBag)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_children_members.events_bundle_children_member.bmd_event_bag")
}

func init() { proto.RegisterFile("bmd_event_bag.proto", fileDescriptor_22002626393f0e45) }

var fileDescriptor_22002626393f0e45 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x56, 0xbd, 0x6e, 0xd4, 0x40,
	0x10, 0xd6, 0xe6, 0xee, 0x02, 0x1e, 0xcb, 0x0a, 0x2c, 0x7f, 0x27, 0x21, 0x44, 0xe4, 0x0a, 0x28,
	0x5c, 0x5c, 0x5e, 0x81, 0x14, 0x11, 0x12, 0x85, 0x43, 0x43, 0xb5, 0xb2, 0xcf, 0x7b, 0xc6, 0xe8,
	0x76, 0xd7, 0x1a, 0x2f, 0x88, 0x94, 0x3c, 0x01, 0x15, 0x2d, 0x4d, 0x24, 0x78, 0x05, 0xc4, 0x0b,
	0xd0, 0xf2, 0x1a, 0xbc, 0x05, 0x9a, 0x59, 0xdb, 0xca, 0xe5, 0x80, 0x3a, 0xa7, 0x54, 0xd6, 0x7c,
	0x3b, 0x3b, 0xf3, 0x7d, 0xdf, 0xde, 0xec, 0x1e, 0xdc, 0x29, 0x4d, 0xa5, 0xf4, 0x7b, 0x6d, 0xbd,
	0x2a, 0x8b, 0x3a, 0x6b, 0xd1, 0x79, 0x27, 0x3f, 0x89, 0x65, 0xd3, 0x2d, 0x9d, 0x6a, 0x5c, 0xa7,
	0x3e, 0xa0, 0x2a, 0xdf, 0xd9, 0x6a, 0xad, 0x4d, 0x8d, 0xca, 0xb5, 0x1a, 0xb3, 0x10, 0xaa, 0xc6,
	0xae, 0x1c, 0x9a, 0xc2, 0x37, 0xce, 0x66, 0x5c, 0xa0, 0xeb, 0x3f, 0x7d, 0xfe, 0xa5, 0x70, 0x33,
	0x52, 0xcb, 0x37, 0xcd, 0xba, 0x42, 0x6d, 0x95, 0xd1, 0xa6, 0xd4, 0xd8, 0xfd, 0x7f, 0x39, 0x7d,
	0x0b, 0x72, 0x83, 0xa8, 0x7a, 0x71, 0xfc, 0xfa, 0x54, 0x3e, 0x85, 0x5b, 0x23, 0x19, 0xaf, 0x71,
	0x55, 0x2c, 0xf5, 0x5c, 0x1c, 0x8a, 0x27, 0x51, 0x7e, 0x10, 0xf0, 0x93, 0x01, 0xa6, 0xd4, 0x50,
	0xea, 0x42, 0xea, 0x5e, 0x48, 0x0d, 0xf8, 0x98, 0x9a, 0x7e, 0x14, 0xa1, 0x19, 0xb5, 0x09, 0x0d,
	0xab, 0xc2, 0x17, 0xf2, 0x21, 0x44, 0xf4, 0x55, 0xfe, 0xac, 0x1d, 0xba, 0xdc, 0x24, 0xe0, 0xd5,
	0x59, 0xab, 0xe5, 0x03, 0xb8, 0x61, 0x1d, 0xe7, 0x71, 0xd5, 0x24, 0xdf, 0xb7, 0xee, 0x39, 0xed,
	0xba, 0x0b, 0x33, 0x8d, 0xe8, 0x70, 0x3e, 0x61, 0x38, 0x04, 0xf2, 0x31, 0xc4, 0x9d, 0xc7, 0xc6,
	0xd6, 0x61, 0xcb, 0x94, 0xab, 0x41, 0x80, 0x68, 0x5b, 0xfa, 0x65, 0x0f, 0xee, 0x6f, 0x72, 0x68,
	0xbc, 0x36, 0xca, 0x94, 0x28, 0x9f, 0xc1, 0xed, 0x5e, 0x49, 0x58, 0xb8, 0xc0, 0xa7, 0x97, 0x72,
	0x4c, 0x38, 0xd3, 0x7a, 0x04, 0xe0, 0x1b, 0xa3, 0x55, 0xe7, 0x0b, 0xd3, 0x32, 0xb3, 0x69, 0x1e,
	0x11, 0x72, 0x4a, 0x80, 0xfc, 0x2e, 0x60, 0xca, 0x04, 0x88, 0x5c, 0xbc, 0x38, 0x17, 0xd9, 0x15,
	0x3b, 0xf7, 0x6c, 0xfb, 0x1c, 0x72, 0x66, 0xfc, 0x2f, 0x83, 0xca, 0x6a, 0x4d, 0x06, 0xf5, 0xd5,
	0xb6, 0x0d, 0x0a, 0x0b, 0xd7, 0xc1, 0xa0, 0xcf, 0x7b, 0x70, 0xef, 0x2f, 0x06, 0x61, 0x2d, 0x53,
	0x48, 0xb0, 0xde, 0xf6, 0x26, 0xc6, 0xfa, 0x3a, 0xf8, 0x72, 0x3e, 0xbb, 0x3c, 0xdd, 0xe4, 0x0b,
	0x09, 0xde, 0x72, 0x24, 0xd2, 0xa3, 0x1f, 0xbf, 0x04, 0x1c, 0x8c, 0x53, 0xe7, 0x59, 0x0a, 0xbb,
	0x12, 0x2f, 0xbe, 0x5d, 0x75, 0xed, 0xc3, 0xc5, 0x91, 0x27, 0xc3, 0xed, 0xe0, 0x4f, 0xec, 0xca,
	0xb1, 0xa4, 0x71, 0x4e, 0x7a, 0x49, 0x93, 0x1d, 0x92, 0x54, 0x56, 0xeb, 0x3c, 0x19, 0xe6, 0x39,
	0x48, 0xfa, 0x29, 0x20, 0xe6, 0x9f, 0x76, 0x2f, 0x67, 0xca, 0x72, 0xbe, 0xee, 0x84, 0x1c, 0xac,
	0xf3, 0x88, 0x26, 0x90, 0x95, 0xa4, 0xbf, 0x05, 0x24, 0x1b, 0x0f, 0x1e, 0x3d, 0x3f, 0x9c, 0x67,
	0x0b, 0xa3, 0xe7, 0x8b, 0xf0, 0xfc, 0x10, 0xf0, 0xb2, 0x30, 0x5a, 0xfe, 0x10, 0x30, 0xa3, 0xa0,
	0x9b, 0x1f, 0x1d, 0x4e, 0x76, 0x61, 0x20, 0x89, 0x6c, 0x1e, 0x28, 0x97, 0xfb, 0xfc, 0xa7, 0xe3,
	0xe8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x7c, 0x23, 0x85, 0x8b, 0x08, 0x00, 0x00,
}
