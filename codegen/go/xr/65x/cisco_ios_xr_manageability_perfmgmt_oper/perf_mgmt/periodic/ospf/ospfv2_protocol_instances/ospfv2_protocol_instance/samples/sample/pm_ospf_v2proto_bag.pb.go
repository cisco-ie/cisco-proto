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
// source: pm_ospf_v2proto_bag.proto

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_periodic_ospf_ospfv2_protocol_instances_ospfv2_protocol_instance_samples_sample

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

type PmOspfV2ProtoBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	SampleId             uint32   `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmOspfV2ProtoBag_KEYS) Reset()         { *m = PmOspfV2ProtoBag_KEYS{} }
func (m *PmOspfV2ProtoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmOspfV2ProtoBag_KEYS) ProtoMessage()    {}
func (*PmOspfV2ProtoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_034184cf9055ec04, []int{0}
}

func (m *PmOspfV2ProtoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmOspfV2ProtoBag_KEYS.Unmarshal(m, b)
}
func (m *PmOspfV2ProtoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmOspfV2ProtoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PmOspfV2ProtoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmOspfV2ProtoBag_KEYS.Merge(m, src)
}
func (m *PmOspfV2ProtoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmOspfV2ProtoBag_KEYS.Size(m)
}
func (m *PmOspfV2ProtoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmOspfV2ProtoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmOspfV2ProtoBag_KEYS proto.InternalMessageInfo

func (m *PmOspfV2ProtoBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *PmOspfV2ProtoBag_KEYS) GetSampleId() uint32 {
	if m != nil {
		return m.SampleId
	}
	return 0
}

type PmOspfV2ProtoBag struct {
	TimeStamp            uint64   `protobuf:"varint,50,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	InputPackets         uint32   `protobuf:"varint,51,opt,name=input_packets,json=inputPackets,proto3" json:"input_packets,omitempty"`
	OutputPackets        uint32   `protobuf:"varint,52,opt,name=output_packets,json=outputPackets,proto3" json:"output_packets,omitempty"`
	InputHelloPackets    uint32   `protobuf:"varint,53,opt,name=input_hello_packets,json=inputHelloPackets,proto3" json:"input_hello_packets,omitempty"`
	OutputHelloPackets   uint32   `protobuf:"varint,54,opt,name=output_hello_packets,json=outputHelloPackets,proto3" json:"output_hello_packets,omitempty"`
	InputDbDs            uint32   `protobuf:"varint,55,opt,name=input_db_ds,json=inputDbDs,proto3" json:"input_db_ds,omitempty"`
	InputDbDsLsa         uint32   `protobuf:"varint,56,opt,name=input_db_ds_lsa,json=inputDbDsLsa,proto3" json:"input_db_ds_lsa,omitempty"`
	OutputDbDs           uint32   `protobuf:"varint,57,opt,name=output_db_ds,json=outputDbDs,proto3" json:"output_db_ds,omitempty"`
	OutputDbDsLsa        uint32   `protobuf:"varint,58,opt,name=output_db_ds_lsa,json=outputDbDsLsa,proto3" json:"output_db_ds_lsa,omitempty"`
	InputLsRequests      uint32   `protobuf:"varint,59,opt,name=input_ls_requests,json=inputLsRequests,proto3" json:"input_ls_requests,omitempty"`
	InputLsRequestsLsa   uint32   `protobuf:"varint,60,opt,name=input_ls_requests_lsa,json=inputLsRequestsLsa,proto3" json:"input_ls_requests_lsa,omitempty"`
	OutputLsRequests     uint32   `protobuf:"varint,61,opt,name=output_ls_requests,json=outputLsRequests,proto3" json:"output_ls_requests,omitempty"`
	OutputLsRequestsLsa  uint32   `protobuf:"varint,62,opt,name=output_ls_requests_lsa,json=outputLsRequestsLsa,proto3" json:"output_ls_requests_lsa,omitempty"`
	InputLsaUpdates      uint32   `protobuf:"varint,63,opt,name=input_lsa_updates,json=inputLsaUpdates,proto3" json:"input_lsa_updates,omitempty"`
	InputLsaUpdatesLsa   uint32   `protobuf:"varint,64,opt,name=input_lsa_updates_lsa,json=inputLsaUpdatesLsa,proto3" json:"input_lsa_updates_lsa,omitempty"`
	OutputLsaUpdates     uint32   `protobuf:"varint,65,opt,name=output_lsa_updates,json=outputLsaUpdates,proto3" json:"output_lsa_updates,omitempty"`
	OutputLsaUpdatesLsa  uint32   `protobuf:"varint,66,opt,name=output_lsa_updates_lsa,json=outputLsaUpdatesLsa,proto3" json:"output_lsa_updates_lsa,omitempty"`
	InputLsaAcks         uint32   `protobuf:"varint,67,opt,name=input_lsa_acks,json=inputLsaAcks,proto3" json:"input_lsa_acks,omitempty"`
	InputLsaAcksLsa      uint32   `protobuf:"varint,68,opt,name=input_lsa_acks_lsa,json=inputLsaAcksLsa,proto3" json:"input_lsa_acks_lsa,omitempty"`
	OutputLsaAcks        uint32   `protobuf:"varint,69,opt,name=output_lsa_acks,json=outputLsaAcks,proto3" json:"output_lsa_acks,omitempty"`
	OutputLsaAcksLsa     uint32   `protobuf:"varint,70,opt,name=output_lsa_acks_lsa,json=outputLsaAcksLsa,proto3" json:"output_lsa_acks_lsa,omitempty"`
	ChecksumErrors       uint32   `protobuf:"varint,71,opt,name=checksum_errors,json=checksumErrors,proto3" json:"checksum_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmOspfV2ProtoBag) Reset()         { *m = PmOspfV2ProtoBag{} }
func (m *PmOspfV2ProtoBag) String() string { return proto.CompactTextString(m) }
func (*PmOspfV2ProtoBag) ProtoMessage()    {}
func (*PmOspfV2ProtoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_034184cf9055ec04, []int{1}
}

func (m *PmOspfV2ProtoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmOspfV2ProtoBag.Unmarshal(m, b)
}
func (m *PmOspfV2ProtoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmOspfV2ProtoBag.Marshal(b, m, deterministic)
}
func (m *PmOspfV2ProtoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmOspfV2ProtoBag.Merge(m, src)
}
func (m *PmOspfV2ProtoBag) XXX_Size() int {
	return xxx_messageInfo_PmOspfV2ProtoBag.Size(m)
}
func (m *PmOspfV2ProtoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmOspfV2ProtoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmOspfV2ProtoBag proto.InternalMessageInfo

func (m *PmOspfV2ProtoBag) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetInputPackets() uint32 {
	if m != nil {
		return m.InputPackets
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetOutputPackets() uint32 {
	if m != nil {
		return m.OutputPackets
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetInputHelloPackets() uint32 {
	if m != nil {
		return m.InputHelloPackets
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetOutputHelloPackets() uint32 {
	if m != nil {
		return m.OutputHelloPackets
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetInputDbDs() uint32 {
	if m != nil {
		return m.InputDbDs
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetInputDbDsLsa() uint32 {
	if m != nil {
		return m.InputDbDsLsa
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetOutputDbDs() uint32 {
	if m != nil {
		return m.OutputDbDs
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetOutputDbDsLsa() uint32 {
	if m != nil {
		return m.OutputDbDsLsa
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetInputLsRequests() uint32 {
	if m != nil {
		return m.InputLsRequests
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetInputLsRequestsLsa() uint32 {
	if m != nil {
		return m.InputLsRequestsLsa
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetOutputLsRequests() uint32 {
	if m != nil {
		return m.OutputLsRequests
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetOutputLsRequestsLsa() uint32 {
	if m != nil {
		return m.OutputLsRequestsLsa
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetInputLsaUpdates() uint32 {
	if m != nil {
		return m.InputLsaUpdates
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetInputLsaUpdatesLsa() uint32 {
	if m != nil {
		return m.InputLsaUpdatesLsa
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetOutputLsaUpdates() uint32 {
	if m != nil {
		return m.OutputLsaUpdates
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetOutputLsaUpdatesLsa() uint32 {
	if m != nil {
		return m.OutputLsaUpdatesLsa
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetInputLsaAcks() uint32 {
	if m != nil {
		return m.InputLsaAcks
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetInputLsaAcksLsa() uint32 {
	if m != nil {
		return m.InputLsaAcksLsa
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetOutputLsaAcks() uint32 {
	if m != nil {
		return m.OutputLsaAcks
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetOutputLsaAcksLsa() uint32 {
	if m != nil {
		return m.OutputLsaAcksLsa
	}
	return 0
}

func (m *PmOspfV2ProtoBag) GetChecksumErrors() uint32 {
	if m != nil {
		return m.ChecksumErrors
	}
	return 0
}

func init() {
	proto.RegisterType((*PmOspfV2ProtoBag_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.ospf.ospfv2_protocol_instances.ospfv2_protocol_instance.samples.sample.pm_ospf_v2proto_bag_KEYS")
	proto.RegisterType((*PmOspfV2ProtoBag)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.ospf.ospfv2_protocol_instances.ospfv2_protocol_instance.samples.sample.pm_ospf_v2proto_bag")
}

func init() { proto.RegisterFile("pm_ospf_v2proto_bag.proto", fileDescriptor_034184cf9055ec04) }

var fileDescriptor_034184cf9055ec04 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5b, 0x6f, 0xd3, 0x4c,
	0x10, 0x86, 0x95, 0x4f, 0x9f, 0x80, 0x0c, 0x6d, 0x02, 0x0e, 0xa0, 0x45, 0x08, 0x14, 0x95, 0x43,
	0x23, 0x0e, 0x11, 0x24, 0x9c, 0xcf, 0x85, 0x84, 0x83, 0x88, 0x10, 0x4a, 0xc5, 0x05, 0x12, 0xd2,
	0x68, 0x63, 0x6f, 0x53, 0x2b, 0x76, 0xd6, 0x78, 0x36, 0x15, 0xdc, 0xf2, 0xf7, 0xf8, 0x53, 0x68,
	0x67, 0xed, 0x74, 0xed, 0x96, 0x9b, 0xd6, 0x7e, 0xf7, 0xd9, 0x67, 0xe6, 0xad, 0x55, 0xb8, 0x98,
	0xa5, 0xa8, 0x29, 0xdb, 0xc3, 0x83, 0x41, 0x96, 0x6b, 0xa3, 0x71, 0x26, 0xe7, 0x7d, 0x7e, 0x0a,
	0x7e, 0x37, 0xc2, 0x98, 0x42, 0x8d, 0xb1, 0x26, 0xfc, 0x99, 0x63, 0x2a, 0x97, 0x72, 0xae, 0xe4,
	0x2c, 0x4e, 0x62, 0xf3, 0x0b, 0x33, 0x95, 0xef, 0xa5, 0xf3, 0xd4, 0xa0, 0xce, 0x54, 0xde, 0xb7,
	0x6f, 0x68, 0x5f, 0xed, 0x53, 0xac, 0xa3, 0x38, 0xec, 0x5b, 0x29, 0xff, 0x38, 0x18, 0x20, 0xfb,
	0x42, 0x9d, 0x60, 0xbc, 0x24, 0x23, 0x97, 0xa1, 0xa2, 0x7f, 0x9e, 0xf4, 0x49, 0xa6, 0x59, 0xa2,
	0xa8, 0xf8, 0xbd, 0xf5, 0x1d, 0xc4, 0x31, 0x1b, 0xe2, 0xa7, 0xf1, 0xb7, 0xdd, 0xe0, 0x2a, 0x6c,
	0x96, 0xd7, 0x70, 0x29, 0x53, 0x25, 0x1a, 0xdd, 0x46, 0xaf, 0x39, 0xdd, 0x28, 0xc3, 0xcf, 0x32,
	0x55, 0xc1, 0x25, 0x68, 0x3a, 0x15, 0xc6, 0x91, 0xf8, 0xaf, 0xdb, 0xe8, 0x6d, 0x4e, 0x4f, 0xb9,
	0xe0, 0x63, 0xb4, 0xf5, 0xe7, 0x24, 0x74, 0x8e, 0xd1, 0x07, 0x97, 0x01, 0x4c, 0x9c, 0x2a, 0x24,
	0x23, 0xd3, 0x4c, 0x0c, 0xba, 0x8d, 0xde, 0xff, 0xd3, 0xa6, 0x4d, 0x76, 0x6d, 0xe0, 0x06, 0x67,
	0x2b, 0x83, 0x99, 0x0c, 0x17, 0xca, 0x90, 0x18, 0xb2, 0x77, 0x83, 0xc3, 0x2f, 0x2e, 0x0b, 0xae,
	0x43, 0x4b, 0xaf, 0x8c, 0x4f, 0xdd, 0x67, 0x6a, 0xd3, 0xa5, 0x25, 0xd6, 0x87, 0x8e, 0x73, 0xed,
	0xab, 0x24, 0xd1, 0x6b, 0xf6, 0x01, 0xb3, 0x67, 0xf9, 0xe8, 0x83, 0x3d, 0x29, 0xf9, 0xbb, 0x70,
	0xae, 0xd0, 0x56, 0x2f, 0x3c, 0xe4, 0x0b, 0x81, 0x3b, 0xab, 0xdc, 0xb8, 0x02, 0xa7, 0xdd, 0x84,
	0x68, 0x86, 0x11, 0x89, 0x47, 0x0c, 0x36, 0x39, 0x1a, 0xcd, 0x46, 0x76, 0xd1, 0xb6, 0x77, 0x8e,
	0x09, 0x49, 0xf1, 0xd8, 0xeb, 0x63, 0x99, 0x09, 0xc9, 0xa0, 0x0b, 0x1b, 0xc5, 0x60, 0xe7, 0x79,
	0xc2, 0x0c, 0xb8, 0x8c, 0x45, 0xdb, 0x70, 0xc6, 0x27, 0xd8, 0xf4, 0xd4, 0xef, 0x5c, 0xaa, 0x6e,
	0x82, 0x2b, 0x86, 0x09, 0x61, 0xae, 0x7e, 0xac, 0x14, 0x19, 0x12, 0xcf, 0x98, 0x74, 0xab, 0x4c,
	0x68, 0x5a, 0xc4, 0xc1, 0x3d, 0x38, 0x7f, 0x84, 0x65, 0xf3, 0x73, 0x57, 0xb8, 0xc6, 0x5b, 0xfd,
	0x6d, 0x28, 0xfe, 0x0c, 0x15, 0xff, 0x0b, 0xe6, 0x8b, 0x0d, 0xbd, 0x01, 0x43, 0xb8, 0x70, 0x94,
	0xe6, 0x09, 0x2f, 0xf9, 0x46, 0xa7, 0x7e, 0xa3, 0xd6, 0x40, 0xe2, 0x2a, 0x8b, 0xa4, 0x51, 0x24,
	0x5e, 0x55, 0x1a, 0xc8, 0xaf, 0x2e, 0xf6, 0x1b, 0xac, 0x59, 0xf6, 0xbf, 0xae, 0x34, 0x28, 0xf9,
	0x7a, 0x83, 0x43, 0xff, 0x4e, 0xb5, 0xc1, 0x7a, 0x80, 0xdf, 0xa0, 0x3a, 0xe1, 0x4d, 0xb5, 0x81,
	0x3f, 0xe2, 0x1a, 0xb4, 0x0e, 0xb7, 0x92, 0xe1, 0x82, 0xc4, 0x5b, 0xef, 0xa3, 0x4f, 0x48, 0xee,
	0x84, 0x0b, 0x0a, 0x6e, 0x41, 0x50, 0xa5, 0x58, 0x3b, 0xaa, 0x16, 0xb5, 0xa4, 0x55, 0xde, 0x80,
	0xb6, 0xb7, 0x07, 0x3b, 0xc7, 0xfe, 0xe7, 0x2f, 0xa5, 0x77, 0xa0, 0x53, 0xe3, 0xd8, 0xfa, 0xae,
	0x56, 0xaf, 0xd4, 0x6e, 0x43, 0x3b, 0xdc, 0x57, 0xe1, 0x82, 0x56, 0x29, 0xaa, 0x3c, 0xd7, 0x39,
	0x89, 0xf7, 0x8c, 0xb6, 0xca, 0x78, 0xcc, 0xe9, 0xec, 0x04, 0xff, 0x03, 0x0f, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x2d, 0x8f, 0x38, 0x12, 0xd4, 0x04, 0x00, 0x00,
}
