// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: blog-service.proto

package pb

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

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	// Using the provided BlogService_NewBlogPost_Params, it creates a new blog post from the
	// argument. Take note that this does not totally persist the data to any persistence storage.
	// That is the job of the rpc SaveBlogPost!
	NewBlogPost(ctx context.Context, in *BlogService_NewBlogPost_Params, opts ...grpc.CallOption) (*Post, error)
	// Persist the provided post inthe BlogService_SaveBlogPost_params argument to any persistent
	// storage.
	SaveBlogPost(ctx context.Context, in *BlogService_SaveBlogPost_Params, opts ...grpc.CallOption) (*BlogService_SaveBlogPost_Response, error)
	// Destructive! Removes a blog post from any persistent storage in the internal service.
	DeleteBlogPost(ctx context.Context, in *BlogService_DeleteBlogPost_Params, opts ...grpc.CallOption) (*BlogService_DeleteBlogPost_Response, error)
	// Gets the specific blog post identified by the argument passed to the GetBlogPost from the persistent storage.
	GetBlogPost(ctx context.Context, in *BlogService_GetBlogPost_Params, opts ...grpc.CallOption) (*Post, error)
	// Creates a new comment for the target uuid provided as an argument. Like NewBlogPost this does
	// not totally save the new comment to any storage.
	NewComment(ctx context.Context, in *BlogService_NewComment_Params, opts ...grpc.CallOption) (*Comment, error)
	// Commits the comment passed as an argument to the BlogService_SaveComment_Params to any persistent storage.
	SaveComment(ctx context.Context, in *BlogService_SaveComment_Params, opts ...grpc.CallOption) (*BlogService_SaveComment_Response, error)
	// Destructive! Removes the comment from any persistence storage in thte internal service. (Sounds similar to the DeleteBlogPost!)
	DeleteComment(ctx context.Context, in *BlogService_DeleteComment_Params, opts ...grpc.CallOption) (*BlogService_DeleteComment_Response, error)
	GlobalSearch(ctx context.Context, in *BlogService_GlobalSearch_Params, opts ...grpc.CallOption) (*SearchResults, error)
	GlobalLatestBlogPosts(ctx context.Context, in *BlogService_GlobalLatestBlogPosts_Params, opts ...grpc.CallOption) (*BlogService_GlobalLatestBlogPosts_Response, error)
	Author_GetAuthorInfo(ctx context.Context, in *BlogService_AuthorGetInfo_Params, opts ...grpc.CallOption) (*User, error)
	Author_LatestBlogPosts(ctx context.Context, in *BlogService_AuthorLatestBlogPosts_Params, opts ...grpc.CallOption) (*BlogService_AuthorLatestBlogPosts_Response, error)
	Author_GetBlogPosts(ctx context.Context, in *BlogService_AuthorGetBlogPosts_Params, opts ...grpc.CallOption) (*BlogService_AuthorGetBlogPosts_Response, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) NewBlogPost(ctx context.Context, in *BlogService_NewBlogPost_Params, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/NewBlogPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) SaveBlogPost(ctx context.Context, in *BlogService_SaveBlogPost_Params, opts ...grpc.CallOption) (*BlogService_SaveBlogPost_Response, error) {
	out := new(BlogService_SaveBlogPost_Response)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/SaveBlogPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlogPost(ctx context.Context, in *BlogService_DeleteBlogPost_Params, opts ...grpc.CallOption) (*BlogService_DeleteBlogPost_Response, error) {
	out := new(BlogService_DeleteBlogPost_Response)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/DeleteBlogPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetBlogPost(ctx context.Context, in *BlogService_GetBlogPost_Params, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/GetBlogPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) NewComment(ctx context.Context, in *BlogService_NewComment_Params, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/NewComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) SaveComment(ctx context.Context, in *BlogService_SaveComment_Params, opts ...grpc.CallOption) (*BlogService_SaveComment_Response, error) {
	out := new(BlogService_SaveComment_Response)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/SaveComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteComment(ctx context.Context, in *BlogService_DeleteComment_Params, opts ...grpc.CallOption) (*BlogService_DeleteComment_Response, error) {
	out := new(BlogService_DeleteComment_Response)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GlobalSearch(ctx context.Context, in *BlogService_GlobalSearch_Params, opts ...grpc.CallOption) (*SearchResults, error) {
	out := new(SearchResults)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/GlobalSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GlobalLatestBlogPosts(ctx context.Context, in *BlogService_GlobalLatestBlogPosts_Params, opts ...grpc.CallOption) (*BlogService_GlobalLatestBlogPosts_Response, error) {
	out := new(BlogService_GlobalLatestBlogPosts_Response)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/GlobalLatestBlogPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Author_GetAuthorInfo(ctx context.Context, in *BlogService_AuthorGetInfo_Params, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/Author_GetAuthorInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Author_LatestBlogPosts(ctx context.Context, in *BlogService_AuthorLatestBlogPosts_Params, opts ...grpc.CallOption) (*BlogService_AuthorLatestBlogPosts_Response, error) {
	out := new(BlogService_AuthorLatestBlogPosts_Response)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/Author_LatestBlogPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Author_GetBlogPosts(ctx context.Context, in *BlogService_AuthorGetBlogPosts_Params, opts ...grpc.CallOption) (*BlogService_AuthorGetBlogPosts_Response, error) {
	out := new(BlogService_AuthorGetBlogPosts_Response)
	err := c.cc.Invoke(ctx, "/blogs.BlogService/Author_GetBlogPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	// Using the provided BlogService_NewBlogPost_Params, it creates a new blog post from the
	// argument. Take note that this does not totally persist the data to any persistence storage.
	// That is the job of the rpc SaveBlogPost!
	NewBlogPost(context.Context, *BlogService_NewBlogPost_Params) (*Post, error)
	// Persist the provided post inthe BlogService_SaveBlogPost_params argument to any persistent
	// storage.
	SaveBlogPost(context.Context, *BlogService_SaveBlogPost_Params) (*BlogService_SaveBlogPost_Response, error)
	// Destructive! Removes a blog post from any persistent storage in the internal service.
	DeleteBlogPost(context.Context, *BlogService_DeleteBlogPost_Params) (*BlogService_DeleteBlogPost_Response, error)
	// Gets the specific blog post identified by the argument passed to the GetBlogPost from the persistent storage.
	GetBlogPost(context.Context, *BlogService_GetBlogPost_Params) (*Post, error)
	// Creates a new comment for the target uuid provided as an argument. Like NewBlogPost this does
	// not totally save the new comment to any storage.
	NewComment(context.Context, *BlogService_NewComment_Params) (*Comment, error)
	// Commits the comment passed as an argument to the BlogService_SaveComment_Params to any persistent storage.
	SaveComment(context.Context, *BlogService_SaveComment_Params) (*BlogService_SaveComment_Response, error)
	// Destructive! Removes the comment from any persistence storage in thte internal service. (Sounds similar to the DeleteBlogPost!)
	DeleteComment(context.Context, *BlogService_DeleteComment_Params) (*BlogService_DeleteComment_Response, error)
	GlobalSearch(context.Context, *BlogService_GlobalSearch_Params) (*SearchResults, error)
	GlobalLatestBlogPosts(context.Context, *BlogService_GlobalLatestBlogPosts_Params) (*BlogService_GlobalLatestBlogPosts_Response, error)
	Author_GetAuthorInfo(context.Context, *BlogService_AuthorGetInfo_Params) (*User, error)
	Author_LatestBlogPosts(context.Context, *BlogService_AuthorLatestBlogPosts_Params) (*BlogService_AuthorLatestBlogPosts_Response, error)
	Author_GetBlogPosts(context.Context, *BlogService_AuthorGetBlogPosts_Params) (*BlogService_AuthorGetBlogPosts_Response, error)
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) NewBlogPost(context.Context, *BlogService_NewBlogPost_Params) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBlogPost not implemented")
}
func (UnimplementedBlogServiceServer) SaveBlogPost(context.Context, *BlogService_SaveBlogPost_Params) (*BlogService_SaveBlogPost_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBlogPost not implemented")
}
func (UnimplementedBlogServiceServer) DeleteBlogPost(context.Context, *BlogService_DeleteBlogPost_Params) (*BlogService_DeleteBlogPost_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlogPost not implemented")
}
func (UnimplementedBlogServiceServer) GetBlogPost(context.Context, *BlogService_GetBlogPost_Params) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlogPost not implemented")
}
func (UnimplementedBlogServiceServer) NewComment(context.Context, *BlogService_NewComment_Params) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewComment not implemented")
}
func (UnimplementedBlogServiceServer) SaveComment(context.Context, *BlogService_SaveComment_Params) (*BlogService_SaveComment_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveComment not implemented")
}
func (UnimplementedBlogServiceServer) DeleteComment(context.Context, *BlogService_DeleteComment_Params) (*BlogService_DeleteComment_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedBlogServiceServer) GlobalSearch(context.Context, *BlogService_GlobalSearch_Params) (*SearchResults, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalSearch not implemented")
}
func (UnimplementedBlogServiceServer) GlobalLatestBlogPosts(context.Context, *BlogService_GlobalLatestBlogPosts_Params) (*BlogService_GlobalLatestBlogPosts_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalLatestBlogPosts not implemented")
}
func (UnimplementedBlogServiceServer) Author_GetAuthorInfo(context.Context, *BlogService_AuthorGetInfo_Params) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Author_GetAuthorInfo not implemented")
}
func (UnimplementedBlogServiceServer) Author_LatestBlogPosts(context.Context, *BlogService_AuthorLatestBlogPosts_Params) (*BlogService_AuthorLatestBlogPosts_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Author_LatestBlogPosts not implemented")
}
func (UnimplementedBlogServiceServer) Author_GetBlogPosts(context.Context, *BlogService_AuthorGetBlogPosts_Params) (*BlogService_AuthorGetBlogPosts_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Author_GetBlogPosts not implemented")
}
func (UnimplementedBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_NewBlogPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_NewBlogPost_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).NewBlogPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/NewBlogPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).NewBlogPost(ctx, req.(*BlogService_NewBlogPost_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_SaveBlogPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_SaveBlogPost_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).SaveBlogPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/SaveBlogPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).SaveBlogPost(ctx, req.(*BlogService_SaveBlogPost_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlogPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_DeleteBlogPost_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlogPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/DeleteBlogPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlogPost(ctx, req.(*BlogService_DeleteBlogPost_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetBlogPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_GetBlogPost_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetBlogPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/GetBlogPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetBlogPost(ctx, req.(*BlogService_GetBlogPost_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_NewComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_NewComment_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).NewComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/NewComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).NewComment(ctx, req.(*BlogService_NewComment_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_SaveComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_SaveComment_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).SaveComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/SaveComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).SaveComment(ctx, req.(*BlogService_SaveComment_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_DeleteComment_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteComment(ctx, req.(*BlogService_DeleteComment_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GlobalSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_GlobalSearch_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GlobalSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/GlobalSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GlobalSearch(ctx, req.(*BlogService_GlobalSearch_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GlobalLatestBlogPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_GlobalLatestBlogPosts_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GlobalLatestBlogPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/GlobalLatestBlogPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GlobalLatestBlogPosts(ctx, req.(*BlogService_GlobalLatestBlogPosts_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Author_GetAuthorInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_AuthorGetInfo_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Author_GetAuthorInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/Author_GetAuthorInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Author_GetAuthorInfo(ctx, req.(*BlogService_AuthorGetInfo_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Author_LatestBlogPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_AuthorLatestBlogPosts_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Author_LatestBlogPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/Author_LatestBlogPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Author_LatestBlogPosts(ctx, req.(*BlogService_AuthorLatestBlogPosts_Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Author_GetBlogPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogService_AuthorGetBlogPosts_Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Author_GetBlogPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogs.BlogService/Author_GetBlogPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Author_GetBlogPosts(ctx, req.(*BlogService_AuthorGetBlogPosts_Params))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blogs.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewBlogPost",
			Handler:    _BlogService_NewBlogPost_Handler,
		},
		{
			MethodName: "SaveBlogPost",
			Handler:    _BlogService_SaveBlogPost_Handler,
		},
		{
			MethodName: "DeleteBlogPost",
			Handler:    _BlogService_DeleteBlogPost_Handler,
		},
		{
			MethodName: "GetBlogPost",
			Handler:    _BlogService_GetBlogPost_Handler,
		},
		{
			MethodName: "NewComment",
			Handler:    _BlogService_NewComment_Handler,
		},
		{
			MethodName: "SaveComment",
			Handler:    _BlogService_SaveComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _BlogService_DeleteComment_Handler,
		},
		{
			MethodName: "GlobalSearch",
			Handler:    _BlogService_GlobalSearch_Handler,
		},
		{
			MethodName: "GlobalLatestBlogPosts",
			Handler:    _BlogService_GlobalLatestBlogPosts_Handler,
		},
		{
			MethodName: "Author_GetAuthorInfo",
			Handler:    _BlogService_Author_GetAuthorInfo_Handler,
		},
		{
			MethodName: "Author_LatestBlogPosts",
			Handler:    _BlogService_Author_LatestBlogPosts_Handler,
		},
		{
			MethodName: "Author_GetBlogPosts",
			Handler:    _BlogService_Author_GetBlogPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog-service.proto",
}
