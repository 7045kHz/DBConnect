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

# MySQL Connection
M_con = pymysql.connect("fqdn.example.com","user","passwd","db" )
M_cur = M_con.cursor()

M_sql_prep = """select FQDN,ICON,CHASSIS,MACHINE_ID,BOOT_ID,VIRTUALIZATION,OS,CPE,KERNEL,ARCH,ENVIRONMENT from INVENTORY"""

M_cur.execute(M_sql_prep)
M_res = M_cur.fetchall()

for m_row in M_res:
    #l = row[2]
    m_r = list()
    for m_c in m_row:
        if type(m_c) is cx_Oracle.LOB or cx_Oracle.CLOB:
         m_r.append(str(m_c))
        else:
            m_r.append(m_c)

    writer.writerow(m_r)

csv_file.close()
M_cur.close()
M_con.close()
# Oracle Connection
O_con = cx_Oracle.connect('user/passwd@fqdn.example.com:1521/sid')
O_cur = O_con.cursor()


# Purge Oracle Table INVENTORY
print("TRUNCATING INVENTORY TABLE ON DB APEX")
O_sql_truncate = """TRUNCATE TABLE SOREPORT.INVENTORY"""
O_cur.execute(O_sql_truncate)
O_con.commit()

print("INSERTING {} INTO INVENTORY table on APEX".format(csv_tmp_file))
csv_data = csv.reader(open(csv_tmp_file)) 
O_sql_prep = """INSERT INTO SOREPORT.INVENTORY(FQDN,ICON,CHASSIS,MACHINE_ID,BOOT_ID,VIRTUALIZATION,OS,CPE,KERNEL,ARCH,ENVIRONMENT) VALUES(:1,:2,:3,:4,:5,:6,:7,:8,:9,:10,:11)"""
for C_row in csv_data:
    C_row=[None if cell == '' else cell for cell in C_row]
    O_cur.execute(O_sql_prep,C_row)


O_con.commit()
O_cur.close()
O_con.close()
print("INSERT/COMMIT INVENTORY TABLE COMPLETE ON DB APEX")

if os.path.exists(csv_tmp_file):
  print("REMOVING/CLEANING UP {}".format(csv_tmp_file))
  os.remove(csv_tmp_file)
print("UPLOAD/INSERT FROM {} INTO TABLE INVENTORY COMPLETE".format(csv_tmp_file))

