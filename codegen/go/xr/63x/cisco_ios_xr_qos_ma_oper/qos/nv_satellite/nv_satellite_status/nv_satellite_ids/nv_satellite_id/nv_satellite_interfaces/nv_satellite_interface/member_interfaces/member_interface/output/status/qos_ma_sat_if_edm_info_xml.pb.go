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
// source: qos_ma_sat_if_edm_info_xml.proto

package cisco_ios_xr_qos_ma_oper_qos_nv_satellite_nv_satellite_status_nv_satellite_ids_nv_satellite_id_nv_satellite_interfaces_nv_satellite_interface_member_interfaces_member_interface_output_status

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

type QosMaSatIfEdmInfoXml_KEYS struct {
	SatelliteId          uint32   `protobuf:"varint,1,opt,name=satellite_id,json=satelliteId,proto3" json:"satellite_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	InterfaceName_1      string   `protobuf:"bytes,3,opt,name=interface_name_1,json=interfaceName1,proto3" json:"interface_name_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QosMaSatIfEdmInfoXml_KEYS) Reset()         { *m = QosMaSatIfEdmInfoXml_KEYS{} }
func (m *QosMaSatIfEdmInfoXml_KEYS) String() string { return proto.CompactTextString(m) }
func (*QosMaSatIfEdmInfoXml_KEYS) ProtoMessage()    {}
func (*QosMaSatIfEdmInfoXml_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb8a0abb728c44e, []int{0}
}

func (m *QosMaSatIfEdmInfoXml_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosMaSatIfEdmInfoXml_KEYS.Unmarshal(m, b)
}
func (m *QosMaSatIfEdmInfoXml_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosMaSatIfEdmInfoXml_KEYS.Marshal(b, m, deterministic)
}
func (m *QosMaSatIfEdmInfoXml_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosMaSatIfEdmInfoXml_KEYS.Merge(m, src)
}
func (m *QosMaSatIfEdmInfoXml_KEYS) XXX_Size() int {
	return xxx_messageInfo_QosMaSatIfEdmInfoXml_KEYS.Size(m)
}
func (m *QosMaSatIfEdmInfoXml_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_QosMaSatIfEdmInfoXml_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_QosMaSatIfEdmInfoXml_KEYS proto.InternalMessageInfo

func (m *QosMaSatIfEdmInfoXml_KEYS) GetSatelliteId() uint32 {
	if m != nil {
		return m.SatelliteId
	}
	return 0
}

func (m *QosMaSatIfEdmInfoXml_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *QosMaSatIfEdmInfoXml_KEYS) GetInterfaceName_1() string {
	if m != nil {
		return m.InterfaceName_1
	}
	return ""
}

type QosMaSatIfEdmInfoXml struct {
	PolicyName           string   `protobuf:"bytes,50,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	Direction            string   `protobuf:"bytes,51,opt,name=direction,proto3" json:"direction,omitempty"`
	LastOperation        string   `protobuf:"bytes,52,opt,name=last_operation,json=lastOperation,proto3" json:"last_operation,omitempty"`
	Status               string   `protobuf:"bytes,53,opt,name=status,proto3" json:"status,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,54,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QosMaSatIfEdmInfoXml) Reset()         { *m = QosMaSatIfEdmInfoXml{} }
func (m *QosMaSatIfEdmInfoXml) String() string { return proto.CompactTextString(m) }
func (*QosMaSatIfEdmInfoXml) ProtoMessage()    {}
func (*QosMaSatIfEdmInfoXml) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb8a0abb728c44e, []int{1}
}

func (m *QosMaSatIfEdmInfoXml) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosMaSatIfEdmInfoXml.Unmarshal(m, b)
}
func (m *QosMaSatIfEdmInfoXml) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosMaSatIfEdmInfoXml.Marshal(b, m, deterministic)
}
func (m *QosMaSatIfEdmInfoXml) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosMaSatIfEdmInfoXml.Merge(m, src)
}
func (m *QosMaSatIfEdmInfoXml) XXX_Size() int {
	return xxx_messageInfo_QosMaSatIfEdmInfoXml.Size(m)
}
func (m *QosMaSatIfEdmInfoXml) XXX_DiscardUnknown() {
	xxx_messageInfo_QosMaSatIfEdmInfoXml.DiscardUnknown(m)
}

var xxx_messageInfo_QosMaSatIfEdmInfoXml proto.InternalMessageInfo

func (m *QosMaSatIfEdmInfoXml) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *QosMaSatIfEdmInfoXml) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *QosMaSatIfEdmInfoXml) GetLastOperation() string {
	if m != nil {
		return m.LastOperation
	}
	return ""
}

