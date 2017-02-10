querystring = require("querystring")
uri = querystring.stringify({"name": "scott", "student": ["abc", "ds"]}, ";", ":")
console.log(uri);
uo = querystring.parse(uri, ";", ":")
console.log(uo);
console.log(querystring.escape("<哈>"))
console.log(querystring.unescape("%3C%E5%93%88%3E"));



// url = require("url")

// link = "https://kang:123@nodejs.org:8989/dist/latest-v6.x/docs/api/url.html?abc=123#url_host"

// uo = url.parse(link, true)



// console.log(uo)
// console.log(url.format(uo));
// console.log(url.resolve("http://abc.com:99/a/", "q/s"));