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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_attributes_inv_eeprom_info

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
	proto.RegisterType((*InvmgrEepromOpaqueData_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.inv_eeprom_info.invmgr_eeprom_opaque_data_KEYS")
	proto.RegisterType((*RmaDetail)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.inv_eeprom_info.rma_detail")
	proto.RegisterType((*DiagEeprom)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.inv_eeprom_info.diag_eeprom")
	proto.RegisterType((*InvmgrEepromOpaqueData)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.inv_eeprom_info.invmgr_eeprom_opaque_data")
}

func init() { proto.RegisterFile("invmgr_eeprom_opaque_data.proto", fileDescriptor_6030716c6d5948f0) }

var fileDescriptor_6030716c6d5948f0 = []byte{
	// 1361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xdf, 0x72, 0x13, 0xb7,
	0x17, 0x1e, 0xff, 0x12, 0x12, 0x22, 0xe3, 0x90, 0xec, 0x0f, 0x82, 0xd2, 0x16, 0x1c, 0x02, 0x94,
	0x40, 0x5a, 0x17, 0xff, 0x21, 0x01, 0xda, 0x5e, 0x84, 0x40, 0xa6, 0x4c, 0x4b, 0xa6, 0x63, 0x53,
	0x66, 0x7a, 0xa5, 0x91, 0x25, 0xd9, 0x56, 0xb3, 0xbb, 0xda, 0x4a, 0x5a, 0xa7, 0xe9, 0x5d, 0xfb,
	0x04, 0xbd, 0xe8, 0x5d, 0x1f, 0xa4, 0x2f, 0xd6, 0x07, 0xe8, 0x9c, 0xa3, 0x5d, 0x7b, 0x87, 0x9a,
	0x6b, 0x6e, 0x1c, 0xed, 0xf7, 0x7d, 0x3a, 0x3a, 0x3a, 0xe7, 0xe8, 0x48, 0x21, 0x4d, 0x9d, 0x4e,
	0x93, 0xb1, 0x65, 0x4a, 0x65, 0xd6, 0x24, 0xcc, 0x64, 0xfc, 0xe7, 0x5c, 0x31, 0xc9, 0x3d, 0x6f,
	0x65, 0xd6, 0x78, 0x13, 0xe5, 0x42, 0x3b, 0x61, 0x98, 0x36, 0x8e, 0xfd, 0x62, 0x59, 0xa1, 0x36,
	0x99, 0xb2, 0x2d, 0x9d, 0x4e, 0x55, 0xea, 0x8d, 0xbd, 0x68, 0x59, 0x2e, 0xce, 0x1c, 0xfe, 0xb6,
	0x54, 0xea, 0xb5, 0xbf, 0x68, 0xb9, 0xd8, 0xf8, 0x96, 0x77, 0xba, 0xed, 0xf0, 0x17, 0x7e, 0x3a,
	0x38, 0xec, 0xc0, 0x4f, 0x17, 0x87, 0xdd, 0x16, 0xf7, 0xde, 0xea, 0x61, 0xee, 0x95, 0x03, 0x63,
	0xa5, 0x0f, 0x3a, 0x1d, 0x99, 0xdd, 0x3f, 0x6a, 0xe4, 0xd6, 0x7b, 0x5d, 0x63, 0xdf, 0xbe, 0xfc,
	0x71, 0x10, 0x45, 0x64, 0x39, 0xe5, 0x89, 0xa2, 0xb5, 0x9d, 0xda, 0xde, 0x5a, 0x1f, 0xc7, 0xd1,
	0x75, 0xb2, 0x02, 0x7f, 0x59, 0x9b, 0xfe, 0x0f, 0xd1, 0x4b, 0xf0, 0xd5, 0x9e, 0xc1, 0x1d, 0xba,
	0x34, 0x87, 0x3b, 0x33, 0xb8, 0x4b, 0x97, 0xe7, 0x70, 0x77, 0x06, 0xf7, 0xe8, 0xa5, 0x39, 0xdc,
	0xdb, 0x35, 0x84, 0xd8, 0x84, 0x33, 0xa9, 0x3c, 0xd7, 0x71, 0x74, 0x9b, 0x5c, 0xf1, 0xca, 0x79,
	0x36, 0xd1, 0x0e, 0x62, 0x50, 0x78, 0x51, 0x07, 0xec, 0x9b, 0x00, 0x45, 0x37, 0xc3, 0x84, 0x34,
	0x4f, 0x86, 0xca, 0x16, 0x0e, 0xad, 0xd9, 0x84, 0x9f, 0x22, 0x10, 0x35, 0x49, 0x1d, 0xe8, 0xd2,
	0x40, 0xf0, 0x0c, 0x66, 0x14, 0xf3, 0x77, 0xff, 0xde, 0x20, 0x75, 0xa9, 0xf9, 0xb8, 0x88, 0x40,
	0xb4, 0x43, 0xea, 0x52, 0x39, 0x61, 0x75, 0xe6, 0xb5, 0x49, 0xcb, 0x15, 0x2b, 0x50, 0xf4, 0x90,
	0x6c, 0x6a, 0x89, 0xd1, 0x1a, 0x19, 0x9b, 0x70, 0xcf, 0xac, 0x9a, 0x16, 0x0b, 0x5f, 0x0d, 0xc4,
	0x09, 0xe2, 0x7d, 0x35, 0x8d, 0xf6, 0xc9, 0xa6, 0x30, 0xa9, 0xb7, 0x26, 0x8e, 0x95, 0x65, 0x23,
	0x9e, 0xe8, 0xb8, 0x74, 0x62, 0x63, 0x4e, 0x9c, 0x20, 0x1e, 0xdd, 0x27, 0x57, 0x2b, 0x62, 0x7f,
	0x91, 0xa9, 0x22, 0x64, 0xeb, 0x73, 0xf8, 0xcd, 0x45, 0xa6, 0xa2, 0x0d, 0xb2, 0x34, 0xd5, 0xb2,
	0x08, 0x1c, 0x0c, 0x21, 0x4d, 0x93, 0x73, 0x2d, 0xe9, 0x4a, 0x48, 0x13, 0x8c, 0x41, 0x95, 0x69,
	0x49, 0x57, 0x83, 0x2a, 0xd3, 0x12, 0x16, 0xc8, 0xa5, 0x66, 0xd5, 0xfd, 0x5d, 0x0e, 0x0b, 0xe4,
	0x52, 0xbf, 0xa8, 0x6c, 0x71, 0x9b, 0x5c, 0x06, 0x21, 0x66, 0x7e, 0x0d, 0x15, 0xab, 0xb9, 0xd4,
	0xa7, 0x90, 0xfc, 0x88, 0x2c, 0x8b, 0x58, 0x69, 0x4a, 0xc2, 0x4a, 0x30, 0x86, 0x95, 0x94, 0xd0,
	0xb4, 0x1e, 0x56, 0x52, 0x42, 0x47, 0xfb, 0x24, 0xf2, 0x26, 0x63, 0xdc, 0x39, 0x95, 0xb0, 0x8c,
	0x5b, 0x0f, 0x09, 0xa2, 0x57, 0x42, 0x90, 0xbc, 0xc9, 0x8e, 0x80, 0xf8, 0x9e, 0x5b, 0x7f, 0x9a,
	0x27, 0xd1, 0x2e, 0x69, 0xcc, 0xc5, 0xb0, 0xb1, 0x46, 0x91, 0xe6, 0x42, 0xf7, 0x56, 0xcb, 0xe8,
	0x06, 0x59, 0xcd, 0x04, 0xa6, 0x99, 0xae, 0x23, 0xbb, 0x92, 0x09, 0xc8, 0x71, 0xb4, 0x45, 0x60,
	0x04, 0xb3, 0xae, 0xce, 0x70, 0x88, 0x48, 0x93, 0xd4, 0xc5, 0x84, 0x3b, 0xa7, 0x1d, 0x73, 0x5a,
	0xd2, 0x8d, 0x90, 0xf8, 0x02, 0x1a, 0x68, 0x09, 0x7b, 0x94, 0x6a, 0x0a, 0x16, 0xdb, 0x74, 0x33,
	0xec, 0x51, 0xaa, 0xe9, 0x69, 0x9e, 0xb4, 0x2b, 0x54, 0x87, 0x46, 0x55, 0xaa, 0x53, 0xa1, 0xba,
	0xf4, 0xff, 0x55, 0xaa, 0x5b, 0xa1, 0x7a, 0xf4, 0x5a, 0x95, 0xea, 0x55, 0xa8, 0xc7, 0xf4, 0x7a,
	0x95, 0x7a, 0x5c, 0xa1, 0x0e, 0xe8, 0x56, 0x95, 0x3a, 0xa8, 0x50, 0x87, 0xf4, 0x46, 0x95, 0x3a,
	0x8c, 0xee, 0x92, 0xf5, 0x84, 0xa7, 0x39, 0xc3, 0xd3, 0x01, 0xa7, 0x95, 0x52, 0x14, 0x5c, 0x01,
	0xf4, 0x8d, 0x72, 0xfe, 0x05, 0xf7, 0x3c, 0xfa, 0xb3, 0x46, 0x96, 0x6c, 0xc2, 0xe9, 0xf6, 0x4e,
	0x6d, 0xaf, 0xde, 0xf9, 0xad, 0xd6, 0xfa, 0x20, 0x6d, 0xa6, 0x35, 0x3f, 0xd0, 0x7d, 0x70, 0x07,
	0xf6, 0x05, 0xb9, 0xf6, 0x4c, 0x4b, 0xfa, 0x51, 0xd8, 0x17, 0x7e, 0xbf, 0xc2, 0xac, 0x05, 0x8a,
	0xc7, 0x9a, 0x3b, 0xfa, 0x71, 0xc8, 0x1a, 0x42, 0x47, 0x80, 0xc0, 0xe1, 0x1b, 0x72, 0xa7, 0x58,
	0xc2, 0x05, 0xe3, 0x52, 0x5a, 0xe5, 0x5c, 0x9b, 0x7e, 0x12, 0xea, 0x0a, 0x88, 0xd7, 0x5c, 0x1c,
	0x15, 0x70, 0xf4, 0x80, 0x6c, 0x16, 0x32, 0x36, 0x8c, 0xcf, 0x98, 0xd3, 0xbf, 0xaa, 0x36, 0xbd,
	0x19, 0x0a, 0x3e, 0x41, 0xdd, 0xf3, 0xf8, 0x6c, 0x00, 0xe8, 0x22, 0xb3, 0x1d, 0x7a, 0x6b, 0x91,
	0xd9, 0xce, 0x22, 0xb3, 0x1d, 0xda, 0x5c, 0x60, 0xb6, 0xb3, 0xc8, 0x6c, 0x97, 0xee, 0x2c, 0x32,
	0xdb, 0x5d, 0x64, 0xb6, 0x4b, 0x6f, 0x2f, 0x30, 0xdb, 0x5d, 0x64, 0xb6, 0x47, 0x77, 0x17, 0x99,
	0xed, 0x2d, 0x32, 0xdb, 0xa3, 0x77, 0x16, 0x98, 0xed, 0x41, 0x51, 0x65, 0x62, 0xc8, 0x9c, 0xb2,
	0x9a, 0xc7, 0x78, 0xd4, 0xee, 0x86, 0xa2, 0xca, 0xc4, 0x70, 0x80, 0x20, 0x1c, 0xb8, 0x87, 0x64,
	0x33, 0x33, 0xe7, 0xca, 0x32, 0x97, 0x67, 0x59, 0x7c, 0x11, 0xfa, 0xd4, 0xbd, 0xb0, 0x38, 0x12,
	0x03, 0xc4, 0xb1, 0x51, 0xed, 0x97, 0x5a, 0x61, 0x52, 0x97, 0x27, 0xa1, 0xe5, 0x7c, 0x1a, 0xda,
	0x1f, 0x12, 0xc7, 0x73, 0x1c, 0xba, 0xd3, 0x30, 0x36, 0x02, 0x7c, 0x1c, 0xa7, 0xdc, 0xe7, 0x56,
	0xd1, 0xfb, 0xc1, 0x4f, 0x84, 0x07, 0x25, 0x1a, 0xdd, 0x21, 0x8d, 0x20, 0x9c, 0x2a, 0xeb, 0xc0,
	0xe2, 0x5e, 0x70, 0x13, 0xc1, 0xb7, 0x01, 0x83, 0xab, 0x23, 0x88, 0x62, 0x95, 0x8e, 0xfd, 0x84,
	0x3e, 0x08, 0x3d, 0x05, 0xb1, 0xef, 0x10, 0x8a, 0xee, 0x91, 0x60, 0x99, 0x89, 0x89, 0x12, 0x67,
	0x2e, 0x4f, 0xe8, 0x43, 0x14, 0x05, 0xeb, 0xc7, 0x05, 0x08, 0x35, 0x59, 0x54, 0x33, 0x04, 0x8f,
	0xee, 0x87, 0x9a, 0x0c, 0x10, 0x04, 0x0e, 0x04, 0x85, 0x1d, 0x93, 0xa7, 0x9e, 0x7e, 0x16, 0x04,
	0xc1, 0x08, 0x20, 0x10, 0xd8, 0x91, 0xcd, 0x59, 0xc2, 0x7f, 0x32, 0x45, 0x5f, 0xff, 0x3c, 0x78,
	0x3c, 0xb2, 0xf9, 0x6b, 0x00, 0x31, 0x58, 0xa5, 0x4a, 0xa7, 0xa5, 0xaa, 0x35, 0x57, 0x01, 0x88,
	0xaa, 0x9b, 0x84, 0x18, 0x95, 0x30, 0xe7, 0xad, 0x4e, 0xc7, 0xf4, 0x8b, 0x70, 0xdf, 0x19, 0x95,
	0x0c, 0x10, 0x00, 0x3a, 0xb3, 0x46, 0xe6, 0x02, 0x4f, 0xd7, 0xa3, 0x40, 0x17, 0xc8, 0x2b, 0x09,
	0xa1, 0x9b, 0xa7, 0x17, 0x2e, 0xcc, 0x76, 0x58, 0xc2, 0x95, 0xe9, 0x2d, 0xee, 0xcc, 0xb2, 0x65,
	0x83, 0xa4, 0x13, 0xf6, 0x93, 0x85, 0x6e, 0x0d, 0x82, 0x3b, 0xa4, 0x81, 0x02, 0xab, 0xa6, 0x1a,
	0x13, 0xd0, 0x2d, 0xea, 0x84, 0x5b, 0xb8, 0xf5, 0x10, 0x03, 0x51, 0x32, 0x1a, 0x33, 0xa9, 0xa6,
	0x9a, 0x63, 0xde, 0x7b, 0x45, 0x87, 0x1a, 0x8d, 0x5f, 0x94, 0x18, 0xb8, 0x3b, 0x39, 0x9f, 0xe5,
	0xf1, 0x71, 0x70, 0x77, 0x72, 0x5e, 0x26, 0x71, 0x9b, 0x5c, 0x06, 0x1b, 0x43, 0xed, 0x1d, 0x3d,
	0x08, 0x9d, 0x22, 0x19, 0x8d, 0x9f, 0x6b, 0xef, 0x20, 0xbf, 0x2a, 0x1d, 0xeb, 0x54, 0x29, 0xcb,
	0x72, 0xa7, 0xe8, 0x61, 0xc8, 0x6f, 0x89, 0xfd, 0xe0, 0x54, 0x44, 0xc9, 0xaa, 0x4b, 0x93, 0xcc,
	0x68, 0x49, 0x9f, 0x84, 0xc9, 0xc5, 0x27, 0xd8, 0x85, 0xa6, 0x24, 0x8c, 0x54, 0xf4, 0x69, 0xa0,
	0x6c, 0xc2, 0x8f, 0x8d, 0x54, 0xd1, 0x1e, 0xd9, 0x50, 0x42, 0x33, 0x1e, 0x67, 0x93, 0xd9, 0xab,
	0xe2, 0x59, 0x28, 0x43, 0x25, 0xf4, 0x11, 0xc0, 0x21, 0x0a, 0xbb, 0xff, 0x2c, 0x93, 0xed, 0xf7,
	0xbe, 0x9e, 0xe0, 0x52, 0x83, 0x3e, 0x28, 0xb8, 0x95, 0x21, 0x99, 0x10, 0xc6, 0x46, 0xbf, 0xae,
	0xd3, 0xe9, 0x31, 0xb7, 0x12, 0x73, 0xd9, 0x24, 0xf5, 0xca, 0x14, 0xda, 0xdd, 0x59, 0xda, 0x6b,
	0xf4, 0x49, 0x80, 0xb0, 0x81, 0xef, 0x91, 0x8d, 0xea, 0x8b, 0x0c, 0xeb, 0xaf, 0x87, 0x76, 0xd6,
	0xe7, 0x2a, 0xac, 0x41, 0x08, 0x24, 0x77, 0x85, 0x23, 0x18, 0xc8, 0x46, 0x7f, 0x6d, 0xc2, 0xdd,
	0xcb, 0xf0, 0xaa, 0xf9, 0xab, 0x46, 0x56, 0x0a, 0xee, 0x00, 0x2f, 0x83, 0xdf, 0x3f, 0xd4, 0x65,
	0x50, 0x79, 0x6b, 0xf5, 0x57, 0x16, 0xbf, 0xb9, 0x0e, 0xff, 0xfb, 0xe6, 0x6a, 0x92, 0x3a, 0x3c,
	0xb6, 0xd8, 0x88, 0x0b, 0x6f, 0x2c, 0xa6, 0xb3, 0xd1, 0x27, 0x00, 0x9d, 0x20, 0x02, 0x67, 0x59,
	0x98, 0x34, 0x55, 0xf0, 0x11, 0xe2, 0xfd, 0x14, 0x35, 0x8d, 0x19, 0x8a, 0x11, 0x7f, 0x44, 0xae,
	0x19, 0x9f, 0x32, 0x9e, 0x65, 0xb1, 0x16, 0x58, 0x82, 0xa1, 0x08, 0x9e, 0xa1, 0x38, 0x32, 0x3e,
	0x3d, 0x9a, 0x53, 0x58, 0x0f, 0x3d, 0xb2, 0xe5, 0x4c, 0x0a, 0x37, 0xd2, 0xbb, 0x73, 0xbe, 0xc4,
	0x39, 0xd7, 0x90, 0x7d, 0x77, 0xd6, 0x13, 0x42, 0x95, 0x9f, 0x28, 0x0b, 0x13, 0x85, 0x49, 0xb2,
	0x58, 0xf3, 0x54, 0xa8, 0x30, 0xef, 0x2b, 0x9c, 0xb7, 0x55, 0xf2, 0xc7, 0x33, 0x1a, 0x67, 0x36,
	0x49, 0x5d, 0x72, 0xaf, 0xca, 0x03, 0xfe, 0x75, 0x38, 0x7c, 0x00, 0x85, 0x13, 0x3e, 0x5c, 0xc1,
	0x7f, 0x19, 0xba, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xc6, 0x84, 0xed, 0x55, 0x0c, 0x00,
	0x00,
}
