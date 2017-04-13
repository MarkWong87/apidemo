package controllers

import (
	"apiproject/models"
	"github.com/astaxie/beego"
)

// Operations about Users
type ArticleController struct {
	beego.Controller
}

// @Title List
// @Description get all Article
// @Success 200 {object} models.GetArticleList
// @router /list [get]
func (this *ArticleController) List() {
	channelId, _ := this.GetInt("channelId")
	page, _ := this.GetInt("page")
	pageSize, _ := this.GetInt("pageSize")
	maps := make(map[string]interface{})
	maps["data"], maps["total"], maps["lastPage"] = models.GetArticleList(channelId, page, pageSize)
	maps["perPage"] = pageSize
	maps["currentPage"] = page
	this.Data["json"] = maps

	this.ServeJSON()
}

// @Title Recommend
// @Description get all Recommend
// @Success 200 {object} models.GetRecommendList
// @router /recommend [get]
func (this *ArticleController) Recommend() {
	channelId, _ := this.GetInt("channelId")
	this.Data["json"] = models.GetRecommendList(channelId)
	this.ServeJSON()
}

// @Title Hot
// @Description get all Hot
// @Success 200 {object} models.GetHots
// @router /hot [get]
func (this *ArticleController) Hot() {
	channelId, _ := this.GetInt("channelId")
	this.Data["json"] = models.GetHots(channelId)
	this.ServeJSON()
}

// @Title Reply
// @Description get all Reply
// @Success 200 {object} models.GetReply
// @router /reply [get]
func (this *ArticleController) Reply() {
	channelId, _ := this.GetInt("channelId")
	this.Data["json"] = models.GetReply(channelId)
	this.ServeJSON()
}
