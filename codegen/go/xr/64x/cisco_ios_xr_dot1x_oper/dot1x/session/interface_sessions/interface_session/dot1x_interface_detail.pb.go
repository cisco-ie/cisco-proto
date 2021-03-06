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
	LastAuthServer       string   `protobuf:"bytes,6,opt,name=last_auth_server,json=lastAuthServer,proto3" json:"last_auth_server,omitempty"`
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

func (m *Dot1XAuthClientInfo) GetLastAuthServer() string {
	if m != nil {
		return m.LastAuthServer
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
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcb, 0x6a, 0x1b, 0x4b,
	0x10, 0x45, 0x96, 0x2d, 0xa4, 0x92, 0xf5, 0x70, 0xfb, 0x62, 0x06, 0xec, 0xcb, 0x35, 0xba, 0x09,
	0x31, 0x31, 0x11, 0xd8, 0xce, 0x6b, 0x1b, 0x93, 0x40, 0x1e, 0xd8, 0x04, 0xc9, 0x9b, 0x90, 0x45,
	0xa7, 0x35, 0x53, 0x8a, 0x1a, 0x69, 0xa6, 0x87, 0xe9, 0x56, 0x90, 0x03, 0x59, 0xe7, 0xaf, 0x42,
	0xbe, 0x21, 0xdb, 0xec, 0xf2, 0x19, 0x59, 0x85, 0xaa, 0x79, 0xc8, 0x96, 0x05, 0xc9, 0xc2, 0x86,
	0xec, 0x5a, 0xa7, 0xcf, 0x74, 0xd7, 0xa9, 0x73, 0xba, 0x04, 0x3b, 0x81, 0x71, 0x07, 0x33, 0xa9,
	0x23, 0x87, 0xc9, 0x50, 0xf9, 0x28, 0x03, 0x74, 0x4a, 0x4f, 0xba, 0x71, 0x62, 0x9c, 0x11, 0x2f,
	0x7d, 0x6d, 0x7d, 0x23, 0xb5, 0xb1, 0x72, 0x96, 0xc8, 0x94, 0x6a, 0x62, 0x4c, 0xba, 0xbc, 0xec,
	0x5a, 0xb4, 0x56, 0x9b, 0xa8, 0x3b, 0xff, 0x3a, 0x43, 0xec, 0x55, 0xa8, 0x73, 0x00, 0xdb, 0xcb,
	0xef, 0x92, 0xaf, 0x9e, 0xbd, 0xe9, 0x0b, 0x01, 0xab, 0x91, 0x0a, 0xd1, 0x2b, 0xed, 0x96, 0xf6,
	0x6a, 0x3d, 0x5e, 0x77, 0x7e, 0x96, 0x60, 0x2b, 0xfd, 0x46, 0x4d, 0xdd, 0x48, 0xfa, 0x13, 0x8d,
	0x91, 0x93, 0x3a, 0x1a, 0x1a, 0xd1, 0x86, 0x72, 0xa8, 0xfc, 0x8c, 0x4d, 0x4b, 0xd1, 0x81, 0x06,
	0xb3, 0x6c, 0x28, 0xad, 0x53, 0x0e, 0xbd, 0x15, 0xde, 0xab, 0x13, 0xd8, 0x0f, 0xfb, 0x04, 0x89,
	0x7d, 0x10, 0xcc, 0x19, 0x60, 0x14, 0xcc, 0x89, 0x65, 0x26, 0xb6, 0x68, 0xe7, 0x18, 0xa3, 0x20,
	0x27, 0xdf, 0x83, 0x4d, 0xa7, 0x43, 0x94, 0xce, 0xc8, 0x08, 0x67, 0x4e, 0x26, 0x48, 0x0c, 0x6f,
	0x95, 0xd9, 0x6d, 0xda, 0x3a, 0x33, 0xa7, 0x38, 0x73, 0x3d, 0xc6, 0xc5, 0x2d, 0x68, 0x4e, 0x94,
	0x75, 0x69, 0xa9, 0xb4, 0xeb, 0xad, 0x31, 0x73, 0x9d, 0xd0, 0x27, 0x53, 0x37, 0x3a, 0xd3, 0x21,
	0x8a, 0x3d, 0x68, 0xcf, 0x59, 0x16, 0x93, 0x0f, 0x98, 0x78, 0x15, 0xe6, 0x35, 0x73, 0x5e, 0x9f,
	0xd1, 0xce, 0xb7, 0x12, 0xb4, 0x2e, 0x88, 0x67, 0xd5, 0x5b, 0x50, 0xc9, 0xaa, 0x48, 0x85, 0x67,
	0xbf, 0xc4, 0x3e, 0x6c, 0xf8, 0x26, 0x1a, 0xea, 0xf7, 0x32, 0xc0, 0x18, 0xa3, 0x00, 0x23, 0xff,
	0x3c, 0xd3, 0xdf, 0x4e, 0x37, 0x9e, 0x16, 0xb8, 0xf8, 0x08, 0x95, 0xb4, 0x93, 0x5e, 0x79, 0xb7,
	0xbc, 0x57, 0x3f, 0x1c, 0x74, 0xaf, 0xcf, 0xe5, 0xee, 0x72, 0xbb, 0x7a, 0xd9, 0x8d, 0x9d, 0xaf,
	0x85, 0xa3, 0x76, 0x1a, 0xc7, 0xbf, 0x71, 0xf4, 0x5f, 0x00, 0x54, 0xb1, 0x0c, 0xd1, 0x8d, 0x4c,
	0x90, 0xc9, 0xa9, 0xa1, 0x8a, 0x4f, 0x18, 0x58, 0xd2, 0xf0, 0xf2, 0x92, 0x86, 0x5f, 0x89, 0xc5,
	0xea, 0x9f, 0xc6, 0x62, 0x6d, 0x69, 0x2c, 0x3a, 0xdf, 0x0b, 0x5f, 0x58, 0x02, 0xd7, 0xfe, 0x1f,
	0xd4, 0xa9, 0xd2, 0x38, 0x31, 0x43, 0x3d, 0xc9, 0x33, 0x4c, 0xc5, 0xbf, 0x4e, 0x91, 0xbf, 0xce,
	0xa0, 0xc5, 0xee, 0x17, 0x06, 0xfd, 0x58, 0x81, 0x7f, 0x16, 0x9f, 0x69, 0x6e, 0x4f, 0xac, 0x72,
	0x69, 0xb4, 0x24, 0xd1, 0xb1, 0x49, 0x1c, 0x77, 0x6b, 0x6a, 0x33, 0x35, 0x40, 0x50, 0x9f, 0x11,
	0xf1, 0x3f, 0x34, 0xd2, 0xa3, 0xf2, 0xbe, 0x64, 0xfe, 0x30, 0x98, 0x77, 0x66, 0x06, 0xb5, 0x22,
	0xdf, 0xec, 0x4d, 0xfd, 0xf0, 0xed, 0x0d, 0x05, 0x92, 0x85, 0x56, 0x69, 0xf9, 0x82, 0x14, 0xcd,
	0xa0, 0x56, 0x38, 0xc8, 0x66, 0xdf, 0xc8, 0xcd, 0xc5, 0x15, 0xbd, 0x2a, 0x2d, 0xe9, 0x66, 0x8a,
	0xd0, 0xe6, 0x62, 0x93, 0xc3, 0xb1, 0xa2, 0x44, 0x3b, 0x8d, 0x72, 0x90, 0xa0, 0x1a, 0xcb, 0xc4,
	0x14, 0x49, 0x5a, 0x77, 0x1a, 0x8f, 0x09, 0xec, 0x99, 0x09, 0x8f, 0x10, 0x0a, 0xdb, 0x40, 0x59,
	0x0c, 0x64, 0xa8, 0x7c, 0x8b, 0x7e, 0xd6, 0xfc, 0x26, 0xaa, 0xf8, 0x98, 0xe0, 0x13, 0x46, 0xe9,
	0xbc, 0x70, 0xac, 0xc8, 0xa0, 0xc4, 0x5d, 0x7a, 0x21, 0xe1, 0x58, 0xf5, 0x09, 0xcc, 0x5f, 0x48,
	0xca, 0x32, 0x71, 0x4a, 0xca, 0x5e, 0x08, 0x93, 0x4c, 0xcc, 0x9c, 0xbb, 0xb0, 0x41, 0x9c, 0x04,
	0x6d, 0x6c, 0x22, 0x8b, 0x17, 0xe7, 0x5b, 0x2b, 0x1c, 0xab, 0x5e, 0x86, 0x13, 0xb7, 0xf3, 0xa5,
	0x9c, 0xbf, 0xf1, 0xc5, 0x49, 0x2f, 0x6e, 0x43, 0x73, 0x8e, 0xf1, 0xb8, 0x3f, 0xe4, 0x33, 0x1a,
	0x05, 0x7a, 0xaa, 0x42, 0x14, 0x77, 0xa0, 0x75, 0xa1, 0x9d, 0xcc, 0x3b, 0x4a, 0x05, 0x16, 0x70,
	0x9f, 0x50, 0xb1, 0x0d, 0x35, 0x3d, 0x94, 0x23, 0x15, 0x05, 0x13, 0xf4, 0xee, 0x33, 0xa5, 0xaa,
	0x87, 0xcf, 0xf9, 0x77, 0x3e, 0x50, 0x1e, 0xcc, 0x07, 0xca, 0x0e, 0xd4, 0xd0, 0x8d, 0x30, 0x71,
	0xe7, 0x31, 0x7a, 0x0f, 0xb3, 0x79, 0x92, 0x03, 0xe2, 0x13, 0xd4, 0x74, 0xe4, 0x86, 0x69, 0x1e,
	0x1e, 0x71, 0x1e, 0xde, 0x5d, 0x7f, 0x1e, 0x2e, 0x3f, 0xab, 0x5e, 0x95, 0xae, 0xe4, 0x38, 0x7e,
	0x2e, 0x41, 0x2b, 0x73, 0xcb, 0x4d, 0x6d, 0x5a, 0xc5, 0x63, 0xae, 0x42, 0xde, 0x64, 0x15, 0xe4,
	0x5e, 0x23, 0xcd, 0x83, 0x9b, 0x5a, 0xaa, 0x64, 0x50, 0xe1, 0x3f, 0xff, 0xa3, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x82, 0xa4, 0xf4, 0x8f, 0x1c, 0x08, 0x00, 0x00,
}
