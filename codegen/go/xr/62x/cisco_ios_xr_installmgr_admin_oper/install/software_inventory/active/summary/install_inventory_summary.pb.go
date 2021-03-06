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
// source: install_inventory_summary.proto

package cisco_ios_xr_installmgr_admin_oper_install_software_inventory_active_summary

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

type InstallInventorySummary_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallInventorySummary_KEYS) Reset()         { *m = InstallInventorySummary_KEYS{} }
func (m *InstallInventorySummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*InstallInventorySummary_KEYS) ProtoMessage()    {}
func (*InstallInventorySummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1276f36e83723dc, []int{0}
}

func (m *InstallInventorySummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallInventorySummary_KEYS.Unmarshal(m, b)
}
func (m *InstallInventorySummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallInventorySummary_KEYS.Marshal(b, m, deterministic)
}
func (m *InstallInventorySummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallInventorySummary_KEYS.Merge(m, src)
}
func (m *InstallInventorySummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_InstallInventorySummary_KEYS.Size(m)
}
func (m *InstallInventorySummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallInventorySummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InstallInventorySummary_KEYS proto.InternalMessageInfo

type PkgGroup struct {
	DeviceName           string   `protobuf:"bytes,1,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PkgGroup) Reset()         { *m = PkgGroup{} }
func (m *PkgGroup) String() string { return proto.CompactTextString(m) }
func (*PkgGroup) ProtoMessage()    {}
func (*PkgGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1276f36e83723dc, []int{1}
}

func (m *PkgGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PkgGroup.Unmarshal(m, b)
}
func (m *PkgGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PkgGroup.Marshal(b, m, deterministic)
}
func (m *PkgGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PkgGroup.Merge(m, src)
}
func (m *PkgGroup) XXX_Size() int {
	return xxx_messageInfo_PkgGroup.Size(m)
}
func (m *PkgGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PkgGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PkgGroup proto.InternalMessageInfo

func (m *PkgGroup) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *PkgGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PkgInfo struct {
	Package              *PkgGroup `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	Version              string    `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	BuildInformation     string    `protobuf:"bytes,3,opt,name=build_information,json=buildInformation,proto3" json:"build_information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PkgInfo) Reset()         { *m = PkgInfo{} }
func (m *PkgInfo) String() string { return proto.CompactTextString(m) }
func (*PkgInfo) ProtoMessage()    {}
func (*PkgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1276f36e83723dc, []int{2}
}

func (m *PkgInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PkgInfo.Unmarshal(m, b)
}
func (m *PkgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PkgInfo.Marshal(b, m, deterministic)
}
func (m *PkgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PkgInfo.Merge(m, src)
}
func (m *PkgInfo) XXX_Size() int {
	return xxx_messageInfo_PkgInfo.Size(m)
}
func (m *PkgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PkgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PkgInfo proto.InternalMessageInfo

func (m *PkgInfo) GetPackage() *PkgGroup {
	if m != nil {
		return m.Package
	}
	return nil
}

func (m *PkgInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PkgInfo) GetBuildInformation() string {
	if m != nil {
		return m.BuildInformation
	}
	return ""
}

type DefaultLoadpath_ struct {
	RequestId              uint32     `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SecureDomainRouterName []string   `protobuf:"bytes,2,rep,name=secure_domain_router_name,json=secureDomainRouterName,proto3" json:"secure_domain_router_name,omitempty"`
	AdminMatch             bool       `protobuf:"varint,3,opt,name=admin_match,json=adminMatch,proto3" json:"admin_match,omitempty"`
	LoadPath               []*PkgInfo `protobuf:"bytes,4,rep,name=load_path,json=loadPath,proto3" json:"load_path,omitempty"`
	StandbyLoadPath        []*PkgInfo `protobuf:"bytes,5,rep,name=standby_load_path,json=standbyLoadPath,proto3" json:"standby_load_path,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}   `json:"-"`
	XXX_unrecognized       []byte     `json:"-"`
	XXX_sizecache          int32      `json:"-"`
}

func (m *DefaultLoadpath_) Reset()         { *m = DefaultLoadpath_{} }
func (m *DefaultLoadpath_) String() string { return proto.CompactTextString(m) }
func (*DefaultLoadpath_) ProtoMessage()    {}
func (*DefaultLoadpath_) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1276f36e83723dc, []int{3}
}

func (m *DefaultLoadpath_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultLoadpath_.Unmarshal(m, b)
}
func (m *DefaultLoadpath_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultLoadpath_.Marshal(b, m, deterministic)
}
func (m *DefaultLoadpath_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultLoadpath_.Merge(m, src)
}
func (m *DefaultLoadpath_) XXX_Size() int {
	return xxx_messageInfo_DefaultLoadpath_.Size(m)
}
func (m *DefaultLoadpath_) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultLoadpath_.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultLoadpath_ proto.InternalMessageInfo

func (m *DefaultLoadpath_) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *DefaultLoadpath_) GetSecureDomainRouterName() []string {
	if m != nil {
		return m.SecureDomainRouterName
	}
	return nil
}

func (m *DefaultLoadpath_) GetAdminMatch() bool {
	if m != nil {
		return m.AdminMatch
	}
	return false
}

func (m *DefaultLoadpath_) GetLoadPath() []*PkgInfo {
	if m != nil {
		return m.LoadPath
	}
	return nil
}

func (m *DefaultLoadpath_) GetStandbyLoadPath() []*PkgInfo {
	if m != nil {
		return m.StandbyLoadPath
	}
	return nil
}

type AdminLoadpath_ struct {
	RequestId            uint32     `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	LoadPath             []*PkgInfo `protobuf:"bytes,2,rep,name=load_path,json=loadPath,proto3" json:"load_path,omitempty"`
	StandbyLoadPath      []*PkgInfo `protobuf:"bytes,3,rep,name=standby_load_path,json=standbyLoadPath,proto3" json:"standby_load_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AdminLoadpath_) Reset()         { *m = AdminLoadpath_{} }
func (m *AdminLoadpath_) String() string { return proto.CompactTextString(m) }
func (*AdminLoadpath_) ProtoMessage()    {}
func (*AdminLoadpath_) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1276f36e83723dc, []int{4}
}

func (m *AdminLoadpath_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminLoadpath_.Unmarshal(m, b)
}
func (m *AdminLoadpath_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminLoadpath_.Marshal(b, m, deterministic)
}
func (m *AdminLoadpath_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminLoadpath_.Merge(m, src)
}
func (m *AdminLoadpath_) XXX_Size() int {
	return xxx_messageInfo_AdminLoadpath_.Size(m)
}
func (m *AdminLoadpath_) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminLoadpath_.DiscardUnknown(m)
}

