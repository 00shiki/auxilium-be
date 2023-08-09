package storage

import (
	"auxilium-be/pkg/utils"
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"io"
	"mime/multipart"
	"os"
)

type Client struct{}

func ClientInit() *Client {
	return &Client{}
}

func (r *Client) Connect() (*storage.Client, error) {
	ctx := context.Background()
	bucketVar := os.Getenv("BUCKET_CREATOR")
	newCreds := utils.ConvertStringToJSON(bucketVar)
	creds := option.WithCredentialsJSON(newCreds)
	client, err := storage.NewClient(ctx, creds)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (r *Client) UploadToBucket(file multipart.File, object string) (string, error) {

	ctx := context.Background()
	conn, err := r.Connect()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	bucketName := os.Getenv("BUCKET_NAME")
	folderName := os.Getenv("FOLDER_NAME")
	bucket := conn.Bucket(bucketName)
	o := bucket.Object(folderName + object)
	o = o.If(storage.Conditions{DoesNotExist: true})

	wc := o.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	url := "https://storage.googleapis.com/" + bucketName + "/" + folderName + object
	return url, nil
}
