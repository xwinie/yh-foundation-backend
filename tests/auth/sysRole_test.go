package auth

import (
	"encoding/json"
	"bytes"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"yh-foundation-backend/tests"
	"yh-foundation-backend/auth"
	"github.com/astaxie/beego/orm"
	"fmt"
	"yh-foundation-backend/filter"
	"yh-foundation-backend/cores"
	"net/url"
	"time"
)

var roleID int64

func TestRolePost(t *testing.T) {

	values := map[string]string{"code": "101", "name": "测试", "description": "哈哈角色"}
	jsonValue, _ := json.Marshal(values)

	method := "POST"
	RequestURL := "/v1/role"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue),timestamp)

	beego.Trace("testing", "TestRolePost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 创建角色\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

func TestRoleGetByCode(t *testing.T) {

	method := "GET"
	RequestURL := "/v1/role/101"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	beego.Trace("testing", "TestRoleGetByCode", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 获取角色信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
			type Response struct {
				Code int
				Msg  string
				auth.SysResource
				cores.Links
			}
			var response Response
			json.Unmarshal(w.Body.Bytes(), &response)
			roleID = response.Id
		})
	})
}

func TestRoleByPageGet(t *testing.T) {

	values := make(url.Values)
	values.Add("p", "3")
	values.Add("perPage", "5")

	method := "GET"
	RequestURL := "/v1/role"+"?"+values.Encode()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	beego.Trace("testing", "TestRoleByPageGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 分页获取角色信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}
func TestRolePut(t *testing.T) {

	values := map[string]interface{}{"name": "测试1"}
	jsonValue, _ := json.Marshal(values)

	method := "PUT"
	RequestURL := "/v1/role/" + fmt.Sprintf("%d", roleID)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue),timestamp)

	beego.Trace("testing", "TestRolePut", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 修改角色信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestRoleDelete(t *testing.T) {

	method := "DELETE"
	RequestURL := "/v1/role/" + fmt.Sprintf("%d", roleID)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	beego.Trace("testing", "TestRoleDelete", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 删除角色信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestRoleDistributorResource(t *testing.T) {

	resource := new(auth.SysResource)
	resource.Code = "2"
	resource.Name = "删除角色"
	resource.Action = "/v1/role/*"
	resource.Method = "DELETE"
	resource2 := new(auth.SysResource)
	resource2.Code = "2"
	resource2.Name = "删除角色"
	resource2.Action = "/v1/role/*"
	resource2.Method = "DELETE"
	role := new(auth.SysRole)
	role.Code = "2"
	role.Name = "管理员"
	o := orm.NewOrm()
	resourceId, _ := o.Insert(resource)
	resourceId2, _ := o.Insert(resource)
	roleId, _ := o.Insert(role)
	roleID = roleId

	values := map[string]interface{}{"resourceId": []int64{resourceId, resourceId2}}
	jsonValue, _ := json.Marshal(values)

	method := "POST"
	RequestURL := "/v1/role/" + fmt.Sprint(roleId) + "/resource"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue),timestamp)

	beego.Trace("testing", "TestRoleDistributorResource", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 角色分配资源\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

func TestRoleGetResourceByRoleId(t *testing.T) {

	method := "GET"
	RequestURL := "/v1/role/" + fmt.Sprint(roleID) + "/resource"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)
	w := tests.Request(method, RequestURL, signature, nil,timestamp)

	beego.Trace("testing", "TestRoleGetResourceByRoleId", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 获取角色对应的资源\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}