var xxx_messageInfo_AdminLoadpath_ proto.InternalMessageInfo

func (m *AdminLoadpath_) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *AdminLoadpath_) GetLoadPath() []*PkgInfo {
	if m != nil {
		return m.LoadPath
	}
	return nil
}

func (m *AdminLoadpath_) GetStandbyLoadPath() []*PkgInfo {
	if m != nil {
		return m.StandbyLoadPath
	}
	return nil
}

type LrSpecificLoadpath struct {
	RequestId              uint32     `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SecureDomainRouterName string     `protobuf:"bytes,2,opt,name=secure_domain_router_name,json=secureDomainRouterName,proto3" json:"secure_domain_router_name,omitempty"`
	LoadPath               []*PkgInfo `protobuf:"bytes,3,rep,name=load_path,json=loadPath,proto3" json:"load_path,omitempty"`
	StandbyLoadPath        []*PkgInfo `protobuf:"bytes,4,rep,name=standby_load_path,json=standbyLoadPath,proto3" json:"standby_load_path,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}   `json:"-"`
	XXX_unrecognized       []byte     `json:"-"`
	XXX_sizecache          int32      `json:"-"`
}

func (m *LrSpecificLoadpath) Reset()         { *m = LrSpecificLoadpath{} }
func (m *LrSpecificLoadpath) String() string { return proto.CompactTextString(m) }
func (*LrSpecificLoadpath) ProtoMessage()    {}
func (*LrSpecificLoadpath) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1276f36e83723dc, []int{5}
}

