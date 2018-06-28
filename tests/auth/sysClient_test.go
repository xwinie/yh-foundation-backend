
package auth

import (
	"encoding/json"
	"testing"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"

	"bytes"
	"yh-foundation-backend/cores"
	"yh-foundation-backend/tests"
	"fmt"
	"yh-foundation-backend/filter"
	"yh-foundation-backend/auth"
	"time"
)

func TestClientPost(t *testing.T) {
	values := map[string]string{"clientId": "test", "name": "测试", "secret": cores.RandStringByLen(10)}
	jsonValue, _ := json.Marshal(values)
	method := "POST"
	RequestURL := "/v1/client"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue),timestamp)

	beego.Trace("testing", "TestClientPost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 创建客户端\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

func TestClientGet(t *testing.T) {

	method := "GET"
	RequestURL := "/v1/client/test"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	type Response struct {
		Code int
		Msg  string
		auth.SysClient
	}
	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)
	beego.Trace("testing", "TestClientGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 获取用户信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
			tests.Id = response.Id
		})
	})
}

func TestClientPut(t *testing.T) {
	values := map[string]interface{}{"clientId": "test1", "name": "测试1"}
	jsonValue, _ := json.Marshal(values)
	method := "PUT"
	RequestURL := "/v1/client/" + fmt.Sprintf("%d", tests.Id)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL,timestamp)

	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue),timestamp)

	beego.Trace("testing", "TestClientPut", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 修改客户端信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestClientDelete(t *testing.T) {

	method := "DELETE"
	RequestURL := "/v1/client/" + fmt.Sprintf("%d", tests.Id)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)

	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	beego.Trace("testing", "TestClientDelete", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 删除客户端信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}
