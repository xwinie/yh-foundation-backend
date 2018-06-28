package filter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
	"yh-foundation-backend/auth"
	"yh-foundation-backend/cores"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func APISecretAuth(timeout int) beego.FilterFunc {
	return func(ctx *context.Context) {
		var errResponse cores.Entity
		if ctx.Input.Header("appid") == "" {

			ctx.Output.Status = http.StatusForbidden
			ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("miss AppId header"), false, false)

			return
		}

		client, _ := auth.GetClientService(ctx.Input.Header("appid"))
		appSecret := client.Secret
		timeoutDuration := 10 * time.Second
		cores.GlobalCaches.Put(ctx.Input.Header("appid"), client.VerifySecret, timeoutDuration)
		if appSecret == "" {

			ctx.Output.Status = http.StatusForbidden
			ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("not exist this AppId"), false, false)

			return
		}
		if ctx.Input.Header("signature") == "" {

			ctx.Output.Status = http.StatusForbidden
			ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("miss signature header"), false, false)

			return
		}
		if ctx.Input.Header("timestamp") == "" {
			ctx.Output.Status = http.StatusForbidden
			ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("miss timestamp header"), false, false)

			return
		}
		u, err := time.Parse("2006-01-02 15:04:05", ctx.Input.Header("timestamp"))
		if err != nil {
			ctx.Output.Status = http.StatusForbidden
			ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("timestamp format is error, should 2006-01-02 15:04:05"), false, false)

			return
		}
		t := time.Now()
		if t.Sub(u).Seconds() > float64(timeout) {
			ctx.Output.Status = http.StatusForbidden
			ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("timeout! the request time is long ago, please try again"), false, false)

			return
		}

		ctx.Request.ParseForm()
		var RequestURL string
		if ctx.Input.IsGet() && len(ctx.Input.URI()) == 0 && len(ctx.Request.Form) != 0 {
			RequestURL = ctx.Input.URL() + "?" + ctx.Request.Form.Encode()
		} else if len(ctx.Input.URI()) == 0 {
			RequestURL = ctx.Input.URL()
		} else {
			RequestURL = ctx.Input.URI()
		}

		clientSignature := ctx.Input.Header("signature")

		serviceSignature := Signature(appSecret, ctx.Input.Method(), ctx.Input.RequestBody, RequestURL, ctx.Input.Header("timestamp"))
		if clientSignature != serviceSignature {
			ctx.Output.Status = http.StatusForbidden
			ctx.Output.JSON(*errResponse.WithCode(http.StatusForbidden).WithMsg("Signature Failed"), false, false)

		}
	}
}

// Signature used to generate signature with the appsecret/method/params/RequestURI
func Signature(appSecret, method string, body []byte, RequestURL string, timestamp string) (result string) {
	beego.Debug("Signature", "appSecret:%s method:%s  params:%s RequestURL:%s timestamp:%s", appSecret, method, string(body), RequestURL, timestamp)

	stringToSign := fmt.Sprintf("%v\n%v\n%v\n%v\n", method, string(body), RequestURL, timestamp)

	sha256 := sha256.New
	hash := hmac.New(sha256, []byte(appSecret))
	hash.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
