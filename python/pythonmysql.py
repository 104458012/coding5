# Date: 2015-05-04
# Authore: Coding5
# Purpose: python 建立與MySQL 資料庫連線 
# Websit: http://www.coding5.com/works?transno=1000000017

# coding=UTF-8

import MySQLdb

try:
  db = MySQLdb.connect
  ("localhost","username","passwd","DBNAME",charset='utf8')

  sql = "SELECT codeA, codeB FROM coding5"

  cursor = db.cursor()
  cursor.execute(sql)

  results = cursor.fetchall()

  for record in results: 
      col1 = record[0]
      col2 = record[1]
    
      print "%s, %s" % (col1, col2)

  db.close()

except MySQLdb.Error as e:
  print "Error %d: %s" % (e.args[0], e.args[1])

