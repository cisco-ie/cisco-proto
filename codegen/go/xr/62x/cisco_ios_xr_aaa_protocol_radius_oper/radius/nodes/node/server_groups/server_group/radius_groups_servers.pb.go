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
// source: radius_groups_servers.proto

package cisco_ios_xr_aaa_protocol_radius_oper_radius_nodes_node_server_groups_server_group

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

type RadiusGroupsServers_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ServerGroupName      string   `protobuf:"bytes,2,opt,name=server_group_name,json=serverGroupName,proto3" json:"server_group_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RadiusGroupsServers_KEYS) Reset()         { *m = RadiusGroupsServers_KEYS{} }
func (m *RadiusGroupsServers_KEYS) String() string { return proto.CompactTextString(m) }
func (*RadiusGroupsServers_KEYS) ProtoMessage()    {}
func (*RadiusGroupsServers_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_529eb045ccaae6fc, []int{0}
}

func (m *RadiusGroupsServers_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusGroupsServers_KEYS.Unmarshal(m, b)
}
func (m *RadiusGroupsServers_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusGroupsServers_KEYS.Marshal(b, m, deterministic)
}
func (m *RadiusGroupsServers_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusGroupsServers_KEYS.Merge(m, src)
}
func (m *RadiusGroupsServers_KEYS) XXX_Size() int {
	return xxx_messageInfo_RadiusGroupsServers_KEYS.Size(m)
}
func (m *RadiusGroupsServers_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusGroupsServers_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusGroupsServers_KEYS proto.InternalMessageInfo

func (m *RadiusGroupsServers_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *RadiusGroupsServers_KEYS) GetServerGroupName() string {
	if m != nil {
		return m.ServerGroupName
	}
	return ""
}

type RadiusAccountingData struct {
	Requests                  uint32   `protobuf:"varint,1,opt,name=requests,proto3" json:"requests,omitempty"`
	PendingRequests           uint32   `protobuf:"varint,2,opt,name=pending_requests,json=pendingRequests,proto3" json:"pending_requests,omitempty"`
	Retransmits               uint32   `protobuf:"varint,3,opt,name=retransmits,proto3" json:"retransmits,omitempty"`
	Responses                 uint32   `protobuf:"varint,4,opt,name=responses,proto3" json:"responses,omitempty"`
	Timeouts                  uint32   `protobuf:"varint,5,opt,name=timeouts,proto3" json:"timeouts,omitempty"`
	BadResponses              uint32   `protobuf:"varint,6,opt,name=bad_responses,json=badResponses,proto3" json:"bad_responses,omitempty"`
	BadAuthenticators         uint32   `protobuf:"varint,7,opt,name=bad_authenticators,json=badAuthenticators,proto3" json:"bad_authenticators,omitempty"`
	UnknownPacketTypes        uint32   `protobuf:"varint,8,opt,name=unknown_packet_types,json=unknownPacketTypes,proto3" json:"unknown_packet_types,omitempty"`
	DroppedResponses          uint32   `protobuf:"varint,9,opt,name=dropped_responses,json=droppedResponses,proto3" json:"dropped_responses,omitempty"`
	Rtt                       uint32   `protobuf:"varint,10,opt,name=rtt,proto3" json:"rtt,omitempty"`
	AcctUnexpectedResponses   uint32   `protobuf:"varint,11,opt,name=acct_unexpected_responses,json=acctUnexpectedResponses,proto3" json:"acct_unexpected_responses,omitempty"`
	AcctServerErrorResponses  uint32   `protobuf:"varint,12,opt,name=acct_server_error_responses,json=acctServerErrorResponses,proto3" json:"acct_server_error_responses,omitempty"`
	AcctIncorrectResponses    uint32   `protobuf:"varint,13,opt,name=acct_incorrect_responses,json=acctIncorrectResponses,proto3" json:"acct_incorrect_responses,omitempty"`
	AcctResponseTime          uint32   `protobuf:"varint,14,opt,name=acct_response_time,json=acctResponseTime,proto3" json:"acct_response_time,omitempty"`
	AcctTransactionSuccessess uint32   `protobuf:"varint,15,opt,name=acct_transaction_successess,json=acctTransactionSuccessess,proto3" json:"acct_transaction_successess,omitempty"`
	AcctTransactionFailure    uint32   `protobuf:"varint,16,opt,name=acct_transaction_failure,json=acctTransactionFailure,proto3" json:"acct_transaction_failure,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *RadiusAccountingData) Reset()         { *m = RadiusAccountingData{} }
func (m *RadiusAccountingData) String() string { return proto.CompactTextString(m) }
func (*RadiusAccountingData) ProtoMessage()    {}
func (*RadiusAccountingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_529eb045ccaae6fc, []int{1}
}

