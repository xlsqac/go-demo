// package main
// 探究在 go 中操作 mysql
// 参考资料：https://mp.weixin.qq.com/s/uUqWzabm6yoKuz3rhAvrpA
// 使用标准库中的 sql

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func insertDate(db *sql.DB) {
	_, err := db.Exec("INSERT INTO department (name) VALUES ('Computer Science'), ('Mathematics')")
	if err != nil {
		log.Fatal(err)
	}

	// 插入instructor数据
	_, err = db.Exec("INSERT INTO instructor (name, dept_id) VALUES ('John Doe', 1), ('Jane Smith', 2)")
	if err != nil {
		log.Fatal(err)
	}

	// 插入course数据
	_, err = db.Exec("INSERT INTO course (title, dept_id) VALUES ('Database Systems', 1), ('Calculus', 2)")
	if err != nil {
		log.Fatal(err)
	}

	// 插入student数据
	_, err = db.Exec("INSERT INTO student (name, dept_id) VALUES ('Alice', 1), ('Bob', 2)")
	if err != nil {
		log.Fatal(err)
	}

	// 插入enrollment数据
	_, err = db.Exec("INSERT INTO enrollment (student_id, course_id, semester, year) VALUES (1, 1, 'Fall', 2024), (2, 2, 'Fall', 2024)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted successfully!")
}

func queryData(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var studentID int
		var name string
		var deptID int
		err := rows.Scan(&studentID, &name, &deptID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Student ID: %d，Name: %s, Dept ID: %d\n", studentID, name, deptID)
	}
}

func updateData(db *sql.DB) {
	_, err := db.Exec("UPDATE student SET name = 'Alice Johnson' WHERE student_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Data updated successfully!")
}

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	dsn := "root:yLKGHEZ+haja6Nnboa5KygoDeA=@tcp(127.0.0.1:3306)/example_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the databases successfully!")

	//insertDate(db)
	queryData(db)

}
