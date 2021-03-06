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
// source: gnssmgr_rx_info.proto

package cisco_ios_xr_gnss_oper_gnss_receiver_nodes_node_receivers_receiver

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

type GnssmgrRxInfo_KEYS struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GnssmgrRxInfo_KEYS) Reset()         { *m = GnssmgrRxInfo_KEYS{} }
func (m *GnssmgrRxInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*GnssmgrRxInfo_KEYS) ProtoMessage()    {}
func (*GnssmgrRxInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d7d628ebf2641e, []int{0}
}

func (m *GnssmgrRxInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GnssmgrRxInfo_KEYS.Unmarshal(m, b)
}
func (m *GnssmgrRxInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GnssmgrRxInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *GnssmgrRxInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GnssmgrRxInfo_KEYS.Merge(m, src)
}
func (m *GnssmgrRxInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_GnssmgrRxInfo_KEYS.Size(m)
}
func (m *GnssmgrRxInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_GnssmgrRxInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_GnssmgrRxInfo_KEYS proto.InternalMessageInfo

func (m *GnssmgrRxInfo_KEYS) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *GnssmgrRxInfo_KEYS) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type GnssmgrBagSatellite struct {
	PrnNumber            uint32   `protobuf:"varint,1,opt,name=prn_number,json=prnNumber,proto3" json:"prn_number,omitempty"`
	ChannelNumber        uint32   `protobuf:"varint,2,opt,name=channel_number,json=channelNumber,proto3" json:"channel_number,omitempty"`
	AcquisitionFlag      bool     `protobuf:"varint,3,opt,name=acquisition_flag,json=acquisitionFlag,proto3" json:"acquisition_flag,omitempty"`
	EphemerisFlag        bool     `protobuf:"varint,4,opt,name=ephemeris_flag,json=ephemerisFlag,proto3" json:"ephemeris_flag,omitempty"`
	SvType               string   `protobuf:"bytes,5,opt,name=sv_type,json=svType,proto3" json:"sv_type,omitempty"`
	SignalStrength       uint32   `protobuf:"varint,6,opt,name=signal_strength,json=signalStrength,proto3" json:"signal_strength,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GnssmgrBagSatellite) Reset()         { *m = GnssmgrBagSatellite{} }
func (m *GnssmgrBagSatellite) String() string { return proto.CompactTextString(m) }
func (*GnssmgrBagSatellite) ProtoMessage()    {}
func (*GnssmgrBagSatellite) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d7d628ebf2641e, []int{1}
}

func (m *GnssmgrBagSatellite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GnssmgrBagSatellite.Unmarshal(m, b)
}
func (m *GnssmgrBagSatellite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GnssmgrBagSatellite.Marshal(b, m, deterministic)
}
func (m *GnssmgrBagSatellite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GnssmgrBagSatellite.Merge(m, src)
}
func (m *GnssmgrBagSatellite) XXX_Size() int {
	return xxx_messageInfo_GnssmgrBagSatellite.Size(m)
}
func (m *GnssmgrBagSatellite) XXX_DiscardUnknown() {
	xxx_messageInfo_GnssmgrBagSatellite.DiscardUnknown(m)
}

var xxx_messageInfo_GnssmgrBagSatellite proto.InternalMessageInfo

func (m *GnssmgrBagSatellite) GetPrnNumber() uint32 {
	if m != nil {
		return m.PrnNumber
	}
	return 0
}

func (m *GnssmgrBagSatellite) GetChannelNumber() uint32 {
	if m != nil {
		return m.ChannelNumber
	}
	return 0
}

func (m *GnssmgrBagSatellite) GetAcquisitionFlag() bool {
	if m != nil {
		return m.AcquisitionFlag
	}
	return false
}

func (m *GnssmgrBagSatellite) GetEphemerisFlag() bool {
	if m != nil {
		return m.EphemerisFlag
	}
	return false
}

func (m *GnssmgrBagSatellite) GetSvType() string {
	if m != nil {
		return m.SvType
	}
	return ""
}

func (m *GnssmgrBagSatellite) GetSignalStrength() uint32 {
	if m != nil {
		return m.SignalStrength
	}
	return 0
}

type GnssmgrRxInfo struct {
	ReceiverNumber         uint32                 `protobuf:"varint,50,opt,name=receiver_number,json=receiverNumber,proto3" json:"receiver_number,omitempty"`
	Node                   string                 `protobuf:"bytes,51,opt,name=node,proto3" json:"node,omitempty"`
	Shutdown               bool                   `protobuf:"varint,52,opt,name=shutdown,proto3" json:"shutdown,omitempty"`
	AntiJamDisable         bool                   `protobuf:"varint,53,opt,name=anti_jam_disable,json=antiJamDisable,proto3" json:"anti_jam_disable,omitempty"`
	Constellation          string                 `protobuf:"bytes,54,opt,name=constellation,proto3" json:"constellation,omitempty"`
	SnrThreshold           uint32                 `protobuf:"varint,55,opt,name=snr_threshold,json=snrThreshold,proto3" json:"snr_threshold,omitempty"`
	ElevationThreshold     uint32                 `protobuf:"varint,56,opt,name=elevation_threshold,json=elevationThreshold,proto3" json:"elevation_threshold,omitempty"`
	PdopThreshold          uint32                 `protobuf:"varint,57,opt,name=pdop_threshold,json=pdopThreshold,proto3" json:"pdop_threshold,omitempty"`
	TraimThreshold         uint32                 `protobuf:"varint,58,opt,name=traim_threshold,json=traimThreshold,proto3" json:"traim_threshold,omitempty"`
	CableDelayCompensation int32                  `protobuf:"zigzag32,59,opt,name=cable_delay_compensation,json=cableDelayCompensation,proto3" json:"cable_delay_compensation,omitempty"`
	Polarity1Pps           string                 `protobuf:"bytes,60,opt,name=polarity1pps,proto3" json:"polarity1pps,omitempty"`
	Available              bool                   `protobuf:"varint,61,opt,name=available,proto3" json:"available,omitempty"`
	LockStatus             string                 `protobuf:"bytes,62,opt,name=lock_status,json=lockStatus,proto3" json:"lock_status,omitempty"`
	ReceiverMode           string                 `protobuf:"bytes,63,opt,name=receiver_mode,json=receiverMode,proto3" json:"receiver_mode,omitempty"`
	SurveyProgress         uint32                 `protobuf:"varint,64,opt,name=survey_progress,json=surveyProgress,proto3" json:"survey_progress,omitempty"`
	HoldoverDuration       uint32                 `protobuf:"varint,65,opt,name=holdover_duration,json=holdoverDuration,proto3" json:"holdover_duration,omitempty"`
	MajorAlarm             string                 `protobuf:"bytes,66,opt,name=major_alarm,json=majorAlarm,proto3" json:"major_alarm,omitempty"`
	MinorAlarm             string                 `protobuf:"bytes,67,opt,name=minor_alarm,json=minorAlarm,proto3" json:"minor_alarm,omitempty"`
	Pdop                   uint32                 `protobuf:"varint,68,opt,name=pdop,proto3" json:"pdop,omitempty"`
	Hdop                   uint32                 `protobuf:"varint,69,opt,name=hdop,proto3" json:"hdop,omitempty"`
	Vdop                   uint32                 `protobuf:"varint,70,opt,name=vdop,proto3" json:"vdop,omitempty"`
	Tdop                   uint32                 `protobuf:"varint,71,opt,name=tdop,proto3" json:"tdop,omitempty"`
	SatelliteDataKnown     bool                   `protobuf:"varint,72,opt,name=satellite_data_known,json=satelliteDataKnown,proto3" json:"satellite_data_known,omitempty"`
	Satellite              []*GnssmgrBagSatellite `protobuf:"bytes,73,rep,name=satellite,proto3" json:"satellite,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}               `json:"-"`
	XXX_unrecognized       []byte                 `json:"-"`
	XXX_sizecache          int32                  `json:"-"`
}

