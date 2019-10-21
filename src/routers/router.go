package routers

import (
	"BeeBlog/src/controllers"
	"github.com/astaxie/beego"
	"os"
)

func init() {
	beego.Router("/", &controllers.MainController{}) //home
	beego.Router("/login",&controllers.LoginController{}) //login
	beego.Router("/category",&controllers.CategoryController{}) //category
	beego.Router("/topic/",&controllers.TopicController{}) //topic
	beego.Router("/topic/reply",&controllers.ReplyController{}) //reply
	beego.Router("/reply/add",&controllers.ReplyController{},"post:Add") //添加评论
	beego.Router("/reply/delete",&controllers.ReplyController{},"get:Delete") //删除评论


	//创建附件目录
	os.Mkdir("attachment",os.ModePerm)
	//作为静态文件
	//beego.SetStaticPath("/attachment","attachment")

	//作为控制器处理文件
	beego.Router("/attachment/:all",&controllers.AttachController{})

	beego.AutoRouter(&controllers.TopicController{}) //自动路由
}
