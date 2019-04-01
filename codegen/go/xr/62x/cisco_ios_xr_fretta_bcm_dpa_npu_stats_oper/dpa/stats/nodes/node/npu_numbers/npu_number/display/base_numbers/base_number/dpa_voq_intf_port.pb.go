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
// source: dpa_voq_intf_port.proto

package cisco_ios_xr_fretta_bcm_dpa_npu_stats_oper_dpa_stats_nodes_node_npu_numbers_npu_number_display_base_numbers_base_number

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

type DpaVoqIntfPort_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	NpuId                uint32   `protobuf:"varint,2,opt,name=npu_id,json=npuId,proto3" json:"npu_id,omitempty"`
	BaseNumber           uint32   `protobuf:"varint,3,opt,name=base_number,json=baseNumber,proto3" json:"base_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DpaVoqIntfPort_KEYS) Reset()         { *m = DpaVoqIntfPort_KEYS{} }
func (m *DpaVoqIntfPort_KEYS) String() string { return proto.CompactTextString(m) }
func (*DpaVoqIntfPort_KEYS) ProtoMessage()    {}
func (*DpaVoqIntfPort_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9db590a833843dd3, []int{0}
}

func (m *DpaVoqIntfPort_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaVoqIntfPort_KEYS.Unmarshal(m, b)
}
func (m *DpaVoqIntfPort_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaVoqIntfPort_KEYS.Marshal(b, m, deterministic)
}
func (m *DpaVoqIntfPort_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaVoqIntfPort_KEYS.Merge(m, src)
}
func (m *DpaVoqIntfPort_KEYS) XXX_Size() int {
	return xxx_messageInfo_DpaVoqIntfPort_KEYS.Size(m)
}
func (m *DpaVoqIntfPort_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaVoqIntfPort_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_DpaVoqIntfPort_KEYS proto.InternalMessageInfo

func (m *DpaVoqIntfPort_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *DpaVoqIntfPort_KEYS) GetNpuId() uint32 {
	if m != nil {
		return m.NpuId
	}
	return 0
}

func (m *DpaVoqIntfPort_KEYS) GetBaseNumber() uint32 {
	if m != nil {
		return m.BaseNumber
	}
	return 0
}

type DpaVoqIngressPacketStatistics struct {
	ReceivedBytes        uint64   `protobuf:"varint,1,opt,name=received_bytes,json=receivedBytes,proto3" json:"received_bytes,omitempty"`
	ReceivedPackets      uint64   `protobuf:"varint,2,opt,name=received_packets,json=receivedPackets,proto3" json:"received_packets,omitempty"`
	DroppedBytes         uint64   `protobuf:"varint,3,opt,name=dropped_bytes,json=droppedBytes,proto3" json:"dropped_bytes,omitempty"`
	DroppedPackets       uint64   `protobuf:"varint,4,opt,name=dropped_packets,json=droppedPackets,proto3" json:"dropped_packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DpaVoqIngressPacketStatistics) Reset()         { *m = DpaVoqIngressPacketStatistics{} }
func (m *DpaVoqIngressPacketStatistics) String() string { return proto.CompactTextString(m) }
func (*DpaVoqIngressPacketStatistics) ProtoMessage()    {}
func (*DpaVoqIngressPacketStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_9db590a833843dd3, []int{1}
}

func (m *DpaVoqIngressPacketStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaVoqIngressPacketStatistics.Unmarshal(m, b)
}
func (m *DpaVoqIngressPacketStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaVoqIngressPacketStatistics.Marshal(b, m, deterministic)
}
func (m *DpaVoqIngressPacketStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaVoqIngressPacketStatistics.Merge(m, src)
}
func (m *DpaVoqIngressPacketStatistics) XXX_Size() int {
	return xxx_messageInfo_DpaVoqIngressPacketStatistics.Size(m)
}
func (m *DpaVoqIngressPacketStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaVoqIngressPacketStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_DpaVoqIngressPacketStatistics proto.InternalMessageInfo

func (m *DpaVoqIngressPacketStatistics) GetReceivedBytes() uint64 {
	if m != nil {
		return m.ReceivedBytes
	}
	return 0
}

func (m *DpaVoqIngressPacketStatistics) GetReceivedPackets() uint64 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *DpaVoqIngressPacketStatistics) GetDroppedBytes() uint64 {
	if m != nil {
		return m.DroppedBytes
	}
	return 0
}

func (m *DpaVoqIngressPacketStatistics) GetDroppedPackets() uint64 {
	if m != nil {
		return m.DroppedPackets
	}
	return 0
}