func (m *GnssmgrRxInfo) Reset()         { *m = GnssmgrRxInfo{} }
func (m *GnssmgrRxInfo) String() string { return proto.CompactTextString(m) }
func (*GnssmgrRxInfo) ProtoMessage()    {}
func (*GnssmgrRxInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d7d628ebf2641e, []int{2}
}

func (m *GnssmgrRxInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GnssmgrRxInfo.Unmarshal(m, b)
}
func (m *GnssmgrRxInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GnssmgrRxInfo.Marshal(b, m, deterministic)
}
func (m *GnssmgrRxInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GnssmgrRxInfo.Merge(m, src)
}
func (m *GnssmgrRxInfo) XXX_Size() int {
	return xxx_messageInfo_GnssmgrRxInfo.Size(m)
}
func (m *GnssmgrRxInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GnssmgrRxInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GnssmgrRxInfo proto.InternalMessageInfo

func (m *GnssmgrRxInfo) GetReceiverNumber() uint32 {
	if m != nil {
		return m.ReceiverNumber
	}
	return 0
}

func (m *GnssmgrRxInfo) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *GnssmgrRxInfo) GetShutdown() bool {
	if m != nil {
		return m.Shutdown
	}
	return false
}

func (m *GnssmgrRxInfo) GetAntiJamDisable() bool {
	if m != nil {
		return m.AntiJamDisable
	}
	return false
}

