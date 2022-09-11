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

func GetMessages(c *gin.Context)  {
	var query models.Query

	err := c.ShouldBindQuery(&query); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	results := services.GetMessages(nil, nil, query)

	c.JSON(http.StatusOK, models.Response{Data: results}); return
}

func GetMessage(c *gin.Context)  {
	id := c.Param("id")

	result := services.GetMessage(bson.M{"id": id}, nil)
	if result == nil {
		result = services.GetMessage(bson.M{"slug": id}, nil)
	}

	c.JSON(http.StatusOK, models.Response{Data: result}); return
}

func CreateMessage(c *gin.Context)  {
	var request models.Message

	err := c.ShouldBindJSON(&request); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}


	request.Id			= uuid.New().String()
	request.CreatedAt	= time.Now()

	res, err := services.CreateMessage(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func UpdateMessage(c *gin.Context)  {
	id := c.Param("id")

	var request models.Message

	request.Id		= id

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	res, err := services.UpdateMessage(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func DeleteMessage(c *gin.Context)  {
	id := c.Param("id")

	res, err := services.DeleteMessage(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}
