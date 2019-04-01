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
// source: ntp_edm_assoc_detail_result.proto

package cisco_ios_xr_ip_ntp_admin_oper_ntp_racks_rack_slots_slot_instances_instance_associations_detail

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

type NtpEdmAssocDetailResult_KEYS struct {
	Number               uint32   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Number_1             uint32   `protobuf:"varint,2,opt,name=number_1,json=number1,proto3" json:"number_1,omitempty"`
	Number_2             uint32   `protobuf:"varint,3,opt,name=number_2,json=number2,proto3" json:"number_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NtpEdmAssocDetailResult_KEYS) Reset()         { *m = NtpEdmAssocDetailResult_KEYS{} }
func (m *NtpEdmAssocDetailResult_KEYS) String() string { return proto.CompactTextString(m) }
func (*NtpEdmAssocDetailResult_KEYS) ProtoMessage()    {}
func (*NtpEdmAssocDetailResult_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_391b9e658e229961, []int{0}
}

func (m *NtpEdmAssocDetailResult_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NtpEdmAssocDetailResult_KEYS.Unmarshal(m, b)
}
func (m *NtpEdmAssocDetailResult_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NtpEdmAssocDetailResult_KEYS.Marshal(b, m, deterministic)
}
func (m *NtpEdmAssocDetailResult_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NtpEdmAssocDetailResult_KEYS.Merge(m, src)
}
func (m *NtpEdmAssocDetailResult_KEYS) XXX_Size() int {
	return xxx_messageInfo_NtpEdmAssocDetailResult_KEYS.Size(m)
}
func (m *NtpEdmAssocDetailResult_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_NtpEdmAssocDetailResult_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_NtpEdmAssocDetailResult_KEYS proto.InternalMessageInfo

func (m *NtpEdmAssocDetailResult_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *NtpEdmAssocDetailResult_KEYS) GetNumber_1() uint32 {
	if m != nil {
		return m.Number_1
	}
	return 0
}

func (m *NtpEdmAssocDetailResult_KEYS) GetNumber_2() uint32 {
	if m != nil {
		return m.Number_2
	}
	return 0
}

type NtpEdmPeerInfoCommon struct {
	HostMode             string   `protobuf:"bytes,1,opt,name=host_mode,json=hostMode,proto3" json:"host_mode,omitempty"`
	IsConfigured         bool     `protobuf:"varint,2,opt,name=is_configured,json=isConfigured,proto3" json:"is_configured,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ReferenceId          string   `protobuf:"bytes,4,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	HostPoll             uint32   `protobuf:"varint,5,opt,name=host_poll,json=hostPoll,proto3" json:"host_poll,omitempty"`
	Reachability         uint32   `protobuf:"varint,6,opt,name=reachability,proto3" json:"reachability,omitempty"`
	Stratum              uint32   `protobuf:"varint,7,opt,name=stratum,proto3" json:"stratum,omitempty"`
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Delay                string   `protobuf:"bytes,9,opt,name=delay,proto3" json:"delay,omitempty"`
	Offset               string   `protobuf:"bytes,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Dispersion           string   `protobuf:"bytes,11,opt,name=dispersion,proto3" json:"dispersion,omitempty"`
	IsSysPeer            bool     `protobuf:"varint,12,opt,name=is_sys_peer,json=isSysPeer,proto3" json:"is_sys_peer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NtpEdmPeerInfoCommon) Reset()         { *m = NtpEdmPeerInfoCommon{} }
func (m *NtpEdmPeerInfoCommon) String() string { return proto.CompactTextString(m) }
func (*NtpEdmPeerInfoCommon) ProtoMessage()    {}
func (*NtpEdmPeerInfoCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_391b9e658e229961, []int{1}
}

func (m *NtpEdmPeerInfoCommon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NtpEdmPeerInfoCommon.Unmarshal(m, b)
}
func (m *NtpEdmPeerInfoCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NtpEdmPeerInfoCommon.Marshal(b, m, deterministic)
}
func (m *NtpEdmPeerInfoCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NtpEdmPeerInfoCommon.Merge(m, src)
}
func (m *NtpEdmPeerInfoCommon) XXX_Size() int {
	return xxx_messageInfo_NtpEdmPeerInfoCommon.Size(m)
}
func (m *NtpEdmPeerInfoCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_NtpEdmPeerInfoCommon.DiscardUnknown(m)
}

var xxx_messageInfo_NtpEdmPeerInfoCommon proto.InternalMessageInfo

func (m *NtpEdmPeerInfoCommon) GetHostMode() string {
	if m != nil {
		return m.HostMode
	}
	return ""
}

func (m *NtpEdmPeerInfoCommon) GetIsConfigured() bool {
	if m != nil {
		return m.IsConfigured
	}
	return false
}

func (m *NtpEdmPeerInfoCommon) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NtpEdmPeerInfoCommon) GetReferenceId() string {
	if m != nil {
		return m.ReferenceId
	}
	return ""
}

