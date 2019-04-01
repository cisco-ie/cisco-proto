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
// source: ospfv3_sh_bad_checksum.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_bad_checksums_bad_checksum

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

type Ospfv3ShBadChecksum_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	PacketNumber         uint32   `protobuf:"varint,3,opt,name=packet_number,json=packetNumber,proto3" json:"packet_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3ShBadChecksum_KEYS) Reset()         { *m = Ospfv3ShBadChecksum_KEYS{} }
func (m *Ospfv3ShBadChecksum_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShBadChecksum_KEYS) ProtoMessage()    {}
func (*Ospfv3ShBadChecksum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_560d5f238b8b8509, []int{0}
}

func (m *Ospfv3ShBadChecksum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShBadChecksum_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3ShBadChecksum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShBadChecksum_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3ShBadChecksum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShBadChecksum_KEYS.Merge(m, src)
}
func (m *Ospfv3ShBadChecksum_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShBadChecksum_KEYS.Size(m)
}
func (m *Ospfv3ShBadChecksum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShBadChecksum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShBadChecksum_KEYS proto.InternalMessageInfo

func (m *Ospfv3ShBadChecksum_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3ShBadChecksum_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3ShBadChecksum_KEYS) GetPacketNumber() uint32 {
	if m != nil {
		return m.PacketNumber
	}
	return 0
}

type Ospfv3EdmTime struct {
	Second               uint32   `protobuf:"varint,1,opt,name=second,proto3" json:"second,omitempty"`
	Nanosecond           uint32   `protobuf:"varint,2,opt,name=nanosecond,proto3" json:"nanosecond,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmTime) Reset()         { *m = Ospfv3EdmTime{} }
func (m *Ospfv3EdmTime) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTime) ProtoMessage()    {}
func (*Ospfv3EdmTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_560d5f238b8b8509, []int{1}
}

func (m *Ospfv3EdmTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTime.Unmarshal(m, b)
}
func (m *Ospfv3EdmTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTime.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTime.Merge(m, src)
}
func (m *Ospfv3EdmTime) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTime.Size(m)
}
func (m *Ospfv3EdmTime) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTime.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTime proto.InternalMessageInfo

func (m *Ospfv3EdmTime) GetSecond() uint32 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *Ospfv3EdmTime) GetNanosecond() uint32 {
	if m != nil {
		return m.Nanosecond
	}
	return 0
}

type Ospfv3ShBadChecksum struct {
	ReceivedChecksum     uint32         `protobuf:"varint,50,opt,name=received_checksum,json=receivedChecksum,proto3" json:"received_checksum,omitempty"`
	ComputedChecksum     uint32         `protobuf:"varint,51,opt,name=computed_checksum,json=computedChecksum,proto3" json:"computed_checksum,omitempty"`
	Timestamp            *Ospfv3EdmTime `protobuf:"bytes,52,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ReceivedData         []byte         `protobuf:"bytes,53,opt,name=received_data,json=receivedData,proto3" json:"received_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Ospfv3ShBadChecksum) Reset()         { *m = Ospfv3ShBadChecksum{} }
func (m *Ospfv3ShBadChecksum) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShBadChecksum) ProtoMessage()    {}
func (*Ospfv3ShBadChecksum) Descriptor() ([]byte, []int) {
	return fileDescriptor_560d5f238b8b8509, []int{2}
}

func (m *Ospfv3ShBadChecksum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShBadChecksum.Unmarshal(m, b)
}
func (m *Ospfv3ShBadChecksum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShBadChecksum.Marshal(b, m, deterministic)
}
func (m *Ospfv3ShBadChecksum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShBadChecksum.Merge(m, src)
}
func (m *Ospfv3ShBadChecksum) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShBadChecksum.Size(m)
}
func (m *Ospfv3ShBadChecksum) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShBadChecksum.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShBadChecksum proto.InternalMessageInfo

func (m *Ospfv3ShBadChecksum) GetReceivedChecksum() uint32 {
	if m != nil {
		return m.ReceivedChecksum
	}
	return 0
}

func (m *Ospfv3ShBadChecksum) GetComputedChecksum() uint32 {
	if m != nil {
		return m.ComputedChecksum
	}
	return 0
}

