package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	// 打印viper读取到的app.yml下的配置项，这里最开始app为空，mysql有值
	fmt.Println("config app inited....")
}

func InitMysql() {

	// 自定义日志模板，打印SQL语句
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        //颜色
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newlogger})

	fmt.Println("mysql inited....")
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})

	// 测试redis是否连接
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	pong, err := Red.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redistest err,", err)
		return
	}

	fmt.Println("redis inited....", pong)
}

const (
	PubkushKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel, message string) (err error) {
	fmt.Println("------------------")

	fmt.Println("publish message :", message)
	err = Red.Publish(ctx, channel, message).Err()
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	fmt.Println("+++++++++++++++++++++++")
	sub := Red.Subscribe(ctx, channel)
	//ch := sub.Channel()
	//
	//for message := range ch {
	//	fmt.Println("Subscribe message :", message.Payload)
	//	return message.Payload, nil
	//}
	//return "", nil

	mes, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe message :", mes.Payload)
	return mes.Payload, err
}
