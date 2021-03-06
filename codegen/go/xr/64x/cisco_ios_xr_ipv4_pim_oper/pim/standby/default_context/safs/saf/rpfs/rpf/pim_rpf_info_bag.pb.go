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
// source: pim_rpf_info_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_safs_saf_rpfs_rpf

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

type PimRpfInfoBag_KEYS struct {
	SafName              string   `protobuf:"bytes,1,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	TopologyName         string   `protobuf:"bytes,2,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	RegisteredAddress    string   `protobuf:"bytes,3,opt,name=registered_address,json=registeredAddress,proto3" json:"registered_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimRpfInfoBag_KEYS) Reset()         { *m = PimRpfInfoBag_KEYS{} }
func (m *PimRpfInfoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimRpfInfoBag_KEYS) ProtoMessage()    {}
func (*PimRpfInfoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1bf566a024695a8, []int{0}
}

func (m *PimRpfInfoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfInfoBag_KEYS.Unmarshal(m, b)
}
func (m *PimRpfInfoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfInfoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimRpfInfoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfInfoBag_KEYS.Merge(m, src)
}
func (m *PimRpfInfoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimRpfInfoBag_KEYS.Size(m)
}
func (m *PimRpfInfoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfInfoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfInfoBag_KEYS proto.InternalMessageInfo

func (m *PimRpfInfoBag_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *PimRpfInfoBag_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *PimRpfInfoBag_KEYS) GetRegisteredAddress() string {
	if m != nil {
		return m.RegisteredAddress
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
	return fileDescriptor_d1bf566a024695a8, []int{1}
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

type PimRpfPathBag struct {
	RpfInterfaceName            string       `protobuf:"bytes,1,opt,name=rpf_interface_name,json=rpfInterfaceName,proto3" json:"rpf_interface_name,omitempty"`
	IsRpfInterfaceDisabled      bool         `protobuf:"varint,2,opt,name=is_rpf_interface_disabled,json=isRpfInterfaceDisabled,proto3" json:"is_rpf_interface_disabled,omitempty"`
	RpfNeighbor                 *PimAddrtype `protobuf:"bytes,3,opt,name=rpf_neighbor,json=rpfNeighbor,proto3" json:"rpf_neighbor,omitempty"`
	IsViaLsm                    bool         `protobuf:"varint,4,opt,name=is_via_lsm,json=isViaLsm,proto3" json:"is_via_lsm,omitempty"`
	IsViaMlsm                   bool         `protobuf:"varint,5,opt,name=is_via_mlsm,json=isViaMlsm,proto3" json:"is_via_mlsm,omitempty"`
	IsConnectorAttributePresent bool         `protobuf:"varint,6,opt,name=is_connector_attribute_present,json=isConnectorAttributePresent,proto3" json:"is_connector_attribute_present,omitempty"`
	Connector                   []uint32     `protobuf:"varint,7,rep,packed,name=connector,proto3" json:"connector,omitempty"`
	ExtranetVrfName             string       `protobuf:"bytes,8,opt,name=extranet_vrf_name,json=extranetVrfName,proto3" json:"extranet_vrf_name,omitempty"`
	RpfNexthop                  *PimAddrtype `protobuf:"bytes,9,opt,name=rpf_nexthop,json=rpfNexthop,proto3" json:"rpf_nexthop,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}     `json:"-"`
	XXX_unrecognized            []byte       `json:"-"`
	XXX_sizecache               int32        `json:"-"`
}

func (m *PimRpfPathBag) Reset()         { *m = PimRpfPathBag{} }
func (m *PimRpfPathBag) String() string { return proto.CompactTextString(m) }
func (*PimRpfPathBag) ProtoMessage()    {}
func (*PimRpfPathBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1bf566a024695a8, []int{2}
}

func (m *PimRpfPathBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfPathBag.Unmarshal(m, b)
}
func (m *PimRpfPathBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfPathBag.Marshal(b, m, deterministic)
}
func (m *PimRpfPathBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfPathBag.Merge(m, src)
}
func (m *PimRpfPathBag) XXX_Size() int {
	return xxx_messageInfo_PimRpfPathBag.Size(m)
}
func (m *PimRpfPathBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfPathBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfPathBag proto.InternalMessageInfo

func (m *PimRpfPathBag) GetRpfInterfaceName() string {
	if m != nil {
		return m.RpfInterfaceName
	}
	return ""
}

func (m *PimRpfPathBag) GetIsRpfInterfaceDisabled() bool {
	if m != nil {
		return m.IsRpfInterfaceDisabled
	}
	return false
}

func (m *PimRpfPathBag) GetRpfNeighbor() *PimAddrtype {
	if m != nil {
		return m.RpfNeighbor
	}
	return nil
}

func (m *PimRpfPathBag) GetIsViaLsm() bool {
	if m != nil {
		return m.IsViaLsm
	}
	return false
}

func (m *PimRpfPathBag) GetIsViaMlsm() bool {
	if m != nil {
		return m.IsViaMlsm
	}
	return false
}

func (m *PimRpfPathBag) GetIsConnectorAttributePresent() bool {
	if m != nil {
		return m.IsConnectorAttributePresent
	}
	return false
}

func (m *PimRpfPathBag) GetConnector() []uint32 {
	if m != nil {
		return m.Connector
	}
	return nil
}

func (m *PimRpfPathBag) GetExtranetVrfName() string {
	if m != nil {
		return m.ExtranetVrfName
	}
	return ""
}

func (m *PimRpfPathBag) GetRpfNexthop() *PimAddrtype {
	if m != nil {
		return m.RpfNexthop
	}
	return nil
}

type PimRpfInfoBag struct {
	RegisteredAddressXr  *PimAddrtype     `protobuf:"bytes,50,opt,name=registered_address_xr,json=registeredAddressXr,proto3" json:"registered_address_xr,omitempty"`
	Metric               uint32           `protobuf:"varint,51,opt,name=metric,proto3" json:"metric,omitempty"`
	MetricPreference     uint32           `protobuf:"varint,52,opt,name=metric_preference,json=metricPreference,proto3" json:"metric_preference,omitempty"`
	IsConnected          uint32           `protobuf:"varint,53,opt,name=is_connected,json=isConnected,proto3" json:"is_connected,omitempty"`
	IsRpfBgpRoute        bool             `protobuf:"varint,54,opt,name=is_rpf_bgp_route,json=isRpfBgpRoute,proto3" json:"is_rpf_bgp_route,omitempty"`
	RpfPath              []*PimRpfPathBag `protobuf:"bytes,55,rep,name=rpf_path,json=rpfPath,proto3" json:"rpf_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PimRpfInfoBag) Reset()         { *m = PimRpfInfoBag{} }
func (m *PimRpfInfoBag) String() string { return proto.CompactTextString(m) }
func (*PimRpfInfoBag) ProtoMessage()    {}
func (*PimRpfInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1bf566a024695a8, []int{3}
}

func (m *PimRpfInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfInfoBag.Unmarshal(m, b)
}
func (m *PimRpfInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfInfoBag.Marshal(b, m, deterministic)
}
func (m *PimRpfInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfInfoBag.Merge(m, src)
}
func (m *PimRpfInfoBag) XXX_Size() int {
	return xxx_messageInfo_PimRpfInfoBag.Size(m)
}
func (m *PimRpfInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfInfoBag proto.InternalMessageInfo

func (m *PimRpfInfoBag) GetRegisteredAddressXr() *PimAddrtype {
	if m != nil {
		return m.RegisteredAddressXr
	}
	return nil
}

func (m *PimRpfInfoBag) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *PimRpfInfoBag) GetMetricPreference() uint32 {
	if m != nil {
		return m.MetricPreference
	}
	return 0
}

func (m *PimRpfInfoBag) GetIsConnected() uint32 {
	if m != nil {
		return m.IsConnected
	}
	return 0
}

func (m *PimRpfInfoBag) GetIsRpfBgpRoute() bool {
	if m != nil {
		return m.IsRpfBgpRoute
	}
	return false
}

func (m *PimRpfInfoBag) GetRpfPath() []*PimRpfPathBag {
	if m != nil {
		return m.RpfPath
	}
	return nil
}

func init() {
	proto.RegisterType((*PimRpfInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.safs.saf.rpfs.rpf.pim_rpf_info_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.safs.saf.rpfs.rpf.pim_addrtype")
	proto.RegisterType((*PimRpfPathBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.safs.saf.rpfs.rpf.pim_rpf_path_bag")
	proto.RegisterType((*PimRpfInfoBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.safs.saf.rpfs.rpf.pim_rpf_info_bag")
}

func init() { proto.RegisterFile("pim_rpf_info_bag.proto", fileDescriptor_d1bf566a024695a8) }

var fileDescriptor_d1bf566a024695a8 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0x48, 0xc9, 0xcf, 0x6c, 0x22, 0x52, 0xa3, 0x96, 0xad, 0xa8, 0xaa, 0x50, 0x0e, 0x44,
	0xfc, 0xec, 0xa1, 0x2d, 0x45, 0x1c, 0x4b, 0x41, 0x02, 0x01, 0x55, 0xb5, 0x48, 0x15, 0x70, 0xb1,
	0xbc, 0xbb, 0xe3, 0xc4, 0x52, 0x76, 0x6d, 0xd9, 0x4e, 0x49, 0x5e, 0x80, 0x03, 0x2f, 0xc2, 0x4b,
	0xf0, 0x70, 0xc8, 0xde, 0xdd, 0x24, 0x6d, 0xaf, 0xed, 0xc5, 0x4a, 0xe6, 0xfb, 0xc6, 0xdf, 0x78,
	0x66, 0xbe, 0x85, 0x6d, 0x25, 0x72, 0xaa, 0x15, 0xa7, 0xa2, 0xe0, 0x92, 0x26, 0x6c, 0x1c, 0x29,
	0x2d, 0xad, 0x24, 0x1f, 0x53, 0x61, 0x52, 0x49, 0x85, 0x34, 0x74, 0xae, 0xa9, 0x50, 0x97, 0x47,
	0xd4, 0x31, 0xa5, 0x42, 0x1d, 0x29, 0x91, 0x47, 0xc6, 0xb2, 0x22, 0x4b, 0x16, 0x51, 0x86, 0x9c,
	0xcd, 0xa6, 0x96, 0xa6, 0xb2, 0xb0, 0x38, 0xb7, 0x91, 0x61, 0xdc, 0xb8, 0x23, 0xd2, 0x8a, 0x1b,
	0x77, 0xec, 0xff, 0x6e, 0xc0, 0xd6, 0x75, 0x11, 0xfa, 0xf9, 0xc3, 0x8f, 0x6f, 0x64, 0x07, 0x3a,
	0x86, 0x71, 0x5a, 0xb0, 0x1c, 0xc3, 0xc6, 0xb0, 0x31, 0xea, 0xc6, 0x6d, 0xc3, 0xf8, 0x19, 0xcb,
	0x91, 0x3c, 0x85, 0xbe, 0x95, 0x4a, 0x4e, 0xe5, 0x78, 0x51, 0xe2, 0xf7, 0x3c, 0xde, 0xab, 0x83,
	0x9e, 0xf4, 0x0a, 0x88, 0xc6, 0xb1, 0x30, 0x16, 0x35, 0x66, 0x94, 0x65, 0x99, 0x46, 0x63, 0xc2,
	0xa6, 0x67, 0x6e, 0xae, 0x90, 0x93, 0x12, 0xd8, 0xcf, 0xa1, 0xe7, 0xea, 0x70, 0x3c, 0xbb, 0x50,
	0x48, 0x1e, 0x41, 0xfb, 0xaa, 0x7a, 0xab, 0x12, 0x7f, 0x02, 0x3d, 0xff, 0xe0, 0xfa, 0xc6, 0x52,
	0x3b, 0x70, 0xb1, 0xea, 0xae, 0x8a, 0x72, 0x7c, 0x4d, 0xd4, 0x51, 0x8e, 0x6b, 0xb9, 0x7f, 0x1b,
	0x30, 0xa8, 0xdf, 0xad, 0x98, 0x9d, 0xb8, 0x77, 0x93, 0x97, 0x40, 0xca, 0x3e, 0x58, 0xd4, 0x9c,
	0xa5, 0xb8, 0x2e, 0x3f, 0xd0, 0x8a, 0x7f, 0xaa, 0x01, 0x5f, 0xc8, 0x5b, 0xd8, 0x11, 0x86, 0x5e,
	0x4d, 0xc8, 0x84, 0x61, 0xc9, 0x14, 0x33, 0x5f, 0x55, 0x27, 0xde, 0x16, 0x26, 0x5e, 0x4b, 0x7b,
	0x5f, 0xa1, 0x64, 0x01, 0x3d, 0x97, 0x57, 0xa0, 0x18, 0x4f, 0x12, 0xa9, 0x7d, 0x81, 0xc1, 0xc1,
	0x45, 0x74, 0x5b, 0x63, 0x8d, 0xd6, 0x5b, 0x19, 0x07, 0x5a, 0xf1, 0xb3, 0x4a, 0x8a, 0xec, 0x02,
	0x08, 0x43, 0x2f, 0x05, 0xa3, 0x53, 0x93, 0x87, 0x1b, 0xbe, 0xcc, 0x8e, 0x30, 0x17, 0x82, 0x7d,
	0x31, 0x39, 0xd9, 0x83, 0xa0, 0x42, 0x73, 0x07, 0xdf, 0xf7, 0x70, 0xd7, 0xc3, 0x5f, 0xa7, 0x26,
	0x27, 0xa7, 0xb0, 0x27, 0x8c, 0xd3, 0x2d, 0x30, 0xb5, 0x52, 0x53, 0x66, 0xad, 0x16, 0xc9, 0xcc,
	0x22, 0x55, 0x1a, 0x0d, 0x16, 0x36, 0x6c, 0xf9, 0x94, 0xc7, 0xc2, 0x9c, 0xd6, 0xa4, 0x93, 0x9a,
	0x73, 0x5e, 0x52, 0xc8, 0x2e, 0x74, 0x97, 0x37, 0x84, 0xed, 0x61, 0x73, 0xd4, 0x8f, 0x57, 0x01,
	0xf2, 0x1c, 0x36, 0x71, 0x6e, 0x35, 0x2b, 0xd0, 0xd2, 0x4b, 0x5d, 0xad, 0x40, 0xc7, 0xcf, 0xe0,
	0x41, 0x0d, 0x5c, 0xe8, 0x72, 0x17, 0x7e, 0x41, 0x50, 0xf6, 0x71, 0x6e, 0x27, 0x52, 0x85, 0xdd,
	0x3b, 0x6d, 0x23, 0xf8, 0x36, 0x7a, 0xa5, 0xfd, 0xbf, 0xcd, 0xd5, 0xfa, 0xd4, 0xb6, 0x21, 0x7f,
	0x1a, 0xb0, 0x75, 0x73, 0xe5, 0xe9, 0x5c, 0x87, 0x07, 0x77, 0x5a, 0xd8, 0xc3, 0x1b, 0x6e, 0xfa,
	0xae, 0xc9, 0x36, 0xb4, 0x72, 0xb4, 0x5a, 0xa4, 0xe1, 0xe1, 0xb0, 0x31, 0xea, 0xc7, 0xd5, 0x3f,
	0xf2, 0x02, 0x36, 0xcb, 0x5f, 0x6e, 0x62, 0x1c, 0x35, 0x16, 0x29, 0x86, 0x47, 0x9e, 0x32, 0x28,
	0x81, 0xf3, 0x65, 0xdc, 0x1b, 0x69, 0x39, 0x6e, 0xcc, 0xc2, 0xd7, 0x9e, 0x17, 0x2c, 0x87, 0x8b,
	0x19, 0x79, 0x06, 0x83, 0xca, 0x05, 0xc9, 0x58, 0x51, 0x2d, 0x67, 0x16, 0xc3, 0x63, 0xbf, 0x03,
	0x7d, 0xbf, 0xfc, 0xef, 0xc6, 0x2a, 0x76, 0x41, 0x32, 0x83, 0x4e, 0x6d, 0xb6, 0xf0, 0xcd, 0xb0,
	0x39, 0x0a, 0x0e, 0x7e, 0xde, 0x6e, 0x3f, 0xd6, 0xad, 0x1c, 0xb7, 0xb5, 0xe2, 0xe7, 0xcc, 0x4e,
	0x92, 0x96, 0xff, 0x62, 0x1e, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x58, 0x81, 0x3e, 0x5f, 0x4b,
	0x05, 0x00, 0x00,
}
