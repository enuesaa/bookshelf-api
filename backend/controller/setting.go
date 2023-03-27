package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/repository/redis"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/validate"
)

func GetAppearance(c *gin.Context) {
	var body v1.SettingGetAppearanceRequest
	if err := validate.Validate(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value := redis.GetValue(c, "aaa")
	c.JSON(http.StatusOK, v1.SettingGetAppearanceResponse {
		Message: value,
	})
}

func PutAppearance(c *gin.Context) {
	var body v1.SettingPutAppearanceRequest
	if err := validate.Validate(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redis.SetValue(c, "aaa", "bbb")
	c.JSON(http.StatusOK, v1.SettingPutAppearanceResponse {})
}