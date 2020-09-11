package Services

import (
	"context"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r * http.Request)(interface{}, error){
	if r.URL.Query().Get("uid") != ""{
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
		return UserRequest{Uid:uid,}, nil
	}
	return nil.errors.New("参数错误")
}
