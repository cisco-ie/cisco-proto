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
	SrteTunnel           string   `protobuf:"bytes,3,opt,name=srte_tunnel,json=srteTunnel,proto3" json:"srte_tunnel,omitempty"`
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

func (m *L2VpnLabelPathBuffer) GetSrteTunnel() string {
	if m != nil {
		return m.SrteTunnel
	}
	return ""
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
	LocalIsStatic                   bool                    `protobuf:"varint,75,opt,name=local_is_static,json=localIsStatic,proto3" json:"local_is_static,omitempty"`
	RemoteIsStatic                  bool                    `protobuf:"varint,76,opt,name=remote_is_static,json=remoteIsStatic,proto3" json:"remote_is_static,omitempty"`
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

func (m *L2VpnEvpnMac) GetLocalIsStatic() bool {
	if m != nil {
		return m.LocalIsStatic
	}
	return false
}

func (m *L2VpnEvpnMac) GetRemoteIsStatic() bool {
	if m != nil {
		return m.RemoteIsStatic
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
	// 792 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0x15, 0x8a, 0xa0, 0x99, 0x7c, 0x6c, 0xd7, 0x0b, 0x62, 0x00, 0xc1, 0x66, 0xa3, 0x05,
	0x45, 0x08, 0xf9, 0xa2, 0x05, 0x96, 0x6f, 0xd8, 0x2e, 0x5d, 0x36, 0x34, 0xbb, 0xaa, 0x9c, 0x5e,
	0x14, 0x09, 0x69, 0x34, 0xb1, 0x4f, 0x92, 0x11, 0xf6, 0x8c, 0x3b, 0x33, 0x8e, 0x92, 0x0b, 0x5e,
	0x80, 0xf7, 0xe2, 0xbd, 0xd0, 0x39, 0x63, 0x3b, 0xd9, 0x8b, 0x72, 0xc5, 0x8d, 0xe5, 0xf9, 0x9f,
	0x5f, 0x4f, 0xff, 0xe7, 0xc3, 0x13, 0xf6, 0x4e, 0x7e, 0xba, 0x29, 0xb5, 0x00, 0x7c, 0x14, 0x32,
	0x8d, 0x4b, 0x6b, 0xbc, 0x89, 0x9e, 0xa5, 0xca, 0xa5, 0x46, 0x28, 0xe3, 0xc4, 0xd6, 0x86, 0xa0,
	0x29, 0xc1, 0xc6, 0xf8, 0x16, 0xcb, 0xd4, 0xab, 0x0d, 0xc4, 0xb0, 0x51, 0x22, 0x03, 0x2f, 0x55,
	0x4e, 0xaf, 0xe9, 0x5a, 0xe5, 0x99, 0x05, 0x1d, 0x17, 0x32, 0x75, 0xf8, 0x18, 0xff, 0xdd, 0x61,
	0x0f, 0x5e, 0xcf, 0x2e, 0x2e, 0x2f, 0x7e, 0x9f, 0x47, 0x27, 0xec, 0x08, 0x36, 0x8a, 0x77, 0x46,
	0x9d, 0xc9, 0x20, 0xc1, 0xd7, 0xe8, 0x11, 0xeb, 0x83, 0x5f, 0x83, 0xd5, 0xe0, 0x85, 0x97, 0x2b,
	0xfe, 0x06, 0x85, 0x7a, 0x8d, 0x76, 0x2d, 0x57, 0xd1, 0x43, 0xd6, 0xc3, 0x04, 0x32, 0xcb, 0x2c,
	0x38, 0xc7, 0x8f, 0x46, 0x9d, 0x49, 0x37, 0x61, 0x85, 0x4c, 0x9f, 0x06, 0x25, 0xfa, 0x88, 0x31,
	0x55, 0xb6, 0xf1, 0x37, 0x29, 0xde, 0x55, 0x65, 0x1d, 0x1e, 0x6f, 0xd9, 0x7b, 0xc1, 0x4b, 0x2e,
	0x17, 0x90, 0x8b, 0x52, 0xfa, 0xb5, 0x58, 0x54, 0xcb, 0x25, 0xd8, 0xe8, 0x7d, 0x76, 0xac, 0x61,
	0xeb, 0xc5, 0xda, 0x94, 0x64, 0xaa, 0x9b, 0xbc, 0x8d, 0xe7, 0x17, 0xa6, 0x44, 0x63, 0xa6, 0xf2,
	0x65, 0xe5, 0xc3, 0x9f, 0x35, 0xc6, 0x82, 0x36, 0x43, 0x09, 0x8d, 0x39, 0xeb, 0x41, 0xf8, 0x4a,
	0x6b, 0xc8, 0x1b, 0x63, 0x28, 0x5d, 0x93, 0x32, 0xfe, 0xa7, 0xcb, 0x86, 0xaf, 0xb7, 0x21, 0xfa,
	0x94, 0xdd, 0x3b, 0xac, 0x57, 0x6c, 0x2d, 0x3f, 0xa5, 0xcc, 0x83, 0x83, 0x92, 0x6f, 0x6c, 0xf4,
	0x98, 0x0d, 0x0f, 0x8a, 0x46, 0xec, 0x8c, 0xd2, 0xf7, 0xf7, 0x75, 0xdf, 0xd8, 0x68, 0xcc, 0x06,
	0xfb, 0xca, 0x11, 0xfa, 0x82, 0xa0, 0x5e, 0x5b, 0xfc, 0x8d, 0x45, 0x97, 0xb9, 0x49, 0x65, 0x5e,
	0xd7, 0xf1, 0x25, 0xfd, 0x37, 0x46, 0x52, 0x28, 0xe3, 0x43, 0xd6, 0xd5, 0x55, 0x41, 0x7d, 0x71,
	0xfc, 0x2b, 0x0a, 0x1f, 0xeb, 0xaa, 0xb8, 0xc2, 0x73, 0xf4, 0x17, 0xeb, 0x1d, 0x34, 0x8c, 0x3f,
	0x19, 0x1d, 0x4d, 0x7a, 0xa7, 0x7f, 0xc4, 0xff, 0xc3, 0x92, 0xc4, 0x77, 0x0c, 0x25, 0x61, 0x78,
	0x38, 0x0f, 0x03, 0x1a, 0xb1, 0xbe, 0x72, 0x22, 0xf8, 0x2f, 0x64, 0xca, 0xbf, 0x1e, 0x75, 0x26,
	0xc7, 0x09, 0x53, 0x6e, 0x86, 0xd2, 0x4b, 0x99, 0x62, 0xa3, 0x94, 0x13, 0xa5, 0x35, 0xdb, 0x9d,
	0x00, 0xed, 0xed, 0x8e, 0x7f, 0x43, 0x4c, 0x5f, 0xb9, 0x2b, 0x14, 0x2f, 0x50, 0xa3, 0x46, 0x39,
	0x61, 0xa1, 0x30, 0x1e, 0x28, 0xd1, 0xb7, 0x04, 0xf5, 0x94, 0x4b, 0x48, 0xc3, 0x4c, 0x38, 0x4e,
	0x63, 0x04, 0x2e, 0x00, 0xee, 0xc3, 0x77, 0xf5, 0x38, 0x8d, 0x79, 0x15, 0x14, 0x5c, 0x09, 0x55,
	0xea, 0x75, 0xbb, 0x69, 0xdf, 0x37, 0xcd, 0xd6, 0xeb, 0x66, 0x15, 0xa7, 0xec, 0x51, 0x30, 0xdb,
	0x0e, 0xd9, 0xc1, 0xaa, 0x00, 0xed, 0x85, 0xca, 0x40, 0x7b, 0xb5, 0x54, 0x60, 0xf9, 0x0f, 0xa3,
	0xa3, 0xc9, 0x20, 0xf9, 0x98, 0xc0, 0x8b, 0x9a, 0x9b, 0x07, 0x6c, 0xda, 0x52, 0xd1, 0x25, 0x1b,
	0xd7, 0x7e, 0xff, 0x2b, 0xd7, 0x8f, 0x94, 0xeb, 0x61, 0x20, 0xef, 0x4e, 0x36, 0x62, 0x7d, 0x70,
	0x4a, 0x94, 0xc6, 0x7a, 0xf1, 0x27, 0xec, 0xf8, 0x4f, 0x61, 0x0b, 0xc0, 0xa9, 0x2b, 0x63, 0xfd,
	0x25, 0xec, 0xa2, 0x09, 0x3b, 0xa9, 0x9d, 0xeb, 0x54, 0x96, 0xc2, 0xef, 0x4a, 0xe0, 0x3f, 0x13,
	0x35, 0x0c, 0x46, 0x51, 0xbe, 0xde, 0x95, 0x10, 0x7d, 0xc6, 0xee, 0x37, 0xc6, 0xf6, 0xe8, 0x53,
	0x42, 0xef, 0xd5, 0x3e, 0x5a, 0xf6, 0x09, 0xe3, 0x39, 0x48, 0xab, 0x21, 0x13, 0x0b, 0xab, 0xb2,
	0x15, 0x04, 0x0b, 0x5a, 0x16, 0xc0, 0xcf, 0xa9, 0x7d, 0xef, 0xd6, 0xf1, 0x73, 0x0a, 0xa3, 0x9b,
	0x57, 0xb2, 0x00, 0x34, 0x1c, 0xec, 0x38, 0xb8, 0x15, 0x2a, 0xe3, 0xcf, 0x0e, 0xd6, 0x76, 0x0e,
	0xb7, 0xd3, 0x0c, 0x47, 0x5a, 0xdb, 0xa8, 0x91, 0x5f, 0xc2, 0x17, 0x1a, 0xc4, 0xc0, 0x3c, 0x66,
	0xc3, 0x7a, 0xf7, 0xcf, 0xea, 0xf5, 0xbf, 0x20, 0x28, 0xe4, 0x9e, 0x9d, 0x85, 0x0f, 0xe0, 0x73,
	0x16, 0x59, 0x53, 0x79, 0xb0, 0xe2, 0xf0, 0x9e, 0x79, 0x4e, 0xf6, 0x4e, 0x42, 0xe4, 0xe5, 0xfe,
	0xb6, 0x89, 0xd9, 0x03, 0xc4, 0x96, 0x79, 0xe5, 0xd6, 0xc2, 0xc2, 0x6d, 0x05, 0xce, 0x43, 0xc6,
	0x7f, 0xa5, 0xc4, 0xf7, 0x0b, 0x99, 0x3e, 0xc7, 0x48, 0xd2, 0x04, 0x30, 0xfb, 0x21, 0x9f, 0x82,
	0xda, 0x40, 0xc6, 0x5f, 0x10, 0x7e, 0xb2, 0xc7, 0x83, 0x1e, 0x7d, 0xc2, 0x86, 0x4a, 0x7b, 0xb0,
	0xba, 0xfd, 0x60, 0xa7, 0xe1, 0x7a, 0x68, 0xd4, 0x60, 0xf9, 0x03, 0x76, 0x6c, 0xc1, 0x99, 0x1c,
	0x53, 0xfd, 0x46, 0xab, 0xdc, 0x9e, 0xf1, 0x8a, 0x09, 0x45, 0x2b, 0x27, 0x9c, 0x97, 0x5e, 0xa5,
	0xfc, 0x92, 0x90, 0x01, 0xc9, 0x53, 0x37, 0x27, 0x11, 0x27, 0x5e, 0x37, 0x70, 0x0f, 0xce, 0x08,
	0x1c, 0x06, 0xbd, 0x21, 0x17, 0x6f, 0xd1, 0x4f, 0xc3, 0xd9, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xd3, 0xa1, 0x7d, 0xaa, 0x32, 0x06, 0x00, 0x00,
}
