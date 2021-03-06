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
// source: alarm_mgr_show_alarm_info_cli.proto

package cisco_ios_xr_alarmgr_server_oper_alarms_detail_detail_card_detail_locations_detail_location_history

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

type AlarmMgrShowAlarmInfoCli_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmMgrShowAlarmInfoCli_KEYS) Reset()         { *m = AlarmMgrShowAlarmInfoCli_KEYS{} }
func (m *AlarmMgrShowAlarmInfoCli_KEYS) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmInfoCli_KEYS) ProtoMessage()    {}
func (*AlarmMgrShowAlarmInfoCli_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8521c1e466bbc1bb, []int{0}
}

func (m *AlarmMgrShowAlarmInfoCli_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmInfoCli_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS.Marshal(b, m, deterministic)
}
func (m *AlarmMgrShowAlarmInfoCli_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS.Merge(m, src)
}
func (m *AlarmMgrShowAlarmInfoCli_KEYS) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS.Size(m)
}
func (m *AlarmMgrShowAlarmInfoCli_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS proto.InternalMessageInfo

func (m *AlarmMgrShowAlarmInfoCli_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type AlarmOtn struct {
	Direction            string   `protobuf:"bytes,1,opt,name=direction,proto3" json:"direction,omitempty"`
	NotificationSource   string   `protobuf:"bytes,2,opt,name=notification_source,json=notificationSource,proto3" json:"notification_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmOtn) Reset()         { *m = AlarmOtn{} }
func (m *AlarmOtn) String() string { return proto.CompactTextString(m) }
func (*AlarmOtn) ProtoMessage()    {}
func (*AlarmOtn) Descriptor() ([]byte, []int) {
	return fileDescriptor_8521c1e466bbc1bb, []int{1}
}

func (m *AlarmOtn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmOtn.Unmarshal(m, b)
}
func (m *AlarmOtn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmOtn.Marshal(b, m, deterministic)
}
func (m *AlarmOtn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmOtn.Merge(m, src)
}
func (m *AlarmOtn) XXX_Size() int {
	return xxx_messageInfo_AlarmOtn.Size(m)
}
func (m *AlarmOtn) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmOtn.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmOtn proto.InternalMessageInfo

func (m *AlarmOtn) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *AlarmOtn) GetNotificationSource() string {
	if m != nil {
		return m.NotificationSource
	}
	return ""
}

type AlarmTca struct {
	ThresholdValue       string   `protobuf:"bytes,1,opt,name=threshold_value,json=thresholdValue,proto3" json:"threshold_value,omitempty"`
	CurrentValue         string   `protobuf:"bytes,2,opt,name=current_value,json=currentValue,proto3" json:"current_value,omitempty"`
	BucketType           string   `protobuf:"bytes,3,opt,name=bucket_type,json=bucketType,proto3" json:"bucket_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmTca) Reset()         { *m = AlarmTca{} }
func (m *AlarmTca) String() string { return proto.CompactTextString(m) }
func (*AlarmTca) ProtoMessage()    {}
func (*AlarmTca) Descriptor() ([]byte, []int) {
	return fileDescriptor_8521c1e466bbc1bb, []int{2}
}

func (m *AlarmTca) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmTca.Unmarshal(m, b)
}
func (m *AlarmTca) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmTca.Marshal(b, m, deterministic)
}
func (m *AlarmTca) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmTca.Merge(m, src)
}
func (m *AlarmTca) XXX_Size() int {
	return xxx_messageInfo_AlarmTca.Size(m)
}
func (m *AlarmTca) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmTca.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmTca proto.InternalMessageInfo

func (m *AlarmTca) GetThresholdValue() string {
	if m != nil {
		return m.ThresholdValue
	}
	return ""
}

func (m *AlarmTca) GetCurrentValue() string {
	if m != nil {
		return m.CurrentValue
	}
	return ""
}

func (m *AlarmTca) GetBucketType() string {
	if m != nil {
		return m.BucketType
	}
	return ""
}

