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
// source: csi_ma_node_info.proto

package cisco_ios_xr_cofo_csi_ma_oper_csima_nodes_node_rack_ids_rack_id_slot_id

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

type CsiMaNodeInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	RackId               uint32   `protobuf:"varint,2,opt,name=rack_id,json=rackId,proto3" json:"rack_id,omitempty"`
	SlotId               uint32   `protobuf:"varint,3,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsiMaNodeInfo_KEYS) Reset()         { *m = CsiMaNodeInfo_KEYS{} }
func (m *CsiMaNodeInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*CsiMaNodeInfo_KEYS) ProtoMessage()    {}
func (*CsiMaNodeInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_677301b1329f36d1, []int{0}
}

func (m *CsiMaNodeInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsiMaNodeInfo_KEYS.Unmarshal(m, b)
}
func (m *CsiMaNodeInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsiMaNodeInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *CsiMaNodeInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsiMaNodeInfo_KEYS.Merge(m, src)
}
func (m *CsiMaNodeInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_CsiMaNodeInfo_KEYS.Size(m)
}
func (m *CsiMaNodeInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_CsiMaNodeInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_CsiMaNodeInfo_KEYS proto.InternalMessageInfo

func (m *CsiMaNodeInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *CsiMaNodeInfo_KEYS) GetRackId() uint32 {
	if m != nil {
		return m.RackId
	}
	return 0
}

func (m *CsiMaNodeInfo_KEYS) GetSlotId() uint32 {
	if m != nil {
		return m.SlotId
	}
	return 0
}

type CsiMaSlice_ struct {
	SliceNodeId          uint32   `protobuf:"varint,1,opt,name=slice_node_id,json=sliceNodeId,proto3" json:"slice_node_id,omitempty"`
	AdminUp              bool     `protobuf:"varint,2,opt,name=admin_up,json=adminUp,proto3" json:"admin_up,omitempty"`
	OperUp               bool     `protobuf:"varint,3,opt,name=oper_up,json=operUp,proto3" json:"oper_up,omitempty"`
	Pic                  uint64   `protobuf:"varint,4,opt,name=pic,proto3" json:"pic,omitempty"`
	CsiPicArr            []uint64 `protobuf:"varint,5,rep,packed,name=csi_pic_arr,json=csiPicArr,proto3" json:"csi_pic_arr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsiMaSlice_) Reset()         { *m = CsiMaSlice_{} }
func (m *CsiMaSlice_) String() string { return proto.CompactTextString(m) }
func (*CsiMaSlice_) ProtoMessage()    {}
func (*CsiMaSlice_) Descriptor() ([]byte, []int) {
	return fileDescriptor_677301b1329f36d1, []int{1}
}

func (m *CsiMaSlice_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsiMaSlice_.Unmarshal(m, b)
}
func (m *CsiMaSlice_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsiMaSlice_.Marshal(b, m, deterministic)
}
func (m *CsiMaSlice_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsiMaSlice_.Merge(m, src)
}
func (m *CsiMaSlice_) XXX_Size() int {
	return xxx_messageInfo_CsiMaSlice_.Size(m)
}
func (m *CsiMaSlice_) XXX_DiscardUnknown() {
	xxx_messageInfo_CsiMaSlice_.DiscardUnknown(m)
}

var xxx_messageInfo_CsiMaSlice_ proto.InternalMessageInfo

func (m *CsiMaSlice_) GetSliceNodeId() uint32 {
	if m != nil {
		return m.SliceNodeId
	}
	return 0
}

func (m *CsiMaSlice_) GetAdminUp() bool {
	if m != nil {
		return m.AdminUp
	}
	return false
}

func (m *CsiMaSlice_) GetOperUp() bool {
	if m != nil {
		return m.OperUp
	}
	return false
}

func (m *CsiMaSlice_) GetPic() uint64 {
	if m != nil {
		return m.Pic
	}
	return 0
}

func (m *CsiMaSlice_) GetCsiPicArr() []uint64 {
	if m != nil {
		return m.CsiPicArr
	}
	return nil
}

