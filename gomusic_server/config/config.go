package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MySQL and MinIO configuration structs
type Config struct {
	MySQL MySQLConfig
	Minio MinioConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	Host      string
	Port      int
	Username  string
	Password  string
	DBName    string
	Charset   string
	ParseTime bool
	Loc       string
}

type MinioConfig struct {
	Endpoint     string
	RootUser     string
	RootPassword string
	UseSSL       bool
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

var DB *gorm.DB
var MinioClient *minio.Client
var RedisClient *redis.Client
var AppConfig Config

// InitConfig initializes configuration from a YAML file
func InitConfig() {
	viper.SetConfigName("config")   // 配置文件名称 (不包含扩展名)
	viper.SetConfigType("yaml")     // 配置文件格式
	viper.AddConfigPath(".")        // 配置文件所在路径
	viper.AddConfigPath("./config") // 添加额外的配置文件路径

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 解析配置文件内容
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

// InitDB initializes the MySQL database connection using GORM
func InitDB() {
	InitConfig()

	// 构建 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		AppConfig.MySQL.Username,
		AppConfig.MySQL.Password,
		AppConfig.MySQL.Host,
		AppConfig.MySQL.Port,
		AppConfig.MySQL.DBName,
		AppConfig.MySQL.Charset,
		AppConfig.MySQL.ParseTime,
		AppConfig.MySQL.Loc,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 添加重试逻辑
		for i := 0; i < 10; i++ {
			DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err == nil {
				log.Println("MySQL connection established successfully.")
				return
			}
			log.Printf("Failed to connect to database: %v. Retrying in 2 seconds...", err)
			time.Sleep(2 * time.Second)
		}
	}
}

// InitMinio initializes the MinIO client
func InitMinio() {
	// 使用从配置文件读取的 MinIO 配置
	client, err := minio.New(AppConfig.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AppConfig.Minio.RootUser, AppConfig.Minio.RootPassword, ""),
		Secure: AppConfig.Minio.UseSSL,
	})
	log.Println("MinIO Endpoint:", AppConfig.Minio.Endpoint)
	if err != nil {
		log.Fatalln("Failed to initialize MinIO:", err)
	}

	MinioClient = client
	log.Println("MinIO client initialized successfully--MinIO 客户端初始化成功！")
}

// InitRedis initializes the Redis client
func InitRedis() {
	// 从配置文件中读取 Redis 的配置信息
	redisAddr := fmt.Sprintf("%s:%d", AppConfig.Redis.Host, AppConfig.Redis.Port)

	// 创建 Redis 客户端
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: AppConfig.Redis.Password, // 没有密码时可留空
		DB:       AppConfig.Redis.DB,       // 使用的 Redis DB
	})

	// 测试 Redis 连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connection established successfully.")
}
