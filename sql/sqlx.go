// package main
// 探究在 go 中操作 mysql
// 参考资料：https://mp.weixin.qq.com/s/uUqWzabm6yoKuz3rhAvrpA
// 使用sqlx

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func insertData(db *sqlx.DB) {
	_, err := db.NamedExec(`INSERT INTO department (name) values (:name)`, []map[string]interface{}{
		{"name": "Computer science"},
		{"name": "Mathematics"},
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.NamedExec(`INSERT INTO department (name) values (:name)`, map[string]interface{}{
		"name": "Computer science1",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Insert Data successfully!")
}

func queryData(db *sqlx.DB) {

}

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	dsn := "root:yLKGHEZ+haja6Nnboa5KygoDeA=@tcp(127.0.0.1:3306)/example_db"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the databases successfully!")
	insertData(db)
}
