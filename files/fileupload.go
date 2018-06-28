package files

import (
	"github.com/astaxie/beego"
	"regexp"
	"strings"
	. "yh-foundation-backend/cores"
	"net/http"
	"github.com/mingzhehao/goutils/filetool"
	"fmt"
	"time"
)

type FileUploadController struct {
	beego.Controller
}

type Sizer interface {
	Size() int64
}

const (
	LocalFileDir   = "static/uploads/file"
	MinFileSize    = 1       // bytes
	MaxFileSize    = 5000000 // bytes
	FileType       = "(jpg|gif|p?jpeg|(x-)?png)"
	AcceptFileType = FileType
)

var (
	imageTypes      = regexp.MustCompile(FileType)
	acceptFileTypes = regexp.MustCompile(AcceptFileType)
)

type FileInfo struct {
	Url          string `json:"url,omitempty"`
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Size         int64  `json:"size"`
	Error        string `json:"error,omitempty"`
	DeleteUrl    string `json:"deleteUrl,omitempty"`
	DeleteType   string `json:"deleteType,omitempty"`
}

func (fi *FileInfo) ValidateType() (valid bool) {
	if acceptFileTypes.MatchString(fi.Type) {
		return true
	}
	fi.Error = "FileType not allowed"
	return false
}

func (fi *FileInfo) ValidateSize() (valid bool) {
	if fi.Size < MinFileSize {
		fi.Error = "File is too small"
	} else if fi.Size > MaxFileSize {
		fi.Error = "File is too big"
	} else {
		return true
	}
	return false
}

func (c *FileUploadController) UploadCtl() {
	//beego.Trace(111,c.Ctx.Request.MultipartForm)
	// GetFiles return multi-upload files
	files, err := c.GetFiles("uploadFile")
	if err != nil {
		c.Data["json"] = BuildEntity(http.StatusNoContent, "getFile err! "+err.Error())
		c.Ctx.Output.Status = http.StatusNoContent
		return
	}
	type Entity struct {
		FileName string
		Url      string
	}
	var Response []Entity
	for i, h := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			c.Data["json"] = BuildEntity(http.StatusNoContent, "invalid file! open file error:"+err.Error())
			c.Ctx.Output.Status = http.StatusNoContent
			c.ServeJSON()
			return
		}
		ext := filetool.Ext(h.Filename)
		fi := &FileInfo{
			Name: h.Filename,
			Type: ext,
		}
		if sizeInterface, ok := file.(Sizer); ok {
			fi.Size = sizeInterface.Size()
		}
		if !fi.ValidateType() || !fi.ValidateSize() {
			c.Data["json"] = BuildEntity(http.StatusNoContent, "invalid file! Validate err:"+fi.Error)
			c.Ctx.Output.Status = http.StatusNoContent
			c.ServeJSON()
			return
		}

		fileExt := strings.TrimLeft(ext, ".")
		fileSaveName := fmt.Sprintf("%s_%s%d%s", fileExt, Md5(h.Filename), time.Now().Unix(), ext)
		filePath := fmt.Sprintf("%s/%s", LocalFileDir, fileSaveName)
		// 保存位置在 static/upload,没有文件夹要先创建
		filetool.InsureDir(LocalFileDir)

		if err3 := c.SaveToFile("uploadFile", filePath); err3 != nil {
			c.Data["json"] = BuildEntity(http.StatusNoContent, "invalid file! save file error:"+err.Error())
			c.Ctx.Output.Status = http.StatusNoContent
			c.ServeJSON()
			return
		}
		Response = append(Response, Entity{h.Filename, filePath})
	}
	c.Data["json"] = Response
	c.ServeJSON()

}
