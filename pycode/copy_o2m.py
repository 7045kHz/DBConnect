#!/usr/bin/python3

import pymysql
import cx_Oracle
import csv
import os

# CSV prep
csv_tmp_file = '/tmp/INVENTORY.csv'
if os.path.exists(csv_tmp_file):
  print("DISCOVERED OLD {}, REMOVING".format(csv_tmp_file))
  os.remove(csv_tmp_file)

csv_file=open(csv_tmp_file,"w")
writer = csv.writer(csv_file, delimiter=',', lineterminator="\n", quoting=csv.QUOTE_NONNUMERIC)


# Oracle Connection
O_con = cx_Oracle.connect('user/passwd@fqdn.example.com:1521/sid')



O_ver = O_con.version.split(".")
O_cur = O_con.cursor()
O_cur.execute('select FQDN,ICON,CHASSIS,MACHINE_ID,BOOT_ID,VIRTUALIZATION,OS,CPE,KERNEL,ARCH,ENVIRONMENT from SOREPORT.INVENTORY')
O_res = O_cur.fetchall()
for o_row in O_res:
    #l = row[2]
    o_r = list()
    for o_c in o_row:
        if type(o_c) is cx_Oracle.LOB or cx_Oracle.CLOB:
         o_r.append(str(o_c))
        else:
            o_r.append(o_c)

    writer.writerow(o_r)

csv_file.close()
O_cur.close()
O_con.close()

# MySQL Connection
M_con = pymysql.connect("fqdn.example.com","user","passwd","db" )
M_cur = M_con.cursor()
csv_data = csv.reader(open(csv_tmp_file)) 
print("TRUNCATING INVENTORY TABLE ON DB DEVOPS")
M_sql_truncate = """TRUNCATE TABLE INVENTORY"""
M_cur.execute(M_sql_truncate)
M_con.commit()

print("INSERTING {} INTO INVENTORY table on DEVOPS".format(csv_tmp_file))

M_sql_prep = """INSERT INTO INVENTORY(FQDN,ICON,CHASSIS,MACHINE_ID,BOOT_ID,VIRTUALIZATION,OS,CPE,KERNEL,ARCH,ENVIRONMENT) VALUES(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)"""
for C_row in csv_data:
    C_row=[None if cell == '' else cell for cell in C_row]
    M_cur.execute(M_sql_prep,C_row)
M_con.commit()
M_cur.close()
M_con.close()

print("INSERT/COMMIT INVENTORY TABLE COMPLETE ON DB DEVOPS")
if os.path.exists(csv_tmp_file):
  print("REMOVING/CLEANING UP {}".format(csv_tmp_file))
  os.remove(csv_tmp_file)
print("UPLOAD/INSERT FROM {} INTO TABLE INVENTORY COMPLETE".format(csv_tmp_file))


