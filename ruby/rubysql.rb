# Date: 2015-04-16
# Authore: Coding5
# Purpose: Ruby 物件化DB 連線client 端呼叫 
# Websit: http://www.coding5.com/works?transno=1000000012

require './dbconn.rb'

begin
  db = DBConn.new
    
  db.getConn

  sSql = "SELECT value FROM table"

  rs = db.runQuery(sSql) 
  rs.each_hash do |row|
    puts row['value']
  end

rescue Mysql::Error => e
  puts e.errno
  puts e.error

ensure
  db.dbClose
end
