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
// source: usergroup_det.proto

package cisco_ios_xr_aaa_locald_oper_aaa_usergroups_usergroup

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

type UsergroupDet_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsergroupDet_KEYS) Reset()         { *m = UsergroupDet_KEYS{} }
func (m *UsergroupDet_KEYS) String() string { return proto.CompactTextString(m) }
func (*UsergroupDet_KEYS) ProtoMessage()    {}
func (*UsergroupDet_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef64b70d856c9bf, []int{0}
}

func (m *UsergroupDet_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsergroupDet_KEYS.Unmarshal(m, b)
}
func (m *UsergroupDet_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsergroupDet_KEYS.Marshal(b, m, deterministic)
}
func (m *UsergroupDet_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsergroupDet_KEYS.Merge(m, src)
}
func (m *UsergroupDet_KEYS) XXX_Size() int {
	return xxx_messageInfo_UsergroupDet_KEYS.Size(m)
}
func (m *UsergroupDet_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_UsergroupDet_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_UsergroupDet_KEYS proto.InternalMessageInfo

func (m *UsergroupDet_KEYS) GetName() string {
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
	return fileDescriptor_3ef64b70d856c9bf, []int{1}
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
	return fileDescriptor_3ef64b70d856c9bf, []int{2}
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

type TaskgroupDet struct {
	NameXr               string   `protobuf:"bytes,1,opt,name=name_xr,json=nameXr,proto3" json:"name_xr,omitempty"`
	IncludedTaskIds      *Taskmap `protobuf:"bytes,2,opt,name=included_task_ids,json=includedTaskIds,proto3" json:"included_task_ids,omitempty"`
	TaskMap              *Taskmap `protobuf:"bytes,3,opt,name=task_map,json=taskMap,proto3" json:"task_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskgroupDet) Reset()         { *m = TaskgroupDet{} }
func (m *TaskgroupDet) String() string { return proto.CompactTextString(m) }
func (*TaskgroupDet) ProtoMessage()    {}
func (*TaskgroupDet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef64b70d856c9bf, []int{3}
}

func (m *TaskgroupDet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskgroupDet.Unmarshal(m, b)
}
func (m *TaskgroupDet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskgroupDet.Marshal(b, m, deterministic)
}
func (m *TaskgroupDet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskgroupDet.Merge(m, src)
}
func (m *TaskgroupDet) XXX_Size() int {
	return xxx_messageInfo_TaskgroupDet.Size(m)
}
func (m *TaskgroupDet) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskgroupDet.DiscardUnknown(m)
}

var xxx_messageInfo_TaskgroupDet proto.InternalMessageInfo

func (m *TaskgroupDet) GetNameXr() string {
	if m != nil {
		return m.NameXr
	}
	return ""
}

func (m *TaskgroupDet) GetIncludedTaskIds() *Taskmap {
	if m != nil {
		return m.IncludedTaskIds
	}
	return nil
}

func (m *TaskgroupDet) GetTaskMap() *Taskmap {
	if m != nil {
		return m.TaskMap
	}
	return nil
}

type UsergroupDet struct {
	NameXr               string          `protobuf:"bytes,50,opt,name=name_xr,json=nameXr,proto3" json:"name_xr,omitempty"`
	Taskgroup            []*TaskgroupDet `protobuf:"bytes,51,rep,name=taskgroup,proto3" json:"taskgroup,omitempty"`
	TaskMap              *Taskmap        `protobuf:"bytes,52,opt,name=task_map,json=taskMap,proto3" json:"task_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UsergroupDet) Reset()         { *m = UsergroupDet{} }
func (m *UsergroupDet) String() string { return proto.CompactTextString(m) }
func (*UsergroupDet) ProtoMessage()    {}
func (*UsergroupDet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef64b70d856c9bf, []int{4}
}

func (m *UsergroupDet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsergroupDet.Unmarshal(m, b)
}
func (m *UsergroupDet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsergroupDet.Marshal(b, m, deterministic)
}
func (m *UsergroupDet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsergroupDet.Merge(m, src)
}
func (m *UsergroupDet) XXX_Size() int {
	return xxx_messageInfo_UsergroupDet.Size(m)
}
func (m *UsergroupDet) XXX_DiscardUnknown() {
	xxx_messageInfo_UsergroupDet.DiscardUnknown(m)
}

var xxx_messageInfo_UsergroupDet proto.InternalMessageInfo

func (m *UsergroupDet) GetNameXr() string {
	if m != nil {
		return m.NameXr
	}
	return ""
}

func (m *UsergroupDet) GetTaskgroup() []*TaskgroupDet {
	if m != nil {
		return m.Taskgroup
	}
	return nil
}

func (m *UsergroupDet) GetTaskMap() *Taskmap {
	if m != nil {
		return m.TaskMap
	}
	return nil
}

func init() {
	proto.RegisterType((*UsergroupDet_KEYS)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.usergroups.usergroup.usergroup_det_KEYS")
	proto.RegisterType((*Taskid)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.usergroups.usergroup.taskid")
	proto.RegisterType((*Taskmap)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.usergroups.usergroup.taskmap")
	proto.RegisterType((*TaskgroupDet)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.usergroups.usergroup.taskgroup_det")
	proto.RegisterType((*UsergroupDet)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.usergroups.usergroup.usergroup_det")
}

func init() { proto.RegisterFile("usergroup_det.proto", fileDescriptor_3ef64b70d856c9bf) }

var fileDescriptor_3ef64b70d856c9bf = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x15, 0xfa, 0xff, 0xaa, 0x0a, 0x61, 0x90, 0xf0, 0x58, 0x65, 0xca, 0x94, 0xa1, 0x85,
	0x11, 0x26, 0x18, 0x10, 0x62, 0x49, 0x19, 0xe8, 0x82, 0x75, 0x8d, 0xad, 0x2a, 0xd0, 0xd4, 0x96,
	0x9d, 0x88, 0x4a, 0x7c, 0x60, 0x46, 0xbe, 0x02, 0xf2, 0xa5, 0x4d, 0xc9, 0x4a, 0xd8, 0xde, 0x39,
	0xe7, 0x7b, 0xf7, 0x7b, 0x91, 0xe1, 0xbc, 0x74, 0xca, 0xae, 0xad, 0x2e, 0x8d, 0x90, 0xaa, 0x88,
	0x8d, 0xd5, 0x85, 0x66, 0xd7, 0x69, 0xe6, 0x52, 0x2d, 0x32, 0xed, 0xc4, 0xce, 0x0a, 0x44, 0x14,
	0x1b, 0x9d, 0xe2, 0x46, 0x0a, 0x6d, 0x94, 0x8d, 0x11, 0x31, 0xae, 0x6f, 0xb9, 0xa3, 0x0c, 0x23,
	0x60, 0x8d, 0x69, 0xe2, 0xf1, 0x7e, 0xb9, 0x60, 0x0c, 0xba, 0x5b, 0xcc, 0x15, 0x0f, 0xa6, 0x41,
	0x34, 0x4a, 0x48, 0x87, 0x9f, 0xd0, 0x2f, 0xd0, 0xbd, 0x67, 0x92, 0x5d, 0xc2, 0xc0, 0x2b, 0x91,
	0xc9, 0x7d, 0x03, 0x7d, 0x78, 0x90, 0xfe, 0x9a, 0x55, 0x28, 0xf9, 0xc9, 0x34, 0x88, 0x86, 0x09,
	0x69, 0x76, 0x01, 0xbd, 0x0f, 0x9b, 0x15, 0x8a, 0x77, 0xe8, 0xb0, 0x2a, 0x18, 0x87, 0x81, 0xda,
	0xa9, 0xb4, 0x2c, 0x14, 0xef, 0xd2, 0xf9, 0xa1, 0xf4, 0xfd, 0x52, 0xad, 0xca, 0x35, 0xef, 0x55,
	0xfd, 0x54, 0x84, 0xaf, 0x95, 0x65, 0x8e, 0x86, 0x2d, 0xa0, 0xe7, 0xa5, 0xe3, 0xc1, 0xb4, 0x13,
	0x8d, 0x67, 0x37, 0xf1, 0x9f, 0xc0, 0xe3, 0x8a, 0x25, 0xa9, 0x66, 0x85, 0xdf, 0x01, 0x4c, 0xbc,
	0xaa, 0x73, 0xf0, 0x90, 0x1e, 0x5b, 0xec, 0xec, 0x01, 0xd2, 0x97, 0x2f, 0x96, 0xbd, 0xc1, 0x59,
	0xb6, 0x4d, 0x37, 0xa5, 0x54, 0x52, 0xec, 0x63, 0x70, 0x44, 0x3c, 0x9e, 0xdd, 0xb6, 0xd8, 0x25,
	0x47, 0x93, 0x9c, 0x1e, 0x06, 0x3f, 0x53, 0x9e, 0x8e, 0x2d, 0x61, 0x48, 0x16, 0x39, 0x1a, 0xca,
	0xaf, 0xbd, 0x05, 0xc5, 0xf8, 0x84, 0x26, 0xfc, 0x0a, 0x60, 0xd2, 0xf8, 0xf3, 0xbf, 0x89, 0x67,
	0x0d, 0xe2, 0x15, 0x8c, 0xea, 0x6c, 0xf8, 0x9c, 0x52, 0xbf, 0x6b, 0xb1, 0x46, 0xed, 0x98, 0x1c,
	0xc7, 0x36, 0x48, 0xaf, 0xfe, 0x95, 0x74, 0xd5, 0xa7, 0x07, 0x32, 0xff, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0x45, 0xd4, 0xa3, 0x37, 0x03, 0x00, 0x00,
}