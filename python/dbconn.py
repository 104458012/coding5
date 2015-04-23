# Date: 2015-04-16
# Authore: Coding5
# Purpose: python 物件導向化DB 建立資料庫連線 
# Websit: http://www.coding5.com/works?transno=1000000018

import MySQLdb

class DBConn:
  def __init__(self):
    self.user = 'username'
    self.host = 'localhost'
    self.passwd = 'passwd'
    self.dbname = 'DBNAME'

  def dbConn(self):
    self.db = MySQLdb.connect(
        self.host,self.user,self.passwd,self.dbname,charset='utf8')
    self.cursor = self.db.cursor()

  def runQuery(self, sql):
    self.cursor.execute(sql)

    self.results = self.cursor.fetchall()

  def dbClose(self):
    self.db.close()
