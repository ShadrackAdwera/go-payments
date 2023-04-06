package api

import (
	"os"
	"testing"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/gin-gonic/gin"
)

func newServer(store db.TxStore) *Server {
	srv := NewServer(store, nil, nil)
	return srv
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
