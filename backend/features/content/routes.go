package content

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func ContentRoutes(r *gin.Engine, db *pgx.Conn) {
	postGroup := r.Group("/v1/content/posts")



	postGroup.POST("", func(c *gin.Context) {
	})
}