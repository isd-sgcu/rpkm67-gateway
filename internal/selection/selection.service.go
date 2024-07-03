package selection

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	selectionProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/selection/v1"
	"go.uber.org/zap"
)

type Service interface {
	Create(req *dto.CreateSelectionRequest) (*dto.CreateSelectionResponse, *apperror.AppError)
	FindByGroupId(req *dto.FindByGroupIdSelectionRequest) (*dto.FindByGroupIdSelectionResponse, *apperror.AppError)
	Delete(req *dto.DeleteSelectionRequest) (*dto.DeleteSelectionResponse, *apperror.AppError)
}

type serviceImpl struct {
	client Client
	log    *zap.Logger
}

func NewService(client Client, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) Create(req *dto.CreateSelectionRequest) (*dto.CreateSelectionResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Create(ctx, &selectionProto.CreateSelectionRequest{
		GroupId: req.GroupId,
		BaanId:  req.BaanId,
	})
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.CreateSelectionResponse{
		Selection: ProtoToDto(res.Selection),
	}, nil
}

func (s *serviceImpl) FindByGroupId(req *dto.FindByGroupIdSelectionRequest) (*dto.FindByGroupIdSelectionResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByGroupId(ctx, &selectionProto.FindByGroupIdSelectionRequest{
		GroupId: req.GroupId,
	})
	if err != nil {
		s.log.Named("FindByGroupId").Error("FindByGroupId: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.FindByGroupIdSelectionResponse{
		Selections: ProtoToDtoList(res.Selections),
	}, nil
}

func (s *serviceImpl) Delete(req *dto.DeleteSelectionRequest) (*dto.DeleteSelectionResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Delete(ctx, &selectionProto.DeleteSelectionRequest{
		GroupId: req.Id,
	})
	if err != nil {
		s.log.Named("Delete").Error("Delete: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.DeleteSelectionResponse{
		Success: res.Success,
	}, nil
}
