// Date: 2015-05-24
// Authore: Coding5
// Purpose: node.js SQL delete statement
// Websit: http://www.coding5.com/works?transno=1000000029

var mysql = require('./dbconn.js');

var sql = "DELETE coding5 WHERE codeA = 'codeC'";

mysql.getDeletet(sql);
