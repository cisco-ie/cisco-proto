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
// source: l2tp_cc_class_data.proto

package cisco_ios_xr_tunnel_l2tun_oper_l2tpv2_nodes_node_classes_class

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

type L2TpCcClassData_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ClassName            string   `protobuf:"bytes,2,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2TpCcClassData_KEYS) Reset()         { *m = L2TpCcClassData_KEYS{} }
func (m *L2TpCcClassData_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2TpCcClassData_KEYS) ProtoMessage()    {}
func (*L2TpCcClassData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e623e0798aa5a3, []int{0}
}

func (m *L2TpCcClassData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCcClassData_KEYS.Unmarshal(m, b)
}
func (m *L2TpCcClassData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCcClassData_KEYS.Marshal(b, m, deterministic)
}
func (m *L2TpCcClassData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCcClassData_KEYS.Merge(m, src)
}
func (m *L2TpCcClassData_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2TpCcClassData_KEYS.Size(m)
}
func (m *L2TpCcClassData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCcClassData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCcClassData_KEYS proto.InternalMessageInfo

func (m *L2TpCcClassData_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *L2TpCcClassData_KEYS) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

type L2TpCcClassData struct {
	IpTos                           uint32   `protobuf:"varint,50,opt,name=ip_tos,json=ipTos,proto3" json:"ip_tos,omitempty"`
	VrfName                         string   `protobuf:"bytes,51,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	ReceiveWindowSize               uint32   `protobuf:"varint,52,opt,name=receive_window_size,json=receiveWindowSize,proto3" json:"receive_window_size,omitempty"`
	ClassNameXr                     string   `protobuf:"bytes,53,opt,name=class_name_xr,json=classNameXr,proto3" json:"class_name_xr,omitempty"`
	DigestHash                      string   `protobuf:"bytes,54,opt,name=digest_hash,json=digestHash,proto3" json:"digest_hash,omitempty"`
	Password                        string   `protobuf:"bytes,55,opt,name=password,proto3" json:"password,omitempty"`
	EncodedPassword                 string   `protobuf:"bytes,56,opt,name=encoded_password,json=encodedPassword,proto3" json:"encoded_password,omitempty"`
	HostName                        string   `protobuf:"bytes,57,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	AccountingMethodList            string   `protobuf:"bytes,58,opt,name=accounting_method_list,json=accountingMethodList,proto3" json:"accounting_method_list,omitempty"`
	HelloTimeout                    uint32   `protobuf:"varint,59,opt,name=hello_timeout,json=helloTimeout,proto3" json:"hello_timeout,omitempty"`
	SetupTimeout                    uint32   `protobuf:"varint,60,opt,name=setup_timeout,json=setupTimeout,proto3" json:"setup_timeout,omitempty"`
	RetransmitMinimumTimeout        uint32   `protobuf:"varint,61,opt,name=retransmit_minimum_timeout,json=retransmitMinimumTimeout,proto3" json:"retransmit_minimum_timeout,omitempty"`
	RetransmitMaximumTimeout        uint32   `protobuf:"varint,62,opt,name=retransmit_maximum_timeout,json=retransmitMaximumTimeout,proto3" json:"retransmit_maximum_timeout,omitempty"`
	InitialRetransmitMinimumTimeout uint32   `protobuf:"varint,63,opt,name=initial_retransmit_minimum_timeout,json=initialRetransmitMinimumTimeout,proto3" json:"initial_retransmit_minimum_timeout,omitempty"`
	InitialRetransmitMaximumTimeout uint32   `protobuf:"varint,64,opt,name=initial_retransmit_maximum_timeout,json=initialRetransmitMaximumTimeout,proto3" json:"initial_retransmit_maximum_timeout,omitempty"`
	TimeoutNoUser                   uint32   `protobuf:"varint,65,opt,name=timeout_no_user,json=timeoutNoUser,proto3" json:"timeout_no_user,omitempty"`
	RetransmitRetries               uint32   `protobuf:"varint,66,opt,name=retransmit_retries,json=retransmitRetries,proto3" json:"retransmit_retries,omitempty"`
	InitialRetransmitRetries        uint32   `protobuf:"varint,67,opt,name=initial_retransmit_retries,json=initialRetransmitRetries,proto3" json:"initial_retransmit_retries,omitempty"`
	IsAuthenticationEnabled         bool     `protobuf:"varint,68,opt,name=is_authentication_enabled,json=isAuthenticationEnabled,proto3" json:"is_authentication_enabled,omitempty"`
	IsHidden                        bool     `protobuf:"varint,69,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden,omitempty"`
	IsDigestEnabled                 bool     `protobuf:"varint,70,opt,name=is_digest_enabled,json=isDigestEnabled,proto3" json:"is_digest_enabled,omitempty"`
	IsDigestCheckEnabled            bool     `protobuf:"varint,71,opt,name=is_digest_check_enabled,json=isDigestCheckEnabled,proto3" json:"is_digest_check_enabled,omitempty"`
	IsCongestionControlEnabled      bool     `protobuf:"varint,72,opt,name=is_congestion_control_enabled,json=isCongestionControlEnabled,proto3" json:"is_congestion_control_enabled,omitempty"`
	IsPeerAddressChecked            bool     `protobuf:"varint,73,opt,name=is_peer_address_checked,json=isPeerAddressChecked,proto3" json:"is_peer_address_checked,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *L2TpCcClassData) Reset()         { *m = L2TpCcClassData{} }
