// Copyright (c) Aptos
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: aptos/tokens/v1/tokens.proto

package pbaptos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHeight     uint64            `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	ChainId         uint32            `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Tokens          []*Token          `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens,omitempty"`
	TokenDatas      []*TokenData      `protobuf:"bytes,4,rep,name=token_datas,json=tokenDatas,proto3" json:"token_datas,omitempty"`
	CollectionDatas []*CollectionData `protobuf:"bytes,5,rep,name=collection_datas,json=collectionDatas,proto3" json:"collection_datas,omitempty"`
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

func (x *Tokens) ProtoReflect() protoreflect.Message {
	mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_aptos_tokens_v1_tokens_proto_rawDescGZIP(), []int{0}
}

func (x *Tokens) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *Tokens) GetChainId() uint32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Tokens) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *Tokens) GetTokenDatas() []*TokenData {
	if x != nil {
		return x.TokenDatas
	}
	return nil
}

func (x *Tokens) GetCollectionDatas() []*CollectionData {
	if x != nil {
		return x.CollectionDatas
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenId            *TokenId `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	TransactionVersion uint64   `protobuf:"varint,2,opt,name=transaction_version,json=transactionVersion,proto3" json:"transaction_version,omitempty"`
	TokenProperties    string   `protobuf:"bytes,3,opt,name=token_properties,json=tokenProperties,proto3" json:"token_properties,omitempty"`
	Amount             uint64   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	OwnerAddress       *string  `protobuf:"bytes,5,opt,name=owner_address,json=ownerAddress,proto3,oneof" json:"owner_address,omitempty"`
	TableHandle        string   `protobuf:"bytes,6,opt,name=table_handle,json=tableHandle,proto3" json:"table_handle,omitempty"`
	TableType          *string  `protobuf:"bytes,7,opt,name=table_type,json=tableType,proto3,oneof" json:"table_type,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_aptos_tokens_v1_tokens_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetTokenId() *TokenId {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *Token) GetTransactionVersion() uint64 {
	if x != nil {
		return x.TransactionVersion
	}
	return 0
}

func (x *Token) GetTokenProperties() string {
	if x != nil {
		return x.TokenProperties
	}
	return ""
}

func (x *Token) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Token) GetOwnerAddress() string {
	if x != nil && x.OwnerAddress != nil {
		return *x.OwnerAddress
	}
	return ""
}

func (x *Token) GetTableHandle() string {
	if x != nil {
		return x.TableHandle
	}
	return ""
}

func (x *Token) GetTableType() string {
	if x != nil && x.TableType != nil {
		return *x.TableType
	}
	return ""
}

type TokenData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenDataId              *TokenDataId `protobuf:"bytes,1,opt,name=token_data_id,json=tokenDataId,proto3" json:"token_data_id,omitempty"`
	TransactionVersion       uint64       `protobuf:"varint,2,opt,name=transaction_version,json=transactionVersion,proto3" json:"transaction_version,omitempty"`
	Maximum                  uint64       `protobuf:"varint,3,opt,name=maximum,proto3" json:"maximum,omitempty"`
	Supply                   uint64       `protobuf:"varint,4,opt,name=supply,proto3" json:"supply,omitempty"`
	LargestPropertyVersion   uint64       `protobuf:"varint,5,opt,name=largest_property_version,json=largestPropertyVersion,proto3" json:"largest_property_version,omitempty"`
	MetadataUri              string       `protobuf:"bytes,6,opt,name=metadata_uri,json=metadataUri,proto3" json:"metadata_uri,omitempty"`
	PayeeAddress             string       `protobuf:"bytes,7,opt,name=payee_address,json=payeeAddress,proto3" json:"payee_address,omitempty"`
	RoyaltyPointsNumerator   uint64       `protobuf:"varint,8,opt,name=royalty_points_numerator,json=royaltyPointsNumerator,proto3" json:"royalty_points_numerator,omitempty"`
	RoyaltyPointsDenominator uint64       `protobuf:"varint,9,opt,name=royalty_points_denominator,json=royaltyPointsDenominator,proto3" json:"royalty_points_denominator,omitempty"`
	MaximumMutable           bool         `protobuf:"varint,10,opt,name=maximum_mutable,json=maximumMutable,proto3" json:"maximum_mutable,omitempty"`
	UriMutable               bool         `protobuf:"varint,11,opt,name=uri_mutable,json=uriMutable,proto3" json:"uri_mutable,omitempty"`
	DescriptionMutable       bool         `protobuf:"varint,12,opt,name=description_mutable,json=descriptionMutable,proto3" json:"description_mutable,omitempty"`
	PropertiesMutable        bool         `protobuf:"varint,13,opt,name=properties_mutable,json=propertiesMutable,proto3" json:"properties_mutable,omitempty"`
	RoyaltyMutable           bool         `protobuf:"varint,14,opt,name=royalty_mutable,json=royaltyMutable,proto3" json:"royalty_mutable,omitempty"`
	DefaultProperties        string       `protobuf:"bytes,15,opt,name=default_properties,json=defaultProperties,proto3" json:"default_properties,omitempty"`
}