func (m *GnssmgrRxInfo) GetConstellation() string {
	if m != nil {
		return m.Constellation
	}
	return ""
}

func (m *GnssmgrRxInfo) GetSnrThreshold() uint32 {
	if m != nil {
		return m.SnrThreshold
	}
	return 0
}

func (m *GnssmgrRxInfo) GetElevationThreshold() uint32 {
	if m != nil {
		return m.ElevationThreshold
	}
	return 0
}

func (m *GnssmgrRxInfo) GetPdopThreshold() uint32 {
	if m != nil {
		return m.PdopThreshold
	}
	return 0
}

func (m *GnssmgrRxInfo) GetTraimThreshold() uint32 {
	if m != nil {
		return m.TraimThreshold
	}
	return 0
}

func (m *GnssmgrRxInfo) GetCableDelayCompensation() int32 {
	if m != nil {
		return m.CableDelayCompensation
	}
	return 0
}

func (m *GnssmgrRxInfo) GetPolarity1Pps() string {
	if m != nil {
		return m.Polarity1Pps
	}
	return ""
}

func (m *GnssmgrRxInfo) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *GnssmgrRxInfo) GetLockStatus() string {
	if m != nil {
		return m.LockStatus
	}
	return ""
}

func (m *GnssmgrRxInfo) GetReceiverMode() string {
	if m != nil {
		return m.ReceiverMode
	}
	return ""
}

func (m *GnssmgrRxInfo) GetSurveyProgress() uint32 {
	if m != nil {
		return m.SurveyProgress
	}
	return 0
}

func (m *GnssmgrRxInfo) GetHoldoverDuration() uint32 {
	if m != nil {
		return m.HoldoverDuration
	}
	return 0
}

func (m *GnssmgrRxInfo) GetMajorAlarm() string {
	if m != nil {
		return m.MajorAlarm
	}
	return ""
}

func (m *GnssmgrRxInfo) GetMinorAlarm() string {
	if m != nil {
		return m.MinorAlarm
	}
	return ""
}

func (m *GnssmgrRxInfo) GetPdop() uint32 {
	if m != nil {
		return m.Pdop
	}
	return 0
}

func (m *GnssmgrRxInfo) GetHdop() uint32 {
	if m != nil {
		return m.Hdop
	}
	return 0
}

func (m *GnssmgrRxInfo) GetVdop() uint32 {
	if m != nil {
		return m.Vdop
	}
	return 0
}

func (m *GnssmgrRxInfo) GetTdop() uint32 {
	if m != nil {
		return m.Tdop
	}
	return 0
}

func (m *GnssmgrRxInfo) GetSatelliteDataKnown() bool {
	if m != nil {
		return m.SatelliteDataKnown
	}
	return false
}

func (m *GnssmgrRxInfo) GetSatellite() []*GnssmgrBagSatellite {
	if m != nil {
		return m.Satellite
	}
	return nil
}

func init() {
	proto.RegisterType((*GnssmgrRxInfo_KEYS)(nil), "cisco_ios_xr_gnss_oper.gnss_receiver.nodes.node.receivers.receiver.gnssmgr_rx_info_KEYS")
	proto.RegisterType((*GnssmgrBagSatellite)(nil), "cisco_ios_xr_gnss_oper.gnss_receiver.nodes.node.receivers.receiver.gnssmgr_bag_satellite")
	proto.RegisterType((*GnssmgrRxInfo)(nil), "cisco_ios_xr_gnss_oper.gnss_receiver.nodes.node.receivers.receiver.gnssmgr_rx_info")
}

func init() { proto.RegisterFile("gnssmgr_rx_info.proto", fileDescriptor_57d7d628ebf2641e) }

