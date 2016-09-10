package redis

import (
	rlib "github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
)

func SetCmd(ctx context.Context, redisname, key, val string) (string, error) {
	rconn := GetWriteConn(ctx, redisname)
	return rlib.String(rconn.Do("SET", key, val))
}

func GetCmd(ctx context.Context, redisname, key string) (string, error) {
	rconn := GetReadConn(ctx, redisname)
	return rlib.String(rconn.Do("GET", key))
}
