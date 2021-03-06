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
// source: icpe_install_status_detail.proto

package cisco_ios_xr_icpe_infra_oper_nv_satellite_install_statuses_install_status

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

type IcpeInstallStatusDetail_KEYS struct {
	SatelliteRange       string   `protobuf:"bytes,1,opt,name=satellite_range,json=satelliteRange,proto3" json:"satellite_range,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IcpeInstallStatusDetail_KEYS) Reset()         { *m = IcpeInstallStatusDetail_KEYS{} }
func (m *IcpeInstallStatusDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*IcpeInstallStatusDetail_KEYS) ProtoMessage()    {}
func (*IcpeInstallStatusDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7482721b22fc993b, []int{0}
}

func (m *IcpeInstallStatusDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IcpeInstallStatusDetail_KEYS.Unmarshal(m, b)
}
func (m *IcpeInstallStatusDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IcpeInstallStatusDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *IcpeInstallStatusDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcpeInstallStatusDetail_KEYS.Merge(m, src)
}
func (m *IcpeInstallStatusDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_IcpeInstallStatusDetail_KEYS.Size(m)
}
func (m *IcpeInstallStatusDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IcpeInstallStatusDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IcpeInstallStatusDetail_KEYS proto.InternalMessageInfo

func (m *IcpeInstallStatusDetail_KEYS) GetSatelliteRange() string {
	if m != nil {
		return m.SatelliteRange
	}
	return ""
}

type IcpeInstallStatusDetail struct {
	SatelliteRangeXr      string   `protobuf:"bytes,50,opt,name=satellite_range_xr,json=satelliteRangeXr,proto3" json:"satellite_range_xr,omitempty"`
	SatsNotInitiated      []uint32 `protobuf:"varint,51,rep,packed,name=sats_not_initiated,json=satsNotInitiated,proto3" json:"sats_not_initiated,omitempty"`
	SatsTransferring      []uint32 `protobuf:"varint,52,rep,packed,name=sats_transferring,json=satsTransferring,proto3" json:"sats_transferring,omitempty"`
	SatsActivating        []uint32 `protobuf:"varint,53,rep,packed,name=sats_activating,json=satsActivating,proto3" json:"sats_activating,omitempty"`
	SatsUpdating          []uint32 `protobuf:"varint,54,rep,packed,name=sats_updating,json=satsUpdating,proto3" json:"sats_updating,omitempty"`
	SatsDeactivating      []uint32 `protobuf:"varint,55,rep,packed,name=sats_deactivating,json=satsDeactivating,proto3" json:"sats_deactivating,omitempty"`
	SatsRemoving          []uint32 `protobuf:"varint,56,rep,packed,name=sats_removing,json=satsRemoving,proto3" json:"sats_removing,omitempty"`
	SatsTransferFailed    []uint32 `protobuf:"varint,57,rep,packed,name=sats_transfer_failed,json=satsTransferFailed,proto3" json:"sats_transfer_failed,omitempty"`
	SatsActivateFailed    []uint32 `protobuf:"varint,58,rep,packed,name=sats_activate_failed,json=satsActivateFailed,proto3" json:"sats_activate_failed,omitempty"`
	SatsUpdateFailed      []uint32 `protobuf:"varint,59,rep,packed,name=sats_update_failed,json=satsUpdateFailed,proto3" json:"sats_update_failed,omitempty"`
	SatsDeactivateFailed  []uint32 `protobuf:"varint,60,rep,packed,name=sats_deactivate_failed,json=satsDeactivateFailed,proto3" json:"sats_deactivate_failed,omitempty"`
	SatsRemoveFailed      []uint32 `protobuf:"varint,61,rep,packed,name=sats_remove_failed,json=satsRemoveFailed,proto3" json:"sats_remove_failed,omitempty"`
	SatsTransferAborted   []uint32 `protobuf:"varint,62,rep,packed,name=sats_transfer_aborted,json=satsTransferAborted,proto3" json:"sats_transfer_aborted,omitempty"`
	SatsActivateAborted   []uint32 `protobuf:"varint,63,rep,packed,name=sats_activate_aborted,json=satsActivateAborted,proto3" json:"sats_activate_aborted,omitempty"`
	SatsUpdateAborted     []uint32 `protobuf:"varint,64,rep,packed,name=sats_update_aborted,json=satsUpdateAborted,proto3" json:"sats_update_aborted,omitempty"`
	SatsDeactivateAborted []uint32 `protobuf:"varint,65,rep,packed,name=sats_deactivate_aborted,json=satsDeactivateAborted,proto3" json:"sats_deactivate_aborted,omitempty"`
	SatsRemoveAborted     []uint32 `protobuf:"varint,66,rep,packed,name=sats_remove_aborted,json=satsRemoveAborted,proto3" json:"sats_remove_aborted,omitempty"`
	SatsNoOperation       []uint32 `protobuf:"varint,67,rep,packed,name=sats_no_operation,json=satsNoOperation,proto3" json:"sats_no_operation,omitempty"`
	SatsCompleted         []uint32 `protobuf:"varint,68,rep,packed,name=sats_completed,json=satsCompleted,proto3" json:"sats_completed,omitempty"`
	OperationId           uint32   `protobuf:"varint,69,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *IcpeInstallStatusDetail) Reset()         { *m = IcpeInstallStatusDetail{} }
func (m *IcpeInstallStatusDetail) String() string { return proto.CompactTextString(m) }
func (*IcpeInstallStatusDetail) ProtoMessage()    {}
func (*IcpeInstallStatusDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_7482721b22fc993b, []int{1}
}

func (m *IcpeInstallStatusDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IcpeInstallStatusDetail.Unmarshal(m, b)
}
func (m *IcpeInstallStatusDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IcpeInstallStatusDetail.Marshal(b, m, deterministic)
}
func (m *IcpeInstallStatusDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcpeInstallStatusDetail.Merge(m, src)
}
func (m *IcpeInstallStatusDetail) XXX_Size() int {
	return xxx_messageInfo_IcpeInstallStatusDetail.Size(m)
}
func (m *IcpeInstallStatusDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_IcpeInstallStatusDetail.DiscardUnknown(m)
}

var xxx_messageInfo_IcpeInstallStatusDetail proto.InternalMessageInfo

func (m *IcpeInstallStatusDetail) GetSatelliteRangeXr() string {
	if m != nil {
		return m.SatelliteRangeXr
	}
	return ""
}

func (m *IcpeInstallStatusDetail) GetSatsNotInitiated() []uint32 {
	if m != nil {
		return m.SatsNotInitiated
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsTransferring() []uint32 {
	if m != nil {
		return m.SatsTransferring
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsActivating() []uint32 {
	if m != nil {
		return m.SatsActivating
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsUpdating() []uint32 {
	if m != nil {
		return m.SatsUpdating
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsDeactivating() []uint32 {
	if m != nil {
		return m.SatsDeactivating
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsRemoving() []uint32 {
	if m != nil {
		return m.SatsRemoving
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsTransferFailed() []uint32 {
	if m != nil {
		return m.SatsTransferFailed
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsActivateFailed() []uint32 {
	if m != nil {
		return m.SatsActivateFailed
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsUpdateFailed() []uint32 {
	if m != nil {
		return m.SatsUpdateFailed
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsDeactivateFailed() []uint32 {
	if m != nil {
		return m.SatsDeactivateFailed
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsRemoveFailed() []uint32 {
	if m != nil {
		return m.SatsRemoveFailed
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsTransferAborted() []uint32 {
	if m != nil {
		return m.SatsTransferAborted
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsActivateAborted() []uint32 {
	if m != nil {
		return m.SatsActivateAborted
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsUpdateAborted() []uint32 {
	if m != nil {
		return m.SatsUpdateAborted
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsDeactivateAborted() []uint32 {
	if m != nil {
		return m.SatsDeactivateAborted
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsRemoveAborted() []uint32 {
	if m != nil {
		return m.SatsRemoveAborted
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsNoOperation() []uint32 {
	if m != nil {
		return m.SatsNoOperation
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetSatsCompleted() []uint32 {
	if m != nil {
		return m.SatsCompleted
	}
	return nil
}

func (m *IcpeInstallStatusDetail) GetOperationId() uint32 {
	if m != nil {
		return m.OperationId
	}
	return 0
}

func init() {
	proto.RegisterType((*IcpeInstallStatusDetail_KEYS)(nil), "cisco_ios_xr_icpe_infra_oper.nv_satellite.install_statuses.install_status.icpe_install_status_detail_KEYS")
	proto.RegisterType((*IcpeInstallStatusDetail)(nil), "cisco_ios_xr_icpe_infra_oper.nv_satellite.install_statuses.install_status.icpe_install_status_detail")
}

func init() { proto.RegisterFile("icpe_install_status_detail.proto", fileDescriptor_7482721b22fc993b) }

var fileDescriptor_7482721b22fc993b = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x86, 0x09, 0x82, 0xe2, 0xd8, 0x58, 0x3b, 0x6d, 0x75, 0xf1, 0xc6, 0xb5, 0x22, 0x06, 0x95,
	0x45, 0xda, 0x5a, 0xbf, 0x3f, 0x62, 0x5b, 0x21, 0x0a, 0x0a, 0xab, 0x82, 0x5e, 0x1d, 0xa6, 0xbb,
	0x93, 0x32, 0xb0, 0x99, 0x59, 0x66, 0x26, 0x21, 0x7f, 0xcb, 0x7f, 0x58, 0xf6, 0xec, 0x7c, 0xec,
	0x2e, 0xe4, 0xf6, 0xdd, 0xe7, 0x3d, 0x67, 0x9f, 0xe4, 0x24, 0x24, 0x15, 0x45, 0xcd, 0x41, 0x48,
	0x63, 0x59, 0x55, 0x81, 0xb1, 0xcc, 0x2e, 0x0d, 0x94, 0xdc, 0x32, 0x51, 0x65, 0xb5, 0x56, 0x56,
	0xd1, 0x59, 0x21, 0x4c, 0xa1, 0x40, 0x28, 0x03, 0x6b, 0x0d, 0x0e, 0x9f, 0x6b, 0x06, 0xaa, 0xe6,
	0x3a, 0x93, 0x2b, 0x30, 0xcc, 0xf2, 0xaa, 0x12, 0x96, 0x67, 0xfd, 0x31, 0xdc, 0x0c, 0x82, 0x83,
	0x6f, 0xe4, 0xc1, 0xe6, 0x75, 0xf0, 0xfd, 0xfc, 0xdf, 0x2f, 0xfa, 0x84, 0x6c, 0x87, 0x79, 0xa0,
	0x99, 0xbc, 0xe4, 0xc9, 0x28, 0x1d, 0x4d, 0x6e, 0xe6, 0xb7, 0x43, 0x9c, 0x37, 0xe9, 0xc1, 0xff,
	0x1b, 0xe4, 0xfe, 0xe6, 0x61, 0xf4, 0x39, 0xa1, 0x83, 0x39, 0xb0, 0xd6, 0xc9, 0x21, 0x8e, 0xba,
	0xd3, 0x1f, 0xf5, 0x57, 0x3b, 0xda, 0x80, 0x54, 0x16, 0x84, 0x14, 0x56, 0x30, 0xcb, 0xcb, 0xe4,
	0x28, 0xbd, 0x36, 0x19, 0x23, 0x6d, 0x7e, 0x28, 0x3b, 0xf3, 0x39, 0x7d, 0x46, 0x76, 0x90, 0xb6,
	0x9a, 0x49, 0x33, 0xe7, 0x5a, 0x0b, 0x79, 0x99, 0x1c, 0x47, 0xf8, 0x77, 0x27, 0x77, 0x42, 0x06,
	0x58, 0x61, 0xc5, 0x8a, 0xd9, 0x06, 0x7d, 0x89, 0x68, 0x23, 0x64, 0xa6, 0x21, 0xa5, 0x8f, 0xc8,
	0x18, 0xc1, 0x65, 0x5d, 0xb6, 0xd8, 0x09, 0x62, 0x5b, 0x4d, 0xf8, 0xc7, 0x65, 0x61, 0x75, 0xc9,
	0x3b, 0xf3, 0x5e, 0xc5, 0xd5, 0x67, 0x9d, 0x3c, 0x4c, 0xd4, 0x7c, 0xa1, 0x56, 0x0d, 0xf8, 0x3a,
	0x4e, 0xcc, 0x5d, 0x46, 0x5f, 0x90, 0xbd, 0x9e, 0x0c, 0xcc, 0x99, 0xa8, 0x78, 0x99, 0xbc, 0x41,
	0x96, 0x76, 0x7d, 0xbe, 0xe2, 0x93, 0xd0, 0x70, 0x9b, 0xb8, 0x6f, 0xbc, 0x8d, 0x0d, 0xa7, 0xc5,
	0x5d, 0xc3, 0x7f, 0xbc, 0xa8, 0x16, 0xf8, 0x77, 0xf1, 0xb5, 0xd1, 0xcf, 0xd3, 0xc7, 0xe4, 0x6e,
	0xdf, 0x31, 0x34, 0xde, 0x63, 0x63, 0xaf, 0x27, 0x3a, 0xdc, 0x81, 0xb2, 0xa1, 0xf1, 0x21, 0xee,
	0x40, 0x63, 0x4f, 0x1f, 0x92, 0xfd, 0xbe, 0x35, 0xbb, 0x50, 0xba, 0xf9, 0xce, 0x3f, 0x62, 0x61,
	0xb7, 0xab, 0x3d, 0x6d, 0x1f, 0x85, 0x4e, 0x78, 0x2b, 0xdf, 0xf9, 0x14, 0x3b, 0x5e, 0xdc, 0x77,
	0x32, 0xb2, 0xdb, 0x35, 0xf7, 0x8d, 0xcf, 0xd8, 0xd8, 0x89, 0xea, 0x9e, 0x3f, 0x21, 0xf7, 0x86,
	0xee, 0xbe, 0x33, 0xc5, 0xce, 0x7e, 0x5f, 0x7e, 0xb8, 0xc7, 0xd9, 0xfb, 0xce, 0x97, 0xb8, 0xa7,
	0xd5, 0xf7, 0xfc, 0x53, 0x77, 0x47, 0x52, 0xe1, 0xcf, 0x98, 0x59, 0xa1, 0x64, 0x72, 0x8a, 0xf4,
	0x76, 0x7b, 0xef, 0x3f, 0x7d, 0x4c, 0x1f, 0x13, 0x3c, 0x55, 0x28, 0xd4, 0xa2, 0xae, 0x78, 0x33,
	0xf6, 0x0c, 0x41, 0x3c, 0xae, 0x53, 0x1f, 0xd2, 0x87, 0x64, 0x2b, 0x8c, 0x02, 0x51, 0x26, 0xe7,
	0xe9, 0x68, 0x32, 0xce, 0x6f, 0x85, 0x6c, 0x56, 0x5e, 0x5c, 0xc7, 0x7f, 0x94, 0xa3, 0xab, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x71, 0xce, 0x71, 0x82, 0x75, 0x04, 0x00, 0x00,
}
