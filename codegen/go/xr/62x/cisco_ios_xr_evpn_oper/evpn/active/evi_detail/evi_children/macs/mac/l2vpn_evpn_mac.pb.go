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
// source: l2vpn_evpn_mac.proto

package cisco_ios_xr_evpn_oper_evpn_active_evi_detail_evi_children_macs_mac

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

type L2VpnEvpnMac_KEYS struct {
	Evi                  uint32   `protobuf:"varint,1,opt,name=evi,proto3" json:"evi,omitempty"`
	EthernetTag          uint32   `protobuf:"varint,2,opt,name=ethernet_tag,json=ethernetTag,proto3" json:"ethernet_tag,omitempty"`
	MacAddress           string   `protobuf:"bytes,3,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	IpAddress            string   `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnMac_KEYS) Reset()         { *m = L2VpnEvpnMac_KEYS{} }
func (m *L2VpnEvpnMac_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnMac_KEYS) ProtoMessage()    {}
func (*L2VpnEvpnMac_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f7d8f00693a691d, []int{0}
}

func (m *L2VpnEvpnMac_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnMac_KEYS.Unmarshal(m, b)
}
func (m *L2VpnEvpnMac_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnMac_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnMac_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnMac_KEYS.Merge(m, src)
}
func (m *L2VpnEvpnMac_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnMac_KEYS.Size(m)
}
func (m *L2VpnEvpnMac_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnMac_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnMac_KEYS proto.InternalMessageInfo

func (m *L2VpnEvpnMac_KEYS) GetEvi() uint32 {
	if m != nil {
		return m.Evi
	}
	return 0
}

func (m *L2VpnEvpnMac_KEYS) GetEthernetTag() uint32 {
	if m != nil {
		return m.EthernetTag
	}
	return 0
}

func (m *L2VpnEvpnMac_KEYS) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *L2VpnEvpnMac_KEYS) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type L2VpnLabelPathBuffer struct {
	NextHop              string   `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	OutputLabel          uint32   `protobuf:"varint,2,opt,name=output_label,json=outputLabel,proto3" json:"output_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnLabelPathBuffer) Reset()         { *m = L2VpnLabelPathBuffer{} }
func (m *L2VpnLabelPathBuffer) String() string { return proto.CompactTextString(m) }
func (*L2VpnLabelPathBuffer) ProtoMessage()    {}
func (*L2VpnLabelPathBuffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f7d8f00693a691d, []int{1}
}

func (m *L2VpnLabelPathBuffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLabelPathBuffer.Unmarshal(m, b)
}
func (m *L2VpnLabelPathBuffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLabelPathBuffer.Marshal(b, m, deterministic)
}
func (m *L2VpnLabelPathBuffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLabelPathBuffer.Merge(m, src)
}
func (m *L2VpnLabelPathBuffer) XXX_Size() int {
	return xxx_messageInfo_L2VpnLabelPathBuffer.Size(m)
}
func (m *L2VpnLabelPathBuffer) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLabelPathBuffer.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLabelPathBuffer proto.InternalMessageInfo

func (m *L2VpnLabelPathBuffer) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *L2VpnLabelPathBuffer) GetOutputLabel() uint32 {
	if m != nil {
		return m.OutputLabel
	}
	return 0
}

