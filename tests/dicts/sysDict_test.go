package dicts

import (
	"encoding/json"
	"testing"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"

	"bytes"
	"yh-foundation-backend/tests"
	"yh-foundation-backend/filter"
	"net/url"
	"yh-foundation-backend/auth"
	"fmt"
	"yh-foundation-backend/dicts"
	"time"
)

func TestDictPost(t *testing.T) {

	values := map[string]interface{}{"code": "100", "City": "5101", "Name": "性别", "Level": 0, "ParentId": -1, "ParentCode": "-1", "Type": "C1"}

	jsonValue, _ := json.Marshal(values)
	method := "POST"
	RequestURL := "/v1/dict"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue),timestamp)

	//beego.Trace("testing", "TestDictPost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 创建数据字典\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

func TestDictGetByPage(t *testing.T) {

	values := make(url.Values)
	values.Add("p", "3")
	values.Add("perPage", "5")


	method := "GET"
	RequestURL := "/v1/dict"+"?"+values.Encode()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method,nil, RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	beego.Trace("testing", "TestDictGetByPage", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 分页获取用户信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})

}

func TestDictGetByPageAndType(t *testing.T) {

	values := make(url.Values)
	values.Add("p", "3")
	values.Add("perPage", "5")
	values.Add("type", "c1")

	method := "GET"
	RequestURL := "/v1/dict"+"?"+values.Encode()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method,nil, RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	beego.Trace("testing", "TestDictGetByPageAndType", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 分页获取用户信息\n", t, func() {
		Convey("Status Code Should Be 400", func() {
			So(w.Code, ShouldEqual, 400)
		})
	})

}

func TestDictPut(t *testing.T) {
	dict, _ := dicts.ReadDict("100")
	values := map[string]interface{}{"name": "测试1"}
	jsonValue, _ := json.Marshal(values)
	method := "PUT"
	RequestURL := "/v1/dict/" + fmt.Sprintf("%d", dict.Id)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL,timestamp)

	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue),timestamp)

	beego.Trace("testing", "TestDictPut", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 修改数据字典信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestDictGet(t *testing.T) {
	dict, _ := dicts.ReadDict("100")

	method := "GET"
	RequestURL := "/v1/dict/" + fmt.Sprintf("%d", dict.Id)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	beego.Trace("testing", "TestDictGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 获取数据字典\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestDictDelete(t *testing.T) {
	dict, _ := dicts.ReadDict("100")
	method := "DELETE"
	RequestURL := "/v1/dict/" + fmt.Sprintf("%d", dict.Id)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)

	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	beego.Trace("testing", "TestDictDelete", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 删除客户端信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}
