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
// source: user_det.proto

package cisco_ios_xr_aaa_locald_oper_aaa_users_user

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

type UserDet_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDet_KEYS) Reset()         { *m = UserDet_KEYS{} }
func (m *UserDet_KEYS) String() string { return proto.CompactTextString(m) }
func (*UserDet_KEYS) ProtoMessage()    {}
func (*UserDet_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8834d045a739849a, []int{0}
}

func (m *UserDet_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDet_KEYS.Unmarshal(m, b)
}
func (m *UserDet_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDet_KEYS.Marshal(b, m, deterministic)
}
func (m *UserDet_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDet_KEYS.Merge(m, src)
}
func (m *UserDet_KEYS) XXX_Size() int {
	return xxx_messageInfo_UserDet_KEYS.Size(m)
}
func (m *UserDet_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDet_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_UserDet_KEYS proto.InternalMessageInfo

func (m *UserDet_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Taskid struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Read                 bool     `protobuf:"varint,2,opt,name=read,proto3" json:"read,omitempty"`
	Write                bool     `protobuf:"varint,3,opt,name=write,proto3" json:"write,omitempty"`
	Execute              bool     `protobuf:"varint,4,opt,name=execute,proto3" json:"execute,omitempty"`
	Debug                bool     `protobuf:"varint,5,opt,name=debug,proto3" json:"debug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Taskid) Reset()         { *m = Taskid{} }
func (m *Taskid) String() string { return proto.CompactTextString(m) }
func (*Taskid) ProtoMessage()    {}
func (*Taskid) Descriptor() ([]byte, []int) {
	return fileDescriptor_8834d045a739849a, []int{1}
}

func (m *Taskid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Taskid.Unmarshal(m, b)
}
func (m *Taskid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Taskid.Marshal(b, m, deterministic)
}
func (m *Taskid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Taskid.Merge(m, src)
}
func (m *Taskid) XXX_Size() int {
	return xxx_messageInfo_Taskid.Size(m)
}
func (m *Taskid) XXX_DiscardUnknown() {
	xxx_messageInfo_Taskid.DiscardUnknown(m)
}

var xxx_messageInfo_Taskid proto.InternalMessageInfo

func (m *Taskid) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *Taskid) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func (m *Taskid) GetWrite() bool {
	if m != nil {
		return m.Write
	}
	return false
}

func (m *Taskid) GetExecute() bool {
	if m != nil {
		return m.Execute
	}
	return false
}

func (m *Taskid) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

type Taskmap struct {
	Tasks                []*Taskid `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Taskmap) Reset()         { *m = Taskmap{} }
func (m *Taskmap) String() string { return proto.CompactTextString(m) }
func (*Taskmap) ProtoMessage()    {}
func (*Taskmap) Descriptor() ([]byte, []int) {
	return fileDescriptor_8834d045a739849a, []int{2}
}

func (m *Taskmap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Taskmap.Unmarshal(m, b)
}
func (m *Taskmap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Taskmap.Marshal(b, m, deterministic)
}
func (m *Taskmap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Taskmap.Merge(m, src)
}
func (m *Taskmap) XXX_Size() int {
	return xxx_messageInfo_Taskmap.Size(m)
}
func (m *Taskmap) XXX_DiscardUnknown() {
	xxx_messageInfo_Taskmap.DiscardUnknown(m)
}

var xxx_messageInfo_Taskmap proto.InternalMessageInfo

func (m *Taskmap) GetTasks() []*Taskid {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type UserDet struct {
	Usergroup            []string `protobuf:"bytes,50,rep,name=usergroup,proto3" json:"usergroup,omitempty"`
	TaskMap              *Taskmap `protobuf:"bytes,51,opt,name=task_map,json=taskMap,proto3" json:"task_map,omitempty"`
	AdminUser            bool     `protobuf:"varint,52,opt,name=admin_user,json=adminUser,proto3" json:"admin_user,omitempty"`
	FirstUser            bool     `protobuf:"varint,53,opt,name=first_user,json=firstUser,proto3" json:"first_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDet) Reset()         { *m = UserDet{} }
func (m *UserDet) String() string { return proto.CompactTextString(m) }
func (*UserDet) ProtoMessage()    {}
func (*UserDet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8834d045a739849a, []int{3}
}

func (m *UserDet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDet.Unmarshal(m, b)
}
func (m *UserDet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDet.Marshal(b, m, deterministic)
}
func (m *UserDet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDet.Merge(m, src)
}
func (m *UserDet) XXX_Size() int {
	return xxx_messageInfo_UserDet.Size(m)
}
func (m *UserDet) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDet.DiscardUnknown(m)
}

var xxx_messageInfo_UserDet proto.InternalMessageInfo

func (m *UserDet) GetUsergroup() []string {
	if m != nil {
		return m.Usergroup
	}
	return nil
}

func (m *UserDet) GetTaskMap() *Taskmap {
	if m != nil {
		return m.TaskMap
	}
	return nil
}

func (m *UserDet) GetAdminUser() bool {
	if m != nil {
		return m.AdminUser
	}
	return false
}

func (m *UserDet) GetFirstUser() bool {
	if m != nil {
		return m.FirstUser
	}
	return false
}

func init() {
	proto.RegisterType((*UserDet_KEYS)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.users.user.user_det_KEYS")
	proto.RegisterType((*Taskid)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.users.user.taskid")
	proto.RegisterType((*Taskmap)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.users.user.taskmap")
	proto.RegisterType((*UserDet)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.users.user.user_det")
}

func init() { proto.RegisterFile("user_det.proto", fileDescriptor_8834d045a739849a) }

var fileDescriptor_8834d045a739849a = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xa9, 0xfb, 0xb7, 0xb3, 0xe8, 0x21, 0x08, 0xe6, 0xa0, 0x50, 0xea, 0xa5, 0x20, 0xf4,
	0xb0, 0xbb, 0x7e, 0x04, 0x0f, 0x8b, 0x88, 0x50, 0xf5, 0xe0, 0x29, 0xcc, 0x36, 0x71, 0x09, 0x6e,
	0x37, 0x21, 0x49, 0x71, 0xc1, 0x0f, 0xe6, 0xd7, 0x93, 0x4c, 0x2c, 0x9e, 0xf7, 0x52, 0xde, 0xbc,
	0x37, 0xaf, 0xc3, 0x8f, 0xc0, 0x45, 0xef, 0x95, 0x13, 0x52, 0x85, 0xda, 0x3a, 0x13, 0x0c, 0xbb,
	0x6b, 0xb5, 0x6f, 0x8d, 0xd0, 0xc6, 0x8b, 0xa3, 0x13, 0x88, 0x28, 0xf6, 0xa6, 0xc5, 0xbd, 0x14,
	0xc6, 0x2a, 0x57, 0x23, 0x62, 0x1d, 0x0b, 0x9e, 0xbe, 0xe5, 0x2d, 0x9c, 0x0f, 0x75, 0xf1, 0xf8,
	0xf0, 0xfe, 0xc2, 0x18, 0x8c, 0x0f, 0xd8, 0x29, 0x9e, 0x15, 0x59, 0x95, 0x37, 0xa4, 0xcb, 0x6f,
	0x98, 0x06, 0xf4, 0x9f, 0x5a, 0xb2, 0x2b, 0x98, 0x45, 0x25, 0xb4, 0xfc, 0x5b, 0xa0, 0x60, 0x23,
	0x63, 0xcd, 0x29, 0x94, 0xfc, 0xac, 0xc8, 0xaa, 0x79, 0x43, 0x9a, 0x5d, 0xc2, 0xe4, 0xcb, 0xe9,
	0xa0, 0xf8, 0x88, 0xcc, 0x34, 0x30, 0x0e, 0x33, 0x75, 0x54, 0x6d, 0x1f, 0x14, 0x1f, 0x93, 0x3f,
	0x8c, 0x71, 0x5f, 0xaa, 0x6d, 0xbf, 0xe3, 0x93, 0xb4, 0x4f, 0x43, 0xf9, 0x9a, 0x4e, 0x76, 0x68,
	0xd9, 0x06, 0x26, 0x51, 0x7a, 0x9e, 0x15, 0xa3, 0x6a, 0xb1, 0x5c, 0xd5, 0x27, 0x90, 0xd6, 0x89,
	0xa0, 0x49, 0x7f, 0x28, 0x7f, 0x32, 0x98, 0x0f, 0xe0, 0xec, 0x1a, 0xf2, 0xa8, 0x77, 0xce, 0xf4,
	0x96, 0x2f, 0x8b, 0x51, 0x95, 0x37, 0xff, 0x06, 0x7b, 0x86, 0x39, 0x31, 0x77, 0x68, 0xf9, 0xaa,
	0xc8, 0xaa, 0xc5, 0x72, 0x7d, 0xf2, 0xe1, 0x0e, 0x6d, 0x43, 0x18, 0x4f, 0x68, 0xd9, 0x0d, 0x00,
	0xca, 0x4e, 0x1f, 0x44, 0x8c, 0xf9, 0x9a, 0x60, 0x73, 0x72, 0xde, 0xbc, 0x72, 0x31, 0xfe, 0xd0,
	0xce, 0x87, 0x14, 0xdf, 0xa7, 0x98, 0x9c, 0x18, 0x6f, 0xa7, 0xf4, 0xca, 0xab, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xec, 0xce, 0x39, 0x8c, 0xf7, 0x01, 0x00, 0x00,
}
