import pyodbc
# Create connection
con = pyodbc.connect(driver="{SQL Server}",server=SERVERNAME,database=DATA_BASE_INFO,uid=username,pwd=password)
cur = con.cursor()
db_cmd = "SELECT * FROM table_name"
res = cur.execute(db_cmd)
# Do something with your result set, for example print out all the results:
for r in res:
    print r
