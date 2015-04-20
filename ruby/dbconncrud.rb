# Date: 2015-04-20
# Authore: Coding5
# Purpose: Ruby 建立DB CRUD 資料庫操作 
# Websit: http://www.coding5.com/works?transno=1000000013

require 'rubygems'
require 'mysql'

class DBConn
  def getConn
    @con = Mysql.new('localhost', 'username', 'passwd', 'DBNAME')
  end

  def runQuery(sql)
    @con.query(sql)
  end

  def runUpdate(sql)
    @con.query(sql)
  end

  def runInsert(sql)
    @con.query(sql)
  end

  def runInsert(sql)
    @con.query(sql)
  end

  def dbClose 
    @con.close
  end
end
