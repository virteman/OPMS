package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func TableName(name string) string {
	return beego.AppConfig.String("mysqlpre") + name
}
func migirateTable() {
	//o := orm.NewOrm()
	//	tSql := "alter table `case` modify `story` varchar(400) NOT NULL"
	//tSql := "alter table `material` modify `name` varchar(50) NOT NULL"
	//tSql := "alter table `user` modify `password` varchar(100)  NULL"
	//o.Raw(tSql).Exec()
}
func Syncdb() {
	name := "default"
	// drop table 后再建表
	force := false
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	//initRole()
	//migirateTable()
}
