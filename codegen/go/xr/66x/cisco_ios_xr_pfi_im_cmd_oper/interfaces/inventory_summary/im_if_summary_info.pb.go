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
// source: im_if_summary_info.proto

package cisco_ios_xr_pfi_im_cmd_oper_interfaces_inventory_summary

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

type ImIfSummaryInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImIfSummaryInfo_KEYS) Reset()         { *m = ImIfSummaryInfo_KEYS{} }
func (m *ImIfSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*ImIfSummaryInfo_KEYS) ProtoMessage()    {}
func (*ImIfSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f1e7e171affb549, []int{0}
}

func (m *ImIfSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImIfSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *ImIfSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImIfSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *ImIfSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImIfSummaryInfo_KEYS.Merge(m, src)
}
func (m *ImIfSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_ImIfSummaryInfo_KEYS.Size(m)
}
func (m *ImIfSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ImIfSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ImIfSummaryInfo_KEYS proto.InternalMessageInfo

type ImIfGroupCountsSt struct {
	InterfaceCount          uint32   `protobuf:"varint,1,opt,name=interface_count,json=interfaceCount,proto3" json:"interface_count,omitempty"`
	UpInterfaceCount        uint32   `protobuf:"varint,2,opt,name=up_interface_count,json=upInterfaceCount,proto3" json:"up_interface_count,omitempty"`
	DownInterfaceCount      uint32   `protobuf:"varint,3,opt,name=down_interface_count,json=downInterfaceCount,proto3" json:"down_interface_count,omitempty"`
	AdminDownInterfaceCount uint32   `protobuf:"varint,4,opt,name=admin_down_interface_count,json=adminDownInterfaceCount,proto3" json:"admin_down_interface_count,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ImIfGroupCountsSt) Reset()         { *m = ImIfGroupCountsSt{} }
func (m *ImIfGroupCountsSt) String() string { return proto.CompactTextString(m) }
func (*ImIfGroupCountsSt) ProtoMessage()    {}
func (*ImIfGroupCountsSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f1e7e171affb549, []int{1}
}

func (m *ImIfGroupCountsSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImIfGroupCountsSt.Unmarshal(m, b)
}
func (m *ImIfGroupCountsSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImIfGroupCountsSt.Marshal(b, m, deterministic)
}
func (m *ImIfGroupCountsSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImIfGroupCountsSt.Merge(m, src)
}
func (m *ImIfGroupCountsSt) XXX_Size() int {
	return xxx_messageInfo_ImIfGroupCountsSt.Size(m)
}
func (m *ImIfGroupCountsSt) XXX_DiscardUnknown() {
	xxx_messageInfo_ImIfGroupCountsSt.DiscardUnknown(m)
}

var xxx_messageInfo_ImIfGroupCountsSt proto.InternalMessageInfo

func (m *ImIfGroupCountsSt) GetInterfaceCount() uint32 {
	if m != nil {
		return m.InterfaceCount
	}
	return 0
}

func (m *ImIfGroupCountsSt) GetUpInterfaceCount() uint32 {
	if m != nil {
		return m.UpInterfaceCount
	}
	return 0
}

func (m *ImIfGroupCountsSt) GetDownInterfaceCount() uint32 {
	if m != nil {
		return m.DownInterfaceCount
	}
	return 0
}

func (m *ImIfGroupCountsSt) GetAdminDownInterfaceCount() uint32 {
	if m != nil {
		return m.AdminDownInterfaceCount
	}
	return 0
}

type ImIfTypeSummarySt struct {
	InterfaceTypeName        string             `protobuf:"bytes,1,opt,name=interface_type_name,json=interfaceTypeName,proto3" json:"interface_type_name,omitempty"`
	InterfaceTypeDescription string             `protobuf:"bytes,2,opt,name=interface_type_description,json=interfaceTypeDescription,proto3" json:"interface_type_description,omitempty"`
	InterfaceCounts          *ImIfGroupCountsSt `protobuf:"bytes,3,opt,name=interface_counts,json=interfaceCounts,proto3" json:"interface_counts,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}           `json:"-"`
	XXX_unrecognized         []byte             `json:"-"`
	XXX_sizecache            int32              `json:"-"`
}

func (m *ImIfTypeSummarySt) Reset()         { *m = ImIfTypeSummarySt{} }
func (m *ImIfTypeSummarySt) String() string { return proto.CompactTextString(m) }
func (*ImIfTypeSummarySt) ProtoMessage()    {}
func (*ImIfTypeSummarySt) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f1e7e171affb549, []int{2}
}

