package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinResponse struct {
	StatusCode int
	Message    gin.H
}

func NewErrorGinResponse() GinResponse {
	return GinResponse{
		StatusCode: http.StatusServiceUnavailable,
		Message:    gin.H{"error": "internal error"},
	}
}
