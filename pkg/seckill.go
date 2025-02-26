package pkg

import (
	"database/sql"
	"golang-example/db"
)

// Seckill 处理商品秒杀业务逻辑
const maxRetries = 10

// Seckill 处理商品秒杀业务逻辑
func Seckill(productID, userID string) (string, int, error) {
	// 调用 db 包中的 InitDB 函数初始化数据库连接
	dbConn, err := db.InitDB()
	if err != nil {
		return "Database connection error", 500, err
	}
	defer dbConn.Close()

	for retry := 0; retry < maxRetries; retry++ {
		// 开启事务
		tx, err := dbConn.Begin()
		if err != nil {
			return "Transaction start error", 500, err
		}
		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()

		var (
			stock   int
			version int
		)

		// 查询商品信息和版本号
		err = tx.QueryRow("SELECT stock, version FROM products WHERE id =?", productID).Scan(&stock, &version)
		if err != nil {
			if err == sql.ErrNoRows {
				return "Product not found", 404, err
			}
			return "Failed to query product information", 500, err
		}

		if stock <= 0 {
			return "Out of stock", 400, nil
		}

		// 尝试减少库存并更新版本号
		result, err := tx.Exec("UPDATE products SET stock = stock - 1, version = version + 1 WHERE id =? AND version =?", productID, version)
		if err != nil {
			return "Failed to update stock", 500, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return "Failed to get rows affected", 500, err
		}

		if rowsAffected > 0 {
			// 创建订单
			_, err = tx.Exec("INSERT INTO orders (product_id, user_id) VALUES (?,?)", productID, userID)
			if err != nil {
				return "Failed to create order", 500, err
			}
			return "Seckill success", 200, nil
		}
	}

	return "Stock has been updated by another user, please try again", 409, nil
}
