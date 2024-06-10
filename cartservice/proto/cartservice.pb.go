// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.2
// source: proto/cartservice.proto

package microshopping

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{0}
}

// 购物车商品
type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 商品id
	ProductSku string `protobuf:"bytes,1,opt,name=product_sku,json=productSku,proto3" json:"product_sku,omitempty"`
	// 商品数量
	Quantity int32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{1}
}

func (x *CartItem) GetProductSku() string {
	if x != nil {
		return x.ProductSku
	}
	return ""
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 引用上面购物车商品
	Item *CartItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *AddItemRequest) Reset() {
	*x = AddItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemRequest) ProtoMessage() {}

func (x *AddItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemRequest.ProtoReflect.Descriptor instead.
func (*AddItemRequest) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{2}
}

func (x *AddItemRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddItemRequest) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items  []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{3}
}

func (x *Cart) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Cart) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{4}
}

func (x *GetCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type EmptyCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *EmptyCartRequest) Reset() {
	*x = EmptyCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyCartRequest) ProtoMessage() {}

func (x *EmptyCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyCartRequest.ProtoReflect.Descriptor instead.
func (*EmptyCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{5}
}

func (x *EmptyCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// ---------------商品服务------------------
// 分页
type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{6}
}

func (x *PageRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32      `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32      `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Total    int64      `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Products []*Product `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *PageResponse) Reset() {
	*x = PageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageResponse) ProtoMessage() {}

func (x *PageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageResponse.ProtoReflect.Descriptor instead.
func (*PageResponse) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{7}
}

func (x *PageResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type Money struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	Units        int64  `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	Nanos        int32  `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Money) Reset() {
	*x = Money{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Money) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Money) ProtoMessage() {}

func (x *Money) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Money.ProtoReflect.Descriptor instead.
func (*Money) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{8}
}

func (x *Money) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Money) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Money) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Picture     string   `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	PriceUsd    *Money   `protobuf:"bytes,5,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	Categories  []string `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{9}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Product) GetPriceUsd() *Money {
	if x != nil {
		return x.PriceUsd
	}
	return nil
}

func (x *Product) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

type ListProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *PageResponse `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{10}
}

func (x *ListProductsResponse) GetPage() *PageResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

// 根据id查询
type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{11}
}

func (x *GetProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SearchProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string       `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Page  *PageRequest `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchProductsRequest) Reset() {
	*x = SearchProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsRequest) ProtoMessage() {}

func (x *SearchProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsRequest.ProtoReflect.Descriptor instead.
func (*SearchProductsRequest) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{12}
}

func (x *SearchProductsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchProductsRequest) GetPage() *PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

// 搜索商品响应
type SearchProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *PageResponse `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchProductsResponse) Reset() {
	*x = SearchProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cartservice_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsResponse) ProtoMessage() {}

func (x *SearchProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartservice_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsResponse.ProtoReflect.Descriptor instead.
func (*SearchProductsResponse) Descriptor() ([]byte, []int) {
	return file_proto_cartservice_proto_rawDescGZIP(), []int{13}
}

func (x *SearchProductsResponse) GetPage() *PageResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_proto_cartservice_proto protoreflect.FileDescriptor

var file_proto_cartservice_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x47, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x56, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x22, 0x4e, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2b, 0x0a,
	0x10, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x0b, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x0c, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x58, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61,
	0x6e, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73,
	0x22, 0xbc, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x09,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x55, 0x73, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22,
	0x47, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a,
	0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x49, 0x0a, 0x16,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0xd6, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x32, 0x95, 0x02, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cartservice_proto_rawDescOnce sync.Once
	file_proto_cartservice_proto_rawDescData = file_proto_cartservice_proto_rawDesc
)

func file_proto_cartservice_proto_rawDescGZIP() []byte {
	file_proto_cartservice_proto_rawDescOnce.Do(func() {
		file_proto_cartservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cartservice_proto_rawDescData)
	})
	return file_proto_cartservice_proto_rawDescData
}

