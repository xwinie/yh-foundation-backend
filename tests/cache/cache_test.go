package cache

import (
	"time"
	"testing"
	"yh-foundation-backend/cores"
	_ "yh-foundation-backend/tests"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

func TestRedisCache(t *testing.T) {
	timeoutDuration := 10 * time.Second
	if err := cores.GlobalCaches.Put("astaxie", "中文", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !cores.GlobalCaches.IsExist("astaxie") {
		t.Error("check err")
	}
	v, _ := redis.String(cores.GlobalCaches.Get("astaxie"), nil);
	beego.Trace("11", v)

}
