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
// source: isis_sh_te_tunnel.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_te_tunnels_te_tunnel

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

type IsisShTeTunnel_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	TopologyName         string   `protobuf:"bytes,4,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	Level                string   `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	SystemId             string   `protobuf:"bytes,6,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,7,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShTeTunnel_KEYS) Reset()         { *m = IsisShTeTunnel_KEYS{} }
func (m *IsisShTeTunnel_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShTeTunnel_KEYS) ProtoMessage()    {}
func (*IsisShTeTunnel_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da9393d0618374e, []int{0}
}

func (m *IsisShTeTunnel_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTeTunnel_KEYS.Unmarshal(m, b)
}
func (m *IsisShTeTunnel_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTeTunnel_KEYS.Marshal(b, m, deterministic)
}
func (m *IsisShTeTunnel_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTeTunnel_KEYS.Merge(m, src)
}
func (m *IsisShTeTunnel_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShTeTunnel_KEYS.Size(m)
}
func (m *IsisShTeTunnel_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTeTunnel_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTeTunnel_KEYS proto.InternalMessageInfo

func (m *IsisShTeTunnel_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IsisShTeTunnel struct {
	TeSystemId                string   `protobuf:"bytes,50,opt,name=te_system_id,json=teSystemId,proto3" json:"te_system_id,omitempty"`
	TeInterface               string   `protobuf:"bytes,51,opt,name=te_interface,json=teInterface,proto3" json:"te_interface,omitempty"`
	TeBandwidth               uint32   `protobuf:"varint,52,opt,name=te_bandwidth,json=teBandwidth,proto3" json:"te_bandwidth,omitempty"`
	TeigpMetric               int32    `protobuf:"zigzag32,53,opt,name=teigp_metric,json=teigpMetric,proto3" json:"teigp_metric,omitempty"`
	TeNextHopIpAddress        string   `protobuf:"bytes,54,opt,name=te_next_hop_ip_address,json=teNextHopIpAddress,proto3" json:"te_next_hop_ip_address,omitempty"`
	TeModeType                string   `protobuf:"bytes,55,opt,name=te_mode_type,json=teModeType,proto3" json:"te_mode_type,omitempty"`
	Teipv4FaEnabled           bool     `protobuf:"varint,56,opt,name=teipv4fa_enabled,json=teipv4faEnabled,proto3" json:"teipv4fa_enabled,omitempty"`
	Teipv6FaEnabled           bool     `protobuf:"varint,57,opt,name=teipv6fa_enabled,json=teipv6faEnabled,proto3" json:"teipv6fa_enabled,omitempty"`
	Teipv4AaEnabled           bool     `protobuf:"varint,58,opt,name=teipv4aa_enabled,json=teipv4aaEnabled,proto3" json:"teipv4aa_enabled,omitempty"`
	Teipv6AaEnabled           bool     `protobuf:"varint,59,opt,name=teipv6aa_enabled,json=teipv6aaEnabled,proto3" json:"teipv6aa_enabled,omitempty"`
	TeCheckpointObjectId      uint32   `protobuf:"varint,60,opt,name=te_checkpoint_object_id,json=teCheckpointObjectId,proto3" json:"te_checkpoint_object_id,omitempty"`
	TeSegmentRoutingEnabled   bool     `protobuf:"varint,61,opt,name=te_segment_routing_enabled,json=teSegmentRoutingEnabled,proto3" json:"te_segment_routing_enabled,omitempty"`
	TeSegmentRoutingStrictSpf bool     `protobuf:"varint,62,opt,name=te_segment_routing_strict_spf,json=teSegmentRoutingStrictSpf,proto3" json:"te_segment_routing_strict_spf,omitempty"`
	TeSegmentRoutingExclude   bool     `protobuf:"varint,63,opt,name=te_segment_routing_exclude,json=teSegmentRoutingExclude,proto3" json:"te_segment_routing_exclude,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *IsisShTeTunnel) Reset()         { *m = IsisShTeTunnel{} }
func (m *IsisShTeTunnel) String() string { return proto.CompactTextString(m) }
func (*IsisShTeTunnel) ProtoMessage()    {}
func (*IsisShTeTunnel) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da9393d0618374e, []int{1}
}

func (m *IsisShTeTunnel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTeTunnel.Unmarshal(m, b)
}
func (m *IsisShTeTunnel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTeTunnel.Marshal(b, m, deterministic)
}
func (m *IsisShTeTunnel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTeTunnel.Merge(m, src)
}
func (m *IsisShTeTunnel) XXX_Size() int {
	return xxx_messageInfo_IsisShTeTunnel.Size(m)
}
func (m *IsisShTeTunnel) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTeTunnel.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTeTunnel proto.InternalMessageInfo

