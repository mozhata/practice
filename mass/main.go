package main

import (
	// . "bitbucket.org/applysquare/applysquare-go/pkg/discussion"

	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

var P func(...interface{}) (int, error) = fmt.Println
var parse func(string, int, int) (int64, error) = strconv.ParseInt
var atoi func(string) (int, error) = strconv.Atoi
var EmailPattern = `^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})$`

type JSONTime time.Time

const (
	JSONTimeFormat = "2006-01-02T15:04:05Z"
)

var formated string = "2015-08-05T04:09:05Z"
var toConverted string = "2015-08-26T05:39:10Z"

type UMPlatform int

const (
	PlantformIos UMPlatform = iota
	PlantformAndriod
)

func main() {
	flag.Parse()
	// s1 := "abcdditui"
	// // s2 := "1234"
	// ss := md5.Sum([]byte(s1))
	// P("md5: ", ss)
	// fmt.Printf("%x", ss)
	// boo, _ := regexp.MatchString(EmailPattern, "1.sg@q.com")
	// P("boo: ", boo)
	// P("format: ", time.Now().Format(JSONTimeFormat))
	// t, err := time.Parse(JSONTimeFormat, toConverted)
	// P("t: ", t, "\nerr : ", err)
	// t, err = time.Parse(JSONTimeFormat, formated)
	// P("t: ", t, "\nerr : ", err)
	// var s string
	// P(s == "")
	// P(formated[:len(formated)-1])
	// P(time.Now().Format("2006-01-02"))
	// P(rand.Int())
	// P(rand.Int())
	// P(PlantformIos)
	// P(PlantformAndriod)
	// source := rand.NewSource(8)
	// r := rand.New(source)
	// P(r.Int())
	// P(r.Int())
	// P(rand.Intn(5))
	// P(rand.Intn(5))
	// P(rand.Intn(5))
	// P(rand.Intn(5))
	// flo := float32(1.0 / 3)
	// P(flo)
	// P(time.Now())
	// P(time.Now().Local())
	// locationC, errC := time.LoadLocation("Asia/Shanghai")
	// locationU, errU := time.LoadLocation("UTC")
	// P(locationC, errC)
	// P(locationU, errU)
	// P(time.Now().In(locationC))
	// P(time.Now().In(locationU))
	// timeStamp := strconv.Itoa(int(time.Now().In(locationC).Unix()))
	// P(timeStamp)
	// v, e := atoi("a")
	// P(v, e)
	// v, e = atoi("")
	// P(v, e)
	// c := cron.New()
	// c.AddFunc("2 0 0 * * *", func() { fmt.Println("...") })
	// c.Start()
	// code := "abcd12"
	// P(code[:4])
	// password := fmt.Sprintf("%x", md5.Sum([]byte("cn.sduhaliluya")))[:4]
	// P(password)
	// var floaNan float64
	// var floatSmall float64 = 11.01
	// P(floaNan)
	// P(floatSmall)
	// P(math.IsNaN(floaNan), math.IsInf(floaNan, 0), math.IsInf(floaNan, 1), math.IsInf(floaNan, -1))
	// P(math.IsNaN(floatSmall), math.IsInf(floatSmall, 0), math.IsInf(floatSmall, 1), math.IsInf(floatSmall, -1))
	// P(&floaNan, &floatSmall)
	// var s map[string]interface{}
	// P(s)
	// P(len(s))
	// P(s == nil)
	// date := time.Now().Add(time.Hour * -24 * 7).Format("2006-01-02")
	// P(date)
	// var f float64
	// var ff float64 = 0.00001
	// P(f, ff)
	// P(ff > f)
	// P(ff > 0)
	// P(f > 0)
	// P(f == 0.0000000000000000000000)
	// var ss = struct {
	// 	a int
	// 	b string
	// }{}
	// P(ss)
	// P(&ss == nil)
	// l := field_of_study.KeyList
	// P(l)
	// P(len(l))

	// dic1 := map[string]string{"a": "aa", "as": "asas", "dd": "dddd"}
	// P(dic1, &dic1)
	// dic2 := dic1
	// dic2["a"] = "cc"
	// P(dic1, dic2)
	// P(&dic2)

	// var dic map[string]string
	// P(len(dic))
	// var s string
	// ll := strings.Split(s, ",")
	// P(ll)
	// P(len(ll))

	// reg := func() *regexp.Regexp {
	// 	wordList := []string{}
	// 	for i := range wordList {
	// 		wordList[i] = regexp.QuoteMeta(wordList[i])
	// 	}
	// 	return regexp.MustCompile(strings.Join(wordList, "|"))
	// }()
	// P(reg.MatchString("abc"))
	// P("reg is nil", reg == nil)
	// P(reg)
	// P(*reg)
	// P("string: ", reg.String(), reg.String() == "")
	// P(MarshalJSONOrDie(reg))
	// dest := []string{"abc", "dsa"}
	// fmt.Println(dest)
	// saveToSlice("source", dest)
	// fmt.Println(dest)
	// P(buldCountySlug(""))
	// P((1 == 2) ^ (3 == 4))
	// sl := make([]string, 0)
	// P(len(sl))
	// sle := make([]string, 7)
	// P(len(sle))
	// hello()

	// osPath()
	// urlParse()
	// loopMap()
	// sorttt()
	// crawl()
	// convertInterface()

	// runes()
	// logFlush()
	// build()
	// intLoop()
	// syncRuntime()
	// closure()
	// timeAdd()
	// tryCall()
	// getEnv()
	// tesMap()
	// dup1()
	// charCount()
	tesFor()
}

// never found a better parterner for a long time, maybe met but I'am not brelliant to deserve it at that time. I need to focus and try to make me better and find a parterner to encourage and take care of each other

// review this
func charCount() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			glog.Infoln("is io.EOF")
			break
		}
		if line, prefix, _ := in.ReadLine(); string(line) == "" || string(line) == "\\n" {
			glog.Infof("line is %q, input should over, prefix: %t", string(line), prefix)
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func tesFor() {
	slice := []string{"abd", "dsad", "abc", "ddd"}
	for v := range slice {
		glog.Infoln(v)
	}
	for i, v := range slice {
		glog.Infoln(i, "\t", v)
	}
}

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		glog.Infoln("one scan, and the input is : ", input.Text())
		if input.Text() == "\n" {
			glog.Infoln("input is `\\n` over")
			break
		}
		if input.Text() == "" {
			glog.Infoln("the input is \"\", over")
			break
		}
		counts[input.Text()] += 1
	}

	for key, v := range counts {
		if v > 1 {
			fmt.Printf("%s:\t%d\n", key, v)
		}
	}
}