func (m *ImIfTypeSummarySt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImIfTypeSummarySt.Unmarshal(m, b)
}
func (m *ImIfTypeSummarySt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImIfTypeSummarySt.Marshal(b, m, deterministic)
}
func (m *ImIfTypeSummarySt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImIfTypeSummarySt.Merge(m, src)
}
func (m *ImIfTypeSummarySt) XXX_Size() int {
	return xxx_messageInfo_ImIfTypeSummarySt.Size(m)
}
func (m *ImIfTypeSummarySt) XXX_DiscardUnknown() {
	xxx_messageInfo_ImIfTypeSummarySt.DiscardUnknown(m)
}

var xxx_messageInfo_ImIfTypeSummarySt proto.InternalMessageInfo

func (m *ImIfTypeSummarySt) GetInterfaceTypeName() string {
	if m != nil {
		return m.InterfaceTypeName
	}
	return ""
}

func (m *ImIfTypeSummarySt) GetInterfaceTypeDescription() string {
	if m != nil {
		return m.InterfaceTypeDescription
	}
	return ""
}

func (m *ImIfTypeSummarySt) GetInterfaceCounts() *ImIfGroupCountsSt {
	if m != nil {
		return m.InterfaceCounts
	}
	return nil
}

type ImIfSummaryInfo struct {
	InterfaceType        []*ImIfTypeSummarySt `protobuf:"bytes,50,rep,name=interface_type,json=interfaceType,proto3" json:"interface_type,omitempty"`
	InterfaceCounts      *ImIfGroupCountsSt   `protobuf:"bytes,51,opt,name=interface_counts,json=interfaceCounts,proto3" json:"interface_counts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ImIfSummaryInfo) Reset()         { *m = ImIfSummaryInfo{} }
func (m *ImIfSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*ImIfSummaryInfo) ProtoMessage()    {}
func (*ImIfSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f1e7e171affb549, []int{3}
}

func (m *ImIfSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImIfSummaryInfo.Unmarshal(m, b)
}
func (m *ImIfSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImIfSummaryInfo.Marshal(b, m, deterministic)
}
func (m *ImIfSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImIfSummaryInfo.Merge(m, src)
}
func (m *ImIfSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_ImIfSummaryInfo.Size(m)
}
func (m *ImIfSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ImIfSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ImIfSummaryInfo proto.InternalMessageInfo

func (m *ImIfSummaryInfo) GetInterfaceType() []*ImIfTypeSummarySt {
	if m != nil {
		return m.InterfaceType
	}
	return nil
}

func (m *ImIfSummaryInfo) GetInterfaceCounts() *ImIfGroupCountsSt {
	if m != nil {
		return m.InterfaceCounts
	}
	return nil
}

func init() {
	proto.RegisterType((*ImIfSummaryInfo_KEYS)(nil), "cisco_ios_xr_pfi_im_cmd_oper.interfaces.inventory_summary.im_if_summary_info_KEYS")
	proto.RegisterType((*ImIfGroupCountsSt)(nil), "cisco_ios_xr_pfi_im_cmd_oper.interfaces.inventory_summary.im_if_group_counts_st")
	proto.RegisterType((*ImIfTypeSummarySt)(nil), "cisco_ios_xr_pfi_im_cmd_oper.interfaces.inventory_summary.im_if_type_summary_st")
	proto.RegisterType((*ImIfSummaryInfo)(nil), "cisco_ios_xr_pfi_im_cmd_oper.interfaces.inventory_summary.im_if_summary_info")
}

func init() { proto.RegisterFile("im_if_summary_info.proto", fileDescriptor_1f1e7e171affb549) }

var fileDescriptor_1f1e7e171affb549 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xbd, 0x4b, 0xc3, 0x40,
	0x14, 0xa7, 0xad, 0x08, 0x3d, 0xe9, 0x87, 0xa7, 0xd2, 0xd8, 0xa9, 0x64, 0xd1, 0x41, 0x82, 0xb4,
	0x93, 0xe8, 0x66, 0x1d, 0x44, 0x10, 0x89, 0x2e, 0x4e, 0x8f, 0x78, 0xb9, 0xc8, 0x1b, 0xee, 0x83,
	0xbb, 0x8b, 0x35, 0xb8, 0x39, 0xfa, 0x3f, 0xfa, 0xbf, 0x48, 0xaf, 0x34, 0x25, 0x49, 0x27, 0x11,
	0xd7, 0xf7, 0xfb, 0xe0, 0xfd, 0x7e, 0xef, 0x8e, 0x04, 0x28, 0x00, 0x33, 0xb0, 0xb9, 0x10, 0x89,
	0x29, 0x00, 0x65, 0xa6, 0x22, 0x6d, 0x94, 0x53, 0xf4, 0x82, 0xa1, 0x65, 0x0a, 0x50, 0x59, 0x78,
	0x37, 0xa0, 0x33, 0x04, 0x14, 0xc0, 0x44, 0x0a, 0x4a, 0x73, 0x13, 0xa1, 0x74, 0xdc, 0x64, 0x09,
	0xe3, 0x36, 0x42, 0xf9, 0xc6, 0xa5, 0x53, 0xa6, 0x58, 0xbb, 0x84, 0xc7, 0x64, 0xd4, 0xb4, 0x85,
	0xbb, 0x9b, 0xe7, 0xc7, 0xf0, 0xbb, 0x45, 0x8e, 0x56, 0xd8, 0xab, 0x51, 0xb9, 0x06, 0xa6, 0x72,
	0xe9, 0x2c, 0x58, 0x47, 0x4f, 0xc8, 0xa0, 0x34, 0x5d, 0x8d, 0x83, 0xd6, 0xa4, 0x75, 0xda, 0x8b,
	0xfb, 0xe5, 0xf8, 0x7a, 0x39, 0xa5, 0x67, 0x84, 0xe6, 0x1a, 0xea, 0xdc, 0xb6, 0xe7, 0x0e, 0x73,
	0x7d, 0x5b, 0x65, 0x9f, 0x93, 0xc3, 0x54, 0x2d, 0x64, 0x83, 0xdf, 0xf1, 0x7c, 0xba, 0xc4, 0x6a,
	0x8a, 0x4b, 0x32, 0x4e, 0x52, 0x81, 0x12, 0xb6, 0xea, 0x76, 0xbc, 0x6e, 0xe4, 0x19, 0xf3, 0x86,
	0x38, 0xfc, 0x6c, 0xaf, 0xf3, 0xb9, 0x42, 0xf3, 0xb2, 0x00, 0xeb, 0x68, 0x44, 0x0e, 0x36, 0x5e,
	0x1e, 0x94, 0x89, 0xe0, 0x3e, 0x63, 0x37, 0xde, 0x2f, 0xa1, 0xa7, 0x42, 0xf3, 0xfb, 0x44, 0x70,
	0x7a, 0x45, 0xc6, 0x35, 0x7e, 0xca, 0x2d, 0x33, 0xa8, 0x1d, 0x2a, 0xe9, 0xe3, 0x76, 0xe3, 0xa0,
	0x22, 0x9b, 0x6f, 0x70, 0xfa, 0x41, 0x86, 0xb5, 0xcd, 0xad, 0x8f, 0xbc, 0x37, 0x7d, 0x88, 0x7e,
	0x7d, 0xd8, 0x68, 0xeb, 0xe5, 0xe2, 0x41, 0xf5, 0x40, 0x36, 0xfc, 0x6a, 0x13, 0xda, 0x7c, 0x00,
	0x74, 0x41, 0xfa, 0xd5, 0x44, 0xc1, 0x74, 0xd2, 0xf9, 0x93, 0x8d, 0x6a, 0x5d, 0xc7, 0xbd, 0x4a,
	0x2f, 0x5b, 0xcb, 0x98, 0xfd, 0x53, 0x19, 0x2f, 0xbb, 0xfe, 0x3b, 0xcd, 0x7e, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x04, 0x6b, 0x7b, 0x19, 0x6a, 0x03, 0x00, 0x00,
}
