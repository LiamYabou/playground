package main

import (
    "fmt"
    "os"
    "time"
    "net/http"
    "strings"
    "math/rand"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
    region = os.Getenv("AWS_S3_REGION")
    bucketName = os.Getenv("AWS_S3_BUCKET_NAME")
    accessId = os.Getenv("AWS_ACCESS_KEY_ID")
    accessSecret = os.Getenv("AWS_SECRET_ACCESS_KEY")
)

func main() {
    var config *aws.Config
    config = &aws.Config{
        Region: aws.String(region),
        Credentials: credentials.NewStaticCredentials(accessId, accessSecret, ""),
    }
    // The session uses the above config.
    sess := session.Must(session.NewSession(config))
    // Creat an uploader with the seeesion and the default options.
    uploader := s3manager.NewUploader(sess)
    // Fetch the stream from the following url
    url := "https://images-na.ssl-images-amazon.com/images/I/6182S7MYC2L._AC_UL200_SR200,200_.jpg"
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Could not request successfully, err: %s\n", err)
    }
    defer resp.Body.Close()
    // # Upload the data stream to the s3 bucket
    // ## Generate the random name of the file
    rand.Seed(time.Now().UnixNano())
    chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
                    "abcdefghijklmnopqrstuvwxyz" +
                    "0123456789")
    length := 16
    var b strings.Builder
    for i := 0; i < length; i++ {
        b.WriteRune(chars[rand.Intn(len(chars))])
    }
    fileNmae := b.String()
    fileExtesion := ".jpg"
    filePath := fmt.Sprintf("images/%s%s", fileNmae, fileExtesion)
    result, err := uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String(bucketName),
        Key: aws.String(filePath),
        Body: resp.Body,
        ContentType: aws.String("image/jpeg"),
    })
    if err != nil {
        fmt.Printf("Could not upload the file into the s3 bucket, err: %s\n", err)
    }
    fmt.Printf("The file has been uploaded to %s\n", result.Location)
}
