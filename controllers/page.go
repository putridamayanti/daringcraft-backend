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

func GetPages(c *gin.Context)  {
	var query models.Query

	err := c.ShouldBindQuery(&query); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	results := services.GetPages(nil, nil, query)

	c.JSON(http.StatusOK, models.Response{Data: results}); return
}

func GetPage(c *gin.Context)  {
	id := c.Param("id")

	result := services.GetPage(bson.M{"id": id}, nil)
	if result == nil {
		result = services.GetPage(bson.M{"code": id}, nil)
	}

	c.JSON(http.StatusOK, models.Response{Data: result}); return
}

func GetPageByCode(c *gin.Context)  {
	code := c.Param("code")

	result := services.GetPage(bson.M{"code": code}, nil)

	c.JSON(http.StatusOK, models.Response{Data: result}); return
}

func CreatePage(c *gin.Context)  {
	var request models.Page

	err := c.ShouldBindJSON(&request); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}


	request.Id			= uuid.New().String()
	request.CreatedAt	= time.Now()

	res, err := services.CreatePage(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func UpdatePage(c *gin.Context)  {
	id := c.Param("id")

	var request models.Page

	request.Id		= id

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	res, err := services.UpdatePage(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func DeletePage(c *gin.Context)  {
	id := c.Param("id")

	res, err := services.DeletePage(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}
