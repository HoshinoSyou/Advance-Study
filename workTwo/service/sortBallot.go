package service

import (
	"Advance/workTwo/response"
	"Advance/workTwo/sqlConn"
	"github.com/gin-gonic/gin"
	"sort"
	"strconv"
)

type candidateSort struct {
	Name   string `json:"string"`
	Ballot int    `json:"ballot"`
}

var ballotSlice []candidateSort

func Sort(c *gin.Context) {
	rdb = sqlConn.Redis()
	for id := 0; id >= 0; id++ {
		name, err := rdb.HGet(string(id), "Name").Result()
		b, err := rdb.HGet(string(id), "Ballot").Result()
		ballot, err := strconv.Atoi(b)
		if err != nil {
			response.Error(c, 213, "查询排行榜失败喵！")
			return
		}
		c := candidateSort{
			Name:   name,
			Ballot: ballot,
		}
		ballotSlice = append(ballotSlice, c)
		sort.Slice(ballotSlice, func(i, j int) bool {
			if ballotSlice[i].Ballot > ballotSlice[j].Ballot {
				return true
			}
			return false
		})
		response.OKWithData(b, "目前比赛名次如下喵：\n", ballotSlice)
	}
}
