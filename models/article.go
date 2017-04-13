package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"math"
	"time"
)

type JcContent struct {
	ContentId   int `orm:"pk"`
	ChannelId   int
	Title       string
	Comments    string
	Views       int
	UserId      int
	Username    string
	Description string
	TypeId      int
	UpdateAt    time.Time
	Status      int
	SortDate    time.Time
	IsRecommend int
}

type ArticleList struct {
	ContentId   int
	Title       string
	Comments    string
	Views       int
	UserId      int
	Username    string
	Description string
	UpdateAt    time.Time
}

type RecommendList struct {
	ContentId int
	Title     string
	SortDate  time.Time
}

type HotList struct {
	ContentId int
	Title     string
	ViewsDay  int
	UpdateAt  time.Time
}

type ReplyList struct {
	ContentId    int
	Title        string
	LastFeedback time.Time
}

func init() {
	orm.RegisterDataBase("default", "mysql", "devops:devops@tcp(192.168.71.117:3306)/system32?charset=utf8", 30)
	//orm.RegisterDataBase("default", "mysql", "root:@/beego?charset=utf8", 30)
	//orm.RegisterModel(new(JcContent))//不使用orm方式就可以不用注册model
	//orm.RegisterModel(new(JcContentCount))

	orm.Debug = true
	//orm.RunSyncdb("default", false, true)//同步数据表结构到数据库
}

func GetArticleList(channelId int, page int, pageSize int) (this []ArticleList, count int64, lastPage int) {
	o := orm.NewOrm()
	sql := "select content_id,title,comments,views,user_id,username,description,update_at from jc_content "
	sql += "where channel_id=? and type_id in(1,2,3,4) and status=2 order by content_id desc limit ?,?"

	var offset int
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}
	count, err := o.Raw(sql, channelId, offset, pageSize).QueryRows(&this)
	if err != nil {
		panic("数据异常")
	}

	lastPage = int(math.Ceil(float64(count) / float64(pageSize)))
	return this, count, lastPage
}

func GetRecommendList(channelId int) (this []RecommendList) {
	o := orm.NewOrm()
	sql := "select * from jc_content where channel_id=? and status=2 and is_recommend=1 and type_id in(1,3) order by sort_date desc limit 10"
	_, err := o.Raw(sql, channelId).QueryRows(&this)
	if err != nil {
		panic("数据异常")
	}
	return this
}

func GetHots(channelId int) (this []HotList) {
	o := orm.NewOrm()
	today := time.Now().Format("2006-01-02")
	sql := "select jc.content_id,jc.title,jcc.views_day,jc.update_at from jc_content jc left join jc_content_count jcc on jc.content_id = jcc.content_id  "
	sql += "where jc.channel_id=? and jc.type_id in(1,3) and jc.status=2 and jc.update_at>? order by jcc.views_day desc, jcc.content_id desc limit 10"
	_, err := o.Raw(sql, channelId, today).QueryRows(&this)
	if err != nil {
		panic("数据异常")
	}
	return this
}

func GetReply(channelId int) (this []ReplyList) {
	o := orm.NewOrm()
	sql := "select * from jc_content where channel_id=? and status=2 and is_recommend=1 and type_id in(1,3) order by last_feedback desc limit 10"
	_, err := o.Raw(sql, channelId).QueryRows(&this)
	if err != nil {
		panic("数据异常")
	}
	return this
}
