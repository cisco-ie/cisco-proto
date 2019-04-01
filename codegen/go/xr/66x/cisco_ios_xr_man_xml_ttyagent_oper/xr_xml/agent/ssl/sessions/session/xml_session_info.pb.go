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
// source: xml_session_info.proto

package cisco_ios_xr_man_xml_ttyagent_oper_xr_xml_agent_ssl_sessions_session

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

type XmlSessionInfo_KEYS struct {
	SessionId            uint32   `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XmlSessionInfo_KEYS) Reset()         { *m = XmlSessionInfo_KEYS{} }
func (m *XmlSessionInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*XmlSessionInfo_KEYS) ProtoMessage()    {}
func (*XmlSessionInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0021efa8da6c904f, []int{0}
}

func (m *XmlSessionInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XmlSessionInfo_KEYS.Unmarshal(m, b)
}
func (m *XmlSessionInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XmlSessionInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *XmlSessionInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XmlSessionInfo_KEYS.Merge(m, src)
}
func (m *XmlSessionInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_XmlSessionInfo_KEYS.Size(m)
}
func (m *XmlSessionInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_XmlSessionInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_XmlSessionInfo_KEYS proto.InternalMessageInfo

func (m *XmlSessionInfo_KEYS) GetSessionId() uint32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type XmlSessionInfo struct {
	Username             string   `protobuf:"bytes,50,opt,name=username,proto3" json:"username,omitempty"`
	State                string   `protobuf:"bytes,51,opt,name=state,proto3" json:"state,omitempty"`
	ClientAddress        string   `protobuf:"bytes,52,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
	ClientPort           uint32   `protobuf:"varint,53,opt,name=client_port,json=clientPort,proto3" json:"client_port,omitempty"`
	ConfigSessionId      string   `protobuf:"bytes,54,opt,name=config_session_id,json=configSessionId,proto3" json:"config_session_id,omitempty"`
	AdminConfigSessionId string   `protobuf:"bytes,55,opt,name=admin_config_session_id,json=adminConfigSessionId,proto3" json:"admin_config_session_id,omitempty"`
	AlarmNotification    string   `protobuf:"bytes,56,opt,name=alarm_notification,json=alarmNotification,proto3" json:"alarm_notification,omitempty"`
	VrfName              string   `protobuf:"bytes,57,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	StartTime            uint32   `protobuf:"varint,58,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	ElapsedTime          uint32   `protobuf:"varint,59,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	LastStateChange      uint32   `protobuf:"varint,60,opt,name=last_state_change,json=lastStateChange,proto3" json:"last_state_change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XmlSessionInfo) Reset()         { *m = XmlSessionInfo{} }
func (m *XmlSessionInfo) String() string { return proto.CompactTextString(m) }
func (*XmlSessionInfo) ProtoMessage()    {}
func (*XmlSessionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0021efa8da6c904f, []int{1}
}

func (m *XmlSessionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XmlSessionInfo.Unmarshal(m, b)
}
func (m *XmlSessionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XmlSessionInfo.Marshal(b, m, deterministic)
}
func (m *XmlSessionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XmlSessionInfo.Merge(m, src)
}
func (m *XmlSessionInfo) XXX_Size() int {
	return xxx_messageInfo_XmlSessionInfo.Size(m)
}
func (m *XmlSessionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_XmlSessionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_XmlSessionInfo proto.InternalMessageInfo

func (m *XmlSessionInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *XmlSessionInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *XmlSessionInfo) GetClientAddress() string {
	if m != nil {
		return m.ClientAddress
	}
	return ""
}

func (m *XmlSessionInfo) GetClientPort() uint32 {
	if m != nil {
		return m.ClientPort
	}
	return 0
}

func (m *XmlSessionInfo) GetConfigSessionId() string {
	if m != nil {
		return m.ConfigSessionId
	}
	return ""
}

func (m *XmlSessionInfo) GetAdminConfigSessionId() string {
	if m != nil {
		return m.AdminConfigSessionId
	}
	return ""
}

func (m *XmlSessionInfo) GetAlarmNotification() string {
	if m != nil {
		return m.AlarmNotification
	}
	return ""
}

func (m *XmlSessionInfo) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *XmlSessionInfo) GetStartTime() uint32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *XmlSessionInfo) GetElapsedTime() uint32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func (m *XmlSessionInfo) GetLastStateChange() uint32 {
	if m != nil {
		return m.LastStateChange
	}
	return 0
}

func init() {
	proto.RegisterType((*XmlSessionInfo_KEYS)(nil), "cisco_ios_xr_man_xml_ttyagent_oper.xr_xml.agent.ssl.sessions.session.xml_session_info_KEYS")
	proto.RegisterType((*XmlSessionInfo)(nil), "cisco_ios_xr_man_xml_ttyagent_oper.xr_xml.agent.ssl.sessions.session.xml_session_info")
}

func init() { proto.RegisterFile("xml_session_info.proto", fileDescriptor_0021efa8da6c904f) }

var fileDescriptor_0021efa8da6c904f = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x5f, 0x4b, 0x02, 0x41,
	0x14, 0xc5, 0x91, 0xa8, 0xf4, 0x9a, 0x99, 0x83, 0xd5, 0x14, 0x44, 0x26, 0x04, 0x22, 0xb4, 0x0f,
	0x99, 0xf6, 0xf7, 0x25, 0xac, 0x87, 0x08, 0x24, 0xb4, 0x97, 0x9e, 0x86, 0x69, 0x77, 0xd6, 0x06,
	0x76, 0x67, 0x96, 0x99, 0x49, 0xec, 0x13, 0xf5, 0x35, 0x63, 0xee, 0xae, 0x22, 0xf6, 0xb6, 0xf3,
	0x3b, 0xe7, 0x5c, 0x0e, 0x87, 0x85, 0x83, 0x79, 0x9a, 0x30, 0x2b, 0xac, 0x95, 0x5a, 0x31, 0xa9,
	0x62, 0x1d, 0x64, 0x46, 0x3b, 0x4d, 0x9e, 0x42, 0x69, 0x43, 0xcd, 0xa4, 0xb6, 0x6c, 0x6e, 0x58,
	0xca, 0x15, 0xf3, 0x46, 0xe7, 0x7e, 0xf8, 0x54, 0x28, 0xc7, 0x74, 0x26, 0x4c, 0x30, 0x37, 0x1e,
	0x06, 0x48, 0x02, 0x6b, 0x93, 0xa0, 0xb8, 0x63, 0x17, 0x1f, 0xed, 0x01, 0xec, 0xaf, 0xdf, 0x67,
	0xaf, 0xcf, 0x1f, 0x13, 0x72, 0x02, 0xb0, 0x84, 0x11, 0x2d, 0xb5, 0x4a, 0x9d, 0xda, 0xb8, 0x52,
	0x90, 0x97, 0xa8, 0xfd, 0xbb, 0x01, 0x7b, 0xeb, 0x41, 0x72, 0x0c, 0xe5, 0x6f, 0x2b, 0x8c, 0xe2,
	0xa9, 0xa0, 0x97, 0xad, 0x52, 0xa7, 0x32, 0x5e, 0xbe, 0x49, 0x13, 0x36, 0xad, 0xe3, 0x4e, 0xd0,
	0x1e, 0x0a, 0xf9, 0x83, 0x9c, 0xc3, 0x6e, 0x98, 0x48, 0xdf, 0x97, 0x47, 0x91, 0x11, 0xd6, 0xd2,
	0x2b, 0x94, 0x6b, 0x39, 0x7d, 0xcc, 0x21, 0x39, 0x85, 0x6a, 0x61, 0xcb, 0xb4, 0x71, 0xb4, 0x8f,
	0x6d, 0x20, 0x47, 0x6f, 0xda, 0x38, 0xd2, 0x85, 0x46, 0xa8, 0x55, 0x2c, 0xa7, 0x6c, 0xa5, 0xf4,
	0x00, 0x4f, 0xd5, 0x73, 0x61, 0xb2, 0xa8, 0x4e, 0xfa, 0x70, 0xc8, 0xa3, 0x54, 0x2a, 0xf6, 0x3f,
	0x71, 0x8d, 0x89, 0x26, 0xca, 0xc3, 0xb5, 0xd8, 0x05, 0x10, 0x9e, 0x70, 0x93, 0x32, 0xa5, 0x9d,
	0x8c, 0x65, 0xc8, 0x9d, 0xd4, 0x8a, 0xde, 0x60, 0xa2, 0x81, 0xca, 0x68, 0x45, 0x20, 0x47, 0x50,
	0x9e, 0x99, 0x98, 0xe1, 0x16, 0xb7, 0x68, 0xda, 0x9e, 0x99, 0x78, 0xe4, 0xa7, 0xf0, 0xd3, 0x3a,
	0x6e, 0x1c, 0x73, 0x32, 0x15, 0xf4, 0xae, 0x98, 0xd6, 0x93, 0x77, 0x99, 0x0a, 0x72, 0x06, 0x3b,
	0x22, 0xe1, 0x99, 0x15, 0x51, 0x6e, 0xb8, 0x47, 0x43, 0xb5, 0x60, 0x68, 0xe9, 0x42, 0x23, 0xe1,
	0xd6, 0x31, 0x1c, 0x91, 0x85, 0x5f, 0x5c, 0x4d, 0x05, 0x7d, 0x40, 0x5f, 0xdd, 0x0b, 0x13, 0xcf,
	0x87, 0x88, 0x3f, 0xb7, 0xf0, 0x77, 0xe9, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0xda, 0x59, 0xfd,
	0xe9, 0x48, 0x02, 0x00, 0x00,
}
