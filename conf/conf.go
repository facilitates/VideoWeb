package conf

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/ini.v1"
	"strings"
	"videoweb/dao"
)

var (
	HttpPort         string
	DbHost           string
	DbPort           string
	DbUser           string
	DbPassWord       string
	DbName           string
	RedisAddr        string
	RedisPW          string
	RedisDbName      int
	RabbitMqUserName string
	RabbitMqPassword string
)

var UserConns = make(map[string]*websocket.Conn)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {

		fmt.Println("配置文件读取错误，请检查文件路径")
	}
	LoadServer(file)
	LoadMysql(file)
	LoadRedis(file)
	LoadRabbitMQ(file)
	mysqlPath := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	rabbitMqPath := strings.Join([]string{"amqp://", RabbitMqUserName, ":", RabbitMqPassword, "@localhost:5672/"}, "")
	dao.Init(mysqlPath, RedisAddr, RedisPW, RedisDbName, rabbitMqPath)
}

func LoadServer(file *ini.File) {
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPW = file.Section("redis").Key("RedisPW").String()
	RedisDbName, _ = file.Section("redis").Key("RedisDbName").Int()
}

func LoadRabbitMQ(file *ini.File) {
	RabbitMqUserName = file.Section("rabbitmq").Key("RabbitMqUserName").String()
	RabbitMqPassword = file.Section("rabbitmq").Key("RabbitMqPassword").String()
}
