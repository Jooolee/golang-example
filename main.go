package main

import (
	"log"
	"net/http"

	"golang-example/controller"
	"golang-example/db"
	"golang-example/pkg"

	"github.com/gin-gonic/gin"
)

// https://gin-gonic.com/zh-cn/docs/quickstart/
// $ curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go

// var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// test status
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// 定义秒杀接口路由
	r.GET("/seckill", controller.SeckillQueueHandler)

	// // Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	// // Authorized group (uses gin.BasicAuth() middleware)
	// // Same than:
	// // authorized := r.Group("/")
	// // authorized.Use(gin.BasicAuth(gin.Credentials{
	// //	  "foo":  "bar",
	// //	  "manu": "123",
	// //}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	// /* example curl for /admin with basicauth header
	//    Zm9vOmJhcg== is base64("foo:bar")

	// 	curl -X POST \
	//   	http://localhost:8080/admin \
	//   	-H 'authorization: Basic Zm9vOmJhcg==' \
	//   	-H 'content-type: application/json' \
	//   	-d '{"value":"bar"}'
	// */
	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return r
}

func main() {
	db.InitRedis()
	if db.RedisClient == nil {
		log.Fatal("Failed to initialize Redis client")
	}
	r := setupRouter()

	go pkg.ProcessSeckillQueue()

	// 启动服务器
	if err := r.Run(":9000"); err != nil {
		panic(err)
	}
}
