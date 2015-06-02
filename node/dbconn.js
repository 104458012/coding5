// Date: 2015-05-24
// Authore: Coding5
// Purpose: node.js MySQL CRUD 操作 
// Websit: http://www.coding5.com/works?transno=1000000029

var mysql = require('mysql');

function getConn(){
  var db  = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: 'pass123',
    database: 'DBNAME',
    port: 3306 
  });

  db.connect();

  console.log("dbconn start.");

  return db;
}

function getQuery(sql) {
  var db = getConn();

  db.query(sql, function(err, results) {
    if (err) {
      console.log("mysql getQuery error");  
    }
  });

  db.end();
}

function getInsert(sql) {
  var db = getConn();

  db.query(sql, function(err) {
    if (err) {
      console.log("mysql getInsert error");  
    }
  });

  db.end();
}

function getUpdate(sql) {
  var db = getConn();

  db.query(sql, function(err) {
    if (err) {
      console.log("mysql getUpdate error");  
    }
  });

  db.end();
}

function getDelete(sql) {
  var db = getConn();

  db.query(sql, function(err) {
    if (err) {
      console.log("mysql getDelete error");  
    }
  });

  db.end();
}

exports.getQuery = getQuery;
exports.getInsert = getInsert;
exports.getUpdate = getUpdate;
exports.getDelete = getDelete;
