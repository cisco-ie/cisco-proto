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

package cisco_ios_xr_bundlemgr_oper_bundle_information_events_events_bundles_events_bundle_events_bundle_ancestor

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
	proto.RegisterType((*BmdEventBag_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_ancestor.bmd_event_bag_KEYS")
	proto.RegisterType((*BmdBagEventData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_ancestor.bmd_bag_event_data")
	proto.RegisterType((*BmdBagEventItemMbr)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_ancestor.bmd_bag_event_item_mbr")
	proto.RegisterType((*BmdBagEventItemBdl)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_ancestor.bmd_bag_event_item_bdl")
	proto.RegisterType((*BmdBagEventItemRg)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_ancestor.bmd_bag_event_item_rg")
	proto.RegisterType((*BmdBagEventItem)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_ancestor.bmd_bag_event_item")
	proto.RegisterType((*BmdEventBag)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.events.events_bundles.events_bundle.events_bundle_ancestor.bmd_event_bag")
}

func init() { proto.RegisterFile("bmd_event_bag.proto", fileDescriptor_22002626393f0e45) }

var fileDescriptor_22002626393f0e45 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x95, 0xbf, 0x8e, 0xd4, 0x30,
	0x10, 0xc6, 0x65, 0x36, 0x77, 0x90, 0x89, 0xa2, 0x05, 0xf3, 0x2f, 0x12, 0x42, 0xac, 0x52, 0x1d,
	0x14, 0x29, 0xf6, 0x1e, 0x80, 0x86, 0x2b, 0x4e, 0x48, 0x14, 0x39, 0x1a, 0x2a, 0xcb, 0xd9, 0x78,
	0xa3, 0x48, 0x6b, 0x3b, 0x9a, 0x18, 0xc4, 0x15, 0x48, 0xdc, 0x23, 0x20, 0x9e, 0x80, 0x96, 0x8e,
	0x57, 0xe0, 0x81, 0x68, 0x69, 0xd1, 0xd8, 0xc9, 0x8a, 0xbd, 0x40, 0xcb, 0xb2, 0x95, 0x35, 0x9f,
	0x3d, 0x93, 0xdf, 0x97, 0xcc, 0x38, 0x70, 0xb7, 0xd2, 0xb5, 0x50, 0xef, 0x94, 0x71, 0xa2, 0x92,
	0x4d, 0xd1, 0xa1, 0x75, 0x96, 0xb7, 0xab, 0xb6, 0x5f, 0x59, 0xd1, 0xda, 0x5e, 0xbc, 0x47, 0x51,
	0xbd, 0x35, 0xf5, 0x46, 0xe9, 0x06, 0x85, 0xed, 0x14, 0x16, 0x21, 0x14, 0xad, 0x59, 0x5b, 0xd4,
	0xd2, 0xb5, 0xd6, 0x14, 0x3e, 0xbf, 0x1f, 0x96, 0xe1, 0xfc, 0xb5, 0x70, 0x37, 0x12, 0xd2, 0xac,
	0x54, 0xef, 0x2c, 0xe6, 0xcf, 0x81, 0xef, 0x10, 0x88, 0x97, 0x67, 0x6f, 0x2e, 0xf8, 0x53, 0xb8,
	0xbd, 0x7d, 0x8c, 0x53, 0xb8, 0x96, 0x2b, 0x95, 0xb1, 0x05, 0x3b, 0x89, 0xcb, 0x79, 0xd0, 0xcf,
	0x47, 0x39, 0xbf, 0x62, 0xa1, 0x02, 0xe5, 0x86, 0x2a, 0xb5, 0x74, 0x92, 0x3f, 0x82, 0x98, 0x56,
	0xe1, 0x2e, 0xbb, 0x31, 0xf5, 0x16, 0x09, 0xaf, 0x2f, 0x3b, 0xc5, 0x1f, 0xc2, 0x4d, 0x63, 0xfd,
	0xb9, 0xec, 0xc6, 0x82, 0x9d, 0xa4, 0xe5, 0xb1, 0xb1, 0x2f, 0x28, 0xeb, 0x1e, 0x1c, 0x29, 0x44,
	0x8b, 0xd9, 0xcc, 0xcb, 0x21, 0xe0, 0x4f, 0x20, 0xe9, 0x1d, 0xb6, 0xa6, 0x09, 0x29, 0x91, 0xaf,
	0x06, 0x41, 0xa2, 0xb4, 0xfc, 0x27, 0x83, 0x07, 0xbb, 0x0c, 0xad, 0x53, 0x5a, 0xe8, 0x0a, 0xf9,
	0x33, 0xb8, 0xa3, 0x95, 0xae, 0x14, 0x0e, 0x1b, 0xbf, 0xf1, 0xcc, 0xc3, 0xc6, 0x19, 0xe9, 0x1e,
	0xeb, 0x31, 0x80, 0x6b, 0xb5, 0x12, 0xbd, 0x93, 0xba, 0xf3, 0x64, 0x51, 0x19, 0x93, 0x72, 0x41,
	0x02, 0xff, 0xc4, 0x20, 0xf2, 0x00, 0x04, 0x97, 0x2c, 0x3f, 0x14, 0xff, 0xec, 0x2b, 0x15, 0xd3,
	0x17, 0x5c, 0x7a, 0x94, 0xbf, 0x39, 0xaf, 0xea, 0x0d, 0x39, 0x1f, 0xca, 0x4c, 0x9d, 0x87, 0x8d,
	0x83, 0x76, 0xfe, 0x83, 0xc1, 0xfd, 0x3f, 0x38, 0xc7, 0x86, 0xe7, 0x90, 0x62, 0x33, 0x35, 0x9d,
	0x60, 0x73, 0xd0, 0x86, 0xbf, 0x45, 0xd7, 0x07, 0x8d, 0x0c, 0x93, 0x93, 0x89, 0xd5, 0x58, 0x6d,
	0x8d, 0x7e, 0x65, 0x30, 0xdf, 0x0e, 0x80, 0xf3, 0x90, 0xde, 0x6e, 0xb2, 0xbc, 0x62, 0x7b, 0x73,
	0x35, 0x4e, 0x67, 0x99, 0x8e, 0x23, 0xe8, 0xce, 0xcd, 0xda, 0x7a, 0xd8, 0x6d, 0xcf, 0x0e, 0xb0,
	0xb3, 0xff, 0x02, 0xb6, 0xaa, 0x37, 0x65, 0x3a, 0x4e, 0x4d, 0x80, 0xfd, 0xc2, 0x20, 0xf1, 0x7d,
	0x36, 0x80, 0x46, 0x1e, 0xf4, 0xe3, 0x9e, 0x41, 0xb1, 0x29, 0x63, 0x6a, 0x74, 0xcf, 0x98, 0x7f,
	0x67, 0x90, 0xee, 0x5c, 0xef, 0x74, 0x2f, 0xfb, 0x73, 0x46, 0x6a, 0x95, 0x2d, 0xc3, 0xbd, 0x4c,
	0xc2, 0x2b, 0xa9, 0x15, 0xff, 0xcc, 0xe0, 0x88, 0x82, 0x3e, 0x3b, 0x5d, 0xcc, 0xf6, 0xda, 0xf7,
	0x44, 0x51, 0x06, 0x96, 0xea, 0xd8, 0xff, 0x14, 0x4f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4a,
	0xa7, 0x87, 0xd1, 0x2b, 0x07, 0x00, 0x00,
}
