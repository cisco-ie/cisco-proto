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
// source: l2vpn_ldp_atom.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_ato_mdb_table_entries_ato_mdb_table_entry

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

type L2VpnLdpAtom_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	PeerId               string   `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	FeCtype              uint32   `protobuf:"varint,3,opt,name=fe_ctype,json=feCtype,proto3" json:"fe_ctype,omitempty"`
	PwId                 uint32   `protobuf:"varint,4,opt,name=pw_id,json=pwId,proto3" json:"pw_id,omitempty"`
	AgiType              uint32   `protobuf:"varint,5,opt,name=agi_type,json=agiType,proto3" json:"agi_type,omitempty"`
	Agi                  uint32   `protobuf:"varint,6,opt,name=agi,proto3" json:"agi,omitempty"`
	SaiiType             uint32   `protobuf:"varint,7,opt,name=saii_type,json=saiiType,proto3" json:"saii_type,omitempty"`
	SaiiLocalId          string   `protobuf:"bytes,8,opt,name=saii_local_id,json=saiiLocalId,proto3" json:"saii_local_id,omitempty"`
	SaiiGlobalId         uint32   `protobuf:"varint,9,opt,name=saii_global_id,json=saiiGlobalId,proto3" json:"saii_global_id,omitempty"`
	SaiiPrefix           string   `protobuf:"bytes,10,opt,name=saii_prefix,json=saiiPrefix,proto3" json:"saii_prefix,omitempty"`
	SaiiAcId             uint32   `protobuf:"varint,11,opt,name=saii_ac_id,json=saiiAcId,proto3" json:"saii_ac_id,omitempty"`
	TaiiType             uint32   `protobuf:"varint,12,opt,name=taii_type,json=taiiType,proto3" json:"taii_type,omitempty"`
	TaiiLocalId          string   `protobuf:"bytes,13,opt,name=taii_local_id,json=taiiLocalId,proto3" json:"taii_local_id,omitempty"`
	TaiiGlobalId         uint32   `protobuf:"varint,14,opt,name=taii_global_id,json=taiiGlobalId,proto3" json:"taii_global_id,omitempty"`
	TaiiPrefix           string   `protobuf:"bytes,15,opt,name=taii_prefix,json=taiiPrefix,proto3" json:"taii_prefix,omitempty"`
	TaiiAcId             uint32   `protobuf:"varint,16,opt,name=taii_ac_id,json=taiiAcId,proto3" json:"taii_ac_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnLdpAtom_KEYS) Reset()         { *m = L2VpnLdpAtom_KEYS{} }
func (m *L2VpnLdpAtom_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpAtom_KEYS) ProtoMessage()    {}
func (*L2VpnLdpAtom_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{0}
}

func (m *L2VpnLdpAtom_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpAtom_KEYS.Unmarshal(m, b)
}
func (m *L2VpnLdpAtom_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpAtom_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpAtom_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpAtom_KEYS.Merge(m, src)
}
func (m *L2VpnLdpAtom_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpAtom_KEYS.Size(m)
}
func (m *L2VpnLdpAtom_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpAtom_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpAtom_KEYS proto.InternalMessageInfo

func (m *L2VpnLdpAtom_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *L2VpnLdpAtom_KEYS) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *L2VpnLdpAtom_KEYS) GetFeCtype() uint32 {
	if m != nil {
		return m.FeCtype
	}
	return 0
}

func (m *L2VpnLdpAtom_KEYS) GetPwId() uint32 {
	if m != nil {
		return m.PwId
	}
	return 0
}

func (m *L2VpnLdpAtom_KEYS) GetAgiType() uint32 {
	if m != nil {
		return m.AgiType
	}
	return 0
}

func (m *L2VpnLdpAtom_KEYS) GetAgi() uint32 {
	if m != nil {
		return m.Agi
	}
	return 0
}

func (m *L2VpnLdpAtom_KEYS) GetSaiiType() uint32 {
	if m != nil {
		return m.SaiiType
	}
	return 0
}

