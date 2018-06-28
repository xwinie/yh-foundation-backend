package routers

import (
	"yh-foundation-backend/auth"
	"yh-foundation-backend/dicts"
	"yh-foundation-backend/files"

	"github.com/astaxie/beego"
)

func AuthRouter() {

	ns :=
		beego.NewNamespace("/v1",
			beego.NSRouter("/login", &auth.LoginController{}, "post:Login"),
			beego.NSNamespace("/user",
				beego.NSRouter("/", &auth.UserController{}, "post:CreateUserCtl"),
				beego.NSRouter("/:id", &auth.UserController{}, "delete:DeleteUserCtl"),
				beego.NSRouter("/:id", &auth.UserController{}, "put:UpdateUserCtl"),
				beego.NSRouter("/:account", &auth.UserController{}, "get:FindUserByAccountCtl"),
				beego.NSRouter("/:id/role", &auth.UserRoleController{}, "post:UserDistributorRoleCtl"),
				beego.NSRouter("/:id/role", &auth.UserRoleController{}, "get:RoleByUserIdCtl"),
				beego.NSRouter("/", &auth.UserController{}, "get:FindUserByPageCtl"),
			),
			beego.NSNamespace("/role",
				beego.NSRouter("/", &auth.RoleController{}, "post:CreateRoleCtl"),
				beego.NSRouter("/:id", &auth.RoleController{}, "delete:DeleteRoleCtl"),
				beego.NSRouter("/:id", &auth.RoleController{}, "put:UpdateRoleCtl"),
				beego.NSRouter("/:code", &auth.RoleController{}, "get:FindRoleByCodeCtl"),
				beego.NSRouter("/:id/resource", &auth.RoleResourceController{}, "post:RoleDistributorResourceCtl"),
				beego.NSRouter("/:id/resource", &auth.RoleResourceController{}, "get:ResourceByRoleIdCtl"),
				beego.NSRouter("/", &auth.RoleController{}, "get:FindRoleByPageCtl"),
			),
			beego.NSNamespace("/resource",
				beego.NSRouter("/", &auth.ResourceController{}, "post:CreateResourceCtl"),
				beego.NSRouter("/:id", &auth.ResourceController{}, "delete:DeleteResourceCtl"),
				beego.NSRouter("/:id", &auth.ResourceController{}, "put:UpdateResourceCtl"),
				beego.NSRouter("/:code", &auth.ResourceController{}, "get:FindResourceByCodeCtl"),
				beego.NSRouter("/", &auth.ResourceController{}, "get:FindResourceByPageCtl"),
				beego.NSRouter("/:userId", &auth.ResourceController{}, "get:MenusByUserId"),
			),

			beego.NSNamespace("/menus",
				beego.NSRouter("/:userId", &auth.ResourceController{}, "get:MenusByUserId"),
			),
			beego.NSNamespace("/client",
				beego.NSRouter("/", &auth.ClientController{}, "post:CreateClientCtl"),
				beego.NSRouter("/", &auth.ClientController{}, "get:FindClientByPageCtl"),
				beego.NSRouter("/:id", &auth.ClientController{}, "delete:DeleteClientCtl"),
				beego.NSRouter("/:id", &auth.ClientController{}, "put:UpdateClientCtl"),
				beego.NSRouter("/:clientId", &auth.ClientController{}, "get:FindClientByClientIdCtl"),
			),
			beego.NSNamespace("/app",
				beego.NSRouter("/", &auth.ClientController{}, "get:FindClientDefault"),
			),
			beego.NSNamespace("/dict",
				beego.NSRouter("/", &dicts.DictController{}, "post:CreateDictCtl"),
				beego.NSRouter("/:id", &dicts.DictController{}, "delete:DeleteDictCtl"),
				beego.NSRouter("/:id", &dicts.DictController{}, "put:UpdateDictCtl"),
				beego.NSRouter("/:id", &dicts.DictController{}, "get:FindDictByIdCtl"),
				beego.NSRouter("/", &dicts.DictController{}, "get:FindDictByPageCtl"),
			),
			beego.NSNamespace("/upload",
				beego.NSRouter("/", &files.FileUploadController{}, "post:UploadCtl"),
			),
		)

	beego.AddNamespace(ns)

	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/images", "images")

}
