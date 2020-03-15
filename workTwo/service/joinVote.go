package service

import (
	"Advance/workTwo/response"
	"Advance/workTwo/sqlConn"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Candidate struct {
	gorm.Model
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	Age          int    `json:"age"`
	Introduction string `json:"introduction"`
	Ballot       int    `json:"ballot"`
}

var rdb *redis.Client

func Join(j *gin.Context) {
	rdb = sqlConn.Redis()
	username := j.Param("username")
	c := Candidate{
		Model:        gorm.Model{},
		Name:         j.PostForm("name"),
		Sex:          j.PostForm("sex"),
		Age:          j.GetInt("age"),
		Introduction: j.PostForm("introduction"),
		Ballot:       0,
	}
	rdb.HMSet(string(c.ID), map[string]interface{}{
		"Name":         c.Name,
		"Sex":          c.Sex,
		"Age":          c.Age,
		"Introduction": c.Introduction,
		"Ballot":       c.Ballot,
	})
	response.OK(j, "参加比赛成功喵！")
}
