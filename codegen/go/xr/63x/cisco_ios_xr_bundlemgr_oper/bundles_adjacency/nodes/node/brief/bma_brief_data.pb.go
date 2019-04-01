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
// source: bma_brief_data.proto

package cisco_ios_xr_bundlemgr_oper_bundles_adjacency_nodes_node_brief

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

type BmaBriefData_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmaBriefData_KEYS) Reset()         { *m = BmaBriefData_KEYS{} }
func (m *BmaBriefData_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmaBriefData_KEYS) ProtoMessage()    {}
func (*BmaBriefData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1289f3712efc355, []int{0}
}

func (m *BmaBriefData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmaBriefData_KEYS.Unmarshal(m, b)
}
func (m *BmaBriefData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmaBriefData_KEYS.Marshal(b, m, deterministic)
}
func (m *BmaBriefData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmaBriefData_KEYS.Merge(m, src)
}
func (m *BmaBriefData_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmaBriefData_KEYS.Size(m)
}
func (m *BmaBriefData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmaBriefData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmaBriefData_KEYS proto.InternalMessageInfo

func (m *BmaBriefData_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type BmaLoadBalanceData struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                uint32   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	LocalLinkThreshold   uint32   `protobuf:"varint,3,opt,name=local_link_threshold,json=localLinkThreshold,proto3" json:"local_link_threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmaLoadBalanceData) Reset()         { *m = BmaLoadBalanceData{} }
func (m *BmaLoadBalanceData) String() string { return proto.CompactTextString(m) }
func (*BmaLoadBalanceData) ProtoMessage()    {}
func (*BmaLoadBalanceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1289f3712efc355, []int{1}
}

func (m *BmaLoadBalanceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmaLoadBalanceData.Unmarshal(m, b)
}
func (m *BmaLoadBalanceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmaLoadBalanceData.Marshal(b, m, deterministic)
}
func (m *BmaLoadBalanceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmaLoadBalanceData.Merge(m, src)
}
func (m *BmaLoadBalanceData) XXX_Size() int {
	return xxx_messageInfo_BmaLoadBalanceData.Size(m)
}
func (m *BmaLoadBalanceData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmaLoadBalanceData.DiscardUnknown(m)
}

var xxx_messageInfo_BmaLoadBalanceData proto.InternalMessageInfo

func (m *BmaLoadBalanceData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *BmaLoadBalanceData) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *BmaLoadBalanceData) GetLocalLinkThreshold() uint32 {
	if m != nil {
		return m.LocalLinkThreshold
	}
	return 0
}

type BundleAdjacencySubInterface struct {
	InterfaceName        string              `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	LoadBalanceData      *BmaLoadBalanceData `protobuf:"bytes,2,opt,name=load_balance_data,json=loadBalanceData,proto3" json:"load_balance_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BundleAdjacencySubInterface) Reset()         { *m = BundleAdjacencySubInterface{} }
func (m *BundleAdjacencySubInterface) String() string { return proto.CompactTextString(m) }
func (*BundleAdjacencySubInterface) ProtoMessage()    {}
func (*BundleAdjacencySubInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1289f3712efc355, []int{2}
}

func (m *BundleAdjacencySubInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleAdjacencySubInterface.Unmarshal(m, b)
}
func (m *BundleAdjacencySubInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleAdjacencySubInterface.Marshal(b, m, deterministic)
}
func (m *BundleAdjacencySubInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleAdjacencySubInterface.Merge(m, src)
}
func (m *BundleAdjacencySubInterface) XXX_Size() int {
	return xxx_messageInfo_BundleAdjacencySubInterface.Size(m)
}
func (m *BundleAdjacencySubInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleAdjacencySubInterface.DiscardUnknown(m)
}

var xxx_messageInfo_BundleAdjacencySubInterface proto.InternalMessageInfo

func (m *BundleAdjacencySubInterface) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *BundleAdjacencySubInterface) GetLoadBalanceData() *BmaLoadBalanceData {
	if m != nil {
		return m.LoadBalanceData
	}
	return nil
}

type BundleAdjacencyBrief struct {
	InterfaceName        string                         `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	SubInterfaceCount    uint32                         `protobuf:"varint,2,opt,name=sub_interface_count,json=subInterfaceCount,proto3" json:"sub_interface_count,omitempty"`
	MemberCount          uint32                         `protobuf:"varint,3,opt,name=member_count,json=memberCount,proto3" json:"member_count,omitempty"`
	TotalWeight          uint32                         `protobuf:"varint,4,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty"`
	SubInterface         []*BundleAdjacencySubInterface `protobuf:"bytes,5,rep,name=sub_interface,json=subInterface,proto3" json:"sub_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *BundleAdjacencyBrief) Reset()         { *m = BundleAdjacencyBrief{} }
func (m *BundleAdjacencyBrief) String() string { return proto.CompactTextString(m) }
func (*BundleAdjacencyBrief) ProtoMessage()    {}
func (*BundleAdjacencyBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1289f3712efc355, []int{3}
}

func (m *BundleAdjacencyBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleAdjacencyBrief.Unmarshal(m, b)
}
func (m *BundleAdjacencyBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleAdjacencyBrief.Marshal(b, m, deterministic)
}
func (m *BundleAdjacencyBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleAdjacencyBrief.Merge(m, src)
}
func (m *BundleAdjacencyBrief) XXX_Size() int {
	return xxx_messageInfo_BundleAdjacencyBrief.Size(m)
}
func (m *BundleAdjacencyBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleAdjacencyBrief.DiscardUnknown(m)
}

var xxx_messageInfo_BundleAdjacencyBrief proto.InternalMessageInfo

func (m *BundleAdjacencyBrief) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *BundleAdjacencyBrief) GetSubInterfaceCount() uint32 {
	if m != nil {
		return m.SubInterfaceCount
	}
	return 0
}

func (m *BundleAdjacencyBrief) GetMemberCount() uint32 {
	if m != nil {
		return m.MemberCount
	}
	return 0
}

func (m *BundleAdjacencyBrief) GetTotalWeight() uint32 {
	if m != nil {
		return m.TotalWeight
	}
	return 0
}

func (m *BundleAdjacencyBrief) GetSubInterface() []*BundleAdjacencySubInterface {
	if m != nil {
		return m.SubInterface
	}
	return nil
}

type BmaBriefData struct {
	BundleData           []*BundleAdjacencyBrief `protobuf:"bytes,50,rep,name=bundle_data,json=bundleData,proto3" json:"bundle_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *BmaBriefData) Reset()         { *m = BmaBriefData{} }
func (m *BmaBriefData) String() string { return proto.CompactTextString(m) }
func (*BmaBriefData) ProtoMessage()    {}
func (*BmaBriefData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1289f3712efc355, []int{4}
}

func (m *BmaBriefData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmaBriefData.Unmarshal(m, b)
}
func (m *BmaBriefData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmaBriefData.Marshal(b, m, deterministic)
}
func (m *BmaBriefData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmaBriefData.Merge(m, src)
}
func (m *BmaBriefData) XXX_Size() int {
	return xxx_messageInfo_BmaBriefData.Size(m)
}
func (m *BmaBriefData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmaBriefData.DiscardUnknown(m)
}

var xxx_messageInfo_BmaBriefData proto.InternalMessageInfo

func (m *BmaBriefData) GetBundleData() []*BundleAdjacencyBrief {
	if m != nil {
		return m.BundleData
	}
	return nil
}

func init() {
	proto.RegisterType((*BmaBriefData_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundles_adjacency.nodes.node.brief.bma_brief_data_KEYS")
	proto.RegisterType((*BmaLoadBalanceData)(nil), "cisco_ios_xr_bundlemgr_oper.bundles_adjacency.nodes.node.brief.bma_load_balance_data")
	proto.RegisterType((*BundleAdjacencySubInterface)(nil), "cisco_ios_xr_bundlemgr_oper.bundles_adjacency.nodes.node.brief.BundleAdjacencySubInterface")
	proto.RegisterType((*BundleAdjacencyBrief)(nil), "cisco_ios_xr_bundlemgr_oper.bundles_adjacency.nodes.node.brief.BundleAdjacencyBrief")
	proto.RegisterType((*BmaBriefData)(nil), "cisco_ios_xr_bundlemgr_oper.bundles_adjacency.nodes.node.brief.bma_brief_data")
}

func init() { proto.RegisterFile("bma_brief_data.proto", fileDescriptor_b1289f3712efc355) }

var fileDescriptor_b1289f3712efc355 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xf6, 0x03, 0xb1, 0x93, 0xed, 0xa2, 0xf5, 0x06, 0x29, 0xd2, 0x5e, 0x4a, 0x24, 0xa4,
	0x9e, 0x22, 0x14, 0xee, 0x48, 0x14, 0x38, 0x20, 0x10, 0x87, 0xec, 0x22, 0x84, 0x38, 0x58, 0x63,
	0xc7, 0xdb, 0x86, 0x3a, 0x76, 0x65, 0x3b, 0x40, 0x6f, 0x70, 0xe2, 0x1f, 0xf1, 0x1b, 0xf8, 0x59,
	0xc8, 0x76, 0x53, 0x5a, 0x40, 0x08, 0xa9, 0x5c, 0xa2, 0xcc, 0x9b, 0x37, 0xf6, 0x7b, 0x6f, 0x12,
	0xc8, 0x58, 0x87, 0x94, 0x99, 0x56, 0xdc, 0xd0, 0x06, 0x1d, 0x96, 0x4b, 0xa3, 0x9d, 0x26, 0x8f,
	0x78, 0x6b, 0xb9, 0xa6, 0xad, 0xb6, 0xf4, 0x93, 0xa1, 0xac, 0x57, 0x8d, 0x14, 0xdd, 0xcc, 0x50,
	0xbd, 0x14, 0xa6, 0x8c, 0xa5, 0xa5, 0xd8, 0xbc, 0x47, 0x2e, 0x14, 0x5f, 0x95, 0x4a, 0x37, 0xc2,
	0x86, 0x67, 0x19, 0x8e, 0x2a, 0x2a, 0xb8, 0xd8, 0x3d, 0x97, 0xbe, 0x78, 0xf6, 0xf6, 0x8a, 0x5c,
	0xc2, 0x89, 0x27, 0x51, 0x85, 0x9d, 0xc8, 0x93, 0x71, 0x32, 0x39, 0xa9, 0x6f, 0x7b, 0xe0, 0x15,
	0x76, 0xa2, 0xb0, 0x70, 0xd7, 0xcf, 0x48, 0x8d, 0x0d, 0x65, 0x28, 0x51, 0x71, 0x11, 0x46, 0x09,
	0x81, 0x23, 0xb7, 0x5a, 0x0e, 0x03, 0xe1, 0x9d, 0x64, 0x70, 0xfc, 0x01, 0x65, 0x2f, 0xf2, 0x83,
	0x71, 0x32, 0x19, 0xd5, 0xb1, 0x20, 0x0f, 0x20, 0x93, 0x9a, 0xa3, 0xa4, 0xb2, 0x55, 0x0b, 0xea,
	0xe6, 0x46, 0xd8, 0xb9, 0x96, 0x4d, 0x7e, 0x18, 0x48, 0x24, 0xf4, 0x5e, 0xb6, 0x6a, 0x71, 0x3d,
	0x74, 0x8a, 0xef, 0x09, 0x5c, 0x4e, 0x83, 0x9f, 0xc7, 0x83, 0x9b, 0xab, 0x9e, 0x3d, 0x57, 0x4e,
	0x98, 0x1b, 0xe4, 0x82, 0xdc, 0x87, 0xb3, 0x76, 0x28, 0xb6, 0x65, 0x8f, 0x36, 0xa8, 0xd7, 0x4e,
	0xbe, 0x24, 0x70, 0xfe, 0x9b, 0xf0, 0xa0, 0x2d, 0xad, 0x5e, 0x97, 0xfb, 0x85, 0x59, 0xfe, 0x31,
	0x95, 0xfa, 0x8e, 0x87, 0xa6, 0x11, 0x79, 0x8a, 0x0e, 0x8b, 0x6f, 0x07, 0x90, 0xfd, 0x62, 0x65,
	0xea, 0xe7, 0xff, 0xd5, 0x43, 0x09, 0x17, 0xb6, 0x67, 0xf4, 0x27, 0x95, 0xeb, 0x5e, 0xb9, 0x75,
	0xc0, 0xe7, 0x76, 0x2b, 0x95, 0x27, 0xbe, 0x41, 0xee, 0xc1, 0x69, 0x27, 0x3a, 0x26, 0xcc, 0x9a,
	0x18, 0x43, 0x4e, 0x23, 0xb6, 0xa1, 0x38, 0xed, 0x50, 0xd2, 0x8f, 0xa2, 0x9d, 0xcd, 0x5d, 0x7e,
	0x14, 0x29, 0x01, 0x7b, 0x13, 0x20, 0xf2, 0x39, 0x81, 0xd1, 0xce, 0xb5, 0xf9, 0xf1, 0xf8, 0x70,
	0x92, 0x56, 0xef, 0xf6, 0x4d, 0xed, 0x2f, 0x5b, 0xad, 0x4f, 0xb7, 0xdd, 0x14, 0x5f, 0x13, 0x38,
	0xdb, 0xfd, 0x5a, 0x49, 0x0f, 0x69, 0xbc, 0x22, 0x2e, 0xb2, 0x0a, 0x92, 0xae, 0xff, 0xb3, 0xa4,
	0xb0, 0x9d, 0x1a, 0xe2, 0xa0, 0x5f, 0x21, 0xbb, 0x15, 0xfe, 0xbe, 0x87, 0x3f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x83, 0xf6, 0xb4, 0xab, 0x95, 0x03, 0x00, 0x00,
}