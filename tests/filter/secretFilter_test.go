package filter

import (
	"bytes"
	"encoding/json"
	"net/url"
	"testing"
	"yh-foundation-backend/auth"
	"yh-foundation-backend/filter"
	"yh-foundation-backend/tests"
	_ "yh-foundation-backend/tests"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"
	"yh-foundation-backend/cores"
	"time"
)

func TestSignatureGet(t *testing.T) {

	role := new(auth.SysRole)
	role.Code = "1"
	role.Name = "管理员"

	o := orm.NewOrm()

	roleId, _ := o.Insert(role)
	if roleId < 0 {
		beego.Trace("添加数据失败")
	}
	method := "GET"
	RequestURL := "/v1/role/1"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	beego.Trace("testing", "TestSignatureGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Get方式校验签名\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestSignaturePost(t *testing.T) {

	values := map[string]string{"code": "2", "name": "测试", "description": "哈哈角色"}
	jsonValue, _ := json.Marshal(values)

	method := "POST"
	RequestURL := "/v1/role"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL,timestamp)

	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue),timestamp)

	beego.Trace("testing", "TestSignaturePost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Post校验签名\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

func TestSignatureGetParams(t *testing.T) {

	values := make(url.Values)
	values.Add("k", "1")

	method := "GET"
	RequestURL := "/v1/role/1"+"?"+values.Encode()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	type Response struct {
		Code int
		Msg  string
		auth.SysRole
		cores.Links
	}
	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)
	beego.Trace("testing", "TestSignatureGetParams", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 获取角色信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
			tests.Id = response.Id
		})
	})
}

func TestSignatureDelete(t *testing.T) {

	values := map[string]string{"code": "2", "name": "测试", "description": "哈哈角色"}
	jsonValue, _ := json.Marshal(values)

	method := "DELETE"
	RequestURL := "/v1/role"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue),timestamp)

	beego.Trace("testing", "TestSignatureDelete", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 删除用户角色\n", t, func() {
		Convey("Status Code Should Be 401", func() {
			So(w.Code, ShouldEqual, 401)
		})
	})
}