var fileDescriptor_57d7d628ebf2641e = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x4e, 0x1b, 0x39,
	0x14, 0xc7, 0x95, 0x85, 0x0d, 0xc4, 0x90, 0x04, 0xbc, 0x2c, 0x6b, 0xad, 0x5a, 0x35, 0x4a, 0x8b,
	0x48, 0x55, 0x29, 0x6d, 0xa1, 0x1f, 0xf4, 0xbb, 0x40, 0xa0, 0x1f, 0xa8, 0x55, 0x15, 0xb8, 0xe1,
	0xca, 0x72, 0x66, 0xcc, 0x8c, 0x61, 0xc6, 0x9e, 0xda, 0xce, 0x40, 0x1e, 0xac, 0x0f, 0xd6, 0x37,
	0xa8, 0x7c, 0x3c, 0x33, 0x09, 0xa8, 0x77, 0xbd, 0x89, 0x3c, 0xbf, 0xf3, 0x3f, 0xf6, 0x39, 0xff,
	0xe3, 0x18, 0xfd, 0x1b, 0x49, 0x63, 0xd2, 0x48, 0x53, 0x7d, 0x45, 0x85, 0x3c, 0x53, 0xfd, 0x4c,
	0x2b, 0xab, 0xf0, 0x5e, 0x20, 0x4c, 0xa0, 0xa8, 0x50, 0x86, 0x5e, 0x69, 0xea, 0x34, 0x54, 0x65,
	0x5c, 0xf7, 0x61, 0xa5, 0x79, 0xc0, 0x45, 0xce, 0x75, 0x5f, 0xaa, 0x90, 0x1b, 0xf8, 0xed, 0x97,
	0xcc, 0x54, 0xab, 0xee, 0x1e, 0x5a, 0xbb, 0xb1, 0x39, 0x3d, 0x3a, 0x38, 0x3d, 0xc6, 0x18, 0xcd,
	0xbb, 0x14, 0x52, 0xeb, 0xd4, 0x7a, 0x8d, 0x21, 0xac, 0xf1, 0x3a, 0xaa, 0xcb, 0x71, 0x3a, 0xe2,
	0x9a, 0xfc, 0xd5, 0xa9, 0xf5, 0x9a, 0xc3, 0xe2, 0xab, 0xfb, 0xb3, 0x36, 0xad, 0x70, 0xc4, 0x22,
	0x6a, 0x98, 0xe5, 0x49, 0x22, 0x2c, 0xc7, 0xb7, 0x11, 0xca, 0xb4, 0xa4, 0x45, 0x56, 0x0d, 0xb2,
	0x1a, 0x99, 0x96, 0x5f, 0x01, 0xe0, 0x0d, 0xd4, 0x0a, 0x62, 0x26, 0x25, 0x4f, 0xe8, 0xb5, 0x8d,
	0x9b, 0x05, 0x2d, 0x64, 0xf7, 0xd1, 0x0a, 0x0b, 0xbe, 0x8f, 0x85, 0x11, 0x56, 0x28, 0x49, 0xcf,
	0x12, 0x16, 0x91, 0xb9, 0x4e, 0xad, 0xb7, 0x38, 0x6c, 0xcf, 0xf0, 0xc3, 0x84, 0x45, 0x6e, 0x47,
	0x9e, 0xc5, 0x3c, 0xe5, 0x5a, 0x18, 0x2f, 0x9c, 0x07, 0x61, 0xb3, 0xa2, 0x20, 0xfb, 0x0f, 0x2d,
	0x98, 0x9c, 0xda, 0x49, 0xc6, 0xc9, 0xdf, 0xd0, 0x60, 0xdd, 0xe4, 0x27, 0x93, 0x8c, 0xe3, 0x4d,
	0xd4, 0x36, 0x22, 0x92, 0x2c, 0xa1, 0xc6, 0x6a, 0x2e, 0x23, 0x1b, 0x93, 0x3a, 0x94, 0xd4, 0xf2,
	0xf8, 0xb8, 0xa0, 0xdd, 0x1f, 0x0b, 0xa8, 0x7d, 0xc3, 0x38, 0x97, 0x5c, 0xfa, 0x5a, 0xf6, 0xb3,
	0xe5, 0x93, 0x4b, 0x5c, 0x34, 0x54, 0x9a, 0xbb, 0x3d, 0x63, 0xee, 0xff, 0x68, 0xd1, 0xc4, 0x63,
	0x1b, 0xaa, 0x4b, 0x49, 0x9e, 0x40, 0xcd, 0xd5, 0x37, 0xee, 0xa1, 0x15, 0x26, 0xad, 0xa0, 0xe7,
	0x2c, 0xa5, 0xa1, 0x30, 0x6c, 0x94, 0x70, 0xf2, 0x14, 0x34, 0x2d, 0xc7, 0x3f, 0xb3, 0x74, 0xe0,
	0x29, 0xbe, 0x87, 0x9a, 0x81, 0x92, 0xc6, 0xf9, 0xcf, 0x9c, 0x29, 0xe4, 0x19, 0x1c, 0x71, 0x1d,
	0xe2, 0xbb, 0xa8, 0x69, 0xa4, 0xa6, 0x36, 0xd6, 0xdc, 0xc4, 0x2a, 0x09, 0xc9, 0x73, 0x28, 0x73,
	0xd9, 0x48, 0x7d, 0x52, 0x32, 0xfc, 0x10, 0xfd, 0xc3, 0x13, 0x9e, 0x43, 0xc6, 0x8c, 0x74, 0x07,
	0xa4, 0xb8, 0x0a, 0x4d, 0x13, 0x36, 0x50, 0x2b, 0x0b, 0x55, 0x36, 0xa3, 0x7d, 0xe1, 0xa7, 0xe9,
	0xe8, 0x54, 0xb6, 0x89, 0xda, 0x56, 0x33, 0x91, 0xce, 0xe8, 0x5e, 0x7a, 0x97, 0x00, 0x4f, 0x85,
	0x3b, 0x88, 0x04, 0xae, 0x29, 0x1a, 0xf2, 0x84, 0x4d, 0x68, 0xa0, 0xd2, 0x8c, 0x4b, 0xe3, 0xdb,
	0x7a, 0xd5, 0xa9, 0xf5, 0x56, 0x87, 0xeb, 0x10, 0x1f, 0xb8, 0xf0, 0xfe, 0x4c, 0x14, 0x77, 0xd1,
	0x72, 0xa6, 0x12, 0xa6, 0x85, 0x9d, 0x3c, 0xce, 0x32, 0x43, 0x5e, 0x83, 0x09, 0xd7, 0x18, 0xbe,
	0x85, 0x1a, 0x2c, 0x67, 0x22, 0x01, 0x33, 0xdf, 0x80, 0x99, 0x53, 0x80, 0xef, 0xa0, 0xa5, 0x44,
	0x05, 0x17, 0xd4, 0x58, 0x66, 0xc7, 0x86, 0xbc, 0x85, 0x0d, 0x90, 0x43, 0xc7, 0x40, 0x9c, 0x85,
	0xd5, 0xac, 0x53, 0x37, 0xcb, 0x77, 0xfe, 0x8c, 0x12, 0x7e, 0x71, 0x33, 0x75, 0xb7, 0x69, 0xac,
	0x73, 0x3e, 0xa1, 0x99, 0x56, 0x91, 0xe6, 0xc6, 0x90, 0xf7, 0xc5, 0x6d, 0x02, 0xfc, 0xad, 0xa0,
	0xf8, 0x01, 0x5a, 0x75, 0x2d, 0x2b, 0xb7, 0x5b, 0x38, 0xd6, 0xbe, 0xc7, 0x5d, 0x90, 0xae, 0x94,
	0x81, 0x41, 0xc1, 0x5d, 0x6d, 0x29, 0x3b, 0x57, 0x9a, 0xb2, 0x84, 0xe9, 0x94, 0xec, 0xf9, 0xda,
	0x00, 0xed, 0x3a, 0x02, 0x02, 0x21, 0x2b, 0xc1, 0x7e, 0x21, 0x70, 0xc8, 0x0b, 0x30, 0x9a, 0x77,
	0x33, 0x21, 0x03, 0x38, 0x01, 0xd6, 0x8e, 0xc5, 0x8e, 0x1d, 0x78, 0x16, 0x17, 0x2c, 0x77, 0xec,
	0xd0, 0xb3, 0xbc, 0x60, 0xd6, 0xb1, 0x0f, 0x9e, 0xb9, 0x35, 0x7e, 0x84, 0xd6, 0xaa, 0xff, 0x3c,
	0x0d, 0x99, 0x65, 0xf4, 0x42, 0xba, 0x7b, 0xfc, 0x11, 0x6c, 0xc5, 0x55, 0x6c, 0xc0, 0x2c, 0x3b,
	0x72, 0x11, 0x7c, 0x89, 0x1a, 0x15, 0x25, 0x9f, 0x3a, 0x73, 0xbd, 0xa5, 0xad, 0xd3, 0xfe, 0x9f,
	0x3f, 0x67, 0xfd, 0xdf, 0x3e, 0x43, 0xc3, 0xe9, 0x59, 0xa3, 0x3a, 0x3c, 0x9d, 0xdb, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x3e, 0xb9, 0xe5, 0x69, 0x53, 0x05, 0x00, 0x00,
}
