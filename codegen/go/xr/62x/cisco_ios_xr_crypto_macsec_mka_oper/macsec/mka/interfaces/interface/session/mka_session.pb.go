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
// source: mka_session.proto

package cisco_ios_xr_crypto_macsec_mka_oper_macsec_mka_interfaces_interface_session

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

type MkaSession_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MkaSession_KEYS) Reset()         { *m = MkaSession_KEYS{} }
func (m *MkaSession_KEYS) String() string { return proto.CompactTextString(m) }
func (*MkaSession_KEYS) ProtoMessage()    {}
func (*MkaSession_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_87216119eb3785c3, []int{0}
}

func (m *MkaSession_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MkaSession_KEYS.Unmarshal(m, b)
}
func (m *MkaSession_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MkaSession_KEYS.Marshal(b, m, deterministic)
}
func (m *MkaSession_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MkaSession_KEYS.Merge(m, src)
}
func (m *MkaSession_KEYS) XXX_Size() int {
	return xxx_messageInfo_MkaSession_KEYS.Size(m)
}
func (m *MkaSession_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MkaSession_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MkaSession_KEYS proto.InternalMessageInfo

func (m *MkaSession_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MkaVlanTag struct {
	Etype                uint32   `protobuf:"varint,1,opt,name=etype,proto3" json:"etype,omitempty"`
	Priority             uint32   `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Cfi                  uint32   `protobuf:"varint,3,opt,name=cfi,proto3" json:"cfi,omitempty"`
	VlanId               uint32   `protobuf:"varint,4,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MkaVlanTag) Reset()         { *m = MkaVlanTag{} }
func (m *MkaVlanTag) String() string { return proto.CompactTextString(m) }
func (*MkaVlanTag) ProtoMessage()    {}
func (*MkaVlanTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_87216119eb3785c3, []int{1}
}

func (m *MkaVlanTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MkaVlanTag.Unmarshal(m, b)
}
func (m *MkaVlanTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MkaVlanTag.Marshal(b, m, deterministic)
}
func (m *MkaVlanTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MkaVlanTag.Merge(m, src)
}
func (m *MkaVlanTag) XXX_Size() int {
	return xxx_messageInfo_MkaVlanTag.Size(m)
}
func (m *MkaVlanTag) XXX_DiscardUnknown() {
	xxx_messageInfo_MkaVlanTag.DiscardUnknown(m)
}

var xxx_messageInfo_MkaVlanTag proto.InternalMessageInfo

func (m *MkaVlanTag) GetEtype() uint32 {
	if m != nil {
		return m.Etype
	}
	return 0
}

func (m *MkaVlanTag) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *MkaVlanTag) GetCfi() uint32 {
	if m != nil {
		return m.Cfi
	}
	return 0
}

func (m *MkaVlanTag) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

type MkaSessionSummary struct {
	InterfaceName         string      `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	InheritedPolicy       bool        `protobuf:"varint,2,opt,name=inherited_policy,json=inheritedPolicy,proto3" json:"inherited_policy,omitempty"`
	Policy                string      `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
	Priority              uint32      `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	MyMac                 string      `protobuf:"bytes,5,opt,name=my_mac,json=myMac,proto3" json:"my_mac,omitempty"`
	DelayProtection       bool        `protobuf:"varint,6,opt,name=delay_protection,json=delayProtection,proto3" json:"delay_protection,omitempty"`
	ReplayProtect         bool        `protobuf:"varint,7,opt,name=replay_protect,json=replayProtect,proto3" json:"replay_protect,omitempty"`
	WindowSize            uint32      `protobuf:"varint,8,opt,name=window_size,json=windowSize,proto3" json:"window_size,omitempty"`
	IncludeIcvIndicator   bool        `protobuf:"varint,9,opt,name=include_icv_indicator,json=includeIcvIndicator,proto3" json:"include_icv_indicator,omitempty"`
	ConfidentialityOffset uint32      `protobuf:"varint,10,opt,name=confidentiality_offset,json=confidentialityOffset,proto3" json:"confidentiality_offset,omitempty"`
	AlgoAgility           uint32      `protobuf:"varint,11,opt,name=algo_agility,json=algoAgility,proto3" json:"algo_agility,omitempty"`
	Capability            uint32      `protobuf:"varint,12,opt,name=capability,proto3" json:"capability,omitempty"`
	CipherStr             string      `protobuf:"bytes,13,opt,name=cipher_str,json=cipherStr,proto3" json:"cipher_str,omitempty"`
	MacSecDesired         bool        `protobuf:"varint,14,opt,name=mac_sec_desired,json=macSecDesired,proto3" json:"mac_sec_desired,omitempty"`
	OuterTag              *MkaVlanTag `protobuf:"bytes,15,opt,name=outer_tag,json=outerTag,proto3" json:"outer_tag,omitempty"`
	InnerTag              *MkaVlanTag `protobuf:"bytes,16,opt,name=inner_tag,json=innerTag,proto3" json:"inner_tag,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}    `json:"-"`
	XXX_unrecognized      []byte      `json:"-"`
	XXX_sizecache         int32       `json:"-"`
}

func (m *MkaSessionSummary) Reset()         { *m = MkaSessionSummary{} }
func (m *MkaSessionSummary) String() string { return proto.CompactTextString(m) }
func (*MkaSessionSummary) ProtoMessage()    {}
func (*MkaSessionSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_87216119eb3785c3, []int{2}
}

func (m *MkaSessionSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MkaSessionSummary.Unmarshal(m, b)
}
func (m *MkaSessionSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MkaSessionSummary.Marshal(b, m, deterministic)
}
func (m *MkaSessionSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MkaSessionSummary.Merge(m, src)
}
func (m *MkaSessionSummary) XXX_Size() int {
	return xxx_messageInfo_MkaSessionSummary.Size(m)
}
func (m *MkaSessionSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_MkaSessionSummary.DiscardUnknown(m)
}

var xxx_messageInfo_MkaSessionSummary proto.InternalMessageInfo

func (m *MkaSessionSummary) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *MkaSessionSummary) GetInheritedPolicy() bool {
	if m != nil {
		return m.InheritedPolicy
	}
	return false
}

func (m *MkaSessionSummary) GetPolicy() string {
	if m != nil {
		return m.Policy
	}
	return ""
}

func (m *MkaSessionSummary) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *MkaSessionSummary) GetMyMac() string {
	if m != nil {
		return m.MyMac
	}
	return ""
}

