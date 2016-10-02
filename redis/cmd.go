package redis

import (
	rlib "github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

func Incr(ctx context.Context, redisname, key string) (int64, error) {
	rconn := GetWriteConn(ctx, redisname)
	return rlib.Int64(rconn.Do("INCR", key))
}

func SetString(ctx context.Context, redisname, key, val string) (string, error) {
	rconn := GetWriteConn(ctx, redisname)
	return rlib.String(rconn.Do("SET", key, val))
}

func GetString(ctx context.Context, redisname, key string) (string, error) {
	rconn := GetReadConn(ctx, redisname)
	return rlib.String(rconn.Do("GET", key))
}

func GetPackedValue(ctx context.Context, redisname, key string, out interface{}) (interface{}, error) {
	rconn := GetReadConn(ctx, redisname)
	value, err := rlib.Bytes(rconn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	if err := msgpack.Unmarshal(value, out); err != nil {
		return nil, err
	}
	return out, nil
}
