package selection

import (
	"context"

	selectionProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/selection/v1"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type SelectionClientMock struct {
	mock.Mock
}

func (c *SelectionClientMock) Create(_ context.Context, req *selectionProto.CreateSelectionRequest, _ ...grpc.CallOption) (res *selectionProto.CreateSelectionResponse, err error) {
	args := c.Called(req)

	if args.Get(0) != nil {
		res = args.Get(0).(*selectionProto.CreateSelectionResponse)
	}

	return res, args.Error(1)
}

func (c *SelectionClientMock) FindByGroupId(_ context.Context, req *selectionProto.FindByGroupIdSelectionRequest, _ ...grpc.CallOption) (res *selectionProto.FindByGroupIdSelectionResponse, err error) {
	args := c.Called(req)

	if args.Get(0) != nil {
		res = args.Get(0).(*selectionProto.FindByGroupIdSelectionResponse)
	}

	return res, args.Error(1)
}

func (c *SelectionClientMock) Delete(_ context.Context, req *selectionProto.DeleteSelectionRequest, _ ...grpc.CallOption) (res *selectionProto.DeleteSelectionResponse, err error) {
	args := c.Called(req)

	if args.Get(0) != nil {
		res = args.Get(0).(*selectionProto.DeleteSelectionResponse)
	}

	return res, args.Error(1)
}
