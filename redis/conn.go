package redis

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	rlib "github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"

	"github.com/satoshi03/go/config"
)

type redisConn struct {
	master *rlib.Pool
	slaves []*rlib.Pool
	name   string
}

func Open(ctx context.Context, rconf config.Redis, redisname string) context.Context {
	// connection pool for master
	master := newPool(rconf.Host, rconf.Port, rconf.DB)

	// connection pool for slaves
	slaves := make([]*rlib.Pool, 0, len(rconf.Slaves))
	for _, sconf := range rconf.Slaves {
		slave := newPool(sconf.Host, sconf.Port, sconf.DB)
		slaves = append(slaves, slave)
	}

	// if no slaves, use master as slave
	if len(slaves) == 0 {
		slaves = append(slaves, master)
	}

	conn := redisConn{
		master: master,
		slaves: slaves,
		name:   redisname,
	}
	return context.WithValue(ctx, redisname, conn)
}

func GetWriteConn(ctx context.Context, redisname string) rlib.Conn {
	rconns := ctx.Value(redisname).(redisConn)
	return rconns.master.Get()
}

func GetReadConn(ctx context.Context, redisname string) rlib.Conn {
	rconns := ctx.Value(redisname).(redisConn)
	// Get connection from slaves randomly
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pool := rconns.slaves[r.Intn(len(rconns.slaves))]
	return pool.Get()
}

func GetAllConn(ctx context.Context, key string) []rlib.Conn {
	rconns := ctx.Value(key).(redisConn)
	allConns := make([]rlib.Conn, 0, len(rconns.slaves)+1)
	allConns = append(allConns, rconns.master.Get())
	for _, pool := range rconns.slaves {
		allConns = append(allConns, pool.Get())
	}
	return allConns
}

func Close(ctx context.Context, key string) context.Context {
	redises := GetAllConn(ctx, key)
	for _, redis := range redises {
		if err := redis.Close(); err != nil {
			log.Println("failed to close redis server:", err)
		}
	}
	return context.WithValue(ctx, key, nil)
}

func newPool(host string, port, db int) *rlib.Pool {
	return &rlib.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (rlib.Conn, error) {
			c, err := rlib.Dial("tcp", fmt.Sprintf("%s:%d", host, port), rlib.DialDatabase(db))
			if err != nil {
				log.Println("failed to dial redis server:", err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c rlib.Conn, t time.Time) error {
			_, err := c.Do("PING")
			log.Println("redis server connection error:", err)
			return err
		},
	}
}
