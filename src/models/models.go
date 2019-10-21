package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/unknwon/com"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	_DB_NAME       = "data/beego.db"
	_SQLITE3_DRIVE = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index""`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopiclastUserId int64
}

type Topic struct { //orm自动增加新加列
	Id               int64
	Uid              int64
	Title            string
	Category         string
	Labels           string
	Content          string `orm:"size(5000)"`
	Attaciment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int64     `orm:"index"`
	Autoher          string
	ReplyTime        time.Time `orm:"index"`
	ReplyCount       int64
	RepleyLastUserId int64
}

//评论
type Conment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

//注册数据库
func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(Conment)) //注册模型
	orm.RegisterDriver(_SQLITE3_DRIVE, orm.DRSqlite)           //注册驱动

	orm.RegisterDataBase("default", _SQLITE3_DRIVE, _DB_NAME, 10) //创建数据库 需要一个default数据库

}

//添加分类列表
func AddCategory(name string) error {
	o := orm.NewOrm() //获取os对象

	//categorySlice:=utils.HandleCategoryAddStr(categoryRes)
	//id,_:=strconv.ParseInt(categorySlice[0],10,64)

	cate := &Category{Title: name}

	qs := o.QueryTable("category")

	err := qs.Filter("title", name).One(cate)

	if err == nil { //如果等于nil,则代表存在，则结束，返回err
		return err
	}

	_, err = o.Insert(cate)

	if err != nil {
		return err
	}

	return nil
}

//删除分类
func DelCategory(id string) error {

	cid, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		beego.Error(err)
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}

	_, err = o.Delete(cate)

	return err

}

//获取分类列表
func GetAllCategorys() ([]*Category, error) {

	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")

	_, err := qs.All(&cates)

	return cates, err
}

func AddTopic(tile, content, label, category ,attachment string) error {

	//处理标签
	//空格作为多个标签和分隔符
	//beego #beego$
	//$bee#$beego#
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	//

	o := orm.NewOrm()

	topic := &Topic{
		Id:               0,
		Uid:              0,
		Title:            tile,
		Content:          content,
		Attaciment:       attachment,
		Labels:           label,
		Created:          time.Now(),
		Updated:          time.Now(),
		Views:            0,
		Autoher:          "clouder",
		ReplyTime:        time.Now(),
		ReplyCount:       0,
		RepleyLastUserId: 0,
		Category:         category,
	}

	_, err := o.Insert(topic)

	if err != nil {
		beego.Error(err)
		return err
	}

	//更新分类
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err != nil {
		//如果不存在，简单的忽略更新操作
		cate.TopicCount++
		_, err = o.Update(cate)
	}

	return err
}

//获取所有的文章列表
func GetAllTopic(cate string, isDesc bool) ([]*Topic, error) {

	o := orm.NewOrm()

	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error

	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		_, err := qs.OrderBy("-created").All(&topics) //排序 —倒叙
		return topics, err
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

//通过文章Id获取文章
func GetTopicById(id string) (*Topic, error) { //id string :url,form表单中大多数传进来的是string类型

	tidNum, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		beego.Error(err)
		//return nil, err
	}
	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")

	err = qs.Filter("id", tidNum).One(topic)

	if err != nil {
		return nil, err
	}

	topic.Views++

	_, err = o.Update(topic)

	//将处理好的label,替换为以空格为空格符的字符串格式
	topic.Labels=strings.Replace(strings.Replace(topic.Labels,"#"," ",-1),"$","",-1)

	/*if err!=nil{
		return nil,err
	}*/

	return topic, err
}

func ModifyTopic(tid, titile, label,content, category,attachment string) error {

	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	tidNum, err := strconv.ParseInt(tid, 10, 64)

	if err != nil {
		beego.Error(err)
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}

	var oldCate,oldAttach string
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttach=topic.Attaciment
		topic.Title = titile
		topic.Content = content
		topic.Updated = time.Now()
		topic.Category = category
		topic.Labels=label
		topic.Attaciment=attachment

		_, err := o.Update(topic)

		if err != nil {
			beego.Error(err)
			return err
		}

		//删除附旧件
		if len(oldAttach)>0{
			os.Remove(path.Join("attachment",oldAttach))
		}


		cate := new(Category)
		qs := o.QueryTable("category")
		qs.Filter("Title", category).One(cate)

		if err != nil {
			cate.TopicCount++
			o.Update(cate)
		}
	}

	//更新分类
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(cate)

		if err != nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}

	return nil
}

//model 删除文章操作
func DeleteTopic(tid string) error {

	tidNum, err := strconv.ParseInt(tid, 10, 64)

	if err != nil {
		beego.Error(err)
		return err
	}

	var oldCate string
	o := orm.NewOrm()

	topic := &Topic{Id: tidNum}

	if o.Read(topic) == nil {
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}

	if len(oldCate) == 0 {
		cate := new(Category)
		qs := o.QueryTable("catefory")

		err = qs.Filter("title", oldCate).One(cate)

		if err != nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}

	//_, err = o.Delete(topic)

	return nil
}

//添加评论
func AddReply(tid, nickname, content string) error {

	tidNum, err := strconv.ParseInt(tid, 10, 64)

	if err != nil {
		beego.Error(err)
		return err
	}

	o := orm.NewOrm()

	conment := &Conment{
		Id:      0,
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}

	_, err = o.Insert(conment)

	if err != nil {
		beego.Error(err)
		return err
	}

	//统计回复数量
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}

	return err
}

//获取评论的内容
func GetAllReplies(tid string) (replies []*Conment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)

	if err != nil {
		beego.Error(err)
		return
	}

	replies = make([]*Conment, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("conment")

	_, err = qs.Filter("tid", tidNum).All(&replies)

	return replies, err
}

//delete reply content
func DeleteReply(rid string) error {

	ridNum, err := strconv.ParseInt(rid, 10, 64)

	if err != nil {
		beego.Error(err)
		return err
	}

	o := orm.NewOrm()

	var tidNum int64

	reply := &Conment{Id: ridNum}

	if o.Read(reply) == nil {
		tidNum = reply.Tid
		_, err = o.Delete(reply) //删除信息

		if err != nil {
			return err
		}
	}

	replies := make([]*Conment, 0)

	qs := o.QueryTable("comment")

	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)

	if err != nil {
		return err
	}

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = replies[0].Created
		topic.ReplyCount = int64(len(replies))
		_, err = o.Update(topic)
	}
	return err
}
