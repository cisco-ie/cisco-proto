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
// source: diameter_gy_sess.proto

package cisco_ios_xr_aaa_diameter_oper_augment_diameter_gy_session_ids_gy_session_id

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

type DiameterGySess_KEYS struct {
	SessionId            uint32   `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiameterGySess_KEYS) Reset()         { *m = DiameterGySess_KEYS{} }
func (m *DiameterGySess_KEYS) String() string { return proto.CompactTextString(m) }
func (*DiameterGySess_KEYS) ProtoMessage()    {}
func (*DiameterGySess_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_63acc6778bb6f1b9, []int{0}
}

func (m *DiameterGySess_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiameterGySess_KEYS.Unmarshal(m, b)
}
func (m *DiameterGySess_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiameterGySess_KEYS.Marshal(b, m, deterministic)
}
func (m *DiameterGySess_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiameterGySess_KEYS.Merge(m, src)
}
func (m *DiameterGySess_KEYS) XXX_Size() int {
	return xxx_messageInfo_DiameterGySess_KEYS.Size(m)
}
func (m *DiameterGySess_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_DiameterGySess_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_DiameterGySess_KEYS proto.InternalMessageInfo

func (m *DiameterGySess_KEYS) GetSessionId() uint32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type DiameterGySess struct {
	AaaSessionId         uint32   `protobuf:"varint,50,opt,name=aaa_session_id,json=aaaSessionId,proto3" json:"aaa_session_id,omitempty"`
	ParentAaaSessionId   uint32   `protobuf:"varint,51,opt,name=parent_aaa_session_id,json=parentAaaSessionId,proto3" json:"parent_aaa_session_id,omitempty"`
	DiameterSessionId    string   `protobuf:"bytes,52,opt,name=diameter_session_id,json=diameterSessionId,proto3" json:"diameter_session_id,omitempty"`
	RequestNumber        uint32   `protobuf:"varint,53,opt,name=request_number,json=requestNumber,proto3" json:"request_number,omitempty"`
	SessionState         string   `protobuf:"bytes,54,opt,name=session_state,json=sessionState,proto3" json:"session_state,omitempty"`
	RequestType          string   `protobuf:"bytes,55,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	RetryCount           uint32   `protobuf:"varint,56,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiameterGySess) Reset()         { *m = DiameterGySess{} }
func (m *DiameterGySess) String() string { return proto.CompactTextString(m) }
func (*DiameterGySess) ProtoMessage()    {}
func (*DiameterGySess) Descriptor() ([]byte, []int) {
	return fileDescriptor_63acc6778bb6f1b9, []int{1}
}

func (m *DiameterGySess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiameterGySess.Unmarshal(m, b)
}
func (m *DiameterGySess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiameterGySess.Marshal(b, m, deterministic)
}
func (m *DiameterGySess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiameterGySess.Merge(m, src)
}
func (m *DiameterGySess) XXX_Size() int {
	return xxx_messageInfo_DiameterGySess.Size(m)
}
func (m *DiameterGySess) XXX_DiscardUnknown() {
	xxx_messageInfo_DiameterGySess.DiscardUnknown(m)
}

var xxx_messageInfo_DiameterGySess proto.InternalMessageInfo

func (m *DiameterGySess) GetAaaSessionId() uint32 {
	if m != nil {
		return m.AaaSessionId
	}
	return 0
}

func (m *DiameterGySess) GetParentAaaSessionId() uint32 {
	if m != nil {
		return m.ParentAaaSessionId
	}
	return 0
}

func (m *DiameterGySess) GetDiameterSessionId() string {
	if m != nil {
		return m.DiameterSessionId
	}
	return ""
}

func (m *DiameterGySess) GetRequestNumber() uint32 {
	if m != nil {
		return m.RequestNumber
	}
	return 0
}

func (m *DiameterGySess) GetSessionState() string {
	if m != nil {
		return m.SessionState
	}
	return ""
}

