package controllers

import (
	"daringcraft-backend/models"
	"daringcraft-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func Upload(c *gin.Context)  {
	response := models.Response{Status: 200}

	file, _ := c.FormFile("file")

	filename := strconv.Itoa(int(time.Now().UnixNano()))

	path := "uploads/" + filename + filepath.Ext(file.Filename)
	err := c.SaveUploadedFile(file, path)
	if err != nil {
		response.Status = 400
		response.Data = err.Error()
		response.Message = "Error Temporary"
	}

	err, url := services.CloudinaryUpload(path, filename)
	if err != nil {
		c.JSON(400, models.Response{
			Data: err.Error(),
		}); return
	}

	err = services.CreateMedia(models.Media{
		Id:        uuid.New().String(),
		Url:       url,
		CreatedAt: time.Now(),
	})

	c.JSON(200, models.Response{
		Data: url,
	}); return
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
