/*
* Date: 2015-04-13
* Authore: Coding5
* Purpose: 建立資料庫新增/修改/刪除物件導向的操作 
* Websit: http://www.coding5.com/works?transno=1000000004
*/

import java.sql.*;

public class DBConn{
  private Connection conn = null;

  public DBConn(){
    try{
      DriverManager.registerDriver
             (new oracle.jdbc.driver.OracleDriver());

      conn = DriverManager.getConnection 
             ("jdbc:oracle:thin:@127.0.0.1:1521:dbname",
              "username", "passwd");

    }catch(Exception e){
       e.printStackTrace();
    }
  }

  public ResultSet runQuery(String sql) throws SQLException {
    Statement sta = conn.createStatement();
    return sta.executeQuery(sql);
  }

  public int runUpdate(String sql) throws SQLException {
    Statement sta = conn.createStatement();
    return sta.executeUpdate(sql);
  }

  public int runInsert(String sql) throws SQLException {
    Statement sta = conn.createStatement();
    return sta.executeUpdate(sql);
  }

  public int runDelete(String sql) throws SQLException {
    Statement sta = conn.createStatement();
    return sta.executeUpdate(sql);
  }

  public void closeDB() throws SQLException {
    conn.close();
  }
}
