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
// source: icpe_oper_sat_remote_version_type.proto

package cisco_ios_xr_icpe_infra_oper_nv_satellite_satellite_versions_satellite_version

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

type IcpeOperSatRemoteVersionType_KEYS struct {
	SatelliteId          uint32   `protobuf:"varint,1,opt,name=satellite_id,json=satelliteId,proto3" json:"satellite_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IcpeOperSatRemoteVersionType_KEYS) Reset()         { *m = IcpeOperSatRemoteVersionType_KEYS{} }
func (m *IcpeOperSatRemoteVersionType_KEYS) String() string { return proto.CompactTextString(m) }
func (*IcpeOperSatRemoteVersionType_KEYS) ProtoMessage()    {}
func (*IcpeOperSatRemoteVersionType_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a488b685ba3cbf, []int{0}
}

func (m *IcpeOperSatRemoteVersionType_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IcpeOperSatRemoteVersionType_KEYS.Unmarshal(m, b)
}
func (m *IcpeOperSatRemoteVersionType_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IcpeOperSatRemoteVersionType_KEYS.Marshal(b, m, deterministic)
}
func (m *IcpeOperSatRemoteVersionType_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcpeOperSatRemoteVersionType_KEYS.Merge(m, src)
}
func (m *IcpeOperSatRemoteVersionType_KEYS) XXX_Size() int {
	return xxx_messageInfo_IcpeOperSatRemoteVersionType_KEYS.Size(m)
}
func (m *IcpeOperSatRemoteVersionType_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IcpeOperSatRemoteVersionType_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IcpeOperSatRemoteVersionType_KEYS proto.InternalMessageInfo

func (m *IcpeOperSatRemoteVersionType_KEYS) GetSatelliteId() uint32 {
	if m != nil {
		return m.SatelliteId
	}
	return 0
}

type IcpeOperSatVersion struct {
	VersionCheckState    string   `protobuf:"bytes,1,opt,name=version_check_state,json=versionCheckState,proto3" json:"version_check_state,omitempty"`
	RemoteVersionPresent bool     `protobuf:"varint,2,opt,name=remote_version_present,json=remoteVersionPresent,proto3" json:"remote_version_present,omitempty"`
	RemoteVersion        []string `protobuf:"bytes,3,rep,name=remote_version,json=remoteVersion,proto3" json:"remote_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IcpeOperSatVersion) Reset()         { *m = IcpeOperSatVersion{} }
func (m *IcpeOperSatVersion) String() string { return proto.CompactTextString(m) }
func (*IcpeOperSatVersion) ProtoMessage()    {}
func (*IcpeOperSatVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a488b685ba3cbf, []int{1}
}

func (m *IcpeOperSatVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IcpeOperSatVersion.Unmarshal(m, b)
}
func (m *IcpeOperSatVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IcpeOperSatVersion.Marshal(b, m, deterministic)
}
func (m *IcpeOperSatVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcpeOperSatVersion.Merge(m, src)
}
func (m *IcpeOperSatVersion) XXX_Size() int {
	return xxx_messageInfo_IcpeOperSatVersion.Size(m)
}
func (m *IcpeOperSatVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_IcpeOperSatVersion.DiscardUnknown(m)
}

var xxx_messageInfo_IcpeOperSatVersion proto.InternalMessageInfo

func (m *IcpeOperSatVersion) GetVersionCheckState() string {
	if m != nil {
		return m.VersionCheckState
	}
	return ""
}

func (m *IcpeOperSatVersion) GetRemoteVersionPresent() bool {
	if m != nil {
		return m.RemoteVersionPresent
	}
	return false
}

func (m *IcpeOperSatVersion) GetRemoteVersion() []string {
	if m != nil {
		return m.RemoteVersion
	}
	return nil
}

type IcpeOperSatRemoteVersionType struct {
	SatelliteIdXr        uint32              `protobuf:"varint,50,opt,name=satellite_id_xr,json=satelliteIdXr,proto3" json:"satellite_id_xr,omitempty"`
	VersionCheckState    string              `protobuf:"bytes,51,opt,name=version_check_state,json=versionCheckState,proto3" json:"version_check_state,omitempty"`
	RemoteVersionPresent bool                `protobuf:"varint,52,opt,name=remote_version_present,json=remoteVersionPresent,proto3" json:"remote_version_present,omitempty"`
	RemoteVersion        []string            `protobuf:"bytes,53,rep,name=remote_version,json=remoteVersion,proto3" json:"remote_version,omitempty"`
	ActiveVersion        *IcpeOperSatVersion `protobuf:"bytes,54,opt,name=active_version,json=activeVersion,proto3" json:"active_version,omitempty"`
	TransferredVersion   *IcpeOperSatVersion `protobuf:"bytes,55,opt,name=transferred_version,json=transferredVersion,proto3" json:"transferred_version,omitempty"`
	CommittedVersion     *IcpeOperSatVersion `protobuf:"bytes,56,opt,name=committed_version,json=committedVersion,proto3" json:"committed_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IcpeOperSatRemoteVersionType) Reset()         { *m = IcpeOperSatRemoteVersionType{} }
func (m *IcpeOperSatRemoteVersionType) String() string { return proto.CompactTextString(m) }
func (*IcpeOperSatRemoteVersionType) ProtoMessage()    {}
func (*IcpeOperSatRemoteVersionType) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a488b685ba3cbf, []int{2}
}

func (m *IcpeOperSatRemoteVersionType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IcpeOperSatRemoteVersionType.Unmarshal(m, b)
}
func (m *IcpeOperSatRemoteVersionType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IcpeOperSatRemoteVersionType.Marshal(b, m, deterministic)
}
func (m *IcpeOperSatRemoteVersionType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcpeOperSatRemoteVersionType.Merge(m, src)
}
func (m *IcpeOperSatRemoteVersionType) XXX_Size() int {
	return xxx_messageInfo_IcpeOperSatRemoteVersionType.Size(m)
}
func (m *IcpeOperSatRemoteVersionType) XXX_DiscardUnknown() {
	xxx_messageInfo_IcpeOperSatRemoteVersionType.DiscardUnknown(m)
}

var xxx_messageInfo_IcpeOperSatRemoteVersionType proto.InternalMessageInfo

func (m *IcpeOperSatRemoteVersionType) GetSatelliteIdXr() uint32 {
	if m != nil {
		return m.SatelliteIdXr
	}
	return 0
}

func (m *IcpeOperSatRemoteVersionType) GetVersionCheckState() string {
	if m != nil {
		return m.VersionCheckState
	}
	return ""
}

func (m *IcpeOperSatRemoteVersionType) GetRemoteVersionPresent() bool {
	if m != nil {
		return m.RemoteVersionPresent
	}
	return false
}

func (m *IcpeOperSatRemoteVersionType) GetRemoteVersion() []string {
	if m != nil {
		return m.RemoteVersion
	}
	return nil
}

func (m *IcpeOperSatRemoteVersionType) GetActiveVersion() *IcpeOperSatVersion {
	if m != nil {
		return m.ActiveVersion
	}
	return nil
}

func (m *IcpeOperSatRemoteVersionType) GetTransferredVersion() *IcpeOperSatVersion {
	if m != nil {
		return m.TransferredVersion
	}
	return nil
}

func (m *IcpeOperSatRemoteVersionType) GetCommittedVersion() *IcpeOperSatVersion {
	if m != nil {
		return m.CommittedVersion
	}
	return nil
}

func init() {
	proto.RegisterType((*IcpeOperSatRemoteVersionType_KEYS)(nil), "cisco_ios_xr_icpe_infra_oper.nv_satellite.satellite_versions.satellite_version.icpe_oper_sat_remote_version_type_KEYS")
	proto.RegisterType((*IcpeOperSatVersion)(nil), "cisco_ios_xr_icpe_infra_oper.nv_satellite.satellite_versions.satellite_version.icpe_oper_sat_version")
	proto.RegisterType((*IcpeOperSatRemoteVersionType)(nil), "cisco_ios_xr_icpe_infra_oper.nv_satellite.satellite_versions.satellite_version.icpe_oper_sat_remote_version_type")
}

func init() {
	proto.RegisterFile("icpe_oper_sat_remote_version_type.proto", fileDescriptor_92a488b685ba3cbf)
}

var fileDescriptor_92a488b685ba3cbf = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x59, 0x8b, 0x62, 0xb7, 0xa6, 0xda, 0xad, 0x4a, 0x8e, 0x69, 0xc1, 0x9a, 0x53, 0x0e,
	0x6d, 0xfd, 0x73, 0x17, 0x0f, 0x52, 0x10, 0x49, 0x41, 0xf4, 0xb4, 0xc4, 0xed, 0x14, 0x17, 0xdb,
	0x6c, 0xd8, 0x1d, 0x4a, 0xfd, 0x0e, 0x9e, 0x04, 0xfd, 0x04, 0x7e, 0x50, 0x69, 0xd2, 0xa4, 0x89,
	0x56, 0x4a, 0x41, 0x7a, 0x0b, 0xf3, 0xe6, 0xbd, 0xf9, 0x3d, 0x92, 0xd0, 0x53, 0x29, 0x22, 0xe0,
	0x2a, 0x02, 0xcd, 0x4d, 0x80, 0x5c, 0xc3, 0x58, 0x21, 0xf0, 0x09, 0x68, 0x23, 0x55, 0xc8, 0xf1,
	0x35, 0x02, 0x2f, 0xd2, 0x0a, 0x15, 0xbb, 0x15, 0xd2, 0x08, 0xc5, 0xa5, 0x32, 0x7c, 0xaa, 0x79,
	0xec, 0x92, 0xe1, 0x50, 0x07, 0xb1, 0xd7, 0x0b, 0x27, 0x33, 0x3b, 0x8c, 0x46, 0x12, 0xc1, 0xcb,
	0x9e, 0xd2, 0x1c, 0xf3, 0x7b, 0xd4, 0xec, 0xd1, 0xd6, 0xca, 0xd3, 0xbc, 0x77, 0xfd, 0xd8, 0x67,
	0x0d, 0xba, 0xb7, 0xb0, 0xcb, 0x81, 0x4d, 0x1c, 0xe2, 0x5a, 0x7e, 0x25, 0x9b, 0xdd, 0x0c, 0x9a,
	0x5f, 0x84, 0x1e, 0x15, 0xd3, 0xe6, 0x31, 0xcc, 0xa3, 0xf5, 0x34, 0x51, 0x3c, 0x83, 0x78, 0xe1,
	0x06, 0x03, 0x84, 0x38, 0xa3, 0xec, 0xd7, 0xe6, 0xd2, 0xd5, 0x4c, 0xe9, 0xcf, 0x04, 0xd6, 0xa5,
	0xc7, 0x3f, 0x40, 0x22, 0x0d, 0x06, 0x42, 0xb4, 0xb7, 0x1c, 0xe2, 0xee, 0xfa, 0x87, 0x89, 0x7a,
	0x9f, 0x88, 0x77, 0x89, 0xc6, 0x4e, 0x68, 0xb5, 0xe8, 0xb2, 0x4b, 0x4e, 0xc9, 0x2d, 0xfb, 0x56,
	0x61, 0xbb, 0xf9, 0xb1, 0x4d, 0x1b, 0x2b, 0x4b, 0xb3, 0x16, 0xdd, 0xcf, 0xf7, 0xe5, 0x53, 0x6d,
	0xb7, 0xe3, 0xca, 0x56, 0xae, 0xf2, 0x83, 0xfe, 0xab, 0x5a, 0x67, 0xfd, 0x6a, 0xdd, 0xb5, 0xaa,
	0x9d, 0x2d, 0xa9, 0xc6, 0xde, 0x08, 0xad, 0x06, 0x02, 0xe5, 0x64, 0xb1, 0x77, 0xee, 0x10, 0xb7,
	0xd2, 0x06, 0xef, 0x7f, 0x3f, 0x1c, 0x6f, 0xe9, 0x7b, 0xf6, 0xad, 0xe4, 0x78, 0x8a, 0xf3, 0x49,
	0x68, 0x1d, 0x75, 0x10, 0x9a, 0x21, 0x68, 0x0d, 0x83, 0x8c, 0xe9, 0x62, 0x93, 0x4c, 0x2c, 0x47,
	0x90, 0x82, 0xbd, 0x13, 0x5a, 0x13, 0x6a, 0x3c, 0x96, 0x88, 0x39, 0xac, 0xcb, 0x4d, 0x62, 0x1d,
	0x64, 0xf7, 0xe7, 0x50, 0x4f, 0x3b, 0xf1, 0x2f, 0xde, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xcd,
	0x86, 0x6a, 0x44, 0x0d, 0x04, 0x00, 0x00,
}