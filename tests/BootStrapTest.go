package tests

import (
	"path/filepath"
	"runtime"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	"yh-foundation-backend/cores"
	"yh-foundation-backend/routers"
	"time"
	"yh-foundation-backend/filter"
	"yh-foundation-backend/auth"
	"net/http"
	"net/http/httptest"
	"io"
	"github.com/astaxie/beego/config"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	appPath, _ := filepath.Abs(filepath.Base(filepath.Join(file, ".."+string(filepath.Separator))))
	dir, _ := filepath.Split(appPath)
	appConf, err := config.NewConfig("ini", dir+"conf/app.conf")
	if err != nil {
		panic(err)
	}

	routers.AuthRouter()
	beego.ErrorController(&cores.ErrorController{})

	beego.InsertFilter("*", beego.BeforeRouter, filter.APISecretAuth(3600))
	beego.InsertFilter("*", beego.BeforeRouter, filter.Authorizer())
	beego.TestBeegoInit(dir)

	orm.DefaultTimeLoc = time.Local
	orm.RegisterDataBase("default", "sqlite3", "file::memory:?mode=memory&cache=shared&loc=Local&parseTime=true")
	name := "default"
	force := false
	verbose := true
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error("init db table error:", err)
	}
	orm.RunCommand()

	InitAuthData()
	cores.InitCache(appConf)
	//orm.Debug=true

}

var Id int64

func Request(method, RequestURL, signature string, body io.Reader, timestamp string) (*httptest.ResponseRecorder) {
	tokenString := GetToken()
	if tokenString == "" {
		beego.Trace("tokenString is empty")
		return nil
	}
	r, _ := http.NewRequest(method, RequestURL, body)
	r.Header.Set("appid", "app1")
	r.Header.Set("timestamp", timestamp)
	r.Header.Set("signature", signature)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)
	return w
}

func RequestByToken(method, RequestURL, signature string, body io.Reader, tokenString string, timestamp string) (*httptest.ResponseRecorder) {
	if tokenString == "" {
		beego.Trace("tokenString is empty")
		return nil
	}
	r, _ := http.NewRequest(method, RequestURL, body)
	r.Header.Set("appid", "app1")
	r.Header.Set("timestamp", timestamp)
	r.Header.Set("signature", signature)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)
	return w
}

func InitAuthData() {

	Secret := cores.RandStringByLen(10)
	client := new(auth.SysClient)
	client.ClientId = "app1"
	client.Name = "测试app"
	client.Secret = Secret
	client.VerifySecret = Secret
	o := orm.NewOrm()
	clientId, _ := o.Insert(client)
	if clientId < 0 {
		beego.Trace("添加数据失败")
	}

	role := new(auth.SysRole)
	salt := cores.RandStringByLen(6)
	role.Code = "1"
	role.Name = "管理员"
	user := new(auth.SysUser)
	user.Account = "12345"
	user.Name = "测试员工"
	user.Password = cores.Md5(cores.Md5(cores.Sha1("12345")+cores.Sha1("Password")) + salt)
	user.Salt = salt
	roleId, _ := o.Insert(role)
	userId, _ := o.Insert(user)
	userRole := new(auth.SysUserRole)
	userRole.RoleId = roleId
	userRole.UserId = userId
	userRoleId, _ := o.Insert(userRole)
	if userRoleId == 0 {

	}
	resource := []auth.SysResource{
		{Code: "10001", Action: "/v1/user", Method: "POST", Name: "添加用户"},
		{Code: "10002", Action: "/v1/user/:id", Method: "DELETE", Name: "删除用户"},
		{Code: "10003", Action: "/v1/user/:id", Method: "PUT", Name: "修改用户"},
		{Code: "10004", Action: "/v1/user/:account", Method: "GET", Name: "根据账号获取用户"},
		{Code: "10005", Action: "/v1/user/:id/role", Method: "POST", Name: "给用户分配角色"},
		{Code: "10006", Action: "/v1/role", Method: "POST", Name: "添加角色"},
		{Code: "10007", Action: "/v1/role/:id", Method: "DELETE", Name: "删除角色"},
		{Code: "10008", Action: "/v1/role/:id", Method: "PUT", Name: "修改角色"},
		{Code: "10009", Action: "/v1/role/:code", Method: "GET", Name: "根据编码获取角色"},
		{Code: "10010", Action: "/v1/role/:id/resource", Method: "POST", Name: "给角色分配资源"},
		{Code: "10011", Action: "/v1/resource", Method: "POST", Name: "添加资源"},
		{Code: "10012", Action: "/v1/resource/:id", Method: "DELETE", Name: "删除资源"},
		{Code: "10013", Action: "/v1/resource/:id", Method: "PUT", Name: "修改资源"},
		{Code: "10014", Action: "/v1/resource/:code", Method: "GET", Name: "根据编码获取资源"},
		{Code: "10015", Action: "/v1/client", Method: "POST", Name: "添加应用"},
		{Code: "10016", Action: "/v1/client/:id", Method: "DELETE", Name: "删除应用"},
		{Code: "10017", Action: "/v1/client/:id", Method: "PUT", Name: "修改应用"},
		{Code: "10018", Action: "/v1/client/:clientId", Method: "GET", Name: "根据应用id获取应用"},
		{Code: "10019", Action: "/v1/user/:id", Method: "PUT", Name: "修改用户"},
		{Code: "10020", Action: "/v1/role/:id/resource", Method: "GET", Name: "根据角色ID获取资源信息"},
		{Code: "10021", Action: "/v1/user/:id/role", Method: "GET", Name: "根据ID获取角色信息"},
		{Code: "10022", Action: "/v1/resource", Method: "GET", Name: "分页获取所有资源"},
		{Code: "10023", Action: "/v1/user", Method: "GET", Name: "分页获取所有用户"},
		{Code: "10024", Action: "/v1/role", Method: "GET", Name: "分页获取所有角色"},
		{Code: "10025", Action: "/v1/dict", Method: "POST", Name: "添加数据字典"},
		{Code: "10026", Action: "/v1/dict/:id", Method: "DELETE", Name: "删除数据字典"},
		{Code: "10027", Action: "/v1/dict/:id", Method: "PUT", Name: "修改数据字典"},
		{Code: "10028", Action: "/v1/dict", Method: "GET", Name: "根据分页获取数据字典"},
		{Code: "10029", Action: "/v1/dict/:id", Method: "GET", Name: "根据ID获取数据字典"},
		{Code: "10030", Action: "/v1/upload", Method: "POST", Name: "上传文件"},
		{Code: "10031", Action: "/static", Method: "GET", Name: "获取文件"},
		{Code: "10032", Action: "/images", Method: "GET", Name: "获取图片"},
		{Code: "10033", Action: "/v1/menus/:userId", Method: "GET", Name: "根据用户获取菜单信息"},
	}
	o.InsertMulti(len(resource), resource)
	var resources []int64
	res1, _ := o.Raw("SELECT id resourceId "+
		"  from sys_resource where code in "+
		"('10000','10001','10002','10003','10004','10005','10006','10007','10008','10009','10010','10011',"+
		"'10012','10013','10014','10015','10016','10017','10018','10019','10020','10021','10022','10023',"+
		"'10024','10025','10026','10027','10028','10029','10030','10031','10032','10033')", roleId).QueryRows(&resources)

	beego.Trace("mysql row affected num:%d,%v ", res1, resources)
	roleResources := auth.RoleDistributorResource{ResourceId: resources}
	response := auth.RoleDistributorResourceService(roleId, &roleResources)
	beego.Trace("init RoleResource Success ", response)

}
