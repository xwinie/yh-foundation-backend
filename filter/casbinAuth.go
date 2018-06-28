package filter

import (
	"net/http"
	"strings"
	"yh-foundation-backend/auth"

	"fmt"
	"yh-foundation-backend/cores"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"time"
	"github.com/garyburd/redigo/redis"
)

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

type Claims struct {
	Exp      float64
	Iat      int64
	Issuer   string
	UserId   int64
	userType int8
	Account  string
	UserName string
	Role     []int64
}

var errResponse cores.Entity

var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
`
var e = casbin.NewEnforcer(casbin.NewModel(text))
var a = &BasicAuthorizer{enforcer: e}
func Authorizer() beego.FilterFunc {
	return func(ctx *context.Context) {

		appId := ctx.Input.Header("appid")

		if ctx.Input.URL() == "/v1/login" || HasIsOpenRestPermission(ctx, appId) {
			return
		}
		VerifySecret, redisErr := redis.String(cores.GlobalCaches.Get(appId), nil)
		if redisErr != nil {
			client, err := auth.GetClientService(appId)
			if err != nil {
				ctx.Output.Status = http.StatusForbidden
				ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("sign is either expired or not active yet"), false, false)
				return
			}
			VerifySecret = client.VerifySecret
			ctx.Output.Status = http.StatusForbidden
			ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("sign is either expired or not active yet"), false, false)
			return
		}
		token := GetToken(ctx, VerifySecret)
		claimsMap, ok := token.Claims.(jwt.MapClaims)
		var claims Claims
		err := mapstructure.Decode(claimsMap, &claims)
		if !ok || err != nil {
			ctx.Output.Status = http.StatusForbidden
			ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("Token is Error"), false, false)
			return
		}

		cores.GlobalCaches.Put(fmt.Sprint(claims.UserId)+"userType", claims.userType, 30*time.Second)
		cores.GlobalCaches.Put(fmt.Sprint(claims.UserId)+"roles", claims.Role, 30*time.Second)
		HasRestPermission(ctx, &claims)

	}
}

func HasIsOpenRestPermission(ctx *context.Context, appId string) bool {
	resources, _, err := auth.ReadAllIsOpenResource()
	if err != nil && resources == nil {
		ctx.Output.Status = http.StatusForbidden
		ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("Token is Error"), false, false)
		return false
	}
	for _, v := range resources {
		e.AddPermissionForUser(appId, v.Action, v.Method)
	}

	if !a.CheckPermission(appId, ctx.Request) {
		return false
	}
	return true
}

func HasRestPermission(ctx *context.Context, claims *Claims) {

	resources, _, err := auth.FindResourceByMultiRole(claims.Role, 0)
	if err != nil && resources == nil {

		ctx.Output.Status = http.StatusForbidden
		ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("Token is Error"), false, false)

		return
	}

	for _, v := range resources {
		e.AddPermissionForUser(claims.Account, v.Action, v.Method)
		//beego.Debug("policy: " + claims.Account + "  " + v.Action + "  " + v.Method)
	}
	if !a.CheckPermission(claims.Account, ctx.Request) {
		ctx.Output.Status = http.StatusUnauthorized
		ctx.Output.JSON(*errResponse.WithCode(http.StatusUnauthorized).WithMsg("Auth Fail"), false, false)
		return
	}
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(user string, r *http.Request) bool {
	method := r.Method
	path := r.URL.Path
	return a.enforcer.Enforce(user, path, method)
}

func GetToken(ctx *context.Context, secret string) (t *jwt.Token) {
	authString := ctx.Input.Header("Authorization")
	if strings.Split(authString, " ")[1] == "" {
		ctx.Output.Status = http.StatusForbidden
		ctx.Output.JSON(*errResponse.WithCode(http.StatusUnauthorized).WithMsg("AuthString invalid,Token:" + authString), false, false)
		return nil
	}
	//beego.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {

		ctx.Output.Status = http.StatusForbidden
		ctx.Output.JSON(*errResponse.WithCode(http.StatusUnauthorized).WithMsg("AuthString invalid,Token:" + authString), false, false)
		return nil
	}
	tokenString := kv[1]
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				ctx.Output.Status = http.StatusForbidden
				ctx.Output.JSON(*errResponse.WithCode(http.StatusUnauthorized).WithMsg("That's not even a token"), false, false)
				return nil
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				ctx.Output.Status = http.StatusForbidden
				ctx.Output.JSON(*errResponse.WithCode(http.StatusUnauthorized).WithMsg("Token is either expired or not active yet"), false, false)
				return
			} else {
				ctx.Output.Status = http.StatusForbidden
				ctx.Output.JSON(*errResponse.WithCode(http.StatusUnauthorized).WithMsg("Couldn‘t handle this token"), false, false)

				return nil
			}
		} else {
			ctx.Output.Status = http.StatusForbidden
			ctx.Output.JSON(*errResponse.WithCode(http.StatusUnauthorized).WithMsg("Parse token is error"), false, false)

			return nil
		}
	}
	if !token.Valid {

		ctx.Output.Status = http.StatusForbidden
		ctx.Output.JSON(*errResponse.WithCode(http.StatusUnauthorized).WithMsg("Token invalid:" + tokenString), false, false)

		return nil
	}
	//beego.Debug("Token:", token)

	return token
}
