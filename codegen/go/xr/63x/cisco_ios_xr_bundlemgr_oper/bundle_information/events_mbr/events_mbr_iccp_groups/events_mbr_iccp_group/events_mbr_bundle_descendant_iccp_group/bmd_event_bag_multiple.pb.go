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
// source: bmd_event_bag_multiple.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_events_mbr_events_mbr_iccp_groups_events_mbr_iccp_group_events_mbr_bundle_descendant_iccp_group

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

type BmdEventBagMultiple_KEYS struct {
	IccpGroup            uint32   `protobuf:"varint,1,opt,name=iccp_group,json=iccpGroup,proto3" json:"iccp_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdEventBagMultiple_KEYS) Reset()         { *m = BmdEventBagMultiple_KEYS{} }
func (m *BmdEventBagMultiple_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmdEventBagMultiple_KEYS) ProtoMessage()    {}
func (*BmdEventBagMultiple_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb75e524eb4a446, []int{0}
}

func (m *BmdEventBagMultiple_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdEventBagMultiple_KEYS.Unmarshal(m, b)
}
func (m *BmdEventBagMultiple_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdEventBagMultiple_KEYS.Marshal(b, m, deterministic)
}
func (m *BmdEventBagMultiple_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdEventBagMultiple_KEYS.Merge(m, src)
}
func (m *BmdEventBagMultiple_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmdEventBagMultiple_KEYS.Size(m)
}
func (m *BmdEventBagMultiple_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdEventBagMultiple_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmdEventBagMultiple_KEYS proto.InternalMessageInfo

func (m *BmdEventBagMultiple_KEYS) GetIccpGroup() uint32 {
	if m != nil {
		return m.IccpGroup
	}
	return 0
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
	return fileDescriptor_2eb75e524eb4a446, []int{1}
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
	return fileDescriptor_2eb75e524eb4a446, []int{2}
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
	return fileDescriptor_2eb75e524eb4a446, []int{3}
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
	return fileDescriptor_2eb75e524eb4a446, []int{4}
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
	return fileDescriptor_2eb75e524eb4a446, []int{5}
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
	ItemName             string             `protobuf:"bytes,1,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	Items                []*BmdBagEventItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BmdEventBag) Reset()         { *m = BmdEventBag{} }
func (m *BmdEventBag) String() string { return proto.CompactTextString(m) }
func (*BmdEventBag) ProtoMessage()    {}
func (*BmdEventBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb75e524eb4a446, []int{6}
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

type BmdEventBagMultiple struct {
	EventsItem           []*BmdEventBag `protobuf:"bytes,50,rep,name=events_item,json=eventsItem,proto3" json:"events_item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BmdEventBagMultiple) Reset()         { *m = BmdEventBagMultiple{} }
func (m *BmdEventBagMultiple) String() string { return proto.CompactTextString(m) }
func (*BmdEventBagMultiple) ProtoMessage()    {}
func (*BmdEventBagMultiple) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb75e524eb4a446, []int{7}
}

func (m *BmdEventBagMultiple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdEventBagMultiple.Unmarshal(m, b)
}
func (m *BmdEventBagMultiple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdEventBagMultiple.Marshal(b, m, deterministic)
}
func (m *BmdEventBagMultiple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdEventBagMultiple.Merge(m, src)
}
func (m *BmdEventBagMultiple) XXX_Size() int {
	return xxx_messageInfo_BmdEventBagMultiple.Size(m)
}
func (m *BmdEventBagMultiple) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdEventBagMultiple.DiscardUnknown(m)
}

var xxx_messageInfo_BmdEventBagMultiple proto.InternalMessageInfo

func (m *BmdEventBagMultiple) GetEventsItem() []*BmdEventBag {
	if m != nil {
		return m.EventsItem
	}
	return nil
}

func init() {
	proto.RegisterType((*BmdEventBagMultiple_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events_mbr.events_mbr_iccp_groups.events_mbr_iccp_group.events_mbr_bundle_descendant_iccp_group.bmd_event_bag_multiple_KEYS")
	proto.RegisterType((*BmdBagEventData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events_mbr.events_mbr_iccp_groups.events_mbr_iccp_group.events_mbr_bundle_descendant_iccp_group.bmd_bag_event_data")
	proto.RegisterType((*BmdBagEventItemMbr)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events_mbr.events_mbr_iccp_groups.events_mbr_iccp_group.events_mbr_bundle_descendant_iccp_group.bmd_bag_event_item_mbr")
	proto.RegisterType((*BmdBagEventItemBdl)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events_mbr.events_mbr_iccp_groups.events_mbr_iccp_group.events_mbr_bundle_descendant_iccp_group.bmd_bag_event_item_bdl")
	proto.RegisterType((*BmdBagEventItemRg)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events_mbr.events_mbr_iccp_groups.events_mbr_iccp_group.events_mbr_bundle_descendant_iccp_group.bmd_bag_event_item_rg")
	proto.RegisterType((*BmdBagEventItem)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events_mbr.events_mbr_iccp_groups.events_mbr_iccp_group.events_mbr_bundle_descendant_iccp_group.bmd_bag_event_item")
	proto.RegisterType((*BmdEventBag)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events_mbr.events_mbr_iccp_groups.events_mbr_iccp_group.events_mbr_bundle_descendant_iccp_group.bmd_event_bag")
	proto.RegisterType((*BmdEventBagMultiple)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events_mbr.events_mbr_iccp_groups.events_mbr_iccp_group.events_mbr_bundle_descendant_iccp_group.bmd_event_bag_multiple")
}

func init() { proto.RegisterFile("bmd_event_bag_multiple.proto", fileDescriptor_2eb75e524eb4a446) }

var fileDescriptor_2eb75e524eb4a446 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x56, 0x3d, 0x6f, 0xd4, 0x4c,
	0x10, 0xd6, 0xde, 0x47, 0xde, 0xd7, 0x63, 0x9d, 0x22, 0x56, 0x7c, 0x58, 0x0a, 0x11, 0x91, 0x2b,
	0x44, 0xe1, 0xe2, 0x68, 0x29, 0x89, 0x50, 0x84, 0x44, 0xe1, 0xd0, 0x50, 0xad, 0xfc, 0xb1, 0xb1,
	0x2c, 0x79, 0x77, 0xad, 0xf5, 0x1e, 0x22, 0x25, 0x7f, 0x80, 0x8a, 0x36, 0x12, 0x42, 0x82, 0xbf,
	0x80, 0xa8, 0xe8, 0x68, 0xe9, 0xf9, 0x05, 0xfc, 0x0c, 0x34, 0xb3, 0x3e, 0x73, 0xc1, 0x47, 0x7f,
	0xa7, 0x54, 0x77, 0xfb, 0xec, 0xcc, 0xdc, 0xf3, 0x3c, 0xbe, 0x99, 0x31, 0xdc, 0xcf, 0x55, 0x29,
	0xe4, 0x6b, 0xa9, 0x9d, 0xc8, 0xb3, 0x4a, 0xa8, 0x55, 0xe3, 0xea, 0xb6, 0x91, 0x49, 0x6b, 0x8d,
	0x33, 0xfc, 0x1d, 0x2b, 0xea, 0xae, 0x30, 0xa2, 0x36, 0x9d, 0x78, 0x63, 0x45, 0xbe, 0xd2, 0x65,
	0x23, 0x55, 0x65, 0x85, 0x69, 0xa5, 0x4d, 0xfc, 0x51, 0xd4, 0xfa, 0xc2, 0x58, 0x95, 0xb9, 0xda,
	0xe8, 0x84, 0x2a, 0x75, 0x42, 0xe5, 0x76, 0xe3, 0xab, 0xa8, 0x8b, 0xa2, 0x15, 0x95, 0x35, 0xab,
	0xb6, 0xdb, 0x0e, 0x6f, 0xa2, 0x7d, 0xd5, 0x52, 0x76, 0x85, 0xd4, 0x65, 0xa6, 0xdd, 0x46, 0x5c,
	0xfc, 0x04, 0x8e, 0xb6, 0x13, 0x16, 0xcf, 0x4f, 0x5f, 0x9d, 0xf3, 0x63, 0x80, 0x3f, 0xc1, 0x11,
	0x3b, 0x61, 0x0f, 0x17, 0x69, 0x80, 0xc8, 0x33, 0xca, 0x7e, 0xcb, 0x80, 0x63, 0x3a, 0x26, 0xfa,
	0x12, 0x65, 0xe6, 0x32, 0x7e, 0x04, 0x01, 0x7e, 0x0a, 0x77, 0xd9, 0x4a, 0x4a, 0x0a, 0xd2, 0xff,
	0x11, 0x78, 0x79, 0xd9, 0x4a, 0x7e, 0x0f, 0xfe, 0xd3, 0x86, 0xe2, 0xa2, 0x09, 0xd5, 0x3b, 0xd0,
	0xe6, 0x29, 0x66, 0xdd, 0x86, 0xb9, 0xb4, 0xd6, 0xd8, 0x68, 0x4a, 0xb0, 0x3f, 0xf0, 0x07, 0x10,
	0x76, 0xce, 0xd6, 0xba, 0xf2, 0x29, 0x33, 0xaa, 0x06, 0x1e, 0xc2, 0xb4, 0xf8, 0x6a, 0x02, 0x77,
	0xaf, 0x73, 0xa8, 0x9d, 0x54, 0xa8, 0x9c, 0x3f, 0x82, 0x5b, 0x4a, 0xaa, 0x5c, 0xda, 0xfe, 0x62,
	0x83, 0xcf, 0xa1, 0xbf, 0x38, 0x45, 0x9c, 0x68, 0x1d, 0x03, 0xb8, 0x5a, 0x49, 0xd1, 0xb9, 0x4c,
	0xb5, 0xc4, 0x6c, 0x96, 0x06, 0x88, 0x9c, 0x23, 0xc0, 0xbf, 0x30, 0x98, 0x11, 0x01, 0x24, 0x17,
	0x2e, 0x3f, 0xb2, 0x64, 0xb7, 0x1e, 0x64, 0x32, 0x7e, 0x0c, 0x29, 0x11, 0xfe, 0x97, 0x3f, 0x79,
	0xd9, 0xa0, 0x3f, 0x7d, 0xcd, 0xb1, 0x3f, 0xfe, 0xe2, 0x06, 0xf8, 0xf3, 0x7e, 0x02, 0x77, 0xb6,
	0xf8, 0x63, 0x2b, 0x1e, 0xc3, 0xc2, 0x56, 0x63, 0x6b, 0x42, 0x5b, 0xdd, 0x00, 0x5b, 0x3e, 0xcc,
	0xff, 0x6e, 0x6d, 0xb4, 0x05, 0xf5, 0x8e, 0x0c, 0x09, 0xe4, 0x60, 0xc7, 0x0f, 0x06, 0x87, 0x43,
	0xcb, 0x39, 0x52, 0x42, 0xa6, 0x84, 0xcb, 0xcf, 0x3b, 0x2e, 0x7d, 0x3d, 0x34, 0xd2, 0xc5, 0x7a,
	0x32, 0xb8, 0x33, 0x7d, 0x61, 0x48, 0xd1, 0xd0, 0x24, 0xbd, 0xa2, 0xe9, 0xfe, 0x28, 0xca, 0xcb,
	0x26, 0x5d, 0xac, 0x7b, 0xd9, 0x2b, 0xfa, 0xce, 0x20, 0xa4, 0xff, 0x75, 0xaf, 0x66, 0x46, 0x6a,
	0x3e, 0xed, 0x83, 0x1a, 0x5b, 0xa5, 0x01, 0x76, 0x1f, 0x09, 0x89, 0x7f, 0x31, 0x58, 0x5c, 0x5b,
	0x5e, 0xb8, 0x78, 0x28, 0x4e, 0x67, 0x6a, 0x58, 0x3c, 0x08, 0xbc, 0xc8, 0x94, 0xe4, 0x5f, 0x19,
	0xcc, 0xf1, 0xd0, 0x45, 0x93, 0x93, 0xe9, 0x1e, 0x34, 0x23, 0x72, 0x4d, 0x3d, 0xe3, 0xf8, 0x27,
	0xf3, 0x43, 0x7c, 0xbc, 0xa7, 0xf9, 0x37, 0x06, 0x61, 0xff, 0x03, 0x18, 0x1b, 0x2d, 0x49, 0xdc,
	0xd5, 0x4e, 0x8a, 0x1b, 0xe8, 0xa7, 0x7e, 0x8a, 0x74, 0x67, 0x4e, 0xaa, 0xfc, 0x80, 0xde, 0x8d,
	0x1e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x42, 0x4b, 0xf7, 0x25, 0x3b, 0x09, 0x00, 0x00,
}
