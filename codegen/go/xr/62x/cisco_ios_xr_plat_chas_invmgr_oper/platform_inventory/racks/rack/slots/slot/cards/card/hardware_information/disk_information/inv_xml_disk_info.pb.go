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
// source: inv_xml_disk_info.proto

package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_slots_slot_cards_card_hardware_information_disk_information

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

type InvXmlDiskInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Name_1               string   `protobuf:"bytes,2,opt,name=name_1,json=name1,proto3" json:"name_1,omitempty"`
	Name_2               string   `protobuf:"bytes,3,opt,name=name_2,json=name2,proto3" json:"name_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvXmlDiskInfo_KEYS) Reset()         { *m = InvXmlDiskInfo_KEYS{} }
func (m *InvXmlDiskInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*InvXmlDiskInfo_KEYS) ProtoMessage()    {}
func (*InvXmlDiskInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdd203fbe2c1fb1, []int{0}
}

func (m *InvXmlDiskInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvXmlDiskInfo_KEYS.Unmarshal(m, b)
}
func (m *InvXmlDiskInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvXmlDiskInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *InvXmlDiskInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvXmlDiskInfo_KEYS.Merge(m, src)
}
func (m *InvXmlDiskInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_InvXmlDiskInfo_KEYS.Size(m)
}
func (m *InvXmlDiskInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InvXmlDiskInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InvXmlDiskInfo_KEYS proto.InternalMessageInfo

func (m *InvXmlDiskInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvXmlDiskInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *InvXmlDiskInfo_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

type InvXmlDiskAttribute struct {
	DiskName             string   `protobuf:"bytes,1,opt,name=disk_name,json=diskName,proto3" json:"disk_name,omitempty"`
	DiskSize             uint32   `protobuf:"varint,2,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
	SectorSize           uint32   `protobuf:"varint,3,opt,name=sector_size,json=sectorSize,proto3" json:"sector_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvXmlDiskAttribute) Reset()         { *m = InvXmlDiskAttribute{} }
func (m *InvXmlDiskAttribute) String() string { return proto.CompactTextString(m) }
func (*InvXmlDiskAttribute) ProtoMessage()    {}
func (*InvXmlDiskAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdd203fbe2c1fb1, []int{1}
}

func (m *InvXmlDiskAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvXmlDiskAttribute.Unmarshal(m, b)
}
func (m *InvXmlDiskAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvXmlDiskAttribute.Marshal(b, m, deterministic)
}
func (m *InvXmlDiskAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvXmlDiskAttribute.Merge(m, src)
}
func (m *InvXmlDiskAttribute) XXX_Size() int {
	return xxx_messageInfo_InvXmlDiskAttribute.Size(m)
}
func (m *InvXmlDiskAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_InvXmlDiskAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_InvXmlDiskAttribute proto.InternalMessageInfo

func (m *InvXmlDiskAttribute) GetDiskName() string {
	if m != nil {
		return m.DiskName
	}
	return ""
}

func (m *InvXmlDiskAttribute) GetDiskSize() uint32 {
	if m != nil {
		return m.DiskSize
	}
	return 0
}

func (m *InvXmlDiskAttribute) GetSectorSize() uint32 {
	if m != nil {
		return m.SectorSize
	}
	return 0
}

type InvXmlDiskInfo struct {
	DiskName             string                 `protobuf:"bytes,50,opt,name=disk_name,json=diskName,proto3" json:"disk_name,omitempty"`
	DiskSize             uint32                 `protobuf:"varint,51,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
	SectorSize           uint32                 `protobuf:"varint,52,opt,name=sector_size,json=sectorSize,proto3" json:"sector_size,omitempty"`
	Disks                []*InvXmlDiskAttribute `protobuf:"bytes,53,rep,name=disks,proto3" json:"disks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *InvXmlDiskInfo) Reset()         { *m = InvXmlDiskInfo{} }
func (m *InvXmlDiskInfo) String() string { return proto.CompactTextString(m) }
func (*InvXmlDiskInfo) ProtoMessage()    {}
func (*InvXmlDiskInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdd203fbe2c1fb1, []int{2}
}

func (m *InvXmlDiskInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvXmlDiskInfo.Unmarshal(m, b)
}
func (m *InvXmlDiskInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvXmlDiskInfo.Marshal(b, m, deterministic)
}
func (m *InvXmlDiskInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvXmlDiskInfo.Merge(m, src)
}
func (m *InvXmlDiskInfo) XXX_Size() int {
	return xxx_messageInfo_InvXmlDiskInfo.Size(m)
}
func (m *InvXmlDiskInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InvXmlDiskInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InvXmlDiskInfo proto.InternalMessageInfo

func (m *InvXmlDiskInfo) GetDiskName() string {
	if m != nil {
		return m.DiskName
	}
	return ""
}

func (m *InvXmlDiskInfo) GetDiskSize() uint32 {
	if m != nil {
		return m.DiskSize
	}
	return 0
}

func (m *InvXmlDiskInfo) GetSectorSize() uint32 {
	if m != nil {
		return m.SectorSize
	}
	return 0
}

func (m *InvXmlDiskInfo) GetDisks() []*InvXmlDiskAttribute {
	if m != nil {
		return m.Disks
	}
	return nil
}

func init() {
	proto.RegisterType((*InvXmlDiskInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.hardware_information.disk_information.inv_xml_disk_info_KEYS")
	proto.RegisterType((*InvXmlDiskAttribute)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.hardware_information.disk_information.inv_xml_disk_attribute")
	proto.RegisterType((*InvXmlDiskInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.hardware_information.disk_information.inv_xml_disk_info")
}

func init() { proto.RegisterFile("inv_xml_disk_info.proto", fileDescriptor_7fdd203fbe2c1fb1) }

var fileDescriptor_7fdd203fbe2c1fb1 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xcf, 0x4e, 0x02, 0x31,
	0x10, 0x87, 0xb3, 0x20, 0x44, 0x4a, 0x3c, 0xd8, 0x44, 0xdd, 0xc4, 0x83, 0x84, 0x13, 0xa7, 0x26,
	0x2c, 0xfa, 0x08, 0x9e, 0x4c, 0x3c, 0xc0, 0x49, 0x2f, 0x4d, 0xe9, 0x16, 0x69, 0xd8, 0xed, 0x6c,
	0x66, 0x2a, 0x22, 0xf1, 0x31, 0x3c, 0xfa, 0x06, 0xbe, 0xa4, 0x69, 0x57, 0x51, 0xc1, 0x70, 0xf5,
	0x32, 0xbb, 0xfb, 0xfd, 0x26, 0xfb, 0x4d, 0xff, 0xb0, 0x33, 0xeb, 0x96, 0x72, 0x55, 0x16, 0x32,
	0xb7, 0xb4, 0x90, 0xd6, 0xcd, 0x40, 0x54, 0x08, 0x1e, 0xf8, 0x8b, 0xb6, 0xa4, 0x41, 0x5a, 0x20,
	0xb9, 0x42, 0x59, 0x15, 0xca, 0x4b, 0x3d, 0x57, 0x24, 0xad, 0x5b, 0x96, 0x0f, 0x28, 0xa1, 0x32,
	0x28, 0x02, 0x9d, 0x01, 0x96, 0x01, 0x1a, 0xe7, 0x01, 0x9f, 0x05, 0x2a, 0xbd, 0xa0, 0x58, 0x05,
	0x15, 0xe0, 0x29, 0x56, 0xa1, 0x15, 0xe6, 0x14, 0xab, 0x98, 0x2b, 0xcc, 0x9f, 0x14, 0x9a, 0xe8,
	0xc2, 0x52, 0x79, 0x0b, 0x4e, 0x6c, 0xe4, 0x9f, 0xa0, 0x7f, 0xcf, 0x4e, 0x77, 0x06, 0x93, 0x37,
	0xd7, 0x77, 0x13, 0xce, 0xd9, 0x81, 0x53, 0xa5, 0x49, 0x93, 0x5e, 0x32, 0xe8, 0x8c, 0xe3, 0x3b,
	0x3f, 0x61, 0xed, 0xf0, 0x94, 0xc3, 0xb4, 0x11, 0x69, 0x2b, 0x7c, 0x0d, 0x37, 0x38, 0x4b, 0x9b,
	0xdf, 0x38, 0xeb, 0xd3, 0xd6, 0xbf, 0x95, 0xf7, 0x68, 0xa7, 0x8f, 0xde, 0xf0, 0x73, 0xd6, 0x89,
	0xe4, 0x87, 0xe0, 0x30, 0x80, 0xdb, 0x20, 0xf9, 0x0a, 0xc9, 0xae, 0x4d, 0xf4, 0x1c, 0xd5, 0xe1,
	0xc4, 0xae, 0x0d, 0xbf, 0x60, 0x5d, 0x32, 0xda, 0x03, 0xd6, 0x71, 0x33, 0xc6, 0xac, 0x46, 0xa1,
	0xa1, 0xff, 0xd6, 0x60, 0xc7, 0x3b, 0x2b, 0xfa, 0x2d, 0xcc, 0xf6, 0x09, 0x47, 0xfb, 0x85, 0x97,
	0xdb, 0x42, 0xfe, 0x9e, 0xb0, 0x56, 0xe8, 0xa6, 0xf4, 0xaa, 0xd7, 0x1c, 0x74, 0xb3, 0xd7, 0x44,
	0xfc, 0xe7, 0x89, 0x8a, 0xbf, 0xb7, 0x7c, 0x5c, 0xcf, 0x38, 0x6d, 0xc7, 0x4b, 0x37, 0xfa, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x2b, 0x59, 0xe8, 0x19, 0x8f, 0x02, 0x00, 0x00,
}