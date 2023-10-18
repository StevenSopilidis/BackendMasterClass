package gapi

import (
	"fmt"

	db "github.com/StevenSopilidis/BackendMasterClass/db/sqlc"
	"github.com/StevenSopilidis/BackendMasterClass/pb"
	"github.com/StevenSopilidis/BackendMasterClass/token"
	"github.com/StevenSopilidis/BackendMasterClass/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer // for forwards compatibility
	config                           util.Config
	store                            db.Store
	tokenMaker                       token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	maker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("connot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: maker,
	}

	return server, nil
}
