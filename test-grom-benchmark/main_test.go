package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn = "root:yLKGHEZ+haja6Nnboa5KygoDeA=@tcp(127.0.0.1:3306)/gorm_benachmark?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
)

type Person struct {
	Id        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

func (p Person) TableName() string {
	return "person"
}

func Benchmark(b *testing.B) {
	limits := []int{
		10,
		100,
		1000,
		10000,
	}

	// 原生 database/sql 连接池
	sqlDB, _ := sql.Open("mysql", dsn)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetMaxIdleConns(100)

	// jmoiron/sqlx 连接池
	sqlxDB, _ := sqlx.Connect("mysql", dsn)
	sqlxDB.SetMaxOpenConns(500)
	sqlxDB.SetMaxIdleConns(100)

	// gorm.io/gorm
	gormDB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		DisableAutomaticPing:   true,
	})
	db, _ := gormDB.DB()
	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(100)

	for _, lim := range limits {
		// Benchmark database/sql
		b.Run(fmt.Sprintf("database/sql limit:%d", lim), func(b *testing.B) {
			var res []Person
			for i := 0; i < b.N; i++ {
				rows, err := sqlDB.Query("SELECT id,first_name,last_name,email FROM person LIMIT ?", lim)
				if err != nil {
					b.Fatal(err)
				}
				defer rows.Close()

				for rows.Next() {
					p := Person{}
					err := rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Email)
					if err != nil {
						b.Fatal(err)
					}
					res = append(res, p)
				}
			}
			_ = res
		})

		// Benchmark jmoiron/sqlx
		b.Run(fmt.Sprintf("jmoiron/sqlx limit:%d", lim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var res []Person
				err := sqlxDB.Select(&res, "SELECT id,first_name,last_name,email FROM person LIMIT ?", lim)
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		// Benchmark gorm.io/gorm
		b.Run(fmt.Sprintf("gorm.io/gorm limit:%d", lim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var res []Person
                //err := gormDB.Raw("SELECT id,first_name,last_name,email FROM person LIMIT ?", lim).Scan(&res).Error
				 err := gormDB.Model(Person{}).Limit(lim).Find(&res).Error
				 //err := gormDB.Limit(lim).Find(&res).Error
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		fmt.Println("#######################################")
	}
}

