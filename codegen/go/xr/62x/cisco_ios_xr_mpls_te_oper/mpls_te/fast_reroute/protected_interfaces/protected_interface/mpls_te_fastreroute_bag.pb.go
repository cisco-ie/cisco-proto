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
// source: mpls_te_fastreroute_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_fast_reroute_protected_interfaces_protected_interface

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

type MplsTeFastrerouteBag_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	BackupTunnelName     string   `protobuf:"bytes,2,opt,name=backup_tunnel_name,json=backupTunnelName,proto3" json:"backup_tunnel_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeFastrerouteBag_KEYS) Reset()         { *m = MplsTeFastrerouteBag_KEYS{} }
func (m *MplsTeFastrerouteBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeFastrerouteBag_KEYS) ProtoMessage()    {}
func (*MplsTeFastrerouteBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a407a6e5f15fcedf, []int{0}
}

func (m *MplsTeFastrerouteBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeFastrerouteBag_KEYS.Unmarshal(m, b)
}
func (m *MplsTeFastrerouteBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeFastrerouteBag_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeFastrerouteBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeFastrerouteBag_KEYS.Merge(m, src)
}
func (m *MplsTeFastrerouteBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeFastrerouteBag_KEYS.Size(m)
}
func (m *MplsTeFastrerouteBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeFastrerouteBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeFastrerouteBag_KEYS proto.InternalMessageInfo

func (m *MplsTeFastrerouteBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *MplsTeFastrerouteBag_KEYS) GetBackupTunnelName() string {
	if m != nil {
		return m.BackupTunnelName
	}
	return ""
}

type MplsTeFastrerouteBag struct {
	BackupTunnelId                uint32   `protobuf:"varint,50,opt,name=backup_tunnel_id,json=backupTunnelId,proto3" json:"backup_tunnel_id,omitempty"`
	BackupTunnelNameXr            string   `protobuf:"bytes,51,opt,name=backup_tunnel_name_xr,json=backupTunnelNameXr,proto3" json:"backup_tunnel_name_xr,omitempty"`
	BackupStatus                  string   `protobuf:"bytes,52,opt,name=backup_status,json=backupStatus,proto3" json:"backup_status,omitempty"`
	BackupType                    string   `protobuf:"bytes,53,opt,name=backup_type,json=backupType,proto3" json:"backup_type,omitempty"`
	BackupUsage                   string   `protobuf:"bytes,54,opt,name=backup_usage,json=backupUsage,proto3" json:"backup_usage,omitempty"`
	ProtInterfaceAutobackupConfig string   `protobuf:"bytes,55,opt,name=prot_interface_autobackup_config,json=protInterfaceAutobackupConfig,proto3" json:"prot_interface_autobackup_config,omitempty"`
	ProtInterfaceSrlgConfig       string   `protobuf:"bytes,56,opt,name=prot_interface_srlg_config,json=protInterfaceSrlgConfig,proto3" json:"prot_interface_srlg_config,omitempty"`
	TunnelAttributeSetName        string   `protobuf:"bytes,57,opt,name=tunnel_attribute_set_name,json=tunnelAttributeSetName,proto3" json:"tunnel_attribute_set_name,omitempty"`
	HasAttributeSet               bool     `protobuf:"varint,58,opt,name=has_attribute_set,json=hasAttributeSet,proto3" json:"has_attribute_set,omitempty"`
	IsAttributeSetInDb            bool     `protobuf:"varint,59,opt,name=is_attribute_set_in_db,json=isAttributeSetInDb,proto3" json:"is_attribute_set_in_db,omitempty"`
	RecreateTimerIsRunning        bool     `protobuf:"varint,60,opt,name=recreate_timer_is_running,json=recreateTimerIsRunning,proto3" json:"recreate_timer_is_running,omitempty"`
	RecreateRemainingTime         uint32   `protobuf:"varint,61,opt,name=recreate_remaining_time,json=recreateRemainingTime,proto3" json:"recreate_remaining_time,omitempty"`
	SourceAddress                 string   `protobuf:"bytes,62,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress            string   `protobuf:"bytes,63,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	GeneralStatus                 string   `protobuf:"bytes,64,opt,name=general_status,json=generalStatus,proto3" json:"general_status,omitempty"`
	ConnectionStatus              string   `protobuf:"bytes,65,opt,name=connection_status,json=connectionStatus,proto3" json:"connection_status,omitempty"`
	OutputInterfaceName           string   `protobuf:"bytes,66,opt,name=output_interface_name,json=outputInterfaceName,proto3" json:"output_interface_name,omitempty"`
	BandwidthPoolType             string   `protobuf:"bytes,67,opt,name=bandwidth_pool_type,json=bandwidthPoolType,proto3" json:"bandwidth_pool_type,omitempty"`
	BandwidthLimitType            string   `protobuf:"bytes,68,opt,name=bandwidth_limit_type,json=bandwidthLimitType,proto3" json:"bandwidth_limit_type,omitempty"`
	Bandwidth                     uint32   `protobuf:"varint,69,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	TunnelInstance                uint32   `protobuf:"varint,70,opt,name=tunnel_instance,json=tunnelInstance,proto3" json:"tunnel_instance,omitempty"`
	InUseBandwidth                uint32   `protobuf:"varint,71,opt,name=in_use_bandwidth,json=inUseBandwidth,proto3" json:"in_use_bandwidth,omitempty"`
	SoftPreemptedInUseBandwidth   uint32   `protobuf:"varint,72,opt,name=soft_preempted_in_use_bandwidth,json=softPreemptedInUseBandwidth,proto3" json:"soft_preempted_in_use_bandwidth,omitempty"`
	LsPs                          uint32   `protobuf:"varint,73,opt,name=ls_ps,json=lsPs,proto3" json:"ls_ps,omitempty"`
	S2LFamilies                   uint32   `protobuf:"varint,74,opt,name=s2l_families,json=s2lFamilies,proto3" json:"s2l_families,omitempty"`
	S2Ls                          uint32   `protobuf:"varint,75,opt,name=s2_ls,json=s2Ls,proto3" json:"s2_ls,omitempty"`
	FrrActiveLsPs                 uint32   `protobuf:"varint,76,opt,name=frr_active_ls_ps,json=frrActiveLsPs,proto3" json:"frr_active_ls_ps,omitempty"`
	FrrActiveSoftPreemptedLsPs    uint32   `protobuf:"varint,77,opt,name=frr_active_soft_preempted_ls_ps,json=frrActiveSoftPreemptedLsPs,proto3" json:"frr_active_soft_preempted_ls_ps,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *MplsTeFastrerouteBag) Reset()         { *m = MplsTeFastrerouteBag{} }
