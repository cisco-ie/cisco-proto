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
// source: pim_gre_info_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_default_context_gre_gre_next_hops_gre_next_hop

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

type PimGreInfoBag_KEYS struct {
	DestinationAddress   string   `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimGreInfoBag_KEYS) Reset()         { *m = PimGreInfoBag_KEYS{} }
func (m *PimGreInfoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimGreInfoBag_KEYS) ProtoMessage()    {}
func (*PimGreInfoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddbb1f78214d771c, []int{0}
}

func (m *PimGreInfoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimGreInfoBag_KEYS.Unmarshal(m, b)
}
func (m *PimGreInfoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimGreInfoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimGreInfoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimGreInfoBag_KEYS.Merge(m, src)
}
func (m *PimGreInfoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimGreInfoBag_KEYS.Size(m)
}
func (m *PimGreInfoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimGreInfoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimGreInfoBag_KEYS proto.InternalMessageInfo

func (m *PimGreInfoBag_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

type PimAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimAddrtype) Reset()         { *m = PimAddrtype{} }
func (m *PimAddrtype) String() string { return proto.CompactTextString(m) }
func (*PimAddrtype) ProtoMessage()    {}
func (*PimAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddbb1f78214d771c, []int{1}
}

func (m *PimAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimAddrtype.Unmarshal(m, b)
}
func (m *PimAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimAddrtype.Marshal(b, m, deterministic)
}
func (m *PimAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimAddrtype.Merge(m, src)
}
func (m *PimAddrtype) XXX_Size() int {
	return xxx_messageInfo_PimAddrtype.Size(m)
}
func (m *PimAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_PimAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_PimAddrtype proto.InternalMessageInfo

func (m *PimAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PimAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *PimAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type PimGrePathBag struct {
	GreInterfaceName            string       `protobuf:"bytes,1,opt,name=gre_interface_name,json=greInterfaceName,proto3" json:"gre_interface_name,omitempty"`
	IsGreInterfaceDisabled      bool         `protobuf:"varint,2,opt,name=is_gre_interface_disabled,json=isGreInterfaceDisabled,proto3" json:"is_gre_interface_disabled,omitempty"`
	GreNeighbor                 *PimAddrtype `protobuf:"bytes,3,opt,name=gre_neighbor,json=greNeighbor,proto3" json:"gre_neighbor,omitempty"`
	IsViaLsm                    bool         `protobuf:"varint,4,opt,name=is_via_lsm,json=isViaLsm,proto3" json:"is_via_lsm,omitempty"`
	IsConnectorAttributePresent bool         `protobuf:"varint,5,opt,name=is_connector_attribute_present,json=isConnectorAttributePresent,proto3" json:"is_connector_attribute_present,omitempty"`
	ExtranetVrfName             string       `protobuf:"bytes,6,opt,name=extranet_vrf_name,json=extranetVrfName,proto3" json:"extranet_vrf_name,omitempty"`
	GreNextHop                  *PimAddrtype `protobuf:"bytes,7,opt,name=gre_next_hop,json=greNextHop,proto3" json:"gre_next_hop,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}     `json:"-"`
	XXX_unrecognized            []byte       `json:"-"`
	XXX_sizecache               int32        `json:"-"`
}

func (m *PimGrePathBag) Reset()         { *m = PimGrePathBag{} }
func (m *PimGrePathBag) String() string { return proto.CompactTextString(m) }
func (*PimGrePathBag) ProtoMessage()    {}
func (*PimGrePathBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddbb1f78214d771c, []int{2}
}

