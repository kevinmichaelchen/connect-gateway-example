// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: users/v1/users.proto

package usersv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/johanbrandhorst/connect-gateway-example/proto/users/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "users.v1.UserService"
)

// UserServiceClient is a client for the users.v1.UserService service.
type UserServiceClient interface {
	// Create a new user.
	AddUser(context.Context, *connect_go.Request[v1.AddUserRequest]) (*connect_go.Response[v1.User], error)
	// Get a user's details.
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.User], error)
	// List users.
	ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest]) (*connect_go.ServerStreamForClient[v1.User], error)
	// Update a user.
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.User], error)
}

// NewUserServiceClient constructs a client for the users.v1.UserService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		addUser: connect_go.NewClient[v1.AddUserRequest, v1.User](
			httpClient,
			baseURL+"/users.v1.UserService/AddUser",
			opts...,
		),
		getUser: connect_go.NewClient[v1.GetUserRequest, v1.User](
			httpClient,
			baseURL+"/users.v1.UserService/GetUser",
			opts...,
		),
		listUsers: connect_go.NewClient[v1.ListUsersRequest, v1.User](
			httpClient,
			baseURL+"/users.v1.UserService/ListUsers",
			opts...,
		),
		updateUser: connect_go.NewClient[v1.UpdateUserRequest, v1.User](
			httpClient,
			baseURL+"/users.v1.UserService/UpdateUser",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	addUser    *connect_go.Client[v1.AddUserRequest, v1.User]
	getUser    *connect_go.Client[v1.GetUserRequest, v1.User]
	listUsers  *connect_go.Client[v1.ListUsersRequest, v1.User]
	updateUser *connect_go.Client[v1.UpdateUserRequest, v1.User]
}

// AddUser calls users.v1.UserService.AddUser.
func (c *userServiceClient) AddUser(ctx context.Context, req *connect_go.Request[v1.AddUserRequest]) (*connect_go.Response[v1.User], error) {
	return c.addUser.CallUnary(ctx, req)
}

// GetUser calls users.v1.UserService.GetUser.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.User], error) {
	return c.getUser.CallUnary(ctx, req)
}

// ListUsers calls users.v1.UserService.ListUsers.
func (c *userServiceClient) ListUsers(ctx context.Context, req *connect_go.Request[v1.ListUsersRequest]) (*connect_go.ServerStreamForClient[v1.User], error) {
	return c.listUsers.CallServerStream(ctx, req)
}

// UpdateUser calls users.v1.UserService.UpdateUser.
func (c *userServiceClient) UpdateUser(ctx context.Context, req *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.User], error) {
	return c.updateUser.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the users.v1.UserService service.
type UserServiceHandler interface {
	// Create a new user.
	AddUser(context.Context, *connect_go.Request[v1.AddUserRequest]) (*connect_go.Response[v1.User], error)
	// Get a user's details.
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.User], error)
	// List users.
	ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest], *connect_go.ServerStream[v1.User]) error
	// Update a user.
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.User], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/users.v1.UserService/AddUser", connect_go.NewUnaryHandler(
		"/users.v1.UserService/AddUser",
		svc.AddUser,
		opts...,
	))
	mux.Handle("/users.v1.UserService/GetUser", connect_go.NewUnaryHandler(
		"/users.v1.UserService/GetUser",
		svc.GetUser,
		opts...,
	))
	mux.Handle("/users.v1.UserService/ListUsers", connect_go.NewServerStreamHandler(
		"/users.v1.UserService/ListUsers",
		svc.ListUsers,
		opts...,
	))
	mux.Handle("/users.v1.UserService/UpdateUser", connect_go.NewUnaryHandler(
		"/users.v1.UserService/UpdateUser",
		svc.UpdateUser,
		opts...,
	))
	return "/users.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) AddUser(context.Context, *connect_go.Request[v1.AddUserRequest]) (*connect_go.Response[v1.User], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("users.v1.UserService.AddUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.User], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("users.v1.UserService.GetUser is not implemented"))
}

func (UnimplementedUserServiceHandler) ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest], *connect_go.ServerStream[v1.User]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("users.v1.UserService.ListUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.User], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("users.v1.UserService.UpdateUser is not implemented"))
}
