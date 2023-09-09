package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// entry point to our tests
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
