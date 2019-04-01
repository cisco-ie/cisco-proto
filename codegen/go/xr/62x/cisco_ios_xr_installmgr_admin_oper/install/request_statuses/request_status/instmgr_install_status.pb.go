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
// source: instmgr_install_status.proto

package cisco_ios_xr_installmgr_admin_oper_install_request_statuses_request_status

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

type InstmgrInstallStatus_KEYS struct {
	RequestId            uint32   `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstmgrInstallStatus_KEYS) Reset()         { *m = InstmgrInstallStatus_KEYS{} }
func (m *InstmgrInstallStatus_KEYS) String() string { return proto.CompactTextString(m) }
func (*InstmgrInstallStatus_KEYS) ProtoMessage()    {}
func (*InstmgrInstallStatus_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f971e566766b905, []int{0}
}

func (m *InstmgrInstallStatus_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrInstallStatus_KEYS.Unmarshal(m, b)
}
func (m *InstmgrInstallStatus_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrInstallStatus_KEYS.Marshal(b, m, deterministic)
}
func (m *InstmgrInstallStatus_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrInstallStatus_KEYS.Merge(m, src)
}
func (m *InstmgrInstallStatus_KEYS) XXX_Size() int {
	return xxx_messageInfo_InstmgrInstallStatus_KEYS.Size(m)
}
func (m *InstmgrInstallStatus_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrInstallStatus_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrInstallStatus_KEYS proto.InternalMessageInfo

func (m *InstmgrInstallStatus_KEYS) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

type InstmgrBagRequestInfo_ struct {
	RequestId            uint32   `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TriggerType          string   `protobuf:"bytes,3,opt,name=trigger_type,json=triggerType,proto3" json:"trigger_type,omitempty"`
	RequestOption        int32    `protobuf:"zigzag32,4,opt,name=request_option,json=requestOption,proto3" json:"request_option,omitempty"`
	OperationType        string   `protobuf:"bytes,5,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	OperationDetail      string   `protobuf:"bytes,6,opt,name=operation_detail,json=operationDetail,proto3" json:"operation_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstmgrBagRequestInfo_) Reset()         { *m = InstmgrBagRequestInfo_{} }
func (m *InstmgrBagRequestInfo_) String() string { return proto.CompactTextString(m) }
func (*InstmgrBagRequestInfo_) ProtoMessage()    {}
func (*InstmgrBagRequestInfo_) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f971e566766b905, []int{1}
}

func (m *InstmgrBagRequestInfo_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrBagRequestInfo_.Unmarshal(m, b)
}
func (m *InstmgrBagRequestInfo_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrBagRequestInfo_.Marshal(b, m, deterministic)
}
func (m *InstmgrBagRequestInfo_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrBagRequestInfo_.Merge(m, src)
}
func (m *InstmgrBagRequestInfo_) XXX_Size() int {
	return xxx_messageInfo_InstmgrBagRequestInfo_.Size(m)
}
func (m *InstmgrBagRequestInfo_) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrBagRequestInfo_.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrBagRequestInfo_ proto.InternalMessageInfo

func (m *InstmgrBagRequestInfo_) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *InstmgrBagRequestInfo_) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *InstmgrBagRequestInfo_) GetTriggerType() string {
	if m != nil {
		return m.TriggerType
	}
	return ""
}

func (m *InstmgrBagRequestInfo_) GetRequestOption() int32 {
	if m != nil {
		return m.RequestOption
	}
	return 0
}

func (m *InstmgrBagRequestInfo_) GetOperationType() string {
	if m != nil {
		return m.OperationType
	}
	return ""
}

func (m *InstmgrBagRequestInfo_) GetOperationDetail() string {
	if m != nil {
		return m.OperationDetail
	}
	return ""
}

