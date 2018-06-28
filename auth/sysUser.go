package auth

import (
	"time"
	. "yh-foundation-backend/cores"
	"github.com/astaxie/beego/orm"
)

type SysUser struct {
	Id           int64
	Account      string    `orm:"size(100);unique"`
	Name         string    `orm:"size(100);null"`
	UserType     int8      `orm:"default(0)"` //0是第三方用户1是self
	Password     string    `orm:"size(200)"`
	Salt         string    `orm:"size(100)"`
	DeleteStatus int8      `orm:"default(0)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
	Locked       int8      `orm:"default(0)"`
}

type QuerySysUser struct {
	Id           int64
	Account      string
	Name         string
	UserType     int8
	DeleteStatus int8
	Created      time.Time
	Updated      time.Time
	Locked       int8
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysUser))
}

//cores.Md5(cores.Md5(cores.Sha1("12345") + cores.Sha1("passwod")) + salt)
func (u SysUser) CheckEqualPassword(password string) bool {
	return u.Password == Md5(password+u.Salt)
}

func (u SysUser) EncryptionPassword(password string) string {
	return Md5(password + u.Salt)
}

func CreateUser(p *SysUser) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	return id, err
}

func CheckAccountIsExist(account string) (bool, error) {
	o := orm.NewOrm()
	var user SysUser
	err := o.QueryTable("sys_user").Filter("account", account).One(&user)
	if &user.Id != nil && err == nil {
		return true, err
	} else {
		return false, err
	}
}

func DeleteUser(p *SysUser) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(p)
	return num, err

}

func UpdateUser(p map[string]interface{}, id int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_user").Filter("id", id).Update(p)
	return num, err

}

func ReadUser(account string) (user *QuerySysUser, err error) {
	o := orm.NewOrm()
	sqlBuild := NewSqlBuild()
	sqlBuild.Table("sys_user")
	sqlBuild.Select("id,account,name,user_type,locked,delete_status,created,updated")
	sqlBuild.Where("account", account)
	err = o.Raw(sqlBuild.GenerateSelectSql()).QueryRow(&user)
	//err := o.QueryTable("sys_user").Filter("account", account).One(&user)
	return user, err
}

func FindUserAllColums(account string) (user SysUser, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("sys_user").Filter("account", account).One(&user)
	return user, err
}

func ReadUserById(id int64) (*SysUser, error) {
	o := orm.NewOrm()
	var user SysUser
	err := o.QueryTable("sys_user").Filter("id", id).One(&user)
	return &user, err
}

func ReadUserByPage(pageSize int, offset int) (users []*QuerySysUser, num int64, err error) {
	o := orm.NewOrm()
	sqlBuild := NewSqlBuild()
	sqlBuild.Table("sys_user")
	sqlBuild.Select("id,account,name,user_type,locked,delete_status,created,updated")
	sqlBuild.Limit(pageSize, offset)
	num, err = o.Raw(sqlBuild.GenerateSelectSql()).QueryRows(&users)

	//num, err = o.QueryTable("sys_user").Limit(pageSize, offset).All(&users,"Id","Account","Name",
	//	"UserType","DeleteStatus","Created","Updated","Locked")

	return users, num, err
}

func ReadUserCountByPage() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_user").Count()
	return num, err
}
