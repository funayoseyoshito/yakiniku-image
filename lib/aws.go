package lib

import (
	"log"

	"image"

	"bytes"
	"image/jpeg"

	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	//S3ACLPublic is setting to public acl
	S3ACLPublic = true
	//S3ACLPrivate is setting to private acl
	S3ACLPrivate = false
)

//AwsIni
type AwsIni struct {
	AccessKeyID     string
	SecretAccessKey string
	S3BucketName    string
}

//SaveImageToS3 is save image to s3 object
func (this *AwsIni) SaveImageToS3(imgID int, img *image.Image, imgQuality int, public bool) {
	var opt jpeg.Options
	opt.Quality = imgQuality

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, &opt); err != nil {
		log.Println("unable to encode image.")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(this.AccessKeyID, this.SecretAccessKey, ""),
	})

	if err != nil {
		FatalExit(err)
	}

	uploader := s3manager.NewUploader(sess)

	s3Input := &s3manager.UploadInput{
		Bucket:      aws.String(Config.Aws.GetAwsBucketName()),
		Key:         aws.String(strconv.Itoa(imgID) + "." + SaveImageExt),
		Body:        bytes.NewBuffer(buffer.Bytes()),
		ContentType: aws.String("image/jpeg"),
	}

	if public {
		s3Input.ACL = aws.String("public-read")
	}

	_, err = uploader.Upload(s3Input)

	if err != nil {
		FatalExit(err)
	}
}
