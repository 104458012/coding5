// Date: 2015-05-18
// Authore: Coding5
// Purpose: 檢查node.js server 執行成功 
// Websit: http://www.coding5.com/works?transno=1000000024

http.createServer(function(request, response) {
  response.writeHead(200, {"Content-Type": "text/plain"});
  response.write("Hello World");
  response.end();
}).listen(8888);
