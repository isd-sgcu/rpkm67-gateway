package store

import (
	"context"

	storeProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/store/object/v1"
	"google.golang.org/grpc"
)

type clientImpl struct {
	client storeProto.ObjectServiceClient
}

type Client interface {
	Upload(ctx context.Context, in *storeProto.UploadObjectRequest, opts ...grpc.CallOption) (*storeProto.UploadObjectResponse, error)
	FindByKey(ctx context.Context, in *storeProto.FindByKeyObjectRequest, opts ...grpc.CallOption) (*storeProto.FindByKeyObjectResponse, error)
	DeleteByKey(ctx context.Context, in *storeProto.DeleteByKeyObjectRequest, opts ...grpc.CallOption) (*storeProto.DeleteByKeyObjectResponse, error)
}

func NewClient(client storeProto.ObjectServiceClient) Client {
	return &clientImpl{
		client: client,
	}
}

func (c *clientImpl) Upload(ctx context.Context, in *storeProto.UploadObjectRequest, opts ...grpc.CallOption) (*storeProto.UploadObjectResponse, error) {
	return c.client.Upload(ctx, in, opts...)
}

func (c *clientImpl) FindByKey(ctx context.Context, in *storeProto.FindByKeyObjectRequest, opts ...grpc.CallOption) (*storeProto.FindByKeyObjectResponse, error) {
	return c.client.FindByKey(ctx, in, opts...)
}

func (c *clientImpl) DeleteByKey(ctx context.Context, in *storeProto.DeleteByKeyObjectRequest, opts ...grpc.CallOption) (*storeProto.DeleteByKeyObjectResponse, error) {
	return c.client.DeleteByKey(ctx, in, opts...)
}
