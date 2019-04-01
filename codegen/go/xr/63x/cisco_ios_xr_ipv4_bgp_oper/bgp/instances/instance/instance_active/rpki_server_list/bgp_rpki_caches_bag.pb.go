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
// source: bgp_rpki_caches_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_rpki_server_list

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

type BgpRpkiCachesBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpRpkiCachesBag_KEYS) Reset()         { *m = BgpRpkiCachesBag_KEYS{} }
func (m *BgpRpkiCachesBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpRpkiCachesBag_KEYS) ProtoMessage()    {}
func (*BgpRpkiCachesBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_71e51de5f701a78d, []int{0}
}

func (m *BgpRpkiCachesBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRpkiCachesBag_KEYS.Unmarshal(m, b)
}
func (m *BgpRpkiCachesBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRpkiCachesBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpRpkiCachesBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRpkiCachesBag_KEYS.Merge(m, src)
}
func (m *BgpRpkiCachesBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpRpkiCachesBag_KEYS.Size(m)
}
func (m *BgpRpkiCachesBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRpkiCachesBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRpkiCachesBag_KEYS proto.InternalMessageInfo

func (m *BgpRpkiCachesBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type BgpEdmRpkiCacheType_ struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Preference           uint32   `protobuf:"varint,2,opt,name=preference,proto3" json:"preference,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	StateTime            uint32   `protobuf:"varint,5,opt,name=state_time,json=stateTime,proto3" json:"state_time,omitempty"`
	Shutdown             bool     `protobuf:"varint,6,opt,name=shutdown,proto3" json:"shutdown,omitempty"`
	Retries              uint32   `protobuf:"varint,7,opt,name=retries,proto3" json:"retries,omitempty"`
	CloseReason          string   `protobuf:"bytes,8,opt,name=close_reason,json=closeReason,proto3" json:"close_reason,omitempty"`
	CloseTime            uint32   `protobuf:"varint,9,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	CloseTimeReal        uint32   `protobuf:"varint,10,opt,name=close_time_real,json=closeTimeReal,proto3" json:"close_time_real,omitempty"`
	ReadBytes            uint32   `protobuf:"varint,11,opt,name=read_bytes,json=readBytes,proto3" json:"read_bytes,omitempty"`
	WriteBytes           uint32   `protobuf:"varint,12,opt,name=write_bytes,json=writeBytes,proto3" json:"write_bytes,omitempty"`
	Transport            uint32   `protobuf:"varint,13,opt,name=transport,proto3" json:"transport,omitempty"`
	Username             string   `protobuf:"bytes,14,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,15,opt,name=password,proto3" json:"password,omitempty"`
	Sshpid               uint32   `protobuf:"varint,16,opt,name=sshpid,proto3" json:"sshpid,omitempty"`
	ProtoState           string   `protobuf:"bytes,17,opt,name=proto_state,json=protoState,proto3" json:"proto_state,omitempty"`
	ProtoStateTime       uint32   `protobuf:"varint,18,opt,name=proto_state_time,json=protoStateTime,proto3" json:"proto_state_time,omitempty"`
	Serial               uint32   `protobuf:"varint,19,opt,name=serial,proto3" json:"serial,omitempty"`
	Nonce                uint32   `protobuf:"varint,20,opt,name=nonce,proto3" json:"nonce,omitempty"`
	RefreshTime          int32    `protobuf:"zigzag32,21,opt,name=refresh_time,json=refreshTime,proto3" json:"refresh_time,omitempty"`
	ResponseTime         int32    `protobuf:"zigzag32,22,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	PurgeTime            int32    `protobuf:"zigzag32,23,opt,name=purge_time,json=purgeTime,proto3" json:"purge_time,omitempty"`
	Ipv4Roa              uint32   `protobuf:"varint,24,opt,name=ipv4roa,proto3" json:"ipv4roa,omitempty"`
	Ipv4RoaAnnounce      uint32   `protobuf:"varint,25,opt,name=ipv4roa_announce,json=ipv4roaAnnounce,proto3" json:"ipv4roa_announce,omitempty"`
	Ipv4RoaWithdraw      uint32   `protobuf:"varint,26,opt,name=ipv4roa_withdraw,json=ipv4roaWithdraw,proto3" json:"ipv4roa_withdraw,omitempty"`
	Ipv6Roa              uint32   `protobuf:"varint,27,opt,name=ipv6roa,proto3" json:"ipv6roa,omitempty"`
	Ipv6RoaAnnounce      uint32   `protobuf:"varint,28,opt,name=ipv6roa_announce,json=ipv6roaAnnounce,proto3" json:"ipv6roa_announce,omitempty"`
	Ipv6RoaWithdraw      uint32   `protobuf:"varint,29,opt,name=ipv6roa_withdraw,json=ipv6roaWithdraw,proto3" json:"ipv6roa_withdraw,omitempty"`
	ProtoError           string   `protobuf:"bytes,30,opt,name=proto_error,json=protoError,proto3" json:"proto_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpEdmRpkiCacheType_) Reset()         { *m = BgpEdmRpkiCacheType_{} }
func (m *BgpEdmRpkiCacheType_) String() string { return proto.CompactTextString(m) }
func (*BgpEdmRpkiCacheType_) ProtoMessage()    {}
func (*BgpEdmRpkiCacheType_) Descriptor() ([]byte, []int) {
	return fileDescriptor_71e51de5f701a78d, []int{1}
}

func (m *BgpEdmRpkiCacheType_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpEdmRpkiCacheType_.Unmarshal(m, b)
}
func (m *BgpEdmRpkiCacheType_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpEdmRpkiCacheType_.Marshal(b, m, deterministic)
}
func (m *BgpEdmRpkiCacheType_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpEdmRpkiCacheType_.Merge(m, src)
}
func (m *BgpEdmRpkiCacheType_) XXX_Size() int {
	return xxx_messageInfo_BgpEdmRpkiCacheType_.Size(m)
}
func (m *BgpEdmRpkiCacheType_) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpEdmRpkiCacheType_.DiscardUnknown(m)
}

var xxx_messageInfo_BgpEdmRpkiCacheType_ proto.InternalMessageInfo

func (m *BgpEdmRpkiCacheType_) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetPreference() uint32 {
	if m != nil {
		return m.Preference
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetStateTime() uint32 {
	if m != nil {
		return m.StateTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetShutdown() bool {
	if m != nil {
		return m.Shutdown
	}
	return false
}

func (m *BgpEdmRpkiCacheType_) GetRetries() uint32 {
	if m != nil {
		return m.Retries
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetCloseReason() string {
	if m != nil {
		return m.CloseReason
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetCloseTime() uint32 {
	if m != nil {
		return m.CloseTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetCloseTimeReal() uint32 {
	if m != nil {
		return m.CloseTimeReal
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetReadBytes() uint32 {
	if m != nil {
		return m.ReadBytes
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetWriteBytes() uint32 {
	if m != nil {
		return m.WriteBytes
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetTransport() uint32 {
	if m != nil {
		return m.Transport
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetSshpid() uint32 {
	if m != nil {
		return m.Sshpid
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetProtoState() string {
	if m != nil {
		return m.ProtoState
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetProtoStateTime() uint32 {
	if m != nil {
		return m.ProtoStateTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetSerial() uint32 {
	if m != nil {
		return m.Serial
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetRefreshTime() int32 {
	if m != nil {
		return m.RefreshTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetResponseTime() int32 {
	if m != nil {
		return m.ResponseTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetPurgeTime() int32 {
	if m != nil {
		return m.PurgeTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv4Roa() uint32 {
	if m != nil {
		return m.Ipv4Roa
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv4RoaAnnounce() uint32 {
	if m != nil {
		return m.Ipv4RoaAnnounce
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv4RoaWithdraw() uint32 {
	if m != nil {
		return m.Ipv4RoaWithdraw
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv6Roa() uint32 {
	if m != nil {
		return m.Ipv6Roa
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv6RoaAnnounce() uint32 {
	if m != nil {
		return m.Ipv6RoaAnnounce
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv6RoaWithdraw() uint32 {
	if m != nil {
		return m.Ipv6RoaWithdraw
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetProtoError() string {
	if m != nil {
		return m.ProtoError
	}
	return ""
}

type BgpRpkiCachesBag struct {
	RpkiServer           []*BgpEdmRpkiCacheType_ `protobuf:"bytes,50,rep,name=rpki_server,json=rpkiServer,proto3" json:"rpki_server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *BgpRpkiCachesBag) Reset()         { *m = BgpRpkiCachesBag{} }
func (m *BgpRpkiCachesBag) String() string { return proto.CompactTextString(m) }
func (*BgpRpkiCachesBag) ProtoMessage()    {}
func (*BgpRpkiCachesBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_71e51de5f701a78d, []int{2}
}

func (m *BgpRpkiCachesBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRpkiCachesBag.Unmarshal(m, b)
}
func (m *BgpRpkiCachesBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRpkiCachesBag.Marshal(b, m, deterministic)
}
func (m *BgpRpkiCachesBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRpkiCachesBag.Merge(m, src)
}
func (m *BgpRpkiCachesBag) XXX_Size() int {
	return xxx_messageInfo_BgpRpkiCachesBag.Size(m)
}
func (m *BgpRpkiCachesBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRpkiCachesBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRpkiCachesBag proto.InternalMessageInfo

func (m *BgpRpkiCachesBag) GetRpkiServer() []*BgpEdmRpkiCacheType_ {
	if m != nil {
		return m.RpkiServer
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpRpkiCachesBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rpki_server_list.bgp_rpki_caches_bag_KEYS")
	proto.RegisterType((*BgpEdmRpkiCacheType_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rpki_server_list.bgp_edm_rpki_cache_type_")
	proto.RegisterType((*BgpRpkiCachesBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rpki_server_list.bgp_rpki_caches_bag")
}

func init() { proto.RegisterFile("bgp_rpki_caches_bag.proto", fileDescriptor_71e51de5f701a78d) }

var fileDescriptor_71e51de5f701a78d = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x6e, 0x13, 0x3f,
	0x10, 0x56, 0x7e, 0xfd, 0x93, 0x66, 0x92, 0xf4, 0x8f, 0xdb, 0x5f, 0x71, 0x4b, 0x5b, 0x42, 0x2b,
	0xa1, 0x70, 0xc9, 0xa1, 0xa0, 0x5c, 0x11, 0x48, 0x3d, 0x21, 0x71, 0xd8, 0x22, 0x21, 0x4e, 0x96,
	0xb3, 0x99, 0x26, 0x16, 0xc9, 0xda, 0x1a, 0x3b, 0x0d, 0x7d, 0x08, 0x1e, 0x83, 0x57, 0xe1, 0xb9,
	0x90, 0xc7, 0x9b, 0xed, 0x46, 0x94, 0x1b, 0xb7, 0x99, 0x6f, 0xbe, 0xf9, 0x3e, 0xef, 0x8c, 0xd7,
	0x70, 0x32, 0x9a, 0x38, 0x45, 0xee, 0x9b, 0x51, 0xb9, 0xce, 0xa7, 0xe8, 0xd5, 0x48, 0x4f, 0x06,
	0x8e, 0x6c, 0xb0, 0x22, 0xcb, 0x8d, 0xcf, 0xad, 0x32, 0xd6, 0xab, 0xef, 0xa4, 0x8c, 0xbb, 0x7f,
	0xab, 0x22, 0xd9, 0x3a, 0xa4, 0xc1, 0x68, 0xe2, 0x06, 0xa6, 0xf0, 0x41, 0x17, 0x39, 0xfa, 0x2a,
	0xaa, 0x02, 0xa5, 0xf3, 0x60, 0xee, 0x71, 0xc0, 0xc2, 0x1e, 0xe9, 0x1e, 0x49, 0xcd, 0x8c, 0x0f,
	0x97, 0xef, 0x40, 0x3e, 0x61, 0xa8, 0x3e, 0xde, 0x7c, 0xbd, 0x15, 0x57, 0xd0, 0xad, 0xfa, 0x0b,
	0x3d, 0x47, 0xd9, 0xe8, 0x35, 0xfa, 0xad, 0xac, 0xb3, 0x02, 0x3f, 0xe9, 0x39, 0x5e, 0xfe, 0x6a,
	0x26, 0x05, 0x1c, 0xcf, 0x6b, 0x2a, 0x2a, 0x3c, 0x38, 0x54, 0x42, 0xc0, 0x66, 0xad, 0x91, 0x63,
	0x71, 0x01, 0xe0, 0x08, 0xef, 0x90, 0xb0, 0xc8, 0x51, 0xfe, 0xd7, 0x6b, 0xf4, 0xbb, 0x59, 0x0d,
	0x89, 0x3d, 0xce, 0x52, 0x90, 0x1b, 0x5c, 0xe1, 0x58, 0x1c, 0xc1, 0x96, 0x0f, 0x3a, 0xa0, 0xdc,
	0x64, 0xa1, 0x94, 0x88, 0x73, 0x00, 0x0e, 0x54, 0x30, 0x73, 0x94, 0x5b, 0xcc, 0x6f, 0x31, 0xf2,
	0xd9, 0xcc, 0x51, 0x9c, 0xc2, 0x8e, 0x9f, 0x2e, 0xc2, 0xd8, 0x2e, 0x0b, 0xb9, 0xdd, 0x6b, 0xf4,
	0x77, 0xb2, 0x2a, 0x17, 0x12, 0x9a, 0x84, 0x81, 0x0c, 0x7a, 0xd9, 0xe4, 0xbe, 0x55, 0x2a, 0x5e,
	0x42, 0x27, 0x9f, 0x59, 0x8f, 0x8a, 0x50, 0x7b, 0x5b, 0xc8, 0x1d, 0x76, 0x6c, 0x33, 0x96, 0x31,
	0x14, 0x7d, 0x13, 0x85, 0x7d, 0x5b, 0xc9, 0x97, 0x11, 0xf6, 0x7d, 0x05, 0x7b, 0x8f, 0xe5, 0x28,
	0x33, 0x93, 0xc0, 0x9c, 0x6e, 0xc5, 0xc9, 0x50, 0xcf, 0xa2, 0x0c, 0xa1, 0x1e, 0xab, 0xd1, 0x43,
	0x40, 0x2f, 0xdb, 0x49, 0x26, 0x22, 0x1f, 0x22, 0x20, 0x5e, 0x40, 0x7b, 0x49, 0x26, 0x60, 0x59,
	0xef, 0xa4, 0x41, 0x31, 0x94, 0x08, 0x67, 0xd0, 0x0a, 0xa4, 0x0b, 0xcf, 0xd3, 0xea, 0xa6, 0xf6,
	0x0a, 0x88, 0x5f, 0xbf, 0xf0, 0x48, 0x3c, 0xfe, 0x5d, 0xfe, 0x86, 0x2a, 0x8f, 0x35, 0xa7, 0xbd,
	0x5f, 0x5a, 0x1a, 0xcb, 0xbd, 0x54, 0x5b, 0xe5, 0xe2, 0x18, 0xb6, 0xbd, 0x9f, 0x3a, 0x33, 0x96,
	0xfb, 0x2c, 0x59, 0x66, 0xf1, 0x38, 0x7c, 0x0b, 0x55, 0x5a, 0xc4, 0x01, 0xb7, 0x01, 0x43, 0xb7,
	0xbc, 0x8d, 0x3e, 0xec, 0xd7, 0x08, 0x69, 0x36, 0x82, 0x25, 0x76, 0x1f, 0x59, 0x3c, 0xa0, 0x68,
	0x81, 0x64, 0xf4, 0x4c, 0x1e, 0x96, 0x16, 0x9c, 0xc5, 0x2d, 0x17, 0x36, 0x5e, 0x8a, 0x23, 0x86,
	0x53, 0x12, 0x17, 0x42, 0x78, 0x47, 0xe8, 0xa7, 0x49, 0xf3, 0xff, 0x5e, 0xa3, 0x7f, 0x90, 0xb5,
	0x4b, 0x8c, 0x05, 0xaf, 0xa0, 0x4b, 0xe8, 0x9d, 0x2d, 0x56, 0x3b, 0x39, 0x66, 0x4e, 0x67, 0x05,
	0x32, 0xe9, 0x1c, 0xc0, 0x2d, 0x68, 0x52, 0x32, 0x9e, 0x31, 0xa3, 0xc5, 0x08, 0x97, 0x25, 0x34,
	0xe3, 0x1f, 0x45, 0x56, 0x4b, 0x99, 0x6e, 0x44, 0x99, 0x8a, 0xd7, 0xb0, 0x5f, 0x86, 0x4a, 0x17,
	0x85, 0x5d, 0xc4, 0x13, 0x9e, 0x30, 0x65, 0xaf, 0xc4, 0xdf, 0x97, 0x70, 0x9d, 0xba, 0x34, 0x61,
	0x3a, 0x26, 0xbd, 0x94, 0xa7, 0x6b, 0xd4, 0x2f, 0x25, 0x5c, 0xfa, 0x0d, 0xa3, 0xdf, 0xf3, 0xca,
	0x6f, 0xf8, 0xe8, 0x37, 0x5c, 0xf3, 0x3b, 0xab, 0x44, 0x86, 0x7f, 0xfa, 0x0d, 0xd7, 0xfc, 0xce,
	0xd7, 0xa8, 0x95, 0x5f, 0xb5, 0x3f, 0x24, 0xb2, 0x24, 0x2f, 0x6a, 0xfb, 0xbb, 0x89, 0xc8, 0xe5,
	0xcf, 0x06, 0x1c, 0x3e, 0xf1, 0x14, 0x88, 0x1f, 0x0d, 0x68, 0xd7, 0x9e, 0x0d, 0x79, 0xdd, 0xdb,
	0xe8, 0xb7, 0xaf, 0x67, 0x83, 0x7f, 0xff, 0x18, 0x0d, 0xfe, 0xf6, 0x8e, 0x64, 0x10, 0x91, 0x5b,
	0x66, 0x8e, 0xb6, 0xf9, 0xcc, 0x6f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xab, 0x5a, 0x35, 0x54,
	0x29, 0x05, 0x00, 0x00,
}
