package dao

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var DB *gorm.DB

var Redisdb *redis.Client

var Conn *amqp.Connection

func initMysql(connstring string) {
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		log.Println(err)
	}
	DB = db
	DB.LogMode(true) //启动模式
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(20) // 设置连接池
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(time.Second * 30)
}

func initRedis(RedisAddr string, RedisPW string, RedisDbName int) {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPW,
		DB:       RedisDbName,
	})
	_, err := Redisdb.Ping().Result()
	if err != nil {
		log.Println(err)
	}
}

func initRabbitMQ(RabbitMqPath string) {
	conn, err := amqp.Dial(RabbitMqPath)
	if err != nil {
		log.Println(err)
	}
	Conn = conn
}

func Init(connstring string, RedisAddr string, RedisPW string, RedisDbName int, RabbitMqPath string) {
	initMysql(connstring)
	initRedis(RedisAddr, RedisPW, RedisDbName)
	initRabbitMQ(RabbitMqPath)
	migration()
}