func (m *L2VpnLdpAtom_KEYS) GetSaiiLocalId() string {
	if m != nil {
		return m.SaiiLocalId
	}
	return ""
}

func (m *L2VpnLdpAtom_KEYS) GetSaiiGlobalId() uint32 {
	if m != nil {
		return m.SaiiGlobalId
	}
	return 0
}

func (m *L2VpnLdpAtom_KEYS) GetSaiiPrefix() string {
	if m != nil {
		return m.SaiiPrefix
	}
	return ""
}

func (m *L2VpnLdpAtom_KEYS) GetSaiiAcId() uint32 {
	if m != nil {
		return m.SaiiAcId
	}
	return 0
}

func (m *L2VpnLdpAtom_KEYS) GetTaiiType() uint32 {
	if m != nil {
		return m.TaiiType
	}
	return 0
}

func (m *L2VpnLdpAtom_KEYS) GetTaiiLocalId() string {
	if m != nil {
		return m.TaiiLocalId
	}
	return ""
}

func (m *L2VpnLdpAtom_KEYS) GetTaiiGlobalId() uint32 {
	if m != nil {
		return m.TaiiGlobalId
	}
	return 0
}

func (m *L2VpnLdpAtom_KEYS) GetTaiiPrefix() string {
	if m != nil {
		return m.TaiiPrefix
	}
	return ""
}

func (m *L2VpnLdpAtom_KEYS) GetTaiiAcId() uint32 {
	if m != nil {
		return m.TaiiAcId
	}
	return 0
}

type L2VpnLdpPwFec_128T struct {
	PseudowireId         uint32   `protobuf:"varint,1,opt,name=pseudowire_id,json=pseudowireId,proto3" json:"pseudowire_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnLdpPwFec_128T) Reset()         { *m = L2VpnLdpPwFec_128T{} }
func (m *L2VpnLdpPwFec_128T) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpPwFec_128T) ProtoMessage()    {}
func (*L2VpnLdpPwFec_128T) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{1}
}

func (m *L2VpnLdpPwFec_128T) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpPwFec_128T.Unmarshal(m, b)
}
func (m *L2VpnLdpPwFec_128T) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpPwFec_128T.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpPwFec_128T) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpPwFec_128T.Merge(m, src)
}
func (m *L2VpnLdpPwFec_128T) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpPwFec_128T.Size(m)
}
func (m *L2VpnLdpPwFec_128T) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpPwFec_128T.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpPwFec_128T proto.InternalMessageInfo

func (m *L2VpnLdpPwFec_128T) GetPseudowireId() uint32 {
	if m != nil {
		return m.PseudowireId
	}
	return 0
}

type L2VpnLdpPwAgi1T struct {
	Rd                   uint64   `protobuf:"varint,1,opt,name=rd,proto3" json:"rd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnLdpPwAgi1T) Reset()         { *m = L2VpnLdpPwAgi1T{} }
func (m *L2VpnLdpPwAgi1T) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpPwAgi1T) ProtoMessage()    {}
func (*L2VpnLdpPwAgi1T) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{2}
}

func (m *L2VpnLdpPwAgi1T) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpPwAgi1T.Unmarshal(m, b)
}
func (m *L2VpnLdpPwAgi1T) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpPwAgi1T.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpPwAgi1T) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpPwAgi1T.Merge(m, src)
}
func (m *L2VpnLdpPwAgi1T) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpPwAgi1T.Size(m)
}
func (m *L2VpnLdpPwAgi1T) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpPwAgi1T.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpPwAgi1T proto.InternalMessageInfo

func (m *L2VpnLdpPwAgi1T) GetRd() uint64 {
	if m != nil {
		return m.Rd
	}
	return 0
}

