# Date: 2015-04-20
# Authore: Coding5
# Purpose: Ruby 建立DB delete 操作 
# Websit: http://www.coding5.com/works?transno=1000000013

require './dbconn.rb'

begin
  db = DBConn.new
    
  db.getConn

  sql = "DELETE coding5 WHERE codeA = 'bb'"
  db.runDelete(sql)

rescue Mysql::Error => e
  puts e.errno
  puts e.error

ensure
  db.dbClose

end