func (m *RadiusAccountingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusAccountingData.Unmarshal(m, b)
}
func (m *RadiusAccountingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusAccountingData.Marshal(b, m, deterministic)
}
func (m *RadiusAccountingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusAccountingData.Merge(m, src)
}
func (m *RadiusAccountingData) XXX_Size() int {
	return xxx_messageInfo_RadiusAccountingData.Size(m)
}
func (m *RadiusAccountingData) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusAccountingData.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusAccountingData proto.InternalMessageInfo

func (m *RadiusAccountingData) GetRequests() uint32 {
	if m != nil {
		return m.Requests
	}
	return 0
}

func (m *RadiusAccountingData) GetPendingRequests() uint32 {
	if m != nil {
		return m.PendingRequests
	}
	return 0
}

func (m *RadiusAccountingData) GetRetransmits() uint32 {
	if m != nil {
		return m.Retransmits
	}
	return 0
}

func (m *RadiusAccountingData) GetResponses() uint32 {
	if m != nil {
		return m.Responses
	}
	return 0
}

func (m *RadiusAccountingData) GetTimeouts() uint32 {
	if m != nil {
		return m.Timeouts
	}
	return 0
}

func (m *RadiusAccountingData) GetBadResponses() uint32 {
	if m != nil {
		return m.BadResponses
	}
	return 0
}

func (m *RadiusAccountingData) GetBadAuthenticators() uint32 {
	if m != nil {
		return m.BadAuthenticators
	}
	return 0
}

func (m *RadiusAccountingData) GetUnknownPacketTypes() uint32 {
	if m != nil {
		return m.UnknownPacketTypes
	}
	return 0
}

func (m *RadiusAccountingData) GetDroppedResponses() uint32 {
	if m != nil {
		return m.DroppedResponses
	}
	return 0
}

func (m *RadiusAccountingData) GetRtt() uint32 {
	if m != nil {
		return m.Rtt
	}
	return 0
}

func (m *RadiusAccountingData) GetAcctUnexpectedResponses() uint32 {
	if m != nil {
		return m.AcctUnexpectedResponses
	}
	return 0
}

func (m *RadiusAccountingData) GetAcctServerErrorResponses() uint32 {
	if m != nil {
		return m.AcctServerErrorResponses
	}
	return 0
}

func (m *RadiusAccountingData) GetAcctIncorrectResponses() uint32 {
	if m != nil {
		return m.AcctIncorrectResponses
	}
	return 0
}

func (m *RadiusAccountingData) GetAcctResponseTime() uint32 {
	if m != nil {
		return m.AcctResponseTime
	}
	return 0
}

func (m *RadiusAccountingData) GetAcctTransactionSuccessess() uint32 {
	if m != nil {
		return m.AcctTransactionSuccessess
	}
	return 0
}

func (m *RadiusAccountingData) GetAcctTransactionFailure() uint32 {
	if m != nil {
		return m.AcctTransactionFailure
	}
	return 0
}

