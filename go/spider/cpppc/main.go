package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

const (
	projectListEntryPoint = "http://www.cpppc.org:8082/efmisweb/ppp/projectLibrary/getPPPList.do?tokenid=null"
	projectBaseRUL        = "http://www.cpppc.org:8082/efmisweb/ppp/projectLibrary/"
	topLevelEntry         = "getProjInfoNational.do"
	otherLevelEntry       = "getProjInfo.do"

	totalPage      = 1536
	supposedPage   = 235
	Stage          = "PROJ_STATE_NAME"
	ProcessStage   = "执行阶段"
	UUIDKey        = "PROJ_RID"
	ProjectNameKey = "PROJ_NAME"
	ProjectLevel   = "PROJ_LEVEL"
)

// file
const (
	rawProjectIndex = "projectindexs.csv"

	// title and link
	allProjectLinks = "all_links.csv"
	sucessLinks     = "success_links.csv"
	failedLinks     = "failed_links.csv"
	selectedLinks   = "selected_links.csv"

	allProjects      = "all_projects.csv"
	trimedProjects   = "trimed_projects.csv"
	selectedProjects = "selected_projects.csv"
	relatedProjects  = "related_projects.csv"
	othersProjects   = "other_projects.csv"
	// all contents

)

var (
	groupKeyWords = []string{
		"水利",
		"生态保护",
		"河道",
		"废水",
		"生态区",
		"河流",
		"污水处理",
	}
)

type projectDetail struct {
	Name       string
	Link       string
	Region     string
	Industry   string
	Qixian     string
	Method     string
	Reply      string
	StartTime  string
	Desc       string
	TotalMoney string
}

func (d *projectDetail) toLine() []string {
	return []string{
		d.Name,
		d.Link,
		d.Region,
		d.Industry,
		d.Qixian,
		d.Method,
		d.Reply,
		d.StartTime,
		d.Desc,
		d.TotalMoney,
	}
}

type projectList struct {
	List        []map[string]interface{} `json:"list"`
	TotalCount  int                      `json:"totalCount"`
	CurrentPage int                      `json:"currentPage"`
	TotalPage   int                      `json:"totalPage"`
}
type project struct {
	Name  string
	UUID  string
	Stage string
	Level string
}

func (p project) toLine() []string {
	return []string{p.Name, p.UUID, p.Stage, p.Level}
}
func (p project) linkLine() []string {
	return []string{p.Name, link(p.UUID, must2Int(p.Level))}
}

func init() {
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("./logs")
	flag.Parse()
}

