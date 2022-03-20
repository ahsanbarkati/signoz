package auth

import "context"

type LoginRequest struct {
	UserID   string
	Password string
}

type LoginResponse struct {
	accessJwt  string
	refrestJwt string
}

type User struct {
	UserID   string
	Password string
	Groups   []Group
}

func Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {

	user, err := authenticateLogin(ctx, request)
	if err != nil {
		return nil, err
	}

	accessJwt, err := generateAccessJwt(user.UserID, user.Groups)
	if err != nil {
		return nil, err
	}

	refreshJwt, err := generateRefreshJwt(user.UserID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{accessJwt: accessJwt, refrestJwt: refreshJwt}, nil
}

func authenticateLogin(ctx context.Context, request *LoginRequest) (*User, error) {
	return &User{
		UserID:   request.UserID,
		Password: request.Password,
		Groups:   nil,
	}, nil
}
