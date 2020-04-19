package main

import (
	"Advance/Crawler/classTable/dao_sql"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var url string
var urls []string

func main() {
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	urls := Urls()
	for sno := 2019210001; sno < 2019215999; sno++ {
		/*go func() {*/
		lessons, res := Regexp(client, urls[sno-2019210001])
		if res.StatusCode != http.StatusOK {
			log.Printf("此学号不存在喵！")
			return
		}
		for range lessons {
			var t = dao_sql.Table{
				Sno:     sno,
				Lessons: lessons,
			}
			dao_sql.AddInSql(t)
			bytes, err := json.Marshal(t)
			if err != nil {
				log.Printf("序列化失败喵！错误信息：", err)
				return
			}
			fmt.Println(string(bytes))
		}
		/*}()*/
	}
}

func Urls() []string {
	for sno := 2019210001; sno < 2019215999; sno++ {
		url = "http://jwc.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + strconv.Itoa(sno)
		urls = append(urls, url)
	}
	return urls
}

func Http(client *http.Client, url string) (res *http.Response) {
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		log.Printf("建立请求失败喵！错误信息:%v", err1)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	res, err2 := client.Do(req)
	if err2 != nil {
		log.Printf("接收响应失败喵！错误信息：%v", err2)
		return
	}
	return res
}

func Regexp(client *http.Client, url string) (lessons []dao_sql.Lesson, res *http.Response) {
	res = Http(client, url)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("找不到此学号页面喵！")
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("读取错误喵！错误信息：%v", err)
		return
	}
	newBody := strings.ReplaceAll(string(body), "\n", " ")
	tableReg := regexp.MustCompile(`<div id="kbStuTabs-list" aria-labelledby="ui-id-2" class="ui-tabs-panel ui-widget-content ui-corner-bottom" role="tabpanel" aria-hidden="false" style="display: block;">(.*?)</div>`)
	tableResult := tableReg.FindAllStringSubmatch(newBody, -1)
	tbodyReg := regexp.MustCompile(`<tbody>(.*?)</tbody>`)
	tbodyResult := tbodyReg.FindAllStringSubmatch(tableResult[0][0], -1)
	trReg := regexp.MustCompile(`<tr>(.*?)</tr>`)
	trResult := trReg.FindAllString(tbodyResult[0][0], -1)
	for i := 0; i < len(trResult); i++ {
		typeReg := regexp.MustCompile(`<td rowspan="\d+" align="center">(.*?)</td>`)
		typeResult := typeReg.FindStringSubmatch(trResult[i])
		if typeResult != nil {
			idName, class, require := TdRowspan(trResult[i])
			teacher, time, place := Td(trResult[i])
			lessons = append(lessons, dao_sql.Lesson{
				Name:          Name(idName),
				Id:            Id(idName),
				Type:          Type(trResult[i]),
				Class:         class,
				Required:      require,
				Teacher:       teacher,
				TimeAndPlaces: TimeAndPlace(time, place),
			})
		}
	}
	return lessons, res
}

func Type(v string) (Type string) {
	typeReg := regexp.MustCompile(`<td rowspan="\d+" align="center">(.*?)</td>`)
	typeResult := typeReg.FindStringSubmatch(v)
	return typeResult[2]
}

func Td(v string) (teacher string, time string, place string) {
	tdReg := regexp.MustCompile(`<td>(.*?)</td>`)
	tdResult := tdReg.FindAllStringSubmatch(v, -1)
	return tdResult[0][1], tdResult[1][1], tdResult[2][1]
}

func TimeAndPlace(time string, place string) (timeAndPlaces []dao_sql.TimeAndPlace) {
	timeSlice1 := strings.Split(time, "第")
	timeSlice2 := strings.Split(timeSlice1[1], " ")
	timeAndPlaces = append(timeAndPlaces, dao_sql.TimeAndPlace{
		Day:     timeSlice1[0],
		Section: timeSlice2[0],
		Week:    timeSlice2[1],
		Place:   place,
	})
	return timeAndPlaces
}

func TdRowspan(v string) (idName string, class string, require string) {
	tdRowspanReg := regexp.MustCompile(`<td rowspan="\d+">(.*?)</td>`)
	tdRowspanResult := tdRowspanReg.FindAllStringSubmatch(v, -1)
	return tdRowspanResult[0][2], tdRowspanResult[1][2], tdRowspanResult[2][2]
}

func Id(idName string) (id string) {
	idNameSlice := strings.Split(idName, "-")
	return idNameSlice[0]
}
func Name(idName string) (name string) {
	idNameSlice := strings.Split(idName, "-")
	if len(idNameSlice) > 2 {
		name = idNameSlice[1] + idNameSlice[2]
	} else {
		name = idNameSlice[1]
	}
	return name
}
