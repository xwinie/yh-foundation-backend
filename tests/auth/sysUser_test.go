package auth

import (
	"encoding/json"
	"bytes"
	"github.com/astaxie/beego"
	"yh-foundation-backend/cores"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"yh-foundation-backend/tests"
	"github.com/astaxie/beego/orm"
	"yh-foundation-backend/auth"
	"fmt"
	"yh-foundation-backend/filter"
	"net/url"
	"time"
)

func TestUserPost(t *testing.T) {

	values := map[string]string{"account": "123456", "name": "测试", "password": cores.Md5(cores.Sha1("123456") + cores.Sha1("Password"))}
	jsonValue, _ := json.Marshal(values)

	method := "POST"
	RequestURL := "/v1/user"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue), timestamp)

	beego.Trace("testing", "TestUserPost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 创建用户\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

func TestUserGet(t *testing.T) {

	method := "GET"
	RequestURL := "/v1/user/123456"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, nil, timestamp)

	beego.Trace("testing", "TestUserGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 获取用户信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)

			type Response struct {
				Code int
				Msg  string
				auth.SysUser
				cores.Links
			}
			var response Response
			json.Unmarshal(w.Body.Bytes(), &response)
			tests.Id = response.Id
		})
	})
}

func TestUserByPageGet(t *testing.T) {

	values := make(url.Values)
	values.Add("p", "3")
	values.Add("perPage", "5")

	method := "GET"
	RequestURL := "/v1/user" + "?" + values.Encode()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, nil, timestamp)

	beego.Trace("testing", "TestUserByPageGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 分页获取用户信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestUserPut(t *testing.T) {

	values := map[string]interface{}{"name": "测试1"}
	jsonValue, _ := json.Marshal(values)

	method := "PUT"
	RequestURL := "/v1/user/" + fmt.Sprintf("%d", tests.Id)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue), timestamp)
	type Response struct {
		Code int
		Msg  string
		cores.Links
	}
	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)
	beego.Trace("testing", "TestUserPut", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 修改用户信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestUserDelete(t *testing.T) {

	method := "DELETE"
	RequestURL := "/v1/user/" + fmt.Sprintf("%d", tests.Id)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, nil, timestamp)
	type Response struct {
		Code int
		Msg  string
		cores.Links
	}
	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)
	beego.Trace("testing", "TestUserDelete", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 删除用户信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

var userID int64

func TestUserDistributorRole(t *testing.T) {

	role := new(auth.SysRole)
	role.Code = "1001"
	role.Name = "管理员"
	role1 := new(auth.SysRole)
	role1.Code = "1000"
	role1.Name = "普通员工"
	user := new(auth.SysUser)
	user.Account = "123456"
	user.Name = "测试员工"
	user.Password = cores.RandStringByLen(10)
	user.Salt = cores.RandStringByLen(6)
	o := orm.NewOrm()
	roleId, _ := o.Insert(role)
	roleId2, _ := o.Insert(role1)
	userId, _ := o.Insert(user)
	userID = userId
	values := map[string]interface{}{"roleId": []int64{roleId, roleId2}}
	jsonValue, _ := json.Marshal(values)

	method := "POST"
	RequestURL := "/v1/user/" + fmt.Sprint(userId) + "/role"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, bytes.NewBuffer(jsonValue), timestamp)

	type Response struct {
		Code int
		Msg  string
		cores.Links
	}
	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)
	beego.Trace("testing", "TestUserDistributorRole", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 用户分配角色\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

func TestUserGetRole(t *testing.T) {

	method := "GET"
	RequestURL := "/v1/user/" + fmt.Sprint(userID) + "/role"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL, timestamp)
	w := tests.Request(method, RequestURL, signature, nil, timestamp)

	type Response struct {
		Code    int
		Msg     string
		SysRole []auth.SysRole
		cores.Links
	}
	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)
	beego.Trace("testing", "TestUserGetRole", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: 获取角色信息\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}
