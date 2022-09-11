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

func GetSubscribers(c *gin.Context)  {
	var query models.Query

	err := c.ShouldBindQuery(&query); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	results := services.GetSubscribers(nil, nil, query)

	c.JSON(http.StatusOK, models.Response{Data: results}); return
}

func GetSubscriber(c *gin.Context)  {
	id := c.Param("id")

	result := services.GetSubscriber(bson.M{"id": id}, nil)
	if result == nil {
		result = services.GetSubscriber(bson.M{"slug": id}, nil)
	}

	c.JSON(http.StatusOK, models.Response{Data: result}); return
}

func CreateSubscriber(c *gin.Context)  {
	var request models.Subscriber

	err := c.ShouldBindJSON(&request); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}


	request.Id			= uuid.New().String()
	request.CreatedAt	= time.Now()

	res, err := services.CreateSubscriber(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func UpdateSubscriber(c *gin.Context)  {
	id := c.Param("id")

	var request models.Subscriber

	request.Id		= id

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	res, err := services.UpdateSubscriber(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func DeleteSubscriber(c *gin.Context)  {
	id := c.Param("id")

	res, err := services.DeleteSubscriber(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}
