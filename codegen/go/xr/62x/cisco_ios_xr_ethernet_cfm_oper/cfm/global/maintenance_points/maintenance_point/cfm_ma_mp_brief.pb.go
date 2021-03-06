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
// source: cfm_ma_mp_brief.proto

package cisco_ios_xr_ethernet_cfm_oper_cfm_global_maintenance_points_maintenance_point

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

type CfmMaMpBrief_KEYS struct {
	Domain               string   `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Interface            string   `protobuf:"bytes,3,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CfmMaMpBrief_KEYS) Reset()         { *m = CfmMaMpBrief_KEYS{} }
func (m *CfmMaMpBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*CfmMaMpBrief_KEYS) ProtoMessage()    {}
func (*CfmMaMpBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e018ca12480c8ac, []int{0}
}

func (m *CfmMaMpBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CfmMaMpBrief_KEYS.Unmarshal(m, b)
}
func (m *CfmMaMpBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CfmMaMpBrief_KEYS.Marshal(b, m, deterministic)
}
func (m *CfmMaMpBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CfmMaMpBrief_KEYS.Merge(m, src)
}
func (m *CfmMaMpBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_CfmMaMpBrief_KEYS.Size(m)
}
func (m *CfmMaMpBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_CfmMaMpBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_CfmMaMpBrief_KEYS proto.InternalMessageInfo

func (m *CfmMaMpBrief_KEYS) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *CfmMaMpBrief_KEYS) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *CfmMaMpBrief_KEYS) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

type CfmMaMp struct {
	DomainName           string   `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	ServiceName          string   `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Interface            string   `protobuf:"bytes,4,opt,name=interface,proto3" json:"interface,omitempty"`
	MaintenancePointType string   `protobuf:"bytes,5,opt,name=maintenance_point_type,json=maintenancePointType,proto3" json:"maintenance_point_type,omitempty"`
	MepId                uint32   `protobuf:"varint,6,opt,name=mep_id,json=mepId,proto3" json:"mep_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CfmMaMp) Reset()         { *m = CfmMaMp{} }
func (m *CfmMaMp) String() string { return proto.CompactTextString(m) }
func (*CfmMaMp) ProtoMessage()    {}
func (*CfmMaMp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e018ca12480c8ac, []int{1}
}

func (m *CfmMaMp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CfmMaMp.Unmarshal(m, b)
}
func (m *CfmMaMp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CfmMaMp.Marshal(b, m, deterministic)
}
func (m *CfmMaMp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CfmMaMp.Merge(m, src)
}
func (m *CfmMaMp) XXX_Size() int {
	return xxx_messageInfo_CfmMaMp.Size(m)
}
func (m *CfmMaMp) XXX_DiscardUnknown() {
	xxx_messageInfo_CfmMaMp.DiscardUnknown(m)
}

var xxx_messageInfo_CfmMaMp proto.InternalMessageInfo

func (m *CfmMaMp) GetDomainName() string {
	if m != nil {
		return m.DomainName
	}
	return ""
}

func (m *CfmMaMp) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *CfmMaMp) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *CfmMaMp) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *CfmMaMp) GetMaintenancePointType() string {
	if m != nil {
		return m.MaintenancePointType
	}
	return ""
}

func (m *CfmMaMp) GetMepId() uint32 {
	if m != nil {
		return m.MepId
	}
	return 0
}

type CfmMaMpBrief struct {
	MaintenancePoint     *CfmMaMp `protobuf:"bytes,50,opt,name=maintenance_point,json=maintenancePoint,proto3" json:"maintenance_point,omitempty"`
	MepHasError          bool     `protobuf:"varint,51,opt,name=mep_has_error,json=mepHasError,proto3" json:"mep_has_error,omitempty"`
	MacAddress           string   `protobuf:"bytes,52,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CfmMaMpBrief) Reset()         { *m = CfmMaMpBrief{} }
func (m *CfmMaMpBrief) String() string { return proto.CompactTextString(m) }
func (*CfmMaMpBrief) ProtoMessage()    {}
func (*CfmMaMpBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e018ca12480c8ac, []int{2}
}

func (m *CfmMaMpBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CfmMaMpBrief.Unmarshal(m, b)
}
func (m *CfmMaMpBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CfmMaMpBrief.Marshal(b, m, deterministic)
}
func (m *CfmMaMpBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CfmMaMpBrief.Merge(m, src)
}
func (m *CfmMaMpBrief) XXX_Size() int {
	return xxx_messageInfo_CfmMaMpBrief.Size(m)
}
func (m *CfmMaMpBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_CfmMaMpBrief.DiscardUnknown(m)
}

var xxx_messageInfo_CfmMaMpBrief proto.InternalMessageInfo

func (m *CfmMaMpBrief) GetMaintenancePoint() *CfmMaMp {
	if m != nil {
		return m.MaintenancePoint
	}
	return nil
}

func (m *CfmMaMpBrief) GetMepHasError() bool {
	if m != nil {
		return m.MepHasError
	}
	return false
}

func (m *CfmMaMpBrief) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*CfmMaMpBrief_KEYS)(nil), "cisco_ios_xr_ethernet_cfm_oper.cfm.global.maintenance_points.maintenance_point.cfm_ma_mp_brief_KEYS")
	proto.RegisterType((*CfmMaMp)(nil), "cisco_ios_xr_ethernet_cfm_oper.cfm.global.maintenance_points.maintenance_point.cfm_ma_mp")
	proto.RegisterType((*CfmMaMpBrief)(nil), "cisco_ios_xr_ethernet_cfm_oper.cfm.global.maintenance_points.maintenance_point.cfm_ma_mp_brief")
}

