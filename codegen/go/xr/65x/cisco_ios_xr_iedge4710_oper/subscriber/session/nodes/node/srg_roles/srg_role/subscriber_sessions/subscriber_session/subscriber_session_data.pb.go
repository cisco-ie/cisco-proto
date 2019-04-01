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

package cisco_ios_xr_iedge4710_oper_subscriber_session_nodes_node_srg_roles_srg_role_subscriber_sessions_subscriber_session

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
	Srg                  string   `protobuf:"bytes,2,opt,name=srg,proto3" json:"srg,omitempty"`
	SessionId            string   `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
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

func (m *SubscriberSessionData_KEYS) GetSrg() string {
	if m != nil {
		return m.Srg
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
	Ipv6InterfaceId              []byte                  `protobuf:"bytes,64,opt,name=ipv6_interface_id,json=ipv6InterfaceId,proto3" json:"ipv6_interface_id,omitempty"`
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

func (m *SubscriberSessionData) GetIpv6InterfaceId() []byte {
	if m != nil {
		return m.Ipv6InterfaceId
	}
	return nil
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
	proto.RegisterType((*SubscriberSessionData_KEYS)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.subscriber_sessions.subscriber_session.subscriber_session_data_KEYS")
	proto.RegisterType((*IedgeOperAccountingData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.subscriber_sessions.subscriber_session.iedge_oper_accounting_data")
	proto.RegisterType((*IedgeAccntData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.subscriber_sessions.subscriber_session.iedge_accnt_data")
	proto.RegisterType((*IedgePolicyData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.subscriber_sessions.subscriber_session.iedge_policy_data")
	proto.RegisterType((*IedgeOperServiceData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.subscriber_sessions.subscriber_session.iedge_oper_service_data")
	proto.RegisterType((*IedgeOperCoaData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.subscriber_sessions.subscriber_session.iedge_oper_coa_data")
	proto.RegisterType((*SubscriberSessionData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.subscriber_sessions.subscriber_session.subscriber_session_data")
}

func init() { proto.RegisterFile("subscriber_session_data.proto", fileDescriptor_0e0d9a3f80b04023) }

var fileDescriptor_0e0d9a3f80b04023 = []byte{
	// 1372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xdd, 0x72, 0x1b, 0xc5,
	0x12, 0xae, 0x8d, 0x73, 0x12, 0xb9, 0x6d, 0xf9, 0x67, 0x9d, 0x93, 0x6c, 0x12, 0xfb, 0x44, 0xf1,
	0xc9, 0x39, 0x28, 0x54, 0x4a, 0x15, 0x9c, 0xc4, 0x09, 0x3f, 0x21, 0x08, 0xc7, 0x01, 0x91, 0x90,
	0x38, 0x52, 0x28, 0x8a, 0xab, 0xad, 0xd1, 0xec, 0xc8, 0x1e, 0x58, 0xed, 0x2c, 0x33, 0xb3, 0xae,
	0x98, 0x7b, 0x28, 0x2e, 0xa8, 0x82, 0xe2, 0x21, 0xb8, 0xe5, 0x0d, 0xe0, 0x15, 0xb8, 0xe0, 0x55,
	0xb8, 0xa7, 0xa6, 0x7b, 0x76, 0xb5, 0xf2, 0x0f, 0xb7, 0xbe, 0x49, 0xb4, 0x5f, 0x7f, 0xdd, 0xd3,
	0x33, 0x3d, 0xf3, 0x75, 0x1b, 0xd6, 0x4c, 0x31, 0x34, 0x5c, 0xcb, 0xa1, 0xd0, 0xb1, 0x11, 0xc6,
	0x48, 0x95, 0xc5, 0x09, 0xb3, 0xac, 0x93, 0x6b, 0x65, 0x55, 0x68, 0xb8, 0x34, 0x5c, 0xc5, 0x52,
	0x99, 0xf8, 0xb5, 0x8e, 0xa5, 0x48, 0x76, 0xc5, 0xdd, 0xfb, 0x6f, 0xdd, 0x8e, 0x55, 0x2e, 0x74,
	0x67, 0xe2, 0xda, 0xf1, 0xae, 0x9d, 0x4c, 0x25, 0xc2, 0xe0, 0xbf, 0x1d, 0xa3, 0x77, 0x63, 0xad,
	0x52, 0x61, 0xaa, 0x5f, 0x9d, 0xa3, 0x6b, 0x99, 0x63, 0xb0, 0xf5, 0x14, 0x56, 0x4f, 0xc8, 0x2a,
	0x7e, 0xba, 0xfd, 0xc5, 0x20, 0xbc, 0x0a, 0xb3, 0x6e, 0x85, 0x38, 0x63, 0x63, 0x11, 0x05, 0xad,
	0xa0, 0x3d, 0xdb, 0x6f, 0x38, 0xe0, 0x39, 0x1b, 0x8b, 0x70, 0x09, 0x66, 0x8c, 0xde, 0x8d, 0xce,
	0x20, 0xec, 0x7e, 0x86, 0x6b, 0x00, 0x65, 0x0c, 0x99, 0x44, 0x33, 0x68, 0x98, 0xf5, 0x48, 0x2f,
	0x59, 0xff, 0xe9, 0x3c, 0x5c, 0xc1, 0x8d, 0xe1, 0xa6, 0x62, 0xc6, 0xb9, 0x2a, 0x32, 0x2b, 0xb3,
	0x5d, 0x5c, 0x31, 0xec, 0xc0, 0x4a, 0x0d, 0x32, 0x96, 0x59, 0x11, 0x6b, 0x8e, 0xcb, 0x36, 0xfb,
	0xcb, 0x13, 0xd3, 0xc0, 0x59, 0xfa, 0x3c, 0xdc, 0x80, 0x7f, 0x4f, 0xf1, 0x55, 0x4e, 0x4e, 0x98,
	0x51, 0xb3, 0xbf, 0x52, 0xf7, 0x50, 0x39, 0x7a, 0xb9, 0x35, 0xb4, 0xe0, 0x4a, 0x27, 0x31, 0x57,
	0x99, 0x15, 0xaf, 0x2d, 0x6d, 0x8d, 0x52, 0x5d, 0x26, 0xd3, 0x16, 0x59, 0x70, 0x8f, 0x6d, 0x58,
	0x1a, 0x0b, 0xbb, 0xa7, 0x92, 0x38, 0x95, 0xc6, 0x93, 0xcf, 0x22, 0x79, 0x81, 0xf0, 0x67, 0xd2,
	0x10, 0xf3, 0x16, 0x84, 0x7e, 0xc1, 0xb8, 0x76, 0x06, 0xff, 0x42, 0xee, 0x92, 0xb7, 0x0c, 0xca,
	0xa3, 0x08, 0x1f, 0xc2, 0xd5, 0xe9, 0xbd, 0x6a, 0x1b, 0x5b, 0x39, 0x16, 0xb1, 0xc8, 0x15, 0xdf,
	0x8b, 0xce, 0xb5, 0x82, 0xf6, 0xd9, 0x7e, 0x34, 0xb5, 0x67, 0x6d, 0x5f, 0xc9, 0xb1, 0xd8, 0x76,
	0xf6, 0xb0, 0x0b, 0x6b, 0xd2, 0xc4, 0x32, 0xb3, 0x42, 0xcb, 0x71, 0xfd, 0x20, 0x45, 0xc6, 0x86,
	0xa9, 0x48, 0xa2, 0xf3, 0xad, 0xa0, 0xdd, 0xe8, 0x5f, 0x91, 0xa6, 0x47, 0x9c, 0x6e, 0x45, 0xd9,
	0x26, 0x46, 0x78, 0x13, 0x96, 0x4a, 0x7f, 0xfc, 0x7f, 0x9f, 0xa5, 0x51, 0x03, 0x0f, 0x6e, 0xd1,
	0xe3, 0x3d, 0x0f, 0x87, 0x9f, 0xc3, 0xcd, 0x94, 0x19, 0x1b, 0x9b, 0x82, 0x73, 0x61, 0xcc, 0xa8,
	0x48, 0xab, 0xa5, 0x8b, 0x3c, 0x71, 0x65, 0xaa, 0xa5, 0x3e, 0x8b, 0xa9, 0xdf, 0x70, 0x0e, 0x83,
	0x8a, 0xef, 0xb3, 0xf8, 0x0c, 0xd9, 0x93, 0x6d, 0x3c, 0x81, 0x56, 0xe6, 0x6a, 0x70, 0x28, 0x1a,
	0xb3, 0x56, 0x8c, 0x73, 0x3a, 0x90, 0x08, 0x30, 0xa7, 0x55, 0xc7, 0x9b, 0x8a, 0xd2, 0x25, 0x92,
	0x0b, 0x16, 0xbe, 0x84, 0xff, 0x63, 0x82, 0xff, 0x10, 0xc7, 0x67, 0x37, 0x87, 0xd9, 0x5d, 0x77,
	0xec, 0x93, 0xa2, 0x51, 0x6a, 0xb7, 0xe1, 0x82, 0x11, 0xd9, 0xe1, 0x90, 0x26, 0x9a, 0xc7, 0x74,
	0x42, 0x67, 0x9b, 0x0a, 0x60, 0xc2, 0x07, 0xe0, 0xea, 0x25, 0x72, 0x2b, 0x92, 0x23, 0x5e, 0x4d,
	0xf4, 0xba, 0x58, 0xda, 0x8f, 0x7a, 0x6a, 0xf1, 0xa5, 0xe0, 0xc7, 0x79, 0x2e, 0x90, 0x67, 0x69,
	0x3f, 0xe4, 0xf9, 0x08, 0x56, 0x8f, 0xc9, 0x32, 0x1e, 0x31, 0x99, 0x16, 0x5a, 0x98, 0x68, 0x11,
	0xbd, 0x2f, 0x1f, 0xc9, 0xf6, 0x89, 0x27, 0xac, 0xff, 0x19, 0xc0, 0x12, 0x3d, 0x49, 0xc6, 0x79,
	0x66, 0xe9, 0x21, 0xfe, 0x1e, 0x54, 0x77, 0x19, 0x6f, 0x27, 0xdd, 0xda, 0x28, 0x68, 0xcd, 0xb4,
	0xe7, 0x36, 0x7e, 0x0c, 0x3a, 0xa7, 0xa0, 0x54, 0x9d, 0x93, 0x75, 0x63, 0x4a, 0x1a, 0xbc, 0xae,
	0x6d, 0xc2, 0x32, 0x39, 0xe4, 0x2a, 0x95, 0xfc, 0x80, 0xb6, 0x75, 0x1d, 0xe6, 0xfd, 0x27, 0xdd,
	0x05, 0xd2, 0xb3, 0x39, 0xc2, 0xb0, 0xea, 0xeb, 0x7f, 0x05, 0x70, 0xa9, 0xb6, 0x92, 0x11, 0x7a,
	0x5f, 0x72, 0x51, 0xb9, 0x97, 0xdf, 0x35, 0x39, 0x9c, 0xf3, 0x18, 0x6a, 0xc0, 0xff, 0x60, 0xa1,
	0xa4, 0xe4, 0x4c, 0xb3, 0xb1, 0xf1, 0xe2, 0xd8, 0xf4, 0xe8, 0x0e, 0x82, 0xf5, 0x48, 0xf6, 0x20,
	0x2f, 0xd5, 0xa7, 0x8c, 0xf4, 0xea, 0x20, 0x9f, 0x8a, 0xe4, 0x34, 0xad, 0x30, 0x5e, 0x75, 0xca,
	0x48, 0x03, 0x04, 0x49, 0x70, 0x89, 0xe6, 0xc5, 0xa6, 0xe9, 0x04, 0x17, 0x91, 0x5e, 0x12, 0xbe,
	0x01, 0x8b, 0x55, 0x3e, 0x5a, 0xe4, 0x4c, 0x26, 0xa8, 0x2c, 0x8d, 0x7e, 0x19, 0x7c, 0x87, 0xd0,
	0xf5, 0x9f, 0x03, 0x58, 0xa9, 0xed, 0x9b, 0x2b, 0x46, 0x7b, 0xfe, 0x2f, 0x34, 0xb5, 0xf8, 0xba,
	0x10, 0xc6, 0xc6, 0x8c, 0x7f, 0x25, 0x12, 0xdc, 0x74, 0xa3, 0x3f, 0xef, 0xc1, 0xae, 0xc3, 0x9c,
	0xf2, 0x95, 0xa4, 0xda, 0x4b, 0x3b, 0x83, 0x2f, 0x6d, 0xc9, 0x5b, 0x26, 0x0f, 0xab, 0x0d, 0x4b,
	0x5a, 0xe4, 0xe9, 0x41, 0x9d, 0x3b, 0x83, 0xdc, 0x05, 0xc4, 0x2b, 0xe6, 0xfa, 0x6f, 0x8b, 0x70,
	0xe9, 0x84, 0xee, 0x44, 0x47, 0x48, 0xdf, 0x78, 0x84, 0x1b, 0xe5, 0x11, 0x22, 0x86, 0x47, 0x78,
	0x03, 0x16, 0xf2, 0x3c, 0x57, 0x22, 0x36, 0xc5, 0x90, 0x48, 0x77, 0x90, 0x34, 0x8f, 0xe8, 0xa0,
	0x18, 0x96, 0x07, 0x8d, 0x8f, 0x67, 0xc4, 0xca, 0xba, 0xde, 0xa5, 0x83, 0xae, 0x50, 0xac, 0xec,
	0x65, 0x68, 0xec, 0xeb, 0x11, 0x11, 0xee, 0x21, 0xe1, 0xfc, 0xbe, 0x1e, 0xa1, 0x69, 0x0d, 0x80,
	0x4b, 0xcd, 0x0b, 0x69, 0x5d, 0x0d, 0x36, 0xa9, 0xe9, 0x79, 0xa4, 0x97, 0xb8, 0x16, 0xaa, 0xc5,
	0x58, 0x59, 0xac, 0xd0, 0x7d, 0x6a, 0xa1, 0x04, 0xf4, 0x92, 0xf0, 0x1a, 0xcc, 0xa5, 0x99, 0x89,
	0x59, 0x92, 0x68, 0x61, 0x4c, 0xf4, 0x00, 0xcd, 0x90, 0x66, 0xa6, 0x4b, 0x08, 0x12, 0x18, 0xaf,
	0x08, 0x6f, 0x7b, 0x02, 0xe3, 0x25, 0x61, 0x1b, 0xae, 0xd9, 0x22, 0xcb, 0x44, 0x1a, 0xf3, 0x54,
	0x3a, 0x29, 0x60, 0x85, 0xdd, 0x13, 0x99, 0x95, 0x9c, 0x59, 0xdf, 0x83, 0xde, 0x41, 0xa7, 0x55,
	0xa2, 0x6d, 0x21, 0xab, 0x3b, 0x45, 0xea, 0x25, 0xb5, 0x30, 0xee, 0x66, 0xb8, 0x47, 0x76, 0x24,
	0xcc, 0xbb, 0xf5, 0x30, 0x03, 0x64, 0x1d, 0x09, 0x73, 0x0b, 0xc2, 0xaa, 0xf9, 0xe5, 0x55, 0xd6,
	0xef, 0x51, 0x13, 0x2c, 0x07, 0x81, 0xbc, 0xcc, 0x1d, 0x35, 0xb6, 0x64, 0xef, 0x6f, 0x56, 0xfc,
	0x87, 0xc8, 0x0f, 0x2b, 0xfe, 0xfe, 0x66, 0xe9, 0xd1, 0x81, 0x95, 0x29, 0x8f, 0x5c, 0x8b, 0x91,
	0x7c, 0x1d, 0xbd, 0x4f, 0xed, 0xbb, 0xe6, 0xb0, 0x83, 0x06, 0x37, 0x22, 0x24, 0x22, 0x15, 0xbb,
	0x0c, 0xa5, 0xb5, 0xe6, 0xf1, 0x08, 0x3d, 0x56, 0x2a, 0x63, 0xcd, 0xe7, 0x4d, 0x58, 0x46, 0xe6,
	0xe4, 0x5a, 0xc8, 0x24, 0xfa, 0xa0, 0x15, 0xb4, 0xe7, 0xfb, 0x8b, 0xce, 0xd0, 0x2b, 0x71, 0xaa,
	0xdf, 0xb8, 0x56, 0x9e, 0x2e, 0x95, 0x67, 0x3c, 0x29, 0xcf, 0xf1, 0x53, 0xc1, 0x87, 0x27, 0x4c,
	0x05, 0x97, 0xa1, 0x91, 0x31, 0x13, 0xe7, 0x4a, 0xdb, 0x68, 0x8b, 0x6e, 0x59, 0xc6, 0xcc, 0x8e,
	0xd2, 0x36, 0xbc, 0x02, 0x8d, 0xc2, 0x08, 0x8d, 0x17, 0xf0, 0x31, 0xdd, 0xa2, 0xf2, 0x3b, 0xfc,
	0x0f, 0x00, 0x15, 0x1f, 0xad, 0xdb, 0x94, 0xc4, 0x04, 0x09, 0x6f, 0x40, 0x73, 0xa4, 0xf4, 0xd8,
	0xb5, 0xc3, 0x04, 0x29, 0x4f, 0xe8, 0x8a, 0x4f, 0x81, 0xae, 0x1a, 0xd2, 0x54, 0x59, 0x56, 0xf5,
	0x8f, 0x3e, 0xc2, 0x27, 0x1f, 0x4a, 0xe3, 0xf3, 0xac, 0x8a, 0x8e, 0x27, 0x35, 0xed, 0xa1, 0x74,
	0xf4, 0x31, 0xd2, 0x17, 0xa7, 0xe8, 0x4a, 0x3b, 0x25, 0x29, 0x89, 0x34, 0xa4, 0xf5, 0xe8, 0x31,
	0x7a, 0x90, 0xa6, 0xb3, 0x87, 0x70, 0xb5, 0x24, 0x71, 0x2d, 0xe8, 0xe6, 0xd5, 0x64, 0xe2, 0x13,
	0x9a, 0x8a, 0x3c, 0x65, 0xcb, 0x33, 0x26, 0xd2, 0x72, 0x07, 0x2e, 0xca, 0x24, 0x15, 0x7e, 0x74,
	0xe4, 0x7b, 0x2c, 0xdb, 0xa5, 0xd1, 0x24, 0x7a, 0x4a, 0xe5, 0x76, 0x56, 0x5c, 0x69, 0x0b, 0x6d,
	0x38, 0x3b, 0xdc, 0x83, 0x4b, 0x56, 0x59, 0x96, 0xd6, 0xea, 0x93, 0x7a, 0xaf, 0x67, 0xa8, 0xa7,
	0x17, 0xd0, 0x5c, 0x15, 0x29, 0x25, 0x37, 0x1a, 0x3e, 0x85, 0x31, 0xf1, 0x21, 0xf9, 0xf8, 0x94,
	0x96, 0x22, 0x63, 0x6f, 0x4a, 0x44, 0x7e, 0x09, 0x00, 0x26, 0xbd, 0x2a, 0x7a, 0xde, 0x0a, 0xda,
	0x73, 0x1b, 0xdf, 0x9e, 0x66, 0x3f, 0x9d, 0x34, 0xfd, 0x7e, 0x2d, 0xb3, 0xf0, 0xd7, 0x00, 0x16,
	0x9d, 0x6a, 0xd6, 0xba, 0x67, 0xf4, 0x02, 0xbb, 0xff, 0x77, 0xa7, 0x99, 0x6d, 0x2d, 0x9d, 0x7e,
	0xd3, 0x14, 0xc3, 0x1d, 0xfc, 0x7e, 0xec, 0x47, 0x96, 0x4a, 0x4b, 0xaa, 0x8e, 0x98, 0x8d, 0x54,
	0xb4, 0x83, 0x69, 0xff, 0x70, 0xea, 0x43, 0x4b, 0x7d, 0x94, 0xa8, 0xa4, 0x6d, 0xe0, 0x3b, 0x75,
	0x36, 0x52, 0xe1, 0x1f, 0x01, 0x5c, 0xab, 0x2e, 0x3f, 0x5d, 0x5d, 0x35, 0xf2, 0x8f, 0x4a, 0x7e,
	0x83, 0x37, 0x3d, 0x7a, 0x89, 0x9b, 0xf9, 0xfe, 0xd4, 0x37, 0x53, 0xce, 0x07, 0xfd, 0xd5, 0xf2,
	0x2d, 0x62, 0xc2, 0x2f, 0x46, 0xdd, 0x7a, 0xba, 0xc3, 0x73, 0xf8, 0x97, 0xed, 0x9d, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x0b, 0x58, 0x3c, 0x18, 0xfa, 0x0e, 0x00, 0x00,
}
