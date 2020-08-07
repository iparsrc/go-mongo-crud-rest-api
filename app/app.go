package app

import (
	"github.com/gin-gonic/gin"
	"github.com/parsaakbari1209/go-mongo-CRUD-web/domain"
)

var (
	router = gin.Default()
)

func StartApp() {
	MapUrls()
	domain.ConnDB()
	router.Run(":8080")
}
