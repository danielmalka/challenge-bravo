package gin

import (
	"fmt"
	"net/http"

	"github.com/danielmalka/challenge-bravo/config"
	"github.com/danielmalka/challenge-bravo/pkg/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GinResponse struct {
	StatusCode int
	Message    gin.H
}

func getDB(c config.Config) (*gorm.DB, error, GinResponse) {
	errorMessage := GinResponse{
		StatusCode: http.StatusServiceUnavailable,
		Message:    gin.H{"error": "internal error"},
	}
	userPass := fmt.Sprintf("%s:%s", c.DBUser, c.DBPass)
	host := fmt.Sprintf("%s:%s", c.DBHost, c.DBPort)
	db, err := storage.ConnectMysql(userPass, c.DBSchema, host)
	return db, err, errorMessage
}
