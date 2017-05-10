package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

const (
	projectListEntryPoint   = "http://www.cpppc.org:8082/efmisweb/ppp/projectLibrary/getPPPList.do?tokenid=null"
	singalProjectEntryPoint = "http://www.cpppc.org:8082/efmisweb/ppp/projectLibrary/getProjInfo.do"

	totalPage      = 73
	Stage          = "PROJ_STATE_NAME"
	ProcessStage   = "执行阶段"
	UUIDKey        = "PROJ_RID"
	ProjectNameKey = "PROJ_NAME"
)

type projectList struct {
	List        []map[string]string `json:"list"`
	TotalCount  int                 `json:"totalCount"`
	CurrentPage int                 `json:"currentPage"`
	TotalPage   int                 `json:"totalPage"`
}
type project struct {
	Name string
	UUID string
}

func init() {
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("./logs")
	flag.Parse()
}

func main() {
	defer glog.Flush()

	// var badLink []string
	// projects := processingProject(projectListEntryPoint, totalPage)

	// glog.Infof("all the processing stage pro: %v", projects)
	// glog.Infoln("begin downloading...")
	// for _, uuid := range projects {
	// 	glog.Infof("downloading %s..", uuid)
	// 	link := fmt.Sprintf("%s?projId=%s", singalProjectEntryPoint, uuid)
	// 	fileName := fmt.Sprintf("%s.html", uuid)
	// 	if err := downloadFile(link, fileName); err != nil {
	// 		glog.Errorf("download page by link %s failed, saving to badLik..", link)
	// 		badLink = append(badLink, link)
	// 		continue
	// 	}
	// }

	// glog.Infoln("getting all table head..")
	// var unformatLink []string
	// // for _, uuid := range projects {

	// // }

	// glog.Infof("finished, badLink: %v", badLink)
	// getTableHead("testdata/project.html")
}

func fileName(uuid string) string {
	return fmt.Sprintf("%s.html", uuid)
}
func link(uuid string) string {
	return fmt.Sprintf("%s?projId=%s", singalProjectEntryPoint, uuid)
}
func processingProject(entryPointURL string, totalPage int) []project {
	var projects []project
	for i := 1; i > totalPage; i++ {
		glog.Infof("query %dsth page", i)
		formData := url.Values{
			"queryPage": {strconv.Itoa(i)},
			"induStr":   {"03"},
			"sortby":    {"proj_state"},
			"orderby":   {"asc"},
		}
		resp, err := http.PostForm(entryPointURL, formData)
		if err != nil {
			panic(fmt.Sprintf("query %dth page, request failed: %v", i, err))
		}
		list, err := getProjectList(resp)
		if err != nil {
			panic(err)
		}
		for _, pro := range list.List {
			if pro[Stage] == ProcessStage {
				uuid, ok := pro[UUIDKey]
				if !ok {
					panic(fmt.Sprintf("%s not found, project: %v", UUIDKey, pro))
				}
				project := project{
					UUID: uuid,
					Name: pro[ProjectNameKey],
				}
				projects = append(projects, project)
			}
		}
	}
	return projects
}
func getProjectList(resp *http.Response) (*projectList, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var list projectList
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

// getAllTableHead return th and unformatLink
func getAllTableHead(projects []project, unformatLink []string) ([]string, []string) {
	result := make([string]bool, 10)
	for _, pro := range projects {
		ths, err := getTableHead(fileName(pro.UUID))
		if err != nil {
			glog.Errorf("project %s(%s) parse failed: %v, save to unformatLink..", pro.Name, pro.UUID, err)
			unformatLink = append(unformatLink, link(pro.UUID))
			continue
		}
		for _, th := range ths {
			result[th] = true
		}
	}
	return result, unformatLink
}
func getTableHead(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("getTableHead: open file failed: %v", err))
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(fmt.Sprintf("NewDocumentFromReader failed: %v", err))
	}
	tbody := doc.Find(".wrap > .Over > .dist > .att > table.view_table")
	if tbody.Length() != 1 {
		return nil, fmt.Errorf("file %s not valid", fileName)
	}
	heads := tbody.Find("tbody > tr > td.view_field_label")
	th := make([]string, 0, heads.Length())
	for i := 0; i < heads.Length(); i++ {
		glog.V(2).Infof("the %dth iterm: %s", i, heads.Eq(i).Text())
		th = append(th, heads.Eq(i).Text())
	}
	return th, nil
}
func parseProject() {

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