func (m *MkaSessionSummary) GetDelayProtection() bool {
	if m != nil {
		return m.DelayProtection
	}
	return false
}

func (m *MkaSessionSummary) GetReplayProtect() bool {
	if m != nil {
		return m.ReplayProtect
	}
	return false
}

func (m *MkaSessionSummary) GetWindowSize() uint32 {
	if m != nil {
		return m.WindowSize
	}
	return 0
}

func (m *MkaSessionSummary) GetIncludeIcvIndicator() bool {
	if m != nil {
		return m.IncludeIcvIndicator
	}
	return false
}

func (m *MkaSessionSummary) GetConfidentialityOffset() uint32 {
	if m != nil {
		return m.ConfidentialityOffset
	}
	return 0
}

func (m *MkaSessionSummary) GetAlgoAgility() uint32 {
	if m != nil {
		return m.AlgoAgility
	}
	return 0
}

func (m *MkaSessionSummary) GetCapability() uint32 {
	if m != nil {
		return m.Capability
	}
	return 0
}

func (m *MkaSessionSummary) GetCipherStr() string {
	if m != nil {
		return m.CipherStr
	}
	return ""
}

func (m *MkaSessionSummary) GetMacSecDesired() bool {
	if m != nil {
		return m.MacSecDesired
	}
	return false
}

func (m *MkaSessionSummary) GetOuterTag() *MkaVlanTag {
	if m != nil {
		return m.OuterTag
	}
	return nil
}

func (m *MkaSessionSummary) GetInnerTag() *MkaVlanTag {
	if m != nil {
		return m.InnerTag
	}
	return nil
}

