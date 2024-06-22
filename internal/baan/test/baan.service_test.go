package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperrors"
	"github.com/isd-sgcu/rpkm67-gateway/internal/baan"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	baanMock "github.com/isd-sgcu/rpkm67-gateway/mocks/client/baan"
	baanProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/baan/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type BaanServiceTest struct {
	suite.Suite
	controller          *gomock.Controller
	logger              *zap.Logger
	BaansProto          []*baanProto.Baan
	BaanProto           *baanProto.Baan
	BaansDto            []*dto.Baan
	BaanDto             *dto.Baan
	FindAllBaanProtoReq *baanProto.FindAllBaanRequest
	FindAllBaanDtoReq   *dto.FindAllBaanRequest
	FindOneBaanProtoReq *baanProto.FindOneBaanRequest
	FindOneBaanDtoReq   *dto.FindOneBaanRequest
	Err                 apperrors.AppError
}

func TestBaanService(t *testing.T) {
	suite.Run(t, new(BaanServiceTest))
}

func (t *BaanServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	t.BaansProto = MockBaansProto()
	t.BaanProto = t.BaansProto[0]
	t.BaansDto = baan.ProtoToDtoList(t.BaansProto)
	t.BaanDto = baan.ProtoToDto(t.BaanProto)

	t.FindAllBaanProtoReq = &baanProto.FindAllBaanRequest{}
	t.FindOneBaanProtoReq = &baanProto.FindOneBaanRequest{
		Id: t.BaanProto.Id,
	}
	t.FindAllBaanDtoReq = &dto.FindAllBaanRequest{}
	t.FindOneBaanDtoReq = &dto.FindOneBaanRequest{
		Id: t.BaanDto.Id,
	}
}

func (t *BaanServiceTest) TestFindAllBaanSuccess() {
	protoResp := &baanProto.FindAllBaanResponse{
		Baans: t.BaansProto,
	}

	findAllBaansDto := baan.ProtoToDtoList(protoResp.Baans)

	expected := &dto.FindAllBaanResponse{
		Baans: findAllBaansDto,
	}

	client := baanMock.BaanClientMock{}
	client.On("FindAllBaan", t.FindAllBaanProtoReq).Return(protoResp, nil)

	svc := baan.NewService(&client, t.logger)
	actual, err := svc.FindAllBaan(t.FindAllBaanDtoReq)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), expected, actual)
}

func (t *BaanServiceTest) TestFindOneBaanSuccess() {
	protoResp := &baanProto.FindOneBaanResponse{
		Baan: t.BaanProto,
	}

	expected := &dto.FindOneBaanResponse{
		Baan: t.BaanDto,
	}

	client := baanMock.BaanClientMock{}
	client.On("FindOneBaan", t.FindOneBaanProtoReq).Return(protoResp, nil)

	svc := baan.NewService(&client, t.logger)
	actual, err := svc.FindOneBaan(t.FindOneBaanDtoReq)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), expected, actual)
}

func (t *BaanServiceTest) TearDownTest() {
	t.controller.Finish()
}
