// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: dbserver.proto

package db

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DbServiceClient is the client API for DbService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DbServiceClient interface {
	CreateExpense(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error)
	GetExpense(ctx context.Context, in *GetExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error)
	UpdateExpense(ctx context.Context, in *UpdateExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error)
	DeleteExpense(ctx context.Context, in *DeleteExpenseRequest, opts ...grpc.CallOption) (*DeleteExpenseResponse, error)
	// INCOMES
	CreateIncome(ctx context.Context, in *CreateIncomeRequest, opts ...grpc.CallOption) (*CreateIncomeResponse, error)
	GetIncome(ctx context.Context, in *GetIncomeRequest, opts ...grpc.CallOption) (*GetIncomeResponse, error)
	UpdateIncome(ctx context.Context, in *UpdateIncomeRequest, opts ...grpc.CallOption) (*UpdateIncomeResponse, error)
	DeleteIncome(ctx context.Context, in *DeleteIncomeRequest, opts ...grpc.CallOption) (*DeleteIncomeResponse, error)
	// BALANCES
	GetRemainingBalance(ctx context.Context, in *RemainingBalanceRequest, opts ...grpc.CallOption) (*RemainingBalanceResponse, error)
	// MONTHLY SUMMARY
	GenerateMonthlySummary(ctx context.Context, in *MonthlySummaryRequest, opts ...grpc.CallOption) (*MonthlySummaryResponse, error)
}

type dbServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDbServiceClient(cc grpc.ClientConnInterface) DbServiceClient {
	return &dbServiceClient{cc}
}

