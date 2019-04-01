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
// source: isis_sh_mesh_groups.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_mesh_groups

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

type IsisShMeshGroups_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShMeshGroups_KEYS) Reset()         { *m = IsisShMeshGroups_KEYS{} }
func (m *IsisShMeshGroups_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShMeshGroups_KEYS) ProtoMessage()    {}
func (*IsisShMeshGroups_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef6f21df520a3481, []int{0}
}

func (m *IsisShMeshGroups_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShMeshGroups_KEYS.Unmarshal(m, b)
}
func (m *IsisShMeshGroups_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShMeshGroups_KEYS.Marshal(b, m, deterministic)
}
func (m *IsisShMeshGroups_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShMeshGroups_KEYS.Merge(m, src)
}
func (m *IsisShMeshGroups_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShMeshGroups_KEYS.Size(m)
}
func (m *IsisShMeshGroups_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShMeshGroups_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShMeshGroups_KEYS proto.InternalMessageInfo

func (m *IsisShMeshGroups_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type IsisShMeshEntryItem struct {
	MeshGroupInterface   string   `protobuf:"bytes,1,opt,name=mesh_group_interface,json=meshGroupInterface,proto3" json:"mesh_group_interface,omitempty"`
	MeshGroupNumber      uint32   `protobuf:"varint,2,opt,name=mesh_group_number,json=meshGroupNumber,proto3" json:"mesh_group_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShMeshEntryItem) Reset()         { *m = IsisShMeshEntryItem{} }
func (m *IsisShMeshEntryItem) String() string { return proto.CompactTextString(m) }
func (*IsisShMeshEntryItem) ProtoMessage()    {}
func (*IsisShMeshEntryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef6f21df520a3481, []int{1}
}

func (m *IsisShMeshEntryItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShMeshEntryItem.Unmarshal(m, b)
}
func (m *IsisShMeshEntryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShMeshEntryItem.Marshal(b, m, deterministic)
}
func (m *IsisShMeshEntryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShMeshEntryItem.Merge(m, src)
}
func (m *IsisShMeshEntryItem) XXX_Size() int {
	return xxx_messageInfo_IsisShMeshEntryItem.Size(m)
}
func (m *IsisShMeshEntryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShMeshEntryItem.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShMeshEntryItem proto.InternalMessageInfo

func (m *IsisShMeshEntryItem) GetMeshGroupInterface() string {
	if m != nil {
		return m.MeshGroupInterface
	}
	return ""
}

func (m *IsisShMeshEntryItem) GetMeshGroupNumber() uint32 {
	if m != nil {
		return m.MeshGroupNumber
	}
	return 0
}

type IsisShMeshEntryEntry struct {
	IsisShMeshEntry      []*IsisShMeshEntryItem `protobuf:"bytes,1,rep,name=isis_sh_mesh_entry,json=isisShMeshEntry,proto3" json:"isis_sh_mesh_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *IsisShMeshEntryEntry) Reset()         { *m = IsisShMeshEntryEntry{} }
func (m *IsisShMeshEntryEntry) String() string { return proto.CompactTextString(m) }
func (*IsisShMeshEntryEntry) ProtoMessage()    {}
func (*IsisShMeshEntryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef6f21df520a3481, []int{2}
}

func (m *IsisShMeshEntryEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShMeshEntryEntry.Unmarshal(m, b)
}
func (m *IsisShMeshEntryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShMeshEntryEntry.Marshal(b, m, deterministic)
}
func (m *IsisShMeshEntryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShMeshEntryEntry.Merge(m, src)
}
func (m *IsisShMeshEntryEntry) XXX_Size() int {
	return xxx_messageInfo_IsisShMeshEntryEntry.Size(m)
}
func (m *IsisShMeshEntryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShMeshEntryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShMeshEntryEntry proto.InternalMessageInfo

func (m *IsisShMeshEntryEntry) GetIsisShMeshEntry() []*IsisShMeshEntryItem {
	if m != nil {
		return m.IsisShMeshEntry
	}
	return nil
}

type IsisShMeshGroups struct {
	MeshGroupConfiguredInterfaceList *IsisShMeshEntryEntry `protobuf:"bytes,50,opt,name=mesh_group_configured_interface_list,json=meshGroupConfiguredInterfaceList,proto3" json:"mesh_group_configured_interface_list,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}              `json:"-"`
	XXX_unrecognized                 []byte                `json:"-"`
	XXX_sizecache                    int32                 `json:"-"`
}

func (m *IsisShMeshGroups) Reset()         { *m = IsisShMeshGroups{} }
func (m *IsisShMeshGroups) String() string { return proto.CompactTextString(m) }
func (*IsisShMeshGroups) ProtoMessage()    {}
func (*IsisShMeshGroups) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef6f21df520a3481, []int{3}
}

func (m *IsisShMeshGroups) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShMeshGroups.Unmarshal(m, b)
}
func (m *IsisShMeshGroups) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShMeshGroups.Marshal(b, m, deterministic)
}
func (m *IsisShMeshGroups) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShMeshGroups.Merge(m, src)
}
func (m *IsisShMeshGroups) XXX_Size() int {
	return xxx_messageInfo_IsisShMeshGroups.Size(m)
}
func (m *IsisShMeshGroups) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShMeshGroups.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShMeshGroups proto.InternalMessageInfo

func (m *IsisShMeshGroups) GetMeshGroupConfiguredInterfaceList() *IsisShMeshEntryEntry {
	if m != nil {
		return m.MeshGroupConfiguredInterfaceList
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShMeshGroups_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.mesh_groups.isis_sh_mesh_groups_KEYS")
	proto.RegisterType((*IsisShMeshEntryItem)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.mesh_groups.isis_sh_mesh_entry_item")
	proto.RegisterType((*IsisShMeshEntryEntry)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.mesh_groups.isis_sh_mesh_entry_entry")
	proto.RegisterType((*IsisShMeshGroups)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.mesh_groups.isis_sh_mesh_groups")
}

func init() { proto.RegisterFile("isis_sh_mesh_groups.proto", fileDescriptor_ef6f21df520a3481) }

var fileDescriptor_ef6f21df520a3481 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x66, 0x15, 0x04, 0xb7, 0x96, 0xe2, 0x2a, 0x18, 0x6f, 0x21, 0x7a, 0x08, 0x1e, 0x16, 0xa9,
	0x0f, 0xd0, 0x83, 0x14, 0x11, 0xb5, 0x87, 0xf4, 0x62, 0x4f, 0x43, 0xba, 0x4e, 0x9b, 0x85, 0x66,
	0xb7, 0xec, 0x6c, 0x50, 0x1f, 0xc0, 0x17, 0xf1, 0xe6, 0x63, 0xf8, 0x66, 0x92, 0x48, 0x7e, 0xd4,
	0x7a, 0x12, 0x2f, 0xc3, 0xf0, 0x7d, 0x33, 0x5f, 0xbe, 0xf9, 0xb2, 0xfc, 0x58, 0x93, 0x26, 0xa0,
	0x0c, 0x72, 0xa4, 0x0c, 0x96, 0xce, 0x16, 0x6b, 0x92, 0x6b, 0x67, 0xbd, 0x15, 0x23, 0xa5, 0x49,
	0x59, 0xd0, 0x96, 0xe0, 0xc9, 0x81, 0x5a, 0x19, 0x82, 0x6a, 0xd8, 0xae, 0xd1, 0xc9, 0xb2, 0x93,
	0xda, 0x90, 0x4f, 0x8d, 0xc2, 0xb6, 0x93, 0x1d, 0x99, 0x68, 0xc4, 0x83, 0x0d, 0xea, 0x70, 0x33,
	0x9e, 0x4d, 0xc5, 0x09, 0xef, 0xd7, 0x3b, 0x60, 0xd2, 0x1c, 0x03, 0x16, 0xb2, 0x78, 0x37, 0xd9,
	0xab, 0xc1, 0x49, 0x9a, 0x63, 0xf4, 0xc8, 0x8f, 0xbe, 0x08, 0xa0, 0xf1, 0xee, 0x19, 0xb4, 0xc7,
	0x5c, 0x9c, 0xf3, 0xc3, 0x56, 0x13, 0xb4, 0xf1, 0xe8, 0x16, 0xa9, 0xaa, 0x65, 0x44, 0xc9, 0x5d,
	0x95, 0xd4, 0x75, 0xcd, 0x88, 0x33, 0xbe, 0xdf, 0xd9, 0x30, 0x45, 0x3e, 0x47, 0x17, 0x6c, 0x85,
	0x2c, 0xee, 0x27, 0x83, 0x66, 0x7c, 0x52, 0xc1, 0xd1, 0x2b, 0xfb, 0x66, 0xfd, 0xf3, 0xcb, 0x55,
	0x15, 0x2f, 0x8c, 0x8b, 0x9f, 0x64, 0xc0, 0xc2, 0xed, 0xb8, 0x37, 0xbc, 0x97, 0x7f, 0x4c, 0x4d,
	0xfe, 0x72, 0x71, 0x32, 0x28, 0x89, 0x69, 0x76, 0x87, 0x94, 0x8d, 0x4b, 0x34, 0x7a, 0x67, 0xfc,
	0x60, 0x43, 0xbe, 0xe2, 0x8d, 0xf1, 0xd3, 0xce, 0xa5, 0xca, 0x9a, 0x85, 0x5e, 0x16, 0x0e, 0x1f,
	0xda, 0x98, 0x60, 0xa5, 0xc9, 0x07, 0xc3, 0x90, 0xc5, 0xbd, 0xe1, 0xec, 0x3f, 0x1c, 0x57, 0x35,
	0x09, 0x9b, 0x5c, 0x2f, 0x1b, 0x13, 0xcd, 0x0f, 0xb9, 0xd5, 0xe4, 0xe7, 0x3b, 0xd5, 0x53, 0xbb,
	0xf8, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x55, 0x7a, 0xa2, 0x87, 0x02, 0x00, 0x00,
}