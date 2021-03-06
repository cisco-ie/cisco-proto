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

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_monitor_ospf_ospfv2_protocol_instances_ospfv2_protocol_instance_samples_sample

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
	SampleId             int32    `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
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

func (m *PmOspfV2ProtoBag_KEYS) GetSampleId() int32 {
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
	proto.RegisterType((*PmOspfV2ProtoBag_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.ospf.ospfv2_protocol_instances.ospfv2_protocol_instance.samples.sample.pm_ospf_v2proto_bag_KEYS")
	proto.RegisterType((*PmOspfV2ProtoBag)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.ospf.ospfv2_protocol_instances.ospfv2_protocol_instance.samples.sample.pm_ospf_v2proto_bag")
}

func init() { proto.RegisterFile("pm_ospf_v2proto_bag.proto", fileDescriptor_034184cf9055ec04) }

var fileDescriptor_034184cf9055ec04 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x69, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0xc4, 0x95, 0xa1, 0x6d, 0xc0, 0x01, 0x64, 0x84, 0x40, 0x51, 0x39, 0x1a, 0x71,
	0x58, 0x90, 0x70, 0xdf, 0x85, 0x84, 0x43, 0x44, 0x08, 0xa5, 0xe2, 0x03, 0x12, 0xd2, 0x68, 0xe3,
	0x6c, 0x53, 0x2b, 0x76, 0xd6, 0x78, 0xd6, 0x15, 0x7c, 0xe4, 0xf7, 0xf1, 0xa7, 0xd0, 0xce, 0xda,
	0xe9, 0xda, 0x2d, 0x5f, 0xda, 0xf8, 0xdd, 0x67, 0x9f, 0x99, 0x37, 0x56, 0xe0, 0x52, 0x96, 0xa2,
	0xa2, 0x6c, 0x17, 0xf7, 0xfb, 0x59, 0xae, 0xb4, 0xc2, 0xa9, 0x98, 0x87, 0xfc, 0xc9, 0xff, 0xe3,
	0x45, 0x31, 0x45, 0x0a, 0x63, 0x45, 0xf8, 0x2b, 0xc7, 0x54, 0x2c, 0xc5, 0x5c, 0x8a, 0x69, 0x9c,
	0xc4, 0xfa, 0x37, 0x66, 0x32, 0xdf, 0x4d, 0xe7, 0xa9, 0x46, 0x95, 0xc9, 0x3c, 0x34, 0x4f, 0x68,
	0x1e, 0xc3, 0x54, 0x2d, 0x63, 0xad, 0xf2, 0xd0, 0x38, 0xf9, 0xcf, 0x7e, 0x1f, 0x59, 0x17, 0xa9,
	0x04, 0xe3, 0x25, 0x69, 0xb1, 0x8c, 0x24, 0xfd, 0xf7, 0x24, 0x24, 0x91, 0x66, 0x89, 0xa4, 0xf2,
	0xff, 0xe6, 0x0f, 0x08, 0x8e, 0x58, 0x10, 0x3f, 0x8f, 0xbe, 0xef, 0xf8, 0xd7, 0x60, 0xbd, 0xba,
	0x86, 0x4b, 0x91, 0xca, 0xc0, 0xeb, 0x7a, 0xbd, 0xd6, 0x64, 0xad, 0x0a, 0xbf, 0x88, 0x54, 0xfa,
	0x97, 0xa1, 0x65, 0x55, 0x18, 0xcf, 0x82, 0x63, 0x5d, 0xaf, 0x77, 0x62, 0x72, 0xda, 0x06, 0x9f,
	0x66, 0x9b, 0x7f, 0x4f, 0x41, 0xe7, 0x08, 0xbd, 0x7f, 0x05, 0x40, 0xc7, 0xa9, 0x44, 0xd2, 0x22,
	0xcd, 0x82, 0x7e, 0xd7, 0xeb, 0x1d, 0x9f, 0xb4, 0x4c, 0xb2, 0x63, 0x02, 0x3b, 0x38, 0x2b, 0x34,
	0x66, 0x22, 0x5a, 0x48, 0x4d, 0xc1, 0xa0, 0xeb, 0xf5, 0xd6, 0xcd, 0xe0, 0xac, 0xd0, 0x5f, 0x6d,
	0xe6, 0xdf, 0x80, 0x0d, 0x55, 0x68, 0x97, 0x7a, 0xc0, 0xd4, 0xba, 0x4d, 0x2b, 0x2c, 0x84, 0x8e,
	0x75, 0xed, 0xc9, 0x24, 0x51, 0x2b, 0xf6, 0x21, 0xb3, 0xe7, 0xf8, 0xe8, 0xa3, 0x39, 0xa9, 0xf8,
	0x7b, 0x70, 0xbe, 0xd4, 0xd6, 0x2f, 0x3c, 0xe2, 0x0b, 0xbe, 0x3d, 0xab, 0xdd, 0xb8, 0x0a, 0x67,
	0xec, 0x84, 0xd9, 0x14, 0x67, 0x14, 0x3c, 0x66, 0xb0, 0xc5, 0xd1, 0x70, 0x3a, 0x34, 0x8b, 0xb6,
	0x9d, 0x73, 0x4c, 0x48, 0x04, 0x4f, 0x9c, 0x3e, 0x86, 0x19, 0x93, 0xf0, 0xbb, 0xb0, 0x56, 0x0e,
	0xb6, 0x9e, 0xa7, 0xcc, 0x80, 0xcd, 0x58, 0xb4, 0x05, 0x67, 0x5d, 0x82, 0x4d, 0xcf, 0xdc, 0xce,
	0x95, 0xea, 0x16, 0xd8, 0x62, 0x98, 0x10, 0xe6, 0xf2, 0x67, 0x21, 0x49, 0x53, 0xf0, 0x9c, 0x49,
	0xbb, 0xca, 0x98, 0x26, 0x65, 0xec, 0xdf, 0x87, 0x0b, 0x87, 0x58, 0x36, 0xbf, 0xb0, 0x85, 0x1b,
	0xbc, 0xd1, 0xdf, 0x81, 0xf2, 0x6b, 0xa8, 0xf9, 0x5f, 0x32, 0x5f, 0x6e, 0xe8, 0x0c, 0x18, 0xc0,
	0xc5, 0xc3, 0x34, 0x4f, 0x78, 0xc5, 0x37, 0x3a, 0xcd, 0x1b, 0x8d, 0x06, 0x02, 0x8b, 0x6c, 0x26,
	0xb4, 0xa4, 0xe0, 0x75, 0xad, 0x81, 0xf8, 0x66, 0x63, 0xb7, 0xc1, 0x8a, 0x65, 0xff, 0x9b, 0x5a,
	0x83, 0x8a, 0x6f, 0x36, 0x38, 0xf0, 0x6f, 0xd7, 0x1b, 0xac, 0x06, 0xb8, 0x0d, 0xea, 0x13, 0xde,
	0xd6, 0x1b, 0xb8, 0x23, 0xae, 0xc3, 0xc6, 0xc1, 0x56, 0x22, 0x5a, 0x50, 0xf0, 0xce, 0x79, 0xe9,
	0x63, 0x12, 0xdb, 0xd1, 0x82, 0xfc, 0xdb, 0xe0, 0xd7, 0x29, 0xd6, 0x0e, 0xeb, 0x45, 0x0d, 0x69,
	0x94, 0x37, 0xa1, 0xed, 0xec, 0xc1, 0xce, 0x91, 0xfb, 0xfa, 0x2b, 0xe9, 0x5d, 0xe8, 0x34, 0x38,
	0xb6, 0xbe, 0x6f, 0xd4, 0xab, 0xb4, 0x5b, 0xd0, 0x8e, 0xf6, 0x64, 0xb4, 0xa0, 0x22, 0x45, 0x99,
	0xe7, 0x2a, 0xa7, 0xe0, 0x03, 0xa3, 0x1b, 0x55, 0x3c, 0xe2, 0x74, 0x7a, 0x92, 0x7f, 0xc0, 0x83,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x68, 0xda, 0xa9, 0xd3, 0x04, 0x00, 0x00,
}
