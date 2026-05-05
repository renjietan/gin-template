package fx_module

import (
	"example.com/t/core/component/redis_drivers"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
)

var FxRedisModule = fx.Module("fx-redis-module",
	fx.Provide(redis_driver.NewRedisDriver),
	fx.Invoke(func(r *redis.Client) {
		// 无过期时间
		//err := r.Set(context.Background(), strconv.Itoa(i), i, 0).Err()
		//if err != nil {
		//	return
		//}

		// 过期时间
		//err := r.SetEX(context.Background(), strconv.Itoa(i), i, 60*time.Second).Err()
		//if err != nil {
		//	return
		//}
	}),
)