type InstmgrIssuAbortStatus struct {
	AbortMethod          string   `protobuf:"bytes,1,opt,name=abort_method,json=abortMethod,proto3" json:"abort_method,omitempty"`
	AbortImpact          string   `protobuf:"bytes,2,opt,name=abort_impact,json=abortImpact,proto3" json:"abort_impact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstmgrIssuAbortStatus) Reset()         { *m = InstmgrIssuAbortStatus{} }
func (m *InstmgrIssuAbortStatus) String() string { return proto.CompactTextString(m) }
func (*InstmgrIssuAbortStatus) ProtoMessage()    {}
func (*InstmgrIssuAbortStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f971e566766b905, []int{2}
}

func (m *InstmgrIssuAbortStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrIssuAbortStatus.Unmarshal(m, b)
}
func (m *InstmgrIssuAbortStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrIssuAbortStatus.Marshal(b, m, deterministic)
}
func (m *InstmgrIssuAbortStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrIssuAbortStatus.Merge(m, src)
}
func (m *InstmgrIssuAbortStatus) XXX_Size() int {
	return xxx_messageInfo_InstmgrIssuAbortStatus.Size(m)
}
func (m *InstmgrIssuAbortStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrIssuAbortStatus.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrIssuAbortStatus proto.InternalMessageInfo

func (m *InstmgrIssuAbortStatus) GetAbortMethod() string {
	if m != nil {
		return m.AbortMethod
	}
	return ""
}

func (m *InstmgrIssuAbortStatus) GetAbortImpact() string {
	if m != nil {
		return m.AbortImpact
	}
	return ""
}

type InstmgrBagUserMsgScope_ struct {
	AdminRead            bool     `protobuf:"varint,1,opt,name=admin_read,json=adminRead,proto3" json:"admin_read,omitempty"`
	AffectedSdRs         uint32   `protobuf:"varint,2,opt,name=affected_sd_rs,json=affectedSdRs,proto3" json:"affected_sd_rs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstmgrBagUserMsgScope_) Reset()         { *m = InstmgrBagUserMsgScope_{} }
func (m *InstmgrBagUserMsgScope_) String() string { return proto.CompactTextString(m) }
func (*InstmgrBagUserMsgScope_) ProtoMessage()    {}
func (*InstmgrBagUserMsgScope_) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f971e566766b905, []int{3}
}

func (m *InstmgrBagUserMsgScope_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrBagUserMsgScope_.Unmarshal(m, b)
}
func (m *InstmgrBagUserMsgScope_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrBagUserMsgScope_.Marshal(b, m, deterministic)
}
func (m *InstmgrBagUserMsgScope_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrBagUserMsgScope_.Merge(m, src)
}
func (m *InstmgrBagUserMsgScope_) XXX_Size() int {
	return xxx_messageInfo_InstmgrBagUserMsgScope_.Size(m)
}
func (m *InstmgrBagUserMsgScope_) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrBagUserMsgScope_.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrBagUserMsgScope_ proto.InternalMessageInfo

func (m *InstmgrBagUserMsgScope_) GetAdminRead() bool {
	if m != nil {
		return m.AdminRead
	}
	return false
}

func (m *InstmgrBagUserMsgScope_) GetAffectedSdRs() uint32 {
	if m != nil {
		return m.AffectedSdRs
	}
	return 0
}

