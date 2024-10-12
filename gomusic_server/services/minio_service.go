package service

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"gomusic_server/config"
)

// UploadFileWithPrefix uploads a file to MinIO with a dynamic prefix (folder structure)
func UploadFileWithPrefix(file *multipart.FileHeader, bucketName string, prefix string) (string, error) {
	// 打开文件
	log.Printf("进入minio上传文件方法，路径:" + prefix)
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// 如果 prefix 为空，则不添加任何前缀
	if prefix != "" {
		// 确保 prefix 以 `/` 结尾
		if !strings.HasSuffix(prefix, "/") {
			prefix = prefix + "/"
		}
	}

	// 生成文件名，包括前缀（即“子目录”）
	objectName := generateObjectName(prefix, file.Filename)

	// 检查桶是否存在，如果不存在则创建
	log.Printf("开始检查是否存在bucket....")
	exists, err := config.MinioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Printf("检查途中报错: %v", err)
		return "", fmt.Errorf("failed to check if bucket exists: %v", err)
	}
	if !exists {
		log.Printf("开始创建bucket " + bucketName)
		err = config.MinioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Printf("创建桶失败报错......:")
			return "", fmt.Errorf("failed to create bucket: %v", err)
		}
		log.Printf("Bucket %s created successfully", bucketName)
	}

	// 上传文件到 MinIO
	log.Printf("开始上传至minio.....")
	_, err = config.MinioClient.PutObject(
		context.Background(),
		bucketName,
		objectName,
		src,
		file.Size,
		minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	log.Printf("File uploaded successfully to %s", objectName)
	return objectName, nil
}

// RemoveFile removes a file from MinIO
func RemoveFile(objectName string, bucketName string) error {
	log.Print("进入Minio删除文件方法")
	err := config.MinioClient.RemoveObject(
		context.Background(),
		bucketName,
		objectName,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		log.Print("Minio移除文件失败！")
		return fmt.Errorf("failed to remove file: %v", err)
	}

	log.Printf("File removed successfully: %s", objectName)
	return nil
}

func generateObjectName(prefix string, fileName string) string {
	// 生成 SHA-1 哈希值
	hash := sha1.New()
	hash.Write([]byte(fileName))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)[:8] // 取前 8 个字符

	return fmt.Sprintf("%s%d-%s%s", prefix, time.Now().Unix(), hashString, filepath.Ext(fileName))
}
