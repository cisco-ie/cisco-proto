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
	Enabled                bool                   `protobuf:"varint,52,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Shutdown               bool                   `protobuf:"varint,53,opt,name=shutdown,proto3" json:"shutdown,omitempty"`
	AntiJamDisable         bool                   `protobuf:"varint,54,opt,name=anti_jam_disable,json=antiJamDisable,proto3" json:"anti_jam_disable,omitempty"`
	Constellation          string                 `protobuf:"bytes,55,opt,name=constellation,proto3" json:"constellation,omitempty"`
	SnrThreshold           uint32                 `protobuf:"varint,56,opt,name=snr_threshold,json=snrThreshold,proto3" json:"snr_threshold,omitempty"`
	ElevationThreshold     uint32                 `protobuf:"varint,57,opt,name=elevation_threshold,json=elevationThreshold,proto3" json:"elevation_threshold,omitempty"`
	PdopThreshold          uint32                 `protobuf:"varint,58,opt,name=pdop_threshold,json=pdopThreshold,proto3" json:"pdop_threshold,omitempty"`
	TraimThreshold         uint32                 `protobuf:"varint,59,opt,name=traim_threshold,json=traimThreshold,proto3" json:"traim_threshold,omitempty"`
	CableDelayCompensation int32                  `protobuf:"zigzag32,60,opt,name=cable_delay_compensation,json=cableDelayCompensation,proto3" json:"cable_delay_compensation,omitempty"`
	Polarity1Pps           string                 `protobuf:"bytes,61,opt,name=polarity1pps,proto3" json:"polarity1pps,omitempty"`
	Available              bool                   `protobuf:"varint,62,opt,name=available,proto3" json:"available,omitempty"`
	LockStatus             string                 `protobuf:"bytes,63,opt,name=lock_status,json=lockStatus,proto3" json:"lock_status,omitempty"`
	ReceiverMode           string                 `protobuf:"bytes,64,opt,name=receiver_mode,json=receiverMode,proto3" json:"receiver_mode,omitempty"`
	SurveyProgress         uint32                 `protobuf:"varint,65,opt,name=survey_progress,json=surveyProgress,proto3" json:"survey_progress,omitempty"`
	HoldoverDuration       uint32                 `protobuf:"varint,66,opt,name=holdover_duration,json=holdoverDuration,proto3" json:"holdover_duration,omitempty"`
	MajorAlarm             string                 `protobuf:"bytes,67,opt,name=major_alarm,json=majorAlarm,proto3" json:"major_alarm,omitempty"`
	MinorAlarm             string                 `protobuf:"bytes,68,opt,name=minor_alarm,json=minorAlarm,proto3" json:"minor_alarm,omitempty"`
	Pdop                   uint32                 `protobuf:"varint,69,opt,name=pdop,proto3" json:"pdop,omitempty"`
	Hdop                   uint32                 `protobuf:"varint,70,opt,name=hdop,proto3" json:"hdop,omitempty"`
	Vdop                   uint32                 `protobuf:"varint,71,opt,name=vdop,proto3" json:"vdop,omitempty"`
	Tdop                   uint32                 `protobuf:"varint,72,opt,name=tdop,proto3" json:"tdop,omitempty"`
	Latitude               int32                  `protobuf:"zigzag32,73,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude              int32                  `protobuf:"zigzag32,74,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Altitude               string                 `protobuf:"bytes,75,opt,name=altitude,proto3" json:"altitude,omitempty"`
	Time                   uint32                 `protobuf:"varint,76,opt,name=time,proto3" json:"time,omitempty"`
	UtcOffset              string                 `protobuf:"bytes,77,opt,name=utc_offset,json=utcOffset,proto3" json:"utc_offset,omitempty"`
	FirmwareVersion        string                 `protobuf:"bytes,78,opt,name=firmware_version,json=firmwareVersion,proto3" json:"firmware_version,omitempty"`
	SatelliteDataKnown     bool                   `protobuf:"varint,79,opt,name=satellite_data_known,json=satelliteDataKnown,proto3" json:"satellite_data_known,omitempty"`
	Satellite              []*GnssmgrBagSatellite `protobuf:"bytes,80,rep,name=satellite,proto3" json:"satellite,omitempty"`
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

