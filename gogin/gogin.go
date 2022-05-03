package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 获取请求参数

//1: 获取querystring参数
// http://127.0.0.1:8080/user/search?username=%E5%B0%8F%E7%8E%8B%E5%AD%9044&address=%E6%B2%99%E6%B2%B311
func requestString(r *gin.Engine) {
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
}

//2: 获取form参数
func requestForm(r *gin.Engine) {
	r.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

}

//3: 获取json参数
func requestJosn(r *gin.Engine) {
	r.POST("/json", func(c *gin.Context) {
		// 注意：下面为了举例子方便，暂时忽略了错误处理
		b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
		// 定义map或结构体
		var m map[string]interface{}
		// 反序列化
		_ = json.Unmarshal(b, &m)

		c.JSON(http.StatusOK, m)
	})
}

// 4:Path
func requestPath(r *gin.Engine) {
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

}

// 5:参数绑定 兼容以上所有
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func requestParam(r *gin.Engine) {
	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginForm?user=q1mi&password=123456)
	//r.GET("/loginForm", func(c *gin.Context) {
	//	var login Login
	//	// ShouldBind()会根据请求的Content-Type自行选择绑定器
	//	if err := c.ShouldBind(&login); err == nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"user":     login.User,
	//			"password": login.Password,
	//		})
	//	} else {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	}
	//})

	// 绑定form表单示例 (user=q1mi&password=123456)
	//r.POST("/loginForm", func(c *gin.Context) {
	//	var login Login
	//	// ShouldBind()会根据请求的Content-Type自行选择绑定器
	//	if err := c.ShouldBind(&login); err == nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"user":     login.User,
	//			"password": login.Password,
	//		})
	//	} else {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	}
	//})
}

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	// 创建一个默认的路由引擎
	//r := gin.Default()
	// 创建一个新的引擎
	r := gin.New()
	r.Use(StatCost())
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	//r.GET("/hello", func(c *gin.Context) {
	//	// c.JSON：返回JSON格式的数据
	//	//c.JSON(200, gin.H{ // H是一个开箱即用的map
	//	//	"message": "Hello world!",
	//	//})
	//	// 方式一：自己拼接JSON
	//	//c.JSONP(http.StatusOK, gin.H{
	//	//	"messages": "Hello world !..",
	//	//})
	//	var msg struct {
	//		Name    string `json:"name"`
	//		Message string
	//		Age     int
	//	}
	//	msg.Name = "小王子"
	//	msg.Message = "Hello world!"
	//	msg.Age = 18
	//	c.JSONP(http.StatusOK, msg)
	//})
	//requestString(r)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	requestParam(r)

	r.Run()
}