type L2VpnLdpPwAgi struct {
	AgiType              string           `protobuf:"bytes,1,opt,name=agi_type,json=agiType,proto3" json:"agi_type,omitempty"`
	Agi1                 *L2VpnLdpPwAgi1T `protobuf:"bytes,2,opt,name=agi1,proto3" json:"agi1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *L2VpnLdpPwAgi) Reset()         { *m = L2VpnLdpPwAgi{} }
func (m *L2VpnLdpPwAgi) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpPwAgi) ProtoMessage()    {}
func (*L2VpnLdpPwAgi) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{3}
}

func (m *L2VpnLdpPwAgi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpPwAgi.Unmarshal(m, b)
}
func (m *L2VpnLdpPwAgi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpPwAgi.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpPwAgi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpPwAgi.Merge(m, src)
}
func (m *L2VpnLdpPwAgi) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpPwAgi.Size(m)
}
func (m *L2VpnLdpPwAgi) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpPwAgi.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpPwAgi proto.InternalMessageInfo

func (m *L2VpnLdpPwAgi) GetAgiType() string {
	if m != nil {
		return m.AgiType
	}
	return ""
}

func (m *L2VpnLdpPwAgi) GetAgi1() *L2VpnLdpPwAgi1T {
	if m != nil {
		return m.Agi1
	}
	return nil
}

type L2VpnLdpPwAii1T struct {
	LocalId              string   `protobuf:"bytes,1,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnLdpPwAii1T) Reset()         { *m = L2VpnLdpPwAii1T{} }
func (m *L2VpnLdpPwAii1T) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpPwAii1T) ProtoMessage()    {}
func (*L2VpnLdpPwAii1T) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{4}
}

func (m *L2VpnLdpPwAii1T) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpPwAii1T.Unmarshal(m, b)
}
func (m *L2VpnLdpPwAii1T) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpPwAii1T.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpPwAii1T) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpPwAii1T.Merge(m, src)
}
func (m *L2VpnLdpPwAii1T) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpPwAii1T.Size(m)
}
func (m *L2VpnLdpPwAii1T) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpPwAii1T.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpPwAii1T proto.InternalMessageInfo

func (m *L2VpnLdpPwAii1T) GetLocalId() string {
	if m != nil {
		return m.LocalId
	}
	return ""
}

type L2VpnLdpPwAii2T struct {
	GobalId              uint32   `protobuf:"varint,1,opt,name=gobal_id,json=gobalId,proto3" json:"gobal_id,omitempty"`
	Prefix               uint32   `protobuf:"varint,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	AcId                 uint32   `protobuf:"varint,3,opt,name=ac_id,json=acId,proto3" json:"ac_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnLdpPwAii2T) Reset()         { *m = L2VpnLdpPwAii2T{} }
func (m *L2VpnLdpPwAii2T) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpPwAii2T) ProtoMessage()    {}
func (*L2VpnLdpPwAii2T) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{5}
}

func (m *L2VpnLdpPwAii2T) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpPwAii2T.Unmarshal(m, b)
}
func (m *L2VpnLdpPwAii2T) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpPwAii2T.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpPwAii2T) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpPwAii2T.Merge(m, src)
}
func (m *L2VpnLdpPwAii2T) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpPwAii2T.Size(m)
}
func (m *L2VpnLdpPwAii2T) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpPwAii2T.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpPwAii2T proto.InternalMessageInfo

func (m *L2VpnLdpPwAii2T) GetGobalId() uint32 {
	if m != nil {
		return m.GobalId
	}
	return 0
}

func (m *L2VpnLdpPwAii2T) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

func (m *L2VpnLdpPwAii2T) GetAcId() uint32 {
	if m != nil {
		return m.AcId
	}
	return 0
}

