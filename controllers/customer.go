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

func GetCustomers(c *gin.Context)  {
	var query models.Query

	err := c.ShouldBindQuery(&query); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	results := services.GetCustomers(nil, nil, query)

	c.JSON(http.StatusOK, models.Response{Data: results}); return
}

func GetCustomer(c *gin.Context)  {
	id := c.Param("id")

	result := services.GetCustomer(bson.M{"id": id}, nil)
	if result == nil {
		result = services.GetCustomer(bson.M{"slug": id}, nil)
	}

	c.JSON(http.StatusOK, models.Response{Data: result}); return
}

func CreateCustomer(c *gin.Context)  {
	var request models.Customer

	err := c.ShouldBindJSON(&request); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}


	request.Id			= uuid.New().String()
	request.CreatedAt	= time.Now()

	res, err := services.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func UpdateCustomer(c *gin.Context)  {
	id := c.Param("id")

	var request models.Customer

	request.Id		= id

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	res, err := services.UpdateCustomer(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func DeleteCustomer(c *gin.Context)  {
	id := c.Param("id")

	res, err := services.DeleteCustomer(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}