func (m *NtpEdmPeerInfoCommon) GetHostPoll() uint32 {
	if m != nil {
		return m.HostPoll
	}
	return 0
}

func (m *NtpEdmPeerInfoCommon) GetReachability() uint32 {
	if m != nil {
		return m.Reachability
	}
	return 0
}

func (m *NtpEdmPeerInfoCommon) GetStratum() uint32 {
	if m != nil {
		return m.Stratum
	}
	return 0
}

func (m *NtpEdmPeerInfoCommon) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *NtpEdmPeerInfoCommon) GetDelay() string {
	if m != nil {
		return m.Delay
	}
	return ""
}

func (m *NtpEdmPeerInfoCommon) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *NtpEdmPeerInfoCommon) GetDispersion() string {
	if m != nil {
		return m.Dispersion
	}
	return ""
}

func (m *NtpEdmPeerInfoCommon) GetIsSysPeer() bool {
	if m != nil {
		return m.IsSysPeer
	}
	return false
}

type EdmUlI struct {
	Int                  uint32   `protobuf:"varint,1,opt,name=int,proto3" json:"int,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdmUlI) Reset()         { *m = EdmUlI{} }
func (m *EdmUlI) String() string { return proto.CompactTextString(m) }
func (*EdmUlI) ProtoMessage()    {}
func (*EdmUlI) Descriptor() ([]byte, []int) {
	return fileDescriptor_391b9e658e229961, []int{2}
}

func (m *EdmUlI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdmUlI.Unmarshal(m, b)
}
func (m *EdmUlI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdmUlI.Marshal(b, m, deterministic)
}
func (m *EdmUlI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdmUlI.Merge(m, src)
}
func (m *EdmUlI) XXX_Size() int {
	return xxx_messageInfo_EdmUlI.Size(m)
}
func (m *EdmUlI) XXX_DiscardUnknown() {
	xxx_messageInfo_EdmUlI.DiscardUnknown(m)
}

var xxx_messageInfo_EdmUlI proto.InternalMessageInfo

func (m *EdmUlI) GetInt() uint32 {
	if m != nil {
		return m.Int
	}
	return 0
}

type EdmUlF struct {
	Frac                 uint32   `protobuf:"varint,1,opt,name=frac,proto3" json:"frac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdmUlF) Reset()         { *m = EdmUlF{} }
func (m *EdmUlF) String() string { return proto.CompactTextString(m) }
func (*EdmUlF) ProtoMessage()    {}
func (*EdmUlF) Descriptor() ([]byte, []int) {
	return fileDescriptor_391b9e658e229961, []int{3}
}

func (m *EdmUlF) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdmUlF.Unmarshal(m, b)
}
func (m *EdmUlF) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdmUlF.Marshal(b, m, deterministic)
}
func (m *EdmUlF) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdmUlF.Merge(m, src)
}
func (m *EdmUlF) XXX_Size() int {
	return xxx_messageInfo_EdmUlF.Size(m)
}
func (m *EdmUlF) XXX_DiscardUnknown() {
	xxx_messageInfo_EdmUlF.DiscardUnknown(m)
}

var xxx_messageInfo_EdmUlF proto.InternalMessageInfo

func (m *EdmUlF) GetFrac() uint32 {
	if m != nil {
		return m.Frac
	}
	return 0
}

