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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_attributes_inv_eeprom_info

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
	proto.RegisterType((*InvmgrEepromOpaqueData_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.inv_eeprom_info.invmgr_eeprom_opaque_data_KEYS")
	proto.RegisterType((*RmaDetail)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.inv_eeprom_info.rma_detail")
	proto.RegisterType((*DiagEeprom)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.inv_eeprom_info.diag_eeprom")
	proto.RegisterType((*InvmgrEepromOpaqueData)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.inv_eeprom_info.invmgr_eeprom_opaque_data")
}

func init() { proto.RegisterFile("invmgr_eeprom_opaque_data.proto", fileDescriptor_6030716c6d5948f0) }

var fileDescriptor_6030716c6d5948f0 = []byte{
	// 1362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xcf, 0x73, 0x13, 0xb7,
	0x17, 0x1f, 0x7f, 0x13, 0x12, 0x22, 0xe3, 0x90, 0xec, 0x17, 0x82, 0xd2, 0x16, 0x1c, 0x02, 0x94,
	0x40, 0x5a, 0x17, 0xff, 0x20, 0x01, 0xda, 0x1e, 0x42, 0x20, 0x53, 0xa6, 0x25, 0xd3, 0xb1, 0x29,
	0x33, 0x3d, 0x69, 0x64, 0x49, 0xb6, 0xd5, 0xec, 0xae, 0xb6, 0x92, 0xd6, 0x69, 0x7a, 0x6f, 0xcf,
	0x3d, 0xf6, 0xd6, 0x3f, 0xa4, 0xff, 0x58, 0x8f, 0x9d, 0xf7, 0xb4, 0x6b, 0xef, 0x50, 0x73, 0xe6,
	0xe2, 0x68, 0x3f, 0x9f, 0x8f, 0x9e, 0x9e, 0xde, 0x7b, 0x7a, 0x52, 0x48, 0x53, 0xa7, 0xd3, 0x64,
	0x6c, 0x99, 0x52, 0x99, 0x35, 0x09, 0x33, 0x19, 0xff, 0x39, 0x57, 0x4c, 0x72, 0xcf, 0x5b, 0x99,
	0x35, 0xde, 0x44, 0x17, 0x42, 0x3b, 0x61, 0x98, 0x36, 0x8e, 0xfd, 0x62, 0x59, 0xa1, 0x36, 0x99,
	0xb2, 0x2d, 0x9d, 0x4e, 0x55, 0xea, 0x8d, 0xbd, 0x68, 0x59, 0x2e, 0xce, 0x1c, 0xfe, 0xb6, 0x32,
	0x73, 0xae, 0xac, 0x9b, 0xa8, 0x78, 0xd4, 0x72, 0xb1, 0xf1, 0x2d, 0xef, 0x74, 0xdb, 0xe1, 0x2f,
	0xfc, 0x74, 0x70, 0xd8, 0x81, 0x9f, 0x2e, 0x0e, 0xbb, 0x2d, 0xee, 0xbd, 0xd5, 0xc3, 0xdc, 0x2b,
	0x07, 0x06, 0x4b, 0x3f, 0x74, 0x3a, 0x32, 0xbb, 0x7f, 0xd4, 0xc8, 0xad, 0xf7, 0xba, 0xc7, 0xbe,
	0x7d, 0xf9, 0xe3, 0x20, 0x8a, 0xc8, 0x72, 0xca, 0x13, 0x45, 0x6b, 0x3b, 0xb5, 0xbd, 0xb5, 0x3e,
	0x8e, 0xa3, 0xeb, 0x64, 0x05, 0xfe, 0xb2, 0x36, 0xfd, 0x1f, 0xa2, 0x97, 0xe0, 0xab, 0x3d, 0x83,
	0x3b, 0x74, 0x69, 0x0e, 0x77, 0x66, 0x70, 0x97, 0x2e, 0xcf, 0xe1, 0xee, 0x0c, 0xee, 0xd1, 0x4b,
	0x73, 0xb8, 0xb7, 0x6b, 0x08, 0xb1, 0x09, 0x67, 0x52, 0x79, 0xae, 0xe3, 0xe8, 0x36, 0xb9, 0xe2,
	0x95, 0xf3, 0x6c, 0xa2, 0x1d, 0xc4, 0xa1, 0xf0, 0xa2, 0x0e, 0xd8, 0x37, 0x01, 0x8a, 0x6e, 0x86,
	0x09, 0x69, 0x9e, 0x0c, 0x95, 0x2d, 0x1c, 0x5a, 0xb3, 0x09, 0x3f, 0x45, 0x20, 0x6a, 0x92, 0x3a,
	0xd0, 0xa5, 0x81, 0xe0, 0x19, 0xcc, 0x28, 0xe6, 0xef, 0xfe, 0xbd, 0x41, 0xea, 0x52, 0xf3, 0x71,
	0x11, 0x81, 0x68, 0x87, 0xd4, 0xa5, 0x72, 0xc2, 0xea, 0xcc, 0x6b, 0x93, 0x96, 0x2b, 0x56, 0xa0,
	0xe8, 0x21, 0xd9, 0xd4, 0x12, 0xa3, 0x35, 0x32, 0x36, 0xe1, 0x9e, 0x59, 0x35, 0x2d, 0x16, 0xbe,
	0x1a, 0x88, 0x13, 0xc4, 0xfb, 0x6a, 0x1a, 0xed, 0x93, 0x4d, 0x61, 0x52, 0x6f, 0x4d, 0x1c, 0x2b,
	0xcb, 0x46, 0x3c, 0xd1, 0x71, 0xe9, 0xc4, 0xc6, 0x9c, 0x38, 0x41, 0x3c, 0xba, 0x4f, 0xae, 0x56,
	0xc4, 0xfe, 0x22, 0x53, 0x45, 0xc8, 0xd6, 0xe7, 0xf0, 0x9b, 0x8b, 0x4c, 0x45, 0x1b, 0x64, 0x69,
	0xaa, 0x65, 0x11, 0x38, 0x18, 0x42, 0x9a, 0x26, 0xe7, 0x5a, 0xd2, 0x95, 0x90, 0x26, 0x18, 0x83,
	0x2a, 0xd3, 0x92, 0xae, 0x06, 0x55, 0xa6, 0x25, 0x2c, 0x90, 0x4b, 0xcd, 0xaa, 0xfb, 0xbb, 0x1c,
	0x16, 0xc8, 0xa5, 0x7e, 0x51, 0xd9, 0xe2, 0x36, 0xb9, 0x0c, 0x42, 0xcc, 0xfc, 0x1a, 0x2a, 0x56,
	0x73, 0xa9, 0x4f, 0x21, 0xf9, 0x11, 0x59, 0x16, 0xb1, 0xd2, 0x94, 0x84, 0x95, 0x60, 0x0c, 0x2b,
	0x29, 0xa1, 0x69, 0x3d, 0xac, 0xa4, 0x84, 0x8e, 0xf6, 0x49, 0xe4, 0x4d, 0xc6, 0xb8, 0x73, 0x2a,
	0x61, 0x19, 0xb7, 0x1e, 0x12, 0x44, 0xaf, 0x84, 0x20, 0x79, 0x93, 0x1d, 0x01, 0xf1, 0x3d, 0xb7,
	0xfe, 0x34, 0x4f, 0xa2, 0x5d, 0xd2, 0x98, 0x8b, 0x61, 0x63, 0x8d, 0x22, 0xcd, 0x85, 0xee, 0xad,
	0x96, 0xd1, 0x0d, 0xb2, 0x9a, 0x09, 0x4c, 0x33, 0x5d, 0x47, 0x76, 0x25, 0x13, 0x90, 0xe3, 0x68,
	0x8b, 0xc0, 0x08, 0x66, 0x5d, 0x9d, 0xe1, 0x10, 0x91, 0x26, 0xa9, 0x8b, 0x09, 0x77, 0x4e, 0x3b,
	0xe6, 0xb4, 0xa4, 0x1b, 0x21, 0xf1, 0x05, 0x34, 0xd0, 0x12, 0xf6, 0x28, 0xd5, 0x14, 0x2c, 0xb6,
	0xe9, 0x66, 0xd8, 0xa3, 0x54, 0xd3, 0xd3, 0x3c, 0x69, 0x57, 0xa8, 0x0e, 0x8d, 0xaa, 0x54, 0xa7,
	0x42, 0x75, 0xe9, 0xff, 0xab, 0x54, 0xb7, 0x42, 0xf5, 0xe8, 0xb5, 0x2a, 0xd5, 0xab, 0x50, 0x8f,
	0xe9, 0xf5, 0x2a, 0xf5, 0xb8, 0x42, 0x1d, 0xd0, 0xad, 0x2a, 0x75, 0x50, 0xa1, 0x0e, 0xe9, 0x8d,
	0x2a, 0x75, 0x18, 0xdd, 0x25, 0xeb, 0x09, 0x4f, 0x73, 0x86, 0xa7, 0x03, 0x4e, 0x2b, 0xa5, 0x28,
	0xb8, 0x02, 0xe8, 0x1b, 0xe5, 0xfc, 0x0b, 0xee, 0x79, 0xf4, 0x67, 0x8d, 0x2c, 0xd9, 0x84, 0xd3,
	0xed, 0x9d, 0xda, 0x5e, 0xbd, 0xf3, 0x5b, 0xad, 0xf5, 0xc1, 0x5a, 0x4d, 0x6b, 0x7e, 0xa8, 0xfb,
	0xe0, 0x12, 0xec, 0x0d, 0xf2, 0xed, 0x99, 0x96, 0xf4, 0xa3, 0xb0, 0x37, 0xfc, 0x7e, 0x85, 0x99,
	0x0b, 0x14, 0x8f, 0x35, 0x77, 0xf4, 0xe3, 0x90, 0x39, 0x84, 0x8e, 0x00, 0x81, 0x03, 0x38, 0xe4,
	0x4e, 0xb1, 0x84, 0x0b, 0xc6, 0xa5, 0xb4, 0xca, 0xb9, 0x36, 0xfd, 0x24, 0xd4, 0x16, 0x10, 0xaf,
	0xb9, 0x38, 0x2a, 0xe0, 0xe8, 0x01, 0xd9, 0x2c, 0x64, 0x6c, 0x18, 0x9f, 0x31, 0xa7, 0x7f, 0x55,
	0x6d, 0x7a, 0x33, 0x14, 0x7d, 0x82, 0xba, 0xe7, 0xf1, 0xd9, 0x00, 0xd0, 0x45, 0x66, 0x3b, 0xf4,
	0xd6, 0x22, 0xb3, 0x9d, 0x45, 0x66, 0x3b, 0xb4, 0xb9, 0xc0, 0x6c, 0x67, 0x91, 0xd9, 0x2e, 0xdd,
	0x59, 0x64, 0xb6, 0xbb, 0xc8, 0x6c, 0x97, 0xde, 0x5e, 0x60, 0xb6, 0xbb, 0xc8, 0x6c, 0x8f, 0xee,
	0x2e, 0x32, 0xdb, 0x5b, 0x64, 0xb6, 0x47, 0xef, 0x2c, 0x30, 0xdb, 0x83, 0xc2, 0xca, 0xc4, 0x90,
	0x39, 0x65, 0x35, 0x8f, 0xf1, 0xb8, 0xdd, 0x0d, 0x85, 0x95, 0x89, 0xe1, 0x00, 0x41, 0x38, 0x74,
	0x0f, 0xc9, 0x26, 0x16, 0x07, 0x73, 0x79, 0x96, 0xc5, 0x17, 0xa1, 0x57, 0xdd, 0x0b, 0x8b, 0x23,
	0x31, 0x40, 0x1c, 0x9b, 0xd5, 0x7e, 0xa9, 0x15, 0x26, 0x75, 0x79, 0x12, 0xda, 0xce, 0xa7, 0xa1,
	0x05, 0x22, 0x71, 0x3c, 0xc7, 0xa1, 0x43, 0x0d, 0x63, 0x23, 0xc0, 0xc7, 0x71, 0xca, 0x7d, 0x6e,
	0x15, 0xbd, 0x1f, 0xfc, 0x44, 0x78, 0x50, 0xa2, 0xd1, 0x1d, 0xd2, 0x08, 0xc2, 0xa9, 0xb2, 0x0e,
	0x2c, 0xee, 0x05, 0x37, 0x11, 0x7c, 0x1b, 0x30, 0xb8, 0x3e, 0x82, 0x28, 0x56, 0xe9, 0xd8, 0x4f,
	0xe8, 0x83, 0xd0, 0x57, 0x10, 0xfb, 0x0e, 0xa1, 0xe8, 0x1e, 0x09, 0x96, 0x99, 0x98, 0x28, 0x71,
	0xe6, 0xf2, 0x84, 0x3e, 0x44, 0x51, 0xb0, 0x7e, 0x5c, 0x80, 0x50, 0x93, 0x45, 0x35, 0x43, 0xf0,
	0xe8, 0x7e, 0xa8, 0xc9, 0x00, 0x41, 0xe0, 0x40, 0x50, 0xd8, 0x31, 0x79, 0xea, 0xe9, 0x67, 0x41,
	0x10, 0x8c, 0x00, 0x02, 0x81, 0x1d, 0xd9, 0x9c, 0x25, 0xfc, 0x27, 0x53, 0xf4, 0xf6, 0xcf, 0x83,
	0xc7, 0x23, 0x9b, 0xbf, 0x06, 0x10, 0x83, 0x55, 0xaa, 0x74, 0x5a, 0xaa, 0x5a, 0x73, 0x15, 0x80,
	0xa8, 0xba, 0x49, 0x88, 0x51, 0x09, 0x73, 0xde, 0xea, 0x74, 0x4c, 0xbf, 0x08, 0x77, 0x9e, 0x51,
	0xc9, 0x00, 0x01, 0xa0, 0x33, 0x6b, 0x64, 0x2e, 0xf0, 0x74, 0x3d, 0x0a, 0x74, 0x81, 0xbc, 0x92,
	0x10, 0xba, 0x79, 0x7a, 0xe1, 0xd2, 0x6c, 0x87, 0x25, 0x5c, 0x99, 0xde, 0xe2, 0xde, 0x2c, 0xdb,
	0x36, 0x48, 0x3a, 0x61, 0x3f, 0x59, 0xe8, 0xd8, 0x20, 0xb8, 0x43, 0x1a, 0x28, 0xb0, 0x6a, 0xaa,
	0x31, 0x01, 0xdd, 0xa2, 0x4e, 0xb8, 0x85, 0x9b, 0x0f, 0x31, 0x10, 0x25, 0xa3, 0x31, 0x93, 0x6a,
	0xaa, 0x39, 0xe6, 0xbd, 0x57, 0x74, 0xa9, 0xd1, 0xf8, 0x45, 0x89, 0x81, 0xbb, 0x93, 0xf3, 0x59,
	0x1e, 0x1f, 0x07, 0x77, 0x27, 0xe7, 0x65, 0x12, 0xb7, 0xc9, 0x65, 0xb0, 0x31, 0xd4, 0xde, 0xd1,
	0x83, 0xd0, 0x29, 0x92, 0xd1, 0xf8, 0xb9, 0xf6, 0x0e, 0xf2, 0xab, 0xd2, 0xb1, 0x4e, 0x95, 0xb2,
	0x2c, 0x77, 0x8a, 0x1e, 0x86, 0xfc, 0x96, 0xd8, 0x0f, 0x4e, 0x45, 0x94, 0xac, 0xba, 0x34, 0xc9,
	0x8c, 0x96, 0xf4, 0x49, 0x98, 0x5c, 0x7c, 0x82, 0x5d, 0x68, 0x4a, 0xc2, 0x48, 0x45, 0x9f, 0x06,
	0xca, 0x26, 0xfc, 0xd8, 0x48, 0x15, 0xed, 0x91, 0x0d, 0x25, 0x34, 0xe3, 0x71, 0x36, 0x99, 0xbd,
	0x2c, 0x9e, 0x85, 0x32, 0x54, 0x42, 0x1f, 0x01, 0x1c, 0xa2, 0xb0, 0xfb, 0xcf, 0x32, 0xd9, 0x7e,
	0xef, 0x0b, 0x0a, 0x2e, 0x36, 0xe8, 0x83, 0x82, 0x5b, 0x19, 0x92, 0x09, 0x61, 0x6c, 0xf4, 0xeb,
	0x3a, 0x9d, 0x1e, 0x73, 0x2b, 0x31, 0x97, 0x4d, 0x52, 0xaf, 0x4c, 0xa1, 0xdd, 0x9d, 0xa5, 0xbd,
	0x46, 0x9f, 0x04, 0x08, 0x9b, 0xf8, 0x1e, 0xd9, 0xa8, 0xbe, 0xca, 0xb0, 0xfe, 0x7a, 0x68, 0x67,
	0x7d, 0xae, 0xc2, 0x1a, 0x84, 0x40, 0x72, 0x57, 0x38, 0x82, 0x81, 0x6c, 0xf4, 0xd7, 0x26, 0xdc,
	0xbd, 0x0c, 0x2f, 0x9b, 0xbf, 0x6a, 0x64, 0xa5, 0xe0, 0x0e, 0xf0, 0x42, 0xf8, 0xfd, 0x43, 0x5e,
	0x08, 0x95, 0x37, 0x57, 0x7f, 0x65, 0xf1, 0xdb, 0xeb, 0xf0, 0xbf, 0x6f, 0xaf, 0x26, 0xa9, 0xc3,
	0xa3, 0x8b, 0x8d, 0xb8, 0xf0, 0xc6, 0x62, 0x4a, 0x1b, 0x7d, 0x02, 0xd0, 0x09, 0x22, 0x70, 0x9e,
	0x85, 0x49, 0x53, 0x05, 0x1f, 0x21, 0xe6, 0x4f, 0x51, 0xd3, 0x98, 0xa1, 0x18, 0xf5, 0x47, 0xe4,
	0x9a, 0xf1, 0x29, 0xe3, 0x59, 0x16, 0x6b, 0x81, 0x65, 0x18, 0x0a, 0xe1, 0x19, 0x8a, 0x23, 0xe3,
	0xd3, 0xa3, 0x39, 0x85, 0x35, 0xd1, 0x23, 0x5b, 0xce, 0xa4, 0x70, 0x2b, 0xbd, 0x3b, 0xe7, 0x4b,
	0x9c, 0x73, 0x0d, 0xd9, 0x77, 0x67, 0x3d, 0x21, 0x54, 0xf9, 0x89, 0xb2, 0x30, 0x51, 0x98, 0x24,
	0x8b, 0x35, 0x4f, 0x85, 0x0a, 0xf3, 0xbe, 0xc2, 0x79, 0x5b, 0x25, 0x7f, 0x3c, 0xa3, 0x71, 0x66,
	0x93, 0xd4, 0x25, 0xf7, 0xaa, 0x3c, 0xe4, 0x5f, 0x87, 0x03, 0x08, 0x50, 0x38, 0xe5, 0xc3, 0x15,
	0xfc, 0xf7, 0xa1, 0xfb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x19, 0x80, 0xce, 0x61, 0x0c,
	0x00, 0x00,
}