func (m *DiameterGySess) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *DiameterGySess) GetRetryCount() uint32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func init() {
	proto.RegisterType((*DiameterGySess_KEYS)(nil), "cisco_ios_xr_aaa_diameter_oper.augment.diameter.gy_session_ids.gy_session_id.diameter_gy_sess_KEYS")
	proto.RegisterType((*DiameterGySess)(nil), "cisco_ios_xr_aaa_diameter_oper.augment.diameter.gy_session_ids.gy_session_id.diameter_gy_sess")
}

func init() { proto.RegisterFile("diameter_gy_sess.proto", fileDescriptor_63acc6778bb6f1b9) }

var fileDescriptor_63acc6778bb6f1b9 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xc9, 0x77, 0xf8, 0xa0, 0xd3, 0xa4, 0xe8, 0x4a, 0x25, 0x17, 0xb1, 0x56, 0x85, 0x9e,
	0x02, 0x5a, 0xad, 0x5e, 0x45, 0x3c, 0x88, 0xe2, 0xa1, 0xf1, 0xe2, 0x69, 0x98, 0x26, 0x43, 0xc9,
	0x21, 0xd9, 0xb8, 0x3b, 0x01, 0xf3, 0x9b, 0xfc, 0x93, 0x92, 0x35, 0x69, 0x63, 0x8e, 0x79, 0xe6,
	0x7d, 0xde, 0x37, 0xb0, 0x70, 0x9c, 0x66, 0x94, 0xb3, 0xb0, 0xc1, 0x6d, 0x8d, 0x96, 0xad, 0x8d,
	0x4a, 0xa3, 0x45, 0xab, 0xd7, 0x24, 0xb3, 0x89, 0xc6, 0x4c, 0x5b, 0xfc, 0x32, 0x48, 0x44, 0xb8,
	0x0b, 0xea, 0x92, 0x4d, 0x44, 0xd5, 0x36, 0xe7, 0x42, 0xa2, 0x8e, 0x46, 0xad, 0x9e, 0xe9, 0x02,
	0xb3, 0xd4, 0xfe, 0xfd, 0x9c, 0xaf, 0x60, 0x3a, 0xdc, 0xc1, 0x97, 0xa7, 0x8f, 0x58, 0x9d, 0x00,
	0xec, 0x63, 0xa1, 0x37, 0xf3, 0x16, 0xc1, 0x7a, 0xd4, 0x92, 0xe7, 0x74, 0xfe, 0xfd, 0x0f, 0x0e,
	0x86, 0xa2, 0xba, 0x80, 0x49, 0xf3, 0x3f, 0x3d, 0xef, 0xda, 0x79, 0x3e, 0x11, 0xc5, 0x9d, 0xaa,
	0xae, 0x60, 0x5a, 0x92, 0xe1, 0x42, 0x70, 0x10, 0x5e, 0xba, 0xb0, 0xfa, 0x3d, 0x3e, 0xf4, 0x95,
	0x08, 0x8e, 0x76, 0x63, 0x3d, 0xe1, 0x66, 0xe6, 0x2d, 0x46, 0xeb, 0xc3, 0xee, 0xb4, 0xcf, 0x5f,
	0xc2, 0xc4, 0xf0, 0x67, 0xc5, 0x56, 0xb0, 0xa8, 0xf2, 0x0d, 0x9b, 0xf0, 0xd6, 0x75, 0x07, 0x2d,
	0x7d, 0x73, 0x50, 0x9d, 0x43, 0xd0, 0xb5, 0x59, 0x21, 0xe1, 0x70, 0xe5, 0x0a, 0xfd, 0x16, 0xc6,
	0x0d, 0x53, 0x67, 0xe0, 0x77, 0x5d, 0x52, 0x97, 0x1c, 0xde, 0xb9, 0xcc, 0xb8, 0x65, 0xef, 0x75,
	0xc9, 0xea, 0x14, 0xc6, 0x86, 0xc5, 0xd4, 0x98, 0xe8, 0xaa, 0x90, 0xf0, 0xde, 0x6d, 0x81, 0x43,
	0x8f, 0x0d, 0xd9, 0xfc, 0x77, 0x4f, 0xb7, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x88, 0xbc,
	0xf4, 0xd4, 0x01, 0x00, 0x00,
}