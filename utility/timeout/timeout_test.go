package utility

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimeoutManager(t *testing.T) {
	// 1. 创建管理器
	tm := NewTimeoutManager()
	defer tm.Close()

	// 2. 并发设置多个定时器
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			name := fmt.Sprintf("timer-%d", id)
			err := tm.SetSync(name,
				time.Second*time.Duration(id+1),
				fmt.Sprintf("定时器 %d", id),
				func() {
					t.Logf("函数回调 %d 在 %v 触发", id, time.Now())
				})

			if err != nil {
				t.Errorf("设置定时器 %d 失败: %v", id, err)
			}
		}(i)
	}

	wg.Wait()

	// 3. 检查定时器是否存在
	if !tm.Exists("timer-1") {
		t.Error("timer-1 应该存在")
	}

	// 4. 获取定时器信息
	info, err := tm.GetTimerInfo("timer-1")
	if err != nil {
		t.Errorf("获取定时器信息失败: %v", err)
	}
	t.Logf("定时器-1 信息: %+v", info)

	// 5. 获取剩余时间
	remaining, err := tm.GetRemainingTime("timer-1")
	if err != nil {
		t.Errorf("获取剩余时间失败: %v", err)
	}
	t.Logf("定时器-1 剩余时间: %v", remaining)

	// 6. 列出所有活跃定时器
	activeTimers := tm.ListActiveTimers()
	t.Logf("活跃定时器: %v", activeTimers)

	// 7. 获取统计信息
	active, created, expired, stopped, reset := tm.GetStats()
	t.Logf("统计信息: 活跃=%d, 已创建=%d, 已过期=%d, 已停止=%d, 已重置=%d",
		active, created, expired, stopped, reset)

	// 8. 重置定时器
	err = tm.ResetTimer("timer-2", 2*time.Second, "重置后的定时器", nil)
	if err != nil {
		t.Errorf("重置定时器失败: %v", err)
	}

	// 9. 使用扩展版本重置
	err = tm.ResetTimerExt("timer-3",
		WithDuration(3*time.Second),
		WithLogStr("扩展重置"),
	)
	if err != nil {
		t.Errorf("使用扩展方式重置定时器失败: %v", err)
	}

	// 10. 停止一个定时器
	stoppedSuccess := tm.StopByName("timer-0")
	if !stoppedSuccess {
		t.Error("停止定时器-0 失败")
	}

	// 11. 等待一段时间
	time.Sleep(4 * time.Second)

	// 12. 最终统计
	active, created, expired, stopped, reset = tm.GetStats()
	t.Logf("最终统计: 活跃=%d, 已创建=%d, 已过期=%d, 已停止=%d, 已重置=%d",
		active, created, expired, stopped, reset)

	// 13. 停止所有定时器
	tm.StopAll()

	// 14. 验证所有定时器已停止
	activeTimers = tm.ListActiveTimers()
	if len(activeTimers) > 0 {
		t.Errorf("期望没有活跃定时器，但得到 %v", activeTimers)
	}
}

func TestConcurrentAccess(t *testing.T) {
	tm := NewTimeoutManager()
	defer tm.Close()

	// 并发读写测试
	var wg sync.WaitGroup
	concurrency := 50

	// 并发写入
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			name := fmt.Sprintf("concurrent-timer-%d", id)
			tm.Set(name, time.Second, "并发测试", func() {
				// 空回调
			})
		}(i)
	}

	// 并发读取
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			name := fmt.Sprintf("concurrent-timer-%d", id)
			tm.Exists(name)
			tm.ListActiveTimers()
		}(i)
	}

	wg.Wait()

	// 验证没有panic发生
	t.Log("并发测试通过，未发生 panic")
}

func TestManagerClose(t *testing.T) {
	tm := NewTimeoutManager()

	// 设置一些定时器
	for i := 0; i < 3; i++ {
		tm.Set(fmt.Sprintf("timer-%d", i), 10*time.Second, "test", func() {})
	}

	// 关闭管理器
	tm.Close()

	// 验证关闭后操作返回错误
	err := tm.SetSync("new-timer", time.Second, "should fail", func() {})
	if err == nil {
		t.Error("关闭后 SetSync 应该失败")
	}

	if !tm.IsClosed() {
		t.Error("管理器应该已关闭")
	}

	// 多次关闭应该是安全的
	tm.Close()
	tm.Close()

	t.Log("关闭测试通过")
}
