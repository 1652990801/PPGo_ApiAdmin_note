/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 10:21:13
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:04:41
***********************************************/
package controllers

import "fmt"

type HomeController struct {
	BaseController
}

func (self *HomeController) Index() {
	self.Data["pageTitle"] = "系统首页"	//view引用时语法是：{{.pageTitle}}
	//self.display()
	fmt.Println("/home路由时，携带的数据：",self.Data)
	//这里的this的数据，应该在parper时已经存储了很多数据，作为全局变量使用；
	fmt.Println("/home路由下的data数据：",self.Data)
	self.TplName = "public/main.html"		//直接返回了一个页面
}

func (self *HomeController) Start() {
	self.Data["pageTitle"] = "控制面板"
	fmt.Println("Start函数中的数据是self.Data：")

	//self.display()
	self.TplName = "home/start.html"
}

