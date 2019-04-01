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
// source: mpls_lm_preemption_event.proto

package cisco_ios_xr_mpls_te_oper_mpls_lcac_preemption_events_preemption_event

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

type MplsLmPreemptionEvent_KEYS struct {
	EventIndex           uint32   `protobuf:"varint,1,opt,name=event_index,json=eventIndex,proto3" json:"event_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmPreemptionEvent_KEYS) Reset()         { *m = MplsLmPreemptionEvent_KEYS{} }
func (m *MplsLmPreemptionEvent_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLmPreemptionEvent_KEYS) ProtoMessage()    {}
func (*MplsLmPreemptionEvent_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_70a021b606bd7715, []int{0}
}

func (m *MplsLmPreemptionEvent_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmPreemptionEvent_KEYS.Unmarshal(m, b)
}
func (m *MplsLmPreemptionEvent_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmPreemptionEvent_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLmPreemptionEvent_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmPreemptionEvent_KEYS.Merge(m, src)
}
func (m *MplsLmPreemptionEvent_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLmPreemptionEvent_KEYS.Size(m)
}
func (m *MplsLmPreemptionEvent_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmPreemptionEvent_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmPreemptionEvent_KEYS proto.InternalMessageInfo

func (m *MplsLmPreemptionEvent_KEYS) GetEventIndex() uint32 {
	if m != nil {
		return m.EventIndex
	}
	return 0
}

type MplsLmPreemptedLsp struct {
	TunnelId                uint32   `protobuf:"varint,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	LspId                   uint32   `protobuf:"varint,2,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	SourceAddress           string   `protobuf:"bytes,3,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress      string   `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	RequestedBandwidth      uint64   `protobuf:"varint,5,opt,name=requested_bandwidth,json=requestedBandwidth,proto3" json:"requested_bandwidth,omitempty"`
	SetupPriority           uint32   `protobuf:"varint,6,opt,name=setup_priority,json=setupPriority,proto3" json:"setup_priority,omitempty"`
	HoldPriority            uint32   `protobuf:"varint,7,opt,name=hold_priority,json=holdPriority,proto3" json:"hold_priority,omitempty"`
	BandwidthType           uint32   `protobuf:"varint,8,opt,name=bandwidth_type,json=bandwidthType,proto3" json:"bandwidth_type,omitempty"`
	SoftPreempted           bool     `protobuf:"varint,9,opt,name=soft_preempted,json=softPreempted,proto3" json:"soft_preempted,omitempty"`
	SoftPreemptionTimeout   uint32   `protobuf:"varint,10,opt,name=soft_preemption_timeout,json=softPreemptionTimeout,proto3" json:"soft_preemption_timeout,omitempty"`
	BandwidthPreempted      bool     `protobuf:"varint,11,opt,name=bandwidth_preempted,json=bandwidthPreempted,proto3" json:"bandwidth_preempted,omitempty"`
	SoftPreemptedFrRrewrite bool     `protobuf:"varint,12,opt,name=soft_preempted_fr_rrewrite,json=softPreemptedFrRrewrite,proto3" json:"soft_preempted_fr_rrewrite,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *MplsLmPreemptedLsp) Reset()         { *m = MplsLmPreemptedLsp{} }
func (m *MplsLmPreemptedLsp) String() string { return proto.CompactTextString(m) }
func (*MplsLmPreemptedLsp) ProtoMessage()    {}
func (*MplsLmPreemptedLsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_70a021b606bd7715, []int{1}
}

func (m *MplsLmPreemptedLsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmPreemptedLsp.Unmarshal(m, b)
}
func (m *MplsLmPreemptedLsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmPreemptedLsp.Marshal(b, m, deterministic)
}
func (m *MplsLmPreemptedLsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmPreemptedLsp.Merge(m, src)
}
func (m *MplsLmPreemptedLsp) XXX_Size() int {
	return xxx_messageInfo_MplsLmPreemptedLsp.Size(m)
}
func (m *MplsLmPreemptedLsp) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmPreemptedLsp.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmPreemptedLsp proto.InternalMessageInfo

func (m *MplsLmPreemptedLsp) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsLmPreemptedLsp) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *MplsLmPreemptedLsp) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsLmPreemptedLsp) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsLmPreemptedLsp) GetRequestedBandwidth() uint64 {
	if m != nil {
		return m.RequestedBandwidth
	}
	return 0
}

func (m *MplsLmPreemptedLsp) GetSetupPriority() uint32 {
	if m != nil {
		return m.SetupPriority
	}
	return 0
}

func (m *MplsLmPreemptedLsp) GetHoldPriority() uint32 {
	if m != nil {
		return m.HoldPriority
	}
	return 0
}

func (m *MplsLmPreemptedLsp) GetBandwidthType() uint32 {
	if m != nil {
		return m.BandwidthType
	}
	return 0
}

func (m *MplsLmPreemptedLsp) GetSoftPreempted() bool {
	if m != nil {
		return m.SoftPreempted
	}
	return false
}

func (m *MplsLmPreemptedLsp) GetSoftPreemptionTimeout() uint32 {
	if m != nil {
		return m.SoftPreemptionTimeout
	}
	return 0
}

func (m *MplsLmPreemptedLsp) GetBandwidthPreempted() bool {
	if m != nil {
		return m.BandwidthPreempted
	}
	return false
}

func (m *MplsLmPreemptedLsp) GetSoftPreemptedFrRrewrite() bool {
	if m != nil {
		return m.SoftPreemptedFrRrewrite
	}
	return false
}

type MplsLmPreemptionEvent struct {
	TunnelId                            uint32                `protobuf:"varint,50,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	LspId                               uint32                `protobuf:"varint,51,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	SourceAddress                       string                `protobuf:"bytes,52,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress                  string                `protobuf:"bytes,53,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	RequestedBandwidth                  uint64                `protobuf:"varint,54,opt,name=requested_bandwidth,json=requestedBandwidth,proto3" json:"requested_bandwidth,omitempty"`
	SetupPriority                       uint32                `protobuf:"varint,55,opt,name=setup_priority,json=setupPriority,proto3" json:"setup_priority,omitempty"`
	HoldPriority                        uint32                `protobuf:"varint,56,opt,name=hold_priority,json=holdPriority,proto3" json:"hold_priority,omitempty"`
	BandwidthType                       uint32                `protobuf:"varint,57,opt,name=bandwidth_type,json=bandwidthType,proto3" json:"bandwidth_type,omitempty"`
	OldBandwidthBc0                     uint64                `protobuf:"varint,58,opt,name=old_bandwidth_bc0,json=oldBandwidthBc0,proto3" json:"old_bandwidth_bc0,omitempty"`
	OldBandwidthBc1                     uint64                `protobuf:"varint,59,opt,name=old_bandwidth_bc1,json=oldBandwidthBc1,proto3" json:"old_bandwidth_bc1,omitempty"`
	NewBandwidthBc0                     uint64                `protobuf:"varint,60,opt,name=new_bandwidth_bc0,json=newBandwidthBc0,proto3" json:"new_bandwidth_bc0,omitempty"`
	NewBandwidthBc1                     uint64                `protobuf:"varint,61,opt,name=new_bandwidth_bc1,json=newBandwidthBc1,proto3" json:"new_bandwidth_bc1,omitempty"`
	BandwidthOvershoot0                 uint64                `protobuf:"varint,62,opt,name=bandwidth_overshoot0,json=bandwidthOvershoot0,proto3" json:"bandwidth_overshoot0,omitempty"`
	BandwidthOvershoot1                 uint64                `protobuf:"varint,63,opt,name=bandwidth_overshoot1,json=bandwidthOvershoot1,proto3" json:"bandwidth_overshoot1,omitempty"`
	InterfaceName                       string                `protobuf:"bytes,64,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	EventTime                           uint32                `protobuf:"varint,65,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	LsPs                                uint32                `protobuf:"varint,66,opt,name=ls_ps,json=lsPs,proto3" json:"ls_ps,omitempty"`
	SoftPreemptedLsPs                   uint32                `protobuf:"varint,67,opt,name=soft_preempted_ls_ps,json=softPreemptedLsPs,proto3" json:"soft_preempted_ls_ps,omitempty"`
	SoftPreemptedLsPsFrRrewrite         uint32                `protobuf:"varint,68,opt,name=soft_preempted_ls_ps_fr_rrewrite,json=softPreemptedLsPsFrRrewrite,proto3" json:"soft_preempted_ls_ps_fr_rrewrite,omitempty"`
	HardPreemptedLsPs                   uint32                `protobuf:"varint,69,opt,name=hard_preempted_ls_ps,json=hardPreemptedLsPs,proto3" json:"hard_preempted_ls_ps,omitempty"`
	TotalPreemptedBandwidthBc0          uint64                `protobuf:"varint,70,opt,name=total_preempted_bandwidth_bc0,json=totalPreemptedBandwidthBc0,proto3" json:"total_preempted_bandwidth_bc0,omitempty"`
	TotalPreemptedBandwidthBc1          uint64                `protobuf:"varint,71,opt,name=total_preempted_bandwidth_bc1,json=totalPreemptedBandwidthBc1,proto3" json:"total_preempted_bandwidth_bc1,omitempty"`
	SoftlyPreemptedBandwidthBc0         uint64                `protobuf:"varint,72,opt,name=softly_preempted_bandwidth_bc0,json=softlyPreemptedBandwidthBc0,proto3" json:"softly_preempted_bandwidth_bc0,omitempty"`
	SoftlyPreemptedBandwidthBc1         uint64                `protobuf:"varint,73,opt,name=softly_preempted_bandwidth_bc1,json=softlyPreemptedBandwidthBc1,proto3" json:"softly_preempted_bandwidth_bc1,omitempty"`
	SoftPreemptedFrRrewriteBandwidthBc0 uint64                `protobuf:"varint,74,opt,name=soft_preempted_fr_rrewrite_bandwidth_bc0,json=softPreemptedFrRrewriteBandwidthBc0,proto3" json:"soft_preempted_fr_rrewrite_bandwidth_bc0,omitempty"`
	SoftPreemptedFrRrewriteBandwidthBc1 uint64                `protobuf:"varint,75,opt,name=soft_preempted_fr_rrewrite_bandwidth_bc1,json=softPreemptedFrRrewriteBandwidthBc1,proto3" json:"soft_preempted_fr_rrewrite_bandwidth_bc1,omitempty"`
	HardPreemptedBandwidthBc0           uint64                `protobuf:"varint,76,opt,name=hard_preempted_bandwidth_bc0,json=hardPreemptedBandwidthBc0,proto3" json:"hard_preempted_bandwidth_bc0,omitempty"`
	HardPreemptedBandwidthBc1           uint64                `protobuf:"varint,77,opt,name=hard_preempted_bandwidth_bc1,json=hardPreemptedBandwidthBc1,proto3" json:"hard_preempted_bandwidth_bc1,omitempty"`
	Lsp                                 []*MplsLmPreemptedLsp `protobuf:"bytes,78,rep,name=lsp,proto3" json:"lsp,omitempty"`
	Tunnels                             uint32                `protobuf:"varint,79,opt,name=tunnels,proto3" json:"tunnels,omitempty"`
	SoftPreemptedTunnels                uint32                `protobuf:"varint,80,opt,name=soft_preempted_tunnels,json=softPreemptedTunnels,proto3" json:"soft_preempted_tunnels,omitempty"`
	SoftPreemptedTunnelsFrRrewrite      uint32                `protobuf:"varint,81,opt,name=soft_preempted_tunnels_fr_rrewrite,json=softPreemptedTunnelsFrRrewrite,proto3" json:"soft_preempted_tunnels_fr_rrewrite,omitempty"`
	HardPreemptedTunnels                uint32                `protobuf:"varint,82,opt,name=hard_preempted_tunnels,json=hardPreemptedTunnels,proto3" json:"hard_preempted_tunnels,omitempty"`
	XXX_NoUnkeyedLiteral                struct{}              `json:"-"`
	XXX_unrecognized                    []byte                `json:"-"`
	XXX_sizecache                       int32                 `json:"-"`
}

