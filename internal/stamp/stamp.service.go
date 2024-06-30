package stamp

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/pin"
	stampProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/stamp/v1"
	"go.uber.org/zap"
)

type Service interface {
	FindByUserId(req *dto.FindByUserIdStampRequest) (*dto.FindByUserIdStampResponse, *apperror.AppError)
	StampByUserId(req *dto.StampByUserIdRequest) (*dto.StampByUserIdResponse, *apperror.AppError)
}

type serviceImpl struct {
	client              stampProto.StampServiceClient
	pinSvc              pin.Service
	pinRequiredActivity map[string]struct{}
	logger              *zap.Logger
}

func NewService(client stampProto.StampServiceClient, pinSvc pin.Service, pinRequiredActivity map[string]struct{}, logger *zap.Logger) Service {
	return &serviceImpl{
		client:              client,
		pinSvc:              pinSvc,
		pinRequiredActivity: pinRequiredActivity,
		logger:              logger,
	}
}

func (s *serviceImpl) FindByUserId(req *dto.FindByUserIdStampRequest) (*dto.FindByUserIdStampResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByUserId(ctx, &stampProto.FindByUserIdStampRequest{
		UserId: req.UserID,
	})
	if err != nil {
		s.logger.Named("FindByUserId").Error("FindByUserId: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.FindByUserIdStampResponse{
		Stamp: ProtoToDto(res.Stamp),
	}, nil
}

func (s *serviceImpl) StampByUserId(req *dto.StampByUserIdRequest) (*dto.StampByUserIdResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, ok := s.pinRequiredActivity[req.ActivityId]; ok {
		checkRes, checkErr := s.pinSvc.CheckPin(&dto.CheckPinRequest{
			ActivityId: req.ActivityId,
			Code:       req.PinCode,
		})
		if checkErr != nil {
			s.logger.Named("StampByUserId").Error("CheckPin: ", zap.Error(checkErr))
			return nil, apperror.HandleServiceError(checkErr)
		}

		if !checkRes.IsMatch {
			s.logger.Named("StampByUserId").Error("invalid pin code")
			return nil, apperror.BadRequestError("invalid pin code")
		}
	}

	res, err := s.client.StampByUserId(ctx, &stampProto.StampByUserIdRequest{
		UserId:     req.UserID,
		ActivityId: req.ActivityId,
	})
	if err != nil {
		s.logger.Named("StampByUserId").Error("StampByUserId: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.StampByUserIdResponse{
		Stamp: ProtoToDto(res.Stamp),
	}, nil
}
