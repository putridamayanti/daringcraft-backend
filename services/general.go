package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

const MediaCollection = "medias"


func CloudinaryUpload(path string, filename string) (error, *string) {
	cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_NAME"), os.Getenv("CLOUDINARY_KEY"), os.Getenv("CLOUDINARY_SECRET"))

	_, err := cld.Upload.Upload(context.Background(), path, uploader.UploadParams{PublicID: filename})
	if err != nil {
		return err, nil
	}

	image, err := cld.Image(filename)
	if err != nil {
		return err, nil
	}

	url, err := image.String()
	if err != nil {
		return err, nil
	}

	err = os.Remove(path)

	return nil, &url
}

func CreateMedia(media models.Media) error {
	_, err := database.InsertOne(MediaCollection, media)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMedia(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(MediaCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}