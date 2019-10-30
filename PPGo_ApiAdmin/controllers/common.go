/**********************************************
** @Des: base controller
** @Author: haodaquan
** @Date:   2017-09-07 16:54:40
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-18 10:28:01
***********************************************/
package controllers

import (
	"PPGo_ApiAdmin/libs"
	"PPGo_ApiAdmin/models"
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"reflect"
	"strconv"
	"strings"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

//重新封装beego.controller的目的是为了重写prepare函数，被之后的所有继承的页面执行
//当一个用户认证通过后，这个用户的的权限信息等都被生成了
type BaseController struct {
	beego.Controller		//继承了beego的基础控制器，又新增了几个字段
	controllerName string	//当前执行的控制器名字
	actionName     string	//当前执行的方法名字
	user           *models.Admin	//这里的数据类型是引用，应用了modeles中的admin类，user接收的应该是内存地址，
	userId         int		//	用户的id
	userName       string	//用户的名字
	loginName      string	//
	pageSize       int
	allowUrl       string	//权限认证，允许访问的url
	noLayout	   bool		//权限认证，不允许访问的url,
}

//重写了prepare函数，在继承了baseController对象的所有方法中执行执行该方法；
//前期准备
//这里的self是请求自身
func (self *BaseController) Prepare() {
	self.pageSize = 20
	controllerName, actionName := self.GetControllerAndAction()	//beego自身的方法，返回当前控制器名字和操作方法

	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])	//转换为小写，并截取controllers之前的内容
	fmt.Println("控制器的名字：",controllerName)
	self.actionName = strings.ToLower(actionName)		//函数的名字或者说是
	fmt.Println("方法名字：",actionName)

	self.Data["version"] = beego.AppConfig.String("version")	//配置文件中没有这个参数呢
	fmt.Println("version:",self.Data["version"])

	//下面这些赋值的信息，可以在以后的get post方法中随意调用；
	self.Data["siteName"] = beego.AppConfig.String("site.name")	//站点名字；获取配置文件参数
	self.Data["curRoute"] = self.controllerName + "." + self.actionName	//控制器.方法；祖父穿组合控制器名字 + 方法名字；例如 www.show
	self.Data["curController"] = self.controllerName	//控制器名字
	self.Data["curAction"] = self.actionName		//方法名字；
	noAuth := "ads,wxApi,www"
	isNoAuth := strings.Contains(noAuth, self.controllerName)	//判断是否包含某个字符，也可以用bytes库判断；必须完全匹配noauth内容；

	//上面这些内容主要是为了初始化一些信息，存储到this.Data中；
	if isNoAuth == false {
		self.auth()		//如果不包含关键字，则执行auth函数；

	}

	self.Data["loginUserId"] = self.userId
	self.Data["loginUserName"] = self.userName
}

//该函数被Prepare函数调用执行，生成结构体属性信息；
//登录权限验证
func (self *BaseController) auth() {

	arr := strings.Split(self.Ctx.GetCookie("auth"), "|")	//得到一个cookie切片，[user password]
	self.userId = 0
	if len(arr) == 2 {	//如果cookie中包含2个元素则为真
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)	//字符串转int类型
		fmt.Println("arr:",arr,"\t idstr:",idstr,"\t,idstr类型：",reflect.TypeOf(idstr),"\t password:",password,"\t password类型：",reflect.TypeOf(password))
		if userId > 0 {
			user, err := models.AdminGetById(userId)		//通过cookie中存储的是用户的id; user返回的是用户信息结构体
			if err == nil && password == libs.Md5([]byte(self.getClientIp()+"|"+user.Password+user.Salt)) {
				self.userId = user.Id
				self.loginName = user.LoginName
				self.userName = user.RealName
				self.user = user
				self.AdminAuth()
			}

			isHasAuth := strings.Contains(self.allowUrl, self.controllerName+"/"+self.actionName)
			noAuth := "ajaxsave/ajaxdel/table/loginin/loginout/getnodes/start/show/ajaxapisave"
			isNoAuth := strings.Contains(noAuth, self.actionName)	//如果请求方法在noauth中，则拒绝执行
			if isHasAuth == false && isNoAuth == false {
				self.Ctx.WriteString("没有权限")
				self.ajaxMsg("没有权限", MSG_ERR)
				return
			}
		}
	}

	if self.userId == 0 && (self.controllerName != "login" && self.actionName != "loginin") {
		self.redirect(beego.URLFor("LoginController.LoginIn"))
	}

}

func (self *BaseController) AdminAuth() {
	// 左侧导航栏
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if self.userId != 1 {
		//普通管理员
		adminAuthIds, _ := models.RoleAuthGetByIds(self.user.RoleIds)
		adminAuthIdArr := strings.Split(adminAuthIds, ",")
		filters = append(filters, "id__in", adminAuthIdArr)
	}
	result, _ := models.AuthGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	list2 := make([]map[string]interface{}, len(result))
	allow_url := ""
	i, j := 0, 0
	for _, v := range result {
		if v.AuthUrl != " " || v.AuthUrl != "/" {
			allow_url += v.AuthUrl
		}
		row := make(map[string]interface{})
		if v.Pid == 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list[i] = row
			i++
		}
		if v.Pid != 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list2[j] = row
			j++
		}
	}

	self.Data["SideMenu1"] = list[:i]  //一级菜单
	self.Data["SideMenu2"] = list2[:j] //二级菜单

	self.allowUrl = allow_url + "/home/index"
}

// 是否POST提交
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"

}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 重定向
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}

//加载模板
func (self *BaseController) display(tpl ...string) {
	var tplname string
	fmt.Println("进入display，开始执行,len(tp1):",len(tpl))
	fmt.Println("tpl:",tpl)
	if len(tpl) > 0 {		//如果元素的数量大于0个
		fmt.Println("display方法执行,tp1内容：",tpl)
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {	//如果切片tpl中无数据，则执行这个html文件。
		//tplname = 当前的控制器名字例如home/start.hmtl
		tplname = self.controllerName + "/" + self.actionName + ".html"		//

	}

	if !self.noLayout {
		if self.Layout == "" {
			self.Layout = "public/layout.html"
		}
	}

	fmt.Println("当前的使用的页面是：",tplname)
	self.TplName = tplname
}

//ajax返回
func (self *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//ajax返回 列表
func (self *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//上传图片
func (self *BaseController) UploadFile(filename string, filepath string) {
	f, h, err := self.GetFile(filename)

	out := make(map[string]interface{})
	if err != nil {
		out["msg"] = "文件读取错误"
	}
	var fileSuffix, newFile string
	fileSuffix = path.Ext(h.Filename) //获取文件后缀
	newFile = libs.GetRandomString(8) + fileSuffix
	err = self.SaveToFile("upfile", filepath+newFile)
	if err != nil {
		out["msg"] = "文件保存错误"
	}
	defer f.Close()
	out["state"] = "SUCCESS"
	out["url"] = filepath + newFile
	out["title"] = newFile
	out["original"] = h.Filename
	out["size"] = h.Size
	out["msg"] = "ok"
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}
