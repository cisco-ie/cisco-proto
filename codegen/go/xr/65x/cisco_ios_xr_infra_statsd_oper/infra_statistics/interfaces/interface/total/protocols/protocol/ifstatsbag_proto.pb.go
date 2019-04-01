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
// source: ifstatsbag_proto.proto

package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_total_protocols_protocol

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

type IfstatsbagProto_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	ProtocolName         string   `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IfstatsbagProto_KEYS) Reset()         { *m = IfstatsbagProto_KEYS{} }
func (m *IfstatsbagProto_KEYS) String() string { return proto.CompactTextString(m) }
func (*IfstatsbagProto_KEYS) ProtoMessage()    {}
func (*IfstatsbagProto_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e305d2d18a94ee6, []int{0}
}

func (m *IfstatsbagProto_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfstatsbagProto_KEYS.Unmarshal(m, b)
}
func (m *IfstatsbagProto_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfstatsbagProto_KEYS.Marshal(b, m, deterministic)
}
func (m *IfstatsbagProto_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfstatsbagProto_KEYS.Merge(m, src)
}
func (m *IfstatsbagProto_KEYS) XXX_Size() int {
	return xxx_messageInfo_IfstatsbagProto_KEYS.Size(m)
}
func (m *IfstatsbagProto_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IfstatsbagProto_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IfstatsbagProto_KEYS proto.InternalMessageInfo

func (m *IfstatsbagProto_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *IfstatsbagProto_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

type IfstatsbagProto struct {
	BytesReceived        uint64   `protobuf:"varint,50,opt,name=bytes_received,json=bytesReceived,proto3" json:"bytes_received,omitempty"`
	PacketsReceived      uint64   `protobuf:"varint,51,opt,name=packets_received,json=packetsReceived,proto3" json:"packets_received,omitempty"`
	BytesSent            uint64   `protobuf:"varint,52,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	PacketsSent          uint64   `protobuf:"varint,53,opt,name=packets_sent,json=packetsSent,proto3" json:"packets_sent,omitempty"`
	Protocol             uint32   `protobuf:"varint,54,opt,name=protocol,proto3" json:"protocol,omitempty"`
	LastDataTime         uint64   `protobuf:"varint,55,opt,name=last_data_time,json=lastDataTime,proto3" json:"last_data_time,omitempty"`
	InputDataRate        uint64   `protobuf:"varint,56,opt,name=input_data_rate,json=inputDataRate,proto3" json:"input_data_rate,omitempty"`
	InputPacketRate      uint64   `protobuf:"varint,57,opt,name=input_packet_rate,json=inputPacketRate,proto3" json:"input_packet_rate,omitempty"`
	OutputDataRate       uint64   `protobuf:"varint,58,opt,name=output_data_rate,json=outputDataRate,proto3" json:"output_data_rate,omitempty"`
	OutputPacketRate     uint64   `protobuf:"varint,59,opt,name=output_packet_rate,json=outputPacketRate,proto3" json:"output_packet_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IfstatsbagProto) Reset()         { *m = IfstatsbagProto{} }
func (m *IfstatsbagProto) String() string { return proto.CompactTextString(m) }
func (*IfstatsbagProto) ProtoMessage()    {}
func (*IfstatsbagProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e305d2d18a94ee6, []int{1}
}

func (m *IfstatsbagProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfstatsbagProto.Unmarshal(m, b)
}
func (m *IfstatsbagProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfstatsbagProto.Marshal(b, m, deterministic)
}
func (m *IfstatsbagProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfstatsbagProto.Merge(m, src)
}
func (m *IfstatsbagProto) XXX_Size() int {
	return xxx_messageInfo_IfstatsbagProto.Size(m)
}
func (m *IfstatsbagProto) XXX_DiscardUnknown() {
	xxx_messageInfo_IfstatsbagProto.DiscardUnknown(m)
}

var xxx_messageInfo_IfstatsbagProto proto.InternalMessageInfo

func (m *IfstatsbagProto) GetBytesReceived() uint64 {
	if m != nil {
		return m.BytesReceived
	}
	return 0
}

func (m *IfstatsbagProto) GetPacketsReceived() uint64 {
	if m != nil {
		return m.PacketsReceived
	}
	return 0
}

func (m *IfstatsbagProto) GetBytesSent() uint64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *IfstatsbagProto) GetPacketsSent() uint64 {
	if m != nil {
		return m.PacketsSent
	}
	return 0
}

func (m *IfstatsbagProto) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *IfstatsbagProto) GetLastDataTime() uint64 {
	if m != nil {
		return m.LastDataTime
	}
	return 0
}

func (m *IfstatsbagProto) GetInputDataRate() uint64 {
	if m != nil {
		return m.InputDataRate
	}
	return 0
}

func (m *IfstatsbagProto) GetInputPacketRate() uint64 {
	if m != nil {
		return m.InputPacketRate
	}
	return 0
}

func (m *IfstatsbagProto) GetOutputDataRate() uint64 {
	if m != nil {
		return m.OutputDataRate
	}
	return 0
}

func (m *IfstatsbagProto) GetOutputPacketRate() uint64 {
	if m != nil {
		return m.OutputPacketRate
	}
	return 0
}

func init() {
	proto.RegisterType((*IfstatsbagProto_KEYS)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.total.protocols.protocol.ifstatsbag_proto_KEYS")
	proto.RegisterType((*IfstatsbagProto)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.total.protocols.protocol.ifstatsbag_proto")
}

func init() { proto.RegisterFile("ifstatsbag_proto.proto", fileDescriptor_4e305d2d18a94ee6) }

var fileDescriptor_4e305d2d18a94ee6 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x40, 0x51, 0x5b, 0x4a, 0xbd, 0xb5, 0x6c, 0x77, 0xa1, 0x45, 0x14, 0x0a, 0xae, 0xf3, 0x81,
	0x12, 0x82, 0x0e, 0x71, 0xbe, 0x73, 0x4d, 0x4e, 0x81, 0x10, 0xe4, 0x5c, 0x72, 0x08, 0xcb, 0x78,
	0x3d, 0x0e, 0x4b, 0x2c, 0xad, 0xd0, 0x8e, 0x43, 0xf2, 0x3f, 0xf2, 0x83, 0x83, 0x66, 0x25, 0xd9,
	0xf1, 0x45, 0xac, 0x9e, 0x9e, 0xde, 0xec, 0x88, 0x3f, 0x66, 0xee, 0x08, 0xc8, 0x4d, 0xe1, 0x49,
	0x15, 0xa5, 0x25, 0x9b, 0xf0, 0x53, 0x3e, 0x6a, 0xe3, 0xb4, 0x55, 0xc6, 0x3a, 0xf5, 0x5a, 0x2a,
	0x93, 0xcf, 0x4b, 0x50, 0x2c, 0xce, 0x94, 0x2d, 0xb0, 0x4c, 0x56, 0xc4, 0x38, 0x32, 0xda, 0x25,
	0x26, 0x27, 0x2c, 0xe7, 0xa0, 0x71, 0xed, 0x98, 0x90, 0x25, 0x58, 0xf8, 0xa2, 0xb6, 0x0b, 0xd7,
	0x9e, 0x46, 0x5a, 0xfc, 0xde, 0x1c, 0xac, 0x6e, 0xae, 0x1f, 0x26, 0x72, 0x47, 0xf4, 0xda, 0xdf,
	0x55, 0x0e, 0x19, 0x46, 0xc1, 0x30, 0x88, 0x3b, 0x69, 0xd8, 0xd2, 0x5b, 0xc8, 0x50, 0x6e, 0x89,
	0xb0, 0x69, 0x79, 0xeb, 0x0b, 0x5b, 0xdd, 0x06, 0x56, 0xd2, 0xe8, 0xfd, 0xab, 0x18, 0x6c, 0x4e,
	0xa9, 0x06, 0x4c, 0xdf, 0x08, 0x9d, 0x2a, 0x51, 0xa3, 0x79, 0xc1, 0x59, 0x74, 0x38, 0x0c, 0xe2,
	0x6f, 0x69, 0xc8, 0x34, 0xad, 0xa1, 0xdc, 0x13, 0x83, 0x02, 0xf4, 0x33, 0xd2, 0x9a, 0x38, 0x66,
	0xb1, 0x5f, 0xf3, 0x56, 0xfd, 0x27, 0x84, 0x2f, 0x3a, 0xcc, 0x29, 0x3a, 0x62, 0xa9, 0xc3, 0x64,
	0x82, 0x39, 0xc9, 0xff, 0xa2, 0xdb, 0x94, 0x58, 0x38, 0x66, 0xe1, 0x67, 0xcd, 0x58, 0xf9, 0x2b,
	0x7e, 0x34, 0x17, 0x8f, 0x4e, 0x86, 0x41, 0x1c, 0xa6, 0xed, 0xbb, 0xdc, 0x16, 0xbd, 0x05, 0x38,
	0x52, 0x33, 0x20, 0x50, 0x64, 0x32, 0x8c, 0x4e, 0x39, 0xd0, 0xad, 0xe8, 0x15, 0x10, 0xdc, 0x9b,
	0x0c, 0xe5, 0xae, 0xe8, 0x9b, 0xbc, 0x58, 0xd6, 0x5a, 0x09, 0x84, 0xd1, 0x99, 0x5f, 0x8b, 0x71,
	0xe5, 0xa5, 0x40, 0x28, 0xf7, 0xc5, 0x2f, 0xef, 0xf9, 0xf1, 0xde, 0x3c, 0xf7, 0x7b, 0xf1, 0x87,
	0x3b, 0xe6, 0xec, 0xc6, 0x62, 0x60, 0x97, 0xf4, 0x39, 0x7a, 0xc1, 0x6a, 0xcf, 0xf3, 0xb6, 0x7a,
	0x20, 0x64, 0x6d, 0xae, 0x67, 0x2f, 0xd9, 0xad, 0x1b, 0xab, 0xee, 0xf4, 0x3b, 0xef, 0x36, 0xfe,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x62, 0x53, 0x97, 0x7b, 0x02, 0x00, 0x00,
}
