package auth

import (
	"Perry_the_Platypus/backend/features/auth/models"
	"Perry_the_Platypus/backend/shared/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
)

func AuthRoutes(r *gin.Engine, db *pgx.Conn) {
	authGroup := r.Group("/v1/auth")

	authGroup.POST("/signup", func(c *gin.Context) {
		var input UserInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		user, err := SignUp(c.Request.Context(), db, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Signup failed: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": sanitizeUser(user)})
	})

	authGroup.POST("/login", func(c *gin.Context) {
		var input LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		user, token, err := Login(c.Request.Context(), db, input)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token, "user": sanitizeUser(user)})
	})

	authGroup.GET("/logout", middleware.JWTMiddleware(os.Getenv("JWT_SECRET")), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Logout successful. Please remove the token from your client.",
		})
	})

	authGroup.GET("/test-user", middleware.JWTMiddleware(os.Getenv("JWT_SECRET")), func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		user, err := GetUserByID(c.Request.Context(), db, userID.(int))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Test query failed: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": sanitizeUser(user)})
	})
}

// sanitizeUser strips sensitive fields from the user before returning
func sanitizeUser(user *models.User) gin.H {
	return gin.H{
		"id":              user.ID,
		"username":        user.Username,
		"email":           user.Email,
		"privacy":         user.Privacy,
		"verified":        user.Verified,
		"follower_count":  user.FollowerCount,
		"following_count": user.FollowingCount,
		"created_at":      user.CreatedAt,
		"updated_at":      user.UpdatedAt,
	}
}