func (x *TokenData) Reset() {
	*x = TokenData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenData) ProtoMessage() {}

func (x *TokenData) ProtoReflect() protoreflect.Message {
	mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenData.ProtoReflect.Descriptor instead.
func (*TokenData) Descriptor() ([]byte, []int) {
	return file_aptos_tokens_v1_tokens_proto_rawDescGZIP(), []int{2}
}

func (x *TokenData) GetTokenDataId() *TokenDataId {
	if x != nil {
		return x.TokenDataId
	}
	return nil
}

func (x *TokenData) GetTransactionVersion() uint64 {
	if x != nil {
		return x.TransactionVersion
	}
	return 0
}

func (x *TokenData) GetMaximum() uint64 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

func (x *TokenData) GetSupply() uint64 {
	if x != nil {
		return x.Supply
	}
	return 0
}

func (x *TokenData) GetLargestPropertyVersion() uint64 {
	if x != nil {
		return x.LargestPropertyVersion
	}
	return 0
}

func (x *TokenData) GetMetadataUri() string {
	if x != nil {
		return x.MetadataUri
	}
	return ""
}

func (x *TokenData) GetPayeeAddress() string {
	if x != nil {
		return x.PayeeAddress
	}
	return ""
}

func (x *TokenData) GetRoyaltyPointsNumerator() uint64 {
	if x != nil {
		return x.RoyaltyPointsNumerator
	}
	return 0
}

func (x *TokenData) GetRoyaltyPointsDenominator() uint64 {
	if x != nil {
		return x.RoyaltyPointsDenominator
	}
	return 0
}

func (x *TokenData) GetMaximumMutable() bool {
	if x != nil {
		return x.MaximumMutable
	}
	return false
}

func (x *TokenData) GetUriMutable() bool {
	if x != nil {
		return x.UriMutable
	}
	return false
}

func (x *TokenData) GetDescriptionMutable() bool {
	if x != nil {
		return x.DescriptionMutable
	}
	return false
}

func (x *TokenData) GetPropertiesMutable() bool {
	if x != nil {
		return x.PropertiesMutable
	}
	return false
}

func (x *TokenData) GetRoyaltyMutable() bool {
	if x != nil {
		return x.RoyaltyMutable
	}
	return false
}

func (x *TokenData) GetDefaultProperties() string {
	if x != nil {
		return x.DefaultProperties
	}
	return ""
}

type TokenId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenDataId     *TokenDataId `protobuf:"bytes,1,opt,name=token_data_id,json=tokenDataId,proto3" json:"token_data_id,omitempty"`
	PropertyVersion uint64       `protobuf:"varint,2,opt,name=property_version,json=propertyVersion,proto3" json:"property_version,omitempty"`
}

func (x *TokenId) Reset() {
	*x = TokenId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenId) ProtoMessage() {}

