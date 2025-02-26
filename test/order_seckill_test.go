package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	"golang-example/db"

	_ "github.com/go-sql-driver/mysql"
)

// http://localhost:9000/seckill?product_id=1&user_id=1
// 模拟 600 人并发抢秒杀的测试用例
func TestSeckillConcurrency(t *testing.T) {
	// 模拟 600 个并发请求
	const concurrency = 10000
	var wg sync.WaitGroup
	successCount := 0
	failureCount := 0
	productID := 5

	// 开始计时
	startTime := time.Now()

	// 启动 600 个 goroutine 并发请求
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			// 发起秒杀请求
			resp, err := http.Get(fmt.Sprintf("http://localhost:9000/seckill?product_id=%d&user_id=%d", productID, userID))
			if err != nil {
				t.Errorf("请求出错: %v", err)
				failureCount++
				return
			}
			defer resp.Body.Close()

			// 读取响应内容
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("读取响应内容出错: %v", err)
				failureCount++
				return
			}

			// 解析响应 JSON
			var result map[string]string
			err = json.Unmarshal(body, &result)
			if err != nil {
				t.Errorf("解析 JSON 出错: %v", err)
				failureCount++
				return
			}

			// 判断请求是否成功
			if resp.StatusCode == http.StatusOK && len(result["message"]) != 0 {
				successCount++
			} else {
				failureCount++
			}
		}(i + 1)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 计算总耗时
	elapsedTime := time.Since(startTime)

	// 输出测试结果
	fmt.Printf("总请求数: %d\n", concurrency)
	fmt.Printf("成功请求数: %d\n", successCount)
	fmt.Printf("失败请求数: %d\n", failureCount)
	fmt.Printf("总耗时: %s\n", elapsedTime)

	// 延时一段时间，等待队列中的请求处理完成
	time.Sleep(15 * time.Second)

	// 验证数据库中的库存和订单数量
	dbConn, err := db.InitDB()
	if err != nil {
		t.Errorf("数据库连接出错: %v", err)
		return
	}
	defer dbConn.Close()

	// 查询商品库存
	var stock int
	err = dbConn.QueryRow(fmt.Sprintf("SELECT stock FROM products WHERE id = %d", productID)).Scan(&stock)
	if err != nil {
		t.Errorf("查询商品库存出错: %v", err)
		return
	}

	// 查询订单数量
	var orderCount int
	err = dbConn.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM orders WHERE product_id = %d", productID)).Scan(&orderCount)
	if err != nil {
		t.Errorf("查询订单数量出错: %v", err)
		return
	}

	// 验证库存和订单数量是否符合预期
	if stock != successCount {
		t.Errorf("商品库存不符合预期，期望: %d，实际: %d", stock, successCount)
	}
	// if orderCount != successCount {
	// 	t.Errorf("订单数量不符合预期，期望: %d，实际: %d", successCount, orderCount)
	// }
}
