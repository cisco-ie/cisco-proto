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
// source: lldp_mgr_global_info.proto

package cisco_ios_xr_ethernet_lldp_oper_lldp_global_lldp_lldp_info

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

type LldpMgrGlobalInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LldpMgrGlobalInfo_KEYS) Reset()         { *m = LldpMgrGlobalInfo_KEYS{} }
func (m *LldpMgrGlobalInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LldpMgrGlobalInfo_KEYS) ProtoMessage()    {}
func (*LldpMgrGlobalInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_74d8998007aaef94, []int{0}
}

func (m *LldpMgrGlobalInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpMgrGlobalInfo_KEYS.Unmarshal(m, b)
}
func (m *LldpMgrGlobalInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpMgrGlobalInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LldpMgrGlobalInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpMgrGlobalInfo_KEYS.Merge(m, src)
}
func (m *LldpMgrGlobalInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LldpMgrGlobalInfo_KEYS.Size(m)
}
func (m *LldpMgrGlobalInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpMgrGlobalInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LldpMgrGlobalInfo_KEYS proto.InternalMessageInfo

type LldpMgrGlobalInfo struct {
	ChassisId            string   `protobuf:"bytes,50,opt,name=chassis_id,json=chassisId,proto3" json:"chassis_id,omitempty"`
	ChassisIdSubType     uint32   `protobuf:"varint,51,opt,name=chassis_id_sub_type,json=chassisIdSubType,proto3" json:"chassis_id_sub_type,omitempty"`
	SystemName           string   `protobuf:"bytes,52,opt,name=system_name,json=systemName,proto3" json:"system_name,omitempty"`
	Timer                uint32   `protobuf:"varint,53,opt,name=timer,proto3" json:"timer,omitempty"`
	HoldTime             uint32   `protobuf:"varint,54,opt,name=hold_time,json=holdTime,proto3" json:"hold_time,omitempty"`
	ReInit               uint32   `protobuf:"varint,55,opt,name=re_init,json=reInit,proto3" json:"re_init,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LldpMgrGlobalInfo) Reset()         { *m = LldpMgrGlobalInfo{} }
func (m *LldpMgrGlobalInfo) String() string { return proto.CompactTextString(m) }
func (*LldpMgrGlobalInfo) ProtoMessage()    {}
func (*LldpMgrGlobalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_74d8998007aaef94, []int{1}
}

func (m *LldpMgrGlobalInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpMgrGlobalInfo.Unmarshal(m, b)
}
func (m *LldpMgrGlobalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpMgrGlobalInfo.Marshal(b, m, deterministic)
}
func (m *LldpMgrGlobalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpMgrGlobalInfo.Merge(m, src)
}
func (m *LldpMgrGlobalInfo) XXX_Size() int {
	return xxx_messageInfo_LldpMgrGlobalInfo.Size(m)
}
func (m *LldpMgrGlobalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpMgrGlobalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LldpMgrGlobalInfo proto.InternalMessageInfo

func (m *LldpMgrGlobalInfo) GetChassisId() string {
	if m != nil {
		return m.ChassisId
	}
	return ""
}

func (m *LldpMgrGlobalInfo) GetChassisIdSubType() uint32 {
	if m != nil {
		return m.ChassisIdSubType
	}
	return 0
}

func (m *LldpMgrGlobalInfo) GetSystemName() string {
	if m != nil {
		return m.SystemName
	}
	return ""
}

func (m *LldpMgrGlobalInfo) GetTimer() uint32 {
	if m != nil {
		return m.Timer
	}
	return 0
}

func (m *LldpMgrGlobalInfo) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *LldpMgrGlobalInfo) GetReInit() uint32 {
	if m != nil {
		return m.ReInit
	}
	return 0
}

func init() {
	proto.RegisterType((*LldpMgrGlobalInfo_KEYS)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.global_lldp.lldp_info.lldp_mgr_global_info_KEYS")
	proto.RegisterType((*LldpMgrGlobalInfo)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.global_lldp.lldp_info.lldp_mgr_global_info")
}

func init() { proto.RegisterFile("lldp_mgr_global_info.proto", fileDescriptor_74d8998007aaef94) }

var fileDescriptor_74d8998007aaef94 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3b, 0x4b, 0x03, 0x41,
	0x14, 0x85, 0xd9, 0xc2, 0xe8, 0x5e, 0x11, 0x64, 0x0c, 0x38, 0x1a, 0xc4, 0x90, 0x2a, 0x8d, 0x5b,
	0x18, 0x1f, 0x60, 0x6f, 0x11, 0x04, 0x8b, 0x24, 0x8d, 0xd5, 0x65, 0x1f, 0xd7, 0xec, 0x85, 0x79,
	0x31, 0x33, 0x01, 0xf7, 0x67, 0xfa, 0x8f, 0x64, 0x27, 0xc1, 0x34, 0xdb, 0xdd, 0x73, 0xbe, 0x39,
	0x5f, 0x31, 0x70, 0xab, 0x54, 0xe3, 0x50, 0x6f, 0x3d, 0x6e, 0x95, 0xad, 0x4a, 0x85, 0x6c, 0xbe,
	0x6d, 0xe1, 0xbc, 0x8d, 0x56, 0xbc, 0xd5, 0x1c, 0x6a, 0x8b, 0x6c, 0x03, 0xfe, 0x78, 0xa4, 0xd8,
	0x92, 0x37, 0x14, 0x31, 0x2d, 0xac, 0x23, 0x5f, 0xf4, 0x57, 0x71, 0xd8, 0xa5, 0x3b, 0xa1, 0xde,
	0x30, 0x9b, 0xc0, 0xcd, 0x90, 0x19, 0x3f, 0xde, 0xbf, 0xd6, 0xb3, 0xdf, 0x0c, 0xc6, 0x43, 0x54,
	0xdc, 0x01, 0xd4, 0x6d, 0x19, 0x02, 0x07, 0xe4, 0x46, 0x3e, 0x4e, 0xb3, 0x79, 0xbe, 0xca, 0x0f,
	0xcd, 0xb2, 0x11, 0x0f, 0x70, 0x75, 0xc4, 0x18, 0x76, 0x15, 0xc6, 0xce, 0x91, 0x5c, 0x4c, 0xb3,
	0xf9, 0xc5, 0xea, 0xf2, 0xff, 0xdd, 0x7a, 0x57, 0x6d, 0x3a, 0x47, 0xe2, 0x1e, 0xce, 0x43, 0x17,
	0x22, 0x69, 0x34, 0xa5, 0x26, 0xf9, 0x94, 0x74, 0xb0, 0xaf, 0x3e, 0x4b, 0x4d, 0x62, 0x0c, 0x27,
	0x91, 0x35, 0x79, 0xf9, 0x9c, 0x0c, 0xfb, 0x20, 0x26, 0x90, 0xb7, 0x56, 0x35, 0xd8, 0x27, 0xf9,
	0x92, 0xc8, 0x59, 0x5f, 0x6c, 0x58, 0x93, 0xb8, 0x86, 0x53, 0x4f, 0xc8, 0x86, 0xa3, 0x7c, 0x4d,
	0x68, 0xe4, 0x69, 0x69, 0x38, 0x56, 0xa3, 0xf4, 0x67, 0x8b, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x98, 0x87, 0x1a, 0x82, 0x51, 0x01, 0x00, 0x00,
}