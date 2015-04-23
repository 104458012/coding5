# Date: 2015-04-23
# Authore: Coding5
# Purpose: python 物件化DB 連線client 端呼叫 
# Websit: http://www.coding5.com/works?transno=1000000018

# coding=UTF-8
from dbconn import DBConn

try:
  dbuse = DBConn()

  dbuse.dbConnect()

  sql = "SELECT codeA FROM coding5"
  dbuse.runQuery(sql)

  for record in dbuse.results:
    codeA = record[0]

    print "%s" % (codeA)

  dbuse.dbClose()
except:
  print "MySQL DB Error"
