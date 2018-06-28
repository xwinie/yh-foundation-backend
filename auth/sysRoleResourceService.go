package auth

import (
	. "yh-foundation-backend/cores"
	"fmt"
	"github.com/astaxie/beego"
)

func RoleDistributorResourceService(roleId int64, p *RoleDistributorResource) (responseEntity ResponseEntity) {
	response := DeleteRoleResourceService(roleId)
	if response.Code == 100016 {
		return *responseEntity.BuildError(BuildEntity(RoleDistributorResourceError, GetMsg(RoleDistributorResourceError)))

	}

	var roleResources []SysRoleResource
	var roleResource SysRoleResource
	for _, value := range p.ResourceId {
		roleResource.RoleId = roleId
		if value != 0 {
			roleResource.ResourceId = value
			roleResources = append(roleResources, roleResource)
		}
	}
	_, err := CreateMultiRoleResource(roleResources)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/role/"+fmt.Sprint(roleId)+"/resource", "self", "GET", "根据编码获取角色信息"))
	if err != nil {
		return *responseEntity.BuildError(BuildEntity(RoleDistributorResourceError, GetMsg(RoleDistributorResourceError)))
	}
	return *responseEntity.BuildPostAndPut(hateoas.AddLinks(links))

}

func DeleteRoleResourceService(id int64) Entity {
	num, err := DeleteRoleResource(id)
	beego.Trace(111, err)
	if num < 0 && err != nil {
		return *BuildEntity(DeleteRoleResourceError, GetMsg(DeleteRoleResourceError))
	} else {
		return *BuildEntity(Success, GetMsg(Success))
	}
}

func ResourceByRoleIdService(roleId int64) (responseEntity ResponseEntity) {
	resources, _, err := FindResourceByRoleId(roleId)
	if err != nil {
		return *responseEntity.BuildError(BuildEntity(QueryError, GetMsg(QueryError)+err.Error()))
	}
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/resource/{code}", "related", "GET", "根据编码获取资源信息"))
	hateoas.AddLinks(links)
	type data struct {
		Resources interface{}
		HateoasTemplate
	}
	d := &data{resources, hateoas}

	return *responseEntity.Build(d)

}
