package controllers

import "github.com/astaxie/beego"

type YjsTest struct {
	beego.Controller
}

func (this *YjsTest) Get()  {
	this.Ctx.WriteString("hell yjs")
	this.GetBool("key1")
	this.GetFloat("key")
	this.GetStrings("key")

}
