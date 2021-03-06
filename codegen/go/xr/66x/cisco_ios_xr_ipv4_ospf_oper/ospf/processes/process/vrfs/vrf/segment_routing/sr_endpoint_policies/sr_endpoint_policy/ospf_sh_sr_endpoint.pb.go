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
// source: ospf_sh_sr_endpoint.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_segment_routing_sr_endpoint_policies_sr_endpoint_policy

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

type OspfShSrEndpoint_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	EndpointIp           string   `protobuf:"bytes,3,opt,name=endpoint_ip,json=endpointIp,proto3" json:"endpoint_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShSrEndpoint_KEYS) Reset()         { *m = OspfShSrEndpoint_KEYS{} }
func (m *OspfShSrEndpoint_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShSrEndpoint_KEYS) ProtoMessage()    {}
func (*OspfShSrEndpoint_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f586c47eaf48c07a, []int{0}
}

func (m *OspfShSrEndpoint_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShSrEndpoint_KEYS.Unmarshal(m, b)
}
func (m *OspfShSrEndpoint_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShSrEndpoint_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShSrEndpoint_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShSrEndpoint_KEYS.Merge(m, src)
}
func (m *OspfShSrEndpoint_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShSrEndpoint_KEYS.Size(m)
}
func (m *OspfShSrEndpoint_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShSrEndpoint_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShSrEndpoint_KEYS proto.InternalMessageInfo

func (m *OspfShSrEndpoint_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShSrEndpoint_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShSrEndpoint_KEYS) GetEndpointIp() string {
	if m != nil {
		return m.EndpointIp
	}
	return ""
}

type OspfShSrPolicy struct {
	PolicyName               string   `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyIfHandle           string   `protobuf:"bytes,2,opt,name=policy_if_handle,json=policyIfHandle,proto3" json:"policy_if_handle,omitempty"`
	PolicyMetric             int32    `protobuf:"zigzag32,3,opt,name=policy_metric,json=policyMetric,proto3" json:"policy_metric,omitempty"`
	PolicyMetricMode         string   `protobuf:"bytes,4,opt,name=policy_metric_mode,json=policyMetricMode,proto3" json:"policy_metric_mode,omitempty"`
	PolicyIsSspf             bool     `protobuf:"varint,5,opt,name=policy_is_sspf,json=policyIsSspf,proto3" json:"policy_is_sspf,omitempty"`
	PolicyIsAutorouteInclude bool     `protobuf:"varint,6,opt,name=policy_is_autoroute_include,json=policyIsAutorouteInclude,proto3" json:"policy_is_autoroute_include,omitempty"`
	PolicyStateIsValid       bool     `protobuf:"varint,7,opt,name=policy_state_is_valid,json=policyStateIsValid,proto3" json:"policy_state_is_valid,omitempty"`
	PolicyStateIsStale       bool     `protobuf:"varint,8,opt,name=policy_state_is_stale,json=policyStateIsStale,proto3" json:"policy_state_is_stale,omitempty"`
	PolicyStateIsIfhError    bool     `protobuf:"varint,9,opt,name=policy_state_is_ifh_error,json=policyStateIsIfhError,proto3" json:"policy_state_is_ifh_error,omitempty"`
	PolicyStateIsIdbPending  bool     `protobuf:"varint,10,opt,name=policy_state_is_idb_pending,json=policyStateIsIdbPending,proto3" json:"policy_state_is_idb_pending,omitempty"`
	PolicyUpdateTimestamp    uint32   `protobuf:"varint,11,opt,name=policy_update_timestamp,json=policyUpdateTimestamp,proto3" json:"policy_update_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *OspfShSrPolicy) Reset()         { *m = OspfShSrPolicy{} }
func (m *OspfShSrPolicy) String() string { return proto.CompactTextString(m) }
func (*OspfShSrPolicy) ProtoMessage()    {}
func (*OspfShSrPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_f586c47eaf48c07a, []int{1}
}

func (m *OspfShSrPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShSrPolicy.Unmarshal(m, b)
}
func (m *OspfShSrPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShSrPolicy.Marshal(b, m, deterministic)
}
func (m *OspfShSrPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShSrPolicy.Merge(m, src)
}
func (m *OspfShSrPolicy) XXX_Size() int {
	return xxx_messageInfo_OspfShSrPolicy.Size(m)
}
func (m *OspfShSrPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShSrPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShSrPolicy proto.InternalMessageInfo

func (m *OspfShSrPolicy) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *OspfShSrPolicy) GetPolicyIfHandle() string {
	if m != nil {
		return m.PolicyIfHandle
	}
	return ""
}

func (m *OspfShSrPolicy) GetPolicyMetric() int32 {
	if m != nil {
		return m.PolicyMetric
	}
	return 0
}

func (m *OspfShSrPolicy) GetPolicyMetricMode() string {
	if m != nil {
		return m.PolicyMetricMode
	}
	return ""
}

func (m *OspfShSrPolicy) GetPolicyIsSspf() bool {
	if m != nil {
		return m.PolicyIsSspf
	}
	return false
}

func (m *OspfShSrPolicy) GetPolicyIsAutorouteInclude() bool {
	if m != nil {
		return m.PolicyIsAutorouteInclude
	}
	return false
}

func (m *OspfShSrPolicy) GetPolicyStateIsValid() bool {
	if m != nil {
		return m.PolicyStateIsValid
	}
	return false
}

func (m *OspfShSrPolicy) GetPolicyStateIsStale() bool {
	if m != nil {
		return m.PolicyStateIsStale
	}
	return false
}

func (m *OspfShSrPolicy) GetPolicyStateIsIfhError() bool {
	if m != nil {
		return m.PolicyStateIsIfhError
	}
	return false
}

func (m *OspfShSrPolicy) GetPolicyStateIsIdbPending() bool {
	if m != nil {
		return m.PolicyStateIsIdbPending
	}
	return false
}

func (m *OspfShSrPolicy) GetPolicyUpdateTimestamp() uint32 {
	if m != nil {
		return m.PolicyUpdateTimestamp
	}
	return 0
}

type OspfShSrEndpoint struct {
	SrEndpoint                 string            `protobuf:"bytes,50,opt,name=sr_endpoint,json=srEndpoint,proto3" json:"sr_endpoint,omitempty"`
	SrRouterId                 uint32            `protobuf:"varint,51,opt,name=sr_router_id,json=srRouterId,proto3" json:"sr_router_id,omitempty"`
	SrAreaIdString             string            `protobuf:"bytes,52,opt,name=sr_area_id_string,json=srAreaIdString,proto3" json:"sr_area_id_string,omitempty"`
	SrPolicy                   []*OspfShSrPolicy `protobuf:"bytes,53,rep,name=sr_policy,json=srPolicy,proto3" json:"sr_policy,omitempty"`
	SrPolicyCount              uint32            `protobuf:"varint,54,opt,name=sr_policy_count,json=srPolicyCount,proto3" json:"sr_policy_count,omitempty"`
	SrValidPolicyCount         uint32            `protobuf:"varint,55,opt,name=sr_valid_policy_count,json=srValidPolicyCount,proto3" json:"sr_valid_policy_count,omitempty"`
	SrEndpStateIsValid         bool              `protobuf:"varint,56,opt,name=sr_endp_state_is_valid,json=srEndpStateIsValid,proto3" json:"sr_endp_state_is_valid,omitempty"`
	SrEndpStateIsStale         bool              `protobuf:"varint,57,opt,name=sr_endp_state_is_stale,json=srEndpStateIsStale,proto3" json:"sr_endp_state_is_stale,omitempty"`
	SrEndpStateIsUnres         bool              `protobuf:"varint,58,opt,name=sr_endp_state_is_unres,json=srEndpStateIsUnres,proto3" json:"sr_endp_state_is_unres,omitempty"`
	SrEndpResolutionFailReason string            `protobuf:"bytes,59,opt,name=sr_endp_resolution_fail_reason,json=srEndpResolutionFailReason,proto3" json:"sr_endp_resolution_fail_reason,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}          `json:"-"`
	XXX_unrecognized           []byte            `json:"-"`
	XXX_sizecache              int32             `json:"-"`
}

