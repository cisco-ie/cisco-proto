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
// source: igmp_edm_summary_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_mld_active_vrfs_vrf_summary

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

type IgmpEdmSummaryBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmSummaryBag_KEYS) Reset()         { *m = IgmpEdmSummaryBag_KEYS{} }
func (m *IgmpEdmSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmSummaryBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_062021e9c297372e, []int{0}
}

func (m *IgmpEdmSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmSummaryBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmSummaryBag_KEYS.Size(m)
}
func (m *IgmpEdmSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmSummaryBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmSummaryBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type IgmpEdmIntfSummaryBag struct {
	InterfaceName                string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	GroupLimit                   uint32   `protobuf:"varint,2,opt,name=group_limit,json=groupLimit,proto3" json:"group_limit,omitempty"`
	GroupCount                   uint32   `protobuf:"varint,3,opt,name=group_count,json=groupCount,proto3" json:"group_count,omitempty"`
	ParentIfhandle               string   `protobuf:"bytes,4,opt,name=parent_ifhandle,json=parentIfhandle,proto3" json:"parent_ifhandle,omitempty"`
	OnOff                        bool     `protobuf:"varint,5,opt,name=on_off,json=onOff,proto3" json:"on_off,omitempty"`
	TimeSinceLastQueryInSeconds  uint32   `protobuf:"varint,6,opt,name=time_since_last_query_in_seconds,json=timeSinceLastQueryInSeconds,proto3" json:"time_since_last_query_in_seconds,omitempty"`
	TimeSinceLastReportInSeconds uint32   `protobuf:"varint,7,opt,name=time_since_last_report_in_seconds,json=timeSinceLastReportInSeconds,proto3" json:"time_since_last_report_in_seconds,omitempty"`
	RouterUptimeInSeconds        uint32   `protobuf:"varint,8,opt,name=router_uptime_in_seconds,json=routerUptimeInSeconds,proto3" json:"router_uptime_in_seconds,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *IgmpEdmIntfSummaryBag) Reset()         { *m = IgmpEdmIntfSummaryBag{} }
func (m *IgmpEdmIntfSummaryBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmIntfSummaryBag) ProtoMessage()    {}
func (*IgmpEdmIntfSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_062021e9c297372e, []int{1}
}

func (m *IgmpEdmIntfSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmIntfSummaryBag.Unmarshal(m, b)
}
func (m *IgmpEdmIntfSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmIntfSummaryBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmIntfSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmIntfSummaryBag.Merge(m, src)
}
func (m *IgmpEdmIntfSummaryBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmIntfSummaryBag.Size(m)
}
func (m *IgmpEdmIntfSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmIntfSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmIntfSummaryBag proto.InternalMessageInfo

func (m *IgmpEdmIntfSummaryBag) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *IgmpEdmIntfSummaryBag) GetGroupLimit() uint32 {
	if m != nil {
		return m.GroupLimit
	}
	return 0
}

func (m *IgmpEdmIntfSummaryBag) GetGroupCount() uint32 {
	if m != nil {
		return m.GroupCount
	}
	return 0
}

func (m *IgmpEdmIntfSummaryBag) GetParentIfhandle() string {
	if m != nil {
		return m.ParentIfhandle
	}
	return ""
}

func (m *IgmpEdmIntfSummaryBag) GetOnOff() bool {
	if m != nil {
		return m.OnOff
	}
	return false
}

func (m *IgmpEdmIntfSummaryBag) GetTimeSinceLastQueryInSeconds() uint32 {
	if m != nil {
		return m.TimeSinceLastQueryInSeconds
	}
	return 0
}

func (m *IgmpEdmIntfSummaryBag) GetTimeSinceLastReportInSeconds() uint32 {
	if m != nil {
		return m.TimeSinceLastReportInSeconds
	}
	return 0
}

func (m *IgmpEdmIntfSummaryBag) GetRouterUptimeInSeconds() uint32 {
	if m != nil {
		return m.RouterUptimeInSeconds
	}
	return 0
}

type IgmpEdmSummaryBag struct {
	Robustness             uint32                   `protobuf:"varint,50,opt,name=robustness,proto3" json:"robustness,omitempty"`
	GroupLimit             uint32                   `protobuf:"varint,51,opt,name=group_limit,json=groupLimit,proto3" json:"group_limit,omitempty"`
	GroupCount             uint32                   `protobuf:"varint,52,opt,name=group_count,json=groupCount,proto3" json:"group_count,omitempty"`
	IsDisabled             bool                     `protobuf:"varint,53,opt,name=is_disabled,json=isDisabled,proto3" json:"is_disabled,omitempty"`
	SupportedInterfaces    uint32                   `protobuf:"varint,54,opt,name=supported_interfaces,json=supportedInterfaces,proto3" json:"supported_interfaces,omitempty"`
	UnsupportedInterfaces  uint32                   `protobuf:"varint,55,opt,name=unsupported_interfaces,json=unsupportedInterfaces,proto3" json:"unsupported_interfaces,omitempty"`
	EnabledInterfaceCount  uint32                   `protobuf:"varint,56,opt,name=enabled_interface_count,json=enabledInterfaceCount,proto3" json:"enabled_interface_count,omitempty"`
	DisabledInterfaceCount uint32                   `protobuf:"varint,57,opt,name=disabled_interface_count,json=disabledInterfaceCount,proto3" json:"disabled_interface_count,omitempty"`
	TunnelMteConfigCount   uint32                   `protobuf:"varint,58,opt,name=tunnel_mte_config_count,json=tunnelMteConfigCount,proto3" json:"tunnel_mte_config_count,omitempty"`
	NodeLowMemory          bool                     `protobuf:"varint,59,opt,name=node_low_memory,json=nodeLowMemory,proto3" json:"node_low_memory,omitempty"`
	Interface              []*IgmpEdmIntfSummaryBag `protobuf:"bytes,60,rep,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                 `json:"-"`
	XXX_unrecognized       []byte                   `json:"-"`
	XXX_sizecache          int32                    `json:"-"`
}

func (m *IgmpEdmSummaryBag) Reset()         { *m = IgmpEdmSummaryBag{} }
func (m *IgmpEdmSummaryBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmSummaryBag) ProtoMessage()    {}
func (*IgmpEdmSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_062021e9c297372e, []int{2}
}

func (m *IgmpEdmSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmSummaryBag.Unmarshal(m, b)
}
func (m *IgmpEdmSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmSummaryBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmSummaryBag.Merge(m, src)
}
func (m *IgmpEdmSummaryBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmSummaryBag.Size(m)
}
func (m *IgmpEdmSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmSummaryBag proto.InternalMessageInfo

func (m *IgmpEdmSummaryBag) GetRobustness() uint32 {
	if m != nil {
		return m.Robustness
	}
	return 0
}

func (m *IgmpEdmSummaryBag) GetGroupLimit() uint32 {
	if m != nil {
		return m.GroupLimit
	}
	return 0
}

func (m *IgmpEdmSummaryBag) GetGroupCount() uint32 {
	if m != nil {
		return m.GroupCount
	}
	return 0
}

func (m *IgmpEdmSummaryBag) GetIsDisabled() bool {
	if m != nil {
		return m.IsDisabled
	}
	return false
}

func (m *IgmpEdmSummaryBag) GetSupportedInterfaces() uint32 {
	if m != nil {
		return m.SupportedInterfaces
	}
	return 0
}

func (m *IgmpEdmSummaryBag) GetUnsupportedInterfaces() uint32 {
	if m != nil {
		return m.UnsupportedInterfaces
	}
	return 0
}

func (m *IgmpEdmSummaryBag) GetEnabledInterfaceCount() uint32 {
	if m != nil {
		return m.EnabledInterfaceCount
	}
	return 0
}

func (m *IgmpEdmSummaryBag) GetDisabledInterfaceCount() uint32 {
	if m != nil {
		return m.DisabledInterfaceCount
	}
	return 0
}

func (m *IgmpEdmSummaryBag) GetTunnelMteConfigCount() uint32 {
	if m != nil {
		return m.TunnelMteConfigCount
	}
	return 0
}

func (m *IgmpEdmSummaryBag) GetNodeLowMemory() bool {
	if m != nil {
		return m.NodeLowMemory
	}
	return false
}

func (m *IgmpEdmSummaryBag) GetInterface() []*IgmpEdmIntfSummaryBag {
	if m != nil {
		return m.Interface
	}
	return nil
}

func init() {
	proto.RegisterType((*IgmpEdmSummaryBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.vrfs.vrf.summary.igmp_edm_summary_bag_KEYS")
	proto.RegisterType((*IgmpEdmIntfSummaryBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.vrfs.vrf.summary.igmp_edm_intf_summary_bag")
	proto.RegisterType((*IgmpEdmSummaryBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.vrfs.vrf.summary.igmp_edm_summary_bag")
}

func init() { proto.RegisterFile("igmp_edm_summary_bag.proto", fileDescriptor_062021e9c297372e) }

var fileDescriptor_062021e9c297372e = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5d, 0x6f, 0xd3, 0x3e,
	0x14, 0xc6, 0xd5, 0xff, 0xde, 0x5d, 0x75, 0x93, 0xfc, 0xef, 0x36, 0x0f, 0x10, 0x2b, 0x95, 0x80,
	0x5e, 0x45, 0x62, 0x5b, 0xdb, 0xf1, 0x72, 0x37, 0x26, 0x54, 0xd1, 0x81, 0x48, 0xc5, 0x05, 0x57,
	0x47, 0x69, 0x72, 0x52, 0x2c, 0xc5, 0x76, 0xb0, 0x9d, 0x8e, 0x7e, 0x08, 0xbe, 0x16, 0x9f, 0x0b,
	0xc5, 0xc9, 0xd2, 0xd0, 0x15, 0x21, 0x71, 0xd3, 0x8b, 0xe7, 0x79, 0x7e, 0x8f, 0x9d, 0x73, 0x9a,
	0x90, 0x07, 0x7c, 0x26, 0x52, 0xc0, 0x48, 0x80, 0xc9, 0x84, 0x08, 0xf4, 0x02, 0xa6, 0xc1, 0xcc,
	0x4b, 0xb5, 0xb2, 0x8a, 0x0e, 0x43, 0x6e, 0x42, 0x05, 0x5c, 0x19, 0xf8, 0xae, 0x81, 0xa7, 0xf3,
	0x0b, 0x70, 0x69, 0x95, 0xa2, 0xf6, 0x44, 0x12, 0x79, 0x41, 0x68, 0xf9, 0x1c, 0xbd, 0xb9, 0x8e,
	0x4d, 0xfe, 0xe3, 0x95, 0x15, 0xdd, 0x01, 0x39, 0x59, 0x57, 0x0b, 0xef, 0xaf, 0xbf, 0x4c, 0xe8,
	0x09, 0xd9, 0x9d, 0xeb, 0x18, 0x64, 0x20, 0x90, 0x35, 0x3a, 0x8d, 0xde, 0x9e, 0xbf, 0x33, 0xd7,
	0xf1, 0x87, 0x40, 0x60, 0xf7, 0xc7, 0x46, 0x0d, 0xe4, 0xd2, 0xc6, 0x75, 0x9a, 0x3e, 0x25, 0xfb,
	0x5c, 0x5a, 0xd4, 0x71, 0x10, 0x62, 0x1d, 0x6f, 0x55, 0x6a, 0x5e, 0x42, 0x4f, 0x49, 0x73, 0xa6,
	0x55, 0x96, 0x42, 0xc2, 0x05, 0xb7, 0xec, 0xbf, 0x4e, 0xa3, 0xd7, 0xf2, 0x89, 0x93, 0xc6, 0xb9,
	0xb2, 0x0c, 0x84, 0x2a, 0x93, 0x96, 0x6d, 0xd4, 0x02, 0x57, 0xb9, 0x42, 0x9f, 0x93, 0x83, 0x34,
	0xd0, 0x28, 0x2d, 0xf0, 0xf8, 0x6b, 0x20, 0xa3, 0x04, 0xd9, 0xa6, 0x3b, 0x69, 0xbf, 0x90, 0x47,
	0xa5, 0x4a, 0x0f, 0xc9, 0xb6, 0x92, 0xa0, 0xe2, 0x98, 0x6d, 0x75, 0x1a, 0xbd, 0x5d, 0x7f, 0x4b,
	0xc9, 0x8f, 0x71, 0x4c, 0xaf, 0x49, 0xc7, 0x72, 0x81, 0x60, 0xb8, 0x0c, 0x11, 0x92, 0xc0, 0x58,
	0xf8, 0x96, 0xa1, 0x5e, 0x00, 0x97, 0x60, 0x30, 0x54, 0x32, 0x32, 0x6c, 0xdb, 0x9d, 0xfa, 0x30,
	0xcf, 0x4d, 0xf2, 0xd8, 0x38, 0x30, 0xf6, 0x53, 0x1e, 0x1a, 0xc9, 0x49, 0x11, 0xa1, 0xef, 0xc8,
	0x93, 0xd5, 0x1a, 0x8d, 0xa9, 0xd2, 0xb6, 0xde, 0xb3, 0xe3, 0x7a, 0x1e, 0xfd, 0xd6, 0xe3, 0xbb,
	0xd4, 0xb2, 0x68, 0x48, 0x98, 0x56, 0x99, 0x45, 0x0d, 0x59, 0xea, 0x1a, 0x6b, 0xfc, 0xae, 0xe3,
	0x0f, 0x0b, 0xff, 0xb3, 0xb3, 0x2b, 0xb0, 0xfb, 0x73, 0x93, 0xb4, 0xd7, 0x2d, 0x92, 0x3e, 0x26,
	0x44, 0xab, 0x69, 0x66, 0xac, 0x44, 0x63, 0xd8, 0x59, 0x31, 0xc1, 0xa5, 0xb2, 0xba, 0x83, 0xf3,
	0xbf, 0xed, 0xe0, 0xe2, 0xde, 0x0e, 0x4e, 0x49, 0x93, 0x1b, 0x88, 0xb8, 0x09, 0xa6, 0x09, 0x46,
	0xac, 0xef, 0xe6, 0x4b, 0xb8, 0x79, 0x5b, 0x2a, 0xf4, 0x05, 0x69, 0x9b, 0x2c, 0xcd, 0x1f, 0x14,
	0x23, 0xa8, 0xfe, 0x01, 0x86, 0x0d, 0x5c, 0xd5, 0xff, 0x95, 0x37, 0xaa, 0x2c, 0xda, 0x27, 0x47,
	0x99, 0x5c, 0x0b, 0x0d, 0x8b, 0x29, 0xd4, 0xdc, 0x1a, 0x36, 0x20, 0xc7, 0x28, 0xdd, 0xa1, 0x4b,
	0xa4, 0xbc, 0xf7, 0x65, 0xc1, 0x95, 0x76, 0xc5, 0x14, 0x8f, 0x70, 0x49, 0xd8, 0xdd, 0xfd, 0xef,
	0x81, 0x2f, 0x1d, 0x78, 0x74, 0xe7, 0xaf, 0x90, 0x7d, 0x72, 0x6c, 0x33, 0x29, 0x31, 0x01, 0x61,
	0x73, 0x42, 0xc6, 0x7c, 0x56, 0x82, 0xaf, 0x1c, 0xd8, 0x2e, 0xec, 0x1b, 0x8b, 0x57, 0xce, 0x2c,
	0xb0, 0x67, 0xe4, 0x40, 0xaa, 0x08, 0x21, 0x51, 0xb7, 0x20, 0x50, 0x28, 0xbd, 0x60, 0xaf, 0xdd,
	0xdc, 0x5a, 0xb9, 0x3c, 0x56, 0xb7, 0x37, 0x4e, 0xa4, 0x29, 0xd9, 0xab, 0xee, 0xc3, 0xde, 0x74,
	0x36, 0x7a, 0xcd, 0x33, 0xdf, 0xfb, 0xc7, 0x77, 0xdd, 0xfb, 0xe3, 0xfb, 0xea, 0x2f, 0x0f, 0x99,
	0x6e, 0xbb, 0x0f, 0xca, 0xf9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xb0, 0x72, 0xa9, 0x6e,
	0x04, 0x00, 0x00,
}
