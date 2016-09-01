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

