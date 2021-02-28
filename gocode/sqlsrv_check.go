package main

import (
	// Import go-mssqldb strictly for side-effects
	"database/sql"
	"fmt"
	"log"
	"./jpw"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	dbConfigFile := "./dblogin.json"
        SQLSrvconnectString := jpw.GetConnectString(dbConfigFile,"SQLSRV_msdb_devops")

	sql_query := "select fqdn,icon,chassis,machine_id,boot_id,virtualization,os,cpe,kernel,arch,environment from INVENTORY"

	println(sql.Drivers())

	// URL connection string formats
	//    sqlserver://sa:mypass@localhost?database=master&connection+timeout=30         // username=sa, password=mypass.
	//    sqlserver://sa:my%7Bpass@somehost?connection+timeout=30                       // password is "my{pass"
	// note: pwd is "myP@55w0rd"

	println("open connection")
	db, err := sql.Open("mssql", SQLSrvconnectString)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(sql_query)
	defer rows.Close()
	for rows.Next() {
		var (
			vfqdn           string
			vicon           string
			vchassis        string
			vmachine_id     string
			vboot_id        string
			vvirtualization string
			vos             string
			vcpe            string
			vkernel         string
			varch           string
			venvironment    string
		)

		rows.Scan(&vfqdn, &vicon, &vchassis, &vmachine_id, &vboot_id, &vvirtualization, &vos, &vcpe, &vkernel, &varch, &venvironment)
		fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n", vfqdn, vicon, vchassis, vmachine_id, vboot_id, vvirtualization, vos, vcpe, vkernel, varch, venvironment)
	}
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