var file_proto_cartservice_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_cartservice_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: microshopping.Empty
	(*CartItem)(nil),               // 1: microshopping.CartItem
	(*AddItemRequest)(nil),         // 2: microshopping.AddItemRequest
	(*Cart)(nil),                   // 3: microshopping.Cart
	(*GetCartRequest)(nil),         // 4: microshopping.GetCartRequest
	(*EmptyCartRequest)(nil),       // 5: microshopping.EmptyCartRequest
	(*PageRequest)(nil),            // 6: microshopping.PageRequest
	(*PageResponse)(nil),           // 7: microshopping.PageResponse
	(*Money)(nil),                  // 8: microshopping.Money
	(*Product)(nil),                // 9: microshopping.Product
	(*ListProductsResponse)(nil),   // 10: microshopping.ListProductsResponse
	(*GetProductRequest)(nil),      // 11: microshopping.GetProductRequest
	(*SearchProductsRequest)(nil),  // 12: microshopping.SearchProductsRequest
	(*SearchProductsResponse)(nil), // 13: microshopping.SearchProductsResponse
}
var file_proto_cartservice_proto_depIdxs = []int32{
	1,  // 0: microshopping.AddItemRequest.item:type_name -> microshopping.CartItem
	1,  // 1: microshopping.Cart.items:type_name -> microshopping.CartItem
	9,  // 2: microshopping.PageResponse.products:type_name -> microshopping.Product
	8,  // 3: microshopping.Product.price_usd:type_name -> microshopping.Money
	7,  // 4: microshopping.ListProductsResponse.page:type_name -> microshopping.PageResponse
	6,  // 5: microshopping.SearchProductsRequest.page:type_name -> microshopping.PageRequest
	7,  // 6: microshopping.SearchProductsResponse.page:type_name -> microshopping.PageResponse
	2,  // 7: microshopping.CartService.AddItem:input_type -> microshopping.AddItemRequest
	4,  // 8: microshopping.CartService.GetCart:input_type -> microshopping.GetCartRequest
	5,  // 9: microshopping.CartService.EmptyCart:input_type -> microshopping.EmptyCartRequest
	6,  // 10: microshopping.ProductCatalogService.ListProducts:input_type -> microshopping.PageRequest
	11, // 11: microshopping.ProductCatalogService.GetProduct:input_type -> microshopping.GetProductRequest
	12, // 12: microshopping.ProductCatalogService.SearchProducts:input_type -> microshopping.SearchProductsRequest
	0,  // 13: microshopping.CartService.AddItem:output_type -> microshopping.Empty
	3,  // 14: microshopping.CartService.GetCart:output_type -> microshopping.Cart
	0,  // 15: microshopping.CartService.EmptyCart:output_type -> microshopping.Empty
	10, // 16: microshopping.ProductCatalogService.ListProducts:output_type -> microshopping.ListProductsResponse
	9,  // 17: microshopping.ProductCatalogService.GetProduct:output_type -> microshopping.Product
	13, // 18: microshopping.ProductCatalogService.SearchProducts:output_type -> microshopping.SearchProductsResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_cartservice_proto_init() }
func file_proto_cartservice_proto_init() {
	if File_proto_cartservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cartservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_cartservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_proto_cartservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemRequest); i {
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
		file_proto_cartservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
		file_proto_cartservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartRequest); i {
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
		file_proto_cartservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyCartRequest); i {
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
		file_proto_cartservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
		file_proto_cartservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageResponse); i {
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
		file_proto_cartservice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Money); i {
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
		file_proto_cartservice_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_proto_cartservice_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsResponse); i {
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
		file_proto_cartservice_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRequest); i {
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
		file_proto_cartservice_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProductsRequest); i {
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
		file_proto_cartservice_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProductsResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_cartservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_cartservice_proto_goTypes,
		DependencyIndexes: file_proto_cartservice_proto_depIdxs,
		MessageInfos:      file_proto_cartservice_proto_msgTypes,
	}.Build()
	File_proto_cartservice_proto = out.File
	file_proto_cartservice_proto_rawDesc = nil
	file_proto_cartservice_proto_goTypes = nil
	file_proto_cartservice_proto_depIdxs = nil
}