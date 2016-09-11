# Common utility package for Golang



## Redis

### Initialize

```
import (
    "golang.org/x/net/context"

    "github.com/satoshi03/go/redis"
)


def initFunc(ctx context.Context) context.Context {
    // redis connection to context
    ctx :=  redis.Open(ctx, "127.0.0.1", "6379", "redis")
    defer redis.Close()

    return ctx
}

```

### GET

```

import (
    "golang.org/x/net/context"

    "github.com/satoshi03/go/redis"
)


def getSomeData(ctx context.Context) string {

    key := "hogehoge"
    cli := redis.GetCon(ctx,"redis")
    value, _ := redis.GetCmd(cli, key)
    return value
}

```

## Sample

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