func (x *TokenId) ProtoReflect() protoreflect.Message {
	mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenId.ProtoReflect.Descriptor instead.
func (*TokenId) Descriptor() ([]byte, []int) {
	return file_aptos_tokens_v1_tokens_proto_rawDescGZIP(), []int{3}
}

func (x *TokenId) GetTokenDataId() *TokenDataId {
	if x != nil {
		return x.TokenDataId
	}
	return nil
}

func (x *TokenId) GetPropertyVersion() uint64 {
	if x != nil {
		return x.PropertyVersion
	}
	return 0
}

type TokenDataId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatorAddress string `protobuf:"bytes,1,opt,name=creator_address,json=creatorAddress,proto3" json:"creator_address,omitempty"`
	CollectionName string `protobuf:"bytes,2,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TokenDataId) Reset() {
	*x = TokenDataId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenDataId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenDataId) ProtoMessage() {}

func (x *TokenDataId) ProtoReflect() protoreflect.Message {
	mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenDataId.ProtoReflect.Descriptor instead.
func (*TokenDataId) Descriptor() ([]byte, []int) {
	return file_aptos_tokens_v1_tokens_proto_rawDescGZIP(), []int{4}
}

func (x *TokenDataId) GetCreatorAddress() string {
	if x != nil {
		return x.CreatorAddress
	}
	return ""
}

func (x *TokenDataId) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

func (x *TokenDataId) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CollectionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatorAddress     string `protobuf:"bytes,1,opt,name=creator_address,json=creatorAddress,proto3" json:"creator_address,omitempty"`
	CollectionName     string `protobuf:"bytes,2,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	Description        string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	TransactionVersion uint64 `protobuf:"varint,4,opt,name=transaction_version,json=transactionVersion,proto3" json:"transaction_version,omitempty"`
	MetadataUri        string `protobuf:"bytes,5,opt,name=metadata_uri,json=metadataUri,proto3" json:"metadata_uri,omitempty"`
	Supply             uint64 `protobuf:"varint,6,opt,name=supply,proto3" json:"supply,omitempty"`
	Maximum            uint64 `protobuf:"varint,7,opt,name=maximum,proto3" json:"maximum,omitempty"`
	MaximumMutable     bool   `protobuf:"varint,8,opt,name=maximum_mutable,json=maximumMutable,proto3" json:"maximum_mutable,omitempty"`
	UriMutable         bool   `protobuf:"varint,9,opt,name=uri_mutable,json=uriMutable,proto3" json:"uri_mutable,omitempty"`
	DescriptionMutable bool   `protobuf:"varint,10,opt,name=description_mutable,json=descriptionMutable,proto3" json:"description_mutable,omitempty"`
}

func (x *CollectionData) Reset() {
	*x = CollectionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionData) ProtoMessage() {}

func (x *CollectionData) ProtoReflect() protoreflect.Message {
	mi := &file_aptos_tokens_v1_tokens_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionData.ProtoReflect.Descriptor instead.
func (*CollectionData) Descriptor() ([]byte, []int) {
	return file_aptos_tokens_v1_tokens_proto_rawDescGZIP(), []int{5}
}

func (x *CollectionData) GetCreatorAddress() string {
	if x != nil {
		return x.CreatorAddress
	}
	return ""
}

func (x *CollectionData) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

func (x *CollectionData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CollectionData) GetTransactionVersion() uint64 {
	if x != nil {
		return x.TransactionVersion
	}
	return 0
}

func (x *CollectionData) GetMetadataUri() string {
	if x != nil {
		return x.MetadataUri
	}
	return ""
}

func (x *CollectionData) GetSupply() uint64 {
	if x != nil {
		return x.Supply
	}
	return 0
}

func (x *CollectionData) GetMaximum() uint64 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

func (x *CollectionData) GetMaximumMutable() bool {
	if x != nil {
		return x.MaximumMutable
	}
	return false
}

func (x *CollectionData) GetUriMutable() bool {
	if x != nil {
		return x.UriMutable
	}
	return false
}

func (x *CollectionData) GetDescriptionMutable() bool {
	if x != nil {
		return x.DescriptionMutable
	}
	return false
}

var File_aptos_tokens_v1_tokens_proto protoreflect.FileDescriptor

var file_aptos_tokens_v1_tokens_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x61, 0x70, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22,
	0xff, 0x01, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x74, 0x6f, 0x73,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x12, 0x4a, 0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x61, 0x70, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x22, 0xc2, 0x02, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x33, 0x0a, 0x08, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x61, 0x70, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64,
	0x12, 0x2f, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x12, 0x22, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xac, 0x05, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70,
	0x74, 0x6f, 0x73, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x64, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x18, 0x6c, 0x61, 0x72,
	0x67, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x6c, 0x61, 0x72,
	0x67, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x55, 0x72, 0x69, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x61, 0x79, 0x65, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x72,
	0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x72,
	0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x1a, 0x72, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x18, 0x72, 0x6f, 0x79, 0x61, 0x6c,
	0x74, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x6d,
	0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x75, 0x72, 0x69, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x75, 0x72, 0x69, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2f, 0x0a,
	0x13, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x75, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2d,
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x6d, 0x75, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x72, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x4d,
	0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x76, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64,
	0x12, 0x40, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x74, 0x6f, 0x73, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x49, 0x64, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a,
	0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x85, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x72, 0x69, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12,
	0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x72, 0x69, 0x5f,
	0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75,
	0x72, 0x69, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2d,
	0x61, 0x70, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x61,
	0x70, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x62, 0x61, 0x70, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_aptos_tokens_v1_tokens_proto_rawDescOnce sync.Once
	file_aptos_tokens_v1_tokens_proto_rawDescData = file_aptos_tokens_v1_tokens_proto_rawDesc
)

