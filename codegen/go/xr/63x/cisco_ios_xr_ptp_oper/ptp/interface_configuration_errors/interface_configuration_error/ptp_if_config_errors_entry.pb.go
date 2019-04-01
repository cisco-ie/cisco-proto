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
// source: ptp_if_config_errors_entry.proto

package cisco_ios_xr_ptp_oper_ptp_interface_configuration_errors_interface_configuration_error

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

type PtpIfConfigErrorsEntry_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpIfConfigErrorsEntry_KEYS) Reset()         { *m = PtpIfConfigErrorsEntry_KEYS{} }
func (m *PtpIfConfigErrorsEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpIfConfigErrorsEntry_KEYS) ProtoMessage()    {}
func (*PtpIfConfigErrorsEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9d8af8579d2211, []int{0}
}

func (m *PtpIfConfigErrorsEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfConfigErrorsEntry_KEYS.Unmarshal(m, b)
}
func (m *PtpIfConfigErrorsEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfConfigErrorsEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpIfConfigErrorsEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfConfigErrorsEntry_KEYS.Merge(m, src)
}
func (m *PtpIfConfigErrorsEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpIfConfigErrorsEntry_KEYS.Size(m)
}
func (m *PtpIfConfigErrorsEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfConfigErrorsEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfConfigErrorsEntry_KEYS proto.InternalMessageInfo

func (m *PtpIfConfigErrorsEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type PtpBagConfigErrors struct {
	GlobalPtp               bool     `protobuf:"varint,1,opt,name=global_ptp,json=globalPtp,proto3" json:"global_ptp,omitempty"`
	EthernetTransport       bool     `protobuf:"varint,2,opt,name=ethernet_transport,json=ethernetTransport,proto3" json:"ethernet_transport,omitempty"`
	OneStep                 bool     `protobuf:"varint,3,opt,name=one_step,json=oneStep,proto3" json:"one_step,omitempty"`
	Slave                   bool     `protobuf:"varint,4,opt,name=slave,proto3" json:"slave,omitempty"`
	Ipv6                    bool     `protobuf:"varint,5,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	Multicast               bool     `protobuf:"varint,6,opt,name=multicast,proto3" json:"multicast,omitempty"`
	ProfileNotGlobal        bool     `protobuf:"varint,7,opt,name=profile_not_global,json=profileNotGlobal,proto3" json:"profile_not_global,omitempty"`
	LocalPriority           bool     `protobuf:"varint,8,opt,name=local_priority,json=localPriority,proto3" json:"local_priority,omitempty"`
	ProfileEthernet         bool     `protobuf:"varint,9,opt,name=profile_ethernet,json=profileEthernet,proto3" json:"profile_ethernet,omitempty"`
	ProfileIpv4             bool     `protobuf:"varint,10,opt,name=profile_ipv4,json=profileIpv4,proto3" json:"profile_ipv4,omitempty"`
	ProfileIpv6             bool     `protobuf:"varint,11,opt,name=profile_ipv6,json=profileIpv6,proto3" json:"profile_ipv6,omitempty"`
	ProfileUnicast          bool     `protobuf:"varint,12,opt,name=profile_unicast,json=profileUnicast,proto3" json:"profile_unicast,omitempty"`
	ProfileMulticast        bool     `protobuf:"varint,13,opt,name=profile_multicast,json=profileMulticast,proto3" json:"profile_multicast,omitempty"`
	ProfileMixed            bool     `protobuf:"varint,14,opt,name=profile_mixed,json=profileMixed,proto3" json:"profile_mixed,omitempty"`
	ProfileMasterUnicast    bool     `protobuf:"varint,15,opt,name=profile_master_unicast,json=profileMasterUnicast,proto3" json:"profile_master_unicast,omitempty"`
	ProfileMasterMulticast  bool     `protobuf:"varint,16,opt,name=profile_master_multicast,json=profileMasterMulticast,proto3" json:"profile_master_multicast,omitempty"`
	ProfileMasterMixed      bool     `protobuf:"varint,17,opt,name=profile_master_mixed,json=profileMasterMixed,proto3" json:"profile_master_mixed,omitempty"`
	TargetAddressIpv4       bool     `protobuf:"varint,18,opt,name=target_address_ipv4,json=targetAddressIpv4,proto3" json:"target_address_ipv4,omitempty"`
	TargetAddressIpv6       bool     `protobuf:"varint,19,opt,name=target_address_ipv6,json=targetAddressIpv6,proto3" json:"target_address_ipv6,omitempty"`
	ProfilePortState        bool     `protobuf:"varint,20,opt,name=profile_port_state,json=profilePortState,proto3" json:"profile_port_state,omitempty"`
	ProfileAnnounceInterval bool     `protobuf:"varint,21,opt,name=profile_announce_interval,json=profileAnnounceInterval,proto3" json:"profile_announce_interval,omitempty"`
	ProfileSyncInterval     bool     `protobuf:"varint,22,opt,name=profile_sync_interval,json=profileSyncInterval,proto3" json:"profile_sync_interval,omitempty"`
	ProfileDelayReqInterval bool     `protobuf:"varint,23,opt,name=profile_delay_req_interval,json=profileDelayReqInterval,proto3" json:"profile_delay_req_interval,omitempty"`
	ProfileSyncTimeout      bool     `protobuf:"varint,24,opt,name=profile_sync_timeout,json=profileSyncTimeout,proto3" json:"profile_sync_timeout,omitempty"`
	ProfileDelayRespTimeout bool     `protobuf:"varint,25,opt,name=profile_delay_resp_timeout,json=profileDelayRespTimeout,proto3" json:"profile_delay_resp_timeout,omitempty"`
	InvalidGrantReduction   bool     `protobuf:"varint,26,opt,name=invalid_grant_reduction,json=invalidGrantReduction,proto3" json:"invalid_grant_reduction,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *PtpBagConfigErrors) Reset()         { *m = PtpBagConfigErrors{} }
func (m *PtpBagConfigErrors) String() string { return proto.CompactTextString(m) }
func (*PtpBagConfigErrors) ProtoMessage()    {}
func (*PtpBagConfigErrors) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9d8af8579d2211, []int{1}
}

func (m *PtpBagConfigErrors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagConfigErrors.Unmarshal(m, b)
}
func (m *PtpBagConfigErrors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagConfigErrors.Marshal(b, m, deterministic)
}
func (m *PtpBagConfigErrors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagConfigErrors.Merge(m, src)
}
func (m *PtpBagConfigErrors) XXX_Size() int {
	return xxx_messageInfo_PtpBagConfigErrors.Size(m)
}
func (m *PtpBagConfigErrors) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagConfigErrors.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagConfigErrors proto.InternalMessageInfo

func (m *PtpBagConfigErrors) GetGlobalPtp() bool {
	if m != nil {
		return m.GlobalPtp
	}
	return false
}

func (m *PtpBagConfigErrors) GetEthernetTransport() bool {
	if m != nil {
		return m.EthernetTransport
	}
	return false
}

func (m *PtpBagConfigErrors) GetOneStep() bool {
	if m != nil {
		return m.OneStep
	}
	return false
}

func (m *PtpBagConfigErrors) GetSlave() bool {
	if m != nil {
		return m.Slave
	}
	return false
}

func (m *PtpBagConfigErrors) GetIpv6() bool {
	if m != nil {
		return m.Ipv6
	}
	return false
}

func (m *PtpBagConfigErrors) GetMulticast() bool {
	if m != nil {
		return m.Multicast
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileNotGlobal() bool {
	if m != nil {
		return m.ProfileNotGlobal
	}
	return false
}

func (m *PtpBagConfigErrors) GetLocalPriority() bool {
	if m != nil {
		return m.LocalPriority
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileEthernet() bool {
	if m != nil {
		return m.ProfileEthernet
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileIpv4() bool {
	if m != nil {
		return m.ProfileIpv4
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileIpv6() bool {
	if m != nil {
		return m.ProfileIpv6
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileUnicast() bool {
	if m != nil {
		return m.ProfileUnicast
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileMulticast() bool {
	if m != nil {
		return m.ProfileMulticast
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileMixed() bool {
	if m != nil {
		return m.ProfileMixed
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileMasterUnicast() bool {
	if m != nil {
		return m.ProfileMasterUnicast
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileMasterMulticast() bool {
	if m != nil {
		return m.ProfileMasterMulticast
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileMasterMixed() bool {
	if m != nil {
		return m.ProfileMasterMixed
	}
	return false
}

func (m *PtpBagConfigErrors) GetTargetAddressIpv4() bool {
	if m != nil {
		return m.TargetAddressIpv4
	}
	return false
}

func (m *PtpBagConfigErrors) GetTargetAddressIpv6() bool {
	if m != nil {
		return m.TargetAddressIpv6
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfilePortState() bool {
	if m != nil {
		return m.ProfilePortState
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileAnnounceInterval() bool {
	if m != nil {
		return m.ProfileAnnounceInterval
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileSyncInterval() bool {
	if m != nil {
		return m.ProfileSyncInterval
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileDelayReqInterval() bool {
	if m != nil {
		return m.ProfileDelayReqInterval
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileSyncTimeout() bool {
	if m != nil {
		return m.ProfileSyncTimeout
	}
	return false
}

func (m *PtpBagConfigErrors) GetProfileDelayRespTimeout() bool {
	if m != nil {
		return m.ProfileDelayRespTimeout
	}
	return false
}

func (m *PtpBagConfigErrors) GetInvalidGrantReduction() bool {
	if m != nil {
		return m.InvalidGrantReduction
	}
	return false
}

type PtpIfConfigErrorsEntry struct {
	ConfigurationProfileName string              `protobuf:"bytes,50,opt,name=configuration_profile_name,json=configurationProfileName,proto3" json:"configuration_profile_name,omitempty"`
	ClockProfile             string              `protobuf:"bytes,51,opt,name=clock_profile,json=clockProfile,proto3" json:"clock_profile,omitempty"`
	TelecomClockType         string              `protobuf:"bytes,52,opt,name=telecom_clock_type,json=telecomClockType,proto3" json:"telecom_clock_type,omitempty"`
	RestrictPortState        string              `protobuf:"bytes,53,opt,name=restrict_port_state,json=restrictPortState,proto3" json:"restrict_port_state,omitempty"`
	ConfigurationErrors      *PtpBagConfigErrors `protobuf:"bytes,54,opt,name=configuration_errors,json=configurationErrors,proto3" json:"configuration_errors,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}            `json:"-"`
	XXX_unrecognized         []byte              `json:"-"`
	XXX_sizecache            int32               `json:"-"`
}

func (m *PtpIfConfigErrorsEntry) Reset()         { *m = PtpIfConfigErrorsEntry{} }
func (m *PtpIfConfigErrorsEntry) String() string { return proto.CompactTextString(m) }
func (*PtpIfConfigErrorsEntry) ProtoMessage()    {}
func (*PtpIfConfigErrorsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9d8af8579d2211, []int{2}
}

func (m *PtpIfConfigErrorsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfConfigErrorsEntry.Unmarshal(m, b)
}
func (m *PtpIfConfigErrorsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfConfigErrorsEntry.Marshal(b, m, deterministic)
}
func (m *PtpIfConfigErrorsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfConfigErrorsEntry.Merge(m, src)
}
func (m *PtpIfConfigErrorsEntry) XXX_Size() int {
	return xxx_messageInfo_PtpIfConfigErrorsEntry.Size(m)
}
func (m *PtpIfConfigErrorsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfConfigErrorsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfConfigErrorsEntry proto.InternalMessageInfo

func (m *PtpIfConfigErrorsEntry) GetConfigurationProfileName() string {
	if m != nil {
		return m.ConfigurationProfileName
	}
	return ""
}

func (m *PtpIfConfigErrorsEntry) GetClockProfile() string {
	if m != nil {
		return m.ClockProfile
	}
	return ""
}

func (m *PtpIfConfigErrorsEntry) GetTelecomClockType() string {
	if m != nil {
		return m.TelecomClockType
	}
	return ""
}

func (m *PtpIfConfigErrorsEntry) GetRestrictPortState() string {
	if m != nil {
		return m.RestrictPortState
	}
	return ""
}

func (m *PtpIfConfigErrorsEntry) GetConfigurationErrors() *PtpBagConfigErrors {
	if m != nil {
		return m.ConfigurationErrors
	}
	return nil
}

func init() {
	proto.RegisterType((*PtpIfConfigErrorsEntry_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_configuration_errors.interface_configuration_error.ptp_if_config_errors_entry_KEYS")
	proto.RegisterType((*PtpBagConfigErrors)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_configuration_errors.interface_configuration_error.ptp_bag_config_errors")
	proto.RegisterType((*PtpIfConfigErrorsEntry)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_configuration_errors.interface_configuration_error.ptp_if_config_errors_entry")
}

func init() { proto.RegisterFile("ptp_if_config_errors_entry.proto", fileDescriptor_6e9d8af8579d2211) }

var fileDescriptor_6e9d8af8579d2211 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcf, 0x6e, 0x2b, 0x35,
	0x14, 0xc6, 0x15, 0xfa, 0x37, 0x6e, 0xd3, 0x36, 0x4e, 0xd2, 0xba, 0x11, 0x88, 0x52, 0x84, 0x28,
	0x02, 0x22, 0xd4, 0x86, 0x08, 0x01, 0x9b, 0x0a, 0xaa, 0x52, 0xa1, 0x56, 0x51, 0x52, 0x90, 0x58,
	0x59, 0xee, 0xe4, 0x24, 0x58, 0x4c, 0x6c, 0xd7, 0xe3, 0x44, 0xcd, 0x8e, 0x17, 0xe0, 0x1d, 0x78,
	0xd4, 0x2b, 0x9f, 0xb1, 0x27, 0x9d, 0x28, 0xbd, 0xbb, 0xbb, 0x4b, 0xbe, 0xef, 0xf7, 0xd9, 0xc7,
	0x47, 0xf6, 0x19, 0x72, 0x66, 0x9c, 0xe1, 0x72, 0xcc, 0x13, 0xad, 0xc6, 0x72, 0xc2, 0xc1, 0x5a,
	0x6d, 0x33, 0x0e, 0xca, 0xd9, 0x45, 0xc7, 0x58, 0xed, 0x34, 0xfd, 0x33, 0x91, 0x59, 0xa2, 0xb9,
	0xd4, 0x19, 0x7f, 0xb1, 0xdc, 0xe3, 0xda, 0x80, 0xed, 0x18, 0x67, 0x3a, 0x52, 0x39, 0xb0, 0x63,
	0x91, 0x40, 0x88, 0xcf, 0xac, 0x70, 0x52, 0xab, 0xb0, 0xca, 0xfb, 0xed, 0xf3, 0xdf, 0xc8, 0xa7,
	0x6f, 0xef, 0xcd, 0x7f, 0xbf, 0xf9, 0x6b, 0x48, 0xbf, 0x20, 0x07, 0xcb, 0x35, 0x94, 0x98, 0x02,
	0xab, 0x9c, 0x55, 0x2e, 0xaa, 0x83, 0x5a, 0xa1, 0x3e, 0x88, 0x29, 0x9c, 0xff, 0x57, 0x25, 0x2d,
	0xbf, 0xd4, 0x93, 0x98, 0x94, 0xd7, 0xa2, 0x9f, 0x10, 0x32, 0x49, 0xf5, 0x93, 0x48, 0x7d, 0xdd,
	0x18, 0xde, 0x1d, 0x54, 0x73, 0xa5, 0xef, 0x0c, 0xfd, 0x96, 0x50, 0x70, 0x7f, 0x83, 0x55, 0xe0,
	0xb8, 0xb3, 0x42, 0x65, 0x46, 0x5b, 0xc7, 0x3e, 0x42, 0xac, 0x1e, 0x9d, 0xc7, 0x68, 0xd0, 0x53,
	0xb2, 0xab, 0x15, 0xf0, 0xcc, 0x81, 0x61, 0x1b, 0x08, 0xed, 0x68, 0x05, 0x43, 0x07, 0x86, 0x36,
	0xc9, 0x56, 0x96, 0x8a, 0x39, 0xb0, 0x4d, 0xd4, 0xf3, 0x3f, 0x94, 0x92, 0x4d, 0x69, 0xe6, 0x3d,
	0xb6, 0x85, 0x22, 0xfe, 0xa6, 0x1f, 0x93, 0xea, 0x74, 0x96, 0x3a, 0x99, 0x88, 0xcc, 0xb1, 0xed,
	0xbc, 0xa2, 0x42, 0xa0, 0xdf, 0x10, 0x6a, 0xac, 0x1e, 0xcb, 0x14, 0xb8, 0xd2, 0x8e, 0xe7, 0xa5,
	0xb2, 0x1d, 0xc4, 0x8e, 0x82, 0xf3, 0xa0, 0xdd, 0x2d, 0xea, 0xbe, 0x3f, 0xa9, 0x4e, 0xfc, 0xe9,
	0xac, 0xd4, 0x56, 0xba, 0x05, 0xdb, 0x45, 0xb2, 0x86, 0x6a, 0x3f, 0x88, 0xf4, 0x2b, 0x12, 0xa3,
	0x3c, 0x1e, 0x8a, 0x55, 0x11, 0x3c, 0x0c, 0xfa, 0x4d, 0x90, 0xe9, 0x67, 0x64, 0x3f, 0xa2, 0xd2,
	0xcc, 0xbb, 0x8c, 0x20, 0xb6, 0x17, 0xb4, 0x3b, 0x33, 0xef, 0xae, 0x20, 0x3d, 0xb6, 0xb7, 0x8a,
	0xf4, 0xe8, 0x97, 0x24, 0x2e, 0xcc, 0x67, 0x2a, 0x3f, 0xe9, 0x3e, 0x52, 0x07, 0x41, 0xfe, 0x23,
	0x57, 0xe9, 0xd7, 0xa4, 0x1e, 0xc1, 0x65, 0x53, 0x6a, 0xa5, 0xd3, 0xde, 0x17, 0xbd, 0xf9, 0x9c,
	0xd4, 0x0a, 0x58, 0xbe, 0xc0, 0x88, 0x1d, 0x20, 0x18, 0xab, 0xb9, 0xf7, 0x1a, 0xed, 0x92, 0xe3,
	0x02, 0x12, 0x99, 0x03, 0x5b, 0x54, 0x70, 0x88, 0x74, 0x33, 0xd2, 0x68, 0xc6, 0x3a, 0x7e, 0x20,
	0x6c, 0x25, 0xb5, 0x2c, 0xe7, 0x08, 0x73, 0xc7, 0xa5, 0xdc, 0xb2, 0xa8, 0xef, 0x48, 0x73, 0x35,
	0x89, 0xb5, 0xd5, 0x31, 0x45, 0xcb, 0x29, 0xac, 0xb0, 0x43, 0x1a, 0x4e, 0xd8, 0x09, 0x38, 0x2e,
	0x46, 0x23, 0x0b, 0x59, 0x96, 0x77, 0x9a, 0xe6, 0xb7, 0x2e, 0xb7, 0xae, 0x73, 0x07, 0xfb, 0xbd,
	0x96, 0xef, 0xb1, 0xc6, 0x7a, 0xbe, 0xf7, 0xfa, 0x0a, 0xf9, 0x5b, 0xcb, 0x33, 0x27, 0x1c, 0xb0,
	0x66, 0xa9, 0xa9, 0x7d, 0x6d, 0xdd, 0xd0, 0xeb, 0xf4, 0x47, 0x72, 0x1a, 0x69, 0xa1, 0x94, 0x9e,
	0xa9, 0x04, 0x38, 0xbe, 0xae, 0xb9, 0x48, 0x59, 0x0b, 0x43, 0x27, 0x01, 0xb8, 0x0e, 0xfe, 0x5d,
	0xb0, 0xe9, 0x25, 0x69, 0xc5, 0x6c, 0xb6, 0x50, 0xc9, 0x32, 0x77, 0x8c, 0xb9, 0x46, 0x30, 0x87,
	0x0b, 0x95, 0x14, 0x99, 0x9f, 0x48, 0x3b, 0x66, 0x46, 0x90, 0x8a, 0x05, 0xb7, 0xf0, 0xbc, 0x0c,
	0x9e, 0x94, 0x36, 0xfc, 0xd5, 0x03, 0x03, 0x78, 0x2e, 0xc2, 0xaf, 0x9a, 0x8d, 0x1b, 0x3a, 0x39,
	0x05, 0x3d, 0x73, 0x8c, 0x95, 0x9a, 0xed, 0xf7, 0x7b, 0xcc, 0x9d, 0x75, 0xdb, 0x65, 0xa6, 0xc8,
	0x9d, 0xae, 0xdb, 0x2e, 0x33, 0x31, 0xdc, 0x23, 0x27, 0x52, 0xcd, 0x45, 0x2a, 0x47, 0x7c, 0x62,
	0x85, 0x72, 0xdc, 0xc2, 0x68, 0x96, 0xf8, 0x11, 0xc6, 0xda, 0x98, 0x6c, 0x05, 0xfb, 0xd6, 0xbb,
	0x83, 0x68, 0x9e, 0xff, 0xbb, 0x41, 0xda, 0x6f, 0x8f, 0x36, 0xfa, 0x33, 0x69, 0x97, 0xe7, 0x61,
	0xf1, 0xe2, 0xfd, 0x84, 0xbb, 0xc4, 0x09, 0xc7, 0x4a, 0x44, 0x3f, 0x3c, 0x7c, 0x31, 0x05, 0xff,
	0x0a, 0x92, 0x54, 0x27, 0xff, 0xc4, 0x14, 0xbb, 0xc2, 0xc0, 0x3e, 0x8a, 0x01, 0xf4, 0x77, 0xc0,
	0x41, 0x0a, 0x89, 0x9e, 0xf2, 0x1c, 0x76, 0x0b, 0x03, 0xac, 0x8b, 0xe4, 0x51, 0x70, 0x7e, 0xf1,
	0xc6, 0xe3, 0xc2, 0x80, 0xbf, 0x61, 0x16, 0x32, 0x67, 0x65, 0xe2, 0x5e, 0x5f, 0x99, 0xef, 0x11,
	0xaf, 0x47, 0x6b, 0x79, 0x67, 0xfe, 0xaf, 0x90, 0xe6, 0xba, 0x81, 0xcf, 0x7a, 0x67, 0x95, 0x8b,
	0xbd, 0xcb, 0x69, 0xe7, 0xc3, 0x7c, 0x31, 0x3a, 0x6b, 0x67, 0xfc, 0xa0, 0x51, 0x42, 0x6f, 0x50,
	0x7c, 0xda, 0xc6, 0x6f, 0xd7, 0xd5, 0xbb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x63, 0x49, 0x19,
	0xdf, 0x06, 0x00, 0x00,
}
