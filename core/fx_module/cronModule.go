package fx_module

import (
	"fmt"
	"time"

	cronManager "example.com/t/core/cron"
	"go.uber.org/fx"
)

var FXCronModule = fx.Module("fx-cron-module",
	fx.Provide(cronManager.NewTimerTask),
	fx.Invoke(func(tm *cronManager.CronManager) {
		var mockFunc = func() {
			time.Sleep(time.Second)
			fmt.Println("1s...")
		}
		if _, err := tm.AddTaskByFunc("func", "@every 1s", mockFunc, "测试"); err != nil {
			fmt.Println("err==================", err)
		}
	}),
)