type InstmgrBagUserMsg_ struct {
	Category             string                   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Scope                *InstmgrBagUserMsgScope_ `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	Message              string                   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *InstmgrBagUserMsg_) Reset()         { *m = InstmgrBagUserMsg_{} }
func (m *InstmgrBagUserMsg_) String() string { return proto.CompactTextString(m) }
func (*InstmgrBagUserMsg_) ProtoMessage()    {}
func (*InstmgrBagUserMsg_) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f971e566766b905, []int{4}
}

func (m *InstmgrBagUserMsg_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrBagUserMsg_.Unmarshal(m, b)
}
func (m *InstmgrBagUserMsg_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrBagUserMsg_.Marshal(b, m, deterministic)
}
func (m *InstmgrBagUserMsg_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrBagUserMsg_.Merge(m, src)
}
func (m *InstmgrBagUserMsg_) XXX_Size() int {
	return xxx_messageInfo_InstmgrBagUserMsg_.Size(m)
}
func (m *InstmgrBagUserMsg_) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrBagUserMsg_.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrBagUserMsg_ proto.InternalMessageInfo

func (m *InstmgrBagUserMsg_) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *InstmgrBagUserMsg_) GetScope() *InstmgrBagUserMsgScope_ {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *InstmgrBagUserMsg_) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type InstmgrBag_IINodeState_ struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstmgrBag_IINodeState_) Reset()         { *m = InstmgrBag_IINodeState_{} }
func (m *InstmgrBag_IINodeState_) String() string { return proto.CompactTextString(m) }
func (*InstmgrBag_IINodeState_) ProtoMessage()    {}
func (*InstmgrBag_IINodeState_) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f971e566766b905, []int{5}
}

func (m *InstmgrBag_IINodeState_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrBag_IINodeState_.Unmarshal(m, b)
}
func (m *InstmgrBag_IINodeState_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrBag_IINodeState_.Marshal(b, m, deterministic)
}
func (m *InstmgrBag_IINodeState_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrBag_IINodeState_.Merge(m, src)
}
func (m *InstmgrBag_IINodeState_) XXX_Size() int {
	return xxx_messageInfo_InstmgrBag_IINodeState_.Size(m)
}
func (m *InstmgrBag_IINodeState_) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrBag_IINodeState_.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrBag_IINodeState_ proto.InternalMessageInfo

func (m *InstmgrBag_IINodeState_) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *InstmgrBag_IINodeState_) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type InstmgrBag_IIInfo_ struct {
	Direction            string                     `protobuf:"bytes,1,opt,name=direction,proto3" json:"direction,omitempty"`
	Nodes                []*InstmgrBag_IINodeState_ `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
	IiError              string                     `protobuf:"bytes,3,opt,name=ii_error,json=iiError,proto3" json:"ii_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *InstmgrBag_IIInfo_) Reset()         { *m = InstmgrBag_IIInfo_{} }
func (m *InstmgrBag_IIInfo_) String() string { return proto.CompactTextString(m) }
func (*InstmgrBag_IIInfo_) ProtoMessage()    {}
func (*InstmgrBag_IIInfo_) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f971e566766b905, []int{6}
}

func (m *InstmgrBag_IIInfo_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrBag_IIInfo_.Unmarshal(m, b)
}
func (m *InstmgrBag_IIInfo_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrBag_IIInfo_.Marshal(b, m, deterministic)
}
func (m *InstmgrBag_IIInfo_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrBag_IIInfo_.Merge(m, src)
}
func (m *InstmgrBag_IIInfo_) XXX_Size() int {
	return xxx_messageInfo_InstmgrBag_IIInfo_.Size(m)
}
func (m *InstmgrBag_IIInfo_) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrBag_IIInfo_.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrBag_IIInfo_ proto.InternalMessageInfo

func (m *InstmgrBag_IIInfo_) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *InstmgrBag_IIInfo_) GetNodes() []*InstmgrBag_IINodeState_ {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *InstmgrBag_IIInfo_) GetIiError() string {
	if m != nil {
		return m.IiError
	}
	return ""
}

type InstmgrInstallStatus struct {
	RequestInformation            *InstmgrBagRequestInfo_ `protobuf:"bytes,50,opt,name=request_information,json=requestInformation,proto3" json:"request_information,omitempty"`
	Percentage                    uint32                  `protobuf:"varint,51,opt,name=percentage,proto3" json:"percentage,omitempty"`
	AbortState                    string                  `protobuf:"bytes,52,opt,name=abort_state,json=abortState,proto3" json:"abort_state,omitempty"`
	AbortStatus                   *InstmgrIssuAbortStatus `protobuf:"bytes,53,opt,name=abort_status,json=abortStatus,proto3" json:"abort_status,omitempty"`
	IssuMessage                   []*InstmgrBagUserMsg_   `protobuf:"bytes,54,rep,name=issu_message,json=issuMessage,proto3" json:"issu_message,omitempty"`
	Message                       []*InstmgrBagUserMsg_   `protobuf:"bytes,55,rep,name=message,proto3" json:"message,omitempty"`
	IncrementalInstallInformation *InstmgrBag_IIInfo_     `protobuf:"bytes,56,opt,name=incremental_install_information,json=incrementalInstallInformation,proto3" json:"incremental_install_information,omitempty"`
	DownloadedBytes               uint32                  `protobuf:"varint,57,opt,name=downloaded_bytes,json=downloadedBytes,proto3" json:"downloaded_bytes,omitempty"`
	UnansweredQuery               bool                    `protobuf:"varint,58,opt,name=unanswered_query,json=unansweredQuery,proto3" json:"unanswered_query,omitempty"`
	OperationPhase                string                  `protobuf:"bytes,59,opt,name=operation_phase,json=operationPhase,proto3" json:"operation_phase,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                `json:"-"`
	XXX_unrecognized              []byte                  `json:"-"`
	XXX_sizecache                 int32                   `json:"-"`
}

func (m *InstmgrInstallStatus) Reset()         { *m = InstmgrInstallStatus{} }
func (m *InstmgrInstallStatus) String() string { return proto.CompactTextString(m) }
func (*InstmgrInstallStatus) ProtoMessage()    {}
func (*InstmgrInstallStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f971e566766b905, []int{7}
}