type CsiMaNodeInfo struct {
	RackSlotId           uint64         `protobuf:"varint,50,opt,name=rack_slot_id,json=rackSlotId,proto3" json:"rack_slot_id,omitempty"`
	NodeId               uint32         `protobuf:"varint,51,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeUp               bool           `protobuf:"varint,52,opt,name=node_up,json=nodeUp,proto3" json:"node_up,omitempty"`
	SliceArr             []*CsiMaSlice_ `protobuf:"bytes,53,rep,name=slice_arr,json=sliceArr,proto3" json:"slice_arr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CsiMaNodeInfo) Reset()         { *m = CsiMaNodeInfo{} }
func (m *CsiMaNodeInfo) String() string { return proto.CompactTextString(m) }
func (*CsiMaNodeInfo) ProtoMessage()    {}
func (*CsiMaNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_677301b1329f36d1, []int{2}
}

func (m *CsiMaNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsiMaNodeInfo.Unmarshal(m, b)
}
func (m *CsiMaNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsiMaNodeInfo.Marshal(b, m, deterministic)
}
func (m *CsiMaNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsiMaNodeInfo.Merge(m, src)
}
func (m *CsiMaNodeInfo) XXX_Size() int {
	return xxx_messageInfo_CsiMaNodeInfo.Size(m)
}
func (m *CsiMaNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CsiMaNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CsiMaNodeInfo proto.InternalMessageInfo

func (m *CsiMaNodeInfo) GetRackSlotId() uint64 {
	if m != nil {
		return m.RackSlotId
	}
	return 0
}

func (m *CsiMaNodeInfo) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *CsiMaNodeInfo) GetNodeUp() bool {
	if m != nil {
		return m.NodeUp
	}
	return false
}

func (m *CsiMaNodeInfo) GetSliceArr() []*CsiMaSlice_ {
	if m != nil {
		return m.SliceArr
	}
	return nil
}

func init() {
	proto.RegisterType((*CsiMaNodeInfo_KEYS)(nil), "cisco_ios_xr_cofo_csi_ma_oper.csima.nodes.node.rack_ids.rack_id.slot_id.csi_ma_node_info_KEYS")
	proto.RegisterType((*CsiMaSlice_)(nil), "cisco_ios_xr_cofo_csi_ma_oper.csima.nodes.node.rack_ids.rack_id.slot_id.csi_ma_slice_")
	proto.RegisterType((*CsiMaNodeInfo)(nil), "cisco_ios_xr_cofo_csi_ma_oper.csima.nodes.node.rack_ids.rack_id.slot_id.csi_ma_node_info")
}

func init() { proto.RegisterFile("csi_ma_node_info.proto", fileDescriptor_677301b1329f36d1) }

var fileDescriptor_677301b1329f36d1 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x6a, 0xe3, 0x30,
	0x14, 0xc4, 0xeb, 0x6c, 0x62, 0x2b, 0x6b, 0x08, 0x82, 0xdd, 0x75, 0x29, 0x14, 0xe3, 0x93, 0x4f,
	0x3a, 0x24, 0xed, 0x07, 0xe4, 0x50, 0x4a, 0x28, 0x84, 0xa2, 0x90, 0x42, 0x4f, 0x42, 0x95, 0x1c,
	0x78, 0x34, 0xb6, 0x84, 0x94, 0x40, 0x7f, 0xa5, 0xdf, 0xd6, 0x9f, 0x29, 0x4f, 0xb2, 0x0f, 0xcd,
	0xb9, 0x17, 0xfb, 0xcd, 0x8c, 0xfd, 0x66, 0x46, 0x36, 0xf9, 0xa7, 0x3c, 0x88, 0x4e, 0x8a, 0xde,
	0xe8, 0x56, 0x40, 0x7f, 0x30, 0xcc, 0x3a, 0x73, 0x32, 0xf4, 0x41, 0x81, 0x57, 0x46, 0x80, 0xf1,
	0xe2, 0xdd, 0x09, 0x65, 0x0e, 0x46, 0x0c, 0x4f, 0x1a, 0xdb, 0x3a, 0xa6, 0x3c, 0x74, 0x92, 0xe1,
	0x4b, 0x3e, 0x5c, 0x99, 0x93, 0xea, 0x4d, 0x80, 0xf6, 0xe3, 0xc0, 0xfc, 0xd1, 0x9c, 0x04, 0xe8,
	0xfa, 0x40, 0xfe, 0x5e, 0x5a, 0x88, 0xc7, 0xfb, 0x97, 0x1d, 0xbd, 0x26, 0x79, 0x60, 0x7a, 0xd9,
	0xb5, 0x65, 0x52, 0x25, 0x4d, 0xce, 0x33, 0x24, 0xb6, 0xb2, 0x6b, 0xe9, 0x7f, 0x32, 0x1b, 0x16,
	0x95, 0xbf, 0xaa, 0xa4, 0x29, 0xf8, 0x14, 0xe1, 0x46, 0xa3, 0x30, 0x6c, 0x2e, 0xd3, 0x28, 0x20,
	0xdc, 0xe8, 0xfa, 0x23, 0x21, 0xc5, 0x60, 0xe4, 0x8f, 0xa0, 0x5a, 0x41, 0x6b, 0x52, 0xc4, 0x29,
	0x1a, 0xeb, 0x60, 0x52, 0xf0, 0x79, 0x20, 0xb7, 0x46, 0xb7, 0x1b, 0x4d, 0xaf, 0x48, 0x26, 0x75,
	0x07, 0xbd, 0x38, 0xdb, 0x60, 0x94, 0xf1, 0x59, 0xc0, 0x7b, 0x8b, 0x4e, 0x58, 0x15, 0x95, 0x34,
	0x28, 0x53, 0x84, 0x7b, 0x4b, 0x17, 0x24, 0xb5, 0xa0, 0xca, 0x49, 0x95, 0x34, 0x13, 0x8e, 0x23,
	0xbd, 0x21, 0x73, 0xb4, 0xb6, 0xa0, 0x84, 0x74, 0xae, 0xfc, 0x5d, 0xa5, 0xcd, 0x84, 0xe7, 0xca,
	0xc3, 0x13, 0xa8, 0xb5, 0x73, 0xf5, 0x67, 0x42, 0x16, 0x97, 0x87, 0x40, 0x2b, 0xf2, 0x27, 0x54,
	0x1c, 0xeb, 0x2c, 0xc3, 0x3e, 0x82, 0xdc, 0x2e, 0x54, 0xc2, 0x04, 0x63, 0xf4, 0x55, 0xec, 0xda,
	0xc7, 0xd4, 0xa3, 0x70, 0xb6, 0xe5, 0x6d, 0x8c, 0x86, 0x70, 0x6f, 0xa9, 0x27, 0x79, 0xac, 0x8c,
	0x31, 0xee, 0xaa, 0xb4, 0x99, 0x2f, 0x9f, 0xd9, 0x0f, 0x7d, 0x49, 0xf6, 0xed, 0x74, 0x79, 0x16,
	0xee, 0x6b, 0xe7, 0x5e, 0xa7, 0xe1, 0x8f, 0x59, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x48, 0xa1,
	0x1f, 0x03, 0x4b, 0x02, 0x00, 0x00,
}