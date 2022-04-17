package root

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"blog-gin_golang_v177_mongo/lib/response"
)

func Index(context *gin.Context) {
	response.Json(context, http.StatusOK, nil)
}
