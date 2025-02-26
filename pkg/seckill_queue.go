package pkg

import (
	"context"
	"encoding/json"
	"golang-example/db"
	"log"
	"time"
)

// AddSeckillRequestToQueue 将秒杀请求添加到 Redis 队列
func AddSeckillRequestToQueue(productID, userID string) (string, int, error) {
	req := map[string]string{
		"product_id": productID,
		"user_id":    userID,
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return "Failed to marshal request", 500, err
	}

	err = db.RedisClient.RPush(context.Background(), "seckill_queue", string(reqJSON)).Err()
	if err != nil {
		return "Failed to add request to queue", 500, err
	}

	return "Request added to queue", 200, nil
}

// 从 Redis 队列中取出请求并处理
func ProcessSeckillQueue() {
	ctx := context.Background()
	// 记录队列开始处理的时间
	startTime := time.Now()

	for {
		// 尝试从队列中取出元素
		result, err := db.RedisClient.BLPop(ctx, 0, "seckill_queue").Result()
		if err != nil {
			// 检查队列长度，如果队列为空，说明消费完成
			queueLen, lenErr := db.RedisClient.LLen(ctx, "seckill_queue").Result()
			if lenErr == nil && queueLen == 0 {
				break
			}
			log.Printf("Failed to pop from Redis queue: %v", err)
			continue
		}

		var req map[string]string
		err = json.Unmarshal([]byte(result[1]), &req)
		if err != nil {
			log.Printf("Failed to unmarshal request: %v", err)
			continue
		}

		productID := req["product_id"]
		userID := req["user_id"]

		_, _, _ = Seckill(productID, userID)
	}

	// 记录队列处理结束的时间
	endTime := time.Now()
	// 计算队列消费的耗时
	elapsedTime := endTime.Sub(startTime)

	log.Printf("Redis 队列消费完的耗时: %s", elapsedTime)
}
