# Date: 2015-04-16
# Authore: Coding5
# Purpose: Ruby 物件化DB 建立資料庫連線 
# Websit: http://www.coding5.com/works?transno=1000000012

require 'rubygems'
require 'mysql'

class DBConn
  def getConn
    @con = Mysql.new('localhost', 'username', 'passwd', 'DBNAME')
  end
    
  def runQuery(sql)
    @con.query(sql)
  end

  def dbClose
    @con.close
  end
end