type L2VpnLdpPwAii struct {
	AiiType              string           `protobuf:"bytes,1,opt,name=aii_type,json=aiiType,proto3" json:"aii_type,omitempty"`
	Aii1                 *L2VpnLdpPwAii1T `protobuf:"bytes,2,opt,name=aii1,proto3" json:"aii1,omitempty"`
	Aii2                 *L2VpnLdpPwAii2T `protobuf:"bytes,3,opt,name=aii2,proto3" json:"aii2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *L2VpnLdpPwAii) Reset()         { *m = L2VpnLdpPwAii{} }
func (m *L2VpnLdpPwAii) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpPwAii) ProtoMessage()    {}
func (*L2VpnLdpPwAii) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{6}
}

func (m *L2VpnLdpPwAii) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpPwAii.Unmarshal(m, b)
}
func (m *L2VpnLdpPwAii) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpPwAii.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpPwAii) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpPwAii.Merge(m, src)
}
func (m *L2VpnLdpPwAii) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpPwAii.Size(m)
}
func (m *L2VpnLdpPwAii) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpPwAii.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpPwAii proto.InternalMessageInfo

func (m *L2VpnLdpPwAii) GetAiiType() string {
	if m != nil {
		return m.AiiType
	}
	return ""
}

func (m *L2VpnLdpPwAii) GetAii1() *L2VpnLdpPwAii1T {
	if m != nil {
		return m.Aii1
	}
	return nil
}

func (m *L2VpnLdpPwAii) GetAii2() *L2VpnLdpPwAii2T {
	if m != nil {
		return m.Aii2
	}
	return nil
}

type L2VpnLdpPwFec_129T struct {
	Agi                  *L2VpnLdpPwAgi `protobuf:"bytes,1,opt,name=agi,proto3" json:"agi,omitempty"`
	Saii                 *L2VpnLdpPwAii `protobuf:"bytes,2,opt,name=saii,proto3" json:"saii,omitempty"`
	Taii                 *L2VpnLdpPwAii `protobuf:"bytes,3,opt,name=taii,proto3" json:"taii,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *L2VpnLdpPwFec_129T) Reset()         { *m = L2VpnLdpPwFec_129T{} }
func (m *L2VpnLdpPwFec_129T) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpPwFec_129T) ProtoMessage()    {}
func (*L2VpnLdpPwFec_129T) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{7}
}

func (m *L2VpnLdpPwFec_129T) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpPwFec_129T.Unmarshal(m, b)
}
func (m *L2VpnLdpPwFec_129T) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpPwFec_129T.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpPwFec_129T) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpPwFec_129T.Merge(m, src)
}
func (m *L2VpnLdpPwFec_129T) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpPwFec_129T.Size(m)
}
func (m *L2VpnLdpPwFec_129T) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpPwFec_129T.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpPwFec_129T proto.InternalMessageInfo

func (m *L2VpnLdpPwFec_129T) GetAgi() *L2VpnLdpPwAgi {
	if m != nil {
		return m.Agi
	}
	return nil
}

func (m *L2VpnLdpPwFec_129T) GetSaii() *L2VpnLdpPwAii {
	if m != nil {
		return m.Saii
	}
	return nil
}

func (m *L2VpnLdpPwFec_129T) GetTaii() *L2VpnLdpPwAii {
	if m != nil {
		return m.Taii
	}
	return nil
}

