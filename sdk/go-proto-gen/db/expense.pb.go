// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: expense.proto

package db

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Expense struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpenseId string                 `protobuf:"bytes,1,opt,name=expense_id,json=expenseId,proto3" json:"expense_id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Amount    float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Category  string                 `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Date      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Expense) Reset() {
	*x = Expense{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expense) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expense) ProtoMessage() {}

func (x *Expense) ProtoReflect() protoreflect.Message {
	mi := &file_expense_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expense.ProtoReflect.Descriptor instead.
func (*Expense) Descriptor() ([]byte, []int) {
	return file_expense_proto_rawDescGZIP(), []int{0}
}

func (x *Expense) GetExpenseId() string {
	if x != nil {
		return x.ExpenseId
	}
	return ""
}

func (x *Expense) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Expense) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Expense) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Expense) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type CreateExpenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expense *Expense `protobuf:"bytes,1,opt,name=expense,proto3" json:"expense,omitempty"`
}

func (x *CreateExpenseRequest) Reset() {
	*x = CreateExpenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExpenseRequest) ProtoMessage() {}

func (x *CreateExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_expense_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExpenseRequest.ProtoReflect.Descriptor instead.
func (*CreateExpenseRequest) Descriptor() ([]byte, []int) {
	return file_expense_proto_rawDescGZIP(), []int{1}
}

func (x *CreateExpenseRequest) GetExpense() *Expense {
	if x != nil {
		return x.Expense
	}
	return nil
}

type CreateExpenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expense *Expense `protobuf:"bytes,1,opt,name=expense,proto3" json:"expense,omitempty"`
}

func (x *CreateExpenseResponse) Reset() {
	*x = CreateExpenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExpenseResponse) ProtoMessage() {}

func (x *CreateExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_expense_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExpenseResponse.ProtoReflect.Descriptor instead.
func (*CreateExpenseResponse) Descriptor() ([]byte, []int) {
	return file_expense_proto_rawDescGZIP(), []int{2}
}

func (x *CreateExpenseResponse) GetExpense() *Expense {
	if x != nil {
		return x.Expense
	}
	return nil
}

type GetExpenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpenseId string `protobuf:"bytes,1,opt,name=expense_id,json=expenseId,proto3" json:"expense_id,omitempty"`
}

func (x *GetExpenseRequest) Reset() {
	*x = GetExpenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExpenseRequest) ProtoMessage() {}

func (x *GetExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_expense_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExpenseRequest.ProtoReflect.Descriptor instead.
func (*GetExpenseRequest) Descriptor() ([]byte, []int) {
	return file_expense_proto_rawDescGZIP(), []int{3}
}

func (x *GetExpenseRequest) GetExpenseId() string {
	if x != nil {
		return x.ExpenseId
	}
	return ""
}

type GetExpenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expense *Expense `protobuf:"bytes,1,opt,name=expense,proto3" json:"expense,omitempty"`
}

func (x *GetExpenseResponse) Reset() {
	*x = GetExpenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExpenseResponse) ProtoMessage() {}

func (x *GetExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_expense_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExpenseResponse.ProtoReflect.Descriptor instead.
func (*GetExpenseResponse) Descriptor() ([]byte, []int) {
	return file_expense_proto_rawDescGZIP(), []int{4}
}

func (x *GetExpenseResponse) GetExpense() *Expense {
	if x != nil {
		return x.Expense
	}
	return nil
}

type UpdateExpenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expense *Expense `protobuf:"bytes,1,opt,name=expense,proto3" json:"expense,omitempty"`
}

func (x *UpdateExpenseRequest) Reset() {
	*x = UpdateExpenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExpenseRequest) ProtoMessage() {}

func (x *UpdateExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_expense_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExpenseRequest.ProtoReflect.Descriptor instead.
func (*UpdateExpenseRequest) Descriptor() ([]byte, []int) {
	return file_expense_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateExpenseRequest) GetExpense() *Expense {
	if x != nil {
		return x.Expense
	}
	return nil
}

type UpdateExpenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expense *Expense `protobuf:"bytes,1,opt,name=expense,proto3" json:"expense,omitempty"`
}

func (x *UpdateExpenseResponse) Reset() {
	*x = UpdateExpenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExpenseResponse) ProtoMessage() {}

