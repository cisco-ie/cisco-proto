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
// source: mpls_te_fastreroute_protection_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_fast_reroute_protections_protection

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

type MplsTeFastrerouteProtectionBag_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	TunnelId             uint32   `protobuf:"varint,3,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,4,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	LspId                uint32   `protobuf:"varint,5,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	SubGroupId           uint32   `protobuf:"varint,6,opt,name=sub_group_id,json=subGroupId,proto3" json:"sub_group_id,omitempty"`
	SubGroupOriginator   string   `protobuf:"bytes,7,opt,name=sub_group_originator,json=subGroupOriginator,proto3" json:"sub_group_originator,omitempty"`
	P2MpId               uint32   `protobuf:"varint,8,opt,name=p2mp_id,json=p2mpId,proto3" json:"p2mp_id,omitempty"`
	SessionType          uint32   `protobuf:"varint,9,opt,name=session_type,json=sessionType,proto3" json:"session_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeFastrerouteProtectionBag_KEYS) Reset()         { *m = MplsTeFastrerouteProtectionBag_KEYS{} }
func (m *MplsTeFastrerouteProtectionBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeFastrerouteProtectionBag_KEYS) ProtoMessage()    {}
func (*MplsTeFastrerouteProtectionBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2e93535459d427, []int{0}
}

func (m *MplsTeFastrerouteProtectionBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeFastrerouteProtectionBag_KEYS.Unmarshal(m, b)
}
func (m *MplsTeFastrerouteProtectionBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeFastrerouteProtectionBag_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeFastrerouteProtectionBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeFastrerouteProtectionBag_KEYS.Merge(m, src)
}
func (m *MplsTeFastrerouteProtectionBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeFastrerouteProtectionBag_KEYS.Size(m)
}
func (m *MplsTeFastrerouteProtectionBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeFastrerouteProtectionBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeFastrerouteProtectionBag_KEYS proto.InternalMessageInfo

func (m *MplsTeFastrerouteProtectionBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag_KEYS) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag_KEYS) GetSubGroupOriginator() string {
	if m != nil {
		return m.SubGroupOriginator
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag_KEYS) GetSessionType() uint32 {
	if m != nil {
		return m.SessionType
	}
	return 0
}

type MplsTeFastrerouteProtectionBag struct {
	SourceAddressXr          string   `protobuf:"bytes,50,opt,name=source_address_xr,json=sourceAddressXr,proto3" json:"source_address_xr,omitempty"`
	DestinationAddressXr     string   `protobuf:"bytes,51,opt,name=destination_address_xr,json=destinationAddressXr,proto3" json:"destination_address_xr,omitempty"`
	TunnelIdXr               uint32   `protobuf:"varint,52,opt,name=tunnel_id_xr,json=tunnelIdXr,proto3" json:"tunnel_id_xr,omitempty"`
	ExtendedTunnelIdXr       string   `protobuf:"bytes,53,opt,name=extended_tunnel_id_xr,json=extendedTunnelIdXr,proto3" json:"extended_tunnel_id_xr,omitempty"`
	TunnelName               string   `protobuf:"bytes,54,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	TunnelInstance           uint32   `protobuf:"varint,55,opt,name=tunnel_instance,json=tunnelInstance,proto3" json:"tunnel_instance,omitempty"`
	Role                     string   `protobuf:"bytes,56,opt,name=role,proto3" json:"role,omitempty"`
	AdminStatus              string   `protobuf:"bytes,57,opt,name=admin_status,json=adminStatus,proto3" json:"admin_status,omitempty"`
	OperationStatus          string   `protobuf:"bytes,58,opt,name=operation_status,json=operationStatus,proto3" json:"operation_status,omitempty"`
	IsSignalled              bool     `protobuf:"varint,59,opt,name=is_signalled,json=isSignalled,proto3" json:"is_signalled,omitempty"`
	IsFrrRequested           bool     `protobuf:"varint,60,opt,name=is_frr_requested,json=isFrrRequested,proto3" json:"is_frr_requested,omitempty"`
	OutboundFrrState         string   `protobuf:"bytes,61,opt,name=outbound_frr_state,json=outboundFrrState,proto3" json:"outbound_frr_state,omitempty"`
	InboundFrrState          string   `protobuf:"bytes,62,opt,name=inbound_frr_state,json=inboundFrrState,proto3" json:"inbound_frr_state,omitempty"`
	OutputInterfaceLsp       string   `protobuf:"bytes,63,opt,name=output_interface_lsp,json=outputInterfaceLsp,proto3" json:"output_interface_lsp,omitempty"`
	OutputLabel              uint32   `protobuf:"varint,64,opt,name=output_label,json=outputLabel,proto3" json:"output_label,omitempty"`
	NextHopAddress           string   `protobuf:"bytes,65,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	FrrOutputTunnelInterface string   `protobuf:"bytes,66,opt,name=frr_output_tunnel_interface,json=frrOutputTunnelInterface,proto3" json:"frr_output_tunnel_interface,omitempty"`
	FrrOutputLabel           uint32   `protobuf:"varint,67,opt,name=frr_output_label,json=frrOutputLabel,proto3" json:"frr_output_label,omitempty"`
	BackupStatus             string   `protobuf:"bytes,68,opt,name=backup_status,json=backupStatus,proto3" json:"backup_status,omitempty"`
	BandwidthPoolType        string   `protobuf:"bytes,69,opt,name=bandwidth_pool_type,json=bandwidthPoolType,proto3" json:"bandwidth_pool_type,omitempty"`
	BandwidthLimitType       string   `protobuf:"bytes,70,opt,name=bandwidth_limit_type,json=bandwidthLimitType,proto3" json:"bandwidth_limit_type,omitempty"`
	OriginalInputInterface   string   `protobuf:"bytes,71,opt,name=original_input_interface,json=originalInputInterface,proto3" json:"original_input_interface,omitempty"`
	InputLabel               uint32   `protobuf:"varint,72,opt,name=input_label,json=inputLabel,proto3" json:"input_label,omitempty"`
	PreviousHopAddress       string   `protobuf:"bytes,73,opt,name=previous_hop_address,json=previousHopAddress,proto3" json:"previous_hop_address,omitempty"`
	BackupBandwidth          uint32   `protobuf:"varint,74,opt,name=backup_bandwidth,json=backupBandwidth,proto3" json:"backup_bandwidth,omitempty"`
	FrrOutputInterface       string   `protobuf:"bytes,75,opt,name=frr_output_interface,json=frrOutputInterface,proto3" json:"frr_output_interface,omitempty"`
	BackupNextHopAddress     string   `protobuf:"bytes,76,opt,name=backup_next_hop_address,json=backupNextHopAddress,proto3" json:"backup_next_hop_address,omitempty"`
	LspBandwidthType         string   `protobuf:"bytes,77,opt,name=lsp_bandwidth_type,json=lspBandwidthType,proto3" json:"lsp_bandwidth_type,omitempty"`
	SharingType              string   `protobuf:"bytes,78,opt,name=sharing_type,json=sharingType,proto3" json:"sharing_type,omitempty"`
	IsP2MpTunnel             bool     `protobuf:"varint,79,opt,name=is_p2mp_tunnel,json=isP2mpTunnel,proto3" json:"is_p2mp_tunnel,omitempty"`
	SubGroupOriginalId       string   `protobuf:"bytes,80,opt,name=sub_group_original_id,json=subGroupOriginalId,proto3" json:"sub_group_original_id,omitempty"`
	SubGroupIdXr             uint32   `protobuf:"varint,81,opt,name=sub_group_id_xr,json=subGroupIdXr,proto3" json:"sub_group_id_xr,omitempty"`
	P2MpIdXr                 uint32   `protobuf:"varint,82,opt,name=p2mp_id_xr,json=p2mpIdXr,proto3" json:"p2mp_id_xr,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *MplsTeFastrerouteProtectionBag) Reset()         { *m = MplsTeFastrerouteProtectionBag{} }
func (m *MplsTeFastrerouteProtectionBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeFastrerouteProtectionBag) ProtoMessage()    {}
func (*MplsTeFastrerouteProtectionBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2e93535459d427, []int{1}
}

func (m *MplsTeFastrerouteProtectionBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeFastrerouteProtectionBag.Unmarshal(m, b)
}
func (m *MplsTeFastrerouteProtectionBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeFastrerouteProtectionBag.Marshal(b, m, deterministic)
}
func (m *MplsTeFastrerouteProtectionBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeFastrerouteProtectionBag.Merge(m, src)
}
func (m *MplsTeFastrerouteProtectionBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeFastrerouteProtectionBag.Size(m)
}
func (m *MplsTeFastrerouteProtectionBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeFastrerouteProtectionBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeFastrerouteProtectionBag proto.InternalMessageInfo

func (m *MplsTeFastrerouteProtectionBag) GetSourceAddressXr() string {
	if m != nil {
		return m.SourceAddressXr
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetDestinationAddressXr() string {
	if m != nil {
		return m.DestinationAddressXr
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetTunnelIdXr() uint32 {
	if m != nil {
		return m.TunnelIdXr
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag) GetExtendedTunnelIdXr() string {
	if m != nil {
		return m.ExtendedTunnelIdXr
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetTunnelInstance() uint32 {
	if m != nil {
		return m.TunnelInstance
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetAdminStatus() string {
	if m != nil {
		return m.AdminStatus
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetOperationStatus() string {
	if m != nil {
		return m.OperationStatus
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetIsSignalled() bool {
	if m != nil {
		return m.IsSignalled
	}
	return false
}

func (m *MplsTeFastrerouteProtectionBag) GetIsFrrRequested() bool {
	if m != nil {
		return m.IsFrrRequested
	}
	return false
}

func (m *MplsTeFastrerouteProtectionBag) GetOutboundFrrState() string {
	if m != nil {
		return m.OutboundFrrState
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetInboundFrrState() string {
	if m != nil {
		return m.InboundFrrState
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetOutputInterfaceLsp() string {
	if m != nil {
		return m.OutputInterfaceLsp
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetOutputLabel() uint32 {
	if m != nil {
		return m.OutputLabel
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetFrrOutputTunnelInterface() string {
	if m != nil {
		return m.FrrOutputTunnelInterface
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetFrrOutputLabel() uint32 {
	if m != nil {
		return m.FrrOutputLabel
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag) GetBackupStatus() string {
	if m != nil {
		return m.BackupStatus
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetBandwidthPoolType() string {
	if m != nil {
		return m.BandwidthPoolType
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetBandwidthLimitType() string {
	if m != nil {
		return m.BandwidthLimitType
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetOriginalInputInterface() string {
	if m != nil {
		return m.OriginalInputInterface
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetInputLabel() uint32 {
	if m != nil {
		return m.InputLabel
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag) GetPreviousHopAddress() string {
	if m != nil {
		return m.PreviousHopAddress
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetBackupBandwidth() uint32 {
	if m != nil {
		return m.BackupBandwidth
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag) GetFrrOutputInterface() string {
	if m != nil {
		return m.FrrOutputInterface
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetBackupNextHopAddress() string {
	if m != nil {
		return m.BackupNextHopAddress
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetLspBandwidthType() string {
	if m != nil {
		return m.LspBandwidthType
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetSharingType() string {
	if m != nil {
		return m.SharingType
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetIsP2MpTunnel() bool {
	if m != nil {
		return m.IsP2MpTunnel
	}
	return false
}

func (m *MplsTeFastrerouteProtectionBag) GetSubGroupOriginalId() string {
	if m != nil {
		return m.SubGroupOriginalId
	}
	return ""
}

func (m *MplsTeFastrerouteProtectionBag) GetSubGroupIdXr() uint32 {
	if m != nil {
		return m.SubGroupIdXr
	}
	return 0
}

func (m *MplsTeFastrerouteProtectionBag) GetP2MpIdXr() uint32 {
	if m != nil {
		return m.P2MpIdXr
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeFastrerouteProtectionBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fast_reroute.protections.protection.mpls_te_fastreroute_protection_bag_KEYS")
	proto.RegisterType((*MplsTeFastrerouteProtectionBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fast_reroute.protections.protection.mpls_te_fastreroute_protection_bag")
}

func init() {
	proto.RegisterFile("mpls_te_fastreroute_protection_bag.proto", fileDescriptor_0a2e93535459d427)
}

var fileDescriptor_0a2e93535459d427 = []byte{
	// 895 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x7f, 0x73, 0x13, 0x37,
	0x10, 0x9d, 0x50, 0x08, 0xc9, 0xda, 0x71, 0x82, 0x08, 0xa0, 0x19, 0x3a, 0xd3, 0x90, 0x96, 0xc1,
	0x74, 0x18, 0x37, 0x0d, 0xd0, 0xd2, 0x1f, 0xb4, 0x85, 0x96, 0x80, 0x4b, 0x48, 0x52, 0x87, 0x3f,
	0xdc, 0xbf, 0x34, 0x67, 0x9f, 0xe2, 0x68, 0x7a, 0x3e, 0x5d, 0x25, 0x5d, 0x9b, 0x7c, 0xa8, 0x7e,
	0xba, 0x7e, 0x01, 0x66, 0x77, 0xa5, 0xb3, 0x4d, 0x98, 0xe1, 0xbf, 0xbb, 0xf7, 0xde, 0xae, 0x56,
	0x4f, 0xab, 0x15, 0x74, 0xa7, 0x55, 0xe1, 0x55, 0xd0, 0xea, 0x24, 0xf3, 0xc1, 0x69, 0x67, 0xeb,
	0xa0, 0x55, 0xe5, 0x6c, 0xd0, 0xe3, 0x60, 0x6c, 0xa9, 0x46, 0xd9, 0xa4, 0x87, 0xbf, 0x56, 0xbc,
	0x19, 0x1b, 0x3f, 0xb6, 0xca, 0x58, 0xaf, 0xce, 0x9c, 0x4a, 0x61, 0xb6, 0xd2, 0xae, 0x97, 0x7e,
	0x7c, 0xc8, 0xca, 0x7c, 0x74, 0xde, 0xc3, 0x5c, 0x2a, 0x26, 0xeb, 0xcd, 0x92, 0xf9, 0xb9, 0xef,
	0xed, 0xff, 0x2f, 0xc1, 0xbd, 0x8f, 0xaf, 0xad, 0x5e, 0xbf, 0xf8, 0xf3, 0x58, 0xdc, 0x85, 0x8e,
	0xb7, 0xb5, 0x1b, 0x6b, 0x95, 0xe5, 0xb9, 0xd3, 0xde, 0xcb, 0xa5, 0xad, 0xa5, 0xee, 0xea, 0x60,
	0x8d, 0xd1, 0x67, 0x0c, 0x8a, 0xaf, 0xe0, 0x7a, 0xae, 0x7d, 0x30, 0x65, 0x46, 0xe1, 0x49, 0x7b,
	0x89, 0xb4, 0x62, 0x8e, 0x4a, 0x01, 0xb7, 0x61, 0x35, 0xd4, 0x65, 0xa9, 0x0b, 0x65, 0x72, 0xf9,
	0xc9, 0xd6, 0x52, 0x77, 0x6d, 0xb0, 0xc2, 0x40, 0x3f, 0x17, 0x0f, 0x40, 0xe8, 0xb3, 0xa0, 0xcb,
	0x5c, 0xe7, 0x6a, 0xa6, 0xba, 0x4c, 0xc9, 0x36, 0x12, 0xf3, 0x36, 0xa9, 0x6f, 0xc0, 0x72, 0xe1,
	0x2b, 0x54, 0x5c, 0xa1, 0x3c, 0x57, 0x0a, 0x5f, 0xf5, 0x73, 0xb1, 0x05, 0x6d, 0x5f, 0x8f, 0xd4,
	0xc4, 0xd9, 0x9a, 0xc8, 0x65, 0x22, 0xc1, 0xd7, 0xa3, 0x97, 0x08, 0xf5, 0x73, 0xb1, 0x03, 0x9b,
	0x33, 0x85, 0x75, 0x66, 0x82, 0x35, 0x5a, 0x27, 0xaf, 0x72, 0xd5, 0x49, 0x79, 0xd8, 0x30, 0xe2,
	0x16, 0x5c, 0xad, 0x76, 0xa7, 0x94, 0x6e, 0x85, 0xd2, 0x2d, 0xe3, 0x6f, 0x3f, 0x17, 0x77, 0xa0,
	0xed, 0xb5, 0xf7, 0xb8, 0xf7, 0x70, 0x5e, 0x69, 0xb9, 0x4a, 0x6c, 0x2b, 0x62, 0x6f, 0xcf, 0x2b,
	0xbd, 0xfd, 0x5f, 0x0b, 0xb6, 0x3f, 0xee, 0xba, 0xf8, 0x12, 0xae, 0x2d, 0x1a, 0xae, 0xce, 0x9c,
	0xdc, 0xa5, 0x8a, 0xd6, 0x17, 0x3c, 0x1f, 0x3a, 0xf1, 0x08, 0x6e, 0x7e, 0xc0, 0x75, 0x0c, 0x78,
	0x48, 0x01, 0x9b, 0x17, 0x8d, 0x1f, 0x3a, 0x34, 0xa6, 0x31, 0x15, 0xb5, 0x8f, 0xd8, 0x98, 0xe4,
	0xfe, 0xd0, 0x89, 0xaf, 0xe1, 0xc6, 0x45, 0xff, 0x51, 0xfa, 0x98, 0x9d, 0x79, 0xff, 0x08, 0x86,
	0x4e, 0x7c, 0x06, 0xad, 0xa8, 0x2c, 0xb3, 0xa9, 0x96, 0xdf, 0x90, 0x30, 0xe6, 0x3c, 0xc8, 0xa6,
	0x5a, 0xdc, 0x83, 0xf5, 0x94, 0xaa, 0xc4, 0x66, 0x1d, 0x6b, 0xf9, 0x2d, 0x2d, 0xdc, 0x89, 0x0b,
	0x47, 0x54, 0x08, 0xb8, 0xec, 0x6c, 0xa1, 0xe5, 0x13, 0x4a, 0x41, 0xdf, 0x68, 0x6f, 0x96, 0x4f,
	0x4d, 0x89, 0x6d, 0x1e, 0x6a, 0x2f, 0xbf, 0x23, 0xae, 0x45, 0xd8, 0x31, 0x41, 0xe2, 0x3e, 0x6c,
	0xe0, 0x85, 0x60, 0x27, 0xa2, 0xec, 0x7b, 0xb6, 0xad, 0xc1, 0xa3, 0xf4, 0x0e, 0xb4, 0x8d, 0x57,
	0xde, 0x4c, 0xca, 0xac, 0x28, 0x74, 0x2e, 0x7f, 0xd8, 0x5a, 0xea, 0xae, 0x0c, 0x5a, 0xc6, 0x1f,
	0x27, 0x48, 0x74, 0x61, 0xc3, 0x78, 0x75, 0xe2, 0x9c, 0x72, 0xfa, 0xef, 0x5a, 0xfb, 0xa0, 0x73,
	0xf9, 0x23, 0xc9, 0x3a, 0xc6, 0xef, 0x39, 0x37, 0x48, 0x28, 0xf6, 0xaa, 0xad, 0xc3, 0xc8, 0xd6,
	0x65, 0x4e, 0x7a, 0x5c, 0x5a, 0xcb, 0xa7, 0xdc, 0xab, 0x89, 0xd9, 0x73, 0x0e, 0xd7, 0xd6, 0x78,
	0xba, 0xa6, 0x7c, 0x5f, 0xfc, 0x13, 0x97, 0x19, 0x89, 0x46, 0xbb, 0x03, 0x9b, 0xb6, 0x0e, 0x55,
	0x1d, 0x94, 0x29, 0x83, 0x76, 0x27, 0xd9, 0x58, 0xab, 0xc2, 0x57, 0xf2, 0x67, 0x3e, 0x04, 0xe6,
	0xfa, 0x89, 0xda, 0xf7, 0x15, 0x6e, 0x2c, 0x46, 0x14, 0xd9, 0x48, 0x17, 0xf2, 0x17, 0xee, 0x42,
	0xc6, 0xf6, 0x11, 0xc2, 0x8d, 0x95, 0xfa, 0x2c, 0xa8, 0x53, 0x5b, 0x35, 0xb7, 0xf4, 0x19, 0x25,
	0xec, 0x20, 0xfe, 0xca, 0x56, 0xe9, 0x86, 0x3e, 0x85, 0xdb, 0x58, 0x62, 0x4c, 0xd8, 0x9c, 0x5d,
	0x5c, 0x4e, 0x3e, 0xa7, 0x20, 0x79, 0xe2, 0xdc, 0x21, 0x29, 0x62, 0x2f, 0x24, 0x1e, 0x17, 0x9a,
	0x0b, 0xe7, 0x7a, 0x7e, 0xe5, 0x03, 0x6f, 0x62, 0xb8, 0xa4, 0xcf, 0x61, 0x6d, 0x94, 0x8d, 0xff,
	0xaa, 0xab, 0x74, 0x6c, 0xbf, 0x51, 0xea, 0x36, 0x83, 0xf1, 0xcc, 0x7a, 0x70, 0x7d, 0x94, 0x95,
	0xf9, 0xbf, 0x26, 0x0f, 0xa7, 0xaa, 0xb2, 0xb6, 0xe0, 0x7b, 0xf6, 0x82, 0xa4, 0xd7, 0x1a, 0xea,
	0xc8, 0xda, 0x02, 0x6f, 0x1b, 0x9a, 0x37, 0xd3, 0x17, 0x66, 0x6a, 0x02, 0x07, 0xec, 0xb1, 0x79,
	0x0d, 0xb7, 0x8f, 0x14, 0x45, 0x3c, 0x01, 0x19, 0x67, 0x00, 0x6e, 0x73, 0xc1, 0x76, 0xf9, 0x92,
	0xa2, 0x6e, 0x26, 0xbe, 0x5f, 0xce, 0x3b, 0x8f, 0xbd, 0xcf, 0x01, 0xbc, 0xcb, 0x57, 0x7c, 0x9f,
	0x08, 0xe2, 0x1d, 0xee, 0xc0, 0x66, 0xe5, 0xf4, 0x3f, 0xc6, 0xd6, 0x7e, 0xc1, 0xf8, 0x3e, 0x17,
	0x93, 0xb8, 0x39, 0xf3, 0xef, 0xc3, 0x46, 0xf4, 0xa4, 0xa9, 0x54, 0xfe, 0x4e, 0x79, 0xd7, 0x19,
	0x7f, 0x9e, 0x60, 0x4c, 0x3e, 0x67, 0xf4, 0xac, 0xe6, 0xd7, 0x9c, 0xbc, 0x31, 0x7b, 0x56, 0xef,
	0x63, 0xb8, 0x15, 0x93, 0x5f, 0x68, 0x85, 0x7d, 0x9e, 0x1b, 0x4c, 0x1f, 0x2c, 0x36, 0xc4, 0x03,
	0x10, 0x38, 0x67, 0x67, 0xb6, 0x92, 0xa1, 0x6f, 0xb8, 0xd3, 0x0b, 0x3f, 0x2b, 0x89, 0xec, 0xc4,
	0x89, 0x78, 0x9a, 0x39, 0x53, 0x4e, 0x58, 0x77, 0xc0, 0x57, 0x36, 0x62, 0x24, 0xf9, 0x02, 0x3a,
	0xc6, 0x2b, 0x1a, 0xa8, 0xdc, 0x5e, 0xf2, 0x90, 0xae, 0x58, 0xdb, 0xf8, 0xa3, 0xdd, 0x69, 0xc5,
	0x1d, 0x85, 0xc3, 0xe8, 0xc2, 0x94, 0xa6, 0xf7, 0xe0, 0xe8, 0x83, 0x63, 0x1a, 0x5f, 0x84, 0xbb,
	0xb0, 0x3e, 0x3f, 0xfa, 0x71, 0x72, 0xfd, 0x41, 0xe6, 0xb5, 0x67, 0xd3, 0x7f, 0xe8, 0xc4, 0xa7,
	0x00, 0x71, 0x9a, 0xa3, 0x62, 0xc0, 0x8f, 0x10, 0x0f, 0xf4, 0xa1, 0x1b, 0x2d, 0xd3, 0xdb, 0xfb,
	0xf0, 0x5d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x38, 0x7a, 0xd0, 0xa7, 0x07, 0x00, 0x00,
}
