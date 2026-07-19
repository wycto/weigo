package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/wycto/weigo/config"
)

// 测试用控制器
type TestController struct {
	Controller
}

func (c *TestController) Index() error {
	return c.Success("index", nil)
}

func (c *TestController) GetUser() error {
	id := c.GetString("id", "0")
	return c.Success("get_user", id)
}

func (c *TestController) AddMember() error {
	name := c.GetParam("name")
	return c.Success("add_member", name)
}

func TestControllerInit(t *testing.T) {
	e := echo.New()
	body := `{"name":"test","age":18}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := &Controller{}
	ctrl.Init(c)

	if ctrl.C == nil {
		t.Error("Controller.C should not be nil after Init")
	}
	if ctrl.data == nil {
		t.Error("Controller.data should be initialized")
	}
	if ctrl.body == nil {
		t.Error("Controller.body should be initialized")
	}

	// 验证 JSON body 解析
	bodyMap := ctrl.GetBody()
	if bodyMap["name"] != "test" {
		t.Errorf("body[name] = %v, want 'test'", bodyMap["name"])
	}
	if bodyMap["age"] != float64(18) {
		t.Errorf("body[age] = %v, want 18", bodyMap["age"])
	}
}

func TestControllerHas(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?key1=val1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("123")

	ctrl := &Controller{}
	ctrl.Init(c)

	if !ctrl.Has("id") {
		t.Error("Has('id') should return true (path param)")
	}
	if !ctrl.Has("key1") {
		t.Error("Has('key1') should return true (query param)")
	}
	if ctrl.Has("nonexistent") {
		t.Error("Has('nonexistent') should return false")
	}
}

func TestControllerHasAndNotEmpty(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?key1=val1&key2=0&key3=", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := &Controller{}
	ctrl.Init(c)

	if !ctrl.HasAndNotEmpty("key1") {
		t.Error("HasAndNotEmpty('key1') should return true")
	}
	if ctrl.HasAndNotEmpty("key2") {
		t.Error("HasAndNotEmpty('key2') should return false (value is '0')")
	}
	if ctrl.HasAndNotEmpty("key3") {
		t.Error("HasAndNotEmpty('key3') should return false (empty value)")
	}
}

func TestControllerGetString(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?q=search", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("42")

	ctrl := &Controller{}
	ctrl.Init(c)

	// 路径参数优先
	if v := ctrl.GetString("id", ""); v != "42" {
		t.Errorf("GetString('id') = %q, want '42'", v)
	}

	// query 参数
	if v := ctrl.GetString("q", ""); v != "search" {
		t.Errorf("GetString('q') = %q, want 'search'", v)
	}

	// 默认值
	if v := ctrl.GetString("missing", "default"); v != "default" {
		t.Errorf("GetString('missing') = %q, want 'default'", v)
	}
}

func TestControllerGetInt(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?count=abc", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("42")

	ctrl := &Controller{}
	ctrl.Init(c)

	// 路径参数
	if v := ctrl.GetInt("id", 0); v != 42 {
		t.Errorf("GetInt('id') = %d, want 42", v)
	}

	// 非数字返回默认值
	if v := ctrl.GetInt("count", 10); v != 10 {
		t.Errorf("GetInt('count') = %d, want 10 (default)", v)
	}

	// 缺失返回默认值
	if v := ctrl.GetInt("missing", 99); v != 99 {
		t.Errorf("GetInt('missing') = %d, want 99", v)
	}
}

func TestControllerSuccess(t *testing.T) {
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{DefaultAppName: "index"},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := &Controller{}
	ctrl.Init(c)

	err := ctrl.Success("ok", map[string]string{"id": "1"})
	if err != nil {
		t.Fatalf("Success() error: %v", err)
	}

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)

	if resp["code"] != float64(0) {
		t.Errorf("code = %v, want 0", resp["code"])
	}
	if resp["msg"] != "ok" {
		t.Errorf("msg = %v, want 'ok'", resp["msg"])
	}
}

func TestControllerError(t *testing.T) {
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{DefaultAppName: "index"},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := &Controller{}
	ctrl.Init(c)

	err := ctrl.Error("failed", nil)
	if err != nil {
		t.Fatalf("Error() error: %v", err)
	}

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)

	if resp["code"] != float64(1) {
		t.Errorf("code = %v, want 1", resp["code"])
	}
}

func TestControllerAssign(t *testing.T) {
	ctrl := &Controller{}
	ctrl.data = make(map[string]interface{})

	ctrl.Assign("title", "Hello")
	ctrl.Assign("count", 42)

	if ctrl.data["title"] != "Hello" {
		t.Errorf("data[title] = %v, want 'Hello'", ctrl.data["title"])
	}
	if ctrl.data["count"] != 42 {
		t.Errorf("data[count] = %v, want 42", ctrl.data["count"])
	}
}

func TestControllerGetParamPriority(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?key=query_val", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("key")
	c.SetParamValues("path_val")

	ctrl := &Controller{}
	ctrl.Init(c)

	// 路径参数优先于 query
	if v := ctrl.GetParam("key"); v != "path_val" {
		t.Errorf("GetParam('key') = %q, want 'path_val' (path param should have priority)", v)
	}
}

func TestControllerGetFloat64(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?price=3.14", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("score")
	c.SetParamValues("99.5")

	ctrl := &Controller{}
	ctrl.Init(c)

	// 路径参数
	if v := ctrl.GetFloat64("score", 0); v != 99.5 {
		t.Errorf("GetFloat64('score') = %v, want 99.5", v)
	}

	// query 参数
	if v := ctrl.GetFloat64("price", 0); v != 3.14 {
		t.Errorf("GetFloat64('price') = %v, want 3.14", v)
	}

	// 非数字返回默认值
	if v := ctrl.GetFloat64("missing", 1.5); v != 1.5 {
		t.Errorf("GetFloat64('missing') = %v, want 1.5", v)
	}
}

func TestControllerGetFloat64FromBody(t *testing.T) {
	e := echo.New()
	body := `{"price":9.99}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := &Controller{}
	ctrl.Init(c)

	if v := ctrl.GetFloat64("price", 0); v != 9.99 {
		t.Errorf("GetFloat64('price') from body = %v, want 9.99", v)
	}
}

func TestControllerGetBool(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		pathVal  string
		defaultV bool
		want     bool
	}{
		{"true string", "active=true", "", false, true},
		{"1 string", "active=1", "", false, true},
		{"yes string", "active=yes", "", false, true},
		{"false string", "active=false", "", true, false},
		{"0 string", "active=0", "", true, false},
		{"empty string", "active=", "", true, true}, // 空值与缺失等价，返回默认值
		{"missing", "", "", true, true},
		{"path param", "", "1", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			url := "/?" + tt.query
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if tt.pathVal != "" {
				c.SetParamNames("active")
				c.SetParamValues(tt.pathVal)
			}

			ctrl := &Controller{}
			ctrl.Init(c)

			if v := ctrl.GetBool("active", tt.defaultV); v != tt.want {
				t.Errorf("GetBool('active') = %v, want %v", v, tt.want)
			}
		})
	}
}

func TestControllerGetBoolFromBody(t *testing.T) {
	e := echo.New()
	body := `{"enabled":true,"disabled":false}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := &Controller{}
	ctrl.Init(c)

	if v := ctrl.GetBool("enabled", false); v != true {
		t.Errorf("GetBool('enabled') = %v, want true", v)
	}
	if v := ctrl.GetBool("disabled", true); v != false {
		t.Errorf("GetBool('disabled') = %v, want false", v)
	}
}