func (x *UpdateExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_expense_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExpenseResponse.ProtoReflect.Descriptor instead.
func (*UpdateExpenseResponse) Descriptor() ([]byte, []int) {
	return file_expense_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateExpenseResponse) GetExpense() *Expense {
	if x != nil {
		return x.Expense
	}
	return nil
}

type DeleteExpenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpenseId string `protobuf:"bytes,1,opt,name=expense_id,json=expenseId,proto3" json:"expense_id,omitempty"`
}

func (x *DeleteExpenseRequest) Reset() {
	*x = DeleteExpenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExpenseRequest) ProtoMessage() {}

func (x *DeleteExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_expense_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExpenseRequest.ProtoReflect.Descriptor instead.
func (*DeleteExpenseRequest) Descriptor() ([]byte, []int) {
	return file_expense_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteExpenseRequest) GetExpenseId() string {
	if x != nil {
		return x.ExpenseId
	}
	return ""
}

type DeleteExpenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteExpenseResponse) Reset() {
	*x = DeleteExpenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExpenseResponse) ProtoMessage() {}

func (x *DeleteExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_expense_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExpenseResponse.ProtoReflect.Descriptor instead.
func (*DeleteExpenseResponse) Descriptor() ([]byte, []int) {
	return file_expense_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteExpenseResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_expense_proto protoreflect.FileDescriptor

var file_expense_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x64, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x3d, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52,
	0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52,
	0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x52, 0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52,
	0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52,
	0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x22,
	0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x69, 0x6b, 0x65, 0x4d, 0x77, 0x69, 0x74, 0x61, 0x2f, 0x66, 0x65, 0x64, 0x68, 0x61,
	0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_expense_proto_rawDescOnce sync.Once
	file_expense_proto_rawDescData = file_expense_proto_rawDesc
)

func file_expense_proto_rawDescGZIP() []byte {
	file_expense_proto_rawDescOnce.Do(func() {
		file_expense_proto_rawDescData = protoimpl.X.CompressGZIP(file_expense_proto_rawDescData)
	})
	return file_expense_proto_rawDescData
}

var file_expense_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_expense_proto_goTypes = []interface{}{
	(*Expense)(nil),               // 0: db.Expense
	(*CreateExpenseRequest)(nil),  // 1: db.CreateExpenseRequest
	(*CreateExpenseResponse)(nil), // 2: db.CreateExpenseResponse
	(*GetExpenseRequest)(nil),     // 3: db.GetExpenseRequest
	(*GetExpenseResponse)(nil),    // 4: db.GetExpenseResponse
	(*UpdateExpenseRequest)(nil),  // 5: db.UpdateExpenseRequest
	(*UpdateExpenseResponse)(nil), // 6: db.UpdateExpenseResponse
	(*DeleteExpenseRequest)(nil),  // 7: db.DeleteExpenseRequest
	(*DeleteExpenseResponse)(nil), // 8: db.DeleteExpenseResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_expense_proto_depIdxs = []int32{
	9, // 0: db.Expense.date:type_name -> google.protobuf.Timestamp
	0, // 1: db.CreateExpenseRequest.expense:type_name -> db.Expense
	0, // 2: db.CreateExpenseResponse.expense:type_name -> db.Expense
	0, // 3: db.GetExpenseResponse.expense:type_name -> db.Expense
	0, // 4: db.UpdateExpenseRequest.expense:type_name -> db.Expense
	0, // 5: db.UpdateExpenseResponse.expense:type_name -> db.Expense
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_expense_proto_init() }
func file_expense_proto_init() {
	if File_expense_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_expense_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expense); i {
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
		file_expense_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExpenseRequest); i {
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
		file_expense_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExpenseResponse); i {
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
		file_expense_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExpenseRequest); i {
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
		file_expense_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExpenseResponse); i {
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
		file_expense_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExpenseRequest); i {
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
		file_expense_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExpenseResponse); i {
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
		file_expense_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExpenseRequest); i {
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
		file_expense_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExpenseResponse); i {
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
			RawDescriptor: file_expense_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_expense_proto_goTypes,
		DependencyIndexes: file_expense_proto_depIdxs,
		MessageInfos:      file_expense_proto_msgTypes,
	}.Build()
	File_expense_proto = out.File
	file_expense_proto_rawDesc = nil
	file_expense_proto_goTypes = nil
	file_expense_proto_depIdxs = nil
}