func (m *L2TpCcClassData) String() string { return proto.CompactTextString(m) }
func (*L2TpCcClassData) ProtoMessage()    {}
func (*L2TpCcClassData) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e623e0798aa5a3, []int{1}
}

func (m *L2TpCcClassData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCcClassData.Unmarshal(m, b)
}
func (m *L2TpCcClassData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCcClassData.Marshal(b, m, deterministic)
}
func (m *L2TpCcClassData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCcClassData.Merge(m, src)
}
func (m *L2TpCcClassData) XXX_Size() int {
	return xxx_messageInfo_L2TpCcClassData.Size(m)
}
func (m *L2TpCcClassData) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCcClassData.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCcClassData proto.InternalMessageInfo

func (m *L2TpCcClassData) GetIpTos() uint32 {
	if m != nil {
		return m.IpTos
	}
	return 0
}

func (m *L2TpCcClassData) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *L2TpCcClassData) GetReceiveWindowSize() uint32 {
	if m != nil {
		return m.ReceiveWindowSize
	}
	return 0
}

func (m *L2TpCcClassData) GetClassNameXr() string {
	if m != nil {
		return m.ClassNameXr
	}
	return ""
}

func (m *L2TpCcClassData) GetDigestHash() string {
	if m != nil {
		return m.DigestHash
	}
	return ""
}

func (m *L2TpCcClassData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *L2TpCcClassData) GetEncodedPassword() string {
	if m != nil {
		return m.EncodedPassword
	}
	return ""
}

func (m *L2TpCcClassData) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *L2TpCcClassData) GetAccountingMethodList() string {
	if m != nil {
		return m.AccountingMethodList
	}
	return ""
}

