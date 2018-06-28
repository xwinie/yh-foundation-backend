package auth

import (
	"fmt"
	"time"
	. "yh-foundation-backend/cores"

	"github.com/astaxie/beego/utils/pagination"
)

func GetSecretService(clientId string) string {
	appSecret, _ := GetSecret(clientId)
	return appSecret
}

func GetClientService(clientId string) (SysClient, error) {
	return GetClient(clientId)
}

func CreateClientService(p *SysClient) (responseEntity ResponseEntity) {

	p.ClientId = RandStringByLen(30)
	IsExistClient, err := CheckClientIsExist(p.ClientId)
	if !IsExistClient && err != nil {
		return *responseEntity.BuildError(BuildEntity(ClientIsExist, GetMsg(ClientIsExist)))
	}
	p.Secret = RandStringByLen(20)
	p.VerifySecret = RandStringByLen(20)
	id, err := CreateClient(p)

	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/client/"+p.ClientId, "self", "GET", "根据应用ID获取应用信息"))
	links.Add(LinkTo("/v1/client/"+fmt.Sprint(id), "self", "DELETE", "根据id删除应用信息"))
	links.Add(LinkTo("/v1/client/"+fmt.Sprint(id), "self", "PUT", "根据id修改应用信息"))
	if err != nil {
		return *responseEntity.BuildError(BuildEntity(CreateClientError, GetMsg(CreateClientError)))
	} else {
		return *responseEntity.BuildPostAndPut(hateoas.AddLinks(links))
	}
}

func DeleteClientService(id int64) (responseEntity ResponseEntity) {

	num, err := DeleteClient(&SysClient{Id: id})

	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/client?perPage={perPage}&p={p}", "self", "GET", "根据分页获取应用信息"))
	hateoas.AddLinks(links)
	if num < 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(DeleteClientError, GetMsg(DeleteClientError)))
	} else {
		return *responseEntity.Build(hateoas.AddLinks(links))
	}
}

func UpdateClientService(id int64, client map[string]interface{}) (responseEntity ResponseEntity) {
	num, err := UpdateClient(client, id)
	var hateoas Hateoas
	var links Links
	links.Add(LinkTo("/v1/client/"+fmt.Sprint(id), "self", "DELETE", "根据id删除应用信息"))
	links.Add(LinkTo("/v1/client/"+fmt.Sprint(id), "self", "PUT", "根据id修改应用信息"))
	if num < 0 && err != nil {
		return *responseEntity.BuildError(BuildEntity(UpdateClientError, GetMsg(UpdateClientError)))
	} else {
		return *responseEntity.Build(hateoas.AddLinks(links))
	}
}

func FindClientByClientIdService(clientId string) (responseEntity ResponseEntity) {
	client, err := ReadClient(clientId)
	if &client.Id == nil && err != nil {
		return *responseEntity.BuildError(BuildEntity(ClientIsExist, GetMsg(ClientIsExist)))
	} else {
		return *responseEntity.Build(client)
	}
}

func FindClientByPageService(p *pagination.Paginator) (responseEntity ResponseEntity) {
	clients, num, err := ReadClientByPage(p.PerPageNums, p.Offset())
	var hateoas HateoasTemplate
	var links Links
	links.Add(LinkTo("/v1/client/{clientid}", "self", "GET", "根据编码获取角色信息"))
	links.Add(LinkTo("/v1/client/{id}", "self", "DELETE", "根据id删除角色信息"))
	links.Add(LinkTo("/v1/client/{id}", "self", "PUT", "根据id修改角色信息"))
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
		Clients []*SysClient
		Total   int64
		*HateoasTemplate
	}
	d := &data{clients, p.Nums(), &hateoas}
	if num == 0 || err != nil {
		return *responseEntity.BuildError(BuildEntity(QueryError, GetMsg(QueryError)))
	}
	return *responseEntity.Build(d)
}

func FindClientCountByPageService() int64 {
	num, err := ReadClientCountByPage()
	if err != nil {
		return 0
	}
	return num
}

func SignAuthService(data *ClientAuthData) (responseEntity ResponseEntity) {
	appSecret, _ := GetSecret(data.ClientID)
	if appSecret != data.Secret {
		return *responseEntity.BuildError(BuildEntity(ClientIsExist, GetMsg(ClientIsExist)))
	}
	secret := Md5(fmt.Sprintf("%s%s%s", RandStringByLen(20), data.Secret, data.ClientID))
	timeoutDuration := 1 * time.Hour
	GlobalCaches.Put(data.ClientID, secret, timeoutDuration)
	return *responseEntity.BuildPostAndPut(map[string]string{"secret": secret})
}
