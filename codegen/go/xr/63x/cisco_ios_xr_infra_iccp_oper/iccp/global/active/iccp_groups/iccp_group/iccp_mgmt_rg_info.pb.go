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
// source: iccp_mgmt_rg_info.proto

package cisco_ios_xr_infra_iccp_oper_iccp_global_active_iccp_groups_iccp_group

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

type IccpMgmtRgInfo_KEYS struct {
	GroupNumber          uint32   `protobuf:"varint,1,opt,name=group_number,json=groupNumber,proto3" json:"group_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IccpMgmtRgInfo_KEYS) Reset()         { *m = IccpMgmtRgInfo_KEYS{} }
func (m *IccpMgmtRgInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*IccpMgmtRgInfo_KEYS) ProtoMessage()    {}
func (*IccpMgmtRgInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccefc9c8a4161ca5, []int{0}
}

func (m *IccpMgmtRgInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpMgmtRgInfo_KEYS.Unmarshal(m, b)
}
func (m *IccpMgmtRgInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpMgmtRgInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *IccpMgmtRgInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpMgmtRgInfo_KEYS.Merge(m, src)
}
func (m *IccpMgmtRgInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_IccpMgmtRgInfo_KEYS.Size(m)
}
func (m *IccpMgmtRgInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpMgmtRgInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IccpMgmtRgInfo_KEYS proto.InternalMessageInfo

func (m *IccpMgmtRgInfo_KEYS) GetGroupNumber() uint32 {
	if m != nil {
		return m.GroupNumber
	}
	return 0
}

type IccpMgmtMemberInfo struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	LdpState             string   `protobuf:"bytes,2,opt,name=ldp_state,json=ldpState,proto3" json:"ldp_state,omitempty"`
	IccpState            string   `protobuf:"bytes,3,opt,name=iccp_state,json=iccpState,proto3" json:"iccp_state,omitempty"`
	RouteWatchState      string   `protobuf:"bytes,4,opt,name=route_watch_state,json=routeWatchState,proto3" json:"route_watch_state,omitempty"`
	HostName             string   `protobuf:"bytes,5,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IccpMgmtMemberInfo) Reset()         { *m = IccpMgmtMemberInfo{} }
func (m *IccpMgmtMemberInfo) String() string { return proto.CompactTextString(m) }
func (*IccpMgmtMemberInfo) ProtoMessage()    {}
func (*IccpMgmtMemberInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccefc9c8a4161ca5, []int{1}
}

func (m *IccpMgmtMemberInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpMgmtMemberInfo.Unmarshal(m, b)
}
func (m *IccpMgmtMemberInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpMgmtMemberInfo.Marshal(b, m, deterministic)
}
func (m *IccpMgmtMemberInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpMgmtMemberInfo.Merge(m, src)
}
func (m *IccpMgmtMemberInfo) XXX_Size() int {
	return xxx_messageInfo_IccpMgmtMemberInfo.Size(m)
}
func (m *IccpMgmtMemberInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpMgmtMemberInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IccpMgmtMemberInfo proto.InternalMessageInfo

func (m *IccpMgmtMemberInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IccpMgmtMemberInfo) GetLdpState() string {
	if m != nil {
		return m.LdpState
	}
	return ""
}

func (m *IccpMgmtMemberInfo) GetIccpState() string {
	if m != nil {
		return m.IccpState
	}
	return ""
}

func (m *IccpMgmtMemberInfo) GetRouteWatchState() string {
	if m != nil {
		return m.RouteWatchState
	}
	return ""
}

func (m *IccpMgmtMemberInfo) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

type IccpMgmtBbiInfo struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IccpMgmtBbiInfo) Reset()         { *m = IccpMgmtBbiInfo{} }
func (m *IccpMgmtBbiInfo) String() string { return proto.CompactTextString(m) }
func (*IccpMgmtBbiInfo) ProtoMessage()    {}
func (*IccpMgmtBbiInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccefc9c8a4161ca5, []int{2}
}

func (m *IccpMgmtBbiInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpMgmtBbiInfo.Unmarshal(m, b)
}
func (m *IccpMgmtBbiInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpMgmtBbiInfo.Marshal(b, m, deterministic)
}
func (m *IccpMgmtBbiInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpMgmtBbiInfo.Merge(m, src)
}
func (m *IccpMgmtBbiInfo) XXX_Size() int {
	return xxx_messageInfo_IccpMgmtBbiInfo.Size(m)
}
func (m *IccpMgmtBbiInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpMgmtBbiInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IccpMgmtBbiInfo proto.InternalMessageInfo

func (m *IccpMgmtBbiInfo) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *IccpMgmtBbiInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type IccpMgmtAppInfo struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IccpMgmtAppInfo) Reset()         { *m = IccpMgmtAppInfo{} }
func (m *IccpMgmtAppInfo) String() string { return proto.CompactTextString(m) }
func (*IccpMgmtAppInfo) ProtoMessage()    {}
func (*IccpMgmtAppInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccefc9c8a4161ca5, []int{3}
}

func (m *IccpMgmtAppInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpMgmtAppInfo.Unmarshal(m, b)
}
func (m *IccpMgmtAppInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpMgmtAppInfo.Marshal(b, m, deterministic)
}
func (m *IccpMgmtAppInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpMgmtAppInfo.Merge(m, src)
}
func (m *IccpMgmtAppInfo) XXX_Size() int {
	return xxx_messageInfo_IccpMgmtAppInfo.Size(m)
}
func (m *IccpMgmtAppInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpMgmtAppInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IccpMgmtAppInfo proto.InternalMessageInfo

func (m *IccpMgmtAppInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type IccpMgmtIsolationRecoveryDelayTimerInfo struct {
	State                string   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Value                uint32   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Elapsed              uint32   `protobuf:"varint,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IccpMgmtIsolationRecoveryDelayTimerInfo) Reset() {
	*m = IccpMgmtIsolationRecoveryDelayTimerInfo{}
}
func (m *IccpMgmtIsolationRecoveryDelayTimerInfo) String() string { return proto.CompactTextString(m) }
func (*IccpMgmtIsolationRecoveryDelayTimerInfo) ProtoMessage()    {}
func (*IccpMgmtIsolationRecoveryDelayTimerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccefc9c8a4161ca5, []int{4}
}

func (m *IccpMgmtIsolationRecoveryDelayTimerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpMgmtIsolationRecoveryDelayTimerInfo.Unmarshal(m, b)
}
func (m *IccpMgmtIsolationRecoveryDelayTimerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpMgmtIsolationRecoveryDelayTimerInfo.Marshal(b, m, deterministic)
}
func (m *IccpMgmtIsolationRecoveryDelayTimerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpMgmtIsolationRecoveryDelayTimerInfo.Merge(m, src)
}
func (m *IccpMgmtIsolationRecoveryDelayTimerInfo) XXX_Size() int {
	return xxx_messageInfo_IccpMgmtIsolationRecoveryDelayTimerInfo.Size(m)
}
func (m *IccpMgmtIsolationRecoveryDelayTimerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpMgmtIsolationRecoveryDelayTimerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IccpMgmtIsolationRecoveryDelayTimerInfo proto.InternalMessageInfo

func (m *IccpMgmtIsolationRecoveryDelayTimerInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *IccpMgmtIsolationRecoveryDelayTimerInfo) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *IccpMgmtIsolationRecoveryDelayTimerInfo) GetElapsed() uint32 {
	if m != nil {
		return m.Elapsed
	}
	return 0
}

type IccpMgmtRgInfo struct {
	GroupId                     uint32                                   `protobuf:"varint,50,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Member                      []*IccpMgmtMemberInfo                    `protobuf:"bytes,51,rep,name=member,proto3" json:"member,omitempty"`
	BackboneInterface           []*IccpMgmtBbiInfo                       `protobuf:"bytes,52,rep,name=backbone_interface,json=backboneInterface,proto3" json:"backbone_interface,omitempty"`
	Application                 []*IccpMgmtAppInfo                       `protobuf:"bytes,53,rep,name=application,proto3" json:"application,omitempty"`
	IsolationRecoveryDelayTimer *IccpMgmtIsolationRecoveryDelayTimerInfo `protobuf:"bytes,54,opt,name=isolation_recovery_delay_timer,json=isolationRecoveryDelayTimer,proto3" json:"isolation_recovery_delay_timer,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                                 `json:"-"`
	XXX_unrecognized            []byte                                   `json:"-"`
	XXX_sizecache               int32                                    `json:"-"`
}

func (m *IccpMgmtRgInfo) Reset()         { *m = IccpMgmtRgInfo{} }
func (m *IccpMgmtRgInfo) String() string { return proto.CompactTextString(m) }
func (*IccpMgmtRgInfo) ProtoMessage()    {}
func (*IccpMgmtRgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccefc9c8a4161ca5, []int{5}
}

func (m *IccpMgmtRgInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpMgmtRgInfo.Unmarshal(m, b)
}
func (m *IccpMgmtRgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpMgmtRgInfo.Marshal(b, m, deterministic)
}
func (m *IccpMgmtRgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpMgmtRgInfo.Merge(m, src)
}
func (m *IccpMgmtRgInfo) XXX_Size() int {
	return xxx_messageInfo_IccpMgmtRgInfo.Size(m)
}
func (m *IccpMgmtRgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpMgmtRgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IccpMgmtRgInfo proto.InternalMessageInfo

func (m *IccpMgmtRgInfo) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *IccpMgmtRgInfo) GetMember() []*IccpMgmtMemberInfo {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *IccpMgmtRgInfo) GetBackboneInterface() []*IccpMgmtBbiInfo {
	if m != nil {
		return m.BackboneInterface
	}
	return nil
}

func (m *IccpMgmtRgInfo) GetApplication() []*IccpMgmtAppInfo {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *IccpMgmtRgInfo) GetIsolationRecoveryDelayTimer() *IccpMgmtIsolationRecoveryDelayTimerInfo {
	if m != nil {
		return m.IsolationRecoveryDelayTimer
	}
	return nil
}

func init() {
	proto.RegisterType((*IccpMgmtRgInfo_KEYS)(nil), "cisco_ios_xr_infra_iccp_oper.iccp.global.active.iccp_groups.iccp_group.iccp_mgmt_rg_info_KEYS")
	proto.RegisterType((*IccpMgmtMemberInfo)(nil), "cisco_ios_xr_infra_iccp_oper.iccp.global.active.iccp_groups.iccp_group.iccp_mgmt_member_info")
	proto.RegisterType((*IccpMgmtBbiInfo)(nil), "cisco_ios_xr_infra_iccp_oper.iccp.global.active.iccp_groups.iccp_group.iccp_mgmt_bbi_info")
	proto.RegisterType((*IccpMgmtAppInfo)(nil), "cisco_ios_xr_infra_iccp_oper.iccp.global.active.iccp_groups.iccp_group.iccp_mgmt_app_info")
	proto.RegisterType((*IccpMgmtIsolationRecoveryDelayTimerInfo)(nil), "cisco_ios_xr_infra_iccp_oper.iccp.global.active.iccp_groups.iccp_group.iccp_mgmt_isolation_recovery_delay_timer_info")
	proto.RegisterType((*IccpMgmtRgInfo)(nil), "cisco_ios_xr_infra_iccp_oper.iccp.global.active.iccp_groups.iccp_group.iccp_mgmt_rg_info")
}

func init() { proto.RegisterFile("iccp_mgmt_rg_info.proto", fileDescriptor_ccefc9c8a4161ca5) }

var fileDescriptor_ccefc9c8a4161ca5 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdf, 0x8a, 0x13, 0x3f,
	0x14, 0x66, 0x7e, 0xdb, 0xdd, 0xdf, 0xf6, 0xd4, 0x2a, 0x0d, 0xfe, 0x19, 0x29, 0x4a, 0x1d, 0x10,
	0x8a, 0xe0, 0x5c, 0x74, 0xd5, 0x1b, 0x6f, 0x55, 0x58, 0x84, 0x05, 0x67, 0x05, 0x51, 0x90, 0x90,
	0xc9, 0x64, 0xbb, 0xc1, 0xcc, 0x24, 0x24, 0x99, 0x6a, 0xc1, 0x17, 0xf0, 0x85, 0xbc, 0xf0, 0x99,
	0x7c, 0x08, 0xc9, 0x49, 0xa7, 0x1d, 0x58, 0x11, 0x2f, 0xea, 0x5d, 0xce, 0x77, 0xbe, 0x73, 0xbe,
	0x8f, 0x73, 0x0e, 0x81, 0x3b, 0x92, 0x73, 0x43, 0xeb, 0x65, 0xed, 0xa9, 0x5d, 0x52, 0xd9, 0x5c,
	0xe8, 0xdc, 0x58, 0xed, 0x35, 0x79, 0xc5, 0xa5, 0xe3, 0x9a, 0x4a, 0xed, 0xe8, 0x17, 0x1b, 0x12,
	0x96, 0x51, 0xe4, 0x6a, 0x23, 0x6c, 0x1e, 0x5e, 0xf9, 0x52, 0xe9, 0x92, 0xa9, 0x9c, 0x71, 0x2f,
	0x57, 0x02, 0x21, 0xba, 0xb4, 0xba, 0x35, 0xae, 0xf7, 0xce, 0x9e, 0xc3, 0xed, 0x2b, 0x12, 0xf4,
	0xf5, 0xcb, 0xf7, 0xe7, 0xe4, 0x01, 0x5c, 0x43, 0x0a, 0x6d, 0xda, 0xba, 0x14, 0x36, 0x4d, 0x66,
	0xc9, 0x7c, 0x5c, 0x8c, 0x10, 0x3b, 0x43, 0x28, 0xfb, 0x9e, 0xc0, 0xad, 0x5d, 0x75, 0x2d, 0x02,
	0x88, 0x1d, 0x48, 0x0a, 0xff, 0xb3, 0xaa, 0xb2, 0xc2, 0x39, 0xac, 0x1b, 0x16, 0x5d, 0x48, 0xa6,
	0x30, 0x54, 0x95, 0xa1, 0xce, 0x33, 0x2f, 0xd2, 0xff, 0x30, 0x77, 0xac, 0x2a, 0x73, 0x1e, 0x62,
	0x72, 0x0f, 0x00, 0xfb, 0xc5, 0xec, 0x01, 0x66, 0x87, 0x01, 0x89, 0xe9, 0x47, 0x30, 0xb1, 0xba,
	0xf5, 0x82, 0x7e, 0x66, 0x9e, 0x5f, 0x6e, 0x58, 0x03, 0x64, 0xdd, 0xc0, 0xc4, 0xbb, 0x80, 0x47,
	0xee, 0x14, 0x86, 0x97, 0xda, 0x79, 0xda, 0xb0, 0x5a, 0xa4, 0x87, 0x51, 0x27, 0x00, 0x67, 0xac,
	0x16, 0xd9, 0x1b, 0x20, 0x3b, 0xdf, 0x65, 0x29, 0xa3, 0xe9, 0x87, 0x70, 0x5d, 0x36, 0x5e, 0xd8,
	0x0b, 0xc6, 0x45, 0xac, 0x8b, 0xde, 0xc7, 0x5b, 0x34, 0x14, 0x93, 0x9b, 0x70, 0xd8, 0x77, 0x1f,
	0x83, 0x6c, 0xde, 0x6f, 0xc9, 0x8c, 0x89, 0x2d, 0x09, 0x0c, 0xfc, 0xda, 0x74, 0x8d, 0xf0, 0x9d,
	0xb5, 0xf0, 0x78, 0xc7, 0x94, 0x4e, 0x2b, 0xe6, 0xa5, 0x6e, 0xa8, 0x15, 0x5c, 0xaf, 0x84, 0x5d,
	0xd3, 0x4a, 0x28, 0xb6, 0xa6, 0x5e, 0xd6, 0xdd, 0x30, 0xb7, 0x82, 0x49, 0x4f, 0x30, 0xa0, 0x2b,
	0xa6, 0xda, 0x68, 0x63, 0x5c, 0xc4, 0x20, 0x0c, 0x5e, 0x28, 0x66, 0x9c, 0xa8, 0x70, 0x7c, 0xe3,
	0xa2, 0x0b, 0xb3, 0x9f, 0x03, 0x98, 0x5c, 0x59, 0x35, 0xb9, 0x0b, 0xc7, 0x71, 0xcb, 0xb2, 0x4a,
	0x17, 0xb1, 0x00, 0xe3, 0xd3, 0x8a, 0xb4, 0x70, 0x14, 0x57, 0x9a, 0x9e, 0xcc, 0x0e, 0xe6, 0xa3,
	0xc5, 0xc7, 0x7c, 0x3f, 0x37, 0x97, 0xff, 0xf6, 0x64, 0x8a, 0x8d, 0x18, 0xf9, 0x96, 0x00, 0x29,
	0x19, 0xff, 0x54, 0xea, 0x46, 0xd0, 0xed, 0xe4, 0xd3, 0x27, 0xe8, 0xe1, 0xc3, 0xfe, 0x3d, 0x74,
	0xeb, 0x2f, 0x26, 0x9d, 0xea, 0x69, 0x27, 0x4a, 0xbe, 0xc2, 0x88, 0x19, 0xa3, 0x24, 0xc7, 0x15,
	0xa5, 0x4f, 0xff, 0x95, 0x87, 0xee, 0x5e, 0x8a, 0xbe, 0x1c, 0xf9, 0x91, 0xc0, 0xfd, 0x3f, 0xdf,
	0x47, 0xfa, 0x6c, 0x96, 0xcc, 0x47, 0x8b, 0x76, 0xff, 0x8e, 0xfe, 0xe2, 0x2e, 0x8b, 0xe9, 0x96,
	0x54, 0x6c, 0x38, 0x2f, 0x02, 0xe5, 0x6d, 0x60, 0x94, 0x47, 0xf8, 0x4f, 0x9d, 0xfc, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x18, 0x8c, 0xa8, 0x7b, 0xc2, 0x04, 0x00, 0x00,
}
