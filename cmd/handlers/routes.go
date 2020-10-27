package handlers

import (
	"database/sql"

	"friends_management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// API routes for APIs
func API(dbconn *sql.DB) http.Handler {
	friendshipService := services.NewManager(dbconn)
	gin.SetMode(gin.ReleaseMode)
	//fmt.Println("%T", friendshipService)
	r := gin.Default()
	// r.POST("api/users", func(c *gin.Context) {
	// 	AddFriendHandler(c, friendshipService)
	// })
	// r.DELETE("api/users", func(c *gin.Context) {
	// 	UnFriendHandler(c, friendshipService)
	// })

	g := r.Group("/api")
	{

		g.GET("/friends/:email_address", func(c *gin.Context) {
			ListFriendHandler(c, friendshipService)
		})

		g.GET("/commonfriends", func(c *gin.Context) {
			ListCommonFriendHandler(c, friendshipService)
		})

		g.POST("/addfriend", func(c *gin.Context) {
			AddFriendHandler(c, friendshipService)
		})

		g.DELETE("/unfriend/:friendship_id", func(c *gin.Context) {
			UnFriendHandler(c, friendshipService)
		})

		g.PUT("/subscribe/:requestor/:target", func(c *gin.Context) {
			SubscribeUpdateHanler(c, friendshipService)
		})

		g.PUT("/block/:requestor/:target", func(c *gin.Context) {
			BlockUpdateHanler(c, friendshipService)
		})

	}

	return r
}
