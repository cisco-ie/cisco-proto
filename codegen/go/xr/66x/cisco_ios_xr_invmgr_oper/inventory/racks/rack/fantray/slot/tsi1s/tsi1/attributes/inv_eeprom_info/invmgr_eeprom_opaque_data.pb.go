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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_attributes_inv_eeprom_info

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
	proto.RegisterType((*InvmgrEepromOpaqueData_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.inv_eeprom_info.invmgr_eeprom_opaque_data_KEYS")
	proto.RegisterType((*RmaDetail)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.inv_eeprom_info.rma_detail")
	proto.RegisterType((*DiagEeprom)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.inv_eeprom_info.diag_eeprom")
	proto.RegisterType((*InvmgrEepromOpaqueData)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.inv_eeprom_info.invmgr_eeprom_opaque_data")
}

func init() { proto.RegisterFile("invmgr_eeprom_opaque_data.proto", fileDescriptor_6030716c6d5948f0) }

var fileDescriptor_6030716c6d5948f0 = []byte{
	// 1331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcf, 0x72, 0x14, 0xb7,
	0x13, 0x2e, 0xff, 0x6c, 0x6c, 0xac, 0x65, 0x8d, 0xad, 0x1f, 0x18, 0x39, 0x09, 0xd8, 0x18, 0x08,
	0x06, 0x27, 0x1b, 0xf6, 0x0f, 0x36, 0x90, 0xe4, 0x60, 0x0c, 0x54, 0xa8, 0x04, 0x57, 0x6a, 0x97,
	0x50, 0x95, 0x93, 0xa2, 0x95, 0xb4, 0xbb, 0x8a, 0x67, 0x46, 0x13, 0x49, 0xb3, 0x8e, 0x73, 0xc8,
	0x2d, 0xef, 0x91, 0x27, 0xcb, 0xb3, 0xa4, 0xba, 0x35, 0xe3, 0x9d, 0x22, 0xcb, 0x91, 0x8b, 0x3d,
	0xfb, 0x7d, 0x5f, 0x77, 0x4b, 0xdd, 0xad, 0x96, 0xc8, 0xb6, 0xc9, 0xa6, 0xe9, 0xd8, 0x71, 0xad,
	0x73, 0x67, 0x53, 0x6e, 0x73, 0xf1, 0x5b, 0xa1, 0xb9, 0x12, 0x41, 0xb4, 0x72, 0x67, 0x83, 0xa5,
	0xbf, 0x48, 0xe3, 0xa5, 0xe5, 0xc6, 0x7a, 0xfe, 0xbb, 0xe3, 0xa5, 0xda, 0xe6, 0xda, 0xb5, 0x4c,
	0x36, 0xd5, 0x59, 0xb0, 0xee, 0xbc, 0xe5, 0x84, 0x3c, 0xf5, 0xf8, 0xb7, 0x35, 0x12, 0x59, 0x70,
	0xe2, 0xbc, 0xe5, 0x13, 0x1b, 0x5a, 0xc1, 0x9b, 0xb6, 0xc7, 0xbf, 0x2d, 0x11, 0x82, 0x33, 0xc3,
	0x22, 0x68, 0x0f, 0x76, 0x55, 0x38, 0x93, 0x8d, 0xec, 0xee, 0x90, 0xdc, 0xfa, 0xe0, 0x22, 0xf8,
	0xf7, 0x2f, 0x7f, 0x1e, 0x50, 0x4a, 0x96, 0x32, 0x91, 0x6a, 0xb6, 0xb0, 0xb3, 0xb0, 0xb7, 0xda,
	0xc7, 0x6f, 0x7a, 0x9d, 0x2c, 0xc3, 0x7f, 0xde, 0x66, 0xff, 0x43, 0xf4, 0x12, 0xfc, 0x6a, 0x5f,
	0xc0, 0x1d, 0xb6, 0x38, 0x83, 0x3b, 0xbb, 0x96, 0x10, 0x97, 0x0a, 0xae, 0x74, 0x10, 0x26, 0xa1,
	0xb7, 0xc9, 0x95, 0xa0, 0x7d, 0xe0, 0x13, 0xe3, 0x61, 0xfd, 0xa5, 0xdf, 0x06, 0x60, 0xdf, 0x45,
	0x88, 0xde, 0x8c, 0x06, 0x59, 0x91, 0x0e, 0xb5, 0x2b, 0x43, 0xac, 0xba, 0x54, 0x9c, 0x20, 0x40,
	0xb7, 0x49, 0x03, 0xe8, 0xca, 0x41, 0x8c, 0x05, 0x16, 0xa5, 0xfd, 0xee, 0xdf, 0xeb, 0xa4, 0xa1,
	0x8c, 0x18, 0x97, 0x7b, 0xa2, 0x3b, 0xa4, 0xa1, 0xb4, 0x97, 0xce, 0xe4, 0xc1, 0xd8, 0xac, 0x8a,
	0x58, 0x83, 0xe8, 0x43, 0xb2, 0x61, 0x14, 0xee, 0x7f, 0x64, 0x5d, 0x2a, 0x02, 0x77, 0x7a, 0x5a,
	0x06, 0xbe, 0x1a, 0x89, 0x57, 0x88, 0xf7, 0xf5, 0x94, 0xee, 0x93, 0x0d, 0x69, 0xb3, 0xe0, 0x6c,
	0x92, 0x68, 0xc7, 0x47, 0x22, 0x35, 0x49, 0xb5, 0x88, 0xf5, 0x19, 0xf1, 0x0a, 0x71, 0x7a, 0x9f,
	0x5c, 0xad, 0x89, 0xc3, 0x79, 0xae, 0xd9, 0x12, 0x4a, 0xd7, 0x66, 0xf0, 0xdb, 0xf3, 0x5c, 0xd3,
	0x75, 0xb2, 0x38, 0x35, 0x8a, 0x5d, 0x42, 0x12, 0x3e, 0x21, 0xf1, 0x93, 0x33, 0xa3, 0xd8, 0x72,
	0x4c, 0x3c, 0x7c, 0x83, 0x2a, 0x37, 0x8a, 0xad, 0x44, 0x55, 0x6e, 0x14, 0x04, 0x28, 0x94, 0xe1,
	0xf5, 0xfd, 0x5d, 0x8e, 0x01, 0x0a, 0x65, 0x5e, 0xd4, 0xb6, 0xb8, 0x45, 0x2e, 0x83, 0x10, 0x6b,
	0xb9, 0x8a, 0x8a, 0x95, 0x42, 0x99, 0x13, 0x28, 0x27, 0x25, 0x4b, 0x32, 0xd1, 0x86, 0x91, 0x18,
	0x09, 0xbe, 0x21, 0x92, 0x96, 0x86, 0x35, 0x62, 0x24, 0x2d, 0x0d, 0xdd, 0x27, 0x34, 0xd8, 0x9c,
	0x0b, 0xef, 0x75, 0xca, 0x73, 0xe1, 0x02, 0x14, 0x88, 0x5d, 0x89, 0x49, 0x0a, 0x36, 0x3f, 0x02,
	0xe2, 0x47, 0xe1, 0xc2, 0x49, 0x91, 0xd2, 0x5d, 0xd2, 0x9c, 0x89, 0x61, 0x63, 0xcd, 0xb2, 0xcc,
	0xa5, 0xee, 0x9d, 0x51, 0xf4, 0x06, 0x59, 0xc9, 0x25, 0x96, 0x99, 0xad, 0x21, 0xbb, 0x9c, 0x4b,
	0xa8, 0x31, 0xdd, 0x24, 0xf0, 0x05, 0x56, 0x57, 0x2f, 0x70, 0xc8, 0xc8, 0x36, 0x69, 0xc8, 0x89,
	0xf0, 0xde, 0x78, 0xee, 0x8d, 0x62, 0xeb, 0xb1, 0xf0, 0x25, 0x34, 0x30, 0x0a, 0xf6, 0xa8, 0xf4,
	0x14, 0x3c, 0xb6, 0xd9, 0x46, 0xdc, 0xa3, 0xd2, 0xd3, 0x93, 0x22, 0x6d, 0xd7, 0xa8, 0x0e, 0xa3,
	0x75, 0xaa, 0x53, 0xa3, 0xba, 0xec, 0xff, 0x75, 0xaa, 0x5b, 0xa3, 0x7a, 0xec, 0x5a, 0x9d, 0xea,
	0xd5, 0xa8, 0xc7, 0xec, 0x7a, 0x9d, 0x7a, 0x5c, 0xa3, 0x0e, 0xd8, 0x66, 0x9d, 0x3a, 0xa8, 0x51,
	0x87, 0xec, 0x46, 0x9d, 0x3a, 0xa4, 0x77, 0xc9, 0x5a, 0x2a, 0xb2, 0x82, 0xe3, 0xe9, 0x80, 0xf3,
	0xc7, 0x18, 0x0a, 0xae, 0x00, 0xfa, 0x56, 0xfb, 0xf0, 0x42, 0x04, 0x41, 0xff, 0x24, 0x8b, 0x2e,
	0x15, 0x6c, 0x6b, 0x67, 0x61, 0xaf, 0xd1, 0x49, 0x5a, 0x1f, 0x7b, 0x40, 0xb4, 0x66, 0x27, 0xb7,
	0x0f, 0x81, 0x61, 0x03, 0x50, 0xd4, 0xc0, 0x8d, 0x62, 0x9f, 0xc4, 0x0d, 0xe0, 0xef, 0xd7, 0x58,
	0x9e, 0x48, 0x89, 0xc4, 0x08, 0xcf, 0x3e, 0x8d, 0xe5, 0x41, 0xe8, 0x08, 0x10, 0x38, 0x65, 0x43,
	0xe1, 0x35, 0x4f, 0x85, 0xe4, 0x42, 0x29, 0xa7, 0xbd, 0x6f, 0xb3, 0xcf, 0x62, 0x03, 0x01, 0xf1,
	0x46, 0xc8, 0xa3, 0x12, 0xa6, 0x0f, 0xc8, 0x46, 0x29, 0xe3, 0xc3, 0xe4, 0x94, 0x7b, 0xf3, 0x87,
	0x6e, 0xb3, 0x9b, 0xb1, 0xb3, 0x53, 0xd4, 0x3d, 0x4f, 0x4e, 0x07, 0x80, 0xce, 0x73, 0xdb, 0x61,
	0xb7, 0xe6, 0xb9, 0xed, 0xcc, 0x73, 0xdb, 0x61, 0xdb, 0x73, 0xdc, 0x76, 0xe6, 0xb9, 0xed, 0xb2,
	0x9d, 0x79, 0x6e, 0xbb, 0xf3, 0xdc, 0x76, 0xd9, 0xed, 0x39, 0x6e, 0xbb, 0xf3, 0xdc, 0xf6, 0xd8,
	0xee, 0x3c, 0xb7, 0xbd, 0x79, 0x6e, 0x7b, 0xec, 0xce, 0x1c, 0xb7, 0x3d, 0xe8, 0x9e, 0x5c, 0x0e,
	0xb9, 0xd7, 0xce, 0x88, 0x04, 0xcf, 0xd4, 0xdd, 0xd8, 0x3d, 0xb9, 0x1c, 0x0e, 0x10, 0x84, 0x93,
	0xf5, 0x90, 0x6c, 0xe4, 0xf6, 0x4c, 0x3b, 0xee, 0x8b, 0x3c, 0x4f, 0xce, 0xe3, 0x40, 0xba, 0x17,
	0x83, 0x23, 0x31, 0x40, 0x1c, 0x27, 0xd2, 0x7e, 0xa5, 0x95, 0x36, 0xf3, 0x45, 0x1a, 0x67, 0xcb,
	0xe7, 0x71, 0xce, 0x21, 0x71, 0x3c, 0xc3, 0x61, 0x0c, 0x0d, 0x13, 0x2b, 0x61, 0x8d, 0xe3, 0x4c,
	0x84, 0xc2, 0x69, 0x76, 0x3f, 0xae, 0x13, 0xe1, 0x41, 0x85, 0xd2, 0x3b, 0xa4, 0x19, 0x85, 0x53,
	0xed, 0x3c, 0x78, 0xdc, 0x8b, 0xcb, 0x44, 0xf0, 0x5d, 0xc4, 0xe0, 0x8e, 0x88, 0xa2, 0x44, 0x67,
	0xe3, 0x30, 0x61, 0x0f, 0xe2, 0xf0, 0x40, 0xec, 0x07, 0x84, 0xe8, 0x3d, 0x12, 0x3d, 0x73, 0x39,
	0xd1, 0xf2, 0xd4, 0x17, 0x29, 0x7b, 0x88, 0xa2, 0xe8, 0xfd, 0xb8, 0x04, 0xa1, 0x27, 0xcb, 0x6e,
	0x86, 0xe4, 0xb1, 0xfd, 0xd8, 0x93, 0x11, 0x82, 0xc4, 0x81, 0xa0, 0xf4, 0x63, 0x8b, 0x2c, 0xb0,
	0x2f, 0xa2, 0x20, 0x3a, 0x01, 0x04, 0x12, 0x3b, 0x72, 0x05, 0x4f, 0xc5, 0xaf, 0xb6, 0x1c, 0xe0,
	0x5f, 0xc6, 0x15, 0x8f, 0x5c, 0xf1, 0x06, 0x40, 0x4c, 0x56, 0xa5, 0x32, 0x59, 0xa5, 0x6a, 0xcd,
	0x54, 0x00, 0xa2, 0xea, 0x26, 0x21, 0x56, 0xa7, 0xdc, 0x07, 0x67, 0xb2, 0x31, 0xfb, 0x2a, 0x5e,
	0x6c, 0x56, 0xa7, 0x03, 0x04, 0x80, 0xce, 0x9d, 0x55, 0x85, 0xc4, 0xd3, 0xf5, 0x28, 0xd2, 0x25,
	0xf2, 0x5a, 0x41, 0xea, 0x66, 0xe5, 0x85, 0x9b, 0xb1, 0x1d, 0x43, 0xf8, 0xaa, 0xbc, 0xe5, 0xe5,
	0x58, 0xcd, 0x66, 0x90, 0x74, 0xe2, 0x7e, 0xf2, 0x38, 0x96, 0x41, 0x70, 0x87, 0x34, 0x51, 0xe0,
	0xf4, 0xd4, 0x60, 0x01, 0xba, 0x65, 0x9f, 0x08, 0x07, 0xd7, 0x1b, 0x62, 0x20, 0x4a, 0x47, 0x63,
	0xae, 0xf4, 0xd4, 0x08, 0xac, 0x7b, 0xaf, 0x1c, 0x45, 0xa3, 0xf1, 0x8b, 0x0a, 0x83, 0xe5, 0x4e,
	0xce, 0x2e, 0xea, 0xf8, 0x38, 0x2e, 0x77, 0x72, 0x56, 0x15, 0x71, 0x8b, 0x5c, 0x06, 0x1f, 0x43,
	0x13, 0x3c, 0x3b, 0x88, 0x93, 0x22, 0x1d, 0x8d, 0x9f, 0x9b, 0xe0, 0xa1, 0xbe, 0x3a, 0x1b, 0x9b,
	0x4c, 0x6b, 0xc7, 0x0b, 0xaf, 0xd9, 0x61, 0xac, 0x6f, 0x85, 0xfd, 0xe4, 0x35, 0x65, 0x64, 0xc5,
	0x67, 0x69, 0x6e, 0x8d, 0x62, 0x4f, 0xa2, 0x71, 0xf9, 0x13, 0xfc, 0xc2, 0x50, 0x92, 0x56, 0x69,
	0xf6, 0x34, 0x52, 0x2e, 0x15, 0xc7, 0x56, 0x69, 0xba, 0x47, 0xd6, 0xb5, 0x34, 0x5c, 0x24, 0xf9,
	0xe4, 0xe2, 0xf9, 0xf0, 0x2c, 0xb6, 0xa1, 0x96, 0xe6, 0x08, 0xe0, 0x98, 0x85, 0xdd, 0x7f, 0x96,
	0xc8, 0xd6, 0x07, 0x1f, 0x3e, 0x70, 0x7b, 0xc1, 0x1c, 0x94, 0xc2, 0xa9, 0x58, 0x4c, 0x48, 0x63,
	0xb3, 0xdf, 0x30, 0xd9, 0xf4, 0x58, 0x38, 0x85, 0xb5, 0xdc, 0x26, 0x8d, 0x9a, 0x09, 0xeb, 0xee,
	0x2c, 0xee, 0x35, 0xfb, 0x24, 0x42, 0x38, 0xa9, 0xf7, 0xc8, 0x7a, 0xfd, 0x31, 0x85, 0xfd, 0xd7,
	0x43, 0x3f, 0x6b, 0x33, 0x15, 0xf6, 0x20, 0x24, 0x52, 0xf8, 0x72, 0x21, 0x98, 0xc8, 0x66, 0x7f,
	0x75, 0x22, 0xfc, 0xcb, 0xf8, 0x7c, 0xf9, 0x6b, 0x81, 0x2c, 0x97, 0xdc, 0x01, 0x8e, 0xfd, 0xf4,
	0xe3, 0x8f, 0xfd, 0xda, 0xf3, 0xa9, 0xbf, 0x3c, 0xff, 0x19, 0x75, 0xf8, 0xdf, 0x67, 0xd4, 0x36,
	0x69, 0xc0, 0xfb, 0x89, 0x8f, 0x84, 0x0c, 0xd6, 0x61, 0xe1, 0x9a, 0x7d, 0x02, 0xd0, 0x2b, 0x44,
	0xe0, 0xd4, 0x4a, 0x9b, 0x65, 0x1a, 0x7e, 0xc4, 0xcc, 0x3e, 0x45, 0x4d, 0xf3, 0x02, 0xc5, 0xdc,
	0x3e, 0x22, 0xd7, 0x6c, 0xc8, 0xb8, 0xc8, 0xf3, 0xc4, 0x48, 0x6c, 0xb6, 0x58, 0xee, 0x67, 0x28,
	0xa6, 0x36, 0x64, 0x47, 0x33, 0x0a, 0x2b, 0xdf, 0x23, 0x9b, 0xde, 0x66, 0x70, 0xf7, 0xbc, 0x6f,
	0xf3, 0x35, 0xda, 0x5c, 0x43, 0xf6, 0x7d, 0xab, 0x27, 0x84, 0xe9, 0x30, 0xd1, 0x0e, 0x0c, 0xa5,
	0x4d, 0xf3, 0xc4, 0x88, 0x4c, 0xea, 0x68, 0xf7, 0x0d, 0xda, 0x6d, 0x56, 0xfc, 0xf1, 0x05, 0x8d,
	0x96, 0xdb, 0xa4, 0xa1, 0x44, 0xd0, 0xd5, 0x51, 0xfe, 0x36, 0x1e, 0x33, 0x80, 0xe2, 0x59, 0x1e,
	0x2e, 0xe3, 0x0b, 0xbe, 0xfb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xd1, 0xd0, 0xa9, 0xe4,
	0x0b, 0x00, 0x00,
}
