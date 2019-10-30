package routers

import (
	"PPGo_ApiAdmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
								//对象							对象的方法    "*:Index"：*表示任意方法都行该函数，
	beego.Router("/", &controllers.WwwController{}, "*:Index")	//读完
						//正则路由，匹配/show/12;变量id等于2
	beego.Router("/show/:id", &controllers.WwwController{}, "*:Show")	//读完
	beego.Router("/list/:class_id", &controllers.WwwController{}, "*:List")	//读完

	//post请求时，提交信息不为空时，根据username查询用户信息，返回结构体数据，数据库中记录的password和用户提交过来的进行比对，提交过来的password经过和随机字符串加密得到的；
	//验证通过后，设置用户cookie,并未使用session;
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")		//读完

	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")		//读完

	//一个简单的警告无权限页面
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")		//读完，

	beego.Router("/home", &controllers.HomeController{}, "*:Index")		//读完
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")	//读完

	//展示数据库中古诗词的内容，首次是展示所有数据
	beego.Router("/news/list", &controllers.NewsController{}, "*:List")		//读完   资讯列表
	beego.Router("/news/edit", &controllers.NewsController{}, "*:Edit")		//读完

	beego.Router("/ads/index", &controllers.AdsController{}, "*:Index")			//读完
	beego.Router("/ads/show", &controllers.AdsController{}, "*:Show")	//读完
	beego.Router("/ads/image_show", &controllers.AdsController{}, "*:ImageShow")	//读完

	beego.AutoRouter(&controllers.ApiController{})	//
	beego.AutoRouter(&controllers.ApiDocController{})
	beego.AutoRouter(&controllers.ApiMonitorController{})
	beego.AutoRouter(&controllers.EnvController{})
	beego.AutoRouter(&controllers.CodeController{})


	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.NewsController{})
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/yjs",&controllers.YjsTest{})
}
