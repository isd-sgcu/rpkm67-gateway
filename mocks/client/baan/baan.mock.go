package baan

import (
	"context"

	baanProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/baan/v1"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type BaanClientMock struct {
	mock.Mock
}

func (c *BaanClientMock) FindAllBaan(_ context.Context, req *baanProto.FindAllBaanRequest, _ ...grpc.CallOption) (res *baanProto.FindAllBaanResponse, err error) {
	args := c.Called(req)

	if args.Get(0) != nil {
		res = args.Get(0).(*baanProto.FindAllBaanResponse)
	}

	return res, args.Error(1)
}

func (c *BaanClientMock) FindOneBaan(_ context.Context, req *baanProto.FindOneBaanRequest, _ ...grpc.CallOption) (res *baanProto.FindOneBaanResponse, err error) {
	args := c.Called(req)

	if args.Get(0) != nil {
		res = args.Get(0).(*baanProto.FindOneBaanResponse)
	}

	return res, args.Error(1)
}
