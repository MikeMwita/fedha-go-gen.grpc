// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: expense_server.proto

package expense

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_expense_server_proto protoreflect.FileDescriptor

var file_expense_server_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x1a,
	0x0d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8d,
	0x0a, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70,
	0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65,
	0x12, 0x1c, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x44, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x19, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x65, 0x78, 0x70,
	0x65, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5b, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x13, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x16, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x65, 0x78, 0x70,
	0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x69, 0x6b,
	0x65, 0x4d, 0x77, 0x69, 0x74, 0x61, 0x2f, 0x66, 0x65, 0x64, 0x68, 0x61, 0x2d, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_expense_server_proto_goTypes = []interface{}{
	(*CreateExpenseRequest)(nil),     // 0: expense.CreateExpenseRequest
	(*GetExpenseRequest)(nil),        // 1: expense.GetExpenseRequest
	(*UpdateExpenseRequest)(nil),     // 2: expense.UpdateExpenseRequest
	(*DeleteExpenseRequest)(nil),     // 3: expense.DeleteExpenseRequest
	(*CreateIncomeRequest)(nil),      // 4: expense.CreateIncomeRequest
	(*GetIncomeRequest)(nil),         // 5: expense.GetIncomeRequest
	(*UpdateIncomeRequest)(nil),      // 6: expense.UpdateIncomeRequest
	(*DeleteIncomeRequest)(nil),      // 7: expense.DeleteIncomeRequest
	(*RemainingBalanceRequest)(nil),  // 8: expense.RemainingBalanceRequest
	(*MonthlySummaryRequest)(nil),    // 9: expense.MonthlySummaryRequest
	(*RegUserReq)(nil),               // 10: expense.RegUserReq
	(*UpdateUserReq)(nil),            // 11: expense.UpdateUserReq
	(*GetPagedUsersReq)(nil),         // 12: expense.GetPagedUsersReq
	(*GetByfieldReq)(nil),            // 13: expense.GetByfieldReq
	(*GetUserByUsernameRequest)(nil), // 14: expense.GetUserByUsernameRequest
	(*GetUserByIDRequest)(nil),       // 15: expense.GetUserByIDRequest
	(*SaveUserRequest)(nil),          // 16: expense.SaveUserRequest
	(*CreateExpenseResponse)(nil),    // 17: expense.CreateExpenseResponse
	(*GetExpenseResponse)(nil),       // 18: expense.GetExpenseResponse
	(*UpdateExpenseResponse)(nil),    // 19: expense.UpdateExpenseResponse
	(*DeleteExpenseResponse)(nil),    // 20: expense.DeleteExpenseResponse
	(*CreateIncomeResponse)(nil),     // 21: expense.CreateIncomeResponse
	(*GetIncomeResponse)(nil),        // 22: expense.GetIncomeResponse
	(*UpdateIncomeResponse)(nil),     // 23: expense.UpdateIncomeResponse
	(*DeleteIncomeResponse)(nil),     // 24: expense.DeleteIncomeResponse
	(*RemainingBalanceResponse)(nil), // 25: expense.RemainingBalanceResponse
	(*MonthlySummaryResponse)(nil),   // 26: expense.MonthlySummaryResponse
	(*RegUserRes)(nil),               // 27: expense.RegUserRes
	(*UpdateUserRes)(nil),            // 28: expense.UpdateUserRes
	(*GetPagedUsersRes)(nil),         // 29: expense.GetPagedUsersRes
	(*GetByfieldRes)(nil),            // 30: expense.GetByfieldRes
	(*User)(nil),                     // 31: expense.User
}
var file_expense_server_proto_depIdxs = []int32{
	0,  // 0: expense.ExpenseService.CreateExpense:input_type -> expense.CreateExpenseRequest
	1,  // 1: expense.ExpenseService.GetExpense:input_type -> expense.GetExpenseRequest
	2,  // 2: expense.ExpenseService.UpdateExpense:input_type -> expense.UpdateExpenseRequest
	3,  // 3: expense.ExpenseService.DeleteExpense:input_type -> expense.DeleteExpenseRequest
	4,  // 4: expense.ExpenseService.CreateIncome:input_type -> expense.CreateIncomeRequest
	5,  // 5: expense.ExpenseService.GetIncome:input_type -> expense.GetIncomeRequest
	6,  // 6: expense.ExpenseService.UpdateIncome:input_type -> expense.UpdateIncomeRequest
	7,  // 7: expense.ExpenseService.DeleteIncome:input_type -> expense.DeleteIncomeRequest
	8,  // 8: expense.ExpenseService.GetRemainingBalance:input_type -> expense.RemainingBalanceRequest
	9,  // 9: expense.ExpenseService.GenerateMonthlySummary:input_type -> expense.MonthlySummaryRequest
	10, // 10: expense.ExpenseService.CreateUser:input_type -> expense.RegUserReq
	11, // 11: expense.ExpenseService.UpdateUser:input_type -> expense.UpdateUserReq
	12, // 12: expense.ExpenseService.GetPagedUsers:input_type -> expense.GetPagedUsersReq
	13, // 13: expense.ExpenseService.GetUserByField:input_type -> expense.GetByfieldReq
	14, // 14: expense.ExpenseService.GetUserByUsername:input_type -> expense.GetUserByUsernameRequest
	15, // 15: expense.ExpenseService.GetUserByID:input_type -> expense.GetUserByIDRequest
	16, // 16: expense.ExpenseService.SaveUser:input_type -> expense.SaveUserRequest
	17, // 17: expense.ExpenseService.CreateExpense:output_type -> expense.CreateExpenseResponse
	18, // 18: expense.ExpenseService.GetExpense:output_type -> expense.GetExpenseResponse
	19, // 19: expense.ExpenseService.UpdateExpense:output_type -> expense.UpdateExpenseResponse
	20, // 20: expense.ExpenseService.DeleteExpense:output_type -> expense.DeleteExpenseResponse
	21, // 21: expense.ExpenseService.CreateIncome:output_type -> expense.CreateIncomeResponse
	22, // 22: expense.ExpenseService.GetIncome:output_type -> expense.GetIncomeResponse
	23, // 23: expense.ExpenseService.UpdateIncome:output_type -> expense.UpdateIncomeResponse
	24, // 24: expense.ExpenseService.DeleteIncome:output_type -> expense.DeleteIncomeResponse
	25, // 25: expense.ExpenseService.GetRemainingBalance:output_type -> expense.RemainingBalanceResponse
	26, // 26: expense.ExpenseService.GenerateMonthlySummary:output_type -> expense.MonthlySummaryResponse
	27, // 27: expense.ExpenseService.CreateUser:output_type -> expense.RegUserRes
	28, // 28: expense.ExpenseService.UpdateUser:output_type -> expense.UpdateUserRes
	29, // 29: expense.ExpenseService.GetPagedUsers:output_type -> expense.GetPagedUsersRes
	30, // 30: expense.ExpenseService.GetUserByField:output_type -> expense.GetByfieldRes
	27, // 31: expense.ExpenseService.GetUserByUsername:output_type -> expense.RegUserRes
	27, // 32: expense.ExpenseService.GetUserByID:output_type -> expense.RegUserRes
	31, // 33: expense.ExpenseService.SaveUser:output_type -> expense.User
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_expense_server_proto_init() }
func file_expense_server_proto_init() {
	if File_expense_server_proto != nil {
		return
	}
	file_balance_proto_init()
	file_expense_proto_init()
	file_Income_proto_init()
	file_monthly_summary_proto_init()
	file_User_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_expense_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_expense_server_proto_goTypes,
		DependencyIndexes: file_expense_server_proto_depIdxs,
	}.Build()
	File_expense_server_proto = out.File
	file_expense_server_proto_rawDesc = nil
	file_expense_server_proto_goTypes = nil
	file_expense_server_proto_depIdxs = nil
}
