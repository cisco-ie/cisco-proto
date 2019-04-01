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
// source: telemetry_show_summary.proto

package cisco_ios_xr_telemetry_model_driven_oper_telemetry_model_driven_summary

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

type TelemetryShowSummary_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetryShowSummary_KEYS) Reset()         { *m = TelemetryShowSummary_KEYS{} }
func (m *TelemetryShowSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*TelemetryShowSummary_KEYS) ProtoMessage()    {}
func (*TelemetryShowSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc55f65eaffa515, []int{0}
}

func (m *TelemetryShowSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryShowSummary_KEYS.Unmarshal(m, b)
}
func (m *TelemetryShowSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryShowSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *TelemetryShowSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryShowSummary_KEYS.Merge(m, src)
}
func (m *TelemetryShowSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_TelemetryShowSummary_KEYS.Size(m)
}
func (m *TelemetryShowSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryShowSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryShowSummary_KEYS proto.InternalMessageInfo

type TelemetryShowSummary struct {
	NumOfSubscriptions          uint32   `protobuf:"varint,50,opt,name=num_of_subscriptions,json=numOfSubscriptions,proto3" json:"num_of_subscriptions,omitempty"`
	NumOfActiveSubscriptions    uint32   `protobuf:"varint,51,opt,name=num_of_active_subscriptions,json=numOfActiveSubscriptions,proto3" json:"num_of_active_subscriptions,omitempty"`
	NumOfPausedSubscriptions    uint32   `protobuf:"varint,52,opt,name=num_of_paused_subscriptions,json=numOfPausedSubscriptions,proto3" json:"num_of_paused_subscriptions,omitempty"`
	NumOfDestinationGroups      uint32   `protobuf:"varint,53,opt,name=num_of_destination_groups,json=numOfDestinationGroups,proto3" json:"num_of_destination_groups,omitempty"`
	NumOfDestinations           uint32   `protobuf:"varint,54,opt,name=num_of_destinations,json=numOfDestinations,proto3" json:"num_of_destinations,omitempty"`
	NumOfTcpDialouts            uint32   `protobuf:"varint,55,opt,name=num_of_tcp_dialouts,json=numOfTcpDialouts,proto3" json:"num_of_tcp_dialouts,omitempty"`
	NumOfUdpDialouts            uint32   `protobuf:"varint,56,opt,name=num_of_udp_dialouts,json=numOfUdpDialouts,proto3" json:"num_of_udp_dialouts,omitempty"`
	NumOfGrpcTlsDialouts        uint32   `protobuf:"varint,57,opt,name=num_of_grpc_tls_dialouts,json=numOfGrpcTlsDialouts,proto3" json:"num_of_grpc_tls_dialouts,omitempty"`
	NumOfGrpcNonTlsDialouts     uint32   `protobuf:"varint,58,opt,name=num_of_grpc_non_tls_dialouts,json=numOfGrpcNonTlsDialouts,proto3" json:"num_of_grpc_non_tls_dialouts,omitempty"`
	NumOfDialins                uint32   `protobuf:"varint,59,opt,name=num_of_dialins,json=numOfDialins,proto3" json:"num_of_dialins,omitempty"`
	NumOfActiveDestinations     uint32   `protobuf:"varint,60,opt,name=num_of_active_destinations,json=numOfActiveDestinations,proto3" json:"num_of_active_destinations,omitempty"`
	NumOfConnectedSessions      uint32   `protobuf:"varint,61,opt,name=num_of_connected_sessions,json=numOfConnectedSessions,proto3" json:"num_of_connected_sessions,omitempty"`
	NumOfConnectingSessions     uint32   `protobuf:"varint,62,opt,name=num_of_connecting_sessions,json=numOfConnectingSessions,proto3" json:"num_of_connecting_sessions,omitempty"`
	NumOfSensorGroups           uint32   `protobuf:"varint,63,opt,name=num_of_sensor_groups,json=numOfSensorGroups,proto3" json:"num_of_sensor_groups,omitempty"`
	NumOfUniqueSensorPaths      uint32   `protobuf:"varint,64,opt,name=num_of_unique_sensor_paths,json=numOfUniqueSensorPaths,proto3" json:"num_of_unique_sensor_paths,omitempty"`
	NumOfSensorPaths            uint32   `protobuf:"varint,65,opt,name=num_of_sensor_paths,json=numOfSensorPaths,proto3" json:"num_of_sensor_paths,omitempty"`
	NumOfNotResolvedSensorPaths uint32   `protobuf:"varint,66,opt,name=num_of_not_resolved_sensor_paths,json=numOfNotResolvedSensorPaths,proto3" json:"num_of_not_resolved_sensor_paths,omitempty"`
	NumOfActiveSensorPaths      uint32   `protobuf:"varint,67,opt,name=num_of_active_sensor_paths,json=numOfActiveSensorPaths,proto3" json:"num_of_active_sensor_paths,omitempty"`
	MaxSensorPaths              uint32   `protobuf:"varint,68,opt,name=max_sensor_paths,json=maxSensorPaths,proto3" json:"max_sensor_paths,omitempty"`
	MaxContainersPerPath        uint32   `protobuf:"varint,69,opt,name=max_containers_per_path,json=maxContainersPerPath,proto3" json:"max_containers_per_path,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *TelemetryShowSummary) Reset()         { *m = TelemetryShowSummary{} }
func (m *TelemetryShowSummary) String() string { return proto.CompactTextString(m) }
func (*TelemetryShowSummary) ProtoMessage()    {}
func (*TelemetryShowSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc55f65eaffa515, []int{1}
}

func (m *TelemetryShowSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryShowSummary.Unmarshal(m, b)
}
func (m *TelemetryShowSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryShowSummary.Marshal(b, m, deterministic)
}
func (m *TelemetryShowSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryShowSummary.Merge(m, src)
}
func (m *TelemetryShowSummary) XXX_Size() int {
	return xxx_messageInfo_TelemetryShowSummary.Size(m)
}
func (m *TelemetryShowSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryShowSummary.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryShowSummary proto.InternalMessageInfo

func (m *TelemetryShowSummary) GetNumOfSubscriptions() uint32 {
	if m != nil {
		return m.NumOfSubscriptions
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfActiveSubscriptions() uint32 {
	if m != nil {
		return m.NumOfActiveSubscriptions
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfPausedSubscriptions() uint32 {
	if m != nil {
		return m.NumOfPausedSubscriptions
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfDestinationGroups() uint32 {
	if m != nil {
		return m.NumOfDestinationGroups
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfDestinations() uint32 {
	if m != nil {
		return m.NumOfDestinations
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfTcpDialouts() uint32 {
	if m != nil {
		return m.NumOfTcpDialouts
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfUdpDialouts() uint32 {
	if m != nil {
		return m.NumOfUdpDialouts
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfGrpcTlsDialouts() uint32 {
	if m != nil {
		return m.NumOfGrpcTlsDialouts
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfGrpcNonTlsDialouts() uint32 {
	if m != nil {
		return m.NumOfGrpcNonTlsDialouts
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfDialins() uint32 {
	if m != nil {
		return m.NumOfDialins
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfActiveDestinations() uint32 {
	if m != nil {
		return m.NumOfActiveDestinations
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfConnectedSessions() uint32 {
	if m != nil {
		return m.NumOfConnectedSessions
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfConnectingSessions() uint32 {
	if m != nil {
		return m.NumOfConnectingSessions
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfSensorGroups() uint32 {
	if m != nil {
		return m.NumOfSensorGroups
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfUniqueSensorPaths() uint32 {
	if m != nil {
		return m.NumOfUniqueSensorPaths
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfSensorPaths() uint32 {
	if m != nil {
		return m.NumOfSensorPaths
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfNotResolvedSensorPaths() uint32 {
	if m != nil {
		return m.NumOfNotResolvedSensorPaths
	}
	return 0
}

func (m *TelemetryShowSummary) GetNumOfActiveSensorPaths() uint32 {
	if m != nil {
		return m.NumOfActiveSensorPaths
	}
	return 0
}

func (m *TelemetryShowSummary) GetMaxSensorPaths() uint32 {
	if m != nil {
		return m.MaxSensorPaths
	}
	return 0
}

func (m *TelemetryShowSummary) GetMaxContainersPerPath() uint32 {
	if m != nil {
		return m.MaxContainersPerPath
	}
	return 0
}

func init() {
	proto.RegisterType((*TelemetryShowSummary_KEYS)(nil), "cisco_ios_xr_telemetry_model_driven_oper.telemetry_model_driven.summary.telemetry_show_summary_KEYS")
	proto.RegisterType((*TelemetryShowSummary)(nil), "cisco_ios_xr_telemetry_model_driven_oper.telemetry_model_driven.summary.telemetry_show_summary")
}

func init() { proto.RegisterFile("telemetry_show_summary.proto", fileDescriptor_edc55f65eaffa515) }

var fileDescriptor_edc55f65eaffa515 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd4, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x07, 0x70, 0x71, 0x41, 0xc8, 0x82, 0x69, 0x84, 0x69, 0x0b, 0x6c, 0x48, 0xd3, 0xc4, 0x61,
	0x17, 0x0a, 0x62, 0x6c, 0xb0, 0x8e, 0x02, 0xa3, 0xad, 0x7a, 0x40, 0x1a, 0xd5, 0xba, 0x1d, 0x38,
	0x59, 0x9e, 0xe3, 0x75, 0x96, 0x12, 0xdb, 0xf8, 0x39, 0xa5, 0xfb, 0xbc, 0x7c, 0x11, 0xd4, 0x17,
	0x37, 0x79, 0x29, 0xe5, 0xea, 0xff, 0xff, 0x67, 0x67, 0x7e, 0x5e, 0xd9, 0x5e, 0x50, 0xb9, 0x2a,
	0x54, 0xf0, 0xf7, 0x1c, 0xee, 0xec, 0x6f, 0x0e, 0x65, 0x51, 0x08, 0x7f, 0xdf, 0x71, 0xde, 0x06,
	0x9b, 0x8c, 0xa4, 0x06, 0x69, 0xb9, 0xb6, 0xc0, 0xe7, 0x9e, 0x37, 0xd5, 0xc2, 0x66, 0x2a, 0xe7,
	0x99, 0xd7, 0x33, 0x65, 0xb8, 0x75, 0xca, 0x77, 0xd6, 0x67, 0x9d, 0xb8, 0xdd, 0xc1, 0x4b, 0xb6,
	0xbb, 0xfe, 0x20, 0xfe, 0x7d, 0xf8, 0x73, 0x72, 0xf0, 0xe7, 0x11, 0xdb, 0x5e, 0x9f, 0x27, 0x6f,
	0xd9, 0x96, 0x29, 0x0b, 0x6e, 0x6f, 0x39, 0x94, 0x37, 0x20, 0xbd, 0x76, 0x41, 0x5b, 0x03, 0xe9,
	0xbb, 0xfd, 0x07, 0x87, 0x4f, 0x2e, 0x13, 0x53, 0x16, 0x3f, 0x6e, 0x27, 0x34, 0x49, 0x7a, 0x6c,
	0x37, 0x0a, 0x21, 0x83, 0x9e, 0xa9, 0x15, 0x78, 0x84, 0x30, 0x45, 0x78, 0x8e, 0x85, 0xff, 0x71,
	0x27, 0x4a, 0x50, 0xd9, 0x0a, 0x7f, 0x4f, 0xf8, 0x18, 0x0b, 0x6d, 0x7e, 0xca, 0x9e, 0x47, 0x9e,
	0x29, 0x08, 0xda, 0x88, 0xc5, 0x32, 0x9f, 0x7a, 0x5b, 0x3a, 0x48, 0x8f, 0x11, 0x6f, 0x23, 0x1e,
	0x34, 0xf1, 0x08, 0xd3, 0xa4, 0xc3, 0x9e, 0xfd, 0x4b, 0x21, 0x3d, 0x41, 0xf4, 0x74, 0x15, 0x41,
	0xf2, 0xba, 0xee, 0x07, 0xe9, 0x78, 0xa6, 0x45, 0x6e, 0xcb, 0x00, 0xe9, 0x07, 0xec, 0x6f, 0x62,
	0xff, 0x4a, 0xba, 0x41, 0x5c, 0x27, 0xf5, 0x32, 0x23, 0xf5, 0x8f, 0xa4, 0x7e, 0x9d, 0x35, 0xf5,
	0x13, 0x96, 0xc6, 0xfa, 0xd4, 0x3b, 0xc9, 0x43, 0x0e, 0x8d, 0x39, 0x45, 0xb3, 0x85, 0x66, 0xe4,
	0x9d, 0xbc, 0xca, 0xa1, 0x76, 0x3d, 0xb6, 0x47, 0x9d, 0xb1, 0xa6, 0x6d, 0xbb, 0x68, 0x77, 0x6a,
	0x7b, 0x61, 0x0d, 0xe5, 0xaf, 0xd8, 0xc6, 0xf2, 0x12, 0xb4, 0xc8, 0xb5, 0x81, 0xf4, 0x0c, 0xc1,
	0xe3, 0xea, 0xef, 0xaf, 0xd6, 0x92, 0x33, 0xf6, 0xa2, 0x3d, 0xe3, 0xd6, 0x8d, 0x7d, 0x22, 0x47,
	0x54, 0x23, 0x6e, 0xdd, 0x5b, 0x33, 0x22, 0x69, 0x8d, 0x51, 0x32, 0x2c, 0x86, 0xac, 0x00, 0xd0,
	0xf6, 0xc8, 0x88, 0xfa, 0xcb, 0x78, 0x12, 0x53, 0x72, 0x6e, 0xa4, 0xda, 0x4c, 0x1b, 0xfb, 0x99,
	0x9c, 0xdb, 0xaf, 0xf3, 0x1a, 0xbf, 0x69, 0x9e, 0xb2, 0x32, 0x60, 0xfd, 0xf2, 0x55, 0x7c, 0x21,
	0x03, 0x9e, 0x60, 0x12, 0x1f, 0x44, 0xb7, 0x3e, 0xad, 0x34, 0xfa, 0x57, 0xa9, 0x96, 0xce, 0x89,
	0x70, 0x07, 0xe9, 0x57, 0xf2, 0xa5, 0xd7, 0x98, 0x57, 0x78, 0xbc, 0x48, 0xc9, 0xb4, 0x5b, 0xe8,
	0x9c, 0x4c, 0x9b, 0xd6, 0x87, 0x6c, 0x3f, 0xd6, 0x8d, 0x0d, 0xdc, 0x2b, 0xb0, 0xf9, 0x0c, 0xaf,
	0x85, 0xd8, 0x6f, 0x68, 0x77, 0xd1, 0x5e, 0xd8, 0x70, 0x19, 0x4b, 0x74, 0x9b, 0xee, 0xea, 0x5c,
	0x5a, 0x1b, 0xf4, 0xc9, 0x17, 0xc7, 0x7f, 0x3d, 0x62, 0x0f, 0xd9, 0x66, 0x21, 0xe6, 0x6d, 0x31,
	0x40, 0xb1, 0x51, 0x88, 0x39, 0x6d, 0x1e, 0xb3, 0x9d, 0x45, 0x53, 0x5a, 0x13, 0x84, 0x36, 0xca,
	0x03, 0x77, 0xaa, 0x12, 0xe9, 0xb0, 0x7a, 0x99, 0x85, 0x98, 0xf7, 0xeb, 0x74, 0xac, 0xd0, 0xdd,
	0x3c, 0xc4, 0x1f, 0xb5, 0xa3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xc4, 0xb4, 0xfc, 0xf4,
	0x04, 0x00, 0x00,
}
