package dto

type Credential struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL3BiZX..."`
	RefreshToken string `json:"refresh_token" example:"e7e84d54-7518-4..."`
	ExpiresIn    int    `json:"expires_in" example:"3600"`
}

type ValidateRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type ValidateResponse struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type GetGoogleLoginUrlResponse struct {
	Url string `json:"url"`
}

type VerifyGoogleLoginRequest struct {
	Code string `json:"code" validate:"required"`
}

type VerifyGoogleLoginResponse struct {
	Credential *Credential `json:"credential"`
	UserId     string      `json:"user_id"`
}
