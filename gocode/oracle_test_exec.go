package main

import (
	"fmt"
	"./jpw"
	"os/exec"
)

func main() {

	sqlFile := "@./exec_sql.sql"
	dbConfigFile := "./dblogin.json"
        // Gets json saved connect strings using databases/functions.go file
        ORCLconnectString := jpw.GetCMDString(dbConfigFile,"ORCL_APEX_devops")
	// db, err := sql.Open("godror", "osreporter/ka1ckt1999@OSR")
	fmt.Printf("C = %s\n", ORCLconnectString)
	output,err:=exec.Command("sqlplus",ORCLconnectString,sqlFile).Output()

	if err != nil {
		fmt.Println("Stderr: ", err, " has issues opening\n")
	}
	fmt.Printf("Output %s\n",output)

}
