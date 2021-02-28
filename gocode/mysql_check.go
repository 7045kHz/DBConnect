package main

import (
    "database/sql"
    "fmt"
    "log"
    "./jpw"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    dbConfigFile := "./dblogin.json"
    MySQLconnectString := jpw.GetConnectString(dbConfigFile,"MYSQL_devops_devops")
    db, err := sql.Open("mysql",MySQLconnectString)
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    var version string

    err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println(version)
}