type L2VpnEvpnMac struct {
	EthernetTagXr                   uint32                  `protobuf:"varint,50,opt,name=ethernet_tag_xr,json=ethernetTagXr,proto3" json:"ethernet_tag_xr,omitempty"`
	MacAddressXr                    string                  `protobuf:"bytes,51,opt,name=mac_address_xr,json=macAddressXr,proto3" json:"mac_address_xr,omitempty"`
	IpAddressXr                     string                  `protobuf:"bytes,52,opt,name=ip_address_xr,json=ipAddressXr,proto3" json:"ip_address_xr,omitempty"`
	LocalLabel                      uint32                  `protobuf:"varint,53,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	NumPaths                        uint32                  `protobuf:"varint,54,opt,name=num_paths,json=numPaths,proto3" json:"num_paths,omitempty"`
	PathBuffer                      []*L2VpnLabelPathBuffer `protobuf:"bytes,55,rep,name=path_buffer,json=pathBuffer,proto3" json:"path_buffer,omitempty"`
	IsLocalMac                      bool                    `protobuf:"varint,56,opt,name=is_local_mac,json=isLocalMac,proto3" json:"is_local_mac,omitempty"`
	IsProxyEntry                    bool                    `protobuf:"varint,57,opt,name=is_proxy_entry,json=isProxyEntry,proto3" json:"is_proxy_entry,omitempty"`
	IsRemoteMac                     bool                    `protobuf:"varint,58,opt,name=is_remote_mac,json=isRemoteMac,proto3" json:"is_remote_mac,omitempty"`
	SooNexthop                      string                  `protobuf:"bytes,59,opt,name=soo_nexthop,json=sooNexthop,proto3" json:"soo_nexthop,omitempty"`
	IpnhAddress                     string                  `protobuf:"bytes,60,opt,name=ipnh_address,json=ipnhAddress,proto3" json:"ipnh_address,omitempty"`
	LocalEthernetSegmentIdentifier  []uint32                `protobuf:"varint,61,rep,packed,name=local_ethernet_segment_identifier,json=localEthernetSegmentIdentifier,proto3" json:"local_ethernet_segment_identifier,omitempty"`
	RemoteEthernetSegmentIdentifier []uint32                `protobuf:"varint,62,rep,packed,name=remote_ethernet_segment_identifier,json=remoteEthernetSegmentIdentifier,proto3" json:"remote_ethernet_segment_identifier,omitempty"`
	EsiPortKey                      uint32                  `protobuf:"varint,63,opt,name=esi_port_key,json=esiPortKey,proto3" json:"esi_port_key,omitempty"`
	LocalEncapType                  uint32                  `protobuf:"varint,64,opt,name=local_encap_type,json=localEncapType,proto3" json:"local_encap_type,omitempty"`
	RemoteEncapType                 uint32                  `protobuf:"varint,65,opt,name=remote_encap_type,json=remoteEncapType,proto3" json:"remote_encap_type,omitempty"`
	LearnedBridgePortName           string                  `protobuf:"bytes,66,opt,name=learned_bridge_port_name,json=learnedBridgePortName,proto3" json:"learned_bridge_port_name,omitempty"`
	LocalSeqId                      uint32                  `protobuf:"varint,67,opt,name=local_seq_id,json=localSeqId,proto3" json:"local_seq_id,omitempty"`
	RemoteSeqId                     uint32                  `protobuf:"varint,68,opt,name=remote_seq_id,json=remoteSeqId,proto3" json:"remote_seq_id,omitempty"`
	LocalL3Label                    uint32                  `protobuf:"varint,69,opt,name=local_l3_label,json=localL3Label,proto3" json:"local_l3_label,omitempty"`
	RouterMacAddress                string                  `protobuf:"bytes,70,opt,name=router_mac_address,json=routerMacAddress,proto3" json:"router_mac_address,omitempty"`
	MacFlushRequested               uint32                  `protobuf:"varint,71,opt,name=mac_flush_requested,json=macFlushRequested,proto3" json:"mac_flush_requested,omitempty"`
	MacFlushReceived                uint32                  `protobuf:"varint,72,opt,name=mac_flush_received,json=macFlushReceived,proto3" json:"mac_flush_received,omitempty"`
	InternalLabel                   uint32                  `protobuf:"varint,73,opt,name=internal_label,json=internalLabel,proto3" json:"internal_label,omitempty"`
	Resolved                        bool                    `protobuf:"varint,74,opt,name=resolved,proto3" json:"resolved,omitempty"`
	IsStatic                        bool                    `protobuf:"varint,75,opt,name=is_static,json=isStatic,proto3" json:"is_static,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                `json:"-"`
	XXX_unrecognized                []byte                  `json:"-"`
	XXX_sizecache                   int32                   `json:"-"`
}

