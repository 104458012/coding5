/*
* Date: 2015-04-12
* Authore: Coding5
* Purpose: Demo java JDBC connect module 
* Websit: http://www.coding5.com
*/

import java.sql.*;

public class DBConn{
  private Connection conn = null;

  public DBConn(){
    try{
      DriverManager.registerDriver
      (new oracle.jdbc.driver.OracleDriver());
      conn = DriverManager.getConnection 
      ("jdbc:oracle:thin:@127.0.0.1:1521:dbname","username", "passwd");
    }catch(Exception e){
      e.printStackTrace();
    }
  }

  public ResultSet runSql(String sql) throws SQLException {
    Statement sta = conn.createStatement();
    return sta.executeQuery(sql);
  }

  public void closeDB() throws SQLException {
    conn.close();
  }
}
