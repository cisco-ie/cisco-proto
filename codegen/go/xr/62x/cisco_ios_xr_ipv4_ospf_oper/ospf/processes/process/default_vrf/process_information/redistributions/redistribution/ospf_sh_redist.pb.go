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
// source: ospf_sh_redist.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_process_information_redistributions_redistribution

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

type OspfShRedist_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	ProtocolName         string   `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	ProcessName_1        string   `protobuf:"bytes,3,opt,name=process_name_1,json=processName1,proto3" json:"process_name_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRedist_KEYS) Reset()         { *m = OspfShRedist_KEYS{} }
func (m *OspfShRedist_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShRedist_KEYS) ProtoMessage()    {}
func (*OspfShRedist_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_128e14098b86f6e9, []int{0}
}

func (m *OspfShRedist_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRedist_KEYS.Unmarshal(m, b)
}
func (m *OspfShRedist_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRedist_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShRedist_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRedist_KEYS.Merge(m, src)
}
func (m *OspfShRedist_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShRedist_KEYS.Size(m)
}
func (m *OspfShRedist_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRedist_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRedist_KEYS proto.InternalMessageInfo

func (m *OspfShRedist_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShRedist_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *OspfShRedist_KEYS) GetProcessName_1() string {
	if m != nil {
		return m.ProcessName_1
	}
	return ""
}

type OspfShRedistProto struct {
	ProtocolType         string   `protobuf:"bytes,1,opt,name=protocol_type,json=protocolType,proto3" json:"protocol_type,omitempty"`
	IsisInstanceId       string   `protobuf:"bytes,2,opt,name=isis_instance_id,json=isisInstanceId,proto3" json:"isis_instance_id,omitempty"`
	OspfProcessId        string   `protobuf:"bytes,3,opt,name=ospf_process_id,json=ospfProcessId,proto3" json:"ospf_process_id,omitempty"`
	BgpAsNumber          string   `protobuf:"bytes,4,opt,name=bgp_as_number,json=bgpAsNumber,proto3" json:"bgp_as_number,omitempty"`
	EigrpAsNumber        string   `protobuf:"bytes,5,opt,name=eigrp_as_number,json=eigrpAsNumber,proto3" json:"eigrp_as_number,omitempty"`
	ConnectedInstance    string   `protobuf:"bytes,6,opt,name=connected_instance,json=connectedInstance,proto3" json:"connected_instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRedistProto) Reset()         { *m = OspfShRedistProto{} }
func (m *OspfShRedistProto) String() string { return proto.CompactTextString(m) }
func (*OspfShRedistProto) ProtoMessage()    {}
func (*OspfShRedistProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_128e14098b86f6e9, []int{1}
}

func (m *OspfShRedistProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRedistProto.Unmarshal(m, b)
}
func (m *OspfShRedistProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRedistProto.Marshal(b, m, deterministic)
}
func (m *OspfShRedistProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRedistProto.Merge(m, src)
}
func (m *OspfShRedistProto) XXX_Size() int {
	return xxx_messageInfo_OspfShRedistProto.Size(m)
}
func (m *OspfShRedistProto) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRedistProto.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRedistProto proto.InternalMessageInfo

func (m *OspfShRedistProto) GetProtocolType() string {
	if m != nil {
		return m.ProtocolType
	}
	return ""
}

func (m *OspfShRedistProto) GetIsisInstanceId() string {
	if m != nil {
		return m.IsisInstanceId
	}
	return ""
}

func (m *OspfShRedistProto) GetOspfProcessId() string {
	if m != nil {
		return m.OspfProcessId
	}
	return ""
}

func (m *OspfShRedistProto) GetBgpAsNumber() string {
	if m != nil {
		return m.BgpAsNumber
	}
	return ""
}

func (m *OspfShRedistProto) GetEigrpAsNumber() string {
	if m != nil {
		return m.EigrpAsNumber
	}
	return ""
}

func (m *OspfShRedistProto) GetConnectedInstance() string {
	if m != nil {
		return m.ConnectedInstance
	}
	return ""
}

type OspfShRedist struct {
	RedistributionProtocol *OspfShRedistProto `protobuf:"bytes,50,opt,name=redistribution_protocol,json=redistributionProtocol,proto3" json:"redistribution_protocol,omitempty"`
	MetricFlag             bool               `protobuf:"varint,51,opt,name=metric_flag,json=metricFlag,proto3" json:"metric_flag,omitempty"`
	Metric                 uint32             `protobuf:"varint,52,opt,name=metric,proto3" json:"metric,omitempty"`
	Classless              bool               `protobuf:"varint,53,opt,name=classless,proto3" json:"classless,omitempty"`
	NssaOnly               bool               `protobuf:"varint,54,opt,name=nssa_only,json=nssaOnly,proto3" json:"nssa_only,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *OspfShRedist) Reset()         { *m = OspfShRedist{} }
func (m *OspfShRedist) String() string { return proto.CompactTextString(m) }
func (*OspfShRedist) ProtoMessage()    {}
func (*OspfShRedist) Descriptor() ([]byte, []int) {
	return fileDescriptor_128e14098b86f6e9, []int{2}
}

func (m *OspfShRedist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRedist.Unmarshal(m, b)
}
func (m *OspfShRedist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRedist.Marshal(b, m, deterministic)
}
func (m *OspfShRedist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRedist.Merge(m, src)
}
func (m *OspfShRedist) XXX_Size() int {
	return xxx_messageInfo_OspfShRedist.Size(m)
}
func (m *OspfShRedist) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRedist.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRedist proto.InternalMessageInfo

func (m *OspfShRedist) GetRedistributionProtocol() *OspfShRedistProto {
	if m != nil {
		return m.RedistributionProtocol
	}
	return nil
}

func (m *OspfShRedist) GetMetricFlag() bool {
	if m != nil {
		return m.MetricFlag
	}
	return false
}

func (m *OspfShRedist) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *OspfShRedist) GetClassless() bool {
	if m != nil {
		return m.Classless
	}
	return false
}

func (m *OspfShRedist) GetNssaOnly() bool {
	if m != nil {
		return m.NssaOnly
	}
	return false
}

func init() {
	proto.RegisterType((*OspfShRedist_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.redistributions.redistribution.ospf_sh_redist_KEYS")
	proto.RegisterType((*OspfShRedistProto)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.redistributions.redistribution.ospf_sh_redist_proto")
	proto.RegisterType((*OspfShRedist)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.redistributions.redistribution.ospf_sh_redist")
}

func init() { proto.RegisterFile("ospf_sh_redist.proto", fileDescriptor_128e14098b86f6e9) }

var fileDescriptor_128e14098b86f6e9 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x95, 0x02, 0xd5, 0x76, 0xba, 0x2d, 0x60, 0xd0, 0x62, 0x09, 0x24, 0x4a, 0x41, 0x28,
	0x17, 0x22, 0xed, 0x1f, 0xb8, 0x73, 0x00, 0xa9, 0x42, 0x5a, 0x56, 0x85, 0x0b, 0xa7, 0x91, 0xe3,
	0xb8, 0xc1, 0x52, 0x62, 0x07, 0x4f, 0xba, 0xa2, 0x67, 0x0e, 0xf0, 0xb5, 0xb8, 0xf1, 0xb1, 0x50,
	0xec, 0x24, 0x34, 0x15, 0xe7, 0xbd, 0x79, 0x7e, 0x7e, 0x79, 0xf3, 0x66, 0x1c, 0x78, 0x68, 0xa9,
	0xda, 0x20, 0x7d, 0x45, 0xa7, 0x32, 0x4d, 0x75, 0x52, 0x39, 0x5b, 0x5b, 0xf6, 0x4d, 0x6a, 0x92,
	0x16, 0xb5, 0x25, 0xfc, 0xee, 0x50, 0x57, 0xd7, 0x17, 0xe8, 0x75, 0xb6, 0x52, 0x2e, 0x69, 0x4e,
	0x8d, 0x4e, 0x2a, 0x22, 0x45, 0xdd, 0x29, 0xc9, 0xd4, 0x46, 0x6c, 0x8b, 0x1a, 0xaf, 0x5d, 0x7f,
	0x8b, 0xda, 0x6c, 0xac, 0x2b, 0x45, 0xad, 0xad, 0x49, 0x42, 0x03, 0xa7, 0xd3, 0x6d, 0x53, 0xd2,
	0x41, 0xbd, 0xfc, 0x11, 0xc1, 0x83, 0x61, 0x16, 0xfc, 0xf0, 0xee, 0xcb, 0x27, 0xf6, 0x0c, 0x8e,
	0x3b, 0x37, 0x23, 0x4a, 0xc5, 0xa3, 0x45, 0x14, 0x4f, 0xd6, 0xd3, 0x96, 0x5d, 0x8a, 0x52, 0xb1,
	0xe7, 0x30, 0xf3, 0xb1, 0xa5, 0x2d, 0x82, 0x66, 0xe4, 0x35, 0xc7, 0x1d, 0xf4, 0xa2, 0x17, 0x30,
	0xdf, 0xf7, 0xc1, 0x53, 0x7e, 0xab, 0x57, 0x75, 0x4e, 0xa7, 0xcb, 0x5f, 0xa3, 0xc3, 0x8d, 0x60,
	0xd8, 0xc8, 0x7e, 0x8f, 0x7a, 0x57, 0x75, 0x39, 0xfa, 0x1e, 0x9f, 0x77, 0x95, 0x62, 0x31, 0xdc,
	0xd3, 0xa4, 0x9b, 0xb1, 0xa9, 0x16, 0x46, 0x2a, 0xd4, 0x59, 0x9b, 0x65, 0xde, 0xf0, 0x55, 0x8b,
	0x57, 0x19, 0x7b, 0x09, 0x77, 0x7d, 0x9b, 0x7e, 0x51, 0x59, 0x1b, 0x67, 0xd6, 0xe0, 0xab, 0x40,
	0x57, 0x19, 0x5b, 0xc2, 0x2c, 0xcd, 0x2b, 0x14, 0x84, 0x66, 0x5b, 0xa6, 0xca, 0xf1, 0xdb, 0x61,
	0xfc, 0x34, 0xaf, 0xde, 0xd2, 0xa5, 0x47, 0x8d, 0x97, 0xd2, 0xb9, 0xdb, 0x57, 0xdd, 0x09, 0x5e,
	0x1e, 0xf7, 0xba, 0x57, 0xc0, 0xa4, 0x35, 0x46, 0xc9, 0x5a, 0x65, 0x7d, 0x44, 0x3e, 0xf6, 0xd2,
	0xfb, 0xfd, 0x4d, 0x17, 0x72, 0xf9, 0x67, 0x04, 0xf3, 0xe1, 0x2a, 0xd8, 0xef, 0x08, 0x1e, 0x0d,
	0x9f, 0x0d, 0xbb, 0xf9, 0xf9, 0xd9, 0x22, 0x8a, 0xa7, 0x67, 0x3f, 0xa3, 0xe4, 0xc6, 0x7f, 0x9d,
	0xe4, 0x7f, 0x0f, 0xb6, 0x3e, 0x19, 0x8a, 0xae, 0xda, 0x9c, 0xec, 0x29, 0x4c, 0x4b, 0x55, 0x3b,
	0x2d, 0x71, 0x53, 0x88, 0x9c, 0x9f, 0x2f, 0xa2, 0xf8, 0x68, 0x0d, 0x01, 0xbd, 0x2f, 0x44, 0xce,
	0x4e, 0x60, 0x1c, 0x2a, 0x7e, 0xb1, 0x88, 0xe2, 0xd9, 0xba, 0xad, 0xd8, 0x13, 0x98, 0xc8, 0x42,
	0x10, 0x15, 0x8a, 0x88, 0xbf, 0xf6, 0x9f, 0xfd, 0x03, 0xec, 0x31, 0x4c, 0x0c, 0x91, 0x40, 0x6b,
	0x8a, 0x1d, 0x7f, 0xe3, 0x6f, 0x8f, 0x1a, 0xf0, 0xd1, 0x14, 0xbb, 0x74, 0xec, 0x23, 0x9d, 0xff,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x45, 0x84, 0x53, 0x79, 0x6d, 0x03, 0x00, 0x00,
}