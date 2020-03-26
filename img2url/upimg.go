package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var endPoint = flag.String("endpoint", "https://s3-cn-east-1.qiniucs.com", "endpoint with http(s):")
var region = flag.String("region", "cn-east-1", "region")
var bucketName = flag.String("bucket", "test", "bucket name")
var absFilePath = flag.String("file", "absFilePath", "the abs path of file that you want to upload to S3.")
var accessKeyID = flag.String("ak", "accesskey", "accesskey")
var secretKeyID = flag.String("sk", "secretkey", "secretkey")

// func upload() {

// }

func main() {
	flag.Parse()
	file, err := os.Open(*absFilePath)
	now := time.Now()
	fileNameString := fmt.Sprintf("/%d/%d/%d/%d-%d-%d-%s.png", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), RandString(8))

	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()
	conf := &aws.Config{
		Region:           aws.String(*region),
		Endpoint:         aws.String(*endPoint),
		Credentials:      credentials.NewStaticCredentials(*accessKeyID, *secretKeyID, ""),
		S3ForcePathStyle: aws.Bool(true),
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{Config: *conf}))
	// service := s3.New(sess)
	uploader := s3manager.NewUploader(sess)
	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	uploadOutput, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*bucketName),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(fileNameString),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: file,
		// ContentType: aws.String("image/png"),
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", *absFilePath, *bucketName, err)
	}
	fmt.Println(uploadOutput.Location)

	// fmt.Printf("Successfully uploaded %q to %q\n", *absFilePath, *bucketName)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandString 生成随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
