package auth

import (
	. "yh-foundation-backend/cores"
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
)

func CreateResourceService(p *SysResource) (responseEntity ResponseEntity) {
	IsExistCode, err := CheckResourceCodeIsExist(p.Code)
	if IsExistCode {
		return *responseEntity.BuildError(BuildEntity(ResourceIsExist, GetMsg(ResourceIsExist)))
	}
	id, err := CreateResource(p)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/resource/"+p.Code, "self", "GET", "根据编码获取资源信息"))
	links.Add(LinkTo("/v1/resource/"+fmt.Sprint(id), "self", "DELETE", "根据id删除资源信息"))
	links.Add(LinkTo("/v1/resource/"+fmt.Sprint(id), "self", "PUT", "根据id修改资源信息"))
	if err != nil {
		return *responseEntity.BuildError(BuildEntity(CreateResourceError, GetMsg(CreateResourceError)))
	} else {
		return *responseEntity.BuildPostAndPut(hateoas.AddLinks(links))
	}
}

func DeleteResourceService(id int64) (responseEntity ResponseEntity) {

	num, err := DeleteResource(&SysResource{Id: id})
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/resource?perPage={perPage}&p={p}", "self", "GET", "根据分页获取资源信息"))
	if num < 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(DeleteResourceError, GetMsg(DeleteResourceError)))
	} else {
		return *responseEntity.Build(hateoas.AddLinks(links))
	}
}

func UpdateResourceService(id int64, p map[string]interface{}) (responseEntity ResponseEntity) {
	num, err := UpdateResource(p, id)
	resource, err := ReadResourceById(id)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/resource/"+resource.Code, "self", "GET", "根据编码获取资源信息"))
	links.Add(LinkTo("/v1/resource/"+fmt.Sprint(id), "self", "DELETE", "根据id删除资源信息"))
	links.Add(LinkTo("/v1/resource/"+fmt.Sprint(id), "self", "PUT", "根据id修改资源信息"))
	if num < 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(UpdateResourceError, GetMsg(UpdateResourceError)))
	} else {
		return *responseEntity.Build(hateoas.AddLinks(links))
	}
}

func FindResourceByCodeService(code string) (responseEntity ResponseEntity) {
	if code == "" {
		return *responseEntity.BuildError(BuildEntity(ParameterError, GetMsg(ParameterError)))
	}
	resource, err := ReadResourceByCode(code)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/resource/"+fmt.Sprint(resource.Id), "self", "DELETE", "根据id删除资源信息"))
	links.Add(LinkTo("/v1/resource/"+fmt.Sprint(resource.Id), "self", "PUT", "根据id修改资源信息"))
	hateoas.AddLinks(links)
	type data struct {
		*SysResource
		*Hateoas
	}
	d := &data{resource, &hateoas}
	if len(resource.Code) == 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(ResourceIsExist, GetMsg(ResourceIsExist)))
	} else {
		return *responseEntity.Build(d)
	}
}

func FindResourceByPageService(p *pagination.Paginator) (responseEntity ResponseEntity) {
	resources, num, err := ReadResourceByPage(p.PerPageNums, p.Offset())
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/resource/{code}", "self", "GET", "根据编码获取资源信息"))
	links.Add(LinkTo("/v1/resource/{id}", "self", "DELETE", "根据id删除资源信息"))
	links.Add(LinkTo("/v1/resource/{id}", "self", "PUT", "根据id修改资源信息"))
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
		Resources []*SysResource
		*HateoasTemplate
	}
	d := &data{resources, &hateoas}
	if num == 0 || err != nil {
		return *responseEntity.BuildError(BuildEntity(QueryError, GetMsg(QueryError)))
	}
	return *responseEntity.Build(d)
}

func FindResourceCountByPageService() int64 {
	num, err := ReadResourceCountByPage()
	if err != nil {
		return 0
	}
	return num
}

func MenuByUserIdService(userId int64) (responseEntity ResponseEntity) {
	//roles := GlobalCaches.Get(fmt.Sprint(userId) + "userType")
	//beego.Trace(111, roles)
	roleId, num, _ := FindRoleIdByUserId(userId)
	menus, num, err := FindResourceByMultiRole(roleId, 1)
	if num == 0 || err != nil {
		return *responseEntity.BuildError(BuildEntity(QueryError, GetMsg(QueryError)))
	}
	return *responseEntity.Build(menus)
}
