package sqlx_

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func TestSqlite3EaseUsage() {
	sqlDB, _ := sql.Open("sqlite3", "./test.db")
	db := sqlx.NewDb(sqlDB, "sqlite3")
	//! 先这样放着，避免后面有问题
	res, err := db.Exec(`Create Table if not exists users(
		id integer primary key AUTOINCREMENT,
		username varchar(20),
		password varchar(20)
	)`)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
	//\ 使用 log.Fatalln 会导致进行死掉
	log.Print("裂开了啊")
}
