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

func init() {
	proto.RegisterType((*DiagEeprom_KEYS)(nil), "cisco_ios_xr_sdr_invmgr_diag_oper.diag.racks.rack.chassis.diag_eeprom_KEYS")
	proto.RegisterType((*RmaDetail)(nil), "cisco_ios_xr_sdr_invmgr_diag_oper.diag.racks.rack.chassis.rma_detail")
	proto.RegisterType((*DiagEeprom)(nil), "cisco_ios_xr_sdr_invmgr_diag_oper.diag.racks.rack.chassis.diag_eeprom")
}

func init() { proto.RegisterFile("diag_eeprom.proto", fileDescriptor_1b4ed4664a25087f) }

var fileDescriptor_1b4ed4664a25087f = []byte{
	// 1043 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x53, 0x1b, 0xb7,
	0x17, 0x1d, 0x7e, 0xfc, 0x26, 0x80, 0x1c, 0x02, 0xe8, 0xa1, 0x55, 0xa6, 0x93, 0x09, 0x25, 0xed,
	0x34, 0x0d, 0x33, 0xee, 0x60, 0x1b, 0x68, 0xfa, 0xcd, 0x67, 0x43, 0x13, 0x68, 0x8a, 0x09, 0x69,
	0xfa, 0xa5, 0xca, 0x92, 0x6c, 0x2b, 0xac, 0x56, 0x3b, 0x92, 0xd6, 0xd4, 0xfd, 0x27, 0xfb, 0x2f,
	0x75, 0xae, 0xee, 0x3a, 0xde, 0x87, 0x7d, 0xea, 0x0b, 0xb3, 0x3e, 0xe7, 0xe8, 0xdc, 0xab, 0x7b,
	0xa5, 0x2b, 0xc8, 0x86, 0x32, 0x62, 0xc4, 0xb5, 0x2e, 0xbc, 0xb3, 0xed, 0xc2, 0xbb, 0xe8, 0xe8,
	0x53, 0x69, 0x82, 0x74, 0xdc, 0xb8, 0xc0, 0xff, 0xf2, 0x3c, 0x28, 0xcf, 0x4d, 0x3e, 0xb1, 0x23,
	0xcf, 0x93, 0xd4, 0x15, 0xda, 0xb7, 0xe1, 0xab, 0xed, 0x85, 0xbc, 0x09, 0xe9, 0x6f, 0x5b, 0x8e,
	0x45, 0x08, 0x26, 0x6c, 0x7d, 0x46, 0xd6, 0x6b, 0x7e, 0xfc, 0xf9, 0xc9, 0x9b, 0x3e, 0xfd, 0x80,
	0xac, 0x80, 0x86, 0xe7, 0xc2, 0x6a, 0xb6, 0xb0, 0xb9, 0xf0, 0x78, 0xe5, 0x72, 0x19, 0x80, 0x0b,
	0x61, 0xf5, 0x96, 0x23, 0xc4, 0x5b, 0xc1, 0x95, 0x8e, 0xc2, 0x64, 0xf4, 0x43, 0x72, 0x37, 0xea,
	0x10, 0xf9, 0xd8, 0x84, 0xe8, 0xfc, 0xb4, 0x52, 0xb7, 0x00, 0x7b, 0x86, 0x10, 0x7d, 0x80, 0x0b,
	0xf2, 0xd2, 0x0e, 0xb4, 0x67, 0xff, 0x4b, 0x82, 0x15, 0x6f, 0xc5, 0x45, 0x02, 0xe8, 0x43, 0xd2,
	0x02, 0x7a, 0x66, 0xb0, 0x98, 0x78, 0x58, 0x51, 0xad, 0xdf, 0xfa, 0x67, 0x8d, 0xb4, 0x6a, 0x29,
	0xd2, 0x4d, 0xd2, 0x52, 0x3a, 0x48, 0x6f, 0x8a, 0x68, 0x5c, 0xce, 0x3a, 0x18, 0xb1, 0x06, 0xd1,
	0x27, 0x64, 0xc3, 0xa8, 0xb4, 0x9d, 0xa1, 0xf3, 0x56, 0x44, 0xee, 0xf5, 0x84, 0x75, 0x93, 0x6e,
	0x0d, 0x89, 0xd3, 0x84, 0x5f, 0xea, 0x09, 0xdd, 0x26, 0x1b, 0xd2, 0xe5, 0xd1, 0xbb, 0x2c, 0xd3,
	0x9e, 0x0f, 0x85, 0x35, 0xd9, 0x94, 0xf5, 0x92, 0x76, 0x7d, 0x4e, 0x9c, 0x26, 0x9c, 0x7e, 0x42,
	0xd6, 0x6a, 0xe2, 0x38, 0x2d, 0x34, 0xdb, 0x4d, 0xd2, 0x7b, 0x73, 0xf8, 0x6a, 0x5a, 0x68, 0xba,
	0x4e, 0x16, 0x27, 0x46, 0xb1, 0xbd, 0x44, 0xc2, 0x27, 0xa5, 0xe4, 0xff, 0xe3, 0x5b, 0xa3, 0xd8,
	0x7e, 0x82, 0xd2, 0x37, 0xa8, 0x0a, 0xa3, 0xd8, 0xe7, 0xa8, 0x2a, 0x8c, 0x82, 0x00, 0xa5, 0x32,
	0xbc, 0xbe, 0xbf, 0xa7, 0x18, 0xa0, 0x54, 0xe6, 0xb8, 0xb6, 0xc5, 0xfb, 0x64, 0x19, 0x84, 0xa9,
	0x43, 0x5f, 0x24, 0xc5, 0x52, 0xa9, 0x0c, 0x34, 0x08, 0x22, 0xc9, 0x4c, 0x1b, 0xf6, 0x25, 0x46,
	0x82, 0x6f, 0x88, 0xa4, 0xa5, 0x61, 0x5f, 0x61, 0x24, 0x2d, 0x0d, 0xdd, 0x26, 0x34, 0xba, 0x82,
	0x8b, 0x10, 0xb4, 0xe5, 0x85, 0xf0, 0x11, 0x1a, 0xc4, 0xbe, 0xc6, 0x22, 0x45, 0x57, 0x1c, 0x00,
	0xf1, 0x52, 0xf8, 0x78, 0x51, 0x5a, 0xba, 0x45, 0x56, 0xe7, 0x62, 0xd8, 0xd8, 0x37, 0x55, 0x9b,
	0x2b, 0xdd, 0xb5, 0x51, 0xf4, 0x7d, 0xb2, 0x54, 0xc8, 0xd4, 0x66, 0xf6, 0x6d, 0x62, 0xef, 0x14,
	0x12, 0x7a, 0x4c, 0xdf, 0x23, 0xf0, 0x05, 0xab, 0xbe, 0x7b, 0x87, 0x43, 0x45, 0x1e, 0x92, 0x56,
	0x75, 0x08, 0x79, 0x30, 0x8a, 0x1d, 0x60, 0xe3, 0x2b, 0xa8, 0x6f, 0x14, 0xec, 0x51, 0xe9, 0x09,
	0x38, 0xee, 0xb0, 0x43, 0xdc, 0xa3, 0xd2, 0x93, 0x8b, 0xd2, 0xee, 0xd4, 0xa8, 0x0e, 0x3b, 0xaa,
	0x53, 0x9d, 0x1a, 0xd5, 0x65, 0xc7, 0x75, 0xaa, 0x5b, 0xa3, 0x7a, 0xec, 0xa4, 0x4e, 0xf5, 0x6a,
	0xd4, 0x2e, 0x3b, 0xad, 0x53, 0xbb, 0x35, 0x6a, 0x8f, 0x7d, 0x5f, 0xa7, 0xf6, 0x6a, 0xd4, 0x3e,
	0x7b, 0x56, 0xa7, 0xf6, 0xe9, 0x47, 0xe4, 0x9e, 0x15, 0x79, 0xc9, 0xd3, 0xed, 0x50, 0x22, 0x0a,
	0x76, 0x96, 0x04, 0x77, 0x01, 0xbd, 0xd2, 0x21, 0x1e, 0x8b, 0x28, 0xe8, 0x6b, 0xb2, 0xe8, 0xad,
	0x60, 0x3f, 0x6c, 0x2e, 0x3c, 0x6e, 0x75, 0x4e, 0xda, 0xff, 0xf9, 0x1a, 0xb7, 0xe7, 0x57, 0xf2,
	0x12, 0x1c, 0x21, 0x33, 0xe8, 0x56, 0xe4, 0x46, 0xb1, 0xe7, 0x98, 0x59, 0xfa, 0x7d, 0x96, 0xea,
	0x8e, 0x94, 0xc8, 0x8c, 0x08, 0xec, 0x05, 0xd6, 0x3d, 0x41, 0x07, 0x80, 0xc0, 0xf5, 0x19, 0x88,
	0xa0, 0xb9, 0x15, 0x92, 0x0b, 0xa5, 0xbc, 0x0e, 0x61, 0x87, 0x9d, 0xe3, 0xc9, 0x00, 0xe2, 0x5c,
	0xc8, 0x83, 0x0a, 0xa6, 0x9f, 0x92, 0x8d, 0x4a, 0xc6, 0x07, 0xd9, 0x0d, 0x0f, 0xe6, 0x6f, 0xbd,
	0xc3, 0x2e, 0xf0, 0xc8, 0xda, 0xa4, 0x3b, 0xcc, 0x6e, 0xfa, 0x80, 0x36, 0xd9, 0x76, 0xd8, 0x8f,
	0x4d, 0xb6, 0x9d, 0x26, 0xdb, 0x0e, 0x7b, 0xd9, 0x60, 0xdb, 0x69, 0xb2, 0xed, 0xb2, 0x9f, 0x9a,
	0x6c, 0xbb, 0x4d, 0xb6, 0x5d, 0x76, 0xd9, 0x60, 0xdb, 0x6d, 0xb2, 0xed, 0xb1, 0x7e, 0x93, 0x6d,
	0xaf, 0xc9, 0xb6, 0xc7, 0xae, 0x1a, 0x6c, 0x7b, 0x70, 0x2c, 0x0a, 0x39, 0xe0, 0x41, 0x7b, 0x23,
	0xb2, 0x74, 0x59, 0x5e, 0xe1, 0xb1, 0x28, 0xe4, 0xa0, 0x9f, 0x40, 0xb8, 0x32, 0x4f, 0xc8, 0x46,
	0xe1, 0x6e, 0xb5, 0xe7, 0xa1, 0x2c, 0x8a, 0x6c, 0x8a, 0x93, 0xe6, 0x1a, 0x83, 0x27, 0xa2, 0x9f,
	0xf0, 0x34, 0x6a, 0xb6, 0x67, 0x5a, 0xe9, 0xf2, 0x50, 0x5a, 0x1c, 0x1a, 0xaf, 0x71, 0x80, 0x25,
	0xe2, 0x68, 0x8e, 0xc3, 0x7c, 0x19, 0x64, 0x4e, 0x42, 0x8e, 0xa3, 0x5c, 0xc4, 0xd2, 0x6b, 0xf6,
	0x33, 0xe6, 0x99, 0xe0, 0xfe, 0x0c, 0xa5, 0x8f, 0xc8, 0x2a, 0x0a, 0x27, 0xda, 0x07, 0x70, 0x7c,
	0x83, 0x69, 0x26, 0xf0, 0x1a, 0x31, 0x18, 0xfe, 0x28, 0xca, 0x74, 0x3e, 0x8a, 0x63, 0xf6, 0x0b,
	0x4e, 0x85, 0x84, 0xbd, 0x48, 0x10, 0xfd, 0x98, 0xa0, 0x33, 0x97, 0x63, 0x2d, 0x6f, 0x42, 0x69,
	0xd9, 0xaf, 0x49, 0x84, 0xee, 0x47, 0x15, 0x08, 0x67, 0xb2, 0x7a, 0x80, 0xa0, 0x78, 0xec, 0x37,
	0x3c, 0x93, 0x08, 0x41, 0xe1, 0x40, 0x50, 0xf9, 0xb8, 0x32, 0x8f, 0xec, 0x77, 0x14, 0xa0, 0x09,
	0x20, 0x50, 0xd8, 0xa1, 0x2f, 0xb9, 0x15, 0x6f, 0x5d, 0x35, 0x99, 0xff, 0xc0, 0x8c, 0x87, 0xbe,
	0x3c, 0x07, 0x30, 0x15, 0x6b, 0xa6, 0x32, 0xf9, 0x4c, 0xc5, 0xe7, 0x2a, 0x00, 0x93, 0xea, 0x01,
	0x21, 0x4e, 0x5b, 0x1e, 0xa2, 0x37, 0xf9, 0x88, 0xfd, 0x89, 0x2f, 0x96, 0xd3, 0xb6, 0x9f, 0x00,
	0xa0, 0x0b, 0xef, 0x54, 0x29, 0xd3, 0xed, 0x12, 0x48, 0x57, 0xc8, 0x99, 0x82, 0xd2, 0xcd, 0xdb,
	0x0b, 0x4f, 0xde, 0x00, 0x43, 0x84, 0x59, 0x7b, 0xab, 0x57, 0x6f, 0x36, 0x74, 0x41, 0x22, 0x71,
	0x3f, 0x05, 0xce, 0x5b, 0x10, 0x3c, 0x22, 0xab, 0x49, 0xe0, 0xf5, 0xc4, 0xa4, 0x06, 0xa8, 0xea,
	0x9c, 0x08, 0x0f, 0xef, 0x56, 0xc2, 0x40, 0x64, 0x87, 0x23, 0xae, 0xf4, 0xc4, 0x88, 0xd4, 0x77,
	0x5d, 0xcd, 0x98, 0xe1, 0xe8, 0x78, 0x86, 0x41, 0xba, 0xe3, 0xdb, 0x77, 0x7d, 0x1c, 0x62, 0xba,
	0xe3, 0xdb, 0x59, 0x13, 0xef, 0x93, 0x65, 0xf0, 0x18, 0x98, 0x18, 0xd8, 0x08, 0x27, 0x85, 0x1d,
	0x8e, 0x0e, 0x4d, 0x0c, 0xd0, 0x5f, 0x9d, 0x8f, 0x4c, 0xae, 0xb5, 0xe7, 0x65, 0xd0, 0x6c, 0x8c,
	0xfd, 0x9d, 0x61, 0xaf, 0x82, 0xa6, 0x8c, 0x2c, 0x85, 0xdc, 0x16, 0xce, 0x28, 0x66, 0x70, 0x71,
	0xf5, 0x13, 0x7c, 0x61, 0x28, 0x49, 0xa7, 0x34, 0x7b, 0x8b, 0x94, 0xb7, 0xe2, 0xc8, 0x29, 0x3d,
	0xb8, 0x93, 0xfe, 0x6b, 0xe9, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x21, 0xab, 0xf5, 0x97, 0xca,
	0x08, 0x00, 0x00,
}