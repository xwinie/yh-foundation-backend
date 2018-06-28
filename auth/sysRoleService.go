package auth

import (
	. "yh-foundation-backend/cores"
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
)

func CreateRoleService(p *SysRole) (responseEntity ResponseEntity) {
	IsExistCode, err := CheckRoleCodeIsExist(p.Code)
	if IsExistCode {
		return *responseEntity.BuildError(BuildEntity(RoleIsExist, GetMsg(RoleIsExist)))
	}
	id, err := CreateRole(p)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/role/"+p.Code, "self", "GET", "根据编码获取角色信息"))
	links.Add(LinkTo("/v1/role/"+fmt.Sprint(id), "self", "DELETE", "根据id删除角色信息"))
	links.Add(LinkTo("/v1/role/"+fmt.Sprint(id), "self", "PUT", "根据id修改用户信息"))
	if err != nil {
		return *responseEntity.BuildError(BuildEntity(CreateRoleError, GetMsg(CreateRoleError)))
	} else {
		return *responseEntity.BuildPostAndPut(hateoas.AddLinks(links))
	}
}

func DeleteRoleService(id int64) (responseEntity ResponseEntity) {
	m := make(map[string]interface{})
	m["DeleteStatus"] = 1
	response := DeleteRoleResourceService(id)
	if response.Code == 100016 {
		return *responseEntity.BuildError(BuildEntity(DeleteRoleError, GetMsg(DeleteRoleError)))

	}
	num, err := UpdateRole(m, id)
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/role?perPage={perPage}&p={p}", "self", "GET", "根据分页获取角色信息"))
	if num < 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(DeleteRoleError, GetMsg(DeleteRoleError)))
	} else {
		return *responseEntity.Build(hateoas.AddLinks(links))
	}
}

func UpdateRoleService(id int64, p map[string]interface{}) (responseEntity ResponseEntity) {

	if _, ok := p["code"]; ok {
		return *responseEntity.BuildError(BuildEntity(RoleIsExist, GetMsg(RoleIsExist)))
	}
	num, err := UpdateRole(p, id)
	role, err := ReadRoleById(id)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/role/"+role.Code, "self", "GET", "根据编码获取角色信息"))
	links.Add(LinkTo("/v1/role/"+fmt.Sprint(id), "self", "DELETE", "根据id删除角色信息"))
	links.Add(LinkTo("/v1/role/"+fmt.Sprint(id), "self", "PUT", "根据id修改用户信息"))
	if num < 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(UpdateRoleError, GetMsg(UpdateRoleError)))
	} else {
		return *responseEntity.Build(hateoas.AddLinks(links))
	}
}

func FindRoleByCodeService(code string) (responseEntity ResponseEntity) {
	if code == "" {
		return *responseEntity.BuildError(BuildEntity(ParameterError, GetMsg(ParameterError)))
	}
	role, err := ReadRoleByCode(code)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/role/"+fmt.Sprint(role.Id), "self", "DELETE", "根据id删除资源信息"))
	links.Add(LinkTo("/v1/role/"+fmt.Sprint(role.Id), "self", "PUT", "根据id修改资源信息"))
	hateoas.AddLinks(links)
	type data struct {
		*SysRole
		*Hateoas
	}
	d := &data{role, &hateoas}
	if len(role.Code) == 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(RoleIsExist, GetMsg(RoleIsExist)))
	} else {
		return *responseEntity.Build(d)
	}
}

func FindRoleByIdService(id int64) (responseEntity ResponseEntity) {
	role, err := ReadRoleById(id)
	if &role.Id == nil && err != nil {
		return *responseEntity.BuildError(BuildEntity(RoleIsExist, GetMsg(RoleIsExist)))
	} else {
		return *responseEntity.Build(role)
	}
}

func FindRoleByPageService(p *pagination.Paginator) (responseEntity ResponseEntity) {
	roles, num, err := ReadRoleByPage(p.PerPageNums, p.Offset())
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/role/{code}", "self", "GET", "根据编码获取角色信息"))
	links.Add(LinkTo("/v1/role/{id}", "self", "DELETE", "根据id删除角色信息"))
	links.Add(LinkTo("/v1/role/{id}", "self", "PUT", "根据id修改角色信息"))
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
		Roles []*SysRole
		Total int64
		*HateoasTemplate
	}
	d := &data{roles, p.Nums(), &hateoas}
	if num == 0 || err != nil {
		return *responseEntity.BuildError(BuildEntity(QueryError, GetMsg(QueryError)))
	}
	return *responseEntity.Build(d)
}

func FindRoleCountByPageService() int64 {
	num, err := ReadRoleCountByPage()
	if err != nil {
		return 0
	}
	return num
}
