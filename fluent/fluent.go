package fluent

import (
	"log"
	"time"

	flib "github.com/fluent/fluent-logger-golang/fluent"
	"github.com/satoshi03/go/config"
	"golang.org/x/net/context"
)

func Open(ctx context.Context, fconf config.Fluent, fluentname string) context.Context {
	// connection to fluent
	f, err := flib.New(flib.Config{FluentPort: fconf.Port, FluentHost: fconf.Host})
	if err != nil {
		panic(err)
	}
	return context.WithValue(ctx, fluentname, f)
}

func Send(ctx context.Context, key, tag string, data map[string]interface{}) {
	f := ctx.Value(key).(*flib.Fluent)
	data["created_at"] = time.Now().Format("2006-01-02 15:04:05 -0700")
	if err := f.Post(tag, data); err != nil {
		log.Println(err)
	}
}

func Close(ctx context.Context, key string) context.Context {
	f := ctx.Value(key).(*flib.Fluent)
	if err := f.Close(); err != nil {
		log.Println(err)
	}
	return context.WithValue(ctx, key, nil)
}
