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
// source: rpl_attachment_info.proto

package cisco_ios_xr_policy_repository_oper_routing_policy_sets_extended_community_seg_nh_sets_set_attached

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

type RplAttachmentInfo_KEYS struct {
	SetName              string   `protobuf:"bytes,1,opt,name=set_name,json=setName,proto3" json:"set_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RplAttachmentInfo_KEYS) Reset()         { *m = RplAttachmentInfo_KEYS{} }
func (m *RplAttachmentInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RplAttachmentInfo_KEYS) ProtoMessage()    {}
func (*RplAttachmentInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_067b47babaed8e3a, []int{0}
}

func (m *RplAttachmentInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RplAttachmentInfo_KEYS.Unmarshal(m, b)
}
func (m *RplAttachmentInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RplAttachmentInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RplAttachmentInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RplAttachmentInfo_KEYS.Merge(m, src)
}
func (m *RplAttachmentInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RplAttachmentInfo_KEYS.Size(m)
}
func (m *RplAttachmentInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RplAttachmentInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RplAttachmentInfo_KEYS proto.InternalMessageInfo

func (m *RplAttachmentInfo_KEYS) GetSetName() string {
	if m != nil {
		return m.SetName
	}
	return ""
}

type ClientInfo struct {
	Protocol                string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	VrfName                 string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	ProtoInstance           string   `protobuf:"bytes,3,opt,name=proto_instance,json=protoInstance,proto3" json:"proto_instance,omitempty"`
	AfName                  string   `protobuf:"bytes,4,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName                 string   `protobuf:"bytes,5,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	NeighborAddress         string   `protobuf:"bytes,6,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	NeighborAfName          string   `protobuf:"bytes,7,opt,name=neighbor_af_name,json=neighborAfName,proto3" json:"neighbor_af_name,omitempty"`
	GroupName               string   `protobuf:"bytes,8,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Direction               string   `protobuf:"bytes,9,opt,name=direction,proto3" json:"direction,omitempty"`
	Group                   string   `protobuf:"bytes,10,opt,name=group,proto3" json:"group,omitempty"`
	SourceProtocol          string   `protobuf:"bytes,11,opt,name=source_protocol,json=sourceProtocol,proto3" json:"source_protocol,omitempty"`
	AggregateNetworkAddress string   `protobuf:"bytes,12,opt,name=aggregate_network_address,json=aggregateNetworkAddress,proto3" json:"aggregate_network_address,omitempty"`
	InterfaceName           string   `protobuf:"bytes,13,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Instance                string   `protobuf:"bytes,14,opt,name=instance,proto3" json:"instance,omitempty"`
	AreaId                  string   `protobuf:"bytes,15,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	PropogateFrom           int32    `protobuf:"zigzag32,16,opt,name=propogate_from,json=propogateFrom,proto3" json:"propogate_from,omitempty"`
	PropogateTo             int32    `protobuf:"zigzag32,17,opt,name=propogate_to,json=propogateTo,proto3" json:"propogate_to,omitempty"`
	RoutePolicyName         string   `protobuf:"bytes,18,opt,name=route_policy_name,json=routePolicyName,proto3" json:"route_policy_name,omitempty"`
	AttachedPolicy          string   `protobuf:"bytes,19,opt,name=attached_policy,json=attachedPolicy,proto3" json:"attached_policy,omitempty"`
	AttachPoint             string   `protobuf:"bytes,20,opt,name=attach_point,json=attachPoint,proto3" json:"attach_point,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_067b47babaed8e3a, []int{1}
}

func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (m *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(m, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *ClientInfo) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *ClientInfo) GetProtoInstance() string {
	if m != nil {
		return m.ProtoInstance
	}
	return ""
}

func (m *ClientInfo) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *ClientInfo) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *ClientInfo) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *ClientInfo) GetNeighborAfName() string {
	if m != nil {
		return m.NeighborAfName
	}
	return ""
}

func (m *ClientInfo) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *ClientInfo) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *ClientInfo) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *ClientInfo) GetSourceProtocol() string {
	if m != nil {
		return m.SourceProtocol
	}
	return ""
}

func (m *ClientInfo) GetAggregateNetworkAddress() string {
	if m != nil {
		return m.AggregateNetworkAddress
	}
	return ""
}

func (m *ClientInfo) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *ClientInfo) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *ClientInfo) GetAreaId() string {
	if m != nil {
		return m.AreaId
	}
	return ""
}

func (m *ClientInfo) GetPropogateFrom() int32 {
	if m != nil {
		return m.PropogateFrom
	}
	return 0
}

func (m *ClientInfo) GetPropogateTo() int32 {
	if m != nil {
		return m.PropogateTo
	}
	return 0
}

func (m *ClientInfo) GetRoutePolicyName() string {
	if m != nil {
		return m.RoutePolicyName
	}
	return ""
}

func (m *ClientInfo) GetAttachedPolicy() string {
	if m != nil {
		return m.AttachedPolicy
	}
	return ""
}

func (m *ClientInfo) GetAttachPoint() string {
	if m != nil {
		return m.AttachPoint
	}
	return ""
}

type RplAttachmentInfo struct {
	Binding              []*ClientInfo `protobuf:"bytes,50,rep,name=binding,proto3" json:"binding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RplAttachmentInfo) Reset()         { *m = RplAttachmentInfo{} }
func (m *RplAttachmentInfo) String() string { return proto.CompactTextString(m) }
func (*RplAttachmentInfo) ProtoMessage()    {}
func (*RplAttachmentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_067b47babaed8e3a, []int{2}
}

func (m *RplAttachmentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RplAttachmentInfo.Unmarshal(m, b)
}
func (m *RplAttachmentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RplAttachmentInfo.Marshal(b, m, deterministic)
}
func (m *RplAttachmentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RplAttachmentInfo.Merge(m, src)
}
func (m *RplAttachmentInfo) XXX_Size() int {
	return xxx_messageInfo_RplAttachmentInfo.Size(m)
}
func (m *RplAttachmentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RplAttachmentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RplAttachmentInfo proto.InternalMessageInfo

func (m *RplAttachmentInfo) GetBinding() []*ClientInfo {
	if m != nil {
		return m.Binding
	}
	return nil
}

func init() {
	proto.RegisterType((*RplAttachmentInfo_KEYS)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.sets.extended_community_seg_nh.sets.set.attached.rpl_attachment_info_KEYS")
	proto.RegisterType((*ClientInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.sets.extended_community_seg_nh.sets.set.attached.client_info")
	proto.RegisterType((*RplAttachmentInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.sets.extended_community_seg_nh.sets.set.attached.rpl_attachment_info")
}

func init() { proto.RegisterFile("rpl_attachment_info.proto", fileDescriptor_067b47babaed8e3a) }

var fileDescriptor_067b47babaed8e3a = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0x15, 0x4a, 0xfe, 0x39, 0x6d, 0xd2, 0xb8, 0x95, 0xea, 0x54, 0x20, 0x85, 0x48, 0x55,
	0x03, 0x87, 0x3d, 0x14, 0x71, 0xe1, 0xc6, 0x01, 0xa4, 0x0a, 0xa9, 0x8a, 0x0a, 0x17, 0x4e, 0x96,
	0xe3, 0x9d, 0xdd, 0x58, 0x64, 0x3d, 0x2b, 0xdb, 0x29, 0xcd, 0x2b, 0xf0, 0x18, 0x3c, 0x0e, 0x4f,
	0x85, 0x76, 0xbc, 0xbb, 0xe4, 0xd0, 0x73, 0x8f, 0xfb, 0x7d, 0xbf, 0x1d, 0x7f, 0x9e, 0x19, 0xb3,
	0x99, 0x2b, 0xb7, 0x52, 0x85, 0xa0, 0xf4, 0xa6, 0x00, 0x1b, 0xa4, 0xb1, 0x19, 0x26, 0xa5, 0xc3,
	0x80, 0x5c, 0x6b, 0xe3, 0x35, 0x4a, 0x83, 0x5e, 0x3e, 0x3a, 0x59, 0xe2, 0xd6, 0xe8, 0xbd, 0x74,
	0x50, 0xa2, 0x37, 0x01, 0xdd, 0x5e, 0x62, 0x09, 0x2e, 0x71, 0xb8, 0x0b, 0xc6, 0xe6, 0xb5, 0x9d,
	0x78, 0x08, 0x3e, 0x81, 0xc7, 0x00, 0x36, 0x85, 0x54, 0x6a, 0x2c, 0x8a, 0x9d, 0x35, 0x61, 0x2f,
	0x3d, 0xe4, 0xd2, 0x6e, 0xa2, 0xed, 0x21, 0x24, 0xf1, 0x44, 0x48, 0x17, 0x1f, 0x98, 0x78, 0x22,
	0x81, 0xfc, 0xfa, 0xf9, 0xc7, 0x37, 0x3e, 0x63, 0x03, 0x0f, 0x41, 0x5a, 0x55, 0x80, 0xe8, 0xcc,
	0x3b, 0xcb, 0xe1, 0x7d, 0xdf, 0x43, 0xb8, 0x53, 0x05, 0x2c, 0xfe, 0x76, 0xd9, 0x48, 0x6f, 0x4d,
	0xc3, 0xf3, 0x4b, 0x36, 0xa0, 0xd0, 0x1a, 0xb7, 0x35, 0xda, 0x7e, 0x57, 0x65, 0x1e, 0x5c, 0x16,
	0xcb, 0xbc, 0x88, 0x65, 0x1e, 0x5c, 0x56, 0x95, 0xe1, 0x57, 0x6c, 0x4c, 0x98, 0x34, 0xd6, 0x07,
	0x65, 0x35, 0x88, 0x23, 0x02, 0x4e, 0x48, 0xbd, 0xad, 0x45, 0x7e, 0xc1, 0xfa, 0xaa, 0x2e, 0xf0,
	0x92, 0xfc, 0x9e, 0x8a, 0xff, 0x57, 0x09, 0x1b, 0xa7, 0x5b, 0x27, 0xac, 0xad, 0xb7, 0xec, 0xd4,
	0x82, 0xc9, 0x37, 0x6b, 0x74, 0x52, 0xa5, 0xa9, 0x03, 0xef, 0x45, 0x8f, 0x90, 0x49, 0xa3, 0x7f,
	0x8a, 0x32, 0x5f, 0x1e, 0xa2, 0x75, 0xb5, 0x3e, 0xa1, 0xe3, 0x16, 0x8d, 0x45, 0x5f, 0x33, 0x96,
	0x3b, 0xdc, 0x95, 0x91, 0x19, 0x10, 0x33, 0x24, 0x85, 0xec, 0x57, 0x6c, 0x98, 0x1a, 0x07, 0x3a,
	0x18, 0xb4, 0x62, 0x18, 0xdd, 0x56, 0xe0, 0xe7, 0xac, 0x4b, 0xa8, 0x60, 0xe4, 0xc4, 0x0f, 0x7e,
	0xcd, 0x26, 0x1e, 0x77, 0x4e, 0x83, 0x6c, 0x1b, 0x38, 0x8a, 0x67, 0x47, 0x79, 0xd5, 0xb4, 0xf1,
	0x23, 0x9b, 0xa9, 0x3c, 0x77, 0x90, 0xab, 0x00, 0xd2, 0x42, 0xf8, 0x85, 0xee, 0x67, 0x7b, 0xb3,
	0x63, 0xfa, 0xe5, 0xa2, 0x05, 0xee, 0xa2, 0xdf, 0xdc, 0xf0, 0x8a, 0x8d, 0x8d, 0x0d, 0xe0, 0x32,
	0xa5, 0x21, 0x66, 0x3f, 0x89, 0x7d, 0x6e, 0x55, 0xca, 0x7f, 0xc9, 0x06, 0xed, 0x20, 0xc6, 0x71,
	0x8a, 0xe6, 0x70, 0x06, 0x0e, 0x94, 0x34, 0xa9, 0x98, 0xd4, 0x33, 0x70, 0xa0, 0x6e, 0xd3, 0x7a,
	0x86, 0x25, 0x52, 0xae, 0xcc, 0x61, 0x21, 0x4e, 0xe7, 0x9d, 0xe5, 0x94, 0x66, 0x18, 0xd5, 0x2f,
	0x0e, 0x0b, 0xfe, 0x86, 0x1d, 0xff, 0xc7, 0x02, 0x8a, 0x29, 0x41, 0xa3, 0x56, 0xfb, 0x8e, 0xfc,
	0x1d, 0x9b, 0x56, 0xeb, 0x0c, 0xcd, 0xae, 0x53, 0x50, 0x1e, 0x67, 0x46, 0xc6, 0x8a, 0x74, 0x8a,
	0x7a, 0xcd, 0x26, 0xcd, 0x0e, 0xd7, 0xb8, 0x38, 0x8b, 0x6d, 0x6b, 0xe4, 0x08, 0x57, 0xe7, 0x46,
	0x45, 0x96, 0x68, 0x6c, 0x10, 0xe7, 0x44, 0x8d, 0xa2, 0xb6, 0xaa, 0xa4, 0xc5, 0x9f, 0x0e, 0x3b,
	0x7b, 0xe2, 0x11, 0xf0, 0xdf, 0x1d, 0xd6, 0x5f, 0x1b, 0x9b, 0x1a, 0x9b, 0x8b, 0x9b, 0xf9, 0xd1,
	0x72, 0x74, 0x53, 0x26, 0xcf, 0xf0, 0x26, 0x93, 0x83, 0x87, 0x75, 0xdf, 0x04, 0x58, 0xf7, 0x68,
	0x3d, 0xde, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xfc, 0xf6, 0x62, 0x31, 0x04, 0x00, 0x00,
}
