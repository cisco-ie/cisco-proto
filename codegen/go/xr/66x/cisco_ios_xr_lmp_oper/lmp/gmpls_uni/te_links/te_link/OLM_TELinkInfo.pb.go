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
// source: OLM_TELinkInfo.proto

package cisco_ios_xr_lmp_oper_lmp_gmpls_uni_te_links_te_link

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

type OLM_TELinkInfo_KEYS struct {
	ControllerName       string   `protobuf:"bytes,1,opt,name=controller_name,json=controllerName,proto3" json:"controller_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OLM_TELinkInfo_KEYS) Reset()         { *m = OLM_TELinkInfo_KEYS{} }
func (m *OLM_TELinkInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*OLM_TELinkInfo_KEYS) ProtoMessage()    {}
func (*OLM_TELinkInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a11ed1bc61a8c37, []int{0}
}

func (m *OLM_TELinkInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OLM_TELinkInfo_KEYS.Unmarshal(m, b)
}
func (m *OLM_TELinkInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OLM_TELinkInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *OLM_TELinkInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OLM_TELinkInfo_KEYS.Merge(m, src)
}
func (m *OLM_TELinkInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_OLM_TELinkInfo_KEYS.Size(m)
}
func (m *OLM_TELinkInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OLM_TELinkInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OLM_TELinkInfo_KEYS proto.InternalMessageInfo

func (m *OLM_TELinkInfo_KEYS) GetControllerName() string {
	if m != nil {
		return m.ControllerName
	}
	return ""
}

type OLMAddrU struct {
	AddressType          string   `protobuf:"bytes,1,opt,name=address_type,json=addressType,proto3" json:"address_type,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	UnnumberedAddress    uint32   `protobuf:"varint,4,opt,name=unnumbered_address,json=unnumberedAddress,proto3" json:"unnumbered_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OLMAddrU) Reset()         { *m = OLMAddrU{} }
func (m *OLMAddrU) String() string { return proto.CompactTextString(m) }
func (*OLMAddrU) ProtoMessage()    {}
func (*OLMAddrU) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a11ed1bc61a8c37, []int{1}
}

func (m *OLMAddrU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OLMAddrU.Unmarshal(m, b)
}
func (m *OLMAddrU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OLMAddrU.Marshal(b, m, deterministic)
}
func (m *OLMAddrU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OLMAddrU.Merge(m, src)
}
func (m *OLMAddrU) XXX_Size() int {
	return xxx_messageInfo_OLMAddrU.Size(m)
}
func (m *OLMAddrU) XXX_DiscardUnknown() {
	xxx_messageInfo_OLMAddrU.DiscardUnknown(m)
}

var xxx_messageInfo_OLMAddrU proto.InternalMessageInfo

func (m *OLMAddrU) GetAddressType() string {
	if m != nil {
		return m.AddressType
	}
	return ""
}

func (m *OLMAddrU) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *OLMAddrU) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

func (m *OLMAddrU) GetUnnumberedAddress() uint32 {
	if m != nil {
		return m.UnnumberedAddress
	}
	return 0
}

