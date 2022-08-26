package utils

import (
	"baas-clean/config"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// Select 选择库
func Select(db uint8) (redis.Conn, error) {
	conn := config.RedisDB.Get()
	if _, err := conn.Do("SELECT", db); err != nil {
		return nil, err
	}
	return conn, nil
}

// Delete 删除
func Delete(key string, db uint8) (bool, error) {
	conn, _ := Select(db)
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	if conn != nil {
		return redis.Bool(conn.Do("DEL", key))
	}
	return false, nil
}

func delete(key string, conn redis.Conn) {
	if conn == nil {
		return
	}
	_, err := conn.Do("DEL", key)
	if err != nil {
		return
	}
}

// LikeDelete 批量删除
func LikeDelete(key string, db uint8) error {
	fmt.Printf("redis delete like key %s, db is %d \n", key, db)
	conn, _ := Select(db)
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		delete(key, conn)
	}
	return nil
}
