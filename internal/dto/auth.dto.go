package dto

type Credential struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL3BiZX..."`
	RefreshToken string `json:"refresh_token" example:"e7e84d54-7518-4..."`
	ExpiresIn    int    `json:"expires_in" example:"3600"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,gte=6,lte=30"`
}
type SignUpRequest struct {
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required,gte=6,lte=30"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
}

type TokenPayloadAuth struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type VerifyGoogleLoginRequest struct {
	Code string `json:"code" validate:"required"`
}

type VerifyGoogleLoginResponse struct {
	Credential *Credential `json:"credential"`
}
