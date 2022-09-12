package controllers

import (
	"daringcraft-backend/models"
	"daringcraft-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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