func tesMap() {
	var dic = map[string]int{}
	dic["a"] = 9
	delete(dic, "b")
	dic["c"] += 2
	fmt.Println(dic)

	// for k := range dic {
	// 	fmt.Println("K: ", k)
	// }
	var dic2 map[string]int
	var dic3 = map[string]int{}
	fmt.Println(len(dic2), len(dic3))
	fmt.Println("lookup: ", dic2["bb"], "-", dic3["bb"])
	delete(dic2, "kk")
	delete(dic3, "bb")
	fmt.Println("loop dic2: ")
	for k := range dic2 {
		fmt.Println("key in dic2: ", k)
	}
	for k := range dic3 {
		fmt.Println("key in dic3: ", k)
	}
	fmt.Println("==nil: ", dic2 == nil, dic3 == nil)
	// dic2["b"] = 8
	// fmt.Println(dic2)
}

func getEnv() {
	gopath := os.Getenv("ABC")
	fmt.Println("ABC: ", gopath)
}

func tryCall() {
	glog.Infoln("tryCall...")
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		glog.Infof("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
	}
}

func timeAdd() {
	one := time.Now()
	glog.Infoln("one: ", one)
	two := one.Add(time.Minute * 10)
	glog.Infoln("two: ", two)
	glog.Infoln("one: ", one)

}

func closure() {
	closure1()
	closure2()
}

func closure1() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i) // 555555
		}()
	}
	wg.Wait()
}

func closure2() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			fmt.Println(j) // 01234(顺序可能会变)
		}(i)
	}
	wg.Wait()
}

