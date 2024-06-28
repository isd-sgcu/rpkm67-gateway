package auth

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	authProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	Validate()
	RefreshToken()
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

func (s *serviceImpl) Validate() {
}
func (s *serviceImpl) RefreshToken() {
}

func (s *serviceImpl) GetGoogleLoginUrl() (*dto.GetGoogleLoginUrlResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetGoogleLoginUrl(ctx, &authProto.GetGoogleLoginUrlRequest{})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
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
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.AlreadyExists:
			return nil, apperror.BadRequestError("User already exists")
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.VerifyGoogleLoginResponse{
		Credential: &dto.Credential{
			AccessToken:  res.Credential.AccessToken,
			RefreshToken: res.Credential.RefreshToken,
			ExpiresIn:    int(res.Credential.ExpiresIn),
		},
	}, nil
}

// func (s *serviceImpl) SignUp(req *dto.SignUpRequest) (*dto.SignupResponse, *apperror.AppError) {
// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// defer cancel()

// res, err := s.client.SignUp(ctx, &authProto.SignUpRequest{
// 	Email:     req.Email,
// 	Password:  req.Password,
// 	Firstname: req.Firstname,
// 	Lastname:  req.Lastname,
// })
// if err != nil {
// 	st, ok := status.FromError(err)
// 	if !ok {
// 		return nil, apperror.InternalServer
// 	}
// 	switch st.Code() {
// 	case codes.AlreadyExists:
// 		return nil, apperror.BadRequestError("User already exists")
// 	case codes.Internal:
// 		return nil, apperror.InternalServerError(err.Error())
// 	default:
// 		return nil, apperror.ServiceUnavailable
// 	}
// }

// return &dto.SignupResponse{
// 	Id:        res.Id,
// 	Email:     res.Email,
// 	Firstname: res.Firstname,
// 	Lastname:  res.Lastname,
// }, nil

// 	return nil, nil
// }

func (s *serviceImpl) SignIn(req *dto.SignInRequest) (*dto.Credential, *apperror.AppError) {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// res, err := s.client.SignIn(ctx, &authProto.SignInRequest{
	// 	Email:    req.Email,
	// 	Password: req.Password,
	// })
	// if err != nil {
	// 	st, ok := status.FromError(err)
	// 	if !ok {
	// 		return nil, apperror.InternalServer
	// 	}
	// 	switch st.Code() {
	// 	case codes.AlreadyExists:
	// 		return nil, apperror.BadRequestError("User already exists")
	// 	case codes.Internal:
	// 		return nil, apperror.InternalServerError(err.Error())
	// 	default:
	// 		return nil, apperror.ServiceUnavailable
	// 	}
	// }

	// return &dto.Credential{
	// 	AccessToken:  res.Credential.AccessToken,
	// 	RefreshToken: res.Credential.RefreshToken,
	// 	ExpiresIn:    int(res.Credential.ExpiresIn),
	// }, nil

	return nil, nil
}
