package auth

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SysClient struct {
	Id           int64
	ClientId     string `orm:"size(100);unique"`
	Name         string `orm:"size(100)"`
	Secret       string `orm:"size(200)"`
	VerifySecret string `orm:"size(200)"`
	Locked       int8   `orm:"default(0)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysClient))
}

func GetSecret(clientId string) (string, error) {
	o := orm.NewOrm()
	var client SysClient
	err := o.QueryTable("sys_client").Filter("client_id", clientId).One(&client)
	if &client.Id == nil && err != nil {
		return "", err
	} else {
		return client.Secret, nil
	}
}

func GetClient(clientId string) (SysClient, error) {
	o := orm.NewOrm()
	var client SysClient
	err := o.QueryTable("sys_client").Filter("client_id", clientId).One(&client)
	if &client.Id == nil && err != nil {
		return client, err
	} else {
		return client, nil
	}
}
func CreateClient(p *SysClient) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	return id, err
}

func CheckClientIsExist(clientId string) (bool, error) {
	o := orm.NewOrm()
	var client SysClient
	err := o.QueryTable("sys_client").Filter("client_id", clientId).One(&client)
	if &client.Id == nil && err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func DeleteClient(p *SysClient) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(p)
	return num, err

}

func UpdateClient(client map[string]interface{}, id int64) (num int64, err error) {
	o := orm.NewOrm()
	beego.Debug("Update Client Data:%s", client)
	num, err = o.QueryTable("sys_client").Filter("id", id).Update(client)
	return num, err

}

func ReadClient(clientId string) (*SysClient, error) {
	o := orm.NewOrm()
	var client SysClient
	err := o.QueryTable("sys_client").Filter("client_id", clientId).One(&client, "Id", "ClientId", "Name", "Secret", "Locked")
	return &client, err
}

func ReadAllClient() (clients []*SysClient, num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_client").All(&clients)
	return clients, num, err
}

func ReadClientCountByPage() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_client").Count()
	return num, err
}

func ReadClientByPage(pageSize int, offset int) (clients []*SysClient, num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_client").Limit(pageSize, offset).All(&clients, "Id", "ClientId", "Name", "Secret", "Locked")
	return clients, num, err
}
