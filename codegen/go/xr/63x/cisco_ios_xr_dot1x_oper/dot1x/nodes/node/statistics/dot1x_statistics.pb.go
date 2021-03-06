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
// source: dot1x_statistics.proto

package cisco_ios_xr_dot1x_oper_dot1x_nodes_node_statistics

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

type Dot1XStatistics_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XStatistics_KEYS) Reset()         { *m = Dot1XStatistics_KEYS{} }
func (m *Dot1XStatistics_KEYS) String() string { return proto.CompactTextString(m) }
func (*Dot1XStatistics_KEYS) ProtoMessage()    {}
func (*Dot1XStatistics_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b69da07a7fbf3f6, []int{0}
}

func (m *Dot1XStatistics_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XStatistics_KEYS.Unmarshal(m, b)
}
func (m *Dot1XStatistics_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XStatistics_KEYS.Marshal(b, m, deterministic)
}
func (m *Dot1XStatistics_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XStatistics_KEYS.Merge(m, src)
}
func (m *Dot1XStatistics_KEYS) XXX_Size() int {
	return xxx_messageInfo_Dot1XStatistics_KEYS.Size(m)
}
func (m *Dot1XStatistics_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XStatistics_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XStatistics_KEYS proto.InternalMessageInfo

func (m *Dot1XStatistics_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Dot1XGlobalStatistics struct {
	TxTotal              uint32   `protobuf:"varint,1,opt,name=tx_total,json=txTotal,proto3" json:"tx_total,omitempty"`
	RxTotal              uint32   `protobuf:"varint,2,opt,name=rx_total,json=rxTotal,proto3" json:"rx_total,omitempty"`
	RxNoIdb              uint32   `protobuf:"varint,3,opt,name=rx_no_idb,json=rxNoIdb,proto3" json:"rx_no_idb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XGlobalStatistics) Reset()         { *m = Dot1XGlobalStatistics{} }
func (m *Dot1XGlobalStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XGlobalStatistics) ProtoMessage()    {}
func (*Dot1XGlobalStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b69da07a7fbf3f6, []int{1}
}

func (m *Dot1XGlobalStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XGlobalStatistics.Unmarshal(m, b)
}
func (m *Dot1XGlobalStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XGlobalStatistics.Marshal(b, m, deterministic)
}
func (m *Dot1XGlobalStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XGlobalStatistics.Merge(m, src)
}
func (m *Dot1XGlobalStatistics) XXX_Size() int {
	return xxx_messageInfo_Dot1XGlobalStatistics.Size(m)
}
func (m *Dot1XGlobalStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XGlobalStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XGlobalStatistics proto.InternalMessageInfo

func (m *Dot1XGlobalStatistics) GetTxTotal() uint32 {
	if m != nil {
		return m.TxTotal
	}
	return 0
}

func (m *Dot1XGlobalStatistics) GetRxTotal() uint32 {
	if m != nil {
		return m.RxTotal
	}
	return 0
}

func (m *Dot1XGlobalStatistics) GetRxNoIdb() uint32 {
	if m != nil {
		return m.RxNoIdb
	}
	return 0
}

type Dot1XIdbStatistics struct {
	RxTotal              uint32   `protobuf:"varint,1,opt,name=rx_total,json=rxTotal,proto3" json:"rx_total,omitempty"`
	TxTotal              uint32   `protobuf:"varint,2,opt,name=tx_total,json=txTotal,proto3" json:"tx_total,omitempty"`
	NoRxOnIntfDown       uint32   `protobuf:"varint,3,opt,name=no_rx_on_intf_down,json=noRxOnIntfDown,proto3" json:"no_rx_on_intf_down,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XIdbStatistics) Reset()         { *m = Dot1XIdbStatistics{} }
func (m *Dot1XIdbStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XIdbStatistics) ProtoMessage()    {}
func (*Dot1XIdbStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b69da07a7fbf3f6, []int{2}
}

func (m *Dot1XIdbStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XIdbStatistics.Unmarshal(m, b)
}
func (m *Dot1XIdbStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XIdbStatistics.Marshal(b, m, deterministic)
}
func (m *Dot1XIdbStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XIdbStatistics.Merge(m, src)
}
func (m *Dot1XIdbStatistics) XXX_Size() int {
	return xxx_messageInfo_Dot1XIdbStatistics.Size(m)
}
func (m *Dot1XIdbStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XIdbStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XIdbStatistics proto.InternalMessageInfo

func (m *Dot1XIdbStatistics) GetRxTotal() uint32 {
	if m != nil {
		return m.RxTotal
	}
	return 0
}

func (m *Dot1XIdbStatistics) GetTxTotal() uint32 {
	if m != nil {
		return m.TxTotal
	}
	return 0
}

func (m *Dot1XIdbStatistics) GetNoRxOnIntfDown() uint32 {
	if m != nil {
		return m.NoRxOnIntfDown
	}
	return 0
}

type Dot1XIfAuthStatistics struct {
	RxStart              uint32   `protobuf:"varint,1,opt,name=rx_start,json=rxStart,proto3" json:"rx_start,omitempty"`
	RxLogoff             uint32   `protobuf:"varint,2,opt,name=rx_logoff,json=rxLogoff,proto3" json:"rx_logoff,omitempty"`
	RxResp               uint32   `protobuf:"varint,3,opt,name=rx_resp,json=rxResp,proto3" json:"rx_resp,omitempty"`
	RxRespId             uint32   `protobuf:"varint,4,opt,name=rx_resp_id,json=rxRespId,proto3" json:"rx_resp_id,omitempty"`
	RxInvalid            uint32   `protobuf:"varint,5,opt,name=rx_invalid,json=rxInvalid,proto3" json:"rx_invalid,omitempty"`
	RxLenErr             uint32   `protobuf:"varint,6,opt,name=rx_len_err,json=rxLenErr,proto3" json:"rx_len_err,omitempty"`
	RxMyMacErr           uint32   `protobuf:"varint,7,opt,name=rx_my_mac_err,json=rxMyMacErr,proto3" json:"rx_my_mac_err,omitempty"`
	RxTotal              uint32   `protobuf:"varint,8,opt,name=rx_total,json=rxTotal,proto3" json:"rx_total,omitempty"`
	TxReq                uint32   `protobuf:"varint,9,opt,name=tx_req,json=txReq,proto3" json:"tx_req,omitempty"`
	TxReqid              uint32   `protobuf:"varint,10,opt,name=tx_reqid,json=txReqid,proto3" json:"tx_reqid,omitempty"`
	TxTotal              uint32   `protobuf:"varint,11,opt,name=tx_total,json=txTotal,proto3" json:"tx_total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XIfAuthStatistics) Reset()         { *m = Dot1XIfAuthStatistics{} }
func (m *Dot1XIfAuthStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XIfAuthStatistics) ProtoMessage()    {}
func (*Dot1XIfAuthStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b69da07a7fbf3f6, []int{3}
}

func (m *Dot1XIfAuthStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XIfAuthStatistics.Unmarshal(m, b)
}
func (m *Dot1XIfAuthStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XIfAuthStatistics.Marshal(b, m, deterministic)
}
func (m *Dot1XIfAuthStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XIfAuthStatistics.Merge(m, src)
}
func (m *Dot1XIfAuthStatistics) XXX_Size() int {
	return xxx_messageInfo_Dot1XIfAuthStatistics.Size(m)
}
func (m *Dot1XIfAuthStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XIfAuthStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XIfAuthStatistics proto.InternalMessageInfo

func (m *Dot1XIfAuthStatistics) GetRxStart() uint32 {
	if m != nil {
		return m.RxStart
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetRxLogoff() uint32 {
	if m != nil {
		return m.RxLogoff
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetRxResp() uint32 {
	if m != nil {
		return m.RxResp
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetRxRespId() uint32 {
	if m != nil {
		return m.RxRespId
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetRxInvalid() uint32 {
	if m != nil {
		return m.RxInvalid
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetRxLenErr() uint32 {
	if m != nil {
		return m.RxLenErr
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetRxMyMacErr() uint32 {
	if m != nil {
		return m.RxMyMacErr
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetRxTotal() uint32 {
	if m != nil {
		return m.RxTotal
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetTxReq() uint32 {
	if m != nil {
		return m.TxReq
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetTxReqid() uint32 {
	if m != nil {
		return m.TxReqid
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetTxTotal() uint32 {
	if m != nil {
		return m.TxTotal
	}
	return 0
}

type Dot1XIfSuppStatistics struct {
	RxReq                uint32   `protobuf:"varint,1,opt,name=rx_req,json=rxReq,proto3" json:"rx_req,omitempty"`
	RxInvalid            uint32   `protobuf:"varint,2,opt,name=rx_invalid,json=rxInvalid,proto3" json:"rx_invalid,omitempty"`
	RxLenErr             uint32   `protobuf:"varint,3,opt,name=rx_len_err,json=rxLenErr,proto3" json:"rx_len_err,omitempty"`
	RxMyMacErr           uint32   `protobuf:"varint,4,opt,name=rx_my_mac_err,json=rxMyMacErr,proto3" json:"rx_my_mac_err,omitempty"`
	RxTotal              uint32   `protobuf:"varint,5,opt,name=rx_total,json=rxTotal,proto3" json:"rx_total,omitempty"`
	TxStart              uint32   `protobuf:"varint,6,opt,name=tx_start,json=txStart,proto3" json:"tx_start,omitempty"`
	TxLogoff             uint32   `protobuf:"varint,7,opt,name=tx_logoff,json=txLogoff,proto3" json:"tx_logoff,omitempty"`
	TxResp               uint32   `protobuf:"varint,8,opt,name=tx_resp,json=txResp,proto3" json:"tx_resp,omitempty"`
	TxTotal              uint32   `protobuf:"varint,9,opt,name=tx_total,json=txTotal,proto3" json:"tx_total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XIfSuppStatistics) Reset()         { *m = Dot1XIfSuppStatistics{} }
func (m *Dot1XIfSuppStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XIfSuppStatistics) ProtoMessage()    {}
func (*Dot1XIfSuppStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b69da07a7fbf3f6, []int{4}
}

func (m *Dot1XIfSuppStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XIfSuppStatistics.Unmarshal(m, b)
}
func (m *Dot1XIfSuppStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XIfSuppStatistics.Marshal(b, m, deterministic)
}
func (m *Dot1XIfSuppStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XIfSuppStatistics.Merge(m, src)
}
func (m *Dot1XIfSuppStatistics) XXX_Size() int {
	return xxx_messageInfo_Dot1XIfSuppStatistics.Size(m)
}
func (m *Dot1XIfSuppStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XIfSuppStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XIfSuppStatistics proto.InternalMessageInfo

func (m *Dot1XIfSuppStatistics) GetRxReq() uint32 {
	if m != nil {
		return m.RxReq
	}
	return 0
}

func (m *Dot1XIfSuppStatistics) GetRxInvalid() uint32 {
	if m != nil {
		return m.RxInvalid
	}
	return 0
}

func (m *Dot1XIfSuppStatistics) GetRxLenErr() uint32 {
	if m != nil {
		return m.RxLenErr
	}
	return 0
}

func (m *Dot1XIfSuppStatistics) GetRxMyMacErr() uint32 {
	if m != nil {
		return m.RxMyMacErr
	}
	return 0
}

func (m *Dot1XIfSuppStatistics) GetRxTotal() uint32 {
	if m != nil {
		return m.RxTotal
	}
	return 0
}

func (m *Dot1XIfSuppStatistics) GetTxStart() uint32 {
	if m != nil {
		return m.TxStart
	}
	return 0
}

func (m *Dot1XIfSuppStatistics) GetTxLogoff() uint32 {
	if m != nil {
		return m.TxLogoff
	}
	return 0
}

func (m *Dot1XIfSuppStatistics) GetTxResp() uint32 {
	if m != nil {
		return m.TxResp
	}
	return 0
}

func (m *Dot1XIfSuppStatistics) GetTxTotal() uint32 {
	if m != nil {
		return m.TxTotal
	}
	return 0
}

type Dot1XIfStatistics struct {
	InterfaceName        string                 `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Pae                  string                 `protobuf:"bytes,2,opt,name=pae,proto3" json:"pae,omitempty"`
	Idb                  *Dot1XIdbStatistics    `protobuf:"bytes,3,opt,name=idb,proto3" json:"idb,omitempty"`
	Auth                 *Dot1XIfAuthStatistics `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty"`
	Supp                 *Dot1XIfSuppStatistics `protobuf:"bytes,5,opt,name=supp,proto3" json:"supp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Dot1XIfStatistics) Reset()         { *m = Dot1XIfStatistics{} }
func (m *Dot1XIfStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XIfStatistics) ProtoMessage()    {}
func (*Dot1XIfStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b69da07a7fbf3f6, []int{5}
}

func (m *Dot1XIfStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XIfStatistics.Unmarshal(m, b)
}
func (m *Dot1XIfStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XIfStatistics.Marshal(b, m, deterministic)
}
func (m *Dot1XIfStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XIfStatistics.Merge(m, src)
}
func (m *Dot1XIfStatistics) XXX_Size() int {
	return xxx_messageInfo_Dot1XIfStatistics.Size(m)
}
func (m *Dot1XIfStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XIfStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XIfStatistics proto.InternalMessageInfo

func (m *Dot1XIfStatistics) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Dot1XIfStatistics) GetPae() string {
	if m != nil {
		return m.Pae
	}
	return ""
}

func (m *Dot1XIfStatistics) GetIdb() *Dot1XIdbStatistics {
	if m != nil {
		return m.Idb
	}
	return nil
}

func (m *Dot1XIfStatistics) GetAuth() *Dot1XIfAuthStatistics {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *Dot1XIfStatistics) GetSupp() *Dot1XIfSuppStatistics {
	if m != nil {
		return m.Supp
	}
	return nil
}

type Dot1XStatistics struct {
	GlStats              *Dot1XGlobalStatistics `protobuf:"bytes,50,opt,name=gl_stats,json=glStats,proto3" json:"gl_stats,omitempty"`
	IfStats              []*Dot1XIfStatistics   `protobuf:"bytes,51,rep,name=if_stats,json=ifStats,proto3" json:"if_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Dot1XStatistics) Reset()         { *m = Dot1XStatistics{} }
func (m *Dot1XStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XStatistics) ProtoMessage()    {}
func (*Dot1XStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b69da07a7fbf3f6, []int{6}
}

func (m *Dot1XStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XStatistics.Unmarshal(m, b)
}
func (m *Dot1XStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XStatistics.Marshal(b, m, deterministic)
}
func (m *Dot1XStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XStatistics.Merge(m, src)
}
func (m *Dot1XStatistics) XXX_Size() int {
	return xxx_messageInfo_Dot1XStatistics.Size(m)
}
func (m *Dot1XStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XStatistics proto.InternalMessageInfo

func (m *Dot1XStatistics) GetGlStats() *Dot1XGlobalStatistics {
	if m != nil {
		return m.GlStats
	}
	return nil
}

func (m *Dot1XStatistics) GetIfStats() []*Dot1XIfStatistics {
	if m != nil {
		return m.IfStats
	}
	return nil
}

func init() {
	proto.RegisterType((*Dot1XStatistics_KEYS)(nil), "cisco_ios_xr_dot1x_oper.dot1x.nodes.node.statistics.dot1x_statistics_KEYS")
	proto.RegisterType((*Dot1XGlobalStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.nodes.node.statistics.dot1x_global_statistics")
	proto.RegisterType((*Dot1XIdbStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.nodes.node.statistics.dot1x_idb_statistics")
	proto.RegisterType((*Dot1XIfAuthStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.nodes.node.statistics.dot1x_if_auth_statistics")
	proto.RegisterType((*Dot1XIfSuppStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.nodes.node.statistics.dot1x_if_supp_statistics")
	proto.RegisterType((*Dot1XIfStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.nodes.node.statistics.dot1x_if_statistics")
	proto.RegisterType((*Dot1XStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.nodes.node.statistics.dot1x_statistics")
}

func init() { proto.RegisterFile("dot1x_statistics.proto", fileDescriptor_0b69da07a7fbf3f6) }

var fileDescriptor_0b69da07a7fbf3f6 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xb5, 0x76, 0x4d, 0x93, 0x53, 0x6d, 0x9a, 0x0c, 0x63, 0xe1, 0x9f, 0x34, 0x2a, 0x21,
	0x4d, 0x5c, 0x54, 0xa2, 0xe3, 0x11, 0x98, 0x44, 0xc5, 0x36, 0xa4, 0x8c, 0x1b, 0xc4, 0x85, 0xe5,
	0x26, 0x4e, 0xb1, 0x48, 0xed, 0xec, 0xc4, 0x63, 0xde, 0x9b, 0xf0, 0x5a, 0xbc, 0x01, 0xaf, 0xc0,
	0x1b, 0x20, 0xdb, 0x49, 0x9b, 0xb6, 0xd3, 0x90, 0xe8, 0x4d, 0x15, 0xe7, 0xf3, 0x39, 0xbf, 0x9c,
	0xef, 0xf8, 0xb8, 0xf0, 0x24, 0x53, 0xfa, 0xad, 0xa1, 0x95, 0x66, 0x5a, 0x54, 0x5a, 0xa4, 0xd5,
	0xa8, 0x44, 0xa5, 0x15, 0x39, 0x4d, 0x45, 0x95, 0x2a, 0x2a, 0x54, 0x45, 0x0d, 0x52, 0xbf, 0x49,
	0x95, 0x1c, 0x47, 0xee, 0x71, 0x24, 0x55, 0xc6, 0x2b, 0xf7, 0x3b, 0x5a, 0x86, 0x0e, 0xdf, 0xc1,
	0xe1, 0x7a, 0x3a, 0xfa, 0xf1, 0xec, 0xcb, 0x15, 0x79, 0x0e, 0x91, 0xdd, 0x4b, 0x25, 0x9b, 0xf3,
	0x78, 0xe7, 0x78, 0xe7, 0x24, 0x4a, 0x42, 0xfb, 0xe2, 0x92, 0xcd, 0xf9, 0xf0, 0x3b, 0x1c, 0xf9,
	0xa8, 0x59, 0xa1, 0xa6, 0xac, 0x68, 0x05, 0x93, 0xa7, 0x10, 0x6a, 0x43, 0xb5, 0xd2, 0xac, 0x70,
	0x61, 0x7b, 0x49, 0x5f, 0x9b, 0xcf, 0x76, 0x69, 0x25, 0x6c, 0xa4, 0x8e, 0x97, 0xb0, 0x96, 0x9e,
	0x41, 0x84, 0x86, 0x4a, 0x45, 0x45, 0x36, 0x8d, 0xbb, 0x8d, 0x76, 0xa9, 0x26, 0xd9, 0x74, 0x68,
	0xe0, 0xb1, 0x87, 0x89, 0x6c, 0xba, 0x46, 0xc2, 0x35, 0x12, 0x2e, 0x49, 0x7a, 0x8d, 0xd4, 0x7c,
	0xc4, 0x1b, 0x20, 0x52, 0x51, 0x34, 0x54, 0x49, 0x2a, 0xa4, 0xce, 0x69, 0xa6, 0x6e, 0x65, 0x8d,
	0xdc, 0x97, 0x2a, 0x31, 0x9f, 0xe4, 0x44, 0xea, 0xfc, 0xbd, 0xba, 0x95, 0xc3, 0x5f, 0x1d, 0x88,
	0x6b, 0x74, 0x4e, 0xd9, 0x8d, 0xfe, 0xb6, 0x89, 0xaf, 0x34, 0x43, 0xbd, 0xc4, 0x5f, 0xd9, 0xa5,
	0xf5, 0x0e, 0x0d, 0x2d, 0xd4, 0x4c, 0xe5, 0x79, 0xcd, 0x0f, 0xd1, 0x9c, 0xbb, 0x35, 0x39, 0x82,
	0x3e, 0x1a, 0x8a, 0xbc, 0x2a, 0x6b, 0x6a, 0x80, 0x26, 0xe1, 0x55, 0x49, 0x5e, 0x00, 0xd4, 0x02,
	0x15, 0x59, 0xbc, 0xdb, 0x84, 0x59, 0x6d, 0x92, 0x91, 0x97, 0x4e, 0x15, 0xf2, 0x07, 0x2b, 0x44,
	0x16, 0xf7, 0x9c, 0x1a, 0xa1, 0x99, 0xf8, 0x17, 0x75, 0x70, 0xc1, 0x25, 0xe5, 0x88, 0x71, 0xb0,
	0x60, 0x72, 0x79, 0x86, 0x48, 0x5e, 0xc1, 0x1e, 0x1a, 0x3a, 0xbf, 0xa3, 0x73, 0x96, 0xba, 0x0d,
	0x7d, 0xb7, 0x01, 0xd0, 0x5c, 0xdc, 0x5d, 0xb0, 0xd4, 0x6e, 0x69, 0xbb, 0x19, 0xae, 0xba, 0x79,
	0x08, 0x81, 0xb6, 0x1f, 0x76, 0x1d, 0x47, 0x4e, 0xe8, 0x69, 0x93, 0xf0, 0xeb, 0xda, 0x64, 0xe4,
	0xd7, 0x22, 0x8b, 0xa1, 0x31, 0x39, 0xb1, 0xcb, 0x15, 0xff, 0x07, 0x2b, 0xfe, 0x0f, 0x7f, 0xb6,
	0x3d, 0xad, 0x6e, 0xca, 0xb2, 0xed, 0xe9, 0x21, 0x04, 0xe8, 0x49, 0xde, 0xd1, 0x1e, 0x3a, 0xd2,
	0x6a, 0xed, 0x9d, 0x87, 0x6b, 0xef, 0xfe, 0xab, 0xf6, 0xdd, 0x07, 0x6b, 0xef, 0xdd, 0x77, 0x92,
	0x7c, 0x97, 0x83, 0xa6, 0x92, 0x45, 0x97, 0xf5, 0xa2, 0xcb, 0xde, 0xd0, 0x50, 0xb7, 0xba, 0xac,
	0xeb, 0x2e, 0x7b, 0x37, 0x03, 0xed, 0xbb, 0xdc, 0xb6, 0x26, 0x5a, 0xb5, 0xe6, 0x4f, 0x07, 0x1e,
	0x2d, 0xad, 0x59, 0xba, 0xf2, 0x1a, 0xf6, 0x85, 0xd4, 0x1c, 0x73, 0x96, 0xae, 0xcc, 0xe3, 0xde,
	0xe2, 0xad, 0x1d, 0x4a, 0x72, 0x00, 0xdd, 0x92, 0x71, 0x67, 0x4f, 0x94, 0xd8, 0x47, 0xf2, 0x15,
	0xba, 0xcd, 0x3c, 0x0d, 0xc6, 0x93, 0xd1, 0x7f, 0xdc, 0x0f, 0xa3, 0xfb, 0x26, 0x2f, 0xb1, 0x59,
	0x09, 0x83, 0x5d, 0x3b, 0x12, 0xce, 0xce, 0xc1, 0xf8, 0x62, 0x9b, 0xec, 0x1b, 0xc3, 0x95, 0xb8,
	0xd4, 0x16, 0x61, 0x4f, 0x88, 0xeb, 0xc9, 0xd6, 0x88, 0xb5, 0xb3, 0x96, 0xb8, 0xd4, 0xc3, 0xdf,
	0x3b, 0x70, 0xb0, 0x7e, 0x01, 0x92, 0x19, 0x84, 0x33, 0x7f, 0xa9, 0x55, 0xf1, 0xd8, 0xb1, 0xcf,
	0xb7, 0x60, 0x6f, 0xdc, 0x91, 0x49, 0x7f, 0x56, 0x5c, 0xd9, 0xe4, 0x24, 0x85, 0xb0, 0x6e, 0x75,
	0x15, 0x9f, 0x1e, 0x77, 0x4f, 0x06, 0xe3, 0x0f, 0x5b, 0x16, 0xd9, 0x82, 0x88, 0xdc, 0x41, 0xa6,
	0x81, 0xfb, 0x7b, 0x38, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x22, 0x7b, 0xbd, 0x38, 0x06,
	0x00, 0x00,
}