type VpData struct {
	MySci                string   `protobuf:"bytes,1,opt,name=my_sci,json=mySci,proto3" json:"my_sci,omitempty"`
	VirtualPortId        uint32   `protobuf:"varint,2,opt,name=virtual_port_id,json=virtualPortId,proto3" json:"virtual_port_id,omitempty"`
	LatestRx             bool     `protobuf:"varint,3,opt,name=latest_rx,json=latestRx,proto3" json:"latest_rx,omitempty"`
	LatestTx             bool     `protobuf:"varint,4,opt,name=latest_tx,json=latestTx,proto3" json:"latest_tx,omitempty"`
	LatestAn             uint32   `protobuf:"varint,5,opt,name=latest_an,json=latestAn,proto3" json:"latest_an,omitempty"`
	LatestKi             string   `protobuf:"bytes,6,opt,name=latest_ki,json=latestKi,proto3" json:"latest_ki,omitempty"`
	LatestKn             uint32   `protobuf:"varint,7,opt,name=latest_kn,json=latestKn,proto3" json:"latest_kn,omitempty"`
	OldRx                bool     `protobuf:"varint,8,opt,name=old_rx,json=oldRx,proto3" json:"old_rx,omitempty"`
	OldTx                bool     `protobuf:"varint,9,opt,name=old_tx,json=oldTx,proto3" json:"old_tx,omitempty"`
	OldAn                uint32   `protobuf:"varint,10,opt,name=old_an,json=oldAn,proto3" json:"old_an,omitempty"`
	OldKi                string   `protobuf:"bytes,11,opt,name=old_ki,json=oldKi,proto3" json:"old_ki,omitempty"`
	OldKn                uint32   `protobuf:"varint,12,opt,name=old_kn,json=oldKn,proto3" json:"old_kn,omitempty"`
	WaitTime             uint32   `protobuf:"varint,13,opt,name=wait_time,json=waitTime,proto3" json:"wait_time,omitempty"`
	RetireTime           uint32   `protobuf:"varint,14,opt,name=retire_time,json=retireTime,proto3" json:"retire_time,omitempty"`
	CipherSuite          uint32   `protobuf:"varint,15,opt,name=cipher_suite,json=cipherSuite,proto3" json:"cipher_suite,omitempty"`
	Ssci                 uint32   `protobuf:"varint,16,opt,name=ssci,proto3" json:"ssci,omitempty"`
	TimeToSakRekey       string   `protobuf:"bytes,17,opt,name=time_to_sak_rekey,json=timeToSakRekey,proto3" json:"time_to_sak_rekey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VpData) Reset()         { *m = VpData{} }
func (m *VpData) String() string { return proto.CompactTextString(m) }
func (*VpData) ProtoMessage()    {}
func (*VpData) Descriptor() ([]byte, []int) {
	return fileDescriptor_87216119eb3785c3, []int{3}
}

func (m *VpData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VpData.Unmarshal(m, b)
}
func (m *VpData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VpData.Marshal(b, m, deterministic)
}
func (m *VpData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VpData.Merge(m, src)
}
func (m *VpData) XXX_Size() int {
	return xxx_messageInfo_VpData.Size(m)
}
func (m *VpData) XXX_DiscardUnknown() {
	xxx_messageInfo_VpData.DiscardUnknown(m)
}

var xxx_messageInfo_VpData proto.InternalMessageInfo

func (m *VpData) GetMySci() string {
	if m != nil {
		return m.MySci
	}
	return ""
}

func (m *VpData) GetVirtualPortId() uint32 {
	if m != nil {
		return m.VirtualPortId
	}
	return 0
}

func (m *VpData) GetLatestRx() bool {
	if m != nil {
		return m.LatestRx
	}
	return false
}

func (m *VpData) GetLatestTx() bool {
	if m != nil {
		return m.LatestTx
	}
	return false
}

func (m *VpData) GetLatestAn() uint32 {
	if m != nil {
		return m.LatestAn
	}
	return 0
}

func (m *VpData) GetLatestKi() string {
	if m != nil {
		return m.LatestKi
	}
	return ""
}

func (m *VpData) GetLatestKn() uint32 {
	if m != nil {
		return m.LatestKn
	}
	return 0
}

func (m *VpData) GetOldRx() bool {
	if m != nil {
		return m.OldRx
	}
	return false
}

func (m *VpData) GetOldTx() bool {
	if m != nil {
		return m.OldTx
	}
	return false
}

func (m *VpData) GetOldAn() uint32 {
	if m != nil {
		return m.OldAn
	}
	return 0
}

func (m *VpData) GetOldKi() string {
	if m != nil {
		return m.OldKi
	}
	return ""
}

func (m *VpData) GetOldKn() uint32 {
	if m != nil {
		return m.OldKn
	}
	return 0
}

func (m *VpData) GetWaitTime() uint32 {
	if m != nil {
		return m.WaitTime
	}
	return 0
}

func (m *VpData) GetRetireTime() uint32 {
	if m != nil {
		return m.RetireTime
	}
	return 0
}

func (m *VpData) GetCipherSuite() uint32 {
	if m != nil {
		return m.CipherSuite
	}
	return 0
}

func (m *VpData) GetSsci() uint32 {
	if m != nil {
		return m.Ssci
	}
	return 0
}

func (m *VpData) GetTimeToSakRekey() string {
	if m != nil {
		return m.TimeToSakRekey
	}
	return ""
}

type PeerData struct {
	Mi                   string   `protobuf:"bytes,1,opt,name=mi,proto3" json:"mi,omitempty"`
	Sci                  string   `protobuf:"bytes,2,opt,name=sci,proto3" json:"sci,omitempty"`
	Mn                   uint32   `protobuf:"varint,3,opt,name=mn,proto3" json:"mn,omitempty"`
	Priority             uint32   `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	Ssci                 uint32   `protobuf:"varint,5,opt,name=ssci,proto3" json:"ssci,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerData) Reset()         { *m = PeerData{} }
func (m *PeerData) String() string { return proto.CompactTextString(m) }
func (*PeerData) ProtoMessage()    {}
func (*PeerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_87216119eb3785c3, []int{4}
}

func (m *PeerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerData.Unmarshal(m, b)
}
func (m *PeerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerData.Marshal(b, m, deterministic)
}
func (m *PeerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerData.Merge(m, src)
}
func (m *PeerData) XXX_Size() int {
	return xxx_messageInfo_PeerData.Size(m)
}
func (m *PeerData) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerData.DiscardUnknown(m)
}

var xxx_messageInfo_PeerData proto.InternalMessageInfo

func (m *PeerData) GetMi() string {
	if m != nil {
		return m.Mi
	}
	return ""
}

func (m *PeerData) GetSci() string {
	if m != nil {
		return m.Sci
	}
	return ""
}

func (m *PeerData) GetMn() uint32 {
	if m != nil {
		return m.Mn
	}
	return 0
}

func (m *PeerData) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *PeerData) GetSsci() uint32 {
	if m != nil {
		return m.Ssci
	}
	return 0
}

type CaData struct {
	IsKeyServer           bool        `protobuf:"varint,1,opt,name=is_key_server,json=isKeyServer,proto3" json:"is_key_server,omitempty"`
	Status                uint32      `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	NumLivePeers          uint32      `protobuf:"varint,3,opt,name=num_live_peers,json=numLivePeers,proto3" json:"num_live_peers,omitempty"`
	FirstCa               bool        `protobuf:"varint,4,opt,name=first_ca,json=firstCa,proto3" json:"first_ca,omitempty"`
	PeerSci               string      `protobuf:"bytes,5,opt,name=peer_sci,json=peerSci,proto3" json:"peer_sci,omitempty"`
	NumLivePeersResponded uint32      `protobuf:"varint,6,opt,name=num_live_peers_responded,json=numLivePeersResponded,proto3" json:"num_live_peers_responded,omitempty"`
	Ckn                   string      `protobuf:"bytes,7,opt,name=ckn,proto3" json:"ckn,omitempty"`
	MyMi                  string      `protobuf:"bytes,8,opt,name=my_mi,json=myMi,proto3" json:"my_mi,omitempty"`
	MyMn                  uint32      `protobuf:"varint,9,opt,name=my_mn,json=myMn,proto3" json:"my_mn,omitempty"`
	Authenticator         bool        `protobuf:"varint,10,opt,name=authenticator,proto3" json:"authenticator,omitempty"`
	LivePeer              []*PeerData `protobuf:"bytes,11,rep,name=live_peer,json=livePeer,proto3" json:"live_peer,omitempty"`
	PotentialPeer         []*PeerData `protobuf:"bytes,12,rep,name=potential_peer,json=potentialPeer,proto3" json:"potential_peer,omitempty"`
	DormantPeer           []*PeerData `protobuf:"bytes,13,rep,name=dormant_peer,json=dormantPeer,proto3" json:"dormant_peer,omitempty"`
	StatusDescription     string      `protobuf:"bytes,14,opt,name=status_description,json=statusDescription,proto3" json:"status_description,omitempty"`
	AuthenticationMode    string      `protobuf:"bytes,15,opt,name=authentication_mode,json=authenticationMode,proto3" json:"authentication_mode,omitempty"`
	KeyChain              string      `protobuf:"bytes,16,opt,name=key_chain,json=keyChain,proto3" json:"key_chain,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}    `json:"-"`
	XXX_unrecognized      []byte      `json:"-"`
	XXX_sizecache         int32       `json:"-"`
}

func (m *CaData) Reset()         { *m = CaData{} }
func (m *CaData) String() string { return proto.CompactTextString(m) }
func (*CaData) ProtoMessage()    {}
func (*CaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_87216119eb3785c3, []int{5}
}

func (m *CaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaData.Unmarshal(m, b)
}
func (m *CaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaData.Marshal(b, m, deterministic)
}
func (m *CaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaData.Merge(m, src)
}
func (m *CaData) XXX_Size() int {
	return xxx_messageInfo_CaData.Size(m)
}
func (m *CaData) XXX_DiscardUnknown() {
	xxx_messageInfo_CaData.DiscardUnknown(m)
}

var xxx_messageInfo_CaData proto.InternalMessageInfo

func (m *CaData) GetIsKeyServer() bool {
	if m != nil {
		return m.IsKeyServer
	}
	return false
}

func (m *CaData) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CaData) GetNumLivePeers() uint32 {
	if m != nil {
		return m.NumLivePeers
	}
	return 0
}

func (m *CaData) GetFirstCa() bool {
	if m != nil {
		return m.FirstCa
	}
	return false
}

func (m *CaData) GetPeerSci() string {
	if m != nil {
		return m.PeerSci
	}
	return ""
}

func (m *CaData) GetNumLivePeersResponded() uint32 {
	if m != nil {
		return m.NumLivePeersResponded
	}
	return 0
}

func (m *CaData) GetCkn() string {
	if m != nil {
		return m.Ckn
	}
	return ""
}

func (m *CaData) GetMyMi() string {
	if m != nil {
		return m.MyMi
	}
	return ""
}

func (m *CaData) GetMyMn() uint32 {
	if m != nil {
		return m.MyMn
	}
	return 0
}

func (m *CaData) GetAuthenticator() bool {
	if m != nil {
		return m.Authenticator
	}
	return false
}

func (m *CaData) GetLivePeer() []*PeerData {
	if m != nil {
		return m.LivePeer
	}
	return nil
}

func (m *CaData) GetPotentialPeer() []*PeerData {
	if m != nil {
		return m.PotentialPeer
	}
	return nil
}

func (m *CaData) GetDormantPeer() []*PeerData {
	if m != nil {
		return m.DormantPeer
	}
	return nil
}

func (m *CaData) GetStatusDescription() string {
	if m != nil {
		return m.StatusDescription
	}
	return ""
}

func (m *CaData) GetAuthenticationMode() string {
	if m != nil {
		return m.AuthenticationMode
	}
	return ""
}

func (m *CaData) GetKeyChain() string {
	if m != nil {
		return m.KeyChain
	}
	return ""
}

type MkaSession struct {
	SessionSummary       *MkaSessionSummary `protobuf:"bytes,50,opt,name=session_summary,json=sessionSummary,proto3" json:"session_summary,omitempty"`
	Vp                   *VpData            `protobuf:"bytes,51,opt,name=vp,proto3" json:"vp,omitempty"`
	Ca                   []*CaData          `protobuf:"bytes,52,rep,name=ca,proto3" json:"ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MkaSession) Reset()         { *m = MkaSession{} }
func (m *MkaSession) String() string { return proto.CompactTextString(m) }
func (*MkaSession) ProtoMessage()    {}
func (*MkaSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_87216119eb3785c3, []int{6}
}

func (m *MkaSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MkaSession.Unmarshal(m, b)
}
func (m *MkaSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MkaSession.Marshal(b, m, deterministic)
}
func (m *MkaSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MkaSession.Merge(m, src)
}
func (m *MkaSession) XXX_Size() int {
	return xxx_messageInfo_MkaSession.Size(m)
}
func (m *MkaSession) XXX_DiscardUnknown() {
	xxx_messageInfo_MkaSession.DiscardUnknown(m)
}

var xxx_messageInfo_MkaSession proto.InternalMessageInfo

func (m *MkaSession) GetSessionSummary() *MkaSessionSummary {
	if m != nil {
		return m.SessionSummary
	}
	return nil
}

func (m *MkaSession) GetVp() *VpData {
	if m != nil {
		return m.Vp
	}
	return nil
}

func (m *MkaSession) GetCa() []*CaData {
	if m != nil {
		return m.Ca
	}
	return nil
}

func init() {
	proto.RegisterType((*MkaSession_KEYS)(nil), "cisco_ios_xr_crypto_macsec_mka_oper.macsec.mka.interfaces.interface.session.mka_session_KEYS")
	proto.RegisterType((*MkaVlanTag)(nil), "cisco_ios_xr_crypto_macsec_mka_oper.macsec.mka.interfaces.interface.session.mka_vlan_tag")
	proto.RegisterType((*MkaSessionSummary)(nil), "cisco_ios_xr_crypto_macsec_mka_oper.macsec.mka.interfaces.interface.session.mka_session_summary")
	proto.RegisterType((*VpData)(nil), "cisco_ios_xr_crypto_macsec_mka_oper.macsec.mka.interfaces.interface.session.vp_data")
	proto.RegisterType((*PeerData)(nil), "cisco_ios_xr_crypto_macsec_mka_oper.macsec.mka.interfaces.interface.session.peer_data")
	proto.RegisterType((*CaData)(nil), "cisco_ios_xr_crypto_macsec_mka_oper.macsec.mka.interfaces.interface.session.ca_data")
	proto.RegisterType((*MkaSession)(nil), "cisco_ios_xr_crypto_macsec_mka_oper.macsec.mka.interfaces.interface.session.mka_session")
}

func init() { proto.RegisterFile("mka_session.proto", fileDescriptor_87216119eb3785c3) }

var fileDescriptor_87216119eb3785c3 = []byte{
	// 1109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xc1, 0x6e, 0x1b, 0x37,
	0x10, 0x85, 0x65, 0x4b, 0xd6, 0x52, 0x5a, 0xd9, 0xa6, 0xeb, 0x74, 0xdb, 0xa2, 0x6d, 0x22, 0xa4,
	0x41, 0x72, 0xa8, 0x0a, 0x38, 0x2d, 0x7a, 0x0e, 0x92, 0x1e, 0x0c, 0x37, 0xad, 0xb1, 0x12, 0x0a,
	0xe4, 0xc4, 0x32, 0x5c, 0xda, 0x1e, 0x68, 0x97, 0x5c, 0x90, 0x94, 0xa2, 0x0d, 0xd0, 0x0f, 0x68,
	0xbf, 0xa4, 0x3f, 0xd6, 0x5b, 0x3f, 0xa2, 0xe0, 0x90, 0x5a, 0xaf, 0x7d, 0xe8, 0xc9, 0xbe, 0x91,
	0xef, 0xcd, 0xce, 0x1b, 0x2e, 0x1f, 0x87, 0x24, 0x47, 0xd5, 0x92, 0x33, 0x2b, 0xad, 0x05, 0xad,
	0x66, 0xb5, 0xd1, 0x4e, 0xd3, 0x73, 0x01, 0x56, 0x68, 0x06, 0xda, 0xb2, 0x8d, 0x61, 0xc2, 0x34,
	0xb5, 0xd3, 0xac, 0xe2, 0xc2, 0x4a, 0xc1, 0x7c, 0xb4, 0xae, 0xa5, 0x99, 0x85, 0xf9, 0xac, 0x5a,
	0xf2, 0x19, 0x28, 0x27, 0xcd, 0x25, 0x17, 0xd2, 0xde, 0x0c, 0x67, 0x31, 0xe5, 0xf4, 0x19, 0x39,
	0xec, 0x28, 0xb0, 0xf3, 0x9f, 0xde, 0xcd, 0x29, 0x25, 0x7b, 0x8a, 0x57, 0x32, 0xdb, 0x79, 0xbc,
	0xf3, 0x3c, 0xc9, 0x71, 0x3c, 0x5d, 0x92, 0xb1, 0x8f, 0x5b, 0x97, 0x5c, 0x31, 0xc7, 0xaf, 0xe8,
	0x27, 0xa4, 0x2f, 0x5d, 0x53, 0x87, 0xa0, 0x34, 0x0f, 0x13, 0xfa, 0x39, 0x19, 0xd6, 0x06, 0xb4,
	0x01, 0xd7, 0x64, 0x3d, 0x24, 0xda, 0x39, 0x3d, 0x24, 0xbb, 0xe2, 0x12, 0xb2, 0x5d, 0x84, 0xfd,
	0x90, 0x7e, 0x4a, 0xf6, 0x31, 0x1f, 0x14, 0xd9, 0x1e, 0xa2, 0x03, 0x3f, 0x3d, 0x2b, 0xa6, 0x7f,
	0x0d, 0xc8, 0x71, 0xb7, 0x2a, 0xbb, 0xaa, 0x2a, 0x6e, 0x1a, 0xfa, 0x0d, 0x99, 0xb4, 0x2b, 0x60,
	0x9d, 0x12, 0xd3, 0x16, 0xfd, 0x85, 0x57, 0x92, 0xbe, 0x20, 0x87, 0xa0, 0xae, 0xa5, 0x01, 0x27,
	0x0b, 0x56, 0xeb, 0x12, 0x44, 0xa8, 0x66, 0x98, 0x1f, 0xb4, 0xf8, 0x05, 0xc2, 0xf4, 0x11, 0x19,
	0xc4, 0x80, 0x5d, 0xcc, 0x14, 0x67, 0xb7, 0x16, 0xb2, 0x77, 0x67, 0x21, 0x27, 0x64, 0x50, 0x35,
	0xfe, 0x87, 0x67, 0x7d, 0xfc, 0xa6, 0x5f, 0x35, 0x6f, 0xb9, 0xf0, 0xaa, 0x85, 0x2c, 0x79, 0xc3,
	0xfc, 0x2e, 0x49, 0xe1, 0x40, 0xab, 0x6c, 0x10, 0x54, 0x11, 0xbf, 0x68, 0x61, 0xbf, 0x0e, 0x23,
	0xeb, 0x4e, 0x6c, 0xb6, 0x8f, 0x81, 0x69, 0x40, 0x63, 0x24, 0xfd, 0x9a, 0x8c, 0x3e, 0x80, 0x2a,
	0xf4, 0x07, 0x66, 0xe1, 0xa3, 0xcc, 0x86, 0x58, 0x07, 0x09, 0xd0, 0x1c, 0x3e, 0x4a, 0x7a, 0x4a,
	0x4e, 0x40, 0x89, 0x72, 0x55, 0x48, 0x06, 0x62, 0xcd, 0x40, 0x15, 0x20, 0xb8, 0xd3, 0x26, 0x4b,
	0x30, 0xdd, 0x71, 0x24, 0xcf, 0xc4, 0xfa, 0x6c, 0x4b, 0xd1, 0x1f, 0xc8, 0x23, 0xa1, 0xd5, 0x25,
	0x14, 0x52, 0x39, 0xe0, 0x25, 0xb8, 0x86, 0xe9, 0xcb, 0x4b, 0x2b, 0x5d, 0x46, 0x30, 0xff, 0xc9,
	0x1d, 0xf6, 0x57, 0x24, 0xe9, 0x13, 0x32, 0xe6, 0xe5, 0x95, 0x66, 0xfc, 0x0a, 0x3c, 0x9a, 0x8d,
	0x30, 0x78, 0xe4, 0xb1, 0x57, 0x01, 0xa2, 0x5f, 0x11, 0x22, 0x78, 0xcd, 0xdf, 0x87, 0x80, 0x71,
	0xa8, 0xf6, 0x06, 0xa1, 0x5f, 0x12, 0x22, 0xa0, 0xbe, 0x96, 0x86, 0x59, 0x67, 0xb2, 0x14, 0xff,
	0x5d, 0x12, 0x90, 0xb9, 0x33, 0xf4, 0x19, 0x39, 0xa8, 0xb8, 0x60, 0xde, 0xc5, 0x85, 0xb4, 0x60,
	0x64, 0x91, 0x4d, 0xc2, 0x5f, 0xa9, 0xb8, 0x98, 0x4b, 0xf1, 0x26, 0x80, 0x74, 0x4d, 0x12, 0xbd,
	0x72, 0xd2, 0x78, 0x1b, 0x66, 0x07, 0x8f, 0x77, 0x9e, 0x8f, 0x4e, 0xdf, 0xcd, 0xee, 0xf1, 0x48,
	0xcc, 0xba, 0x3e, 0xcf, 0x87, 0xa8, 0xb5, 0xe0, 0x57, 0x5e, 0x17, 0x94, 0x8a, 0xba, 0x87, 0x0f,
	0xae, 0x8b, 0x5a, 0x0b, 0x7e, 0x35, 0xfd, 0x77, 0x97, 0xec, 0xaf, 0x6b, 0x56, 0x70, 0xc7, 0xa3,
	0xf5, 0xac, 0x80, 0x68, 0xfc, 0x7e, 0xd5, 0xcc, 0x05, 0xf8, 0x5f, 0xb7, 0x06, 0xe3, 0x56, 0xbc,
	0x64, 0xb5, 0x36, 0xce, 0x1f, 0xa8, 0x70, 0xfa, 0xd2, 0x08, 0x5f, 0x68, 0xe3, 0xce, 0x0a, 0xfa,
	0x05, 0x49, 0x4a, 0xee, 0xa4, 0x75, 0xcc, 0x6c, 0xd0, 0xf0, 0xc3, 0x7c, 0x18, 0x80, 0x7c, 0xd3,
	0x21, 0xdd, 0x06, 0x3d, 0xdf, 0x92, 0x8b, 0x2e, 0xc9, 0x15, 0xda, 0x3e, 0xdd, 0x92, 0xaf, 0x54,
	0x87, 0x5c, 0x02, 0x5a, 0x3e, 0xd9, 0x92, 0xe7, 0xd0, 0x25, 0x15, 0xda, 0xbc, 0xfd, 0xf2, 0x5c,
	0xf9, 0xf5, 0xe8, 0xb2, 0xf0, 0xd5, 0x0c, 0x51, 0xb0, 0xaf, 0xcb, 0x22, 0xdf, 0x6c, 0x61, 0xb7,
	0x89, 0x46, 0xf6, 0xf0, 0xa2, 0x85, 0xb9, 0x8a, 0x56, 0xf5, 0xf0, 0xab, 0x36, 0xc9, 0x12, 0xd0,
	0x94, 0x09, 0xc2, 0xe7, 0xd0, 0xc2, 0x2a, 0x5a, 0x11, 0x61, 0x2c, 0xf6, 0x03, 0x07, 0xc7, 0x1c,
	0x54, 0x12, 0x4d, 0x98, 0xe6, 0x43, 0x0f, 0x2c, 0xa0, 0x92, 0xfe, 0xc4, 0x19, 0xe9, 0xc0, 0xc8,
	0x40, 0x4f, 0x82, 0x87, 0x03, 0x84, 0x01, 0x4f, 0xc8, 0x78, 0xeb, 0xe1, 0x15, 0x38, 0x89, 0xfe,
	0x4b, 0xf3, 0x51, 0x74, 0xb1, 0x87, 0x7c, 0xf7, 0xb4, 0x7e, 0x87, 0x0e, 0x91, 0xc2, 0x31, 0x7d,
	0x41, 0x8e, 0x7c, 0x42, 0xe6, 0x34, 0xb3, 0x7c, 0xc9, 0x8c, 0x5c, 0xca, 0x26, 0x3b, 0xc2, 0x6a,
	0x27, 0x9e, 0x58, 0xe8, 0x39, 0x5f, 0xe6, 0x1e, 0x9d, 0x56, 0x24, 0xa9, 0xa5, 0x34, 0x61, 0xbf,
	0x27, 0xa4, 0x57, 0x6d, 0xf7, 0xba, 0x57, 0x81, 0xef, 0xa1, 0x3e, 0x75, 0x0f, 0x01, 0x3f, 0xc4,
	0x08, 0x15, 0x9b, 0x6a, 0xaf, 0x52, 0xff, 0xdb, 0xb8, 0xb6, 0x95, 0xf5, 0x6f, 0x2a, 0x9b, 0xfe,
	0x3d, 0x20, 0xfb, 0x82, 0x07, 0xb5, 0x29, 0x49, 0xc1, 0xb2, 0xa5, 0x6c, 0x98, 0x95, 0x66, 0x2d,
	0x0d, 0x0a, 0x0f, 0xf3, 0x11, 0xd8, 0x73, 0xd9, 0xcc, 0x11, 0xf2, 0x0d, 0xd3, 0x3a, 0xee, 0x56,
	0x36, 0x3a, 0x2c, 0xce, 0xe8, 0x53, 0x32, 0x51, 0xab, 0x8a, 0x95, 0xb0, 0x96, 0xcc, 0xd7, 0x6f,
	0x63, 0x4d, 0x63, 0xb5, 0xaa, 0x7e, 0x86, 0xb5, 0xbc, 0xf0, 0x18, 0xfd, 0x8c, 0x0c, 0x2f, 0xc1,
	0x58, 0xc7, 0x04, 0x8f, 0x16, 0xdb, 0xc7, 0xf9, 0x6b, 0xee, 0x29, 0x5c, 0xf7, 0xb6, 0xc0, 0x24,
	0xdf, 0xf7, 0x73, 0x6f, 0xef, 0x1f, 0x49, 0x76, 0x3b, 0x37, 0x33, 0xd2, 0xd6, 0x5a, 0x15, 0xb2,
	0x40, 0xbb, 0xa5, 0xf9, 0x49, 0x57, 0x25, 0xdf, 0x92, 0x78, 0xe5, 0x44, 0xd7, 0x25, 0xb9, 0x1f,
	0xd2, 0x63, 0xd2, 0xf7, 0xbd, 0x1b, 0xd0, 0x6f, 0x49, 0xbe, 0x57, 0x35, 0x6f, 0x61, 0x0b, 0x2a,
	0x74, 0x5b, 0x8a, 0xa0, 0xa2, 0x4f, 0x49, 0xca, 0x57, 0xee, 0xda, 0xf7, 0xc1, 0xd0, 0x53, 0x49,
	0x68, 0x46, 0xb7, 0x40, 0x6a, 0x49, 0xd2, 0x96, 0x95, 0x8d, 0x1e, 0xef, 0x3e, 0x1f, 0x9d, 0xfe,
	0x76, 0xaf, 0x4d, 0xa1, 0xf5, 0x42, 0x3e, 0x2c, 0xe3, 0x02, 0xe9, 0x1f, 0x64, 0x52, 0x6b, 0x17,
	0x3a, 0x74, 0x50, 0x1e, 0x3f, 0xa8, 0x72, 0xda, 0xaa, 0xa1, 0x7c, 0x43, 0xc6, 0x85, 0x36, 0x15,
	0x57, 0x2e, 0x88, 0xa7, 0x0f, 0x2a, 0x3e, 0x8a, 0x5a, 0x28, 0xfd, 0x2d, 0xa1, 0xc1, 0x6f, 0xfe,
	0x8a, 0x10, 0x06, 0x6a, 0xbc, 0x65, 0x27, 0xb8, 0x97, 0x47, 0x81, 0x79, 0x73, 0x43, 0xd0, 0xef,
	0xc8, 0x71, 0x67, 0xbb, 0xfc, 0x4b, 0xa2, 0xd2, 0x45, 0x38, 0xb4, 0x49, 0x4e, 0x6f, 0x53, 0x6f,
	0x75, 0x21, 0x7d, 0x73, 0xf0, 0xf6, 0x17, 0xd7, 0x1c, 0x14, 0x1e, 0xe0, 0x24, 0x1f, 0x2e, 0x65,
	0xf3, 0xda, 0xcf, 0xa7, 0xff, 0xf4, 0xc8, 0xa8, 0xf3, 0x2a, 0xa1, 0x7f, 0xee, 0x90, 0x83, 0x3b,
	0x2f, 0x94, 0xec, 0x14, 0xef, 0x85, 0xdf, 0xef, 0xfd, 0x5e, 0xb8, 0xa3, 0x93, 0x4f, 0x22, 0x30,
	0x8f, 0x2f, 0xa3, 0x82, 0xf4, 0xd6, 0x75, 0xf6, 0x12, 0xd5, 0x17, 0xf7, 0xaa, 0x1e, 0xaf, 0x9e,
	0xbc, 0xb7, 0xae, 0xbd, 0x8a, 0xe0, 0xd9, 0xf7, 0xb8, 0xdf, 0xf7, 0xab, 0x12, 0x5b, 0x50, 0xde,
	0x13, 0xfc, 0xfd, 0x00, 0x9f, 0xb9, 0x2f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x19, 0xff, 0x42,
	0x09, 0xfb, 0x0a, 0x00, 0x00,
}