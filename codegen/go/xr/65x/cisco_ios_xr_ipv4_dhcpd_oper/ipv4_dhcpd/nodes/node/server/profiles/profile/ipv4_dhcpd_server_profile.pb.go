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
// source: ipv4_dhcpd_server_profile.proto

package cisco_ios_xr_ipv4_dhcpd_oper_ipv4_dhcpd_nodes_node_server_profiles_profile

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

type Ipv4DhcpdServerProfile_KEYS struct {
	Nodeid               string   `protobuf:"bytes,1,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	ServerProfileName    string   `protobuf:"bytes,2,opt,name=server_profile_name,json=serverProfileName,proto3" json:"server_profile_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdServerProfile_KEYS) Reset()         { *m = Ipv4DhcpdServerProfile_KEYS{} }
func (m *Ipv4DhcpdServerProfile_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdServerProfile_KEYS) ProtoMessage()    {}
func (*Ipv4DhcpdServerProfile_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c75a5ddc1135fea3, []int{0}
}

func (m *Ipv4DhcpdServerProfile_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdServerProfile_KEYS.Unmarshal(m, b)
}
func (m *Ipv4DhcpdServerProfile_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdServerProfile_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdServerProfile_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdServerProfile_KEYS.Merge(m, src)
}
func (m *Ipv4DhcpdServerProfile_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdServerProfile_KEYS.Size(m)
}
func (m *Ipv4DhcpdServerProfile_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdServerProfile_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdServerProfile_KEYS proto.InternalMessageInfo

func (m *Ipv4DhcpdServerProfile_KEYS) GetNodeid() string {
	if m != nil {
		return m.Nodeid
	}
	return ""
}

func (m *Ipv4DhcpdServerProfile_KEYS) GetServerProfileName() string {
	if m != nil {
		return m.ServerProfileName
	}
	return ""
}

type Ipv4DhcpdServerProfile struct {
	ServerProfileNameXr              string   `protobuf:"bytes,50,opt,name=server_profile_name_xr,json=serverProfileNameXr,proto3" json:"server_profile_name_xr,omitempty"`
	SecureArp                        bool     `protobuf:"varint,51,opt,name=secure_arp,json=secureArp,proto3" json:"secure_arp,omitempty"`
	RequestedAddressCheck            bool     `protobuf:"varint,52,opt,name=requested_address_check,json=requestedAddressCheck,proto3" json:"requested_address_check,omitempty"`
	ServerIdCheck                    bool     `protobuf:"varint,53,opt,name=server_id_check,json=serverIdCheck,proto3" json:"server_id_check,omitempty"`
	DuplicateMacAddressCheck         bool     `protobuf:"varint,54,opt,name=duplicate_mac_address_check,json=duplicateMacAddressCheck,proto3" json:"duplicate_mac_address_check,omitempty"`
	DuplicateIpAddressCheck          bool     `protobuf:"varint,55,opt,name=duplicate_ip_address_check,json=duplicateIpAddressCheck,proto3" json:"duplicate_ip_address_check,omitempty"`
	IsMoveAllowed                    bool     `protobuf:"varint,56,opt,name=is_move_allowed,json=isMoveAllowed,proto3" json:"is_move_allowed,omitempty"`
	BcastPolicy                      uint32   `protobuf:"varint,57,opt,name=bcast_policy,json=bcastPolicy,proto3" json:"bcast_policy,omitempty"`
	GiaddrPolicy                     uint32   `protobuf:"varint,58,opt,name=giaddr_policy,json=giaddrPolicy,proto3" json:"giaddr_policy,omitempty"`
	SubnetMask                       string   `protobuf:"bytes,59,opt,name=subnet_mask,json=subnetMask,proto3" json:"subnet_mask,omitempty"`
	ServerPoolName                   string   `protobuf:"bytes,60,opt,name=server_pool_name,json=serverPoolName,proto3" json:"server_pool_name,omitempty"`
	ServerProfileLease               uint32   `protobuf:"varint,61,opt,name=server_profile_lease,json=serverProfileLease,proto3" json:"server_profile_lease,omitempty"`
	ServerProfileNetbiosNodeType     uint32   `protobuf:"varint,62,opt,name=server_profile_netbios_node_type,json=serverProfileNetbiosNodeType,proto3" json:"server_profile_netbios_node_type,omitempty"`
	ServerBootfileName               string   `protobuf:"bytes,63,opt,name=server_bootfile_name,json=serverBootfileName,proto3" json:"server_bootfile_name,omitempty"`
	ServerDomainName                 string   `protobuf:"bytes,64,opt,name=server_domain_name,json=serverDomainName,proto3" json:"server_domain_name,omitempty"`
	ServerProfileiedgeCheck          uint32   `protobuf:"varint,65,opt,name=server_profileiedge_check,json=serverProfileiedgeCheck,proto3" json:"server_profileiedge_check,omitempty"`
	ServerProfileServerDnsCount      uint32   `protobuf:"varint,66,opt,name=server_profile_server_dns_count,json=serverProfileServerDnsCount,proto3" json:"server_profile_server_dns_count,omitempty"`
	ServerProfiledefaultRouterCount  uint32   `protobuf:"varint,67,opt,name=server_profiledefault_router_count,json=serverProfiledefaultRouterCount,proto3" json:"server_profiledefault_router_count,omitempty"`
	ServerProfileNetbiosNameSvrCount uint32   `protobuf:"varint,68,opt,name=server_profile_netbios_name_svr_count,json=serverProfileNetbiosNameSvrCount,proto3" json:"server_profile_netbios_name_svr_count,omitempty"`
	ServerProfileTimeSvrCount        uint32   `protobuf:"varint,69,opt,name=server_profile_time_svr_count,json=serverProfileTimeSvrCount,proto3" json:"server_profile_time_svr_count,omitempty"`
	ServerProfileDns                 []string `protobuf:"bytes,70,rep,name=server_profile_dns,json=serverProfileDns,proto3" json:"server_profile_dns,omitempty"`
	ServerProfileDefaultRouter       []string `protobuf:"bytes,71,rep,name=server_profile_default_router,json=serverProfileDefaultRouter,proto3" json:"server_profile_default_router,omitempty"`
	ServerProfileNetbiousNameServer  []string `protobuf:"bytes,72,rep,name=server_profile_netbious_name_server,json=serverProfileNetbiousNameServer,proto3" json:"server_profile_netbious_name_server,omitempty"`
	ServerProfileTimeServer          []string `protobuf:"bytes,73,rep,name=server_profile_time_server,json=serverProfileTimeServer,proto3" json:"server_profile_time_server,omitempty"`
	LeaseLimitType                   uint32   `protobuf:"varint,74,opt,name=lease_limit_type,json=leaseLimitType,proto3" json:"lease_limit_type,omitempty"`
	LeaseLimitCount                  uint32   `protobuf:"varint,75,opt,name=lease_limit_count,json=leaseLimitCount,proto3" json:"lease_limit_count,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *Ipv4DhcpdServerProfile) Reset()         { *m = Ipv4DhcpdServerProfile{} }
func (m *Ipv4DhcpdServerProfile) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdServerProfile) ProtoMessage()    {}
func (*Ipv4DhcpdServerProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_c75a5ddc1135fea3, []int{1}
}

func (m *Ipv4DhcpdServerProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdServerProfile.Unmarshal(m, b)
}
func (m *Ipv4DhcpdServerProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdServerProfile.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdServerProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdServerProfile.Merge(m, src)
}
func (m *Ipv4DhcpdServerProfile) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdServerProfile.Size(m)
}
func (m *Ipv4DhcpdServerProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdServerProfile.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdServerProfile proto.InternalMessageInfo

func (m *Ipv4DhcpdServerProfile) GetServerProfileNameXr() string {
	if m != nil {
		return m.ServerProfileNameXr
	}
	return ""
}

func (m *Ipv4DhcpdServerProfile) GetSecureArp() bool {
	if m != nil {
		return m.SecureArp
	}
	return false
}

func (m *Ipv4DhcpdServerProfile) GetRequestedAddressCheck() bool {
	if m != nil {
		return m.RequestedAddressCheck
	}
	return false
}

func (m *Ipv4DhcpdServerProfile) GetServerIdCheck() bool {
	if m != nil {
		return m.ServerIdCheck
	}
	return false
}

func (m *Ipv4DhcpdServerProfile) GetDuplicateMacAddressCheck() bool {
	if m != nil {
		return m.DuplicateMacAddressCheck
	}
	return false
}

func (m *Ipv4DhcpdServerProfile) GetDuplicateIpAddressCheck() bool {
	if m != nil {
		return m.DuplicateIpAddressCheck
	}
	return false
}

func (m *Ipv4DhcpdServerProfile) GetIsMoveAllowed() bool {
	if m != nil {
		return m.IsMoveAllowed
	}
	return false
}

func (m *Ipv4DhcpdServerProfile) GetBcastPolicy() uint32 {
	if m != nil {
		return m.BcastPolicy
	}
	return 0
}

func (m *Ipv4DhcpdServerProfile) GetGiaddrPolicy() uint32 {
	if m != nil {
		return m.GiaddrPolicy
	}
	return 0
}

func (m *Ipv4DhcpdServerProfile) GetSubnetMask() string {
	if m != nil {
		return m.SubnetMask
	}
	return ""
}

func (m *Ipv4DhcpdServerProfile) GetServerPoolName() string {
	if m != nil {
		return m.ServerPoolName
	}
	return ""
}

func (m *Ipv4DhcpdServerProfile) GetServerProfileLease() uint32 {
	if m != nil {
		return m.ServerProfileLease
	}
	return 0
}

func (m *Ipv4DhcpdServerProfile) GetServerProfileNetbiosNodeType() uint32 {
	if m != nil {
		return m.ServerProfileNetbiosNodeType
	}
	return 0
}

func (m *Ipv4DhcpdServerProfile) GetServerBootfileName() string {
	if m != nil {
		return m.ServerBootfileName
	}
	return ""
}

func (m *Ipv4DhcpdServerProfile) GetServerDomainName() string {
	if m != nil {
		return m.ServerDomainName
	}
	return ""
}

func (m *Ipv4DhcpdServerProfile) GetServerProfileiedgeCheck() uint32 {
	if m != nil {
		return m.ServerProfileiedgeCheck
	}
	return 0
}

func (m *Ipv4DhcpdServerProfile) GetServerProfileServerDnsCount() uint32 {
	if m != nil {
		return m.ServerProfileServerDnsCount
	}
	return 0
}

func (m *Ipv4DhcpdServerProfile) GetServerProfiledefaultRouterCount() uint32 {
	if m != nil {
		return m.ServerProfiledefaultRouterCount
	}
	return 0
}

func (m *Ipv4DhcpdServerProfile) GetServerProfileNetbiosNameSvrCount() uint32 {
	if m != nil {
		return m.ServerProfileNetbiosNameSvrCount
	}
	return 0
}

func (m *Ipv4DhcpdServerProfile) GetServerProfileTimeSvrCount() uint32 {
	if m != nil {
		return m.ServerProfileTimeSvrCount
	}
	return 0
}

func (m *Ipv4DhcpdServerProfile) GetServerProfileDns() []string {
	if m != nil {
		return m.ServerProfileDns
	}
	return nil
}

func (m *Ipv4DhcpdServerProfile) GetServerProfileDefaultRouter() []string {
	if m != nil {
		return m.ServerProfileDefaultRouter
	}
	return nil
}

func (m *Ipv4DhcpdServerProfile) GetServerProfileNetbiousNameServer() []string {
	if m != nil {
		return m.ServerProfileNetbiousNameServer
	}
	return nil
}

func (m *Ipv4DhcpdServerProfile) GetServerProfileTimeServer() []string {
	if m != nil {
		return m.ServerProfileTimeServer
	}
	return nil
}

func (m *Ipv4DhcpdServerProfile) GetLeaseLimitType() uint32 {
	if m != nil {
		return m.LeaseLimitType
	}
	return 0
}

func (m *Ipv4DhcpdServerProfile) GetLeaseLimitCount() uint32 {
	if m != nil {
		return m.LeaseLimitCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4DhcpdServerProfile_KEYS)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.server.profiles.profile.ipv4_dhcpd_server_profile_KEYS")
	proto.RegisterType((*Ipv4DhcpdServerProfile)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.server.profiles.profile.ipv4_dhcpd_server_profile")
}

func init() { proto.RegisterFile("ipv4_dhcpd_server_profile.proto", fileDescriptor_c75a5ddc1135fea3) }

var fileDescriptor_c75a5ddc1135fea3 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0x6d, 0x4f, 0xdb, 0x48,
	0x10, 0xc7, 0x95, 0x3b, 0x09, 0x1d, 0xcb, 0xf3, 0x72, 0x47, 0x16, 0x38, 0x2e, 0x39, 0xd0, 0x9d,
	0xa2, 0xd3, 0x29, 0xaa, 0x0a, 0xa5, 0x2d, 0x94, 0x96, 0x40, 0xa0, 0xe5, 0xb1, 0x28, 0xf0, 0xa2,
	0x7d, 0xb5, 0xda, 0x78, 0x07, 0x58, 0x61, 0x7b, 0xdd, 0x5d, 0x3b, 0x85, 0xef, 0xd1, 0x0f, 0x5c,
	0x79, 0x76, 0x43, 0x62, 0x13, 0xde, 0x20, 0x3c, 0xf3, 0x9f, 0xdf, 0x8c, 0xff, 0x33, 0x56, 0x48,
	0x4d, 0x25, 0xbd, 0x0d, 0x2e, 0x6f, 0x83, 0x44, 0x72, 0x0b, 0xa6, 0x07, 0x86, 0x27, 0x46, 0x5f,
	0xab, 0x10, 0x9a, 0x89, 0xd1, 0xa9, 0xa6, 0xc7, 0x81, 0xb2, 0x81, 0xe6, 0x4a, 0x5b, 0x7e, 0x6f,
	0xf8, 0x90, 0x5a, 0x27, 0x60, 0x9a, 0x83, 0xe7, 0x66, 0xac, 0x25, 0x58, 0xfc, 0xdb, 0x74, 0xa0,
	0xa6, 0x07, 0xd9, 0xfe, 0x3f, 0xab, 0xb7, 0xe4, 0xaf, 0x67, 0xdb, 0xf1, 0x93, 0x83, 0xaf, 0x97,
	0x74, 0x81, 0x8c, 0xe5, 0x04, 0x25, 0x59, 0xa5, 0x5e, 0x69, 0x8c, 0x77, 0xfc, 0x13, 0x6d, 0x92,
	0xf9, 0x92, 0x3c, 0x16, 0x11, 0xb0, 0x5f, 0x50, 0x34, 0xe7, 0x52, 0x17, 0x2e, 0x73, 0x2e, 0x22,
	0x58, 0xfd, 0x31, 0x41, 0x16, 0x9f, 0x6d, 0x45, 0xd7, 0xc9, 0xc2, 0x08, 0x1a, 0xbf, 0x37, 0xec,
	0x25, 0x02, 0xe7, 0x9f, 0x00, 0xbf, 0x18, 0xba, 0x42, 0x88, 0x85, 0x20, 0x33, 0xc0, 0x85, 0x49,
	0xd8, 0x7a, 0xbd, 0xd2, 0xf8, 0xad, 0x33, 0xee, 0x22, 0x2d, 0x93, 0xd0, 0x4d, 0x52, 0x35, 0xf0,
	0x2d, 0x03, 0x9b, 0x82, 0xe4, 0x42, 0x4a, 0x03, 0xd6, 0xf2, 0xe0, 0x16, 0x82, 0x3b, 0xb6, 0x81,
	0xda, 0x3f, 0x1e, 0xd3, 0x2d, 0x97, 0xdd, 0xcf, 0x93, 0xf4, 0x5f, 0x32, 0xe3, 0x67, 0x51, 0xd2,
	0xeb, 0x5f, 0xa1, 0x7e, 0xca, 0x85, 0x8f, 0xa4, 0xd3, 0xed, 0x90, 0x65, 0x99, 0x25, 0xa1, 0x0a,
	0x44, 0x0a, 0x3c, 0x12, 0x41, 0xa9, 0xc7, 0x26, 0xd6, 0xb0, 0x47, 0xc9, 0x99, 0x08, 0x0a, 0x6d,
	0xb6, 0xc9, 0xd2, 0xa0, 0x5c, 0x25, 0xa5, 0xea, 0xd7, 0x58, 0x5d, 0x7d, 0x54, 0x1c, 0x25, 0xe5,
	0x19, 0x95, 0xe5, 0x91, 0xee, 0x01, 0x17, 0x61, 0xa8, 0xbf, 0x83, 0x64, 0x6f, 0xdc, 0x8c, 0xca,
	0x9e, 0xe9, 0x1e, 0xb4, 0x5c, 0x90, 0xfe, 0x4d, 0x26, 0xbb, 0x81, 0xb0, 0x29, 0x4f, 0x74, 0xa8,
	0x82, 0x07, 0xf6, 0xb6, 0x5e, 0x69, 0x4c, 0x75, 0x26, 0x30, 0x76, 0x81, 0x21, 0xba, 0x46, 0xa6,
	0x6e, 0x54, 0xde, 0xbc, 0xaf, 0xd9, 0x42, 0xcd, 0xa4, 0x0b, 0x7a, 0x51, 0x8d, 0x4c, 0xd8, 0xac,
	0x1b, 0x43, 0xca, 0x23, 0x61, 0xef, 0xd8, 0x36, 0x2e, 0x85, 0xb8, 0xd0, 0x99, 0xb0, 0x77, 0xb4,
	0x41, 0x66, 0xfb, 0x0b, 0xd4, 0x3a, 0x74, 0xb7, 0xf0, 0x0e, 0x55, 0xd3, 0x7e, 0x75, 0x5a, 0x87,
	0xf9, 0xde, 0xe8, 0x0b, 0xf2, 0x7b, 0x69, 0xd5, 0x21, 0x08, 0x0b, 0x6c, 0x07, 0xdb, 0xd2, 0xc2,
	0xa2, 0x4f, 0xf3, 0x0c, 0x3d, 0x24, 0xf5, 0xf2, 0x71, 0x40, 0xda, 0xcd, 0xef, 0x3f, 0xbf, 0x45,
	0x9e, 0x3e, 0x24, 0xc0, 0xde, 0x63, 0xf5, 0x9f, 0xc5, 0x33, 0x71, 0xaa, 0x73, 0x2d, 0xe1, 0xea,
	0x21, 0x19, 0xee, 0xdc, 0xd5, 0x3a, 0x1d, 0xdc, 0xec, 0x07, 0x9c, 0xd3, 0x77, 0xde, 0xf3, 0x29,
	0x9c, 0xf5, 0x7f, 0xe2, 0xa3, 0x5c, 0xea, 0x48, 0xa8, 0xd8, 0xe9, 0x77, 0x51, 0xef, 0xdf, 0xb7,
	0x8d, 0x09, 0x54, 0x6f, 0x91, 0xc5, 0xe2, 0x9c, 0x0a, 0xe4, 0x0d, 0xf8, 0x85, 0xb6, 0x70, 0xc0,
	0x6a, 0x61, 0x40, 0xcc, 0xbb, 0x85, 0xb6, 0x49, 0xad, 0xf4, 0x8e, 0xfd, 0xc6, 0xb1, 0xe5, 0x81,
	0xce, 0xe2, 0x94, 0xed, 0x21, 0x61, 0xb9, 0x40, 0xb8, 0x74, 0x33, 0xc4, 0x76, 0x3f, 0x97, 0xd0,
	0x13, 0xb2, 0x5a, 0xa4, 0x48, 0xb8, 0x16, 0x59, 0x98, 0x72, 0xa3, 0xb3, 0x14, 0x8c, 0x07, 0xed,
	0x23, 0xa8, 0x56, 0x00, 0x79, 0x61, 0x07, 0x75, 0x0e, 0xf6, 0x99, 0xfc, 0xf3, 0x9c, 0xed, 0xf9,
	0xb7, 0x69, 0x7b, 0x7d, 0x5e, 0x1b, 0x79, 0xf5, 0x91, 0xde, 0x8b, 0x08, 0x2e, 0x7b, 0x1e, 0xb8,
	0x4b, 0x56, 0x4a, 0xc0, 0x54, 0x15, 0x40, 0x07, 0x08, 0x5a, 0x2c, 0x80, 0xae, 0xd4, 0x10, 0x61,
	0xb0, 0x8f, 0x3e, 0x41, 0xc6, 0x96, 0x1d, 0xd6, 0x7f, 0x1d, 0xec, 0xc3, 0x97, 0xb5, 0x63, 0x4b,
	0x5b, 0x4f, 0xfa, 0x15, 0xed, 0x60, 0x1f, 0xb1, 0x70, 0xa9, 0x58, 0x38, 0x6c, 0x04, 0x3d, 0x25,
	0x6b, 0x23, 0x3d, 0xc8, 0xfa, 0x26, 0x60, 0x92, 0x7d, 0x42, 0x50, 0x6d, 0x84, 0x03, 0x99, 0xb3,
	0x00, 0x13, 0xf9, 0x27, 0x3f, 0xd2, 0x00, 0x07, 0x39, 0x42, 0x48, 0xf5, 0xe9, 0xdb, 0xbb, 0xe2,
	0x06, 0x99, 0xc5, 0x0f, 0x85, 0x87, 0x2a, 0x52, 0xa9, 0xbb, 0xfa, 0x63, 0x34, 0x6c, 0x1a, 0xe3,
	0xa7, 0x79, 0x18, 0xef, 0xfc, 0x3f, 0x32, 0x37, 0xac, 0x74, 0xde, 0x9e, 0xa0, 0x74, 0x66, 0x20,
	0x45, 0x47, 0xbb, 0x63, 0xf8, 0x9b, 0xb2, 0xfe, 0x33, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x72, 0x3d,
	0x53, 0x76, 0x06, 0x00, 0x00,
}
