package controllers

import (
	"BeeBlog/src/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "home.html"

	c.Data["IsLogin"] = checkAccount(c.Ctx) //检测request cookie中值，是否为真，真则将导航栏，登录改为退出，否则相反

	var err error
	topics,err:=models.GetAllTopic("",true) //获取全部文章

	if err!=nil{
		beego.Error(err)
	}else {
		c.Data["Topics"]=topics
	}

	categories,err:=models.GetAllCategorys()

	if err!=nil{
		beego.Error(err)
	}

	c.Data["Categories"]=categories

}
