// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: news.proto

package api

import (
	api "awesomeProject"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContentCheckServiceClient is the client API for ContentCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentCheckServiceClient interface {
	// Проверка жизнеспособности сервиса
	CheckHealth(ctx context.Context, in *api.EmptyRequest, opts ...grpc.CallOption) (*api.HealthResponse, error)
}

type contentCheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentCheckServiceClient(cc grpc.ClientConnInterface) ContentCheckServiceClient {
	return &contentCheckServiceClient{cc}
}

func (c *contentCheckServiceClient) CheckHealth(ctx context.Context, in *api.EmptyRequest, opts ...grpc.CallOption) (*api.HealthResponse, error) {
	out := new(api.HealthResponse)
	err := c.cc.Invoke(ctx, "/content_v1.ContentCheckService/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentCheckServiceServer is the server API for ContentCheckService service.
// All implementations must embed UnimplementedContentCheckServiceServer
// for forward compatibility
type ContentCheckServiceServer interface {
	// Проверка жизнеспособности сервиса
	CheckHealth(context.Context, *api.EmptyRequest) (*api.HealthResponse, error)
	mustEmbedUnimplementedContentCheckServiceServer()
}

// UnimplementedContentCheckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentCheckServiceServer struct {
}

func (UnimplementedContentCheckServiceServer) CheckHealth(context.Context, *api.EmptyRequest) (*api.HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (UnimplementedContentCheckServiceServer) mustEmbedUnimplementedContentCheckServiceServer() {}

// UnsafeContentCheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentCheckServiceServer will
// result in compilation errors.
type UnsafeContentCheckServiceServer interface {
	mustEmbedUnimplementedContentCheckServiceServer()
}

func RegisterContentCheckServiceServer(s grpc.ServiceRegistrar, srv ContentCheckServiceServer) {
	s.RegisterService(&ContentCheckService_ServiceDesc, srv)
}

func _ContentCheckService_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentCheckServiceServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_v1.ContentCheckService/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentCheckServiceServer).CheckHealth(ctx, req.(*api.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentCheckService_ServiceDesc is the grpc.ServiceDesc for ContentCheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentCheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content_v1.ContentCheckService",
	HandlerType: (*ContentCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckHealth",
			Handler:    _ContentCheckService_CheckHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news.proto",
}

// NewsServiceClient is the client API for NewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsServiceClient interface {
	// Получение списка новостей
	GetNews(ctx context.Context, in *api.NewsRequestParams, opts ...grpc.CallOption) (*api.NewsList, error)
	// Получение детальной информации по новости
	GetOne(ctx context.Context, in *api.ObjectId, opts ...grpc.CallOption) (*api.NewsObject, error)
	// Получение детальной информации по новости для отображения при переходе по письму
	GetOneBySlug(ctx context.Context, in *api.ObjectSlug, opts ...grpc.CallOption) (*api.NewsObject, error)
	// Создание новости
	Create(ctx context.Context, in *api.RequestNewsObject, opts ...grpc.CallOption) (*api.BaseResponse, error)
	// Обновление новости
	Update(ctx context.Context, in *api.RequestNewsObject, opts ...grpc.CallOption) (*api.BaseResponse, error)
	// Удаление новости
	Delete(ctx context.Context, in *api.ObjectId, opts ...grpc.CallOption) (*api.BaseResponse, error)
	//Получить ссылку на файл для скачивания
	GetFileLink(ctx context.Context, in *api.FileId, opts ...grpc.CallOption) (*api.FileLink, error)
}

type newsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsServiceClient(cc grpc.ClientConnInterface) NewsServiceClient {
	return &newsServiceClient{cc}
}

func (c *newsServiceClient) GetNews(ctx context.Context, in *api.NewsRequestParams, opts ...grpc.CallOption) (*api.NewsList, error) {
	out := new(api.NewsList)
	err := c.cc.Invoke(ctx, "/content_v1.NewsService/GetNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetOne(ctx context.Context, in *api.ObjectId, opts ...grpc.CallOption) (*api.NewsObject, error) {
	out := new(api.NewsObject)
	err := c.cc.Invoke(ctx, "/content_v1.NewsService/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetOneBySlug(ctx context.Context, in *api.ObjectSlug, opts ...grpc.CallOption) (*api.NewsObject, error) {
	out := new(api.NewsObject)
	err := c.cc.Invoke(ctx, "/content_v1.NewsService/GetOneBySlug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) Create(ctx context.Context, in *api.RequestNewsObject, opts ...grpc.CallOption) (*api.BaseResponse, error) {
	out := new(api.BaseResponse)
	err := c.cc.Invoke(ctx, "/content_v1.NewsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) Update(ctx context.Context, in *api.RequestNewsObject, opts ...grpc.CallOption) (*api.BaseResponse, error) {
	out := new(api.BaseResponse)
	err := c.cc.Invoke(ctx, "/content_v1.NewsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) Delete(ctx context.Context, in *api.ObjectId, opts ...grpc.CallOption) (*api.BaseResponse, error) {
	out := new(api.BaseResponse)
	err := c.cc.Invoke(ctx, "/content_v1.NewsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetFileLink(ctx context.Context, in *api.FileId, opts ...grpc.CallOption) (*api.FileLink, error) {
	out := new(api.FileLink)
	err := c.cc.Invoke(ctx, "/content_v1.NewsService/GetFileLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServiceServer is the server API for NewsService service.
// All implementations must embed UnimplementedNewsServiceServer
// for forward compatibility
type NewsServiceServer interface {
	// Получение списка новостей
	GetNews(context.Context, *api.NewsRequestParams) (*api.NewsList, error)
	// Получение детальной информации по новости
	GetOne(context.Context, *api.ObjectId) (*api.NewsObject, error)
	// Получение детальной информации по новости для отображения при переходе по письму
	GetOneBySlug(context.Context, *api.ObjectSlug) (*api.NewsObject, error)
	// Создание новости
	Create(context.Context, *api.RequestNewsObject) (*api.BaseResponse, error)
	// Обновление новости
	Update(context.Context, *api.RequestNewsObject) (*api.BaseResponse, error)
	// Удаление новости
	Delete(context.Context, *api.ObjectId) (*api.BaseResponse, error)
	//Получить ссылку на файл для скачивания
	GetFileLink(context.Context, *api.FileId) (*api.FileLink, error)
	mustEmbedUnimplementedNewsServiceServer()
}

// UnimplementedNewsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNewsServiceServer struct {
}

func (UnimplementedNewsServiceServer) GetNews(context.Context, *api.NewsRequestParams) (*api.NewsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNews not implemented")
}
func (UnimplementedNewsServiceServer) GetOne(context.Context, *api.ObjectId) (*api.NewsObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedNewsServiceServer) GetOneBySlug(context.Context, *api.ObjectSlug) (*api.NewsObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneBySlug not implemented")
}
func (UnimplementedNewsServiceServer) Create(context.Context, *api.RequestNewsObject) (*api.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNewsServiceServer) Update(context.Context, *api.RequestNewsObject) (*api.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNewsServiceServer) Delete(context.Context, *api.ObjectId) (*api.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNewsServiceServer) GetFileLink(context.Context, *api.FileId) (*api.FileLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileLink not implemented")
}
func (UnimplementedNewsServiceServer) mustEmbedUnimplementedNewsServiceServer() {}

// UnsafeNewsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsServiceServer will
// result in compilation errors.
type UnsafeNewsServiceServer interface {
	mustEmbedUnimplementedNewsServiceServer()
}

func RegisterNewsServiceServer(s grpc.ServiceRegistrar, srv *ContentCheckServiceServer) {
	s.RegisterService(&NewsService_ServiceDesc, srv)
}

func _NewsService_GetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.NewsRequestParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_v1.NewsService/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetNews(ctx, req.(*api.NewsRequestParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ObjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_v1.NewsService/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetOne(ctx, req.(*api.ObjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetOneBySlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ObjectSlug)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetOneBySlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_v1.NewsService/GetOneBySlug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetOneBySlug(ctx, req.(*api.ObjectSlug))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.RequestNewsObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_v1.NewsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).Create(ctx, req.(*api.RequestNewsObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.RequestNewsObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_v1.NewsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).Update(ctx, req.(*api.RequestNewsObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ObjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_v1.NewsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).Delete(ctx, req.(*api.ObjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetFileLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.FileId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetFileLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_v1.NewsService/GetFileLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetFileLink(ctx, req.(*api.FileId))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsService_ServiceDesc is the grpc.ServiceDesc for NewsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content_v1.NewsService",
	HandlerType: (*NewsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNews",
			Handler:    _NewsService_GetNews_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _NewsService_GetOne_Handler,
		},
		{
			MethodName: "GetOneBySlug",
			Handler:    _NewsService_GetOneBySlug_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _NewsService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NewsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NewsService_Delete_Handler,
		},
		{
			MethodName: "GetFileLink",
			Handler:    _NewsService_GetFileLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news.proto",
}

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	// Получение тега
	Get(ctx context.Context, in *api.EmptyRequest, opts ...grpc.CallOption) (*api.TagList, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) Get(ctx context.Context, in *api.EmptyRequest, opts ...grpc.CallOption) (*api.TagList, error) {
	out := new(api.TagList)
	err := c.cc.Invoke(ctx, "/content_v1.TagService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	// Получение тега
	Get(context.Context, *api.EmptyRequest) (*api.TagList, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) Get(context.Context, *api.EmptyRequest) (*api.TagList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s grpc.ServiceRegistrar, srv TagServiceServer) {
	s.RegisterService(&TagService_ServiceDesc, srv)
}

func _TagService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_v1.TagService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).Get(ctx, req.(*api.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagService_ServiceDesc is the grpc.ServiceDesc for TagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content_v1.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TagService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news.proto",
}
