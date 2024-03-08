package storage

// func (s *Storage) UploadToS3(bucketName string, imageID int, imageData []byte) (string, error) {
// 	// Upload imageData to S3 bucket 'bucketName' and return the URL
// 	// Example using AWS SDK for Go (replace 'your-region' with your actual AWS region)
// 	uploader := s3manager.NewUploaderWithClient(s3.New(s.session.NewSession(&aws.Config{
// 		Region: aws.String("your-region"),
// 	})))
// 	result, err := uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(fmt.Sprintf("images/%d.jpg", imageID)),
// 		Body:   bytes.NewReader(imageData),
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	return result.Location, nil
// }

//https://aws.github.io/aws-sdk-go-v2/docs/sdk-utilities/s3/
