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
// source: cdp_neighbor_entry.proto

package cisco_ios_xr_cdp_oper_cdp_nodes_node_neighbors_devices_device

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

type CdpNeighborEntry_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	DeviceId             string   `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CdpNeighborEntry_KEYS) Reset()         { *m = CdpNeighborEntry_KEYS{} }
func (m *CdpNeighborEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*CdpNeighborEntry_KEYS) ProtoMessage()    {}
func (*CdpNeighborEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3180752218293e, []int{0}
}

func (m *CdpNeighborEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdpNeighborEntry_KEYS.Unmarshal(m, b)
}
func (m *CdpNeighborEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdpNeighborEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *CdpNeighborEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdpNeighborEntry_KEYS.Merge(m, src)
}
func (m *CdpNeighborEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_CdpNeighborEntry_KEYS.Size(m)
}
func (m *CdpNeighborEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_CdpNeighborEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_CdpNeighborEntry_KEYS proto.InternalMessageInfo

func (m *CdpNeighborEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *CdpNeighborEntry_KEYS) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

type CdpL3Addr struct {
	AddressType          string   `protobuf:"bytes,1,opt,name=address_type,json=addressType,proto3" json:"address_type,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CdpL3Addr) Reset()         { *m = CdpL3Addr{} }
func (m *CdpL3Addr) String() string { return proto.CompactTextString(m) }
func (*CdpL3Addr) ProtoMessage()    {}
func (*CdpL3Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3180752218293e, []int{1}
}

func (m *CdpL3Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdpL3Addr.Unmarshal(m, b)
}
func (m *CdpL3Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdpL3Addr.Marshal(b, m, deterministic)
}
func (m *CdpL3Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdpL3Addr.Merge(m, src)
}
func (m *CdpL3Addr) XXX_Size() int {
	return xxx_messageInfo_CdpL3Addr.Size(m)
}
func (m *CdpL3Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_CdpL3Addr.DiscardUnknown(m)
}

var xxx_messageInfo_CdpL3Addr proto.InternalMessageInfo

func (m *CdpL3Addr) GetAddressType() string {
	if m != nil {
		return m.AddressType
	}
	return ""
}

