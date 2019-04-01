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
// source: mpls_lsd_lbl_sum.proto

package cisco_ios_xr_mpls_lsd_oper_mpls_lsd_nodes_mpls_lsd_node_label_summary_vrfs_label_summary_vrf

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

type MplsLsdLblSum_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdLblSum_KEYS) Reset()         { *m = MplsLsdLblSum_KEYS{} }
func (m *MplsLsdLblSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLsdLblSum_KEYS) ProtoMessage()    {}
func (*MplsLsdLblSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c07ef7bacd841e, []int{0}
}

func (m *MplsLsdLblSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdLblSum_KEYS.Unmarshal(m, b)
}
func (m *MplsLsdLblSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdLblSum_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLsdLblSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdLblSum_KEYS.Merge(m, src)
}
func (m *MplsLsdLblSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLsdLblSum_KEYS.Size(m)
}
func (m *MplsLsdLblSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdLblSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdLblSum_KEYS proto.InternalMessageInfo

func (m *MplsLsdLblSum_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MplsLsdLblSum_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type MplsLsdLblSumElem struct {
	ApplicationType        string   `protobuf:"bytes,1,opt,name=application_type,json=applicationType,proto3" json:"application_type,omitempty"`
	ApplicationName        string   `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	ApplicationRolePrimary int32    `protobuf:"zigzag32,3,opt,name=application_role_primary,json=applicationRolePrimary,proto3" json:"application_role_primary,omitempty"`
	ApplicationInstance    string   `protobuf:"bytes,4,opt,name=application_instance,json=applicationInstance,proto3" json:"application_instance,omitempty"`
	NumberOfLabels         uint32   `protobuf:"varint,5,opt,name=number_of_labels,json=numberOfLabels,proto3" json:"number_of_labels,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *MplsLsdLblSumElem) Reset()         { *m = MplsLsdLblSumElem{} }
func (m *MplsLsdLblSumElem) String() string { return proto.CompactTextString(m) }
func (*MplsLsdLblSumElem) ProtoMessage()    {}
func (*MplsLsdLblSumElem) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c07ef7bacd841e, []int{1}
}

func (m *MplsLsdLblSumElem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdLblSumElem.Unmarshal(m, b)
}
func (m *MplsLsdLblSumElem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdLblSumElem.Marshal(b, m, deterministic)
}
func (m *MplsLsdLblSumElem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdLblSumElem.Merge(m, src)
}
func (m *MplsLsdLblSumElem) XXX_Size() int {
	return xxx_messageInfo_MplsLsdLblSumElem.Size(m)
}
func (m *MplsLsdLblSumElem) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdLblSumElem.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdLblSumElem proto.InternalMessageInfo

func (m *MplsLsdLblSumElem) GetApplicationType() string {
	if m != nil {
		return m.ApplicationType
	}
	return ""
}

func (m *MplsLsdLblSumElem) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *MplsLsdLblSumElem) GetApplicationRolePrimary() int32 {
	if m != nil {
		return m.ApplicationRolePrimary
	}
	return 0
}

func (m *MplsLsdLblSumElem) GetApplicationInstance() string {
	if m != nil {
		return m.ApplicationInstance
	}
	return ""
}

func (m *MplsLsdLblSumElem) GetNumberOfLabels() uint32 {
	if m != nil {
		return m.NumberOfLabels
	}
	return 0
}

type MplsLsdLblSum struct {
	VrfNameXr            string               `protobuf:"bytes,50,opt,name=vrf_name_xr,json=vrfNameXr,proto3" json:"vrf_name_xr,omitempty"`
	OwnerCount           []*MplsLsdLblSumElem `protobuf:"bytes,51,rep,name=owner_count,json=ownerCount,proto3" json:"owner_count,omitempty"`
	TotalLabels          uint32               `protobuf:"varint,52,opt,name=total_labels,json=totalLabels,proto3" json:"total_labels,omitempty"`
	RsiConnected         int32                `protobuf:"zigzag32,53,opt,name=rsi_connected,json=rsiConnected,proto3" json:"rsi_connected,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MplsLsdLblSum) Reset()         { *m = MplsLsdLblSum{} }
func (m *MplsLsdLblSum) String() string { return proto.CompactTextString(m) }
func (*MplsLsdLblSum) ProtoMessage()    {}
func (*MplsLsdLblSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c07ef7bacd841e, []int{2}
}

func (m *MplsLsdLblSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdLblSum.Unmarshal(m, b)
}
func (m *MplsLsdLblSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdLblSum.Marshal(b, m, deterministic)
}
func (m *MplsLsdLblSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdLblSum.Merge(m, src)
}
func (m *MplsLsdLblSum) XXX_Size() int {
	return xxx_messageInfo_MplsLsdLblSum.Size(m)
}
func (m *MplsLsdLblSum) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdLblSum.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdLblSum proto.InternalMessageInfo

func (m *MplsLsdLblSum) GetVrfNameXr() string {
	if m != nil {
		return m.VrfNameXr
	}
	return ""
}

func (m *MplsLsdLblSum) GetOwnerCount() []*MplsLsdLblSumElem {
	if m != nil {
		return m.OwnerCount
	}
	return nil
}

func (m *MplsLsdLblSum) GetTotalLabels() uint32 {
	if m != nil {
		return m.TotalLabels
	}
	return 0
}

func (m *MplsLsdLblSum) GetRsiConnected() int32 {
	if m != nil {
		return m.RsiConnected
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsLsdLblSum_KEYS)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.label_summary_vrfs.label_summary_vrf.mpls_lsd_lbl_sum_KEYS")
	proto.RegisterType((*MplsLsdLblSumElem)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.label_summary_vrfs.label_summary_vrf.mpls_lsd_lbl_sum_elem")
	proto.RegisterType((*MplsLsdLblSum)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.label_summary_vrfs.label_summary_vrf.mpls_lsd_lbl_sum")
}

func init() { proto.RegisterFile("mpls_lsd_lbl_sum.proto", fileDescriptor_54c07ef7bacd841e) }

var fileDescriptor_54c07ef7bacd841e = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcf, 0x6b, 0xdb, 0x30,
	0x18, 0xc5, 0xc9, 0x7e, 0x24, 0x72, 0xb2, 0x65, 0xda, 0x16, 0x34, 0x06, 0xc3, 0xcb, 0x2e, 0xde,
	0xc5, 0xb0, 0x64, 0x83, 0xdd, 0xc3, 0x0e, 0xa3, 0xa5, 0x29, 0x6e, 0x0f, 0x2d, 0x14, 0x3e, 0x64,
	0x47, 0x06, 0x81, 0x2c, 0x09, 0x49, 0x49, 0x93, 0x7b, 0xff, 0x81, 0x42, 0xff, 0xe0, 0x62, 0x39,
	0x0e, 0x6e, 0x9a, 0x6b, 0x8f, 0x7a, 0xef, 0xf3, 0x7b, 0xdf, 0x7b, 0xfe, 0xd0, 0xb8, 0xd4, 0xc2,
	0x82, 0xb0, 0x4b, 0x10, 0x99, 0x00, 0xbb, 0x2a, 0x13, 0x6d, 0x94, 0x53, 0xf8, 0x26, 0xe7, 0x36,
	0x57, 0xc0, 0x95, 0x85, 0x8d, 0x81, 0xfd, 0x90, 0xd2, 0xcc, 0x24, 0xfb, 0x97, 0x54, 0x4b, 0x66,
	0x9f, 0x3e, 0x13, 0x41, 0x33, 0xe6, 0x85, 0x4a, 0x6a, 0xb6, 0xb0, 0x36, 0x85, 0x7d, 0x0e, 0x4d,
	0x16, 0xe8, 0xf3, 0xa1, 0x2f, 0x9c, 0xfc, 0xbb, 0xbe, 0xc0, 0x5f, 0x51, 0xbf, 0x52, 0x01, 0x49,
	0x4b, 0x46, 0x82, 0x28, 0x88, 0xfb, 0x69, 0xaf, 0x02, 0xce, 0x68, 0xc9, 0xf0, 0x17, 0xd4, 0x5b,
	0x9b, 0xa2, 0xe6, 0x3a, 0x9e, 0x7b, 0xbb, 0x36, 0x45, 0x45, 0x4d, 0xee, 0x3a, 0x47, 0x14, 0x99,
	0x60, 0x25, 0xfe, 0x89, 0x46, 0x54, 0x6b, 0xc1, 0x73, 0xea, 0xb8, 0x92, 0xe0, 0xb6, 0xba, 0x11,
	0x7e, 0xdf, 0xc2, 0x2f, 0xb7, 0x9a, 0x1d, 0x8e, 0xb6, 0x7c, 0xda, 0xa3, 0x7e, 0x95, 0xbf, 0x88,
	0xb4, 0x47, 0x8d, 0x12, 0x0c, 0xb4, 0xe1, 0x55, 0x40, 0xd2, 0x8d, 0x82, 0xf8, 0x43, 0x3a, 0x6e,
	0xf1, 0xa9, 0x12, 0xec, 0xbc, 0x66, 0xf1, 0x2f, 0xf4, 0xa9, 0xfd, 0x25, 0x97, 0xd6, 0x51, 0x99,
	0x33, 0xf2, 0xca, 0x1b, 0x7d, 0x6c, 0x71, 0xff, 0x77, 0x14, 0x8e, 0xd1, 0x48, 0xae, 0xca, 0x8c,
	0x19, 0x50, 0x05, 0xf8, 0x32, 0x2d, 0x79, 0x1d, 0x05, 0xf1, 0x30, 0x7d, 0x57, 0xe3, 0x8b, 0xe2,
	0xd4, 0xa3, 0x93, 0xfb, 0x0e, 0x1a, 0x1d, 0xd6, 0x80, 0xbf, 0xa1, 0xb0, 0xa9, 0x0d, 0x36, 0x86,
	0x4c, 0xbd, 0x51, 0x7f, 0xd7, 0xdc, 0x95, 0xc1, 0x0f, 0x01, 0x0a, 0xd5, 0xad, 0x64, 0x06, 0x72,
	0xb5, 0x92, 0x8e, 0xcc, 0xa2, 0x6e, 0x1c, 0x4e, 0x6d, 0xf2, 0x92, 0x17, 0x90, 0x1c, 0xfd, 0x59,
	0x29, 0xf2, 0x7b, 0xcc, 0xab, 0x35, 0xf0, 0x77, 0x34, 0x70, 0xca, 0x51, 0xd1, 0x24, 0xfe, 0xed,
	0x13, 0x87, 0x1e, 0xab, 0xe3, 0xe2, 0x1f, 0x68, 0x68, 0x2c, 0x87, 0x5c, 0x49, 0xc9, 0x72, 0xc7,
	0x96, 0xe4, 0x8f, 0xaf, 0x7e, 0x60, 0x2c, 0x9f, 0x37, 0x58, 0xf6, 0xc6, 0x1f, 0xf4, 0xec, 0x31,
	0x00, 0x00, 0xff, 0xff, 0xac, 0x25, 0x35, 0xb9, 0xea, 0x02, 0x00, 0x00,
}
