package selection

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	selectionProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/selection/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	CreateSelection(req *dto.CreateSelectionRequest) (*dto.CreateSelectionResponse, *apperror.AppError)
	FindByGroupIdSelection(req *dto.FindByGroupIdSelectionRequest) (*dto.FindByGroupIdSelectionResponse, *apperror.AppError)
	UpdateSelection(req *dto.UpdateSelectionRequest) (*dto.UpdateSelectionResponse, *apperror.AppError)
}

type serviceImpl struct {
	client selectionProto.SelectionServiceClient
	log    *zap.Logger
}

func NewService(client selectionProto.SelectionServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) CreateSelection(req *dto.CreateSelectionRequest) (*dto.CreateSelectionResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Create(ctx, &selectionProto.CreateSelectionRequest{
		GroupId: req.GroupId,
		BaanIds: req.BaanIds,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			s.log.Named("CreateSelection").Error("FromeError: ", zap.Error(err))
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			s.log.Named("CreateSelection").Error("Create: ", zap.Error(err))
			return nil, apperror.BadRequestError("Invalid argument")
		case codes.Internal:
			s.log.Named("CreateSelection").Error("Create: ", zap.Error(err))
			return nil, apperror.InternalServerError(err.Error())
		default:
			s.log.Named("CreateSelection").Error("Create: ", zap.Error(err))
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.CreateSelectionResponse{
		Selection: ProtoToDto(res.Selection),
	}, nil
}

func (s *serviceImpl) FindByGroupIdSelection(req *dto.FindByGroupIdSelectionRequest) (*dto.FindByGroupIdSelectionResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByGroupId(ctx, &selectionProto.FindByGroupIdSelectionRequest{
		GroupId: req.GroupId,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			s.log.Named("FindByGroupIdSelection").Error("FromeError: ", zap.Error(err))
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			s.log.Named("FindByGroupIdSelection").Error("FindByGroupId: ", zap.Error(err))
			return nil, apperror.BadRequestError("Invalid argument")
		case codes.Internal:
			s.log.Named("FindByGroupIdSelection").Error("FindByGroupId: ", zap.Error(err))
			return nil, apperror.InternalServerError(err.Error())
		default:
			s.log.Named("FindByGroupIdSelection").Error("FindByGroupId: ", zap.Error(err))
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.FindByGroupIdSelectionResponse{
		Selection: ProtoToDto(res.Selection),
	}, nil
}

func (s *serviceImpl) UpdateSelection(req *dto.UpdateSelectionRequest) (*dto.UpdateSelectionResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Update(ctx, &selectionProto.UpdateSelectionRequest{
		Selection: DtoToProto(req.Selection),
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			s.log.Named("UpdateSelection").Error("FromError: ", zap.Error(err))
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			s.log.Named("UpdateSelection").Error("Update: ", zap.Error(err))
			return nil, apperror.BadRequestError("Invalid argument")
		case codes.Internal:
			s.log.Named("UpdateSelection").Error("Update: ", zap.Error(err))
			return nil, apperror.InternalServerError(err.Error())
		default:
			s.log.Named("UpdateSelection").Error("Update: ", zap.Error(err))
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.UpdateSelectionResponse{
		Success: res.Success,
	}, nil
}
