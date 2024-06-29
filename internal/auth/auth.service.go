package auth

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	authProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
	"go.uber.org/zap"
)

type Service interface {
	Validate(req *dto.ValidateRequest) (*dto.ValidateResponse, *apperror.AppError)
	RefreshToken(req *dto.RefreshTokenRequest) (*dto.Credential, *apperror.AppError)
	GetGoogleLoginUrl() (*dto.GetGoogleLoginUrlResponse, *apperror.AppError)
	VerifyGoogleLogin(req *dto.VerifyGoogleLoginRequest) (*dto.VerifyGoogleLoginResponse, *apperror.AppError)
}

type serviceImpl struct {
	client authProto.AuthServiceClient
	log    *zap.Logger
}

func NewService(client authProto.AuthServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) Validate(req *dto.ValidateRequest) (*dto.ValidateResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Validate(ctx, &authProto.ValidateRequest{AccessToken: req.AccessToken})
	if err != nil {
		s.log.Named("Validate").Error("Validate: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.ValidateResponse{
		UserId: res.UserId,
		Role:   res.Role,
	}, nil
}
func (s *serviceImpl) RefreshToken(req *dto.RefreshTokenRequest) (*dto.Credential, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.RefreshToken(ctx, &authProto.RefreshTokenRequest{RefreshToken: req.RefreshToken})
	if err != nil {
		s.log.Named("RefreshToken").Error("RefreshToken: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.Credential{
		AccessToken:  res.Credential.AccessToken,
		RefreshToken: res.Credential.RefreshToken,
		ExpiresIn:    int(res.Credential.ExpiresIn),
	}, nil
}

func (s *serviceImpl) GetGoogleLoginUrl() (*dto.GetGoogleLoginUrlResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetGoogleLoginUrl(ctx, &authProto.GetGoogleLoginUrlRequest{})
	if err != nil {
		s.log.Named("GetGoogleLoginUrl").Error("GetGoogleLoginUrl: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.GetGoogleLoginUrlResponse{
		Url: res.Url,
	}, nil
}

func (s *serviceImpl) VerifyGoogleLogin(req *dto.VerifyGoogleLoginRequest) (*dto.VerifyGoogleLoginResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.VerifyGoogleLogin(ctx, &authProto.VerifyGoogleLoginRequest{
		Code: req.Code,
	})
	if err != nil {
		s.log.Named("VerifyGoogleLogin").Error("VerifyGoogleLogin: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.VerifyGoogleLoginResponse{
		Credential: &dto.Credential{
			AccessToken:  res.Credential.AccessToken,
			RefreshToken: res.Credential.RefreshToken,
			ExpiresIn:    int(res.Credential.ExpiresIn),
		},
	}, nil
}