func (m *PimGrePathBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimGrePathBag.Unmarshal(m, b)
}
func (m *PimGrePathBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimGrePathBag.Marshal(b, m, deterministic)
}
func (m *PimGrePathBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimGrePathBag.Merge(m, src)
}
func (m *PimGrePathBag) XXX_Size() int {
	return xxx_messageInfo_PimGrePathBag.Size(m)
}
func (m *PimGrePathBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimGrePathBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimGrePathBag proto.InternalMessageInfo

func (m *PimGrePathBag) GetGreInterfaceName() string {
	if m != nil {
		return m.GreInterfaceName
	}
	return ""
}

func (m *PimGrePathBag) GetIsGreInterfaceDisabled() bool {
	if m != nil {
		return m.IsGreInterfaceDisabled
	}
	return false
}

func (m *PimGrePathBag) GetGreNeighbor() *PimAddrtype {
	if m != nil {
		return m.GreNeighbor
	}
	return nil
}

func (m *PimGrePathBag) GetIsViaLsm() bool {
	if m != nil {
		return m.IsViaLsm
	}
	return false
}

func (m *PimGrePathBag) GetIsConnectorAttributePresent() bool {
	if m != nil {
		return m.IsConnectorAttributePresent
	}
	return false
}

func (m *PimGrePathBag) GetExtranetVrfName() string {
	if m != nil {
		return m.ExtranetVrfName
	}
	return ""
}

func (m *PimGrePathBag) GetGreNextHop() *PimAddrtype {
	if m != nil {
		return m.GreNextHop
	}
	return nil
}

type PimGreInfoBag struct {
	RegisteredAddress    *PimAddrtype     `protobuf:"bytes,50,opt,name=registered_address,json=registeredAddress,proto3" json:"registered_address,omitempty"`
	Metric               uint32           `protobuf:"varint,51,opt,name=metric,proto3" json:"metric,omitempty"`
	MetricPreference     uint32           `protobuf:"varint,52,opt,name=metric_preference,json=metricPreference,proto3" json:"metric_preference,omitempty"`
	IsConnected          uint32           `protobuf:"varint,53,opt,name=is_connected,json=isConnected,proto3" json:"is_connected,omitempty"`
	GrePath              []*PimGrePathBag `protobuf:"bytes,54,rep,name=gre_path,json=grePath,proto3" json:"gre_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PimGreInfoBag) Reset()         { *m = PimGreInfoBag{} }
func (m *PimGreInfoBag) String() string { return proto.CompactTextString(m) }
func (*PimGreInfoBag) ProtoMessage()    {}
func (*PimGreInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddbb1f78214d771c, []int{3}
}

func (m *PimGreInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimGreInfoBag.Unmarshal(m, b)
}
func (m *PimGreInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimGreInfoBag.Marshal(b, m, deterministic)
}
func (m *PimGreInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimGreInfoBag.Merge(m, src)
}
func (m *PimGreInfoBag) XXX_Size() int {
	return xxx_messageInfo_PimGreInfoBag.Size(m)
}
func (m *PimGreInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimGreInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimGreInfoBag proto.InternalMessageInfo

func (m *PimGreInfoBag) GetRegisteredAddress() *PimAddrtype {
	if m != nil {
		return m.RegisteredAddress
	}
	return nil
}

func (m *PimGreInfoBag) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *PimGreInfoBag) GetMetricPreference() uint32 {
	if m != nil {
		return m.MetricPreference
	}
	return 0
}

func (m *PimGreInfoBag) GetIsConnected() uint32 {
	if m != nil {
		return m.IsConnected
	}
	return 0
}

func (m *PimGreInfoBag) GetGrePath() []*PimGrePathBag {
	if m != nil {
		return m.GrePath
	}
	return nil
}

func init() {
	proto.RegisterType((*PimGreInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.gre.gre_next_hops.gre_next_hop.pim_gre_info_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.gre.gre_next_hops.gre_next_hop.pim_addrtype")
	proto.RegisterType((*PimGrePathBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.gre.gre_next_hops.gre_next_hop.pim_gre_path_bag")
	proto.RegisterType((*PimGreInfoBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.gre.gre_next_hops.gre_next_hop.pim_gre_info_bag")
}

func init() { proto.RegisterFile("pim_gre_info_bag.proto", fileDescriptor_ddbb1f78214d771c) }

var fileDescriptor_ddbb1f78214d771c = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x65, 0x4d, 0x4d, 0xe2, 0x24, 0x62, 0x32, 0x62, 0x5c, 0x51, 0x24, 0xe6, 0x29, 0xa8, 0xac,
	0x90, 0xd6, 0x80, 0x8f, 0xa5, 0x8a, 0x15, 0xa5, 0x84, 0x08, 0x05, 0x7d, 0x19, 0x66, 0x77, 0xef,
	0x6e, 0x06, 0xb2, 0x33, 0xcb, 0x9d, 0xdb, 0x90, 0xfe, 0x02, 0x41, 0x7f, 0x80, 0x6f, 0xbe, 0xf9,
	0x3f, 0x65, 0x67, 0x3f, 0xba, 0xf5, 0xbd, 0x7d, 0x08, 0x64, 0xe6, 0x9c, 0x3b, 0xf7, 0xcc, 0x39,
	0x77, 0x96, 0x4d, 0x72, 0x95, 0x89, 0x14, 0x41, 0x28, 0x9d, 0x18, 0x11, 0xca, 0x34, 0xc8, 0xd1,
	0x90, 0xe1, 0xdf, 0x23, 0x65, 0x23, 0x23, 0x94, 0xb1, 0x62, 0x8f, 0x42, 0xe5, 0xbb, 0x23, 0x51,
	0x30, 0x4d, 0x0e, 0x18, 0xa8, 0x7c, 0xb7, 0x2c, 0x56, 0x81, 0x25, 0xa9, 0xe3, 0xf0, 0x32, 0x88,
	0x21, 0x91, 0x17, 0x5b, 0x12, 0x91, 0xd1, 0x04, 0x7b, 0x0a, 0x52, 0x84, 0xe2, 0x27, 0x34, 0xec,
	0x49, 0x6c, 0x4c, 0x6e, 0xaf, 0xad, 0x66, 0xa7, 0xec, 0xd1, 0xff, 0x5d, 0xc5, 0xe7, 0x0f, 0xdf,
	0xbe, 0xf2, 0x37, 0xec, 0x61, 0x0c, 0x96, 0x94, 0x96, 0xa4, 0x8c, 0x16, 0x32, 0x8e, 0x11, 0xac,
	0xf5, 0xbd, 0xa9, 0x37, 0xbf, 0xb7, 0xe6, 0x2d, 0xe8, 0xb8, 0x44, 0x66, 0x19, 0x1b, 0x16, 0x27,
	0x15, 0x44, 0xba, 0xcc, 0x81, 0x3f, 0x66, 0x3d, 0x99, 0x08, 0x2d, 0x33, 0xa8, 0x8a, 0xba, 0x32,
	0x39, 0x93, 0x19, 0xf0, 0x17, 0x6c, 0xe8, 0xee, 0x50, 0x1f, 0x79, 0xc7, 0xa1, 0x83, 0x62, 0xaf,
	0x3a, 0xab, 0xa2, 0x2c, 0x1b, 0x4a, 0xa7, 0xa1, 0x2c, 0xeb, 0x76, 0x7f, 0x0f, 0xd8, 0xa8, 0x56,
	0x9e, 0x4b, 0xda, 0x14, 0xca, 0xf9, 0x6b, 0xc6, 0xcb, 0x9b, 0x10, 0x60, 0x22, 0x23, 0x68, 0xb7,
	0x1f, 0xa5, 0x08, 0x9f, 0x6a, 0xc0, 0x09, 0x79, 0xc7, 0x9e, 0x28, 0x2b, 0xae, 0x17, 0xc4, 0xca,
	0xca, 0x70, 0x0b, 0xb1, 0x53, 0xd5, 0x5f, 0x4f, 0x94, 0xfd, 0xd8, 0x2a, 0x7b, 0x5f, 0xa1, 0xfc,
	0x97, 0xc7, 0x86, 0xa5, 0x8f, 0x2a, 0xdd, 0x84, 0x06, 0x9d, 0xc2, 0xc1, 0x62, 0x13, 0xdc, 0x5c,
	0x54, 0x41, 0xdb, 0xdd, 0xf5, 0x20, 0x45, 0x38, 0xab, 0x9a, 0xf3, 0x67, 0x8c, 0x29, 0x2b, 0x76,
	0x4a, 0x8a, 0xad, 0xcd, 0xfc, 0x03, 0xa7, 0xbc, 0xaf, 0xec, 0xb9, 0x92, 0x5f, 0x6c, 0xc6, 0x4f,
	0xd8, 0x73, 0x65, 0x8b, 0x26, 0x1a, 0x22, 0x32, 0x28, 0x24, 0x11, 0xaa, 0xf0, 0x82, 0x40, 0xe4,
	0x08, 0x16, 0x34, 0xf9, 0x77, 0x5d, 0xc5, 0x53, 0x65, 0x4f, 0x6a, 0xd2, 0x71, 0xcd, 0x59, 0x95,
	0x14, 0xfe, 0x92, 0x8d, 0x61, 0x4f, 0x28, 0x35, 0x90, 0xd8, 0x61, 0x95, 0x6b, 0xd7, 0x19, 0xfb,
	0xa0, 0x06, 0xce, 0xb1, 0x0c, 0xf8, 0x67, 0x63, 0x4e, 0xa9, 0xdc, 0xef, 0xdd, 0xb2, 0x39, 0xcc,
	0x99, 0xb3, 0xa7, 0x53, 0x93, 0xcf, 0xfe, 0x74, 0xae, 0xe6, 0xa4, 0x9e, 0x70, 0xfe, 0xdb, 0x63,
	0x1c, 0x21, 0x55, 0x96, 0x00, 0x21, 0x6e, 0xc6, 0x6c, 0x71, 0xcb, 0x3a, 0xc7, 0x57, 0x1a, 0xea,
	0xc9, 0x9f, 0xb0, 0x6e, 0x06, 0x84, 0x2a, 0xf2, 0x0f, 0xa7, 0xde, 0xfc, 0xfe, 0xba, 0x5a, 0xf1,
	0x57, 0x6c, 0x5c, 0xfe, 0x2b, 0x42, 0x4b, 0x00, 0x41, 0x47, 0xe0, 0x1f, 0x39, 0xca, 0xa8, 0x04,
	0x56, 0xcd, 0xbe, 0x7b, 0x3e, 0x4d, 0xe2, 0x10, 0xfb, 0x6f, 0x1d, 0x6f, 0xd0, 0xe4, 0x0b, 0x31,
	0xff, 0xe1, 0xb1, 0x7e, 0xfd, 0x74, 0xfc, 0xe5, 0xb4, 0x33, 0x1f, 0x2c, 0xb6, 0x37, 0x7d, 0xef,
	0xf6, 0x53, 0x5d, 0xf7, 0x52, 0x84, 0x95, 0xa4, 0x4d, 0xd8, 0x75, 0x1f, 0xb9, 0xc3, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x43, 0xcf, 0xe7, 0xee, 0xfe, 0x04, 0x00, 0x00,
}