type AlarmMgrShowAlarmData struct {
	Description          string    `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Location             string    `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Aid                  string    `protobuf:"bytes,3,opt,name=aid,proto3" json:"aid,omitempty"`
	Tag                  string    `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	Module               string    `protobuf:"bytes,5,opt,name=module,proto3" json:"module,omitempty"`
	Eid                  string    `protobuf:"bytes,6,opt,name=eid,proto3" json:"eid,omitempty"`
	ReportingAgentId     uint32    `protobuf:"varint,7,opt,name=reporting_agent_id,json=reportingAgentId,proto3" json:"reporting_agent_id,omitempty"`
	PendingSync          bool      `protobuf:"varint,8,opt,name=pending_sync,json=pendingSync,proto3" json:"pending_sync,omitempty"`
	Severity             string    `protobuf:"bytes,9,opt,name=severity,proto3" json:"severity,omitempty"`
	Status               string    `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Group                string    `protobuf:"bytes,11,opt,name=group,proto3" json:"group,omitempty"`
	SetTime              string    `protobuf:"bytes,12,opt,name=set_time,json=setTime,proto3" json:"set_time,omitempty"`
	SetTimestamp         uint64    `protobuf:"varint,13,opt,name=set_timestamp,json=setTimestamp,proto3" json:"set_timestamp,omitempty"`
	ClearTime            string    `protobuf:"bytes,14,opt,name=clear_time,json=clearTime,proto3" json:"clear_time,omitempty"`
	ClearTimestamp       uint64    `protobuf:"varint,15,opt,name=clear_timestamp,json=clearTimestamp,proto3" json:"clear_timestamp,omitempty"`
	ServiceAffecting     string    `protobuf:"bytes,16,opt,name=service_affecting,json=serviceAffecting,proto3" json:"service_affecting,omitempty"`
	Otn                  *AlarmOtn `protobuf:"bytes,17,opt,name=otn,proto3" json:"otn,omitempty"`
	Tca                  *AlarmTca `protobuf:"bytes,18,opt,name=tca,proto3" json:"tca,omitempty"`
	Type                 string    `protobuf:"bytes,19,opt,name=type,proto3" json:"type,omitempty"`
	Interface            string    `protobuf:"bytes,20,opt,name=interface,proto3" json:"interface,omitempty"`
	AlarmName            string    `protobuf:"bytes,21,opt,name=alarm_name,json=alarmName,proto3" json:"alarm_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AlarmMgrShowAlarmData) Reset()         { *m = AlarmMgrShowAlarmData{} }
func (m *AlarmMgrShowAlarmData) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmData) ProtoMessage()    {}
func (*AlarmMgrShowAlarmData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8521c1e466bbc1bb, []int{3}
}

func (m *AlarmMgrShowAlarmData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmData.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmData.Marshal(b, m, deterministic)
}
func (m *AlarmMgrShowAlarmData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmData.Merge(m, src)
}
func (m *AlarmMgrShowAlarmData) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmData.Size(m)
}
func (m *AlarmMgrShowAlarmData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmData proto.InternalMessageInfo

func (m *AlarmMgrShowAlarmData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetAid() string {
	if m != nil {
		return m.Aid
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetEid() string {
	if m != nil {
		return m.Eid
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetReportingAgentId() uint32 {
	if m != nil {
		return m.ReportingAgentId
	}
	return 0
}

func (m *AlarmMgrShowAlarmData) GetPendingSync() bool {
	if m != nil {
		return m.PendingSync
	}
	return false
}

func (m *AlarmMgrShowAlarmData) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetSetTime() string {
	if m != nil {
		return m.SetTime
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetSetTimestamp() uint64 {
	if m != nil {
		return m.SetTimestamp
	}
	return 0
}

func (m *AlarmMgrShowAlarmData) GetClearTime() string {
	if m != nil {
		return m.ClearTime
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetClearTimestamp() uint64 {
	if m != nil {
		return m.ClearTimestamp
	}
	return 0
}

func (m *AlarmMgrShowAlarmData) GetServiceAffecting() string {
	if m != nil {
		return m.ServiceAffecting
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetOtn() *AlarmOtn {
	if m != nil {
		return m.Otn
	}
	return nil
}

func (m *AlarmMgrShowAlarmData) GetTca() *AlarmTca {
	if m != nil {
		return m.Tca
	}
	return nil
}

func (m *AlarmMgrShowAlarmData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetAlarmName() string {
	if m != nil {
		return m.AlarmName
	}
	return ""
}

type AlarmMgrShowAlarmInfoCli struct {
	AlarmInfo            []*AlarmMgrShowAlarmData `protobuf:"bytes,50,rep,name=alarm_info,json=alarmInfo,proto3" json:"alarm_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AlarmMgrShowAlarmInfoCli) Reset()         { *m = AlarmMgrShowAlarmInfoCli{} }
func (m *AlarmMgrShowAlarmInfoCli) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmInfoCli) ProtoMessage()    {}
func (*AlarmMgrShowAlarmInfoCli) Descriptor() ([]byte, []int) {
	return fileDescriptor_8521c1e466bbc1bb, []int{4}
}

func (m *AlarmMgrShowAlarmInfoCli) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmInfoCli) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli.Marshal(b, m, deterministic)
}
func (m *AlarmMgrShowAlarmInfoCli) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmInfoCli.Merge(m, src)
}
func (m *AlarmMgrShowAlarmInfoCli) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli.Size(m)
}
func (m *AlarmMgrShowAlarmInfoCli) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmInfoCli.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmInfoCli proto.InternalMessageInfo

