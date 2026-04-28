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
		if _, err := tm.AddTaskByFunc("func", "@every 1s", func() {
			time.Sleep(time.Second)
			fmt.Println("1s...")
		}, "测试"); err != nil {
			fmt.Println("err==================", err)
		}
	}),
)
