// package main
// 使用 orm 操作数据库
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type Department struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
}

type Instructor struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:100;not null"`
	DeptID uint
	Dept   Department `gorm:"foreignKey:DeptID"`
}

type Course struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"size:100;not null"`
	DeptID uint
	Dept   Department `gorm:"foreignKey:DeptID"`
}

type Student struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:100;not null"`
	DeptID uint
	Dept   Department `gorm:"foreignKey:DeptID"`
}

type Enrollment struct {
	ID        uint `gorm:"primaryKey"`
	StudentID uint
	CourseID  uint
	Semester  string  `gorm:"size:50;not null"`
	Year      int     `gorm:"not null"`
	Student   Student `gorm:"foreignKey:StudentID"`
	Course    Course  `gorm:"foreignKey:CourseID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	dsn := "root:yLKGHEZ+haja6Nnboa5KygoDeA=@tcp(127.0.0.1:3306)/example_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		log.Fatal(err)
	}
	//err = db.AutoMigrate(&Enrollment{})
	//if err != nil {
	//	log.Fatal(err)
	//}

	log.Println("Connected to the databases successfully!")

	// createData(db)
	queryData(db)
	updateData(db)
	queryData(db)
}

func createData(db *gorm.DB) {
	// 创建院系
	cs := Department{Name: "Computer Science"}
	math := Department{Name: "Mathematics"}
	db.Create(&cs)
	db.Create(&math)

	// 创建教师
	db.Create(&Instructor{Name: "John Do", DeptID: cs.ID})
	db.Create(&Instructor{Name: "Jane Smith", DeptID: math.ID})

	// 创建课程
	db.Create(&Course{Title: "Database Systems", DeptID: cs.ID})
	db.Create(&Course{Title: "Calculus", DeptID: math.ID})

	// 创建学生
	db.Create(&Student{Name: "Alice", DeptID: cs.ID})
	db.Create(&Student{Name: "Bob", DeptID: math.ID})

	// 学生选课
	db.Create(&Enrollment{StudentID: 1, CourseID: 1, Semester: "Fall", Year: 2024})
	db.Create(&Enrollment{StudentID: 2, CourseID: 2, Semester: "Fall", Year: 2024})
}

func queryData(db *gorm.DB) {
	log.Println("query data")
	var students []Student
	db.Find(&students)
	for _, student := range students {
		log.Printf("student ID: %d, Name: %s, department ID: %d\n", student.ID, student.Name, student.DeptID)
	}

	var courses []Course
	db.Where("id = ?", 1).Find(&courses)
	for _, course := range courses {
		log.Printf("course ID: %d, Title: %s, department ID: %d\n", course.ID, course.Title, course.DeptID)
	}
}

func updateData(db *gorm.DB) {
	db.Model(&Student{}).Where("id=?", 1).Update("name", "Alice Johnson")
}
