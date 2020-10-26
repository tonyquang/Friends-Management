package handlers

import (
	"friends_management/models/request"
	"friends_management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddFriendHandler(c *gin.Context, service services.Service) {
	var friends request.AddFriendRequest
	if err := c.BindJSON(&friends); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err_res := service.AddFriend(friends)

	if res == nil {
		c.JSON(err_res.Code, err_res.Description)
		return
	}

	c.JSON(http.StatusOK, res)
}

func UnFriendHandler(c *gin.Context, service services.Service) {
	var unfriends request.UnFriendRequest
	if err := c.BindJSON(&unfriends); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err_res := service.UnFriend(unfriends)

	if res == nil {
		c.JSON(err_res.Code, err_res.Description)
		return
	}

	c.JSON(http.StatusOK, res)

}
