package api

import (
	"os"
	"testing"
	"time"

	db "github.com/StevenSopilidis/BackendMasterClass/db/sqlc"
	"github.com/StevenSopilidis/BackendMasterClass/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	t.Helper()
	config := util.Config{
		TokenSymmetricKey:     util.RandomString(32),
		ACCESS_TOKEN_DURATION: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

// entry point to our tests
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
