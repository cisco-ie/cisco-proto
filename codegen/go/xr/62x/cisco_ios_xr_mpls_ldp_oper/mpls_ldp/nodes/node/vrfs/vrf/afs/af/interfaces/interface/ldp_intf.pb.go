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
// source: ldp_intf.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_vrfs_vrf_afs_af_interfaces_interface

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

type LdpIntf_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,3,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpIntf_KEYS) Reset()         { *m = LdpIntf_KEYS{} }
func (m *LdpIntf_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpIntf_KEYS) ProtoMessage()    {}
func (*LdpIntf_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_de2a8d7666321eba, []int{0}
}

func (m *LdpIntf_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIntf_KEYS.Unmarshal(m, b)
}
func (m *LdpIntf_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIntf_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpIntf_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIntf_KEYS.Merge(m, src)
}
func (m *LdpIntf_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpIntf_KEYS.Size(m)
}
func (m *LdpIntf_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIntf_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIntf_KEYS proto.InternalMessageInfo

func (m *LdpIntf_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpIntf_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *LdpIntf_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *LdpIntf_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type LdpVrfInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpVrfInfo) Reset()         { *m = LdpVrfInfo{} }
func (m *LdpVrfInfo) String() string { return proto.CompactTextString(m) }
func (*LdpVrfInfo) ProtoMessage()    {}
func (*LdpVrfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_de2a8d7666321eba, []int{1}
}

func (m *LdpVrfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpVrfInfo.Unmarshal(m, b)
}
func (m *LdpVrfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpVrfInfo.Marshal(b, m, deterministic)
}
func (m *LdpVrfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpVrfInfo.Merge(m, src)
}
func (m *LdpVrfInfo) XXX_Size() int {
	return xxx_messageInfo_LdpVrfInfo.Size(m)
}
func (m *LdpVrfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpVrfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpVrfInfo proto.InternalMessageInfo

func (m *LdpVrfInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LdpVrfInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LdpIntfTeMeshGrp struct {
	LdpTeMeshGroupAllCfgd bool     `protobuf:"varint,1,opt,name=ldp_te_mesh_group_all_cfgd,json=ldpTeMeshGroupAllCfgd,proto3" json:"ldp_te_mesh_group_all_cfgd,omitempty"`
	LdpMeshGroupEnabled   bool     `protobuf:"varint,2,opt,name=ldp_mesh_group_enabled,json=ldpMeshGroupEnabled,proto3" json:"ldp_mesh_group_enabled,omitempty"`
	TeMeshGroupId         uint32   `protobuf:"varint,3,opt,name=te_mesh_group_id,json=teMeshGroupId,proto3" json:"te_mesh_group_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *LdpIntfTeMeshGrp) Reset()         { *m = LdpIntfTeMeshGrp{} }
func (m *LdpIntfTeMeshGrp) String() string { return proto.CompactTextString(m) }
func (*LdpIntfTeMeshGrp) ProtoMessage()    {}
func (*LdpIntfTeMeshGrp) Descriptor() ([]byte, []int) {
	return fileDescriptor_de2a8d7666321eba, []int{2}
}

func (m *LdpIntfTeMeshGrp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIntfTeMeshGrp.Unmarshal(m, b)
}
func (m *LdpIntfTeMeshGrp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIntfTeMeshGrp.Marshal(b, m, deterministic)
}
func (m *LdpIntfTeMeshGrp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIntfTeMeshGrp.Merge(m, src)
}
func (m *LdpIntfTeMeshGrp) XXX_Size() int {
	return xxx_messageInfo_LdpIntfTeMeshGrp.Size(m)
}
func (m *LdpIntfTeMeshGrp) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIntfTeMeshGrp.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIntfTeMeshGrp proto.InternalMessageInfo

func (m *LdpIntfTeMeshGrp) GetLdpTeMeshGroupAllCfgd() bool {
	if m != nil {
		return m.LdpTeMeshGroupAllCfgd
	}
	return false
}

func (m *LdpIntfTeMeshGrp) GetLdpMeshGroupEnabled() bool {
	if m != nil {
		return m.LdpMeshGroupEnabled
	}
	return false
}

func (m *LdpIntfTeMeshGrp) GetTeMeshGroupId() uint32 {
	if m != nil {
		return m.TeMeshGroupId
	}
	return 0
}

type LdpIntfAutocfg struct {
	Tuple                string   `protobuf:"bytes,1,opt,name=tuple,proto3" json:"tuple,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpIntfAutocfg) Reset()         { *m = LdpIntfAutocfg{} }
func (m *LdpIntfAutocfg) String() string { return proto.CompactTextString(m) }
func (*LdpIntfAutocfg) ProtoMessage()    {}
func (*LdpIntfAutocfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_de2a8d7666321eba, []int{3}
}

func (m *LdpIntfAutocfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIntfAutocfg.Unmarshal(m, b)
}
func (m *LdpIntfAutocfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIntfAutocfg.Marshal(b, m, deterministic)
}
func (m *LdpIntfAutocfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIntfAutocfg.Merge(m, src)
}
func (m *LdpIntfAutocfg) XXX_Size() int {
	return xxx_messageInfo_LdpIntfAutocfg.Size(m)
}
func (m *LdpIntfAutocfg) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIntfAutocfg.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIntfAutocfg proto.InternalMessageInfo

func (m *LdpIntfAutocfg) GetTuple() string {
	if m != nil {
		return m.Tuple
	}
	return ""
}

type LdpIntf struct {
	Interface            string              `protobuf:"bytes,50,opt,name=interface,proto3" json:"interface,omitempty"`
	InterfaceNameXr      string              `protobuf:"bytes,51,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	Vrf                  *LdpVrfInfo         `protobuf:"bytes,52,opt,name=vrf,proto3" json:"vrf,omitempty"`
	LdpEnabled           bool                `protobuf:"varint,53,opt,name=ldp_enabled,json=ldpEnabled,proto3" json:"ldp_enabled,omitempty"`
	IsImStale            bool                `protobuf:"varint,54,opt,name=is_im_stale,json=isImStale,proto3" json:"is_im_stale,omitempty"`
	LdpConfigMode        bool                `protobuf:"varint,55,opt,name=ldp_config_mode,json=ldpConfigMode,proto3" json:"ldp_config_mode,omitempty"`
	LdpAutoconfigDisable bool                `protobuf:"varint,56,opt,name=ldp_autoconfig_disable,json=ldpAutoconfigDisable,proto3" json:"ldp_autoconfig_disable,omitempty"`
	TeMeshGrp            []*LdpIntfTeMeshGrp `protobuf:"bytes,57,rep,name=te_mesh_grp,json=teMeshGrp,proto3" json:"te_mesh_grp,omitempty"`
	AutoConfig           []*LdpIntfAutocfg   `protobuf:"bytes,58,rep,name=auto_config,json=autoConfig,proto3" json:"auto_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LdpIntf) Reset()         { *m = LdpIntf{} }
func (m *LdpIntf) String() string { return proto.CompactTextString(m) }
func (*LdpIntf) ProtoMessage()    {}
func (*LdpIntf) Descriptor() ([]byte, []int) {
	return fileDescriptor_de2a8d7666321eba, []int{4}
}

func (m *LdpIntf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIntf.Unmarshal(m, b)
}
func (m *LdpIntf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIntf.Marshal(b, m, deterministic)
}
func (m *LdpIntf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIntf.Merge(m, src)
}
func (m *LdpIntf) XXX_Size() int {
	return xxx_messageInfo_LdpIntf.Size(m)
}
func (m *LdpIntf) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIntf.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIntf proto.InternalMessageInfo

func (m *LdpIntf) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *LdpIntf) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *LdpIntf) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpIntf) GetLdpEnabled() bool {
	if m != nil {
		return m.LdpEnabled
	}
	return false
}

func (m *LdpIntf) GetIsImStale() bool {
	if m != nil {
		return m.IsImStale
	}
	return false
}

func (m *LdpIntf) GetLdpConfigMode() bool {
	if m != nil {
		return m.LdpConfigMode
	}
	return false
}

func (m *LdpIntf) GetLdpAutoconfigDisable() bool {
	if m != nil {
		return m.LdpAutoconfigDisable
	}
	return false
}

func (m *LdpIntf) GetTeMeshGrp() []*LdpIntfTeMeshGrp {
	if m != nil {
		return m.TeMeshGrp
	}
	return nil
}

func (m *LdpIntf) GetAutoConfig() []*LdpIntfAutocfg {
	if m != nil {
		return m.AutoConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpIntf_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.afs.af.interfaces.interface.ldp_intf_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.afs.af.interfaces.interface.ldp_vrf_info")
	proto.RegisterType((*LdpIntfTeMeshGrp)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.afs.af.interfaces.interface.ldp_intf_te_mesh_grp")
	proto.RegisterType((*LdpIntfAutocfg)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.afs.af.interfaces.interface.ldp_intf_autocfg")
	proto.RegisterType((*LdpIntf)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.afs.af.interfaces.interface.ldp_intf")
}

func init() { proto.RegisterFile("ldp_intf.proto", fileDescriptor_de2a8d7666321eba) }

var fileDescriptor_de2a8d7666321eba = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0xd6, 0xb2, 0xb5, 0x37, 0xb4, 0x1b, 0x66, 0x40, 0xf8, 0x10, 0x54, 0x95, 0x80, 0x8a,
	0x87, 0x3c, 0xb4, 0xe3, 0x63, 0xbc, 0x4d, 0x63, 0x42, 0x13, 0x1a, 0x0f, 0x29, 0x0f, 0xf0, 0x64,
	0x79, 0xb1, 0x9d, 0x59, 0x72, 0x62, 0xcb, 0x4e, 0xab, 0xfd, 0x02, 0x10, 0xbf, 0x86, 0x5f, 0x88,
	0x84, 0x6c, 0x37, 0x69, 0xcb, 0x7b, 0x5f, 0x22, 0xfb, 0x9c, 0x73, 0x7d, 0xae, 0x8f, 0xaf, 0x02,
	0x43, 0x49, 0x35, 0x16, 0x55, 0xcd, 0x53, 0x6d, 0x54, 0xad, 0xd0, 0x3c, 0x17, 0x36, 0x57, 0x58,
	0x28, 0x8b, 0x6f, 0x0d, 0x2e, 0xb5, 0xb4, 0xd8, 0x29, 0x94, 0x66, 0x26, 0x6d, 0x76, 0x69, 0xa5,
	0x28, 0xb3, 0xfe, 0x9b, 0x2e, 0x0d, 0xb7, 0xee, 0x93, 0x12, 0x6e, 0x53, 0xc2, 0x53, 0x51, 0xd5,
	0xcc, 0x70, 0x92, 0x33, 0xbb, 0x5e, 0x8e, 0x7f, 0x45, 0x30, 0x68, 0x7c, 0xf0, 0x97, 0x8b, 0x1f,
	0x73, 0xf4, 0x14, 0xfa, 0xae, 0x1c, 0x57, 0xa4, 0x64, 0x49, 0x34, 0x8a, 0x26, 0xfd, 0xac, 0xe7,
	0x80, 0xaf, 0xa4, 0x64, 0xe8, 0x31, 0xf4, 0x96, 0x86, 0x07, 0x6e, 0xcf, 0x73, 0x07, 0x4b, 0xc3,
	0x3d, 0xf5, 0x08, 0x0e, 0xc8, 0x8a, 0xe9, 0x78, 0x66, 0x9f, 0x04, 0xe2, 0x25, 0x0c, 0x5b, 0xbf,
	0xc0, 0x77, 0x3d, 0x3f, 0x68, 0x51, 0x27, 0x1b, 0x4f, 0xe1, 0xae, 0x6b, 0xc4, 0x1d, 0x2f, 0x2a,
	0xae, 0x10, 0x82, 0xee, 0x46, 0x0b, 0x7e, 0x8d, 0x86, 0xb0, 0x27, 0xa8, 0x37, 0x1e, 0x64, 0x7b,
	0x82, 0x8e, 0xff, 0x44, 0x70, 0xdc, 0x76, 0x5f, 0x33, 0x5c, 0x32, 0x7b, 0x83, 0x0b, 0xa3, 0xd1,
	0x29, 0x3c, 0x71, 0xf8, 0x1a, 0x52, 0x0b, 0x8d, 0x89, 0x94, 0x38, 0xe7, 0x05, 0xf5, 0x47, 0xf6,
	0xb2, 0x07, 0x92, 0xea, 0x6f, 0xec, 0x8a, 0xd9, 0x9b, 0xcf, 0x8e, 0x3e, 0x93, 0xf2, 0x9c, 0x17,
	0x14, 0xcd, 0xe0, 0xa1, 0x2b, 0xdd, 0xa8, 0x63, 0x15, 0xb9, 0x96, 0x2c, 0xf8, 0xf6, 0xb2, 0xfb,
	0x92, 0xea, 0xb6, 0xe8, 0x22, 0x50, 0xe8, 0x35, 0x1c, 0x6d, 0x7b, 0x09, 0xea, 0x53, 0x18, 0x64,
	0x83, 0x7a, 0x6d, 0x71, 0x49, 0xc7, 0x13, 0x38, 0x6a, 0x1b, 0x26, 0x8b, 0x5a, 0xe5, 0xbc, 0x40,
	0xc7, 0x70, 0xa7, 0x5e, 0x68, 0xd9, 0x5c, 0x35, 0x6c, 0xc6, 0x7f, 0xbb, 0xd0, 0x6b, 0xa4, 0xe8,
	0x19, 0xf4, 0xdb, 0xb4, 0x92, 0xa9, 0x97, 0xad, 0x01, 0xf4, 0x06, 0xee, 0x6d, 0x27, 0x8c, 0x6f,
	0x4d, 0x32, 0xf3, 0xaa, 0xc3, 0xad, 0x90, 0xbf, 0x1b, 0x64, 0xa1, 0xb3, 0x34, 0x3c, 0x39, 0x19,
	0x45, 0x93, 0x78, 0x4a, 0xd2, 0x1d, 0xcc, 0x54, 0xba, 0xf9, 0x8c, 0x99, 0x73, 0x43, 0x2f, 0x20,
	0x76, 0x60, 0x13, 0xe4, 0x5b, 0x1f, 0x24, 0x48, 0xda, 0xe6, 0xf7, 0x1c, 0x62, 0x61, 0xb1, 0x28,
	0xb1, 0xad, 0x89, 0x64, 0xc9, 0x3b, 0x2f, 0xe8, 0x0b, 0x7b, 0x59, 0xce, 0x1d, 0x80, 0x5e, 0xc1,
	0xa1, 0x3b, 0x20, 0x57, 0x15, 0x17, 0x05, 0x2e, 0x15, 0x65, 0xc9, 0x7b, 0xaf, 0x71, 0xc3, 0x7b,
	0xee, 0xd1, 0x2b, 0x45, 0x19, 0x3a, 0x09, 0x8f, 0xe7, 0x93, 0x0d, 0x5a, 0x2a, 0xac, 0xb3, 0x48,
	0x3e, 0x78, 0xb9, 0x9b, 0x96, 0xb3, 0x96, 0xfc, 0x14, 0x38, 0xf4, 0x3b, 0x82, 0x78, 0x63, 0x7a,
	0x92, 0xd3, 0x51, 0x67, 0x12, 0x4f, 0xc5, 0xce, 0xc2, 0xf9, 0x7f, 0x5c, 0xb3, 0x7e, 0x33, 0x24,
	0x1a, 0xfd, 0x8c, 0x20, 0x76, 0xed, 0xaf, 0xee, 0x9a, 0x7c, 0xf4, 0xbd, 0xb0, 0xdd, 0xf6, 0xb2,
	0x9a, 0xc4, 0x0c, 0xdc, 0x22, 0xc4, 0x79, 0xbd, 0xef, 0xff, 0x3a, 0xb3, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x1b, 0xe7, 0x1f, 0xa7, 0x87, 0x04, 0x00, 0x00,
}
