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
// source: macsec_secy_stats.proto

package cisco_ios_xr_crypto_macsec_secy_oper_macsec_secy_interfaces_interface_stats

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

type MacsecSecyStats_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MacsecSecyStats_KEYS) Reset()         { *m = MacsecSecyStats_KEYS{} }
func (m *MacsecSecyStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*MacsecSecyStats_KEYS) ProtoMessage()    {}
func (*MacsecSecyStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d619ae602856a09, []int{0}
}

func (m *MacsecSecyStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MacsecSecyStats_KEYS.Unmarshal(m, b)
}
func (m *MacsecSecyStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MacsecSecyStats_KEYS.Marshal(b, m, deterministic)
}
func (m *MacsecSecyStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MacsecSecyStats_KEYS.Merge(m, src)
}
func (m *MacsecSecyStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_MacsecSecyStats_KEYS.Size(m)
}
func (m *MacsecSecyStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MacsecSecyStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MacsecSecyStats_KEYS proto.InternalMessageInfo

func (m *MacsecSecyStats_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MacsecIfStats struct {
	InPktsUntagged       uint64   `protobuf:"varint,1,opt,name=in_pkts_untagged,json=inPktsUntagged,proto3" json:"in_pkts_untagged,omitempty"`
	InPktsNoTag          uint64   `protobuf:"varint,2,opt,name=in_pkts_no_tag,json=inPktsNoTag,proto3" json:"in_pkts_no_tag,omitempty"`
	InPktsBadTag         uint64   `protobuf:"varint,3,opt,name=in_pkts_bad_tag,json=inPktsBadTag,proto3" json:"in_pkts_bad_tag,omitempty"`
	InPktsUnknownSci     uint64   `protobuf:"varint,4,opt,name=in_pkts_unknown_sci,json=inPktsUnknownSci,proto3" json:"in_pkts_unknown_sci,omitempty"`
	InPktsNoSci          uint64   `protobuf:"varint,5,opt,name=in_pkts_no_sci,json=inPktsNoSci,proto3" json:"in_pkts_no_sci,omitempty"`
	InPktsOverrun        uint64   `protobuf:"varint,6,opt,name=in_pkts_overrun,json=inPktsOverrun,proto3" json:"in_pkts_overrun,omitempty"`
	InOctetsValidated    uint64   `protobuf:"varint,7,opt,name=in_octets_validated,json=inOctetsValidated,proto3" json:"in_octets_validated,omitempty"`
	InOctetsDecrypted    uint64   `protobuf:"varint,8,opt,name=in_octets_decrypted,json=inOctetsDecrypted,proto3" json:"in_octets_decrypted,omitempty"`
	OutPktsUntagged      uint64   `protobuf:"varint,9,opt,name=out_pkts_untagged,json=outPktsUntagged,proto3" json:"out_pkts_untagged,omitempty"`
	OutPktsTooLong       uint64   `protobuf:"varint,10,opt,name=out_pkts_too_long,json=outPktsTooLong,proto3" json:"out_pkts_too_long,omitempty"`
	OutOctetsProtected   uint64   `protobuf:"varint,11,opt,name=out_octets_protected,json=outOctetsProtected,proto3" json:"out_octets_protected,omitempty"`
	OutOctetsEncrypted   uint64   `protobuf:"varint,12,opt,name=out_octets_encrypted,json=outOctetsEncrypted,proto3" json:"out_octets_encrypted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MacsecIfStats) Reset()         { *m = MacsecIfStats{} }
func (m *MacsecIfStats) String() string { return proto.CompactTextString(m) }
func (*MacsecIfStats) ProtoMessage()    {}
func (*MacsecIfStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d619ae602856a09, []int{1}
}

func (m *MacsecIfStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MacsecIfStats.Unmarshal(m, b)
}
func (m *MacsecIfStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MacsecIfStats.Marshal(b, m, deterministic)
}
func (m *MacsecIfStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MacsecIfStats.Merge(m, src)
}
func (m *MacsecIfStats) XXX_Size() int {
	return xxx_messageInfo_MacsecIfStats.Size(m)
}
func (m *MacsecIfStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MacsecIfStats.DiscardUnknown(m)
}

var xxx_messageInfo_MacsecIfStats proto.InternalMessageInfo

func (m *MacsecIfStats) GetInPktsUntagged() uint64 {
	if m != nil {
		return m.InPktsUntagged
	}
	return 0
}

func (m *MacsecIfStats) GetInPktsNoTag() uint64 {
	if m != nil {
		return m.InPktsNoTag
	}
	return 0
}

func (m *MacsecIfStats) GetInPktsBadTag() uint64 {
	if m != nil {
		return m.InPktsBadTag
	}
	return 0
}

func (m *MacsecIfStats) GetInPktsUnknownSci() uint64 {
	if m != nil {
		return m.InPktsUnknownSci
	}
	return 0
}

func (m *MacsecIfStats) GetInPktsNoSci() uint64 {
	if m != nil {
		return m.InPktsNoSci
	}
	return 0
}

func (m *MacsecIfStats) GetInPktsOverrun() uint64 {
	if m != nil {
		return m.InPktsOverrun
	}
	return 0
}

func (m *MacsecIfStats) GetInOctetsValidated() uint64 {
	if m != nil {
		return m.InOctetsValidated
	}
	return 0
}

func (m *MacsecIfStats) GetInOctetsDecrypted() uint64 {
	if m != nil {
		return m.InOctetsDecrypted
	}
	return 0
}

func (m *MacsecIfStats) GetOutPktsUntagged() uint64 {
	if m != nil {
		return m.OutPktsUntagged
	}
	return 0
}

func (m *MacsecIfStats) GetOutPktsTooLong() uint64 {
	if m != nil {
		return m.OutPktsTooLong
	}
	return 0
}

func (m *MacsecIfStats) GetOutOctetsProtected() uint64 {
	if m != nil {
		return m.OutOctetsProtected
	}
	return 0
}

func (m *MacsecIfStats) GetOutOctetsEncrypted() uint64 {
	if m != nil {
		return m.OutOctetsEncrypted
	}
	return 0
}

type MacsecTxSaStats struct {
	OutPktsProtected     uint64   `protobuf:"varint,1,opt,name=out_pkts_protected,json=outPktsProtected,proto3" json:"out_pkts_protected,omitempty"`
	OutPktsEncrypted     uint64   `protobuf:"varint,2,opt,name=out_pkts_encrypted,json=outPktsEncrypted,proto3" json:"out_pkts_encrypted,omitempty"`
	NextPn               uint64   `protobuf:"varint,3,opt,name=next_pn,json=nextPn,proto3" json:"next_pn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MacsecTxSaStats) Reset()         { *m = MacsecTxSaStats{} }
func (m *MacsecTxSaStats) String() string { return proto.CompactTextString(m) }
func (*MacsecTxSaStats) ProtoMessage()    {}
func (*MacsecTxSaStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d619ae602856a09, []int{2}
}

func (m *MacsecTxSaStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MacsecTxSaStats.Unmarshal(m, b)
}
func (m *MacsecTxSaStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MacsecTxSaStats.Marshal(b, m, deterministic)
}
func (m *MacsecTxSaStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MacsecTxSaStats.Merge(m, src)
}
func (m *MacsecTxSaStats) XXX_Size() int {
	return xxx_messageInfo_MacsecTxSaStats.Size(m)
}
func (m *MacsecTxSaStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MacsecTxSaStats.DiscardUnknown(m)
}

var xxx_messageInfo_MacsecTxSaStats proto.InternalMessageInfo

func (m *MacsecTxSaStats) GetOutPktsProtected() uint64 {
	if m != nil {
		return m.OutPktsProtected
	}
	return 0
}

func (m *MacsecTxSaStats) GetOutPktsEncrypted() uint64 {
	if m != nil {
		return m.OutPktsEncrypted
	}
	return 0
}

func (m *MacsecTxSaStats) GetNextPn() uint64 {
	if m != nil {
		return m.NextPn
	}
	return 0
}

type MacsecTxScStats struct {
	TxSci                uint64             `protobuf:"varint,1,opt,name=tx_sci,json=txSci,proto3" json:"tx_sci,omitempty"`
	OutPktsProtected     uint64             `protobuf:"varint,2,opt,name=out_pkts_protected,json=outPktsProtected,proto3" json:"out_pkts_protected,omitempty"`
	OutPktsEncrypted     uint64             `protobuf:"varint,3,opt,name=out_pkts_encrypted,json=outPktsEncrypted,proto3" json:"out_pkts_encrypted,omitempty"`
	OutOctetsProtected   uint64             `protobuf:"varint,4,opt,name=out_octets_protected,json=outOctetsProtected,proto3" json:"out_octets_protected,omitempty"`
	OutOctetsEncrypted   uint64             `protobuf:"varint,5,opt,name=out_octets_encrypted,json=outOctetsEncrypted,proto3" json:"out_octets_encrypted,omitempty"`
	OutPktsTooLong       uint64             `protobuf:"varint,6,opt,name=out_pkts_too_long,json=outPktsTooLong,proto3" json:"out_pkts_too_long,omitempty"`
	TxsaStat             []*MacsecTxSaStats `protobuf:"bytes,7,rep,name=txsa_stat,json=txsaStat,proto3" json:"txsa_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MacsecTxScStats) Reset()         { *m = MacsecTxScStats{} }
func (m *MacsecTxScStats) String() string { return proto.CompactTextString(m) }
func (*MacsecTxScStats) ProtoMessage()    {}
func (*MacsecTxScStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d619ae602856a09, []int{3}
}

func (m *MacsecTxScStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MacsecTxScStats.Unmarshal(m, b)
}
func (m *MacsecTxScStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MacsecTxScStats.Marshal(b, m, deterministic)
}
func (m *MacsecTxScStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MacsecTxScStats.Merge(m, src)
}
func (m *MacsecTxScStats) XXX_Size() int {
	return xxx_messageInfo_MacsecTxScStats.Size(m)
}
func (m *MacsecTxScStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MacsecTxScStats.DiscardUnknown(m)
}

var xxx_messageInfo_MacsecTxScStats proto.InternalMessageInfo

func (m *MacsecTxScStats) GetTxSci() uint64 {
	if m != nil {
		return m.TxSci
	}
	return 0
}

func (m *MacsecTxScStats) GetOutPktsProtected() uint64 {
	if m != nil {
		return m.OutPktsProtected
	}
	return 0
}

func (m *MacsecTxScStats) GetOutPktsEncrypted() uint64 {
	if m != nil {
		return m.OutPktsEncrypted
	}
	return 0
}

func (m *MacsecTxScStats) GetOutOctetsProtected() uint64 {
	if m != nil {
		return m.OutOctetsProtected
	}
	return 0
}

func (m *MacsecTxScStats) GetOutOctetsEncrypted() uint64 {
	if m != nil {
		return m.OutOctetsEncrypted
	}
	return 0
}

func (m *MacsecTxScStats) GetOutPktsTooLong() uint64 {
	if m != nil {
		return m.OutPktsTooLong
	}
	return 0
}

func (m *MacsecTxScStats) GetTxsaStat() []*MacsecTxSaStats {
	if m != nil {
		return m.TxsaStat
	}
	return nil
}

type MacsecRxSaStats struct {
	InPktsOk             uint64   `protobuf:"varint,1,opt,name=in_pkts_ok,json=inPktsOk,proto3" json:"in_pkts_ok,omitempty"`
	InPktsInvalid        uint64   `protobuf:"varint,2,opt,name=in_pkts_invalid,json=inPktsInvalid,proto3" json:"in_pkts_invalid,omitempty"`
	InPktsNotValid       uint64   `protobuf:"varint,3,opt,name=in_pkts_not_valid,json=inPktsNotValid,proto3" json:"in_pkts_not_valid,omitempty"`
	InPktsNotUsingSa     uint64   `protobuf:"varint,4,opt,name=in_pkts_not_using_sa,json=inPktsNotUsingSa,proto3" json:"in_pkts_not_using_sa,omitempty"`
	InPktsUnusedSa       uint64   `protobuf:"varint,5,opt,name=in_pkts_unused_sa,json=inPktsUnusedSa,proto3" json:"in_pkts_unused_sa,omitempty"`
	NextPn               uint64   `protobuf:"varint,6,opt,name=next_pn,json=nextPn,proto3" json:"next_pn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MacsecRxSaStats) Reset()         { *m = MacsecRxSaStats{} }
func (m *MacsecRxSaStats) String() string { return proto.CompactTextString(m) }
func (*MacsecRxSaStats) ProtoMessage()    {}
func (*MacsecRxSaStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d619ae602856a09, []int{4}
}

func (m *MacsecRxSaStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MacsecRxSaStats.Unmarshal(m, b)
}
func (m *MacsecRxSaStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MacsecRxSaStats.Marshal(b, m, deterministic)
}
func (m *MacsecRxSaStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MacsecRxSaStats.Merge(m, src)
}
func (m *MacsecRxSaStats) XXX_Size() int {
	return xxx_messageInfo_MacsecRxSaStats.Size(m)
}
func (m *MacsecRxSaStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MacsecRxSaStats.DiscardUnknown(m)
}

var xxx_messageInfo_MacsecRxSaStats proto.InternalMessageInfo

func (m *MacsecRxSaStats) GetInPktsOk() uint64 {
	if m != nil {
		return m.InPktsOk
	}
	return 0
}

func (m *MacsecRxSaStats) GetInPktsInvalid() uint64 {
	if m != nil {
		return m.InPktsInvalid
	}
	return 0
}

func (m *MacsecRxSaStats) GetInPktsNotValid() uint64 {
	if m != nil {
		return m.InPktsNotValid
	}
	return 0
}

func (m *MacsecRxSaStats) GetInPktsNotUsingSa() uint64 {
	if m != nil {
		return m.InPktsNotUsingSa
	}
	return 0
}

func (m *MacsecRxSaStats) GetInPktsUnusedSa() uint64 {
	if m != nil {
		return m.InPktsUnusedSa
	}
	return 0
}

func (m *MacsecRxSaStats) GetNextPn() uint64 {
	if m != nil {
		return m.NextPn
	}
	return 0
}

type MacsecRxScStats struct {
	RxSci                uint64             `protobuf:"varint,1,opt,name=rx_sci,json=rxSci,proto3" json:"rx_sci,omitempty"`
	InPktsUnchecked      uint64             `protobuf:"varint,2,opt,name=in_pkts_unchecked,json=inPktsUnchecked,proto3" json:"in_pkts_unchecked,omitempty"`
	InPktsDelayed        uint64             `protobuf:"varint,3,opt,name=in_pkts_delayed,json=inPktsDelayed,proto3" json:"in_pkts_delayed,omitempty"`
	InPktsLate           uint64             `protobuf:"varint,4,opt,name=in_pkts_late,json=inPktsLate,proto3" json:"in_pkts_late,omitempty"`
	InPktsOk             uint64             `protobuf:"varint,5,opt,name=in_pkts_ok,json=inPktsOk,proto3" json:"in_pkts_ok,omitempty"`
	InPktsInvalid        uint64             `protobuf:"varint,6,opt,name=in_pkts_invalid,json=inPktsInvalid,proto3" json:"in_pkts_invalid,omitempty"`
	InPktsNotValid       uint64             `protobuf:"varint,7,opt,name=in_pkts_not_valid,json=inPktsNotValid,proto3" json:"in_pkts_not_valid,omitempty"`
	InPktsNotUsingSa     uint64             `protobuf:"varint,8,opt,name=in_pkts_not_using_sa,json=inPktsNotUsingSa,proto3" json:"in_pkts_not_using_sa,omitempty"`
	InPktsUnusedSa       uint64             `protobuf:"varint,9,opt,name=in_pkts_unused_sa,json=inPktsUnusedSa,proto3" json:"in_pkts_unused_sa,omitempty"`
	InPktsUntaggedHit    uint64             `protobuf:"varint,10,opt,name=in_pkts_untagged_hit,json=inPktsUntaggedHit,proto3" json:"in_pkts_untagged_hit,omitempty"`
	InOctetsValidated    uint64             `protobuf:"varint,11,opt,name=in_octets_validated,json=inOctetsValidated,proto3" json:"in_octets_validated,omitempty"`
	InOctetsDecrypted    uint64             `protobuf:"varint,12,opt,name=in_octets_decrypted,json=inOctetsDecrypted,proto3" json:"in_octets_decrypted,omitempty"`
	RxsaStat             []*MacsecRxSaStats `protobuf:"bytes,13,rep,name=rxsa_stat,json=rxsaStat,proto3" json:"rxsa_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MacsecRxScStats) Reset()         { *m = MacsecRxScStats{} }
func (m *MacsecRxScStats) String() string { return proto.CompactTextString(m) }
func (*MacsecRxScStats) ProtoMessage()    {}
func (*MacsecRxScStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d619ae602856a09, []int{5}
}

func (m *MacsecRxScStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MacsecRxScStats.Unmarshal(m, b)
}
func (m *MacsecRxScStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MacsecRxScStats.Marshal(b, m, deterministic)
}
func (m *MacsecRxScStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MacsecRxScStats.Merge(m, src)
}
func (m *MacsecRxScStats) XXX_Size() int {
	return xxx_messageInfo_MacsecRxScStats.Size(m)
}
func (m *MacsecRxScStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MacsecRxScStats.DiscardUnknown(m)
}

var xxx_messageInfo_MacsecRxScStats proto.InternalMessageInfo

func (m *MacsecRxScStats) GetRxSci() uint64 {
	if m != nil {
		return m.RxSci
	}
	return 0
}

func (m *MacsecRxScStats) GetInPktsUnchecked() uint64 {
	if m != nil {
		return m.InPktsUnchecked
	}
	return 0
}

func (m *MacsecRxScStats) GetInPktsDelayed() uint64 {
	if m != nil {
		return m.InPktsDelayed
	}
	return 0
}

func (m *MacsecRxScStats) GetInPktsLate() uint64 {
	if m != nil {
		return m.InPktsLate
	}
	return 0
}

func (m *MacsecRxScStats) GetInPktsOk() uint64 {
	if m != nil {
		return m.InPktsOk
	}
	return 0
}

func (m *MacsecRxScStats) GetInPktsInvalid() uint64 {
	if m != nil {
		return m.InPktsInvalid
	}
	return 0
}

func (m *MacsecRxScStats) GetInPktsNotValid() uint64 {
	if m != nil {
		return m.InPktsNotValid
	}
	return 0
}

func (m *MacsecRxScStats) GetInPktsNotUsingSa() uint64 {
	if m != nil {
		return m.InPktsNotUsingSa
	}
	return 0
}

func (m *MacsecRxScStats) GetInPktsUnusedSa() uint64 {
	if m != nil {
		return m.InPktsUnusedSa
	}
	return 0
}

func (m *MacsecRxScStats) GetInPktsUntaggedHit() uint64 {
	if m != nil {
		return m.InPktsUntaggedHit
	}
	return 0
}

func (m *MacsecRxScStats) GetInOctetsValidated() uint64 {
	if m != nil {
		return m.InOctetsValidated
	}
	return 0
}

func (m *MacsecRxScStats) GetInOctetsDecrypted() uint64 {
	if m != nil {
		return m.InOctetsDecrypted
	}
	return 0
}

func (m *MacsecRxScStats) GetRxsaStat() []*MacsecRxSaStats {
	if m != nil {
		return m.RxsaStat
	}
	return nil
}

type MacsecSecyStats struct {
	IntfStats            *MacsecIfStats     `protobuf:"bytes,50,opt,name=intf_stats,json=intfStats,proto3" json:"intf_stats,omitempty"`
	TxScStats            *MacsecTxScStats   `protobuf:"bytes,51,opt,name=tx_sc_stats,json=txScStats,proto3" json:"tx_sc_stats,omitempty"`
	RxScStats            []*MacsecRxScStats `protobuf:"bytes,52,rep,name=rx_sc_stats,json=rxScStats,proto3" json:"rx_sc_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MacsecSecyStats) Reset()         { *m = MacsecSecyStats{} }
func (m *MacsecSecyStats) String() string { return proto.CompactTextString(m) }
func (*MacsecSecyStats) ProtoMessage()    {}
func (*MacsecSecyStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d619ae602856a09, []int{6}
}

func (m *MacsecSecyStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MacsecSecyStats.Unmarshal(m, b)
}
func (m *MacsecSecyStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MacsecSecyStats.Marshal(b, m, deterministic)
}
func (m *MacsecSecyStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MacsecSecyStats.Merge(m, src)
}
func (m *MacsecSecyStats) XXX_Size() int {
	return xxx_messageInfo_MacsecSecyStats.Size(m)
}
func (m *MacsecSecyStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MacsecSecyStats.DiscardUnknown(m)
}

var xxx_messageInfo_MacsecSecyStats proto.InternalMessageInfo

func (m *MacsecSecyStats) GetIntfStats() *MacsecIfStats {
	if m != nil {
		return m.IntfStats
	}
	return nil
}

func (m *MacsecSecyStats) GetTxScStats() *MacsecTxScStats {
	if m != nil {
		return m.TxScStats
	}
	return nil
}

func (m *MacsecSecyStats) GetRxScStats() []*MacsecRxScStats {
	if m != nil {
		return m.RxScStats
	}
	return nil
}

func init() {
	proto.RegisterType((*MacsecSecyStats_KEYS)(nil), "cisco_ios_xr_crypto_macsec_secy_oper.macsec.secy.interfaces.interface.stats.macsec_secy_stats_KEYS")
	proto.RegisterType((*MacsecIfStats)(nil), "cisco_ios_xr_crypto_macsec_secy_oper.macsec.secy.interfaces.interface.stats.macsec_if_stats")
	proto.RegisterType((*MacsecTxSaStats)(nil), "cisco_ios_xr_crypto_macsec_secy_oper.macsec.secy.interfaces.interface.stats.macsec_tx_sa_stats")
	proto.RegisterType((*MacsecTxScStats)(nil), "cisco_ios_xr_crypto_macsec_secy_oper.macsec.secy.interfaces.interface.stats.macsec_tx_sc_stats")
	proto.RegisterType((*MacsecRxSaStats)(nil), "cisco_ios_xr_crypto_macsec_secy_oper.macsec.secy.interfaces.interface.stats.macsec_rx_sa_stats")
	proto.RegisterType((*MacsecRxScStats)(nil), "cisco_ios_xr_crypto_macsec_secy_oper.macsec.secy.interfaces.interface.stats.macsec_rx_sc_stats")
	proto.RegisterType((*MacsecSecyStats)(nil), "cisco_ios_xr_crypto_macsec_secy_oper.macsec.secy.interfaces.interface.stats.macsec_secy_stats")
}

func init() { proto.RegisterFile("macsec_secy_stats.proto", fileDescriptor_3d619ae602856a09) }

var fileDescriptor_3d619ae602856a09 = []byte{
	// 776 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x96, 0x5d, 0x6b, 0x13, 0x4d,
	0x14, 0xc7, 0x49, 0xf3, 0x3e, 0x49, 0x9b, 0x66, 0x9e, 0x3e, 0x76, 0x2f, 0xbc, 0x08, 0x11, 0x25,
	0x96, 0xba, 0x4a, 0xeb, 0x27, 0x90, 0x16, 0x94, 0x96, 0xb6, 0x24, 0xad, 0x20, 0x08, 0xc3, 0x74,
	0x77, 0x9a, 0x2e, 0x49, 0x67, 0xc2, 0xce, 0xd9, 0x9a, 0x62, 0xc1, 0x6b, 0x6f, 0xc5, 0x6b, 0xbf,
	0xa7, 0xe0, 0x85, 0xcc, 0xec, 0xcc, 0xbe, 0xa4, 0x09, 0x1a, 0x30, 0x77, 0xc9, 0x99, 0xdf, 0xee,
	0xff, 0xec, 0x39, 0xff, 0x73, 0x76, 0xd1, 0xf6, 0x0d, 0xf5, 0x24, 0xf3, 0x88, 0x64, 0xde, 0x1d,
	0x91, 0x40, 0x41, 0xba, 0x93, 0x50, 0x80, 0xc0, 0x47, 0x5e, 0x20, 0x3d, 0x41, 0x02, 0x21, 0xc9,
	0x34, 0x24, 0x5e, 0x78, 0x37, 0x01, 0x41, 0xb2, 0xb0, 0x98, 0xb0, 0xd0, 0x8d, 0x03, 0xae, 0x0a,
	0xb8, 0x01, 0x07, 0x16, 0x5e, 0x51, 0x8f, 0xc9, 0xf4, 0xa7, 0xab, 0x6f, 0xd9, 0xdd, 0x45, 0x8f,
	0x1e, 0xe8, 0x90, 0xa3, 0xc3, 0x0f, 0x03, 0x8c, 0x51, 0x89, 0xd3, 0x1b, 0xe6, 0x14, 0x3a, 0x85,
	0x5e, 0xbd, 0xaf, 0x7f, 0x77, 0x7f, 0x94, 0x50, 0xcb, 0xe0, 0xc1, 0x55, 0x0c, 0xe3, 0x1e, 0xda,
	0x0c, 0x38, 0x99, 0x8c, 0x40, 0x92, 0x88, 0x03, 0x1d, 0x0e, 0x99, 0xaf, 0xaf, 0x29, 0xf5, 0x37,
	0x02, 0x7e, 0x36, 0x02, 0x79, 0x61, 0xa2, 0xf8, 0x09, 0xda, 0xb0, 0x24, 0x17, 0x04, 0xe8, 0xd0,
	0x59, 0xd3, 0x5c, 0x23, 0xe6, 0x4e, 0xc4, 0x39, 0x1d, 0xe2, 0xa7, 0xa8, 0x65, 0xa1, 0x4b, 0xea,
	0x6b, 0xaa, 0xa8, 0xa9, 0x66, 0x4c, 0xbd, 0xa1, 0xbe, 0xc2, 0x5e, 0xa0, 0xff, 0x52, 0xd5, 0x11,
	0x17, 0x9f, 0x38, 0x91, 0x5e, 0xe0, 0x94, 0x34, 0xba, 0x69, 0x85, 0xf5, 0xc1, 0xc0, 0x0b, 0x66,
	0xa4, 0x15, 0x59, 0xce, 0x4b, 0x2b, 0xe8, 0x59, 0x2a, 0x2d, 0x6e, 0x59, 0x18, 0x46, 0xdc, 0xa9,
	0x68, 0x6a, 0x3d, 0xa6, 0x4e, 0xe3, 0x20, 0x76, 0xb5, 0xb6, 0xf0, 0x80, 0x81, 0x24, 0xb7, 0x74,
	0x1c, 0xf8, 0x14, 0x98, 0xef, 0x54, 0x35, 0xdb, 0x0e, 0xf8, 0xa9, 0x3e, 0x79, 0x6f, 0x0f, 0xf2,
	0xbc, 0xcf, 0x74, 0xc7, 0x98, 0xef, 0xd4, 0xf2, 0xfc, 0x81, 0x3d, 0xc0, 0x3b, 0xa8, 0x2d, 0x22,
	0x98, 0x29, 0x69, 0x5d, 0xd3, 0x2d, 0x11, 0x41, 0xae, 0xa6, 0xcf, 0x33, 0x2c, 0x08, 0x41, 0xc6,
	0x82, 0x0f, 0x1d, 0x14, 0x97, 0xdf, 0xb0, 0xe7, 0x42, 0x1c, 0x0b, 0x3e, 0xc4, 0xaf, 0xd0, 0x96,
	0x42, 0x4d, 0x1e, 0xca, 0x4b, 0xcc, 0x53, 0x79, 0x34, 0x34, 0x8d, 0x45, 0x04, 0x71, 0x22, 0x67,
	0xf6, 0x64, 0xe6, 0x0a, 0xc6, 0x6d, 0xe6, 0xcd, 0x99, 0x2b, 0x0e, 0xed, 0x49, 0xf7, 0x6b, 0x01,
	0x61, 0x63, 0x10, 0x98, 0x12, 0x49, 0x8d, 0x47, 0x76, 0x11, 0x4e, 0xb2, 0x4c, 0x85, 0x63, 0x97,
	0x6c, 0x9a, 0x34, 0x53, 0xd9, 0x2c, 0x9d, 0x8a, 0xae, 0xe5, 0xe8, 0x44, 0x12, 0x6f, 0xa3, 0x2a,
	0x67, 0x53, 0x20, 0x13, 0x6e, 0x8c, 0x52, 0x51, 0x7f, 0xcf, 0x78, 0xf7, 0x7b, 0x31, 0x97, 0x8b,
	0x67, 0x72, 0xf9, 0x1f, 0x55, 0xf4, 0xdf, 0xc0, 0xe8, 0x97, 0x61, 0xaa, 0x9a, 0x3f, 0x3f, 0xc5,
	0xb5, 0xa5, 0x52, 0x2c, 0x2e, 0x48, 0x71, 0x51, 0xe5, 0x4b, 0x4b, 0x57, 0xbe, 0xbc, 0xa8, 0xf2,
	0xf3, 0x8d, 0x50, 0x99, 0x6b, 0x84, 0x7b, 0x54, 0x87, 0xa9, 0xe9, 0x8d, 0x53, 0xed, 0x14, 0x7b,
	0x8d, 0x3d, 0xe2, 0xfe, 0xc3, 0xa5, 0xe2, 0x3e, 0x74, 0x40, 0xbf, 0xa6, 0x14, 0x07, 0x40, 0xa1,
	0xfb, 0x33, 0xb5, 0x48, 0x98, 0xb1, 0xc8, 0x63, 0x84, 0x92, 0xe1, 0x1b, 0x99, 0xd6, 0xd4, 0xcc,
	0xdc, 0x8d, 0xb2, 0xa3, 0x19, 0x70, 0x3d, 0x72, 0xa6, 0x35, 0x66, 0x34, 0xdf, 0xc5, 0x41, 0x55,
	0x85, 0x74, 0xce, 0x21, 0x1e, 0x4e, 0xd3, 0x96, 0x0d, 0x3b, 0xea, 0xa0, 0x27, 0x13, 0xbb, 0x68,
	0x2b, 0x8b, 0x46, 0x32, 0xe0, 0x43, 0x22, 0x69, 0x7e, 0x85, 0x9c, 0x08, 0xb8, 0x50, 0x07, 0x03,
	0x9a, 0xbd, 0x75, 0xc4, 0x23, 0xc9, 0x7c, 0x05, 0x97, 0xf3, 0x8b, 0x4e, 0x85, 0x07, 0x34, 0x6b,
	0xc9, 0x4a, 0xce, 0x92, 0xbf, 0x4a, 0xb9, 0x67, 0xcf, 0x58, 0x32, 0xcc, 0x59, 0x32, 0xd4, 0x96,
	0xdc, 0xc9, 0x2a, 0x7a, 0xd7, 0xcc, 0x1b, 0x25, 0x8e, 0x6c, 0x59, 0x45, 0x13, 0xce, 0x16, 0xc8,
	0x67, 0x63, 0x7a, 0x97, 0xb8, 0xd1, 0x14, 0xe8, 0x20, 0x0e, 0xe2, 0x0e, 0x6a, 0x5a, 0x6e, 0x4c,
	0x81, 0x99, 0xa7, 0x45, 0x31, 0x74, 0x4c, 0x81, 0xcd, 0x34, 0xa2, 0xfc, 0xe7, 0x46, 0x54, 0xfe,
	0xba, 0x11, 0xd5, 0xa5, 0x1a, 0x51, 0x5b, 0xa6, 0x11, 0xf5, 0xb9, 0x8d, 0x78, 0x99, 0xde, 0xda,
	0x2e, 0x52, 0x72, 0x1d, 0x80, 0x59, 0x90, 0xed, 0xfc, 0xfb, 0xe9, 0x6d, 0x00, 0x8b, 0x56, 0x7b,
	0x63, 0xc9, 0xd5, 0xde, 0x5c, 0xb4, 0xda, 0xef, 0x51, 0x3d, 0x4c, 0x46, 0x6f, 0x7d, 0x75, 0xa3,
	0x17, 0x66, 0x47, 0x2f, 0xb4, 0xa3, 0xf7, 0xad, 0x88, 0xda, 0x0f, 0xde, 0xf6, 0xf8, 0xb3, 0x6a,
	0x38, 0x98, 0xd7, 0xb9, 0xb3, 0xd7, 0x29, 0xf4, 0x1a, 0x7b, 0x1f, 0x57, 0x91, 0x94, 0xfd, 0x64,
	0xe8, 0xd7, 0x95, 0xde, 0x40, 0x8b, 0x7f, 0x41, 0x8d, 0xcc, 0x72, 0x76, 0xf6, 0xb5, 0xfa, 0xca,
	0xb6, 0x91, 0x67, 0x13, 0x50, 0x3b, 0x3f, 0x49, 0x20, 0x33, 0x8a, 0xce, 0xeb, 0xd5, 0xf6, 0x24,
	0x49, 0x20, 0xb4, 0x09, 0x5c, 0x56, 0xf4, 0x57, 0xdd, 0xfe, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0x73, 0x94, 0x4d, 0xf0, 0x09, 0x00, 0x00,
}