package main

// 根据表结构生成结构体
// https://blog.csdn.net/shida219/article/details/115338873

import (
	"fmt"

	"github.com/gohouse/converter"
)

var dsn = "root:12369@tcp(127.0.0.1:3306)/zproject?parseTime=True&loc=Local"
var savePath = "sql/zproject/z_entity/user.go"

func main() {
	err := converter.NewTable2Struct().
		SavePath(savePath).
		Dsn(dsn).TagKey("gorm").
		EnableJsonTag(true).
		Table("user").
		Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
