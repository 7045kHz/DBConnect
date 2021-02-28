import cx_Oracle

con = cx_Oracle.connect('user/passwd@fqdn.example.com:1521/sid')
ver = con.version.split(".")
cur = con.cursor()
cur.execute('select * from v_items@SOR')
res = cur.fetchall()
print("Looping through results")
for row in res:
    #l = row[2]
    r = list()
    for c in row:
        if type(c) is cx_Oracle.LOB or cx_Oracle.CLOB:
         r.append(str(c))
        else:
            r.append(c)

    print(r)
con.close()
