package model

import "github.com/astaxie/beego/orm"

type Article struct {
	Id                  int    `orm:"column(id)"`
	SeriesId            string `orm:"column(seriesId)"`
	SeriesUrlMd5        string `orm:"column(seriesUrlMd5)"`
	ArticleLink         string `orm:"column(articleLink)"`
	ArticleTitle        string `orm:"column(articleTitle)"`
	ArticleLogo         string `orm:"column(articleLogo)"`
	ArticleLogoLoc      string `orm:"column(articleLogoLoc)"`
	ArticleAuthor       string `orm:"column(articleAuthor)"`
	ArticlePubDate      string `orm:"column(articlePubDate)"`
	ArticleViewCount    string `orm:"column(articleViewCount)"`
	ArticleCommentCount string `orm:"column(articleCommentCount)"`
	ArticleTags         string `orm:"column(articleTags)"`
	ArticleProvider     string `orm:"column(articleProvider)"`
	ArticleType         string `orm:"column(articleType)"`
	ArticleCategoryTag  string `orm:"column(articleCategoryTag)"`
	ArticleContent      string `orm:"column(articleContent)"`
	Urlmd5              string `orm:"column(urlmd5)"`
	SourceURL           string `orm:"column(sourceURL)"`
	TotoalPages         int    `orm:"column(totoalPages)"`
	UpdateTime          string `orm:"column(updateTime)"`
}

func init() {
	orm.RegisterModel(new(Article))
}
