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
// source: subscriber_session_data.proto

package cisco_ios_xr_iedge4710_oper_subscriber_session_nodes_node_sessions_session

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

type SubscriberSessionData_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	SessionId            string   `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriberSessionData_KEYS) Reset()         { *m = SubscriberSessionData_KEYS{} }
func (m *SubscriberSessionData_KEYS) String() string { return proto.CompactTextString(m) }
func (*SubscriberSessionData_KEYS) ProtoMessage()    {}
func (*SubscriberSessionData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e0d9a3f80b04023, []int{0}
}

func (m *SubscriberSessionData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberSessionData_KEYS.Unmarshal(m, b)
}
func (m *SubscriberSessionData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberSessionData_KEYS.Marshal(b, m, deterministic)
}
func (m *SubscriberSessionData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberSessionData_KEYS.Merge(m, src)
}
func (m *SubscriberSessionData_KEYS) XXX_Size() int {
	return xxx_messageInfo_SubscriberSessionData_KEYS.Size(m)
}
func (m *SubscriberSessionData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberSessionData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberSessionData_KEYS proto.InternalMessageInfo

func (m *SubscriberSessionData_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *SubscriberSessionData_KEYS) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type IedgeOperAccountingData struct {
	AccountingStateRc                    uint32   `protobuf:"varint,1,opt,name=accounting_state_rc,json=accountingStateRc,proto3" json:"accounting_state_rc,omitempty"`
	AccountingStopState                  uint32   `protobuf:"varint,2,opt,name=accounting_stop_state,json=accountingStopState,proto3" json:"accounting_stop_state,omitempty"`
	RecordContextName                    string   `protobuf:"bytes,3,opt,name=record_context_name,json=recordContextName,proto3" json:"record_context_name,omitempty"`
	MethodListName                       string   `protobuf:"bytes,4,opt,name=method_list_name,json=methodListName,proto3" json:"method_list_name,omitempty"`
	AccountSessionId                     string   `protobuf:"bytes,5,opt,name=account_session_id,json=accountSessionId,proto3" json:"account_session_id,omitempty"`
	AccountingStartTimeEpoch             uint64   `protobuf:"varint,6,opt,name=accounting_start_time_epoch,json=accountingStartTimeEpoch,proto3" json:"accounting_start_time_epoch,omitempty"`
	IsInterimAccountingEnabled           bool     `protobuf:"varint,7,opt,name=is_interim_accounting_enabled,json=isInterimAccountingEnabled,proto3" json:"is_interim_accounting_enabled,omitempty"`
	InterimInterval                      uint32   `protobuf:"varint,8,opt,name=interim_interval,json=interimInterval,proto3" json:"interim_interval,omitempty"`
	LastSuccessfulInterimUpdateTimeEpoch uint64   `protobuf:"varint,9,opt,name=last_successful_interim_update_time_epoch,json=lastSuccessfulInterimUpdateTimeEpoch,proto3" json:"last_successful_interim_update_time_epoch,omitempty"`
	NextInterimUpdateAttemptTime         uint32   `protobuf:"varint,10,opt,name=next_interim_update_attempt_time,json=nextInterimUpdateAttemptTime,proto3" json:"next_interim_update_attempt_time,omitempty"`
	LastInterimUpdateAttemptTimeEpoch    uint64   `protobuf:"varint,11,opt,name=last_interim_update_attempt_time_epoch,json=lastInterimUpdateAttemptTimeEpoch,proto3" json:"last_interim_update_attempt_time_epoch,omitempty"`
	SentInterimUpdates                   uint32   `protobuf:"varint,12,opt,name=sent_interim_updates,json=sentInterimUpdates,proto3" json:"sent_interim_updates,omitempty"`
	AcceptedInterimUpdates               uint32   `protobuf:"varint,13,opt,name=accepted_interim_updates,json=acceptedInterimUpdates,proto3" json:"accepted_interim_updates,omitempty"`
	RejectedInterimUpdates               uint32   `protobuf:"varint,14,opt,name=rejected_interim_updates,json=rejectedInterimUpdates,proto3" json:"rejected_interim_updates,omitempty"`
	SentInterimUpdateFailures            uint32   `protobuf:"varint,15,opt,name=sent_interim_update_failures,json=sentInterimUpdateFailures,proto3" json:"sent_interim_update_failures,omitempty"`
	XXX_NoUnkeyedLiteral                 struct{} `json:"-"`
	XXX_unrecognized                     []byte   `json:"-"`
	XXX_sizecache                        int32    `json:"-"`
}

func (m *IedgeOperAccountingData) Reset()         { *m = IedgeOperAccountingData{} }
func (m *IedgeOperAccountingData) String() string { return proto.CompactTextString(m) }
func (*IedgeOperAccountingData) ProtoMessage()    {}
func (*IedgeOperAccountingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e0d9a3f80b04023, []int{1}
}

func (m *IedgeOperAccountingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IedgeOperAccountingData.Unmarshal(m, b)
}
func (m *IedgeOperAccountingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IedgeOperAccountingData.Marshal(b, m, deterministic)
}
func (m *IedgeOperAccountingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IedgeOperAccountingData.Merge(m, src)
}
func (m *IedgeOperAccountingData) XXX_Size() int {
	return xxx_messageInfo_IedgeOperAccountingData.Size(m)
}
func (m *IedgeOperAccountingData) XXX_DiscardUnknown() {
	xxx_messageInfo_IedgeOperAccountingData.DiscardUnknown(m)
}

var xxx_messageInfo_IedgeOperAccountingData proto.InternalMessageInfo

func (m *IedgeOperAccountingData) GetAccountingStateRc() uint32 {
	if m != nil {
		return m.AccountingStateRc
	}
	return 0
}

func (m *IedgeOperAccountingData) GetAccountingStopState() uint32 {
	if m != nil {
		return m.AccountingStopState
	}
	return 0
}

func (m *IedgeOperAccountingData) GetRecordContextName() string {
	if m != nil {
		return m.RecordContextName
	}
	return ""
}

func (m *IedgeOperAccountingData) GetMethodListName() string {
	if m != nil {
		return m.MethodListName
	}
	return ""
}

func (m *IedgeOperAccountingData) GetAccountSessionId() string {
	if m != nil {
		return m.AccountSessionId
	}
	return ""
}

func (m *IedgeOperAccountingData) GetAccountingStartTimeEpoch() uint64 {
	if m != nil {
		return m.AccountingStartTimeEpoch
	}
	return 0
}

func (m *IedgeOperAccountingData) GetIsInterimAccountingEnabled() bool {
	if m != nil {
		return m.IsInterimAccountingEnabled
	}
	return false
}

func (m *IedgeOperAccountingData) GetInterimInterval() uint32 {
	if m != nil {
		return m.InterimInterval
	}
	return 0
}

func (m *IedgeOperAccountingData) GetLastSuccessfulInterimUpdateTimeEpoch() uint64 {
	if m != nil {
		return m.LastSuccessfulInterimUpdateTimeEpoch
	}
	return 0
}

func (m *IedgeOperAccountingData) GetNextInterimUpdateAttemptTime() uint32 {
	if m != nil {
		return m.NextInterimUpdateAttemptTime
	}
	return 0
}

func (m *IedgeOperAccountingData) GetLastInterimUpdateAttemptTimeEpoch() uint64 {
	if m != nil {
		return m.LastInterimUpdateAttemptTimeEpoch
	}
	return 0
}

func (m *IedgeOperAccountingData) GetSentInterimUpdates() uint32 {
	if m != nil {
		return m.SentInterimUpdates
	}
	return 0
}

func (m *IedgeOperAccountingData) GetAcceptedInterimUpdates() uint32 {
	if m != nil {
		return m.AcceptedInterimUpdates
	}
	return 0
}

func (m *IedgeOperAccountingData) GetRejectedInterimUpdates() uint32 {
	if m != nil {
		return m.RejectedInterimUpdates
	}
	return 0
}

func (m *IedgeOperAccountingData) GetSentInterimUpdateFailures() uint32 {
	if m != nil {
		return m.SentInterimUpdateFailures
	}
	return 0
}

type IedgeAccntData struct {
	AccountingSession    []*IedgeOperAccountingData `protobuf:"bytes,1,rep,name=accounting_session,json=accountingSession,proto3" json:"accounting_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *IedgeAccntData) Reset()         { *m = IedgeAccntData{} }
