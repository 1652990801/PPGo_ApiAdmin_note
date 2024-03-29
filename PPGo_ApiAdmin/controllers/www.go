/**********************************************
** @Des: 文章
** @Author: wangsy
** @Date:   2017-12-09 14:17:37
***********************************************/
package controllers

import (
	"PPGo_ApiAdmin/models"
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"math/rand"
)

type WwwController struct {
	BaseController		//这里又是重写了基础控制器，做了一层封装；
}
//所有方法都执行该函数，get。post,delete
func (self *WwwController) Index() {

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	filters = append(filters, "class_id", 5)//filters是一个切片，类型为空接口，目前又数值 0 1 5
	result, _ := models.NewsGetList(1, 6, filters...)
	//fmt.Println("-------------------------------------------------------------------------------")
	//fmt.Println("result:",result)
	//result: [0xc0426a4180 0xc0426a4240 0xc0426a4300 0xc0426a43c0 0xc0426a4480 0xc0426a4540]  返回的是InfoList对象
	list := make([]map[string]interface{}, len(result))	//list是一个map key是字符串，value是空接口

	//下面这个循环是，取出所有信息，返回的是InfoList对象，将每个对象的信息，生成map信息，每个map存储到list切片中，格式为["0":"['id':'0','title':'故事']]
	for k, v := range result {		//循环取出来的数据
		row := make(map[string]interface{})		//将信息存储到 row中
		row["id"] = v.Id
		row["title"] = v.Title
		row["class_id"] = v.ClassId

		if(string(v.Picurl) == "") {		//如果图片地址为空
			var r = rand.Intn(10)	//返回一个随机数
			v.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"	//组合成一个图片路径
		}
		row["picurl"] = v.Picurl
		row["media"] = v.Media

		if (v.Desc != "") {
			nameRune := []rune(v.Desc)
			lth := len(nameRune)
			if(lth > 30) {
				lth = 30
			}
			row["desc"] = string(nameRune[:lth])
		}

		row["linkurl"] = v.Linkurl
		row["author"] = v.Author
		list[k] = row
	}
	//fmt.Println("------------------------------------------------------------------")
	//fmt.Println("list内容：",list)
	//fmt.Println("list[1]:",list[1])
	//fmt.Println("------------------------------------------------------------------")

	filters2 := make([]interface{}, 0)
	filters2 = append(filters2, "status", 1)
	filters2 = append(filters2, "class_id", 3)
	result2, _ := models.NewsGetList(1, 6, filters2...)
	list2 := make([]map[string]interface{}, len(result2))
	for k, v := range result2 {
		row2 := make(map[string]interface{})
		row2["id"] = v.Id
		row2["title"] = v.Title
		row2["class_id"] = v.ClassId
		if(string(v.Picurl) == "") {
			var r = rand.Intn(10)
			v.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row2["picurl"] = v.Picurl
		row2["media"] = v.Media

		if (v.Desc != "") {
			nameRune := []rune(v.Desc)
			lth := len(nameRune)
			if(lth > 30) {
				lth = 30
			}
			row2["desc"] = string(nameRune[:lth])
		}

		row2["linkurl"] = v.Linkurl
		row2["author"] = v.Author
		list2[k] = row2
	}
	//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	//fmt.Println("list2:",list2)
	//fmt.Println("list2[1]:",list2[1])
	//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	//上面两端的逻辑都是从数据库返回一个切片，切片中存储的是结构体对象(内存地址),一个切片，然后循环切片获取对象内容，重组后再次存储到一个切片中,切片数据类型是map

	//国学经典
	filters3 := make([]interface{}, 0)
	filters3 = append(filters3, "status", 1)
	filters3 = append(filters3, "class_id", 1)
	result3, _ := models.NewsGetList(1, 16, filters3...)//1，16，[0,1,1]
	list3 := make([]map[string]interface{}, len(result3))	//list3 = ["":"","":""]
	for k, v := range result3 {
		row2 := make(map[string]interface{})
		row2["id"] = v.Id
		row2["title"] = v.Title
		row2["class_id"] = v.ClassId
		if(string(v.Picurl) == "") {
			var r = rand.Intn(10)
			v.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row2["picurl"] = v.Picurl
		row2["media"] = v.Media

		if (v.Desc != "") {
			nameRune := []rune(v.Desc)
			lth := len(nameRune)
			if(lth > 30) {
				lth = 30
			}
			row2["desc"] = string(nameRune[:lth])
		}

		row2["linkurl"] = v.Linkurl
		row2["author"] = v.Author
		list3[k] = row2
	}

	//诗词古韵
	filters4 := make([]interface{}, 0)
	filters4 = append(filters4, "status", 1)
	filters4 = append(filters4, "class_id", 2)
	result4, _ := models.NewsGetList(1, 6, filters4...)
	list4 := make([]map[string]interface{}, len(result4))
	for k, v := range result4 {
		row2 := make(map[string]interface{})
		row2["id"] = v.Id
		row2["title"] = v.Title
		row2["class_id"] = v.ClassId
		if(string(v.Picurl) == "") {
			var r = rand.Intn(10)
			v.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row2["picurl"] = v.Picurl
		row2["media"] = v.Media

		if (v.Desc != "") {
			nameRune := []rune(v.Desc)
			lth := len(nameRune)
			if(lth > 30) {
				lth = 30
			}
			row2["desc"] = string(nameRune[:lth])
		}

		row2["linkurl"] = v.Linkurl
		row2["author"] = v.Author
		list4[k] = row2
	}


	out := make(map[string]interface{})		//out=map ["string":"interface"]
	out["list"] = list		//["list":"map[]"]
	out["list2"] = list2
	out["list3"] = list3
	out["list4"] = list4
	out["class_id"] = 0
	self.Data["data"] = out
	//fmt.Println("==============================================================")
	//fmt.Println("out = ",out)
	//fmt.Println("out[list2] = ",out["list2"])
	//fmt.Println("==============================================================")


	self.Layout = "public/www_layout.html"
	fmt.Println("开始执行display函数")
	self.display()	//返回一个this.tpLanem = self.controllerName + "/" + self.actionName + ".html"	和上面进行合并渲染
}

//访问/show/整数执行该函数
func (self *WwwController) Show() {
	fmt.Println("执行show函数：")

	id, _ := self.GetInt(":id")		//获取url中的变量
	News, _ := models.NewsGetById(id)		//根据id返回古诗
	NewsNext, _ := models.NewsGetNextById(id)	//返回下一首古诗

	nextRow := make(map[string]interface{})

	if(NewsNext != nil ){
		nextRow["id"] = NewsNext.Id
		nextRow["title"] = NewsNext.Title
	}
	row := make(map[string]interface{})
	row["class_id"] = 0
	if (News != nil) {
		row["id"] = News.Id
		row["title"] = News.Title
		row["class_id"] = News.ClassId
		row["desc"] = News.Desc
		row["content"] = News.Content
		if(string(News.Picurl) == "") {
			var r = rand.Intn(10)
			News.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row["picurl"] = News.Picurl
		row["linkurl"] = News.Linkurl
		row["media"] = News.Media
		row["author"] = News.Author
		row["posttime"] = beego.Date(time.Unix(News.Posttime, 0), "Y/m/d")
	}
	fmt.Println(row)
	row["next"] = nextRow
	self.Data["data"] = row
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("row:",row)
	fmt.Println("------------------------------------------------------------")

	self.Layout = "public/www_layout.html"
	self.display()
}
//访问/list/整数，时执行该函数
//首页的导航，都会引导到这里，
func (self *WwwController) List() {

	page, err := self.GetInt("page")
	fmt.Println("page:",page,"err:",err)
	tmp,_ := self.GetInt(":class_id")
	fmt.Println("class_id:",tmp)

	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 16
	}
	catId, cerr_ := self.GetInt(":class_id")
	fmt.Println(catId)
	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if cerr_ == nil {
		filters = append(filters, "class_id", catId)
	}
	result, count := models.NewsGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["class_id"] = v.ClassId

		if(string(v.Picurl) == "") {
			var r = rand.Intn(10)
			v.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row["picurl"] = v.Picurl
		row["media"] = v.Media
		if (v.Desc != "") {
			nameRune := []rune(v.Desc)
			lth := len(nameRune)
			if(lth > 30) {
				lth = 30
			}
			row["desc"] = string(nameRune[:lth])
		}


		row["linkurl"] = v.Linkurl
		row["author"] = v.Author
		row["posttime"] = beego.Date(time.Unix(v.Posttime, 0), "Y-m-d")
		list[k] = row
	}

	classArr := make(map[int]string)
	classArr[5] = "开心儿歌"
	classArr[3] = "儿童古诗"
	classArr[2] = "诗词古韵"
	classArr[1] = "经典国学1111"

	out := make(map[string]interface{})
	out["count"] = count
	out["class_id"] = catId
	out["page"] = page
	out["class_name"] = classArr[catId]
	out["title"] = classArr[catId]
	out["list"] = list
	self.Data["data"] = out

	self.Layout = "public/www_layout.html"
	self.display()		//先渲染该模板，然后返回给Layout
}


