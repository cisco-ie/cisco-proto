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
// source: igmp_edm_traffic_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_mld_standby_vrfs_vrf_traffic_counters

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

type IgmpEdmTrafficBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmTrafficBag_KEYS) Reset()         { *m = IgmpEdmTrafficBag_KEYS{} }
func (m *IgmpEdmTrafficBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmTrafficBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmTrafficBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3a8d350bb40983, []int{0}
}

func (m *IgmpEdmTrafficBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmTrafficBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmTrafficBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmTrafficBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmTrafficBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmTrafficBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmTrafficBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmTrafficBag_KEYS.Size(m)
}
func (m *IgmpEdmTrafficBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmTrafficBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmTrafficBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmTrafficBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type IgmpEdmTrafficBag struct {
	ElapsedTime                   uint32   `protobuf:"varint,50,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	PacketsIn                     uint32   `protobuf:"varint,51,opt,name=packets_in,json=packetsIn,proto3" json:"packets_in,omitempty"`
	PacketsOut                    uint32   `protobuf:"varint,52,opt,name=packets_out,json=packetsOut,proto3" json:"packets_out,omitempty"`
	FormatErrors                  uint32   `protobuf:"varint,53,opt,name=format_errors,json=formatErrors,proto3" json:"format_errors,omitempty"`
	PacketManagerInputErrors      uint32   `protobuf:"varint,54,opt,name=packet_manager_input_errors,json=packetManagerInputErrors,proto3" json:"packet_manager_input_errors,omitempty"`
	PacketManagerOutputErrors     uint32   `protobuf:"varint,55,opt,name=packet_manager_output_errors,json=packetManagerOutputErrors,proto3" json:"packet_manager_output_errors,omitempty"`
	ChecksumErrors                uint32   `protobuf:"varint,56,opt,name=checksum_errors,json=checksumErrors,proto3" json:"checksum_errors,omitempty"`
	ReceiveSocketErrors           uint32   `protobuf:"varint,57,opt,name=receive_socket_errors,json=receiveSocketErrors,proto3" json:"receive_socket_errors,omitempty"`
	SocketErrors                  uint32   `protobuf:"varint,58,opt,name=socket_errors,json=socketErrors,proto3" json:"socket_errors,omitempty"`
	BadScopeErrors                uint32   `protobuf:"varint,59,opt,name=bad_scope_errors,json=badScopeErrors,proto3" json:"bad_scope_errors,omitempty"`
	AuxillaryDataLengthErrors     uint32   `protobuf:"varint,60,opt,name=auxillary_data_length_errors,json=auxillaryDataLengthErrors,proto3" json:"auxillary_data_length_errors,omitempty"`
	InvalidSourceAddressErrors    uint32   `protobuf:"varint,61,opt,name=invalid_source_address_errors,json=invalidSourceAddressErrors,proto3" json:"invalid_source_address_errors,omitempty"`
	NoSocketConnection            uint32   `protobuf:"varint,62,opt,name=no_socket_connection,json=noSocketConnection,proto3" json:"no_socket_connection,omitempty"`
	MiscellaneousErrors           uint32   `protobuf:"varint,63,opt,name=miscellaneous_errors,json=miscellaneousErrors,proto3" json:"miscellaneous_errors,omitempty"`
	InputQueries                  uint32   `protobuf:"varint,64,opt,name=input_queries,json=inputQueries,proto3" json:"input_queries,omitempty"`
	InputReports                  uint32   `protobuf:"varint,65,opt,name=input_reports,json=inputReports,proto3" json:"input_reports,omitempty"`
	InputLeaves                   uint32   `protobuf:"varint,66,opt,name=input_leaves,json=inputLeaves,proto3" json:"input_leaves,omitempty"`
	InputMtrace                   uint32   `protobuf:"varint,67,opt,name=input_mtrace,json=inputMtrace,proto3" json:"input_mtrace,omitempty"`
	InputDvmrp                    uint32   `protobuf:"varint,68,opt,name=input_dvmrp,json=inputDvmrp,proto3" json:"input_dvmrp,omitempty"`
	InputPim                      uint32   `protobuf:"varint,69,opt,name=input_pim,json=inputPim,proto3" json:"input_pim,omitempty"`
	OutputQueries                 uint32   `protobuf:"varint,70,opt,name=output_queries,json=outputQueries,proto3" json:"output_queries,omitempty"`
	OutputReports                 uint32   `protobuf:"varint,71,opt,name=output_reports,json=outputReports,proto3" json:"output_reports,omitempty"`
	OutputLeaves                  uint32   `protobuf:"varint,72,opt,name=output_leaves,json=outputLeaves,proto3" json:"output_leaves,omitempty"`
	OutputMtrace                  uint32   `protobuf:"varint,73,opt,name=output_mtrace,json=outputMtrace,proto3" json:"output_mtrace,omitempty"`
	OutputDvmrp                   uint32   `protobuf:"varint,74,opt,name=output_dvmrp,json=outputDvmrp,proto3" json:"output_dvmrp,omitempty"`
	OutputPim                     uint32   `protobuf:"varint,75,opt,name=output_pim,json=outputPim,proto3" json:"output_pim,omitempty"`
	GetPacketFailure              uint32   `protobuf:"varint,76,opt,name=get_packet_failure,json=getPacketFailure,proto3" json:"get_packet_failure,omitempty"`
	OutputNoParentInterfaceHandle uint32   `protobuf:"varint,77,opt,name=output_no_parent_interface_handle,json=outputNoParentInterfaceHandle,proto3" json:"output_no_parent_interface_handle,omitempty"`
	InputNoIdb                    uint32   `protobuf:"varint,78,opt,name=input_no_idb,json=inputNoIdb,proto3" json:"input_no_idb,omitempty"`
	InputNoVrfInIdb               uint32   `protobuf:"varint,79,opt,name=input_no_vrf_in_idb,json=inputNoVrfInIdb,proto3" json:"input_no_vrf_in_idb,omitempty"`
	InputDisabledIdb              uint32   `protobuf:"varint,80,opt,name=input_disabled_idb,json=inputDisabledIdb,proto3" json:"input_disabled_idb,omitempty"`
	InputMartianAddress           uint32   `protobuf:"varint,81,opt,name=input_martian_address,json=inputMartianAddress,proto3" json:"input_martian_address,omitempty"`
	InputNoAssignedVrfId          uint32   `protobuf:"varint,82,opt,name=input_no_assigned_vrf_id,json=inputNoAssignedVrfId,proto3" json:"input_no_assigned_vrf_id,omitempty"`
	InputNoVrfMtrace              uint32   `protobuf:"varint,83,opt,name=input_no_vrf_mtrace,json=inputNoVrfMtrace,proto3" json:"input_no_vrf_mtrace,omitempty"`
	InputNoPlatformSupportMtrace  uint32   `protobuf:"varint,84,opt,name=input_no_platform_support_mtrace,json=inputNoPlatformSupportMtrace,proto3" json:"input_no_platform_support_mtrace,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *IgmpEdmTrafficBag) Reset()         { *m = IgmpEdmTrafficBag{} }
func (m *IgmpEdmTrafficBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmTrafficBag) ProtoMessage()    {}
func (*IgmpEdmTrafficBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3a8d350bb40983, []int{1}
}

func (m *IgmpEdmTrafficBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmTrafficBag.Unmarshal(m, b)
}
func (m *IgmpEdmTrafficBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmTrafficBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmTrafficBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmTrafficBag.Merge(m, src)
}
func (m *IgmpEdmTrafficBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmTrafficBag.Size(m)
}
func (m *IgmpEdmTrafficBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmTrafficBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmTrafficBag proto.InternalMessageInfo

func (m *IgmpEdmTrafficBag) GetElapsedTime() uint32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetPacketsIn() uint32 {
	if m != nil {
		return m.PacketsIn
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetPacketsOut() uint32 {
	if m != nil {
		return m.PacketsOut
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetFormatErrors() uint32 {
	if m != nil {
		return m.FormatErrors
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetPacketManagerInputErrors() uint32 {
	if m != nil {
		return m.PacketManagerInputErrors
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetPacketManagerOutputErrors() uint32 {
	if m != nil {
		return m.PacketManagerOutputErrors
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetChecksumErrors() uint32 {
	if m != nil {
		return m.ChecksumErrors
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetReceiveSocketErrors() uint32 {
	if m != nil {
		return m.ReceiveSocketErrors
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetSocketErrors() uint32 {
	if m != nil {
		return m.SocketErrors
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetBadScopeErrors() uint32 {
	if m != nil {
		return m.BadScopeErrors
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetAuxillaryDataLengthErrors() uint32 {
	if m != nil {
		return m.AuxillaryDataLengthErrors
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInvalidSourceAddressErrors() uint32 {
	if m != nil {
		return m.InvalidSourceAddressErrors
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetNoSocketConnection() uint32 {
	if m != nil {
		return m.NoSocketConnection
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetMiscellaneousErrors() uint32 {
	if m != nil {
		return m.MiscellaneousErrors
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputQueries() uint32 {
	if m != nil {
		return m.InputQueries
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputReports() uint32 {
	if m != nil {
		return m.InputReports
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputLeaves() uint32 {
	if m != nil {
		return m.InputLeaves
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputMtrace() uint32 {
	if m != nil {
		return m.InputMtrace
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputDvmrp() uint32 {
	if m != nil {
		return m.InputDvmrp
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputPim() uint32 {
	if m != nil {
		return m.InputPim
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetOutputQueries() uint32 {
	if m != nil {
		return m.OutputQueries
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetOutputReports() uint32 {
	if m != nil {
		return m.OutputReports
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetOutputLeaves() uint32 {
	if m != nil {
		return m.OutputLeaves
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetOutputMtrace() uint32 {
	if m != nil {
		return m.OutputMtrace
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetOutputDvmrp() uint32 {
	if m != nil {
		return m.OutputDvmrp
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetOutputPim() uint32 {
	if m != nil {
		return m.OutputPim
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetGetPacketFailure() uint32 {
	if m != nil {
		return m.GetPacketFailure
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetOutputNoParentInterfaceHandle() uint32 {
	if m != nil {
		return m.OutputNoParentInterfaceHandle
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputNoIdb() uint32 {
	if m != nil {
		return m.InputNoIdb
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputNoVrfInIdb() uint32 {
	if m != nil {
		return m.InputNoVrfInIdb
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputDisabledIdb() uint32 {
	if m != nil {
		return m.InputDisabledIdb
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputMartianAddress() uint32 {
	if m != nil {
		return m.InputMartianAddress
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputNoAssignedVrfId() uint32 {
	if m != nil {
		return m.InputNoAssignedVrfId
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputNoVrfMtrace() uint32 {
	if m != nil {
		return m.InputNoVrfMtrace
	}
	return 0
}

func (m *IgmpEdmTrafficBag) GetInputNoPlatformSupportMtrace() uint32 {
	if m != nil {
		return m.InputNoPlatformSupportMtrace
	}
	return 0
}

func init() {
	proto.RegisterType((*IgmpEdmTrafficBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.standby.vrfs.vrf.traffic_counters.igmp_edm_traffic_bag_KEYS")
	proto.RegisterType((*IgmpEdmTrafficBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.standby.vrfs.vrf.traffic_counters.igmp_edm_traffic_bag")
}

func init() { proto.RegisterFile("igmp_edm_traffic_bag.proto", fileDescriptor_3d3a8d350bb40983) }

var fileDescriptor_3d3a8d350bb40983 = []byte{
	// 824 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x95, 0x5f, 0x6f, 0x14, 0x37,
	0x14, 0xc5, 0x95, 0x97, 0x96, 0x98, 0x04, 0x90, 0x09, 0x92, 0x03, 0x44, 0x4d, 0x88, 0xaa, 0xe6,
	0x81, 0xae, 0xda, 0x40, 0xd3, 0xbf, 0x94, 0x6e, 0x49, 0xd2, 0x6c, 0xc9, 0x9f, 0x65, 0x17, 0x55,
	0xea, 0x93, 0xe5, 0x1d, 0xdf, 0xd9, 0x58, 0xcc, 0xd8, 0x53, 0xdb, 0xb3, 0x82, 0xef, 0xd5, 0x0f,
	0x88, 0xe6, 0x5e, 0x7b, 0x77, 0x82, 0x78, 0x89, 0x94, 0x73, 0x7f, 0xe7, 0x8e, 0x8f, 0x7d, 0xed,
	0x65, 0x0f, 0xcd, 0xbc, 0x6e, 0x24, 0xe8, 0x5a, 0x46, 0xaf, 0xca, 0xd2, 0x14, 0x72, 0xa6, 0xe6,
	0x83, 0xc6, 0xbb, 0xe8, 0xf8, 0xb0, 0x30, 0xa1, 0x70, 0xd2, 0xb8, 0x20, 0xdf, 0x7b, 0x69, 0x9a,
	0xc5, 0x73, 0x89, 0xb4, 0x6b, 0xc0, 0x0f, 0xea, 0x4a, 0x0f, 0x42, 0x54, 0x56, 0xcf, 0x3e, 0x0c,
	0x16, 0xbe, 0x0c, 0xdd, 0x9f, 0x41, 0xee, 0x51, 0xb8, 0xd6, 0x46, 0xf0, 0xe1, 0xc9, 0x11, 0xdb,
	0xfe, 0xdc, 0x07, 0xe4, 0xeb, 0x93, 0x7f, 0xa7, 0x7c, 0x9b, 0xdd, 0x5a, 0xf8, 0x52, 0x5a, 0x55,
	0x83, 0x58, 0xdb, 0x5d, 0x3b, 0x58, 0x9f, 0x7c, 0xb9, 0xf0, 0xe5, 0xa5, 0xaa, 0xe1, 0xc9, 0xff,
	0x1b, 0x6c, 0xeb, 0x73, 0x46, 0xbe, 0xc7, 0x36, 0xa0, 0x52, 0x4d, 0x00, 0x2d, 0xa3, 0xa9, 0x41,
	0x1c, 0xee, 0xae, 0x1d, 0x6c, 0x4e, 0x6e, 0x27, 0xed, 0xad, 0xa9, 0x81, 0xef, 0x30, 0xd6, 0xa8,
	0xe2, 0x1d, 0xc4, 0x20, 0x8d, 0x15, 0xcf, 0x10, 0x58, 0x4f, 0xca, 0xc8, 0xf2, 0xaf, 0xd8, 0xed,
	0x5c, 0x76, 0x6d, 0x14, 0xcf, 0xb1, 0x9e, 0x1d, 0x57, 0x6d, 0xe4, 0xfb, 0x6c, 0xb3, 0x74, 0xbe,
	0x56, 0x51, 0x82, 0xf7, 0xce, 0x07, 0xf1, 0x03, 0x22, 0x1b, 0x24, 0x9e, 0xa0, 0xc6, 0x5f, 0xb0,
	0x47, 0x64, 0x91, 0xb5, 0xb2, 0x6a, 0x0e, 0x5e, 0x1a, 0xdb, 0xb4, 0x4b, 0xcb, 0x11, 0x5a, 0x04,
	0x21, 0x17, 0x44, 0x8c, 0x3a, 0x20, 0xd9, 0x5f, 0xb2, 0xc7, 0x9f, 0xd8, 0x5d, 0x1b, 0x7b, 0xfe,
	0x1f, 0xd1, 0xbf, 0x7d, 0xc3, 0x7f, 0x85, 0x44, 0x6a, 0xf0, 0x0d, 0xbb, 0x5b, 0x5c, 0x43, 0xf1,
	0x2e, 0xb4, 0x75, 0xf6, 0xfc, 0x84, 0x9e, 0x3b, 0x59, 0x4e, 0xe0, 0x21, 0x7b, 0xe0, 0xa1, 0x00,
	0xb3, 0x00, 0x19, 0x1c, 0x7e, 0x31, 0xe1, 0x3f, 0x23, 0x7e, 0x3f, 0x15, 0xa7, 0x58, 0x4b, 0x9e,
	0x7d, 0xb6, 0x79, 0x93, 0xfd, 0x85, 0x76, 0x20, 0xf4, 0xa1, 0x03, 0x76, 0x6f, 0xa6, 0xb4, 0x0c,
	0x85, 0x6b, 0x20, 0x73, 0xbf, 0xd2, 0x12, 0x66, 0x4a, 0x4f, 0x3b, 0x79, 0x15, 0x56, 0xb5, 0xef,
	0x4d, 0x55, 0x29, 0xff, 0x41, 0x6a, 0x15, 0x95, 0xac, 0xc0, 0xce, 0xe3, 0x75, 0x76, 0xfd, 0x46,
	0x61, 0x97, 0xcc, 0xb1, 0x8a, 0xea, 0x1c, 0x89, 0xd4, 0x60, 0xc8, 0x76, 0x8c, 0x5d, 0xa8, 0xca,
	0x68, 0x19, 0x5c, 0xeb, 0x0b, 0x90, 0x4a, 0x6b, 0x0f, 0x21, 0xe4, 0x0e, 0x2f, 0xb0, 0xc3, 0xc3,
	0x04, 0x4d, 0x91, 0x19, 0x12, 0x92, 0x5a, 0x7c, 0xc7, 0xb6, 0xac, 0xcb, 0x3b, 0x50, 0x38, 0x6b,
	0xa1, 0x88, 0xc6, 0x59, 0xf1, 0x3b, 0x3a, 0xb9, 0x75, 0xb4, 0x01, 0xaf, 0x96, 0x15, 0xfe, 0x3d,
	0xdb, 0xaa, 0x4d, 0x28, 0xa0, 0xaa, 0x94, 0x05, 0xd7, 0x2e, 0xbf, 0xf5, 0x92, 0xf6, 0xed, 0x46,
	0x6d, 0xb5, 0x6f, 0x34, 0x05, 0xff, 0xb5, 0xe0, 0x0d, 0x04, 0xf1, 0x07, 0xed, 0x1b, 0x8a, 0x6f,
	0x48, 0x5b, 0x41, 0x1e, 0x1a, 0xe7, 0x63, 0x10, 0xc3, 0x1e, 0x34, 0x21, 0xad, 0x1b, 0x73, 0x82,
	0x2a, 0x50, 0x0b, 0x08, 0xe2, 0x4f, 0x1a, 0x73, 0xd4, 0xce, 0x51, 0x5a, 0x21, 0x75, 0xf4, 0xaa,
	0x00, 0xf1, 0xaa, 0x87, 0x5c, 0xa0, 0xd4, 0x8d, 0x3a, 0x21, 0x7a, 0x51, 0xfb, 0x46, 0x1c, 0xd3,
	0xa8, 0xa3, 0x74, 0xdc, 0x29, 0xfc, 0x11, 0x5b, 0x27, 0xa0, 0x31, 0xb5, 0x38, 0xc1, 0xf2, 0x2d,
	0x14, 0xc6, 0xa6, 0xe6, 0x5f, 0xb3, 0x3b, 0x69, 0x28, 0x73, 0x9c, 0x53, 0x24, 0x36, 0x49, 0xcd,
	0x79, 0x56, 0x58, 0x0e, 0xf4, 0x57, 0x1f, 0xcb, 0x89, 0xf6, 0x59, 0x12, 0x72, 0xa4, 0x33, 0x8a,
	0x4d, 0x62, 0xca, 0xb4, 0x82, 0x52, 0xa8, 0x51, 0x1f, 0x4a, 0xa9, 0xf6, 0x58, 0xfa, 0x3f, 0xc5,
	0xfa, 0x9b, 0x82, 0x93, 0x46, 0xb9, 0x76, 0x18, 0x4b, 0x48, 0x17, 0xec, 0x35, 0x3d, 0x01, 0xa4,
	0x74, 0xc9, 0x9e, 0x32, 0x3e, 0x87, 0x28, 0xd3, 0x0d, 0x2c, 0x95, 0xa9, 0x5a, 0x0f, 0xe2, 0x1c,
	0xb1, 0x7b, 0x73, 0x88, 0x63, 0x2c, 0x9c, 0x92, 0xce, 0xcf, 0xd8, 0x5e, 0x6a, 0x66, 0x9d, 0x6c,
	0x94, 0x07, 0x1b, 0xa5, 0xe9, 0x9e, 0xb7, 0x52, 0x15, 0x20, 0xaf, 0x95, 0xd5, 0x15, 0x88, 0x0b,
	0x34, 0xef, 0x10, 0x78, 0xe9, 0xc6, 0x88, 0x8d, 0x32, 0x75, 0x86, 0x10, 0xdf, 0xcd, 0x47, 0x66,
	0x9d, 0x34, 0x7a, 0x26, 0x2e, 0x7b, 0x07, 0x72, 0xe9, 0x46, 0x7a, 0xc6, 0x9f, 0xb2, 0xfb, 0x4b,
	0xa2, 0x7b, 0x1b, 0x8d, 0x45, 0xf0, 0x0a, 0xc1, 0xbb, 0x09, 0xfc, 0xc7, 0x97, 0x23, 0x4b, 0x34,
	0x4f, 0xe7, 0x6b, 0x82, 0x9a, 0x55, 0xa0, 0x11, 0x1e, 0x53, 0x0e, 0x3a, 0xe6, 0x54, 0xe8, 0xe8,
	0x43, 0xf6, 0x20, 0x0d, 0x8c, 0xf2, 0xd1, 0x28, 0x9b, 0x2f, 0x91, 0x78, 0x43, 0x13, 0x4d, 0x93,
	0x43, 0xb5, 0x74, 0x79, 0xf8, 0x11, 0x13, 0xcb, 0xf5, 0xa8, 0x10, 0xcc, 0xdc, 0x82, 0xa6, 0x85,
	0x69, 0x31, 0x41, 0xdb, 0x56, 0x5a, 0xd4, 0x30, 0x55, 0xbb, 0xc5, 0x69, 0xfe, 0xed, 0x27, 0x39,
	0xd2, 0x71, 0x4e, 0x7b, 0x4b, 0xc3, 0x1c, 0xe9, 0x48, 0x4f, 0xd9, 0xee, 0x12, 0x6f, 0x2a, 0x15,
	0xbb, 0xa7, 0x56, 0x86, 0xb6, 0xe9, 0x46, 0x27, 0x7b, 0xdf, 0xa2, 0xf7, 0x71, 0xf2, 0x8e, 0x13,
	0x35, 0x25, 0x88, 0xfa, 0xcc, 0xbe, 0xc0, 0x1f, 0xae, 0x67, 0x1f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x76, 0x8a, 0xc4, 0xe3, 0xd6, 0x06, 0x00, 0x00,
}