type OLMAddr_ struct {
	Address              *OLMAddrU `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OLMAddr_) Reset()         { *m = OLMAddr_{} }
func (m *OLMAddr_) String() string { return proto.CompactTextString(m) }
func (*OLMAddr_) ProtoMessage()    {}
func (*OLMAddr_) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a11ed1bc61a8c37, []int{2}
}

func (m *OLMAddr_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OLMAddr_.Unmarshal(m, b)
}
func (m *OLMAddr_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OLMAddr_.Marshal(b, m, deterministic)
}
func (m *OLMAddr_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OLMAddr_.Merge(m, src)
}
func (m *OLMAddr_) XXX_Size() int {
	return xxx_messageInfo_OLMAddr_.Size(m)
}
func (m *OLMAddr_) XXX_DiscardUnknown() {
	xxx_messageInfo_OLMAddr_.DiscardUnknown(m)
}

var xxx_messageInfo_OLMAddr_ proto.InternalMessageInfo

func (m *OLMAddr_) GetAddress() *OLMAddrU {
	if m != nil {
		return m.Address
	}
	return nil
}

type OLM_TELinkInfo struct {
	InterfaceName                       string    `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	ProtocolOwner                       string    `protobuf:"bytes,51,opt,name=protocol_owner,json=protocolOwner,proto3" json:"protocol_owner,omitempty"`
	LocalLinkId                         *OLMAddr_ `protobuf:"bytes,52,opt,name=local_link_id,json=localLinkId,proto3" json:"local_link_id,omitempty"`
	RemoteLinkId                        *OLMAddr_ `protobuf:"bytes,53,opt,name=remote_link_id,json=remoteLinkId,proto3" json:"remote_link_id,omitempty"`
	LocalTeLinkId                       *OLMAddr_ `protobuf:"bytes,54,opt,name=local_te_link_id,json=localTeLinkId,proto3" json:"local_te_link_id,omitempty"`
	RemoteTeLinkId                      *OLMAddr_ `protobuf:"bytes,55,opt,name=remote_te_link_id,json=remoteTeLinkId,proto3" json:"remote_te_link_id,omitempty"`
	NeighborName                        string    `protobuf:"bytes,56,opt,name=neighbor_name,json=neighborName,proto3" json:"neighbor_name,omitempty"`
	NeighborAddress                     *OLMAddr_ `protobuf:"bytes,57,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	IpccId                              uint32    `protobuf:"varint,58,opt,name=ipcc_id,json=ipccId,proto3" json:"ipcc_id,omitempty"`
	IpcCtype                            string    `protobuf:"bytes,59,opt,name=ipc_ctype,json=ipcCtype,proto3" json:"ipc_ctype,omitempty"`
	IpccName                            string    `protobuf:"bytes,60,opt,name=ipcc_name,json=ipccName,proto3" json:"ipcc_name,omitempty"`
	RemoteIpccAddress                   *OLMAddr_ `protobuf:"bytes,61,opt,name=remote_ipcc_address,json=remoteIpccAddress,proto3" json:"remote_ipcc_address,omitempty"`
	LocalMuxCap                         string    `protobuf:"bytes,62,opt,name=local_mux_cap,json=localMuxCap,proto3" json:"local_mux_cap,omitempty"`
	RemoteMuxCap                        string    `protobuf:"bytes,63,opt,name=remote_mux_cap,json=remoteMuxCap,proto3" json:"remote_mux_cap,omitempty"`
	ImState                             string    `protobuf:"bytes,64,opt,name=im_state,json=imState,proto3" json:"im_state,omitempty"`
	LmpState                            string    `protobuf:"bytes,65,opt,name=lmp_state,json=lmpState,proto3" json:"lmp_state,omitempty"`
	TeLinkLmpState                      string    `protobuf:"bytes,66,opt,name=te_link_lmp_state,json=teLinkLmpState,proto3" json:"te_link_lmp_state,omitempty"`
	GmplsTeLinkLocalMinimumBandwidth    uint64    `protobuf:"varint,67,opt,name=gmpls_te_link_local_minimum_bandwidth,json=gmplsTeLinkLocalMinimumBandwidth,proto3" json:"gmpls_te_link_local_minimum_bandwidth,omitempty"`
	GmplsTeLinkLocalMaximumBandwidth    uint64    `protobuf:"varint,68,opt,name=gmpls_te_link_local_maximum_bandwidth,json=gmplsTeLinkLocalMaximumBandwidth,proto3" json:"gmpls_te_link_local_maximum_bandwidth,omitempty"`
	GmplsTeLinkNeighborMinimumBandwidth uint64    `protobuf:"varint,69,opt,name=gmpls_te_link_neighbor_minimum_bandwidth,json=gmplsTeLinkNeighborMinimumBandwidth,proto3" json:"gmpls_te_link_neighbor_minimum_bandwidth,omitempty"`
	GmplsTeLinkNeighborMaximumBandwidth uint64    `protobuf:"varint,70,opt,name=gmpls_te_link_neighbor_maximum_bandwidth,json=gmplsTeLinkNeighborMaximumBandwidth,proto3" json:"gmpls_te_link_neighbor_maximum_bandwidth,omitempty"`
	GmplsTeLinkLocalEncodingType        string    `protobuf:"bytes,71,opt,name=gmpls_te_link_local_encoding_type,json=gmplsTeLinkLocalEncodingType,proto3" json:"gmpls_te_link_local_encoding_type,omitempty"`
	GmplsTeLinkNeighborEncodingType     string    `protobuf:"bytes,72,opt,name=gmpls_te_link_neighbor_encoding_type,json=gmplsTeLinkNeighborEncodingType,proto3" json:"gmpls_te_link_neighbor_encoding_type,omitempty"`
	IsLmpEnabled                        bool      `protobuf:"varint,73,opt,name=is_lmp_enabled,json=isLmpEnabled,proto3" json:"is_lmp_enabled,omitempty"`
	LmpCompLinkStatus                   []string  `protobuf:"bytes,74,rep,name=lmp_comp_link_status,json=lmpCompLinkStatus,proto3" json:"lmp_comp_link_status,omitempty"`
	LmpTransmitMsgId                    uint32    `protobuf:"varint,75,opt,name=lmp_transmit_msg_id,json=lmpTransmitMsgId,proto3" json:"lmp_transmit_msg_id,omitempty"`
	LmpReceiveMsgId                     uint32    `protobuf:"varint,76,opt,name=lmp_receive_msg_id,json=lmpReceiveMsgId,proto3" json:"lmp_receive_msg_id,omitempty"`
	XXX_NoUnkeyedLiteral                struct{}  `json:"-"`
	XXX_unrecognized                    []byte    `json:"-"`
	XXX_sizecache                       int32     `json:"-"`
}

func (m *OLM_TELinkInfo) Reset()         { *m = OLM_TELinkInfo{} }
func (m *OLM_TELinkInfo) String() string { return proto.CompactTextString(m) }
func (*OLM_TELinkInfo) ProtoMessage()    {}
func (*OLM_TELinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a11ed1bc61a8c37, []int{3}
}

func (m *OLM_TELinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OLM_TELinkInfo.Unmarshal(m, b)
}
func (m *OLM_TELinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OLM_TELinkInfo.Marshal(b, m, deterministic)
}
func (m *OLM_TELinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OLM_TELinkInfo.Merge(m, src)
}
func (m *OLM_TELinkInfo) XXX_Size() int {
	return xxx_messageInfo_OLM_TELinkInfo.Size(m)
}
func (m *OLM_TELinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OLM_TELinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OLM_TELinkInfo proto.InternalMessageInfo

func (m *OLM_TELinkInfo) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OLM_TELinkInfo) GetProtocolOwner() string {
	if m != nil {
		return m.ProtocolOwner
	}
	return ""
}

func (m *OLM_TELinkInfo) GetLocalLinkId() *OLMAddr_ {
	if m != nil {
		return m.LocalLinkId
	}
	return nil
}

func (m *OLM_TELinkInfo) GetRemoteLinkId() *OLMAddr_ {
	if m != nil {
		return m.RemoteLinkId
	}
	return nil
}

func (m *OLM_TELinkInfo) GetLocalTeLinkId() *OLMAddr_ {
	if m != nil {
		return m.LocalTeLinkId
	}
	return nil
}

func (m *OLM_TELinkInfo) GetRemoteTeLinkId() *OLMAddr_ {
	if m != nil {
		return m.RemoteTeLinkId
	}
	return nil
}

func (m *OLM_TELinkInfo) GetNeighborName() string {
	if m != nil {
		return m.NeighborName
	}
	return ""
}

func (m *OLM_TELinkInfo) GetNeighborAddress() *OLMAddr_ {
	if m != nil {
		return m.NeighborAddress
	}
	return nil
}

func (m *OLM_TELinkInfo) GetIpccId() uint32 {
	if m != nil {
		return m.IpccId
	}
	return 0
}

func (m *OLM_TELinkInfo) GetIpcCtype() string {
	if m != nil {
		return m.IpcCtype
	}
	return ""
}

func (m *OLM_TELinkInfo) GetIpccName() string {
	if m != nil {
		return m.IpccName
	}
	return ""
}

func (m *OLM_TELinkInfo) GetRemoteIpccAddress() *OLMAddr_ {
	if m != nil {
		return m.RemoteIpccAddress
	}
	return nil
}

func (m *OLM_TELinkInfo) GetLocalMuxCap() string {
	if m != nil {
		return m.LocalMuxCap
	}
	return ""
}

func (m *OLM_TELinkInfo) GetRemoteMuxCap() string {
	if m != nil {
		return m.RemoteMuxCap
	}
	return ""
}

func (m *OLM_TELinkInfo) GetImState() string {
	if m != nil {
		return m.ImState
	}
	return ""
}

func (m *OLM_TELinkInfo) GetLmpState() string {
	if m != nil {
		return m.LmpState
	}
	return ""
}

func (m *OLM_TELinkInfo) GetTeLinkLmpState() string {
	if m != nil {
		return m.TeLinkLmpState
	}
	return ""
}

func (m *OLM_TELinkInfo) GetGmplsTeLinkLocalMinimumBandwidth() uint64 {
	if m != nil {
		return m.GmplsTeLinkLocalMinimumBandwidth
	}
	return 0
}

func (m *OLM_TELinkInfo) GetGmplsTeLinkLocalMaximumBandwidth() uint64 {
	if m != nil {
		return m.GmplsTeLinkLocalMaximumBandwidth
	}
	return 0
}

func (m *OLM_TELinkInfo) GetGmplsTeLinkNeighborMinimumBandwidth() uint64 {
	if m != nil {
		return m.GmplsTeLinkNeighborMinimumBandwidth
	}
	return 0
}

func (m *OLM_TELinkInfo) GetGmplsTeLinkNeighborMaximumBandwidth() uint64 {
	if m != nil {
		return m.GmplsTeLinkNeighborMaximumBandwidth
	}
	return 0
}

func (m *OLM_TELinkInfo) GetGmplsTeLinkLocalEncodingType() string {
	if m != nil {
		return m.GmplsTeLinkLocalEncodingType
	}
	return ""
}

func (m *OLM_TELinkInfo) GetGmplsTeLinkNeighborEncodingType() string {
	if m != nil {
		return m.GmplsTeLinkNeighborEncodingType
	}
	return ""
}

func (m *OLM_TELinkInfo) GetIsLmpEnabled() bool {
	if m != nil {
		return m.IsLmpEnabled
	}
	return false
}

func (m *OLM_TELinkInfo) GetLmpCompLinkStatus() []string {
	if m != nil {
		return m.LmpCompLinkStatus
	}
	return nil
}

func (m *OLM_TELinkInfo) GetLmpTransmitMsgId() uint32 {
	if m != nil {
		return m.LmpTransmitMsgId
	}
	return 0
}

func (m *OLM_TELinkInfo) GetLmpReceiveMsgId() uint32 {
	if m != nil {
		return m.LmpReceiveMsgId
	}
	return 0
}

func init() {
	proto.RegisterType((*OLM_TELinkInfo_KEYS)(nil), "cisco_ios_xr_lmp_oper.lmp.gmpls_uni.te_links.te_link.OLM_TELinkInfo_KEYS")
	proto.RegisterType((*OLMAddrU)(nil), "cisco_ios_xr_lmp_oper.lmp.gmpls_uni.te_links.te_link.OLMAddrU")
	proto.RegisterType((*OLMAddr_)(nil), "cisco_ios_xr_lmp_oper.lmp.gmpls_uni.te_links.te_link.OLM_addr_")
	proto.RegisterType((*OLM_TELinkInfo)(nil), "cisco_ios_xr_lmp_oper.lmp.gmpls_uni.te_links.te_link.OLM_TELinkInfo")
}

func init() { proto.RegisterFile("OLM_TELinkInfo.proto", fileDescriptor_3a11ed1bc61a8c37) }

var fileDescriptor_3a11ed1bc61a8c37 = []byte{
	// 782 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x6d, 0x4f, 0x1a, 0x4b,
	0x14, 0xc7, 0xc3, 0xd5, 0x28, 0x8c, 0x80, 0xb2, 0x9a, 0xdc, 0xbd, 0xb9, 0x37, 0xb9, 0x88, 0x9a,
	0xd2, 0x34, 0xd2, 0x44, 0xad, 0x7d, 0xd6, 0x2a, 0xa5, 0x96, 0x0a, 0x92, 0xac, 0x98, 0xb4, 0xaf,
	0x26, 0xcb, 0xec, 0x88, 0xa3, 0x3b, 0x0f, 0xd9, 0x5d, 0x14, 0x3f, 0x4c, 0xdf, 0xf4, 0x93, 0x36,
	0xf3, 0xb4, 0x88, 0x4a, 0x5f, 0x18, 0xde, 0xb1, 0xe7, 0xfc, 0xf6, 0x7f, 0xce, 0xff, 0xcc, 0x19,
	0x16, 0xac, 0x74, 0x5a, 0x6d, 0xd8, 0x6d, 0xb4, 0x08, 0xbb, 0x6a, 0xb2, 0x73, 0x5e, 0x13, 0x11,
	0x4f, 0xb8, 0xb3, 0x83, 0x48, 0x8c, 0x38, 0x24, 0x3c, 0x86, 0xc3, 0x08, 0x86, 0x54, 0x40, 0x2e,
	0x70, 0x54, 0x0b, 0xa9, 0xa8, 0xf5, 0xa9, 0x08, 0x63, 0x38, 0x60, 0xa4, 0x96, 0x60, 0x18, 0x12,
	0x76, 0x15, 0xdb, 0x1f, 0x95, 0x3d, 0xb0, 0x3c, 0xae, 0x06, 0x8f, 0x1b, 0x3f, 0x4e, 0x9d, 0x67,
	0x60, 0x11, 0x71, 0x96, 0x44, 0x3c, 0x0c, 0x71, 0x04, 0x99, 0x4f, 0xb1, 0x9b, 0x29, 0x67, 0xaa,
	0x39, 0xaf, 0x38, 0x0a, 0x9f, 0xf8, 0x14, 0x57, 0x7e, 0x65, 0x40, 0xb6, 0xd3, 0x6a, 0x1f, 0x04,
	0x41, 0x74, 0xe6, 0xac, 0x82, 0xbc, 0x1f, 0x04, 0x11, 0x8e, 0x63, 0x98, 0xdc, 0x0a, 0xfb, 0xca,
	0x82, 0x89, 0x75, 0x6f, 0x05, 0x96, 0x08, 0x11, 0xd7, 0x3b, 0xd0, 0xc4, 0xdc, 0xbf, 0x34, 0x22,
	0x63, 0x07, 0x3a, 0x64, 0x90, 0xdd, 0x14, 0x99, 0x49, 0x91, 0x5d, 0x8b, 0x6c, 0x02, 0x67, 0xc0,
	0xd8, 0x80, 0xf6, 0x70, 0x84, 0x83, 0x14, 0x9c, 0x2d, 0x67, 0xaa, 0x05, 0xaf, 0x34, 0xca, 0x18,
	0xbc, 0x82, 0x41, 0x4e, 0x9a, 0x94, 0x1c, 0x74, 0xbe, 0x83, 0x79, 0xfb, 0x82, 0xec, 0x6f, 0x61,
	0x6b, 0xaf, 0xf6, 0x94, 0xc9, 0xd5, 0xac, 0x6b, 0xcf, 0xca, 0x55, 0x7e, 0x16, 0x40, 0x71, 0x7c,
	0x98, 0xce, 0x06, 0x28, 0x12, 0x96, 0xe0, 0xe8, 0xdc, 0x47, 0x58, 0x8f, 0x71, 0x4b, 0xb9, 0x29,
	0xa4, 0x51, 0x39, 0x45, 0x89, 0xa9, 0x43, 0x44, 0x3c, 0x84, 0xfc, 0x86, 0xe1, 0xc8, 0xdd, 0xd6,
	0x98, 0x8d, 0x76, 0x64, 0xd0, 0x41, 0xa0, 0x10, 0x72, 0xe4, 0x87, 0xaa, 0x01, 0x48, 0x02, 0x77,
	0x47, 0x19, 0xd8, 0x7f, 0xb2, 0x01, 0x3d, 0x12, 0x6f, 0x41, 0xa9, 0xaa, 0xa6, 0x03, 0x07, 0x83,
	0x62, 0x84, 0x29, 0x37, 0x94, 0xac, 0xf2, 0x6a, 0x3a, 0x55, 0xf2, 0x5a, 0xd6, 0x94, 0xb9, 0x00,
	0x4b, 0xda, 0xcb, 0x9d, 0x42, 0xbb, 0xd3, 0x29, 0xa4, 0x87, 0xd4, 0xb5, 0x95, 0x2e, 0x41, 0xc9,
	0x18, 0xba, 0x53, 0xea, 0xf5, 0x74, 0x4a, 0x99, 0x51, 0xa5, 0xb5, 0xd6, 0x40, 0x81, 0x61, 0xd2,
	0xbf, 0xe8, 0x71, 0x73, 0x6b, 0xde, 0xa8, 0x73, 0xcc, 0xdb, 0xa0, 0x3a, 0xed, 0x4b, 0xb0, 0x94,
	0x42, 0x76, 0x15, 0xdf, 0x4e, 0xa7, 0x9f, 0x45, 0x2b, 0x6c, 0x6f, 0xca, 0xdf, 0x60, 0x9e, 0x08,
	0x84, 0xa4, 0xe5, 0x77, 0xea, 0x7a, 0xcc, 0xc9, 0xc7, 0x66, 0xe0, 0xfc, 0x0b, 0x72, 0x44, 0x20,
	0x88, 0xd4, 0x45, 0x7d, 0xaf, 0xba, 0xcc, 0x12, 0x81, 0xea, 0xf2, 0xd9, 0x24, 0x91, 0xb6, 0xf0,
	0x21, 0x4d, 0x22, 0xd5, 0x3e, 0x07, 0xcb, 0x66, 0x9e, 0x8a, 0xb1, 0x0e, 0x3e, 0x4e, 0xc7, 0x81,
	0x39, 0xab, 0xa6, 0x40, 0xc8, 0x7a, 0xa8, 0xd8, 0xb5, 0xa7, 0x83, 0x21, 0x44, 0xbe, 0x70, 0xf7,
	0xf4, 0x3f, 0x82, 0x0a, 0xb6, 0x07, 0xc3, 0xba, 0x2f, 0x9c, 0xf5, 0x74, 0x6b, 0x2d, 0xb4, 0xaf,
	0x27, 0xaf, 0xa3, 0x86, 0xfa, 0x07, 0x64, 0x09, 0x85, 0x71, 0xe2, 0x27, 0xd8, 0xfd, 0xa4, 0xf2,
	0xf3, 0x84, 0x9e, 0xca, 0x47, 0x69, 0x59, 0x36, 0xab, 0x73, 0x07, 0xda, 0x72, 0x48, 0x85, 0x4e,
	0x3e, 0x07, 0x25, 0xbb, 0x3b, 0x23, 0xe8, 0x50, 0xff, 0x21, 0xea, 0x8d, 0x6e, 0x59, 0xb4, 0x03,
	0x36, 0xb4, 0xcf, 0xf4, 0x05, 0xdd, 0x3a, 0x61, 0x84, 0x0e, 0x28, 0xec, 0xf9, 0x2c, 0xb8, 0x21,
	0x41, 0x72, 0xe1, 0xd6, 0xcb, 0x99, 0xea, 0xac, 0x57, 0x56, 0xb0, 0xde, 0x9f, 0x96, 0xf2, 0xa3,
	0xc1, 0x43, 0xcb, 0x4d, 0x14, 0xf4, 0x87, 0xf7, 0x04, 0x3f, 0x4f, 0x10, 0xd4, 0xe0, 0x48, 0xf0,
	0x0c, 0x54, 0xc7, 0x05, 0xd3, 0x65, 0x7c, 0xd8, 0x64, 0x43, 0x69, 0xae, 0xdd, 0xd1, 0x3c, 0x31,
	0xf0, 0x83, 0x3e, 0xff, 0x20, 0xfb, 0xa0, 0xd5, 0x2f, 0x93, 0x65, 0xef, 0x77, 0x7b, 0x04, 0x56,
	0x1f, 0xb3, 0x8f, 0x19, 0xe2, 0x01, 0x61, 0x7d, 0xfd, 0xa1, 0x39, 0x52, 0x47, 0xf1, 0xdf, 0x7d,
	0xeb, 0x0d, 0x03, 0xa9, 0x2f, 0x4f, 0x1b, 0xac, 0x4f, 0xe8, 0x6f, 0x5c, 0xeb, 0xab, 0xd2, 0xfa,
	0xff, 0x91, 0xde, 0xc6, 0xe4, 0xd6, 0x41, 0x91, 0xc4, 0x6a, 0x1b, 0x30, 0xf3, 0x7b, 0x21, 0x0e,
	0xdc, 0x66, 0x39, 0x53, 0xcd, 0x7a, 0x79, 0x12, 0xb7, 0xa8, 0x68, 0xe8, 0x98, 0xf3, 0x12, 0xac,
	0x48, 0x04, 0x71, 0x2a, 0x74, 0x51, 0xb9, 0x3a, 0x83, 0xd8, 0xfd, 0x56, 0x9e, 0xa9, 0xe6, 0xbc,
	0x52, 0x48, 0x45, 0x9d, 0x53, 0x21, 0x8b, 0x9c, 0xaa, 0x84, 0xb3, 0x09, 0x96, 0xe5, 0x0b, 0x49,
	0xe4, 0xb3, 0x98, 0x92, 0x04, 0xd2, 0xb8, 0x2f, 0xef, 0xee, 0xb1, 0xba, 0xbb, 0x4b, 0x21, 0x15,
	0x5d, 0x93, 0x69, 0xc7, 0xfd, 0x66, 0xe0, 0xbc, 0x00, 0x8e, 0xc4, 0x23, 0x8c, 0x30, 0xb9, 0xc6,
	0x96, 0x6e, 0x29, 0x7a, 0x31, 0xa4, 0xc2, 0xd3, 0x09, 0x05, 0xf7, 0xe6, 0xd4, 0xd7, 0x64, 0xfb,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x1d, 0x22, 0x83, 0x40, 0x08, 0x00, 0x00,
}
