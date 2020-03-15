package service

import (
	"Advance/workTwo/response"
	"Advance/workTwo/sqlConn"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var id []string
var b UserData

func Vote(v *gin.Context) {
	b.Username = v.Param("username")
	rdb = sqlConn.Redis()
	db = sqlConn.Mysql()
	db.Where("username =?", b.Username).First(&b)
	t := time.Now()
	h := t.Hour()
	m := t.Minute()
	s := t.Second()
	if h == 0 && m == 0 && s == 0 {
		db.Model(&b).Update("Ballot", 3)
	}
	Ballot(v)
}

func Ballot(v *gin.Context) {
	rdb = sqlConn.Redis()
	for i := 0; i <= 2; i++ {
		id = append(id, v.Param("id"))
		b, err := rdb.HGet(id[i], "Ballot").Result()
		if err != nil {
			response.Error(v, 212, "投票失败喵！")
			return
		} else {
			ballot, _ := strconv.Atoi(b)
			rdb.HDel(id[i], "Ballot")
			rdb.HSet(id[i], "Ballot", ballot+1)
		}
	}
	response.OK(v, "投票成功喵！")
}
