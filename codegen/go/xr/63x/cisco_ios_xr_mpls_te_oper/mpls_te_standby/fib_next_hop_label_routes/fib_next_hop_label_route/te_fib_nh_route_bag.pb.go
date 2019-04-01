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
// source: te_fib_nh_route_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_fib_next_hop_label_routes_fib_next_hop_label_route

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

type TeFibNhRouteBag_KEYS struct {
	Label                uint32   `protobuf:"varint,1,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeFibNhRouteBag_KEYS) Reset()         { *m = TeFibNhRouteBag_KEYS{} }
func (m *TeFibNhRouteBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeFibNhRouteBag_KEYS) ProtoMessage()    {}
func (*TeFibNhRouteBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_45fc777012568099, []int{0}
}

func (m *TeFibNhRouteBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeFibNhRouteBag_KEYS.Unmarshal(m, b)
}
func (m *TeFibNhRouteBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeFibNhRouteBag_KEYS.Marshal(b, m, deterministic)
}
func (m *TeFibNhRouteBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeFibNhRouteBag_KEYS.Merge(m, src)
}
func (m *TeFibNhRouteBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeFibNhRouteBag_KEYS.Size(m)
}
func (m *TeFibNhRouteBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeFibNhRouteBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeFibNhRouteBag_KEYS proto.InternalMessageInfo

func (m *TeFibNhRouteBag_KEYS) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

type TeTargetAddr struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	IpAddress            string   `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Label                uint32   `protobuf:"varint,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeTargetAddr) Reset()         { *m = TeTargetAddr{} }
func (m *TeTargetAddr) String() string { return proto.CompactTextString(m) }
func (*TeTargetAddr) ProtoMessage()    {}
func (*TeTargetAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_45fc777012568099, []int{1}
}

func (m *TeTargetAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeTargetAddr.Unmarshal(m, b)
}
func (m *TeTargetAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeTargetAddr.Marshal(b, m, deterministic)
}
func (m *TeTargetAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeTargetAddr.Merge(m, src)
}
func (m *TeTargetAddr) XXX_Size() int {
	return xxx_messageInfo_TeTargetAddr.Size(m)
}
func (m *TeTargetAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_TeTargetAddr.DiscardUnknown(m)
}

var xxx_messageInfo_TeTargetAddr proto.InternalMessageInfo

func (m *TeTargetAddr) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TeTargetAddr) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *TeTargetAddr) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

type TeFibNhPathBag struct {
	Interface            string   `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	NextHopAddress       string   `protobuf:"bytes,2,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	LabelStack           []uint32 `protobuf:"varint,3,rep,packed,name=label_stack,json=labelStack,proto3" json:"label_stack,omitempty"`
	PathId               uint32   `protobuf:"varint,4,opt,name=path_id,json=pathId,proto3" json:"path_id,omitempty"`
	BackupPathId         uint32   `protobuf:"varint,5,opt,name=backup_path_id,json=backupPathId,proto3" json:"backup_path_id,omitempty"`
	PureBackup           bool     `protobuf:"varint,6,opt,name=pure_backup,json=pureBackup,proto3" json:"pure_backup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeFibNhPathBag) Reset()         { *m = TeFibNhPathBag{} }
func (m *TeFibNhPathBag) String() string { return proto.CompactTextString(m) }
func (*TeFibNhPathBag) ProtoMessage()    {}
func (*TeFibNhPathBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_45fc777012568099, []int{2}
}

func (m *TeFibNhPathBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeFibNhPathBag.Unmarshal(m, b)
}
func (m *TeFibNhPathBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeFibNhPathBag.Marshal(b, m, deterministic)
}
func (m *TeFibNhPathBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeFibNhPathBag.Merge(m, src)
}
func (m *TeFibNhPathBag) XXX_Size() int {
	return xxx_messageInfo_TeFibNhPathBag.Size(m)
}
func (m *TeFibNhPathBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TeFibNhPathBag.DiscardUnknown(m)
}

var xxx_messageInfo_TeFibNhPathBag proto.InternalMessageInfo

func (m *TeFibNhPathBag) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *TeFibNhPathBag) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *TeFibNhPathBag) GetLabelStack() []uint32 {
	if m != nil {
		return m.LabelStack
	}
	return nil
}

