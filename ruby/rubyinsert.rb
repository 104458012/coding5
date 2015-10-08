# Date: 2015-04-20
# Authore: Coding5
# Purpose: Ruby 建立DB insert 操作 
# Websit: http://www.coding5.com/works?transno=1000000013

require './dbconncrud.rb'

begin
  db = DBConn.new
    
  db.getConn

  sql = "INSERT coding5 INTO (codeA) VALUES ('aa')"
  db.runInsert(sql)

rescue Mysql::Error => e
  puts e.errno
  puts e.error

ensure
  db.dbClose

end
