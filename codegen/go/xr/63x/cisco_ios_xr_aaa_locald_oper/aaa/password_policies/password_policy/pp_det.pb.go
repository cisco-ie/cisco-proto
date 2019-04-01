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
// source: pp_det.proto

package cisco_ios_xr_aaa_locald_oper_aaa_password_policies_password_policy

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

type PpDet_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PpDet_KEYS) Reset()         { *m = PpDet_KEYS{} }
func (m *PpDet_KEYS) String() string { return proto.CompactTextString(m) }
func (*PpDet_KEYS) ProtoMessage()    {}
func (*PpDet_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fbca33dc58e5d21, []int{0}
}

func (m *PpDet_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PpDet_KEYS.Unmarshal(m, b)
}
func (m *PpDet_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PpDet_KEYS.Marshal(b, m, deterministic)
}
func (m *PpDet_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PpDet_KEYS.Merge(m, src)
}
func (m *PpDet_KEYS) XXX_Size() int {
	return xxx_messageInfo_PpDet_KEYS.Size(m)
}
func (m *PpDet_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PpDet_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PpDet_KEYS proto.InternalMessageInfo

func (m *PpDet_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PpDuration struct {
	Years                uint32   `protobuf:"varint,1,opt,name=years,proto3" json:"years,omitempty"`
	Months               uint32   `protobuf:"varint,2,opt,name=months,proto3" json:"months,omitempty"`
	Days                 uint32   `protobuf:"varint,3,opt,name=days,proto3" json:"days,omitempty"`
	Hours                uint32   `protobuf:"varint,4,opt,name=hours,proto3" json:"hours,omitempty"`
	Mins                 uint32   `protobuf:"varint,5,opt,name=mins,proto3" json:"mins,omitempty"`
	Secs                 uint32   `protobuf:"varint,6,opt,name=secs,proto3" json:"secs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PpDuration) Reset()         { *m = PpDuration{} }
func (m *PpDuration) String() string { return proto.CompactTextString(m) }
func (*PpDuration) ProtoMessage()    {}
func (*PpDuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fbca33dc58e5d21, []int{1}
}

func (m *PpDuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PpDuration.Unmarshal(m, b)
}
func (m *PpDuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PpDuration.Marshal(b, m, deterministic)
}
func (m *PpDuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PpDuration.Merge(m, src)
}
func (m *PpDuration) XXX_Size() int {
	return xxx_messageInfo_PpDuration.Size(m)
}
func (m *PpDuration) XXX_DiscardUnknown() {
	xxx_messageInfo_PpDuration.DiscardUnknown(m)
}

var xxx_messageInfo_PpDuration proto.InternalMessageInfo

func (m *PpDuration) GetYears() uint32 {
	if m != nil {
		return m.Years
	}
	return 0
}

func (m *PpDuration) GetMonths() uint32 {
	if m != nil {
		return m.Months
	}
	return 0
}

func (m *PpDuration) GetDays() uint32 {
	if m != nil {
		return m.Days
	}
	return 0
}

func (m *PpDuration) GetHours() uint32 {
	if m != nil {
		return m.Hours
	}
	return 0
}

func (m *PpDuration) GetMins() uint32 {
	if m != nil {
		return m.Mins
	}
	return 0
}

func (m *PpDuration) GetSecs() uint32 {
	if m != nil {
		return m.Secs
	}
	return 0
}

type PpDet struct {
	NameXr               string      `protobuf:"bytes,50,opt,name=name_xr,json=nameXr,proto3" json:"name_xr,omitempty"`
	MinLen               uint32      `protobuf:"varint,51,opt,name=min_len,json=minLen,proto3" json:"min_len,omitempty"`
	MaxLen               uint32      `protobuf:"varint,52,opt,name=max_len,json=maxLen,proto3" json:"max_len,omitempty"`
	SplChar              uint32      `protobuf:"varint,53,opt,name=spl_char,json=splChar,proto3" json:"spl_char,omitempty"`
	UpperCase            uint32      `protobuf:"varint,54,opt,name=upper_case,json=upperCase,proto3" json:"upper_case,omitempty"`
	LowerCase            uint32      `protobuf:"varint,55,opt,name=lower_case,json=lowerCase,proto3" json:"lower_case,omitempty"`
	Numeric              uint32      `protobuf:"varint,56,opt,name=numeric,proto3" json:"numeric,omitempty"`
	LifeTime             *PpDuration `protobuf:"bytes,57,opt,name=life_time,json=lifeTime,proto3" json:"life_time,omitempty"`
	LockOutTime          *PpDuration `protobuf:"bytes,58,opt,name=lock_out_time,json=lockOutTime,proto3" json:"lock_out_time,omitempty"`
	MinCharChange        uint32      `protobuf:"varint,59,opt,name=min_char_change,json=minCharChange,proto3" json:"min_char_change,omitempty"`
	NumOfUsers           uint32      `protobuf:"varint,60,opt,name=num_of_users,json=numOfUsers,proto3" json:"num_of_users,omitempty"`
	MaxFailAttempts      uint32      `protobuf:"varint,61,opt,name=max_fail_attempts,json=maxFailAttempts,proto3" json:"max_fail_attempts,omitempty"`
	UsrCount             uint32      `protobuf:"varint,62,opt,name=usr_count,json=usrCount,proto3" json:"usr_count,omitempty"`
	ErrCount             uint32      `protobuf:"varint,63,opt,name=err_count,json=errCount,proto3" json:"err_count,omitempty"`
	LockOutCount         uint32      `protobuf:"varint,64,opt,name=lock_out_count,json=lockOutCount,proto3" json:"lock_out_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PpDet) Reset()         { *m = PpDet{} }
func (m *PpDet) String() string { return proto.CompactTextString(m) }
func (*PpDet) ProtoMessage()    {}
func (*PpDet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fbca33dc58e5d21, []int{2}
}

func (m *PpDet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PpDet.Unmarshal(m, b)
}
func (m *PpDet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PpDet.Marshal(b, m, deterministic)
}
func (m *PpDet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PpDet.Merge(m, src)
}
func (m *PpDet) XXX_Size() int {
	return xxx_messageInfo_PpDet.Size(m)
}
func (m *PpDet) XXX_DiscardUnknown() {
	xxx_messageInfo_PpDet.DiscardUnknown(m)
}

var xxx_messageInfo_PpDet proto.InternalMessageInfo

func (m *PpDet) GetNameXr() string {
	if m != nil {
		return m.NameXr
	}
	return ""
}

func (m *PpDet) GetMinLen() uint32 {
	if m != nil {
		return m.MinLen
	}
	return 0
}

func (m *PpDet) GetMaxLen() uint32 {
	if m != nil {
		return m.MaxLen
	}
	return 0
}

func (m *PpDet) GetSplChar() uint32 {
	if m != nil {
		return m.SplChar
	}
	return 0
}

func (m *PpDet) GetUpperCase() uint32 {
	if m != nil {
		return m.UpperCase
	}
	return 0
}

func (m *PpDet) GetLowerCase() uint32 {
	if m != nil {
		return m.LowerCase
	}
	return 0
}

func (m *PpDet) GetNumeric() uint32 {
	if m != nil {
		return m.Numeric
	}
	return 0
}

func (m *PpDet) GetLifeTime() *PpDuration {
	if m != nil {
		return m.LifeTime
	}
	return nil
}

func (m *PpDet) GetLockOutTime() *PpDuration {
	if m != nil {
		return m.LockOutTime
	}
	return nil
}

func (m *PpDet) GetMinCharChange() uint32 {
	if m != nil {
		return m.MinCharChange
	}
	return 0
}

func (m *PpDet) GetNumOfUsers() uint32 {
	if m != nil {
		return m.NumOfUsers
	}
	return 0
}

func (m *PpDet) GetMaxFailAttempts() uint32 {
	if m != nil {
		return m.MaxFailAttempts
	}
	return 0
}

func (m *PpDet) GetUsrCount() uint32 {
	if m != nil {
		return m.UsrCount
	}
	return 0
}

func (m *PpDet) GetErrCount() uint32 {
	if m != nil {
		return m.ErrCount
	}
	return 0
}

func (m *PpDet) GetLockOutCount() uint32 {
	if m != nil {
		return m.LockOutCount
	}
	return 0
}

func init() {
	proto.RegisterType((*PpDet_KEYS)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.password_policies.password_policy.pp_det_KEYS")
	proto.RegisterType((*PpDuration)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.password_policies.password_policy.pp_duration")
	proto.RegisterType((*PpDet)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.password_policies.password_policy.pp_det")
}

func init() { proto.RegisterFile("pp_det.proto", fileDescriptor_8fbca33dc58e5d21) }

var fileDescriptor_8fbca33dc58e5d21 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x5b, 0x8b, 0x53, 0x31,
	0x10, 0xc7, 0xa9, 0x76, 0x7b, 0x49, 0x5b, 0x17, 0x83, 0x68, 0x44, 0x84, 0x5a, 0x44, 0x16, 0x1f,
	0xfa, 0xb0, 0xeb, 0xfd, 0xae, 0x45, 0x5f, 0x14, 0x0a, 0x55, 0x41, 0x9f, 0xc2, 0x78, 0x9a, 0xda,
	0x60, 0x6e, 0x64, 0x12, 0xb6, 0xfd, 0x10, 0x7e, 0x66, 0x25, 0x93, 0xb3, 0x22, 0xbe, 0xca, 0xbe,
	0x1c, 0x66, 0xfe, 0xbf, 0x3f, 0x33, 0x73, 0x66, 0x08, 0x1b, 0x87, 0x20, 0xd7, 0x2a, 0xcd, 0x43,
	0xf4, 0xc9, 0xf3, 0x37, 0x8d, 0xc6, 0xc6, 0x4b, 0xed, 0x51, 0xee, 0xa2, 0x04, 0x00, 0x69, 0x7c,
	0x03, 0x66, 0x2d, 0x7d, 0x50, 0x71, 0x0e, 0x00, 0xf3, 0x00, 0x88, 0xa7, 0x3e, 0xae, 0x65, 0xf0,
	0x46, 0x37, 0x5a, 0xe1, 0x3f, 0xca, 0x7e, 0x76, 0x8b, 0x8d, 0x6a, 0x4d, 0xf9, 0xfe, 0xed, 0xd7,
	0x8f, 0x9c, 0xb3, 0xae, 0x03, 0xab, 0x44, 0x67, 0xda, 0x39, 0x1a, 0xae, 0x28, 0x9e, 0xfd, 0xec,
	0x54, 0x4f, 0x8e, 0x90, 0xb4, 0x77, 0xfc, 0x0a, 0x3b, 0xd8, 0x2b, 0x88, 0x48, 0xa6, 0xc9, 0xaa,
	0x26, 0xfc, 0x2a, 0xeb, 0x59, 0xef, 0xd2, 0x16, 0xc5, 0x05, 0x92, 0xdb, 0xac, 0x54, 0x5c, 0xc3,
	0x1e, 0xc5, 0x45, 0x52, 0x29, 0x2e, 0x15, 0xb6, 0x3e, 0x47, 0x14, 0xdd, 0x5a, 0x81, 0x92, 0xe2,
	0xb4, 0xda, 0xa1, 0x38, 0xa8, 0xce, 0x12, 0x17, 0x0d, 0x55, 0x83, 0xa2, 0x57, 0xb5, 0x12, 0xcf,
	0x7e, 0x75, 0x59, 0xaf, 0xce, 0xcc, 0xaf, 0xb1, 0x7e, 0x19, 0x51, 0xee, 0xa2, 0x38, 0xa6, 0x89,
	0x7b, 0x25, 0xfd, 0x12, 0x0b, 0xb0, 0xda, 0x49, 0xa3, 0x9c, 0x38, 0x69, 0xc7, 0xd1, 0xee, 0x83,
	0x72, 0x04, 0x60, 0x47, 0xe0, 0x5e, 0x0b, 0x60, 0x57, 0xc0, 0x75, 0x36, 0xc0, 0x60, 0x64, 0xb3,
	0x85, 0x28, 0xee, 0x13, 0xe9, 0x63, 0x30, 0x8b, 0x2d, 0x44, 0x7e, 0x93, 0xb1, 0x1c, 0x82, 0x8a,
	0xb2, 0x01, 0x54, 0xe2, 0x01, 0xc1, 0x21, 0x29, 0x0b, 0x40, 0x55, 0xb0, 0xf1, 0xa7, 0x67, 0xf8,
	0x61, 0xc5, 0xa4, 0x10, 0x16, 0xac, 0xef, 0xb2, 0x55, 0x51, 0x37, 0xe2, 0x51, 0xad, 0xdb, 0xa6,
	0xdc, 0xb0, 0xa1, 0xd1, 0x1b, 0x25, 0x93, 0xb6, 0x4a, 0x3c, 0x9e, 0x76, 0x8e, 0x46, 0xc7, 0xcb,
	0xf9, 0xff, 0xdf, 0x74, 0xfe, 0xd7, 0xb1, 0x56, 0x83, 0xd2, 0xe1, 0x93, 0xb6, 0x8a, 0x23, 0x9b,
	0x18, 0xdf, 0xfc, 0x90, 0x3e, 0xa7, 0xda, 0xf1, 0xc9, 0xf9, 0x74, 0x1c, 0x95, 0x2e, 0xcb, 0x9c,
	0xa8, 0xe9, 0x1d, 0x76, 0x58, 0xee, 0x50, 0xb6, 0x5a, 0x3e, 0xee, 0xbb, 0x12, 0x4f, 0x69, 0x09,
	0x13, 0xab, 0x5d, 0x59, 0xee, 0x82, 0x44, 0x3e, 0x65, 0x63, 0x97, 0xad, 0xf4, 0x1b, 0x99, 0x51,
	0x45, 0x14, 0xcf, 0xc8, 0xc4, 0x5c, 0xb6, 0xcb, 0xcd, 0xe7, 0xa2, 0xf0, 0xbb, 0xec, 0x72, 0x39,
	0xdc, 0x06, 0xb4, 0x91, 0x90, 0x92, 0xb2, 0x21, 0xa1, 0x78, 0x4e, 0xb6, 0x43, 0x0b, 0xbb, 0x77,
	0xa0, 0xcd, 0xeb, 0x56, 0xe6, 0x37, 0xd8, 0x30, 0x63, 0x94, 0x8d, 0xcf, 0x2e, 0x89, 0x17, 0xe4,
	0x19, 0x64, 0x8c, 0x8b, 0x92, 0x17, 0xa8, 0xe2, 0x19, 0x7c, 0x59, 0xa1, 0x8a, 0x2d, 0xbc, 0xcd,
	0x2e, 0xfd, 0x59, 0x52, 0x75, 0xbc, 0x22, 0xc7, 0xb8, 0xfd, 0x29, 0x72, 0x7d, 0xeb, 0xd1, 0xfb,
	0x3b, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x30, 0xfa, 0xd8, 0x21, 0x8f, 0x03, 0x00, 0x00,
}