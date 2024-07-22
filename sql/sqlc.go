// package main
// 使用 sqlc 操作数据库
package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sql/db"
)

func main() {
	dsn := "root:yLKGHEZ+haja6Nnboa5KygoDeA=@tcp(127.0.0.1:3306)/example_db"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queries := db.New(conn)
	createData(queries)
}

func createData(queries *db.Queries) {
	ctx := context.Background()

	_, err := queries.CreateDepartment(ctx, "Computer science")
	if err != nil {
		log.Fatal(err)
	}

	_, err = queries.CreateInstructor(
		ctx,
		db.CreateInstructorParams{Name: "John Doe", DeptID: sql.NullInt32{Int32: 1, Valid: true}},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func queryData(queries *db.Queries) {
	ctx := context.Background()
	students, err := queries.GetStudents(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, student := range students {
		log.Printf("Student ID: %d, Name: %s, Department ID: %d\n", student.ID, student.Name, student.DeptID.Int32)
	}

	courses, err := queries.GetCoursesByDept(ctx, sql.NullInt32{Int32: 1, Valid: true})
	if err != nil {
		log.Fatal(err)
	}
	for _, course := range courses {
		fmt.Printf("Course ID: %d, Title: %s, Department ID: %d\n", course.ID, course.Title, course.DeptID.Int32)
	}
}
