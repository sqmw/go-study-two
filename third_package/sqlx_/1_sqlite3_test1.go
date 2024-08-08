package sqlx_

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // 这个是数据库驱动相关的，必须要有不然不能进行操作
)

func Sqlite3Test1() {
	var db *sqlx.DB

	//\ exactly the same as the built-in
	//\ 用来打开一个数据库连接，这里连接的是 内存
	//\ 正常情况下，这里打开的事 可执行文件 对应的目录下面的一个叫做 test.db 的文件，但是这里进行了处理
	db, _ = sqlx.Connect("sqlite3", "./test.db")

	// from a pre-existing sql.DB; note the required driverName
	buildInDB, _ := sql.Open("sqlite3", "./test.db")
	db = sqlx.NewDb(buildInDB, "sqlite3")

	// force a connection and test that it worked
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}
}
