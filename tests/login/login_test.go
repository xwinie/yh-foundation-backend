package login

import (
	"testing"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	_ "yh-foundation-backend/tests"
	"encoding/json"
	"bytes"
	"yh-foundation-backend/filter"
	"yh-foundation-backend/tests"
	"yh-foundation-backend/cores"
	"yh-foundation-backend/auth"
	"time"
)

func TestLogin(t *testing.T) {

	values := map[string]string{"UserName": "12345", "Password": cores.Md5(cores.Sha1("12345") + cores.Sha1("Password"))}
	jsonValue, _ := json.Marshal(values)

	method := "POST"
	RequestURL := "/v1/login"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)

	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue),timestamp)

	beego.Trace("testing", "TestLogin", "Code[%d]\n%s", w.Code, w.Body.String())
	type Response struct {
		Token string
	}
	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)
	Convey("Subject: 用户登录\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}
