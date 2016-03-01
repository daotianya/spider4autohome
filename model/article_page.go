package model

import "github.com/astaxie/beego/orm"

type ArticlePage struct {
	Id            int    `orm:"column(id)"`
	ArticleUrlMd5 string `orm:"column(articleUrlMd5)"`
	PageNo        string `orm:"column(pageNo)"`
	PageContent   string `orm:"column(pageContent)"`
	PageLink      string `orm:"column(pageLink)"`
	Urlmd5        string `orm:"column(urlmd5)"`
	UpdateTime    string `orm:"column(updateTime)"`
}

func init() {
	orm.RegisterModel(new(ArticlePage))
}

func (self *ArticlePage) TableName() string {
	return "article_page"
}
