package controllers

import (
	"daringcraft-backend/models"
	"daringcraft-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func GetMedias(c *gin.Context)  {
	var query models.Query

	err := c.ShouldBindQuery(&query); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	results := services.GetBanners(nil, nil, query)

	c.JSON(http.StatusOK, models.Response{Data: results}); return
}

func CreateMedia(c *gin.Context)  {
	var request models.Media

	err := c.ShouldBindJSON(&request); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	request.Id			= uuid.New().String()
	request.CreatedAt	= time.Now()

	err = services.CreateMedia(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: "Success"}); return
}

func DeleteMedia(c *gin.Context)  {
	id := c.Param("id")

	res, err := services.DeleteMedia(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}