func (m *Ospfv3ShBadChecksum) GetTimestamp() *Ospfv3EdmTime {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Ospfv3ShBadChecksum) GetReceivedData() []byte {
	if m != nil {
		return m.ReceivedData
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3ShBadChecksum_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.bad_checksums.bad_checksum.ospfv3_sh_bad_checksum_KEYS")
	proto.RegisterType((*Ospfv3EdmTime)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.bad_checksums.bad_checksum.ospfv3_edm_time")
	proto.RegisterType((*Ospfv3ShBadChecksum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.bad_checksums.bad_checksum.ospfv3_sh_bad_checksum")
}

func init() { proto.RegisterFile("ospfv3_sh_bad_checksum.proto", fileDescriptor_560d5f238b8b8509) }

var fileDescriptor_560d5f238b8b8509 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x84, 0xe9, 0xde, 0x3a, 0xd4, 0x1e, 0x46, 0x45, 0x91, 0x3a, 0x2f, 0x05, 0xa1,
	0x87, 0x4d, 0xfd, 0x07, 0xd4, 0x83, 0x08, 0x3b, 0xd4, 0x93, 0x5e, 0x42, 0x96, 0xbe, 0xb2, 0x32,
	0xd2, 0x84, 0x24, 0x2b, 0x1e, 0xbd, 0xfa, 0x77, 0xf8, 0x8f, 0xca, 0xd2, 0x74, 0xab, 0xb2, 0xab,
	0x97, 0xf2, 0xde, 0xf7, 0xf3, 0xe5, 0xfd, 0x6a, 0xe0, 0x42, 0x1a, 0x55, 0xd4, 0x33, 0x6a, 0x96,
	0x74, 0xc1, 0x72, 0xca, 0x97, 0xc8, 0x57, 0x66, 0x2d, 0x52, 0xa5, 0xa5, 0x95, 0xe1, 0x3b, 0x2f,
	0x0d, 0x97, 0xb4, 0x94, 0x86, 0x7e, 0x68, 0x5a, 0xaa, 0xfa, 0x9e, 0x7a, 0xbf, 0x54, 0xa8, 0xd3,
	0x26, 0xde, 0x78, 0x39, 0x1a, 0x83, 0xa6, 0x8d, 0xd2, 0x5a, 0x17, 0xee, 0x93, 0x76, 0x6b, 0x9a,
	0x5f, 0xd9, 0xe4, 0x93, 0xc0, 0xf9, 0xfe, 0xe6, 0xf4, 0xe5, 0xe9, 0xed, 0x35, 0xbc, 0x82, 0xc0,
	0x97, 0xa3, 0x15, 0x13, 0x18, 0x91, 0x98, 0x24, 0x83, 0x6c, 0xe8, 0xb5, 0x39, 0x13, 0x18, 0x9e,
	0xc1, 0x51, 0xad, 0x8b, 0x06, 0xf7, 0x1c, 0x3e, 0xac, 0x75, 0xe1, 0xd0, 0x35, 0x8c, 0x14, 0xe3,
	0x2b, 0xb4, 0xb4, 0x5a, 0x8b, 0x05, 0xea, 0xe8, 0x20, 0x26, 0xc9, 0x28, 0x0b, 0x1a, 0x71, 0xee,
	0xb4, 0xc9, 0x33, 0x1c, 0xfb, 0x09, 0x30, 0x17, 0xd4, 0x96, 0x02, 0xc3, 0x31, 0xf4, 0x0d, 0x72,
	0x59, 0xe5, 0xae, 0xdf, 0x28, 0xf3, 0x59, 0x78, 0x09, 0x50, 0xb1, 0x4a, 0x7a, 0xd6, 0x73, 0xac,
	0xa3, 0x4c, 0xbe, 0x7b, 0x30, 0xde, 0xbf, 0x4d, 0x78, 0x03, 0xa7, 0x1a, 0x39, 0x96, 0x35, 0xee,
	0xc4, 0x68, 0xea, 0x2a, 0x9c, 0xb4, 0xe0, 0xa1, 0x63, 0xe6, 0x52, 0xa8, 0xb5, 0xed, 0x9a, 0x67,
	0x8d, 0xb9, 0x05, 0x5b, 0xf3, 0x17, 0x81, 0xc1, 0x66, 0x6a, 0x63, 0x99, 0x50, 0xd1, 0x6d, 0x4c,
	0x92, 0xe1, 0x74, 0x95, 0xfe, 0xdf, 0x3f, 0x4b, 0xff, 0x5c, 0x2b, 0xdb, 0x75, 0xdf, 0x1c, 0x7c,
	0xbb, 0x65, 0xce, 0x2c, 0x8b, 0xee, 0x62, 0x92, 0x04, 0x59, 0xd0, 0x8a, 0x8f, 0xcc, 0xb2, 0x45,
	0xdf, 0x3d, 0xab, 0xd9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xb4, 0x95, 0xa4, 0x76, 0x02,
	0x00, 0x00,
}