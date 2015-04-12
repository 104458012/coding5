/*
* Date: 2015-04-12
* Authore: Coding5
* Purpose: Demo java DB module call 
* Websit: http://www.coding5.com
*/

public class MainCall{
  public static void main(String[] args){
    String sql;
    String sCodeA;

    DBConn db = new DBConn();

    try{
      sql = "SELECT codeA FROM coding5 ";
      
      ResultSet rs = db.runSql(sql); 

      while(rs.next()){
        sCodeA = rs.getString("codeA");
        System.out.println("["+sCodeA+"]");
      }
      
      db.closeDB();
    }catch(Exception e){
      System.out.println ("Exception: " + e);
    }
  }
}
