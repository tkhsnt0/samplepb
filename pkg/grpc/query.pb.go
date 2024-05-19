// ライセンスヘッダ:バージョン3を利用

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.9
// source: api/protobuf/query.proto

// パッケージの宣言

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *****************************************
// 商品カテゴリ検索用param型とResult型の定義
// *****************************************
//
//	カテゴリ検索用Param型
type CategoryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // カテゴリ番号
}

func (x *CategoryParam) Reset() {
	*x = CategoryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryParam) ProtoMessage() {}

func (x *CategoryParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryParam.ProtoReflect.Descriptor instead.
func (*CategoryParam) Descriptor() ([]byte, []int) {
	return file_api_protobuf_query_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryParam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// 商品カテゴリ複数件を返すResult型
type CategoriesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*Category            `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"` // 商品カテゴリ複数
	Error      *Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`           // エラー
	Timestamp  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`   // タイムスタンプ
}

func (x *CategoriesResult) Reset() {
	*x = CategoriesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoriesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesResult) ProtoMessage() {}

func (x *CategoriesResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesResult.ProtoReflect.Descriptor instead.
func (*CategoriesResult) Descriptor() ([]byte, []int) {
	return file_api_protobuf_query_proto_rawDescGZIP(), []int{1}
}

func (x *CategoriesResult) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *CategoriesResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CategoriesResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// 商品カテゴリ1件を返すResult型
type CategoryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// エラーか検索結果のいずれかを返す
	//
	// Types that are assignable to Result:
	//
	//	*CategoryResult_Category
	//	*CategoryResult_Error
	Result    isCategoryResult_Result `protobuf_oneof:"result"`
	Timestamp *timestamppb.Timestamp  `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // タイムスタンプ
}

func (x *CategoryResult) Reset() {
	*x = CategoryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResult) ProtoMessage() {}

func (x *CategoryResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResult.ProtoReflect.Descriptor instead.
func (*CategoryResult) Descriptor() ([]byte, []int) {
	return file_api_protobuf_query_proto_rawDescGZIP(), []int{2}
}

func (m *CategoryResult) GetResult() isCategoryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *CategoryResult) GetCategory() *Category {
	if x, ok := x.GetResult().(*CategoryResult_Category); ok {
		return x.Category
	}
	return nil
}

func (x *CategoryResult) GetError() *Error {
	if x, ok := x.GetResult().(*CategoryResult_Error); ok {
		return x.Error
	}
	return nil
}

func (x *CategoryResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type isCategoryResult_Result interface {
	isCategoryResult_Result()
}

type CategoryResult_Category struct {
	Category *Category `protobuf:"bytes,1,opt,name=category,proto3,oneof"` // 商品カテゴリ
}

type CategoryResult_Error struct {
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"` // エラー
}

func (*CategoryResult_Category) isCategoryResult_Result() {}

func (*CategoryResult_Error) isCategoryResult_Result() {}

// *****************************************
// 商品検索用param型とResult型の定義
// *****************************************
//
//	商品検索用Param型
type ProductParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`           // 商品番号
	Keyword string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"` // キーワード
}

func (x *ProductParam) Reset() {
	*x = ProductParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductParam) ProtoMessage() {}

func (x *ProductParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductParam.ProtoReflect.Descriptor instead.
func (*ProductParam) Descriptor() ([]byte, []int) {
	return file_api_protobuf_query_proto_rawDescGZIP(), []int{3}
}

func (x *ProductParam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductParam) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// 商品複数件を返すResult型
type ProductsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products  []*Product             `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`   // 商品複数
	Error     *Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`         // エラー
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // タイムスタンプ
}

func (x *ProductsResult) Reset() {
	*x = ProductsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsResult) ProtoMessage() {}

func (x *ProductsResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsResult.ProtoReflect.Descriptor instead.
func (*ProductsResult) Descriptor() ([]byte, []int) {
	return file_api_protobuf_query_proto_rawDescGZIP(), []int{4}
}

func (x *ProductsResult) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductsResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ProductsResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// 商品1件を返すResult型
type ProductResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// エラーか検索結果のいずれかを返す
	//
	// Types that are assignable to Result:
	//
	//	*ProductResult_Product
	//	*ProductResult_Error
	Result    isProductResult_Result `protobuf_oneof:"result"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // タイムスタンプ
}

func (x *ProductResult) Reset() {
	*x = ProductResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResult) ProtoMessage() {}

func (x *ProductResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResult.ProtoReflect.Descriptor instead.
func (*ProductResult) Descriptor() ([]byte, []int) {
	return file_api_protobuf_query_proto_rawDescGZIP(), []int{5}
}

func (m *ProductResult) GetResult() isProductResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ProductResult) GetProduct() *Product {
	if x, ok := x.GetResult().(*ProductResult_Product); ok {
		return x.Product
	}
	return nil
}

func (x *ProductResult) GetError() *Error {
	if x, ok := x.GetResult().(*ProductResult_Error); ok {
		return x.Error
	}
	return nil
}

