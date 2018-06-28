package auth

import (
	"encoding/json"
	"bytes"
	"github.com/astaxie/beego"
	"yh-foundation-backend/cores"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"yh-foundation-backend/tests"
	"yh-foundation-backend/filter"
	"yh-foundation-backend/auth"
	"net/url"
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
)

func TestResourcePost(t *testing.T) {

	values := map[string]string{"code": "101", "name": "测试", "action": "/v1/resource/*", "method": "get"}
	jsonValue, _ := json.Marshal(values)

	method := "POST"
	RequestURL := "/v1/resource"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue), timestamp)

	beego.Trace("testing", "TestResourcePost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 创建资源\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

func TestResourceGet(t *testing.T) {

	method := "GET"
	RequestURL := "/v1/resource/101"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, nil, timestamp)

	type Response struct {
		Code int
		Msg  string
		auth.SysResource
		cores.Links
	}
	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)
	beego.Trace("testing", "TestResourceGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 获取资源信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
			tests.Id = response.Id
		})
	})
}

func TestResourceByPageGet(t *testing.T) {

	values := make(url.Values)
	values.Add("p", "3")
	values.Add("perPage", "5")

	method := "GET"
	RequestURL := "/v1/resource"+"?"+values.Encode()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method,nil, RequestURL, timestamp)
	beego.Trace(111111,RequestURL)
	w := tests.Request(method, RequestURL, signature, nil, timestamp)

	beego.Trace("testing", "TestResourceGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 获取资源信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestResourcePut(t *testing.T) {

	values := map[string]interface{}{"name": "测试1"}
	jsonValue, _ := json.Marshal(values)

	method := "PUT"
	RequestURL := "/v1/resource/" + fmt.Sprint(tests.Id)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue), timestamp)

	beego.Trace("testing", "TestResourcePut", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 修改资源信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestResourceDelete(t *testing.T) {

	method := "DELETE"
	RequestURL := "/v1/resource/" + fmt.Sprint(tests.Id)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, nil, timestamp)

	beego.Trace("testing", "TestResourceDelete", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 删除资源信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestMenuByUserIdGet(t *testing.T) {

	o := orm.NewOrm()
	role, err := auth.ReadRoleByCode("1")
	if err != nil {
		beego.Trace("获取角色数据异常")
	}
	resource := []auth.SysResource{
		{Code: "100001", Action: "/user", Method: "", Name: "用户管理", ResType: 1},
	}
	o.InsertMulti(len(resource), resource)
	var resources []int64
	res1, _ := o.Raw("SELECT id resourceId "+
		"  from sys_resource where code in "+
		"('100001')", role.Id).QueryRows(&resources)

	beego.Trace("mysql row affected num:%d,%v ", res1, resources)
	roleResources := auth.RoleDistributorResource{ResourceId: resources}
	response := auth.RoleDistributorResourceService(role.Id, &roleResources)
	beego.Trace("init RoleResource Success ", response)
	var user auth.SysUser
	err1 := o.QueryTable("sys_user").Filter("account", "12345").One(&user)
	if err1 != nil {
		beego.Trace("获取用户数据异常")
	}
	method := "GET"
	RequestURL := "/v1/menus/" + fmt.Sprint(user.Id)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, nil, timestamp)

	beego.Trace("testing", "TestMenuByUserIdGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 获取用户菜单\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}