func (m *L2VpnEvpnMac) Reset()         { *m = L2VpnEvpnMac{} }
func (m *L2VpnEvpnMac) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnMac) ProtoMessage()    {}
func (*L2VpnEvpnMac) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f7d8f00693a691d, []int{2}
}

func (m *L2VpnEvpnMac) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnMac.Unmarshal(m, b)
}
func (m *L2VpnEvpnMac) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnMac.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnMac) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnMac.Merge(m, src)
}
func (m *L2VpnEvpnMac) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnMac.Size(m)
}
func (m *L2VpnEvpnMac) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnMac.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnMac proto.InternalMessageInfo

func (m *L2VpnEvpnMac) GetEthernetTagXr() uint32 {
	if m != nil {
		return m.EthernetTagXr
	}
	return 0
}

func (m *L2VpnEvpnMac) GetMacAddressXr() string {
	if m != nil {
		return m.MacAddressXr
	}
	return ""
}

func (m *L2VpnEvpnMac) GetIpAddressXr() string {
	if m != nil {
		return m.IpAddressXr
	}
	return ""
}

func (m *L2VpnEvpnMac) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

func (m *L2VpnEvpnMac) GetNumPaths() uint32 {
	if m != nil {
		return m.NumPaths
	}
	return 0
}

func (m *L2VpnEvpnMac) GetPathBuffer() []*L2VpnLabelPathBuffer {
	if m != nil {
		return m.PathBuffer
	}
	return nil
}

func (m *L2VpnEvpnMac) GetIsLocalMac() bool {
	if m != nil {
		return m.IsLocalMac
	}
	return false
}

func (m *L2VpnEvpnMac) GetIsProxyEntry() bool {
	if m != nil {
		return m.IsProxyEntry
	}
	return false
}

func (m *L2VpnEvpnMac) GetIsRemoteMac() bool {
	if m != nil {
		return m.IsRemoteMac
	}
	return false
}

func (m *L2VpnEvpnMac) GetSooNexthop() string {
	if m != nil {
		return m.SooNexthop
	}
	return ""
}

func (m *L2VpnEvpnMac) GetIpnhAddress() string {
	if m != nil {
		return m.IpnhAddress
	}
	return ""
}

func (m *L2VpnEvpnMac) GetLocalEthernetSegmentIdentifier() []uint32 {
	if m != nil {
		return m.LocalEthernetSegmentIdentifier
	}
	return nil
}

func (m *L2VpnEvpnMac) GetRemoteEthernetSegmentIdentifier() []uint32 {
	if m != nil {
		return m.RemoteEthernetSegmentIdentifier
	}
	return nil
}

func (m *L2VpnEvpnMac) GetEsiPortKey() uint32 {
	if m != nil {
		return m.EsiPortKey
	}
	return 0
}

func (m *L2VpnEvpnMac) GetLocalEncapType() uint32 {
	if m != nil {
		return m.LocalEncapType
	}
	return 0
}

func (m *L2VpnEvpnMac) GetRemoteEncapType() uint32 {
	if m != nil {
		return m.RemoteEncapType
	}
	return 0
}

func (m *L2VpnEvpnMac) GetLearnedBridgePortName() string {
	if m != nil {
		return m.LearnedBridgePortName
	}
	return ""
}

func (m *L2VpnEvpnMac) GetLocalSeqId() uint32 {
	if m != nil {
		return m.LocalSeqId
	}
	return 0
}

func (m *L2VpnEvpnMac) GetRemoteSeqId() uint32 {
	if m != nil {
		return m.RemoteSeqId
	}
	return 0
}

func (m *L2VpnEvpnMac) GetLocalL3Label() uint32 {
	if m != nil {
		return m.LocalL3Label
	}
	return 0
}

func (m *L2VpnEvpnMac) GetRouterMacAddress() string {
	if m != nil {
		return m.RouterMacAddress
	}
	return ""
}

