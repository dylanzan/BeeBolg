package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {

	//判断是否退出
	isExit:=this.Input().Get("exit")=="true"

	if isExit{
		this.Ctx.SetCookie("uname","",-1,"/") //清空cookie
		this.Ctx.SetCookie("pwd","",-1,"/")  //清空cookie
		this.Redirect("/",301) //重定向到首页
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("password") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<32 - 1
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	this.Redirect("/", 301)

	return
}

func checkAccount(ctx *context.Context) bool { //新版本的需要引入 "github.com/astaxie/beego/context" context.Context()
	ck, err := ctx.Request.Cookie("uname")

	if err != nil {
		log.Println("checkAccount ck request err is ",err )
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")

	pwd := ck.Value

	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("password") == pwd
}