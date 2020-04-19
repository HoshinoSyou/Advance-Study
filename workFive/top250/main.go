package main

import (
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

type movie struct {
	Rank       int    `json:"rank"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Director   string `json:"director"`
	Evaluation string `json:"evaluation"`
	Remark     string `json:"remark"`
}

var url string
var urls []string

func main() {
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	urls := Urls()
	for _, url := range urls {
		movies := Regexp(client, url)
		for _, movie := range movies {
			bytes, err := json.Marshal(movie)
			if err != nil {
				log.Printf("序列化失败喵！错误信息：", err)
				return
			}
			fmt.Println(string(bytes))
		}
	}
}

func Urls() []string {
	for i := 0; i < 10; i++ {
		if i == 0 {
			url = "https://movie.douban.com/top250"
		} else {
			url = "https://movie.douban.com/top250" + "?start=" + strconv.Itoa(25*i) + "&filter="
		}
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

func Regexp(client *http.Client, url string) (movies []movie) {
	res := Http(client, url)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("读取错误喵！错误信息：%v", err)
		return
	}
	b := strings.ReplaceAll(string(body), "\n", " ")
	olReg := regexp.MustCompile(`<ol class="grid_view">(.*?)</ol>`)
	if olReg == nil {
		log.Printf("解析ol失败喵！")
		return
	}
	olResult := olReg.FindAllStringSubmatch(b, -1)
	liReg := regexp.MustCompile(`<li>(.*?)</li>`)
	if liReg == nil {
		log.Printf("解析li失败喵！")
		return
	}
	liResult := liReg.FindAllString(olResult[0][0], -1)
	for _, v := range liResult {
		movies = append(movies, movie{
			Rank:       Rank(v),
			Name:       Name(v),
			Image:      Image(v),
			Director:   Director(v),
			Evaluation: Evaluation(v),
			Remark:     Remark(v),
		})
	}
	return movies
}

func Rank(v string) (rank int) {
	rankReg := regexp.MustCompile(`<em class="">(.*?)</em>`)
	rankSlice := rankReg.FindStringSubmatch(v)
	rank, _ = strconv.Atoi(rankSlice[1])
	return rank
}

func Name(v string) (name string) {
	nameReg := regexp.MustCompile(`<span class="title">(.*?)</span>`)
	nameSlice := nameReg.FindStringSubmatch(v)
	return nameSlice[1]
}

func Image(v string) (image string) {
	imageReg := regexp.MustCompile(`<img width="\d+" alt="(.*?)" src="(.*?)" class="">`)
	imageSlice := imageReg.FindStringSubmatch(v)
	return imageSlice[2]
}

func Director(v string) (director string) {
	val := strings.ReplaceAll(v, " ", "")
	textReg := regexp.MustCompile(`<pclass="">(.*?)<br>`)
	texts := textReg.FindStringSubmatch(val)
	textSlice := strings.Split(texts[1], "&nbsp;&nbsp;&nbsp;")
	directorSlice := strings.Split(textSlice[0], "导演:")
	return directorSlice[1]
}

func Evaluation(v string) (evaluation string) {
	evaluationReg := regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
	evaluationSlice := evaluationReg.FindStringSubmatch(v)
	return evaluationSlice[1]
}

func Remark(v string) (remark string) {
	remarkReg := regexp.MustCompile(`<span class="inq">(.*?)</span>`)
	remarkSlice := remarkReg.FindStringSubmatch(v)
	if remarkSlice == nil {
		return "无"
	}
	return remarkSlice[1]
}
