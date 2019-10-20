package controllers

import (
	"BeeBlog/src/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

//添加评论
func (this *ReplyController)Add(){
	tid:=this.Input().Get("tid")
	nickname:=this.Input().Get("nickname")
	content:=this.Input().Get("content")

	err:=models.AddReply(tid,nickname,content)

	if err!=nil{
		beego.Error(err)
	}

	this.Redirect("/topic/view/",302)

}

//删除评论
func (this *ReplyController)Delete()  {


	tid:=this.Input().Get("tid")
	err:=models.DeleteReply(this.Input().Get("rid"))

	if err!=nil{
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid,302)
}