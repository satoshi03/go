# Common utility package for Golang



## Redis

### Configuration

```.yml
redis:
    r1:
        host: 127.0.0.1
        port: 6379
        db: 1
    r2:
        host: 127.0.0.1
        port: 6379
        db: 2
        slaves:
            - host: 127.0.0.1
              port: 6379
              db: 3
            - host: 127.0.0.1
              port: 6379
              db: 4
```

### Initialize

```.go
import (
    "golang.org/x/net/context"

    "github.com/satoshi03/go/redis"
)

func main() {
	conf := config.Read("config.yml")
	// initialize context
	ctx := context.Background()

	// Redis
	ctx = redisContext(ctx, conf.Redis[redisname])
	defer redis.Close(ctx, redisname)

    // Do something
}

```

### GET

```.go
import (
    "golang.org/x/net/context"

    "github.com/satoshi03/go/redis"
)


def getSomeData(ctx context.Context, redisname string) string {
    key := "hogehoge"
    redis.GetCmd(ctx, redisname, key)
}
```

### GET

```.go
import (
    "golang.org/x/net/context"

    "github.com/satoshi03/go/redis"
)


def getSomeData(ctx context.Context, redisname string) string {
    key := "hogehoge"
    val := "piyopiyo"
    redis.SetCmd(ctx, redisname, key, val)
}
```



## Example

```.go
package main

import (
	"golang.org/x/net/context"

	"github.com/satoshi03/go/config"
	"github.com/satoshi03/go/fluent"
	"github.com/satoshi03/go/redis"
)

var redisname = "r1"
var fluentname = "f1"

func main() {
	conf := config.Read("config.yml")
	// initialize context
	ctx := context.Background()

	// Redis
	ctx = redisContext(ctx, conf.Redis[redisname])
	defer redis.Close(ctx, redisname)
	redis.SetCmd(ctx, redisname, "test", "test")
	r, _ := redis.GetCmd(ctx, redisname, "test")

	// Fluentd
	ctx = fluentContext(ctx, conf.Fluent)
	defer fluent.Close(ctx, fluentname)
	fluent.Send(ctx, fluentname, "debug.test", map[string]interface{}{"test": "test"})
}

func redisContext(ctx context.Context, conf config.Redis) context.Context {
	ctx = redis.Open(ctx, conf, redisname)
	return ctx
}

func fluentContext(ctx context.Context, conf config.Fluent) context.Context {
	ctx = fluent.Open(ctx, conf, fluentname)
	return ctx
}
```
