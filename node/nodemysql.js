var mysql = require('mysql');

var db = mysql.createConnection({
  host:'localhost',
  user: 'root',
  password: 'pass123',
  database: 'DBNAME',
  port:3306
});

db.connect();

db.query('select codeA as codeA from coding5;',function(err, results){
  if(err){
    console.log("mysql error");
  }

  console.log(results[0].codeA);
});

db.end();
