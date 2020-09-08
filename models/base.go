package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init(){
	//orm.RegisterModel(new(Genrialize))d
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
}
type Base struct{

}
