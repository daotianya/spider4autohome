package model

import "github.com/astaxie/beego/orm"

type Producer struct {
	Id           int    `orm:"column(id)"`
	ProducerName string `orm:"column(producerName)"`
	UpdateTime   string `orm:"column(updateTime)"`
}

func init() {
	orm.RegisterModel(new(Producer))
}
