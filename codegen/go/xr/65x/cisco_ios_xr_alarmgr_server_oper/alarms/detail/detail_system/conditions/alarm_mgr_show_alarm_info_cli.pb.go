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

package cisco_ios_xr_alarmgr_server_oper_alarms_detail_detail_system_conditions

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
	proto.RegisterType((*AlarmMgrShowAlarmInfoCli_KEYS)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.conditions.alarm_mgr_show_alarm_info_cli_KEYS")
	proto.RegisterType((*AlarmOtn)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.conditions.alarm_otn")
	proto.RegisterType((*AlarmTca)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.conditions.alarm_tca")
	proto.RegisterType((*AlarmMgrShowAlarmData)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.conditions.alarm_mgr_show_alarm_data")
	proto.RegisterType((*AlarmMgrShowAlarmInfoCli)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.conditions.alarm_mgr_show_alarm_info_cli")
}

func init() {
	proto.RegisterFile("alarm_mgr_show_alarm_info_cli.proto", fileDescriptor_8521c1e466bbc1bb)
}

var fileDescriptor_8521c1e466bbc1bb = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xb5, 0xa4, 0x4d, 0x1b, 0x27, 0x6d, 0x53, 0xb7, 0x20, 0x17, 0x51, 0xb1, 0xa4, 0x48,
	0x8d, 0x04, 0x0a, 0x52, 0x79, 0x82, 0x1e, 0x10, 0xaa, 0x90, 0x38, 0x6c, 0x2b, 0x24, 0xb8, 0x58,
	0xae, 0x3d, 0xd9, 0x5a, 0xec, 0xda, 0x2b, 0x7b, 0xb6, 0x10, 0x6e, 0xbc, 0x03, 0x6f, 0xc0, 0x8b,
	0x22, 0xdb, 0x9b, 0x34, 0x87, 0xaa, 0xa7, 0x72, 0xca, 0xce, 0xf7, 0xff, 0xfa, 0x67, 0x66, 0x63,
	0x2f, 0x39, 0x11, 0x95, 0x70, 0x35, 0xaf, 0x4b, 0xc7, 0xfd, 0x8d, 0xfd, 0xc1, 0x53, 0xa9, 0xcd,
	0xdc, 0x72, 0x59, 0xe9, 0x59, 0xe3, 0x2c, 0x5a, 0xfa, 0x51, 0x6a, 0x2f, 0x2d, 0xd7, 0xd6, 0xf3,
	0x9f, 0x2e, 0x59, 0x82, 0x1f, 0xdc, 0x2d, 0x38, 0x6e, 0x1b, 0x70, 0xb3, 0xc8, 0xfc, 0x4c, 0x01,
	0x0a, 0x5d, 0x75, 0x3f, 0xdc, 0x2f, 0x3c, 0x42, 0x3d, 0x93, 0xd6, 0x28, 0x8d, 0xda, 0x1a, 0x3f,
	0x79, 0x4d, 0x26, 0x0f, 0xf6, 0xe3, 0x9f, 0x3e, 0x7c, 0xbd, 0x9c, 0x7c, 0x23, 0x83, 0x84, 0x2d,
	0x1a, 0xfa, 0x82, 0x0c, 0x94, 0x76, 0x20, 0x43, 0x00, 0xcb, 0xf2, 0x6c, 0x3a, 0x28, 0xee, 0x00,
	0x7d, 0x47, 0x0e, 0x8c, 0x45, 0x3d, 0xd7, 0x52, 0x84, 0x9a, 0x7b, 0xdb, 0x3a, 0x09, 0xec, 0x49,
	0xf4, 0xd1, 0x75, 0xe9, 0x32, 0x2a, 0x93, 0x5f, 0xcb, 0x6c, 0x94, 0x82, 0x9e, 0x92, 0x3d, 0xbc,
	0x71, 0xe0, 0x6f, 0x6c, 0xa5, 0xf8, 0xad, 0xa8, 0x5a, 0xe8, 0x3a, 0xec, 0xae, 0xf0, 0x97, 0x40,
	0xe9, 0x09, 0xd9, 0x91, 0xad, 0x73, 0x60, 0xb0, 0xb3, 0xa5, 0x06, 0xa3, 0x0e, 0x26, 0xd3, 0x4b,
	0x32, 0xbc, 0x6e, 0xe5, 0x77, 0x40, 0x8e, 0x8b, 0x06, 0x58, 0x2f, 0x5a, 0x48, 0x42, 0x57, 0x8b,
	0x06, 0x26, 0x7f, 0xfa, 0xe4, 0xe8, 0xde, 0xf5, 0x95, 0x40, 0x41, 0x73, 0x32, 0x54, 0xe0, 0xa5,
	0xd3, 0xcd, 0xda, 0xaa, 0xeb, 0x88, 0x3e, 0x27, 0xdb, 0x95, 0x4d, 0xdb, 0x74, 0x03, 0xac, 0x6a,
	0x3a, 0x26, 0x3d, 0xa1, 0x55, 0xd7, 0x34, 0x3c, 0x06, 0x82, 0xa2, 0x64, 0x1b, 0x89, 0xa0, 0x28,
	0xe9, 0x33, 0xd2, 0xaf, 0xad, 0x6a, 0x2b, 0x60, 0x9b, 0x11, 0x76, 0x55, 0x70, 0x82, 0x56, 0xac,
	0x9f, 0x9c, 0xa0, 0x15, 0x7d, 0x4b, 0xa8, 0x83, 0xc6, 0x3a, 0xd4, 0xa6, 0xe4, 0xa2, 0x0c, 0x7b,
	0x6b, 0xc5, 0xb6, 0xf2, 0x6c, 0xba, 0x53, 0x8c, 0x57, 0xca, 0x79, 0x10, 0x2e, 0x14, 0x7d, 0x45,
	0x46, 0x0d, 0x18, 0x15, 0xbc, 0x7e, 0x61, 0x24, 0xdb, 0xce, 0xb3, 0xe9, 0x76, 0x31, 0xec, 0xd8,
	0xe5, 0xc2, 0xc8, 0x30, 0xba, 0x87, 0x5b, 0x70, 0x1a, 0x17, 0x6c, 0x90, 0x46, 0x5f, 0xd6, 0x61,
	0x2c, 0x8f, 0x02, 0x5b, 0xcf, 0x48, 0x1a, 0x2b, 0x55, 0xf4, 0x90, 0x6c, 0x96, 0xce, 0xb6, 0x0d,
	0x1b, 0x46, 0x9c, 0x0a, 0x7a, 0x14, 0x92, 0x90, 0xa3, 0xae, 0x81, 0x8d, 0xa2, 0xb0, 0xe5, 0x01,
	0xaf, 0x74, 0x1d, 0xff, 0xa5, 0xa5, 0xe4, 0x51, 0xd4, 0x0d, 0xdb, 0xc9, 0xb3, 0xe9, 0x46, 0x31,
	0xea, 0xf4, 0xc8, 0xe8, 0x31, 0x21, 0xb2, 0x02, 0xe1, 0x52, 0xc2, 0x6e, 0x3a, 0x50, 0x91, 0xc4,
	0x8c, 0x53, 0xb2, 0x77, 0x27, 0xa7, 0x94, 0xbd, 0x98, 0xb2, 0xbb, 0xf2, 0xa4, 0x9c, 0x37, 0x64,
	0x3f, 0x5c, 0x00, 0x2d, 0x81, 0x8b, 0xf9, 0x3c, 0x1c, 0x47, 0x53, 0xb2, 0x71, 0x8c, 0x1b, 0x77,
	0xc2, 0xf9, 0x92, 0x53, 0x45, 0x7a, 0x16, 0x0d, 0xdb, 0xcf, 0xb3, 0xe9, 0xf0, 0xac, 0x98, 0x3d,
	0xd2, 0x75, 0x9a, 0xad, 0x6e, 0x49, 0x11, 0xe2, 0x43, 0x17, 0x94, 0x82, 0xd1, 0xff, 0xd2, 0x05,
	0xa5, 0x28, 0x42, 0x3c, 0xa5, 0x64, 0x23, 0x9e, 0xef, 0x83, 0xb8, 0x6b, 0x7c, 0x0e, 0x97, 0x54,
	0x1b, 0x04, 0x37, 0x17, 0x12, 0xd8, 0x61, 0x7a, 0xa7, 0x2b, 0x10, 0x5e, 0x79, 0xca, 0x30, 0xa2,
	0x06, 0xf6, 0x34, 0xc9, 0x91, 0x7c, 0x16, 0x35, 0x4c, 0xfe, 0x66, 0xe4, 0xf8, 0xc1, 0xaf, 0x02,
	0xfd, 0x9d, 0x2d, 0x13, 0x02, 0x62, 0x67, 0x79, 0x6f, 0x3a, 0x3c, 0xbb, 0x7e, 0xe4, 0x05, 0xef,
	0xb9, 0x93, 0xdd, 0x94, 0x17, 0x66, 0x6e, 0xaf, 0xfb, 0xf1, 0x53, 0xf8, 0xfe, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb6, 0xf4, 0x23, 0xa6, 0x31, 0x05, 0x00, 0x00,
}
