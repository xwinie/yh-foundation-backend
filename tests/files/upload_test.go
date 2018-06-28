package files

import (
	"testing"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	_ "yh-foundation-backend/tests"
	"yh-foundation-backend/filter"
	"yh-foundation-backend/tests"
	"yh-foundation-backend/auth"
	"time"
	"github.com/mingzhehao/goutils/filetool"
	"net/http/httptest"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"os"
)

func TestUpload(t *testing.T) {

	method := "POST"
	RequestURL := "/v1/upload"

	fileName := "upload.jpg"
	if !filetool.IsExist(fileName) {
		filetool.WriteStringToFile(fileName, "测试upload")
	}
	tokenSting := tests.GetToken()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	b := httplib.Post(RequestURL)
	b.Header("appid", "app1")
	b.Header("timestamp", timestamp)
	b.Header("Authorization", "Bearer "+tokenSting)
	b.PostFile("uploadFile", fileName)
	signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)
	b.Header("signature", signature)
	b.DoRequest()
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, b.GetRequest())

	beego.Trace("testing", "TestUpload", "Code[%d]\n%s", w.Code, w.Body.String())
	Convey("Subject: 文件上传\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)

			Convey("download file should be exists", func() {
				type Entity struct {
					FileName string
					Url      string
				}
				var Response []Entity
				json.Unmarshal(w.Body.Bytes(), &Response)
				beego.Trace("testing", "TestDownload", "%s", Response[0].Url)
				method := "GET"
				RequestURL := "/" + Response[0].Url
				b := httplib.Get(RequestURL)
				b.Header("appid", "app1")
				b.Header("timestamp", timestamp)
				b.Header("Authorization", "Bearer "+tokenSting)
				signature := filter.Signature(auth.GetSecretService("app1"), method, nil, RequestURL,timestamp)
				b.Header("signature", signature)
				b.DoRequest()
				w := httptest.NewRecorder()
				beego.BeeApp.Handlers.ServeHTTP(w, b.GetRequest())
				if filetool.IsExist("download" + Response[0].FileName) {
					os.Remove("download" + Response[0].FileName)
				}
				filetool.WriteBytesToFile("download"+Response[0].FileName, w.Body.Bytes())
				So(w.Code, ShouldEqual, 200)
				So(filetool.IsExist("download"+Response[0].FileName), ShouldEqual, true)
			})
		})
	})
}
