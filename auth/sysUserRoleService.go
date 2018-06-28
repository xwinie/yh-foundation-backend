package auth

import (
	. "yh-foundation-backend/cores"
	"fmt"
)

func UserDistributorRoleService(userId int64, p *UserDistributorRole) (responseEntity ResponseEntity) {

	response := DeleteUserRoleService(userId)
	if response.Code == 100018 {
		return *responseEntity.BuildError(BuildEntity(UserDistributorRoleError, GetMsg(UserDistributorRoleError)))
	}
	sysUserRole := []SysUserRole{}
	userRole := SysUserRole{}
	for _, value := range p.RoleId {
		userRole.UserId = userId
		userRole.RoleId = value
		sysUserRole = append(sysUserRole, userRole)
	}
	_, err := CreateMultiUserRole(sysUserRole)

	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/user/"+fmt.Sprint(userId)+"/role", "self", "GET", "根据ID获取角色信息"))

	if err != nil {
		return *responseEntity.BuildError(BuildEntity(UserDistributorRoleError, GetMsg(UserDistributorRoleError)))
	}

	return *responseEntity.BuildPostAndPut(hateoas.AddLinks(links))

}

func DeleteUserRoleService(id int64) Entity {
	num, err := DeleteUserRole(id)
	if num < 0 && err != nil {
		return *BuildEntity(DeleteUserRoleError, GetMsg(DeleteUserRoleError))
	} else {
		return *BuildEntity(Success, GetMsg(Success))
	}
}

func RoleByUserIdService(userId int64) (responseEntity ResponseEntity) {
	roles, _, err := FindRoleByUserId(userId)
	if err != nil {
		return *responseEntity.BuildError(BuildEntity(QueryError, GetMsg(QueryError)+err.Error()))
	}
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/role/{code}", "related", "GET", "根据编码获取信息"))
	hateoas.AddLinks(links)
	type data struct {
		Roles interface{}
		HateoasTemplate
	}
	d := &data{roles, hateoas}

	return *responseEntity.Build(d)

}
