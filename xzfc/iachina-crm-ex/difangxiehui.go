// package main
// 导出机构类型为地方协会的数据
package main

import (
	"database/sql"
	"fmt"
	"github.com/xuri/excelize/v2"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Type2Id int = 28
)

type company struct {
	id                int64
	name              sql.NullString
	creditCode        sql.NullString
	jian              sql.NullString
	en                sql.NullString
	registeredCapital sql.NullString // 注册资金
	regAddress        sql.NullString // 注册地
	lpName            sql.NullString // 法人姓名
	foundDate         sql.NullString // 注册时间
	tel               sql.NullString // 电话
	address           sql.NullString // 办公地址
}

type xlsxFile struct {
	f    *excelize.File
	name string
}

func connectMySQL() *sql.DB {
	dsn := "root:yLKGHEZ+haja6Nnboa5KygoDeA=@tcp(127.0.0.1:3306)/iachina_crm"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("[open mysql failed]", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("[ping mysql failed]", err)
		panic(err)
	}

	return db
}

func queryRowsCount(db *sql.DB) int {
	var count int
	sqlStr := "select count(id) from company where type_2_id=? and is_valid=0"
	err := db.QueryRow(sqlStr, Type2Id).Scan(&count)
	fmt.Println("[count]", count)
	if err != nil {
		fmt.Println("[query count failed]", err)
		return count
	}
	return count
}

// queryRows 查询机构类型是地方协会的数据 type_2_id = 20
func queryRows(db *sql.DB) *sql.Rows {
	sqlStr := "select id, name, credit_code, jian, en, registered_capital, reg_address, lp_name, found_date, tel, address from company where type_2_id=? and is_valid=0"
	rows, err := db.Query(sqlStr, Type2Id)
	if err != nil {
		fmt.Println("[query failed]", err)
	}
	return rows
}

func queryRowsForMySql() (*sql.Rows, int) {
	db := connectMySQL()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	count := queryRowsCount(db)
	if count == 0 {
		fmt.Println("没有符合条件的数据")
		return nil, count
	}

	rows := queryRows(db)
	return rows, count
}

func (x *xlsxFile) setHeader() {
	tableHeader := []interface{}{
		"系统机构代码", "机构代码", "统一社会信用代码", "机构中文全称", "机构中文简称", "机构英文全称", "机构英文简称",
		"披露机构类型", "中资外资类型", "经济性质或类型", "注册资本(万元)", "经营区域范围", "业务范围", "注册地",
		"曾用名", "法定代表人", "币种", "成立时间", "机构电话", "机构地址",
	}
	// 把表头写进去
	err := x.f.SetSheetRow("Sheet1", "A1", &tableHeader)
	if err != nil {
		fmt.Println("[set table header failed]", err)
		return
	}
}

func (x *xlsxFile) save() {
	if err := x.f.SaveAs(x.name); err != nil {
		panic(err)
	}
}

func (x *xlsxFile) setData(rows *sql.Rows) {
	// 设置表头
	x.setHeader()

	// 设置数据
	key := 0
	for rows.Next() {
		var c company

		// 数据写到结构体中
		err := rows.Scan(
			&c.id, &c.name, &c.creditCode, &c.jian, &c.en, &c.registeredCapital, &c.regAddress, &c.lpName, &c.foundDate, &c.tel, &c.address,
		)
		if err != nil {
			fmt.Println("[scan failed]", err)
			return
		}

		// 构造数据，数据顺序要和表头对应
		s := []interface{}{
			c.id, "", c.creditCode.String, c.name.String,
			c.jian.String, c.en.String, "", "", "", "", c.registeredCapital.String,
			"", "", c.regAddress.String, "", c.lpName.String, "", c.foundDate.String,
			c.tel.String, c.address.String,
		}
		// 要写入的行
		axis := fmt.Sprintf("A%d", key+2)
		// 写入一行数据
		err = x.f.SetSheetRow("Sheet1", axis, &s)
		if err != nil {
			fmt.Printf("[set row data err] err: %s, cell: %s, id: %d\n", err, axis, c.id)
			return
		}
		key++
	}

	x.save()
}

func main() {
	rows, count := queryRowsForMySql()
	if count == 0 {
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	fileName := "地方协会.xlsx"
	xlsx := xlsxFile{f: excelize.NewFile(), name: fileName}
	xlsx.setData(rows)
}
