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
// source: atm_vcm_vc.proto

package cisco_ios_xr_atm_vcm_oper_atm_vcm_nodes_node_vcs_vc

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

type AtmVcmVc_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Vpi                  uint32   `protobuf:"varint,3,opt,name=vpi,proto3" json:"vpi,omitempty"`
	Vci                  uint32   `protobuf:"varint,4,opt,name=vci,proto3" json:"vci,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AtmVcmVc_KEYS) Reset()         { *m = AtmVcmVc_KEYS{} }
func (m *AtmVcmVc_KEYS) String() string { return proto.CompactTextString(m) }
func (*AtmVcmVc_KEYS) ProtoMessage()    {}
func (*AtmVcmVc_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_606e93107615db12, []int{0}
}

func (m *AtmVcmVc_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtmVcmVc_KEYS.Unmarshal(m, b)
}
func (m *AtmVcmVc_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtmVcmVc_KEYS.Marshal(b, m, deterministic)
}
func (m *AtmVcmVc_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtmVcmVc_KEYS.Merge(m, src)
}
func (m *AtmVcmVc_KEYS) XXX_Size() int {
	return xxx_messageInfo_AtmVcmVc_KEYS.Size(m)
}
func (m *AtmVcmVc_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AtmVcmVc_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AtmVcmVc_KEYS proto.InternalMessageInfo

func (m *AtmVcmVc_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *AtmVcmVc_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *AtmVcmVc_KEYS) GetVpi() uint32 {
	if m != nil {
		return m.Vpi
	}
	return 0
}

func (m *AtmVcmVc_KEYS) GetVci() uint32 {
	if m != nil {
		return m.Vci
	}
	return 0
}

type AtmVcmCellPackingInfo struct {
	LocalMaxCellsPackedPerPacket      uint32   `protobuf:"varint,1,opt,name=local_max_cells_packed_per_packet,json=localMaxCellsPackedPerPacket,proto3" json:"local_max_cells_packed_per_packet,omitempty"`
	NegotiatedMaxCellsPackedPerPacket uint32   `protobuf:"varint,2,opt,name=negotiated_max_cells_packed_per_packet,json=negotiatedMaxCellsPackedPerPacket,proto3" json:"negotiated_max_cells_packed_per_packet,omitempty"`
	MaxCellPackedTimeout              uint32   `protobuf:"varint,3,opt,name=max_cell_packed_timeout,json=maxCellPackedTimeout,proto3" json:"max_cell_packed_timeout,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *AtmVcmCellPackingInfo) Reset()         { *m = AtmVcmCellPackingInfo{} }
func (m *AtmVcmCellPackingInfo) String() string { return proto.CompactTextString(m) }
func (*AtmVcmCellPackingInfo) ProtoMessage()    {}
func (*AtmVcmCellPackingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_606e93107615db12, []int{1}
}

