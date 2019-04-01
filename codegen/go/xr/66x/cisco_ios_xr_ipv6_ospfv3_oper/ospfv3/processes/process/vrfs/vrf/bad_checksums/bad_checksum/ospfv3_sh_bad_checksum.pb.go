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
	ReceivedData         string         `protobuf:"bytes,53,opt,name=received_data,json=receivedData,proto3" json:"received_data,omitempty"`
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

func (m *Ospfv3ShBadChecksum) GetReceivedData() string {
	if m != nil {
		return m.ReceivedData
	}
	return ""
}

func init() {
	proto.RegisterType((*Ospfv3ShBadChecksum_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.bad_checksums.bad_checksum.ospfv3_sh_bad_checksum_KEYS")
	proto.RegisterType((*Ospfv3EdmTime)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.bad_checksums.bad_checksum.ospfv3_edm_time")
	proto.RegisterType((*Ospfv3ShBadChecksum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.bad_checksums.bad_checksum.ospfv3_sh_bad_checksum")
}

func init() { proto.RegisterFile("ospfv3_sh_bad_checksum.proto", fileDescriptor_560d5f238b8b8509) }

var fileDescriptor_560d5f238b8b8509 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0x84, 0xe9, 0xde, 0x6d, 0xa8, 0x3d, 0x8c, 0x8a, 0x22, 0x73, 0x5e, 0x06, 0x42,
	0x0e, 0x9b, 0xfa, 0x05, 0xd4, 0x83, 0x08, 0x3b, 0xd4, 0x93, 0x5e, 0x42, 0x96, 0xbe, 0x65, 0x65,
	0xa4, 0x09, 0x49, 0x16, 0x3c, 0x7a, 0xf5, 0x73, 0xf8, 0x45, 0xa5, 0x69, 0xba, 0x55, 0xd9, 0xd5,
	0x4b, 0xc9, 0xfb, 0xfc, 0x1e, 0xde, 0xbf, 0x85, 0x0b, 0x69, 0x54, 0xee, 0xe6, 0xd4, 0xac, 0xe8,
	0x92, 0x65, 0x94, 0xaf, 0x90, 0xaf, 0xcd, 0x46, 0x10, 0xa5, 0xa5, 0x95, 0xf1, 0x3b, 0x2f, 0x0c,
	0x97, 0xb4, 0x90, 0x86, 0x7e, 0x68, 0x5a, 0x28, 0x77, 0x4f, 0x83, 0x5f, 0x2a, 0xd4, 0xa4, 0x7e,
	0x57, 0x5e, 0x8e, 0xc6, 0xa0, 0x69, 0x5e, 0xc4, 0xe9, 0xdc, 0x7f, 0x48, 0x3b, 0xa7, 0xf9, 0x15,
	0x4d, 0x3e, 0x23, 0x38, 0xdf, 0x5f, 0x9c, 0xbe, 0x3c, 0xbd, 0xbd, 0xc6, 0x57, 0x30, 0x08, 0xe9,
	0x68, 0xc9, 0x04, 0x26, 0xd1, 0x38, 0x9a, 0xf6, 0xd2, 0x7e, 0xd0, 0x16, 0x4c, 0x60, 0x7c, 0x06,
	0x47, 0x4e, 0xe7, 0x35, 0xee, 0x78, 0x7c, 0xe8, 0x74, 0xee, 0xd1, 0x35, 0x0c, 0x15, 0xe3, 0x6b,
	0xb4, 0xb4, 0xdc, 0x88, 0x25, 0xea, 0xe4, 0x60, 0x1c, 0x4d, 0x87, 0xe9, 0xa0, 0x16, 0x17, 0x5e,
	0x9b, 0x3c, 0xc3, 0x71, 0xe8, 0x00, 0x33, 0x41, 0x6d, 0x21, 0x30, 0x1e, 0x41, 0xd7, 0x20, 0x97,
	0x65, 0xe6, 0xeb, 0x0d, 0xd3, 0x10, 0xc5, 0x97, 0x00, 0x25, 0x2b, 0x65, 0x60, 0x1d, 0xcf, 0x5a,
	0xca, 0xe4, 0xbb, 0x03, 0xa3, 0xfd, 0xd3, 0xc4, 0x37, 0x70, 0xaa, 0x91, 0x63, 0xe1, 0x70, 0x27,
	0x26, 0x33, 0x9f, 0xe1, 0xa4, 0x01, 0x0f, 0x2d, 0x33, 0x97, 0x42, 0x6d, 0x6c, 0xdb, 0x3c, 0xaf,
	0xcd, 0x0d, 0xd8, 0x9a, 0xbf, 0x22, 0xe8, 0x55, 0x5d, 0x1b, 0xcb, 0x84, 0x4a, 0x6e, 0xc7, 0xd1,
	0xb4, 0x3f, 0x5b, 0x93, 0xff, 0xbb, 0x19, 0xf9, 0xb3, 0xad, 0x74, 0x57, 0xbd, 0x5a, 0xf8, 0x76,
	0xca, 0x8c, 0x59, 0x96, 0xdc, 0xf9, 0x83, 0x0c, 0x1a, 0xf1, 0x91, 0x59, 0xb6, 0xec, 0xfa, 0xdf,
	0x6a, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xff, 0xba, 0x3f, 0x76, 0x02, 0x00, 0x00,
}
