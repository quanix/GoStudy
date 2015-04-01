package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {

	c.Data["IsTopic"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplNames = "topic.html"

	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
}

//文章添加
func (c *TopicController) Add() {
	c.TplNames = "topic_add.html"
}

//添加或更新
func (c *TopicController) Post() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	// 解析表单
	tid := c.Input().Get("tid")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	lable := c.Input().Get("lable")

	category := c.Input().Get("category")

	var err error
	//err = models.AddTopic(title, content)
	if len(tid) == 0 {
		err = models.AddTopic(title, category, lable, content)
	} else {
		err = models.ModifyTopic(tid, title, category, lable, content)
	}
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)
}

//文章浏览
func (c *TopicController) View() {
	c.TplNames = "topic_view.html"

	tid := c.Ctx.Input.Params["0"]

	topic, err := models.GetTopic(c.Ctx.Input.Params["0"])
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Lables"] = strings.Split(topic.Lables, " ")

	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Replies"] = replies
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}

func (this *TopicController) Modify() {
	this.TplNames = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}