func (m *QosMaSatIfEdmInfoXml) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *QosMaSatIfEdmInfoXml) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*QosMaSatIfEdmInfoXml_KEYS)(nil), "cisco_ios_xr_qos_ma_oper.qos.nv_satellite.nv_satellite_status.nv_satellite_ids.nv_satellite_id.nv_satellite_interfaces.nv_satellite_interface.member_interfaces.member_interface.output.status.qos_ma_sat_if_edm_info_xml_KEYS")
	proto.RegisterType((*QosMaSatIfEdmInfoXml)(nil), "cisco_ios_xr_qos_ma_oper.qos.nv_satellite.nv_satellite_status.nv_satellite_ids.nv_satellite_id.nv_satellite_interfaces.nv_satellite_interface.member_interfaces.member_interface.output.status.qos_ma_sat_if_edm_info_xml")
}

func init() { proto.RegisterFile("qos_ma_sat_if_edm_info_xml.proto", fileDescriptor_eeb8a0abb728c44e) }

var fileDescriptor_eeb8a0abb728c44e = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0xc9, 0xf7, 0x41, 0xa1, 0xe9, 0x0f, 0x92, 0x85, 0x04, 0x11, 0x5a, 0x2b, 0x42, 0x57,
	0x81, 0x5a, 0xf5, 0x0e, 0x5c, 0x88, 0xa8, 0x50, 0x57, 0xae, 0x0e, 0xe9, 0xcc, 0xa9, 0x04, 0x26,
	0x73, 0xa6, 0x49, 0x2a, 0xf5, 0x42, 0xbc, 0x1d, 0xbd, 0x35, 0x69, 0x52, 0x6b, 0xa7, 0x32, 0xcb,
	0xf3, 0x9c, 0x27, 0x39, 0x6f, 0x7e, 0xf8, 0x70, 0x49, 0x1e, 0xac, 0x06, 0xaf, 0x03, 0x98, 0x05,
	0x60, 0x6e, 0xc1, 0x94, 0x0b, 0x82, 0xb5, 0x2d, 0x54, 0xe5, 0x28, 0x90, 0xf8, 0x64, 0x99, 0xf1,
	0x19, 0x81, 0x21, 0x0f, 0x6b, 0x07, 0x5b, 0x9f, 0x2a, 0x74, 0x6a, 0x49, 0x5e, 0x95, 0x6f, 0x9b,
	0xb5, 0x58, 0x14, 0x26, 0x60, 0xad, 0x00, 0x1f, 0x74, 0x58, 0xd5, 0x05, 0x30, 0xf9, 0x1f, 0x70,
	0x50, 0x97, 0x01, 0xdd, 0x42, 0x67, 0xe8, 0x1b, 0xb8, 0xb2, 0x68, 0xe7, 0xe8, 0xf6, 0xc5, 0x43,
	0xa2, 0x68, 0x15, 0xaa, 0x55, 0x50, 0x29, 0xc0, 0xe8, 0x83, 0xf1, 0x41, 0xf3, 0x29, 0xe1, 0xfe,
	0xf6, 0xe5, 0x59, 0x9c, 0xf1, 0xee, 0x7e, 0x26, 0xc9, 0x86, 0x6c, 0xdc, 0x9b, 0x75, 0x76, 0xec,
	0x2e, 0x17, 0x17, 0xbc, 0xbf, 0x9b, 0x00, 0xa5, 0xb6, 0x28, 0xff, 0x0d, 0xd9, 0xb8, 0x3d, 0xeb,
	0xed, 0xe8, 0xa3, 0xb6, 0x28, 0xc6, 0xfc, 0xa8, 0xae, 0xc1, 0x44, 0xfe, 0x8f, 0x62, 0xbf, 0x26,
	0x4e, 0x46, 0x5f, 0x8c, 0x9f, 0x34, 0xe7, 0x12, 0x03, 0xde, 0xa9, 0xa8, 0x30, 0xd9, 0x7b, 0x1a,
	0x76, 0x19, 0xf7, 0xe0, 0x09, 0xc5, 0x49, 0xa7, 0xbc, 0x9d, 0x1b, 0x87, 0x59, 0x30, 0x54, 0xca,
	0x69, 0x6c, 0xff, 0x82, 0x4d, 0xdc, 0x42, 0xfb, 0x10, 0x1f, 0x4a, 0x47, 0xe5, 0x2a, 0xc5, 0xdd,
	0xd0, 0xa7, 0x1f, 0x28, 0x8e, 0x79, 0x2b, 0x5d, 0x93, 0xbc, 0x8e, 0xed, 0x6d, 0x25, 0xce, 0x79,
	0x0f, 0x9d, 0x23, 0x07, 0x16, 0xbd, 0xd7, 0xaf, 0x28, 0x6f, 0x62, 0xbb, 0x1b, 0xe1, 0x43, 0x62,
	0xf3, 0x56, 0xfc, 0x21, 0xd3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x91, 0x9f, 0xd8, 0x45,
	0x02, 0x00, 0x00,
}
