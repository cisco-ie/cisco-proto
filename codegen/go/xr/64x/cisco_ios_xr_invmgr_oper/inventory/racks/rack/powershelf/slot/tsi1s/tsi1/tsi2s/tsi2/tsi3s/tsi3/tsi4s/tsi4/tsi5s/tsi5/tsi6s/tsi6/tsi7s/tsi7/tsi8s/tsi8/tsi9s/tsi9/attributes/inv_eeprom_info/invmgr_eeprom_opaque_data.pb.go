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
// source: invmgr_eeprom_opaque_data.proto

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_attributes_inv_eeprom_info

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

type InvmgrEepromOpaqueData_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Name_1               string   `protobuf:"bytes,2,opt,name=name_1,json=name1,proto3" json:"name_1,omitempty"`
	Name_2               string   `protobuf:"bytes,3,opt,name=name_2,json=name2,proto3" json:"name_2,omitempty"`
	Name_3               string   `protobuf:"bytes,4,opt,name=name_3,json=name3,proto3" json:"name_3,omitempty"`
	Name_4               string   `protobuf:"bytes,5,opt,name=name_4,json=name4,proto3" json:"name_4,omitempty"`
	Name_5               string   `protobuf:"bytes,6,opt,name=name_5,json=name5,proto3" json:"name_5,omitempty"`
	Name_6               string   `protobuf:"bytes,7,opt,name=name_6,json=name6,proto3" json:"name_6,omitempty"`
	Name_7               string   `protobuf:"bytes,8,opt,name=name_7,json=name7,proto3" json:"name_7,omitempty"`
	Name_8               string   `protobuf:"bytes,9,opt,name=name_8,json=name8,proto3" json:"name_8,omitempty"`
	Name_9               string   `protobuf:"bytes,10,opt,name=name_9,json=name9,proto3" json:"name_9,omitempty"`
	Name_10              string   `protobuf:"bytes,11,opt,name=name_10,json=name10,proto3" json:"name_10,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvmgrEepromOpaqueData_KEYS) Reset()         { *m = InvmgrEepromOpaqueData_KEYS{} }
func (m *InvmgrEepromOpaqueData_KEYS) String() string { return proto.CompactTextString(m) }
func (*InvmgrEepromOpaqueData_KEYS) ProtoMessage()    {}
func (*InvmgrEepromOpaqueData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6030716c6d5948f0, []int{0}
}

func (m *InvmgrEepromOpaqueData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvmgrEepromOpaqueData_KEYS.Unmarshal(m, b)
}
func (m *InvmgrEepromOpaqueData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvmgrEepromOpaqueData_KEYS.Marshal(b, m, deterministic)
}
func (m *InvmgrEepromOpaqueData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvmgrEepromOpaqueData_KEYS.Merge(m, src)
}
func (m *InvmgrEepromOpaqueData_KEYS) XXX_Size() int {
	return xxx_messageInfo_InvmgrEepromOpaqueData_KEYS.Size(m)
}
func (m *InvmgrEepromOpaqueData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InvmgrEepromOpaqueData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InvmgrEepromOpaqueData_KEYS proto.InternalMessageInfo

func (m *InvmgrEepromOpaqueData_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvmgrEepromOpaqueData_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *InvmgrEepromOpaqueData_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *InvmgrEepromOpaqueData_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
	}
	return ""
}

func (m *InvmgrEepromOpaqueData_KEYS) GetName_4() string {
	if m != nil {
		return m.Name_4
	}
	return ""
}

func (m *InvmgrEepromOpaqueData_KEYS) GetName_5() string {
	if m != nil {
		return m.Name_5
	}
	return ""
}

func (m *InvmgrEepromOpaqueData_KEYS) GetName_6() string {
	if m != nil {
		return m.Name_6
	}
	return ""
}

func (m *InvmgrEepromOpaqueData_KEYS) GetName_7() string {
	if m != nil {
		return m.Name_7
	}
	return ""
}

func (m *InvmgrEepromOpaqueData_KEYS) GetName_8() string {
	if m != nil {
		return m.Name_8
	}
	return ""
}

func (m *InvmgrEepromOpaqueData_KEYS) GetName_9() string {
	if m != nil {
		return m.Name_9
	}
	return ""
}

func (m *InvmgrEepromOpaqueData_KEYS) GetName_10() string {
	if m != nil {
		return m.Name_10
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
	return fileDescriptor_6030716c6d5948f0, []int{1}
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
	Description          string     `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	IdpromFormatRev      string     `protobuf:"bytes,2,opt,name=idprom_format_rev,json=idpromFormatRev,proto3" json:"idprom_format_rev,omitempty"`
	ControllerFamily     string     `protobuf:"bytes,3,opt,name=controller_family,json=controllerFamily,proto3" json:"controller_family,omitempty"`
	ControllerType       string     `protobuf:"bytes,4,opt,name=controller_type,json=controllerType,proto3" json:"controller_type,omitempty"`
	Vid                  string     `protobuf:"bytes,5,opt,name=vid,proto3" json:"vid,omitempty"`
	Hwid                 string     `protobuf:"bytes,6,opt,name=hwid,proto3" json:"hwid,omitempty"`
	Pid                  string     `protobuf:"bytes,7,opt,name=pid,proto3" json:"pid,omitempty"`
	UdiDescription       string     `protobuf:"bytes,8,opt,name=udi_description,json=udiDescription,proto3" json:"udi_description,omitempty"`
	UdiName              string     `protobuf:"bytes,9,opt,name=udi_name,json=udiName,proto3" json:"udi_name,omitempty"`
	Clei                 string     `protobuf:"bytes,10,opt,name=clei,proto3" json:"clei,omitempty"`
	Eci                  string     `protobuf:"bytes,11,opt,name=eci,proto3" json:"eci,omitempty"`
	TopAssemPartNum      string     `protobuf:"bytes,12,opt,name=top_assem_part_num,json=topAssemPartNum,proto3" json:"top_assem_part_num,omitempty"`
	TopAssemVid          string     `protobuf:"bytes,13,opt,name=top_assem_vid,json=topAssemVid,proto3" json:"top_assem_vid,omitempty"`
	PcaNum               string     `protobuf:"bytes,14,opt,name=pca_num,json=pcaNum,proto3" json:"pca_num,omitempty"`
	Pcavid               string     `protobuf:"bytes,15,opt,name=pcavid,proto3" json:"pcavid,omitempty"`
	ChassisSid           string     `protobuf:"bytes,16,opt,name=chassis_sid,json=chassisSid,proto3" json:"chassis_sid,omitempty"`
	DevNum1              string     `protobuf:"bytes,17,opt,name=dev_num1,json=devNum1,proto3" json:"dev_num1,omitempty"`
	DevNum2              string     `protobuf:"bytes,18,opt,name=dev_num2,json=devNum2,proto3" json:"dev_num2,omitempty"`
	DevNum3              string     `protobuf:"bytes,19,opt,name=dev_num3,json=devNum3,proto3" json:"dev_num3,omitempty"`
	DevNum4              string     `protobuf:"bytes,20,opt,name=dev_num4,json=devNum4,proto3" json:"dev_num4,omitempty"`
	DevNum5              string     `protobuf:"bytes,21,opt,name=dev_num5,json=devNum5,proto3" json:"dev_num5,omitempty"`
	DevNum6              string     `protobuf:"bytes,22,opt,name=dev_num6,json=devNum6,proto3" json:"dev_num6,omitempty"`
	DevNum7              string     `protobuf:"bytes,23,opt,name=dev_num7,json=devNum7,proto3" json:"dev_num7,omitempty"`
	ManuTestData         string     `protobuf:"bytes,24,opt,name=manu_test_data,json=manuTestData,proto3" json:"manu_test_data,omitempty"`
	Rma                  *RmaDetail `protobuf:"bytes,25,opt,name=rma,proto3" json:"rma,omitempty"`
	AssetId              string     `protobuf:"bytes,26,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	AssetAlias           string     `protobuf:"bytes,27,opt,name=asset_alias,json=assetAlias,proto3" json:"asset_alias,omitempty"`
	BaseMacAddress1      string     `protobuf:"bytes,28,opt,name=base_mac_address1,json=baseMacAddress1,proto3" json:"base_mac_address1,omitempty"`
	MacAddBlkSize1       string     `protobuf:"bytes,29,opt,name=mac_add_blk_size1,json=macAddBlkSize1,proto3" json:"mac_add_blk_size1,omitempty"`
	BaseMacAddress2      string     `protobuf:"bytes,30,opt,name=base_mac_address2,json=baseMacAddress2,proto3" json:"base_mac_address2,omitempty"`
	MacAddBlkSize2       string     `protobuf:"bytes,31,opt,name=mac_add_blk_size2,json=macAddBlkSize2,proto3" json:"mac_add_blk_size2,omitempty"`
	BaseMacAddress3      string     `protobuf:"bytes,32,opt,name=base_mac_address3,json=baseMacAddress3,proto3" json:"base_mac_address3,omitempty"`
	MacAddBlkSize3       string     `protobuf:"bytes,33,opt,name=mac_add_blk_size3,json=macAddBlkSize3,proto3" json:"mac_add_blk_size3,omitempty"`
	BaseMacAddress4      string     `protobuf:"bytes,34,opt,name=base_mac_address4,json=baseMacAddress4,proto3" json:"base_mac_address4,omitempty"`
	MacAddBlkSize4       string     `protobuf:"bytes,35,opt,name=mac_add_blk_size4,json=macAddBlkSize4,proto3" json:"mac_add_blk_size4,omitempty"`
	PcbSerialNum         string     `protobuf:"bytes,36,opt,name=pcb_serial_num,json=pcbSerialNum,proto3" json:"pcb_serial_num,omitempty"`
	PowerSupplyType      string     `protobuf:"bytes,37,opt,name=power_supply_type,json=powerSupplyType,proto3" json:"power_supply_type,omitempty"`
	PowerConsumption     string     `protobuf:"bytes,38,opt,name=power_consumption,json=powerConsumption,proto3" json:"power_consumption,omitempty"`
	BlockSignature       string     `protobuf:"bytes,39,opt,name=block_signature,json=blockSignature,proto3" json:"block_signature,omitempty"`
	BlockVersion         string     `protobuf:"bytes,40,opt,name=block_version,json=blockVersion,proto3" json:"block_version,omitempty"`
	BlockLength          string     `protobuf:"bytes,41,opt,name=block_length,json=blockLength,proto3" json:"block_length,omitempty"`
	BlockChecksum        string     `protobuf:"bytes,42,opt,name=block_checksum,json=blockChecksum,proto3" json:"block_checksum,omitempty"`
	EepromSize           string     `protobuf:"bytes,43,opt,name=eeprom_size,json=eepromSize,proto3" json:"eeprom_size,omitempty"`
	BlockCount           string     `protobuf:"bytes,44,opt,name=block_count,json=blockCount,proto3" json:"block_count,omitempty"`
	FruMajorType         string     `protobuf:"bytes,45,opt,name=fru_major_type,json=fruMajorType,proto3" json:"fru_major_type,omitempty"`
	FruMinorType         string     `protobuf:"bytes,46,opt,name=fru_minor_type,json=fruMinorType,proto3" json:"fru_minor_type,omitempty"`
	OemString            string     `protobuf:"bytes,47,opt,name=oem_string,json=oemString,proto3" json:"oem_string,omitempty"`
	ProductId            string     `protobuf:"bytes,48,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SerialNumber         string     `protobuf:"bytes,49,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	PartNumber           string     `protobuf:"bytes,50,opt,name=part_number,json=partNumber,proto3" json:"part_number,omitempty"`
	PartRevision         string     `protobuf:"bytes,51,opt,name=part_revision,json=partRevision,proto3" json:"part_revision,omitempty"`
	MfgDeviation         string     `protobuf:"bytes,52,opt,name=mfg_deviation,json=mfgDeviation,proto3" json:"mfg_deviation,omitempty"`
	HwVersion            string     `protobuf:"bytes,53,opt,name=hw_version,json=hwVersion,proto3" json:"hw_version,omitempty"`
	MfgBits              string     `protobuf:"bytes,54,opt,name=mfg_bits,json=mfgBits,proto3" json:"mfg_bits,omitempty"`
	EngineerUse          string     `protobuf:"bytes,55,opt,name=engineer_use,json=engineerUse,proto3" json:"engineer_use,omitempty"`
	Snmpoid              string     `protobuf:"bytes,56,opt,name=snmpoid,proto3" json:"snmpoid,omitempty"`
	RmaCode              string     `protobuf:"bytes,57,opt,name=rma_code,json=rmaCode,proto3" json:"rma_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DiagEeprom) Reset()         { *m = DiagEeprom{} }
func (m *DiagEeprom) String() string { return proto.CompactTextString(m) }
func (*DiagEeprom) ProtoMessage()    {}
func (*DiagEeprom) Descriptor() ([]byte, []int) {
	return fileDescriptor_6030716c6d5948f0, []int{2}
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

type InvmgrEepromOpaqueData struct {
	InvCardType          uint32      `protobuf:"varint,50,opt,name=inv_card_type,json=invCardType,proto3" json:"inv_card_type,omitempty"`
	OpaqueData           []uint32    `protobuf:"varint,51,rep,packed,name=opaque_data,json=opaqueData,proto3" json:"opaque_data,omitempty"`
	OpaqueDataSize       uint32      `protobuf:"varint,52,opt,name=opaque_data_size,json=opaqueDataSize,proto3" json:"opaque_data_size,omitempty"`
	HasEeprom            uint32      `protobuf:"varint,53,opt,name=has_eeprom,json=hasEeprom,proto3" json:"has_eeprom,omitempty"`
	Eeprom               *DiagEeprom `protobuf:"bytes,54,opt,name=eeprom,proto3" json:"eeprom,omitempty"`
	Description          string      `protobuf:"bytes,55,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *InvmgrEepromOpaqueData) Reset()         { *m = InvmgrEepromOpaqueData{} }
