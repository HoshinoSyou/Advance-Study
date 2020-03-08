package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"strings"
)

func CheckToken(token string) (username string, err error) {
	split := strings.Split(token, ".")
	if len(split) != 3 {
		err = errors.New("token构建错误")
		log.Println(err)
		return
	}
	_, err = base64.StdEncoding.DecodeString(split[0])
	if err != nil {
		err = errors.New("header解析错误")
		log.Println(err)
		return
	}
	p, err := base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		err = errors.New("payload解析错误")
		log.Println(err)
		return
	}
	_, err = base64.StdEncoding.DecodeString(split[2])
	if err != nil {
		err = errors.New("signature解析错误")
		log.Println(err)
		return
	}
	var payload Payload
	err = json.Unmarshal(p, &payload)
	if err != nil {
		log.Printf("反序列化失败喵！错误信息：%v", err)
		return
	}
	username = payload.Username
	return
}
