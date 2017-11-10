package zuiqiang

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/axgle/mahonia"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

var novelMap = map[string]string{
	"49981": "tulong",
}

const (
	host = "http://www.33zw.com"
	size = 200
)

const (
	novelNum = "49981" // tulong
	from     = 545
)

type chapterLink struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

func init() {
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("./logs")
	flag.Parse()
}

var decoder = mahonia.NewDecoder("gbk")

func Tulong() {
	catalogURL := fmt.Sprintf("%s/xiaoshuo/%s/", host, novelNum)
	links, err := getChaperList(from, size, catalogURL)
	if err != nil {
		glog.Fatalf("getChapterList failed: %v", err)
	}

	// download
	for i, l := range links {
		downloadFile(l.Href, l.Title+".html")
		glog.Infof("\n\nthe %dst chap %s downloaded", i, l.Title)
	}

	novelFile := fmt.Sprintf("%s_%d_%d.txt", novelMap[novelNum], from, size)
	// parse
	for i, l := range links {
		glog.Infof("parsing the %dst chapter: %s", i, l.Title)
		content, err := parseChapter(l.Title + ".html")
		if err != nil {
			glog.Errorf("ignore parseChapter err: %v, continue", err)
			continue
		}
		prefixTrims := []string{}
		suffixTrims := []string{
			"聽聽聽聽",
			"聽",
		}
		content = modifyContent(content, l.Title, prefixTrims, suffixTrims)
		err = writeContentToFile(content, novelFile)
		if err != nil {
			glog.Errorf("ignore writeContentToFile failed: %v", err)
			continue
		}
	}

}

func modifyContent(content, title string, prefixTrims, suffixTrims []string) string {
	// content = strings.TrimPrefixr(content, prefix)
	for _, pre := range prefixTrims {
		content = strings.Replace(content, pre, "", -1)
	}
	for _, suffix := range suffixTrims {
		content = strings.Replace(content, suffix, "", -1)
	}
	content = fmt.Sprintf("%s\n\n%s\n\n\n\n\n\n\n", title, content)
	return content
}

func writeContentToFile(content, filename string) error {
	dest, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(fmt.Sprintf("writeContentToFile failed: %v", err))
	}
	defer dest.Close()
	dest.WriteString(content)
	return nil
}

func parseChapter(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("parseChapter: open file failed: %v", err))
	}
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(fmt.Sprintf("NewDocumentFromReader failed: %v", err))
	}
	content := doc.Find("#content").Text()
	index := strings.LastIndex(content, "手机用户")
	if index > 0 {
		if len(content[index:]) < 500 {
			content = content[:index]
		}
	}
	return content, nil
}

func getChaperList(from, size int, catalogURL string) ([]chapterLink, error) {
	doc, err := goquery.NewDocument(catalogURL)
	if err != nil {
		panic(fmt.Sprintf("failed to catalogURL. err: %s", err))
	}

	result := make([]chapterLink, 0, size)
	chapList := doc.Find("#list > ul._chapter > li > a")
	for i := from; i < chapList.Length(); i++ {
		if i > from+size {
			break
		}
		title := chapList.Eq(i).Text()
		if decoder == nil {
			panic("ops..")
		}
		title = decoder.ConvertString(title)
		href, found := chapList.Eq(i).Attr("href")
		if !found {
			continue
		}
		result = append(result, chapterLink{href, title})

		glog.Infof("the %dst item: %q href: %s", i, title, href)
	}
	return result, nil
}

func downloadFile(url, filename string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	str, err := doc.Html()
	if err != nil {
		panic(fmt.Sprintf("downloadFile: %v", err))
	}
	str = decoder.ConvertString(str)
	savePage(str, filename)
	return nil
}

func savePage(str string, filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		panic(fmt.Sprintf("savePage:open failed: %v", err))
	}
	defer f.Close()
	f.WriteString(str)
	glog.Infoln("savePage: ok")
}
