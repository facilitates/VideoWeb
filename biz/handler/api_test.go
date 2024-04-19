package handler

import (
	"bytes"
	"testing"
	"videoweb/biz/handler/videoweb"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestPerformRequest(t *testing.T) {
	h := server.Default()
	h.POST("/user/register", videoweb.Register)
	jsonStr := `{"username":"test","password":"123456"}`
	w := ut.PerformRequest(h.Engine, "POST", "/user/register", &ut.Body{Body: bytes.NewBufferString(jsonStr), Len: len(jsonStr)})
	resp := w.Result()
	assert.DeepEqual(t, "{\"base\":{\"code\":200,\"msg\":\"注册成功\"}}", string(resp.Body()))
}
