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
// source: fpd_info_list.proto

package cisco_ios_xr_show_fpd_loc_ng_oper_show_fpd_locations_location_fpd

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

type FpdInfoList_KEYS struct {
	LocationName         string   `protobuf:"bytes,1,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	FpdName              string   `protobuf:"bytes,2,opt,name=fpd_name,json=fpdName,proto3" json:"fpd_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FpdInfoList_KEYS) Reset()         { *m = FpdInfoList_KEYS{} }
func (m *FpdInfoList_KEYS) String() string { return proto.CompactTextString(m) }
func (*FpdInfoList_KEYS) ProtoMessage()    {}
func (*FpdInfoList_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b1254036d90f634, []int{0}
}

func (m *FpdInfoList_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FpdInfoList_KEYS.Unmarshal(m, b)
}
func (m *FpdInfoList_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FpdInfoList_KEYS.Marshal(b, m, deterministic)
}
func (m *FpdInfoList_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FpdInfoList_KEYS.Merge(m, src)
}
func (m *FpdInfoList_KEYS) XXX_Size() int {
	return xxx_messageInfo_FpdInfoList_KEYS.Size(m)
}
func (m *FpdInfoList_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FpdInfoList_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FpdInfoList_KEYS proto.InternalMessageInfo

func (m *FpdInfoList_KEYS) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *FpdInfoList_KEYS) GetFpdName() string {
	if m != nil {
		return m.FpdName
	}
	return ""
}

type FpdInfo struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	CardName             string   `protobuf:"bytes,2,opt,name=card_name,json=cardName,proto3" json:"card_name,omitempty"`
	FpdName              string   `protobuf:"bytes,3,opt,name=fpd_name,json=fpdName,proto3" json:"fpd_name,omitempty"`
	HwVersion            string   `protobuf:"bytes,4,opt,name=hw_version,json=hwVersion,proto3" json:"hw_version,omitempty"`
	SecureBootAttr       string   `protobuf:"bytes,5,opt,name=secure_boot_attr,json=secureBootAttr,proto3" json:"secure_boot_attr,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	RunningVersion       string   `protobuf:"bytes,7,opt,name=running_version,json=runningVersion,proto3" json:"running_version,omitempty"`
	ProgramdVersion      string   `protobuf:"bytes,8,opt,name=programd_version,json=programdVersion,proto3" json:"programd_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FpdInfo) Reset()         { *m = FpdInfo{} }
func (m *FpdInfo) String() string { return proto.CompactTextString(m) }
func (*FpdInfo) ProtoMessage()    {}
func (*FpdInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b1254036d90f634, []int{1}
}

func (m *FpdInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FpdInfo.Unmarshal(m, b)
}
func (m *FpdInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FpdInfo.Marshal(b, m, deterministic)
}
func (m *FpdInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FpdInfo.Merge(m, src)
}
func (m *FpdInfo) XXX_Size() int {
	return xxx_messageInfo_FpdInfo.Size(m)
}
func (m *FpdInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FpdInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FpdInfo proto.InternalMessageInfo

func (m *FpdInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *FpdInfo) GetCardName() string {
	if m != nil {
		return m.CardName
	}
	return ""
}

func (m *FpdInfo) GetFpdName() string {
	if m != nil {
		return m.FpdName
	}
	return ""
}

func (m *FpdInfo) GetHwVersion() string {
	if m != nil {
		return m.HwVersion
	}
	return ""
}

func (m *FpdInfo) GetSecureBootAttr() string {
	if m != nil {
		return m.SecureBootAttr
	}
	return ""
}

func (m *FpdInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FpdInfo) GetRunningVersion() string {
	if m != nil {
		return m.RunningVersion
	}
	return ""
}

func (m *FpdInfo) GetProgramdVersion() string {
	if m != nil {
		return m.ProgramdVersion
	}
	return ""
}

type FpdInfoList struct {
	FpdInfoDetaile       []*FpdInfo `protobuf:"bytes,50,rep,name=fpd_info_detaile,json=fpdInfoDetaile,proto3" json:"fpd_info_detaile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FpdInfoList) Reset()         { *m = FpdInfoList{} }
func (m *FpdInfoList) String() string { return proto.CompactTextString(m) }
func (*FpdInfoList) ProtoMessage()    {}
func (*FpdInfoList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b1254036d90f634, []int{2}
}

func (m *FpdInfoList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FpdInfoList.Unmarshal(m, b)
}
func (m *FpdInfoList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FpdInfoList.Marshal(b, m, deterministic)
}
func (m *FpdInfoList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FpdInfoList.Merge(m, src)
}
func (m *FpdInfoList) XXX_Size() int {
	return xxx_messageInfo_FpdInfoList.Size(m)
}
func (m *FpdInfoList) XXX_DiscardUnknown() {
	xxx_messageInfo_FpdInfoList.DiscardUnknown(m)
}

var xxx_messageInfo_FpdInfoList proto.InternalMessageInfo

func (m *FpdInfoList) GetFpdInfoDetaile() []*FpdInfo {
	if m != nil {
		return m.FpdInfoDetaile
	}
	return nil
}

func init() {
	proto.RegisterType((*FpdInfoList_KEYS)(nil), "cisco_ios_xr_show_fpd_loc_ng_oper.show_fpd.locations.location.fpd.fpd_info_list_KEYS")
	proto.RegisterType((*FpdInfo)(nil), "cisco_ios_xr_show_fpd_loc_ng_oper.show_fpd.locations.location.fpd.fpd_info")
	proto.RegisterType((*FpdInfoList)(nil), "cisco_ios_xr_show_fpd_loc_ng_oper.show_fpd.locations.location.fpd.fpd_info_list")
}

func init() { proto.RegisterFile("fpd_info_list.proto", fileDescriptor_5b1254036d90f634) }

var fileDescriptor_5b1254036d90f634 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xd9, 0xa6, 0x5b, 0xf7, 0xea, 0xfe, 0x10, 0x41, 0xaa, 0x22, 0x8c, 0x79, 0x70, 0x5e,
	0x7a, 0x98, 0x9f, 0x60, 0xa2, 0x07, 0x19, 0x78, 0x98, 0x22, 0x78, 0x0a, 0x59, 0x9b, 0x6e, 0x81,
	0x2e, 0x6f, 0x48, 0x52, 0xeb, 0x27, 0xf0, 0x0b, 0xf8, 0x85, 0xa5, 0x69, 0x53, 0x28, 0x1e, 0xbd,
	0xe5, 0xfd, 0x3d, 0xcf, 0xfb, 0x3c, 0x90, 0x04, 0xce, 0x52, 0x95, 0x50, 0x21, 0x53, 0xa4, 0x99,
	0x30, 0x36, 0x52, 0x1a, 0x2d, 0x92, 0x55, 0x2c, 0x4c, 0x8c, 0x54, 0xa0, 0xa1, 0x5f, 0x9a, 0x9a,
	0x3d, 0x16, 0xb4, 0xb4, 0x65, 0x18, 0x53, 0xb9, 0xa3, 0xa8, 0xb8, 0x8e, 0x3c, 0x8c, 0x32, 0x8c,
	0x99, 0x15, 0x28, 0x4d, 0x73, 0x8a, 0x52, 0x95, 0xcc, 0xdf, 0x80, 0xb4, 0x92, 0xe9, 0xfa, 0xe9,
	0xe3, 0x95, 0xdc, 0xc0, 0xc8, 0xbb, 0xa8, 0x64, 0x07, 0x1e, 0x76, 0x66, 0x9d, 0xc5, 0x70, 0x73,
	0xea, 0xe1, 0x0b, 0x3b, 0x70, 0x72, 0x01, 0x41, 0xb9, 0xea, 0xf4, 0xae, 0xd3, 0x07, 0xa9, 0x4a,
	0x4a, 0x69, 0xfe, 0xd3, 0xad, 0xb4, 0x32, 0x96, 0x5c, 0x42, 0xe0, 0xf7, 0xea, 0x9c, 0x66, 0x26,
	0x57, 0x30, 0x8c, 0x99, 0x6e, 0x85, 0x04, 0x25, 0xf8, 0x53, 0xd0, 0x6b, 0x15, 0x90, 0x6b, 0x80,
	0x7d, 0x41, 0x3f, 0xb9, 0x36, 0x65, 0xea, 0x91, 0x13, 0x87, 0xfb, 0xe2, 0xbd, 0x02, 0x64, 0x01,
	0x53, 0xc3, 0xe3, 0x5c, 0x73, 0xba, 0x45, 0xb4, 0x94, 0x59, 0xab, 0xc3, 0x63, 0x67, 0x1a, 0x57,
	0xfc, 0x01, 0xd1, 0xae, 0xac, 0xd5, 0xe4, 0x1c, 0xfa, 0xc6, 0x32, 0x9b, 0x9b, 0xb0, 0xef, 0xf4,
	0x7a, 0x22, 0xb7, 0x30, 0xd1, 0xb9, 0x94, 0x42, 0xee, 0x9a, 0x96, 0x41, 0x15, 0x50, 0x63, 0x5f,
	0x75, 0x07, 0x53, 0xa5, 0x71, 0xa7, 0xd9, 0x21, 0x69, 0x9c, 0x81, 0x73, 0x4e, 0x3c, 0xaf, 0xad,
	0xf3, 0xef, 0x0e, 0x8c, 0x5a, 0x97, 0x4d, 0x72, 0x98, 0x36, 0x20, 0xe1, 0x96, 0x89, 0x8c, 0x87,
	0xcb, 0x59, 0x6f, 0x71, 0xb2, 0x5c, 0x47, 0xff, 0x7e, 0xdb, 0xc8, 0x47, 0x6f, 0xc6, 0xa9, 0x4a,
	0x9e, 0x65, 0x8a, 0x8f, 0x55, 0xc5, 0xb6, 0xef, 0xbe, 0xcf, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0x60, 0x2a, 0x63, 0x55, 0x02, 0x00, 0x00,
}