type DpaVoqIntfPort struct {
	InUse                bool                             `protobuf:"varint,50,opt,name=in_use,json=inUse,proto3" json:"in_use,omitempty"`
	RackNum              uint32                           `protobuf:"varint,51,opt,name=rack_num,json=rackNum,proto3" json:"rack_num,omitempty"`
	SlotNum              uint32                           `protobuf:"varint,52,opt,name=slot_num,json=slotNum,proto3" json:"slot_num,omitempty"`
	NpuNum               uint32                           `protobuf:"varint,53,opt,name=npu_num,json=npuNum,proto3" json:"npu_num,omitempty"`
	NpuCore              uint32                           `protobuf:"varint,54,opt,name=npu_core,json=npuCore,proto3" json:"npu_core,omitempty"`
	PortNum              uint32                           `protobuf:"varint,55,opt,name=port_num,json=portNum,proto3" json:"port_num,omitempty"`
	IfHandle             uint32                           `protobuf:"varint,56,opt,name=if_handle,json=ifHandle,proto3" json:"if_handle,omitempty"`
	SysPort              uint32                           `protobuf:"varint,57,opt,name=sys_port,json=sysPort,proto3" json:"sys_port,omitempty"`
	PpPort               uint32                           `protobuf:"varint,58,opt,name=pp_port,json=ppPort,proto3" json:"pp_port,omitempty"`
	PortSpeed            uint32                           `protobuf:"varint,59,opt,name=port_speed,json=portSpeed,proto3" json:"port_speed,omitempty"`
	VoqBase              uint32                           `protobuf:"varint,60,opt,name=voq_base,json=voqBase,proto3" json:"voq_base,omitempty"`
	ConnectorId          uint32                           `protobuf:"varint,61,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	IsLocalPort          bool                             `protobuf:"varint,62,opt,name=is_local_port,json=isLocalPort,proto3" json:"is_local_port,omitempty"`
	VoqStat              []*DpaVoqIngressPacketStatistics `protobuf:"bytes,63,rep,name=voq_stat,json=voqStat,proto3" json:"voq_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *DpaVoqIntfPort) Reset()         { *m = DpaVoqIntfPort{} }
func (m *DpaVoqIntfPort) String() string { return proto.CompactTextString(m) }
func (*DpaVoqIntfPort) ProtoMessage()    {}
func (*DpaVoqIntfPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_9db590a833843dd3, []int{2}
}

func (m *DpaVoqIntfPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaVoqIntfPort.Unmarshal(m, b)
}
func (m *DpaVoqIntfPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaVoqIntfPort.Marshal(b, m, deterministic)
}
func (m *DpaVoqIntfPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaVoqIntfPort.Merge(m, src)
}
func (m *DpaVoqIntfPort) XXX_Size() int {
	return xxx_messageInfo_DpaVoqIntfPort.Size(m)
}
func (m *DpaVoqIntfPort) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaVoqIntfPort.DiscardUnknown(m)
}

var xxx_messageInfo_DpaVoqIntfPort proto.InternalMessageInfo

func (m *DpaVoqIntfPort) GetInUse() bool {
	if m != nil {
		return m.InUse
	}
	return false
}

func (m *DpaVoqIntfPort) GetRackNum() uint32 {
	if m != nil {
		return m.RackNum
	}
	return 0
}

func (m *DpaVoqIntfPort) GetSlotNum() uint32 {
	if m != nil {
		return m.SlotNum
	}
	return 0
}

func (m *DpaVoqIntfPort) GetNpuNum() uint32 {
	if m != nil {
		return m.NpuNum
	}
	return 0
}

func (m *DpaVoqIntfPort) GetNpuCore() uint32 {
	if m != nil {
		return m.NpuCore
	}
	return 0
}

func (m *DpaVoqIntfPort) GetPortNum() uint32 {
	if m != nil {
		return m.PortNum
	}
	return 0
}

func (m *DpaVoqIntfPort) GetIfHandle() uint32 {
	if m != nil {
		return m.IfHandle
	}
	return 0
}

func (m *DpaVoqIntfPort) GetSysPort() uint32 {
	if m != nil {
		return m.SysPort
	}
	return 0
}

func (m *DpaVoqIntfPort) GetPpPort() uint32 {
	if m != nil {
		return m.PpPort
	}
	return 0
}

func (m *DpaVoqIntfPort) GetPortSpeed() uint32 {
	if m != nil {
		return m.PortSpeed
	}
	return 0
}

