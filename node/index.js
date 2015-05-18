// Date: 2015-05-18
// Authore: Coding5
// Purpose: 設定預設路徑 
// Websit: http://www.coding5.com/works?transno=1000000024

var server = require("./server");
var router = require("./router");
var requestHandlers = require("./requestHandlers");

var handle = {}
handle["/"] = requestHandlers.start;
handle["/start"] = requestHandlers.start;
handle["/upload"] = requestHandlers.upload;

server.start(router.route, handle);