// go run main.go -logtostderr
func syncRuntime() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.baidu.com/",
		"http://dict.youdao.com/w/currency/#keyfrom=dict2.top",
		"https://docs.mongodb.com/manual/mongo/",
		"http://www.runoob.com/mongodb/mongodb-q,uery.html",
		"http://studygolang.com/articles/2059",
	}
	glog.Infoln("fetching url..")
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			glog.Infoln("fetch url: ", url)
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			r, _ := http.Get(url)
			glog.Infoln("status: %s, code: %d, url is %s", r.Status, r.StatusCode, url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

// test loop
type MyInt int

func (m MyInt) String() string {
	return fmt.Sprint(int(m))
}

// panic
func intLoop() {
	var m MyInt = 888
	s := m.String()
	fmt.Println(s)
}

func build() {
	uids := []int{1256, 1457, 2799, 9448, 27078, 27090, 27095, 27104, 27148, 27160, 27988}
	// AND creator=%d;
	sql := "SELECT id FROM eventtopic WHERE `delete`=0"
	for _, uid := range uids {
		sql = fmt.Sprintf("%s%s", sql, fmt.Sprintf(" OR creator=%d", uid))
	}
	fmt.Println("sql: \n", sql+";")

}

func logFlush() {
	// go run main.go -alsologtostderr -log_dir="./"
	flag.Parse()
	glog.Infoln("abc..")
	glog.Infof("abc..%d", 123)
	glog.Flush()
}

func runes() {
	str := "hello中国"
	conv := []rune(str)
	P(str)
	for _, v := range conv {
		P(fmt.Sprintf("%d", v))
		P(string(v))
	}
	ss := []rune(str)[0]
	P(string(ss))
}

func convertInterface() {
	foo := Interface()
	tryOne, ok := foo.(map[string]string)
	fmt.Println(foo, tryOne, ok)
	tryTwo, ok := foo.(map[string]interface{})
	fmt.Println(tryTwo, ok)
}

func Interface() interface{} {
	foo := make(map[string]interface{})
	foo["key"] = "value"
	return foo
}

func makeNew() {

}

func crawl() {
	P("begin...")
	url := "http://www.applysquare.com/cn/"
	doc, err := goquery.NewDocument(url)
	P(doc, err)
}

func sorttt() {
	slice := []string{"a", "c", "b"}
	slice2 := []string{"a", "c", "b", "d"}

	sort.Strings(slice)
	sort.Strings(slice2)
	P(slice)
	P(slice2)
}

func loopMap() {
	dict := map[string]string{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	}
	for key, val := range dict {
		P(key, val)
	}
}

func urlParse() {
	backend, err := url.Parse("http://localhost:8888/dir1/dir2/tail?q=123&q2=abc#alt")
	P(backend, err)
	P(MarshalJSONOrDie(backend))
}

func osPath() {
	d, err := os.Getwd()
	// 得到运行改命令时所在的目录, 比如在上一级目录运行该命令, 得到的目录就是上一级的
	P(d, err)
	dir1 := "/Users/mozhata/work/src/practice"
	dir2 := "/Users/mozhata/work/src/practice/"
	// /Users/mozhata/work/src
	P(path.Dir(dir1))
	// /Users/mozhata/work/src/practice
	P(path.Dir(dir2))
}

func buldCountySlug(slug string) string {
	countryKey := strings.Split(slug, ".")[0]
	if countryKey == "" {
		countryKey = "us"
	}
	return fmt.Sprintf("country_%s", countryKey)
}
func saveToSlice(source string, dest []string) {
	dest = append(dest, source)
}
func MarshalJSONOrDie(v interface{}) string {
	b, err := json.Marshal(v)
	Check(err)
	return string(b)
}
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

/*
map-reduce
*/
/*
(doc) ->
	 return if doc.doc_type isnt 'Profile'
	 emit [doc.app_info.device_identifier.substr(0,2),doc.app_info.device_identifier.substr(0,4) ],null if doc.app_info.device_identifier
*/
// func Tt(s1,s2 string){
// 	print s1,s2
// }
func hello() {
	fmt.Println("hello world")
}
