package jpw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Database struct {
	Connection          string `json:"Connection"`
	DB            string `json:"dbType"`
	Description   string `json:"Description"`
	ConnectString string `json:"ConnectString"`
}

type Databases struct {
	Databases []Database `json:"Databases`
}
/* Example Call
x := GetConnectString("MYSQL_devops_devops")
fmt.Println(x)
*/

/*
		Example JSON File for DB passwords
		[
		  {
	        "Title": "DB_Connection",
	        "Description": "some info about connection",
	        "dbType": "DB Type",
	        "ConnectString": "db connect string"
	      },
		  {
	        "Title": "ORCL_OSR_osreporter",
	        "dbType": "ORCL",
	        "Description": "Connects to OC.RSYSLAB.COM OSR Oracle DB as user osreporter",
	        "ConnectString": "user='osreporter' password='ka1ckt1999' connectString='oc.rsyslab.com:1521/OSR'"
	 	   }
	  	]
*/

var dbConfigFile string

func GetConnectString(dbConfigFile string, ConnectName string) string {
	var ConnectString string
	var databases Databases

	dbList, err := ioutil.ReadFile(dbConfigFile)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(dbList, &databases)
	if err != nil {
		fmt.Println("error:", err)
	}
	for i := 0; i < len(databases.Databases); i++ {
		if databases.Databases[i].Connection == ConnectName {
			//fmt.Println("Name: " + databases.Databases[i].Name)
			//fmt.Println("DB: " + databases.Databases[i].DB)
			//fmt.Println("Description: " + databases.Databases[i].Description)
			//fmt.Println("Connect String: " + databases.Databases[i].ConnectString)
			ConnectString = databases.Databases[i].ConnectString
		}
	}
	if len(ConnectString) < 1 {
		ConnectString = "MISSING"
	}

	return ConnectString
}
