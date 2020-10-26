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
	// lấy tất cả  transactions
	// lấy tất transactions của 1 users theo accountid nếu accountid rỗng thì lấy hết trans của user đó
	g := r.Group("api/users/")
	{
		// chỉ áp dụng Midleware này cho group này
		//g.Use(MyMidleWare())
		g.POST("/", func(c *gin.Context) {
			AddFriendHandler(c, friendshipService)
		})
	}

	return r
}