func (m *InvmgrEepromOpaqueData) String() string { return proto.CompactTextString(m) }
func (*InvmgrEepromOpaqueData) ProtoMessage()    {}
func (*InvmgrEepromOpaqueData) Descriptor() ([]byte, []int) {
	return fileDescriptor_6030716c6d5948f0, []int{3}
}

func (m *InvmgrEepromOpaqueData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvmgrEepromOpaqueData.Unmarshal(m, b)
}
func (m *InvmgrEepromOpaqueData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvmgrEepromOpaqueData.Marshal(b, m, deterministic)
}
func (m *InvmgrEepromOpaqueData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvmgrEepromOpaqueData.Merge(m, src)
}
func (m *InvmgrEepromOpaqueData) XXX_Size() int {
	return xxx_messageInfo_InvmgrEepromOpaqueData.Size(m)
}
func (m *InvmgrEepromOpaqueData) XXX_DiscardUnknown() {
	xxx_messageInfo_InvmgrEepromOpaqueData.DiscardUnknown(m)
}

var xxx_messageInfo_InvmgrEepromOpaqueData proto.InternalMessageInfo

func (m *InvmgrEepromOpaqueData) GetInvCardType() uint32 {
	if m != nil {
		return m.InvCardType
	}
	return 0
}

func (m *InvmgrEepromOpaqueData) GetOpaqueData() []uint32 {
	if m != nil {
		return m.OpaqueData
	}
	return nil
}

