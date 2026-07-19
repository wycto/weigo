package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := Success(c, "ok", map[string]string{"key": "value"})
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

func TestError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := Error(c, "failed", nil)
	if err != nil {
		t.Fatalf("Error() error: %v", err)
	}

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)

	if resp["code"] != float64(1) {
		t.Errorf("code = %v, want 1", resp["code"])
	}
	if resp["msg"] != "failed" {
		t.Errorf("msg = %v, want 'failed'", resp["msg"])
	}
}

func TestResponseMessage(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ResponseMessage(c, ParamsMissing, nil)
	if err != nil {
		t.Fatalf("ResponseMessage() error: %v", err)
	}

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)

	if resp["code"] != float64(200101) {
		t.Errorf("code = %v, want 200101", resp["code"])
	}
	if resp["msg"] != "参数缺失" {
		t.Errorf("msg = %v, want '参数缺失'", resp["msg"])
	}
}

func TestMessageConstants(t *testing.T) {
	if ParamsMissing.Code != 200101 {
		t.Errorf("ParamsMissing.Code = %d, want 200101", ParamsMissing.Code)
	}
	if DataNone.Code != 200201 {
		t.Errorf("DataNone.Code = %d, want 200201", DataNone.Code)
	}
	if SysError.Code != 200501 {
		t.Errorf("SysError.Code = %d, want 200501", SysError.Code)
	}
}

func TestErrorWithStatus(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ErrorWithStatus(c, http.StatusNotFound, 40400, "not found", nil)
	if err != nil {
		t.Fatalf("ErrorWithStatus() error: %v", err)
	}

	if rec.Code != http.StatusNotFound {
		t.Errorf("HTTP status = %d, want %d", rec.Code, http.StatusNotFound)
	}

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)

	if resp["code"] != float64(40400) {
		t.Errorf("code = %v, want 40400", resp["code"])
	}
}

func TestBadRequest(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := BadRequest(c, "参数错误")
	if err != nil {
		t.Fatalf("BadRequest() error: %v", err)
	}

	if rec.Code != http.StatusBadRequest {
		t.Errorf("HTTP status = %d, want %d", rec.Code, http.StatusBadRequest)
	}

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)

	if resp["msg"] != "参数错误" {
		t.Errorf("msg = %v, want '参数错误'", resp["msg"])
	}
}

func TestUnauthorized(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := Unauthorized(c, "未登录")
	if err != nil {
		t.Fatalf("Unauthorized() error: %v", err)
	}

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("HTTP status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

func TestForbidden(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := Forbidden(c, "无权限")
	if err != nil {
		t.Fatalf("Forbidden() error: %v", err)
	}

	if rec.Code != http.StatusForbidden {
		t.Errorf("HTTP status = %d, want %d", rec.Code, http.StatusForbidden)
	}
}

func TestNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := NotFound(c, "数据不存在")
	if err != nil {
		t.Fatalf("NotFound() error: %v", err)
	}

	if rec.Code != http.StatusNotFound {
		t.Errorf("HTTP status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func TestServerError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ServerError(c, "服务器开小差了")
	if err != nil {
		t.Fatalf("ServerError() error: %v", err)
	}

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("HTTP status = %d, want %d", rec.Code, http.StatusInternalServerError)
	}
}
