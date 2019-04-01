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
// source: rib_edm_proto_route_summ_detail.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_rib_table_ids_rib_table_id_summary_protos_summary_proto

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

type RibEdmProtoRouteSummDetail_KEYS struct {
	Tableid              string   `protobuf:"bytes,1,opt,name=tableid,proto3" json:"tableid,omitempty"`
	Protoid              uint32   `protobuf:"varint,2,opt,name=protoid,proto3" json:"protoid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RibEdmProtoRouteSummDetail_KEYS) Reset()         { *m = RibEdmProtoRouteSummDetail_KEYS{} }
func (m *RibEdmProtoRouteSummDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*RibEdmProtoRouteSummDetail_KEYS) ProtoMessage()    {}
func (*RibEdmProtoRouteSummDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67f6a6d08872ebe, []int{0}
}

func (m *RibEdmProtoRouteSummDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS.Unmarshal(m, b)
}
func (m *RibEdmProtoRouteSummDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *RibEdmProtoRouteSummDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS.Merge(m, src)
}
func (m *RibEdmProtoRouteSummDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS.Size(m)
}
func (m *RibEdmProtoRouteSummDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS proto.InternalMessageInfo

func (m *RibEdmProtoRouteSummDetail_KEYS) GetTableid() string {
	if m != nil {
		return m.Tableid
	}
	return ""
}

func (m *RibEdmProtoRouteSummDetail_KEYS) GetProtoid() uint32 {
	if m != nil {
		return m.Protoid
	}
	return 0
}

type RibEdmRouteCount struct {
	ActiveRoutesCount    uint32   `protobuf:"varint,1,opt,name=active_routes_count,json=activeRoutesCount,proto3" json:"active_routes_count,omitempty"`
	NumBackupRoutes      uint32   `protobuf:"varint,2,opt,name=num_backup_routes,json=numBackupRoutes,proto3" json:"num_backup_routes,omitempty"`
	NumActivePaths       uint32   `protobuf:"varint,3,opt,name=num_active_paths,json=numActivePaths,proto3" json:"num_active_paths,omitempty"`
	NumBackupPaths       uint32   `protobuf:"varint,4,opt,name=num_backup_paths,json=numBackupPaths,proto3" json:"num_backup_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RibEdmRouteCount) Reset()         { *m = RibEdmRouteCount{} }
func (m *RibEdmRouteCount) String() string { return proto.CompactTextString(m) }
func (*RibEdmRouteCount) ProtoMessage()    {}
func (*RibEdmRouteCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67f6a6d08872ebe, []int{1}
}

func (m *RibEdmRouteCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RibEdmRouteCount.Unmarshal(m, b)
}
func (m *RibEdmRouteCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RibEdmRouteCount.Marshal(b, m, deterministic)
}
func (m *RibEdmRouteCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RibEdmRouteCount.Merge(m, src)
}
func (m *RibEdmRouteCount) XXX_Size() int {
	return xxx_messageInfo_RibEdmRouteCount.Size(m)
}
func (m *RibEdmRouteCount) XXX_DiscardUnknown() {
	xxx_messageInfo_RibEdmRouteCount.DiscardUnknown(m)
}

var xxx_messageInfo_RibEdmRouteCount proto.InternalMessageInfo

func (m *RibEdmRouteCount) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *RibEdmRouteCount) GetNumBackupRoutes() uint32 {
	if m != nil {
		return m.NumBackupRoutes
	}
	return 0
}

func (m *RibEdmRouteCount) GetNumActivePaths() uint32 {
	if m != nil {
		return m.NumActivePaths
	}
	return 0
}

func (m *RibEdmRouteCount) GetNumBackupPaths() uint32 {
	if m != nil {
		return m.NumBackupPaths
	}
	return 0
}