func (m *InvmgrEepromOpaqueData) GetOpaqueDataSize() uint32 {
	if m != nil {
		return m.OpaqueDataSize
	}
	return 0
}

func (m *InvmgrEepromOpaqueData) GetHasEeprom() uint32 {
	if m != nil {
		return m.HasEeprom
	}
	return 0
}

func (m *InvmgrEepromOpaqueData) GetEeprom() *DiagEeprom {
	if m != nil {
		return m.Eeprom
	}
	return nil
}

func (m *InvmgrEepromOpaqueData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*InvmgrEepromOpaqueData_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.inv_eeprom_info.invmgr_eeprom_opaque_data_KEYS")
	proto.RegisterType((*RmaDetail)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.inv_eeprom_info.rma_detail")
	proto.RegisterType((*DiagEeprom)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.inv_eeprom_info.diag_eeprom")
	proto.RegisterType((*InvmgrEepromOpaqueData)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.attributes.inv_eeprom_info.invmgr_eeprom_opaque_data")
}

func init() { proto.RegisterFile("invmgr_eeprom_opaque_data.proto", fileDescriptor_6030716c6d5948f0) }

var fileDescriptor_6030716c6d5948f0 = []byte{
	// 1296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x97, 0xdf, 0x72, 0x13, 0xb7,
	0x17, 0xc7, 0xc7, 0x04, 0x12, 0x22, 0xe3, 0x90, 0xec, 0xef, 0x07, 0x28, 0x6d, 0x21, 0x21, 0x40,
	0x09, 0xa4, 0x75, 0xf1, 0xff, 0xe4, 0x12, 0x02, 0x4c, 0x99, 0x96, 0x4c, 0xc7, 0xa6, 0xcc, 0xf4,
	0x4a, 0x23, 0xaf, 0x64, 0x5b, 0xcd, 0xee, 0x6a, 0x2b, 0x69, 0x9d, 0xa6, 0xcf, 0xd2, 0x77, 0xe8,
	0x03, 0xf4, 0x01, 0xda, 0xfb, 0xde, 0xf5, 0xae, 0x6f, 0xd2, 0x39, 0x47, 0xb2, 0x77, 0x67, 0x6a,
	0x1e, 0x80, 0x1b, 0xe5, 0xf8, 0x73, 0xce, 0x9e, 0x5d, 0x9d, 0xef, 0xd9, 0xa3, 0x0d, 0xd9, 0x53,
	0xd9, 0x3c, 0x9d, 0x1a, 0x26, 0x65, 0x6e, 0x74, 0xca, 0x74, 0xce, 0x7f, 0x2a, 0x24, 0x13, 0xdc,
	0xf1, 0x66, 0x6e, 0xb4, 0xd3, 0xd1, 0xef, 0xb5, 0x58, 0xd9, 0x58, 0x33, 0xa5, 0x2d, 0xfb, 0xd9,
	0xb0, 0x10, 0xae, 0x73, 0x69, 0x9a, 0x2a, 0x9b, 0xcb, 0xcc, 0x69, 0x73, 0xd9, 0x34, 0x3c, 0x3e,
	0xb7, 0xb8, 0x36, 0x73, 0x7d, 0x21, 0x8d, 0x9d, 0xc9, 0x64, 0xd2, 0xb4, 0x89, 0x76, 0x4d, 0x67,
	0x55, 0xcb, 0xe2, 0x0a, 0x4b, 0x1b, 0xcd, 0x36, 0x2c, 0x1d, 0x34, 0x3b, 0xb0, 0x74, 0xd1, 0xec,
	0xc2, 0xd2, 0x43, 0xb3, 0x07, 0x4b, 0x1f, 0xcd, 0x3e, 0x2c, 0x03, 0x34, 0x07, 0xb0, 0x1c, 0xa3,
	0x79, 0x0c, 0xcb, 0x09, 0x9a, 0x27, 0x4d, 0xee, 0x9c, 0x51, 0xe3, 0xc2, 0x49, 0x0b, 0x8f, 0xb4,
	0xd8, 0x8a, 0xca, 0x26, 0xfa, 0xe0, 0xd7, 0x2b, 0xe4, 0xde, 0x07, 0x77, 0xc8, 0xbe, 0x79, 0xf5,
	0xc3, 0x28, 0x8a, 0xc8, 0xd5, 0x8c, 0xa7, 0x92, 0xd6, 0xf6, 0x6b, 0x87, 0x9b, 0x43, 0xb4, 0xa3,
	0x5b, 0x64, 0x1d, 0xfe, 0xb2, 0x16, 0xbd, 0x82, 0xf4, 0x1a, 0xfc, 0x6a, 0x2d, 0x71, 0x9b, 0xae,
	0x95, 0xb8, 0xbd, 0xc4, 0x1d, 0x7a, 0xb5, 0xc4, 0x9d, 0x25, 0xee, 0xd2, 0x6b, 0x25, 0xee, 0x2e,
	0x71, 0x8f, 0xae, 0x97, 0xb8, 0xb7, 0xc4, 0x7d, 0xba, 0x51, 0xe2, 0xfe, 0x12, 0x0f, 0xe8, 0xf5,
	0x12, 0x0f, 0x96, 0xf8, 0x98, 0x6e, 0x96, 0xf8, 0x78, 0x89, 0x4f, 0x28, 0x29, 0xf1, 0x49, 0x74,
	0x87, 0x6c, 0xf8, 0xed, 0x3c, 0xa3, 0x75, 0xe4, 0x18, 0xd5, 0x7a, 0x76, 0xa0, 0x09, 0x31, 0x29,
	0x67, 0x42, 0x3a, 0xae, 0x92, 0xe8, 0x3e, 0xb9, 0xe1, 0xa4, 0x75, 0x6c, 0xa6, 0x2c, 0xa8, 0x1a,
	0x2a, 0x52, 0x07, 0xf6, 0xb5, 0x47, 0xd1, 0x5d, 0x7f, 0x41, 0x56, 0xa4, 0x63, 0x69, 0x42, 0x71,
	0x36, 0x4d, 0xca, 0xcf, 0x10, 0x44, 0x7b, 0xa4, 0x0e, 0xee, 0x45, 0x02, 0x5f, 0x25, 0xb8, 0x22,
	0x5c, 0x7f, 0xf0, 0xd7, 0x36, 0xa9, 0x0b, 0xc5, 0xa7, 0x41, 0x8d, 0x68, 0x9f, 0xd4, 0x85, 0xb4,
	0xb1, 0x51, 0xb9, 0x53, 0x3a, 0x5b, 0xdc, 0xb1, 0x82, 0xa2, 0xa7, 0x64, 0x47, 0x09, 0x54, 0x6e,
	0xa2, 0x4d, 0xca, 0x1d, 0x33, 0x72, 0x1e, 0x6e, 0x7c, 0xd3, 0x3b, 0x5e, 0x23, 0x1f, 0xca, 0x79,
	0x74, 0x44, 0x76, 0x62, 0x9d, 0x39, 0xa3, 0x93, 0x44, 0x1a, 0x36, 0xe1, 0xa9, 0x4a, 0x16, 0x0f,
	0xb1, 0x5d, 0x3a, 0x5e, 0x23, 0x8f, 0x1e, 0x93, 0x9b, 0x95, 0x60, 0x77, 0x99, 0xcb, 0x20, 0xdf,
	0x56, 0x89, 0xdf, 0x5d, 0xe6, 0x32, 0xda, 0x26, 0x6b, 0x73, 0x25, 0x82, 0x88, 0x60, 0x42, 0xcb,
	0xcc, 0x2e, 0x94, 0x08, 0x02, 0xa2, 0x0d, 0x51, 0xb9, 0x12, 0x41, 0x3c, 0x30, 0xe1, 0x06, 0x85,
	0x50, 0xac, 0xba, 0x3f, 0xaf, 0xe1, 0x56, 0x21, 0xd4, 0xcb, 0xca, 0x16, 0x77, 0xc9, 0x75, 0x08,
	0xc4, 0x2e, 0xf4, 0x72, 0x6e, 0x14, 0x42, 0x9d, 0x41, 0x23, 0x46, 0xe4, 0x6a, 0x9c, 0x48, 0x15,
	0xe4, 0x44, 0x1b, 0xee, 0x24, 0x63, 0x15, 0x94, 0x04, 0x33, 0x3a, 0x22, 0x91, 0xd3, 0x39, 0xe3,
	0xd6, 0xca, 0x94, 0xe5, 0xdc, 0x38, 0x10, 0x88, 0xde, 0xf0, 0x45, 0x72, 0x3a, 0x7f, 0x0e, 0x8e,
	0xef, 0xb8, 0x71, 0x67, 0x45, 0x1a, 0x1d, 0x90, 0x46, 0x19, 0x0c, 0x1b, 0x6b, 0x04, 0x99, 0x43,
	0xdc, 0x7b, 0x25, 0xa0, 0x61, 0xf2, 0x18, 0x65, 0xa6, 0x5b, 0xbe, 0x61, 0xf2, 0x18, 0x34, 0x8e,
	0x6e, 0x13, 0xb0, 0xe0, 0xaa, 0x9b, 0x4b, 0x0e, 0x15, 0xd9, 0x23, 0xf5, 0x78, 0xc6, 0xad, 0x55,
	0x96, 0x59, 0x25, 0xe8, 0xb6, 0x17, 0x3e, 0xa0, 0x91, 0x12, 0xb0, 0x47, 0x21, 0xe7, 0x90, 0xb1,
	0x45, 0x77, 0xfc, 0x1e, 0x85, 0x9c, 0x9f, 0x15, 0x69, 0xab, 0xe2, 0x6a, 0xd3, 0xa8, 0xea, 0x6a,
	0x57, 0x5c, 0x1d, 0xfa, 0xbf, 0xaa, 0xab, 0x53, 0x71, 0x75, 0xe9, 0xff, 0xab, 0xae, 0x6e, 0xc5,
	0xd5, 0xa3, 0xb7, 0xaa, 0xae, 0x5e, 0xc5, 0xd5, 0xa7, 0xb7, 0xab, 0xae, 0x7e, 0xc5, 0x35, 0xa0,
	0x77, 0xaa, 0xae, 0x41, 0xf4, 0x90, 0x6c, 0xa5, 0x3c, 0x2b, 0x18, 0xbe, 0x1d, 0x30, 0x39, 0x28,
	0xc5, 0x80, 0x1b, 0x40, 0xdf, 0x49, 0xeb, 0x5e, 0x72, 0xc7, 0xa3, 0xbf, 0x6b, 0x64, 0xcd, 0xa4,
	0x9c, 0xee, 0xee, 0xd7, 0x0e, 0xeb, 0xed, 0x3f, 0x6a, 0xcd, 0x8f, 0x78, 0x70, 0x36, 0xcb, 0xb1,
	0x30, 0x84, 0x4d, 0x41, 0x75, 0xa0, 0x63, 0x1c, 0x53, 0x82, 0x7e, 0xe2, 0xab, 0x83, 0xbf, 0xdf,
	0xa0, 0xf6, 0xde, 0xc5, 0x13, 0xc5, 0x2d, 0xfd, 0xd4, 0x6b, 0x8f, 0xe8, 0x39, 0x10, 0x78, 0x85,
	0xc7, 0xdc, 0x4a, 0x96, 0xf2, 0x98, 0x71, 0x21, 0x8c, 0xb4, 0xb6, 0x45, 0x3f, 0xf3, 0xdd, 0x09,
	0x8e, 0xb7, 0x3c, 0x7e, 0x1e, 0x70, 0xf4, 0x84, 0xec, 0x84, 0x30, 0x36, 0x4e, 0xce, 0x99, 0x55,
	0xbf, 0xc8, 0x16, 0xbd, 0xeb, 0x5f, 0x9b, 0x14, 0xe3, 0x5e, 0x24, 0xe7, 0x23, 0xa0, 0xab, 0xd2,
	0xb6, 0xe9, 0xbd, 0x55, 0x69, 0xdb, 0xab, 0xd2, 0xb6, 0xe9, 0xde, 0x8a, 0xb4, 0xed, 0x55, 0x69,
	0x3b, 0x74, 0x7f, 0x55, 0xda, 0xce, 0xaa, 0xb4, 0x1d, 0x7a, 0x7f, 0x45, 0xda, 0xce, 0xaa, 0xb4,
	0x5d, 0x7a, 0xb0, 0x2a, 0x6d, 0x77, 0x55, 0xda, 0x2e, 0x7d, 0xb0, 0x22, 0x6d, 0x17, 0x5a, 0x33,
	0x8f, 0xc7, 0xcc, 0x4a, 0xa3, 0x78, 0x82, 0x2f, 0xec, 0x43, 0xdf, 0x9a, 0x79, 0x3c, 0x1e, 0x21,
	0x84, 0xd7, 0xf6, 0x29, 0xd9, 0xc1, 0xf6, 0x62, 0xb6, 0xc8, 0xf3, 0xe4, 0xd2, 0x4f, 0xbb, 0x47,
	0xfe, 0xe6, 0xe8, 0x18, 0x21, 0xc7, 0x71, 0x77, 0xb4, 0x88, 0x8d, 0x75, 0x66, 0x8b, 0xd4, 0x0f,
	0xae, 0xcf, 0xfd, 0x10, 0x45, 0xc7, 0x69, 0xc9, 0x61, 0xc6, 0x8d, 0x13, 0x1d, 0xc3, 0x33, 0x4e,
	0x33, 0xee, 0x0a, 0x23, 0xe9, 0x63, 0xff, 0x9c, 0x88, 0x47, 0x0b, 0x1a, 0x3d, 0x20, 0x0d, 0x1f,
	0x38, 0x97, 0xc6, 0x42, 0xc6, 0x43, 0xff, 0x98, 0x08, 0xdf, 0x7b, 0x06, 0x07, 0x90, 0x0f, 0x4a,
	0x64, 0x36, 0x75, 0x33, 0xfa, 0xc4, 0x4f, 0x26, 0x64, 0xdf, 0x22, 0x8a, 0x1e, 0x11, 0x9f, 0x99,
	0xc5, 0x33, 0x19, 0x9f, 0xdb, 0x22, 0xa5, 0x4f, 0x31, 0xc8, 0x67, 0x3f, 0x0d, 0x10, 0x7a, 0x32,
	0x74, 0x33, 0x14, 0x8f, 0x1e, 0xf9, 0x9e, 0xf4, 0x08, 0x0a, 0x07, 0x01, 0x21, 0x8f, 0x2e, 0x32,
	0x47, 0xbf, 0xf0, 0x01, 0x3e, 0x09, 0x10, 0x28, 0xec, 0xc4, 0x14, 0x2c, 0xe5, 0x3f, 0xea, 0x70,
	0x3a, 0x7c, 0xe9, 0x9f, 0x78, 0x62, 0x8a, 0xb7, 0x00, 0xb1, 0x58, 0x8b, 0x28, 0x95, 0x2d, 0xa2,
	0x9a, 0x65, 0x14, 0x40, 0x8c, 0xba, 0x4b, 0x88, 0x96, 0x29, 0xb3, 0xce, 0xa8, 0x6c, 0x4a, 0xbf,
	0xf2, 0xa7, 0xa6, 0x96, 0xe9, 0x08, 0x01, 0xb8, 0x73, 0xa3, 0x45, 0x11, 0xe3, 0xdb, 0xf5, 0xcc,
	0xbb, 0x03, 0x79, 0x23, 0xa0, 0x74, 0xa5, 0xbc, 0x70, 0xec, 0xb6, 0xfc, 0x2d, 0xec, 0x42, 0xde,
	0x70, 0xf2, 0x2e, 0x06, 0x3f, 0x84, 0xb4, 0xfd, 0x7e, 0x72, 0x3f, 0xf3, 0x21, 0xe0, 0x01, 0x69,
	0x60, 0x80, 0x91, 0x73, 0x85, 0x02, 0x74, 0x42, 0x9f, 0x70, 0x03, 0x67, 0x27, 0x32, 0x08, 0x4a,
	0x27, 0x53, 0x26, 0xe4, 0x5c, 0x71, 0xd4, 0xbd, 0x1b, 0xe6, 0xdc, 0x64, 0xfa, 0x72, 0xc1, 0xe0,
	0x71, 0x67, 0x17, 0x4b, 0x1d, 0x7b, 0xfe, 0x71, 0x67, 0x17, 0x0b, 0x11, 0x77, 0xc9, 0x75, 0xc8,
	0x31, 0x56, 0xce, 0xd2, 0xbe, 0x9f, 0x14, 0xe9, 0x64, 0xfa, 0x42, 0x39, 0x0b, 0xfa, 0xca, 0x6c,
	0xaa, 0x32, 0x29, 0x0d, 0x2b, 0xac, 0xa4, 0x03, 0xaf, 0xef, 0x82, 0x7d, 0x6f, 0x65, 0x44, 0xc9,
	0x86, 0xcd, 0xd2, 0x5c, 0x2b, 0x41, 0x8f, 0xfd, 0xc5, 0xe1, 0x27, 0xe4, 0x85, 0xa1, 0x14, 0x6b,
	0x21, 0xe9, 0x89, 0x77, 0x99, 0x94, 0x9f, 0x6a, 0x21, 0x0f, 0x7e, 0x5b, 0x23, 0xbb, 0x1f, 0xfc,
	0xca, 0x83, 0x03, 0x0f, 0xa6, 0x5b, 0xcc, 0x8d, 0xf0, 0x12, 0x41, 0x71, 0x1a, 0xc3, 0xba, 0xca,
	0xe6, 0xa7, 0xdc, 0x08, 0x54, 0x68, 0x8f, 0xd4, 0x2b, 0x97, 0xd0, 0xce, 0xfe, 0xda, 0x61, 0x63,
	0x48, 0x3c, 0xc2, 0xe1, 0x7e, 0x48, 0xb6, 0xab, 0x5f, 0x8e, 0xd8, 0x55, 0x5d, 0xcc, 0xb3, 0x55,
	0x46, 0x61, 0x67, 0x41, 0x79, 0xb8, 0x0d, 0x0f, 0x82, 0xe5, 0x69, 0x0c, 0x37, 0x67, 0xdc, 0xbe,
	0xf2, 0x5f, 0x3c, 0xff, 0xd4, 0xc8, 0x7a, 0xf0, 0xf5, 0xf1, 0xa0, 0xf8, 0xf3, 0xe3, 0x3e, 0x28,
	0x2a, 0x5f, 0x73, 0xc3, 0xf5, 0xd5, 0x5f, 0x75, 0x83, 0xff, 0x7c, 0xd5, 0x8d, 0xd7, 0xf1, 0x9f,
	0x8b, 0xce, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x92, 0xaa, 0x37, 0x7f, 0x0c, 0x00, 0x00,
}
