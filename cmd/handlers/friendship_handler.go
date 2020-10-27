package handlers

import (
	"friends_management/models/request"
	"friends_management/services"
	"friends_management/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddFriendHandler(c *gin.Context, service services.Service) {
	var friends request.AddFriendRequest
	if err := c.BindJSON(&friends); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.AddFriend(friends)

	utils.Respone(res, err, c)
}

func UnFriendHandler(c *gin.Context, service services.Service) {
	var friendship_id string = ""
	friendship_id = c.Param("friendship_id")

	if friendship_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not valid!"})
		return
	}

	res, err := service.UnFriend(friendship_id)

	utils.Respone(res, err, c)
}

func ListFriendHandler(c *gin.Context, service services.Service) {
	var mailAddress string = ""
	mailAddress = c.Param("email_address")

	if mailAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not valid!"})
		return
	}

	res, err_res := service.ViewListFriendsByEmail(mailAddress)

	utils.Respone(res, err_res, c)

}

func ListCommonFriendHandler(c *gin.Context, service services.Service) {
	var userOne, userTwo string = "", ""

	userOne = c.Query("user_one")
	userTwo = c.Query("user_two")

	if userOne == "" || userTwo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not valid!"})
		return
	}
	var friends request.AddFriendRequest
	friends.Friends = append(friends.Friends, userOne, userTwo)
	res, err := service.ViewListCommonFriendsByEmail(friends)

	utils.Respone(res, err, c)
}

func SubscribeUpdateHanler(c *gin.Context, service services.Service) {
	var requestor, target string = "", ""

	requestor = c.Param("requestor")
	target = c.Param("target")

	if requestor == "" || target == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not valid!"})
		return
	}

	res, err := service.SubscribeUpdate(request.HandleUpdateRequest{Requestor: requestor, Target: target})
	utils.Respone(res, err, c)
}

func BlockUpdateHanler(c *gin.Context, service services.Service) {
	var requestor, target string = "", ""

	requestor = c.Param("requestor")
	target = c.Param("target")

	if requestor == "" || target == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not valid!"})
		return
	}

	res, err := service.BlockUpdate(request.HandleUpdateRequest{Requestor: requestor, Target: target})
	utils.Respone(res, err, c)
}

func ReceviceUserCanUpdate(c *gin.Context, service services.Service) {
	var sender string = ""

	sender = c.Param("sender")

	if sender == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not valid!"})
		return
	}

	res, err := service.ViewListFriendsRecvUpdate(sender)

	utils.Respone(res, err, c)
}
