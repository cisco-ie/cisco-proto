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
// source: pim_ifstats_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_interface_statistics_interface_statistic

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

type PimIfstatsBag_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimIfstatsBag_KEYS) Reset()         { *m = PimIfstatsBag_KEYS{} }
func (m *PimIfstatsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimIfstatsBag_KEYS) ProtoMessage()    {}
func (*PimIfstatsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_42110ba1d1e0463b, []int{0}
}

func (m *PimIfstatsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimIfstatsBag_KEYS.Unmarshal(m, b)
}
func (m *PimIfstatsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimIfstatsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimIfstatsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimIfstatsBag_KEYS.Merge(m, src)
}
func (m *PimIfstatsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimIfstatsBag_KEYS.Size(m)
}
func (m *PimIfstatsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimIfstatsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimIfstatsBag_KEYS proto.InternalMessageInfo

func (m *PimIfstatsBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type PimIfstatsBag struct {
	InputHello                     uint32   `protobuf:"varint,50,opt,name=input_hello,json=inputHello,proto3" json:"input_hello,omitempty"`
	OutputHello                    uint32   `protobuf:"varint,51,opt,name=output_hello,json=outputHello,proto3" json:"output_hello,omitempty"`
	InputRegister                  uint32   `protobuf:"varint,52,opt,name=input_register,json=inputRegister,proto3" json:"input_register,omitempty"`
	OutputRegister                 uint32   `protobuf:"varint,53,opt,name=output_register,json=outputRegister,proto3" json:"output_register,omitempty"`
	InputRegisterStop              uint32   `protobuf:"varint,54,opt,name=input_register_stop,json=inputRegisterStop,proto3" json:"input_register_stop,omitempty"`
	OutputRegisterStop             uint32   `protobuf:"varint,55,opt,name=output_register_stop,json=outputRegisterStop,proto3" json:"output_register_stop,omitempty"`
	InputJp                        uint32   `protobuf:"varint,56,opt,name=input_jp,json=inputJp,proto3" json:"input_jp,omitempty"`
	OutputJp                       uint32   `protobuf:"varint,57,opt,name=output_jp,json=outputJp,proto3" json:"output_jp,omitempty"`
	InputBsrMessage                uint32   `protobuf:"varint,58,opt,name=input_bsr_message,json=inputBsrMessage,proto3" json:"input_bsr_message,omitempty"`
	OutputBsrMessage               uint32   `protobuf:"varint,59,opt,name=output_bsr_message,json=outputBsrMessage,proto3" json:"output_bsr_message,omitempty"`
	InputAssert                    uint32   `protobuf:"varint,60,opt,name=input_assert,json=inputAssert,proto3" json:"input_assert,omitempty"`
	OutputAssert                   uint32   `protobuf:"varint,61,opt,name=output_assert,json=outputAssert,proto3" json:"output_assert,omitempty"`
	InputGraftMessage              uint32   `protobuf:"varint,62,opt,name=input_graft_message,json=inputGraftMessage,proto3" json:"input_graft_message,omitempty"`
	OutputGraftMessage             uint32   `protobuf:"varint,63,opt,name=output_graft_message,json=outputGraftMessage,proto3" json:"output_graft_message,omitempty"`
	InputGraftAckMessage           uint32   `protobuf:"varint,64,opt,name=input_graft_ack_message,json=inputGraftAckMessage,proto3" json:"input_graft_ack_message,omitempty"`
	OutputGraftAckMessage          uint32   `protobuf:"varint,65,opt,name=output_graft_ack_message,json=outputGraftAckMessage,proto3" json:"output_graft_ack_message,omitempty"`
	InputCandidateRpAdvertisement  uint32   `protobuf:"varint,66,opt,name=input_candidate_rp_advertisement,json=inputCandidateRpAdvertisement,proto3" json:"input_candidate_rp_advertisement,omitempty"`
	OutputCandidateRpAdvertisement uint32   `protobuf:"varint,67,opt,name=output_candidate_rp_advertisement,json=outputCandidateRpAdvertisement,proto3" json:"output_candidate_rp_advertisement,omitempty"`
	InputDfElection                uint32   `protobuf:"varint,68,opt,name=input_df_election,json=inputDfElection,proto3" json:"input_df_election,omitempty"`
	OutputDfElection               uint32   `protobuf:"varint,69,opt,name=output_df_election,json=outputDfElection,proto3" json:"output_df_election,omitempty"`
	InputMiscellaneous             uint32   `protobuf:"varint,70,opt,name=input_miscellaneous,json=inputMiscellaneous,proto3" json:"input_miscellaneous,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *PimIfstatsBag) Reset()         { *m = PimIfstatsBag{} }
func (m *PimIfstatsBag) String() string { return proto.CompactTextString(m) }
func (*PimIfstatsBag) ProtoMessage()    {}
func (*PimIfstatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_42110ba1d1e0463b, []int{1}
}

func (m *PimIfstatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimIfstatsBag.Unmarshal(m, b)
}
func (m *PimIfstatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimIfstatsBag.Marshal(b, m, deterministic)
}
func (m *PimIfstatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimIfstatsBag.Merge(m, src)
}
func (m *PimIfstatsBag) XXX_Size() int {
	return xxx_messageInfo_PimIfstatsBag.Size(m)
}
func (m *PimIfstatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimIfstatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimIfstatsBag proto.InternalMessageInfo

func (m *PimIfstatsBag) GetInputHello() uint32 {
	if m != nil {
		return m.InputHello
	}
	return 0
}

func (m *PimIfstatsBag) GetOutputHello() uint32 {
	if m != nil {
		return m.OutputHello
	}
	return 0
}

func (m *PimIfstatsBag) GetInputRegister() uint32 {
	if m != nil {
		return m.InputRegister
	}
	return 0
}

func (m *PimIfstatsBag) GetOutputRegister() uint32 {
	if m != nil {
		return m.OutputRegister
	}
	return 0
}

func (m *PimIfstatsBag) GetInputRegisterStop() uint32 {
	if m != nil {
		return m.InputRegisterStop
	}
	return 0
}

func (m *PimIfstatsBag) GetOutputRegisterStop() uint32 {
	if m != nil {
		return m.OutputRegisterStop
	}
	return 0
}

func (m *PimIfstatsBag) GetInputJp() uint32 {
	if m != nil {
		return m.InputJp
	}
	return 0
}

func (m *PimIfstatsBag) GetOutputJp() uint32 {
	if m != nil {
		return m.OutputJp
	}
	return 0
}

func (m *PimIfstatsBag) GetInputBsrMessage() uint32 {
	if m != nil {
		return m.InputBsrMessage
	}
	return 0
}

func (m *PimIfstatsBag) GetOutputBsrMessage() uint32 {
	if m != nil {
		return m.OutputBsrMessage
	}
	return 0
}

func (m *PimIfstatsBag) GetInputAssert() uint32 {
	if m != nil {
		return m.InputAssert
	}
	return 0
}

func (m *PimIfstatsBag) GetOutputAssert() uint32 {
	if m != nil {
		return m.OutputAssert
	}
	return 0
}

func (m *PimIfstatsBag) GetInputGraftMessage() uint32 {
	if m != nil {
		return m.InputGraftMessage
	}
	return 0
}

func (m *PimIfstatsBag) GetOutputGraftMessage() uint32 {
	if m != nil {
		return m.OutputGraftMessage
	}
	return 0
}

func (m *PimIfstatsBag) GetInputGraftAckMessage() uint32 {
	if m != nil {
		return m.InputGraftAckMessage
	}
	return 0
}

func (m *PimIfstatsBag) GetOutputGraftAckMessage() uint32 {
	if m != nil {
		return m.OutputGraftAckMessage
	}
	return 0
}

func (m *PimIfstatsBag) GetInputCandidateRpAdvertisement() uint32 {
	if m != nil {
		return m.InputCandidateRpAdvertisement
	}
	return 0
}

func (m *PimIfstatsBag) GetOutputCandidateRpAdvertisement() uint32 {
	if m != nil {
		return m.OutputCandidateRpAdvertisement
	}
	return 0
}

func (m *PimIfstatsBag) GetInputDfElection() uint32 {
	if m != nil {
		return m.InputDfElection
	}
	return 0
}

func (m *PimIfstatsBag) GetOutputDfElection() uint32 {
	if m != nil {
		return m.OutputDfElection
	}
	return 0
}

func (m *PimIfstatsBag) GetInputMiscellaneous() uint32 {
	if m != nil {
		return m.InputMiscellaneous
	}
	return 0
}

func init() {
	proto.RegisterType((*PimIfstatsBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.interface_statistics.interface_statistic.pim_ifstats_bag_KEYS")
	proto.RegisterType((*PimIfstatsBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.interface_statistics.interface_statistic.pim_ifstats_bag")
}

func init() { proto.RegisterFile("pim_ifstats_bag.proto", fileDescriptor_42110ba1d1e0463b) }

var fileDescriptor_42110ba1d1e0463b = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x4f, 0x6f, 0x13, 0x31,
	0x10, 0xc5, 0x95, 0x0b, 0x6d, 0xdd, 0xa6, 0x01, 0x93, 0x0a, 0x23, 0x04, 0x24, 0x45, 0x88, 0x0a,
	0xa1, 0x05, 0xd1, 0x96, 0xf2, 0xaf, 0x40, 0xda, 0x86, 0x42, 0x51, 0x39, 0xa4, 0x27, 0x2e, 0x58,
	0x8e, 0x33, 0x1b, 0xdc, 0xee, 0xae, 0x2d, 0xdb, 0x89, 0xfa, 0x39, 0xf9, 0x44, 0x28, 0xe3, 0xcd,
	0xc6, 0x1b, 0x01, 0xc7, 0xbc, 0xf9, 0xbd, 0xf7, 0xa2, 0xf1, 0x68, 0xc9, 0x96, 0x51, 0x39, 0x57,
	0xa9, 0xf3, 0xc2, 0x3b, 0x3e, 0x14, 0xe3, 0xc4, 0x58, 0xed, 0x35, 0xfd, 0x29, 0x95, 0x93, 0x9a,
	0x2b, 0xed, 0xf8, 0xb5, 0xe5, 0xca, 0x4c, 0xf7, 0xf8, 0x0c, 0xd4, 0x06, 0x6c, 0x62, 0x54, 0x9e,
	0x08, 0xe9, 0xd5, 0x14, 0x92, 0x11, 0xa4, 0x62, 0x92, 0x79, 0x2e, 0x75, 0xe1, 0xe1, 0xda, 0x27,
	0xaa, 0xf0, 0x60, 0x53, 0x21, 0x81, 0xcf, 0x02, 0x95, 0xf3, 0x4a, 0xba, 0xbf, 0x89, 0xdb, 0x87,
	0xa4, 0xbd, 0x54, 0xcc, 0xbf, 0xf5, 0x7f, 0x5c, 0xd0, 0xc7, 0x64, 0x73, 0x81, 0x17, 0x22, 0x07,
	0xd6, 0xe8, 0x34, 0x76, 0xd6, 0x06, 0xcd, 0x4a, 0xfd, 0x2e, 0x72, 0xd8, 0xfe, 0xbd, 0x42, 0x5a,
	0x4b, 0x7e, 0xfa, 0x90, 0xac, 0xab, 0xc2, 0x4c, 0x3c, 0xff, 0x05, 0x59, 0xa6, 0xd9, 0xcb, 0x4e,
	0x63, 0xa7, 0x39, 0x20, 0x28, 0x7d, 0x99, 0x29, 0xb4, 0x4b, 0x36, 0xf4, 0xc4, 0x2f, 0x88, 0x5d,
	0x24, 0xd6, 0x83, 0x16, 0x10, 0xac, 0x9f, 0x11, 0x16, 0xc6, 0xca, 0x79, 0xb0, 0x6c, 0x0f, 0xa1,
	0x26, 0xaa, 0x83, 0x52, 0xa4, 0x4f, 0x48, 0xab, 0x4c, 0xaa, 0xb8, 0x7d, 0xe4, 0x36, 0x83, 0x5c,
	0x81, 0x09, 0xb9, 0x5d, 0xcf, 0xe3, 0xce, 0x6b, 0xc3, 0x5e, 0x21, 0x7c, 0xab, 0x16, 0x7a, 0xe1,
	0xb5, 0xa1, 0x2f, 0x48, 0x7b, 0x29, 0x38, 0x18, 0x0e, 0xd0, 0x40, 0xeb, 0xe9, 0xe8, 0xb8, 0x4b,
	0x56, 0x43, 0xc3, 0xa5, 0x61, 0xaf, 0x91, 0x5a, 0xc1, 0xdf, 0x67, 0x86, 0xde, 0x23, 0x6b, 0x65,
	0xd8, 0xa5, 0x61, 0x6f, 0x70, 0xb6, 0x1a, 0x84, 0x33, 0x43, 0x9f, 0x92, 0x50, 0xcf, 0x87, 0xce,
	0xf2, 0x1c, 0x9c, 0x13, 0x63, 0x60, 0x6f, 0x11, 0x6a, 0xe1, 0xe0, 0xc8, 0xd9, 0xf3, 0x20, 0xd3,
	0x67, 0xa4, 0x6c, 0xae, 0xc1, 0xef, 0x10, 0xbe, 0x19, 0x26, 0x11, 0xdd, 0x25, 0x1b, 0x21, 0x59,
	0x38, 0x07, 0xd6, 0xb3, 0xf7, 0x61, 0xcd, 0xa8, 0xf5, 0x50, 0xa2, 0x8f, 0x48, 0xb3, 0x0c, 0x2c,
	0x99, 0x43, 0x64, 0xca, 0xe7, 0x29, 0xa1, 0x6a, 0x77, 0x63, 0x2b, 0x52, 0x5f, 0xd5, 0x7e, 0x88,
	0x76, 0x77, 0x3a, 0x9b, 0xcc, 0x7b, 0x17, 0xbb, 0xab, 0x1b, 0x3e, 0xc6, 0xbb, 0xab, 0x39, 0xf6,
	0xc9, 0x9d, 0xb8, 0x41, 0xc8, 0xab, 0xca, 0xf4, 0x09, 0x4d, 0xed, 0x45, 0x4b, 0x4f, 0x5e, 0xcd,
	0x6d, 0x07, 0x84, 0xd5, 0x8a, 0x62, 0x5f, 0x0f, 0x7d, 0x5b, 0x51, 0x59, 0x64, 0x3c, 0x25, 0x9d,
	0xd0, 0x27, 0x45, 0x31, 0x52, 0x23, 0xe1, 0x81, 0x5b, 0xc3, 0xc5, 0x68, 0x0a, 0xd6, 0x2b, 0x07,
	0x39, 0x14, 0x9e, 0x1d, 0x61, 0xc0, 0x7d, 0xe4, 0x8e, 0xe7, 0xd8, 0xc0, 0xf4, 0x62, 0x88, 0x7e,
	0x25, 0xdd, 0xf2, 0x1f, 0xfc, 0x27, 0xe9, 0x18, 0x93, 0x1e, 0x04, 0xf0, 0x9f, 0x51, 0xd5, 0x1d,
	0x8c, 0x52, 0x0e, 0x19, 0x48, 0xaf, 0x74, 0xc1, 0x4e, 0xa2, 0x3b, 0x38, 0x49, 0xfb, 0xa5, 0x1c,
	0xdd, 0x41, 0x0c, 0xf7, 0xe3, 0x3b, 0x88, 0xe8, 0xe7, 0xf3, 0xf7, 0xcb, 0x95, 0x93, 0x90, 0x65,
	0xa2, 0x00, 0x3d, 0x71, 0xec, 0x73, 0x78, 0x0e, 0x1c, 0x9d, 0xc7, 0x93, 0xe1, 0x0d, 0xfc, 0xf4,
	0xec, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xf5, 0x5f, 0x92, 0x93, 0x04, 0x00, 0x00,
}