type RadiusAuthenticationData struct {
	AccessRequests              uint32   `protobuf:"varint,1,opt,name=access_requests,json=accessRequests,proto3" json:"access_requests,omitempty"`
	PendingAccessRequests       uint32   `protobuf:"varint,2,opt,name=pending_access_requests,json=pendingAccessRequests,proto3" json:"pending_access_requests,omitempty"`
	AccessRequestRetransmits    uint32   `protobuf:"varint,3,opt,name=access_request_retransmits,json=accessRequestRetransmits,proto3" json:"access_request_retransmits,omitempty"`
	AccessAccepts               uint32   `protobuf:"varint,4,opt,name=access_accepts,json=accessAccepts,proto3" json:"access_accepts,omitempty"`
	AccessRejects               uint32   `protobuf:"varint,5,opt,name=access_rejects,json=accessRejects,proto3" json:"access_rejects,omitempty"`
	AccessChallenges            uint32   `protobuf:"varint,6,opt,name=access_challenges,json=accessChallenges,proto3" json:"access_challenges,omitempty"`
	AccessTimeouts              uint32   `protobuf:"varint,7,opt,name=access_timeouts,json=accessTimeouts,proto3" json:"access_timeouts,omitempty"`
	BadAccessResponses          uint32   `protobuf:"varint,8,opt,name=bad_access_responses,json=badAccessResponses,proto3" json:"bad_access_responses,omitempty"`
	BadAccessAuthenticators     uint32   `protobuf:"varint,9,opt,name=bad_access_authenticators,json=badAccessAuthenticators,proto3" json:"bad_access_authenticators,omitempty"`
	UnknownAccessTypes          uint32   `protobuf:"varint,10,opt,name=unknown_access_types,json=unknownAccessTypes,proto3" json:"unknown_access_types,omitempty"`
	DroppedAccessResponses      uint32   `protobuf:"varint,11,opt,name=dropped_access_responses,json=droppedAccessResponses,proto3" json:"dropped_access_responses,omitempty"`
	Rtt                         uint32   `protobuf:"varint,12,opt,name=rtt,proto3" json:"rtt,omitempty"`
	AuthenResponseTime          uint32   `protobuf:"varint,13,opt,name=authen_response_time,json=authenResponseTime,proto3" json:"authen_response_time,omitempty"`
	AuthenTransactionSuccessess uint32   `protobuf:"varint,14,opt,name=authen_transaction_successess,json=authenTransactionSuccessess,proto3" json:"authen_transaction_successess,omitempty"`
	AuthenTransactionFailure    uint32   `protobuf:"varint,15,opt,name=authen_transaction_failure,json=authenTransactionFailure,proto3" json:"authen_transaction_failure,omitempty"`
	AuthenUnexpectedResponses   uint32   `protobuf:"varint,16,opt,name=authen_unexpected_responses,json=authenUnexpectedResponses,proto3" json:"authen_unexpected_responses,omitempty"`
	AuthenServerErrorResponses  uint32   `protobuf:"varint,17,opt,name=authen_server_error_responses,json=authenServerErrorResponses,proto3" json:"authen_server_error_responses,omitempty"`
	AuthenIncorrectResponses    uint32   `protobuf:"varint,18,opt,name=authen_incorrect_responses,json=authenIncorrectResponses,proto3" json:"authen_incorrect_responses,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *RadiusAuthenticationData) Reset()         { *m = RadiusAuthenticationData{} }
func (m *RadiusAuthenticationData) String() string { return proto.CompactTextString(m) }
func (*RadiusAuthenticationData) ProtoMessage()    {}
func (*RadiusAuthenticationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_529eb045ccaae6fc, []int{2}
}

func (m *RadiusAuthenticationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusAuthenticationData.Unmarshal(m, b)
}
func (m *RadiusAuthenticationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusAuthenticationData.Marshal(b, m, deterministic)
}
func (m *RadiusAuthenticationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusAuthenticationData.Merge(m, src)
}
func (m *RadiusAuthenticationData) XXX_Size() int {
	return xxx_messageInfo_RadiusAuthenticationData.Size(m)
}
func (m *RadiusAuthenticationData) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusAuthenticationData.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusAuthenticationData proto.InternalMessageInfo

func (m *RadiusAuthenticationData) GetAccessRequests() uint32 {
	if m != nil {
		return m.AccessRequests
	}
	return 0
}

func (m *RadiusAuthenticationData) GetPendingAccessRequests() uint32 {
	if m != nil {
		return m.PendingAccessRequests
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAccessRequestRetransmits() uint32 {
	if m != nil {
		return m.AccessRequestRetransmits
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAccessAccepts() uint32 {
	if m != nil {
		return m.AccessAccepts
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAccessRejects() uint32 {
	if m != nil {
		return m.AccessRejects
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAccessChallenges() uint32 {
	if m != nil {
		return m.AccessChallenges
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAccessTimeouts() uint32 {
	if m != nil {
		return m.AccessTimeouts
	}
	return 0
}

func (m *RadiusAuthenticationData) GetBadAccessResponses() uint32 {
	if m != nil {
		return m.BadAccessResponses
	}
	return 0
}

func (m *RadiusAuthenticationData) GetBadAccessAuthenticators() uint32 {
	if m != nil {
		return m.BadAccessAuthenticators
	}
	return 0
}

func (m *RadiusAuthenticationData) GetUnknownAccessTypes() uint32 {
	if m != nil {
		return m.UnknownAccessTypes
	}
	return 0
}

func (m *RadiusAuthenticationData) GetDroppedAccessResponses() uint32 {
	if m != nil {
		return m.DroppedAccessResponses
	}
	return 0
}

func (m *RadiusAuthenticationData) GetRtt() uint32 {
	if m != nil {
		return m.Rtt
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAuthenResponseTime() uint32 {
	if m != nil {
		return m.AuthenResponseTime
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAuthenTransactionSuccessess() uint32 {
	if m != nil {
		return m.AuthenTransactionSuccessess
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAuthenTransactionFailure() uint32 {
	if m != nil {
		return m.AuthenTransactionFailure
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAuthenUnexpectedResponses() uint32 {
	if m != nil {
		return m.AuthenUnexpectedResponses
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAuthenServerErrorResponses() uint32 {
	if m != nil {
		return m.AuthenServerErrorResponses
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAuthenIncorrectResponses() uint32 {
	if m != nil {
		return m.AuthenIncorrectResponses
	}
	return 0
}

type RadiusAuthorizationData struct {
	AuthorRequests              uint32   `protobuf:"varint,1,opt,name=author_requests,json=authorRequests,proto3" json:"author_requests,omitempty"`
	AuthorRequestTimeouts       uint32   `protobuf:"varint,2,opt,name=author_request_timeouts,json=authorRequestTimeouts,proto3" json:"author_request_timeouts,omitempty"`
	AuthorUnexpectedResponses   uint32   `protobuf:"varint,3,opt,name=author_unexpected_responses,json=authorUnexpectedResponses,proto3" json:"author_unexpected_responses,omitempty"`
	AuthorServerErrorResponses  uint32   `protobuf:"varint,4,opt,name=author_server_error_responses,json=authorServerErrorResponses,proto3" json:"author_server_error_responses,omitempty"`
	AuthorIncorrectResponses    uint32   `protobuf:"varint,5,opt,name=author_incorrect_responses,json=authorIncorrectResponses,proto3" json:"author_incorrect_responses,omitempty"`
	AuthorResponseTime          uint32   `protobuf:"varint,6,opt,name=author_response_time,json=authorResponseTime,proto3" json:"author_response_time,omitempty"`
	AuthorTransactionSuccessess uint32   `protobuf:"varint,7,opt,name=author_transaction_successess,json=authorTransactionSuccessess,proto3" json:"author_transaction_successess,omitempty"`
	AuthorTransactionFailure    uint32   `protobuf:"varint,8,opt,name=author_transaction_failure,json=authorTransactionFailure,proto3" json:"author_transaction_failure,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *RadiusAuthorizationData) Reset()         { *m = RadiusAuthorizationData{} }
func (m *RadiusAuthorizationData) String() string { return proto.CompactTextString(m) }
func (*RadiusAuthorizationData) ProtoMessage()    {}
func (*RadiusAuthorizationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_529eb045ccaae6fc, []int{3}
}

func (m *RadiusAuthorizationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusAuthorizationData.Unmarshal(m, b)
}
func (m *RadiusAuthorizationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusAuthorizationData.Marshal(b, m, deterministic)
}
func (m *RadiusAuthorizationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusAuthorizationData.Merge(m, src)
}
func (m *RadiusAuthorizationData) XXX_Size() int {
	return xxx_messageInfo_RadiusAuthorizationData.Size(m)
}
func (m *RadiusAuthorizationData) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusAuthorizationData.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusAuthorizationData proto.InternalMessageInfo

func (m *RadiusAuthorizationData) GetAuthorRequests() uint32 {
	if m != nil {
		return m.AuthorRequests
	}
	return 0
}

func (m *RadiusAuthorizationData) GetAuthorRequestTimeouts() uint32 {
	if m != nil {
		return m.AuthorRequestTimeouts
	}
	return 0
}

func (m *RadiusAuthorizationData) GetAuthorUnexpectedResponses() uint32 {
	if m != nil {
		return m.AuthorUnexpectedResponses
	}
	return 0
}

func (m *RadiusAuthorizationData) GetAuthorServerErrorResponses() uint32 {
	if m != nil {
		return m.AuthorServerErrorResponses
	}
	return 0
}

func (m *RadiusAuthorizationData) GetAuthorIncorrectResponses() uint32 {
	if m != nil {
		return m.AuthorIncorrectResponses
	}
	return 0
}

func (m *RadiusAuthorizationData) GetAuthorResponseTime() uint32 {
	if m != nil {
		return m.AuthorResponseTime
	}
	return 0
}

func (m *RadiusAuthorizationData) GetAuthorTransactionSuccessess() uint32 {
	if m != nil {
		return m.AuthorTransactionSuccessess
	}
	return 0
}

func (m *RadiusAuthorizationData) GetAuthorTransactionFailure() uint32 {
	if m != nil {
		return m.AuthorTransactionFailure
	}
	return 0
}

type RadiusServerGroupData struct {
	ServerAddress        string                    `protobuf:"bytes,1,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	AuthenticationPort   uint32                    `protobuf:"varint,2,opt,name=authentication_port,json=authenticationPort,proto3" json:"authentication_port,omitempty"`
	AccountingPort       uint32                    `protobuf:"varint,3,opt,name=accounting_port,json=accountingPort,proto3" json:"accounting_port,omitempty"`
	IsPrivate            bool                      `protobuf:"varint,4,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	Accounting           *RadiusAccountingData     `protobuf:"bytes,5,opt,name=accounting,proto3" json:"accounting,omitempty"`
	Authentication       *RadiusAuthenticationData `protobuf:"bytes,6,opt,name=authentication,proto3" json:"authentication,omitempty"`
	Authorization        *RadiusAuthorizationData  `protobuf:"bytes,7,opt,name=authorization,proto3" json:"authorization,omitempty"`
	IpAddress            string                    `protobuf:"bytes,8,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Family               string                    `protobuf:"bytes,9,opt,name=family,proto3" json:"family,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RadiusServerGroupData) Reset()         { *m = RadiusServerGroupData{} }
func (m *RadiusServerGroupData) String() string { return proto.CompactTextString(m) }
func (*RadiusServerGroupData) ProtoMessage()    {}
func (*RadiusServerGroupData) Descriptor() ([]byte, []int) {
	return fileDescriptor_529eb045ccaae6fc, []int{4}
}

func (m *RadiusServerGroupData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusServerGroupData.Unmarshal(m, b)
}
func (m *RadiusServerGroupData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusServerGroupData.Marshal(b, m, deterministic)
}
func (m *RadiusServerGroupData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusServerGroupData.Merge(m, src)
}
func (m *RadiusServerGroupData) XXX_Size() int {
	return xxx_messageInfo_RadiusServerGroupData.Size(m)
}
func (m *RadiusServerGroupData) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusServerGroupData.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusServerGroupData proto.InternalMessageInfo

func (m *RadiusServerGroupData) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

func (m *RadiusServerGroupData) GetAuthenticationPort() uint32 {
	if m != nil {
		return m.AuthenticationPort
	}
	return 0
}

func (m *RadiusServerGroupData) GetAccountingPort() uint32 {
	if m != nil {
		return m.AccountingPort
	}
	return 0
}

func (m *RadiusServerGroupData) GetIsPrivate() bool {
	if m != nil {
		return m.IsPrivate
	}
	return false
}

func (m *RadiusServerGroupData) GetAccounting() *RadiusAccountingData {
	if m != nil {
		return m.Accounting
	}
	return nil
}

func (m *RadiusServerGroupData) GetAuthentication() *RadiusAuthenticationData {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *RadiusServerGroupData) GetAuthorization() *RadiusAuthorizationData {
	if m != nil {
		return m.Authorization
	}
	return nil
}

func (m *RadiusServerGroupData) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *RadiusServerGroupData) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

type RadiusGroupsServers struct {
	Groups               uint32                   `protobuf:"varint,50,opt,name=groups,proto3" json:"groups,omitempty"`
	VrfName              string                   `protobuf:"bytes,51,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	DeadTime             uint32                   `protobuf:"varint,52,opt,name=dead_time,json=deadTime,proto3" json:"dead_time,omitempty"`
	Servers              uint32                   `protobuf:"varint,53,opt,name=servers,proto3" json:"servers,omitempty"`
	ServerGroup          []*RadiusServerGroupData `protobuf:"bytes,54,rep,name=server_group,json=serverGroup,proto3" json:"server_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RadiusGroupsServers) Reset()         { *m = RadiusGroupsServers{} }
func (m *RadiusGroupsServers) String() string { return proto.CompactTextString(m) }
func (*RadiusGroupsServers) ProtoMessage()    {}
func (*RadiusGroupsServers) Descriptor() ([]byte, []int) {
	return fileDescriptor_529eb045ccaae6fc, []int{5}
}

func (m *RadiusGroupsServers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusGroupsServers.Unmarshal(m, b)
}
func (m *RadiusGroupsServers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusGroupsServers.Marshal(b, m, deterministic)
}
func (m *RadiusGroupsServers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusGroupsServers.Merge(m, src)
}
func (m *RadiusGroupsServers) XXX_Size() int {
	return xxx_messageInfo_RadiusGroupsServers.Size(m)
}
func (m *RadiusGroupsServers) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusGroupsServers.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusGroupsServers proto.InternalMessageInfo

func (m *RadiusGroupsServers) GetGroups() uint32 {
	if m != nil {
		return m.Groups
	}
	return 0
}

func (m *RadiusGroupsServers) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *RadiusGroupsServers) GetDeadTime() uint32 {
	if m != nil {
		return m.DeadTime
	}
	return 0
}

func (m *RadiusGroupsServers) GetServers() uint32 {
	if m != nil {
		return m.Servers
	}
	return 0
}

func (m *RadiusGroupsServers) GetServerGroup() []*RadiusServerGroupData {
	if m != nil {
		return m.ServerGroup
	}
	return nil
}

func init() {
	proto.RegisterType((*RadiusGroupsServers_KEYS)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.server_groups.server_group.radius_groups_servers_KEYS")
	proto.RegisterType((*RadiusAccountingData)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.server_groups.server_group.radius_accounting_data")
	proto.RegisterType((*RadiusAuthenticationData)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.server_groups.server_group.radius_authentication_data")
	proto.RegisterType((*RadiusAuthorizationData)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.server_groups.server_group.radius_authorization_data")
	proto.RegisterType((*RadiusServerGroupData)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.server_groups.server_group.radius_server_group_data")
	proto.RegisterType((*RadiusGroupsServers)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.server_groups.server_group.radius_groups_servers")
}

func init() { proto.RegisterFile("radius_groups_servers.proto", fileDescriptor_529eb045ccaae6fc) }

var fileDescriptor_529eb045ccaae6fc = []byte{
	// 1102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xdd, 0x6e, 0x1c, 0x35,
	0x14, 0xd6, 0x66, 0x9b, 0xec, 0xee, 0xd9, 0x6c, 0xb2, 0x31, 0x69, 0x32, 0x49, 0xa8, 0x14, 0x05,
	0x55, 0x94, 0xbf, 0x25, 0xda, 0x42, 0x55, 0x55, 0x80, 0xb4, 0xa0, 0x82, 0x10, 0x12, 0x8a, 0x26,
	0xe1, 0x82, 0x2b, 0xcb, 0x99, 0x71, 0xd2, 0x69, 0x77, 0xc7, 0x83, 0xed, 0x09, 0x2d, 0x8f, 0xc0,
	0x0d, 0x17, 0x95, 0x78, 0x0b, 0xde, 0x81, 0x17, 0xe1, 0x21, 0x78, 0x03, 0xe4, 0x63, 0xcf, 0x8c,
	0x67, 0x76, 0x36, 0x57, 0xe4, 0xa6, 0xca, 0x9c, 0xf3, 0x7d, 0xc7, 0xc7, 0xc7, 0xdf, 0xe7, 0x75,
	0xe1, 0x48, 0xb2, 0x38, 0xc9, 0x15, 0xbd, 0x96, 0x22, 0xcf, 0x14, 0x55, 0x5c, 0xde, 0x70, 0xa9,
	0x26, 0x99, 0x14, 0x5a, 0x90, 0x30, 0x4a, 0x54, 0x24, 0x68, 0x22, 0x14, 0x7d, 0x2d, 0x29, 0x63,
	0x8c, 0x62, 0x3c, 0x12, 0x73, 0xea, 0x68, 0x22, 0xe3, 0x72, 0x62, 0xff, 0x9e, 0xa4, 0x22, 0xe6,
	0xf6, 0xdf, 0x89, 0x2d, 0xe3, 0x8a, 0xd6, 0xbe, 0x4e, 0x38, 0x1c, 0xb6, 0x2e, 0x49, 0x7f, 0x78,
	0xfe, 0xf3, 0x39, 0x39, 0x82, 0x81, 0x29, 0x40, 0x53, 0xb6, 0xe0, 0x41, 0xe7, 0xb8, 0xf3, 0x68,
	0x10, 0xf6, 0x4d, 0xe0, 0x47, 0xb6, 0xe0, 0xe4, 0x43, 0xd8, 0xf1, 0x4b, 0x59, 0xd0, 0x1a, 0x82,
	0xb6, 0x6d, 0xe2, 0x3b, 0x13, 0x37, 0xd8, 0x93, 0x7f, 0xd6, 0x61, 0xcf, 0xad, 0xc3, 0xa2, 0x48,
	0xe4, 0xa9, 0x4e, 0xd2, 0x6b, 0x1a, 0x33, 0xcd, 0xc8, 0x21, 0xf4, 0x25, 0xff, 0x25, 0xe7, 0x4a,
	0x2b, 0x5c, 0x62, 0x14, 0x96, 0xdf, 0xe4, 0x03, 0x18, 0x67, 0x3c, 0x8d, 0x0d, 0xb6, 0xc4, 0xac,
	0x21, 0x66, 0xdb, 0xc5, 0xc3, 0x02, 0x7a, 0x0c, 0x43, 0xc9, 0xb5, 0x64, 0xa9, 0x5a, 0x24, 0x5a,
	0x05, 0x5d, 0x44, 0xf9, 0x21, 0xf2, 0x2e, 0x0c, 0x24, 0x57, 0x99, 0x48, 0x15, 0x57, 0xc1, 0x3d,
	0xcc, 0x57, 0x01, 0xd3, 0x86, 0x4e, 0x16, 0x5c, 0xe4, 0x5a, 0x05, 0xeb, 0xb6, 0x8d, 0xe2, 0x9b,
	0xbc, 0x07, 0xa3, 0x4b, 0x16, 0xd3, 0x8a, 0xbd, 0x81, 0x80, 0xcd, 0x4b, 0x16, 0x87, 0x65, 0x81,
	0x4f, 0x80, 0x18, 0x10, 0xcb, 0xf5, 0x0b, 0x9e, 0xea, 0x24, 0x62, 0x5a, 0x48, 0x15, 0xf4, 0x10,
	0xb9, 0x73, 0xc9, 0xe2, 0x59, 0x2d, 0x41, 0x4e, 0x61, 0x37, 0x4f, 0x5f, 0xa5, 0xe2, 0xd7, 0x94,
	0x66, 0x2c, 0x7a, 0xc5, 0x35, 0xd5, 0x6f, 0x32, 0xae, 0x82, 0x3e, 0x12, 0x88, 0xcb, 0x9d, 0x61,
	0xea, 0xc2, 0x64, 0xc8, 0x47, 0xb0, 0x13, 0x4b, 0x91, 0x65, 0xdc, 0xef, 0x64, 0x80, 0xf0, 0xb1,
	0x4b, 0x54, 0xdd, 0x8c, 0xa1, 0x2b, 0xb5, 0x0e, 0x00, 0xd3, 0xe6, 0x4f, 0xf2, 0x0c, 0x0e, 0x58,
	0x14, 0x69, 0x9a, 0xa7, 0xfc, 0x75, 0xc6, 0x23, 0x5d, 0x2b, 0x33, 0x44, 0xdc, 0xbe, 0x01, 0xfc,
	0x54, 0xe6, 0xab, 0x6a, 0x5f, 0xc2, 0x11, 0x72, 0xdd, 0x79, 0x73, 0x29, 0x85, 0xf4, 0xd8, 0x9b,
	0xc8, 0x0e, 0x0c, 0xe4, 0x1c, 0x11, 0xcf, 0x0d, 0xa0, 0xa2, 0x3f, 0x05, 0xcc, 0xd1, 0x24, 0x8d,
	0x84, 0x94, 0x3c, 0xd2, 0x1e, 0x77, 0x84, 0xdc, 0x3d, 0x93, 0xff, 0xbe, 0x48, 0x57, 0xcc, 0x8f,
	0x81, 0x20, 0xb3, 0xc0, 0x53, 0x73, 0x26, 0xc1, 0x96, 0xdd, 0xb4, 0xc9, 0x14, 0xd0, 0x8b, 0x64,
	0xc1, 0xc9, 0x57, 0xae, 0x4d, 0x3c, 0x73, 0x16, 0xe9, 0x44, 0xa4, 0x54, 0xe5, 0x51, 0xc4, 0x95,
	0xe2, 0x4a, 0x05, 0xdb, 0x48, 0xc3, 0x29, 0x5c, 0x54, 0x88, 0xf3, 0x12, 0x50, 0xf6, 0xe9, 0xf3,
	0xaf, 0x58, 0x32, 0xcf, 0x25, 0x0f, 0xc6, 0x55, 0x9f, 0x1e, 0xf9, 0x5b, 0x9b, 0x3d, 0xf9, 0xbb,
	0x57, 0xfa, 0xc8, 0x13, 0x80, 0xe1, 0xa3, 0xc6, 0xdf, 0x87, 0x6d, 0x86, 0xab, 0xd0, 0x86, 0xd4,
	0xb7, 0x6c, 0xb8, 0x54, 0xf1, 0x13, 0xd8, 0x2f, 0x04, 0xdf, 0x24, 0x58, 0xdd, 0xdf, 0x77, 0xe9,
	0x59, 0x9d, 0xf7, 0x05, 0x1c, 0xd6, 0xf1, 0x74, 0xd9, 0x0c, 0x41, 0x6d, 0xad, 0xd0, 0x73, 0xc6,
	0x43, 0x70, 0x7d, 0xe0, 0xa2, 0x99, 0x2e, 0xec, 0x31, 0xb2, 0xd1, 0x99, 0x0d, 0x7a, 0x30, 0xc9,
	0x5f, 0xf2, 0xa8, 0x34, 0xca, 0xa8, 0x28, 0x8c, 0x41, 0xa3, 0x53, 0x07, 0x8b, 0x5e, 0xb0, 0xf9,
	0x9c, 0xa7, 0xd7, 0xa5, 0x63, 0xc6, 0x36, 0xf1, 0x4d, 0x19, 0xf7, 0x26, 0x53, 0xba, 0xaf, 0xe7,
	0x4f, 0xe6, 0xa2, 0xf0, 0xe0, 0x29, 0xec, 0xa2, 0xbd, 0x8a, 0x06, 0x0a, 0xfd, 0x38, 0xbf, 0x18,
	0x83, 0xb9, 0x2e, 0x0a, 0xed, 0x3c, 0x83, 0x03, 0x8f, 0xd1, 0xf0, 0xa5, 0xf5, 0xcd, 0x7e, 0x49,
	0x5b, 0xed, 0xce, 0xa2, 0x3d, 0x74, 0x27, 0xd4, 0xdc, 0x69, 0xa9, 0xd6, 0x9d, 0x4f, 0x21, 0x28,
	0xdc, 0xb9, 0xd4, 0xa3, 0x75, 0xd7, 0x9e, 0xcb, 0x37, 0xfb, 0x74, 0x56, 0xdd, 0xac, 0xac, 0x7a,
	0x0a, 0xbb, 0xb6, 0xdd, 0x86, 0xee, 0xad, 0x57, 0x88, 0xcd, 0xd5, 0x94, 0xff, 0x35, 0x3c, 0x70,
	0x8c, 0x15, 0xda, 0xb7, 0x96, 0x39, 0xb2, 0xa0, 0x76, 0xf5, 0x1b, 0x0d, 0x2d, 0xd7, 0x28, 0xf4,
	0xbf, 0xed, 0x34, 0xd4, 0x2c, 0xe0, 0x1c, 0x80, 0xde, 0xb3, 0xec, 0xd6, 0x0b, 0x66, 0xec, 0xbc,
	0x87, 0x90, 0xb6, 0x2b, 0x66, 0x56, 0xee, 0x60, 0xc5, 0x25, 0xb3, 0x83, 0x15, 0x5c, 0x8b, 0xad,
	0xd7, 0x4c, 0xb5, 0x81, 0xb6, 0x8b, 0x86, 0xf8, 0x1b, 0x58, 0xbe, 0x6a, 0x4e, 0xfe, 0xed, 0xc2,
	0x81, 0x67, 0x61, 0x21, 0x93, 0xdf, 0x1a, 0x0e, 0xc6, 0xe8, 0xb2, 0x83, 0x31, 0xec, 0x3b, 0xb8,
	0x0e, 0xac, 0x84, 0xed, 0x1c, 0x5c, 0x23, 0x94, 0xfa, 0x76, 0xf3, 0x13, 0xb2, 0x7d, 0x7e, 0xdd,
	0x6a, 0x7e, 0x42, 0xde, 0x32, 0x3f, 0x21, 0x57, 0xcd, 0xef, 0x5e, 0x35, 0x3f, 0x21, 0x6f, 0x9b,
	0x9f, 0x90, 0xad, 0xf3, 0x5b, 0xaf, 0xe6, 0x27, 0x64, 0xcb, 0x55, 0xed, 0x44, 0xeb, 0xad, 0x69,
	0x45, 0xbb, 0x51, 0x89, 0xb6, 0x5a, 0xcc, 0x17, 0xad, 0x90, 0xab, 0x44, 0xdb, 0xab, 0x44, 0x2b,
	0xe4, 0xad, 0xa2, 0x6d, 0xd4, 0x28, 0x44, 0xdb, 0xf7, 0x7b, 0x6e, 0xb9, 0xb6, 0xff, 0x5a, 0x87,
	0xc0, 0x9d, 0x79, 0xed, 0x29, 0x83, 0x47, 0xfe, 0x10, 0xb6, 0x5c, 0x90, 0xc5, 0xb1, 0x34, 0xfd,
	0xd8, 0x17, 0xd0, 0xc8, 0x46, 0x67, 0x36, 0x48, 0x3e, 0x85, 0x77, 0x1a, 0x57, 0x7e, 0x26, 0xa4,
	0x76, 0x87, 0x4d, 0xea, 0xa9, 0x33, 0x21, 0xb5, 0xbb, 0xf2, 0x8a, 0x37, 0x10, 0x82, 0xbb, 0xe5,
	0x95, 0xe7, 0xc2, 0x08, 0x7c, 0x00, 0x90, 0x28, 0x9a, 0xc9, 0xe4, 0x86, 0x69, 0x8e, 0xe7, 0xd7,
	0x0f, 0x07, 0x89, 0x3a, 0xb3, 0x01, 0xf2, 0x7b, 0x07, 0xa0, 0x62, 0xe0, 0xf9, 0x0c, 0xa7, 0x2f,
	0x27, 0xff, 0xff, 0x23, 0x71, 0xd2, 0xfe, 0x72, 0x0b, 0xbd, 0xd5, 0xc9, 0x9f, 0x1d, 0xd8, 0xaa,
	0xef, 0x15, 0x0f, 0x7e, 0x38, 0x4d, 0xef, 0xb2, 0xa1, 0xe5, 0x9f, 0xda, 0xb0, 0xd1, 0x05, 0x79,
	0xdb, 0x81, 0x51, 0xcd, 0xcf, 0xa8, 0xaa, 0xe1, 0x74, 0x71, 0xc7, 0x7d, 0xd5, 0xef, 0x8f, 0xb0,
	0xde, 0x03, 0x1e, 0x6d, 0x56, 0xea, 0xaa, 0x8f, 0xba, 0x1a, 0x24, 0x59, 0xa1, 0xa9, 0x3d, 0xd8,
	0xb8, 0x62, 0x8b, 0x64, 0xfe, 0x06, 0x7f, 0xa7, 0x06, 0xa1, 0xfb, 0x3a, 0x79, 0xbb, 0x06, 0xf7,
	0x5b, 0x9f, 0xeb, 0x86, 0x61, 0x23, 0xc1, 0x14, 0xb5, 0xe4, 0xbe, 0xc8, 0x01, 0xf4, 0x6f, 0xe4,
	0x95, 0x7d, 0x9b, 0x3f, 0xc6, 0x5a, 0xbd, 0x1b, 0x79, 0x85, 0xef, 0xf7, 0x23, 0x18, 0xc4, 0x9c,
	0xc5, 0xd6, 0xa5, 0x9f, 0xd9, 0x27, 0xaf, 0x09, 0xa0, 0x37, 0x03, 0xe8, 0xb9, 0xd2, 0xc1, 0xe7,
	0x98, 0x2a, 0x3e, 0xc9, 0x1f, 0x1d, 0xd8, 0xf4, 0x37, 0x1d, 0x3c, 0x39, 0xee, 0x3e, 0x1a, 0x4e,
	0xe7, 0x77, 0x38, 0xcf, 0x25, 0x6f, 0x86, 0x43, 0xef, 0x3f, 0x18, 0x97, 0x1b, 0xb8, 0xd0, 0xe3,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x06, 0xdd, 0x5d, 0x2d, 0x3d, 0x0d, 0x00, 0x00,
}
