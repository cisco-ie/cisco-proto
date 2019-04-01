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
// source: optics_edm_wave_info.proto

package cisco_ios_xr_controller_optics_oper_optics_oper_optics_ports_optics_port_optics_dwdm_carrier_channel_map

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

type OpticsEdmWaveInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpticsEdmWaveInfo_KEYS) Reset()         { *m = OpticsEdmWaveInfo_KEYS{} }
func (m *OpticsEdmWaveInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*OpticsEdmWaveInfo_KEYS) ProtoMessage()    {}
func (*OpticsEdmWaveInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc52397cdce6e3f2, []int{0}
}

func (m *OpticsEdmWaveInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpticsEdmWaveInfo_KEYS.Unmarshal(m, b)
}
func (m *OpticsEdmWaveInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpticsEdmWaveInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *OpticsEdmWaveInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpticsEdmWaveInfo_KEYS.Merge(m, src)
}
func (m *OpticsEdmWaveInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_OpticsEdmWaveInfo_KEYS.Size(m)
}
func (m *OpticsEdmWaveInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OpticsEdmWaveInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OpticsEdmWaveInfo_KEYS proto.InternalMessageInfo

func (m *OpticsEdmWaveInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type OpticsEdmChannelWavelengthMapInfo struct {
	ItuChanNum           uint32   `protobuf:"varint,1,opt,name=itu_chan_num,json=ituChanNum,proto3" json:"itu_chan_num,omitempty"`
	G694ChanNum          int32    `protobuf:"zigzag32,2,opt,name=g694_chan_num,json=g694ChanNum,proto3" json:"g694_chan_num,omitempty"`
	Frequency            string   `protobuf:"bytes,3,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Wavelength           string   `protobuf:"bytes,4,opt,name=wavelength,proto3" json:"wavelength,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpticsEdmChannelWavelengthMapInfo) Reset()         { *m = OpticsEdmChannelWavelengthMapInfo{} }
func (m *OpticsEdmChannelWavelengthMapInfo) String() string { return proto.CompactTextString(m) }
func (*OpticsEdmChannelWavelengthMapInfo) ProtoMessage()    {}
func (*OpticsEdmChannelWavelengthMapInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc52397cdce6e3f2, []int{1}
}

func (m *OpticsEdmChannelWavelengthMapInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpticsEdmChannelWavelengthMapInfo.Unmarshal(m, b)
}
func (m *OpticsEdmChannelWavelengthMapInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpticsEdmChannelWavelengthMapInfo.Marshal(b, m, deterministic)
}
func (m *OpticsEdmChannelWavelengthMapInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpticsEdmChannelWavelengthMapInfo.Merge(m, src)
}
func (m *OpticsEdmChannelWavelengthMapInfo) XXX_Size() int {
	return xxx_messageInfo_OpticsEdmChannelWavelengthMapInfo.Size(m)
}
func (m *OpticsEdmChannelWavelengthMapInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OpticsEdmChannelWavelengthMapInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OpticsEdmChannelWavelengthMapInfo proto.InternalMessageInfo

func (m *OpticsEdmChannelWavelengthMapInfo) GetItuChanNum() uint32 {
	if m != nil {
		return m.ItuChanNum
	}
	return 0
}

func (m *OpticsEdmChannelWavelengthMapInfo) GetG694ChanNum() int32 {
	if m != nil {
		return m.G694ChanNum
	}
	return 0
}

func (m *OpticsEdmChannelWavelengthMapInfo) GetFrequency() string {
	if m != nil {
		return m.Frequency
	}
	return ""
}

func (m *OpticsEdmChannelWavelengthMapInfo) GetWavelength() string {
	if m != nil {
		return m.Wavelength
	}
	return ""
}

type OpticsEdmWaveInfo struct {
	DwdmCarrierBand      string                               `protobuf:"bytes,50,opt,name=dwdm_carrier_band,json=dwdmCarrierBand,proto3" json:"dwdm_carrier_band,omitempty"`
	DwdmCarrierMin       uint32                               `protobuf:"varint,51,opt,name=dwdm_carrier_min,json=dwdmCarrierMin,proto3" json:"dwdm_carrier_min,omitempty"`
	DwdmCarrierMax       uint32                               `protobuf:"varint,52,opt,name=dwdm_carrier_max,json=dwdmCarrierMax,proto3" json:"dwdm_carrier_max,omitempty"`
	DwdmCarrierMapInfo   []*OpticsEdmChannelWavelengthMapInfo `protobuf:"bytes,53,rep,name=dwdm_carrier_map_info,json=dwdmCarrierMapInfo,proto3" json:"dwdm_carrier_map_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *OpticsEdmWaveInfo) Reset()         { *m = OpticsEdmWaveInfo{} }
func (m *OpticsEdmWaveInfo) String() string { return proto.CompactTextString(m) }
func (*OpticsEdmWaveInfo) ProtoMessage()    {}
func (*OpticsEdmWaveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc52397cdce6e3f2, []int{2}
}

func (m *OpticsEdmWaveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpticsEdmWaveInfo.Unmarshal(m, b)
}
func (m *OpticsEdmWaveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpticsEdmWaveInfo.Marshal(b, m, deterministic)
}
func (m *OpticsEdmWaveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpticsEdmWaveInfo.Merge(m, src)
}
func (m *OpticsEdmWaveInfo) XXX_Size() int {
	return xxx_messageInfo_OpticsEdmWaveInfo.Size(m)
}
func (m *OpticsEdmWaveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OpticsEdmWaveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OpticsEdmWaveInfo proto.InternalMessageInfo

func (m *OpticsEdmWaveInfo) GetDwdmCarrierBand() string {
	if m != nil {
		return m.DwdmCarrierBand
	}
	return ""
}

func (m *OpticsEdmWaveInfo) GetDwdmCarrierMin() uint32 {
	if m != nil {
		return m.DwdmCarrierMin
	}
	return 0
}

func (m *OpticsEdmWaveInfo) GetDwdmCarrierMax() uint32 {
	if m != nil {
		return m.DwdmCarrierMax
	}
	return 0
}

func (m *OpticsEdmWaveInfo) GetDwdmCarrierMapInfo() []*OpticsEdmChannelWavelengthMapInfo {
	if m != nil {
		return m.DwdmCarrierMapInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*OpticsEdmWaveInfo_KEYS)(nil), "cisco_ios_xr_controller_optics_oper.optics_oper.optics_ports.optics_port.optics_dwdm_carrier_channel_map.optics_edm_wave_info_KEYS")
	proto.RegisterType((*OpticsEdmChannelWavelengthMapInfo)(nil), "cisco_ios_xr_controller_optics_oper.optics_oper.optics_ports.optics_port.optics_dwdm_carrier_channel_map.optics_edm_channel_wavelength_map_info")
	proto.RegisterType((*OpticsEdmWaveInfo)(nil), "cisco_ios_xr_controller_optics_oper.optics_oper.optics_ports.optics_port.optics_dwdm_carrier_channel_map.optics_edm_wave_info")
}

func init() { proto.RegisterFile("optics_edm_wave_info.proto", fileDescriptor_cc52397cdce6e3f2) }

var fileDescriptor_cc52397cdce6e3f2 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x49, 0x5b, 0xfe, 0xd0, 0xe9, 0xbf, 0x6a, 0x17, 0x85, 0x28, 0x22, 0x21, 0x07, 0x09,
	0x1e, 0x22, 0xb4, 0x55, 0xf0, 0x6a, 0xf1, 0x20, 0xa2, 0x87, 0x78, 0xf2, 0xb4, 0x6c, 0x93, 0x6d,
	0xb3, 0x90, 0x9d, 0x8d, 0x9b, 0x8d, 0xad, 0x6f, 0xe1, 0x83, 0xf8, 0x18, 0xbe, 0x86, 0xef, 0x22,
	0xd9, 0x5a, 0x9a, 0xda, 0x1e, 0x3c, 0x79, 0x9b, 0x7c, 0xfb, 0xfb, 0xbe, 0xec, 0xcc, 0x0e, 0x1c,
	0xa9, 0xdc, 0x88, 0xb8, 0xa0, 0x3c, 0x91, 0x74, 0xc6, 0x5e, 0x38, 0x15, 0x38, 0x51, 0x61, 0xae,
	0x95, 0x51, 0x24, 0x8d, 0x45, 0x11, 0x2b, 0x2a, 0x54, 0x41, 0xe7, 0x9a, 0xc6, 0x0a, 0x8d, 0x56,
	0x59, 0xc6, 0x35, 0xfd, 0xf6, 0xa8, 0x9c, 0xeb, 0x70, 0x4b, 0x9d, 0x2b, 0x6d, 0x8a, 0xfa, 0xc7,
	0xb2, 0x4e, 0x66, 0x89, 0xa4, 0x31, 0xd3, 0x5a, 0x70, 0x4d, 0xe3, 0x94, 0x21, 0xf2, 0x8c, 0x4a,
	0x96, 0xfb, 0xe7, 0x70, 0xb8, 0xed, 0x1e, 0xf4, 0xee, 0xe6, 0xe9, 0x91, 0x10, 0x68, 0x21, 0x93,
	0xdc, 0x75, 0x3c, 0x27, 0x68, 0x47, 0xb6, 0xf6, 0xdf, 0x1d, 0x38, 0xad, 0x39, 0x96, 0x59, 0x95,
	0x33, 0xe3, 0x38, 0x35, 0x69, 0x15, 0x6b, 0x33, 0x88, 0x07, 0xff, 0x85, 0x29, 0x2d, 0x42, 0xb1,
	0x94, 0x36, 0xa6, 0x1b, 0x81, 0x30, 0xe5, 0x28, 0x65, 0xf8, 0x50, 0x4a, 0xe2, 0x43, 0x77, 0x7a,
	0x79, 0x35, 0x5c, 0x21, 0x0d, 0xcf, 0x09, 0x7a, 0x51, 0xa7, 0x12, 0x97, 0xcc, 0x31, 0xb4, 0x27,
	0x9a, 0x3f, 0x97, 0x1c, 0xe3, 0x57, 0xb7, 0x69, 0x6f, 0xb2, 0x12, 0xc8, 0x09, 0xc0, 0xea, 0xd7,
	0x6e, 0xcb, 0x1e, 0xd7, 0x14, 0xff, 0xb3, 0x01, 0xfb, 0xdb, 0x1a, 0x24, 0x67, 0xd0, 0x5b, 0x1b,
	0xca, 0x98, 0x61, 0xe2, 0xf6, 0xad, 0x7f, 0xb7, 0x3a, 0x18, 0x2d, 0xf4, 0x6b, 0x86, 0x09, 0x09,
	0x60, 0x6f, 0x8d, 0x95, 0x02, 0xdd, 0x81, 0x6d, 0x66, 0xa7, 0x86, 0xde, 0x0b, 0xdc, 0x24, 0xd9,
	0xdc, 0x1d, 0x6e, 0x92, 0x6c, 0x4e, 0x3e, 0x1c, 0x38, 0xf8, 0x81, 0x2e, 0xc6, 0xe6, 0x5e, 0x78,
	0xcd, 0xa0, 0xd3, 0x7f, 0x73, 0xc2, 0xbf, 0x5a, 0x82, 0xf0, 0x77, 0xef, 0x19, 0x91, 0xb5, 0x16,
	0xf2, 0x5b, 0x9c, 0xa8, 0xf1, 0x3f, 0xbb, 0xb0, 0x83, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9,
	0xab, 0x83, 0x04, 0xce, 0x02, 0x00, 0x00,
}