func (m *L2TpCcClassData) GetHelloTimeout() uint32 {
	if m != nil {
		return m.HelloTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetSetupTimeout() uint32 {
	if m != nil {
		return m.SetupTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetRetransmitMinimumTimeout() uint32 {
	if m != nil {
		return m.RetransmitMinimumTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetRetransmitMaximumTimeout() uint32 {
	if m != nil {
		return m.RetransmitMaximumTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetInitialRetransmitMinimumTimeout() uint32 {
	if m != nil {
		return m.InitialRetransmitMinimumTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetInitialRetransmitMaximumTimeout() uint32 {
	if m != nil {
		return m.InitialRetransmitMaximumTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetTimeoutNoUser() uint32 {
	if m != nil {
		return m.TimeoutNoUser
	}
	return 0
}

func (m *L2TpCcClassData) GetRetransmitRetries() uint32 {
	if m != nil {
		return m.RetransmitRetries
	}
	return 0
}

func (m *L2TpCcClassData) GetInitialRetransmitRetries() uint32 {
	if m != nil {
		return m.InitialRetransmitRetries
	}
	return 0
}

func (m *L2TpCcClassData) GetIsAuthenticationEnabled() bool {
	if m != nil {
		return m.IsAuthenticationEnabled
	}
	return false
}

func (m *L2TpCcClassData) GetIsHidden() bool {
	if m != nil {
		return m.IsHidden
	}
	return false
}

func (m *L2TpCcClassData) GetIsDigestEnabled() bool {
	if m != nil {
		return m.IsDigestEnabled
	}
	return false
}

func (m *L2TpCcClassData) GetIsDigestCheckEnabled() bool {
	if m != nil {
		return m.IsDigestCheckEnabled
	}
	return false
}

func (m *L2TpCcClassData) GetIsCongestionControlEnabled() bool {
	if m != nil {
		return m.IsCongestionControlEnabled
	}
	return false
}

func (m *L2TpCcClassData) GetIsPeerAddressChecked() bool {
	if m != nil {
		return m.IsPeerAddressChecked
	}
	return false
}

func init() {
	proto.RegisterType((*L2TpCcClassData_KEYS)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.nodes.node.classes.class.l2tp_cc_class_data_KEYS")
	proto.RegisterType((*L2TpCcClassData)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.nodes.node.classes.class.l2tp_cc_class_data")
}

func init() { proto.RegisterFile("l2tp_cc_class_data.proto", fileDescriptor_91e623e0798aa5a3) }

var fileDescriptor_91e623e0798aa5a3 = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x5f, 0x4f, 0x13, 0x4d,
	0x14, 0xc6, 0xd3, 0x37, 0x79, 0x79, 0xdb, 0xe1, 0x6d, 0x90, 0x15, 0x65, 0x80, 0x10, 0x48, 0x4d,
	0x0c, 0x9a, 0xd8, 0x8b, 0x02, 0xfe, 0x41, 0x44, 0x6b, 0x41, 0x31, 0x08, 0x21, 0x05, 0xa2, 0x5e,
	0x4d, 0x86, 0xdd, 0x03, 0x7b, 0xe2, 0xee, 0xcc, 0x66, 0xce, 0x6c, 0x21, 0x5c, 0xf9, 0xd1, 0xcd,
	0xce, 0x74, 0xb7, 0xd6, 0x5a, 0x6f, 0xa0, 0xfb, 0x3c, 0xcf, 0xef, 0x39, 0x67, 0x26, 0xdb, 0x32,
	0x9e, 0x74, 0x6c, 0x26, 0xc2, 0x50, 0x84, 0x89, 0x24, 0x12, 0x91, 0xb4, 0xb2, 0x9d, 0x19, 0x6d,
	0x75, 0xb0, 0x17, 0x22, 0x85, 0x5a, 0xa0, 0x26, 0x71, 0x6b, 0x84, 0xcd, 0x95, 0x82, 0x44, 0x24,
	0x1d, 0x9b, 0x2b, 0xa1, 0x33, 0x30, 0xed, 0x02, 0x1c, 0x74, 0xda, 0x4a, 0x47, 0x40, 0xee, 0x6f,
	0xdb, 0x55, 0x00, 0xf9, 0xff, 0xad, 0x0b, 0xb6, 0x38, 0xd9, 0x2d, 0x8e, 0x0e, 0xbe, 0x9d, 0x05,
	0x2b, 0xac, 0x51, 0x00, 0x42, 0xc9, 0x14, 0x78, 0x6d, 0xbd, 0xb6, 0xd1, 0xe8, 0xd7, 0x0b, 0xe1,
	0x44, 0xa6, 0x10, 0xac, 0x32, 0xe6, 0xf3, 0xce, 0xfd, 0xc7, 0xb9, 0x0d, 0xa7, 0x14, 0x76, 0xeb,
	0x47, 0x83, 0x05, 0x93, 0xbd, 0xc1, 0x03, 0x36, 0x83, 0x99, 0xb0, 0x9a, 0x78, 0x67, 0xbd, 0xb6,
	0xd1, 0xec, 0xff, 0x8b, 0xd9, 0xb9, 0xa6, 0x60, 0x89, 0xd5, 0x07, 0xe6, 0xca, 0x57, 0x6d, 0xba,
	0xaa, 0xff, 0x06, 0xe6, 0xca, 0xcd, 0x69, 0xb3, 0xfb, 0x06, 0x42, 0xc0, 0x01, 0x88, 0x1b, 0x54,
	0x91, 0xbe, 0x11, 0x84, 0x77, 0xc0, 0xb7, 0x1c, 0x3e, 0x3f, 0xb4, 0xbe, 0x38, 0xe7, 0x0c, 0xef,
	0x20, 0x68, 0xb1, 0xe6, 0x68, 0x2f, 0x71, 0x6b, 0xf8, 0xb6, 0xeb, 0x9b, 0xad, 0x56, 0xfb, 0x6a,
	0x82, 0x35, 0x36, 0x1b, 0xe1, 0x35, 0x90, 0x15, 0xb1, 0xa4, 0x98, 0x3f, 0x77, 0x09, 0xe6, 0xa5,
	0x43, 0x49, 0x71, 0xb0, 0xcc, 0xea, 0x99, 0x24, 0xba, 0xd1, 0x26, 0xe2, 0x2f, 0xfc, 0xc1, 0xcb,
	0xe7, 0xe0, 0x09, 0xbb, 0x07, 0x2a, 0xd4, 0x11, 0x44, 0xa2, 0xca, 0xbc, 0x74, 0x99, 0xb9, 0xa1,
	0x7e, 0x5a, 0x46, 0x57, 0x58, 0x23, 0xd6, 0x64, 0xfd, 0xb9, 0x5e, 0xf9, 0x9e, 0x42, 0x70, 0x07,
	0xdb, 0x62, 0x0f, 0x65, 0x18, 0xea, 0x5c, 0x59, 0x54, 0xd7, 0x22, 0x05, 0x1b, 0xeb, 0x48, 0x24,
	0x48, 0x96, 0xef, 0xb8, 0xe4, 0xc2, 0xc8, 0x3d, 0x76, 0xe6, 0x67, 0x24, 0x1b, 0x3c, 0x62, 0xcd,
	0x18, 0x92, 0x44, 0x0b, 0x8b, 0x29, 0xe8, 0xdc, 0xf2, 0xd7, 0xee, 0x22, 0xfe, 0x77, 0xe2, 0xb9,
	0xd7, 0x8a, 0x10, 0x81, 0xcd, 0xb3, 0x2a, 0xb4, 0xeb, 0x43, 0x4e, 0x2c, 0x43, 0xbb, 0x6c, 0xd9,
	0x80, 0x35, 0x52, 0x51, 0x8a, 0x56, 0xa4, 0xa8, 0x30, 0xcd, 0xd3, 0x8a, 0x78, 0xe3, 0x08, 0x3e,
	0x4a, 0x1c, 0xfb, 0xc0, 0x14, 0x5a, 0xde, 0x8e, 0xd1, 0x7b, 0x13, 0xb4, 0x0f, 0x94, 0xf4, 0x11,
	0x6b, 0xa1, 0x42, 0x8b, 0x32, 0x11, 0x7f, 0xd9, 0xe1, 0xad, 0x6b, 0x59, 0x1b, 0x26, 0xfb, 0xd3,
	0x56, 0x99, 0x52, 0xf6, 0xdb, 0x4a, 0xef, 0xa6, 0x95, 0x8d, 0x6f, 0xf6, 0x98, 0xcd, 0x0d, 0x09,
	0xa1, 0xb4, 0xc8, 0x09, 0x0c, 0xef, 0x3a, 0xb2, 0x39, 0x94, 0x4f, 0xf4, 0x05, 0x81, 0x09, 0x9e,
	0xb1, 0xe0, 0x97, 0x61, 0xc5, 0x47, 0x04, 0xe2, 0xef, 0xcb, 0xb7, 0xb2, 0x74, 0xfa, 0xde, 0x28,
	0xae, 0xeb, 0x0f, 0x3b, 0x96, 0x58, 0xcf, 0x5f, 0xd7, 0xc4, 0x6e, 0x25, 0xbd, 0xc3, 0x96, 0x90,
	0x84, 0xcc, 0x6d, 0x0c, 0xca, 0x62, 0x28, 0x2d, 0x6a, 0x25, 0x40, 0xc9, 0xcb, 0x04, 0x22, 0xbe,
	0xbf, 0x5e, 0xdb, 0xa8, 0xf7, 0x17, 0x91, 0xba, 0x63, 0xfe, 0x81, 0xb7, 0x8b, 0x77, 0x10, 0x49,
	0xc4, 0x18, 0x45, 0xa0, 0xf8, 0x81, 0xcb, 0xd6, 0x91, 0x0e, 0xdd, 0x73, 0xf0, 0x94, 0xcd, 0x23,
	0x89, 0xe1, 0x77, 0xa1, 0x2c, 0xfc, 0xe0, 0x42, 0x73, 0x48, 0xfb, 0x4e, 0x2f, 0x8b, 0xb6, 0xd9,
	0xe2, 0x28, 0x1b, 0xc6, 0x10, 0x7e, 0xaf, 0x88, 0x8f, 0x8e, 0x58, 0x28, 0x89, 0x5e, 0x61, 0x96,
	0x58, 0x97, 0xad, 0x22, 0x89, 0x50, 0xab, 0xc2, 0x29, 0xf6, 0x0e, 0xb5, 0xb2, 0x46, 0x27, 0x15,
	0x7c, 0xe8, 0xe0, 0x65, 0xa4, 0x5e, 0x95, 0xe9, 0xf9, 0xc8, 0xf8, 0xe4, 0x0c, 0xc0, 0x08, 0x19,
	0x45, 0x06, 0x88, 0xfc, 0x7c, 0x88, 0xf8, 0xa7, 0x72, 0xf2, 0x29, 0x80, 0xe9, 0x7a, 0xb3, 0xe7,
	0xbd, 0xcb, 0x19, 0xf7, 0x03, 0xb9, 0xf9, 0x33, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x96, 0xc6, 0x58,
	0x3c, 0x05, 0x00, 0x00,
}