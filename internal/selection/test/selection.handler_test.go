package test

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/selection"
	routerMock "github.com/isd-sgcu/rpkm67-gateway/mocks/router"
	selectionMock "github.com/isd-sgcu/rpkm67-gateway/mocks/selection"
	validatorMock "github.com/isd-sgcu/rpkm67-gateway/mocks/validator"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SelectionHandlerTest struct {
	suite.Suite
	controller                *gomock.Controller
	logger                    *zap.Logger
	Selections                []*dto.Selection
	Selection                 *dto.Selection
	CreateSelectionReq        *dto.CreateSelectionRequest
	FindByGroupIdSelectionReq *dto.FindByGroupIdSelectionRequest
	UpdateSelectionReq        *dto.UpdateSelectionRequest
	Err                       *apperror.AppError
	QueriesMock               map[string]string
}

func TestSelectionHandler(t *testing.T) {
	suite.Run(t, new(SelectionHandlerTest))
}

func (t *SelectionHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	selectionsProto := MockSelectionsProto()
	selectionProto := selectionsProto[0]

	t.Selections = selection.ProtoToDtoList(selectionsProto)
	t.Selection = selection.ProtoToDto(selectionProto)

	t.CreateSelectionReq = &dto.CreateSelectionRequest{}
	t.FindByGroupIdSelectionReq = &dto.FindByGroupIdSelectionRequest{
		GroupId: t.Selection.GroupId,
	}
	t.UpdateSelectionReq = &dto.UpdateSelectionRequest{}
}

func (t *SelectionHandlerTest) TestCreateSelectionSuccess() {
	expectedResp := &dto.CreateSelectionResponse{
		Selection: t.Selection,
	}

	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)

	context.EXPECT().Bind(t.CreateSelectionReq).Return(nil)
	validator.EXPECT().Validate(t.CreateSelectionReq).Return(nil)
	selectionSvc.EXPECT().CreateSelection(t.CreateSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusCreated, expectedResp)

	handler := selection.NewHandler(selectionSvc, validator, t.logger)
	handler.CreateSelection(context)
}

func (t *SelectionHandlerTest) TestCreateSelectionInvalidArgument() {
	expectedErr := apperror.BadRequestError("invalid argument")

	context := routerMock.NewMockContext(t.controller)

	context.EXPECT().Bind(t.CreateSelectionReq).Return(expectedErr)
	context.EXPECT().BadRequestError(expectedErr.Error())

	handler := selection.NewHandler(nil, nil, t.logger)
	handler.CreateSelection(context)
}

func (t *SelectionHandlerTest) TestCreateSelectionInternalError() {
	expectedErr := apperror.InternalServerError("internal error")

	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)

	context.EXPECT().Bind(t.CreateSelectionReq).Return(nil)
	validator.EXPECT().Validate(t.CreateSelectionReq).Return(nil)

	selectionSvc.EXPECT().CreateSelection(t.CreateSelectionReq).Return(nil, expectedErr)
	context.EXPECT().ResponseError(expectedErr)

	handler := selection.NewHandler(selectionSvc, validator, t.logger)
	handler.CreateSelection(context)
}

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionSuccess() {
	expectedResp := &dto.FindByGroupIdSelectionResponse{
		Selection: t.Selection,
	}

	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)

	context.EXPECT().Param("id").Return(t.Selection.GroupId)
	validator.EXPECT().Validate(t.FindByGroupIdSelectionReq).Return(nil)
	selectionSvc.EXPECT().FindByGroupIdSelection(t.FindByGroupIdSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler := selection.NewHandler(selectionSvc, validator, t.logger)
	handler.FindByGroupIdSelection(context)

}

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionInvalidArgument() {
	expectedErr := apperror.BadRequestError("url parameter 'id' not found")

	context := routerMock.NewMockContext(t.controller)

	context.EXPECT().Param("id").Return("")
	context.EXPECT().BadRequestError(expectedErr.Error())

	handler := selection.NewHandler(nil, nil, t.logger)
	handler.FindByGroupIdSelection(context)
}

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionInternalError() {
	expectedErr := apperror.InternalServerError("internal error")

	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)

	context.EXPECT().Param("id").Return(t.Selection.GroupId)
	validator.EXPECT().Validate(t.FindByGroupIdSelectionReq).Return(nil)

	selectionSvc.EXPECT().FindByGroupIdSelection(t.FindByGroupIdSelectionReq).Return(nil, expectedErr)
	context.EXPECT().ResponseError(expectedErr)

	handler := selection.NewHandler(selectionSvc, validator, t.logger)
	handler.FindByGroupIdSelection(context)
}

func (t *SelectionHandlerTest) TestUpdateSelectionSuccess() {
	expectedResp := &dto.UpdateSelectionResponse{
		Success: true,
	}

	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)

	context.EXPECT().Bind(t.UpdateSelectionReq).Return(nil)
	validator.EXPECT().Validate(t.UpdateSelectionReq).Return(nil)
	selectionSvc.EXPECT().UpdateSelection(t.UpdateSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler := selection.NewHandler(selectionSvc, validator, t.logger)
	handler.UpdateSelection(context)
}

func (t *SelectionHandlerTest) TestUpdateSelectionInvalidArgument() {
	expectedErr := apperror.BadRequestError("invalid argument")

	context := routerMock.NewMockContext(t.controller)

	context.EXPECT().Bind(t.UpdateSelectionReq).Return(expectedErr)
	context.EXPECT().BadRequestError(expectedErr.Error())

	handler := selection.NewHandler(nil, nil, t.logger)
	handler.UpdateSelection(context)
}

func (t *SelectionHandlerTest) TestUpdateSelectionInternalError() {
	expectedErr := apperror.InternalServerError("internal error")

	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)

	context.EXPECT().Bind(t.UpdateSelectionReq).Return(nil)
	validator.EXPECT().Validate(t.UpdateSelectionReq).Return(nil)

	selectionSvc.EXPECT().UpdateSelection(t.UpdateSelectionReq).Return(nil, expectedErr)
	context.EXPECT().ResponseError(expectedErr)

	handler := selection.NewHandler(selectionSvc, validator, t.logger)
	handler.UpdateSelection(context)
}
