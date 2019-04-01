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
// source: coh_time_stats.proto

package cisco_ios_xr_ncs5500_coherent_node_oper_coherent_nodes_node_coherent_time_stats

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

type CohTimeStats_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CohTimeStats_KEYS) Reset()         { *m = CohTimeStats_KEYS{} }
func (m *CohTimeStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*CohTimeStats_KEYS) ProtoMessage()    {}
func (*CohTimeStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d33926ee67332a6e, []int{0}
}

func (m *CohTimeStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CohTimeStats_KEYS.Unmarshal(m, b)
}
func (m *CohTimeStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CohTimeStats_KEYS.Marshal(b, m, deterministic)
}
func (m *CohTimeStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CohTimeStats_KEYS.Merge(m, src)
}
func (m *CohTimeStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_CohTimeStats_KEYS.Size(m)
}
func (m *CohTimeStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_CohTimeStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_CohTimeStats_KEYS proto.InternalMessageInfo

func (m *CohTimeStats_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type CohOpTimeStats struct {
	Start                []string `protobuf:"bytes,1,rep,name=start,proto3" json:"start,omitempty"`
	End                  []string `protobuf:"bytes,2,rep,name=end,proto3" json:"end,omitempty"`
	TimeTaken            []string `protobuf:"bytes,3,rep,name=time_taken,json=timeTaken,proto3" json:"time_taken,omitempty"`
	WorstTime            []string `protobuf:"bytes,4,rep,name=worst_time,json=worstTime,proto3" json:"worst_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CohOpTimeStats) Reset()         { *m = CohOpTimeStats{} }
func (m *CohOpTimeStats) String() string { return proto.CompactTextString(m) }
func (*CohOpTimeStats) ProtoMessage()    {}
func (*CohOpTimeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_d33926ee67332a6e, []int{1}
}

func (m *CohOpTimeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CohOpTimeStats.Unmarshal(m, b)
}
func (m *CohOpTimeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CohOpTimeStats.Marshal(b, m, deterministic)
}
func (m *CohOpTimeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CohOpTimeStats.Merge(m, src)
}
func (m *CohOpTimeStats) XXX_Size() int {
	return xxx_messageInfo_CohOpTimeStats.Size(m)
}
func (m *CohOpTimeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_CohOpTimeStats.DiscardUnknown(m)
}

var xxx_messageInfo_CohOpTimeStats proto.InternalMessageInfo

func (m *CohOpTimeStats) GetStart() []string {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *CohOpTimeStats) GetEnd() []string {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *CohOpTimeStats) GetTimeTaken() []string {
	if m != nil {
		return m.TimeTaken
	}
	return nil
}

func (m *CohOpTimeStats) GetWorstTime() []string {
	if m != nil {
		return m.WorstTime
	}
	return nil
}

type CohPortOpStats struct {
	LaserState           bool            `protobuf:"varint,1,opt,name=laser_state,json=laserState,proto3" json:"laser_state,omitempty"`
	LaserOnStats         *CohOpTimeStats `protobuf:"bytes,2,opt,name=laser_on_stats,json=laserOnStats,proto3" json:"laser_on_stats,omitempty"`
	LaserOffStats        *CohOpTimeStats `protobuf:"bytes,3,opt,name=laser_off_stats,json=laserOffStats,proto3" json:"laser_off_stats,omitempty"`
	ProvisionedFrequency uint32          `protobuf:"varint,4,opt,name=provisioned_frequency,json=provisionedFrequency,proto3" json:"provisioned_frequency,omitempty"`
	WlOpStats            *CohOpTimeStats `protobuf:"bytes,5,opt,name=wl_op_stats,json=wlOpStats,proto3" json:"wl_op_stats,omitempty"`
	TxPower              uint32          `protobuf:"varint,6,opt,name=tx_power,json=txPower,proto3" json:"tx_power,omitempty"`
	TxpwrOpStats         *CohOpTimeStats `protobuf:"bytes,7,opt,name=txpwr_op_stats,json=txpwrOpStats,proto3" json:"txpwr_op_stats,omitempty"`
	CdMin                uint32          `protobuf:"varint,8,opt,name=cd_min,json=cdMin,proto3" json:"cd_min,omitempty"`
	CdminOpStats         *CohOpTimeStats `protobuf:"bytes,9,opt,name=cdmin_op_stats,json=cdminOpStats,proto3" json:"cdmin_op_stats,omitempty"`
	CdMax                uint32          `protobuf:"varint,10,opt,name=cd_max,json=cdMax,proto3" json:"cd_max,omitempty"`
	CdmaxOpStats         *CohOpTimeStats `protobuf:"bytes,11,opt,name=cdmax_op_stats,json=cdmaxOpStats,proto3" json:"cdmax_op_stats,omitempty"`
	TrafficType          uint32          `protobuf:"varint,12,opt,name=traffic_type,json=trafficType,proto3" json:"traffic_type,omitempty"`
	TraffictypeOpStats   *CohOpTimeStats `protobuf:"bytes,13,opt,name=traffictype_op_stats,json=traffictypeOpStats,proto3" json:"traffictype_op_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CohPortOpStats) Reset()         { *m = CohPortOpStats{} }
func (m *CohPortOpStats) String() string { return proto.CompactTextString(m) }
func (*CohPortOpStats) ProtoMessage()    {}
func (*CohPortOpStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_d33926ee67332a6e, []int{2}
}

func (m *CohPortOpStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CohPortOpStats.Unmarshal(m, b)
}
func (m *CohPortOpStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CohPortOpStats.Marshal(b, m, deterministic)
}
func (m *CohPortOpStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CohPortOpStats.Merge(m, src)
}
func (m *CohPortOpStats) XXX_Size() int {
	return xxx_messageInfo_CohPortOpStats.Size(m)
}
func (m *CohPortOpStats) XXX_DiscardUnknown() {
	xxx_messageInfo_CohPortOpStats.DiscardUnknown(m)
}

var xxx_messageInfo_CohPortOpStats proto.InternalMessageInfo

func (m *CohPortOpStats) GetLaserState() bool {
	if m != nil {
		return m.LaserState
	}
	return false
}

func (m *CohPortOpStats) GetLaserOnStats() *CohOpTimeStats {
	if m != nil {
		return m.LaserOnStats
	}
	return nil
}

func (m *CohPortOpStats) GetLaserOffStats() *CohOpTimeStats {
	if m != nil {
		return m.LaserOffStats
	}
	return nil
}

func (m *CohPortOpStats) GetProvisionedFrequency() uint32 {
	if m != nil {
		return m.ProvisionedFrequency
	}
	return 0
}

func (m *CohPortOpStats) GetWlOpStats() *CohOpTimeStats {
	if m != nil {
		return m.WlOpStats
	}
	return nil
}

func (m *CohPortOpStats) GetTxPower() uint32 {
	if m != nil {
		return m.TxPower
	}
	return 0
}

func (m *CohPortOpStats) GetTxpwrOpStats() *CohOpTimeStats {
	if m != nil {
		return m.TxpwrOpStats
	}
	return nil
}

func (m *CohPortOpStats) GetCdMin() uint32 {
	if m != nil {
		return m.CdMin
	}
	return 0
}

func (m *CohPortOpStats) GetCdminOpStats() *CohOpTimeStats {
	if m != nil {
		return m.CdminOpStats
	}
	return nil
}

func (m *CohPortOpStats) GetCdMax() uint32 {
	if m != nil {
		return m.CdMax
	}
	return 0
}

func (m *CohPortOpStats) GetCdmaxOpStats() *CohOpTimeStats {
	if m != nil {
		return m.CdmaxOpStats
	}
	return nil
}

func (m *CohPortOpStats) GetTrafficType() uint32 {
	if m != nil {
		return m.TrafficType
	}
	return 0
}

func (m *CohPortOpStats) GetTraffictypeOpStats() *CohOpTimeStats {
	if m != nil {
		return m.TraffictypeOpStats
	}
	return nil
}

type CohTimeStats struct {
	DriverInit               []string          `protobuf:"bytes,50,rep,name=driver_init,json=driverInit,proto3" json:"driver_init,omitempty"`
	DriverOperational        []string          `protobuf:"bytes,51,rep,name=driver_operational,json=driverOperational,proto3" json:"driver_operational,omitempty"`
	DeviceCreated            []string          `protobuf:"bytes,52,rep,name=device_created,json=deviceCreated,proto3" json:"device_created,omitempty"`
	OpticsControllersCreated []string          `protobuf:"bytes,53,rep,name=optics_controllers_created,json=opticsControllersCreated,proto3" json:"optics_controllers_created,omitempty"`
	DspControllersCreated    []string          `protobuf:"bytes,54,rep,name=dsp_controllers_created,json=dspControllersCreated,proto3" json:"dsp_controllers_created,omitempty"`
	EthIntfCreated           []string          `protobuf:"bytes,55,rep,name=eth_intf_created,json=ethIntfCreated,proto3" json:"eth_intf_created,omitempty"`
	OptsEaBulkCreate         *CohOpTimeStats   `protobuf:"bytes,56,opt,name=opts_ea_bulk_create,json=optsEaBulkCreate,proto3" json:"opts_ea_bulk_create,omitempty"`
	OptsEaBulkUpdate         *CohOpTimeStats   `protobuf:"bytes,57,opt,name=opts_ea_bulk_update,json=optsEaBulkUpdate,proto3" json:"opts_ea_bulk_update,omitempty"`
	DspEaBulkCreate          *CohOpTimeStats   `protobuf:"bytes,58,opt,name=dsp_ea_bulk_create,json=dspEaBulkCreate,proto3" json:"dsp_ea_bulk_create,omitempty"`
	DspEaBulkUpdate          *CohOpTimeStats   `protobuf:"bytes,59,opt,name=dsp_ea_bulk_update,json=dspEaBulkUpdate,proto3" json:"dsp_ea_bulk_update,omitempty"`
	PortStat                 []*CohPortOpStats `protobuf:"bytes,60,rep,name=port_stat,json=portStat,proto3" json:"port_stat,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *CohTimeStats) Reset()         { *m = CohTimeStats{} }
func (m *CohTimeStats) String() string { return proto.CompactTextString(m) }
func (*CohTimeStats) ProtoMessage()    {}
func (*CohTimeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_d33926ee67332a6e, []int{3}
}

func (m *CohTimeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CohTimeStats.Unmarshal(m, b)
}
func (m *CohTimeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CohTimeStats.Marshal(b, m, deterministic)
}
func (m *CohTimeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CohTimeStats.Merge(m, src)
}
func (m *CohTimeStats) XXX_Size() int {
	return xxx_messageInfo_CohTimeStats.Size(m)
}
func (m *CohTimeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_CohTimeStats.DiscardUnknown(m)
}

var xxx_messageInfo_CohTimeStats proto.InternalMessageInfo

func (m *CohTimeStats) GetDriverInit() []string {
	if m != nil {
		return m.DriverInit
	}
	return nil
}

func (m *CohTimeStats) GetDriverOperational() []string {
	if m != nil {
		return m.DriverOperational
	}
	return nil
}

func (m *CohTimeStats) GetDeviceCreated() []string {
	if m != nil {
		return m.DeviceCreated
	}
	return nil
}

func (m *CohTimeStats) GetOpticsControllersCreated() []string {
	if m != nil {
		return m.OpticsControllersCreated
	}
	return nil
}

func (m *CohTimeStats) GetDspControllersCreated() []string {
	if m != nil {
		return m.DspControllersCreated
	}
	return nil
}

func (m *CohTimeStats) GetEthIntfCreated() []string {
	if m != nil {
		return m.EthIntfCreated
	}
	return nil
}

func (m *CohTimeStats) GetOptsEaBulkCreate() *CohOpTimeStats {
	if m != nil {
		return m.OptsEaBulkCreate
	}
	return nil
}

func (m *CohTimeStats) GetOptsEaBulkUpdate() *CohOpTimeStats {
	if m != nil {
		return m.OptsEaBulkUpdate
	}
	return nil
}

func (m *CohTimeStats) GetDspEaBulkCreate() *CohOpTimeStats {
	if m != nil {
		return m.DspEaBulkCreate
	}
	return nil
}

func (m *CohTimeStats) GetDspEaBulkUpdate() *CohOpTimeStats {
	if m != nil {
		return m.DspEaBulkUpdate
	}
	return nil
}

func (m *CohTimeStats) GetPortStat() []*CohPortOpStats {
	if m != nil {
		return m.PortStat
	}
	return nil
}

func init() {
	proto.RegisterType((*CohTimeStats_KEYS)(nil), "cisco_ios_xr_ncs5500_coherent_node_oper.coherent.nodes.node.coherent_time_stats.coh_time_stats_KEYS")
	proto.RegisterType((*CohOpTimeStats)(nil), "cisco_ios_xr_ncs5500_coherent_node_oper.coherent.nodes.node.coherent_time_stats.coh_op_time_stats")
	proto.RegisterType((*CohPortOpStats)(nil), "cisco_ios_xr_ncs5500_coherent_node_oper.coherent.nodes.node.coherent_time_stats.coh_port_op_stats")
	proto.RegisterType((*CohTimeStats)(nil), "cisco_ios_xr_ncs5500_coherent_node_oper.coherent.nodes.node.coherent_time_stats.coh_time_stats")
}

func init() { proto.RegisterFile("coh_time_stats.proto", fileDescriptor_d33926ee67332a6e) }

var fileDescriptor_d33926ee67332a6e = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x96, 0x4d, 0x6b, 0x1b, 0x3b,
	0x14, 0x86, 0x99, 0x38, 0x1f, 0x1e, 0x39, 0x76, 0x12, 0xc5, 0xe1, 0xea, 0xde, 0x4b, 0xa9, 0x6b,
	0x28, 0x78, 0x53, 0x13, 0x92, 0x26, 0xfd, 0xca, 0xaa, 0x21, 0x85, 0x50, 0x5a, 0x17, 0x27, 0x5d,
	0x74, 0x25, 0x94, 0x19, 0x0d, 0x16, 0x19, 0x4b, 0xaa, 0x24, 0xc7, 0xe3, 0x55, 0xa1, 0xab, 0xd2,
	0x4d, 0x29, 0xf4, 0x57, 0xf4, 0xf7, 0xf5, 0x07, 0x14, 0x49, 0x33, 0xfe, 0x48, 0xb2, 0x6c, 0x27,
	0x9b, 0x90, 0xf3, 0x9e, 0xf3, 0xce, 0xfb, 0x70, 0x84, 0x85, 0x40, 0x33, 0x12, 0x03, 0x6c, 0xd8,
	0x90, 0x62, 0x6d, 0x88, 0xd1, 0x5d, 0xa9, 0x84, 0x11, 0xb0, 0x17, 0x31, 0x1d, 0x09, 0xcc, 0x84,
	0xc6, 0x99, 0xc2, 0x3c, 0xd2, 0x07, 0x07, 0xbb, 0xbb, 0x38, 0x12, 0x03, 0xaa, 0x28, 0x37, 0x98,
	0x8b, 0x98, 0x62, 0x21, 0xa9, 0xea, 0x16, 0x52, 0xd7, 0x4a, 0xda, 0xfd, 0x9d, 0x6a, 0x73, 0x9f,
	0x6d, 0xef, 0x81, 0xed, 0xc5, 0x20, 0xfc, 0xfa, 0xe4, 0xc3, 0x19, 0xfc, 0x1f, 0x84, 0xee, 0x5b,
	0x9c, 0x0c, 0x29, 0x0a, 0x5a, 0x41, 0x27, 0xec, 0x57, 0xad, 0xf0, 0x96, 0x0c, 0x69, 0x7b, 0x02,
	0xb6, 0xac, 0x47, 0xc8, 0x39, 0x1b, 0x6c, 0x82, 0x15, 0x6d, 0x88, 0x32, 0x28, 0x68, 0x55, 0x3a,
	0x61, 0xdf, 0x17, 0x70, 0x13, 0x54, 0x28, 0x8f, 0xd1, 0x92, 0xd3, 0xec, 0xbf, 0xf0, 0x1e, 0x00,
	0xce, 0x65, 0xc8, 0x25, 0xe5, 0xa8, 0xe2, 0x1a, 0xa1, 0x55, 0xce, 0xad, 0x60, 0xdb, 0x63, 0xa1,
	0xb4, 0x67, 0x44, 0xcb, 0xbe, 0xed, 0x94, 0x73, 0x36, 0xa4, 0xed, 0x9f, 0xa1, 0xcf, 0x96, 0x42,
	0x19, 0x0b, 0xe0, 0xb3, 0xef, 0x83, 0x5a, 0x4a, 0x34, 0x55, 0xae, 0xf4, 0xbc, 0xd5, 0x3e, 0x70,
	0xd2, 0x99, 0x55, 0xe0, 0x97, 0x00, 0x34, 0xfc, 0x84, 0xe0, 0xde, 0x83, 0x96, 0x5a, 0x41, 0xa7,
	0xb6, 0x77, 0xd1, 0xfd, 0xc3, 0x0b, 0xed, 0xde, 0xd8, 0x4c, 0x7f, 0xdd, 0x25, 0xf7, 0xf8, 0x99,
	0x63, 0xfd, 0x1a, 0x80, 0x8d, 0x1c, 0x25, 0x49, 0x72, 0x96, 0x4a, 0x69, 0x2c, 0x75, 0xcf, 0x92,
	0x24, 0x1e, 0x66, 0x1f, 0xec, 0x48, 0x25, 0xae, 0x98, 0x66, 0x82, 0xd3, 0x18, 0x27, 0x8a, 0x7e,
	0x1c, 0x51, 0x1e, 0x4d, 0xd0, 0x72, 0x2b, 0xe8, 0xd4, 0xfb, 0xcd, 0xb9, 0xe6, 0xab, 0xa2, 0x07,
	0x3f, 0x07, 0xa0, 0x36, 0x4e, 0xa7, 0xdb, 0x47, 0x2b, 0xa5, 0xd1, 0x87, 0xe3, 0xb4, 0x27, 0x3d,
	0xf9, 0xbf, 0xa0, 0x6a, 0x32, 0x2c, 0xc5, 0x98, 0x2a, 0xb4, 0xea, 0x60, 0xd7, 0x4c, 0xf6, 0xce,
	0x96, 0xee, 0xb0, 0x4d, 0x26, 0xc7, 0x6a, 0x86, 0xb8, 0x56, 0xde, 0x61, 0xbb, 0xe4, 0x82, 0x72,
	0x07, 0xac, 0x46, 0x31, 0x1e, 0x32, 0x8e, 0xaa, 0x8e, 0x71, 0x25, 0x8a, 0xdf, 0x30, 0xee, 0x08,
	0xa3, 0x78, 0xc8, 0xf8, 0x8c, 0x30, 0x2c, 0x8f, 0xd0, 0x25, 0x5f, 0x23, 0x24, 0x19, 0x02, 0x53,
	0x42, 0x92, 0x15, 0x84, 0x24, 0x9b, 0x11, 0xd6, 0x4a, 0x25, 0x24, 0x59, 0x41, 0xf8, 0x00, 0xac,
	0x1b, 0x45, 0x92, 0x84, 0x45, 0xd8, 0x4c, 0x24, 0x45, 0xeb, 0x8e, 0xb3, 0x96, 0x6b, 0xe7, 0x13,
	0x49, 0xe1, 0x8f, 0x00, 0x34, 0xf3, 0xda, 0x8e, 0xcc, 0x98, 0xeb, 0xa5, 0x31, 0xc3, 0xb9, 0xfc,
	0x9c, 0xbc, 0xfd, 0x6b, 0x0d, 0x34, 0x16, 0x2f, 0x57, 0x7b, 0x53, 0xc5, 0x8a, 0x5d, 0x51, 0x85,
	0x19, 0x67, 0x06, 0xed, 0xb9, 0xfb, 0x0d, 0x78, 0xe9, 0x94, 0x33, 0x03, 0x1f, 0x01, 0x98, 0x0f,
	0x58, 0x20, 0x62, 0x98, 0xe0, 0x24, 0x45, 0xfb, 0x6e, 0x6e, 0xcb, 0x77, 0x7a, 0xb3, 0x06, 0x7c,
	0x08, 0x1a, 0x31, 0xbd, 0x62, 0x11, 0xc5, 0x91, 0xa2, 0xc4, 0xd0, 0x18, 0x3d, 0x76, 0xa3, 0x75,
	0xaf, 0x1e, 0x7b, 0x11, 0x1e, 0x81, 0xff, 0x84, 0x34, 0x2c, 0xd2, 0x38, 0x12, 0xdc, 0x28, 0x91,
	0xa6, 0x54, 0xe9, 0xa9, 0xe5, 0xc0, 0x59, 0x90, 0x9f, 0x38, 0x9e, 0x0d, 0x14, 0xee, 0x43, 0xf0,
	0x4f, 0xac, 0xe5, 0xad, 0xd6, 0x43, 0x67, 0xdd, 0x89, 0xb5, 0xbc, 0xc5, 0xd7, 0x01, 0x9b, 0xd4,
	0x0c, 0x30, 0xe3, 0x26, 0x99, 0x1a, 0x9e, 0x38, 0x43, 0x83, 0x9a, 0xc1, 0x29, 0x37, 0x49, 0x31,
	0xf9, 0x3d, 0x00, 0xdb, 0x42, 0x1a, 0x8d, 0x29, 0xc1, 0x17, 0xa3, 0xf4, 0x32, 0x1f, 0x47, 0x4f,
	0x4b, 0x3b, 0xbf, 0x4d, 0x1b, 0x7f, 0x42, 0x5e, 0x8e, 0xd2, 0x4b, 0x0f, 0x75, 0x93, 0x69, 0x24,
	0x63, 0xcb, 0xf4, 0xec, 0x2e, 0x98, 0xde, 0xbb, 0x6c, 0xf8, 0x2d, 0x00, 0xd0, 0x1e, 0xc5, 0xb5,
	0x35, 0x3d, 0x2f, 0x0d, 0x69, 0x23, 0xd6, 0x72, 0x61, 0x4b, 0xd7, 0x89, 0xf2, 0x25, 0xbd, 0xb8,
	0x03, 0xa2, 0x7c, 0x47, 0x9f, 0x40, 0xe8, 0x5e, 0x07, 0xb6, 0x8d, 0x8e, 0x5a, 0x95, 0xbf, 0xc6,
	0xb1, 0xf0, 0x06, 0xe9, 0x57, 0x6d, 0x69, 0x7f, 0xf7, 0x17, 0xab, 0xee, 0xa9, 0xb6, 0xff, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0xc0, 0x2a, 0xbf, 0xeb, 0xc2, 0x09, 0x00, 0x00,
}