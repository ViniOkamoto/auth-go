package models

type (
	AuthenticationRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	AuthenticationResponse struct {
		AccessToken string `json:"accessToken"`
		ExpiresIn   int    `json:"expiresIn"`
	}

	RefreshRequest struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}
)
