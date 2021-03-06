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
// source: grid_svr_show_client.proto

package cisco_ios_xr_fretta_grid_svr_oper_grid_nodes_node_clients_client

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

type GridSvrShowClient_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ClientName           string   `protobuf:"bytes,2,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GridSvrShowClient_KEYS) Reset()         { *m = GridSvrShowClient_KEYS{} }
func (m *GridSvrShowClient_KEYS) String() string { return proto.CompactTextString(m) }
func (*GridSvrShowClient_KEYS) ProtoMessage()    {}
func (*GridSvrShowClient_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b2539926a329392, []int{0}
}

func (m *GridSvrShowClient_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GridSvrShowClient_KEYS.Unmarshal(m, b)
}
func (m *GridSvrShowClient_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GridSvrShowClient_KEYS.Marshal(b, m, deterministic)
}
func (m *GridSvrShowClient_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GridSvrShowClient_KEYS.Merge(m, src)
}
func (m *GridSvrShowClient_KEYS) XXX_Size() int {
	return xxx_messageInfo_GridSvrShowClient_KEYS.Size(m)
}
func (m *GridSvrShowClient_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_GridSvrShowClient_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_GridSvrShowClient_KEYS proto.InternalMessageInfo

func (m *GridSvrShowClient_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *GridSvrShowClient_KEYS) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

type GridSvrShowClientData struct {
	ResId                uint32   `protobuf:"varint,1,opt,name=res_id,json=resId,proto3" json:"res_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GridSvrShowClientData) Reset()         { *m = GridSvrShowClientData{} }
func (m *GridSvrShowClientData) String() string { return proto.CompactTextString(m) }
func (*GridSvrShowClientData) ProtoMessage()    {}
func (*GridSvrShowClientData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b2539926a329392, []int{1}
}

func (m *GridSvrShowClientData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GridSvrShowClientData.Unmarshal(m, b)
}
func (m *GridSvrShowClientData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GridSvrShowClientData.Marshal(b, m, deterministic)
}
func (m *GridSvrShowClientData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GridSvrShowClientData.Merge(m, src)
}
func (m *GridSvrShowClientData) XXX_Size() int {
	return xxx_messageInfo_GridSvrShowClientData.Size(m)
}
func (m *GridSvrShowClientData) XXX_DiscardUnknown() {
	xxx_messageInfo_GridSvrShowClientData.DiscardUnknown(m)
}

var xxx_messageInfo_GridSvrShowClientData proto.InternalMessageInfo

func (m *GridSvrShowClientData) GetResId() uint32 {
	if m != nil {
		return m.ResId
	}
	return 0
}

type GridSvrShowClient struct {
	ClientData           []*GridSvrShowClientData `protobuf:"bytes,50,rep,name=client_data,json=clientData,proto3" json:"client_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GridSvrShowClient) Reset()         { *m = GridSvrShowClient{} }
func (m *GridSvrShowClient) String() string { return proto.CompactTextString(m) }
func (*GridSvrShowClient) ProtoMessage()    {}
func (*GridSvrShowClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b2539926a329392, []int{2}
}

func (m *GridSvrShowClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GridSvrShowClient.Unmarshal(m, b)
}
func (m *GridSvrShowClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GridSvrShowClient.Marshal(b, m, deterministic)
}
func (m *GridSvrShowClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GridSvrShowClient.Merge(m, src)
}
func (m *GridSvrShowClient) XXX_Size() int {
	return xxx_messageInfo_GridSvrShowClient.Size(m)
}
func (m *GridSvrShowClient) XXX_DiscardUnknown() {
	xxx_messageInfo_GridSvrShowClient.DiscardUnknown(m)
}

var xxx_messageInfo_GridSvrShowClient proto.InternalMessageInfo

func (m *GridSvrShowClient) GetClientData() []*GridSvrShowClientData {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func init() {
	proto.RegisterType((*GridSvrShowClient_KEYS)(nil), "cisco_ios_xr_fretta_grid_svr_oper.grid.nodes.node.clients.client.grid_svr_show_client_KEYS")
	proto.RegisterType((*GridSvrShowClientData)(nil), "cisco_ios_xr_fretta_grid_svr_oper.grid.nodes.node.clients.client.grid_svr_show_client_data")
	proto.RegisterType((*GridSvrShowClient)(nil), "cisco_ios_xr_fretta_grid_svr_oper.grid.nodes.node.clients.client.grid_svr_show_client")
}

func init() { proto.RegisterFile("grid_svr_show_client.proto", fileDescriptor_9b2539926a329392) }

var fileDescriptor_9b2539926a329392 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2f, 0xca, 0x4c,
	0x89, 0x2f, 0x2e, 0x2b, 0x8a, 0x2f, 0xce, 0xc8, 0x2f, 0x8f, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x72, 0x48, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc,
	0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x4f, 0x2b, 0x4a, 0x2d, 0x29, 0x49, 0x8c, 0x87, 0xab, 0xcf, 0x2f,
	0x48, 0x2d, 0xd2, 0x03, 0xf1, 0xf4, 0xf2, 0xf2, 0x53, 0x52, 0x8b, 0xc1, 0xa4, 0x1e, 0x44, 0x7f,
	0x31, 0x94, 0x56, 0x8a, 0xe4, 0x92, 0xc4, 0x66, 0x7e, 0xbc, 0xb7, 0x6b, 0x64, 0xb0, 0x90, 0x34,
	0x17, 0x27, 0x48, 0x4f, 0x7c, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x07, 0x48, 0xc0, 0x2f, 0x31, 0x37, 0x55, 0x48, 0x9e, 0x8b, 0x1b, 0xaa, 0x16, 0x2c, 0xcd, 0x04,
	0x96, 0xe6, 0x82, 0x08, 0x81, 0x14, 0x28, 0x19, 0xe1, 0x30, 0x3a, 0x25, 0xb1, 0x24, 0x51, 0x48,
	0x94, 0x8b, 0xad, 0x28, 0xb5, 0x38, 0x3e, 0x33, 0x05, 0x6c, 0x2e, 0x6f, 0x10, 0x6b, 0x51, 0x6a,
	0xb1, 0x67, 0x8a, 0xd2, 0x14, 0x46, 0x2e, 0x11, 0x6c, 0x9a, 0x84, 0x6a, 0xe0, 0xb6, 0x81, 0xb4,
	0x4b, 0x18, 0x29, 0x30, 0x6b, 0x70, 0x1b, 0x45, 0xeb, 0x51, 0xea, 0x7f, 0x3d, 0x9c, 0x2e, 0x84,
	0x79, 0xc5, 0x25, 0xb1, 0x24, 0x31, 0x89, 0x0d, 0x1c, 0xdc, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xbb, 0xbe, 0x4e, 0x44, 0x8c, 0x01, 0x00, 0x00,
}