func main() {
	defer glog.Flush()

	// projects := processingProject(projectListEntryPoint)

	// glog.Infof("all the processing stage pro: %v", projects)
	// saveIndexs(projects, "projectindexs.csv")
	// generateAllLinks(rawProjectIndex, allProjectLinks)
	// download file
	// ReadLine(allProjectLinks, downloadByLink)

	// parse local html
	// ReadLine(sucessLinks, parseProjectBySuccessLine)
	// trim space to regin
	// ReadLine(allProjects, trimSpace)
	// grop projects
	groupProjects()
}
func ReadLine(fileName string, handler func([]string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	reader := csv.NewReader(f)
	// count := 0
	for {
		// if count > 10 {
		// 	return nil
		// }
		// count++
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			glog.Errorf("err is: %v", err)
			return err
		}
		handler(line)
	}
	return nil
}
func groupProjects() {
	selected := make(map[string][]string)
	selectedOrder := make([]string, 0, 40)
	getSelectedMapping := func(line []string) {
		if len(line) == 0 {
			glog.Infoln("get an empty line, ignore it")
			return
		}
		if len(line) != 2 {
			panic(fmt.Sprintf("line should hase %v item, this: %v", 2, len(line)))
		}
		title := strings.TrimSpace(line[0])
		selected[title] = []string{}
		selectedOrder = append(selectedOrder, title)
	}
	// populate selected
	ReadLine(selectedLinks, getSelectedMapping)
	// group
	groupHandler := func(line []string) {
		if len(line) == 0 {
			glog.Infoln("get an empty line, ignore it")
			return
		}
		detail := projectDetail{}
		length := len(detail.toLine())
		if len(line) != length {
			panic(fmt.Sprintf("like line should hase %v item, this: %v", length, line))
		}
		title := strings.TrimSpace(line[0])
		if _, found := selected[title]; found {
			// save to selected when finish travel all projects
			selected[title] = line
			return
		}
		if matchProject(title) {
			// save to related
			writeLine(line, relatedProjects)
			return
		}
		// save to others
		writeLine(line, othersProjects)
	}
	ReadLine(trimedProjects, groupHandler)
	// write selected to file
	for _, title := range selectedOrder {
		writeLine(selected[title], selectedProjects)
	}
}
func matchProject(proStr string) bool {
	for _, kw := range groupKeyWords {
		if strings.Contains(proStr, kw) {
			return true
		}
	}
	return false
}
func trimSpace(line []string) {
	if len(line) == 0 {
		glog.Infoln("get an empty line, ignore it")
		return
	}
	detail := projectDetail{}
	length := len(detail.toLine())
	if len(line) != length {
		panic(fmt.Sprintf("like line should hase %v item, this: %v", length, line))
	}
	// region
	regionIndex := 2
	parts := strings.Split(line[regionIndex], "-->")
	for i, _ := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	line[regionIndex] = strings.Join(parts, "->")
	writeLine(line, trimedProjects)
}
func generateAllLinks(rawFile, destFile string) {
	f, err := os.Open(rawFile)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(f)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				return
			}
			glog.Fatalf("err is: %v", err)
			panic(err)
		}
		if len(line) == 0 {
			glog.Infoln("get an empty line, ignore it")
			return
		}
		linkLine := []string{line[0], link(line[1], must2Int(line[3]))}
		writeLine(linkLine, destFile)
	}
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

func downloadByLink(line []string) {
	if len(line) == 0 {
		glog.Infoln("get an empty line, ignore it")
		return
	}
	if len(line) != 2 {
		panic(fmt.Sprintf("like line should hase two item, this: %v", line))
	}
	uuid := strings.Split(line[1], "projId=")[1]
	err := downloadFile(line[1], fileName(uuid))
	if err != nil {
		badRecord := []string{line[0], line[1], uuid, err.Error()}
		writeLine(badRecord, failedLinks)
	}
	record := []string{line[0], line[1], uuid, strconv.Itoa(http.StatusOK)}
	writeLine(record, sucessLinks)
}
func parseProjectBySuccessLine(line []string) {
	if len(line) == 0 {
		glog.Infoln("get an empty line, ignore it")
		return
	}
	if len(line) != 4 {
		panic(fmt.Sprintf("like line should hase 4 item, this: %v", line))
	}
	detail := &projectDetail{
		Name: line[0],
		Link: line[1],
	}
	detail.parseProject(fileName(line[2]))
	writeLine(detail.toLine(), allProjects)
}

