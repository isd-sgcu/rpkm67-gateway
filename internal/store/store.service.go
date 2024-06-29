package store

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	storeProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/store/object/v1"
	"go.uber.org/zap"
)

type Service interface {
	Upload(req *dto.UploadObjectRequest) (*dto.UploadObjectResponse, *apperror.AppError)
	FindByKey(req *dto.FindByKeyObjectRequest) (*dto.FindByKeyObjectResponse, *apperror.AppError)
	DeleteByKey(req *dto.DeleteObjectRequest) (*dto.DeleteObjectResponse, *apperror.AppError)
}

type serviceImpl struct {
	client storeProto.ObjectServiceClient
	log    *zap.Logger
}

func NewService(client storeProto.ObjectServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) Upload(req *dto.UploadObjectRequest) (*dto.UploadObjectResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Upload(ctx, &storeProto.UploadObjectRequest{
		Filename: req.Filename,
	})
	if err != nil {
		s.log.Named("Upload").Error("Upload: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.UploadObjectResponse{
		Object: &dto.Object{
			Url: res.Object.Url,
			Key: res.Object.Key,
		},
	}, nil
}

func (s *serviceImpl) FindByKey(req *dto.FindByKeyObjectRequest) (*dto.FindByKeyObjectResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByKey(ctx, &storeProto.FindByKeyObjectRequest{
		Key: req.Key,
	})
	if err != nil {
		s.log.Named("FindByKey").Error("FindByKey: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.FindByKeyObjectResponse{
		Object: &dto.Object{
			Url: res.Object.Url,
			Key: res.Object.Key,
		},
	}, nil
}

func (s *serviceImpl) DeleteByKey(req *dto.DeleteObjectRequest) (*dto.DeleteObjectResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.DeleteByKey(ctx, &storeProto.DeleteByKeyObjectRequest{
		Key: req.Key,
	})
	if err != nil {
		s.log.Named("DeleteByKey").Error("DeleteByKey: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.DeleteObjectResponse{
		Success: res.Success,
	}, nil
}