func (m *AtmVcmCellPackingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtmVcmCellPackingInfo.Unmarshal(m, b)
}
func (m *AtmVcmCellPackingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtmVcmCellPackingInfo.Marshal(b, m, deterministic)
}
func (m *AtmVcmCellPackingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtmVcmCellPackingInfo.Merge(m, src)
}
func (m *AtmVcmCellPackingInfo) XXX_Size() int {
	return xxx_messageInfo_AtmVcmCellPackingInfo.Size(m)
}
func (m *AtmVcmCellPackingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AtmVcmCellPackingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AtmVcmCellPackingInfo proto.InternalMessageInfo

func (m *AtmVcmCellPackingInfo) GetLocalMaxCellsPackedPerPacket() uint32 {
	if m != nil {
		return m.LocalMaxCellsPackedPerPacket
	}
	return 0
}

func (m *AtmVcmCellPackingInfo) GetNegotiatedMaxCellsPackedPerPacket() uint32 {
	if m != nil {
		return m.NegotiatedMaxCellsPackedPerPacket
	}
	return 0
}

func (m *AtmVcmCellPackingInfo) GetMaxCellPackedTimeout() uint32 {
	if m != nil {
		return m.MaxCellPackedTimeout
	}
	return 0
}

type AtmVcmVc struct {
	MainInterface        string                 `protobuf:"bytes,50,opt,name=main_interface,json=mainInterface,proto3" json:"main_interface,omitempty"`
	SubInterface         string                 `protobuf:"bytes,51,opt,name=sub_interface,json=subInterface,proto3" json:"sub_interface,omitempty"`
	VcInterface          string                 `protobuf:"bytes,52,opt,name=vc_interface,json=vcInterface,proto3" json:"vc_interface,omitempty"`
	VpiXr                uint32                 `protobuf:"varint,53,opt,name=vpi_xr,json=vpiXr,proto3" json:"vpi_xr,omitempty"`
	VciXr                uint32                 `protobuf:"varint,54,opt,name=vci_xr,json=vciXr,proto3" json:"vci_xr,omitempty"`
	Type                 string                 `protobuf:"bytes,55,opt,name=type,proto3" json:"type,omitempty"`
	Encapsulation        string                 `protobuf:"bytes,56,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	Shape                string                 `protobuf:"bytes,57,opt,name=shape,proto3" json:"shape,omitempty"`
	PeakCellRate         uint32                 `protobuf:"varint,58,opt,name=peak_cell_rate,json=peakCellRate,proto3" json:"peak_cell_rate,omitempty"`
	SustainedCellRate    uint32                 `protobuf:"varint,59,opt,name=sustained_cell_rate,json=sustainedCellRate,proto3" json:"sustained_cell_rate,omitempty"`
	BurstRate            uint32                 `protobuf:"varint,60,opt,name=burst_rate,json=burstRate,proto3" json:"burst_rate,omitempty"`
	EncapsInheritLevel   string                 `protobuf:"bytes,61,opt,name=encaps_inherit_level,json=encapsInheritLevel,proto3" json:"encaps_inherit_level,omitempty"`
	QosInheritLevel      string                 `protobuf:"bytes,62,opt,name=qos_inherit_level,json=qosInheritLevel,proto3" json:"qos_inherit_level,omitempty"`
	TransmitMtu          uint32                 `protobuf:"varint,63,opt,name=transmit_mtu,json=transmitMtu,proto3" json:"transmit_mtu,omitempty"`
	ReceiveMtu           uint32                 `protobuf:"varint,64,opt,name=receive_mtu,json=receiveMtu,proto3" json:"receive_mtu,omitempty"`
	VcOnvpTunnel         bool                   `protobuf:"varint,65,opt,name=vc_onvp_tunnel,json=vcOnvpTunnel,proto3" json:"vc_onvp_tunnel,omitempty"`
	VcOnP2PSubInterface  bool                   `protobuf:"varint,66,opt,name=vc_on_p2p_sub_interface,json=vcOnP2pSubInterface,proto3" json:"vc_on_p2p_sub_interface,omitempty"`
	OperStatus           bool                   `protobuf:"varint,67,opt,name=oper_status,json=operStatus,proto3" json:"oper_status,omitempty"`
	AminStatus           bool                   `protobuf:"varint,68,opt,name=amin_status,json=aminStatus,proto3" json:"amin_status,omitempty"`
	InternalState        string                 `protobuf:"bytes,69,opt,name=internal_state,json=internalState,proto3" json:"internal_state,omitempty"`
	LastStateChangeTime  uint32                 `protobuf:"varint,70,opt,name=last_state_change_time,json=lastStateChangeTime,proto3" json:"last_state_change_time,omitempty"`
	CellPackingData      *AtmVcmCellPackingInfo `protobuf:"bytes,71,opt,name=cell_packing_data,json=cellPackingData,proto3" json:"cell_packing_data,omitempty"`
	TestMode             string                 `protobuf:"bytes,72,opt,name=test_mode,json=testMode,proto3" json:"test_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AtmVcmVc) Reset()         { *m = AtmVcmVc{} }
func (m *AtmVcmVc) String() string { return proto.CompactTextString(m) }
func (*AtmVcmVc) ProtoMessage()    {}
func (*AtmVcmVc) Descriptor() ([]byte, []int) {
	return fileDescriptor_606e93107615db12, []int{2}
}

func (m *AtmVcmVc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtmVcmVc.Unmarshal(m, b)
}
func (m *AtmVcmVc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtmVcmVc.Marshal(b, m, deterministic)
}
func (m *AtmVcmVc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtmVcmVc.Merge(m, src)
}
func (m *AtmVcmVc) XXX_Size() int {
	return xxx_messageInfo_AtmVcmVc.Size(m)
}
func (m *AtmVcmVc) XXX_DiscardUnknown() {
	xxx_messageInfo_AtmVcmVc.DiscardUnknown(m)
}

var xxx_messageInfo_AtmVcmVc proto.InternalMessageInfo

func (m *AtmVcmVc) GetMainInterface() string {
	if m != nil {
		return m.MainInterface
	}
	return ""
}

func (m *AtmVcmVc) GetSubInterface() string {
	if m != nil {
		return m.SubInterface
	}
	return ""
}

func (m *AtmVcmVc) GetVcInterface() string {
	if m != nil {
		return m.VcInterface
	}
	return ""
}

func (m *AtmVcmVc) GetVpiXr() uint32 {
	if m != nil {
		return m.VpiXr
	}
	return 0
}

func (m *AtmVcmVc) GetVciXr() uint32 {
	if m != nil {
		return m.VciXr
	}
	return 0
}

func (m *AtmVcmVc) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AtmVcmVc) GetEncapsulation() string {
	if m != nil {
		return m.Encapsulation
	}
	return ""
}

func (m *AtmVcmVc) GetShape() string {
	if m != nil {
		return m.Shape
	}
	return ""
}

func (m *AtmVcmVc) GetPeakCellRate() uint32 {
	if m != nil {
		return m.PeakCellRate
	}
	return 0
}

func (m *AtmVcmVc) GetSustainedCellRate() uint32 {
	if m != nil {
		return m.SustainedCellRate
	}
	return 0
}

func (m *AtmVcmVc) GetBurstRate() uint32 {
	if m != nil {
		return m.BurstRate
	}
	return 0
}

func (m *AtmVcmVc) GetEncapsInheritLevel() string {
	if m != nil {
		return m.EncapsInheritLevel
	}
	return ""
}

func (m *AtmVcmVc) GetQosInheritLevel() string {
	if m != nil {
		return m.QosInheritLevel
	}
	return ""
}

func (m *AtmVcmVc) GetTransmitMtu() uint32 {
	if m != nil {
		return m.TransmitMtu
	}
	return 0
}

func (m *AtmVcmVc) GetReceiveMtu() uint32 {
	if m != nil {
		return m.ReceiveMtu
	}
	return 0
}

func (m *AtmVcmVc) GetVcOnvpTunnel() bool {
	if m != nil {
		return m.VcOnvpTunnel
	}
	return false
}

func (m *AtmVcmVc) GetVcOnP2PSubInterface() bool {
	if m != nil {
		return m.VcOnP2PSubInterface
	}
	return false
}

func (m *AtmVcmVc) GetOperStatus() bool {
	if m != nil {
		return m.OperStatus
	}
	return false
}

func (m *AtmVcmVc) GetAminStatus() bool {
	if m != nil {
		return m.AminStatus
	}
	return false
}

func (m *AtmVcmVc) GetInternalState() string {
	if m != nil {
		return m.InternalState
	}
	return ""
}

func (m *AtmVcmVc) GetLastStateChangeTime() uint32 {
	if m != nil {
		return m.LastStateChangeTime
	}
	return 0
}

func (m *AtmVcmVc) GetCellPackingData() *AtmVcmCellPackingInfo {
	if m != nil {
		return m.CellPackingData
	}
	return nil
}

func (m *AtmVcmVc) GetTestMode() string {
	if m != nil {
		return m.TestMode
	}
	return ""
}

func init() {
	proto.RegisterType((*AtmVcmVc_KEYS)(nil), "cisco_ios_xr_atm_vcm_oper.atm_vcm.nodes.node.vcs.vc.atm_vcm_vc_KEYS")
	proto.RegisterType((*AtmVcmCellPackingInfo)(nil), "cisco_ios_xr_atm_vcm_oper.atm_vcm.nodes.node.vcs.vc.atm_vcm_cell_packing_info")
	proto.RegisterType((*AtmVcmVc)(nil), "cisco_ios_xr_atm_vcm_oper.atm_vcm.nodes.node.vcs.vc.atm_vcm_vc")
}

func init() { proto.RegisterFile("atm_vcm_vc.proto", fileDescriptor_606e93107615db12) }

var fileDescriptor_606e93107615db12 = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5b, 0x6f, 0x23, 0x35,
	0x14, 0xc7, 0x95, 0xbd, 0xb1, 0x3d, 0x49, 0xda, 0xad, 0x5b, 0x58, 0x23, 0x40, 0x9b, 0x94, 0x82,
	0x22, 0x1e, 0x46, 0x28, 0xd9, 0x72, 0xbf, 0xb7, 0xa5, 0x54, 0xd0, 0x52, 0xd2, 0x3e, 0xc0, 0x93,
	0xe5, 0x38, 0xa7, 0xa9, 0xd5, 0x19, 0xdb, 0x1d, 0x7b, 0x46, 0x29, 0x5f, 0x87, 0xef, 0xc7, 0x67,
	0x40, 0x3e, 0x33, 0x93, 0x0b, 0x52, 0x79, 0xe0, 0x21, 0x91, 0xe7, 0xf7, 0xff, 0xcd, 0x19, 0xcf,
	0xd1, 0x1c, 0xc3, 0x0b, 0x19, 0x32, 0x51, 0xaa, 0xf8, 0x4b, 0x5c, 0x6e, 0x83, 0x65, 0x23, 0xa5,
	0xbd, 0xb2, 0x42, 0x5b, 0x2f, 0xe6, 0xb9, 0x68, 0x62, 0xeb, 0x30, 0x4f, 0xea, 0x8b, 0xc4, 0xd8,
	0x29, 0x7a, 0xfa, 0x4f, 0x4a, 0xe5, 0x93, 0x52, 0xed, 0xdd, 0xc3, 0xd6, 0xb2, 0x90, 0xf8, 0xf9,
	0xf8, 0x8f, 0x4b, 0xf6, 0x0e, 0x6c, 0x44, 0x43, 0x18, 0x99, 0x21, 0x6f, 0xf5, 0x5a, 0x83, 0x8d,
	0xf1, 0xf3, 0x08, 0xce, 0x65, 0x86, 0xec, 0x03, 0xd8, 0xd4, 0x26, 0x60, 0x7e, 0x2d, 0x55, 0x6d,
	0x3c, 0x22, 0xa3, 0xbb, 0xa0, 0xa4, 0xbd, 0x80, 0xc7, 0xa5, 0xd3, 0xfc, 0x71, 0xaf, 0x35, 0xe8,
	0x8e, 0xe3, 0x92, 0x88, 0xd2, 0xfc, 0x49, 0x4d, 0x94, 0xde, 0xfb, 0xbb, 0x05, 0x6f, 0x37, 0xcf,
	0x56, 0x98, 0xa6, 0xc2, 0x49, 0x75, 0xab, 0xcd, 0x4c, 0x68, 0x73, 0x6d, 0xd9, 0x09, 0xf4, 0x53,
	0xab, 0x64, 0x2a, 0x32, 0x39, 0xa7, 0xd8, 0x53, 0x8e, 0x53, 0xe1, 0x30, 0xaf, 0x96, 0x81, 0x76,
	0xd7, 0x1d, 0xbf, 0x4b, 0xe2, 0x99, 0x9c, 0x1f, 0x46, 0xed, 0x82, 0xac, 0x0b, 0xcc, 0x69, 0x11,
	0xd8, 0x6f, 0xf0, 0xa1, 0xc1, 0x99, 0x0d, 0x5a, 0x06, 0x9c, 0xfe, 0x67, 0xb5, 0x47, 0x54, 0xad,
	0xbf, 0xb4, 0x1f, 0x2a, 0x79, 0x00, 0x2f, 0x9b, 0x3a, 0x4d, 0x99, 0xa0, 0x33, 0xb4, 0x45, 0xa8,
	0xdf, 0x78, 0x37, 0xab, 0xee, 0xac, 0x6e, 0xbc, 0xaa, 0xb2, 0xbd, 0xbf, 0xde, 0x00, 0x58, 0x36,
	0x3b, 0xb6, 0x32, 0x93, 0xda, 0x88, 0x45, 0xe7, 0xf8, 0xb0, 0x6a, 0x65, 0xa4, 0xa7, 0x0d, 0x64,
	0xef, 0x43, 0xd7, 0x17, 0x93, 0x15, 0x6b, 0x44, 0x56, 0xc7, 0x17, 0x93, 0xa5, 0xd4, 0x87, 0x4e,
	0xa9, 0x56, 0x9c, 0xd7, 0xe4, 0xb4, 0x4b, 0xb5, 0x54, 0xde, 0x84, 0x67, 0xa5, 0xd3, 0x62, 0x9e,
	0xf3, 0x03, 0xda, 0xe3, 0xd3, 0xd2, 0xe9, 0xdf, 0x73, 0xc2, 0x8a, 0xf0, 0x27, 0x35, 0x56, 0x11,
	0x33, 0x78, 0x12, 0xee, 0x1d, 0xf2, 0x4f, 0xa9, 0x10, 0xad, 0xd9, 0x3e, 0x74, 0xd1, 0x28, 0xe9,
	0x7c, 0x91, 0xca, 0xa0, 0xad, 0xe1, 0x9f, 0x55, 0xfb, 0x5d, 0x83, 0x6c, 0x17, 0x9e, 0xfa, 0x1b,
	0xe9, 0x90, 0x7f, 0x4e, 0x69, 0x75, 0xc1, 0xf6, 0x61, 0xd3, 0xa1, 0xbc, 0xad, 0x7a, 0x96, 0xcb,
	0x80, 0xfc, 0x0b, 0x7a, 0x5c, 0x27, 0xd2, 0xd8, 0xaa, 0xb1, 0x0c, 0xc8, 0x12, 0xd8, 0xf1, 0x85,
	0x0f, 0x52, 0x1b, 0x9c, 0xae, 0xa8, 0x5f, 0x92, 0xba, 0xbd, 0x88, 0x16, 0xfe, 0x7b, 0x00, 0x93,
	0x22, 0xf7, 0xa1, 0xd2, 0xbe, 0x22, 0x6d, 0x83, 0x08, 0xc5, 0x1f, 0xc3, 0x6e, 0xb5, 0x37, 0xa1,
	0xcd, 0x0d, 0xe6, 0x3a, 0x88, 0x14, 0x4b, 0x4c, 0xf9, 0xd7, 0xb4, 0x33, 0x56, 0x65, 0xa7, 0x55,
	0xf4, 0x4b, 0x4c, 0xd8, 0x47, 0xb0, 0x7d, 0x67, 0xff, 0xad, 0x7f, 0x43, 0xfa, 0xd6, 0x9d, 0x5d,
	0x77, 0xfb, 0xd0, 0x09, 0xb9, 0x34, 0x3e, 0xd3, 0x41, 0x64, 0xa1, 0xe0, 0xdf, 0xd2, 0xe3, 0xdb,
	0x0d, 0x3b, 0x0b, 0x05, 0x7b, 0x05, 0xed, 0x1c, 0x15, 0xea, 0x12, 0xc9, 0xf8, 0x8e, 0x0c, 0xa8,
	0x51, 0x14, 0xf6, 0x61, 0xb3, 0x54, 0xc2, 0x9a, 0xd2, 0x89, 0x50, 0x18, 0x83, 0x29, 0xff, 0xbe,
	0xd7, 0x1a, 0x3c, 0x1f, 0x77, 0x4a, 0xf5, 0xab, 0x29, 0xdd, 0x15, 0x31, 0xf6, 0x1a, 0x5e, 0x92,
	0x25, 0xdc, 0xd0, 0x89, 0xf5, 0x8f, 0xe1, 0x07, 0xd2, 0x77, 0xa2, 0x7e, 0x31, 0x74, 0x97, 0xab,
	0xdf, 0xc4, 0x2b, 0x68, 0xc7, 0xe1, 0x17, 0x3e, 0xc8, 0x50, 0x78, 0x7e, 0x48, 0x26, 0x44, 0x74,
	0x49, 0x24, 0x0a, 0x32, 0xd3, 0xa6, 0x11, 0x8e, 0x2a, 0x21, 0xa2, 0x5a, 0x68, 0x86, 0xdd, 0xc8,
	0x94, 0x24, 0xe4, 0xc7, 0x2b, 0xc3, 0x6e, 0x64, 0x1a, 0x3d, 0x64, 0x23, 0x78, 0x2b, 0x95, 0x3e,
	0x54, 0x8a, 0x50, 0x37, 0xd2, 0xcc, 0x90, 0x06, 0x82, 0xff, 0x48, 0x2f, 0xbc, 0x13, 0x53, 0x52,
	0x0f, 0x29, 0x8b, 0xf3, 0xc0, 0xfe, 0x84, 0xed, 0xb5, 0xa1, 0x9f, 0xca, 0x20, 0xf9, 0x49, 0xaf,
	0x35, 0x68, 0x0f, 0xcf, 0x93, 0xff, 0x71, 0x92, 0x25, 0x0f, 0x1e, 0x25, 0xe3, 0x2d, 0x55, 0x8f,
	0xa2, 0x36, 0xb3, 0x23, 0x19, 0x64, 0x3c, 0xe1, 0x02, 0xfa, 0x20, 0x32, 0x3b, 0x45, 0xfe, 0x53,
	0x75, 0xc2, 0x45, 0x70, 0x66, 0xa7, 0x38, 0x79, 0x46, 0xa7, 0xe9, 0xe8, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9f, 0x5a, 0xb3, 0xe4, 0x61, 0x05, 0x00, 0x00,
}