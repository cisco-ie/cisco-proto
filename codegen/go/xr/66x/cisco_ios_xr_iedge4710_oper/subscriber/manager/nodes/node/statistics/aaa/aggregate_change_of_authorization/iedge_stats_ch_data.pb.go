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
// source: iedge_stats_ch_data.proto

package cisco_ios_xr_iedge4710_oper_subscriber_manager_nodes_node_statistics_aaa_aggregate_change_of_authorization

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

type IedgeStatsChData_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IedgeStatsChData_KEYS) Reset()         { *m = IedgeStatsChData_KEYS{} }
func (m *IedgeStatsChData_KEYS) String() string { return proto.CompactTextString(m) }
func (*IedgeStatsChData_KEYS) ProtoMessage()    {}
func (*IedgeStatsChData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8089d26fbc9ea0d, []int{0}
}

func (m *IedgeStatsChData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IedgeStatsChData_KEYS.Unmarshal(m, b)
}
func (m *IedgeStatsChData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IedgeStatsChData_KEYS.Marshal(b, m, deterministic)
}
func (m *IedgeStatsChData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IedgeStatsChData_KEYS.Merge(m, src)
}
func (m *IedgeStatsChData_KEYS) XXX_Size() int {
	return xxx_messageInfo_IedgeStatsChData_KEYS.Size(m)
}
func (m *IedgeStatsChData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IedgeStatsChData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IedgeStatsChData_KEYS proto.InternalMessageInfo

func (m *IedgeStatsChData_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type IedgeStatsChReqStats struct {
	ReceivedRequests        uint64   `protobuf:"varint,1,opt,name=received_requests,json=receivedRequests,proto3" json:"received_requests,omitempty"`
	AcknowledgedRequests    uint64   `protobuf:"varint,2,opt,name=acknowledged_requests,json=acknowledgedRequests,proto3" json:"acknowledged_requests,omitempty"`
	NonAcknowledgedRequests uint64   `protobuf:"varint,3,opt,name=non_acknowledged_requests,json=nonAcknowledgedRequests,proto3" json:"non_acknowledged_requests,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *IedgeStatsChReqStats) Reset()         { *m = IedgeStatsChReqStats{} }
func (m *IedgeStatsChReqStats) String() string { return proto.CompactTextString(m) }
func (*IedgeStatsChReqStats) ProtoMessage()    {}
func (*IedgeStatsChReqStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8089d26fbc9ea0d, []int{1}
}

func (m *IedgeStatsChReqStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IedgeStatsChReqStats.Unmarshal(m, b)
}
func (m *IedgeStatsChReqStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IedgeStatsChReqStats.Marshal(b, m, deterministic)
}
func (m *IedgeStatsChReqStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IedgeStatsChReqStats.Merge(m, src)
}
func (m *IedgeStatsChReqStats) XXX_Size() int {
	return xxx_messageInfo_IedgeStatsChReqStats.Size(m)
}
func (m *IedgeStatsChReqStats) XXX_DiscardUnknown() {
	xxx_messageInfo_IedgeStatsChReqStats.DiscardUnknown(m)
}

var xxx_messageInfo_IedgeStatsChReqStats proto.InternalMessageInfo

func (m *IedgeStatsChReqStats) GetReceivedRequests() uint64 {
	if m != nil {
		return m.ReceivedRequests
	}
	return 0
}

func (m *IedgeStatsChReqStats) GetAcknowledgedRequests() uint64 {
	if m != nil {
		return m.AcknowledgedRequests
	}
	return 0
}

func (m *IedgeStatsChReqStats) GetNonAcknowledgedRequests() uint64 {
	if m != nil {
		return m.NonAcknowledgedRequests
	}
	return 0
}

type IedgeStatsChData struct {
	AccountLogon                   *IedgeStatsChReqStats `protobuf:"bytes,50,opt,name=account_logon,json=accountLogon,proto3" json:"account_logon,omitempty"`
	AccountLogoff                  *IedgeStatsChReqStats `protobuf:"bytes,51,opt,name=account_logoff,json=accountLogoff,proto3" json:"account_logoff,omitempty"`
	AccountUpdate                  *IedgeStatsChReqStats `protobuf:"bytes,52,opt,name=account_update,json=accountUpdate,proto3" json:"account_update,omitempty"`
	SessionDisconnect              *IedgeStatsChReqStats `protobuf:"bytes,53,opt,name=session_disconnect,json=sessionDisconnect,proto3" json:"session_disconnect,omitempty"`
	SingleServiceLogon             *IedgeStatsChReqStats `protobuf:"bytes,54,opt,name=single_service_logon,json=singleServiceLogon,proto3" json:"single_service_logon,omitempty"`
	SingleServiceLogoff            *IedgeStatsChReqStats `protobuf:"bytes,55,opt,name=single_service_logoff,json=singleServiceLogoff,proto3" json:"single_service_logoff,omitempty"`
	SingleServiceModify            *IedgeStatsChReqStats `protobuf:"bytes,56,opt,name=single_service_modify,json=singleServiceModify,proto3" json:"single_service_modify,omitempty"`
	ServiceMulti                   *IedgeStatsChReqStats `protobuf:"bytes,57,opt,name=service_multi,json=serviceMulti,proto3" json:"service_multi,omitempty"`
	UnknownAccountCmdResps         uint64                `protobuf:"varint,58,opt,name=unknown_account_cmd_resps,json=unknownAccountCmdResps,proto3" json:"unknown_account_cmd_resps,omitempty"`
	UnknownServiceCmdResps         uint64                `protobuf:"varint,59,opt,name=unknown_service_cmd_resps,json=unknownServiceCmdResps,proto3" json:"unknown_service_cmd_resps,omitempty"`
	UnknownCmdResps                uint64                `protobuf:"varint,60,opt,name=unknown_cmd_resps,json=unknownCmdResps,proto3" json:"unknown_cmd_resps,omitempty"`
	AttrListRetrieveFailureResps   uint64                `protobuf:"varint,61,opt,name=attr_list_retrieve_failure_resps,json=attrListRetrieveFailureResps,proto3" json:"attr_list_retrieve_failure_resps,omitempty"`
	RespSendFailure                uint64                `protobuf:"varint,62,opt,name=resp_send_failure,json=respSendFailure,proto3" json:"resp_send_failure,omitempty"`
	InternalErrResps               uint64                `protobuf:"varint,63,opt,name=internal_err_resps,json=internalErrResps,proto3" json:"internal_err_resps,omitempty"`
	ServiceProfilePushFailureResps uint64                `protobuf:"varint,64,opt,name=service_profile_push_failure_resps,json=serviceProfilePushFailureResps,proto3" json:"service_profile_push_failure_resps,omitempty"`
	NoCmdResps                     uint64                `protobuf:"varint,65,opt,name=no_cmd_resps,json=noCmdResps,proto3" json:"no_cmd_resps,omitempty"`
	NoSessionFoundResps            uint64                `protobuf:"varint,66,opt,name=no_session_found_resps,json=noSessionFoundResps,proto3" json:"no_session_found_resps,omitempty"`
	NoSessionPeerResps             uint64                `protobuf:"varint,67,opt,name=no_session_peer_resps,json=noSessionPeerResps,proto3" json:"no_session_peer_resps,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}              `json:"-"`
	XXX_unrecognized               []byte                `json:"-"`
	XXX_sizecache                  int32                 `json:"-"`
}

func (m *IedgeStatsChData) Reset()         { *m = IedgeStatsChData{} }
func (m *IedgeStatsChData) String() string { return proto.CompactTextString(m) }
func (*IedgeStatsChData) ProtoMessage()    {}
func (*IedgeStatsChData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8089d26fbc9ea0d, []int{2}
}

func (m *IedgeStatsChData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IedgeStatsChData.Unmarshal(m, b)
}
func (m *IedgeStatsChData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IedgeStatsChData.Marshal(b, m, deterministic)
}
func (m *IedgeStatsChData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IedgeStatsChData.Merge(m, src)
}
func (m *IedgeStatsChData) XXX_Size() int {
	return xxx_messageInfo_IedgeStatsChData.Size(m)
}
func (m *IedgeStatsChData) XXX_DiscardUnknown() {
	xxx_messageInfo_IedgeStatsChData.DiscardUnknown(m)
}

var xxx_messageInfo_IedgeStatsChData proto.InternalMessageInfo

func (m *IedgeStatsChData) GetAccountLogon() *IedgeStatsChReqStats {
	if m != nil {
		return m.AccountLogon
	}
	return nil
}

func (m *IedgeStatsChData) GetAccountLogoff() *IedgeStatsChReqStats {
	if m != nil {
		return m.AccountLogoff
	}
	return nil
}

func (m *IedgeStatsChData) GetAccountUpdate() *IedgeStatsChReqStats {
	if m != nil {
		return m.AccountUpdate
	}
	return nil
}

func (m *IedgeStatsChData) GetSessionDisconnect() *IedgeStatsChReqStats {
	if m != nil {
		return m.SessionDisconnect
	}
	return nil
}

func (m *IedgeStatsChData) GetSingleServiceLogon() *IedgeStatsChReqStats {
	if m != nil {
		return m.SingleServiceLogon
	}
	return nil
}

func (m *IedgeStatsChData) GetSingleServiceLogoff() *IedgeStatsChReqStats {
	if m != nil {
		return m.SingleServiceLogoff
	}
	return nil
}

func (m *IedgeStatsChData) GetSingleServiceModify() *IedgeStatsChReqStats {
	if m != nil {
		return m.SingleServiceModify
	}
	return nil
}

func (m *IedgeStatsChData) GetServiceMulti() *IedgeStatsChReqStats {
	if m != nil {
		return m.ServiceMulti
	}
	return nil
}

func (m *IedgeStatsChData) GetUnknownAccountCmdResps() uint64 {
	if m != nil {
		return m.UnknownAccountCmdResps
	}
	return 0
}

func (m *IedgeStatsChData) GetUnknownServiceCmdResps() uint64 {
	if m != nil {
		return m.UnknownServiceCmdResps
	}
	return 0
}

func (m *IedgeStatsChData) GetUnknownCmdResps() uint64 {
	if m != nil {
		return m.UnknownCmdResps
	}
	return 0
}

func (m *IedgeStatsChData) GetAttrListRetrieveFailureResps() uint64 {
	if m != nil {
		return m.AttrListRetrieveFailureResps
	}
	return 0
}

func (m *IedgeStatsChData) GetRespSendFailure() uint64 {
	if m != nil {
		return m.RespSendFailure
	}
	return 0
}

func (m *IedgeStatsChData) GetInternalErrResps() uint64 {
	if m != nil {
		return m.InternalErrResps
	}
	return 0
}

func (m *IedgeStatsChData) GetServiceProfilePushFailureResps() uint64 {
	if m != nil {
		return m.ServiceProfilePushFailureResps
	}
	return 0
}

func (m *IedgeStatsChData) GetNoCmdResps() uint64 {
	if m != nil {
		return m.NoCmdResps
	}
	return 0
}

func (m *IedgeStatsChData) GetNoSessionFoundResps() uint64 {
	if m != nil {
		return m.NoSessionFoundResps
	}
	return 0
}

func (m *IedgeStatsChData) GetNoSessionPeerResps() uint64 {
	if m != nil {
		return m.NoSessionPeerResps
	}
	return 0
}

func init() {
	proto.RegisterType((*IedgeStatsChData_KEYS)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.manager.nodes.node.statistics.aaa.aggregate_change_of_authorization.iedge_stats_ch_data_KEYS")
	proto.RegisterType((*IedgeStatsChReqStats)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.manager.nodes.node.statistics.aaa.aggregate_change_of_authorization.iedge_stats_ch_req_stats")
	proto.RegisterType((*IedgeStatsChData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.manager.nodes.node.statistics.aaa.aggregate_change_of_authorization.iedge_stats_ch_data")
}

func init() { proto.RegisterFile("iedge_stats_ch_data.proto", fileDescriptor_d8089d26fbc9ea0d) }

var fileDescriptor_d8089d26fbc9ea0d = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xcb, 0x6e, 0xd4, 0x3c,
	0x14, 0xc7, 0x95, 0xef, 0x43, 0x88, 0x9a, 0x96, 0x52, 0xf7, 0x42, 0x2a, 0x10, 0x1a, 0xcd, 0xaa,
	0x02, 0x14, 0xd1, 0x4e, 0xa1, 0xb4, 0x5c, 0x4b, 0x69, 0x17, 0xd0, 0xa2, 0x6a, 0x46, 0x2c, 0x58,
	0x59, 0x6e, 0x72, 0x92, 0x31, 0x64, 0x8e, 0x07, 0xdb, 0x29, 0x97, 0x35, 0x4f, 0x83, 0x58, 0x22,
	0x01, 0x6f, 0xc1, 0x23, 0x21, 0xdb, 0x49, 0x9a, 0xd2, 0x61, 0x8b, 0xb2, 0x19, 0x69, 0x7c, 0xfe,
	0x3f, 0xfb, 0xe7, 0x33, 0xd1, 0x99, 0x90, 0x65, 0x01, 0x49, 0x06, 0x4c, 0x1b, 0x6e, 0x34, 0x8b,
	0x87, 0x2c, 0xe1, 0x86, 0x47, 0x63, 0x25, 0x8d, 0xa4, 0x6f, 0x62, 0xa1, 0x63, 0xc9, 0x84, 0xd4,
	0xec, 0x83, 0x62, 0x2e, 0xb7, 0xbe, 0xb1, 0x7a, 0x9b, 0xc9, 0x31, 0xa8, 0x48, 0x17, 0x47, 0x3a,
	0x56, 0xe2, 0x08, 0x54, 0x34, 0xe2, 0xc8, 0x33, 0x50, 0x11, 0xca, 0x04, 0xb4, 0xfb, 0x8c, 0xec,
	0x76, 0x42, 0x1b, 0x11, 0xeb, 0x88, 0x73, 0x1e, 0xf1, 0x2c, 0x53, 0x90, 0x71, 0x03, 0x2c, 0x1e,
	0x72, 0xcc, 0x80, 0xc9, 0x94, 0xf1, 0xc2, 0x0c, 0xa5, 0x12, 0x9f, 0xb8, 0x11, 0x12, 0xbb, 0x1b,
	0x24, 0x9c, 0x20, 0xc2, 0x5e, 0xec, 0xbe, 0x1e, 0xd0, 0xab, 0x64, 0xca, 0x6e, 0xca, 0x90, 0x8f,
	0x20, 0x0c, 0x3a, 0xc1, 0xca, 0x54, 0xff, 0x82, 0x5d, 0x78, 0xc9, 0x47, 0xd0, 0xfd, 0x11, 0x9c,
	0x21, 0x15, 0xbc, 0xf3, 0x5f, 0xe8, 0x4d, 0x32, 0xa7, 0x20, 0x06, 0x71, 0x0c, 0x89, 0x5d, 0x2d,
	0x40, 0x1b, 0xed, 0x76, 0x38, 0xd7, 0xbf, 0x5c, 0x15, 0xfa, 0xe5, 0x3a, 0xed, 0x91, 0x45, 0x1e,
	0xbf, 0x45, 0xf9, 0x3e, 0xb7, 0xdb, 0x35, 0x80, 0xff, 0x1c, 0xb0, 0xd0, 0x2c, 0xd6, 0xd0, 0x16,
	0x59, 0x46, 0x89, 0x6c, 0x32, 0xf8, 0xbf, 0x03, 0xaf, 0xa0, 0xc4, 0xed, 0x09, 0x6c, 0xf7, 0xd7,
	0x2c, 0x99, 0x9f, 0x70, 0x69, 0xfa, 0x25, 0x20, 0x33, 0x3c, 0x8e, 0x65, 0x81, 0x86, 0xe5, 0x32,
	0x93, 0x18, 0xae, 0x75, 0x82, 0x95, 0x8b, 0x6b, 0x9f, 0x83, 0xe8, 0xdf, 0xfd, 0x22, 0xd1, 0xdf,
	0x9a, 0xda, 0x9f, 0x2e, 0xdd, 0xf6, 0xad, 0x1a, 0xfd, 0x1a, 0x90, 0x4b, 0x4d, 0xd9, 0x34, 0x0d,
	0x7b, 0x6d, 0xb2, 0x9d, 0x69, 0xd8, 0xa6, 0xe9, 0x29, 0xdd, 0x62, 0x9c, 0x70, 0x03, 0xe1, 0x7a,
	0x1b, 0x75, 0x5f, 0x39, 0x37, 0xfa, 0x2d, 0x20, 0x54, 0x83, 0xd6, 0x42, 0x22, 0x4b, 0xac, 0x1e,
	0x22, 0xc4, 0x26, 0xbc, 0xd3, 0x26, 0xe5, 0xb9, 0x52, 0xf0, 0x59, 0xed, 0x47, 0xbf, 0x07, 0x64,
	0x41, 0x0b, 0xcc, 0x72, 0x60, 0x1a, 0xd4, 0xb1, 0x88, 0xa1, 0x7c, 0x90, 0xef, 0xb6, 0x49, 0x9c,
	0x7a, 0xc5, 0x81, 0x37, 0xf4, 0x8f, 0xf3, 0xcf, 0x80, 0x2c, 0x4e, 0x30, 0x4f, 0xd3, 0x70, 0xa3,
	0x4d, 0xea, 0xf3, 0x67, 0xd4, 0xd3, 0x74, 0x92, 0xfb, 0x48, 0x26, 0x22, 0xfd, 0x18, 0xde, 0x6b,
	0xaf, 0xfb, 0x81, 0x33, 0x74, 0x33, 0xaf, 0x96, 0x2e, 0x72, 0x23, 0xc2, 0xcd, 0x56, 0xcd, 0xbc,
	0xd2, 0xed, 0xc0, 0xaa, 0xd1, 0x4d, 0xb2, 0x5c, 0xa0, 0x1d, 0xe8, 0x76, 0xf0, 0xfb, 0x59, 0x12,
	0x8f, 0xec, 0xdc, 0xd7, 0x63, 0x1d, 0x6e, 0xb9, 0xa1, 0xbf, 0x54, 0x06, 0xb6, 0x7d, 0x7d, 0x67,
	0x94, 0xf4, 0x6d, 0xb5, 0x89, 0x56, 0xd7, 0x3d, 0x41, 0xef, 0x9f, 0x42, 0xcb, 0x06, 0xd5, 0xe8,
	0x0d, 0x32, 0x57, 0xa1, 0x27, 0xc8, 0x03, 0x87, 0xcc, 0x96, 0x85, 0x3a, 0xbb, 0x47, 0x3a, 0xdc,
	0x18, 0xc5, 0x72, 0xa1, 0x0d, 0x53, 0x60, 0x94, 0x80, 0x63, 0x60, 0x29, 0x17, 0x79, 0xa1, 0xa0,
	0x44, 0x1f, 0x3a, 0xf4, 0x9a, 0xcd, 0xed, 0x0b, 0x6d, 0xfa, 0x65, 0x6a, 0xcf, 0x87, 0xea, 0x33,
	0x6d, 0x98, 0x69, 0xc0, 0xa4, 0xc2, 0xc3, 0x47, 0xfe, 0x4c, 0x5b, 0x18, 0x00, 0x26, 0x25, 0x40,
	0x6f, 0x11, 0x2a, 0xd0, 0x80, 0x42, 0x9e, 0x33, 0x50, 0xaa, 0x3c, 0xe5, 0xb1, 0xff, 0xb7, 0xad,
	0x2a, 0xbb, 0x4a, 0xf9, 0x9d, 0x9f, 0x93, 0x6e, 0xd5, 0x80, 0xb1, 0x92, 0xa9, 0xc8, 0x81, 0x8d,
	0x0b, 0x3d, 0xfc, 0xc3, 0xf1, 0x89, 0xa3, 0xaf, 0x97, 0xc9, 0x43, 0x1f, 0x3c, 0x2c, 0xf4, 0xf0,
	0x94, 0x65, 0x87, 0x4c, 0xa3, 0x6c, 0x34, 0x65, 0xdb, 0x51, 0x04, 0x65, 0xdd, 0x8f, 0x1e, 0x59,
	0x42, 0xc9, 0xaa, 0x49, 0x9a, 0xca, 0x02, 0xab, 0xec, 0x53, 0x97, 0x9d, 0x47, 0x39, 0xf0, 0xc5,
	0x3d, 0x5b, 0xf3, 0xd0, 0x2a, 0x59, 0x6c, 0x40, 0x63, 0x80, 0xea, 0x4e, 0x3b, 0x8e, 0xa1, 0x35,
	0x73, 0x08, 0xe0, 0x6f, 0x75, 0x74, 0xde, 0xbd, 0x39, 0xf5, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x2d, 0xe1, 0x8e, 0x27, 0x56, 0x09, 0x00, 0x00,
}