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

type Dot1XIfPortControlStatistics struct {
	EnableSucc           uint32   `protobuf:"varint,1,opt,name=enable_succ,json=enableSucc,proto3" json:"enable_succ,omitempty"`
	EnableFail           uint32   `protobuf:"varint,2,opt,name=enable_fail,json=enableFail,proto3" json:"enable_fail,omitempty"`
	AddClientSucc        uint32   `protobuf:"varint,3,opt,name=add_client_succ,json=addClientSucc,proto3" json:"add_client_succ,omitempty"`
	AddClientFail        uint32   `protobuf:"varint,4,opt,name=add_client_fail,json=addClientFail,proto3" json:"add_client_fail,omitempty"`
	RemoveClientSucc     uint32   `protobuf:"varint,5,opt,name=remove_client_succ,json=removeClientSucc,proto3" json:"remove_client_succ,omitempty"`
	RemoveClientFail     uint32   `protobuf:"varint,6,opt,name=remove_client_fail,json=removeClientFail,proto3" json:"remove_client_fail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XIfPortControlStatistics) Reset()         { *m = Dot1XIfPortControlStatistics{} }
func (m *Dot1XIfPortControlStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XIfPortControlStatistics) ProtoMessage()    {}
func (*Dot1XIfPortControlStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb6866fc1c150562, []int{2}
}

func (m *Dot1XIfPortControlStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XIfPortControlStatistics.Unmarshal(m, b)
}
func (m *Dot1XIfPortControlStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XIfPortControlStatistics.Marshal(b, m, deterministic)
}
func (m *Dot1XIfPortControlStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XIfPortControlStatistics.Merge(m, src)
}
func (m *Dot1XIfPortControlStatistics) XXX_Size() int {
	return xxx_messageInfo_Dot1XIfPortControlStatistics.Size(m)
}
func (m *Dot1XIfPortControlStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XIfPortControlStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XIfPortControlStatistics proto.InternalMessageInfo

func (m *Dot1XIfPortControlStatistics) GetEnableSucc() uint32 {
	if m != nil {
		return m.EnableSucc
	}
	return 0
}

func (m *Dot1XIfPortControlStatistics) GetEnableFail() uint32 {
	if m != nil {
		return m.EnableFail
	}
	return 0
}

func (m *Dot1XIfPortControlStatistics) GetAddClientSucc() uint32 {
	if m != nil {
		return m.AddClientSucc
	}
	return 0
}

func (m *Dot1XIfPortControlStatistics) GetAddClientFail() uint32 {
	if m != nil {
		return m.AddClientFail
	}
	return 0
}

func (m *Dot1XIfPortControlStatistics) GetRemoveClientSucc() uint32 {
	if m != nil {
		return m.RemoveClientSucc
	}
	return 0
}

func (m *Dot1XIfPortControlStatistics) GetRemoveClientFail() uint32 {
	if m != nil {
		return m.RemoveClientFail
	}
	return 0
}

type Dot1XIfAuthStatistics struct {
	RxStart                    uint32                        `protobuf:"varint,1,opt,name=rx_start,json=rxStart,proto3" json:"rx_start,omitempty"`
	RxLogoff                   uint32                        `protobuf:"varint,2,opt,name=rx_logoff,json=rxLogoff,proto3" json:"rx_logoff,omitempty"`
	RxResp                     uint32                        `protobuf:"varint,3,opt,name=rx_resp,json=rxResp,proto3" json:"rx_resp,omitempty"`
	RxRespId                   uint32                        `protobuf:"varint,4,opt,name=rx_resp_id,json=rxRespId,proto3" json:"rx_resp_id,omitempty"`
	RxInvalid                  uint32                        `protobuf:"varint,5,opt,name=rx_invalid,json=rxInvalid,proto3" json:"rx_invalid,omitempty"`
	RxLenErr                   uint32                        `protobuf:"varint,6,opt,name=rx_len_err,json=rxLenErr,proto3" json:"rx_len_err,omitempty"`
	RxMyMacErr                 uint32                        `protobuf:"varint,7,opt,name=rx_my_mac_err,json=rxMyMacErr,proto3" json:"rx_my_mac_err,omitempty"`
	RxTotal                    uint32                        `protobuf:"varint,8,opt,name=rx_total,json=rxTotal,proto3" json:"rx_total,omitempty"`
	TxReq                      uint32                        `protobuf:"varint,9,opt,name=tx_req,json=txReq,proto3" json:"tx_req,omitempty"`
	TxReqid                    uint32                        `protobuf:"varint,10,opt,name=tx_reqid,json=txReqid,proto3" json:"tx_reqid,omitempty"`
	TxTotal                    uint32                        `protobuf:"varint,11,opt,name=tx_total,json=txTotal,proto3" json:"tx_total,omitempty"`
	PacketDropNoConfigReceived uint32                        `protobuf:"varint,12,opt,name=packet_drop_no_config_received,json=packetDropNoConfigReceived,proto3" json:"packet_drop_no_config_received,omitempty"`
	PortControl                *Dot1XIfPortControlStatistics `protobuf:"bytes,13,opt,name=port_control,json=portControl,proto3" json:"port_control,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                      `json:"-"`
	XXX_unrecognized           []byte                        `json:"-"`
	XXX_sizecache              int32                         `json:"-"`
}

