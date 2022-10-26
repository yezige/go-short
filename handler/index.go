package handler

import (
	"net/http"
	"time"

	"github.com/yezige/go-short/logx"
	"github.com/yezige/go-short/redis"
	"github.com/yezige/go-short/request"
)

func Index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	logx.LogAccess.Infoln("Path: " + path)

	if path == "" {
		IsError(w, "short path empty")
		return
	}

	// 先查询 redis 缓存
	cache, err := redis.New().Get("goshort:path:cache:" + path)
	if err == nil {
		IsSuccess(w, string(cache))
		return
	}

	// 查询 redis
	search, err := redis.New().Get("goshort:path:" + path)
	if err == nil {
		// 获取url内容
		res, err := request.New(search).Get().GetBody()
		if err != nil {
			IsError(w, err)
			return
		}

		// 存储到redis
		if err := redis.New().SetTTL("goshort:path:cache:"+path, res, time.Hour*1); err != nil {
			IsError(w, "Add To Redis Error: "+"goshort:path:cache:"+path)
			return
		}

		IsSuccess(w, string(res))
		return
	}
	IsError(w, "Short Path: "+path+" Does Not Exist")
}