func init() { proto.RegisterFile("cfm_ma_mp_brief.proto", fileDescriptor_8e018ca12480c8ac) }

var fileDescriptor_8e018ca12480c8ac = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4e, 0xeb, 0x30,
	0x10, 0xc6, 0xe5, 0xf7, 0x5e, 0xf3, 0xe8, 0x84, 0x0a, 0xb0, 0xda, 0xca, 0x0b, 0x24, 0x42, 0x56,
	0x59, 0x65, 0xd1, 0xf6, 0x02, 0x2c, 0x2a, 0x81, 0x90, 0x2a, 0x14, 0xd8, 0x74, 0x65, 0xb9, 0xce,
	0x84, 0x46, 0x8a, 0x63, 0xcb, 0x8e, 0x2a, 0x7a, 0x01, 0xce, 0xc7, 0x0d, 0xb8, 0x0a, 0xca, 0x1f,
	0x0a, 0xa4, 0x5b, 0x96, 0xf3, 0x7d, 0xe3, 0xf9, 0x7d, 0x9a, 0x31, 0x4c, 0x64, 0xa6, 0xb8, 0x12,
	0x5c, 0x19, 0xbe, 0xb1, 0x39, 0x66, 0xb1, 0xb1, 0xba, 0xd2, 0x74, 0x25, 0x73, 0x27, 0x35, 0xcf,
	0xb5, 0xe3, 0x2f, 0x96, 0x63, 0xb5, 0x45, 0x5b, 0x62, 0xc5, 0xeb, 0x66, 0x6d, 0xd0, 0xc6, 0x32,
	0x53, 0xf1, 0x73, 0xa1, 0x37, 0xa2, 0x88, 0x95, 0xc8, 0xcb, 0x0a, 0x4b, 0x51, 0x4a, 0xe4, 0x46,
	0xe7, 0x65, 0xe5, 0x8e, 0xa5, 0x30, 0x83, 0x71, 0x0f, 0xc4, 0xef, 0x97, 0xeb, 0x47, 0x3a, 0x05,
	0x2f, 0xd5, 0x75, 0x3b, 0x23, 0x01, 0x89, 0x86, 0x49, 0x57, 0x51, 0x06, 0xff, 0x1d, 0xda, 0x5d,
	0x2e, 0x91, 0xfd, 0x69, 0x8c, 0xcf, 0x92, 0x5e, 0xc2, 0xb0, 0x1e, 0x6e, 0x33, 0x21, 0x91, 0xfd,
	0x6d, 0xbc, 0x2f, 0x21, 0x7c, 0x23, 0x30, 0x3c, 0x80, 0xe8, 0x15, 0xf8, 0xed, 0x3c, 0x5e, 0x0a,
	0x85, 0x1d, 0x02, 0x5a, 0x69, 0x25, 0x14, 0xd2, 0x31, 0x0c, 0x0a, 0xdc, 0x61, 0xd1, 0x41, 0xda,
	0x82, 0x5e, 0xc3, 0x69, 0x47, 0x6b, 0xdf, 0xb5, 0x14, 0xbf, 0xd3, 0x9a, 0x87, 0x3f, 0x52, 0xfc,
	0xeb, 0xa5, 0xa0, 0x0b, 0x98, 0x1e, 0xad, 0x80, 0x57, 0x7b, 0x83, 0x6c, 0xd0, 0xb4, 0x8e, 0xbf,
	0xb9, 0x0f, 0xb5, 0xf9, 0xb4, 0x37, 0x48, 0x27, 0xe0, 0x29, 0x34, 0x3c, 0x4f, 0x99, 0x17, 0x90,
	0x68, 0x94, 0x0c, 0x14, 0x9a, 0xbb, 0x34, 0x7c, 0x27, 0x70, 0xd6, 0xdb, 0x1d, 0x7d, 0x25, 0x70,
	0x71, 0x44, 0x60, 0xb3, 0x80, 0x44, 0xfe, 0x6c, 0x1d, 0xff, 0xee, 0xed, 0xe2, 0x03, 0x3c, 0x39,
	0xef, 0xe7, 0xa6, 0x21, 0x8c, 0xea, 0xcc, 0x5b, 0xe1, 0x38, 0x5a, 0xab, 0x2d, 0x9b, 0x07, 0x24,
	0x3a, 0x49, 0x7c, 0x85, 0xe6, 0x56, 0xb8, 0x65, 0x2d, 0xd5, 0x57, 0x50, 0x42, 0x72, 0x91, 0xa6,
	0x16, 0x9d, 0x63, 0x8b, 0xf6, 0x0a, 0x4a, 0xc8, 0x9b, 0x56, 0xd9, 0x78, 0xcd, 0x9f, 0x9b, 0x7f,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x73, 0x8b, 0xbe, 0x8c, 0x02, 0x00, 0x00,
}