func (m *Dot1XIfAuthStatistics) Reset()         { *m = Dot1XIfAuthStatistics{} }
func (m *Dot1XIfAuthStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XIfAuthStatistics) ProtoMessage()    {}
func (*Dot1XIfAuthStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb6866fc1c150562, []int{3}
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

func (m *Dot1XIfAuthStatistics) GetPacketDropNoConfigReceived() uint32 {
	if m != nil {
		return m.PacketDropNoConfigReceived
	}
	return 0
}

func (m *Dot1XIfAuthStatistics) GetPortControl() *Dot1XIfPortControlStatistics {
	if m != nil {
		return m.PortControl
	}
	return nil
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
	return fileDescriptor_eb6866fc1c150562, []int{4}
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

type Dot1XIfLocalEapStatistics struct {
	Requests             uint32   `protobuf:"varint,1,opt,name=requests,proto3" json:"requests,omitempty"`
	Replies              uint32   `protobuf:"varint,2,opt,name=replies,proto3" json:"replies,omitempty"`
	Timeout              uint32   `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	DroppedNoEap         uint32   `protobuf:"varint,4,opt,name=dropped_no_eap,json=droppedNoEap,proto3" json:"dropped_no_eap,omitempty"`
	Dropped              uint32   `protobuf:"varint,5,opt,name=dropped,proto3" json:"dropped,omitempty"`
	Success              uint32   `protobuf:"varint,6,opt,name=success,proto3" json:"success,omitempty"`
	Failed               uint32   `protobuf:"varint,7,opt,name=failed,proto3" json:"failed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dot1XIfLocalEapStatistics) Reset()         { *m = Dot1XIfLocalEapStatistics{} }
func (m *Dot1XIfLocalEapStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XIfLocalEapStatistics) ProtoMessage()    {}
func (*Dot1XIfLocalEapStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb6866fc1c150562, []int{5}
}

func (m *Dot1XIfLocalEapStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dot1XIfLocalEapStatistics.Unmarshal(m, b)
}
func (m *Dot1XIfLocalEapStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dot1XIfLocalEapStatistics.Marshal(b, m, deterministic)
}
func (m *Dot1XIfLocalEapStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dot1XIfLocalEapStatistics.Merge(m, src)
}
func (m *Dot1XIfLocalEapStatistics) XXX_Size() int {
	return xxx_messageInfo_Dot1XIfLocalEapStatistics.Size(m)
}
func (m *Dot1XIfLocalEapStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Dot1XIfLocalEapStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_Dot1XIfLocalEapStatistics proto.InternalMessageInfo

func (m *Dot1XIfLocalEapStatistics) GetRequests() uint32 {
	if m != nil {
		return m.Requests
	}
	return 0
}

func (m *Dot1XIfLocalEapStatistics) GetReplies() uint32 {
	if m != nil {
		return m.Replies
	}
	return 0
}

func (m *Dot1XIfLocalEapStatistics) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Dot1XIfLocalEapStatistics) GetDroppedNoEap() uint32 {
	if m != nil {
		return m.DroppedNoEap
	}
	return 0
}

func (m *Dot1XIfLocalEapStatistics) GetDropped() uint32 {
	if m != nil {
		return m.Dropped
	}
	return 0
}

func (m *Dot1XIfLocalEapStatistics) GetSuccess() uint32 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *Dot1XIfLocalEapStatistics) GetFailed() uint32 {
	if m != nil {
		return m.Failed
	}
	return 0
}

type Dot1XIfStatistics struct {
	InterfaceName        string                     `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Pae                  string                     `protobuf:"bytes,51,opt,name=pae,proto3" json:"pae,omitempty"`
	Idb                  *Dot1XIdbStatistics        `protobuf:"bytes,52,opt,name=idb,proto3" json:"idb,omitempty"`
	Auth                 *Dot1XIfAuthStatistics     `protobuf:"bytes,53,opt,name=auth,proto3" json:"auth,omitempty"`
	Supp                 *Dot1XIfSuppStatistics     `protobuf:"bytes,54,opt,name=supp,proto3" json:"supp,omitempty"`
	LocalEap             *Dot1XIfLocalEapStatistics `protobuf:"bytes,55,opt,name=local_eap,json=localEap,proto3" json:"local_eap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Dot1XIfStatistics) Reset()         { *m = Dot1XIfStatistics{} }
func (m *Dot1XIfStatistics) String() string { return proto.CompactTextString(m) }
func (*Dot1XIfStatistics) ProtoMessage()    {}
func (*Dot1XIfStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb6866fc1c150562, []int{6}
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

func (m *Dot1XIfStatistics) GetLocalEap() *Dot1XIfLocalEapStatistics {
	if m != nil {
		return m.LocalEap
	}
	return nil
}

func init() {
	proto.RegisterType((*Dot1XIfStatistics_KEYS)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_if_statistics_KEYS")
	proto.RegisterType((*Dot1XIdbStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_idb_statistics")
	proto.RegisterType((*Dot1XIfPortControlStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_if_port_control_statistics")
	proto.RegisterType((*Dot1XIfAuthStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_if_auth_statistics")
	proto.RegisterType((*Dot1XIfSuppStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_if_supp_statistics")
	proto.RegisterType((*Dot1XIfLocalEapStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_if_local_eap_statistics")
	proto.RegisterType((*Dot1XIfStatistics)(nil), "cisco_ios_xr_dot1x_oper.dot1x.statistics.interface_statistics.interface_statistic.dot1x_if_statistics")
}

func init() { proto.RegisterFile("dot1x_if_statistics.proto", fileDescriptor_eb6866fc1c150562) }

var fileDescriptor_eb6866fc1c150562 = []byte{
	// 815 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0x96, 0x77, 0xed, 0xb1, 0x5d, 0x5e, 0x2f, 0x51, 0x43, 0x60, 0x12, 0x08, 0x2c, 0x16, 0xa0,
	0x08, 0x21, 0x4b, 0x24, 0xfc, 0xdc, 0xd9, 0x18, 0x69, 0x05, 0x59, 0xc4, 0x2c, 0x17, 0x4e, 0xad,
	0xf6, 0x74, 0xcd, 0xd2, 0xca, 0xb8, 0xbb, 0xdd, 0xd3, 0xde, 0xcc, 0x9e, 0x78, 0x02, 0x24, 0x4e,
	0x88, 0xa7, 0xe0, 0x85, 0x38, 0xf1, 0x26, 0xa8, 0x7f, 0xc6, 0x99, 0xd9, 0xb5, 0x36, 0x97, 0xf8,
	0x36, 0x55, 0xdf, 0xd7, 0xd5, 0xd5, 0x55, 0x5f, 0x95, 0x0d, 0x0f, 0xb8, 0xb2, 0x5f, 0xd6, 0x54,
	0x14, 0xb4, 0xb2, 0xcc, 0x8a, 0xca, 0x8a, 0xbc, 0x9a, 0x6b, 0xa3, 0xac, 0x22, 0x3f, 0xe7, 0xa2,
	0xca, 0x15, 0x15, 0xaa, 0xa2, 0xb5, 0xa1, 0x81, 0xa7, 0x34, 0x9a, 0xb9, 0xff, 0x9c, 0xb7, 0xf8,
	0x42, 0x5a, 0x34, 0x05, 0xcb, 0x91, 0xde, 0xed, 0x9c, 0xcd, 0x21, 0xdd, 0x71, 0x1f, 0xfd, 0x61,
	0xf1, 0xeb, 0x05, 0x21, 0xd0, 0x97, 0x6c, 0x85, 0x69, 0xef, 0xa4, 0xf7, 0x78, 0x9c, 0xf9, 0xef,
	0x59, 0x0d, 0xef, 0x44, 0x3e, 0x5f, 0xb6, 0x0e, 0x90, 0x07, 0x30, 0x32, 0x35, 0xb5, 0xca, 0xb2,
	0xd2, 0xf3, 0xa7, 0xd9, 0xd0, 0xd4, 0xbf, 0x38, 0xd3, 0x41, 0xb6, 0x81, 0x0e, 0x02, 0x64, 0x23,
	0xf4, 0x39, 0x10, 0xa9, 0xa8, 0xa9, 0xa9, 0x92, 0x54, 0x48, 0x5b, 0x50, 0xae, 0x5e, 0xca, 0xf4,
	0xd0, 0x93, 0x8e, 0xa5, 0xca, 0xea, 0x9f, 0xe4, 0x99, 0xb4, 0xc5, 0x33, 0xf5, 0x52, 0xce, 0xfe,
	0x3c, 0x80, 0x93, 0x6d, 0xaa, 0x5a, 0x19, 0x4b, 0x73, 0x25, 0xad, 0x51, 0x65, 0x3b, 0x8d, 0x8f,
	0x60, 0x82, 0x92, 0x2d, 0x4b, 0xa4, 0xd5, 0x26, 0xcf, 0x63, 0x26, 0x10, 0x5c, 0x17, 0x9b, 0x3c,
	0x6f, 0x11, 0x0a, 0x26, 0x9a, 0x7c, 0x22, 0xe1, 0x7b, 0x26, 0x4a, 0xf2, 0x19, 0xbc, 0xc5, 0x38,
	0xa7, 0x79, 0x29, 0x50, 0xda, 0x10, 0x25, 0xe4, 0x33, 0x65, 0x9c, 0x9f, 0x7a, 0xaf, 0x0f, 0xd4,
	0xe5, 0xf9, 0x60, 0xfd, 0x1b, 0x3c, 0x1f, 0xef, 0x0b, 0x20, 0x06, 0x57, 0xea, 0x0a, 0x3b, 0x21,
	0x07, 0x9e, 0x7a, 0x2f, 0x20, 0xad, 0xa8, 0xb7, 0xd8, 0x3e, 0x70, 0x72, 0x9b, 0xed, 0x62, 0xcf,
	0xfe, 0xe9, 0xb7, 0xba, 0xc7, 0x36, 0xf6, 0xb7, 0xdb, 0x1d, 0xa9, 0x2c, 0x33, 0xf6, 0x55, 0x47,
	0x2e, 0x9c, 0x49, 0xde, 0x87, 0xb1, 0xa9, 0x69, 0xa9, 0x2e, 0x55, 0x51, 0xc4, 0x12, 0x8c, 0x4c,
	0xfd, 0xa3, 0xb7, 0xc9, 0x7b, 0x30, 0x34, 0x35, 0x35, 0x58, 0xe9, 0xf8, 0xf0, 0xc4, 0xd4, 0x19,
	0x56, 0x9a, 0x7c, 0x00, 0x10, 0x01, 0x2a, 0x78, 0x7c, 0xec, 0x28, 0x60, 0x67, 0x9c, 0x3c, 0xf2,
	0xa8, 0x90, 0x57, 0xac, 0x14, 0x3c, 0xbe, 0x6f, 0x6c, 0xea, 0xb3, 0xe0, 0x88, 0x87, 0x4b, 0x94,
	0x14, 0x8d, 0x89, 0x0f, 0x72, 0x77, 0xa2, 0x5c, 0x18, 0x43, 0x3e, 0x86, 0xa9, 0xa9, 0xe9, 0xea,
	0x9a, 0xae, 0x58, 0xee, 0x09, 0xc3, 0xd0, 0x17, 0x53, 0x3f, 0xbf, 0x7e, 0xce, 0x72, 0x47, 0x69,
	0x0b, 0x6c, 0xd4, 0x15, 0xd8, 0x7d, 0x48, 0xac, 0x4b, 0x6c, 0x9d, 0x8e, 0x3d, 0x30, 0xb0, 0x75,
	0x86, 0xeb, 0xa8, 0x3b, 0x83, 0x6b, 0xc1, 0x53, 0x68, 0x74, 0x97, 0x39, 0xb3, 0x23, 0xc9, 0x49,
	0x57, 0x92, 0xdf, 0xc1, 0x87, 0x9a, 0xe5, 0x2f, 0xd0, 0x52, 0x6e, 0x94, 0xa6, 0x52, 0x39, 0x99,
	0x15, 0xe2, 0x92, 0x1a, 0xcc, 0x51, 0x5c, 0x21, 0x4f, 0x8f, 0xfc, 0x81, 0x87, 0x81, 0xf5, 0xcc,
	0x28, 0x7d, 0xae, 0x4e, 0x3d, 0x25, 0x8b, 0x0c, 0xf2, 0x57, 0x0f, 0x8e, 0xda, 0x0a, 0x4d, 0xa7,
	0x27, 0xbd, 0xc7, 0x93, 0x27, 0xd5, 0xfc, 0x8d, 0xcf, 0xef, 0xfc, 0x75, 0x13, 0x91, 0x4d, 0x1c,
	0x70, 0x1a, 0xfc, 0xb3, 0xbf, 0x0f, 0xda, 0xe3, 0xbe, 0xd1, 0xba, 0x2d, 0x98, 0xfb, 0x90, 0x98,
	0x50, 0xc6, 0x20, 0x97, 0x81, 0xf1, 0x65, 0xec, 0x36, 0xf6, 0xe0, 0xee, 0xc6, 0x1e, 0xbe, 0xae,
	0xb1, 0xfd, 0x3b, 0x1b, 0x3b, 0xd8, 0xb5, 0x39, 0x82, 0x84, 0x93, 0xa6, 0x4d, 0x5b, 0x09, 0xdb,
	0xad, 0x84, 0x83, 0x5a, 0x46, 0xb6, 0x25, 0x61, 0x1b, 0x25, 0x1c, 0xa4, 0x92, 0xd8, 0x20, 0xe1,
	0x76, 0xdf, 0xc7, 0x9d, 0xbe, 0xcf, 0xfe, 0xeb, 0xc1, 0xa3, 0x6d, 0x69, 0x4a, 0x95, 0xb3, 0x92,
	0x22, 0xeb, 0xd4, 0xe7, 0x21, 0x8c, 0x0c, 0xae, 0x37, 0x58, 0xd9, 0x2a, 0x56, 0x68, 0x6b, 0x93,
	0x14, 0x86, 0x06, 0x75, 0x29, 0xb0, 0x6a, 0x56, 0x5c, 0x34, 0x1d, 0x62, 0xc5, 0x0a, 0xd5, 0xc6,
	0xc6, 0xe2, 0x34, 0x26, 0xf9, 0x04, 0x8e, 0x9d, 0xc4, 0x34, 0x72, 0xa7, 0x32, 0x64, 0x3a, 0x16,
	0xe7, 0x28, 0x7a, 0xcf, 0xd5, 0x82, 0x69, 0x77, 0x3e, 0xda, 0x4d, 0x75, 0xa2, 0xe9, 0x10, 0xb7,
	0x4b, 0xb0, 0xaa, 0x9a, 0xe2, 0x44, 0x93, 0xbc, 0x0b, 0x89, 0xdb, 0x1b, 0xc8, 0x63, 0x65, 0xa2,
	0x35, 0xfb, 0xb7, 0x0f, 0x6f, 0xef, 0xd8, 0xf6, 0xe4, 0x53, 0x38, 0x7e, 0xa5, 0x2d, 0xbf, 0xf2,
	0x9f, 0xf8, 0x95, 0x3f, 0xdd, 0x7a, 0xcf, 0xd9, 0x0a, 0xc9, 0x3d, 0x38, 0xd4, 0x0c, 0xd3, 0xa7,
	0x1e, 0x73, 0x9f, 0xe4, 0x1a, 0x0e, 0x05, 0x5f, 0xa6, 0x5f, 0x79, 0x79, 0x5f, 0xee, 0x4f, 0xde,
	0x9d, 0xdf, 0x9a, 0xcc, 0xdd, 0x49, 0x7e, 0x87, 0xbe, 0xdb, 0x78, 0xe9, 0xd7, 0xfe, 0xee, 0x17,
	0xfb, 0x1c, 0xad, 0x1b, 0x9b, 0x35, 0xf3, 0x17, 0xbb, 0x04, 0xdc, 0x04, 0xa5, 0xdf, 0xec, 0x3f,
	0x81, 0x1b, 0x93, 0x9a, 0xf9, 0x8b, 0xc9, 0x1f, 0x3d, 0x18, 0x6f, 0x85, 0x9a, 0x7e, 0xeb, 0xd3,
	0xd0, 0xfb, 0x4c, 0x63, 0xd7, 0x54, 0x64, 0x23, 0xef, 0x5d, 0x30, 0xbd, 0x4c, 0xfc, 0x9f, 0x94,
	0xa7, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xaf, 0x7f, 0xc7, 0xc1, 0x08, 0x00, 0x00,
}