type RibEdmProtoRouteSummDetail struct {
	Name                 string            `protobuf:"bytes,50,opt,name=name,proto3" json:"name,omitempty"`
	Instance             string            `protobuf:"bytes,51,opt,name=instance,proto3" json:"instance,omitempty"`
	ProtoRouteCount      *RibEdmRouteCount `protobuf:"bytes,52,opt,name=proto_route_count,json=protoRouteCount,proto3" json:"proto_route_count,omitempty"`
	RtypeNone            *RibEdmRouteCount `protobuf:"bytes,53,opt,name=rtype_none,json=rtypeNone,proto3" json:"rtype_none,omitempty"`
	RtypeOther           *RibEdmRouteCount `protobuf:"bytes,54,opt,name=rtype_other,json=rtypeOther,proto3" json:"rtype_other,omitempty"`
	RtypeOspfIntra       *RibEdmRouteCount `protobuf:"bytes,55,opt,name=rtype_ospf_intra,json=rtypeOspfIntra,proto3" json:"rtype_ospf_intra,omitempty"`
	RtypeOspfInter       *RibEdmRouteCount `protobuf:"bytes,56,opt,name=rtype_ospf_inter,json=rtypeOspfInter,proto3" json:"rtype_ospf_inter,omitempty"`
	RtypeOspfExtern1     *RibEdmRouteCount `protobuf:"bytes,57,opt,name=rtype_ospf_extern1,json=rtypeOspfExtern1,proto3" json:"rtype_ospf_extern1,omitempty"`
	RtypeOspfExtern2     *RibEdmRouteCount `protobuf:"bytes,58,opt,name=rtype_ospf_extern2,json=rtypeOspfExtern2,proto3" json:"rtype_ospf_extern2,omitempty"`
	RtypeIsisSum         *RibEdmRouteCount `protobuf:"bytes,59,opt,name=rtype_isis_sum,json=rtypeIsisSum,proto3" json:"rtype_isis_sum,omitempty"`
	RtypeIsisL1          *RibEdmRouteCount `protobuf:"bytes,60,opt,name=rtype_isis_l1,json=rtypeIsisL1,proto3" json:"rtype_isis_l1,omitempty"`
	RtypeIsisL2          *RibEdmRouteCount `protobuf:"bytes,61,opt,name=rtype_isis_l2,json=rtypeIsisL2,proto3" json:"rtype_isis_l2,omitempty"`
	RtypeIsisL1Ia        *RibEdmRouteCount `protobuf:"bytes,62,opt,name=rtype_isis_l1_ia,json=rtypeIsisL1Ia,proto3" json:"rtype_isis_l1_ia,omitempty"`
	RtypeBgpInt          *RibEdmRouteCount `protobuf:"bytes,63,opt,name=rtype_bgp_int,json=rtypeBgpInt,proto3" json:"rtype_bgp_int,omitempty"`
	RtypeBgpExt          *RibEdmRouteCount `protobuf:"bytes,64,opt,name=rtype_bgp_ext,json=rtypeBgpExt,proto3" json:"rtype_bgp_ext,omitempty"`
	RtypeBgpLoc          *RibEdmRouteCount `protobuf:"bytes,65,opt,name=rtype_bgp_loc,json=rtypeBgpLoc,proto3" json:"rtype_bgp_loc,omitempty"`
	RtypeOspfNssa1       *RibEdmRouteCount `protobuf:"bytes,66,opt,name=rtype_ospf_nssa1,json=rtypeOspfNssa1,proto3" json:"rtype_ospf_nssa1,omitempty"`
	RtypeOspfNssa2       *RibEdmRouteCount `protobuf:"bytes,67,opt,name=rtype_ospf_nssa2,json=rtypeOspfNssa2,proto3" json:"rtype_ospf_nssa2,omitempty"`
	RtypeIgrp2Int        *RibEdmRouteCount `protobuf:"bytes,68,opt,name=rtype_igrp2_int,json=rtypeIgrp2Int,proto3" json:"rtype_igrp2_int,omitempty"`
	RtypeIgrp2Ext        *RibEdmRouteCount `protobuf:"bytes,69,opt,name=rtype_igrp2_ext,json=rtypeIgrp2Ext,proto3" json:"rtype_igrp2_ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RibEdmProtoRouteSummDetail) Reset()         { *m = RibEdmProtoRouteSummDetail{} }
func (m *RibEdmProtoRouteSummDetail) String() string { return proto.CompactTextString(m) }
func (*RibEdmProtoRouteSummDetail) ProtoMessage()    {}
func (*RibEdmProtoRouteSummDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67f6a6d08872ebe, []int{2}
}

func (m *RibEdmProtoRouteSummDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail.Unmarshal(m, b)
}
func (m *RibEdmProtoRouteSummDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail.Marshal(b, m, deterministic)
}
func (m *RibEdmProtoRouteSummDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RibEdmProtoRouteSummDetail.Merge(m, src)
}
func (m *RibEdmProtoRouteSummDetail) XXX_Size() int {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail.Size(m)
}
func (m *RibEdmProtoRouteSummDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RibEdmProtoRouteSummDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RibEdmProtoRouteSummDetail proto.InternalMessageInfo

func (m *RibEdmProtoRouteSummDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RibEdmProtoRouteSummDetail) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *RibEdmProtoRouteSummDetail) GetProtoRouteCount() *RibEdmRouteCount {
	if m != nil {
		return m.ProtoRouteCount
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeNone() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeNone
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOther() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOther
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfIntra() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfIntra
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfInter() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfInter
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfExtern1() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfExtern1
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfExtern2() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfExtern2
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIsisSum() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIsisSum
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIsisL1() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIsisL1
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIsisL2() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIsisL2
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIsisL1Ia() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIsisL1Ia
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeBgpInt() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeBgpInt
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeBgpExt() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeBgpExt
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeBgpLoc() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeBgpLoc
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfNssa1() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfNssa1
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfNssa2() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfNssa2
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIgrp2Int() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIgrp2Int
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIgrp2Ext() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIgrp2Ext
	}
	return nil
}

func init() {
	proto.RegisterType((*RibEdmProtoRouteSummDetail_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.rib_table_ids.rib_table_id.summary_protos.summary_proto.rib_edm_proto_route_summ_detail_KEYS")
	proto.RegisterType((*RibEdmRouteCount)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.rib_table_ids.rib_table_id.summary_protos.summary_proto.rib_edm_route_count")
	proto.RegisterType((*RibEdmProtoRouteSummDetail)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.rib_table_ids.rib_table_id.summary_protos.summary_proto.rib_edm_proto_route_summ_detail")
}

func init() {
	proto.RegisterFile("rib_edm_proto_route_summ_detail.proto", fileDescriptor_c67f6a6d08872ebe)
}

var fileDescriptor_c67f6a6d08872ebe = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0x5f, 0x6b, 0x53, 0x31,
	0x18, 0xc6, 0x89, 0x0e, 0x75, 0x99, 0x5d, 0xdb, 0x0c, 0x24, 0x78, 0x63, 0x29, 0x0a, 0xc5, 0x8b,
	0x03, 0x3d, 0x9b, 0xff, 0xff, 0xae, 0xb3, 0x17, 0xc5, 0x31, 0xa5, 0xbb, 0x9a, 0x37, 0x21, 0xe7,
	0x34, 0xeb, 0x82, 0x3d, 0x49, 0x48, 0x72, 0x46, 0xf7, 0x05, 0xbc, 0x13, 0x41, 0x44, 0x45, 0xfc,
	0x3a, 0x7e, 0x2f, 0x49, 0x72, 0x5a, 0xba, 0x5a, 0xd8, 0x95, 0x64, 0x17, 0x85, 0xbe, 0x6f, 0x9e,
	0x3c, 0xe7, 0xc7, 0x93, 0xe4, 0x85, 0xf7, 0x34, 0xcf, 0x08, 0x1b, 0x15, 0x44, 0x69, 0x69, 0x25,
	0xd1, 0xb2, 0xb4, 0x8c, 0x98, 0xb2, 0x28, 0xc8, 0x88, 0x59, 0xca, 0x27, 0x89, 0xef, 0xa3, 0xa3,
	0x9c, 0x9b, 0x5c, 0x12, 0x2e, 0x0d, 0x99, 0x6a, 0xc2, 0x15, 0x71, 0xdb, 0xb8, 0x3a, 0xdd, 0x21,
	0x52, 0x31, 0x9d, 0x68, 0x9e, 0xb9, 0x1f, 0xb1, 0x34, 0x9b, 0x30, 0xc2, 0x47, 0xe6, 0x5c, 0x95,
	0x38, 0x3f, 0xaa, 0xcf, 0xc2, 0x37, 0xcc, 0xf9, 0xb2, 0xfd, 0x01, 0xde, 0xbd, 0x80, 0x81, 0xbc,
	0xed, 0x1f, 0x1d, 0x22, 0x0c, 0xaf, 0x7b, 0x43, 0x3e, 0xc2, 0xa0, 0x05, 0x3a, 0xeb, 0xc3, 0x59,
	0xe9, 0x56, 0xfc, 0x4e, 0x3e, 0xc2, 0x57, 0x5a, 0xa0, 0x53, 0x1b, 0xce, 0xca, 0xf6, 0x1f, 0x00,
	0xb7, 0x66, 0xe6, 0xc1, 0x36, 0x97, 0xa5, 0xb0, 0x28, 0x81, 0x5b, 0x34, 0xb7, 0xfc, 0x94, 0x85,
	0xae, 0x09, 0x6d, 0xef, 0x5b, 0x1b, 0x36, 0xc3, 0xd2, 0xd0, 0xaf, 0xec, 0x79, 0xfd, 0x7d, 0xd8,
	0x14, 0x65, 0x41, 0x32, 0x9a, 0x7f, 0x2c, 0x55, 0xb5, 0xa7, 0xfa, 0x56, 0x5d, 0x94, 0x45, 0xcf,
	0xf7, 0xc3, 0x06, 0xd4, 0x81, 0x0d, 0xa7, 0xad, 0xfc, 0x15, 0xb5, 0x27, 0x06, 0x5f, 0xf5, 0xd2,
	0x4d, 0x51, 0x16, 0xbb, 0xbe, 0xfd, 0xde, 0x75, 0x67, 0xca, 0xca, 0x35, 0x28, 0xd7, 0xe6, 0xca,
	0x60, 0xea, 0x95, 0xed, 0x4f, 0xb7, 0xe0, 0x9d, 0x0b, 0x42, 0x42, 0x08, 0xae, 0x09, 0x5a, 0x30,
	0x9c, 0xfa, 0x70, 0xfc, 0x7f, 0x74, 0x1b, 0xde, 0xe0, 0xc2, 0x58, 0x2a, 0x72, 0x86, 0xb7, 0x7d,
	0x7f, 0x5e, 0xa3, 0x5f, 0x00, 0x36, 0x17, 0xbd, 0x42, 0x04, 0x3b, 0x2d, 0xd0, 0xd9, 0x48, 0x45,
	0xf2, 0xdf, 0xce, 0x3b, 0x59, 0x71, 0x1e, 0xc3, 0xba, 0x5f, 0xf2, 0x01, 0x86, 0xc0, 0x3f, 0x03,
	0x08, 0xb5, 0x3d, 0x53, 0x8c, 0x08, 0x29, 0x18, 0x7e, 0x10, 0x85, 0x6a, 0xdd, 0x13, 0x1c, 0x48,
	0xc1, 0xd0, 0x17, 0x00, 0x37, 0x02, 0x8f, 0xb4, 0x27, 0x4c, 0xe3, 0x87, 0x51, 0x80, 0x42, 0x24,
	0xef, 0x1c, 0x01, 0xfa, 0x09, 0x60, 0xa3, 0x22, 0x32, 0xea, 0x98, 0x70, 0x61, 0x35, 0xc5, 0x8f,
	0xa2, 0x60, 0x6d, 0x06, 0x2c, 0xa3, 0x8e, 0x07, 0x8e, 0x62, 0x05, 0x1a, 0xd3, 0xf8, 0x71, 0x7c,
	0x34, 0xa6, 0xd1, 0x6f, 0x00, 0xd1, 0x02, 0x1a, 0x9b, 0x5a, 0xa6, 0x45, 0x17, 0x3f, 0x89, 0x02,
	0xd7, 0x98, 0xc3, 0xf5, 0x03, 0xc7, 0x6a, 0xbc, 0x14, 0x3f, 0xbd, 0x14, 0x78, 0x29, 0xfa, 0x06,
	0x60, 0x08, 0x94, 0x70, 0xc3, 0x8d, 0x9b, 0x3e, 0xf8, 0x59, 0x14, 0xb4, 0x9b, 0x9e, 0x62, 0x60,
	0xb8, 0x39, 0x2c, 0x0b, 0xf4, 0x15, 0xc0, 0xda, 0x02, 0xd6, 0xa4, 0x8b, 0x9f, 0x47, 0xa1, 0xda,
	0x98, 0x53, 0xed, 0x77, 0xff, 0x81, 0x4a, 0xf1, 0x8b, 0xd8, 0x50, 0x29, 0xfa, 0x31, 0x7f, 0x99,
	0x55, 0x52, 0x84, 0x53, 0xfc, 0x32, 0x0a, 0x57, 0x6d, 0x21, 0xac, 0x01, 0x5d, 0x88, 0x2b, 0x1b,
	0x2b, 0x37, 0x32, 0xf0, 0xab, 0x88, 0x71, 0xf5, 0xc6, 0x6a, 0x20, 0xec, 0x12, 0x14, 0x9b, 0x5a,
	0xfc, 0x3a, 0x2e, 0x54, 0x7f, 0xba, 0x0c, 0x35, 0x91, 0x39, 0xde, 0x8d, 0x0b, 0xb5, 0x2f, 0xf3,
	0xe5, 0x91, 0x2f, 0x8c, 0xa1, 0x5d, 0xdc, 0x8b, 0x3c, 0xf2, 0x0f, 0x1c, 0xc5, 0x2a, 0xb4, 0x14,
	0xef, 0x5d, 0x02, 0xb4, 0x14, 0x7d, 0x07, 0xb0, 0x5e, 0x3d, 0xc7, 0xb1, 0x56, 0xa9, 0xbf, 0xf6,
	0x6f, 0x62, 0xbe, 0x46, 0x47, 0xe1, 0x2e, 0xfe, 0x32, 0x98, 0xbb, 0xfa, 0xfd, 0xd8, 0x60, 0xfd,
	0xa9, 0xcd, 0xae, 0x79, 0xe5, 0xf6, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x70, 0xfd, 0x5f,
	0xb7, 0x0c, 0x00, 0x00,
}