func (m *MplsLmPreemptionEvent) Reset()         { *m = MplsLmPreemptionEvent{} }
func (m *MplsLmPreemptionEvent) String() string { return proto.CompactTextString(m) }
func (*MplsLmPreemptionEvent) ProtoMessage()    {}
func (*MplsLmPreemptionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_70a021b606bd7715, []int{2}
}

func (m *MplsLmPreemptionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmPreemptionEvent.Unmarshal(m, b)
}
func (m *MplsLmPreemptionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmPreemptionEvent.Marshal(b, m, deterministic)
}
func (m *MplsLmPreemptionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmPreemptionEvent.Merge(m, src)
}
func (m *MplsLmPreemptionEvent) XXX_Size() int {
	return xxx_messageInfo_MplsLmPreemptionEvent.Size(m)
}
func (m *MplsLmPreemptionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmPreemptionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmPreemptionEvent proto.InternalMessageInfo

func (m *MplsLmPreemptionEvent) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsLmPreemptionEvent) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsLmPreemptionEvent) GetRequestedBandwidth() uint64 {
	if m != nil {
		return m.RequestedBandwidth
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetSetupPriority() uint32 {
	if m != nil {
		return m.SetupPriority
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetHoldPriority() uint32 {
	if m != nil {
		return m.HoldPriority
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetBandwidthType() uint32 {
	if m != nil {
		return m.BandwidthType
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetOldBandwidthBc0() uint64 {
	if m != nil {
		return m.OldBandwidthBc0
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetOldBandwidthBc1() uint64 {
	if m != nil {
		return m.OldBandwidthBc1
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetNewBandwidthBc0() uint64 {
	if m != nil {
		return m.NewBandwidthBc0
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetNewBandwidthBc1() uint64 {
	if m != nil {
		return m.NewBandwidthBc1
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetBandwidthOvershoot0() uint64 {
	if m != nil {
		return m.BandwidthOvershoot0
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetBandwidthOvershoot1() uint64 {
	if m != nil {
		return m.BandwidthOvershoot1
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *MplsLmPreemptionEvent) GetEventTime() uint32 {
	if m != nil {
		return m.EventTime
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetLsPs() uint32 {
	if m != nil {
		return m.LsPs
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetSoftPreemptedLsPs() uint32 {
	if m != nil {
		return m.SoftPreemptedLsPs
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetSoftPreemptedLsPsFrRrewrite() uint32 {
	if m != nil {
		return m.SoftPreemptedLsPsFrRrewrite
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetHardPreemptedLsPs() uint32 {
	if m != nil {
		return m.HardPreemptedLsPs
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetTotalPreemptedBandwidthBc0() uint64 {
	if m != nil {
		return m.TotalPreemptedBandwidthBc0
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetTotalPreemptedBandwidthBc1() uint64 {
	if m != nil {
		return m.TotalPreemptedBandwidthBc1
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetSoftlyPreemptedBandwidthBc0() uint64 {
	if m != nil {
		return m.SoftlyPreemptedBandwidthBc0
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetSoftlyPreemptedBandwidthBc1() uint64 {
	if m != nil {
		return m.SoftlyPreemptedBandwidthBc1
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetSoftPreemptedFrRrewriteBandwidthBc0() uint64 {
	if m != nil {
		return m.SoftPreemptedFrRrewriteBandwidthBc0
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetSoftPreemptedFrRrewriteBandwidthBc1() uint64 {
	if m != nil {
		return m.SoftPreemptedFrRrewriteBandwidthBc1
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetHardPreemptedBandwidthBc0() uint64 {
	if m != nil {
		return m.HardPreemptedBandwidthBc0
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetHardPreemptedBandwidthBc1() uint64 {
	if m != nil {
		return m.HardPreemptedBandwidthBc1
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetLsp() []*MplsLmPreemptedLsp {
	if m != nil {
		return m.Lsp
	}
	return nil
}

func (m *MplsLmPreemptionEvent) GetTunnels() uint32 {
	if m != nil {
		return m.Tunnels
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetSoftPreemptedTunnels() uint32 {
	if m != nil {
		return m.SoftPreemptedTunnels
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetSoftPreemptedTunnelsFrRrewrite() uint32 {
	if m != nil {
		return m.SoftPreemptedTunnelsFrRrewrite
	}
	return 0
}

func (m *MplsLmPreemptionEvent) GetHardPreemptedTunnels() uint32 {
	if m != nil {
		return m.HardPreemptedTunnels
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsLmPreemptionEvent_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.preemption_events.preemption_event.mpls_lm_preemption_event_KEYS")
	proto.RegisterType((*MplsLmPreemptedLsp)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.preemption_events.preemption_event.mpls_lm_preempted_lsp")
	proto.RegisterType((*MplsLmPreemptionEvent)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.preemption_events.preemption_event.mpls_lm_preemption_event")
}

func init() { proto.RegisterFile("mpls_lm_preemption_event.proto", fileDescriptor_70a021b606bd7715) }

var fileDescriptor_70a021b606bd7715 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5b, 0x4f, 0xdb, 0x48,
	0x18, 0x55, 0x96, 0x70, 0xc9, 0x40, 0x58, 0x31, 0xdc, 0x66, 0x61, 0x61, 0xa3, 0x20, 0xa4, 0x68,
	0x1f, 0xb2, 0x18, 0x58, 0xf6, 0xc2, 0x6e, 0xb9, 0x15, 0xda, 0x00, 0x85, 0x34, 0xa5, 0x0f, 0x7d,
	0xa8, 0x46, 0xc6, 0x9e, 0x28, 0x96, 0x1c, 0x8f, 0x3b, 0x33, 0x21, 0xe4, 0x7f, 0xf4, 0x3f, 0xf6,
	0x6f, 0x54, 0x33, 0x93, 0xd8, 0x19, 0x63, 0x93, 0x20, 0xf5, 0xd1, 0xdf, 0x77, 0xce, 0xf9, 0x8e,
	0x3f, 0x9f, 0x19, 0x19, 0x6c, 0xb6, 0x43, 0x9f, 0x63, 0xbf, 0x8d, 0x43, 0x46, 0x48, 0x3b, 0x14,
	0x1e, 0x0d, 0x30, 0x79, 0x20, 0x81, 0xa8, 0x86, 0x8c, 0x0a, 0x0a, 0x2f, 0x1c, 0x8f, 0x3b, 0x14,
	0x7b, 0x94, 0xe3, 0x47, 0x86, 0x15, 0x58, 0x10, 0x4c, 0x43, 0xc2, 0xaa, 0x9a, 0xe9, 0xd8, 0x4e,
	0x35, 0xc9, 0xe5, 0x4f, 0x2a, 0xe5, 0x63, 0xb0, 0x91, 0x35, 0x09, 0x5f, 0x9d, 0x7f, 0xfa, 0x00,
	0x7f, 0x03, 0xb3, 0xfa, 0xc9, 0x0b, 0x5c, 0xf2, 0x88, 0x72, 0xa5, 0x5c, 0xa5, 0xd8, 0x00, 0xaa,
	0x54, 0x93, 0x95, 0xf2, 0xd7, 0x3c, 0x58, 0x4e, 0x48, 0x10, 0x17, 0xfb, 0x3c, 0x84, 0xeb, 0xa0,
	0x20, 0x3a, 0x41, 0x40, 0x7c, 0xec, 0xb9, 0x7d, 0xe2, 0x8c, 0x2e, 0xd4, 0x5c, 0xb8, 0x0c, 0xa6,
	0x7c, 0x1e, 0xca, 0xce, 0x4f, 0xaa, 0x33, 0xe9, 0xf3, 0xb0, 0xe6, 0xc2, 0x6d, 0x30, 0xcf, 0x69,
	0x87, 0x39, 0x04, 0xdb, 0xae, 0xcb, 0x08, 0xe7, 0x68, 0xa2, 0x94, 0xab, 0x14, 0x1a, 0x45, 0x5d,
	0x3d, 0xd1, 0x45, 0xf8, 0x07, 0x58, 0x74, 0x09, 0x17, 0x5e, 0x60, 0x2b, 0xbf, 0x03, 0x6c, 0x5e,
	0x61, 0xe1, 0x50, 0x6b, 0x88, 0xc0, 0xc8, 0x97, 0x0e, 0xe1, 0xd2, 0xdc, 0xbd, 0x1d, 0xb8, 0x5d,
	0xcf, 0x15, 0x2d, 0x34, 0x59, 0xca, 0x55, 0xf2, 0x0d, 0x18, 0xb5, 0x4e, 0x07, 0x1d, 0x65, 0x84,
	0x88, 0x4e, 0x88, 0x43, 0xe6, 0x51, 0xe6, 0x89, 0x1e, 0x9a, 0x52, 0x3e, 0x8b, 0xaa, 0x5a, 0xef,
	0x17, 0xe1, 0x16, 0x28, 0xb6, 0xa8, 0xef, 0xc6, 0xa8, 0x69, 0x85, 0x9a, 0x93, 0xc5, 0x08, 0xb4,
	0x0d, 0xe6, 0xa3, 0x91, 0x58, 0xf4, 0x42, 0x82, 0x66, 0xb4, 0x56, 0x54, 0xbd, 0xeb, 0x85, 0x44,
	0xbf, 0x7b, 0x53, 0xc4, 0x5b, 0x44, 0x85, 0x52, 0xae, 0x32, 0x23, 0xdf, 0xbd, 0x29, 0xea, 0x83,
	0x22, 0x3c, 0x00, 0xab, 0xc3, 0x30, 0xf9, 0xfe, 0xc2, 0x6b, 0x13, 0xda, 0x11, 0x08, 0x28, 0xd9,
	0xe5, 0x21, 0xbc, 0x47, 0x83, 0x3b, 0xdd, 0x94, 0x2b, 0x88, 0x5d, 0xc4, 0x33, 0x66, 0xd5, 0x0c,
	0x18, 0xb5, 0xe2, 0x41, 0x87, 0x60, 0xcd, 0xf4, 0x83, 0x9b, 0x0c, 0x33, 0x46, 0xba, 0xcc, 0x13,
	0x04, 0xcd, 0x29, 0xde, 0xaa, 0xe1, 0xed, 0x82, 0x35, 0xfa, 0xed, 0xf2, 0xb7, 0x22, 0x40, 0x59,
	0xc9, 0x32, 0x93, 0xb1, 0x9b, 0x99, 0x8c, 0xbd, 0xe7, 0x93, 0xb1, 0xff, 0x82, 0x64, 0xfc, 0xf9,
	0xd2, 0x64, 0x1c, 0xbc, 0x20, 0x19, 0x7f, 0x8d, 0x95, 0x8c, 0xbf, 0xc7, 0x4a, 0xc6, 0x3f, 0x69,
	0xc9, 0xf8, 0x1d, 0x2c, 0x48, 0xa9, 0x18, 0x7a, 0xef, 0xec, 0xa0, 0x7f, 0x95, 0xc3, 0x9f, 0xa9,
	0x1f, 0x7b, 0x3b, 0x75, 0x76, 0xd2, 0xb0, 0x16, 0x3a, 0x4c, 0xc3, 0x5a, 0x12, 0x1b, 0x90, 0x6e,
	0x42, 0xf7, 0x3f, 0x8d, 0x0d, 0x48, 0x37, 0xa9, 0x9b, 0xc4, 0x5a, 0xe8, 0xff, 0x34, 0xac, 0x05,
	0x2d, 0xb0, 0x14, 0xe3, 0xe8, 0x03, 0x61, 0xbc, 0x45, 0xa9, 0xd8, 0x41, 0xaf, 0x14, 0x3c, 0x8e,
	0xe1, 0x6d, 0xd4, 0xca, 0xa0, 0x58, 0xe8, 0x28, 0x8b, 0x62, 0xc9, 0xe5, 0x79, 0x81, 0x20, 0xac,
	0x69, 0x3b, 0x04, 0x07, 0x76, 0x9b, 0xa0, 0x63, 0x9d, 0x88, 0xa8, 0x7a, 0x63, 0xb7, 0x09, 0xdc,
	0x00, 0xfa, 0xba, 0x52, 0xa7, 0x04, 0x9d, 0xa8, 0xfd, 0x16, 0x54, 0x45, 0x9e, 0x0c, 0xb8, 0x08,
	0x26, 0x7d, 0x8e, 0x43, 0x8e, 0x4e, 0x55, 0x27, 0xef, 0xf3, 0xba, 0x0c, 0xc5, 0x52, 0x22, 0xfa,
	0x1a, 0x73, 0xa6, 0x30, 0x0b, 0x46, 0xe8, 0xaf, 0x25, 0xe1, 0x1c, 0x94, 0xd2, 0x08, 0xc6, 0x89,
	0x79, 0xad, 0xc8, 0xeb, 0x4f, 0xc8, 0xf1, 0xa9, 0x91, 0x73, 0x5b, 0x36, 0x73, 0x9f, 0xcc, 0x3d,
	0xd7, 0x73, 0x65, 0xcf, 0x9c, 0x7b, 0x02, 0x36, 0x04, 0x15, 0xb6, 0x3f, 0xc4, 0x30, 0xbf, 0xe6,
	0x85, 0xda, 0xdf, 0x9a, 0x02, 0x45, 0x54, 0xe3, 0xc3, 0x8e, 0x90, 0xb0, 0xd0, 0x9b, 0x11, 0x12,
	0x16, 0x3c, 0x03, 0x9b, 0xf2, 0xad, 0xfc, 0x5e, 0xa6, 0x8d, 0xb7, 0x4a, 0x63, 0x5d, 0xa3, 0xd2,
	0x7d, 0x8c, 0x12, 0xb1, 0x50, 0x6d, 0x94, 0x88, 0x05, 0x3f, 0x82, 0x4a, 0xf6, 0x9d, 0x95, 0xf0,
	0x74, 0xa9, 0xe4, 0xb6, 0x32, 0x6e, 0x30, 0xc3, 0xdb, 0xf8, 0xb2, 0x16, 0xba, 0x1a, 0x57, 0xd6,
	0x82, 0x47, 0xe0, 0xd7, 0xc4, 0xe7, 0x36, 0x1d, 0x5e, 0x2b, 0xa9, 0x5f, 0x8c, 0xcf, 0x6e, 0xf8,
	0x7a, 0x5e, 0xc0, 0x42, 0xef, 0x9e, 0x17, 0xb0, 0x20, 0x05, 0x13, 0x3e, 0x0f, 0xd1, 0x4d, 0x69,
	0xa2, 0x32, 0xbb, 0xfb, 0xb9, 0xfa, 0x63, 0xfe, 0x2a, 0xaa, 0xa9, 0xff, 0x03, 0x0d, 0x39, 0x09,
	0x22, 0x30, 0xad, 0x6f, 0x7a, 0x8e, 0x6e, 0x55, 0xa8, 0x07, 0x8f, 0x70, 0x1f, 0xac, 0x24, 0x76,
	0x3c, 0x00, 0xd6, 0x15, 0x70, 0xc9, 0xd8, 0xe8, 0x5d, 0x9f, 0x75, 0x09, 0xca, 0xe9, 0x2c, 0xe3,
	0xe8, 0xbd, 0x57, 0x0a, 0x9b, 0x69, 0x0a, 0x43, 0xa7, 0x6f, 0x1f, 0xac, 0x24, 0xb6, 0x39, 0x70,
	0xd0, 0xd0, 0x0e, 0x8c, 0x3d, 0xf6, 0xf9, 0xf7, 0x53, 0xea, 0x8f, 0x6c, 0xef, 0x7b, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa6, 0xe6, 0x03, 0x95, 0xb3, 0x09, 0x00, 0x00,
}
