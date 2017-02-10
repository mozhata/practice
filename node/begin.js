function showTime(str){
				console.log("str: " + str);
				if (str.substring(0, 4) === "0000"){
					console.log("is 0000");
					return null
				}
				return str
			}
console.log(showTime("000000"))
console.log(showTime("lsdagalsg"))



/*
var user_infos = [
    ['小A', '女', 21, '大一'],
    ['小B', '男', 23, '大三'],
    ['小C', '男', 24, '大四'],
    ['小D', '女', 21, '大一'],
    ['小E', '女', 22, '大四'],
    ['小F', '男', 21, '大一'],
    ['小G', '女', 22, '大二'],
    ['小H', '女', 20, '大三'],
    ['小I', '女', 20, '大一'],
    ['小J', '男', 20, '大三'],
]
for (i = 0; i < user_infos.length; i++) {
    if (user_infos[i][3] == "大一" && user_infos[i][1] == "女") {
        console.log(user_infos[i]);
    }
}
*/

/*// 测试break跳出几层循环
console.log("test break");
for (i = 0; i < 10; i++) {
    console.log("first level");
    for (j = 0; j < 10; j++) {
        console.log("the secode level");
        if (i > 1) {
        	console.log("break..");
            break
        }
        console.log("i: " + i + "j: " + j);
    }
}
*/

/*var http=require('http');
http.createServer(function(req,res){
  res.writeHead(200,{'Content-Type':'text/plain'});
  res.end('Hello \n');
}).listen(1337,'127.0.0.1');
console.log('Server running at http://127.0.0.1:1337/');
*/
