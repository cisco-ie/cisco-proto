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
	AcctTransactionSuccessess uint32   `protobuf:"varint,12,opt,name=acct_transaction_successess,json=acctTransactionSuccessess,proto3" json:"acct_transaction_successess,omitempty"`
	AcctTransactionFailure    uint32   `protobuf:"varint,13,opt,name=acct_transaction_failure,json=acctTransactionFailure,proto3" json:"acct_transaction_failure,omitempty"`
	AcctThrottledTransactions uint32   `protobuf:"varint,14,opt,name=acct_throttled_transactions,json=acctThrottledTransactions,proto3" json:"acct_throttled_transactions,omitempty"`
	AcctMaxThrottleTrans      uint32   `protobuf:"varint,15,opt,name=acct_max_throttle_trans,json=acctMaxThrottleTrans,proto3" json:"acct_max_throttle_trans,omitempty"`
	TotalTestAcctReqs         uint32   `protobuf:"varint,16,opt,name=total_test_acct_reqs,json=totalTestAcctReqs,proto3" json:"total_test_acct_reqs,omitempty"`
	TotalTestAcctTimeouts     uint32   `protobuf:"varint,17,opt,name=total_test_acct_timeouts,json=totalTestAcctTimeouts,proto3" json:"total_test_acct_timeouts,omitempty"`
	TotalTestAcctResponse     uint32   `protobuf:"varint,18,opt,name=total_test_acct_response,json=totalTestAcctResponse,proto3" json:"total_test_acct_response,omitempty"`
	TotalTestAcctPending      uint32   `protobuf:"varint,19,opt,name=total_test_acct_pending,json=totalTestAcctPending,proto3" json:"total_test_acct_pending,omitempty"`
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

func (m *RadiusAccountingData) GetAcctThrottledTransactions() uint32 {
	if m != nil {
		return m.AcctThrottledTransactions
	}
	return 0
}

func (m *RadiusAccountingData) GetAcctMaxThrottleTrans() uint32 {
	if m != nil {
		return m.AcctMaxThrottleTrans
	}
	return 0
}

func (m *RadiusAccountingData) GetTotalTestAcctReqs() uint32 {
	if m != nil {
		return m.TotalTestAcctReqs
	}
	return 0
}

func (m *RadiusAccountingData) GetTotalTestAcctTimeouts() uint32 {
	if m != nil {
		return m.TotalTestAcctTimeouts
	}
	return 0
}

func (m *RadiusAccountingData) GetTotalTestAcctResponse() uint32 {
	if m != nil {
		return m.TotalTestAcctResponse
	}
	return 0
}