func (x *ProductResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type isProductResult_Result interface {
	isProductResult_Result()
}

type ProductResult_Product struct {
	Product *Product `protobuf:"bytes,1,opt,name=product,proto3,oneof"` // 検索結果
}

type ProductResult_Error struct {
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"` // 検索エラー
}

func (*ProductResult_Product) isProductResult_Result() {}

func (*ProductResult_Error) isProductResult_Result() {}

var File_api_protobuf_query_proto protoreflect.FileDescriptor

var file_api_protobuf_query_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x10, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0xaf, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x48, 0x00, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x38, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xa0,
	0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0x86, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x3a, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x39, 0x0a,
	0x04, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xfb, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37,
	0x0a, 0x04, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x42, 0x79, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protobuf_query_proto_rawDescOnce sync.Once
	file_api_protobuf_query_proto_rawDescData = file_api_protobuf_query_proto_rawDesc
)

func file_api_protobuf_query_proto_rawDescGZIP() []byte {
	file_api_protobuf_query_proto_rawDescOnce.Do(func() {
		file_api_protobuf_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protobuf_query_proto_rawDescData)
	})
	return file_api_protobuf_query_proto_rawDescData
}

var file_api_protobuf_query_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_protobuf_query_proto_goTypes = []interface{}{
	(*CategoryParam)(nil),         // 0: protobuf.CategoryParam
	(*CategoriesResult)(nil),      // 1: protobuf.CategoriesResult
	(*CategoryResult)(nil),        // 2: protobuf.CategoryResult
	(*ProductParam)(nil),          // 3: protobuf.ProductParam
	(*ProductsResult)(nil),        // 4: protobuf.ProductsResult
	(*ProductResult)(nil),         // 5: protobuf.ProductResult
	(*Category)(nil),              // 6: protobuf.Category
	(*Error)(nil),                 // 7: protobuf.Error
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*Product)(nil),               // 9: protobuf.Product
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_api_protobuf_query_proto_depIdxs = []int32{
	6,  // 0: protobuf.CategoriesResult.categories:type_name -> protobuf.Category
	7,  // 1: protobuf.CategoriesResult.error:type_name -> protobuf.Error
	8,  // 2: protobuf.CategoriesResult.timestamp:type_name -> google.protobuf.Timestamp
	6,  // 3: protobuf.CategoryResult.category:type_name -> protobuf.Category
	7,  // 4: protobuf.CategoryResult.error:type_name -> protobuf.Error
	8,  // 5: protobuf.CategoryResult.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 6: protobuf.ProductsResult.products:type_name -> protobuf.Product
	7,  // 7: protobuf.ProductsResult.error:type_name -> protobuf.Error
	8,  // 8: protobuf.ProductsResult.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 9: protobuf.ProductResult.product:type_name -> protobuf.Product
	7,  // 10: protobuf.ProductResult.error:type_name -> protobuf.Error
	8,  // 11: protobuf.ProductResult.timestamp:type_name -> google.protobuf.Timestamp
	10, // 12: protobuf.CategoryQuery.List:input_type -> google.protobuf.Empty
	0,  // 13: protobuf.CategoryQuery.ById:input_type -> protobuf.CategoryParam
	10, // 14: protobuf.ProductQuery.ListStream:input_type -> google.protobuf.Empty
	10, // 15: protobuf.ProductQuery.List:input_type -> google.protobuf.Empty
	3,  // 16: protobuf.ProductQuery.ById:input_type -> protobuf.ProductParam
	3,  // 17: protobuf.ProductQuery.ByKeyword:input_type -> protobuf.ProductParam
	1,  // 18: protobuf.CategoryQuery.List:output_type -> protobuf.CategoriesResult
	2,  // 19: protobuf.CategoryQuery.ById:output_type -> protobuf.CategoryResult
	9,  // 20: protobuf.ProductQuery.ListStream:output_type -> protobuf.Product
	4,  // 21: protobuf.ProductQuery.List:output_type -> protobuf.ProductsResult
	5,  // 22: protobuf.ProductQuery.ById:output_type -> protobuf.ProductResult
	4,  // 23: protobuf.ProductQuery.ByKeyword:output_type -> protobuf.ProductsResult
	18, // [18:24] is the sub-list for method output_type
	12, // [12:18] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_api_protobuf_query_proto_init() }
func file_api_protobuf_query_proto_init() {
	if File_api_protobuf_query_proto != nil {
		return
	}
	file_api_protobuf_error_proto_init()
	file_api_protobuf_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_protobuf_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryParam); i {
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
		file_api_protobuf_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoriesResult); i {
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
		file_api_protobuf_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryResult); i {
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
		file_api_protobuf_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductParam); i {
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
		file_api_protobuf_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsResult); i {
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
		file_api_protobuf_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResult); i {
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
	file_api_protobuf_query_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CategoryResult_Category)(nil),
		(*CategoryResult_Error)(nil),
	}
	file_api_protobuf_query_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ProductResult_Product)(nil),
		(*ProductResult_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_protobuf_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_protobuf_query_proto_goTypes,
		DependencyIndexes: file_api_protobuf_query_proto_depIdxs,
		MessageInfos:      file_api_protobuf_query_proto_msgTypes,
	}.Build()
	File_api_protobuf_query_proto = out.File
	file_api_protobuf_query_proto_rawDesc = nil
	file_api_protobuf_query_proto_goTypes = nil
	file_api_protobuf_query_proto_depIdxs = nil
}
