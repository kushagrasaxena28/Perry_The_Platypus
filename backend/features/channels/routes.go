package channels

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func ChannelsRoutes(r *gin.Engine, db *pgx.Conn) {
	channelGroup := r.Group("/v1/channels")

	//Core Channel Routes
	channelGroup.GET("", func(c *gin.Context) {})
	channelGroup.POST("", func(c *gin.Context) {})
	channelGroup.GET("/:id", func(c *gin.Context) {})
	channelGroup.PUT("/:id", func(c *gin.Context) {})
	channelGroup.DELETE("/:id", func(c *gin.Context) {})

	// Channel Membership Routes
	channelGroup.POST("/:id/subscribe", func(c *gin.Context) {})
	channelGroup.DELETE("/:id/unsubscribe", func(c *gin.Context) {})
	channelGroup.GET("/:id/subscribers", func(c *gin.Context) {})

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
