package cores

import "net/http"

//Entity 结构体
type Entity struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (r Entity) New(newCode int, newMsg string) Entity {
	r.Msg = newMsg
	r.Code = newCode
	return r
}

func (r *Entity) WithMsg(newMsg string) *Entity {
	r.Msg = newMsg
	return r
}

func (r *Entity) WithAttachMsg(newMsg string) *Entity {
	r.Msg = r.Msg + newMsg
	return r
}
func (r *Entity) WithCode(newCode int) *Entity {
	r.Code = newCode
	return r
}

type ResponseEntity struct {
	StatusCode int
	Data       interface{}
}

func (r *ResponseEntity) NewBuild(StatusCode int, Data interface{}) *ResponseEntity {
	r.StatusCode = StatusCode
	r.Data = Data
	return r
}

func (r *ResponseEntity) Build(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusOK
	r.Data = Data
	return r
}

func (r *ResponseEntity) BuildError(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusBadRequest
	r.Data = Data
	return r
}

func (r *ResponseEntity) BuildFormatError(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusNotAcceptable
	r.Data = Data
	return r
}
func (r *ResponseEntity) BuildPostAndPut(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusCreated
	r.Data = Data
	return r
}

func (r *ResponseEntity) BuildDelete(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusNoContent
	r.Data = Data
	return r
}

func (r *ResponseEntity) BuildDeleteGone(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusGone
	r.Data = Data
	return r
}

func BuildEntity(newCode int, newMsg string) *Entity {
	return &Entity{newCode, newMsg}
}
