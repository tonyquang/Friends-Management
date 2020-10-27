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
	var friendship_id string = ""
	friendship_id = c.Param("friendship_id")

	if friendship_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not valid!"})
		return
	}

	res, err_res := service.UnFriend(friendship_id)

	if res == nil {
		c.JSON(err_res.Code, err_res.Description)
		return
	}

	c.JSON(http.StatusOK, res)
}

func ListFriendHandler(c *gin.Context, service services.Service) {
	var mailAddress string = ""
	mailAddress = c.Param("email_address")

	if mailAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not valid!"})
		return
	}

	res, err_res := service.ViewListFriendsByEmail(mailAddress)

	if err_res != nil {
		c.JSON(err_res.Code, err_res.Description)
		return
	}

	c.JSON(http.StatusOK, res)

}

func ListCommonFriendHandler(c *gin.Context, service services.Service) {
	var UserOne, UserTwo string = "", ""

	UserOne = c.Query("user_one")
	UserTwo = c.Query("user_two")

	if UserOne == "" || UserTwo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not valid!"})
		return
	}
	var friends request.AddFriendRequest
	friends.Friends = append(friends.Friends, UserOne, UserTwo)
	res, err := service.ViewListCommonFriendsByEmail(friends)

	if err != nil {
		c.JSON(err.Code, err.Description)
		return
	}

	c.JSON(http.StatusOK, res)
}
