package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/wycto/weigo/config"
)

func TestCORS_AllowConfiguredOrigin(t *testing.T) {
	// 设置测试配置
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{
			Debug:        false,
			AllowOrigins: []string{"http://example.com", "http://test.com"},
		},
	}

	e := echo.New()
	e.Use(CORS())

	// 测试允许的来源
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "http://example.com")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(200, "ok")
	}

	if err := CORS()(h)(c); err != nil {
		t.Fatalf("CORS() error: %v", err)
	}

	if rec.Header().Get("Access-Control-Allow-Origin") != "http://example.com" {
		t.Errorf("Allow-Origin = %q, want %q", rec.Header().Get("Access-Control-Allow-Origin"), "http://example.com")
	}
}

func TestCORS_RejectUnconfiguredOrigin(t *testing.T) {
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{
			Debug:        false,
			AllowOrigins: []string{"http://example.com"},
		},
	}

	e := echo.New()
	e.Use(CORS())

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "http://evil.com")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(200, "ok")
	}

	if err := CORS()(h)(c); err != nil {
		t.Fatalf("CORS() error: %v", err)
	}

	if rec.Header().Get("Access-Control-Allow-Origin") != "" {
		t.Errorf("Allow-Origin should be empty for rejected origin, got %q", rec.Header().Get("Access-Control-Allow-Origin"))
	}
}

func TestCORS_DebugModeAllowsAll(t *testing.T) {
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{
			Debug:        true,
			AllowOrigins: nil, // 未配置白名单
		},
	}

	e := echo.New()
	e.Use(CORS())

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "http://any-origin.com")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(200, "ok")
	}

	if err := CORS()(h)(c); err != nil {
		t.Fatalf("CORS() error: %v", err)
	}

	if rec.Header().Get("Access-Control-Allow-Origin") != "http://any-origin.com" {
		t.Errorf("Debug mode should allow all origins, got %q", rec.Header().Get("Access-Control-Allow-Origin"))
	}
}

func TestCORS_DebugModeWithConfigRestricts(t *testing.T) {
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{
			Debug:        true,
			AllowOrigins: []string{"http://example.com"}, // 配置了白名单后应遵守
		},
	}

	e := echo.New()
	e.Use(CORS())

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "http://evil.com")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(200, "ok")
	}

	if err := CORS()(h)(c); err != nil {
		t.Fatalf("CORS() error: %v", err)
	}

	if rec.Header().Get("Access-Control-Allow-Origin") != "" {
		t.Errorf("Should reject unconfigured origin even in debug mode when AllowOrigins is set")
	}
}

func TestCORS_NoOriginHeader(t *testing.T) {
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{
			AllowOrigins: []string{"http://example.com"},
		},
	}

	e := echo.New()
	e.Use(CORS())

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// 不设置 Origin header
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(200, "ok")
	}

	if err := CORS()(h)(c); err != nil {
		t.Fatalf("CORS() error: %v", err)
	}

	// 没有 Origin 时不应设置 CORS 头
	if rec.Header().Get("Access-Control-Allow-Origin") != "" {
		t.Errorf("Should not set CORS headers when no Origin header present")
	}
}

func TestCORS_PreflightReturns204(t *testing.T) {
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{
			AllowOrigins: []string{"*"},
		},
	}

	e := echo.New()
	e.Use(CORS())

	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	req.Header.Set("Origin", "http://example.com")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(200, "ok")
	}

	if err := CORS()(h)(c); err != nil {
		t.Fatalf("CORS() error: %v", err)
	}

	if rec.Code != 204 {
		t.Errorf("Preflight status = %d, want 204", rec.Code)
	}
}
