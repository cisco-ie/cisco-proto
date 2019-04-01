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
// source: pm_ospf_v3proto_bag.proto

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_periodic_ospf_ospfv3_protocol_instances_ospfv3_protocol_instance_samples_sample

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

type PmOspfV3ProtoBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	SampleId             int32    `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmOspfV3ProtoBag_KEYS) Reset()         { *m = PmOspfV3ProtoBag_KEYS{} }
func (m *PmOspfV3ProtoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmOspfV3ProtoBag_KEYS) ProtoMessage()    {}
func (*PmOspfV3ProtoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c6480a8399a46b, []int{0}
}

func (m *PmOspfV3ProtoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmOspfV3ProtoBag_KEYS.Unmarshal(m, b)
}
func (m *PmOspfV3ProtoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmOspfV3ProtoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PmOspfV3ProtoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmOspfV3ProtoBag_KEYS.Merge(m, src)
}
func (m *PmOspfV3ProtoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmOspfV3ProtoBag_KEYS.Size(m)
}
func (m *PmOspfV3ProtoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmOspfV3ProtoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmOspfV3ProtoBag_KEYS proto.InternalMessageInfo

func (m *PmOspfV3ProtoBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *PmOspfV3ProtoBag_KEYS) GetSampleId() int32 {
	if m != nil {
		return m.SampleId
	}
	return 0
}

type PmOspfV3ProtoBag struct {
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmOspfV3ProtoBag) Reset()         { *m = PmOspfV3ProtoBag{} }
func (m *PmOspfV3ProtoBag) String() string { return proto.CompactTextString(m) }
func (*PmOspfV3ProtoBag) ProtoMessage()    {}
func (*PmOspfV3ProtoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c6480a8399a46b, []int{1}
}

func (m *PmOspfV3ProtoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmOspfV3ProtoBag.Unmarshal(m, b)
}
func (m *PmOspfV3ProtoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmOspfV3ProtoBag.Marshal(b, m, deterministic)
}
func (m *PmOspfV3ProtoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmOspfV3ProtoBag.Merge(m, src)
}
func (m *PmOspfV3ProtoBag) XXX_Size() int {
	return xxx_messageInfo_PmOspfV3ProtoBag.Size(m)
}
func (m *PmOspfV3ProtoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmOspfV3ProtoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmOspfV3ProtoBag proto.InternalMessageInfo

func (m *PmOspfV3ProtoBag) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetInputPackets() uint32 {
	if m != nil {
		return m.InputPackets
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetOutputPackets() uint32 {
	if m != nil {
		return m.OutputPackets
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetInputHelloPackets() uint32 {
	if m != nil {
		return m.InputHelloPackets
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetOutputHelloPackets() uint32 {
	if m != nil {
		return m.OutputHelloPackets
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetInputDbDs() uint32 {
	if m != nil {
		return m.InputDbDs
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetInputDbDsLsa() uint32 {
	if m != nil {
		return m.InputDbDsLsa
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetOutputDbDs() uint32 {
	if m != nil {
		return m.OutputDbDs
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetOutputDbDsLsa() uint32 {
	if m != nil {
		return m.OutputDbDsLsa
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetInputLsRequests() uint32 {
	if m != nil {
		return m.InputLsRequests
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetInputLsRequestsLsa() uint32 {
	if m != nil {
		return m.InputLsRequestsLsa
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetOutputLsRequests() uint32 {
	if m != nil {
		return m.OutputLsRequests
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetOutputLsRequestsLsa() uint32 {
	if m != nil {
		return m.OutputLsRequestsLsa
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetInputLsaUpdates() uint32 {
	if m != nil {
		return m.InputLsaUpdates
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetInputLsaUpdatesLsa() uint32 {
	if m != nil {
		return m.InputLsaUpdatesLsa
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetOutputLsaUpdates() uint32 {
	if m != nil {
		return m.OutputLsaUpdates
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetOutputLsaUpdatesLsa() uint32 {
	if m != nil {
		return m.OutputLsaUpdatesLsa
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetInputLsaAcks() uint32 {
	if m != nil {
		return m.InputLsaAcks
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetInputLsaAcksLsa() uint32 {
	if m != nil {
		return m.InputLsaAcksLsa
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetOutputLsaAcks() uint32 {
	if m != nil {
		return m.OutputLsaAcks
	}
	return 0
}

func (m *PmOspfV3ProtoBag) GetOutputLsaAcksLsa() uint32 {
	if m != nil {
		return m.OutputLsaAcksLsa
	}
	return 0
}

func init() {
	proto.RegisterType((*PmOspfV3ProtoBag_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.ospf.ospfv3_protocol_instances.ospfv3_protocol_instance.samples.sample.pm_ospf_v3proto_bag_KEYS")
	proto.RegisterType((*PmOspfV3ProtoBag)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.ospf.ospfv3_protocol_instances.ospfv3_protocol_instance.samples.sample.pm_ospf_v3proto_bag")
}

func init() { proto.RegisterFile("pm_ospf_v3proto_bag.proto", fileDescriptor_12c6480a8399a46b) }

var fileDescriptor_12c6480a8399a46b = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x59, 0x6f, 0xd3, 0x40,
	0x10, 0xc7, 0x65, 0xc4, 0x95, 0xa1, 0x69, 0xc1, 0x01, 0x64, 0x84, 0x40, 0x51, 0xb9, 0x22, 0x8e,
	0x08, 0x08, 0xf7, 0x5d, 0x48, 0x11, 0x88, 0x08, 0xa1, 0x54, 0x3c, 0x20, 0x21, 0x8d, 0x36, 0xf6,
	0x36, 0x58, 0xb1, 0xb3, 0x4b, 0x66, 0x53, 0xc1, 0x2b, 0x1f, 0x85, 0x4f, 0x8a, 0x76, 0xd6, 0x4e,
	0x77, 0xdd, 0xf2, 0x92, 0xd8, 0xb3, 0xbf, 0xf9, 0xcd, 0xfc, 0xb3, 0x0a, 0x5c, 0xd0, 0x25, 0x2a,
	0xd2, 0xbb, 0xb8, 0x37, 0xd0, 0x0b, 0x65, 0x14, 0x4e, 0xc4, 0xb4, 0xcf, 0x4f, 0xf1, 0x9f, 0x28,
	0xcd, 0x29, 0x55, 0x98, 0x2b, 0xc2, 0x5f, 0x0b, 0x2c, 0xc5, 0x5c, 0x4c, 0xa5, 0x98, 0xe4, 0x45,
	0x6e, 0x7e, 0xa3, 0x96, 0x8b, 0xdd, 0x72, 0x5a, 0x1a, 0x54, 0x5a, 0x2e, 0xfa, 0xf6, 0x0d, 0xed,
	0xab, 0x7d, 0xca, 0x55, 0x96, 0xa7, 0x7d, 0x2b, 0xe5, 0x8f, 0xbd, 0x01, 0xb2, 0x2f, 0x55, 0x05,
	0xe6, 0x73, 0x32, 0x62, 0x9e, 0x4a, 0xfa, 0xef, 0x49, 0x9f, 0x44, 0xa9, 0x0b, 0x49, 0xd5, 0xf7,
	0xe6, 0x77, 0x48, 0x0e, 0xd9, 0x10, 0x3f, 0x6d, 0x7f, 0xdb, 0x89, 0xaf, 0x40, 0xbb, 0x6e, 0xc3,
	0xb9, 0x28, 0x65, 0x12, 0x75, 0xa3, 0x5e, 0x6b, 0xbc, 0x56, 0x17, 0x3f, 0x8b, 0x52, 0xc6, 0x17,
	0xa1, 0xe5, 0x54, 0x98, 0x67, 0xc9, 0x91, 0x6e, 0xd4, 0x3b, 0x36, 0x3e, 0xe9, 0x0a, 0x1f, 0xb3,
	0xcd, 0xbf, 0x27, 0xa0, 0x73, 0x88, 0x3e, 0xbe, 0x04, 0x60, 0xf2, 0x52, 0x22, 0x19, 0x51, 0xea,
	0xe4, 0x7e, 0x37, 0xea, 0x1d, 0x1d, 0xb7, 0x6c, 0x65, 0xc7, 0x16, 0xdc, 0x60, 0xbd, 0x34, 0xa8,
	0x45, 0x3a, 0x93, 0x86, 0x92, 0x41, 0x37, 0xea, 0xb5, 0xed, 0x60, 0xbd, 0x34, 0x5f, 0x5c, 0x2d,
	0xbe, 0x06, 0xeb, 0x6a, 0x69, 0x7c, 0xea, 0x01, 0x53, 0x6d, 0x57, 0xad, 0xb1, 0x3e, 0x74, 0x9c,
	0xeb, 0x87, 0x2c, 0x0a, 0xb5, 0x62, 0x1f, 0x32, 0x7b, 0x86, 0x8f, 0x3e, 0xd8, 0x93, 0x9a, 0xbf,
	0x0b, 0x67, 0x2b, 0x6d, 0xd8, 0xf0, 0x88, 0x1b, 0x62, 0x77, 0x16, 0x74, 0x5c, 0x86, 0x53, 0x6e,
	0x42, 0x36, 0xc1, 0x8c, 0x92, 0xc7, 0x0c, 0xb6, 0xb8, 0x34, 0x9c, 0x0c, 0xed, 0xa2, 0x1b, 0xde,
	0x39, 0x16, 0x24, 0x92, 0x27, 0x5e, 0x1e, 0xcb, 0x8c, 0x48, 0xc4, 0x5d, 0x58, 0xab, 0x06, 0x3b,
	0xcf, 0x53, 0x66, 0xc0, 0xd5, 0x58, 0x74, 0x03, 0x4e, 0xfb, 0x04, 0x9b, 0x9e, 0xf9, 0x99, 0x6b,
	0xd5, 0x4d, 0x70, 0xc1, 0xb0, 0x20, 0x5c, 0xc8, 0x9f, 0x4b, 0x49, 0x86, 0x92, 0xe7, 0x4c, 0xba,
	0x55, 0x46, 0x34, 0xae, 0xca, 0xf1, 0x3d, 0x38, 0x77, 0x80, 0x65, 0xf3, 0x0b, 0x17, 0xb8, 0xc1,
	0x5b, 0xfd, 0x6d, 0xa8, 0x7e, 0x86, 0xc0, 0xff, 0x92, 0xf9, 0x6a, 0x43, 0x6f, 0xc0, 0x00, 0xce,
	0x1f, 0xa4, 0x79, 0xc2, 0x2b, 0xee, 0xe8, 0x34, 0x3b, 0x1a, 0x09, 0x04, 0x2e, 0x75, 0x26, 0x8c,
	0xa4, 0xe4, 0x75, 0x90, 0x40, 0x7c, 0x75, 0x65, 0x3f, 0xc1, 0x8a, 0x65, 0xff, 0x9b, 0x20, 0x41,
	0xcd, 0x37, 0x13, 0xec, 0xfb, 0xb7, 0xc2, 0x04, 0xab, 0x01, 0x7e, 0x82, 0x70, 0xc2, 0xdb, 0x30,
	0x81, 0x3f, 0xe2, 0x2a, 0xac, 0xef, 0x6f, 0x25, 0xd2, 0x19, 0x25, 0xef, 0xbc, 0x4b, 0x1f, 0x91,
	0xd8, 0x4a, 0x67, 0x14, 0xdf, 0x82, 0x38, 0xa4, 0x58, 0x3b, 0x0c, 0x83, 0x5a, 0xd2, 0x2a, 0xaf,
	0xc3, 0x86, 0xb7, 0x07, 0x3b, 0xb7, 0xfd, 0xeb, 0xaf, 0xa5, 0x77, 0xa0, 0xd3, 0xe0, 0xd8, 0xfa,
	0xbe, 0x11, 0xaf, 0xd2, 0x4e, 0x8e, 0xf3, 0xff, 0x72, 0xf0, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x1f,
	0x62, 0x59, 0x75, 0xab, 0x04, 0x00, 0x00,
}
