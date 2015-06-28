<?php
/*
 Date: 2015-06-28
 Authore: Coding5
 Purpose: php 實作網頁程式 
 Websit: http://www.coding5.com/works?transno=1000000009
*/

require_once("DBConn.php"); 
$db = new DB(); 
$db->connect_db($_DB['host'], $_DB['username'], $_DB['password'], $_DB['dbname']);

$sinput1="{$_POST['input1']}";

$sql = "SELECT codeB FROM coding5 WHERE codeA = '".$sinput1."'"; 
$db->query($sql); 

$array=0;
while($result = $db->fetch_array()){
  $codeB$array]="{$result["codeB"]}";

  print $codeB[$array];

  $array++;
}

$count1=$array;

$db->closeDB(); 

?>
