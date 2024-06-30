package stamp

import (
	"context"

	stampProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/stamp/v1"
	"google.golang.org/grpc"
)

type clientImpl struct {
	client stampProto.StampServiceClient
}

type Client interface {
	FindByUserId(ctx context.Context, in *stampProto.FindByUserIdStampRequest, opts ...grpc.CallOption) (*stampProto.FindByUserIdStampResponse, error)
	StampByUserId(ctx context.Context, in *stampProto.StampByUserIdRequest, opts ...grpc.CallOption) (*stampProto.StampByUserIdResponse, error)
}

func NewClient(client stampProto.StampServiceClient) Client {
	return &clientImpl{
		client: client,
	}
}

func (c *clientImpl) FindByUserId(ctx context.Context, in *stampProto.FindByUserIdStampRequest, opts ...grpc.CallOption) (*stampProto.FindByUserIdStampResponse, error) {
	return c.client.FindByUserId(ctx, in, opts...)
}

func (c *clientImpl) StampByUserId(ctx context.Context, in *stampProto.StampByUserIdRequest, opts ...grpc.CallOption) (*stampProto.StampByUserIdResponse, error) {
	return c.client.StampByUserId(ctx, in, opts...)
}
