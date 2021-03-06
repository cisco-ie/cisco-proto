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
// source: mpls_te_s2l_tun_rw_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_p2mp_tunnel_rewrites_tunnel_rewrite

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

type MplsTeS2LTunRwBag_KEYS struct {
	TunnelId             uint32   `protobuf:"varint,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	LspId                uint32   `protobuf:"varint,2,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	Ctype                string   `protobuf:"bytes,3,opt,name=ctype,proto3" json:"ctype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeS2LTunRwBag_KEYS) Reset()         { *m = MplsTeS2LTunRwBag_KEYS{} }
func (m *MplsTeS2LTunRwBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeS2LTunRwBag_KEYS) ProtoMessage()    {}
func (*MplsTeS2LTunRwBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d097dad63b57d04, []int{0}
}

func (m *MplsTeS2LTunRwBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeS2LTunRwBag_KEYS.Unmarshal(m, b)
}
func (m *MplsTeS2LTunRwBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeS2LTunRwBag_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeS2LTunRwBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeS2LTunRwBag_KEYS.Merge(m, src)
}
func (m *MplsTeS2LTunRwBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeS2LTunRwBag_KEYS.Size(m)
}
func (m *MplsTeS2LTunRwBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeS2LTunRwBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeS2LTunRwBag_KEYS proto.InternalMessageInfo

func (m *MplsTeS2LTunRwBag_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsTeS2LTunRwBag_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *MplsTeS2LTunRwBag_KEYS) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

type TeS2LOutputRwOutputInfoBag struct {
	PhysicaInterfaceName string   `protobuf:"bytes,1,opt,name=physica_interface_name,json=physicaInterfaceName,proto3" json:"physica_interface_name,omitempty"`
	TunnelInterfaceName  string   `protobuf:"bytes,2,opt,name=tunnel_interface_name,json=tunnelInterfaceName,proto3" json:"tunnel_interface_name,omitempty"`
	ParentInterfaceName  string   `protobuf:"bytes,3,opt,name=parent_interface_name,json=parentInterfaceName,proto3" json:"parent_interface_name,omitempty"`
	NextHopAddress       string   `protobuf:"bytes,4,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	OutLabel             uint32   `protobuf:"varint,5,opt,name=out_label,json=outLabel,proto3" json:"out_label,omitempty"`
	SrLabelStack         []uint32 `protobuf:"varint,6,rep,packed,name=sr_label_stack,json=srLabelStack,proto3" json:"sr_label_stack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeS2LOutputRwOutputInfoBag) Reset()         { *m = TeS2LOutputRwOutputInfoBag{} }
func (m *TeS2LOutputRwOutputInfoBag) String() string { return proto.CompactTextString(m) }
func (*TeS2LOutputRwOutputInfoBag) ProtoMessage()    {}
func (*TeS2LOutputRwOutputInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d097dad63b57d04, []int{1}
}

func (m *TeS2LOutputRwOutputInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LOutputRwOutputInfoBag.Unmarshal(m, b)
}
func (m *TeS2LOutputRwOutputInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LOutputRwOutputInfoBag.Marshal(b, m, deterministic)
}
func (m *TeS2LOutputRwOutputInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LOutputRwOutputInfoBag.Merge(m, src)
}
func (m *TeS2LOutputRwOutputInfoBag) XXX_Size() int {
	return xxx_messageInfo_TeS2LOutputRwOutputInfoBag.Size(m)
}
func (m *TeS2LOutputRwOutputInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LOutputRwOutputInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LOutputRwOutputInfoBag proto.InternalMessageInfo

func (m *TeS2LOutputRwOutputInfoBag) GetPhysicaInterfaceName() string {
	if m != nil {
		return m.PhysicaInterfaceName
	}
	return ""
}

func (m *TeS2LOutputRwOutputInfoBag) GetTunnelInterfaceName() string {
	if m != nil {
		return m.TunnelInterfaceName
	}
	return ""
}

func (m *TeS2LOutputRwOutputInfoBag) GetParentInterfaceName() string {
	if m != nil {
		return m.ParentInterfaceName
	}
	return ""
}

func (m *TeS2LOutputRwOutputInfoBag) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *TeS2LOutputRwOutputInfoBag) GetOutLabel() uint32 {
	if m != nil {
		return m.OutLabel
	}
	return 0
}

func (m *TeS2LOutputRwOutputInfoBag) GetSrLabelStack() []uint32 {
	if m != nil {
		return m.SrLabelStack
	}
	return nil
}

type TeS2LTunRwFieldsP2PBag struct {
	LocalLabel             uint32                      `protobuf:"varint,1,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	TunnelId               uint32                      `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	SourcePe               string                      `protobuf:"bytes,3,opt,name=source_pe,json=sourcePe,proto3" json:"source_pe,omitempty"`
	DestinationAddress     string                      `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	ExtendedTunnelId       string                      `protobuf:"bytes,5,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	PrimaryS2L             *TeS2LOutputRwOutputInfoBag `protobuf:"bytes,6,opt,name=primary_s2l,json=primaryS2l,proto3" json:"primary_s2l,omitempty"`
	BackupTunnelRewrite    *TeS2LOutputRwOutputInfoBag `protobuf:"bytes,7,opt,name=backup_tunnel_rewrite,json=backupTunnelRewrite,proto3" json:"backup_tunnel_rewrite,omitempty"`
	BackupActive           bool                        `protobuf:"varint,8,opt,name=backup_active,json=backupActive,proto3" json:"backup_active,omitempty"`
	OriginalInputInterface string                      `protobuf:"bytes,9,opt,name=original_input_interface,json=originalInputInterface,proto3" json:"original_input_interface,omitempty"`
	PreviousHopAddress     string                      `protobuf:"bytes,10,opt,name=previous_hop_address,json=previousHopAddress,proto3" json:"previous_hop_address,omitempty"`
	OutputInterfaceName    string                      `protobuf:"bytes,11,opt,name=output_interface_name,json=outputInterfaceName,proto3" json:"output_interface_name,omitempty"`
	BackupTunnelName       string                      `protobuf:"bytes,12,opt,name=backup_tunnel_name,json=backupTunnelName,proto3" json:"backup_tunnel_name,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                    `json:"-"`
	XXX_unrecognized       []byte                      `json:"-"`
	XXX_sizecache          int32                       `json:"-"`
}

func (m *TeS2LTunRwFieldsP2PBag) Reset()         { *m = TeS2LTunRwFieldsP2PBag{} }
func (m *TeS2LTunRwFieldsP2PBag) String() string { return proto.CompactTextString(m) }
func (*TeS2LTunRwFieldsP2PBag) ProtoMessage()    {}
func (*TeS2LTunRwFieldsP2PBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d097dad63b57d04, []int{2}
}

func (m *TeS2LTunRwFieldsP2PBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LTunRwFieldsP2PBag.Unmarshal(m, b)
}
func (m *TeS2LTunRwFieldsP2PBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LTunRwFieldsP2PBag.Marshal(b, m, deterministic)
}
func (m *TeS2LTunRwFieldsP2PBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LTunRwFieldsP2PBag.Merge(m, src)
}
func (m *TeS2LTunRwFieldsP2PBag) XXX_Size() int {
	return xxx_messageInfo_TeS2LTunRwFieldsP2PBag.Size(m)
}
func (m *TeS2LTunRwFieldsP2PBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LTunRwFieldsP2PBag.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LTunRwFieldsP2PBag proto.InternalMessageInfo

func (m *TeS2LTunRwFieldsP2PBag) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

func (m *TeS2LTunRwFieldsP2PBag) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *TeS2LTunRwFieldsP2PBag) GetSourcePe() string {
	if m != nil {
		return m.SourcePe
	}
	return ""
}

func (m *TeS2LTunRwFieldsP2PBag) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *TeS2LTunRwFieldsP2PBag) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *TeS2LTunRwFieldsP2PBag) GetPrimaryS2L() *TeS2LOutputRwOutputInfoBag {
	if m != nil {
		return m.PrimaryS2L
	}
	return nil
}

func (m *TeS2LTunRwFieldsP2PBag) GetBackupTunnelRewrite() *TeS2LOutputRwOutputInfoBag {
	if m != nil {
		return m.BackupTunnelRewrite
	}
	return nil
}

func (m *TeS2LTunRwFieldsP2PBag) GetBackupActive() bool {
	if m != nil {
		return m.BackupActive
	}
	return false
}

func (m *TeS2LTunRwFieldsP2PBag) GetOriginalInputInterface() string {
	if m != nil {
		return m.OriginalInputInterface
	}
	return ""
}

func (m *TeS2LTunRwFieldsP2PBag) GetPreviousHopAddress() string {
	if m != nil {
		return m.PreviousHopAddress
	}
	return ""
}

func (m *TeS2LTunRwFieldsP2PBag) GetOutputInterfaceName() string {
	if m != nil {
		return m.OutputInterfaceName
	}
	return ""
}

func (m *TeS2LTunRwFieldsP2PBag) GetBackupTunnelName() string {
	if m != nil {
		return m.BackupTunnelName
	}
	return ""
}

type TeS2LTunRwFieldsP2MpBag struct {
	LocalLabel           uint32   `protobuf:"varint,1,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeS2LTunRwFieldsP2MpBag) Reset()         { *m = TeS2LTunRwFieldsP2MpBag{} }
func (m *TeS2LTunRwFieldsP2MpBag) String() string { return proto.CompactTextString(m) }
func (*TeS2LTunRwFieldsP2MpBag) ProtoMessage()    {}
func (*TeS2LTunRwFieldsP2MpBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d097dad63b57d04, []int{3}
}

func (m *TeS2LTunRwFieldsP2MpBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LTunRwFieldsP2MpBag.Unmarshal(m, b)
}
func (m *TeS2LTunRwFieldsP2MpBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LTunRwFieldsP2MpBag.Marshal(b, m, deterministic)
}
func (m *TeS2LTunRwFieldsP2MpBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LTunRwFieldsP2MpBag.Merge(m, src)
}
func (m *TeS2LTunRwFieldsP2MpBag) XXX_Size() int {
	return xxx_messageInfo_TeS2LTunRwFieldsP2MpBag.Size(m)
}
func (m *TeS2LTunRwFieldsP2MpBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LTunRwFieldsP2MpBag.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LTunRwFieldsP2MpBag proto.InternalMessageInfo

func (m *TeS2LTunRwFieldsP2MpBag) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

type TeS2LTunRwFieldCtypeBag struct {
	FecCtype             string                   `protobuf:"bytes,1,opt,name=fec_ctype,json=fecCtype,proto3" json:"fec_ctype,omitempty"`
	P2PTunnelRewrite     *TeS2LTunRwFieldsP2PBag  `protobuf:"bytes,2,opt,name=p2p_tunnel_rewrite,json=p2pTunnelRewrite,proto3" json:"p2p_tunnel_rewrite,omitempty"`
	P2MpTunnelRewrite    *TeS2LTunRwFieldsP2MpBag `protobuf:"bytes,3,opt,name=p2mp_tunnel_rewrite,json=p2mpTunnelRewrite,proto3" json:"p2mp_tunnel_rewrite,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TeS2LTunRwFieldCtypeBag) Reset()         { *m = TeS2LTunRwFieldCtypeBag{} }
func (m *TeS2LTunRwFieldCtypeBag) String() string { return proto.CompactTextString(m) }
func (*TeS2LTunRwFieldCtypeBag) ProtoMessage()    {}
func (*TeS2LTunRwFieldCtypeBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d097dad63b57d04, []int{4}
}

func (m *TeS2LTunRwFieldCtypeBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LTunRwFieldCtypeBag.Unmarshal(m, b)
}
func (m *TeS2LTunRwFieldCtypeBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LTunRwFieldCtypeBag.Marshal(b, m, deterministic)
}
func (m *TeS2LTunRwFieldCtypeBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LTunRwFieldCtypeBag.Merge(m, src)
}
func (m *TeS2LTunRwFieldCtypeBag) XXX_Size() int {
	return xxx_messageInfo_TeS2LTunRwFieldCtypeBag.Size(m)
}
func (m *TeS2LTunRwFieldCtypeBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LTunRwFieldCtypeBag.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LTunRwFieldCtypeBag proto.InternalMessageInfo

func (m *TeS2LTunRwFieldCtypeBag) GetFecCtype() string {
	if m != nil {
		return m.FecCtype
	}
	return ""
}

func (m *TeS2LTunRwFieldCtypeBag) GetP2PTunnelRewrite() *TeS2LTunRwFieldsP2PBag {
	if m != nil {
		return m.P2PTunnelRewrite
	}
	return nil
}

func (m *TeS2LTunRwFieldCtypeBag) GetP2MpTunnelRewrite() *TeS2LTunRwFieldsP2MpBag {
	if m != nil {
		return m.P2MpTunnelRewrite
	}
	return nil
}

type TeS2LTunRwFieldsBag struct {
	Timestamp            uint32                   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TunnelRewrite        *TeS2LTunRwFieldCtypeBag `protobuf:"bytes,2,opt,name=tunnel_rewrite,json=tunnelRewrite,proto3" json:"tunnel_rewrite,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TeS2LTunRwFieldsBag) Reset()         { *m = TeS2LTunRwFieldsBag{} }
func (m *TeS2LTunRwFieldsBag) String() string { return proto.CompactTextString(m) }
func (*TeS2LTunRwFieldsBag) ProtoMessage()    {}
func (*TeS2LTunRwFieldsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d097dad63b57d04, []int{5}
}

func (m *TeS2LTunRwFieldsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LTunRwFieldsBag.Unmarshal(m, b)
}
func (m *TeS2LTunRwFieldsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LTunRwFieldsBag.Marshal(b, m, deterministic)
}
func (m *TeS2LTunRwFieldsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LTunRwFieldsBag.Merge(m, src)
}
func (m *TeS2LTunRwFieldsBag) XXX_Size() int {
	return xxx_messageInfo_TeS2LTunRwFieldsBag.Size(m)
}
func (m *TeS2LTunRwFieldsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LTunRwFieldsBag.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LTunRwFieldsBag proto.InternalMessageInfo

func (m *TeS2LTunRwFieldsBag) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TeS2LTunRwFieldsBag) GetTunnelRewrite() *TeS2LTunRwFieldCtypeBag {
	if m != nil {
		return m.TunnelRewrite
	}
	return nil
}

type MplsTeS2LTunRwBag struct {
	TunnelInterface      string               `protobuf:"bytes,50,opt,name=tunnel_interface,json=tunnelInterface,proto3" json:"tunnel_interface,omitempty"`
	TunnelSignalName     string               `protobuf:"bytes,51,opt,name=tunnel_signal_name,json=tunnelSignalName,proto3" json:"tunnel_signal_name,omitempty"`
	SuccessfulRewrite    *TeS2LTunRwFieldsBag `protobuf:"bytes,52,opt,name=successful_rewrite,json=successfulRewrite,proto3" json:"successful_rewrite,omitempty"`
	FailedRewrite        *TeS2LTunRwFieldsBag `protobuf:"bytes,53,opt,name=failed_rewrite,json=failedRewrite,proto3" json:"failed_rewrite,omitempty"`
	PendingRewrite       *TeS2LTunRwFieldsBag `protobuf:"bytes,54,opt,name=pending_rewrite,json=pendingRewrite,proto3" json:"pending_rewrite,omitempty"`
	IsSegmentRouting     bool                 `protobuf:"varint,55,opt,name=is_segment_routing,json=isSegmentRouting,proto3" json:"is_segment_routing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MplsTeS2LTunRwBag) Reset()         { *m = MplsTeS2LTunRwBag{} }
func (m *MplsTeS2LTunRwBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeS2LTunRwBag) ProtoMessage()    {}
func (*MplsTeS2LTunRwBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d097dad63b57d04, []int{6}
}

func (m *MplsTeS2LTunRwBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeS2LTunRwBag.Unmarshal(m, b)
}
func (m *MplsTeS2LTunRwBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeS2LTunRwBag.Marshal(b, m, deterministic)
}
func (m *MplsTeS2LTunRwBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeS2LTunRwBag.Merge(m, src)
}
func (m *MplsTeS2LTunRwBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeS2LTunRwBag.Size(m)
}
func (m *MplsTeS2LTunRwBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeS2LTunRwBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeS2LTunRwBag proto.InternalMessageInfo

func (m *MplsTeS2LTunRwBag) GetTunnelInterface() string {
	if m != nil {
		return m.TunnelInterface
	}
	return ""
}

func (m *MplsTeS2LTunRwBag) GetTunnelSignalName() string {
	if m != nil {
		return m.TunnelSignalName
	}
	return ""
}

func (m *MplsTeS2LTunRwBag) GetSuccessfulRewrite() *TeS2LTunRwFieldsBag {
	if m != nil {
		return m.SuccessfulRewrite
	}
	return nil
}

func (m *MplsTeS2LTunRwBag) GetFailedRewrite() *TeS2LTunRwFieldsBag {
	if m != nil {
		return m.FailedRewrite
	}
	return nil
}

func (m *MplsTeS2LTunRwBag) GetPendingRewrite() *TeS2LTunRwFieldsBag {
	if m != nil {
		return m.PendingRewrite
	}
	return nil
}

func (m *MplsTeS2LTunRwBag) GetIsSegmentRouting() bool {
	if m != nil {
		return m.IsSegmentRouting
	}
	return false
}

func init() {
	proto.RegisterType((*MplsTeS2LTunRwBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.p2mp.tunnel_rewrites.tunnel_rewrite.mpls_te_s2l_tun_rw_bag_KEYS")
	proto.RegisterType((*TeS2LOutputRwOutputInfoBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.p2mp.tunnel_rewrites.tunnel_rewrite.te_s2l_output_rw_output_info_bag")
	proto.RegisterType((*TeS2LTunRwFieldsP2PBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.p2mp.tunnel_rewrites.tunnel_rewrite.te_s2l_tun_rw_fields_p2p_bag")
	proto.RegisterType((*TeS2LTunRwFieldsP2MpBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.p2mp.tunnel_rewrites.tunnel_rewrite.te_s2l_tun_rw_fields_p2mp_bag")
	proto.RegisterType((*TeS2LTunRwFieldCtypeBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.p2mp.tunnel_rewrites.tunnel_rewrite.te_s2l_tun_rw_field_ctype_bag")
	proto.RegisterType((*TeS2LTunRwFieldsBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.p2mp.tunnel_rewrites.tunnel_rewrite.te_s2l_tun_rw_fields_bag")
	proto.RegisterType((*MplsTeS2LTunRwBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.p2mp.tunnel_rewrites.tunnel_rewrite.mpls_te_s2l_tun_rw_bag")
}

func init() { proto.RegisterFile("mpls_te_s2l_tun_rw_bag.proto", fileDescriptor_2d097dad63b57d04) }

var fileDescriptor_2d097dad63b57d04 = []byte{
	// 825 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x8a, 0xeb, 0x36,
	0x14, 0xc6, 0x99, 0x9b, 0x34, 0x39, 0x99, 0xe4, 0xa6, 0xca, 0x64, 0x30, 0x64, 0x4a, 0x43, 0xda,
	0x45, 0x0a, 0x43, 0x5a, 0x7c, 0x6f, 0x7f, 0x96, 0xbd, 0x94, 0x81, 0x86, 0x96, 0x52, 0x9c, 0x6e,
	0xba, 0x12, 0x8a, 0xad, 0x78, 0xc4, 0xd8, 0x96, 0x90, 0xe4, 0xf9, 0x79, 0x81, 0x32, 0x50, 0x28,
	0x14, 0xba, 0xeb, 0x33, 0xf4, 0x2d, 0x4a, 0x5f, 0xa5, 0x2f, 0xd0, 0x07, 0x28, 0x92, 0xec, 0x64,
	0x6c, 0x52, 0x7a, 0x17, 0x33, 0xb3, 0xb3, 0xce, 0xa7, 0xa3, 0xef, 0xd3, 0x39, 0xdf, 0xb1, 0x0d,
	0x67, 0x99, 0x48, 0x15, 0xd6, 0x14, 0xab, 0x20, 0xc5, 0xba, 0xc8, 0xb1, 0xbc, 0xc1, 0x1b, 0x92,
	0x2c, 0x85, 0xe4, 0x9a, 0xa3, 0x8b, 0x88, 0xa9, 0x88, 0x63, 0xc6, 0x15, 0xbe, 0x95, 0xb8, 0xda,
	0xca, 0x05, 0x95, 0xcb, 0x72, 0xb1, 0x14, 0x41, 0x26, 0x96, 0xba, 0xc8, 0x73, 0x9a, 0x62, 0x49,
	0x6f, 0x24, 0xd3, 0x54, 0x35, 0xd6, 0xf3, 0x04, 0xa6, 0x87, 0x69, 0xf0, 0x37, 0x17, 0x3f, 0xae,
	0xd1, 0x14, 0x7a, 0x65, 0x02, 0x8b, 0x7d, 0x6f, 0xe6, 0x2d, 0x06, 0x61, 0xd7, 0x05, 0x56, 0x31,
	0x9a, 0x40, 0x27, 0x55, 0xc2, 0x20, 0x2d, 0x8b, 0xb4, 0x53, 0x25, 0x56, 0x31, 0x3a, 0x81, 0x76,
	0xa4, 0xef, 0x04, 0xf5, 0x8f, 0x66, 0xde, 0xa2, 0x17, 0xba, 0xc5, 0xfc, 0x8f, 0x16, 0xcc, 0x4a,
	0x12, 0x5e, 0x68, 0x51, 0x68, 0xc3, 0x53, 0x3e, 0xb1, 0x7c, 0xcb, 0x0d, 0x27, 0x7a, 0x0d, 0xa7,
	0xe2, 0xf2, 0x4e, 0xb1, 0x88, 0x60, 0x96, 0x6b, 0x2a, 0xb7, 0x24, 0xa2, 0x38, 0x27, 0x19, 0xb5,
	0xdc, 0xbd, 0xf0, 0xa4, 0x44, 0x57, 0x15, 0xf8, 0x1d, 0xc9, 0x28, 0x0a, 0x60, 0x52, 0x89, 0xac,
	0x27, 0xb5, 0x6c, 0xd2, 0xb8, 0x14, 0xdc, 0xcc, 0x11, 0x44, 0xd2, 0x5c, 0x37, 0x73, 0x9c, 0xe8,
	0xb1, 0x03, 0xeb, 0x39, 0x0b, 0x18, 0xe5, 0xf4, 0x56, 0xe3, 0x4b, 0x2e, 0x30, 0x89, 0x63, 0x49,
	0x95, 0xf2, 0x5f, 0xd8, 0xed, 0x43, 0x13, 0xff, 0x9a, 0x8b, 0x37, 0x2e, 0x6a, 0xca, 0xc6, 0x0b,
	0x8d, 0x53, 0xb2, 0xa1, 0xa9, 0xdf, 0x76, 0x65, 0xe3, 0x85, 0xfe, 0xd6, 0xac, 0xd1, 0x87, 0x30,
	0x54, 0xd2, 0x61, 0x58, 0x69, 0x12, 0x5d, 0xf9, 0x9d, 0xd9, 0xd1, 0x62, 0x10, 0x1e, 0x2b, 0x69,
	0x37, 0xac, 0x4d, 0x6c, 0xfe, 0x4f, 0x1b, 0xce, 0xea, 0x4d, 0xd9, 0x32, 0x9a, 0xc6, 0x0a, 0x8b,
	0x40, 0xd8, 0x5a, 0xbd, 0x0f, 0xfd, 0x94, 0x47, 0x24, 0x2d, 0x59, 0x5c, 0x73, 0xc0, 0x86, 0x1c,
	0x4f, 0xad, 0x77, 0xad, 0x46, 0xef, 0xa6, 0xd0, 0x53, 0xbc, 0x90, 0x11, 0xc5, 0xbb, 0x46, 0x75,
	0x5d, 0xe0, 0x7b, 0x8a, 0x3e, 0x86, 0x71, 0x4c, 0x95, 0x66, 0x39, 0xd1, 0x8c, 0xe7, 0x8d, 0xbb,
	0xa2, 0x07, 0x50, 0x75, 0xdf, 0x73, 0x40, 0xf4, 0x56, 0xd3, 0x3c, 0xa6, 0x31, 0xde, 0x73, 0xb6,
	0xed, 0xfe, 0x51, 0x85, 0xfc, 0x50, 0x71, 0xdf, 0x7b, 0xd0, 0x17, 0x92, 0x65, 0x44, 0xde, 0x99,
	0xfb, 0xf9, 0x9d, 0x99, 0xb7, 0xe8, 0x07, 0xc9, 0xf2, 0x51, 0x1c, 0xbd, 0xfc, 0x3f, 0x93, 0x85,
	0x50, 0x72, 0xaf, 0x83, 0x14, 0xfd, 0xee, 0xc1, 0x64, 0x43, 0xa2, 0xab, 0x42, 0xe0, 0xfa, 0x31,
	0xfe, 0x3b, 0xcf, 0x2b, 0x6a, 0xec, 0x54, 0xb8, 0x1a, 0x85, 0x2e, 0x0b, 0x7d, 0x00, 0x83, 0x52,
	0x1c, 0x89, 0x34, 0xbb, 0xa6, 0x7e, 0x77, 0xe6, 0x2d, 0xba, 0xe1, 0xb1, 0x0b, 0xbe, 0xb1, 0x31,
	0xf4, 0x05, 0xf8, 0x5c, 0xb2, 0x84, 0xe5, 0xc4, 0xf8, 0xdf, 0x1d, 0x5a, 0xba, 0xd6, 0xef, 0xd9,
	0x0e, 0x9c, 0x56, 0xf8, 0xca, 0xc0, 0x3b, 0x4f, 0xa3, 0x4f, 0xe0, 0x44, 0x48, 0x7a, 0xcd, 0x78,
	0xa1, 0x6a, 0x9e, 0x06, 0xd7, 0xe7, 0x0a, 0x7b, 0xe0, 0xeb, 0x00, 0x26, 0x3b, 0xe1, 0xb5, 0xa9,
	0xe9, 0xbb, 0xa9, 0x71, 0x60, 0x7d, 0x6a, 0xce, 0x01, 0xd5, 0x2b, 0x6c, 0x13, 0x8e, 0x9d, 0x37,
	0x1e, 0xde, 0xda, 0xec, 0x9e, 0x7f, 0x09, 0xef, 0xfd, 0x87, 0xeb, 0xb3, 0xb7, 0xb3, 0xfd, 0xfc,
	0xef, 0xd6, 0xc1, 0x23, 0xb0, 0x7d, 0x0d, 0xd9, 0x23, 0xa6, 0xd0, 0xdb, 0xd2, 0xc8, 0x05, 0xca,
	0x17, 0x4b, 0x77, 0x4b, 0xa3, 0xaf, 0xcc, 0x1a, 0xfd, 0xea, 0x01, 0x32, 0x23, 0xd6, 0xb0, 0x43,
	0xcb, 0xda, 0x21, 0x7a, 0x5c, 0x3b, 0x1c, 0x1c, 0xec, 0x70, 0x24, 0x82, 0x86, 0x0f, 0x7e, 0xf3,
	0x60, 0x6c, 0x0b, 0xd0, 0x10, 0x75, 0x64, 0x45, 0xc5, 0x4f, 0x2b, 0xca, 0xd5, 0x3d, 0x7c, 0xd7,
	0x3c, 0xd5, 0x64, 0xcd, 0xff, 0xf2, 0xc0, 0x3f, 0x98, 0x64, 0x8a, 0x7c, 0x06, 0x3d, 0xcd, 0x32,
	0xaa, 0x34, 0xc9, 0x44, 0xd9, 0xa5, 0x7d, 0x00, 0xfd, 0xec, 0xc1, 0xf0, 0x60, 0x85, 0x9f, 0xf0,
	0x32, 0x7b, 0x07, 0x84, 0x03, 0x5d, 0xbb, 0xc8, 0x9f, 0x2f, 0xe0, 0xf4, 0xf0, 0x57, 0x10, 0x7d,
	0x04, 0xa3, 0xe6, 0xb7, 0xc5, 0x0f, 0xac, 0x65, 0x5e, 0x36, 0x3e, 0x2b, 0xc6, 0xe8, 0xe5, 0x56,
	0xc5, 0x12, 0x33, 0x8d, 0xd6, 0xe8, 0xaf, 0x9c, 0xd1, 0x1d, 0xb2, 0xb6, 0x80, 0x1d, 0x8b, 0x5f,
	0x3c, 0x40, 0xaa, 0x88, 0x22, 0xaa, 0xd4, 0xb6, 0xd8, 0x57, 0xe1, 0xb5, 0xad, 0x02, 0x7e, 0xca,
	0x96, 0xda, 0x6e, 0xee, 0xa9, 0x2b, 0x93, 0xfd, 0xe4, 0xc1, 0x70, 0x4b, 0x58, 0x4a, 0xe3, 0x9d,
	0x98, 0x4f, 0x9f, 0x47, 0xcc, 0xc0, 0xd1, 0x56, 0x42, 0xee, 0x3d, 0x78, 0x29, 0x68, 0x1e, 0xb3,
	0x3c, 0xd9, 0x29, 0xf9, 0xec, 0x79, 0x94, 0x0c, 0x4b, 0xde, 0x4a, 0xca, 0x39, 0x20, 0xa6, 0xb0,
	0xa2, 0x49, 0x66, 0xfe, 0x14, 0x24, 0x2f, 0x34, 0xcb, 0x13, 0xff, 0x73, 0xfb, 0x16, 0x1e, 0x31,
	0xb5, 0x76, 0x40, 0xe8, 0xe2, 0x9b, 0x8e, 0xfd, 0x33, 0x7b, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xd7, 0x82, 0x21, 0xdf, 0xb9, 0x09, 0x00, 0x00,
}
