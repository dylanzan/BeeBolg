package main

import (
	"BeeBlog/src/models"
	_ "BeeBlog/src/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init(){
	models.RegisterDB()
}

func main() {

	orm.Debug=true
	orm.RunSyncdb("default",false,true) //自动建表

	//beego.SetStaticPath("/static","static")
	//beego.SetStaticPath("/css","css")
	//beego.SetStaticPath("/js","js")
	beego.Run()

}

