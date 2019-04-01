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
// source: icpe_install_op_status_detail.proto

package cisco_ios_xr_icpe_infra_oper_nv_satellite_install_op_statuses_install_op_status

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

type IcpeInstallOpStatusDetail_KEYS struct {
	OperationId          uint32   `protobuf:"varint,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IcpeInstallOpStatusDetail_KEYS) Reset()         { *m = IcpeInstallOpStatusDetail_KEYS{} }
func (m *IcpeInstallOpStatusDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*IcpeInstallOpStatusDetail_KEYS) ProtoMessage()    {}
func (*IcpeInstallOpStatusDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_443865083f96737b, []int{0}
}

func (m *IcpeInstallOpStatusDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IcpeInstallOpStatusDetail_KEYS.Unmarshal(m, b)
}
func (m *IcpeInstallOpStatusDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IcpeInstallOpStatusDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *IcpeInstallOpStatusDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcpeInstallOpStatusDetail_KEYS.Merge(m, src)
}
func (m *IcpeInstallOpStatusDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_IcpeInstallOpStatusDetail_KEYS.Size(m)
}
func (m *IcpeInstallOpStatusDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IcpeInstallOpStatusDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IcpeInstallOpStatusDetail_KEYS proto.InternalMessageInfo

func (m *IcpeInstallOpStatusDetail_KEYS) GetOperationId() uint32 {
	if m != nil {
		return m.OperationId
	}
	return 0
}

type IcpeInstallOpStatusDetail struct {
	OperationIdXr         uint32   `protobuf:"varint,50,opt,name=operation_id_xr,json=operationIdXr,proto3" json:"operation_id_xr,omitempty"`
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
	SatelliteRange        string   `protobuf:"bytes,69,opt,name=satellite_range,json=satelliteRange,proto3" json:"satellite_range,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *IcpeInstallOpStatusDetail) Reset()         { *m = IcpeInstallOpStatusDetail{} }
func (m *IcpeInstallOpStatusDetail) String() string { return proto.CompactTextString(m) }
func (*IcpeInstallOpStatusDetail) ProtoMessage()    {}
func (*IcpeInstallOpStatusDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_443865083f96737b, []int{1}
}

func (m *IcpeInstallOpStatusDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IcpeInstallOpStatusDetail.Unmarshal(m, b)
}
func (m *IcpeInstallOpStatusDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IcpeInstallOpStatusDetail.Marshal(b, m, deterministic)
}
func (m *IcpeInstallOpStatusDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcpeInstallOpStatusDetail.Merge(m, src)
}
func (m *IcpeInstallOpStatusDetail) XXX_Size() int {
	return xxx_messageInfo_IcpeInstallOpStatusDetail.Size(m)
}
func (m *IcpeInstallOpStatusDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_IcpeInstallOpStatusDetail.DiscardUnknown(m)
}

var xxx_messageInfo_IcpeInstallOpStatusDetail proto.InternalMessageInfo

func (m *IcpeInstallOpStatusDetail) GetOperationIdXr() uint32 {
	if m != nil {
		return m.OperationIdXr
	}
	return 0
}

func (m *IcpeInstallOpStatusDetail) GetSatsNotInitiated() []uint32 {
	if m != nil {
		return m.SatsNotInitiated
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsTransferring() []uint32 {
	if m != nil {
		return m.SatsTransferring
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsActivating() []uint32 {
	if m != nil {
		return m.SatsActivating
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsUpdating() []uint32 {
	if m != nil {
		return m.SatsUpdating
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsDeactivating() []uint32 {
	if m != nil {
		return m.SatsDeactivating
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsRemoving() []uint32 {
	if m != nil {
		return m.SatsRemoving
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsTransferFailed() []uint32 {
	if m != nil {
		return m.SatsTransferFailed
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsActivateFailed() []uint32 {
	if m != nil {
		return m.SatsActivateFailed
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsUpdateFailed() []uint32 {
	if m != nil {
		return m.SatsUpdateFailed
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsDeactivateFailed() []uint32 {
	if m != nil {
		return m.SatsDeactivateFailed
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsRemoveFailed() []uint32 {
	if m != nil {
		return m.SatsRemoveFailed
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsTransferAborted() []uint32 {
	if m != nil {
		return m.SatsTransferAborted
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsActivateAborted() []uint32 {
	if m != nil {
		return m.SatsActivateAborted
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsUpdateAborted() []uint32 {
	if m != nil {
		return m.SatsUpdateAborted
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsDeactivateAborted() []uint32 {
	if m != nil {
		return m.SatsDeactivateAborted
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsRemoveAborted() []uint32 {
	if m != nil {
		return m.SatsRemoveAborted
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsNoOperation() []uint32 {
	if m != nil {
		return m.SatsNoOperation
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatsCompleted() []uint32 {
	if m != nil {
		return m.SatsCompleted
	}
	return nil
}

func (m *IcpeInstallOpStatusDetail) GetSatelliteRange() string {
	if m != nil {
		return m.SatelliteRange
	}
	return ""
}

func init() {
	proto.RegisterType((*IcpeInstallOpStatusDetail_KEYS)(nil), "cisco_ios_xr_icpe_infra_oper.nv_satellite.install_op_statuses.install_op_status.icpe_install_op_status_detail_KEYS")
	proto.RegisterType((*IcpeInstallOpStatusDetail)(nil), "cisco_ios_xr_icpe_infra_oper.nv_satellite.install_op_statuses.install_op_status.icpe_install_op_status_detail")
}

func init() {
	proto.RegisterFile("icpe_install_op_status_detail.proto", fileDescriptor_443865083f96737b)
}

var fileDescriptor_443865083f96737b = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xed, 0x6b, 0x53, 0x31,
	0x14, 0xc6, 0x29, 0x82, 0x62, 0x5c, 0xad, 0xcb, 0x36, 0xbd, 0x5f, 0x84, 0xda, 0xa1, 0x16, 0x95,
	0x22, 0xdb, 0x9c, 0xef, 0x2f, 0x75, 0x9b, 0x32, 0x04, 0x07, 0x55, 0x41, 0x3f, 0x1d, 0xb2, 0x7b,
	0xd3, 0x11, 0xb8, 0x4b, 0x2e, 0x49, 0x56, 0xf6, 0x87, 0xf9, 0x07, 0x8e, 0x9c, 0x9b, 0xb7, 0xdb,
	0x41, 0xbf, 0x3e, 0x79, 0x9e, 0xf3, 0xe4, 0xd7, 0x9e, 0x1b, 0xb2, 0x2d, 0xca, 0x86, 0x83, 0x90,
	0xc6, 0xb2, 0xba, 0x06, 0xd5, 0x80, 0xb1, 0xcc, 0x5e, 0x18, 0xa8, 0xb8, 0x65, 0xa2, 0x9e, 0x34,
	0x5a, 0x59, 0x45, 0x4f, 0x4a, 0x61, 0x4a, 0x05, 0x42, 0x19, 0xb8, 0xd4, 0xe0, 0x13, 0x73, 0xcd,
	0x40, 0x35, 0x5c, 0x4f, 0xe4, 0x02, 0x0c, 0xb3, 0xbc, 0xae, 0x85, 0xe5, 0x93, 0x6b, 0x93, 0xb8,
	0xb9, 0xae, 0x8d, 0xbe, 0x93, 0xd1, 0xca, 0x5e, 0xf8, 0x71, 0xf4, 0xef, 0x17, 0x7d, 0x44, 0xd6,
	0x5c, 0x01, 0xb3, 0x42, 0x49, 0x10, 0x55, 0xd1, 0x1b, 0xf6, 0xc6, 0xfd, 0xd9, 0x9d, 0xa8, 0x1d,
	0x57, 0xa3, 0xff, 0xb7, 0xc8, 0xc3, 0x95, 0x93, 0xe8, 0x13, 0x32, 0xc8, 0x87, 0xc0, 0xa5, 0x2e,
	0x76, 0x70, 0x4e, 0x3f, 0x9b, 0xf3, 0x57, 0xd3, 0x17, 0x84, 0x1a, 0x66, 0x0d, 0x48, 0x65, 0x41,
	0x48, 0x61, 0x05, 0xb3, 0xbc, 0x2a, 0x76, 0x87, 0x37, 0xc6, 0xfd, 0xd9, 0x3d, 0x77, 0xf2, 0x53,
	0xd9, 0xe3, 0xa0, 0xd3, 0xe7, 0x64, 0x1d, 0xdd, 0x56, 0x33, 0x69, 0xe6, 0x5c, 0x6b, 0x21, 0xcf,
	0x8a, 0xbd, 0x64, 0xfe, 0x9d, 0xe9, 0xf4, 0x29, 0x19, 0xa0, 0x99, 0x95, 0x56, 0x2c, 0x98, 0x75,
	0xd6, 0x57, 0x68, 0xbd, 0xeb, 0xe4, 0x69, 0x54, 0xe9, 0x36, 0xe9, 0xa3, 0xf1, 0xa2, 0xa9, 0x5a,
	0xdb, 0x3e, 0xda, 0xd6, 0x9c, 0xf8, 0xc7, 0x6b, 0xb1, 0xba, 0xe2, 0xd9, 0xbc, 0xd7, 0xa9, 0xfa,
	0x30, 0xd3, 0xe3, 0x44, 0xcd, 0xcf, 0xd5, 0xc2, 0x19, 0xdf, 0xa4, 0x89, 0x33, 0xaf, 0xd1, 0x97,
	0x64, 0xb3, 0x03, 0x03, 0x73, 0x26, 0x6a, 0x5e, 0x15, 0x6f, 0xd1, 0x4b, 0x73, 0x9e, 0x6f, 0x78,
	0x12, 0x13, 0xbe, 0x89, 0x87, 0xc4, 0xbb, 0x94, 0xf0, 0x58, 0xdc, 0x27, 0xc2, 0xcf, 0x8b, 0x68,
	0xd1, 0xff, 0x3e, 0x5d, 0x1b, 0xf9, 0x82, 0x7b, 0x8f, 0xdc, 0xef, 0x32, 0xc6, 0xc4, 0x07, 0x4c,
	0x6c, 0x76, 0x40, 0x97, 0x3b, 0x10, 0x36, 0x26, 0x3e, 0xa6, 0x0e, 0x24, 0x0e, 0xee, 0x1d, 0xb2,
	0xd5, 0xa5, 0x66, 0xa7, 0x4a, 0xbb, 0xff, 0xfc, 0x13, 0x06, 0x36, 0x72, 0xec, 0x69, 0x7b, 0x14,
	0x33, 0xf1, 0x56, 0x21, 0xf3, 0x39, 0x65, 0x02, 0x78, 0xc8, 0x4c, 0xc8, 0x46, 0x4e, 0x1e, 0x12,
	0x5f, 0x30, 0xb1, 0x9e, 0xd0, 0x83, 0x7f, 0x9f, 0x3c, 0x58, 0x66, 0x0f, 0x99, 0x29, 0x66, 0xb6,
	0xba, 0xf0, 0xcb, 0x3d, 0x9e, 0x3e, 0x64, 0xbe, 0xa6, 0x9e, 0x16, 0x3f, 0xf8, 0x9f, 0xf9, 0x3d,
	0x92, 0x0a, 0xe2, 0x97, 0x50, 0x1c, 0xa0, 0x7b, 0xd0, 0xee, 0xfb, 0x49, 0x90, 0xe9, 0x63, 0x82,
	0xab, 0x0a, 0xa5, 0x3a, 0x6f, 0x6a, 0xee, 0xc6, 0x1e, 0xa2, 0x11, 0x97, 0xeb, 0x20, 0x88, 0x7e,
	0xd1, 0xdb, 0x97, 0x00, 0x34, 0x93, 0x67, 0xbc, 0x38, 0x1a, 0xf6, 0xc6, 0xb7, 0x71, 0xd1, 0x5b,
	0x79, 0xe6, 0xd4, 0xd3, 0x9b, 0xf8, 0xae, 0xec, 0x5e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xeb,
	0x18, 0x5f, 0x7e, 0x04, 0x00, 0x00,
}