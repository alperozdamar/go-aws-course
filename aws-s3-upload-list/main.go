package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io/ioutil"
	"os"
)

const s3BucketName = "go-aws--s3-test-alpertest1"
const regionName = "us-east-1"

func main() {
	var (
		s3Client *s3.Client
		err      error
		out      []byte
	)
	if s3Client, err = initS3Client(context.Background()); err != nil {
		fmt.Printf("initConfig error: %s", err)
		os.Exit(1)
	}
	if err = createS3Bucket(context.Background(), s3Client); err != nil {
		fmt.Printf("createS3Bucket error: %s", err)
		os.Exit(1)
	}
	if err = uploadFileToS3(context.Background(), s3Client); err != nil {
		fmt.Printf("uploadFileToS3 error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Uploaded file.\n")
	if out, err = downloadFileFromS3(context.Background(), s3Client); err != nil {
		fmt.Printf("uploadFileToS3 error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Downloaded file with contents: %s", out)
}

func initS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(regionName))
	if err != nil {
		return nil, fmt.Errorf("Config error: %s", err)
	}
	return s3.NewFromConfig(cfg), nil
}

func createS3Bucket(ctx context.Context, s3Client *s3.Client) error {
	allBuckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("ListBuckets error: %s", err)
	}
	found := false
	for _, bucket := range allBuckets.Buckets {
		if *bucket.Name == s3BucketName {
			found = true
		}
	}
	if found == false {
		_, err := s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(s3BucketName),
			CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: regionName,
			},
		})
		if err != nil {
			return fmt.Errorf("CreateBucket error: %s", err)
		}
	}
	return nil
}

func uploadFileToS3(ctx context.Context, s3Client *s3.Client) error {
	//ioutil.ReadFile()//If you wanna read from file.
	testFile, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return fmt.Errorf("Read File Error: %s", err)
	}

	uploader := manager.NewUploader(s3Client)
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s3BucketName),
		Key:    aws.String("test.txt"),
		//Body:   strings.NewReader("This is Alper's test-2 "),
		Body: bytes.NewReader(testFile),
	})
	if err != nil {
		return fmt.Errorf("Upload error: %s", err)
	}
	return nil
}

func downloadFileFromS3(ctx context.Context, s3Client *s3.Client) ([]byte, error) {
	buffer := manager.NewWriteAtBuffer([]byte{})
	downloader := manager.NewDownloader(s3Client)
	numBytes, err := downloader.Download(ctx, buffer, &s3.GetObjectInput{
		Bucket: aws.String(s3BucketName),
		Key:    aws.String("test.txt"),
	})
	if err != nil {
		return buffer.Bytes(), fmt.Errorf("Upload error: %s", err)
	}

	if bytesReceived := int64(len(buffer.Bytes())); numBytes != bytesReceived {
		return buffer.Bytes(), fmt.Errorf("Incorrect number of bytes returned. Got %d, but expected %d", numBytes, bytesReceived)
	}
	return buffer.Bytes(), nil
}