func (m *LrSpecificLoadpath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LrSpecificLoadpath.Unmarshal(m, b)
}
func (m *LrSpecificLoadpath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LrSpecificLoadpath.Marshal(b, m, deterministic)
}
func (m *LrSpecificLoadpath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LrSpecificLoadpath.Merge(m, src)
}
func (m *LrSpecificLoadpath) XXX_Size() int {
	return xxx_messageInfo_LrSpecificLoadpath.Size(m)
}
func (m *LrSpecificLoadpath) XXX_DiscardUnknown() {
	xxx_messageInfo_LrSpecificLoadpath.DiscardUnknown(m)
}

var xxx_messageInfo_LrSpecificLoadpath proto.InternalMessageInfo

func (m *LrSpecificLoadpath) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *LrSpecificLoadpath) GetSecureDomainRouterName() string {
	if m != nil {
		return m.SecureDomainRouterName
	}
	return ""
}

func (m *LrSpecificLoadpath) GetLoadPath() []*PkgInfo {
	if m != nil {
		return m.LoadPath
	}
	return nil
}

func (m *LrSpecificLoadpath) GetStandbyLoadPath() []*PkgInfo {
	if m != nil {
		return m.StandbyLoadPath
	}
	return nil
}

type LocationSpecificLoadpath struct {
	RequestId              uint32     `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SecureDomainRouterName string     `protobuf:"bytes,2,opt,name=secure_domain_router_name,json=secureDomainRouterName,proto3" json:"secure_domain_router_name,omitempty"`
	NodeName               string     `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	LoadPath               []*PkgInfo `protobuf:"bytes,4,rep,name=load_path,json=loadPath,proto3" json:"load_path,omitempty"`
	StandbyLoadPath        []*PkgInfo `protobuf:"bytes,5,rep,name=standby_load_path,json=standbyLoadPath,proto3" json:"standby_load_path,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}   `json:"-"`
	XXX_unrecognized       []byte     `json:"-"`
	XXX_sizecache          int32      `json:"-"`
}

func (m *LocationSpecificLoadpath) Reset()         { *m = LocationSpecificLoadpath{} }
func (m *LocationSpecificLoadpath) String() string { return proto.CompactTextString(m) }
func (*LocationSpecificLoadpath) ProtoMessage()    {}
func (*LocationSpecificLoadpath) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1276f36e83723dc, []int{6}
}

func (m *LocationSpecificLoadpath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationSpecificLoadpath.Unmarshal(m, b)
}
func (m *LocationSpecificLoadpath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationSpecificLoadpath.Marshal(b, m, deterministic)
}
func (m *LocationSpecificLoadpath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationSpecificLoadpath.Merge(m, src)
}
func (m *LocationSpecificLoadpath) XXX_Size() int {
	return xxx_messageInfo_LocationSpecificLoadpath.Size(m)
}
func (m *LocationSpecificLoadpath) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationSpecificLoadpath.DiscardUnknown(m)
}

var xxx_messageInfo_LocationSpecificLoadpath proto.InternalMessageInfo

func (m *LocationSpecificLoadpath) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *LocationSpecificLoadpath) GetSecureDomainRouterName() string {
	if m != nil {
		return m.SecureDomainRouterName
	}
	return ""
}

func (m *LocationSpecificLoadpath) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LocationSpecificLoadpath) GetLoadPath() []*PkgInfo {
	if m != nil {
		return m.LoadPath
	}
	return nil
}

func (m *LocationSpecificLoadpath) GetStandbyLoadPath() []*PkgInfo {
	if m != nil {
		return m.StandbyLoadPath
	}
	return nil
}

type InstallInventorySummary struct {
	DefaultLoadPath      *DefaultLoadpath_           `protobuf:"bytes,50,opt,name=default_load_path,json=defaultLoadPath,proto3" json:"default_load_path,omitempty"`
	AdminLoadPath        *AdminLoadpath_             `protobuf:"bytes,51,opt,name=admin_load_path,json=adminLoadPath,proto3" json:"admin_load_path,omitempty"`
	SdrLoadPath          []*LrSpecificLoadpath       `protobuf:"bytes,52,rep,name=sdr_load_path,json=sdrLoadPath,proto3" json:"sdr_load_path,omitempty"`
	LocationLoadPath     []*LocationSpecificLoadpath `protobuf:"bytes,53,rep,name=location_load_path,json=locationLoadPath,proto3" json:"location_load_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *InstallInventorySummary) Reset()         { *m = InstallInventorySummary{} }
func (m *InstallInventorySummary) String() string { return proto.CompactTextString(m) }
func (*InstallInventorySummary) ProtoMessage()    {}
func (*InstallInventorySummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1276f36e83723dc, []int{7}
}

func (m *InstallInventorySummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallInventorySummary.Unmarshal(m, b)
}
func (m *InstallInventorySummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallInventorySummary.Marshal(b, m, deterministic)
}
func (m *InstallInventorySummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallInventorySummary.Merge(m, src)
}
func (m *InstallInventorySummary) XXX_Size() int {
	return xxx_messageInfo_InstallInventorySummary.Size(m)
}
func (m *InstallInventorySummary) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallInventorySummary.DiscardUnknown(m)
}

