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

func GetMantras(c *gin.Context)  {
	var query models.Query

	err := c.ShouldBindQuery(&query); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	results := services.GetMantras(nil, nil, query)

	c.JSON(http.StatusOK, models.Response{Data: results}); return
}

func GetMantra(c *gin.Context)  {
	id := c.Param("id")

	result := services.GetMantra(bson.M{"id": id}, nil)
	if result == nil {
		result = services.GetMantra(bson.M{"slug": id}, nil)
	}

	c.JSON(http.StatusOK, models.Response{Data: result}); return
}

func CreateMantra(c *gin.Context)  {
	var request models.Mantra

	err := c.ShouldBindJSON(&request); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}


	request.Id			= uuid.New().String()
	request.CreatedAt	= time.Now()

	res, err := services.CreateMantra(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func UpdateMantra(c *gin.Context)  {
	id := c.Param("id")

	var request models.Mantra

	request.Id		= id

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	res, err := services.UpdateMantra(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func DeleteMantra(c *gin.Context)  {
	id := c.Param("id")

	res, err := services.DeleteMantra(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}
