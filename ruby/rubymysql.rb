# Date: 2015-04-14
# Authore: Coding5
# Purpose: Ruby 連接MySQL 資料庫 
# Websit: http://www.coding5.com/works?transno=1000000011

require 'rubygems'
require 'mysql'

begin
  con = Mysql.new('localhost', 'username', 'passwd', 'DBMANE')

  # 建立SQL Statement
  rs = con.query("SELECT codeA FROM coding5")
  
  # 撈取欄位資料
  rs.each_hash do |row|
    puts row['codeA']
  end
  
rescue Mysql::Error => e
  puts e.errno
  puts e.error

ensure
  con.close if con
end
