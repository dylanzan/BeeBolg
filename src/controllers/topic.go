package controllers

import (
	"BeeBlog/src/models"
	"github.com/astaxie/beego"
	"log"
	"path"
)

type TopicController struct {
	beego.Controller
}


func (this *TopicController)Get() {

	this.Data["IsLogin"]=checkAccount(this.Ctx)
	this.TplName="topic.html"
	this.Data["IsTopic"]=true

	topics,err:=models.GetAllTopic("",false) //获取所有文章列表

	if err!=nil{
		beego.Error(err.Error())
	}else {
		this.Data["Topics"]=topics
	}

}

//添加文章
func (this *TopicController)Post(){
	if !checkAccount(this.Ctx){
		this.Redirect("/login",302)
		return
	}

	//解析表单
	title:=this.Input().Get("title")
	content:=this.Input().Get("content")
	tid:=this.Input().Get("tid")
	category:=this.Input().Get("category")
	label:=this.Input().Get("label")


	//判断用户是否上传附件
 	_,fh,err:=this.GetFile("attachment")

 	if err!=nil{
 		beego.Error(err)
	}

 	var attachment string

 	if fh!=nil{
 		attachment=fh.Filename
 		beego.Info(attachment)
 		err=this.SaveToFile("attachment",path.Join("attachment",attachment))
 		//filename:tmp.go
 		//attachment/tmp.go

 		if err!=nil{
 			beego.Error(err)
		}
	}


	//var err error

	if len(tid)==0{
		err=models.AddTopic(title,content,label,category,attachment)
	}else {
		err = models.ModifyTopic(tid,title,label,content,category,attachment)
	}

	if err!=nil{
		beego.Error(err)
	}

	this.Redirect("/topic",302)

}

//添加文章
func (this *TopicController) Add(){
	this.Data["IsLogin"]=checkAccount(this.Ctx)
	this.TplName="topic_add.html"
}

//查看文章

func (this *TopicController) View(){
	this.TplName="topic_view.html"

	topic,err:=models.GetTopicById(this.Ctx.Input.Param("0")) //获取url中的参数 //http://localhost:8080/topic/views/{"0"}
	if err!=nil{
		beego.Error(err)
		this.Redirect("/",302)
		return
	}

	this.Data["Topic"]=topic

	this.Data["Tid"]=this.Ctx.Input.Param("0")

	replies,err:=models.GetAllReplies(this.Ctx.Input.Param("0"))

	if err!=nil{
		beego.Error(err)
		return
	}

	this.Data["Replies"]=replies
	//log.Println(replies[1])
	this.Data["IsLogin"]=checkAccount(this.Ctx)



}

//修改文章
func (this *TopicController)Modify(){
	this.TplName="topic_modify.html"
	tid:=this.Input().Get("tid")

	log.Println("tid is ",tid)

	topic,err:=models.GetTopicById(tid)

	if err!=nil{
		beego.Error(err)
		this.Redirect("/",302)
		return
	}

	this.Data["Topic"]=topic
	this.Data["Tid"]=tid

}

//删除文章
func (this *TopicController)Delete(){

	if !checkAccount(this.Ctx){
		this.Redirect("/login",302)
		return
	}

	tid:=this.Input().Get("tid")
	if !checkAccount(this.Ctx){ //如果没有登录，则先跳转到登录界面
		this.Redirect("/login",302)
		return
	}

	//tidNum,err:=strconv.ParseInt(tid,10,64)

/*	if err!=nil{
		beego.Error(err)
		return
	}*/

	err:=models.DeleteTopic(tid)

	if err!=nil{
		beego.Error(err)
		//this.Redirect("/",302)
		return
	}
	this.Redirect("/topic",302)

}