func (m *GnssmgrRxInfo) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
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

func (m *GnssmgrRxInfo) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *GnssmgrRxInfo) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *GnssmgrRxInfo) GetAltitude() string {
	if m != nil {
		return m.Altitude
	}
	return ""
}

func (m *GnssmgrRxInfo) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *GnssmgrRxInfo) GetUtcOffset() string {
	if m != nil {
		return m.UtcOffset
	}
	return ""
}

func (m *GnssmgrRxInfo) GetFirmwareVersion() string {
	if m != nil {
		return m.FirmwareVersion
	}
	return ""
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
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x5b, 0x8f, 0x1b, 0x35,
	0x14, 0xc7, 0x15, 0xba, 0xa4, 0x8d, 0xdb, 0x24, 0x5b, 0x53, 0x8a, 0x85, 0x40, 0x44, 0x81, 0xaa,
	0x41, 0x48, 0x01, 0x5a, 0x2e, 0xe5, 0x4e, 0xb7, 0x69, 0x81, 0x2e, 0xbd, 0x28, 0xbb, 0x42, 0xda,
	0x27, 0xcb, 0x99, 0x71, 0x66, 0xbc, 0xeb, 0xb1, 0x07, 0xdb, 0x33, 0xbb, 0x79, 0xe6, 0x13, 0xf3,
	0x0d, 0xd0, 0x39, 0x9e, 0x99, 0x64, 0x57, 0xbc, 0xf5, 0x25, 0xb2, 0x7f, 0xe7, 0x7f, 0x6c, 0x9f,
	0xbf, 0x9d, 0x33, 0xe4, 0xdd, 0xcc, 0x78, 0x5f, 0x64, 0x8e, 0xbb, 0x0b, 0xae, 0xcc, 0xda, 0xce,
	0x4b, 0x67, 0x83, 0xa5, 0x07, 0x89, 0xf2, 0x89, 0xe5, 0xca, 0x7a, 0x7e, 0xe1, 0x38, 0x68, 0xb8,
	0x2d, 0xa5, 0x9b, 0xe3, 0xc8, 0xc9, 0x44, 0xaa, 0x5a, 0xba, 0xb9, 0xb1, 0xa9, 0xf4, 0xf8, 0x3b,
	0x6f, 0x99, 0xef, 0x46, 0xd3, 0x03, 0x72, 0xe7, 0xca, 0xe2, 0xfc, 0xf0, 0xe9, 0xc9, 0x11, 0xa5,
	0x64, 0x0f, 0x52, 0x58, 0x6f, 0xd2, 0x9b, 0x0d, 0x96, 0x38, 0xa6, 0x77, 0x49, 0xdf, 0x54, 0xc5,
	0x4a, 0x3a, 0xf6, 0xd6, 0xa4, 0x37, 0x1b, 0x2e, 0x9b, 0xd9, 0xf4, 0xdf, 0xde, 0xf6, 0x84, 0x2b,
	0x91, 0x71, 0x2f, 0x82, 0xd4, 0x5a, 0x05, 0x49, 0x3f, 0x24, 0xa4, 0x74, 0x86, 0x37, 0x59, 0x3d,
	0xcc, 0x1a, 0x94, 0xce, 0xbc, 0x44, 0x40, 0xef, 0x91, 0x51, 0x92, 0x0b, 0x63, 0xa4, 0xe6, 0x97,
	0x16, 0x1e, 0x36, 0xb4, 0x91, 0x7d, 0x4a, 0xf6, 0x45, 0xf2, 0x77, 0xa5, 0xbc, 0x0a, 0xca, 0x1a,
	0xbe, 0xd6, 0x22, 0x63, 0xd7, 0x26, 0xbd, 0xd9, 0x8d, 0xe5, 0x78, 0x87, 0x3f, 0xd3, 0x22, 0x83,
	0x15, 0x65, 0x99, 0xcb, 0x42, 0x3a, 0xe5, 0xa3, 0x70, 0x0f, 0x85, 0xc3, 0x8e, 0xa2, 0xec, 0x3d,
	0x72, 0xdd, 0xd7, 0x3c, 0x6c, 0x4a, 0xc9, 0xde, 0xc6, 0x02, 0xfb, 0xbe, 0x3e, 0xde, 0x94, 0x92,
	0xde, 0x27, 0x63, 0xaf, 0x32, 0x23, 0x34, 0xf7, 0xc1, 0x49, 0x93, 0x85, 0x9c, 0xf5, 0xf1, 0x48,
	0xa3, 0x88, 0x8f, 0x1a, 0x3a, 0xfd, 0x67, 0x40, 0xc6, 0x57, 0x8c, 0x83, 0xe4, 0xd6, 0xd7, 0xb6,
	0x9e, 0x07, 0x31, 0xb9, 0xc5, 0x4d, 0x41, 0xad, 0xb9, 0x0f, 0x77, 0xcc, 0x65, 0xe4, 0xba, 0x34,
	0x62, 0xa5, 0x65, 0xca, 0xbe, 0xc2, 0x23, 0xb7, 0x53, 0xfa, 0x3e, 0xb9, 0xe1, 0xf3, 0x2a, 0xa4,
	0xf6, 0xdc, 0xb0, 0xaf, 0x31, 0xd4, 0xcd, 0xe9, 0x8c, 0xec, 0x0b, 0x13, 0x14, 0x3f, 0x15, 0x05,
	0x4f, 0x95, 0x87, 0x04, 0xf6, 0x0d, 0x6a, 0x46, 0xc0, 0x9f, 0x8b, 0x62, 0x11, 0x29, 0xfd, 0x84,
	0x0c, 0x13, 0x6b, 0x3c, 0xdc, 0x8c, 0x00, 0xbb, 0xd8, 0xb7, 0xb8, 0xf9, 0x65, 0x48, 0x3f, 0x26,
	0x43, 0x6f, 0x1c, 0x0f, 0xb9, 0x93, 0x3e, 0xb7, 0x3a, 0x65, 0x8f, 0xb0, 0x80, 0x5b, 0xde, 0xb8,
	0xe3, 0x96, 0xd1, 0xcf, 0xc9, 0x3b, 0x52, 0xcb, 0x1a, 0x33, 0x76, 0xa4, 0xdf, 0xa1, 0x94, 0x76,
	0xa1, 0x6d, 0xc2, 0x3d, 0x32, 0x2a, 0x53, 0x5b, 0xee, 0x68, 0xbf, 0x8f, 0xf7, 0x0c, 0x74, 0x2b,
	0xbb, 0x4f, 0xc6, 0xc1, 0x09, 0x55, 0xec, 0xe8, 0x7e, 0x88, 0xfe, 0x21, 0xde, 0x0a, 0x1f, 0x11,
	0x96, 0x40, 0x51, 0x3c, 0x95, 0x5a, 0x6c, 0x78, 0x62, 0x8b, 0x52, 0x1a, 0x1f, 0xcb, 0xfa, 0x71,
	0xd2, 0x9b, 0xdd, 0x5e, 0xde, 0xc5, 0xf8, 0x02, 0xc2, 0x4f, 0x76, 0xa2, 0x74, 0x4a, 0x6e, 0x95,
	0x56, 0x0b, 0xa7, 0xc2, 0xe6, 0xcb, 0xb2, 0xf4, 0xec, 0x27, 0x34, 0xe1, 0x12, 0xa3, 0x1f, 0x90,
	0x81, 0xa8, 0x85, 0xd2, 0x68, 0xe6, 0xcf, 0x68, 0xe6, 0x16, 0xd0, 0x8f, 0xc8, 0x4d, 0x6d, 0x93,
	0x33, 0xee, 0x83, 0x08, 0x95, 0x67, 0xbf, 0xe0, 0x02, 0x04, 0xd0, 0x11, 0x12, 0xb0, 0xb0, 0x7b,
	0x05, 0x05, 0xdc, 0xf2, 0xaf, 0x71, 0x8f, 0x16, 0xbe, 0x80, 0xdb, 0x86, 0x77, 0x56, 0xb9, 0x5a,
	0x6e, 0x78, 0xe9, 0x6c, 0xe6, 0xa4, 0xf7, 0xec, 0x71, 0xf3, 0xce, 0x10, 0xbf, 0x6e, 0x28, 0xfd,
	0x8c, 0xdc, 0x86, 0x92, 0x2d, 0xac, 0x96, 0x56, 0x2e, 0xd6, 0x78, 0x80, 0xd2, 0xfd, 0x36, 0xb0,
	0x68, 0x38, 0x9c, 0xad, 0x10, 0xa7, 0xd6, 0x71, 0xa1, 0x85, 0x2b, 0xd8, 0x93, 0x78, 0x36, 0x44,
	0x8f, 0x81, 0xa0, 0x40, 0x99, 0x4e, 0xb0, 0x68, 0x04, 0x80, 0xa2, 0x80, 0x92, 0x3d, 0xb8, 0x13,
	0xf6, 0x14, 0x77, 0xc0, 0x31, 0xb0, 0x1c, 0xd8, 0xb3, 0xc8, 0xf2, 0x86, 0xd5, 0xc0, 0x7e, 0x8b,
	0xac, 0x6e, 0x58, 0x00, 0xf6, 0x7b, 0x64, 0x30, 0x86, 0xb7, 0x0b, 0x2f, 0x2b, 0x54, 0xa9, 0x64,
	0x7f, 0xe0, 0xcd, 0x74, 0x73, 0xf0, 0x59, 0x5b, 0x93, 0xc5, 0xe0, 0x73, 0x0c, 0x6e, 0x01, 0x64,
	0x0a, 0xdd, 0x64, 0x1e, 0xe2, 0x39, 0xbb, 0x39, 0xee, 0xa4, 0x0a, 0xc9, 0xfe, 0x6c, 0x76, 0x52,
	0x05, 0xb6, 0x9a, 0x2a, 0x24, 0xdc, 0xae, 0xd7, 0x5e, 0x06, 0xf6, 0x02, 0x33, 0x06, 0x55, 0x48,
	0x5e, 0x21, 0x80, 0x1e, 0xb2, 0x56, 0xae, 0x38, 0x17, 0x4e, 0x72, 0xe8, 0x80, 0x60, 0xe3, 0x4b,
	0x14, 0x8d, 0x5b, 0xfe, 0x57, 0xc4, 0xf4, 0x0b, 0x72, 0xa7, 0xeb, 0x60, 0x3c, 0x15, 0x41, 0xf0,
	0x33, 0x03, 0xff, 0xbd, 0x57, 0xf8, 0x14, 0x68, 0x17, 0x5b, 0x88, 0x20, 0x0e, 0x21, 0x42, 0xcf,
	0xc9, 0xa0, 0xa3, 0xec, 0xf5, 0xe4, 0xda, 0xec, 0xe6, 0x83, 0x93, 0xf9, 0x9b, 0x37, 0xe7, 0xf9,
	0xff, 0x36, 0xd5, 0xe5, 0x76, 0xaf, 0x55, 0x1f, 0x3f, 0x04, 0x0f, 0xff, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x9e, 0xd3, 0xa3, 0xf7, 0x21, 0x06, 0x00, 0x00,
}