func (m *L2VpnEvpnMac) GetMacFlushRequested() uint32 {
	if m != nil {
		return m.MacFlushRequested
	}
	return 0
}

func (m *L2VpnEvpnMac) GetMacFlushReceived() uint32 {
	if m != nil {
		return m.MacFlushReceived
	}
	return 0
}

func (m *L2VpnEvpnMac) GetInternalLabel() uint32 {
	if m != nil {
		return m.InternalLabel
	}
	return 0
}

func (m *L2VpnEvpnMac) GetResolved() bool {
	if m != nil {
		return m.Resolved
	}
	return false
}

func (m *L2VpnEvpnMac) GetIsStatic() bool {
	if m != nil {
		return m.IsStatic
	}
	return false
}

func init() {
	proto.RegisterType((*L2VpnEvpnMac_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.macs.mac.l2vpn_evpn_mac_KEYS")
	proto.RegisterType((*L2VpnLabelPathBuffer)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.macs.mac.l2vpn_label_path_buffer")
	proto.RegisterType((*L2VpnEvpnMac)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.macs.mac.l2vpn_evpn_mac")
}

func init() { proto.RegisterFile("l2vpn_evpn_mac.proto", fileDescriptor_3f7d8f00693a691d) }

var fileDescriptor_3f7d8f00693a691d = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5d, 0x8f, 0x1b, 0x35,
	0x14, 0x55, 0x08, 0x82, 0xe4, 0xe6, 0xa3, 0xa9, 0x0b, 0xc2, 0x80, 0x60, 0xd3, 0xa8, 0xa0, 0x08,
	0xa1, 0x79, 0xd8, 0x00, 0xe5, 0x1b, 0xba, 0x25, 0xa5, 0x21, 0xdb, 0x6a, 0x35, 0xa9, 0xc4, 0x22,
	0x21, 0x59, 0xce, 0xcc, 0x4d, 0x62, 0x31, 0x63, 0xcf, 0xda, 0x9e, 0x28, 0x79, 0xe0, 0x0f, 0xf0,
	0xc4, 0x4f, 0x46, 0xb6, 0x27, 0x93, 0xec, 0xc3, 0xf2, 0xd4, 0x97, 0xd1, 0xf8, 0x9c, 0x33, 0x67,
	0xce, 0xbd, 0xbe, 0x36, 0xbc, 0x93, 0x9d, 0x6f, 0x0b, 0xc9, 0xd0, 0x3d, 0x72, 0x9e, 0x44, 0x85,
	0x56, 0x56, 0x91, 0xa7, 0x89, 0x30, 0x89, 0x62, 0x42, 0x19, 0xb6, 0xd3, 0x81, 0x54, 0x05, 0xea,
	0xc8, 0xbd, 0x45, 0x3c, 0xb1, 0x62, 0x8b, 0x11, 0x6e, 0x05, 0x4b, 0xd1, 0x72, 0x91, 0xf9, 0xd7,
	0x64, 0x23, 0xb2, 0x54, 0xa3, 0x8c, 0x72, 0x9e, 0x18, 0xf7, 0x18, 0xfd, 0xd3, 0x80, 0x07, 0xb7,
	0xdd, 0xd9, 0x7c, 0xfa, 0xc7, 0x82, 0x0c, 0xa0, 0x89, 0x5b, 0x41, 0x1b, 0xc3, 0xc6, 0xb8, 0x17,
	0xbb, 0x57, 0xf2, 0x10, 0xba, 0x68, 0x37, 0xa8, 0x25, 0x5a, 0x66, 0xf9, 0x9a, 0xbe, 0xe1, 0xa9,
	0xce, 0x01, 0x7b, 0xc5, 0xd7, 0xe4, 0x0c, 0x3a, 0xce, 0x80, 0xa7, 0xa9, 0x46, 0x63, 0x68, 0x73,
	0xd8, 0x18, 0xb7, 0x63, 0xc8, 0x79, 0xf2, 0x24, 0x20, 0xe4, 0x23, 0x00, 0x51, 0xd4, 0xfc, 0x9b,
	0x9e, 0x6f, 0x8b, 0xa2, 0xa2, 0x47, 0xbf, 0xc3, 0x7b, 0x21, 0x4b, 0xc6, 0x97, 0x98, 0xb1, 0x82,
	0xdb, 0x0d, 0x5b, 0x96, 0xab, 0x15, 0x6a, 0xf2, 0x3e, 0xb4, 0x24, 0xee, 0x2c, 0xdb, 0xa8, 0xc2,
	0x87, 0x6a, 0xc7, 0x6f, 0xbb, 0xf5, 0x73, 0x55, 0xb8, 0x60, 0xaa, 0xb4, 0x45, 0x69, 0xc3, 0x67,
	0x87, 0x60, 0x01, 0xbb, 0x74, 0xd0, 0xe8, 0xdf, 0x36, 0xf4, 0x6f, 0x57, 0x49, 0x3e, 0x85, 0x7b,
	0xa7, 0xe5, 0xb0, 0x9d, 0xa6, 0xe7, 0xfe, 0xc3, 0xde, 0x49, 0x45, 0xd7, 0x9a, 0x3c, 0x82, 0xfe,
	0x49, 0x4d, 0x4e, 0x36, 0xf1, 0xbf, 0xef, 0x1e, 0xcb, 0xba, 0xd6, 0x64, 0x04, 0xbd, 0x63, 0x61,
	0x4e, 0xf4, 0x85, 0x17, 0x75, 0xea, 0xda, 0xae, 0xb5, 0xeb, 0x4e, 0xa6, 0x12, 0x9e, 0x55, 0x31,
	0xbf, 0xf4, 0x7f, 0x03, 0x0f, 0xf9, 0x94, 0xe4, 0x43, 0x68, 0xcb, 0x32, 0xf7, 0x65, 0x1b, 0xfa,
	0x95, 0xa7, 0x5b, 0xb2, 0xcc, 0xaf, 0xdc, 0x9a, 0xfc, 0x0d, 0x9d, 0x93, 0x7e, 0xd0, 0xc7, 0xc3,
	0xe6, 0xb8, 0x73, 0xfe, 0x67, 0xf4, 0x1a, 0x66, 0x20, 0xba, 0xa3, 0xe7, 0x31, 0xb8, 0xc5, 0x45,
	0xe8, 0xff, 0x10, 0xba, 0xc2, 0xb0, 0x90, 0x3f, 0xe7, 0x09, 0xfd, 0x7a, 0xd8, 0x18, 0xb7, 0x62,
	0x10, 0xe6, 0xd2, 0x41, 0x2f, 0x78, 0xe2, 0x1a, 0x25, 0x0c, 0x2b, 0xb4, 0xda, 0xed, 0x19, 0x4a,
	0xab, 0xf7, 0xf4, 0x1b, 0xaf, 0xe9, 0x0a, 0x73, 0xe5, 0xc0, 0xa9, 0xc3, 0x7c, 0xa3, 0x0c, 0xd3,
	0x98, 0x2b, 0x8b, 0xde, 0xe8, 0x5b, 0x2f, 0xea, 0x08, 0x13, 0x7b, 0xcc, 0x39, 0x9d, 0x41, 0xc7,
	0x28, 0xc5, 0xdc, 0xfe, 0xba, 0xed, 0xfe, 0x2e, 0x8c, 0x91, 0x51, 0xea, 0x65, 0x40, 0xdc, 0x8e,
	0x8b, 0x42, 0x6e, 0xea, 0x41, 0xfa, 0xfe, 0xd0, 0x6c, 0xb9, 0x39, 0x4c, 0xda, 0x0c, 0x1e, 0x86,
	0xb0, 0xf5, 0x26, 0x1b, 0x5c, 0xe7, 0x28, 0x2d, 0x13, 0x29, 0x4a, 0x2b, 0x56, 0x02, 0x35, 0xfd,
	0x61, 0xd8, 0x1c, 0xf7, 0xe2, 0x8f, 0xbd, 0x70, 0x5a, 0xe9, 0x16, 0x41, 0x36, 0xab, 0x55, 0x64,
	0x0e, 0xa3, 0x2a, 0xef, 0xff, 0x79, 0xfd, 0xe8, 0xbd, 0xce, 0x82, 0xf2, 0x6e, 0xb3, 0x21, 0x74,
	0xd1, 0x08, 0x56, 0x28, 0x6d, 0xd9, 0x5f, 0xb8, 0xa7, 0x3f, 0x85, 0x29, 0x40, 0x23, 0xae, 0x94,
	0xb6, 0x73, 0xdc, 0x93, 0x31, 0x0c, 0xaa, 0xe4, 0x32, 0xe1, 0x05, 0xb3, 0xfb, 0x02, 0xe9, 0xcf,
	0x5e, 0xd5, 0x0f, 0x41, 0x1d, 0xfc, 0x6a, 0x5f, 0x20, 0xf9, 0x0c, 0xee, 0x1f, 0x82, 0x1d, 0xa5,
	0x4f, 0xbc, 0xf4, 0x5e, 0x95, 0xa3, 0xd6, 0x3e, 0x06, 0x9a, 0x21, 0xd7, 0x12, 0x53, 0xb6, 0xd4,
	0x22, 0x5d, 0x63, 0x88, 0x20, 0x79, 0x8e, 0xf4, 0xc2, 0xb7, 0xef, 0xdd, 0x8a, 0xbf, 0xf0, 0xb4,
	0x4b, 0xf3, 0x92, 0xe7, 0xe8, 0x02, 0x87, 0x38, 0x06, 0x6f, 0x98, 0x48, 0xe9, 0xd3, 0x93, 0xb1,
	0x5d, 0xe0, 0xcd, 0x2c, 0x75, 0x5b, 0x5a, 0xc5, 0xa8, 0x24, 0xbf, 0x84, 0x03, 0x18, 0xc0, 0xa0,
	0x79, 0x04, 0xfd, 0x6a, 0xf6, 0x27, 0xd5, 0xf8, 0x4f, 0xbd, 0x28, 0x78, 0x5f, 0x4e, 0xc2, 0x01,
	0xf8, 0x1c, 0x88, 0x56, 0xa5, 0x45, 0xcd, 0x4e, 0xaf, 0x91, 0x67, 0x3e, 0xde, 0x20, 0x30, 0x2f,
	0x8e, 0x97, 0x49, 0x04, 0x0f, 0x9c, 0x6c, 0x95, 0x95, 0x66, 0xc3, 0x34, 0xde, 0x94, 0x68, 0x2c,
	0xa6, 0xf4, 0x57, 0x6f, 0x7c, 0x3f, 0xe7, 0xc9, 0x33, 0xc7, 0xc4, 0x07, 0xc2, 0xb9, 0x9f, 0xea,
	0x13, 0x14, 0x5b, 0x4c, 0xe9, 0x73, 0x2f, 0x1f, 0x1c, 0xe5, 0x01, 0x27, 0x9f, 0x40, 0x5f, 0x48,
	0x8b, 0x5a, 0xd6, 0x07, 0x76, 0x16, 0xae, 0x87, 0x03, 0x1a, 0x22, 0x7f, 0x00, 0x2d, 0x8d, 0x46,
	0x65, 0xce, 0xea, 0x37, 0x3f, 0xca, 0xf5, 0xda, 0x9d, 0x67, 0x61, 0x98, 0xb1, 0xdc, 0x8a, 0x84,
	0xce, 0x03, 0x29, 0xcc, 0xc2, 0xaf, 0x97, 0x6f, 0xf9, 0x4b, 0x7c, 0xf2, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf6, 0x6b, 0x21, 0x1c, 0xdc, 0x05, 0x00, 0x00,
}
