package config

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var MySqlDB *gorm.DB

// ConnectMySql 初始化MySql连接
func ConnectMySql() {
	var err error
	dialect := DatabaseSetting.Dialect
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.DatabaseName)

	MySqlDB, err = gorm.Open(dialect, url)
	if err != nil {
		log.Fatalf("gorm open mysql failed, err: %v", err)
		return
	}

	MySqlDB.SingularTable(true)
	MySqlDB.DB().SetMaxIdleConns(10)
	MySqlDB.DB().SetMaxOpenConns(100)
}

// CloseMySqlDB 关闭连接
func CloseMySqlDB() {
	err := MySqlDB.Close()
	if err != nil {
		log.Fatalf("mysql close connect failed, err: %v", err)
	}
}

var RedisDB *redis.Pool

// ConnectRedis 连接Redis
func ConnectRedis() error {
	RedisDB = &redis.Pool{
		MaxIdle:     RedisSetting.MaxIdle,
		MaxActive:   RedisSetting.MaxActive,
		IdleTimeout: RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if RedisSetting.Password != "" {
				_, err := conn.Do("AUTH", RedisSetting.Password)
				if err != nil {
					err := conn.Close()
					if err != nil {
						return nil, err
					}
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}