func (m *DpaVoqIntfPort) GetVoqBase() uint32 {
	if m != nil {
		return m.VoqBase
	}
	return 0
}

func (m *DpaVoqIntfPort) GetConnectorId() uint32 {
	if m != nil {
		return m.ConnectorId
	}
	return 0
}

func (m *DpaVoqIntfPort) GetIsLocalPort() bool {
	if m != nil {
		return m.IsLocalPort
	}
	return false
}

func (m *DpaVoqIntfPort) GetVoqStat() []*DpaVoqIngressPacketStatistics {
	if m != nil {
		return m.VoqStat
	}
	return nil
}

func init() {
	proto.RegisterType((*DpaVoqIntfPort_KEYS)(nil), "cisco_ios_xr_fretta_bcm_dpa_npu_stats_oper.dpa.stats.nodes.node.npu_numbers.npu_number.display.base_numbers.base_number.dpa_voq_intf_port_KEYS")
	proto.RegisterType((*DpaVoqIngressPacketStatistics)(nil), "cisco_ios_xr_fretta_bcm_dpa_npu_stats_oper.dpa.stats.nodes.node.npu_numbers.npu_number.display.base_numbers.base_number.dpa_voq_ingress_packet_statistics")
	proto.RegisterType((*DpaVoqIntfPort)(nil), "cisco_ios_xr_fretta_bcm_dpa_npu_stats_oper.dpa.stats.nodes.node.npu_numbers.npu_number.display.base_numbers.base_number.dpa_voq_intf_port")
}

func init() { proto.RegisterFile("dpa_voq_intf_port.proto", fileDescriptor_9db590a833843dd3) }

