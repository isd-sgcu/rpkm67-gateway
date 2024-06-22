package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperrors"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/utils"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type BaanHandlerTest struct {
	suite.Suite
	controller *gomock.Controller
	logger     *zap.Logger

	Baans          []*dto.Baan
	Baan           *dto.Baan
	FindAllBaanReq *dto.FindAllBaanRequest
	FindOneBaanReq *dto.FindOneBaanRequest
	Err            *apperrors.AppError
}

func TestBaanHandler(t *testing.T) {
	suite.Run(t, new(BaanHandlerTest))
}

func (t *BaanHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	baansProto := utils.MockBaansProto()
	baanProto := baansProto[0]

	t.Baans = utils.ProtoToDtoList(baansProto)
	t.Baan = utils.ProtoToDto(baanProto)

	t.FindAllBaanReq = &dto.FindAllBaanRequest{}
	t.FindOneBaanReq = &dto.FindOneBaanRequest{
		Id: t.Baan.Id,
	}
}

func (t *BaanHandlerTest) TestFindAllBaan() {
}

func (t *BaanHandlerTest) TestFindOneBaan() {

}

func (t *BaanHandlerTest) TearDownTest() {
	t.controller.Finish()
}
