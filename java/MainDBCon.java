/*
* Date: 2015-04-12
* Authore: Coding5
* Purpose: Demo java program use JDBC connect oracle DB 
* Websit: http://www.coding5.com
*/

import java.sql.*;

public class MainDBCon{
  public static void main(String[] args){
    ResultSet rs=null;
    String sql;
    String sValue;

    try{
      DriverManager.registerDriver (new oracle.jdbc.driver.OracleDriver());

      Connection conn = DriverManager.getConnection 
                        ("jdbc:oracle:thin:@127.0.0.1:1521:dbname","user","passwd");

      Statement stmt= conn.createStatement
                      (ResultSet.TYPE_SCROLL_SENSITIVE,ResultSet.CONCUR_UPDATABLE);

      sql = "SELECT VALUE FROM TESTDB WHERE ROWNUM < 3"; 

      rs=stmt.executeQuery(sql);

      while(rs.next()){
        sValue = rs.getString("VALUE");
        System.out.println("["+sValue+"]");
      }

      conn.close();
    }catch(Exception e){
      System.out.println ("Exception: " + e);
    }
  }
}
