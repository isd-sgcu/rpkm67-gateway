package checkin

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	checkinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	Create(req *dto.CreateCheckInRequest) (*dto.CreateCheckInResponse, *apperror.AppError)
	FindByUserID(req *dto.FindByUserIdCheckInRequest) (*dto.FindByUserIdCheckInResponse, *apperror.AppError)
	FindByEmail(req *dto.FindByEmailCheckInRequest) (*dto.FindByEmailCheckInResponse, *apperror.AppError)
}

type serviceImpl struct {
	client checkinProto.CheckInServiceClient
	log    *zap.Logger
}

func NewService(client checkinProto.CheckInServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) Create(req *dto.CreateCheckInRequest) (*dto.CreateCheckInResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Create(ctx, &checkinProto.CreateCheckInRequest{
		Email:  req.Email,
		UserId: req.UserID,
		Event:  req.Event,
	})
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			return nil, apperror.BadRequest
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.CreateCheckInResponse{
		CheckIn: &dto.CheckIn{
			ID:     res.CheckIn.Id,
			UserID: res.CheckIn.UserId,
			Email:  res.CheckIn.Email,
			Event:  res.CheckIn.Event,
		},
	}, nil
}

func (s *serviceImpl) FindByEmail(req *dto.FindByEmailCheckInRequest) (*dto.FindByEmailCheckInResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByEmail(ctx, &checkinProto.FindByEmailCheckInRequest{
		Email: req.Email,
	})
	if err != nil {
		s.log.Named("FindByEmail").Error("FindByEmail: ", zap.Error(err))
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			return nil, apperror.BadRequest
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.FindByEmailCheckInResponse{
		CheckIns: ProtoToDtos(res.CheckIns),
	}, nil
}

func (s *serviceImpl) FindByUserID(req *dto.FindByUserIdCheckInRequest) (*dto.FindByUserIdCheckInResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByUserId(ctx, &checkinProto.FindByUserIdCheckInRequest{
		UserId: req.UserID,
	})
	if err != nil {
		s.log.Named("FindByUserID").Error("FindByUserID: ", zap.Error(err))
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			return nil, apperror.BadRequest
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.FindByUserIdCheckInResponse{
		CheckIns: ProtoToDtos(res.CheckIns),
	}, nil
}
