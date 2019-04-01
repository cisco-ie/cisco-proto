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
// source: islm_summary_data.proto

package cisco_ios_xr_iedge4710_oper_iedge_license_manager_nodes_node_iedge_license_manager_summary

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

type IslmSummaryData_KEYS struct {
	Nodeid               string   `protobuf:"bytes,1,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IslmSummaryData_KEYS) Reset()         { *m = IslmSummaryData_KEYS{} }
func (m *IslmSummaryData_KEYS) String() string { return proto.CompactTextString(m) }
func (*IslmSummaryData_KEYS) ProtoMessage()    {}
func (*IslmSummaryData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1667d12f050f149e, []int{0}
}

func (m *IslmSummaryData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IslmSummaryData_KEYS.Unmarshal(m, b)
}
func (m *IslmSummaryData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IslmSummaryData_KEYS.Marshal(b, m, deterministic)
}
func (m *IslmSummaryData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IslmSummaryData_KEYS.Merge(m, src)
}
func (m *IslmSummaryData_KEYS) XXX_Size() int {
	return xxx_messageInfo_IslmSummaryData_KEYS.Size(m)
}
func (m *IslmSummaryData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IslmSummaryData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IslmSummaryData_KEYS proto.InternalMessageInfo

func (m *IslmSummaryData_KEYS) GetNodeid() string {
	if m != nil {
		return m.Nodeid
	}
	return ""
}

type IslmSummaryData struct {
	SessionLimit         uint32   `protobuf:"varint,50,opt,name=session_limit,json=sessionLimit,proto3" json:"session_limit,omitempty"`
	SessionThreshold     uint32   `protobuf:"varint,51,opt,name=session_threshold,json=sessionThreshold,proto3" json:"session_threshold,omitempty"`
	SessionLicenseCount  uint32   `protobuf:"varint,52,opt,name=session_license_count,json=sessionLicenseCount,proto3" json:"session_license_count,omitempty"`
	SessionCount         uint32   `protobuf:"varint,53,opt,name=session_count,json=sessionCount,proto3" json:"session_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IslmSummaryData) Reset()         { *m = IslmSummaryData{} }
func (m *IslmSummaryData) String() string { return proto.CompactTextString(m) }
func (*IslmSummaryData) ProtoMessage()    {}
func (*IslmSummaryData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1667d12f050f149e, []int{1}
}

func (m *IslmSummaryData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IslmSummaryData.Unmarshal(m, b)
}
func (m *IslmSummaryData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IslmSummaryData.Marshal(b, m, deterministic)
}
func (m *IslmSummaryData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IslmSummaryData.Merge(m, src)
}
func (m *IslmSummaryData) XXX_Size() int {
	return xxx_messageInfo_IslmSummaryData.Size(m)
}
func (m *IslmSummaryData) XXX_DiscardUnknown() {
	xxx_messageInfo_IslmSummaryData.DiscardUnknown(m)
}

var xxx_messageInfo_IslmSummaryData proto.InternalMessageInfo

func (m *IslmSummaryData) GetSessionLimit() uint32 {
	if m != nil {
		return m.SessionLimit
	}
	return 0
}

func (m *IslmSummaryData) GetSessionThreshold() uint32 {
	if m != nil {
		return m.SessionThreshold
	}
	return 0
}

func (m *IslmSummaryData) GetSessionLicenseCount() uint32 {
	if m != nil {
		return m.SessionLicenseCount
	}
	return 0
}

func (m *IslmSummaryData) GetSessionCount() uint32 {
	if m != nil {
		return m.SessionCount
	}
	return 0
}

func init() {
	proto.RegisterType((*IslmSummaryData_KEYS)(nil), "cisco_ios_xr_iedge4710_oper.iedge_license_manager.nodes.node.iedge_license_manager_summary.islm_summary_data_KEYS")
	proto.RegisterType((*IslmSummaryData)(nil), "cisco_ios_xr_iedge4710_oper.iedge_license_manager.nodes.node.iedge_license_manager_summary.islm_summary_data")
}

func init() { proto.RegisterFile("islm_summary_data.proto", fileDescriptor_1667d12f050f149e) }

var fileDescriptor_1667d12f050f149e = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xd9, 0x4b, 0xc1, 0x60, 0xc1, 0x46, 0xac, 0x7b, 0x2c, 0xf5, 0x52, 0x10, 0x96, 0xda,
	0x56, 0xfc, 0x00, 0xe2, 0x49, 0x4f, 0xd5, 0x8b, 0x5e, 0x86, 0xb8, 0x19, 0xda, 0x81, 0x24, 0x53,
	0x32, 0x29, 0xe8, 0x97, 0xf3, 0xb3, 0xc9, 0xa6, 0xbb, 0xfe, 0x61, 0x7b, 0x09, 0xcc, 0x7b, 0xbf,
	0x37, 0xf3, 0x88, 0xba, 0x24, 0x71, 0x1e, 0x64, 0xef, 0xbd, 0x89, 0x9f, 0x60, 0x4d, 0x32, 0xd5,
	0x2e, 0x72, 0x62, 0xfd, 0x56, 0x93, 0xd4, 0x0c, 0xc4, 0x02, 0x1f, 0x11, 0x08, 0xed, 0x06, 0x57,
	0x77, 0x37, 0x73, 0xe0, 0x1d, 0xc6, 0x2a, 0x8f, 0xe0, 0xa8, 0xc6, 0x20, 0x08, 0xde, 0x04, 0xb3,
	0xc1, 0x58, 0x05, 0xb6, 0x28, 0xf9, 0x3d, 0x0e, 0x74, 0x67, 0xa6, 0x73, 0x35, 0xee, 0x9d, 0x85,
	0xc7, 0x87, 0xd7, 0x67, 0x3d, 0x56, 0x83, 0x26, 0x4f, 0xb6, 0x2c, 0x26, 0xc5, 0xec, 0x64, 0xdd,
	0x4e, 0xd3, 0xaf, 0x42, 0x8d, 0x7a, 0x11, 0x7d, 0xa5, 0x86, 0x82, 0x22, 0xc4, 0x01, 0x1c, 0x79,
	0x4a, 0xe5, 0x62, 0x52, 0xcc, 0x86, 0xeb, 0xd3, 0x56, 0x7c, 0x6a, 0x34, 0x7d, 0xad, 0x46, 0x1d,
	0x94, 0xb6, 0x11, 0x65, 0xcb, 0xce, 0x96, 0xcb, 0x0c, 0x9e, 0xb5, 0xc6, 0x4b, 0xa7, 0xeb, 0x85,
	0xba, 0xf8, 0xdd, 0x78, 0x28, 0x5f, 0xf3, 0x3e, 0xa4, 0x72, 0x95, 0x03, 0xe7, 0x3f, 0x9b, 0xb3,
	0x77, 0xdf, 0x58, 0x7f, 0x5b, 0x1c, 0xd8, 0xdb, 0x7f, 0x2d, 0x32, 0xf4, 0x3e, 0xc8, 0xbf, 0xba,
	0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xfb, 0xec, 0x7f, 0x70, 0x01, 0x00, 0x00,
}
