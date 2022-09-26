package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

// 通过Context的Param方法来获取API参数
func Api() {
	r := gin.Default()
	r.GET("/share/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	//默认为监听8080端口
	r.Run()
}

// URL参数可以通过DefaultQuery()或Query()方法获取
func UrlQuery() {
	r := gin.Default()
	r.GET("/share", func(c *gin.Context) {
		//指定默认值
		name := c.DefaultQuery("name", "司图")
		action := c.DefaultQuery("action", "技术分享会")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name+action))
	})
	r.Run()
}

// 表单传输为post请求
func FromPost() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	r.Run()
}

// 定义接收数据的结构体
//type Login struct {
//	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
//	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
//	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
//}
func JsonBind() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// JSON绑定
	r.POST("loginJSON", func(c *gin.Context) {
		// 声明接收的变量
		var json Login
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if json.User != "root0" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")

}

// 加载HTML
func HTML() {
	r := gin.Default()
	r.LoadHTMLGlob("D:\\gosummaryCode\\gosummary\\gogin\\views\\*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index2.html", gin.H{"ce": "8099"})
	})
	r.Run()
}

func Async() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 1.异步
	r.GET("/long_async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})
	// 2.同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})

	r.Run(":8000")
}

func MiddleLogic() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

// 全局中间件
func MiddleWare() {
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleLogic())
	// {}为了代码规范
	{
		r.GET("/share", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})

	}
	r.Run()
}

// 中间件案例
func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时：", since)
}
func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
func MiddleCase() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	//r.Use(myTime)
	// {}为了代码规范
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", myTime, shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	{
		shoppingGroup.Group("/share")
	}
	r.Run()

}

// cookie
func Cookie() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 服务端要给客户端cookie
	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			//cookie = "NotSet"
			// 给客户端设置cookie
			//  maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			//   secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})
	r.Run()
}

// 日志文件
func Logs() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("D:\\gosummaryCode\\gosummary\\gogin\\logs\\gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run()
}
func main() {
	//// 创建一个默认的路由引擎
	//r := gin.Default()
	//// 创建一个新的引擎
	////r := gin.New()
	////r.Use(StatCost())
	//
	//r.GET("/share", func(c *gin.Context) {
	//	var msg struct {
	//		Name    string `json:"name"`
	//		Message string
	//	}
	//	msg.Name = "司图部门"
	//	msg.Message = "司图部门技术分享"
	//	c.JSONP(http.StatusOK, msg)
	//})
	////requestString(r)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	//requestParam(r)
	//r.Run()
	//Api()
	//UrlQuery()
	//FromPost()

	//JsonBind()
	//HTML()
	//Async()
	//MiddleWare()
	//MiddleCase()
	//Cookie()
	//Logs()
}
