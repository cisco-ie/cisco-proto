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
// source: srv6_locator_info.proto

package cisco_ios_xr_segment_routing_srv6_oper_srv6_standby_locators_locator_info

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

type Srv6LocatorInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Srv6LocatorInfo_KEYS) Reset()         { *m = Srv6LocatorInfo_KEYS{} }
func (m *Srv6LocatorInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*Srv6LocatorInfo_KEYS) ProtoMessage()    {}
func (*Srv6LocatorInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d77284990899884, []int{0}
}

func (m *Srv6LocatorInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6LocatorInfo_KEYS.Unmarshal(m, b)
}
func (m *Srv6LocatorInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6LocatorInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *Srv6LocatorInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6LocatorInfo_KEYS.Merge(m, src)
}
func (m *Srv6LocatorInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_Srv6LocatorInfo_KEYS.Size(m)
}
func (m *Srv6LocatorInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6LocatorInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6LocatorInfo_KEYS proto.InternalMessageInfo

func (m *Srv6LocatorInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Srv6LocatorIntfInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IfHandle             string   `protobuf:"bytes,2,opt,name=if_handle,json=ifHandle,proto3" json:"if_handle,omitempty"`
	ProgrammedPrefix     string   `protobuf:"bytes,3,opt,name=programmed_prefix,json=programmedPrefix,proto3" json:"programmed_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Srv6LocatorIntfInfo) Reset()         { *m = Srv6LocatorIntfInfo{} }
func (m *Srv6LocatorIntfInfo) String() string { return proto.CompactTextString(m) }
func (*Srv6LocatorIntfInfo) ProtoMessage()    {}
func (*Srv6LocatorIntfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d77284990899884, []int{1}
}

func (m *Srv6LocatorIntfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6LocatorIntfInfo.Unmarshal(m, b)
}
func (m *Srv6LocatorIntfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6LocatorIntfInfo.Marshal(b, m, deterministic)
}
func (m *Srv6LocatorIntfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6LocatorIntfInfo.Merge(m, src)
}
func (m *Srv6LocatorIntfInfo) XXX_Size() int {
	return xxx_messageInfo_Srv6LocatorIntfInfo.Size(m)
}
func (m *Srv6LocatorIntfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6LocatorIntfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6LocatorIntfInfo proto.InternalMessageInfo

func (m *Srv6LocatorIntfInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Srv6LocatorIntfInfo) GetIfHandle() string {
	if m != nil {
		return m.IfHandle
	}
	return ""
}

func (m *Srv6LocatorIntfInfo) GetProgrammedPrefix() string {
	if m != nil {
		return m.ProgrammedPrefix
	}
	return ""
}

type MgmtSidmgrTimestamp struct {
	TimeInNanoSeconds    uint64   `protobuf:"varint,1,opt,name=time_in_nano_seconds,json=timeInNanoSeconds,proto3" json:"time_in_nano_seconds,omitempty"`
	AgeInNanoSeconds     uint64   `protobuf:"varint,2,opt,name=age_in_nano_seconds,json=ageInNanoSeconds,proto3" json:"age_in_nano_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MgmtSidmgrTimestamp) Reset()         { *m = MgmtSidmgrTimestamp{} }
func (m *MgmtSidmgrTimestamp) String() string { return proto.CompactTextString(m) }
func (*MgmtSidmgrTimestamp) ProtoMessage()    {}
func (*MgmtSidmgrTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d77284990899884, []int{2}
}

func (m *MgmtSidmgrTimestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MgmtSidmgrTimestamp.Unmarshal(m, b)
}
func (m *MgmtSidmgrTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MgmtSidmgrTimestamp.Marshal(b, m, deterministic)
}
func (m *MgmtSidmgrTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MgmtSidmgrTimestamp.Merge(m, src)
}
func (m *MgmtSidmgrTimestamp) XXX_Size() int {
	return xxx_messageInfo_MgmtSidmgrTimestamp.Size(m)
}
func (m *MgmtSidmgrTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_MgmtSidmgrTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_MgmtSidmgrTimestamp proto.InternalMessageInfo

func (m *MgmtSidmgrTimestamp) GetTimeInNanoSeconds() uint64 {
	if m != nil {
		return m.TimeInNanoSeconds
	}
	return 0
}

func (m *MgmtSidmgrTimestamp) GetAgeInNanoSeconds() uint64 {
	if m != nil {
		return m.AgeInNanoSeconds
	}
	return 0
}

type Srv6LocatorInfo struct {
	Name                 string               `protobuf:"bytes,50,opt,name=name,proto3" json:"name,omitempty"`
	Id                   uint32               `protobuf:"varint,51,opt,name=id,proto3" json:"id,omitempty"`
	Prefix               string               `protobuf:"bytes,52,opt,name=prefix,proto3" json:"prefix,omitempty"`
	IsOperational        bool                 `protobuf:"varint,53,opt,name=is_operational,json=isOperational,proto3" json:"is_operational,omitempty"`
	IsDefault            bool                 `protobuf:"varint,54,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	OutOfResourcesState  string               `protobuf:"bytes,55,opt,name=out_of_resources_state,json=outOfResourcesState,proto3" json:"out_of_resources_state,omitempty"`
	Interface            *Srv6LocatorIntfInfo `protobuf:"bytes,56,opt,name=interface,proto3" json:"interface,omitempty"`
	CreateTimestamp      *MgmtSidmgrTimestamp `protobuf:"bytes,57,opt,name=create_timestamp,json=createTimestamp,proto3" json:"create_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Srv6LocatorInfo) Reset()         { *m = Srv6LocatorInfo{} }
func (m *Srv6LocatorInfo) String() string { return proto.CompactTextString(m) }
func (*Srv6LocatorInfo) ProtoMessage()    {}
func (*Srv6LocatorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d77284990899884, []int{3}
}

func (m *Srv6LocatorInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6LocatorInfo.Unmarshal(m, b)
}
func (m *Srv6LocatorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6LocatorInfo.Marshal(b, m, deterministic)
}
func (m *Srv6LocatorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6LocatorInfo.Merge(m, src)
}
func (m *Srv6LocatorInfo) XXX_Size() int {
	return xxx_messageInfo_Srv6LocatorInfo.Size(m)
}
func (m *Srv6LocatorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6LocatorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6LocatorInfo proto.InternalMessageInfo

func (m *Srv6LocatorInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Srv6LocatorInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Srv6LocatorInfo) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Srv6LocatorInfo) GetIsOperational() bool {
	if m != nil {
		return m.IsOperational
	}
	return false
}

func (m *Srv6LocatorInfo) GetIsDefault() bool {
	if m != nil {
		return m.IsDefault
	}
	return false
}

func (m *Srv6LocatorInfo) GetOutOfResourcesState() string {
	if m != nil {
		return m.OutOfResourcesState
	}
	return ""
}

func (m *Srv6LocatorInfo) GetInterface() *Srv6LocatorIntfInfo {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *Srv6LocatorInfo) GetCreateTimestamp() *MgmtSidmgrTimestamp {
	if m != nil {
		return m.CreateTimestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Srv6LocatorInfo_KEYS)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.standby.locators.locator.info.srv6_locator_info_KEYS")
	proto.RegisterType((*Srv6LocatorIntfInfo)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.standby.locators.locator.info.srv6_locator_intf_info")
	proto.RegisterType((*MgmtSidmgrTimestamp)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.standby.locators.locator.info.mgmt_sidmgr_timestamp")
	proto.RegisterType((*Srv6LocatorInfo)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.standby.locators.locator.info.srv6_locator_info")
}

func init() { proto.RegisterFile("srv6_locator_info.proto", fileDescriptor_7d77284990899884) }

var fileDescriptor_7d77284990899884 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0xc7, 0xb9, 0xb4, 0x94, 0x66, 0xa5, 0x35, 0xd9, 0x6a, 0x5c, 0x10, 0x21, 0x04, 0x84, 0x80,
	0x7a, 0x42, 0xa3, 0x55, 0xdf, 0x15, 0x2c, 0x82, 0x95, 0x8b, 0x2f, 0x3e, 0x8d, 0xdb, 0xbb, 0xdd,
	0x73, 0x20, 0xbb, 0x73, 0xec, 0xce, 0xd5, 0xfa, 0xe4, 0x07, 0xf0, 0x13, 0xfa, 0x6d, 0xe4, 0x36,
	0x89, 0x51, 0x92, 0xc7, 0xbc, 0xcd, 0xcc, 0xff, 0xb7, 0xfb, 0x9f, 0xb9, 0x9d, 0x13, 0x0f, 0x62,
	0xb8, 0xb9, 0x80, 0x05, 0x95, 0x9a, 0x29, 0x00, 0x7a, 0x4b, 0x79, 0x13, 0x88, 0x49, 0x5e, 0x96,
	0x18, 0x4b, 0x02, 0xa4, 0x08, 0xb7, 0x01, 0xa2, 0xa9, 0x9d, 0xf1, 0x0c, 0x81, 0x5a, 0x46, 0x5f,
	0x43, 0x3a, 0x45, 0x8d, 0x09, 0x79, 0x17, 0xe5, 0x91, 0xb5, 0xaf, 0xae, 0x7f, 0xe4, 0xab, 0x7b,
	0xe2, 0x3a, 0xc8, 0xbb, 0x0b, 0x27, 0x4f, 0xc5, 0x68, 0xcb, 0x05, 0x3e, 0xbc, 0xfb, 0x32, 0x97,
	0x52, 0x1c, 0x7a, 0xed, 0x8c, 0xca, 0xc6, 0xd9, 0xb4, 0x5f, 0xa4, 0x78, 0x72, 0xb3, 0x45, 0xb3,
	0x4d, 0x47, 0x76, 0xd1, 0xf2, 0xa1, 0xe8, 0xa3, 0x85, 0x6f, 0xda, 0x57, 0x0b, 0xa3, 0x7a, 0x49,
	0x38, 0x46, 0xfb, 0x3e, 0xe5, 0xf2, 0x89, 0x18, 0x36, 0x81, 0xea, 0xa0, 0x9d, 0x33, 0x15, 0x34,
	0xc1, 0x58, 0xbc, 0x55, 0x07, 0x09, 0x1a, 0x6c, 0x84, 0x4f, 0xa9, 0x3e, 0xf9, 0x2e, 0xee, 0xbb,
	0xda, 0x31, 0x44, 0xac, 0x5c, 0x1d, 0x80, 0xd1, 0x99, 0xc8, 0xda, 0x35, 0xf2, 0xb9, 0xb8, 0xd7,
	0x25, 0x80, 0x1e, 0xbc, 0xf6, 0x04, 0xd1, 0x94, 0xe4, 0xab, 0x98, 0xda, 0x38, 0x2c, 0x86, 0x9d,
	0x76, 0xe9, 0x3f, 0x6a, 0x4f, 0xf3, 0xa5, 0x20, 0x9f, 0x89, 0x33, 0x5d, 0x6f, 0xf3, 0xbd, 0xc4,
	0x0f, 0x74, 0xfd, 0x3f, 0x3e, 0xf9, 0x7d, 0x20, 0x86, 0x5b, 0xdf, 0xe7, 0xef, 0xb0, 0xe7, 0xff,
	0x0c, 0x7b, 0x2a, 0x7a, 0x58, 0xa9, 0xd9, 0x38, 0x9b, 0x9e, 0x14, 0x3d, 0xac, 0xe4, 0x48, 0x1c,
	0xad, 0x86, 0x7a, 0x91, 0xa8, 0x55, 0x26, 0x1f, 0x8b, 0x53, 0x8c, 0xe9, 0x79, 0x34, 0x23, 0x79,
	0xbd, 0x50, 0x2f, 0xc7, 0xd9, 0xf4, 0xb8, 0x38, 0xc1, 0x78, 0xb5, 0x29, 0xca, 0x47, 0x42, 0x60,
	0x84, 0xca, 0x58, 0xdd, 0x2e, 0x58, 0x5d, 0x24, 0xa4, 0x8f, 0xf1, 0xed, 0xb2, 0x20, 0x67, 0x62,
	0x44, 0x2d, 0x03, 0x59, 0x08, 0x26, 0x52, 0x1b, 0x4a, 0x13, 0x21, 0xb2, 0x66, 0xa3, 0x5e, 0x25,
	0xb7, 0x33, 0x6a, 0xf9, 0xca, 0x16, 0x6b, 0x6d, 0xde, 0x49, 0xf2, 0xa7, 0xe8, 0xa3, 0x67, 0x13,
	0xac, 0x2e, 0x8d, 0x7a, 0x3d, 0xce, 0xa6, 0x77, 0xce, 0x75, 0xbe, 0xb7, 0x55, 0xca, 0x77, 0x6f,
	0x46, 0xb1, 0xf1, 0x94, 0xbf, 0x32, 0x31, 0x28, 0x83, 0xd1, 0x6c, 0x36, 0x4f, 0xa8, 0xde, 0xa4,
	0x46, 0xbe, 0xee, 0xb1, 0x91, 0x9d, 0xab, 0x52, 0xdc, 0x5d, 0x3a, 0x7f, 0x5e, 0x17, 0xae, 0x8f,
	0xd2, 0xcf, 0x34, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x51, 0x47, 0xa1, 0xf2, 0x67, 0x03, 0x00,
	0x00,
}
