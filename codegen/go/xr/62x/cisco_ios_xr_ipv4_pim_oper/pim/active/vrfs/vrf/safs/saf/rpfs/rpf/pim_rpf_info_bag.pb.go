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

package cisco_ios_xr_ipv4_pim_oper_pim_active_vrfs_vrf_safs_saf_rpfs_rpf

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	SafName              string   `protobuf:"bytes,2,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	TopologyName         string   `protobuf:"bytes,3,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	RegisteredAddress    string   `protobuf:"bytes,4,opt,name=registered_address,json=registeredAddress,proto3" json:"registered_address,omitempty"`
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

func (m *PimRpfInfoBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*PimRpfInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.safs.saf.rpfs.rpf.pim_rpf_info_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.safs.saf.rpfs.rpf.pim_addrtype")
	proto.RegisterType((*PimRpfPathBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.safs.saf.rpfs.rpf.pim_rpf_path_bag")
	proto.RegisterType((*PimRpfInfoBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.safs.saf.rpfs.rpf.pim_rpf_info_bag")
}

func init() { proto.RegisterFile("pim_rpf_info_bag.proto", fileDescriptor_d1bf566a024695a8) }

var fileDescriptor_d1bf566a024695a8 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xc7, 0x95, 0x27, 0x7d, 0xf2, 0x32, 0x9b, 0xe8, 0x49, 0xfd, 0xa8, 0x65, 0x2b, 0xaa, 0x2a,
	0x94, 0x03, 0x11, 0x2f, 0x7b, 0x68, 0x4b, 0x11, 0x37, 0x4a, 0xe1, 0x80, 0x80, 0xaa, 0x5a, 0xa4,
	0x0a, 0x4e, 0x96, 0xb3, 0x3b, 0x9b, 0x58, 0xca, 0xae, 0xcd, 0xd8, 0x8d, 0xd2, 0x2b, 0x5f, 0x82,
	0x33, 0x77, 0x3e, 0x24, 0xb2, 0xf7, 0xa5, 0x29, 0x3d, 0x42, 0x2f, 0x56, 0x32, 0xbf, 0xbf, 0x3d,
	0xa3, 0xff, 0xcc, 0x2c, 0x6c, 0x6b, 0x99, 0x73, 0xd2, 0x19, 0x97, 0x45, 0xa6, 0xf8, 0x54, 0xcc,
	0x22, 0x4d, 0xca, 0x2a, 0xf6, 0x2a, 0x91, 0x26, 0x51, 0x5c, 0x2a, 0xc3, 0x57, 0xc4, 0xa5, 0x5e,
	0x1e, 0x71, 0xa7, 0x54, 0x1a, 0x29, 0xd2, 0x32, 0x8f, 0x44, 0x62, 0xe5, 0x12, 0xa3, 0x25, 0x65,
	0xc6, 0x1d, 0x91, 0x11, 0x99, 0x71, 0x47, 0x44, 0x3a, 0x33, 0xee, 0xd8, 0xff, 0xd1, 0x82, 0xad,
	0xdf, 0x1f, 0xe7, 0xef, 0xdf, 0x7e, 0xf9, 0xc4, 0x76, 0xa0, 0xb7, 0xa4, 0x8c, 0x17, 0x22, 0xc7,
	0xb0, 0x35, 0x6e, 0x4d, 0xfa, 0x71, 0x77, 0x49, 0xd9, 0x99, 0xc8, 0xd1, 0x21, 0x23, 0x2a, 0xf4,
	0x4f, 0x89, 0x8c, 0x28, 0xd1, 0x43, 0x18, 0x5a, 0xa5, 0xd5, 0x42, 0xcd, 0xae, 0x4a, 0xde, 0xf6,
	0x7c, 0x50, 0x07, 0xbd, 0xe8, 0x19, 0x30, 0xc2, 0x99, 0x34, 0x16, 0x09, 0x53, 0x2e, 0xd2, 0x94,
	0xd0, 0x98, 0x70, 0xc3, 0x2b, 0x37, 0xaf, 0xc9, 0x49, 0x09, 0xf6, 0x73, 0x18, 0xb8, 0x12, 0x9d,
	0xce, 0x5e, 0x69, 0x64, 0xf7, 0xa0, 0x2b, 0x6e, 0x14, 0xd6, 0xa9, 0x92, 0x3f, 0x80, 0x81, 0xf7,
	0xa0, 0x7e, 0xb1, 0xac, 0x2d, 0x70, 0xb1, 0xea, 0xad, 0x4a, 0x72, 0xdc, 0x48, 0xda, 0x8d, 0xe4,
	0xb8, 0x4e, 0xf7, 0x73, 0x03, 0x46, 0xb5, 0x25, 0x5a, 0xd8, 0xb9, 0xb3, 0x84, 0x3d, 0x05, 0x56,
	0x5a, 0x64, 0x91, 0x32, 0x91, 0xe0, 0x7a, 0xfa, 0x11, 0xe9, 0xec, 0x5d, 0x0d, 0x7c, 0x21, 0x2f,
	0x61, 0x47, 0x1a, 0x7e, 0xf3, 0x42, 0x2a, 0x8d, 0x98, 0x2e, 0x30, 0xf5, 0x55, 0xf5, 0xe2, 0x6d,
	0x69, 0xe2, 0xb5, 0x6b, 0x6f, 0x2a, 0xca, 0xbe, 0xc2, 0xc0, 0xdd, 0x2b, 0x50, 0xce, 0xe6, 0x53,
	0x45, 0xbe, 0xc0, 0xe0, 0xe0, 0x2c, 0xfa, 0xd3, 0x4e, 0x47, 0xeb, 0x16, 0xc6, 0x01, 0xe9, 0xec,
	0xac, 0x4a, 0xc1, 0x76, 0x01, 0xa4, 0xe1, 0x4b, 0x29, 0xf8, 0xc2, 0xe4, 0xbe, 0x0d, 0xbd, 0xb8,
	0x27, 0xcd, 0x85, 0x14, 0x1f, 0x4c, 0xce, 0xf6, 0x20, 0xa8, 0x68, 0xee, 0xf0, 0xbf, 0x1e, 0xf7,
	0x3d, 0xfe, 0xb8, 0x30, 0x39, 0x3b, 0x85, 0x3d, 0x69, 0x78, 0xa2, 0x8a, 0x02, 0x13, 0xab, 0x88,
	0x0b, 0x6b, 0x49, 0x4e, 0x2f, 0x2d, 0x72, 0x4d, 0x68, 0xb0, 0xb0, 0x61, 0xc7, 0x5f, 0xb9, 0x2f,
	0xcd, 0x69, 0x2d, 0x3a, 0xa9, 0x35, 0xe7, 0xa5, 0x84, 0xed, 0x42, 0xbf, 0x79, 0x21, 0xec, 0x8e,
	0xdb, 0x93, 0x61, 0x7c, 0x1d, 0x60, 0x8f, 0x61, 0x13, 0x57, 0x96, 0x44, 0x81, 0x96, 0x37, 0x33,
	0xd9, 0xf3, 0xde, 0xff, 0x57, 0x83, 0x8b, 0x6a, 0x36, 0x15, 0x04, 0xa5, 0x7f, 0x2b, 0x3b, 0x57,
	0x3a, 0xec, 0xdf, 0x89, 0x7d, 0xe0, 0xed, 0xf3, 0x19, 0xf6, 0xbf, 0xb7, 0xaf, 0xc7, 0xa5, 0xde,
	0x20, 0xf6, 0xad, 0x05, 0x5b, 0xb7, 0x47, 0x9c, 0xaf, 0x28, 0x3c, 0xb8, 0x93, 0x82, 0xfe, 0xbf,
	0xb5, 0x35, 0x9f, 0x89, 0x6d, 0x43, 0x27, 0x47, 0x4b, 0x32, 0x09, 0x0f, 0xc7, 0xad, 0xc9, 0x30,
	0xae, 0xfe, 0xb1, 0x27, 0xb0, 0x59, 0xfe, 0x72, 0x1d, 0xca, 0x90, 0xb0, 0x48, 0x30, 0x3c, 0xf2,
	0x92, 0x51, 0x09, 0xce, 0x9b, 0xb8, 0x5f, 0x98, 0xa6, 0xbd, 0x98, 0x86, 0xcf, 0xbd, 0x2e, 0x68,
	0x9a, 0x89, 0x29, 0x7b, 0x04, 0xa3, 0x6a, 0xda, 0xa7, 0x33, 0xcd, 0x49, 0x5d, 0x5a, 0x0c, 0x8f,
	0x7d, 0xcf, 0x87, 0x7e, 0xc8, 0x5f, 0xcf, 0x74, 0xec, 0x82, 0x2c, 0x87, 0x5e, 0xbd, 0x54, 0xe1,
	0x8b, 0x71, 0x7b, 0x12, 0x1c, 0xc4, 0x7f, 0xc7, 0x87, 0xf5, 0x55, 0x8d, 0xbb, 0xa4, 0xb3, 0x73,
	0x61, 0xe7, 0xd3, 0x8e, 0xff, 0x48, 0x1e, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x90, 0xe6, 0x40,
	0x06, 0x3e, 0x05, 0x00, 0x00,
}
