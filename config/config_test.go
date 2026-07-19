package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadValidConfig(t *testing.T) {
	dir := t.TempDir()
	cfg := Config{
		App: AppConfig{
			Debug:        true,
			Port:         "8080",
			AllowOrigins: []string{"http://localhost:3000"},
		},
		DB: DBConfig{
			Driver:   "mysql",
			Host:     "127.0.0.1",
			Port:     "3306",
			Username: "root",
			Password: "123456",
			Database: "testdb",
			Charset:  "utf8mb4",
			Prefix:   "w_",
		},
		View: ViewConfig{
			RootPath: "view",
		},
	}
	writeConfig(t, dir, cfg)

	origRoot := RootPath
	RootPath = dir
	defer func() { RootPath = origRoot }()

	result := Load("")

	if result.App.Port != "8080" {
		t.Errorf("Port = %q, want '8080'", result.App.Port)
	}
	if result.App.Debug != true {
		t.Error("Debug should be true")
	}
	if result.DB.Driver != "mysql" {
		t.Errorf("Driver = %q, want 'mysql'", result.DB.Driver)
	}
	if result.DB.Host != "127.0.0.1" {
		t.Errorf("Host = %q, want '127.0.0.1'", result.DB.Host)
	}
	if result.DB.Prefix != "w_" {
		t.Errorf("Prefix = %q, want 'w_'", result.DB.Prefix)
	}
	if len(result.App.AllowOrigins) != 1 || result.App.AllowOrigins[0] != "http://localhost:3000" {
		t.Errorf("AllowOrigins = %v, want [http://localhost:3000]", result.App.AllowOrigins)
	}
}

func TestLoadSetsDefaults(t *testing.T) {
	dir := t.TempDir()
	cfg := Config{
		App: AppConfig{
			Port: "9090",
		},
	}
	writeConfig(t, dir, cfg)

	origRoot := RootPath
	RootPath = dir
	defer func() { RootPath = origRoot }()

	result := Load("")

	if result.App.DefaultControllerName != "index" {
		t.Errorf("DefaultControllerName = %q, want 'index'", result.App.DefaultControllerName)
	}
	if result.App.DefaultActionName != "index" {
		t.Errorf("DefaultActionName = %q, want 'index'", result.App.DefaultActionName)
	}
	if result.App.DefaultAppName != "index" {
		t.Errorf("DefaultAppName = %q, want 'index'", result.App.DefaultAppName)
	}
	if result.View.RootPath != "view" {
		t.Errorf("View.RootPath = %q, want 'view'", result.View.RootPath)
	}
}

func TestLoadPreservesExistingValues(t *testing.T) {
	dir := t.TempDir()
	cfg := Config{
		App: AppConfig{
			DefaultControllerName: "home",
			DefaultActionName:     "show",
			DefaultAppName:        "admin",
		},
		View: ViewConfig{
			RootPath: "templates",
		},
	}
	writeConfig(t, dir, cfg)

	origRoot := RootPath
	RootPath = dir
	defer func() { RootPath = origRoot }()

	result := Load("")

	if result.App.DefaultControllerName != "home" {
		t.Errorf("DefaultControllerName = %q, want 'home'", result.App.DefaultControllerName)
	}
	if result.App.DefaultActionName != "show" {
		t.Errorf("DefaultActionName = %q, want 'show'", result.App.DefaultActionName)
	}
	if result.App.DefaultAppName != "admin" {
		t.Errorf("DefaultAppName = %q, want 'admin'", result.App.DefaultAppName)
	}
	if result.View.RootPath != "templates" {
		t.Errorf("View.RootPath = %q, want 'templates'", result.View.RootPath)
	}
}

func TestLoadSetsGlobalConfig(t *testing.T) {
	dir := t.TempDir()
	cfg := Config{App: AppConfig{Port: "7777"}}
	writeConfig(t, dir, cfg)

	origRoot := RootPath
	RootPath = dir
	defer func() { RootPath = origRoot }()

	result := Load("")

	if GlobalConfig != result {
		t.Error("GlobalConfig should be set to the loaded config")
	}
	if GlobalConfig.App.Port != "7777" {
		t.Errorf("GlobalConfig.App.Port = %q, want '7777'", GlobalConfig.App.Port)
	}
}

func TestLoadWithSubDir(t *testing.T) {
	dir := t.TempDir()
	subDir := filepath.Join(dir, "config")
	os.MkdirAll(subDir, 0755)

	cfg := Config{App: AppConfig{Port: "5555"}}
	writeConfig(t, subDir, cfg)

	origRoot := RootPath
	RootPath = dir
	defer func() { RootPath = origRoot }()

	result := Load("config")

	if result.App.Port != "5555" {
		t.Errorf("Port = %q, want '5555'", result.App.Port)
	}
}

func TestLoadPanicOnMissingFile(t *testing.T) {
	dir := t.TempDir()
	origRoot := RootPath
	RootPath = dir
	defer func() { RootPath = origRoot }()

	defer func() {
		r := recover()
		if r == nil {
			t.Error("Load() should panic when config file is missing")
		}
		msg, ok := r.(string)
		if !ok {
			t.Errorf("panic message should be string, got %T", r)
		}
		if !contains(msg, "配置文件打开失败") {
			t.Errorf("panic message = %q, should contain '配置文件打开失败'", msg)
		}
	}()

	Load("nonexistent")
}

func TestLoadPanicOnInvalidJSON(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "config.json"), []byte("{invalid json}"), 0644)

	origRoot := RootPath
	RootPath = dir
	defer func() { RootPath = origRoot }()

	defer func() {
		r := recover()
		if r == nil {
			t.Error("Load() should panic on invalid JSON")
		}
		msg, ok := r.(string)
		if !ok {
			t.Errorf("panic message should be string, got %T", r)
		}
		if !contains(msg, "配置加载失败") {
			t.Errorf("panic message = %q, should contain '配置加载失败'", msg)
		}
	}()

	Load("")
}

func TestLoadPanicOnEmptyDir(t *testing.T) {
	dir := t.TempDir()
	origRoot := RootPath
	RootPath = dir
	defer func() { RootPath = origRoot }()

	defer func() {
		r := recover()
		if r == nil {
			t.Error("Load() should panic when config dir is empty and no config.json exists")
		}
	}()

	Load("")
}

func TestDS(t *testing.T) {
	if DS == "" {
		t.Error("DS should not be empty")
	}
	if DS != "/" && DS != "\\" {
		t.Errorf("DS = %q, unexpected value", DS)
	}
}

func TestRootPath(t *testing.T) {
	if RootPath == "" {
		t.Error("RootPath should not be empty")
	}
}

// --- helpers ---

func writeConfig(t *testing.T, dir string, cfg Config) {
	t.Helper()
	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("failed to marshal config: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "config.json"), data, 0644); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
