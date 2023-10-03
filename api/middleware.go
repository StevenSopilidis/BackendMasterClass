package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/StevenSopilidis/BackendMasterClass/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("Authorization header not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse((err)))
			return
		}

		fields := strings.Fields(authorizationHeader)
		// Bearer token
		if len(fields) < 2 {
			err := errors.New("Invalid format of authorization header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse((err)))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("Authorization type not supported")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse((err)))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse((err)))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
