<?php
/*
 Date: 2015-04-29
 Authore: Coding5
 Purpose: php 建立資料庫物件導向化 
 Websit: http://www.coding5.com/works?transno=1000000007
*/

global $_DB;

$_DB['host'] = "localhost:3306";
$_DB['username'] = "user";
$_DB['password'] = "p123";
$_DB['dbname'] = "DBNAME";

class DB
{
  var $_dbConn = 0;
  var $_queryResource = 0;

  function connect_db($host, $user, $pwd, $dbname){
    $dbConn = mysql_connect($host, $user, $pwd);
    if (! $dbConn){
      echo "db connect error !!";
    }

    mysql_query("SET NAMES 'UTF8'"); 

    if (! mysql_select_db($dbname, $dbConn)){
       echo "mysql_select_db error !!";
    }

    $this->_dbConn = $dbConn;
    return true;
  }

  function runQuery($sql){
   if (! $queryResource = mysql_query($sql, $this->_dbConn)){
     echo "mysql_query error !!";
   }
   $this->_queryResource = $queryResource;
   return $queryResource;
  }

  function fetch_array(){
    return mysql_fetch_array($this->_queryResource, MYSQL_ASSOC);
  }

  function closeDB(){
    mysql_close($this->_dbConn);
  }
}

?>
