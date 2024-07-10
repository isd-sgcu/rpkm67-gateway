package test

import (
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/selection"
	ctxMock "github.com/isd-sgcu/rpkm67-gateway/mocks/context"
	groupMock "github.com/isd-sgcu/rpkm67-gateway/mocks/group"
	selectionMock "github.com/isd-sgcu/rpkm67-gateway/mocks/selection"
	validatorMock "github.com/isd-sgcu/rpkm67-gateway/mocks/validator"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SelectionHandlerTest struct {
	suite.Suite
	controller                *gomock.Controller
	logger                    *zap.Logger
	userId                    string
	Selections                []*dto.Selection
	Selection                 *dto.Selection
	CreateSelectionReq        *dto.CreateSelectionRequest
	FindByGroupIdSelectionReq *dto.FindByGroupIdSelectionRequest
	DeleteSelectionReq        *dto.DeleteSelectionRequest
}

func TestSelectionHandler(t *testing.T) {
	suite.Run(t, new(SelectionHandlerTest))
}

func (t *SelectionHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	t.userId = faker.UUIDHyphenated()

	selectionsProto := MockSelectionsProto()
	selectionProto := selectionsProto[0]

	t.Selections = selection.ProtoToDtoList(selectionsProto)
	t.Selection = selection.ProtoToDto(selectionProto)

	t.CreateSelectionReq = &dto.CreateSelectionRequest{
		GroupId: t.Selection.GroupId,
		BaanId:  t.Selection.BaanId,
	}
	t.FindByGroupIdSelectionReq = &dto.FindByGroupIdSelectionRequest{
		GroupId: t.Selection.GroupId,
	}
	t.DeleteSelectionReq = &dto.DeleteSelectionRequest{
		Id: t.Selection.Id,
	}
}

func (t *SelectionHandlerTest) TestCreateSelectionSuccess() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	groupSvc := groupMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := selection.NewHandler(selectionSvc, groupSvc, validator, t.logger)

	expectedResp := &dto.CreateSelectionResponse{
		Selection: t.Selection,
	}

	context.EXPECT().GetString("userId").Return(t.userId)
	groupSvc.EXPECT().FindByUserId(&dto.FindByUserIdGroupRequest{UserId: t.userId}).
		Return(&dto.FindByUserIdGroupResponse{Group: &dto.Group{LeaderID: t.userId}}, nil)
	context.EXPECT().Next()
	context.EXPECT().Bind(&dto.CreateSelectionRequest{}).SetArg(0, *t.CreateSelectionReq)
	validator.EXPECT().Validate(t.CreateSelectionReq).Return(nil)
	selectionSvc.EXPECT().Create(t.CreateSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusCreated, expectedResp)

	handler.Create(context)
}

func (t *SelectionHandlerTest) TestCreateSelectionBindError() {
	context := ctxMock.NewMockCtx(t.controller)
	groupSvc := groupMock.NewMockService(t.controller)
	handler := selection.NewHandler(nil, groupSvc, nil, t.logger)

	context.EXPECT().GetString("userId").Return(t.userId)
	groupSvc.EXPECT().FindByUserId(&dto.FindByUserIdGroupRequest{UserId: t.userId}).
		Return(&dto.FindByUserIdGroupResponse{Group: &dto.Group{LeaderID: t.userId}}, nil)
	context.EXPECT().Next()
	context.EXPECT().Bind(&dto.CreateSelectionRequest{}).Return(apperror.BadRequest)
	context.EXPECT().BadRequestError(apperror.BadRequest.Error())

	handler.Create(context)
}