func (m *RadiusAccountingData) GetTotalTestAcctPending() uint32 {
	if m != nil {
		return m.TotalTestAcctPending
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
	AuthenTransactionSuccessess uint32   `protobuf:"varint,13,opt,name=authen_transaction_successess,json=authenTransactionSuccessess,proto3" json:"authen_transaction_successess,omitempty"`
	AuthenTransactionFailure    uint32   `protobuf:"varint,14,opt,name=authen_transaction_failure,json=authenTransactionFailure,proto3" json:"authen_transaction_failure,omitempty"`
	AuthenUnexpectedResponses   uint32   `protobuf:"varint,15,opt,name=authen_unexpected_responses,json=authenUnexpectedResponses,proto3" json:"authen_unexpected_responses,omitempty"`
	AuthenServerErrorResponses  uint32   `protobuf:"varint,16,opt,name=authen_server_error_responses,json=authenServerErrorResponses,proto3" json:"authen_server_error_responses,omitempty"`
	AuthenIncorrectResponses    uint32   `protobuf:"varint,17,opt,name=authen_incorrect_responses,json=authenIncorrectResponses,proto3" json:"authen_incorrect_responses,omitempty"`
	AuthThrottledTransactions   uint32   `protobuf:"varint,18,opt,name=auth_throttled_transactions,json=authThrottledTransactions,proto3" json:"auth_throttled_transactions,omitempty"`
	AuthMaxTransactions         uint32   `protobuf:"varint,19,opt,name=auth_max_transactions,json=authMaxTransactions,proto3" json:"auth_max_transactions,omitempty"`
	TotalTestAuthReqs           uint32   `protobuf:"varint,20,opt,name=total_test_auth_reqs,json=totalTestAuthReqs,proto3" json:"total_test_auth_reqs,omitempty"`
	TotalTestAuthTimeouts       uint32   `protobuf:"varint,21,opt,name=total_test_auth_timeouts,json=totalTestAuthTimeouts,proto3" json:"total_test_auth_timeouts,omitempty"`
	TotalTestAuthResponse       uint32   `protobuf:"varint,22,opt,name=total_test_auth_response,json=totalTestAuthResponse,proto3" json:"total_test_auth_response,omitempty"`
	TotalTestAuthPending        uint32   `protobuf:"varint,23,opt,name=total_test_auth_pending,json=totalTestAuthPending,proto3" json:"total_test_auth_pending,omitempty"`
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

func (m *RadiusAuthenticationData) GetAuthThrottledTransactions() uint32 {
	if m != nil {
		return m.AuthThrottledTransactions
	}
	return 0
}

func (m *RadiusAuthenticationData) GetAuthMaxTransactions() uint32 {
	if m != nil {
		return m.AuthMaxTransactions
	}
	return 0
}

func (m *RadiusAuthenticationData) GetTotalTestAuthReqs() uint32 {
	if m != nil {
		return m.TotalTestAuthReqs
	}
	return 0
}

func (m *RadiusAuthenticationData) GetTotalTestAuthTimeouts() uint32 {
	if m != nil {
		return m.TotalTestAuthTimeouts
	}
	return 0
}

func (m *RadiusAuthenticationData) GetTotalTestAuthResponse() uint32 {
	if m != nil {
		return m.TotalTestAuthResponse
	}
	return 0
}

func (m *RadiusAuthenticationData) GetTotalTestAuthPending() uint32 {
	if m != nil {
		return m.TotalTestAuthPending
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
	AuthenticationPort   uint32                    `protobuf:"varint,1,opt,name=authentication_port,json=authenticationPort,proto3" json:"authentication_port,omitempty"`
	AccountingPort       uint32                    `protobuf:"varint,2,opt,name=accounting_port,json=accountingPort,proto3" json:"accounting_port,omitempty"`
	IsPrivate            bool                      `protobuf:"varint,3,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	Accounting           *RadiusAccountingData     `protobuf:"bytes,4,opt,name=accounting,proto3" json:"accounting,omitempty"`
	Authentication       *RadiusAuthenticationData `protobuf:"bytes,5,opt,name=authentication,proto3" json:"authentication,omitempty"`
	Authorization        *RadiusAuthorizationData  `protobuf:"bytes,6,opt,name=authorization,proto3" json:"authorization,omitempty"`
	IpAddress            string                    `protobuf:"bytes,7,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Family               string                    `protobuf:"bytes,8,opt,name=family,proto3" json:"family,omitempty"`
	RedirectedRequests   uint32                    `protobuf:"varint,9,opt,name=redirected_requests,json=redirectedRequests,proto3" json:"redirected_requests,omitempty"`
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

func (m *RadiusServerGroupData) GetRedirectedRequests() uint32 {
	if m != nil {
		return m.RedirectedRequests
	}
	return 0
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
	// 1239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xcf, 0x6e, 0xdb, 0xc6,
	0x13, 0x86, 0xec, 0x24, 0x92, 0x46, 0x96, 0xff, 0x6c, 0x64, 0x9b, 0x8e, 0x7f, 0x01, 0x02, 0xff,
	0x50, 0x34, 0x6d, 0x51, 0x25, 0x50, 0x9a, 0x34, 0x08, 0x8a, 0x02, 0x6a, 0x91, 0x16, 0x45, 0xd1,
	0xc2, 0xa0, 0xdd, 0x43, 0x4f, 0xc4, 0x9a, 0x5c, 0x5b, 0x4c, 0x24, 0x2e, 0xbd, 0xbb, 0x74, 0x9d,
	0xbe, 0x40, 0x81, 0x5e, 0x7a, 0x08, 0xd0, 0x57, 0xe9, 0xcb, 0xf4, 0x25, 0xfa, 0x06, 0xc5, 0xce,
	0xee, 0x92, 0x4b, 0x8a, 0xf2, 0xa9, 0xbe, 0x18, 0xde, 0x99, 0xef, 0x9b, 0x9d, 0xe1, 0xcc, 0x7c,
	0x22, 0xe1, 0x50, 0xd0, 0x24, 0x2d, 0x64, 0x74, 0x21, 0x78, 0x91, 0xcb, 0x48, 0x32, 0x71, 0xc5,
	0x84, 0x1c, 0xe7, 0x82, 0x2b, 0x4e, 0xc2, 0x38, 0x95, 0x31, 0x8f, 0x52, 0x2e, 0xa3, 0x6b, 0x11,
	0x51, 0x4a, 0x23, 0xb4, 0xc7, 0x7c, 0x1e, 0x59, 0x1a, 0xcf, 0x99, 0x18, 0x9b, 0xff, 0xc7, 0x19,
	0x4f, 0x98, 0xf9, 0x3b, 0x36, 0x61, 0x6c, 0xd0, 0xda, 0xe9, 0x88, 0xc1, 0x83, 0xd6, 0x2b, 0xa3,
	0xef, 0x5f, 0xff, 0x7c, 0x42, 0x0e, 0xa1, 0xaf, 0x03, 0x44, 0x19, 0x5d, 0xb0, 0xa0, 0xf3, 0xa8,
	0xf3, 0xb8, 0x1f, 0xf6, 0xb4, 0xe1, 0x47, 0xba, 0x60, 0xe4, 0x63, 0xd8, 0xf1, 0x43, 0x19, 0xd0,
	0x1a, 0x82, 0xb6, 0x8c, 0xe3, 0x5b, 0x6d, 0xd7, 0xd8, 0xa3, 0xdf, 0xba, 0xb0, 0x67, 0xef, 0xa1,
	0x71, 0xcc, 0x8b, 0x4c, 0xa5, 0xd9, 0x45, 0x94, 0x50, 0x45, 0xc9, 0x03, 0xe8, 0x09, 0x76, 0x59,
	0x30, 0xa9, 0x24, 0x5e, 0x31, 0x0c, 0xcb, 0x33, 0xf9, 0x08, 0xb6, 0x73, 0x96, 0x25, 0x1a, 0x5b,
	0x62, 0xd6, 0x10, 0xb3, 0x65, 0xed, 0xa1, 0x83, 0x3e, 0x82, 0x81, 0x60, 0x4a, 0xd0, 0x4c, 0x2e,
	0x52, 0x25, 0x83, 0x75, 0x44, 0xf9, 0x26, 0xf2, 0x3f, 0xe8, 0x0b, 0x26, 0x73, 0x9e, 0x49, 0x26,
	0x83, 0x3b, 0xe8, 0xaf, 0x0c, 0x3a, 0x0d, 0x95, 0x2e, 0x18, 0x2f, 0x94, 0x0c, 0xee, 0x9a, 0x34,
	0xdc, 0x99, 0xfc, 0x1f, 0x86, 0x67, 0x34, 0x89, 0x2a, 0xf6, 0x3d, 0x04, 0x6c, 0x9c, 0xd1, 0x24,
	0x2c, 0x03, 0x7c, 0x0a, 0x44, 0x83, 0x68, 0xa1, 0x66, 0x2c, 0x53, 0x69, 0x4c, 0x15, 0x17, 0x32,
	0xe8, 0x22, 0x72, 0xe7, 0x8c, 0x26, 0xd3, 0x9a, 0x83, 0x3c, 0x85, 0x51, 0x91, 0xbd, 0xcd, 0xf8,
	0x2f, 0x59, 0x94, 0xd3, 0xf8, 0x2d, 0x53, 0x91, 0x7a, 0x97, 0x33, 0x19, 0xf4, 0x90, 0x40, 0xac,
	0xef, 0x18, 0x5d, 0xa7, 0xda, 0x43, 0x3e, 0x81, 0x9d, 0x44, 0xf0, 0x3c, 0x67, 0x7e, 0x26, 0x7d,
	0x84, 0x6f, 0x5b, 0x47, 0x95, 0xcd, 0x36, 0xac, 0x0b, 0xa5, 0x02, 0x40, 0xb7, 0xfe, 0x97, 0xbc,
	0x82, 0x03, 0x1a, 0xc7, 0x2a, 0x2a, 0x32, 0x76, 0x9d, 0xb3, 0x58, 0xd5, 0xc2, 0x0c, 0x10, 0xb7,
	0xaf, 0x01, 0x3f, 0x95, 0xfe, 0x2a, 0xda, 0x97, 0x70, 0x88, 0x5c, 0x7c, 0x98, 0x34, 0x56, 0x29,
	0xcf, 0x22, 0x59, 0xc4, 0x31, 0x93, 0x92, 0x49, 0x19, 0x6c, 0x20, 0x1b, 0xc3, 0x9f, 0x56, 0x88,
	0x93, 0x12, 0x40, 0x5e, 0x42, 0xb0, 0xc4, 0x3f, 0xa7, 0xe9, 0xbc, 0x10, 0x2c, 0x18, 0x22, 0x79,
	0xaf, 0x41, 0xfe, 0xc6, 0x78, 0xab, 0x9b, 0x67, 0x82, 0x2b, 0x35, 0x67, 0x89, 0x1f, 0x43, 0x06,
	0x9b, 0xde, 0xcd, 0x0e, 0xe1, 0x45, 0x91, 0xe4, 0x39, 0x60, 0x51, 0xd1, 0x82, 0x5e, 0x97, 0x31,
	0x4c, 0x88, 0x60, 0x0b, 0xb9, 0x23, 0xed, 0xfe, 0x81, 0x5e, 0x3b, 0x3a, 0xb2, 0xc9, 0x13, 0x18,
	0x29, 0xae, 0xe8, 0x3c, 0x52, 0x4c, 0xaa, 0x08, 0x23, 0x08, 0x76, 0x29, 0x83, 0x6d, 0xd3, 0x4e,
	0xf4, 0x9d, 0x32, 0xa9, 0xa6, 0x71, 0xac, 0x42, 0x76, 0x29, 0xc9, 0xe7, 0x10, 0x34, 0x09, 0xe5,
	0x38, 0xed, 0x20, 0x69, 0xb7, 0x46, 0x3a, 0x75, 0xb3, 0xd5, 0x42, 0x74, 0x6d, 0x09, 0x48, 0x0b,
	0xd1, 0x35, 0x45, 0x57, 0xd6, 0x24, 0xda, 0x9d, 0x08, 0xee, 0x9b, 0xca, 0x6a, 0xbc, 0x63, 0xe3,
	0x3b, 0xfa, 0xbb, 0x5f, 0x6e, 0xbc, 0x37, 0xaa, 0xba, 0x21, 0xb8, 0x8d, 0x1f, 0xc2, 0x16, 0xc5,
	0xb6, 0x45, 0x8d, 0xa5, 0xdc, 0x34, 0xe6, 0x72, 0xdf, 0x5e, 0xc0, 0xbe, 0x5b, 0xcd, 0x26, 0xc1,
	0x6c, 0xe8, 0xae, 0x75, 0x4f, 0xeb, 0xbc, 0x2f, 0xe0, 0x41, 0x1d, 0x1f, 0x2d, 0xaf, 0x6d, 0x50,
	0xbb, 0x2b, 0xf4, 0x76, 0xf8, 0x03, 0xb0, 0x79, 0xe0, 0xa5, 0xb9, 0x72, 0x8b, 0x3c, 0x34, 0xd6,
	0xa9, 0x31, 0x7a, 0x30, 0xc1, 0xde, 0xb0, 0xb8, 0x5c, 0xe9, 0xa1, 0x0b, 0x8c, 0x46, 0xbd, 0x51,
	0x16, 0x16, 0xcf, 0xe8, 0x7c, 0xce, 0xb2, 0x8b, 0x72, 0xb7, 0xb7, 0x8d, 0xe3, 0xeb, 0xd2, 0xee,
	0x3d, 0x99, 0xb2, 0xb1, 0x5d, 0xff, 0xc9, 0x94, 0x1d, 0x7d, 0x0a, 0x23, 0x14, 0x02, 0x97, 0x80,
	0xdb, 0x31, 0xbb, 0xd9, 0x5a, 0x0a, 0x6c, 0x16, 0x6e, 0xbd, 0x5e, 0xc1, 0x81, 0xc7, 0x68, 0x28,
	0x88, 0xd9, 0xf0, 0xfd, 0x92, 0xb6, 0x5a, 0x47, 0x5c, 0x7a, 0xa8, 0x23, 0x50, 0xd3, 0x11, 0x43,
	0x35, 0x3a, 0xf2, 0x12, 0x02, 0xa7, 0x23, 0x4b, 0x39, 0x1a, 0x1d, 0xd8, 0xb3, 0xfe, 0x66, 0x9e,
	0x56, 0x54, 0x36, 0x2a, 0x51, 0xf9, 0x0a, 0x1e, 0x9a, 0x74, 0x57, 0x49, 0x83, 0xd9, 0xee, 0x43,
	0x03, 0x6a, 0x17, 0x07, 0x3d, 0x11, 0xcb, 0x31, 0x9c, 0x3c, 0x6c, 0xda, 0x89, 0x68, 0x06, 0xf0,
	0x05, 0xc2, 0xb0, 0x5b, 0x85, 0x6d, 0xcb, 0x0a, 0x04, 0x42, 0xda, 0xa4, 0x6d, 0x5a, 0x56, 0x60,
	0x7f, 0xcc, 0x98, 0x10, 0x5c, 0x78, 0x11, 0xcc, 0xca, 0xdb, 0x14, 0x4f, 0x10, 0xf3, 0x5a, 0x43,
	0xaa, 0x10, 0x55, 0x01, 0x69, 0x16, 0x73, 0x21, 0x98, 0xb7, 0xc3, 0x6e, 0xfb, 0x6d, 0x01, 0xdf,
	0x39, 0x40, 0x5d, 0x5b, 0x0b, 0x35, 0x5b, 0xa5, 0x70, 0xa4, 0x2a, 0xa0, 0x5d, 0xe1, 0x26, 0xb0,
	0x8b, 0x7c, 0x54, 0x38, 0x9f, 0x69, 0x54, 0xe0, 0xbe, 0x76, 0x6a, 0x7d, 0xf3, 0x39, 0x0d, 0x79,
	0xd3, 0x74, 0x94, 0xb7, 0x51, 0x53, 0xde, 0x0a, 0x35, 0x6b, 0x93, 0x37, 0xcc, 0xd7, 0x6d, 0xc1,
	0x6e, 0x53, 0xa5, 0x74, 0xaa, 0x2b, 0xe4, 0xcd, 0xdc, 0x64, 0xe5, 0x6d, 0xaf, 0x85, 0xb8, 0x4a,
	0xde, 0x34, 0xd1, 0xc9, 0xdb, 0x7e, 0x53, 0xde, 0x0a, 0x35, 0x73, 0xf2, 0xf6, 0xcf, 0x3a, 0x1c,
	0x78, 0xf2, 0xc6, 0x45, 0xfa, 0x6b, 0x43, 0xdd, 0xd0, 0xba, 0xac, 0x6e, 0x68, 0xf6, 0xd5, 0xad,
	0x0e, 0xac, 0xca, 0xb5, 0xea, 0x56, 0x23, 0x94, 0xe5, 0xda, 0x66, 0x72, 0xd1, 0x3e, 0x8d, 0xeb,
	0x55, 0x33, 0xb9, 0xb8, 0x61, 0x1a, 0xb9, 0x58, 0x35, 0x8d, 0x77, 0xaa, 0x69, 0xe4, 0xe2, 0xa6,
	0x69, 0xe4, 0xa2, 0x75, 0x1a, 0xef, 0x56, 0xd3, 0xc8, 0x45, 0xcb, 0x34, 0x3e, 0x85, 0x51, 0x59,
	0xb8, 0xb1, 0x61, 0xe5, 0x56, 0x15, 0x89, 0xab, 0xda, 0xb8, 0x74, 0xd9, 0x4e, 0x02, 0xb8, 0x58,
	0x25, 0x01, 0xdd, 0x4a, 0x02, 0xb8, 0xb8, 0x51, 0x02, 0x1a, 0x31, 0x9c, 0x04, 0xf4, 0xfc, 0x9c,
	0x97, 0x25, 0xe0, 0xe8, 0xaf, 0xbb, 0x10, 0xd8, 0x9e, 0xd7, 0x5e, 0x48, 0xb1, 0xe5, 0x4f, 0xe0,
	0x7e, 0xe3, 0x77, 0x2e, 0xe7, 0x42, 0xd9, 0xb6, 0x93, 0xba, 0xeb, 0x98, 0x0b, 0x65, 0x75, 0xde,
	0xbd, 0xa2, 0x22, 0x78, 0xad, 0xd4, 0x79, 0x6b, 0x46, 0xe0, 0x43, 0x80, 0x54, 0x46, 0xb9, 0x48,
	0xaf, 0xa8, 0x62, 0xd8, 0xda, 0x5e, 0xd8, 0x4f, 0xe5, 0xb1, 0x31, 0x90, 0xdf, 0x3b, 0x00, 0x15,
	0x03, 0x1b, 0x37, 0x98, 0xbc, 0x19, 0xff, 0xf7, 0xef, 0xf0, 0xe3, 0xf6, 0x17, 0xeb, 0xd0, 0xbb,
	0x9d, 0xfc, 0xd9, 0x81, 0xcd, 0x7a, 0xad, 0x38, 0x09, 0x83, 0x49, 0x76, 0x9b, 0x09, 0x2d, 0xbf,
	0x5f, 0x84, 0x8d, 0x2c, 0xc8, 0xfb, 0x0e, 0x0c, 0x6b, 0x8b, 0x8a, 0x93, 0x36, 0x98, 0x2c, 0x6e,
	0x39, 0xaf, 0xba, 0x30, 0x84, 0xf5, 0x1c, 0xb0, 0xb5, 0x79, 0x44, 0x93, 0x44, 0xb8, 0x01, 0xee,
	0x87, 0xfd, 0x34, 0x9f, 0x1a, 0x03, 0xd9, 0x83, 0x7b, 0xe7, 0x74, 0x91, 0xce, 0xdf, 0xe1, 0x68,
	0xf6, 0x43, 0x7b, 0xd2, 0xb3, 0x26, 0x58, 0x92, 0x0a, 0xb7, 0xf6, 0x56, 0x62, 0xcc, 0x2f, 0x38,
	0xa9, 0x5c, 0x4e, 0x66, 0x8e, 0xde, 0xaf, 0xc1, 0x6e, 0xeb, 0xe7, 0x97, 0xbe, 0xc2, 0x58, 0x82,
	0x09, 0xb2, 0xed, 0x89, 0x1c, 0x40, 0xef, 0x4a, 0x9c, 0x9b, 0x6f, 0xad, 0x67, 0x78, 0x79, 0xf7,
	0x4a, 0x9c, 0xe3, 0xf7, 0xd8, 0x21, 0xf4, 0x13, 0x46, 0x13, 0xb3, 0xaf, 0x9f, 0x99, 0x4f, 0x18,
	0x6d, 0xc0, 0x2d, 0x0d, 0xa0, 0x6b, 0x43, 0x07, 0xcf, 0xd1, 0xe5, 0x8e, 0xe4, 0x8f, 0x0e, 0x6c,
	0xf8, 0x4f, 0x29, 0x78, 0xf1, 0x68, 0xfd, 0xf1, 0x60, 0x32, 0xbf, 0xc5, 0x06, 0x2c, 0x6d, 0x69,
	0x38, 0xf0, 0x3e, 0x18, 0xcf, 0xee, 0xe1, 0x45, 0xcf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x13,
	0xbc, 0xae, 0xcb, 0x0d, 0x0f, 0x00, 0x00,
}