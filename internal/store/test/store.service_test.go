package test

import (
	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/store"
	objectMock "github.com/isd-sgcu/rpkm67-gateway/mocks/store"
	objectProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/store/object/v1"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type StoreServiceTest struct {
	suite.Suite
	controller                    *gomock.Controller
	logger                        *zap.Logger
	mockUrl                       string
	mockKey                       string
	mockData                      []byte
	objectProto                   *objectProto.Object
	objectDto                     *dto.Object
	createObjectProtoRequest      *objectProto.UploadObjectRequest
	createObjectDtoRequest        *dto.UploadObjectRequest
	findByKeyObjectProtoRequest   *objectProto.FindByKeyObjectRequest
	findByKeyObjectDtoRequest     *dto.FindByKeyObjectRequest
	deleteByKeyObjectProtoRequest *objectProto.DeleteByKeyObjectRequest
	deleteObjectDtoRequest        *dto.DeleteObjectRequest
}

func (t *StoreServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
	t.mockUrl = "mockUrl"
	t.mockKey = "mockKey"
	t.mockData = []byte("mockData")
	t.objectProto = &objectProto.Object{
		Url: t.mockUrl,
		Key: "key",
	}
	t.objectDto = &dto.Object{
		Url: t.mockUrl,
		Key: "key",
	}
	t.createObjectProtoRequest = &objectProto.UploadObjectRequest{
		Filename: "filename",
	}
	t.createObjectDtoRequest = &dto.UploadObjectRequest{
		Filename: "filename",
	}
	t.findByKeyObjectProtoRequest = &objectProto.FindByKeyObjectRequest{
		Key: "key",
	}
	t.findByKeyObjectDtoRequest = &dto.FindByKeyObjectRequest{
		Key: "key",
	}
	t.deleteByKeyObjectProtoRequest = &objectProto.DeleteByKeyObjectRequest{
		Key: "key",
	}
	t.deleteObjectDtoRequest = &dto.DeleteObjectRequest{
		Key: "key",
	}
}

func (t *StoreServiceTest) TestCreateObjectSuccess() {
	client := objectMock.NewMockClient(t.controller)
	svc := store.NewService(client, t.logger)

	protoResp := &objectProto.UploadObjectResponse{
		Object: t.objectProto,
	}

	expected := &dto.UploadObjectResponse{
		Object: t.objectDto,
	}

	client.EXPECT().Upload(gomock.Any(), t.createObjectProtoRequest).Return(protoResp, nil)
	actual, err := svc.Upload(t.createObjectDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *StoreServiceTest) TestCreateObjectInvalidArgument() {
	client := objectMock.NewMockClient(t.controller)
	svc := store.NewService(client, t.logger)

	clientErr := apperror.BadRequest

	client.EXPECT().Upload(gomock.Any(), t.createObjectProtoRequest).Return(nil, clientErr)
	actual, err := svc.Upload(t.createObjectDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.BadRequest, err)
}

func (t *StoreServiceTest) TestCreateObjectInternalError() {
	client := objectMock.NewMockClient(t.controller)
	svc := store.NewService(client, t.logger)

	clientErr := apperror.InternalServer

	client.EXPECT().Upload(gomock.Any(), t.createObjectProtoRequest).Return(nil, clientErr)
	actual, err := svc.Upload(t.createObjectDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.InternalServer, err)
}

func (t *StoreServiceTest) TestFindObjectByKeySuccess() {
	client := objectMock.NewMockClient(t.controller)
	svc := store.NewService(client, t.logger)

	protoResp := &objectProto.FindByKeyObjectResponse{
		Object: t.objectProto,
	}

	expected := &dto.FindByKeyObjectResponse{
		Object: t.objectDto,
	}

	client.EXPECT().FindByKey(gomock.Any(), t.findByKeyObjectProtoRequest).Return(protoResp, nil)
	actual, err := svc.FindByKey(t.findByKeyObjectDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *StoreServiceTest) TestFindObjectByKeyInvalidArgument() {
	client := objectMock.NewMockClient(t.controller)
	svc := store.NewService(client, t.logger)

	clientErr := apperror.BadRequest

	client.EXPECT().FindByKey(gomock.Any(), t.findByKeyObjectProtoRequest).Return(nil, clientErr)
	actual, err := svc.FindByKey(t.findByKeyObjectDtoRequest)

	t.NotNil(err)
	t.Nil(actual)
}

func (t *StoreServiceTest) TestFindObjectByKeyInternalError() {
	client := objectMock.NewMockClient(t.controller)
	svc := store.NewService(client, t.logger)

	clientErr := apperror.InternalServer

	client.EXPECT().FindByKey(gomock.Any(), t.findByKeyObjectProtoRequest).Return(nil, clientErr)
	actual, err := svc.FindByKey(t.findByKeyObjectDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.InternalServer, err)
}

func (t *StoreServiceTest) TestDeleteObjectByKeySuccess() {
	client := objectMock.NewMockClient(t.controller)
	svc := store.NewService(client, t.logger)

	protoResp := &objectProto.DeleteByKeyObjectResponse{
		Success: true,
	}

	expected := &dto.DeleteObjectResponse{
		Success: true,
	}

	client.EXPECT().DeleteByKey(gomock.Any(), t.deleteByKeyObjectProtoRequest).Return(protoResp, nil)
	actual, err := svc.DeleteByKey(t.deleteObjectDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *StoreServiceTest) TestDeleteObjectByKeyInvalidArgument() {
	client := objectMock.NewMockClient(t.controller)
	svc := store.NewService(client, t.logger)

	clientErr := apperror.BadRequest

	client.EXPECT().DeleteByKey(gomock.Any(), t.deleteByKeyObjectProtoRequest).Return(nil, clientErr)
	actual, err := svc.DeleteByKey(t.deleteObjectDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.BadRequest, err)
}

func (t *StoreServiceTest) TestDeleteObjectByKeyInternalError() {
	client := objectMock.NewMockClient(t.controller)
	svc := store.NewService(client, t.logger)

	clientErr := apperror.InternalServer

	client.EXPECT().DeleteByKey(gomock.Any(), t.deleteByKeyObjectProtoRequest).Return(nil, clientErr)
	actual, err := svc.DeleteByKey(t.deleteObjectDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.InternalServer, err)
}
