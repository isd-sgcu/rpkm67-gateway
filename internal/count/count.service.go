package count

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	countProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/count/v1"

	"go.uber.org/zap"
)

type Service interface {
	Create(req *dto.CreateCountRequest) (*dto.CreateCountResponse, *apperror.AppError)
}

type serviceImpl struct {
	client countProto.CountServiceClient
	log    *zap.Logger
}

func NewService(client countProto.CountServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) Create(req *dto.CreateCountRequest) (*dto.CreateCountResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Create(ctx, &countProto.CreateCountRequest{
		Name: req.Name,
	})
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.CreateCountResponse{
		Count: &dto.Count{
			ID:   res.Count.Id,
			Name: res.Count.Name,
		},
	}, nil
}
