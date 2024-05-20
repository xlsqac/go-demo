package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type icAccountField struct {
    inner_name string
}

func queryAllField(db *sql.DB) []string {
    sqlStr := "select distinct `inner_name` from `ic_account_field`"
    rows, err := db.Query(sqlStr)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    var fields []string
    var iaf icAccountField
    for rows.Next() {
        err := rows.Scan(&iaf.inner_name)
        if err != nil {
            panic(err)
        }

        fields = append(fields, iaf.inner_name)

    }
    return fields
}

func queryIcField(db *sql.DB, icCode string) []string {
    sqlStr := fmt.Sprintf("select distinct `inner_name` from `ic_account_field` where `ic_code`='%s'", icCode)
    rows, err := db.Query(sqlStr)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    var fields []string
    var iaf icAccountField
    for rows.Next() {
        err := rows.Scan(&iaf.inner_name)
        if err != nil {
            panic(err)
        }

        fields = append(fields, iaf.inner_name)

    }
    return fields

}
func diffOneWay(ss1, ss2raw []string) (result []string) {
    set := make(map[string]struct{}, len(ss1))

    for _, s := range ss1 {
        set[s] = struct{}{}
    }

    for _, s := range ss2raw {
        if _, ok := set[s]; ok {
            delete(set, s) // remove duplicates
        } else {
            result = append(result, s)
        }
    }
    return
}

func main() {
    dsn := "root:yLKGHEZ+haja6Nnboa5KygoDeA=@tcp(127.0.0.1:3306)/yobee_task"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err)

    }
    defer db.Close()

    allField := queryAllField(db)

    icCode := "pingan"
    icField := queryIcField(db, icCode)

    fmt.Printf("allField: %v\n", allField)
    fmt.Printf("icField: %v\n", icField)
    added := diffOneWay(icField, allField)
    fmt.Printf("added %v\n", added)

}
