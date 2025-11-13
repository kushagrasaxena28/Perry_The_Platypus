package channels

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func ChannelRoutes(r *gin.Engine, db *pgx.Conn) {
	ChannelGroup := r.Group("/v1/channels")

	//Core Channel Routes
	ChannelGroup.GET("", func(c *gin.Context) {})
	ChannelGroup.POST("", func(c *gin.Context) {})
	ChannelGroup.GET("/:id", func(c *gin.Context) {})
	ChannelGroup.PUT("/:id", func(c *gin.Context) {})
	ChannelGroup.DELETE("/:id", func(c *gin.Context) {})

	// Channel Membership Routes
	ChannelGroup.POST("/:id/subscribe", func(c *gin.Context) {})
	ChannelGroup.DELETE("/:id/unsubscribe", func(c *gin.Context) {})
	ChannelGroup.GET("/:id/subscribers", func(c *gin.Context) {})

	/*Community Routes
	channelGroup.GET("/:id/communities", ListCommunities)
	channelGroup.POST("/:id/communities", AuthMiddleware, CreateCommunity)
	*/
	/*Posts , Reels , Videos , Stream Routes
	  channelGroup.GET("/:id/posts", ListChannelPosts)
	  channelGroup.GET("/:id/reels", ListChannelReels)
	  channelGroup.GET("/:id/videos", ListChannelVideos)
	  channelGroup.GET("/:id/streams", ListChannelStreams)
	*/

}
