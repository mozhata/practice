package jianzhuang

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

var novelMap = map[string]string{
	"79_79596": "qinli",
	"79_79877": "daxiaji",
}

const (
	host = "http://www.xxbiquge.com"
	size = 200
)

// const (
// 	novelNum = "79_79596" // qinli
// 	from     = 284
// )

const (
	novelNum = "79_79877" // daxiaji
	from     = 265
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

func Jianzhuang() {
	catalogURL := fmt.Sprintf("%s/%s/", host, novelNum)
	links, err := getChaperList(from, size, catalogURL)
	if err != nil {
		glog.Fatalf("getChapterList failed: %v", err)
	}

	// download
	for i, l := range links {
		if strings.HasPrefix(l.Href, "/"+novelNum+"/") {
			filename := strings.TrimPrefix(l.Href, "/"+novelNum+"/")
			if filename != "" {
				chapterURI := host + l.Href
				downloadFile(chapterURI, filename)
				glog.Infof("\n\nthe %dst chap %s downloaded", i, l.Title)
			}
		}
	}

	novelFile := fmt.Sprintf("%s_%d_%d.txt", novelMap[novelNum], from, size)
	// parse
	for i, l := range links {
		glog.Infof("parsing the %dst chapter: %s", i, l.Title)
		if strings.HasPrefix(l.Href, "/"+novelNum+"/") {
			filename := strings.TrimPrefix(l.Href, "/"+novelNum+"/")
			if filename != "" {
				content, err := parseChapter(filename)
				if err != nil {
					glog.Errorf("ignore parseChapter err: %v, continue", err)
					continue
				}
				prefixTrims := []string{
					"chaptererror();",
					"(未完待续。。)",
					"本站重要通知：你还在用网页版追小说吗？使用本站的免费小说APP，会员同步书架，文字大小调节、阅读亮度调整、更好的阅读体验，请关注微信公众号 jiakonglishi (按住三秒复制) 下载免费阅读器!!",
					"请关注微信公众号",
					"在线看:meinvmei222",
					"(长按三秒复制)",
					"公告：笔趣阁免费APP上线了，支持安卓，苹果。请关注微信公众号进入下载安装 wanbenheji (按住三秒复制)!!",
					"混血女主播直播后忘关摄像头私_生活视频遭曝光",
					"公告：免费小说app安卓，支持安卓，苹果，告别一切广告，请关注微信公众号进入下载安装 zuopingshuji 按住三秒复制!!",
					"请关注微信公众号进入下载安装",
					"免费小说app安卓，支持安卓，苹果，告别一切广告",
					"请关注微信公众号在线看:",
				}
				suffixTrims := []string{}
				content = modifyContent(content, l.Title, prefixTrims, suffixTrims)
				err = writeContentToFile(content, novelFile)
				if err != nil {
					glog.Errorf("ignore writeContentToFile failed: %v", err)
					continue
				}
			}
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
	return content, nil
}

func getChaperList(from, size int, catalogURL string) ([]chapterLink, error) {
	doc, err := goquery.NewDocument(catalogURL)
	if err != nil {
		panic(fmt.Sprintf("failed to catalogURL. err: %s", err))
	}

	result := make([]chapterLink, 0, size)
	chapList := doc.Find("#list >dl > dd > a")
	for i := from; i < chapList.Length(); i++ {
		if i > from+size {
			break
		}
		title := chapList.Eq(i).Text()
		href, found := chapList.Eq(i).Attr("href")
		if !found {
			continue
		}
		result = append(result, chapterLink{href, title})

		// glog.Infof("the %dst item: %q href: %s", i, title, href)
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