func file_aptos_tokens_v1_tokens_proto_rawDescGZIP() []byte {
	file_aptos_tokens_v1_tokens_proto_rawDescOnce.Do(func() {
		file_aptos_tokens_v1_tokens_proto_rawDescData = protoimpl.X.CompressGZIP(file_aptos_tokens_v1_tokens_proto_rawDescData)
	})
	return file_aptos_tokens_v1_tokens_proto_rawDescData
}

var file_aptos_tokens_v1_tokens_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aptos_tokens_v1_tokens_proto_goTypes = []interface{}{
	(*Tokens)(nil),         // 0: aptos.tokens.v1.Tokens
	(*Token)(nil),          // 1: aptos.tokens.v1.Token
	(*TokenData)(nil),      // 2: aptos.tokens.v1.TokenData
	(*TokenId)(nil),        // 3: aptos.tokens.v1.TokenId
	(*TokenDataId)(nil),    // 4: aptos.tokens.v1.TokenDataId
	(*CollectionData)(nil), // 5: aptos.tokens.v1.CollectionData
}
var file_aptos_tokens_v1_tokens_proto_depIdxs = []int32{
	1, // 0: aptos.tokens.v1.Tokens.tokens:type_name -> aptos.tokens.v1.Token
	2, // 1: aptos.tokens.v1.Tokens.token_datas:type_name -> aptos.tokens.v1.TokenData
	5, // 2: aptos.tokens.v1.Tokens.collection_datas:type_name -> aptos.tokens.v1.CollectionData
	3, // 3: aptos.tokens.v1.Token.token_id:type_name -> aptos.tokens.v1.TokenId
	4, // 4: aptos.tokens.v1.TokenData.token_data_id:type_name -> aptos.tokens.v1.TokenDataId
	4, // 5: aptos.tokens.v1.TokenId.token_data_id:type_name -> aptos.tokens.v1.TokenDataId
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_aptos_tokens_v1_tokens_proto_init() }
func file_aptos_tokens_v1_tokens_proto_init() {
	if File_aptos_tokens_v1_tokens_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aptos_tokens_v1_tokens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tokens); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aptos_tokens_v1_tokens_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aptos_tokens_v1_tokens_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aptos_tokens_v1_tokens_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aptos_tokens_v1_tokens_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenDataId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aptos_tokens_v1_tokens_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_aptos_tokens_v1_tokens_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aptos_tokens_v1_tokens_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aptos_tokens_v1_tokens_proto_goTypes,
		DependencyIndexes: file_aptos_tokens_v1_tokens_proto_depIdxs,
		MessageInfos:      file_aptos_tokens_v1_tokens_proto_msgTypes,
	}.Build()
	File_aptos_tokens_v1_tokens_proto = out.File
	file_aptos_tokens_v1_tokens_proto_rawDesc = nil
	file_aptos_tokens_v1_tokens_proto_goTypes = nil
	file_aptos_tokens_v1_tokens_proto_depIdxs = nil
}
