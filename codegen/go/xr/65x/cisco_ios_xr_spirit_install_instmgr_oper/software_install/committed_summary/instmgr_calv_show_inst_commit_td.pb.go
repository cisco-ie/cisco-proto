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
// source: instmgr_calv_show_inst_commit_td.proto

package cisco_ios_xr_spirit_install_instmgr_oper_software_install_committed_summary

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

type InstmgrCalvShowInstCommitTd_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstmgrCalvShowInstCommitTd_KEYS) Reset()         { *m = InstmgrCalvShowInstCommitTd_KEYS{} }
func (m *InstmgrCalvShowInstCommitTd_KEYS) String() string { return proto.CompactTextString(m) }
func (*InstmgrCalvShowInstCommitTd_KEYS) ProtoMessage()    {}
func (*InstmgrCalvShowInstCommitTd_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9dc69a01bf98677, []int{0}
}

func (m *InstmgrCalvShowInstCommitTd_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS.Unmarshal(m, b)
}
func (m *InstmgrCalvShowInstCommitTd_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS.Marshal(b, m, deterministic)
}
func (m *InstmgrCalvShowInstCommitTd_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS.Merge(m, src)
}
func (m *InstmgrCalvShowInstCommitTd_KEYS) XXX_Size() int {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS.Size(m)
}
func (m *InstmgrCalvShowInstCommitTd_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS proto.InternalMessageInfo

type InstmgrCalvShowInstCommitRow struct {
	ErrorMessage              string   `protobuf:"bytes,1,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Location                  string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	NodeType                  string   `protobuf:"bytes,3,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	BootPartitionName         string   `protobuf:"bytes,4,opt,name=boot_partition_name,json=bootPartitionName,proto3" json:"boot_partition_name,omitempty"`
	NumberOfCommittedPackages uint32   `protobuf:"varint,5,opt,name=number_of_committed_packages,json=numberOfCommittedPackages,proto3" json:"number_of_committed_packages,omitempty"`
	CommittedPackages         string   `protobuf:"bytes,6,opt,name=committed_packages,json=committedPackages,proto3" json:"committed_packages,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *InstmgrCalvShowInstCommitRow) Reset()         { *m = InstmgrCalvShowInstCommitRow{} }
func (m *InstmgrCalvShowInstCommitRow) String() string { return proto.CompactTextString(m) }
func (*InstmgrCalvShowInstCommitRow) ProtoMessage()    {}
func (*InstmgrCalvShowInstCommitRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9dc69a01bf98677, []int{1}
}

func (m *InstmgrCalvShowInstCommitRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrCalvShowInstCommitRow.Unmarshal(m, b)
}
func (m *InstmgrCalvShowInstCommitRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrCalvShowInstCommitRow.Marshal(b, m, deterministic)
}
func (m *InstmgrCalvShowInstCommitRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrCalvShowInstCommitRow.Merge(m, src)
}
func (m *InstmgrCalvShowInstCommitRow) XXX_Size() int {
	return xxx_messageInfo_InstmgrCalvShowInstCommitRow.Size(m)
}
func (m *InstmgrCalvShowInstCommitRow) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrCalvShowInstCommitRow.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrCalvShowInstCommitRow proto.InternalMessageInfo

func (m *InstmgrCalvShowInstCommitRow) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *InstmgrCalvShowInstCommitRow) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *InstmgrCalvShowInstCommitRow) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func (m *InstmgrCalvShowInstCommitRow) GetBootPartitionName() string {
	if m != nil {
		return m.BootPartitionName
	}
	return ""
}

func (m *InstmgrCalvShowInstCommitRow) GetNumberOfCommittedPackages() uint32 {
	if m != nil {
		return m.NumberOfCommittedPackages
	}
	return 0
}

func (m *InstmgrCalvShowInstCommitRow) GetCommittedPackages() string {
	if m != nil {
		return m.CommittedPackages
	}
	return ""
}

type InstmgrCalvShowInstCommitTd struct {
	CommittedPackageInfo []*InstmgrCalvShowInstCommitRow `protobuf:"bytes,50,rep,name=committed_package_info,json=committedPackageInfo,proto3" json:"committed_package_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *InstmgrCalvShowInstCommitTd) Reset()         { *m = InstmgrCalvShowInstCommitTd{} }
func (m *InstmgrCalvShowInstCommitTd) String() string { return proto.CompactTextString(m) }
func (*InstmgrCalvShowInstCommitTd) ProtoMessage()    {}
func (*InstmgrCalvShowInstCommitTd) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9dc69a01bf98677, []int{2}
}

func (m *InstmgrCalvShowInstCommitTd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd.Unmarshal(m, b)
}
func (m *InstmgrCalvShowInstCommitTd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd.Marshal(b, m, deterministic)
}
func (m *InstmgrCalvShowInstCommitTd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrCalvShowInstCommitTd.Merge(m, src)
}
func (m *InstmgrCalvShowInstCommitTd) XXX_Size() int {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd.Size(m)
}
func (m *InstmgrCalvShowInstCommitTd) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrCalvShowInstCommitTd.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrCalvShowInstCommitTd proto.InternalMessageInfo

func (m *InstmgrCalvShowInstCommitTd) GetCommittedPackageInfo() []*InstmgrCalvShowInstCommitRow {
	if m != nil {
		return m.CommittedPackageInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*InstmgrCalvShowInstCommitTd_KEYS)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.committed_summary.instmgr_calv_show_inst_commit_td_KEYS")
	proto.RegisterType((*InstmgrCalvShowInstCommitRow)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.committed_summary.instmgr_calv_show_inst_commit_row")
	proto.RegisterType((*InstmgrCalvShowInstCommitTd)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.committed_summary.instmgr_calv_show_inst_commit_td")
}

func init() {
	proto.RegisterFile("instmgr_calv_show_inst_commit_td.proto", fileDescriptor_a9dc69a01bf98677)
}

var fileDescriptor_a9dc69a01bf98677 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbf, 0x4e, 0x23, 0x31,
	0x10, 0x87, 0xb5, 0xc9, 0x5d, 0x74, 0xf1, 0x5d, 0x8a, 0x33, 0x08, 0x2d, 0x7f, 0x8a, 0x25, 0x08,
	0x48, 0xc3, 0x16, 0xe1, 0x01, 0x28, 0x10, 0x05, 0x8a, 0x80, 0x28, 0xd0, 0x50, 0x8d, 0x1c, 0xc7,
	0x1b, 0x2c, 0xd6, 0x1e, 0x6b, 0xec, 0x10, 0xf2, 0x38, 0x48, 0xbc, 0x0c, 0x6f, 0x85, 0x76, 0x97,
	0x05, 0x29, 0x20, 0xd2, 0x50, 0xfa, 0xf7, 0x7d, 0x9e, 0xd1, 0x78, 0xcc, 0x0e, 0xb4, 0xf5, 0xc1,
	0x4c, 0x09, 0xa4, 0xc8, 0x1f, 0xc0, 0xdf, 0xe1, 0x1c, 0x8a, 0x04, 0x24, 0x1a, 0xa3, 0x03, 0x84,
	0x49, 0xea, 0x08, 0x03, 0xf2, 0x81, 0xd4, 0x5e, 0x22, 0x68, 0xf4, 0xf0, 0x48, 0xe0, 0x9d, 0x26,
	0x1d, 0x4a, 0x53, 0xe4, 0x39, 0xd4, 0x35, 0xd0, 0x29, 0x4a, 0x3d, 0x66, 0x61, 0x2e, 0x48, 0xd5,
	0x34, 0xad, 0x4a, 0x05, 0x35, 0x01, 0x3f, 0x33, 0x46, 0xd0, 0xa2, 0x7b, 0xc8, 0xf6, 0x57, 0xb5,
	0x85, 0xc1, 0xd9, 0xed, 0x75, 0xf7, 0xa9, 0xc1, 0x76, 0xbf, 0x37, 0x09, 0xe7, 0x7c, 0x8f, 0x75,
	0x14, 0x11, 0x12, 0x18, 0xe5, 0xbd, 0x98, 0xaa, 0x38, 0x4a, 0xa2, 0x5e, 0x7b, 0xf4, 0xaf, 0x0c,
	0x2f, 0xaa, 0x8c, 0x6f, 0xb1, 0x3f, 0x39, 0x4a, 0x11, 0x34, 0xda, 0xb8, 0x51, 0xf2, 0xf7, 0x33,
	0xdf, 0x66, 0x6d, 0x8b, 0x13, 0x05, 0x61, 0xe1, 0x54, 0xdc, 0xac, 0x60, 0x11, 0xdc, 0x2c, 0x9c,
	0xe2, 0x29, 0x5b, 0x1b, 0x23, 0x06, 0x70, 0x82, 0x82, 0x2e, 0x74, 0xb0, 0xc2, 0xa8, 0xf8, 0x57,
	0xa9, 0xfd, 0x2f, 0xd0, 0xb0, 0x26, 0x97, 0xc2, 0x28, 0x7e, 0xc2, 0x76, 0xec, 0xcc, 0x8c, 0x15,
	0x01, 0x66, 0xf0, 0x31, 0xbb, 0x13, 0xf2, 0x5e, 0x4c, 0x95, 0x8f, 0x7f, 0x27, 0x51, 0xaf, 0x33,
	0xda, 0xac, 0x9c, 0xab, 0xec, 0xb4, 0x36, 0x86, 0x6f, 0x02, 0x3f, 0x62, 0xfc, 0x8b, 0x6b, 0xad,
	0xaa, 0x9f, 0x5c, 0xd6, 0xbb, 0x2f, 0x11, 0x4b, 0x56, 0xbd, 0x26, 0x7f, 0x8e, 0xd8, 0xc6, 0xa7,
	0xa2, 0xa0, 0x6d, 0x86, 0x71, 0x3f, 0x69, 0xf6, 0xfe, 0xf6, 0x6d, 0xfa, 0x83, 0x0b, 0x4e, 0x57,
	0xee, 0x6c, 0xb4, 0xbe, 0x3c, 0xc8, 0xb9, 0xcd, 0x70, 0xdc, 0x2a, 0x3f, 0xdb, 0xf1, 0x6b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x16, 0x28, 0xd5, 0xe2, 0x96, 0x02, 0x00, 0x00,
}