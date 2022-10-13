// package main

// import "github.com/gin-gonic/gin"

// type Person struct {
// 	ID   string `uri:"id" binding:"required"`
// 	Name string `uri:"name" binding:"required"`
// }

// func main() {
// 	r := gin.Default()

// 	// http://localhost:8081/hello
// 	r.GET("/hello", func(ctx *gin.Context) {
// 		ctx.Writer.Write([]byte("hello world"))
// 	})

// 	// http://localhost:8081/uri/hello/123
// 	r.GET("/uri/:name/:id", func(c *gin.Context) {
// 		var persion Person
// 		err := c.ShouldBindUri(&persion)
// 		if err != nil {
// 			c.JSON(400, gin.H{"msg": err})
// 			return
// 		}
// 		c.JSON(200, gin.H{"name": persion.Name, "id": persion.ID})
// 	})

// 	r.Run(":8081") // 监听并在 0.0.0.0:8081 上启动服务
// }
