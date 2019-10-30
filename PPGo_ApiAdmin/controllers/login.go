/**********************************************
** @Des: login
** @Author: haodaquan
** @Date:   2017-09-07 16:30:10
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:55:21
***********************************************/
package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"PPGo_ApiAdmin/libs"
	"PPGo_ApiAdmin/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

//登录 TODO:XSRF过滤
func (self *LoginController) LoginIn() {
	fmt.Println("创建新用户的密码：",libs.Md5([]byte("666666kmcB")))
	//INSERT INTO `pp_uc_admin` (`id`, `login_name`, `real_name`, `password`, `role_ids`, `phone`, `email`, `salt`, `last_login`, `last_ip`, `status`, `create_id`, `update_id`, `create_time`, `update_time`)
	//VALUES(5,'root','超管-叶藏','cc7b7761e6b1f7a956242f7efb552b64','0','18210954486','1652990801@qq.com','kmcB',1517993417,'[',1,0,0,0,1506128438);
	if self.userId > 0 {
		self.redirect(beego.URLFor("HomeController.Index"))
	}

	fmt.Println("beego.URLFOR地址：url:",beego.URLFor("HomeController.Index"))
	beego.ReadFromRequest(&self.Controller)
	if self.isPost() {		//如果是post请求，通过this.ctx.requese.mether获取请求方式，

		username := strings.TrimSpace(self.GetString("username"))		//获取用户信息，并去除两侧空格strings.Trimspace
		password := strings.TrimSpace(self.GetString("password"))

		if username != "" && password != "" {
			user, err := models.AdminGetByName(username)		//根据用户名字查询用户信息，返回用户信息结构体
			fmt.Println("user用户信息：",user)
			flash := beego.NewFlash()
			errorMsg := ""
			//查询用户的信息返回用户信息结构体，如果查询没有错误，则比对密码，
			//数据库中取出来的密码是经过md5加密的，加密使用的使用的是密码和随机字符串
			//user.Password  等于 password+user.Salt（随机字符串）
			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {		//查询没有错误，密码是MD5加密的，密码+随机字符串
				errorMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {	//验证正确，如果验证成功，则跳转道/home
				user.LastIp = self.getClientIp()	//client的IP地址
				user.LastLogin = time.Now().Unix()	//这里返回的是当前事件戳 ，例如1571983072
				user.Update()
				//ip|password|salt
				authkey := libs.Md5([]byte(self.getClientIp() + "|" + user.Password + user.Salt))

				//这里是设置的session = auth:1|key
				self.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
				//重定向
				fmt.Println("当前用户的cookie是：",self.Ctx.GetCookie("auth"))
				self.redirect(beego.URLFor("HomeController.Index"))
			}
			flash.Error(errorMsg)
			flash.Store(&self.Controller)
			self.redirect(beego.URLFor("LoginController.LoginIn"))
		}
	}
	self.TplName = "login/login.html"
}

//登出
func (self *LoginController) LoginOut() {
	fmt.Println("要登出的用户的cookie是：",self.Ctx.GetCookie("auth"))	//1|6ed4bb467de17bb9dd97bc8dd65a109b
	self.Ctx.SetCookie("auth", "")	//重新设置了cookie，为空
	self.redirect(beego.URLFor("LoginController.LoginIn"))
}

func (self *LoginController) NoAuth() {
	self.Ctx.WriteString("没有权限")
}
