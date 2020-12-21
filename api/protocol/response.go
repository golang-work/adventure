package protocol

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-work/adventure/support"
	"net/http"
)

type response struct {
	ctx       *gin.Context
	stateCode int
	sendBody  gin.H
}

func Response(ctx *gin.Context) *response {
	return &response{
		ctx:       ctx,
		stateCode: http.StatusOK,
		sendBody: gin.H{
			"code": 10200,
			"data": nil,
		},
	}
}

func (r *response) Result(data interface{}) *response {
	r.sendBody["data"] = data
	return r
}

func (r *response) Success() *response {
	r.sendBody["data"] = 1
	return r
}

func (r *response) Fatal(data string) *response {
	r.stateCode = http.StatusBadRequest
	r.sendBody["code"] = 10400
	r.sendBody["data"] = data
	return r
}

func (r *response) ValidaFatal(data string) *response {
	r.stateCode = http.StatusUnprocessableEntity
	r.sendBody["code"] = 10422
	r.sendBody["data"] = data
	return r
}

func (r *response) Created(data interface{}) *response {
	r.stateCode = http.StatusCreated
	r.sendBody["code"] = 10201
	r.sendBody["data"] = data
	return r
}

func (r *response) Extra(field gin.H) *response {
	for k, v := range field {
		r.sendBody[k] = v
	}
	return r
}

func (r *response) Abort(e interface{}) *response {
	var abort support.Abort
	switch v := e.(type) {
	case support.Abort:
		abort = v
	case string:
		abort = support.Throw(v)
	}
	r.stateCode = support.LevelMap[abort.Level]
	r.sendBody["code"] = abort.Code
	r.sendBody["data"] = abort.Message
	return r
}

func (r *response) Json() {
	r.ctx.JSON(r.stateCode, r.sendBody)
}
