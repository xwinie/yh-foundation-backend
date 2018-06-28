package tests

import (
	"yh-foundation-backend/cores"
	"encoding/json"
	"bytes"
	"github.com/astaxie/beego"
	"yh-foundation-backend/filter"
	"net/http"
	"time"
	"net/http/httptest"
	"yh-foundation-backend/auth"
)

func GetToken() string {
	values := map[string]string{"UserName": "12345", "Password": cores.Md5(cores.Sha1("12345") + cores.Sha1("Password"))}
	jsonValue, _ := json.Marshal(values)

	method := "POST"
	RequestURL := "/v1/login"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signature := filter.Signature(auth.GetSecretService("app1"), method, bytes.NewBuffer(jsonValue).Bytes(), RequestURL,timestamp)

	r, _ := http.NewRequest(method, RequestURL, bytes.NewBuffer(jsonValue))
	r.Header.Set("appid", "app1")
	r.Header.Set("timestamp", timestamp)
	r.Header.Set("signature", signature)
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("testing", "GetToken", "Code[%d]\n%s", w.Code, w.Body.String())
	type Response struct {
		Account string
		Name    string
		Token   string
		Exp     int64
	}
	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)
	return response.Token
}
