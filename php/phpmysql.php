<?php
/*
 Date: 2015-05-06
 Authore: Coding5
 Purpose: php 建立MySQL 資料庫連線 
 Websit: http://www.coding5.com/works?transno=1000000006
*/

$dbhost = 'localhost:3306';
$dbuser = 'root';
$dbpass = 'p123';
$dbname = 'DBNAME';

$conn = mysql_connect($dbhost, $dbuser, $dbpass);

mysql_query("SET NAMES 'UTF8'");
mysql_select_db($dbname);

$sql = "select codeA from coding5 "; 
$result = mysql_query($sql); 

while($row = mysql_fetch_array($result)){
  echo $row['codeA']."<br>";
}

mysql_close($conn);

?>
