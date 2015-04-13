/*
* Date: 2015-04-13
* Authore: Coding5
* Purpose: client 端呼叫DB delete function 操作 
* Websit: http://www.coding5.com/works?transno=1000000004
*/

import java.sql.*;

public class MainDelete{
  public static void main(String[] args){
    String sql="";
    int ret = 0;

    DBConn db = new DBConn();

    try{
      sql = "DELETE coding5 WHERE codeA = 'a' ";
      
      ret = db.runDelete(sql);

      if(ret != 0){
         System.out.println("OK ret = "+ret);
      }else{
         System.out.println("sql statement error !!");
      }

      db.colseDB();
    }catch(Exception e){
      System.out.println ("Exception: " + e);
    }
  }
}
