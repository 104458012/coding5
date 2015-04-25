# Date: 2015-04-25
# Authore: Coding5
# Purpose: python Client 端delete 呼叫 
# Websit: http://www.coding5.com/works?transno=1000000019

# coding=UTF-8
from dbconn import DBConn

try:
  dbuse = DBConn()

  dbuse.dbConnect()

  sql = "DELETE coding5 WHERE codeA = 'aaa'"
  dbuse.runDelete(sql)

  dbuse.dbClose()
except:
  print "MySQL DB Error"
