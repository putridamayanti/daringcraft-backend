package services

import (
	"context"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
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
