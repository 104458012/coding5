// Date: 2015-05-24
// Authore: Coding5
// Purpose: node.js SQL update statement 
// Websit: http://www.coding5.com/works?transno=1000000029

var mysql = require('./dbconn.js');

var sql = "UPDATE coding5 SET codeA = 'codeC' WHERE codeA = 'codeA'";

mysql.getUpdate(sql);
