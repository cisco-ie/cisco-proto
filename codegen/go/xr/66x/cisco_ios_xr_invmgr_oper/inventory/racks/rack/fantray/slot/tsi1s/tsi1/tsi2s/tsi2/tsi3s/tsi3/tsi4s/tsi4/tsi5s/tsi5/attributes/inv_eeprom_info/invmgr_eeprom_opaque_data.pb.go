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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_attributes_inv_eeprom_info

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
	EciAlphaNumber       string     `protobuf:"bytes,58,opt,name=eci_alpha_number,json=eciAlphaNumber,proto3" json:"eci_alpha_number,omitempty"`
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

func (m *DiagEeprom) GetEciAlphaNumber() string {
	if m != nil {
		return m.EciAlphaNumber
	}
	return ""
}

type InvmgrEepromOpaqueData struct {
	InvCardType            uint32      `protobuf:"varint,50,opt,name=inv_card_type,json=invCardType,proto3" json:"inv_card_type,omitempty"`
	OpaqueData             []uint32    `protobuf:"varint,51,rep,packed,name=opaque_data,json=opaqueData,proto3" json:"opaque_data,omitempty"`
	OpaqueDataSize         uint32      `protobuf:"varint,52,opt,name=opaque_data_size,json=opaqueDataSize,proto3" json:"opaque_data_size,omitempty"`
	HasEeprom              uint32      `protobuf:"varint,53,opt,name=has_eeprom,json=hasEeprom,proto3" json:"has_eeprom,omitempty"`
	Eeprom                 *DiagEeprom `protobuf:"bytes,54,opt,name=eeprom,proto3" json:"eeprom,omitempty"`
	Description            string      `protobuf:"bytes,55,opt,name=description,proto3" json:"description,omitempty"`
	FormFactor             uint32      `protobuf:"varint,56,opt,name=form_factor,json=formFactor,proto3" json:"form_factor,omitempty"`
	ConnectorType          uint32      `protobuf:"varint,57,opt,name=connector_type,json=connectorType,proto3" json:"connector_type,omitempty"`
	OtnApplicationCode     uint32      `protobuf:"varint,58,opt,name=otn_application_code,json=otnApplicationCode,proto3" json:"otn_application_code,omitempty"`
	SonetApplicationCode   uint32      `protobuf:"varint,59,opt,name=sonet_application_code,json=sonetApplicationCode,proto3" json:"sonet_application_code,omitempty"`
	EthernetComplianceCode uint32      `protobuf:"varint,60,opt,name=ethernet_compliance_code,json=ethernetComplianceCode,proto3" json:"ethernet_compliance_code,omitempty"`
	DateString             string      `protobuf:"bytes,61,opt,name=date_string,json=dateString,proto3" json:"date_string,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}    `json:"-"`
	XXX_unrecognized       []byte      `json:"-"`
	XXX_sizecache          int32       `json:"-"`
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

func (m *InvmgrEepromOpaqueData) GetFormFactor() uint32 {
	if m != nil {
		return m.FormFactor
	}
	return 0
}

func (m *InvmgrEepromOpaqueData) GetConnectorType() uint32 {
	if m != nil {
		return m.ConnectorType
	}
	return 0
}

func (m *InvmgrEepromOpaqueData) GetOtnApplicationCode() uint32 {
	if m != nil {
		return m.OtnApplicationCode
	}
	return 0
}

func (m *InvmgrEepromOpaqueData) GetSonetApplicationCode() uint32 {
	if m != nil {
		return m.SonetApplicationCode
	}
	return 0
}

func (m *InvmgrEepromOpaqueData) GetEthernetComplianceCode() uint32 {
	if m != nil {
		return m.EthernetComplianceCode
	}
	return 0
}

func (m *InvmgrEepromOpaqueData) GetDateString() string {
	if m != nil {
		return m.DateString
	}
	return ""
}

func init() {
	proto.RegisterType((*InvmgrEepromOpaqueData_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.inv_eeprom_info.invmgr_eeprom_opaque_data_KEYS")
	proto.RegisterType((*RmaDetail)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.inv_eeprom_info.rma_detail")
	proto.RegisterType((*DiagEeprom)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.inv_eeprom_info.diag_eeprom")
	proto.RegisterType((*InvmgrEepromOpaqueData)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.attributes.inv_eeprom_info.invmgr_eeprom_opaque_data")
}

func init() { proto.RegisterFile("invmgr_eeprom_opaque_data.proto", fileDescriptor_6030716c6d5948f0) }

var fileDescriptor_6030716c6d5948f0 = []byte{
	// 1385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xdd, 0x72, 0x1b, 0x35,
	0x14, 0x1e, 0x93, 0x36, 0x69, 0xe4, 0x3a, 0x4d, 0x96, 0x36, 0x55, 0x80, 0x36, 0x69, 0xda, 0xd2,
	0xb4, 0x01, 0x53, 0xff, 0x24, 0x69, 0x0b, 0x5c, 0xa4, 0x69, 0x3b, 0x74, 0xa0, 0x19, 0xc6, 0x29,
	0x9d, 0xe1, 0x4a, 0x23, 0x4b, 0xb2, 0x2d, 0xb2, 0xbb, 0x5a, 0x24, 0xad, 0x43, 0x78, 0x05, 0x78,
	0x00, 0x78, 0x06, 0xee, 0x78, 0x00, 0x1e, 0x81, 0x67, 0x62, 0xce, 0xd1, 0xae, 0x77, 0xa7, 0xb8,
	0xf7, 0xbd, 0xd9, 0xc8, 0xdf, 0xf7, 0xe9, 0xe8, 0xe8, 0x9c, 0xa3, 0x23, 0x85, 0x6c, 0xea, 0x74,
	0x9a, 0x8c, 0x2d, 0x53, 0x2a, 0xb3, 0x26, 0x61, 0x26, 0xe3, 0x3f, 0xe7, 0x8a, 0x49, 0xee, 0x79,
	0x3b, 0xb3, 0xc6, 0x9b, 0xe8, 0xf7, 0x86, 0xd0, 0x4e, 0x18, 0xa6, 0x8d, 0x63, 0xbf, 0x58, 0x56,
	0xc8, 0x4d, 0xa6, 0x6c, 0x5b, 0xa7, 0x53, 0x95, 0x7a, 0x63, 0xcf, 0xdb, 0x96, 0x8b, 0x53, 0x87,
	0xdf, 0xf6, 0x88, 0xa7, 0xde, 0xf2, 0xf3, 0xb6, 0x8b, 0x8d, 0x6f, 0x7b, 0xa7, 0x3b, 0x0e, 0xbf,
	0xf0, 0xe9, 0xe2, 0xb0, 0x0b, 0x9f, 0x1e, 0x0e, 0x7b, 0xf0, 0xe9, 0xe3, 0xb0, 0x0f, 0x9f, 0x3d,
	0x1c, 0xee, 0xb5, 0xb9, 0xf7, 0x56, 0x0f, 0x73, 0xaf, 0x1c, 0xac, 0x51, 0xfa, 0xa6, 0xd3, 0x91,
	0xd9, 0xfe, 0xa7, 0x41, 0x6e, 0xbe, 0xd3, 0x65, 0xf6, 0xed, 0xf3, 0x1f, 0x4f, 0xa2, 0x88, 0x5c,
	0x48, 0x79, 0xa2, 0x68, 0x63, 0xab, 0xb1, 0xb3, 0x3c, 0xc0, 0x71, 0x74, 0x8d, 0x2c, 0xc2, 0x5f,
	0xd6, 0xa1, 0x1f, 0x20, 0x7a, 0x11, 0x7e, 0x75, 0x66, 0x70, 0x97, 0x2e, 0x54, 0x70, 0x77, 0x06,
	0xf7, 0xe8, 0x85, 0x0a, 0xee, 0xcd, 0xe0, 0x3e, 0xbd, 0x58, 0xc1, 0xfd, 0x19, 0xbc, 0x47, 0x17,
	0x2b, 0x78, 0x6f, 0x06, 0xef, 0xd3, 0xa5, 0x0a, 0xde, 0xdf, 0x36, 0x84, 0xd8, 0x84, 0x33, 0xa9,
	0x3c, 0xd7, 0x71, 0x74, 0x8b, 0x5c, 0xf6, 0xca, 0x79, 0x36, 0xd1, 0x0e, 0x02, 0x59, 0xf8, 0xdc,
	0x04, 0xec, 0x9b, 0x00, 0x45, 0x37, 0xc2, 0x84, 0x34, 0x4f, 0x86, 0xca, 0x16, 0xee, 0x2f, 0xdb,
	0x84, 0x1f, 0x23, 0x10, 0x6d, 0x92, 0x26, 0xd0, 0xa5, 0x81, 0xb0, 0x0f, 0x98, 0x51, 0xcc, 0xdf,
	0xfe, 0x77, 0x95, 0x34, 0xa5, 0xe6, 0xe3, 0x22, 0x5e, 0xd1, 0x16, 0x69, 0x4a, 0xe5, 0x84, 0xd5,
	0x99, 0xd7, 0x26, 0x2d, 0x57, 0xac, 0x41, 0xd1, 0x03, 0xb2, 0xa6, 0x25, 0xc6, 0x76, 0x64, 0x6c,
	0xc2, 0x3d, 0xb3, 0x6a, 0x5a, 0x2c, 0x7c, 0x25, 0x10, 0x2f, 0x10, 0x1f, 0xa8, 0x69, 0xb4, 0x4b,
	0xd6, 0x84, 0x49, 0xbd, 0x35, 0x71, 0xac, 0x2c, 0x1b, 0xf1, 0x44, 0xc7, 0xa5, 0x13, 0xab, 0x15,
	0xf1, 0x02, 0xf1, 0xe8, 0x1e, 0xb9, 0x52, 0x13, 0xfb, 0xf3, 0x4c, 0x15, 0x01, 0x5e, 0xa9, 0xe0,
	0xd7, 0xe7, 0x99, 0x8a, 0x56, 0xc9, 0xc2, 0x54, 0xcb, 0x22, 0xcc, 0x30, 0x84, 0xa4, 0x4e, 0xce,
	0xb4, 0x2c, 0x42, 0x8c, 0x63, 0x50, 0x65, 0x5a, 0x16, 0xe1, 0x85, 0x21, 0x2c, 0x90, 0x4b, 0xcd,
	0xea, 0xfb, 0xbb, 0x14, 0x16, 0xc8, 0xa5, 0x7e, 0x56, 0xdb, 0xe2, 0x06, 0xb9, 0x04, 0x42, 0xac,
	0x93, 0x65, 0x54, 0x2c, 0xe5, 0x52, 0x1f, 0x43, 0xa9, 0x44, 0xe4, 0x82, 0x88, 0x95, 0xa6, 0x24,
	0xac, 0x04, 0x63, 0x58, 0x49, 0x09, 0x4d, 0x9b, 0x61, 0x25, 0x25, 0x74, 0xb4, 0x4b, 0x22, 0x6f,
	0x32, 0xc6, 0x9d, 0x53, 0x09, 0xcb, 0xb8, 0xf5, 0x90, 0x20, 0x7a, 0x39, 0x04, 0xc9, 0x9b, 0xec,
	0x10, 0x88, 0xef, 0xb9, 0xf5, 0xc7, 0x79, 0x12, 0x6d, 0x93, 0x56, 0x25, 0x86, 0x8d, 0xb5, 0x8a,
	0x34, 0x17, 0xba, 0x37, 0x5a, 0x46, 0xd7, 0xc9, 0x52, 0x26, 0x30, 0xcd, 0x74, 0x05, 0xd9, 0xc5,
	0x4c, 0x40, 0x8e, 0xa3, 0x75, 0x02, 0x23, 0x98, 0x75, 0x65, 0x86, 0x43, 0x44, 0x36, 0x49, 0x53,
	0x4c, 0xb8, 0x73, 0xda, 0x31, 0xa7, 0x25, 0x5d, 0x0d, 0x89, 0x2f, 0xa0, 0x13, 0x2d, 0x61, 0x8f,
	0x52, 0x4d, 0xc1, 0x62, 0x87, 0xae, 0x85, 0x3d, 0x4a, 0x35, 0x3d, 0xce, 0x93, 0x4e, 0x8d, 0xea,
	0xd2, 0xa8, 0x4e, 0x75, 0x6b, 0x54, 0x8f, 0x7e, 0x58, 0xa7, 0x7a, 0x35, 0xaa, 0x4f, 0xaf, 0xd6,
	0xa9, 0x7e, 0x8d, 0xda, 0xa3, 0xd7, 0xea, 0xd4, 0x5e, 0x8d, 0xda, 0xa7, 0xeb, 0x75, 0x6a, 0xbf,
	0x46, 0x1d, 0xd0, 0xeb, 0x75, 0xea, 0x20, 0xba, 0x43, 0x56, 0x12, 0x9e, 0xe6, 0x0c, 0x4f, 0x07,
	0x9c, 0x6d, 0x4a, 0x51, 0x70, 0x19, 0xd0, 0xd7, 0xca, 0xf9, 0x67, 0xdc, 0xf3, 0xe8, 0xaf, 0x06,
	0x59, 0xb0, 0x09, 0xa7, 0x1b, 0x5b, 0x8d, 0x9d, 0x66, 0xf7, 0x8f, 0x46, 0xfb, 0x7d, 0xea, 0x55,
	0xed, 0xea, 0x9c, 0x0f, 0xc0, 0x4b, 0xd8, 0x2e, 0x94, 0x80, 0x67, 0x5a, 0xd2, 0x8f, 0xc2, 0x76,
	0xf1, 0xf7, 0x4b, 0x4c, 0x66, 0xa0, 0x78, 0xac, 0xb9, 0xa3, 0x1f, 0x87, 0x64, 0x22, 0x74, 0x08,
	0x08, 0x9c, 0xc9, 0x21, 0x77, 0x8a, 0x25, 0x5c, 0x30, 0x2e, 0xa5, 0x55, 0xce, 0x75, 0xe8, 0x27,
	0xa1, 0xdc, 0x80, 0x78, 0xc5, 0xc5, 0x61, 0x01, 0x47, 0xf7, 0xc9, 0x5a, 0x21, 0x63, 0xc3, 0xf8,
	0x94, 0x39, 0xfd, 0xab, 0xea, 0xd0, 0x1b, 0xe1, 0x1c, 0x24, 0xa8, 0x7b, 0x1a, 0x9f, 0x9e, 0x00,
	0x3a, 0xcf, 0x6c, 0x97, 0xde, 0x9c, 0x67, 0xb6, 0x3b, 0xcf, 0x6c, 0x97, 0x6e, 0xce, 0x31, 0xdb,
	0x9d, 0x67, 0xb6, 0x47, 0xb7, 0xe6, 0x99, 0xed, 0xcd, 0x33, 0xdb, 0xa3, 0xb7, 0xe6, 0x98, 0xed,
	0xcd, 0x33, 0xdb, 0xa7, 0xdb, 0xf3, 0xcc, 0xf6, 0xe7, 0x99, 0xed, 0xd3, 0xdb, 0x73, 0xcc, 0xf6,
	0xa1, 0xd6, 0x32, 0x31, 0x64, 0x4e, 0x59, 0xcd, 0x63, 0x3c, 0x81, 0x77, 0x42, 0xad, 0x65, 0x62,
	0x78, 0x82, 0x20, 0x9c, 0xc3, 0x07, 0x64, 0x2d, 0x33, 0x67, 0xca, 0x32, 0x97, 0x67, 0x59, 0x7c,
	0x1e, 0xda, 0xd7, 0xdd, 0xb0, 0x38, 0x12, 0x27, 0x88, 0x63, 0xff, 0xda, 0x2d, 0xb5, 0xc2, 0xa4,
	0x2e, 0x4f, 0x42, 0x27, 0xfa, 0x34, 0x74, 0x45, 0x24, 0x8e, 0x2a, 0x1c, 0x9a, 0xd6, 0x30, 0x36,
	0x02, 0x7c, 0x1c, 0xa7, 0xdc, 0xe7, 0x56, 0xd1, 0x7b, 0xc1, 0x4f, 0x84, 0x4f, 0x4a, 0x34, 0xba,
	0x4d, 0x5a, 0x41, 0x38, 0x55, 0xd6, 0x81, 0xc5, 0x9d, 0xe0, 0x26, 0x82, 0x6f, 0x02, 0x06, 0x37,
	0x4a, 0x10, 0xc5, 0x2a, 0x1d, 0xfb, 0x09, 0xbd, 0x1f, 0x5a, 0x0d, 0x62, 0xdf, 0x21, 0x14, 0xdd,
	0x25, 0xc1, 0x32, 0x13, 0x13, 0x25, 0x4e, 0x5d, 0x9e, 0xd0, 0x07, 0x28, 0x0a, 0xd6, 0x8f, 0x0a,
	0x10, 0x6a, 0xb2, 0xa8, 0x66, 0x08, 0x1e, 0xdd, 0x0d, 0x35, 0x19, 0x20, 0x08, 0x1c, 0x08, 0x0a,
	0x3b, 0x26, 0x4f, 0x3d, 0xfd, 0x2c, 0x08, 0x82, 0x11, 0x40, 0x20, 0xb0, 0x23, 0x9b, 0xb3, 0x84,
	0xff, 0x64, 0x8a, 0x76, 0xff, 0x79, 0xf0, 0x78, 0x64, 0xf3, 0x57, 0x00, 0x62, 0xb0, 0x4a, 0x95,
	0x4e, 0x4b, 0x55, 0xbb, 0x52, 0x01, 0x88, 0xaa, 0x1b, 0x84, 0x18, 0x95, 0x30, 0xe7, 0xad, 0x4e,
	0xc7, 0xf4, 0x8b, 0x70, 0x0d, 0x1a, 0x95, 0x9c, 0x20, 0x00, 0x74, 0x66, 0x8d, 0xcc, 0x05, 0x9e,
	0xae, 0x87, 0x81, 0x2e, 0x90, 0x97, 0x12, 0x42, 0x57, 0xa5, 0x17, 0xee, 0xd1, 0x4e, 0x58, 0xc2,
	0x95, 0xe9, 0x2d, 0xae, 0xd2, 0xb2, 0x93, 0x83, 0xa4, 0x1b, 0xf6, 0x93, 0x85, 0x26, 0x0e, 0x82,
	0xdb, 0xa4, 0x85, 0x02, 0xab, 0xa6, 0x1a, 0x13, 0xd0, 0x2b, 0xea, 0x84, 0x5b, 0xb8, 0x0c, 0x11,
	0x03, 0x51, 0x32, 0x1a, 0x33, 0xa9, 0xa6, 0x9a, 0x63, 0xde, 0xfb, 0x45, 0xe3, 0x1a, 0x8d, 0x9f,
	0x95, 0x18, 0xb8, 0x3b, 0x39, 0x9b, 0xe5, 0x71, 0x2f, 0xb8, 0x3b, 0x39, 0x2b, 0x93, 0xb8, 0x41,
	0x2e, 0x81, 0x8d, 0xa1, 0xf6, 0x8e, 0xee, 0x87, 0x4e, 0x91, 0x8c, 0xc6, 0x4f, 0xb5, 0x77, 0x90,
	0x5f, 0x95, 0x8e, 0x75, 0xaa, 0x94, 0x65, 0xb9, 0x53, 0xf4, 0x20, 0xe4, 0xb7, 0xc4, 0x7e, 0x70,
	0x2a, 0xa2, 0x64, 0xc9, 0xa5, 0x49, 0x66, 0xb4, 0xa4, 0x8f, 0xc2, 0xe4, 0xe2, 0x27, 0xd8, 0x85,
	0xa6, 0x24, 0x8c, 0x54, 0xf4, 0x71, 0xa0, 0x6c, 0xc2, 0x8f, 0x8c, 0x54, 0xd1, 0x0e, 0x59, 0x55,
	0x42, 0x33, 0x1e, 0x67, 0x93, 0xd9, 0x63, 0xe3, 0x49, 0x28, 0x43, 0x25, 0xf4, 0x21, 0xc0, 0x21,
	0x0a, 0xdb, 0xbf, 0x5d, 0x24, 0x1b, 0xef, 0x7c, 0x82, 0xc1, 0x5d, 0x07, 0x7d, 0x50, 0x70, 0x2b,
	0x43, 0x32, 0x21, 0x8c, 0xad, 0x41, 0x53, 0xa7, 0xd3, 0x23, 0x6e, 0x25, 0xe6, 0x72, 0x93, 0x34,
	0x6b, 0x53, 0x68, 0x6f, 0x6b, 0x61, 0xa7, 0x35, 0x20, 0x01, 0xc2, 0xbe, 0xbe, 0x43, 0x56, 0xeb,
	0xcf, 0x3a, 0xac, 0xbf, 0x3e, 0xda, 0x59, 0xa9, 0x54, 0x58, 0x83, 0x10, 0x48, 0xee, 0x0a, 0x47,
	0x30, 0x90, 0xad, 0xc1, 0xf2, 0x84, 0xbb, 0xe7, 0xe1, 0xb1, 0xf3, 0x77, 0x83, 0x2c, 0x16, 0xdc,
	0x3e, 0xde, 0x11, 0x7f, 0xbe, 0x67, 0x77, 0x44, 0xed, 0x65, 0x36, 0x58, 0x9c, 0xff, 0x42, 0x3b,
	0xf8, 0xff, 0x0b, 0x6d, 0x93, 0x34, 0xe1, 0x69, 0xc6, 0x46, 0x5c, 0x78, 0x63, 0x31, 0xcb, 0xad,
	0x01, 0x01, 0xe8, 0x05, 0x22, 0x70, 0xc4, 0x85, 0x49, 0x53, 0x05, 0x3f, 0x42, 0x1a, 0x1e, 0xa3,
	0xa6, 0x35, 0x43, 0x31, 0x11, 0x0f, 0xc9, 0x55, 0xe3, 0x53, 0xc6, 0xb3, 0x2c, 0xd6, 0x02, 0x2b,
	0x33, 0xd4, 0xc6, 0x13, 0x14, 0x47, 0xc6, 0xa7, 0x87, 0x15, 0x85, 0x65, 0xd2, 0x27, 0xeb, 0xce,
	0xa4, 0x70, 0x51, 0xbd, 0x3d, 0xe7, 0x4b, 0x9c, 0x73, 0x15, 0xd9, 0xb7, 0x67, 0x3d, 0x22, 0x54,
	0xf9, 0x89, 0xb2, 0x30, 0x51, 0x98, 0x24, 0x8b, 0x35, 0x4f, 0x85, 0x0a, 0xf3, 0xbe, 0xc2, 0x79,
	0xeb, 0x25, 0x7f, 0x34, 0xa3, 0x71, 0xe6, 0x26, 0x69, 0x4a, 0xee, 0x55, 0x79, 0xee, 0xbf, 0x0e,
	0x67, 0x12, 0xa0, 0x70, 0xf0, 0x87, 0x8b, 0xf8, 0x6f, 0x4a, 0xef, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x32, 0x52, 0x82, 0x3a, 0xc9, 0x0c, 0x00, 0x00,
}