func (m *AlarmMgrShowAlarmInfoCli) GetAlarmInfo() []*AlarmMgrShowAlarmData {
	if m != nil {
		return m.AlarmInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*AlarmMgrShowAlarmInfoCli_KEYS)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_card.detail_locations.detail_location.history.alarm_mgr_show_alarm_info_cli_KEYS")
	proto.RegisterType((*AlarmOtn)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_card.detail_locations.detail_location.history.alarm_otn")
	proto.RegisterType((*AlarmTca)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_card.detail_locations.detail_location.history.alarm_tca")
	proto.RegisterType((*AlarmMgrShowAlarmData)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_card.detail_locations.detail_location.history.alarm_mgr_show_alarm_data")
	proto.RegisterType((*AlarmMgrShowAlarmInfoCli)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_card.detail_locations.detail_location.history.alarm_mgr_show_alarm_info_cli")
}

func init() {
	proto.RegisterFile("alarm_mgr_show_alarm_info_cli.proto", fileDescriptor_8521c1e466bbc1bb)
}

var fileDescriptor_8521c1e466bbc1bb = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xc1, 0x6e, 0x13, 0x3b,
	0x14, 0xd5, 0xbc, 0xb4, 0x69, 0xe3, 0xa4, 0x6d, 0xea, 0xf6, 0xbd, 0xe7, 0x22, 0x2a, 0x86, 0x74,
	0xd1, 0x48, 0xa0, 0x20, 0x95, 0x35, 0x8b, 0x2e, 0x58, 0x54, 0x48, 0x2c, 0xd2, 0x0a, 0x09, 0x36,
	0x96, 0x6b, 0xdf, 0x4c, 0x2c, 0x66, 0xec, 0x91, 0x7d, 0xa7, 0x10, 0x16, 0x88, 0x1f, 0xe1, 0x5b,
	0xf8, 0x03, 0xbe, 0x09, 0xd9, 0x9e, 0x49, 0x2b, 0x54, 0x75, 0xd9, 0xd5, 0xf8, 0x9e, 0x73, 0x74,
	0xee, 0xb1, 0x3d, 0xd7, 0xe4, 0x44, 0x94, 0xc2, 0x55, 0xbc, 0x2a, 0x1c, 0xf7, 0x4b, 0xfb, 0x85,
	0xa7, 0x52, 0x9b, 0x85, 0xe5, 0xb2, 0xd4, 0xb3, 0xda, 0x59, 0xb4, 0x54, 0x4a, 0xed, 0xa5, 0xe5,
	0xda, 0x7a, 0xfe, 0xd5, 0x25, 0x49, 0xd0, 0x83, 0xbb, 0x01, 0xc7, 0x6d, 0x0d, 0x6e, 0x16, 0x31,
	0x3f, 0x53, 0x80, 0x42, 0x97, 0xed, 0x87, 0x4b, 0xe1, 0x54, 0xb7, 0x2e, 0xad, 0x14, 0xa8, 0xad,
	0xf1, 0x7f, 0x03, 0xb3, 0xa5, 0xf6, 0x68, 0xdd, 0x6a, 0xf2, 0x86, 0x4c, 0x1e, 0xcc, 0xc2, 0xdf,
	0xbd, 0xfd, 0x78, 0x49, 0xff, 0x27, 0x5b, 0xc6, 0x2a, 0xe0, 0x5a, 0xb1, 0x2c, 0xcf, 0xa6, 0x83,
	0x79, 0x3f, 0x94, 0x17, 0x6a, 0xf2, 0x89, 0x0c, 0x92, 0xde, 0xa2, 0xa1, 0x4f, 0xc9, 0x40, 0x69,
	0x07, 0x32, 0x34, 0x68, 0x75, 0xb7, 0x00, 0x7d, 0x45, 0x0e, 0x8c, 0x45, 0xbd, 0xd0, 0x29, 0x01,
	0xf7, 0xb6, 0x71, 0x12, 0xd8, 0x3f, 0x51, 0x47, 0xef, 0x52, 0x97, 0x91, 0x99, 0x7c, 0xeb, 0xbc,
	0x51, 0x0a, 0x7a, 0x4a, 0xf6, 0x70, 0xe9, 0xc0, 0x2f, 0x6d, 0xa9, 0xf8, 0x8d, 0x28, 0x1b, 0x68,
	0x3b, 0xec, 0xae, 0xe1, 0x0f, 0x01, 0xa5, 0x27, 0x64, 0x47, 0x36, 0xce, 0x81, 0xc1, 0x56, 0x96,
	0x1a, 0x8c, 0x5a, 0x30, 0x89, 0x9e, 0x91, 0xe1, 0x75, 0x23, 0x3f, 0x03, 0x72, 0x5c, 0xd5, 0xc0,
	0x7a, 0x51, 0x42, 0x12, 0x74, 0xb5, 0xaa, 0x61, 0xf2, 0xbb, 0x4f, 0x8e, 0xee, 0x3d, 0x17, 0x25,
	0x50, 0xd0, 0x9c, 0x0c, 0x15, 0x78, 0xe9, 0x74, 0x7d, 0x67, 0xab, 0x77, 0x21, 0xfa, 0x84, 0x6c,
	0x77, 0x47, 0xdd, 0x06, 0x58, 0xd7, 0x74, 0x4c, 0x7a, 0x42, 0xab, 0xb6, 0x69, 0x58, 0x06, 0x04,
	0x45, 0xc1, 0x36, 0x12, 0x82, 0xa2, 0xa0, 0xff, 0x91, 0x7e, 0x65, 0x55, 0x53, 0x02, 0xdb, 0x4c,
	0xe7, 0x9d, 0xaa, 0xa0, 0x04, 0xad, 0x58, 0x3f, 0x29, 0x41, 0x2b, 0xfa, 0x92, 0x50, 0x07, 0xb5,
	0x75, 0xa8, 0x4d, 0xc1, 0x45, 0x11, 0xf6, 0xad, 0x15, 0xdb, 0xca, 0xb3, 0xe9, 0xce, 0x7c, 0xbc,
	0x66, 0xce, 0x03, 0x71, 0xa1, 0xe8, 0x73, 0x32, 0xaa, 0xc1, 0xa8, 0xa0, 0xf5, 0x2b, 0x23, 0xd9,
	0x76, 0x9e, 0x4d, 0xb7, 0xe7, 0xc3, 0x16, 0xbb, 0x5c, 0x19, 0x19, 0xa2, 0x7b, 0xb8, 0x01, 0xa7,
	0x71, 0xc5, 0x06, 0x29, 0x7a, 0x57, 0x87, 0x58, 0x1e, 0x05, 0x36, 0x9e, 0x91, 0x14, 0x2b, 0x55,
	0xf4, 0x90, 0x6c, 0x16, 0xce, 0x36, 0x35, 0x1b, 0x46, 0x38, 0x15, 0xf4, 0x28, 0x38, 0x21, 0x47,
	0x5d, 0x01, 0x1b, 0x45, 0x62, 0xcb, 0x03, 0x5e, 0xe9, 0x2a, 0xde, 0x52, 0x47, 0x79, 0x14, 0x55,
	0xcd, 0x76, 0xf2, 0x6c, 0xba, 0x31, 0x1f, 0xb5, 0x7c, 0xc4, 0xe8, 0x31, 0x21, 0xb2, 0x04, 0xe1,
	0x92, 0xc3, 0x6e, 0xfa, 0xa1, 0x22, 0x12, 0x3d, 0x4e, 0xc9, 0xde, 0x2d, 0x9d, 0x5c, 0xf6, 0xa2,
	0xcb, 0xee, 0x5a, 0x93, 0x7c, 0x5e, 0x90, 0xfd, 0x30, 0x35, 0x5a, 0x02, 0x17, 0x8b, 0x45, 0xf8,
	0x1d, 0x4d, 0xc1, 0xc6, 0xd1, 0x6e, 0xdc, 0x12, 0xe7, 0x1d, 0x4e, 0x7f, 0x64, 0xa4, 0x67, 0xd1,
	0xb0, 0xfd, 0x3c, 0x9b, 0x0e, 0xcf, 0xcc, 0xec, 0x11, 0x86, 0x70, 0xb6, 0x1e, 0xa1, 0x79, 0x68,
	0x1d, 0x23, 0xa0, 0x14, 0x8c, 0x3e, 0x7a, 0x04, 0x94, 0x62, 0x1e, 0x5a, 0x53, 0x4a, 0x36, 0xe2,
	0x64, 0x1c, 0xc4, 0x53, 0x8a, 0xeb, 0x30, 0xde, 0xda, 0x20, 0xb8, 0x85, 0x90, 0xc0, 0x0e, 0xd3,
	0x6d, 0xac, 0x81, 0x70, 0x59, 0xc9, 0xc3, 0x88, 0x0a, 0xd8, 0xbf, 0x89, 0x8e, 0xc8, 0x7b, 0x51,
	0xc1, 0xe4, 0x57, 0x46, 0x8e, 0x1f, 0x7c, 0x68, 0xe8, 0xcf, 0xac, 0x73, 0x08, 0x10, 0x3b, 0xcb,
	0x7b, 0xd3, 0xe1, 0xd9, 0xf7, 0x47, 0xdc, 0xfc, 0x3d, 0x93, 0xde, 0xee, 0xe0, 0xc2, 0x2c, 0xec,
	0x75, 0x3f, 0xbe, 0xca, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x07, 0x22, 0x29, 0xbc,
	0x05, 0x00, 0x00,
}