func (t *SelectionHandlerTest) TestCreateSelectionServiceError() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	groupSvc := groupMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := selection.NewHandler(selectionSvc, groupSvc, validator, t.logger)

	context.EXPECT().GetString("userId").Return(t.userId)
	groupSvc.EXPECT().FindByUserId(&dto.FindByUserIdGroupRequest{UserId: t.userId}).
		Return(&dto.FindByUserIdGroupResponse{Group: &dto.Group{LeaderID: t.userId}}, nil)
	context.EXPECT().Next()
	context.EXPECT().Bind(&dto.CreateSelectionRequest{}).SetArg(0, *t.CreateSelectionReq)
	validator.EXPECT().Validate(t.CreateSelectionReq).Return(nil)
	selectionSvc.EXPECT().Create(t.CreateSelectionReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.Create(context)
}

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionSuccess() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	groupSvc := groupMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := selection.NewHandler(selectionSvc, groupSvc, validator, t.logger)

	expectedResp := &dto.FindByGroupIdSelectionResponse{
		Selections: t.Selections,
	}

	context.EXPECT().GetString("userId").Return(t.userId)
	groupSvc.EXPECT().FindByUserId(&dto.FindByUserIdGroupRequest{UserId: t.userId}).
		Return(&dto.FindByUserIdGroupResponse{Group: &dto.Group{LeaderID: t.userId}}, nil)
	context.EXPECT().Next()
	context.EXPECT().Param("id").Return(t.Selection.GroupId)
	validator.EXPECT().Validate(t.FindByGroupIdSelectionReq).Return(nil)
	selectionSvc.EXPECT().FindByGroupId(t.FindByGroupIdSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.FindByGroupId(context)
}

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionUrlParamEmpty() {
	context := ctxMock.NewMockCtx(t.controller)
	groupSvc := groupMock.NewMockService(t.controller)
	handler := selection.NewHandler(nil, groupSvc, nil, t.logger)

	expectedErr := apperror.BadRequestError("url parameter 'id' not found")

	context.EXPECT().GetString("userId").Return(t.userId)
	groupSvc.EXPECT().FindByUserId(&dto.FindByUserIdGroupRequest{UserId: t.userId}).
		Return(&dto.FindByUserIdGroupResponse{Group: &dto.Group{LeaderID: t.userId}}, nil)
	context.EXPECT().Next()
	context.EXPECT().Param("id").Return("")
	context.EXPECT().BadRequestError(expectedErr.Error())

	handler.FindByGroupId(context)
}

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionServiceError() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	groupSvc := groupMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := selection.NewHandler(selectionSvc, groupSvc, validator, t.logger)

	context.EXPECT().GetString("userId").Return(t.userId)
	groupSvc.EXPECT().FindByUserId(&dto.FindByUserIdGroupRequest{UserId: t.userId}).
		Return(&dto.FindByUserIdGroupResponse{Group: &dto.Group{LeaderID: t.userId}}, nil)
	context.EXPECT().Next()
	context.EXPECT().Param("id").Return(t.Selection.GroupId)
	validator.EXPECT().Validate(t.FindByGroupIdSelectionReq).Return(nil)
	selectionSvc.EXPECT().FindByGroupId(t.FindByGroupIdSelectionReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.FindByGroupId(context)
}

func (t *SelectionHandlerTest) TestDeleteSelectionSuccess() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	groupSvc := groupMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := selection.NewHandler(selectionSvc, groupSvc, validator, t.logger)

	expectedResp := &dto.DeleteSelectionResponse{
		Success: true,
	}

	context.EXPECT().GetString("userId").Return(t.userId)
	groupSvc.EXPECT().FindByUserId(&dto.FindByUserIdGroupRequest{UserId: t.userId}).
		Return(&dto.FindByUserIdGroupResponse{Group: &dto.Group{LeaderID: t.userId}}, nil)
	context.EXPECT().Next()
	context.EXPECT().Bind(&dto.DeleteSelectionRequest{}).SetArg(0, *t.DeleteSelectionReq)
	validator.EXPECT().Validate(t.DeleteSelectionReq).Return(nil)
	selectionSvc.EXPECT().Delete(t.DeleteSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.Delete(context)
}

func (t *SelectionHandlerTest) TestDeleteSelectionBindError() {
	context := ctxMock.NewMockCtx(t.controller)
	groupSvc := groupMock.NewMockService(t.controller)
	handler := selection.NewHandler(nil, groupSvc, nil, t.logger)

	context.EXPECT().GetString("userId").Return(t.userId)
	groupSvc.EXPECT().FindByUserId(&dto.FindByUserIdGroupRequest{UserId: t.userId}).
		Return(&dto.FindByUserIdGroupResponse{Group: &dto.Group{LeaderID: t.userId}}, nil)
	context.EXPECT().Next()
	context.EXPECT().Bind(&dto.DeleteSelectionRequest{}).Return(apperror.BadRequest)
	context.EXPECT().BadRequestError(apperror.BadRequest.Error())

	handler.Delete(context)
}

func (t *SelectionHandlerTest) TestDeleteSelectionServiceError() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	groupSvc := groupMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := selection.NewHandler(selectionSvc, groupSvc, validator, t.logger)

	context.EXPECT().GetString("userId").Return(t.userId)
	groupSvc.EXPECT().FindByUserId(&dto.FindByUserIdGroupRequest{UserId: t.userId}).
		Return(&dto.FindByUserIdGroupResponse{Group: &dto.Group{LeaderID: t.userId}}, nil)
	context.EXPECT().Next()
	context.EXPECT().Bind(&dto.DeleteSelectionRequest{}).SetArg(0, *t.DeleteSelectionReq)
	validator.EXPECT().Validate(t.DeleteSelectionReq).Return(nil)
	selectionSvc.EXPECT().Delete(t.DeleteSelectionReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.Delete(context)
}