func (m *CdpL3Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *CdpL3Addr) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type CdpAddrEntryItem struct {
	Address              *CdpL3Addr `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CdpAddrEntryItem) Reset()         { *m = CdpAddrEntryItem{} }
func (m *CdpAddrEntryItem) String() string { return proto.CompactTextString(m) }
func (*CdpAddrEntryItem) ProtoMessage()    {}
func (*CdpAddrEntryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3180752218293e, []int{2}
}

func (m *CdpAddrEntryItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdpAddrEntryItem.Unmarshal(m, b)
}
func (m *CdpAddrEntryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdpAddrEntryItem.Marshal(b, m, deterministic)
}
func (m *CdpAddrEntryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdpAddrEntryItem.Merge(m, src)
}
func (m *CdpAddrEntryItem) XXX_Size() int {
	return xxx_messageInfo_CdpAddrEntryItem.Size(m)
}
func (m *CdpAddrEntryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CdpAddrEntryItem.DiscardUnknown(m)
}

var xxx_messageInfo_CdpAddrEntryItem proto.InternalMessageInfo

func (m *CdpAddrEntryItem) GetAddress() *CdpL3Addr {
	if m != nil {
		return m.Address
	}
	return nil
}

type CdpAddrEntryEntry struct {
	CdpAddrEntry         []*CdpAddrEntryItem `protobuf:"bytes,1,rep,name=cdp_addr_entry,json=cdpAddrEntry,proto3" json:"cdp_addr_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CdpAddrEntryEntry) Reset()         { *m = CdpAddrEntryEntry{} }
func (m *CdpAddrEntryEntry) String() string { return proto.CompactTextString(m) }
func (*CdpAddrEntryEntry) ProtoMessage()    {}
func (*CdpAddrEntryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3180752218293e, []int{3}
}

func (m *CdpAddrEntryEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdpAddrEntryEntry.Unmarshal(m, b)
}
func (m *CdpAddrEntryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdpAddrEntryEntry.Marshal(b, m, deterministic)
}
func (m *CdpAddrEntryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdpAddrEntryEntry.Merge(m, src)
}
func (m *CdpAddrEntryEntry) XXX_Size() int {
	return xxx_messageInfo_CdpAddrEntryEntry.Size(m)
}
func (m *CdpAddrEntryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CdpAddrEntryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CdpAddrEntryEntry proto.InternalMessageInfo

func (m *CdpAddrEntryEntry) GetCdpAddrEntry() []*CdpAddrEntryItem {
	if m != nil {
		return m.CdpAddrEntry
	}
	return nil
}

type CdpProtHelloEntryItem struct {
	HelloMessage         []byte   `protobuf:"bytes,1,opt,name=hello_message,json=helloMessage,proto3" json:"hello_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CdpProtHelloEntryItem) Reset()         { *m = CdpProtHelloEntryItem{} }
func (m *CdpProtHelloEntryItem) String() string { return proto.CompactTextString(m) }
func (*CdpProtHelloEntryItem) ProtoMessage()    {}
func (*CdpProtHelloEntryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3180752218293e, []int{4}
}

func (m *CdpProtHelloEntryItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdpProtHelloEntryItem.Unmarshal(m, b)
}
func (m *CdpProtHelloEntryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdpProtHelloEntryItem.Marshal(b, m, deterministic)
}
func (m *CdpProtHelloEntryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdpProtHelloEntryItem.Merge(m, src)
}
func (m *CdpProtHelloEntryItem) XXX_Size() int {
	return xxx_messageInfo_CdpProtHelloEntryItem.Size(m)
}
func (m *CdpProtHelloEntryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CdpProtHelloEntryItem.DiscardUnknown(m)
}

var xxx_messageInfo_CdpProtHelloEntryItem proto.InternalMessageInfo

func (m *CdpProtHelloEntryItem) GetHelloMessage() []byte {
	if m != nil {
		return m.HelloMessage
	}
	return nil
}

type CdpProtHelloEntryEntry struct {
	CdpProtHelloEntry    []*CdpProtHelloEntryItem `protobuf:"bytes,1,rep,name=cdp_prot_hello_entry,json=cdpProtHelloEntry,proto3" json:"cdp_prot_hello_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CdpProtHelloEntryEntry) Reset()         { *m = CdpProtHelloEntryEntry{} }
func (m *CdpProtHelloEntryEntry) String() string { return proto.CompactTextString(m) }
func (*CdpProtHelloEntryEntry) ProtoMessage()    {}
func (*CdpProtHelloEntryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3180752218293e, []int{5}
}

func (m *CdpProtHelloEntryEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdpProtHelloEntryEntry.Unmarshal(m, b)
}
func (m *CdpProtHelloEntryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdpProtHelloEntryEntry.Marshal(b, m, deterministic)
}
func (m *CdpProtHelloEntryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdpProtHelloEntryEntry.Merge(m, src)
}
func (m *CdpProtHelloEntryEntry) XXX_Size() int {
	return xxx_messageInfo_CdpProtHelloEntryEntry.Size(m)
}
func (m *CdpProtHelloEntryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CdpProtHelloEntryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CdpProtHelloEntryEntry proto.InternalMessageInfo

func (m *CdpProtHelloEntryEntry) GetCdpProtHelloEntry() []*CdpProtHelloEntryItem {
	if m != nil {
		return m.CdpProtHelloEntry
	}
	return nil
}

type CdpNeighborDetail struct {
	NetworkAddresses     *CdpAddrEntryEntry      `protobuf:"bytes,1,opt,name=network_addresses,json=networkAddresses,proto3" json:"network_addresses,omitempty"`
	Version              string                  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	ProtocolHelloList    *CdpProtHelloEntryEntry `protobuf:"bytes,3,opt,name=protocol_hello_list,json=protocolHelloList,proto3" json:"protocol_hello_list,omitempty"`
	VtpDomain            string                  `protobuf:"bytes,4,opt,name=vtp_domain,json=vtpDomain,proto3" json:"vtp_domain,omitempty"`
	NativeVlan           uint32                  `protobuf:"varint,5,opt,name=native_vlan,json=nativeVlan,proto3" json:"native_vlan,omitempty"`
	Duplex               string                  `protobuf:"bytes,6,opt,name=duplex,proto3" json:"duplex,omitempty"`
	SystemName           string                  `protobuf:"bytes,7,opt,name=system_name,json=systemName,proto3" json:"system_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CdpNeighborDetail) Reset()         { *m = CdpNeighborDetail{} }
func (m *CdpNeighborDetail) String() string { return proto.CompactTextString(m) }
func (*CdpNeighborDetail) ProtoMessage()    {}
func (*CdpNeighborDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3180752218293e, []int{6}
}

func (m *CdpNeighborDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdpNeighborDetail.Unmarshal(m, b)
}
func (m *CdpNeighborDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdpNeighborDetail.Marshal(b, m, deterministic)
}
func (m *CdpNeighborDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdpNeighborDetail.Merge(m, src)
}
func (m *CdpNeighborDetail) XXX_Size() int {
	return xxx_messageInfo_CdpNeighborDetail.Size(m)
}
func (m *CdpNeighborDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_CdpNeighborDetail.DiscardUnknown(m)
}

var xxx_messageInfo_CdpNeighborDetail proto.InternalMessageInfo

func (m *CdpNeighborDetail) GetNetworkAddresses() *CdpAddrEntryEntry {
	if m != nil {
		return m.NetworkAddresses
	}
	return nil
}

func (m *CdpNeighborDetail) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CdpNeighborDetail) GetProtocolHelloList() *CdpProtHelloEntryEntry {
	if m != nil {
		return m.ProtocolHelloList
	}
	return nil
}

func (m *CdpNeighborDetail) GetVtpDomain() string {
	if m != nil {
		return m.VtpDomain
	}
	return ""
}

func (m *CdpNeighborDetail) GetNativeVlan() uint32 {
	if m != nil {
		return m.NativeVlan
	}
	return 0
}

func (m *CdpNeighborDetail) GetDuplex() string {
	if m != nil {
		return m.Duplex
	}
	return ""
}

func (m *CdpNeighborDetail) GetSystemName() string {
	if m != nil {
		return m.SystemName
	}
	return ""
}

type CdpNeighborItem struct {
	ReceivingInterfaceName string             `protobuf:"bytes,1,opt,name=receiving_interface_name,json=receivingInterfaceName,proto3" json:"receiving_interface_name,omitempty"`
	DeviceId               string             `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	PortId                 string             `protobuf:"bytes,3,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	HeaderVersion          uint32             `protobuf:"varint,4,opt,name=header_version,json=headerVersion,proto3" json:"header_version,omitempty"`
	HoldTime               uint32             `protobuf:"varint,5,opt,name=hold_time,json=holdTime,proto3" json:"hold_time,omitempty"`
	Capabilities           string             `protobuf:"bytes,6,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
	Platform               string             `protobuf:"bytes,7,opt,name=platform,proto3" json:"platform,omitempty"`
	Detail                 *CdpNeighborDetail `protobuf:"bytes,8,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *CdpNeighborItem) Reset()         { *m = CdpNeighborItem{} }
func (m *CdpNeighborItem) String() string { return proto.CompactTextString(m) }
func (*CdpNeighborItem) ProtoMessage()    {}
func (*CdpNeighborItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3180752218293e, []int{7}
}

func (m *CdpNeighborItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdpNeighborItem.Unmarshal(m, b)
}
func (m *CdpNeighborItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdpNeighborItem.Marshal(b, m, deterministic)
}
func (m *CdpNeighborItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdpNeighborItem.Merge(m, src)
}
func (m *CdpNeighborItem) XXX_Size() int {
	return xxx_messageInfo_CdpNeighborItem.Size(m)
}
func (m *CdpNeighborItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CdpNeighborItem.DiscardUnknown(m)
}

var xxx_messageInfo_CdpNeighborItem proto.InternalMessageInfo

func (m *CdpNeighborItem) GetReceivingInterfaceName() string {
	if m != nil {
		return m.ReceivingInterfaceName
	}
	return ""
}

func (m *CdpNeighborItem) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *CdpNeighborItem) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *CdpNeighborItem) GetHeaderVersion() uint32 {
	if m != nil {
		return m.HeaderVersion
	}
	return 0
}

func (m *CdpNeighborItem) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *CdpNeighborItem) GetCapabilities() string {
	if m != nil {
		return m.Capabilities
	}
	return ""
}

func (m *CdpNeighborItem) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *CdpNeighborItem) GetDetail() *CdpNeighborDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

type CdpNeighborEntry struct {
	CdpNeighbor          []*CdpNeighborItem `protobuf:"bytes,50,rep,name=cdp_neighbor,json=cdpNeighbor,proto3" json:"cdp_neighbor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CdpNeighborEntry) Reset()         { *m = CdpNeighborEntry{} }
func (m *CdpNeighborEntry) String() string { return proto.CompactTextString(m) }
func (*CdpNeighborEntry) ProtoMessage()    {}
func (*CdpNeighborEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3180752218293e, []int{8}
}

func (m *CdpNeighborEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdpNeighborEntry.Unmarshal(m, b)
}
func (m *CdpNeighborEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdpNeighborEntry.Marshal(b, m, deterministic)
}
func (m *CdpNeighborEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdpNeighborEntry.Merge(m, src)
}
func (m *CdpNeighborEntry) XXX_Size() int {
	return xxx_messageInfo_CdpNeighborEntry.Size(m)
}
func (m *CdpNeighborEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CdpNeighborEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CdpNeighborEntry proto.InternalMessageInfo

func (m *CdpNeighborEntry) GetCdpNeighbor() []*CdpNeighborItem {
	if m != nil {
		return m.CdpNeighbor
	}
	return nil
}

func init() {
	proto.RegisterType((*CdpNeighborEntry_KEYS)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.devices.device.cdp_neighbor_entry_KEYS")
	proto.RegisterType((*CdpL3Addr)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.devices.device.cdp_l3_addr")
	proto.RegisterType((*CdpAddrEntryItem)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.devices.device.cdp_addr_entry_item")
	proto.RegisterType((*CdpAddrEntryEntry)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.devices.device.cdp_addr_entry_entry")
	proto.RegisterType((*CdpProtHelloEntryItem)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.devices.device.cdp_prot_hello_entry_item")
	proto.RegisterType((*CdpProtHelloEntryEntry)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.devices.device.cdp_prot_hello_entry_entry")
	proto.RegisterType((*CdpNeighborDetail)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.devices.device.cdp_neighbor_detail")
	proto.RegisterType((*CdpNeighborItem)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.devices.device.cdp_neighbor_item")
	proto.RegisterType((*CdpNeighborEntry)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.devices.device.cdp_neighbor_entry")
}

func init() { proto.RegisterFile("cdp_neighbor_entry.proto", fileDescriptor_9e3180752218293e) }

var fileDescriptor_9e3180752218293e = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x4e, 0x14, 0x41,
	0x10, 0xce, 0x00, 0xee, 0x4f, 0xed, 0x42, 0xdc, 0xc6, 0xc0, 0x08, 0x31, 0xe2, 0x18, 0x13, 0x4e,
	0x73, 0x58, 0x0c, 0xf1, 0x62, 0x22, 0x89, 0x24, 0xe2, 0x0f, 0x21, 0x03, 0x21, 0x72, 0xea, 0x0c,
	0xd3, 0x05, 0xdb, 0x3a, 0x33, 0xdd, 0x99, 0x6e, 0x47, 0x36, 0x5e, 0xbc, 0x72, 0xf3, 0x31, 0x7c,
	0x03, 0x1f, 0xcc, 0x17, 0x30, 0xfd, 0x33, 0x1b, 0x76, 0x21, 0xf1, 0xe0, 0x7a, 0xe9, 0x49, 0x7f,
	0x55, 0x5d, 0xd5, 0x5f, 0xd5, 0xd7, 0x35, 0x10, 0x66, 0x4c, 0xd2, 0x12, 0xf9, 0xe5, 0xe8, 0x5c,
	0x54, 0x14, 0x4b, 0x5d, 0x8d, 0x63, 0x59, 0x09, 0x2d, 0xc8, 0xcb, 0x8c, 0xab, 0x4c, 0x50, 0x2e,
	0x14, 0xbd, 0xaa, 0xa8, 0x71, 0x13, 0x12, 0xab, 0x38, 0x63, 0x32, 0x2e, 0x05, 0x43, 0x65, 0xd7,
	0xb8, 0x39, 0xaa, 0x62, 0x86, 0x35, 0xcf, 0xb0, 0xf9, 0x46, 0xc7, 0xb0, 0x7e, 0x3b, 0x34, 0x7d,
	0xb7, 0x7f, 0x76, 0x4c, 0x36, 0xa1, 0x6b, 0xce, 0xd2, 0x32, 0x2d, 0x30, 0x0c, 0xb6, 0x82, 0xed,
	0x6e, 0xd2, 0x31, 0xc0, 0x61, 0x5a, 0xa0, 0x31, 0xba, 0x08, 0x94, 0xb3, 0x70, 0xc1, 0x19, 0x1d,
	0x70, 0xc0, 0xa2, 0x1a, 0x7a, 0x26, 0x68, 0xbe, 0x43, 0x53, 0xc6, 0x2a, 0xf2, 0x04, 0xfa, 0xe6,
	0x8b, 0x4a, 0x51, 0x3d, 0x96, 0x4d, 0xac, 0x9e, 0xc7, 0x4e, 0xc6, 0x12, 0x8d, 0x0b, 0x97, 0xf5,
	0x73, 0xea, 0x31, 0x1f, 0xb1, 0x67, 0xb0, 0x3d, 0x07, 0x79, 0x97, 0xdd, 0x89, 0xcb, 0xe2, 0xc4,
	0x65, 0xd7, 0xbb, 0x44, 0xdf, 0x60, 0xd5, 0xe4, 0x35, 0x1e, 0x9e, 0x08, 0xd7, 0x58, 0x10, 0x06,
	0xed, 0xe6, 0x90, 0x49, 0xdd, 0x1b, 0xbe, 0x8d, 0xff, 0xa9, 0x68, 0xf1, 0x0d, 0x72, 0x49, 0x13,
	0x3a, 0xfa, 0x11, 0xc0, 0x83, 0x99, 0xec, 0x76, 0x25, 0x57, 0xb0, 0x32, 0x8d, 0x87, 0xc1, 0xd6,
	0xe2, 0x76, 0x6f, 0x98, 0xcc, 0xe1, 0x16, 0x33, 0x54, 0x93, 0x7e, 0xc6, 0xa4, 0xa9, 0xc6, 0xbe,
	0x81, 0xa2, 0x57, 0xf0, 0xd0, 0x38, 0x19, 0xa1, 0xd0, 0x11, 0xe6, 0xb9, 0xb8, 0x59, 0x95, 0xa7,
	0xb0, 0xec, 0xb0, 0x02, 0x95, 0x4a, 0x2f, 0x5d, 0x5b, 0xfa, 0x49, 0xdf, 0x82, 0x1f, 0x1c, 0x16,
	0xfd, 0x0c, 0x60, 0xe3, 0xce, 0x10, 0x8e, 0xda, 0xb5, 0xe7, 0x3c, 0x6b, 0xf6, 0x0c, 0x3f, 0xce,
	0x81, 0xe1, 0x9d, 0x97, 0x4f, 0x06, 0x19, 0x93, 0x47, 0x95, 0xd0, 0x6f, 0x8c, 0xc1, 0x91, 0xfd,
	0xb5, 0xe8, 0xba, 0x3f, 0x91, 0x32, 0x43, 0x9d, 0xf2, 0x9c, 0x7c, 0x0f, 0x60, 0x50, 0xa2, 0xfe,
	0x2a, 0xaa, 0xcf, 0x8d, 0x76, 0xb0, 0x11, 0xc2, 0xf1, 0x7c, 0x5b, 0x60, 0xd7, 0xe4, 0xbe, 0xcf,
	0xb6, 0xd7, 0x24, 0x23, 0x21, 0xb4, 0x6b, 0xac, 0x14, 0x17, 0xa5, 0x17, 0x76, 0xb3, 0x35, 0x05,
	0x5c, 0xb5, 0xef, 0x38, 0x13, 0xb9, 0x67, 0x99, 0x73, 0xa5, 0xad, 0xb8, 0x7b, 0xc3, 0xb3, 0xff,
	0x51, 0x3f, 0x77, 0xc9, 0x41, 0x93, 0xd5, 0x56, 0xf0, 0x3d, 0x57, 0x9a, 0x3c, 0x02, 0xa8, 0xb5,
	0xa4, 0x4c, 0x14, 0x29, 0x2f, 0xc3, 0x25, 0x7b, 0xd1, 0x6e, 0xad, 0xe5, 0x6b, 0x0b, 0x90, 0xc7,
	0xd0, 0x2b, 0x53, 0xcd, 0x6b, 0xa4, 0x75, 0x9e, 0x96, 0xe1, 0xbd, 0xad, 0x60, 0x7b, 0x39, 0x01,
	0x07, 0x9d, 0xe6, 0x69, 0x49, 0xd6, 0xa0, 0xc5, 0xbe, 0xc8, 0x1c, 0xaf, 0xc2, 0x96, 0x3d, 0xeb,
	0x77, 0xe6, 0xa0, 0x1a, 0x2b, 0x8d, 0x85, 0x9b, 0x24, 0x6d, 0x6b, 0x04, 0x07, 0x99, 0x59, 0x12,
	0xfd, 0x5e, 0x80, 0xc1, 0x54, 0xe7, 0xac, 0x3e, 0x5f, 0x40, 0x58, 0x61, 0x86, 0xbc, 0xe6, 0xe5,
	0x25, 0xe5, 0xa5, 0xc6, 0xea, 0x22, 0xcd, 0xa6, 0xa6, 0xd1, 0xda, 0xc4, 0x7e, 0xd0, 0x98, 0xff,
	0x3a, 0x9b, 0xc8, 0x3a, 0xb4, 0xa5, 0xa8, 0xb4, 0x31, 0xb9, 0x09, 0xd2, 0x32, 0xdb, 0x03, 0x46,
	0x9e, 0xc1, 0xca, 0x08, 0x53, 0x86, 0x15, 0x6d, 0x7a, 0xb5, 0x64, 0x29, 0x2e, 0x3b, 0xf4, 0xd4,
	0x77, 0x6c, 0x13, 0xba, 0x23, 0x91, 0x33, 0xaa, 0x79, 0x81, 0xbe, 0x08, 0x1d, 0x03, 0x9c, 0xf0,
	0x02, 0x49, 0x04, 0xfd, 0x2c, 0x95, 0xe9, 0x39, 0xcf, 0xb9, 0xe6, 0xa8, 0x7c, 0x21, 0xa6, 0x30,
	0xb2, 0x01, 0x1d, 0x99, 0xa7, 0xfa, 0x42, 0x54, 0x85, 0xaf, 0xc5, 0x64, 0x4f, 0x3e, 0x41, 0xcb,
	0xa9, 0x36, 0xec, 0x58, 0x01, 0xcc, 0x63, 0x44, 0xcc, 0xbc, 0x87, 0xc4, 0x67, 0x88, 0xae, 0x03,
	0x20, 0xb7, 0x47, 0x3f, 0x51, 0xd0, 0xbf, 0x89, 0x86, 0x43, 0xfb, 0x92, 0x8f, 0xe6, 0x79, 0x11,
	0xfb, 0x82, 0xcd, 0x1f, 0xe2, 0xd0, 0x23, 0xe7, 0x2d, 0xab, 0xc6, 0x9d, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x13, 0x21, 0x31, 0xe7, 0x06, 0x00, 0x00,
}