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
// source: inventory.proto

package cisco_ios_xr_installmgr_admin_oper_install_software_inventory_inactive_inventories_inventory

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

type Inventory_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Inventory_KEYS) Reset()         { *m = Inventory_KEYS{} }
func (m *Inventory_KEYS) String() string { return proto.CompactTextString(m) }
func (*Inventory_KEYS) ProtoMessage()    {}
func (*Inventory_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7173caedb7c6ae96, []int{0}
}

func (m *Inventory_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inventory_KEYS.Unmarshal(m, b)
}
func (m *Inventory_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inventory_KEYS.Marshal(b, m, deterministic)
}
func (m *Inventory_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inventory_KEYS.Merge(m, src)
}
func (m *Inventory_KEYS) XXX_Size() int {
	return xxx_messageInfo_Inventory_KEYS.Size(m)
}
func (m *Inventory_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Inventory_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Inventory_KEYS proto.InternalMessageInfo

func (m *Inventory_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type PkgGroup struct {
	DeviceName           string   `protobuf:"bytes,1,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PkgGroup) Reset()         { *m = PkgGroup{} }
func (m *PkgGroup) String() string { return proto.CompactTextString(m) }
func (*PkgGroup) ProtoMessage()    {}
func (*PkgGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_7173caedb7c6ae96, []int{1}
}

func (m *PkgGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PkgGroup.Unmarshal(m, b)
}
func (m *PkgGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PkgGroup.Marshal(b, m, deterministic)
}
func (m *PkgGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PkgGroup.Merge(m, src)
}
func (m *PkgGroup) XXX_Size() int {
	return xxx_messageInfo_PkgGroup.Size(m)
}
func (m *PkgGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PkgGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PkgGroup proto.InternalMessageInfo

func (m *PkgGroup) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *PkgGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PkgInfo struct {
	Package              *PkgGroup `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	Version              string    `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	BuildInformation     string    `protobuf:"bytes,3,opt,name=build_information,json=buildInformation,proto3" json:"build_information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PkgInfo) Reset()         { *m = PkgInfo{} }
func (m *PkgInfo) String() string { return proto.CompactTextString(m) }
func (*PkgInfo) ProtoMessage()    {}
func (*PkgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7173caedb7c6ae96, []int{2}
}

func (m *PkgInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PkgInfo.Unmarshal(m, b)
}
func (m *PkgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PkgInfo.Marshal(b, m, deterministic)
}
func (m *PkgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PkgInfo.Merge(m, src)
}
func (m *PkgInfo) XXX_Size() int {
	return xxx_messageInfo_PkgInfo.Size(m)
}
func (m *PkgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PkgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PkgInfo proto.InternalMessageInfo

func (m *PkgInfo) GetPackage() *PkgGroup {
	if m != nil {
		return m.Package
	}
	return nil
}

func (m *PkgInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PkgInfo) GetBuildInformation() string {
	if m != nil {
		return m.BuildInformation
	}
	return ""
}

type Inventory struct {
	Major                  uint32     `protobuf:"varint,50,opt,name=major,proto3" json:"major,omitempty"`
	Minor                  uint32     `protobuf:"varint,51,opt,name=minor,proto3" json:"minor,omitempty"`
	BootImageName          string     `protobuf:"bytes,52,opt,name=boot_image_name,json=bootImageName,proto3" json:"boot_image_name,omitempty"`
	LoadPath               []*PkgInfo `protobuf:"bytes,53,rep,name=load_path,json=loadPath,proto3" json:"load_path,omitempty"`
	NodeType               uint64     `protobuf:"varint,54,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	SecureDomainRouterName string     `protobuf:"bytes,55,opt,name=secure_domain_router_name,json=secureDomainRouterName,proto3" json:"secure_domain_router_name,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}   `json:"-"`
	XXX_unrecognized       []byte     `json:"-"`
	XXX_sizecache          int32      `json:"-"`
}

func (m *Inventory) Reset()         { *m = Inventory{} }
func (m *Inventory) String() string { return proto.CompactTextString(m) }
func (*Inventory) ProtoMessage()    {}
func (*Inventory) Descriptor() ([]byte, []int) {
	return fileDescriptor_7173caedb7c6ae96, []int{3}
}

func (m *Inventory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inventory.Unmarshal(m, b)
}
func (m *Inventory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inventory.Marshal(b, m, deterministic)
}
func (m *Inventory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inventory.Merge(m, src)
}
func (m *Inventory) XXX_Size() int {
	return xxx_messageInfo_Inventory.Size(m)
}
func (m *Inventory) XXX_DiscardUnknown() {
	xxx_messageInfo_Inventory.DiscardUnknown(m)
}

var xxx_messageInfo_Inventory proto.InternalMessageInfo

func (m *Inventory) GetMajor() uint32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *Inventory) GetMinor() uint32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *Inventory) GetBootImageName() string {
	if m != nil {
		return m.BootImageName
	}
	return ""
}

func (m *Inventory) GetLoadPath() []*PkgInfo {
	if m != nil {
		return m.LoadPath
	}
	return nil
}

func (m *Inventory) GetNodeType() uint64 {
	if m != nil {
		return m.NodeType
	}
	return 0
}

func (m *Inventory) GetSecureDomainRouterName() string {
	if m != nil {
		return m.SecureDomainRouterName
	}
	return ""
}

func init() {
	proto.RegisterType((*Inventory_KEYS)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.inactive.inventories.inventory.inventory_KEYS")
	proto.RegisterType((*PkgGroup)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.inactive.inventories.inventory.pkg_group")
	proto.RegisterType((*PkgInfo)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.inactive.inventories.inventory.pkg_info")
	proto.RegisterType((*Inventory)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.inactive.inventories.inventory.inventory")
}

func init() { proto.RegisterFile("inventory.proto", fileDescriptor_7173caedb7c6ae96) }

var fileDescriptor_7173caedb7c6ae96 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0xc9, 0xbd, 0x57, 0xdb, 0x9c, 0x52, 0xab, 0x83, 0x48, 0xc4, 0x85, 0x21, 0x0b, 0x09,
	0x88, 0x59, 0xb4, 0xfe, 0xc1, 0x9d, 0x0b, 0x5d, 0x14, 0x41, 0x24, 0xba, 0x11, 0x84, 0x61, 0x9a,
	0x4c, 0xd3, 0xb1, 0x99, 0x39, 0xc3, 0x64, 0x12, 0xcd, 0x52, 0x7c, 0x2e, 0x1f, 0xc3, 0xf7, 0x91,
	0x99, 0xa4, 0x29, 0xf7, 0x01, 0xba, 0xcb, 0xf9, 0x7d, 0x5f, 0xe6, 0x9c, 0xef, 0x70, 0x60, 0x25,
	0x54, 0xc7, 0x95, 0x45, 0xd3, 0x67, 0xda, 0xa0, 0x45, 0xf2, 0xbd, 0x10, 0x4d, 0x81, 0x54, 0x60,
	0x43, 0x7f, 0x19, 0x2a, 0x54, 0x63, 0x59, 0x5d, 0xcb, 0xca, 0x50, 0x56, 0x4a, 0xa1, 0x28, 0x6a,
	0x6e, 0xb2, 0x91, 0x66, 0x0d, 0xee, 0xed, 0x4f, 0x66, 0x38, 0x3d, 0x3f, 0x22, 0x14, 0x2b, 0xac,
	0xe8, 0x78, 0x76, 0x42, 0x82, 0x37, 0xd3, 0x77, 0x9f, 0xbc, 0x80, 0x7b, 0x53, 0x41, 0x3f, 0x7e,
	0xf8, 0xf6, 0x85, 0x3c, 0x81, 0x50, 0x61, 0xc9, 0xa9, 0x62, 0x92, 0x47, 0x41, 0x1c, 0xa4, 0x61,
	0x3e, 0x77, 0xe0, 0x13, 0x93, 0x3c, 0x79, 0x07, 0xa1, 0x3e, 0x56, 0xb4, 0x32, 0xd8, 0x6a, 0xf2,
	0x14, 0x16, 0x25, 0xef, 0x44, 0x71, 0xcb, 0x0b, 0x03, 0x72, 0x6e, 0x42, 0xe0, 0xc6, 0x2b, 0x57,
	0x5e, 0xf1, 0xdf, 0xc9, 0xbf, 0x00, 0xe6, 0xee, 0x09, 0xa1, 0xf6, 0x48, 0x7e, 0x07, 0x30, 0xd3,
	0xac, 0x38, 0xb2, 0x6a, 0xf8, 0x7d, 0xb1, 0xae, 0xb2, 0x4b, 0xc6, 0xcd, 0xa6, 0xe1, 0xf3, 0x53,
	0x5f, 0x12, 0xc1, 0xac, 0xe3, 0xa6, 0x11, 0xa8, 0xc6, 0x39, 0x4f, 0x25, 0x79, 0x0e, 0x0f, 0x76,
	0xad, 0xa8, 0x4b, 0x3f, 0xab, 0x91, 0xcc, 0x3a, 0xcf, 0xb5, 0xf7, 0xdc, 0xf7, 0xc2, 0xf6, 0xcc,
	0x93, 0xbf, 0x57, 0x10, 0x4e, 0x7d, 0xc8, 0x43, 0xb8, 0x23, 0xd9, 0x0f, 0x34, 0xd1, 0x3a, 0x0e,
	0xd2, 0x65, 0x3e, 0x14, 0x9e, 0x0a, 0x85, 0x26, 0xda, 0x8c, 0xd4, 0x15, 0xe4, 0x19, 0xac, 0x76,
	0x88, 0x96, 0x0a, 0xc9, 0xaa, 0x71, 0x95, 0x2f, 0x7d, 0x93, 0xa5, 0xc3, 0x5b, 0x47, 0xfd, 0x36,
	0xff, 0x04, 0x10, 0xd6, 0xc8, 0x4a, 0xaa, 0x99, 0x3d, 0x44, 0xaf, 0xe2, 0xeb, 0x74, 0xb1, 0xde,
	0x5f, 0x7e, 0x5d, 0x2e, 0x7c, 0x3e, 0x77, 0x8d, 0x3f, 0x33, 0x7b, 0x98, 0xce, 0xc3, 0xf6, 0x9a,
	0x47, 0xaf, 0xe3, 0x20, 0xbd, 0x19, 0xce, 0xe3, 0x6b, 0xaf, 0x39, 0x79, 0x0b, 0x8f, 0x1b, 0x5e,
	0xb4, 0x86, 0xd3, 0x12, 0x25, 0x13, 0x8a, 0x1a, 0x6c, 0x2d, 0x37, 0x43, 0xa8, 0x37, 0x3e, 0xd4,
	0xa3, 0xc1, 0xf0, 0xde, 0xeb, 0xb9, 0x97, 0x5d, 0xba, 0xdd, 0x5d, 0x7f, 0xed, 0x9b, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x5b, 0x47, 0xdc, 0xb0, 0x00, 0x03, 0x00, 0x00,
}
