package main

import (
	"context"
	"fmt"
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

	//4.设置模板目标

	//出现异常跳转到指定页面

	//连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		logs.Error(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(db)
	fmt.Println(ctx)

	//5.注册控制器
	productRepository := repositories.NewProductManger("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	//6.启动服务
	app.Run(
		iris.Addr("localhost:8080"))

}
