package config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

// 定义常量
const (
	FilePath        = "config/application.ini"
	SectionDatabase = "database"
	SectionRedis    = "redis"
	SectionServer   = "server"
)

type Server struct {
	RunMode      string
	HttpPort     uint16
	ReadTimeout  uint16
	WriteTimeout uint16
}

var ServerSetting = new(Server)

type Database struct {
	Dialect      string
	User         string
	Password     string
	Host         string
	DatabaseName string
	TablePrefix  string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

// ReadProperties 读取配置文件属性
func ReadProperties() {
	var err error
	cfg, err = ini.Load(FilePath)
	if err != nil {
		log.Fatalf("setting read file application.ini failed, err: %v", err)
	}

	mapTo(SectionServer, ServerSetting)
	mapTo(SectionDatabase, DatabaseSetting)
	mapTo(SectionRedis, RedisSetting)
}

// 读取配置文件的指定部分
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("setting mapTo %s failed, err: %v", section, err)
	}
}