func fileName(uuid string) string {
	return fmt.Sprintf("%s.html", uuid)
}
func link(uuid string, level int) string {
	if level < 1 {
		panic(fmt.Sprintf("level < 1: %d", level))
	}
	entryPoint := otherLevelEntry
	if level == 1 {
		entryPoint = topLevelEntry
	}
	return fmt.Sprintf("%s%s?projId=%s", projectBaseRUL, entryPoint, uuid)
}
func saveIndexs(indexs []project, raw string) {
	fRaw, err := os.OpenFile(raw, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(fmt.Sprintln(raw, err))
	}
	defer fRaw.Close()

	bufRaw := new(bytes.Buffer)
	wRaw := csv.NewWriter(bufRaw)
	for _, p := range indexs {
		line := p.toLine()
		for _, i := range line {
			if len(i) < 1 {
				i = " "
			}
		}
		wRaw.Write(line)
	}
	wRaw.Flush()
	fRaw.WriteString(bufRaw.String())

	glog.Infoln("saveIndexs:writen")
}
func processingProject(entryPointURL string) []project {
	var projects []project
	for page := 1; page < totalPage; page++ {
		if page > supposedPage {
			glog.Fatalf("page %d not supposed to larger than %d", page, supposedPage)
		}
		glog.Infof("query %dsth page", page)
		formData := url.Values{
			"queryPage": {strconv.Itoa(page)},
			"sortby":    {"proj_state"},
			"orderby":   {"desc"},
		}
		resp, err := http.PostForm(entryPointURL, formData)
		if err != nil {
			panic(fmt.Sprintf("query %dth page, request failed: %v", page, err))
		}
		list, err := getProjectList(resp)
		if err != nil {
			panic(err)
		}
		for _, pro := range list.List {
			stage := mustGet(Stage, pro)
			if stage != ProcessStage {
				glog.Infof("\n\nat page %d foud other stage: %s, stop.", page, stage)
				return projects
			}
			uuid := mustGet(UUIDKey, pro)
			title := mustGet(ProjectNameKey, pro)
			levelStr := mustGet(ProjectLevel, pro)
			item := project{
				UUID:  uuid,
				Name:  title,
				Stage: stage,
				Level: levelStr,
			}
			projects = append(projects, item)
		}
	}
	return projects
}
func mustGet(key string, dict map[string]interface{}) string {
	val, ok := dict[key]
	if !ok {
		panic(fmt.Sprintf("key %s not found in dict %v", key, dict))
	}
	return val.(string)
}
func must2Int(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}
func getProjectList(resp *http.Response) (*projectList, error) {
	if resp.StatusCode != http.StatusOK {
		glog.Fatalf("response not normal")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var list projectList
	err = json.Unmarshal(body, &list)
	if err != nil {
		glog.Infof("unmarshal body failed: %v", err)
		return nil, err
	}
	return &list, nil
}

func (d *projectDetail) parseProject(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("parseProject: open file failed: %v", err))
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(fmt.Sprintf("NewDocumentFromReader failed: %v", err))
	}
	viewTable := doc.Find(".wrap > .Over > .dist > .att > table.view_table")
	if viewTable.Length() != 1 {
		// parse national
		glog.Infof("%s is national project, will use national method", fileName)
		d.parseNational(doc)
	} else {
		// parse normal
		d.parseNormal(viewTable)
	}
	return
}
func (d *projectDetail) parseNormal(table *goquery.Selection) {
	trs := table.Find("tbody > tr")
	for i := 0; i < trs.Length(); i++ {
		label := trs.Eq(i).ChildrenFiltered("td.view_field_label").Eq(0).Text()
		value := trs.Eq(i).ChildrenFiltered("td.view_field_con").Eq(0).Text()
		switch strings.TrimSpace(label) {
		case "所在地区":
			d.Region = value
		case "所属行业":
			d.Industry = value
		case "合作期限":
			d.Qixian = value
		case "运作方式":
			d.Method = value
		case "回报机制":
			d.Reply = value
		case "发起时间":
			d.StartTime = value
		case "项目概况":
			d.Desc = value
		case "项目总投资":
			d.TotalMoney = value
		}
	}
}
func (d *projectDetail) parseNational(doc *goquery.Document) {
	tds := doc.Find("table.view_table > tbody > tr >td")
	for i := 0; i < tds.Length(); i++ {
		if strings.TrimSpace(tds.Eq(i).Text()) == "所在地区" {
			d.Region = tds.Eq(i + 1).Text()
			continue
		}
		if label := tds.Eq(i).ChildrenFiltered("b"); label.Length() > 0 {
			key := strings.TrimSpace(label.Eq(0).Text())
			value := tds.Eq(i + 1).Text()
			switch key {
			case "所属行业":
				d.Industry = value
			case "回报机制":
				d.Reply = value
			case "发起时间":
				d.StartTime = value
			case "项目总投资":
				d.TotalMoney = value
			default:
				if strings.Contains(key, "项目运作方式(") {
					d.Method = value
				} else if strings.Contains(key, "合作期限") {
					d.Qixian = value
				} else if strings.Contains(key, "项目概况") {
					d.Desc = value
				}
				continue
			}
		}
	}
}

func downloadFile(url, filename string) error {
	res, e := http.Get(url)
	if e != nil {
		return e
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("res status code %d by url %s", res.StatusCode, url)
	}

	doc, err := goquery.NewDocumentFromResponse(res)
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
	glog.Infof("file %s saved", filename)
}