func (m *InstmgrInstallStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrInstallStatus.Unmarshal(m, b)
}
func (m *InstmgrInstallStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrInstallStatus.Marshal(b, m, deterministic)
}
func (m *InstmgrInstallStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrInstallStatus.Merge(m, src)
}
func (m *InstmgrInstallStatus) XXX_Size() int {
	return xxx_messageInfo_InstmgrInstallStatus.Size(m)
}
func (m *InstmgrInstallStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrInstallStatus.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrInstallStatus proto.InternalMessageInfo

func (m *InstmgrInstallStatus) GetRequestInformation() *InstmgrBagRequestInfo_ {
	if m != nil {
		return m.RequestInformation
	}
	return nil
}

func (m *InstmgrInstallStatus) GetPercentage() uint32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *InstmgrInstallStatus) GetAbortState() string {
	if m != nil {
		return m.AbortState
	}
	return ""
}

func (m *InstmgrInstallStatus) GetAbortStatus() *InstmgrIssuAbortStatus {
	if m != nil {
		return m.AbortStatus
	}
	return nil
}

func (m *InstmgrInstallStatus) GetIssuMessage() []*InstmgrBagUserMsg_ {
	if m != nil {
		return m.IssuMessage
	}
	return nil
}

func (m *InstmgrInstallStatus) GetMessage() []*InstmgrBagUserMsg_ {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *InstmgrInstallStatus) GetIncrementalInstallInformation() *InstmgrBag_IIInfo_ {
	if m != nil {
		return m.IncrementalInstallInformation
	}
	return nil
}

func (m *InstmgrInstallStatus) GetDownloadedBytes() uint32 {
	if m != nil {
		return m.DownloadedBytes
	}
	return 0
}

func (m *InstmgrInstallStatus) GetUnansweredQuery() bool {
	if m != nil {
		return m.UnansweredQuery
	}
	return false
}

func (m *InstmgrInstallStatus) GetOperationPhase() string {
	if m != nil {
		return m.OperationPhase
	}
	return ""
}

func init() {
	proto.RegisterType((*InstmgrInstallStatus_KEYS)(nil), "cisco_ios_xr_installmgr_admin_oper.install.request_statuses.request_status.instmgr_install_status_KEYS")
	proto.RegisterType((*InstmgrBagRequestInfo_)(nil), "cisco_ios_xr_installmgr_admin_oper.install.request_statuses.request_status.instmgr_bag_request_info_")
	proto.RegisterType((*InstmgrIssuAbortStatus)(nil), "cisco_ios_xr_installmgr_admin_oper.install.request_statuses.request_status.instmgr_issu_abort_status")
	proto.RegisterType((*InstmgrBagUserMsgScope_)(nil), "cisco_ios_xr_installmgr_admin_oper.install.request_statuses.request_status.instmgr_bag_user_msg_scope_")
	proto.RegisterType((*InstmgrBagUserMsg_)(nil), "cisco_ios_xr_installmgr_admin_oper.install.request_statuses.request_status.instmgr_bag_user_msg_")
	proto.RegisterType((*InstmgrBag_IINodeState_)(nil), "cisco_ios_xr_installmgr_admin_oper.install.request_statuses.request_status.instmgr_bag_II_node_state_")
	proto.RegisterType((*InstmgrBag_IIInfo_)(nil), "cisco_ios_xr_installmgr_admin_oper.install.request_statuses.request_status.instmgr_bag_II_info_")
	proto.RegisterType((*InstmgrInstallStatus)(nil), "cisco_ios_xr_installmgr_admin_oper.install.request_statuses.request_status.instmgr_install_status")
}

func init() { proto.RegisterFile("instmgr_install_status.proto", fileDescriptor_4f971e566766b905) }