func (m *OspfShSrEndpoint) Reset()         { *m = OspfShSrEndpoint{} }
func (m *OspfShSrEndpoint) String() string { return proto.CompactTextString(m) }
func (*OspfShSrEndpoint) ProtoMessage()    {}
func (*OspfShSrEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_f586c47eaf48c07a, []int{2}
}

func (m *OspfShSrEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShSrEndpoint.Unmarshal(m, b)
}
func (m *OspfShSrEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShSrEndpoint.Marshal(b, m, deterministic)
}
func (m *OspfShSrEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShSrEndpoint.Merge(m, src)
}
func (m *OspfShSrEndpoint) XXX_Size() int {
	return xxx_messageInfo_OspfShSrEndpoint.Size(m)
}
func (m *OspfShSrEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShSrEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShSrEndpoint proto.InternalMessageInfo

func (m *OspfShSrEndpoint) GetSrEndpoint() string {
	if m != nil {
		return m.SrEndpoint
	}
	return ""
}

func (m *OspfShSrEndpoint) GetSrRouterId() uint32 {
	if m != nil {
		return m.SrRouterId
	}
	return 0
}

func (m *OspfShSrEndpoint) GetSrAreaIdString() string {
	if m != nil {
		return m.SrAreaIdString
	}
	return ""
}

func (m *OspfShSrEndpoint) GetSrPolicy() []*OspfShSrPolicy {
	if m != nil {
		return m.SrPolicy
	}
	return nil
}

func (m *OspfShSrEndpoint) GetSrPolicyCount() uint32 {
	if m != nil {
		return m.SrPolicyCount
	}
	return 0
}

func (m *OspfShSrEndpoint) GetSrValidPolicyCount() uint32 {
	if m != nil {
		return m.SrValidPolicyCount
	}
	return 0
}

func (m *OspfShSrEndpoint) GetSrEndpStateIsValid() bool {
	if m != nil {
		return m.SrEndpStateIsValid
	}
	return false
}

func (m *OspfShSrEndpoint) GetSrEndpStateIsStale() bool {
	if m != nil {
		return m.SrEndpStateIsStale
	}
	return false
}

func (m *OspfShSrEndpoint) GetSrEndpStateIsUnres() bool {
	if m != nil {
		return m.SrEndpStateIsUnres
	}
	return false
}

func (m *OspfShSrEndpoint) GetSrEndpResolutionFailReason() string {
	if m != nil {
		return m.SrEndpResolutionFailReason
	}
	return ""
}

func init() {
	proto.RegisterType((*OspfShSrEndpoint_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.segment_routing.sr_endpoint_policies.sr_endpoint_policy.ospf_sh_sr_endpoint_KEYS")
	proto.RegisterType((*OspfShSrPolicy)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.segment_routing.sr_endpoint_policies.sr_endpoint_policy.ospf_sh_sr_policy")
	proto.RegisterType((*OspfShSrEndpoint)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.segment_routing.sr_endpoint_policies.sr_endpoint_policy.ospf_sh_sr_endpoint")
}

func init() { proto.RegisterFile("ospf_sh_sr_endpoint.proto", fileDescriptor_f586c47eaf48c07a) }

var fileDescriptor_f586c47eaf48c07a = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xdf, 0x6e, 0xd3, 0x3c,
	0x18, 0xc6, 0xd5, 0x6f, 0xfb, 0xb6, 0xce, 0xdd, 0xf6, 0x6d, 0xfe, 0x34, 0xf0, 0x40, 0x82, 0x32,
	0x10, 0x2a, 0x12, 0x8a, 0xc4, 0x36, 0xc6, 0xf8, 0x77, 0x30, 0xd0, 0x10, 0x11, 0x1a, 0x9a, 0x5a,
	0x86, 0xc4, 0x91, 0xe5, 0xc5, 0xce, 0x6a, 0x29, 0x89, 0xa3, 0xf7, 0x75, 0x2a, 0x76, 0x03, 0x5c,
	0x09, 0x87, 0xdc, 0x16, 0xf7, 0x81, 0x6c, 0x27, 0x5d, 0xbb, 0xe5, 0x9c, 0x93, 0x2a, 0x7d, 0x9e,
	0xe7, 0xe7, 0xd7, 0x6e, 0x1e, 0x97, 0x6c, 0x1b, 0x2c, 0x53, 0x8e, 0x63, 0x8e, 0xc0, 0x55, 0x21,
	0x4b, 0xa3, 0x0b, 0x1b, 0x95, 0x60, 0xac, 0xa1, 0x98, 0x68, 0x4c, 0x0c, 0xd7, 0x06, 0xf9, 0x77,
	0xe0, 0xba, 0x9c, 0xec, 0x73, 0x1f, 0x36, 0xa5, 0x82, 0xc8, 0x3d, 0xb9, 0x5c, 0xa2, 0x10, 0x15,
	0x36, 0x4f, 0xd1, 0x04, 0x52, 0xff, 0x11, 0xa1, 0xba, 0xc8, 0x55, 0x61, 0x39, 0x98, 0xca, 0xea,
	0xe2, 0x22, 0x9a, 0x59, 0x9e, 0x97, 0x26, 0xd3, 0x89, 0x56, 0x78, 0x53, 0xbc, 0xdc, 0xb9, 0x24,
	0xac, 0x65, 0x47, 0xfc, 0xd3, 0xf1, 0xb7, 0x11, 0x7d, 0x40, 0x56, 0xeb, 0x39, 0xbc, 0x10, 0xb9,
	0x62, 0x9d, 0x7e, 0x67, 0xb0, 0x32, 0xec, 0xd5, 0xda, 0x67, 0x91, 0x2b, 0xba, 0x4d, 0xba, 0x13,
	0x48, 0x83, 0xfd, 0x8f, 0xb7, 0x97, 0x27, 0x90, 0x7a, 0xeb, 0x3e, 0xe9, 0x4d, 0x97, 0xd3, 0x25,
	0x5b, 0xf0, 0x2e, 0x69, 0xa4, 0xb8, 0xdc, 0xf9, 0xb5, 0x48, 0x36, 0x67, 0x66, 0x87, 0x0d, 0x39,
	0x2c, 0x3c, 0xcd, 0xce, 0x24, 0x41, 0xf2, 0xeb, 0x0e, 0xc8, 0x46, 0x1d, 0xd0, 0x29, 0x1f, 0x8b,
	0x42, 0x66, 0xcd, 0xe8, 0xf5, 0xa0, 0xc7, 0xe9, 0x47, 0xaf, 0xd2, 0x87, 0x64, 0xad, 0x4e, 0xe6,
	0xca, 0x82, 0x4e, 0xfc, 0x1e, 0x36, 0x87, 0xab, 0x41, 0x3c, 0xf1, 0x1a, 0x7d, 0x4a, 0xe8, 0x5c,
	0x88, 0xe7, 0x46, 0x2a, 0xb6, 0xe8, 0x17, 0xdc, 0x98, 0x4d, 0x9e, 0x18, 0xa9, 0xe8, 0x23, 0xb2,
	0xde, 0x0c, 0x47, 0x8e, 0x58, 0xa6, 0xec, 0xdf, 0x7e, 0x67, 0xd0, 0x6d, 0xd6, 0x8c, 0x71, 0x84,
	0x65, 0x4a, 0xdf, 0x92, 0xbb, 0x57, 0x29, 0x51, 0x59, 0xe3, 0xde, 0x8a, 0xe2, 0xba, 0x48, 0xb2,
	0x4a, 0x2a, 0xb6, 0xe4, 0x11, 0xd6, 0x20, 0x47, 0x4d, 0x20, 0x0e, 0x3e, 0x7d, 0x46, 0xb6, 0x6a,
	0x1c, 0xad, 0x70, 0x1c, 0xf2, 0x89, 0xc8, 0xb4, 0x64, 0xcb, 0x1e, 0xac, 0xf7, 0x3b, 0x72, 0x5e,
	0x8c, 0x5f, 0x9d, 0xd3, 0x86, 0xa0, 0x15, 0x99, 0x62, 0xdd, 0x16, 0x64, 0xe4, 0x1c, 0x7a, 0x48,
	0xb6, 0xaf, 0x23, 0x3a, 0x1d, 0x73, 0x05, 0x60, 0x80, 0xad, 0x78, 0x6c, 0x6b, 0x0e, 0x8b, 0xd3,
	0xf1, 0xb1, 0x33, 0xe9, 0x9b, 0xe9, 0xf1, 0xae, 0x48, 0x79, 0xce, 0x4b, 0x55, 0x48, 0x5d, 0x5c,
	0x30, 0xe2, 0xd9, 0xdb, 0xf3, 0xac, 0x3c, 0x3f, 0x0d, 0x36, 0x3d, 0x20, 0xb5, 0xc5, 0xab, 0x52,
	0x3a, 0xdc, 0xea, 0x5c, 0xa1, 0x15, 0x79, 0xc9, 0x7a, 0xfd, 0xce, 0x60, 0xad, 0x99, 0x7a, 0xe6,
	0xdd, 0x2f, 0x8d, 0xb9, 0xf3, 0x7b, 0x91, 0xfc, 0xdf, 0x52, 0x55, 0x57, 0x98, 0x99, 0xaf, 0x6c,
	0x37, 0x14, 0x06, 0xe1, 0xb8, 0x09, 0xf4, 0xc9, 0x2a, 0x82, 0xbf, 0x18, 0x0a, 0xb8, 0x96, 0x6c,
	0xcf, 0x4f, 0x21, 0x08, 0x43, 0x2f, 0xc5, 0x92, 0x3e, 0x21, 0x9b, 0x08, 0x5c, 0x80, 0x12, 0x5c,
	0x4b, 0x8e, 0x16, 0xdc, 0x31, 0xf6, 0x43, 0xa7, 0x10, 0x8e, 0x40, 0x89, 0x58, 0x8e, 0xbc, 0x4a,
	0x7f, 0x76, 0xc8, 0xca, 0xb4, 0xac, 0xec, 0x79, 0x7f, 0x61, 0xd0, 0xdb, 0xfd, 0xd1, 0x89, 0xfe,
	0xc2, 0xd5, 0x8d, 0x6e, 0xdc, 0x9d, 0x61, 0x17, 0xe1, 0x34, 0xdc, 0xa2, 0xc7, 0xe4, 0xbf, 0xa9,
	0xcc, 0x13, 0x53, 0x15, 0x96, 0x1d, 0xf8, 0x63, 0xaf, 0x35, 0x91, 0xf7, 0x4e, 0x74, 0xbd, 0x41,
	0x08, 0xed, 0x9a, 0x4f, 0xbf, 0xf0, 0x69, 0x8a, 0xe0, 0xfb, 0x35, 0x8b, 0xec, 0x92, 0x5b, 0xf5,
	0x66, 0xae, 0xd7, 0xf3, 0x30, 0x74, 0x2d, 0xfc, 0xf4, 0x73, 0xf5, 0x6c, 0x63, 0x42, 0x3f, 0x5f,
	0xb6, 0x30, 0xa1, 0x9f, 0x6d, 0x4c, 0x55, 0x80, 0x42, 0xf6, 0xaa, 0x85, 0x39, 0x73, 0x0e, 0x7d,
	0x47, 0xee, 0x35, 0x0c, 0x28, 0x34, 0x59, 0x65, 0xb5, 0x29, 0x78, 0x2a, 0x74, 0xc6, 0x41, 0x09,
	0x34, 0x05, 0x7b, 0xed, 0xdf, 0xea, 0x9d, 0xc0, 0x0e, 0xa7, 0x99, 0x0f, 0x42, 0x67, 0x43, 0x9f,
	0x38, 0x5f, 0xf2, 0xff, 0xc6, 0x7b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x78, 0x26, 0xfa,
	0xaa, 0x05, 0x00, 0x00,
}
