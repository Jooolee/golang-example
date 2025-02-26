package test

import (
	"fmt"
	"sync"
	"testing"
)

// 假设这是一个共享资源
var sharedResource int
var locker sync.Mutex

func increment() {
	// 获取锁，阻止其他 goroutine 同时访问共享资源
	locker.Lock()
	// 使用 defer 确保在函数返回时自动释放锁
	defer locker.Unlock()

	// 对共享资源进行操作
	sharedResource++
	fmt.Println("Incremented sharedResource to:", sharedResource)
}
func TestGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	// 启动多个 goroutine 并发执行 increment 函数
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("Final value of sharedResource:", sharedResource)
}
