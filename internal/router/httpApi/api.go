package httpApi

import (
	"eye/internal/service"
	"github.com/julienschmidt/httprouter"
)

func Register(router *httprouter.Router){
	// 设置路由方式一
	router.Handle("GET","/1",service.GetHello)

	//设置路由方式二
	router.GET("/",service.GetHello)
}
