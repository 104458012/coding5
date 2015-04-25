# Date: 2015-04-25
# Authore: Coding5
# Purpose: python 建立CRUD 操作function 
# Websit: http://www.coding5.com/works?transno=1000000019

import MySQLdb

class DBConn:
  def __init__(self):
    self.user = 'username'
    self.host = 'localhost'
    self.passwd = 'passwd'
    self.dbname = 'DBNAME'

  def dbConnect(self):
    self.db = MySQLdb.connect(
      self.host,self.user,self.passwd,self.dbname,charset='utf8')
    self.cursor = self.db.cursor()

  def runQuery(self, sql):
    self.cursor.execute(sql)
    self.results = self.cursor.fetchall()

  def runInsert(self, sql):
    self.cursor.execute(sql)
    self.db.commit()

  def runUpdate(self, sql):
    self.cursor.execute(sql)
    self.db.commit()

  def runDelete(self, sql):
    self.cursor.execute(sql)
    self.db.commit()

  def dbClose(self):
    self.db.close()
