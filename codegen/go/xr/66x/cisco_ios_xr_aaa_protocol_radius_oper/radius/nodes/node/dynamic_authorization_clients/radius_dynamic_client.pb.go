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
// source: radius_dynamic_client.proto

package cisco_ios_xr_aaa_protocol_radius_oper_radius_nodes_node_dynamic_authorization_clients

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

type RadiusDynamicClient_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RadiusDynamicClient_KEYS) Reset()         { *m = RadiusDynamicClient_KEYS{} }
func (m *RadiusDynamicClient_KEYS) String() string { return proto.CompactTextString(m) }
func (*RadiusDynamicClient_KEYS) ProtoMessage()    {}
func (*RadiusDynamicClient_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4914d387ba6a101, []int{0}
}

func (m *RadiusDynamicClient_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusDynamicClient_KEYS.Unmarshal(m, b)
}
func (m *RadiusDynamicClient_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusDynamicClient_KEYS.Marshal(b, m, deterministic)
}
func (m *RadiusDynamicClient_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusDynamicClient_KEYS.Merge(m, src)
}
func (m *RadiusDynamicClient_KEYS) XXX_Size() int {
	return xxx_messageInfo_RadiusDynamicClient_KEYS.Size(m)
}
func (m *RadiusDynamicClient_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusDynamicClient_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusDynamicClient_KEYS proto.InternalMessageInfo

func (m *RadiusDynamicClient_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type RadiusDynamicClientData struct {
	DiscReqs             uint32   `protobuf:"varint,1,opt,name=disc_reqs,json=discReqs,proto3" json:"disc_reqs,omitempty"`
	DiscAcks             uint32   `protobuf:"varint,2,opt,name=disc_acks,json=discAcks,proto3" json:"disc_acks,omitempty"`
	DiscNaks             uint32   `protobuf:"varint,3,opt,name=disc_naks,json=discNaks,proto3" json:"disc_naks,omitempty"`
	DiscBadAuth          uint32   `protobuf:"varint,4,opt,name=disc_bad_auth,json=discBadAuth,proto3" json:"disc_bad_auth,omitempty"`
	DropDiscReqs         uint32   `protobuf:"varint,5,opt,name=drop_disc_reqs,json=dropDiscReqs,proto3" json:"drop_disc_reqs,omitempty"`
	CoaReqs              uint32   `protobuf:"varint,6,opt,name=coa_reqs,json=coaReqs,proto3" json:"coa_reqs,omitempty"`
	CoaAcks              uint32   `protobuf:"varint,7,opt,name=coa_acks,json=coaAcks,proto3" json:"coa_acks,omitempty"`
	CoaNaks              uint32   `protobuf:"varint,8,opt,name=coa_naks,json=coaNaks,proto3" json:"coa_naks,omitempty"`
	CoaBadAuth           uint32   `protobuf:"varint,9,opt,name=coa_bad_auth,json=coaBadAuth,proto3" json:"coa_bad_auth,omitempty"`
	DropCoaReqs          uint32   `protobuf:"varint,10,opt,name=drop_coa_reqs,json=dropCoaReqs,proto3" json:"drop_coa_reqs,omitempty"`
	UnknownTypes         uint32   `protobuf:"varint,11,opt,name=unknown_types,json=unknownTypes,proto3" json:"unknown_types,omitempty"`
	InternalError        uint32   `protobuf:"varint,12,opt,name=internal_error,json=internalError,proto3" json:"internal_error,omitempty"`
	PakDecodeFail        uint32   `protobuf:"varint,13,opt,name=pak_decode_fail,json=pakDecodeFail,proto3" json:"pak_decode_fail,omitempty"`
	VrfParseFailErr      uint32   `protobuf:"varint,14,opt,name=vrf_parse_fail_err,json=vrfParseFailErr,proto3" json:"vrf_parse_fail_err,omitempty"`
	UnknownVsaError      uint32   `protobuf:"varint,15,opt,name=unknown_vsa_error,json=unknownVsaError,proto3" json:"unknown_vsa_error,omitempty"`
	SendMsgFailed        uint32   `protobuf:"varint,16,opt,name=send_msg_failed,json=sendMsgFailed,proto3" json:"send_msg_failed,omitempty"`
	RadiusToCh           uint32   `protobuf:"varint,17,opt,name=radius_to_ch,json=radiusToCh,proto3" json:"radius_to_ch,omitempty"`
	ChToRadius           uint32   `protobuf:"varint,18,opt,name=ch_to_radius,json=chToRadius,proto3" json:"ch_to_radius,omitempty"`
	ServiceParseFail     uint32   `protobuf:"varint,19,opt,name=service_parse_fail,json=serviceParseFail,proto3" json:"service_parse_fail,omitempty"`
	MultiSubsError       uint32   `protobuf:"varint,20,opt,name=multi_subs_error,json=multiSubsError,proto3" json:"multi_subs_error,omitempty"`
	ServiceNotPresent    uint32   `protobuf:"varint,21,opt,name=service_not_present,json=serviceNotPresent,proto3" json:"service_not_present,omitempty"`
	SendToChFail         uint32   `protobuf:"varint,22,opt,name=send_to_ch_fail,json=sendToChFail,proto3" json:"send_to_ch_fail,omitempty"`
	VrfName              string   `protobuf:"bytes,23,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	ClientAddress        string   `protobuf:"bytes,24,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RadiusDynamicClientData) Reset()         { *m = RadiusDynamicClientData{} }
func (m *RadiusDynamicClientData) String() string { return proto.CompactTextString(m) }
func (*RadiusDynamicClientData) ProtoMessage()    {}
func (*RadiusDynamicClientData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4914d387ba6a101, []int{1}
}

func (m *RadiusDynamicClientData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusDynamicClientData.Unmarshal(m, b)
}
func (m *RadiusDynamicClientData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusDynamicClientData.Marshal(b, m, deterministic)
}
func (m *RadiusDynamicClientData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusDynamicClientData.Merge(m, src)
}
func (m *RadiusDynamicClientData) XXX_Size() int {
	return xxx_messageInfo_RadiusDynamicClientData.Size(m)
}
func (m *RadiusDynamicClientData) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusDynamicClientData.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusDynamicClientData proto.InternalMessageInfo

func (m *RadiusDynamicClientData) GetDiscReqs() uint32 {
	if m != nil {
		return m.DiscReqs
	}
	return 0
}

func (m *RadiusDynamicClientData) GetDiscAcks() uint32 {
	if m != nil {
		return m.DiscAcks
	}
	return 0
}

func (m *RadiusDynamicClientData) GetDiscNaks() uint32 {
	if m != nil {
		return m.DiscNaks
	}
	return 0
}

func (m *RadiusDynamicClientData) GetDiscBadAuth() uint32 {
	if m != nil {
		return m.DiscBadAuth
	}
	return 0
}

func (m *RadiusDynamicClientData) GetDropDiscReqs() uint32 {
	if m != nil {
		return m.DropDiscReqs
	}
	return 0
}

func (m *RadiusDynamicClientData) GetCoaReqs() uint32 {
	if m != nil {
		return m.CoaReqs
	}
	return 0
}

func (m *RadiusDynamicClientData) GetCoaAcks() uint32 {
	if m != nil {
		return m.CoaAcks
	}
	return 0
}

func (m *RadiusDynamicClientData) GetCoaNaks() uint32 {
	if m != nil {
		return m.CoaNaks
	}
	return 0
}

func (m *RadiusDynamicClientData) GetCoaBadAuth() uint32 {
	if m != nil {
		return m.CoaBadAuth
	}
	return 0
}

func (m *RadiusDynamicClientData) GetDropCoaReqs() uint32 {
	if m != nil {
		return m.DropCoaReqs
	}
	return 0
}

func (m *RadiusDynamicClientData) GetUnknownTypes() uint32 {
	if m != nil {
		return m.UnknownTypes
	}
	return 0
}

func (m *RadiusDynamicClientData) GetInternalError() uint32 {
	if m != nil {
		return m.InternalError
	}
	return 0
}

func (m *RadiusDynamicClientData) GetPakDecodeFail() uint32 {
	if m != nil {
		return m.PakDecodeFail
	}
	return 0
}

func (m *RadiusDynamicClientData) GetVrfParseFailErr() uint32 {
	if m != nil {
		return m.VrfParseFailErr
	}
	return 0
}

func (m *RadiusDynamicClientData) GetUnknownVsaError() uint32 {
	if m != nil {
		return m.UnknownVsaError
	}
	return 0
}

func (m *RadiusDynamicClientData) GetSendMsgFailed() uint32 {
	if m != nil {
		return m.SendMsgFailed
	}
	return 0
}

func (m *RadiusDynamicClientData) GetRadiusToCh() uint32 {
	if m != nil {
		return m.RadiusToCh
	}
	return 0
}

func (m *RadiusDynamicClientData) GetChToRadius() uint32 {
	if m != nil {
		return m.ChToRadius
	}
	return 0
}

func (m *RadiusDynamicClientData) GetServiceParseFail() uint32 {
	if m != nil {
		return m.ServiceParseFail
	}
	return 0
}

func (m *RadiusDynamicClientData) GetMultiSubsError() uint32 {
	if m != nil {
		return m.MultiSubsError
	}
	return 0
}

func (m *RadiusDynamicClientData) GetServiceNotPresent() uint32 {
	if m != nil {
		return m.ServiceNotPresent
	}
	return 0
}

func (m *RadiusDynamicClientData) GetSendToChFail() uint32 {
	if m != nil {
		return m.SendToChFail
	}
	return 0
}

func (m *RadiusDynamicClientData) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *RadiusDynamicClientData) GetClientAddress() string {
	if m != nil {
		return m.ClientAddress
	}
	return ""
}

type RadiusDynamicClient struct {
	DynamicAuthorClient  []*RadiusDynamicClientData `protobuf:"bytes,50,rep,name=dynamic_author_client,json=dynamicAuthorClient,proto3" json:"dynamic_author_client,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RadiusDynamicClient) Reset()         { *m = RadiusDynamicClient{} }
func (m *RadiusDynamicClient) String() string { return proto.CompactTextString(m) }
func (*RadiusDynamicClient) ProtoMessage()    {}
func (*RadiusDynamicClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4914d387ba6a101, []int{2}
}

func (m *RadiusDynamicClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusDynamicClient.Unmarshal(m, b)
}
func (m *RadiusDynamicClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusDynamicClient.Marshal(b, m, deterministic)
}
func (m *RadiusDynamicClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusDynamicClient.Merge(m, src)
}
func (m *RadiusDynamicClient) XXX_Size() int {
	return xxx_messageInfo_RadiusDynamicClient.Size(m)
}
func (m *RadiusDynamicClient) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusDynamicClient.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusDynamicClient proto.InternalMessageInfo

func (m *RadiusDynamicClient) GetDynamicAuthorClient() []*RadiusDynamicClientData {
	if m != nil {
		return m.DynamicAuthorClient
	}
	return nil
}

func init() {
	proto.RegisterType((*RadiusDynamicClient_KEYS)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.dynamic_authorization_clients.radius_dynamic_client_KEYS")
	proto.RegisterType((*RadiusDynamicClientData)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.dynamic_authorization_clients.radius_dynamic_client_data")
	proto.RegisterType((*RadiusDynamicClient)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.dynamic_authorization_clients.radius_dynamic_client")
}

func init() { proto.RegisterFile("radius_dynamic_client.proto", fileDescriptor_a4914d387ba6a101) }

var fileDescriptor_a4914d387ba6a101 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x4e, 0x54, 0x31,
	0x14, 0xc6, 0x33, 0xa2, 0x30, 0x14, 0x66, 0x80, 0x22, 0x5a, 0x64, 0x43, 0x46, 0x31, 0x44, 0xcd,
	0x5d, 0xe0, 0xca, 0xe5, 0xc8, 0x9f, 0x8d, 0x91, 0x90, 0x01, 0x4d, 0x5c, 0x35, 0x67, 0xda, 0x0e,
	0xd3, 0xdc, 0x99, 0xf6, 0xd2, 0xf6, 0x8e, 0xe2, 0xfb, 0xf8, 0x1c, 0xbe, 0x99, 0x31, 0x3d, 0xed,
	0xdc, 0xc1, 0x04, 0x96, 0x6e, 0x08, 0xf9, 0x7e, 0x5f, 0xcf, 0xf9, 0x4e, 0x3b, 0xe7, 0x92, 0x3d,
	0x07, 0x52, 0xd7, 0x9e, 0xcb, 0x5b, 0x03, 0x53, 0x2d, 0xb8, 0x98, 0x68, 0x65, 0x42, 0x51, 0x39,
	0x1b, 0x2c, 0xfd, 0x22, 0xb4, 0x17, 0x96, 0x6b, 0xeb, 0xf9, 0x0f, 0xc7, 0x01, 0x80, 0xa3, 0x2e,
	0xec, 0x84, 0xe7, 0x63, 0xb6, 0x52, 0xae, 0x48, 0xff, 0x17, 0xc6, 0x4a, 0x95, 0xfe, 0x16, 0xf3,
	0x6a, 0x50, 0x87, 0xb1, 0x75, 0xfa, 0x27, 0x04, 0x6d, 0x4d, 0xae, 0xed, 0x7b, 0x1f, 0xc8, 0x8b,
	0x7b, 0xbb, 0xf2, 0x4f, 0xa7, 0xdf, 0x2e, 0xe9, 0x1e, 0x59, 0x8d, 0x35, 0xb8, 0x81, 0xa9, 0x62,
	0xad, 0xfd, 0xd6, 0xe1, 0xea, 0xa0, 0x1d, 0x85, 0x73, 0x98, 0xaa, 0xde, 0x9f, 0xe5, 0x87, 0xce,
	0x4a, 0x08, 0x10, 0xcf, 0x4a, 0xed, 0x05, 0x77, 0xea, 0xc6, 0xe3, 0xd9, 0xce, 0xa0, 0x1d, 0x85,
	0x81, 0xba, 0xf1, 0x0d, 0x04, 0x51, 0x7a, 0xf6, 0x68, 0x01, 0xfb, 0xa2, 0x5c, 0x40, 0x03, 0xa5,
	0x67, 0x4b, 0x0b, 0x78, 0x0e, 0xa5, 0xa7, 0x3d, 0xd2, 0x41, 0x38, 0x04, 0x89, 0x23, 0xb1, 0xc7,
	0x68, 0x58, 0x8b, 0xe2, 0x47, 0x90, 0xfd, 0x3a, 0x8c, 0xe9, 0x2b, 0xd2, 0x95, 0xce, 0x56, 0x7c,
	0xd1, 0xff, 0x09, 0x9a, 0xd6, 0xa3, 0x7a, 0x32, 0xcf, 0xb0, 0x4b, 0xda, 0xc2, 0x42, 0xe2, 0xcb,
	0xc8, 0x57, 0x84, 0x85, 0xbb, 0x08, 0xd3, 0xad, 0x34, 0x08, 0xc3, 0x65, 0x84, 0xd9, 0xda, 0x0d,
	0xc2, 0x68, 0xfb, 0x64, 0x3d, 0xa2, 0x26, 0xd9, 0x2a, 0x62, 0x22, 0x2c, 0xcc, 0x83, 0xc5, 0xf0,
	0x31, 0x58, 0xd3, 0x97, 0xe4, 0xf0, 0xce, 0x56, 0xc7, 0xb9, 0xf7, 0x4b, 0xd2, 0xa9, 0x4d, 0x69,
	0xec, 0x77, 0xc3, 0xc3, 0x6d, 0xa5, 0x3c, 0x5b, 0x4b, 0xd9, 0xb3, 0x78, 0x15, 0x35, 0x7a, 0x40,
	0xba, 0xda, 0x04, 0xe5, 0x0c, 0x4c, 0xb8, 0x72, 0xce, 0x3a, 0xb6, 0x8e, 0xae, 0xce, 0x5c, 0x3d,
	0x8d, 0x22, 0x7d, 0x4d, 0x36, 0x2a, 0x28, 0xb9, 0x54, 0x22, 0xbe, 0xe2, 0x08, 0xf4, 0x84, 0x75,
	0x92, 0xaf, 0x82, 0xf2, 0x04, 0xd5, 0x33, 0xd0, 0x13, 0xfa, 0x96, 0xd0, 0x99, 0x1b, 0xf1, 0x0a,
	0x9c, 0x4f, 0xb6, 0x58, 0x94, 0x75, 0xd1, 0xba, 0x31, 0x73, 0xa3, 0x8b, 0x08, 0xa2, 0xf3, 0xd4,
	0x39, 0xfa, 0x86, 0x6c, 0xcd, 0x03, 0xce, 0x3c, 0xe4, 0xf6, 0x1b, 0xc9, 0x9b, 0xc1, 0x57, 0x0f,
	0x4d, 0x00, 0xaf, 0x8c, 0xe4, 0x53, 0x7f, 0x8d, 0x75, 0x95, 0x64, 0x9b, 0x29, 0x40, 0x94, 0x3f,
	0xfb, 0xeb, 0x33, 0x14, 0xe3, 0xd5, 0xe5, 0x9f, 0x52, 0xb0, 0x5c, 0x8c, 0xd9, 0x56, 0xba, 0xba,
	0xa4, 0x5d, 0xd9, 0xe3, 0x31, 0x5e, 0xee, 0x38, 0xd2, 0xa4, 0x31, 0x9a, 0x2f, 0x77, 0x7c, 0x65,
	0x07, 0xa8, 0xd0, 0x77, 0x84, 0x7a, 0xe5, 0x66, 0x5a, 0xa8, 0x3b, 0x83, 0xb0, 0x6d, 0xf4, 0x6d,
	0x66, 0xd2, 0x0c, 0x42, 0x0f, 0xc9, 0xe6, 0xb4, 0x9e, 0x04, 0xcd, 0x7d, 0x3d, 0xf4, 0x79, 0x88,
	0xa7, 0xe8, 0xed, 0xa2, 0x7e, 0x59, 0x0f, 0x7d, 0x9a, 0xa1, 0x20, 0xdb, 0xf3, 0xba, 0xc6, 0x06,
	0x5e, 0x39, 0xe5, 0x95, 0x09, 0x6c, 0x07, 0xcd, 0x5b, 0x19, 0x9d, 0xdb, 0x70, 0x91, 0x00, 0x3d,
	0xc8, 0x33, 0xe3, 0x24, 0x29, 0xc4, 0xb3, 0xf4, 0x84, 0x51, 0x8e, 0xc3, 0x60, 0x80, 0x5d, 0xd2,
	0x8e, 0x77, 0x8e, 0xab, 0xf5, 0x1c, 0x57, 0x6b, 0x65, 0xe6, 0x46, 0x71, 0xb3, 0xe2, 0xeb, 0xe6,
	0x4d, 0x02, 0x29, 0x9d, 0xf2, 0x9e, 0x31, 0x34, 0x74, 0x92, 0xda, 0x4f, 0x62, 0xef, 0x77, 0x8b,
	0xec, 0xdc, 0xbb, 0x80, 0xf4, 0x57, 0x8b, 0xec, 0xfc, 0xbb, 0xf7, 0x99, 0xb0, 0xa3, 0xfd, 0xa5,
	0xc3, 0xb5, 0xa3, 0x9b, 0xe2, 0xbf, 0x7c, 0x4d, 0x8a, 0x87, 0x3f, 0x07, 0x83, 0xed, 0x2c, 0xf6,
	0xf1, 0xe0, 0x31, 0x92, 0xe1, 0x32, 0x76, 0x7d, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x81,
	0xbd, 0x28, 0xfa, 0x04, 0x00, 0x00,
}