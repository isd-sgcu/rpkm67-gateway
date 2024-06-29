package user

import (
	"context"

	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	"google.golang.org/grpc"
)

type clientImpl struct {
	client userProto.UserServiceClient
}

type Client interface {
	Create(ctx context.Context, in *userProto.CreateUserRequest, opts ...grpc.CallOption) (*userProto.CreateUserResponse, error)
	FindOne(ctx context.Context, in *userProto.FindOneUserRequest, opts ...grpc.CallOption) (*userProto.FindOneUserResponse, error)
	Update(ctx context.Context, in *userProto.UpdateUserRequest, opts ...grpc.CallOption) (*userProto.UpdateUserResponse, error)
}

func NewClient(client userProto.UserServiceClient) Client {
	return &clientImpl{
		client: client,
	}
}

func (c *clientImpl) Create(ctx context.Context, in *userProto.CreateUserRequest, opts ...grpc.CallOption) (*userProto.CreateUserResponse, error) {
	return c.client.Create(ctx, in, opts...)
}

func (c *clientImpl) FindOne(ctx context.Context, in *userProto.FindOneUserRequest, opts ...grpc.CallOption) (*userProto.FindOneUserResponse, error) {
	return c.client.FindOne(ctx, in, opts...)
}

func (c *clientImpl) Update(ctx context.Context, in *userProto.UpdateUserRequest, opts ...grpc.CallOption) (*userProto.UpdateUserResponse, error) {
	return c.client.Update(ctx, in, opts...)
}
