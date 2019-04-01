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
// source: diag_eeprom.proto

package cisco_ios_xr_sdr_invmgr_diag_oper_diag_racks_rack_chassis

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

type DiagEeprom_KEYS struct {
	RackName             string   `protobuf:"bytes,1,opt,name=rack_name,json=rackName,proto3" json:"rack_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiagEeprom_KEYS) Reset()         { *m = DiagEeprom_KEYS{} }
func (m *DiagEeprom_KEYS) String() string { return proto.CompactTextString(m) }
func (*DiagEeprom_KEYS) ProtoMessage()    {}
func (*DiagEeprom_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4ed4664a25087f, []int{0}
}

func (m *DiagEeprom_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiagEeprom_KEYS.Unmarshal(m, b)
}
func (m *DiagEeprom_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiagEeprom_KEYS.Marshal(b, m, deterministic)
}
func (m *DiagEeprom_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiagEeprom_KEYS.Merge(m, src)
}
func (m *DiagEeprom_KEYS) XXX_Size() int {
	return xxx_messageInfo_DiagEeprom_KEYS.Size(m)
}
func (m *DiagEeprom_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_DiagEeprom_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_DiagEeprom_KEYS proto.InternalMessageInfo

func (m *DiagEeprom_KEYS) GetRackName() string {
	if m != nil {
		return m.RackName
	}
	return ""
}

type RmaDetail struct {
	TestHistory          string   `protobuf:"bytes,1,opt,name=test_history,json=testHistory,proto3" json:"test_history,omitempty"`
	RmaNumber            string   `protobuf:"bytes,2,opt,name=rma_number,json=rmaNumber,proto3" json:"rma_number,omitempty"`
	RmaHistory           string   `protobuf:"bytes,3,opt,name=rma_history,json=rmaHistory,proto3" json:"rma_history,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RmaDetail) Reset()         { *m = RmaDetail{} }
func (m *RmaDetail) String() string { return proto.CompactTextString(m) }
func (*RmaDetail) ProtoMessage()    {}
func (*RmaDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4ed4664a25087f, []int{1}
}

func (m *RmaDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RmaDetail.Unmarshal(m, b)
}
func (m *RmaDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RmaDetail.Marshal(b, m, deterministic)
}
func (m *RmaDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RmaDetail.Merge(m, src)
}
func (m *RmaDetail) XXX_Size() int {
	return xxx_messageInfo_RmaDetail.Size(m)
}
func (m *RmaDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RmaDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RmaDetail proto.InternalMessageInfo

func (m *RmaDetail) GetTestHistory() string {
	if m != nil {
		return m.TestHistory
	}
	return ""
}

func (m *RmaDetail) GetRmaNumber() string {
	if m != nil {
		return m.RmaNumber
	}
	return ""
}

func (m *RmaDetail) GetRmaHistory() string {
	if m != nil {
		return m.RmaHistory
	}
	return ""
}

type DiagEeprom struct {
	Description          string     `protobuf:"bytes,50,opt,name=description,proto3" json:"description,omitempty"`
	IdpromFormatRev      string     `protobuf:"bytes,51,opt,name=idprom_format_rev,json=idpromFormatRev,proto3" json:"idprom_format_rev,omitempty"`
	ControllerFamily     string     `protobuf:"bytes,52,opt,name=controller_family,json=controllerFamily,proto3" json:"controller_family,omitempty"`
	ControllerType       string     `protobuf:"bytes,53,opt,name=controller_type,json=controllerType,proto3" json:"controller_type,omitempty"`
	Vid                  string     `protobuf:"bytes,54,opt,name=vid,proto3" json:"vid,omitempty"`
	Hwid                 string     `protobuf:"bytes,55,opt,name=hwid,proto3" json:"hwid,omitempty"`
	Pid                  string     `protobuf:"bytes,56,opt,name=pid,proto3" json:"pid,omitempty"`
	UdiDescription       string     `protobuf:"bytes,57,opt,name=udi_description,json=udiDescription,proto3" json:"udi_description,omitempty"`
	UdiName              string     `protobuf:"bytes,58,opt,name=udi_name,json=udiName,proto3" json:"udi_name,omitempty"`
	Clei                 string     `protobuf:"bytes,59,opt,name=clei,proto3" json:"clei,omitempty"`
	Eci                  string     `protobuf:"bytes,60,opt,name=eci,proto3" json:"eci,omitempty"`
	TopAssemPartNum      string     `protobuf:"bytes,61,opt,name=top_assem_part_num,json=topAssemPartNum,proto3" json:"top_assem_part_num,omitempty"`
	TopAssemVid          string     `protobuf:"bytes,62,opt,name=top_assem_vid,json=topAssemVid,proto3" json:"top_assem_vid,omitempty"`
	PcaNum               string     `protobuf:"bytes,63,opt,name=pca_num,json=pcaNum,proto3" json:"pca_num,omitempty"`
	Pcavid               string     `protobuf:"bytes,64,opt,name=pcavid,proto3" json:"pcavid,omitempty"`
	ChassisSid           string     `protobuf:"bytes,65,opt,name=chassis_sid,json=chassisSid,proto3" json:"chassis_sid,omitempty"`
	DevNum1              string     `protobuf:"bytes,66,opt,name=dev_num1,json=devNum1,proto3" json:"dev_num1,omitempty"`
	DevNum2              string     `protobuf:"bytes,67,opt,name=dev_num2,json=devNum2,proto3" json:"dev_num2,omitempty"`
	DevNum3              string     `protobuf:"bytes,68,opt,name=dev_num3,json=devNum3,proto3" json:"dev_num3,omitempty"`
	DevNum4              string     `protobuf:"bytes,69,opt,name=dev_num4,json=devNum4,proto3" json:"dev_num4,omitempty"`
	DevNum5              string     `protobuf:"bytes,70,opt,name=dev_num5,json=devNum5,proto3" json:"dev_num5,omitempty"`
	DevNum6              string     `protobuf:"bytes,71,opt,name=dev_num6,json=devNum6,proto3" json:"dev_num6,omitempty"`
	DevNum7              string     `protobuf:"bytes,72,opt,name=dev_num7,json=devNum7,proto3" json:"dev_num7,omitempty"`
	ManuTestData         string     `protobuf:"bytes,73,opt,name=manu_test_data,json=manuTestData,proto3" json:"manu_test_data,omitempty"`
	Rma                  *RmaDetail `protobuf:"bytes,74,opt,name=rma,proto3" json:"rma,omitempty"`
	AssetId              string     `protobuf:"bytes,75,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	AssetAlias           string     `protobuf:"bytes,76,opt,name=asset_alias,json=assetAlias,proto3" json:"asset_alias,omitempty"`
	BaseMacAddress1      string     `protobuf:"bytes,77,opt,name=base_mac_address1,json=baseMacAddress1,proto3" json:"base_mac_address1,omitempty"`
	MacAddBlkSize1       string     `protobuf:"bytes,78,opt,name=mac_add_blk_size1,json=macAddBlkSize1,proto3" json:"mac_add_blk_size1,omitempty"`
	BaseMacAddress2      string     `protobuf:"bytes,79,opt,name=base_mac_address2,json=baseMacAddress2,proto3" json:"base_mac_address2,omitempty"`
	MacAddBlkSize2       string     `protobuf:"bytes,80,opt,name=mac_add_blk_size2,json=macAddBlkSize2,proto3" json:"mac_add_blk_size2,omitempty"`
	BaseMacAddress3      string     `protobuf:"bytes,81,opt,name=base_mac_address3,json=baseMacAddress3,proto3" json:"base_mac_address3,omitempty"`
	MacAddBlkSize3       string     `protobuf:"bytes,82,opt,name=mac_add_blk_size3,json=macAddBlkSize3,proto3" json:"mac_add_blk_size3,omitempty"`
	BaseMacAddress4      string     `protobuf:"bytes,83,opt,name=base_mac_address4,json=baseMacAddress4,proto3" json:"base_mac_address4,omitempty"`
	MacAddBlkSize4       string     `protobuf:"bytes,84,opt,name=mac_add_blk_size4,json=macAddBlkSize4,proto3" json:"mac_add_blk_size4,omitempty"`
	PcbSerialNum         string     `protobuf:"bytes,85,opt,name=pcb_serial_num,json=pcbSerialNum,proto3" json:"pcb_serial_num,omitempty"`
	PowerSupplyType      string     `protobuf:"bytes,86,opt,name=power_supply_type,json=powerSupplyType,proto3" json:"power_supply_type,omitempty"`
	PowerConsumption     string     `protobuf:"bytes,87,opt,name=power_consumption,json=powerConsumption,proto3" json:"power_consumption,omitempty"`
	BlockSignature       string     `protobuf:"bytes,88,opt,name=block_signature,json=blockSignature,proto3" json:"block_signature,omitempty"`
	BlockVersion         string     `protobuf:"bytes,89,opt,name=block_version,json=blockVersion,proto3" json:"block_version,omitempty"`
	BlockLength          string     `protobuf:"bytes,90,opt,name=block_length,json=blockLength,proto3" json:"block_length,omitempty"`
	BlockChecksum        string     `protobuf:"bytes,91,opt,name=block_checksum,json=blockChecksum,proto3" json:"block_checksum,omitempty"`
	EepromSize           string     `protobuf:"bytes,92,opt,name=eeprom_size,json=eepromSize,proto3" json:"eeprom_size,omitempty"`
	BlockCount           string     `protobuf:"bytes,93,opt,name=block_count,json=blockCount,proto3" json:"block_count,omitempty"`
	FruMajorType         string     `protobuf:"bytes,94,opt,name=fru_major_type,json=fruMajorType,proto3" json:"fru_major_type,omitempty"`
	FruMinorType         string     `protobuf:"bytes,95,opt,name=fru_minor_type,json=fruMinorType,proto3" json:"fru_minor_type,omitempty"`
	OemString            string     `protobuf:"bytes,96,opt,name=oem_string,json=oemString,proto3" json:"oem_string,omitempty"`
	ProductId            string     `protobuf:"bytes,97,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SerialNumber         string     `protobuf:"bytes,98,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	PartNumber           string     `protobuf:"bytes,99,opt,name=part_number,json=partNumber,proto3" json:"part_number,omitempty"`
	PartRevision         string     `protobuf:"bytes,100,opt,name=part_revision,json=partRevision,proto3" json:"part_revision,omitempty"`
	MfgDeviation         string     `protobuf:"bytes,101,opt,name=mfg_deviation,json=mfgDeviation,proto3" json:"mfg_deviation,omitempty"`
	HwVersion            string     `protobuf:"bytes,102,opt,name=hw_version,json=hwVersion,proto3" json:"hw_version,omitempty"`
	MfgBits              string     `protobuf:"bytes,103,opt,name=mfg_bits,json=mfgBits,proto3" json:"mfg_bits,omitempty"`
	EngineerUse          string     `protobuf:"bytes,104,opt,name=engineer_use,json=engineerUse,proto3" json:"engineer_use,omitempty"`
	Snmpoid              string     `protobuf:"bytes,105,opt,name=snmpoid,proto3" json:"snmpoid,omitempty"`
	RmaCode              string     `protobuf:"bytes,106,opt,name=rma_code,json=rmaCode,proto3" json:"rma_code,omitempty"`
	EciAlphaNumber       string     `protobuf:"bytes,107,opt,name=eci_alpha_number,json=eciAlphaNumber,proto3" json:"eci_alpha_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DiagEeprom) Reset()         { *m = DiagEeprom{} }
func (m *DiagEeprom) String() string { return proto.CompactTextString(m) }
func (*DiagEeprom) ProtoMessage()    {}
func (*DiagEeprom) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4ed4664a25087f, []int{2}
}

func (m *DiagEeprom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiagEeprom.Unmarshal(m, b)
}
func (m *DiagEeprom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiagEeprom.Marshal(b, m, deterministic)
}
func (m *DiagEeprom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiagEeprom.Merge(m, src)
}
func (m *DiagEeprom) XXX_Size() int {
	return xxx_messageInfo_DiagEeprom.Size(m)
}
func (m *DiagEeprom) XXX_DiscardUnknown() {
	xxx_messageInfo_DiagEeprom.DiscardUnknown(m)
}

var xxx_messageInfo_DiagEeprom proto.InternalMessageInfo

func (m *DiagEeprom) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DiagEeprom) GetIdpromFormatRev() string {
	if m != nil {
		return m.IdpromFormatRev
	}
	return ""
}

func (m *DiagEeprom) GetControllerFamily() string {
	if m != nil {
		return m.ControllerFamily
	}
	return ""
}

func (m *DiagEeprom) GetControllerType() string {
	if m != nil {
		return m.ControllerType
	}
	return ""
}

func (m *DiagEeprom) GetVid() string {
	if m != nil {
		return m.Vid
	}
	return ""
}

func (m *DiagEeprom) GetHwid() string {
	if m != nil {
		return m.Hwid
	}
	return ""
}

func (m *DiagEeprom) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *DiagEeprom) GetUdiDescription() string {
	if m != nil {
		return m.UdiDescription
	}
	return ""
}

func (m *DiagEeprom) GetUdiName() string {
	if m != nil {
		return m.UdiName
	}
	return ""
}

func (m *DiagEeprom) GetClei() string {
	if m != nil {
		return m.Clei
	}
	return ""
}

func (m *DiagEeprom) GetEci() string {
	if m != nil {
		return m.Eci
	}
	return ""
}

func (m *DiagEeprom) GetTopAssemPartNum() string {
	if m != nil {
		return m.TopAssemPartNum
	}
	return ""
}

func (m *DiagEeprom) GetTopAssemVid() string {
	if m != nil {
		return m.TopAssemVid
	}
	return ""
}

func (m *DiagEeprom) GetPcaNum() string {
	if m != nil {
		return m.PcaNum
	}
	return ""
}

func (m *DiagEeprom) GetPcavid() string {
	if m != nil {
		return m.Pcavid
	}
	return ""
}

func (m *DiagEeprom) GetChassisSid() string {
	if m != nil {
		return m.ChassisSid
	}
	return ""
}

func (m *DiagEeprom) GetDevNum1() string {
	if m != nil {
		return m.DevNum1
	}
	return ""
}

func (m *DiagEeprom) GetDevNum2() string {
	if m != nil {
		return m.DevNum2
	}
	return ""
}

func (m *DiagEeprom) GetDevNum3() string {
	if m != nil {
		return m.DevNum3
	}
	return ""
}

func (m *DiagEeprom) GetDevNum4() string {
	if m != nil {
		return m.DevNum4
	}
	return ""
}

func (m *DiagEeprom) GetDevNum5() string {
	if m != nil {
		return m.DevNum5
	}
	return ""
}

func (m *DiagEeprom) GetDevNum6() string {
	if m != nil {
		return m.DevNum6
	}
	return ""
}

func (m *DiagEeprom) GetDevNum7() string {
	if m != nil {
		return m.DevNum7
	}
	return ""
}

func (m *DiagEeprom) GetManuTestData() string {
	if m != nil {
		return m.ManuTestData
	}
	return ""
}

func (m *DiagEeprom) GetRma() *RmaDetail {
	if m != nil {
		return m.Rma
	}
	return nil
}

func (m *DiagEeprom) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *DiagEeprom) GetAssetAlias() string {
	if m != nil {
		return m.AssetAlias
	}
	return ""
}

func (m *DiagEeprom) GetBaseMacAddress1() string {
	if m != nil {
		return m.BaseMacAddress1
	}
	return ""
}

func (m *DiagEeprom) GetMacAddBlkSize1() string {
	if m != nil {
		return m.MacAddBlkSize1
	}
	return ""
}

func (m *DiagEeprom) GetBaseMacAddress2() string {
	if m != nil {
		return m.BaseMacAddress2
	}
	return ""
}

func (m *DiagEeprom) GetMacAddBlkSize2() string {
	if m != nil {
		return m.MacAddBlkSize2
	}
	return ""
}

func (m *DiagEeprom) GetBaseMacAddress3() string {
	if m != nil {
		return m.BaseMacAddress3
	}
	return ""
}

func (m *DiagEeprom) GetMacAddBlkSize3() string {
	if m != nil {
		return m.MacAddBlkSize3
	}
	return ""
}

func (m *DiagEeprom) GetBaseMacAddress4() string {
	if m != nil {
		return m.BaseMacAddress4
	}
	return ""
}

func (m *DiagEeprom) GetMacAddBlkSize4() string {
	if m != nil {
		return m.MacAddBlkSize4
	}
	return ""
}

func (m *DiagEeprom) GetPcbSerialNum() string {
	if m != nil {
		return m.PcbSerialNum
	}
	return ""
}

func (m *DiagEeprom) GetPowerSupplyType() string {
	if m != nil {
		return m.PowerSupplyType
	}
	return ""
}

func (m *DiagEeprom) GetPowerConsumption() string {
	if m != nil {
		return m.PowerConsumption
	}
	return ""
}

func (m *DiagEeprom) GetBlockSignature() string {
	if m != nil {
		return m.BlockSignature
	}
	return ""
}

func (m *DiagEeprom) GetBlockVersion() string {
	if m != nil {
		return m.BlockVersion
	}
	return ""
}

func (m *DiagEeprom) GetBlockLength() string {
	if m != nil {
		return m.BlockLength
	}
	return ""
}

func (m *DiagEeprom) GetBlockChecksum() string {
	if m != nil {
		return m.BlockChecksum
	}
	return ""
}

func (m *DiagEeprom) GetEepromSize() string {
	if m != nil {
		return m.EepromSize
	}
	return ""
}

func (m *DiagEeprom) GetBlockCount() string {
	if m != nil {
		return m.BlockCount
	}
	return ""
}

func (m *DiagEeprom) GetFruMajorType() string {
	if m != nil {
		return m.FruMajorType
	}
	return ""
}

func (m *DiagEeprom) GetFruMinorType() string {
	if m != nil {
		return m.FruMinorType
	}
	return ""
}

func (m *DiagEeprom) GetOemString() string {
	if m != nil {
		return m.OemString
	}
	return ""
}

func (m *DiagEeprom) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *DiagEeprom) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *DiagEeprom) GetPartNumber() string {
	if m != nil {
		return m.PartNumber
	}
	return ""
}

func (m *DiagEeprom) GetPartRevision() string {
	if m != nil {
		return m.PartRevision
	}
	return ""
}

func (m *DiagEeprom) GetMfgDeviation() string {
	if m != nil {
		return m.MfgDeviation
	}
	return ""
}

func (m *DiagEeprom) GetHwVersion() string {
	if m != nil {
		return m.HwVersion
	}
	return ""
}

func (m *DiagEeprom) GetMfgBits() string {
	if m != nil {
		return m.MfgBits
	}
	return ""
}

func (m *DiagEeprom) GetEngineerUse() string {
	if m != nil {
		return m.EngineerUse
	}
	return ""
}

func (m *DiagEeprom) GetSnmpoid() string {
	if m != nil {
		return m.Snmpoid
	}
	return ""
}

func (m *DiagEeprom) GetRmaCode() string {
	if m != nil {
		return m.RmaCode
	}
	return ""
}

func (m *DiagEeprom) GetEciAlphaNumber() string {
	if m != nil {
		return m.EciAlphaNumber
	}
	return ""
}

func init() {
	proto.RegisterType((*DiagEeprom_KEYS)(nil), "cisco_ios_xr_sdr_invmgr_diag_oper.diag.racks.rack.chassis.diag_eeprom_KEYS")
	proto.RegisterType((*RmaDetail)(nil), "cisco_ios_xr_sdr_invmgr_diag_oper.diag.racks.rack.chassis.rma_detail")
	proto.RegisterType((*DiagEeprom)(nil), "cisco_ios_xr_sdr_invmgr_diag_oper.diag.racks.rack.chassis.diag_eeprom")
}

func init() { proto.RegisterFile("diag_eeprom.proto", fileDescriptor_1b4ed4664a25087f) }

var fileDescriptor_1b4ed4664a25087f = []byte{
	// 1063 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x53, 0x1b, 0xb7,
	0x17, 0x1d, 0x7e, 0xfc, 0x26, 0x80, 0x1c, 0x02, 0xe8, 0xa1, 0x55, 0xa6, 0x93, 0x09, 0x25, 0xed,
	0x94, 0x86, 0x19, 0x77, 0xb0, 0x0d, 0x34, 0xfd, 0xe6, 0xb3, 0xa1, 0x09, 0x34, 0xc5, 0x84, 0x34,
	0xfd, 0x52, 0x65, 0x49, 0xb6, 0x15, 0x56, 0xab, 0x1d, 0x49, 0x6b, 0x4a, 0xff, 0xf3, 0xbe, 0x75,
	0xae, 0xee, 0x1a, 0xef, 0xc3, 0x3e, 0xf5, 0xc5, 0x23, 0x9f, 0x73, 0x74, 0x24, 0xdd, 0x7b, 0x75,
	0xb5, 0x64, 0x4d, 0x19, 0x31, 0xe2, 0x5a, 0x17, 0xde, 0xd9, 0x76, 0xe1, 0x5d, 0x74, 0xf4, 0x99,
	0x34, 0x41, 0x3a, 0x6e, 0x5c, 0xe0, 0x7f, 0x79, 0x1e, 0x94, 0xe7, 0x26, 0x9f, 0xd8, 0x91, 0xe7,
	0x49, 0xea, 0x0a, 0xed, 0xdb, 0x30, 0x6a, 0x7b, 0x21, 0xaf, 0x43, 0xfa, 0x6d, 0xcb, 0xb1, 0x08,
	0xc1, 0x84, 0x8d, 0xcf, 0xc8, 0x6a, 0xcd, 0x8f, 0xbf, 0x38, 0x7e, 0xdb, 0xa7, 0x1f, 0x90, 0x25,
	0xd0, 0xf0, 0x5c, 0x58, 0xcd, 0xe6, 0xd6, 0xe7, 0x36, 0x97, 0x2e, 0x16, 0x01, 0x38, 0x17, 0x56,
	0x6f, 0x38, 0x42, 0xbc, 0x15, 0x5c, 0xe9, 0x28, 0x4c, 0x46, 0x3f, 0x24, 0xf7, 0xa3, 0x0e, 0x91,
	0x8f, 0x4d, 0x88, 0xce, 0xdf, 0x56, 0xea, 0x16, 0x60, 0xcf, 0x11, 0xa2, 0x8f, 0x70, 0x42, 0x5e,
	0xda, 0x81, 0xf6, 0xec, 0x7f, 0x49, 0xb0, 0xe4, 0xad, 0x38, 0x4f, 0x00, 0x7d, 0x4c, 0x5a, 0x40,
	0x4f, 0x0d, 0xe6, 0x13, 0x0f, 0x33, 0xaa, 0xf9, 0x1b, 0xff, 0xac, 0x90, 0x56, 0x6d, 0x8b, 0x74,
	0x9d, 0xb4, 0x94, 0x0e, 0xd2, 0x9b, 0x22, 0x1a, 0x97, 0xb3, 0x0e, 0xae, 0x58, 0x83, 0xe8, 0x53,
	0xb2, 0x66, 0x54, 0x3a, 0xce, 0xd0, 0x79, 0x2b, 0x22, 0xf7, 0x7a, 0xc2, 0xba, 0x49, 0xb7, 0x82,
	0xc4, 0x49, 0xc2, 0x2f, 0xf4, 0x84, 0x6e, 0x91, 0x35, 0xe9, 0xf2, 0xe8, 0x5d, 0x96, 0x69, 0xcf,
	0x87, 0xc2, 0x9a, 0xec, 0x96, 0xf5, 0x92, 0x76, 0x75, 0x46, 0x9c, 0x24, 0x9c, 0x7e, 0x42, 0x56,
	0x6a, 0xe2, 0x78, 0x5b, 0x68, 0xb6, 0x93, 0xa4, 0x0f, 0x66, 0xf0, 0xe5, 0x6d, 0xa1, 0xe9, 0x2a,
	0x99, 0x9f, 0x18, 0xc5, 0x76, 0x13, 0x09, 0x43, 0x4a, 0xc9, 0xff, 0xc7, 0x37, 0x46, 0xb1, 0xbd,
	0x04, 0xa5, 0x31, 0xa8, 0x0a, 0xa3, 0xd8, 0xe7, 0xa8, 0x2a, 0x8c, 0x82, 0x05, 0x4a, 0x65, 0x78,
	0xfd, 0x7c, 0xcf, 0x70, 0x81, 0x52, 0x99, 0xa3, 0xda, 0x11, 0x1f, 0x92, 0x45, 0x10, 0xa6, 0x0c,
	0x7d, 0x91, 0x14, 0x0b, 0xa5, 0x32, 0x90, 0x20, 0x58, 0x49, 0x66, 0xda, 0xb0, 0x2f, 0x71, 0x25,
	0x18, 0xc3, 0x4a, 0x5a, 0x1a, 0xf6, 0x15, 0xae, 0xa4, 0xa5, 0xa1, 0x5b, 0x84, 0x46, 0x57, 0x70,
	0x11, 0x82, 0xb6, 0xbc, 0x10, 0x3e, 0x42, 0x82, 0xd8, 0xd7, 0x18, 0xa4, 0xe8, 0x8a, 0x7d, 0x20,
	0x5e, 0x09, 0x1f, 0xcf, 0x4b, 0x4b, 0x37, 0xc8, 0xf2, 0x4c, 0x0c, 0x07, 0xfb, 0xa6, 0x4a, 0x73,
	0xa5, 0xbb, 0x32, 0x8a, 0xbe, 0x4f, 0x16, 0x0a, 0x99, 0xd2, 0xcc, 0xbe, 0x4d, 0xec, 0xbd, 0x42,
	0x42, 0x8e, 0xe9, 0x7b, 0x04, 0x46, 0x30, 0xeb, 0xbb, 0x3b, 0x1c, 0x22, 0xf2, 0x98, 0xb4, 0xaa,
	0x22, 0xe4, 0xc1, 0x28, 0xb6, 0x8f, 0x89, 0xaf, 0xa0, 0xbe, 0x51, 0x70, 0x46, 0xa5, 0x27, 0xe0,
	0xb8, 0xcd, 0x0e, 0xf0, 0x8c, 0x4a, 0x4f, 0xce, 0x4b, 0xbb, 0x5d, 0xa3, 0x3a, 0xec, 0xb0, 0x4e,
	0x75, 0x6a, 0x54, 0x97, 0x1d, 0xd5, 0xa9, 0x6e, 0x8d, 0xea, 0xb1, 0xe3, 0x3a, 0xd5, 0xab, 0x51,
	0x3b, 0xec, 0xa4, 0x4e, 0xed, 0xd4, 0xa8, 0x5d, 0xf6, 0x7d, 0x9d, 0xda, 0xad, 0x51, 0x7b, 0xec,
	0x79, 0x9d, 0xda, 0xa3, 0x1f, 0x91, 0x07, 0x56, 0xe4, 0x25, 0x4f, 0xb7, 0x43, 0x89, 0x28, 0xd8,
	0x69, 0x12, 0xdc, 0x07, 0xf4, 0x52, 0x87, 0x78, 0x24, 0xa2, 0xa0, 0x6f, 0xc8, 0xbc, 0xb7, 0x82,
	0xfd, 0xb0, 0x3e, 0xb7, 0xd9, 0xea, 0x1c, 0xb7, 0xff, 0xf3, 0x35, 0x6e, 0xcf, 0xae, 0xe4, 0x05,
	0x38, 0xc2, 0xce, 0x20, 0x5b, 0x91, 0x1b, 0xc5, 0x5e, 0xe0, 0xce, 0xd2, 0xff, 0xd3, 0x14, 0x77,
	0xa4, 0x44, 0x66, 0x44, 0x60, 0x2f, 0x31, 0xee, 0x09, 0xda, 0x07, 0x04, 0xae, 0xcf, 0x40, 0x04,
	0xcd, 0xad, 0x90, 0x5c, 0x28, 0xe5, 0x75, 0x08, 0xdb, 0xec, 0x0c, 0x2b, 0x03, 0x88, 0x33, 0x21,
	0xf7, 0x2b, 0x98, 0x7e, 0x4a, 0xd6, 0x2a, 0x19, 0x1f, 0x64, 0xd7, 0x3c, 0x98, 0xbf, 0xf5, 0x36,
	0x3b, 0xc7, 0x92, 0xb5, 0x49, 0x77, 0x90, 0x5d, 0xf7, 0x01, 0x6d, 0xb2, 0xed, 0xb0, 0x1f, 0x9b,
	0x6c, 0x3b, 0x4d, 0xb6, 0x1d, 0xf6, 0xaa, 0xc1, 0xb6, 0xd3, 0x64, 0xdb, 0x65, 0x3f, 0x35, 0xd9,
	0x76, 0x9b, 0x6c, 0xbb, 0xec, 0xa2, 0xc1, 0xb6, 0xdb, 0x64, 0xdb, 0x63, 0xfd, 0x26, 0xdb, 0x5e,
	0x93, 0x6d, 0x8f, 0x5d, 0x36, 0xd8, 0xf6, 0xa0, 0x2c, 0x0a, 0x39, 0xe0, 0x41, 0x7b, 0x23, 0xb2,
	0x74, 0x59, 0x5e, 0x63, 0x59, 0x14, 0x72, 0xd0, 0x4f, 0x20, 0x5c, 0x99, 0xa7, 0x64, 0xad, 0x70,
	0x37, 0xda, 0xf3, 0x50, 0x16, 0x45, 0x76, 0x8b, 0x9d, 0xe6, 0x0a, 0x17, 0x4f, 0x44, 0x3f, 0xe1,
	0xa9, 0xd5, 0x6c, 0x4d, 0xb5, 0xd2, 0xe5, 0xa1, 0xb4, 0xd8, 0x34, 0xde, 0x60, 0x03, 0x4b, 0xc4,
	0xe1, 0x0c, 0x87, 0xfe, 0x32, 0xc8, 0x9c, 0x84, 0x3d, 0x8e, 0x72, 0x11, 0x4b, 0xaf, 0xd9, 0xcf,
	0xb8, 0xcf, 0x04, 0xf7, 0xa7, 0x28, 0x7d, 0x42, 0x96, 0x51, 0x38, 0xd1, 0x3e, 0x80, 0xe3, 0x5b,
	0xdc, 0x66, 0x02, 0xaf, 0x10, 0x83, 0xe6, 0x8f, 0xa2, 0x4c, 0xe7, 0xa3, 0x38, 0x66, 0xbf, 0x60,
	0x57, 0x48, 0xd8, 0xcb, 0x04, 0xd1, 0x8f, 0x09, 0x3a, 0x73, 0x39, 0xd6, 0xf2, 0x3a, 0x94, 0x96,
	0xfd, 0x9a, 0x44, 0xe8, 0x7e, 0x58, 0x81, 0x50, 0x93, 0xd5, 0x03, 0x04, 0xc1, 0x63, 0xbf, 0x61,
	0x4d, 0x22, 0x04, 0x81, 0x03, 0x41, 0xe5, 0xe3, 0xca, 0x3c, 0xb2, 0xdf, 0x51, 0x80, 0x26, 0x80,
	0x40, 0x60, 0x87, 0xbe, 0xe4, 0x56, 0xbc, 0x73, 0x55, 0x67, 0xfe, 0x03, 0x77, 0x3c, 0xf4, 0xe5,
	0x19, 0x80, 0x29, 0x58, 0x53, 0x95, 0xc9, 0xa7, 0x2a, 0x3e, 0x53, 0x01, 0x98, 0x54, 0x8f, 0x08,
	0x71, 0xda, 0xf2, 0x10, 0xbd, 0xc9, 0x47, 0xec, 0x4f, 0x7c, 0xb1, 0x9c, 0xb6, 0xfd, 0x04, 0x00,
	0x5d, 0x78, 0xa7, 0x4a, 0x99, 0x6e, 0x97, 0x40, 0xba, 0x42, 0x4e, 0x15, 0x84, 0x6e, 0x96, 0x5e,
	0x78, 0xf2, 0x06, 0xb8, 0x44, 0x98, 0xa6, 0xb7, 0x7a, 0xf5, 0xa6, 0x4d, 0x17, 0x24, 0x12, 0xcf,
	0x53, 0x60, 0xbf, 0x05, 0xc1, 0x13, 0xb2, 0x9c, 0x04, 0x5e, 0x4f, 0x4c, 0x4a, 0x80, 0xaa, 0xea,
	0x44, 0x78, 0x78, 0xb7, 0x12, 0x06, 0x22, 0x3b, 0x1c, 0x71, 0xa5, 0x27, 0x46, 0xa4, 0xbc, 0xeb,
	0xaa, 0xc7, 0x0c, 0x47, 0x47, 0x53, 0x0c, 0xb6, 0x3b, 0xbe, 0xb9, 0xcb, 0xe3, 0x10, 0xb7, 0x3b,
	0xbe, 0x99, 0x26, 0xf1, 0x21, 0x59, 0x04, 0x8f, 0x81, 0x89, 0x81, 0x8d, 0xb0, 0x53, 0xd8, 0xe1,
	0xe8, 0xc0, 0xc4, 0x00, 0xf9, 0xd5, 0xf9, 0xc8, 0xe4, 0x5a, 0x7b, 0x5e, 0x06, 0xcd, 0xc6, 0x98,
	0xdf, 0x29, 0xf6, 0x3a, 0x68, 0xca, 0xc8, 0x42, 0xc8, 0x6d, 0xe1, 0x8c, 0x62, 0x06, 0x27, 0x57,
	0x7f, 0xc1, 0x17, 0x9a, 0x92, 0x74, 0x4a, 0xb3, 0x77, 0x48, 0x79, 0x2b, 0x0e, 0x9d, 0xd2, 0x74,
	0x93, 0xac, 0x6a, 0x69, 0xb8, 0xc8, 0x8a, 0xf1, 0xdd, 0x77, 0xc1, 0x35, 0x96, 0xa1, 0x96, 0x66,
	0x1f, 0x60, 0x8c, 0xc2, 0xe0, 0x5e, 0xfa, 0xbe, 0xe9, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x2c,
	0x12, 0x0d, 0x40, 0xf4, 0x08, 0x00, 0x00,
}
