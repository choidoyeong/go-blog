// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package blog

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

// BlogClient is the client API for Blog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogClient interface {
	GetPost(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Post, error)
	CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	ModifyPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	ListPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (Blog_ListPostsClient, error)
	CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error)
	ListComments(ctx context.Context, in *PostId, opts ...grpc.CallOption) (Blog_ListCommentsClient, error)
}

type blogClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogClient(cc grpc.ClientConnInterface) BlogClient {
	return &blogClient{cc}
}

func (c *blogClient) GetPost(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.Blog/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.Blog/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) ModifyPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.Blog/ModifyPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) ListPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (Blog_ListPostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Blog_ServiceDesc.Streams[0], "/blog.Blog/ListPosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogListPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Blog_ListPostsClient interface {
	Recv() (*Post, error)
	grpc.ClientStream
}

type blogListPostsClient struct {
	grpc.ClientStream
}

func (x *blogListPostsClient) Recv() (*Post, error) {
	m := new(Post)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blogClient) CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, "/blog.Blog/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) ListComments(ctx context.Context, in *PostId, opts ...grpc.CallOption) (Blog_ListCommentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Blog_ServiceDesc.Streams[1], "/blog.Blog/ListComments", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogListCommentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Blog_ListCommentsClient interface {
	Recv() (*Comment, error)
	grpc.ClientStream
}

type blogListCommentsClient struct {
	grpc.ClientStream
}

func (x *blogListCommentsClient) Recv() (*Comment, error) {
	m := new(Comment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServer is the server API for Blog service.
// All implementations must embed UnimplementedBlogServer
// for forward compatibility
type BlogServer interface {
	GetPost(context.Context, *PostId) (*Post, error)
	CreatePost(context.Context, *Post) (*Post, error)
	ModifyPost(context.Context, *Post) (*Post, error)
	ListPosts(*EmptyRequest, Blog_ListPostsServer) error
	CreateComment(context.Context, *Comment) (*Comment, error)
	ListComments(*PostId, Blog_ListCommentsServer) error
	mustEmbedUnimplementedBlogServer()
}

// UnimplementedBlogServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServer struct {
}

func (UnimplementedBlogServer) GetPost(context.Context, *PostId) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedBlogServer) CreatePost(context.Context, *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedBlogServer) ModifyPost(context.Context, *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyPost not implemented")
}
func (UnimplementedBlogServer) ListPosts(*EmptyRequest, Blog_ListPostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (UnimplementedBlogServer) CreateComment(context.Context, *Comment) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedBlogServer) ListComments(*PostId, Blog_ListCommentsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListComments not implemented")
}
func (UnimplementedBlogServer) mustEmbedUnimplementedBlogServer() {}

// UnsafeBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServer will
// result in compilation errors.
type UnsafeBlogServer interface {
	mustEmbedUnimplementedBlogServer()
}

func RegisterBlogServer(s grpc.ServiceRegistrar, srv BlogServer) {
	s.RegisterService(&Blog_ServiceDesc, srv)
}

func _Blog_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).GetPost(ctx, req.(*PostId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).CreatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_ModifyPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).ModifyPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/ModifyPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).ModifyPost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_ListPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServer).ListPosts(m, &blogListPostsServer{stream})
}

type Blog_ListPostsServer interface {
	Send(*Post) error
	grpc.ServerStream
}

type blogListPostsServer struct {
	grpc.ServerStream
}

func (x *blogListPostsServer) Send(m *Post) error {
	return x.ServerStream.SendMsg(m)
}

func _Blog_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).CreateComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_ListComments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PostId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServer).ListComments(m, &blogListCommentsServer{stream})
}

type Blog_ListCommentsServer interface {
	Send(*Comment) error
	grpc.ServerStream
}

type blogListCommentsServer struct {
	grpc.ServerStream
}

func (x *blogListCommentsServer) Send(m *Comment) error {
	return x.ServerStream.SendMsg(m)
}

// Blog_ServiceDesc is the grpc.ServiceDesc for Blog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.Blog",
	HandlerType: (*BlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPost",
			Handler:    _Blog_GetPost_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _Blog_CreatePost_Handler,
		},
		{
			MethodName: "ModifyPost",
			Handler:    _Blog_ModifyPost_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _Blog_CreateComment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPosts",
			Handler:       _Blog_ListPosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListComments",
			Handler:       _Blog_ListComments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blog/blog.proto",
}
