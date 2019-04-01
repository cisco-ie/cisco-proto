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

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_fast_reroute_protected_interfaces_protected_interface

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
	proto.RegisterType((*MplsTeFastrerouteBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fast_reroute.protected_interfaces.protected_interface.mpls_te_fastreroute_bag_KEYS")
	proto.RegisterType((*MplsTeFastrerouteBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fast_reroute.protected_interfaces.protected_interface.mpls_te_fastreroute_bag")
}

func init() { proto.RegisterFile("mpls_te_fastreroute_bag.proto", fileDescriptor_a407a6e5f15fcedf) }

var fileDescriptor_a407a6e5f15fcedf = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0xeb, 0x6e, 0x2b, 0x35,
	0x14, 0x85, 0x15, 0x04, 0x88, 0xe3, 0x43, 0x7a, 0x5a, 0x87, 0xb6, 0xe6, 0xd0, 0xaa, 0xa1, 0xa8,
	0x6a, 0x04, 0x28, 0x40, 0x0a, 0x85, 0x52, 0x6e, 0xe9, 0x95, 0xa1, 0x01, 0x55, 0x49, 0x2b, 0xc1,
	0x2f, 0xcb, 0x33, 0xb3, 0x93, 0x5a, 0xcc, 0xd8, 0x23, 0x6f, 0x0f, 0xb4, 0x4f, 0xc4, 0x6b, 0x22,
	0xdb, 0x33, 0x93, 0x0b, 0xf4, 0x67, 0xd7, 0xfa, 0xd6, 0x72, 0x62, 0xef, 0x9d, 0x92, 0xdd, 0xbc,
	0xc8, 0x90, 0x5b, 0xe0, 0x53, 0x81, 0xd6, 0x80, 0xd1, 0xa5, 0x05, 0x1e, 0x8b, 0x59, 0xbf, 0x30,
	0xda, 0x6a, 0xca, 0x13, 0x89, 0x89, 0xe6, 0x52, 0x23, 0x7f, 0x34, 0xbc, 0x66, 0x75, 0x01, 0xa6,
	0x5f, 0xff, 0x81, 0x56, 0xa8, 0x34, 0x7e, 0xea, 0xbb, 0x02, 0x5e, 0x35, 0xf8, 0x34, 0x24, 0x16,
	0x52, 0x2e, 0x95, 0x05, 0x33, 0x15, 0x09, 0xe0, 0xff, 0x89, 0xfb, 0x48, 0x76, 0x9e, 0xf9, 0x04,
	0xfc, 0xe6, 0xf2, 0x8f, 0x09, 0x3d, 0x20, 0x6b, 0x0d, 0xcc, 0x95, 0xc8, 0x81, 0xb5, 0xba, 0xad,
	0xde, 0x8b, 0x71, 0xbb, 0x51, 0x7f, 0x13, 0x39, 0xd0, 0x4f, 0x09, 0x8d, 0x45, 0xf2, 0x67, 0x59,
	0x70, 0x5b, 0x2a, 0x05, 0x59, 0x40, 0xdf, 0xf0, 0xe8, 0x7a, 0x70, 0xee, 0xbc, 0xe1, 0xe8, 0xfd,
	0x7f, 0x08, 0xd9, 0x7e, 0xe6, 0x54, 0xda, 0x23, 0xeb, 0xcb, 0x4d, 0x32, 0x65, 0x83, 0x6e, 0xab,
	0xd7, 0x1e, 0xaf, 0x2d, 0xf6, 0x44, 0x29, 0xfd, 0x82, 0x6c, 0xfe, 0xf7, 0x4c, 0xfe, 0x68, 0xd8,
	0x91, 0x3f, 0x96, 0xae, 0x1e, 0xfb, 0xbb, 0xa1, 0x1f, 0x91, 0x76, 0x15, 0x41, 0x2b, 0x6c, 0x89,
	0xec, 0x4b, 0x8f, 0xbe, 0x1b, 0xc4, 0x89, 0xd7, 0xe8, 0x1e, 0x79, 0x59, 0xf7, 0x3e, 0x15, 0xc0,
	0xbe, 0xf2, 0x08, 0xa9, 0xda, 0x9e, 0x0a, 0xa0, 0x1f, 0x92, 0x2a, 0xc0, 0x4b, 0x14, 0x33, 0x60,
	0xc7, 0x9e, 0xa8, 0x42, 0xf7, 0x4e, 0xa2, 0xd7, 0xa4, 0xeb, 0x6e, 0x7b, 0x7e, 0xd1, 0x5c, 0x94,
	0x56, 0x57, 0xa9, 0x44, 0xab, 0xa9, 0x9c, 0xb1, 0xaf, 0x7d, 0x6c, 0xd7, 0x71, 0x51, 0x8d, 0x0d,
	0x1b, 0xea, 0xdc, 0x43, 0xf4, 0x94, 0xbc, 0x5e, 0x29, 0x42, 0x93, 0xcd, 0xea, 0x8a, 0x6f, 0x7c,
	0xc5, 0xf6, 0x52, 0xc5, 0xc4, 0x64, 0xb3, 0x2a, 0x7c, 0x42, 0xde, 0xaf, 0xae, 0x46, 0x58, 0x6b,
	0x64, 0xec, 0xee, 0x18, 0xc1, 0x86, 0xc7, 0x39, 0xf1, 0xd9, 0xad, 0x00, 0x0c, 0x6b, 0x7f, 0x02,
	0xd6, 0x3f, 0xe8, 0xc7, 0x64, 0xe3, 0x41, 0xe0, 0x72, 0x8e, 0x7d, 0xdb, 0x6d, 0xf5, 0xde, 0x19,
	0xbf, 0x7a, 0x10, 0xb8, 0xc8, 0xd3, 0x01, 0xd9, 0x92, 0x2b, 0x28, 0x97, 0x8a, 0xa7, 0x31, 0x3b,
	0xf5, 0x01, 0x2a, 0x97, 0xf8, 0x48, 0x5d, 0xc4, 0xee, 0xa3, 0x19, 0x48, 0x0c, 0x08, 0x0b, 0xdc,
	0xca, 0x1c, 0x0c, 0x97, 0xc8, 0x4d, 0xa9, 0x94, 0x54, 0x33, 0xf6, 0x9d, 0x8f, 0x6d, 0xd5, 0xc0,
	0x9d, 0xf3, 0x23, 0x1c, 0x07, 0x97, 0x1e, 0x93, 0xed, 0x26, 0x6a, 0x20, 0x17, 0xd2, 0xa9, 0xbe,
	0x84, 0x7d, 0xef, 0x07, 0x65, 0xb3, 0xb6, 0xc7, 0xb5, 0xeb, 0x1a, 0xdc, 0x28, 0xa3, 0x2e, 0x8d,
	0x7b, 0x8b, 0x34, 0x35, 0x80, 0xc8, 0x7e, 0x08, 0xa3, 0x1c, 0xd4, 0x61, 0x10, 0xe9, 0x67, 0xa4,
	0x93, 0x02, 0x5a, 0xa9, 0x84, 0x95, 0x5a, 0x35, 0xec, 0x8f, 0x61, 0xa8, 0x16, 0xac, 0x3a, 0x70,
	0x40, 0xd6, 0x66, 0xa0, 0xc0, 0x88, 0xac, 0x9e, 0xaa, 0x9f, 0x42, 0x6f, 0xa5, 0x56, 0x63, 0xf5,
	0x09, 0xd9, 0x48, 0xb4, 0x52, 0x90, 0xf8, 0xda, 0x8a, 0x1c, 0x86, 0x0d, 0x99, 0x1b, 0x15, 0x3c,
	0x20, 0x9b, 0xba, 0xb4, 0x45, 0xb9, 0xf8, 0xf0, 0xfe, 0xd5, 0xce, 0x7c, 0xa0, 0x13, 0xcc, 0x68,
	0x69, 0x07, 0xfb, 0xa4, 0x13, 0x0b, 0x95, 0xfe, 0x2d, 0x53, 0xfb, 0xc0, 0x0b, 0xad, 0xb3, 0x30,
	0xbf, 0xe7, 0x3e, 0xb1, 0xd1, 0x58, 0xb7, 0x5a, 0x67, 0x7e, 0x8c, 0x3f, 0x27, 0xef, 0xcd, 0xf9,
	0x4c, 0xe6, 0xd2, 0x86, 0xc0, 0x45, 0xbd, 0x3e, 0x95, 0x37, 0x72, 0x96, 0x4f, 0xec, 0x90, 0x17,
	0x8d, 0xca, 0x2e, 0xfd, 0x5d, 0xcf, 0x05, 0x7a, 0x48, 0x5e, 0xd5, 0x2b, 0xab, 0xdc, 0x8f, 0x52,
	0x02, 0xec, 0x2a, 0x2c, 0x6e, 0x90, 0xa3, 0x4a, 0x75, 0x2b, 0x2e, 0x15, 0x2f, 0xd1, 0x2d, 0x7c,
	0xdd, 0x76, 0x1d, 0x48, 0xa9, 0xee, 0x11, 0xce, 0x9a, 0xca, 0x0b, 0xb2, 0x87, 0x7a, 0x6a, 0x79,
	0x61, 0x00, 0xf2, 0x22, 0xfc, 0x72, 0xad, 0x04, 0x7f, 0xf6, 0xc1, 0x0f, 0x1c, 0x76, 0x5b, 0x53,
	0xd1, 0x72, 0x4b, 0x87, 0xbc, 0x95, 0x21, 0x2f, 0x90, 0x45, 0x9e, 0x7d, 0x33, 0xc3, 0x5b, 0x74,
	0x4b, 0x8c, 0x83, 0x8c, 0x4f, 0x45, 0x2e, 0x33, 0x09, 0xc8, 0x7e, 0xf1, 0xde, 0x4b, 0x1c, 0x64,
	0x57, 0x95, 0xe4, 0x72, 0x38, 0xe0, 0x19, 0xb2, 0x9b, 0x90, 0xc3, 0xc1, 0x08, 0xe9, 0x21, 0x59,
	0x9f, 0x1a, 0xc3, 0x45, 0x62, 0xe5, 0x5f, 0xc0, 0x43, 0xef, 0xc8, 0xfb, 0xed, 0xa9, 0x31, 0x43,
	0x2f, 0x8f, 0xdc, 0x01, 0xe7, 0x64, 0x6f, 0x01, 0x5c, 0xf9, 0x1a, 0x21, 0xf7, 0xab, 0xcf, 0xbd,
	0x6e, 0x72, 0x93, 0xc5, 0x2f, 0xe1, 0x4a, 0xe2, 0xb7, 0xfd, 0xbf, 0x81, 0xa3, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xcf, 0xfd, 0xa7, 0x5e, 0x27, 0x06, 0x00, 0x00,
}