type L2VpnLdpPwFec_130T struct {
	Agi                  *L2VpnLdpPwAgi `protobuf:"bytes,1,opt,name=agi,proto3" json:"agi,omitempty"`
	Saii                 *L2VpnLdpPwAii `protobuf:"bytes,2,opt,name=saii,proto3" json:"saii,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *L2VpnLdpPwFec_130T) Reset()         { *m = L2VpnLdpPwFec_130T{} }
func (m *L2VpnLdpPwFec_130T) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpPwFec_130T) ProtoMessage()    {}
func (*L2VpnLdpPwFec_130T) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{8}
}

func (m *L2VpnLdpPwFec_130T) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpPwFec_130T.Unmarshal(m, b)
}
func (m *L2VpnLdpPwFec_130T) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpPwFec_130T.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpPwFec_130T) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpPwFec_130T.Merge(m, src)
}
func (m *L2VpnLdpPwFec_130T) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpPwFec_130T.Size(m)
}
func (m *L2VpnLdpPwFec_130T) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpPwFec_130T.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpPwFec_130T proto.InternalMessageInfo

func (m *L2VpnLdpPwFec_130T) GetAgi() *L2VpnLdpPwAgi {
	if m != nil {
		return m.Agi
	}
	return nil
}

func (m *L2VpnLdpPwFec_130T) GetSaii() *L2VpnLdpPwAii {
	if m != nil {
		return m.Saii
	}
	return nil
}

type L2VpnLdpPwFecInfo struct {
	FeCtype              string              `protobuf:"bytes,1,opt,name=fe_ctype,json=feCtype,proto3" json:"fe_ctype,omitempty"`
	Fec128               *L2VpnLdpPwFec_128T `protobuf:"bytes,2,opt,name=fec128,proto3" json:"fec128,omitempty"`
	Fec129               *L2VpnLdpPwFec_129T `protobuf:"bytes,3,opt,name=fec129,proto3" json:"fec129,omitempty"`
	Fec130               *L2VpnLdpPwFec_130T `protobuf:"bytes,4,opt,name=fec130,proto3" json:"fec130,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *L2VpnLdpPwFecInfo) Reset()         { *m = L2VpnLdpPwFecInfo{} }
func (m *L2VpnLdpPwFecInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpPwFecInfo) ProtoMessage()    {}
func (*L2VpnLdpPwFecInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{9}
}

func (m *L2VpnLdpPwFecInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpPwFecInfo.Unmarshal(m, b)
}
func (m *L2VpnLdpPwFecInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpPwFecInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpPwFecInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpPwFecInfo.Merge(m, src)
}
func (m *L2VpnLdpPwFecInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpPwFecInfo.Size(m)
}
func (m *L2VpnLdpPwFecInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpPwFecInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpPwFecInfo proto.InternalMessageInfo

func (m *L2VpnLdpPwFecInfo) GetFeCtype() string {
	if m != nil {
		return m.FeCtype
	}
	return ""
}

func (m *L2VpnLdpPwFecInfo) GetFec128() *L2VpnLdpPwFec_128T {
	if m != nil {
		return m.Fec128
	}
	return nil
}

func (m *L2VpnLdpPwFecInfo) GetFec129() *L2VpnLdpPwFec_129T {
	if m != nil {
		return m.Fec129
	}
	return nil
}

func (m *L2VpnLdpPwFecInfo) GetFec130() *L2VpnLdpPwFec_130T {
	if m != nil {
		return m.Fec130
	}
	return nil
}

type L2VpnLdpAtom struct {
	PeerIdXr             string             `protobuf:"bytes,50,opt,name=peer_id_xr,json=peerIdXr,proto3" json:"peer_id_xr,omitempty"`
	FecInfo              *L2VpnLdpPwFecInfo `protobuf:"bytes,51,opt,name=fec_info,json=fecInfo,proto3" json:"fec_info,omitempty"`
	MappingTlvCount      uint32             `protobuf:"varint,52,opt,name=mapping_tlv_count,json=mappingTlvCount,proto3" json:"mapping_tlv_count,omitempty"`
	NotificationTlvCount uint32             `protobuf:"varint,53,opt,name=notification_tlv_count,json=notificationTlvCount,proto3" json:"notification_tlv_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L2VpnLdpAtom) Reset()         { *m = L2VpnLdpAtom{} }
func (m *L2VpnLdpAtom) String() string { return proto.CompactTextString(m) }
func (*L2VpnLdpAtom) ProtoMessage()    {}
func (*L2VpnLdpAtom) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db8ca76eb701464, []int{10}
}

func (m *L2VpnLdpAtom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLdpAtom.Unmarshal(m, b)
}
func (m *L2VpnLdpAtom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLdpAtom.Marshal(b, m, deterministic)
}
func (m *L2VpnLdpAtom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLdpAtom.Merge(m, src)
}
func (m *L2VpnLdpAtom) XXX_Size() int {
	return xxx_messageInfo_L2VpnLdpAtom.Size(m)
}
func (m *L2VpnLdpAtom) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLdpAtom.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLdpAtom proto.InternalMessageInfo

func (m *L2VpnLdpAtom) GetPeerIdXr() string {
	if m != nil {
		return m.PeerIdXr
	}
	return ""
}

func (m *L2VpnLdpAtom) GetFecInfo() *L2VpnLdpPwFecInfo {
	if m != nil {
		return m.FecInfo
	}
	return nil
}

func (m *L2VpnLdpAtom) GetMappingTlvCount() uint32 {
	if m != nil {
		return m.MappingTlvCount
	}
	return 0
}

func (m *L2VpnLdpAtom) GetNotificationTlvCount() uint32 {
	if m != nil {
		return m.NotificationTlvCount
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnLdpAtom_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_atom_KEYS")
	proto.RegisterType((*L2VpnLdpPwFec_128T)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_pw_fec_128_t")
	proto.RegisterType((*L2VpnLdpPwAgi1T)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_pw_agi1_t")
	proto.RegisterType((*L2VpnLdpPwAgi)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_pw_agi")
	proto.RegisterType((*L2VpnLdpPwAii1T)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_pw_aii1_t")
	proto.RegisterType((*L2VpnLdpPwAii2T)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_pw_aii2_t")
	proto.RegisterType((*L2VpnLdpPwAii)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_pw_aii")
	proto.RegisterType((*L2VpnLdpPwFec_129T)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_pw_fec_129_t")
	proto.RegisterType((*L2VpnLdpPwFec_130T)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_pw_fec_130_t")
	proto.RegisterType((*L2VpnLdpPwFecInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_pw_fec_info")
	proto.RegisterType((*L2VpnLdpAtom)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.ato_mdb_table_entries.ato_mdb_table_entry.l2vpn_ldp_atom")
}

func init() { proto.RegisterFile("l2vpn_ldp_atom.proto", fileDescriptor_0db8ca76eb701464) }

var fileDescriptor_0db8ca76eb701464 = []byte{
	// 759 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x96, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe5, 0x24, 0x4d, 0xd2, 0x6d, 0x93, 0xf6, 0xe7, 0xf6, 0x57, 0x5c, 0x81, 0x44, 0x65,
	0x40, 0xaa, 0x38, 0x44, 0x8d, 0x53, 0xa4, 0xe6, 0xc0, 0x01, 0x55, 0x08, 0x45, 0x20, 0x84, 0x42,
	0x0f, 0xe5, 0x80, 0x56, 0x1b, 0x7b, 0x63, 0x6d, 0xe5, 0x78, 0x57, 0xce, 0xb6, 0x69, 0x25, 0x84,
	0xc4, 0x0d, 0xf1, 0x00, 0xbd, 0xf1, 0x10, 0x1c, 0x78, 0x12, 0x5e, 0x85, 0x07, 0x40, 0x33, 0xbb,
	0xf9, 0xe3, 0x34, 0x57, 0x7c, 0xe1, 0x12, 0x65, 0x67, 0x67, 0x76, 0x3e, 0x9e, 0xef, 0xec, 0x1f,
	0xb2, 0x9b, 0x04, 0x57, 0x2a, 0xa5, 0x49, 0xa4, 0x28, 0xd3, 0x72, 0xd4, 0x52, 0x99, 0xd4, 0xd2,
	0x3d, 0x0f, 0xc5, 0x38, 0x94, 0x54, 0xc8, 0x31, 0xbd, 0xce, 0xe8, 0x48, 0x25, 0x63, 0xf4, 0x90,
	0x8a, 0x67, 0xad, 0xe9, 0xa8, 0x95, 0xca, 0x88, 0x8f, 0xf1, 0xb7, 0xc5, 0xb4, 0xa4, 0xa3, 0x68,
	0x40, 0x35, 0x1b, 0x24, 0x9c, 0xf2, 0x54, 0x67, 0x82, 0x8f, 0x57, 0x58, 0x6f, 0xfc, 0xdf, 0x65,
	0xb2, 0x93, 0x4f, 0x49, 0x5f, 0xbf, 0xfc, 0xf0, 0xde, 0xbd, 0x4f, 0xd6, 0x61, 0x25, 0x9a, 0xb2,
	0x11, 0xf7, 0x9c, 0x03, 0xe7, 0x70, 0xbd, 0x5f, 0x07, 0xc3, 0x5b, 0x36, 0xe2, 0xee, 0x3d, 0x52,
	0x53, 0x9c, 0x67, 0x54, 0x44, 0x5e, 0x09, 0xa7, 0xaa, 0x30, 0xec, 0x45, 0xee, 0x3e, 0xa9, 0x0f,
	0x39, 0x0d, 0xf5, 0x8d, 0xe2, 0x5e, 0xf9, 0xc0, 0x39, 0x6c, 0xf4, 0x6b, 0x43, 0x7e, 0x0a, 0x43,
	0x77, 0x87, 0xac, 0xa9, 0x09, 0x44, 0x54, 0xd0, 0x5e, 0x51, 0x13, 0xe3, 0xcf, 0x62, 0x41, 0xd1,
	0x7f, 0xcd, 0xf8, 0xb3, 0x58, 0x9c, 0x81, 0xff, 0x36, 0x29, 0xb3, 0x58, 0x78, 0x55, 0xb4, 0xc2,
	0x5f, 0x40, 0x1a, 0x33, 0x61, 0xbd, 0x6b, 0x68, 0xaf, 0x83, 0x01, 0xdd, 0x7d, 0xd2, 0xc0, 0xc9,
	0x44, 0x86, 0x2c, 0x81, 0x34, 0x75, 0x04, 0xdb, 0x00, 0xe3, 0x1b, 0xb0, 0xf5, 0x22, 0xf7, 0x31,
	0x69, 0xa2, 0x4f, 0x9c, 0xc8, 0x81, 0x71, 0x5a, 0xc7, 0x55, 0x36, 0xc1, 0xfa, 0x0a, 0x8d, 0xbd,
	0xc8, 0x7d, 0x48, 0x30, 0x88, 0xaa, 0x8c, 0x0f, 0xc5, 0xb5, 0x47, 0x70, 0x1d, 0x02, 0xa6, 0x77,
	0x68, 0x71, 0x1f, 0x10, 0x1c, 0x51, 0x16, 0xc2, 0x12, 0x1b, 0x73, 0x90, 0x17, 0x61, 0x2f, 0x02,
	0x4a, 0x3d, 0xa3, 0xdc, 0x34, 0x93, 0x7a, 0x81, 0x52, 0xe7, 0x28, 0x1b, 0x86, 0x52, 0xe7, 0x29,
	0x75, 0x9e, 0xb2, 0x69, 0x28, 0xf5, 0x12, 0xa5, 0x5e, 0xa0, 0xdc, 0x32, 0x94, 0x3a, 0x47, 0xa9,
	0xe7, 0x94, 0xdb, 0x73, 0x10, 0xa0, 0xf4, 0x9f, 0x93, 0xbd, 0xb9, 0xea, 0x6a, 0x42, 0x87, 0x3c,
	0xa4, 0xed, 0xe0, 0x84, 0x6a, 0xf7, 0x11, 0x69, 0xa8, 0x31, 0xbf, 0x8c, 0xe4, 0x44, 0x64, 0x1c,
	0x42, 0x1d, 0x93, 0x7d, 0x6e, 0xec, 0x45, 0xfe, 0x93, 0xc5, 0xa6, 0x51, 0x13, 0xca, 0x62, 0xd1,
	0xa6, 0xda, 0x6d, 0x92, 0x52, 0x66, 0x02, 0x2a, 0xfd, 0x52, 0x16, 0xf9, 0x3f, 0x1c, 0xb2, 0xbd,
	0xec, 0x97, 0xd3, 0xdc, 0x34, 0xd6, 0x4c, 0xf3, 0x2f, 0x0e, 0xa9, 0xc0, 0x52, 0xd8, 0x55, 0x1b,
	0xc1, 0xa8, 0xf5, 0xb7, 0xda, 0xbe, 0xb5, 0x82, 0xbe, 0x8f, 0xa9, 0xfd, 0xa3, 0xe5, 0x4f, 0x13,
	0xf8, 0x69, 0xfb, 0xa4, 0x3e, 0x13, 0xcd, 0x52, 0x27, 0x46, 0x30, 0xff, 0xe3, 0xdd, 0x88, 0xc0,
	0x44, 0xc4, 0x53, 0x05, 0x4d, 0x0d, 0x6b, 0xb1, 0x15, 0x6f, 0x8f, 0x54, 0xad, 0x6e, 0x25, 0x9c,
	0xb0, 0x23, 0xd8, 0x23, 0x46, 0x2e, 0xb3, 0x77, 0x2a, 0x0c, 0xa4, 0xfa, 0x59, 0x5a, 0x2e, 0xa2,
	0x30, 0x45, 0x14, 0x4b, 0x45, 0x14, 0x0b, 0x45, 0x14, 0xc5, 0x17, 0x51, 0xd8, 0x22, 0x0a, 0xd1,
	0x9e, 0x32, 0x04, 0xf8, 0x21, 0x85, 0x32, 0x04, 0x96, 0x21, 0xf0, 0x6f, 0xcb, 0x2b, 0x7b, 0xbc,
	0x4b, 0xb5, 0xfb, 0xc9, 0x9c, 0x2d, 0x0e, 0xc2, 0x5d, 0x14, 0xd7, 0x65, 0xe6, 0x1c, 0xfb, 0x4c,
	0x2a, 0x70, 0x5a, 0x58, 0x7d, 0x2e, 0x8a, 0xab, 0x4d, 0x1f, 0xf3, 0x42, 0x7e, 0x38, 0x07, 0xac,
	0x36, 0x85, 0xe6, 0x87, 0xbc, 0xfe, 0x6d, 0x69, 0x95, 0x30, 0x9d, 0xa3, 0x7f, 0x5d, 0x18, 0xff,
	0x57, 0x99, 0xfc, 0x7f, 0xa7, 0x30, 0x22, 0x1d, 0xca, 0xdc, 0xbd, 0x6a, 0xb7, 0xfb, 0xf4, 0x5e,
	0xfd, 0xea, 0x90, 0xea, 0x90, 0x87, 0xed, 0xe0, 0xc4, 0x72, 0xab, 0x82, 0xb8, 0x67, 0x57, 0x46,
	0xdf, 0xe6, 0x9f, 0xa3, 0x74, 0x6d, 0x6f, 0x15, 0x8b, 0xd2, 0x9d, 0xa1, 0x74, 0x67, 0x28, 0x9d,
	0x23, 0x7c, 0x6f, 0x14, 0x8c, 0x02, 0xbd, 0xdc, 0xb7, 0xf9, 0xfd, 0xef, 0x25, 0xd2, 0xcc, 0xbf,
	0xb0, 0xe0, 0x6e, 0xb6, 0xef, 0x27, 0x7a, 0x9d, 0x79, 0x81, 0x79, 0x5d, 0x99, 0x27, 0xd4, 0x79,
	0xe6, 0x7e, 0x73, 0x40, 0x6d, 0xa3, 0xbc, 0xd7, 0x41, 0x7a, 0x59, 0x20, 0x3d, 0xa4, 0x85, 0xf6,
	0x0a, 0x7b, 0xd0, 0x79, 0x4f, 0xc9, 0x7f, 0x23, 0xa6, 0x94, 0x48, 0x63, 0xaa, 0x93, 0x2b, 0x1a,
	0xca, 0xcb, 0x54, 0x7b, 0xc7, 0x78, 0x3d, 0x6d, 0xd9, 0x89, 0xb3, 0xe4, 0xea, 0x14, 0xcc, 0xee,
	0x31, 0xd9, 0x4b, 0xa5, 0x16, 0x43, 0x11, 0x32, 0x2d, 0x64, 0xba, 0x10, 0xf0, 0x0c, 0x03, 0x76,
	0x17, 0x67, 0xa7, 0x51, 0x83, 0x2a, 0x3e, 0x71, 0x3b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x41, 0xe5, 0x39, 0xfa, 0x0a, 0x00, 0x00,
}