func (m *IsisShTeTunnel) GetTeSystemId() string {
	if m != nil {
		return m.TeSystemId
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeInterface() string {
	if m != nil {
		return m.TeInterface
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeBandwidth() uint32 {
	if m != nil {
		return m.TeBandwidth
	}
	return 0
}

func (m *IsisShTeTunnel) GetTeigpMetric() int32 {
	if m != nil {
		return m.TeigpMetric
	}
	return 0
}

func (m *IsisShTeTunnel) GetTeNextHopIpAddress() string {
	if m != nil {
		return m.TeNextHopIpAddress
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeModeType() string {
	if m != nil {
		return m.TeModeType
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeipv4FaEnabled() bool {
	if m != nil {
		return m.Teipv4FaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeipv6FaEnabled() bool {
	if m != nil {
		return m.Teipv6FaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeipv4AaEnabled() bool {
	if m != nil {
		return m.Teipv4AaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeipv6AaEnabled() bool {
	if m != nil {
		return m.Teipv6AaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeCheckpointObjectId() uint32 {
	if m != nil {
		return m.TeCheckpointObjectId
	}
	return 0
}

func (m *IsisShTeTunnel) GetTeSegmentRoutingEnabled() bool {
	if m != nil {
		return m.TeSegmentRoutingEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeSegmentRoutingStrictSpf() bool {
	if m != nil {
		return m.TeSegmentRoutingStrictSpf
	}
	return false
}

func (m *IsisShTeTunnel) GetTeSegmentRoutingExclude() bool {
	if m != nil {
		return m.TeSegmentRoutingExclude
	}
	return false
}

func init() {
	proto.RegisterType((*IsisShTeTunnel_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_tunnels.te_tunnel.isis_sh_te_tunnel_KEYS")
	proto.RegisterType((*IsisShTeTunnel)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_tunnels.te_tunnel.isis_sh_te_tunnel")
}

func init() { proto.RegisterFile("isis_sh_te_tunnel.proto", fileDescriptor_6da9393d0618374e) }

var fileDescriptor_6da9393d0618374e = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x5f, 0x6f, 0x12, 0x41,
	0x14, 0xc5, 0xb3, 0x6a, 0x81, 0x8e, 0x45, 0xed, 0xa6, 0x29, 0x5b, 0x8d, 0x09, 0xb6, 0x31, 0xc1,
	0x17, 0x1e, 0x5a, 0x8a, 0x7f, 0xea, 0x7f, 0xd3, 0x44, 0x62, 0x5a, 0x13, 0xf0, 0xc5, 0xa7, 0x9b,
	0x61, 0xe7, 0x2e, 0x8c, 0xee, 0xce, 0x4c, 0x76, 0x2e, 0x15, 0xe2, 0xc7, 0xf0, 0x8b, 0xfa, 0x11,
	0xcc, 0xce, 0x32, 0x0b, 0x29, 0xf6, 0xed, 0xce, 0x39, 0xbf, 0x39, 0xe7, 0xc2, 0x24, 0xcb, 0x5a,
	0xd2, 0x4a, 0x0b, 0x76, 0x0a, 0x84, 0x40, 0x33, 0xa5, 0x30, 0xed, 0x9a, 0x5c, 0x93, 0x0e, 0x7f,
	0xc7, 0xd2, 0xc6, 0x1a, 0xa4, 0xb6, 0x30, 0xcf, 0x21, 0x4e, 0x95, 0x05, 0x87, 0x6a, 0x83, 0x79,
	0xb7, 0x98, 0xba, 0x52, 0x59, 0xe2, 0x2a, 0xc6, 0xd5, 0xd4, 0x25, 0x6d, 0x74, 0xaa, 0x27, 0x12,
	0xad, 0x1f, 0x17, 0xd5, 0x00, 0x29, 0x5e, 0x61, 0x6a, 0xaf, 0x9d, 0xbb, 0x55, 0xaf, 0x5d, 0x8d,
	0x87, 0x7f, 0x03, 0xb6, 0xbf, 0xb1, 0x18, 0x7c, 0x39, 0xff, 0x3e, 0x0a, 0x8f, 0x58, 0xd3, 0xd7,
	0x81, 0xe2, 0x19, 0x46, 0x41, 0x3b, 0xe8, 0x6c, 0x0f, 0x77, 0xbc, 0x78, 0xc9, 0x33, 0x0c, 0x5b,
	0xac, 0xce, 0x93, 0xd2, 0xbe, 0xe5, 0xec, 0x1a, 0x4f, 0x9c, 0x71, 0xc0, 0x1a, 0xd6, 0x3b, 0xb7,
	0x9d, 0x53, 0xb7, 0x4b, 0xeb, 0x88, 0x35, 0xab, 0xd5, 0x9c, 0x7f, 0xa7, 0x0c, 0xf6, 0xa2, 0x83,
	0xf6, 0xd8, 0x96, 0x5b, 0x3b, 0xda, 0x72, 0x66, 0x79, 0x08, 0x1f, 0xb1, 0x6d, 0xbb, 0xb0, 0x84,
	0x19, 0x48, 0x11, 0xd5, 0x9c, 0xd3, 0x28, 0x85, 0x81, 0x08, 0x9f, 0xb2, 0x7b, 0x52, 0x11, 0xe6,
	0x09, 0xf7, 0x1b, 0xd7, 0x1d, 0xd1, 0xac, 0xd4, 0x22, 0xf9, 0xf0, 0xcf, 0x16, 0xdb, 0xdd, 0xf8,
	0xc9, 0x61, 0x9b, 0xed, 0x10, 0xc2, 0x2a, 0xfc, 0xd8, 0x5d, 0x65, 0x84, 0x23, 0x1f, 0xff, 0xc4,
	0x11, 0x55, 0x56, 0x74, 0xe2, 0x88, 0xbb, 0x84, 0x03, 0x2f, 0x2d, 0x91, 0x31, 0x57, 0xe2, 0x97,
	0x14, 0x34, 0x8d, 0x7a, 0xed, 0xa0, 0xd3, 0x2c, 0x90, 0x8f, 0x5e, 0x2a, 0x11, 0x39, 0x31, 0x90,
	0x21, 0xe5, 0x32, 0x8e, 0x4e, 0xdb, 0x41, 0x67, 0xb7, 0x40, 0xe4, 0xc4, 0x5c, 0x38, 0x29, 0x3c,
	0x66, 0xfb, 0x84, 0xa0, 0x70, 0x4e, 0x30, 0xd5, 0x06, 0xa4, 0x01, 0x2e, 0x44, 0x8e, 0xd6, 0x46,
	0x7d, 0x57, 0x19, 0x12, 0x5e, 0xe2, 0x9c, 0x3e, 0x6b, 0x33, 0x30, 0x1f, 0x4a, 0x67, 0xb9, 0x7e,
	0xa6, 0x05, 0x02, 0x2d, 0x0c, 0x46, 0xcf, 0xfd, 0xfa, 0x17, 0x5a, 0xe0, 0xb7, 0x85, 0xc1, 0xf0,
	0x19, 0x7b, 0x40, 0x28, 0xcd, 0x55, 0x2f, 0xe1, 0x80, 0x8a, 0x8f, 0x53, 0x14, 0xd1, 0x8b, 0x76,
	0xd0, 0x69, 0x0c, 0xef, 0x7b, 0xfd, 0xbc, 0x94, 0x2b, 0xb4, 0xbf, 0x86, 0xbe, 0x5c, 0x43, 0xfb,
	0x9b, 0x68, 0x8f, 0xaf, 0xd0, 0x57, 0xeb, 0xa9, 0x7c, 0x33, 0x75, 0x0d, 0x3d, 0x5b, 0x4f, 0x5d,
	0xa1, 0xa7, 0xac, 0x45, 0x08, 0xf1, 0x14, 0xe3, 0x9f, 0x46, 0x4b, 0x45, 0xa0, 0xc7, 0x3f, 0x30,
	0xa6, 0xe2, 0x5d, 0x5e, 0xbb, 0xbf, 0x74, 0x8f, 0xf0, 0x53, 0xe5, 0x7e, 0x75, 0xe6, 0x40, 0x84,
	0x67, 0xec, 0x61, 0xf1, 0x86, 0x38, 0xc9, 0x50, 0x11, 0xe4, 0x7a, 0x46, 0x52, 0x4d, 0xaa, 0xae,
	0x37, 0xae, 0xab, 0x45, 0x38, 0x2a, 0x81, 0x61, 0xe9, 0xfb, 0xce, 0xf7, 0xec, 0xf1, 0x7f, 0x2e,
	0xdb, 0xe2, 0x45, 0x08, 0xac, 0x49, 0xa2, 0xb7, 0xee, 0xfe, 0xc1, 0xf5, 0xfb, 0x23, 0x47, 0x8c,
	0x4c, 0x72, 0x53, 0xfd, 0x3c, 0x4e, 0x67, 0x02, 0xa3, 0x77, 0x37, 0xd4, 0x97, 0xf6, 0xb8, 0xe6,
	0x3e, 0x06, 0x27, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0xaf, 0x81, 0x8f, 0x27, 0x04, 0x00,
	0x00,
}
