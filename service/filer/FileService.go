package filer

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
	"study_room_management_backend/config"
	"time"
)

var (
	cfg    = config.C.Storage
	s3sess *session.Session
)

func init() {
	if cfg.Type == "minio" || cfg.Type == "s3" {
		s3sess, _ = session.NewSession(&aws.Config{
			Region:           aws.String(cfg.S3.Region),
			Endpoint:         aws.String(cfg.Minio.Endpoint),
			Credentials:      credentials.NewStaticCredentials(cfg.Minio.AccessKey, cfg.Minio.SecretKey, ""),
			DisableSSL:       aws.Bool(!cfg.Minio.UseSSL),
			S3ForcePathStyle: aws.Bool(true), // minio 需要
		})
	}
}

func GenObjectKey(folder string, uid int64, ext string) string {
	return fmt.Sprintf("%s/%d_%d%s", folder, uid, time.Now().Unix(), ext)
}

func SaveFile(r io.Reader, size int64, key string) (string, error) {
	switch cfg.Type {
	case "local":
		return saveLocal(r, key)
	case "minio", "s3":
		return saveS3(r, key, size)
	}
	return "", fmt.Errorf("unknown storage type")
}

func saveLocal(r io.Reader, key string) (string, error) {
	p := path.Join(cfg.Local.Path, key)
	if err := os.MkdirAll(path.Dir(p), 0755); err != nil {
		return "", err
	}
	f, err := os.Create(p)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	if _, err := io.Copy(f, r); err != nil {
		return "", err
	}
	// 返回可访问的 URL，这里用 nginx 挂静态目录
	return fmt.Sprintf("http://"+viper.GetString("server.port")+"/static/%s", key), nil
}

func saveS3(r io.Reader, key string, size int64) (string, error) {
	bucket := cfg.Minio.Bucket
	_, err := s3.New(s3sess).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(key),
		Body:          aws.ReadSeekCloser(r),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String("image/jpeg"),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", cfg.Minio.Endpoint, bucket, key), nil
}
