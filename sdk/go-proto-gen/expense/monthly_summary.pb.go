// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: monthly_summary.proto

package expense

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

type MonthlySummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *MonthlySummaryRequest) Reset() {
	*x = MonthlySummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monthly_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonthlySummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthlySummaryRequest) ProtoMessage() {}

func (x *MonthlySummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monthly_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthlySummaryRequest.ProtoReflect.Descriptor instead.
func (*MonthlySummaryRequest) Descriptor() ([]byte, []int) {
	return file_monthly_summary_proto_rawDescGZIP(), []int{0}
}

func (x *MonthlySummaryRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *MonthlySummaryRequest) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

type MonthlySummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year             int32             `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month            int32             `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	TotalExpenses    float64           `protobuf:"fixed64,3,opt,name=total_expenses,json=totalExpenses,proto3" json:"total_expenses,omitempty"`
	TotalIncome      float64           `protobuf:"fixed64,4,opt,name=total_income,json=totalIncome,proto3" json:"total_income,omitempty"`
	RemainingBalance float64           `protobuf:"fixed64,5,opt,name=remaining_balance,json=remainingBalance,proto3" json:"remaining_balance,omitempty"`
	Expenses         []*ExpenseSummary `protobuf:"bytes,6,rep,name=expenses,proto3" json:"expenses,omitempty"`
}

func (x *MonthlySummaryResponse) Reset() {
	*x = MonthlySummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monthly_summary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonthlySummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthlySummaryResponse) ProtoMessage() {}

func (x *MonthlySummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monthly_summary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthlySummaryResponse.ProtoReflect.Descriptor instead.
func (*MonthlySummaryResponse) Descriptor() ([]byte, []int) {
	return file_monthly_summary_proto_rawDescGZIP(), []int{1}
}

func (x *MonthlySummaryResponse) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *MonthlySummaryResponse) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *MonthlySummaryResponse) GetTotalExpenses() float64 {
	if x != nil {
		return x.TotalExpenses
	}
	return 0
}

func (x *MonthlySummaryResponse) GetTotalIncome() float64 {
	if x != nil {
		return x.TotalIncome
	}
	return 0
}

func (x *MonthlySummaryResponse) GetRemainingBalance() float64 {
	if x != nil {
		return x.RemainingBalance
	}
	return 0
}

func (x *MonthlySummaryResponse) GetExpenses() []*ExpenseSummary {
	if x != nil {
		return x.Expenses
	}
	return nil
}

type ExpenseSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Amount   float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Category string  `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *ExpenseSummary) Reset() {
	*x = ExpenseSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monthly_summary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpenseSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpenseSummary) ProtoMessage() {}

func (x *ExpenseSummary) ProtoReflect() protoreflect.Message {
	mi := &file_monthly_summary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpenseSummary.ProtoReflect.Descriptor instead.
func (*ExpenseSummary) Descriptor() ([]byte, []int) {
	return file_monthly_summary_proto_rawDescGZIP(), []int{2}
}

func (x *ExpenseSummary) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ExpenseSummary) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ExpenseSummary) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_monthly_summary_proto protoreflect.FileDescriptor

var file_monthly_summary_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x22, 0x41, 0x0a, 0x15, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x22, 0xee, 0x01, 0x0a, 0x16, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x33, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x08, 0x65, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d,
	0x69, 0x6b, 0x65, 0x4d, 0x77, 0x69, 0x74, 0x61, 0x2f, 0x66, 0x65, 0x64, 0x68, 0x61, 0x2d, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monthly_summary_proto_rawDescOnce sync.Once
	file_monthly_summary_proto_rawDescData = file_monthly_summary_proto_rawDesc
)

func file_monthly_summary_proto_rawDescGZIP() []byte {
	file_monthly_summary_proto_rawDescOnce.Do(func() {
		file_monthly_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_monthly_summary_proto_rawDescData)
	})
	return file_monthly_summary_proto_rawDescData
}

var file_monthly_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_monthly_summary_proto_goTypes = []interface{}{
	(*MonthlySummaryRequest)(nil),  // 0: expense.MonthlySummaryRequest
	(*MonthlySummaryResponse)(nil), // 1: expense.MonthlySummaryResponse
	(*ExpenseSummary)(nil),         // 2: expense.ExpenseSummary
}
var file_monthly_summary_proto_depIdxs = []int32{
	2, // 0: expense.MonthlySummaryResponse.expenses:type_name -> expense.ExpenseSummary
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_monthly_summary_proto_init() }
func file_monthly_summary_proto_init() {
	if File_monthly_summary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monthly_summary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonthlySummaryRequest); i {
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
		file_monthly_summary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonthlySummaryResponse); i {
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
		file_monthly_summary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpenseSummary); i {
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
			RawDescriptor: file_monthly_summary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_monthly_summary_proto_goTypes,
		DependencyIndexes: file_monthly_summary_proto_depIdxs,
		MessageInfos:      file_monthly_summary_proto_msgTypes,
	}.Build()
	File_monthly_summary_proto = out.File
	file_monthly_summary_proto_rawDesc = nil
	file_monthly_summary_proto_goTypes = nil
	file_monthly_summary_proto_depIdxs = nil
}