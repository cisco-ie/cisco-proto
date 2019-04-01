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
// source: dot1x_interface_detail.proto

package cisco_ios_xr_dot1x_oper_dot1x_session_interface_sessions_interface_session

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

type Dot1XInterfaceDetail_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XInterfaceDetail_KEYS) Reset()         { *m = Dot1XInterfaceDetail_KEYS{} }
func (m *Dot1XInterfaceDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*Dot1XInterfaceDetail_KEYS) ProtoMessage()    {}
func (*Dot1XInterfaceDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f160d63776e65b5, []int{0}
}

func (m *Dot1XInterfaceDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XInterfaceDetail_KEYS.Unmarshal(m, b)
}
func (m *Dot1XInterfaceDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XInterfaceDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *Dot1XInterfaceDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XInterfaceDetail_KEYS.Merge(m, src)
}
func (m *Dot1XInterfaceDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_Dot1XInterfaceDetail_KEYS.Size(m)
}
func (m *Dot1XInterfaceDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XInterfaceDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XInterfaceDetail_KEYS proto.InternalMessageInfo

func (m *Dot1XInterfaceDetail_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Dot1XAuthClientInfo struct {
	Mac                  string   `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	AuthSmState          string   `protobuf:"bytes,2,opt,name=auth_sm_state,json=authSmState,proto3" json:"auth_sm_state,omitempty"`
	AuthBendSmState      string   `protobuf:"bytes,3,opt,name=auth_bend_sm_state,json=authBendSmState,proto3" json:"auth_bend_sm_state,omitempty"`
	TimeToNextReauth     string   `protobuf:"bytes,4,opt,name=time_to_next_reauth,json=timeToNextReauth,proto3" json:"time_to_next_reauth,omitempty"`
	LastAuthTime         string   `protobuf:"bytes,5,opt,name=last_auth_time,json=lastAuthTime,proto3" json:"last_auth_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XAuthClientInfo) Reset()         { *m = Dot1XAuthClientInfo{} }
func (m *Dot1XAuthClientInfo) String() string { return proto.CompactTextString(m) }
func (*Dot1XAuthClientInfo) ProtoMessage()    {}
func (*Dot1XAuthClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f160d63776e65b5, []int{1}
}

func (m *Dot1XAuthClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XAuthClientInfo.Unmarshal(m, b)
}
func (m *Dot1XAuthClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XAuthClientInfo.Marshal(b, m, deterministic)
}
func (m *Dot1XAuthClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XAuthClientInfo.Merge(m, src)
}
func (m *Dot1XAuthClientInfo) XXX_Size() int {
	return xxx_messageInfo_Dot1XAuthClientInfo.Size(m)
}
func (m *Dot1XAuthClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XAuthClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XAuthClientInfo proto.InternalMessageInfo

func (m *Dot1XAuthClientInfo) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Dot1XAuthClientInfo) GetAuthSmState() string {
	if m != nil {
		return m.AuthSmState
	}
	return ""
}

func (m *Dot1XAuthClientInfo) GetAuthBendSmState() string {
	if m != nil {
		return m.AuthBendSmState
	}
	return ""
}

func (m *Dot1XAuthClientInfo) GetTimeToNextReauth() string {
	if m != nil {
		return m.TimeToNextReauth
	}
	return ""
}

func (m *Dot1XAuthClientInfo) GetLastAuthTime() string {
	if m != nil {
		return m.LastAuthTime
	}
	return ""
}

type Dot1XAuthInfo struct {
	Reauth               string                 `protobuf:"bytes,1,opt,name=reauth,proto3" json:"reauth,omitempty"`
	ConfigDependency     string                 `protobuf:"bytes,2,opt,name=config_dependency,json=configDependency,proto3" json:"config_dependency,omitempty"`
	Client               []*Dot1XAuthClientInfo `protobuf:"bytes,3,rep,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Dot1XAuthInfo) Reset()         { *m = Dot1XAuthInfo{} }
func (m *Dot1XAuthInfo) String() string { return proto.CompactTextString(m) }
func (*Dot1XAuthInfo) ProtoMessage()    {}
func (*Dot1XAuthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f160d63776e65b5, []int{2}
}

func (m *Dot1XAuthInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XAuthInfo.Unmarshal(m, b)
}
func (m *Dot1XAuthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XAuthInfo.Marshal(b, m, deterministic)
}
func (m *Dot1XAuthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XAuthInfo.Merge(m, src)
}
func (m *Dot1XAuthInfo) XXX_Size() int {
	return xxx_messageInfo_Dot1XAuthInfo.Size(m)
}
func (m *Dot1XAuthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XAuthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XAuthInfo proto.InternalMessageInfo

func (m *Dot1XAuthInfo) GetReauth() string {
	if m != nil {
		return m.Reauth
	}
	return ""
}

func (m *Dot1XAuthInfo) GetConfigDependency() string {
	if m != nil {
		return m.ConfigDependency
	}
	return ""
}

func (m *Dot1XAuthInfo) GetClient() []*Dot1XAuthClientInfo {
	if m != nil {
		return m.Client
	}
	return nil
}

type Dot1XSuppClientInfo struct {
	Mac                  string   `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	EapMethod            string   `protobuf:"bytes,2,opt,name=eap_method,json=eapMethod,proto3" json:"eap_method,omitempty"`
	LastAuthTime         string   `protobuf:"bytes,3,opt,name=last_auth_time,json=lastAuthTime,proto3" json:"last_auth_time,omitempty"`
	AuthSmState          string   `protobuf:"bytes,4,opt,name=auth_sm_state,json=authSmState,proto3" json:"auth_sm_state,omitempty"`
	AuthBendSmState      string   `protobuf:"bytes,5,opt,name=auth_bend_sm_state,json=authBendSmState,proto3" json:"auth_bend_sm_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XSuppClientInfo) Reset()         { *m = Dot1XSuppClientInfo{} }
func (m *Dot1XSuppClientInfo) String() string { return proto.CompactTextString(m) }
func (*Dot1XSuppClientInfo) ProtoMessage()    {}
func (*Dot1XSuppClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f160d63776e65b5, []int{3}
}

func (m *Dot1XSuppClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XSuppClientInfo.Unmarshal(m, b)
}
func (m *Dot1XSuppClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XSuppClientInfo.Marshal(b, m, deterministic)
}
func (m *Dot1XSuppClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XSuppClientInfo.Merge(m, src)
}
func (m *Dot1XSuppClientInfo) XXX_Size() int {
	return xxx_messageInfo_Dot1XSuppClientInfo.Size(m)
}
func (m *Dot1XSuppClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XSuppClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XSuppClientInfo proto.InternalMessageInfo

func (m *Dot1XSuppClientInfo) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Dot1XSuppClientInfo) GetEapMethod() string {
	if m != nil {
		return m.EapMethod
	}
	return ""
}

func (m *Dot1XSuppClientInfo) GetLastAuthTime() string {
	if m != nil {
		return m.LastAuthTime
	}
	return ""
}

func (m *Dot1XSuppClientInfo) GetAuthSmState() string {
	if m != nil {
		return m.AuthSmState
	}
	return ""
}

func (m *Dot1XSuppClientInfo) GetAuthBendSmState() string {
	if m != nil {
		return m.AuthBendSmState
	}
	return ""
}

type Dot1XSuppInfo struct {
	EapProfile           string                 `protobuf:"bytes,1,opt,name=eap_profile,json=eapProfile,proto3" json:"eap_profile,omitempty"`
	ConfigDependency     string                 `protobuf:"bytes,2,opt,name=config_dependency,json=configDependency,proto3" json:"config_dependency,omitempty"`
	Client               []*Dot1XSuppClientInfo `protobuf:"bytes,3,rep,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Dot1XSuppInfo) Reset()         { *m = Dot1XSuppInfo{} }
func (m *Dot1XSuppInfo) String() string { return proto.CompactTextString(m) }
func (*Dot1XSuppInfo) ProtoMessage()    {}
func (*Dot1XSuppInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f160d63776e65b5, []int{4}
}

func (m *Dot1XSuppInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XSuppInfo.Unmarshal(m, b)
}
func (m *Dot1XSuppInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XSuppInfo.Marshal(b, m, deterministic)
}
func (m *Dot1XSuppInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XSuppInfo.Merge(m, src)
}
func (m *Dot1XSuppInfo) XXX_Size() int {
	return xxx_messageInfo_Dot1XSuppInfo.Size(m)
}
func (m *Dot1XSuppInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XSuppInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XSuppInfo proto.InternalMessageInfo

func (m *Dot1XSuppInfo) GetEapProfile() string {
	if m != nil {
		return m.EapProfile
	}
	return ""
}

func (m *Dot1XSuppInfo) GetConfigDependency() string {
	if m != nil {
		return m.ConfigDependency
	}
	return ""
}

func (m *Dot1XSuppInfo) GetClient() []*Dot1XSuppClientInfo {
	if m != nil {
		return m.Client
	}
	return nil
}

type Dot1XInterfaceInfo struct {
	Pae                  string         `protobuf:"bytes,1,opt,name=pae,proto3" json:"pae,omitempty"`
	PortStatus           string         `protobuf:"bytes,2,opt,name=port_status,json=portStatus,proto3" json:"port_status,omitempty"`
	Dot1XProfile         string         `protobuf:"bytes,3,opt,name=dot1x_profile,json=dot1xProfile,proto3" json:"dot1x_profile,omitempty"`
	AuthInfo             *Dot1XAuthInfo `protobuf:"bytes,4,opt,name=auth_info,json=authInfo,proto3" json:"auth_info,omitempty"`
	SuppInfo             *Dot1XSuppInfo `protobuf:"bytes,5,opt,name=supp_info,json=suppInfo,proto3" json:"supp_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Dot1XInterfaceInfo) Reset()         { *m = Dot1XInterfaceInfo{} }
func (m *Dot1XInterfaceInfo) String() string { return proto.CompactTextString(m) }
func (*Dot1XInterfaceInfo) ProtoMessage()    {}
func (*Dot1XInterfaceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f160d63776e65b5, []int{5}
}

func (m *Dot1XInterfaceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XInterfaceInfo.Unmarshal(m, b)
}
func (m *Dot1XInterfaceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XInterfaceInfo.Marshal(b, m, deterministic)
}
func (m *Dot1XInterfaceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XInterfaceInfo.Merge(m, src)
}
func (m *Dot1XInterfaceInfo) XXX_Size() int {
	return xxx_messageInfo_Dot1XInterfaceInfo.Size(m)
}
func (m *Dot1XInterfaceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XInterfaceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XInterfaceInfo proto.InternalMessageInfo

func (m *Dot1XInterfaceInfo) GetPae() string {
	if m != nil {
		return m.Pae
	}
	return ""
}

func (m *Dot1XInterfaceInfo) GetPortStatus() string {
	if m != nil {
		return m.PortStatus
	}
	return ""
}

func (m *Dot1XInterfaceInfo) GetDot1XProfile() string {
	if m != nil {
		return m.Dot1XProfile
	}
	return ""
}

func (m *Dot1XInterfaceInfo) GetAuthInfo() *Dot1XAuthInfo {
	if m != nil {
		return m.AuthInfo
	}
	return nil
}

func (m *Dot1XInterfaceInfo) GetSuppInfo() *Dot1XSuppInfo {
	if m != nil {
		return m.SuppInfo
	}
	return nil
}

type Dot1XInterfaceMka struct {
	TieBreakRole         string   `protobuf:"bytes,1,opt,name=tie_break_role,json=tieBreakRole,proto3" json:"tie_break_role,omitempty"`
	EapBasedMacsec       string   `protobuf:"bytes,2,opt,name=eap_based_macsec,json=eapBasedMacsec,proto3" json:"eap_based_macsec,omitempty"`
	MkaStartTime         string   `protobuf:"bytes,3,opt,name=mka_start_time,json=mkaStartTime,proto3" json:"mka_start_time,omitempty"`
	MkaStopTime          string   `protobuf:"bytes,4,opt,name=mka_stop_time,json=mkaStopTime,proto3" json:"mka_stop_time,omitempty"`
	MkaResponseTime      string   `protobuf:"bytes,5,opt,name=mka_response_time,json=mkaResponseTime,proto3" json:"mka_response_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XInterfaceMka) Reset()         { *m = Dot1XInterfaceMka{} }
func (m *Dot1XInterfaceMka) String() string { return proto.CompactTextString(m) }
func (*Dot1XInterfaceMka) ProtoMessage()    {}
func (*Dot1XInterfaceMka) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f160d63776e65b5, []int{6}
}

func (m *Dot1XInterfaceMka) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XInterfaceMka.Unmarshal(m, b)
}
func (m *Dot1XInterfaceMka) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XInterfaceMka.Marshal(b, m, deterministic)
}
func (m *Dot1XInterfaceMka) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XInterfaceMka.Merge(m, src)
}
func (m *Dot1XInterfaceMka) XXX_Size() int {
	return xxx_messageInfo_Dot1XInterfaceMka.Size(m)
}
func (m *Dot1XInterfaceMka) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XInterfaceMka.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XInterfaceMka proto.InternalMessageInfo

func (m *Dot1XInterfaceMka) GetTieBreakRole() string {
	if m != nil {
		return m.TieBreakRole
	}
	return ""
}

func (m *Dot1XInterfaceMka) GetEapBasedMacsec() string {
	if m != nil {
		return m.EapBasedMacsec
	}
	return ""
}

func (m *Dot1XInterfaceMka) GetMkaStartTime() string {
	if m != nil {
		return m.MkaStartTime
	}
	return ""
}

func (m *Dot1XInterfaceMka) GetMkaStopTime() string {
	if m != nil {
		return m.MkaStopTime
	}
	return ""
}

func (m *Dot1XInterfaceMka) GetMkaResponseTime() string {
	if m != nil {
		return m.MkaResponseTime
	}
	return ""
}

type Dot1XInterfaceDetail struct {
	InterfaceName        string              `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	InterfaceSname       string              `protobuf:"bytes,51,opt,name=interface_sname,json=interfaceSname,proto3" json:"interface_sname,omitempty"`
	IfHandle             string              `protobuf:"bytes,52,opt,name=if_handle,json=ifHandle,proto3" json:"if_handle,omitempty"`
	Mac                  string              `protobuf:"bytes,53,opt,name=mac,proto3" json:"mac,omitempty"`
	Ethertype            string              `protobuf:"bytes,54,opt,name=ethertype,proto3" json:"ethertype,omitempty"`
	IntfInfo             *Dot1XInterfaceInfo `protobuf:"bytes,55,opt,name=intf_info,json=intfInfo,proto3" json:"intf_info,omitempty"`
	MkaStatusInfo        *Dot1XInterfaceMka  `protobuf:"bytes,56,opt,name=mka_status_info,json=mkaStatusInfo,proto3" json:"mka_status_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Dot1XInterfaceDetail) Reset()         { *m = Dot1XInterfaceDetail{} }
func (m *Dot1XInterfaceDetail) String() string { return proto.CompactTextString(m) }
func (*Dot1XInterfaceDetail) ProtoMessage()    {}
func (*Dot1XInterfaceDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f160d63776e65b5, []int{7}
}

func (m *Dot1XInterfaceDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XInterfaceDetail.Unmarshal(m, b)
}
func (m *Dot1XInterfaceDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XInterfaceDetail.Marshal(b, m, deterministic)
}
func (m *Dot1XInterfaceDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XInterfaceDetail.Merge(m, src)
}
func (m *Dot1XInterfaceDetail) XXX_Size() int {
	return xxx_messageInfo_Dot1XInterfaceDetail.Size(m)
}
func (m *Dot1XInterfaceDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XInterfaceDetail.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XInterfaceDetail proto.InternalMessageInfo

func (m *Dot1XInterfaceDetail) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Dot1XInterfaceDetail) GetInterfaceSname() string {
	if m != nil {
		return m.InterfaceSname
	}
	return ""
}

func (m *Dot1XInterfaceDetail) GetIfHandle() string {
	if m != nil {
		return m.IfHandle
	}
	return ""
}

func (m *Dot1XInterfaceDetail) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Dot1XInterfaceDetail) GetEthertype() string {
	if m != nil {
		return m.Ethertype
	}
	return ""
}

func (m *Dot1XInterfaceDetail) GetIntfInfo() *Dot1XInterfaceInfo {
	if m != nil {
		return m.IntfInfo
	}
	return nil
}

func (m *Dot1XInterfaceDetail) GetMkaStatusInfo() *Dot1XInterfaceMka {
	if m != nil {
		return m.MkaStatusInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Dot1XInterfaceDetail_KEYS)(nil), "cisco_ios_xr_dot1x_oper.dot1x.session.interface_sessions.interface_session.dot1x_interface_detail_KEYS")
	proto.RegisterType((*Dot1XAuthClientInfo)(nil), "cisco_ios_xr_dot1x_oper.dot1x.session.interface_sessions.interface_session.dot1x_auth_client_info")
	proto.RegisterType((*Dot1XAuthInfo)(nil), "cisco_ios_xr_dot1x_oper.dot1x.session.interface_sessions.interface_session.dot1x_auth_info")
	proto.RegisterType((*Dot1XSuppClientInfo)(nil), "cisco_ios_xr_dot1x_oper.dot1x.session.interface_sessions.interface_session.dot1x_supp_client_info")
	proto.RegisterType((*Dot1XSuppInfo)(nil), "cisco_ios_xr_dot1x_oper.dot1x.session.interface_sessions.interface_session.dot1x_supp_info")
	proto.RegisterType((*Dot1XInterfaceInfo)(nil), "cisco_ios_xr_dot1x_oper.dot1x.session.interface_sessions.interface_session.dot1x_interface_info")
	proto.RegisterType((*Dot1XInterfaceMka)(nil), "cisco_ios_xr_dot1x_oper.dot1x.session.interface_sessions.interface_session.dot1x_interface_mka")
	proto.RegisterType((*Dot1XInterfaceDetail)(nil), "cisco_ios_xr_dot1x_oper.dot1x.session.interface_sessions.interface_session.dot1x_interface_detail")
}

func init() { proto.RegisterFile("dot1x_interface_detail.proto", fileDescriptor_2f160d63776e65b5) }

var fileDescriptor_2f160d63776e65b5 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x4f, 0x6b, 0x53, 0x4f,
	0x14, 0x25, 0x4d, 0x5b, 0x92, 0x9b, 0xa6, 0x49, 0xa7, 0x3f, 0x4a, 0xa0, 0xfd, 0x61, 0x89, 0x8a,
	0xc5, 0x62, 0xa0, 0xad, 0xff, 0xb6, 0x16, 0x05, 0xff, 0xd0, 0x22, 0x49, 0x37, 0xe2, 0x62, 0x9c,
	0xbc, 0x77, 0x63, 0x86, 0xe4, 0xbd, 0x19, 0xde, 0x4c, 0x20, 0x15, 0x5c, 0xfb, 0xad, 0xc4, 0x8f,
	0x20, 0x6e, 0xdd, 0xf9, 0x49, 0xe4, 0xde, 0xf7, 0x27, 0x6d, 0x1a, 0xd0, 0x45, 0x0b, 0xee, 0x26,
	0x67, 0xce, 0x9b, 0xb9, 0xe7, 0x9e, 0x33, 0x37, 0xb0, 0x13, 0x1a, 0x7f, 0x30, 0x95, 0x3a, 0xf6,
	0x98, 0x0c, 0x54, 0x80, 0x32, 0x44, 0xaf, 0xf4, 0xb8, 0x63, 0x13, 0xe3, 0x8d, 0x78, 0x1d, 0x68,
	0x17, 0x18, 0xa9, 0x8d, 0x93, 0xd3, 0x44, 0xa6, 0x54, 0x63, 0x31, 0xe9, 0xf0, 0xb2, 0xe3, 0xd0,
	0x39, 0x6d, 0xe2, 0xce, 0xec, 0xeb, 0x0c, 0x71, 0x57, 0xa1, 0xf6, 0x01, 0x6c, 0x2f, 0xbe, 0x4b,
	0xbe, 0x79, 0xf1, 0xae, 0x27, 0x04, 0x2c, 0xc7, 0x2a, 0xc2, 0x56, 0x69, 0xb7, 0xb4, 0x57, 0xed,
	0xf2, 0xba, 0xfd, 0xbd, 0x04, 0x5b, 0xe9, 0x37, 0x6a, 0xe2, 0x87, 0x32, 0x18, 0x6b, 0x8c, 0xbd,
	0xd4, 0xf1, 0xc0, 0x88, 0x26, 0x94, 0x23, 0x15, 0x64, 0x6c, 0x5a, 0x8a, 0x36, 0xd4, 0x99, 0xe5,
	0x22, 0xe9, 0xbc, 0xf2, 0xd8, 0x5a, 0xe2, 0xbd, 0x1a, 0x81, 0xbd, 0xa8, 0x47, 0x90, 0xd8, 0x07,
	0xc1, 0x9c, 0x3e, 0xc6, 0xe1, 0x8c, 0x58, 0x66, 0x62, 0x83, 0x76, 0x8e, 0x31, 0x0e, 0x73, 0xf2,
	0x03, 0xd8, 0xf4, 0x3a, 0x42, 0xe9, 0x8d, 0x8c, 0x71, 0xea, 0x65, 0x82, 0xc4, 0x68, 0x2d, 0x33,
	0xbb, 0x49, 0x5b, 0x67, 0xe6, 0x14, 0xa7, 0xbe, 0xcb, 0xb8, 0xb8, 0x03, 0xeb, 0x63, 0xe5, 0x7c,
	0x5a, 0x2a, 0xed, 0xb6, 0x56, 0x98, 0xb9, 0x46, 0xe8, 0xb3, 0x89, 0x1f, 0x9e, 0xe9, 0x08, 0xdb,
	0x3f, 0x4a, 0xd0, 0xb8, 0x20, 0x89, 0xb5, 0x6c, 0xc1, 0x6a, 0x76, 0x76, 0x2a, 0x27, 0xfb, 0x25,
	0xf6, 0x61, 0x23, 0x30, 0xf1, 0x40, 0x7f, 0x94, 0x21, 0x5a, 0x8c, 0x43, 0x8c, 0x83, 0xf3, 0x4c,
	0x55, 0x33, 0xdd, 0x78, 0x5e, 0xe0, 0xe2, 0x13, 0xac, 0xa6, 0xfd, 0x69, 0x95, 0x77, 0xcb, 0x7b,
	0xb5, 0xc3, 0x7e, 0xe7, 0xfa, 0xbc, 0xeb, 0x2c, 0x36, 0xa1, 0x9b, 0xdd, 0xd8, 0xfe, 0x56, 0xf8,
	0xe4, 0x26, 0xd6, 0xfe, 0xc1, 0xa7, 0xff, 0x01, 0x50, 0x59, 0x19, 0xa1, 0x1f, 0x9a, 0x30, 0x93,
	0x53, 0x45, 0x65, 0x4f, 0x18, 0x58, 0xd0, 0xc6, 0xf2, 0xd5, 0x36, 0x5e, 0x35, 0x7b, 0xf9, 0x6f,
	0xcd, 0x5e, 0x59, 0x68, 0x76, 0xfb, 0x67, 0xe1, 0x0b, 0x4b, 0xe0, 0xda, 0x6f, 0x41, 0x8d, 0x2a,
	0xb5, 0x89, 0x19, 0xe8, 0x71, 0x9e, 0x4c, 0x2a, 0xfe, 0x6d, 0x8a, 0xfc, 0x73, 0x06, 0xcd, 0x77,
	0xbf, 0x30, 0xe8, 0xd7, 0x12, 0xfc, 0x37, 0xff, 0xf8, 0x72, 0x7b, 0xac, 0xca, 0xa5, 0xd1, 0x92,
	0x44, 0x5b, 0x93, 0x78, 0xee, 0xd6, 0xc4, 0x65, 0x6a, 0x80, 0xa0, 0x1e, 0x23, 0xe2, 0x36, 0xd4,
	0xd3, 0xa3, 0xf2, 0xbe, 0x64, 0xfe, 0x30, 0x98, 0x77, 0x66, 0x0a, 0xd5, 0x22, 0xdf, 0xec, 0x4d,
	0xed, 0xf0, 0xfd, 0x0d, 0x05, 0x92, 0x85, 0x56, 0x68, 0xf9, 0x8a, 0x14, 0x4d, 0xa1, 0x5a, 0x38,
	0xc8, 0x66, 0xdf, 0xc8, 0xcd, 0xc5, 0x15, 0xdd, 0x0a, 0x2d, 0xe9, 0x66, 0x8a, 0xd0, 0xe6, 0x7c,
	0x93, 0xa3, 0x91, 0xa2, 0x44, 0x7b, 0x8d, 0xb2, 0x9f, 0xa0, 0x1a, 0xc9, 0xc4, 0x14, 0x49, 0x5a,
	0xf3, 0x1a, 0x8f, 0x09, 0xec, 0x9a, 0x31, 0x8a, 0x3d, 0x68, 0x52, 0xd8, 0xfa, 0xca, 0x61, 0x28,
	0x23, 0x15, 0x38, 0x0c, 0xb2, 0xe6, 0xaf, 0xa3, 0xb2, 0xc7, 0x04, 0x9f, 0x30, 0x4a, 0xe7, 0x45,
	0x23, 0x45, 0x06, 0x25, 0xfe, 0xd2, 0x0b, 0x89, 0x46, 0xaa, 0x47, 0x60, 0xfe, 0x42, 0x52, 0x96,
	0xb1, 0x29, 0x29, 0x7b, 0x21, 0x4c, 0x32, 0x96, 0x39, 0xf7, 0x61, 0x83, 0x38, 0x09, 0x3a, 0x6b,
	0x62, 0x87, 0x17, 0xa7, 0x56, 0x23, 0x1a, 0xa9, 0x6e, 0x86, 0xf3, 0xe0, 0xfa, 0x5a, 0xce, 0xdf,
	0xf8, 0xfc, 0xfc, 0x16, 0x77, 0x61, 0x7d, 0x86, 0xf1, 0x10, 0x3f, 0xe4, 0x33, 0xea, 0x05, 0x7a,
	0xaa, 0x22, 0x14, 0xf7, 0xa0, 0x71, 0xa1, 0x9d, 0xcc, 0x3b, 0x4a, 0x05, 0x16, 0x70, 0x8f, 0x50,
	0xb1, 0x0d, 0x55, 0x3d, 0x90, 0x43, 0x15, 0x87, 0x63, 0x6c, 0x3d, 0x64, 0x4a, 0x45, 0x0f, 0x5e,
	0xf2, 0xef, 0x7c, 0xa0, 0x3c, 0x9a, 0x0d, 0x94, 0x1d, 0xa8, 0xa2, 0x1f, 0x62, 0xe2, 0xcf, 0x2d,
	0xb6, 0x1e, 0x67, 0xf3, 0x24, 0x07, 0xc4, 0x67, 0xa8, 0xea, 0xd8, 0x0f, 0xd2, 0x3c, 0x3c, 0xe1,
	0x3c, 0x7c, 0xb8, 0xfe, 0x3c, 0x5c, 0x7e, 0x56, 0xdd, 0x0a, 0x5d, 0xc9, 0x71, 0xfc, 0x52, 0x82,
	0x46, 0xe6, 0x96, 0x9f, 0xb8, 0xb4, 0x8a, 0xa7, 0x5c, 0x85, 0xbc, 0xc9, 0x2a, 0xc8, 0xbd, 0x7a,
	0x9a, 0x07, 0x3f, 0x71, 0x54, 0x49, 0x7f, 0x95, 0xff, 0xd2, 0x8f, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x8c, 0x14, 0xf9, 0x54, 0xf2, 0x07, 0x00, 0x00,
}