func (m *IedgeAccntData) String() string { return proto.CompactTextString(m) }
func (*IedgeAccntData) ProtoMessage()    {}
func (*IedgeAccntData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e0d9a3f80b04023, []int{2}
}

func (m *IedgeAccntData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IedgeAccntData.Unmarshal(m, b)
}
func (m *IedgeAccntData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IedgeAccntData.Marshal(b, m, deterministic)
}
func (m *IedgeAccntData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IedgeAccntData.Merge(m, src)
}
func (m *IedgeAccntData) XXX_Size() int {
	return xxx_messageInfo_IedgeAccntData.Size(m)
}
func (m *IedgeAccntData) XXX_DiscardUnknown() {
	xxx_messageInfo_IedgeAccntData.DiscardUnknown(m)
}

var xxx_messageInfo_IedgeAccntData proto.InternalMessageInfo

func (m *IedgeAccntData) GetAccountingSession() []*IedgeOperAccountingData {
	if m != nil {
		return m.AccountingSession
	}
	return nil
}

type IedgePolicyData struct {
	PolicyEpoch          string   `protobuf:"bytes,1,opt,name=policy_epoch,json=policyEpoch,proto3" json:"policy_epoch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IedgePolicyData) Reset()         { *m = IedgePolicyData{} }
func (m *IedgePolicyData) String() string { return proto.CompactTextString(m) }
func (*IedgePolicyData) ProtoMessage()    {}
func (*IedgePolicyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e0d9a3f80b04023, []int{3}
}

func (m *IedgePolicyData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IedgePolicyData.Unmarshal(m, b)
}
func (m *IedgePolicyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IedgePolicyData.Marshal(b, m, deterministic)
}
func (m *IedgePolicyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IedgePolicyData.Merge(m, src)
}
func (m *IedgePolicyData) XXX_Size() int {
	return xxx_messageInfo_IedgePolicyData.Size(m)
}
func (m *IedgePolicyData) XXX_DiscardUnknown() {
	xxx_messageInfo_IedgePolicyData.DiscardUnknown(m)
}

var xxx_messageInfo_IedgePolicyData proto.InternalMessageInfo

func (m *IedgePolicyData) GetPolicyEpoch() string {
	if m != nil {
		return m.PolicyEpoch
	}
	return ""
}

type IedgeOperServiceData struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	ServiceParams        string   `protobuf:"bytes,2,opt,name=service_params,json=serviceParams,proto3" json:"service_params,omitempty"`
	ServiceType          string   `protobuf:"bytes,3,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceStatus        string   `protobuf:"bytes,4,opt,name=service_status,json=serviceStatus,proto3" json:"service_status,omitempty"`
	ServiceId            uint32   `protobuf:"varint,5,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ServicePrepaid       bool     `protobuf:"varint,6,opt,name=service_prepaid,json=servicePrepaid,proto3" json:"service_prepaid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IedgeOperServiceData) Reset()         { *m = IedgeOperServiceData{} }
func (m *IedgeOperServiceData) String() string { return proto.CompactTextString(m) }
func (*IedgeOperServiceData) ProtoMessage()    {}
func (*IedgeOperServiceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e0d9a3f80b04023, []int{4}
}

func (m *IedgeOperServiceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IedgeOperServiceData.Unmarshal(m, b)
}
func (m *IedgeOperServiceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IedgeOperServiceData.Marshal(b, m, deterministic)
}
func (m *IedgeOperServiceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IedgeOperServiceData.Merge(m, src)
}
func (m *IedgeOperServiceData) XXX_Size() int {
	return xxx_messageInfo_IedgeOperServiceData.Size(m)
}
func (m *IedgeOperServiceData) XXX_DiscardUnknown() {
	xxx_messageInfo_IedgeOperServiceData.DiscardUnknown(m)
}

var xxx_messageInfo_IedgeOperServiceData proto.InternalMessageInfo

func (m *IedgeOperServiceData) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *IedgeOperServiceData) GetServiceParams() string {
	if m != nil {
		return m.ServiceParams
	}
	return ""
}

