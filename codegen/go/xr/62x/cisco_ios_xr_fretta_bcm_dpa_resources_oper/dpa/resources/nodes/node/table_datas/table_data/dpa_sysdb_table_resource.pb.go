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
// source: dpa_sysdb_table_resource.proto

package cisco_ios_xr_fretta_bcm_dpa_resources_oper_dpa_resources_nodes_node_table_datas_table_data

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

type DpaSysdbTableResource_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Resource             uint32   `protobuf:"varint,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DpaSysdbTableResource_KEYS) Reset()         { *m = DpaSysdbTableResource_KEYS{} }
func (m *DpaSysdbTableResource_KEYS) String() string { return proto.CompactTextString(m) }
func (*DpaSysdbTableResource_KEYS) ProtoMessage()    {}
func (*DpaSysdbTableResource_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b72bbb1ac20a6d8, []int{0}
}

func (m *DpaSysdbTableResource_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaSysdbTableResource_KEYS.Unmarshal(m, b)
}
func (m *DpaSysdbTableResource_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaSysdbTableResource_KEYS.Marshal(b, m, deterministic)
}
func (m *DpaSysdbTableResource_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaSysdbTableResource_KEYS.Merge(m, src)
}
func (m *DpaSysdbTableResource_KEYS) XXX_Size() int {
	return xxx_messageInfo_DpaSysdbTableResource_KEYS.Size(m)
}
func (m *DpaSysdbTableResource_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaSysdbTableResource_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_DpaSysdbTableResource_KEYS proto.InternalMessageInfo

func (m *DpaSysdbTableResource_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *DpaSysdbTableResource_KEYS) GetResource() uint32 {
	if m != nil {
		return m.Resource
	}
	return 0
}

type DpaSysdbNpuTableResource struct {
	NpuId                uint32   `protobuf:"varint,1,opt,name=npu_id,json=npuId,proto3" json:"npu_id,omitempty"`
	NumObjs              uint32   `protobuf:"varint,2,opt,name=num_objs,json=numObjs,proto3" json:"num_objs,omitempty"`
	CreateReq            uint32   `protobuf:"varint,3,opt,name=create_req,json=createReq,proto3" json:"create_req,omitempty"`
	CreateOk             uint32   `protobuf:"varint,4,opt,name=create_ok,json=createOk,proto3" json:"create_ok,omitempty"`
	DeleteReq            uint32   `protobuf:"varint,5,opt,name=delete_req,json=deleteReq,proto3" json:"delete_req,omitempty"`
	DeleteOk             uint32   `protobuf:"varint,6,opt,name=delete_ok,json=deleteOk,proto3" json:"delete_ok,omitempty"`
	UpdateReq            uint32   `protobuf:"varint,7,opt,name=update_req,json=updateReq,proto3" json:"update_req,omitempty"`
	UpdateOk             uint32   `protobuf:"varint,8,opt,name=update_ok,json=updateOk,proto3" json:"update_ok,omitempty"`
	EodReq               uint32   `protobuf:"varint,9,opt,name=eod_req,json=eodReq,proto3" json:"eod_req,omitempty"`
	EodOk                uint32   `protobuf:"varint,10,opt,name=eod_ok,json=eodOk,proto3" json:"eod_ok,omitempty"`
	ErrHwFail            uint32   `protobuf:"varint,11,opt,name=err_hw_fail,json=errHwFail,proto3" json:"err_hw_fail,omitempty"`
	ErrRefResolve        uint32   `protobuf:"varint,12,opt,name=err_ref_resolve,json=errRefResolve,proto3" json:"err_ref_resolve,omitempty"`
	ErrDbNotfound        uint32   `protobuf:"varint,13,opt,name=err_db_notfound,json=errDbNotfound,proto3" json:"err_db_notfound,omitempty"`
	ErrDbExists          uint32   `protobuf:"varint,14,opt,name=err_db_exists,json=errDbExists,proto3" json:"err_db_exists,omitempty"`
	ErrDbNomem           uint32   `protobuf:"varint,15,opt,name=err_db_nomem,json=errDbNomem,proto3" json:"err_db_nomem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DpaSysdbNpuTableResource) Reset()         { *m = DpaSysdbNpuTableResource{} }
