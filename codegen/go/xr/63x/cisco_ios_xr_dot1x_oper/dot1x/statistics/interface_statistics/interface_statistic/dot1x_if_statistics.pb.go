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
// source: dot1x_if_statistics.proto

package cisco_ios_xr_dot1x_oper_dot1x_statistics_interface_statistics_interface_statistic

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

type Dot1XIfStatistics_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XIfStatistics_KEYS) Reset()         { *m = Dot1XIfStatistics_KEYS{} }
func (m *Dot1XIfStatistics_KEYS) String() string { return proto.CompactTextString(m) }
func (*Dot1XIfStatistics_KEYS) ProtoMessage()    {}
func (*Dot1XIfStatistics_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb6866fc1c150562, []int{0}
}

func (m *Dot1XIfStatistics_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XIfStatistics_KEYS.Unmarshal(m, b)
}
func (m *Dot1XIfStatistics_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XIfStatistics_KEYS.Marshal(b, m, deterministic)
}
func (m *Dot1XIfStatistics_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XIfStatistics_KEYS.Merge(m, src)
}
func (m *Dot1XIfStatistics_KEYS) XXX_Size() int {
	return xxx_messageInfo_Dot1XIfStatistics_KEYS.Size(m)
}
func (m *Dot1XIfStatistics_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XIfStatistics_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XIfStatistics_KEYS proto.InternalMessageInfo

func (m *Dot1XIfStatistics_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
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
	return fileDescriptor_eb6866fc1c150562, []int{1}
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
	return fileDescriptor_eb6866fc1c150562, []int{2}
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
	return fileDescriptor_eb6866fc1c150562, []int{3}
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
	InterfaceName        string                 `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Pae                  string                 `protobuf:"bytes,51,opt,name=pae,proto3" json:"pae,omitempty"`
	Idb                  *Dot1XIdbStatistics    `protobuf:"bytes,52,opt,name=idb,proto3" json:"idb,omitempty"`
	Auth                 *Dot1XIfAuthStatistics `protobuf:"bytes,53,opt,name=auth,proto3" json:"auth,omitempty"`
	Supp                 *Dot1XIfSuppStatistics `protobuf:"bytes,54,opt,name=supp,proto3" json:"supp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Dot1XIfStatistics) Reset()         { *m = Dot1XIfStatistics{} }
func (m *Dot1XIfStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XIfStatistics) ProtoMessage()    {}
func (*Dot1XIfStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb6866fc1c150562, []int{4}
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

func init() {
	proto.RegisterType((*Dot1XIfStatistics_KEYS)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_if_statistics_KEYS")
	proto.RegisterType((*Dot1XIdbStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_idb_statistics")
	proto.RegisterType((*Dot1XIfAuthStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_if_auth_statistics")
	proto.RegisterType((*Dot1XIfSuppStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_if_supp_statistics")
	proto.RegisterType((*Dot1XIfStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_if_statistics")
}

func init() { proto.RegisterFile("dot1x_if_statistics.proto", fileDescriptor_eb6866fc1c150562) }

var fileDescriptor_eb6866fc1c150562 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0xbe, 0x36, 0xc9, 0x54, 0xa9, 0x90, 0xa1, 0xc2, 0x15, 0x20, 0x95, 0x48, 0x48, 0x15,
	0x87, 0x95, 0x68, 0x81, 0x5f, 0x40, 0x0f, 0x11, 0x14, 0xc4, 0x96, 0x0b, 0x27, 0xcb, 0xd9, 0xf5,
	0x16, 0x8b, 0xc4, 0xde, 0xcc, 0x4e, 0xa9, 0x73, 0xe2, 0x6f, 0xf0, 0x9f, 0x38, 0xf1, 0x8f, 0x90,
	0xbd, 0x9b, 0x74, 0x43, 0xa3, 0x70, 0x69, 0x6f, 0x3b, 0x7e, 0xcf, 0x7a, 0x33, 0x6f, 0xfc, 0x16,
	0x0e, 0x33, 0x4b, 0xaf, 0x9c, 0xd0, 0xb9, 0x28, 0x49, 0x92, 0x2e, 0x49, 0xa7, 0x65, 0x5c, 0xa0,
	0x25, 0xcb, 0x3e, 0xa7, 0xba, 0x4c, 0xad, 0xd0, 0xb6, 0x14, 0x0e, 0x45, 0xc5, 0xb3, 0x85, 0xc2,
	0x38, 0x7c, 0xc6, 0x0d, 0xbe, 0x36, 0xa4, 0x30, 0x97, 0xa9, 0x12, 0xbb, 0x0f, 0xc7, 0x31, 0xf0,
	0x2d, 0x7a, 0xe2, 0xfd, 0xd9, 0xd7, 0x0b, 0xc6, 0xa0, 0x6b, 0xe4, 0x5c, 0xf1, 0xd6, 0x51, 0xeb,
	0x78, 0x98, 0x84, 0xef, 0xb1, 0x83, 0x47, 0x35, 0x3f, 0x9b, 0x36, 0x2e, 0xb0, 0x43, 0x18, 0xa0,
	0x13, 0x64, 0x49, 0xce, 0x02, 0x7f, 0x94, 0xf4, 0xd1, 0x7d, 0xf1, 0xa5, 0x87, 0x68, 0x05, 0xb5,
	0x2b, 0x88, 0x6a, 0xe8, 0x25, 0x30, 0x63, 0x05, 0x3a, 0x61, 0x8d, 0xd0, 0x86, 0x72, 0x91, 0xd9,
	0x6b, 0xc3, 0x3b, 0x81, 0xb4, 0x6f, 0x6c, 0xe2, 0x3e, 0x99, 0x89, 0xa1, 0xfc, 0x9d, 0xbd, 0x36,
	0xe3, 0x3f, 0xed, 0x46, 0xab, 0xf2, 0x8a, 0xbe, 0xdd, 0x96, 0x2f, 0x49, 0x22, 0xdd, 0xc8, 0x5f,
	0xf8, 0x92, 0x3d, 0x81, 0x21, 0x3a, 0x31, 0xb3, 0x97, 0x36, 0xcf, 0x6b, 0xfd, 0x01, 0xba, 0x0f,
	0xa1, 0x66, 0x8f, 0xa1, 0x8f, 0x4e, 0xa0, 0x2a, 0x8b, 0x5a, 0x35, 0x42, 0x97, 0xa8, 0xb2, 0x60,
	0x4f, 0x01, 0x6a, 0x40, 0xe8, 0x8c, 0x77, 0x57, 0xd7, 0x3c, 0x36, 0xc9, 0xd8, 0xb3, 0x80, 0x6a,
	0xf3, 0x43, 0xce, 0x74, 0xc6, 0x7b, 0x01, 0x1d, 0xa2, 0x9b, 0x54, 0x07, 0xf5, 0xe5, 0x99, 0x32,
	0x42, 0x21, 0xf2, 0x68, 0xad, 0xa9, 0xcc, 0x19, 0x22, 0x7b, 0x0e, 0x23, 0x74, 0x62, 0xbe, 0x14,
	0x73, 0x99, 0x06, 0x42, 0x3f, 0x10, 0x00, 0xdd, 0xf9, 0xf2, 0x5c, 0xa6, 0x9e, 0xd2, 0x74, 0x73,
	0xb0, 0xe9, 0xe6, 0x01, 0x44, 0xe4, 0x1b, 0x5b, 0xf0, 0x61, 0x00, 0x7a, 0xe4, 0x12, 0xb5, 0xa8,
	0x4d, 0x46, 0xb5, 0xd0, 0x19, 0x87, 0x95, 0xc9, 0x89, 0x2f, 0x37, 0xfc, 0xdf, 0xdb, 0xf0, 0x7f,
	0xfc, 0xab, 0xe9, 0x69, 0x79, 0x55, 0x14, 0x4d, 0x4f, 0x0f, 0x20, 0xc2, 0x4a, 0xa9, 0x72, 0xb4,
	0x87, 0x41, 0x69, 0x73, 0xf6, 0xf6, 0xee, 0xd9, 0x3b, 0xff, 0x9b, 0xbd, 0xbb, 0x73, 0xf6, 0xde,
	0xb6, 0x97, 0x54, 0x6d, 0x39, 0x5a, 0x4d, 0xb2, 0xde, 0x32, 0xad, 0xb7, 0x5c, 0x19, 0x3a, 0xa0,
	0xc6, 0x96, 0xa9, 0xde, 0x72, 0xe5, 0x66, 0x44, 0xd5, 0x96, 0x9b, 0xd6, 0x0c, 0x37, 0xad, 0xf9,
	0xdd, 0x81, 0x87, 0x5b, 0x92, 0xc1, 0x5e, 0xc0, 0xfe, 0x4d, 0x8e, 0x42, 0x3c, 0x4e, 0x42, 0x3c,
	0x46, 0xeb, 0xd3, 0x8f, 0x72, 0xae, 0xd8, 0x03, 0xe8, 0x14, 0x52, 0xf1, 0xd3, 0x80, 0xf9, 0x4f,
	0xb6, 0x84, 0x8e, 0xce, 0xa6, 0xfc, 0xf5, 0x51, 0xeb, 0x78, 0xef, 0xe4, 0x32, 0xbe, 0xf3, 0x28,
	0xc7, 0xdb, 0x72, 0x99, 0x78, 0x4d, 0xf6, 0x13, 0xba, 0x3e, 0x30, 0xfc, 0x4d, 0xd0, 0xfe, 0x7e,
	0x7f, 0xda, 0xb7, 0x82, 0x99, 0x04, 0x61, 0xdf, 0x80, 0x7f, 0x5d, 0xfc, 0xed, 0xfd, 0x37, 0xf0,
	0xcf, 0x2b, 0x4e, 0x82, 0xf0, 0x34, 0x0a, 0x3f, 0xd0, 0xd3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x2e, 0x76, 0xd6, 0x6c, 0x5d, 0x05, 0x00, 0x00,
}