func (m *IedgeOperServiceData) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *IedgeOperServiceData) GetServiceStatus() string {
	if m != nil {
		return m.ServiceStatus
	}
	return ""
}

func (m *IedgeOperServiceData) GetServiceId() uint32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *IedgeOperServiceData) GetServicePrepaid() bool {
	if m != nil {
		return m.ServicePrepaid
	}
	return false
}

type IedgeOperCoaData struct {
	RequestAcked         bool     `protobuf:"varint,1,opt,name=request_acked,json=requestAcked,proto3" json:"request_acked,omitempty"`
	RequestTimeEpoch     uint64   `protobuf:"varint,2,opt,name=request_time_epoch,json=requestTimeEpoch,proto3" json:"request_time_epoch,omitempty"`
	ReplyTimeEpoch       uint64   `protobuf:"varint,3,opt,name=reply_time_epoch,json=replyTimeEpoch,proto3" json:"reply_time_epoch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IedgeOperCoaData) Reset()         { *m = IedgeOperCoaData{} }
func (m *IedgeOperCoaData) String() string { return proto.CompactTextString(m) }
func (*IedgeOperCoaData) ProtoMessage()    {}
func (*IedgeOperCoaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e0d9a3f80b04023, []int{5}
}

func (m *IedgeOperCoaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IedgeOperCoaData.Unmarshal(m, b)
}
func (m *IedgeOperCoaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IedgeOperCoaData.Marshal(b, m, deterministic)
}
func (m *IedgeOperCoaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IedgeOperCoaData.Merge(m, src)
}
func (m *IedgeOperCoaData) XXX_Size() int {
	return xxx_messageInfo_IedgeOperCoaData.Size(m)
}
func (m *IedgeOperCoaData) XXX_DiscardUnknown() {
	xxx_messageInfo_IedgeOperCoaData.DiscardUnknown(m)
}

var xxx_messageInfo_IedgeOperCoaData proto.InternalMessageInfo

func (m *IedgeOperCoaData) GetRequestAcked() bool {
	if m != nil {
		return m.RequestAcked
	}
	return false
}

func (m *IedgeOperCoaData) GetRequestTimeEpoch() uint64 {
	if m != nil {
		return m.RequestTimeEpoch
	}
	return 0
}

func (m *IedgeOperCoaData) GetReplyTimeEpoch() uint64 {
	if m != nil {
		return m.ReplyTimeEpoch
	}
	return 0
}

type SubscriberSessionData struct {
	SessionType                  string                  `protobuf:"bytes,50,opt,name=session_type,json=sessionType,proto3" json:"session_type,omitempty"`
	PppoeSubType                 string                  `protobuf:"bytes,51,opt,name=pppoe_sub_type,json=pppoeSubType,proto3" json:"pppoe_sub_type,omitempty"`
	InterfaceName                string                  `protobuf:"bytes,52,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	VrfName                      string                  `protobuf:"bytes,53,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	CircuitId                    string                  `protobuf:"bytes,54,opt,name=circuit_id,json=circuitId,proto3" json:"circuit_id,omitempty"`
	RemoteId                     string                  `protobuf:"bytes,55,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
	LnsAddress                   string                  `protobuf:"bytes,56,opt,name=lns_address,json=lnsAddress,proto3" json:"lns_address,omitempty"`
	LacAddress                   string                  `protobuf:"bytes,57,opt,name=lac_address,json=lacAddress,proto3" json:"lac_address,omitempty"`
	TunnelClientAuthenticationId string                  `protobuf:"bytes,58,opt,name=tunnel_client_authentication_id,json=tunnelClientAuthenticationId,proto3" json:"tunnel_client_authentication_id,omitempty"`
	TunnelServerAuthenticationId string                  `protobuf:"bytes,59,opt,name=tunnel_server_authentication_id,json=tunnelServerAuthenticationId,proto3" json:"tunnel_server_authentication_id,omitempty"`
	SessionIpAddress             string                  `protobuf:"bytes,60,opt,name=session_ip_address,json=sessionIpAddress,proto3" json:"session_ip_address,omitempty"`
	SessionIpv6Address           string                  `protobuf:"bytes,61,opt,name=session_ipv6_address,json=sessionIpv6Address,proto3" json:"session_ipv6_address,omitempty"`
	SessionIpv6Prefix            string                  `protobuf:"bytes,62,opt,name=session_ipv6_prefix,json=sessionIpv6Prefix,proto3" json:"session_ipv6_prefix,omitempty"`
	DelegatedIpv6Prefix          string                  `protobuf:"bytes,63,opt,name=delegated_ipv6_prefix,json=delegatedIpv6Prefix,proto3" json:"delegated_ipv6_prefix,omitempty"`
	Ipv6InterfaceId              string                  `protobuf:"bytes,64,opt,name=ipv6_interface_id,json=ipv6InterfaceId,proto3" json:"ipv6_interface_id,omitempty"`
	MacAddress                   string                  `protobuf:"bytes,65,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	AccountSessionId             string                  `protobuf:"bytes,66,opt,name=account_session_id,json=accountSessionId,proto3" json:"account_session_id,omitempty"`
	NasPort                      string                  `protobuf:"bytes,67,opt,name=nas_port,json=nasPort,proto3" json:"nas_port,omitempty"`
	Username                     string                  `protobuf:"bytes,68,opt,name=username,proto3" json:"username,omitempty"`
	Clientname                   string                  `protobuf:"bytes,69,opt,name=clientname,proto3" json:"clientname,omitempty"`
	Formattedname                string                  `protobuf:"bytes,70,opt,name=formattedname,proto3" json:"formattedname,omitempty"`
	IsSessionAuthentic           bool                    `protobuf:"varint,71,opt,name=is_session_authentic,json=isSessionAuthentic,proto3" json:"is_session_authentic,omitempty"`
	IsSessionAuthor              bool                    `protobuf:"varint,72,opt,name=is_session_author,json=isSessionAuthor,proto3" json:"is_session_author,omitempty"`
	SessionState                 string                  `protobuf:"bytes,73,opt,name=session_state,json=sessionState,proto3" json:"session_state,omitempty"`
	SessionCreationTimeEpoch     uint64                  `protobuf:"varint,74,opt,name=session_creation_time_epoch,json=sessionCreationTimeEpoch,proto3" json:"session_creation_time_epoch,omitempty"`
	IdleStateChangeTime          string                  `protobuf:"bytes,75,opt,name=idle_state_change_time,json=idleStateChangeTime,proto3" json:"idle_state_change_time,omitempty"`
	TotalSessionIdleTime         uint32                  `protobuf:"varint,76,opt,name=total_session_idle_time,json=totalSessionIdleTime,proto3" json:"total_session_idle_time,omitempty"`
	AccessInterfaceName          string                  `protobuf:"bytes,77,opt,name=access_interface_name,json=accessInterfaceName,proto3" json:"access_interface_name,omitempty"`
	Accounting                   *IedgeAccntData         `protobuf:"bytes,78,opt,name=accounting,proto3" json:"accounting,omitempty"`
	SubPolicyData                []*IedgePolicyData      `protobuf:"bytes,79,rep,name=sub_policy_data,json=subPolicyData,proto3" json:"sub_policy_data,omitempty"`
	SessionServiceInfo           []*IedgeOperServiceData `protobuf:"bytes,80,rep,name=session_service_info,json=sessionServiceInfo,proto3" json:"session_service_info,omitempty"`
	SessionChangeOfAuthorization []*IedgeOperCoaData     `protobuf:"bytes,81,rep,name=session_change_of_authorization,json=sessionChangeOfAuthorization,proto3" json:"session_change_of_authorization,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                `json:"-"`
	XXX_unrecognized             []byte                  `json:"-"`
	XXX_sizecache                int32                   `json:"-"`
}

func (m *SubscriberSessionData) Reset()         { *m = SubscriberSessionData{} }
func (m *SubscriberSessionData) String() string { return proto.CompactTextString(m) }
func (*SubscriberSessionData) ProtoMessage()    {}
func (*SubscriberSessionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e0d9a3f80b04023, []int{6}
}

func (m *SubscriberSessionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberSessionData.Unmarshal(m, b)
}
func (m *SubscriberSessionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberSessionData.Marshal(b, m, deterministic)
}
func (m *SubscriberSessionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberSessionData.Merge(m, src)
}
func (m *SubscriberSessionData) XXX_Size() int {
	return xxx_messageInfo_SubscriberSessionData.Size(m)
}
func (m *SubscriberSessionData) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberSessionData.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberSessionData proto.InternalMessageInfo

func (m *SubscriberSessionData) GetSessionType() string {
	if m != nil {
		return m.SessionType
	}
	return ""
}

func (m *SubscriberSessionData) GetPppoeSubType() string {
	if m != nil {
		return m.PppoeSubType
	}
	return ""
}

func (m *SubscriberSessionData) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *SubscriberSessionData) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *SubscriberSessionData) GetCircuitId() string {
	if m != nil {
		return m.CircuitId
	}
	return ""
}

func (m *SubscriberSessionData) GetRemoteId() string {
	if m != nil {
		return m.RemoteId
	}
	return ""
}

func (m *SubscriberSessionData) GetLnsAddress() string {
	if m != nil {
		return m.LnsAddress
	}
	return ""
}

func (m *SubscriberSessionData) GetLacAddress() string {
	if m != nil {
		return m.LacAddress
	}
	return ""
}

func (m *SubscriberSessionData) GetTunnelClientAuthenticationId() string {
	if m != nil {
		return m.TunnelClientAuthenticationId
	}
	return ""
}

func (m *SubscriberSessionData) GetTunnelServerAuthenticationId() string {
	if m != nil {
		return m.TunnelServerAuthenticationId
	}
	return ""
}

func (m *SubscriberSessionData) GetSessionIpAddress() string {
	if m != nil {
		return m.SessionIpAddress
	}
	return ""
}

func (m *SubscriberSessionData) GetSessionIpv6Address() string {
	if m != nil {
		return m.SessionIpv6Address
	}
	return ""
}

func (m *SubscriberSessionData) GetSessionIpv6Prefix() string {
	if m != nil {
		return m.SessionIpv6Prefix
	}
	return ""
}

func (m *SubscriberSessionData) GetDelegatedIpv6Prefix() string {
	if m != nil {
		return m.DelegatedIpv6Prefix
	}
	return ""
}

func (m *SubscriberSessionData) GetIpv6InterfaceId() string {
	if m != nil {
		return m.Ipv6InterfaceId
	}
	return ""
}

func (m *SubscriberSessionData) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *SubscriberSessionData) GetAccountSessionId() string {
	if m != nil {
		return m.AccountSessionId
	}
	return ""
}

func (m *SubscriberSessionData) GetNasPort() string {
	if m != nil {
		return m.NasPort
	}
	return ""
}

func (m *SubscriberSessionData) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SubscriberSessionData) GetClientname() string {
	if m != nil {
		return m.Clientname
	}
	return ""
}

func (m *SubscriberSessionData) GetFormattedname() string {
	if m != nil {
		return m.Formattedname
	}
	return ""
}

func (m *SubscriberSessionData) GetIsSessionAuthentic() bool {
	if m != nil {
		return m.IsSessionAuthentic
	}
	return false
}

func (m *SubscriberSessionData) GetIsSessionAuthor() bool {
	if m != nil {
		return m.IsSessionAuthor
	}
	return false
}

func (m *SubscriberSessionData) GetSessionState() string {
	if m != nil {
		return m.SessionState
	}
	return ""
}

func (m *SubscriberSessionData) GetSessionCreationTimeEpoch() uint64 {
	if m != nil {
		return m.SessionCreationTimeEpoch
	}
	return 0
}

func (m *SubscriberSessionData) GetIdleStateChangeTime() string {
	if m != nil {
		return m.IdleStateChangeTime
	}
	return ""
}

func (m *SubscriberSessionData) GetTotalSessionIdleTime() uint32 {
	if m != nil {
		return m.TotalSessionIdleTime
	}
	return 0
}

func (m *SubscriberSessionData) GetAccessInterfaceName() string {
	if m != nil {
		return m.AccessInterfaceName
	}
	return ""
}

func (m *SubscriberSessionData) GetAccounting() *IedgeAccntData {
	if m != nil {
		return m.Accounting
	}
	return nil
}

func (m *SubscriberSessionData) GetSubPolicyData() []*IedgePolicyData {
	if m != nil {
		return m.SubPolicyData
	}
	return nil
}

func (m *SubscriberSessionData) GetSessionServiceInfo() []*IedgeOperServiceData {
	if m != nil {
		return m.SessionServiceInfo
	}
	return nil
}

func (m *SubscriberSessionData) GetSessionChangeOfAuthorization() []*IedgeOperCoaData {
	if m != nil {
		return m.SessionChangeOfAuthorization
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscriberSessionData_KEYS)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.sessions.session.subscriber_session_data_KEYS")
	proto.RegisterType((*IedgeOperAccountingData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.sessions.session.iedge_oper_accounting_data")
	proto.RegisterType((*IedgeAccntData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.sessions.session.iedge_accnt_data")
	proto.RegisterType((*IedgePolicyData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.sessions.session.iedge_policy_data")
	proto.RegisterType((*IedgeOperServiceData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.sessions.session.iedge_oper_service_data")
	proto.RegisterType((*IedgeOperCoaData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.sessions.session.iedge_oper_coa_data")
	proto.RegisterType((*SubscriberSessionData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.sessions.session.subscriber_session_data")
}

func init() { proto.RegisterFile("subscriber_session_data.proto", fileDescriptor_0e0d9a3f80b04023) }

var fileDescriptor_0e0d9a3f80b04023 = []byte{
	// 1333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xdf, 0x73, 0x14, 0x45,
	0x10, 0xae, 0x03, 0x84, 0x4b, 0x27, 0x97, 0x1f, 0x1b, 0x84, 0x05, 0x82, 0x84, 0x88, 0x7a, 0x58,
	0xd4, 0x15, 0x06, 0x08, 0xf8, 0x03, 0xf1, 0x0c, 0x41, 0x0f, 0x10, 0xc2, 0x1d, 0x96, 0xa5, 0xa5,
	0xb5, 0x35, 0x37, 0x3b, 0x97, 0x8c, 0xee, 0xed, 0xac, 0x33, 0xb3, 0x29, 0x62, 0xf9, 0xea, 0xbb,
	0x65, 0x15, 0xff, 0x82, 0x0f, 0xfe, 0x7f, 0xbe, 0x5b, 0xd3, 0x3d, 0xb3, 0xb7, 0x97, 0x10, 0x9f,
	0xe0, 0x25, 0xc9, 0x76, 0x7f, 0x5f, 0xcf, 0x37, 0xd3, 0x3d, 0xdd, 0x13, 0xb8, 0x68, 0xca, 0xa1,
	0xe1, 0x5a, 0x0e, 0x85, 0x4e, 0x8c, 0x30, 0x46, 0xaa, 0x3c, 0x49, 0x99, 0x65, 0x9d, 0x42, 0x2b,
	0xab, 0xa2, 0x87, 0x5c, 0x1a, 0xae, 0x12, 0xa9, 0x4c, 0xf2, 0x42, 0x27, 0x52, 0xa4, 0x3b, 0xe2,
	0xe6, 0xed, 0x8f, 0xae, 0x27, 0xaa, 0x10, 0xba, 0x33, 0xa1, 0x76, 0x3c, 0xb5, 0x93, 0xab, 0x54,
	0x18, 0xfc, 0x19, 0x4c, 0x26, 0xfc, 0xb1, 0xf6, 0x03, 0xac, 0x1c, 0xb1, 0x58, 0xf2, 0x68, 0xeb,
	0xfb, 0x41, 0x74, 0x01, 0x66, 0x1c, 0x31, 0xc9, 0xd9, 0x58, 0xc4, 0x8d, 0xd5, 0x46, 0x7b, 0xa6,
	0xdf, 0x74, 0x86, 0x27, 0x6c, 0x2c, 0xa2, 0x8b, 0x00, 0x81, 0x21, 0xd3, 0xf8, 0x18, 0x7a, 0x67,
	0xbc, 0xa5, 0x97, 0xae, 0xfd, 0x79, 0x0a, 0xce, 0xa3, 0x3a, 0x54, 0x96, 0x30, 0xce, 0x55, 0x99,
	0x5b, 0x99, 0xef, 0x60, 0xfc, 0xa8, 0x03, 0xcb, 0x35, 0x93, 0xb1, 0xcc, 0x8a, 0x44, 0x73, 0x5c,
	0xa4, 0xd5, 0x5f, 0x9a, 0xb8, 0x06, 0xce, 0xd3, 0xe7, 0xd1, 0x3a, 0xbc, 0x3d, 0x85, 0x57, 0x05,
	0x91, 0x70, 0xe1, 0x56, 0x7f, 0xb9, 0xce, 0x50, 0x05, 0xb2, 0xdc, 0x1a, 0x5a, 0x70, 0xa5, 0xd3,
	0x84, 0xab, 0xdc, 0x8a, 0x17, 0x96, 0x36, 0x72, 0x1c, 0xa5, 0x2e, 0x91, 0x6b, 0x93, 0x3c, 0xb8,
	0xa3, 0x36, 0x2c, 0x8e, 0x85, 0xdd, 0x55, 0x69, 0x92, 0x49, 0xe3, 0xc1, 0x27, 0x10, 0x3c, 0x4f,
	0xf6, 0xc7, 0xd2, 0x10, 0xf2, 0x1a, 0x44, 0x7e, 0xc1, 0xa4, 0x76, 0x06, 0x6f, 0x21, 0x76, 0xd1,
	0x7b, 0x06, 0xe1, 0x28, 0xa2, 0xbb, 0x70, 0x61, 0x7a, 0xaf, 0xda, 0x26, 0x56, 0x8e, 0x45, 0x22,
	0x0a, 0xc5, 0x77, 0xe3, 0x93, 0xab, 0x8d, 0xf6, 0x89, 0x7e, 0x3c, 0xb5, 0x67, 0x6d, 0x9f, 0xcb,
	0xb1, 0xd8, 0x72, 0xfe, 0xa8, 0x0b, 0x17, 0xa5, 0x49, 0x64, 0x6e, 0x85, 0x96, 0xe3, 0xfa, 0x41,
	0x8a, 0x9c, 0x0d, 0x33, 0x91, 0xc6, 0xa7, 0x56, 0x1b, 0xed, 0x66, 0xff, 0xbc, 0x34, 0x3d, 0xc2,
	0x74, 0x2b, 0xc8, 0x16, 0x21, 0xa2, 0xab, 0xb0, 0x18, 0xf8, 0xf8, 0x7b, 0x8f, 0x65, 0x71, 0x13,
	0x0f, 0x6e, 0xc1, 0xdb, 0x7b, 0xde, 0x1c, 0x7d, 0x07, 0x57, 0x33, 0x66, 0x6c, 0x62, 0x4a, 0xce,
	0x85, 0x31, 0xa3, 0x32, 0xab, 0x96, 0x2e, 0x8b, 0xd4, 0xa5, 0xa9, 0x26, 0x7d, 0x06, 0xa5, 0x5f,
	0x71, 0x84, 0x41, 0x85, 0xf7, 0x2a, 0xbe, 0x45, 0xf4, 0x64, 0x1b, 0x0f, 0x60, 0x35, 0x77, 0x39,
	0x38, 0x10, 0x8d, 0x59, 0x2b, 0xc6, 0x05, 0x1d, 0x48, 0x0c, 0xa8, 0x69, 0xc5, 0xe1, 0xa6, 0xa2,
	0x74, 0x09, 0xe4, 0x82, 0x45, 0xcf, 0xe0, 0x7d, 0x14, 0xf8, 0x3f, 0x71, 0xbc, 0xba, 0x59, 0x54,
	0x77, 0xd9, 0xa1, 0x8f, 0x8a, 0x46, 0xd2, 0xae, 0xc3, 0x69, 0x23, 0xf2, 0x83, 0x21, 0x4d, 0x3c,
	0x87, 0x72, 0x22, 0xe7, 0x9b, 0x0a, 0x60, 0xa2, 0x3b, 0xe0, 0xf2, 0x25, 0x0a, 0x2b, 0xd2, 0x43,
	0xac, 0x16, 0xb2, 0xce, 0x04, 0xff, 0x61, 0xa6, 0x16, 0x3f, 0x0b, 0xfe, 0x2a, 0xe6, 0x3c, 0x31,
	0x83, 0xff, 0x00, 0xf3, 0x1e, 0xac, 0xbc, 0x42, 0x65, 0x32, 0x62, 0x32, 0x2b, 0xb5, 0x30, 0xf1,
	0x02, 0xb2, 0xcf, 0x1d, 0x52, 0xfb, 0xc0, 0x03, 0xd6, 0xfe, 0x69, 0xc0, 0x22, 0x5d, 0x49, 0xc6,
	0x79, 0x6e, 0xe9, 0x22, 0xbe, 0x6c, 0x54, 0xb5, 0x8c, 0xd5, 0x49, 0x55, 0x1b, 0x37, 0x56, 0x8f,
	0xb7, 0x67, 0xd7, 0x47, 0x9d, 0xd7, 0xd7, 0x6d, 0x3a, 0x47, 0x77, 0x83, 0xa9, 0x0b, 0xef, 0x7b,
	0xd3, 0x06, 0x2c, 0x11, 0xa1, 0x50, 0x99, 0xe4, 0xfb, 0x24, 0xf6, 0x32, 0xcc, 0xf9, 0x4f, 0xca,
	0x30, 0xf5, 0xa4, 0x59, 0xb2, 0x61, 0x2e, 0xd7, 0xfe, 0x6d, 0xc0, 0xd9, 0xda, 0x4a, 0x46, 0xe8,
	0x3d, 0xc9, 0x45, 0x45, 0x0f, 0xdf, 0xb5, 0x96, 0x36, 0xeb, 0x6d, 0x78, 0xb3, 0xdf, 0x83, 0xf9,
	0x00, 0x29, 0x98, 0x66, 0x63, 0xe3, 0x3b, 0x5b, 0xcb, 0x5b, 0xb7, 0xd1, 0x58, 0x8f, 0x64, 0xf7,
	0x8b, 0xd0, 0x53, 0x42, 0xa4, 0xe7, 0xfb, 0xc5, 0x54, 0x24, 0xd7, 0xa9, 0x4a, 0xe3, 0x7b, 0x49,
	0x88, 0x34, 0x40, 0x23, 0xb5, 0x51, 0x82, 0xf9, 0x16, 0xd2, 0x72, 0x6d, 0x14, 0x2d, 0xbd, 0x34,
	0xfa, 0x00, 0x16, 0x2a, 0x3d, 0x5a, 0x14, 0x4c, 0xa6, 0xd8, 0x2f, 0x9a, 0xfd, 0x10, 0x7c, 0x9b,
	0xac, 0x6b, 0x7f, 0x35, 0x60, 0xb9, 0xb6, 0x6f, 0xae, 0x18, 0xed, 0xf9, 0x5d, 0x68, 0x69, 0xf1,
	0x6b, 0x29, 0x8c, 0x4d, 0x18, 0xff, 0x45, 0xa4, 0xb8, 0xe9, 0x66, 0x7f, 0xce, 0x1b, 0xbb, 0xce,
	0xe6, 0xfa, 0x59, 0x00, 0xd5, 0xee, 0xcf, 0x31, 0xbc, 0x3f, 0x8b, 0xde, 0x33, 0xb9, 0x2e, 0x6d,
	0x58, 0xd4, 0xa2, 0xc8, 0xf6, 0xeb, 0xd8, 0xe3, 0x88, 0x9d, 0x47, 0x7b, 0x85, 0x5c, 0x7b, 0x39,
	0x0f, 0x67, 0x8f, 0x98, 0x30, 0x74, 0x84, 0xf4, 0x8d, 0x47, 0xb8, 0x1e, 0x8e, 0x10, 0x6d, 0x78,
	0x84, 0x57, 0x60, 0xbe, 0x28, 0x0a, 0x25, 0x12, 0x53, 0x0e, 0x09, 0x74, 0x03, 0x41, 0x73, 0x68,
	0x1d, 0x94, 0xc3, 0x70, 0xd0, 0x78, 0x25, 0x46, 0x2c, 0xe4, 0xf5, 0x26, 0x1d, 0x74, 0x65, 0xc5,
	0xcc, 0x9e, 0x83, 0xe6, 0x9e, 0x1e, 0x11, 0xe0, 0x16, 0x02, 0x4e, 0xed, 0xe9, 0x51, 0x18, 0x65,
	0x5c, 0x6a, 0x5e, 0x4a, 0xeb, 0x72, 0xb0, 0x41, 0xa3, 0xcc, 0x5b, 0x7a, 0xa9, 0x1b, 0x83, 0x5a,
	0x8c, 0x95, 0xc5, 0x0c, 0xdd, 0xa6, 0x31, 0x48, 0x86, 0x5e, 0x1a, 0x5d, 0x82, 0xd9, 0x2c, 0x37,
	0x09, 0x4b, 0x53, 0x2d, 0x8c, 0x89, 0xef, 0xa0, 0x1b, 0xb2, 0xdc, 0x74, 0xc9, 0x82, 0x00, 0xc6,
	0x2b, 0xc0, 0xc7, 0x1e, 0xc0, 0x78, 0x00, 0x6c, 0xc1, 0x25, 0x5b, 0xe6, 0xb9, 0xc8, 0x12, 0x9e,
	0x49, 0x77, 0xc1, 0x59, 0x69, 0x77, 0x45, 0x6e, 0x25, 0x67, 0xd6, 0x4f, 0x96, 0x4f, 0x90, 0xb4,
	0x42, 0xb0, 0x4d, 0x44, 0x75, 0xa7, 0x40, 0xbd, 0xb4, 0x16, 0xc6, 0x55, 0x86, 0xbb, 0x64, 0x87,
	0xc2, 0x7c, 0x5a, 0x0f, 0x33, 0x40, 0xd4, 0xa1, 0x30, 0xd7, 0x20, 0xaa, 0x46, 0x5a, 0x51, 0xa9,
	0xfe, 0x8c, 0x46, 0x5b, 0x18, 0xef, 0x45, 0xd0, 0x8e, 0x9d, 0x33, 0xa0, 0xf7, 0x36, 0x2a, 0xfc,
	0x5d, 0xc4, 0x47, 0x15, 0x7e, 0x6f, 0x23, 0x30, 0x3a, 0xb0, 0x3c, 0xc5, 0x28, 0xb4, 0x18, 0xc9,
	0x17, 0xf1, 0xe7, 0x34, 0x94, 0x6b, 0x84, 0x6d, 0x74, 0xb8, 0xc1, 0x9f, 0x8a, 0x4c, 0xec, 0x30,
	0x6c, 0x98, 0x35, 0xc6, 0x3d, 0x64, 0x2c, 0x57, 0xce, 0x1a, 0xe7, 0x43, 0x58, 0x42, 0xe4, 0xa4,
	0x2c, 0x64, 0x1a, 0x7f, 0x81, 0xf8, 0x05, 0xe7, 0xe8, 0x05, 0x3b, 0xe5, 0x6f, 0x5c, 0x4b, 0x4f,
	0x97, 0xd2, 0x33, 0x9e, 0xa4, 0xe7, 0xd5, 0xb3, 0xfe, 0xcb, 0x23, 0x66, 0xfd, 0x39, 0x68, 0xe6,
	0xcc, 0x24, 0x85, 0xd2, 0x36, 0xde, 0xa4, 0x2a, 0xcb, 0x99, 0xd9, 0x56, 0xda, 0x46, 0xe7, 0xa1,
	0x59, 0x1a, 0xa1, 0xb1, 0x00, 0xef, 0x53, 0x15, 0x85, 0xef, 0xe8, 0x1d, 0x00, 0x4a, 0x3e, 0x7a,
	0xb7, 0x48, 0xc4, 0xc4, 0x12, 0x5d, 0x81, 0xd6, 0x48, 0xe9, 0xb1, 0x1b, 0x72, 0x29, 0x42, 0x1e,
	0x50, 0x89, 0x4f, 0x19, 0x5d, 0x36, 0xa4, 0xa9, 0x54, 0x56, 0xf9, 0x8f, 0xbf, 0xc2, 0x2b, 0x1f,
	0x49, 0xe3, 0x75, 0x56, 0x49, 0xc7, 0x93, 0x9a, 0x66, 0x28, 0x1d, 0x7f, 0x8d, 0xf0, 0x85, 0x29,
	0xb8, 0xd2, 0xae, 0x93, 0x04, 0x20, 0x3d, 0xbd, 0x7a, 0x74, 0x19, 0xbd, 0x91, 0xde, 0x5c, 0x77,
	0xe1, 0x42, 0x00, 0x71, 0x2d, 0xa8, 0xf2, 0x6a, 0x6d, 0xe2, 0x21, 0xbd, 0x75, 0x3c, 0x64, 0xd3,
	0x23, 0x26, 0xad, 0xe5, 0x06, 0x9c, 0x91, 0x69, 0x26, 0xfc, 0x83, 0x90, 0xef, 0xb2, 0x7c, 0x87,
	0x1e, 0x1c, 0xf1, 0x23, 0x4a, 0xb7, 0xf3, 0xe2, 0x4a, 0x9b, 0xe8, 0xc3, 0x17, 0xc1, 0x2d, 0x38,
	0x6b, 0x95, 0x65, 0x59, 0x2d, 0x3f, 0x99, 0x67, 0x3d, 0xc6, 0x7e, 0x7a, 0x1a, 0xdd, 0x55, 0x92,
	0x32, 0xa2, 0xd1, 0x93, 0x52, 0x18, 0x93, 0x1c, 0x68, 0x1f, 0xdf, 0xd0, 0x52, 0xe4, 0xec, 0x4d,
	0x35, 0x91, 0xdf, 0x01, 0x26, 0xa3, 0x2a, 0x7e, 0xb2, 0xda, 0x68, 0xcf, 0xae, 0xff, 0xf8, 0xfa,
	0x87, 0xe4, 0x64, 0x3e, 0xf7, 0x6b, 0xeb, 0x45, 0x7f, 0x34, 0x60, 0xc1, 0xb5, 0xc2, 0xda, 0x48,
	0x8c, 0x9f, 0xe2, 0xa0, 0xfe, 0xe9, 0xf5, 0x6b, 0xa8, 0x2d, 0xd2, 0x6f, 0x99, 0x72, 0xb8, 0x8d,
	0xdf, 0xf7, 0xfd, 0x9b, 0xa1, 0xba, 0xf6, 0xd5, 0xf0, 0xca, 0x47, 0x2a, 0xde, 0x46, 0x31, 0xfc,
	0x0d, 0xbd, 0x1a, 0xea, 0xb3, 0xbc, 0xea, 0x2d, 0x03, 0x3f, 0x2a, 0xf3, 0x91, 0x8a, 0xfe, 0x6e,
	0xc0, 0xa5, 0xaa, 0xfa, 0xa8, 0x76, 0xd4, 0xc8, 0x57, 0xb5, 0xfc, 0x0d, 0x4b, 0x2d, 0x7e, 0x86,
	0x12, 0x93, 0x37, 0x24, 0x31, 0x8c, 0xdd, 0xfe, 0x4a, 0x28, 0x71, 0x94, 0xf1, 0x74, 0xd4, 0xad,
	0x8b, 0x18, 0x9e, 0xc4, 0xff, 0xe5, 0x6e, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x78, 0xbf, 0x50,
	0x06, 0xec, 0x0d, 0x00, 0x00,
}