func (m *DpaSysdbNpuTableResource) String() string { return proto.CompactTextString(m) }
func (*DpaSysdbNpuTableResource) ProtoMessage()    {}
func (*DpaSysdbNpuTableResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b72bbb1ac20a6d8, []int{1}
}

func (m *DpaSysdbNpuTableResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaSysdbNpuTableResource.Unmarshal(m, b)
}
func (m *DpaSysdbNpuTableResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaSysdbNpuTableResource.Marshal(b, m, deterministic)
}
func (m *DpaSysdbNpuTableResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaSysdbNpuTableResource.Merge(m, src)
}
func (m *DpaSysdbNpuTableResource) XXX_Size() int {
	return xxx_messageInfo_DpaSysdbNpuTableResource.Size(m)
}
func (m *DpaSysdbNpuTableResource) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaSysdbNpuTableResource.DiscardUnknown(m)
}

var xxx_messageInfo_DpaSysdbNpuTableResource proto.InternalMessageInfo

func (m *DpaSysdbNpuTableResource) GetNpuId() uint32 {
	if m != nil {
		return m.NpuId
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetNumObjs() uint32 {
	if m != nil {
		return m.NumObjs
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetCreateReq() uint32 {
	if m != nil {
		return m.CreateReq
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetCreateOk() uint32 {
	if m != nil {
		return m.CreateOk
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetDeleteReq() uint32 {
	if m != nil {
		return m.DeleteReq
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetDeleteOk() uint32 {
	if m != nil {
		return m.DeleteOk
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetUpdateReq() uint32 {
	if m != nil {
		return m.UpdateReq
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetUpdateOk() uint32 {
	if m != nil {
		return m.UpdateOk
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetEodReq() uint32 {
	if m != nil {
		return m.EodReq
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetEodOk() uint32 {
	if m != nil {
		return m.EodOk
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetErrHwFail() uint32 {
	if m != nil {
		return m.ErrHwFail
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetErrRefResolve() uint32 {
	if m != nil {
		return m.ErrRefResolve
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetErrDbNotfound() uint32 {
	if m != nil {
		return m.ErrDbNotfound
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetErrDbExists() uint32 {
	if m != nil {
		return m.ErrDbExists
	}
	return 0
}

func (m *DpaSysdbNpuTableResource) GetErrDbNomem() uint32 {
	if m != nil {
		return m.ErrDbNomem
	}
	return 0
}

type DpaSysdbTableResource struct {
	TableId              uint32                      `protobuf:"varint,50,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	Name                 []string                    `protobuf:"bytes,51,rep,name=name,proto3" json:"name,omitempty"`
	IsGlobal             bool                        `protobuf:"varint,52,opt,name=is_global,json=isGlobal,proto3" json:"is_global,omitempty"`
	NumNpus              uint32                      `protobuf:"varint,53,opt,name=num_npus,json=numNpus,proto3" json:"num_npus,omitempty"`
	TableSpecificList    []string                    `protobuf:"bytes,54,rep,name=table_specific_list,json=tableSpecificList,proto3" json:"table_specific_list,omitempty"`
	NpuTblr              []*DpaSysdbNpuTableResource `protobuf:"bytes,55,rep,name=npu_tblr,json=npuTblr,proto3" json:"npu_tblr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *DpaSysdbTableResource) Reset()         { *m = DpaSysdbTableResource{} }
func (m *DpaSysdbTableResource) String() string { return proto.CompactTextString(m) }
func (*DpaSysdbTableResource) ProtoMessage()    {}
func (*DpaSysdbTableResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b72bbb1ac20a6d8, []int{2}
}

func (m *DpaSysdbTableResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaSysdbTableResource.Unmarshal(m, b)
}
func (m *DpaSysdbTableResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaSysdbTableResource.Marshal(b, m, deterministic)
}
func (m *DpaSysdbTableResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaSysdbTableResource.Merge(m, src)
}
func (m *DpaSysdbTableResource) XXX_Size() int {
	return xxx_messageInfo_DpaSysdbTableResource.Size(m)
}
func (m *DpaSysdbTableResource) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaSysdbTableResource.DiscardUnknown(m)
}

var xxx_messageInfo_DpaSysdbTableResource proto.InternalMessageInfo

func (m *DpaSysdbTableResource) GetTableId() uint32 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *DpaSysdbTableResource) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *DpaSysdbTableResource) GetIsGlobal() bool {
	if m != nil {
		return m.IsGlobal
	}
	return false
}

func (m *DpaSysdbTableResource) GetNumNpus() uint32 {
	if m != nil {
		return m.NumNpus
	}
	return 0
}

func (m *DpaSysdbTableResource) GetTableSpecificList() []string {
	if m != nil {
		return m.TableSpecificList
	}
	return nil
}

func (m *DpaSysdbTableResource) GetNpuTblr() []*DpaSysdbNpuTableResource {
	if m != nil {
		return m.NpuTblr
	}
	return nil
}

func init() {
	proto.RegisterType((*DpaSysdbTableResource_KEYS)(nil), "cisco_ios_xr_fretta_bcm_dpa_resources_oper.dpa.resources.nodes.node.table_datas.table_data.dpa_sysdb_table_resource_KEYS")
	proto.RegisterType((*DpaSysdbNpuTableResource)(nil), "cisco_ios_xr_fretta_bcm_dpa_resources_oper.dpa.resources.nodes.node.table_datas.table_data.dpa_sysdb_npu_table_resource")
	proto.RegisterType((*DpaSysdbTableResource)(nil), "cisco_ios_xr_fretta_bcm_dpa_resources_oper.dpa.resources.nodes.node.table_datas.table_data.dpa_sysdb_table_resource")
}

func init() { proto.RegisterFile("dpa_sysdb_table_resource.proto", fileDescriptor_9b72bbb1ac20a6d8) }

var fileDescriptor_9b72bbb1ac20a6d8 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x55, 0x9a, 0x36, 0x99, 0xdc, 0x34, 0x54, 0x18, 0x55, 0x98, 0x47, 0xab, 0x28, 0x0b, 0x94,
	0xd5, 0x2c, 0x5a, 0x1e, 0x3f, 0x40, 0x81, 0x0a, 0xd4, 0x48, 0x53, 0x16, 0xc0, 0xc6, 0xf2, 0x8c,
	0xef, 0x80, 0x3b, 0x0f, 0xbb, 0xf6, 0x0c, 0x0d, 0xdf, 0xc1, 0x47, 0xf1, 0x2f, 0x7c, 0x05, 0xf2,
	0x23, 0x01, 0x21, 0x95, 0x1d, 0x9b, 0x91, 0x7d, 0x5e, 0xd7, 0x63, 0xdf, 0x0b, 0xc7, 0x42, 0x73,
	0x66, 0xbf, 0x59, 0x91, 0xb3, 0x8e, 0xe7, 0x35, 0x32, 0x83, 0x56, 0xf5, 0xa6, 0xc0, 0x54, 0x1b,
	0xd5, 0x29, 0xf2, 0xa9, 0x90, 0xb6, 0x50, 0x4c, 0x2a, 0xcb, 0xd6, 0x86, 0x95, 0x06, 0xbb, 0x8e,
	0xb3, 0xbc, 0x68, 0x98, 0xf3, 0x6d, 0xb4, 0x96, 0x29, 0x8d, 0x26, 0x15, 0x9a, 0xa7, 0x5b, 0x28,
	0x6d, 0x95, 0x88, 0xdf, 0x34, 0x24, 0x0b, 0xde, 0x71, 0xfb, 0xc7, 0x7a, 0xf1, 0x01, 0x8e, 0x6e,
	0xab, 0xce, 0xde, 0x9e, 0x7d, 0xbc, 0x24, 0x8f, 0x60, 0xe2, 0xfc, 0xac, 0xe5, 0x0d, 0xd2, 0xc1,
	0x7c, 0xb0, 0x9c, 0x64, 0x89, 0x03, 0x2e, 0x78, 0x83, 0xe4, 0x21, 0x24, 0x1b, 0x35, 0xdd, 0x99,
	0x0f, 0x96, 0xb3, 0x6c, 0xbb, 0x5f, 0xfc, 0x1c, 0xc2, 0xe3, 0xdf, 0xd1, 0xad, 0xee, 0xff, 0x8a,
	0x27, 0x87, 0x30, 0x72, 0xa8, 0x14, 0x3e, 0x76, 0x96, 0xed, 0xb5, 0xba, 0x3f, 0x17, 0xe4, 0x01,
	0x24, 0x6d, 0xdf, 0x30, 0x95, 0x5f, 0xd9, 0x98, 0x39, 0x6e, 0xfb, 0x66, 0x95, 0x5f, 0x59, 0x72,
	0x04, 0x50, 0x18, 0xe4, 0x9d, 0x0b, 0xb9, 0xa6, 0x43, 0x4f, 0x4e, 0x02, 0x92, 0xe1, 0xb5, 0x3b,
	0x6a, 0xa4, 0x55, 0x45, 0x77, 0xc3, 0x71, 0x02, 0xb0, 0xaa, 0x9c, 0x57, 0x60, 0x8d, 0xd1, 0xbb,
	0x17, 0xbc, 0x01, 0x89, 0xde, 0x48, 0xab, 0x8a, 0x8e, 0x82, 0x37, 0x00, 0xc1, 0xdb, 0x6b, 0xb1,
	0xa9, 0x3b, 0x0e, 0xde, 0x80, 0x44, 0x6f, 0xa4, 0x55, 0x45, 0x93, 0xe0, 0x0d, 0xc0, 0xaa, 0x22,
	0xf7, 0x61, 0x8c, 0x4a, 0x78, 0xe3, 0xc4, 0x53, 0x23, 0x54, 0xc2, 0xb9, 0x0e, 0xc1, 0xad, 0x9c,
	0x05, 0xc2, 0xef, 0xa3, 0x12, 0xab, 0x8a, 0x1c, 0xc3, 0x14, 0x8d, 0x61, 0x5f, 0x6e, 0x58, 0xc9,
	0x65, 0x4d, 0xa7, 0xa1, 0x18, 0x1a, 0xf3, 0xe6, 0xe6, 0x15, 0x97, 0x35, 0x79, 0x02, 0x07, 0x8e,
	0x37, 0x58, 0xfa, 0x9b, 0xac, 0xbf, 0x22, 0xdd, 0xf7, 0x9a, 0x19, 0x1a, 0x93, 0x61, 0x99, 0x05,
	0x70, 0xa3, 0x73, 0x57, 0xaf, 0xba, 0x52, 0xf5, 0xad, 0xa0, 0xb3, 0xad, 0xee, 0x65, 0x7e, 0x11,
	0x41, 0xb2, 0x80, 0x59, 0xd4, 0xe1, 0x5a, 0xda, 0xce, 0xd2, 0x3b, 0x5e, 0x35, 0xf5, 0xaa, 0x33,
	0x0f, 0x91, 0x39, 0xec, 0x6f, 0xb3, 0x1a, 0x6c, 0xe8, 0x81, 0x97, 0x40, 0x0c, 0x6a, 0xb0, 0x59,
	0xfc, 0xd8, 0x01, 0x7a, 0x5b, 0x1f, 0xb9, 0x17, 0x0d, 0x88, 0x14, 0xf4, 0x24, 0xbc, 0xa8, 0xdf,
	0x9f, 0x0b, 0x42, 0x60, 0xd7, 0x37, 0xd6, 0xe9, 0x7c, 0xb8, 0x9c, 0x64, 0x7e, 0xed, 0xae, 0x53,
	0x5a, 0xf6, 0xb9, 0x56, 0x39, 0xaf, 0xe9, 0xd3, 0xf9, 0x60, 0x99, 0x64, 0x89, 0xb4, 0xaf, 0xfd,
	0x7e, 0xd3, 0x1d, 0xad, 0xee, 0x2d, 0x7d, 0xb6, 0xed, 0x8e, 0x0b, 0xdd, 0x5b, 0x92, 0xc2, 0xbd,
	0x50, 0xc6, 0x6a, 0x2c, 0x64, 0x29, 0x0b, 0x56, 0x4b, 0xdb, 0xd1, 0xe7, 0x3e, 0xfa, 0xae, 0xa7,
	0x2e, 0x23, 0xf3, 0x4e, 0xda, 0x8e, 0x7c, 0x1f, 0x40, 0xe2, 0xdb, 0x32, 0xaf, 0x0d, 0x7d, 0x31,
	0x1f, 0x2e, 0xa7, 0x27, 0xeb, 0xf4, 0xff, 0x8d, 0x5a, 0xfa, 0xaf, 0x61, 0xc8, 0xc6, 0xad, 0xee,
	0xdf, 0xe7, 0xb5, 0xc9, 0x47, 0x7e, 0xe6, 0x4f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xf2,
	0x84, 0x29, 0x15, 0x04, 0x00, 0x00,
}