func (m *TeFibNhPathBag) GetPathId() uint32 {
	if m != nil {
		return m.PathId
	}
	return 0
}

func (m *TeFibNhPathBag) GetBackupPathId() uint32 {
	if m != nil {
		return m.BackupPathId
	}
	return 0
}

func (m *TeFibNhPathBag) GetPureBackup() bool {
	if m != nil {
		return m.PureBackup
	}
	return false
}

type TeFibNhRouteBag struct {
	PrefixAddr           string            `protobuf:"bytes,50,opt,name=prefix_addr,json=prefixAddr,proto3" json:"prefix_addr,omitempty"`
	TargetAddress        *TeTargetAddr     `protobuf:"bytes,51,opt,name=target_address,json=targetAddress,proto3" json:"target_address,omitempty"`
	CreateTime           uint32            `protobuf:"varint,52,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	LastUsedTime         uint32            `protobuf:"varint,53,opt,name=last_used_time,json=lastUsedTime,proto3" json:"last_used_time,omitempty"`
	LastUpdateTime       uint32            `protobuf:"varint,54,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	ExpireTime           uint32            `protobuf:"varint,55,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	IsRegistered         bool              `protobuf:"varint,56,opt,name=is_registered,json=isRegistered,proto3" json:"is_registered,omitempty"`
	IsNotified           bool              `protobuf:"varint,57,opt,name=is_notified,json=isNotified,proto3" json:"is_notified,omitempty"`
	IsStale              bool              `protobuf:"varint,58,opt,name=is_stale,json=isStale,proto3" json:"is_stale,omitempty"`
	RouteVersion         uint64            `protobuf:"varint,59,opt,name=route_version,json=routeVersion,proto3" json:"route_version,omitempty"`
	NextHopPath          []*TeFibNhPathBag `protobuf:"bytes,60,rep,name=next_hop_path,json=nextHopPath,proto3" json:"next_hop_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TeFibNhRouteBag) Reset()         { *m = TeFibNhRouteBag{} }
func (m *TeFibNhRouteBag) String() string { return proto.CompactTextString(m) }
func (*TeFibNhRouteBag) ProtoMessage()    {}
func (*TeFibNhRouteBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_45fc777012568099, []int{3}
}

func (m *TeFibNhRouteBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeFibNhRouteBag.Unmarshal(m, b)
}
func (m *TeFibNhRouteBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeFibNhRouteBag.Marshal(b, m, deterministic)
}
func (m *TeFibNhRouteBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeFibNhRouteBag.Merge(m, src)
}
func (m *TeFibNhRouteBag) XXX_Size() int {
	return xxx_messageInfo_TeFibNhRouteBag.Size(m)
}
func (m *TeFibNhRouteBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TeFibNhRouteBag.DiscardUnknown(m)
}

var xxx_messageInfo_TeFibNhRouteBag proto.InternalMessageInfo

func (m *TeFibNhRouteBag) GetPrefixAddr() string {
	if m != nil {
		return m.PrefixAddr
	}
	return ""
}

func (m *TeFibNhRouteBag) GetTargetAddress() *TeTargetAddr {
	if m != nil {
		return m.TargetAddress
	}
	return nil
}

func (m *TeFibNhRouteBag) GetCreateTime() uint32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *TeFibNhRouteBag) GetLastUsedTime() uint32 {
	if m != nil {
		return m.LastUsedTime
	}
	return 0
}

func (m *TeFibNhRouteBag) GetLastUpdateTime() uint32 {
	if m != nil {
		return m.LastUpdateTime
	}
	return 0
}

func (m *TeFibNhRouteBag) GetExpireTime() uint32 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *TeFibNhRouteBag) GetIsRegistered() bool {
	if m != nil {
		return m.IsRegistered
	}
	return false
}

func (m *TeFibNhRouteBag) GetIsNotified() bool {
	if m != nil {
		return m.IsNotified
	}
	return false
}

func (m *TeFibNhRouteBag) GetIsStale() bool {
	if m != nil {
		return m.IsStale
	}
	return false
}

func (m *TeFibNhRouteBag) GetRouteVersion() uint64 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *TeFibNhRouteBag) GetNextHopPath() []*TeFibNhPathBag {
	if m != nil {
		return m.NextHopPath
	}
	return nil
}

func init() {
	proto.RegisterType((*TeFibNhRouteBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fib_next_hop_label_routes.fib_next_hop_label_route.te_fib_nh_route_bag_KEYS")
	proto.RegisterType((*TeTargetAddr)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fib_next_hop_label_routes.fib_next_hop_label_route.te_target_addr")
	proto.RegisterType((*TeFibNhPathBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fib_next_hop_label_routes.fib_next_hop_label_route.te_fib_nh_path_bag")
	proto.RegisterType((*TeFibNhRouteBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fib_next_hop_label_routes.fib_next_hop_label_route.te_fib_nh_route_bag")
}

func init() { proto.RegisterFile("te_fib_nh_route_bag.proto", fileDescriptor_45fc777012568099) }

var fileDescriptor_45fc777012568099 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0xc7, 0x39, 0x93, 0xa6, 0xcd, 0x5c, 0x12, 0xca, 0x2a, 0xb8, 0x05, 0x25, 0x47, 0xf4, 0xe1,
	0x9e, 0x82, 0xa4, 0xfe, 0xd6, 0x17, 0x05, 0x41, 0x11, 0x44, 0x2e, 0x56, 0x28, 0x08, 0xcb, 0xde,
	0xdd, 0x24, 0x59, 0x7a, 0xc9, 0x2d, 0xbb, 0x1b, 0x49, 0xff, 0x0e, 0xfd, 0xf7, 0xfc, 0x4b, 0x7c,
	0x91, 0x9d, 0xcd, 0xaf, 0x4a, 0xfb, 0xd8, 0xb7, 0xcc, 0x67, 0xbe, 0x37, 0x33, 0xfb, 0x9d, 0xdd,
	0xc0, 0x89, 0x43, 0x31, 0x51, 0xb9, 0x58, 0xcc, 0x84, 0xa9, 0x97, 0x0e, 0x45, 0x2e, 0xa7, 0x43,
	0x6d, 0x6a, 0x57, 0xb3, 0x1f, 0x85, 0xb2, 0x45, 0x2d, 0x54, 0x6d, 0xc5, 0xca, 0x88, 0xb9, 0xae,
	0xac, 0x70, 0x28, 0x6a, 0x8d, 0x66, 0xb8, 0x09, 0xac, 0x93, 0x8b, 0x32, 0xbf, 0x1c, 0x52, 0x05,
	0x5c, 0x39, 0x31, 0xab, 0xb5, 0xa8, 0x64, 0x8e, 0x55, 0xa8, 0x66, 0x6f, 0xcc, 0x0c, 0x9e, 0x00,
	0xbf, 0xa6, 0xb5, 0xf8, 0xfc, 0xe1, 0x7c, 0xcc, 0xee, 0xc1, 0x01, 0x49, 0x79, 0x94, 0x44, 0x69,
	0x37, 0x0b, 0xc1, 0xe0, 0x1c, 0x7a, 0x0e, 0x85, 0x93, 0x66, 0x8a, 0x4e, 0xc8, 0xb2, 0x34, 0x8c,
	0x41, 0xd3, 0x5d, 0x6a, 0x24, 0x59, 0x3b, 0xa3, 0xdf, 0xec, 0x21, 0x80, 0xd2, 0x94, 0x46, 0x6b,
	0xf9, 0x1d, 0xca, 0xb4, 0x95, 0x7e, 0x17, 0xc0, 0xae, 0x74, 0x63, 0xbf, 0xf4, 0x9f, 0x08, 0xd8,
	0x6e, 0x1a, 0x2d, 0xdd, 0xcc, 0x0f, 0xc3, 0x1e, 0x40, 0x5b, 0x2d, 0x1c, 0x9a, 0x89, 0x2c, 0x36,
	0x4d, 0x76, 0x80, 0xa5, 0x70, 0xbc, 0x3d, 0xd9, 0xd5, 0x7e, 0x3d, 0xcf, 0x3f, 0xd6, 0xdb, 0xa6,
	0x7d, 0x88, 0xc3, 0xd1, 0xad, 0x93, 0xc5, 0x05, 0x6f, 0x24, 0x8d, 0xb4, 0x9b, 0x01, 0xa1, 0xb1,
	0x27, 0xec, 0x3e, 0x1c, 0x52, 0x53, 0x55, 0xf2, 0x26, 0xcd, 0xd5, 0xf2, 0xe1, 0xa7, 0x92, 0x3d,
	0x86, 0x5e, 0x2e, 0x8b, 0x8b, 0xa5, 0x16, 0x9b, 0xfc, 0x01, 0xe5, 0x3b, 0x81, 0x7e, 0x0d, 0xaa,
	0x3e, 0xc4, 0x7a, 0x69, 0xbc, 0x81, 0x1e, 0xf2, 0x56, 0x12, 0xa5, 0x47, 0x19, 0x78, 0xf4, 0x9e,
	0xc8, 0xe0, 0x6f, 0x13, 0xee, 0x5e, 0xe3, 0x36, 0x7d, 0x68, 0x70, 0xa2, 0x56, 0x74, 0x00, 0x3e,
	0xa2, 0xe9, 0x21, 0x20, 0x3f, 0x3c, 0xfb, 0x15, 0x41, 0x6f, 0xcf, 0x71, 0x7f, 0xc4, 0xd3, 0x24,
	0x4a, 0xe3, 0x51, 0x35, 0xbc, 0xcd, 0xdb, 0x31, 0xbc, 0xba, 0xe8, 0xac, 0x1b, 0x82, 0x3d, 0x3f,
	0x0b, 0x83, 0xd2, 0x8b, 0xd4, 0x1c, 0xf9, 0x53, 0xb2, 0x04, 0x02, 0xfa, 0xa6, 0xe6, 0xe8, 0x6d,
	0xab, 0xa4, 0x75, 0x62, 0x69, 0xb1, 0x0c, 0x9a, 0x67, 0xc1, 0x36, 0x4f, 0xcf, 0x2c, 0x96, 0xa4,
	0x4a, 0xe1, 0x38, 0xa8, 0x74, 0xb9, 0xad, 0xf5, 0x9c, 0x74, 0xf4, 0xf5, 0x19, 0x61, 0x52, 0xf6,
	0x21, 0xc6, 0x95, 0x56, 0x66, 0x2d, 0x7a, 0x11, 0x1a, 0x06, 0x44, 0x82, 0x47, 0xd0, 0x55, 0x56,
	0x18, 0x9c, 0x2a, 0xeb, 0xd0, 0x60, 0xc9, 0x5f, 0xd2, 0x0e, 0x3a, 0xca, 0x66, 0x5b, 0xe6, 0xab,
	0x28, 0x2b, 0x16, 0xb5, 0x53, 0x13, 0x85, 0x25, 0x7f, 0x15, 0xd6, 0xa4, 0xec, 0x97, 0x35, 0x61,
	0x27, 0x70, 0xa4, 0xac, 0xb7, 0xad, 0x42, 0xfe, 0x9a, 0xb2, 0x87, 0xca, 0x8e, 0x7d, 0xe8, 0x1b,
	0x84, 0xb5, 0xfd, 0x44, 0x63, 0x55, 0xbd, 0xe0, 0x6f, 0x92, 0x28, 0x6d, 0x66, 0x1d, 0x82, 0xdf,
	0x03, 0x63, 0xbf, 0x23, 0xe8, 0x6e, 0xed, 0xf4, 0x17, 0x86, 0xbf, 0x4d, 0x1a, 0x69, 0x3c, 0xd2,
	0xb7, 0xbe, 0xac, 0xff, 0x5e, 0x4e, 0x16, 0xaf, 0x5f, 0x80, 0xbf, 0xa0, 0x79, 0x8b, 0xfe, 0x4f,
	0x4e, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x1c, 0x87, 0x1a, 0x6c, 0x04, 0x00, 0x00,
}
