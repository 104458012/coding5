// Date: 2015-05-24
// Authore: Coding5
// Purpose: node.js SQL insert statement 
// Websit: http://www.coding5.com/works?transno=1000000029

var mysql = require('./dbconn.js');

var sql = "INSERT INTO coding5 (codeA) VALUES ('codeA')";

mysql.getInsert(sql);
