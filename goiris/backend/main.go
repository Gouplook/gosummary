package main

import (
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gosummary/goiris/backend/web/controllers"
	"gosummary/goiris/common"
	"gosummary/goiris/repositories"
	"gosummary/goiris/services"
)

func main() {
	//1.创建iris 实例
	app := iris.New()
	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")

	//3.注册模板
	// D:\gosummaryCode\gosummary\goiris\backend\web\views
	tmplate := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	//4.设置模板目标
	app.HandleDir("/assets", "./backend/web/assets")

	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
		// D:\gosummaryCode\gosummary\goiris\backend\web\views\shared
	})

	//连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		logs.Error(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//5.注册控制器
	productRepository := repositories.NewProductManger("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	//6.启动服务
	app.Run(
		iris.Addr("localhost:8083"))

}
