package controllers

import (
	"BeeBlog/src/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController)Get(){

	this.Data["IsLogin"]=checkAccount(this.Ctx)
	op:=this.Input().Get("op")

	switch op {
	case "add":
		name:=this.Input().Get("name")
		if len(name)==0{
			break
		}

		err:=models.AddCategory(name)
		if err!=nil{
			beego.Error(err)
		}
		this.Redirect("/category",301)
		return
	case "del":
		id:=this.Input().Get("id")
		if len(id)==0 {
			break
		}

		err:=models.DelCategory(id)

		if err!=nil{
			beego.Error(err)
		}
	}
	
	this.TplName="category.html"
	this.Data["IsCategory"]=true

	defer func() {
		categorysErr:=recover()
		if categorysErr!=nil{
			beego.Error(categorysErr)
		}
	}()

	//var getAllErr error
	this.Data["Categories"],_=models.GetAllCategorys()

	return
}