var fileDescriptor_9db590a833843dd3 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x4f, 0x6f, 0x13, 0x3f,
	0x10, 0xd5, 0xfe, 0x92, 0x36, 0x89, 0xd3, 0xb4, 0x3f, 0x56, 0x82, 0x18, 0x55, 0x88, 0x34, 0x08,
	0x11, 0x2e, 0x7b, 0x68, 0xf9, 0xff, 0x57, 0x2a, 0x42, 0xa2, 0x02, 0x45, 0x55, 0x22, 0x0e, 0x9c,
	0x2c, 0x67, 0xd7, 0x01, 0xab, 0x59, 0xdb, 0xf5, 0x78, 0x03, 0xf9, 0x2a, 0x7c, 0x10, 0x0e, 0x5c,
	0xf9, 0x62, 0x68, 0xc6, 0x9b, 0x10, 0xa9, 0x07, 0x8e, 0x5c, 0x56, 0x3b, 0xef, 0xcd, 0xbc, 0x37,
	0x33, 0xb6, 0x59, 0xbf, 0x70, 0x52, 0x2c, 0xed, 0xa5, 0xd0, 0x26, 0xcc, 0x85, 0xb3, 0x3e, 0x64,
	0xce, 0xdb, 0x60, 0xd3, 0xaf, 0xb9, 0x86, 0xdc, 0x0a, 0x6d, 0x41, 0x7c, 0xf3, 0x62, 0xee, 0x55,
	0x08, 0x52, 0xcc, 0xf2, 0x52, 0x60, 0x81, 0x71, 0x95, 0x80, 0x20, 0x03, 0x08, 0xeb, 0x94, 0xcf,
	0x0a, 0x27, 0x33, 0x0a, 0x33, 0x63, 0x0b, 0x15, 0xbf, 0x19, 0xe6, 0x98, 0xaa, 0x9c, 0x29, 0x0f,
	0x5b, 0xff, 0x59, 0xa1, 0xc1, 0x2d, 0xe4, 0x2a, 0x9b, 0x49, 0x50, 0x1b, 0x7e, 0x2b, 0x18, 0x96,
	0xec, 0xc6, 0x95, 0x9e, 0xc4, 0xfb, 0xb7, 0x9f, 0xa6, 0xe9, 0x21, 0xeb, 0xa0, 0xb2, 0x30, 0xb2,
	0x54, 0x3c, 0x19, 0x24, 0xa3, 0xce, 0xa4, 0x8d, 0xc0, 0x58, 0x96, 0x2a, 0xbd, 0xce, 0x76, 0xd1,
	0x45, 0x17, 0xfc, 0xbf, 0x41, 0x32, 0xea, 0x4d, 0x76, 0x8c, 0xab, 0xce, 0x8a, 0xf4, 0x36, 0xeb,
	0x6e, 0x89, 0xf3, 0x06, 0x71, 0x0c, 0xa1, 0x71, 0xb4, 0xfb, 0x95, 0xb0, 0xa3, 0x3f, 0x7e, 0x9f,
	0xbd, 0x02, 0x10, 0x4e, 0xe6, 0x17, 0x2a, 0xd0, 0x84, 0x1a, 0x82, 0xce, 0x21, 0xbd, 0xcb, 0xf6,
	0xbd, 0xca, 0x95, 0x5e, 0xaa, 0x42, 0xcc, 0x56, 0x41, 0x01, 0xf9, 0x37, 0x27, 0xbd, 0x35, 0x7a,
	0x8a, 0x60, 0x7a, 0x9f, 0xfd, 0xbf, 0x49, 0x8b, 0x22, 0x40, 0xed, 0x34, 0x27, 0x07, 0x6b, 0xfc,
	0x3c, 0xc2, 0xe9, 0x1d, 0xd6, 0x2b, 0xbc, 0x75, 0x6e, 0x23, 0xd8, 0xa0, 0xbc, 0xbd, 0x1a, 0x8c,
	0x7a, 0xf7, 0xd8, 0xc1, 0x3a, 0x69, 0x2d, 0xd7, 0xa4, 0xb4, 0xfd, 0x1a, 0xae, 0xd5, 0x86, 0x3f,
	0x9b, 0xec, 0xda, 0x95, 0xad, 0xe1, 0x4e, 0xb4, 0x11, 0x15, 0x28, 0x7e, 0x3c, 0x48, 0x46, 0xed,
	0xc9, 0x8e, 0x36, 0x1f, 0x41, 0xa5, 0x37, 0x59, 0xdb, 0xcb, 0xfc, 0x02, 0x77, 0xc2, 0x4f, 0x68,
	0x21, 0x2d, 0x8c, 0xc7, 0x55, 0x89, 0x14, 0x2c, 0x6c, 0x20, 0xea, 0x41, 0xa4, 0x30, 0x46, 0xaa,
	0xcf, 0x5a, 0xf5, 0x31, 0xf2, 0x87, 0xc4, 0xe0, 0xbe, 0xeb, 0x1a, 0x24, 0x72, 0xeb, 0x15, 0x7f,
	0x14, 0x6b, 0x8c, 0xab, 0xde, 0x58, 0x4f, 0x4e, 0x74, 0x7c, 0x58, 0xf4, 0x38, 0x52, 0x18, 0x63,
	0xd5, 0x21, 0xeb, 0xe8, 0xb9, 0xf8, 0x22, 0x4d, 0xb1, 0x50, 0xfc, 0x09, 0x71, 0x6d, 0x3d, 0x7f,
	0x47, 0x31, 0xb5, 0xb1, 0x02, 0x1a, 0x82, 0x3f, 0xad, 0xdb, 0x58, 0xc1, 0x39, 0xce, 0xd4, 0x67,
	0x2d, 0xe7, 0x22, 0xf3, 0x2c, 0xb6, 0xe1, 0x1c, 0x11, 0xb7, 0x18, 0x23, 0x2f, 0x70, 0x4a, 0x15,
	0xfc, 0x39, 0x71, 0x1d, 0x44, 0xa6, 0x08, 0xa0, 0x24, 0x2e, 0x07, 0x4f, 0x9e, 0xbf, 0x88, 0x92,
	0x4b, 0x7b, 0x79, 0x2a, 0x41, 0xa5, 0x47, 0x6c, 0x2f, 0xb7, 0xc6, 0xa8, 0x3c, 0x58, 0x8f, 0x17,
	0xe8, 0x25, 0xd1, 0xdd, 0x0d, 0x76, 0x56, 0xa4, 0x43, 0xd6, 0xd3, 0x20, 0x16, 0x36, 0x97, 0x8b,
	0xe8, 0xfd, 0x8a, 0x16, 0xda, 0xd5, 0xf0, 0x01, 0x31, 0x6a, 0xe0, 0x47, 0x12, 0x2d, 0xf0, 0xda,
	0xf0, 0xd7, 0x83, 0xc6, 0xa8, 0x7b, 0xfc, 0x3d, 0xc9, 0xfe, 0xd1, 0x33, 0xca, 0xfe, 0x7a, 0xa7,
	0x69, 0xfe, 0x69, 0x90, 0x61, 0xb6, 0x4b, 0x2f, 0xfe, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5c, 0x98, 0x9c, 0x07, 0x0c, 0x04, 0x00, 0x00,
}