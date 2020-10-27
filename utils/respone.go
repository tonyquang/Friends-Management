package utils

import (
	model_res "friends_management/models/respone"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Respone(res interface{}, err *model_res.ResponseError, c *gin.Context) {
	if err != nil {
		c.JSON(err.Code, err.Description)
		return
	}
	c.JSON(http.StatusOK, res)
}
