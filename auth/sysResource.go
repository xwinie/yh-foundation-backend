package auth

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SysResource struct {
	Id           int64
	Code         string    `orm:"size(100);unique"`
	Name         string    `orm:"size(100)"`
	Action       string    `orm:"size(200)"`
	Method       string    `orm:"size(100)"`
	IsOpen       int8      `orm:"default(0)"` //0非开放1开放
	ResType      int8      `orm:"default(0)"` //0代表是接口1代表菜单
	ParentId     int64     `orm:"default(0)"`
	DeleteStatus int8      `orm:"default(0)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysResource))
}

func CreateResource(p *SysResource) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	return id, err
}

func CheckResourceCodeIsExist(code string) (bool, error) {
	o := orm.NewOrm()
	var resource SysResource
	err := o.QueryTable("sys_resource").Filter("code", code).One(&resource)
	if &resource.Id != nil && err == nil {
		return true, err
	} else {
		return false, err
	}
}

func DeleteResource(p *SysResource) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(p)
	return num, err

}

func UpdateResource(p map[string]interface{}, id int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_resource").Filter("id", id).Update(p)
	return num, err

}

func ReadResourceByCode(code string) (*SysResource, error) {
	o := orm.NewOrm()
	var resource SysResource
	err := o.QueryTable("sys_resource").Filter("code", code).One(&resource)
	return &resource, err
}

func ReadResourceById(id int64) (*SysResource, error) {
	o := orm.NewOrm()
	var resource SysResource
	err := o.QueryTable("sys_resource").Filter("id", id).One(&resource)
	return &resource, err
}
func ReadAllIsOpenResource() (resources []*SysResource, num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_resource").Filter("is_open", 1).All(&resources)
	return resources, num, err
}

func ReadResourceByPage(pageSize int, offset int) (resources []*SysResource, num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_resource").Limit(pageSize, offset).All(&resources)
	return resources, num, err
}

func ReadResourceCountByPage() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_resource").Count()
	return num, err
}
