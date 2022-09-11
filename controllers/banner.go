package controllers

import (
	"daringcraft-backend/models"
	"daringcraft-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func GetBanners(c *gin.Context)  {
	var query models.Query

	err := c.ShouldBindQuery(&query); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	results := services.GetBanners(nil, nil, query)

	c.JSON(http.StatusOK, models.Response{Data: results}); return
}

func GetBanner(c *gin.Context)  {
	id := c.Param("id")

	result := services.GetBanner(bson.M{"id": id}, nil)
	if result == nil {
		result = services.GetBanner(bson.M{"slug": id}, nil)
	}

	c.JSON(http.StatusOK, models.Response{Data: result}); return
}

func CreateBanner(c *gin.Context)  {
	var request models.Banner

	err := c.ShouldBindJSON(&request); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}


	request.Id			= uuid.New().String()
	request.CreatedAt	= time.Now()

	res, err := services.CreateBanner(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func UpdateBanner(c *gin.Context)  {
	id := c.Param("id")

	var request models.Banner

	request.Id		= id

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	res, err := services.UpdateBanner(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func DeleteBanner(c *gin.Context)  {
	id := c.Param("id")

	res, err := services.DeleteBanner(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}