var fileDescriptor_4f971e566766b905 = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcd, 0x72, 0xd3, 0x48,
	0x10, 0xc7, 0x4b, 0x9b, 0x75, 0x12, 0xb7, 0xe3, 0x24, 0x3b, 0x9b, 0xdd, 0x55, 0xbe, 0x36, 0x59,
	0xd5, 0x6e, 0xad, 0x73, 0xf1, 0x21, 0xd9, 0x4f, 0xe0, 0x44, 0x91, 0x83, 0xa1, 0x42, 0x40, 0xe6,
	0xc2, 0x69, 0x18, 0x4b, 0x6d, 0x65, 0xaa, 0x2c, 0x8d, 0x32, 0x23, 0x57, 0x70, 0x01, 0x37, 0x0e,
	0x3c, 0x01, 0xcf, 0x40, 0x15, 0xaf, 0x02, 0xcf, 0xc2, 0x2b, 0x50, 0xd3, 0x23, 0x59, 0x4e, 0x48,
	0xc1, 0x05, 0x73, 0x72, 0xf5, 0xcf, 0x3d, 0xdd, 0xd3, 0xdd, 0xff, 0x1e, 0xc1, 0x8e, 0xcc, 0x4c,
	0x91, 0x26, 0x9a, 0xdb, 0x5f, 0x31, 0x1a, 0x71, 0x53, 0x88, 0x62, 0x6c, 0xba, 0xb9, 0x56, 0x85,
	0x62, 0x77, 0x23, 0x69, 0x22, 0xc5, 0xa5, 0x32, 0xfc, 0xe9, 0xd4, 0xc5, 0x7a, 0x8b, 0x38, 0x95,
	0x19, 0x57, 0x39, 0xea, 0x6e, 0x49, 0xbb, 0x1a, 0xcf, 0xc7, 0x68, 0x8a, 0x32, 0x00, 0x9a, 0x2b,
	0x20, 0xb8, 0x05, 0xdb, 0xd7, 0xe7, 0xe2, 0xf7, 0x8e, 0x1f, 0xf7, 0xd9, 0x2e, 0x40, 0x75, 0x40,
	0xc6, 0xbe, 0xb7, 0xef, 0x75, 0xda, 0x61, 0xb3, 0x24, 0xbd, 0x38, 0xf8, 0xe0, 0xc1, 0x66, 0x75,
	0x7c, 0x20, 0x12, 0x3e, 0xf5, 0xcd, 0x86, 0x8a, 0x7f, 0xe1, 0x30, 0xfb, 0x05, 0x96, 0xc6, 0x06,
	0xb5, 0xfd, 0xef, 0xbb, 0x7d, 0xaf, 0xd3, 0x0c, 0x17, 0xad, 0xd9, 0x8b, 0xd9, 0x6f, 0xb0, 0x52,
	0x68, 0x99, 0x24, 0xa8, 0x79, 0x31, 0xc9, 0xd1, 0x5f, 0xa0, 0x7f, 0x5b, 0x25, 0x7b, 0x34, 0xc9,
	0x91, 0xfd, 0x01, 0xab, 0x55, 0x68, 0x95, 0x17, 0x52, 0x65, 0xfe, 0xf7, 0xfb, 0x5e, 0xe7, 0x87,
	0xb0, 0x5d, 0xd2, 0x53, 0x82, 0xd6, 0xcd, 0x76, 0x43, 0x58, 0xc3, 0xc5, 0x6a, 0x50, 0xac, 0xf6,
	0x94, 0x52, 0xb4, 0x03, 0x58, 0xaf, 0xdd, 0x62, 0x2c, 0x84, 0x1c, 0xf9, 0x8b, 0xe4, 0xb8, 0x36,
	0xe5, 0x77, 0x08, 0x07, 0xa2, 0x2e, 0x58, 0x1a, 0x33, 0xe6, 0x62, 0xa0, 0x74, 0xd5, 0x4c, 0x7b,
	0x71, 0x67, 0xa7, 0x58, 0x9c, 0x29, 0x57, 0x72, 0x33, 0x6c, 0x11, 0x3b, 0x21, 0x54, 0xbb, 0xc8,
	0x34, 0x17, 0x51, 0x51, 0x56, 0xee, 0x5c, 0x7a, 0x84, 0x82, 0x41, 0x3d, 0x12, 0xdb, 0x53, 0xea,
	0x51, 0x6a, 0x12, 0x6e, 0x22, 0x95, 0x23, 0x75, 0xd5, 0xcd, 0x59, 0xa3, 0x70, 0x29, 0x96, 0xc3,
	0x26, 0x91, 0x10, 0x45, 0xcc, 0x7e, 0x87, 0x55, 0x31, 0x1c, 0x62, 0x54, 0x60, 0xcc, 0x4d, 0xcc,
	0xb5, 0xa1, 0x14, 0xed, 0x70, 0xa5, 0xa2, 0xfd, 0x38, 0x34, 0xc1, 0x3b, 0x0f, 0x7e, 0xba, 0x36,
	0x09, 0xdb, 0x82, 0xe5, 0x48, 0x14, 0x98, 0x28, 0x3d, 0x29, 0xef, 0x3f, 0xb5, 0xd9, 0x0b, 0x68,
	0xd0, 0x25, 0x28, 0x64, 0xeb, 0x30, 0xe9, 0x7e, 0x3d, 0x21, 0x76, 0x3f, 0x53, 0x72, 0xe8, 0xb2,
	0x32, 0x1f, 0x96, 0x52, 0x34, 0x46, 0x24, 0x95, 0x24, 0x2a, 0x33, 0x38, 0x85, 0xad, 0xd9, 0xf3,
	0xbd, 0x1e, 0xcf, 0x54, 0x8c, 0x14, 0x1a, 0x39, 0xdb, 0x86, 0x26, 0x99, 0x99, 0x48, 0xb1, 0xaa,
	0xc9, 0x82, 0xfb, 0x22, 0x45, 0xb6, 0x01, 0x0d, 0x72, 0x2b, 0x27, 0xe1, 0x8c, 0xe0, 0xbd, 0x07,
	0x1b, 0x57, 0x22, 0x3a, 0x4d, 0xef, 0x40, 0x33, 0x96, 0x1a, 0x23, 0xd2, 0x9c, 0x8b, 0x55, 0x03,
	0xf6, 0x1c, 0x1a, 0x36, 0xb0, 0xed, 0xf9, 0x42, 0xa7, 0x75, 0x38, 0x9c, 0x57, 0x83, 0x2e, 0x17,
	0x18, 0xba, 0xa4, 0x6c, 0x13, 0x96, 0xa5, 0xe4, 0xa8, 0xb5, 0xd2, 0x55, 0x83, 0xa4, 0x3c, 0xb6,
	0x66, 0xf0, 0x76, 0x09, 0x7e, 0xbe, 0x7e, 0xcf, 0xd9, 0x6b, 0x0f, 0x7e, 0x9c, 0xdd, 0x5b, 0x9d,
	0x92, 0xde, 0xfd, 0x43, 0x9a, 0x31, 0xce, 0xab, 0x84, 0x4b, 0x4f, 0x45, 0xc8, 0xaa, 0x67, 0xa1,
	0xbe, 0x00, 0xfb, 0x15, 0x20, 0x47, 0x1d, 0x61, 0x56, 0xd8, 0x89, 0x1f, 0x91, 0x8a, 0x67, 0x08,
	0xdb, 0x83, 0x56, 0xbd, 0x7d, 0xe8, 0xff, 0x45, 0x15, 0x03, 0xa1, 0xbe, 0x25, 0xec, 0x95, 0x57,
	0x2d, 0x9b, 0xcb, 0xef, 0xff, 0x3d, 0xbf, 0x92, 0x3e, 0x79, 0x0c, 0xca, 0x9d, 0xee, 0xbb, 0x26,
	0xbf, 0xf4, 0x60, 0x85, 0x5c, 0x2a, 0x01, 0xff, 0x43, 0x02, 0x11, 0x73, 0xdf, 0xa0, 0xb0, 0x65,
	0xd3, 0x9e, 0xb8, 0xac, 0xec, 0x59, 0xbd, 0x41, 0xff, 0x7e, 0xab, 0x0b, 0x54, 0x19, 0xd9, 0x1b,
	0x0f, 0xf6, 0x64, 0x16, 0x69, 0x4c, 0xed, 0x00, 0x47, 0x53, 0x1d, 0xce, 0x8a, 0xee, 0x3f, 0x9a,
	0xd0, 0x93, 0x39, 0xee, 0x8d, 0xd3, 0xdb, 0xee, 0xcc, 0x45, 0x7a, 0x2e, 0xd0, 0xac, 0xf4, 0x0e,
	0x60, 0x3d, 0x56, 0x17, 0xd9, 0x48, 0x89, 0x18, 0x63, 0x3e, 0x98, 0x14, 0x68, 0xfc, 0xff, 0x49,
	0x80, 0x6b, 0x35, 0xbf, 0x6d, 0xb1, 0x75, 0x1d, 0x67, 0x22, 0x33, 0x17, 0xa8, 0x31, 0xe6, 0xe7,
	0x63, 0xd4, 0x13, 0xff, 0x06, 0x3d, 0xca, 0x6b, 0x35, 0x7f, 0x68, 0x31, 0xfb, 0x13, 0xea, 0xcf,
	0x09, 0xcf, 0xcf, 0x84, 0x41, 0xff, 0x26, 0x89, 0xb6, 0xfe, 0x48, 0x3d, 0xb0, 0x74, 0xb0, 0x48,
	0xdf, 0xf9, 0xa3, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xdc, 0x2c, 0xfc, 0x07, 0x08, 0x00,
	0x00,
}
