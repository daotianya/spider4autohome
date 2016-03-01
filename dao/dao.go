package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"spider/conf"
	"spider/model"
)

var o orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", conf.DB["user"], conf.DB["pass"], conf.DB["host"], conf.DB["port"], conf.DB["name"], conf.DB["char"])
	orm.RegisterDataBase("default", "mysql", dburl)
	o = orm.NewOrm()
}

//保存品牌
func SaveBrand(brand *model.Brand) (int64, error) {
	dbBrand := &model.Brand{Urlmd5: brand.Urlmd5}
	if o.Read(dbBrand, "Urlmd5") == orm.ErrNoRows { //纪录不存在
		return o.Insert(brand)
	} else {
		brand.Id = dbBrand.Id
		return o.Update(brand)
	}
}

func SaveSeries(series *model.Series) (int64, error) {
	dbSeries := &model.Series{Urlmd5: series.Urlmd5}
	err := o.Read(dbSeries, "Urlmd5")
	if err == orm.ErrNoRows { //纪录不存在
		return o.Insert(series)
	} else {
		series.Id = dbSeries.Id
		return o.Update(series, "SeriesId", "SeriesName", "SeriesLink", "SeriesImg", "SeriesImgLoc", "SeriesPrice", "ProducerName", "SourceURL", "UpdateTime")
	}
}

func UpdateSeries(series *model.Series) (int64, error) {
	dbSeries := &model.Series{Urlmd5: series.Urlmd5}
	err := o.Read(dbSeries, "Urlmd5")
	if err == nil {
		dbSeries.SeriesScore = series.SeriesScore
		dbSeries.SeriesBigImg = series.SeriesBigImg
		dbSeries.SeriesBigImgLoc = series.SeriesBigImgLoc
		dbSeries.SeriesEngine = series.SeriesEngine
		dbSeries.SeriesGearbox = series.SeriesGearbox
		dbSeries.SeriesCarStruct = series.SeriesCarStruct
		dbSeries.UpdateTime = series.UpdateTime
		return o.Update(dbSeries)
	}

	return int64(-1), err
}

func UpdateArticel(a *model.Article) (int64, error) {
	dbArticle := &model.Article{Urlmd5: a.Urlmd5}
	err := o.Read(dbArticle, "Urlmd5")
	if err == nil {
		dbArticle.ArticleTags = a.ArticleTags
		dbArticle.ArticleProvider = a.ArticleProvider
		dbArticle.ArticleType = a.ArticleType
		dbArticle.ArticleCategoryTag = a.ArticleCategoryTag
		dbArticle.ArticleContent = a.ArticleContent
		dbArticle.TotoalPages = a.TotoalPages
		dbArticle.UpdateTime = a.UpdateTime
		return o.Update(dbArticle)
	}

	return int64(-1), err
}

func SaveProducer(p *model.Producer) (int64, error) {
	dbProducer := &model.Producer{ProducerName: p.ProducerName}
	err := o.Read(dbProducer, "ProducerName")
	if err == orm.ErrNoRows { //纪录不存在
		return o.Insert(p)
	} else {
		p.Id = dbProducer.Id
		return o.Update(p)
	}
}

func SaveArticle(a *model.Article) (int64, error) {
	dbArticle := &model.Article{Urlmd5: a.Urlmd5}
	err := o.Read(dbArticle, "Urlmd5")

	if err == orm.ErrNoRows { //纪录不存在
		return o.Insert(a)
	} else {
		a.Id = dbArticle.Id
		return o.Update(a, "SeriesId", "SeriesUrlMd5", "ArticleLink", "ArticleTitle", "ArticleLogo", "ArticleLogoLoc", "ArticleAuthor", "ArticlePubDate", "ArticleViewCount", "ArticleCommentCount", "TotoalPages", "SourceURL", "UpdateTime")
	}
}

func SaveArticlePage(a *model.ArticlePage) (int64, error) {
	dbArticle := &model.ArticlePage{Urlmd5: a.Urlmd5}
	err := o.Read(dbArticle, "Urlmd5")

	if err == orm.ErrNoRows { //纪录不存在
		return o.Insert(a)
	} else {
		a.Id = dbArticle.Id
		return o.Update(a)
	}
}
