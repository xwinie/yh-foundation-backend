package auth

import (
	. "yh-foundation-backend/cores"
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
)

func CreateUserService(sysUser *SysUser) (responseEntity ResponseEntity) {
	IsExistUser, _ := CheckAccountIsExist(sysUser.Account)
	if IsExistUser {
		return *responseEntity.BuildError(BuildEntity(UserIsExist, GetMsg(UserIsExist)))
	}
	sysUser.Salt = RandStringByLen(6)
	sysUser.Password = sysUser.EncryptionPassword(sysUser.Password);
	id, err1 := CreateUser(sysUser)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/user/"+sysUser.Account, "self", "GET", "根据账号获取用户信息"))
	links.Add(LinkTo("/v1/user/"+fmt.Sprint(id), "self", "DELETE", "根据id删除用户信息"))
	links.Add(LinkTo("/v1/user/"+fmt.Sprint(id), "self", "PUT", "根据id修改用户信息"))
	if err1 != nil {
		return *responseEntity.BuildError(BuildEntity(CreateUserError, GetMsg(CreateUserError)))
	} else {
		return *responseEntity.BuildPostAndPut(hateoas.AddLinks(links))
	}
}

func DeleteUserService(id int64) (responseEntity ResponseEntity) {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if r := recover(); r != nil {
			responseEntity.BuildError(BuildEntity(DeleteUserError, GetMsg(DeleteUserError)))
			return
		}
	}()
	m := make(map[string]interface{})
	m["DeleteStatus"] = 1
	num, err := UpdateUser(m, id)
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/user?perPage={perPage}&p={p}", "self", "GET", "根据分页获取用户信息"))
	if num < 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(DeleteUserError, GetMsg(DeleteUserError)))
	} else {
		return *responseEntity.Build(hateoas.AddLinks(links))
	}
}

func UpdateUserService(id int64, sysUser map[string]interface{}) (responseEntity ResponseEntity) {
	if len(sysUser) == 0 {
		return *responseEntity.BuildError(BuildEntity(ParameterError, GetMsg(ParameterError)))
	}
	u, _ := ReadUserById(id)
	if val, ok := sysUser["Password"]; ok {
		sysUser["Password"] = u.EncryptionPassword(val.(string));
	}
	num, err := UpdateUser(sysUser, id)

	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/user/"+u.Account, "self", "GET", "根据账号获取用户信息"))
	links.Add(LinkTo("/v1/user/"+fmt.Sprint(id), "self", "DELETE", "根据id删除用户信息"))
	links.Add(LinkTo("/v1/user/"+fmt.Sprint(id), "self", "PUT", "根据id修改用户信息"))
	if num < 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(UpdateUserError, GetMsg(UpdateUserError)))
	} else {
		return *responseEntity.Build(hateoas.AddLinks(links))
	}
}

func FindUserByAccountService(account string) (responseEntity ResponseEntity) {
	if account == "" {
		return *responseEntity.BuildError(BuildEntity(ParameterError, GetMsg(ParameterError)))
	}
	user, err := ReadUser(account)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/user/"+fmt.Sprint(user.Id), "self", "DELETE", "根据id删除用户信息"))
	links.Add(LinkTo("/v1/user/"+fmt.Sprint(user.Id), "self", "PUT", "根据id修改用户信息"))
	hateoas.AddLinks(links)
	type data struct {
		*QuerySysUser
		*Hateoas
	}
	d := &data{user, &hateoas}
	if len(user.Account) == 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(UserIsExist, GetMsg(UserIsExist)))
	} else {
		return *responseEntity.Build(d)
	}
}

func FindUserByPageService(p *pagination.Paginator) (responseEntity ResponseEntity) {
	users, num, err := ReadUserByPage(p.PerPageNums, p.Offset())
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/user/{account}", "self", "GET", "根据用户账号获取用户信息"))
	links.Add(LinkTo("/v1/user/{id}", "self", "DELETE", "根据id删除用户信息"))
	links.Add(LinkTo("/v1/user/{id}", "self", "PUT", "根据id修改用户信息"))
	links.Add(LinkTo(p.PageLinkFirst(), "first", "GET", ""))
	links.Add(LinkTo(p.PageLinkLast(), "last", "GET", ""))
	if p.HasNext() {
		links.Add(LinkTo(p.PageLinkNext(), "next", "GET", ""))
	}
	if p.HasPrev() {
		links.Add(LinkTo(p.PageLinkPrev(), "prev", "GET", ""))
	}
	hateoas.AddLinks(links)
	type data struct {
		Users []*QuerySysUser
		Total int64
		*HateoasTemplate
	}
	d := &data{users, p.Nums(), &hateoas}
	if num == 0 || err != nil {
		return *responseEntity.BuildError(BuildEntity(QueryError, GetMsg(QueryError)))
	}
	return *responseEntity.Build(d)
}

func FindUserCountByPageService() int64 {
	num, err := ReadUserCountByPage()
	if err != nil {
		return 0
	}
	return num
}
