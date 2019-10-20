package routers

import (
	"BeeBlog/src/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}) //home
	beego.Router("/login",&controllers.LoginController{}) //login
	beego.Router("/category",&controllers.CategoryController{}) //category
	beego.Router("/topic/",&controllers.TopicController{}) //topic
	beego.Router("/topic/reply",&controllers.ReplyController{}) //reply
	beego.Router("/reply/add",&controllers.ReplyController{},"post:Add") //添加评论
	beego.Router("/reply/delete",&controllers.ReplyController{},"get:Delete") //删除评论

	beego.AutoRouter(&controllers.TopicController{}) //自动路由
}
