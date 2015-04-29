<?php
/*
 Date: 2015-04-29
 Authore: Coding5
 Purpose: php 建立Client 端呼叫 
 Websit: http://www.coding5.com/works?transno=1000000007
*/

require_once("DBConn.php"); 

$db = new DB(); 

$db->connect_db(
  $_DB['host'], $_DB['username'], $_DB['password'], $_DB['dbname']);

$sql = "SELECT codeA FROM coding5 ";

$db->runQuery($sql); 

$array=0;
while($result = $db->fetch_array()){
  $_aCodeA[$array]="{$result["codeA"]}";

  print $_aCodeA[$array];

  $array++;
}

$count1=$array;

$db->closeDB(); 

?>
