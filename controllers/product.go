package controllers

import (
	"daringcraft-backend/lib"
	"daringcraft-backend/models"
	"daringcraft-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
	"time"
)

func GetProducts(c *gin.Context)  {
	var query models.Query

	err := c.ShouldBindQuery(&query); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	results := services.GetProducts(nil, nil, query)

	c.JSON(http.StatusOK, models.Response{Data: results}); return
}

func GetProduct(c *gin.Context)  {
	id := c.Param("id")

	result := services.GetProduct(bson.M{"id": id}, nil)
	if result == nil {
		result = services.GetProduct(bson.M{"slug": id}, nil)
	}

	c.JSON(http.StatusOK, models.Response{Data: result}); return
}

func CreateProduct(c *gin.Context)  {
	var request models.Product

	err := c.ShouldBindJSON(&request); if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}


	request.Id			= uuid.New().String()
	request.Slug		= lib.SlugGenerator(request.Name)
	request.CreatedAt	= time.Now()

	res, err := services.CreateProduct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func UpdateProduct(c *gin.Context)  {
	id := c.Param("id")

	var request models.Product

	request.Id		= id

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	res, err := services.UpdateProduct(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func DeleteProduct(c *gin.Context)  {
	id := c.Param("id")

	res, err := services.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Response{Data: res}); return
}

func PrintfulGetProducts(c *gin.Context)  {
	res, err := lib.PrintfulSendRequest("GET", "store/products")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	var printfulResult models.PrintfulResult

	result, _ := json.Marshal(res)
	err = json.Unmarshal(result, &printfulResult)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	product := printfulResult.Result

	c.JSON(http.StatusOK, models.Response{Data: product}); return
}

func PrintfulGetProductById(c *gin.Context)  {
	id := c.Param("id")

	res, err := lib.PrintfulSendRequest("GET", "store/products/" + id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	var printfulResult models.PrintfulResult

	result, _ := json.Marshal(res)
	err = json.Unmarshal(result, &printfulResult)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Data: err.Error()}); return
	}

	product := printfulResult.Result

	variants := make([]models.ProductVariant, 0)
	images := make([]string, 0)

	for _, val := range product.SyncVariants {
		variant := models.ProductVariant{
			Id:     	strconv.Itoa(val.Id),
			ProductId: 	strconv.Itoa(val.ProductId),
			Name:   	val.Name,
			Price:  	val.RetailPrice,
			Status: 	val.IsIgnored,
			SKU:    	val.Sku,
		}

		variants = append(variants, variant)
	}

	images = append(images, product.SyncProduct.Thumbnail)

	status := "PUBLISHED"
	if !product.SyncProduct.IsIgnored {
		status = "ARCHIVED"
	}

	id = strconv.Itoa(product.SyncProduct.Id)

	final := models.Product{
		Id: id,
		Name: product.SyncProduct.Name,
		Images: images,
		Variants: variants,
		PreOrder: 7,
		Description: product.SyncProduct.Name,
		Status: status,
		CreatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, models.Response{Data: final}); return
}

