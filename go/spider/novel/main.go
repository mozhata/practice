package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

var novelMap = map[string]string{
	"86_86745": "shengxu",
	"1_1280":   "zaohuazhiwang",
	"0_444":    "xiuxiankuangtu",
	"3_3271":   "haoren",
	"78_78031": "yinianyongheng",
}

const (
	host = "http://m.biquge.tw"
	size = 200
)

// const (
// 	novelNum = "3_3271" // 郝仁
// 	from     = 1741
// )

// const (
// 	novelNum = "86_86745" // 圣墟
// 	from     = 594
// )

// const (
// 	novelNum = "0_444" // 修仙狂徒
// 	from     = 1200
// )

const (
	novelNum = "1_1280" // 造化之王
	from     = 2121
)

// const (
// 	novelNum = "78_78031" // yinianyongheng
// 	from     = 1000
// )

type chapterLink struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

func init() {
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("./logs")
	flag.Parse()
}

func main() {
	defer glog.Flush()
	common()
	// chaoshen.Chaoshen()
	// chaoshen.Gudaochangsheng()
	// zuiqiang.Tulong()
	// jianzhuang.Jianzhuang()
}

func common() {
	catalogURL := fmt.Sprintf("%s/%s/all.html", host, novelNum)
	links, err := getChaperList(from, size, catalogURL)
	if err != nil {
		glog.Fatalf("getChapterList failed: %v", err)
	}

	// download
	for i, l := range links {
		if strings.HasPrefix(l.Href, "/"+novelNum) {
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
		if strings.HasPrefix(l.Href, "/"+novelNum) {
			filename := strings.TrimPrefix(l.Href, "/"+novelNum+"/")
			if filename != "" {
				content, err := parseChapter(filename)
				if err != nil {
					glog.Errorf("ignore parseChapter err: %v, continue", err)
					continue
				}
				prefixTrims := []string{
					"wz1()",
					"『章节错误,点此举报』",
					"一秒记住【爱去小说网.】，为您提供精彩小说阅读。",
					"<>",
					"纯文字在线阅读本站域名手机同步阅读请访问",
					"%一%本%读%小说",
					"混血女主播直播后忘关摄像头",
					"私_生活视频遭曝光!!请关注微信公众号在线看：meinvmei222（长按三秒复制）！！",
					"泰国胸最女主播衣服都快包不住了视频在线看!!",
					"请关注微信公众号：meinvmei222（长按三秒复制）！！",
					"app2()",
					"app1()",
				}
				suffixTrims := []string{
					"(未完待续……)",
					"　　hp:手机用户请访问m.",
					"　　最快更新，无弹窗阅读请。",
					"『加入书签，方便阅读』",
					"最快更新，无弹窗阅读请。",
				}
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
	content = fmt.Sprintf("%s\n%s\n\n", title, content)
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
	content := doc.Find("div#chaptercontent").Text()
	return content, nil
}

func getChaperList(from, size int, catalogURL string) ([]chapterLink, error) {
	doc, err := goquery.NewDocument(catalogURL)
	if err != nil {
		panic(fmt.Sprintf("failed to catalogURL. err: %s", err))
	}

	result := make([]chapterLink, 0, size)
	chapList := doc.Find("div#chapterlist > p > a")
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