func (m *MplsTeFastrerouteBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeFastrerouteBag) ProtoMessage()    {}
func (*MplsTeFastrerouteBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_a407a6e5f15fcedf, []int{1}
}

func (m *MplsTeFastrerouteBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeFastrerouteBag.Unmarshal(m, b)
}
func (m *MplsTeFastrerouteBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeFastrerouteBag.Marshal(b, m, deterministic)
}
func (m *MplsTeFastrerouteBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeFastrerouteBag.Merge(m, src)
}
func (m *MplsTeFastrerouteBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeFastrerouteBag.Size(m)
}
func (m *MplsTeFastrerouteBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeFastrerouteBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeFastrerouteBag proto.InternalMessageInfo

func (m *MplsTeFastrerouteBag) GetBackupTunnelId() uint32 {
	if m != nil {
		return m.BackupTunnelId
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetBackupTunnelNameXr() string {
	if m != nil {
		return m.BackupTunnelNameXr
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBackupStatus() string {
	if m != nil {
		return m.BackupStatus
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBackupType() string {
	if m != nil {
		return m.BackupType
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBackupUsage() string {
	if m != nil {
		return m.BackupUsage
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetProtInterfaceAutobackupConfig() string {
	if m != nil {
		return m.ProtInterfaceAutobackupConfig
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetProtInterfaceSrlgConfig() string {
	if m != nil {
		return m.ProtInterfaceSrlgConfig
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetTunnelAttributeSetName() string {
	if m != nil {
		return m.TunnelAttributeSetName
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetHasAttributeSet() bool {
	if m != nil {
		return m.HasAttributeSet
	}
	return false
}

func (m *MplsTeFastrerouteBag) GetIsAttributeSetInDb() bool {
	if m != nil {
		return m.IsAttributeSetInDb
	}
	return false
}

func (m *MplsTeFastrerouteBag) GetRecreateTimerIsRunning() bool {
	if m != nil {
		return m.RecreateTimerIsRunning
	}
	return false
}

func (m *MplsTeFastrerouteBag) GetRecreateRemainingTime() uint32 {
	if m != nil {
		return m.RecreateRemainingTime
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetGeneralStatus() string {
	if m != nil {
		return m.GeneralStatus
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetConnectionStatus() string {
	if m != nil {
		return m.ConnectionStatus
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetOutputInterfaceName() string {
	if m != nil {
		return m.OutputInterfaceName
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBandwidthPoolType() string {
	if m != nil {
		return m.BandwidthPoolType
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBandwidthLimitType() string {
	if m != nil {
		return m.BandwidthLimitType
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetTunnelInstance() uint32 {
	if m != nil {
		return m.TunnelInstance
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetInUseBandwidth() uint32 {
	if m != nil {
		return m.InUseBandwidth
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetSoftPreemptedInUseBandwidth() uint32 {
	if m != nil {
		return m.SoftPreemptedInUseBandwidth
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetLsPs() uint32 {
	if m != nil {
		return m.LsPs
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetS2LFamilies() uint32 {
	if m != nil {
		return m.S2LFamilies
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetS2Ls() uint32 {
	if m != nil {
		return m.S2Ls
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetFrrActiveLsPs() uint32 {
	if m != nil {
		return m.FrrActiveLsPs
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetFrrActiveSoftPreemptedLsPs() uint32 {
	if m != nil {
		return m.FrrActiveSoftPreemptedLsPs
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeFastrerouteBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.fast_reroute.protected_interfaces.protected_interface.mpls_te_fastreroute_bag_KEYS")
	proto.RegisterType((*MplsTeFastrerouteBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.fast_reroute.protected_interfaces.protected_interface.mpls_te_fastreroute_bag")
}

func init() { proto.RegisterFile("mpls_te_fastreroute_bag.proto", fileDescriptor_a407a6e5f15fcedf) }

var fileDescriptor_a407a6e5f15fcedf = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x7b, 0x6f, 0x23, 0x35,
	0x14, 0xc5, 0x15, 0x04, 0x88, 0xf5, 0x92, 0x6e, 0xeb, 0xd0, 0xd6, 0x2c, 0xbb, 0x6a, 0x58, 0xb4,
	0xda, 0x08, 0x50, 0x80, 0x14, 0x0a, 0xa5, 0xbc, 0xd2, 0x27, 0x43, 0x03, 0xaa, 0x92, 0x56, 0xc0,
	0x5f, 0x96, 0x33, 0x73, 0x93, 0x5a, 0xcc, 0xd8, 0x23, 0x5f, 0x0f, 0xb4, 0x9f, 0x88, 0xaf, 0x89,
	0xfc, 0x98, 0xc9, 0x83, 0xed, 0x9f, 0x3d, 0xe7, 0xfc, 0x8e, 0x1b, 0xfb, 0xde, 0x21, 0xcf, 0x8b,
	0x32, 0x47, 0x6e, 0x81, 0xcf, 0x04, 0x5a, 0x03, 0x46, 0x57, 0x16, 0xf8, 0x54, 0xcc, 0xfb, 0xa5,
	0xd1, 0x56, 0xd3, 0xdf, 0x53, 0x89, 0xa9, 0xe6, 0x52, 0x23, 0xbf, 0x33, 0xbc, 0xce, 0xea, 0x12,
	0x4c, 0x3f, 0xfe, 0xd1, 0x77, 0x20, 0x8f, 0xa4, 0xa7, 0x20, 0xb5, 0x90, 0x71, 0xa9, 0x2c, 0x98,
	0x99, 0x48, 0x01, 0x5f, 0x27, 0xbe, 0x40, 0xf2, 0xec, 0x81, 0x93, 0xf9, 0xe5, 0xd9, 0x9f, 0x13,
	0xfa, 0x92, 0x6c, 0x34, 0x61, 0xae, 0x44, 0x01, 0xac, 0xd5, 0x6d, 0xf5, 0x1e, 0x8d, 0xdb, 0x8d,
	0xfa, 0x9b, 0x28, 0x80, 0x7e, 0x4a, 0xe8, 0x54, 0xa4, 0x7f, 0x55, 0x25, 0xb7, 0x95, 0x52, 0x90,
	0x87, 0xe8, 0x1b, 0x3e, 0xba, 0x19, 0x9c, 0x6b, 0x6f, 0xb8, 0xf4, 0x8b, 0x7f, 0x09, 0xd9, 0x7d,
	0xe0, 0x54, 0xda, 0x23, 0x9b, 0xab, 0x4d, 0x32, 0x63, 0x83, 0x6e, 0xab, 0xd7, 0x1e, 0x6f, 0x2c,
	0xf7, 0x24, 0x19, 0xfd, 0x82, 0x6c, 0xff, 0xff, 0x4c, 0x7e, 0x67, 0xd8, 0xbe, 0x3f, 0x96, 0xae,
	0x1f, 0xfb, 0x87, 0xa1, 0x1f, 0x91, 0x76, 0x44, 0xd0, 0x0a, 0x5b, 0x21, 0xfb, 0xd2, 0x47, 0xdf,
	0x0d, 0xe2, 0xc4, 0x6b, 0x74, 0x8f, 0x3c, 0xae, 0x7b, 0xef, 0x4b, 0x60, 0x5f, 0xf9, 0x08, 0x89,
	0x6d, 0xf7, 0x25, 0xd0, 0x0f, 0x49, 0x04, 0x78, 0x85, 0x62, 0x0e, 0xec, 0xc0, 0x27, 0x22, 0x74,
	0xe3, 0x24, 0x7a, 0x41, 0xba, 0xee, 0xb6, 0x17, 0x17, 0xcd, 0x45, 0x65, 0x75, 0xa4, 0x52, 0xad,
	0x66, 0x72, 0xce, 0xbe, 0xf6, 0xd8, 0x73, 0x97, 0x4b, 0xea, 0xd8, 0xb0, 0x49, 0x9d, 0xf8, 0x10,
	0x3d, 0x22, 0x4f, 0xd7, 0x8a, 0xd0, 0xe4, 0xf3, 0xba, 0xe2, 0x1b, 0x5f, 0xb1, 0xbb, 0x52, 0x31,
	0x31, 0xf9, 0x3c, 0xc2, 0x87, 0xe4, 0xfd, 0x78, 0x35, 0xc2, 0x5a, 0x23, 0xa7, 0xee, 0x8e, 0x11,
	0x6c, 0x78, 0x9c, 0x43, 0xcf, 0xee, 0x84, 0xc0, 0xb0, 0xf6, 0x27, 0x60, 0xfd, 0x83, 0x7e, 0x4c,
	0xb6, 0x6e, 0x05, 0xae, 0x72, 0xec, 0xdb, 0x6e, 0xab, 0xf7, 0xce, 0xf8, 0xc9, 0xad, 0xc0, 0xe5,
	0x3c, 0x1d, 0x90, 0x1d, 0xb9, 0x16, 0xe5, 0x52, 0xf1, 0x6c, 0xca, 0x8e, 0x3c, 0x40, 0xe5, 0x4a,
	0x3e, 0x51, 0xa7, 0x53, 0xf7, 0xaf, 0x19, 0x48, 0x0d, 0x08, 0x0b, 0xdc, 0xca, 0x02, 0x0c, 0x97,
	0xc8, 0x4d, 0xa5, 0x94, 0x54, 0x73, 0xf6, 0x9d, 0xc7, 0x76, 0xea, 0xc0, 0xb5, 0xf3, 0x13, 0x1c,
	0x07, 0x97, 0x1e, 0x90, 0xdd, 0x06, 0x35, 0x50, 0x08, 0xe9, 0x54, 0x5f, 0xc2, 0xbe, 0xf7, 0x83,
	0xb2, 0x5d, 0xdb, 0xe3, 0xda, 0x75, 0x0d, 0x6e, 0x94, 0x51, 0x57, 0xc6, 0xbd, 0x45, 0x96, 0x19,
	0x40, 0x64, 0x3f, 0x84, 0x51, 0x0e, 0xea, 0x30, 0x88, 0xf4, 0x33, 0xd2, 0xc9, 0x00, 0xad, 0x54,
	0xc2, 0x4a, 0xad, 0x9a, 0xec, 0x8f, 0x61, 0xa8, 0x96, 0xac, 0x1a, 0x78, 0x49, 0x36, 0xe6, 0xa0,
	0xc0, 0x88, 0xbc, 0x9e, 0xaa, 0x9f, 0x42, 0x6f, 0x54, 0xe3, 0x58, 0x7d, 0x42, 0xb6, 0x52, 0xad,
	0x14, 0xa4, 0xbe, 0x36, 0x26, 0x87, 0x61, 0x43, 0x16, 0x46, 0x0c, 0x0f, 0xc8, 0xb6, 0xae, 0x6c,
	0x59, 0x2d, 0x3f, 0xbc, 0x7f, 0xb5, 0x63, 0x0f, 0x74, 0x82, 0x99, 0xac, 0xec, 0x60, 0x9f, 0x74,
	0xa6, 0x42, 0x65, 0xff, 0xc8, 0xcc, 0xde, 0xf2, 0x52, 0xeb, 0x3c, 0xcc, 0xef, 0x89, 0x27, 0xb6,
	0x1a, 0xeb, 0x4a, 0xeb, 0xdc, 0x8f, 0xf1, 0xe7, 0xe4, 0xbd, 0x45, 0x3e, 0x97, 0x85, 0xb4, 0x01,
	0x38, 0xad, 0xd7, 0x27, 0x7a, 0x23, 0x67, 0x79, 0xe2, 0x19, 0x79, 0xd4, 0xa8, 0xec, 0xcc, 0xdf,
	0xf5, 0x42, 0xa0, 0xaf, 0xc8, 0x93, 0x7a, 0x65, 0x15, 0x5a, 0xa1, 0x52, 0x60, 0xe7, 0x61, 0x71,
	0x83, 0x9c, 0x44, 0xd5, 0xad, 0xb8, 0x54, 0xbc, 0x42, 0xb7, 0xf0, 0x75, 0xdb, 0x45, 0x48, 0x4a,
	0x75, 0x83, 0x70, 0xdc, 0x54, 0x9e, 0x92, 0x3d, 0xd4, 0x33, 0xcb, 0x4b, 0x03, 0x50, 0x94, 0xe1,
	0xcb, 0xb5, 0x06, 0xfe, 0xec, 0xc1, 0x0f, 0x5c, 0xec, 0xaa, 0x4e, 0x25, 0xab, 0x2d, 0x1d, 0xf2,
	0x56, 0x8e, 0xbc, 0x44, 0x96, 0xf8, 0xec, 0x9b, 0x39, 0x5e, 0xa1, 0x5b, 0x62, 0x1c, 0xe4, 0x7c,
	0x26, 0x0a, 0x99, 0x4b, 0x40, 0xf6, 0x8b, 0xf7, 0x1e, 0xe3, 0x20, 0x3f, 0x8f, 0x92, 0xe3, 0x70,
	0xc0, 0x73, 0x64, 0x97, 0x81, 0xc3, 0xc1, 0x08, 0xe9, 0x2b, 0xb2, 0x39, 0x33, 0x86, 0x8b, 0xd4,
	0xca, 0xbf, 0x81, 0x87, 0xde, 0x91, 0xf7, 0xdb, 0x33, 0x63, 0x86, 0x5e, 0x1e, 0xb9, 0x03, 0x4e,
	0xc8, 0xde, 0x52, 0x70, 0xed, 0x67, 0x04, 0xee, 0x57, 0xcf, 0x3d, 0x6d, 0xb8, 0xc9, 0xf2, 0x8f,
	0x70, 0x25, 0xd3, 0xb7, 0xfd, 0xe7, 0x7f, 0xff, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xab,
	0x39, 0x46, 0x1f, 0x06, 0x00, 0x00,
}
