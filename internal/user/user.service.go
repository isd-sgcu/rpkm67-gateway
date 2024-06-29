package user

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	"go.uber.org/zap"
)

type Service interface {
	FindOne(req *dto.FindOneUserRequest) (*dto.FindOneUserResponse, *apperror.AppError)
	UpdateProfile(req *dto.UpdateUserProfileRequest) (*dto.UpdateUserProfileResponse, *apperror.AppError)
	UpdatePicture(req *dto.UpdateUserPictureRequest) (*dto.UpdateUserPictureResponse, *apperror.AppError)
}

type serviceImpl struct {
	client userProto.UserServiceClient
	// object svc
	log *zap.Logger
}

func NewService(client userProto.UserServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) FindOne(req *dto.FindOneUserRequest) (*dto.FindOneUserResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindOne(ctx, &userProto.FindOneUserRequest{
		Id: req.Id,
	})
	if err != nil {
		s.log.Named("FindOne").Error("FindOne: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.FindOneUserResponse{
		User: ProtoToDto(res.User),
	}, nil
}

func (s *serviceImpl) UpdateProfile(req *dto.UpdateUserProfileRequest) (*dto.UpdateUserProfileResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateReq := CreateUpdateUserRequestProto(req)

	res, err := s.client.Update(ctx, updateReq)
	if err != nil {
		s.log.Named("Update").Error("Update: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.UpdateUserProfileResponse{
		User: ProtoToDto(res.User),
	}, nil
}

func (s *serviceImpl) UpdatePicture(req *dto.UpdateUserPictureRequest) (*dto.UpdateUserPictureResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//object code

	updateReq := &userProto.UpdateUserRequest{
		Id: req.Id,
		// PhotoKey: req.PhotoKey, from object
	}

	res, err := s.client.Update(ctx, updateReq)
	if err != nil {
		s.log.Named("UpdatePicture").Error("Update: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.UpdateUserPictureResponse{
		User: ProtoToDto(res.User),
	}, nil
}
