package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"net/http"
	"os"
	"practice/go/githublib"
	"runtime"
	"strings"
	"sync"
	"time"

	// "go.net/html"
	"golang.org/x/net/html"

	"github.com/golang/glog"
	"github.com/zykzhang/handy"
)

var P func(...interface{}) (int, error) = fmt.Println

const StudentNum int = 30

type HomeWork struct {
}

func main() {
	// shareVariable()
	// useChannel()
	// timeOut()
	// blockWithout()
	// notBlockWithInit()
	// bufferChannel()
	// HomeWorkThings()
	// chapOne()
	// chapTwo()
	// chapThree()
	// chapfour()
	chapFive()
}

func chapFive() {
	// findLinks()
	// switchCase()
}

func switchCase() {
	var arg = flag.String("foo", "", "...")
	var bar int
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()
	switch *arg {
	case "":

	case "+1":
		bar = bar + 1
	case "+2":
		bar = bar + 2
	default:
		bar = bar + 3
	}
	switch bar {
	case 0:
		glog.Infoln("foo is nil, do nothing")
	case 1:
		glog.Infoln("foo is +1")
	case 2:
		glog.Infoln("foo is +2")
	case 3:
		glog.Infoln("foo is other case")
	}
}

func findLinks() {
	doc, err := html.Parse(os.Stdin)
	glog.Infoln("begin.., doc: ", doc.Namespace, doc.Data)
	if err != nil {
		glog.Fatal(err)
	}
	links := make([]string, 0)
	links = visit(links, doc)
	glog.Infoln("loop.., links: ", links)
	for _, link := range links {
		glog.Infoln(link)
	}
}

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, atrr := range node.Attr {
			if atrr.Key == "href" {
				links = append(links, atrr.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func chapfour() {
	// issues()
	// treeSort()
	// issuesReport()
	// issuesReportHTML()
	// autoEscape()
}

func autoEscape() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

func issuesReportHTML() {
	// go run main.go repo:golang/go commenter:gopherbot json encoder >issues.html
	var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))
	result, err := githublib.SearchIssues(os.Args[1:])
	glog.Infoln(handy.MarshalJSONOrDie(result))
	if err != nil {
		glog.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		glog.Fatal(err)
	}
}

func issuesReport() {
	const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreateAt | daysAgo}} days
{{end}}`
	daysAgo := func(t time.Time) int {
		return int(time.Since(t).Hours() / 24)
	}
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		glog.Fatal(err)
	}
	result, err := githublib.SearchIssues(os.Args[1:])
	glog.Infoln(handy.MarshalJSONOrDie(result))
	if err != nil {
		glog.Fatal(err)
	}
	if err := report.Execute(os.Stderr, result); err != nil {
		glog.Fatal(err)
	}
}

func treeSort() {
	flag.Parse()
	sort := func(values []int) {
		var root *tree
		for _, v := range values {
			root = add(root, v)
		}
		appendValue(values[:0], root)
	}
	numbers := []int{2, 3, 5, 1, 7, 8, 33, 11, 44, 112}
	glog.Infoln("numbers before sort: ", numbers)
	sort(numbers)
	glog.Infoln("numbers after sort: ", numbers)
	glog.Infoln(numbers[:0])
}

type tree struct {
	value int
	left  *tree
	right *tree
}

func add(root *tree, v int) *tree {
	if root == nil {
		root = new(tree)
		root.value = v
		return root
	}
	if v < root.value {
		root.left = add(root.left, v)
	}
	if v > root.value {
		root.right = add(root.right, v)
	}
	return root
}

func appendValue(values []int, root *tree) []int {
	if root != nil {
		values = appendValue(values, root.left)
		values = append(values, root.value)
		values = appendValue(values, root.right)
	}
	return values
}

func issues() {
	result, err := githublib.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func chapThree() {
	// Operation()
	// BitOperation()
	// format()
	// float()
}

func float() {
	var z float64
	P(z, -z, 1/z, -1/z, z/z)          // 0 -0 +Inf -Inf NaN
	P(math.IsNaN(z), math.IsNaN(z/z)) // false true
	nan := math.NaN()
	P(nan == nan, nan < nan, nan > nan) // false false false
}

func format() {
	o := 0666 // 八进制
	x := int(0xdeadbeef)
	fmt.Printf("%d %[1]o %#[1]o %#[1]x\n", o) // 438 666 0666 0x1b6
	fmt.Printf("%d %[1]x %#[1]x\n", x)        // 3735928559 deadbeef 0xdeadbeef
}

func BitOperation() {
	var foo int8 = 6
	var bar int8 = 5
	fmt.Printf("%08b\t%d\n", ^bar, ^bar) //-0000110 -6
	fmt.Printf("%08b\n", foo^bar)        //00000011
	fmt.Printf("%08b\n", foo)            //00000110
	fmt.Printf("%08b\t%d\n", bar, bar)   //00000101 5
	fmt.Printf("%08b\n", foo&^bar)       //00000010
}

func Operation() {
	P(-5 % 2)
	P(-5 % -2)
	P(-5 / 2)
	P(-5 / 2.0)
}

func chapTwo() {
	// tesNil()
	// echo()
	// diffTypeOperation()
}

func diffTypeOperation() {
	type One int
	var one One = 11
	P(one - 1)

	var three One = 3
	P(one - three)

	// type Two int
	// var two = 2
	// P(one > two) // 失败
	type ComOne struct {
		One int
		Two int
	}
	type ComTwo struct {
		One int
		Two int
	}
	comTwo := ComTwo{1, 1}
	P(comTwo)
	val := ComOne(comTwo)
	P(val)
	// val2 := ComOne(one)	// failed
}

func echo() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")

	flag.Parse()
	fmt.Println(*n, *sep)
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func tesNil() {
	var x, y *int
	fmt.Println(x == y)
	d := 2
	x = &d
	y = &d
	fmt.Println(x == y)
}

func chapOne() {
	// hello()
	// echo1()
	// echo2()
	// dup1()
	// dup2()
	// dup3()
	// lissajous(os.Stdout)
	// fetch()
	// fetch2()
	// fetchAll()
	// tinyServer()
}

func tinyServer() {
	var mu sync.Mutex
	var count int

	handler := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
		fmt.Fprintf(w, "%s\t%s\t%s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "From[%q] = %q\n", k, v)
		}
	}

	image := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}

	counter := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Count: %d\n", count)
		mu.Unlock()
	}

	svg := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		svgSurface(w)
	}

	mandelbrot := func(w http.ResponseWriter, r *http.Request) {
		mandelbrot(w)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/img", image)
	http.HandleFunc("/svg", svg)
	http.HandleFunc("/mandelbrot", mandelbrot)
	log.Fatal(http.ListenAndServe(":7777", nil))
	// log.Fatal(http.ListenAndServe("localhost:7777", nil))

}

func svgSurface(out io.Writer) {
	const (
		width, height = 600, 320            // canvas size in pixels
		cells         = 100                 // number of grid cells
		xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
		xyscale       = width / 2 / xyrange // pixels per x or y unit
		zscale        = height * 0.4        // pixels per z unit
		angle         = math.Pi / 6         // angle of x, y axes (=30°)
	)
	var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

	f := func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return math.Sin(r) / r
	}

	corner := func(i, j int) (float64, float64) {
		// Find point (x,y) at corner of cell (i,j).
		x := xyrange * (float64(i)/cells - 0.5)
		y := xyrange * (float64(j)/cells - 0.5)

		// Compute surface height z.
		z := f(x, y)

		// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
		sx := width/2 + (x-y)*cos30*xyscale
		sy := height/2 + (x+y)*sin30*xyscale - z*zscale
		return sx, sy
	}

	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	content := ""
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			content = fmt.Sprintf("%s<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", content, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	svg = fmt.Sprintf("%s%s%s", svg, content, "</svg>")
	fmt.Fprintf(out, "%s", svg)
}

func mandelbrot(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	mandelbrot := func(z complex128) color.Color {
		const iterations = 200
		const contrast = 15

		var v complex128
		for n := uint8(0); n < iterations; n++ {
			v = v*v + z
			if cmplx.Abs(v) > 2 {
				return color.Gray{255 - contrast*n}
			}
		}
		return color.Black
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func fetchAll() {
	start := time.Now()
	fmt.Println("begin..")
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go subFetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func subFetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func fetch2() {
	fmt.Println("begin..")
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
	}
}
func fetch() {
	fmt.Println("begin..")
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		// fmt.Println(b)
	}
}

func lissajous(out io.Writer) {
	var palette = []color.Color{color.White, color.Black}
	const (
		whiteIndex = 0
		blackIndex = 1
	)

	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func dup3() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3 : %v", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func echo2() {
	// s, sep := "", ""
	var s = "" // 很少这样用
	var sep = ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func HomeWorkThings() {
	hwChan := make(chan HomeWork, 30)
	// hwChan := make(chan HomeWork, 20)  缓冲区可以改小
	exitChan := make(chan struct{})
	for i := 0; i < StudentNum; i++ {
		go student(hwChan)
	}
	go teacher(hwChan, exitChan)
	<-exitChan
	P("all things done")
}
func teacher(hwChan chan HomeWork, exitChan chan struct{}) {
	for i := 0; i < StudentNum; i++ {
		<-hwChan
		P("dealing homework..")
	}
	P("home work finished ")
	close(exitChan)
}
func student(hwChan chan HomeWork) {
	//学生提交作业
	P("sending home work..")
	hwChan <- HomeWork{}
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func hello() {
	fmt.Println("hello, $, 中文")
}

func blockWithout() {
	var ch chan int
	P("step one")
	go func() {
		ch <- 1
	}()
	P("step two")
	// 到此阻塞
	signal := <-ch
	P("step three", signal)

}
func notBlockWithInit() {
	ch := make(chan int)
	P("step one")
	go func() {
		ch <- 1
	}()
	P("step two")
	// 成功执行
	signal := <-ch
	P("step three", signal)

}

// timeout
func timeOut() {
	ch := make(chan int)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 1)
		timeout <- true
	}()
	select {
	case <-ch:
		fmt.Println("blabla..")
	case <-timeout:
		fmt.Println("timeout..")
	}
}

// 共享变量（内存）
func shareVariable() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()

		c := counter

		lock.Unlock()

		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}

// channel
func useChannel() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		// watch this
		chs[i] = make(chan int)
		go Count2(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func Count2(ch chan int) {
	fmt.Println("Count2ing")
	ch <- 1
}
