package main

import (
	"database/sql"
	"fmt"
	"log"
        "./jpw"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/godror/godror"
)

var Connection string
var ConnectString string
//var databases Databases


func main() {

	dbConfigFile := "./dblogin.json"
	// Gets json saved connect strings using databases/functions.go file
	ORCLconnectString := jpw.GetConnectString(dbConfigFile,"ORCL_SOR_devops")
	fmt.Println("ORCLconnectString: ", ORCLconnectString)
	SQLSrvconnectString := jpw.GetConnectString(dbConfigFile,"SQLSRV_msdb_devops")
	fmt.Println("SQLSrvconnectString: ", SQLSrvconnectString)

	// Setup Oracle Connection
	// db, err := sql.Open("godror", "osreporter/ka1ckt1999@OSR")
	ORCLdb, err := sql.Open("godror", ORCLconnectString)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ORCLdb.Close()

	// Setup MS SqlServer connection
	// URL connection string formats
	//    sqlserver://sa:mypass@localhost?database=master&connection+timeout=30         // username=sa, password=mypass.
	//    sqlserver://sa:my%7Bpass@somehost?connection+timeout=30                       // password is "my{pass"
	// note: pwd is "myP@55w0rd"

	db, err := sql.Open("mssql", SQLSrvconnectString)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	} else {
	  println("open connection")
	}

	//MS_Insert := "INSERT INTO INVENTORY (fqdn,icon,chassis,machine_id,boot_id,virtualization,os,cpe,kernel,arch,environment) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');"
	// Get data from Oracle
	SQLString := "select fqdn,icon,chassis,machine_id,boot_id,virtualization,os,cpe,kernel,arch,environment from SOREPORT.INVENTORY"

	ORCLrows, err := ORCLdb.Query(SQLString)
	if err != nil {
		fmt.Println("Oracle Error running query")
		fmt.Println(err)
		return
	}
	defer ORCLrows.Close()
	var vfqdn string
	var vicon string
	var vchassis string
	var vmachineid string
	var vbootid string
	var vvirtualization string
	var vos string
	var vcpe string
	var vkernel string
	var varch string
	var venvironment string
	var tsql string
	_, err = db.Exec("truncate table inventory")
	if err != nil {
		fmt.Println("SQLserver Error running query")
		fmt.Println(err)
		return
	}
	for ORCLrows.Next() {
		ORCLrows.Scan(&vfqdn, &vicon, &vchassis, &vmachineid, &vbootid, &vvirtualization, &vos, &vcpe, &vkernel, &varch, &venvironment)
		tsql = fmt.Sprintf("INSERT INTO INVENTORY (fqdn,icon,chassis,machine_id,boot_id,virtualization,os,cpe,kernel,arch,environment) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", vfqdn, vicon, vchassis, vmachineid, vbootid, vvirtualization, vos, vcpe, vkernel, varch, venvironment)
		_, err = db.Exec(tsql)
		if err != nil {
			fmt.Println("SQLserver Error running query")
			fmt.Println(err)
			return
		}
	}
}
