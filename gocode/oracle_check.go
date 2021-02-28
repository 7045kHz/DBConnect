package main

import (
	"database/sql"
	"fmt"
	"./jpw"

	_ "github.com/godror/godror"
)

func main() {

	dbConfigFile := "./dblogin.json"
        // Gets json saved connect strings using databases/functions.go file
        ORCLconnectString := jpw.GetConnectString(dbConfigFile,"ORCL_OSR_osreporter")
	// db, err := sql.Open("godror", "osreporter/ka1ckt1999@OSR")
	db, err := sql.Open("godror", ORCLconnectString)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select fqdn,package,version,arch,summary from packages")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var vfqdn string
	var vpackage string
	var vversion string
	var varch string
	var vsummary string
	for rows.Next() {

		rows.Scan(&vfqdn, &vpackage, &vversion, &varch, &vsummary)
		fmt.Printf("%s,%s,%s,%s,%s\n", vfqdn, vpackage, vversion, varch, vsummary)
	}
}