type EdmLFp struct {
	Sec                  *EdmUlI  `protobuf:"bytes,1,opt,name=sec,proto3" json:"sec,omitempty"`
	FracSecs             *EdmUlF  `protobuf:"bytes,2,opt,name=frac_secs,json=fracSecs,proto3" json:"frac_secs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdmLFp) Reset()         { *m = EdmLFp{} }
func (m *EdmLFp) String() string { return proto.CompactTextString(m) }
func (*EdmLFp) ProtoMessage()    {}
func (*EdmLFp) Descriptor() ([]byte, []int) {
	return fileDescriptor_391b9e658e229961, []int{4}
}

func (m *EdmLFp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdmLFp.Unmarshal(m, b)
}
func (m *EdmLFp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdmLFp.Marshal(b, m, deterministic)
}
func (m *EdmLFp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdmLFp.Merge(m, src)
}
func (m *EdmLFp) XXX_Size() int {
	return xxx_messageInfo_EdmLFp.Size(m)
}
func (m *EdmLFp) XXX_DiscardUnknown() {
	xxx_messageInfo_EdmLFp.DiscardUnknown(m)
}

var xxx_messageInfo_EdmLFp proto.InternalMessageInfo

func (m *EdmLFp) GetSec() *EdmUlI {
	if m != nil {
		return m.Sec
	}
	return nil
}

func (m *EdmLFp) GetFracSecs() *EdmUlF {
	if m != nil {
		return m.FracSecs
	}
	return nil
}

type FilterDetails struct {
	FilterDelay          string   `protobuf:"bytes,1,opt,name=filter_delay,json=filterDelay,proto3" json:"filter_delay,omitempty"`
	FilterOffset         string   `protobuf:"bytes,2,opt,name=filter_offset,json=filterOffset,proto3" json:"filter_offset,omitempty"`
	FilterDisp           string   `protobuf:"bytes,3,opt,name=filter_disp,json=filterDisp,proto3" json:"filter_disp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterDetails) Reset()         { *m = FilterDetails{} }
func (m *FilterDetails) String() string { return proto.CompactTextString(m) }
func (*FilterDetails) ProtoMessage()    {}
func (*FilterDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_391b9e658e229961, []int{5}
}

func (m *FilterDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterDetails.Unmarshal(m, b)
}
func (m *FilterDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterDetails.Marshal(b, m, deterministic)
}
func (m *FilterDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterDetails.Merge(m, src)
}
func (m *FilterDetails) XXX_Size() int {
	return xxx_messageInfo_FilterDetails.Size(m)
}
func (m *FilterDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterDetails.DiscardUnknown(m)
}

var xxx_messageInfo_FilterDetails proto.InternalMessageInfo

func (m *FilterDetails) GetFilterDelay() string {
	if m != nil {
		return m.FilterDelay
	}
	return ""
}

func (m *FilterDetails) GetFilterOffset() string {
	if m != nil {
		return m.FilterOffset
	}
	return ""
}

func (m *FilterDetails) GetFilterDisp() string {
	if m != nil {
		return m.FilterDisp
	}
	return ""
}

type NtpEdmPeerDetailInfo struct {
	PeerInfoCommon       *NtpEdmPeerInfoCommon `protobuf:"bytes,1,opt,name=peer_info_common,json=peerInfoCommon,proto3" json:"peer_info_common,omitempty"`
	Leap                 string                `protobuf:"bytes,2,opt,name=leap,proto3" json:"leap,omitempty"`
	PeerMode             string                `protobuf:"bytes,3,opt,name=peer_mode,json=peerMode,proto3" json:"peer_mode,omitempty"`
	PollInterval         uint32                `protobuf:"varint,4,opt,name=poll_interval,json=pollInterval,proto3" json:"poll_interval,omitempty"`
	IsRefClock           bool                  `protobuf:"varint,5,opt,name=is_ref_clock,json=isRefClock,proto3" json:"is_ref_clock,omitempty"`
	IsAuthenticated      bool                  `protobuf:"varint,6,opt,name=is_authenticated,json=isAuthenticated,proto3" json:"is_authenticated,omitempty"`
	RootDelay            string                `protobuf:"bytes,7,opt,name=root_delay,json=rootDelay,proto3" json:"root_delay,omitempty"`
	RootDispersion       string                `protobuf:"bytes,8,opt,name=root_dispersion,json=rootDispersion,proto3" json:"root_dispersion,omitempty"`
	SynchDistance        string                `protobuf:"bytes,9,opt,name=synch_distance,json=synchDistance,proto3" json:"synch_distance,omitempty"`
	RefTime              *EdmLFp               `protobuf:"bytes,10,opt,name=ref_time,json=refTime,proto3" json:"ref_time,omitempty"`
	Precision            string                `protobuf:"bytes,11,opt,name=precision,proto3" json:"precision,omitempty"`
	Version              uint32                `protobuf:"varint,12,opt,name=version,proto3" json:"version,omitempty"`
	OriginateTime        *EdmLFp               `protobuf:"bytes,13,opt,name=originate_time,json=originateTime,proto3" json:"originate_time,omitempty"`
	ReceiveTime          *EdmLFp               `protobuf:"bytes,14,opt,name=receive_time,json=receiveTime,proto3" json:"receive_time,omitempty"`
	TransmitTime         *EdmLFp               `protobuf:"bytes,15,opt,name=transmit_time,json=transmitTime,proto3" json:"transmit_time,omitempty"`
	FilterIndex          uint32                `protobuf:"varint,16,opt,name=filter_index,json=filterIndex,proto3" json:"filter_index,omitempty"`
	FilterDetail         []*FilterDetails      `protobuf:"bytes,17,rep,name=filter_detail,json=filterDetail,proto3" json:"filter_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NtpEdmPeerDetailInfo) Reset()         { *m = NtpEdmPeerDetailInfo{} }
func (m *NtpEdmPeerDetailInfo) String() string { return proto.CompactTextString(m) }
func (*NtpEdmPeerDetailInfo) ProtoMessage()    {}
func (*NtpEdmPeerDetailInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_391b9e658e229961, []int{6}
}

func (m *NtpEdmPeerDetailInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NtpEdmPeerDetailInfo.Unmarshal(m, b)
}
func (m *NtpEdmPeerDetailInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NtpEdmPeerDetailInfo.Marshal(b, m, deterministic)
}
func (m *NtpEdmPeerDetailInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NtpEdmPeerDetailInfo.Merge(m, src)
}
func (m *NtpEdmPeerDetailInfo) XXX_Size() int {
	return xxx_messageInfo_NtpEdmPeerDetailInfo.Size(m)
}
func (m *NtpEdmPeerDetailInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NtpEdmPeerDetailInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NtpEdmPeerDetailInfo proto.InternalMessageInfo

func (m *NtpEdmPeerDetailInfo) GetPeerInfoCommon() *NtpEdmPeerInfoCommon {
	if m != nil {
		return m.PeerInfoCommon
	}
	return nil
}

func (m *NtpEdmPeerDetailInfo) GetLeap() string {
	if m != nil {
		return m.Leap
	}
	return ""
}

func (m *NtpEdmPeerDetailInfo) GetPeerMode() string {
	if m != nil {
		return m.PeerMode
	}
	return ""
}

func (m *NtpEdmPeerDetailInfo) GetPollInterval() uint32 {
	if m != nil {
		return m.PollInterval
	}
	return 0
}

func (m *NtpEdmPeerDetailInfo) GetIsRefClock() bool {
	if m != nil {
		return m.IsRefClock
	}
	return false
}

func (m *NtpEdmPeerDetailInfo) GetIsAuthenticated() bool {
	if m != nil {
		return m.IsAuthenticated
	}
	return false
}

func (m *NtpEdmPeerDetailInfo) GetRootDelay() string {
	if m != nil {
		return m.RootDelay
	}
	return ""
}

func (m *NtpEdmPeerDetailInfo) GetRootDispersion() string {
	if m != nil {
		return m.RootDispersion
	}
	return ""
}

func (m *NtpEdmPeerDetailInfo) GetSynchDistance() string {
	if m != nil {
		return m.SynchDistance
	}
	return ""
}

func (m *NtpEdmPeerDetailInfo) GetRefTime() *EdmLFp {
	if m != nil {
		return m.RefTime
	}
	return nil
}

func (m *NtpEdmPeerDetailInfo) GetPrecision() string {
	if m != nil {
		return m.Precision
	}
	return ""
}

func (m *NtpEdmPeerDetailInfo) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *NtpEdmPeerDetailInfo) GetOriginateTime() *EdmLFp {
	if m != nil {
		return m.OriginateTime
	}
	return nil
}

func (m *NtpEdmPeerDetailInfo) GetReceiveTime() *EdmLFp {
	if m != nil {
		return m.ReceiveTime
	}
	return nil
}

func (m *NtpEdmPeerDetailInfo) GetTransmitTime() *EdmLFp {
	if m != nil {
		return m.TransmitTime
	}
	return nil
}

func (m *NtpEdmPeerDetailInfo) GetFilterIndex() uint32 {
	if m != nil {
		return m.FilterIndex
	}
	return 0
}

func (m *NtpEdmPeerDetailInfo) GetFilterDetail() []*FilterDetails {
	if m != nil {
		return m.FilterDetail
	}
	return nil
}

type NtpEdmAssocDetailResult struct {
	IsNtpEnabled         bool                    `protobuf:"varint,50,opt,name=is_ntp_enabled,json=isNtpEnabled,proto3" json:"is_ntp_enabled,omitempty"`
	SysLeap              string                  `protobuf:"bytes,51,opt,name=sys_leap,json=sysLeap,proto3" json:"sys_leap,omitempty"`
	PeerDetailInfo       []*NtpEdmPeerDetailInfo `protobuf:"bytes,52,rep,name=peer_detail_info,json=peerDetailInfo,proto3" json:"peer_detail_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NtpEdmAssocDetailResult) Reset()         { *m = NtpEdmAssocDetailResult{} }
func (m *NtpEdmAssocDetailResult) String() string { return proto.CompactTextString(m) }
func (*NtpEdmAssocDetailResult) ProtoMessage()    {}
func (*NtpEdmAssocDetailResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_391b9e658e229961, []int{7}
}

func (m *NtpEdmAssocDetailResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NtpEdmAssocDetailResult.Unmarshal(m, b)
}
func (m *NtpEdmAssocDetailResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NtpEdmAssocDetailResult.Marshal(b, m, deterministic)
}
func (m *NtpEdmAssocDetailResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NtpEdmAssocDetailResult.Merge(m, src)
}
func (m *NtpEdmAssocDetailResult) XXX_Size() int {
	return xxx_messageInfo_NtpEdmAssocDetailResult.Size(m)
}
func (m *NtpEdmAssocDetailResult) XXX_DiscardUnknown() {
	xxx_messageInfo_NtpEdmAssocDetailResult.DiscardUnknown(m)
}

var xxx_messageInfo_NtpEdmAssocDetailResult proto.InternalMessageInfo

func (m *NtpEdmAssocDetailResult) GetIsNtpEnabled() bool {
	if m != nil {
		return m.IsNtpEnabled
	}
	return false
}

func (m *NtpEdmAssocDetailResult) GetSysLeap() string {
	if m != nil {
		return m.SysLeap
	}
	return ""
}

func (m *NtpEdmAssocDetailResult) GetPeerDetailInfo() []*NtpEdmPeerDetailInfo {
	if m != nil {
		return m.PeerDetailInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*NtpEdmAssocDetailResult_KEYS)(nil), "cisco_ios_xr_ip_ntp_admin_oper.ntp.racks.rack.slots.slot.instances.instance.associations_detail.ntp_edm_assoc_detail_result_KEYS")
	proto.RegisterType((*NtpEdmPeerInfoCommon)(nil), "cisco_ios_xr_ip_ntp_admin_oper.ntp.racks.rack.slots.slot.instances.instance.associations_detail.ntp_edm_peer_info_common")
	proto.RegisterType((*EdmUlI)(nil), "cisco_ios_xr_ip_ntp_admin_oper.ntp.racks.rack.slots.slot.instances.instance.associations_detail.edm_ul_i")
	proto.RegisterType((*EdmUlF)(nil), "cisco_ios_xr_ip_ntp_admin_oper.ntp.racks.rack.slots.slot.instances.instance.associations_detail.edm_ul_f")
	proto.RegisterType((*EdmLFp)(nil), "cisco_ios_xr_ip_ntp_admin_oper.ntp.racks.rack.slots.slot.instances.instance.associations_detail.edm_l_fp")
	proto.RegisterType((*FilterDetails)(nil), "cisco_ios_xr_ip_ntp_admin_oper.ntp.racks.rack.slots.slot.instances.instance.associations_detail.filter_details")
	proto.RegisterType((*NtpEdmPeerDetailInfo)(nil), "cisco_ios_xr_ip_ntp_admin_oper.ntp.racks.rack.slots.slot.instances.instance.associations_detail.ntp_edm_peer_detail_info")
	proto.RegisterType((*NtpEdmAssocDetailResult)(nil), "cisco_ios_xr_ip_ntp_admin_oper.ntp.racks.rack.slots.slot.instances.instance.associations_detail.ntp_edm_assoc_detail_result")
}

func init() { proto.RegisterFile("ntp_edm_assoc_detail_result.proto", fileDescriptor_391b9e658e229961) }

var fileDescriptor_391b9e658e229961 = []byte{
	// 899 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4f, 0x6f, 0x1b, 0xb7,
	0x13, 0x85, 0x64, 0xc7, 0x5a, 0x51, 0x7f, 0xec, 0x1f, 0xf1, 0x43, 0xb1, 0x45, 0x52, 0xd7, 0x51,
	0x5a, 0x34, 0xbd, 0x08, 0x88, 0xd2, 0x2f, 0x50, 0xd8, 0x39, 0x18, 0xfd, 0x17, 0xac, 0x7b, 0xe9,
	0x89, 0xa0, 0x77, 0x67, 0xe3, 0x41, 0x76, 0xc9, 0x05, 0x87, 0x32, 0x22, 0xf4, 0x1a, 0xf4, 0x54,
	0x14, 0x39, 0xf4, 0x1b, 0xf4, 0x5b, 0xf6, 0x50, 0x14, 0x1c, 0x52, 0x6b, 0x25, 0x40, 0x73, 0xb3,
	0x2e, 0x02, 0xf9, 0x86, 0xdc, 0x79, 0xc3, 0xe1, 0x7b, 0xa2, 0x78, 0x6c, 0x7c, 0xa7, 0xa0, 0x6a,
	0x95, 0x26, 0xb2, 0xa5, 0xaa, 0xc0, 0x6b, 0x6c, 0x94, 0x03, 0x5a, 0x37, 0x7e, 0xd9, 0x39, 0xeb,
	0xad, 0x54, 0x25, 0x52, 0x69, 0x15, 0x5a, 0x52, 0x6f, 0x9c, 0xc2, 0x4e, 0x85, 0x2d, 0xba, 0x6a,
	0xd1, 0x28, 0xdb, 0x81, 0x5b, 0x1a, 0xdf, 0x2d, 0x9d, 0x2e, 0x5f, 0x13, 0xff, 0x2e, 0xa9, 0xb1,
	0x9e, 0xf8, 0x77, 0x89, 0x86, 0xbc, 0x36, 0x25, 0x50, 0x3f, 0x5a, 0x72, 0x1a, 0xd4, 0x1e, 0xad,
	0xa1, 0x94, 0x6d, 0xd1, 0x89, 0xb3, 0x8f, 0xb0, 0x50, 0xdf, 0xbd, 0xf8, 0xe5, 0x4a, 0x7e, 0x22,
	0x8e, 0xcc, 0xba, 0xbd, 0x06, 0x97, 0x0f, 0xce, 0x06, 0x4f, 0x67, 0x45, 0x9a, 0xc9, 0x4f, 0x45,
	0x16, 0x47, 0xea, 0x59, 0x3e, 0xe4, 0xc8, 0x28, 0xce, 0x9f, 0xed, 0x84, 0x56, 0xf9, 0xc1, 0x6e,
	0x68, 0xb5, 0xf8, 0x7b, 0x28, 0xf2, 0x6d, 0xca, 0x0e, 0xc0, 0x29, 0x34, 0xb5, 0x55, 0xa5, 0x6d,
	0x5b, 0x6b, 0xe4, 0x43, 0x31, 0xbe, 0xb1, 0xe4, 0x55, 0x6b, 0x2b, 0xe0, 0x6c, 0xe3, 0x22, 0x0b,
	0xc0, 0x0f, 0xb6, 0x02, 0xf9, 0x44, 0xcc, 0x90, 0x54, 0x69, 0x4d, 0x8d, 0xaf, 0xd6, 0x0e, 0x2a,
	0x4e, 0x9a, 0x15, 0x53, 0xa4, 0xf3, 0x1e, 0x93, 0xb9, 0x18, 0xe9, 0xaa, 0x72, 0x40, 0xc4, 0x89,
	0xc7, 0xc5, 0x76, 0x2a, 0x1f, 0x8b, 0xa9, 0x83, 0x1a, 0x1c, 0x98, 0x12, 0x14, 0x56, 0xf9, 0x21,
	0x87, 0x27, 0x3d, 0x76, 0x59, 0xf5, 0xe9, 0x3b, 0xdb, 0x34, 0xf9, 0x03, 0xe6, 0xcd, 0xe9, 0x5f,
	0xda, 0xa6, 0x91, 0x8b, 0xb0, 0x5f, 0x97, 0x37, 0xfa, 0x1a, 0x1b, 0xf4, 0x9b, 0xfc, 0x88, 0xe3,
	0xef, 0x61, 0x21, 0x3b, 0x79, 0xa7, 0xfd, 0xba, 0xcd, 0x47, 0xb1, 0xec, 0x34, 0x0d, 0x87, 0x48,
	0x5e, 0xfb, 0x35, 0xe5, 0x19, 0xe7, 0x4d, 0x33, 0xf9, 0x7f, 0xf1, 0xa0, 0x82, 0x46, 0x6f, 0xf2,
	0x31, 0xc3, 0x71, 0x12, 0x56, 0xdb, 0xba, 0x26, 0xf0, 0xb9, 0x88, 0xab, 0xe3, 0x4c, 0x9e, 0x0a,
	0x51, 0x21, 0x75, 0xe0, 0x08, 0xad, 0xc9, 0x27, 0x1c, 0xdb, 0x41, 0xe4, 0xa9, 0x98, 0x20, 0x29,
	0xda, 0x10, 0x1f, 0x6d, 0x3e, 0xe5, 0x03, 0x1a, 0x23, 0x5d, 0x6d, 0xe8, 0x25, 0x80, 0x5b, 0x3c,
	0x12, 0x59, 0x38, 0xf7, 0x75, 0xa3, 0x50, 0x9e, 0x88, 0x03, 0x34, 0x3e, 0xf5, 0x34, 0x0c, 0x17,
	0xa7, 0x7d, 0xb4, 0x96, 0x52, 0x1c, 0xd6, 0x4e, 0x97, 0x29, 0xcc, 0xe3, 0xc5, 0xbb, 0x61, 0x5c,
	0xd0, 0xa8, 0xba, 0x93, 0xbf, 0x8a, 0x03, 0x82, 0x18, 0x9f, 0xac, 0x70, 0x79, 0xcf, 0x17, 0x75,
	0xb9, 0xa5, 0x5d, 0x84, 0xac, 0xf2, 0xb7, 0x81, 0x18, 0x07, 0x4a, 0x8a, 0xa0, 0x24, 0xbe, 0x07,
	0x7b, 0xe4, 0x50, 0x17, 0x59, 0xc8, 0x7d, 0x05, 0x25, 0x2d, 0x36, 0x62, 0x5e, 0x63, 0xe3, 0xc1,
	0xa5, 0x35, 0x7c, 0xcd, 0x7a, 0x24, 0xf4, 0x35, 0xde, 0xe2, 0x49, 0xc4, 0x2e, 0xb8, 0xbb, 0x4f,
	0xc4, 0x2c, 0x2d, 0x49, 0x4d, 0x1e, 0xf2, 0x9a, 0xb4, 0xef, 0xa7, 0xd8, 0xea, 0xcf, 0xc5, 0x64,
	0xfb, 0x1d, 0xa4, 0x2e, 0x5d, 0x66, 0x91, 0x3e, 0x83, 0xd4, 0x2d, 0xfe, 0x19, 0x7f, 0x20, 0xa4,
	0x24, 0xdd, 0xa0, 0x27, 0xf9, 0xd7, 0x40, 0x9c, 0x7c, 0xa8, 0xae, 0xd4, 0xab, 0xcd, 0xbd, 0x9f,
	0xd3, 0x7f, 0xc9, 0xbb, 0x98, 0x07, 0xe4, 0xd2, 0xd4, 0xf6, 0x3c, 0xca, 0x5d, 0x8a, 0xc3, 0x06,
	0x74, 0x97, 0xea, 0xe7, 0x71, 0xd0, 0x20, 0xef, 0x63, 0x0b, 0x88, 0x55, 0x67, 0x01, 0xd8, 0x5a,
	0x40, 0xd0, 0xa6, 0x42, 0xe3, 0xc1, 0xdd, 0xea, 0x86, 0x45, 0x3c, 0x2b, 0xa6, 0x01, 0xbc, 0x4c,
	0x98, 0x3c, 0x13, 0x53, 0x24, 0xe5, 0xa0, 0x56, 0x65, 0x63, 0xcb, 0xd7, 0x2c, 0xe4, 0xac, 0x10,
	0x48, 0x05, 0xd4, 0xe7, 0x01, 0x91, 0x5f, 0x8b, 0x13, 0x24, 0xa5, 0xd7, 0xfe, 0x06, 0x8c, 0xc7,
	0x52, 0x7b, 0xa8, 0x58, 0xce, 0x59, 0x71, 0x8c, 0xf4, 0xed, 0x2e, 0x2c, 0x3f, 0x13, 0xc2, 0x59,
	0xeb, 0x53, 0x33, 0x47, 0xcc, 0x67, 0x1c, 0x90, 0xd8, 0xca, 0xaf, 0xc4, 0x71, 0x0c, 0xdf, 0xa9,
	0x32, 0xea, 0x7b, 0xce, 0x6b, 0xee, 0x94, 0xf9, 0xa5, 0x98, 0xd3, 0xc6, 0x94, 0x37, 0x61, 0x25,
	0x9f, 0x5c, 0x12, 0xfc, 0x8c, 0xd1, 0x8b, 0x04, 0xca, 0xb7, 0x03, 0x91, 0x05, 0xe6, 0x1e, 0x5b,
	0x60, 0xed, 0xef, 0xeb, 0x5e, 0x07, 0x4d, 0x17, 0x23, 0x07, 0xf5, 0xcf, 0xd8, 0x82, 0x7c, 0x24,
	0xc6, 0x9d, 0x83, 0x12, 0x77, 0x6c, 0xe6, 0x0e, 0x08, 0x2e, 0x77, 0x9b, 0x8a, 0x9d, 0x46, 0x97,
	0x4b, 0x53, 0xf9, 0x6e, 0x20, 0xe6, 0xd6, 0xe1, 0x2b, 0x34, 0xda, 0x43, 0x2c, 0x62, 0xb6, 0xef,
	0x22, 0x66, 0x3d, 0x01, 0x2e, 0xe5, 0xf7, 0x41, 0xf0, 0xed, 0x12, 0xf0, 0x36, 0x11, 0x9a, 0xef,
	0x9b, 0xd0, 0x24, 0xa5, 0x67, 0x3a, 0x7f, 0x0c, 0xc4, 0xcc, 0x3b, 0x6d, 0xa8, 0x45, 0x1f, 0xf9,
	0x1c, 0xef, 0x9b, 0xcf, 0x74, 0x9b, 0x9f, 0x09, 0xdd, 0xf9, 0x15, 0x9a, 0x0a, 0xde, 0xe4, 0x27,
	0xdc, 0xd1, 0xe4, 0x3d, 0x97, 0x01, 0x92, 0x7f, 0x0e, 0x7a, 0xc3, 0x8a, 0xdf, 0xca, 0xff, 0x77,
	0x76, 0xf0, 0x74, 0xb2, 0xb2, 0xf7, 0xce, 0xf9, 0x7d, 0x6f, 0xdd, 0x3a, 0xe4, 0x45, 0x7c, 0xbb,
	0xbc, 0x1d, 0x8a, 0x87, 0x1f, 0x79, 0xbc, 0xc8, 0x2f, 0xc4, 0x1c, 0x89, 0x29, 0x81, 0xd1, 0xd7,
	0x0d, 0x54, 0xf9, 0x6a, 0xfb, 0x60, 0xf8, 0xd1, 0x77, 0x2f, 0x22, 0x16, 0x9e, 0x2a, 0xe1, 0xff,
	0x92, 0x7d, 0xe8, 0x79, 0x7c, 0x31, 0xd0, 0x86, 0xbe, 0x0f, 0x56, 0xd4, 0x9b, 0xe8, 0x8e, 0xb3,
	0xe6, 0xdf, 0x70, 0xe9, 0x7b, 0x36, 0xd1, 0x1d, 0x02, 0xd1, 0x44, 0xe3, 0x11, 0x04, 0x2b, 0xbd,
	0x3e, 0xe2, 0xa7, 0xe2, 0xf3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x61, 0x0e, 0x64, 0x16, 0x4f,
	0x0a, 0x00, 0x00,
}
