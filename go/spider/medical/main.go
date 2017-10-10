package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/zykzhang/handy"

	"github.com/mozhata/merr"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

const (
	baseEntry = "http://www.qxw18.com"
	maxPage   = 62
	// http://www.qxw18.com/company/list-areaid-1-62.html
	badURISummary  = "bad_uri.csv"
	medicalSummary = "medical.csv"
)

type badURI struct {
	url    string
	reason string
}

var (
	faildURIs []badURI
)

func init() {
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("./logs")
	flag.Parse()
}

func main() {
	defer glog.Flush()
	// download list page
	pageLists := downloadListPage()
	glog.Infof("pageLists: %s", handy.MarshalJSONOrDie(pageLists))

	// parse list page to get detail page entrypoint
	companyEntries := getCompanyEntries(pageLists)
	glog.Infof("total %d companies", len(companyEntries))

	// download detail page
	cmpLinks := downloadDetailPages(companyEntries)
	glog.Infof("%d cmpLinks downloaded", len(cmpLinks))

	// parse detail page to get infos
	titles := make(map[string]bool, 5)
	allProducts := make([]map[string]string, 0, len(cmpLinks))
	for _, url := range cmpLinks {
		filename := fmt.Sprintf("%s.html", getIDByURI(url))
		contents := parseDetailContent(filename)
		for key := range contents {
			titles[key] = true
		}
		contents["url"] = url
		contents["company"] = companyEntries[url]
		allProducts = append(allProducts, contents)
	}
	// save bad-url
	if len(faildURIs) > 0 {
		for _, u := range faildURIs {
			line := []string{u.url, u.reason}
			writeLine(line, badURISummary)
		}
	}
	// save product infos
	thead := make([]string, 0, len(titles))
	thead = append(thead, "company", "url")
	for k := range titles {
		thead = append(thead, k)
	}
	writeLine(thead, medicalSummary)
	for _, dic := range allProducts {
		line := make([]string, 0, len(thead))
		for _, key := range thead {
			line = append(line, dic[key])
		}
		writeLine(line, medicalSummary)
	}
}

func parseDetailContent(filename string) map[string]string {
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %s: %v", filename, err))
	}
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(fmt.Sprintf("NewDocumentFromReader failed: %v. file: %s", err, filename))
	}
	result := make(map[string]string, 2)

	contacts := doc.Find(".text_04 > div").Eq(1).Text()
	lines := strings.Split(contacts, "\n")
	for _, l := range lines {
		parts := strings.Split(l, "：")
		if len(parts) == 2 {
			glog.Infof("title: %s\tcontent: %s", parts[0], parts[1])
			result[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	divs := doc.Find(".text_02")
	for i := 0; i < divs.Length(); i++ {
		titile := divs.Eq(i).Find(".biaoti").Text()
		if strings.Contains(titile, "产品展示") {
			products := divs.Eq(i).Find("a")
			lines := make([]string, 0, products.Length())
			for i := 0; i < products.Length(); i++ {
				text := strings.TrimSpace(products.Eq(i).Text())
				href, _ := products.Eq(i).Attr("href")
				href = strings.TrimSpace(href)
				lines = append(lines, fmt.Sprintf("%s : %s", text, href))
				glog.Infof("text: %s\tlink: %s", text, href)
			}
			result["products"] = strings.Join(lines, "\n")
			break
		}
	}
	glog.Infof("result: %s", handy.MarshalJSONOrDie(result))

	return result
}

func downloadDetailPages(companyEntries map[string]string) []string {
	links := make([]string, 0, len(companyEntries))
	for link := range companyEntries {
		cmpID := getIDByURI(link)
		filename := fmt.Sprintf("%s.html", cmpID)
		err := downloadFile(link, filename)
		if err != nil {
			logErr(err)
			continue
		}
		links = append(links, link)
	}
	return links
}

func getCompanyEntries(pageListFiles []string) map[string]string {
	detailEntres := make(map[string]string, 500)
	for _, filename := range pageListFiles {
		f, err := os.Open(filename)
		if err != nil {
			panic(fmt.Sprintf("failed to parse %s", filename))
		}
		doc, err := goquery.NewDocumentFromReader(f)
		if err != nil {
			panic(fmt.Sprintf("NewDocumentFromReader failed: %v, file: %s", err, filename))
		}
		list := doc.Find(".qiye_box > div > a")
		for i := 0; i < list.Length(); i++ {
			href, found := list.Eq(i).Attr("href")
			if !found {
				continue
			}
			companyName := strings.TrimSpace(list.Eq(i).Text())
			if companyName == "" {
				continue
			}
			detailEntres[href] = companyName
		}
	}
	return detailEntres
}

func downloadListPage() []string {
	pageListEntries := make([]string, 0, maxPage)
	for i := 1; i <= maxPage; i++ {
		fileName := fmt.Sprintf("list-areaid-1-%d.html", i)
		pageListEntries = append(pageListEntries, fileName)
		entry := fmt.Sprintf("%s/company/%s", baseEntry, fileName)
		glog.Infof("entry: %s", entry)
		err := downloadFile(entry, fileName)
		if err != nil {
			logErr(err)
		}
	}
	return pageListEntries
}

func downloadFile(url, filename string) error {
	res, err := http.Get(url)
	if err != nil {
		makrBadURL(url, err.Error())
		return merr.WrapErr(err)
	}
	if res.StatusCode != http.StatusOK {
		makrBadURL(url, fmt.Sprintf("status code: %d", res.StatusCode))
		return merr.NotFoundError(nil, "url %s status code is %d", url, res.StatusCode)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		makrBadURL(url, err.Error())
		return merr.WrapErr(err)
	}
	str, err := doc.Html()
	if err != nil {
		panic(fmt.Sprintf("failed to download file by url %s. err: %v", url, err))
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

func makrBadURL(url, reason string) {
	faildURIs = append(faildURIs, badURI{
		url:    url,
		reason: reason,
	})
}

func logErr(err error) {
	e := merr.WrapErr(err)
	glog.Infof("err: %s\n origin err: %s\ncallstack: %s",
		e.Message,
		e.RawErr,
		e.CallStack(),
	)
}

func writeLine(line []string, filename string) {
	if len(line) == 0 {
		glog.Infoln("get an empty line, ignore it")
		return
	}
	for _, i := range line {
		if len(i) < 1 {
			i = " "
		}
	}
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)
	r2.Write(line)
	r2.Flush()
	fout, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer fout.Close()
	if err != nil {
		panic(fmt.Sprintln(filename, err))
	}
	fout.WriteString(buf.String())
	glog.Infoln("writeLine:writen")
}

func getIDByURI(href string) string {
	parts := strings.Split(href, "/")
	return strings.TrimSpace(parts[len(parts)-1])
}
