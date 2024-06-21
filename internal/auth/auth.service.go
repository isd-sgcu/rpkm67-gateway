package auth

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperrors"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	authProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	client authProto.AuthServiceClient
}

type serviceImpl struct {
	client authProto.AuthServiceClient
	log    *zap.Logger
}

func NewService(client authProto.AuthServiceClient, log *zap.Logger) *Service {
	return &serviceImpl{
		repo: repo,
		log:  log,
	}
}

func (s *serviceImpl) Signin(req *dto.SignInRequest) (*dto.Credential, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.SignIn(ctx, &authProto.SignInRequest{
		StudentId: req.StudentId,
		Password:  req.Password,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperrors.InternalServer
		}
		switch st.Code() {
		case codes.AlreadyExists:
			return nil, apperrors.BadRequestError("User already exists")
		case codes.Internal:
			return nil, apperrors.InternalServer
		default:
			return nil, apperrors.ServiceUnavailable
		}
	}

	return &dto.Credential{
		AccessToken:  res.Credential.AccessToken,
		RefreshToken: res.Credential.RefreshToken,
		ExpiresIn:    int(res.Credential.ExpiresIn),
	}, nil
}
