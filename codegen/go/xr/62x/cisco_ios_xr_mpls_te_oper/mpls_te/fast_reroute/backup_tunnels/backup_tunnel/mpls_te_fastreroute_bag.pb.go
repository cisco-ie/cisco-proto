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

package cisco_ios_xr_mpls_te_oper_mpls_te_fast_reroute_backup_tunnels_backup_tunnel

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
	TunnelName           string   `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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

func (m *MplsTeFastrerouteBag_KEYS) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *MplsTeFastrerouteBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
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
	proto.RegisterType((*MplsTeFastrerouteBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.fast_reroute.backup_tunnels.backup_tunnel.mpls_te_fastreroute_bag_KEYS")
	proto.RegisterType((*MplsTeFastrerouteBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.fast_reroute.backup_tunnels.backup_tunnel.mpls_te_fastreroute_bag")
}

func init() { proto.RegisterFile("mpls_te_fastreroute_bag.proto", fileDescriptor_a407a6e5f15fcedf) }

var fileDescriptor_a407a6e5f15fcedf = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xfd, 0x4f, 0x1b, 0x37,
	0x1c, 0xc6, 0x95, 0x69, 0x9b, 0x86, 0x59, 0x78, 0x71, 0x06, 0x78, 0x0c, 0x44, 0xc6, 0x84, 0x88,
	0x36, 0x29, 0xdb, 0xc2, 0xc6, 0xc6, 0xd8, 0xd6, 0x86, 0xd7, 0x5e, 0x49, 0x2b, 0x94, 0x80, 0xd4,
	0xfe, 0x64, 0x39, 0x77, 0xbe, 0x60, 0xf5, 0xce, 0x3e, 0xf9, 0xeb, 0x6b, 0xe1, 0x2f, 0xea, 0xbf,
	0x59, 0xf9, 0xe5, 0x2e, 0x09, 0x15, 0x3f, 0xe6, 0x79, 0x3e, 0xcf, 0xe3, 0xb3, 0xfd, 0x75, 0xd0,
	0x76, 0x5e, 0x64, 0x40, 0x0d, 0xa7, 0x29, 0x03, 0xa3, 0xb9, 0x56, 0xa5, 0xe1, 0x74, 0xcc, 0x26,
	0xdd, 0x42, 0x2b, 0xa3, 0xf0, 0x55, 0x2c, 0x20, 0x56, 0x54, 0x28, 0xa0, 0xf7, 0x9a, 0x56, 0xac,
	0x2a, 0xb8, 0xee, 0x86, 0x1f, 0x5d, 0x1b, 0xa4, 0x21, 0xd9, 0x1d, 0xb3, 0xf8, 0x5d, 0x59, 0x50,
	0x53, 0x4a, 0xc9, 0x33, 0x98, 0xff, 0xb9, 0x9b, 0xa2, 0xad, 0x27, 0x56, 0xa3, 0x57, 0xe7, 0x6f,
	0x47, 0x78, 0x07, 0x2d, 0x7a, 0x92, 0x4a, 0x96, 0x73, 0xd2, 0x68, 0x37, 0x3a, 0x0b, 0x43, 0xe4,
	0xa5, 0xd7, 0x2c, 0xe7, 0x78, 0x0f, 0x2d, 0x09, 0x69, 0xb8, 0x4e, 0x59, 0xcc, 0x3d, 0xf3, 0x85,
	0x63, 0x9a, 0xb5, 0x6a, 0xb1, 0xdd, 0x8f, 0x08, 0x6d, 0x3c, 0xb1, 0x10, 0xee, 0xa0, 0x95, 0xb9,
	0x8f, 0xa2, 0x22, 0x21, 0xbd, 0x76, 0xa3, 0xd3, 0x1c, 0x2e, 0x79, 0xfd, 0xc6, 0xc9, 0x51, 0x82,
	0x7f, 0x47, 0x6b, 0xf3, 0xa4, 0x5d, 0x90, 0xde, 0x6b, 0x72, 0xe0, 0xd6, 0xc4, 0xb3, 0xb8, 0x5d,
	0xf6, 0x8d, 0xc6, 0x3f, 0xa1, 0x66, 0x88, 0x80, 0x61, 0xa6, 0x04, 0xf2, 0x87, 0x43, 0xbf, 0xf5,
	0xe2, 0xc8, 0x69, 0x76, 0x97, 0x55, 0xef, 0x43, 0xc1, 0xc9, 0x9f, 0x7e, 0x97, 0xa1, 0xed, 0xa1,
	0xe0, 0xf8, 0x47, 0x14, 0x02, 0xb4, 0x04, 0x36, 0xe1, 0xe4, 0xd0, 0x11, 0x21, 0x74, 0x6b, 0x25,
	0x7c, 0x89, 0xda, 0xf6, 0x7e, 0xe8, 0xf4, 0x34, 0x58, 0x69, 0x54, 0x48, 0xc5, 0x4a, 0xa6, 0x62,
	0x42, 0xfe, 0x72, 0xb1, 0x6d, 0xcb, 0x45, 0x15, 0xd6, 0xaf, 0xa9, 0x53, 0x07, 0xe1, 0x63, 0xb4,
	0xf9, 0xa8, 0x08, 0x74, 0x36, 0xa9, 0x2a, 0xfe, 0x76, 0x15, 0x1b, 0x73, 0x15, 0x23, 0x9d, 0x4d,
	0x42, 0xf8, 0x08, 0x7d, 0x1f, 0x8e, 0x86, 0x19, 0xa3, 0xc5, 0xd8, 0x9e, 0x31, 0x70, 0xe3, 0x6f,
	0xe6, 0xc8, 0x65, 0xd7, 0x3d, 0xd0, 0xaf, 0xfc, 0x11, 0x37, 0xee, 0x26, 0x7f, 0x46, 0xab, 0x77,
	0x0c, 0xe6, 0x73, 0xe4, 0x9f, 0x76, 0xa3, 0xf3, 0xcd, 0x70, 0xf9, 0x8e, 0xc1, 0x2c, 0x8f, 0x7b,
	0x68, 0x5d, 0x3c, 0x42, 0xa9, 0x90, 0x34, 0x19, 0x93, 0x63, 0x17, 0xc0, 0x62, 0x8e, 0x8f, 0xe4,
	0xd9, 0xd8, 0x7e, 0x9a, 0xe6, 0xb1, 0xe6, 0xcc, 0x70, 0x6a, 0x44, 0xce, 0x35, 0x15, 0x40, 0x75,
	0x29, 0xa5, 0x90, 0x13, 0xf2, 0xaf, 0x8b, 0xad, 0x57, 0xc0, 0x8d, 0xf5, 0x23, 0x18, 0x7a, 0x17,
	0x1f, 0xa2, 0x8d, 0x3a, 0xaa, 0x79, 0xce, 0x84, 0x55, 0x5d, 0x09, 0xf9, 0xcf, 0x0d, 0xca, 0x5a,
	0x65, 0x0f, 0x2b, 0xd7, 0x36, 0xd8, 0xe1, 0x04, 0x55, 0x6a, 0x7b, 0x17, 0x49, 0xa2, 0x39, 0x00,
	0xf9, 0xdf, 0x0f, 0xa7, 0x57, 0xfb, 0x5e, 0xc4, 0xbf, 0xa2, 0x56, 0xc2, 0xc1, 0x08, 0xc9, 0x8c,
	0x50, 0xb2, 0x66, 0x9f, 0xf9, 0xa1, 0x9a, 0xb1, 0xaa, 0xc0, 0x1e, 0x5a, 0x9a, 0x70, 0xc9, 0x35,
	0xcb, 0xaa, 0xa9, 0x7a, 0xee, 0x7b, 0x83, 0x1a, 0xc6, 0xea, 0x17, 0xb4, 0x1a, 0x2b, 0x29, 0x79,
	0xec, 0x6a, 0x03, 0xd9, 0x77, 0xe4, 0xca, 0xd4, 0x08, 0x70, 0x0f, 0xad, 0xa9, 0xd2, 0x14, 0xe5,
	0xec, 0xc5, 0xbb, 0x5b, 0x3b, 0x71, 0x81, 0x96, 0x37, 0xa3, 0xd9, 0x57, 0x85, 0xbb, 0xa8, 0x35,
	0x66, 0x32, 0xf9, 0x20, 0x12, 0x73, 0x47, 0x0b, 0xa5, 0x32, 0x3f, 0xbf, 0xa7, 0x2e, 0xb1, 0x5a,
	0x5b, 0xd7, 0x4a, 0x65, 0x6e, 0x8c, 0x7f, 0x43, 0xdf, 0x4d, 0xf9, 0x4c, 0xe4, 0xc2, 0xf8, 0xc0,
	0x59, 0xf5, 0x7c, 0x82, 0x37, 0xb0, 0x96, 0x4b, 0x6c, 0xa1, 0x85, 0x5a, 0x25, 0xe7, 0xee, 0xac,
	0xa7, 0x02, 0xde, 0x47, 0xcb, 0xd5, 0x93, 0x95, 0x60, 0x98, 0x8c, 0x39, 0xb9, 0xf0, 0x0f, 0xd7,
	0xcb, 0x51, 0x50, 0xed, 0x13, 0x17, 0x92, 0x96, 0x60, 0x1f, 0x7c, 0xd5, 0x76, 0xe9, 0x49, 0x21,
	0x6f, 0x81, 0x9f, 0xd4, 0x95, 0x67, 0x68, 0x07, 0x54, 0x6a, 0x68, 0xa1, 0x39, 0xcf, 0x0b, 0xc3,
	0x13, 0xfa, 0x59, 0xf0, 0x85, 0x0b, 0xfe, 0x60, 0xb1, 0xeb, 0x8a, 0x8a, 0xe6, 0x5b, 0x5a, 0xe8,
	0xab, 0x0c, 0x68, 0x01, 0x24, 0x72, 0xec, 0x97, 0x19, 0x5c, 0x83, 0x7d, 0xc4, 0xd0, 0xcb, 0x68,
	0xca, 0x72, 0x91, 0x09, 0x0e, 0xe4, 0xa5, 0xf3, 0x16, 0xa1, 0x97, 0x5d, 0x04, 0xc9, 0xe6, 0xa0,
	0x47, 0x33, 0x20, 0x57, 0x3e, 0x07, 0xbd, 0x01, 0xe0, 0x7d, 0xb4, 0x92, 0x6a, 0x4d, 0x59, 0x6c,
	0xc4, 0x7b, 0x4e, 0x7d, 0xef, 0xc0, 0xf9, 0xcd, 0x54, 0xeb, 0xbe, 0x93, 0x07, 0x76, 0x81, 0x53,
	0xb4, 0x33, 0x03, 0x3e, 0xda, 0x86, 0xcf, 0xbd, 0x72, 0xb9, 0xcd, 0x3a, 0x37, 0x9a, 0xdd, 0x84,
	0x2d, 0x19, 0x7f, 0xed, 0xfe, 0xe5, 0x0f, 0x3e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xca, 0x96,
	0x4b, 0x06, 0x06, 0x00, 0x00,
}