func (c *dbServiceClient) CreateExpense(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error) {
	out := new(ExpenseResponse)
	err := c.cc.Invoke(ctx, "/db.DbService/CreateExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) GetExpense(ctx context.Context, in *GetExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error) {
	out := new(ExpenseResponse)
	err := c.cc.Invoke(ctx, "/db.DbService/GetExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) UpdateExpense(ctx context.Context, in *UpdateExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error) {
	out := new(ExpenseResponse)
	err := c.cc.Invoke(ctx, "/db.DbService/UpdateExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) DeleteExpense(ctx context.Context, in *DeleteExpenseRequest, opts ...grpc.CallOption) (*DeleteExpenseResponse, error) {
	out := new(DeleteExpenseResponse)
	err := c.cc.Invoke(ctx, "/db.DbService/DeleteExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) CreateIncome(ctx context.Context, in *CreateIncomeRequest, opts ...grpc.CallOption) (*CreateIncomeResponse, error) {
	out := new(CreateIncomeResponse)
	err := c.cc.Invoke(ctx, "/db.DbService/CreateIncome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) GetIncome(ctx context.Context, in *GetIncomeRequest, opts ...grpc.CallOption) (*GetIncomeResponse, error) {
	out := new(GetIncomeResponse)
	err := c.cc.Invoke(ctx, "/db.DbService/GetIncome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) UpdateIncome(ctx context.Context, in *UpdateIncomeRequest, opts ...grpc.CallOption) (*UpdateIncomeResponse, error) {
	out := new(UpdateIncomeResponse)
	err := c.cc.Invoke(ctx, "/db.DbService/UpdateIncome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) DeleteIncome(ctx context.Context, in *DeleteIncomeRequest, opts ...grpc.CallOption) (*DeleteIncomeResponse, error) {
	out := new(DeleteIncomeResponse)
	err := c.cc.Invoke(ctx, "/db.DbService/DeleteIncome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) GetRemainingBalance(ctx context.Context, in *RemainingBalanceRequest, opts ...grpc.CallOption) (*RemainingBalanceResponse, error) {
	out := new(RemainingBalanceResponse)
	err := c.cc.Invoke(ctx, "/db.DbService/GetRemainingBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) GenerateMonthlySummary(ctx context.Context, in *MonthlySummaryRequest, opts ...grpc.CallOption) (*MonthlySummaryResponse, error) {
	out := new(MonthlySummaryResponse)
	err := c.cc.Invoke(ctx, "/db.DbService/GenerateMonthlySummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DbServiceServer is the server API for DbService service.
// All implementations must embed UnimplementedDbServiceServer
// for forward compatibility
type DbServiceServer interface {
	CreateExpense(context.Context, *ExpenseRequest) (*ExpenseResponse, error)
	GetExpense(context.Context, *GetExpenseRequest) (*ExpenseResponse, error)
	UpdateExpense(context.Context, *UpdateExpenseRequest) (*ExpenseResponse, error)
	DeleteExpense(context.Context, *DeleteExpenseRequest) (*DeleteExpenseResponse, error)
	// INCOMES
	CreateIncome(context.Context, *CreateIncomeRequest) (*CreateIncomeResponse, error)
	GetIncome(context.Context, *GetIncomeRequest) (*GetIncomeResponse, error)
	UpdateIncome(context.Context, *UpdateIncomeRequest) (*UpdateIncomeResponse, error)
	DeleteIncome(context.Context, *DeleteIncomeRequest) (*DeleteIncomeResponse, error)
	// BALANCES
	GetRemainingBalance(context.Context, *RemainingBalanceRequest) (*RemainingBalanceResponse, error)
	// MONTHLY SUMMARY
	GenerateMonthlySummary(context.Context, *MonthlySummaryRequest) (*MonthlySummaryResponse, error)
	mustEmbedUnimplementedDbServiceServer()
}

// UnimplementedDbServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDbServiceServer struct {
}

func (UnimplementedDbServiceServer) CreateExpense(context.Context, *ExpenseRequest) (*ExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpense not implemented")
}
func (UnimplementedDbServiceServer) GetExpense(context.Context, *GetExpenseRequest) (*ExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpense not implemented")
}
func (UnimplementedDbServiceServer) UpdateExpense(context.Context, *UpdateExpenseRequest) (*ExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExpense not implemented")
}
func (UnimplementedDbServiceServer) DeleteExpense(context.Context, *DeleteExpenseRequest) (*DeleteExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExpense not implemented")
}
func (UnimplementedDbServiceServer) CreateIncome(context.Context, *CreateIncomeRequest) (*CreateIncomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIncome not implemented")
}
func (UnimplementedDbServiceServer) GetIncome(context.Context, *GetIncomeRequest) (*GetIncomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncome not implemented")
}
func (UnimplementedDbServiceServer) UpdateIncome(context.Context, *UpdateIncomeRequest) (*UpdateIncomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIncome not implemented")
}
func (UnimplementedDbServiceServer) DeleteIncome(context.Context, *DeleteIncomeRequest) (*DeleteIncomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIncome not implemented")
}
func (UnimplementedDbServiceServer) GetRemainingBalance(context.Context, *RemainingBalanceRequest) (*RemainingBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRemainingBalance not implemented")
}
func (UnimplementedDbServiceServer) GenerateMonthlySummary(context.Context, *MonthlySummaryRequest) (*MonthlySummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateMonthlySummary not implemented")
}
func (UnimplementedDbServiceServer) mustEmbedUnimplementedDbServiceServer() {}

// UnsafeDbServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DbServiceServer will
// result in compilation errors.
type UnsafeDbServiceServer interface {
	mustEmbedUnimplementedDbServiceServer()
}

func RegisterDbServiceServer(s grpc.ServiceRegistrar, srv DbServiceServer) {
	s.RegisterService(&DbService_ServiceDesc, srv)
}

func _DbService_CreateExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).CreateExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DbService/CreateExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).CreateExpense(ctx, req.(*ExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_GetExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).GetExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DbService/GetExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).GetExpense(ctx, req.(*GetExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_UpdateExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).UpdateExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DbService/UpdateExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).UpdateExpense(ctx, req.(*UpdateExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_DeleteExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).DeleteExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DbService/DeleteExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).DeleteExpense(ctx, req.(*DeleteExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_CreateIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIncomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).CreateIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DbService/CreateIncome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).CreateIncome(ctx, req.(*CreateIncomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_GetIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).GetIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DbService/GetIncome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).GetIncome(ctx, req.(*GetIncomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_UpdateIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIncomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).UpdateIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DbService/UpdateIncome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).UpdateIncome(ctx, req.(*UpdateIncomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_DeleteIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIncomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).DeleteIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DbService/DeleteIncome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).DeleteIncome(ctx, req.(*DeleteIncomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_GetRemainingBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemainingBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).GetRemainingBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DbService/GetRemainingBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).GetRemainingBalance(ctx, req.(*RemainingBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_GenerateMonthlySummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonthlySummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).GenerateMonthlySummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DbService/GenerateMonthlySummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).GenerateMonthlySummary(ctx, req.(*MonthlySummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DbService_ServiceDesc is the grpc.ServiceDesc for DbService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DbService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.DbService",
	HandlerType: (*DbServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExpense",
			Handler:    _DbService_CreateExpense_Handler,
		},
		{
			MethodName: "GetExpense",
			Handler:    _DbService_GetExpense_Handler,
		},
		{
			MethodName: "UpdateExpense",
			Handler:    _DbService_UpdateExpense_Handler,
		},
		{
			MethodName: "DeleteExpense",
			Handler:    _DbService_DeleteExpense_Handler,
		},
		{
			MethodName: "CreateIncome",
			Handler:    _DbService_CreateIncome_Handler,
		},
		{
			MethodName: "GetIncome",
			Handler:    _DbService_GetIncome_Handler,
		},
		{
			MethodName: "UpdateIncome",
			Handler:    _DbService_UpdateIncome_Handler,
		},
		{
			MethodName: "DeleteIncome",
			Handler:    _DbService_DeleteIncome_Handler,
		},
		{
			MethodName: "GetRemainingBalance",
			Handler:    _DbService_GetRemainingBalance_Handler,
		},
		{
			MethodName: "GenerateMonthlySummary",
			Handler:    _DbService_GenerateMonthlySummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dbserver.proto",
}
