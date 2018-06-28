package dicts

import (
	. "yh-foundation-backend/cores"
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
)

func CreateDictService(sysDict *SysDict) (responseEntity ResponseEntity) {
	IsExistUser, _ := CheckCodeIsExist(sysDict.Code)
	if IsExistUser {
		return *responseEntity.BuildError(BuildEntity(DictIsExist, GetMsg(DictIsExist)))
	}
	id, err1 := CreateDict(sysDict)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/dict/"+fmt.Sprint(id), "self", "GET", "根据ID获取字典信息"))
	links.Add(LinkTo("/v1/dict/"+fmt.Sprint(id), "self", "DELETE", "根据id删除字典信息"))
	links.Add(LinkTo("/v1/dict/"+fmt.Sprint(id), "self", "PUT", "根据id修改字典信息"))
	if err1 != nil {
		return *responseEntity.BuildError(BuildEntity(CreateDictError, GetMsg(CreateDictError)))
	} else {
		return *responseEntity.BuildPostAndPut(hateoas.AddLinks(links))
	}
}

func DeleteDictService(id int64) (responseEntity ResponseEntity) {

	num, err := DeleteDict(&SysDict{Id: id})
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/dict?perPage={perPage}&p={p}", "self", "GET", "根据分页获取数据字典信息"))
	if num < 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(DeleteDictError, GetMsg(DeleteDictError)))
	} else {
		return *responseEntity.Build(hateoas.AddLinks(links))
	}
}

func UpdateDictService(id int64, sysDict map[string]interface{}) (responseEntity ResponseEntity) {
	num, err := UpdateDict(sysDict, id)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/dict/"+fmt.Sprint(id), "self", "GET", "根据ID获取字典信息"))
	links.Add(LinkTo("/v1/dict/"+fmt.Sprint(id), "self", "DELETE", "根据id删除字典信息"))
	links.Add(LinkTo("/v1/dict/"+fmt.Sprint(id), "self", "PUT", "根据id修改字典信息"))
	if num < 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(UpdateDictError, GetMsg(UpdateDictError)))
	} else {
		return *responseEntity.Build(hateoas.AddLinks(links))
	}
}

func ReadDictByIdService(id int64) (responseEntity ResponseEntity) {
	dict, err := ReadDictById(id)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/dict/"+fmt.Sprint(id), "self", "GET", "根据ID获取字典信息"))
	links.Add(LinkTo("/v1/dict/"+fmt.Sprint(id), "self", "DELETE", "根据id删除字典信息"))
	links.Add(LinkTo("/v1/dict/"+fmt.Sprint(id), "self", "PUT", "根据id修改字典信息"))

	type data struct {
		*SysDict
		*Hateoas
	}
	d := &data{dict, &hateoas}
	if &dict.Id == nil && err != nil {
		return *responseEntity.BuildError(BuildEntity(NotFoundDict, GetMsg(NotFoundDict)))
	} else {
		return *responseEntity.Build(d)
	}
}

func FindDictByPageService(p *pagination.Paginator, pg PageQueryDict) (responseEntity ResponseEntity) {

	dicts, num, err := ReadDictByPage(pg, p.PerPageNums, p.Offset())
	if num == 0 || err != nil {
		return *responseEntity.BuildError(BuildEntity(QueryError, GetMsg(QueryError)))
	}

	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/dict/{id}", "self", "GET", "根据ID获取字典信息"))
	links.Add(LinkTo("/v1/dict/{id}", "self", "DELETE", "根据id删除字典信息"))
	links.Add(LinkTo("/v1/dict/{id}", "self", "PUT", "根据id修改字典信息"))
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
		Dicts []*SysDict
		*HateoasTemplate
	}
	d := &data{dicts, &hateoas}
	return *responseEntity.Build(d)
}

func FindDictCountByPageService(pg PageQueryDict) int64 {
	num, err := ReadDictCountByPage(pg)
	if err != nil {
		return 0
	}
	return num
}
