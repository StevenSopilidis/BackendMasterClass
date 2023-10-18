package gapi

import (
	"context"
	"database/sql"

	db "github.com/StevenSopilidis/BackendMasterClass/db/sqlc"
	"github.com/StevenSopilidis/BackendMasterClass/pb"
	"github.com/StevenSopilidis/BackendMasterClass/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "User not Found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to find user")
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Incorrect password")
	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(req.Username, server.config.ACCESS_TOKEN_DURATION)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create access token")
	}

	refreshToken, refreshTokenPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.RefreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create refresh token")
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshTokenPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create session")
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  timestamppb.New(accessTokenPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshTokenPayload.ExpiredAt),
		SessionId:             session.ID.String(),
	}

	return rsp, nil
}
