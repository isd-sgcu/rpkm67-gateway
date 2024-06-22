package selection

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperrors"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	selectionProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/selection/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	CreateSelection(req *dto.CreateSelectionRequest) (*dto.CreateSelectionResponse, *apperrors.AppError)
	FindByGroupIdSelection(req *dto.FindByGroupIdSelectionRequest) (*dto.FindByGroupIdSelectionResponse, *apperrors.AppError)
	UpdateSelection(req *dto.UpdateSelectionRequest) (*dto.UpdateSelectionResponse, *apperrors.AppError)
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

func (s *serviceImpl) CreateSelection(req *dto.CreateSelectionRequest) (*dto.CreateSelectionResponse, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Create(ctx, &selectionProto.CreateSelectionRequest{
		GroupId: req.GroupId,
		BaanIds: req.BaanIds,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperrors.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			return nil, apperrors.BadRequestError("Invalid argument")
		case codes.Internal:
			return nil, apperrors.InternalServerError(err.Error())
		default:
			return nil, apperrors.ServiceUnavailable
		}
	}

	return &dto.CreateSelectionResponse{
		Selection: ProtoToDto(res.Selection),
	}, nil
}

func (s *serviceImpl) FindByGroupIdSelection(req *dto.FindByGroupIdSelectionRequest) (*dto.FindByGroupIdSelectionResponse, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByGroupId(ctx, &selectionProto.FindByGroupIdSelectionRequest{
		UserId: req.UserId,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperrors.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			return nil, apperrors.BadRequestError("Invalid argument")
		case codes.Internal:
			return nil, apperrors.InternalServerError(err.Error())
		default:
			return nil, apperrors.ServiceUnavailable
		}
	}

	return &dto.FindByGroupIdSelectionResponse{
		Selection: ProtoToDto(res.Selection),
	}, nil
}

func (s *serviceImpl) UpdateSelection(req *dto.UpdateSelectionRequest) (*dto.UpdateSelectionResponse, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Update(ctx, &selectionProto.UpdateSelectionRequest{
		Selection: DtoToProto(req.Selection),
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperrors.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			return nil, apperrors.BadRequestError("Invalid argument")
		case codes.Internal:
			return nil, apperrors.InternalServerError(err.Error())
		default:
			return nil, apperrors.ServiceUnavailable
		}
	}

	return &dto.UpdateSelectionResponse{
		Success: res.Success,
	}, nil
}
