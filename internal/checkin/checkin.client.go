package checkin

import (
	"context"

	checkinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"google.golang.org/grpc"
)

type clientImpl struct {
	client checkinProto.CheckInServiceClient
}

type Client interface {
	Create(ctx context.Context, in *checkinProto.CreateCheckInRequest, opts ...grpc.CallOption) (*checkinProto.CreateCheckInResponse, error)
	FindByUserId(ctx context.Context, in *checkinProto.FindByUserIdCheckInRequest, opts ...grpc.CallOption) (*checkinProto.FindByUserIdCheckInResponse, error)
	FindByEmail(ctx context.Context, in *checkinProto.FindByEmailCheckInRequest, opts ...grpc.CallOption) (*checkinProto.FindByEmailCheckInResponse, error)
}

func NewClient(client checkinProto.CheckInServiceClient) Client {
	return &clientImpl{
		client: client,
	}
}

func (c *clientImpl) Create(ctx context.Context, in *checkinProto.CreateCheckInRequest, opts ...grpc.CallOption) (*checkinProto.CreateCheckInResponse, error) {
	return c.client.Create(ctx, in, opts...)
}

func (c *clientImpl) FindByUserId(ctx context.Context, in *checkinProto.FindByUserIdCheckInRequest, opts ...grpc.CallOption) (*checkinProto.FindByUserIdCheckInResponse, error) {
	return c.client.FindByUserId(ctx, in, opts...)
}

func (c *clientImpl) FindByEmail(ctx context.Context, in *checkinProto.FindByEmailCheckInRequest, opts ...grpc.CallOption) (*checkinProto.FindByEmailCheckInResponse, error) {
	return c.client.FindByEmail(ctx, in, opts...)
}
