package selection

import (
	"context"

	selProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/selection/v1"
	"google.golang.org/grpc"
)

type Client interface {
	Create(ctx context.Context, in *selProto.CreateSelectionRequest, opts ...grpc.CallOption) (*selProto.CreateSelectionResponse, error)
	FindByGroupId(ctx context.Context, in *selProto.FindByGroupIdSelectionRequest, opts ...grpc.CallOption) (*selProto.FindByGroupIdSelectionResponse, error)
	Delete(ctx context.Context, in *selProto.DeleteSelectionRequest, opts ...grpc.CallOption) (*selProto.DeleteSelectionResponse, error)
	CountByBaanId(ctx context.Context, in *selProto.CountByBaanIdSelectionRequest, opts ...grpc.CallOption) (*selProto.CountByBaanIdSelectionResponse, error)
}

type clientImpl struct {
	client selProto.SelectionServiceClient
}

func NewClient(client selProto.SelectionServiceClient) Client {
	return &clientImpl{
		client: client,
	}
}

func (c *clientImpl) Create(ctx context.Context, in *selProto.CreateSelectionRequest, opts ...grpc.CallOption) (*selProto.CreateSelectionResponse, error) {
	return c.client.Create(ctx, in, opts...)
}

func (c *clientImpl) FindByGroupId(ctx context.Context, in *selProto.FindByGroupIdSelectionRequest, opts ...grpc.CallOption) (*selProto.FindByGroupIdSelectionResponse, error) {
	return c.client.FindByGroupId(ctx, in, opts...)
}

func (c *clientImpl) Delete(ctx context.Context, in *selProto.DeleteSelectionRequest, opts ...grpc.CallOption) (*selProto.DeleteSelectionResponse, error) {
	return c.client.Delete(ctx, in, opts...)
}

func (c *clientImpl) CountByBaanId(ctx context.Context, in *selProto.CountByBaanIdSelectionRequest, opts ...grpc.CallOption) (*selProto.CountByBaanIdSelectionResponse, error) {
	return c.client.CountByBaanId(ctx, in, opts...)
}