var xxx_messageInfo_InstallInventorySummary proto.InternalMessageInfo

func (m *InstallInventorySummary) GetDefaultLoadPath() *DefaultLoadpath_ {
	if m != nil {
		return m.DefaultLoadPath
	}
	return nil
}

func (m *InstallInventorySummary) GetAdminLoadPath() *AdminLoadpath_ {
	if m != nil {
		return m.AdminLoadPath
	}
	return nil
}

func (m *InstallInventorySummary) GetSdrLoadPath() []*LrSpecificLoadpath {
	if m != nil {
		return m.SdrLoadPath
	}
	return nil
}

func (m *InstallInventorySummary) GetLocationLoadPath() []*LocationSpecificLoadpath {
	if m != nil {
		return m.LocationLoadPath
	}
	return nil
}

func init() {
	proto.RegisterType((*InstallInventorySummary_KEYS)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.active.summary.install_inventory_summary_KEYS")
	proto.RegisterType((*PkgGroup)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.active.summary.pkg_group")
	proto.RegisterType((*PkgInfo)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.active.summary.pkg_info")
	proto.RegisterType((*DefaultLoadpath_)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.active.summary.default_loadpath_")
	proto.RegisterType((*AdminLoadpath_)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.active.summary.admin_loadpath_")
	proto.RegisterType((*LrSpecificLoadpath)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.active.summary.lr_specific_loadpath")
	proto.RegisterType((*LocationSpecificLoadpath)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.active.summary.location_specific_loadpath")
	proto.RegisterType((*InstallInventorySummary)(nil), "cisco_ios_xr_installmgr_admin_oper.install.software_inventory.active.summary.install_inventory_summary")
}

func init() { proto.RegisterFile("install_inventory_summary.proto", fileDescriptor_d1276f36e83723dc) }

var fileDescriptor_d1276f36e83723dc = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x96, 0xc1, 0x8a, 0x13, 0x31,
	0x18, 0xc7, 0x99, 0x4e, 0x75, 0x3b, 0x5f, 0x29, 0xdd, 0x06, 0x91, 0xae, 0xa2, 0x2d, 0x3d, 0x15,
	0x84, 0x39, 0x74, 0xf5, 0xe0, 0xcd, 0x83, 0x1e, 0x16, 0x57, 0x91, 0x08, 0x8a, 0x07, 0x09, 0xe9,
	0x24, 0x6d, 0xc3, 0xce, 0x4c, 0x66, 0x93, 0x4c, 0xb5, 0x57, 0x41, 0x11, 0xc4, 0xab, 0x0f, 0xe1,
	0x73, 0xf8, 0x06, 0xbe, 0x85, 0x4f, 0x21, 0x93, 0x49, 0xdb, 0xad, 0x6b, 0x45, 0x70, 0xec, 0x1e,
	0xf6, 0xd6, 0xfe, 0xbf, 0xf4, 0xfb, 0x7e, 0xf3, 0xcf, 0xe4, 0x9f, 0x42, 0x4f, 0xa4, 0xda, 0xd0,
	0x38, 0x26, 0x22, 0x9d, 0xf3, 0xd4, 0x48, 0xb5, 0x20, 0x3a, 0x4f, 0x12, 0xaa, 0x16, 0x61, 0xa6,
	0xa4, 0x91, 0xe8, 0x38, 0x12, 0x3a, 0x92, 0x44, 0x48, 0x4d, 0xde, 0x2a, 0xe2, 0x56, 0x27, 0x53,
	0x45, 0x28, 0x4b, 0x44, 0x4a, 0x64, 0xc6, 0x55, 0xe8, 0xd4, 0x50, 0xcb, 0x89, 0x79, 0x43, 0x15,
	0x5f, 0x37, 0x0b, 0x69, 0x64, 0xc4, 0x9c, 0x87, 0xae, 0xe7, 0xa0, 0x0f, 0xb7, 0xb7, 0x0e, 0x24,
	0x8f, 0x1f, 0xbd, 0x7a, 0x3e, 0x78, 0x00, 0x41, 0x76, 0x32, 0x25, 0x53, 0x25, 0xf3, 0x0c, 0xf5,
	0xa0, 0xc9, 0xf8, 0x5c, 0x44, 0x9c, 0xa4, 0x34, 0xe1, 0x5d, 0xaf, 0xef, 0x0d, 0x03, 0x0c, 0xa5,
	0xf4, 0x94, 0x26, 0x1c, 0x21, 0xa8, 0xdb, 0x4a, 0xcd, 0x56, 0xec, 0xe7, 0xc1, 0x37, 0x0f, 0x1a,
	0x45, 0x0b, 0x91, 0x4e, 0x24, 0x3a, 0x85, 0xbd, 0x8c, 0x46, 0x27, 0x74, 0x5a, 0xfe, 0xba, 0x39,
	0x7a, 0x19, 0x56, 0xf9, 0x40, 0xe1, 0x8a, 0x15, 0x2f, 0xe7, 0xa0, 0x2e, 0xec, 0xcd, 0xb9, 0xd2,
	0x42, 0xa6, 0x0e, 0x6b, 0xf9, 0x15, 0xdd, 0x81, 0xce, 0x38, 0x17, 0x31, 0xb3, 0x68, 0x2a, 0xa1,
	0xa6, 0x58, 0xe3, 0xdb, 0x35, 0xfb, 0xb6, 0x70, 0xb4, 0xd6, 0x07, 0x1f, 0x7d, 0xe8, 0x30, 0x3e,
	0xa1, 0x79, 0x6c, 0x48, 0x2c, 0x29, 0xcb, 0xa8, 0x99, 0x11, 0x74, 0x0b, 0x40, 0xf1, 0xd3, 0x9c,
	0x6b, 0x43, 0x04, 0xb3, 0x8f, 0xd4, 0xc2, 0x81, 0x53, 0x8e, 0x18, 0xba, 0x0f, 0x07, 0x9a, 0x47,
	0xb9, 0xe2, 0x84, 0xc9, 0x84, 0x8a, 0x94, 0x28, 0x99, 0x1b, 0xae, 0x88, 0x33, 0xc9, 0x1f, 0x06,
	0xf8, 0x7a, 0xb9, 0xe0, 0xa1, 0xad, 0x63, 0x5b, 0xb6, 0x56, 0xf6, 0xa0, 0x59, 0x3a, 0x90, 0x50,
	0x13, 0xcd, 0x2c, 0x56, 0x03, 0x83, 0x95, 0x9e, 0x14, 0x0a, 0xd2, 0x10, 0x14, 0x1c, 0xa4, 0x00,
	0xe9, 0xd6, 0xfb, 0xfe, 0xb0, 0x39, 0x7a, 0x51, 0xbd, 0x99, 0x85, 0x35, 0xb8, 0x51, 0x0c, 0x7a,
	0x46, 0xcd, 0x0c, 0xbd, 0xf3, 0xa0, 0xa3, 0x0d, 0x4d, 0xd9, 0x78, 0x41, 0xd6, 0xd3, 0xaf, 0xfc,
	0xd7, 0xe9, 0x6d, 0x37, 0xf0, 0xd8, 0x41, 0x0c, 0xbe, 0xd6, 0xa0, 0x5d, 0xb6, 0xfc, 0xeb, 0x8d,
	0xd8, 0x30, 0xab, 0x76, 0xa1, 0x66, 0xf9, 0xbb, 0x35, 0xeb, 0x47, 0x0d, 0xae, 0xc5, 0x8a, 0xe8,
	0x8c, 0x47, 0x62, 0x22, 0xa2, 0x95, 0x65, 0xff, 0xf8, 0xea, 0x7a, 0x7f, 0x78, 0x75, 0x37, 0xcc,
	0xf6, 0x2f, 0xd4, 0xec, 0xfa, 0x6e, 0xcd, 0xfe, 0xec, 0xc3, 0x8d, 0x58, 0x46, 0x36, 0x31, 0x76,
	0x6a, 0xf9, 0x4d, 0x08, 0x52, 0xc9, 0x5c, 0x2e, 0x97, 0x11, 0xd6, 0x28, 0x84, 0xf3, 0xfb, 0x71,
	0xa9, 0x92, 0xe2, 0x7b, 0x1d, 0x0e, 0xb6, 0x5e, 0x70, 0xe8, 0x93, 0xb7, 0x19, 0xe9, 0x25, 0xe2,
	0xc8, 0xde, 0x4b, 0xa4, 0x5a, 0xc4, 0x73, 0x37, 0x07, 0x6e, 0x3b, 0x69, 0xc9, 0x8a, 0xde, 0x7b,
	0x67, 0x53, 0xad, 0x64, 0x39, 0xb4, 0x2c, 0xaf, 0xab, 0x65, 0xf9, 0x25, 0x3a, 0x71, 0xcb, 0x0a,
	0x2b, 0x8e, 0x0f, 0x1e, 0xb4, 0x34, 0x53, 0x67, 0x28, 0xee, 0xda, 0x4d, 0x1b, 0x57, 0x4b, 0xf1,
	0xbb, 0x4c, 0xc2, 0x4d, 0xcd, 0xd4, 0x0a, 0xe4, 0x8b, 0x07, 0x68, 0x75, 0x98, 0xd6, 0x34, 0xf7,
	0x2c, 0xcd, 0xac, 0x62, 0x9a, 0xad, 0x87, 0x16, 0xef, 0x2f, 0x6b, 0x4b, 0xb0, 0xf1, 0x55, 0xfb,
	0x57, 0xec, 0xf0, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x6c, 0xc7, 0x20, 0xad, 0x09, 0x00,
	0x00,
}
